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
package apikeys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Key) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Restrictions) {
		if err := r.Restrictions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *KeyRestrictions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BrowserKeyRestrictions) {
		if err := r.BrowserKeyRestrictions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ServerKeyRestrictions) {
		if err := r.ServerKeyRestrictions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AndroidKeyRestrictions) {
		if err := r.AndroidKeyRestrictions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.IosKeyRestrictions) {
		if err := r.IosKeyRestrictions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *KeyRestrictionsBrowserKeyRestrictions) validate() error {
	return nil
}
func (r *KeyRestrictionsServerKeyRestrictions) validate() error {
	return nil
}
func (r *KeyRestrictionsAndroidKeyRestrictions) validate() error {
	return nil
}
func (r *KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) validate() error {
	return nil
}
func (r *KeyRestrictionsIosKeyRestrictions) validate() error {
	return nil
}
func (r *KeyRestrictionsApiTargets) validate() error {
	return nil
}

func keyGetURL(userBasePath string, r *Key) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/global/keys/{{name}}", "https://apikeys.googleapis.com/v2/", userBasePath, params), nil
}

func keyListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/locations/global/keys", "https://apikeys.googleapis.com/v2/", userBasePath, params), nil

}

func keyCreateURL(userBasePath, project, name string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"name":    name,
	}
	return dcl.URL("projects/{{project}}/locations/global/keys?keyId={{name}}", "https://apikeys.googleapis.com/v2/", userBasePath, params), nil

}

func keyDeleteURL(userBasePath string, r *Key) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/global/keys/{{name}}", "https://apikeys.googleapis.com/v2/", userBasePath, params), nil
}

// keyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type keyApiOperation interface {
	do(context.Context, *Key, *Client) error
}

// newUpdateKeyUpdateKeyRequest creates a request for an
// Key resource's UpdateKey update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateKeyUpdateKeyRequest(ctx context.Context, f *Key, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v, err := expandKeyRestrictions(c, f.Restrictions); err != nil {
		return nil, fmt.Errorf("error expanding Restrictions into restrictions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["restrictions"] = v
	}
	return req, nil
}

// marshalUpdateKeyUpdateKeyRequest converts the update into
// the final JSON request body.
func marshalUpdateKeyUpdateKeyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateKeyUpdateKeyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateKeyUpdateKeyOperation) do(ctx context.Context, r *Key, c *Client) error {
	_, err := c.GetKey(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateKey")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"displayName", "restrictions"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateKeyUpdateKeyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateKeyUpdateKeyRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://apikeys.googleapis.com/v2/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listKeyRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := keyListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != KeyMaxPage {
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

type listKeyOperation struct {
	Keys  []map[string]interface{} `json:"keys"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listKey(ctx context.Context, project, pageToken string, pageSize int32) ([]*Key, string, error) {
	b, err := c.listKeyRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listKeyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Key
	for _, v := range m.Keys {
		res := flattenKey(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllKey(ctx context.Context, f func(*Key) bool, resources []*Key) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteKey(ctx, res)
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

type deleteKeyOperation struct{}

func (op *deleteKeyOperation) do(ctx context.Context, r *Key, c *Client) error {

	_, err := c.GetKey(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Key not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetKey checking for existence. error: %v", err)
		return err
	}

	u, err := keyDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://apikeys.googleapis.com/v2/", "GET"); err != nil {
		return err
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createKeyOperation struct {
	response map[string]interface{}
}

func (op *createKeyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createKeyOperation) do(ctx context.Context, r *Key, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, name := r.createFields()
	u, err := keyCreateURL(c.Config.BasePath, project, name)

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
	if err := o.Wait(ctx, c.Config, "https://apikeys.googleapis.com/v2/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetKey(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getKeyRaw(ctx context.Context, r *Key) ([]byte, error) {

	u, err := keyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) keyDiffsForRawDesired(ctx context.Context, rawDesired *Key, opts ...dcl.ApplyOption) (initial, desired *Key, diffs []keyDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Key
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Key); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Key, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetKey(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Key resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Key resource: %v", err)
		}
		c.Config.Logger.Info("Found that Key resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeKeyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Key: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Key: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeKeyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Key: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeKeyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Key: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffKey(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeKeyInitialState(rawInitial, rawDesired *Key) (*Key, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeKeyDesiredState(rawDesired, rawInitial *Key, opts ...dcl.ApplyOption) (*Key, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Restrictions = canonicalizeKeyRestrictions(rawDesired.Restrictions, nil, opts...)

		return rawDesired, nil
	}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	rawDesired.Restrictions = canonicalizeKeyRestrictions(rawDesired.Restrictions, rawInitial.Restrictions, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeKeyNewState(c *Client, rawNew, rawDesired *Key) (*Key, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Uid) && dcl.IsEmptyValueIndirect(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.KeyString) && dcl.IsEmptyValueIndirect(rawDesired.KeyString) {
		rawNew.KeyString = rawDesired.KeyString
	} else {
		if dcl.StringCanonicalize(rawDesired.KeyString, rawNew.KeyString) {
			rawNew.KeyString = rawDesired.KeyString
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

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Restrictions) && dcl.IsEmptyValueIndirect(rawDesired.Restrictions) {
		rawNew.Restrictions = rawDesired.Restrictions
	} else {
		rawNew.Restrictions = canonicalizeNewKeyRestrictions(c, rawDesired.Restrictions, rawNew.Restrictions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeKeyRestrictions(des, initial *KeyRestrictions, opts ...dcl.ApplyOption) *KeyRestrictions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BrowserKeyRestrictions = canonicalizeKeyRestrictionsBrowserKeyRestrictions(des.BrowserKeyRestrictions, initial.BrowserKeyRestrictions, opts...)
	des.ServerKeyRestrictions = canonicalizeKeyRestrictionsServerKeyRestrictions(des.ServerKeyRestrictions, initial.ServerKeyRestrictions, opts...)
	des.AndroidKeyRestrictions = canonicalizeKeyRestrictionsAndroidKeyRestrictions(des.AndroidKeyRestrictions, initial.AndroidKeyRestrictions, opts...)
	des.IosKeyRestrictions = canonicalizeKeyRestrictionsIosKeyRestrictions(des.IosKeyRestrictions, initial.IosKeyRestrictions, opts...)
	if dcl.IsZeroValue(des.ApiTargets) {
		des.ApiTargets = initial.ApiTargets
	}

	return des
}

func canonicalizeNewKeyRestrictions(c *Client, des, nw *KeyRestrictions) *KeyRestrictions {
	if des == nil || nw == nil {
		return nw
	}

	nw.BrowserKeyRestrictions = canonicalizeNewKeyRestrictionsBrowserKeyRestrictions(c, des.BrowserKeyRestrictions, nw.BrowserKeyRestrictions)
	nw.ServerKeyRestrictions = canonicalizeNewKeyRestrictionsServerKeyRestrictions(c, des.ServerKeyRestrictions, nw.ServerKeyRestrictions)
	nw.AndroidKeyRestrictions = canonicalizeNewKeyRestrictionsAndroidKeyRestrictions(c, des.AndroidKeyRestrictions, nw.AndroidKeyRestrictions)
	nw.IosKeyRestrictions = canonicalizeNewKeyRestrictionsIosKeyRestrictions(c, des.IosKeyRestrictions, nw.IosKeyRestrictions)
	nw.ApiTargets = canonicalizeNewKeyRestrictionsApiTargetsSlice(c, des.ApiTargets, nw.ApiTargets)

	return nw
}

func canonicalizeNewKeyRestrictionsSet(c *Client, des, nw []KeyRestrictions) []KeyRestrictions {
	if des == nil {
		return nw
	}
	var reorderedNew []KeyRestrictions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKeyRestrictionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKeyRestrictionsSlice(c *Client, des, nw []KeyRestrictions) []KeyRestrictions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KeyRestrictions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKeyRestrictions(c, &d, &n))
	}

	return items
}

func canonicalizeKeyRestrictionsBrowserKeyRestrictions(des, initial *KeyRestrictionsBrowserKeyRestrictions, opts ...dcl.ApplyOption) *KeyRestrictionsBrowserKeyRestrictions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AllowedReferrers) {
		des.AllowedReferrers = initial.AllowedReferrers
	}

	return des
}

func canonicalizeNewKeyRestrictionsBrowserKeyRestrictions(c *Client, des, nw *KeyRestrictionsBrowserKeyRestrictions) *KeyRestrictionsBrowserKeyRestrictions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AllowedReferrers) {
		nw.AllowedReferrers = des.AllowedReferrers
	}

	return nw
}

func canonicalizeNewKeyRestrictionsBrowserKeyRestrictionsSet(c *Client, des, nw []KeyRestrictionsBrowserKeyRestrictions) []KeyRestrictionsBrowserKeyRestrictions {
	if des == nil {
		return nw
	}
	var reorderedNew []KeyRestrictionsBrowserKeyRestrictions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKeyRestrictionsBrowserKeyRestrictionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKeyRestrictionsBrowserKeyRestrictionsSlice(c *Client, des, nw []KeyRestrictionsBrowserKeyRestrictions) []KeyRestrictionsBrowserKeyRestrictions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KeyRestrictionsBrowserKeyRestrictions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKeyRestrictionsBrowserKeyRestrictions(c, &d, &n))
	}

	return items
}

func canonicalizeKeyRestrictionsServerKeyRestrictions(des, initial *KeyRestrictionsServerKeyRestrictions, opts ...dcl.ApplyOption) *KeyRestrictionsServerKeyRestrictions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AllowedIps) {
		des.AllowedIps = initial.AllowedIps
	}

	return des
}

func canonicalizeNewKeyRestrictionsServerKeyRestrictions(c *Client, des, nw *KeyRestrictionsServerKeyRestrictions) *KeyRestrictionsServerKeyRestrictions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AllowedIps) {
		nw.AllowedIps = des.AllowedIps
	}

	return nw
}

func canonicalizeNewKeyRestrictionsServerKeyRestrictionsSet(c *Client, des, nw []KeyRestrictionsServerKeyRestrictions) []KeyRestrictionsServerKeyRestrictions {
	if des == nil {
		return nw
	}
	var reorderedNew []KeyRestrictionsServerKeyRestrictions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKeyRestrictionsServerKeyRestrictionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKeyRestrictionsServerKeyRestrictionsSlice(c *Client, des, nw []KeyRestrictionsServerKeyRestrictions) []KeyRestrictionsServerKeyRestrictions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KeyRestrictionsServerKeyRestrictions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKeyRestrictionsServerKeyRestrictions(c, &d, &n))
	}

	return items
}

func canonicalizeKeyRestrictionsAndroidKeyRestrictions(des, initial *KeyRestrictionsAndroidKeyRestrictions, opts ...dcl.ApplyOption) *KeyRestrictionsAndroidKeyRestrictions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AllowedApplications) {
		des.AllowedApplications = initial.AllowedApplications
	}

	return des
}

func canonicalizeNewKeyRestrictionsAndroidKeyRestrictions(c *Client, des, nw *KeyRestrictionsAndroidKeyRestrictions) *KeyRestrictionsAndroidKeyRestrictions {
	if des == nil || nw == nil {
		return nw
	}

	nw.AllowedApplications = canonicalizeNewKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSlice(c, des.AllowedApplications, nw.AllowedApplications)

	return nw
}

func canonicalizeNewKeyRestrictionsAndroidKeyRestrictionsSet(c *Client, des, nw []KeyRestrictionsAndroidKeyRestrictions) []KeyRestrictionsAndroidKeyRestrictions {
	if des == nil {
		return nw
	}
	var reorderedNew []KeyRestrictionsAndroidKeyRestrictions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKeyRestrictionsAndroidKeyRestrictionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKeyRestrictionsAndroidKeyRestrictionsSlice(c *Client, des, nw []KeyRestrictionsAndroidKeyRestrictions) []KeyRestrictionsAndroidKeyRestrictions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KeyRestrictionsAndroidKeyRestrictions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKeyRestrictionsAndroidKeyRestrictions(c, &d, &n))
	}

	return items
}

func canonicalizeKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(des, initial *KeyRestrictionsAndroidKeyRestrictionsAllowedApplications, opts ...dcl.ApplyOption) *KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Sha1Fingerprint, initial.Sha1Fingerprint) || dcl.IsZeroValue(des.Sha1Fingerprint) {
		des.Sha1Fingerprint = initial.Sha1Fingerprint
	}
	if dcl.StringCanonicalize(des.PackageName, initial.PackageName) || dcl.IsZeroValue(des.PackageName) {
		des.PackageName = initial.PackageName
	}

	return des
}

func canonicalizeNewKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(c *Client, des, nw *KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) *KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Sha1Fingerprint, nw.Sha1Fingerprint) {
		nw.Sha1Fingerprint = des.Sha1Fingerprint
	}
	if dcl.StringCanonicalize(des.PackageName, nw.PackageName) {
		nw.PackageName = des.PackageName
	}

	return nw
}

func canonicalizeNewKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSet(c *Client, des, nw []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if des == nil {
		return nw
	}
	var reorderedNew []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSlice(c *Client, des, nw []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(c, &d, &n))
	}

	return items
}

func canonicalizeKeyRestrictionsIosKeyRestrictions(des, initial *KeyRestrictionsIosKeyRestrictions, opts ...dcl.ApplyOption) *KeyRestrictionsIosKeyRestrictions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AllowedBundleIds) {
		des.AllowedBundleIds = initial.AllowedBundleIds
	}

	return des
}

func canonicalizeNewKeyRestrictionsIosKeyRestrictions(c *Client, des, nw *KeyRestrictionsIosKeyRestrictions) *KeyRestrictionsIosKeyRestrictions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AllowedBundleIds) {
		nw.AllowedBundleIds = des.AllowedBundleIds
	}

	return nw
}

func canonicalizeNewKeyRestrictionsIosKeyRestrictionsSet(c *Client, des, nw []KeyRestrictionsIosKeyRestrictions) []KeyRestrictionsIosKeyRestrictions {
	if des == nil {
		return nw
	}
	var reorderedNew []KeyRestrictionsIosKeyRestrictions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKeyRestrictionsIosKeyRestrictionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKeyRestrictionsIosKeyRestrictionsSlice(c *Client, des, nw []KeyRestrictionsIosKeyRestrictions) []KeyRestrictionsIosKeyRestrictions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KeyRestrictionsIosKeyRestrictions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKeyRestrictionsIosKeyRestrictions(c, &d, &n))
	}

	return items
}

func canonicalizeKeyRestrictionsApiTargets(des, initial *KeyRestrictionsApiTargets, opts ...dcl.ApplyOption) *KeyRestrictionsApiTargets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		des.Service = initial.Service
	}
	if dcl.IsZeroValue(des.Methods) {
		des.Methods = initial.Methods
	}

	return des
}

func canonicalizeNewKeyRestrictionsApiTargets(c *Client, des, nw *KeyRestrictionsApiTargets) *KeyRestrictionsApiTargets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.IsZeroValue(nw.Methods) {
		nw.Methods = des.Methods
	}

	return nw
}

func canonicalizeNewKeyRestrictionsApiTargetsSet(c *Client, des, nw []KeyRestrictionsApiTargets) []KeyRestrictionsApiTargets {
	if des == nil {
		return nw
	}
	var reorderedNew []KeyRestrictionsApiTargets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKeyRestrictionsApiTargetsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKeyRestrictionsApiTargetsSlice(c *Client, des, nw []KeyRestrictionsApiTargets) []KeyRestrictionsApiTargets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KeyRestrictionsApiTargets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKeyRestrictionsApiTargets(c, &d, &n))
	}

	return items
}

type keyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         keyApiOperation
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
func diffKey(c *Client, desired, actual *Key, opts ...dcl.ApplyOption) ([]keyDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []keyDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKeyUpdateKeyOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.KeyString, actual.KeyString, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyString")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Restrictions, actual.Restrictions, dcl.Info{ObjectFunction: compareKeyRestrictionsNewStyle, OperationSelector: dcl.TriggersOperation("updateKeyUpdateKeyOperation")}, fn.AddNest("Restrictions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToKeyDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
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
	var deduped []keyDiff
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
func compareKeyRestrictionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KeyRestrictions)
	if !ok {
		desiredNotPointer, ok := d.(KeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictions or *KeyRestrictions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KeyRestrictions)
	if !ok {
		actualNotPointer, ok := a.(KeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BrowserKeyRestrictions, actual.BrowserKeyRestrictions, dcl.Info{ObjectFunction: compareKeyRestrictionsBrowserKeyRestrictionsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BrowserKeyRestrictions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServerKeyRestrictions, actual.ServerKeyRestrictions, dcl.Info{ObjectFunction: compareKeyRestrictionsServerKeyRestrictionsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServerKeyRestrictions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AndroidKeyRestrictions, actual.AndroidKeyRestrictions, dcl.Info{ObjectFunction: compareKeyRestrictionsAndroidKeyRestrictionsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AndroidKeyRestrictions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IosKeyRestrictions, actual.IosKeyRestrictions, dcl.Info{ObjectFunction: compareKeyRestrictionsIosKeyRestrictionsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IosKeyRestrictions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ApiTargets, actual.ApiTargets, dcl.Info{ObjectFunction: compareKeyRestrictionsApiTargetsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ApiTargets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareKeyRestrictionsBrowserKeyRestrictionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KeyRestrictionsBrowserKeyRestrictions)
	if !ok {
		desiredNotPointer, ok := d.(KeyRestrictionsBrowserKeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsBrowserKeyRestrictions or *KeyRestrictionsBrowserKeyRestrictions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KeyRestrictionsBrowserKeyRestrictions)
	if !ok {
		actualNotPointer, ok := a.(KeyRestrictionsBrowserKeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsBrowserKeyRestrictions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowedReferrers, actual.AllowedReferrers, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedReferrers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareKeyRestrictionsServerKeyRestrictionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KeyRestrictionsServerKeyRestrictions)
	if !ok {
		desiredNotPointer, ok := d.(KeyRestrictionsServerKeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsServerKeyRestrictions or *KeyRestrictionsServerKeyRestrictions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KeyRestrictionsServerKeyRestrictions)
	if !ok {
		actualNotPointer, ok := a.(KeyRestrictionsServerKeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsServerKeyRestrictions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowedIps, actual.AllowedIps, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedIps")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareKeyRestrictionsAndroidKeyRestrictionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KeyRestrictionsAndroidKeyRestrictions)
	if !ok {
		desiredNotPointer, ok := d.(KeyRestrictionsAndroidKeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsAndroidKeyRestrictions or *KeyRestrictionsAndroidKeyRestrictions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KeyRestrictionsAndroidKeyRestrictions)
	if !ok {
		actualNotPointer, ok := a.(KeyRestrictionsAndroidKeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsAndroidKeyRestrictions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowedApplications, actual.AllowedApplications, dcl.Info{ObjectFunction: compareKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedApplications")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KeyRestrictionsAndroidKeyRestrictionsAllowedApplications)
	if !ok {
		desiredNotPointer, ok := d.(KeyRestrictionsAndroidKeyRestrictionsAllowedApplications)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications or *KeyRestrictionsAndroidKeyRestrictionsAllowedApplications", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KeyRestrictionsAndroidKeyRestrictionsAllowedApplications)
	if !ok {
		actualNotPointer, ok := a.(KeyRestrictionsAndroidKeyRestrictionsAllowedApplications)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Sha1Fingerprint, actual.Sha1Fingerprint, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha1Fingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PackageName, actual.PackageName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PackageName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareKeyRestrictionsIosKeyRestrictionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KeyRestrictionsIosKeyRestrictions)
	if !ok {
		desiredNotPointer, ok := d.(KeyRestrictionsIosKeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsIosKeyRestrictions or *KeyRestrictionsIosKeyRestrictions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KeyRestrictionsIosKeyRestrictions)
	if !ok {
		actualNotPointer, ok := a.(KeyRestrictionsIosKeyRestrictions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsIosKeyRestrictions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowedBundleIds, actual.AllowedBundleIds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedBundleIds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareKeyRestrictionsApiTargetsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KeyRestrictionsApiTargets)
	if !ok {
		desiredNotPointer, ok := d.(KeyRestrictionsApiTargets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsApiTargets or *KeyRestrictionsApiTargets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KeyRestrictionsApiTargets)
	if !ok {
		actualNotPointer, ok := a.(KeyRestrictionsApiTargets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KeyRestrictionsApiTargets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Methods, actual.Methods, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Methods")); len(ds) != 0 || err != nil {
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
func (r *Key) urlNormalized() *Key {
	normalized := dcl.Copy(*r).(Key)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.KeyString = dcl.SelfLinkToName(r.KeyString)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Key) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Key) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Key) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Key) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateKey" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/global/keys/{{name}}", "https://apikeys.googleapis.com/v2/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Key resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Key) marshal(c *Client) ([]byte, error) {
	m, err := expandKey(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Key: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalKey decodes JSON responses into the Key resource schema.
func unmarshalKey(b []byte, c *Client) (*Key, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapKey(m, c)
}

func unmarshalMapKey(m map[string]interface{}, c *Client) (*Key, error) {

	return flattenKey(c, m), nil
}

// expandKey expands Key into a JSON request object.
func expandKey(c *Client, f *Key) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Uid; !dcl.IsEmptyValueIndirect(v) {
		m["uid"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.KeyString; !dcl.IsEmptyValueIndirect(v) {
		m["keyString"] = v
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
	if v, err := expandKeyRestrictions(c, f.Restrictions); err != nil {
		return nil, fmt.Errorf("error expanding Restrictions into restrictions: %w", err)
	} else if v != nil {
		m["restrictions"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenKey flattens Key from a JSON request object into the
// Key type.
func flattenKey(c *Client, i interface{}) *Key {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Key{}
	res.Name = dcl.FlattenString(m["name"])
	res.Uid = dcl.FlattenString(m["uid"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.KeyString = dcl.FlattenString(m["keyString"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.DeleteTime = dcl.FlattenString(m["deleteTime"])
	res.Restrictions = flattenKeyRestrictions(c, m["restrictions"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandKeyRestrictionsMap expands the contents of KeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsMap(c *Client, f map[string]KeyRestrictions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKeyRestrictionsSlice expands the contents of KeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsSlice(c *Client, f []KeyRestrictions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKeyRestrictionsMap flattens the contents of KeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsMap(c *Client, i interface{}) map[string]KeyRestrictions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KeyRestrictions{}
	}

	if len(a) == 0 {
		return map[string]KeyRestrictions{}
	}

	items := make(map[string]KeyRestrictions)
	for k, item := range a {
		items[k] = *flattenKeyRestrictions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKeyRestrictionsSlice flattens the contents of KeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsSlice(c *Client, i interface{}) []KeyRestrictions {
	a, ok := i.([]interface{})
	if !ok {
		return []KeyRestrictions{}
	}

	if len(a) == 0 {
		return []KeyRestrictions{}
	}

	items := make([]KeyRestrictions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKeyRestrictions(c, item.(map[string]interface{})))
	}

	return items
}

// expandKeyRestrictions expands an instance of KeyRestrictions into a JSON
// request object.
func expandKeyRestrictions(c *Client, f *KeyRestrictions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandKeyRestrictionsBrowserKeyRestrictions(c, f.BrowserKeyRestrictions); err != nil {
		return nil, fmt.Errorf("error expanding BrowserKeyRestrictions into browserKeyRestrictions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["browserKeyRestrictions"] = v
	}
	if v, err := expandKeyRestrictionsServerKeyRestrictions(c, f.ServerKeyRestrictions); err != nil {
		return nil, fmt.Errorf("error expanding ServerKeyRestrictions into serverKeyRestrictions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["serverKeyRestrictions"] = v
	}
	if v, err := expandKeyRestrictionsAndroidKeyRestrictions(c, f.AndroidKeyRestrictions); err != nil {
		return nil, fmt.Errorf("error expanding AndroidKeyRestrictions into androidKeyRestrictions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["androidKeyRestrictions"] = v
	}
	if v, err := expandKeyRestrictionsIosKeyRestrictions(c, f.IosKeyRestrictions); err != nil {
		return nil, fmt.Errorf("error expanding IosKeyRestrictions into iosKeyRestrictions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["iosKeyRestrictions"] = v
	}
	if v, err := expandKeyRestrictionsApiTargetsSlice(c, f.ApiTargets); err != nil {
		return nil, fmt.Errorf("error expanding ApiTargets into apiTargets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["apiTargets"] = v
	}

	return m, nil
}

// flattenKeyRestrictions flattens an instance of KeyRestrictions from a JSON
// response object.
func flattenKeyRestrictions(c *Client, i interface{}) *KeyRestrictions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KeyRestrictions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKeyRestrictions
	}
	r.BrowserKeyRestrictions = flattenKeyRestrictionsBrowserKeyRestrictions(c, m["browserKeyRestrictions"])
	r.ServerKeyRestrictions = flattenKeyRestrictionsServerKeyRestrictions(c, m["serverKeyRestrictions"])
	r.AndroidKeyRestrictions = flattenKeyRestrictionsAndroidKeyRestrictions(c, m["androidKeyRestrictions"])
	r.IosKeyRestrictions = flattenKeyRestrictionsIosKeyRestrictions(c, m["iosKeyRestrictions"])
	r.ApiTargets = flattenKeyRestrictionsApiTargetsSlice(c, m["apiTargets"])

	return r
}

// expandKeyRestrictionsBrowserKeyRestrictionsMap expands the contents of KeyRestrictionsBrowserKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsBrowserKeyRestrictionsMap(c *Client, f map[string]KeyRestrictionsBrowserKeyRestrictions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKeyRestrictionsBrowserKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKeyRestrictionsBrowserKeyRestrictionsSlice expands the contents of KeyRestrictionsBrowserKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsBrowserKeyRestrictionsSlice(c *Client, f []KeyRestrictionsBrowserKeyRestrictions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKeyRestrictionsBrowserKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKeyRestrictionsBrowserKeyRestrictionsMap flattens the contents of KeyRestrictionsBrowserKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsBrowserKeyRestrictionsMap(c *Client, i interface{}) map[string]KeyRestrictionsBrowserKeyRestrictions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KeyRestrictionsBrowserKeyRestrictions{}
	}

	if len(a) == 0 {
		return map[string]KeyRestrictionsBrowserKeyRestrictions{}
	}

	items := make(map[string]KeyRestrictionsBrowserKeyRestrictions)
	for k, item := range a {
		items[k] = *flattenKeyRestrictionsBrowserKeyRestrictions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKeyRestrictionsBrowserKeyRestrictionsSlice flattens the contents of KeyRestrictionsBrowserKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsBrowserKeyRestrictionsSlice(c *Client, i interface{}) []KeyRestrictionsBrowserKeyRestrictions {
	a, ok := i.([]interface{})
	if !ok {
		return []KeyRestrictionsBrowserKeyRestrictions{}
	}

	if len(a) == 0 {
		return []KeyRestrictionsBrowserKeyRestrictions{}
	}

	items := make([]KeyRestrictionsBrowserKeyRestrictions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKeyRestrictionsBrowserKeyRestrictions(c, item.(map[string]interface{})))
	}

	return items
}

// expandKeyRestrictionsBrowserKeyRestrictions expands an instance of KeyRestrictionsBrowserKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsBrowserKeyRestrictions(c *Client, f *KeyRestrictionsBrowserKeyRestrictions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowedReferrers; !dcl.IsEmptyValueIndirect(v) {
		m["allowedReferrers"] = v
	}

	return m, nil
}

// flattenKeyRestrictionsBrowserKeyRestrictions flattens an instance of KeyRestrictionsBrowserKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsBrowserKeyRestrictions(c *Client, i interface{}) *KeyRestrictionsBrowserKeyRestrictions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KeyRestrictionsBrowserKeyRestrictions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKeyRestrictionsBrowserKeyRestrictions
	}
	r.AllowedReferrers = dcl.FlattenStringSlice(m["allowedReferrers"])

	return r
}

// expandKeyRestrictionsServerKeyRestrictionsMap expands the contents of KeyRestrictionsServerKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsServerKeyRestrictionsMap(c *Client, f map[string]KeyRestrictionsServerKeyRestrictions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKeyRestrictionsServerKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKeyRestrictionsServerKeyRestrictionsSlice expands the contents of KeyRestrictionsServerKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsServerKeyRestrictionsSlice(c *Client, f []KeyRestrictionsServerKeyRestrictions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKeyRestrictionsServerKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKeyRestrictionsServerKeyRestrictionsMap flattens the contents of KeyRestrictionsServerKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsServerKeyRestrictionsMap(c *Client, i interface{}) map[string]KeyRestrictionsServerKeyRestrictions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KeyRestrictionsServerKeyRestrictions{}
	}

	if len(a) == 0 {
		return map[string]KeyRestrictionsServerKeyRestrictions{}
	}

	items := make(map[string]KeyRestrictionsServerKeyRestrictions)
	for k, item := range a {
		items[k] = *flattenKeyRestrictionsServerKeyRestrictions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKeyRestrictionsServerKeyRestrictionsSlice flattens the contents of KeyRestrictionsServerKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsServerKeyRestrictionsSlice(c *Client, i interface{}) []KeyRestrictionsServerKeyRestrictions {
	a, ok := i.([]interface{})
	if !ok {
		return []KeyRestrictionsServerKeyRestrictions{}
	}

	if len(a) == 0 {
		return []KeyRestrictionsServerKeyRestrictions{}
	}

	items := make([]KeyRestrictionsServerKeyRestrictions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKeyRestrictionsServerKeyRestrictions(c, item.(map[string]interface{})))
	}

	return items
}

// expandKeyRestrictionsServerKeyRestrictions expands an instance of KeyRestrictionsServerKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsServerKeyRestrictions(c *Client, f *KeyRestrictionsServerKeyRestrictions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowedIps; !dcl.IsEmptyValueIndirect(v) {
		m["allowedIps"] = v
	}

	return m, nil
}

// flattenKeyRestrictionsServerKeyRestrictions flattens an instance of KeyRestrictionsServerKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsServerKeyRestrictions(c *Client, i interface{}) *KeyRestrictionsServerKeyRestrictions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KeyRestrictionsServerKeyRestrictions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKeyRestrictionsServerKeyRestrictions
	}
	r.AllowedIps = dcl.FlattenStringSlice(m["allowedIps"])

	return r
}

// expandKeyRestrictionsAndroidKeyRestrictionsMap expands the contents of KeyRestrictionsAndroidKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsAndroidKeyRestrictionsMap(c *Client, f map[string]KeyRestrictionsAndroidKeyRestrictions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKeyRestrictionsAndroidKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKeyRestrictionsAndroidKeyRestrictionsSlice expands the contents of KeyRestrictionsAndroidKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsAndroidKeyRestrictionsSlice(c *Client, f []KeyRestrictionsAndroidKeyRestrictions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKeyRestrictionsAndroidKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKeyRestrictionsAndroidKeyRestrictionsMap flattens the contents of KeyRestrictionsAndroidKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsAndroidKeyRestrictionsMap(c *Client, i interface{}) map[string]KeyRestrictionsAndroidKeyRestrictions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KeyRestrictionsAndroidKeyRestrictions{}
	}

	if len(a) == 0 {
		return map[string]KeyRestrictionsAndroidKeyRestrictions{}
	}

	items := make(map[string]KeyRestrictionsAndroidKeyRestrictions)
	for k, item := range a {
		items[k] = *flattenKeyRestrictionsAndroidKeyRestrictions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKeyRestrictionsAndroidKeyRestrictionsSlice flattens the contents of KeyRestrictionsAndroidKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsAndroidKeyRestrictionsSlice(c *Client, i interface{}) []KeyRestrictionsAndroidKeyRestrictions {
	a, ok := i.([]interface{})
	if !ok {
		return []KeyRestrictionsAndroidKeyRestrictions{}
	}

	if len(a) == 0 {
		return []KeyRestrictionsAndroidKeyRestrictions{}
	}

	items := make([]KeyRestrictionsAndroidKeyRestrictions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKeyRestrictionsAndroidKeyRestrictions(c, item.(map[string]interface{})))
	}

	return items
}

// expandKeyRestrictionsAndroidKeyRestrictions expands an instance of KeyRestrictionsAndroidKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsAndroidKeyRestrictions(c *Client, f *KeyRestrictionsAndroidKeyRestrictions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSlice(c, f.AllowedApplications); err != nil {
		return nil, fmt.Errorf("error expanding AllowedApplications into allowedApplications: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["allowedApplications"] = v
	}

	return m, nil
}

// flattenKeyRestrictionsAndroidKeyRestrictions flattens an instance of KeyRestrictionsAndroidKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsAndroidKeyRestrictions(c *Client, i interface{}) *KeyRestrictionsAndroidKeyRestrictions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KeyRestrictionsAndroidKeyRestrictions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKeyRestrictionsAndroidKeyRestrictions
	}
	r.AllowedApplications = flattenKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSlice(c, m["allowedApplications"])

	return r
}

// expandKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsMap expands the contents of KeyRestrictionsAndroidKeyRestrictionsAllowedApplications into a JSON
// request object.
func expandKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsMap(c *Client, f map[string]KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSlice expands the contents of KeyRestrictionsAndroidKeyRestrictionsAllowedApplications into a JSON
// request object.
func expandKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSlice(c *Client, f []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsMap flattens the contents of KeyRestrictionsAndroidKeyRestrictionsAllowedApplications from a JSON
// response object.
func flattenKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsMap(c *Client, i interface{}) map[string]KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{}
	}

	if len(a) == 0 {
		return map[string]KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{}
	}

	items := make(map[string]KeyRestrictionsAndroidKeyRestrictionsAllowedApplications)
	for k, item := range a {
		items[k] = *flattenKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSlice flattens the contents of KeyRestrictionsAndroidKeyRestrictionsAllowedApplications from a JSON
// response object.
func flattenKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSlice(c *Client, i interface{}) []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	a, ok := i.([]interface{})
	if !ok {
		return []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{}
	}

	if len(a) == 0 {
		return []KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{}
	}

	items := make([]KeyRestrictionsAndroidKeyRestrictionsAllowedApplications, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(c, item.(map[string]interface{})))
	}

	return items
}

// expandKeyRestrictionsAndroidKeyRestrictionsAllowedApplications expands an instance of KeyRestrictionsAndroidKeyRestrictionsAllowedApplications into a JSON
// request object.
func expandKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(c *Client, f *KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Sha1Fingerprint; !dcl.IsEmptyValueIndirect(v) {
		m["sha1Fingerprint"] = v
	}
	if v := f.PackageName; !dcl.IsEmptyValueIndirect(v) {
		m["packageName"] = v
	}

	return m, nil
}

// flattenKeyRestrictionsAndroidKeyRestrictionsAllowedApplications flattens an instance of KeyRestrictionsAndroidKeyRestrictionsAllowedApplications from a JSON
// response object.
func flattenKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(c *Client, i interface{}) *KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKeyRestrictionsAndroidKeyRestrictionsAllowedApplications
	}
	r.Sha1Fingerprint = dcl.FlattenString(m["sha1Fingerprint"])
	r.PackageName = dcl.FlattenString(m["packageName"])

	return r
}

// expandKeyRestrictionsIosKeyRestrictionsMap expands the contents of KeyRestrictionsIosKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsIosKeyRestrictionsMap(c *Client, f map[string]KeyRestrictionsIosKeyRestrictions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKeyRestrictionsIosKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKeyRestrictionsIosKeyRestrictionsSlice expands the contents of KeyRestrictionsIosKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsIosKeyRestrictionsSlice(c *Client, f []KeyRestrictionsIosKeyRestrictions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKeyRestrictionsIosKeyRestrictions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKeyRestrictionsIosKeyRestrictionsMap flattens the contents of KeyRestrictionsIosKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsIosKeyRestrictionsMap(c *Client, i interface{}) map[string]KeyRestrictionsIosKeyRestrictions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KeyRestrictionsIosKeyRestrictions{}
	}

	if len(a) == 0 {
		return map[string]KeyRestrictionsIosKeyRestrictions{}
	}

	items := make(map[string]KeyRestrictionsIosKeyRestrictions)
	for k, item := range a {
		items[k] = *flattenKeyRestrictionsIosKeyRestrictions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKeyRestrictionsIosKeyRestrictionsSlice flattens the contents of KeyRestrictionsIosKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsIosKeyRestrictionsSlice(c *Client, i interface{}) []KeyRestrictionsIosKeyRestrictions {
	a, ok := i.([]interface{})
	if !ok {
		return []KeyRestrictionsIosKeyRestrictions{}
	}

	if len(a) == 0 {
		return []KeyRestrictionsIosKeyRestrictions{}
	}

	items := make([]KeyRestrictionsIosKeyRestrictions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKeyRestrictionsIosKeyRestrictions(c, item.(map[string]interface{})))
	}

	return items
}

// expandKeyRestrictionsIosKeyRestrictions expands an instance of KeyRestrictionsIosKeyRestrictions into a JSON
// request object.
func expandKeyRestrictionsIosKeyRestrictions(c *Client, f *KeyRestrictionsIosKeyRestrictions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowedBundleIds; !dcl.IsEmptyValueIndirect(v) {
		m["allowedBundleIds"] = v
	}

	return m, nil
}

// flattenKeyRestrictionsIosKeyRestrictions flattens an instance of KeyRestrictionsIosKeyRestrictions from a JSON
// response object.
func flattenKeyRestrictionsIosKeyRestrictions(c *Client, i interface{}) *KeyRestrictionsIosKeyRestrictions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KeyRestrictionsIosKeyRestrictions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKeyRestrictionsIosKeyRestrictions
	}
	r.AllowedBundleIds = dcl.FlattenStringSlice(m["allowedBundleIds"])

	return r
}

// expandKeyRestrictionsApiTargetsMap expands the contents of KeyRestrictionsApiTargets into a JSON
// request object.
func expandKeyRestrictionsApiTargetsMap(c *Client, f map[string]KeyRestrictionsApiTargets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKeyRestrictionsApiTargets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKeyRestrictionsApiTargetsSlice expands the contents of KeyRestrictionsApiTargets into a JSON
// request object.
func expandKeyRestrictionsApiTargetsSlice(c *Client, f []KeyRestrictionsApiTargets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKeyRestrictionsApiTargets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKeyRestrictionsApiTargetsMap flattens the contents of KeyRestrictionsApiTargets from a JSON
// response object.
func flattenKeyRestrictionsApiTargetsMap(c *Client, i interface{}) map[string]KeyRestrictionsApiTargets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KeyRestrictionsApiTargets{}
	}

	if len(a) == 0 {
		return map[string]KeyRestrictionsApiTargets{}
	}

	items := make(map[string]KeyRestrictionsApiTargets)
	for k, item := range a {
		items[k] = *flattenKeyRestrictionsApiTargets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKeyRestrictionsApiTargetsSlice flattens the contents of KeyRestrictionsApiTargets from a JSON
// response object.
func flattenKeyRestrictionsApiTargetsSlice(c *Client, i interface{}) []KeyRestrictionsApiTargets {
	a, ok := i.([]interface{})
	if !ok {
		return []KeyRestrictionsApiTargets{}
	}

	if len(a) == 0 {
		return []KeyRestrictionsApiTargets{}
	}

	items := make([]KeyRestrictionsApiTargets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKeyRestrictionsApiTargets(c, item.(map[string]interface{})))
	}

	return items
}

// expandKeyRestrictionsApiTargets expands an instance of KeyRestrictionsApiTargets into a JSON
// request object.
func expandKeyRestrictionsApiTargets(c *Client, f *KeyRestrictionsApiTargets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.Methods; !dcl.IsEmptyValueIndirect(v) {
		m["methods"] = v
	}

	return m, nil
}

// flattenKeyRestrictionsApiTargets flattens an instance of KeyRestrictionsApiTargets from a JSON
// response object.
func flattenKeyRestrictionsApiTargets(c *Client, i interface{}) *KeyRestrictionsApiTargets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KeyRestrictionsApiTargets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKeyRestrictionsApiTargets
	}
	r.Service = dcl.FlattenString(m["service"])
	r.Methods = dcl.FlattenStringSlice(m["methods"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Key) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalKey(b, c)
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

func convertFieldDiffToKeyDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]keyDiff, error) {
	var diffs []keyDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := keyDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameTokeyApiOperation(op, opts...)
				if err != nil {
					return nil, err
				}
				diff.UpdateOp = op
			}
			diffs = append(diffs, diff)
		}
	}
	return diffs, nil
}

func convertOpNameTokeyApiOperation(op string, opts ...dcl.ApplyOption) (keyApiOperation, error) {
	switch op {

	case "updateKeyUpdateKeyOperation":
		return &updateKeyUpdateKeyOperation{}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
