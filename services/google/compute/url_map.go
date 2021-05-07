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

type UrlMap struct {
	DefaultRouteAction *UrlMapDefaultRouteAction `json:"defaultRouteAction"`
	DefaultService     *string                   `json:"defaultService"`
	DefaultUrlRedirect *UrlMapDefaultUrlRedirect `json:"defaultUrlRedirect"`
	Description        *string                   `json:"description"`
	SelfLink           *string                   `json:"selfLink"`
	HeaderAction       *UrlMapHeaderAction       `json:"headerAction"`
	HostRule           []UrlMapHostRule          `json:"hostRule"`
	Name               *string                   `json:"name"`
	PathMatcher        []UrlMapPathMatcher       `json:"pathMatcher"`
	Region             *string                   `json:"region"`
	Test               []UrlMapTest              `json:"test"`
	Project            *string                   `json:"project"`
}

func (r *UrlMap) String() string {
	return dcl.SprintResource(r)
}

// The enum UrlMapDefaultUrlRedirectRedirectResponseCodeEnum.
type UrlMapDefaultUrlRedirectRedirectResponseCodeEnum string

// UrlMapDefaultUrlRedirectRedirectResponseCodeEnumRef returns a *UrlMapDefaultUrlRedirectRedirectResponseCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func UrlMapDefaultUrlRedirectRedirectResponseCodeEnumRef(s string) *UrlMapDefaultUrlRedirectRedirectResponseCodeEnum {
	if s == "" {
		return nil
	}

	v := UrlMapDefaultUrlRedirectRedirectResponseCodeEnum(s)
	return &v
}

func (v UrlMapDefaultUrlRedirectRedirectResponseCodeEnum) Validate() error {
	for _, s := range []string{"MOVED_PERMANENTLY_DEFAULT", "FOUND", "SEE_OTHER", "TEMPORARY_REDIRECT", "PERMANENT_REDIRECT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UrlMapDefaultUrlRedirectRedirectResponseCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum.
type UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum string

// UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumRef returns a *UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumRef(s string) *UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum {
	if s == "" {
		return nil
	}

	v := UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(s)
	return &v
}

func (v UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum) Validate() error {
	for _, s := range []string{"MOVED_PERMANENTLY_DEFAULT", "FOUND", "SEE_OTHER", "TEMPORARY_REDIRECT", "PERMANENT_REDIRECT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum.
type UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum string

// UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumRef returns a *UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumRef(s string) *UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum {
	if s == "" {
		return nil
	}

	v := UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(s)
	return &v
}

func (v UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum) Validate() error {
	for _, s := range []string{"MOVED_PERMANENTLY_DEFAULT", "FOUND", "SEE_OTHER", "TEMPORARY_REDIRECT", "PERMANENT_REDIRECT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum.
type UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum string

// UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumRef returns a *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum with the value of string s
// If the empty string is provided, nil is returned.
func UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumRef(s string) *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum {
	if s == "" {
		return nil
	}

	v := UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(s)
	return &v
}

func (v UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum) Validate() error {
	for _, s := range []string{"NOT_SET", "MATCH_ALL", "MATCH_ANY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum.
type UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum string

// UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumRef returns a *UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumRef(s string) *UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum {
	if s == "" {
		return nil
	}

	v := UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(s)
	return &v
}

func (v UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum) Validate() error {
	for _, s := range []string{"MOVED_PERMANENTLY_DEFAULT", "FOUND", "SEE_OTHER", "TEMPORARY_REDIRECT", "PERMANENT_REDIRECT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type UrlMapDefaultRouteAction struct {
	empty                  bool                                             `json:"-"`
	WeightedBackendService []UrlMapDefaultRouteActionWeightedBackendService `json:"weightedBackendService"`
	UrlRewrite             *UrlMapDefaultRouteActionUrlRewrite              `json:"urlRewrite"`
	Timeout                *UrlMapDefaultRouteActionTimeout                 `json:"timeout"`
	RetryPolicy            *UrlMapDefaultRouteActionRetryPolicy             `json:"retryPolicy"`
	RequestMirrorPolicy    *UrlMapDefaultRouteActionRequestMirrorPolicy     `json:"requestMirrorPolicy"`
	CorsPolicy             *UrlMapDefaultRouteActionCorsPolicy              `json:"corsPolicy"`
	FaultInjectionPolicy   *UrlMapDefaultRouteActionFaultInjectionPolicy    `json:"faultInjectionPolicy"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteAction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteAction *UrlMapDefaultRouteAction = &UrlMapDefaultRouteAction{empty: true}

func (r *UrlMapDefaultRouteAction) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionWeightedBackendService struct {
	empty          bool                `json:"-"`
	BackendService *string             `json:"backendService"`
	Weight         *int64              `json:"weight"`
	HeaderAction   *UrlMapHeaderAction `json:"headerAction"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionWeightedBackendService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionWeightedBackendService *UrlMapDefaultRouteActionWeightedBackendService = &UrlMapDefaultRouteActionWeightedBackendService{empty: true}

func (r *UrlMapDefaultRouteActionWeightedBackendService) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionWeightedBackendService) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapHeaderAction struct {
	empty                   bool                                     `json:"-"`
	RequestHeadersToRemove  []string                                 `json:"requestHeadersToRemove"`
	RequestHeadersToAdd     []UrlMapHeaderActionRequestHeadersToAdd  `json:"requestHeadersToAdd"`
	ResponseHeadersToRemove []string                                 `json:"responseHeadersToRemove"`
	ResponseHeadersToAdd    []UrlMapHeaderActionResponseHeadersToAdd `json:"responseHeadersToAdd"`
}

// This object is used to assert a desired state where this UrlMapHeaderAction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapHeaderAction *UrlMapHeaderAction = &UrlMapHeaderAction{empty: true}

func (r *UrlMapHeaderAction) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapHeaderAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapHeaderActionRequestHeadersToAdd struct {
	empty       bool    `json:"-"`
	HeaderName  *string `json:"headerName"`
	HeaderValue *string `json:"headerValue"`
	Replace     *bool   `json:"replace"`
}

// This object is used to assert a desired state where this UrlMapHeaderActionRequestHeadersToAdd is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapHeaderActionRequestHeadersToAdd *UrlMapHeaderActionRequestHeadersToAdd = &UrlMapHeaderActionRequestHeadersToAdd{empty: true}

func (r *UrlMapHeaderActionRequestHeadersToAdd) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapHeaderActionRequestHeadersToAdd) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapHeaderActionResponseHeadersToAdd struct {
	empty       bool    `json:"-"`
	HeaderName  *string `json:"headerName"`
	HeaderValue *string `json:"headerValue"`
	Replace     *bool   `json:"replace"`
}

// This object is used to assert a desired state where this UrlMapHeaderActionResponseHeadersToAdd is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapHeaderActionResponseHeadersToAdd *UrlMapHeaderActionResponseHeadersToAdd = &UrlMapHeaderActionResponseHeadersToAdd{empty: true}

func (r *UrlMapHeaderActionResponseHeadersToAdd) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapHeaderActionResponseHeadersToAdd) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionUrlRewrite struct {
	empty             bool    `json:"-"`
	PathPrefixRewrite *string `json:"pathPrefixRewrite"`
	HostRewrite       *string `json:"hostRewrite"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionUrlRewrite is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionUrlRewrite *UrlMapDefaultRouteActionUrlRewrite = &UrlMapDefaultRouteActionUrlRewrite{empty: true}

func (r *UrlMapDefaultRouteActionUrlRewrite) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionUrlRewrite) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionTimeout struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionTimeout *UrlMapDefaultRouteActionTimeout = &UrlMapDefaultRouteActionTimeout{empty: true}

func (r *UrlMapDefaultRouteActionTimeout) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionTimeout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionRetryPolicy struct {
	empty          bool                                              `json:"-"`
	RetryCondition []string                                          `json:"retryCondition"`
	NumRetries     *int64                                            `json:"numRetries"`
	PerTryTimeout  *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout `json:"perTryTimeout"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionRetryPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionRetryPolicy *UrlMapDefaultRouteActionRetryPolicy = &UrlMapDefaultRouteActionRetryPolicy{empty: true}

func (r *UrlMapDefaultRouteActionRetryPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionRetryPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionRetryPolicyPerTryTimeout struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionRetryPolicyPerTryTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionRetryPolicyPerTryTimeout *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout = &UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{empty: true}

func (r *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionRequestMirrorPolicy struct {
	empty          bool    `json:"-"`
	BackendService *string `json:"backendService"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionRequestMirrorPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionRequestMirrorPolicy *UrlMapDefaultRouteActionRequestMirrorPolicy = &UrlMapDefaultRouteActionRequestMirrorPolicy{empty: true}

func (r *UrlMapDefaultRouteActionRequestMirrorPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionRequestMirrorPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionCorsPolicy struct {
	empty            bool     `json:"-"`
	AllowOrigin      []string `json:"allowOrigin"`
	AllowOriginRegex []string `json:"allowOriginRegex"`
	AllowMethod      []string `json:"allowMethod"`
	AllowHeader      []string `json:"allowHeader"`
	ExposeHeader     []string `json:"exposeHeader"`
	MaxAge           *int64   `json:"maxAge"`
	AllowCredentials *bool    `json:"allowCredentials"`
	Disabled         *bool    `json:"disabled"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionCorsPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionCorsPolicy *UrlMapDefaultRouteActionCorsPolicy = &UrlMapDefaultRouteActionCorsPolicy{empty: true}

func (r *UrlMapDefaultRouteActionCorsPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionCorsPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionFaultInjectionPolicy struct {
	empty bool                                               `json:"-"`
	Delay *UrlMapDefaultRouteActionFaultInjectionPolicyDelay `json:"delay"`
	Abort *UrlMapDefaultRouteActionFaultInjectionPolicyAbort `json:"abort"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionFaultInjectionPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionFaultInjectionPolicy *UrlMapDefaultRouteActionFaultInjectionPolicy = &UrlMapDefaultRouteActionFaultInjectionPolicy{empty: true}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionFaultInjectionPolicyDelay struct {
	empty      bool                                                         `json:"-"`
	FixedDelay *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay `json:"fixedDelay"`
	Percentage *float64                                                     `json:"percentage"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionFaultInjectionPolicyDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelay *UrlMapDefaultRouteActionFaultInjectionPolicyDelay = &UrlMapDefaultRouteActionFaultInjectionPolicyDelay{empty: true}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelay) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelay) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay = &UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{empty: true}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultRouteActionFaultInjectionPolicyAbort struct {
	empty      bool     `json:"-"`
	HttpStatus *int64   `json:"httpStatus"`
	Percentage *float64 `json:"percentage"`
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionFaultInjectionPolicyAbort is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionFaultInjectionPolicyAbort *UrlMapDefaultRouteActionFaultInjectionPolicyAbort = &UrlMapDefaultRouteActionFaultInjectionPolicyAbort{empty: true}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyAbort) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyAbort) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapDefaultUrlRedirect struct {
	empty                bool                                              `json:"-"`
	HostRedirect         *string                                           `json:"hostRedirect"`
	PathRedirect         *string                                           `json:"pathRedirect"`
	PrefixRedirect       *string                                           `json:"prefixRedirect"`
	RedirectResponseCode *UrlMapDefaultUrlRedirectRedirectResponseCodeEnum `json:"redirectResponseCode"`
	HttpsRedirect        *bool                                             `json:"httpsRedirect"`
	StripQuery           *bool                                             `json:"stripQuery"`
}

// This object is used to assert a desired state where this UrlMapDefaultUrlRedirect is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultUrlRedirect *UrlMapDefaultUrlRedirect = &UrlMapDefaultUrlRedirect{empty: true}

func (r *UrlMapDefaultUrlRedirect) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapDefaultUrlRedirect) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapHostRule struct {
	empty       bool     `json:"-"`
	Description *string  `json:"description"`
	Host        []string `json:"host"`
	PathMatcher *string  `json:"pathMatcher"`
}

// This object is used to assert a desired state where this UrlMapHostRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapHostRule *UrlMapHostRule = &UrlMapHostRule{empty: true}

func (r *UrlMapHostRule) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapHostRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcher struct {
	empty              bool                                 `json:"-"`
	Name               *string                              `json:"name"`
	Description        *string                              `json:"description"`
	DefaultService     *string                              `json:"defaultService"`
	DefaultRouteAction *UrlMapDefaultRouteAction            `json:"defaultRouteAction"`
	DefaultUrlRedirect *UrlMapPathMatcherDefaultUrlRedirect `json:"defaultUrlRedirect"`
	PathRule           []UrlMapPathMatcherPathRule          `json:"pathRule"`
	RouteRule          []UrlMapPathMatcherRouteRule         `json:"routeRule"`
	HeaderAction       *UrlMapHeaderAction                  `json:"headerAction"`
}

// This object is used to assert a desired state where this UrlMapPathMatcher is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcher *UrlMapPathMatcher = &UrlMapPathMatcher{empty: true}

func (r *UrlMapPathMatcher) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcher) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherDefaultUrlRedirect struct {
	empty                bool                                                         `json:"-"`
	HostRedirect         *string                                                      `json:"hostRedirect"`
	PathRedirect         *string                                                      `json:"pathRedirect"`
	PrefixRedirect       *string                                                      `json:"prefixRedirect"`
	RedirectResponseCode *UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum `json:"redirectResponseCode"`
	HttpsRedirect        *bool                                                        `json:"httpsRedirect"`
	StripQuery           *bool                                                        `json:"stripQuery"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherDefaultUrlRedirect is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherDefaultUrlRedirect *UrlMapPathMatcherDefaultUrlRedirect = &UrlMapPathMatcherDefaultUrlRedirect{empty: true}

func (r *UrlMapPathMatcherDefaultUrlRedirect) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherDefaultUrlRedirect) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRule struct {
	empty          bool                                  `json:"-"`
	BackendService *string                               `json:"backendService"`
	RouteAction    *UrlMapPathMatcherPathRuleRouteAction `json:"routeAction"`
	UrlRedirect    *UrlMapPathMatcherPathRuleUrlRedirect `json:"urlRedirect"`
	Path           []string                              `json:"path"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRule *UrlMapPathMatcherPathRule = &UrlMapPathMatcherPathRule{empty: true}

func (r *UrlMapPathMatcherPathRule) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteAction struct {
	empty                  bool                                                         `json:"-"`
	WeightedBackendService []UrlMapPathMatcherPathRuleRouteActionWeightedBackendService `json:"weightedBackendService"`
	UrlRewrite             *UrlMapPathMatcherPathRuleRouteActionUrlRewrite              `json:"urlRewrite"`
	Timeout                *UrlMapPathMatcherPathRuleRouteActionTimeout                 `json:"timeout"`
	RetryPolicy            *UrlMapPathMatcherPathRuleRouteActionRetryPolicy             `json:"retryPolicy"`
	RequestMirrorPolicy    *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy     `json:"requestMirrorPolicy"`
	CorsPolicy             *UrlMapPathMatcherPathRuleRouteActionCorsPolicy              `json:"corsPolicy"`
	FaultInjectionPolicy   *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy    `json:"faultInjectionPolicy"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteAction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteAction *UrlMapPathMatcherPathRuleRouteAction = &UrlMapPathMatcherPathRuleRouteAction{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteAction) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionWeightedBackendService struct {
	empty          bool                `json:"-"`
	BackendService *string             `json:"backendService"`
	Weight         *int64              `json:"weight"`
	HeaderAction   *UrlMapHeaderAction `json:"headerAction"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionWeightedBackendService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionWeightedBackendService *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService = &UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionUrlRewrite struct {
	empty             bool    `json:"-"`
	PathPrefixRewrite *string `json:"pathPrefixRewrite"`
	HostRewrite       *string `json:"hostRewrite"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionUrlRewrite is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionUrlRewrite *UrlMapPathMatcherPathRuleRouteActionUrlRewrite = &UrlMapPathMatcherPathRuleRouteActionUrlRewrite{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionUrlRewrite) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionUrlRewrite) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionTimeout struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionTimeout *UrlMapPathMatcherPathRuleRouteActionTimeout = &UrlMapPathMatcherPathRuleRouteActionTimeout{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionTimeout) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionTimeout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionRetryPolicy struct {
	empty          bool                                                          `json:"-"`
	RetryCondition []string                                                      `json:"retryCondition"`
	NumRetries     *int64                                                        `json:"numRetries"`
	PerTryTimeout  *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout `json:"perTryTimeout"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionRetryPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicy *UrlMapPathMatcherPathRuleRouteActionRetryPolicy = &UrlMapPathMatcherPathRuleRouteActionRetryPolicy{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout = &UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy struct {
	empty          bool    `json:"-"`
	BackendService *string `json:"backendService"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy = &UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionCorsPolicy struct {
	empty            bool     `json:"-"`
	AllowOrigin      []string `json:"allowOrigin"`
	AllowOriginRegex []string `json:"allowOriginRegex"`
	AllowMethod      []string `json:"allowMethod"`
	AllowHeader      []string `json:"allowHeader"`
	ExposeHeader     []string `json:"exposeHeader"`
	MaxAge           *int64   `json:"maxAge"`
	AllowCredentials *bool    `json:"allowCredentials"`
	Disabled         *bool    `json:"disabled"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionCorsPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionCorsPolicy *UrlMapPathMatcherPathRuleRouteActionCorsPolicy = &UrlMapPathMatcherPathRuleRouteActionCorsPolicy{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionCorsPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionCorsPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy struct {
	empty bool                                                           `json:"-"`
	Delay *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay `json:"delay"`
	Abort *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort `json:"abort"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy = &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay struct {
	empty      bool                                                                     `json:"-"`
	FixedDelay *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay `json:"fixedDelay"`
	Percentage *float64                                                                 `json:"percentage"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay = &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay = &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort struct {
	empty      bool     `json:"-"`
	HttpStatus *int64   `json:"httpStatus"`
	Percentage *float64 `json:"percentage"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort = &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherPathRuleUrlRedirect struct {
	empty                bool                                                          `json:"-"`
	HostRedirect         *string                                                       `json:"hostRedirect"`
	PathRedirect         *string                                                       `json:"pathRedirect"`
	PrefixRedirect       *string                                                       `json:"prefixRedirect"`
	RedirectResponseCode *UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum `json:"redirectResponseCode"`
	HttpsRedirect        *bool                                                         `json:"httpsRedirect"`
	StripQuery           *bool                                                         `json:"stripQuery"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleUrlRedirect is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleUrlRedirect *UrlMapPathMatcherPathRuleUrlRedirect = &UrlMapPathMatcherPathRuleUrlRedirect{empty: true}

func (r *UrlMapPathMatcherPathRuleUrlRedirect) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherPathRuleUrlRedirect) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRule struct {
	empty          bool                                   `json:"-"`
	Priority       *int64                                 `json:"priority"`
	Description    *string                                `json:"description"`
	MatchRule      []UrlMapPathMatcherRouteRuleMatchRule  `json:"matchRule"`
	BackendService *string                                `json:"backendService"`
	RouteAction    *UrlMapPathMatcherRouteRuleRouteAction `json:"routeAction"`
	UrlRedirect    *UrlMapPathMatcherRouteRuleUrlRedirect `json:"urlRedirect"`
	HeaderAction   *UrlMapHeaderAction                    `json:"headerAction"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRule *UrlMapPathMatcherRouteRule = &UrlMapPathMatcherRouteRule{empty: true}

func (r *UrlMapPathMatcherRouteRule) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleMatchRule struct {
	empty               bool                                                     `json:"-"`
	PrefixMatch         *string                                                  `json:"prefixMatch"`
	FullPathMatch       *string                                                  `json:"fullPathMatch"`
	RegexMatch          *string                                                  `json:"regexMatch"`
	IgnoreCase          *bool                                                    `json:"ignoreCase"`
	HeaderMatch         []UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch         `json:"headerMatch"`
	QueryParameterMatch []UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch `json:"queryParameterMatch"`
	MetadataFilter      []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter      `json:"metadataFilter"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRule *UrlMapPathMatcherRouteRuleMatchRule = &UrlMapPathMatcherRouteRuleMatchRule{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRule) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleMatchRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch struct {
	empty        bool                                                      `json:"-"`
	HeaderName   *string                                                   `json:"headerName"`
	ExactMatch   *string                                                   `json:"exactMatch"`
	RegexMatch   *string                                                   `json:"regexMatch"`
	RangeMatch   *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch `json:"rangeMatch"`
	PresentMatch *bool                                                     `json:"presentMatch"`
	PrefixMatch  *string                                                   `json:"prefixMatch"`
	SuffixMatch  *string                                                   `json:"suffixMatch"`
	InvertMatch  *bool                                                     `json:"invertMatch"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch = &UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch struct {
	empty      bool   `json:"-"`
	RangeStart *int64 `json:"rangeStart"`
	RangeEnd   *int64 `json:"rangeEnd"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch = &UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch struct {
	empty        bool    `json:"-"`
	Name         *string `json:"name"`
	PresentMatch *bool   `json:"presentMatch"`
	ExactMatch   *string `json:"exactMatch"`
	RegexMatch   *string `json:"regexMatch"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch = &UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter struct {
	empty               bool                                                                      `json:"-"`
	FilterMatchCriteria *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum `json:"filterMatchCriteria"`
	FilterLabel         []UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel            `json:"filterLabel"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter = &UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel = &UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteAction struct {
	empty                  bool                                                          `json:"-"`
	WeightedBackendService []UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService `json:"weightedBackendService"`
	UrlRewrite             *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite              `json:"urlRewrite"`
	Timeout                *UrlMapPathMatcherRouteRuleRouteActionTimeout                 `json:"timeout"`
	RetryPolicy            *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy             `json:"retryPolicy"`
	RequestMirrorPolicy    *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy     `json:"requestMirrorPolicy"`
	CorsPolicy             *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy              `json:"corsPolicy"`
	FaultInjectionPolicy   *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy    `json:"faultInjectionPolicy"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteAction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteAction *UrlMapPathMatcherRouteRuleRouteAction = &UrlMapPathMatcherRouteRuleRouteAction{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteAction) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService struct {
	empty          bool                `json:"-"`
	BackendService *string             `json:"backendService"`
	Weight         *int64              `json:"weight"`
	HeaderAction   *UrlMapHeaderAction `json:"headerAction"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService = &UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionUrlRewrite struct {
	empty             bool    `json:"-"`
	PathPrefixRewrite *string `json:"pathPrefixRewrite"`
	HostRewrite       *string `json:"hostRewrite"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionUrlRewrite is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionUrlRewrite *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite = &UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionTimeout struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionTimeout *UrlMapPathMatcherRouteRuleRouteActionTimeout = &UrlMapPathMatcherRouteRuleRouteActionTimeout{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionTimeout) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionTimeout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionRetryPolicy struct {
	empty          bool                                                           `json:"-"`
	RetryCondition []string                                                       `json:"retryCondition"`
	NumRetries     *int64                                                         `json:"numRetries"`
	PerTryTimeout  *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout `json:"perTryTimeout"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionRetryPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicy *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy = &UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout = &UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy struct {
	empty          bool    `json:"-"`
	BackendService *string `json:"backendService"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy = &UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionCorsPolicy struct {
	empty            bool     `json:"-"`
	AllowOrigin      []string `json:"allowOrigin"`
	AllowOriginRegex []string `json:"allowOriginRegex"`
	AllowMethod      []string `json:"allowMethod"`
	AllowHeader      []string `json:"allowHeader"`
	ExposeHeader     []string `json:"exposeHeader"`
	MaxAge           *int64   `json:"maxAge"`
	AllowCredentials *bool    `json:"allowCredentials"`
	Disabled         *bool    `json:"disabled"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionCorsPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionCorsPolicy *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy = &UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy struct {
	empty bool                                                            `json:"-"`
	Delay *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay `json:"delay"`
	Abort *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort `json:"abort"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy = &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay struct {
	empty      bool                                                                      `json:"-"`
	FixedDelay *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay `json:"fixedDelay"`
	Percentage *float64                                                                  `json:"percentage"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay = &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay = &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort struct {
	empty      bool     `json:"-"`
	HttpStatus *int64   `json:"httpStatus"`
	Percentage *float64 `json:"percentage"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort = &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapPathMatcherRouteRuleUrlRedirect struct {
	empty                bool                                                           `json:"-"`
	HostRedirect         *string                                                        `json:"hostRedirect"`
	PathRedirect         *string                                                        `json:"pathRedirect"`
	PrefixRedirect       *string                                                        `json:"prefixRedirect"`
	RedirectResponseCode *UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum `json:"redirectResponseCode"`
	HttpsRedirect        *bool                                                          `json:"httpsRedirect"`
	StripQuery           *bool                                                          `json:"stripQuery"`
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleUrlRedirect is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleUrlRedirect *UrlMapPathMatcherRouteRuleUrlRedirect = &UrlMapPathMatcherRouteRuleUrlRedirect{empty: true}

func (r *UrlMapPathMatcherRouteRuleUrlRedirect) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapPathMatcherRouteRuleUrlRedirect) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type UrlMapTest struct {
	empty                  bool    `json:"-"`
	Description            *string `json:"description"`
	Host                   *string `json:"host"`
	Path                   *string `json:"path"`
	ExpectedBackendService *string `json:"expectedBackendService"`
}

// This object is used to assert a desired state where this UrlMapTest is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapTest *UrlMapTest = &UrlMapTest{empty: true}

func (r *UrlMapTest) String() string {
	return dcl.SprintResource(r)
}

func (r *UrlMapTest) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *UrlMap) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "UrlMap",
		Version: "compute",
	}
}

const UrlMapMaxPage = -1

type UrlMapList struct {
	Items []*UrlMap

	nextToken string

	pageSize int32

	project string
}

func (l *UrlMapList) HasNext() bool {
	return l.nextToken != ""
}

func (l *UrlMapList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listUrlMap(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListUrlMap(ctx context.Context, project string) (*UrlMapList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListUrlMapWithMaxResults(ctx, project, UrlMapMaxPage)

}

func (c *Client) ListUrlMapWithMaxResults(ctx context.Context, project string, pageSize int32) (*UrlMapList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listUrlMap(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &UrlMapList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetUrlMap(ctx context.Context, r *UrlMap) (*UrlMap, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getUrlMapRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalUrlMap(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeUrlMapNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteUrlMap(ctx context.Context, r *UrlMap) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("UrlMap resource is nil")
	}
	c.Config.Logger.Info("Deleting UrlMap...")
	deleteOp := deleteUrlMapOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllUrlMap deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllUrlMap(ctx context.Context, project string, filter func(*UrlMap) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListUrlMap(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllUrlMap(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllUrlMap(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyUrlMap(ctx context.Context, rawDesired *UrlMap, opts ...dcl.ApplyOption) (*UrlMap, error) {
	c.Config.Logger.Info("Beginning ApplyUrlMap...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.urlMapDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []urlMapApiOperation
	if create {
		ops = append(ops, &createUrlMapOperation{})
	} else if recreate {

		ops = append(ops, &deleteUrlMapOperation{})

		ops = append(ops, &createUrlMapOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeUrlMapDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetUrlMap(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createUrlMapOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapUrlMap(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeUrlMapNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeUrlMapNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeUrlMapDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffUrlMap(c, newDesired, newState)
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
