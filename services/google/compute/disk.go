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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Disk struct {
	SelfLink                    *string               `json:"selfLink"`
	Description                 *string               `json:"description"`
	DiskEncryptionKey           *DiskEncryptionKey    `json:"diskEncryptionKey"`
	GuestOsFeature              []DiskGuestOsFeature  `json:"guestOsFeature"`
	Labels                      map[string]string     `json:"labels"`
	LabelFingerprint            *string               `json:"labelFingerprint"`
	License                     []string              `json:"license"`
	Name                        *string               `json:"name"`
	Region                      *string               `json:"region"`
	ReplicaZones                []string              `json:"replicaZones"`
	ResourcePolicy              []string              `json:"resourcePolicy"`
	SizeGb                      *int64                `json:"sizeGb"`
	SourceImage                 *string               `json:"sourceImage"`
	SourceImageEncryptionKey    *DiskEncryptionKey    `json:"sourceImageEncryptionKey"`
	SourceImageId               *string               `json:"sourceImageId"`
	SourceSnapshot              *string               `json:"sourceSnapshot"`
	SourceSnapshotEncryptionKey *DiskEncryptionKey    `json:"sourceSnapshotEncryptionKey"`
	SourceSnapshotId            *string               `json:"sourceSnapshotId"`
	Type                        *string               `json:"type"`
	Zone                        *string               `json:"zone"`
	Project                     *string               `json:"project"`
	Id                          *int64                `json:"id"`
	Status                      *DiskStatusEnum       `json:"status"`
	Options                     *string               `json:"options"`
	Licenses                    []string              `json:"licenses"`
	GuestOsFeatures             []DiskGuestOsFeatures `json:"guestOsFeatures"`
	LastAttachTimestamp         *string               `json:"lastAttachTimestamp"`
	LastDetachTimestamp         *string               `json:"lastDetachTimestamp"`
	Users                       []string              `json:"users"`
	LicenseCodes                []int64               `json:"licenseCodes"`
	PhysicalBlockSizeBytes      *int64                `json:"physicalBlockSizeBytes"`
	ResourcePolicies            []string              `json:"resourcePolicies"`
	SourceDisk                  *string               `json:"sourceDisk"`
	SourceDiskId                *string               `json:"sourceDiskId"`
	Location                    *string               `json:"location"`
}

func (r *Disk) String() string {
	return dcl.SprintResource(r)
}

// The enum DiskGuestOsFeatureTypeEnum.
type DiskGuestOsFeatureTypeEnum string

// DiskGuestOsFeatureTypeEnumRef returns a *DiskGuestOsFeatureTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func DiskGuestOsFeatureTypeEnumRef(s string) *DiskGuestOsFeatureTypeEnum {
	if s == "" {
		return nil
	}

	v := DiskGuestOsFeatureTypeEnum(s)
	return &v
}

func (v DiskGuestOsFeatureTypeEnum) Validate() error {
	for _, s := range []string{"FEATURE_TYPE_UNSPECIFIED", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "MULTI_IP_SUBNET", "UEFI_COMPATIBLE", "SECURE_BOOT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DiskGuestOsFeatureTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DiskGuestOsFeatureTypeAltEnum.
type DiskGuestOsFeatureTypeAltEnum string

// DiskGuestOsFeatureTypeAltEnumRef returns a *DiskGuestOsFeatureTypeAltEnum with the value of string s
// If the empty string is provided, nil is returned.
func DiskGuestOsFeatureTypeAltEnumRef(s string) *DiskGuestOsFeatureTypeAltEnum {
	if s == "" {
		return nil
	}

	v := DiskGuestOsFeatureTypeAltEnum(s)
	return &v
}

func (v DiskGuestOsFeatureTypeAltEnum) Validate() error {
	for _, s := range []string{"FEATURE_TYPE_UNSPECIFIED", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "MULTI_IP_SUBNET", "UEFI_COMPATIBLE", "SECURE_BOOT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DiskGuestOsFeatureTypeAltEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DiskStatusEnum.
type DiskStatusEnum string

// DiskStatusEnumRef returns a *DiskStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func DiskStatusEnumRef(s string) *DiskStatusEnum {
	if s == "" {
		return nil
	}

	v := DiskStatusEnum(s)
	return &v
}

func (v DiskStatusEnum) Validate() error {
	for _, s := range []string{"PENDING", "RUNNING", "DONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DiskStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DiskGuestOsFeaturesTypeEnum.
type DiskGuestOsFeaturesTypeEnum string

// DiskGuestOsFeaturesTypeEnumRef returns a *DiskGuestOsFeaturesTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func DiskGuestOsFeaturesTypeEnumRef(s string) *DiskGuestOsFeaturesTypeEnum {
	if s == "" {
		return nil
	}

	v := DiskGuestOsFeaturesTypeEnum(s)
	return &v
}

func (v DiskGuestOsFeaturesTypeEnum) Validate() error {
	for _, s := range []string{"PATH", "OTHER", "PARAMETER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DiskGuestOsFeaturesTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DiskGuestOsFeaturesTypeAltsEnum.
type DiskGuestOsFeaturesTypeAltsEnum string

// DiskGuestOsFeaturesTypeAltsEnumRef returns a *DiskGuestOsFeaturesTypeAltsEnum with the value of string s
// If the empty string is provided, nil is returned.
func DiskGuestOsFeaturesTypeAltsEnumRef(s string) *DiskGuestOsFeaturesTypeAltsEnum {
	if s == "" {
		return nil
	}

	v := DiskGuestOsFeaturesTypeAltsEnum(s)
	return &v
}

func (v DiskGuestOsFeaturesTypeAltsEnum) Validate() error {
	for _, s := range []string{"FEATURE_TYPE_UNSPECIFIED", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "MULTI_IP_SUBNET", "UEFI_COMPATIBLE", "SECURE_BOOT", "SEV_CAPABLE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DiskGuestOsFeaturesTypeAltsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type DiskGuestOsFeature struct {
	empty   bool                            `json:"-"`
	Type    *DiskGuestOsFeatureTypeEnum     `json:"type"`
	TypeAlt []DiskGuestOsFeatureTypeAltEnum `json:"typeAlt"`
}

type jsonDiskGuestOsFeature DiskGuestOsFeature

func (r *DiskGuestOsFeature) UnmarshalJSON(data []byte) error {
	var res jsonDiskGuestOsFeature
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDiskGuestOsFeature
	} else {

		r.Type = res.Type

		r.TypeAlt = res.TypeAlt

	}
	return nil
}

// This object is used to assert a desired state where this DiskGuestOsFeature is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDiskGuestOsFeature *DiskGuestOsFeature = &DiskGuestOsFeature{empty: true}

func (r *DiskGuestOsFeature) Empty() bool {
	return r.empty
}

func (r *DiskGuestOsFeature) String() string {
	return dcl.SprintResource(r)
}

func (r *DiskGuestOsFeature) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DiskEncryptionKey struct {
	empty                bool    `json:"-"`
	RawKey               *string `json:"rawKey"`
	KmsKeyName           *string `json:"kmsKeyName"`
	Sha256               *string `json:"sha256"`
	KmsKeyServiceAccount *string `json:"kmsKeyServiceAccount"`
}

type jsonDiskEncryptionKey DiskEncryptionKey

func (r *DiskEncryptionKey) UnmarshalJSON(data []byte) error {
	var res jsonDiskEncryptionKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDiskEncryptionKey
	} else {

		r.RawKey = res.RawKey

		r.KmsKeyName = res.KmsKeyName

		r.Sha256 = res.Sha256

		r.KmsKeyServiceAccount = res.KmsKeyServiceAccount

	}
	return nil
}

// This object is used to assert a desired state where this DiskEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDiskEncryptionKey *DiskEncryptionKey = &DiskEncryptionKey{empty: true}

func (r *DiskEncryptionKey) Empty() bool {
	return r.empty
}

func (r *DiskEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *DiskEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DiskGuestOsFeatures struct {
	empty    bool                              `json:"-"`
	Type     *DiskGuestOsFeaturesTypeEnum      `json:"type"`
	TypeAlts []DiskGuestOsFeaturesTypeAltsEnum `json:"typeAlts"`
}

type jsonDiskGuestOsFeatures DiskGuestOsFeatures

func (r *DiskGuestOsFeatures) UnmarshalJSON(data []byte) error {
	var res jsonDiskGuestOsFeatures
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDiskGuestOsFeatures
	} else {

		r.Type = res.Type

		r.TypeAlts = res.TypeAlts

	}
	return nil
}

// This object is used to assert a desired state where this DiskGuestOsFeatures is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDiskGuestOsFeatures *DiskGuestOsFeatures = &DiskGuestOsFeatures{empty: true}

func (r *DiskGuestOsFeatures) Empty() bool {
	return r.empty
}

func (r *DiskGuestOsFeatures) String() string {
	return dcl.SprintResource(r)
}

func (r *DiskGuestOsFeatures) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Disk) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Disk",
		Version: "compute",
	}
}

const DiskMaxPage = -1

type DiskList struct {
	Items []*Disk

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *DiskList) HasNext() bool {
	return l.nextToken != ""
}

func (l *DiskList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listDisk(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListDisk(ctx context.Context, project, location string) (*DiskList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListDiskWithMaxResults(ctx, project, location, DiskMaxPage)

}

func (c *Client) ListDiskWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*DiskList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listDisk(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &DiskList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Disk) URLNormalized() *Disk {
	normalized := dcl.Copy(*r).(Disk)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.LabelFingerprint = dcl.SelfLinkToName(r.LabelFingerprint)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.SourceImage = dcl.SelfLinkToName(r.SourceImage)
	normalized.SourceImageId = dcl.SelfLinkToName(r.SourceImageId)
	normalized.SourceSnapshot = dcl.SelfLinkToName(r.SourceSnapshot)
	normalized.SourceSnapshotId = dcl.SelfLinkToName(r.SourceSnapshotId)
	normalized.Type = dcl.SelfLinkToName(r.Type)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Options = dcl.SelfLinkToName(r.Options)
	normalized.LastAttachTimestamp = dcl.SelfLinkToName(r.LastAttachTimestamp)
	normalized.LastDetachTimestamp = dcl.SelfLinkToName(r.LastDetachTimestamp)
	normalized.SourceDisk = dcl.SelfLinkToName(r.SourceDisk)
	normalized.SourceDiskId = dcl.SelfLinkToName(r.SourceDiskId)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetDisk(ctx context.Context, r *Disk) (*Disk, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getDiskRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalDisk(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeDiskNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteDisk(ctx context.Context, r *Disk) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Disk resource is nil")
	}
	c.Config.Logger.Info("Deleting Disk...")
	deleteOp := deleteDiskOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllDisk deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllDisk(ctx context.Context, project, location string, filter func(*Disk) bool) error {
	listObj, err := c.ListDisk(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllDisk(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllDisk(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyDisk(ctx context.Context, rawDesired *Disk, opts ...dcl.ApplyOption) (*Disk, error) {

	var resultNewState *Disk
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyDiskHelper(c, ctx, rawDesired, opts...)
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

func applyDiskHelper(c *Client, ctx context.Context, rawDesired *Disk, opts ...dcl.ApplyOption) (*Disk, error) {
	c.Config.Logger.Info("Beginning ApplyDisk...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.diskDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToDiskDiffs(c.Config, fieldDiffs, opts)
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
	var ops []diskApiOperation
	if create {
		ops = append(ops, &createDiskOperation{})
	} else if recreate {
		ops = append(ops, &deleteDiskOperation{})
		ops = append(ops, &createDiskOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeDiskDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetDisk(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createDiskOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapDisk(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeDiskNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeDiskNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeDiskDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffDisk(c, newDesired, newState)
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
