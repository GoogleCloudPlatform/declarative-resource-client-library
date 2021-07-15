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
package appengine

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type DomainMapping struct {
	SelfLink        *string                        `json:"selfLink"`
	Name            *string                        `json:"name"`
	SslSettings     *DomainMappingSslSettings      `json:"sslSettings"`
	ResourceRecords []DomainMappingResourceRecords `json:"resourceRecords"`
	App             *string                        `json:"app"`
}

func (r *DomainMapping) String() string {
	return dcl.SprintResource(r)
}

// The enum DomainMappingSslSettingsSslManagementTypeEnum.
type DomainMappingSslSettingsSslManagementTypeEnum string

// DomainMappingSslSettingsSslManagementTypeEnumRef returns a *DomainMappingSslSettingsSslManagementTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func DomainMappingSslSettingsSslManagementTypeEnumRef(s string) *DomainMappingSslSettingsSslManagementTypeEnum {
	if s == "" {
		return nil
	}

	v := DomainMappingSslSettingsSslManagementTypeEnum(s)
	return &v
}

func (v DomainMappingSslSettingsSslManagementTypeEnum) Validate() error {
	for _, s := range []string{"AUTOMATIC", "MANUAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DomainMappingSslSettingsSslManagementTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DomainMappingResourceRecordsTypeEnum.
type DomainMappingResourceRecordsTypeEnum string

// DomainMappingResourceRecordsTypeEnumRef returns a *DomainMappingResourceRecordsTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func DomainMappingResourceRecordsTypeEnumRef(s string) *DomainMappingResourceRecordsTypeEnum {
	if s == "" {
		return nil
	}

	v := DomainMappingResourceRecordsTypeEnum(s)
	return &v
}

func (v DomainMappingResourceRecordsTypeEnum) Validate() error {
	for _, s := range []string{"DATABASE_TYPE_UNSPECIFIED", "CLOUD_DATASTORE", "CLOUD_FIRESTORE", "CLOUD_DATASTORE_COMPATIBILITY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DomainMappingResourceRecordsTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type DomainMappingSslSettings struct {
	empty                       bool                                           `json:"-"`
	CertificateId               *string                                        `json:"certificateId"`
	SslManagementType           *DomainMappingSslSettingsSslManagementTypeEnum `json:"sslManagementType"`
	PendingManagedCertificateId *string                                        `json:"pendingManagedCertificateId"`
}

type jsonDomainMappingSslSettings DomainMappingSslSettings

func (r *DomainMappingSslSettings) UnmarshalJSON(data []byte) error {
	var res jsonDomainMappingSslSettings
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDomainMappingSslSettings
	} else {

		r.CertificateId = res.CertificateId

		r.SslManagementType = res.SslManagementType

		r.PendingManagedCertificateId = res.PendingManagedCertificateId

	}
	return nil
}

// This object is used to assert a desired state where this DomainMappingSslSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDomainMappingSslSettings *DomainMappingSslSettings = &DomainMappingSslSettings{empty: true}

func (r *DomainMappingSslSettings) Empty() bool {
	return r.empty
}

func (r *DomainMappingSslSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *DomainMappingSslSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DomainMappingResourceRecords struct {
	empty  bool                                  `json:"-"`
	Name   *string                               `json:"name"`
	Rrdata *string                               `json:"rrdata"`
	Type   *DomainMappingResourceRecordsTypeEnum `json:"type"`
}

type jsonDomainMappingResourceRecords DomainMappingResourceRecords

func (r *DomainMappingResourceRecords) UnmarshalJSON(data []byte) error {
	var res jsonDomainMappingResourceRecords
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDomainMappingResourceRecords
	} else {

		r.Name = res.Name

		r.Rrdata = res.Rrdata

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this DomainMappingResourceRecords is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDomainMappingResourceRecords *DomainMappingResourceRecords = &DomainMappingResourceRecords{empty: true}

func (r *DomainMappingResourceRecords) Empty() bool {
	return r.empty
}

func (r *DomainMappingResourceRecords) String() string {
	return dcl.SprintResource(r)
}

func (r *DomainMappingResourceRecords) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *DomainMapping) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "app_engine",
		Type:    "DomainMapping",
		Version: "appengine",
	}
}

const DomainMappingMaxPage = -1

type DomainMappingList struct {
	Items []*DomainMapping

	nextToken string

	pageSize int32

	app string
}

func (l *DomainMappingList) HasNext() bool {
	return l.nextToken != ""
}

func (l *DomainMappingList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listDomainMapping(ctx, l.app, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListDomainMapping(ctx context.Context, app string) (*DomainMappingList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListDomainMappingWithMaxResults(ctx, app, DomainMappingMaxPage)

}

func (c *Client) ListDomainMappingWithMaxResults(ctx context.Context, app string, pageSize int32) (*DomainMappingList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listDomainMapping(ctx, app, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &DomainMappingList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		app: app,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *DomainMapping) URLNormalized() *DomainMapping {
	normalized := dcl.Copy(*r).(DomainMapping)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.App = dcl.SelfLinkToName(r.App)
	return &normalized
}

func (c *Client) GetDomainMapping(ctx context.Context, r *DomainMapping) (*DomainMapping, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getDomainMappingRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalDomainMapping(b, c)
	if err != nil {
		return nil, err
	}
	result.App = r.App
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeDomainMappingNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteDomainMapping(ctx context.Context, r *DomainMapping) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("DomainMapping resource is nil")
	}
	c.Config.Logger.Info("Deleting DomainMapping...")
	deleteOp := deleteDomainMappingOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllDomainMapping deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllDomainMapping(ctx context.Context, app string, filter func(*DomainMapping) bool) error {
	listObj, err := c.ListDomainMapping(ctx, app)
	if err != nil {
		return err
	}

	err = c.deleteAllDomainMapping(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllDomainMapping(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyDomainMapping(ctx context.Context, rawDesired *DomainMapping, opts ...dcl.ApplyOption) (*DomainMapping, error) {
	var resultNewState *DomainMapping
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyDomainMappingHelper(c, ctx, rawDesired, opts...)
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

func applyDomainMappingHelper(c *Client, ctx context.Context, rawDesired *DomainMapping, opts ...dcl.ApplyOption) (*DomainMapping, error) {
	c.Config.Logger.Info("Beginning ApplyDomainMapping...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.domainMappingDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToDomainMappingDiffs(c.Config, fieldDiffs, opts)
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
	var ops []domainMappingApiOperation
	if create {
		ops = append(ops, &createDomainMappingOperation{})
	} else if recreate {
		ops = append(ops, &deleteDomainMappingOperation{})
		ops = append(ops, &createDomainMappingOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeDomainMappingDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetDomainMapping(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createDomainMappingOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapDomainMapping(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeDomainMappingNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeDomainMappingNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeDomainMappingDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffDomainMapping(c, newDesired, newState)
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
