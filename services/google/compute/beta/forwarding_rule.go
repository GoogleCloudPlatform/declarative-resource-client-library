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

type ForwardingRule struct {
	AllPorts             *bool                                  `json:"allPorts"`
	AllowGlobalAccess    *bool                                  `json:"allowGlobalAccess"`
	BackendService       *string                                `json:"backendService"`
	CreationTimestamp    *string                                `json:"creationTimestamp"`
	Description          *string                                `json:"description"`
	IPAddress            *string                                `json:"ipAddress"`
	IPProtocol           *ForwardingRuleIPProtocolEnum          `json:"ipProtocol"`
	IPVersion            *ForwardingRuleIPVersionEnum           `json:"ipVersion"`
	IsMirroringCollector *bool                                  `json:"isMirroringCollector"`
	LoadBalancingScheme  *ForwardingRuleLoadBalancingSchemeEnum `json:"loadBalancingScheme"`
	MetadataFilter       []ForwardingRuleMetadataFilter         `json:"metadataFilter"`
	Name                 *string                                `json:"name"`
	Network              *string                                `json:"network"`
	NetworkTier          *ForwardingRuleNetworkTierEnum         `json:"networkTier"`
	PortRange            *string                                `json:"portRange"`
	Ports                []string                               `json:"ports"`
	Region               *string                                `json:"region"`
	SelfLink             *string                                `json:"selfLink"`
	ServiceLabel         *string                                `json:"serviceLabel"`
	ServiceName          *string                                `json:"serviceName"`
	Subnetwork           *string                                `json:"subnetwork"`
	Target               *string                                `json:"target"`
	Project              *string                                `json:"project"`
	Location             *string                                `json:"location"`
}

func (r *ForwardingRule) String() string {
	return dcl.SprintResource(r)
}

// The enum ForwardingRuleIPProtocolEnum.
type ForwardingRuleIPProtocolEnum string

// ForwardingRuleIPProtocolEnumRef returns a *ForwardingRuleIPProtocolEnum with the value of string s
// If the empty string is provided, nil is returned.
func ForwardingRuleIPProtocolEnumRef(s string) *ForwardingRuleIPProtocolEnum {
	if s == "" {
		return nil
	}

	v := ForwardingRuleIPProtocolEnum(s)
	return &v
}

func (v ForwardingRuleIPProtocolEnum) Validate() error {
	for _, s := range []string{"TCP", "UDP", "ESP", "AH", "SCTP", "ICMP"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ForwardingRuleIPProtocolEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ForwardingRuleIPVersionEnum.
type ForwardingRuleIPVersionEnum string

// ForwardingRuleIPVersionEnumRef returns a *ForwardingRuleIPVersionEnum with the value of string s
// If the empty string is provided, nil is returned.
func ForwardingRuleIPVersionEnumRef(s string) *ForwardingRuleIPVersionEnum {
	if s == "" {
		return nil
	}

	v := ForwardingRuleIPVersionEnum(s)
	return &v
}

func (v ForwardingRuleIPVersionEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED_VERSION", "IPV4", "IPV6"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ForwardingRuleIPVersionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ForwardingRuleLoadBalancingSchemeEnum.
type ForwardingRuleLoadBalancingSchemeEnum string

// ForwardingRuleLoadBalancingSchemeEnumRef returns a *ForwardingRuleLoadBalancingSchemeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ForwardingRuleLoadBalancingSchemeEnumRef(s string) *ForwardingRuleLoadBalancingSchemeEnum {
	if s == "" {
		return nil
	}

	v := ForwardingRuleLoadBalancingSchemeEnum(s)
	return &v
}

func (v ForwardingRuleLoadBalancingSchemeEnum) Validate() error {
	for _, s := range []string{"INVALID", "INTERNAL", "INTERNAL_MANAGED", "INTERNAL_SELF_MANAGED", "EXTERNAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ForwardingRuleLoadBalancingSchemeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ForwardingRuleMetadataFilterFilterMatchCriteriaEnum.
type ForwardingRuleMetadataFilterFilterMatchCriteriaEnum string

// ForwardingRuleMetadataFilterFilterMatchCriteriaEnumRef returns a *ForwardingRuleMetadataFilterFilterMatchCriteriaEnum with the value of string s
// If the empty string is provided, nil is returned.
func ForwardingRuleMetadataFilterFilterMatchCriteriaEnumRef(s string) *ForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if s == "" {
		return nil
	}

	v := ForwardingRuleMetadataFilterFilterMatchCriteriaEnum(s)
	return &v
}

func (v ForwardingRuleMetadataFilterFilterMatchCriteriaEnum) Validate() error {
	for _, s := range []string{"NOT_SET", "MATCH_ALL", "MATCH_ANY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ForwardingRuleMetadataFilterFilterMatchCriteriaEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ForwardingRuleNetworkTierEnum.
type ForwardingRuleNetworkTierEnum string

// ForwardingRuleNetworkTierEnumRef returns a *ForwardingRuleNetworkTierEnum with the value of string s
// If the empty string is provided, nil is returned.
func ForwardingRuleNetworkTierEnumRef(s string) *ForwardingRuleNetworkTierEnum {
	if s == "" {
		return nil
	}

	v := ForwardingRuleNetworkTierEnum(s)
	return &v
}

func (v ForwardingRuleNetworkTierEnum) Validate() error {
	for _, s := range []string{"PREMIUM", "STANDARD"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ForwardingRuleNetworkTierEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ForwardingRuleMetadataFilter struct {
	empty               bool                                                 `json:"-"`
	FilterMatchCriteria *ForwardingRuleMetadataFilterFilterMatchCriteriaEnum `json:"filterMatchCriteria"`
	FilterLabel         []ForwardingRuleMetadataFilterFilterLabel            `json:"filterLabel"`
}

type jsonForwardingRuleMetadataFilter ForwardingRuleMetadataFilter

func (r *ForwardingRuleMetadataFilter) UnmarshalJSON(data []byte) error {
	var res jsonForwardingRuleMetadataFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyForwardingRuleMetadataFilter
	} else {

		r.FilterMatchCriteria = res.FilterMatchCriteria

		r.FilterLabel = res.FilterLabel

	}
	return nil
}

// This object is used to assert a desired state where this ForwardingRuleMetadataFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyForwardingRuleMetadataFilter *ForwardingRuleMetadataFilter = &ForwardingRuleMetadataFilter{empty: true}

func (r *ForwardingRuleMetadataFilter) Empty() bool {
	return r.empty
}

func (r *ForwardingRuleMetadataFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *ForwardingRuleMetadataFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ForwardingRuleMetadataFilterFilterLabel struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

type jsonForwardingRuleMetadataFilterFilterLabel ForwardingRuleMetadataFilterFilterLabel

func (r *ForwardingRuleMetadataFilterFilterLabel) UnmarshalJSON(data []byte) error {
	var res jsonForwardingRuleMetadataFilterFilterLabel
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyForwardingRuleMetadataFilterFilterLabel
	} else {

		r.Name = res.Name

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this ForwardingRuleMetadataFilterFilterLabel is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyForwardingRuleMetadataFilterFilterLabel *ForwardingRuleMetadataFilterFilterLabel = &ForwardingRuleMetadataFilterFilterLabel{empty: true}

func (r *ForwardingRuleMetadataFilterFilterLabel) Empty() bool {
	return r.empty
}

func (r *ForwardingRuleMetadataFilterFilterLabel) String() string {
	return dcl.SprintResource(r)
}

func (r *ForwardingRuleMetadataFilterFilterLabel) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ForwardingRule) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "ForwardingRule",
		Version: "beta",
	}
}

const ForwardingRuleMaxPage = -1

type ForwardingRuleList struct {
	Items []*ForwardingRule

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *ForwardingRuleList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ForwardingRuleList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listForwardingRule(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListForwardingRule(ctx context.Context, project, location string) (*ForwardingRuleList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListForwardingRuleWithMaxResults(ctx, project, location, ForwardingRuleMaxPage)

}

func (c *Client) ListForwardingRuleWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ForwardingRuleList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listForwardingRule(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ForwardingRuleList{
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
func (r *ForwardingRule) URLNormalized() *ForwardingRule {
	normalized := dcl.Copy(*r).(ForwardingRule)
	normalized.BackendService = dcl.SelfLinkToName(r.BackendService)
	normalized.CreationTimestamp = dcl.SelfLinkToName(r.CreationTimestamp)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.IPAddress = dcl.SelfLinkToName(r.IPAddress)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.PortRange = dcl.SelfLinkToName(r.PortRange)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.ServiceLabel = dcl.SelfLinkToName(r.ServiceLabel)
	normalized.ServiceName = dcl.SelfLinkToName(r.ServiceName)
	normalized.Subnetwork = dcl.SelfLinkToName(r.Subnetwork)
	normalized.Target = dcl.SelfLinkToName(r.Target)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetForwardingRule(ctx context.Context, r *ForwardingRule) (*ForwardingRule, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getForwardingRuleRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalForwardingRule(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeForwardingRuleNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteForwardingRule(ctx context.Context, r *ForwardingRule) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("ForwardingRule resource is nil")
	}
	c.Config.Logger.Info("Deleting ForwardingRule...")
	deleteOp := deleteForwardingRuleOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllForwardingRule deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllForwardingRule(ctx context.Context, project, location string, filter func(*ForwardingRule) bool) error {
	listObj, err := c.ListForwardingRule(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllForwardingRule(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllForwardingRule(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyForwardingRule(ctx context.Context, rawDesired *ForwardingRule, opts ...dcl.ApplyOption) (*ForwardingRule, error) {

	var resultNewState *ForwardingRule
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyForwardingRuleHelper(c, ctx, rawDesired, opts...)
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

func applyForwardingRuleHelper(c *Client, ctx context.Context, rawDesired *ForwardingRule, opts ...dcl.ApplyOption) (*ForwardingRule, error) {
	c.Config.Logger.Info("Beginning ApplyForwardingRule...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.forwardingRuleDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToForwardingRuleOp(opStrings, fieldDiffs, opts)
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
	var ops []forwardingRuleApiOperation
	if create {
		ops = append(ops, &createForwardingRuleOperation{})
	} else if recreate {
		ops = append(ops, &deleteForwardingRuleOperation{})
		ops = append(ops, &createForwardingRuleOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeForwardingRuleDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetForwardingRule(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createForwardingRuleOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapForwardingRule(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeForwardingRuleNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeForwardingRuleNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeForwardingRuleDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffForwardingRule(c, newDesired, newState)
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
