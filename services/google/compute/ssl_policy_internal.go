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
	"strings"
	"time"

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
	b, err := c.getSslPolicyRaw(ctx, f.URLNormalized())
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
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateSslPolicyPatchOperation) do(ctx context.Context, r *SslPolicy, c *Client) error {
	_, err := c.GetSslPolicy(ctx, r.URLNormalized())
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
		res, err := unmarshalMapSslPolicy(v, c)
		if err != nil {
			return nil, m.Token, err
		}
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
	r, err := c.GetSslPolicy(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("SslPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetSslPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := sslPolicyDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetSslPolicy(ctx, r.URLNormalized())
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

	if _, err := c.GetSslPolicy(ctx, r.URLNormalized()); err != nil {
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

	u, err := sslPolicyGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) sslPolicyDiffsForRawDesired(ctx context.Context, rawDesired *SslPolicy, opts ...dcl.ApplyOption) (initial, desired *SslPolicy, diffs []*dcl.FieldDiff, err error) {
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
	rawInitial, err := c.GetSslPolicy(ctx, fetchState.URLNormalized())
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
	canonicalDesired := &SslPolicy{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Profile) {
		canonicalDesired.Profile = rawInitial.Profile
	} else {
		canonicalDesired.Profile = rawDesired.Profile
	}
	if dcl.IsZeroValue(rawDesired.MinTlsVersion) {
		canonicalDesired.MinTlsVersion = rawInitial.MinTlsVersion
	} else {
		canonicalDesired.MinTlsVersion = rawDesired.MinTlsVersion
	}
	if dcl.IsZeroValue(rawDesired.CustomFeature) {
		canonicalDesired.CustomFeature = rawInitial.CustomFeature
	} else {
		canonicalDesired.CustomFeature = rawDesired.CustomFeature
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
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

	cDes := &SslPolicyWarning{}

	if dcl.StringCanonicalize(des.Code, initial.Code) || dcl.IsZeroValue(des.Code) {
		cDes.Code = initial.Code
	} else {
		cDes.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		cDes.Message = initial.Message
	} else {
		cDes.Message = des.Message
	}
	if dcl.IsZeroValue(des.Data) {
		des.Data = initial.Data
	} else {
		cDes.Data = des.Data
	}

	return cDes
}

func canonicalizeNewSslPolicyWarning(c *Client, des, nw *SslPolicyWarning) *SslPolicyWarning {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Code, nw.Code) {
		nw.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, nw.Message) {
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
			if diffs, _ := compareSslPolicyWarningNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &SslPolicyWarningData{}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		cDes.Key = initial.Key
	} else {
		cDes.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}

	return cDes
}

func canonicalizeNewSslPolicyWarningData(c *Client, des, nw *SslPolicyWarningData) *SslPolicyWarningData {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
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
			if diffs, _ := compareSslPolicyWarningDataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
	}

	var items []SslPolicyWarningData
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSslPolicyWarningData(c, &d, &n))
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
func diffSslPolicy(c *Client, desired, actual *SslPolicy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Profile, actual.Profile, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateSslPolicyPatchOperation")}, fn.AddNest("Profile")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinTlsVersion, actual.MinTlsVersion, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateSslPolicyPatchOperation")}, fn.AddNest("MinTlsVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnabledFeature, actual.EnabledFeature, dcl.Info{OutputOnly: true, Type: "Set", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnabledFeatures")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomFeature, actual.CustomFeature, dcl.Info{Type: "Set", OperationSelector: dcl.TriggersOperation("updateSslPolicyPatchOperation")}, fn.AddNest("CustomFeatures")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Warning, actual.Warning, dcl.Info{OutputOnly: true, ObjectFunction: compareSslPolicyWarningNewStyle, EmptyObject: EmptySslPolicyWarning, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Warnings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareSslPolicyWarningNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SslPolicyWarning)
	if !ok {
		desiredNotPointer, ok := d.(SslPolicyWarning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SslPolicyWarning or *SslPolicyWarning", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SslPolicyWarning)
	if !ok {
		actualNotPointer, ok := a.(SslPolicyWarning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SslPolicyWarning", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Data, actual.Data, dcl.Info{ObjectFunction: compareSslPolicyWarningDataNewStyle, EmptyObject: EmptySslPolicyWarningData, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Data")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareSslPolicyWarningDataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SslPolicyWarningData)
	if !ok {
		desiredNotPointer, ok := d.(SslPolicyWarningData)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SslPolicyWarningData or *SslPolicyWarningData", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SslPolicyWarningData)
	if !ok {
		actualNotPointer, ok := a.(SslPolicyWarningData)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SslPolicyWarningData", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *SslPolicy) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *SslPolicy) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *SslPolicy) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *SslPolicy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
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
	m["enabledFeatures"] = f.EnabledFeature
	m["customFeatures"] = f.CustomFeature
	if v, err := expandSslPolicyWarningSlice(c, f.Warning); err != nil {
		return nil, fmt.Errorf("error expanding Warning into warnings: %w", err)
	} else {
		m["warnings"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
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

	res := &SslPolicy{}
	res.Id = dcl.FlattenInteger(m["id"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Profile = flattenSslPolicyProfileEnum(m["profile"])
	if _, ok := m["profile"]; !ok {
		c.Config.Logger.Info("Using default value for profile")
		res.Profile = SslPolicyProfileEnumRef("COMPATIBLE")
	}
	res.MinTlsVersion = flattenSslPolicyMinTlsVersionEnum(m["minTlsVersion"])
	if _, ok := m["minTlsVersion"]; !ok {
		c.Config.Logger.Info("Using default value for minTlsVersion")
		res.MinTlsVersion = SslPolicyMinTlsVersionEnumRef("TLS_1_0")
	}
	res.EnabledFeature = dcl.FlattenStringSlice(m["enabledFeatures"])
	res.CustomFeature = dcl.FlattenStringSlice(m["customFeatures"])
	res.Warning = flattenSslPolicyWarningSlice(c, m["warnings"])
	res.Project = dcl.FlattenString(m["project"])

	return res
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySslPolicyWarning
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySslPolicyWarningData
	}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// flattenSslPolicyProfileEnumMap flattens the contents of SslPolicyProfileEnum from a JSON
// response object.
func flattenSslPolicyProfileEnumMap(c *Client, i interface{}) map[string]SslPolicyProfileEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SslPolicyProfileEnum{}
	}

	if len(a) == 0 {
		return map[string]SslPolicyProfileEnum{}
	}

	items := make(map[string]SslPolicyProfileEnum)
	for k, item := range a {
		items[k] = *flattenSslPolicyProfileEnum(item.(interface{}))
	}

	return items
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

// flattenSslPolicyMinTlsVersionEnumMap flattens the contents of SslPolicyMinTlsVersionEnum from a JSON
// response object.
func flattenSslPolicyMinTlsVersionEnumMap(c *Client, i interface{}) map[string]SslPolicyMinTlsVersionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SslPolicyMinTlsVersionEnum{}
	}

	if len(a) == 0 {
		return map[string]SslPolicyMinTlsVersionEnum{}
	}

	items := make(map[string]SslPolicyMinTlsVersionEnum)
	for k, item := range a {
		items[k] = *flattenSslPolicyMinTlsVersionEnum(item.(interface{}))
	}

	return items
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

type sslPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         sslPolicyApiOperation
}

func convertFieldDiffsToSslPolicyDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]sslPolicyDiff, error) {
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
	var diffs []sslPolicyDiff
	// For each operation name, create a sslPolicyDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := sslPolicyDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToSslPolicyApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToSslPolicyApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (sslPolicyApiOperation, error) {
	switch opName {

	case "updateSslPolicyPatchOperation":
		return &updateSslPolicyPatchOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
