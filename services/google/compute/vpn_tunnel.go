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
	"fmt"
	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type VpnTunnel struct {
	Id                           *int64               `json:"id"`
	Name                         *string              `json:"name"`
	Description                  *string              `json:"description"`
	Region                       *string              `json:"region"`
	TargetVpnGateway             *string              `json:"targetVpnGateway"`
	VpnGateway                   *string              `json:"vpnGateway"`
	VpnGatewayInterface          *int64               `json:"vpnGatewayInterface"`
	PeerExternalGateway          *string              `json:"peerExternalGateway"`
	PeerExternalGatewayInterface *int64               `json:"peerExternalGatewayInterface"`
	PeerGcpGateway               *string              `json:"peerGcpGateway"`
	Router                       *string              `json:"router"`
	PeerIP                       *string              `json:"peerIP"`
	SharedSecret                 *string              `json:"sharedSecret"`
	SharedSecretHash             *string              `json:"sharedSecretHash"`
	Status                       *VpnTunnelStatusEnum `json:"status"`
	SelfLink                     *string              `json:"selfLink"`
	IkeVersion                   *int64               `json:"ikeVersion"`
	DetailedStatus               *string              `json:"detailedStatus"`
	LocalTrafficSelector         []string             `json:"localTrafficSelector"`
	RemoteTrafficSelector        []string             `json:"remoteTrafficSelector"`
	Project                      *string              `json:"project"`
}

func (r *VpnTunnel) String() string {
	return dcl.SprintResource(r)
}

// The enum VpnTunnelStatusEnum.
type VpnTunnelStatusEnum string

// VpnTunnelStatusEnumRef returns a *VpnTunnelStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func VpnTunnelStatusEnumRef(s string) *VpnTunnelStatusEnum {
	if s == "" {
		return nil
	}

	v := VpnTunnelStatusEnum(s)
	return &v
}

func (v VpnTunnelStatusEnum) Validate() error {
	for _, s := range []string{"PROVISIONING", "WAITING_FOR_FULL_CONFIG", "FIRST_HANDSHAKE", "ESTABLISHED", "NO_INCOMING_PACKETS", "AUTHORIZATION_ERROR", "NEGOTIATION_FAILURE", "DEPROVISIONING", "FAILED", "REJECTED", "ALLOCATING_RESOURCES", "STOPPED", "PEER_IDENTITY_MISMATCH", "TS_NARROWING_NOT_ALLOWED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "VpnTunnelStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *VpnTunnel) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "VpnTunnel",
		Version: "compute",
	}
}

const VpnTunnelMaxPage = -1

type VpnTunnelList struct {
	Items []*VpnTunnel

	nextToken string

	pageSize int32

	project string

	region string
}

func (l *VpnTunnelList) HasNext() bool {
	return l.nextToken != ""
}

func (l *VpnTunnelList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listVpnTunnel(ctx, l.project, l.region, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListVpnTunnel(ctx context.Context, project, region string) (*VpnTunnelList, error) {

	return c.ListVpnTunnelWithMaxResults(ctx, project, region, VpnTunnelMaxPage)

}

func (c *Client) ListVpnTunnelWithMaxResults(ctx context.Context, project, region string, pageSize int32) (*VpnTunnelList, error) {
	items, token, err := c.listVpnTunnel(ctx, project, region, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &VpnTunnelList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		region: region,
	}, nil
}

func (c *Client) GetVpnTunnel(ctx context.Context, r *VpnTunnel) (*VpnTunnel, error) {
	b, err := c.getVpnTunnelRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalVpnTunnel(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Region = r.Region
	result.Name = r.Name
	if dcl.IsZeroValue(result.IkeVersion) {
		result.IkeVersion = dcl.Int64(2)
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeVpnTunnelNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteVpnTunnel(ctx context.Context, r *VpnTunnel) error {
	if r == nil {
		return fmt.Errorf("VpnTunnel resource is nil")
	}
	c.Config.Logger.Info("Deleting VpnTunnel...")
	deleteOp := deleteVpnTunnelOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllVpnTunnel deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllVpnTunnel(ctx context.Context, project, region string, filter func(*VpnTunnel) bool) error {
	listObj, err := c.ListVpnTunnel(ctx, project, region)
	if err != nil {
		return err
	}

	err = c.deleteAllVpnTunnel(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllVpnTunnel(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyVpnTunnel(ctx context.Context, rawDesired *VpnTunnel, opts ...dcl.ApplyOption) (*VpnTunnel, error) {
	c.Config.Logger.Info("Beginning ApplyVpnTunnel...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.vpnTunnelDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []vpnTunnelApiOperation
	if create {
		ops = append(ops, &createVpnTunnelOperation{})
	} else if recreate {
		ops = append(ops, &deleteVpnTunnelOperation{})
		ops = append(ops, &createVpnTunnelOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeVpnTunnelDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetVpnTunnel(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeVpnTunnelNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeVpnTunnelDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffVpnTunnel(c, newDesired, newState)
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
