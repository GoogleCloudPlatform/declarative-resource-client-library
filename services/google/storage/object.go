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
	"encoding/json"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Object struct {
	Name                    *string                   `json:"name"`
	Bucket                  *string                   `json:"bucket"`
	Generation              *int64                    `json:"generation"`
	Metageneration          *int64                    `json:"metageneration"`
	Id                      *string                   `json:"id"`
	SelfLink                *string                   `json:"selfLink"`
	ContentType             *string                   `json:"contentType"`
	TimeCreated             *string                   `json:"timeCreated"`
	Updated                 *string                   `json:"updated"`
	CustomTime              *string                   `json:"customTime"`
	TimeDeleted             *string                   `json:"timeDeleted"`
	TemporaryHold           *bool                     `json:"temporaryHold"`
	EventBasedHold          *bool                     `json:"eventBasedHold"`
	RetentionExpirationTime *string                   `json:"retentionExpirationTime"`
	StorageClass            *string                   `json:"storageClass"`
	TimeStorageClassUpdated *string                   `json:"timeStorageClassUpdated"`
	Size                    *int64                    `json:"size"`
	Md5Hash                 *string                   `json:"md5Hash"`
	MediaLink               *string                   `json:"mediaLink"`
	Metadata                map[string]string         `json:"metadata"`
	Owner                   *ObjectOwner              `json:"owner"`
	Crc32c                  *string                   `json:"crc32c"`
	ComponentCount          *int64                    `json:"componentCount"`
	Etag                    *string                   `json:"etag"`
	CustomerEncryption      *ObjectCustomerEncryption `json:"customerEncryption"`
	KmsKeyName              *string                   `json:"kmsKeyName"`
	Content                 *string                   `json:"content"`
}

func (r *Object) String() string {
	return dcl.SprintResource(r)
}

type ObjectOwner struct {
	empty    bool    `json:"-"`
	Entity   *string `json:"entity"`
	EntityId *string `json:"entityId"`
}

type jsonObjectOwner ObjectOwner

func (r *ObjectOwner) UnmarshalJSON(data []byte) error {
	var res jsonObjectOwner
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyObjectOwner
	} else {

		r.Entity = res.Entity

		r.EntityId = res.EntityId

	}
	return nil
}

// This object is used to assert a desired state where this ObjectOwner is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyObjectOwner *ObjectOwner = &ObjectOwner{empty: true}

func (r *ObjectOwner) Empty() bool {
	return r.empty
}

func (r *ObjectOwner) String() string {
	return dcl.SprintResource(r)
}

func (r *ObjectOwner) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ObjectCustomerEncryption struct {
	empty               bool    `json:"-"`
	EncryptionAlgorithm *string `json:"encryptionAlgorithm"`
	KeySha256           *string `json:"keySha256"`
	Key                 *string `json:"key"`
}

type jsonObjectCustomerEncryption ObjectCustomerEncryption

func (r *ObjectCustomerEncryption) UnmarshalJSON(data []byte) error {
	var res jsonObjectCustomerEncryption
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyObjectCustomerEncryption
	} else {

		r.EncryptionAlgorithm = res.EncryptionAlgorithm

		r.KeySha256 = res.KeySha256

		r.Key = res.Key

	}
	return nil
}

// This object is used to assert a desired state where this ObjectCustomerEncryption is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyObjectCustomerEncryption *ObjectCustomerEncryption = &ObjectCustomerEncryption{empty: true}

func (r *ObjectCustomerEncryption) Empty() bool {
	return r.empty
}

func (r *ObjectCustomerEncryption) String() string {
	return dcl.SprintResource(r)
}

func (r *ObjectCustomerEncryption) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Object) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "storage",
		Type:    "Object",
		Version: "storage",
	}
}

const ObjectMaxPage = -1

type ObjectList struct {
	Items []*Object

	nextToken string

	pageSize int32

	bucket string
}

func (l *ObjectList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ObjectList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listObject(ctx, l.bucket, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListObject(ctx context.Context, bucket string) (*ObjectList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListObjectWithMaxResults(ctx, bucket, ObjectMaxPage)

}

func (c *Client) ListObjectWithMaxResults(ctx context.Context, bucket string, pageSize int32) (*ObjectList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listObject(ctx, bucket, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ObjectList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		bucket: bucket,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Object) URLNormalized() *Object {
	normalized := dcl.Copy(*r).(Object)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Bucket = dcl.SelfLinkToName(r.Bucket)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.ContentType = dcl.SelfLinkToName(r.ContentType)
	normalized.StorageClass = dcl.SelfLinkToName(r.StorageClass)
	normalized.Md5Hash = dcl.SelfLinkToName(r.Md5Hash)
	normalized.MediaLink = dcl.SelfLinkToName(r.MediaLink)
	normalized.Crc32c = dcl.SelfLinkToName(r.Crc32c)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.KmsKeyName = dcl.SelfLinkToName(r.KmsKeyName)
	normalized.Content = dcl.SelfLinkToName(r.Content)
	return &normalized
}

func (c *Client) DeleteObject(ctx context.Context, r *Object) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Object resource is nil")
	}
	c.Config.Logger.Info("Deleting Object...")
	deleteOp := deleteObjectOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllObject deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllObject(ctx context.Context, bucket string, filter func(*Object) bool) error {
	listObj, err := c.ListObject(ctx, bucket)
	if err != nil {
		return err
	}

	err = c.deleteAllObject(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllObject(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyObject(ctx context.Context, rawDesired *Object, opts ...dcl.ApplyOption) (*Object, error) {

	var resultNewState *Object
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyObjectHelper(c, ctx, rawDesired, opts...)
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

func applyObjectHelper(c *Client, ctx context.Context, rawDesired *Object, opts ...dcl.ApplyOption) (*Object, error) {
	c.Config.Logger.Info("Beginning ApplyObject...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.objectDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToObjectDiffs(c.Config, fieldDiffs, opts)
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
	var ops []objectApiOperation
	if create {
		ops = append(ops, &createObjectOperation{})
	} else if recreate {
		ops = append(ops, &deleteObjectOperation{})
		ops = append(ops, &createObjectOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeObjectDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetObject(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createObjectOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapObject(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeObjectNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeObjectNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeObjectDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffObject(c, newDesired, newState)
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
