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
package alpha

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type KrmApiHost struct {
	Name               *string                     `json:"name"`
	Labels             map[string]string           `json:"labels"`
	BundlesConfig      *KrmApiHostBundlesConfig    `json:"bundlesConfig"`
	UsePrivateEndpoint *bool                       `json:"usePrivateEndpoint"`
	GkeResourceLink    *string                     `json:"gkeResourceLink"`
	State              *KrmApiHostStateEnum        `json:"state"`
	ManagementConfig   *KrmApiHostManagementConfig `json:"managementConfig"`
	Project            *string                     `json:"project"`
	Location           *string                     `json:"location"`
}

func (r *KrmApiHost) String() string {
	return dcl.SprintResource(r)
}

// The enum KrmApiHostStateEnum.
type KrmApiHostStateEnum string

// KrmApiHostStateEnumRef returns a *KrmApiHostStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func KrmApiHostStateEnumRef(s string) *KrmApiHostStateEnum {
	if s == "" {
		return nil
	}

	v := KrmApiHostStateEnum(s)
	return &v
}

func (v KrmApiHostStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "CREATING", "RUNNING", "DELETING", "SUSPENDED", "READ_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "KrmApiHostStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type KrmApiHostBundlesConfig struct {
	empty                  bool                                           `json:"-"`
	ConfigControllerConfig *KrmApiHostBundlesConfigConfigControllerConfig `json:"configControllerConfig"`
}

type jsonKrmApiHostBundlesConfig KrmApiHostBundlesConfig

func (r *KrmApiHostBundlesConfig) UnmarshalJSON(data []byte) error {
	var res jsonKrmApiHostBundlesConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyKrmApiHostBundlesConfig
	} else {

		r.ConfigControllerConfig = res.ConfigControllerConfig

	}
	return nil
}

// This object is used to assert a desired state where this KrmApiHostBundlesConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyKrmApiHostBundlesConfig *KrmApiHostBundlesConfig = &KrmApiHostBundlesConfig{empty: true}

func (r *KrmApiHostBundlesConfig) Empty() bool {
	return r.empty
}

func (r *KrmApiHostBundlesConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *KrmApiHostBundlesConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type KrmApiHostBundlesConfigConfigControllerConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

type jsonKrmApiHostBundlesConfigConfigControllerConfig KrmApiHostBundlesConfigConfigControllerConfig

func (r *KrmApiHostBundlesConfigConfigControllerConfig) UnmarshalJSON(data []byte) error {
	var res jsonKrmApiHostBundlesConfigConfigControllerConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyKrmApiHostBundlesConfigConfigControllerConfig
	} else {

		r.Enabled = res.Enabled

	}
	return nil
}

// This object is used to assert a desired state where this KrmApiHostBundlesConfigConfigControllerConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyKrmApiHostBundlesConfigConfigControllerConfig *KrmApiHostBundlesConfigConfigControllerConfig = &KrmApiHostBundlesConfigConfigControllerConfig{empty: true}

func (r *KrmApiHostBundlesConfigConfigControllerConfig) Empty() bool {
	return r.empty
}

func (r *KrmApiHostBundlesConfigConfigControllerConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *KrmApiHostBundlesConfigConfigControllerConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type KrmApiHostManagementConfig struct {
	empty                    bool                                                `json:"-"`
	StandardManagementConfig *KrmApiHostManagementConfigStandardManagementConfig `json:"standardManagementConfig"`
}

type jsonKrmApiHostManagementConfig KrmApiHostManagementConfig

func (r *KrmApiHostManagementConfig) UnmarshalJSON(data []byte) error {
	var res jsonKrmApiHostManagementConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyKrmApiHostManagementConfig
	} else {

		r.StandardManagementConfig = res.StandardManagementConfig

	}
	return nil
}

// This object is used to assert a desired state where this KrmApiHostManagementConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyKrmApiHostManagementConfig *KrmApiHostManagementConfig = &KrmApiHostManagementConfig{empty: true}

func (r *KrmApiHostManagementConfig) Empty() bool {
	return r.empty
}

func (r *KrmApiHostManagementConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *KrmApiHostManagementConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type KrmApiHostManagementConfigStandardManagementConfig struct {
	empty               bool    `json:"-"`
	Network             *string `json:"network"`
	MasterIPv4CidrBlock *string `json:"masterIPv4CidrBlock"`
	ManBlock            *string `json:"manBlock"`
	ClusterCidrBlock    *string `json:"clusterCidrBlock"`
	ServicesCidrBlock   *string `json:"servicesCidrBlock"`
	ClusterNamedRange   *string `json:"clusterNamedRange"`
	ServicesNamedRange  *string `json:"servicesNamedRange"`
}

type jsonKrmApiHostManagementConfigStandardManagementConfig KrmApiHostManagementConfigStandardManagementConfig

func (r *KrmApiHostManagementConfigStandardManagementConfig) UnmarshalJSON(data []byte) error {
	var res jsonKrmApiHostManagementConfigStandardManagementConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyKrmApiHostManagementConfigStandardManagementConfig
	} else {

		r.Network = res.Network

		r.MasterIPv4CidrBlock = res.MasterIPv4CidrBlock

		r.ManBlock = res.ManBlock

		r.ClusterCidrBlock = res.ClusterCidrBlock

		r.ServicesCidrBlock = res.ServicesCidrBlock

		r.ClusterNamedRange = res.ClusterNamedRange

		r.ServicesNamedRange = res.ServicesNamedRange

	}
	return nil
}

// This object is used to assert a desired state where this KrmApiHostManagementConfigStandardManagementConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyKrmApiHostManagementConfigStandardManagementConfig *KrmApiHostManagementConfigStandardManagementConfig = &KrmApiHostManagementConfigStandardManagementConfig{empty: true}

func (r *KrmApiHostManagementConfigStandardManagementConfig) Empty() bool {
	return r.empty
}

func (r *KrmApiHostManagementConfigStandardManagementConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *KrmApiHostManagementConfigStandardManagementConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *KrmApiHost) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "krmapihosting",
		Type:    "KrmApiHost",
		Version: "alpha",
	}
}

const KrmApiHostMaxPage = -1

type KrmApiHostList struct {
	Items []*KrmApiHost

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *KrmApiHostList) HasNext() bool {
	return l.nextToken != ""
}

func (l *KrmApiHostList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listKrmApiHost(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListKrmApiHost(ctx context.Context, project, location string) (*KrmApiHostList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListKrmApiHostWithMaxResults(ctx, project, location, KrmApiHostMaxPage)

}

func (c *Client) ListKrmApiHostWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*KrmApiHostList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listKrmApiHost(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &KrmApiHostList{
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
func (r *KrmApiHost) URLNormalized() *KrmApiHost {
	normalized := dcl.Copy(*r).(KrmApiHost)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.GkeResourceLink = dcl.SelfLinkToName(r.GkeResourceLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetKrmApiHost(ctx context.Context, r *KrmApiHost) (*KrmApiHost, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getKrmApiHostRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalKrmApiHost(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeKrmApiHostNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteKrmApiHost(ctx context.Context, r *KrmApiHost) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("KrmApiHost resource is nil")
	}
	c.Config.Logger.Info("Deleting KrmApiHost...")
	deleteOp := deleteKrmApiHostOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllKrmApiHost deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllKrmApiHost(ctx context.Context, project, location string, filter func(*KrmApiHost) bool) error {
	listObj, err := c.ListKrmApiHost(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllKrmApiHost(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllKrmApiHost(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyKrmApiHost(ctx context.Context, rawDesired *KrmApiHost, opts ...dcl.ApplyOption) (*KrmApiHost, error) {
	var resultNewState *KrmApiHost
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyKrmApiHostHelper(c, ctx, rawDesired, opts...)
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

func applyKrmApiHostHelper(c *Client, ctx context.Context, rawDesired *KrmApiHost, opts ...dcl.ApplyOption) (*KrmApiHost, error) {
	c.Config.Logger.Info("Beginning ApplyKrmApiHost...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.krmApiHostDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToKrmApiHostDiffs(c.Config, fieldDiffs, opts)
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
	var ops []krmApiHostApiOperation
	if create {
		ops = append(ops, &createKrmApiHostOperation{})
	} else if recreate {
		ops = append(ops, &deleteKrmApiHostOperation{})
		ops = append(ops, &createKrmApiHostOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeKrmApiHostDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetKrmApiHost(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createKrmApiHostOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapKrmApiHost(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeKrmApiHostNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeKrmApiHostNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeKrmApiHostDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffKrmApiHost(c, newDesired, newState)
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
