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
package dns

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type ManagedZone struct {
	Description             *string                             `json:"description"`
	DnsName                 *string                             `json:"dnsName"`
	DnssecConfig            *ManagedZoneDnssecConfig            `json:"dnssecConfig"`
	Name                    *string                             `json:"name"`
	NameServers             []string                            `json:"nameServers"`
	Labels                  map[string]string                   `json:"labels"`
	Visibility              *ManagedZoneVisibilityEnum          `json:"visibility"`
	PrivateVisibilityConfig *ManagedZonePrivateVisibilityConfig `json:"privateVisibilityConfig"`
	ForwardingConfig        *ManagedZoneForwardingConfig        `json:"forwardingConfig"`
	ReverseLookup           *bool                               `json:"reverseLookup"`
	PeeringConfig           *ManagedZonePeeringConfig           `json:"peeringConfig"`
	Project                 *string                             `json:"project"`
}

func (r *ManagedZone) String() string {
	return dcl.SprintResource(r)
}

// The enum ManagedZoneDnssecConfigNonExistenceEnum.
type ManagedZoneDnssecConfigNonExistenceEnum string

// ManagedZoneDnssecConfigNonExistenceEnumRef returns a *ManagedZoneDnssecConfigNonExistenceEnum with the value of string s
// If the empty string is provided, nil is returned.
func ManagedZoneDnssecConfigNonExistenceEnumRef(s string) *ManagedZoneDnssecConfigNonExistenceEnum {
	if s == "" {
		return nil
	}

	v := ManagedZoneDnssecConfigNonExistenceEnum(s)
	return &v
}

func (v ManagedZoneDnssecConfigNonExistenceEnum) Validate() error {
	for _, s := range []string{"nsec", "nsec3"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ManagedZoneDnssecConfigNonExistenceEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ManagedZoneDnssecConfigStateEnum.
type ManagedZoneDnssecConfigStateEnum string

// ManagedZoneDnssecConfigStateEnumRef returns a *ManagedZoneDnssecConfigStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ManagedZoneDnssecConfigStateEnumRef(s string) *ManagedZoneDnssecConfigStateEnum {
	if s == "" {
		return nil
	}

	v := ManagedZoneDnssecConfigStateEnum(s)
	return &v
}

func (v ManagedZoneDnssecConfigStateEnum) Validate() error {
	for _, s := range []string{"off", "on", "transfer"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ManagedZoneDnssecConfigStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum.
type ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum string

// ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumRef returns a *ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum with the value of string s
// If the empty string is provided, nil is returned.
func ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumRef(s string) *ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum {
	if s == "" {
		return nil
	}

	v := ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(s)
	return &v
}

func (v ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum) Validate() error {
	for _, s := range []string{"ecdsap256sha256", "ecdsap384sha384", "rsasha1", "rsasha256", "rsasha512"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum.
type ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum string

// ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumRef returns a *ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumRef(s string) *ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum {
	if s == "" {
		return nil
	}

	v := ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(s)
	return &v
}

func (v ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum) Validate() error {
	for _, s := range []string{"keySigning", "zoneSigning"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ManagedZoneVisibilityEnum.
type ManagedZoneVisibilityEnum string

// ManagedZoneVisibilityEnumRef returns a *ManagedZoneVisibilityEnum with the value of string s
// If the empty string is provided, nil is returned.
func ManagedZoneVisibilityEnumRef(s string) *ManagedZoneVisibilityEnum {
	if s == "" {
		return nil
	}

	v := ManagedZoneVisibilityEnum(s)
	return &v
}

func (v ManagedZoneVisibilityEnum) Validate() error {
	for _, s := range []string{"private", "public"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ManagedZoneVisibilityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum.
type ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum string

// ManagedZoneForwardingConfigTargetNameServersForwardingPathEnumRef returns a *ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum with the value of string s
// If the empty string is provided, nil is returned.
func ManagedZoneForwardingConfigTargetNameServersForwardingPathEnumRef(s string) *ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum {
	if s == "" {
		return nil
	}

	v := ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(s)
	return &v
}

func (v ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum) Validate() error {
	for _, s := range []string{"default", "private"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ManagedZoneDnssecConfig struct {
	empty           bool                                     `json:"-"`
	Kind            *string                                  `json:"kind"`
	NonExistence    *ManagedZoneDnssecConfigNonExistenceEnum `json:"nonExistence"`
	State           *ManagedZoneDnssecConfigStateEnum        `json:"state"`
	DefaultKeySpecs []ManagedZoneDnssecConfigDefaultKeySpecs `json:"defaultKeySpecs"`
}

type jsonManagedZoneDnssecConfig ManagedZoneDnssecConfig

func (r *ManagedZoneDnssecConfig) UnmarshalJSON(data []byte) error {
	var res jsonManagedZoneDnssecConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyManagedZoneDnssecConfig
	} else {

		r.Kind = res.Kind

		r.NonExistence = res.NonExistence

		r.State = res.State

		r.DefaultKeySpecs = res.DefaultKeySpecs

	}
	return nil
}

// This object is used to assert a desired state where this ManagedZoneDnssecConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyManagedZoneDnssecConfig *ManagedZoneDnssecConfig = &ManagedZoneDnssecConfig{empty: true}

func (r *ManagedZoneDnssecConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ManagedZoneDnssecConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ManagedZoneDnssecConfigDefaultKeySpecs struct {
	empty     bool                                                 `json:"-"`
	Algorithm *ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum `json:"algorithm"`
	KeyLength *int64                                               `json:"keyLength"`
	KeyType   *ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum   `json:"keyType"`
	Kind      *string                                              `json:"kind"`
}

type jsonManagedZoneDnssecConfigDefaultKeySpecs ManagedZoneDnssecConfigDefaultKeySpecs

func (r *ManagedZoneDnssecConfigDefaultKeySpecs) UnmarshalJSON(data []byte) error {
	var res jsonManagedZoneDnssecConfigDefaultKeySpecs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyManagedZoneDnssecConfigDefaultKeySpecs
	} else {

		r.Algorithm = res.Algorithm

		r.KeyLength = res.KeyLength

		r.KeyType = res.KeyType

		r.Kind = res.Kind

	}
	return nil
}

// This object is used to assert a desired state where this ManagedZoneDnssecConfigDefaultKeySpecs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyManagedZoneDnssecConfigDefaultKeySpecs *ManagedZoneDnssecConfigDefaultKeySpecs = &ManagedZoneDnssecConfigDefaultKeySpecs{empty: true}

func (r *ManagedZoneDnssecConfigDefaultKeySpecs) String() string {
	return dcl.SprintResource(r)
}

func (r *ManagedZoneDnssecConfigDefaultKeySpecs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ManagedZonePrivateVisibilityConfig struct {
	empty    bool                                         `json:"-"`
	Networks []ManagedZonePrivateVisibilityConfigNetworks `json:"networks"`
}

type jsonManagedZonePrivateVisibilityConfig ManagedZonePrivateVisibilityConfig

func (r *ManagedZonePrivateVisibilityConfig) UnmarshalJSON(data []byte) error {
	var res jsonManagedZonePrivateVisibilityConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyManagedZonePrivateVisibilityConfig
	} else {

		r.Networks = res.Networks

	}
	return nil
}

// This object is used to assert a desired state where this ManagedZonePrivateVisibilityConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyManagedZonePrivateVisibilityConfig *ManagedZonePrivateVisibilityConfig = &ManagedZonePrivateVisibilityConfig{empty: true}

func (r *ManagedZonePrivateVisibilityConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ManagedZonePrivateVisibilityConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ManagedZonePrivateVisibilityConfigNetworks struct {
	empty      bool    `json:"-"`
	NetworkUrl *string `json:"networkUrl"`
}

type jsonManagedZonePrivateVisibilityConfigNetworks ManagedZonePrivateVisibilityConfigNetworks

func (r *ManagedZonePrivateVisibilityConfigNetworks) UnmarshalJSON(data []byte) error {
	var res jsonManagedZonePrivateVisibilityConfigNetworks
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyManagedZonePrivateVisibilityConfigNetworks
	} else {

		r.NetworkUrl = res.NetworkUrl

	}
	return nil
}

// This object is used to assert a desired state where this ManagedZonePrivateVisibilityConfigNetworks is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyManagedZonePrivateVisibilityConfigNetworks *ManagedZonePrivateVisibilityConfigNetworks = &ManagedZonePrivateVisibilityConfigNetworks{empty: true}

func (r *ManagedZonePrivateVisibilityConfigNetworks) String() string {
	return dcl.SprintResource(r)
}

func (r *ManagedZonePrivateVisibilityConfigNetworks) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ManagedZoneForwardingConfig struct {
	empty             bool                                           `json:"-"`
	TargetNameServers []ManagedZoneForwardingConfigTargetNameServers `json:"targetNameServers"`
}

type jsonManagedZoneForwardingConfig ManagedZoneForwardingConfig

func (r *ManagedZoneForwardingConfig) UnmarshalJSON(data []byte) error {
	var res jsonManagedZoneForwardingConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyManagedZoneForwardingConfig
	} else {

		r.TargetNameServers = res.TargetNameServers

	}
	return nil
}

// This object is used to assert a desired state where this ManagedZoneForwardingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyManagedZoneForwardingConfig *ManagedZoneForwardingConfig = &ManagedZoneForwardingConfig{empty: true}

func (r *ManagedZoneForwardingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ManagedZoneForwardingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ManagedZoneForwardingConfigTargetNameServers struct {
	empty          bool                                                            `json:"-"`
	IPv4Address    *string                                                         `json:"ipv4Address"`
	ForwardingPath *ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum `json:"forwardingPath"`
}

type jsonManagedZoneForwardingConfigTargetNameServers ManagedZoneForwardingConfigTargetNameServers

func (r *ManagedZoneForwardingConfigTargetNameServers) UnmarshalJSON(data []byte) error {
	var res jsonManagedZoneForwardingConfigTargetNameServers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyManagedZoneForwardingConfigTargetNameServers
	} else {

		r.IPv4Address = res.IPv4Address

		r.ForwardingPath = res.ForwardingPath

	}
	return nil
}

// This object is used to assert a desired state where this ManagedZoneForwardingConfigTargetNameServers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyManagedZoneForwardingConfigTargetNameServers *ManagedZoneForwardingConfigTargetNameServers = &ManagedZoneForwardingConfigTargetNameServers{empty: true}

func (r *ManagedZoneForwardingConfigTargetNameServers) String() string {
	return dcl.SprintResource(r)
}

func (r *ManagedZoneForwardingConfigTargetNameServers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ManagedZonePeeringConfig struct {
	empty         bool                                   `json:"-"`
	TargetNetwork *ManagedZonePeeringConfigTargetNetwork `json:"targetNetwork"`
}

type jsonManagedZonePeeringConfig ManagedZonePeeringConfig

func (r *ManagedZonePeeringConfig) UnmarshalJSON(data []byte) error {
	var res jsonManagedZonePeeringConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyManagedZonePeeringConfig
	} else {

		r.TargetNetwork = res.TargetNetwork

	}
	return nil
}

// This object is used to assert a desired state where this ManagedZonePeeringConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyManagedZonePeeringConfig *ManagedZonePeeringConfig = &ManagedZonePeeringConfig{empty: true}

func (r *ManagedZonePeeringConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ManagedZonePeeringConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ManagedZonePeeringConfigTargetNetwork struct {
	empty      bool    `json:"-"`
	NetworkUrl *string `json:"networkUrl"`
}

type jsonManagedZonePeeringConfigTargetNetwork ManagedZonePeeringConfigTargetNetwork

func (r *ManagedZonePeeringConfigTargetNetwork) UnmarshalJSON(data []byte) error {
	var res jsonManagedZonePeeringConfigTargetNetwork
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyManagedZonePeeringConfigTargetNetwork
	} else {

		r.NetworkUrl = res.NetworkUrl

	}
	return nil
}

// This object is used to assert a desired state where this ManagedZonePeeringConfigTargetNetwork is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyManagedZonePeeringConfigTargetNetwork *ManagedZonePeeringConfigTargetNetwork = &ManagedZonePeeringConfigTargetNetwork{empty: true}

func (r *ManagedZonePeeringConfigTargetNetwork) String() string {
	return dcl.SprintResource(r)
}

func (r *ManagedZonePeeringConfigTargetNetwork) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ManagedZone) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "dns",
		Type:    "ManagedZone",
		Version: "dns",
	}
}

const ManagedZoneMaxPage = -1

type ManagedZoneList struct {
	Items []*ManagedZone

	nextToken string

	pageSize int32

	project string
}

func (l *ManagedZoneList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ManagedZoneList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listManagedZone(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListManagedZone(ctx context.Context, project string) (*ManagedZoneList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListManagedZoneWithMaxResults(ctx, project, ManagedZoneMaxPage)

}

func (c *Client) ListManagedZoneWithMaxResults(ctx context.Context, project string, pageSize int32) (*ManagedZoneList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listManagedZone(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ManagedZoneList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetManagedZone(ctx context.Context, r *ManagedZone) (*ManagedZone, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getManagedZoneRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalManagedZone(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeManagedZoneNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteManagedZone(ctx context.Context, r *ManagedZone) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("ManagedZone resource is nil")
	}
	c.Config.Logger.Info("Deleting ManagedZone...")
	deleteOp := deleteManagedZoneOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllManagedZone deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllManagedZone(ctx context.Context, project string, filter func(*ManagedZone) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListManagedZone(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllManagedZone(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllManagedZone(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyManagedZone(ctx context.Context, rawDesired *ManagedZone, opts ...dcl.ApplyOption) (*ManagedZone, error) {
	c.Config.Logger.Info("Beginning ApplyManagedZone...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.managedZoneDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []managedZoneApiOperation
	if create {
		ops = append(ops, &createManagedZoneOperation{})
	} else if recreate {

		ops = append(ops, &deleteManagedZoneOperation{})

		ops = append(ops, &createManagedZoneOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeManagedZoneDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetManagedZone(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createManagedZoneOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapManagedZone(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeManagedZoneNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeManagedZoneNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeManagedZoneDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffManagedZone(c, newDesired, newState)
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
