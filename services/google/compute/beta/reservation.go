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
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Reservation struct {
	Id                          *int64                          `json:"id"`
	SelfLink                    *string                         `json:"selfLink"`
	Zone                        *string                         `json:"zone"`
	Description                 *string                         `json:"description"`
	Name                        *string                         `json:"name"`
	SpecificReservation         *ReservationSpecificReservation `json:"specificReservation"`
	Commitment                  *string                         `json:"commitment"`
	SpecificReservationRequired *bool                           `json:"specificReservationRequired"`
	Status                      *ReservationStatusEnum          `json:"status"`
	Project                     *string                         `json:"project"`
}

func (r *Reservation) String() string {
	return dcl.SprintResource(r)
}

// The enum ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum.
type ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum string

// ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumRef returns a *ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum with the value of string s
// If the empty string is provided, nil is returned.
func ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumRef(s string) *ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum {
	if s == "" {
		return nil
	}

	v := ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(s)
	return &v
}

func (v ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum) Validate() error {
	for _, s := range []string{"SCSI", "NVME", "NVDIMM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ReservationStatusEnum.
type ReservationStatusEnum string

// ReservationStatusEnumRef returns a *ReservationStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func ReservationStatusEnumRef(s string) *ReservationStatusEnum {
	if s == "" {
		return nil
	}

	v := ReservationStatusEnum(s)
	return &v
}

func (v ReservationStatusEnum) Validate() error {
	for _, s := range []string{"PENDING", "RUNNING", "DONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ReservationStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ReservationSpecificReservation struct {
	empty              bool                                              `json:"-"`
	InstanceProperties *ReservationSpecificReservationInstanceProperties `json:"instanceProperties"`
	Count              *int64                                            `json:"count"`
	InUseCount         *int64                                            `json:"inUseCount"`
}

// This object is used to assert a desired state where this ReservationSpecificReservation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyReservationSpecificReservation *ReservationSpecificReservation = &ReservationSpecificReservation{empty: true}

func (r *ReservationSpecificReservation) String() string {
	return dcl.SprintResource(r)
}

func (r *ReservationSpecificReservation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ReservationSpecificReservationInstanceProperties struct {
	empty             bool                                                                `json:"-"`
	MachineType       *string                                                             `json:"machineType"`
	GuestAccelerators []ReservationSpecificReservationInstancePropertiesGuestAccelerators `json:"guestAccelerators"`
	MinCpuPlatform    *string                                                             `json:"minCpuPlatform"`
	LocalSsds         []ReservationSpecificReservationInstancePropertiesLocalSsds         `json:"localSsds"`
}

// This object is used to assert a desired state where this ReservationSpecificReservationInstanceProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyReservationSpecificReservationInstanceProperties *ReservationSpecificReservationInstanceProperties = &ReservationSpecificReservationInstanceProperties{empty: true}

func (r *ReservationSpecificReservationInstanceProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *ReservationSpecificReservationInstanceProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ReservationSpecificReservationInstancePropertiesGuestAccelerators struct {
	empty            bool    `json:"-"`
	AcceleratorType  *string `json:"acceleratorType"`
	AcceleratorCount *int64  `json:"acceleratorCount"`
}

// This object is used to assert a desired state where this ReservationSpecificReservationInstancePropertiesGuestAccelerators is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyReservationSpecificReservationInstancePropertiesGuestAccelerators *ReservationSpecificReservationInstancePropertiesGuestAccelerators = &ReservationSpecificReservationInstancePropertiesGuestAccelerators{empty: true}

func (r *ReservationSpecificReservationInstancePropertiesGuestAccelerators) String() string {
	return dcl.SprintResource(r)
}

func (r *ReservationSpecificReservationInstancePropertiesGuestAccelerators) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ReservationSpecificReservationInstancePropertiesLocalSsds struct {
	empty      bool                                                                    `json:"-"`
	DiskSizeGb *int64                                                                  `json:"diskSizeGb"`
	Interface  *ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum `json:"interface"`
}

// This object is used to assert a desired state where this ReservationSpecificReservationInstancePropertiesLocalSsds is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyReservationSpecificReservationInstancePropertiesLocalSsds *ReservationSpecificReservationInstancePropertiesLocalSsds = &ReservationSpecificReservationInstancePropertiesLocalSsds{empty: true}

func (r *ReservationSpecificReservationInstancePropertiesLocalSsds) String() string {
	return dcl.SprintResource(r)
}

func (r *ReservationSpecificReservationInstancePropertiesLocalSsds) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Reservation) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Reservation",
		Version: "beta",
	}
}

const ReservationMaxPage = -1

type ReservationList struct {
	Items []*Reservation

	nextToken string

	pageSize int32

	project string

	zone string
}

func (l *ReservationList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ReservationList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listReservation(ctx, l.project, l.zone, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListReservation(ctx context.Context, project, zone string) (*ReservationList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListReservationWithMaxResults(ctx, project, zone, ReservationMaxPage)

}

func (c *Client) ListReservationWithMaxResults(ctx context.Context, project, zone string, pageSize int32) (*ReservationList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listReservation(ctx, project, zone, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ReservationList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		zone: zone,
	}, nil
}

func (c *Client) GetReservation(ctx context.Context, r *Reservation) (*Reservation, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getReservationRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalReservation(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Zone = r.Zone
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeReservationNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteReservation(ctx context.Context, r *Reservation) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Reservation resource is nil")
	}
	c.Config.Logger.Info("Deleting Reservation...")
	deleteOp := deleteReservationOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllReservation deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllReservation(ctx context.Context, project, zone string, filter func(*Reservation) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListReservation(ctx, project, zone)
	if err != nil {
		return err
	}

	err = c.deleteAllReservation(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllReservation(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyReservation(ctx context.Context, rawDesired *Reservation, opts ...dcl.ApplyOption) (*Reservation, error) {
	c.Config.Logger.Info("Beginning ApplyReservation...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.reservationDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []reservationApiOperation
	if create {
		ops = append(ops, &createReservationOperation{})
	} else if recreate {

		ops = append(ops, &deleteReservationOperation{})

		ops = append(ops, &createReservationOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeReservationDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetReservation(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createReservationOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapReservation(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeReservationNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeReservationNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeReservationDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffReservation(c, newDesired, newState)
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
