// Copyright 2025 Google LLC. All Rights Reserved.
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
package clouddeploy

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Target struct {
	Name               *string                             `json:"name"`
	TargetId           *string                             `json:"targetId"`
	Uid                *string                             `json:"uid"`
	Description        *string                             `json:"description"`
	Annotations        map[string]string                   `json:"annotations"`
	Labels             map[string]string                   `json:"labels"`
	RequireApproval    *bool                               `json:"requireApproval"`
	CreateTime         *string                             `json:"createTime"`
	UpdateTime         *string                             `json:"updateTime"`
	Gke                *TargetGke                          `json:"gke"`
	AnthosCluster      *TargetAnthosCluster                `json:"anthosCluster"`
	Etag               *string                             `json:"etag"`
	ExecutionConfigs   []TargetExecutionConfigs            `json:"executionConfigs"`
	Project            *string                             `json:"project"`
	Location           *string                             `json:"location"`
	Run                *TargetRun                          `json:"run"`
	MultiTarget        *TargetMultiTarget                  `json:"multiTarget"`
	DeployParameters   map[string]string                   `json:"deployParameters"`
	CustomTarget       *TargetCustomTarget                 `json:"customTarget"`
	AssociatedEntities map[string]TargetAssociatedEntities `json:"associatedEntities"`
}

func (r *Target) String() string {
	return dcl.SprintResource(r)
}

// The enum TargetExecutionConfigsUsagesEnum.
type TargetExecutionConfigsUsagesEnum string

// TargetExecutionConfigsUsagesEnumRef returns a *TargetExecutionConfigsUsagesEnum with the value of string s
// If the empty string is provided, nil is returned.
func TargetExecutionConfigsUsagesEnumRef(s string) *TargetExecutionConfigsUsagesEnum {
	v := TargetExecutionConfigsUsagesEnum(s)
	return &v
}

func (v TargetExecutionConfigsUsagesEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EXECUTION_ENVIRONMENT_USAGE_UNSPECIFIED", "RENDER", "DEPLOY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TargetExecutionConfigsUsagesEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type TargetGke struct {
	empty       bool    `json:"-"`
	Cluster     *string `json:"cluster"`
	InternalIP  *bool   `json:"internalIP"`
	ProxyUrl    *string `json:"proxyUrl"`
	DnsEndpoint *bool   `json:"dnsEndpoint"`
}

type jsonTargetGke TargetGke

func (r *TargetGke) UnmarshalJSON(data []byte) error {
	var res jsonTargetGke
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTargetGke
	} else {

		r.Cluster = res.Cluster

		r.InternalIP = res.InternalIP

		r.ProxyUrl = res.ProxyUrl

		r.DnsEndpoint = res.DnsEndpoint

	}
	return nil
}

// This object is used to assert a desired state where this TargetGke is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTargetGke *TargetGke = &TargetGke{empty: true}

func (r *TargetGke) Empty() bool {
	return r.empty
}

func (r *TargetGke) String() string {
	return dcl.SprintResource(r)
}

func (r *TargetGke) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TargetAnthosCluster struct {
	empty      bool    `json:"-"`
	Membership *string `json:"membership"`
}

type jsonTargetAnthosCluster TargetAnthosCluster

func (r *TargetAnthosCluster) UnmarshalJSON(data []byte) error {
	var res jsonTargetAnthosCluster
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTargetAnthosCluster
	} else {

		r.Membership = res.Membership

	}
	return nil
}

// This object is used to assert a desired state where this TargetAnthosCluster is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTargetAnthosCluster *TargetAnthosCluster = &TargetAnthosCluster{empty: true}

func (r *TargetAnthosCluster) Empty() bool {
	return r.empty
}

func (r *TargetAnthosCluster) String() string {
	return dcl.SprintResource(r)
}

func (r *TargetAnthosCluster) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TargetExecutionConfigs struct {
	empty            bool                               `json:"-"`
	Usages           []TargetExecutionConfigsUsagesEnum `json:"usages"`
	WorkerPool       *string                            `json:"workerPool"`
	ServiceAccount   *string                            `json:"serviceAccount"`
	ArtifactStorage  *string                            `json:"artifactStorage"`
	ExecutionTimeout *string                            `json:"executionTimeout"`
	Verbose          *bool                              `json:"verbose"`
}

type jsonTargetExecutionConfigs TargetExecutionConfigs

func (r *TargetExecutionConfigs) UnmarshalJSON(data []byte) error {
	var res jsonTargetExecutionConfigs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTargetExecutionConfigs
	} else {

		r.Usages = res.Usages

		r.WorkerPool = res.WorkerPool

		r.ServiceAccount = res.ServiceAccount

		r.ArtifactStorage = res.ArtifactStorage

		r.ExecutionTimeout = res.ExecutionTimeout

		r.Verbose = res.Verbose

	}
	return nil
}

// This object is used to assert a desired state where this TargetExecutionConfigs is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTargetExecutionConfigs *TargetExecutionConfigs = &TargetExecutionConfigs{empty: true}

func (r *TargetExecutionConfigs) Empty() bool {
	return r.empty
}

func (r *TargetExecutionConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *TargetExecutionConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TargetRun struct {
	empty    bool    `json:"-"`
	Location *string `json:"location"`
}

type jsonTargetRun TargetRun

func (r *TargetRun) UnmarshalJSON(data []byte) error {
	var res jsonTargetRun
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTargetRun
	} else {

		r.Location = res.Location

	}
	return nil
}

// This object is used to assert a desired state where this TargetRun is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTargetRun *TargetRun = &TargetRun{empty: true}

func (r *TargetRun) Empty() bool {
	return r.empty
}

func (r *TargetRun) String() string {
	return dcl.SprintResource(r)
}

func (r *TargetRun) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TargetMultiTarget struct {
	empty     bool     `json:"-"`
	TargetIds []string `json:"targetIds"`
}

type jsonTargetMultiTarget TargetMultiTarget

func (r *TargetMultiTarget) UnmarshalJSON(data []byte) error {
	var res jsonTargetMultiTarget
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTargetMultiTarget
	} else {

		r.TargetIds = res.TargetIds

	}
	return nil
}

// This object is used to assert a desired state where this TargetMultiTarget is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTargetMultiTarget *TargetMultiTarget = &TargetMultiTarget{empty: true}

func (r *TargetMultiTarget) Empty() bool {
	return r.empty
}

func (r *TargetMultiTarget) String() string {
	return dcl.SprintResource(r)
}

func (r *TargetMultiTarget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TargetCustomTarget struct {
	empty            bool    `json:"-"`
	CustomTargetType *string `json:"customTargetType"`
}

type jsonTargetCustomTarget TargetCustomTarget

func (r *TargetCustomTarget) UnmarshalJSON(data []byte) error {
	var res jsonTargetCustomTarget
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTargetCustomTarget
	} else {

		r.CustomTargetType = res.CustomTargetType

	}
	return nil
}

// This object is used to assert a desired state where this TargetCustomTarget is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTargetCustomTarget *TargetCustomTarget = &TargetCustomTarget{empty: true}

func (r *TargetCustomTarget) Empty() bool {
	return r.empty
}

func (r *TargetCustomTarget) String() string {
	return dcl.SprintResource(r)
}

func (r *TargetCustomTarget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TargetAssociatedEntities struct {
	empty          bool                                     `json:"-"`
	GkeClusters    []TargetAssociatedEntitiesGkeClusters    `json:"gkeClusters"`
	AnthosClusters []TargetAssociatedEntitiesAnthosClusters `json:"anthosClusters"`
}

type jsonTargetAssociatedEntities TargetAssociatedEntities

func (r *TargetAssociatedEntities) UnmarshalJSON(data []byte) error {
	var res jsonTargetAssociatedEntities
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTargetAssociatedEntities
	} else {

		r.GkeClusters = res.GkeClusters

		r.AnthosClusters = res.AnthosClusters

	}
	return nil
}

// This object is used to assert a desired state where this TargetAssociatedEntities is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTargetAssociatedEntities *TargetAssociatedEntities = &TargetAssociatedEntities{empty: true}

func (r *TargetAssociatedEntities) Empty() bool {
	return r.empty
}

func (r *TargetAssociatedEntities) String() string {
	return dcl.SprintResource(r)
}

func (r *TargetAssociatedEntities) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TargetAssociatedEntitiesGkeClusters struct {
	empty      bool    `json:"-"`
	Cluster    *string `json:"cluster"`
	InternalIP *bool   `json:"internalIP"`
	ProxyUrl   *string `json:"proxyUrl"`
}

type jsonTargetAssociatedEntitiesGkeClusters TargetAssociatedEntitiesGkeClusters

func (r *TargetAssociatedEntitiesGkeClusters) UnmarshalJSON(data []byte) error {
	var res jsonTargetAssociatedEntitiesGkeClusters
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTargetAssociatedEntitiesGkeClusters
	} else {

		r.Cluster = res.Cluster

		r.InternalIP = res.InternalIP

		r.ProxyUrl = res.ProxyUrl

	}
	return nil
}

// This object is used to assert a desired state where this TargetAssociatedEntitiesGkeClusters is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTargetAssociatedEntitiesGkeClusters *TargetAssociatedEntitiesGkeClusters = &TargetAssociatedEntitiesGkeClusters{empty: true}

func (r *TargetAssociatedEntitiesGkeClusters) Empty() bool {
	return r.empty
}

func (r *TargetAssociatedEntitiesGkeClusters) String() string {
	return dcl.SprintResource(r)
}

func (r *TargetAssociatedEntitiesGkeClusters) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TargetAssociatedEntitiesAnthosClusters struct {
	empty      bool    `json:"-"`
	Membership *string `json:"membership"`
}

type jsonTargetAssociatedEntitiesAnthosClusters TargetAssociatedEntitiesAnthosClusters

func (r *TargetAssociatedEntitiesAnthosClusters) UnmarshalJSON(data []byte) error {
	var res jsonTargetAssociatedEntitiesAnthosClusters
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTargetAssociatedEntitiesAnthosClusters
	} else {

		r.Membership = res.Membership

	}
	return nil
}

// This object is used to assert a desired state where this TargetAssociatedEntitiesAnthosClusters is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTargetAssociatedEntitiesAnthosClusters *TargetAssociatedEntitiesAnthosClusters = &TargetAssociatedEntitiesAnthosClusters{empty: true}

func (r *TargetAssociatedEntitiesAnthosClusters) Empty() bool {
	return r.empty
}

func (r *TargetAssociatedEntitiesAnthosClusters) String() string {
	return dcl.SprintResource(r)
}

func (r *TargetAssociatedEntitiesAnthosClusters) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Target) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "clouddeploy",
		Type:    "Target",
		Version: "clouddeploy",
	}
}

func (r *Target) ID() (string, error) {
	if err := extractTargetFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                dcl.ValueOrEmptyString(nr.Name),
		"target_id":           dcl.ValueOrEmptyString(nr.TargetId),
		"uid":                 dcl.ValueOrEmptyString(nr.Uid),
		"description":         dcl.ValueOrEmptyString(nr.Description),
		"annotations":         dcl.ValueOrEmptyString(nr.Annotations),
		"labels":              dcl.ValueOrEmptyString(nr.Labels),
		"require_approval":    dcl.ValueOrEmptyString(nr.RequireApproval),
		"create_time":         dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":         dcl.ValueOrEmptyString(nr.UpdateTime),
		"gke":                 dcl.ValueOrEmptyString(nr.Gke),
		"anthos_cluster":      dcl.ValueOrEmptyString(nr.AnthosCluster),
		"etag":                dcl.ValueOrEmptyString(nr.Etag),
		"execution_configs":   dcl.ValueOrEmptyString(nr.ExecutionConfigs),
		"project":             dcl.ValueOrEmptyString(nr.Project),
		"location":            dcl.ValueOrEmptyString(nr.Location),
		"run":                 dcl.ValueOrEmptyString(nr.Run),
		"multi_target":        dcl.ValueOrEmptyString(nr.MultiTarget),
		"deploy_parameters":   dcl.ValueOrEmptyString(nr.DeployParameters),
		"custom_target":       dcl.ValueOrEmptyString(nr.CustomTarget),
		"associated_entities": dcl.ValueOrEmptyString(nr.AssociatedEntities),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/targets/{{name}}", params), nil
}

const TargetMaxPage = -1

type TargetList struct {
	Items []*Target

	nextToken string

	pageSize int32

	resource *Target
}

func (l *TargetList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TargetList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTarget(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTarget(ctx context.Context, project, location string) (*TargetList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTargetWithMaxResults(ctx, project, location, TargetMaxPage)

}

func (c *Client) ListTargetWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*TargetList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Target{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listTarget(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TargetList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetTarget(ctx context.Context, r *Target) (*Target, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractTargetFields(r)

	b, err := c.getTargetRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTarget(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTargetNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractTargetFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTarget(ctx context.Context, r *Target) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Target resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Target...")
	deleteOp := deleteTargetOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTarget deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTarget(ctx context.Context, project, location string, filter func(*Target) bool) error {
	listObj, err := c.ListTarget(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllTarget(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTarget(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTarget(ctx context.Context, rawDesired *Target, opts ...dcl.ApplyOption) (*Target, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Target
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTargetHelper(c, ctx, rawDesired, opts...)
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

func applyTargetHelper(c *Client, ctx context.Context, rawDesired *Target, opts ...dcl.ApplyOption) (*Target, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyTarget...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractTargetFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.targetDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTargetDiffs(c.Config, fieldDiffs, opts)
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
	var ops []targetApiOperation
	if create {
		ops = append(ops, &createTargetOperation{})
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
	return applyTargetDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyTargetDiff(c *Client, ctx context.Context, desired *Target, rawDesired *Target, ops []targetApiOperation, opts ...dcl.ApplyOption) (*Target, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetTarget(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTargetOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTarget(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTargetNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTargetNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTargetDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractTargetFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractTargetFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTarget(c, newDesired, newState)
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

func (r *Target) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "", body, nil
}
