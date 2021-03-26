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
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *SslPolicy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}
func (r *SslPolicyWarning) validate() error {
	return nil
}
func (r *SslPolicyWarningData) validate() error {
	return nil
}

func sslPolicyGetURL(userBasePath string, r *SslPolicy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/sslPolicies/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func sslPolicyListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/sslPolicies", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func sslPolicyCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/sslPolicies", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func sslPolicyDeleteURL(userBasePath string, r *SslPolicy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/sslPolicies/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// sslPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type sslPolicyApiOperation interface {
	do(context.Context, *SslPolicy, *Client) error
}

// newUpdateSslPolicyPatchRequest creates a request for an
// SslPolicy resource's Patch update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateSslPolicyPatchRequest(ctx context.Context, f *SslPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Profile; !dcl.IsEmptyValueIndirect(v) {
		req["profile"] = v
	}
	if v := f.MinTlsVersion; !dcl.IsEmptyValueIndirect(v) {
		req["minTlsVersion"] = v
	}
	if v := f.CustomFeature; !dcl.IsEmptyValueIndirect(v) {
		req["customFeatures"] = v
	}
	b, err := c.getSslPolicyRaw(ctx, f.urlNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawFingerprint, err := dcl.GetMapEntry(
		m,
		[]string{"fingerprint"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["fingerprint"] = rawFingerprint.(string)
	}
	return req, nil
}

// marshalUpdateSslPolicyPatchRequest converts the update into
// the final JSON request body.
func marshalUpdateSslPolicyPatchRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateSslPolicyPatchOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateSslPolicyPatchOperation) do(ctx context.Context, r *SslPolicy, c *Client) error {
	_, err := c.GetSslPolicy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "Patch")
	if err != nil {
		return err
	}

	req, err := newUpdateSslPolicyPatchRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateSslPolicyPatchRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

func (c *Client) listSslPolicyRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := sslPolicyListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != SslPolicyMaxPage {
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

type listSslPolicyOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listSslPolicy(ctx context.Context, project, pageToken string, pageSize int32) ([]*SslPolicy, string, error) {
	b, err := c.listSslPolicyRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listSslPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*SslPolicy
	for _, v := range m.Items {
		res := flattenSslPolicy(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllSslPolicy(ctx context.Context, f func(*SslPolicy) bool, resources []*SslPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteSslPolicy(ctx, res)
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

type deleteSslPolicyOperation struct{}

func (op *deleteSslPolicyOperation) do(ctx context.Context, r *SslPolicy, c *Client) error {

	_, err := c.GetSslPolicy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("SslPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetSslPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := sslPolicyDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetSslPolicy(ctx, r.urlNormalized())
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
type createSslPolicyOperation struct {
	response map[string]interface{}
}

func (op *createSslPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createSslPolicyOperation) do(ctx context.Context, r *SslPolicy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := sslPolicyCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetSslPolicy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getSslPolicyRaw(ctx context.Context, r *SslPolicy) ([]byte, error) {
	if dcl.IsZeroValue(r.Profile) {
		r.Profile = SslPolicyProfileEnumRef("COMPATIBLE")
	}
	if dcl.IsZeroValue(r.MinTlsVersion) {
		r.MinTlsVersion = SslPolicyMinTlsVersionEnumRef("TLS_1_0")
	}

	u, err := sslPolicyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) sslPolicyDiffsForRawDesired(ctx context.Context, rawDesired *SslPolicy, opts ...dcl.ApplyOption) (initial, desired *SslPolicy, diffs []sslPolicyDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *SslPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*SslPolicy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected SslPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetSslPolicy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a SslPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve SslPolicy resource: %v", err)
		}
		c.Config.Logger.Info("Found that SslPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeSslPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for SslPolicy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for SslPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeSslPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for SslPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeSslPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for SslPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffSslPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeSslPolicyInitialState(rawInitial, rawDesired *SslPolicy) (*SslPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeSslPolicyDesiredState(rawDesired, rawInitial *SslPolicy, opts ...dcl.ApplyOption) (*SslPolicy, error) {

	if dcl.IsZeroValue(rawDesired.Profile) {
		rawDesired.Profile = SslPolicyProfileEnumRef("COMPATIBLE")
	}

	if dcl.IsZeroValue(rawDesired.MinTlsVersion) {
		rawDesired.MinTlsVersion = SslPolicyMinTlsVersionEnumRef("TLS_1_0")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Profile) {
		rawDesired.Profile = rawInitial.Profile
	}
	if dcl.IsZeroValue(rawDesired.MinTlsVersion) {
		rawDesired.MinTlsVersion = rawInitial.MinTlsVersion
	}
	if dcl.IsZeroValue(rawDesired.EnabledFeature) {
		rawDesired.EnabledFeature = rawInitial.EnabledFeature
	}
	if dcl.IsZeroValue(rawDesired.CustomFeature) {
		rawDesired.CustomFeature = rawInitial.CustomFeature
	}
	if dcl.IsZeroValue(rawDesired.Warning) {
		rawDesired.Warning = rawInitial.Warning
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeSslPolicyNewState(c *Client, rawNew, rawDesired *SslPolicy) (*SslPolicy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
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

	if dcl.IsEmptyValueIndirect(rawNew.Profile) && dcl.IsEmptyValueIndirect(rawDesired.Profile) {
		rawNew.Profile = rawDesired.Profile
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MinTlsVersion) && dcl.IsEmptyValueIndirect(rawDesired.MinTlsVersion) {
		rawNew.MinTlsVersion = rawDesired.MinTlsVersion
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnabledFeature) && dcl.IsEmptyValueIndirect(rawDesired.EnabledFeature) {
		rawNew.EnabledFeature = rawDesired.EnabledFeature
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CustomFeature) && dcl.IsEmptyValueIndirect(rawDesired.CustomFeature) {
		rawNew.CustomFeature = rawDesired.CustomFeature
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Warning) && dcl.IsEmptyValueIndirect(rawDesired.Warning) {
		rawNew.Warning = rawDesired.Warning
	} else {
		rawNew.Warning = canonicalizeNewSslPolicyWarningSlice(c, rawDesired.Warning, rawNew.Warning)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeSslPolicyWarning(des, initial *SslPolicyWarning, opts ...dcl.ApplyOption) *SslPolicyWarning {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Code, initial.Code) || dcl.IsZeroValue(des.Code) {
		des.Code = initial.Code
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		des.Message = initial.Message
	}
	if dcl.IsZeroValue(des.Data) {
		des.Data = initial.Data
	}

	return des
}

func canonicalizeNewSslPolicyWarning(c *Client, des, nw *SslPolicyWarning) *SslPolicyWarning {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Code, nw.Code) || dcl.IsZeroValue(des.Code) {
		nw.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, nw.Message) || dcl.IsZeroValue(des.Message) {
		nw.Message = des.Message
	}
	nw.Data = canonicalizeNewSslPolicyWarningDataSlice(c, des.Data, nw.Data)

	return nw
}

func canonicalizeNewSslPolicyWarningSet(c *Client, des, nw []SslPolicyWarning) []SslPolicyWarning {
	if des == nil {
		return nw
	}
	var reorderedNew []SslPolicyWarning
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareSslPolicyWarning(c, &d, &n) {
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

func canonicalizeNewSslPolicyWarningSlice(c *Client, des, nw []SslPolicyWarning) []SslPolicyWarning {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []SslPolicyWarning
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSslPolicyWarning(c, &d, &n))
	}

	return items
}

func canonicalizeSslPolicyWarningData(des, initial *SslPolicyWarningData, opts ...dcl.ApplyOption) *SslPolicyWarningData {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewSslPolicyWarningData(c *Client, des, nw *SslPolicyWarningData) *SslPolicyWarningData {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) || dcl.IsZeroValue(des.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) || dcl.IsZeroValue(des.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewSslPolicyWarningDataSet(c *Client, des, nw []SslPolicyWarningData) []SslPolicyWarningData {
	if des == nil {
		return nw
	}
	var reorderedNew []SslPolicyWarningData
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareSslPolicyWarningData(c, &d, &n) {
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

func canonicalizeNewSslPolicyWarningDataSlice(c *Client, des, nw []SslPolicyWarningData) []SslPolicyWarningData {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []SslPolicyWarningData
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSslPolicyWarningData(c, &d, &n))
	}

	return items
}

type sslPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         sslPolicyApiOperation
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
func diffSslPolicy(c *Client, desired, actual *SslPolicy, opts ...dcl.ApplyOption) ([]sslPolicyDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []sslPolicyDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, sslPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, sslPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !reflect.DeepEqual(desired.Profile, actual.Profile) {
		c.Config.Logger.Infof("Detected diff in Profile.\nDESIRED: %v\nACTUAL: %v", desired.Profile, actual.Profile)

		diffs = append(diffs, sslPolicyDiff{
			UpdateOp:  &updateSslPolicyPatchOperation{},
			FieldName: "Profile",
		})

	}
	if !reflect.DeepEqual(desired.MinTlsVersion, actual.MinTlsVersion) {
		c.Config.Logger.Infof("Detected diff in MinTlsVersion.\nDESIRED: %v\nACTUAL: %v", desired.MinTlsVersion, actual.MinTlsVersion)

		diffs = append(diffs, sslPolicyDiff{
			UpdateOp:  &updateSslPolicyPatchOperation{},
			FieldName: "MinTlsVersion",
		})

	}
	if !dcl.StringSliceEquals(desired.CustomFeature, actual.CustomFeature) {
		c.Config.Logger.Infof("Detected diff in CustomFeature.\nDESIRED: %v\nACTUAL: %v", desired.CustomFeature, actual.CustomFeature)

		toAdd, toRemove := dcl.CompareStringSets(desired.CustomFeature, actual.CustomFeature)
		c.Config.Logger.Infof("diff in CustomFeature is a set field - recomparing with set logic. \nto add: %#v\nto remove: %#v", toAdd, toRemove)
		if len(toAdd) != 0 || len(toRemove) != 0 {
			c.Config.Logger.Info("diff in CustomFeature persists after set logic analysis.")
			diffs = append(diffs, sslPolicyDiff{
				UpdateOp:  &updateSslPolicyPatchOperation{},
				FieldName: "CustomFeature",
			})
		}

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
	var deduped []sslPolicyDiff
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
func compareSslPolicyWarning(c *Client, desired, actual *SslPolicyWarning) bool {
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
	if !dcl.StringCanonicalize(desired.Code, actual.Code) && !dcl.IsZeroValue(desired.Code) {
		c.Config.Logger.Infof("Diff in Code. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Code), dcl.SprintResource(actual.Code))
		return true
	}
	if actual.Message == nil && desired.Message != nil && !dcl.IsEmptyValueIndirect(desired.Message) {
		c.Config.Logger.Infof("desired Message %s - but actually nil", dcl.SprintResource(desired.Message))
		return true
	}
	if !dcl.StringCanonicalize(desired.Message, actual.Message) && !dcl.IsZeroValue(desired.Message) {
		c.Config.Logger.Infof("Diff in Message. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Message), dcl.SprintResource(actual.Message))
		return true
	}
	if actual.Data == nil && desired.Data != nil && !dcl.IsEmptyValueIndirect(desired.Data) {
		c.Config.Logger.Infof("desired Data %s - but actually nil", dcl.SprintResource(desired.Data))
		return true
	}
	if compareSslPolicyWarningDataSlice(c, desired.Data, actual.Data) && !dcl.IsZeroValue(desired.Data) {
		c.Config.Logger.Infof("Diff in Data. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Data), dcl.SprintResource(actual.Data))
		return true
	}
	return false
}

func compareSslPolicyWarningSlice(c *Client, desired, actual []SslPolicyWarning) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in SslPolicyWarning, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareSslPolicyWarning(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in SslPolicyWarning, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareSslPolicyWarningMap(c *Client, desired, actual map[string]SslPolicyWarning) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in SslPolicyWarning, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in SslPolicyWarning, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareSslPolicyWarning(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in SslPolicyWarning, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareSslPolicyWarningData(c *Client, desired, actual *SslPolicyWarningData) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
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

func compareSslPolicyWarningDataSlice(c *Client, desired, actual []SslPolicyWarningData) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in SslPolicyWarningData, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareSslPolicyWarningData(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in SslPolicyWarningData, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareSslPolicyWarningDataMap(c *Client, desired, actual map[string]SslPolicyWarningData) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in SslPolicyWarningData, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in SslPolicyWarningData, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareSslPolicyWarningData(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in SslPolicyWarningData, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareSslPolicyProfileEnumSlice(c *Client, desired, actual []SslPolicyProfileEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in SslPolicyProfileEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareSslPolicyProfileEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in SslPolicyProfileEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareSslPolicyProfileEnum(c *Client, desired, actual *SslPolicyProfileEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareSslPolicyMinTlsVersionEnumSlice(c *Client, desired, actual []SslPolicyMinTlsVersionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in SslPolicyMinTlsVersionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareSslPolicyMinTlsVersionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in SslPolicyMinTlsVersionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareSslPolicyMinTlsVersionEnum(c *Client, desired, actual *SslPolicyMinTlsVersionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *SslPolicy) urlNormalized() *SslPolicy {
	normalized := deepcopy.Copy(*r).(SslPolicy)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *SslPolicy) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *SslPolicy) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *SslPolicy) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *SslPolicy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "Patch" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/sslPolicies/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the SslPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *SslPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandSslPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling SslPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalSslPolicy decodes JSON responses into the SslPolicy resource schema.
func unmarshalSslPolicy(b []byte, c *Client) (*SslPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapSslPolicy(m, c)
}

func unmarshalMapSslPolicy(m map[string]interface{}, c *Client) (*SslPolicy, error) {

	return flattenSslPolicy(c, m), nil
}

// expandSslPolicy expands SslPolicy into a JSON request object.
func expandSslPolicy(c *Client, f *SslPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Profile; !dcl.IsEmptyValueIndirect(v) {
		m["profile"] = v
	}
	if v := f.MinTlsVersion; !dcl.IsEmptyValueIndirect(v) {
		m["minTlsVersion"] = v
	}
	if v := f.EnabledFeature; !dcl.IsEmptyValueIndirect(v) {
		m["enabledFeatures"] = v
	}
	if v := f.CustomFeature; !dcl.IsEmptyValueIndirect(v) {
		m["customFeatures"] = v
	}
	if v, err := expandSslPolicyWarningSlice(c, f.Warning); err != nil {
		return nil, fmt.Errorf("error expanding Warning into warnings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["warnings"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenSslPolicy flattens SslPolicy from a JSON request object into the
// SslPolicy type.
func flattenSslPolicy(c *Client, i interface{}) *SslPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &SslPolicy{}
	r.Id = dcl.FlattenInteger(m["id"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.Profile = flattenSslPolicyProfileEnum(m["profile"])
	if _, ok := m["profile"]; !ok {
		c.Config.Logger.Info("Using default value for profile")
		r.Profile = SslPolicyProfileEnumRef("COMPATIBLE")
	}
	r.MinTlsVersion = flattenSslPolicyMinTlsVersionEnum(m["minTlsVersion"])
	if _, ok := m["minTlsVersion"]; !ok {
		c.Config.Logger.Info("Using default value for minTlsVersion")
		r.MinTlsVersion = SslPolicyMinTlsVersionEnumRef("TLS_1_0")
	}
	r.EnabledFeature = dcl.FlattenStringSlice(m["enabledFeatures"])
	r.CustomFeature = dcl.FlattenStringSlice(m["customFeatures"])
	r.Warning = flattenSslPolicyWarningSlice(c, m["warnings"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandSslPolicyWarningMap expands the contents of SslPolicyWarning into a JSON
// request object.
func expandSslPolicyWarningMap(c *Client, f map[string]SslPolicyWarning) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSslPolicyWarning(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSslPolicyWarningSlice expands the contents of SslPolicyWarning into a JSON
// request object.
func expandSslPolicyWarningSlice(c *Client, f []SslPolicyWarning) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSslPolicyWarning(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSslPolicyWarningMap flattens the contents of SslPolicyWarning from a JSON
// response object.
func flattenSslPolicyWarningMap(c *Client, i interface{}) map[string]SslPolicyWarning {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SslPolicyWarning{}
	}

	if len(a) == 0 {
		return map[string]SslPolicyWarning{}
	}

	items := make(map[string]SslPolicyWarning)
	for k, item := range a {
		items[k] = *flattenSslPolicyWarning(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSslPolicyWarningSlice flattens the contents of SslPolicyWarning from a JSON
// response object.
func flattenSslPolicyWarningSlice(c *Client, i interface{}) []SslPolicyWarning {
	a, ok := i.([]interface{})
	if !ok {
		return []SslPolicyWarning{}
	}

	if len(a) == 0 {
		return []SslPolicyWarning{}
	}

	items := make([]SslPolicyWarning, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSslPolicyWarning(c, item.(map[string]interface{})))
	}

	return items
}

// expandSslPolicyWarning expands an instance of SslPolicyWarning into a JSON
// request object.
func expandSslPolicyWarning(c *Client, f *SslPolicyWarning) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v, err := expandSslPolicyWarningDataSlice(c, f.Data); err != nil {
		return nil, fmt.Errorf("error expanding Data into data: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["data"] = v
	}

	return m, nil
}

// flattenSslPolicyWarning flattens an instance of SslPolicyWarning from a JSON
// response object.
func flattenSslPolicyWarning(c *Client, i interface{}) *SslPolicyWarning {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SslPolicyWarning{}
	r.Code = dcl.FlattenString(m["code"])
	r.Message = dcl.FlattenString(m["message"])
	r.Data = flattenSslPolicyWarningDataSlice(c, m["data"])

	return r
}

// expandSslPolicyWarningDataMap expands the contents of SslPolicyWarningData into a JSON
// request object.
func expandSslPolicyWarningDataMap(c *Client, f map[string]SslPolicyWarningData) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSslPolicyWarningData(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSslPolicyWarningDataSlice expands the contents of SslPolicyWarningData into a JSON
// request object.
func expandSslPolicyWarningDataSlice(c *Client, f []SslPolicyWarningData) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSslPolicyWarningData(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSslPolicyWarningDataMap flattens the contents of SslPolicyWarningData from a JSON
// response object.
func flattenSslPolicyWarningDataMap(c *Client, i interface{}) map[string]SslPolicyWarningData {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SslPolicyWarningData{}
	}

	if len(a) == 0 {
		return map[string]SslPolicyWarningData{}
	}

	items := make(map[string]SslPolicyWarningData)
	for k, item := range a {
		items[k] = *flattenSslPolicyWarningData(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSslPolicyWarningDataSlice flattens the contents of SslPolicyWarningData from a JSON
// response object.
func flattenSslPolicyWarningDataSlice(c *Client, i interface{}) []SslPolicyWarningData {
	a, ok := i.([]interface{})
	if !ok {
		return []SslPolicyWarningData{}
	}

	if len(a) == 0 {
		return []SslPolicyWarningData{}
	}

	items := make([]SslPolicyWarningData, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSslPolicyWarningData(c, item.(map[string]interface{})))
	}

	return items
}

// expandSslPolicyWarningData expands an instance of SslPolicyWarningData into a JSON
// request object.
func expandSslPolicyWarningData(c *Client, f *SslPolicyWarningData) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenSslPolicyWarningData flattens an instance of SslPolicyWarningData from a JSON
// response object.
func flattenSslPolicyWarningData(c *Client, i interface{}) *SslPolicyWarningData {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SslPolicyWarningData{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// flattenSslPolicyProfileEnumSlice flattens the contents of SslPolicyProfileEnum from a JSON
// response object.
func flattenSslPolicyProfileEnumSlice(c *Client, i interface{}) []SslPolicyProfileEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []SslPolicyProfileEnum{}
	}

	if len(a) == 0 {
		return []SslPolicyProfileEnum{}
	}

	items := make([]SslPolicyProfileEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSslPolicyProfileEnum(item.(interface{})))
	}

	return items
}

// flattenSslPolicyProfileEnum asserts that an interface is a string, and returns a
// pointer to a *SslPolicyProfileEnum with the same value as that string.
func flattenSslPolicyProfileEnum(i interface{}) *SslPolicyProfileEnum {
	s, ok := i.(string)
	if !ok {
		return SslPolicyProfileEnumRef("")
	}

	return SslPolicyProfileEnumRef(s)
}

// flattenSslPolicyMinTlsVersionEnumSlice flattens the contents of SslPolicyMinTlsVersionEnum from a JSON
// response object.
func flattenSslPolicyMinTlsVersionEnumSlice(c *Client, i interface{}) []SslPolicyMinTlsVersionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []SslPolicyMinTlsVersionEnum{}
	}

	if len(a) == 0 {
		return []SslPolicyMinTlsVersionEnum{}
	}

	items := make([]SslPolicyMinTlsVersionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSslPolicyMinTlsVersionEnum(item.(interface{})))
	}

	return items
}

// flattenSslPolicyMinTlsVersionEnum asserts that an interface is a string, and returns a
// pointer to a *SslPolicyMinTlsVersionEnum with the same value as that string.
func flattenSslPolicyMinTlsVersionEnum(i interface{}) *SslPolicyMinTlsVersionEnum {
	s, ok := i.(string)
	if !ok {
		return SslPolicyMinTlsVersionEnumRef("")
	}

	return SslPolicyMinTlsVersionEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *SslPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalSslPolicy(b, c)
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
