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

type InstanceGroupManager struct {
	BaseInstanceName    *string                                   `json:"baseInstanceName"`
	CreationTimestamp   *string                                   `json:"creationTimestamp"`
	DistributionPolicy  *InstanceGroupManagerDistributionPolicy   `json:"distributionPolicy"`
	CurrentActions      *InstanceGroupManagerCurrentActions       `json:"currentActions"`
	Description         *string                                   `json:"description"`
	Versions            []InstanceGroupManagerVersions            `json:"versions"`
	Id                  *int64                                    `json:"id"`
	InstanceGroup       *string                                   `json:"instanceGroup"`
	InstanceTemplate    *string                                   `json:"instanceTemplate"`
	Name                *string                                   `json:"name"`
	NamedPorts          []InstanceGroupManagerNamedPorts          `json:"namedPorts"`
	Status              *InstanceGroupManagerStatus               `json:"status"`
	TargetPools         []string                                  `json:"targetPools"`
	AutoHealingPolicies []InstanceGroupManagerAutoHealingPolicies `json:"autoHealingPolicies"`
	UpdatePolicy        *InstanceGroupManagerUpdatePolicy         `json:"updatePolicy"`
	TargetSize          *int64                                    `json:"targetSize"`
	Zone                *string                                   `json:"zone"`
	Region              *string                                   `json:"region"`
	Project             *string                                   `json:"project"`
	Location            *string                                   `json:"location"`
}

func (r *InstanceGroupManager) String() string {
	return dcl.SprintResource(r)
}

// The enum InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.
type InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum string

// InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumRef returns a *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumRef(s string) *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(s)
	return &v
}

func (v InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) Validate() error {
	for _, s := range []string{"PROACTIVE", "NONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceGroupManagerUpdatePolicyMinimalActionEnum.
type InstanceGroupManagerUpdatePolicyMinimalActionEnum string

// InstanceGroupManagerUpdatePolicyMinimalActionEnumRef returns a *InstanceGroupManagerUpdatePolicyMinimalActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceGroupManagerUpdatePolicyMinimalActionEnumRef(s string) *InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceGroupManagerUpdatePolicyMinimalActionEnum(s)
	return &v
}

func (v InstanceGroupManagerUpdatePolicyMinimalActionEnum) Validate() error {
	for _, s := range []string{"RESTART", "REPLACE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceGroupManagerUpdatePolicyMinimalActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InstanceGroupManagerDistributionPolicy struct {
	empty bool                                          `json:"-"`
	Zones []InstanceGroupManagerDistributionPolicyZones `json:"zones"`
}

type jsonInstanceGroupManagerDistributionPolicy InstanceGroupManagerDistributionPolicy

func (r *InstanceGroupManagerDistributionPolicy) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerDistributionPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerDistributionPolicy
	} else {

		r.Zones = res.Zones

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerDistributionPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerDistributionPolicy *InstanceGroupManagerDistributionPolicy = &InstanceGroupManagerDistributionPolicy{empty: true}

func (r *InstanceGroupManagerDistributionPolicy) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerDistributionPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerDistributionPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerDistributionPolicyZones struct {
	empty bool    `json:"-"`
	Zone  *string `json:"zone"`
}

type jsonInstanceGroupManagerDistributionPolicyZones InstanceGroupManagerDistributionPolicyZones

func (r *InstanceGroupManagerDistributionPolicyZones) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerDistributionPolicyZones
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerDistributionPolicyZones
	} else {

		r.Zone = res.Zone

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerDistributionPolicyZones is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerDistributionPolicyZones *InstanceGroupManagerDistributionPolicyZones = &InstanceGroupManagerDistributionPolicyZones{empty: true}

func (r *InstanceGroupManagerDistributionPolicyZones) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerDistributionPolicyZones) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerDistributionPolicyZones) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerCurrentActions struct {
	empty                  bool   `json:"-"`
	Abandoning             *int64 `json:"abandoning"`
	Creating               *int64 `json:"creating"`
	CreatingWithoutRetries *int64 `json:"creatingWithoutRetries"`
	Deleting               *int64 `json:"deleting"`
	None                   *int64 `json:"none"`
	Recreating             *int64 `json:"recreating"`
	Refreshing             *int64 `json:"refreshing"`
	Restarting             *int64 `json:"restarting"`
}

type jsonInstanceGroupManagerCurrentActions InstanceGroupManagerCurrentActions

func (r *InstanceGroupManagerCurrentActions) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerCurrentActions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerCurrentActions
	} else {

		r.Abandoning = res.Abandoning

		r.Creating = res.Creating

		r.CreatingWithoutRetries = res.CreatingWithoutRetries

		r.Deleting = res.Deleting

		r.None = res.None

		r.Recreating = res.Recreating

		r.Refreshing = res.Refreshing

		r.Restarting = res.Restarting

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerCurrentActions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerCurrentActions *InstanceGroupManagerCurrentActions = &InstanceGroupManagerCurrentActions{empty: true}

func (r *InstanceGroupManagerCurrentActions) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerCurrentActions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerCurrentActions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerVersions struct {
	empty            bool                                    `json:"-"`
	Name             *string                                 `json:"name"`
	InstanceTemplate *string                                 `json:"instanceTemplate"`
	TargetSize       *InstanceGroupManagerVersionsTargetSize `json:"targetSize"`
}

type jsonInstanceGroupManagerVersions InstanceGroupManagerVersions

func (r *InstanceGroupManagerVersions) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerVersions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerVersions
	} else {

		r.Name = res.Name

		r.InstanceTemplate = res.InstanceTemplate

		r.TargetSize = res.TargetSize

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerVersions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerVersions *InstanceGroupManagerVersions = &InstanceGroupManagerVersions{empty: true}

func (r *InstanceGroupManagerVersions) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerVersions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerVersions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerVersionsTargetSize struct {
	empty      bool   `json:"-"`
	Fixed      *int64 `json:"fixed"`
	Percent    *int64 `json:"percent"`
	Calculated *int64 `json:"calculated"`
}

type jsonInstanceGroupManagerVersionsTargetSize InstanceGroupManagerVersionsTargetSize

func (r *InstanceGroupManagerVersionsTargetSize) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerVersionsTargetSize
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerVersionsTargetSize
	} else {

		r.Fixed = res.Fixed

		r.Percent = res.Percent

		r.Calculated = res.Calculated

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerVersionsTargetSize is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerVersionsTargetSize *InstanceGroupManagerVersionsTargetSize = &InstanceGroupManagerVersionsTargetSize{empty: true}

func (r *InstanceGroupManagerVersionsTargetSize) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerVersionsTargetSize) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerVersionsTargetSize) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerNamedPorts struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Port  *int64  `json:"port"`
}

type jsonInstanceGroupManagerNamedPorts InstanceGroupManagerNamedPorts

func (r *InstanceGroupManagerNamedPorts) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerNamedPorts
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerNamedPorts
	} else {

		r.Name = res.Name

		r.Port = res.Port

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerNamedPorts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerNamedPorts *InstanceGroupManagerNamedPorts = &InstanceGroupManagerNamedPorts{empty: true}

func (r *InstanceGroupManagerNamedPorts) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerNamedPorts) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerNamedPorts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerStatus struct {
	empty         bool                                     `json:"-"`
	IsStable      *bool                                    `json:"isStable"`
	VersionTarget *InstanceGroupManagerStatusVersionTarget `json:"versionTarget"`
	Autoscalar    *string                                  `json:"autoscalar"`
}

type jsonInstanceGroupManagerStatus InstanceGroupManagerStatus

func (r *InstanceGroupManagerStatus) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerStatus
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerStatus
	} else {

		r.IsStable = res.IsStable

		r.VersionTarget = res.VersionTarget

		r.Autoscalar = res.Autoscalar

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerStatus *InstanceGroupManagerStatus = &InstanceGroupManagerStatus{empty: true}

func (r *InstanceGroupManagerStatus) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerStatusVersionTarget struct {
	empty     bool  `json:"-"`
	IsReached *bool `json:"isReached"`
}

type jsonInstanceGroupManagerStatusVersionTarget InstanceGroupManagerStatusVersionTarget

func (r *InstanceGroupManagerStatusVersionTarget) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerStatusVersionTarget
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerStatusVersionTarget
	} else {

		r.IsReached = res.IsReached

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerStatusVersionTarget is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerStatusVersionTarget *InstanceGroupManagerStatusVersionTarget = &InstanceGroupManagerStatusVersionTarget{empty: true}

func (r *InstanceGroupManagerStatusVersionTarget) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerStatusVersionTarget) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerStatusVersionTarget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerAutoHealingPolicies struct {
	empty           bool    `json:"-"`
	HealthCheck     *string `json:"healthCheck"`
	InitialDelaySec *int64  `json:"initialDelaySec"`
}

type jsonInstanceGroupManagerAutoHealingPolicies InstanceGroupManagerAutoHealingPolicies

func (r *InstanceGroupManagerAutoHealingPolicies) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerAutoHealingPolicies
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerAutoHealingPolicies
	} else {

		r.HealthCheck = res.HealthCheck

		r.InitialDelaySec = res.InitialDelaySec

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerAutoHealingPolicies is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerAutoHealingPolicies *InstanceGroupManagerAutoHealingPolicies = &InstanceGroupManagerAutoHealingPolicies{empty: true}

func (r *InstanceGroupManagerAutoHealingPolicies) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerAutoHealingPolicies) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerAutoHealingPolicies) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerUpdatePolicy struct {
	empty                      bool                                                            `json:"-"`
	InstanceRedistributionType *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum `json:"instanceRedistributionType"`
	MinimalAction              *InstanceGroupManagerUpdatePolicyMinimalActionEnum              `json:"minimalAction"`
	MaxSurge                   *InstanceGroupManagerUpdatePolicyMaxSurge                       `json:"maxSurge"`
	MaxUnavailable             *InstanceGroupManagerUpdatePolicyMaxUnavailable                 `json:"maxUnavailable"`
}

type jsonInstanceGroupManagerUpdatePolicy InstanceGroupManagerUpdatePolicy

func (r *InstanceGroupManagerUpdatePolicy) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerUpdatePolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerUpdatePolicy
	} else {

		r.InstanceRedistributionType = res.InstanceRedistributionType

		r.MinimalAction = res.MinimalAction

		r.MaxSurge = res.MaxSurge

		r.MaxUnavailable = res.MaxUnavailable

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerUpdatePolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerUpdatePolicy *InstanceGroupManagerUpdatePolicy = &InstanceGroupManagerUpdatePolicy{empty: true}

func (r *InstanceGroupManagerUpdatePolicy) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerUpdatePolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerUpdatePolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerUpdatePolicyMaxSurge struct {
	empty      bool   `json:"-"`
	Fixed      *int64 `json:"fixed"`
	Percent    *int64 `json:"percent"`
	Calculated *int64 `json:"calculated"`
}

type jsonInstanceGroupManagerUpdatePolicyMaxSurge InstanceGroupManagerUpdatePolicyMaxSurge

func (r *InstanceGroupManagerUpdatePolicyMaxSurge) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerUpdatePolicyMaxSurge
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerUpdatePolicyMaxSurge
	} else {

		r.Fixed = res.Fixed

		r.Percent = res.Percent

		r.Calculated = res.Calculated

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerUpdatePolicyMaxSurge is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerUpdatePolicyMaxSurge *InstanceGroupManagerUpdatePolicyMaxSurge = &InstanceGroupManagerUpdatePolicyMaxSurge{empty: true}

func (r *InstanceGroupManagerUpdatePolicyMaxSurge) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerUpdatePolicyMaxSurge) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerUpdatePolicyMaxSurge) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerUpdatePolicyMaxUnavailable struct {
	empty      bool   `json:"-"`
	Fixed      *int64 `json:"fixed"`
	Percent    *int64 `json:"percent"`
	Calculated *int64 `json:"calculated"`
}

type jsonInstanceGroupManagerUpdatePolicyMaxUnavailable InstanceGroupManagerUpdatePolicyMaxUnavailable

func (r *InstanceGroupManagerUpdatePolicyMaxUnavailable) UnmarshalJSON(data []byte) error {
	var res jsonInstanceGroupManagerUpdatePolicyMaxUnavailable
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceGroupManagerUpdatePolicyMaxUnavailable
	} else {

		r.Fixed = res.Fixed

		r.Percent = res.Percent

		r.Calculated = res.Calculated

	}
	return nil
}

// This object is used to assert a desired state where this InstanceGroupManagerUpdatePolicyMaxUnavailable is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerUpdatePolicyMaxUnavailable *InstanceGroupManagerUpdatePolicyMaxUnavailable = &InstanceGroupManagerUpdatePolicyMaxUnavailable{empty: true}

func (r *InstanceGroupManagerUpdatePolicyMaxUnavailable) Empty() bool {
	return r.empty
}

func (r *InstanceGroupManagerUpdatePolicyMaxUnavailable) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerUpdatePolicyMaxUnavailable) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *InstanceGroupManager) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "InstanceGroupManager",
		Version: "compute",
	}
}

const InstanceGroupManagerMaxPage = -1

type InstanceGroupManagerList struct {
	Items []*InstanceGroupManager

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *InstanceGroupManagerList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InstanceGroupManagerList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInstanceGroupManager(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInstanceGroupManager(ctx context.Context, project, location string) (*InstanceGroupManagerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListInstanceGroupManagerWithMaxResults(ctx, project, location, InstanceGroupManagerMaxPage)

}

func (c *Client) ListInstanceGroupManagerWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*InstanceGroupManagerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listInstanceGroupManager(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InstanceGroupManagerList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetInstanceGroupManager(ctx context.Context, r *InstanceGroupManager) (*InstanceGroupManager, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getInstanceGroupManagerRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalInstanceGroupManager(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeInstanceGroupManagerNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteInstanceGroupManager(ctx context.Context, r *InstanceGroupManager) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("InstanceGroupManager resource is nil")
	}
	c.Config.Logger.Info("Deleting InstanceGroupManager...")
	deleteOp := deleteInstanceGroupManagerOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInstanceGroupManager deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInstanceGroupManager(ctx context.Context, project, location string, filter func(*InstanceGroupManager) bool) error {
	listObj, err := c.ListInstanceGroupManager(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllInstanceGroupManager(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllInstanceGroupManager(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyInstanceGroupManager(ctx context.Context, rawDesired *InstanceGroupManager, opts ...dcl.ApplyOption) (*InstanceGroupManager, error) {

	var resultNewState *InstanceGroupManager
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyInstanceGroupManagerHelper(c, ctx, rawDesired, opts...)
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

func applyInstanceGroupManagerHelper(c *Client, ctx context.Context, rawDesired *InstanceGroupManager, opts ...dcl.ApplyOption) (*InstanceGroupManager, error) {
	c.Config.Logger.Info("Beginning ApplyInstanceGroupManager...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.instanceGroupManagerDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToInstanceGroupManagerOp(opStrings, fieldDiffs, opts)
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
	var ops []instanceGroupManagerApiOperation
	if create {
		ops = append(ops, &createInstanceGroupManagerOperation{})
	} else if recreate {

		ops = append(ops, &deleteInstanceGroupManagerOperation{})

		ops = append(ops, &createInstanceGroupManagerOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeInstanceGroupManagerDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetInstanceGroupManager(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createInstanceGroupManagerOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapInstanceGroupManager(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeInstanceGroupManagerNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeInstanceGroupManagerNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeInstanceGroupManagerDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffInstanceGroupManager(c, newDesired, newState)
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
