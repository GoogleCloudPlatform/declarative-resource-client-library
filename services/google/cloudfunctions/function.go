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
package cloudfunctions

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Function struct {
	Name                       *string                                 `json:"name"`
	Description                *string                                 `json:"description"`
	SourceArchiveUrl           *string                                 `json:"sourceArchiveUrl"`
	SourceRepository           *FunctionSourceRepository               `json:"sourceRepository"`
	HttpsTrigger               *FunctionHttpsTrigger                   `json:"httpsTrigger"`
	EventTrigger               *FunctionEventTrigger                   `json:"eventTrigger"`
	Status                     *FunctionStatusEnum                     `json:"status"`
	EntryPoint                 *string                                 `json:"entryPoint"`
	Runtime                    *string                                 `json:"runtime"`
	Timeout                    *string                                 `json:"timeout"`
	AvailableMemoryMb          *int64                                  `json:"availableMemoryMb"`
	ServiceAccountEmail        *string                                 `json:"serviceAccountEmail"`
	UpdateTime                 *string                                 `json:"updateTime"`
	VersionId                  *int64                                  `json:"versionId"`
	Labels                     map[string]string                       `json:"labels"`
	EnvironmentVariables       map[string]string                       `json:"environmentVariables"`
	Network                    *string                                 `json:"network"`
	MaxInstances               *int64                                  `json:"maxInstances"`
	VPCConnector               *string                                 `json:"vpcConnector"`
	VPCConnectorEgressSettings *FunctionVPCConnectorEgressSettingsEnum `json:"vpcConnectorEgressSettings"`
	IngressSettings            *FunctionIngressSettingsEnum            `json:"ingressSettings"`
	Region                     *string                                 `json:"region"`
	Project                    *string                                 `json:"project"`
}

func (r *Function) String() string {
	return dcl.SprintResource(r)
}

// The enum FunctionStatusEnum.
type FunctionStatusEnum string

// FunctionStatusEnumRef returns a *FunctionStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func FunctionStatusEnumRef(s string) *FunctionStatusEnum {
	if s == "" {
		return nil
	}

	v := FunctionStatusEnum(s)
	return &v
}

func (v FunctionStatusEnum) Validate() error {
	for _, s := range []string{"CLOUD_FUNCTION_STATUS_UNSPECIFIED", "ACTIVE", "OFFLINE", "DEPLOY_IN_PROGRESS", "DELETE_IN_PROGRESS", "UNKNOWN"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FunctionStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FunctionVPCConnectorEgressSettingsEnum.
type FunctionVPCConnectorEgressSettingsEnum string

// FunctionVPCConnectorEgressSettingsEnumRef returns a *FunctionVPCConnectorEgressSettingsEnum with the value of string s
// If the empty string is provided, nil is returned.
func FunctionVPCConnectorEgressSettingsEnumRef(s string) *FunctionVPCConnectorEgressSettingsEnum {
	if s == "" {
		return nil
	}

	v := FunctionVPCConnectorEgressSettingsEnum(s)
	return &v
}

func (v FunctionVPCConnectorEgressSettingsEnum) Validate() error {
	for _, s := range []string{"VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED", "PRIVATE_RANGES_ONLY", "ALL_TRAFFIC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FunctionVPCConnectorEgressSettingsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FunctionIngressSettingsEnum.
type FunctionIngressSettingsEnum string

// FunctionIngressSettingsEnumRef returns a *FunctionIngressSettingsEnum with the value of string s
// If the empty string is provided, nil is returned.
func FunctionIngressSettingsEnumRef(s string) *FunctionIngressSettingsEnum {
	if s == "" {
		return nil
	}

	v := FunctionIngressSettingsEnum(s)
	return &v
}

func (v FunctionIngressSettingsEnum) Validate() error {
	for _, s := range []string{"INGRESS_SETTINGS_UNSPECIFIED", "ALLOW_ALL", "ALLOW_INTERNAL_ONLY", "ALLOW_INTERNAL_AND_GCLB"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FunctionIngressSettingsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type FunctionSourceRepository struct {
	empty       bool    `json:"-"`
	Url         *string `json:"url"`
	DeployedUrl *string `json:"deployedUrl"`
}

// This object is used to assert a desired state where this FunctionSourceRepository is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFunctionSourceRepository *FunctionSourceRepository = &FunctionSourceRepository{empty: true}

func (r *FunctionSourceRepository) String() string {
	return dcl.SprintResource(r)
}

func (r *FunctionSourceRepository) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FunctionHttpsTrigger struct {
	empty bool    `json:"-"`
	Url   *string `json:"url"`
}

// This object is used to assert a desired state where this FunctionHttpsTrigger is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFunctionHttpsTrigger *FunctionHttpsTrigger = &FunctionHttpsTrigger{empty: true}

func (r *FunctionHttpsTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *FunctionHttpsTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FunctionEventTrigger struct {
	empty         bool    `json:"-"`
	EventType     *string `json:"eventType"`
	Resource      *string `json:"resource"`
	Service       *string `json:"service"`
	FailurePolicy *bool   `json:"failurePolicy"`
}

// This object is used to assert a desired state where this FunctionEventTrigger is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFunctionEventTrigger *FunctionEventTrigger = &FunctionEventTrigger{empty: true}

func (r *FunctionEventTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *FunctionEventTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Function) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloudfunctions",
		Type:    "Function",
		Version: "cloudfunctions",
	}
}

const FunctionMaxPage = -1

type FunctionList struct {
	Items []*Function

	nextToken string

	pageSize int32

	project string

	region string
}

func (l *FunctionList) HasNext() bool {
	return l.nextToken != ""
}

func (l *FunctionList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listFunction(ctx, l.project, l.region, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListFunction(ctx context.Context, project, region string) (*FunctionList, error) {

	return c.ListFunctionWithMaxResults(ctx, project, region, FunctionMaxPage)

}

func (c *Client) ListFunctionWithMaxResults(ctx context.Context, project, region string, pageSize int32) (*FunctionList, error) {
	items, token, err := c.listFunction(ctx, project, region, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &FunctionList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		region: region,
	}, nil
}

func (c *Client) GetFunction(ctx context.Context, r *Function) (*Function, error) {
	b, err := c.getFunctionRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalFunction(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Region = r.Region
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFunctionNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteFunction(ctx context.Context, r *Function) error {
	if r == nil {
		return fmt.Errorf("Function resource is nil")
	}
	c.Config.Logger.Info("Deleting Function...")
	deleteOp := deleteFunctionOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllFunction deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllFunction(ctx context.Context, project, region string, filter func(*Function) bool) error {
	listObj, err := c.ListFunction(ctx, project, region)
	if err != nil {
		return err
	}

	err = c.deleteAllFunction(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllFunction(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyFunction(ctx context.Context, rawDesired *Function, opts ...dcl.ApplyOption) (*Function, error) {
	c.Config.Logger.Info("Beginning ApplyFunction...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.functionDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
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
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []functionApiOperation
	if create {
		ops = append(ops, &createFunctionOperation{})
	} else if recreate {

		ops = append(ops, &deleteFunctionOperation{})

		ops = append(ops, &createFunctionOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeFunctionDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %#v", op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %#v: %v", op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %#v", op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetFunction(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFunctionNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFunctionDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFunction(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
