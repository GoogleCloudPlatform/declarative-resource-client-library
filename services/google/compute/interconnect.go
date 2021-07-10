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

type Interconnect struct {
	Description             *string                            `json:"description"`
	SelfLink                *string                            `json:"selfLink"`
	Id                      *int64                             `json:"id"`
	Name                    *string                            `json:"name"`
	Location                *string                            `json:"location"`
	LinkType                *InterconnectLinkTypeEnum          `json:"linkType"`
	RequestedLinkCount      *int64                             `json:"requestedLinkCount"`
	InterconnectType        *InterconnectInterconnectTypeEnum  `json:"interconnectType"`
	AdminEnabled            *bool                              `json:"adminEnabled"`
	NocContactEmail         *string                            `json:"nocContactEmail"`
	CustomerName            *string                            `json:"customerName"`
	OperationalStatus       *InterconnectOperationalStatusEnum `json:"operationalStatus"`
	ProvisionedLinkCount    *int64                             `json:"provisionedLinkCount"`
	InterconnectAttachments []string                           `json:"interconnectAttachments"`
	PeerIPAddress           *string                            `json:"peerIPAddress"`
	GoogleIPAddress         *string                            `json:"googleIPAddress"`
	GoogleReferenceId       *string                            `json:"googleReferenceId"`
	ExpectedOutages         []InterconnectExpectedOutages      `json:"expectedOutages"`
	CircuitInfos            []InterconnectCircuitInfos         `json:"circuitInfos"`
	State                   *InterconnectStateEnum             `json:"state"`
	Project                 *string                            `json:"project"`
}

func (r *Interconnect) String() string {
	return dcl.SprintResource(r)
}

// The enum InterconnectLinkTypeEnum.
type InterconnectLinkTypeEnum string

// InterconnectLinkTypeEnumRef returns a *InterconnectLinkTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InterconnectLinkTypeEnumRef(s string) *InterconnectLinkTypeEnum {
	if s == "" {
		return nil
	}

	v := InterconnectLinkTypeEnum(s)
	return &v
}

func (v InterconnectLinkTypeEnum) Validate() error {
	for _, s := range []string{"LINK_TYPE_ETHERNET_10G_LR", "LINK_TYPE_ETHERNET_100G_LR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InterconnectLinkTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InterconnectInterconnectTypeEnum.
type InterconnectInterconnectTypeEnum string

// InterconnectInterconnectTypeEnumRef returns a *InterconnectInterconnectTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InterconnectInterconnectTypeEnumRef(s string) *InterconnectInterconnectTypeEnum {
	if s == "" {
		return nil
	}

	v := InterconnectInterconnectTypeEnum(s)
	return &v
}

func (v InterconnectInterconnectTypeEnum) Validate() error {
	for _, s := range []string{"IT_PRIVATE", "PARTNER", "DEDICATED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InterconnectInterconnectTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InterconnectOperationalStatusEnum.
type InterconnectOperationalStatusEnum string

// InterconnectOperationalStatusEnumRef returns a *InterconnectOperationalStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func InterconnectOperationalStatusEnumRef(s string) *InterconnectOperationalStatusEnum {
	if s == "" {
		return nil
	}

	v := InterconnectOperationalStatusEnum(s)
	return &v
}

func (v InterconnectOperationalStatusEnum) Validate() error {
	for _, s := range []string{"OS_ACTIVE", "OS_UNPROVISIONED", "OS_UNDER_MAINTENANCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InterconnectOperationalStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InterconnectExpectedOutagesSourceEnum.
type InterconnectExpectedOutagesSourceEnum string

// InterconnectExpectedOutagesSourceEnumRef returns a *InterconnectExpectedOutagesSourceEnum with the value of string s
// If the empty string is provided, nil is returned.
func InterconnectExpectedOutagesSourceEnumRef(s string) *InterconnectExpectedOutagesSourceEnum {
	if s == "" {
		return nil
	}

	v := InterconnectExpectedOutagesSourceEnum(s)
	return &v
}

func (v InterconnectExpectedOutagesSourceEnum) Validate() error {
	for _, s := range []string{"GOOGLE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InterconnectExpectedOutagesSourceEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InterconnectExpectedOutagesStateEnum.
type InterconnectExpectedOutagesStateEnum string

// InterconnectExpectedOutagesStateEnumRef returns a *InterconnectExpectedOutagesStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func InterconnectExpectedOutagesStateEnumRef(s string) *InterconnectExpectedOutagesStateEnum {
	if s == "" {
		return nil
	}

	v := InterconnectExpectedOutagesStateEnum(s)
	return &v
}

func (v InterconnectExpectedOutagesStateEnum) Validate() error {
	for _, s := range []string{"ACTIVE", "CANCELLED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InterconnectExpectedOutagesStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InterconnectExpectedOutagesIssueTypeEnum.
type InterconnectExpectedOutagesIssueTypeEnum string

// InterconnectExpectedOutagesIssueTypeEnumRef returns a *InterconnectExpectedOutagesIssueTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InterconnectExpectedOutagesIssueTypeEnumRef(s string) *InterconnectExpectedOutagesIssueTypeEnum {
	if s == "" {
		return nil
	}

	v := InterconnectExpectedOutagesIssueTypeEnum(s)
	return &v
}

func (v InterconnectExpectedOutagesIssueTypeEnum) Validate() error {
	for _, s := range []string{"OUTAGE", "PARTIAL_OUTAGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InterconnectExpectedOutagesIssueTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InterconnectStateEnum.
type InterconnectStateEnum string

// InterconnectStateEnumRef returns a *InterconnectStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func InterconnectStateEnumRef(s string) *InterconnectStateEnum {
	if s == "" {
		return nil
	}

	v := InterconnectStateEnum(s)
	return &v
}

func (v InterconnectStateEnum) Validate() error {
	for _, s := range []string{"DEPRECATED", "OBSOLETE", "DELETED", "ACTIVE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InterconnectStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InterconnectExpectedOutages struct {
	empty            bool                                      `json:"-"`
	Name             *string                                   `json:"name"`
	Description      *string                                   `json:"description"`
	Source           *InterconnectExpectedOutagesSourceEnum    `json:"source"`
	State            *InterconnectExpectedOutagesStateEnum     `json:"state"`
	IssueType        *InterconnectExpectedOutagesIssueTypeEnum `json:"issueType"`
	AffectedCircuits []string                                  `json:"affectedCircuits"`
	StartTime        *int64                                    `json:"startTime"`
	EndTime          *int64                                    `json:"endTime"`
}

type jsonInterconnectExpectedOutages InterconnectExpectedOutages

func (r *InterconnectExpectedOutages) UnmarshalJSON(data []byte) error {
	var res jsonInterconnectExpectedOutages
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInterconnectExpectedOutages
	} else {

		r.Name = res.Name

		r.Description = res.Description

		r.Source = res.Source

		r.State = res.State

		r.IssueType = res.IssueType

		r.AffectedCircuits = res.AffectedCircuits

		r.StartTime = res.StartTime

		r.EndTime = res.EndTime

	}
	return nil
}

// This object is used to assert a desired state where this InterconnectExpectedOutages is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInterconnectExpectedOutages *InterconnectExpectedOutages = &InterconnectExpectedOutages{empty: true}

func (r *InterconnectExpectedOutages) Empty() bool {
	return r.empty
}

func (r *InterconnectExpectedOutages) String() string {
	return dcl.SprintResource(r)
}

func (r *InterconnectExpectedOutages) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InterconnectCircuitInfos struct {
	empty            bool    `json:"-"`
	GoogleCircuitId  *string `json:"googleCircuitId"`
	GoogleDemarcId   *string `json:"googleDemarcId"`
	CustomerDemarcId *string `json:"customerDemarcId"`
}

type jsonInterconnectCircuitInfos InterconnectCircuitInfos

func (r *InterconnectCircuitInfos) UnmarshalJSON(data []byte) error {
	var res jsonInterconnectCircuitInfos
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInterconnectCircuitInfos
	} else {

		r.GoogleCircuitId = res.GoogleCircuitId

		r.GoogleDemarcId = res.GoogleDemarcId

		r.CustomerDemarcId = res.CustomerDemarcId

	}
	return nil
}

// This object is used to assert a desired state where this InterconnectCircuitInfos is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInterconnectCircuitInfos *InterconnectCircuitInfos = &InterconnectCircuitInfos{empty: true}

func (r *InterconnectCircuitInfos) Empty() bool {
	return r.empty
}

func (r *InterconnectCircuitInfos) String() string {
	return dcl.SprintResource(r)
}

func (r *InterconnectCircuitInfos) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Interconnect) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Interconnect",
		Version: "compute",
	}
}

const InterconnectMaxPage = -1

type InterconnectList struct {
	Items []*Interconnect

	nextToken string

	pageSize int32

	project string
}

func (l *InterconnectList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InterconnectList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInterconnect(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInterconnect(ctx context.Context, project string) (*InterconnectList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListInterconnectWithMaxResults(ctx, project, InterconnectMaxPage)

}

func (c *Client) ListInterconnectWithMaxResults(ctx context.Context, project string, pageSize int32) (*InterconnectList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listInterconnect(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InterconnectList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Interconnect) URLNormalized() *Interconnect {
	normalized := dcl.Copy(*r).(Interconnect)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.NocContactEmail = dcl.SelfLinkToName(r.NocContactEmail)
	normalized.CustomerName = dcl.SelfLinkToName(r.CustomerName)
	normalized.PeerIPAddress = dcl.SelfLinkToName(r.PeerIPAddress)
	normalized.GoogleIPAddress = dcl.SelfLinkToName(r.GoogleIPAddress)
	normalized.GoogleReferenceId = dcl.SelfLinkToName(r.GoogleReferenceId)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetInterconnect(ctx context.Context, r *Interconnect) (*Interconnect, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getInterconnectRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalInterconnect(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeInterconnectNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteInterconnect(ctx context.Context, r *Interconnect) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Interconnect resource is nil")
	}
	c.Config.Logger.Info("Deleting Interconnect...")
	deleteOp := deleteInterconnectOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInterconnect deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInterconnect(ctx context.Context, project string, filter func(*Interconnect) bool) error {
	listObj, err := c.ListInterconnect(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllInterconnect(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllInterconnect(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyInterconnect(ctx context.Context, rawDesired *Interconnect, opts ...dcl.ApplyOption) (*Interconnect, error) {

	var resultNewState *Interconnect
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyInterconnectHelper(c, ctx, rawDesired, opts...)
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

func applyInterconnectHelper(c *Client, ctx context.Context, rawDesired *Interconnect, opts ...dcl.ApplyOption) (*Interconnect, error) {
	c.Config.Logger.Info("Beginning ApplyInterconnect...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.interconnectDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToInterconnectOp(opStrings, fieldDiffs, opts)
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
	var ops []interconnectApiOperation
	if create {
		ops = append(ops, &createInterconnectOperation{})
	} else if recreate {
		ops = append(ops, &deleteInterconnectOperation{})
		ops = append(ops, &createInterconnectOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeInterconnectDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetInterconnect(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createInterconnectOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapInterconnect(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeInterconnectNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeInterconnectNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeInterconnectDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffInterconnect(c, newDesired, newState)
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
