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
package gkemulticloud

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type AzureNodePool struct {
	Name                  *string                         `json:"name"`
	Version               *string                         `json:"version"`
	Config                *AzureNodePoolConfig            `json:"config"`
	SubnetId              *string                         `json:"subnetId"`
	Autoscaling           *AzureNodePoolAutoscaling       `json:"autoscaling"`
	State                 *AzureNodePoolStateEnum         `json:"state"`
	Uid                   *string                         `json:"uid"`
	Reconciling           *bool                           `json:"reconciling"`
	CreateTime            *string                         `json:"createTime"`
	UpdateTime            *string                         `json:"updateTime"`
	Etag                  *string                         `json:"etag"`
	Annotations           map[string]string               `json:"annotations"`
	MaxPodsConstraint     *AzureNodePoolMaxPodsConstraint `json:"maxPodsConstraint"`
	AzureAvailabilityZone *string                         `json:"azureAvailabilityZone"`
	Project               *string                         `json:"project"`
	Location              *string                         `json:"location"`
	AzureCluster          *string                         `json:"azureCluster"`
}

func (r *AzureNodePool) String() string {
	return dcl.SprintResource(r)
}

// The enum AzureNodePoolStateEnum.
type AzureNodePoolStateEnum string

// AzureNodePoolStateEnumRef returns a *AzureNodePoolStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func AzureNodePoolStateEnumRef(s string) *AzureNodePoolStateEnum {
	if s == "" {
		return nil
	}

	v := AzureNodePoolStateEnum(s)
	return &v
}

func (v AzureNodePoolStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "PROVISIONING", "RUNNING", "RECONCILING", "STOPPING", "ERROR", "DEGRADED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AzureNodePoolStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AzureNodePoolConfig struct {
	empty      bool                           `json:"-"`
	VmSize     *string                        `json:"vmSize"`
	RootVolume *AzureNodePoolConfigRootVolume `json:"rootVolume"`
	Tags       map[string]string              `json:"tags"`
	SshConfig  *AzureNodePoolConfigSshConfig  `json:"sshConfig"`
}

type jsonAzureNodePoolConfig AzureNodePoolConfig

func (r *AzureNodePoolConfig) UnmarshalJSON(data []byte) error {
	var res jsonAzureNodePoolConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureNodePoolConfig
	} else {

		r.VmSize = res.VmSize

		r.RootVolume = res.RootVolume

		r.Tags = res.Tags

		r.SshConfig = res.SshConfig

	}
	return nil
}

// This object is used to assert a desired state where this AzureNodePoolConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureNodePoolConfig *AzureNodePoolConfig = &AzureNodePoolConfig{empty: true}

func (r *AzureNodePoolConfig) Empty() bool {
	return r.empty
}

func (r *AzureNodePoolConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureNodePoolConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureNodePoolConfigRootVolume struct {
	empty   bool   `json:"-"`
	SizeGib *int64 `json:"sizeGib"`
}

type jsonAzureNodePoolConfigRootVolume AzureNodePoolConfigRootVolume

func (r *AzureNodePoolConfigRootVolume) UnmarshalJSON(data []byte) error {
	var res jsonAzureNodePoolConfigRootVolume
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureNodePoolConfigRootVolume
	} else {

		r.SizeGib = res.SizeGib

	}
	return nil
}

// This object is used to assert a desired state where this AzureNodePoolConfigRootVolume is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureNodePoolConfigRootVolume *AzureNodePoolConfigRootVolume = &AzureNodePoolConfigRootVolume{empty: true}

func (r *AzureNodePoolConfigRootVolume) Empty() bool {
	return r.empty
}

func (r *AzureNodePoolConfigRootVolume) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureNodePoolConfigRootVolume) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureNodePoolConfigSshConfig struct {
	empty         bool    `json:"-"`
	AuthorizedKey *string `json:"authorizedKey"`
}

type jsonAzureNodePoolConfigSshConfig AzureNodePoolConfigSshConfig

func (r *AzureNodePoolConfigSshConfig) UnmarshalJSON(data []byte) error {
	var res jsonAzureNodePoolConfigSshConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureNodePoolConfigSshConfig
	} else {

		r.AuthorizedKey = res.AuthorizedKey

	}
	return nil
}

// This object is used to assert a desired state where this AzureNodePoolConfigSshConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureNodePoolConfigSshConfig *AzureNodePoolConfigSshConfig = &AzureNodePoolConfigSshConfig{empty: true}

func (r *AzureNodePoolConfigSshConfig) Empty() bool {
	return r.empty
}

func (r *AzureNodePoolConfigSshConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureNodePoolConfigSshConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureNodePoolAutoscaling struct {
	empty        bool   `json:"-"`
	MinNodeCount *int64 `json:"minNodeCount"`
	MaxNodeCount *int64 `json:"maxNodeCount"`
}

type jsonAzureNodePoolAutoscaling AzureNodePoolAutoscaling

func (r *AzureNodePoolAutoscaling) UnmarshalJSON(data []byte) error {
	var res jsonAzureNodePoolAutoscaling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureNodePoolAutoscaling
	} else {

		r.MinNodeCount = res.MinNodeCount

		r.MaxNodeCount = res.MaxNodeCount

	}
	return nil
}

// This object is used to assert a desired state where this AzureNodePoolAutoscaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureNodePoolAutoscaling *AzureNodePoolAutoscaling = &AzureNodePoolAutoscaling{empty: true}

func (r *AzureNodePoolAutoscaling) Empty() bool {
	return r.empty
}

func (r *AzureNodePoolAutoscaling) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureNodePoolAutoscaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AzureNodePoolMaxPodsConstraint struct {
	empty          bool   `json:"-"`
	MaxPodsPerNode *int64 `json:"maxPodsPerNode"`
}

type jsonAzureNodePoolMaxPodsConstraint AzureNodePoolMaxPodsConstraint

func (r *AzureNodePoolMaxPodsConstraint) UnmarshalJSON(data []byte) error {
	var res jsonAzureNodePoolMaxPodsConstraint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAzureNodePoolMaxPodsConstraint
	} else {

		r.MaxPodsPerNode = res.MaxPodsPerNode

	}
	return nil
}

// This object is used to assert a desired state where this AzureNodePoolMaxPodsConstraint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAzureNodePoolMaxPodsConstraint *AzureNodePoolMaxPodsConstraint = &AzureNodePoolMaxPodsConstraint{empty: true}

func (r *AzureNodePoolMaxPodsConstraint) Empty() bool {
	return r.empty
}

func (r *AzureNodePoolMaxPodsConstraint) String() string {
	return dcl.SprintResource(r)
}

func (r *AzureNodePoolMaxPodsConstraint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *AzureNodePool) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "gkemulticloud",
		Type:    "AzureNodePool",
		Version: "gkemulticloud",
	}
}

func (r *AzureNodePool) ID() (string, error) {
	if err := extractAzureNodePoolFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                  dcl.ValueOrEmptyString(nr.Name),
		"version":               dcl.ValueOrEmptyString(nr.Version),
		"config":                dcl.ValueOrEmptyString(nr.Config),
		"subnetId":              dcl.ValueOrEmptyString(nr.SubnetId),
		"autoscaling":           dcl.ValueOrEmptyString(nr.Autoscaling),
		"state":                 dcl.ValueOrEmptyString(nr.State),
		"uid":                   dcl.ValueOrEmptyString(nr.Uid),
		"reconciling":           dcl.ValueOrEmptyString(nr.Reconciling),
		"createTime":            dcl.ValueOrEmptyString(nr.CreateTime),
		"updateTime":            dcl.ValueOrEmptyString(nr.UpdateTime),
		"etag":                  dcl.ValueOrEmptyString(nr.Etag),
		"annotations":           dcl.ValueOrEmptyString(nr.Annotations),
		"maxPodsConstraint":     dcl.ValueOrEmptyString(nr.MaxPodsConstraint),
		"azureAvailabilityZone": dcl.ValueOrEmptyString(nr.AzureAvailabilityZone),
		"project":               dcl.ValueOrEmptyString(nr.Project),
		"location":              dcl.ValueOrEmptyString(nr.Location),
		"azureCluster":          dcl.ValueOrEmptyString(nr.AzureCluster),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/azureClusters/{{azure_cluster}}/azureNodePools/{{name}}", params), nil
}

const AzureNodePoolMaxPage = -1

type AzureNodePoolList struct {
	Items []*AzureNodePool

	nextToken string

	pageSize int32

	resource *AzureNodePool
}

func (l *AzureNodePoolList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AzureNodePoolList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAzureNodePool(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAzureNodePool(ctx context.Context, r *AzureNodePool) (*AzureNodePoolList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAzureNodePoolWithMaxResults(ctx, r, AzureNodePoolMaxPage)

}

func (c *Client) ListAzureNodePoolWithMaxResults(ctx context.Context, r *AzureNodePool, pageSize int32) (*AzureNodePoolList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAzureNodePool(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AzureNodePoolList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetAzureNodePool(ctx context.Context, r *AzureNodePool) (*AzureNodePool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAzureNodePoolRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAzureNodePool(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.AzureCluster = r.AzureCluster
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAzureNodePoolNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAzureNodePool(ctx context.Context, r *AzureNodePool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("AzureNodePool resource is nil")
	}
	c.Config.Logger.Info("Deleting AzureNodePool...")
	deleteOp := deleteAzureNodePoolOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAzureNodePool deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAzureNodePool(ctx context.Context, project, location, azureCluster string, filter func(*AzureNodePool) bool) error {
	r := &AzureNodePool{

		Project: &project,

		Location: &location,

		AzureCluster: &azureCluster,
	}
	listObj, err := c.ListAzureNodePool(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllAzureNodePool(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAzureNodePool(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAzureNodePool(ctx context.Context, rawDesired *AzureNodePool, opts ...dcl.ApplyOption) (*AzureNodePool, error) {
	var resultNewState *AzureNodePool
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAzureNodePoolHelper(c, ctx, rawDesired, opts...)
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

func applyAzureNodePoolHelper(c *Client, ctx context.Context, rawDesired *AzureNodePool, opts ...dcl.ApplyOption) (*AzureNodePool, error) {
	c.Config.Logger.Info("Beginning ApplyAzureNodePool...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractAzureNodePoolFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.azureNodePoolDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToAzureNodePoolDiffs(c.Config, fieldDiffs, opts)
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
	var ops []azureNodePoolApiOperation
	if create {
		ops = append(ops, &createAzureNodePoolOperation{})
	} else if recreate {
		ops = append(ops, &deleteAzureNodePoolOperation{})
		ops = append(ops, &createAzureNodePoolOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAzureNodePoolDesiredState(rawDesired, nil)
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
		c.Config.Logger.Infof("Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetAzureNodePool(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAzureNodePoolOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAzureNodePool(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAzureNodePoolNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAzureNodePoolNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAzureNodePoolDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAzureNodePool(c, newDesired, newState)
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
