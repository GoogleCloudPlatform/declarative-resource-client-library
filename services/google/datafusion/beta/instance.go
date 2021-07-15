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

type Instance struct {
	Name                        *string                    `json:"name"`
	Description                 *string                    `json:"description"`
	Type                        *InstanceTypeEnum          `json:"type"`
	EnableStackdriverLogging    *bool                      `json:"enableStackdriverLogging"`
	EnableStackdriverMonitoring *bool                      `json:"enableStackdriverMonitoring"`
	PrivateInstance             *bool                      `json:"privateInstance"`
	NetworkConfig               *InstanceNetworkConfig     `json:"networkConfig"`
	Labels                      map[string]string          `json:"labels"`
	Options                     map[string]string          `json:"options"`
	CreateTime                  *string                    `json:"createTime"`
	UpdateTime                  *string                    `json:"updateTime"`
	State                       *InstanceStateEnum         `json:"state"`
	StateMessage                *string                    `json:"stateMessage"`
	ServiceEndpoint             *string                    `json:"serviceEndpoint"`
	Zone                        *string                    `json:"zone"`
	Version                     *string                    `json:"version"`
	DisplayName                 *string                    `json:"displayName"`
	AvailableVersion            []InstanceAvailableVersion `json:"availableVersion"`
	ApiEndpoint                 *string                    `json:"apiEndpoint"`
	GcsBucket                   *string                    `json:"gcsBucket"`
	P4ServiceAccount            *string                    `json:"p4ServiceAccount"`
	TenantProjectId             *string                    `json:"tenantProjectId"`
	DataprocServiceAccount      *string                    `json:"dataprocServiceAccount"`
	Project                     *string                    `json:"project"`
	Location                    *string                    `json:"location"`
}

func (r *Instance) String() string {
	return dcl.SprintResource(r)
}

// The enum InstanceTypeEnum.
type InstanceTypeEnum string

// InstanceTypeEnumRef returns a *InstanceTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceTypeEnumRef(s string) *InstanceTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceTypeEnum(s)
	return &v
}

func (v InstanceTypeEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "BASIC", "ENTERPRISE", "DEVELOPER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceStateEnum.
type InstanceStateEnum string

// InstanceStateEnumRef returns a *InstanceStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceStateEnumRef(s string) *InstanceStateEnum {
	if s == "" {
		return nil
	}

	v := InstanceStateEnum(s)
	return &v
}

func (v InstanceStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "ENABLED", "DISABLED", "UNKNOWN"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InstanceNetworkConfig struct {
	empty        bool    `json:"-"`
	Network      *string `json:"network"`
	IPAllocation *string `json:"ipAllocation"`
}

type jsonInstanceNetworkConfig InstanceNetworkConfig

func (r *InstanceNetworkConfig) UnmarshalJSON(data []byte) error {
	var res jsonInstanceNetworkConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceNetworkConfig
	} else {

		r.Network = res.Network

		r.IPAllocation = res.IPAllocation

	}
	return nil
}

// This object is used to assert a desired state where this InstanceNetworkConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNetworkConfig *InstanceNetworkConfig = &InstanceNetworkConfig{empty: true}

func (r *InstanceNetworkConfig) Empty() bool {
	return r.empty
}

func (r *InstanceNetworkConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNetworkConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceAvailableVersion struct {
	empty             bool     `json:"-"`
	VersionNumber     *string  `json:"versionNumber"`
	DefaultVersion    *bool    `json:"defaultVersion"`
	AvailableFeatures []string `json:"availableFeatures"`
}

type jsonInstanceAvailableVersion InstanceAvailableVersion

func (r *InstanceAvailableVersion) UnmarshalJSON(data []byte) error {
	var res jsonInstanceAvailableVersion
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceAvailableVersion
	} else {

		r.VersionNumber = res.VersionNumber

		r.DefaultVersion = res.DefaultVersion

		r.AvailableFeatures = res.AvailableFeatures

	}
	return nil
}

// This object is used to assert a desired state where this InstanceAvailableVersion is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceAvailableVersion *InstanceAvailableVersion = &InstanceAvailableVersion{empty: true}

func (r *InstanceAvailableVersion) Empty() bool {
	return r.empty
}

func (r *InstanceAvailableVersion) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceAvailableVersion) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Instance) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "data_fusion",
		Type:    "Instance",
		Version: "beta",
	}
}

const InstanceMaxPage = -1

type InstanceList struct {
	Items []*Instance

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *InstanceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InstanceList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInstance(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInstance(ctx context.Context, project, location string) (*InstanceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListInstanceWithMaxResults(ctx, project, location, InstanceMaxPage)

}

func (c *Client) ListInstanceWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*InstanceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listInstance(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InstanceList{
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
func (r *Instance) URLNormalized() *Instance {
	normalized := dcl.Copy(*r).(Instance)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.StateMessage = dcl.SelfLinkToName(r.StateMessage)
	normalized.ServiceEndpoint = dcl.SelfLinkToName(r.ServiceEndpoint)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.Version = dcl.SelfLinkToName(r.Version)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.ApiEndpoint = dcl.SelfLinkToName(r.ApiEndpoint)
	normalized.GcsBucket = dcl.SelfLinkToName(r.GcsBucket)
	normalized.P4ServiceAccount = dcl.SelfLinkToName(r.P4ServiceAccount)
	normalized.TenantProjectId = dcl.SelfLinkToName(r.TenantProjectId)
	normalized.DataprocServiceAccount = dcl.SelfLinkToName(r.DataprocServiceAccount)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetInstance(ctx context.Context, r *Instance) (*Instance, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getInstanceRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalInstance(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeInstanceNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteInstance(ctx context.Context, r *Instance) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Instance resource is nil")
	}
	c.Config.Logger.Info("Deleting Instance...")
	deleteOp := deleteInstanceOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInstance deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInstance(ctx context.Context, project, location string, filter func(*Instance) bool) error {
	listObj, err := c.ListInstance(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllInstance(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllInstance(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyInstance(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (*Instance, error) {

	var resultNewState *Instance
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyInstanceHelper(c, ctx, rawDesired, opts...)
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

func applyInstanceHelper(c *Client, ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (*Instance, error) {
	c.Config.Logger.Info("Beginning ApplyInstance...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.instanceDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToInstanceDiffs(c.Config, fieldDiffs, opts)
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
	var ops []instanceApiOperation
	if create {
		ops = append(ops, &createInstanceOperation{})
	} else if recreate {
		ops = append(ops, &deleteInstanceOperation{})
		ops = append(ops, &createInstanceOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeInstanceDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetInstance(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createInstanceOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapInstance(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeInstanceNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeInstanceNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeInstanceDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffInstance(c, newDesired, newState)
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
