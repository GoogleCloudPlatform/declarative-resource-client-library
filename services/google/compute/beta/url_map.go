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
	"time"

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

type jsonUrlMapDefaultRouteAction UrlMapDefaultRouteAction

func (r *UrlMapDefaultRouteAction) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteAction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteAction
	} else {

		r.WeightedBackendService = res.WeightedBackendService

		r.UrlRewrite = res.UrlRewrite

		r.Timeout = res.Timeout

		r.RetryPolicy = res.RetryPolicy

		r.RequestMirrorPolicy = res.RequestMirrorPolicy

		r.CorsPolicy = res.CorsPolicy

		r.FaultInjectionPolicy = res.FaultInjectionPolicy

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteAction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteAction *UrlMapDefaultRouteAction = &UrlMapDefaultRouteAction{empty: true}

func (r *UrlMapDefaultRouteAction) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionWeightedBackendService UrlMapDefaultRouteActionWeightedBackendService

func (r *UrlMapDefaultRouteActionWeightedBackendService) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionWeightedBackendService
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionWeightedBackendService
	} else {

		r.BackendService = res.BackendService

		r.Weight = res.Weight

		r.HeaderAction = res.HeaderAction

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionWeightedBackendService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionWeightedBackendService *UrlMapDefaultRouteActionWeightedBackendService = &UrlMapDefaultRouteActionWeightedBackendService{empty: true}

func (r *UrlMapDefaultRouteActionWeightedBackendService) Empty() bool {
	return r.empty
}

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

type jsonUrlMapHeaderAction UrlMapHeaderAction

func (r *UrlMapHeaderAction) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapHeaderAction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapHeaderAction
	} else {

		r.RequestHeadersToRemove = res.RequestHeadersToRemove

		r.RequestHeadersToAdd = res.RequestHeadersToAdd

		r.ResponseHeadersToRemove = res.ResponseHeadersToRemove

		r.ResponseHeadersToAdd = res.ResponseHeadersToAdd

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapHeaderAction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapHeaderAction *UrlMapHeaderAction = &UrlMapHeaderAction{empty: true}

func (r *UrlMapHeaderAction) Empty() bool {
	return r.empty
}

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

type jsonUrlMapHeaderActionRequestHeadersToAdd UrlMapHeaderActionRequestHeadersToAdd

func (r *UrlMapHeaderActionRequestHeadersToAdd) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapHeaderActionRequestHeadersToAdd
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapHeaderActionRequestHeadersToAdd
	} else {

		r.HeaderName = res.HeaderName

		r.HeaderValue = res.HeaderValue

		r.Replace = res.Replace

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapHeaderActionRequestHeadersToAdd is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapHeaderActionRequestHeadersToAdd *UrlMapHeaderActionRequestHeadersToAdd = &UrlMapHeaderActionRequestHeadersToAdd{empty: true}

func (r *UrlMapHeaderActionRequestHeadersToAdd) Empty() bool {
	return r.empty
}

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

type jsonUrlMapHeaderActionResponseHeadersToAdd UrlMapHeaderActionResponseHeadersToAdd

func (r *UrlMapHeaderActionResponseHeadersToAdd) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapHeaderActionResponseHeadersToAdd
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapHeaderActionResponseHeadersToAdd
	} else {

		r.HeaderName = res.HeaderName

		r.HeaderValue = res.HeaderValue

		r.Replace = res.Replace

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapHeaderActionResponseHeadersToAdd is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapHeaderActionResponseHeadersToAdd *UrlMapHeaderActionResponseHeadersToAdd = &UrlMapHeaderActionResponseHeadersToAdd{empty: true}

func (r *UrlMapHeaderActionResponseHeadersToAdd) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionUrlRewrite UrlMapDefaultRouteActionUrlRewrite

func (r *UrlMapDefaultRouteActionUrlRewrite) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionUrlRewrite
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionUrlRewrite
	} else {

		r.PathPrefixRewrite = res.PathPrefixRewrite

		r.HostRewrite = res.HostRewrite

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionUrlRewrite is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionUrlRewrite *UrlMapDefaultRouteActionUrlRewrite = &UrlMapDefaultRouteActionUrlRewrite{empty: true}

func (r *UrlMapDefaultRouteActionUrlRewrite) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionTimeout UrlMapDefaultRouteActionTimeout

func (r *UrlMapDefaultRouteActionTimeout) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionTimeout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionTimeout
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionTimeout *UrlMapDefaultRouteActionTimeout = &UrlMapDefaultRouteActionTimeout{empty: true}

func (r *UrlMapDefaultRouteActionTimeout) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionRetryPolicy UrlMapDefaultRouteActionRetryPolicy

func (r *UrlMapDefaultRouteActionRetryPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionRetryPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionRetryPolicy
	} else {

		r.RetryCondition = res.RetryCondition

		r.NumRetries = res.NumRetries

		r.PerTryTimeout = res.PerTryTimeout

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionRetryPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionRetryPolicy *UrlMapDefaultRouteActionRetryPolicy = &UrlMapDefaultRouteActionRetryPolicy{empty: true}

func (r *UrlMapDefaultRouteActionRetryPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionRetryPolicyPerTryTimeout UrlMapDefaultRouteActionRetryPolicyPerTryTimeout

func (r *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionRetryPolicyPerTryTimeout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionRetryPolicyPerTryTimeout
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionRetryPolicyPerTryTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionRetryPolicyPerTryTimeout *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout = &UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{empty: true}

func (r *UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionRequestMirrorPolicy UrlMapDefaultRouteActionRequestMirrorPolicy

func (r *UrlMapDefaultRouteActionRequestMirrorPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionRequestMirrorPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionRequestMirrorPolicy
	} else {

		r.BackendService = res.BackendService

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionRequestMirrorPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionRequestMirrorPolicy *UrlMapDefaultRouteActionRequestMirrorPolicy = &UrlMapDefaultRouteActionRequestMirrorPolicy{empty: true}

func (r *UrlMapDefaultRouteActionRequestMirrorPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionCorsPolicy UrlMapDefaultRouteActionCorsPolicy

func (r *UrlMapDefaultRouteActionCorsPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionCorsPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionCorsPolicy
	} else {

		r.AllowOrigin = res.AllowOrigin

		r.AllowOriginRegex = res.AllowOriginRegex

		r.AllowMethod = res.AllowMethod

		r.AllowHeader = res.AllowHeader

		r.ExposeHeader = res.ExposeHeader

		r.MaxAge = res.MaxAge

		r.AllowCredentials = res.AllowCredentials

		r.Disabled = res.Disabled

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionCorsPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionCorsPolicy *UrlMapDefaultRouteActionCorsPolicy = &UrlMapDefaultRouteActionCorsPolicy{empty: true}

func (r *UrlMapDefaultRouteActionCorsPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionFaultInjectionPolicy UrlMapDefaultRouteActionFaultInjectionPolicy

func (r *UrlMapDefaultRouteActionFaultInjectionPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionFaultInjectionPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionFaultInjectionPolicy
	} else {

		r.Delay = res.Delay

		r.Abort = res.Abort

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionFaultInjectionPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionFaultInjectionPolicy *UrlMapDefaultRouteActionFaultInjectionPolicy = &UrlMapDefaultRouteActionFaultInjectionPolicy{empty: true}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionFaultInjectionPolicyDelay UrlMapDefaultRouteActionFaultInjectionPolicyDelay

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelay) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionFaultInjectionPolicyDelay
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelay
	} else {

		r.FixedDelay = res.FixedDelay

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionFaultInjectionPolicyDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelay *UrlMapDefaultRouteActionFaultInjectionPolicyDelay = &UrlMapDefaultRouteActionFaultInjectionPolicyDelay{empty: true}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelay) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay = &UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{empty: true}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultRouteActionFaultInjectionPolicyAbort UrlMapDefaultRouteActionFaultInjectionPolicyAbort

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyAbort) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultRouteActionFaultInjectionPolicyAbort
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultRouteActionFaultInjectionPolicyAbort
	} else {

		r.HttpStatus = res.HttpStatus

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultRouteActionFaultInjectionPolicyAbort is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultRouteActionFaultInjectionPolicyAbort *UrlMapDefaultRouteActionFaultInjectionPolicyAbort = &UrlMapDefaultRouteActionFaultInjectionPolicyAbort{empty: true}

func (r *UrlMapDefaultRouteActionFaultInjectionPolicyAbort) Empty() bool {
	return r.empty
}

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

type jsonUrlMapDefaultUrlRedirect UrlMapDefaultUrlRedirect

func (r *UrlMapDefaultUrlRedirect) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapDefaultUrlRedirect
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapDefaultUrlRedirect
	} else {

		r.HostRedirect = res.HostRedirect

		r.PathRedirect = res.PathRedirect

		r.PrefixRedirect = res.PrefixRedirect

		r.RedirectResponseCode = res.RedirectResponseCode

		r.HttpsRedirect = res.HttpsRedirect

		r.StripQuery = res.StripQuery

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapDefaultUrlRedirect is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapDefaultUrlRedirect *UrlMapDefaultUrlRedirect = &UrlMapDefaultUrlRedirect{empty: true}

func (r *UrlMapDefaultUrlRedirect) Empty() bool {
	return r.empty
}

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

type jsonUrlMapHostRule UrlMapHostRule

func (r *UrlMapHostRule) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapHostRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapHostRule
	} else {

		r.Description = res.Description

		r.Host = res.Host

		r.PathMatcher = res.PathMatcher

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapHostRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapHostRule *UrlMapHostRule = &UrlMapHostRule{empty: true}

func (r *UrlMapHostRule) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcher UrlMapPathMatcher

func (r *UrlMapPathMatcher) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcher
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcher
	} else {

		r.Name = res.Name

		r.Description = res.Description

		r.DefaultService = res.DefaultService

		r.DefaultRouteAction = res.DefaultRouteAction

		r.DefaultUrlRedirect = res.DefaultUrlRedirect

		r.PathRule = res.PathRule

		r.RouteRule = res.RouteRule

		r.HeaderAction = res.HeaderAction

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcher is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcher *UrlMapPathMatcher = &UrlMapPathMatcher{empty: true}

func (r *UrlMapPathMatcher) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherDefaultUrlRedirect UrlMapPathMatcherDefaultUrlRedirect

func (r *UrlMapPathMatcherDefaultUrlRedirect) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherDefaultUrlRedirect
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherDefaultUrlRedirect
	} else {

		r.HostRedirect = res.HostRedirect

		r.PathRedirect = res.PathRedirect

		r.PrefixRedirect = res.PrefixRedirect

		r.RedirectResponseCode = res.RedirectResponseCode

		r.HttpsRedirect = res.HttpsRedirect

		r.StripQuery = res.StripQuery

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherDefaultUrlRedirect is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherDefaultUrlRedirect *UrlMapPathMatcherDefaultUrlRedirect = &UrlMapPathMatcherDefaultUrlRedirect{empty: true}

func (r *UrlMapPathMatcherDefaultUrlRedirect) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRule UrlMapPathMatcherPathRule

func (r *UrlMapPathMatcherPathRule) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRule
	} else {

		r.BackendService = res.BackendService

		r.RouteAction = res.RouteAction

		r.UrlRedirect = res.UrlRedirect

		r.Path = res.Path

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRule *UrlMapPathMatcherPathRule = &UrlMapPathMatcherPathRule{empty: true}

func (r *UrlMapPathMatcherPathRule) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteAction UrlMapPathMatcherPathRuleRouteAction

func (r *UrlMapPathMatcherPathRuleRouteAction) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteAction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteAction
	} else {

		r.WeightedBackendService = res.WeightedBackendService

		r.UrlRewrite = res.UrlRewrite

		r.Timeout = res.Timeout

		r.RetryPolicy = res.RetryPolicy

		r.RequestMirrorPolicy = res.RequestMirrorPolicy

		r.CorsPolicy = res.CorsPolicy

		r.FaultInjectionPolicy = res.FaultInjectionPolicy

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteAction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteAction *UrlMapPathMatcherPathRuleRouteAction = &UrlMapPathMatcherPathRuleRouteAction{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteAction) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionWeightedBackendService UrlMapPathMatcherPathRuleRouteActionWeightedBackendService

func (r *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionWeightedBackendService
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionWeightedBackendService
	} else {

		r.BackendService = res.BackendService

		r.Weight = res.Weight

		r.HeaderAction = res.HeaderAction

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionWeightedBackendService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionWeightedBackendService *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService = &UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionUrlRewrite UrlMapPathMatcherPathRuleRouteActionUrlRewrite

func (r *UrlMapPathMatcherPathRuleRouteActionUrlRewrite) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionUrlRewrite
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionUrlRewrite
	} else {

		r.PathPrefixRewrite = res.PathPrefixRewrite

		r.HostRewrite = res.HostRewrite

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionUrlRewrite is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionUrlRewrite *UrlMapPathMatcherPathRuleRouteActionUrlRewrite = &UrlMapPathMatcherPathRuleRouteActionUrlRewrite{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionUrlRewrite) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionTimeout UrlMapPathMatcherPathRuleRouteActionTimeout

func (r *UrlMapPathMatcherPathRuleRouteActionTimeout) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionTimeout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionTimeout
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionTimeout *UrlMapPathMatcherPathRuleRouteActionTimeout = &UrlMapPathMatcherPathRuleRouteActionTimeout{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionTimeout) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionRetryPolicy UrlMapPathMatcherPathRuleRouteActionRetryPolicy

func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionRetryPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicy
	} else {

		r.RetryCondition = res.RetryCondition

		r.NumRetries = res.NumRetries

		r.PerTryTimeout = res.PerTryTimeout

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionRetryPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicy *UrlMapPathMatcherPathRuleRouteActionRetryPolicy = &UrlMapPathMatcherPathRuleRouteActionRetryPolicy{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout

func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout = &UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy

func (r *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy
	} else {

		r.BackendService = res.BackendService

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy = &UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionCorsPolicy UrlMapPathMatcherPathRuleRouteActionCorsPolicy

func (r *UrlMapPathMatcherPathRuleRouteActionCorsPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionCorsPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionCorsPolicy
	} else {

		r.AllowOrigin = res.AllowOrigin

		r.AllowOriginRegex = res.AllowOriginRegex

		r.AllowMethod = res.AllowMethod

		r.AllowHeader = res.AllowHeader

		r.ExposeHeader = res.ExposeHeader

		r.MaxAge = res.MaxAge

		r.AllowCredentials = res.AllowCredentials

		r.Disabled = res.Disabled

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionCorsPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionCorsPolicy *UrlMapPathMatcherPathRuleRouteActionCorsPolicy = &UrlMapPathMatcherPathRuleRouteActionCorsPolicy{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionCorsPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy
	} else {

		r.Delay = res.Delay

		r.Abort = res.Abort

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy = &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay
	} else {

		r.FixedDelay = res.FixedDelay

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay = &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay = &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort
	} else {

		r.HttpStatus = res.HttpStatus

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort = &UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{empty: true}

func (r *UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherPathRuleUrlRedirect UrlMapPathMatcherPathRuleUrlRedirect

func (r *UrlMapPathMatcherPathRuleUrlRedirect) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherPathRuleUrlRedirect
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherPathRuleUrlRedirect
	} else {

		r.HostRedirect = res.HostRedirect

		r.PathRedirect = res.PathRedirect

		r.PrefixRedirect = res.PrefixRedirect

		r.RedirectResponseCode = res.RedirectResponseCode

		r.HttpsRedirect = res.HttpsRedirect

		r.StripQuery = res.StripQuery

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherPathRuleUrlRedirect is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherPathRuleUrlRedirect *UrlMapPathMatcherPathRuleUrlRedirect = &UrlMapPathMatcherPathRuleUrlRedirect{empty: true}

func (r *UrlMapPathMatcherPathRuleUrlRedirect) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRule UrlMapPathMatcherRouteRule

func (r *UrlMapPathMatcherRouteRule) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRule
	} else {

		r.Priority = res.Priority

		r.Description = res.Description

		r.MatchRule = res.MatchRule

		r.BackendService = res.BackendService

		r.RouteAction = res.RouteAction

		r.UrlRedirect = res.UrlRedirect

		r.HeaderAction = res.HeaderAction

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRule *UrlMapPathMatcherRouteRule = &UrlMapPathMatcherRouteRule{empty: true}

func (r *UrlMapPathMatcherRouteRule) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleMatchRule UrlMapPathMatcherRouteRuleMatchRule

func (r *UrlMapPathMatcherRouteRuleMatchRule) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleMatchRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleMatchRule
	} else {

		r.PrefixMatch = res.PrefixMatch

		r.FullPathMatch = res.FullPathMatch

		r.RegexMatch = res.RegexMatch

		r.IgnoreCase = res.IgnoreCase

		r.HeaderMatch = res.HeaderMatch

		r.QueryParameterMatch = res.QueryParameterMatch

		r.MetadataFilter = res.MetadataFilter

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRule *UrlMapPathMatcherRouteRuleMatchRule = &UrlMapPathMatcherRouteRuleMatchRule{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRule) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch

func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch
	} else {

		r.HeaderName = res.HeaderName

		r.ExactMatch = res.ExactMatch

		r.RegexMatch = res.RegexMatch

		r.RangeMatch = res.RangeMatch

		r.PresentMatch = res.PresentMatch

		r.PrefixMatch = res.PrefixMatch

		r.SuffixMatch = res.SuffixMatch

		r.InvertMatch = res.InvertMatch

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch = &UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch

func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch
	} else {

		r.RangeStart = res.RangeStart

		r.RangeEnd = res.RangeEnd

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch = &UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch

func (r *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch
	} else {

		r.Name = res.Name

		r.PresentMatch = res.PresentMatch

		r.ExactMatch = res.ExactMatch

		r.RegexMatch = res.RegexMatch

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch = &UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter

func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter
	} else {

		r.FilterMatchCriteria = res.FilterMatchCriteria

		r.FilterLabel = res.FilterLabel

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter = &UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel

func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel
	} else {

		r.Name = res.Name

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel = &UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{empty: true}

func (r *UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteAction UrlMapPathMatcherRouteRuleRouteAction

func (r *UrlMapPathMatcherRouteRuleRouteAction) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteAction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteAction
	} else {

		r.WeightedBackendService = res.WeightedBackendService

		r.UrlRewrite = res.UrlRewrite

		r.Timeout = res.Timeout

		r.RetryPolicy = res.RetryPolicy

		r.RequestMirrorPolicy = res.RequestMirrorPolicy

		r.CorsPolicy = res.CorsPolicy

		r.FaultInjectionPolicy = res.FaultInjectionPolicy

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteAction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteAction *UrlMapPathMatcherRouteRuleRouteAction = &UrlMapPathMatcherRouteRuleRouteAction{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteAction) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService

func (r *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService
	} else {

		r.BackendService = res.BackendService

		r.Weight = res.Weight

		r.HeaderAction = res.HeaderAction

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService = &UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionUrlRewrite UrlMapPathMatcherRouteRuleRouteActionUrlRewrite

func (r *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionUrlRewrite
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionUrlRewrite
	} else {

		r.PathPrefixRewrite = res.PathPrefixRewrite

		r.HostRewrite = res.HostRewrite

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionUrlRewrite is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionUrlRewrite *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite = &UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionTimeout UrlMapPathMatcherRouteRuleRouteActionTimeout

func (r *UrlMapPathMatcherRouteRuleRouteActionTimeout) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionTimeout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionTimeout
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionTimeout *UrlMapPathMatcherRouteRuleRouteActionTimeout = &UrlMapPathMatcherRouteRuleRouteActionTimeout{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionTimeout) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionRetryPolicy UrlMapPathMatcherRouteRuleRouteActionRetryPolicy

func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionRetryPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicy
	} else {

		r.RetryCondition = res.RetryCondition

		r.NumRetries = res.NumRetries

		r.PerTryTimeout = res.PerTryTimeout

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionRetryPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicy *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy = &UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout

func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout = &UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy

func (r *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy
	} else {

		r.BackendService = res.BackendService

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy = &UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionCorsPolicy UrlMapPathMatcherRouteRuleRouteActionCorsPolicy

func (r *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionCorsPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionCorsPolicy
	} else {

		r.AllowOrigin = res.AllowOrigin

		r.AllowOriginRegex = res.AllowOriginRegex

		r.AllowMethod = res.AllowMethod

		r.AllowHeader = res.AllowHeader

		r.ExposeHeader = res.ExposeHeader

		r.MaxAge = res.MaxAge

		r.AllowCredentials = res.AllowCredentials

		r.Disabled = res.Disabled

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionCorsPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionCorsPolicy *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy = &UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy
	} else {

		r.Delay = res.Delay

		r.Abort = res.Abort

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy = &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay
	} else {

		r.FixedDelay = res.FixedDelay

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay = &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay = &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort
	} else {

		r.HttpStatus = res.HttpStatus

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort = &UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{empty: true}

func (r *UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) Empty() bool {
	return r.empty
}

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

type jsonUrlMapPathMatcherRouteRuleUrlRedirect UrlMapPathMatcherRouteRuleUrlRedirect

func (r *UrlMapPathMatcherRouteRuleUrlRedirect) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapPathMatcherRouteRuleUrlRedirect
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapPathMatcherRouteRuleUrlRedirect
	} else {

		r.HostRedirect = res.HostRedirect

		r.PathRedirect = res.PathRedirect

		r.PrefixRedirect = res.PrefixRedirect

		r.RedirectResponseCode = res.RedirectResponseCode

		r.HttpsRedirect = res.HttpsRedirect

		r.StripQuery = res.StripQuery

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapPathMatcherRouteRuleUrlRedirect is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapPathMatcherRouteRuleUrlRedirect *UrlMapPathMatcherRouteRuleUrlRedirect = &UrlMapPathMatcherRouteRuleUrlRedirect{empty: true}

func (r *UrlMapPathMatcherRouteRuleUrlRedirect) Empty() bool {
	return r.empty
}

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

type jsonUrlMapTest UrlMapTest

func (r *UrlMapTest) UnmarshalJSON(data []byte) error {
	var res jsonUrlMapTest
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyUrlMapTest
	} else {

		r.Description = res.Description

		r.Host = res.Host

		r.Path = res.Path

		r.ExpectedBackendService = res.ExpectedBackendService

	}
	return nil
}

// This object is used to assert a desired state where this UrlMapTest is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUrlMapTest *UrlMapTest = &UrlMapTest{empty: true}

func (r *UrlMapTest) Empty() bool {
	return r.empty
}

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
		Version: "beta",
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
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
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
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListUrlMapWithMaxResults(ctx, project, UrlMapMaxPage)

}

func (c *Client) ListUrlMapWithMaxResults(ctx context.Context, project string, pageSize int32) (*UrlMapList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
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

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *UrlMap) URLNormalized() *UrlMap {
	normalized := dcl.Copy(*r).(UrlMap)
	normalized.DefaultService = dcl.SelfLinkToName(r.DefaultService)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetUrlMap(ctx context.Context, r *UrlMap) (*UrlMap, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
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
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
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
	var resultNewState *UrlMap
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyUrlMapHelper(c, ctx, rawDesired, opts...)
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

func applyUrlMapHelper(c *Client, ctx context.Context, rawDesired *UrlMap, opts ...dcl.ApplyOption) (*UrlMap, error) {
	c.Config.Logger.Info("Beginning ApplyUrlMap...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.urlMapDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToUrlMapDiffs(c.Config, fieldDiffs, opts)
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
	rawNew, err := c.GetUrlMap(ctx, desired.URLNormalized())
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
