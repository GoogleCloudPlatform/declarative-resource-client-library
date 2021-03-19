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
	BackendType                 *InstanceBackendTypeEnum             `json:"backendType"`
	ConnectionName              *string                              `json:"connectionName"`
	DatabaseVersion             *InstanceDatabaseVersionEnum         `json:"databaseVersion"`
	Etag                        *string                              `json:"etag"`
	GceZone                     *string                              `json:"gceZone"`
	InstanceType                *InstanceInstanceTypeEnum            `json:"instanceType"`
	MasterInstanceName          *string                              `json:"masterInstanceName"`
	MaxDiskSize                 *InstanceMaxDiskSize                 `json:"maxDiskSize"`
	Name                        *string                              `json:"name"`
	Project                     *string                              `json:"project"`
	Region                      *string                              `json:"region"`
	RootPassword                *string                              `json:"rootPassword"`
	CurrentDiskSize             *InstanceCurrentDiskSize             `json:"currentDiskSize"`
	DiskEncryptionConfiguration *InstanceDiskEncryptionConfiguration `json:"diskEncryptionConfiguration"`
	FailoverReplica             *InstanceFailoverReplica             `json:"failoverReplica"`
	IPAddresses                 []InstanceIPAddresses                `json:"ipAddresses"`
	MasterInstance              *InstanceMasterInstance              `json:"masterInstance"`
	ReplicaConfiguration        *InstanceReplicaConfiguration        `json:"replicaConfiguration"`
	ScheduledMaintenance        *InstanceScheduledMaintenance        `json:"scheduledMaintenance"`
	Settings                    *InstanceSettings                    `json:"settings"`
	State                       *string                              `json:"state"`
	ReplicaInstances            []InstanceReplicaInstances           `json:"replicaInstances"`
	ServerCaCert                *InstanceServerCaCert                `json:"serverCaCert"`
	IPv6Address                 *string                              `json:"ipv6Address"`
	ServiceAccountEmailAddress  *string                              `json:"serviceAccountEmailAddress"`
	OnPremisesConfiguration     *InstanceOnPremisesConfiguration     `json:"onPremisesConfiguration"`
	SuspensionReason            []string                             `json:"suspensionReason"`
	DiskEncryptionStatus        *InstanceDiskEncryptionStatus        `json:"diskEncryptionStatus"`
	InstanceUid                 *string                              `json:"instanceUid"`
}

func (r *Instance) String() string {
	return dcl.SprintResource(r)
}

// The enum InstanceBackendTypeEnum.
type InstanceBackendTypeEnum string

// InstanceBackendTypeEnumRef returns a *InstanceBackendTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceBackendTypeEnumRef(s string) *InstanceBackendTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceBackendTypeEnum(s)
	return &v
}

func (v InstanceBackendTypeEnum) Validate() error {
	for _, s := range []string{"SQL_BACKEND_TYPE_UNSPECIFIED", "FIRST_GEN", "SECOND_GEN", "EXTERNAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceBackendTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceDatabaseVersionEnum.
type InstanceDatabaseVersionEnum string

// InstanceDatabaseVersionEnumRef returns a *InstanceDatabaseVersionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceDatabaseVersionEnumRef(s string) *InstanceDatabaseVersionEnum {
	if s == "" {
		return nil
	}

	v := InstanceDatabaseVersionEnum(s)
	return &v
}

func (v InstanceDatabaseVersionEnum) Validate() error {
	for _, s := range []string{"SQL_DATABASE_VERSION_UNSPECIFIED", "MYSQL_5_6", "MYSQL_5_7", "MYSQL_8_0", "SQLSERVER_2017_STANDARD", "SQLSERVER_2017_ENTERPRISE", "SQLSERVER_2017_EXPRESS", "SQLSERVER_2017_WEB", "POSTGRES_9_6", "POSTGRES_10", "POSTGRES_11", "POSTGRES_12"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceDatabaseVersionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceInstanceTypeEnum.
type InstanceInstanceTypeEnum string

// InstanceInstanceTypeEnumRef returns a *InstanceInstanceTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInstanceTypeEnumRef(s string) *InstanceInstanceTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceInstanceTypeEnum(s)
	return &v
}

func (v InstanceInstanceTypeEnum) Validate() error {
	for _, s := range []string{"SQL_INSTANCE_TYPE_UNSPECIFIED", "CLOUD_SQL_INSTANCE", "ON_PREMISES_INSTANCE", "READ_REPLICA_INSTANCE", "READ_REPLICA_POOL_INSTANCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInstanceTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceIPAddressesTypeEnum.
type InstanceIPAddressesTypeEnum string

// InstanceIPAddressesTypeEnumRef returns a *InstanceIPAddressesTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceIPAddressesTypeEnumRef(s string) *InstanceIPAddressesTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceIPAddressesTypeEnum(s)
	return &v
}

func (v InstanceIPAddressesTypeEnum) Validate() error {
	for _, s := range []string{"SQL_IP_ADDRESS_TYPE_UNSPECIFIED", "PRIMARY", "OUTGOING", "PRIVATE", "MIGRATED_1ST_GEN"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceIPAddressesTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceSettingsAvailabilityTypeEnum.
type InstanceSettingsAvailabilityTypeEnum string

// InstanceSettingsAvailabilityTypeEnumRef returns a *InstanceSettingsAvailabilityTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceSettingsAvailabilityTypeEnumRef(s string) *InstanceSettingsAvailabilityTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceSettingsAvailabilityTypeEnum(s)
	return &v
}

func (v InstanceSettingsAvailabilityTypeEnum) Validate() error {
	for _, s := range []string{"SQL_AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceSettingsAvailabilityTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceSettingsPricingPlanEnum.
type InstanceSettingsPricingPlanEnum string

// InstanceSettingsPricingPlanEnumRef returns a *InstanceSettingsPricingPlanEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceSettingsPricingPlanEnumRef(s string) *InstanceSettingsPricingPlanEnum {
	if s == "" {
		return nil
	}

	v := InstanceSettingsPricingPlanEnum(s)
	return &v
}

func (v InstanceSettingsPricingPlanEnum) Validate() error {
	for _, s := range []string{"SQL_PRICING_PLAN_UNSPECIFIED", "PACKAGE", "PER_USE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceSettingsPricingPlanEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceSettingsReplicationTypeEnum.
type InstanceSettingsReplicationTypeEnum string

// InstanceSettingsReplicationTypeEnumRef returns a *InstanceSettingsReplicationTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceSettingsReplicationTypeEnumRef(s string) *InstanceSettingsReplicationTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceSettingsReplicationTypeEnum(s)
	return &v
}

func (v InstanceSettingsReplicationTypeEnum) Validate() error {
	for _, s := range []string{"SQL_REPLICATION_TYPE_UNSPECIFIED", "SYNCHRONOUS", "ASYNCHRONOUS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceSettingsReplicationTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceSettingsActivationPolicyEnum.
type InstanceSettingsActivationPolicyEnum string

// InstanceSettingsActivationPolicyEnumRef returns a *InstanceSettingsActivationPolicyEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceSettingsActivationPolicyEnumRef(s string) *InstanceSettingsActivationPolicyEnum {
	if s == "" {
		return nil
	}

	v := InstanceSettingsActivationPolicyEnum(s)
	return &v
}

func (v InstanceSettingsActivationPolicyEnum) Validate() error {
	for _, s := range []string{"SQL_ACTIVATION_POLICY_UNSPECIFIED", "ALWAYS", "NEVER", "ON_DEMAND"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceSettingsActivationPolicyEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceSettingsDataDiskTypeEnum.
type InstanceSettingsDataDiskTypeEnum string

// InstanceSettingsDataDiskTypeEnumRef returns a *InstanceSettingsDataDiskTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceSettingsDataDiskTypeEnumRef(s string) *InstanceSettingsDataDiskTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceSettingsDataDiskTypeEnum(s)
	return &v
}

func (v InstanceSettingsDataDiskTypeEnum) Validate() error {
	for _, s := range []string{"SQL_DATA_DISK_TYPE_UNSPECIFIED", "PD_SSD", "PD_HDD", "OBSOLETE_LOCAL_SSD"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceSettingsDataDiskTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceSettingsMaintenanceWindowUpdateTrackEnum.
type InstanceSettingsMaintenanceWindowUpdateTrackEnum string

// InstanceSettingsMaintenanceWindowUpdateTrackEnumRef returns a *InstanceSettingsMaintenanceWindowUpdateTrackEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceSettingsMaintenanceWindowUpdateTrackEnumRef(s string) *InstanceSettingsMaintenanceWindowUpdateTrackEnum {
	if s == "" {
		return nil
	}

	v := InstanceSettingsMaintenanceWindowUpdateTrackEnum(s)
	return &v
}

func (v InstanceSettingsMaintenanceWindowUpdateTrackEnum) Validate() error {
	for _, s := range []string{"SQL_UPDATE_TRACK_UNSPECIFIED", "canary", "stable"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceSettingsMaintenanceWindowUpdateTrackEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum.
type InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum string

// InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumRef returns a *InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumRef(s string) *InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum {
	if s == "" {
		return nil
	}

	v := InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(s)
	return &v
}

func (v InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum) Validate() error {
	for _, s := range []string{"RETENTION_UNIT_UNSPECIFIED", "COUNT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InstanceMaxDiskSize struct {
	empty bool   `json:"-"`
	Value *int64 `json:"value"`
}

// This object is used to assert a desired state where this InstanceMaxDiskSize is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceMaxDiskSize *InstanceMaxDiskSize = &InstanceMaxDiskSize{empty: true}

func (r *InstanceMaxDiskSize) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceMaxDiskSize) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCurrentDiskSize struct {
	empty bool   `json:"-"`
	Value *int64 `json:"value"`
}

// This object is used to assert a desired state where this InstanceCurrentDiskSize is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCurrentDiskSize *InstanceCurrentDiskSize = &InstanceCurrentDiskSize{empty: true}

func (r *InstanceCurrentDiskSize) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCurrentDiskSize) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDiskEncryptionConfiguration struct {
	empty      bool    `json:"-"`
	KmsKeyName *string `json:"kmsKeyName"`
	Kind       *string `json:"kind"`
}

// This object is used to assert a desired state where this InstanceDiskEncryptionConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDiskEncryptionConfiguration *InstanceDiskEncryptionConfiguration = &InstanceDiskEncryptionConfiguration{empty: true}

func (r *InstanceDiskEncryptionConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDiskEncryptionConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFailoverReplica struct {
	empty            bool                                     `json:"-"`
	Name             *string                                  `json:"name"`
	Available        *bool                                    `json:"available"`
	FailoverInstance *InstanceFailoverReplicaFailoverInstance `json:"failoverInstance"`
}

// This object is used to assert a desired state where this InstanceFailoverReplica is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFailoverReplica *InstanceFailoverReplica = &InstanceFailoverReplica{empty: true}

func (r *InstanceFailoverReplica) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFailoverReplica) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFailoverReplicaFailoverInstance struct {
	empty  bool    `json:"-"`
	Name   *string `json:"name"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceFailoverReplicaFailoverInstance is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFailoverReplicaFailoverInstance *InstanceFailoverReplicaFailoverInstance = &InstanceFailoverReplicaFailoverInstance{empty: true}

func (r *InstanceFailoverReplicaFailoverInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFailoverReplicaFailoverInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceIPAddresses struct {
	empty        bool                             `json:"-"`
	Type         *InstanceIPAddressesTypeEnum     `json:"type"`
	IPAddress    *string                          `json:"ipAddress"`
	TimeToRetire *InstanceIPAddressesTimeToRetire `json:"timeToRetire"`
}

// This object is used to assert a desired state where this InstanceIPAddresses is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceIPAddresses *InstanceIPAddresses = &InstanceIPAddresses{empty: true}

func (r *InstanceIPAddresses) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceIPAddresses) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceIPAddressesTimeToRetire struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this InstanceIPAddressesTimeToRetire is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceIPAddressesTimeToRetire *InstanceIPAddressesTimeToRetire = &InstanceIPAddressesTimeToRetire{empty: true}

func (r *InstanceIPAddressesTimeToRetire) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceIPAddressesTimeToRetire) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceMasterInstance struct {
	empty  bool    `json:"-"`
	Name   *string `json:"name"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceMasterInstance is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceMasterInstance *InstanceMasterInstance = &InstanceMasterInstance{empty: true}

func (r *InstanceMasterInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceMasterInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReplicaConfiguration struct {
	empty                     bool                                                   `json:"-"`
	Kind                      *string                                                `json:"kind"`
	MysqlReplicaConfiguration *InstanceReplicaConfigurationMysqlReplicaConfiguration `json:"mysqlReplicaConfiguration"`
	FailoverTarget            *bool                                                  `json:"failoverTarget"`
	ReplicaPoolConfiguration  *InstanceReplicaConfigurationReplicaPoolConfiguration  `json:"replicaPoolConfiguration"`
}

// This object is used to assert a desired state where this InstanceReplicaConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReplicaConfiguration *InstanceReplicaConfiguration = &InstanceReplicaConfiguration{empty: true}

func (r *InstanceReplicaConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReplicaConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReplicaConfigurationMysqlReplicaConfiguration struct {
	empty                   bool                                                                        `json:"-"`
	DumpFilePath            *string                                                                     `json:"dumpFilePath"`
	Username                *string                                                                     `json:"username"`
	Password                *string                                                                     `json:"password"`
	ConnectRetryInterval    *int64                                                                      `json:"connectRetryInterval"`
	MasterHeartbeatPeriod   *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod `json:"masterHeartbeatPeriod"`
	CaCertificate           *string                                                                     `json:"caCertificate"`
	ClientCertificate       *string                                                                     `json:"clientCertificate"`
	ClientKey               *string                                                                     `json:"clientKey"`
	SslCipher               *string                                                                     `json:"sslCipher"`
	VerifyServerCertificate *bool                                                                       `json:"verifyServerCertificate"`
	Kind                    *string                                                                     `json:"kind"`
}

// This object is used to assert a desired state where this InstanceReplicaConfigurationMysqlReplicaConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReplicaConfigurationMysqlReplicaConfiguration *InstanceReplicaConfigurationMysqlReplicaConfiguration = &InstanceReplicaConfigurationMysqlReplicaConfiguration{empty: true}

func (r *InstanceReplicaConfigurationMysqlReplicaConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReplicaConfigurationMysqlReplicaConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod struct {
	empty bool   `json:"-"`
	Value *int64 `json:"value"`
}

// This object is used to assert a desired state where this InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod = &InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{empty: true}

func (r *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReplicaConfigurationReplicaPoolConfiguration struct {
	empty                        bool                                                                              `json:"-"`
	Kind                         *string                                                                           `json:"kind"`
	StaticPoolConfiguration      *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration      `json:"staticPoolConfiguration"`
	AutoscalingPoolConfiguration *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration `json:"autoscalingPoolConfiguration"`
	ReplicaCount                 *int64                                                                            `json:"replicaCount"`
	ExposeReplicaIP              *bool                                                                             `json:"exposeReplicaIP"`
}

// This object is used to assert a desired state where this InstanceReplicaConfigurationReplicaPoolConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReplicaConfigurationReplicaPoolConfiguration *InstanceReplicaConfigurationReplicaPoolConfiguration = &InstanceReplicaConfigurationReplicaPoolConfiguration{empty: true}

func (r *InstanceReplicaConfigurationReplicaPoolConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReplicaConfigurationReplicaPoolConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration struct {
	empty           bool    `json:"-"`
	Kind            *string `json:"kind"`
	ReplicaCount    *int64  `json:"replicaCount"`
	ExposeReplicaIP *bool   `json:"exposeReplicaIP"`
}

// This object is used to assert a desired state where this InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration = &InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{empty: true}

func (r *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration struct {
	empty           bool     `json:"-"`
	Kind            *string  `json:"kind"`
	MinReplicaCount *int64   `json:"minReplicaCount"`
	MaxReplicaCount *int64   `json:"maxReplicaCount"`
	TargetCpuUtil   *float64 `json:"targetCpuUtil"`
}

// This object is used to assert a desired state where this InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration = &InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{empty: true}

func (r *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceScheduledMaintenance struct {
	empty         bool                                   `json:"-"`
	StartTime     *InstanceScheduledMaintenanceStartTime `json:"startTime"`
	CanDefer      *bool                                  `json:"canDefer"`
	CanReschedule *bool                                  `json:"canReschedule"`
}

// This object is used to assert a desired state where this InstanceScheduledMaintenance is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceScheduledMaintenance *InstanceScheduledMaintenance = &InstanceScheduledMaintenance{empty: true}

func (r *InstanceScheduledMaintenance) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceScheduledMaintenance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceScheduledMaintenanceStartTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this InstanceScheduledMaintenanceStartTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceScheduledMaintenanceStartTime *InstanceScheduledMaintenanceStartTime = &InstanceScheduledMaintenanceStartTime{empty: true}

func (r *InstanceScheduledMaintenanceStartTime) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceScheduledMaintenanceStartTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettings struct {
	empty                       bool                                     `json:"-"`
	AuthorizedGaeApplications   []string                                 `json:"authorizedGaeApplications"`
	Tier                        *string                                  `json:"tier"`
	Kind                        *string                                  `json:"kind"`
	AvailabilityType            *InstanceSettingsAvailabilityTypeEnum    `json:"availabilityType"`
	PricingPlan                 *InstanceSettingsPricingPlanEnum         `json:"pricingPlan"`
	ReplicationType             *InstanceSettingsReplicationTypeEnum     `json:"replicationType"`
	ActivationPolicy            *InstanceSettingsActivationPolicyEnum    `json:"activationPolicy"`
	StorageAutoResize           *bool                                    `json:"storageAutoResize"`
	DataDiskType                *InstanceSettingsDataDiskTypeEnum        `json:"dataDiskType"`
	DatabaseReplicationEnabled  *bool                                    `json:"databaseReplicationEnabled"`
	CrashSafeReplicationEnabled *bool                                    `json:"crashSafeReplicationEnabled"`
	SettingsVersion             *InstanceSettingsSettingsVersion         `json:"settingsVersion"`
	UserLabels                  map[string]string                        `json:"userLabels"`
	StorageAutoResizeLimit      *InstanceSettingsStorageAutoResizeLimit  `json:"storageAutoResizeLimit"`
	IPConfiguration             *InstanceSettingsIPConfiguration         `json:"ipConfiguration"`
	LocationPreference          *InstanceSettingsLocationPreference      `json:"locationPreference"`
	DatabaseFlags               []InstanceSettingsDatabaseFlags          `json:"databaseFlags"`
	MaintenanceWindow           *InstanceSettingsMaintenanceWindow       `json:"maintenanceWindow"`
	BackupConfiguration         *InstanceSettingsBackupConfiguration     `json:"backupConfiguration"`
	DataDiskSizeGb              *InstanceSettingsDataDiskSizeGb          `json:"dataDiskSizeGb"`
	ActiveDirectoryConfig       *InstanceSettingsActiveDirectoryConfig   `json:"activeDirectoryConfig"`
	Collation                   *string                                  `json:"collation"`
	DenyMaintenancePeriods      []InstanceSettingsDenyMaintenancePeriods `json:"denyMaintenancePeriods"`
	InsightsConfig              *InstanceSettingsInsightsConfig          `json:"insightsConfig"`
}

// This object is used to assert a desired state where this InstanceSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettings *InstanceSettings = &InstanceSettings{empty: true}

func (r *InstanceSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsSettingsVersion struct {
	empty bool   `json:"-"`
	Value *int64 `json:"value"`
}

// This object is used to assert a desired state where this InstanceSettingsSettingsVersion is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsSettingsVersion *InstanceSettingsSettingsVersion = &InstanceSettingsSettingsVersion{empty: true}

func (r *InstanceSettingsSettingsVersion) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsSettingsVersion) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsStorageAutoResizeLimit struct {
	empty bool   `json:"-"`
	Value *int64 `json:"value"`
}

// This object is used to assert a desired state where this InstanceSettingsStorageAutoResizeLimit is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsStorageAutoResizeLimit *InstanceSettingsStorageAutoResizeLimit = &InstanceSettingsStorageAutoResizeLimit{empty: true}

func (r *InstanceSettingsStorageAutoResizeLimit) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsStorageAutoResizeLimit) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsIPConfiguration struct {
	empty              bool                                                `json:"-"`
	IPv4Enabled        *bool                                               `json:"ipv4Enabled"`
	PrivateNetwork     *string                                             `json:"privateNetwork"`
	RequireSsl         *bool                                               `json:"requireSsl"`
	AuthorizedNetworks []InstanceSettingsIPConfigurationAuthorizedNetworks `json:"authorizedNetworks"`
}

// This object is used to assert a desired state where this InstanceSettingsIPConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsIPConfiguration *InstanceSettingsIPConfiguration = &InstanceSettingsIPConfiguration{empty: true}

func (r *InstanceSettingsIPConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsIPConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsIPConfigurationAuthorizedNetworks struct {
	empty          bool    `json:"-"`
	Value          *string `json:"value"`
	ExpirationTime *string `json:"expirationTime"`
	Name           *string `json:"name"`
	Kind           *string `json:"kind"`
}

// This object is used to assert a desired state where this InstanceSettingsIPConfigurationAuthorizedNetworks is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsIPConfigurationAuthorizedNetworks *InstanceSettingsIPConfigurationAuthorizedNetworks = &InstanceSettingsIPConfigurationAuthorizedNetworks{empty: true}

func (r *InstanceSettingsIPConfigurationAuthorizedNetworks) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsIPConfigurationAuthorizedNetworks) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsLocationPreference struct {
	empty bool    `json:"-"`
	Zone  *string `json:"zone"`
	Kind  *string `json:"kind"`
}

// This object is used to assert a desired state where this InstanceSettingsLocationPreference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsLocationPreference *InstanceSettingsLocationPreference = &InstanceSettingsLocationPreference{empty: true}

func (r *InstanceSettingsLocationPreference) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsLocationPreference) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsDatabaseFlags struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceSettingsDatabaseFlags is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsDatabaseFlags *InstanceSettingsDatabaseFlags = &InstanceSettingsDatabaseFlags{empty: true}

func (r *InstanceSettingsDatabaseFlags) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsDatabaseFlags) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsMaintenanceWindow struct {
	empty       bool                                              `json:"-"`
	Hour        *int64                                            `json:"hour"`
	Day         *int64                                            `json:"day"`
	UpdateTrack *InstanceSettingsMaintenanceWindowUpdateTrackEnum `json:"updateTrack"`
	Kind        *string                                           `json:"kind"`
}

// This object is used to assert a desired state where this InstanceSettingsMaintenanceWindow is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsMaintenanceWindow *InstanceSettingsMaintenanceWindow = &InstanceSettingsMaintenanceWindow{empty: true}

func (r *InstanceSettingsMaintenanceWindow) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsMaintenanceWindow) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsBackupConfiguration struct {
	empty                       bool                                                        `json:"-"`
	StartTime                   *string                                                     `json:"startTime"`
	Enabled                     *bool                                                       `json:"enabled"`
	Kind                        *string                                                     `json:"kind"`
	BinaryLogEnabled            *bool                                                       `json:"binaryLogEnabled"`
	Location                    *string                                                     `json:"location"`
	BackupRetentionSettings     *InstanceSettingsBackupConfigurationBackupRetentionSettings `json:"backupRetentionSettings"`
	TransactionLogRetentionDays *int64                                                      `json:"transactionLogRetentionDays"`
}

// This object is used to assert a desired state where this InstanceSettingsBackupConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsBackupConfiguration *InstanceSettingsBackupConfiguration = &InstanceSettingsBackupConfiguration{empty: true}

func (r *InstanceSettingsBackupConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsBackupConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsBackupConfigurationBackupRetentionSettings struct {
	empty           bool                                                                         `json:"-"`
	RetentionUnit   *InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum `json:"retentionUnit"`
	RetainedBackups *int64                                                                       `json:"retainedBackups"`
}

// This object is used to assert a desired state where this InstanceSettingsBackupConfigurationBackupRetentionSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsBackupConfigurationBackupRetentionSettings *InstanceSettingsBackupConfigurationBackupRetentionSettings = &InstanceSettingsBackupConfigurationBackupRetentionSettings{empty: true}

func (r *InstanceSettingsBackupConfigurationBackupRetentionSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsBackupConfigurationBackupRetentionSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsDataDiskSizeGb struct {
	empty bool   `json:"-"`
	Value *int64 `json:"value"`
}

// This object is used to assert a desired state where this InstanceSettingsDataDiskSizeGb is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsDataDiskSizeGb *InstanceSettingsDataDiskSizeGb = &InstanceSettingsDataDiskSizeGb{empty: true}

func (r *InstanceSettingsDataDiskSizeGb) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsDataDiskSizeGb) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsActiveDirectoryConfig struct {
	empty  bool    `json:"-"`
	Kind   *string `json:"kind"`
	Domain *string `json:"domain"`
}

// This object is used to assert a desired state where this InstanceSettingsActiveDirectoryConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsActiveDirectoryConfig *InstanceSettingsActiveDirectoryConfig = &InstanceSettingsActiveDirectoryConfig{empty: true}

func (r *InstanceSettingsActiveDirectoryConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsActiveDirectoryConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsDenyMaintenancePeriods struct {
	empty     bool    `json:"-"`
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
	Time      *string `json:"time"`
}

// This object is used to assert a desired state where this InstanceSettingsDenyMaintenancePeriods is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsDenyMaintenancePeriods *InstanceSettingsDenyMaintenancePeriods = &InstanceSettingsDenyMaintenancePeriods{empty: true}

func (r *InstanceSettingsDenyMaintenancePeriods) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsDenyMaintenancePeriods) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceSettingsInsightsConfig struct {
	empty                 bool   `json:"-"`
	QueryInsightsEnabled  *bool  `json:"queryInsightsEnabled"`
	RecordClientAddress   *bool  `json:"recordClientAddress"`
	RecordApplicationTags *bool  `json:"recordApplicationTags"`
	QueryStringLength     *int64 `json:"queryStringLength"`
}

// This object is used to assert a desired state where this InstanceSettingsInsightsConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSettingsInsightsConfig *InstanceSettingsInsightsConfig = &InstanceSettingsInsightsConfig{empty: true}

func (r *InstanceSettingsInsightsConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSettingsInsightsConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReplicaInstances struct {
	empty  bool    `json:"-"`
	Name   *string `json:"name"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceReplicaInstances is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReplicaInstances *InstanceReplicaInstances = &InstanceReplicaInstances{empty: true}

func (r *InstanceReplicaInstances) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReplicaInstances) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceServerCaCert struct {
	empty            bool    `json:"-"`
	Kind             *string `json:"kind"`
	CertSerialNumber *string `json:"certSerialNumber"`
	Cert             *string `json:"cert"`
	CreateTime       *string `json:"createTime"`
	CommonName       *string `json:"commonName"`
	ExpirationTime   *string `json:"expirationTime"`
	Sha1Fingerprint  *string `json:"sha1Fingerprint"`
	Instance         *string `json:"instance"`
}

// This object is used to assert a desired state where this InstanceServerCaCert is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceServerCaCert *InstanceServerCaCert = &InstanceServerCaCert{empty: true}

func (r *InstanceServerCaCert) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceServerCaCert) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceOnPremisesConfiguration struct {
	empty               bool     `json:"-"`
	HostPort            *string  `json:"hostPort"`
	Kind                *string  `json:"kind"`
	Username            *string  `json:"username"`
	Password            *string  `json:"password"`
	CaCertificate       *string  `json:"caCertificate"`
	ClientCertificate   *string  `json:"clientCertificate"`
	ClientKey           *string  `json:"clientKey"`
	DumpFilePath        *string  `json:"dumpFilePath"`
	Database            *string  `json:"database"`
	ReplicatedDatabases []string `json:"replicatedDatabases"`
}

// This object is used to assert a desired state where this InstanceOnPremisesConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceOnPremisesConfiguration *InstanceOnPremisesConfiguration = &InstanceOnPremisesConfiguration{empty: true}

func (r *InstanceOnPremisesConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceOnPremisesConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDiskEncryptionStatus struct {
	empty             bool    `json:"-"`
	KmsKeyVersionName *string `json:"kmsKeyVersionName"`
	Kind              *string `json:"kind"`
}

// This object is used to assert a desired state where this InstanceDiskEncryptionStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDiskEncryptionStatus *InstanceDiskEncryptionStatus = &InstanceDiskEncryptionStatus{empty: true}

func (r *InstanceDiskEncryptionStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDiskEncryptionStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Instance) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "sql",
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
}

func (l *InstanceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InstanceList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInstance(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInstance(ctx context.Context, project string) (*InstanceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListInstanceWithMaxResults(ctx, project, InstanceMaxPage)

}

func (c *Client) ListInstanceWithMaxResults(ctx context.Context, project string, pageSize int32) (*InstanceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listInstance(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InstanceList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetInstance(ctx context.Context, r *Instance) (*Instance, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

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
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Instance resource is nil")
	}
	c.Config.Logger.Info("Deleting Instance...")
	deleteOp := deleteInstanceOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInstance deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInstance(ctx context.Context, project string, filter func(*Instance) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListInstance(ctx, project)
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

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

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

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createInstanceOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapInstance(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeInstanceNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
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
