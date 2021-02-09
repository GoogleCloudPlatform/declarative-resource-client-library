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
package pubsub

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Subscription struct {
	Name                     *string                       `json:"name"`
	Topic                    *string                       `json:"topic"`
	Labels                   map[string]string             `json:"labels"`
	MessageRetentionDuration *string                       `json:"messageRetentionDuration"`
	RetainAckedMessages      *bool                         `json:"retainAckedMessages"`
	ExpirationPolicy         *SubscriptionExpirationPolicy `json:"expirationPolicy"`
	Project                  *string                       `json:"project"`
	DeadLetterPolicy         *SubscriptionDeadLetterPolicy `json:"deadLetterPolicy"`
	PushConfig               *SubscriptionPushConfig       `json:"pushConfig"`
	AckDeadlineSeconds       *int64                        `json:"ackDeadlineSeconds"`
}

func (r *Subscription) String() string {
	return dcl.SprintResource(r)
}

type SubscriptionExpirationPolicy struct {
	empty bool    `json:"-"`
	Ttl   *string `json:"ttl"`
}

// This object is used to assert a desired state where this SubscriptionExpirationPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySubscriptionExpirationPolicy *SubscriptionExpirationPolicy = &SubscriptionExpirationPolicy{empty: true}

func (r *SubscriptionExpirationPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *SubscriptionExpirationPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type SubscriptionDeadLetterPolicy struct {
	empty               bool    `json:"-"`
	DeadLetterTopic     *string `json:"deadLetterTopic"`
	MaxDeliveryAttempts *int64  `json:"maxDeliveryAttempts"`
}

// This object is used to assert a desired state where this SubscriptionDeadLetterPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySubscriptionDeadLetterPolicy *SubscriptionDeadLetterPolicy = &SubscriptionDeadLetterPolicy{empty: true}

func (r *SubscriptionDeadLetterPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *SubscriptionDeadLetterPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type SubscriptionPushConfig struct {
	empty        bool                             `json:"-"`
	PushEndpoint *string                          `json:"pushEndpoint"`
	Attributes   map[string]string                `json:"attributes"`
	OidcToken    *SubscriptionPushConfigOidcToken `json:"oidcToken"`
}

// This object is used to assert a desired state where this SubscriptionPushConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySubscriptionPushConfig *SubscriptionPushConfig = &SubscriptionPushConfig{empty: true}

func (r *SubscriptionPushConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *SubscriptionPushConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type SubscriptionPushConfigOidcToken struct {
	empty               bool    `json:"-"`
	ServiceAccountEmail *string `json:"serviceAccountEmail"`
	Audience            *string `json:"audience"`
}

// This object is used to assert a desired state where this SubscriptionPushConfigOidcToken is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySubscriptionPushConfigOidcToken *SubscriptionPushConfigOidcToken = &SubscriptionPushConfigOidcToken{empty: true}

func (r *SubscriptionPushConfigOidcToken) String() string {
	return dcl.SprintResource(r)
}

func (r *SubscriptionPushConfigOidcToken) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Subscription) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "pubsub",
		Type:    "Subscription",
		Version: "pubsub",
	}
}

const SubscriptionMaxPage = -1

type SubscriptionList struct {
	Items []*Subscription

	nextToken string

	pageSize int32

	project string
}

func (l *SubscriptionList) HasNext() bool {
	return l.nextToken != ""
}

func (l *SubscriptionList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listSubscription(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListSubscription(ctx context.Context, project string) (*SubscriptionList, error) {

	return c.ListSubscriptionWithMaxResults(ctx, project, SubscriptionMaxPage)

}

func (c *Client) ListSubscriptionWithMaxResults(ctx context.Context, project string, pageSize int32) (*SubscriptionList, error) {
	items, token, err := c.listSubscription(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &SubscriptionList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetSubscription(ctx context.Context, r *Subscription) (*Subscription, error) {
	b, err := c.getSubscriptionRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalSubscription(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeSubscriptionNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteSubscription(ctx context.Context, r *Subscription) error {
	if r == nil {
		return fmt.Errorf("Subscription resource is nil")
	}
	c.Config.Logger.Info("Deleting Subscription...")
	deleteOp := deleteSubscriptionOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllSubscription deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllSubscription(ctx context.Context, project string, filter func(*Subscription) bool) error {
	listObj, err := c.ListSubscription(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllSubscription(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllSubscription(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplySubscription(ctx context.Context, rawDesired *Subscription, opts ...dcl.ApplyOption) (*Subscription, error) {
	c.Config.Logger.Info("Beginning ApplySubscription...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.subscriptionDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []subscriptionApiOperation
	if create {
		ops = append(ops, &createSubscriptionOperation{})
	} else if recreate {

		ops = append(ops, &deleteSubscriptionOperation{})

		ops = append(ops, &createSubscriptionOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeSubscriptionDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetSubscription(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeSubscriptionNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeSubscriptionDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffSubscription(c, newDesired, newState)
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
