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
package accesscontextmanager

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type AccessLevel struct {
	Title       *string           `json:"title"`
	CreateTime  *string           `json:"createTime"`
	UpdateTime  *string           `json:"updateTime"`
	Description *string           `json:"description"`
	Basic       *AccessLevelBasic `json:"basic"`
	Name        *string           `json:"name"`
	Policy      *string           `json:"policy"`
}

func (r *AccessLevel) String() string {
	return dcl.SprintResource(r)
}

// The enum AccessLevelBasicCombiningFunctionEnum.
type AccessLevelBasicCombiningFunctionEnum string

// AccessLevelBasicCombiningFunctionEnumRef returns a *AccessLevelBasicCombiningFunctionEnum with the value of string s
// If the empty string is provided, nil is returned.
func AccessLevelBasicCombiningFunctionEnumRef(s string) *AccessLevelBasicCombiningFunctionEnum {
	if s == "" {
		return nil
	}

	v := AccessLevelBasicCombiningFunctionEnum(s)
	return &v
}

func (v AccessLevelBasicCombiningFunctionEnum) Validate() error {
	for _, s := range []string{"AND", "OR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AccessLevelBasicCombiningFunctionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum.
type AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum string

// AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumRef returns a *AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum with the value of string s
// If the empty string is provided, nil is returned.
func AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumRef(s string) *AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum {
	if s == "" {
		return nil
	}

	v := AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(s)
	return &v
}

func (v AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum) Validate() error {
	for _, s := range []string{"ENCRYPTION_UNSPECIFIED", "ENCRYPTION_UNSUPPORTED", "UNENCRYPTED", "ENCRYPTED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum.
type AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum string

// AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumRef returns a *AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum with the value of string s
// If the empty string is provided, nil is returned.
func AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumRef(s string) *AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum {
	if s == "" {
		return nil
	}

	v := AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(s)
	return &v
}

func (v AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum) Validate() error {
	for _, s := range []string{"MANAGEMENT_UNSPECIFIED", "NONE", "BASIC", "COMPLETE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum.
type AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum string

// AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnumRef returns a *AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnumRef(s string) *AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum {
	if s == "" {
		return nil
	}

	v := AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum(s)
	return &v
}

func (v AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum) Validate() error {
	for _, s := range []string{"OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AccessLevelBasic struct {
	empty             bool                                   `json:"-"`
	CombiningFunction *AccessLevelBasicCombiningFunctionEnum `json:"combiningFunction"`
	Conditions        []AccessLevelBasicConditions           `json:"conditions"`
}

type jsonAccessLevelBasic AccessLevelBasic

func (r *AccessLevelBasic) UnmarshalJSON(data []byte) error {
	var res jsonAccessLevelBasic
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAccessLevelBasic
	} else {

		r.CombiningFunction = res.CombiningFunction

		r.Conditions = res.Conditions

	}
	return nil
}

// This object is used to assert a desired state where this AccessLevelBasic is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAccessLevelBasic *AccessLevelBasic = &AccessLevelBasic{empty: true}

func (r *AccessLevelBasic) Empty() bool {
	return r.empty
}

func (r *AccessLevelBasic) String() string {
	return dcl.SprintResource(r)
}

func (r *AccessLevelBasic) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AccessLevelBasicConditions struct {
	empty                bool                                    `json:"-"`
	Regions              []string                                `json:"regions"`
	IPSubnetworks        []string                                `json:"ipSubnetworks"`
	RequiredAccessLevels []string                                `json:"requiredAccessLevels"`
	Members              []string                                `json:"members"`
	Negate               *bool                                   `json:"negate"`
	DevicePolicy         *AccessLevelBasicConditionsDevicePolicy `json:"devicePolicy"`
}

type jsonAccessLevelBasicConditions AccessLevelBasicConditions

func (r *AccessLevelBasicConditions) UnmarshalJSON(data []byte) error {
	var res jsonAccessLevelBasicConditions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAccessLevelBasicConditions
	} else {

		r.Regions = res.Regions

		r.IPSubnetworks = res.IPSubnetworks

		r.RequiredAccessLevels = res.RequiredAccessLevels

		r.Members = res.Members

		r.Negate = res.Negate

		r.DevicePolicy = res.DevicePolicy

	}
	return nil
}

// This object is used to assert a desired state where this AccessLevelBasicConditions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAccessLevelBasicConditions *AccessLevelBasicConditions = &AccessLevelBasicConditions{empty: true}

func (r *AccessLevelBasicConditions) Empty() bool {
	return r.empty
}

func (r *AccessLevelBasicConditions) String() string {
	return dcl.SprintResource(r)
}

func (r *AccessLevelBasicConditions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AccessLevelBasicConditionsDevicePolicy struct {
	empty                         bool                                                                      `json:"-"`
	RequireScreenlock             *bool                                                                     `json:"requireScreenlock"`
	RequireAdminApproval          *bool                                                                     `json:"requireAdminApproval"`
	RequireCorpOwned              *bool                                                                     `json:"requireCorpOwned"`
	AllowedEncryptionStatuses     []AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum     `json:"allowedEncryptionStatuses"`
	AllowedDeviceManagementLevels []AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum `json:"allowedDeviceManagementLevels"`
	OsConstraints                 []AccessLevelBasicConditionsDevicePolicyOsConstraints                     `json:"osConstraints"`
}

type jsonAccessLevelBasicConditionsDevicePolicy AccessLevelBasicConditionsDevicePolicy

func (r *AccessLevelBasicConditionsDevicePolicy) UnmarshalJSON(data []byte) error {
	var res jsonAccessLevelBasicConditionsDevicePolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAccessLevelBasicConditionsDevicePolicy
	} else {

		r.RequireScreenlock = res.RequireScreenlock

		r.RequireAdminApproval = res.RequireAdminApproval

		r.RequireCorpOwned = res.RequireCorpOwned

		r.AllowedEncryptionStatuses = res.AllowedEncryptionStatuses

		r.AllowedDeviceManagementLevels = res.AllowedDeviceManagementLevels

		r.OsConstraints = res.OsConstraints

	}
	return nil
}

// This object is used to assert a desired state where this AccessLevelBasicConditionsDevicePolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAccessLevelBasicConditionsDevicePolicy *AccessLevelBasicConditionsDevicePolicy = &AccessLevelBasicConditionsDevicePolicy{empty: true}

func (r *AccessLevelBasicConditionsDevicePolicy) Empty() bool {
	return r.empty
}

func (r *AccessLevelBasicConditionsDevicePolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *AccessLevelBasicConditionsDevicePolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AccessLevelBasicConditionsDevicePolicyOsConstraints struct {
	empty                   bool                                                           `json:"-"`
	MinimumVersion          *string                                                        `json:"minimumVersion"`
	OsType                  *AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum `json:"osType"`
	RequireVerifiedChromeOs *bool                                                          `json:"requireVerifiedChromeOs"`
}

type jsonAccessLevelBasicConditionsDevicePolicyOsConstraints AccessLevelBasicConditionsDevicePolicyOsConstraints

func (r *AccessLevelBasicConditionsDevicePolicyOsConstraints) UnmarshalJSON(data []byte) error {
	var res jsonAccessLevelBasicConditionsDevicePolicyOsConstraints
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAccessLevelBasicConditionsDevicePolicyOsConstraints
	} else {

		r.MinimumVersion = res.MinimumVersion

		r.OsType = res.OsType

		r.RequireVerifiedChromeOs = res.RequireVerifiedChromeOs

	}
	return nil
}

// This object is used to assert a desired state where this AccessLevelBasicConditionsDevicePolicyOsConstraints is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAccessLevelBasicConditionsDevicePolicyOsConstraints *AccessLevelBasicConditionsDevicePolicyOsConstraints = &AccessLevelBasicConditionsDevicePolicyOsConstraints{empty: true}

func (r *AccessLevelBasicConditionsDevicePolicyOsConstraints) Empty() bool {
	return r.empty
}

func (r *AccessLevelBasicConditionsDevicePolicyOsConstraints) String() string {
	return dcl.SprintResource(r)
}

func (r *AccessLevelBasicConditionsDevicePolicyOsConstraints) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *AccessLevel) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "access_context_manager",
		Type:    "AccessLevel",
		Version: "accesscontextmanager",
	}
}

const AccessLevelMaxPage = -1

type AccessLevelList struct {
	Items []*AccessLevel

	nextToken string

	pageSize int32

	policy string
}

func (l *AccessLevelList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AccessLevelList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAccessLevel(ctx, l.policy, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAccessLevel(ctx context.Context, policy string) (*AccessLevelList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAccessLevelWithMaxResults(ctx, policy, AccessLevelMaxPage)

}

func (c *Client) ListAccessLevelWithMaxResults(ctx context.Context, policy string, pageSize int32) (*AccessLevelList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAccessLevel(ctx, policy, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AccessLevelList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		policy: policy,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *AccessLevel) URLNormalized() *AccessLevel {
	normalized := dcl.Copy(*r).(AccessLevel)
	normalized.Title = dcl.SelfLinkToName(r.Title)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Policy = dcl.SelfLinkToName(r.Policy)
	return &normalized
}
func (c *Client) GetAccessLevel(ctx context.Context, r *AccessLevel) (*AccessLevel, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAccessLevelRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAccessLevel(b, c)
	if err != nil {
		return nil, err
	}
	result.Policy = r.Policy
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAccessLevelNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAccessLevel(ctx context.Context, r *AccessLevel) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("AccessLevel resource is nil")
	}
	c.Config.Logger.Info("Deleting AccessLevel...")
	deleteOp := deleteAccessLevelOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAccessLevel deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAccessLevel(ctx context.Context, policy string, filter func(*AccessLevel) bool) error {
	listObj, err := c.ListAccessLevel(ctx, policy)
	if err != nil {
		return err
	}

	err = c.deleteAllAccessLevel(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAccessLevel(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAccessLevel(ctx context.Context, rawDesired *AccessLevel, opts ...dcl.ApplyOption) (*AccessLevel, error) {

	var resultNewState *AccessLevel
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAccessLevelHelper(c, ctx, rawDesired, opts...)
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

func applyAccessLevelHelper(c *Client, ctx context.Context, rawDesired *AccessLevel, opts ...dcl.ApplyOption) (*AccessLevel, error) {
	c.Config.Logger.Info("Beginning ApplyAccessLevel...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.accessLevelDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToAccessLevelOp(opStrings, fieldDiffs, opts)
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
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []accessLevelApiOperation
	if create {
		ops = append(ops, &createAccessLevelOperation{})
	} else if recreate {
		ops = append(ops, &deleteAccessLevelOperation{})
		ops = append(ops, &createAccessLevelOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAccessLevelDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetAccessLevel(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAccessLevelOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAccessLevel(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAccessLevelNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAccessLevelNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAccessLevelDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAccessLevel(c, newDesired, newState)
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
