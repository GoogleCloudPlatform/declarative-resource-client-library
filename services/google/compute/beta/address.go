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
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Address struct {
	Id                *int64                  `json:"id"`
	Name              *string                 `json:"name"`
	Description       *string                 `json:"description"`
	Address           *string                 `json:"address"`
	PrefixLength      *int64                  `json:"prefixLength"`
	Status            *AddressStatusEnum      `json:"status"`
	Region            *string                 `json:"region"`
	SelfLink          *string                 `json:"selfLink"`
	NetworkTier       *AddressNetworkTierEnum `json:"networkTier"`
	IPVersion         *AddressIPVersionEnum   `json:"ipVersion"`
	AddressType       *AddressAddressTypeEnum `json:"addressType"`
	Purpose           *AddressPurposeEnum     `json:"purpose"`
	Subnetwork        *string                 `json:"subnetwork"`
	Network           *string                 `json:"network"`
	Project           *string                 `json:"project"`
	CreationTimestamp *string                 `json:"creationTimestamp"`
	Users             []string                `json:"users"`
	Labels            map[string]string       `json:"labels"`
	LabelFingerprint  *string                 `json:"labelFingerprint"`
	Location          *string                 `json:"location"`
}

func (r *Address) String() string {
	return dcl.SprintResource(r)
}

// The enum AddressStatusEnum.
type AddressStatusEnum string

// AddressStatusEnumRef returns a *AddressStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func AddressStatusEnumRef(s string) *AddressStatusEnum {
	if s == "" {
		return nil
	}

	v := AddressStatusEnum(s)
	return &v
}

func (v AddressStatusEnum) Validate() error {
	for _, s := range []string{"PENDING", "RUNNING", "DONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AddressStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AddressNetworkTierEnum.
type AddressNetworkTierEnum string

// AddressNetworkTierEnumRef returns a *AddressNetworkTierEnum with the value of string s
// If the empty string is provided, nil is returned.
func AddressNetworkTierEnumRef(s string) *AddressNetworkTierEnum {
	if s == "" {
		return nil
	}

	v := AddressNetworkTierEnum(s)
	return &v
}

func (v AddressNetworkTierEnum) Validate() error {
	for _, s := range []string{"PREMIUM", "STANDARD"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AddressNetworkTierEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AddressIPVersionEnum.
type AddressIPVersionEnum string

// AddressIPVersionEnumRef returns a *AddressIPVersionEnum with the value of string s
// If the empty string is provided, nil is returned.
func AddressIPVersionEnumRef(s string) *AddressIPVersionEnum {
	if s == "" {
		return nil
	}

	v := AddressIPVersionEnum(s)
	return &v
}

func (v AddressIPVersionEnum) Validate() error {
	for _, s := range []string{"IPV4", "IPV6"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AddressIPVersionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AddressAddressTypeEnum.
type AddressAddressTypeEnum string

// AddressAddressTypeEnumRef returns a *AddressAddressTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AddressAddressTypeEnumRef(s string) *AddressAddressTypeEnum {
	if s == "" {
		return nil
	}

	v := AddressAddressTypeEnum(s)
	return &v
}

func (v AddressAddressTypeEnum) Validate() error {
	for _, s := range []string{"INTERNAL", "EXTERNAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AddressAddressTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AddressPurposeEnum.
type AddressPurposeEnum string

// AddressPurposeEnumRef returns a *AddressPurposeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AddressPurposeEnumRef(s string) *AddressPurposeEnum {
	if s == "" {
		return nil
	}

	v := AddressPurposeEnum(s)
	return &v
}

func (v AddressPurposeEnum) Validate() error {
	for _, s := range []string{"GCE_ENDPOINT", "DNS_RESOLVER", "VPC_PEERING", "NAT_AUTO"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AddressPurposeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Address) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Address",
		Version: "beta",
	}
}

const AddressMaxPage = -1

type AddressList struct {
	Items []*Address

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *AddressList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AddressList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAddress(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAddress(ctx context.Context, project, location string) (*AddressList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAddressWithMaxResults(ctx, project, location, AddressMaxPage)

}

func (c *Client) ListAddressWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*AddressList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAddress(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AddressList{
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
func (r *Address) URLNormalized() *Address {
	normalized := dcl.Copy(*r).(Address)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Address = dcl.SelfLinkToName(r.Address)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Subnetwork = dcl.SelfLinkToName(r.Subnetwork)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.LabelFingerprint = dcl.SelfLinkToName(r.LabelFingerprint)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetAddress(ctx context.Context, r *Address) (*Address, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAddressRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAddress(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name
	if dcl.IsZeroValue(result.AddressType) {
		result.AddressType = AddressAddressTypeEnumRef("EXTERNAL")
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAddressNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAddress(ctx context.Context, r *Address) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Address resource is nil")
	}
	c.Config.Logger.Info("Deleting Address...")
	deleteOp := deleteAddressOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAddress deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAddress(ctx context.Context, project, location string, filter func(*Address) bool) error {
	listObj, err := c.ListAddress(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllAddress(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAddress(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAddress(ctx context.Context, rawDesired *Address, opts ...dcl.ApplyOption) (*Address, error) {

	var resultNewState *Address
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAddressHelper(c, ctx, rawDesired, opts...)
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

func applyAddressHelper(c *Client, ctx context.Context, rawDesired *Address, opts ...dcl.ApplyOption) (*Address, error) {
	c.Config.Logger.Info("Beginning ApplyAddress...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.addressDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToAddressOp(opStrings, fieldDiffs, opts)
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
	var ops []addressApiOperation
	if create {
		ops = append(ops, &createAddressOperation{})
	} else if recreate {
		ops = append(ops, &deleteAddressOperation{})
		ops = append(ops, &createAddressOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAddressDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetAddress(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAddressOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAddress(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAddressNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAddressNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAddressDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAddress(c, newDesired, newState)
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
