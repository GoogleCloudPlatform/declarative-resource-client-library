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
package appengine

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Application) validate() error {

	if !dcl.IsEmptyValueIndirect(r.Iap) {
		if err := r.Iap.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FeatureSettings) {
		if err := r.FeatureSettings.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ApplicationDispatchRules) validate() error {
	return nil
}
func (r *ApplicationIap) validate() error {
	return nil
}
func (r *ApplicationFeatureSettings) validate() error {
	return nil
}

func applicationGetURL(userBasePath string, r *Application) (string, error) {
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("apps/{{name}}", "https://appengine.googleapis.com/v1/", userBasePath, params), nil
}

func applicationCreateURL(userBasePath string, _ ...interface{}) (string, error) {
	return dcl.URL("apps", "https://appengine.googleapis.com/v1/", userBasePath, map[string]interface{}{}), nil
}

// applicationApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type applicationApiOperation interface {
	do(context.Context, *Application, *Client) error
}

// newUpdateApplicationPatchApplicationRequest creates a request for an
// Application resource's PatchApplication update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateApplicationPatchApplicationRequest(ctx context.Context, f *Application, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandApplicationDispatchRulesSlice(c, f.DispatchRules); err != nil {
		return nil, fmt.Errorf("error expanding DispatchRules into dispatchRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["dispatchRules"] = v
	}
	if v := f.AuthDomain; !dcl.IsEmptyValueIndirect(v) {
		req["authDomain"] = v
	}
	if v := f.DefaultCookieExpiration; !dcl.IsEmptyValueIndirect(v) {
		req["defaultCookieExpiration"] = v
	}
	if v := f.ServingStatus; !dcl.IsEmptyValueIndirect(v) {
		req["servingStatus"] = v
	}
	if v := f.GcrDomain; !dcl.IsEmptyValueIndirect(v) {
		req["gcrDomain"] = v
	}
	if v := f.DatabaseType; !dcl.IsEmptyValueIndirect(v) {
		req["databaseType"] = v
	}
	if v, err := expandApplicationFeatureSettings(c, f.FeatureSettings); err != nil {
		return nil, fmt.Errorf("error expanding FeatureSettings into featureSettings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["featureSettings"] = v
	}
	return req, nil
}

// marshalUpdateApplicationPatchApplicationRequest converts the update into
// the final JSON request body.
func marshalUpdateApplicationPatchApplicationRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateApplicationPatchApplicationOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateApplicationPatchApplicationOperation) do(ctx context.Context, r *Application, c *Client) error {
	_, err := c.GetApplication(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchApplication")
	if err != nil {
		return err
	}
	mask := dcl.SnakeCaseUpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateApplicationPatchApplicationRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateApplicationPatchApplicationRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://appengine.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createApplicationOperation struct {
	response map[string]interface{}
}

func (op *createApplicationOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createApplicationOperation) do(ctx context.Context, r *Application, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	u, err := applicationCreateURL(c.Config.BasePath)

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
	if err := o.Wait(ctx, c.Config, "https://appengine.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetApplication(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getApplicationRaw(ctx context.Context, r *Application) ([]byte, error) {

	u, err := applicationGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) applicationDiffsForRawDesired(ctx context.Context, rawDesired *Application, opts ...dcl.ApplyOption) (initial, desired *Application, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Application
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Application); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Application, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetApplication(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Application resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Application resource: %v", err)
		}
		c.Config.Logger.Info("Found that Application resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeApplicationDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Application: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Application: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeApplicationInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Application: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeApplicationDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Application: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffApplication(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeApplicationInitialState(rawInitial, rawDesired *Application) (*Application, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeApplicationDesiredState(rawDesired, rawInitial *Application, opts ...dcl.ApplyOption) (*Application, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Iap = canonicalizeApplicationIap(rawDesired.Iap, nil, opts...)
		rawDesired.FeatureSettings = canonicalizeApplicationFeatureSettings(rawDesired.FeatureSettings, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Application{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.DispatchRules) {
		canonicalDesired.DispatchRules = rawInitial.DispatchRules
	} else {
		canonicalDesired.DispatchRules = rawDesired.DispatchRules
	}
	if dcl.StringCanonicalize(rawDesired.AuthDomain, rawInitial.AuthDomain) {
		canonicalDesired.AuthDomain = rawInitial.AuthDomain
	} else {
		canonicalDesired.AuthDomain = rawDesired.AuthDomain
	}
	if dcl.StringCanonicalize(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	if dcl.StringCanonicalize(rawDesired.DefaultCookieExpiration, rawInitial.DefaultCookieExpiration) {
		canonicalDesired.DefaultCookieExpiration = rawInitial.DefaultCookieExpiration
	} else {
		canonicalDesired.DefaultCookieExpiration = rawDesired.DefaultCookieExpiration
	}
	if dcl.IsZeroValue(rawDesired.ServingStatus) {
		canonicalDesired.ServingStatus = rawInitial.ServingStatus
	} else {
		canonicalDesired.ServingStatus = rawDesired.ServingStatus
	}
	canonicalDesired.Iap = canonicalizeApplicationIap(rawDesired.Iap, rawInitial.Iap, opts...)
	if dcl.StringCanonicalize(rawDesired.GcrDomain, rawInitial.GcrDomain) {
		canonicalDesired.GcrDomain = rawInitial.GcrDomain
	} else {
		canonicalDesired.GcrDomain = rawDesired.GcrDomain
	}
	if dcl.IsZeroValue(rawDesired.DatabaseType) {
		canonicalDesired.DatabaseType = rawInitial.DatabaseType
	} else {
		canonicalDesired.DatabaseType = rawDesired.DatabaseType
	}
	canonicalDesired.FeatureSettings = canonicalizeApplicationFeatureSettings(rawDesired.FeatureSettings, rawInitial.FeatureSettings, opts...)

	return canonicalDesired, nil
}

func canonicalizeApplicationNewState(c *Client, rawNew, rawDesired *Application) (*Application, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DispatchRules) && dcl.IsEmptyValueIndirect(rawDesired.DispatchRules) {
		rawNew.DispatchRules = rawDesired.DispatchRules
	} else {
		rawNew.DispatchRules = canonicalizeNewApplicationDispatchRulesSlice(c, rawDesired.DispatchRules, rawNew.DispatchRules)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AuthDomain) && dcl.IsEmptyValueIndirect(rawDesired.AuthDomain) {
		rawNew.AuthDomain = rawDesired.AuthDomain
	} else {
		if dcl.StringCanonicalize(rawDesired.AuthDomain, rawNew.AuthDomain) {
			rawNew.AuthDomain = rawDesired.AuthDomain
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Location) && dcl.IsEmptyValueIndirect(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
		if dcl.StringCanonicalize(rawDesired.Location, rawNew.Location) {
			rawNew.Location = rawDesired.Location
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CodeBucket) && dcl.IsEmptyValueIndirect(rawDesired.CodeBucket) {
		rawNew.CodeBucket = rawDesired.CodeBucket
	} else {
		if dcl.StringCanonicalize(rawDesired.CodeBucket, rawNew.CodeBucket) {
			rawNew.CodeBucket = rawDesired.CodeBucket
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultCookieExpiration) && dcl.IsEmptyValueIndirect(rawDesired.DefaultCookieExpiration) {
		rawNew.DefaultCookieExpiration = rawDesired.DefaultCookieExpiration
	} else {
		if dcl.StringCanonicalize(rawDesired.DefaultCookieExpiration, rawNew.DefaultCookieExpiration) {
			rawNew.DefaultCookieExpiration = rawDesired.DefaultCookieExpiration
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServingStatus) && dcl.IsEmptyValueIndirect(rawDesired.ServingStatus) {
		rawNew.ServingStatus = rawDesired.ServingStatus
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultHostname) && dcl.IsEmptyValueIndirect(rawDesired.DefaultHostname) {
		rawNew.DefaultHostname = rawDesired.DefaultHostname
	} else {
		if dcl.StringCanonicalize(rawDesired.DefaultHostname, rawNew.DefaultHostname) {
			rawNew.DefaultHostname = rawDesired.DefaultHostname
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultBucket) && dcl.IsEmptyValueIndirect(rawDesired.DefaultBucket) {
		rawNew.DefaultBucket = rawDesired.DefaultBucket
	} else {
		if dcl.StringCanonicalize(rawDesired.DefaultBucket, rawNew.DefaultBucket) {
			rawNew.DefaultBucket = rawDesired.DefaultBucket
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Iap) && dcl.IsEmptyValueIndirect(rawDesired.Iap) {
		rawNew.Iap = rawDesired.Iap
	} else {
		rawNew.Iap = canonicalizeNewApplicationIap(c, rawDesired.Iap, rawNew.Iap)
	}

	if dcl.IsEmptyValueIndirect(rawNew.GcrDomain) && dcl.IsEmptyValueIndirect(rawDesired.GcrDomain) {
		rawNew.GcrDomain = rawDesired.GcrDomain
	} else {
		if dcl.StringCanonicalize(rawDesired.GcrDomain, rawNew.GcrDomain) {
			rawNew.GcrDomain = rawDesired.GcrDomain
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DatabaseType) && dcl.IsEmptyValueIndirect(rawDesired.DatabaseType) {
		rawNew.DatabaseType = rawDesired.DatabaseType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.FeatureSettings) && dcl.IsEmptyValueIndirect(rawDesired.FeatureSettings) {
		rawNew.FeatureSettings = rawDesired.FeatureSettings
	} else {
		rawNew.FeatureSettings = canonicalizeNewApplicationFeatureSettings(c, rawDesired.FeatureSettings, rawNew.FeatureSettings)
	}

	return rawNew, nil
}

func canonicalizeApplicationDispatchRules(des, initial *ApplicationDispatchRules, opts ...dcl.ApplyOption) *ApplicationDispatchRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ApplicationDispatchRules{}

	if dcl.StringCanonicalize(des.Domain, initial.Domain) || dcl.IsZeroValue(des.Domain) {
		cDes.Domain = initial.Domain
	} else {
		cDes.Domain = des.Domain
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		cDes.Service = initial.Service
	} else {
		cDes.Service = des.Service
	}

	return cDes
}

func canonicalizeNewApplicationDispatchRules(c *Client, des, nw *ApplicationDispatchRules) *ApplicationDispatchRules {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Domain, nw.Domain) {
		nw.Domain = des.Domain
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}

	return nw
}

func canonicalizeNewApplicationDispatchRulesSet(c *Client, des, nw []ApplicationDispatchRules) []ApplicationDispatchRules {
	if des == nil {
		return nw
	}
	var reorderedNew []ApplicationDispatchRules
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareApplicationDispatchRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewApplicationDispatchRulesSlice(c *Client, des, nw []ApplicationDispatchRules) []ApplicationDispatchRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ApplicationDispatchRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewApplicationDispatchRules(c, &d, &n))
	}

	return items
}

func canonicalizeApplicationIap(des, initial *ApplicationIap, opts ...dcl.ApplyOption) *ApplicationIap {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ApplicationIap{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}
	if dcl.StringCanonicalize(des.OAuth2ClientId, initial.OAuth2ClientId) || dcl.IsZeroValue(des.OAuth2ClientId) {
		cDes.OAuth2ClientId = initial.OAuth2ClientId
	} else {
		cDes.OAuth2ClientId = des.OAuth2ClientId
	}
	if dcl.StringCanonicalize(des.OAuth2ClientSecret, initial.OAuth2ClientSecret) || dcl.IsZeroValue(des.OAuth2ClientSecret) {
		cDes.OAuth2ClientSecret = initial.OAuth2ClientSecret
	} else {
		cDes.OAuth2ClientSecret = des.OAuth2ClientSecret
	}

	return cDes
}

func canonicalizeNewApplicationIap(c *Client, des, nw *ApplicationIap) *ApplicationIap {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.StringCanonicalize(des.OAuth2ClientId, nw.OAuth2ClientId) {
		nw.OAuth2ClientId = des.OAuth2ClientId
	}
	nw.OAuth2ClientSecret = des.OAuth2ClientSecret
	if dcl.StringCanonicalize(des.OAuth2ClientSecretSha256, nw.OAuth2ClientSecretSha256) {
		nw.OAuth2ClientSecretSha256 = des.OAuth2ClientSecretSha256
	}

	return nw
}

func canonicalizeNewApplicationIapSet(c *Client, des, nw []ApplicationIap) []ApplicationIap {
	if des == nil {
		return nw
	}
	var reorderedNew []ApplicationIap
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareApplicationIapNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewApplicationIapSlice(c *Client, des, nw []ApplicationIap) []ApplicationIap {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ApplicationIap
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewApplicationIap(c, &d, &n))
	}

	return items
}

func canonicalizeApplicationFeatureSettings(des, initial *ApplicationFeatureSettings, opts ...dcl.ApplyOption) *ApplicationFeatureSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ApplicationFeatureSettings{}

	if dcl.BoolCanonicalize(des.SplitHealthChecks, initial.SplitHealthChecks) || dcl.IsZeroValue(des.SplitHealthChecks) {
		cDes.SplitHealthChecks = initial.SplitHealthChecks
	} else {
		cDes.SplitHealthChecks = des.SplitHealthChecks
	}
	if dcl.BoolCanonicalize(des.UseContainerOptimizedOs, initial.UseContainerOptimizedOs) || dcl.IsZeroValue(des.UseContainerOptimizedOs) {
		cDes.UseContainerOptimizedOs = initial.UseContainerOptimizedOs
	} else {
		cDes.UseContainerOptimizedOs = des.UseContainerOptimizedOs
	}

	return cDes
}

func canonicalizeNewApplicationFeatureSettings(c *Client, des, nw *ApplicationFeatureSettings) *ApplicationFeatureSettings {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.SplitHealthChecks, nw.SplitHealthChecks) {
		nw.SplitHealthChecks = des.SplitHealthChecks
	}
	if dcl.BoolCanonicalize(des.UseContainerOptimizedOs, nw.UseContainerOptimizedOs) {
		nw.UseContainerOptimizedOs = des.UseContainerOptimizedOs
	}

	return nw
}

func canonicalizeNewApplicationFeatureSettingsSet(c *Client, des, nw []ApplicationFeatureSettings) []ApplicationFeatureSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []ApplicationFeatureSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareApplicationFeatureSettingsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewApplicationFeatureSettingsSlice(c *Client, des, nw []ApplicationFeatureSettings) []ApplicationFeatureSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ApplicationFeatureSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewApplicationFeatureSettings(c, &d, &n))
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
func diffApplication(c *Client, desired, actual *Application, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DispatchRules, actual.DispatchRules, dcl.Info{ObjectFunction: compareApplicationDispatchRulesNewStyle, EmptyObject: EmptyApplicationDispatchRules, OperationSelector: dcl.TriggersOperation("updateApplicationPatchApplicationOperation")}, fn.AddNest("DispatchRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthDomain, actual.AuthDomain, dcl.Info{OperationSelector: dcl.TriggersOperation("updateApplicationPatchApplicationOperation")}, fn.AddNest("AuthDomain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocationId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CodeBucket, actual.CodeBucket, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CodeBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultCookieExpiration, actual.DefaultCookieExpiration, dcl.Info{OperationSelector: dcl.TriggersOperation("updateApplicationPatchApplicationOperation")}, fn.AddNest("DefaultCookieExpiration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServingStatus, actual.ServingStatus, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateApplicationPatchApplicationOperation")}, fn.AddNest("ServingStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultHostname, actual.DefaultHostname, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultHostname")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultBucket, actual.DefaultBucket, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Iap, actual.Iap, dcl.Info{ObjectFunction: compareApplicationIapNewStyle, EmptyObject: EmptyApplicationIap, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Iap")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GcrDomain, actual.GcrDomain, dcl.Info{OperationSelector: dcl.TriggersOperation("updateApplicationPatchApplicationOperation")}, fn.AddNest("GcrDomain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatabaseType, actual.DatabaseType, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateApplicationPatchApplicationOperation")}, fn.AddNest("DatabaseType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FeatureSettings, actual.FeatureSettings, dcl.Info{ObjectFunction: compareApplicationFeatureSettingsNewStyle, EmptyObject: EmptyApplicationFeatureSettings, OperationSelector: dcl.TriggersOperation("updateApplicationPatchApplicationOperation")}, fn.AddNest("FeatureSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareApplicationDispatchRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ApplicationDispatchRules)
	if !ok {
		desiredNotPointer, ok := d.(ApplicationDispatchRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ApplicationDispatchRules or *ApplicationDispatchRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ApplicationDispatchRules)
	if !ok {
		actualNotPointer, ok := a.(ApplicationDispatchRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ApplicationDispatchRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Domain, actual.Domain, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Domain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareApplicationIapNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ApplicationIap)
	if !ok {
		desiredNotPointer, ok := d.(ApplicationIap)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ApplicationIap or *ApplicationIap", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ApplicationIap)
	if !ok {
		actualNotPointer, ok := a.(ApplicationIap)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ApplicationIap", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuth2ClientId, actual.OAuth2ClientId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Oauth2ClientId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuth2ClientSecret, actual.OAuth2ClientSecret, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Oauth2ClientSecret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuth2ClientSecretSha256, actual.OAuth2ClientSecretSha256, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Oauth2ClientSecretSha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareApplicationFeatureSettingsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ApplicationFeatureSettings)
	if !ok {
		desiredNotPointer, ok := d.(ApplicationFeatureSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ApplicationFeatureSettings or *ApplicationFeatureSettings", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ApplicationFeatureSettings)
	if !ok {
		actualNotPointer, ok := a.(ApplicationFeatureSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ApplicationFeatureSettings", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SplitHealthChecks, actual.SplitHealthChecks, dcl.Info{OperationSelector: dcl.TriggersOperation("updateApplicationPatchApplicationOperation")}, fn.AddNest("SplitHealthChecks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UseContainerOptimizedOs, actual.UseContainerOptimizedOs, dcl.Info{OperationSelector: dcl.TriggersOperation("updateApplicationPatchApplicationOperation")}, fn.AddNest("UseContainerOptimizedOs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Application) getFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

func (r *Application) createFields() string {
	return ""
}

func (r *Application) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "PatchApplication" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("apps/{{name}}", "https://appengine.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Application resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Application) marshal(c *Client) ([]byte, error) {
	m, err := expandApplication(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Application: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalApplication decodes JSON responses into the Application resource schema.
func unmarshalApplication(b []byte, c *Client) (*Application, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapApplication(m, c)
}

func unmarshalMapApplication(m map[string]interface{}, c *Client) (*Application, error) {

	return flattenApplication(c, m), nil
}

// expandApplication expands Application into a JSON request object.
func expandApplication(c *Client, f *Application) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v, err := expandApplicationDispatchRulesSlice(c, f.DispatchRules); err != nil {
		return nil, fmt.Errorf("error expanding DispatchRules into dispatchRules: %w", err)
	} else if v != nil {
		m["dispatchRules"] = v
	}
	if v := f.AuthDomain; !dcl.IsEmptyValueIndirect(v) {
		m["authDomain"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["locationId"] = v
	}
	if v := f.CodeBucket; !dcl.IsEmptyValueIndirect(v) {
		m["codeBucket"] = v
	}
	if v := f.DefaultCookieExpiration; !dcl.IsEmptyValueIndirect(v) {
		m["defaultCookieExpiration"] = v
	}
	if v := f.ServingStatus; !dcl.IsEmptyValueIndirect(v) {
		m["servingStatus"] = v
	}
	if v := f.DefaultHostname; !dcl.IsEmptyValueIndirect(v) {
		m["defaultHostname"] = v
	}
	if v := f.DefaultBucket; !dcl.IsEmptyValueIndirect(v) {
		m["defaultBucket"] = v
	}
	if v, err := expandApplicationIap(c, f.Iap); err != nil {
		return nil, fmt.Errorf("error expanding Iap into iap: %w", err)
	} else if v != nil {
		m["iap"] = v
	}
	if v := f.GcrDomain; !dcl.IsEmptyValueIndirect(v) {
		m["gcrDomain"] = v
	}
	if v := f.DatabaseType; !dcl.IsEmptyValueIndirect(v) {
		m["databaseType"] = v
	}
	if v, err := expandApplicationFeatureSettings(c, f.FeatureSettings); err != nil {
		return nil, fmt.Errorf("error expanding FeatureSettings into featureSettings: %w", err)
	} else if v != nil {
		m["featureSettings"] = v
	}

	return m, nil
}

// flattenApplication flattens Application from a JSON request object into the
// Application type.
func flattenApplication(c *Client, i interface{}) *Application {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Application{}
	res.Name = dcl.FlattenString(m["id"])
	res.DispatchRules = flattenApplicationDispatchRulesSlice(c, m["dispatchRules"])
	res.AuthDomain = dcl.FlattenString(m["authDomain"])
	res.Location = dcl.FlattenString(m["locationId"])
	res.CodeBucket = dcl.FlattenString(m["codeBucket"])
	res.DefaultCookieExpiration = dcl.FlattenString(m["defaultCookieExpiration"])
	res.ServingStatus = flattenApplicationServingStatusEnum(m["servingStatus"])
	res.DefaultHostname = dcl.FlattenString(m["defaultHostname"])
	res.DefaultBucket = dcl.FlattenString(m["defaultBucket"])
	res.Iap = flattenApplicationIap(c, m["iap"])
	res.GcrDomain = dcl.FlattenString(m["gcrDomain"])
	res.DatabaseType = flattenApplicationDatabaseTypeEnum(m["databaseType"])
	res.FeatureSettings = flattenApplicationFeatureSettings(c, m["featureSettings"])

	return res
}

// expandApplicationDispatchRulesMap expands the contents of ApplicationDispatchRules into a JSON
// request object.
func expandApplicationDispatchRulesMap(c *Client, f map[string]ApplicationDispatchRules) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandApplicationDispatchRules(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandApplicationDispatchRulesSlice expands the contents of ApplicationDispatchRules into a JSON
// request object.
func expandApplicationDispatchRulesSlice(c *Client, f []ApplicationDispatchRules) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandApplicationDispatchRules(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenApplicationDispatchRulesMap flattens the contents of ApplicationDispatchRules from a JSON
// response object.
func flattenApplicationDispatchRulesMap(c *Client, i interface{}) map[string]ApplicationDispatchRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ApplicationDispatchRules{}
	}

	if len(a) == 0 {
		return map[string]ApplicationDispatchRules{}
	}

	items := make(map[string]ApplicationDispatchRules)
	for k, item := range a {
		items[k] = *flattenApplicationDispatchRules(c, item.(map[string]interface{}))
	}

	return items
}

// flattenApplicationDispatchRulesSlice flattens the contents of ApplicationDispatchRules from a JSON
// response object.
func flattenApplicationDispatchRulesSlice(c *Client, i interface{}) []ApplicationDispatchRules {
	a, ok := i.([]interface{})
	if !ok {
		return []ApplicationDispatchRules{}
	}

	if len(a) == 0 {
		return []ApplicationDispatchRules{}
	}

	items := make([]ApplicationDispatchRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenApplicationDispatchRules(c, item.(map[string]interface{})))
	}

	return items
}

// expandApplicationDispatchRules expands an instance of ApplicationDispatchRules into a JSON
// request object.
func expandApplicationDispatchRules(c *Client, f *ApplicationDispatchRules) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Domain; !dcl.IsEmptyValueIndirect(v) {
		m["domain"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}

	return m, nil
}

// flattenApplicationDispatchRules flattens an instance of ApplicationDispatchRules from a JSON
// response object.
func flattenApplicationDispatchRules(c *Client, i interface{}) *ApplicationDispatchRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ApplicationDispatchRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyApplicationDispatchRules
	}
	r.Domain = dcl.FlattenString(m["domain"])
	r.Path = dcl.FlattenString(m["path"])
	r.Service = dcl.FlattenString(m["service"])

	return r
}

// expandApplicationIapMap expands the contents of ApplicationIap into a JSON
// request object.
func expandApplicationIapMap(c *Client, f map[string]ApplicationIap) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandApplicationIap(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandApplicationIapSlice expands the contents of ApplicationIap into a JSON
// request object.
func expandApplicationIapSlice(c *Client, f []ApplicationIap) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandApplicationIap(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenApplicationIapMap flattens the contents of ApplicationIap from a JSON
// response object.
func flattenApplicationIapMap(c *Client, i interface{}) map[string]ApplicationIap {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ApplicationIap{}
	}

	if len(a) == 0 {
		return map[string]ApplicationIap{}
	}

	items := make(map[string]ApplicationIap)
	for k, item := range a {
		items[k] = *flattenApplicationIap(c, item.(map[string]interface{}))
	}

	return items
}

// flattenApplicationIapSlice flattens the contents of ApplicationIap from a JSON
// response object.
func flattenApplicationIapSlice(c *Client, i interface{}) []ApplicationIap {
	a, ok := i.([]interface{})
	if !ok {
		return []ApplicationIap{}
	}

	if len(a) == 0 {
		return []ApplicationIap{}
	}

	items := make([]ApplicationIap, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenApplicationIap(c, item.(map[string]interface{})))
	}

	return items
}

// expandApplicationIap expands an instance of ApplicationIap into a JSON
// request object.
func expandApplicationIap(c *Client, f *ApplicationIap) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.OAuth2ClientId; !dcl.IsEmptyValueIndirect(v) {
		m["oauth2ClientId"] = v
	}
	if v := f.OAuth2ClientSecret; !dcl.IsEmptyValueIndirect(v) {
		m["oauth2ClientSecret"] = v
	}
	if v := f.OAuth2ClientSecretSha256; !dcl.IsEmptyValueIndirect(v) {
		m["oauth2ClientSecretSha256"] = v
	}

	return m, nil
}

// flattenApplicationIap flattens an instance of ApplicationIap from a JSON
// response object.
func flattenApplicationIap(c *Client, i interface{}) *ApplicationIap {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ApplicationIap{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyApplicationIap
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.OAuth2ClientId = dcl.FlattenString(m["oauth2ClientId"])
	r.OAuth2ClientSecret = dcl.FlattenString(m["oauth2ClientSecret"])
	r.OAuth2ClientSecretSha256 = dcl.FlattenString(m["oauth2ClientSecretSha256"])

	return r
}

// expandApplicationFeatureSettingsMap expands the contents of ApplicationFeatureSettings into a JSON
// request object.
func expandApplicationFeatureSettingsMap(c *Client, f map[string]ApplicationFeatureSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandApplicationFeatureSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandApplicationFeatureSettingsSlice expands the contents of ApplicationFeatureSettings into a JSON
// request object.
func expandApplicationFeatureSettingsSlice(c *Client, f []ApplicationFeatureSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandApplicationFeatureSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenApplicationFeatureSettingsMap flattens the contents of ApplicationFeatureSettings from a JSON
// response object.
func flattenApplicationFeatureSettingsMap(c *Client, i interface{}) map[string]ApplicationFeatureSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ApplicationFeatureSettings{}
	}

	if len(a) == 0 {
		return map[string]ApplicationFeatureSettings{}
	}

	items := make(map[string]ApplicationFeatureSettings)
	for k, item := range a {
		items[k] = *flattenApplicationFeatureSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenApplicationFeatureSettingsSlice flattens the contents of ApplicationFeatureSettings from a JSON
// response object.
func flattenApplicationFeatureSettingsSlice(c *Client, i interface{}) []ApplicationFeatureSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []ApplicationFeatureSettings{}
	}

	if len(a) == 0 {
		return []ApplicationFeatureSettings{}
	}

	items := make([]ApplicationFeatureSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenApplicationFeatureSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandApplicationFeatureSettings expands an instance of ApplicationFeatureSettings into a JSON
// request object.
func expandApplicationFeatureSettings(c *Client, f *ApplicationFeatureSettings) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SplitHealthChecks; !dcl.IsEmptyValueIndirect(v) {
		m["splitHealthChecks"] = v
	}
	if v := f.UseContainerOptimizedOs; !dcl.IsEmptyValueIndirect(v) {
		m["useContainerOptimizedOs"] = v
	}

	return m, nil
}

// flattenApplicationFeatureSettings flattens an instance of ApplicationFeatureSettings from a JSON
// response object.
func flattenApplicationFeatureSettings(c *Client, i interface{}) *ApplicationFeatureSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ApplicationFeatureSettings{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyApplicationFeatureSettings
	}
	r.SplitHealthChecks = dcl.FlattenBool(m["splitHealthChecks"])
	r.UseContainerOptimizedOs = dcl.FlattenBool(m["useContainerOptimizedOs"])

	return r
}

// flattenApplicationServingStatusEnumMap flattens the contents of ApplicationServingStatusEnum from a JSON
// response object.
func flattenApplicationServingStatusEnumMap(c *Client, i interface{}) map[string]ApplicationServingStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ApplicationServingStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]ApplicationServingStatusEnum{}
	}

	items := make(map[string]ApplicationServingStatusEnum)
	for k, item := range a {
		items[k] = *flattenApplicationServingStatusEnum(item.(interface{}))
	}

	return items
}

// flattenApplicationServingStatusEnumSlice flattens the contents of ApplicationServingStatusEnum from a JSON
// response object.
func flattenApplicationServingStatusEnumSlice(c *Client, i interface{}) []ApplicationServingStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ApplicationServingStatusEnum{}
	}

	if len(a) == 0 {
		return []ApplicationServingStatusEnum{}
	}

	items := make([]ApplicationServingStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenApplicationServingStatusEnum(item.(interface{})))
	}

	return items
}

// flattenApplicationServingStatusEnum asserts that an interface is a string, and returns a
// pointer to a *ApplicationServingStatusEnum with the same value as that string.
func flattenApplicationServingStatusEnum(i interface{}) *ApplicationServingStatusEnum {
	s, ok := i.(string)
	if !ok {
		return ApplicationServingStatusEnumRef("")
	}

	return ApplicationServingStatusEnumRef(s)
}

// flattenApplicationDatabaseTypeEnumMap flattens the contents of ApplicationDatabaseTypeEnum from a JSON
// response object.
func flattenApplicationDatabaseTypeEnumMap(c *Client, i interface{}) map[string]ApplicationDatabaseTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ApplicationDatabaseTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ApplicationDatabaseTypeEnum{}
	}

	items := make(map[string]ApplicationDatabaseTypeEnum)
	for k, item := range a {
		items[k] = *flattenApplicationDatabaseTypeEnum(item.(interface{}))
	}

	return items
}

// flattenApplicationDatabaseTypeEnumSlice flattens the contents of ApplicationDatabaseTypeEnum from a JSON
// response object.
func flattenApplicationDatabaseTypeEnumSlice(c *Client, i interface{}) []ApplicationDatabaseTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ApplicationDatabaseTypeEnum{}
	}

	if len(a) == 0 {
		return []ApplicationDatabaseTypeEnum{}
	}

	items := make([]ApplicationDatabaseTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenApplicationDatabaseTypeEnum(item.(interface{})))
	}

	return items
}

// flattenApplicationDatabaseTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ApplicationDatabaseTypeEnum with the same value as that string.
func flattenApplicationDatabaseTypeEnum(i interface{}) *ApplicationDatabaseTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ApplicationDatabaseTypeEnumRef("")
	}

	return ApplicationDatabaseTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Application) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalApplication(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

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

type applicationDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         applicationApiOperation
}

func convertFieldDiffsToApplicationDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]applicationDiff, error) {
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
	var diffs []applicationDiff
	// For each operation name, create a applicationDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := applicationDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToApplicationApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToApplicationApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (applicationApiOperation, error) {
	switch opName {

	case "updateApplicationPatchApplicationOperation":
		return &updateApplicationPatchApplicationOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
