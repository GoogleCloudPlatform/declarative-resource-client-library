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

type ManagedSslCertificate struct {
	Id                      *int64                         `json:"id"`
	Name                    *string                        `json:"name"`
	Description             *string                        `json:"description"`
	SelfLink                *string                        `json:"selfLink"`
	Managed                 *ManagedSslCertificateManaged  `json:"managed"`
	Type                    *ManagedSslCertificateTypeEnum `json:"type"`
	SubjectAlternativeNames []string                       `json:"subjectAlternativeNames"`
	ExpireTime              *string                        `json:"expireTime"`
	Project                 *string                        `json:"project"`
}

func (r *ManagedSslCertificate) String() string {
	return dcl.SprintResource(r)
}

// The enum ManagedSslCertificateManagedStatusEnum.
type ManagedSslCertificateManagedStatusEnum string

// ManagedSslCertificateManagedStatusEnumRef returns a *ManagedSslCertificateManagedStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func ManagedSslCertificateManagedStatusEnumRef(s string) *ManagedSslCertificateManagedStatusEnum {
	if s == "" {
		return nil
	}

	v := ManagedSslCertificateManagedStatusEnum(s)
	return &v
}

func (v ManagedSslCertificateManagedStatusEnum) Validate() error {
	for _, s := range []string{"PENDING", "RUNNING", "DONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ManagedSslCertificateManagedStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ManagedSslCertificateTypeEnum.
type ManagedSslCertificateTypeEnum string

// ManagedSslCertificateTypeEnumRef returns a *ManagedSslCertificateTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ManagedSslCertificateTypeEnumRef(s string) *ManagedSslCertificateTypeEnum {
	if s == "" {
		return nil
	}

	v := ManagedSslCertificateTypeEnum(s)
	return &v
}

func (v ManagedSslCertificateTypeEnum) Validate() error {
	for _, s := range []string{"MANAGED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ManagedSslCertificateTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ManagedSslCertificateManaged struct {
	empty        bool                                    `json:"-"`
	Domains      []string                                `json:"domains"`
	Status       *ManagedSslCertificateManagedStatusEnum `json:"status"`
	DomainStatus map[string]string                       `json:"domainStatus"`
}

type jsonManagedSslCertificateManaged ManagedSslCertificateManaged

func (r *ManagedSslCertificateManaged) UnmarshalJSON(data []byte) error {
	var res jsonManagedSslCertificateManaged
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyManagedSslCertificateManaged
	} else {

		r.Domains = res.Domains

		r.Status = res.Status

		r.DomainStatus = res.DomainStatus

	}
	return nil
}

// This object is used to assert a desired state where this ManagedSslCertificateManaged is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyManagedSslCertificateManaged *ManagedSslCertificateManaged = &ManagedSslCertificateManaged{empty: true}

func (r *ManagedSslCertificateManaged) Empty() bool {
	return r.empty
}

func (r *ManagedSslCertificateManaged) String() string {
	return dcl.SprintResource(r)
}

func (r *ManagedSslCertificateManaged) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ManagedSslCertificate) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "ManagedSslCertificate",
		Version: "compute",
	}
}

const ManagedSslCertificateMaxPage = -1

type ManagedSslCertificateList struct {
	Items []*ManagedSslCertificate

	nextToken string

	pageSize int32

	project string
}

func (l *ManagedSslCertificateList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ManagedSslCertificateList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listManagedSslCertificate(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListManagedSslCertificate(ctx context.Context, project string) (*ManagedSslCertificateList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListManagedSslCertificateWithMaxResults(ctx, project, ManagedSslCertificateMaxPage)

}

func (c *Client) ListManagedSslCertificateWithMaxResults(ctx context.Context, project string, pageSize int32) (*ManagedSslCertificateList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listManagedSslCertificate(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ManagedSslCertificateList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetManagedSslCertificate(ctx context.Context, r *ManagedSslCertificate) (*ManagedSslCertificate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getManagedSslCertificateRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalManagedSslCertificate(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name
	if dcl.IsZeroValue(result.Type) {
		result.Type = ManagedSslCertificateTypeEnumRef("MANAGED")
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeManagedSslCertificateNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteManagedSslCertificate(ctx context.Context, r *ManagedSslCertificate) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("ManagedSslCertificate resource is nil")
	}
	c.Config.Logger.Info("Deleting ManagedSslCertificate...")
	deleteOp := deleteManagedSslCertificateOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllManagedSslCertificate deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllManagedSslCertificate(ctx context.Context, project string, filter func(*ManagedSslCertificate) bool) error {
	listObj, err := c.ListManagedSslCertificate(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllManagedSslCertificate(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllManagedSslCertificate(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyManagedSslCertificate(ctx context.Context, rawDesired *ManagedSslCertificate, opts ...dcl.ApplyOption) (*ManagedSslCertificate, error) {

	var resultNewState *ManagedSslCertificate
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyManagedSslCertificateHelper(c, ctx, rawDesired, opts...)
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

func applyManagedSslCertificateHelper(c *Client, ctx context.Context, rawDesired *ManagedSslCertificate, opts ...dcl.ApplyOption) (*ManagedSslCertificate, error) {
	c.Config.Logger.Info("Beginning ApplyManagedSslCertificate...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.managedSslCertificateDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []managedSslCertificateApiOperation
	if create {
		ops = append(ops, &createManagedSslCertificateOperation{})
	} else if recreate {

		ops = append(ops, &deleteManagedSslCertificateOperation{})

		ops = append(ops, &createManagedSslCertificateOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeManagedSslCertificateDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetManagedSslCertificate(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createManagedSslCertificateOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapManagedSslCertificate(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeManagedSslCertificateNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeManagedSslCertificateNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeManagedSslCertificateDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffManagedSslCertificate(c, newDesired, newState)
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
