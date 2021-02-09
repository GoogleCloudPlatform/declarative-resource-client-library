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
package appengine

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Version struct {
	ConsumerName              *string                      `json:"consumerName"`
	Name                      *string                      `json:"name"`
	AutomaticScaling          *VersionAutomaticScaling     `json:"automaticScaling"`
	BasicScaling              *VersionBasicScaling         `json:"basicScaling"`
	ManualScaling             *VersionManualScaling        `json:"manualScaling"`
	JobScaling                *VersionJobScaling           `json:"jobScaling"`
	PoolScaling               *VersionPoolScaling          `json:"poolScaling"`
	InboundServices           []VersionInboundServicesEnum `json:"inboundServices"`
	InstanceClass             *string                      `json:"instanceClass"`
	Network                   *VersionNetwork              `json:"network"`
	Zones                     []string                     `json:"zones"`
	Resources                 *VersionResources            `json:"resources"`
	Runtime                   *string                      `json:"runtime"`
	RuntimeChannel            *string                      `json:"runtimeChannel"`
	Threadsafe                *bool                        `json:"threadsafe"`
	Vm                        *bool                        `json:"vm"`
	BetaSettings              map[string]string            `json:"betaSettings"`
	Env                       *string                      `json:"env"`
	ServingStatus             *VersionServingStatusEnum    `json:"servingStatus"`
	CreatedBy                 *string                      `json:"createdBy"`
	CreateTime                *string                      `json:"createTime"`
	DiskUsageBytes            *int64                       `json:"diskUsageBytes"`
	RuntimeApiVersion         *string                      `json:"runtimeApiVersion"`
	RuntimeMainExecutablePath *string                      `json:"runtimeMainExecutablePath"`
	Handlers                  []VersionHandlers            `json:"handlers"`
	ErrorHandlers             []VersionErrorHandlers       `json:"errorHandlers"`
	Libraries                 []VersionLibraries           `json:"libraries"`
	ApiConfig                 *VersionApiConfig            `json:"apiConfig"`
	EnvVariables              map[string]string            `json:"envVariables"`
	BuildEnvVariables         map[string]string            `json:"buildEnvVariables"`
	DefaultExpiration         *string                      `json:"defaultExpiration"`
	Deployment                *VersionDeployment           `json:"deployment"`
	HealthCheck               *VersionHealthCheck          `json:"healthCheck"`
	ReadinessCheck            *VersionReadinessCheck       `json:"readinessCheck"`
	LivenessCheck             *VersionLivenessCheck        `json:"livenessCheck"`
	NobuildFilesRegex         *string                      `json:"nobuildFilesRegex"`
	VersionUrl                *string                      `json:"versionUrl"`
	ServiceAuthSpec           *VersionServiceAuthSpec      `json:"serviceAuthSpec"`
	ServiceCorsSpec           *VersionServiceCorsSpec      `json:"serviceCorsSpec"`
	RouteHash                 *string                      `json:"routeHash"`
	Entrypoint                *VersionEntrypoint           `json:"entrypoint"`
	VPCAccessConnector        *VersionVPCAccessConnector   `json:"vpcAccessConnector"`
	NetworkSettings           *VersionNetworkSettings      `json:"networkSettings"`
	InstanceSpec              *VersionInstanceSpec         `json:"instanceSpec"`
	App                       *string                      `json:"app"`
	Service                   *string                      `json:"service"`
}

func (r *Version) String() string {
	return dcl.SprintResource(r)
}

// The enum VersionInboundServicesEnum.
type VersionInboundServicesEnum string

// VersionInboundServicesEnumRef returns a *VersionInboundServicesEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionInboundServicesEnumRef(s string) *VersionInboundServicesEnum {
	if s == "" {
		return nil
	}

	v := VersionInboundServicesEnum(s)
	return &v
}

func (v VersionInboundServicesEnum) Validate() error {
	for _, s := range []string{"INBOUND_SERVICE_UNSPECIFIED", "INBOUND_SERVICE_MAIL", "INBOUND_SERVICE_MAIL_BOUNCE", "INBOUND_SERVICE_XMPP_ERROR", "INBOUND_SERVICE_XMPP_MESSAGE", "INBOUND_SERVICE_XMPP_SUBSCRIBE", "INBOUND_SERVICE_XMPP_PRESENCE", "INBOUND_SERVICE_CHANNEL_PRESENCE", "INBOUND_SERVICE_WARMUP"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionInboundServicesEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionServingStatusEnum.
type VersionServingStatusEnum string

// VersionServingStatusEnumRef returns a *VersionServingStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionServingStatusEnumRef(s string) *VersionServingStatusEnum {
	if s == "" {
		return nil
	}

	v := VersionServingStatusEnum(s)
	return &v
}

func (v VersionServingStatusEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "SERVING", "USER_DISABLED", "SYSTEM_DISABLED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionServingStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionHandlersSecurityLevelEnum.
type VersionHandlersSecurityLevelEnum string

// VersionHandlersSecurityLevelEnumRef returns a *VersionHandlersSecurityLevelEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionHandlersSecurityLevelEnumRef(s string) *VersionHandlersSecurityLevelEnum {
	if s == "" {
		return nil
	}

	v := VersionHandlersSecurityLevelEnum(s)
	return &v
}

func (v VersionHandlersSecurityLevelEnum) Validate() error {
	for _, s := range []string{"SECURE_UNSPECIFIED", "SECURE_DEFAULT", "SECURE_NEVER", "SECURE_OPTIONAL", "SECURE_ALWAYS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionHandlersSecurityLevelEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionHandlersLoginEnum.
type VersionHandlersLoginEnum string

// VersionHandlersLoginEnumRef returns a *VersionHandlersLoginEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionHandlersLoginEnumRef(s string) *VersionHandlersLoginEnum {
	if s == "" {
		return nil
	}

	v := VersionHandlersLoginEnum(s)
	return &v
}

func (v VersionHandlersLoginEnum) Validate() error {
	for _, s := range []string{"LOGIN_UNSPECIFIED", "LOGIN_OPTIONAL", "LOGIN_ADMIN", "LOGIN_REQUIRED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionHandlersLoginEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionHandlersAuthFailActionEnum.
type VersionHandlersAuthFailActionEnum string

// VersionHandlersAuthFailActionEnumRef returns a *VersionHandlersAuthFailActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionHandlersAuthFailActionEnumRef(s string) *VersionHandlersAuthFailActionEnum {
	if s == "" {
		return nil
	}

	v := VersionHandlersAuthFailActionEnum(s)
	return &v
}

func (v VersionHandlersAuthFailActionEnum) Validate() error {
	for _, s := range []string{"AUTH_FAIL_ACTION_UNSPECIFIED", "AUTH_FAIL_ACTION_REDIRECT", "AUTH_FAIL_ACTION_UNAUTHORIZED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionHandlersAuthFailActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionHandlersRedirectHttpResponseCodeEnum.
type VersionHandlersRedirectHttpResponseCodeEnum string

// VersionHandlersRedirectHttpResponseCodeEnumRef returns a *VersionHandlersRedirectHttpResponseCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionHandlersRedirectHttpResponseCodeEnumRef(s string) *VersionHandlersRedirectHttpResponseCodeEnum {
	if s == "" {
		return nil
	}

	v := VersionHandlersRedirectHttpResponseCodeEnum(s)
	return &v
}

func (v VersionHandlersRedirectHttpResponseCodeEnum) Validate() error {
	for _, s := range []string{"REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED", "REDIRECT_HTTP_RESPONSE_CODE_301", "REDIRECT_HTTP_RESPONSE_CODE_302", "REDIRECT_HTTP_RESPONSE_CODE_303", "REDIRECT_HTTP_RESPONSE_CODE_307"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionHandlersRedirectHttpResponseCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionErrorHandlersErrorCodeEnum.
type VersionErrorHandlersErrorCodeEnum string

// VersionErrorHandlersErrorCodeEnumRef returns a *VersionErrorHandlersErrorCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionErrorHandlersErrorCodeEnumRef(s string) *VersionErrorHandlersErrorCodeEnum {
	if s == "" {
		return nil
	}

	v := VersionErrorHandlersErrorCodeEnum(s)
	return &v
}

func (v VersionErrorHandlersErrorCodeEnum) Validate() error {
	for _, s := range []string{"ERROR_CODE_UNSPECIFIED", "ERROR_CODE_DEFAULT", "ERROR_CODE_OVER_QUOTA", "ERROR_CODE_DOS_API_DENIAL", "ERROR_CODE_TIMEOUT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionErrorHandlersErrorCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionApiConfigAuthFailActionEnum.
type VersionApiConfigAuthFailActionEnum string

// VersionApiConfigAuthFailActionEnumRef returns a *VersionApiConfigAuthFailActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionApiConfigAuthFailActionEnumRef(s string) *VersionApiConfigAuthFailActionEnum {
	if s == "" {
		return nil
	}

	v := VersionApiConfigAuthFailActionEnum(s)
	return &v
}

func (v VersionApiConfigAuthFailActionEnum) Validate() error {
	for _, s := range []string{"AUTH_FAIL_ACTION_UNSPECIFIED", "AUTH_FAIL_ACTION_REDIRECT", "AUTH_FAIL_ACTION_UNAUTHORIZED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionApiConfigAuthFailActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionApiConfigLoginEnum.
type VersionApiConfigLoginEnum string

// VersionApiConfigLoginEnumRef returns a *VersionApiConfigLoginEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionApiConfigLoginEnumRef(s string) *VersionApiConfigLoginEnum {
	if s == "" {
		return nil
	}

	v := VersionApiConfigLoginEnum(s)
	return &v
}

func (v VersionApiConfigLoginEnum) Validate() error {
	for _, s := range []string{"LOGIN_UNSPECIFIED", "LOGIN_OPTIONAL", "LOGIN_ADMIN", "LOGIN_REQUIRED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionApiConfigLoginEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionApiConfigSecurityLevelEnum.
type VersionApiConfigSecurityLevelEnum string

// VersionApiConfigSecurityLevelEnumRef returns a *VersionApiConfigSecurityLevelEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionApiConfigSecurityLevelEnumRef(s string) *VersionApiConfigSecurityLevelEnum {
	if s == "" {
		return nil
	}

	v := VersionApiConfigSecurityLevelEnum(s)
	return &v
}

func (v VersionApiConfigSecurityLevelEnum) Validate() error {
	for _, s := range []string{"SECURE_UNSPECIFIED", "SECURE_DEFAULT", "SECURE_NEVER", "SECURE_OPTIONAL", "SECURE_ALWAYS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionApiConfigSecurityLevelEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionVPCAccessConnectorEgressSettingEnum.
type VersionVPCAccessConnectorEgressSettingEnum string

// VersionVPCAccessConnectorEgressSettingEnumRef returns a *VersionVPCAccessConnectorEgressSettingEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionVPCAccessConnectorEgressSettingEnumRef(s string) *VersionVPCAccessConnectorEgressSettingEnum {
	if s == "" {
		return nil
	}

	v := VersionVPCAccessConnectorEgressSettingEnum(s)
	return &v
}

func (v VersionVPCAccessConnectorEgressSettingEnum) Validate() error {
	for _, s := range []string{"EGRESS_SETTING_UNSPECIFIED", "ALL_TRAFFIC", "PRIVATE_IP_RANGES"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionVPCAccessConnectorEgressSettingEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionNetworkSettingsIngressTrafficAllowedEnum.
type VersionNetworkSettingsIngressTrafficAllowedEnum string

// VersionNetworkSettingsIngressTrafficAllowedEnumRef returns a *VersionNetworkSettingsIngressTrafficAllowedEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionNetworkSettingsIngressTrafficAllowedEnumRef(s string) *VersionNetworkSettingsIngressTrafficAllowedEnum {
	if s == "" {
		return nil
	}

	v := VersionNetworkSettingsIngressTrafficAllowedEnum(s)
	return &v
}

func (v VersionNetworkSettingsIngressTrafficAllowedEnum) Validate() error {
	for _, s := range []string{"INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED", "INGRESS_TRAFFIC_ALLOWED_ALL", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionNetworkSettingsIngressTrafficAllowedEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum VersionInstanceSpecPortsProtocolEnum.
type VersionInstanceSpecPortsProtocolEnum string

// VersionInstanceSpecPortsProtocolEnumRef returns a *VersionInstanceSpecPortsProtocolEnum with the value of string s
// If the empty string is provided, nil is returned.
func VersionInstanceSpecPortsProtocolEnumRef(s string) *VersionInstanceSpecPortsProtocolEnum {
	if s == "" {
		return nil
	}

	v := VersionInstanceSpecPortsProtocolEnum(s)
	return &v
}

func (v VersionInstanceSpecPortsProtocolEnum) Validate() error {
	for _, s := range []string{"PROTOCOL_UNSPECIFIED", "PROTOCOL_HTTP1", "PROTOCOL_H2C", "PROTOCOL_TCP"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VersionInstanceSpecPortsProtocolEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type VersionAutomaticScaling struct {
	empty                     bool                                              `json:"-"`
	CoolDownPeriod            *string                                           `json:"coolDownPeriod"`
	CpuUtilization            *VersionAutomaticScalingCpuUtilization            `json:"cpuUtilization"`
	MaxConcurrentRequests     *int64                                            `json:"maxConcurrentRequests"`
	MaxIdleInstances          *int64                                            `json:"maxIdleInstances"`
	MaxTotalInstances         *int64                                            `json:"maxTotalInstances"`
	MaxPendingLatency         *string                                           `json:"maxPendingLatency"`
	MinIdleInstances          *int64                                            `json:"minIdleInstances"`
	MinTotalInstances         *int64                                            `json:"minTotalInstances"`
	MinPendingLatency         *string                                           `json:"minPendingLatency"`
	RequestUtilization        *VersionAutomaticScalingRequestUtilization        `json:"requestUtilization"`
	DiskUtilization           *VersionAutomaticScalingDiskUtilization           `json:"diskUtilization"`
	NetworkUtilization        *VersionAutomaticScalingNetworkUtilization        `json:"networkUtilization"`
	StandardSchedulerSettings *VersionAutomaticScalingStandardSchedulerSettings `json:"standardSchedulerSettings"`
}

// This object is used to assert a desired state where this VersionAutomaticScaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionAutomaticScaling *VersionAutomaticScaling = &VersionAutomaticScaling{empty: true}

func (r *VersionAutomaticScaling) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionAutomaticScaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionAutomaticScalingCpuUtilization struct {
	empty                   bool     `json:"-"`
	AggregationWindowLength *string  `json:"aggregationWindowLength"`
	TargetUtilization       *float64 `json:"targetUtilization"`
}

// This object is used to assert a desired state where this VersionAutomaticScalingCpuUtilization is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionAutomaticScalingCpuUtilization *VersionAutomaticScalingCpuUtilization = &VersionAutomaticScalingCpuUtilization{empty: true}

func (r *VersionAutomaticScalingCpuUtilization) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionAutomaticScalingCpuUtilization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionAutomaticScalingRequestUtilization struct {
	empty                       bool   `json:"-"`
	TargetRequestCountPerSecond *int64 `json:"targetRequestCountPerSecond"`
	TargetConcurrentRequests    *int64 `json:"targetConcurrentRequests"`
}

// This object is used to assert a desired state where this VersionAutomaticScalingRequestUtilization is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionAutomaticScalingRequestUtilization *VersionAutomaticScalingRequestUtilization = &VersionAutomaticScalingRequestUtilization{empty: true}

func (r *VersionAutomaticScalingRequestUtilization) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionAutomaticScalingRequestUtilization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionAutomaticScalingDiskUtilization struct {
	empty                     bool   `json:"-"`
	TargetWriteBytesPerSecond *int64 `json:"targetWriteBytesPerSecond"`
	TargetWriteOpsPerSecond   *int64 `json:"targetWriteOpsPerSecond"`
	TargetReadBytesPerSecond  *int64 `json:"targetReadBytesPerSecond"`
	TargetReadOpsPerSecond    *int64 `json:"targetReadOpsPerSecond"`
}

// This object is used to assert a desired state where this VersionAutomaticScalingDiskUtilization is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionAutomaticScalingDiskUtilization *VersionAutomaticScalingDiskUtilization = &VersionAutomaticScalingDiskUtilization{empty: true}

func (r *VersionAutomaticScalingDiskUtilization) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionAutomaticScalingDiskUtilization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionAutomaticScalingNetworkUtilization struct {
	empty                          bool   `json:"-"`
	TargetSentBytesPerSecond       *int64 `json:"targetSentBytesPerSecond"`
	TargetSentPacketsPerSecond     *int64 `json:"targetSentPacketsPerSecond"`
	TargetReceivedBytesPerSecond   *int64 `json:"targetReceivedBytesPerSecond"`
	TargetReceivedPacketsPerSecond *int64 `json:"targetReceivedPacketsPerSecond"`
}

// This object is used to assert a desired state where this VersionAutomaticScalingNetworkUtilization is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionAutomaticScalingNetworkUtilization *VersionAutomaticScalingNetworkUtilization = &VersionAutomaticScalingNetworkUtilization{empty: true}

func (r *VersionAutomaticScalingNetworkUtilization) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionAutomaticScalingNetworkUtilization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionAutomaticScalingStandardSchedulerSettings struct {
	empty                       bool     `json:"-"`
	TargetCpuUtilization        *float64 `json:"targetCpuUtilization"`
	TargetThroughputUtilization *float64 `json:"targetThroughputUtilization"`
	MinInstances                *int64   `json:"minInstances"`
	MaxInstances                *int64   `json:"maxInstances"`
}

// This object is used to assert a desired state where this VersionAutomaticScalingStandardSchedulerSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionAutomaticScalingStandardSchedulerSettings *VersionAutomaticScalingStandardSchedulerSettings = &VersionAutomaticScalingStandardSchedulerSettings{empty: true}

func (r *VersionAutomaticScalingStandardSchedulerSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionAutomaticScalingStandardSchedulerSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionBasicScaling struct {
	empty        bool    `json:"-"`
	IdleTimeout  *string `json:"idleTimeout"`
	MaxInstances *int64  `json:"maxInstances"`
}

// This object is used to assert a desired state where this VersionBasicScaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionBasicScaling *VersionBasicScaling = &VersionBasicScaling{empty: true}

func (r *VersionBasicScaling) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionBasicScaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionManualScaling struct {
	empty     bool   `json:"-"`
	Instances *int64 `json:"instances"`
}

// This object is used to assert a desired state where this VersionManualScaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionManualScaling *VersionManualScaling = &VersionManualScaling{empty: true}

func (r *VersionManualScaling) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionManualScaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionJobScaling struct {
	empty                     bool    `json:"-"`
	Completions               *int64  `json:"completions"`
	Parallelism               *int64  `json:"parallelism"`
	JobDeadline               *string `json:"jobDeadline"`
	InstanceRetries           *int64  `json:"instanceRetries"`
	InstanceDeadline          *string `json:"instanceDeadline"`
	InstanceTerminationWindow *string `json:"instanceTerminationWindow"`
}

// This object is used to assert a desired state where this VersionJobScaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionJobScaling *VersionJobScaling = &VersionJobScaling{empty: true}

func (r *VersionJobScaling) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionJobScaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionPoolScaling struct {
	empty          bool   `json:"-"`
	Replicas       *int64 `json:"replicas"`
	MaxUnavailable *int64 `json:"maxUnavailable"`
	MaxSurge       *int64 `json:"maxSurge"`
}

// This object is used to assert a desired state where this VersionPoolScaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionPoolScaling *VersionPoolScaling = &VersionPoolScaling{empty: true}

func (r *VersionPoolScaling) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionPoolScaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionNetwork struct {
	empty           bool     `json:"-"`
	ForwardedPorts  []string `json:"forwardedPorts"`
	InstanceTag     *string  `json:"instanceTag"`
	Name            *string  `json:"name"`
	SubnetworkName  *string  `json:"subnetworkName"`
	SessionAffinity *bool    `json:"sessionAffinity"`
}

// This object is used to assert a desired state where this VersionNetwork is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionNetwork *VersionNetwork = &VersionNetwork{empty: true}

func (r *VersionNetwork) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionNetwork) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionResources struct {
	empty           bool                      `json:"-"`
	Cpu             *float64                  `json:"cpu"`
	DiskGb          *float64                  `json:"diskGb"`
	MemoryGb        *float64                  `json:"memoryGb"`
	Volumes         []VersionResourcesVolumes `json:"volumes"`
	KmsKeyReference *string                   `json:"kmsKeyReference"`
}

// This object is used to assert a desired state where this VersionResources is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionResources *VersionResources = &VersionResources{empty: true}

func (r *VersionResources) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionResourcesVolumes struct {
	empty      bool     `json:"-"`
	Name       *string  `json:"name"`
	VolumeType *string  `json:"volumeType"`
	SizeGb     *float64 `json:"sizeGb"`
}

// This object is used to assert a desired state where this VersionResourcesVolumes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionResourcesVolumes *VersionResourcesVolumes = &VersionResourcesVolumes{empty: true}

func (r *VersionResourcesVolumes) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionResourcesVolumes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionHandlers struct {
	empty                    bool                                         `json:"-"`
	UrlRegex                 *string                                      `json:"urlRegex"`
	StaticFiles              *VersionHandlersStaticFiles                  `json:"staticFiles"`
	Script                   *VersionHandlersScript                       `json:"script"`
	ApiEndpoint              *VersionHandlersApiEndpoint                  `json:"apiEndpoint"`
	SecurityLevel            *VersionHandlersSecurityLevelEnum            `json:"securityLevel"`
	Login                    *VersionHandlersLoginEnum                    `json:"login"`
	AuthFailAction           *VersionHandlersAuthFailActionEnum           `json:"authFailAction"`
	RedirectHttpResponseCode *VersionHandlersRedirectHttpResponseCodeEnum `json:"redirectHttpResponseCode"`
}

// This object is used to assert a desired state where this VersionHandlers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionHandlers *VersionHandlers = &VersionHandlers{empty: true}

func (r *VersionHandlers) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionHandlers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionHandlersStaticFiles struct {
	empty               bool              `json:"-"`
	Path                *string           `json:"path"`
	UploadPathRegex     *string           `json:"uploadPathRegex"`
	HttpHeaders         map[string]string `json:"httpHeaders"`
	MimeType            *string           `json:"mimeType"`
	Expiration          *string           `json:"expiration"`
	RequireMatchingFile *bool             `json:"requireMatchingFile"`
	ApplicationReadable *bool             `json:"applicationReadable"`
}

// This object is used to assert a desired state where this VersionHandlersStaticFiles is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionHandlersStaticFiles *VersionHandlersStaticFiles = &VersionHandlersStaticFiles{empty: true}

func (r *VersionHandlersStaticFiles) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionHandlersStaticFiles) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionHandlersScript struct {
	empty      bool    `json:"-"`
	ScriptPath *string `json:"scriptPath"`
}

// This object is used to assert a desired state where this VersionHandlersScript is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionHandlersScript *VersionHandlersScript = &VersionHandlersScript{empty: true}

func (r *VersionHandlersScript) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionHandlersScript) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionHandlersApiEndpoint struct {
	empty      bool    `json:"-"`
	ScriptPath *string `json:"scriptPath"`
}

// This object is used to assert a desired state where this VersionHandlersApiEndpoint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionHandlersApiEndpoint *VersionHandlersApiEndpoint = &VersionHandlersApiEndpoint{empty: true}

func (r *VersionHandlersApiEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionHandlersApiEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionErrorHandlers struct {
	empty      bool                               `json:"-"`
	ErrorCode  *VersionErrorHandlersErrorCodeEnum `json:"errorCode"`
	StaticFile *string                            `json:"staticFile"`
	MimeType   *string                            `json:"mimeType"`
}

// This object is used to assert a desired state where this VersionErrorHandlers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionErrorHandlers *VersionErrorHandlers = &VersionErrorHandlers{empty: true}

func (r *VersionErrorHandlers) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionErrorHandlers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionLibraries struct {
	empty   bool    `json:"-"`
	Name    *string `json:"name"`
	Version *string `json:"version"`
}

// This object is used to assert a desired state where this VersionLibraries is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionLibraries *VersionLibraries = &VersionLibraries{empty: true}

func (r *VersionLibraries) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionLibraries) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionApiConfig struct {
	empty          bool                                `json:"-"`
	AuthFailAction *VersionApiConfigAuthFailActionEnum `json:"authFailAction"`
	Login          *VersionApiConfigLoginEnum          `json:"login"`
	Script         *string                             `json:"script"`
	SecurityLevel  *VersionApiConfigSecurityLevelEnum  `json:"securityLevel"`
	Url            *string                             `json:"url"`
}

// This object is used to assert a desired state where this VersionApiConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionApiConfig *VersionApiConfig = &VersionApiConfig{empty: true}

func (r *VersionApiConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionApiConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionDeployment struct {
	empty             bool                                `json:"-"`
	Files             map[string]VersionDeploymentFiles   `json:"files"`
	Container         *VersionDeploymentContainer         `json:"container"`
	Zip               *VersionDeploymentZip               `json:"zip"`
	CloudBuildOptions *VersionDeploymentCloudBuildOptions `json:"cloudBuildOptions"`
}

// This object is used to assert a desired state where this VersionDeployment is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionDeployment *VersionDeployment = &VersionDeployment{empty: true}

func (r *VersionDeployment) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionDeployment) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionDeploymentFiles struct {
	empty     bool    `json:"-"`
	SourceUrl *string `json:"sourceUrl"`
	Sha1Sum   *string `json:"sha1Sum"`
	MimeType  *string `json:"mimeType"`
}

// This object is used to assert a desired state where this VersionDeploymentFiles is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionDeploymentFiles *VersionDeploymentFiles = &VersionDeploymentFiles{empty: true}

func (r *VersionDeploymentFiles) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionDeploymentFiles) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionDeploymentContainer struct {
	empty bool    `json:"-"`
	Image *string `json:"image"`
}

// This object is used to assert a desired state where this VersionDeploymentContainer is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionDeploymentContainer *VersionDeploymentContainer = &VersionDeploymentContainer{empty: true}

func (r *VersionDeploymentContainer) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionDeploymentContainer) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionDeploymentZip struct {
	empty      bool    `json:"-"`
	SourceUrl  *string `json:"sourceUrl"`
	FilesCount *int64  `json:"filesCount"`
}

// This object is used to assert a desired state where this VersionDeploymentZip is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionDeploymentZip *VersionDeploymentZip = &VersionDeploymentZip{empty: true}

func (r *VersionDeploymentZip) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionDeploymentZip) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionDeploymentCloudBuildOptions struct {
	empty             bool    `json:"-"`
	AppYamlPath       *string `json:"appYamlPath"`
	CloudBuildTimeout *string `json:"cloudBuildTimeout"`
}

// This object is used to assert a desired state where this VersionDeploymentCloudBuildOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionDeploymentCloudBuildOptions *VersionDeploymentCloudBuildOptions = &VersionDeploymentCloudBuildOptions{empty: true}

func (r *VersionDeploymentCloudBuildOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionDeploymentCloudBuildOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionHealthCheck struct {
	empty              bool    `json:"-"`
	DisableHealthCheck *bool   `json:"disableHealthCheck"`
	Host               *string `json:"host"`
	HealthyThreshold   *int64  `json:"healthyThreshold"`
	UnhealthyThreshold *int64  `json:"unhealthyThreshold"`
	RestartThreshold   *int64  `json:"restartThreshold"`
	CheckInterval      *string `json:"checkInterval"`
	Timeout            *string `json:"timeout"`
}

// This object is used to assert a desired state where this VersionHealthCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionHealthCheck *VersionHealthCheck = &VersionHealthCheck{empty: true}

func (r *VersionHealthCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionHealthCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionReadinessCheck struct {
	empty            bool    `json:"-"`
	Path             *string `json:"path"`
	Host             *string `json:"host"`
	FailureThreshold *int64  `json:"failureThreshold"`
	SuccessThreshold *int64  `json:"successThreshold"`
	CheckInterval    *string `json:"checkInterval"`
	Timeout          *string `json:"timeout"`
	AppStartTimeout  *string `json:"appStartTimeout"`
}

// This object is used to assert a desired state where this VersionReadinessCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionReadinessCheck *VersionReadinessCheck = &VersionReadinessCheck{empty: true}

func (r *VersionReadinessCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionReadinessCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionLivenessCheck struct {
	empty            bool    `json:"-"`
	Path             *string `json:"path"`
	Host             *string `json:"host"`
	FailureThreshold *int64  `json:"failureThreshold"`
	SuccessThreshold *int64  `json:"successThreshold"`
	CheckInterval    *string `json:"checkInterval"`
	Timeout          *string `json:"timeout"`
	InitialDelay     *string `json:"initialDelay"`
}

// This object is used to assert a desired state where this VersionLivenessCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionLivenessCheck *VersionLivenessCheck = &VersionLivenessCheck{empty: true}

func (r *VersionLivenessCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionLivenessCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionServiceAuthSpec struct {
	empty                bool     `json:"-"`
	Audiences            []string `json:"audiences"`
	IamServiceName       *string  `json:"iamServiceName"`
	IamResourceName      *string  `json:"iamResourceName"`
	IamPolicyId          *string  `json:"iamPolicyId"`
	IamPolicyType        *string  `json:"iamPolicyType"`
	IamPermission        *string  `json:"iamPermission"`
	AcceptGcloudClientId *bool    `json:"acceptGcloudClientId"`
	Clear                *bool    `json:"clear"`
}

// This object is used to assert a desired state where this VersionServiceAuthSpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionServiceAuthSpec *VersionServiceAuthSpec = &VersionServiceAuthSpec{empty: true}

func (r *VersionServiceAuthSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionServiceAuthSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionServiceCorsSpec struct {
	empty           bool     `json:"-"`
	Origin          []string `json:"origin"`
	Method          []string `json:"method"`
	Header          []string `json:"header"`
	ExposedHeader   []string `json:"exposedHeader"`
	AllowCredential *bool    `json:"allowCredential"`
	MaxAgeSeconds   *int64   `json:"maxAgeSeconds"`
}

// This object is used to assert a desired state where this VersionServiceCorsSpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionServiceCorsSpec *VersionServiceCorsSpec = &VersionServiceCorsSpec{empty: true}

func (r *VersionServiceCorsSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionServiceCorsSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionEntrypoint struct {
	empty bool    `json:"-"`
	Shell *string `json:"shell"`
}

// This object is used to assert a desired state where this VersionEntrypoint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionEntrypoint *VersionEntrypoint = &VersionEntrypoint{empty: true}

func (r *VersionEntrypoint) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionEntrypoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionVPCAccessConnector struct {
	empty         bool                                        `json:"-"`
	Name          *string                                     `json:"name"`
	EgressSetting *VersionVPCAccessConnectorEgressSettingEnum `json:"egressSetting"`
}

// This object is used to assert a desired state where this VersionVPCAccessConnector is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionVPCAccessConnector *VersionVPCAccessConnector = &VersionVPCAccessConnector{empty: true}

func (r *VersionVPCAccessConnector) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionVPCAccessConnector) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionNetworkSettings struct {
	empty                 bool                                             `json:"-"`
	IngressTrafficAllowed *VersionNetworkSettingsIngressTrafficAllowedEnum `json:"ingressTrafficAllowed"`
}

// This object is used to assert a desired state where this VersionNetworkSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionNetworkSettings *VersionNetworkSettings = &VersionNetworkSettings{empty: true}

func (r *VersionNetworkSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionNetworkSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionInstanceSpec struct {
	empty     bool                           `json:"-"`
	Sandboxes []VersionInstanceSpecSandboxes `json:"sandboxes"`
	Ports     []VersionInstanceSpecPorts     `json:"ports"`
}

// This object is used to assert a desired state where this VersionInstanceSpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionInstanceSpec *VersionInstanceSpec = &VersionInstanceSpec{empty: true}

func (r *VersionInstanceSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionInstanceSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionInstanceSpecSandboxes struct {
	empty      bool                                     `json:"-"`
	Name       *string                                  `json:"name"`
	Containers []VersionInstanceSpecSandboxesContainers `json:"containers"`
}

// This object is used to assert a desired state where this VersionInstanceSpecSandboxes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionInstanceSpecSandboxes *VersionInstanceSpecSandboxes = &VersionInstanceSpecSandboxes{empty: true}

func (r *VersionInstanceSpecSandboxes) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionInstanceSpecSandboxes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionInstanceSpecSandboxesContainers struct {
	empty bool    `json:"-"`
	Ports []int64 `json:"ports"`
}

// This object is used to assert a desired state where this VersionInstanceSpecSandboxesContainers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionInstanceSpecSandboxesContainers *VersionInstanceSpecSandboxesContainers = &VersionInstanceSpecSandboxesContainers{empty: true}

func (r *VersionInstanceSpecSandboxesContainers) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionInstanceSpecSandboxesContainers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type VersionInstanceSpecPorts struct {
	empty     bool                                  `json:"-"`
	Name      *string                               `json:"name"`
	Sandbox   *string                               `json:"sandbox"`
	Port      *int64                                `json:"port"`
	Protocol  *VersionInstanceSpecPortsProtocolEnum `json:"protocol"`
	IsDefault *bool                                 `json:"isDefault"`
}

// This object is used to assert a desired state where this VersionInstanceSpecPorts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyVersionInstanceSpecPorts *VersionInstanceSpecPorts = &VersionInstanceSpecPorts{empty: true}

func (r *VersionInstanceSpecPorts) String() string {
	return dcl.SprintResource(r)
}

func (r *VersionInstanceSpecPorts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Version) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "appengine",
		Type:    "Version",
		Version: "appengine",
	}
}

const VersionMaxPage = -1

type VersionList struct {
	Items []*Version

	nextToken string

	pageSize int32

	app string

	service string
}

func (l *VersionList) HasNext() bool {
	return l.nextToken != ""
}

func (l *VersionList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listVersion(ctx, l.app, l.service, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListVersion(ctx context.Context, app, service string) (*VersionList, error) {

	return c.ListVersionWithMaxResults(ctx, app, service, VersionMaxPage)

}

func (c *Client) ListVersionWithMaxResults(ctx context.Context, app, service string, pageSize int32) (*VersionList, error) {
	items, token, err := c.listVersion(ctx, app, service, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &VersionList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		app: app,

		service: service,
	}, nil
}

func (c *Client) GetVersion(ctx context.Context, r *Version) (*Version, error) {
	b, err := c.getVersionRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalVersion(b, c)
	if err != nil {
		return nil, err
	}
	result.App = r.App
	result.Service = r.Service
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeVersionNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteVersion(ctx context.Context, r *Version) error {
	if r == nil {
		return fmt.Errorf("Version resource is nil")
	}
	c.Config.Logger.Info("Deleting Version...")
	deleteOp := deleteVersionOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllVersion deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllVersion(ctx context.Context, app, service string, filter func(*Version) bool) error {
	listObj, err := c.ListVersion(ctx, app, service)
	if err != nil {
		return err
	}

	err = c.deleteAllVersion(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllVersion(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyVersion(ctx context.Context, rawDesired *Version, opts ...dcl.ApplyOption) (*Version, error) {
	c.Config.Logger.Info("Beginning ApplyVersion...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.versionDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []versionApiOperation
	if create {
		ops = append(ops, &createVersionOperation{})
	} else if recreate {

		ops = append(ops, &deleteVersionOperation{})

		ops = append(ops, &createVersionOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeVersionDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetVersion(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeVersionNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeVersionDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffVersion(c, newDesired, newState)
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
