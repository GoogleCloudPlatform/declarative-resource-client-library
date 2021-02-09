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
package accesscontextmanager

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type ServicePerimeter struct {
	Title                 *string                            `json:"title"`
	Description           *string                            `json:"description"`
	CreateTime            *string                            `json:"createTime"`
	UpdateTime            *string                            `json:"updateTime"`
	PerimeterType         *ServicePerimeterPerimeterTypeEnum `json:"perimeterType"`
	Status                *ServicePerimeterStatus            `json:"status"`
	Policy                *string                            `json:"policy"`
	Name                  *string                            `json:"name"`
	UseExplicitDryRunSpec *bool                              `json:"useExplicitDryRunSpec"`
	Spec                  *ServicePerimeterSpec              `json:"spec"`
}

func (r *ServicePerimeter) String() string {
	return dcl.SprintResource(r)
}

// The enum ServicePerimeterPerimeterTypeEnum.
type ServicePerimeterPerimeterTypeEnum string

// ServicePerimeterPerimeterTypeEnumRef returns a *ServicePerimeterPerimeterTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServicePerimeterPerimeterTypeEnumRef(s string) *ServicePerimeterPerimeterTypeEnum {
	if s == "" {
		return nil
	}

	v := ServicePerimeterPerimeterTypeEnum(s)
	return &v
}

func (v ServicePerimeterPerimeterTypeEnum) Validate() error {
	for _, s := range []string{"PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ServicePerimeterPerimeterTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ServicePerimeterStatus struct {
	empty                 bool                                         `json:"-"`
	Resources             []string                                     `json:"resources"`
	AccessLevels          []string                                     `json:"accessLevels"`
	RestrictedServices    []string                                     `json:"restrictedServices"`
	VPCAccessibleServices *ServicePerimeterStatusVPCAccessibleServices `json:"vpcAccessibleServices"`
}

// This object is used to assert a desired state where this ServicePerimeterStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServicePerimeterStatus *ServicePerimeterStatus = &ServicePerimeterStatus{empty: true}

func (r *ServicePerimeterStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *ServicePerimeterStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServicePerimeterStatusVPCAccessibleServices struct {
	empty             bool     `json:"-"`
	EnableRestriction *bool    `json:"enableRestriction"`
	AllowedServices   []string `json:"allowedServices"`
}

// This object is used to assert a desired state where this ServicePerimeterStatusVPCAccessibleServices is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServicePerimeterStatusVPCAccessibleServices *ServicePerimeterStatusVPCAccessibleServices = &ServicePerimeterStatusVPCAccessibleServices{empty: true}

func (r *ServicePerimeterStatusVPCAccessibleServices) String() string {
	return dcl.SprintResource(r)
}

func (r *ServicePerimeterStatusVPCAccessibleServices) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServicePerimeterSpec struct {
	empty                 bool                                       `json:"-"`
	Resources             []string                                   `json:"resources"`
	AccessLevels          []string                                   `json:"accessLevels"`
	RestrictedServices    []string                                   `json:"restrictedServices"`
	VPCAccessibleServices *ServicePerimeterSpecVPCAccessibleServices `json:"vpcAccessibleServices"`
}

// This object is used to assert a desired state where this ServicePerimeterSpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServicePerimeterSpec *ServicePerimeterSpec = &ServicePerimeterSpec{empty: true}

func (r *ServicePerimeterSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *ServicePerimeterSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServicePerimeterSpecVPCAccessibleServices struct {
	empty             bool     `json:"-"`
	EnableRestriction *bool    `json:"enableRestriction"`
	AllowedServices   []string `json:"allowedServices"`
}

// This object is used to assert a desired state where this ServicePerimeterSpecVPCAccessibleServices is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServicePerimeterSpecVPCAccessibleServices *ServicePerimeterSpecVPCAccessibleServices = &ServicePerimeterSpecVPCAccessibleServices{empty: true}

func (r *ServicePerimeterSpecVPCAccessibleServices) String() string {
	return dcl.SprintResource(r)
}

func (r *ServicePerimeterSpecVPCAccessibleServices) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ServicePerimeter) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "access_context_manager",
		Type:    "ServicePerimeter",
		Version: "accesscontextmanager",
	}
}

const ServicePerimeterMaxPage = -1

type ServicePerimeterList struct {
	Items []*ServicePerimeter

	nextToken string

	pageSize int32

	policy string
}

func (l *ServicePerimeterList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ServicePerimeterList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listServicePerimeter(ctx, l.policy, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListServicePerimeter(ctx context.Context, policy string) (*ServicePerimeterList, error) {

	return c.ListServicePerimeterWithMaxResults(ctx, policy, ServicePerimeterMaxPage)

}

func (c *Client) ListServicePerimeterWithMaxResults(ctx context.Context, policy string, pageSize int32) (*ServicePerimeterList, error) {
	items, token, err := c.listServicePerimeter(ctx, policy, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ServicePerimeterList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		policy: policy,
	}, nil
}

func (c *Client) GetServicePerimeter(ctx context.Context, r *ServicePerimeter) (*ServicePerimeter, error) {
	b, err := c.getServicePerimeterRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalServicePerimeter(b, c)
	if err != nil {
		return nil, err
	}
	result.Policy = r.Policy
	result.Name = r.Name
	if dcl.IsZeroValue(result.PerimeterType) {
		result.PerimeterType = ServicePerimeterPerimeterTypeEnumRef("PERIMETER_TYPE_REGULAR")
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeServicePerimeterNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteServicePerimeter(ctx context.Context, r *ServicePerimeter) error {
	if r == nil {
		return fmt.Errorf("ServicePerimeter resource is nil")
	}
	c.Config.Logger.Info("Deleting ServicePerimeter...")
	deleteOp := deleteServicePerimeterOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllServicePerimeter deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllServicePerimeter(ctx context.Context, policy string, filter func(*ServicePerimeter) bool) error {
	listObj, err := c.ListServicePerimeter(ctx, policy)
	if err != nil {
		return err
	}

	err = c.deleteAllServicePerimeter(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllServicePerimeter(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyServicePerimeter(ctx context.Context, rawDesired *ServicePerimeter, opts ...dcl.ApplyOption) (*ServicePerimeter, error) {
	c.Config.Logger.Info("Beginning ApplyServicePerimeter...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.servicePerimeterDiffsForRawDesired(ctx, rawDesired, opts...)
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
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []servicePerimeterApiOperation
	if create {
		ops = append(ops, &createServicePerimeterOperation{})
	} else if recreate {

		ops = append(ops, &deleteServicePerimeterOperation{})

		ops = append(ops, &createServicePerimeterOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeServicePerimeterDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetServicePerimeter(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeServicePerimeterNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeServicePerimeterDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffServicePerimeter(c, newDesired, newState)
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
