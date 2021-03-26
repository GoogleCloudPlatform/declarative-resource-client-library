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
package identitytoolkit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *OAuthIdpConfig) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ResponseType) {
		if err := r.ResponseType.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OAuthIdpConfigResponseType) validate() error {
	return nil
}

func oAuthIdpConfigGetURL(userBasePath string, r *OAuthIdpConfig) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/oauthIdpConfigs/{{name}}", "https://identitytoolkit.googleapis.com/v2/", userBasePath, params), nil
}

func oAuthIdpConfigListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/oauthIdpConfigs", "https://identitytoolkit.googleapis.com/v2/", userBasePath, params), nil

}

func oAuthIdpConfigCreateURL(userBasePath, project, name string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"name":    name,
	}
	return dcl.URL("projects/{{project}}/oauthIdpConfigs?oauthIdpConfigId={{name}}", "https://identitytoolkit.googleapis.com/v2/", userBasePath, params), nil

}

func oAuthIdpConfigDeleteURL(userBasePath string, r *OAuthIdpConfig) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/oauthIdpConfigs/{{name}}", "https://identitytoolkit.googleapis.com/v2/", userBasePath, params), nil
}

// oAuthIdpConfigApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type oAuthIdpConfigApiOperation interface {
	do(context.Context, *OAuthIdpConfig, *Client) error
}

// newUpdateOAuthIdpConfigUpdateConfigRequest creates a request for an
// OAuthIdpConfig resource's UpdateConfig update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateOAuthIdpConfigUpdateConfigRequest(ctx context.Context, f *OAuthIdpConfig, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.ClientId; !dcl.IsEmptyValueIndirect(v) {
		req["clientId"] = v
	}
	if v := f.Issuer; !dcl.IsEmptyValueIndirect(v) {
		req["issuer"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		req["enabled"] = v
	}
	if v := f.ClientSecret; !dcl.IsEmptyValueIndirect(v) {
		req["clientSecret"] = v
	}
	if v, err := expandOAuthIdpConfigResponseType(c, f.ResponseType); err != nil {
		return nil, fmt.Errorf("error expanding ResponseType into responseType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["responseType"] = v
	}
	return req, nil
}

// marshalUpdateOAuthIdpConfigUpdateConfigRequest converts the update into
// the final JSON request body.
func marshalUpdateOAuthIdpConfigUpdateConfigRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateOAuthIdpConfigUpdateConfigOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateOAuthIdpConfigUpdateConfigOperation) do(ctx context.Context, r *OAuthIdpConfig, c *Client) error {
	_, err := c.GetOAuthIdpConfig(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateConfig")
	if err != nil {
		return err
	}

	req, err := newUpdateOAuthIdpConfigUpdateConfigRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateOAuthIdpConfigUpdateConfigRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listOAuthIdpConfigRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := oAuthIdpConfigListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != OAuthIdpConfigMaxPage {
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

type listOAuthIdpConfigOperation struct {
	OauthIdpConfigs []map[string]interface{} `json:"oauthIdpConfigs"`
	Token           string                   `json:"nextPageToken"`
}

func (c *Client) listOAuthIdpConfig(ctx context.Context, project, pageToken string, pageSize int32) ([]*OAuthIdpConfig, string, error) {
	b, err := c.listOAuthIdpConfigRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listOAuthIdpConfigOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*OAuthIdpConfig
	for _, v := range m.OauthIdpConfigs {
		res := flattenOAuthIdpConfig(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllOAuthIdpConfig(ctx context.Context, f func(*OAuthIdpConfig) bool, resources []*OAuthIdpConfig) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteOAuthIdpConfig(ctx, res)
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

type deleteOAuthIdpConfigOperation struct{}

func (op *deleteOAuthIdpConfigOperation) do(ctx context.Context, r *OAuthIdpConfig, c *Client) error {

	_, err := c.GetOAuthIdpConfig(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("OAuthIdpConfig not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetOAuthIdpConfig checking for existence. error: %v", err)
		return err
	}

	u, err := oAuthIdpConfigDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete OAuthIdpConfig: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetOAuthIdpConfig(ctx, r.urlNormalized())
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
type createOAuthIdpConfigOperation struct {
	response map[string]interface{}
}

func (op *createOAuthIdpConfigOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createOAuthIdpConfigOperation) do(ctx context.Context, r *OAuthIdpConfig, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, name := r.createFields()
	u, err := oAuthIdpConfigCreateURL(c.Config.BasePath, project, name)

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

	if _, err := c.GetOAuthIdpConfig(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getOAuthIdpConfigRaw(ctx context.Context, r *OAuthIdpConfig) ([]byte, error) {

	u, err := oAuthIdpConfigGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) oAuthIdpConfigDiffsForRawDesired(ctx context.Context, rawDesired *OAuthIdpConfig, opts ...dcl.ApplyOption) (initial, desired *OAuthIdpConfig, diffs []oAuthIdpConfigDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *OAuthIdpConfig
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*OAuthIdpConfig); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected OAuthIdpConfig, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetOAuthIdpConfig(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a OAuthIdpConfig resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve OAuthIdpConfig resource: %v", err)
		}
		c.Config.Logger.Info("Found that OAuthIdpConfig resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeOAuthIdpConfigDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for OAuthIdpConfig: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for OAuthIdpConfig: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeOAuthIdpConfigInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for OAuthIdpConfig: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeOAuthIdpConfigDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for OAuthIdpConfig: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffOAuthIdpConfig(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeOAuthIdpConfigInitialState(rawInitial, rawDesired *OAuthIdpConfig) (*OAuthIdpConfig, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeOAuthIdpConfigDesiredState(rawDesired, rawInitial *OAuthIdpConfig, opts ...dcl.ApplyOption) (*OAuthIdpConfig, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ResponseType = canonicalizeOAuthIdpConfigResponseType(rawDesired.ResponseType, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.ClientId, rawInitial.ClientId) {
		rawDesired.ClientId = rawInitial.ClientId
	}
	if dcl.StringCanonicalize(rawDesired.Issuer, rawInitial.Issuer) {
		rawDesired.Issuer = rawInitial.Issuer
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	if dcl.BoolCanonicalize(rawDesired.Enabled, rawInitial.Enabled) {
		rawDesired.Enabled = rawInitial.Enabled
	}
	if dcl.StringCanonicalize(rawDesired.ClientSecret, rawInitial.ClientSecret) {
		rawDesired.ClientSecret = rawInitial.ClientSecret
	}
	rawDesired.ResponseType = canonicalizeOAuthIdpConfigResponseType(rawDesired.ResponseType, rawInitial.ResponseType, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeOAuthIdpConfigNewState(c *Client, rawNew, rawDesired *OAuthIdpConfig) (*OAuthIdpConfig, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClientId) && dcl.IsEmptyValueIndirect(rawDesired.ClientId) {
		rawNew.ClientId = rawDesired.ClientId
	} else {
		if dcl.StringCanonicalize(rawDesired.ClientId, rawNew.ClientId) {
			rawNew.ClientId = rawDesired.ClientId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Issuer) && dcl.IsEmptyValueIndirect(rawDesired.Issuer) {
		rawNew.Issuer = rawDesired.Issuer
	} else {
		if dcl.StringCanonicalize(rawDesired.Issuer, rawNew.Issuer) {
			rawNew.Issuer = rawDesired.Issuer
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Enabled) && dcl.IsEmptyValueIndirect(rawDesired.Enabled) {
		rawNew.Enabled = rawDesired.Enabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Enabled, rawNew.Enabled) {
			rawNew.Enabled = rawDesired.Enabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClientSecret) && dcl.IsEmptyValueIndirect(rawDesired.ClientSecret) {
		rawNew.ClientSecret = rawDesired.ClientSecret
	} else {
		if dcl.StringCanonicalize(rawDesired.ClientSecret, rawNew.ClientSecret) {
			rawNew.ClientSecret = rawDesired.ClientSecret
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResponseType) && dcl.IsEmptyValueIndirect(rawDesired.ResponseType) {
		rawNew.ResponseType = rawDesired.ResponseType
	} else {
		rawNew.ResponseType = canonicalizeNewOAuthIdpConfigResponseType(c, rawDesired.ResponseType, rawNew.ResponseType)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeOAuthIdpConfigResponseType(des, initial *OAuthIdpConfigResponseType, opts ...dcl.ApplyOption) *OAuthIdpConfigResponseType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.IdToken, initial.IdToken) || dcl.IsZeroValue(des.IdToken) {
		des.IdToken = initial.IdToken
	}
	if dcl.BoolCanonicalize(des.Code, initial.Code) || dcl.IsZeroValue(des.Code) {
		des.Code = initial.Code
	}
	if dcl.BoolCanonicalize(des.Token, initial.Token) || dcl.IsZeroValue(des.Token) {
		des.Token = initial.Token
	}

	return des
}

func canonicalizeNewOAuthIdpConfigResponseType(c *Client, des, nw *OAuthIdpConfigResponseType) *OAuthIdpConfigResponseType {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.IdToken, nw.IdToken) || dcl.IsZeroValue(des.IdToken) {
		nw.IdToken = des.IdToken
	}
	if dcl.BoolCanonicalize(des.Code, nw.Code) || dcl.IsZeroValue(des.Code) {
		nw.Code = des.Code
	}
	if dcl.BoolCanonicalize(des.Token, nw.Token) || dcl.IsZeroValue(des.Token) {
		nw.Token = des.Token
	}

	return nw
}

func canonicalizeNewOAuthIdpConfigResponseTypeSet(c *Client, des, nw []OAuthIdpConfigResponseType) []OAuthIdpConfigResponseType {
	if des == nil {
		return nw
	}
	var reorderedNew []OAuthIdpConfigResponseType
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareOAuthIdpConfigResponseType(c, &d, &n) {
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

func canonicalizeNewOAuthIdpConfigResponseTypeSlice(c *Client, des, nw []OAuthIdpConfigResponseType) []OAuthIdpConfigResponseType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []OAuthIdpConfigResponseType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOAuthIdpConfigResponseType(c, &d, &n))
	}

	return items
}

type oAuthIdpConfigDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         oAuthIdpConfigApiOperation
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
func diffOAuthIdpConfig(c *Client, desired, actual *OAuthIdpConfig, opts ...dcl.ApplyOption) ([]oAuthIdpConfigDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []oAuthIdpConfigDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, oAuthIdpConfigDiff{
			UpdateOp:  &updateOAuthIdpConfigUpdateConfigOperation{},
			FieldName: "Name",
		})

	}
	if !dcl.IsZeroValue(desired.ClientId) && !dcl.StringCanonicalize(desired.ClientId, actual.ClientId) {
		c.Config.Logger.Infof("Detected diff in ClientId.\nDESIRED: %v\nACTUAL: %v", desired.ClientId, actual.ClientId)

		diffs = append(diffs, oAuthIdpConfigDiff{
			UpdateOp:  &updateOAuthIdpConfigUpdateConfigOperation{},
			FieldName: "ClientId",
		})

	}
	if !dcl.IsZeroValue(desired.Issuer) && !dcl.StringCanonicalize(desired.Issuer, actual.Issuer) {
		c.Config.Logger.Infof("Detected diff in Issuer.\nDESIRED: %v\nACTUAL: %v", desired.Issuer, actual.Issuer)

		diffs = append(diffs, oAuthIdpConfigDiff{
			UpdateOp:  &updateOAuthIdpConfigUpdateConfigOperation{},
			FieldName: "Issuer",
		})

	}
	if !dcl.IsZeroValue(desired.DisplayName) && !dcl.StringCanonicalize(desired.DisplayName, actual.DisplayName) {
		c.Config.Logger.Infof("Detected diff in DisplayName.\nDESIRED: %v\nACTUAL: %v", desired.DisplayName, actual.DisplayName)

		diffs = append(diffs, oAuthIdpConfigDiff{
			UpdateOp:  &updateOAuthIdpConfigUpdateConfigOperation{},
			FieldName: "DisplayName",
		})

	}
	if !dcl.IsZeroValue(desired.Enabled) && !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) {
		c.Config.Logger.Infof("Detected diff in Enabled.\nDESIRED: %v\nACTUAL: %v", desired.Enabled, actual.Enabled)

		diffs = append(diffs, oAuthIdpConfigDiff{
			UpdateOp:  &updateOAuthIdpConfigUpdateConfigOperation{},
			FieldName: "Enabled",
		})

	}
	if !dcl.IsZeroValue(desired.ClientSecret) && !dcl.StringCanonicalize(desired.ClientSecret, actual.ClientSecret) {
		c.Config.Logger.Infof("Detected diff in ClientSecret.\nDESIRED: %v\nACTUAL: %v", desired.ClientSecret, actual.ClientSecret)

		diffs = append(diffs, oAuthIdpConfigDiff{
			UpdateOp:  &updateOAuthIdpConfigUpdateConfigOperation{},
			FieldName: "ClientSecret",
		})

	}
	if compareOAuthIdpConfigResponseType(c, desired.ResponseType, actual.ResponseType) {
		c.Config.Logger.Infof("Detected diff in ResponseType.\nDESIRED: %v\nACTUAL: %v", desired.ResponseType, actual.ResponseType)

		diffs = append(diffs, oAuthIdpConfigDiff{
			UpdateOp:  &updateOAuthIdpConfigUpdateConfigOperation{},
			FieldName: "ResponseType",
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
	var deduped []oAuthIdpConfigDiff
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
func compareOAuthIdpConfigResponseType(c *Client, desired, actual *OAuthIdpConfigResponseType) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.IdToken == nil && desired.IdToken != nil && !dcl.IsEmptyValueIndirect(desired.IdToken) {
		c.Config.Logger.Infof("desired IdToken %s - but actually nil", dcl.SprintResource(desired.IdToken))
		return true
	}
	if !dcl.BoolCanonicalize(desired.IdToken, actual.IdToken) && !dcl.IsZeroValue(desired.IdToken) {
		c.Config.Logger.Infof("Diff in IdToken. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IdToken), dcl.SprintResource(actual.IdToken))
		return true
	}
	if actual.Code == nil && desired.Code != nil && !dcl.IsEmptyValueIndirect(desired.Code) {
		c.Config.Logger.Infof("desired Code %s - but actually nil", dcl.SprintResource(desired.Code))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Code, actual.Code) && !dcl.IsZeroValue(desired.Code) {
		c.Config.Logger.Infof("Diff in Code. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Code), dcl.SprintResource(actual.Code))
		return true
	}
	if actual.Token == nil && desired.Token != nil && !dcl.IsEmptyValueIndirect(desired.Token) {
		c.Config.Logger.Infof("desired Token %s - but actually nil", dcl.SprintResource(desired.Token))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Token, actual.Token) && !dcl.IsZeroValue(desired.Token) {
		c.Config.Logger.Infof("Diff in Token. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Token), dcl.SprintResource(actual.Token))
		return true
	}
	return false
}

func compareOAuthIdpConfigResponseTypeSlice(c *Client, desired, actual []OAuthIdpConfigResponseType) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OAuthIdpConfigResponseType, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareOAuthIdpConfigResponseType(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in OAuthIdpConfigResponseType, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareOAuthIdpConfigResponseTypeMap(c *Client, desired, actual map[string]OAuthIdpConfigResponseType) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OAuthIdpConfigResponseType, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in OAuthIdpConfigResponseType, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareOAuthIdpConfigResponseType(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in OAuthIdpConfigResponseType, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *OAuthIdpConfig) urlNormalized() *OAuthIdpConfig {
	normalized := deepcopy.Copy(*r).(OAuthIdpConfig)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.ClientId = dcl.SelfLinkToName(r.ClientId)
	normalized.Issuer = dcl.SelfLinkToName(r.Issuer)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.ClientSecret = dcl.SelfLinkToName(r.ClientSecret)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *OAuthIdpConfig) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *OAuthIdpConfig) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *OAuthIdpConfig) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *OAuthIdpConfig) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateConfig" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/oauthIdpConfigs/{{name}}", "https://identitytoolkit.googleapis.com/v2/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the OAuthIdpConfig resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *OAuthIdpConfig) marshal(c *Client) ([]byte, error) {
	m, err := expandOAuthIdpConfig(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling OAuthIdpConfig: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{},
	)

	return json.Marshal(m)
}

// unmarshalOAuthIdpConfig decodes JSON responses into the OAuthIdpConfig resource schema.
func unmarshalOAuthIdpConfig(b []byte, c *Client) (*OAuthIdpConfig, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapOAuthIdpConfig(m, c)
}

func unmarshalMapOAuthIdpConfig(m map[string]interface{}, c *Client) (*OAuthIdpConfig, error) {

	return flattenOAuthIdpConfig(c, m), nil
}

// expandOAuthIdpConfig expands OAuthIdpConfig into a JSON request object.
func expandOAuthIdpConfig(c *Client, f *OAuthIdpConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.ClientId; !dcl.IsEmptyValueIndirect(v) {
		m["clientId"] = v
	}
	if v := f.Issuer; !dcl.IsEmptyValueIndirect(v) {
		m["issuer"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.ClientSecret; !dcl.IsEmptyValueIndirect(v) {
		m["clientSecret"] = v
	}
	if v, err := expandOAuthIdpConfigResponseType(c, f.ResponseType); err != nil {
		return nil, fmt.Errorf("error expanding ResponseType into responseType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["responseType"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenOAuthIdpConfig flattens OAuthIdpConfig from a JSON request object into the
// OAuthIdpConfig type.
func flattenOAuthIdpConfig(c *Client, i interface{}) *OAuthIdpConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &OAuthIdpConfig{}
	r.Name = dcl.FlattenString(m["name"])
	r.ClientId = dcl.FlattenString(m["clientId"])
	r.Issuer = dcl.FlattenString(m["issuer"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.ClientSecret = dcl.FlattenString(m["clientSecret"])
	r.ResponseType = flattenOAuthIdpConfigResponseType(c, m["responseType"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandOAuthIdpConfigResponseTypeMap expands the contents of OAuthIdpConfigResponseType into a JSON
// request object.
func expandOAuthIdpConfigResponseTypeMap(c *Client, f map[string]OAuthIdpConfigResponseType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOAuthIdpConfigResponseType(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOAuthIdpConfigResponseTypeSlice expands the contents of OAuthIdpConfigResponseType into a JSON
// request object.
func expandOAuthIdpConfigResponseTypeSlice(c *Client, f []OAuthIdpConfigResponseType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOAuthIdpConfigResponseType(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOAuthIdpConfigResponseTypeMap flattens the contents of OAuthIdpConfigResponseType from a JSON
// response object.
func flattenOAuthIdpConfigResponseTypeMap(c *Client, i interface{}) map[string]OAuthIdpConfigResponseType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OAuthIdpConfigResponseType{}
	}

	if len(a) == 0 {
		return map[string]OAuthIdpConfigResponseType{}
	}

	items := make(map[string]OAuthIdpConfigResponseType)
	for k, item := range a {
		items[k] = *flattenOAuthIdpConfigResponseType(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOAuthIdpConfigResponseTypeSlice flattens the contents of OAuthIdpConfigResponseType from a JSON
// response object.
func flattenOAuthIdpConfigResponseTypeSlice(c *Client, i interface{}) []OAuthIdpConfigResponseType {
	a, ok := i.([]interface{})
	if !ok {
		return []OAuthIdpConfigResponseType{}
	}

	if len(a) == 0 {
		return []OAuthIdpConfigResponseType{}
	}

	items := make([]OAuthIdpConfigResponseType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOAuthIdpConfigResponseType(c, item.(map[string]interface{})))
	}

	return items
}

// expandOAuthIdpConfigResponseType expands an instance of OAuthIdpConfigResponseType into a JSON
// request object.
func expandOAuthIdpConfigResponseType(c *Client, f *OAuthIdpConfigResponseType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IdToken; !dcl.IsEmptyValueIndirect(v) {
		m["idToken"] = v
	}
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Token; !dcl.IsEmptyValueIndirect(v) {
		m["token"] = v
	}

	return m, nil
}

// flattenOAuthIdpConfigResponseType flattens an instance of OAuthIdpConfigResponseType from a JSON
// response object.
func flattenOAuthIdpConfigResponseType(c *Client, i interface{}) *OAuthIdpConfigResponseType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OAuthIdpConfigResponseType{}
	r.IdToken = dcl.FlattenBool(m["idToken"])
	r.Code = dcl.FlattenBool(m["code"])
	r.Token = dcl.FlattenBool(m["token"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *OAuthIdpConfig) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalOAuthIdpConfig(b, c)
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
