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
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type InstanceTemplate struct {
	CreationTimestamp *string                     `json:"creationTimestamp"`
	Description       *string                     `json:"description"`
	Id                *int64                      `json:"id"`
	SelfLink          *string                     `json:"selfLink"`
	Name              *string                     `json:"name"`
	Properties        *InstanceTemplateProperties `json:"properties"`
	Project           *string                     `json:"project"`
}

func (r *InstanceTemplate) String() string {
	return dcl.SprintResource(r)
}

// The enum InstanceTemplatePropertiesDisksInterfaceEnum.
type InstanceTemplatePropertiesDisksInterfaceEnum string

// InstanceTemplatePropertiesDisksInterfaceEnumRef returns a *InstanceTemplatePropertiesDisksInterfaceEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceTemplatePropertiesDisksInterfaceEnumRef(s string) *InstanceTemplatePropertiesDisksInterfaceEnum {
	if s == "" {
		return nil
	}

	v := InstanceTemplatePropertiesDisksInterfaceEnum(s)
	return &v
}

func (v InstanceTemplatePropertiesDisksInterfaceEnum) Validate() error {
	for _, s := range []string{"SCSI", "NVME"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceTemplatePropertiesDisksInterfaceEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceTemplatePropertiesDisksModeEnum.
type InstanceTemplatePropertiesDisksModeEnum string

// InstanceTemplatePropertiesDisksModeEnumRef returns a *InstanceTemplatePropertiesDisksModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceTemplatePropertiesDisksModeEnumRef(s string) *InstanceTemplatePropertiesDisksModeEnum {
	if s == "" {
		return nil
	}

	v := InstanceTemplatePropertiesDisksModeEnum(s)
	return &v
}

func (v InstanceTemplatePropertiesDisksModeEnum) Validate() error {
	for _, s := range []string{"READ_WRITE", "READ_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceTemplatePropertiesDisksModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceTemplatePropertiesDisksTypeEnum.
type InstanceTemplatePropertiesDisksTypeEnum string

// InstanceTemplatePropertiesDisksTypeEnumRef returns a *InstanceTemplatePropertiesDisksTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceTemplatePropertiesDisksTypeEnumRef(s string) *InstanceTemplatePropertiesDisksTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceTemplatePropertiesDisksTypeEnum(s)
	return &v
}

func (v InstanceTemplatePropertiesDisksTypeEnum) Validate() error {
	for _, s := range []string{"SCRATCH", "PERSISTENT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceTemplatePropertiesDisksTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum.
type InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum string

// InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumRef returns a *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumRef(s string) *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(s)
	return &v
}

func (v InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum) Validate() error {
	for _, s := range []string{"ONE_TO_ONE_NAT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum.
type InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum string

// InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumRef returns a *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumRef(s string) *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum {
	if s == "" {
		return nil
	}

	v := InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(s)
	return &v
}

func (v InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum) Validate() error {
	for _, s := range []string{"PREMIUM", "STANDARD"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum.
type InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum string

// InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumRef returns a *InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumRef(s string) *InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum {
	if s == "" {
		return nil
	}

	v := InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(s)
	return &v
}

func (v InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum) Validate() error {
	for _, s := range []string{"IN", "NOT_IN"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InstanceTemplateProperties struct {
	empty                  bool                                              `json:"-"`
	CanIPForward           *bool                                             `json:"canIPForward"`
	Description            *string                                           `json:"description"`
	Disks                  []InstanceTemplatePropertiesDisks                 `json:"disks"`
	Labels                 map[string]string                                 `json:"labels"`
	MachineType            *string                                           `json:"machineType"`
	MinCpuPlatform         *string                                           `json:"minCpuPlatform"`
	Metadata               map[string]string                                 `json:"metadata"`
	ReservationAffinity    *InstanceTemplatePropertiesReservationAffinity    `json:"reservationAffinity"`
	GuestAccelerators      []InstanceTemplatePropertiesGuestAccelerators     `json:"guestAccelerators"`
	NetworkInterfaces      []InstanceTemplatePropertiesNetworkInterfaces     `json:"networkInterfaces"`
	ShieldedInstanceConfig *InstanceTemplatePropertiesShieldedInstanceConfig `json:"shieldedInstanceConfig"`
	Scheduling             *InstanceTemplatePropertiesScheduling             `json:"scheduling"`
	ServiceAccounts        []InstanceTemplatePropertiesServiceAccounts       `json:"serviceAccounts"`
	Tags                   []string                                          `json:"tags"`
}

// This object is used to assert a desired state where this InstanceTemplateProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplateProperties *InstanceTemplateProperties = &InstanceTemplateProperties{empty: true}

func (r *InstanceTemplateProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplateProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesDisks struct {
	empty             bool                                              `json:"-"`
	AutoDelete        *bool                                             `json:"autoDelete"`
	Boot              *bool                                             `json:"boot"`
	DeviceName        *string                                           `json:"deviceName"`
	DiskEncryptionKey *InstanceTemplatePropertiesDisksDiskEncryptionKey `json:"diskEncryptionKey"`
	Index             *int64                                            `json:"index"`
	InitializeParams  *InstanceTemplatePropertiesDisksInitializeParams  `json:"initializeParams"`
	GuestOsFeatures   []InstanceTemplatePropertiesDisksGuestOsFeatures  `json:"guestOsFeatures"`
	Interface         *InstanceTemplatePropertiesDisksInterfaceEnum     `json:"interface"`
	Mode              *InstanceTemplatePropertiesDisksModeEnum          `json:"mode"`
	Source            *string                                           `json:"source"`
	Type              *InstanceTemplatePropertiesDisksTypeEnum          `json:"type"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesDisks is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesDisks *InstanceTemplatePropertiesDisks = &InstanceTemplatePropertiesDisks{empty: true}

func (r *InstanceTemplatePropertiesDisks) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesDisks) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesDisksDiskEncryptionKey struct {
	empty           bool    `json:"-"`
	RawKey          *string `json:"rawKey"`
	RsaEncryptedKey *string `json:"rsaEncryptedKey"`
	Sha256          *string `json:"sha256"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesDisksDiskEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesDisksDiskEncryptionKey *InstanceTemplatePropertiesDisksDiskEncryptionKey = &InstanceTemplatePropertiesDisksDiskEncryptionKey{empty: true}

func (r *InstanceTemplatePropertiesDisksDiskEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesDisksDiskEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesDisksInitializeParams struct {
	empty                       bool                                                                        `json:"-"`
	DiskName                    *string                                                                     `json:"diskName"`
	DiskSizeGb                  *int64                                                                      `json:"diskSizeGb"`
	DiskType                    *string                                                                     `json:"diskType"`
	SourceImage                 *string                                                                     `json:"sourceImage"`
	Labels                      map[string]string                                                           `json:"labels"`
	SourceSnapshot              *string                                                                     `json:"sourceSnapshot"`
	SourceSnapshotEncryptionKey *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey"`
	Description                 *string                                                                     `json:"description"`
	ResourcePolicies            []string                                                                    `json:"resourcePolicies"`
	OnUpdateAction              *string                                                                     `json:"onUpdateAction"`
	SourceImageEncryptionKey    *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey    `json:"sourceImageEncryptionKey"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesDisksInitializeParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesDisksInitializeParams *InstanceTemplatePropertiesDisksInitializeParams = &InstanceTemplatePropertiesDisksInitializeParams{empty: true}

func (r *InstanceTemplatePropertiesDisksInitializeParams) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesDisksInitializeParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey struct {
	empty      bool    `json:"-"`
	RawKey     *string `json:"rawKey"`
	Sha256     *string `json:"sha256"`
	KmsKeyName *string `json:"kmsKeyName"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey = &InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{empty: true}

func (r *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey struct {
	empty      bool    `json:"-"`
	RawKey     *string `json:"rawKey"`
	Sha256     *string `json:"sha256"`
	KmsKeyName *string `json:"kmsKeyName"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey = &InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{empty: true}

func (r *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesDisksGuestOsFeatures struct {
	empty bool    `json:"-"`
	Type  *string `json:"type"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesDisksGuestOsFeatures is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesDisksGuestOsFeatures *InstanceTemplatePropertiesDisksGuestOsFeatures = &InstanceTemplatePropertiesDisksGuestOsFeatures{empty: true}

func (r *InstanceTemplatePropertiesDisksGuestOsFeatures) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesDisksGuestOsFeatures) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesReservationAffinity struct {
	empty bool     `json:"-"`
	Key   *string  `json:"key"`
	Value []string `json:"value"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesReservationAffinity is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesReservationAffinity *InstanceTemplatePropertiesReservationAffinity = &InstanceTemplatePropertiesReservationAffinity{empty: true}

func (r *InstanceTemplatePropertiesReservationAffinity) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesReservationAffinity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesGuestAccelerators struct {
	empty            bool    `json:"-"`
	AcceleratorCount *int64  `json:"acceleratorCount"`
	AcceleratorType  *string `json:"acceleratorType"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesGuestAccelerators is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesGuestAccelerators *InstanceTemplatePropertiesGuestAccelerators = &InstanceTemplatePropertiesGuestAccelerators{empty: true}

func (r *InstanceTemplatePropertiesGuestAccelerators) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesGuestAccelerators) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesNetworkInterfaces struct {
	empty         bool                                                       `json:"-"`
	AccessConfigs []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs `json:"accessConfigs"`
	AliasIPRanges []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges `json:"aliasIPRanges"`
	Name          *string                                                    `json:"name"`
	Network       *string                                                    `json:"network"`
	NetworkIP     *string                                                    `json:"networkIP"`
	Subnetwork    *string                                                    `json:"subnetwork"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesNetworkInterfaces is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesNetworkInterfaces *InstanceTemplatePropertiesNetworkInterfaces = &InstanceTemplatePropertiesNetworkInterfaces{empty: true}

func (r *InstanceTemplatePropertiesNetworkInterfaces) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesNetworkInterfaces) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesNetworkInterfacesAccessConfigs struct {
	empty               bool                                                                     `json:"-"`
	Name                *string                                                                  `json:"name"`
	NatIP               *string                                                                  `json:"natIP"`
	Type                *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum        `json:"type"`
	SetPublicPtr        *bool                                                                    `json:"setPublicPtr"`
	PublicPtrDomainName *string                                                                  `json:"publicPtrDomainName"`
	NetworkTier         *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum `json:"networkTier"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesNetworkInterfacesAccessConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesNetworkInterfacesAccessConfigs *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs = &InstanceTemplatePropertiesNetworkInterfacesAccessConfigs{empty: true}

func (r *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges struct {
	empty               bool    `json:"-"`
	IPCidrRange         *string `json:"ipCidrRange"`
	SubnetworkRangeName *string `json:"subnetworkRangeName"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges = &InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{empty: true}

func (r *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesShieldedInstanceConfig struct {
	empty                     bool  `json:"-"`
	EnableSecureBoot          *bool `json:"enableSecureBoot"`
	EnableVtpm                *bool `json:"enableVtpm"`
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesShieldedInstanceConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesShieldedInstanceConfig *InstanceTemplatePropertiesShieldedInstanceConfig = &InstanceTemplatePropertiesShieldedInstanceConfig{empty: true}

func (r *InstanceTemplatePropertiesShieldedInstanceConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesShieldedInstanceConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesScheduling struct {
	empty             bool                                                 `json:"-"`
	AutomaticRestart  *bool                                                `json:"automaticRestart"`
	OnHostMaintenance *string                                              `json:"onHostMaintenance"`
	Preemptible       *bool                                                `json:"preemptible"`
	NodeAffinities    []InstanceTemplatePropertiesSchedulingNodeAffinities `json:"nodeAffinities"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesScheduling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesScheduling *InstanceTemplatePropertiesScheduling = &InstanceTemplatePropertiesScheduling{empty: true}

func (r *InstanceTemplatePropertiesScheduling) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesScheduling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesSchedulingNodeAffinities struct {
	empty    bool                                                            `json:"-"`
	Key      *string                                                         `json:"key"`
	Operator *InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum `json:"operator"`
	Values   []string                                                        `json:"values"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesSchedulingNodeAffinities is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesSchedulingNodeAffinities *InstanceTemplatePropertiesSchedulingNodeAffinities = &InstanceTemplatePropertiesSchedulingNodeAffinities{empty: true}

func (r *InstanceTemplatePropertiesSchedulingNodeAffinities) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesSchedulingNodeAffinities) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceTemplatePropertiesServiceAccounts struct {
	empty  bool     `json:"-"`
	Email  *string  `json:"email"`
	Scopes []string `json:"scopes"`
}

// This object is used to assert a desired state where this InstanceTemplatePropertiesServiceAccounts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceTemplatePropertiesServiceAccounts *InstanceTemplatePropertiesServiceAccounts = &InstanceTemplatePropertiesServiceAccounts{empty: true}

func (r *InstanceTemplatePropertiesServiceAccounts) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceTemplatePropertiesServiceAccounts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *InstanceTemplate) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "InstanceTemplate",
		Version: "compute",
	}
}

const InstanceTemplateMaxPage = -1

type InstanceTemplateList struct {
	Items []*InstanceTemplate

	nextToken string

	pageSize int32

	project string
}

func (l *InstanceTemplateList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InstanceTemplateList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInstanceTemplate(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInstanceTemplate(ctx context.Context, project string) (*InstanceTemplateList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListInstanceTemplateWithMaxResults(ctx, project, InstanceTemplateMaxPage)

}

func (c *Client) ListInstanceTemplateWithMaxResults(ctx context.Context, project string, pageSize int32) (*InstanceTemplateList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listInstanceTemplate(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InstanceTemplateList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetInstanceTemplate(ctx context.Context, r *InstanceTemplate) (*InstanceTemplate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getInstanceTemplateRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalInstanceTemplate(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeInstanceTemplateNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteInstanceTemplate(ctx context.Context, r *InstanceTemplate) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("InstanceTemplate resource is nil")
	}
	c.Config.Logger.Info("Deleting InstanceTemplate...")
	deleteOp := deleteInstanceTemplateOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInstanceTemplate deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInstanceTemplate(ctx context.Context, project string, filter func(*InstanceTemplate) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListInstanceTemplate(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllInstanceTemplate(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllInstanceTemplate(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyInstanceTemplate(ctx context.Context, rawDesired *InstanceTemplate, opts ...dcl.ApplyOption) (*InstanceTemplate, error) {
	c.Config.Logger.Info("Beginning ApplyInstanceTemplate...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.instanceTemplateDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []instanceTemplateApiOperation
	if create {
		ops = append(ops, &createInstanceTemplateOperation{})
	} else if recreate {

		ops = append(ops, &deleteInstanceTemplateOperation{})

		ops = append(ops, &createInstanceTemplateOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeInstanceTemplateDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetInstanceTemplate(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createInstanceTemplateOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapInstanceTemplate(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeInstanceTemplateNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeInstanceTemplateNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeInstanceTemplateDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffInstanceTemplate(c, newDesired, newState)
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
