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
package apigee

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Organization struct {
	Name                             *string                           `json:"name"`
	DisplayName                      *string                           `json:"displayName"`
	Description                      *string                           `json:"description"`
	CreatedAt                        *int64                            `json:"createdAt"`
	LastModifiedAt                   *int64                            `json:"lastModifiedAt"`
	ExpiresAt                        *int64                            `json:"expiresAt"`
	Environments                     []string                          `json:"environments"`
	Properties                       *OrganizationProperties           `json:"properties"`
	AnalyticsRegion                  *string                           `json:"analyticsRegion"`
	AuthorizedNetwork                *string                           `json:"authorizedNetwork"`
	RuntimeType                      *OrganizationRuntimeTypeEnum      `json:"runtimeType"`
	SubscriptionType                 *OrganizationSubscriptionTypeEnum `json:"subscriptionType"`
	BillingType                      *OrganizationBillingTypeEnum      `json:"billingType"`
	CaCertificate                    *string                           `json:"caCertificate"`
	RuntimeDatabaseEncryptionKeyName *string                           `json:"runtimeDatabaseEncryptionKeyName"`
	ProjectId                        *string                           `json:"projectId"`
	State                            *OrganizationStateEnum            `json:"state"`
	Parent                           *string                           `json:"parent"`
}

func (r *Organization) String() string {
	return dcl.SprintResource(r)
}

// The enum OrganizationRuntimeTypeEnum.
type OrganizationRuntimeTypeEnum string

// OrganizationRuntimeTypeEnumRef returns a *OrganizationRuntimeTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func OrganizationRuntimeTypeEnumRef(s string) *OrganizationRuntimeTypeEnum {
	if s == "" {
		return nil
	}

	v := OrganizationRuntimeTypeEnum(s)
	return &v
}

func (v OrganizationRuntimeTypeEnum) Validate() error {
	for _, s := range []string{"RUNTIME_TYPE_UNSPECIFIED", "CLOUD", "HYBRID"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OrganizationRuntimeTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OrganizationSubscriptionTypeEnum.
type OrganizationSubscriptionTypeEnum string

// OrganizationSubscriptionTypeEnumRef returns a *OrganizationSubscriptionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func OrganizationSubscriptionTypeEnumRef(s string) *OrganizationSubscriptionTypeEnum {
	if s == "" {
		return nil
	}

	v := OrganizationSubscriptionTypeEnum(s)
	return &v
}

func (v OrganizationSubscriptionTypeEnum) Validate() error {
	for _, s := range []string{"SUBSCRIPTION_TYPE_UNSPECIFIED", "PAID", "TRIAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OrganizationSubscriptionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OrganizationBillingTypeEnum.
type OrganizationBillingTypeEnum string

// OrganizationBillingTypeEnumRef returns a *OrganizationBillingTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func OrganizationBillingTypeEnumRef(s string) *OrganizationBillingTypeEnum {
	if s == "" {
		return nil
	}

	v := OrganizationBillingTypeEnum(s)
	return &v
}

func (v OrganizationBillingTypeEnum) Validate() error {
	for _, s := range []string{"BILLING_TYPE_UNSPECIFIED", "SUBSCRIPTION", "EVALUATION", "PAYG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OrganizationBillingTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OrganizationStateEnum.
type OrganizationStateEnum string

// OrganizationStateEnumRef returns a *OrganizationStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func OrganizationStateEnumRef(s string) *OrganizationStateEnum {
	if s == "" {
		return nil
	}

	v := OrganizationStateEnum(s)
	return &v
}

func (v OrganizationStateEnum) Validate() error {
	for _, s := range []string{"SNAPSHOT_STATE_UNSPECIFIED", "MISSING", "OK_DOCSTORE", "OK_SUBMITTED", "OK_EXTERNAL", "DELETED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OrganizationStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type OrganizationProperties struct {
	empty    bool                             `json:"-"`
	Property []OrganizationPropertiesProperty `json:"property"`
}

type jsonOrganizationProperties OrganizationProperties

func (r *OrganizationProperties) UnmarshalJSON(data []byte) error {
	var res jsonOrganizationProperties
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOrganizationProperties
	} else {

		r.Property = res.Property

	}
	return nil
}

// This object is used to assert a desired state where this OrganizationProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOrganizationProperties *OrganizationProperties = &OrganizationProperties{empty: true}

func (r *OrganizationProperties) Empty() bool {
	return r.empty
}

func (r *OrganizationProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *OrganizationProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OrganizationPropertiesProperty struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

type jsonOrganizationPropertiesProperty OrganizationPropertiesProperty

func (r *OrganizationPropertiesProperty) UnmarshalJSON(data []byte) error {
	var res jsonOrganizationPropertiesProperty
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOrganizationPropertiesProperty
	} else {

		r.Name = res.Name

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this OrganizationPropertiesProperty is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyOrganizationPropertiesProperty *OrganizationPropertiesProperty = &OrganizationPropertiesProperty{empty: true}

func (r *OrganizationPropertiesProperty) Empty() bool {
	return r.empty
}

func (r *OrganizationPropertiesProperty) String() string {
	return dcl.SprintResource(r)
}

func (r *OrganizationPropertiesProperty) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Organization) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "apigee",
		Type:    "Organization",
		Version: "apigee",
	}
}

const OrganizationMaxPage = -1

type OrganizationList struct {
	Items []*Organization

	nextToken string

	pageSize int32
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Organization) URLNormalized() *Organization {
	normalized := dcl.Copy(*r).(Organization)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AnalyticsRegion = dcl.SelfLinkToName(r.AnalyticsRegion)
	normalized.AuthorizedNetwork = dcl.SelfLinkToName(r.AuthorizedNetwork)
	normalized.CaCertificate = dcl.SelfLinkToName(r.CaCertificate)
	normalized.RuntimeDatabaseEncryptionKeyName = dcl.SelfLinkToName(r.RuntimeDatabaseEncryptionKeyName)
	normalized.ProjectId = dcl.SelfLinkToName(r.ProjectId)
	normalized.Parent = r.Parent
	return &normalized
}

func (c *Client) GetOrganization(ctx context.Context, r *Organization) (*Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getOrganizationRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalOrganization(b, c)
	if err != nil {
		return nil, err
	}
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeOrganizationNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteOrganization(ctx context.Context, r *Organization) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Organization resource is nil")
	}
	c.Config.Logger.Info("Deleting Organization...")
	deleteOp := deleteOrganizationOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllOrganization deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllOrganization(ctx context.Context, filter func(*Organization) bool) error {
	listObj, err := c.ListOrganization(ctx)
	if err != nil {
		return err
	}

	err = c.deleteAllOrganization(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllOrganization(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyOrganization(ctx context.Context, rawDesired *Organization, opts ...dcl.ApplyOption) (*Organization, error) {

	var resultNewState *Organization
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyOrganizationHelper(c, ctx, rawDesired, opts...)
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

func applyOrganizationHelper(c *Client, ctx context.Context, rawDesired *Organization, opts ...dcl.ApplyOption) (*Organization, error) {
	c.Config.Logger.Info("Beginning ApplyOrganization...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.organizationDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToOrganizationDiffs(c.Config, fieldDiffs, opts)
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
	var ops []organizationApiOperation
	if create {
		ops = append(ops, &createOrganizationOperation{})
	} else if recreate {
		ops = append(ops, &deleteOrganizationOperation{})
		ops = append(ops, &createOrganizationOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeOrganizationDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetOrganization(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createOrganizationOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapOrganization(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeOrganizationNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeOrganizationNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeOrganizationDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffOrganization(c, newDesired, newState)
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
