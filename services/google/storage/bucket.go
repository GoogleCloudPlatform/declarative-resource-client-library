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
package storage

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Bucket struct {
	Project      *string                 `json:"project"`
	Location     *string                 `json:"location"`
	Name         *string                 `json:"name"`
	Cors         []BucketCors            `json:"cors"`
	Lifecycle    *BucketLifecycle        `json:"lifecycle"`
	Logging      *BucketLogging          `json:"logging"`
	StorageClass *BucketStorageClassEnum `json:"storageClass"`
	Versioning   *BucketVersioning       `json:"versioning"`
	Website      *BucketWebsite          `json:"website"`
}

func (r *Bucket) String() string {
	return dcl.SprintResource(r)
}

// The enum BucketLifecycleRuleActionTypeEnum.
type BucketLifecycleRuleActionTypeEnum string

// BucketLifecycleRuleActionTypeEnumRef returns a *BucketLifecycleRuleActionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func BucketLifecycleRuleActionTypeEnumRef(s string) *BucketLifecycleRuleActionTypeEnum {
	if s == "" {
		return nil
	}

	v := BucketLifecycleRuleActionTypeEnum(s)
	return &v
}

func (v BucketLifecycleRuleActionTypeEnum) Validate() error {
	for _, s := range []string{"Delete", "SetStorageClass"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BucketLifecycleRuleActionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BucketLifecycleRuleConditionWithStateEnum.
type BucketLifecycleRuleConditionWithStateEnum string

// BucketLifecycleRuleConditionWithStateEnumRef returns a *BucketLifecycleRuleConditionWithStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func BucketLifecycleRuleConditionWithStateEnumRef(s string) *BucketLifecycleRuleConditionWithStateEnum {
	if s == "" {
		return nil
	}

	v := BucketLifecycleRuleConditionWithStateEnum(s)
	return &v
}

func (v BucketLifecycleRuleConditionWithStateEnum) Validate() error {
	for _, s := range []string{"LIVE", "ARCHIVED", "ANY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BucketLifecycleRuleConditionWithStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BucketStorageClassEnum.
type BucketStorageClassEnum string

// BucketStorageClassEnumRef returns a *BucketStorageClassEnum with the value of string s
// If the empty string is provided, nil is returned.
func BucketStorageClassEnumRef(s string) *BucketStorageClassEnum {
	if s == "" {
		return nil
	}

	v := BucketStorageClassEnum(s)
	return &v
}

func (v BucketStorageClassEnum) Validate() error {
	for _, s := range []string{"MULTI_REGIONAL", "REGIONAL", "STANDARD", "NEARLINE", "COLDLINE", "ARCHIVE", "DURABLE_REDUCED_AVAILABILITY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BucketStorageClassEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type BucketCors struct {
	empty          bool     `json:"-"`
	MaxAgeSeconds  *int64   `json:"maxAgeSeconds"`
	Method         []string `json:"method"`
	Origin         []string `json:"origin"`
	ResponseHeader []string `json:"responseHeader"`
}

// This object is used to assert a desired state where this BucketCors is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBucketCors *BucketCors = &BucketCors{empty: true}

func (r *BucketCors) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketCors) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLifecycle struct {
	empty bool                  `json:"-"`
	Rule  []BucketLifecycleRule `json:"rule"`
}

// This object is used to assert a desired state where this BucketLifecycle is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBucketLifecycle *BucketLifecycle = &BucketLifecycle{empty: true}

func (r *BucketLifecycle) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLifecycle) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLifecycleRule struct {
	empty     bool                          `json:"-"`
	Action    *BucketLifecycleRuleAction    `json:"action"`
	Condition *BucketLifecycleRuleCondition `json:"condition"`
}

// This object is used to assert a desired state where this BucketLifecycleRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBucketLifecycleRule *BucketLifecycleRule = &BucketLifecycleRule{empty: true}

func (r *BucketLifecycleRule) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLifecycleRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLifecycleRuleAction struct {
	empty        bool                               `json:"-"`
	StorageClass *string                            `json:"storageClass"`
	Type         *BucketLifecycleRuleActionTypeEnum `json:"type"`
}

// This object is used to assert a desired state where this BucketLifecycleRuleAction is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBucketLifecycleRuleAction *BucketLifecycleRuleAction = &BucketLifecycleRuleAction{empty: true}

func (r *BucketLifecycleRuleAction) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLifecycleRuleAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLifecycleRuleCondition struct {
	empty               bool                                       `json:"-"`
	Age                 *int64                                     `json:"age"`
	CreatedBefore       *string                                    `json:"createdBefore"`
	WithState           *BucketLifecycleRuleConditionWithStateEnum `json:"withState"`
	MatchesStorageClass []string                                   `json:"matchesStorageClass"`
	NumNewerVersions    *int64                                     `json:"numNewerVersions"`
}

// This object is used to assert a desired state where this BucketLifecycleRuleCondition is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBucketLifecycleRuleCondition *BucketLifecycleRuleCondition = &BucketLifecycleRuleCondition{empty: true}

func (r *BucketLifecycleRuleCondition) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLifecycleRuleCondition) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLogging struct {
	empty           bool    `json:"-"`
	LogBucket       *string `json:"logBucket"`
	LogObjectPrefix *string `json:"logObjectPrefix"`
}

// This object is used to assert a desired state where this BucketLogging is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBucketLogging *BucketLogging = &BucketLogging{empty: true}

func (r *BucketLogging) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLogging) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketVersioning struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this BucketVersioning is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBucketVersioning *BucketVersioning = &BucketVersioning{empty: true}

func (r *BucketVersioning) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketVersioning) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketWebsite struct {
	empty          bool    `json:"-"`
	MainPageSuffix *string `json:"mainPageSuffix"`
	NotFoundPage   *string `json:"notFoundPage"`
}

// This object is used to assert a desired state where this BucketWebsite is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBucketWebsite *BucketWebsite = &BucketWebsite{empty: true}

func (r *BucketWebsite) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketWebsite) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Bucket) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "storage",
		Type:    "Bucket",
		Version: "storage",
	}
}

const BucketMaxPage = -1

type BucketList struct {
	Items []*Bucket

	nextToken string

	pageSize int32

	project string
}

func (l *BucketList) HasNext() bool {
	return l.nextToken != ""
}

func (l *BucketList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listBucket(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListBucket(ctx context.Context, project string) (*BucketList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListBucketWithMaxResults(ctx, project, BucketMaxPage)

}

func (c *Client) ListBucketWithMaxResults(ctx context.Context, project string, pageSize int32) (*BucketList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listBucket(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &BucketList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetBucket(ctx context.Context, r *Bucket) (*Bucket, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getBucketRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalBucket(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeBucketNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteBucket(ctx context.Context, r *Bucket) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Bucket resource is nil")
	}
	c.Config.Logger.Info("Deleting Bucket...")
	deleteOp := deleteBucketOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllBucket deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllBucket(ctx context.Context, project string, filter func(*Bucket) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListBucket(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllBucket(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllBucket(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyBucket(ctx context.Context, rawDesired *Bucket, opts ...dcl.ApplyOption) (*Bucket, error) {
	c.Config.Logger.Info("Beginning ApplyBucket...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.bucketDiffsForRawDesired(ctx, rawDesired, opts...)
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
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []bucketApiOperation
	if create {
		ops = append(ops, &createBucketOperation{})
	} else if recreate {

		ops = append(ops, &deleteBucketOperation{})

		ops = append(ops, &createBucketOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeBucketDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetBucket(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createBucketOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapBucket(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeBucketNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeBucketNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeBucketDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffBucket(c, newDesired, newState)
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
