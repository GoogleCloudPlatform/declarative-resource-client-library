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
	"fmt"
	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Instance struct {
	Name                   *string                            `json:"name"`
	DisplayName            *string                            `json:"displayName"`
	Labels                 []InstanceLabels                   `json:"labels"`
	LocationId             *string                            `json:"locationId"`
	AlternativeLocationId  *string                            `json:"alternativeLocationId"`
	RedisVersion           *string                            `json:"redisVersion"`
	ReservedIPRange        *string                            `json:"reservedIPRange"`
	Host                   *string                            `json:"host"`
	Port                   *int64                             `json:"port"`
	CurrentLocationId      *string                            `json:"currentLocationId"`
	CreateTime             *string                            `json:"createTime"`
	State                  *InstanceStateEnum                 `json:"state"`
	StatusMessage          *string                            `json:"statusMessage"`
	RedisConfigs           []InstanceRedisConfigs             `json:"redisConfigs"`
	Tier                   *InstanceTierEnum                  `json:"tier"`
	MemorySizeGb           *int64                             `json:"memorySizeGb"`
	AuthorizedNetwork      *string                            `json:"authorizedNetwork"`
	PersistenceIamIdentity *string                            `json:"persistenceIamIdentity"`
	ConnectMode            *InstanceConnectModeEnum           `json:"connectMode"`
	AuthEnabled            *bool                              `json:"authEnabled"`
	ServerCaCerts          []InstanceServerCaCerts            `json:"serverCaCerts"`
	TransitEncryptionMode  *InstanceTransitEncryptionModeEnum `json:"transitEncryptionMode"`
	MaintenancePolicy      *InstanceMaintenancePolicy         `json:"maintenancePolicy"`
	MaintenanceSchedule    *InstanceMaintenanceSchedule       `json:"maintenanceSchedule"`
	Project                *string                            `json:"project"`
	Location               *string                            `json:"location"`
	Region                 *string                            `json:"region"`
}

func (r *Instance) String() string {
	return dcl.SprintResource(r)
}

// The enum InstanceStateEnum.
type InstanceStateEnum string

// InstanceStateEnumRef returns a *InstanceStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceStateEnumRef(s string) *InstanceStateEnum {
	if s == "" {
		return nil
	}

	v := InstanceStateEnum(s)
	return &v
}

func (v InstanceStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "CREATING", "READY", "UPDATING", "DELETING", "REPAIRING", "PERFORMING_MAINTENANCE", "IMPORTING", "FAILING_OVER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceTierEnum.
type InstanceTierEnum string

// InstanceTierEnumRef returns a *InstanceTierEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceTierEnumRef(s string) *InstanceTierEnum {
	if s == "" {
		return nil
	}

	v := InstanceTierEnum(s)
	return &v
}

func (v InstanceTierEnum) Validate() error {
	for _, s := range []string{"TIER_UNSPECIFIED", "BASIC", "STANDARD_HA"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceTierEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceConnectModeEnum.
type InstanceConnectModeEnum string

// InstanceConnectModeEnumRef returns a *InstanceConnectModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceConnectModeEnumRef(s string) *InstanceConnectModeEnum {
	if s == "" {
		return nil
	}

	v := InstanceConnectModeEnum(s)
	return &v
}

func (v InstanceConnectModeEnum) Validate() error {
	for _, s := range []string{"CONNECT_MODE_UNSPECIFIED", "DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceConnectModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceTransitEncryptionModeEnum.
type InstanceTransitEncryptionModeEnum string

// InstanceTransitEncryptionModeEnumRef returns a *InstanceTransitEncryptionModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceTransitEncryptionModeEnumRef(s string) *InstanceTransitEncryptionModeEnum {
	if s == "" {
		return nil
	}

	v := InstanceTransitEncryptionModeEnum(s)
	return &v
}

func (v InstanceTransitEncryptionModeEnum) Validate() error {
	for _, s := range []string{"TRANSIT_ENCRYPTION_MODE_UNSPECIFIED", "SERVER_AUTHENTICATION", "DISABLED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceTransitEncryptionModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum.
type InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum string

// InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumRef returns a *InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumRef(s string) *InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum {
	if s == "" {
		return nil
	}

	v := InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(s)
	return &v
}

func (v InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum) Validate() error {
	for _, s := range []string{"DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InstanceLabels struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceLabels *InstanceLabels = &InstanceLabels{empty: true}

func (r *InstanceLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRedisConfigs struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceRedisConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRedisConfigs *InstanceRedisConfigs = &InstanceRedisConfigs{empty: true}

func (r *InstanceRedisConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRedisConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceServerCaCerts struct {
	empty           bool    `json:"-"`
	SerialNumber    *string `json:"serialNumber"`
	Cert            *string `json:"cert"`
	CreateTime      *string `json:"createTime"`
	ExpireTime      *string `json:"expireTime"`
	Sha1Fingerprint *string `json:"sha1Fingerprint"`
}

// This object is used to assert a desired state where this InstanceServerCaCerts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceServerCaCerts *InstanceServerCaCerts = &InstanceServerCaCerts{empty: true}

func (r *InstanceServerCaCerts) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceServerCaCerts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceMaintenancePolicy struct {
	empty                   bool                                               `json:"-"`
	CreateTime              *string                                            `json:"createTime"`
	UpdateTime              *string                                            `json:"updateTime"`
	Description             *string                                            `json:"description"`
	WeeklyMaintenanceWindow []InstanceMaintenancePolicyWeeklyMaintenanceWindow `json:"weeklyMaintenanceWindow"`
}

// This object is used to assert a desired state where this InstanceMaintenancePolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceMaintenancePolicy *InstanceMaintenancePolicy = &InstanceMaintenancePolicy{empty: true}

func (r *InstanceMaintenancePolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceMaintenancePolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceMaintenancePolicyWeeklyMaintenanceWindow struct {
	empty     bool                                                       `json:"-"`
	Day       *InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum   `json:"day"`
	StartTime *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime `json:"startTime"`
	Duration  *InstanceMaintenancePolicyWeeklyMaintenanceWindowDuration  `json:"duration"`
}

// This object is used to assert a desired state where this InstanceMaintenancePolicyWeeklyMaintenanceWindow is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceMaintenancePolicyWeeklyMaintenanceWindow *InstanceMaintenancePolicyWeeklyMaintenanceWindow = &InstanceMaintenancePolicyWeeklyMaintenanceWindow{empty: true}

func (r *InstanceMaintenancePolicyWeeklyMaintenanceWindow) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceMaintenancePolicyWeeklyMaintenanceWindow) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime struct {
	empty   bool   `json:"-"`
	Hours   *int64 `json:"hours"`
	Minutes *int64 `json:"minutes"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime = &InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{empty: true}

func (r *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceMaintenancePolicyWeeklyMaintenanceWindowDuration struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this InstanceMaintenancePolicyWeeklyMaintenanceWindowDuration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceMaintenancePolicyWeeklyMaintenanceWindowDuration *InstanceMaintenancePolicyWeeklyMaintenanceWindowDuration = &InstanceMaintenancePolicyWeeklyMaintenanceWindowDuration{empty: true}

func (r *InstanceMaintenancePolicyWeeklyMaintenanceWindowDuration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceMaintenancePolicyWeeklyMaintenanceWindowDuration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceMaintenanceSchedule struct {
	empty         bool    `json:"-"`
	StartTime     *string `json:"startTime"`
	EndTime       *string `json:"endTime"`
	CanReschedule *bool   `json:"canReschedule"`
}

// This object is used to assert a desired state where this InstanceMaintenanceSchedule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceMaintenanceSchedule *InstanceMaintenanceSchedule = &InstanceMaintenanceSchedule{empty: true}

func (r *InstanceMaintenanceSchedule) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceMaintenanceSchedule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Instance) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "redis",
		Type:    "Instance",
		Version: "beta",
	}
}

const InstanceMaxPage = -1

type InstanceList struct {
	Items []*Instance

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *InstanceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InstanceList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInstance(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInstance(ctx context.Context, project, location string) (*InstanceList, error) {

	return c.ListInstanceWithMaxResults(ctx, project, location, InstanceMaxPage)

}

func (c *Client) ListInstanceWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*InstanceList, error) {
	items, token, err := c.listInstance(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InstanceList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetInstance(ctx context.Context, r *Instance) (*Instance, error) {
	b, err := c.getInstanceRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalInstance(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeInstanceNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteInstance(ctx context.Context, r *Instance) error {
	if r == nil {
		return fmt.Errorf("Instance resource is nil")
	}
	c.Config.Logger.Info("Deleting Instance...")
	deleteOp := deleteInstanceOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInstance deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInstance(ctx context.Context, project, location string, filter func(*Instance) bool) error {
	listObj, err := c.ListInstance(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllInstance(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllInstance(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyInstance(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (*Instance, error) {
	c.Config.Logger.Info("Beginning ApplyInstance...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.instanceDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []instanceApiOperation
	if create {
		ops = append(ops, &createInstanceOperation{})
	} else if recreate {
		ops = append(ops, &deleteInstanceOperation{})
		ops = append(ops, &createInstanceOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeInstanceDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetInstance(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeInstanceNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeInstanceDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffInstance(c, newDesired, newState)
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
