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
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *KrmApiHost) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.BundlesConfig) {
		if err := r.BundlesConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ManagementConfig) {
		if err := r.ManagementConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *KrmApiHostBundlesConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ConfigControllerConfig) {
		if err := r.ConfigControllerConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *KrmApiHostBundlesConfigConfigControllerConfig) validate() error {
	return nil
}
func (r *KrmApiHostManagementConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.StandardManagementConfig) {
		if err := r.StandardManagementConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *KrmApiHostManagementConfigStandardManagementConfig) validate() error {
	return nil
}

func krmApiHostGetURL(userBasePath string, r *KrmApiHost) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/krmApiHosts/{{name}}", "https://krmapihosting.googleapis.com/v1alpha1/", userBasePath, params), nil
}

func krmApiHostListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/krmApiHosts", "https://krmapihosting.googleapis.com/v1alpha1/", userBasePath, params), nil

}

func krmApiHostCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/krmApiHosts?krm_api_host_id={{name}}", "https://krmapihosting.googleapis.com/v1alpha1/", userBasePath, params), nil

}

func krmApiHostDeleteURL(userBasePath string, r *KrmApiHost) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/krmApiHosts/{{name}}", "https://krmapihosting.googleapis.com/v1alpha1/", userBasePath, params), nil
}

// krmApiHostApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type krmApiHostApiOperation interface {
	do(context.Context, *KrmApiHost, *Client) error
}

// newUpdateKrmApiHostUpdateKrmApiHostRequest creates a request for an
// KrmApiHost resource's UpdateKrmApiHost update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateKrmApiHostUpdateKrmApiHostRequest(ctx context.Context, f *KrmApiHost, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandKrmApiHostBundlesConfig(c, f.BundlesConfig); err != nil {
		return nil, fmt.Errorf("error expanding BundlesConfig into bundlesConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["bundlesConfig"] = v
	}
	if v := f.UsePrivateEndpoint; !dcl.IsEmptyValueIndirect(v) {
		req["usePrivateEndpoint"] = v
	}
	if v, err := expandKrmApiHostManagementConfig(c, f.ManagementConfig); err != nil {
		return nil, fmt.Errorf("error expanding ManagementConfig into managementConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["managementConfig"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/krmApiHosts/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateKrmApiHostUpdateKrmApiHostRequest converts the update into
// the final JSON request body.
func marshalUpdateKrmApiHostUpdateKrmApiHostRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateKrmApiHostUpdateKrmApiHostOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateKrmApiHostUpdateKrmApiHostOperation) do(ctx context.Context, r *KrmApiHost, c *Client) error {
	_, err := c.GetKrmApiHost(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateKrmApiHost")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateKrmApiHostUpdateKrmApiHostRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateKrmApiHostUpdateKrmApiHostRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://krmapihosting.googleapis.com/v1alpha1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listKrmApiHostRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := krmApiHostListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != KrmApiHostMaxPage {
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

type listKrmApiHostOperation struct {
	KrmApiHosts []map[string]interface{} `json:"krmApiHosts"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listKrmApiHost(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*KrmApiHost, string, error) {
	b, err := c.listKrmApiHostRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listKrmApiHostOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*KrmApiHost
	for _, v := range m.KrmApiHosts {
		res, err := unmarshalMapKrmApiHost(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllKrmApiHost(ctx context.Context, f func(*KrmApiHost) bool, resources []*KrmApiHost) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteKrmApiHost(ctx, res)
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

type deleteKrmApiHostOperation struct{}

func (op *deleteKrmApiHostOperation) do(ctx context.Context, r *KrmApiHost, c *Client) error {
	r, err := c.GetKrmApiHost(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("KrmApiHost not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetKrmApiHost checking for existence. error: %v", err)
		return err
	}

	u, err := krmApiHostDeleteURL(c.Config.BasePath, r.URLNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://krmapihosting.googleapis.com/v1alpha1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetKrmApiHost(ctx, r.URLNormalized())
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
type createKrmApiHostOperation struct {
	response map[string]interface{}
}

func (op *createKrmApiHostOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createKrmApiHostOperation) do(ctx context.Context, r *KrmApiHost, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := krmApiHostCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://krmapihosting.googleapis.com/v1alpha1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string, was %T", name)
	}
	r.Name = &name

	if _, err := c.GetKrmApiHost(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getKrmApiHostRaw(ctx context.Context, r *KrmApiHost) ([]byte, error) {

	u, err := krmApiHostGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) krmApiHostDiffsForRawDesired(ctx context.Context, rawDesired *KrmApiHost, opts ...dcl.ApplyOption) (initial, desired *KrmApiHost, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *KrmApiHost
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*KrmApiHost); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected KrmApiHost, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeKrmApiHostDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetKrmApiHost(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a KrmApiHost resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve KrmApiHost resource: %v", err)
		}
		c.Config.Logger.Info("Found that KrmApiHost resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeKrmApiHostDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for KrmApiHost: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for KrmApiHost: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeKrmApiHostInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for KrmApiHost: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeKrmApiHostDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for KrmApiHost: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffKrmApiHost(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeKrmApiHostInitialState(rawInitial, rawDesired *KrmApiHost) (*KrmApiHost, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeKrmApiHostDesiredState(rawDesired, rawInitial *KrmApiHost, opts ...dcl.ApplyOption) (*KrmApiHost, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.BundlesConfig = canonicalizeKrmApiHostBundlesConfig(rawDesired.BundlesConfig, nil, opts...)
		rawDesired.ManagementConfig = canonicalizeKrmApiHostManagementConfig(rawDesired.ManagementConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &KrmApiHost{}
	if dcl.IsZeroValue(rawDesired.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.BundlesConfig = canonicalizeKrmApiHostBundlesConfig(rawDesired.BundlesConfig, rawInitial.BundlesConfig, opts...)
	if dcl.BoolCanonicalize(rawDesired.UsePrivateEndpoint, rawInitial.UsePrivateEndpoint) {
		canonicalDesired.UsePrivateEndpoint = rawInitial.UsePrivateEndpoint
	} else {
		canonicalDesired.UsePrivateEndpoint = rawDesired.UsePrivateEndpoint
	}
	canonicalDesired.ManagementConfig = canonicalizeKrmApiHostManagementConfig(rawDesired.ManagementConfig, rawInitial.ManagementConfig, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeKrmApiHostNewState(c *Client, rawNew, rawDesired *KrmApiHost) (*KrmApiHost, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BundlesConfig) && dcl.IsEmptyValueIndirect(rawDesired.BundlesConfig) {
		rawNew.BundlesConfig = rawDesired.BundlesConfig
	} else {
		rawNew.BundlesConfig = canonicalizeNewKrmApiHostBundlesConfig(c, rawDesired.BundlesConfig, rawNew.BundlesConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.UsePrivateEndpoint) && dcl.IsEmptyValueIndirect(rawDesired.UsePrivateEndpoint) {
		rawNew.UsePrivateEndpoint = rawDesired.UsePrivateEndpoint
	} else {
		if dcl.BoolCanonicalize(rawDesired.UsePrivateEndpoint, rawNew.UsePrivateEndpoint) {
			rawNew.UsePrivateEndpoint = rawDesired.UsePrivateEndpoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GkeResourceLink) && dcl.IsEmptyValueIndirect(rawDesired.GkeResourceLink) {
		rawNew.GkeResourceLink = rawDesired.GkeResourceLink
	} else {
		if dcl.StringCanonicalize(rawDesired.GkeResourceLink, rawNew.GkeResourceLink) {
			rawNew.GkeResourceLink = rawDesired.GkeResourceLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ManagementConfig) && dcl.IsEmptyValueIndirect(rawDesired.ManagementConfig) {
		rawNew.ManagementConfig = rawDesired.ManagementConfig
	} else {
		rawNew.ManagementConfig = canonicalizeNewKrmApiHostManagementConfig(c, rawDesired.ManagementConfig, rawNew.ManagementConfig)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeKrmApiHostBundlesConfig(des, initial *KrmApiHostBundlesConfig, opts ...dcl.ApplyOption) *KrmApiHostBundlesConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &KrmApiHostBundlesConfig{}

	cDes.ConfigControllerConfig = canonicalizeKrmApiHostBundlesConfigConfigControllerConfig(des.ConfigControllerConfig, initial.ConfigControllerConfig, opts...)

	return cDes
}

func canonicalizeNewKrmApiHostBundlesConfig(c *Client, des, nw *KrmApiHostBundlesConfig) *KrmApiHostBundlesConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.ConfigControllerConfig = canonicalizeNewKrmApiHostBundlesConfigConfigControllerConfig(c, des.ConfigControllerConfig, nw.ConfigControllerConfig)

	return nw
}

func canonicalizeNewKrmApiHostBundlesConfigSet(c *Client, des, nw []KrmApiHostBundlesConfig) []KrmApiHostBundlesConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []KrmApiHostBundlesConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKrmApiHostBundlesConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKrmApiHostBundlesConfigSlice(c *Client, des, nw []KrmApiHostBundlesConfig) []KrmApiHostBundlesConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KrmApiHostBundlesConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKrmApiHostBundlesConfig(c, &d, &n))
	}

	return items
}

func canonicalizeKrmApiHostBundlesConfigConfigControllerConfig(des, initial *KrmApiHostBundlesConfigConfigControllerConfig, opts ...dcl.ApplyOption) *KrmApiHostBundlesConfigConfigControllerConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &KrmApiHostBundlesConfigConfigControllerConfig{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}

	return cDes
}

func canonicalizeNewKrmApiHostBundlesConfigConfigControllerConfig(c *Client, des, nw *KrmApiHostBundlesConfigConfigControllerConfig) *KrmApiHostBundlesConfigConfigControllerConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewKrmApiHostBundlesConfigConfigControllerConfigSet(c *Client, des, nw []KrmApiHostBundlesConfigConfigControllerConfig) []KrmApiHostBundlesConfigConfigControllerConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []KrmApiHostBundlesConfigConfigControllerConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKrmApiHostBundlesConfigConfigControllerConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKrmApiHostBundlesConfigConfigControllerConfigSlice(c *Client, des, nw []KrmApiHostBundlesConfigConfigControllerConfig) []KrmApiHostBundlesConfigConfigControllerConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KrmApiHostBundlesConfigConfigControllerConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKrmApiHostBundlesConfigConfigControllerConfig(c, &d, &n))
	}

	return items
}

func canonicalizeKrmApiHostManagementConfig(des, initial *KrmApiHostManagementConfig, opts ...dcl.ApplyOption) *KrmApiHostManagementConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &KrmApiHostManagementConfig{}

	cDes.StandardManagementConfig = canonicalizeKrmApiHostManagementConfigStandardManagementConfig(des.StandardManagementConfig, initial.StandardManagementConfig, opts...)

	return cDes
}

func canonicalizeNewKrmApiHostManagementConfig(c *Client, des, nw *KrmApiHostManagementConfig) *KrmApiHostManagementConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.StandardManagementConfig = canonicalizeNewKrmApiHostManagementConfigStandardManagementConfig(c, des.StandardManagementConfig, nw.StandardManagementConfig)

	return nw
}

func canonicalizeNewKrmApiHostManagementConfigSet(c *Client, des, nw []KrmApiHostManagementConfig) []KrmApiHostManagementConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []KrmApiHostManagementConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKrmApiHostManagementConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKrmApiHostManagementConfigSlice(c *Client, des, nw []KrmApiHostManagementConfig) []KrmApiHostManagementConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KrmApiHostManagementConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKrmApiHostManagementConfig(c, &d, &n))
	}

	return items
}

func canonicalizeKrmApiHostManagementConfigStandardManagementConfig(des, initial *KrmApiHostManagementConfigStandardManagementConfig, opts ...dcl.ApplyOption) *KrmApiHostManagementConfigStandardManagementConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &KrmApiHostManagementConfigStandardManagementConfig{}

	if dcl.StringCanonicalize(des.Network, initial.Network) || dcl.IsZeroValue(des.Network) {
		cDes.Network = initial.Network
	} else {
		cDes.Network = des.Network
	}
	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, initial.MasterIPv4CidrBlock) || dcl.IsZeroValue(des.MasterIPv4CidrBlock) {
		cDes.MasterIPv4CidrBlock = initial.MasterIPv4CidrBlock
	} else {
		cDes.MasterIPv4CidrBlock = des.MasterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ManBlock, initial.ManBlock) || dcl.IsZeroValue(des.ManBlock) {
		cDes.ManBlock = initial.ManBlock
	} else {
		cDes.ManBlock = des.ManBlock
	}
	if dcl.StringCanonicalize(des.ClusterCidrBlock, initial.ClusterCidrBlock) || dcl.IsZeroValue(des.ClusterCidrBlock) {
		cDes.ClusterCidrBlock = initial.ClusterCidrBlock
	} else {
		cDes.ClusterCidrBlock = des.ClusterCidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesCidrBlock, initial.ServicesCidrBlock) || dcl.IsZeroValue(des.ServicesCidrBlock) {
		cDes.ServicesCidrBlock = initial.ServicesCidrBlock
	} else {
		cDes.ServicesCidrBlock = des.ServicesCidrBlock
	}
	if dcl.StringCanonicalize(des.ClusterNamedRange, initial.ClusterNamedRange) || dcl.IsZeroValue(des.ClusterNamedRange) {
		cDes.ClusterNamedRange = initial.ClusterNamedRange
	} else {
		cDes.ClusterNamedRange = des.ClusterNamedRange
	}
	if dcl.StringCanonicalize(des.ServicesNamedRange, initial.ServicesNamedRange) || dcl.IsZeroValue(des.ServicesNamedRange) {
		cDes.ServicesNamedRange = initial.ServicesNamedRange
	} else {
		cDes.ServicesNamedRange = des.ServicesNamedRange
	}

	return cDes
}

func canonicalizeNewKrmApiHostManagementConfigStandardManagementConfig(c *Client, des, nw *KrmApiHostManagementConfigStandardManagementConfig) *KrmApiHostManagementConfigStandardManagementConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Network, nw.Network) {
		nw.Network = des.Network
	}
	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, nw.MasterIPv4CidrBlock) {
		nw.MasterIPv4CidrBlock = des.MasterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ManBlock, nw.ManBlock) {
		nw.ManBlock = des.ManBlock
	}
	if dcl.StringCanonicalize(des.ClusterCidrBlock, nw.ClusterCidrBlock) {
		nw.ClusterCidrBlock = des.ClusterCidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesCidrBlock, nw.ServicesCidrBlock) {
		nw.ServicesCidrBlock = des.ServicesCidrBlock
	}
	if dcl.StringCanonicalize(des.ClusterNamedRange, nw.ClusterNamedRange) {
		nw.ClusterNamedRange = des.ClusterNamedRange
	}
	if dcl.StringCanonicalize(des.ServicesNamedRange, nw.ServicesNamedRange) {
		nw.ServicesNamedRange = des.ServicesNamedRange
	}

	return nw
}

func canonicalizeNewKrmApiHostManagementConfigStandardManagementConfigSet(c *Client, des, nw []KrmApiHostManagementConfigStandardManagementConfig) []KrmApiHostManagementConfigStandardManagementConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []KrmApiHostManagementConfigStandardManagementConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareKrmApiHostManagementConfigStandardManagementConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewKrmApiHostManagementConfigStandardManagementConfigSlice(c *Client, des, nw []KrmApiHostManagementConfigStandardManagementConfig) []KrmApiHostManagementConfigStandardManagementConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []KrmApiHostManagementConfigStandardManagementConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewKrmApiHostManagementConfigStandardManagementConfig(c, &d, &n))
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
func diffKrmApiHost(c *Client, desired, actual *KrmApiHost, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BundlesConfig, actual.BundlesConfig, dcl.Info{ObjectFunction: compareKrmApiHostBundlesConfigNewStyle, EmptyObject: EmptyKrmApiHostBundlesConfig, OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("BundlesConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UsePrivateEndpoint, actual.UsePrivateEndpoint, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("UsePrivateEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GkeResourceLink, actual.GkeResourceLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GkeResourceLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManagementConfig, actual.ManagementConfig, dcl.Info{ObjectFunction: compareKrmApiHostManagementConfigNewStyle, EmptyObject: EmptyKrmApiHostManagementConfig, OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("ManagementConfig")); len(ds) != 0 || err != nil {
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
func compareKrmApiHostBundlesConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KrmApiHostBundlesConfig)
	if !ok {
		desiredNotPointer, ok := d.(KrmApiHostBundlesConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KrmApiHostBundlesConfig or *KrmApiHostBundlesConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KrmApiHostBundlesConfig)
	if !ok {
		actualNotPointer, ok := a.(KrmApiHostBundlesConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KrmApiHostBundlesConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigControllerConfig, actual.ConfigControllerConfig, dcl.Info{ObjectFunction: compareKrmApiHostBundlesConfigConfigControllerConfigNewStyle, EmptyObject: EmptyKrmApiHostBundlesConfigConfigControllerConfig, OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("ConfigControllerConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareKrmApiHostBundlesConfigConfigControllerConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KrmApiHostBundlesConfigConfigControllerConfig)
	if !ok {
		desiredNotPointer, ok := d.(KrmApiHostBundlesConfigConfigControllerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KrmApiHostBundlesConfigConfigControllerConfig or *KrmApiHostBundlesConfigConfigControllerConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KrmApiHostBundlesConfigConfigControllerConfig)
	if !ok {
		actualNotPointer, ok := a.(KrmApiHostBundlesConfigConfigControllerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KrmApiHostBundlesConfigConfigControllerConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareKrmApiHostManagementConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KrmApiHostManagementConfig)
	if !ok {
		desiredNotPointer, ok := d.(KrmApiHostManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KrmApiHostManagementConfig or *KrmApiHostManagementConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KrmApiHostManagementConfig)
	if !ok {
		actualNotPointer, ok := a.(KrmApiHostManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KrmApiHostManagementConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StandardManagementConfig, actual.StandardManagementConfig, dcl.Info{ObjectFunction: compareKrmApiHostManagementConfigStandardManagementConfigNewStyle, EmptyObject: EmptyKrmApiHostManagementConfigStandardManagementConfig, OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("StandardManagementConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareKrmApiHostManagementConfigStandardManagementConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*KrmApiHostManagementConfigStandardManagementConfig)
	if !ok {
		desiredNotPointer, ok := d.(KrmApiHostManagementConfigStandardManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KrmApiHostManagementConfigStandardManagementConfig or *KrmApiHostManagementConfigStandardManagementConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*KrmApiHostManagementConfigStandardManagementConfig)
	if !ok {
		actualNotPointer, ok := a.(KrmApiHostManagementConfigStandardManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a KrmApiHostManagementConfigStandardManagementConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MasterIPv4CidrBlock, actual.MasterIPv4CidrBlock, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("MasterIpv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManBlock, actual.ManBlock, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("ManBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterCidrBlock, actual.ClusterCidrBlock, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("ClusterCidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesCidrBlock, actual.ServicesCidrBlock, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("ServicesCidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterNamedRange, actual.ClusterNamedRange, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("ClusterNamedRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesNamedRange, actual.ServicesNamedRange, dcl.Info{OperationSelector: dcl.TriggersOperation("updateKrmApiHostUpdateKrmApiHostOperation")}, fn.AddNest("ServicesNamedRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *KrmApiHost) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *KrmApiHost) createFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *KrmApiHost) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *KrmApiHost) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "UpdateKrmApiHost" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/krmApiHosts/{{name}}", "https://krmapihosting.googleapis.com/v1alpha1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the KrmApiHost resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *KrmApiHost) marshal(c *Client) ([]byte, error) {
	m, err := expandKrmApiHost(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling KrmApiHost: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalKrmApiHost decodes JSON responses into the KrmApiHost resource schema.
func unmarshalKrmApiHost(b []byte, c *Client) (*KrmApiHost, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapKrmApiHost(m, c)
}

func unmarshalMapKrmApiHost(m map[string]interface{}, c *Client) (*KrmApiHost, error) {

	return flattenKrmApiHost(c, m), nil
}

// expandKrmApiHost expands KrmApiHost into a JSON request object.
func expandKrmApiHost(c *Client, f *KrmApiHost) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/krmApiHosts/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandKrmApiHostBundlesConfig(c, f.BundlesConfig); err != nil {
		return nil, fmt.Errorf("error expanding BundlesConfig into bundlesConfig: %w", err)
	} else if v != nil {
		m["bundlesConfig"] = v
	}
	if v := f.UsePrivateEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["usePrivateEndpoint"] = v
	}
	if v := f.GkeResourceLink; !dcl.IsEmptyValueIndirect(v) {
		m["gkeResourceLink"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := expandKrmApiHostManagementConfig(c, f.ManagementConfig); err != nil {
		return nil, fmt.Errorf("error expanding ManagementConfig into managementConfig: %w", err)
	} else if v != nil {
		m["managementConfig"] = v
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

// flattenKrmApiHost flattens KrmApiHost from a JSON request object into the
// KrmApiHost type.
func flattenKrmApiHost(c *Client, i interface{}) *KrmApiHost {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &KrmApiHost{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.BundlesConfig = flattenKrmApiHostBundlesConfig(c, m["bundlesConfig"])
	res.UsePrivateEndpoint = dcl.FlattenBool(m["usePrivateEndpoint"])
	res.GkeResourceLink = dcl.FlattenString(m["gkeResourceLink"])
	res.State = flattenKrmApiHostStateEnum(m["state"])
	res.ManagementConfig = flattenKrmApiHostManagementConfig(c, m["managementConfig"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandKrmApiHostBundlesConfigMap expands the contents of KrmApiHostBundlesConfig into a JSON
// request object.
func expandKrmApiHostBundlesConfigMap(c *Client, f map[string]KrmApiHostBundlesConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKrmApiHostBundlesConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKrmApiHostBundlesConfigSlice expands the contents of KrmApiHostBundlesConfig into a JSON
// request object.
func expandKrmApiHostBundlesConfigSlice(c *Client, f []KrmApiHostBundlesConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKrmApiHostBundlesConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKrmApiHostBundlesConfigMap flattens the contents of KrmApiHostBundlesConfig from a JSON
// response object.
func flattenKrmApiHostBundlesConfigMap(c *Client, i interface{}) map[string]KrmApiHostBundlesConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KrmApiHostBundlesConfig{}
	}

	if len(a) == 0 {
		return map[string]KrmApiHostBundlesConfig{}
	}

	items := make(map[string]KrmApiHostBundlesConfig)
	for k, item := range a {
		items[k] = *flattenKrmApiHostBundlesConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKrmApiHostBundlesConfigSlice flattens the contents of KrmApiHostBundlesConfig from a JSON
// response object.
func flattenKrmApiHostBundlesConfigSlice(c *Client, i interface{}) []KrmApiHostBundlesConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []KrmApiHostBundlesConfig{}
	}

	if len(a) == 0 {
		return []KrmApiHostBundlesConfig{}
	}

	items := make([]KrmApiHostBundlesConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKrmApiHostBundlesConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandKrmApiHostBundlesConfig expands an instance of KrmApiHostBundlesConfig into a JSON
// request object.
func expandKrmApiHostBundlesConfig(c *Client, f *KrmApiHostBundlesConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandKrmApiHostBundlesConfigConfigControllerConfig(c, f.ConfigControllerConfig); err != nil {
		return nil, fmt.Errorf("error expanding ConfigControllerConfig into configControllerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["configControllerConfig"] = v
	}

	return m, nil
}

// flattenKrmApiHostBundlesConfig flattens an instance of KrmApiHostBundlesConfig from a JSON
// response object.
func flattenKrmApiHostBundlesConfig(c *Client, i interface{}) *KrmApiHostBundlesConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KrmApiHostBundlesConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKrmApiHostBundlesConfig
	}
	r.ConfigControllerConfig = flattenKrmApiHostBundlesConfigConfigControllerConfig(c, m["configControllerConfig"])

	return r
}

// expandKrmApiHostBundlesConfigConfigControllerConfigMap expands the contents of KrmApiHostBundlesConfigConfigControllerConfig into a JSON
// request object.
func expandKrmApiHostBundlesConfigConfigControllerConfigMap(c *Client, f map[string]KrmApiHostBundlesConfigConfigControllerConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKrmApiHostBundlesConfigConfigControllerConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKrmApiHostBundlesConfigConfigControllerConfigSlice expands the contents of KrmApiHostBundlesConfigConfigControllerConfig into a JSON
// request object.
func expandKrmApiHostBundlesConfigConfigControllerConfigSlice(c *Client, f []KrmApiHostBundlesConfigConfigControllerConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKrmApiHostBundlesConfigConfigControllerConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKrmApiHostBundlesConfigConfigControllerConfigMap flattens the contents of KrmApiHostBundlesConfigConfigControllerConfig from a JSON
// response object.
func flattenKrmApiHostBundlesConfigConfigControllerConfigMap(c *Client, i interface{}) map[string]KrmApiHostBundlesConfigConfigControllerConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KrmApiHostBundlesConfigConfigControllerConfig{}
	}

	if len(a) == 0 {
		return map[string]KrmApiHostBundlesConfigConfigControllerConfig{}
	}

	items := make(map[string]KrmApiHostBundlesConfigConfigControllerConfig)
	for k, item := range a {
		items[k] = *flattenKrmApiHostBundlesConfigConfigControllerConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKrmApiHostBundlesConfigConfigControllerConfigSlice flattens the contents of KrmApiHostBundlesConfigConfigControllerConfig from a JSON
// response object.
func flattenKrmApiHostBundlesConfigConfigControllerConfigSlice(c *Client, i interface{}) []KrmApiHostBundlesConfigConfigControllerConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []KrmApiHostBundlesConfigConfigControllerConfig{}
	}

	if len(a) == 0 {
		return []KrmApiHostBundlesConfigConfigControllerConfig{}
	}

	items := make([]KrmApiHostBundlesConfigConfigControllerConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKrmApiHostBundlesConfigConfigControllerConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandKrmApiHostBundlesConfigConfigControllerConfig expands an instance of KrmApiHostBundlesConfigConfigControllerConfig into a JSON
// request object.
func expandKrmApiHostBundlesConfigConfigControllerConfig(c *Client, f *KrmApiHostBundlesConfigConfigControllerConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenKrmApiHostBundlesConfigConfigControllerConfig flattens an instance of KrmApiHostBundlesConfigConfigControllerConfig from a JSON
// response object.
func flattenKrmApiHostBundlesConfigConfigControllerConfig(c *Client, i interface{}) *KrmApiHostBundlesConfigConfigControllerConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KrmApiHostBundlesConfigConfigControllerConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKrmApiHostBundlesConfigConfigControllerConfig
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandKrmApiHostManagementConfigMap expands the contents of KrmApiHostManagementConfig into a JSON
// request object.
func expandKrmApiHostManagementConfigMap(c *Client, f map[string]KrmApiHostManagementConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKrmApiHostManagementConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKrmApiHostManagementConfigSlice expands the contents of KrmApiHostManagementConfig into a JSON
// request object.
func expandKrmApiHostManagementConfigSlice(c *Client, f []KrmApiHostManagementConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKrmApiHostManagementConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKrmApiHostManagementConfigMap flattens the contents of KrmApiHostManagementConfig from a JSON
// response object.
func flattenKrmApiHostManagementConfigMap(c *Client, i interface{}) map[string]KrmApiHostManagementConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KrmApiHostManagementConfig{}
	}

	if len(a) == 0 {
		return map[string]KrmApiHostManagementConfig{}
	}

	items := make(map[string]KrmApiHostManagementConfig)
	for k, item := range a {
		items[k] = *flattenKrmApiHostManagementConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKrmApiHostManagementConfigSlice flattens the contents of KrmApiHostManagementConfig from a JSON
// response object.
func flattenKrmApiHostManagementConfigSlice(c *Client, i interface{}) []KrmApiHostManagementConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []KrmApiHostManagementConfig{}
	}

	if len(a) == 0 {
		return []KrmApiHostManagementConfig{}
	}

	items := make([]KrmApiHostManagementConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKrmApiHostManagementConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandKrmApiHostManagementConfig expands an instance of KrmApiHostManagementConfig into a JSON
// request object.
func expandKrmApiHostManagementConfig(c *Client, f *KrmApiHostManagementConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandKrmApiHostManagementConfigStandardManagementConfig(c, f.StandardManagementConfig); err != nil {
		return nil, fmt.Errorf("error expanding StandardManagementConfig into standardManagementConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["standardManagementConfig"] = v
	}

	return m, nil
}

// flattenKrmApiHostManagementConfig flattens an instance of KrmApiHostManagementConfig from a JSON
// response object.
func flattenKrmApiHostManagementConfig(c *Client, i interface{}) *KrmApiHostManagementConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KrmApiHostManagementConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKrmApiHostManagementConfig
	}
	r.StandardManagementConfig = flattenKrmApiHostManagementConfigStandardManagementConfig(c, m["standardManagementConfig"])

	return r
}

// expandKrmApiHostManagementConfigStandardManagementConfigMap expands the contents of KrmApiHostManagementConfigStandardManagementConfig into a JSON
// request object.
func expandKrmApiHostManagementConfigStandardManagementConfigMap(c *Client, f map[string]KrmApiHostManagementConfigStandardManagementConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandKrmApiHostManagementConfigStandardManagementConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandKrmApiHostManagementConfigStandardManagementConfigSlice expands the contents of KrmApiHostManagementConfigStandardManagementConfig into a JSON
// request object.
func expandKrmApiHostManagementConfigStandardManagementConfigSlice(c *Client, f []KrmApiHostManagementConfigStandardManagementConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandKrmApiHostManagementConfigStandardManagementConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenKrmApiHostManagementConfigStandardManagementConfigMap flattens the contents of KrmApiHostManagementConfigStandardManagementConfig from a JSON
// response object.
func flattenKrmApiHostManagementConfigStandardManagementConfigMap(c *Client, i interface{}) map[string]KrmApiHostManagementConfigStandardManagementConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KrmApiHostManagementConfigStandardManagementConfig{}
	}

	if len(a) == 0 {
		return map[string]KrmApiHostManagementConfigStandardManagementConfig{}
	}

	items := make(map[string]KrmApiHostManagementConfigStandardManagementConfig)
	for k, item := range a {
		items[k] = *flattenKrmApiHostManagementConfigStandardManagementConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenKrmApiHostManagementConfigStandardManagementConfigSlice flattens the contents of KrmApiHostManagementConfigStandardManagementConfig from a JSON
// response object.
func flattenKrmApiHostManagementConfigStandardManagementConfigSlice(c *Client, i interface{}) []KrmApiHostManagementConfigStandardManagementConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []KrmApiHostManagementConfigStandardManagementConfig{}
	}

	if len(a) == 0 {
		return []KrmApiHostManagementConfigStandardManagementConfig{}
	}

	items := make([]KrmApiHostManagementConfigStandardManagementConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKrmApiHostManagementConfigStandardManagementConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandKrmApiHostManagementConfigStandardManagementConfig expands an instance of KrmApiHostManagementConfigStandardManagementConfig into a JSON
// request object.
func expandKrmApiHostManagementConfigStandardManagementConfig(c *Client, f *KrmApiHostManagementConfigStandardManagementConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.MasterIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["masterIpv4CidrBlock"] = v
	}
	if v := f.ManBlock; !dcl.IsEmptyValueIndirect(v) {
		m["manBlock"] = v
	}
	if v := f.ClusterCidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["clusterCidrBlock"] = v
	}
	if v := f.ServicesCidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["servicesCidrBlock"] = v
	}
	if v := f.ClusterNamedRange; !dcl.IsEmptyValueIndirect(v) {
		m["clusterNamedRange"] = v
	}
	if v := f.ServicesNamedRange; !dcl.IsEmptyValueIndirect(v) {
		m["servicesNamedRange"] = v
	}

	return m, nil
}

// flattenKrmApiHostManagementConfigStandardManagementConfig flattens an instance of KrmApiHostManagementConfigStandardManagementConfig from a JSON
// response object.
func flattenKrmApiHostManagementConfigStandardManagementConfig(c *Client, i interface{}) *KrmApiHostManagementConfigStandardManagementConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &KrmApiHostManagementConfigStandardManagementConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyKrmApiHostManagementConfigStandardManagementConfig
	}
	r.Network = dcl.FlattenString(m["network"])
	r.MasterIPv4CidrBlock = dcl.FlattenString(m["masterIpv4CidrBlock"])
	r.ManBlock = dcl.FlattenString(m["manBlock"])
	r.ClusterCidrBlock = dcl.FlattenString(m["clusterCidrBlock"])
	r.ServicesCidrBlock = dcl.FlattenString(m["servicesCidrBlock"])
	r.ClusterNamedRange = dcl.FlattenString(m["clusterNamedRange"])
	r.ServicesNamedRange = dcl.FlattenString(m["servicesNamedRange"])

	return r
}

// flattenKrmApiHostStateEnumMap flattens the contents of KrmApiHostStateEnum from a JSON
// response object.
func flattenKrmApiHostStateEnumMap(c *Client, i interface{}) map[string]KrmApiHostStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]KrmApiHostStateEnum{}
	}

	if len(a) == 0 {
		return map[string]KrmApiHostStateEnum{}
	}

	items := make(map[string]KrmApiHostStateEnum)
	for k, item := range a {
		items[k] = *flattenKrmApiHostStateEnum(item.(interface{}))
	}

	return items
}

// flattenKrmApiHostStateEnumSlice flattens the contents of KrmApiHostStateEnum from a JSON
// response object.
func flattenKrmApiHostStateEnumSlice(c *Client, i interface{}) []KrmApiHostStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []KrmApiHostStateEnum{}
	}

	if len(a) == 0 {
		return []KrmApiHostStateEnum{}
	}

	items := make([]KrmApiHostStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenKrmApiHostStateEnum(item.(interface{})))
	}

	return items
}

// flattenKrmApiHostStateEnum asserts that an interface is a string, and returns a
// pointer to a *KrmApiHostStateEnum with the same value as that string.
func flattenKrmApiHostStateEnum(i interface{}) *KrmApiHostStateEnum {
	s, ok := i.(string)
	if !ok {
		return KrmApiHostStateEnumRef("")
	}

	return KrmApiHostStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *KrmApiHost) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalKrmApiHost(b, c)
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

type krmApiHostDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         krmApiHostApiOperation
}

func convertFieldDiffsToKrmApiHostDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]krmApiHostDiff, error) {
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
	var diffs []krmApiHostDiff
	// For each operation name, create a krmApiHostDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := krmApiHostDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToKrmApiHostApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToKrmApiHostApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (krmApiHostApiOperation, error) {
	switch opName {

	case "updateKrmApiHostUpdateKrmApiHostOperation":
		return &updateKrmApiHostUpdateKrmApiHostOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
