// Copyright 2022 Google LLC. All Rights Reserved.
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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *MetadataStore) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.EncryptionSpec) {
		if err := r.EncryptionSpec.validate(); err != nil {
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
func (r *MetadataStoreEncryptionSpec) validate() error {
	if err := dcl.Required(r, "kmsKeyName"); err != nil {
		return err
	}
	return nil
}
func (r *MetadataStoreState) validate() error {
	return nil
}
func (r *MetadataStore) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-aiplatform.googleapis.com/v1beta1/", params)
}

func (r *MetadataStore) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/metadataStores/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *MetadataStore) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/metadataStores", nr.basePath(), userBasePath, params), nil

}

func (r *MetadataStore) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/metadataStores?metadataStoreId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *MetadataStore) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/metadataStores/{{name}}", nr.basePath(), userBasePath, params), nil
}

// metadataStoreApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type metadataStoreApiOperation interface {
	do(context.Context, *MetadataStore, *Client) error
}

func (c *Client) listMetadataStoreRaw(ctx context.Context, r *MetadataStore, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != MetadataStoreMaxPage {
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

type listMetadataStoreOperation struct {
	MetadataStores []map[string]interface{} `json:"metadataStores"`
	Token          string                   `json:"nextPageToken"`
}

func (c *Client) listMetadataStore(ctx context.Context, r *MetadataStore, pageToken string, pageSize int32) ([]*MetadataStore, string, error) {
	b, err := c.listMetadataStoreRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listMetadataStoreOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*MetadataStore
	for _, v := range m.MetadataStores {
		res, err := unmarshalMapMetadataStore(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllMetadataStore(ctx context.Context, f func(*MetadataStore) bool, resources []*MetadataStore) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteMetadataStore(ctx, res)
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

type deleteMetadataStoreOperation struct{}

func (op *deleteMetadataStoreOperation) do(ctx context.Context, r *MetadataStore, c *Client) error {
	r, err := c.GetMetadataStore(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "MetadataStore not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetMetadataStore checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	u, err = dcl.AddQueryParams(u, map[string]string{"force": "true"})
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
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetMetadataStore(ctx, r)
		if dcl.IsNotFound(err) {
			return nil, nil
		}
		if retriesRemaining > 0 {
			retriesRemaining--
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, dcl.NotDeletedError{ExistingResource: r}
	}, c.Config.RetryProvider)
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createMetadataStoreOperation struct {
	response map[string]interface{}
}

func (op *createMetadataStoreOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createMetadataStoreOperation) do(ctx context.Context, r *MetadataStore, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
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
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetMetadataStore(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getMetadataStoreRaw(ctx context.Context, r *MetadataStore) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
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

func (c *Client) metadataStoreDiffsForRawDesired(ctx context.Context, rawDesired *MetadataStore, opts ...dcl.ApplyOption) (initial, desired *MetadataStore, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *MetadataStore
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*MetadataStore); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected MetadataStore, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetMetadataStore(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a MetadataStore resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve MetadataStore resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that MetadataStore resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeMetadataStoreDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for MetadataStore: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for MetadataStore: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractMetadataStoreFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeMetadataStoreInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for MetadataStore: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeMetadataStoreDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for MetadataStore: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffMetadataStore(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeMetadataStoreInitialState(rawInitial, rawDesired *MetadataStore) (*MetadataStore, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeMetadataStoreDesiredState(rawDesired, rawInitial *MetadataStore, opts ...dcl.ApplyOption) (*MetadataStore, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.EncryptionSpec = canonicalizeMetadataStoreEncryptionSpec(rawDesired.EncryptionSpec, nil, opts...)
		rawDesired.State = canonicalizeMetadataStoreState(rawDesired.State, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &MetadataStore{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.EncryptionSpec = canonicalizeMetadataStoreEncryptionSpec(rawDesired.EncryptionSpec, rawInitial.EncryptionSpec, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
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

func canonicalizeMetadataStoreNewState(c *Client, rawNew, rawDesired *MetadataStore) (*MetadataStore, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
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

	if dcl.IsEmptyValueIndirect(rawNew.EncryptionSpec) && dcl.IsEmptyValueIndirect(rawDesired.EncryptionSpec) {
		rawNew.EncryptionSpec = rawDesired.EncryptionSpec
	} else {
		rawNew.EncryptionSpec = canonicalizeNewMetadataStoreEncryptionSpec(c, rawDesired.EncryptionSpec, rawNew.EncryptionSpec)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
		rawNew.State = canonicalizeNewMetadataStoreState(c, rawDesired.State, rawNew.State)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeMetadataStoreEncryptionSpec(des, initial *MetadataStoreEncryptionSpec, opts ...dcl.ApplyOption) *MetadataStoreEncryptionSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MetadataStoreEncryptionSpec{}

	if dcl.IsZeroValue(des.KmsKeyName) || (dcl.IsEmptyValueIndirect(des.KmsKeyName) && dcl.IsEmptyValueIndirect(initial.KmsKeyName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}

	return cDes
}

func canonicalizeMetadataStoreEncryptionSpecSlice(des, initial []MetadataStoreEncryptionSpec, opts ...dcl.ApplyOption) []MetadataStoreEncryptionSpec {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MetadataStoreEncryptionSpec, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMetadataStoreEncryptionSpec(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MetadataStoreEncryptionSpec, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMetadataStoreEncryptionSpec(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMetadataStoreEncryptionSpec(c *Client, des, nw *MetadataStoreEncryptionSpec) *MetadataStoreEncryptionSpec {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MetadataStoreEncryptionSpec while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewMetadataStoreEncryptionSpecSet(c *Client, des, nw []MetadataStoreEncryptionSpec) []MetadataStoreEncryptionSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []MetadataStoreEncryptionSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMetadataStoreEncryptionSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMetadataStoreEncryptionSpecSlice(c *Client, des, nw []MetadataStoreEncryptionSpec) []MetadataStoreEncryptionSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MetadataStoreEncryptionSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMetadataStoreEncryptionSpec(c, &d, &n))
	}

	return items
}

func canonicalizeMetadataStoreState(des, initial *MetadataStoreState, opts ...dcl.ApplyOption) *MetadataStoreState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MetadataStoreState{}

	if dcl.IsZeroValue(des.DiskUtilizationBytes) || (dcl.IsEmptyValueIndirect(des.DiskUtilizationBytes) && dcl.IsEmptyValueIndirect(initial.DiskUtilizationBytes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DiskUtilizationBytes = initial.DiskUtilizationBytes
	} else {
		cDes.DiskUtilizationBytes = des.DiskUtilizationBytes
	}

	return cDes
}

func canonicalizeMetadataStoreStateSlice(des, initial []MetadataStoreState, opts ...dcl.ApplyOption) []MetadataStoreState {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MetadataStoreState, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMetadataStoreState(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MetadataStoreState, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMetadataStoreState(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMetadataStoreState(c *Client, des, nw *MetadataStoreState) *MetadataStoreState {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MetadataStoreState while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewMetadataStoreStateSet(c *Client, des, nw []MetadataStoreState) []MetadataStoreState {
	if des == nil {
		return nw
	}
	var reorderedNew []MetadataStoreState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMetadataStoreStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewMetadataStoreStateSlice(c *Client, des, nw []MetadataStoreState) []MetadataStoreState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MetadataStoreState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMetadataStoreState(c, &d, &n))
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
func diffMetadataStore(c *Client, desired, actual *MetadataStore, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncryptionSpec, actual.EncryptionSpec, dcl.DiffInfo{ObjectFunction: compareMetadataStoreEncryptionSpecNewStyle, EmptyObject: EmptyMetadataStoreEncryptionSpec, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncryptionSpec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareMetadataStoreStateNewStyle, EmptyObject: EmptyMetadataStoreState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareMetadataStoreEncryptionSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MetadataStoreEncryptionSpec)
	if !ok {
		desiredNotPointer, ok := d.(MetadataStoreEncryptionSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MetadataStoreEncryptionSpec or *MetadataStoreEncryptionSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MetadataStoreEncryptionSpec)
	if !ok {
		actualNotPointer, ok := a.(MetadataStoreEncryptionSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MetadataStoreEncryptionSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMetadataStoreStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MetadataStoreState)
	if !ok {
		desiredNotPointer, ok := d.(MetadataStoreState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MetadataStoreState or *MetadataStoreState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MetadataStoreState)
	if !ok {
		actualNotPointer, ok := a.(MetadataStoreState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MetadataStoreState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DiskUtilizationBytes, actual.DiskUtilizationBytes, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskUtilizationBytes")); len(ds) != 0 || err != nil {
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
func (r *MetadataStore) urlNormalized() *MetadataStore {
	normalized := dcl.Copy(*r).(MetadataStore)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *MetadataStore) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the MetadataStore resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *MetadataStore) marshal(c *Client) ([]byte, error) {
	m, err := expandMetadataStore(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling MetadataStore: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalMetadataStore decodes JSON responses into the MetadataStore resource schema.
func unmarshalMetadataStore(b []byte, c *Client, res *MetadataStore) (*MetadataStore, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapMetadataStore(m, c, res)
}

func unmarshalMapMetadataStore(m map[string]interface{}, c *Client, res *MetadataStore) (*MetadataStore, error) {

	flattened := flattenMetadataStore(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandMetadataStore expands MetadataStore into a JSON request object.
func expandMetadataStore(c *Client, f *MetadataStore) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/metadataStores/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandMetadataStoreEncryptionSpec(c, f.EncryptionSpec, res); err != nil {
		return nil, fmt.Errorf("error expanding EncryptionSpec into encryptionSpec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["encryptionSpec"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
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

// flattenMetadataStore flattens MetadataStore from a JSON request object into the
// MetadataStore type.
func flattenMetadataStore(c *Client, i interface{}, res *MetadataStore) *MetadataStore {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &MetadataStore{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.EncryptionSpec = flattenMetadataStoreEncryptionSpec(c, m["encryptionSpec"], res)
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.State = flattenMetadataStoreState(c, m["state"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandMetadataStoreEncryptionSpecMap expands the contents of MetadataStoreEncryptionSpec into a JSON
// request object.
func expandMetadataStoreEncryptionSpecMap(c *Client, f map[string]MetadataStoreEncryptionSpec, res *MetadataStore) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMetadataStoreEncryptionSpec(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMetadataStoreEncryptionSpecSlice expands the contents of MetadataStoreEncryptionSpec into a JSON
// request object.
func expandMetadataStoreEncryptionSpecSlice(c *Client, f []MetadataStoreEncryptionSpec, res *MetadataStore) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMetadataStoreEncryptionSpec(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMetadataStoreEncryptionSpecMap flattens the contents of MetadataStoreEncryptionSpec from a JSON
// response object.
func flattenMetadataStoreEncryptionSpecMap(c *Client, i interface{}, res *MetadataStore) map[string]MetadataStoreEncryptionSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MetadataStoreEncryptionSpec{}
	}

	if len(a) == 0 {
		return map[string]MetadataStoreEncryptionSpec{}
	}

	items := make(map[string]MetadataStoreEncryptionSpec)
	for k, item := range a {
		items[k] = *flattenMetadataStoreEncryptionSpec(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMetadataStoreEncryptionSpecSlice flattens the contents of MetadataStoreEncryptionSpec from a JSON
// response object.
func flattenMetadataStoreEncryptionSpecSlice(c *Client, i interface{}, res *MetadataStore) []MetadataStoreEncryptionSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []MetadataStoreEncryptionSpec{}
	}

	if len(a) == 0 {
		return []MetadataStoreEncryptionSpec{}
	}

	items := make([]MetadataStoreEncryptionSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetadataStoreEncryptionSpec(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMetadataStoreEncryptionSpec expands an instance of MetadataStoreEncryptionSpec into a JSON
// request object.
func expandMetadataStoreEncryptionSpec(c *Client, f *MetadataStoreEncryptionSpec, res *MetadataStore) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}

	return m, nil
}

// flattenMetadataStoreEncryptionSpec flattens an instance of MetadataStoreEncryptionSpec from a JSON
// response object.
func flattenMetadataStoreEncryptionSpec(c *Client, i interface{}, res *MetadataStore) *MetadataStoreEncryptionSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MetadataStoreEncryptionSpec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMetadataStoreEncryptionSpec
	}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])

	return r
}

// expandMetadataStoreStateMap expands the contents of MetadataStoreState into a JSON
// request object.
func expandMetadataStoreStateMap(c *Client, f map[string]MetadataStoreState, res *MetadataStore) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMetadataStoreState(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMetadataStoreStateSlice expands the contents of MetadataStoreState into a JSON
// request object.
func expandMetadataStoreStateSlice(c *Client, f []MetadataStoreState, res *MetadataStore) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMetadataStoreState(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMetadataStoreStateMap flattens the contents of MetadataStoreState from a JSON
// response object.
func flattenMetadataStoreStateMap(c *Client, i interface{}, res *MetadataStore) map[string]MetadataStoreState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MetadataStoreState{}
	}

	if len(a) == 0 {
		return map[string]MetadataStoreState{}
	}

	items := make(map[string]MetadataStoreState)
	for k, item := range a {
		items[k] = *flattenMetadataStoreState(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMetadataStoreStateSlice flattens the contents of MetadataStoreState from a JSON
// response object.
func flattenMetadataStoreStateSlice(c *Client, i interface{}, res *MetadataStore) []MetadataStoreState {
	a, ok := i.([]interface{})
	if !ok {
		return []MetadataStoreState{}
	}

	if len(a) == 0 {
		return []MetadataStoreState{}
	}

	items := make([]MetadataStoreState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetadataStoreState(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMetadataStoreState expands an instance of MetadataStoreState into a JSON
// request object.
func expandMetadataStoreState(c *Client, f *MetadataStoreState, res *MetadataStore) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DiskUtilizationBytes; !dcl.IsEmptyValueIndirect(v) {
		m["diskUtilizationBytes"] = v
	}

	return m, nil
}

// flattenMetadataStoreState flattens an instance of MetadataStoreState from a JSON
// response object.
func flattenMetadataStoreState(c *Client, i interface{}, res *MetadataStore) *MetadataStoreState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MetadataStoreState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMetadataStoreState
	}
	r.DiskUtilizationBytes = dcl.FlattenInteger(m["diskUtilizationBytes"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *MetadataStore) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalMetadataStore(b, c, r)
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

type metadataStoreDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         metadataStoreApiOperation
}

func convertFieldDiffsToMetadataStoreDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]metadataStoreDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []metadataStoreDiff
	// For each operation name, create a metadataStoreDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := metadataStoreDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToMetadataStoreApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToMetadataStoreApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (metadataStoreApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractMetadataStoreFields(r *MetadataStore) error {
	vEncryptionSpec := r.EncryptionSpec
	if vEncryptionSpec == nil {
		// note: explicitly not the empty object.
		vEncryptionSpec = &MetadataStoreEncryptionSpec{}
	}
	if err := extractMetadataStoreEncryptionSpecFields(r, vEncryptionSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEncryptionSpec) {
		r.EncryptionSpec = vEncryptionSpec
	}
	vState := r.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &MetadataStoreState{}
	}
	if err := extractMetadataStoreStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		r.State = vState
	}
	return nil
}
func extractMetadataStoreEncryptionSpecFields(r *MetadataStore, o *MetadataStoreEncryptionSpec) error {
	return nil
}
func extractMetadataStoreStateFields(r *MetadataStore, o *MetadataStoreState) error {
	return nil
}

func postReadExtractMetadataStoreFields(r *MetadataStore) error {
	vEncryptionSpec := r.EncryptionSpec
	if vEncryptionSpec == nil {
		// note: explicitly not the empty object.
		vEncryptionSpec = &MetadataStoreEncryptionSpec{}
	}
	if err := postReadExtractMetadataStoreEncryptionSpecFields(r, vEncryptionSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEncryptionSpec) {
		r.EncryptionSpec = vEncryptionSpec
	}
	vState := r.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &MetadataStoreState{}
	}
	if err := postReadExtractMetadataStoreStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		r.State = vState
	}
	return nil
}
func postReadExtractMetadataStoreEncryptionSpecFields(r *MetadataStore, o *MetadataStoreEncryptionSpec) error {
	return nil
}
func postReadExtractMetadataStoreStateFields(r *MetadataStore, o *MetadataStoreState) error {
	return nil
}
