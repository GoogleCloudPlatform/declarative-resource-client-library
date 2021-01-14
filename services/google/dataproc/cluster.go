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
package dataproc

import (
	"context"
	"crypto/sha256"
	"fmt"
	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Cluster struct {
	Project       *string                `json:"project"`
	Name          *string                `json:"name"`
	Config        *ClusterClusterConfig  `json:"config"`
	Labels        []ClusterLabels        `json:"labels"`
	Status        *ClusterStatus         `json:"status"`
	StatusHistory []ClusterStatusHistory `json:"statusHistory"`
	ClusterUuid   *string                `json:"clusterUuid"`
	Metrics       *ClusterMetrics        `json:"metrics"`
	Location      *string                `json:"location"`
}

func (r *Cluster) String() string {
	return dcl.SprintResource(r)
}

// The enum ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.
type ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum string

// ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumRef returns a *ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumRef(s string) *ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if s == "" {
		return nil
	}

	v := ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(s)
	return &v
}

func (v ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) Validate() error {
	for _, s := range []string{"PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED", "INHERIT_FROM_SUBNETWORK", "OUTBOUND", "BIDIRECTIONAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.
type ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum string

// ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumRef returns a *ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumRef(s string) *ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if s == "" {
		return nil
	}

	v := ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(s)
	return &v
}

func (v ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "NO_RESERVATION", "ANY_RESERVATION", "SPECIFIC_RESERVATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterInstanceGroupConfigPreemptibilityEnum.
type ClusterInstanceGroupConfigPreemptibilityEnum string

// ClusterInstanceGroupConfigPreemptibilityEnumRef returns a *ClusterInstanceGroupConfigPreemptibilityEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterInstanceGroupConfigPreemptibilityEnumRef(s string) *ClusterInstanceGroupConfigPreemptibilityEnum {
	if s == "" {
		return nil
	}

	v := ClusterInstanceGroupConfigPreemptibilityEnum(s)
	return &v
}

func (v ClusterInstanceGroupConfigPreemptibilityEnum) Validate() error {
	for _, s := range []string{"PREEMPTIBILITY_UNSPECIFIED", "NON_PREEMPTIBLE", "PREEMPTIBLE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterInstanceGroupConfigPreemptibilityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterClusterConfigSoftwareConfigOptionalComponentsEnum.
type ClusterClusterConfigSoftwareConfigOptionalComponentsEnum string

// ClusterClusterConfigSoftwareConfigOptionalComponentsEnumRef returns a *ClusterClusterConfigSoftwareConfigOptionalComponentsEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterClusterConfigSoftwareConfigOptionalComponentsEnumRef(s string) *ClusterClusterConfigSoftwareConfigOptionalComponentsEnum {
	if s == "" {
		return nil
	}

	v := ClusterClusterConfigSoftwareConfigOptionalComponentsEnum(s)
	return &v
}

func (v ClusterClusterConfigSoftwareConfigOptionalComponentsEnum) Validate() error {
	for _, s := range []string{"COMPONENT_UNSPECIFIED", "ANACONDA", "DOCKER", "DRUID", "FLINK", "HBASE", "HIVE_WEBHCAT", "JUPYTER", "KERBEROS", "PRESTO", "RANGER", "SOLR", "ZEPPELIN", "ZOOKEEPER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterClusterConfigSoftwareConfigOptionalComponentsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterStatusStateEnum.
type ClusterStatusStateEnum string

// ClusterStatusStateEnumRef returns a *ClusterStatusStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterStatusStateEnumRef(s string) *ClusterStatusStateEnum {
	if s == "" {
		return nil
	}

	v := ClusterStatusStateEnum(s)
	return &v
}

func (v ClusterStatusStateEnum) Validate() error {
	for _, s := range []string{"UNKNOWN", "CREATING", "RUNNING", "ERROR", "DELETING", "UPDATING", "STOPPING", "STOPPED", "STARTING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterStatusStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterStatusSubstateEnum.
type ClusterStatusSubstateEnum string

// ClusterStatusSubstateEnumRef returns a *ClusterStatusSubstateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterStatusSubstateEnumRef(s string) *ClusterStatusSubstateEnum {
	if s == "" {
		return nil
	}

	v := ClusterStatusSubstateEnum(s)
	return &v
}

func (v ClusterStatusSubstateEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "UNHEALTHY", "STALE_STATUS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterStatusSubstateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterStatusHistoryStateEnum.
type ClusterStatusHistoryStateEnum string

// ClusterStatusHistoryStateEnumRef returns a *ClusterStatusHistoryStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterStatusHistoryStateEnumRef(s string) *ClusterStatusHistoryStateEnum {
	if s == "" {
		return nil
	}

	v := ClusterStatusHistoryStateEnum(s)
	return &v
}

func (v ClusterStatusHistoryStateEnum) Validate() error {
	for _, s := range []string{"UNKNOWN", "CREATING", "RUNNING", "ERROR", "DELETING", "UPDATING", "STOPPING", "STOPPED", "STARTING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterStatusHistoryStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterStatusHistorySubstateEnum.
type ClusterStatusHistorySubstateEnum string

// ClusterStatusHistorySubstateEnumRef returns a *ClusterStatusHistorySubstateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterStatusHistorySubstateEnumRef(s string) *ClusterStatusHistorySubstateEnum {
	if s == "" {
		return nil
	}

	v := ClusterStatusHistorySubstateEnum(s)
	return &v
}

func (v ClusterStatusHistorySubstateEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "UNHEALTHY", "STALE_STATUS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterStatusHistorySubstateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ClusterClusterConfig struct {
	empty                 bool                                        `json:"-"`
	StagingBucket         *string                                     `json:"stagingBucket"`
	TempBucket            *string                                     `json:"tempBucket"`
	GceClusterConfig      *ClusterClusterConfigGceClusterConfig       `json:"gceClusterConfig"`
	MasterConfig          *ClusterInstanceGroupConfig                 `json:"masterConfig"`
	WorkerConfig          *ClusterInstanceGroupConfig                 `json:"workerConfig"`
	SecondaryWorkerConfig *ClusterInstanceGroupConfig                 `json:"secondaryWorkerConfig"`
	SoftwareConfig        *ClusterClusterConfigSoftwareConfig         `json:"softwareConfig"`
	InitializationActions []ClusterClusterConfigInitializationActions `json:"initializationActions"`
	EncryptionConfig      *ClusterClusterConfigEncryptionConfig       `json:"encryptionConfig"`
	AutoscalingConfig     *ClusterClusterConfigAutoscalingConfig      `json:"autoscalingConfig"`
	SecurityConfig        *ClusterClusterConfigSecurityConfig         `json:"securityConfig"`
	LifecycleConfig       *ClusterClusterConfigLifecycleConfig        `json:"lifecycleConfig"`
	EndpointConfig        *ClusterClusterConfigEndpointConfig         `json:"endpointConfig"`
}

// This object is used to assert a desired state where this ClusterClusterConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfig *ClusterClusterConfig = &ClusterClusterConfig{empty: true}

func (r *ClusterClusterConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigGceClusterConfig struct {
	empty                   bool                                                             `json:"-"`
	Zone                    *string                                                          `json:"zone"`
	Network                 *string                                                          `json:"network"`
	Subnetwork              *string                                                          `json:"subnetwork"`
	InternalIPOnly          *bool                                                            `json:"internalIPOnly"`
	PrivateIPv6GoogleAccess *ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum `json:"privateIPv6GoogleAccess"`
	ServiceAccount          *string                                                          `json:"serviceAccount"`
	ServiceAccountScopes    []string                                                         `json:"serviceAccountScopes"`
	Tags                    []string                                                         `json:"tags"`
	Metadata                []ClusterClusterConfigGceClusterConfigMetadata                   `json:"metadata"`
	ReservationAffinity     *ClusterClusterConfigGceClusterConfigReservationAffinity         `json:"reservationAffinity"`
	NodeGroupAffinity       *ClusterClusterConfigGceClusterConfigNodeGroupAffinity           `json:"nodeGroupAffinity"`
}

// This object is used to assert a desired state where this ClusterClusterConfigGceClusterConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigGceClusterConfig *ClusterClusterConfigGceClusterConfig = &ClusterClusterConfigGceClusterConfig{empty: true}

func (r *ClusterClusterConfigGceClusterConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigGceClusterConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigGceClusterConfigMetadata struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this ClusterClusterConfigGceClusterConfigMetadata is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigGceClusterConfigMetadata *ClusterClusterConfigGceClusterConfigMetadata = &ClusterClusterConfigGceClusterConfigMetadata{empty: true}

func (r *ClusterClusterConfigGceClusterConfigMetadata) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigGceClusterConfigMetadata) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigGceClusterConfigReservationAffinity struct {
	empty                  bool                                                                               `json:"-"`
	ConsumeReservationType *ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum `json:"consumeReservationType"`
	Key                    *string                                                                            `json:"key"`
	Values                 []string                                                                           `json:"values"`
}

// This object is used to assert a desired state where this ClusterClusterConfigGceClusterConfigReservationAffinity is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigGceClusterConfigReservationAffinity *ClusterClusterConfigGceClusterConfigReservationAffinity = &ClusterClusterConfigGceClusterConfigReservationAffinity{empty: true}

func (r *ClusterClusterConfigGceClusterConfigReservationAffinity) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigGceClusterConfigReservationAffinity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigGceClusterConfigNodeGroupAffinity struct {
	empty     bool    `json:"-"`
	NodeGroup *string `json:"nodeGroup"`
}

// This object is used to assert a desired state where this ClusterClusterConfigGceClusterConfigNodeGroupAffinity is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigGceClusterConfigNodeGroupAffinity *ClusterClusterConfigGceClusterConfigNodeGroupAffinity = &ClusterClusterConfigGceClusterConfigNodeGroupAffinity{empty: true}

func (r *ClusterClusterConfigGceClusterConfigNodeGroupAffinity) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigGceClusterConfigNodeGroupAffinity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterInstanceGroupConfig struct {
	empty              bool                                           `json:"-"`
	NumInstances       *int64                                         `json:"numInstances"`
	InstanceNames      []string                                       `json:"instanceNames"`
	InstanceReferences []ClusterInstanceGroupConfigInstanceReferences `json:"instanceReferences"`
	Image              *string                                        `json:"image"`
	MachineType        *string                                        `json:"machineType"`
	DiskConfig         *ClusterInstanceGroupConfigDiskConfig          `json:"diskConfig"`
	IsPreemptible      *bool                                          `json:"isPreemptible"`
	Preemptibility     *ClusterInstanceGroupConfigPreemptibilityEnum  `json:"preemptibility"`
	ManagedGroupConfig *ClusterInstanceGroupConfigManagedGroupConfig  `json:"managedGroupConfig"`
	Accelerators       []ClusterInstanceGroupConfigAccelerators       `json:"accelerators"`
	MinCpuPlatform     *string                                        `json:"minCpuPlatform"`
}

// This object is used to assert a desired state where this ClusterInstanceGroupConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterInstanceGroupConfig *ClusterInstanceGroupConfig = &ClusterInstanceGroupConfig{empty: true}

func (r *ClusterInstanceGroupConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterInstanceGroupConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterInstanceGroupConfigInstanceReferences struct {
	empty        bool    `json:"-"`
	InstanceName *string `json:"instanceName"`
	InstanceId   *string `json:"instanceId"`
	PublicKey    *string `json:"publicKey"`
}

// This object is used to assert a desired state where this ClusterInstanceGroupConfigInstanceReferences is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterInstanceGroupConfigInstanceReferences *ClusterInstanceGroupConfigInstanceReferences = &ClusterInstanceGroupConfigInstanceReferences{empty: true}

func (r *ClusterInstanceGroupConfigInstanceReferences) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterInstanceGroupConfigInstanceReferences) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterInstanceGroupConfigDiskConfig struct {
	empty          bool    `json:"-"`
	BootDiskType   *string `json:"bootDiskType"`
	BootDiskSizeGb *int64  `json:"bootDiskSizeGb"`
	NumLocalSsds   *int64  `json:"numLocalSsds"`
}

// This object is used to assert a desired state where this ClusterInstanceGroupConfigDiskConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterInstanceGroupConfigDiskConfig *ClusterInstanceGroupConfigDiskConfig = &ClusterInstanceGroupConfigDiskConfig{empty: true}

func (r *ClusterInstanceGroupConfigDiskConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterInstanceGroupConfigDiskConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterInstanceGroupConfigManagedGroupConfig struct {
	empty                    bool    `json:"-"`
	InstanceTemplateName     *string `json:"instanceTemplateName"`
	InstanceGroupManagerName *string `json:"instanceGroupManagerName"`
}

// This object is used to assert a desired state where this ClusterInstanceGroupConfigManagedGroupConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterInstanceGroupConfigManagedGroupConfig *ClusterInstanceGroupConfigManagedGroupConfig = &ClusterInstanceGroupConfigManagedGroupConfig{empty: true}

func (r *ClusterInstanceGroupConfigManagedGroupConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterInstanceGroupConfigManagedGroupConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterInstanceGroupConfigAccelerators struct {
	empty            bool    `json:"-"`
	AcceleratorType  *string `json:"acceleratorType"`
	AcceleratorCount *int64  `json:"acceleratorCount"`
}

// This object is used to assert a desired state where this ClusterInstanceGroupConfigAccelerators is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterInstanceGroupConfigAccelerators *ClusterInstanceGroupConfigAccelerators = &ClusterInstanceGroupConfigAccelerators{empty: true}

func (r *ClusterInstanceGroupConfigAccelerators) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterInstanceGroupConfigAccelerators) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigSoftwareConfig struct {
	empty              bool                                                       `json:"-"`
	ImageVersion       *string                                                    `json:"imageVersion"`
	Properties         map[string]string                                          `json:"properties"`
	OptionalComponents []ClusterClusterConfigSoftwareConfigOptionalComponentsEnum `json:"optionalComponents"`
}

// This object is used to assert a desired state where this ClusterClusterConfigSoftwareConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigSoftwareConfig *ClusterClusterConfigSoftwareConfig = &ClusterClusterConfigSoftwareConfig{empty: true}

func (r *ClusterClusterConfigSoftwareConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigSoftwareConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigInitializationActions struct {
	empty            bool                                                       `json:"-"`
	ExecutableFile   *string                                                    `json:"executableFile"`
	ExecutionTimeout *ClusterClusterConfigInitializationActionsExecutionTimeout `json:"executionTimeout"`
}

// This object is used to assert a desired state where this ClusterClusterConfigInitializationActions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigInitializationActions *ClusterClusterConfigInitializationActions = &ClusterClusterConfigInitializationActions{empty: true}

func (r *ClusterClusterConfigInitializationActions) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigInitializationActions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigInitializationActionsExecutionTimeout struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this ClusterClusterConfigInitializationActionsExecutionTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigInitializationActionsExecutionTimeout *ClusterClusterConfigInitializationActionsExecutionTimeout = &ClusterClusterConfigInitializationActionsExecutionTimeout{empty: true}

func (r *ClusterClusterConfigInitializationActionsExecutionTimeout) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigInitializationActionsExecutionTimeout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigEncryptionConfig struct {
	empty           bool    `json:"-"`
	GcePdKmsKeyName *string `json:"gcePdKmsKeyName"`
}

// This object is used to assert a desired state where this ClusterClusterConfigEncryptionConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigEncryptionConfig *ClusterClusterConfigEncryptionConfig = &ClusterClusterConfigEncryptionConfig{empty: true}

func (r *ClusterClusterConfigEncryptionConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigEncryptionConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigAutoscalingConfig struct {
	empty  bool    `json:"-"`
	Policy *string `json:"policy"`
}

// This object is used to assert a desired state where this ClusterClusterConfigAutoscalingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigAutoscalingConfig *ClusterClusterConfigAutoscalingConfig = &ClusterClusterConfigAutoscalingConfig{empty: true}

func (r *ClusterClusterConfigAutoscalingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigAutoscalingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigSecurityConfig struct {
	empty          bool                                              `json:"-"`
	KerberosConfig *ClusterClusterConfigSecurityConfigKerberosConfig `json:"kerberosConfig"`
}

// This object is used to assert a desired state where this ClusterClusterConfigSecurityConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigSecurityConfig *ClusterClusterConfigSecurityConfig = &ClusterClusterConfigSecurityConfig{empty: true}

func (r *ClusterClusterConfigSecurityConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigSecurityConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigSecurityConfigKerberosConfig struct {
	empty                         bool    `json:"-"`
	EnableKerberos                *bool   `json:"enableKerberos"`
	RootPrincipalPassword         *string `json:"rootPrincipalPassword"`
	KmsKey                        *string `json:"kmsKey"`
	Keystore                      *string `json:"keystore"`
	Truststore                    *string `json:"truststore"`
	KeystorePassword              *string `json:"keystorePassword"`
	KeyPassword                   *string `json:"keyPassword"`
	TruststorePassword            *string `json:"truststorePassword"`
	CrossRealmTrustRealm          *string `json:"crossRealmTrustRealm"`
	CrossRealmTrustKdc            *string `json:"crossRealmTrustKdc"`
	CrossRealmTrustAdminServer    *string `json:"crossRealmTrustAdminServer"`
	CrossRealmTrustSharedPassword *string `json:"crossRealmTrustSharedPassword"`
	KdcDbKey                      *string `json:"kdcDbKey"`
	TgtLifetimeHours              *int64  `json:"tgtLifetimeHours"`
	Realm                         *string `json:"realm"`
}

// This object is used to assert a desired state where this ClusterClusterConfigSecurityConfigKerberosConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigSecurityConfigKerberosConfig *ClusterClusterConfigSecurityConfigKerberosConfig = &ClusterClusterConfigSecurityConfigKerberosConfig{empty: true}

func (r *ClusterClusterConfigSecurityConfigKerberosConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigSecurityConfigKerberosConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigLifecycleConfig struct {
	empty          bool                                              `json:"-"`
	IdleDeleteTtl  *ClusterClusterConfigLifecycleConfigIdleDeleteTtl `json:"idleDeleteTtl"`
	AutoDeleteTime *string                                           `json:"autoDeleteTime"`
	AutoDeleteTtl  *ClusterClusterConfigLifecycleConfigAutoDeleteTtl `json:"autoDeleteTtl"`
	IdleStartTime  *string                                           `json:"idleStartTime"`
}

// This object is used to assert a desired state where this ClusterClusterConfigLifecycleConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigLifecycleConfig *ClusterClusterConfigLifecycleConfig = &ClusterClusterConfigLifecycleConfig{empty: true}

func (r *ClusterClusterConfigLifecycleConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigLifecycleConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigLifecycleConfigIdleDeleteTtl struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this ClusterClusterConfigLifecycleConfigIdleDeleteTtl is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigLifecycleConfigIdleDeleteTtl *ClusterClusterConfigLifecycleConfigIdleDeleteTtl = &ClusterClusterConfigLifecycleConfigIdleDeleteTtl{empty: true}

func (r *ClusterClusterConfigLifecycleConfigIdleDeleteTtl) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigLifecycleConfigIdleDeleteTtl) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigLifecycleConfigAutoDeleteTtl struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this ClusterClusterConfigLifecycleConfigAutoDeleteTtl is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigLifecycleConfigAutoDeleteTtl *ClusterClusterConfigLifecycleConfigAutoDeleteTtl = &ClusterClusterConfigLifecycleConfigAutoDeleteTtl{empty: true}

func (r *ClusterClusterConfigLifecycleConfigAutoDeleteTtl) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigLifecycleConfigAutoDeleteTtl) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigEndpointConfig struct {
	empty                bool                                          `json:"-"`
	HttpPorts            []ClusterClusterConfigEndpointConfigHttpPorts `json:"httpPorts"`
	EnableHttpPortAccess *bool                                         `json:"enableHttpPortAccess"`
}

// This object is used to assert a desired state where this ClusterClusterConfigEndpointConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigEndpointConfig *ClusterClusterConfigEndpointConfig = &ClusterClusterConfigEndpointConfig{empty: true}

func (r *ClusterClusterConfigEndpointConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigEndpointConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterConfigEndpointConfigHttpPorts struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this ClusterClusterConfigEndpointConfigHttpPorts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterConfigEndpointConfigHttpPorts *ClusterClusterConfigEndpointConfigHttpPorts = &ClusterClusterConfigEndpointConfigHttpPorts{empty: true}

func (r *ClusterClusterConfigEndpointConfigHttpPorts) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterConfigEndpointConfigHttpPorts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterLabels struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this ClusterLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterLabels *ClusterLabels = &ClusterLabels{empty: true}

func (r *ClusterLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterStatus struct {
	empty          bool                       `json:"-"`
	State          *ClusterStatusStateEnum    `json:"state"`
	Detail         *string                    `json:"detail"`
	StateStartTime *string                    `json:"stateStartTime"`
	Substate       *ClusterStatusSubstateEnum `json:"substate"`
}

// This object is used to assert a desired state where this ClusterStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterStatus *ClusterStatus = &ClusterStatus{empty: true}

func (r *ClusterStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterStatusHistory struct {
	empty          bool                              `json:"-"`
	State          *ClusterStatusHistoryStateEnum    `json:"state"`
	Detail         *string                           `json:"detail"`
	StateStartTime *string                           `json:"stateStartTime"`
	Substate       *ClusterStatusHistorySubstateEnum `json:"substate"`
}

// This object is used to assert a desired state where this ClusterStatusHistory is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterStatusHistory *ClusterStatusHistory = &ClusterStatusHistory{empty: true}

func (r *ClusterStatusHistory) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterStatusHistory) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMetrics struct {
	empty       bool                        `json:"-"`
	HdfsMetrics []ClusterMetricsHdfsMetrics `json:"hdfsMetrics"`
	YarnMetrics []ClusterMetricsYarnMetrics `json:"yarnMetrics"`
}

// This object is used to assert a desired state where this ClusterMetrics is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMetrics *ClusterMetrics = &ClusterMetrics{empty: true}

func (r *ClusterMetrics) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMetrics) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMetricsHdfsMetrics struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *int64  `json:"value"`
}

// This object is used to assert a desired state where this ClusterMetricsHdfsMetrics is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMetricsHdfsMetrics *ClusterMetricsHdfsMetrics = &ClusterMetricsHdfsMetrics{empty: true}

func (r *ClusterMetricsHdfsMetrics) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMetricsHdfsMetrics) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMetricsYarnMetrics struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *int64  `json:"value"`
}

// This object is used to assert a desired state where this ClusterMetricsYarnMetrics is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMetricsYarnMetrics *ClusterMetricsYarnMetrics = &ClusterMetricsYarnMetrics{empty: true}

func (r *ClusterMetricsYarnMetrics) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMetricsYarnMetrics) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Cluster) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "dataproc",
		Type:    "Cluster",
		Version: "dataproc",
	}
}

const ClusterMaxPage = -1

type ClusterList struct {
	Items []*Cluster

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *ClusterList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ClusterList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listCluster(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListCluster(ctx context.Context, project, location string) (*ClusterList, error) {

	return c.ListClusterWithMaxResults(ctx, project, location, ClusterMaxPage)

}

func (c *Client) ListClusterWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ClusterList, error) {
	items, token, err := c.listCluster(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ClusterList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetCluster(ctx context.Context, r *Cluster) (*Cluster, error) {
	b, err := c.getClusterRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalCluster(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeClusterNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteCluster(ctx context.Context, r *Cluster) error {
	if r == nil {
		return fmt.Errorf("Cluster resource is nil")
	}
	c.Config.Logger.Info("Deleting Cluster...")
	deleteOp := deleteClusterOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllCluster deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllCluster(ctx context.Context, project, location string, filter func(*Cluster) bool) error {
	listObj, err := c.ListCluster(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllCluster(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllCluster(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyCluster(ctx context.Context, rawDesired *Cluster, opts ...dcl.ApplyOption) (*Cluster, error) {
	c.Config.Logger.Info("Beginning ApplyCluster...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.clusterDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []clusterApiOperation
	if create {
		ops = append(ops, &createClusterOperation{})
	} else if recreate {
		ops = append(ops, &deleteClusterOperation{})
		ops = append(ops, &createClusterOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeClusterDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetCluster(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeClusterNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeClusterDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffCluster(c, newDesired, newState)
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
