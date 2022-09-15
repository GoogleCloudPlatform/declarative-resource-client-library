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
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *FhirStore) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "version"); err != nil {
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
	if !dcl.IsEmptyValueIndirect(r.ValidationConfig) {
		if err := r.ValidationConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FhirStoreNotificationConfig) validate() error {
	return nil
}
func (r *FhirStoreStreamConfigs) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BigqueryDestination) {
		if err := r.BigqueryDestination.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FhirStoreStreamConfigsBigqueryDestination) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SchemaConfig) {
		if err := r.SchemaConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) validate() error {
	return nil
}
func (r *FhirStoreValidationConfig) validate() error {
	return nil
}
func (r *FhirStore) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://healthcare.googleapis.com/v1/", params)
}

func (r *FhirStore) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/fhirStores/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *FhirStore) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/fhirStores", nr.basePath(), userBasePath, params), nil

}

func (r *FhirStore) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/fhirStores?fhirStoreId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *FhirStore) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/fhirStores/{{name}}", nr.basePath(), userBasePath, params), nil
}

// fhirStoreApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type fhirStoreApiOperation interface {
	do(context.Context, *FhirStore, *Client) error
}

// newUpdateFhirStoreUpdateFhirStoreRequest creates a request for an
// FhirStore resource's UpdateFhirStore update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFhirStoreUpdateFhirStoreRequest(ctx context.Context, f *FhirStore, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.EnableUpdateCreate; !dcl.IsEmptyValueIndirect(v) {
		req["enableUpdateCreate"] = v
	}
	if v, err := expandFhirStoreNotificationConfig(c, f.NotificationConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding NotificationConfig into notificationConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["notificationConfig"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandFhirStoreStreamConfigsSlice(c, f.StreamConfigs, res); err != nil {
		return nil, fmt.Errorf("error expanding StreamConfigs into streamConfigs: %w", err)
	} else if v != nil {
		req["streamConfigs"] = v
	}
	if v, err := expandFhirStoreValidationConfig(c, f.ValidationConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding ValidationConfig into validationConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["validationConfig"] = v
	}
	if v := f.DefaultSearchHandlingStrict; !dcl.IsEmptyValueIndirect(v) {
		req["defaultSearchHandlingStrict"] = v
	}
	if v := f.ComplexDataTypeReferenceParsing; !dcl.IsEmptyValueIndirect(v) {
		req["complexDataTypeReferenceParsing"] = v
	}
	return req, nil
}

// marshalUpdateFhirStoreUpdateFhirStoreRequest converts the update into
// the final JSON request body.
func marshalUpdateFhirStoreUpdateFhirStoreRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFhirStoreUpdateFhirStoreOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateFhirStoreUpdateFhirStoreOperation) do(ctx context.Context, r *FhirStore, c *Client) error {
	_, err := c.GetFhirStore(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateFhirStore")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateFhirStoreUpdateFhirStoreRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateFhirStoreUpdateFhirStoreRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listFhirStoreRaw(ctx context.Context, r *FhirStore, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FhirStoreMaxPage {
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

type listFhirStoreOperation struct {
	FhirStores []map[string]interface{} `json:"fhirStores"`
	Token      string                   `json:"nextPageToken"`
}

func (c *Client) listFhirStore(ctx context.Context, r *FhirStore, pageToken string, pageSize int32) ([]*FhirStore, string, error) {
	b, err := c.listFhirStoreRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFhirStoreOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*FhirStore
	for _, v := range m.FhirStores {
		res, err := unmarshalMapFhirStore(v, c, r)
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

func (c *Client) deleteAllFhirStore(ctx context.Context, f func(*FhirStore) bool, resources []*FhirStore) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFhirStore(ctx, res)
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

type deleteFhirStoreOperation struct{}

func (op *deleteFhirStoreOperation) do(ctx context.Context, r *FhirStore, c *Client) error {
	r, err := c.GetFhirStore(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.InfoWithContextf(ctx, "FhirStore not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetFhirStore checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete FhirStore: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createFhirStoreOperation struct {
	response map[string]interface{}
}

func (op *createFhirStoreOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFhirStoreOperation) do(ctx context.Context, r *FhirStore, c *Client) error {
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

	if _, err := c.GetFhirStore(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFhirStoreRaw(ctx context.Context, r *FhirStore) ([]byte, error) {

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

func (c *Client) fhirStoreDiffsForRawDesired(ctx context.Context, rawDesired *FhirStore, opts ...dcl.ApplyOption) (initial, desired *FhirStore, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *FhirStore
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*FhirStore); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected FhirStore, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFhirStore(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a FhirStore resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve FhirStore resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that FhirStore resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFhirStoreDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for FhirStore: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for FhirStore: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractFhirStoreFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFhirStoreInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for FhirStore: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFhirStoreDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for FhirStore: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFhirStore(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFhirStoreInitialState(rawInitial, rawDesired *FhirStore) (*FhirStore, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFhirStoreDesiredState(rawDesired, rawInitial *FhirStore, opts ...dcl.ApplyOption) (*FhirStore, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.NotificationConfig = canonicalizeFhirStoreNotificationConfig(rawDesired.NotificationConfig, nil, opts...)
		rawDesired.ValidationConfig = canonicalizeFhirStoreValidationConfig(rawDesired.ValidationConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &FhirStore{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.BoolCanonicalize(rawDesired.EnableUpdateCreate, rawInitial.EnableUpdateCreate) {
		canonicalDesired.EnableUpdateCreate = rawInitial.EnableUpdateCreate
	} else {
		canonicalDesired.EnableUpdateCreate = rawDesired.EnableUpdateCreate
	}
	canonicalDesired.NotificationConfig = canonicalizeFhirStoreNotificationConfig(rawDesired.NotificationConfig, rawInitial.NotificationConfig, opts...)
	if dcl.BoolCanonicalize(rawDesired.DisableReferentialIntegrity, rawInitial.DisableReferentialIntegrity) {
		canonicalDesired.DisableReferentialIntegrity = rawInitial.DisableReferentialIntegrity
	} else {
		canonicalDesired.DisableReferentialIntegrity = rawDesired.DisableReferentialIntegrity
	}
	if dcl.IsZeroValue(rawDesired.ShardNum) || (dcl.IsEmptyValueIndirect(rawDesired.ShardNum) && dcl.IsEmptyValueIndirect(rawInitial.ShardNum)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ShardNum = rawInitial.ShardNum
	} else {
		canonicalDesired.ShardNum = rawDesired.ShardNum
	}
	if dcl.BoolCanonicalize(rawDesired.DisableResourceVersioning, rawInitial.DisableResourceVersioning) {
		canonicalDesired.DisableResourceVersioning = rawInitial.DisableResourceVersioning
	} else {
		canonicalDesired.DisableResourceVersioning = rawDesired.DisableResourceVersioning
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.IsZeroValue(rawDesired.Version) || (dcl.IsEmptyValueIndirect(rawDesired.Version) && dcl.IsEmptyValueIndirect(rawInitial.Version)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Version = rawInitial.Version
	} else {
		canonicalDesired.Version = rawDesired.Version
	}
	canonicalDesired.StreamConfigs = canonicalizeFhirStoreStreamConfigsSlice(rawDesired.StreamConfigs, rawInitial.StreamConfigs, opts...)
	canonicalDesired.ValidationConfig = canonicalizeFhirStoreValidationConfig(rawDesired.ValidationConfig, rawInitial.ValidationConfig, opts...)
	if dcl.BoolCanonicalize(rawDesired.DefaultSearchHandlingStrict, rawInitial.DefaultSearchHandlingStrict) {
		canonicalDesired.DefaultSearchHandlingStrict = rawInitial.DefaultSearchHandlingStrict
	} else {
		canonicalDesired.DefaultSearchHandlingStrict = rawDesired.DefaultSearchHandlingStrict
	}
	if dcl.IsZeroValue(rawDesired.ComplexDataTypeReferenceParsing) || (dcl.IsEmptyValueIndirect(rawDesired.ComplexDataTypeReferenceParsing) && dcl.IsEmptyValueIndirect(rawInitial.ComplexDataTypeReferenceParsing)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ComplexDataTypeReferenceParsing = rawInitial.ComplexDataTypeReferenceParsing
	} else {
		canonicalDesired.ComplexDataTypeReferenceParsing = rawDesired.ComplexDataTypeReferenceParsing
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

func canonicalizeFhirStoreNewState(c *Client, rawNew, rawDesired *FhirStore) (*FhirStore, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.EnableUpdateCreate) && dcl.IsEmptyValueIndirect(rawDesired.EnableUpdateCreate) {
		rawNew.EnableUpdateCreate = rawDesired.EnableUpdateCreate
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableUpdateCreate, rawNew.EnableUpdateCreate) {
			rawNew.EnableUpdateCreate = rawDesired.EnableUpdateCreate
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NotificationConfig) && dcl.IsEmptyValueIndirect(rawDesired.NotificationConfig) {
		rawNew.NotificationConfig = rawDesired.NotificationConfig
	} else {
		rawNew.NotificationConfig = canonicalizeNewFhirStoreNotificationConfig(c, rawDesired.NotificationConfig, rawNew.NotificationConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisableReferentialIntegrity) && dcl.IsEmptyValueIndirect(rawDesired.DisableReferentialIntegrity) {
		rawNew.DisableReferentialIntegrity = rawDesired.DisableReferentialIntegrity
	} else {
		if dcl.BoolCanonicalize(rawDesired.DisableReferentialIntegrity, rawNew.DisableReferentialIntegrity) {
			rawNew.DisableReferentialIntegrity = rawDesired.DisableReferentialIntegrity
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ShardNum) && dcl.IsEmptyValueIndirect(rawDesired.ShardNum) {
		rawNew.ShardNum = rawDesired.ShardNum
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisableResourceVersioning) && dcl.IsEmptyValueIndirect(rawDesired.DisableResourceVersioning) {
		rawNew.DisableResourceVersioning = rawDesired.DisableResourceVersioning
	} else {
		if dcl.BoolCanonicalize(rawDesired.DisableResourceVersioning, rawNew.DisableResourceVersioning) {
			rawNew.DisableResourceVersioning = rawDesired.DisableResourceVersioning
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Version) && dcl.IsEmptyValueIndirect(rawDesired.Version) {
		rawNew.Version = rawDesired.Version
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StreamConfigs) && dcl.IsEmptyValueIndirect(rawDesired.StreamConfigs) {
		rawNew.StreamConfigs = rawDesired.StreamConfigs
	} else {
		rawNew.StreamConfigs = canonicalizeNewFhirStoreStreamConfigsSlice(c, rawDesired.StreamConfigs, rawNew.StreamConfigs)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ValidationConfig) && dcl.IsEmptyValueIndirect(rawDesired.ValidationConfig) {
		rawNew.ValidationConfig = rawDesired.ValidationConfig
	} else {
		rawNew.ValidationConfig = canonicalizeNewFhirStoreValidationConfig(c, rawDesired.ValidationConfig, rawNew.ValidationConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultSearchHandlingStrict) && dcl.IsEmptyValueIndirect(rawDesired.DefaultSearchHandlingStrict) {
		rawNew.DefaultSearchHandlingStrict = rawDesired.DefaultSearchHandlingStrict
	} else {
		if dcl.BoolCanonicalize(rawDesired.DefaultSearchHandlingStrict, rawNew.DefaultSearchHandlingStrict) {
			rawNew.DefaultSearchHandlingStrict = rawDesired.DefaultSearchHandlingStrict
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ComplexDataTypeReferenceParsing) && dcl.IsEmptyValueIndirect(rawDesired.ComplexDataTypeReferenceParsing) {
		rawNew.ComplexDataTypeReferenceParsing = rawDesired.ComplexDataTypeReferenceParsing
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.Dataset = rawDesired.Dataset

	return rawNew, nil
}

func canonicalizeFhirStoreNotificationConfig(des, initial *FhirStoreNotificationConfig, opts ...dcl.ApplyOption) *FhirStoreNotificationConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FhirStoreNotificationConfig{}

	if dcl.IsZeroValue(des.PubsubTopic) || (dcl.IsEmptyValueIndirect(des.PubsubTopic) && dcl.IsEmptyValueIndirect(initial.PubsubTopic)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.PubsubTopic = initial.PubsubTopic
	} else {
		cDes.PubsubTopic = des.PubsubTopic
	}

	return cDes
}

func canonicalizeFhirStoreNotificationConfigSlice(des, initial []FhirStoreNotificationConfig, opts ...dcl.ApplyOption) []FhirStoreNotificationConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FhirStoreNotificationConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFhirStoreNotificationConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FhirStoreNotificationConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFhirStoreNotificationConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFhirStoreNotificationConfig(c *Client, des, nw *FhirStoreNotificationConfig) *FhirStoreNotificationConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FhirStoreNotificationConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFhirStoreNotificationConfigSet(c *Client, des, nw []FhirStoreNotificationConfig) []FhirStoreNotificationConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []FhirStoreNotificationConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFhirStoreNotificationConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFhirStoreNotificationConfigSlice(c *Client, des, nw []FhirStoreNotificationConfig) []FhirStoreNotificationConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FhirStoreNotificationConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFhirStoreNotificationConfig(c, &d, &n))
	}

	return items
}

func canonicalizeFhirStoreStreamConfigs(des, initial *FhirStoreStreamConfigs, opts ...dcl.ApplyOption) *FhirStoreStreamConfigs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FhirStoreStreamConfigs{}

	if dcl.StringArrayCanonicalize(des.ResourceTypes, initial.ResourceTypes) {
		cDes.ResourceTypes = initial.ResourceTypes
	} else {
		cDes.ResourceTypes = des.ResourceTypes
	}
	cDes.BigqueryDestination = canonicalizeFhirStoreStreamConfigsBigqueryDestination(des.BigqueryDestination, initial.BigqueryDestination, opts...)

	return cDes
}

func canonicalizeFhirStoreStreamConfigsSlice(des, initial []FhirStoreStreamConfigs, opts ...dcl.ApplyOption) []FhirStoreStreamConfigs {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FhirStoreStreamConfigs, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFhirStoreStreamConfigs(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FhirStoreStreamConfigs, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFhirStoreStreamConfigs(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFhirStoreStreamConfigs(c *Client, des, nw *FhirStoreStreamConfigs) *FhirStoreStreamConfigs {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FhirStoreStreamConfigs while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.ResourceTypes, nw.ResourceTypes) {
		nw.ResourceTypes = des.ResourceTypes
	}
	nw.BigqueryDestination = canonicalizeNewFhirStoreStreamConfigsBigqueryDestination(c, des.BigqueryDestination, nw.BigqueryDestination)

	return nw
}

func canonicalizeNewFhirStoreStreamConfigsSet(c *Client, des, nw []FhirStoreStreamConfigs) []FhirStoreStreamConfigs {
	if des == nil {
		return nw
	}
	var reorderedNew []FhirStoreStreamConfigs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFhirStoreStreamConfigsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFhirStoreStreamConfigsSlice(c *Client, des, nw []FhirStoreStreamConfigs) []FhirStoreStreamConfigs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FhirStoreStreamConfigs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFhirStoreStreamConfigs(c, &d, &n))
	}

	return items
}

func canonicalizeFhirStoreStreamConfigsBigqueryDestination(des, initial *FhirStoreStreamConfigsBigqueryDestination, opts ...dcl.ApplyOption) *FhirStoreStreamConfigsBigqueryDestination {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FhirStoreStreamConfigsBigqueryDestination{}

	if dcl.StringCanonicalize(des.DatasetUri, initial.DatasetUri) || dcl.IsZeroValue(des.DatasetUri) {
		cDes.DatasetUri = initial.DatasetUri
	} else {
		cDes.DatasetUri = des.DatasetUri
	}
	cDes.SchemaConfig = canonicalizeFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(des.SchemaConfig, initial.SchemaConfig, opts...)

	return cDes
}

func canonicalizeFhirStoreStreamConfigsBigqueryDestinationSlice(des, initial []FhirStoreStreamConfigsBigqueryDestination, opts ...dcl.ApplyOption) []FhirStoreStreamConfigsBigqueryDestination {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FhirStoreStreamConfigsBigqueryDestination, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFhirStoreStreamConfigsBigqueryDestination(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FhirStoreStreamConfigsBigqueryDestination, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFhirStoreStreamConfigsBigqueryDestination(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFhirStoreStreamConfigsBigqueryDestination(c *Client, des, nw *FhirStoreStreamConfigsBigqueryDestination) *FhirStoreStreamConfigsBigqueryDestination {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FhirStoreStreamConfigsBigqueryDestination while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.DatasetUri, nw.DatasetUri) {
		nw.DatasetUri = des.DatasetUri
	}
	nw.SchemaConfig = canonicalizeNewFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c, des.SchemaConfig, nw.SchemaConfig)

	return nw
}

func canonicalizeNewFhirStoreStreamConfigsBigqueryDestinationSet(c *Client, des, nw []FhirStoreStreamConfigsBigqueryDestination) []FhirStoreStreamConfigsBigqueryDestination {
	if des == nil {
		return nw
	}
	var reorderedNew []FhirStoreStreamConfigsBigqueryDestination
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFhirStoreStreamConfigsBigqueryDestinationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFhirStoreStreamConfigsBigqueryDestinationSlice(c *Client, des, nw []FhirStoreStreamConfigsBigqueryDestination) []FhirStoreStreamConfigsBigqueryDestination {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FhirStoreStreamConfigsBigqueryDestination
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFhirStoreStreamConfigsBigqueryDestination(c, &d, &n))
	}

	return items
}

func canonicalizeFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(des, initial *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig, opts ...dcl.ApplyOption) *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}

	if dcl.IsZeroValue(des.SchemaType) || (dcl.IsEmptyValueIndirect(des.SchemaType) && dcl.IsEmptyValueIndirect(initial.SchemaType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SchemaType = initial.SchemaType
	} else {
		cDes.SchemaType = des.SchemaType
	}
	if dcl.IsZeroValue(des.RecursiveStructureDepth) || (dcl.IsEmptyValueIndirect(des.RecursiveStructureDepth) && dcl.IsEmptyValueIndirect(initial.RecursiveStructureDepth)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RecursiveStructureDepth = initial.RecursiveStructureDepth
	} else {
		cDes.RecursiveStructureDepth = des.RecursiveStructureDepth
	}

	return cDes
}

func canonicalizeFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSlice(des, initial []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig, opts ...dcl.ApplyOption) []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FhirStoreStreamConfigsBigqueryDestinationSchemaConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FhirStoreStreamConfigsBigqueryDestinationSchemaConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c *Client, des, nw *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FhirStoreStreamConfigsBigqueryDestinationSchemaConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSet(c *Client, des, nw []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSlice(c *Client, des, nw []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c, &d, &n))
	}

	return items
}

func canonicalizeFhirStoreValidationConfig(des, initial *FhirStoreValidationConfig, opts ...dcl.ApplyOption) *FhirStoreValidationConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FhirStoreValidationConfig{}

	if dcl.BoolCanonicalize(des.DisableProfileValidation, initial.DisableProfileValidation) || dcl.IsZeroValue(des.DisableProfileValidation) {
		cDes.DisableProfileValidation = initial.DisableProfileValidation
	} else {
		cDes.DisableProfileValidation = des.DisableProfileValidation
	}
	if dcl.StringArrayCanonicalize(des.EnabledImplementationGuides, initial.EnabledImplementationGuides) {
		cDes.EnabledImplementationGuides = initial.EnabledImplementationGuides
	} else {
		cDes.EnabledImplementationGuides = des.EnabledImplementationGuides
	}
	if dcl.BoolCanonicalize(des.DisableRequiredFieldValidation, initial.DisableRequiredFieldValidation) || dcl.IsZeroValue(des.DisableRequiredFieldValidation) {
		cDes.DisableRequiredFieldValidation = initial.DisableRequiredFieldValidation
	} else {
		cDes.DisableRequiredFieldValidation = des.DisableRequiredFieldValidation
	}
	if dcl.BoolCanonicalize(des.DisableReferenceTypeValidation, initial.DisableReferenceTypeValidation) || dcl.IsZeroValue(des.DisableReferenceTypeValidation) {
		cDes.DisableReferenceTypeValidation = initial.DisableReferenceTypeValidation
	} else {
		cDes.DisableReferenceTypeValidation = des.DisableReferenceTypeValidation
	}
	if dcl.BoolCanonicalize(des.DisableFhirpathValidation, initial.DisableFhirpathValidation) || dcl.IsZeroValue(des.DisableFhirpathValidation) {
		cDes.DisableFhirpathValidation = initial.DisableFhirpathValidation
	} else {
		cDes.DisableFhirpathValidation = des.DisableFhirpathValidation
	}

	return cDes
}

func canonicalizeFhirStoreValidationConfigSlice(des, initial []FhirStoreValidationConfig, opts ...dcl.ApplyOption) []FhirStoreValidationConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FhirStoreValidationConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFhirStoreValidationConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FhirStoreValidationConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFhirStoreValidationConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFhirStoreValidationConfig(c *Client, des, nw *FhirStoreValidationConfig) *FhirStoreValidationConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FhirStoreValidationConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.DisableProfileValidation, nw.DisableProfileValidation) {
		nw.DisableProfileValidation = des.DisableProfileValidation
	}
	if dcl.StringArrayCanonicalize(des.EnabledImplementationGuides, nw.EnabledImplementationGuides) {
		nw.EnabledImplementationGuides = des.EnabledImplementationGuides
	}
	if dcl.BoolCanonicalize(des.DisableRequiredFieldValidation, nw.DisableRequiredFieldValidation) {
		nw.DisableRequiredFieldValidation = des.DisableRequiredFieldValidation
	}
	if dcl.BoolCanonicalize(des.DisableReferenceTypeValidation, nw.DisableReferenceTypeValidation) {
		nw.DisableReferenceTypeValidation = des.DisableReferenceTypeValidation
	}
	if dcl.BoolCanonicalize(des.DisableFhirpathValidation, nw.DisableFhirpathValidation) {
		nw.DisableFhirpathValidation = des.DisableFhirpathValidation
	}

	return nw
}

func canonicalizeNewFhirStoreValidationConfigSet(c *Client, des, nw []FhirStoreValidationConfig) []FhirStoreValidationConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []FhirStoreValidationConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFhirStoreValidationConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFhirStoreValidationConfigSlice(c *Client, des, nw []FhirStoreValidationConfig) []FhirStoreValidationConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FhirStoreValidationConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFhirStoreValidationConfig(c, &d, &n))
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
func diffFhirStore(c *Client, desired, actual *FhirStore, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.EnableUpdateCreate, actual.EnableUpdateCreate, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("EnableUpdateCreate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NotificationConfig, actual.NotificationConfig, dcl.DiffInfo{ObjectFunction: compareFhirStoreNotificationConfigNewStyle, EmptyObject: EmptyFhirStoreNotificationConfig, OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("NotificationConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisableReferentialIntegrity, actual.DisableReferentialIntegrity, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisableReferentialIntegrity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ShardNum, actual.ShardNum, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ShardNum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisableResourceVersioning, actual.DisableResourceVersioning, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisableResourceVersioning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StreamConfigs, actual.StreamConfigs, dcl.DiffInfo{ObjectFunction: compareFhirStoreStreamConfigsNewStyle, EmptyObject: EmptyFhirStoreStreamConfigs, OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("StreamConfigs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValidationConfig, actual.ValidationConfig, dcl.DiffInfo{ObjectFunction: compareFhirStoreValidationConfigNewStyle, EmptyObject: EmptyFhirStoreValidationConfig, OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("ValidationConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultSearchHandlingStrict, actual.DefaultSearchHandlingStrict, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("DefaultSearchHandlingStrict")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ComplexDataTypeReferenceParsing, actual.ComplexDataTypeReferenceParsing, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("ComplexDataTypeReferenceParsing")); len(ds) != 0 || err != nil {
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
func compareFhirStoreNotificationConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FhirStoreNotificationConfig)
	if !ok {
		desiredNotPointer, ok := d.(FhirStoreNotificationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreNotificationConfig or *FhirStoreNotificationConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FhirStoreNotificationConfig)
	if !ok {
		actualNotPointer, ok := a.(FhirStoreNotificationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreNotificationConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PubsubTopic, actual.PubsubTopic, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("PubsubTopic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFhirStoreStreamConfigsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FhirStoreStreamConfigs)
	if !ok {
		desiredNotPointer, ok := d.(FhirStoreStreamConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreStreamConfigs or *FhirStoreStreamConfigs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FhirStoreStreamConfigs)
	if !ok {
		actualNotPointer, ok := a.(FhirStoreStreamConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreStreamConfigs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceTypes, actual.ResourceTypes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("ResourceTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BigqueryDestination, actual.BigqueryDestination, dcl.DiffInfo{ObjectFunction: compareFhirStoreStreamConfigsBigqueryDestinationNewStyle, EmptyObject: EmptyFhirStoreStreamConfigsBigqueryDestination, OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("BigqueryDestination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFhirStoreStreamConfigsBigqueryDestinationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FhirStoreStreamConfigsBigqueryDestination)
	if !ok {
		desiredNotPointer, ok := d.(FhirStoreStreamConfigsBigqueryDestination)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreStreamConfigsBigqueryDestination or *FhirStoreStreamConfigsBigqueryDestination", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FhirStoreStreamConfigsBigqueryDestination)
	if !ok {
		actualNotPointer, ok := a.(FhirStoreStreamConfigsBigqueryDestination)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreStreamConfigsBigqueryDestination", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DatasetUri, actual.DatasetUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("DatasetUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SchemaConfig, actual.SchemaConfig, dcl.DiffInfo{ObjectFunction: compareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigNewStyle, EmptyObject: EmptyFhirStoreStreamConfigsBigqueryDestinationSchemaConfig, OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("SchemaConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FhirStoreStreamConfigsBigqueryDestinationSchemaConfig)
	if !ok {
		desiredNotPointer, ok := d.(FhirStoreStreamConfigsBigqueryDestinationSchemaConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreStreamConfigsBigqueryDestinationSchemaConfig or *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FhirStoreStreamConfigsBigqueryDestinationSchemaConfig)
	if !ok {
		actualNotPointer, ok := a.(FhirStoreStreamConfigsBigqueryDestinationSchemaConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreStreamConfigsBigqueryDestinationSchemaConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SchemaType, actual.SchemaType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("SchemaType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RecursiveStructureDepth, actual.RecursiveStructureDepth, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("RecursiveStructureDepth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFhirStoreValidationConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FhirStoreValidationConfig)
	if !ok {
		desiredNotPointer, ok := d.(FhirStoreValidationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreValidationConfig or *FhirStoreValidationConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FhirStoreValidationConfig)
	if !ok {
		actualNotPointer, ok := a.(FhirStoreValidationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FhirStoreValidationConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DisableProfileValidation, actual.DisableProfileValidation, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("DisableProfileValidation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnabledImplementationGuides, actual.EnabledImplementationGuides, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("EnabledImplementationGuides")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisableRequiredFieldValidation, actual.DisableRequiredFieldValidation, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("DisableRequiredFieldValidation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisableReferenceTypeValidation, actual.DisableReferenceTypeValidation, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("DisableReferenceTypeValidation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisableFhirpathValidation, actual.DisableFhirpathValidation, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFhirStoreUpdateFhirStoreOperation")}, fn.AddNest("DisableFhirpathValidation")); len(ds) != 0 || err != nil {
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
func (r *FhirStore) urlNormalized() *FhirStore {
	normalized := dcl.Copy(*r).(FhirStore)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Dataset = dcl.SelfLinkToName(r.Dataset)
	return &normalized
}

func (r *FhirStore) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateFhirStore" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"dataset":  dcl.ValueOrEmptyString(nr.Dataset),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/fhirStores/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the FhirStore resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *FhirStore) marshal(c *Client) ([]byte, error) {
	m, err := expandFhirStore(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling FhirStore: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFhirStore decodes JSON responses into the FhirStore resource schema.
func unmarshalFhirStore(b []byte, c *Client, res *FhirStore) (*FhirStore, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFhirStore(m, c, res)
}

func unmarshalMapFhirStore(m map[string]interface{}, c *Client, res *FhirStore) (*FhirStore, error) {

	flattened := flattenFhirStore(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandFhirStore expands FhirStore into a JSON request object.
func expandFhirStore(c *Client, f *FhirStore) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.EnableUpdateCreate; dcl.ValueShouldBeSent(v) {
		m["enableUpdateCreate"] = v
	}
	if v, err := expandFhirStoreNotificationConfig(c, f.NotificationConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding NotificationConfig into notificationConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["notificationConfig"] = v
	}
	if v := f.DisableReferentialIntegrity; dcl.ValueShouldBeSent(v) {
		m["disableReferentialIntegrity"] = v
	}
	if v := f.ShardNum; dcl.ValueShouldBeSent(v) {
		m["shardNum"] = v
	}
	if v := f.DisableResourceVersioning; dcl.ValueShouldBeSent(v) {
		m["disableResourceVersioning"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v := f.Version; dcl.ValueShouldBeSent(v) {
		m["version"] = v
	}
	if v, err := expandFhirStoreStreamConfigsSlice(c, f.StreamConfigs, res); err != nil {
		return nil, fmt.Errorf("error expanding StreamConfigs into streamConfigs: %w", err)
	} else if v != nil {
		m["streamConfigs"] = v
	}
	if v, err := expandFhirStoreValidationConfig(c, f.ValidationConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding ValidationConfig into validationConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["validationConfig"] = v
	}
	if v := f.DefaultSearchHandlingStrict; dcl.ValueShouldBeSent(v) {
		m["defaultSearchHandlingStrict"] = v
	}
	if v := f.ComplexDataTypeReferenceParsing; dcl.ValueShouldBeSent(v) {
		m["complexDataTypeReferenceParsing"] = v
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

// flattenFhirStore flattens FhirStore from a JSON request object into the
// FhirStore type.
func flattenFhirStore(c *Client, i interface{}, res *FhirStore) *FhirStore {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &FhirStore{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.EnableUpdateCreate = dcl.FlattenBool(m["enableUpdateCreate"])
	resultRes.NotificationConfig = flattenFhirStoreNotificationConfig(c, m["notificationConfig"], res)
	resultRes.DisableReferentialIntegrity = dcl.FlattenBool(m["disableReferentialIntegrity"])
	resultRes.ShardNum = dcl.FlattenInteger(m["shardNum"])
	resultRes.DisableResourceVersioning = dcl.FlattenBool(m["disableResourceVersioning"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Version = flattenFhirStoreVersionEnum(m["version"])
	resultRes.StreamConfigs = flattenFhirStoreStreamConfigsSlice(c, m["streamConfigs"], res)
	resultRes.ValidationConfig = flattenFhirStoreValidationConfig(c, m["validationConfig"], res)
	resultRes.DefaultSearchHandlingStrict = dcl.FlattenBool(m["defaultSearchHandlingStrict"])
	resultRes.ComplexDataTypeReferenceParsing = flattenFhirStoreComplexDataTypeReferenceParsingEnum(m["complexDataTypeReferenceParsing"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Dataset = dcl.FlattenString(m["dataset"])

	return resultRes
}

// expandFhirStoreNotificationConfigMap expands the contents of FhirStoreNotificationConfig into a JSON
// request object.
func expandFhirStoreNotificationConfigMap(c *Client, f map[string]FhirStoreNotificationConfig, res *FhirStore) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFhirStoreNotificationConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFhirStoreNotificationConfigSlice expands the contents of FhirStoreNotificationConfig into a JSON
// request object.
func expandFhirStoreNotificationConfigSlice(c *Client, f []FhirStoreNotificationConfig, res *FhirStore) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFhirStoreNotificationConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFhirStoreNotificationConfigMap flattens the contents of FhirStoreNotificationConfig from a JSON
// response object.
func flattenFhirStoreNotificationConfigMap(c *Client, i interface{}, res *FhirStore) map[string]FhirStoreNotificationConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FhirStoreNotificationConfig{}
	}

	if len(a) == 0 {
		return map[string]FhirStoreNotificationConfig{}
	}

	items := make(map[string]FhirStoreNotificationConfig)
	for k, item := range a {
		items[k] = *flattenFhirStoreNotificationConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFhirStoreNotificationConfigSlice flattens the contents of FhirStoreNotificationConfig from a JSON
// response object.
func flattenFhirStoreNotificationConfigSlice(c *Client, i interface{}, res *FhirStore) []FhirStoreNotificationConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []FhirStoreNotificationConfig{}
	}

	if len(a) == 0 {
		return []FhirStoreNotificationConfig{}
	}

	items := make([]FhirStoreNotificationConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFhirStoreNotificationConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFhirStoreNotificationConfig expands an instance of FhirStoreNotificationConfig into a JSON
// request object.
func expandFhirStoreNotificationConfig(c *Client, f *FhirStoreNotificationConfig, res *FhirStore) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PubsubTopic; !dcl.IsEmptyValueIndirect(v) {
		m["pubsubTopic"] = v
	}

	return m, nil
}

// flattenFhirStoreNotificationConfig flattens an instance of FhirStoreNotificationConfig from a JSON
// response object.
func flattenFhirStoreNotificationConfig(c *Client, i interface{}, res *FhirStore) *FhirStoreNotificationConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FhirStoreNotificationConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFhirStoreNotificationConfig
	}
	r.PubsubTopic = dcl.FlattenString(m["pubsubTopic"])

	return r
}

// expandFhirStoreStreamConfigsMap expands the contents of FhirStoreStreamConfigs into a JSON
// request object.
func expandFhirStoreStreamConfigsMap(c *Client, f map[string]FhirStoreStreamConfigs, res *FhirStore) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFhirStoreStreamConfigs(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFhirStoreStreamConfigsSlice expands the contents of FhirStoreStreamConfigs into a JSON
// request object.
func expandFhirStoreStreamConfigsSlice(c *Client, f []FhirStoreStreamConfigs, res *FhirStore) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFhirStoreStreamConfigs(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFhirStoreStreamConfigsMap flattens the contents of FhirStoreStreamConfigs from a JSON
// response object.
func flattenFhirStoreStreamConfigsMap(c *Client, i interface{}, res *FhirStore) map[string]FhirStoreStreamConfigs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FhirStoreStreamConfigs{}
	}

	if len(a) == 0 {
		return map[string]FhirStoreStreamConfigs{}
	}

	items := make(map[string]FhirStoreStreamConfigs)
	for k, item := range a {
		items[k] = *flattenFhirStoreStreamConfigs(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFhirStoreStreamConfigsSlice flattens the contents of FhirStoreStreamConfigs from a JSON
// response object.
func flattenFhirStoreStreamConfigsSlice(c *Client, i interface{}, res *FhirStore) []FhirStoreStreamConfigs {
	a, ok := i.([]interface{})
	if !ok {
		return []FhirStoreStreamConfigs{}
	}

	if len(a) == 0 {
		return []FhirStoreStreamConfigs{}
	}

	items := make([]FhirStoreStreamConfigs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFhirStoreStreamConfigs(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFhirStoreStreamConfigs expands an instance of FhirStoreStreamConfigs into a JSON
// request object.
func expandFhirStoreStreamConfigs(c *Client, f *FhirStoreStreamConfigs, res *FhirStore) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceTypes; v != nil {
		m["resourceTypes"] = v
	}
	if v, err := expandFhirStoreStreamConfigsBigqueryDestination(c, f.BigqueryDestination, res); err != nil {
		return nil, fmt.Errorf("error expanding BigqueryDestination into bigqueryDestination: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bigqueryDestination"] = v
	}

	return m, nil
}

// flattenFhirStoreStreamConfigs flattens an instance of FhirStoreStreamConfigs from a JSON
// response object.
func flattenFhirStoreStreamConfigs(c *Client, i interface{}, res *FhirStore) *FhirStoreStreamConfigs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FhirStoreStreamConfigs{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFhirStoreStreamConfigs
	}
	r.ResourceTypes = dcl.FlattenStringSlice(m["resourceTypes"])
	r.BigqueryDestination = flattenFhirStoreStreamConfigsBigqueryDestination(c, m["bigqueryDestination"], res)

	return r
}

// expandFhirStoreStreamConfigsBigqueryDestinationMap expands the contents of FhirStoreStreamConfigsBigqueryDestination into a JSON
// request object.
func expandFhirStoreStreamConfigsBigqueryDestinationMap(c *Client, f map[string]FhirStoreStreamConfigsBigqueryDestination, res *FhirStore) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFhirStoreStreamConfigsBigqueryDestination(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFhirStoreStreamConfigsBigqueryDestinationSlice expands the contents of FhirStoreStreamConfigsBigqueryDestination into a JSON
// request object.
func expandFhirStoreStreamConfigsBigqueryDestinationSlice(c *Client, f []FhirStoreStreamConfigsBigqueryDestination, res *FhirStore) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFhirStoreStreamConfigsBigqueryDestination(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFhirStoreStreamConfigsBigqueryDestinationMap flattens the contents of FhirStoreStreamConfigsBigqueryDestination from a JSON
// response object.
func flattenFhirStoreStreamConfigsBigqueryDestinationMap(c *Client, i interface{}, res *FhirStore) map[string]FhirStoreStreamConfigsBigqueryDestination {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FhirStoreStreamConfigsBigqueryDestination{}
	}

	if len(a) == 0 {
		return map[string]FhirStoreStreamConfigsBigqueryDestination{}
	}

	items := make(map[string]FhirStoreStreamConfigsBigqueryDestination)
	for k, item := range a {
		items[k] = *flattenFhirStoreStreamConfigsBigqueryDestination(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFhirStoreStreamConfigsBigqueryDestinationSlice flattens the contents of FhirStoreStreamConfigsBigqueryDestination from a JSON
// response object.
func flattenFhirStoreStreamConfigsBigqueryDestinationSlice(c *Client, i interface{}, res *FhirStore) []FhirStoreStreamConfigsBigqueryDestination {
	a, ok := i.([]interface{})
	if !ok {
		return []FhirStoreStreamConfigsBigqueryDestination{}
	}

	if len(a) == 0 {
		return []FhirStoreStreamConfigsBigqueryDestination{}
	}

	items := make([]FhirStoreStreamConfigsBigqueryDestination, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFhirStoreStreamConfigsBigqueryDestination(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFhirStoreStreamConfigsBigqueryDestination expands an instance of FhirStoreStreamConfigsBigqueryDestination into a JSON
// request object.
func expandFhirStoreStreamConfigsBigqueryDestination(c *Client, f *FhirStoreStreamConfigsBigqueryDestination, res *FhirStore) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DatasetUri; !dcl.IsEmptyValueIndirect(v) {
		m["datasetUri"] = v
	}
	if v, err := expandFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c, f.SchemaConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding SchemaConfig into schemaConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["schemaConfig"] = v
	}

	return m, nil
}

// flattenFhirStoreStreamConfigsBigqueryDestination flattens an instance of FhirStoreStreamConfigsBigqueryDestination from a JSON
// response object.
func flattenFhirStoreStreamConfigsBigqueryDestination(c *Client, i interface{}, res *FhirStore) *FhirStoreStreamConfigsBigqueryDestination {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FhirStoreStreamConfigsBigqueryDestination{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFhirStoreStreamConfigsBigqueryDestination
	}
	r.DatasetUri = dcl.FlattenString(m["datasetUri"])
	r.SchemaConfig = flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c, m["schemaConfig"], res)

	return r
}

// expandFhirStoreStreamConfigsBigqueryDestinationSchemaConfigMap expands the contents of FhirStoreStreamConfigsBigqueryDestinationSchemaConfig into a JSON
// request object.
func expandFhirStoreStreamConfigsBigqueryDestinationSchemaConfigMap(c *Client, f map[string]FhirStoreStreamConfigsBigqueryDestinationSchemaConfig, res *FhirStore) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSlice expands the contents of FhirStoreStreamConfigsBigqueryDestinationSchemaConfig into a JSON
// request object.
func expandFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSlice(c *Client, f []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig, res *FhirStore) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigMap flattens the contents of FhirStoreStreamConfigsBigqueryDestinationSchemaConfig from a JSON
// response object.
func flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigMap(c *Client, i interface{}, res *FhirStore) map[string]FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
	}

	if len(a) == 0 {
		return map[string]FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
	}

	items := make(map[string]FhirStoreStreamConfigsBigqueryDestinationSchemaConfig)
	for k, item := range a {
		items[k] = *flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSlice flattens the contents of FhirStoreStreamConfigsBigqueryDestinationSchemaConfig from a JSON
// response object.
func flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSlice(c *Client, i interface{}, res *FhirStore) []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
	}

	if len(a) == 0 {
		return []FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
	}

	items := make([]FhirStoreStreamConfigsBigqueryDestinationSchemaConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFhirStoreStreamConfigsBigqueryDestinationSchemaConfig expands an instance of FhirStoreStreamConfigsBigqueryDestinationSchemaConfig into a JSON
// request object.
func expandFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c *Client, f *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig, res *FhirStore) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SchemaType; !dcl.IsEmptyValueIndirect(v) {
		m["schemaType"] = v
	}
	if v := f.RecursiveStructureDepth; !dcl.IsEmptyValueIndirect(v) {
		m["recursiveStructureDepth"] = v
	}

	return m, nil
}

// flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfig flattens an instance of FhirStoreStreamConfigsBigqueryDestinationSchemaConfig from a JSON
// response object.
func flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(c *Client, i interface{}, res *FhirStore) *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFhirStoreStreamConfigsBigqueryDestinationSchemaConfig
	}
	r.SchemaType = flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(m["schemaType"])
	r.RecursiveStructureDepth = dcl.FlattenInteger(m["recursiveStructureDepth"])

	return r
}

// expandFhirStoreValidationConfigMap expands the contents of FhirStoreValidationConfig into a JSON
// request object.
func expandFhirStoreValidationConfigMap(c *Client, f map[string]FhirStoreValidationConfig, res *FhirStore) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFhirStoreValidationConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFhirStoreValidationConfigSlice expands the contents of FhirStoreValidationConfig into a JSON
// request object.
func expandFhirStoreValidationConfigSlice(c *Client, f []FhirStoreValidationConfig, res *FhirStore) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFhirStoreValidationConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFhirStoreValidationConfigMap flattens the contents of FhirStoreValidationConfig from a JSON
// response object.
func flattenFhirStoreValidationConfigMap(c *Client, i interface{}, res *FhirStore) map[string]FhirStoreValidationConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FhirStoreValidationConfig{}
	}

	if len(a) == 0 {
		return map[string]FhirStoreValidationConfig{}
	}

	items := make(map[string]FhirStoreValidationConfig)
	for k, item := range a {
		items[k] = *flattenFhirStoreValidationConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFhirStoreValidationConfigSlice flattens the contents of FhirStoreValidationConfig from a JSON
// response object.
func flattenFhirStoreValidationConfigSlice(c *Client, i interface{}, res *FhirStore) []FhirStoreValidationConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []FhirStoreValidationConfig{}
	}

	if len(a) == 0 {
		return []FhirStoreValidationConfig{}
	}

	items := make([]FhirStoreValidationConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFhirStoreValidationConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFhirStoreValidationConfig expands an instance of FhirStoreValidationConfig into a JSON
// request object.
func expandFhirStoreValidationConfig(c *Client, f *FhirStoreValidationConfig, res *FhirStore) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DisableProfileValidation; !dcl.IsEmptyValueIndirect(v) {
		m["disableProfileValidation"] = v
	}
	if v := f.EnabledImplementationGuides; v != nil {
		m["enabledImplementationGuides"] = v
	}
	if v := f.DisableRequiredFieldValidation; !dcl.IsEmptyValueIndirect(v) {
		m["disableRequiredFieldValidation"] = v
	}
	if v := f.DisableReferenceTypeValidation; !dcl.IsEmptyValueIndirect(v) {
		m["disableReferenceTypeValidation"] = v
	}
	if v := f.DisableFhirpathValidation; !dcl.IsEmptyValueIndirect(v) {
		m["disableFhirpathValidation"] = v
	}

	return m, nil
}

// flattenFhirStoreValidationConfig flattens an instance of FhirStoreValidationConfig from a JSON
// response object.
func flattenFhirStoreValidationConfig(c *Client, i interface{}, res *FhirStore) *FhirStoreValidationConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FhirStoreValidationConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFhirStoreValidationConfig
	}
	r.DisableProfileValidation = dcl.FlattenBool(m["disableProfileValidation"])
	r.EnabledImplementationGuides = dcl.FlattenStringSlice(m["enabledImplementationGuides"])
	r.DisableRequiredFieldValidation = dcl.FlattenBool(m["disableRequiredFieldValidation"])
	r.DisableReferenceTypeValidation = dcl.FlattenBool(m["disableReferenceTypeValidation"])
	r.DisableFhirpathValidation = dcl.FlattenBool(m["disableFhirpathValidation"])

	return r
}

// flattenFhirStoreVersionEnumMap flattens the contents of FhirStoreVersionEnum from a JSON
// response object.
func flattenFhirStoreVersionEnumMap(c *Client, i interface{}, res *FhirStore) map[string]FhirStoreVersionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FhirStoreVersionEnum{}
	}

	if len(a) == 0 {
		return map[string]FhirStoreVersionEnum{}
	}

	items := make(map[string]FhirStoreVersionEnum)
	for k, item := range a {
		items[k] = *flattenFhirStoreVersionEnum(item.(interface{}))
	}

	return items
}

// flattenFhirStoreVersionEnumSlice flattens the contents of FhirStoreVersionEnum from a JSON
// response object.
func flattenFhirStoreVersionEnumSlice(c *Client, i interface{}, res *FhirStore) []FhirStoreVersionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FhirStoreVersionEnum{}
	}

	if len(a) == 0 {
		return []FhirStoreVersionEnum{}
	}

	items := make([]FhirStoreVersionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFhirStoreVersionEnum(item.(interface{})))
	}

	return items
}

// flattenFhirStoreVersionEnum asserts that an interface is a string, and returns a
// pointer to a *FhirStoreVersionEnum with the same value as that string.
func flattenFhirStoreVersionEnum(i interface{}) *FhirStoreVersionEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FhirStoreVersionEnumRef(s)
}

// flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumMap flattens the contents of FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum from a JSON
// response object.
func flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumMap(c *Client, i interface{}, res *FhirStore) map[string]FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum{}
	}

	items := make(map[string]FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum)
	for k, item := range a {
		items[k] = *flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(item.(interface{}))
	}

	return items
}

// flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumSlice flattens the contents of FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum from a JSON
// response object.
func flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumSlice(c *Client, i interface{}, res *FhirStore) []FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum{}
	}

	if len(a) == 0 {
		return []FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum{}
	}

	items := make([]FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(item.(interface{})))
	}

	return items
}

// flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum asserts that an interface is a string, and returns a
// pointer to a *FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum with the same value as that string.
func flattenFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(i interface{}) *FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumRef(s)
}

// flattenFhirStoreComplexDataTypeReferenceParsingEnumMap flattens the contents of FhirStoreComplexDataTypeReferenceParsingEnum from a JSON
// response object.
func flattenFhirStoreComplexDataTypeReferenceParsingEnumMap(c *Client, i interface{}, res *FhirStore) map[string]FhirStoreComplexDataTypeReferenceParsingEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FhirStoreComplexDataTypeReferenceParsingEnum{}
	}

	if len(a) == 0 {
		return map[string]FhirStoreComplexDataTypeReferenceParsingEnum{}
	}

	items := make(map[string]FhirStoreComplexDataTypeReferenceParsingEnum)
	for k, item := range a {
		items[k] = *flattenFhirStoreComplexDataTypeReferenceParsingEnum(item.(interface{}))
	}

	return items
}

// flattenFhirStoreComplexDataTypeReferenceParsingEnumSlice flattens the contents of FhirStoreComplexDataTypeReferenceParsingEnum from a JSON
// response object.
func flattenFhirStoreComplexDataTypeReferenceParsingEnumSlice(c *Client, i interface{}, res *FhirStore) []FhirStoreComplexDataTypeReferenceParsingEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FhirStoreComplexDataTypeReferenceParsingEnum{}
	}

	if len(a) == 0 {
		return []FhirStoreComplexDataTypeReferenceParsingEnum{}
	}

	items := make([]FhirStoreComplexDataTypeReferenceParsingEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFhirStoreComplexDataTypeReferenceParsingEnum(item.(interface{})))
	}

	return items
}

// flattenFhirStoreComplexDataTypeReferenceParsingEnum asserts that an interface is a string, and returns a
// pointer to a *FhirStoreComplexDataTypeReferenceParsingEnum with the same value as that string.
func flattenFhirStoreComplexDataTypeReferenceParsingEnum(i interface{}) *FhirStoreComplexDataTypeReferenceParsingEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FhirStoreComplexDataTypeReferenceParsingEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *FhirStore) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFhirStore(b, c, r)
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

type fhirStoreDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         fhirStoreApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToFhirStoreDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]fhirStoreDiff, error) {
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
	var diffs []fhirStoreDiff
	// For each operation name, create a fhirStoreDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := fhirStoreDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToFhirStoreApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToFhirStoreApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (fhirStoreApiOperation, error) {
	switch opName {

	case "updateFhirStoreUpdateFhirStoreOperation":
		return &updateFhirStoreUpdateFhirStoreOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractFhirStoreFields(r *FhirStore) error {
	vNotificationConfig := r.NotificationConfig
	if vNotificationConfig == nil {
		// note: explicitly not the empty object.
		vNotificationConfig = &FhirStoreNotificationConfig{}
	}
	if err := extractFhirStoreNotificationConfigFields(r, vNotificationConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vNotificationConfig) {
		r.NotificationConfig = vNotificationConfig
	}
	vValidationConfig := r.ValidationConfig
	if vValidationConfig == nil {
		// note: explicitly not the empty object.
		vValidationConfig = &FhirStoreValidationConfig{}
	}
	if err := extractFhirStoreValidationConfigFields(r, vValidationConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vValidationConfig) {
		r.ValidationConfig = vValidationConfig
	}
	return nil
}
func extractFhirStoreNotificationConfigFields(r *FhirStore, o *FhirStoreNotificationConfig) error {
	return nil
}
func extractFhirStoreStreamConfigsFields(r *FhirStore, o *FhirStoreStreamConfigs) error {
	vBigqueryDestination := o.BigqueryDestination
	if vBigqueryDestination == nil {
		// note: explicitly not the empty object.
		vBigqueryDestination = &FhirStoreStreamConfigsBigqueryDestination{}
	}
	if err := extractFhirStoreStreamConfigsBigqueryDestinationFields(r, vBigqueryDestination); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBigqueryDestination) {
		o.BigqueryDestination = vBigqueryDestination
	}
	return nil
}
func extractFhirStoreStreamConfigsBigqueryDestinationFields(r *FhirStore, o *FhirStoreStreamConfigsBigqueryDestination) error {
	vSchemaConfig := o.SchemaConfig
	if vSchemaConfig == nil {
		// note: explicitly not the empty object.
		vSchemaConfig = &FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
	}
	if err := extractFhirStoreStreamConfigsBigqueryDestinationSchemaConfigFields(r, vSchemaConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSchemaConfig) {
		o.SchemaConfig = vSchemaConfig
	}
	return nil
}
func extractFhirStoreStreamConfigsBigqueryDestinationSchemaConfigFields(r *FhirStore, o *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) error {
	return nil
}
func extractFhirStoreValidationConfigFields(r *FhirStore, o *FhirStoreValidationConfig) error {
	return nil
}

func postReadExtractFhirStoreFields(r *FhirStore) error {
	vNotificationConfig := r.NotificationConfig
	if vNotificationConfig == nil {
		// note: explicitly not the empty object.
		vNotificationConfig = &FhirStoreNotificationConfig{}
	}
	if err := postReadExtractFhirStoreNotificationConfigFields(r, vNotificationConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vNotificationConfig) {
		r.NotificationConfig = vNotificationConfig
	}
	vValidationConfig := r.ValidationConfig
	if vValidationConfig == nil {
		// note: explicitly not the empty object.
		vValidationConfig = &FhirStoreValidationConfig{}
	}
	if err := postReadExtractFhirStoreValidationConfigFields(r, vValidationConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vValidationConfig) {
		r.ValidationConfig = vValidationConfig
	}
	return nil
}
func postReadExtractFhirStoreNotificationConfigFields(r *FhirStore, o *FhirStoreNotificationConfig) error {
	return nil
}
func postReadExtractFhirStoreStreamConfigsFields(r *FhirStore, o *FhirStoreStreamConfigs) error {
	vBigqueryDestination := o.BigqueryDestination
	if vBigqueryDestination == nil {
		// note: explicitly not the empty object.
		vBigqueryDestination = &FhirStoreStreamConfigsBigqueryDestination{}
	}
	if err := extractFhirStoreStreamConfigsBigqueryDestinationFields(r, vBigqueryDestination); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBigqueryDestination) {
		o.BigqueryDestination = vBigqueryDestination
	}
	return nil
}
func postReadExtractFhirStoreStreamConfigsBigqueryDestinationFields(r *FhirStore, o *FhirStoreStreamConfigsBigqueryDestination) error {
	vSchemaConfig := o.SchemaConfig
	if vSchemaConfig == nil {
		// note: explicitly not the empty object.
		vSchemaConfig = &FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
	}
	if err := extractFhirStoreStreamConfigsBigqueryDestinationSchemaConfigFields(r, vSchemaConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSchemaConfig) {
		o.SchemaConfig = vSchemaConfig
	}
	return nil
}
func postReadExtractFhirStoreStreamConfigsBigqueryDestinationSchemaConfigFields(r *FhirStore, o *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) error {
	return nil
}
func postReadExtractFhirStoreValidationConfigFields(r *FhirStore, o *FhirStoreValidationConfig) error {
	return nil
}
