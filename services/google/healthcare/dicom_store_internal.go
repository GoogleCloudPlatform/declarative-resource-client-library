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
package healthcare

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *DicomStore) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Dataset, "Dataset"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.NotificationConfig) {
		if err := r.NotificationConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DicomStoreNotificationConfig) validate() error {
	return nil
}
func (r *DicomStore) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://healthcare.googleapis.com/v1/", params)
}

func (r *DicomStore) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/dicomStores/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *DicomStore) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/dicomStores", nr.basePath(), userBasePath, params), nil

}

func (r *DicomStore) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/dicomStores?dicomStoreId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *DicomStore) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/dicomStores/{{name}}", nr.basePath(), userBasePath, params), nil
}

// dicomStoreApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type dicomStoreApiOperation interface {
	do(context.Context, *DicomStore, *Client) error
}

// newUpdateDicomStoreUpdateDicomStoreRequest creates a request for an
// DicomStore resource's UpdateDicomStore update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDicomStoreUpdateDicomStoreRequest(ctx context.Context, f *DicomStore, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := expandDicomStoreNotificationConfig(c, f.NotificationConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding NotificationConfig into notificationConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["notificationConfig"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	return req, nil
}

// marshalUpdateDicomStoreUpdateDicomStoreRequest converts the update into
// the final JSON request body.
func marshalUpdateDicomStoreUpdateDicomStoreRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateDicomStoreUpdateDicomStoreOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDicomStoreUpdateDicomStoreOperation) do(ctx context.Context, r *DicomStore, c *Client) error {
	_, err := c.GetDicomStore(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateDicomStore")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateDicomStoreUpdateDicomStoreRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateDicomStoreUpdateDicomStoreRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listDicomStoreRaw(ctx context.Context, r *DicomStore, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != DicomStoreMaxPage {
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

type listDicomStoreOperation struct {
	DicomStores []map[string]interface{} `json:"dicomStores"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listDicomStore(ctx context.Context, r *DicomStore, pageToken string, pageSize int32) ([]*DicomStore, string, error) {
	b, err := c.listDicomStoreRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listDicomStoreOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*DicomStore
	for _, v := range m.DicomStores {
		res, err := unmarshalMapDicomStore(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		res.Dataset = r.Dataset
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllDicomStore(ctx context.Context, f func(*DicomStore) bool, resources []*DicomStore) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteDicomStore(ctx, res)
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

type deleteDicomStoreOperation struct{}

func (op *deleteDicomStoreOperation) do(ctx context.Context, r *DicomStore, c *Client) error {
	r, err := c.GetDicomStore(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.InfoWithContextf(ctx, "DicomStore not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetDicomStore checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete DicomStore: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createDicomStoreOperation struct {
	response map[string]interface{}
}

func (op *createDicomStoreOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createDicomStoreOperation) do(ctx context.Context, r *DicomStore, c *Client) error {
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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetDicomStore(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDicomStoreRaw(ctx context.Context, r *DicomStore) ([]byte, error) {

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

func (c *Client) dicomStoreDiffsForRawDesired(ctx context.Context, rawDesired *DicomStore, opts ...dcl.ApplyOption) (initial, desired *DicomStore, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *DicomStore
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*DicomStore); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected DicomStore, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetDicomStore(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a DicomStore resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve DicomStore resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that DicomStore resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeDicomStoreDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for DicomStore: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for DicomStore: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractDicomStoreFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeDicomStoreInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for DicomStore: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeDicomStoreDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for DicomStore: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffDicomStore(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeDicomStoreInitialState(rawInitial, rawDesired *DicomStore) (*DicomStore, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeDicomStoreDesiredState(rawDesired, rawInitial *DicomStore, opts ...dcl.ApplyOption) (*DicomStore, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.NotificationConfig = canonicalizeDicomStoreNotificationConfig(rawDesired.NotificationConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &DicomStore{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.NotificationConfig = canonicalizeDicomStoreNotificationConfig(rawDesired.NotificationConfig, rawInitial.NotificationConfig, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
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
	if dcl.NameToSelfLink(rawDesired.Dataset, rawInitial.Dataset) {
		canonicalDesired.Dataset = rawInitial.Dataset
	} else {
		canonicalDesired.Dataset = rawDesired.Dataset
	}

	return canonicalDesired, nil
}

func canonicalizeDicomStoreNewState(c *Client, rawNew, rawDesired *DicomStore) (*DicomStore, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.NotificationConfig) && dcl.IsEmptyValueIndirect(rawDesired.NotificationConfig) {
		rawNew.NotificationConfig = rawDesired.NotificationConfig
	} else {
		rawNew.NotificationConfig = canonicalizeNewDicomStoreNotificationConfig(c, rawDesired.NotificationConfig, rawNew.NotificationConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.Dataset = rawDesired.Dataset

	return rawNew, nil
}

func canonicalizeDicomStoreNotificationConfig(des, initial *DicomStoreNotificationConfig, opts ...dcl.ApplyOption) *DicomStoreNotificationConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DicomStoreNotificationConfig{}

	if dcl.IsZeroValue(des.PubsubTopic) || (dcl.IsEmptyValueIndirect(des.PubsubTopic) && dcl.IsEmptyValueIndirect(initial.PubsubTopic)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.PubsubTopic = initial.PubsubTopic
	} else {
		cDes.PubsubTopic = des.PubsubTopic
	}

	return cDes
}

func canonicalizeDicomStoreNotificationConfigSlice(des, initial []DicomStoreNotificationConfig, opts ...dcl.ApplyOption) []DicomStoreNotificationConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DicomStoreNotificationConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDicomStoreNotificationConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DicomStoreNotificationConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDicomStoreNotificationConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDicomStoreNotificationConfig(c *Client, des, nw *DicomStoreNotificationConfig) *DicomStoreNotificationConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DicomStoreNotificationConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDicomStoreNotificationConfigSet(c *Client, des, nw []DicomStoreNotificationConfig) []DicomStoreNotificationConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []DicomStoreNotificationConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDicomStoreNotificationConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDicomStoreNotificationConfigSlice(c *Client, des, nw []DicomStoreNotificationConfig) []DicomStoreNotificationConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DicomStoreNotificationConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDicomStoreNotificationConfig(c, &d, &n))
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
func diffDicomStore(c *Client, desired, actual *DicomStore, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.NotificationConfig, actual.NotificationConfig, dcl.DiffInfo{ObjectFunction: compareDicomStoreNotificationConfigNewStyle, EmptyObject: EmptyDicomStoreNotificationConfig, OperationSelector: dcl.TriggersOperation("updateDicomStoreUpdateDicomStoreOperation")}, fn.AddNest("NotificationConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDicomStoreUpdateDicomStoreOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Dataset, actual.Dataset, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Dataset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareDicomStoreNotificationConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DicomStoreNotificationConfig)
	if !ok {
		desiredNotPointer, ok := d.(DicomStoreNotificationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DicomStoreNotificationConfig or *DicomStoreNotificationConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DicomStoreNotificationConfig)
	if !ok {
		actualNotPointer, ok := a.(DicomStoreNotificationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DicomStoreNotificationConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PubsubTopic, actual.PubsubTopic, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateDicomStoreUpdateDicomStoreOperation")}, fn.AddNest("PubsubTopic")); len(ds) != 0 || err != nil {
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
func (r *DicomStore) urlNormalized() *DicomStore {
	normalized := dcl.Copy(*r).(DicomStore)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Dataset = dcl.SelfLinkToName(r.Dataset)
	return &normalized
}

func (r *DicomStore) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateDicomStore" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/dicomStores/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the DicomStore resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *DicomStore) marshal(c *Client) ([]byte, error) {
	m, err := expandDicomStore(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling DicomStore: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalDicomStore decodes JSON responses into the DicomStore resource schema.
func unmarshalDicomStore(b []byte, c *Client, res *DicomStore) (*DicomStore, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapDicomStore(m, c, res)
}

func unmarshalMapDicomStore(m map[string]interface{}, c *Client, res *DicomStore) (*DicomStore, error) {

	flattened := flattenDicomStore(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandDicomStore expands DicomStore into a JSON request object.
func expandDicomStore(c *Client, f *DicomStore) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandDicomStoreNotificationConfig(c, f.NotificationConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding NotificationConfig into notificationConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["notificationConfig"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Dataset into dataset: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dataset"] = v
	}

	return m, nil
}

// flattenDicomStore flattens DicomStore from a JSON request object into the
// DicomStore type.
func flattenDicomStore(c *Client, i interface{}, res *DicomStore) *DicomStore {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &DicomStore{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.NotificationConfig = flattenDicomStoreNotificationConfig(c, m["notificationConfig"], res)
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Dataset = dcl.FlattenString(m["dataset"])

	return resultRes
}

// expandDicomStoreNotificationConfigMap expands the contents of DicomStoreNotificationConfig into a JSON
// request object.
func expandDicomStoreNotificationConfigMap(c *Client, f map[string]DicomStoreNotificationConfig, res *DicomStore) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDicomStoreNotificationConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDicomStoreNotificationConfigSlice expands the contents of DicomStoreNotificationConfig into a JSON
// request object.
func expandDicomStoreNotificationConfigSlice(c *Client, f []DicomStoreNotificationConfig, res *DicomStore) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDicomStoreNotificationConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDicomStoreNotificationConfigMap flattens the contents of DicomStoreNotificationConfig from a JSON
// response object.
func flattenDicomStoreNotificationConfigMap(c *Client, i interface{}, res *DicomStore) map[string]DicomStoreNotificationConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DicomStoreNotificationConfig{}
	}

	if len(a) == 0 {
		return map[string]DicomStoreNotificationConfig{}
	}

	items := make(map[string]DicomStoreNotificationConfig)
	for k, item := range a {
		items[k] = *flattenDicomStoreNotificationConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDicomStoreNotificationConfigSlice flattens the contents of DicomStoreNotificationConfig from a JSON
// response object.
func flattenDicomStoreNotificationConfigSlice(c *Client, i interface{}, res *DicomStore) []DicomStoreNotificationConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []DicomStoreNotificationConfig{}
	}

	if len(a) == 0 {
		return []DicomStoreNotificationConfig{}
	}

	items := make([]DicomStoreNotificationConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDicomStoreNotificationConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDicomStoreNotificationConfig expands an instance of DicomStoreNotificationConfig into a JSON
// request object.
func expandDicomStoreNotificationConfig(c *Client, f *DicomStoreNotificationConfig, res *DicomStore) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PubsubTopic; !dcl.IsEmptyValueIndirect(v) {
		m["pubsubTopic"] = v
	}

	return m, nil
}

// flattenDicomStoreNotificationConfig flattens an instance of DicomStoreNotificationConfig from a JSON
// response object.
func flattenDicomStoreNotificationConfig(c *Client, i interface{}, res *DicomStore) *DicomStoreNotificationConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DicomStoreNotificationConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDicomStoreNotificationConfig
	}
	r.PubsubTopic = dcl.FlattenString(m["pubsubTopic"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *DicomStore) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalDicomStore(b, c, r)
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
		if nr.Dataset == nil && ncr.Dataset == nil {
			c.Config.Logger.Info("Both Dataset fields null - considering equal.")
		} else if nr.Dataset == nil || ncr.Dataset == nil {
			c.Config.Logger.Info("Only one Dataset field is null - considering unequal.")
			return false
		} else if *nr.Dataset != *ncr.Dataset {
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

type dicomStoreDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         dicomStoreApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToDicomStoreDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]dicomStoreDiff, error) {
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
	var diffs []dicomStoreDiff
	// For each operation name, create a dicomStoreDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := dicomStoreDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToDicomStoreApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToDicomStoreApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (dicomStoreApiOperation, error) {
	switch opName {

	case "updateDicomStoreUpdateDicomStoreOperation":
		return &updateDicomStoreUpdateDicomStoreOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractDicomStoreFields(r *DicomStore) error {
	vNotificationConfig := r.NotificationConfig
	if vNotificationConfig == nil {
		// note: explicitly not the empty object.
		vNotificationConfig = &DicomStoreNotificationConfig{}
	}
	if err := extractDicomStoreNotificationConfigFields(r, vNotificationConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vNotificationConfig) {
		r.NotificationConfig = vNotificationConfig
	}
	return nil
}
func extractDicomStoreNotificationConfigFields(r *DicomStore, o *DicomStoreNotificationConfig) error {
	return nil
}

func postReadExtractDicomStoreFields(r *DicomStore) error {
	vNotificationConfig := r.NotificationConfig
	if vNotificationConfig == nil {
		// note: explicitly not the empty object.
		vNotificationConfig = &DicomStoreNotificationConfig{}
	}
	if err := postReadExtractDicomStoreNotificationConfigFields(r, vNotificationConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vNotificationConfig) {
		r.NotificationConfig = vNotificationConfig
	}
	return nil
}
func postReadExtractDicomStoreNotificationConfigFields(r *DicomStore, o *DicomStoreNotificationConfig) error {
	return nil
}
