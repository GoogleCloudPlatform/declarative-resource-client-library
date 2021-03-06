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

func (r *Snapshot) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "sourceDisk"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SnapshotEncryptionKey) {
		if err := r.SnapshotEncryptionKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SourceDiskEncryptionKey) {
		if err := r.SourceDiskEncryptionKey.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *SnapshotSnapshotEncryptionKey) validate() error {
	if err := dcl.Required(r, "rawKey"); err != nil {
		return err
	}
	return nil
}
func (r *SnapshotSourceDiskEncryptionKey) validate() error {
	return nil
}

func snapshotGetURL(userBasePath string, r *Snapshot) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/snapshots/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func snapshotListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/snapshots", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func snapshotCreateURL(userBasePath, project, zone, sourceDisk string) (string, error) {
	params := map[string]interface{}{
		"project":    project,
		"zone":       zone,
		"sourceDisk": sourceDisk,
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/disks/{{sourceDisk}}/createSnapshot", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func snapshotDeleteURL(userBasePath string, r *Snapshot) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/snapshots/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// snapshotApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type snapshotApiOperation interface {
	do(context.Context, *Snapshot, *Client) error
}

// newUpdateSnapshotSetLabelsRequest creates a request for an
// Snapshot resource's setLabels update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateSnapshotSetLabelsRequest(ctx context.Context, f *Snapshot, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	b, err := c.getSnapshotRaw(ctx, f.URLNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawLabelFingerprint, err := dcl.GetMapEntry(
		m,
		[]string{"labelFingerprint"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["labelFingerprint"] = rawLabelFingerprint.(string)
	}
	return req, nil
}

// marshalUpdateSnapshotSetLabelsRequest converts the update into
// the final JSON request body.
func marshalUpdateSnapshotSetLabelsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateSnapshotSetLabelsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateSnapshotSetLabelsOperation) do(ctx context.Context, r *Snapshot, c *Client) error {
	_, err := c.GetSnapshot(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setLabels")
	if err != nil {
		return err
	}

	req, err := newUpdateSnapshotSetLabelsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateSnapshotSetLabelsRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

func (c *Client) listSnapshotRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := snapshotListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != SnapshotMaxPage {
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

type listSnapshotOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listSnapshot(ctx context.Context, project, pageToken string, pageSize int32) ([]*Snapshot, string, error) {
	b, err := c.listSnapshotRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listSnapshotOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Snapshot
	for _, v := range m.Items {
		res, err := unmarshalMapSnapshot(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllSnapshot(ctx context.Context, f func(*Snapshot) bool, resources []*Snapshot) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteSnapshot(ctx, res)
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

type deleteSnapshotOperation struct{}

func (op *deleteSnapshotOperation) do(ctx context.Context, r *Snapshot, c *Client) error {
	r, err := c.GetSnapshot(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Snapshot not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetSnapshot checking for existence. error: %v", err)
		return err
	}

	u, err := snapshotDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetSnapshot(ctx, r.URLNormalized())
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
type createSnapshotOperation struct {
	response map[string]interface{}
}

func (op *createSnapshotOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createSnapshotOperation) do(ctx context.Context, r *Snapshot, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, zone, sourceDisk := r.createFields()
	u, err := snapshotCreateURL(c.Config.BasePath, project, zone, sourceDisk)

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

	if _, err := c.GetSnapshot(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getSnapshotRaw(ctx context.Context, r *Snapshot) ([]byte, error) {

	u, err := snapshotGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) snapshotDiffsForRawDesired(ctx context.Context, rawDesired *Snapshot, opts ...dcl.ApplyOption) (initial, desired *Snapshot, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Snapshot
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Snapshot); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Snapshot, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetSnapshot(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Snapshot resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Snapshot resource: %v", err)
		}
		c.Config.Logger.Info("Found that Snapshot resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeSnapshotDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Snapshot: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Snapshot: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeSnapshotInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Snapshot: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeSnapshotDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Snapshot: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffSnapshot(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeSnapshotInitialState(rawInitial, rawDesired *Snapshot) (*Snapshot, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeSnapshotDesiredState(rawDesired, rawInitial *Snapshot, opts ...dcl.ApplyOption) (*Snapshot, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SnapshotEncryptionKey = canonicalizeSnapshotSnapshotEncryptionKey(rawDesired.SnapshotEncryptionKey, nil, opts...)
		rawDesired.SourceDiskEncryptionKey = canonicalizeSnapshotSourceDiskEncryptionKey(rawDesired.SourceDiskEncryptionKey, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Snapshot{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.NameToSelfLink(rawDesired.SourceDisk, rawInitial.SourceDisk) {
		canonicalDesired.SourceDisk = rawInitial.SourceDisk
	} else {
		canonicalDesired.SourceDisk = rawDesired.SourceDisk
	}
	canonicalDesired.SnapshotEncryptionKey = canonicalizeSnapshotSnapshotEncryptionKey(rawDesired.SnapshotEncryptionKey, rawInitial.SnapshotEncryptionKey, opts...)
	canonicalDesired.SourceDiskEncryptionKey = canonicalizeSnapshotSourceDiskEncryptionKey(rawDesired.SourceDiskEncryptionKey, rawInitial.SourceDiskEncryptionKey, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Zone, rawInitial.Zone) {
		canonicalDesired.Zone = rawInitial.Zone
	} else {
		canonicalDesired.Zone = rawDesired.Zone
	}

	return canonicalDesired, nil
}

func canonicalizeSnapshotNewState(c *Client, rawNew, rawDesired *Snapshot) (*Snapshot, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceDisk) && dcl.IsEmptyValueIndirect(rawDesired.SourceDisk) {
		rawNew.SourceDisk = rawDesired.SourceDisk
	} else {
		if dcl.NameToSelfLink(rawDesired.SourceDisk, rawNew.SourceDisk) {
			rawNew.SourceDisk = rawDesired.SourceDisk
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskSizeGb) && dcl.IsEmptyValueIndirect(rawDesired.DiskSizeGb) {
		rawNew.DiskSizeGb = rawDesired.DiskSizeGb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StorageBytes) && dcl.IsEmptyValueIndirect(rawDesired.StorageBytes) {
		rawNew.StorageBytes = rawDesired.StorageBytes
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.License) && dcl.IsEmptyValueIndirect(rawDesired.License) {
		rawNew.License = rawDesired.License
	} else {
	}

	rawNew.SnapshotEncryptionKey = rawDesired.SnapshotEncryptionKey

	rawNew.SourceDiskEncryptionKey = rawDesired.SourceDiskEncryptionKey

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Zone = rawDesired.Zone

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	return rawNew, nil
}

func canonicalizeSnapshotSnapshotEncryptionKey(des, initial *SnapshotSnapshotEncryptionKey, opts ...dcl.ApplyOption) *SnapshotSnapshotEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &SnapshotSnapshotEncryptionKey{}

	if dcl.StringCanonicalize(des.RawKey, initial.RawKey) || dcl.IsZeroValue(des.RawKey) {
		cDes.RawKey = initial.RawKey
	} else {
		cDes.RawKey = des.RawKey
	}

	return cDes
}

func canonicalizeNewSnapshotSnapshotEncryptionKey(c *Client, des, nw *SnapshotSnapshotEncryptionKey) *SnapshotSnapshotEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	nw.RawKey = des.RawKey
	if dcl.StringCanonicalize(des.Sha256, nw.Sha256) {
		nw.Sha256 = des.Sha256
	}

	return nw
}

func canonicalizeNewSnapshotSnapshotEncryptionKeySet(c *Client, des, nw []SnapshotSnapshotEncryptionKey) []SnapshotSnapshotEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []SnapshotSnapshotEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareSnapshotSnapshotEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewSnapshotSnapshotEncryptionKeySlice(c *Client, des, nw []SnapshotSnapshotEncryptionKey) []SnapshotSnapshotEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []SnapshotSnapshotEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSnapshotSnapshotEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeSnapshotSourceDiskEncryptionKey(des, initial *SnapshotSourceDiskEncryptionKey, opts ...dcl.ApplyOption) *SnapshotSourceDiskEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &SnapshotSourceDiskEncryptionKey{}

	if dcl.StringCanonicalize(des.RawKey, initial.RawKey) || dcl.IsZeroValue(des.RawKey) {
		cDes.RawKey = initial.RawKey
	} else {
		cDes.RawKey = des.RawKey
	}

	return cDes
}

func canonicalizeNewSnapshotSourceDiskEncryptionKey(c *Client, des, nw *SnapshotSourceDiskEncryptionKey) *SnapshotSourceDiskEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RawKey, nw.RawKey) {
		nw.RawKey = des.RawKey
	}

	return nw
}

func canonicalizeNewSnapshotSourceDiskEncryptionKeySet(c *Client, des, nw []SnapshotSourceDiskEncryptionKey) []SnapshotSourceDiskEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []SnapshotSourceDiskEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareSnapshotSourceDiskEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewSnapshotSourceDiskEncryptionKeySlice(c *Client, des, nw []SnapshotSourceDiskEncryptionKey) []SnapshotSourceDiskEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []SnapshotSourceDiskEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSnapshotSourceDiskEncryptionKey(c, &d, &n))
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
func diffSnapshot(c *Client, desired, actual *Snapshot, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceDisk, actual.SourceDisk, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceDisk")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskSizeGb, actual.DiskSizeGb, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StorageBytes, actual.StorageBytes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StorageBytes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.License, actual.License, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Licenses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SnapshotEncryptionKey, actual.SnapshotEncryptionKey, dcl.Info{Ignore: true, ObjectFunction: compareSnapshotSnapshotEncryptionKeyNewStyle, EmptyObject: EmptySnapshotSnapshotEncryptionKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SnapshotEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceDiskEncryptionKey, actual.SourceDiskEncryptionKey, dcl.Info{Ignore: true, ObjectFunction: compareSnapshotSourceDiskEncryptionKeyNewStyle, EmptyObject: EmptySnapshotSourceDiskEncryptionKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceDiskEncryptionKey")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateSnapshotSetLabelsOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.Info{Ignore: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zone")); len(ds) != 0 || err != nil {
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

	return newDiffs, nil
}
func compareSnapshotSnapshotEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SnapshotSnapshotEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(SnapshotSnapshotEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SnapshotSnapshotEncryptionKey or *SnapshotSnapshotEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SnapshotSnapshotEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(SnapshotSnapshotEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SnapshotSnapshotEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{Ignore: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256, actual.Sha256, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareSnapshotSourceDiskEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SnapshotSourceDiskEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(SnapshotSourceDiskEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SnapshotSourceDiskEncryptionKey or *SnapshotSourceDiskEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SnapshotSourceDiskEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(SnapshotSourceDiskEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SnapshotSourceDiskEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Snapshot) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Snapshot) createFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Zone), dcl.ValueOrEmptyString(n.SourceDisk)
}

func (r *Snapshot) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Snapshot) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "setLabels" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/snapshots/{{name}}/setLabels", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Snapshot resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Snapshot) marshal(c *Client) ([]byte, error) {
	m, err := expandSnapshot(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Snapshot: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalSnapshot decodes JSON responses into the Snapshot resource schema.
func unmarshalSnapshot(b []byte, c *Client) (*Snapshot, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapSnapshot(m, c)
}

func unmarshalMapSnapshot(m map[string]interface{}, c *Client) (*Snapshot, error) {

	return flattenSnapshot(c, m), nil
}

// expandSnapshot expands Snapshot into a JSON request object.
func expandSnapshot(c *Client, f *Snapshot) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.SourceDisk; !dcl.IsEmptyValueIndirect(v) {
		m["sourceDisk"] = v
	}
	if v := f.DiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskSizeGb"] = v
	}
	if v := f.StorageBytes; !dcl.IsEmptyValueIndirect(v) {
		m["storageBytes"] = v
	}
	if v := f.License; !dcl.IsEmptyValueIndirect(v) {
		m["licenses"] = v
	}
	if v, err := expandSnapshotSnapshotEncryptionKey(c, f.SnapshotEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding SnapshotEncryptionKey into snapshotEncryptionKey: %w", err)
	} else if v != nil {
		m["snapshotEncryptionKey"] = v
	}
	if v, err := expandSnapshotSourceDiskEncryptionKey(c, f.SourceDiskEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding SourceDiskEncryptionKey into sourceDiskEncryptionKey: %w", err)
	} else if v != nil {
		m["sourceDiskEncryptionKey"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Zone into zone: %w", err)
	} else if v != nil {
		m["zone"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}

	return m, nil
}

// flattenSnapshot flattens Snapshot from a JSON request object into the
// Snapshot type.
func flattenSnapshot(c *Client, i interface{}) *Snapshot {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Snapshot{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.SourceDisk = dcl.FlattenString(m["sourceDisk"])
	res.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	res.StorageBytes = dcl.FlattenInteger(m["storageBytes"])
	res.License = dcl.FlattenStringSlice(m["licenses"])
	res.SnapshotEncryptionKey = flattenSnapshotSnapshotEncryptionKey(c, m["snapshotEncryptionKey"])
	res.SourceDiskEncryptionKey = flattenSnapshotSourceDiskEncryptionKey(c, m["sourceDiskEncryptionKey"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Project = dcl.FlattenString(m["project"])
	res.Zone = dcl.FlattenSecretValue(m["zone"])
	res.Id = dcl.FlattenInteger(m["id"])

	return res
}

// expandSnapshotSnapshotEncryptionKeyMap expands the contents of SnapshotSnapshotEncryptionKey into a JSON
// request object.
func expandSnapshotSnapshotEncryptionKeyMap(c *Client, f map[string]SnapshotSnapshotEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSnapshotSnapshotEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSnapshotSnapshotEncryptionKeySlice expands the contents of SnapshotSnapshotEncryptionKey into a JSON
// request object.
func expandSnapshotSnapshotEncryptionKeySlice(c *Client, f []SnapshotSnapshotEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSnapshotSnapshotEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSnapshotSnapshotEncryptionKeyMap flattens the contents of SnapshotSnapshotEncryptionKey from a JSON
// response object.
func flattenSnapshotSnapshotEncryptionKeyMap(c *Client, i interface{}) map[string]SnapshotSnapshotEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SnapshotSnapshotEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]SnapshotSnapshotEncryptionKey{}
	}

	items := make(map[string]SnapshotSnapshotEncryptionKey)
	for k, item := range a {
		items[k] = *flattenSnapshotSnapshotEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSnapshotSnapshotEncryptionKeySlice flattens the contents of SnapshotSnapshotEncryptionKey from a JSON
// response object.
func flattenSnapshotSnapshotEncryptionKeySlice(c *Client, i interface{}) []SnapshotSnapshotEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []SnapshotSnapshotEncryptionKey{}
	}

	if len(a) == 0 {
		return []SnapshotSnapshotEncryptionKey{}
	}

	items := make([]SnapshotSnapshotEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSnapshotSnapshotEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandSnapshotSnapshotEncryptionKey expands an instance of SnapshotSnapshotEncryptionKey into a JSON
// request object.
func expandSnapshotSnapshotEncryptionKey(c *Client, f *SnapshotSnapshotEncryptionKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RawKey; !dcl.IsEmptyValueIndirect(v) {
		m["rawKey"] = v
	}
	if v := f.Sha256; !dcl.IsEmptyValueIndirect(v) {
		m["sha256"] = v
	}

	return m, nil
}

// flattenSnapshotSnapshotEncryptionKey flattens an instance of SnapshotSnapshotEncryptionKey from a JSON
// response object.
func flattenSnapshotSnapshotEncryptionKey(c *Client, i interface{}) *SnapshotSnapshotEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SnapshotSnapshotEncryptionKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySnapshotSnapshotEncryptionKey
	}
	r.RawKey = dcl.FlattenSecretValue(m["rawKey"])
	r.Sha256 = dcl.FlattenString(m["sha256"])

	return r
}

// expandSnapshotSourceDiskEncryptionKeyMap expands the contents of SnapshotSourceDiskEncryptionKey into a JSON
// request object.
func expandSnapshotSourceDiskEncryptionKeyMap(c *Client, f map[string]SnapshotSourceDiskEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSnapshotSourceDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSnapshotSourceDiskEncryptionKeySlice expands the contents of SnapshotSourceDiskEncryptionKey into a JSON
// request object.
func expandSnapshotSourceDiskEncryptionKeySlice(c *Client, f []SnapshotSourceDiskEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSnapshotSourceDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSnapshotSourceDiskEncryptionKeyMap flattens the contents of SnapshotSourceDiskEncryptionKey from a JSON
// response object.
func flattenSnapshotSourceDiskEncryptionKeyMap(c *Client, i interface{}) map[string]SnapshotSourceDiskEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SnapshotSourceDiskEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]SnapshotSourceDiskEncryptionKey{}
	}

	items := make(map[string]SnapshotSourceDiskEncryptionKey)
	for k, item := range a {
		items[k] = *flattenSnapshotSourceDiskEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSnapshotSourceDiskEncryptionKeySlice flattens the contents of SnapshotSourceDiskEncryptionKey from a JSON
// response object.
func flattenSnapshotSourceDiskEncryptionKeySlice(c *Client, i interface{}) []SnapshotSourceDiskEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []SnapshotSourceDiskEncryptionKey{}
	}

	if len(a) == 0 {
		return []SnapshotSourceDiskEncryptionKey{}
	}

	items := make([]SnapshotSourceDiskEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSnapshotSourceDiskEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandSnapshotSourceDiskEncryptionKey expands an instance of SnapshotSourceDiskEncryptionKey into a JSON
// request object.
func expandSnapshotSourceDiskEncryptionKey(c *Client, f *SnapshotSourceDiskEncryptionKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RawKey; !dcl.IsEmptyValueIndirect(v) {
		m["rawKey"] = v
	}

	return m, nil
}

// flattenSnapshotSourceDiskEncryptionKey flattens an instance of SnapshotSourceDiskEncryptionKey from a JSON
// response object.
func flattenSnapshotSourceDiskEncryptionKey(c *Client, i interface{}) *SnapshotSourceDiskEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SnapshotSourceDiskEncryptionKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySnapshotSourceDiskEncryptionKey
	}
	r.RawKey = dcl.FlattenString(m["rawKey"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Snapshot) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalSnapshot(b, c)
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

type snapshotDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         snapshotApiOperation
}

func convertFieldDiffsToSnapshotDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]snapshotDiff, error) {
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
	var diffs []snapshotDiff
	// For each operation name, create a snapshotDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := snapshotDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToSnapshotApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToSnapshotApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (snapshotApiOperation, error) {
	switch opName {

	case "updateSnapshotSetLabelsOperation":
		return &updateSnapshotSetLabelsOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
