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

type NodePool struct {
	Name              *string                    `json:"name"`
	Config            *NodePoolConfig            `json:"config"`
	NodeCount         *int64                     `json:"nodeCount"`
	Version           *string                    `json:"version"`
	Status            *string                    `json:"status"`
	StatusMessage     *string                    `json:"statusMessage"`
	Locations         []string                   `json:"locations"`
	Autoscaling       *NodePoolAutoscaling       `json:"autoscaling"`
	Management        *NodePoolManagement        `json:"management"`
	MaxPodsConstraint *NodePoolMaxPodsConstraint `json:"maxPodsConstraint"`
	Conditions        []NodePoolConditions       `json:"conditions"`
	PodIPv4CidrSize   *int64                     `json:"podIPv4CidrSize"`
	UpgradeSettings   *NodePoolUpgradeSettings   `json:"upgradeSettings"`
	Cluster           *string                    `json:"cluster"`
	Project           *string                    `json:"project"`
	Location          *string                    `json:"location"`
}

func (r *NodePool) String() string {
	return dcl.SprintResource(r)
}

// The enum NodePoolConfigSandboxConfigTypeEnum.
type NodePoolConfigSandboxConfigTypeEnum string

// NodePoolConfigSandboxConfigTypeEnumRef returns a *NodePoolConfigSandboxConfigTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func NodePoolConfigSandboxConfigTypeEnumRef(s string) *NodePoolConfigSandboxConfigTypeEnum {
	if s == "" {
		return nil
	}

	v := NodePoolConfigSandboxConfigTypeEnum(s)
	return &v
}

func (v NodePoolConfigSandboxConfigTypeEnum) Validate() error {
	for _, s := range []string{"GVISOR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NodePoolConfigSandboxConfigTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NodePoolConfigReservationAffinityConsumeReservationTypeEnum.
type NodePoolConfigReservationAffinityConsumeReservationTypeEnum string

// NodePoolConfigReservationAffinityConsumeReservationTypeEnumRef returns a *NodePoolConfigReservationAffinityConsumeReservationTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func NodePoolConfigReservationAffinityConsumeReservationTypeEnumRef(s string) *NodePoolConfigReservationAffinityConsumeReservationTypeEnum {
	if s == "" {
		return nil
	}

	v := NodePoolConfigReservationAffinityConsumeReservationTypeEnum(s)
	return &v
}

func (v NodePoolConfigReservationAffinityConsumeReservationTypeEnum) Validate() error {
	for _, s := range []string{"NO_RESERVATION", "ANY_RESERVATION", "SPECIFIC_RESERVATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NodePoolConfigReservationAffinityConsumeReservationTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NodePoolConditionsCodeEnum.
type NodePoolConditionsCodeEnum string

// NodePoolConditionsCodeEnumRef returns a *NodePoolConditionsCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func NodePoolConditionsCodeEnumRef(s string) *NodePoolConditionsCodeEnum {
	if s == "" {
		return nil
	}

	v := NodePoolConditionsCodeEnum(s)
	return &v
}

func (v NodePoolConditionsCodeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN", "GCE_STOCKOUT", "GKE_SERVICE_ACCOUNT_DELETED", "GCE_QUOTA_EXCEEDED", "SET_BY_OPERATOR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NodePoolConditionsCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type NodePoolConfig struct {
	empty                  bool                                  `json:"-"`
	MachineType            *string                               `json:"machineType"`
	DiskSizeGb             *int64                                `json:"diskSizeGb"`
	OAuthScopes            []string                              `json:"oauthScopes"`
	ServiceAccount         *string                               `json:"serviceAccount"`
	Metadata               map[string]string                     `json:"metadata"`
	ImageType              *string                               `json:"imageType"`
	Labels                 map[string]string                     `json:"labels"`
	LocalSsdCount          *int64                                `json:"localSsdCount"`
	Tags                   []string                              `json:"tags"`
	Preemptible            *bool                                 `json:"preemptible"`
	Accelerators           []NodePoolConfigAccelerators          `json:"accelerators"`
	DiskType               *string                               `json:"diskType"`
	MinCpuPlatform         *string                               `json:"minCpuPlatform"`
	Taints                 []NodePoolConfigTaints                `json:"taints"`
	SandboxConfig          *NodePoolConfigSandboxConfig          `json:"sandboxConfig"`
	ReservationAffinity    *NodePoolConfigReservationAffinity    `json:"reservationAffinity"`
	ShieldedInstanceConfig *NodePoolConfigShieldedInstanceConfig `json:"shieldedInstanceConfig"`
}

type jsonNodePoolConfig NodePoolConfig

func (r *NodePoolConfig) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolConfig
	} else {

		r.MachineType = res.MachineType

		r.DiskSizeGb = res.DiskSizeGb

		r.OAuthScopes = res.OAuthScopes

		r.ServiceAccount = res.ServiceAccount

		r.Metadata = res.Metadata

		r.ImageType = res.ImageType

		r.Labels = res.Labels

		r.LocalSsdCount = res.LocalSsdCount

		r.Tags = res.Tags

		r.Preemptible = res.Preemptible

		r.Accelerators = res.Accelerators

		r.DiskType = res.DiskType

		r.MinCpuPlatform = res.MinCpuPlatform

		r.Taints = res.Taints

		r.SandboxConfig = res.SandboxConfig

		r.ReservationAffinity = res.ReservationAffinity

		r.ShieldedInstanceConfig = res.ShieldedInstanceConfig

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolConfig *NodePoolConfig = &NodePoolConfig{empty: true}

func (r *NodePoolConfig) Empty() bool {
	return r.empty
}

func (r *NodePoolConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolConfigAccelerators struct {
	empty            bool    `json:"-"`
	AcceleratorCount *int64  `json:"acceleratorCount"`
	AcceleratorType  *string `json:"acceleratorType"`
}

type jsonNodePoolConfigAccelerators NodePoolConfigAccelerators

func (r *NodePoolConfigAccelerators) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolConfigAccelerators
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolConfigAccelerators
	} else {

		r.AcceleratorCount = res.AcceleratorCount

		r.AcceleratorType = res.AcceleratorType

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolConfigAccelerators is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolConfigAccelerators *NodePoolConfigAccelerators = &NodePoolConfigAccelerators{empty: true}

func (r *NodePoolConfigAccelerators) Empty() bool {
	return r.empty
}

func (r *NodePoolConfigAccelerators) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolConfigAccelerators) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolConfigTaints struct {
	empty  bool    `json:"-"`
	Key    *string `json:"key"`
	Value  *string `json:"value"`
	Effect *string `json:"effect"`
}

type jsonNodePoolConfigTaints NodePoolConfigTaints

func (r *NodePoolConfigTaints) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolConfigTaints
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolConfigTaints
	} else {

		r.Key = res.Key

		r.Value = res.Value

		r.Effect = res.Effect

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolConfigTaints is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolConfigTaints *NodePoolConfigTaints = &NodePoolConfigTaints{empty: true}

func (r *NodePoolConfigTaints) Empty() bool {
	return r.empty
}

func (r *NodePoolConfigTaints) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolConfigTaints) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolConfigSandboxConfig struct {
	empty bool                                 `json:"-"`
	Type  *NodePoolConfigSandboxConfigTypeEnum `json:"type"`
}

type jsonNodePoolConfigSandboxConfig NodePoolConfigSandboxConfig

func (r *NodePoolConfigSandboxConfig) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolConfigSandboxConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolConfigSandboxConfig
	} else {

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolConfigSandboxConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolConfigSandboxConfig *NodePoolConfigSandboxConfig = &NodePoolConfigSandboxConfig{empty: true}

func (r *NodePoolConfigSandboxConfig) Empty() bool {
	return r.empty
}

func (r *NodePoolConfigSandboxConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolConfigSandboxConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolConfigReservationAffinity struct {
	empty                  bool                                                         `json:"-"`
	ConsumeReservationType *NodePoolConfigReservationAffinityConsumeReservationTypeEnum `json:"consumeReservationType"`
	Key                    *string                                                      `json:"key"`
	Values                 []string                                                     `json:"values"`
}

type jsonNodePoolConfigReservationAffinity NodePoolConfigReservationAffinity

func (r *NodePoolConfigReservationAffinity) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolConfigReservationAffinity
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolConfigReservationAffinity
	} else {

		r.ConsumeReservationType = res.ConsumeReservationType

		r.Key = res.Key

		r.Values = res.Values

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolConfigReservationAffinity is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolConfigReservationAffinity *NodePoolConfigReservationAffinity = &NodePoolConfigReservationAffinity{empty: true}

func (r *NodePoolConfigReservationAffinity) Empty() bool {
	return r.empty
}

func (r *NodePoolConfigReservationAffinity) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolConfigReservationAffinity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolConfigShieldedInstanceConfig struct {
	empty                     bool  `json:"-"`
	EnableSecureBoot          *bool `json:"enableSecureBoot"`
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring"`
}

type jsonNodePoolConfigShieldedInstanceConfig NodePoolConfigShieldedInstanceConfig

func (r *NodePoolConfigShieldedInstanceConfig) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolConfigShieldedInstanceConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolConfigShieldedInstanceConfig
	} else {

		r.EnableSecureBoot = res.EnableSecureBoot

		r.EnableIntegrityMonitoring = res.EnableIntegrityMonitoring

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolConfigShieldedInstanceConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolConfigShieldedInstanceConfig *NodePoolConfigShieldedInstanceConfig = &NodePoolConfigShieldedInstanceConfig{empty: true}

func (r *NodePoolConfigShieldedInstanceConfig) Empty() bool {
	return r.empty
}

func (r *NodePoolConfigShieldedInstanceConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolConfigShieldedInstanceConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolAutoscaling struct {
	empty           bool   `json:"-"`
	Enabled         *bool  `json:"enabled"`
	MinNodeCount    *int64 `json:"minNodeCount"`
	MaxNodeCount    *int64 `json:"maxNodeCount"`
	Autoprovisioned *bool  `json:"autoprovisioned"`
}

type jsonNodePoolAutoscaling NodePoolAutoscaling

func (r *NodePoolAutoscaling) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolAutoscaling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolAutoscaling
	} else {

		r.Enabled = res.Enabled

		r.MinNodeCount = res.MinNodeCount

		r.MaxNodeCount = res.MaxNodeCount

		r.Autoprovisioned = res.Autoprovisioned

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolAutoscaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolAutoscaling *NodePoolAutoscaling = &NodePoolAutoscaling{empty: true}

func (r *NodePoolAutoscaling) Empty() bool {
	return r.empty
}

func (r *NodePoolAutoscaling) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolAutoscaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolManagement struct {
	empty          bool                              `json:"-"`
	AutoUpgrade    *bool                             `json:"autoUpgrade"`
	AutoRepair     *bool                             `json:"autoRepair"`
	UpgradeOptions *NodePoolManagementUpgradeOptions `json:"upgradeOptions"`
}

type jsonNodePoolManagement NodePoolManagement

func (r *NodePoolManagement) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolManagement
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolManagement
	} else {

		r.AutoUpgrade = res.AutoUpgrade

		r.AutoRepair = res.AutoRepair

		r.UpgradeOptions = res.UpgradeOptions

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolManagement is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolManagement *NodePoolManagement = &NodePoolManagement{empty: true}

func (r *NodePoolManagement) Empty() bool {
	return r.empty
}

func (r *NodePoolManagement) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolManagement) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolManagementUpgradeOptions struct {
	empty                bool    `json:"-"`
	AutoUpgradeStartTime *string `json:"autoUpgradeStartTime"`
	Description          *string `json:"description"`
}

type jsonNodePoolManagementUpgradeOptions NodePoolManagementUpgradeOptions

func (r *NodePoolManagementUpgradeOptions) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolManagementUpgradeOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolManagementUpgradeOptions
	} else {

		r.AutoUpgradeStartTime = res.AutoUpgradeStartTime

		r.Description = res.Description

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolManagementUpgradeOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolManagementUpgradeOptions *NodePoolManagementUpgradeOptions = &NodePoolManagementUpgradeOptions{empty: true}

func (r *NodePoolManagementUpgradeOptions) Empty() bool {
	return r.empty
}

func (r *NodePoolManagementUpgradeOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolManagementUpgradeOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolMaxPodsConstraint struct {
	empty          bool   `json:"-"`
	MaxPodsPerNode *int64 `json:"maxPodsPerNode"`
}

type jsonNodePoolMaxPodsConstraint NodePoolMaxPodsConstraint

func (r *NodePoolMaxPodsConstraint) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolMaxPodsConstraint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolMaxPodsConstraint
	} else {

		r.MaxPodsPerNode = res.MaxPodsPerNode

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolMaxPodsConstraint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolMaxPodsConstraint *NodePoolMaxPodsConstraint = &NodePoolMaxPodsConstraint{empty: true}

func (r *NodePoolMaxPodsConstraint) Empty() bool {
	return r.empty
}

func (r *NodePoolMaxPodsConstraint) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolMaxPodsConstraint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolConditions struct {
	empty   bool                        `json:"-"`
	Code    *NodePoolConditionsCodeEnum `json:"code"`
	Message *string                     `json:"message"`
}

type jsonNodePoolConditions NodePoolConditions

func (r *NodePoolConditions) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolConditions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolConditions
	} else {

		r.Code = res.Code

		r.Message = res.Message

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolConditions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolConditions *NodePoolConditions = &NodePoolConditions{empty: true}

func (r *NodePoolConditions) Empty() bool {
	return r.empty
}

func (r *NodePoolConditions) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolConditions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodePoolUpgradeSettings struct {
	empty          bool   `json:"-"`
	MaxSurge       *int64 `json:"maxSurge"`
	MaxUnavailable *int64 `json:"maxUnavailable"`
}

type jsonNodePoolUpgradeSettings NodePoolUpgradeSettings

func (r *NodePoolUpgradeSettings) UnmarshalJSON(data []byte) error {
	var res jsonNodePoolUpgradeSettings
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodePoolUpgradeSettings
	} else {

		r.MaxSurge = res.MaxSurge

		r.MaxUnavailable = res.MaxUnavailable

	}
	return nil
}

// This object is used to assert a desired state where this NodePoolUpgradeSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodePoolUpgradeSettings *NodePoolUpgradeSettings = &NodePoolUpgradeSettings{empty: true}

func (r *NodePoolUpgradeSettings) Empty() bool {
	return r.empty
}

func (r *NodePoolUpgradeSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *NodePoolUpgradeSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *NodePool) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "container",
		Type:    "NodePool",
		Version: "beta",
	}
}

const NodePoolMaxPage = -1

type NodePoolList struct {
	Items []*NodePool

	nextToken string

	project string

	location string

	cluster string
}

func (l *NodePoolList) HasNext() bool {
	return l.nextToken != ""
}

func (l *NodePoolList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listNodePool(ctx, l.project, l.location, l.cluster, l.nextToken)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListNodePool(ctx context.Context, project, location, cluster string) (*NodePoolList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listNodePool(ctx, project, location, cluster, "")
	if err != nil {
		return nil, err
	}
	return &NodePoolList{
		Items:     items,
		nextToken: token,

		project: project,

		location: location,

		cluster: cluster,
	}, nil

}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *NodePool) URLNormalized() *NodePool {
	normalized := dcl.Copy(*r).(NodePool)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Version = dcl.SelfLinkToName(r.Version)
	normalized.StatusMessage = dcl.SelfLinkToName(r.StatusMessage)
	normalized.Cluster = dcl.SelfLinkToName(r.Cluster)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetNodePool(ctx context.Context, r *NodePool) (*NodePool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getNodePoolRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalNodePool(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Cluster = r.Cluster
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeNodePoolNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteNodePool(ctx context.Context, r *NodePool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("NodePool resource is nil")
	}
	c.Config.Logger.Info("Deleting NodePool...")
	deleteOp := deleteNodePoolOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllNodePool deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllNodePool(ctx context.Context, project, location, cluster string, filter func(*NodePool) bool) error {
	listObj, err := c.ListNodePool(ctx, project, location, cluster)
	if err != nil {
		return err
	}

	err = c.deleteAllNodePool(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllNodePool(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyNodePool(ctx context.Context, rawDesired *NodePool, opts ...dcl.ApplyOption) (*NodePool, error) {

	var resultNewState *NodePool
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyNodePoolHelper(c, ctx, rawDesired, opts...)
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

func applyNodePoolHelper(c *Client, ctx context.Context, rawDesired *NodePool, opts ...dcl.ApplyOption) (*NodePool, error) {
	c.Config.Logger.Info("Beginning ApplyNodePool...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.nodePoolDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToNodePoolDiffs(c.Config, fieldDiffs, opts)
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
	var ops []nodePoolApiOperation
	if create {
		ops = append(ops, &createNodePoolOperation{})
	} else if recreate {
		ops = append(ops, &deleteNodePoolOperation{})
		ops = append(ops, &createNodePoolOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeNodePoolDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetNodePool(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createNodePoolOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapNodePool(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeNodePoolNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeNodePoolNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeNodePoolDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffNodePool(c, newDesired, newState)
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
