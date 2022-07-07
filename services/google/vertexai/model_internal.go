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
package vertexai

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

func (r *Model) validate() error {

	if err := dcl.Required(r, "displayName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "containerSpec"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.OriginalModelInfo) {
		if err := r.OriginalModelInfo.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ContainerSpec) {
		if err := r.ContainerSpec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EncryptionSpec) {
		if err := r.EncryptionSpec.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ModelSupportedExportFormats) validate() error {
	return nil
}
func (r *ModelOriginalModelInfo) validate() error {
	return nil
}
func (r *ModelContainerSpec) validate() error {
	if err := dcl.Required(r, "imageUri"); err != nil {
		return err
	}
	return nil
}
func (r *ModelContainerSpecEnv) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	return nil
}
func (r *ModelContainerSpecPorts) validate() error {
	return nil
}
func (r *ModelDeployedModels) validate() error {
	return nil
}
func (r *ModelEncryptionSpec) validate() error {
	if err := dcl.Required(r, "kmsKeyName"); err != nil {
		return err
	}
	return nil
}
func (r *Model) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-aiplatform.googleapis.com/v1/", params)
}

func (r *Model) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/models/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Model) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/models", nr.basePath(), userBasePath, params), nil

}

func (r *Model) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/models:upload", nr.basePath(), userBasePath, params), nil

}

func (r *Model) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/models/{{name}}", nr.basePath(), userBasePath, params), nil
}

// modelApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type modelApiOperation interface {
	do(context.Context, *Model, *Client) error
}

// newUpdateModelUpdateModelRequest creates a request for an
// Model resource's UpdateModel update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateModelUpdateModelRequest(ctx context.Context, f *Model, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	b, err := c.getModelRaw(ctx, f)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	return req, nil
}

// marshalUpdateModelUpdateModelRequest converts the update into
// the final JSON request body.
func marshalUpdateModelUpdateModelRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateModelUpdateModelOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateModelUpdateModelOperation) do(ctx context.Context, r *Model, c *Client) error {
	_, err := c.GetModel(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateModel")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateModelUpdateModelRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateModelUpdateModelRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listModelRaw(ctx context.Context, r *Model, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ModelMaxPage {
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

type listModelOperation struct {
	Models []map[string]interface{} `json:"models"`
	Token  string                   `json:"nextPageToken"`
}

func (c *Client) listModel(ctx context.Context, r *Model, pageToken string, pageSize int32) ([]*Model, string, error) {
	b, err := c.listModelRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listModelOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Model
	for _, v := range m.Models {
		res, err := unmarshalMapModel(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllModel(ctx context.Context, f func(*Model) bool, resources []*Model) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteModel(ctx, res)
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

type deleteModelOperation struct{}

func (op *deleteModelOperation) do(ctx context.Context, r *Model, c *Client) error {
	r, err := c.GetModel(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Model not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetModel checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
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
		_, err := c.GetModel(ctx, r)
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
type createModelOperation struct {
	response map[string]interface{}
}

func (op *createModelOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getModelRaw(ctx context.Context, r *Model) ([]byte, error) {

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

func (c *Client) modelDiffsForRawDesired(ctx context.Context, rawDesired *Model, opts ...dcl.ApplyOption) (initial, desired *Model, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Model
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Model); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Model, got %T", sh)
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
		desired, err := canonicalizeModelDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetModel(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Model resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Model resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Model resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeModelDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Model: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Model: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractModelFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeModelInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Model: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeModelDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Model: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffModel(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeModelInitialState(rawInitial, rawDesired *Model) (*Model, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeModelDesiredState(rawDesired, rawInitial *Model, opts ...dcl.ApplyOption) (*Model, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.OriginalModelInfo = canonicalizeModelOriginalModelInfo(rawDesired.OriginalModelInfo, nil, opts...)
		rawDesired.ContainerSpec = canonicalizeModelContainerSpec(rawDesired.ContainerSpec, nil, opts...)
		rawDesired.EncryptionSpec = canonicalizeModelEncryptionSpec(rawDesired.EncryptionSpec, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Model{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.StringCanonicalize(rawDesired.VersionDescription, rawInitial.VersionDescription) {
		canonicalDesired.VersionDescription = rawInitial.VersionDescription
	} else {
		canonicalDesired.VersionDescription = rawDesired.VersionDescription
	}
	canonicalDesired.ContainerSpec = canonicalizeModelContainerSpec(rawDesired.ContainerSpec, rawInitial.ContainerSpec, opts...)
	if dcl.StringCanonicalize(rawDesired.ArtifactUri, rawInitial.ArtifactUri) {
		canonicalDesired.ArtifactUri = rawInitial.ArtifactUri
	} else {
		canonicalDesired.ArtifactUri = rawDesired.ArtifactUri
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.EncryptionSpec = canonicalizeModelEncryptionSpec(rawDesired.EncryptionSpec, rawInitial.EncryptionSpec, opts...)
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

func canonicalizeModelNewState(c *Client, rawNew, rawDesired *Model) (*Model, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.VersionId) && dcl.IsEmptyValueIndirect(rawDesired.VersionId) {
		rawNew.VersionId = rawDesired.VersionId
	} else {
		if dcl.StringCanonicalize(rawDesired.VersionId, rawNew.VersionId) {
			rawNew.VersionId = rawDesired.VersionId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VersionCreateTime) && dcl.IsEmptyValueIndirect(rawDesired.VersionCreateTime) {
		rawNew.VersionCreateTime = rawDesired.VersionCreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.VersionUpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.VersionUpdateTime) {
		rawNew.VersionUpdateTime = rawDesired.VersionUpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VersionDescription) && dcl.IsEmptyValueIndirect(rawDesired.VersionDescription) {
		rawNew.VersionDescription = rawDesired.VersionDescription
	} else {
		if dcl.StringCanonicalize(rawDesired.VersionDescription, rawNew.VersionDescription) {
			rawNew.VersionDescription = rawDesired.VersionDescription
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SupportedExportFormats) && dcl.IsEmptyValueIndirect(rawDesired.SupportedExportFormats) {
		rawNew.SupportedExportFormats = rawDesired.SupportedExportFormats
	} else {
		rawNew.SupportedExportFormats = canonicalizeNewModelSupportedExportFormatsSlice(c, rawDesired.SupportedExportFormats, rawNew.SupportedExportFormats)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TrainingPipeline) && dcl.IsEmptyValueIndirect(rawDesired.TrainingPipeline) {
		rawNew.TrainingPipeline = rawDesired.TrainingPipeline
	} else {
		if dcl.StringCanonicalize(rawDesired.TrainingPipeline, rawNew.TrainingPipeline) {
			rawNew.TrainingPipeline = rawDesired.TrainingPipeline
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.OriginalModelInfo) && dcl.IsEmptyValueIndirect(rawDesired.OriginalModelInfo) {
		rawNew.OriginalModelInfo = rawDesired.OriginalModelInfo
	} else {
		rawNew.OriginalModelInfo = canonicalizeNewModelOriginalModelInfo(c, rawDesired.OriginalModelInfo, rawNew.OriginalModelInfo)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ContainerSpec) && dcl.IsEmptyValueIndirect(rawDesired.ContainerSpec) {
		rawNew.ContainerSpec = rawDesired.ContainerSpec
	} else {
		rawNew.ContainerSpec = rawDesired.ContainerSpec
	}

	if dcl.IsEmptyValueIndirect(rawNew.ArtifactUri) && dcl.IsEmptyValueIndirect(rawDesired.ArtifactUri) {
		rawNew.ArtifactUri = rawDesired.ArtifactUri
	} else {
		if dcl.StringCanonicalize(rawDesired.ArtifactUri, rawNew.ArtifactUri) {
			rawNew.ArtifactUri = rawDesired.ArtifactUri
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SupportedDeploymentResourcesTypes) && dcl.IsEmptyValueIndirect(rawDesired.SupportedDeploymentResourcesTypes) {
		rawNew.SupportedDeploymentResourcesTypes = rawDesired.SupportedDeploymentResourcesTypes
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SupportedInputStorageFormats) && dcl.IsEmptyValueIndirect(rawDesired.SupportedInputStorageFormats) {
		rawNew.SupportedInputStorageFormats = rawDesired.SupportedInputStorageFormats
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.SupportedInputStorageFormats, rawNew.SupportedInputStorageFormats) {
			rawNew.SupportedInputStorageFormats = rawDesired.SupportedInputStorageFormats
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SupportedOutputStorageFormats) && dcl.IsEmptyValueIndirect(rawDesired.SupportedOutputStorageFormats) {
		rawNew.SupportedOutputStorageFormats = rawDesired.SupportedOutputStorageFormats
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.SupportedOutputStorageFormats, rawNew.SupportedOutputStorageFormats) {
			rawNew.SupportedOutputStorageFormats = rawDesired.SupportedOutputStorageFormats
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

	if dcl.IsEmptyValueIndirect(rawNew.DeployedModels) && dcl.IsEmptyValueIndirect(rawDesired.DeployedModels) {
		rawNew.DeployedModels = rawDesired.DeployedModels
	} else {
		rawNew.DeployedModels = canonicalizeNewModelDeployedModelsSlice(c, rawDesired.DeployedModels, rawNew.DeployedModels)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EncryptionSpec) && dcl.IsEmptyValueIndirect(rawDesired.EncryptionSpec) {
		rawNew.EncryptionSpec = rawDesired.EncryptionSpec
	} else {
		rawNew.EncryptionSpec = canonicalizeNewModelEncryptionSpec(c, rawDesired.EncryptionSpec, rawNew.EncryptionSpec)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeModelSupportedExportFormats(des, initial *ModelSupportedExportFormats, opts ...dcl.ApplyOption) *ModelSupportedExportFormats {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ModelSupportedExportFormats{}

	return cDes
}

func canonicalizeModelSupportedExportFormatsSlice(des, initial []ModelSupportedExportFormats, opts ...dcl.ApplyOption) []ModelSupportedExportFormats {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ModelSupportedExportFormats, 0, len(des))
		for _, d := range des {
			cd := canonicalizeModelSupportedExportFormats(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ModelSupportedExportFormats, 0, len(des))
	for i, d := range des {
		cd := canonicalizeModelSupportedExportFormats(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewModelSupportedExportFormats(c *Client, des, nw *ModelSupportedExportFormats) *ModelSupportedExportFormats {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ModelSupportedExportFormats while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}

	return nw
}

func canonicalizeNewModelSupportedExportFormatsSet(c *Client, des, nw []ModelSupportedExportFormats) []ModelSupportedExportFormats {
	if des == nil {
		return nw
	}
	var reorderedNew []ModelSupportedExportFormats
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareModelSupportedExportFormatsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewModelSupportedExportFormatsSlice(c *Client, des, nw []ModelSupportedExportFormats) []ModelSupportedExportFormats {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ModelSupportedExportFormats
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewModelSupportedExportFormats(c, &d, &n))
	}

	return items
}

func canonicalizeModelOriginalModelInfo(des, initial *ModelOriginalModelInfo, opts ...dcl.ApplyOption) *ModelOriginalModelInfo {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ModelOriginalModelInfo{}

	return cDes
}

func canonicalizeModelOriginalModelInfoSlice(des, initial []ModelOriginalModelInfo, opts ...dcl.ApplyOption) []ModelOriginalModelInfo {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ModelOriginalModelInfo, 0, len(des))
		for _, d := range des {
			cd := canonicalizeModelOriginalModelInfo(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ModelOriginalModelInfo, 0, len(des))
	for i, d := range des {
		cd := canonicalizeModelOriginalModelInfo(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewModelOriginalModelInfo(c *Client, des, nw *ModelOriginalModelInfo) *ModelOriginalModelInfo {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ModelOriginalModelInfo while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewModelOriginalModelInfoSet(c *Client, des, nw []ModelOriginalModelInfo) []ModelOriginalModelInfo {
	if des == nil {
		return nw
	}
	var reorderedNew []ModelOriginalModelInfo
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareModelOriginalModelInfoNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewModelOriginalModelInfoSlice(c *Client, des, nw []ModelOriginalModelInfo) []ModelOriginalModelInfo {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ModelOriginalModelInfo
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewModelOriginalModelInfo(c, &d, &n))
	}

	return items
}

func canonicalizeModelContainerSpec(des, initial *ModelContainerSpec, opts ...dcl.ApplyOption) *ModelContainerSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ModelContainerSpec{}

	if dcl.StringCanonicalize(des.ImageUri, initial.ImageUri) || dcl.IsZeroValue(des.ImageUri) {
		cDes.ImageUri = initial.ImageUri
	} else {
		cDes.ImageUri = des.ImageUri
	}
	if dcl.StringArrayCanonicalize(des.Command, initial.Command) {
		cDes.Command = initial.Command
	} else {
		cDes.Command = des.Command
	}
	if dcl.StringArrayCanonicalize(des.Args, initial.Args) {
		cDes.Args = initial.Args
	} else {
		cDes.Args = des.Args
	}
	cDes.Env = canonicalizeModelContainerSpecEnvSlice(des.Env, initial.Env, opts...)
	cDes.Ports = canonicalizeModelContainerSpecPortsSlice(des.Ports, initial.Ports, opts...)
	if dcl.StringCanonicalize(des.PredictRoute, initial.PredictRoute) || dcl.IsZeroValue(des.PredictRoute) {
		cDes.PredictRoute = initial.PredictRoute
	} else {
		cDes.PredictRoute = des.PredictRoute
	}
	if dcl.StringCanonicalize(des.HealthRoute, initial.HealthRoute) || dcl.IsZeroValue(des.HealthRoute) {
		cDes.HealthRoute = initial.HealthRoute
	} else {
		cDes.HealthRoute = des.HealthRoute
	}

	return cDes
}

func canonicalizeModelContainerSpecSlice(des, initial []ModelContainerSpec, opts ...dcl.ApplyOption) []ModelContainerSpec {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ModelContainerSpec, 0, len(des))
		for _, d := range des {
			cd := canonicalizeModelContainerSpec(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ModelContainerSpec, 0, len(des))
	for i, d := range des {
		cd := canonicalizeModelContainerSpec(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewModelContainerSpec(c *Client, des, nw *ModelContainerSpec) *ModelContainerSpec {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ModelContainerSpec while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ImageUri, nw.ImageUri) {
		nw.ImageUri = des.ImageUri
	}
	if dcl.StringArrayCanonicalize(des.Command, nw.Command) {
		nw.Command = des.Command
	}
	if dcl.StringArrayCanonicalize(des.Args, nw.Args) {
		nw.Args = des.Args
	}
	nw.Env = canonicalizeNewModelContainerSpecEnvSlice(c, des.Env, nw.Env)
	nw.Ports = canonicalizeNewModelContainerSpecPortsSlice(c, des.Ports, nw.Ports)
	if dcl.StringCanonicalize(des.PredictRoute, nw.PredictRoute) {
		nw.PredictRoute = des.PredictRoute
	}
	if dcl.StringCanonicalize(des.HealthRoute, nw.HealthRoute) {
		nw.HealthRoute = des.HealthRoute
	}

	return nw
}

func canonicalizeNewModelContainerSpecSet(c *Client, des, nw []ModelContainerSpec) []ModelContainerSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []ModelContainerSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareModelContainerSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewModelContainerSpecSlice(c *Client, des, nw []ModelContainerSpec) []ModelContainerSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ModelContainerSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewModelContainerSpec(c, &d, &n))
	}

	return items
}

func canonicalizeModelContainerSpecEnv(des, initial *ModelContainerSpecEnv, opts ...dcl.ApplyOption) *ModelContainerSpecEnv {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ModelContainerSpecEnv{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}

	return cDes
}

func canonicalizeModelContainerSpecEnvSlice(des, initial []ModelContainerSpecEnv, opts ...dcl.ApplyOption) []ModelContainerSpecEnv {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ModelContainerSpecEnv, 0, len(des))
		for _, d := range des {
			cd := canonicalizeModelContainerSpecEnv(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ModelContainerSpecEnv, 0, len(des))
	for i, d := range des {
		cd := canonicalizeModelContainerSpecEnv(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewModelContainerSpecEnv(c *Client, des, nw *ModelContainerSpecEnv) *ModelContainerSpecEnv {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ModelContainerSpecEnv while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewModelContainerSpecEnvSet(c *Client, des, nw []ModelContainerSpecEnv) []ModelContainerSpecEnv {
	if des == nil {
		return nw
	}
	var reorderedNew []ModelContainerSpecEnv
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareModelContainerSpecEnvNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewModelContainerSpecEnvSlice(c *Client, des, nw []ModelContainerSpecEnv) []ModelContainerSpecEnv {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ModelContainerSpecEnv
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewModelContainerSpecEnv(c, &d, &n))
	}

	return items
}

func canonicalizeModelContainerSpecPorts(des, initial *ModelContainerSpecPorts, opts ...dcl.ApplyOption) *ModelContainerSpecPorts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ModelContainerSpecPorts{}

	if dcl.IsZeroValue(des.ContainerPort) || (dcl.IsEmptyValueIndirect(des.ContainerPort) && dcl.IsEmptyValueIndirect(initial.ContainerPort)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ContainerPort = initial.ContainerPort
	} else {
		cDes.ContainerPort = des.ContainerPort
	}

	return cDes
}

func canonicalizeModelContainerSpecPortsSlice(des, initial []ModelContainerSpecPorts, opts ...dcl.ApplyOption) []ModelContainerSpecPorts {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ModelContainerSpecPorts, 0, len(des))
		for _, d := range des {
			cd := canonicalizeModelContainerSpecPorts(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ModelContainerSpecPorts, 0, len(des))
	for i, d := range des {
		cd := canonicalizeModelContainerSpecPorts(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewModelContainerSpecPorts(c *Client, des, nw *ModelContainerSpecPorts) *ModelContainerSpecPorts {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ModelContainerSpecPorts while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewModelContainerSpecPortsSet(c *Client, des, nw []ModelContainerSpecPorts) []ModelContainerSpecPorts {
	if des == nil {
		return nw
	}
	var reorderedNew []ModelContainerSpecPorts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareModelContainerSpecPortsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewModelContainerSpecPortsSlice(c *Client, des, nw []ModelContainerSpecPorts) []ModelContainerSpecPorts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ModelContainerSpecPorts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewModelContainerSpecPorts(c, &d, &n))
	}

	return items
}

func canonicalizeModelDeployedModels(des, initial *ModelDeployedModels, opts ...dcl.ApplyOption) *ModelDeployedModels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ModelDeployedModels{}

	if dcl.IsZeroValue(des.Endpoint) || (dcl.IsEmptyValueIndirect(des.Endpoint) && dcl.IsEmptyValueIndirect(initial.Endpoint)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Endpoint = initial.Endpoint
	} else {
		cDes.Endpoint = des.Endpoint
	}
	if dcl.StringCanonicalize(des.DeployedModelId, initial.DeployedModelId) || dcl.IsZeroValue(des.DeployedModelId) {
		cDes.DeployedModelId = initial.DeployedModelId
	} else {
		cDes.DeployedModelId = des.DeployedModelId
	}

	return cDes
}

func canonicalizeModelDeployedModelsSlice(des, initial []ModelDeployedModels, opts ...dcl.ApplyOption) []ModelDeployedModels {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ModelDeployedModels, 0, len(des))
		for _, d := range des {
			cd := canonicalizeModelDeployedModels(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ModelDeployedModels, 0, len(des))
	for i, d := range des {
		cd := canonicalizeModelDeployedModels(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewModelDeployedModels(c *Client, des, nw *ModelDeployedModels) *ModelDeployedModels {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ModelDeployedModels while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.DeployedModelId, nw.DeployedModelId) {
		nw.DeployedModelId = des.DeployedModelId
	}

	return nw
}

func canonicalizeNewModelDeployedModelsSet(c *Client, des, nw []ModelDeployedModels) []ModelDeployedModels {
	if des == nil {
		return nw
	}
	var reorderedNew []ModelDeployedModels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareModelDeployedModelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewModelDeployedModelsSlice(c *Client, des, nw []ModelDeployedModels) []ModelDeployedModels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ModelDeployedModels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewModelDeployedModels(c, &d, &n))
	}

	return items
}

func canonicalizeModelEncryptionSpec(des, initial *ModelEncryptionSpec, opts ...dcl.ApplyOption) *ModelEncryptionSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ModelEncryptionSpec{}

	if dcl.IsZeroValue(des.KmsKeyName) || (dcl.IsEmptyValueIndirect(des.KmsKeyName) && dcl.IsEmptyValueIndirect(initial.KmsKeyName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}

	return cDes
}

func canonicalizeModelEncryptionSpecSlice(des, initial []ModelEncryptionSpec, opts ...dcl.ApplyOption) []ModelEncryptionSpec {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ModelEncryptionSpec, 0, len(des))
		for _, d := range des {
			cd := canonicalizeModelEncryptionSpec(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ModelEncryptionSpec, 0, len(des))
	for i, d := range des {
		cd := canonicalizeModelEncryptionSpec(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewModelEncryptionSpec(c *Client, des, nw *ModelEncryptionSpec) *ModelEncryptionSpec {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ModelEncryptionSpec while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewModelEncryptionSpecSet(c *Client, des, nw []ModelEncryptionSpec) []ModelEncryptionSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []ModelEncryptionSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareModelEncryptionSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewModelEncryptionSpecSlice(c *Client, des, nw []ModelEncryptionSpec) []ModelEncryptionSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ModelEncryptionSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewModelEncryptionSpec(c, &d, &n))
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
func diffModel(c *Client, desired, actual *Model, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VersionId, actual.VersionId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VersionCreateTime, actual.VersionCreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionCreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VersionUpdateTime, actual.VersionUpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionUpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateModelUpdateModelOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateModelUpdateModelOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VersionDescription, actual.VersionDescription, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionDescription")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SupportedExportFormats, actual.SupportedExportFormats, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareModelSupportedExportFormatsNewStyle, EmptyObject: EmptyModelSupportedExportFormats, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SupportedExportFormats")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrainingPipeline, actual.TrainingPipeline, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TrainingPipeline")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OriginalModelInfo, actual.OriginalModelInfo, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareModelOriginalModelInfoNewStyle, EmptyObject: EmptyModelOriginalModelInfo, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OriginalModelInfo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContainerSpec, actual.ContainerSpec, dcl.DiffInfo{ObjectFunction: compareModelContainerSpecNewStyle, EmptyObject: EmptyModelContainerSpec, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ContainerSpec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ArtifactUri, actual.ArtifactUri, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArtifactUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SupportedDeploymentResourcesTypes, actual.SupportedDeploymentResourcesTypes, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SupportedDeploymentResourcesTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SupportedInputStorageFormats, actual.SupportedInputStorageFormats, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SupportedInputStorageFormats")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SupportedOutputStorageFormats, actual.SupportedOutputStorageFormats, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SupportedOutputStorageFormats")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.DeployedModels, actual.DeployedModels, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareModelDeployedModelsNewStyle, EmptyObject: EmptyModelDeployedModels, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeployedModels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateModelUpdateModelOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncryptionSpec, actual.EncryptionSpec, dcl.DiffInfo{ObjectFunction: compareModelEncryptionSpecNewStyle, EmptyObject: EmptyModelEncryptionSpec, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncryptionSpec")); len(ds) != 0 || err != nil {
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
func compareModelSupportedExportFormatsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ModelSupportedExportFormats)
	if !ok {
		desiredNotPointer, ok := d.(ModelSupportedExportFormats)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelSupportedExportFormats or *ModelSupportedExportFormats", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ModelSupportedExportFormats)
	if !ok {
		actualNotPointer, ok := a.(ModelSupportedExportFormats)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelSupportedExportFormats", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExportableContents, actual.ExportableContents, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExportableContents")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareModelOriginalModelInfoNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ModelOriginalModelInfo)
	if !ok {
		desiredNotPointer, ok := d.(ModelOriginalModelInfo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelOriginalModelInfo or *ModelOriginalModelInfo", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ModelOriginalModelInfo)
	if !ok {
		actualNotPointer, ok := a.(ModelOriginalModelInfo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelOriginalModelInfo", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Model, actual.Model, dcl.DiffInfo{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Model")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareModelContainerSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ModelContainerSpec)
	if !ok {
		desiredNotPointer, ok := d.(ModelContainerSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelContainerSpec or *ModelContainerSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ModelContainerSpec)
	if !ok {
		actualNotPointer, ok := a.(ModelContainerSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelContainerSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ImageUri, actual.ImageUri, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImageUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Env, actual.Env, dcl.DiffInfo{ObjectFunction: compareModelContainerSpecEnvNewStyle, EmptyObject: EmptyModelContainerSpecEnv, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Env")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.DiffInfo{ObjectFunction: compareModelContainerSpecPortsNewStyle, EmptyObject: EmptyModelContainerSpecPorts, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PredictRoute, actual.PredictRoute, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PredictRoute")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HealthRoute, actual.HealthRoute, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HealthRoute")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareModelContainerSpecEnvNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ModelContainerSpecEnv)
	if !ok {
		desiredNotPointer, ok := d.(ModelContainerSpecEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelContainerSpecEnv or *ModelContainerSpecEnv", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ModelContainerSpecEnv)
	if !ok {
		actualNotPointer, ok := a.(ModelContainerSpecEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelContainerSpecEnv", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareModelContainerSpecPortsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ModelContainerSpecPorts)
	if !ok {
		desiredNotPointer, ok := d.(ModelContainerSpecPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelContainerSpecPorts or *ModelContainerSpecPorts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ModelContainerSpecPorts)
	if !ok {
		actualNotPointer, ok := a.(ModelContainerSpecPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelContainerSpecPorts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ContainerPort, actual.ContainerPort, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ContainerPort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareModelDeployedModelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ModelDeployedModels)
	if !ok {
		desiredNotPointer, ok := d.(ModelDeployedModels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelDeployedModels or *ModelDeployedModels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ModelDeployedModels)
	if !ok {
		actualNotPointer, ok := a.(ModelDeployedModels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelDeployedModels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Endpoint, actual.Endpoint, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Endpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeployedModelId, actual.DeployedModelId, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeployedModelId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareModelEncryptionSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ModelEncryptionSpec)
	if !ok {
		desiredNotPointer, ok := d.(ModelEncryptionSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelEncryptionSpec or *ModelEncryptionSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ModelEncryptionSpec)
	if !ok {
		actualNotPointer, ok := a.(ModelEncryptionSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelEncryptionSpec", a)
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Model) urlNormalized() *Model {
	normalized := dcl.Copy(*r).(Model)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.VersionId = dcl.SelfLinkToName(r.VersionId)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.VersionDescription = dcl.SelfLinkToName(r.VersionDescription)
	normalized.TrainingPipeline = dcl.SelfLinkToName(r.TrainingPipeline)
	normalized.ArtifactUri = dcl.SelfLinkToName(r.ArtifactUri)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Model) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateModel" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/models/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Model resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Model) marshal(c *Client) ([]byte, error) {
	m, err := expandModel(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Model: %w", err)
	}
	m = encodeUploadModelRequest(m)

	return json.Marshal(m)
}

// unmarshalModel decodes JSON responses into the Model resource schema.
func unmarshalModel(b []byte, c *Client, res *Model) (*Model, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapModel(m, c, res)
}

func unmarshalMapModel(m map[string]interface{}, c *Client, res *Model) (*Model, error) {

	flattened := flattenModel(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandModel expands Model into a JSON request object.
func expandModel(c *Client, f *Model) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.VersionDescription; dcl.ValueShouldBeSent(v) {
		m["versionDescription"] = v
	}
	if v, err := expandModelContainerSpec(c, f.ContainerSpec, res); err != nil {
		return nil, fmt.Errorf("error expanding ContainerSpec into containerSpec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["containerSpec"] = v
	}
	if v := f.ArtifactUri; dcl.ValueShouldBeSent(v) {
		m["artifactUri"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandModelEncryptionSpec(c, f.EncryptionSpec, res); err != nil {
		return nil, fmt.Errorf("error expanding EncryptionSpec into encryptionSpec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["encryptionSpec"] = v
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

// flattenModel flattens Model from a JSON request object into the
// Model type.
func flattenModel(c *Client, i interface{}, res *Model) *Model {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Model{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.VersionId = dcl.FlattenString(m["versionId"])
	resultRes.VersionCreateTime = dcl.FlattenString(m["versionCreateTime"])
	resultRes.VersionUpdateTime = dcl.FlattenString(m["versionUpdateTime"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.VersionDescription = dcl.FlattenString(m["versionDescription"])
	resultRes.SupportedExportFormats = flattenModelSupportedExportFormatsSlice(c, m["supportedExportFormats"], res)
	resultRes.TrainingPipeline = dcl.FlattenString(m["trainingPipeline"])
	resultRes.OriginalModelInfo = flattenModelOriginalModelInfo(c, m["originalModelInfo"], res)
	resultRes.ContainerSpec = flattenModelContainerSpec(c, m["containerSpec"], res)
	resultRes.ArtifactUri = dcl.FlattenString(m["artifactUri"])
	resultRes.SupportedDeploymentResourcesTypes = flattenModelSupportedDeploymentResourcesTypesEnumSlice(c, m["supportedDeploymentResourcesTypes"], res)
	resultRes.SupportedInputStorageFormats = dcl.FlattenStringSlice(m["supportedInputStorageFormats"])
	resultRes.SupportedOutputStorageFormats = dcl.FlattenStringSlice(m["supportedOutputStorageFormats"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.DeployedModels = flattenModelDeployedModelsSlice(c, m["deployedModels"], res)
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.EncryptionSpec = flattenModelEncryptionSpec(c, m["encryptionSpec"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandModelSupportedExportFormatsMap expands the contents of ModelSupportedExportFormats into a JSON
// request object.
func expandModelSupportedExportFormatsMap(c *Client, f map[string]ModelSupportedExportFormats, res *Model) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandModelSupportedExportFormats(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandModelSupportedExportFormatsSlice expands the contents of ModelSupportedExportFormats into a JSON
// request object.
func expandModelSupportedExportFormatsSlice(c *Client, f []ModelSupportedExportFormats, res *Model) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandModelSupportedExportFormats(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenModelSupportedExportFormatsMap flattens the contents of ModelSupportedExportFormats from a JSON
// response object.
func flattenModelSupportedExportFormatsMap(c *Client, i interface{}, res *Model) map[string]ModelSupportedExportFormats {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelSupportedExportFormats{}
	}

	if len(a) == 0 {
		return map[string]ModelSupportedExportFormats{}
	}

	items := make(map[string]ModelSupportedExportFormats)
	for k, item := range a {
		items[k] = *flattenModelSupportedExportFormats(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenModelSupportedExportFormatsSlice flattens the contents of ModelSupportedExportFormats from a JSON
// response object.
func flattenModelSupportedExportFormatsSlice(c *Client, i interface{}, res *Model) []ModelSupportedExportFormats {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelSupportedExportFormats{}
	}

	if len(a) == 0 {
		return []ModelSupportedExportFormats{}
	}

	items := make([]ModelSupportedExportFormats, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelSupportedExportFormats(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandModelSupportedExportFormats expands an instance of ModelSupportedExportFormats into a JSON
// request object.
func expandModelSupportedExportFormats(c *Client, f *ModelSupportedExportFormats, res *Model) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenModelSupportedExportFormats flattens an instance of ModelSupportedExportFormats from a JSON
// response object.
func flattenModelSupportedExportFormats(c *Client, i interface{}, res *Model) *ModelSupportedExportFormats {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ModelSupportedExportFormats{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyModelSupportedExportFormats
	}
	r.Id = dcl.FlattenString(m["id"])
	r.ExportableContents = flattenModelSupportedExportFormatsExportableContentsEnumSlice(c, m["exportableContents"], res)

	return r
}

// expandModelOriginalModelInfoMap expands the contents of ModelOriginalModelInfo into a JSON
// request object.
func expandModelOriginalModelInfoMap(c *Client, f map[string]ModelOriginalModelInfo, res *Model) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandModelOriginalModelInfo(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandModelOriginalModelInfoSlice expands the contents of ModelOriginalModelInfo into a JSON
// request object.
func expandModelOriginalModelInfoSlice(c *Client, f []ModelOriginalModelInfo, res *Model) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandModelOriginalModelInfo(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenModelOriginalModelInfoMap flattens the contents of ModelOriginalModelInfo from a JSON
// response object.
func flattenModelOriginalModelInfoMap(c *Client, i interface{}, res *Model) map[string]ModelOriginalModelInfo {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelOriginalModelInfo{}
	}

	if len(a) == 0 {
		return map[string]ModelOriginalModelInfo{}
	}

	items := make(map[string]ModelOriginalModelInfo)
	for k, item := range a {
		items[k] = *flattenModelOriginalModelInfo(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenModelOriginalModelInfoSlice flattens the contents of ModelOriginalModelInfo from a JSON
// response object.
func flattenModelOriginalModelInfoSlice(c *Client, i interface{}, res *Model) []ModelOriginalModelInfo {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelOriginalModelInfo{}
	}

	if len(a) == 0 {
		return []ModelOriginalModelInfo{}
	}

	items := make([]ModelOriginalModelInfo, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelOriginalModelInfo(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandModelOriginalModelInfo expands an instance of ModelOriginalModelInfo into a JSON
// request object.
func expandModelOriginalModelInfo(c *Client, f *ModelOriginalModelInfo, res *Model) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenModelOriginalModelInfo flattens an instance of ModelOriginalModelInfo from a JSON
// response object.
func flattenModelOriginalModelInfo(c *Client, i interface{}, res *Model) *ModelOriginalModelInfo {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ModelOriginalModelInfo{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyModelOriginalModelInfo
	}
	r.Model = dcl.FlattenString(m["model"])

	return r
}

// expandModelContainerSpecMap expands the contents of ModelContainerSpec into a JSON
// request object.
func expandModelContainerSpecMap(c *Client, f map[string]ModelContainerSpec, res *Model) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandModelContainerSpec(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandModelContainerSpecSlice expands the contents of ModelContainerSpec into a JSON
// request object.
func expandModelContainerSpecSlice(c *Client, f []ModelContainerSpec, res *Model) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandModelContainerSpec(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenModelContainerSpecMap flattens the contents of ModelContainerSpec from a JSON
// response object.
func flattenModelContainerSpecMap(c *Client, i interface{}, res *Model) map[string]ModelContainerSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelContainerSpec{}
	}

	if len(a) == 0 {
		return map[string]ModelContainerSpec{}
	}

	items := make(map[string]ModelContainerSpec)
	for k, item := range a {
		items[k] = *flattenModelContainerSpec(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenModelContainerSpecSlice flattens the contents of ModelContainerSpec from a JSON
// response object.
func flattenModelContainerSpecSlice(c *Client, i interface{}, res *Model) []ModelContainerSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelContainerSpec{}
	}

	if len(a) == 0 {
		return []ModelContainerSpec{}
	}

	items := make([]ModelContainerSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelContainerSpec(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandModelContainerSpec expands an instance of ModelContainerSpec into a JSON
// request object.
func expandModelContainerSpec(c *Client, f *ModelContainerSpec, res *Model) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ImageUri; !dcl.IsEmptyValueIndirect(v) {
		m["imageUri"] = v
	}
	if v := f.Command; v != nil {
		m["command"] = v
	}
	if v := f.Args; v != nil {
		m["args"] = v
	}
	if v, err := expandModelContainerSpecEnvSlice(c, f.Env, res); err != nil {
		return nil, fmt.Errorf("error expanding Env into env: %w", err)
	} else if v != nil {
		m["env"] = v
	}
	if v, err := expandModelContainerSpecPortsSlice(c, f.Ports, res); err != nil {
		return nil, fmt.Errorf("error expanding Ports into ports: %w", err)
	} else if v != nil {
		m["ports"] = v
	}
	if v := f.PredictRoute; !dcl.IsEmptyValueIndirect(v) {
		m["predictRoute"] = v
	}
	if v := f.HealthRoute; !dcl.IsEmptyValueIndirect(v) {
		m["healthRoute"] = v
	}

	return m, nil
}

// flattenModelContainerSpec flattens an instance of ModelContainerSpec from a JSON
// response object.
func flattenModelContainerSpec(c *Client, i interface{}, res *Model) *ModelContainerSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ModelContainerSpec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyModelContainerSpec
	}
	r.ImageUri = dcl.FlattenString(m["imageUri"])
	r.Command = dcl.FlattenStringSlice(m["command"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.Env = flattenModelContainerSpecEnvSlice(c, m["env"], res)
	r.Ports = flattenModelContainerSpecPortsSlice(c, m["ports"], res)
	r.PredictRoute = dcl.FlattenString(m["predictRoute"])
	r.HealthRoute = dcl.FlattenString(m["healthRoute"])

	return r
}

// expandModelContainerSpecEnvMap expands the contents of ModelContainerSpecEnv into a JSON
// request object.
func expandModelContainerSpecEnvMap(c *Client, f map[string]ModelContainerSpecEnv, res *Model) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandModelContainerSpecEnv(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandModelContainerSpecEnvSlice expands the contents of ModelContainerSpecEnv into a JSON
// request object.
func expandModelContainerSpecEnvSlice(c *Client, f []ModelContainerSpecEnv, res *Model) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandModelContainerSpecEnv(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenModelContainerSpecEnvMap flattens the contents of ModelContainerSpecEnv from a JSON
// response object.
func flattenModelContainerSpecEnvMap(c *Client, i interface{}, res *Model) map[string]ModelContainerSpecEnv {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelContainerSpecEnv{}
	}

	if len(a) == 0 {
		return map[string]ModelContainerSpecEnv{}
	}

	items := make(map[string]ModelContainerSpecEnv)
	for k, item := range a {
		items[k] = *flattenModelContainerSpecEnv(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenModelContainerSpecEnvSlice flattens the contents of ModelContainerSpecEnv from a JSON
// response object.
func flattenModelContainerSpecEnvSlice(c *Client, i interface{}, res *Model) []ModelContainerSpecEnv {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelContainerSpecEnv{}
	}

	if len(a) == 0 {
		return []ModelContainerSpecEnv{}
	}

	items := make([]ModelContainerSpecEnv, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelContainerSpecEnv(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandModelContainerSpecEnv expands an instance of ModelContainerSpecEnv into a JSON
// request object.
func expandModelContainerSpecEnv(c *Client, f *ModelContainerSpecEnv, res *Model) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenModelContainerSpecEnv flattens an instance of ModelContainerSpecEnv from a JSON
// response object.
func flattenModelContainerSpecEnv(c *Client, i interface{}, res *Model) *ModelContainerSpecEnv {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ModelContainerSpecEnv{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyModelContainerSpecEnv
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandModelContainerSpecPortsMap expands the contents of ModelContainerSpecPorts into a JSON
// request object.
func expandModelContainerSpecPortsMap(c *Client, f map[string]ModelContainerSpecPorts, res *Model) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandModelContainerSpecPorts(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandModelContainerSpecPortsSlice expands the contents of ModelContainerSpecPorts into a JSON
// request object.
func expandModelContainerSpecPortsSlice(c *Client, f []ModelContainerSpecPorts, res *Model) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandModelContainerSpecPorts(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenModelContainerSpecPortsMap flattens the contents of ModelContainerSpecPorts from a JSON
// response object.
func flattenModelContainerSpecPortsMap(c *Client, i interface{}, res *Model) map[string]ModelContainerSpecPorts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelContainerSpecPorts{}
	}

	if len(a) == 0 {
		return map[string]ModelContainerSpecPorts{}
	}

	items := make(map[string]ModelContainerSpecPorts)
	for k, item := range a {
		items[k] = *flattenModelContainerSpecPorts(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenModelContainerSpecPortsSlice flattens the contents of ModelContainerSpecPorts from a JSON
// response object.
func flattenModelContainerSpecPortsSlice(c *Client, i interface{}, res *Model) []ModelContainerSpecPorts {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelContainerSpecPorts{}
	}

	if len(a) == 0 {
		return []ModelContainerSpecPorts{}
	}

	items := make([]ModelContainerSpecPorts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelContainerSpecPorts(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandModelContainerSpecPorts expands an instance of ModelContainerSpecPorts into a JSON
// request object.
func expandModelContainerSpecPorts(c *Client, f *ModelContainerSpecPorts, res *Model) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ContainerPort; !dcl.IsEmptyValueIndirect(v) {
		m["containerPort"] = v
	}

	return m, nil
}

// flattenModelContainerSpecPorts flattens an instance of ModelContainerSpecPorts from a JSON
// response object.
func flattenModelContainerSpecPorts(c *Client, i interface{}, res *Model) *ModelContainerSpecPorts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ModelContainerSpecPorts{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyModelContainerSpecPorts
	}
	r.ContainerPort = dcl.FlattenInteger(m["containerPort"])

	return r
}

// expandModelDeployedModelsMap expands the contents of ModelDeployedModels into a JSON
// request object.
func expandModelDeployedModelsMap(c *Client, f map[string]ModelDeployedModels, res *Model) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandModelDeployedModels(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandModelDeployedModelsSlice expands the contents of ModelDeployedModels into a JSON
// request object.
func expandModelDeployedModelsSlice(c *Client, f []ModelDeployedModels, res *Model) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandModelDeployedModels(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenModelDeployedModelsMap flattens the contents of ModelDeployedModels from a JSON
// response object.
func flattenModelDeployedModelsMap(c *Client, i interface{}, res *Model) map[string]ModelDeployedModels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelDeployedModels{}
	}

	if len(a) == 0 {
		return map[string]ModelDeployedModels{}
	}

	items := make(map[string]ModelDeployedModels)
	for k, item := range a {
		items[k] = *flattenModelDeployedModels(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenModelDeployedModelsSlice flattens the contents of ModelDeployedModels from a JSON
// response object.
func flattenModelDeployedModelsSlice(c *Client, i interface{}, res *Model) []ModelDeployedModels {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelDeployedModels{}
	}

	if len(a) == 0 {
		return []ModelDeployedModels{}
	}

	items := make([]ModelDeployedModels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelDeployedModels(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandModelDeployedModels expands an instance of ModelDeployedModels into a JSON
// request object.
func expandModelDeployedModels(c *Client, f *ModelDeployedModels, res *Model) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Endpoint; !dcl.IsEmptyValueIndirect(v) {
		m["endpoint"] = v
	}
	if v := f.DeployedModelId; !dcl.IsEmptyValueIndirect(v) {
		m["deployedModelId"] = v
	}

	return m, nil
}

// flattenModelDeployedModels flattens an instance of ModelDeployedModels from a JSON
// response object.
func flattenModelDeployedModels(c *Client, i interface{}, res *Model) *ModelDeployedModels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ModelDeployedModels{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyModelDeployedModels
	}
	r.Endpoint = dcl.FlattenString(m["endpoint"])
	r.DeployedModelId = dcl.FlattenString(m["deployedModelId"])

	return r
}

// expandModelEncryptionSpecMap expands the contents of ModelEncryptionSpec into a JSON
// request object.
func expandModelEncryptionSpecMap(c *Client, f map[string]ModelEncryptionSpec, res *Model) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandModelEncryptionSpec(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandModelEncryptionSpecSlice expands the contents of ModelEncryptionSpec into a JSON
// request object.
func expandModelEncryptionSpecSlice(c *Client, f []ModelEncryptionSpec, res *Model) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandModelEncryptionSpec(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenModelEncryptionSpecMap flattens the contents of ModelEncryptionSpec from a JSON
// response object.
func flattenModelEncryptionSpecMap(c *Client, i interface{}, res *Model) map[string]ModelEncryptionSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelEncryptionSpec{}
	}

	if len(a) == 0 {
		return map[string]ModelEncryptionSpec{}
	}

	items := make(map[string]ModelEncryptionSpec)
	for k, item := range a {
		items[k] = *flattenModelEncryptionSpec(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenModelEncryptionSpecSlice flattens the contents of ModelEncryptionSpec from a JSON
// response object.
func flattenModelEncryptionSpecSlice(c *Client, i interface{}, res *Model) []ModelEncryptionSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelEncryptionSpec{}
	}

	if len(a) == 0 {
		return []ModelEncryptionSpec{}
	}

	items := make([]ModelEncryptionSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelEncryptionSpec(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandModelEncryptionSpec expands an instance of ModelEncryptionSpec into a JSON
// request object.
func expandModelEncryptionSpec(c *Client, f *ModelEncryptionSpec, res *Model) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}

	return m, nil
}

// flattenModelEncryptionSpec flattens an instance of ModelEncryptionSpec from a JSON
// response object.
func flattenModelEncryptionSpec(c *Client, i interface{}, res *Model) *ModelEncryptionSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ModelEncryptionSpec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyModelEncryptionSpec
	}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])

	return r
}

// flattenModelSupportedExportFormatsExportableContentsEnumMap flattens the contents of ModelSupportedExportFormatsExportableContentsEnum from a JSON
// response object.
func flattenModelSupportedExportFormatsExportableContentsEnumMap(c *Client, i interface{}, res *Model) map[string]ModelSupportedExportFormatsExportableContentsEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelSupportedExportFormatsExportableContentsEnum{}
	}

	if len(a) == 0 {
		return map[string]ModelSupportedExportFormatsExportableContentsEnum{}
	}

	items := make(map[string]ModelSupportedExportFormatsExportableContentsEnum)
	for k, item := range a {
		items[k] = *flattenModelSupportedExportFormatsExportableContentsEnum(item.(interface{}))
	}

	return items
}

// flattenModelSupportedExportFormatsExportableContentsEnumSlice flattens the contents of ModelSupportedExportFormatsExportableContentsEnum from a JSON
// response object.
func flattenModelSupportedExportFormatsExportableContentsEnumSlice(c *Client, i interface{}, res *Model) []ModelSupportedExportFormatsExportableContentsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelSupportedExportFormatsExportableContentsEnum{}
	}

	if len(a) == 0 {
		return []ModelSupportedExportFormatsExportableContentsEnum{}
	}

	items := make([]ModelSupportedExportFormatsExportableContentsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelSupportedExportFormatsExportableContentsEnum(item.(interface{})))
	}

	return items
}

// flattenModelSupportedExportFormatsExportableContentsEnum asserts that an interface is a string, and returns a
// pointer to a *ModelSupportedExportFormatsExportableContentsEnum with the same value as that string.
func flattenModelSupportedExportFormatsExportableContentsEnum(i interface{}) *ModelSupportedExportFormatsExportableContentsEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ModelSupportedExportFormatsExportableContentsEnumRef(s)
}

// flattenModelSupportedDeploymentResourcesTypesEnumMap flattens the contents of ModelSupportedDeploymentResourcesTypesEnum from a JSON
// response object.
func flattenModelSupportedDeploymentResourcesTypesEnumMap(c *Client, i interface{}, res *Model) map[string]ModelSupportedDeploymentResourcesTypesEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelSupportedDeploymentResourcesTypesEnum{}
	}

	if len(a) == 0 {
		return map[string]ModelSupportedDeploymentResourcesTypesEnum{}
	}

	items := make(map[string]ModelSupportedDeploymentResourcesTypesEnum)
	for k, item := range a {
		items[k] = *flattenModelSupportedDeploymentResourcesTypesEnum(item.(interface{}))
	}

	return items
}

// flattenModelSupportedDeploymentResourcesTypesEnumSlice flattens the contents of ModelSupportedDeploymentResourcesTypesEnum from a JSON
// response object.
func flattenModelSupportedDeploymentResourcesTypesEnumSlice(c *Client, i interface{}, res *Model) []ModelSupportedDeploymentResourcesTypesEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelSupportedDeploymentResourcesTypesEnum{}
	}

	if len(a) == 0 {
		return []ModelSupportedDeploymentResourcesTypesEnum{}
	}

	items := make([]ModelSupportedDeploymentResourcesTypesEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelSupportedDeploymentResourcesTypesEnum(item.(interface{})))
	}

	return items
}

// flattenModelSupportedDeploymentResourcesTypesEnum asserts that an interface is a string, and returns a
// pointer to a *ModelSupportedDeploymentResourcesTypesEnum with the same value as that string.
func flattenModelSupportedDeploymentResourcesTypesEnum(i interface{}) *ModelSupportedDeploymentResourcesTypesEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ModelSupportedDeploymentResourcesTypesEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Model) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalModel(b, c, r)
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

type modelDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         modelApiOperation
}

func convertFieldDiffsToModelDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]modelDiff, error) {
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
	var diffs []modelDiff
	// For each operation name, create a modelDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := modelDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToModelApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToModelApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (modelApiOperation, error) {
	switch opName {

	case "updateModelUpdateModelOperation":
		return &updateModelUpdateModelOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractModelFields(r *Model) error {
	vOriginalModelInfo := r.OriginalModelInfo
	if vOriginalModelInfo == nil {
		// note: explicitly not the empty object.
		vOriginalModelInfo = &ModelOriginalModelInfo{}
	}
	if err := extractModelOriginalModelInfoFields(r, vOriginalModelInfo); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOriginalModelInfo) {
		r.OriginalModelInfo = vOriginalModelInfo
	}
	vContainerSpec := r.ContainerSpec
	if vContainerSpec == nil {
		// note: explicitly not the empty object.
		vContainerSpec = &ModelContainerSpec{}
	}
	if err := extractModelContainerSpecFields(r, vContainerSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vContainerSpec) {
		r.ContainerSpec = vContainerSpec
	}
	vEncryptionSpec := r.EncryptionSpec
	if vEncryptionSpec == nil {
		// note: explicitly not the empty object.
		vEncryptionSpec = &ModelEncryptionSpec{}
	}
	if err := extractModelEncryptionSpecFields(r, vEncryptionSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEncryptionSpec) {
		r.EncryptionSpec = vEncryptionSpec
	}
	return nil
}
func extractModelSupportedExportFormatsFields(r *Model, o *ModelSupportedExportFormats) error {
	return nil
}
func extractModelOriginalModelInfoFields(r *Model, o *ModelOriginalModelInfo) error {
	return nil
}
func extractModelContainerSpecFields(r *Model, o *ModelContainerSpec) error {
	return nil
}
func extractModelContainerSpecEnvFields(r *Model, o *ModelContainerSpecEnv) error {
	return nil
}
func extractModelContainerSpecPortsFields(r *Model, o *ModelContainerSpecPorts) error {
	return nil
}
func extractModelDeployedModelsFields(r *Model, o *ModelDeployedModels) error {
	return nil
}
func extractModelEncryptionSpecFields(r *Model, o *ModelEncryptionSpec) error {
	return nil
}

func postReadExtractModelFields(r *Model) error {
	vOriginalModelInfo := r.OriginalModelInfo
	if vOriginalModelInfo == nil {
		// note: explicitly not the empty object.
		vOriginalModelInfo = &ModelOriginalModelInfo{}
	}
	if err := postReadExtractModelOriginalModelInfoFields(r, vOriginalModelInfo); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOriginalModelInfo) {
		r.OriginalModelInfo = vOriginalModelInfo
	}
	vContainerSpec := r.ContainerSpec
	if vContainerSpec == nil {
		// note: explicitly not the empty object.
		vContainerSpec = &ModelContainerSpec{}
	}
	if err := postReadExtractModelContainerSpecFields(r, vContainerSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vContainerSpec) {
		r.ContainerSpec = vContainerSpec
	}
	vEncryptionSpec := r.EncryptionSpec
	if vEncryptionSpec == nil {
		// note: explicitly not the empty object.
		vEncryptionSpec = &ModelEncryptionSpec{}
	}
	if err := postReadExtractModelEncryptionSpecFields(r, vEncryptionSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEncryptionSpec) {
		r.EncryptionSpec = vEncryptionSpec
	}
	return nil
}
func postReadExtractModelSupportedExportFormatsFields(r *Model, o *ModelSupportedExportFormats) error {
	return nil
}
func postReadExtractModelOriginalModelInfoFields(r *Model, o *ModelOriginalModelInfo) error {
	return nil
}
func postReadExtractModelContainerSpecFields(r *Model, o *ModelContainerSpec) error {
	return nil
}
func postReadExtractModelContainerSpecEnvFields(r *Model, o *ModelContainerSpecEnv) error {
	return nil
}
func postReadExtractModelContainerSpecPortsFields(r *Model, o *ModelContainerSpecPorts) error {
	return nil
}
func postReadExtractModelDeployedModelsFields(r *Model, o *ModelDeployedModels) error {
	return nil
}
func postReadExtractModelEncryptionSpecFields(r *Model, o *ModelEncryptionSpec) error {
	return nil
}
