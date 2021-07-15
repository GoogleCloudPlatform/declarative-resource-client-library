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
package identitytoolkit

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Tenant struct {
	Name                  *string           `json:"name"`
	DisplayName           *string           `json:"displayName"`
	AllowPasswordSignup   *bool             `json:"allowPasswordSignup"`
	EnableEmailLinkSignin *bool             `json:"enableEmailLinkSignin"`
	DisableAuth           *bool             `json:"disableAuth"`
	EnableAnonymousUser   *bool             `json:"enableAnonymousUser"`
	MfaConfig             *TenantMfaConfig  `json:"mfaConfig"`
	TestPhoneNumbers      map[string]string `json:"testPhoneNumbers"`
	Project               *string           `json:"project"`
}

func (r *Tenant) String() string {
	return dcl.SprintResource(r)
}

// The enum TenantMfaConfigStateEnum.
type TenantMfaConfigStateEnum string

// TenantMfaConfigStateEnumRef returns a *TenantMfaConfigStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func TenantMfaConfigStateEnumRef(s string) *TenantMfaConfigStateEnum {
	if s == "" {
		return nil
	}

	v := TenantMfaConfigStateEnum(s)
	return &v
}

func (v TenantMfaConfigStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "DISABLED", "ENABLED", "MANDATORY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TenantMfaConfigStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum TenantMfaConfigEnabledProvidersEnum.
type TenantMfaConfigEnabledProvidersEnum string

// TenantMfaConfigEnabledProvidersEnumRef returns a *TenantMfaConfigEnabledProvidersEnum with the value of string s
// If the empty string is provided, nil is returned.
func TenantMfaConfigEnabledProvidersEnumRef(s string) *TenantMfaConfigEnabledProvidersEnum {
	if s == "" {
		return nil
	}

	v := TenantMfaConfigEnabledProvidersEnum(s)
	return &v
}

func (v TenantMfaConfigEnabledProvidersEnum) Validate() error {
	for _, s := range []string{"PROVIDER_UNSPECIFIED", "PHONE_SMS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TenantMfaConfigEnabledProvidersEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type TenantMfaConfig struct {
	empty            bool                                  `json:"-"`
	State            *TenantMfaConfigStateEnum             `json:"state"`
	EnabledProviders []TenantMfaConfigEnabledProvidersEnum `json:"enabledProviders"`
}

type jsonTenantMfaConfig TenantMfaConfig

func (r *TenantMfaConfig) UnmarshalJSON(data []byte) error {
	var res jsonTenantMfaConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTenantMfaConfig
	} else {

		r.State = res.State

		r.EnabledProviders = res.EnabledProviders

	}
	return nil
}

// This object is used to assert a desired state where this TenantMfaConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTenantMfaConfig *TenantMfaConfig = &TenantMfaConfig{empty: true}

func (r *TenantMfaConfig) Empty() bool {
	return r.empty
}

func (r *TenantMfaConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *TenantMfaConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Tenant) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "identity_toolkit",
		Type:    "Tenant",
		Version: "identitytoolkit",
	}
}

const TenantMaxPage = -1

type TenantList struct {
	Items []*Tenant

	nextToken string

	pageSize int32

	project string
}

func (l *TenantList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TenantList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTenant(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTenant(ctx context.Context, project string) (*TenantList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTenantWithMaxResults(ctx, project, TenantMaxPage)

}

func (c *Client) ListTenantWithMaxResults(ctx context.Context, project string, pageSize int32) (*TenantList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listTenant(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TenantList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Tenant) URLNormalized() *Tenant {
	normalized := dcl.Copy(*r).(Tenant)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetTenant(ctx context.Context, r *Tenant) (*Tenant, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getTenantRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTenant(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTenantNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTenant(ctx context.Context, r *Tenant) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Tenant resource is nil")
	}
	c.Config.Logger.Info("Deleting Tenant...")
	deleteOp := deleteTenantOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTenant deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTenant(ctx context.Context, project string, filter func(*Tenant) bool) error {
	listObj, err := c.ListTenant(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllTenant(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTenant(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTenant(ctx context.Context, rawDesired *Tenant, opts ...dcl.ApplyOption) (*Tenant, error) {
	var resultNewState *Tenant
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTenantHelper(c, ctx, rawDesired, opts...)
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

func applyTenantHelper(c *Client, ctx context.Context, rawDesired *Tenant, opts ...dcl.ApplyOption) (*Tenant, error) {
	c.Config.Logger.Info("Beginning ApplyTenant...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.tenantDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTenantDiffs(c.Config, fieldDiffs, opts)
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
	var ops []tenantApiOperation
	if create {
		ops = append(ops, &createTenantOperation{})
	} else if recreate {
		ops = append(ops, &deleteTenantOperation{})
		ops = append(ops, &createTenantOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeTenantDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetTenant(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTenantOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTenant(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTenantNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTenantNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTenantDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTenant(c, newDesired, newState)
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
