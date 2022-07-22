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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Model struct {
	Name                              *string                                      `json:"name"`
	VersionId                         *string                                      `json:"versionId"`
	VersionCreateTime                 *string                                      `json:"versionCreateTime"`
	VersionUpdateTime                 *string                                      `json:"versionUpdateTime"`
	DisplayName                       *string                                      `json:"displayName"`
	Description                       *string                                      `json:"description"`
	VersionDescription                *string                                      `json:"versionDescription"`
	SupportedExportFormats            []ModelSupportedExportFormats                `json:"supportedExportFormats"`
	TrainingPipeline                  *string                                      `json:"trainingPipeline"`
	OriginalModelInfo                 *ModelOriginalModelInfo                      `json:"originalModelInfo"`
	ContainerSpec                     *ModelContainerSpec                          `json:"containerSpec"`
	ArtifactUri                       *string                                      `json:"artifactUri"`
	SupportedDeploymentResourcesTypes []ModelSupportedDeploymentResourcesTypesEnum `json:"supportedDeploymentResourcesTypes"`
	SupportedInputStorageFormats      []string                                     `json:"supportedInputStorageFormats"`
	SupportedOutputStorageFormats     []string                                     `json:"supportedOutputStorageFormats"`
	CreateTime                        *string                                      `json:"createTime"`
	UpdateTime                        *string                                      `json:"updateTime"`
	DeployedModels                    []ModelDeployedModels                        `json:"deployedModels"`
	Etag                              *string                                      `json:"etag"`
	Labels                            map[string]string                            `json:"labels"`
	EncryptionSpec                    *ModelEncryptionSpec                         `json:"encryptionSpec"`
	Project                           *string                                      `json:"project"`
	Location                          *string                                      `json:"location"`
}

func (r *Model) String() string {
	return dcl.SprintResource(r)
}

// The enum ModelSupportedExportFormatsExportableContentsEnum.
type ModelSupportedExportFormatsExportableContentsEnum string

// ModelSupportedExportFormatsExportableContentsEnumRef returns a *ModelSupportedExportFormatsExportableContentsEnum with the value of string s
// If the empty string is provided, nil is returned.
func ModelSupportedExportFormatsExportableContentsEnumRef(s string) *ModelSupportedExportFormatsExportableContentsEnum {
	v := ModelSupportedExportFormatsExportableContentsEnum(s)
	return &v
}

func (v ModelSupportedExportFormatsExportableContentsEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EXPORTABLE_CONTENT_UNSPECIFIED", "ARTIFACT", "IMAGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ModelSupportedExportFormatsExportableContentsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ModelSupportedDeploymentResourcesTypesEnum.
type ModelSupportedDeploymentResourcesTypesEnum string

// ModelSupportedDeploymentResourcesTypesEnumRef returns a *ModelSupportedDeploymentResourcesTypesEnum with the value of string s
// If the empty string is provided, nil is returned.
func ModelSupportedDeploymentResourcesTypesEnumRef(s string) *ModelSupportedDeploymentResourcesTypesEnum {
	v := ModelSupportedDeploymentResourcesTypesEnum(s)
	return &v
}

func (v ModelSupportedDeploymentResourcesTypesEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"DEPLOYMENT_RESOURCES_TYPE_UNSPECIFIED", "DEDICATED_RESOURCES", "AUTOMATIC_RESOURCES"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ModelSupportedDeploymentResourcesTypesEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ModelSupportedExportFormats struct {
	empty              bool                                                `json:"-"`
	Id                 *string                                             `json:"id"`
	ExportableContents []ModelSupportedExportFormatsExportableContentsEnum `json:"exportableContents"`
}

type jsonModelSupportedExportFormats ModelSupportedExportFormats

func (r *ModelSupportedExportFormats) UnmarshalJSON(data []byte) error {
	var res jsonModelSupportedExportFormats
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyModelSupportedExportFormats
	} else {

		r.Id = res.Id

		r.ExportableContents = res.ExportableContents

	}
	return nil
}

// This object is used to assert a desired state where this ModelSupportedExportFormats is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyModelSupportedExportFormats *ModelSupportedExportFormats = &ModelSupportedExportFormats{empty: true}

func (r *ModelSupportedExportFormats) Empty() bool {
	return r.empty
}

func (r *ModelSupportedExportFormats) String() string {
	return dcl.SprintResource(r)
}

func (r *ModelSupportedExportFormats) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ModelOriginalModelInfo struct {
	empty bool    `json:"-"`
	Model *string `json:"model"`
}

type jsonModelOriginalModelInfo ModelOriginalModelInfo

func (r *ModelOriginalModelInfo) UnmarshalJSON(data []byte) error {
	var res jsonModelOriginalModelInfo
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyModelOriginalModelInfo
	} else {

		r.Model = res.Model

	}
	return nil
}

// This object is used to assert a desired state where this ModelOriginalModelInfo is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyModelOriginalModelInfo *ModelOriginalModelInfo = &ModelOriginalModelInfo{empty: true}

func (r *ModelOriginalModelInfo) Empty() bool {
	return r.empty
}

func (r *ModelOriginalModelInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *ModelOriginalModelInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ModelContainerSpec struct {
	empty        bool                      `json:"-"`
	ImageUri     *string                   `json:"imageUri"`
	Command      []string                  `json:"command"`
	Args         []string                  `json:"args"`
	Env          []ModelContainerSpecEnv   `json:"env"`
	Ports        []ModelContainerSpecPorts `json:"ports"`
	PredictRoute *string                   `json:"predictRoute"`
	HealthRoute  *string                   `json:"healthRoute"`
}

type jsonModelContainerSpec ModelContainerSpec

func (r *ModelContainerSpec) UnmarshalJSON(data []byte) error {
	var res jsonModelContainerSpec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyModelContainerSpec
	} else {

		r.ImageUri = res.ImageUri

		r.Command = res.Command

		r.Args = res.Args

		r.Env = res.Env

		r.Ports = res.Ports

		r.PredictRoute = res.PredictRoute

		r.HealthRoute = res.HealthRoute

	}
	return nil
}

// This object is used to assert a desired state where this ModelContainerSpec is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyModelContainerSpec *ModelContainerSpec = &ModelContainerSpec{empty: true}

func (r *ModelContainerSpec) Empty() bool {
	return r.empty
}

func (r *ModelContainerSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *ModelContainerSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ModelContainerSpecEnv struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

type jsonModelContainerSpecEnv ModelContainerSpecEnv

func (r *ModelContainerSpecEnv) UnmarshalJSON(data []byte) error {
	var res jsonModelContainerSpecEnv
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyModelContainerSpecEnv
	} else {

		r.Name = res.Name

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this ModelContainerSpecEnv is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyModelContainerSpecEnv *ModelContainerSpecEnv = &ModelContainerSpecEnv{empty: true}

func (r *ModelContainerSpecEnv) Empty() bool {
	return r.empty
}

func (r *ModelContainerSpecEnv) String() string {
	return dcl.SprintResource(r)
}

func (r *ModelContainerSpecEnv) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ModelContainerSpecPorts struct {
	empty         bool   `json:"-"`
	ContainerPort *int64 `json:"containerPort"`
}

type jsonModelContainerSpecPorts ModelContainerSpecPorts

func (r *ModelContainerSpecPorts) UnmarshalJSON(data []byte) error {
	var res jsonModelContainerSpecPorts
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyModelContainerSpecPorts
	} else {

		r.ContainerPort = res.ContainerPort

	}
	return nil
}

// This object is used to assert a desired state where this ModelContainerSpecPorts is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyModelContainerSpecPorts *ModelContainerSpecPorts = &ModelContainerSpecPorts{empty: true}

func (r *ModelContainerSpecPorts) Empty() bool {
	return r.empty
}

func (r *ModelContainerSpecPorts) String() string {
	return dcl.SprintResource(r)
}

func (r *ModelContainerSpecPorts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ModelDeployedModels struct {
	empty           bool    `json:"-"`
	Endpoint        *string `json:"endpoint"`
	DeployedModelId *string `json:"deployedModelId"`
}

type jsonModelDeployedModels ModelDeployedModels

func (r *ModelDeployedModels) UnmarshalJSON(data []byte) error {
	var res jsonModelDeployedModels
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyModelDeployedModels
	} else {

		r.Endpoint = res.Endpoint

		r.DeployedModelId = res.DeployedModelId

	}
	return nil
}

// This object is used to assert a desired state where this ModelDeployedModels is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyModelDeployedModels *ModelDeployedModels = &ModelDeployedModels{empty: true}

func (r *ModelDeployedModels) Empty() bool {
	return r.empty
}

func (r *ModelDeployedModels) String() string {
	return dcl.SprintResource(r)
}

func (r *ModelDeployedModels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ModelEncryptionSpec struct {
	empty      bool    `json:"-"`
	KmsKeyName *string `json:"kmsKeyName"`
}

type jsonModelEncryptionSpec ModelEncryptionSpec

func (r *ModelEncryptionSpec) UnmarshalJSON(data []byte) error {
	var res jsonModelEncryptionSpec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyModelEncryptionSpec
	} else {

		r.KmsKeyName = res.KmsKeyName

	}
	return nil
}

// This object is used to assert a desired state where this ModelEncryptionSpec is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyModelEncryptionSpec *ModelEncryptionSpec = &ModelEncryptionSpec{empty: true}

func (r *ModelEncryptionSpec) Empty() bool {
	return r.empty
}

func (r *ModelEncryptionSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *ModelEncryptionSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Model) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "vertex_ai",
		Type:    "Model",
		Version: "vertexai",
	}
}

func (r *Model) ID() (string, error) {
	if err := extractModelFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                                 dcl.ValueOrEmptyString(nr.Name),
		"version_id":                           dcl.ValueOrEmptyString(nr.VersionId),
		"version_create_time":                  dcl.ValueOrEmptyString(nr.VersionCreateTime),
		"version_update_time":                  dcl.ValueOrEmptyString(nr.VersionUpdateTime),
		"display_name":                         dcl.ValueOrEmptyString(nr.DisplayName),
		"description":                          dcl.ValueOrEmptyString(nr.Description),
		"version_description":                  dcl.ValueOrEmptyString(nr.VersionDescription),
		"supported_export_formats":             dcl.ValueOrEmptyString(nr.SupportedExportFormats),
		"training_pipeline":                    dcl.ValueOrEmptyString(nr.TrainingPipeline),
		"original_model_info":                  dcl.ValueOrEmptyString(nr.OriginalModelInfo),
		"container_spec":                       dcl.ValueOrEmptyString(nr.ContainerSpec),
		"artifact_uri":                         dcl.ValueOrEmptyString(nr.ArtifactUri),
		"supported_deployment_resources_types": dcl.ValueOrEmptyString(nr.SupportedDeploymentResourcesTypes),
		"supported_input_storage_formats":      dcl.ValueOrEmptyString(nr.SupportedInputStorageFormats),
		"supported_output_storage_formats":     dcl.ValueOrEmptyString(nr.SupportedOutputStorageFormats),
		"create_time":                          dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":                          dcl.ValueOrEmptyString(nr.UpdateTime),
		"deployed_models":                      dcl.ValueOrEmptyString(nr.DeployedModels),
		"etag":                                 dcl.ValueOrEmptyString(nr.Etag),
		"labels":                               dcl.ValueOrEmptyString(nr.Labels),
		"encryption_spec":                      dcl.ValueOrEmptyString(nr.EncryptionSpec),
		"project":                              dcl.ValueOrEmptyString(nr.Project),
		"location":                             dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/models/{{name}}", params), nil
}

const ModelMaxPage = -1

type ModelList struct {
	Items []*Model

	nextToken string

	pageSize int32

	resource *Model
}

func (l *ModelList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ModelList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listModel(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListModel(ctx context.Context, project, location string) (*ModelList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListModelWithMaxResults(ctx, project, location, ModelMaxPage)

}

func (c *Client) ListModelWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ModelList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Model{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listModel(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ModelList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetModel(ctx context.Context, r *Model) (*Model, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractModelFields(r)

	b, err := c.getModelRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalModel(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeModelNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractModelFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteModel(ctx context.Context, r *Model) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Model resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Model...")
	deleteOp := deleteModelOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllModel deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllModel(ctx context.Context, project, location string, filter func(*Model) bool) error {
	listObj, err := c.ListModel(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllModel(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllModel(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyModel(ctx context.Context, rawDesired *Model, opts ...dcl.ApplyOption) (*Model, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Model
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyModelHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyModelHelper(c *Client, ctx context.Context, rawDesired *Model, opts ...dcl.ApplyOption) (*Model, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyModel...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractModelFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.modelDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToModelDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []modelApiOperation
	if create {
		ops = append(ops, &createModelOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyModelDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyModelDiff(c *Client, ctx context.Context, desired *Model, rawDesired *Model, ops []modelApiOperation, opts ...dcl.ApplyOption) (*Model, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetModel(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createModelOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapModel(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeModelNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeModelNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeModelDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractModelFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractModelFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffModel(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
