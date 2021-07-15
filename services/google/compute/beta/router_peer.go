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

type RouterPeer struct {
	CreationTimestamp       *string                        `json:"creationTimestamp"`
	Router                  *string                        `json:"router"`
	Name                    *string                        `json:"name"`
	InterfaceName           *string                        `json:"interfaceName"`
	IPAddress               *string                        `json:"ipAddress"`
	PeerIPAddress           *string                        `json:"peerIPAddress"`
	PeerAsn                 *int64                         `json:"peerAsn"`
	AdvertisedRoutePriority *int64                         `json:"advertisedRoutePriority"`
	AdvertiseMode           *string                        `json:"advertiseMode"`
	ManagementType          *string                        `json:"managementType"`
	AdvertisedGroups        []string                       `json:"advertisedGroups"`
	AdvertisedIPRanges      []RouterPeerAdvertisedIPRanges `json:"advertisedIPRanges"`
	Region                  *string                        `json:"region"`
	Project                 *string                        `json:"project"`
}

func (r *RouterPeer) String() string {
	return dcl.SprintResource(r)
}

type RouterPeerAdvertisedIPRanges struct {
	empty       bool    `json:"-"`
	Range       *string `json:"range"`
	Description *string `json:"description"`
}

type jsonRouterPeerAdvertisedIPRanges RouterPeerAdvertisedIPRanges

func (r *RouterPeerAdvertisedIPRanges) UnmarshalJSON(data []byte) error {
	var res jsonRouterPeerAdvertisedIPRanges
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouterPeerAdvertisedIPRanges
	} else {

		r.Range = res.Range

		r.Description = res.Description

	}
	return nil
}

// This object is used to assert a desired state where this RouterPeerAdvertisedIPRanges is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouterPeerAdvertisedIPRanges *RouterPeerAdvertisedIPRanges = &RouterPeerAdvertisedIPRanges{empty: true}

func (r *RouterPeerAdvertisedIPRanges) Empty() bool {
	return r.empty
}

func (r *RouterPeerAdvertisedIPRanges) String() string {
	return dcl.SprintResource(r)
}

func (r *RouterPeerAdvertisedIPRanges) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *RouterPeer) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "RouterPeer",
		Version: "beta",
	}
}

const RouterPeerMaxPage = -1

type RouterPeerList struct {
	Items []*RouterPeer

	nextToken string

	pageSize int32

	project string

	region string
}

func (l *RouterPeerList) HasNext() bool {
	return l.nextToken != ""
}

func (l *RouterPeerList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listRouterPeer(ctx, l.project, l.region, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListRouterPeer(ctx context.Context, project, region string) (*RouterPeerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListRouterPeerWithMaxResults(ctx, project, region, RouterPeerMaxPage)

}

func (c *Client) ListRouterPeerWithMaxResults(ctx context.Context, project, region string, pageSize int32) (*RouterPeerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listRouterPeer(ctx, project, region, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &RouterPeerList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		region: region,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *RouterPeer) URLNormalized() *RouterPeer {
	normalized := dcl.Copy(*r).(RouterPeer)
	normalized.Router = dcl.SelfLinkToName(r.Router)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.InterfaceName = dcl.SelfLinkToName(r.InterfaceName)
	normalized.IPAddress = dcl.SelfLinkToName(r.IPAddress)
	normalized.PeerIPAddress = dcl.SelfLinkToName(r.PeerIPAddress)
	normalized.AdvertiseMode = dcl.SelfLinkToName(r.AdvertiseMode)
	normalized.ManagementType = dcl.SelfLinkToName(r.ManagementType)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetRouterPeer(ctx context.Context, r *RouterPeer) (*RouterPeer, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getRouterPeerRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	item, err := routerPeerFromParent(r, m)
	if err != nil {
		return nil, err
	}
	b, err = json.Marshal(item)
	if err != nil {
		return nil, err
	}

	result, err := unmarshalRouterPeer(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Region = r.Region
	result.Router = r.Router
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeRouterPeerNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteRouterPeer(ctx context.Context, r *RouterPeer) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("RouterPeer resource is nil")
	}
	c.Config.Logger.Info("Deleting RouterPeer...")
	deleteOp := deleteRouterPeerOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllRouterPeer deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllRouterPeer(ctx context.Context, project, region string, filter func(*RouterPeer) bool) error {
	listObj, err := c.ListRouterPeer(ctx, project, region)
	if err != nil {
		return err
	}

	err = c.deleteAllRouterPeer(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllRouterPeer(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyRouterPeer(ctx context.Context, rawDesired *RouterPeer, opts ...dcl.ApplyOption) (*RouterPeer, error) {
	var resultNewState *RouterPeer
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyRouterPeerHelper(c, ctx, rawDesired, opts...)
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

func applyRouterPeerHelper(c *Client, ctx context.Context, rawDesired *RouterPeer, opts ...dcl.ApplyOption) (*RouterPeer, error) {
	c.Config.Logger.Info("Beginning ApplyRouterPeer...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.routerPeerDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToRouterPeerDiffs(c.Config, fieldDiffs, opts)
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
	var ops []routerPeerApiOperation
	if create {
		ops = append(ops, &createRouterPeerOperation{})
	} else if recreate {
		ops = append(ops, &deleteRouterPeerOperation{})
		ops = append(ops, &createRouterPeerOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeRouterPeerDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetRouterPeer(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createRouterPeerOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapRouterPeer(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeRouterPeerNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeRouterPeerNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeRouterPeerDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffRouterPeer(c, newDesired, newState)
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
