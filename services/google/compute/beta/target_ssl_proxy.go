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
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type TargetSslProxy struct {
	Id              *int64                         `json:"id"`
	Name            *string                        `json:"name"`
	Description     *string                        `json:"description"`
	SelfLink        *string                        `json:"selfLink"`
	Service         *string                        `json:"service"`
	SslCertificates []string                       `json:"sslCertificates"`
	ProxyHeader     *TargetSslProxyProxyHeaderEnum `json:"proxyHeader"`
	SslPolicy       *string                        `json:"sslPolicy"`
	Project         *string                        `json:"project"`
}

func (r *TargetSslProxy) String() string {
	return dcl.SprintResource(r)
}

// The enum TargetSslProxyProxyHeaderEnum.
type TargetSslProxyProxyHeaderEnum string

// TargetSslProxyProxyHeaderEnumRef returns a *TargetSslProxyProxyHeaderEnum with the value of string s
// If the empty string is provided, nil is returned.
func TargetSslProxyProxyHeaderEnumRef(s string) *TargetSslProxyProxyHeaderEnum {
	if s == "" {
		return nil
	}

	v := TargetSslProxyProxyHeaderEnum(s)
	return &v
}

func (v TargetSslProxyProxyHeaderEnum) Validate() error {
	for _, s := range []string{"NONE", "PROXY_V1"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TargetSslProxyProxyHeaderEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *TargetSslProxy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "TargetSslProxy",
		Version: "beta",
	}
}

const TargetSslProxyMaxPage = -1

type TargetSslProxyList struct {
	Items []*TargetSslProxy

	nextToken string

	pageSize int32

	project string
}

func (l *TargetSslProxyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TargetSslProxyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTargetSslProxy(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTargetSslProxy(ctx context.Context, project string) (*TargetSslProxyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTargetSslProxyWithMaxResults(ctx, project, TargetSslProxyMaxPage)

}

func (c *Client) ListTargetSslProxyWithMaxResults(ctx context.Context, project string, pageSize int32) (*TargetSslProxyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listTargetSslProxy(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TargetSslProxyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *TargetSslProxy) URLNormalized() *TargetSslProxy {
	normalized := dcl.Copy(*r).(TargetSslProxy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Service = dcl.SelfLinkToName(r.Service)
	normalized.SslPolicy = dcl.SelfLinkToName(r.SslPolicy)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetTargetSslProxy(ctx context.Context, r *TargetSslProxy) (*TargetSslProxy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getTargetSslProxyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTargetSslProxy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name
	if dcl.IsZeroValue(result.ProxyHeader) {
		result.ProxyHeader = TargetSslProxyProxyHeaderEnumRef("NONE")
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTargetSslProxyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTargetSslProxy(ctx context.Context, r *TargetSslProxy) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("TargetSslProxy resource is nil")
	}
	c.Config.Logger.Info("Deleting TargetSslProxy...")
	deleteOp := deleteTargetSslProxyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTargetSslProxy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTargetSslProxy(ctx context.Context, project string, filter func(*TargetSslProxy) bool) error {
	listObj, err := c.ListTargetSslProxy(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllTargetSslProxy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTargetSslProxy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTargetSslProxy(ctx context.Context, rawDesired *TargetSslProxy, opts ...dcl.ApplyOption) (*TargetSslProxy, error) {
	var resultNewState *TargetSslProxy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTargetSslProxyHelper(c, ctx, rawDesired, opts...)
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

func applyTargetSslProxyHelper(c *Client, ctx context.Context, rawDesired *TargetSslProxy, opts ...dcl.ApplyOption) (*TargetSslProxy, error) {
	c.Config.Logger.Info("Beginning ApplyTargetSslProxy...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.targetSslProxyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTargetSslProxyDiffs(c.Config, fieldDiffs, opts)
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
	var ops []targetSslProxyApiOperation
	if create {
		ops = append(ops, &createTargetSslProxyOperation{})
	} else if recreate {
		ops = append(ops, &deleteTargetSslProxyOperation{})
		ops = append(ops, &createTargetSslProxyOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeTargetSslProxyDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetTargetSslProxy(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTargetSslProxyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTargetSslProxy(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTargetSslProxyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTargetSslProxyNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTargetSslProxyDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTargetSslProxy(c, newDesired, newState)
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
