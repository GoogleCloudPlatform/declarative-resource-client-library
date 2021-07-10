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
package tpu

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Node struct {
	Name                 *string                `json:"name"`
	Description          *string                `json:"description"`
	AcceleratorType      *string                `json:"acceleratorType"`
	IPAddress            *string                `json:"ipAddress"`
	Port                 *string                `json:"port"`
	State                *NodeStateEnum         `json:"state"`
	HealthDescription    *string                `json:"healthDescription"`
	TensorflowVersion    *string                `json:"tensorflowVersion"`
	Network              *string                `json:"network"`
	CidrBlock            *string                `json:"cidrBlock"`
	ServiceAccount       *string                `json:"serviceAccount"`
	CreateTime           *NodeCreateTime        `json:"createTime"`
	SchedulingConfig     *NodeSchedulingConfig  `json:"schedulingConfig"`
	NetworkEndpoints     []NodeNetworkEndpoints `json:"networkEndpoints"`
	Health               *NodeHealthEnum        `json:"health"`
	Labels               map[string]string      `json:"labels"`
	UseServiceNetworking *bool                  `json:"useServiceNetworking"`
	Symptoms             []NodeSymptoms         `json:"symptoms"`
	Project              *string                `json:"project"`
	Location             *string                `json:"location"`
}

func (r *Node) String() string {
	return dcl.SprintResource(r)
}

// The enum NodeStateEnum.
type NodeStateEnum string

// NodeStateEnumRef returns a *NodeStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func NodeStateEnumRef(s string) *NodeStateEnum {
	if s == "" {
		return nil
	}

	v := NodeStateEnum(s)
	return &v
}

func (v NodeStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "CREATING", "READY", "RESTARTING", "REIMAGING", "DELETING", "REPAIRING", "STOPPED", "STOPPING", "STARTING", "PREEMPTED", "TERMINATED", "HIDING", "HIDDEN", "UNHIDING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NodeStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NodeHealthEnum.
type NodeHealthEnum string

// NodeHealthEnumRef returns a *NodeHealthEnum with the value of string s
// If the empty string is provided, nil is returned.
func NodeHealthEnumRef(s string) *NodeHealthEnum {
	if s == "" {
		return nil
	}

	v := NodeHealthEnum(s)
	return &v
}

func (v NodeHealthEnum) Validate() error {
	for _, s := range []string{"HEALTH_UNSPECIFIED", "HEALTHY", "DEPRECATED_UNHEALTHY", "TIMEOUT", "UNHEALTHY_TENSORFLOW", "UNHEALTHY_MAINTENANCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NodeHealthEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NodeSymptomsSymptomTypeEnum.
type NodeSymptomsSymptomTypeEnum string

// NodeSymptomsSymptomTypeEnumRef returns a *NodeSymptomsSymptomTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func NodeSymptomsSymptomTypeEnumRef(s string) *NodeSymptomsSymptomTypeEnum {
	if s == "" {
		return nil
	}

	v := NodeSymptomsSymptomTypeEnum(s)
	return &v
}

func (v NodeSymptomsSymptomTypeEnum) Validate() error {
	for _, s := range []string{"SYMPTOM_TYPE_UNSPECIFIED", "LOW_MEMORY", "OUT_OF_MEMORY", "EXECUTE_TIMED_OUT", "MESH_BUILD_FAIL", "HBM_OUT_OF_MEMORY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NodeSymptomsSymptomTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type NodeCreateTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonNodeCreateTime NodeCreateTime

func (r *NodeCreateTime) UnmarshalJSON(data []byte) error {
	var res jsonNodeCreateTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodeCreateTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this NodeCreateTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodeCreateTime *NodeCreateTime = &NodeCreateTime{empty: true}

func (r *NodeCreateTime) Empty() bool {
	return r.empty
}

func (r *NodeCreateTime) String() string {
	return dcl.SprintResource(r)
}

func (r *NodeCreateTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodeSchedulingConfig struct {
	empty       bool  `json:"-"`
	Preemptible *bool `json:"preemptible"`
	Reserved    *bool `json:"reserved"`
}

type jsonNodeSchedulingConfig NodeSchedulingConfig

func (r *NodeSchedulingConfig) UnmarshalJSON(data []byte) error {
	var res jsonNodeSchedulingConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodeSchedulingConfig
	} else {

		r.Preemptible = res.Preemptible

		r.Reserved = res.Reserved

	}
	return nil
}

// This object is used to assert a desired state where this NodeSchedulingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodeSchedulingConfig *NodeSchedulingConfig = &NodeSchedulingConfig{empty: true}

func (r *NodeSchedulingConfig) Empty() bool {
	return r.empty
}

func (r *NodeSchedulingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *NodeSchedulingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodeNetworkEndpoints struct {
	empty     bool    `json:"-"`
	IPAddress *string `json:"ipAddress"`
	Port      *int64  `json:"port"`
}

type jsonNodeNetworkEndpoints NodeNetworkEndpoints

func (r *NodeNetworkEndpoints) UnmarshalJSON(data []byte) error {
	var res jsonNodeNetworkEndpoints
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodeNetworkEndpoints
	} else {

		r.IPAddress = res.IPAddress

		r.Port = res.Port

	}
	return nil
}

// This object is used to assert a desired state where this NodeNetworkEndpoints is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodeNetworkEndpoints *NodeNetworkEndpoints = &NodeNetworkEndpoints{empty: true}

func (r *NodeNetworkEndpoints) Empty() bool {
	return r.empty
}

func (r *NodeNetworkEndpoints) String() string {
	return dcl.SprintResource(r)
}

func (r *NodeNetworkEndpoints) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodeSymptoms struct {
	empty       bool                         `json:"-"`
	CreateTime  *NodeSymptomsCreateTime      `json:"createTime"`
	SymptomType *NodeSymptomsSymptomTypeEnum `json:"symptomType"`
	Details     *string                      `json:"details"`
	WorkerId    *string                      `json:"workerId"`
}

type jsonNodeSymptoms NodeSymptoms

func (r *NodeSymptoms) UnmarshalJSON(data []byte) error {
	var res jsonNodeSymptoms
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodeSymptoms
	} else {

		r.CreateTime = res.CreateTime

		r.SymptomType = res.SymptomType

		r.Details = res.Details

		r.WorkerId = res.WorkerId

	}
	return nil
}

// This object is used to assert a desired state where this NodeSymptoms is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodeSymptoms *NodeSymptoms = &NodeSymptoms{empty: true}

func (r *NodeSymptoms) Empty() bool {
	return r.empty
}

func (r *NodeSymptoms) String() string {
	return dcl.SprintResource(r)
}

func (r *NodeSymptoms) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NodeSymptomsCreateTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonNodeSymptomsCreateTime NodeSymptomsCreateTime

func (r *NodeSymptomsCreateTime) UnmarshalJSON(data []byte) error {
	var res jsonNodeSymptomsCreateTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNodeSymptomsCreateTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this NodeSymptomsCreateTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNodeSymptomsCreateTime *NodeSymptomsCreateTime = &NodeSymptomsCreateTime{empty: true}

func (r *NodeSymptomsCreateTime) Empty() bool {
	return r.empty
}

func (r *NodeSymptomsCreateTime) String() string {
	return dcl.SprintResource(r)
}

func (r *NodeSymptomsCreateTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Node) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "tpu",
		Type:    "Node",
		Version: "tpu",
	}
}

const NodeMaxPage = -1

type NodeList struct {
	Items []*Node

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *NodeList) HasNext() bool {
	return l.nextToken != ""
}

func (l *NodeList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listNode(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListNode(ctx context.Context, project, location string) (*NodeList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListNodeWithMaxResults(ctx, project, location, NodeMaxPage)

}

func (c *Client) ListNodeWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*NodeList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listNode(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &NodeList{
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
func (r *Node) URLNormalized() *Node {
	normalized := dcl.Copy(*r).(Node)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AcceleratorType = dcl.SelfLinkToName(r.AcceleratorType)
	normalized.IPAddress = dcl.SelfLinkToName(r.IPAddress)
	normalized.Port = dcl.SelfLinkToName(r.Port)
	normalized.HealthDescription = dcl.SelfLinkToName(r.HealthDescription)
	normalized.TensorflowVersion = dcl.SelfLinkToName(r.TensorflowVersion)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.CidrBlock = dcl.SelfLinkToName(r.CidrBlock)
	normalized.ServiceAccount = dcl.SelfLinkToName(r.ServiceAccount)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetNode(ctx context.Context, r *Node) (*Node, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getNodeRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalNode(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeNodeNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteNode(ctx context.Context, r *Node) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Node resource is nil")
	}
	c.Config.Logger.Info("Deleting Node...")
	deleteOp := deleteNodeOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllNode deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllNode(ctx context.Context, project, location string, filter func(*Node) bool) error {
	listObj, err := c.ListNode(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllNode(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllNode(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyNode(ctx context.Context, rawDesired *Node, opts ...dcl.ApplyOption) (*Node, error) {

	var resultNewState *Node
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyNodeHelper(c, ctx, rawDesired, opts...)
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

func applyNodeHelper(c *Client, ctx context.Context, rawDesired *Node, opts ...dcl.ApplyOption) (*Node, error) {
	c.Config.Logger.Info("Beginning ApplyNode...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.nodeDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToNodeOp(opStrings, fieldDiffs, opts)
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
	var ops []nodeApiOperation
	if create {
		ops = append(ops, &createNodeOperation{})
	} else if recreate {
		ops = append(ops, &deleteNodeOperation{})
		ops = append(ops, &createNodeOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeNodeDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetNode(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createNodeOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapNode(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeNodeNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeNodeNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeNodeDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffNode(c, newDesired, newState)
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
