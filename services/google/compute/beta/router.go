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

type Router struct {
	CreationTimestamp *string            `json:"creationTimestamp"`
	Nats              []RouterNats       `json:"nats"`
	Name              *string            `json:"name"`
	Network           *string            `json:"network"`
	Interfaces        []RouterInterfaces `json:"interfaces"`
	Description       *string            `json:"description"`
	BgpPeers          []RouterBgpPeers   `json:"bgpPeers"`
	Bgp               *RouterBgp         `json:"bgp"`
	Region            *string            `json:"region"`
	Project           *string            `json:"project"`
	SelfLink          *string            `json:"selfLink"`
}

func (r *Router) String() string {
	return dcl.SprintResource(r)
}

// The enum RouterNatsLogConfigFilterEnum.
type RouterNatsLogConfigFilterEnum string

// RouterNatsLogConfigFilterEnumRef returns a *RouterNatsLogConfigFilterEnum with the value of string s
// If the empty string is provided, nil is returned.
func RouterNatsLogConfigFilterEnumRef(s string) *RouterNatsLogConfigFilterEnum {
	if s == "" {
		return nil
	}

	v := RouterNatsLogConfigFilterEnum(s)
	return &v
}

func (v RouterNatsLogConfigFilterEnum) Validate() error {
	for _, s := range []string{"ERRORS_ONLY", "TRANSLATIONS_ONLY", "ALL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RouterNatsLogConfigFilterEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RouterNatsSourceSubnetworkIPRangesToNatEnum.
type RouterNatsSourceSubnetworkIPRangesToNatEnum string

// RouterNatsSourceSubnetworkIPRangesToNatEnumRef returns a *RouterNatsSourceSubnetworkIPRangesToNatEnum with the value of string s
// If the empty string is provided, nil is returned.
func RouterNatsSourceSubnetworkIPRangesToNatEnumRef(s string) *RouterNatsSourceSubnetworkIPRangesToNatEnum {
	if s == "" {
		return nil
	}

	v := RouterNatsSourceSubnetworkIPRangesToNatEnum(s)
	return &v
}

func (v RouterNatsSourceSubnetworkIPRangesToNatEnum) Validate() error {
	for _, s := range []string{"ALL_SUBNETWORKS_ALL_IP_RANGES", "ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES", "LIST_OF_SUBNETWORKS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RouterNatsSourceSubnetworkIPRangesToNatEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RouterNatsNatIPAllocateOptionEnum.
type RouterNatsNatIPAllocateOptionEnum string

// RouterNatsNatIPAllocateOptionEnumRef returns a *RouterNatsNatIPAllocateOptionEnum with the value of string s
// If the empty string is provided, nil is returned.
func RouterNatsNatIPAllocateOptionEnumRef(s string) *RouterNatsNatIPAllocateOptionEnum {
	if s == "" {
		return nil
	}

	v := RouterNatsNatIPAllocateOptionEnum(s)
	return &v
}

func (v RouterNatsNatIPAllocateOptionEnum) Validate() error {
	for _, s := range []string{"MANUAL_ONLY", "AUTO_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RouterNatsNatIPAllocateOptionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RouterInterfacesManagementTypeEnum.
type RouterInterfacesManagementTypeEnum string

// RouterInterfacesManagementTypeEnumRef returns a *RouterInterfacesManagementTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func RouterInterfacesManagementTypeEnumRef(s string) *RouterInterfacesManagementTypeEnum {
	if s == "" {
		return nil
	}

	v := RouterInterfacesManagementTypeEnum(s)
	return &v
}

func (v RouterInterfacesManagementTypeEnum) Validate() error {
	for _, s := range []string{"MANAGED_BY_USER", "MANAGED_BY_ATTACHMENT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RouterInterfacesManagementTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RouterBgpPeersAdvertisedGroupsEnum.
type RouterBgpPeersAdvertisedGroupsEnum string

// RouterBgpPeersAdvertisedGroupsEnumRef returns a *RouterBgpPeersAdvertisedGroupsEnum with the value of string s
// If the empty string is provided, nil is returned.
func RouterBgpPeersAdvertisedGroupsEnumRef(s string) *RouterBgpPeersAdvertisedGroupsEnum {
	if s == "" {
		return nil
	}

	v := RouterBgpPeersAdvertisedGroupsEnum(s)
	return &v
}

func (v RouterBgpPeersAdvertisedGroupsEnum) Validate() error {
	for _, s := range []string{"ALL_SUBNETS", "ALL_VPC_SUBNETS", "ALL_PEER_VPC_SUBNETS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RouterBgpPeersAdvertisedGroupsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RouterBgpAdvertiseModeEnum.
type RouterBgpAdvertiseModeEnum string

// RouterBgpAdvertiseModeEnumRef returns a *RouterBgpAdvertiseModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func RouterBgpAdvertiseModeEnumRef(s string) *RouterBgpAdvertiseModeEnum {
	if s == "" {
		return nil
	}

	v := RouterBgpAdvertiseModeEnum(s)
	return &v
}

func (v RouterBgpAdvertiseModeEnum) Validate() error {
	for _, s := range []string{"DEFAULT", "CUSTOM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RouterBgpAdvertiseModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type RouterNats struct {
	empty                         bool                                         `json:"-"`
	Name                          *string                                      `json:"name"`
	LogConfig                     *RouterNatsLogConfig                         `json:"logConfig"`
	SourceSubnetworkIPRangesToNat *RouterNatsSourceSubnetworkIPRangesToNatEnum `json:"sourceSubnetworkIPRangesToNat"`
	NatIps                        []string                                     `json:"natIps"`
	DrainNatIps                   []string                                     `json:"drainNatIps"`
	NatIPAllocateOption           []RouterNatsNatIPAllocateOptionEnum          `json:"natIPAllocateOption"`
	MinPortsPerVm                 *int64                                       `json:"minPortsPerVm"`
	UdpIdleTimeoutSec             *int64                                       `json:"udpIdleTimeoutSec"`
	IcmpIdleTimeoutSec            *int64                                       `json:"icmpIdleTimeoutSec"`
	TcpEstablishedIdleTimeoutSec  *int64                                       `json:"tcpEstablishedIdleTimeoutSec"`
	TcpTransitoryIdleTimeoutSec   *int64                                       `json:"tcpTransitoryIdleTimeoutSec"`
	Subnetworks                   []RouterNatsSubnetworks                      `json:"subnetworks"`
}

type jsonRouterNats RouterNats

func (r *RouterNats) UnmarshalJSON(data []byte) error {
	var res jsonRouterNats
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouterNats
	} else {

		r.Name = res.Name

		r.LogConfig = res.LogConfig

		r.SourceSubnetworkIPRangesToNat = res.SourceSubnetworkIPRangesToNat

		r.NatIps = res.NatIps

		r.DrainNatIps = res.DrainNatIps

		r.NatIPAllocateOption = res.NatIPAllocateOption

		r.MinPortsPerVm = res.MinPortsPerVm

		r.UdpIdleTimeoutSec = res.UdpIdleTimeoutSec

		r.IcmpIdleTimeoutSec = res.IcmpIdleTimeoutSec

		r.TcpEstablishedIdleTimeoutSec = res.TcpEstablishedIdleTimeoutSec

		r.TcpTransitoryIdleTimeoutSec = res.TcpTransitoryIdleTimeoutSec

		r.Subnetworks = res.Subnetworks

	}
	return nil
}

// This object is used to assert a desired state where this RouterNats is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouterNats *RouterNats = &RouterNats{empty: true}

func (r *RouterNats) Empty() bool {
	return r.empty
}

func (r *RouterNats) String() string {
	return dcl.SprintResource(r)
}

func (r *RouterNats) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RouterNatsLogConfig struct {
	empty  bool                           `json:"-"`
	Enable *bool                          `json:"enable"`
	Filter *RouterNatsLogConfigFilterEnum `json:"filter"`
}

type jsonRouterNatsLogConfig RouterNatsLogConfig

func (r *RouterNatsLogConfig) UnmarshalJSON(data []byte) error {
	var res jsonRouterNatsLogConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouterNatsLogConfig
	} else {

		r.Enable = res.Enable

		r.Filter = res.Filter

	}
	return nil
}

// This object is used to assert a desired state where this RouterNatsLogConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouterNatsLogConfig *RouterNatsLogConfig = &RouterNatsLogConfig{empty: true}

func (r *RouterNatsLogConfig) Empty() bool {
	return r.empty
}

func (r *RouterNatsLogConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *RouterNatsLogConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RouterNatsSubnetworks struct {
	empty                 bool    `json:"-"`
	Name                  *string `json:"name"`
	SourceIPRangesToNat   *string `json:"sourceIPRangesToNat"`
	SecondaryIPRangeNames *string `json:"secondaryIPRangeNames"`
}

type jsonRouterNatsSubnetworks RouterNatsSubnetworks

func (r *RouterNatsSubnetworks) UnmarshalJSON(data []byte) error {
	var res jsonRouterNatsSubnetworks
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouterNatsSubnetworks
	} else {

		r.Name = res.Name

		r.SourceIPRangesToNat = res.SourceIPRangesToNat

		r.SecondaryIPRangeNames = res.SecondaryIPRangeNames

	}
	return nil
}

// This object is used to assert a desired state where this RouterNatsSubnetworks is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouterNatsSubnetworks *RouterNatsSubnetworks = &RouterNatsSubnetworks{empty: true}

func (r *RouterNatsSubnetworks) Empty() bool {
	return r.empty
}

func (r *RouterNatsSubnetworks) String() string {
	return dcl.SprintResource(r)
}

func (r *RouterNatsSubnetworks) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RouterInterfaces struct {
	empty           bool                                `json:"-"`
	Name            *string                             `json:"name"`
	LinkedVpnTunnel *string                             `json:"linkedVpnTunnel"`
	IPRange         *string                             `json:"ipRange"`
	ManagementType  *RouterInterfacesManagementTypeEnum `json:"managementType"`
}

type jsonRouterInterfaces RouterInterfaces

func (r *RouterInterfaces) UnmarshalJSON(data []byte) error {
	var res jsonRouterInterfaces
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouterInterfaces
	} else {

		r.Name = res.Name

		r.LinkedVpnTunnel = res.LinkedVpnTunnel

		r.IPRange = res.IPRange

		r.ManagementType = res.ManagementType

	}
	return nil
}

// This object is used to assert a desired state where this RouterInterfaces is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouterInterfaces *RouterInterfaces = &RouterInterfaces{empty: true}

func (r *RouterInterfaces) Empty() bool {
	return r.empty
}

func (r *RouterInterfaces) String() string {
	return dcl.SprintResource(r)
}

func (r *RouterInterfaces) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RouterBgpPeers struct {
	empty                   bool                                 `json:"-"`
	Name                    *string                              `json:"name"`
	InterfaceName           *string                              `json:"interfaceName"`
	IPAddress               *string                              `json:"ipAddress"`
	PeerIPAddress           *string                              `json:"peerIPAddress"`
	PeerAsn                 *int64                               `json:"peerAsn"`
	AdvertisedRoutePriority *int64                               `json:"advertisedRoutePriority"`
	AdvertiseMode           *string                              `json:"advertiseMode"`
	ManagementType          *string                              `json:"managementType"`
	AdvertisedGroups        []RouterBgpPeersAdvertisedGroupsEnum `json:"advertisedGroups"`
	AdvertisedIPRanges      []RouterBgpPeersAdvertisedIPRanges   `json:"advertisedIPRanges"`
}

type jsonRouterBgpPeers RouterBgpPeers

func (r *RouterBgpPeers) UnmarshalJSON(data []byte) error {
	var res jsonRouterBgpPeers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouterBgpPeers
	} else {

		r.Name = res.Name

		r.InterfaceName = res.InterfaceName

		r.IPAddress = res.IPAddress

		r.PeerIPAddress = res.PeerIPAddress

		r.PeerAsn = res.PeerAsn

		r.AdvertisedRoutePriority = res.AdvertisedRoutePriority

		r.AdvertiseMode = res.AdvertiseMode

		r.ManagementType = res.ManagementType

		r.AdvertisedGroups = res.AdvertisedGroups

		r.AdvertisedIPRanges = res.AdvertisedIPRanges

	}
	return nil
}

// This object is used to assert a desired state where this RouterBgpPeers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouterBgpPeers *RouterBgpPeers = &RouterBgpPeers{empty: true}

func (r *RouterBgpPeers) Empty() bool {
	return r.empty
}

func (r *RouterBgpPeers) String() string {
	return dcl.SprintResource(r)
}

func (r *RouterBgpPeers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RouterBgpPeersAdvertisedIPRanges struct {
	empty       bool    `json:"-"`
	Range       *string `json:"range"`
	Description *string `json:"description"`
}

type jsonRouterBgpPeersAdvertisedIPRanges RouterBgpPeersAdvertisedIPRanges

func (r *RouterBgpPeersAdvertisedIPRanges) UnmarshalJSON(data []byte) error {
	var res jsonRouterBgpPeersAdvertisedIPRanges
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouterBgpPeersAdvertisedIPRanges
	} else {

		r.Range = res.Range

		r.Description = res.Description

	}
	return nil
}

// This object is used to assert a desired state where this RouterBgpPeersAdvertisedIPRanges is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouterBgpPeersAdvertisedIPRanges *RouterBgpPeersAdvertisedIPRanges = &RouterBgpPeersAdvertisedIPRanges{empty: true}

func (r *RouterBgpPeersAdvertisedIPRanges) Empty() bool {
	return r.empty
}

func (r *RouterBgpPeersAdvertisedIPRanges) String() string {
	return dcl.SprintResource(r)
}

func (r *RouterBgpPeersAdvertisedIPRanges) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RouterBgp struct {
	empty              bool                          `json:"-"`
	Asn                *int64                        `json:"asn"`
	AdvertiseMode      *RouterBgpAdvertiseModeEnum   `json:"advertiseMode"`
	AdvertisedGroups   []string                      `json:"advertisedGroups"`
	AdvertisedIPRanges []RouterBgpAdvertisedIPRanges `json:"advertisedIPRanges"`
}

type jsonRouterBgp RouterBgp

func (r *RouterBgp) UnmarshalJSON(data []byte) error {
	var res jsonRouterBgp
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouterBgp
	} else {

		r.Asn = res.Asn

		r.AdvertiseMode = res.AdvertiseMode

		r.AdvertisedGroups = res.AdvertisedGroups

		r.AdvertisedIPRanges = res.AdvertisedIPRanges

	}
	return nil
}

// This object is used to assert a desired state where this RouterBgp is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouterBgp *RouterBgp = &RouterBgp{empty: true}

func (r *RouterBgp) Empty() bool {
	return r.empty
}

func (r *RouterBgp) String() string {
	return dcl.SprintResource(r)
}

func (r *RouterBgp) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RouterBgpAdvertisedIPRanges struct {
	empty       bool    `json:"-"`
	Range       *string `json:"range"`
	Description *string `json:"description"`
}

type jsonRouterBgpAdvertisedIPRanges RouterBgpAdvertisedIPRanges

func (r *RouterBgpAdvertisedIPRanges) UnmarshalJSON(data []byte) error {
	var res jsonRouterBgpAdvertisedIPRanges
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouterBgpAdvertisedIPRanges
	} else {

		r.Range = res.Range

		r.Description = res.Description

	}
	return nil
}

// This object is used to assert a desired state where this RouterBgpAdvertisedIPRanges is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouterBgpAdvertisedIPRanges *RouterBgpAdvertisedIPRanges = &RouterBgpAdvertisedIPRanges{empty: true}

func (r *RouterBgpAdvertisedIPRanges) Empty() bool {
	return r.empty
}

func (r *RouterBgpAdvertisedIPRanges) String() string {
	return dcl.SprintResource(r)
}

func (r *RouterBgpAdvertisedIPRanges) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Router) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Router",
		Version: "beta",
	}
}

const RouterMaxPage = -1

type RouterList struct {
	Items []*Router

	nextToken string

	pageSize int32

	project string

	region string
}

func (l *RouterList) HasNext() bool {
	return l.nextToken != ""
}

func (l *RouterList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listRouter(ctx, l.project, l.region, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListRouter(ctx context.Context, project, region string) (*RouterList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListRouterWithMaxResults(ctx, project, region, RouterMaxPage)

}

func (c *Client) ListRouterWithMaxResults(ctx context.Context, project, region string, pageSize int32) (*RouterList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listRouter(ctx, project, region, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &RouterList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		region: region,
	}, nil
}

func (c *Client) GetRouter(ctx context.Context, r *Router) (*Router, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getRouterRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalRouter(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Region = r.Region
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeRouterNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteRouter(ctx context.Context, r *Router) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Router resource is nil")
	}
	c.Config.Logger.Info("Deleting Router...")
	deleteOp := deleteRouterOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllRouter deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllRouter(ctx context.Context, project, region string, filter func(*Router) bool) error {
	listObj, err := c.ListRouter(ctx, project, region)
	if err != nil {
		return err
	}

	err = c.deleteAllRouter(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllRouter(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyRouter(ctx context.Context, rawDesired *Router, opts ...dcl.ApplyOption) (*Router, error) {

	var resultNewState *Router
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyRouterHelper(c, ctx, rawDesired, opts...)
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

func applyRouterHelper(c *Client, ctx context.Context, rawDesired *Router, opts ...dcl.ApplyOption) (*Router, error) {
	c.Config.Logger.Info("Beginning ApplyRouter...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.routerDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []routerApiOperation
	if create {
		ops = append(ops, &createRouterOperation{})
	} else if recreate {

		ops = append(ops, &deleteRouterOperation{})

		ops = append(ops, &createRouterOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeRouterDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetRouter(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createRouterOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapRouter(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeRouterNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeRouterNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeRouterDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffRouter(c, newDesired, newState)
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
