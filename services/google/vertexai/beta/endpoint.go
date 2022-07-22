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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Endpoint struct {
	Name                         *string                  `json:"name"`
	DisplayName                  *string                  `json:"displayName"`
	Description                  *string                  `json:"description"`
	DeployedModels               []EndpointDeployedModels `json:"deployedModels"`
	Etag                         *string                  `json:"etag"`
	Labels                       map[string]string        `json:"labels"`
	CreateTime                   *string                  `json:"createTime"`
	UpdateTime                   *string                  `json:"updateTime"`
	EncryptionSpec               *EndpointEncryptionSpec  `json:"encryptionSpec"`
	Network                      *string                  `json:"network"`
	ModelDeploymentMonitoringJob *string                  `json:"modelDeploymentMonitoringJob"`
	Project                      *string                  `json:"project"`
	Location                     *string                  `json:"location"`
}

func (r *Endpoint) String() string {
	return dcl.SprintResource(r)
}

// The enum EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum.
type EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum string

// EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumRef returns a *EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumRef(s string) *EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	v := EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(s)
	return &v
}

func (v EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ACCELERATOR_TYPE_UNSPECIFIED", "NVIDIA_TESLA_K80", "NVIDIA_TESLA_P100", "NVIDIA_TESLA_V100", "NVIDIA_TESLA_P4", "NVIDIA_TESLA_T4", "NVIDIA_TESLA_A100", "TPU_V2", "TPU_V3"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type EndpointDeployedModels struct {
	empty                   bool                                      `json:"-"`
	DedicatedResources      *EndpointDeployedModelsDedicatedResources `json:"dedicatedResources"`
	AutomaticResources      *EndpointDeployedModelsAutomaticResources `json:"automaticResources"`
	Id                      *string                                   `json:"id"`
	Model                   *string                                   `json:"model"`
	ModelVersionId          *string                                   `json:"modelVersionId"`
	DisplayName             *string                                   `json:"displayName"`
	CreateTime              *string                                   `json:"createTime"`
	ServiceAccount          *string                                   `json:"serviceAccount"`
	DisableContainerLogging *bool                                     `json:"disableContainerLogging"`
	EnableAccessLogging     *bool                                     `json:"enableAccessLogging"`
	PrivateEndpoints        *EndpointDeployedModelsPrivateEndpoints   `json:"privateEndpoints"`
	SharedResources         *string                                   `json:"sharedResources"`
	EnableContainerLogging  *bool                                     `json:"enableContainerLogging"`
}

type jsonEndpointDeployedModels EndpointDeployedModels

func (r *EndpointDeployedModels) UnmarshalJSON(data []byte) error {
	var res jsonEndpointDeployedModels
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointDeployedModels
	} else {

		r.DedicatedResources = res.DedicatedResources

		r.AutomaticResources = res.AutomaticResources

		r.Id = res.Id

		r.Model = res.Model

		r.ModelVersionId = res.ModelVersionId

		r.DisplayName = res.DisplayName

		r.CreateTime = res.CreateTime

		r.ServiceAccount = res.ServiceAccount

		r.DisableContainerLogging = res.DisableContainerLogging

		r.EnableAccessLogging = res.EnableAccessLogging

		r.PrivateEndpoints = res.PrivateEndpoints

		r.SharedResources = res.SharedResources

		r.EnableContainerLogging = res.EnableContainerLogging

	}
	return nil
}

// This object is used to assert a desired state where this EndpointDeployedModels is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointDeployedModels *EndpointDeployedModels = &EndpointDeployedModels{empty: true}

func (r *EndpointDeployedModels) Empty() bool {
	return r.empty
}

func (r *EndpointDeployedModels) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointDeployedModels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointDeployedModelsDedicatedResources struct {
	empty                  bool                                                             `json:"-"`
	MachineSpec            *EndpointDeployedModelsDedicatedResourcesMachineSpec             `json:"machineSpec"`
	MinReplicaCount        *int64                                                           `json:"minReplicaCount"`
	MaxReplicaCount        *int64                                                           `json:"maxReplicaCount"`
	AutoscalingMetricSpecs []EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs `json:"autoscalingMetricSpecs"`
}

type jsonEndpointDeployedModelsDedicatedResources EndpointDeployedModelsDedicatedResources

func (r *EndpointDeployedModelsDedicatedResources) UnmarshalJSON(data []byte) error {
	var res jsonEndpointDeployedModelsDedicatedResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointDeployedModelsDedicatedResources
	} else {

		r.MachineSpec = res.MachineSpec

		r.MinReplicaCount = res.MinReplicaCount

		r.MaxReplicaCount = res.MaxReplicaCount

		r.AutoscalingMetricSpecs = res.AutoscalingMetricSpecs

	}
	return nil
}

// This object is used to assert a desired state where this EndpointDeployedModelsDedicatedResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointDeployedModelsDedicatedResources *EndpointDeployedModelsDedicatedResources = &EndpointDeployedModelsDedicatedResources{empty: true}

func (r *EndpointDeployedModelsDedicatedResources) Empty() bool {
	return r.empty
}

func (r *EndpointDeployedModelsDedicatedResources) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointDeployedModelsDedicatedResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointDeployedModelsDedicatedResourcesMachineSpec struct {
	empty            bool                                                                    `json:"-"`
	MachineType      *string                                                                 `json:"machineType"`
	AcceleratorType  *EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum `json:"acceleratorType"`
	AcceleratorCount *int64                                                                  `json:"acceleratorCount"`
}

type jsonEndpointDeployedModelsDedicatedResourcesMachineSpec EndpointDeployedModelsDedicatedResourcesMachineSpec

func (r *EndpointDeployedModelsDedicatedResourcesMachineSpec) UnmarshalJSON(data []byte) error {
	var res jsonEndpointDeployedModelsDedicatedResourcesMachineSpec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointDeployedModelsDedicatedResourcesMachineSpec
	} else {

		r.MachineType = res.MachineType

		r.AcceleratorType = res.AcceleratorType

		r.AcceleratorCount = res.AcceleratorCount

	}
	return nil
}

// This object is used to assert a desired state where this EndpointDeployedModelsDedicatedResourcesMachineSpec is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointDeployedModelsDedicatedResourcesMachineSpec *EndpointDeployedModelsDedicatedResourcesMachineSpec = &EndpointDeployedModelsDedicatedResourcesMachineSpec{empty: true}

func (r *EndpointDeployedModelsDedicatedResourcesMachineSpec) Empty() bool {
	return r.empty
}

func (r *EndpointDeployedModelsDedicatedResourcesMachineSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointDeployedModelsDedicatedResourcesMachineSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Target     *int64  `json:"target"`
}

type jsonEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs

func (r *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) UnmarshalJSON(data []byte) error {
	var res jsonEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs
	} else {

		r.MetricName = res.MetricName

		r.Target = res.Target

	}
	return nil
}

// This object is used to assert a desired state where this EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs = &EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{empty: true}

func (r *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) Empty() bool {
	return r.empty
}

func (r *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointDeployedModelsAutomaticResources struct {
	empty           bool   `json:"-"`
	MinReplicaCount *int64 `json:"minReplicaCount"`
	MaxReplicaCount *int64 `json:"maxReplicaCount"`
}

type jsonEndpointDeployedModelsAutomaticResources EndpointDeployedModelsAutomaticResources

func (r *EndpointDeployedModelsAutomaticResources) UnmarshalJSON(data []byte) error {
	var res jsonEndpointDeployedModelsAutomaticResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointDeployedModelsAutomaticResources
	} else {

		r.MinReplicaCount = res.MinReplicaCount

		r.MaxReplicaCount = res.MaxReplicaCount

	}
	return nil
}

// This object is used to assert a desired state where this EndpointDeployedModelsAutomaticResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointDeployedModelsAutomaticResources *EndpointDeployedModelsAutomaticResources = &EndpointDeployedModelsAutomaticResources{empty: true}

func (r *EndpointDeployedModelsAutomaticResources) Empty() bool {
	return r.empty
}

func (r *EndpointDeployedModelsAutomaticResources) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointDeployedModelsAutomaticResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointDeployedModelsPrivateEndpoints struct {
	empty             bool    `json:"-"`
	PredictHttpUri    *string `json:"predictHttpUri"`
	ExplainHttpUri    *string `json:"explainHttpUri"`
	HealthHttpUri     *string `json:"healthHttpUri"`
	ServiceAttachment *string `json:"serviceAttachment"`
}

type jsonEndpointDeployedModelsPrivateEndpoints EndpointDeployedModelsPrivateEndpoints

func (r *EndpointDeployedModelsPrivateEndpoints) UnmarshalJSON(data []byte) error {
	var res jsonEndpointDeployedModelsPrivateEndpoints
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointDeployedModelsPrivateEndpoints
	} else {

		r.PredictHttpUri = res.PredictHttpUri

		r.ExplainHttpUri = res.ExplainHttpUri

		r.HealthHttpUri = res.HealthHttpUri

		r.ServiceAttachment = res.ServiceAttachment

	}
	return nil
}

// This object is used to assert a desired state where this EndpointDeployedModelsPrivateEndpoints is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointDeployedModelsPrivateEndpoints *EndpointDeployedModelsPrivateEndpoints = &EndpointDeployedModelsPrivateEndpoints{empty: true}

func (r *EndpointDeployedModelsPrivateEndpoints) Empty() bool {
	return r.empty
}

func (r *EndpointDeployedModelsPrivateEndpoints) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointDeployedModelsPrivateEndpoints) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointEncryptionSpec struct {
	empty      bool    `json:"-"`
	KmsKeyName *string `json:"kmsKeyName"`
}

type jsonEndpointEncryptionSpec EndpointEncryptionSpec

func (r *EndpointEncryptionSpec) UnmarshalJSON(data []byte) error {
	var res jsonEndpointEncryptionSpec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointEncryptionSpec
	} else {

		r.KmsKeyName = res.KmsKeyName

	}
	return nil
}

// This object is used to assert a desired state where this EndpointEncryptionSpec is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointEncryptionSpec *EndpointEncryptionSpec = &EndpointEncryptionSpec{empty: true}

func (r *EndpointEncryptionSpec) Empty() bool {
	return r.empty
}

func (r *EndpointEncryptionSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointEncryptionSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Endpoint) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "vertex_ai",
		Type:    "Endpoint",
		Version: "beta",
	}
}

func (r *Endpoint) ID() (string, error) {
	if err := extractEndpointFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                            dcl.ValueOrEmptyString(nr.Name),
		"display_name":                    dcl.ValueOrEmptyString(nr.DisplayName),
		"description":                     dcl.ValueOrEmptyString(nr.Description),
		"deployed_models":                 dcl.ValueOrEmptyString(nr.DeployedModels),
		"etag":                            dcl.ValueOrEmptyString(nr.Etag),
		"labels":                          dcl.ValueOrEmptyString(nr.Labels),
		"create_time":                     dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":                     dcl.ValueOrEmptyString(nr.UpdateTime),
		"encryption_spec":                 dcl.ValueOrEmptyString(nr.EncryptionSpec),
		"network":                         dcl.ValueOrEmptyString(nr.Network),
		"model_deployment_monitoring_job": dcl.ValueOrEmptyString(nr.ModelDeploymentMonitoringJob),
		"project":                         dcl.ValueOrEmptyString(nr.Project),
		"location":                        dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/endpoints/{{name}}", params), nil
}

const EndpointMaxPage = -1

type EndpointList struct {
	Items []*Endpoint

	nextToken string

	pageSize int32

	resource *Endpoint
}

func (l *EndpointList) HasNext() bool {
	return l.nextToken != ""
}

func (l *EndpointList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listEndpoint(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListEndpoint(ctx context.Context, project, location string) (*EndpointList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListEndpointWithMaxResults(ctx, project, location, EndpointMaxPage)

}

func (c *Client) ListEndpointWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*EndpointList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Endpoint{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listEndpoint(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &EndpointList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetEndpoint(ctx context.Context, r *Endpoint) (*Endpoint, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractEndpointFields(r)

	b, err := c.getEndpointRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEndpoint(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEndpointNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractEndpointFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteEndpoint(ctx context.Context, r *Endpoint) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Endpoint resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Endpoint...")
	deleteOp := deleteEndpointOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllEndpoint deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllEndpoint(ctx context.Context, project, location string, filter func(*Endpoint) bool) error {
	listObj, err := c.ListEndpoint(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllEndpoint(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllEndpoint(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyEndpoint(ctx context.Context, rawDesired *Endpoint, opts ...dcl.ApplyOption) (*Endpoint, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Endpoint
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyEndpointHelper(c, ctx, rawDesired, opts...)
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

func applyEndpointHelper(c *Client, ctx context.Context, rawDesired *Endpoint, opts ...dcl.ApplyOption) (*Endpoint, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyEndpoint...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractEndpointFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.endpointDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToEndpointDiffs(c.Config, fieldDiffs, opts)
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
	var ops []endpointApiOperation
	if create {
		ops = append(ops, &createEndpointOperation{})
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
	return applyEndpointDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyEndpointDiff(c *Client, ctx context.Context, desired *Endpoint, rawDesired *Endpoint, ops []endpointApiOperation, opts ...dcl.ApplyOption) (*Endpoint, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetEndpoint(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createEndpointOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapEndpoint(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeEndpointNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEndpointNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEndpointDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractEndpointFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractEndpointFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEndpoint(c, newDesired, newState)
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
