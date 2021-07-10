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

type HealthCheck struct {
	CheckIntervalSec   *int64                       `json:"checkIntervalSec"`
	Description        *string                      `json:"description"`
	HealthyThreshold   *int64                       `json:"healthyThreshold"`
	Http2HealthCheck   *HealthCheckHttp2HealthCheck `json:"http2HealthCheck"`
	HttpHealthCheck    *HealthCheckHttpHealthCheck  `json:"httpHealthCheck"`
	HttpsHealthCheck   *HealthCheckHttpsHealthCheck `json:"httpsHealthCheck"`
	Name               *string                      `json:"name"`
	SslHealthCheck     *HealthCheckSslHealthCheck   `json:"sslHealthCheck"`
	TcpHealthCheck     *HealthCheckTcpHealthCheck   `json:"tcpHealthCheck"`
	Type               *HealthCheckTypeEnum         `json:"type"`
	UnhealthyThreshold *int64                       `json:"unhealthyThreshold"`
	TimeoutSec         *int64                       `json:"timeoutSec"`
	Region             *string                      `json:"region"`
	Project            *string                      `json:"project"`
	SelfLink           *string                      `json:"selfLink"`
	Location           *string                      `json:"location"`
}

func (r *HealthCheck) String() string {
	return dcl.SprintResource(r)
}

// The enum HealthCheckHttp2HealthCheckPortSpecificationEnum.
type HealthCheckHttp2HealthCheckPortSpecificationEnum string

// HealthCheckHttp2HealthCheckPortSpecificationEnumRef returns a *HealthCheckHttp2HealthCheckPortSpecificationEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckHttp2HealthCheckPortSpecificationEnumRef(s string) *HealthCheckHttp2HealthCheckPortSpecificationEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckHttp2HealthCheckPortSpecificationEnum(s)
	return &v
}

func (v HealthCheckHttp2HealthCheckPortSpecificationEnum) Validate() error {
	for _, s := range []string{"USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckHttp2HealthCheckPortSpecificationEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckHttp2HealthCheckProxyHeaderEnum.
type HealthCheckHttp2HealthCheckProxyHeaderEnum string

// HealthCheckHttp2HealthCheckProxyHeaderEnumRef returns a *HealthCheckHttp2HealthCheckProxyHeaderEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckHttp2HealthCheckProxyHeaderEnumRef(s string) *HealthCheckHttp2HealthCheckProxyHeaderEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckHttp2HealthCheckProxyHeaderEnum(s)
	return &v
}

func (v HealthCheckHttp2HealthCheckProxyHeaderEnum) Validate() error {
	for _, s := range []string{"NONE", "PROXY_V1"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckHttp2HealthCheckProxyHeaderEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckHttpHealthCheckPortSpecificationEnum.
type HealthCheckHttpHealthCheckPortSpecificationEnum string

// HealthCheckHttpHealthCheckPortSpecificationEnumRef returns a *HealthCheckHttpHealthCheckPortSpecificationEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckHttpHealthCheckPortSpecificationEnumRef(s string) *HealthCheckHttpHealthCheckPortSpecificationEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckHttpHealthCheckPortSpecificationEnum(s)
	return &v
}

func (v HealthCheckHttpHealthCheckPortSpecificationEnum) Validate() error {
	for _, s := range []string{"USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckHttpHealthCheckPortSpecificationEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckHttpHealthCheckProxyHeaderEnum.
type HealthCheckHttpHealthCheckProxyHeaderEnum string

// HealthCheckHttpHealthCheckProxyHeaderEnumRef returns a *HealthCheckHttpHealthCheckProxyHeaderEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckHttpHealthCheckProxyHeaderEnumRef(s string) *HealthCheckHttpHealthCheckProxyHeaderEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckHttpHealthCheckProxyHeaderEnum(s)
	return &v
}

func (v HealthCheckHttpHealthCheckProxyHeaderEnum) Validate() error {
	for _, s := range []string{"NONE", "PROXY_V1"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckHttpHealthCheckProxyHeaderEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckHttpsHealthCheckPortSpecificationEnum.
type HealthCheckHttpsHealthCheckPortSpecificationEnum string

// HealthCheckHttpsHealthCheckPortSpecificationEnumRef returns a *HealthCheckHttpsHealthCheckPortSpecificationEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckHttpsHealthCheckPortSpecificationEnumRef(s string) *HealthCheckHttpsHealthCheckPortSpecificationEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckHttpsHealthCheckPortSpecificationEnum(s)
	return &v
}

func (v HealthCheckHttpsHealthCheckPortSpecificationEnum) Validate() error {
	for _, s := range []string{"USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckHttpsHealthCheckPortSpecificationEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckHttpsHealthCheckProxyHeaderEnum.
type HealthCheckHttpsHealthCheckProxyHeaderEnum string

// HealthCheckHttpsHealthCheckProxyHeaderEnumRef returns a *HealthCheckHttpsHealthCheckProxyHeaderEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckHttpsHealthCheckProxyHeaderEnumRef(s string) *HealthCheckHttpsHealthCheckProxyHeaderEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckHttpsHealthCheckProxyHeaderEnum(s)
	return &v
}

func (v HealthCheckHttpsHealthCheckProxyHeaderEnum) Validate() error {
	for _, s := range []string{"NONE", "PROXY_V1"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckHttpsHealthCheckProxyHeaderEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckSslHealthCheckPortSpecificationEnum.
type HealthCheckSslHealthCheckPortSpecificationEnum string

// HealthCheckSslHealthCheckPortSpecificationEnumRef returns a *HealthCheckSslHealthCheckPortSpecificationEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckSslHealthCheckPortSpecificationEnumRef(s string) *HealthCheckSslHealthCheckPortSpecificationEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckSslHealthCheckPortSpecificationEnum(s)
	return &v
}

func (v HealthCheckSslHealthCheckPortSpecificationEnum) Validate() error {
	for _, s := range []string{"USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckSslHealthCheckPortSpecificationEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckSslHealthCheckProxyHeaderEnum.
type HealthCheckSslHealthCheckProxyHeaderEnum string

// HealthCheckSslHealthCheckProxyHeaderEnumRef returns a *HealthCheckSslHealthCheckProxyHeaderEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckSslHealthCheckProxyHeaderEnumRef(s string) *HealthCheckSslHealthCheckProxyHeaderEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckSslHealthCheckProxyHeaderEnum(s)
	return &v
}

func (v HealthCheckSslHealthCheckProxyHeaderEnum) Validate() error {
	for _, s := range []string{"NONE", "PROXY_V1"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckSslHealthCheckProxyHeaderEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckTcpHealthCheckPortSpecificationEnum.
type HealthCheckTcpHealthCheckPortSpecificationEnum string

// HealthCheckTcpHealthCheckPortSpecificationEnumRef returns a *HealthCheckTcpHealthCheckPortSpecificationEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckTcpHealthCheckPortSpecificationEnumRef(s string) *HealthCheckTcpHealthCheckPortSpecificationEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckTcpHealthCheckPortSpecificationEnum(s)
	return &v
}

func (v HealthCheckTcpHealthCheckPortSpecificationEnum) Validate() error {
	for _, s := range []string{"USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckTcpHealthCheckPortSpecificationEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckTcpHealthCheckProxyHeaderEnum.
type HealthCheckTcpHealthCheckProxyHeaderEnum string

// HealthCheckTcpHealthCheckProxyHeaderEnumRef returns a *HealthCheckTcpHealthCheckProxyHeaderEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckTcpHealthCheckProxyHeaderEnumRef(s string) *HealthCheckTcpHealthCheckProxyHeaderEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckTcpHealthCheckProxyHeaderEnum(s)
	return &v
}

func (v HealthCheckTcpHealthCheckProxyHeaderEnum) Validate() error {
	for _, s := range []string{"NONE", "PROXY_V1"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckTcpHealthCheckProxyHeaderEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum HealthCheckTypeEnum.
type HealthCheckTypeEnum string

// HealthCheckTypeEnumRef returns a *HealthCheckTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func HealthCheckTypeEnumRef(s string) *HealthCheckTypeEnum {
	if s == "" {
		return nil
	}

	v := HealthCheckTypeEnum(s)
	return &v
}

func (v HealthCheckTypeEnum) Validate() error {
	for _, s := range []string{"HTTP", "HTTPS", "HTTP2", "TCP", "SSL", "INVALID"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HealthCheckTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type HealthCheckHttp2HealthCheck struct {
	empty             bool                                              `json:"-"`
	Port              *int64                                            `json:"port"`
	PortName          *string                                           `json:"portName"`
	PortSpecification *HealthCheckHttp2HealthCheckPortSpecificationEnum `json:"portSpecification"`
	Host              *string                                           `json:"host"`
	RequestPath       *string                                           `json:"requestPath"`
	ProxyHeader       *HealthCheckHttp2HealthCheckProxyHeaderEnum       `json:"proxyHeader"`
	Response          *string                                           `json:"response"`
}

type jsonHealthCheckHttp2HealthCheck HealthCheckHttp2HealthCheck

func (r *HealthCheckHttp2HealthCheck) UnmarshalJSON(data []byte) error {
	var res jsonHealthCheckHttp2HealthCheck
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHealthCheckHttp2HealthCheck
	} else {

		r.Port = res.Port

		r.PortName = res.PortName

		r.PortSpecification = res.PortSpecification

		r.Host = res.Host

		r.RequestPath = res.RequestPath

		r.ProxyHeader = res.ProxyHeader

		r.Response = res.Response

	}
	return nil
}

// This object is used to assert a desired state where this HealthCheckHttp2HealthCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyHealthCheckHttp2HealthCheck *HealthCheckHttp2HealthCheck = &HealthCheckHttp2HealthCheck{empty: true}

func (r *HealthCheckHttp2HealthCheck) Empty() bool {
	return r.empty
}

func (r *HealthCheckHttp2HealthCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *HealthCheckHttp2HealthCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HealthCheckHttpHealthCheck struct {
	empty             bool                                             `json:"-"`
	Port              *int64                                           `json:"port"`
	PortName          *string                                          `json:"portName"`
	PortSpecification *HealthCheckHttpHealthCheckPortSpecificationEnum `json:"portSpecification"`
	Host              *string                                          `json:"host"`
	RequestPath       *string                                          `json:"requestPath"`
	ProxyHeader       *HealthCheckHttpHealthCheckProxyHeaderEnum       `json:"proxyHeader"`
	Response          *string                                          `json:"response"`
}

type jsonHealthCheckHttpHealthCheck HealthCheckHttpHealthCheck

func (r *HealthCheckHttpHealthCheck) UnmarshalJSON(data []byte) error {
	var res jsonHealthCheckHttpHealthCheck
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHealthCheckHttpHealthCheck
	} else {

		r.Port = res.Port

		r.PortName = res.PortName

		r.PortSpecification = res.PortSpecification

		r.Host = res.Host

		r.RequestPath = res.RequestPath

		r.ProxyHeader = res.ProxyHeader

		r.Response = res.Response

	}
	return nil
}

// This object is used to assert a desired state where this HealthCheckHttpHealthCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyHealthCheckHttpHealthCheck *HealthCheckHttpHealthCheck = &HealthCheckHttpHealthCheck{empty: true}

func (r *HealthCheckHttpHealthCheck) Empty() bool {
	return r.empty
}

func (r *HealthCheckHttpHealthCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *HealthCheckHttpHealthCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HealthCheckHttpsHealthCheck struct {
	empty             bool                                              `json:"-"`
	Port              *int64                                            `json:"port"`
	PortName          *string                                           `json:"portName"`
	PortSpecification *HealthCheckHttpsHealthCheckPortSpecificationEnum `json:"portSpecification"`
	Host              *string                                           `json:"host"`
	RequestPath       *string                                           `json:"requestPath"`
	ProxyHeader       *HealthCheckHttpsHealthCheckProxyHeaderEnum       `json:"proxyHeader"`
	Response          *string                                           `json:"response"`
}

type jsonHealthCheckHttpsHealthCheck HealthCheckHttpsHealthCheck

func (r *HealthCheckHttpsHealthCheck) UnmarshalJSON(data []byte) error {
	var res jsonHealthCheckHttpsHealthCheck
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHealthCheckHttpsHealthCheck
	} else {

		r.Port = res.Port

		r.PortName = res.PortName

		r.PortSpecification = res.PortSpecification

		r.Host = res.Host

		r.RequestPath = res.RequestPath

		r.ProxyHeader = res.ProxyHeader

		r.Response = res.Response

	}
	return nil
}

// This object is used to assert a desired state where this HealthCheckHttpsHealthCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyHealthCheckHttpsHealthCheck *HealthCheckHttpsHealthCheck = &HealthCheckHttpsHealthCheck{empty: true}

func (r *HealthCheckHttpsHealthCheck) Empty() bool {
	return r.empty
}

func (r *HealthCheckHttpsHealthCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *HealthCheckHttpsHealthCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HealthCheckSslHealthCheck struct {
	empty             bool                                            `json:"-"`
	Port              *int64                                          `json:"port"`
	PortName          *string                                         `json:"portName"`
	PortSpecification *HealthCheckSslHealthCheckPortSpecificationEnum `json:"portSpecification"`
	Request           *string                                         `json:"request"`
	Response          *string                                         `json:"response"`
	ProxyHeader       *HealthCheckSslHealthCheckProxyHeaderEnum       `json:"proxyHeader"`
}

type jsonHealthCheckSslHealthCheck HealthCheckSslHealthCheck

func (r *HealthCheckSslHealthCheck) UnmarshalJSON(data []byte) error {
	var res jsonHealthCheckSslHealthCheck
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHealthCheckSslHealthCheck
	} else {

		r.Port = res.Port

		r.PortName = res.PortName

		r.PortSpecification = res.PortSpecification

		r.Request = res.Request

		r.Response = res.Response

		r.ProxyHeader = res.ProxyHeader

	}
	return nil
}

// This object is used to assert a desired state where this HealthCheckSslHealthCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyHealthCheckSslHealthCheck *HealthCheckSslHealthCheck = &HealthCheckSslHealthCheck{empty: true}

func (r *HealthCheckSslHealthCheck) Empty() bool {
	return r.empty
}

func (r *HealthCheckSslHealthCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *HealthCheckSslHealthCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HealthCheckTcpHealthCheck struct {
	empty             bool                                            `json:"-"`
	Port              *int64                                          `json:"port"`
	PortName          *string                                         `json:"portName"`
	PortSpecification *HealthCheckTcpHealthCheckPortSpecificationEnum `json:"portSpecification"`
	Request           *string                                         `json:"request"`
	Response          *string                                         `json:"response"`
	ProxyHeader       *HealthCheckTcpHealthCheckProxyHeaderEnum       `json:"proxyHeader"`
}

type jsonHealthCheckTcpHealthCheck HealthCheckTcpHealthCheck

func (r *HealthCheckTcpHealthCheck) UnmarshalJSON(data []byte) error {
	var res jsonHealthCheckTcpHealthCheck
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHealthCheckTcpHealthCheck
	} else {

		r.Port = res.Port

		r.PortName = res.PortName

		r.PortSpecification = res.PortSpecification

		r.Request = res.Request

		r.Response = res.Response

		r.ProxyHeader = res.ProxyHeader

	}
	return nil
}

// This object is used to assert a desired state where this HealthCheckTcpHealthCheck is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyHealthCheckTcpHealthCheck *HealthCheckTcpHealthCheck = &HealthCheckTcpHealthCheck{empty: true}

func (r *HealthCheckTcpHealthCheck) Empty() bool {
	return r.empty
}

func (r *HealthCheckTcpHealthCheck) String() string {
	return dcl.SprintResource(r)
}

func (r *HealthCheckTcpHealthCheck) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *HealthCheck) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "HealthCheck",
		Version: "compute",
	}
}

const HealthCheckMaxPage = -1

type HealthCheckList struct {
	Items []*HealthCheck

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *HealthCheckList) HasNext() bool {
	return l.nextToken != ""
}

func (l *HealthCheckList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listHealthCheck(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListHealthCheck(ctx context.Context, project, location string) (*HealthCheckList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListHealthCheckWithMaxResults(ctx, project, location, HealthCheckMaxPage)

}

func (c *Client) ListHealthCheckWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*HealthCheckList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listHealthCheck(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &HealthCheckList{
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
func (r *HealthCheck) URLNormalized() *HealthCheck {
	normalized := dcl.Copy(*r).(HealthCheck)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetHealthCheck(ctx context.Context, r *HealthCheck) (*HealthCheck, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getHealthCheckRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalHealthCheck(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeHealthCheckNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteHealthCheck(ctx context.Context, r *HealthCheck) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("HealthCheck resource is nil")
	}
	c.Config.Logger.Info("Deleting HealthCheck...")
	deleteOp := deleteHealthCheckOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllHealthCheck deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllHealthCheck(ctx context.Context, project, location string, filter func(*HealthCheck) bool) error {
	listObj, err := c.ListHealthCheck(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllHealthCheck(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllHealthCheck(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyHealthCheck(ctx context.Context, rawDesired *HealthCheck, opts ...dcl.ApplyOption) (*HealthCheck, error) {

	var resultNewState *HealthCheck
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyHealthCheckHelper(c, ctx, rawDesired, opts...)
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

func applyHealthCheckHelper(c *Client, ctx context.Context, rawDesired *HealthCheck, opts ...dcl.ApplyOption) (*HealthCheck, error) {
	c.Config.Logger.Info("Beginning ApplyHealthCheck...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.healthCheckDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToHealthCheckOp(opStrings, fieldDiffs, opts)
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
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []healthCheckApiOperation
	if create {
		ops = append(ops, &createHealthCheckOperation{})
	} else if recreate {
		ops = append(ops, &deleteHealthCheckOperation{})
		ops = append(ops, &createHealthCheckOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeHealthCheckDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetHealthCheck(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createHealthCheckOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapHealthCheck(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeHealthCheckNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeHealthCheckNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeHealthCheckDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffHealthCheck(c, newDesired, newState)
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
