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
package alpha

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

func (r *AuthorizationPolicy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "action"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}
func (r *AuthorizationPolicyRules) validate() error {
	return nil
}
func (r *AuthorizationPolicyRulesSources) validate() error {
	return nil
}
func (r *AuthorizationPolicyRulesDestinations) validate() error {
	if err := dcl.Required(r, "hosts"); err != nil {
		return err
	}
	if err := dcl.Required(r, "ports"); err != nil {
		return err
	}
	if err := dcl.Required(r, "paths"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.HttpHeaderMatch) {
		if err := r.HttpHeaderMatch.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) validate() error {
	if err := dcl.Required(r, "headerName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "regexMatch"); err != nil {
		return err
	}
	return nil
}

func authorizationPolicyGetURL(userBasePath string, r *AuthorizationPolicy) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, params), nil
}

func authorizationPolicyListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, params), nil

}

func authorizationPolicyCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies?authorizationPolicyId={{name}}", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, params), nil

}

func authorizationPolicyDeleteURL(userBasePath string, r *AuthorizationPolicy) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, params), nil
}

func (r *AuthorizationPolicy) SetPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *n.Project,
		"location": *n.Location,
		"name":     *n.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}:setIamPolicy", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, fields)
}

func (r *AuthorizationPolicy) getPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *n.Project,
		"location": *n.Location,
		"name":     *n.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}:getIamPolicy", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, fields)
}

func (r *AuthorizationPolicy) IAMPolicyVersion() int {
	return 3
}

// authorizationPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type authorizationPolicyApiOperation interface {
	do(context.Context, *AuthorizationPolicy, *Client) error
}

// newUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest creates a request for an
// AuthorizationPolicy resource's UpdateAuthorizationPolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest(ctx context.Context, f *AuthorizationPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Action; !dcl.IsEmptyValueIndirect(v) {
		req["action"] = v
	}
	if v, err := expandAuthorizationPolicyRulesSlice(c, f.Rules); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["rules"] = v
	}
	return req, nil
}

// marshalUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAuthorizationPolicyUpdateAuthorizationPolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAuthorizationPolicyUpdateAuthorizationPolicyOperation) do(ctx context.Context, r *AuthorizationPolicy, c *Client) error {
	_, err := c.GetAuthorizationPolicy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateAuthorizationPolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateAuthorizationPolicyUpdateAuthorizationPolicyRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://networksecurity.googleapis.com/v1alpha1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listAuthorizationPolicyRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := authorizationPolicyListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AuthorizationPolicyMaxPage {
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

type listAuthorizationPolicyOperation struct {
	AuthorizationPolicies []map[string]interface{} `json:"authorizationPolicies"`
	Token                 string                   `json:"nextPageToken"`
}

func (c *Client) listAuthorizationPolicy(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*AuthorizationPolicy, string, error) {
	b, err := c.listAuthorizationPolicyRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAuthorizationPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AuthorizationPolicy
	for _, v := range m.AuthorizationPolicies {
		res := flattenAuthorizationPolicy(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAuthorizationPolicy(ctx context.Context, f func(*AuthorizationPolicy) bool, resources []*AuthorizationPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAuthorizationPolicy(ctx, res)
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

type deleteAuthorizationPolicyOperation struct{}

func (op *deleteAuthorizationPolicyOperation) do(ctx context.Context, r *AuthorizationPolicy, c *Client) error {

	_, err := c.GetAuthorizationPolicy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("AuthorizationPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAuthorizationPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := authorizationPolicyDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://networksecurity.googleapis.com/v1alpha1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAuthorizationPolicy(ctx, r.urlNormalized())
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
type createAuthorizationPolicyOperation struct {
	response map[string]interface{}
}

func (op *createAuthorizationPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAuthorizationPolicyOperation) do(ctx context.Context, r *AuthorizationPolicy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := authorizationPolicyCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://networksecurity.googleapis.com/v1alpha1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetAuthorizationPolicy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAuthorizationPolicyRaw(ctx context.Context, r *AuthorizationPolicy) ([]byte, error) {

	u, err := authorizationPolicyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) authorizationPolicyDiffsForRawDesired(ctx context.Context, rawDesired *AuthorizationPolicy, opts ...dcl.ApplyOption) (initial, desired *AuthorizationPolicy, diffs []authorizationPolicyDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AuthorizationPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AuthorizationPolicy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected AuthorizationPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAuthorizationPolicy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a AuthorizationPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AuthorizationPolicy resource: %v", err)
		}
		c.Config.Logger.Info("Found that AuthorizationPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAuthorizationPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for AuthorizationPolicy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for AuthorizationPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAuthorizationPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for AuthorizationPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAuthorizationPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for AuthorizationPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAuthorizationPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAuthorizationPolicyInitialState(rawInitial, rawDesired *AuthorizationPolicy) (*AuthorizationPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAuthorizationPolicyDesiredState(rawDesired, rawInitial *AuthorizationPolicy, opts ...dcl.ApplyOption) (*AuthorizationPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.Action) {
		rawDesired.Action = rawInitial.Action
	}
	if dcl.IsZeroValue(rawDesired.Rules) {
		rawDesired.Rules = rawInitial.Rules
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeAuthorizationPolicyNewState(c *Client, rawNew, rawDesired *AuthorizationPolicy) (*AuthorizationPolicy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
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

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Action) && dcl.IsEmptyValueIndirect(rawDesired.Action) {
		rawNew.Action = rawDesired.Action
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Rules) && dcl.IsEmptyValueIndirect(rawDesired.Rules) {
		rawNew.Rules = rawDesired.Rules
	} else {
		rawNew.Rules = canonicalizeNewAuthorizationPolicyRulesSlice(c, rawDesired.Rules, rawNew.Rules)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeAuthorizationPolicyRules(des, initial *AuthorizationPolicyRules, opts ...dcl.ApplyOption) *AuthorizationPolicyRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Sources) {
		des.Sources = initial.Sources
	}
	if dcl.IsZeroValue(des.Destinations) {
		des.Destinations = initial.Destinations
	}

	return des
}

func canonicalizeNewAuthorizationPolicyRules(c *Client, des, nw *AuthorizationPolicyRules) *AuthorizationPolicyRules {
	if des == nil || nw == nil {
		return nw
	}

	nw.Sources = canonicalizeNewAuthorizationPolicyRulesSourcesSlice(c, des.Sources, nw.Sources)
	nw.Destinations = canonicalizeNewAuthorizationPolicyRulesDestinationsSlice(c, des.Destinations, nw.Destinations)

	return nw
}

func canonicalizeNewAuthorizationPolicyRulesSet(c *Client, des, nw []AuthorizationPolicyRules) []AuthorizationPolicyRules {
	if des == nil {
		return nw
	}
	var reorderedNew []AuthorizationPolicyRules
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAuthorizationPolicyRules(c, &d, &n) {
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

func canonicalizeNewAuthorizationPolicyRulesSlice(c *Client, des, nw []AuthorizationPolicyRules) []AuthorizationPolicyRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AuthorizationPolicyRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAuthorizationPolicyRules(c, &d, &n))
	}

	return items
}

func canonicalizeAuthorizationPolicyRulesSources(des, initial *AuthorizationPolicyRulesSources, opts ...dcl.ApplyOption) *AuthorizationPolicyRulesSources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Principals) {
		des.Principals = initial.Principals
	}
	if dcl.IsZeroValue(des.IPBlocks) {
		des.IPBlocks = initial.IPBlocks
	}

	return des
}

func canonicalizeNewAuthorizationPolicyRulesSources(c *Client, des, nw *AuthorizationPolicyRulesSources) *AuthorizationPolicyRulesSources {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewAuthorizationPolicyRulesSourcesSet(c *Client, des, nw []AuthorizationPolicyRulesSources) []AuthorizationPolicyRulesSources {
	if des == nil {
		return nw
	}
	var reorderedNew []AuthorizationPolicyRulesSources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAuthorizationPolicyRulesSources(c, &d, &n) {
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

func canonicalizeNewAuthorizationPolicyRulesSourcesSlice(c *Client, des, nw []AuthorizationPolicyRulesSources) []AuthorizationPolicyRulesSources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AuthorizationPolicyRulesSources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAuthorizationPolicyRulesSources(c, &d, &n))
	}

	return items
}

func canonicalizeAuthorizationPolicyRulesDestinations(des, initial *AuthorizationPolicyRulesDestinations, opts ...dcl.ApplyOption) *AuthorizationPolicyRulesDestinations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Hosts) {
		des.Hosts = initial.Hosts
	}
	if dcl.IsZeroValue(des.Ports) {
		des.Ports = initial.Ports
	}
	if dcl.IsZeroValue(des.Paths) {
		des.Paths = initial.Paths
	}
	if dcl.IsZeroValue(des.Methods) {
		des.Methods = initial.Methods
	}
	des.HttpHeaderMatch = canonicalizeAuthorizationPolicyRulesDestinationsHttpHeaderMatch(des.HttpHeaderMatch, initial.HttpHeaderMatch, opts...)

	return des
}

func canonicalizeNewAuthorizationPolicyRulesDestinations(c *Client, des, nw *AuthorizationPolicyRulesDestinations) *AuthorizationPolicyRulesDestinations {
	if des == nil || nw == nil {
		return nw
	}

	nw.HttpHeaderMatch = canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, des.HttpHeaderMatch, nw.HttpHeaderMatch)

	return nw
}

func canonicalizeNewAuthorizationPolicyRulesDestinationsSet(c *Client, des, nw []AuthorizationPolicyRulesDestinations) []AuthorizationPolicyRulesDestinations {
	if des == nil {
		return nw
	}
	var reorderedNew []AuthorizationPolicyRulesDestinations
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAuthorizationPolicyRulesDestinations(c, &d, &n) {
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

func canonicalizeNewAuthorizationPolicyRulesDestinationsSlice(c *Client, des, nw []AuthorizationPolicyRulesDestinations) []AuthorizationPolicyRulesDestinations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AuthorizationPolicyRulesDestinations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAuthorizationPolicyRulesDestinations(c, &d, &n))
	}

	return items
}

func canonicalizeAuthorizationPolicyRulesDestinationsHttpHeaderMatch(des, initial *AuthorizationPolicyRulesDestinationsHttpHeaderMatch, opts ...dcl.ApplyOption) *AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.HeaderName, initial.HeaderName) || dcl.IsZeroValue(des.HeaderName) {
		des.HeaderName = initial.HeaderName
	}
	if dcl.StringCanonicalize(des.RegexMatch, initial.RegexMatch) || dcl.IsZeroValue(des.RegexMatch) {
		des.RegexMatch = initial.RegexMatch
	}

	return des
}

func canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c *Client, des, nw *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) *AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HeaderName, nw.HeaderName) {
		nw.HeaderName = des.HeaderName
	}
	if dcl.StringCanonicalize(des.RegexMatch, nw.RegexMatch) {
		nw.RegexMatch = des.RegexMatch
	}

	return nw
}

func canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatchSet(c *Client, des, nw []AuthorizationPolicyRulesDestinationsHttpHeaderMatch) []AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if des == nil {
		return nw
	}
	var reorderedNew []AuthorizationPolicyRulesDestinationsHttpHeaderMatch
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, &d, &n) {
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

func canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice(c *Client, des, nw []AuthorizationPolicyRulesDestinationsHttpHeaderMatch) []AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AuthorizationPolicyRulesDestinationsHttpHeaderMatch
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, &d, &n))
	}

	return items
}

type authorizationPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         authorizationPolicyApiOperation
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
func diffAuthorizationPolicy(c *Client, desired, actual *AuthorizationPolicy, opts ...dcl.ApplyOption) ([]authorizationPolicyDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []authorizationPolicyDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, authorizationPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, authorizationPolicyDiff{
			UpdateOp:  &updateAuthorizationPolicyUpdateAuthorizationPolicyOperation{},
			FieldName: "Description",
		})

	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, authorizationPolicyDiff{
			UpdateOp:  &updateAuthorizationPolicyUpdateAuthorizationPolicyOperation{},
			FieldName: "Labels",
		})

	}
	if !reflect.DeepEqual(desired.Action, actual.Action) {
		c.Config.Logger.Infof("Detected diff in Action.\nDESIRED: %v\nACTUAL: %v", desired.Action, actual.Action)

		diffs = append(diffs, authorizationPolicyDiff{
			UpdateOp:  &updateAuthorizationPolicyUpdateAuthorizationPolicyOperation{},
			FieldName: "Action",
		})

	}
	if compareAuthorizationPolicyRulesSlice(c, desired.Rules, actual.Rules) {
		c.Config.Logger.Infof("Detected diff in Rules.\nDESIRED: %v\nACTUAL: %v", desired.Rules, actual.Rules)

		diffs = append(diffs, authorizationPolicyDiff{
			UpdateOp:  &updateAuthorizationPolicyUpdateAuthorizationPolicyOperation{},
			FieldName: "Rules",
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
	var deduped []authorizationPolicyDiff
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
func compareAuthorizationPolicyRules(c *Client, desired, actual *AuthorizationPolicyRules) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Sources == nil && desired.Sources != nil && !dcl.IsEmptyValueIndirect(desired.Sources) {
		c.Config.Logger.Infof("desired Sources %s - but actually nil", dcl.SprintResource(desired.Sources))
		return true
	}
	if compareAuthorizationPolicyRulesSourcesSlice(c, desired.Sources, actual.Sources) && !dcl.IsZeroValue(desired.Sources) {
		c.Config.Logger.Infof("Diff in Sources. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sources), dcl.SprintResource(actual.Sources))
		return true
	}
	if actual.Destinations == nil && desired.Destinations != nil && !dcl.IsEmptyValueIndirect(desired.Destinations) {
		c.Config.Logger.Infof("desired Destinations %s - but actually nil", dcl.SprintResource(desired.Destinations))
		return true
	}
	if compareAuthorizationPolicyRulesDestinationsSlice(c, desired.Destinations, actual.Destinations) && !dcl.IsZeroValue(desired.Destinations) {
		c.Config.Logger.Infof("Diff in Destinations. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Destinations), dcl.SprintResource(actual.Destinations))
		return true
	}
	return false
}

func compareAuthorizationPolicyRulesSlice(c *Client, desired, actual []AuthorizationPolicyRules) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AuthorizationPolicyRules, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAuthorizationPolicyRules(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRules, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAuthorizationPolicyRulesMap(c *Client, desired, actual map[string]AuthorizationPolicyRules) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AuthorizationPolicyRules, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRules, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAuthorizationPolicyRules(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRules, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAuthorizationPolicyRulesSources(c *Client, desired, actual *AuthorizationPolicyRulesSources) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Principals == nil && desired.Principals != nil && !dcl.IsEmptyValueIndirect(desired.Principals) {
		c.Config.Logger.Infof("desired Principals %s - but actually nil", dcl.SprintResource(desired.Principals))
		return true
	}
	if !dcl.StringSliceEquals(desired.Principals, actual.Principals) && !dcl.IsZeroValue(desired.Principals) {
		c.Config.Logger.Infof("Diff in Principals. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Principals), dcl.SprintResource(actual.Principals))
		return true
	}
	if actual.IPBlocks == nil && desired.IPBlocks != nil && !dcl.IsEmptyValueIndirect(desired.IPBlocks) {
		c.Config.Logger.Infof("desired IPBlocks %s - but actually nil", dcl.SprintResource(desired.IPBlocks))
		return true
	}
	if !dcl.StringSliceEquals(desired.IPBlocks, actual.IPBlocks) && !dcl.IsZeroValue(desired.IPBlocks) {
		c.Config.Logger.Infof("Diff in IPBlocks. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPBlocks), dcl.SprintResource(actual.IPBlocks))
		return true
	}
	return false
}

func compareAuthorizationPolicyRulesSourcesSlice(c *Client, desired, actual []AuthorizationPolicyRulesSources) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AuthorizationPolicyRulesSources, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAuthorizationPolicyRulesSources(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRulesSources, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAuthorizationPolicyRulesSourcesMap(c *Client, desired, actual map[string]AuthorizationPolicyRulesSources) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AuthorizationPolicyRulesSources, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRulesSources, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAuthorizationPolicyRulesSources(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRulesSources, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAuthorizationPolicyRulesDestinations(c *Client, desired, actual *AuthorizationPolicyRulesDestinations) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Hosts == nil && desired.Hosts != nil && !dcl.IsEmptyValueIndirect(desired.Hosts) {
		c.Config.Logger.Infof("desired Hosts %s - but actually nil", dcl.SprintResource(desired.Hosts))
		return true
	}
	if !dcl.StringSliceEquals(desired.Hosts, actual.Hosts) && !dcl.IsZeroValue(desired.Hosts) {
		c.Config.Logger.Infof("Diff in Hosts. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Hosts), dcl.SprintResource(actual.Hosts))
		return true
	}
	if actual.Ports == nil && desired.Ports != nil && !dcl.IsEmptyValueIndirect(desired.Ports) {
		c.Config.Logger.Infof("desired Ports %s - but actually nil", dcl.SprintResource(desired.Ports))
		return true
	}
	if !dcl.IntSliceEquals(desired.Ports, actual.Ports) && !dcl.IsZeroValue(desired.Ports) {
		c.Config.Logger.Infof("Diff in Ports. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Ports), dcl.SprintResource(actual.Ports))
		return true
	}
	if actual.Paths == nil && desired.Paths != nil && !dcl.IsEmptyValueIndirect(desired.Paths) {
		c.Config.Logger.Infof("desired Paths %s - but actually nil", dcl.SprintResource(desired.Paths))
		return true
	}
	if !dcl.StringSliceEquals(desired.Paths, actual.Paths) && !dcl.IsZeroValue(desired.Paths) {
		c.Config.Logger.Infof("Diff in Paths. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Paths), dcl.SprintResource(actual.Paths))
		return true
	}
	if actual.Methods == nil && desired.Methods != nil && !dcl.IsEmptyValueIndirect(desired.Methods) {
		c.Config.Logger.Infof("desired Methods %s - but actually nil", dcl.SprintResource(desired.Methods))
		return true
	}
	if !dcl.StringSliceEquals(desired.Methods, actual.Methods) && !dcl.IsZeroValue(desired.Methods) {
		c.Config.Logger.Infof("Diff in Methods. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Methods), dcl.SprintResource(actual.Methods))
		return true
	}
	if actual.HttpHeaderMatch == nil && desired.HttpHeaderMatch != nil && !dcl.IsEmptyValueIndirect(desired.HttpHeaderMatch) {
		c.Config.Logger.Infof("desired HttpHeaderMatch %s - but actually nil", dcl.SprintResource(desired.HttpHeaderMatch))
		return true
	}
	if compareAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, desired.HttpHeaderMatch, actual.HttpHeaderMatch) && !dcl.IsZeroValue(desired.HttpHeaderMatch) {
		c.Config.Logger.Infof("Diff in HttpHeaderMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpHeaderMatch), dcl.SprintResource(actual.HttpHeaderMatch))
		return true
	}
	return false
}

func compareAuthorizationPolicyRulesDestinationsSlice(c *Client, desired, actual []AuthorizationPolicyRulesDestinations) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AuthorizationPolicyRulesDestinations, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAuthorizationPolicyRulesDestinations(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRulesDestinations, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAuthorizationPolicyRulesDestinationsMap(c *Client, desired, actual map[string]AuthorizationPolicyRulesDestinations) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AuthorizationPolicyRulesDestinations, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRulesDestinations, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAuthorizationPolicyRulesDestinations(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRulesDestinations, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c *Client, desired, actual *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HeaderName == nil && desired.HeaderName != nil && !dcl.IsEmptyValueIndirect(desired.HeaderName) {
		c.Config.Logger.Infof("desired HeaderName %s - but actually nil", dcl.SprintResource(desired.HeaderName))
		return true
	}
	if !dcl.StringCanonicalize(desired.HeaderName, actual.HeaderName) && !dcl.IsZeroValue(desired.HeaderName) {
		c.Config.Logger.Infof("Diff in HeaderName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HeaderName), dcl.SprintResource(actual.HeaderName))
		return true
	}
	if actual.RegexMatch == nil && desired.RegexMatch != nil && !dcl.IsEmptyValueIndirect(desired.RegexMatch) {
		c.Config.Logger.Infof("desired RegexMatch %s - but actually nil", dcl.SprintResource(desired.RegexMatch))
		return true
	}
	if !dcl.StringCanonicalize(desired.RegexMatch, actual.RegexMatch) && !dcl.IsZeroValue(desired.RegexMatch) {
		c.Config.Logger.Infof("Diff in RegexMatch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RegexMatch), dcl.SprintResource(actual.RegexMatch))
		return true
	}
	return false
}

func compareAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice(c *Client, desired, actual []AuthorizationPolicyRulesDestinationsHttpHeaderMatch) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AuthorizationPolicyRulesDestinationsHttpHeaderMatch, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRulesDestinationsHttpHeaderMatch, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAuthorizationPolicyRulesDestinationsHttpHeaderMatchMap(c *Client, desired, actual map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AuthorizationPolicyRulesDestinationsHttpHeaderMatch, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRulesDestinationsHttpHeaderMatch, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyRulesDestinationsHttpHeaderMatch, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAuthorizationPolicyActionEnumSlice(c *Client, desired, actual []AuthorizationPolicyActionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AuthorizationPolicyActionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAuthorizationPolicyActionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AuthorizationPolicyActionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAuthorizationPolicyActionEnum(c *Client, desired, actual *AuthorizationPolicyActionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *AuthorizationPolicy) urlNormalized() *AuthorizationPolicy {
	normalized := deepcopy.Copy(*r).(AuthorizationPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *AuthorizationPolicy) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *AuthorizationPolicy) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *AuthorizationPolicy) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *AuthorizationPolicy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateAuthorizationPolicy" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}", "https://networksecurity.googleapis.com/v1alpha1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AuthorizationPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AuthorizationPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandAuthorizationPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AuthorizationPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAuthorizationPolicy decodes JSON responses into the AuthorizationPolicy resource schema.
func unmarshalAuthorizationPolicy(b []byte, c *Client) (*AuthorizationPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAuthorizationPolicy(m, c)
}

func unmarshalMapAuthorizationPolicy(m map[string]interface{}, c *Client) (*AuthorizationPolicy, error) {

	return flattenAuthorizationPolicy(c, m), nil
}

// expandAuthorizationPolicy expands AuthorizationPolicy into a JSON request object.
func expandAuthorizationPolicy(c *Client, f *AuthorizationPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/*/locations/%s/authorizationPolicies/%s", f.Name, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Action; !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}
	if v, err := expandAuthorizationPolicyRulesSlice(c, f.Rules); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rules"] = v
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

// flattenAuthorizationPolicy flattens AuthorizationPolicy from a JSON request object into the
// AuthorizationPolicy type.
func flattenAuthorizationPolicy(c *Client, i interface{}) *AuthorizationPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &AuthorizationPolicy{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Action = flattenAuthorizationPolicyActionEnum(m["action"])
	r.Rules = flattenAuthorizationPolicyRulesSlice(c, m["rules"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandAuthorizationPolicyRulesMap expands the contents of AuthorizationPolicyRules into a JSON
// request object.
func expandAuthorizationPolicyRulesMap(c *Client, f map[string]AuthorizationPolicyRules) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAuthorizationPolicyRules(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAuthorizationPolicyRulesSlice expands the contents of AuthorizationPolicyRules into a JSON
// request object.
func expandAuthorizationPolicyRulesSlice(c *Client, f []AuthorizationPolicyRules) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAuthorizationPolicyRules(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAuthorizationPolicyRulesMap flattens the contents of AuthorizationPolicyRules from a JSON
// response object.
func flattenAuthorizationPolicyRulesMap(c *Client, i interface{}) map[string]AuthorizationPolicyRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AuthorizationPolicyRules{}
	}

	if len(a) == 0 {
		return map[string]AuthorizationPolicyRules{}
	}

	items := make(map[string]AuthorizationPolicyRules)
	for k, item := range a {
		items[k] = *flattenAuthorizationPolicyRules(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAuthorizationPolicyRulesSlice flattens the contents of AuthorizationPolicyRules from a JSON
// response object.
func flattenAuthorizationPolicyRulesSlice(c *Client, i interface{}) []AuthorizationPolicyRules {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyRules{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyRules{}
	}

	items := make([]AuthorizationPolicyRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyRules(c, item.(map[string]interface{})))
	}

	return items
}

// expandAuthorizationPolicyRules expands an instance of AuthorizationPolicyRules into a JSON
// request object.
func expandAuthorizationPolicyRules(c *Client, f *AuthorizationPolicyRules) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAuthorizationPolicyRulesSourcesSlice(c, f.Sources); err != nil {
		return nil, fmt.Errorf("error expanding Sources into sources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sources"] = v
	}
	if v, err := expandAuthorizationPolicyRulesDestinationsSlice(c, f.Destinations); err != nil {
		return nil, fmt.Errorf("error expanding Destinations into destinations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["destinations"] = v
	}

	return m, nil
}

// flattenAuthorizationPolicyRules flattens an instance of AuthorizationPolicyRules from a JSON
// response object.
func flattenAuthorizationPolicyRules(c *Client, i interface{}) *AuthorizationPolicyRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AuthorizationPolicyRules{}
	r.Sources = flattenAuthorizationPolicyRulesSourcesSlice(c, m["sources"])
	r.Destinations = flattenAuthorizationPolicyRulesDestinationsSlice(c, m["destinations"])

	return r
}

// expandAuthorizationPolicyRulesSourcesMap expands the contents of AuthorizationPolicyRulesSources into a JSON
// request object.
func expandAuthorizationPolicyRulesSourcesMap(c *Client, f map[string]AuthorizationPolicyRulesSources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAuthorizationPolicyRulesSources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAuthorizationPolicyRulesSourcesSlice expands the contents of AuthorizationPolicyRulesSources into a JSON
// request object.
func expandAuthorizationPolicyRulesSourcesSlice(c *Client, f []AuthorizationPolicyRulesSources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAuthorizationPolicyRulesSources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAuthorizationPolicyRulesSourcesMap flattens the contents of AuthorizationPolicyRulesSources from a JSON
// response object.
func flattenAuthorizationPolicyRulesSourcesMap(c *Client, i interface{}) map[string]AuthorizationPolicyRulesSources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AuthorizationPolicyRulesSources{}
	}

	if len(a) == 0 {
		return map[string]AuthorizationPolicyRulesSources{}
	}

	items := make(map[string]AuthorizationPolicyRulesSources)
	for k, item := range a {
		items[k] = *flattenAuthorizationPolicyRulesSources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAuthorizationPolicyRulesSourcesSlice flattens the contents of AuthorizationPolicyRulesSources from a JSON
// response object.
func flattenAuthorizationPolicyRulesSourcesSlice(c *Client, i interface{}) []AuthorizationPolicyRulesSources {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyRulesSources{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyRulesSources{}
	}

	items := make([]AuthorizationPolicyRulesSources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyRulesSources(c, item.(map[string]interface{})))
	}

	return items
}

// expandAuthorizationPolicyRulesSources expands an instance of AuthorizationPolicyRulesSources into a JSON
// request object.
func expandAuthorizationPolicyRulesSources(c *Client, f *AuthorizationPolicyRulesSources) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Principals; !dcl.IsEmptyValueIndirect(v) {
		m["principals"] = v
	}
	if v := f.IPBlocks; !dcl.IsEmptyValueIndirect(v) {
		m["ipBlocks"] = v
	}

	return m, nil
}

// flattenAuthorizationPolicyRulesSources flattens an instance of AuthorizationPolicyRulesSources from a JSON
// response object.
func flattenAuthorizationPolicyRulesSources(c *Client, i interface{}) *AuthorizationPolicyRulesSources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AuthorizationPolicyRulesSources{}
	r.Principals = dcl.FlattenStringSlice(m["principals"])
	r.IPBlocks = dcl.FlattenStringSlice(m["ipBlocks"])

	return r
}

// expandAuthorizationPolicyRulesDestinationsMap expands the contents of AuthorizationPolicyRulesDestinations into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsMap(c *Client, f map[string]AuthorizationPolicyRulesDestinations) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAuthorizationPolicyRulesDestinations(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAuthorizationPolicyRulesDestinationsSlice expands the contents of AuthorizationPolicyRulesDestinations into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsSlice(c *Client, f []AuthorizationPolicyRulesDestinations) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAuthorizationPolicyRulesDestinations(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAuthorizationPolicyRulesDestinationsMap flattens the contents of AuthorizationPolicyRulesDestinations from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsMap(c *Client, i interface{}) map[string]AuthorizationPolicyRulesDestinations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AuthorizationPolicyRulesDestinations{}
	}

	if len(a) == 0 {
		return map[string]AuthorizationPolicyRulesDestinations{}
	}

	items := make(map[string]AuthorizationPolicyRulesDestinations)
	for k, item := range a {
		items[k] = *flattenAuthorizationPolicyRulesDestinations(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAuthorizationPolicyRulesDestinationsSlice flattens the contents of AuthorizationPolicyRulesDestinations from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsSlice(c *Client, i interface{}) []AuthorizationPolicyRulesDestinations {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyRulesDestinations{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyRulesDestinations{}
	}

	items := make([]AuthorizationPolicyRulesDestinations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyRulesDestinations(c, item.(map[string]interface{})))
	}

	return items
}

// expandAuthorizationPolicyRulesDestinations expands an instance of AuthorizationPolicyRulesDestinations into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinations(c *Client, f *AuthorizationPolicyRulesDestinations) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Hosts; !dcl.IsEmptyValueIndirect(v) {
		m["hosts"] = v
	}
	if v := f.Ports; !dcl.IsEmptyValueIndirect(v) {
		m["ports"] = v
	}
	if v := f.Paths; !dcl.IsEmptyValueIndirect(v) {
		m["paths"] = v
	}
	if v := f.Methods; !dcl.IsEmptyValueIndirect(v) {
		m["methods"] = v
	}
	if v, err := expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, f.HttpHeaderMatch); err != nil {
		return nil, fmt.Errorf("error expanding HttpHeaderMatch into httpHeaderMatch: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpHeaderMatch"] = v
	}

	return m, nil
}

// flattenAuthorizationPolicyRulesDestinations flattens an instance of AuthorizationPolicyRulesDestinations from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinations(c *Client, i interface{}) *AuthorizationPolicyRulesDestinations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AuthorizationPolicyRulesDestinations{}
	r.Hosts = dcl.FlattenStringSlice(m["hosts"])
	r.Ports = dcl.FlattenIntSlice(m["ports"])
	r.Paths = dcl.FlattenStringSlice(m["paths"])
	r.Methods = dcl.FlattenStringSlice(m["methods"])
	r.HttpHeaderMatch = flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, m["httpHeaderMatch"])

	return r
}

// expandAuthorizationPolicyRulesDestinationsHttpHeaderMatchMap expands the contents of AuthorizationPolicyRulesDestinationsHttpHeaderMatch into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsHttpHeaderMatchMap(c *Client, f map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice expands the contents of AuthorizationPolicyRulesDestinationsHttpHeaderMatch into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice(c *Client, f []AuthorizationPolicyRulesDestinationsHttpHeaderMatch) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatchMap flattens the contents of AuthorizationPolicyRulesDestinationsHttpHeaderMatch from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatchMap(c *Client, i interface{}) map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	}

	if len(a) == 0 {
		return map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	}

	items := make(map[string]AuthorizationPolicyRulesDestinationsHttpHeaderMatch)
	for k, item := range a {
		items[k] = *flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice flattens the contents of AuthorizationPolicyRulesDestinationsHttpHeaderMatch from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatchSlice(c *Client, i interface{}) []AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	}

	items := make([]AuthorizationPolicyRulesDestinationsHttpHeaderMatch, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c, item.(map[string]interface{})))
	}

	return items
}

// expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch expands an instance of AuthorizationPolicyRulesDestinationsHttpHeaderMatch into a JSON
// request object.
func expandAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c *Client, f *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HeaderName; !dcl.IsEmptyValueIndirect(v) {
		m["headerName"] = v
	}
	if v := f.RegexMatch; !dcl.IsEmptyValueIndirect(v) {
		m["regexMatch"] = v
	}

	return m, nil
}

// flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch flattens an instance of AuthorizationPolicyRulesDestinationsHttpHeaderMatch from a JSON
// response object.
func flattenAuthorizationPolicyRulesDestinationsHttpHeaderMatch(c *Client, i interface{}) *AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	r.HeaderName = dcl.FlattenString(m["headerName"])
	r.RegexMatch = dcl.FlattenString(m["regexMatch"])

	return r
}

// flattenAuthorizationPolicyActionEnumSlice flattens the contents of AuthorizationPolicyActionEnum from a JSON
// response object.
func flattenAuthorizationPolicyActionEnumSlice(c *Client, i interface{}) []AuthorizationPolicyActionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AuthorizationPolicyActionEnum{}
	}

	if len(a) == 0 {
		return []AuthorizationPolicyActionEnum{}
	}

	items := make([]AuthorizationPolicyActionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAuthorizationPolicyActionEnum(item.(interface{})))
	}

	return items
}

// flattenAuthorizationPolicyActionEnum asserts that an interface is a string, and returns a
// pointer to a *AuthorizationPolicyActionEnum with the same value as that string.
func flattenAuthorizationPolicyActionEnum(i interface{}) *AuthorizationPolicyActionEnum {
	s, ok := i.(string)
	if !ok {
		return AuthorizationPolicyActionEnumRef("")
	}

	return AuthorizationPolicyActionEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AuthorizationPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAuthorizationPolicy(b, c)
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
