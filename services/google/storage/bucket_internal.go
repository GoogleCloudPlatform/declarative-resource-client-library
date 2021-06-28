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
)

func (r *Bucket) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "location"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Lifecycle) {
		if err := r.Lifecycle.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Logging) {
		if err := r.Logging.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Versioning) {
		if err := r.Versioning.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Website) {
		if err := r.Website.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BucketCors) validate() error {
	return nil
}
func (r *BucketLifecycle) validate() error {
	return nil
}
func (r *BucketLifecycleRule) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Action) {
		if err := r.Action.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Condition) {
		if err := r.Condition.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BucketLifecycleRuleAction) validate() error {
	return nil
}
func (r *BucketLifecycleRuleCondition) validate() error {
	return nil
}
func (r *BucketLogging) validate() error {
	return nil
}
func (r *BucketVersioning) validate() error {
	return nil
}
func (r *BucketWebsite) validate() error {
	return nil
}

func bucketGetURL(userBasePath string, r *Bucket) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("b/{{name}}?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

func bucketListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("b?project={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil

}

func bucketCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("b?project={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil

}

func bucketDeleteURL(userBasePath string, r *Bucket) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("b/{{name}}?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

func (r *Bucket) SetPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"name": *n.Name,
	}
	return dcl.URL("b/{{name}}/iam", "https://www.googleapis.com/storage/v1/", userBasePath, fields)
}

func (r *Bucket) SetPolicyVerb() string {
	return "PUT"
}

func (r *Bucket) getPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"name": *n.Name,
	}
	return dcl.URL("b/{{name}}/iam", "https://www.googleapis.com/storage/v1/", userBasePath, fields)
}

func (r *Bucket) IAMPolicyVersion() int {
	return 3
}

// bucketApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type bucketApiOperation interface {
	do(context.Context, *Bucket, *Client) error
}

// newUpdateBucketUpdateRequest creates a request for an
// Bucket resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateBucketUpdateRequest(ctx context.Context, f *Bucket, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandBucketCorsSlice(c, f.Cors); err != nil {
		return nil, fmt.Errorf("error expanding Cors into cors: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["cors"] = v
	}
	if v, err := expandBucketLifecycle(c, f.Lifecycle); err != nil {
		return nil, fmt.Errorf("error expanding Lifecycle into lifecycle: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["lifecycle"] = v
	}
	if v, err := expandBucketLogging(c, f.Logging); err != nil {
		return nil, fmt.Errorf("error expanding Logging into logging: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["logging"] = v
	}
	if v := f.StorageClass; !dcl.IsEmptyValueIndirect(v) {
		req["storageClass"] = v
	}
	if v, err := expandBucketVersioning(c, f.Versioning); err != nil {
		return nil, fmt.Errorf("error expanding Versioning into versioning: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["versioning"] = v
	}
	if v, err := expandBucketWebsite(c, f.Website); err != nil {
		return nil, fmt.Errorf("error expanding Website into website: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["website"] = v
	}
	return req, nil
}

// marshalUpdateBucketUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateBucketUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateBucketUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateBucketUpdateOperation) do(ctx context.Context, r *Bucket, c *Client) error {
	_, err := c.GetBucket(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateBucketUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateBucketUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listBucketRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := bucketListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != BucketMaxPage {
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

type listBucketOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listBucket(ctx context.Context, project, pageToken string, pageSize int32) ([]*Bucket, string, error) {
	b, err := c.listBucketRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listBucketOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Bucket
	for _, v := range m.Items {
		res, err := unmarshalMapBucket(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllBucket(ctx context.Context, f func(*Bucket) bool, resources []*Bucket) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteBucket(ctx, res)
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

type deleteBucketOperation struct{}

func (op *deleteBucketOperation) do(ctx context.Context, r *Bucket, c *Client) error {

	_, err := c.GetBucket(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Bucket not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetBucket checking for existence. error: %v", err)
		return err
	}

	u, err := bucketDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Bucket: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetBucket(ctx, r.urlNormalized())
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
type createBucketOperation struct {
	response map[string]interface{}
}

func (op *createBucketOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createBucketOperation) do(ctx context.Context, r *Bucket, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := bucketCreateURL(c.Config.BasePath, project)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetBucket(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getBucketRaw(ctx context.Context, r *Bucket) ([]byte, error) {

	u, err := bucketGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) bucketDiffsForRawDesired(ctx context.Context, rawDesired *Bucket, opts ...dcl.ApplyOption) (initial, desired *Bucket, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Bucket
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Bucket); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Bucket, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetBucket(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Bucket resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Bucket resource: %v", err)
		}
		c.Config.Logger.Info("Found that Bucket resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeBucketDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Bucket: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Bucket: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeBucketInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Bucket: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeBucketDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Bucket: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffBucket(c, desired, initial, opts...)
	fmt.Printf("newDiffs: %v\n", diffs)
	return initial, desired, diffs, err
}

func canonicalizeBucketInitialState(rawInitial, rawDesired *Bucket) (*Bucket, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeBucketDesiredState(rawDesired, rawInitial *Bucket, opts ...dcl.ApplyOption) (*Bucket, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Lifecycle = canonicalizeBucketLifecycle(rawDesired.Lifecycle, nil, opts...)
		rawDesired.Logging = canonicalizeBucketLogging(rawDesired.Logging, nil, opts...)
		rawDesired.Versioning = canonicalizeBucketVersioning(rawDesired.Versioning, nil, opts...)
		rawDesired.Website = canonicalizeBucketWebsite(rawDesired.Website, nil, opts...)

		return rawDesired, nil
	}

	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.StringCanonicalize(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Cors) {
		rawDesired.Cors = rawInitial.Cors
	}
	rawDesired.Lifecycle = canonicalizeBucketLifecycle(rawDesired.Lifecycle, rawInitial.Lifecycle, opts...)
	rawDesired.Logging = canonicalizeBucketLogging(rawDesired.Logging, rawInitial.Logging, opts...)
	if dcl.IsZeroValue(rawDesired.StorageClass) {
		rawDesired.StorageClass = rawInitial.StorageClass
	}
	rawDesired.Versioning = canonicalizeBucketVersioning(rawDesired.Versioning, rawInitial.Versioning, opts...)
	rawDesired.Website = canonicalizeBucketWebsite(rawDesired.Website, rawInitial.Website, opts...)

	return rawDesired, nil
}

func canonicalizeBucketNewState(c *Client, rawNew, rawDesired *Bucket) (*Bucket, error) {

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Location) && dcl.IsEmptyValueIndirect(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
		if dcl.StringCanonicalize(rawDesired.Location, rawNew.Location) {
			rawNew.Location = rawDesired.Location
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Cors) && dcl.IsEmptyValueIndirect(rawDesired.Cors) {
		rawNew.Cors = rawDesired.Cors
	} else {
		rawNew.Cors = canonicalizeNewBucketCorsSlice(c, rawDesired.Cors, rawNew.Cors)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Lifecycle) && dcl.IsEmptyValueIndirect(rawDesired.Lifecycle) {
		rawNew.Lifecycle = rawDesired.Lifecycle
	} else {
		rawNew.Lifecycle = canonicalizeNewBucketLifecycle(c, rawDesired.Lifecycle, rawNew.Lifecycle)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Logging) && dcl.IsEmptyValueIndirect(rawDesired.Logging) {
		rawNew.Logging = rawDesired.Logging
	} else {
		rawNew.Logging = canonicalizeNewBucketLogging(c, rawDesired.Logging, rawNew.Logging)
	}

	if dcl.IsEmptyValueIndirect(rawNew.StorageClass) && dcl.IsEmptyValueIndirect(rawDesired.StorageClass) {
		rawNew.StorageClass = rawDesired.StorageClass
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Versioning) && dcl.IsEmptyValueIndirect(rawDesired.Versioning) {
		rawNew.Versioning = rawDesired.Versioning
	} else {
		rawNew.Versioning = canonicalizeNewBucketVersioning(c, rawDesired.Versioning, rawNew.Versioning)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Website) && dcl.IsEmptyValueIndirect(rawDesired.Website) {
		rawNew.Website = rawDesired.Website
	} else {
		rawNew.Website = canonicalizeNewBucketWebsite(c, rawDesired.Website, rawNew.Website)
	}

	return rawNew, nil
}

func canonicalizeBucketCors(des, initial *BucketCors, opts ...dcl.ApplyOption) *BucketCors {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MaxAgeSeconds) {
		des.MaxAgeSeconds = initial.MaxAgeSeconds
	}
	if dcl.IsZeroValue(des.Method) {
		des.Method = initial.Method
	}
	if dcl.IsZeroValue(des.Origin) {
		des.Origin = initial.Origin
	}
	if dcl.IsZeroValue(des.ResponseHeader) {
		des.ResponseHeader = initial.ResponseHeader
	}

	return des
}

func canonicalizeNewBucketCors(c *Client, des, nw *BucketCors) *BucketCors {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.MaxAgeSeconds) {
		nw.MaxAgeSeconds = des.MaxAgeSeconds
	}
	if dcl.IsZeroValue(nw.Method) {
		nw.Method = des.Method
	}
	if dcl.IsZeroValue(nw.Origin) {
		nw.Origin = des.Origin
	}
	if dcl.IsZeroValue(nw.ResponseHeader) {
		nw.ResponseHeader = des.ResponseHeader
	}

	return nw
}

func canonicalizeNewBucketCorsSet(c *Client, des, nw []BucketCors) []BucketCors {
	if des == nil {
		return nw
	}
	var reorderedNew []BucketCors
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBucketCorsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBucketCorsSlice(c *Client, des, nw []BucketCors) []BucketCors {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketCors
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketCors(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLifecycle(des, initial *BucketLifecycle, opts ...dcl.ApplyOption) *BucketLifecycle {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Rule) {
		des.Rule = initial.Rule
	}

	return des
}

func canonicalizeNewBucketLifecycle(c *Client, des, nw *BucketLifecycle) *BucketLifecycle {
	if des == nil || nw == nil {
		return nw
	}

	nw.Rule = canonicalizeNewBucketLifecycleRuleSlice(c, des.Rule, nw.Rule)

	return nw
}

func canonicalizeNewBucketLifecycleSet(c *Client, des, nw []BucketLifecycle) []BucketLifecycle {
	if des == nil {
		return nw
	}
	var reorderedNew []BucketLifecycle
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBucketLifecycleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBucketLifecycleSlice(c *Client, des, nw []BucketLifecycle) []BucketLifecycle {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLifecycle
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLifecycle(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLifecycleRule(des, initial *BucketLifecycleRule, opts ...dcl.ApplyOption) *BucketLifecycleRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Action = canonicalizeBucketLifecycleRuleAction(des.Action, initial.Action, opts...)
	des.Condition = canonicalizeBucketLifecycleRuleCondition(des.Condition, initial.Condition, opts...)

	return des
}

func canonicalizeNewBucketLifecycleRule(c *Client, des, nw *BucketLifecycleRule) *BucketLifecycleRule {
	if des == nil || nw == nil {
		return nw
	}

	nw.Action = canonicalizeNewBucketLifecycleRuleAction(c, des.Action, nw.Action)
	nw.Condition = canonicalizeNewBucketLifecycleRuleCondition(c, des.Condition, nw.Condition)

	return nw
}

func canonicalizeNewBucketLifecycleRuleSet(c *Client, des, nw []BucketLifecycleRule) []BucketLifecycleRule {
	if des == nil {
		return nw
	}
	var reorderedNew []BucketLifecycleRule
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBucketLifecycleRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBucketLifecycleRuleSlice(c *Client, des, nw []BucketLifecycleRule) []BucketLifecycleRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLifecycleRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLifecycleRule(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLifecycleRuleAction(des, initial *BucketLifecycleRuleAction, opts ...dcl.ApplyOption) *BucketLifecycleRuleAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.StorageClass, initial.StorageClass) || dcl.IsZeroValue(des.StorageClass) {
		des.StorageClass = initial.StorageClass
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewBucketLifecycleRuleAction(c *Client, des, nw *BucketLifecycleRuleAction) *BucketLifecycleRuleAction {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.StorageClass, nw.StorageClass) {
		nw.StorageClass = des.StorageClass
	}
	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
	}

	return nw
}

func canonicalizeNewBucketLifecycleRuleActionSet(c *Client, des, nw []BucketLifecycleRuleAction) []BucketLifecycleRuleAction {
	if des == nil {
		return nw
	}
	var reorderedNew []BucketLifecycleRuleAction
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBucketLifecycleRuleActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBucketLifecycleRuleActionSlice(c *Client, des, nw []BucketLifecycleRuleAction) []BucketLifecycleRuleAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLifecycleRuleAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLifecycleRuleAction(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLifecycleRuleCondition(des, initial *BucketLifecycleRuleCondition, opts ...dcl.ApplyOption) *BucketLifecycleRuleCondition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Age) {
		des.Age = initial.Age
	}
	if dcl.IsZeroValue(des.CreatedBefore) {
		des.CreatedBefore = initial.CreatedBefore
	}
	if dcl.IsZeroValue(des.WithState) {
		des.WithState = initial.WithState
	}
	if dcl.IsZeroValue(des.MatchesStorageClass) {
		des.MatchesStorageClass = initial.MatchesStorageClass
	}
	if dcl.IsZeroValue(des.NumNewerVersions) {
		des.NumNewerVersions = initial.NumNewerVersions
	}

	return des
}

func canonicalizeNewBucketLifecycleRuleCondition(c *Client, des, nw *BucketLifecycleRuleCondition) *BucketLifecycleRuleCondition {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Age) {
		nw.Age = des.Age
	}
	if dcl.IsZeroValue(nw.CreatedBefore) {
		nw.CreatedBefore = des.CreatedBefore
	}
	if dcl.IsZeroValue(nw.WithState) {
		nw.WithState = des.WithState
	}
	if dcl.IsZeroValue(nw.MatchesStorageClass) {
		nw.MatchesStorageClass = des.MatchesStorageClass
	}
	if dcl.IsZeroValue(nw.NumNewerVersions) {
		nw.NumNewerVersions = des.NumNewerVersions
	}

	return nw
}

func canonicalizeNewBucketLifecycleRuleConditionSet(c *Client, des, nw []BucketLifecycleRuleCondition) []BucketLifecycleRuleCondition {
	if des == nil {
		return nw
	}
	var reorderedNew []BucketLifecycleRuleCondition
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBucketLifecycleRuleConditionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBucketLifecycleRuleConditionSlice(c *Client, des, nw []BucketLifecycleRuleCondition) []BucketLifecycleRuleCondition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLifecycleRuleCondition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLifecycleRuleCondition(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLogging(des, initial *BucketLogging, opts ...dcl.ApplyOption) *BucketLogging {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.LogBucket, initial.LogBucket) || dcl.IsZeroValue(des.LogBucket) {
		des.LogBucket = initial.LogBucket
	}
	if dcl.StringCanonicalize(des.LogObjectPrefix, initial.LogObjectPrefix) || dcl.IsZeroValue(des.LogObjectPrefix) {
		des.LogObjectPrefix = initial.LogObjectPrefix
	}

	return des
}

func canonicalizeNewBucketLogging(c *Client, des, nw *BucketLogging) *BucketLogging {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.LogBucket, nw.LogBucket) {
		nw.LogBucket = des.LogBucket
	}
	if dcl.StringCanonicalize(des.LogObjectPrefix, nw.LogObjectPrefix) {
		nw.LogObjectPrefix = des.LogObjectPrefix
	}

	return nw
}

func canonicalizeNewBucketLoggingSet(c *Client, des, nw []BucketLogging) []BucketLogging {
	if des == nil {
		return nw
	}
	var reorderedNew []BucketLogging
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBucketLoggingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBucketLoggingSlice(c *Client, des, nw []BucketLogging) []BucketLogging {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLogging
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLogging(c, &d, &n))
	}

	return items
}

func canonicalizeBucketVersioning(des, initial *BucketVersioning, opts ...dcl.ApplyOption) *BucketVersioning {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewBucketVersioning(c *Client, des, nw *BucketVersioning) *BucketVersioning {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewBucketVersioningSet(c *Client, des, nw []BucketVersioning) []BucketVersioning {
	if des == nil {
		return nw
	}
	var reorderedNew []BucketVersioning
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBucketVersioningNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBucketVersioningSlice(c *Client, des, nw []BucketVersioning) []BucketVersioning {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketVersioning
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketVersioning(c, &d, &n))
	}

	return items
}

func canonicalizeBucketWebsite(des, initial *BucketWebsite, opts ...dcl.ApplyOption) *BucketWebsite {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MainPageSuffix, initial.MainPageSuffix) || dcl.IsZeroValue(des.MainPageSuffix) {
		des.MainPageSuffix = initial.MainPageSuffix
	}
	if dcl.StringCanonicalize(des.NotFoundPage, initial.NotFoundPage) || dcl.IsZeroValue(des.NotFoundPage) {
		des.NotFoundPage = initial.NotFoundPage
	}

	return des
}

func canonicalizeNewBucketWebsite(c *Client, des, nw *BucketWebsite) *BucketWebsite {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MainPageSuffix, nw.MainPageSuffix) {
		nw.MainPageSuffix = des.MainPageSuffix
	}
	if dcl.StringCanonicalize(des.NotFoundPage, nw.NotFoundPage) {
		nw.NotFoundPage = des.NotFoundPage
	}

	return nw
}

func canonicalizeNewBucketWebsiteSet(c *Client, des, nw []BucketWebsite) []BucketWebsite {
	if des == nil {
		return nw
	}
	var reorderedNew []BucketWebsite
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBucketWebsiteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBucketWebsiteSlice(c *Client, des, nw []BucketWebsite) []BucketWebsite {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketWebsite
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketWebsite(c, &d, &n))
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
func diffBucket(c *Client, desired, actual *Bucket, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Cors, actual.Cors, dcl.Info{ObjectFunction: compareBucketCorsNewStyle, EmptyObject: EmptyBucketCors, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Cors")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Lifecycle, actual.Lifecycle, dcl.Info{ObjectFunction: compareBucketLifecycleNewStyle, EmptyObject: EmptyBucketLifecycle, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Lifecycle")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Logging, actual.Logging, dcl.Info{ObjectFunction: compareBucketLoggingNewStyle, EmptyObject: EmptyBucketLogging, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Logging")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StorageClass, actual.StorageClass, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("StorageClass")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Versioning, actual.Versioning, dcl.Info{ObjectFunction: compareBucketVersioningNewStyle, EmptyObject: EmptyBucketVersioning, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Versioning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Website, actual.Website, dcl.Info{ObjectFunction: compareBucketWebsiteNewStyle, EmptyObject: EmptyBucketWebsite, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Website")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareBucketCorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketCors)
	if !ok {
		desiredNotPointer, ok := d.(BucketCors)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketCors or *BucketCors", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketCors)
	if !ok {
		actualNotPointer, ok := a.(BucketCors)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketCors", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxAgeSeconds, actual.MaxAgeSeconds, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("MaxAgeSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Method, actual.Method, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Method")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Origin, actual.Origin, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Origin")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResponseHeader, actual.ResponseHeader, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("ResponseHeader")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLifecycleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLifecycle)
	if !ok {
		desiredNotPointer, ok := d.(BucketLifecycle)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycle or *BucketLifecycle", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLifecycle)
	if !ok {
		actualNotPointer, ok := a.(BucketLifecycle)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycle", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Rule, actual.Rule, dcl.Info{ObjectFunction: compareBucketLifecycleRuleNewStyle, EmptyObject: EmptyBucketLifecycleRule, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Rule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLifecycleRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLifecycleRule)
	if !ok {
		desiredNotPointer, ok := d.(BucketLifecycleRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRule or *BucketLifecycleRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLifecycleRule)
	if !ok {
		actualNotPointer, ok := a.(BucketLifecycleRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Action, actual.Action, dcl.Info{ObjectFunction: compareBucketLifecycleRuleActionNewStyle, EmptyObject: EmptyBucketLifecycleRuleAction, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Action")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Condition, actual.Condition, dcl.Info{ObjectFunction: compareBucketLifecycleRuleConditionNewStyle, EmptyObject: EmptyBucketLifecycleRuleCondition, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Condition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLifecycleRuleActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLifecycleRuleAction)
	if !ok {
		desiredNotPointer, ok := d.(BucketLifecycleRuleAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRuleAction or *BucketLifecycleRuleAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLifecycleRuleAction)
	if !ok {
		actualNotPointer, ok := a.(BucketLifecycleRuleAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRuleAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StorageClass, actual.StorageClass, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("StorageClass")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLifecycleRuleConditionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLifecycleRuleCondition)
	if !ok {
		desiredNotPointer, ok := d.(BucketLifecycleRuleCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRuleCondition or *BucketLifecycleRuleCondition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLifecycleRuleCondition)
	if !ok {
		actualNotPointer, ok := a.(BucketLifecycleRuleCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRuleCondition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Age, actual.Age, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Age")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreatedBefore, actual.CreatedBefore, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("CreatedBefore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WithState, actual.WithState, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("IsLive")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MatchesStorageClass, actual.MatchesStorageClass, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("MatchesStorageClass")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumNewerVersions, actual.NumNewerVersions, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("NumNewerVersions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLoggingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLogging)
	if !ok {
		desiredNotPointer, ok := d.(BucketLogging)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLogging or *BucketLogging", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLogging)
	if !ok {
		actualNotPointer, ok := a.(BucketLogging)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLogging", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LogBucket, actual.LogBucket, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("LogBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LogObjectPrefix, actual.LogObjectPrefix, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("LogObjectPrefix")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketVersioningNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketVersioning)
	if !ok {
		desiredNotPointer, ok := d.(BucketVersioning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketVersioning or *BucketVersioning", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketVersioning)
	if !ok {
		actualNotPointer, ok := a.(BucketVersioning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketVersioning", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketWebsiteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketWebsite)
	if !ok {
		desiredNotPointer, ok := d.(BucketWebsite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketWebsite or *BucketWebsite", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketWebsite)
	if !ok {
		actualNotPointer, ok := a.(BucketWebsite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketWebsite", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MainPageSuffix, actual.MainPageSuffix, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("MainPageSuffix")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NotFoundPage, actual.NotFoundPage, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("NotFoundPage")); len(ds) != 0 || err != nil {
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
func (r *Bucket) urlNormalized() *Bucket {
	normalized := dcl.Copy(*r).(Bucket)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	return &normalized
}

func (r *Bucket) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Bucket) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Bucket) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Bucket) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("b/{{name}}", "https://www.googleapis.com/storage/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Bucket resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Bucket) marshal(c *Client) ([]byte, error) {
	m, err := expandBucket(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Bucket: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalBucket decodes JSON responses into the Bucket resource schema.
func unmarshalBucket(b []byte, c *Client) (*Bucket, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapBucket(m, c)
}

func unmarshalMapBucket(m map[string]interface{}, c *Client) (*Bucket, error) {

	return flattenBucket(c, m), nil
}

// expandBucket expands Bucket into a JSON request object.
func expandBucket(c *Client, f *Bucket) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandBucketCorsSlice(c, f.Cors); err != nil {
		return nil, fmt.Errorf("error expanding Cors into cors: %w", err)
	} else if v != nil {
		m["cors"] = v
	}
	if v, err := expandBucketLifecycle(c, f.Lifecycle); err != nil {
		return nil, fmt.Errorf("error expanding Lifecycle into lifecycle: %w", err)
	} else if v != nil {
		m["lifecycle"] = v
	}
	if v, err := expandBucketLogging(c, f.Logging); err != nil {
		return nil, fmt.Errorf("error expanding Logging into logging: %w", err)
	} else if v != nil {
		m["logging"] = v
	}
	if v := f.StorageClass; !dcl.IsEmptyValueIndirect(v) {
		m["storageClass"] = v
	}
	if v, err := expandBucketVersioning(c, f.Versioning); err != nil {
		return nil, fmt.Errorf("error expanding Versioning into versioning: %w", err)
	} else if v != nil {
		m["versioning"] = v
	}
	if v, err := expandBucketWebsite(c, f.Website); err != nil {
		return nil, fmt.Errorf("error expanding Website into website: %w", err)
	} else if v != nil {
		m["website"] = v
	}

	return m, nil
}

// flattenBucket flattens Bucket from a JSON request object into the
// Bucket type.
func flattenBucket(c *Client, i interface{}) *Bucket {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Bucket{}
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])
	res.Name = dcl.FlattenString(m["name"])
	res.Cors = flattenBucketCorsSlice(c, m["cors"])
	res.Lifecycle = flattenBucketLifecycle(c, m["lifecycle"])
	res.Logging = flattenBucketLogging(c, m["logging"])
	res.StorageClass = flattenBucketStorageClassEnum(m["storageClass"])
	res.Versioning = flattenBucketVersioning(c, m["versioning"])
	res.Website = flattenBucketWebsite(c, m["website"])

	return res
}

// expandBucketCorsMap expands the contents of BucketCors into a JSON
// request object.
func expandBucketCorsMap(c *Client, f map[string]BucketCors) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketCors(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketCorsSlice expands the contents of BucketCors into a JSON
// request object.
func expandBucketCorsSlice(c *Client, f []BucketCors) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketCors(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketCorsMap flattens the contents of BucketCors from a JSON
// response object.
func flattenBucketCorsMap(c *Client, i interface{}) map[string]BucketCors {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketCors{}
	}

	if len(a) == 0 {
		return map[string]BucketCors{}
	}

	items := make(map[string]BucketCors)
	for k, item := range a {
		items[k] = *flattenBucketCors(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBucketCorsSlice flattens the contents of BucketCors from a JSON
// response object.
func flattenBucketCorsSlice(c *Client, i interface{}) []BucketCors {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketCors{}
	}

	if len(a) == 0 {
		return []BucketCors{}
	}

	items := make([]BucketCors, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketCors(c, item.(map[string]interface{})))
	}

	return items
}

// expandBucketCors expands an instance of BucketCors into a JSON
// request object.
func expandBucketCors(c *Client, f *BucketCors) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxAgeSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["maxAgeSeconds"] = v
	}
	if v := f.Method; !dcl.IsEmptyValueIndirect(v) {
		m["method"] = v
	}
	if v := f.Origin; !dcl.IsEmptyValueIndirect(v) {
		m["origin"] = v
	}
	if v := f.ResponseHeader; !dcl.IsEmptyValueIndirect(v) {
		m["responseHeader"] = v
	}

	return m, nil
}

// flattenBucketCors flattens an instance of BucketCors from a JSON
// response object.
func flattenBucketCors(c *Client, i interface{}) *BucketCors {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketCors{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketCors
	}
	r.MaxAgeSeconds = dcl.FlattenInteger(m["maxAgeSeconds"])
	r.Method = dcl.FlattenStringSlice(m["method"])
	r.Origin = dcl.FlattenStringSlice(m["origin"])
	r.ResponseHeader = dcl.FlattenStringSlice(m["responseHeader"])

	return r
}

// expandBucketLifecycleMap expands the contents of BucketLifecycle into a JSON
// request object.
func expandBucketLifecycleMap(c *Client, f map[string]BucketLifecycle) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLifecycle(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLifecycleSlice expands the contents of BucketLifecycle into a JSON
// request object.
func expandBucketLifecycleSlice(c *Client, f []BucketLifecycle) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLifecycle(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLifecycleMap flattens the contents of BucketLifecycle from a JSON
// response object.
func flattenBucketLifecycleMap(c *Client, i interface{}) map[string]BucketLifecycle {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycle{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycle{}
	}

	items := make(map[string]BucketLifecycle)
	for k, item := range a {
		items[k] = *flattenBucketLifecycle(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBucketLifecycleSlice flattens the contents of BucketLifecycle from a JSON
// response object.
func flattenBucketLifecycleSlice(c *Client, i interface{}) []BucketLifecycle {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycle{}
	}

	if len(a) == 0 {
		return []BucketLifecycle{}
	}

	items := make([]BucketLifecycle, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycle(c, item.(map[string]interface{})))
	}

	return items
}

// expandBucketLifecycle expands an instance of BucketLifecycle into a JSON
// request object.
func expandBucketLifecycle(c *Client, f *BucketLifecycle) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandBucketLifecycleRuleSlice(c, f.Rule); err != nil {
		return nil, fmt.Errorf("error expanding Rule into rule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rule"] = v
	}

	return m, nil
}

// flattenBucketLifecycle flattens an instance of BucketLifecycle from a JSON
// response object.
func flattenBucketLifecycle(c *Client, i interface{}) *BucketLifecycle {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLifecycle{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLifecycle
	}
	r.Rule = flattenBucketLifecycleRuleSlice(c, m["rule"])

	return r
}

// expandBucketLifecycleRuleMap expands the contents of BucketLifecycleRule into a JSON
// request object.
func expandBucketLifecycleRuleMap(c *Client, f map[string]BucketLifecycleRule) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLifecycleRule(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLifecycleRuleSlice expands the contents of BucketLifecycleRule into a JSON
// request object.
func expandBucketLifecycleRuleSlice(c *Client, f []BucketLifecycleRule) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLifecycleRule(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLifecycleRuleMap flattens the contents of BucketLifecycleRule from a JSON
// response object.
func flattenBucketLifecycleRuleMap(c *Client, i interface{}) map[string]BucketLifecycleRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycleRule{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycleRule{}
	}

	items := make(map[string]BucketLifecycleRule)
	for k, item := range a {
		items[k] = *flattenBucketLifecycleRule(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBucketLifecycleRuleSlice flattens the contents of BucketLifecycleRule from a JSON
// response object.
func flattenBucketLifecycleRuleSlice(c *Client, i interface{}) []BucketLifecycleRule {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRule{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRule{}
	}

	items := make([]BucketLifecycleRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRule(c, item.(map[string]interface{})))
	}

	return items
}

// expandBucketLifecycleRule expands an instance of BucketLifecycleRule into a JSON
// request object.
func expandBucketLifecycleRule(c *Client, f *BucketLifecycleRule) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandBucketLifecycleRuleAction(c, f.Action); err != nil {
		return nil, fmt.Errorf("error expanding Action into action: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}
	if v, err := expandBucketLifecycleRuleCondition(c, f.Condition); err != nil {
		return nil, fmt.Errorf("error expanding Condition into condition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["condition"] = v
	}

	return m, nil
}

// flattenBucketLifecycleRule flattens an instance of BucketLifecycleRule from a JSON
// response object.
func flattenBucketLifecycleRule(c *Client, i interface{}) *BucketLifecycleRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLifecycleRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLifecycleRule
	}
	r.Action = flattenBucketLifecycleRuleAction(c, m["action"])
	r.Condition = flattenBucketLifecycleRuleCondition(c, m["condition"])

	return r
}

// expandBucketLifecycleRuleActionMap expands the contents of BucketLifecycleRuleAction into a JSON
// request object.
func expandBucketLifecycleRuleActionMap(c *Client, f map[string]BucketLifecycleRuleAction) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLifecycleRuleAction(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLifecycleRuleActionSlice expands the contents of BucketLifecycleRuleAction into a JSON
// request object.
func expandBucketLifecycleRuleActionSlice(c *Client, f []BucketLifecycleRuleAction) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLifecycleRuleAction(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLifecycleRuleActionMap flattens the contents of BucketLifecycleRuleAction from a JSON
// response object.
func flattenBucketLifecycleRuleActionMap(c *Client, i interface{}) map[string]BucketLifecycleRuleAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycleRuleAction{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycleRuleAction{}
	}

	items := make(map[string]BucketLifecycleRuleAction)
	for k, item := range a {
		items[k] = *flattenBucketLifecycleRuleAction(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBucketLifecycleRuleActionSlice flattens the contents of BucketLifecycleRuleAction from a JSON
// response object.
func flattenBucketLifecycleRuleActionSlice(c *Client, i interface{}) []BucketLifecycleRuleAction {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRuleAction{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRuleAction{}
	}

	items := make([]BucketLifecycleRuleAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRuleAction(c, item.(map[string]interface{})))
	}

	return items
}

// expandBucketLifecycleRuleAction expands an instance of BucketLifecycleRuleAction into a JSON
// request object.
func expandBucketLifecycleRuleAction(c *Client, f *BucketLifecycleRuleAction) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.StorageClass; !dcl.IsEmptyValueIndirect(v) {
		m["storageClass"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenBucketLifecycleRuleAction flattens an instance of BucketLifecycleRuleAction from a JSON
// response object.
func flattenBucketLifecycleRuleAction(c *Client, i interface{}) *BucketLifecycleRuleAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLifecycleRuleAction{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLifecycleRuleAction
	}
	r.StorageClass = dcl.FlattenString(m["storageClass"])
	r.Type = flattenBucketLifecycleRuleActionTypeEnum(m["type"])

	return r
}

// expandBucketLifecycleRuleConditionMap expands the contents of BucketLifecycleRuleCondition into a JSON
// request object.
func expandBucketLifecycleRuleConditionMap(c *Client, f map[string]BucketLifecycleRuleCondition) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLifecycleRuleCondition(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLifecycleRuleConditionSlice expands the contents of BucketLifecycleRuleCondition into a JSON
// request object.
func expandBucketLifecycleRuleConditionSlice(c *Client, f []BucketLifecycleRuleCondition) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLifecycleRuleCondition(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLifecycleRuleConditionMap flattens the contents of BucketLifecycleRuleCondition from a JSON
// response object.
func flattenBucketLifecycleRuleConditionMap(c *Client, i interface{}) map[string]BucketLifecycleRuleCondition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycleRuleCondition{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycleRuleCondition{}
	}

	items := make(map[string]BucketLifecycleRuleCondition)
	for k, item := range a {
		items[k] = *flattenBucketLifecycleRuleCondition(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBucketLifecycleRuleConditionSlice flattens the contents of BucketLifecycleRuleCondition from a JSON
// response object.
func flattenBucketLifecycleRuleConditionSlice(c *Client, i interface{}) []BucketLifecycleRuleCondition {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRuleCondition{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRuleCondition{}
	}

	items := make([]BucketLifecycleRuleCondition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRuleCondition(c, item.(map[string]interface{})))
	}

	return items
}

// expandBucketLifecycleRuleCondition expands an instance of BucketLifecycleRuleCondition into a JSON
// request object.
func expandBucketLifecycleRuleCondition(c *Client, f *BucketLifecycleRuleCondition) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Age; !dcl.IsEmptyValueIndirect(v) {
		m["age"] = v
	}
	if v := f.CreatedBefore; !dcl.IsEmptyValueIndirect(v) {
		m["createdBefore"] = v
	}
	if v, err := expandStorageBucketLifecycleWithState(f, f.WithState); err != nil {
		return nil, fmt.Errorf("error expanding WithState into isLive: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["isLive"] = v
	}
	if v := f.MatchesStorageClass; !dcl.IsEmptyValueIndirect(v) {
		m["matchesStorageClass"] = v
	}
	if v := f.NumNewerVersions; !dcl.IsEmptyValueIndirect(v) {
		m["numNewerVersions"] = v
	}

	return m, nil
}

// flattenBucketLifecycleRuleCondition flattens an instance of BucketLifecycleRuleCondition from a JSON
// response object.
func flattenBucketLifecycleRuleCondition(c *Client, i interface{}) *BucketLifecycleRuleCondition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLifecycleRuleCondition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLifecycleRuleCondition
	}
	r.Age = dcl.FlattenInteger(m["age"])
	r.CreatedBefore = dcl.FlattenString(m["createdBefore"])
	r.WithState = flattenStorageBucketLifecycleWithState(m["isLive"])
	r.MatchesStorageClass = dcl.FlattenStringSlice(m["matchesStorageClass"])
	r.NumNewerVersions = dcl.FlattenInteger(m["numNewerVersions"])

	return r
}

// expandBucketLoggingMap expands the contents of BucketLogging into a JSON
// request object.
func expandBucketLoggingMap(c *Client, f map[string]BucketLogging) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLogging(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLoggingSlice expands the contents of BucketLogging into a JSON
// request object.
func expandBucketLoggingSlice(c *Client, f []BucketLogging) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLogging(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLoggingMap flattens the contents of BucketLogging from a JSON
// response object.
func flattenBucketLoggingMap(c *Client, i interface{}) map[string]BucketLogging {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLogging{}
	}

	if len(a) == 0 {
		return map[string]BucketLogging{}
	}

	items := make(map[string]BucketLogging)
	for k, item := range a {
		items[k] = *flattenBucketLogging(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBucketLoggingSlice flattens the contents of BucketLogging from a JSON
// response object.
func flattenBucketLoggingSlice(c *Client, i interface{}) []BucketLogging {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLogging{}
	}

	if len(a) == 0 {
		return []BucketLogging{}
	}

	items := make([]BucketLogging, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLogging(c, item.(map[string]interface{})))
	}

	return items
}

// expandBucketLogging expands an instance of BucketLogging into a JSON
// request object.
func expandBucketLogging(c *Client, f *BucketLogging) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.LogBucket; !dcl.IsEmptyValueIndirect(v) {
		m["logBucket"] = v
	}
	if v := f.LogObjectPrefix; !dcl.IsEmptyValueIndirect(v) {
		m["logObjectPrefix"] = v
	}

	return m, nil
}

// flattenBucketLogging flattens an instance of BucketLogging from a JSON
// response object.
func flattenBucketLogging(c *Client, i interface{}) *BucketLogging {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLogging{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLogging
	}
	r.LogBucket = dcl.FlattenString(m["logBucket"])
	r.LogObjectPrefix = dcl.FlattenString(m["logObjectPrefix"])

	return r
}

// expandBucketVersioningMap expands the contents of BucketVersioning into a JSON
// request object.
func expandBucketVersioningMap(c *Client, f map[string]BucketVersioning) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketVersioning(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketVersioningSlice expands the contents of BucketVersioning into a JSON
// request object.
func expandBucketVersioningSlice(c *Client, f []BucketVersioning) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketVersioning(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketVersioningMap flattens the contents of BucketVersioning from a JSON
// response object.
func flattenBucketVersioningMap(c *Client, i interface{}) map[string]BucketVersioning {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketVersioning{}
	}

	if len(a) == 0 {
		return map[string]BucketVersioning{}
	}

	items := make(map[string]BucketVersioning)
	for k, item := range a {
		items[k] = *flattenBucketVersioning(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBucketVersioningSlice flattens the contents of BucketVersioning from a JSON
// response object.
func flattenBucketVersioningSlice(c *Client, i interface{}) []BucketVersioning {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketVersioning{}
	}

	if len(a) == 0 {
		return []BucketVersioning{}
	}

	items := make([]BucketVersioning, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketVersioning(c, item.(map[string]interface{})))
	}

	return items
}

// expandBucketVersioning expands an instance of BucketVersioning into a JSON
// request object.
func expandBucketVersioning(c *Client, f *BucketVersioning) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenBucketVersioning flattens an instance of BucketVersioning from a JSON
// response object.
func flattenBucketVersioning(c *Client, i interface{}) *BucketVersioning {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketVersioning{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketVersioning
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandBucketWebsiteMap expands the contents of BucketWebsite into a JSON
// request object.
func expandBucketWebsiteMap(c *Client, f map[string]BucketWebsite) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketWebsite(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketWebsiteSlice expands the contents of BucketWebsite into a JSON
// request object.
func expandBucketWebsiteSlice(c *Client, f []BucketWebsite) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketWebsite(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketWebsiteMap flattens the contents of BucketWebsite from a JSON
// response object.
func flattenBucketWebsiteMap(c *Client, i interface{}) map[string]BucketWebsite {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketWebsite{}
	}

	if len(a) == 0 {
		return map[string]BucketWebsite{}
	}

	items := make(map[string]BucketWebsite)
	for k, item := range a {
		items[k] = *flattenBucketWebsite(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBucketWebsiteSlice flattens the contents of BucketWebsite from a JSON
// response object.
func flattenBucketWebsiteSlice(c *Client, i interface{}) []BucketWebsite {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketWebsite{}
	}

	if len(a) == 0 {
		return []BucketWebsite{}
	}

	items := make([]BucketWebsite, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketWebsite(c, item.(map[string]interface{})))
	}

	return items
}

// expandBucketWebsite expands an instance of BucketWebsite into a JSON
// request object.
func expandBucketWebsite(c *Client, f *BucketWebsite) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MainPageSuffix; !dcl.IsEmptyValueIndirect(v) {
		m["mainPageSuffix"] = v
	}
	if v := f.NotFoundPage; !dcl.IsEmptyValueIndirect(v) {
		m["notFoundPage"] = v
	}

	return m, nil
}

// flattenBucketWebsite flattens an instance of BucketWebsite from a JSON
// response object.
func flattenBucketWebsite(c *Client, i interface{}) *BucketWebsite {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketWebsite{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketWebsite
	}
	r.MainPageSuffix = dcl.FlattenString(m["mainPageSuffix"])
	r.NotFoundPage = dcl.FlattenString(m["notFoundPage"])

	return r
}

// flattenBucketLifecycleRuleActionTypeEnumSlice flattens the contents of BucketLifecycleRuleActionTypeEnum from a JSON
// response object.
func flattenBucketLifecycleRuleActionTypeEnumSlice(c *Client, i interface{}) []BucketLifecycleRuleActionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRuleActionTypeEnum{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRuleActionTypeEnum{}
	}

	items := make([]BucketLifecycleRuleActionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRuleActionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenBucketLifecycleRuleActionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *BucketLifecycleRuleActionTypeEnum with the same value as that string.
func flattenBucketLifecycleRuleActionTypeEnum(i interface{}) *BucketLifecycleRuleActionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return BucketLifecycleRuleActionTypeEnumRef("")
	}

	return BucketLifecycleRuleActionTypeEnumRef(s)
}

// flattenBucketLifecycleRuleConditionWithStateEnumSlice flattens the contents of BucketLifecycleRuleConditionWithStateEnum from a JSON
// response object.
func flattenBucketLifecycleRuleConditionWithStateEnumSlice(c *Client, i interface{}) []BucketLifecycleRuleConditionWithStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRuleConditionWithStateEnum{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRuleConditionWithStateEnum{}
	}

	items := make([]BucketLifecycleRuleConditionWithStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRuleConditionWithStateEnum(item.(interface{})))
	}

	return items
}

// flattenBucketLifecycleRuleConditionWithStateEnum asserts that an interface is a string, and returns a
// pointer to a *BucketLifecycleRuleConditionWithStateEnum with the same value as that string.
func flattenBucketLifecycleRuleConditionWithStateEnum(i interface{}) *BucketLifecycleRuleConditionWithStateEnum {
	s, ok := i.(string)
	if !ok {
		return BucketLifecycleRuleConditionWithStateEnumRef("")
	}

	return BucketLifecycleRuleConditionWithStateEnumRef(s)
}

// flattenBucketStorageClassEnumSlice flattens the contents of BucketStorageClassEnum from a JSON
// response object.
func flattenBucketStorageClassEnumSlice(c *Client, i interface{}) []BucketStorageClassEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketStorageClassEnum{}
	}

	if len(a) == 0 {
		return []BucketStorageClassEnum{}
	}

	items := make([]BucketStorageClassEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketStorageClassEnum(item.(interface{})))
	}

	return items
}

// flattenBucketStorageClassEnum asserts that an interface is a string, and returns a
// pointer to a *BucketStorageClassEnum with the same value as that string.
func flattenBucketStorageClassEnum(i interface{}) *BucketStorageClassEnum {
	s, ok := i.(string)
	if !ok {
		return BucketStorageClassEnumRef("")
	}

	return BucketStorageClassEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Bucket) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalBucket(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
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

type bucketDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         bucketApiOperation
}

func convertFieldDiffToBucketOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]bucketDiff, error) {
	var diffs []bucketDiff
	for _, op := range ops {
		diff := bucketDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTobucketApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTobucketApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (bucketApiOperation, error) {
	switch op {

	case "updateBucketUpdateOperation":
		return &updateBucketUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
