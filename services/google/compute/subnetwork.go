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

type Subnetwork struct {
	CreationTimestamp     *string                       `json:"creationTimestamp"`
	Description           *string                       `json:"description"`
	GatewayAddress        *string                       `json:"gatewayAddress"`
	IPCidrRange           *string                       `json:"ipCidrRange"`
	Name                  *string                       `json:"name"`
	Network               *string                       `json:"network"`
	Fingerprint           *string                       `json:"fingerprint"`
	Purpose               *SubnetworkPurposeEnum        `json:"purpose"`
	Role                  *SubnetworkRoleEnum           `json:"role"`
	SecondaryIPRanges     []SubnetworkSecondaryIPRanges `json:"secondaryIPRanges"`
	PrivateIPGoogleAccess *bool                         `json:"privateIPGoogleAccess"`
	Region                *string                       `json:"region"`
	LogConfig             *SubnetworkLogConfig          `json:"logConfig"`
	Project               *string                       `json:"project"`
	SelfLink              *string                       `json:"selfLink"`
	EnableFlowLogs        *bool                         `json:"enableFlowLogs"`
}

func (r *Subnetwork) String() string {
	return dcl.SprintResource(r)
}

// The enum SubnetworkPurposeEnum.
type SubnetworkPurposeEnum string

// SubnetworkPurposeEnumRef returns a *SubnetworkPurposeEnum with the value of string s
// If the empty string is provided, nil is returned.
func SubnetworkPurposeEnumRef(s string) *SubnetworkPurposeEnum {
	if s == "" {
		return nil
	}

	v := SubnetworkPurposeEnum(s)
	return &v
}

func (v SubnetworkPurposeEnum) Validate() error {
	for _, s := range []string{"INTERNAL_HTTPS_LOAD_BALANCER", "PRIVATE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "SubnetworkPurposeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum SubnetworkRoleEnum.
type SubnetworkRoleEnum string

// SubnetworkRoleEnumRef returns a *SubnetworkRoleEnum with the value of string s
// If the empty string is provided, nil is returned.
func SubnetworkRoleEnumRef(s string) *SubnetworkRoleEnum {
	if s == "" {
		return nil
	}

	v := SubnetworkRoleEnum(s)
	return &v
}

func (v SubnetworkRoleEnum) Validate() error {
	for _, s := range []string{"ACTIVE", "BACKUP"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "SubnetworkRoleEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum SubnetworkLogConfigAggregationIntervalEnum.
type SubnetworkLogConfigAggregationIntervalEnum string

// SubnetworkLogConfigAggregationIntervalEnumRef returns a *SubnetworkLogConfigAggregationIntervalEnum with the value of string s
// If the empty string is provided, nil is returned.
func SubnetworkLogConfigAggregationIntervalEnumRef(s string) *SubnetworkLogConfigAggregationIntervalEnum {
	if s == "" {
		return nil
	}

	v := SubnetworkLogConfigAggregationIntervalEnum(s)
	return &v
}

func (v SubnetworkLogConfigAggregationIntervalEnum) Validate() error {
	for _, s := range []string{"INTERVAL_5_SEC", "INTERVAL_30_SEC", "INTERVAL_1_MIN", "INTERVAL_5_MIN", "INTERVAL_10_MIN", "INTERVAL_15_MIN"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "SubnetworkLogConfigAggregationIntervalEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum SubnetworkLogConfigMetadataEnum.
type SubnetworkLogConfigMetadataEnum string

// SubnetworkLogConfigMetadataEnumRef returns a *SubnetworkLogConfigMetadataEnum with the value of string s
// If the empty string is provided, nil is returned.
func SubnetworkLogConfigMetadataEnumRef(s string) *SubnetworkLogConfigMetadataEnum {
	if s == "" {
		return nil
	}

	v := SubnetworkLogConfigMetadataEnum(s)
	return &v
}

func (v SubnetworkLogConfigMetadataEnum) Validate() error {
	for _, s := range []string{"EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "SubnetworkLogConfigMetadataEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type SubnetworkSecondaryIPRanges struct {
	empty       bool    `json:"-"`
	RangeName   *string `json:"rangeName"`
	IPCidrRange *string `json:"ipCidrRange"`
}

type jsonSubnetworkSecondaryIPRanges SubnetworkSecondaryIPRanges

func (r *SubnetworkSecondaryIPRanges) UnmarshalJSON(data []byte) error {
	var res jsonSubnetworkSecondaryIPRanges
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptySubnetworkSecondaryIPRanges
	} else {

		r.RangeName = res.RangeName

		r.IPCidrRange = res.IPCidrRange

	}
	return nil
}

// This object is used to assert a desired state where this SubnetworkSecondaryIPRanges is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySubnetworkSecondaryIPRanges *SubnetworkSecondaryIPRanges = &SubnetworkSecondaryIPRanges{empty: true}

func (r *SubnetworkSecondaryIPRanges) Empty() bool {
	return r.empty
}

func (r *SubnetworkSecondaryIPRanges) String() string {
	return dcl.SprintResource(r)
}

func (r *SubnetworkSecondaryIPRanges) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type SubnetworkLogConfig struct {
	empty               bool                                        `json:"-"`
	AggregationInterval *SubnetworkLogConfigAggregationIntervalEnum `json:"aggregationInterval"`
	FlowSampling        *float64                                    `json:"flowSampling"`
	Metadata            *SubnetworkLogConfigMetadataEnum            `json:"metadata"`
}

type jsonSubnetworkLogConfig SubnetworkLogConfig

func (r *SubnetworkLogConfig) UnmarshalJSON(data []byte) error {
	var res jsonSubnetworkLogConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptySubnetworkLogConfig
	} else {

		r.AggregationInterval = res.AggregationInterval

		r.FlowSampling = res.FlowSampling

		r.Metadata = res.Metadata

	}
	return nil
}

// This object is used to assert a desired state where this SubnetworkLogConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySubnetworkLogConfig *SubnetworkLogConfig = &SubnetworkLogConfig{empty: true}

func (r *SubnetworkLogConfig) Empty() bool {
	return r.empty
}

func (r *SubnetworkLogConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *SubnetworkLogConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Subnetwork) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Subnetwork",
		Version: "compute",
	}
}

const SubnetworkMaxPage = -1

type SubnetworkList struct {
	Items []*Subnetwork

	nextToken string

	pageSize int32

	project string

	region string
}

func (l *SubnetworkList) HasNext() bool {
	return l.nextToken != ""
}

func (l *SubnetworkList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listSubnetwork(ctx, l.project, l.region, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListSubnetwork(ctx context.Context, project, region string) (*SubnetworkList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListSubnetworkWithMaxResults(ctx, project, region, SubnetworkMaxPage)

}

func (c *Client) ListSubnetworkWithMaxResults(ctx context.Context, project, region string, pageSize int32) (*SubnetworkList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listSubnetwork(ctx, project, region, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &SubnetworkList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		region: region,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Subnetwork) URLNormalized() *Subnetwork {
	normalized := dcl.Copy(*r).(Subnetwork)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.GatewayAddress = dcl.SelfLinkToName(r.GatewayAddress)
	normalized.IPCidrRange = dcl.SelfLinkToName(r.IPCidrRange)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.Fingerprint = dcl.SelfLinkToName(r.Fingerprint)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	return &normalized
}

func (c *Client) GetSubnetwork(ctx context.Context, r *Subnetwork) (*Subnetwork, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getSubnetworkRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalSubnetwork(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Region = r.Region
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeSubnetworkNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteSubnetwork(ctx context.Context, r *Subnetwork) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Subnetwork resource is nil")
	}
	c.Config.Logger.Info("Deleting Subnetwork...")
	deleteOp := deleteSubnetworkOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllSubnetwork deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllSubnetwork(ctx context.Context, project, region string, filter func(*Subnetwork) bool) error {
	listObj, err := c.ListSubnetwork(ctx, project, region)
	if err != nil {
		return err
	}

	err = c.deleteAllSubnetwork(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllSubnetwork(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplySubnetwork(ctx context.Context, rawDesired *Subnetwork, opts ...dcl.ApplyOption) (*Subnetwork, error) {
	var resultNewState *Subnetwork
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applySubnetworkHelper(c, ctx, rawDesired, opts...)
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

func applySubnetworkHelper(c *Client, ctx context.Context, rawDesired *Subnetwork, opts ...dcl.ApplyOption) (*Subnetwork, error) {
	c.Config.Logger.Info("Beginning ApplySubnetwork...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.subnetworkDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToSubnetworkDiffs(c.Config, fieldDiffs, opts)
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
	var ops []subnetworkApiOperation
	if create {
		ops = append(ops, &createSubnetworkOperation{})
	} else if recreate {
		ops = append(ops, &deleteSubnetworkOperation{})
		ops = append(ops, &createSubnetworkOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeSubnetworkDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetSubnetwork(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createSubnetworkOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapSubnetwork(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeSubnetworkNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeSubnetworkNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeSubnetworkDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffSubnetwork(c, newDesired, newState)
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
