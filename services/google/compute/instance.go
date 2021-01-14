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
package compute

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Instance struct {
	CanIPForward           *bool                           `json:"canIPForward"`
	CpuPlatform            *string                         `json:"cpuPlatform"`
	CreationTimestamp      *string                         `json:"creationTimestamp"`
	DeletionProtection     *bool                           `json:"deletionProtection"`
	Description            *string                         `json:"description"`
	Disks                  []InstanceDisks                 `json:"disks"`
	GuestAccelerators      []InstanceGuestAccelerators     `json:"guestAccelerators"`
	Hostname               *string                         `json:"hostname"`
	Id                     *string                         `json:"id"`
	Labels                 map[string]string               `json:"labels"`
	Metadata               map[string]string               `json:"metadata"`
	MachineType            *string                         `json:"machineType"`
	MinCpuPlatform         *string                         `json:"minCpuPlatform"`
	Name                   *string                         `json:"name"`
	NetworkInterfaces      []InstanceNetworkInterfaces     `json:"networkInterfaces"`
	Scheduling             *InstanceScheduling             `json:"scheduling"`
	ServiceAccounts        []InstanceServiceAccounts       `json:"serviceAccounts"`
	ShieldedInstanceConfig *InstanceShieldedInstanceConfig `json:"shieldedInstanceConfig"`
	Status                 *InstanceStatusEnum             `json:"status"`
	StatusMessage          *string                         `json:"statusMessage"`
	Tags                   []string                        `json:"tags"`
	Zone                   *string                         `json:"zone"`
	Project                *string                         `json:"project"`
	SelfLink               *string                         `json:"selfLink"`
}

func (r *Instance) String() string {
	return dcl.SprintResource(r)
}

// The enum InstanceDisksInterfaceEnum.
type InstanceDisksInterfaceEnum string

// InstanceDisksInterfaceEnumRef returns a *InstanceDisksInterfaceEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceDisksInterfaceEnumRef(s string) *InstanceDisksInterfaceEnum {
	if s == "" {
		return nil
	}

	v := InstanceDisksInterfaceEnum(s)
	return &v
}

func (v InstanceDisksInterfaceEnum) Validate() error {
	for _, s := range []string{"SCSI", "NVME"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceDisksInterfaceEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceDisksModeEnum.
type InstanceDisksModeEnum string

// InstanceDisksModeEnumRef returns a *InstanceDisksModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceDisksModeEnumRef(s string) *InstanceDisksModeEnum {
	if s == "" {
		return nil
	}

	v := InstanceDisksModeEnum(s)
	return &v
}

func (v InstanceDisksModeEnum) Validate() error {
	for _, s := range []string{"READ_WRITE", "READ_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceDisksModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceDisksTypeEnum.
type InstanceDisksTypeEnum string

// InstanceDisksTypeEnumRef returns a *InstanceDisksTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceDisksTypeEnumRef(s string) *InstanceDisksTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceDisksTypeEnum(s)
	return &v
}

func (v InstanceDisksTypeEnum) Validate() error {
	for _, s := range []string{"SCRATCH", "PERSISTENT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceDisksTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceNetworkInterfacesAccessConfigsTypeEnum.
type InstanceNetworkInterfacesAccessConfigsTypeEnum string

// InstanceNetworkInterfacesAccessConfigsTypeEnumRef returns a *InstanceNetworkInterfacesAccessConfigsTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceNetworkInterfacesAccessConfigsTypeEnumRef(s string) *InstanceNetworkInterfacesAccessConfigsTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceNetworkInterfacesAccessConfigsTypeEnum(s)
	return &v
}

func (v InstanceNetworkInterfacesAccessConfigsTypeEnum) Validate() error {
	for _, s := range []string{"ONE_TO_ONE_NAT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceNetworkInterfacesAccessConfigsTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceStatusEnum.
type InstanceStatusEnum string

// InstanceStatusEnumRef returns a *InstanceStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceStatusEnumRef(s string) *InstanceStatusEnum {
	if s == "" {
		return nil
	}

	v := InstanceStatusEnum(s)
	return &v
}

func (v InstanceStatusEnum) Validate() error {
	for _, s := range []string{"PROVISIONING", "STAGING", "RUNNING", "STOPPING", "SUSPENDING", "SUSPENDED", "TERMINATED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InstanceDisks struct {
	empty             bool                            `json:"-"`
	AutoDelete        *bool                           `json:"autoDelete"`
	Boot              *bool                           `json:"boot"`
	DeviceName        *string                         `json:"deviceName"`
	DiskEncryptionKey *InstanceDisksDiskEncryptionKey `json:"diskEncryptionKey"`
	Index             *int64                          `json:"index"`
	InitializeParams  *InstanceDisksInitializeParams  `json:"initializeParams"`
	Interface         *InstanceDisksInterfaceEnum     `json:"interface"`
	Mode              *InstanceDisksModeEnum          `json:"mode"`
	Source            *string                         `json:"source"`
	Type              *InstanceDisksTypeEnum          `json:"type"`
}

// This object is used to assert a desired state where this InstanceDisks is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDisks *InstanceDisks = &InstanceDisks{empty: true}

func (r *InstanceDisks) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDisks) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDisksDiskEncryptionKey struct {
	empty           bool    `json:"-"`
	RawKey          *string `json:"rawKey"`
	RsaEncryptedKey *string `json:"rsaEncryptedKey"`
	Sha256          *string `json:"sha256"`
}

// This object is used to assert a desired state where this InstanceDisksDiskEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDisksDiskEncryptionKey *InstanceDisksDiskEncryptionKey = &InstanceDisksDiskEncryptionKey{empty: true}

func (r *InstanceDisksDiskEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDisksDiskEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDisksInitializeParams struct {
	empty                    bool                                                   `json:"-"`
	DiskName                 *string                                                `json:"diskName"`
	DiskSizeGb               *int64                                                 `json:"diskSizeGb"`
	DiskType                 *string                                                `json:"diskType"`
	SourceImage              *string                                                `json:"sourceImage"`
	SourceImageEncryptionKey *InstanceDisksInitializeParamsSourceImageEncryptionKey `json:"sourceImageEncryptionKey"`
}

// This object is used to assert a desired state where this InstanceDisksInitializeParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDisksInitializeParams *InstanceDisksInitializeParams = &InstanceDisksInitializeParams{empty: true}

func (r *InstanceDisksInitializeParams) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDisksInitializeParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDisksInitializeParamsSourceImageEncryptionKey struct {
	empty  bool    `json:"-"`
	RawKey *string `json:"rawKey"`
	Sha256 *string `json:"sha256"`
}

// This object is used to assert a desired state where this InstanceDisksInitializeParamsSourceImageEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDisksInitializeParamsSourceImageEncryptionKey *InstanceDisksInitializeParamsSourceImageEncryptionKey = &InstanceDisksInitializeParamsSourceImageEncryptionKey{empty: true}

func (r *InstanceDisksInitializeParamsSourceImageEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDisksInitializeParamsSourceImageEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGuestAccelerators struct {
	empty            bool    `json:"-"`
	AcceleratorCount *int64  `json:"acceleratorCount"`
	AcceleratorType  *string `json:"acceleratorType"`
}

// This object is used to assert a desired state where this InstanceGuestAccelerators is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGuestAccelerators *InstanceGuestAccelerators = &InstanceGuestAccelerators{empty: true}

func (r *InstanceGuestAccelerators) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGuestAccelerators) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNetworkInterfaces struct {
	empty         bool                                     `json:"-"`
	AccessConfigs []InstanceNetworkInterfacesAccessConfigs `json:"accessConfigs"`
	AliasIPRanges []InstanceNetworkInterfacesAliasIPRanges `json:"aliasIPRanges"`
	Name          *string                                  `json:"name"`
	Network       *string                                  `json:"network"`
	NetworkIP     *string                                  `json:"networkIP"`
	Subnetwork    *string                                  `json:"subnetwork"`
}

// This object is used to assert a desired state where this InstanceNetworkInterfaces is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNetworkInterfaces *InstanceNetworkInterfaces = &InstanceNetworkInterfaces{empty: true}

func (r *InstanceNetworkInterfaces) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNetworkInterfaces) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNetworkInterfacesAccessConfigs struct {
	empty bool                                            `json:"-"`
	Name  *string                                         `json:"name"`
	NatIP *string                                         `json:"natIP"`
	Type  *InstanceNetworkInterfacesAccessConfigsTypeEnum `json:"type"`
}

// This object is used to assert a desired state where this InstanceNetworkInterfacesAccessConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNetworkInterfacesAccessConfigs *InstanceNetworkInterfacesAccessConfigs = &InstanceNetworkInterfacesAccessConfigs{empty: true}

func (r *InstanceNetworkInterfacesAccessConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNetworkInterfacesAccessConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNetworkInterfacesAliasIPRanges struct {
	empty               bool    `json:"-"`
	IPCidrRange         *string `json:"ipCidrRange"`
	SubnetworkRangeName *string `json:"subnetworkRangeName"`
}

// This object is used to assert a desired state where this InstanceNetworkInterfacesAliasIPRanges is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNetworkInterfacesAliasIPRanges *InstanceNetworkInterfacesAliasIPRanges = &InstanceNetworkInterfacesAliasIPRanges{empty: true}

func (r *InstanceNetworkInterfacesAliasIPRanges) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNetworkInterfacesAliasIPRanges) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceScheduling struct {
	empty             bool    `json:"-"`
	AutomaticRestart  *bool   `json:"automaticRestart"`
	OnHostMaintenance *string `json:"onHostMaintenance"`
	Preemptible       *bool   `json:"preemptible"`
}

// This object is used to assert a desired state where this InstanceScheduling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceScheduling *InstanceScheduling = &InstanceScheduling{empty: true}

func (r *InstanceScheduling) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceScheduling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceServiceAccounts struct {
	empty  bool     `json:"-"`
	Email  *string  `json:"email"`
	Scopes []string `json:"scopes"`
}

// This object is used to assert a desired state where this InstanceServiceAccounts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceServiceAccounts *InstanceServiceAccounts = &InstanceServiceAccounts{empty: true}

func (r *InstanceServiceAccounts) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceServiceAccounts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceShieldedInstanceConfig struct {
	empty                     bool  `json:"-"`
	EnableSecureBoot          *bool `json:"enableSecureBoot"`
	EnableVtpm                *bool `json:"enableVtpm"`
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring"`
}

// This object is used to assert a desired state where this InstanceShieldedInstanceConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceShieldedInstanceConfig *InstanceShieldedInstanceConfig = &InstanceShieldedInstanceConfig{empty: true}

func (r *InstanceShieldedInstanceConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceShieldedInstanceConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Instance) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Instance",
		Version: "compute",
	}
}

const InstanceMaxPage = -1

type InstanceList struct {
	Items []*Instance

	nextToken string

	pageSize int32

	project string

	zone string
}

func (l *InstanceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InstanceList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInstance(ctx, l.project, l.zone, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInstance(ctx context.Context, project, zone string) (*InstanceList, error) {

	return c.ListInstanceWithMaxResults(ctx, project, zone, InstanceMaxPage)

}

func (c *Client) ListInstanceWithMaxResults(ctx context.Context, project, zone string, pageSize int32) (*InstanceList, error) {
	items, token, err := c.listInstance(ctx, project, zone, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InstanceList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		zone: zone,
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
	result.Zone = r.Zone
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
func (c *Client) DeleteAllInstance(ctx context.Context, project, zone string, filter func(*Instance) bool) error {
	listObj, err := c.ListInstance(ctx, project, zone)
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
func (r *Instance) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
