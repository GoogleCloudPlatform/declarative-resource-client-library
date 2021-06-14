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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Object) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "bucket"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Owner) {
		if err := r.Owner.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CustomerEncryption) {
		if err := r.CustomerEncryption.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ObjectOwner) validate() error {
	return nil
}
func (r *ObjectCustomerEncryption) validate() error {
	return nil
}

func objectGetURL(userBasePath string, r *Object) (string, error) {
	params := map[string]interface{}{
		"bucket": dcl.ValueOrEmptyString(r.Bucket),
		"name":   dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("b/{{bucket}}/o/{{name}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

func objectListURL(userBasePath, bucket string) (string, error) {
	params := map[string]interface{}{
		"bucket": bucket,
	}
	return dcl.URL("b/{{bucket}}/o", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil

}

func objectDeleteURL(userBasePath string, r *Object) (string, error) {
	params := map[string]interface{}{
		"bucket": dcl.ValueOrEmptyString(r.Bucket),
		"name":   dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("b/{{bucket}}/o/{{name}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

// objectApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type objectApiOperation interface {
	do(context.Context, *Object, *Client) error
}

// newUpdateObjectPatchObjectRequest creates a request for an
// Object resource's PatchObject update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateObjectPatchObjectRequest(ctx context.Context, f *Object, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateObjectPatchObjectRequest converts the update into
// the final JSON request body.
func marshalUpdateObjectPatchObjectRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"customer_encryption"},
		[]string{},
	)
	return json.Marshal(m)
}

type updateObjectPatchObjectOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateObjectPatchObjectOperation) do(ctx context.Context, r *Object, c *Client) error {
	_, err := c.GetObject(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchObject")
	if err != nil {
		return err
	}

	req, err := newUpdateObjectPatchObjectRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateObjectPatchObjectRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/storage/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listObjectRaw(ctx context.Context, bucket, pageToken string, pageSize int32) ([]byte, error) {
	u, err := objectListURL(c.Config.BasePath, bucket)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ObjectMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listObjectOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listObject(ctx context.Context, bucket, pageToken string, pageSize int32) ([]*Object, string, error) {
	b, err := c.listObjectRaw(ctx, bucket, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listObjectOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Object
	for _, v := range m.Items {
		res, err := unmarshalMapObject(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Bucket = &bucket
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllObject(ctx context.Context, f func(*Object) bool, resources []*Object) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteObject(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteObjectOperation struct{}

func (op *deleteObjectOperation) do(ctx context.Context, r *Object, c *Client) error {

	_, err := c.GetObject(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Object not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetObject checking for existence. error: %v", err)
		return err
	}

	u, err := objectDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Object: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetObject(ctx, r.urlNormalized())
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createObjectOperation struct {
	response map[string]interface{}
}

func (op *createObjectOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) objectDiffsForRawDesired(ctx context.Context, rawDesired *Object, opts ...dcl.ApplyOption) (initial, desired *Object, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Object
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Object); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Object, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetObject(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Object resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Object resource: %v", err)
		}
		c.Config.Logger.Info("Found that Object resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeObjectDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Object: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Object: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeObjectInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Object: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeObjectDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Object: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffObject(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeObjectInitialState(rawInitial, rawDesired *Object) (*Object, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeObjectDesiredState(rawDesired, rawInitial *Object, opts ...dcl.ApplyOption) (*Object, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Owner = canonicalizeObjectOwner(rawDesired.Owner, nil, opts...)
		rawDesired.CustomerEncryption = canonicalizeObjectCustomerEncryption(rawDesired.CustomerEncryption, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Bucket, rawInitial.Bucket) {
		rawDesired.Bucket = rawInitial.Bucket
	}
	if dcl.StringCanonicalize(rawDesired.ContentType, rawInitial.ContentType) {
		rawDesired.ContentType = rawInitial.ContentType
	}
	if dcl.IsZeroValue(rawDesired.CustomTime) {
		rawDesired.CustomTime = rawInitial.CustomTime
	}
	if dcl.BoolCanonicalize(rawDesired.TemporaryHold, rawInitial.TemporaryHold) {
		rawDesired.TemporaryHold = rawInitial.TemporaryHold
	}
	if dcl.BoolCanonicalize(rawDesired.EventBasedHold, rawInitial.EventBasedHold) {
		rawDesired.EventBasedHold = rawInitial.EventBasedHold
	}
	if dcl.StringCanonicalize(rawDesired.StorageClass, rawInitial.StorageClass) {
		rawDesired.StorageClass = rawInitial.StorageClass
	}
	if dcl.StringCanonicalize(rawDesired.Md5Hash, rawInitial.Md5Hash) {
		rawDesired.Md5Hash = rawInitial.Md5Hash
	}
	if dcl.IsZeroValue(rawDesired.Metadata) {
		rawDesired.Metadata = rawInitial.Metadata
	}
	if dcl.StringCanonicalize(rawDesired.Crc32c, rawInitial.Crc32c) {
		rawDesired.Crc32c = rawInitial.Crc32c
	}
	rawDesired.CustomerEncryption = canonicalizeObjectCustomerEncryption(rawDesired.CustomerEncryption, rawInitial.CustomerEncryption, opts...)
	if dcl.NameToSelfLink(rawDesired.KmsKeyName, rawInitial.KmsKeyName) {
		rawDesired.KmsKeyName = rawInitial.KmsKeyName
	}
	if dcl.StringCanonicalize(rawDesired.Content, rawInitial.Content) {
		rawDesired.Content = rawInitial.Content
	}

	return rawDesired, nil
}

func canonicalizeObjectNewState(c *Client, rawNew, rawDesired *Object) (*Object, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Bucket) && dcl.IsEmptyValueIndirect(rawDesired.Bucket) {
		rawNew.Bucket = rawDesired.Bucket
	} else {
		if dcl.NameToSelfLink(rawDesired.Bucket, rawNew.Bucket) {
			rawNew.Bucket = rawDesired.Bucket
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Generation) && dcl.IsEmptyValueIndirect(rawDesired.Generation) {
		rawNew.Generation = rawDesired.Generation
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Metageneration) && dcl.IsEmptyValueIndirect(rawDesired.Metageneration) {
		rawNew.Metageneration = rawDesired.Metageneration
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ContentType) && dcl.IsEmptyValueIndirect(rawDesired.ContentType) {
		rawNew.ContentType = rawDesired.ContentType
	} else {
		if dcl.StringCanonicalize(rawDesired.ContentType, rawNew.ContentType) {
			rawNew.ContentType = rawDesired.ContentType
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeCreated) && dcl.IsEmptyValueIndirect(rawDesired.TimeCreated) {
		rawNew.TimeCreated = rawDesired.TimeCreated
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Updated) && dcl.IsEmptyValueIndirect(rawDesired.Updated) {
		rawNew.Updated = rawDesired.Updated
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CustomTime) && dcl.IsEmptyValueIndirect(rawDesired.CustomTime) {
		rawNew.CustomTime = rawDesired.CustomTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeDeleted) && dcl.IsEmptyValueIndirect(rawDesired.TimeDeleted) {
		rawNew.TimeDeleted = rawDesired.TimeDeleted
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TemporaryHold) && dcl.IsEmptyValueIndirect(rawDesired.TemporaryHold) {
		rawNew.TemporaryHold = rawDesired.TemporaryHold
	} else {
		if dcl.BoolCanonicalize(rawDesired.TemporaryHold, rawNew.TemporaryHold) {
			rawNew.TemporaryHold = rawDesired.TemporaryHold
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EventBasedHold) && dcl.IsEmptyValueIndirect(rawDesired.EventBasedHold) {
		rawNew.EventBasedHold = rawDesired.EventBasedHold
	} else {
		if dcl.BoolCanonicalize(rawDesired.EventBasedHold, rawNew.EventBasedHold) {
			rawNew.EventBasedHold = rawDesired.EventBasedHold
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RetentionExpirationTime) && dcl.IsEmptyValueIndirect(rawDesired.RetentionExpirationTime) {
		rawNew.RetentionExpirationTime = rawDesired.RetentionExpirationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StorageClass) && dcl.IsEmptyValueIndirect(rawDesired.StorageClass) {
		rawNew.StorageClass = rawDesired.StorageClass
	} else {
		if dcl.StringCanonicalize(rawDesired.StorageClass, rawNew.StorageClass) {
			rawNew.StorageClass = rawDesired.StorageClass
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeStorageClassUpdated) && dcl.IsEmptyValueIndirect(rawDesired.TimeStorageClassUpdated) {
		rawNew.TimeStorageClassUpdated = rawDesired.TimeStorageClassUpdated
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Size) && dcl.IsEmptyValueIndirect(rawDesired.Size) {
		rawNew.Size = rawDesired.Size
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Md5Hash) && dcl.IsEmptyValueIndirect(rawDesired.Md5Hash) {
		rawNew.Md5Hash = rawDesired.Md5Hash
	} else {
		if dcl.StringCanonicalize(rawDesired.Md5Hash, rawNew.Md5Hash) {
			rawNew.Md5Hash = rawDesired.Md5Hash
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MediaLink) && dcl.IsEmptyValueIndirect(rawDesired.MediaLink) {
		rawNew.MediaLink = rawDesired.MediaLink
	} else {
		if dcl.StringCanonicalize(rawDesired.MediaLink, rawNew.MediaLink) {
			rawNew.MediaLink = rawDesired.MediaLink
		}
	}

	rawNew.Metadata = rawDesired.Metadata

	if dcl.IsEmptyValueIndirect(rawNew.Owner) && dcl.IsEmptyValueIndirect(rawDesired.Owner) {
		rawNew.Owner = rawDesired.Owner
	} else {
		rawNew.Owner = canonicalizeNewObjectOwner(c, rawDesired.Owner, rawNew.Owner)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Crc32c) && dcl.IsEmptyValueIndirect(rawDesired.Crc32c) {
		rawNew.Crc32c = rawDesired.Crc32c
	} else {
		if dcl.StringCanonicalize(rawDesired.Crc32c, rawNew.Crc32c) {
			rawNew.Crc32c = rawDesired.Crc32c
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ComponentCount) && dcl.IsEmptyValueIndirect(rawDesired.ComponentCount) {
		rawNew.ComponentCount = rawDesired.ComponentCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CustomerEncryption) && dcl.IsEmptyValueIndirect(rawDesired.CustomerEncryption) {
		rawNew.CustomerEncryption = rawDesired.CustomerEncryption
	} else {
		rawNew.CustomerEncryption = canonicalizeNewObjectCustomerEncryption(c, rawDesired.CustomerEncryption, rawNew.CustomerEncryption)
	}

	if dcl.IsEmptyValueIndirect(rawNew.KmsKeyName) && dcl.IsEmptyValueIndirect(rawDesired.KmsKeyName) {
		rawNew.KmsKeyName = rawDesired.KmsKeyName
	} else {
		if dcl.NameToSelfLink(rawDesired.KmsKeyName, rawNew.KmsKeyName) {
			rawNew.KmsKeyName = rawDesired.KmsKeyName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Content) && dcl.IsEmptyValueIndirect(rawDesired.Content) {
		rawNew.Content = rawDesired.Content
	} else {
		if dcl.StringCanonicalize(rawDesired.Content, rawNew.Content) {
			rawNew.Content = rawDesired.Content
		}
	}

	return rawNew, nil
}

func canonicalizeObjectOwner(des, initial *ObjectOwner, opts ...dcl.ApplyOption) *ObjectOwner {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	return des
}

func canonicalizeNewObjectOwner(c *Client, des, nw *ObjectOwner) *ObjectOwner {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Entity, nw.Entity) {
		nw.Entity = des.Entity
	}
	if dcl.StringCanonicalize(des.EntityId, nw.EntityId) {
		nw.EntityId = des.EntityId
	}

	return nw
}

func canonicalizeNewObjectOwnerSet(c *Client, des, nw []ObjectOwner) []ObjectOwner {
	if des == nil {
		return nw
	}
	var reorderedNew []ObjectOwner
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareObjectOwnerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewObjectOwnerSlice(c *Client, des, nw []ObjectOwner) []ObjectOwner {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ObjectOwner
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewObjectOwner(c, &d, &n))
	}

	return items
}

func canonicalizeObjectCustomerEncryption(des, initial *ObjectCustomerEncryption, opts ...dcl.ApplyOption) *ObjectCustomerEncryption {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.EncryptionAlgorithm, initial.EncryptionAlgorithm) || dcl.IsZeroValue(des.EncryptionAlgorithm) {
		des.EncryptionAlgorithm = initial.EncryptionAlgorithm
	}
	if dcl.StringCanonicalize(des.KeySha256, initial.KeySha256) || dcl.IsZeroValue(des.KeySha256) {
		des.KeySha256 = initial.KeySha256
	}
	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}

	return des
}

func canonicalizeNewObjectCustomerEncryption(c *Client, des, nw *ObjectCustomerEncryption) *ObjectCustomerEncryption {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.EncryptionAlgorithm, nw.EncryptionAlgorithm) {
		nw.EncryptionAlgorithm = des.EncryptionAlgorithm
	}
	if dcl.StringCanonicalize(des.KeySha256, nw.KeySha256) {
		nw.KeySha256 = des.KeySha256
	}
	nw.Key = des.Key

	return nw
}

func canonicalizeNewObjectCustomerEncryptionSet(c *Client, des, nw []ObjectCustomerEncryption) []ObjectCustomerEncryption {
	if des == nil {
		return nw
	}
	var reorderedNew []ObjectCustomerEncryption
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareObjectCustomerEncryptionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewObjectCustomerEncryptionSlice(c *Client, des, nw []ObjectCustomerEncryption) []ObjectCustomerEncryption {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ObjectCustomerEncryption
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewObjectCustomerEncryption(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffObject(c *Client, desired, actual *Object, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Bucket, actual.Bucket, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Bucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metageneration, actual.Metageneration, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Metageneration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContentType, actual.ContentType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ContentType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeCreated, actual.TimeCreated, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TimeCreated")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Updated, actual.Updated, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Updated")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomTime, actual.CustomTime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CustomTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeDeleted, actual.TimeDeleted, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TimeDeleted")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TemporaryHold, actual.TemporaryHold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TemporaryHold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EventBasedHold, actual.EventBasedHold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EventBasedHold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RetentionExpirationTime, actual.RetentionExpirationTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RetentionExpirationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StorageClass, actual.StorageClass, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StorageClass")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeStorageClassUpdated, actual.TimeStorageClassUpdated, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TimeStorageClassUpdated")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Size, actual.Size, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Size")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Md5Hash, actual.Md5Hash, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Md5Hash")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MediaLink, actual.MediaLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MediaLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{Ignore: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Owner, actual.Owner, dcl.Info{OutputOnly: true, ObjectFunction: compareObjectOwnerNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Owner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Crc32c, actual.Crc32c, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Crc32c")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ComponentCount, actual.ComponentCount, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ComponentCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomerEncryption, actual.CustomerEncryption, dcl.Info{ObjectFunction: compareObjectCustomerEncryptionNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CustomerEncryption")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareObjectOwnerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ObjectOwner)
	if !ok {
		desiredNotPointer, ok := d.(ObjectOwner)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ObjectOwner or *ObjectOwner", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ObjectOwner)
	if !ok {
		actualNotPointer, ok := a.(ObjectOwner)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ObjectOwner", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Entity, actual.Entity, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Entity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EntityId, actual.EntityId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EntityId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareObjectCustomerEncryptionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ObjectCustomerEncryption)
	if !ok {
		desiredNotPointer, ok := d.(ObjectCustomerEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ObjectCustomerEncryption or *ObjectCustomerEncryption", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ObjectCustomerEncryption)
	if !ok {
		actualNotPointer, ok := a.(ObjectCustomerEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ObjectCustomerEncryption", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EncryptionAlgorithm, actual.EncryptionAlgorithm, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncryptionAlgorithm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeySha256, actual.KeySha256, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeySha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Object) urlNormalized() *Object {
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

func (r *Object) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Name)
}

func (r *Object) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Name)
}

func (r *Object) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Name)
}

func (r *Object) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "PatchObject" {
		fields := map[string]interface{}{
			"bucket": dcl.ValueOrEmptyString(n.Bucket),
			"name":   dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("b/{{bucket}}/o/{{name}}", "https://www.googleapis.com/storage/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Object resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Object) marshal(c *Client) ([]byte, error) {
	m, err := expandObject(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Object: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"customer_encryption"},
		[]string{},
	)

	return json.Marshal(m)
}

// unmarshalObject decodes JSON responses into the Object resource schema.
func unmarshalObject(b []byte, c *Client) (*Object, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapObject(m, c)
}

func unmarshalMapObject(m map[string]interface{}, c *Client) (*Object, error) {

	return flattenObject(c, m), nil
}

// expandObject expands Object into a JSON request object.
func expandObject(c *Client, f *Object) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Bucket; !dcl.IsEmptyValueIndirect(v) {
		m["bucket"] = v
	}
	if v := f.Generation; !dcl.IsEmptyValueIndirect(v) {
		m["generation"] = v
	}
	if v := f.Metageneration; !dcl.IsEmptyValueIndirect(v) {
		m["metageneration"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.ContentType; !dcl.IsEmptyValueIndirect(v) {
		m["contentType"] = v
	}
	if v := f.TimeCreated; !dcl.IsEmptyValueIndirect(v) {
		m["timeCreated"] = v
	}
	if v := f.Updated; !dcl.IsEmptyValueIndirect(v) {
		m["updated"] = v
	}
	if v := f.CustomTime; !dcl.IsEmptyValueIndirect(v) {
		m["customTime"] = v
	}
	if v := f.TimeDeleted; !dcl.IsEmptyValueIndirect(v) {
		m["timeDeleted"] = v
	}
	if v := f.TemporaryHold; !dcl.IsEmptyValueIndirect(v) {
		m["temporaryHold"] = v
	}
	if v := f.EventBasedHold; !dcl.IsEmptyValueIndirect(v) {
		m["eventBasedHold"] = v
	}
	if v := f.RetentionExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		m["retentionExpirationTime"] = v
	}
	if v := f.StorageClass; !dcl.IsEmptyValueIndirect(v) {
		m["storageClass"] = v
	}
	if v := f.TimeStorageClassUpdated; !dcl.IsEmptyValueIndirect(v) {
		m["timeStorageClassUpdated"] = v
	}
	if v := f.Size; !dcl.IsEmptyValueIndirect(v) {
		m["size"] = v
	}
	if v := f.Md5Hash; !dcl.IsEmptyValueIndirect(v) {
		m["md5Hash"] = v
	}
	if v := f.MediaLink; !dcl.IsEmptyValueIndirect(v) {
		m["mediaLink"] = v
	}
	if v := f.Metadata; !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v, err := expandObjectOwner(c, f.Owner); err != nil {
		return nil, fmt.Errorf("error expanding Owner into owner: %w", err)
	} else if v != nil {
		m["owner"] = v
	}
	if v := f.Crc32c; !dcl.IsEmptyValueIndirect(v) {
		m["crc32c"] = v
	}
	if v := f.ComponentCount; !dcl.IsEmptyValueIndirect(v) {
		m["componentCount"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v, err := expandObjectCustomerEncryption(c, f.CustomerEncryption); err != nil {
		return nil, fmt.Errorf("error expanding CustomerEncryption into customerEncryption: %w", err)
	} else if v != nil {
		m["customerEncryption"] = v
	}
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}

	return m, nil
}

// flattenObject flattens Object from a JSON request object into the
// Object type.
func flattenObject(c *Client, i interface{}) *Object {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Object{}
	res.Name = dcl.FlattenString(m["name"])
	res.Bucket = dcl.FlattenString(m["bucket"])
	res.Generation = dcl.FlattenInteger(m["generation"])
	res.Metageneration = dcl.FlattenInteger(m["metageneration"])
	res.Id = dcl.FlattenString(m["id"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.ContentType = dcl.FlattenString(m["contentType"])
	res.TimeCreated = dcl.FlattenString(m["timeCreated"])
	res.Updated = dcl.FlattenString(m["updated"])
	res.CustomTime = dcl.FlattenString(m["customTime"])
	res.TimeDeleted = dcl.FlattenString(m["timeDeleted"])
	res.TemporaryHold = dcl.FlattenBool(m["temporaryHold"])
	res.EventBasedHold = dcl.FlattenBool(m["eventBasedHold"])
	res.RetentionExpirationTime = dcl.FlattenString(m["retentionExpirationTime"])
	res.StorageClass = dcl.FlattenString(m["storageClass"])
	res.TimeStorageClassUpdated = dcl.FlattenString(m["timeStorageClassUpdated"])
	res.Size = dcl.FlattenInteger(m["size"])
	res.Md5Hash = dcl.FlattenString(m["md5Hash"])
	res.MediaLink = dcl.FlattenString(m["mediaLink"])
	res.Metadata = dcl.FlattenKeyValuePairs(m["metadata"])
	res.Owner = flattenObjectOwner(c, m["owner"])
	res.Crc32c = dcl.FlattenString(m["crc32c"])
	res.ComponentCount = dcl.FlattenInteger(m["componentCount"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.CustomerEncryption = flattenObjectCustomerEncryption(c, m["customerEncryption"])
	res.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	res.Content = dcl.FlattenString(m["content"])

	return res
}

// expandObjectOwnerMap expands the contents of ObjectOwner into a JSON
// request object.
func expandObjectOwnerMap(c *Client, f map[string]ObjectOwner) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandObjectOwner(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandObjectOwnerSlice expands the contents of ObjectOwner into a JSON
// request object.
func expandObjectOwnerSlice(c *Client, f []ObjectOwner) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandObjectOwner(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenObjectOwnerMap flattens the contents of ObjectOwner from a JSON
// response object.
func flattenObjectOwnerMap(c *Client, i interface{}) map[string]ObjectOwner {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ObjectOwner{}
	}

	if len(a) == 0 {
		return map[string]ObjectOwner{}
	}

	items := make(map[string]ObjectOwner)
	for k, item := range a {
		items[k] = *flattenObjectOwner(c, item.(map[string]interface{}))
	}

	return items
}

// flattenObjectOwnerSlice flattens the contents of ObjectOwner from a JSON
// response object.
func flattenObjectOwnerSlice(c *Client, i interface{}) []ObjectOwner {
	a, ok := i.([]interface{})
	if !ok {
		return []ObjectOwner{}
	}

	if len(a) == 0 {
		return []ObjectOwner{}
	}

	items := make([]ObjectOwner, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenObjectOwner(c, item.(map[string]interface{})))
	}

	return items
}

// expandObjectOwner expands an instance of ObjectOwner into a JSON
// request object.
func expandObjectOwner(c *Client, f *ObjectOwner) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Entity; !dcl.IsEmptyValueIndirect(v) {
		m["entity"] = v
	}
	if v := f.EntityId; !dcl.IsEmptyValueIndirect(v) {
		m["entityId"] = v
	}

	return m, nil
}

// flattenObjectOwner flattens an instance of ObjectOwner from a JSON
// response object.
func flattenObjectOwner(c *Client, i interface{}) *ObjectOwner {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ObjectOwner{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyObjectOwner
	}
	r.Entity = dcl.FlattenString(m["entity"])
	r.EntityId = dcl.FlattenString(m["entityId"])

	return r
}

// expandObjectCustomerEncryptionMap expands the contents of ObjectCustomerEncryption into a JSON
// request object.
func expandObjectCustomerEncryptionMap(c *Client, f map[string]ObjectCustomerEncryption) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandObjectCustomerEncryption(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandObjectCustomerEncryptionSlice expands the contents of ObjectCustomerEncryption into a JSON
// request object.
func expandObjectCustomerEncryptionSlice(c *Client, f []ObjectCustomerEncryption) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandObjectCustomerEncryption(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenObjectCustomerEncryptionMap flattens the contents of ObjectCustomerEncryption from a JSON
// response object.
func flattenObjectCustomerEncryptionMap(c *Client, i interface{}) map[string]ObjectCustomerEncryption {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ObjectCustomerEncryption{}
	}

	if len(a) == 0 {
		return map[string]ObjectCustomerEncryption{}
	}

	items := make(map[string]ObjectCustomerEncryption)
	for k, item := range a {
		items[k] = *flattenObjectCustomerEncryption(c, item.(map[string]interface{}))
	}

	return items
}

// flattenObjectCustomerEncryptionSlice flattens the contents of ObjectCustomerEncryption from a JSON
// response object.
func flattenObjectCustomerEncryptionSlice(c *Client, i interface{}) []ObjectCustomerEncryption {
	a, ok := i.([]interface{})
	if !ok {
		return []ObjectCustomerEncryption{}
	}

	if len(a) == 0 {
		return []ObjectCustomerEncryption{}
	}

	items := make([]ObjectCustomerEncryption, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenObjectCustomerEncryption(c, item.(map[string]interface{})))
	}

	return items
}

// expandObjectCustomerEncryption expands an instance of ObjectCustomerEncryption into a JSON
// request object.
func expandObjectCustomerEncryption(c *Client, f *ObjectCustomerEncryption) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EncryptionAlgorithm; !dcl.IsEmptyValueIndirect(v) {
		m["encryptionAlgorithm"] = v
	}
	if v := f.KeySha256; !dcl.IsEmptyValueIndirect(v) {
		m["keySha256"] = v
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}

	return m, nil
}

// flattenObjectCustomerEncryption flattens an instance of ObjectCustomerEncryption from a JSON
// response object.
func flattenObjectCustomerEncryption(c *Client, i interface{}) *ObjectCustomerEncryption {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ObjectCustomerEncryption{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyObjectCustomerEncryption
	}
	r.EncryptionAlgorithm = dcl.FlattenString(m["encryptionAlgorithm"])
	r.KeySha256 = dcl.FlattenString(m["keySha256"])
	r.Key = dcl.FlattenString(m["key"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Object) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalObject(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Bucket == nil && ncr.Bucket == nil {
			c.Config.Logger.Info("Both Bucket fields null - considering equal.")
		} else if nr.Bucket == nil || ncr.Bucket == nil {
			c.Config.Logger.Info("Only one Bucket field is null - considering unequal.")
			return false
		} else if *nr.Bucket != *ncr.Bucket {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type objectDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         objectApiOperation
}

func convertFieldDiffToObjectOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]objectDiff, error) {
	var diffs []objectDiff
	for _, op := range ops {
		diff := objectDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToobjectApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToobjectApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (objectApiOperation, error) {
	switch op {

	case "updateObjectPatchObjectOperation":
		return &updateObjectPatchObjectOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
