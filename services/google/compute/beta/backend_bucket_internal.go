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

func (r *BackendBucket) validate() error {

	if !dcl.IsEmptyValueIndirect(r.CdnPolicy) {
		if err := r.CdnPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BackendBucketCdnPolicy) validate() error {
	return nil
}

func backendBucketGetURL(userBasePath string, r *BackendBucket) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/backendBuckets/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func backendBucketListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/backendBuckets", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func backendBucketCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/backendBuckets", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func backendBucketDeleteURL(userBasePath string, r *BackendBucket) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/backendBuckets/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// backendBucketApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type backendBucketApiOperation interface {
	do(context.Context, *BackendBucket, *Client) error
}

// newUpdateBackendBucketUpdateRequest creates a request for an
// BackendBucket resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateBackendBucketUpdateRequest(ctx context.Context, f *BackendBucket, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.BucketName; !dcl.IsEmptyValueIndirect(v) {
		req["bucketName"] = v
	}
	if v, err := expandBackendBucketCdnPolicy(c, f.CdnPolicy); err != nil {
		return nil, fmt.Errorf("error expanding CdnPolicy into cdnPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["cdnPolicy"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.EnableCdn; !dcl.IsEmptyValueIndirect(v) {
		req["enableCdn"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	return req, nil
}

// marshalUpdateBackendBucketUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateBackendBucketUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateBackendBucketUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateBackendBucketUpdateOperation) do(ctx context.Context, r *BackendBucket, c *Client) error {
	_, err := c.GetBackendBucket(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateBackendBucketUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateBackendBucketUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listBackendBucketRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := backendBucketListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != BackendBucketMaxPage {
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

type listBackendBucketOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listBackendBucket(ctx context.Context, project, pageToken string, pageSize int32) ([]*BackendBucket, string, error) {
	b, err := c.listBackendBucketRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listBackendBucketOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*BackendBucket
	for _, v := range m.Items {
		res, err := unmarshalMapBackendBucket(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllBackendBucket(ctx context.Context, f func(*BackendBucket) bool, resources []*BackendBucket) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteBackendBucket(ctx, res)
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

type deleteBackendBucketOperation struct{}

func (op *deleteBackendBucketOperation) do(ctx context.Context, r *BackendBucket, c *Client) error {
	r, err := c.GetBackendBucket(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("BackendBucket not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetBackendBucket checking for existence. error: %v", err)
		return err
	}

	u, err := backendBucketDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetBackendBucket(ctx, r.URLNormalized())
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
type createBackendBucketOperation struct {
	response map[string]interface{}
}

func (op *createBackendBucketOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createBackendBucketOperation) do(ctx context.Context, r *BackendBucket, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := backendBucketCreateURL(c.Config.BasePath, project)

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
	// wait for object to be created.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetBackendBucket(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getBackendBucketRaw(ctx context.Context, r *BackendBucket) ([]byte, error) {

	u, err := backendBucketGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) backendBucketDiffsForRawDesired(ctx context.Context, rawDesired *BackendBucket, opts ...dcl.ApplyOption) (initial, desired *BackendBucket, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *BackendBucket
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*BackendBucket); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected BackendBucket, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetBackendBucket(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a BackendBucket resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve BackendBucket resource: %v", err)
		}
		c.Config.Logger.Info("Found that BackendBucket resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeBackendBucketDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for BackendBucket: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for BackendBucket: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeBackendBucketInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for BackendBucket: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeBackendBucketDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for BackendBucket: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffBackendBucket(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeBackendBucketInitialState(rawInitial, rawDesired *BackendBucket) (*BackendBucket, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeBackendBucketDesiredState(rawDesired, rawInitial *BackendBucket, opts ...dcl.ApplyOption) (*BackendBucket, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.CdnPolicy = canonicalizeBackendBucketCdnPolicy(rawDesired.CdnPolicy, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &BackendBucket{}
	if dcl.NameToSelfLink(rawDesired.BucketName, rawInitial.BucketName) {
		canonicalDesired.BucketName = rawInitial.BucketName
	} else {
		canonicalDesired.BucketName = rawDesired.BucketName
	}
	canonicalDesired.CdnPolicy = canonicalizeBackendBucketCdnPolicy(rawDesired.CdnPolicy, rawInitial.CdnPolicy, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.BoolCanonicalize(rawDesired.EnableCdn, rawInitial.EnableCdn) {
		canonicalDesired.EnableCdn = rawInitial.EnableCdn
	} else {
		canonicalDesired.EnableCdn = rawDesired.EnableCdn
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeBackendBucketNewState(c *Client, rawNew, rawDesired *BackendBucket) (*BackendBucket, error) {

	if dcl.IsEmptyValueIndirect(rawNew.BucketName) && dcl.IsEmptyValueIndirect(rawDesired.BucketName) {
		rawNew.BucketName = rawDesired.BucketName
	} else {
		if dcl.NameToSelfLink(rawDesired.BucketName, rawNew.BucketName) {
			rawNew.BucketName = rawDesired.BucketName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CdnPolicy) && dcl.IsEmptyValueIndirect(rawDesired.CdnPolicy) {
		rawNew.CdnPolicy = rawDesired.CdnPolicy
	} else {
		rawNew.CdnPolicy = canonicalizeNewBackendBucketCdnPolicy(c, rawDesired.CdnPolicy, rawNew.CdnPolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableCdn) && dcl.IsEmptyValueIndirect(rawDesired.EnableCdn) {
		rawNew.EnableCdn = rawDesired.EnableCdn
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableCdn, rawNew.EnableCdn) {
			rawNew.EnableCdn = rawDesired.EnableCdn
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	return rawNew, nil
}

func canonicalizeBackendBucketCdnPolicy(des, initial *BackendBucketCdnPolicy, opts ...dcl.ApplyOption) *BackendBucketCdnPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BackendBucketCdnPolicy{}

	if dcl.IsZeroValue(des.SignedUrlKeyNames) {
		des.SignedUrlKeyNames = initial.SignedUrlKeyNames
	} else {
		cDes.SignedUrlKeyNames = des.SignedUrlKeyNames
	}
	if dcl.IsZeroValue(des.SignedUrlCacheMaxAgeSec) {
		des.SignedUrlCacheMaxAgeSec = initial.SignedUrlCacheMaxAgeSec
	} else {
		cDes.SignedUrlCacheMaxAgeSec = des.SignedUrlCacheMaxAgeSec
	}

	return cDes
}

func canonicalizeNewBackendBucketCdnPolicy(c *Client, des, nw *BackendBucketCdnPolicy) *BackendBucketCdnPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.SignedUrlKeyNames) {
		nw.SignedUrlKeyNames = des.SignedUrlKeyNames
	}
	if dcl.IsZeroValue(nw.SignedUrlCacheMaxAgeSec) {
		nw.SignedUrlCacheMaxAgeSec = des.SignedUrlCacheMaxAgeSec
	}

	return nw
}

func canonicalizeNewBackendBucketCdnPolicySet(c *Client, des, nw []BackendBucketCdnPolicy) []BackendBucketCdnPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendBucketCdnPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendBucketCdnPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendBucketCdnPolicySlice(c *Client, des, nw []BackendBucketCdnPolicy) []BackendBucketCdnPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendBucketCdnPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendBucketCdnPolicy(c, &d, &n))
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
func diffBackendBucket(c *Client, desired, actual *BackendBucket, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.BucketName, actual.BucketName, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateBackendBucketUpdateOperation")}, fn.AddNest("BucketName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CdnPolicy, actual.CdnPolicy, dcl.Info{ObjectFunction: compareBackendBucketCdnPolicyNewStyle, EmptyObject: EmptyBackendBucketCdnPolicy, OperationSelector: dcl.TriggersOperation("updateBackendBucketUpdateOperation")}, fn.AddNest("CdnPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendBucketUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableCdn, actual.EnableCdn, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendBucketUpdateOperation")}, fn.AddNest("EnableCdn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendBucketUpdateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
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

	return newDiffs, nil
}
func compareBackendBucketCdnPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendBucketCdnPolicy)
	if !ok {
		desiredNotPointer, ok := d.(BackendBucketCdnPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendBucketCdnPolicy or *BackendBucketCdnPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendBucketCdnPolicy)
	if !ok {
		actualNotPointer, ok := a.(BackendBucketCdnPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendBucketCdnPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SignedUrlKeyNames, actual.SignedUrlKeyNames, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SignedUrlKeyNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SignedUrlCacheMaxAgeSec, actual.SignedUrlCacheMaxAgeSec, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SignedUrlCacheMaxAgeSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *BackendBucket) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *BackendBucket) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *BackendBucket) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *BackendBucket) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/backendBuckets/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the BackendBucket resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *BackendBucket) marshal(c *Client) ([]byte, error) {
	m, err := expandBackendBucket(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling BackendBucket: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalBackendBucket decodes JSON responses into the BackendBucket resource schema.
func unmarshalBackendBucket(b []byte, c *Client) (*BackendBucket, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapBackendBucket(m, c)
}

func unmarshalMapBackendBucket(m map[string]interface{}, c *Client) (*BackendBucket, error) {

	return flattenBackendBucket(c, m), nil
}

// expandBackendBucket expands BackendBucket into a JSON request object.
func expandBackendBucket(c *Client, f *BackendBucket) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.BucketName; !dcl.IsEmptyValueIndirect(v) {
		m["bucketName"] = v
	}
	if v, err := expandBackendBucketCdnPolicy(c, f.CdnPolicy); err != nil {
		return nil, fmt.Errorf("error expanding CdnPolicy into cdnPolicy: %w", err)
	} else if v != nil {
		m["cdnPolicy"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.EnableCdn; !dcl.IsEmptyValueIndirect(v) {
		m["enableCdn"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}

	return m, nil
}

// flattenBackendBucket flattens BackendBucket from a JSON request object into the
// BackendBucket type.
func flattenBackendBucket(c *Client, i interface{}) *BackendBucket {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &BackendBucket{}
	res.BucketName = dcl.FlattenString(m["bucketName"])
	res.CdnPolicy = flattenBackendBucketCdnPolicy(c, m["cdnPolicy"])
	res.Description = dcl.FlattenString(m["description"])
	res.EnableCdn = dcl.FlattenBool(m["enableCdn"])
	res.Name = dcl.FlattenString(m["name"])
	res.Project = dcl.FlattenString(m["project"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])

	return res
}

// expandBackendBucketCdnPolicyMap expands the contents of BackendBucketCdnPolicy into a JSON
// request object.
func expandBackendBucketCdnPolicyMap(c *Client, f map[string]BackendBucketCdnPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendBucketCdnPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendBucketCdnPolicySlice expands the contents of BackendBucketCdnPolicy into a JSON
// request object.
func expandBackendBucketCdnPolicySlice(c *Client, f []BackendBucketCdnPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendBucketCdnPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendBucketCdnPolicyMap flattens the contents of BackendBucketCdnPolicy from a JSON
// response object.
func flattenBackendBucketCdnPolicyMap(c *Client, i interface{}) map[string]BackendBucketCdnPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendBucketCdnPolicy{}
	}

	if len(a) == 0 {
		return map[string]BackendBucketCdnPolicy{}
	}

	items := make(map[string]BackendBucketCdnPolicy)
	for k, item := range a {
		items[k] = *flattenBackendBucketCdnPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendBucketCdnPolicySlice flattens the contents of BackendBucketCdnPolicy from a JSON
// response object.
func flattenBackendBucketCdnPolicySlice(c *Client, i interface{}) []BackendBucketCdnPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendBucketCdnPolicy{}
	}

	if len(a) == 0 {
		return []BackendBucketCdnPolicy{}
	}

	items := make([]BackendBucketCdnPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendBucketCdnPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendBucketCdnPolicy expands an instance of BackendBucketCdnPolicy into a JSON
// request object.
func expandBackendBucketCdnPolicy(c *Client, f *BackendBucketCdnPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SignedUrlKeyNames; !dcl.IsEmptyValueIndirect(v) {
		m["signedUrlKeyNames"] = v
	}
	if v := f.SignedUrlCacheMaxAgeSec; !dcl.IsEmptyValueIndirect(v) {
		m["signedUrlCacheMaxAgeSec"] = v
	}

	return m, nil
}

// flattenBackendBucketCdnPolicy flattens an instance of BackendBucketCdnPolicy from a JSON
// response object.
func flattenBackendBucketCdnPolicy(c *Client, i interface{}) *BackendBucketCdnPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendBucketCdnPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendBucketCdnPolicy
	}
	r.SignedUrlKeyNames = dcl.FlattenStringSlice(m["signedUrlKeyNames"])
	r.SignedUrlCacheMaxAgeSec = dcl.FlattenInteger(m["signedUrlCacheMaxAgeSec"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *BackendBucket) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalBackendBucket(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
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

type backendBucketDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         backendBucketApiOperation
}

func convertFieldDiffsToBackendBucketDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]backendBucketDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []backendBucketDiff
	// For each operation name, create a backendBucketDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := backendBucketDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToBackendBucketApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToBackendBucketApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (backendBucketApiOperation, error) {
	switch opName {

	case "updateBackendBucketUpdateOperation":
		return &updateBackendBucketUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
