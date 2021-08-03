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
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *DomainMapping) validate() error {

	if err := dcl.RequiredParameter(r.App, "App"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SslSettings) {
		if err := r.SslSettings.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DomainMappingSslSettings) validate() error {
	return nil
}
func (r *DomainMappingResourceRecords) validate() error {
	return nil
}

func domainMappingGetURL(userBasePath string, r *DomainMapping) (string, error) {
	params := map[string]interface{}{
		"app":  dcl.ValueOrEmptyString(r.App),
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("apps/{{app}}/domainMappings/{{name}}", "https://appengine.googleapis.com/v1/", userBasePath, params), nil
}

func domainMappingListURL(userBasePath, app string) (string, error) {
	params := map[string]interface{}{
		"app": app,
	}
	return dcl.URL("apps/{{app}}/domainMappings", "https://appengine.googleapis.com/v1/", userBasePath, params), nil

}

func domainMappingCreateURL(userBasePath, app string) (string, error) {
	params := map[string]interface{}{
		"app": app,
	}
	return dcl.URL("apps/{{app}}/domainMappings", "https://appengine.googleapis.com/v1/", userBasePath, params), nil

}

func domainMappingDeleteURL(userBasePath string, r *DomainMapping) (string, error) {
	params := map[string]interface{}{
		"app":  dcl.ValueOrEmptyString(r.App),
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("apps/{{app}}/domainMappings/{{name}}", "https://appengine.googleapis.com/v1/", userBasePath, params), nil
}

// domainMappingApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type domainMappingApiOperation interface {
	do(context.Context, *DomainMapping, *Client) error
}

// newUpdateDomainMappingUpdateDomainMappingRequest creates a request for an
// DomainMapping resource's UpdateDomainMapping update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDomainMappingUpdateDomainMappingRequest(ctx context.Context, f *DomainMapping, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandDomainMappingSslSettings(c, f.SslSettings); err != nil {
		return nil, fmt.Errorf("error expanding SslSettings into sslSettings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["sslSettings"] = v
	}
	return req, nil
}

// marshalUpdateDomainMappingUpdateDomainMappingRequest converts the update into
// the final JSON request body.
func marshalUpdateDomainMappingUpdateDomainMappingRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateDomainMappingUpdateDomainMappingOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDomainMappingUpdateDomainMappingOperation) do(ctx context.Context, r *DomainMapping, c *Client) error {
	_, err := c.GetDomainMapping(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateDomainMapping")
	if err != nil {
		return err
	}

	req, err := newUpdateDomainMappingUpdateDomainMappingRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateDomainMappingUpdateDomainMappingRequest(c, req)
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

func (c *Client) listDomainMappingRaw(ctx context.Context, app, pageToken string, pageSize int32) ([]byte, error) {
	u, err := domainMappingListURL(c.Config.BasePath, app)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != DomainMappingMaxPage {
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

type listDomainMappingOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listDomainMapping(ctx context.Context, app, pageToken string, pageSize int32) ([]*DomainMapping, string, error) {
	b, err := c.listDomainMappingRaw(ctx, app, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listDomainMappingOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*DomainMapping
	for _, v := range m.Items {
		res, err := unmarshalMapDomainMapping(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.App = &app
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllDomainMapping(ctx context.Context, f func(*DomainMapping) bool, resources []*DomainMapping) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteDomainMapping(ctx, res)
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

type deleteDomainMappingOperation struct{}

func (op *deleteDomainMappingOperation) do(ctx context.Context, r *DomainMapping, c *Client) error {
	r, err := c.GetDomainMapping(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("DomainMapping not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetDomainMapping checking for existence. error: %v", err)
		return err
	}

	u, err := domainMappingDeleteURL(c.Config.BasePath, r.URLNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://appengine.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetDomainMapping(ctx, r.URLNormalized())
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
type createDomainMappingOperation struct {
	response map[string]interface{}
}

func (op *createDomainMappingOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createDomainMappingOperation) do(ctx context.Context, r *DomainMapping, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	app := r.createFields()
	u, err := domainMappingCreateURL(c.Config.BasePath, app)

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

	if _, err := c.GetDomainMapping(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDomainMappingRaw(ctx context.Context, r *DomainMapping) ([]byte, error) {

	u, err := domainMappingGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) domainMappingDiffsForRawDesired(ctx context.Context, rawDesired *DomainMapping, opts ...dcl.ApplyOption) (initial, desired *DomainMapping, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *DomainMapping
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*DomainMapping); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected DomainMapping, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetDomainMapping(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a DomainMapping resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve DomainMapping resource: %v", err)
		}
		c.Config.Logger.Info("Found that DomainMapping resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeDomainMappingDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for DomainMapping: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for DomainMapping: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeDomainMappingInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for DomainMapping: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeDomainMappingDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for DomainMapping: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffDomainMapping(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeDomainMappingInitialState(rawInitial, rawDesired *DomainMapping) (*DomainMapping, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeDomainMappingDesiredState(rawDesired, rawInitial *DomainMapping, opts ...dcl.ApplyOption) (*DomainMapping, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SslSettings = canonicalizeDomainMappingSslSettings(rawDesired.SslSettings, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &DomainMapping{}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		canonicalDesired.SelfLink = rawInitial.SelfLink
	} else {
		canonicalDesired.SelfLink = rawDesired.SelfLink
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.SslSettings = canonicalizeDomainMappingSslSettings(rawDesired.SslSettings, rawInitial.SslSettings, opts...)
	if dcl.NameToSelfLink(rawDesired.App, rawInitial.App) {
		canonicalDesired.App = rawInitial.App
	} else {
		canonicalDesired.App = rawDesired.App
	}

	return canonicalDesired, nil
}

func canonicalizeDomainMappingNewState(c *Client, rawNew, rawDesired *DomainMapping) (*DomainMapping, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.SslSettings) && dcl.IsEmptyValueIndirect(rawDesired.SslSettings) {
		rawNew.SslSettings = rawDesired.SslSettings
	} else {
		rawNew.SslSettings = canonicalizeNewDomainMappingSslSettings(c, rawDesired.SslSettings, rawNew.SslSettings)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourceRecords) && dcl.IsEmptyValueIndirect(rawDesired.ResourceRecords) {
		rawNew.ResourceRecords = rawDesired.ResourceRecords
	} else {
		rawNew.ResourceRecords = canonicalizeNewDomainMappingResourceRecordsSlice(c, rawDesired.ResourceRecords, rawNew.ResourceRecords)
	}

	rawNew.App = rawDesired.App

	return rawNew, nil
}

func canonicalizeDomainMappingSslSettings(des, initial *DomainMappingSslSettings, opts ...dcl.ApplyOption) *DomainMappingSslSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DomainMappingSslSettings{}

	if dcl.StringCanonicalize(des.CertificateId, initial.CertificateId) || dcl.IsZeroValue(des.CertificateId) {
		cDes.CertificateId = initial.CertificateId
	} else {
		cDes.CertificateId = des.CertificateId
	}
	if dcl.IsZeroValue(des.SslManagementType) {
		des.SslManagementType = initial.SslManagementType
	} else {
		cDes.SslManagementType = des.SslManagementType
	}

	return cDes
}

func canonicalizeNewDomainMappingSslSettings(c *Client, des, nw *DomainMappingSslSettings) *DomainMappingSslSettings {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CertificateId, nw.CertificateId) {
		nw.CertificateId = des.CertificateId
	}
	if dcl.StringCanonicalize(des.PendingManagedCertificateId, nw.PendingManagedCertificateId) {
		nw.PendingManagedCertificateId = des.PendingManagedCertificateId
	}

	return nw
}

func canonicalizeNewDomainMappingSslSettingsSet(c *Client, des, nw []DomainMappingSslSettings) []DomainMappingSslSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []DomainMappingSslSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDomainMappingSslSettingsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDomainMappingSslSettingsSlice(c *Client, des, nw []DomainMappingSslSettings) []DomainMappingSslSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DomainMappingSslSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDomainMappingSslSettings(c, &d, &n))
	}

	return items
}

func canonicalizeDomainMappingResourceRecords(des, initial *DomainMappingResourceRecords, opts ...dcl.ApplyOption) *DomainMappingResourceRecords {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DomainMappingResourceRecords{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Rrdata, initial.Rrdata) || dcl.IsZeroValue(des.Rrdata) {
		cDes.Rrdata = initial.Rrdata
	} else {
		cDes.Rrdata = des.Rrdata
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}

	return cDes
}

func canonicalizeNewDomainMappingResourceRecords(c *Client, des, nw *DomainMappingResourceRecords) *DomainMappingResourceRecords {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Rrdata, nw.Rrdata) {
		nw.Rrdata = des.Rrdata
	}

	return nw
}

func canonicalizeNewDomainMappingResourceRecordsSet(c *Client, des, nw []DomainMappingResourceRecords) []DomainMappingResourceRecords {
	if des == nil {
		return nw
	}
	var reorderedNew []DomainMappingResourceRecords
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDomainMappingResourceRecordsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDomainMappingResourceRecordsSlice(c *Client, des, nw []DomainMappingResourceRecords) []DomainMappingResourceRecords {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DomainMappingResourceRecords
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDomainMappingResourceRecords(c, &d, &n))
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
func diffDomainMapping(c *Client, desired, actual *DomainMapping, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SslSettings, actual.SslSettings, dcl.Info{ObjectFunction: compareDomainMappingSslSettingsNewStyle, EmptyObject: EmptyDomainMappingSslSettings, OperationSelector: dcl.TriggersOperation("updateDomainMappingUpdateDomainMappingOperation")}, fn.AddNest("SslSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceRecords, actual.ResourceRecords, dcl.Info{OutputOnly: true, ObjectFunction: compareDomainMappingResourceRecordsNewStyle, EmptyObject: EmptyDomainMappingResourceRecords, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceRecords")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.App, actual.App, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("App")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareDomainMappingSslSettingsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DomainMappingSslSettings)
	if !ok {
		desiredNotPointer, ok := d.(DomainMappingSslSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DomainMappingSslSettings or *DomainMappingSslSettings", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DomainMappingSslSettings)
	if !ok {
		actualNotPointer, ok := a.(DomainMappingSslSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DomainMappingSslSettings", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CertificateId, actual.CertificateId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertificateId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SslManagementType, actual.SslManagementType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SslManagementType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PendingManagedCertificateId, actual.PendingManagedCertificateId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PendingManagedCertificateId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDomainMappingResourceRecordsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DomainMappingResourceRecords)
	if !ok {
		desiredNotPointer, ok := d.(DomainMappingResourceRecords)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DomainMappingResourceRecords or *DomainMappingResourceRecords", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DomainMappingResourceRecords)
	if !ok {
		actualNotPointer, ok := a.(DomainMappingResourceRecords)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DomainMappingResourceRecords", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rrdata, actual.Rrdata, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Rrdata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *DomainMapping) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.App), dcl.ValueOrEmptyString(n.Name)
}

func (r *DomainMapping) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.App)
}

func (r *DomainMapping) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.App), dcl.ValueOrEmptyString(n.Name)
}

func (r *DomainMapping) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "UpdateDomainMapping" {
		fields := map[string]interface{}{
			"app":  dcl.ValueOrEmptyString(n.App),
			"name": dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("apps/{{app}}/domainMappings/{{name}}", "https://appengine.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the DomainMapping resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *DomainMapping) marshal(c *Client) ([]byte, error) {
	m, err := expandDomainMapping(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling DomainMapping: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalDomainMapping decodes JSON responses into the DomainMapping resource schema.
func unmarshalDomainMapping(b []byte, c *Client) (*DomainMapping, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapDomainMapping(m, c)
}

func unmarshalMapDomainMapping(m map[string]interface{}, c *Client) (*DomainMapping, error) {

	return flattenDomainMapping(c, m), nil
}

// expandDomainMapping expands DomainMapping into a JSON request object.
func expandDomainMapping(c *Client, f *DomainMapping) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v, err := expandDomainMappingSslSettings(c, f.SslSettings); err != nil {
		return nil, fmt.Errorf("error expanding SslSettings into sslSettings: %w", err)
	} else if v != nil {
		m["sslSettings"] = v
	}
	if v, err := expandDomainMappingResourceRecordsSlice(c, f.ResourceRecords); err != nil {
		return nil, fmt.Errorf("error expanding ResourceRecords into resourceRecords: %w", err)
	} else {
		m["resourceRecords"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding App into app: %w", err)
	} else if v != nil {
		m["app"] = v
	}

	return m, nil
}

// flattenDomainMapping flattens DomainMapping from a JSON request object into the
// DomainMapping type.
func flattenDomainMapping(c *Client, i interface{}) *DomainMapping {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &DomainMapping{}
	res.SelfLink = dcl.FlattenString(m["name"])
	res.Name = dcl.FlattenString(m["id"])
	res.SslSettings = flattenDomainMappingSslSettings(c, m["sslSettings"])
	res.ResourceRecords = flattenDomainMappingResourceRecordsSlice(c, m["resourceRecords"])
	res.App = dcl.FlattenString(m["app"])

	return res
}

// expandDomainMappingSslSettingsMap expands the contents of DomainMappingSslSettings into a JSON
// request object.
func expandDomainMappingSslSettingsMap(c *Client, f map[string]DomainMappingSslSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDomainMappingSslSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDomainMappingSslSettingsSlice expands the contents of DomainMappingSslSettings into a JSON
// request object.
func expandDomainMappingSslSettingsSlice(c *Client, f []DomainMappingSslSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDomainMappingSslSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDomainMappingSslSettingsMap flattens the contents of DomainMappingSslSettings from a JSON
// response object.
func flattenDomainMappingSslSettingsMap(c *Client, i interface{}) map[string]DomainMappingSslSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DomainMappingSslSettings{}
	}

	if len(a) == 0 {
		return map[string]DomainMappingSslSettings{}
	}

	items := make(map[string]DomainMappingSslSettings)
	for k, item := range a {
		items[k] = *flattenDomainMappingSslSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDomainMappingSslSettingsSlice flattens the contents of DomainMappingSslSettings from a JSON
// response object.
func flattenDomainMappingSslSettingsSlice(c *Client, i interface{}) []DomainMappingSslSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []DomainMappingSslSettings{}
	}

	if len(a) == 0 {
		return []DomainMappingSslSettings{}
	}

	items := make([]DomainMappingSslSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDomainMappingSslSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandDomainMappingSslSettings expands an instance of DomainMappingSslSettings into a JSON
// request object.
func expandDomainMappingSslSettings(c *Client, f *DomainMappingSslSettings) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CertificateId; !dcl.IsEmptyValueIndirect(v) {
		m["certificateId"] = v
	}
	if v := f.SslManagementType; !dcl.IsEmptyValueIndirect(v) {
		m["sslManagementType"] = v
	}
	if v := f.PendingManagedCertificateId; !dcl.IsEmptyValueIndirect(v) {
		m["pendingManagedCertificateId"] = v
	}

	return m, nil
}

// flattenDomainMappingSslSettings flattens an instance of DomainMappingSslSettings from a JSON
// response object.
func flattenDomainMappingSslSettings(c *Client, i interface{}) *DomainMappingSslSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DomainMappingSslSettings{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDomainMappingSslSettings
	}
	r.CertificateId = dcl.FlattenString(m["certificateId"])
	r.SslManagementType = flattenDomainMappingSslSettingsSslManagementTypeEnum(m["sslManagementType"])
	r.PendingManagedCertificateId = dcl.FlattenString(m["pendingManagedCertificateId"])

	return r
}

// expandDomainMappingResourceRecordsMap expands the contents of DomainMappingResourceRecords into a JSON
// request object.
func expandDomainMappingResourceRecordsMap(c *Client, f map[string]DomainMappingResourceRecords) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDomainMappingResourceRecords(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDomainMappingResourceRecordsSlice expands the contents of DomainMappingResourceRecords into a JSON
// request object.
func expandDomainMappingResourceRecordsSlice(c *Client, f []DomainMappingResourceRecords) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDomainMappingResourceRecords(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDomainMappingResourceRecordsMap flattens the contents of DomainMappingResourceRecords from a JSON
// response object.
func flattenDomainMappingResourceRecordsMap(c *Client, i interface{}) map[string]DomainMappingResourceRecords {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DomainMappingResourceRecords{}
	}

	if len(a) == 0 {
		return map[string]DomainMappingResourceRecords{}
	}

	items := make(map[string]DomainMappingResourceRecords)
	for k, item := range a {
		items[k] = *flattenDomainMappingResourceRecords(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDomainMappingResourceRecordsSlice flattens the contents of DomainMappingResourceRecords from a JSON
// response object.
func flattenDomainMappingResourceRecordsSlice(c *Client, i interface{}) []DomainMappingResourceRecords {
	a, ok := i.([]interface{})
	if !ok {
		return []DomainMappingResourceRecords{}
	}

	if len(a) == 0 {
		return []DomainMappingResourceRecords{}
	}

	items := make([]DomainMappingResourceRecords, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDomainMappingResourceRecords(c, item.(map[string]interface{})))
	}

	return items
}

// expandDomainMappingResourceRecords expands an instance of DomainMappingResourceRecords into a JSON
// request object.
func expandDomainMappingResourceRecords(c *Client, f *DomainMappingResourceRecords) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Rrdata; !dcl.IsEmptyValueIndirect(v) {
		m["rrdata"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenDomainMappingResourceRecords flattens an instance of DomainMappingResourceRecords from a JSON
// response object.
func flattenDomainMappingResourceRecords(c *Client, i interface{}) *DomainMappingResourceRecords {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DomainMappingResourceRecords{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDomainMappingResourceRecords
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Rrdata = dcl.FlattenString(m["rrdata"])
	r.Type = flattenDomainMappingResourceRecordsTypeEnum(m["type"])

	return r
}

// flattenDomainMappingSslSettingsSslManagementTypeEnumMap flattens the contents of DomainMappingSslSettingsSslManagementTypeEnum from a JSON
// response object.
func flattenDomainMappingSslSettingsSslManagementTypeEnumMap(c *Client, i interface{}) map[string]DomainMappingSslSettingsSslManagementTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DomainMappingSslSettingsSslManagementTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]DomainMappingSslSettingsSslManagementTypeEnum{}
	}

	items := make(map[string]DomainMappingSslSettingsSslManagementTypeEnum)
	for k, item := range a {
		items[k] = *flattenDomainMappingSslSettingsSslManagementTypeEnum(item.(interface{}))
	}

	return items
}

// flattenDomainMappingSslSettingsSslManagementTypeEnumSlice flattens the contents of DomainMappingSslSettingsSslManagementTypeEnum from a JSON
// response object.
func flattenDomainMappingSslSettingsSslManagementTypeEnumSlice(c *Client, i interface{}) []DomainMappingSslSettingsSslManagementTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DomainMappingSslSettingsSslManagementTypeEnum{}
	}

	if len(a) == 0 {
		return []DomainMappingSslSettingsSslManagementTypeEnum{}
	}

	items := make([]DomainMappingSslSettingsSslManagementTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDomainMappingSslSettingsSslManagementTypeEnum(item.(interface{})))
	}

	return items
}

// flattenDomainMappingSslSettingsSslManagementTypeEnum asserts that an interface is a string, and returns a
// pointer to a *DomainMappingSslSettingsSslManagementTypeEnum with the same value as that string.
func flattenDomainMappingSslSettingsSslManagementTypeEnum(i interface{}) *DomainMappingSslSettingsSslManagementTypeEnum {
	s, ok := i.(string)
	if !ok {
		return DomainMappingSslSettingsSslManagementTypeEnumRef("")
	}

	return DomainMappingSslSettingsSslManagementTypeEnumRef(s)
}

// flattenDomainMappingResourceRecordsTypeEnumMap flattens the contents of DomainMappingResourceRecordsTypeEnum from a JSON
// response object.
func flattenDomainMappingResourceRecordsTypeEnumMap(c *Client, i interface{}) map[string]DomainMappingResourceRecordsTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DomainMappingResourceRecordsTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]DomainMappingResourceRecordsTypeEnum{}
	}

	items := make(map[string]DomainMappingResourceRecordsTypeEnum)
	for k, item := range a {
		items[k] = *flattenDomainMappingResourceRecordsTypeEnum(item.(interface{}))
	}

	return items
}

// flattenDomainMappingResourceRecordsTypeEnumSlice flattens the contents of DomainMappingResourceRecordsTypeEnum from a JSON
// response object.
func flattenDomainMappingResourceRecordsTypeEnumSlice(c *Client, i interface{}) []DomainMappingResourceRecordsTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DomainMappingResourceRecordsTypeEnum{}
	}

	if len(a) == 0 {
		return []DomainMappingResourceRecordsTypeEnum{}
	}

	items := make([]DomainMappingResourceRecordsTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDomainMappingResourceRecordsTypeEnum(item.(interface{})))
	}

	return items
}

// flattenDomainMappingResourceRecordsTypeEnum asserts that an interface is a string, and returns a
// pointer to a *DomainMappingResourceRecordsTypeEnum with the same value as that string.
func flattenDomainMappingResourceRecordsTypeEnum(i interface{}) *DomainMappingResourceRecordsTypeEnum {
	s, ok := i.(string)
	if !ok {
		return DomainMappingResourceRecordsTypeEnumRef("")
	}

	return DomainMappingResourceRecordsTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *DomainMapping) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalDomainMapping(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.App == nil && ncr.App == nil {
			c.Config.Logger.Info("Both App fields null - considering equal.")
		} else if nr.App == nil || ncr.App == nil {
			c.Config.Logger.Info("Only one App field is null - considering unequal.")
			return false
		} else if *nr.App != *ncr.App {
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

type domainMappingDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         domainMappingApiOperation
}

func convertFieldDiffsToDomainMappingDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]domainMappingDiff, error) {
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
	var diffs []domainMappingDiff
	// For each operation name, create a domainMappingDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := domainMappingDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToDomainMappingApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToDomainMappingApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (domainMappingApiOperation, error) {
	switch opName {

	case "updateDomainMappingUpdateDomainMappingOperation":
		return &updateDomainMappingUpdateDomainMappingOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
