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
package dns

import (
	"context"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type ResourceRecordSet struct {
	DnsName     *string  `json:"dnsName"`
	DnsType     *string  `json:"dnsType"`
	Ttl         *int64   `json:"ttl"`
	Target      []string `json:"target"`
	ManagedZone *string  `json:"managedZone"`
	Project     *string  `json:"project"`
}

func (r *ResourceRecordSet) String() string {
	return dcl.SprintResource(r)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ResourceRecordSet) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "dns",
		Type:    "ResourceRecordSet",
		Version: "dns",
	}
}

const ResourceRecordSetMaxPage = -1

type ResourceRecordSetList struct {
	Items []*ResourceRecordSet

	nextToken string

	pageSize int32

	project string

	managedZone string
}

func (l *ResourceRecordSetList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ResourceRecordSetList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listResourceRecordSet(ctx, l.project, l.managedZone, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListResourceRecordSet(ctx context.Context, project, managedZone string) (*ResourceRecordSetList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListResourceRecordSetWithMaxResults(ctx, project, managedZone, ResourceRecordSetMaxPage)

}

func (c *Client) ListResourceRecordSetWithMaxResults(ctx context.Context, project, managedZone string, pageSize int32) (*ResourceRecordSetList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listResourceRecordSet(ctx, project, managedZone, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ResourceRecordSetList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		managedZone: managedZone,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *ResourceRecordSet) URLNormalized() *ResourceRecordSet {
	normalized := dcl.Copy(*r).(ResourceRecordSet)
	normalized.DnsName = dcl.SelfLinkToName(r.DnsName)
	normalized.DnsType = dcl.SelfLinkToName(r.DnsType)
	normalized.ManagedZone = dcl.SelfLinkToName(r.ManagedZone)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) DeleteResourceRecordSet(ctx context.Context, r *ResourceRecordSet) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("ResourceRecordSet resource is nil")
	}
	c.Config.Logger.Info("Deleting ResourceRecordSet...")
	deleteOp := deleteResourceRecordSetOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllResourceRecordSet deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllResourceRecordSet(ctx context.Context, project, managedZone string, filter func(*ResourceRecordSet) bool) error {
	listObj, err := c.ListResourceRecordSet(ctx, project, managedZone)
	if err != nil {
		return err
	}

	err = c.deleteAllResourceRecordSet(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllResourceRecordSet(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyResourceRecordSet(ctx context.Context, rawDesired *ResourceRecordSet, opts ...dcl.ApplyOption) (*ResourceRecordSet, error) {
	var resultNewState *ResourceRecordSet
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyResourceRecordSetHelper(c, ctx, rawDesired, opts...)
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

func applyResourceRecordSetHelper(c *Client, ctx context.Context, rawDesired *ResourceRecordSet, opts ...dcl.ApplyOption) (*ResourceRecordSet, error) {
	c.Config.Logger.Info("Beginning ApplyResourceRecordSet...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.resourceRecordSetDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToResourceRecordSetDiffs(c.Config, fieldDiffs, opts)
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
	var ops []resourceRecordSetApiOperation
	if create {
		ops = append(ops, &createResourceRecordSetOperation{})
	} else if recreate {
		ops = append(ops, &deleteResourceRecordSetOperation{})
		ops = append(ops, &createResourceRecordSetOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeResourceRecordSetDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetResourceRecordSet(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createResourceRecordSetOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapResourceRecordSet(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeResourceRecordSetNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeResourceRecordSetNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeResourceRecordSetDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffResourceRecordSet(c, newDesired, newState)
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
