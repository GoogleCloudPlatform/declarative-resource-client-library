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
package monitoring

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type NotificationChannel struct {
	Description        *string                                    `json:"description"`
	DisplayName        *string                                    `json:"displayName"`
	Enabled            *bool                                      `json:"enabled"`
	Labels             map[string]string                          `json:"labels"`
	Name               *string                                    `json:"name"`
	Type               *string                                    `json:"type"`
	UserLabels         map[string]string                          `json:"userLabels"`
	VerificationStatus *NotificationChannelVerificationStatusEnum `json:"verificationStatus"`
	Project            *string                                    `json:"project"`
}

func (r *NotificationChannel) String() string {
	return dcl.SprintResource(r)
}

// The enum NotificationChannelVerificationStatusEnum.
type NotificationChannelVerificationStatusEnum string

// NotificationChannelVerificationStatusEnumRef returns a *NotificationChannelVerificationStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func NotificationChannelVerificationStatusEnumRef(s string) *NotificationChannelVerificationStatusEnum {
	if s == "" {
		return nil
	}

	v := NotificationChannelVerificationStatusEnum(s)
	return &v
}

func (v NotificationChannelVerificationStatusEnum) Validate() error {
	for _, s := range []string{"VERIFICATION_STATUS_UNSPECIFIED", "UNVERIFIED", "VERIFIED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NotificationChannelVerificationStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *NotificationChannel) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "monitoring",
		Type:    "NotificationChannel",
		Version: "monitoring",
	}
}

const NotificationChannelMaxPage = -1

type NotificationChannelList struct {
	Items []*NotificationChannel

	nextToken string

	pageSize int32

	project string
}

func (l *NotificationChannelList) HasNext() bool {
	return l.nextToken != ""
}

func (l *NotificationChannelList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listNotificationChannel(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListNotificationChannel(ctx context.Context, project string) (*NotificationChannelList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListNotificationChannelWithMaxResults(ctx, project, NotificationChannelMaxPage)

}

func (c *Client) ListNotificationChannelWithMaxResults(ctx context.Context, project string, pageSize int32) (*NotificationChannelList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listNotificationChannel(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &NotificationChannelList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetNotificationChannel(ctx context.Context, r *NotificationChannel) (*NotificationChannel, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getNotificationChannelRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalNotificationChannel(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name
	if dcl.IsZeroValue(result.Enabled) {
		result.Enabled = dcl.Bool(true)
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeNotificationChannelNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteNotificationChannel(ctx context.Context, r *NotificationChannel) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("NotificationChannel resource is nil")
	}
	c.Config.Logger.Info("Deleting NotificationChannel...")
	deleteOp := deleteNotificationChannelOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllNotificationChannel deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllNotificationChannel(ctx context.Context, project string, filter func(*NotificationChannel) bool) error {
	listObj, err := c.ListNotificationChannel(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllNotificationChannel(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllNotificationChannel(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyNotificationChannel(ctx context.Context, rawDesired *NotificationChannel, opts ...dcl.ApplyOption) (*NotificationChannel, error) {

	var resultNewState *NotificationChannel
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyNotificationChannelHelper(c, ctx, rawDesired, opts...)
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

func applyNotificationChannelHelper(c *Client, ctx context.Context, rawDesired *NotificationChannel, opts ...dcl.ApplyOption) (*NotificationChannel, error) {
	c.Config.Logger.Info("Beginning ApplyNotificationChannel...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.notificationChannelDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToNotificationChannelOp(opStrings, fieldDiffs, opts)
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
	var ops []notificationChannelApiOperation
	if create {
		ops = append(ops, &createNotificationChannelOperation{})
	} else if recreate {

		ops = append(ops, &deleteNotificationChannelOperation{})

		ops = append(ops, &createNotificationChannelOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeNotificationChannelDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetNotificationChannel(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createNotificationChannelOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapNotificationChannel(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeNotificationChannelNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeNotificationChannelNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeNotificationChannelDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffNotificationChannel(c, newDesired, newState)
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
