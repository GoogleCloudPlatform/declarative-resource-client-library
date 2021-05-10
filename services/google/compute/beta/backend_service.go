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
	"encoding/json"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type BackendService struct {
	Name                     *string                                 `json:"name"`
	Description              *string                                 `json:"description"`
	SelfLink                 *string                                 `json:"selfLink"`
	SelfLinkWithId           *string                                 `json:"selfLinkWithId"`
	Backends                 []BackendServiceBackends                `json:"backends"`
	HealthChecks             []string                                `json:"healthChecks"`
	TimeoutSec               *int64                                  `json:"timeoutSec"`
	Port                     *int64                                  `json:"port"`
	Protocol                 *BackendServiceProtocolEnum             `json:"protocol"`
	Fingerprint              *string                                 `json:"fingerprint"`
	PortName                 *string                                 `json:"portName"`
	EnableCdn                *bool                                   `json:"enableCdn"`
	SessionAffinity          *BackendServiceSessionAffinityEnum      `json:"sessionAffinity"`
	AffinityCookieTtlSec     *int64                                  `json:"affinityCookieTtlSec"`
	Location                 *string                                 `json:"location"`
	FailoverPolicy           *BackendServiceFailoverPolicy           `json:"failoverPolicy"`
	LoadBalancingScheme      *BackendServiceLoadBalancingSchemeEnum  `json:"loadBalancingScheme"`
	ConnectionDraining       *BackendServiceConnectionDraining       `json:"connectionDraining"`
	Iap                      *BackendServiceIap                      `json:"iap"`
	CdnPolicy                *BackendServiceCdnPolicy                `json:"cdnPolicy"`
	CustomRequestHeaders     []string                                `json:"customRequestHeaders"`
	CustomResponseHeaders    []string                                `json:"customResponseHeaders"`
	SecurityPolicy           *string                                 `json:"securityPolicy"`
	LogConfig                *BackendServiceLogConfig                `json:"logConfig"`
	SecuritySettings         *BackendServiceSecuritySettings         `json:"securitySettings"`
	LocalityLbPolicy         *BackendServiceLocalityLbPolicyEnum     `json:"localityLbPolicy"`
	ConsistentHash           *BackendServiceConsistentHash           `json:"consistentHash"`
	CircuitBreakers          *BackendServiceCircuitBreakers          `json:"circuitBreakers"`
	OutlierDetection         *BackendServiceOutlierDetection         `json:"outlierDetection"`
	Network                  *string                                 `json:"network"`
	Subsetting               *BackendServiceSubsetting               `json:"subsetting"`
	ConnectionTrackingPolicy *BackendServiceConnectionTrackingPolicy `json:"connectionTrackingPolicy"`
	MaxStreamDuration        *BackendServiceMaxStreamDuration        `json:"maxStreamDuration"`
	Project                  *string                                 `json:"project"`
}

func (r *BackendService) String() string {
	return dcl.SprintResource(r)
}

// The enum BackendServiceBackendsBalancingModeEnum.
type BackendServiceBackendsBalancingModeEnum string

// BackendServiceBackendsBalancingModeEnumRef returns a *BackendServiceBackendsBalancingModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackendServiceBackendsBalancingModeEnumRef(s string) *BackendServiceBackendsBalancingModeEnum {
	if s == "" {
		return nil
	}

	v := BackendServiceBackendsBalancingModeEnum(s)
	return &v
}

func (v BackendServiceBackendsBalancingModeEnum) Validate() error {
	for _, s := range []string{"UTILIZATION", "RATE", "CONNECTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackendServiceBackendsBalancingModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BackendServiceProtocolEnum.
type BackendServiceProtocolEnum string

// BackendServiceProtocolEnumRef returns a *BackendServiceProtocolEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackendServiceProtocolEnumRef(s string) *BackendServiceProtocolEnum {
	if s == "" {
		return nil
	}

	v := BackendServiceProtocolEnum(s)
	return &v
}

func (v BackendServiceProtocolEnum) Validate() error {
	for _, s := range []string{"HTTP", "HTTPS", "HTTP2", "TCP", "SSL", "UDP", "GRPC", "ALL", "UNSPECIFIED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackendServiceProtocolEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BackendServiceSessionAffinityEnum.
type BackendServiceSessionAffinityEnum string

// BackendServiceSessionAffinityEnumRef returns a *BackendServiceSessionAffinityEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackendServiceSessionAffinityEnumRef(s string) *BackendServiceSessionAffinityEnum {
	if s == "" {
		return nil
	}

	v := BackendServiceSessionAffinityEnum(s)
	return &v
}

func (v BackendServiceSessionAffinityEnum) Validate() error {
	for _, s := range []string{"NONE", "CLIENT_IP", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "CLIENT_IP_PORT_PROTO", "HTTP_COOKIE", "HEADER_FIELD", "CLIENT_IP_NO_DESTINATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackendServiceSessionAffinityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BackendServiceLoadBalancingSchemeEnum.
type BackendServiceLoadBalancingSchemeEnum string

// BackendServiceLoadBalancingSchemeEnumRef returns a *BackendServiceLoadBalancingSchemeEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackendServiceLoadBalancingSchemeEnumRef(s string) *BackendServiceLoadBalancingSchemeEnum {
	if s == "" {
		return nil
	}

	v := BackendServiceLoadBalancingSchemeEnum(s)
	return &v
}

func (v BackendServiceLoadBalancingSchemeEnum) Validate() error {
	for _, s := range []string{"INVALID_LOAD_BALANCING_SCHEME", "INTERNAL", "INTERNAL_MANAGED", "INTERNAL_SELF_MANAGED", "EXTERNAL", "EXTERNAL_MANAGED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackendServiceLoadBalancingSchemeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BackendServiceCdnPolicyCacheModeEnum.
type BackendServiceCdnPolicyCacheModeEnum string

// BackendServiceCdnPolicyCacheModeEnumRef returns a *BackendServiceCdnPolicyCacheModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackendServiceCdnPolicyCacheModeEnumRef(s string) *BackendServiceCdnPolicyCacheModeEnum {
	if s == "" {
		return nil
	}

	v := BackendServiceCdnPolicyCacheModeEnum(s)
	return &v
}

func (v BackendServiceCdnPolicyCacheModeEnum) Validate() error {
	for _, s := range []string{"INVALID_CACHE_MODE", "USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackendServiceCdnPolicyCacheModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BackendServiceLocalityLbPolicyEnum.
type BackendServiceLocalityLbPolicyEnum string

// BackendServiceLocalityLbPolicyEnumRef returns a *BackendServiceLocalityLbPolicyEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackendServiceLocalityLbPolicyEnumRef(s string) *BackendServiceLocalityLbPolicyEnum {
	if s == "" {
		return nil
	}

	v := BackendServiceLocalityLbPolicyEnum(s)
	return &v
}

func (v BackendServiceLocalityLbPolicyEnum) Validate() error {
	for _, s := range []string{"INVALID_LB_POLICY", "ROUND_ROBIN", "LEAST_REQUEST", "RING_HASH", "RANDOM", "ORIGINAL_DESTINATION", "MAGLEV"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackendServiceLocalityLbPolicyEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BackendServiceSubsettingPolicyEnum.
type BackendServiceSubsettingPolicyEnum string

// BackendServiceSubsettingPolicyEnumRef returns a *BackendServiceSubsettingPolicyEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackendServiceSubsettingPolicyEnumRef(s string) *BackendServiceSubsettingPolicyEnum {
	if s == "" {
		return nil
	}

	v := BackendServiceSubsettingPolicyEnum(s)
	return &v
}

func (v BackendServiceSubsettingPolicyEnum) Validate() error {
	for _, s := range []string{"NONE", "CONSISTENT_HASH_SUBSETTING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackendServiceSubsettingPolicyEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BackendServiceConnectionTrackingPolicyTrackingModeEnum.
type BackendServiceConnectionTrackingPolicyTrackingModeEnum string

// BackendServiceConnectionTrackingPolicyTrackingModeEnumRef returns a *BackendServiceConnectionTrackingPolicyTrackingModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackendServiceConnectionTrackingPolicyTrackingModeEnumRef(s string) *BackendServiceConnectionTrackingPolicyTrackingModeEnum {
	if s == "" {
		return nil
	}

	v := BackendServiceConnectionTrackingPolicyTrackingModeEnum(s)
	return &v
}

func (v BackendServiceConnectionTrackingPolicyTrackingModeEnum) Validate() error {
	for _, s := range []string{"INVALID_TRACKING_MODE", "PER_CONNECTION", "PER_SESSION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackendServiceConnectionTrackingPolicyTrackingModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum.
type BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum string

// BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnumRef returns a *BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnumRef(s string) *BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum {
	if s == "" {
		return nil
	}

	v := BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum(s)
	return &v
}

func (v BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum) Validate() error {
	for _, s := range []string{"DEFAULT_FOR_PROTOCOL", "NEVER_PERSIST", "ALWAYS_PERSIST"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type BackendServiceBackends struct {
	empty                     bool                                     `json:"-"`
	Description               *string                                  `json:"description"`
	Group                     *string                                  `json:"group"`
	BalancingMode             *BackendServiceBackendsBalancingModeEnum `json:"balancingMode"`
	MaxUtilization            *float64                                 `json:"maxUtilization"`
	MaxRate                   *int64                                   `json:"maxRate"`
	MaxRatePerInstance        *float64                                 `json:"maxRatePerInstance"`
	MaxRatePerEndpoint        *float64                                 `json:"maxRatePerEndpoint"`
	MaxConnections            *int64                                   `json:"maxConnections"`
	MaxConnectionsPerInstance *int64                                   `json:"maxConnectionsPerInstance"`
	MaxConnectionsPerEndpoint *int64                                   `json:"maxConnectionsPerEndpoint"`
	CapacityScaler            *float64                                 `json:"capacityScaler"`
	Failover                  *bool                                    `json:"failover"`
}

type jsonBackendServiceBackends BackendServiceBackends

func (r *BackendServiceBackends) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceBackends
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceBackends
	} else {

		r.Description = res.Description

		r.Group = res.Group

		r.BalancingMode = res.BalancingMode

		r.MaxUtilization = res.MaxUtilization

		r.MaxRate = res.MaxRate

		r.MaxRatePerInstance = res.MaxRatePerInstance

		r.MaxRatePerEndpoint = res.MaxRatePerEndpoint

		r.MaxConnections = res.MaxConnections

		r.MaxConnectionsPerInstance = res.MaxConnectionsPerInstance

		r.MaxConnectionsPerEndpoint = res.MaxConnectionsPerEndpoint

		r.CapacityScaler = res.CapacityScaler

		r.Failover = res.Failover

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceBackends is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceBackends *BackendServiceBackends = &BackendServiceBackends{empty: true}

func (r *BackendServiceBackends) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceBackends) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceFailoverPolicy struct {
	empty                            bool     `json:"-"`
	DisableConnectionDrainOnFailover *bool    `json:"disableConnectionDrainOnFailover"`
	DropTrafficIfUnhealthy           *bool    `json:"dropTrafficIfUnhealthy"`
	FailoverRatio                    *float64 `json:"failoverRatio"`
}

type jsonBackendServiceFailoverPolicy BackendServiceFailoverPolicy

func (r *BackendServiceFailoverPolicy) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceFailoverPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceFailoverPolicy
	} else {

		r.DisableConnectionDrainOnFailover = res.DisableConnectionDrainOnFailover

		r.DropTrafficIfUnhealthy = res.DropTrafficIfUnhealthy

		r.FailoverRatio = res.FailoverRatio

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceFailoverPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceFailoverPolicy *BackendServiceFailoverPolicy = &BackendServiceFailoverPolicy{empty: true}

func (r *BackendServiceFailoverPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceFailoverPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceConnectionDraining struct {
	empty              bool   `json:"-"`
	DrainingTimeoutSec *int64 `json:"drainingTimeoutSec"`
}

type jsonBackendServiceConnectionDraining BackendServiceConnectionDraining

func (r *BackendServiceConnectionDraining) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceConnectionDraining
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceConnectionDraining
	} else {

		r.DrainingTimeoutSec = res.DrainingTimeoutSec

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceConnectionDraining is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceConnectionDraining *BackendServiceConnectionDraining = &BackendServiceConnectionDraining{empty: true}

func (r *BackendServiceConnectionDraining) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceConnectionDraining) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceIap struct {
	empty                    bool    `json:"-"`
	Enabled                  *bool   `json:"enabled"`
	OAuth2ClientId           *string `json:"oauth2ClientId"`
	OAuth2ClientSecret       *string `json:"oauth2ClientSecret"`
	OAuth2ClientSecretSha256 *string `json:"oauth2ClientSecretSha256"`
}

type jsonBackendServiceIap BackendServiceIap

func (r *BackendServiceIap) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceIap
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceIap
	} else {

		r.Enabled = res.Enabled

		r.OAuth2ClientId = res.OAuth2ClientId

		r.OAuth2ClientSecret = res.OAuth2ClientSecret

		r.OAuth2ClientSecretSha256 = res.OAuth2ClientSecretSha256

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceIap is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceIap *BackendServiceIap = &BackendServiceIap{empty: true}

func (r *BackendServiceIap) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceIap) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceCdnPolicy struct {
	empty                       bool                                                 `json:"-"`
	CacheKeyPolicy              *BackendServiceCdnPolicyCacheKeyPolicy               `json:"cacheKeyPolicy"`
	SignedUrlKeyNames           []string                                             `json:"signedUrlKeyNames"`
	SignedUrlCacheMaxAgeSec     *int64                                               `json:"signedUrlCacheMaxAgeSec"`
	RequestCoalescing           *bool                                                `json:"requestCoalescing"`
	CacheMode                   *BackendServiceCdnPolicyCacheModeEnum                `json:"cacheMode"`
	DefaultTtl                  *int64                                               `json:"defaultTtl"`
	MaxTtl                      *int64                                               `json:"maxTtl"`
	ClientTtl                   *int64                                               `json:"clientTtl"`
	NegativeCaching             *bool                                                `json:"negativeCaching"`
	NegativeCachingPolicy       []BackendServiceCdnPolicyNegativeCachingPolicy       `json:"negativeCachingPolicy"`
	BypassCacheOnRequestHeaders []BackendServiceCdnPolicyBypassCacheOnRequestHeaders `json:"bypassCacheOnRequestHeaders"`
	ServeWhileStale             *int64                                               `json:"serveWhileStale"`
}

type jsonBackendServiceCdnPolicy BackendServiceCdnPolicy

func (r *BackendServiceCdnPolicy) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceCdnPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceCdnPolicy
	} else {

		r.CacheKeyPolicy = res.CacheKeyPolicy

		r.SignedUrlKeyNames = res.SignedUrlKeyNames

		r.SignedUrlCacheMaxAgeSec = res.SignedUrlCacheMaxAgeSec

		r.RequestCoalescing = res.RequestCoalescing

		r.CacheMode = res.CacheMode

		r.DefaultTtl = res.DefaultTtl

		r.MaxTtl = res.MaxTtl

		r.ClientTtl = res.ClientTtl

		r.NegativeCaching = res.NegativeCaching

		r.NegativeCachingPolicy = res.NegativeCachingPolicy

		r.BypassCacheOnRequestHeaders = res.BypassCacheOnRequestHeaders

		r.ServeWhileStale = res.ServeWhileStale

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceCdnPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceCdnPolicy *BackendServiceCdnPolicy = &BackendServiceCdnPolicy{empty: true}

func (r *BackendServiceCdnPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceCdnPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceCdnPolicyCacheKeyPolicy struct {
	empty                bool     `json:"-"`
	IncludeProtocol      *bool    `json:"includeProtocol"`
	IncludeHost          *bool    `json:"includeHost"`
	IncludeQueryString   *bool    `json:"includeQueryString"`
	QueryStringWhitelist []string `json:"queryStringWhitelist"`
	QueryStringBlacklist []string `json:"queryStringBlacklist"`
	IncludeHttpHeaders   []string `json:"includeHttpHeaders"`
	IncludeNamedCookies  []string `json:"includeNamedCookies"`
}

type jsonBackendServiceCdnPolicyCacheKeyPolicy BackendServiceCdnPolicyCacheKeyPolicy

func (r *BackendServiceCdnPolicyCacheKeyPolicy) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceCdnPolicyCacheKeyPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceCdnPolicyCacheKeyPolicy
	} else {

		r.IncludeProtocol = res.IncludeProtocol

		r.IncludeHost = res.IncludeHost

		r.IncludeQueryString = res.IncludeQueryString

		r.QueryStringWhitelist = res.QueryStringWhitelist

		r.QueryStringBlacklist = res.QueryStringBlacklist

		r.IncludeHttpHeaders = res.IncludeHttpHeaders

		r.IncludeNamedCookies = res.IncludeNamedCookies

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceCdnPolicyCacheKeyPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceCdnPolicyCacheKeyPolicy *BackendServiceCdnPolicyCacheKeyPolicy = &BackendServiceCdnPolicyCacheKeyPolicy{empty: true}

func (r *BackendServiceCdnPolicyCacheKeyPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceCdnPolicyCacheKeyPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceCdnPolicyNegativeCachingPolicy struct {
	empty bool   `json:"-"`
	Code  *int64 `json:"code"`
	Ttl   *int64 `json:"ttl"`
}

type jsonBackendServiceCdnPolicyNegativeCachingPolicy BackendServiceCdnPolicyNegativeCachingPolicy

func (r *BackendServiceCdnPolicyNegativeCachingPolicy) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceCdnPolicyNegativeCachingPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceCdnPolicyNegativeCachingPolicy
	} else {

		r.Code = res.Code

		r.Ttl = res.Ttl

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceCdnPolicyNegativeCachingPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceCdnPolicyNegativeCachingPolicy *BackendServiceCdnPolicyNegativeCachingPolicy = &BackendServiceCdnPolicyNegativeCachingPolicy{empty: true}

func (r *BackendServiceCdnPolicyNegativeCachingPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceCdnPolicyNegativeCachingPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceCdnPolicyBypassCacheOnRequestHeaders struct {
	empty      bool    `json:"-"`
	HeaderName *string `json:"headerName"`
}

type jsonBackendServiceCdnPolicyBypassCacheOnRequestHeaders BackendServiceCdnPolicyBypassCacheOnRequestHeaders

func (r *BackendServiceCdnPolicyBypassCacheOnRequestHeaders) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceCdnPolicyBypassCacheOnRequestHeaders
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceCdnPolicyBypassCacheOnRequestHeaders
	} else {

		r.HeaderName = res.HeaderName

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceCdnPolicyBypassCacheOnRequestHeaders is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceCdnPolicyBypassCacheOnRequestHeaders *BackendServiceCdnPolicyBypassCacheOnRequestHeaders = &BackendServiceCdnPolicyBypassCacheOnRequestHeaders{empty: true}

func (r *BackendServiceCdnPolicyBypassCacheOnRequestHeaders) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceCdnPolicyBypassCacheOnRequestHeaders) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceLogConfig struct {
	empty      bool     `json:"-"`
	Enable     *bool    `json:"enable"`
	SampleRate *float64 `json:"sampleRate"`
}

type jsonBackendServiceLogConfig BackendServiceLogConfig

func (r *BackendServiceLogConfig) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceLogConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceLogConfig
	} else {

		r.Enable = res.Enable

		r.SampleRate = res.SampleRate

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceLogConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceLogConfig *BackendServiceLogConfig = &BackendServiceLogConfig{empty: true}

func (r *BackendServiceLogConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceLogConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceSecuritySettings struct {
	empty           bool     `json:"-"`
	ClientTlsPolicy *string  `json:"clientTlsPolicy"`
	Authentication  *string  `json:"authentication"`
	SubjectAltNames []string `json:"subjectAltNames"`
}

type jsonBackendServiceSecuritySettings BackendServiceSecuritySettings

func (r *BackendServiceSecuritySettings) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceSecuritySettings
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceSecuritySettings
	} else {

		r.ClientTlsPolicy = res.ClientTlsPolicy

		r.Authentication = res.Authentication

		r.SubjectAltNames = res.SubjectAltNames

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceSecuritySettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceSecuritySettings *BackendServiceSecuritySettings = &BackendServiceSecuritySettings{empty: true}

func (r *BackendServiceSecuritySettings) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceSecuritySettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceConsistentHash struct {
	empty           bool                                    `json:"-"`
	HttpCookie      *BackendServiceConsistentHashHttpCookie `json:"httpCookie"`
	HttpHeaderName  *string                                 `json:"httpHeaderName"`
	MinimumRingSize *int64                                  `json:"minimumRingSize"`
}

type jsonBackendServiceConsistentHash BackendServiceConsistentHash

func (r *BackendServiceConsistentHash) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceConsistentHash
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceConsistentHash
	} else {

		r.HttpCookie = res.HttpCookie

		r.HttpHeaderName = res.HttpHeaderName

		r.MinimumRingSize = res.MinimumRingSize

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceConsistentHash is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceConsistentHash *BackendServiceConsistentHash = &BackendServiceConsistentHash{empty: true}

func (r *BackendServiceConsistentHash) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceConsistentHash) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceConsistentHashHttpCookie struct {
	empty bool                                       `json:"-"`
	Name  *string                                    `json:"name"`
	Path  *string                                    `json:"path"`
	Ttl   *BackendServiceConsistentHashHttpCookieTtl `json:"ttl"`
}

type jsonBackendServiceConsistentHashHttpCookie BackendServiceConsistentHashHttpCookie

func (r *BackendServiceConsistentHashHttpCookie) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceConsistentHashHttpCookie
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceConsistentHashHttpCookie
	} else {

		r.Name = res.Name

		r.Path = res.Path

		r.Ttl = res.Ttl

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceConsistentHashHttpCookie is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceConsistentHashHttpCookie *BackendServiceConsistentHashHttpCookie = &BackendServiceConsistentHashHttpCookie{empty: true}

func (r *BackendServiceConsistentHashHttpCookie) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceConsistentHashHttpCookie) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceConsistentHashHttpCookieTtl struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonBackendServiceConsistentHashHttpCookieTtl BackendServiceConsistentHashHttpCookieTtl

func (r *BackendServiceConsistentHashHttpCookieTtl) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceConsistentHashHttpCookieTtl
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceConsistentHashHttpCookieTtl
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceConsistentHashHttpCookieTtl is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceConsistentHashHttpCookieTtl *BackendServiceConsistentHashHttpCookieTtl = &BackendServiceConsistentHashHttpCookieTtl{empty: true}

func (r *BackendServiceConsistentHashHttpCookieTtl) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceConsistentHashHttpCookieTtl) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceCircuitBreakers struct {
	empty                    bool                                         `json:"-"`
	ConnectTimeout           *BackendServiceCircuitBreakersConnectTimeout `json:"connectTimeout"`
	MaxRequestsPerConnection *int64                                       `json:"maxRequestsPerConnection"`
	MaxConnections           *int64                                       `json:"maxConnections"`
	MaxPendingRequests       *int64                                       `json:"maxPendingRequests"`
	MaxRequests              *int64                                       `json:"maxRequests"`
	MaxRetries               *int64                                       `json:"maxRetries"`
}

type jsonBackendServiceCircuitBreakers BackendServiceCircuitBreakers

func (r *BackendServiceCircuitBreakers) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceCircuitBreakers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceCircuitBreakers
	} else {

		r.ConnectTimeout = res.ConnectTimeout

		r.MaxRequestsPerConnection = res.MaxRequestsPerConnection

		r.MaxConnections = res.MaxConnections

		r.MaxPendingRequests = res.MaxPendingRequests

		r.MaxRequests = res.MaxRequests

		r.MaxRetries = res.MaxRetries

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceCircuitBreakers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceCircuitBreakers *BackendServiceCircuitBreakers = &BackendServiceCircuitBreakers{empty: true}

func (r *BackendServiceCircuitBreakers) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceCircuitBreakers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceCircuitBreakersConnectTimeout struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonBackendServiceCircuitBreakersConnectTimeout BackendServiceCircuitBreakersConnectTimeout

func (r *BackendServiceCircuitBreakersConnectTimeout) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceCircuitBreakersConnectTimeout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceCircuitBreakersConnectTimeout
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceCircuitBreakersConnectTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceCircuitBreakersConnectTimeout *BackendServiceCircuitBreakersConnectTimeout = &BackendServiceCircuitBreakersConnectTimeout{empty: true}

func (r *BackendServiceCircuitBreakersConnectTimeout) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceCircuitBreakersConnectTimeout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceOutlierDetection struct {
	empty                              bool                                            `json:"-"`
	ConsecutiveErrors                  *int64                                          `json:"consecutiveErrors"`
	Interval                           *BackendServiceOutlierDetectionInterval         `json:"interval"`
	BaseEjectionTime                   *BackendServiceOutlierDetectionBaseEjectionTime `json:"baseEjectionTime"`
	MaxEjectionPercent                 *int64                                          `json:"maxEjectionPercent"`
	EnforcingConsecutiveErrors         *int64                                          `json:"enforcingConsecutiveErrors"`
	EnforcingSuccessRate               *int64                                          `json:"enforcingSuccessRate"`
	SuccessRateMinimumHosts            *int64                                          `json:"successRateMinimumHosts"`
	SuccessRateRequestVolume           *int64                                          `json:"successRateRequestVolume"`
	SuccessRateStdevFactor             *int64                                          `json:"successRateStdevFactor"`
	ConsecutiveGatewayFailure          *int64                                          `json:"consecutiveGatewayFailure"`
	EnforcingConsecutiveGatewayFailure *int64                                          `json:"enforcingConsecutiveGatewayFailure"`
}

type jsonBackendServiceOutlierDetection BackendServiceOutlierDetection

func (r *BackendServiceOutlierDetection) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceOutlierDetection
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceOutlierDetection
	} else {

		r.ConsecutiveErrors = res.ConsecutiveErrors

		r.Interval = res.Interval

		r.BaseEjectionTime = res.BaseEjectionTime

		r.MaxEjectionPercent = res.MaxEjectionPercent

		r.EnforcingConsecutiveErrors = res.EnforcingConsecutiveErrors

		r.EnforcingSuccessRate = res.EnforcingSuccessRate

		r.SuccessRateMinimumHosts = res.SuccessRateMinimumHosts

		r.SuccessRateRequestVolume = res.SuccessRateRequestVolume

		r.SuccessRateStdevFactor = res.SuccessRateStdevFactor

		r.ConsecutiveGatewayFailure = res.ConsecutiveGatewayFailure

		r.EnforcingConsecutiveGatewayFailure = res.EnforcingConsecutiveGatewayFailure

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceOutlierDetection is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceOutlierDetection *BackendServiceOutlierDetection = &BackendServiceOutlierDetection{empty: true}

func (r *BackendServiceOutlierDetection) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceOutlierDetection) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceOutlierDetectionInterval struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonBackendServiceOutlierDetectionInterval BackendServiceOutlierDetectionInterval

func (r *BackendServiceOutlierDetectionInterval) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceOutlierDetectionInterval
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceOutlierDetectionInterval
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceOutlierDetectionInterval is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceOutlierDetectionInterval *BackendServiceOutlierDetectionInterval = &BackendServiceOutlierDetectionInterval{empty: true}

func (r *BackendServiceOutlierDetectionInterval) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceOutlierDetectionInterval) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceOutlierDetectionBaseEjectionTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonBackendServiceOutlierDetectionBaseEjectionTime BackendServiceOutlierDetectionBaseEjectionTime

func (r *BackendServiceOutlierDetectionBaseEjectionTime) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceOutlierDetectionBaseEjectionTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceOutlierDetectionBaseEjectionTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceOutlierDetectionBaseEjectionTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceOutlierDetectionBaseEjectionTime *BackendServiceOutlierDetectionBaseEjectionTime = &BackendServiceOutlierDetectionBaseEjectionTime{empty: true}

func (r *BackendServiceOutlierDetectionBaseEjectionTime) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceOutlierDetectionBaseEjectionTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceSubsetting struct {
	empty  bool                                `json:"-"`
	Policy *BackendServiceSubsettingPolicyEnum `json:"policy"`
}

type jsonBackendServiceSubsetting BackendServiceSubsetting

func (r *BackendServiceSubsetting) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceSubsetting
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceSubsetting
	} else {

		r.Policy = res.Policy

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceSubsetting is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceSubsetting *BackendServiceSubsetting = &BackendServiceSubsetting{empty: true}

func (r *BackendServiceSubsetting) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceSubsetting) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceConnectionTrackingPolicy struct {
	empty                                    bool                                                                                `json:"-"`
	TrackingMode                             *BackendServiceConnectionTrackingPolicyTrackingModeEnum                             `json:"trackingMode"`
	ConnectionPersistenceOnUnhealthyBackends *BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum `json:"connectionPersistenceOnUnhealthyBackends"`
	IdleTimeoutSec                           *int64                                                                              `json:"idleTimeoutSec"`
}

type jsonBackendServiceConnectionTrackingPolicy BackendServiceConnectionTrackingPolicy

func (r *BackendServiceConnectionTrackingPolicy) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceConnectionTrackingPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceConnectionTrackingPolicy
	} else {

		r.TrackingMode = res.TrackingMode

		r.ConnectionPersistenceOnUnhealthyBackends = res.ConnectionPersistenceOnUnhealthyBackends

		r.IdleTimeoutSec = res.IdleTimeoutSec

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceConnectionTrackingPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceConnectionTrackingPolicy *BackendServiceConnectionTrackingPolicy = &BackendServiceConnectionTrackingPolicy{empty: true}

func (r *BackendServiceConnectionTrackingPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceConnectionTrackingPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BackendServiceMaxStreamDuration struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonBackendServiceMaxStreamDuration BackendServiceMaxStreamDuration

func (r *BackendServiceMaxStreamDuration) UnmarshalJSON(data []byte) error {
	var res jsonBackendServiceMaxStreamDuration
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBackendServiceMaxStreamDuration
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this BackendServiceMaxStreamDuration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBackendServiceMaxStreamDuration *BackendServiceMaxStreamDuration = &BackendServiceMaxStreamDuration{empty: true}

func (r *BackendServiceMaxStreamDuration) String() string {
	return dcl.SprintResource(r)
}

func (r *BackendServiceMaxStreamDuration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *BackendService) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "BackendService",
		Version: "beta",
	}
}

const BackendServiceMaxPage = -1

type BackendServiceList struct {
	Items []*BackendService

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *BackendServiceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *BackendServiceList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listBackendService(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListBackendService(ctx context.Context, project, location string) (*BackendServiceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListBackendServiceWithMaxResults(ctx, project, location, BackendServiceMaxPage)

}

func (c *Client) ListBackendServiceWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*BackendServiceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listBackendService(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &BackendServiceList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetBackendService(ctx context.Context, r *BackendService) (*BackendService, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getBackendServiceRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalBackendService(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeBackendServiceNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteBackendService(ctx context.Context, r *BackendService) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("BackendService resource is nil")
	}
	c.Config.Logger.Info("Deleting BackendService...")
	deleteOp := deleteBackendServiceOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllBackendService deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllBackendService(ctx context.Context, project, location string, filter func(*BackendService) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListBackendService(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllBackendService(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllBackendService(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyBackendService(ctx context.Context, rawDesired *BackendService, opts ...dcl.ApplyOption) (*BackendService, error) {
	c.Config.Logger.Info("Beginning ApplyBackendService...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.backendServiceDiffsForRawDesired(ctx, rawDesired, opts...)
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
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []backendServiceApiOperation
	if create {
		ops = append(ops, &createBackendServiceOperation{})
	} else if recreate {

		ops = append(ops, &deleteBackendServiceOperation{})

		ops = append(ops, &createBackendServiceOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeBackendServiceDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetBackendService(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createBackendServiceOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapBackendService(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeBackendServiceNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeBackendServiceNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeBackendServiceDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffBackendService(c, newDesired, newState)
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
