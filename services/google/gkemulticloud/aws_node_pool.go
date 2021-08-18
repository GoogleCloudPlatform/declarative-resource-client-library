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

type AwsNodePool struct {
	Name              *string                       `json:"name"`
	Version           *string                       `json:"version"`
	Config            *AwsNodePoolConfig            `json:"config"`
	Autoscaling       *AwsNodePoolAutoscaling       `json:"autoscaling"`
	SubnetId          *string                       `json:"subnetId"`
	State             *AwsNodePoolStateEnum         `json:"state"`
	Uid               *string                       `json:"uid"`
	Reconciling       *bool                         `json:"reconciling"`
	CreateTime        *string                       `json:"createTime"`
	UpdateTime        *string                       `json:"updateTime"`
	Etag              *string                       `json:"etag"`
	Annotations       map[string]string             `json:"annotations"`
	MaxPodsConstraint *AwsNodePoolMaxPodsConstraint `json:"maxPodsConstraint"`
	Project           *string                       `json:"project"`
	Location          *string                       `json:"location"`
	AwsCluster        *string                       `json:"awsCluster"`
}

func (r *AwsNodePool) String() string {
	return dcl.SprintResource(r)
}

// The enum AwsNodePoolConfigRootVolumeVolumeTypeEnum.
type AwsNodePoolConfigRootVolumeVolumeTypeEnum string

// AwsNodePoolConfigRootVolumeVolumeTypeEnumRef returns a *AwsNodePoolConfigRootVolumeVolumeTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AwsNodePoolConfigRootVolumeVolumeTypeEnumRef(s string) *AwsNodePoolConfigRootVolumeVolumeTypeEnum {
	if s == "" {
		return nil
	}

	v := AwsNodePoolConfigRootVolumeVolumeTypeEnum(s)
	return &v
}

func (v AwsNodePoolConfigRootVolumeVolumeTypeEnum) Validate() error {
	for _, s := range []string{"VOLUME_TYPE_UNSPECIFIED", "GP2", "GP3"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AwsNodePoolConfigRootVolumeVolumeTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AwsNodePoolConfigTaintsEffectEnum.
type AwsNodePoolConfigTaintsEffectEnum string

// AwsNodePoolConfigTaintsEffectEnumRef returns a *AwsNodePoolConfigTaintsEffectEnum with the value of string s
// If the empty string is provided, nil is returned.
func AwsNodePoolConfigTaintsEffectEnumRef(s string) *AwsNodePoolConfigTaintsEffectEnum {
	if s == "" {
		return nil
	}

	v := AwsNodePoolConfigTaintsEffectEnum(s)
	return &v
}

func (v AwsNodePoolConfigTaintsEffectEnum) Validate() error {
	for _, s := range []string{"EFFECT_UNSPECIFIED", "NO_SCHEDULE", "PREFER_NO_SCHEDULE", "NO_EXECUTE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AwsNodePoolConfigTaintsEffectEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AwsNodePoolStateEnum.
type AwsNodePoolStateEnum string

// AwsNodePoolStateEnumRef returns a *AwsNodePoolStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func AwsNodePoolStateEnumRef(s string) *AwsNodePoolStateEnum {
	if s == "" {
		return nil
	}

	v := AwsNodePoolStateEnum(s)
	return &v
}

func (v AwsNodePoolStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "PROVISIONING", "RUNNING", "RECONCILING", "STOPPING", "ERROR", "DEGRADED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AwsNodePoolStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AwsNodePoolConfig struct {
	empty              bool                         `json:"-"`
	InstanceType       *string                      `json:"instanceType"`
	RootVolume         *AwsNodePoolConfigRootVolume `json:"rootVolume"`
	Taints             []AwsNodePoolConfigTaints    `json:"taints"`
	Labels             map[string]string            `json:"labels"`
	Tags               map[string]string            `json:"tags"`
	IamInstanceProfile *string                      `json:"iamInstanceProfile"`
	SshConfig          *AwsNodePoolConfigSshConfig  `json:"sshConfig"`
	SecurityGroupIds   []string                     `json:"securityGroupIds"`
}

type jsonAwsNodePoolConfig AwsNodePoolConfig

func (r *AwsNodePoolConfig) UnmarshalJSON(data []byte) error {
	var res jsonAwsNodePoolConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsNodePoolConfig
	} else {

		r.InstanceType = res.InstanceType

		r.RootVolume = res.RootVolume

		r.Taints = res.Taints

		r.Labels = res.Labels

		r.Tags = res.Tags

		r.IamInstanceProfile = res.IamInstanceProfile

		r.SshConfig = res.SshConfig

		r.SecurityGroupIds = res.SecurityGroupIds

	}
	return nil
}

// This object is used to assert a desired state where this AwsNodePoolConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsNodePoolConfig *AwsNodePoolConfig = &AwsNodePoolConfig{empty: true}

func (r *AwsNodePoolConfig) Empty() bool {
	return r.empty
}

func (r *AwsNodePoolConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsNodePoolConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsNodePoolConfigRootVolume struct {
	empty      bool                                       `json:"-"`
	SizeGib    *int64                                     `json:"sizeGib"`
	VolumeType *AwsNodePoolConfigRootVolumeVolumeTypeEnum `json:"volumeType"`
	Iops       *int64                                     `json:"iops"`
	KmsKeyArn  *string                                    `json:"kmsKeyArn"`
}

type jsonAwsNodePoolConfigRootVolume AwsNodePoolConfigRootVolume

func (r *AwsNodePoolConfigRootVolume) UnmarshalJSON(data []byte) error {
	var res jsonAwsNodePoolConfigRootVolume
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsNodePoolConfigRootVolume
	} else {

		r.SizeGib = res.SizeGib

		r.VolumeType = res.VolumeType

		r.Iops = res.Iops

		r.KmsKeyArn = res.KmsKeyArn

	}
	return nil
}

// This object is used to assert a desired state where this AwsNodePoolConfigRootVolume is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsNodePoolConfigRootVolume *AwsNodePoolConfigRootVolume = &AwsNodePoolConfigRootVolume{empty: true}

func (r *AwsNodePoolConfigRootVolume) Empty() bool {
	return r.empty
}

func (r *AwsNodePoolConfigRootVolume) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsNodePoolConfigRootVolume) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsNodePoolConfigTaints struct {
	empty  bool                               `json:"-"`
	Key    *string                            `json:"key"`
	Value  *string                            `json:"value"`
	Effect *AwsNodePoolConfigTaintsEffectEnum `json:"effect"`
}

type jsonAwsNodePoolConfigTaints AwsNodePoolConfigTaints

func (r *AwsNodePoolConfigTaints) UnmarshalJSON(data []byte) error {
	var res jsonAwsNodePoolConfigTaints
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsNodePoolConfigTaints
	} else {

		r.Key = res.Key

		r.Value = res.Value

		r.Effect = res.Effect

	}
	return nil
}

// This object is used to assert a desired state where this AwsNodePoolConfigTaints is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsNodePoolConfigTaints *AwsNodePoolConfigTaints = &AwsNodePoolConfigTaints{empty: true}

func (r *AwsNodePoolConfigTaints) Empty() bool {
	return r.empty
}

func (r *AwsNodePoolConfigTaints) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsNodePoolConfigTaints) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsNodePoolConfigSshConfig struct {
	empty      bool    `json:"-"`
	Ec2KeyPair *string `json:"ec2KeyPair"`
}

type jsonAwsNodePoolConfigSshConfig AwsNodePoolConfigSshConfig

func (r *AwsNodePoolConfigSshConfig) UnmarshalJSON(data []byte) error {
	var res jsonAwsNodePoolConfigSshConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsNodePoolConfigSshConfig
	} else {

		r.Ec2KeyPair = res.Ec2KeyPair

	}
	return nil
}

// This object is used to assert a desired state where this AwsNodePoolConfigSshConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsNodePoolConfigSshConfig *AwsNodePoolConfigSshConfig = &AwsNodePoolConfigSshConfig{empty: true}

func (r *AwsNodePoolConfigSshConfig) Empty() bool {
	return r.empty
}

func (r *AwsNodePoolConfigSshConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsNodePoolConfigSshConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsNodePoolAutoscaling struct {
	empty        bool   `json:"-"`
	MinNodeCount *int64 `json:"minNodeCount"`
	MaxNodeCount *int64 `json:"maxNodeCount"`
}

type jsonAwsNodePoolAutoscaling AwsNodePoolAutoscaling

func (r *AwsNodePoolAutoscaling) UnmarshalJSON(data []byte) error {
	var res jsonAwsNodePoolAutoscaling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsNodePoolAutoscaling
	} else {

		r.MinNodeCount = res.MinNodeCount

		r.MaxNodeCount = res.MaxNodeCount

	}
	return nil
}

// This object is used to assert a desired state where this AwsNodePoolAutoscaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsNodePoolAutoscaling *AwsNodePoolAutoscaling = &AwsNodePoolAutoscaling{empty: true}

func (r *AwsNodePoolAutoscaling) Empty() bool {
	return r.empty
}

func (r *AwsNodePoolAutoscaling) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsNodePoolAutoscaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AwsNodePoolMaxPodsConstraint struct {
	empty          bool   `json:"-"`
	MaxPodsPerNode *int64 `json:"maxPodsPerNode"`
}

type jsonAwsNodePoolMaxPodsConstraint AwsNodePoolMaxPodsConstraint

func (r *AwsNodePoolMaxPodsConstraint) UnmarshalJSON(data []byte) error {
	var res jsonAwsNodePoolMaxPodsConstraint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAwsNodePoolMaxPodsConstraint
	} else {

		r.MaxPodsPerNode = res.MaxPodsPerNode

	}
	return nil
}

// This object is used to assert a desired state where this AwsNodePoolMaxPodsConstraint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAwsNodePoolMaxPodsConstraint *AwsNodePoolMaxPodsConstraint = &AwsNodePoolMaxPodsConstraint{empty: true}

func (r *AwsNodePoolMaxPodsConstraint) Empty() bool {
	return r.empty
}

func (r *AwsNodePoolMaxPodsConstraint) String() string {
	return dcl.SprintResource(r)
}

func (r *AwsNodePoolMaxPodsConstraint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *AwsNodePool) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "gkemulticloud",
		Type:    "AwsNodePool",
		Version: "gkemulticloud",
	}
}

func (r *AwsNodePool) ID() (string, error) {
	if err := extractAwsNodePoolFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":              dcl.ValueOrEmptyString(nr.Name),
		"version":           dcl.ValueOrEmptyString(nr.Version),
		"config":            dcl.ValueOrEmptyString(nr.Config),
		"autoscaling":       dcl.ValueOrEmptyString(nr.Autoscaling),
		"subnetId":          dcl.ValueOrEmptyString(nr.SubnetId),
		"state":             dcl.ValueOrEmptyString(nr.State),
		"uid":               dcl.ValueOrEmptyString(nr.Uid),
		"reconciling":       dcl.ValueOrEmptyString(nr.Reconciling),
		"createTime":        dcl.ValueOrEmptyString(nr.CreateTime),
		"updateTime":        dcl.ValueOrEmptyString(nr.UpdateTime),
		"etag":              dcl.ValueOrEmptyString(nr.Etag),
		"annotations":       dcl.ValueOrEmptyString(nr.Annotations),
		"maxPodsConstraint": dcl.ValueOrEmptyString(nr.MaxPodsConstraint),
		"project":           dcl.ValueOrEmptyString(nr.Project),
		"location":          dcl.ValueOrEmptyString(nr.Location),
		"awsCluster":        dcl.ValueOrEmptyString(nr.AwsCluster),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/awsClusters/{{aws_cluster}}/awsNodePools/{{name}}", params), nil
}

const AwsNodePoolMaxPage = -1

type AwsNodePoolList struct {
	Items []*AwsNodePool

	nextToken string

	pageSize int32

	resource *AwsNodePool
}

func (l *AwsNodePoolList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AwsNodePoolList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAwsNodePool(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAwsNodePool(ctx context.Context, r *AwsNodePool) (*AwsNodePoolList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAwsNodePoolWithMaxResults(ctx, r, AwsNodePoolMaxPage)

}

func (c *Client) ListAwsNodePoolWithMaxResults(ctx context.Context, r *AwsNodePool, pageSize int32) (*AwsNodePoolList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAwsNodePool(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AwsNodePoolList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetAwsNodePool(ctx context.Context, r *AwsNodePool) (*AwsNodePool, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAwsNodePoolRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAwsNodePool(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.AwsCluster = r.AwsCluster
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAwsNodePoolNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAwsNodePool(ctx context.Context, r *AwsNodePool) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("AwsNodePool resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting AwsNodePool...")
	deleteOp := deleteAwsNodePoolOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAwsNodePool deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAwsNodePool(ctx context.Context, project, location, awsCluster string, filter func(*AwsNodePool) bool) error {
	r := &AwsNodePool{
		Project:    &project,
		Location:   &location,
		AwsCluster: &awsCluster,
	}
	listObj, err := c.ListAwsNodePool(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllAwsNodePool(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAwsNodePool(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAwsNodePool(ctx context.Context, rawDesired *AwsNodePool, opts ...dcl.ApplyOption) (*AwsNodePool, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *AwsNodePool
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAwsNodePoolHelper(c, ctx, rawDesired, opts...)
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

func applyAwsNodePoolHelper(c *Client, ctx context.Context, rawDesired *AwsNodePool, opts ...dcl.ApplyOption) (*AwsNodePool, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyAwsNodePool...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractAwsNodePoolFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.awsNodePoolDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToAwsNodePoolDiffs(c.Config, fieldDiffs, opts)
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
	var ops []awsNodePoolApiOperation
	if create {
		ops = append(ops, &createAwsNodePoolOperation{})
	} else if recreate {
		ops = append(ops, &deleteAwsNodePoolOperation{})
		ops = append(ops, &createAwsNodePoolOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAwsNodePoolDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetAwsNodePool(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAwsNodePoolOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAwsNodePool(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAwsNodePoolNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAwsNodePoolNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAwsNodePoolDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAwsNodePool(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
