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

type Workload struct {
	Name                       *string                       `json:"name"`
	DisplayName                *string                       `json:"displayName"`
	Resources                  []WorkloadResources           `json:"resources"`
	ComplianceRegime           *WorkloadComplianceRegimeEnum `json:"complianceRegime"`
	CreateTime                 *string                       `json:"createTime"`
	BillingAccount             *string                       `json:"billingAccount"`
	Labels                     map[string]string             `json:"labels"`
	ProvisionedResourcesParent *string                       `json:"provisionedResourcesParent"`
	KmsSettings                *WorkloadKmsSettings          `json:"kmsSettings"`
	ResourceSettings           []WorkloadResourceSettings    `json:"resourceSettings"`
	Organization               *string                       `json:"organization"`
	Location                   *string                       `json:"location"`
}

func (r *Workload) String() string {
	return dcl.SprintResource(r)
}

// The enum WorkloadResourcesResourceTypeEnum.
type WorkloadResourcesResourceTypeEnum string

// WorkloadResourcesResourceTypeEnumRef returns a *WorkloadResourcesResourceTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadResourcesResourceTypeEnumRef(s string) *WorkloadResourcesResourceTypeEnum {
	if s == "" {
		return nil
	}

	v := WorkloadResourcesResourceTypeEnum(s)
	return &v
}

func (v WorkloadResourcesResourceTypeEnum) Validate() error {
	for _, s := range []string{"RESOURCE_TYPE_UNSPECIFIED", "CONSUMER_PROJECT", "ENCRYPTION_KEYS_PROJECT", "KEYRING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadResourcesResourceTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadComplianceRegimeEnum.
type WorkloadComplianceRegimeEnum string

// WorkloadComplianceRegimeEnumRef returns a *WorkloadComplianceRegimeEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadComplianceRegimeEnumRef(s string) *WorkloadComplianceRegimeEnum {
	if s == "" {
		return nil
	}

	v := WorkloadComplianceRegimeEnum(s)
	return &v
}

func (v WorkloadComplianceRegimeEnum) Validate() error {
	for _, s := range []string{"COMPLIANCE_REGIME_UNSPECIFIED", "IL4", "CJIS", "FEDRAMP_HIGH", "FEDRAMP_MODERATE", "US_REGIONAL_ACCESS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadComplianceRegimeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadResourceSettingsResourceTypeEnum.
type WorkloadResourceSettingsResourceTypeEnum string

// WorkloadResourceSettingsResourceTypeEnumRef returns a *WorkloadResourceSettingsResourceTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadResourceSettingsResourceTypeEnumRef(s string) *WorkloadResourceSettingsResourceTypeEnum {
	if s == "" {
		return nil
	}

	v := WorkloadResourceSettingsResourceTypeEnum(s)
	return &v
}

func (v WorkloadResourceSettingsResourceTypeEnum) Validate() error {
	for _, s := range []string{"RESOURCE_TYPE_UNSPECIFIED", "CONSUMER_PROJECT", "ENCRYPTION_KEYS_PROJECT", "KEYRING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadResourceSettingsResourceTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type WorkloadResources struct {
	empty        bool                               `json:"-"`
	ResourceId   *int64                             `json:"resourceId"`
	ResourceType *WorkloadResourcesResourceTypeEnum `json:"resourceType"`
}

type jsonWorkloadResources WorkloadResources

func (r *WorkloadResources) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadResources
	} else {

		r.ResourceId = res.ResourceId

		r.ResourceType = res.ResourceType

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadResources is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkloadResources *WorkloadResources = &WorkloadResources{empty: true}

func (r *WorkloadResources) Empty() bool {
	return r.empty
}

func (r *WorkloadResources) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkloadKmsSettings struct {
	empty            bool    `json:"-"`
	NextRotationTime *string `json:"nextRotationTime"`
	RotationPeriod   *string `json:"rotationPeriod"`
}

type jsonWorkloadKmsSettings WorkloadKmsSettings

func (r *WorkloadKmsSettings) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadKmsSettings
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadKmsSettings
	} else {

		r.NextRotationTime = res.NextRotationTime

		r.RotationPeriod = res.RotationPeriod

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadKmsSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkloadKmsSettings *WorkloadKmsSettings = &WorkloadKmsSettings{empty: true}

func (r *WorkloadKmsSettings) Empty() bool {
	return r.empty
}

func (r *WorkloadKmsSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadKmsSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkloadResourceSettings struct {
	empty        bool                                      `json:"-"`
	ResourceId   *string                                   `json:"resourceId"`
	ResourceType *WorkloadResourceSettingsResourceTypeEnum `json:"resourceType"`
}

type jsonWorkloadResourceSettings WorkloadResourceSettings

func (r *WorkloadResourceSettings) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadResourceSettings
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadResourceSettings
	} else {

		r.ResourceId = res.ResourceId

		r.ResourceType = res.ResourceType

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadResourceSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkloadResourceSettings *WorkloadResourceSettings = &WorkloadResourceSettings{empty: true}

func (r *WorkloadResourceSettings) Empty() bool {
	return r.empty
}

func (r *WorkloadResourceSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadResourceSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Workload) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "assured_workloads",
		Type:    "Workload",
		Version: "beta",
	}
}

const WorkloadMaxPage = -1

type WorkloadList struct {
	Items []*Workload

	nextToken string

	pageSize int32

	organization string

	location string
}

func (l *WorkloadList) HasNext() bool {
	return l.nextToken != ""
}

func (l *WorkloadList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listWorkload(ctx, l.organization, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListWorkload(ctx context.Context, organization, location string) (*WorkloadList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListWorkloadWithMaxResults(ctx, organization, location, WorkloadMaxPage)

}

func (c *Client) ListWorkloadWithMaxResults(ctx context.Context, organization, location string, pageSize int32) (*WorkloadList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listWorkload(ctx, organization, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &WorkloadList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		organization: organization,

		location: location,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Workload) URLNormalized() *Workload {
	normalized := dcl.Copy(*r).(Workload)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.BillingAccount = dcl.SelfLinkToName(r.BillingAccount)
	normalized.ProvisionedResourcesParent = dcl.SelfLinkToName(r.ProvisionedResourcesParent)
	normalized.Organization = dcl.SelfLinkToName(r.Organization)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetWorkload(ctx context.Context, r *Workload) (*Workload, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getWorkloadRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalWorkload(b, c)
	if err != nil {
		return nil, err
	}
	result.Organization = r.Organization
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeWorkloadNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteWorkload(ctx context.Context, r *Workload) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Workload resource is nil")
	}
	c.Config.Logger.Info("Deleting Workload...")
	deleteOp := deleteWorkloadOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllWorkload deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllWorkload(ctx context.Context, organization, location string, filter func(*Workload) bool) error {
	listObj, err := c.ListWorkload(ctx, organization, location)
	if err != nil {
		return err
	}

	err = c.deleteAllWorkload(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllWorkload(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyWorkload(ctx context.Context, rawDesired *Workload, opts ...dcl.ApplyOption) (*Workload, error) {
	var resultNewState *Workload
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyWorkloadHelper(c, ctx, rawDesired, opts...)
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

func applyWorkloadHelper(c *Client, ctx context.Context, rawDesired *Workload, opts ...dcl.ApplyOption) (*Workload, error) {
	c.Config.Logger.Info("Beginning ApplyWorkload...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.workloadDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToWorkloadDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
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
	var ops []workloadApiOperation
	if create {
		ops = append(ops, &createWorkloadOperation{})
	} else if recreate {
		ops = append(ops, &deleteWorkloadOperation{})
		ops = append(ops, &createWorkloadOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeWorkloadDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetWorkload(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createWorkloadOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapWorkload(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeWorkloadNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeWorkloadNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeWorkloadDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffWorkload(c, newDesired, newState)
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
