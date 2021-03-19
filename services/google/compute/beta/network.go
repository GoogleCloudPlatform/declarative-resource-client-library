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

type Network struct {
	Description           *string               `json:"description"`
	GatewayIPv4           *string               `json:"gatewayIPv4"`
	IPv4Range             *string               `json:"ipv4Range"`
	Name                  *string               `json:"name"`
	AutoCreateSubnetworks *bool                 `json:"autoCreateSubnetworks"`
	RoutingConfig         *NetworkRoutingConfig `json:"routingConfig"`
	Project               *string               `json:"project"`
	SelfLink              *string               `json:"selfLink"`
}

func (r *Network) String() string {
	return dcl.SprintResource(r)
}

// The enum NetworkRoutingConfigRoutingModeEnum.
type NetworkRoutingConfigRoutingModeEnum string

// NetworkRoutingConfigRoutingModeEnumRef returns a *NetworkRoutingConfigRoutingModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func NetworkRoutingConfigRoutingModeEnumRef(s string) *NetworkRoutingConfigRoutingModeEnum {
	if s == "" {
		return nil
	}

	v := NetworkRoutingConfigRoutingModeEnum(s)
	return &v
}

func (v NetworkRoutingConfigRoutingModeEnum) Validate() error {
	for _, s := range []string{"REGIONAL", "GLOBAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NetworkRoutingConfigRoutingModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type NetworkRoutingConfig struct {
	empty       bool                                 `json:"-"`
	RoutingMode *NetworkRoutingConfigRoutingModeEnum `json:"routingMode"`
}

// This object is used to assert a desired state where this NetworkRoutingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNetworkRoutingConfig *NetworkRoutingConfig = &NetworkRoutingConfig{empty: true}

func (r *NetworkRoutingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *NetworkRoutingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Network) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Network",
		Version: "beta",
	}
}

const NetworkMaxPage = -1

type NetworkList struct {
	Items []*Network

	nextToken string

	pageSize int32

	project string
}

func (l *NetworkList) HasNext() bool {
	return l.nextToken != ""
}

func (l *NetworkList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listNetwork(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListNetwork(ctx context.Context, project string) (*NetworkList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListNetworkWithMaxResults(ctx, project, NetworkMaxPage)

}

func (c *Client) ListNetworkWithMaxResults(ctx context.Context, project string, pageSize int32) (*NetworkList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listNetwork(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &NetworkList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetNetwork(ctx context.Context, r *Network) (*Network, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getNetworkRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalNetwork(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name
	if dcl.IsZeroValue(result.AutoCreateSubnetworks) {
		result.AutoCreateSubnetworks = dcl.Bool(true)
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeNetworkNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteNetwork(ctx context.Context, r *Network) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Network resource is nil")
	}
	c.Config.Logger.Info("Deleting Network...")
	deleteOp := deleteNetworkOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllNetwork deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllNetwork(ctx context.Context, project string, filter func(*Network) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListNetwork(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllNetwork(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllNetwork(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyNetwork(ctx context.Context, rawDesired *Network, opts ...dcl.ApplyOption) (*Network, error) {
	c.Config.Logger.Info("Beginning ApplyNetwork...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.networkDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []networkApiOperation
	if create {
		ops = append(ops, &createNetworkOperation{})
	} else if recreate {

		ops = append(ops, &deleteNetworkOperation{})

		ops = append(ops, &createNetworkOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeNetworkDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetNetwork(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createNetworkOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapNetwork(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeNetworkNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeNetworkNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeNetworkDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffNetwork(c, newDesired, newState)
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
