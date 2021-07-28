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

func (r *Image) validate() error {

	if !dcl.IsEmptyValueIndirect(r.ImageEncryptionKey) {
		if err := r.ImageEncryptionKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RawDisk) {
		if err := r.RawDisk.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ShieldedInstanceInitialState) {
		if err := r.ShieldedInstanceInitialState.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SourceDiskEncryptionKey) {
		if err := r.SourceDiskEncryptionKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SourceImageEncryptionKey) {
		if err := r.SourceImageEncryptionKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SourceSnapshotEncryptionKey) {
		if err := r.SourceSnapshotEncryptionKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Deprecated) {
		if err := r.Deprecated.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ImageGuestOSFeature) validate() error {
	return nil
}
func (r *ImageImageEncryptionKey) validate() error {
	return nil
}
func (r *ImageRawDisk) validate() error {
	return nil
}
func (r *ImageShieldedInstanceInitialState) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Pk) {
		if err := r.Pk.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ImageShieldedInstanceInitialStatePk) validate() error {
	return nil
}
func (r *ImageShieldedInstanceInitialStateKek) validate() error {
	return nil
}
func (r *ImageShieldedInstanceInitialStateDb) validate() error {
	return nil
}
func (r *ImageShieldedInstanceInitialStateDbx) validate() error {
	return nil
}
func (r *ImageSourceDiskEncryptionKey) validate() error {
	return nil
}
func (r *ImageSourceImageEncryptionKey) validate() error {
	return nil
}
func (r *ImageSourceSnapshotEncryptionKey) validate() error {
	return nil
}
func (r *ImageDeprecated) validate() error {
	return nil
}

func imageGetURL(userBasePath string, r *Image) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/images/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func imageListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/images", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func imageCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/images", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func imageDeleteURL(userBasePath string, r *Image) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/images/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// imageApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type imageApiOperation interface {
	do(context.Context, *Image, *Client) error
}

// newUpdateImageDeprecateRequest creates a request for an
// Image resource's deprecate update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateImageDeprecateRequest(ctx context.Context, f *Image, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandImageDeprecated(c, f.Deprecated); err != nil {
		return nil, fmt.Errorf("error expanding Deprecated into deprecated: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["deprecated"] = v
	}
	return req, nil
}

// marshalUpdateImageDeprecateRequest converts the update into
// the final JSON request body.
func marshalUpdateImageDeprecateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeImageDeprecateRequest(m)
	return json.Marshal(m)
}

type updateImageDeprecateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateImageDeprecateOperation) do(ctx context.Context, r *Image, c *Client) error {
	_, err := c.GetImage(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "deprecate")
	if err != nil {
		return err
	}

	req, err := newUpdateImageDeprecateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateImageDeprecateRequest(c, req)
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

// newUpdateImageSetLabelsRequest creates a request for an
// Image resource's setLabels update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateImageSetLabelsRequest(ctx context.Context, f *Image, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	b, err := c.getImageRaw(ctx, f.URLNormalized())
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

// marshalUpdateImageSetLabelsRequest converts the update into
// the final JSON request body.
func marshalUpdateImageSetLabelsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateImageSetLabelsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateImageSetLabelsOperation) do(ctx context.Context, r *Image, c *Client) error {
	_, err := c.GetImage(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setLabels")
	if err != nil {
		return err
	}

	req, err := newUpdateImageSetLabelsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateImageSetLabelsRequest(c, req)
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

func (c *Client) listImageRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := imageListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ImageMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "LIST", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listImageOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listImage(ctx context.Context, project, pageToken string, pageSize int32) ([]*Image, string, error) {
	b, err := c.listImageRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listImageOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Image
	for _, v := range m.Items {
		res, err := unmarshalMapImage(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllImage(ctx context.Context, f func(*Image) bool, resources []*Image) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteImage(ctx, res)
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

type deleteImageOperation struct{}

func (op *deleteImageOperation) do(ctx context.Context, r *Image, c *Client) error {
	r, err := c.GetImage(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Image not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetImage checking for existence. error: %v", err)
		return err
	}

	u, err := imageDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetImage(ctx, r.URLNormalized())
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
type createImageOperation struct {
	response map[string]interface{}
}

func (op *createImageOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createImageOperation) do(ctx context.Context, r *Image, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := imageCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetImage(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getImageRaw(ctx context.Context, r *Image) ([]byte, error) {

	u, err := imageGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) imageDiffsForRawDesired(ctx context.Context, rawDesired *Image, opts ...dcl.ApplyOption) (initial, desired *Image, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Image
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Image); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Image, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetImage(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Image resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Image resource: %v", err)
		}
		c.Config.Logger.Info("Found that Image resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeImageDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Image: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Image: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeImageInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Image: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeImageDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Image: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffImage(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeImageInitialState(rawInitial, rawDesired *Image) (*Image, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeImageDesiredState(rawDesired, rawInitial *Image, opts ...dcl.ApplyOption) (*Image, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ImageEncryptionKey = canonicalizeImageImageEncryptionKey(rawDesired.ImageEncryptionKey, nil, opts...)
		rawDesired.RawDisk = canonicalizeImageRawDisk(rawDesired.RawDisk, nil, opts...)
		rawDesired.ShieldedInstanceInitialState = canonicalizeImageShieldedInstanceInitialState(rawDesired.ShieldedInstanceInitialState, nil, opts...)
		rawDesired.SourceDiskEncryptionKey = canonicalizeImageSourceDiskEncryptionKey(rawDesired.SourceDiskEncryptionKey, nil, opts...)
		rawDesired.SourceImageEncryptionKey = canonicalizeImageSourceImageEncryptionKey(rawDesired.SourceImageEncryptionKey, nil, opts...)
		rawDesired.SourceSnapshotEncryptionKey = canonicalizeImageSourceSnapshotEncryptionKey(rawDesired.SourceSnapshotEncryptionKey, nil, opts...)
		rawDesired.Deprecated = canonicalizeImageDeprecated(rawDesired.Deprecated, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Image{}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.DiskSizeGb) {
		canonicalDesired.DiskSizeGb = rawInitial.DiskSizeGb
	} else {
		canonicalDesired.DiskSizeGb = rawDesired.DiskSizeGb
	}
	if dcl.StringCanonicalize(rawDesired.Family, rawInitial.Family) {
		canonicalDesired.Family = rawInitial.Family
	} else {
		canonicalDesired.Family = rawDesired.Family
	}
	if dcl.IsZeroValue(rawDesired.GuestOSFeature) {
		canonicalDesired.GuestOSFeature = rawInitial.GuestOSFeature
	} else {
		canonicalDesired.GuestOSFeature = rawDesired.GuestOSFeature
	}
	canonicalDesired.ImageEncryptionKey = canonicalizeImageImageEncryptionKey(rawDesired.ImageEncryptionKey, rawInitial.ImageEncryptionKey, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.IsZeroValue(rawDesired.License) {
		canonicalDesired.License = rawInitial.License
	} else {
		canonicalDesired.License = rawDesired.License
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.RawDisk = canonicalizeImageRawDisk(rawDesired.RawDisk, rawInitial.RawDisk, opts...)
	canonicalDesired.ShieldedInstanceInitialState = canonicalizeImageShieldedInstanceInitialState(rawDesired.ShieldedInstanceInitialState, rawInitial.ShieldedInstanceInitialState, opts...)
	if dcl.StringCanonicalize(rawDesired.SourceDisk, rawInitial.SourceDisk) {
		canonicalDesired.SourceDisk = rawInitial.SourceDisk
	} else {
		canonicalDesired.SourceDisk = rawDesired.SourceDisk
	}
	canonicalDesired.SourceDiskEncryptionKey = canonicalizeImageSourceDiskEncryptionKey(rawDesired.SourceDiskEncryptionKey, rawInitial.SourceDiskEncryptionKey, opts...)
	if dcl.StringCanonicalize(rawDesired.SourceImage, rawInitial.SourceImage) {
		canonicalDesired.SourceImage = rawInitial.SourceImage
	} else {
		canonicalDesired.SourceImage = rawDesired.SourceImage
	}
	canonicalDesired.SourceImageEncryptionKey = canonicalizeImageSourceImageEncryptionKey(rawDesired.SourceImageEncryptionKey, rawInitial.SourceImageEncryptionKey, opts...)
	if dcl.StringCanonicalize(rawDesired.SourceImageId, rawInitial.SourceImageId) {
		canonicalDesired.SourceImageId = rawInitial.SourceImageId
	} else {
		canonicalDesired.SourceImageId = rawDesired.SourceImageId
	}
	if dcl.StringCanonicalize(rawDesired.SourceSnapshot, rawInitial.SourceSnapshot) {
		canonicalDesired.SourceSnapshot = rawInitial.SourceSnapshot
	} else {
		canonicalDesired.SourceSnapshot = rawDesired.SourceSnapshot
	}
	canonicalDesired.SourceSnapshotEncryptionKey = canonicalizeImageSourceSnapshotEncryptionKey(rawDesired.SourceSnapshotEncryptionKey, rawInitial.SourceSnapshotEncryptionKey, opts...)
	if dcl.StringCanonicalize(rawDesired.SourceSnapshotId, rawInitial.SourceSnapshotId) {
		canonicalDesired.SourceSnapshotId = rawInitial.SourceSnapshotId
	} else {
		canonicalDesired.SourceSnapshotId = rawDesired.SourceSnapshotId
	}
	if dcl.IsZeroValue(rawDesired.SourceType) {
		canonicalDesired.SourceType = rawInitial.SourceType
	} else {
		canonicalDesired.SourceType = rawDesired.SourceType
	}
	if dcl.IsZeroValue(rawDesired.StorageLocation) {
		canonicalDesired.StorageLocation = rawInitial.StorageLocation
	} else {
		canonicalDesired.StorageLocation = rawDesired.StorageLocation
	}
	canonicalDesired.Deprecated = canonicalizeImageDeprecated(rawDesired.Deprecated, rawInitial.Deprecated, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeImageNewState(c *Client, rawNew, rawDesired *Image) (*Image, error) {

	if dcl.IsEmptyValueIndirect(rawNew.ArchiveSizeBytes) && dcl.IsEmptyValueIndirect(rawDesired.ArchiveSizeBytes) {
		rawNew.ArchiveSizeBytes = rawDesired.ArchiveSizeBytes
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskSizeGb) && dcl.IsEmptyValueIndirect(rawDesired.DiskSizeGb) {
		rawNew.DiskSizeGb = rawDesired.DiskSizeGb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Family) && dcl.IsEmptyValueIndirect(rawDesired.Family) {
		rawNew.Family = rawDesired.Family
	} else {
		if dcl.StringCanonicalize(rawDesired.Family, rawNew.Family) {
			rawNew.Family = rawDesired.Family
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GuestOSFeature) && dcl.IsEmptyValueIndirect(rawDesired.GuestOSFeature) {
		rawNew.GuestOSFeature = rawDesired.GuestOSFeature
	} else {
		rawNew.GuestOSFeature = canonicalizeNewImageGuestOSFeatureSet(c, rawDesired.GuestOSFeature, rawNew.GuestOSFeature)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ImageEncryptionKey) && dcl.IsEmptyValueIndirect(rawDesired.ImageEncryptionKey) {
		rawNew.ImageEncryptionKey = rawDesired.ImageEncryptionKey
	} else {
		rawNew.ImageEncryptionKey = canonicalizeNewImageImageEncryptionKey(c, rawDesired.ImageEncryptionKey, rawNew.ImageEncryptionKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.License) && dcl.IsEmptyValueIndirect(rawDesired.License) {
		rawNew.License = rawDesired.License
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	rawNew.RawDisk = rawDesired.RawDisk

	if dcl.IsEmptyValueIndirect(rawNew.ShieldedInstanceInitialState) && dcl.IsEmptyValueIndirect(rawDesired.ShieldedInstanceInitialState) {
		rawNew.ShieldedInstanceInitialState = rawDesired.ShieldedInstanceInitialState
	} else {
		rawNew.ShieldedInstanceInitialState = canonicalizeNewImageShieldedInstanceInitialState(c, rawDesired.ShieldedInstanceInitialState, rawNew.ShieldedInstanceInitialState)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceDisk) && dcl.IsEmptyValueIndirect(rawDesired.SourceDisk) {
		rawNew.SourceDisk = rawDesired.SourceDisk
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceDisk, rawNew.SourceDisk) {
			rawNew.SourceDisk = rawDesired.SourceDisk
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceDiskEncryptionKey) && dcl.IsEmptyValueIndirect(rawDesired.SourceDiskEncryptionKey) {
		rawNew.SourceDiskEncryptionKey = rawDesired.SourceDiskEncryptionKey
	} else {
		rawNew.SourceDiskEncryptionKey = canonicalizeNewImageSourceDiskEncryptionKey(c, rawDesired.SourceDiskEncryptionKey, rawNew.SourceDiskEncryptionKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceDiskId) && dcl.IsEmptyValueIndirect(rawDesired.SourceDiskId) {
		rawNew.SourceDiskId = rawDesired.SourceDiskId
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceDiskId, rawNew.SourceDiskId) {
			rawNew.SourceDiskId = rawDesired.SourceDiskId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceImage) && dcl.IsEmptyValueIndirect(rawDesired.SourceImage) {
		rawNew.SourceImage = rawDesired.SourceImage
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceImage, rawNew.SourceImage) {
			rawNew.SourceImage = rawDesired.SourceImage
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceImageEncryptionKey) && dcl.IsEmptyValueIndirect(rawDesired.SourceImageEncryptionKey) {
		rawNew.SourceImageEncryptionKey = rawDesired.SourceImageEncryptionKey
	} else {
		rawNew.SourceImageEncryptionKey = canonicalizeNewImageSourceImageEncryptionKey(c, rawDesired.SourceImageEncryptionKey, rawNew.SourceImageEncryptionKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceImageId) && dcl.IsEmptyValueIndirect(rawDesired.SourceImageId) {
		rawNew.SourceImageId = rawDesired.SourceImageId
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceImageId, rawNew.SourceImageId) {
			rawNew.SourceImageId = rawDesired.SourceImageId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceSnapshot) && dcl.IsEmptyValueIndirect(rawDesired.SourceSnapshot) {
		rawNew.SourceSnapshot = rawDesired.SourceSnapshot
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceSnapshot, rawNew.SourceSnapshot) {
			rawNew.SourceSnapshot = rawDesired.SourceSnapshot
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceSnapshotEncryptionKey) && dcl.IsEmptyValueIndirect(rawDesired.SourceSnapshotEncryptionKey) {
		rawNew.SourceSnapshotEncryptionKey = rawDesired.SourceSnapshotEncryptionKey
	} else {
		rawNew.SourceSnapshotEncryptionKey = canonicalizeNewImageSourceSnapshotEncryptionKey(c, rawDesired.SourceSnapshotEncryptionKey, rawNew.SourceSnapshotEncryptionKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceSnapshotId) && dcl.IsEmptyValueIndirect(rawDesired.SourceSnapshotId) {
		rawNew.SourceSnapshotId = rawDesired.SourceSnapshotId
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceSnapshotId, rawNew.SourceSnapshotId) {
			rawNew.SourceSnapshotId = rawDesired.SourceSnapshotId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceType) && dcl.IsEmptyValueIndirect(rawDesired.SourceType) {
		rawNew.SourceType = rawDesired.SourceType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StorageLocation) && dcl.IsEmptyValueIndirect(rawDesired.StorageLocation) {
		rawNew.StorageLocation = rawDesired.StorageLocation
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Deprecated) && dcl.IsEmptyValueIndirect(rawDesired.Deprecated) {
		rawNew.Deprecated = rawDesired.Deprecated
	} else {
		rawNew.Deprecated = canonicalizeNewImageDeprecated(c, rawDesired.Deprecated, rawNew.Deprecated)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeImageGuestOSFeature(des, initial *ImageGuestOSFeature, opts ...dcl.ApplyOption) *ImageGuestOSFeature {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageGuestOSFeature{}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}

	return cDes
}

func canonicalizeNewImageGuestOSFeature(c *Client, des, nw *ImageGuestOSFeature) *ImageGuestOSFeature {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
	}

	return nw
}

func canonicalizeNewImageGuestOSFeatureSet(c *Client, des, nw []ImageGuestOSFeature) []ImageGuestOSFeature {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageGuestOSFeature
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageGuestOSFeatureNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageGuestOSFeatureSlice(c *Client, des, nw []ImageGuestOSFeature) []ImageGuestOSFeature {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageGuestOSFeature
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageGuestOSFeature(c, &d, &n))
	}

	return items
}

func canonicalizeImageImageEncryptionKey(des, initial *ImageImageEncryptionKey, opts ...dcl.ApplyOption) *ImageImageEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageImageEncryptionKey{}

	if dcl.StringCanonicalize(des.RawKey, initial.RawKey) || dcl.IsZeroValue(des.RawKey) {
		cDes.RawKey = initial.RawKey
	} else {
		cDes.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Sha256, initial.Sha256) || dcl.IsZeroValue(des.Sha256) {
		cDes.Sha256 = initial.Sha256
	} else {
		cDes.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyServiceAccount, initial.KmsKeyServiceAccount) || dcl.IsZeroValue(des.KmsKeyServiceAccount) {
		cDes.KmsKeyServiceAccount = initial.KmsKeyServiceAccount
	} else {
		cDes.KmsKeyServiceAccount = des.KmsKeyServiceAccount
	}

	return cDes
}

func canonicalizeNewImageImageEncryptionKey(c *Client, des, nw *ImageImageEncryptionKey) *ImageImageEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RawKey, nw.RawKey) {
		nw.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Sha256, nw.Sha256) {
		nw.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyServiceAccount, nw.KmsKeyServiceAccount) {
		nw.KmsKeyServiceAccount = des.KmsKeyServiceAccount
	}

	return nw
}

func canonicalizeNewImageImageEncryptionKeySet(c *Client, des, nw []ImageImageEncryptionKey) []ImageImageEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageImageEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageImageEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageImageEncryptionKeySlice(c *Client, des, nw []ImageImageEncryptionKey) []ImageImageEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageImageEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageImageEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeImageRawDisk(des, initial *ImageRawDisk, opts ...dcl.ApplyOption) *ImageRawDisk {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.ContainerType) {
		des.ContainerType = ImageRawDiskContainerTypeEnumRef("TAR")
	}

	if initial == nil {
		return des
	}

	cDes := &ImageRawDisk{}

	if dcl.StringCanonicalize(des.Source, initial.Source) || dcl.IsZeroValue(des.Source) {
		cDes.Source = initial.Source
	} else {
		cDes.Source = des.Source
	}
	if dcl.StringCanonicalize(des.Sha1Checksum, initial.Sha1Checksum) || dcl.IsZeroValue(des.Sha1Checksum) {
		cDes.Sha1Checksum = initial.Sha1Checksum
	} else {
		cDes.Sha1Checksum = des.Sha1Checksum
	}
	if dcl.IsZeroValue(des.ContainerType) {
		des.ContainerType = initial.ContainerType
	} else {
		cDes.ContainerType = des.ContainerType
	}

	return cDes
}

func canonicalizeNewImageRawDisk(c *Client, des, nw *ImageRawDisk) *ImageRawDisk {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ContainerType) {
		nw.ContainerType = ImageRawDiskContainerTypeEnumRef("TAR")
	}

	if dcl.StringCanonicalize(des.Source, nw.Source) {
		nw.Source = des.Source
	}
	if dcl.StringCanonicalize(des.Sha1Checksum, nw.Sha1Checksum) {
		nw.Sha1Checksum = des.Sha1Checksum
	}
	if dcl.IsZeroValue(nw.ContainerType) {
		nw.ContainerType = des.ContainerType
	}

	return nw
}

func canonicalizeNewImageRawDiskSet(c *Client, des, nw []ImageRawDisk) []ImageRawDisk {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageRawDisk
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageRawDiskNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageRawDiskSlice(c *Client, des, nw []ImageRawDisk) []ImageRawDisk {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageRawDisk
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageRawDisk(c, &d, &n))
	}

	return items
}

func canonicalizeImageShieldedInstanceInitialState(des, initial *ImageShieldedInstanceInitialState, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageShieldedInstanceInitialState{}

	cDes.Pk = canonicalizeImageShieldedInstanceInitialStatePk(des.Pk, initial.Pk, opts...)
	if dcl.IsZeroValue(des.Kek) {
		des.Kek = initial.Kek
	} else {
		cDes.Kek = des.Kek
	}
	if dcl.IsZeroValue(des.Db) {
		des.Db = initial.Db
	} else {
		cDes.Db = des.Db
	}
	if dcl.IsZeroValue(des.Dbx) {
		des.Dbx = initial.Dbx
	} else {
		cDes.Dbx = des.Dbx
	}

	return cDes
}

func canonicalizeNewImageShieldedInstanceInitialState(c *Client, des, nw *ImageShieldedInstanceInitialState) *ImageShieldedInstanceInitialState {
	if des == nil || nw == nil {
		return nw
	}

	nw.Pk = canonicalizeNewImageShieldedInstanceInitialStatePk(c, des.Pk, nw.Pk)
	nw.Kek = canonicalizeNewImageShieldedInstanceInitialStateKekSlice(c, des.Kek, nw.Kek)
	nw.Db = canonicalizeNewImageShieldedInstanceInitialStateDbSlice(c, des.Db, nw.Db)
	nw.Dbx = canonicalizeNewImageShieldedInstanceInitialStateDbxSlice(c, des.Dbx, nw.Dbx)

	return nw
}

func canonicalizeNewImageShieldedInstanceInitialStateSet(c *Client, des, nw []ImageShieldedInstanceInitialState) []ImageShieldedInstanceInitialState {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageShieldedInstanceInitialState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageShieldedInstanceInitialStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageShieldedInstanceInitialStateSlice(c *Client, des, nw []ImageShieldedInstanceInitialState) []ImageShieldedInstanceInitialState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageShieldedInstanceInitialState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageShieldedInstanceInitialState(c, &d, &n))
	}

	return items
}

func canonicalizeImageShieldedInstanceInitialStatePk(des, initial *ImageShieldedInstanceInitialStatePk, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialStatePk {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageShieldedInstanceInitialStatePk{}

	if dcl.StringCanonicalize(des.Content, initial.Content) || dcl.IsZeroValue(des.Content) {
		cDes.Content = initial.Content
	} else {
		cDes.Content = des.Content
	}
	if dcl.IsZeroValue(des.FileType) {
		des.FileType = initial.FileType
	} else {
		cDes.FileType = des.FileType
	}

	return cDes
}

func canonicalizeNewImageShieldedInstanceInitialStatePk(c *Client, des, nw *ImageShieldedInstanceInitialStatePk) *ImageShieldedInstanceInitialStatePk {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}
	if dcl.IsZeroValue(nw.FileType) {
		nw.FileType = des.FileType
	}

	return nw
}

func canonicalizeNewImageShieldedInstanceInitialStatePkSet(c *Client, des, nw []ImageShieldedInstanceInitialStatePk) []ImageShieldedInstanceInitialStatePk {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageShieldedInstanceInitialStatePk
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageShieldedInstanceInitialStatePkNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageShieldedInstanceInitialStatePkSlice(c *Client, des, nw []ImageShieldedInstanceInitialStatePk) []ImageShieldedInstanceInitialStatePk {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageShieldedInstanceInitialStatePk
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageShieldedInstanceInitialStatePk(c, &d, &n))
	}

	return items
}

func canonicalizeImageShieldedInstanceInitialStateKek(des, initial *ImageShieldedInstanceInitialStateKek, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialStateKek {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageShieldedInstanceInitialStateKek{}

	if dcl.StringCanonicalize(des.Content, initial.Content) || dcl.IsZeroValue(des.Content) {
		cDes.Content = initial.Content
	} else {
		cDes.Content = des.Content
	}
	if dcl.IsZeroValue(des.FileType) {
		des.FileType = initial.FileType
	} else {
		cDes.FileType = des.FileType
	}

	return cDes
}

func canonicalizeNewImageShieldedInstanceInitialStateKek(c *Client, des, nw *ImageShieldedInstanceInitialStateKek) *ImageShieldedInstanceInitialStateKek {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}
	if dcl.IsZeroValue(nw.FileType) {
		nw.FileType = des.FileType
	}

	return nw
}

func canonicalizeNewImageShieldedInstanceInitialStateKekSet(c *Client, des, nw []ImageShieldedInstanceInitialStateKek) []ImageShieldedInstanceInitialStateKek {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageShieldedInstanceInitialStateKek
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageShieldedInstanceInitialStateKekNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageShieldedInstanceInitialStateKekSlice(c *Client, des, nw []ImageShieldedInstanceInitialStateKek) []ImageShieldedInstanceInitialStateKek {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageShieldedInstanceInitialStateKek
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageShieldedInstanceInitialStateKek(c, &d, &n))
	}

	return items
}

func canonicalizeImageShieldedInstanceInitialStateDb(des, initial *ImageShieldedInstanceInitialStateDb, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialStateDb {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageShieldedInstanceInitialStateDb{}

	if dcl.StringCanonicalize(des.Content, initial.Content) || dcl.IsZeroValue(des.Content) {
		cDes.Content = initial.Content
	} else {
		cDes.Content = des.Content
	}
	if dcl.IsZeroValue(des.FileType) {
		des.FileType = initial.FileType
	} else {
		cDes.FileType = des.FileType
	}

	return cDes
}

func canonicalizeNewImageShieldedInstanceInitialStateDb(c *Client, des, nw *ImageShieldedInstanceInitialStateDb) *ImageShieldedInstanceInitialStateDb {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}
	if dcl.IsZeroValue(nw.FileType) {
		nw.FileType = des.FileType
	}

	return nw
}

func canonicalizeNewImageShieldedInstanceInitialStateDbSet(c *Client, des, nw []ImageShieldedInstanceInitialStateDb) []ImageShieldedInstanceInitialStateDb {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageShieldedInstanceInitialStateDb
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageShieldedInstanceInitialStateDbNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageShieldedInstanceInitialStateDbSlice(c *Client, des, nw []ImageShieldedInstanceInitialStateDb) []ImageShieldedInstanceInitialStateDb {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageShieldedInstanceInitialStateDb
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageShieldedInstanceInitialStateDb(c, &d, &n))
	}

	return items
}

func canonicalizeImageShieldedInstanceInitialStateDbx(des, initial *ImageShieldedInstanceInitialStateDbx, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialStateDbx {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageShieldedInstanceInitialStateDbx{}

	if dcl.StringCanonicalize(des.Content, initial.Content) || dcl.IsZeroValue(des.Content) {
		cDes.Content = initial.Content
	} else {
		cDes.Content = des.Content
	}
	if dcl.IsZeroValue(des.FileType) {
		des.FileType = initial.FileType
	} else {
		cDes.FileType = des.FileType
	}

	return cDes
}

func canonicalizeNewImageShieldedInstanceInitialStateDbx(c *Client, des, nw *ImageShieldedInstanceInitialStateDbx) *ImageShieldedInstanceInitialStateDbx {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}
	if dcl.IsZeroValue(nw.FileType) {
		nw.FileType = des.FileType
	}

	return nw
}

func canonicalizeNewImageShieldedInstanceInitialStateDbxSet(c *Client, des, nw []ImageShieldedInstanceInitialStateDbx) []ImageShieldedInstanceInitialStateDbx {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageShieldedInstanceInitialStateDbx
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageShieldedInstanceInitialStateDbxNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageShieldedInstanceInitialStateDbxSlice(c *Client, des, nw []ImageShieldedInstanceInitialStateDbx) []ImageShieldedInstanceInitialStateDbx {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageShieldedInstanceInitialStateDbx
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageShieldedInstanceInitialStateDbx(c, &d, &n))
	}

	return items
}

func canonicalizeImageSourceDiskEncryptionKey(des, initial *ImageSourceDiskEncryptionKey, opts ...dcl.ApplyOption) *ImageSourceDiskEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageSourceDiskEncryptionKey{}

	if dcl.StringCanonicalize(des.RawKey, initial.RawKey) || dcl.IsZeroValue(des.RawKey) {
		cDes.RawKey = initial.RawKey
	} else {
		cDes.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Sha256, initial.Sha256) || dcl.IsZeroValue(des.Sha256) {
		cDes.Sha256 = initial.Sha256
	} else {
		cDes.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyServiceAccount, initial.KmsKeyServiceAccount) || dcl.IsZeroValue(des.KmsKeyServiceAccount) {
		cDes.KmsKeyServiceAccount = initial.KmsKeyServiceAccount
	} else {
		cDes.KmsKeyServiceAccount = des.KmsKeyServiceAccount
	}

	return cDes
}

func canonicalizeNewImageSourceDiskEncryptionKey(c *Client, des, nw *ImageSourceDiskEncryptionKey) *ImageSourceDiskEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RawKey, nw.RawKey) {
		nw.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Sha256, nw.Sha256) {
		nw.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyServiceAccount, nw.KmsKeyServiceAccount) {
		nw.KmsKeyServiceAccount = des.KmsKeyServiceAccount
	}

	return nw
}

func canonicalizeNewImageSourceDiskEncryptionKeySet(c *Client, des, nw []ImageSourceDiskEncryptionKey) []ImageSourceDiskEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageSourceDiskEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageSourceDiskEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageSourceDiskEncryptionKeySlice(c *Client, des, nw []ImageSourceDiskEncryptionKey) []ImageSourceDiskEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageSourceDiskEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageSourceDiskEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeImageSourceImageEncryptionKey(des, initial *ImageSourceImageEncryptionKey, opts ...dcl.ApplyOption) *ImageSourceImageEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageSourceImageEncryptionKey{}

	if dcl.StringCanonicalize(des.RawKey, initial.RawKey) || dcl.IsZeroValue(des.RawKey) {
		cDes.RawKey = initial.RawKey
	} else {
		cDes.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Sha256, initial.Sha256) || dcl.IsZeroValue(des.Sha256) {
		cDes.Sha256 = initial.Sha256
	} else {
		cDes.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyServiceAccount, initial.KmsKeyServiceAccount) || dcl.IsZeroValue(des.KmsKeyServiceAccount) {
		cDes.KmsKeyServiceAccount = initial.KmsKeyServiceAccount
	} else {
		cDes.KmsKeyServiceAccount = des.KmsKeyServiceAccount
	}

	return cDes
}

func canonicalizeNewImageSourceImageEncryptionKey(c *Client, des, nw *ImageSourceImageEncryptionKey) *ImageSourceImageEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RawKey, nw.RawKey) {
		nw.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Sha256, nw.Sha256) {
		nw.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyServiceAccount, nw.KmsKeyServiceAccount) {
		nw.KmsKeyServiceAccount = des.KmsKeyServiceAccount
	}

	return nw
}

func canonicalizeNewImageSourceImageEncryptionKeySet(c *Client, des, nw []ImageSourceImageEncryptionKey) []ImageSourceImageEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageSourceImageEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageSourceImageEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageSourceImageEncryptionKeySlice(c *Client, des, nw []ImageSourceImageEncryptionKey) []ImageSourceImageEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageSourceImageEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageSourceImageEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeImageSourceSnapshotEncryptionKey(des, initial *ImageSourceSnapshotEncryptionKey, opts ...dcl.ApplyOption) *ImageSourceSnapshotEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageSourceSnapshotEncryptionKey{}

	if dcl.StringCanonicalize(des.RawKey, initial.RawKey) || dcl.IsZeroValue(des.RawKey) {
		cDes.RawKey = initial.RawKey
	} else {
		cDes.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Sha256, initial.Sha256) || dcl.IsZeroValue(des.Sha256) {
		cDes.Sha256 = initial.Sha256
	} else {
		cDes.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyServiceAccount, initial.KmsKeyServiceAccount) || dcl.IsZeroValue(des.KmsKeyServiceAccount) {
		cDes.KmsKeyServiceAccount = initial.KmsKeyServiceAccount
	} else {
		cDes.KmsKeyServiceAccount = des.KmsKeyServiceAccount
	}

	return cDes
}

func canonicalizeNewImageSourceSnapshotEncryptionKey(c *Client, des, nw *ImageSourceSnapshotEncryptionKey) *ImageSourceSnapshotEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RawKey, nw.RawKey) {
		nw.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Sha256, nw.Sha256) {
		nw.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyServiceAccount, nw.KmsKeyServiceAccount) {
		nw.KmsKeyServiceAccount = des.KmsKeyServiceAccount
	}

	return nw
}

func canonicalizeNewImageSourceSnapshotEncryptionKeySet(c *Client, des, nw []ImageSourceSnapshotEncryptionKey) []ImageSourceSnapshotEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageSourceSnapshotEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageSourceSnapshotEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageSourceSnapshotEncryptionKeySlice(c *Client, des, nw []ImageSourceSnapshotEncryptionKey) []ImageSourceSnapshotEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageSourceSnapshotEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageSourceSnapshotEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeImageDeprecated(des, initial *ImageDeprecated, opts ...dcl.ApplyOption) *ImageDeprecated {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ImageDeprecated{}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	} else {
		cDes.State = des.State
	}
	if dcl.StringCanonicalize(des.Replacement, initial.Replacement) || dcl.IsZeroValue(des.Replacement) {
		cDes.Replacement = initial.Replacement
	} else {
		cDes.Replacement = des.Replacement
	}
	if dcl.StringCanonicalize(des.Deprecated, initial.Deprecated) || dcl.IsZeroValue(des.Deprecated) {
		cDes.Deprecated = initial.Deprecated
	} else {
		cDes.Deprecated = des.Deprecated
	}
	if dcl.StringCanonicalize(des.Obsolete, initial.Obsolete) || dcl.IsZeroValue(des.Obsolete) {
		cDes.Obsolete = initial.Obsolete
	} else {
		cDes.Obsolete = des.Obsolete
	}
	if dcl.StringCanonicalize(des.Deleted, initial.Deleted) || dcl.IsZeroValue(des.Deleted) {
		cDes.Deleted = initial.Deleted
	} else {
		cDes.Deleted = des.Deleted
	}

	return cDes
}

func canonicalizeNewImageDeprecated(c *Client, des, nw *ImageDeprecated) *ImageDeprecated {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.State) {
		nw.State = des.State
	}
	if dcl.StringCanonicalize(des.Replacement, nw.Replacement) {
		nw.Replacement = des.Replacement
	}
	if dcl.StringCanonicalize(des.Deprecated, nw.Deprecated) {
		nw.Deprecated = des.Deprecated
	}
	if dcl.StringCanonicalize(des.Obsolete, nw.Obsolete) {
		nw.Obsolete = des.Obsolete
	}
	if dcl.StringCanonicalize(des.Deleted, nw.Deleted) {
		nw.Deleted = des.Deleted
	}

	return nw
}

func canonicalizeNewImageDeprecatedSet(c *Client, des, nw []ImageDeprecated) []ImageDeprecated {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageDeprecated
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareImageDeprecatedNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewImageDeprecatedSlice(c *Client, des, nw []ImageDeprecated) []ImageDeprecated {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ImageDeprecated
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewImageDeprecated(c, &d, &n))
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
func diffImage(c *Client, desired, actual *Image, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.ArchiveSizeBytes, actual.ArchiveSizeBytes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArchiveSizeBytes")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.DiskSizeGb, actual.DiskSizeGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Family, actual.Family, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Family")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GuestOSFeature, actual.GuestOSFeature, dcl.Info{Type: "Set", ObjectFunction: compareImageGuestOSFeatureNewStyle, EmptyObject: EmptyImageGuestOSFeature, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GuestOsFeatures")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImageEncryptionKey, actual.ImageEncryptionKey, dcl.Info{ObjectFunction: compareImageImageEncryptionKeyNewStyle, EmptyObject: EmptyImageImageEncryptionKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImageEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateImageSetLabelsOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.License, actual.License, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Licenses")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.RawDisk, actual.RawDisk, dcl.Info{Ignore: true, ObjectFunction: compareImageRawDiskNewStyle, EmptyObject: EmptyImageRawDisk, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawDisk")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ShieldedInstanceInitialState, actual.ShieldedInstanceInitialState, dcl.Info{ObjectFunction: compareImageShieldedInstanceInitialStateNewStyle, EmptyObject: EmptyImageShieldedInstanceInitialState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ShieldedInstanceInitialState")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.SourceDisk, actual.SourceDisk, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceDisk")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceDiskEncryptionKey, actual.SourceDiskEncryptionKey, dcl.Info{ObjectFunction: compareImageSourceDiskEncryptionKeyNewStyle, EmptyObject: EmptyImageSourceDiskEncryptionKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceDiskEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceDiskId, actual.SourceDiskId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceDiskId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceImage, actual.SourceImage, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceImage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceImageEncryptionKey, actual.SourceImageEncryptionKey, dcl.Info{ObjectFunction: compareImageSourceImageEncryptionKeyNewStyle, EmptyObject: EmptyImageSourceImageEncryptionKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceImageEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceImageId, actual.SourceImageId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceImageId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceSnapshot, actual.SourceSnapshot, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceSnapshot")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceSnapshotEncryptionKey, actual.SourceSnapshotEncryptionKey, dcl.Info{ObjectFunction: compareImageSourceSnapshotEncryptionKeyNewStyle, EmptyObject: EmptyImageSourceSnapshotEncryptionKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceSnapshotEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceSnapshotId, actual.SourceSnapshotId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceSnapshotId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceType, actual.SourceType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StorageLocation, actual.StorageLocation, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StorageLocations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deprecated, actual.Deprecated, dcl.Info{ObjectFunction: compareImageDeprecatedNewStyle, EmptyObject: EmptyImageDeprecated, OperationSelector: dcl.TriggersOperation("updateImageDeprecateOperation")}, fn.AddNest("Deprecated")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareImageGuestOSFeatureNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageGuestOSFeature)
	if !ok {
		desiredNotPointer, ok := d.(ImageGuestOSFeature)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageGuestOSFeature or *ImageGuestOSFeature", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageGuestOSFeature)
	if !ok {
		actualNotPointer, ok := a.(ImageGuestOSFeature)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageGuestOSFeature", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageImageEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageImageEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(ImageImageEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageImageEncryptionKey or *ImageImageEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageImageEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(ImageImageEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageImageEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256, actual.Sha256, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyServiceAccount, actual.KmsKeyServiceAccount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageRawDiskNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageRawDisk)
	if !ok {
		desiredNotPointer, ok := d.(ImageRawDisk)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageRawDisk or *ImageRawDisk", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageRawDisk)
	if !ok {
		actualNotPointer, ok := a.(ImageRawDisk)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageRawDisk", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Source, actual.Source, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Source")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha1Checksum, actual.Sha1Checksum, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha1Checksum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContainerType, actual.ContainerType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ContainerType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageShieldedInstanceInitialStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageShieldedInstanceInitialState)
	if !ok {
		desiredNotPointer, ok := d.(ImageShieldedInstanceInitialState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialState or *ImageShieldedInstanceInitialState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageShieldedInstanceInitialState)
	if !ok {
		actualNotPointer, ok := a.(ImageShieldedInstanceInitialState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pk, actual.Pk, dcl.Info{ObjectFunction: compareImageShieldedInstanceInitialStatePkNewStyle, EmptyObject: EmptyImageShieldedInstanceInitialStatePk, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Pk")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kek, actual.Kek, dcl.Info{ObjectFunction: compareImageShieldedInstanceInitialStateKekNewStyle, EmptyObject: EmptyImageShieldedInstanceInitialStateKek, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Keks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Db, actual.Db, dcl.Info{ObjectFunction: compareImageShieldedInstanceInitialStateDbNewStyle, EmptyObject: EmptyImageShieldedInstanceInitialStateDb, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Dbs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dbx, actual.Dbx, dcl.Info{ObjectFunction: compareImageShieldedInstanceInitialStateDbxNewStyle, EmptyObject: EmptyImageShieldedInstanceInitialStateDbx, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Dbxs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageShieldedInstanceInitialStatePkNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageShieldedInstanceInitialStatePk)
	if !ok {
		desiredNotPointer, ok := d.(ImageShieldedInstanceInitialStatePk)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialStatePk or *ImageShieldedInstanceInitialStatePk", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageShieldedInstanceInitialStatePk)
	if !ok {
		actualNotPointer, ok := a.(ImageShieldedInstanceInitialStatePk)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialStatePk", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FileType, actual.FileType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FileType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageShieldedInstanceInitialStateKekNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageShieldedInstanceInitialStateKek)
	if !ok {
		desiredNotPointer, ok := d.(ImageShieldedInstanceInitialStateKek)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialStateKek or *ImageShieldedInstanceInitialStateKek", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageShieldedInstanceInitialStateKek)
	if !ok {
		actualNotPointer, ok := a.(ImageShieldedInstanceInitialStateKek)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialStateKek", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FileType, actual.FileType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FileType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageShieldedInstanceInitialStateDbNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageShieldedInstanceInitialStateDb)
	if !ok {
		desiredNotPointer, ok := d.(ImageShieldedInstanceInitialStateDb)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialStateDb or *ImageShieldedInstanceInitialStateDb", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageShieldedInstanceInitialStateDb)
	if !ok {
		actualNotPointer, ok := a.(ImageShieldedInstanceInitialStateDb)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialStateDb", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FileType, actual.FileType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FileType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageShieldedInstanceInitialStateDbxNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageShieldedInstanceInitialStateDbx)
	if !ok {
		desiredNotPointer, ok := d.(ImageShieldedInstanceInitialStateDbx)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialStateDbx or *ImageShieldedInstanceInitialStateDbx", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageShieldedInstanceInitialStateDbx)
	if !ok {
		actualNotPointer, ok := a.(ImageShieldedInstanceInitialStateDbx)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageShieldedInstanceInitialStateDbx", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FileType, actual.FileType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FileType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageSourceDiskEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageSourceDiskEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(ImageSourceDiskEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageSourceDiskEncryptionKey or *ImageSourceDiskEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageSourceDiskEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(ImageSourceDiskEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageSourceDiskEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256, actual.Sha256, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyServiceAccount, actual.KmsKeyServiceAccount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageSourceImageEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageSourceImageEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(ImageSourceImageEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageSourceImageEncryptionKey or *ImageSourceImageEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageSourceImageEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(ImageSourceImageEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageSourceImageEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256, actual.Sha256, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyServiceAccount, actual.KmsKeyServiceAccount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageSourceSnapshotEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageSourceSnapshotEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(ImageSourceSnapshotEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageSourceSnapshotEncryptionKey or *ImageSourceSnapshotEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageSourceSnapshotEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(ImageSourceSnapshotEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageSourceSnapshotEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256, actual.Sha256, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyServiceAccount, actual.KmsKeyServiceAccount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareImageDeprecatedNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ImageDeprecated)
	if !ok {
		desiredNotPointer, ok := d.(ImageDeprecated)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageDeprecated or *ImageDeprecated", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ImageDeprecated)
	if !ok {
		actualNotPointer, ok := a.(ImageDeprecated)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ImageDeprecated", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Replacement, actual.Replacement, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Replacement")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deprecated, actual.Deprecated, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Deprecated")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Obsolete, actual.Obsolete, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Obsolete")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deleted, actual.Deleted, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Deleted")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Image) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Image) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Image) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Image) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "deprecate" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/images/{{name}}/deprecate", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	if updateName == "setLabels" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/images/{{name}}/setLabels", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Image resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Image) marshal(c *Client) ([]byte, error) {
	m, err := expandImage(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Image: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalImage decodes JSON responses into the Image resource schema.
func unmarshalImage(b []byte, c *Client) (*Image, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapImage(m, c)
}

func unmarshalMapImage(m map[string]interface{}, c *Client) (*Image, error) {

	return flattenImage(c, m), nil
}

// expandImage expands Image into a JSON request object.
func expandImage(c *Client, f *Image) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.ArchiveSizeBytes; !dcl.IsEmptyValueIndirect(v) {
		m["archiveSizeBytes"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.DiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskSizeGb"] = v
	}
	if v := f.Family; !dcl.IsEmptyValueIndirect(v) {
		m["family"] = v
	}
	if v, err := expandImageGuestOSFeatureSlice(c, f.GuestOSFeature); err != nil {
		return nil, fmt.Errorf("error expanding GuestOSFeature into guestOsFeatures: %w", err)
	} else {
		m["guestOsFeatures"] = v
	}
	if v, err := expandImageImageEncryptionKey(c, f.ImageEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding ImageEncryptionKey into imageEncryptionKey: %w", err)
	} else if v != nil {
		m["imageEncryptionKey"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	m["licenses"] = f.License
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandImageRawDisk(c, f.RawDisk); err != nil {
		return nil, fmt.Errorf("error expanding RawDisk into rawDisk: %w", err)
	} else if v != nil {
		m["rawDisk"] = v
	}
	if v, err := expandImageShieldedInstanceInitialState(c, f.ShieldedInstanceInitialState); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedInstanceInitialState into shieldedInstanceInitialState: %w", err)
	} else if v != nil {
		m["shieldedInstanceInitialState"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.SourceDisk; !dcl.IsEmptyValueIndirect(v) {
		m["sourceDisk"] = v
	}
	if v, err := expandImageSourceDiskEncryptionKey(c, f.SourceDiskEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding SourceDiskEncryptionKey into sourceDiskEncryptionKey: %w", err)
	} else if v != nil {
		m["sourceDiskEncryptionKey"] = v
	}
	if v := f.SourceDiskId; !dcl.IsEmptyValueIndirect(v) {
		m["sourceDiskId"] = v
	}
	if v := f.SourceImage; !dcl.IsEmptyValueIndirect(v) {
		m["sourceImage"] = v
	}
	if v, err := expandImageSourceImageEncryptionKey(c, f.SourceImageEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding SourceImageEncryptionKey into sourceImageEncryptionKey: %w", err)
	} else if v != nil {
		m["sourceImageEncryptionKey"] = v
	}
	if v := f.SourceImageId; !dcl.IsEmptyValueIndirect(v) {
		m["sourceImageId"] = v
	}
	if v := f.SourceSnapshot; !dcl.IsEmptyValueIndirect(v) {
		m["sourceSnapshot"] = v
	}
	if v, err := expandImageSourceSnapshotEncryptionKey(c, f.SourceSnapshotEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding SourceSnapshotEncryptionKey into sourceSnapshotEncryptionKey: %w", err)
	} else if v != nil {
		m["sourceSnapshotEncryptionKey"] = v
	}
	if v := f.SourceSnapshotId; !dcl.IsEmptyValueIndirect(v) {
		m["sourceSnapshotId"] = v
	}
	if v := f.SourceType; !dcl.IsEmptyValueIndirect(v) {
		m["sourceType"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	m["storageLocations"] = f.StorageLocation
	if v, err := expandImageDeprecated(c, f.Deprecated); err != nil {
		return nil, fmt.Errorf("error expanding Deprecated into deprecated: %w", err)
	} else if v != nil {
		m["deprecated"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenImage flattens Image from a JSON request object into the
// Image type.
func flattenImage(c *Client, i interface{}) *Image {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Image{}
	res.ArchiveSizeBytes = dcl.FlattenInteger(m["archiveSizeBytes"])
	res.Description = dcl.FlattenString(m["description"])
	res.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	res.Family = dcl.FlattenString(m["family"])
	res.GuestOSFeature = flattenImageGuestOSFeatureSlice(c, m["guestOsFeatures"])
	res.ImageEncryptionKey = flattenImageImageEncryptionKey(c, m["imageEncryptionKey"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.License = dcl.FlattenStringSlice(m["licenses"])
	res.Name = dcl.FlattenString(m["name"])
	res.RawDisk = flattenImageRawDisk(c, m["rawDisk"])
	res.ShieldedInstanceInitialState = flattenImageShieldedInstanceInitialState(c, m["shieldedInstanceInitialState"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.SourceDisk = dcl.FlattenString(m["sourceDisk"])
	res.SourceDiskEncryptionKey = flattenImageSourceDiskEncryptionKey(c, m["sourceDiskEncryptionKey"])
	res.SourceDiskId = dcl.FlattenString(m["sourceDiskId"])
	res.SourceImage = dcl.FlattenString(m["sourceImage"])
	res.SourceImageEncryptionKey = flattenImageSourceImageEncryptionKey(c, m["sourceImageEncryptionKey"])
	res.SourceImageId = dcl.FlattenString(m["sourceImageId"])
	res.SourceSnapshot = dcl.FlattenString(m["sourceSnapshot"])
	res.SourceSnapshotEncryptionKey = flattenImageSourceSnapshotEncryptionKey(c, m["sourceSnapshotEncryptionKey"])
	res.SourceSnapshotId = dcl.FlattenString(m["sourceSnapshotId"])
	res.SourceType = flattenImageSourceTypeEnum(m["sourceType"])
	res.Status = flattenImageStatusEnum(m["status"])
	res.StorageLocation = dcl.FlattenStringSlice(m["storageLocations"])
	res.Deprecated = flattenImageDeprecated(c, m["deprecated"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandImageGuestOSFeatureMap expands the contents of ImageGuestOSFeature into a JSON
// request object.
func expandImageGuestOSFeatureMap(c *Client, f map[string]ImageGuestOSFeature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageGuestOSFeature(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageGuestOSFeatureSlice expands the contents of ImageGuestOSFeature into a JSON
// request object.
func expandImageGuestOSFeatureSlice(c *Client, f []ImageGuestOSFeature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageGuestOSFeature(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageGuestOSFeatureMap flattens the contents of ImageGuestOSFeature from a JSON
// response object.
func flattenImageGuestOSFeatureMap(c *Client, i interface{}) map[string]ImageGuestOSFeature {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageGuestOSFeature{}
	}

	if len(a) == 0 {
		return map[string]ImageGuestOSFeature{}
	}

	items := make(map[string]ImageGuestOSFeature)
	for k, item := range a {
		items[k] = *flattenImageGuestOSFeature(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageGuestOSFeatureSlice flattens the contents of ImageGuestOSFeature from a JSON
// response object.
func flattenImageGuestOSFeatureSlice(c *Client, i interface{}) []ImageGuestOSFeature {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageGuestOSFeature{}
	}

	if len(a) == 0 {
		return []ImageGuestOSFeature{}
	}

	items := make([]ImageGuestOSFeature, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageGuestOSFeature(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageGuestOSFeature expands an instance of ImageGuestOSFeature into a JSON
// request object.
func expandImageGuestOSFeature(c *Client, f *ImageGuestOSFeature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenImageGuestOSFeature flattens an instance of ImageGuestOSFeature from a JSON
// response object.
func flattenImageGuestOSFeature(c *Client, i interface{}) *ImageGuestOSFeature {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageGuestOSFeature{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageGuestOSFeature
	}
	r.Type = flattenImageGuestOSFeatureTypeEnum(m["type"])

	return r
}

// expandImageImageEncryptionKeyMap expands the contents of ImageImageEncryptionKey into a JSON
// request object.
func expandImageImageEncryptionKeyMap(c *Client, f map[string]ImageImageEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageImageEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageImageEncryptionKeySlice expands the contents of ImageImageEncryptionKey into a JSON
// request object.
func expandImageImageEncryptionKeySlice(c *Client, f []ImageImageEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageImageEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageImageEncryptionKeyMap flattens the contents of ImageImageEncryptionKey from a JSON
// response object.
func flattenImageImageEncryptionKeyMap(c *Client, i interface{}) map[string]ImageImageEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageImageEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]ImageImageEncryptionKey{}
	}

	items := make(map[string]ImageImageEncryptionKey)
	for k, item := range a {
		items[k] = *flattenImageImageEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageImageEncryptionKeySlice flattens the contents of ImageImageEncryptionKey from a JSON
// response object.
func flattenImageImageEncryptionKeySlice(c *Client, i interface{}) []ImageImageEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageImageEncryptionKey{}
	}

	if len(a) == 0 {
		return []ImageImageEncryptionKey{}
	}

	items := make([]ImageImageEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageImageEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageImageEncryptionKey expands an instance of ImageImageEncryptionKey into a JSON
// request object.
func expandImageImageEncryptionKey(c *Client, f *ImageImageEncryptionKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RawKey; !dcl.IsEmptyValueIndirect(v) {
		m["rawKey"] = v
	}
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}
	if v := f.Sha256; !dcl.IsEmptyValueIndirect(v) {
		m["sha256"] = v
	}
	if v := f.KmsKeyServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyServiceAccount"] = v
	}

	return m, nil
}

// flattenImageImageEncryptionKey flattens an instance of ImageImageEncryptionKey from a JSON
// response object.
func flattenImageImageEncryptionKey(c *Client, i interface{}) *ImageImageEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageImageEncryptionKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageImageEncryptionKey
	}
	r.RawKey = dcl.FlattenString(m["rawKey"])
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	r.Sha256 = dcl.FlattenString(m["sha256"])
	r.KmsKeyServiceAccount = dcl.FlattenString(m["kmsKeyServiceAccount"])

	return r
}

// expandImageRawDiskMap expands the contents of ImageRawDisk into a JSON
// request object.
func expandImageRawDiskMap(c *Client, f map[string]ImageRawDisk) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageRawDisk(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageRawDiskSlice expands the contents of ImageRawDisk into a JSON
// request object.
func expandImageRawDiskSlice(c *Client, f []ImageRawDisk) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageRawDisk(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageRawDiskMap flattens the contents of ImageRawDisk from a JSON
// response object.
func flattenImageRawDiskMap(c *Client, i interface{}) map[string]ImageRawDisk {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageRawDisk{}
	}

	if len(a) == 0 {
		return map[string]ImageRawDisk{}
	}

	items := make(map[string]ImageRawDisk)
	for k, item := range a {
		items[k] = *flattenImageRawDisk(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageRawDiskSlice flattens the contents of ImageRawDisk from a JSON
// response object.
func flattenImageRawDiskSlice(c *Client, i interface{}) []ImageRawDisk {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageRawDisk{}
	}

	if len(a) == 0 {
		return []ImageRawDisk{}
	}

	items := make([]ImageRawDisk, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageRawDisk(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageRawDisk expands an instance of ImageRawDisk into a JSON
// request object.
func expandImageRawDisk(c *Client, f *ImageRawDisk) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Source; !dcl.IsEmptyValueIndirect(v) {
		m["source"] = v
	}
	if v := f.Sha1Checksum; !dcl.IsEmptyValueIndirect(v) {
		m["sha1Checksum"] = v
	}
	if v := f.ContainerType; !dcl.IsEmptyValueIndirect(v) {
		m["containerType"] = v
	}

	return m, nil
}

// flattenImageRawDisk flattens an instance of ImageRawDisk from a JSON
// response object.
func flattenImageRawDisk(c *Client, i interface{}) *ImageRawDisk {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageRawDisk{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageRawDisk
	}
	r.Source = dcl.FlattenString(m["source"])
	r.Sha1Checksum = dcl.FlattenString(m["sha1Checksum"])
	r.ContainerType = flattenImageRawDiskContainerTypeEnum(m["containerType"])
	if dcl.IsEmptyValueIndirect(m["containerType"]) {
		c.Config.Logger.Info("Using default value for containerType.")
		r.ContainerType = ImageRawDiskContainerTypeEnumRef("TAR")
	}

	return r
}

// expandImageShieldedInstanceInitialStateMap expands the contents of ImageShieldedInstanceInitialState into a JSON
// request object.
func expandImageShieldedInstanceInitialStateMap(c *Client, f map[string]ImageShieldedInstanceInitialState) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageShieldedInstanceInitialState(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageShieldedInstanceInitialStateSlice expands the contents of ImageShieldedInstanceInitialState into a JSON
// request object.
func expandImageShieldedInstanceInitialStateSlice(c *Client, f []ImageShieldedInstanceInitialState) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageShieldedInstanceInitialState(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageShieldedInstanceInitialStateMap flattens the contents of ImageShieldedInstanceInitialState from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateMap(c *Client, i interface{}) map[string]ImageShieldedInstanceInitialState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageShieldedInstanceInitialState{}
	}

	if len(a) == 0 {
		return map[string]ImageShieldedInstanceInitialState{}
	}

	items := make(map[string]ImageShieldedInstanceInitialState)
	for k, item := range a {
		items[k] = *flattenImageShieldedInstanceInitialState(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateSlice flattens the contents of ImageShieldedInstanceInitialState from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateSlice(c *Client, i interface{}) []ImageShieldedInstanceInitialState {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageShieldedInstanceInitialState{}
	}

	if len(a) == 0 {
		return []ImageShieldedInstanceInitialState{}
	}

	items := make([]ImageShieldedInstanceInitialState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageShieldedInstanceInitialState(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageShieldedInstanceInitialState expands an instance of ImageShieldedInstanceInitialState into a JSON
// request object.
func expandImageShieldedInstanceInitialState(c *Client, f *ImageShieldedInstanceInitialState) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandImageShieldedInstanceInitialStatePk(c, f.Pk); err != nil {
		return nil, fmt.Errorf("error expanding Pk into pk: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pk"] = v
	}
	if v, err := expandImageShieldedInstanceInitialStateKekSlice(c, f.Kek); err != nil {
		return nil, fmt.Errorf("error expanding Kek into keks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["keks"] = v
	}
	if v, err := expandImageShieldedInstanceInitialStateDbSlice(c, f.Db); err != nil {
		return nil, fmt.Errorf("error expanding Db into dbs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dbs"] = v
	}
	if v, err := expandImageShieldedInstanceInitialStateDbxSlice(c, f.Dbx); err != nil {
		return nil, fmt.Errorf("error expanding Dbx into dbxs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dbxs"] = v
	}

	return m, nil
}

// flattenImageShieldedInstanceInitialState flattens an instance of ImageShieldedInstanceInitialState from a JSON
// response object.
func flattenImageShieldedInstanceInitialState(c *Client, i interface{}) *ImageShieldedInstanceInitialState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageShieldedInstanceInitialState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageShieldedInstanceInitialState
	}
	r.Pk = flattenImageShieldedInstanceInitialStatePk(c, m["pk"])
	r.Kek = flattenImageShieldedInstanceInitialStateKekSlice(c, m["keks"])
	r.Db = flattenImageShieldedInstanceInitialStateDbSlice(c, m["dbs"])
	r.Dbx = flattenImageShieldedInstanceInitialStateDbxSlice(c, m["dbxs"])

	return r
}

// expandImageShieldedInstanceInitialStatePkMap expands the contents of ImageShieldedInstanceInitialStatePk into a JSON
// request object.
func expandImageShieldedInstanceInitialStatePkMap(c *Client, f map[string]ImageShieldedInstanceInitialStatePk) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageShieldedInstanceInitialStatePk(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageShieldedInstanceInitialStatePkSlice expands the contents of ImageShieldedInstanceInitialStatePk into a JSON
// request object.
func expandImageShieldedInstanceInitialStatePkSlice(c *Client, f []ImageShieldedInstanceInitialStatePk) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageShieldedInstanceInitialStatePk(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageShieldedInstanceInitialStatePkMap flattens the contents of ImageShieldedInstanceInitialStatePk from a JSON
// response object.
func flattenImageShieldedInstanceInitialStatePkMap(c *Client, i interface{}) map[string]ImageShieldedInstanceInitialStatePk {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageShieldedInstanceInitialStatePk{}
	}

	if len(a) == 0 {
		return map[string]ImageShieldedInstanceInitialStatePk{}
	}

	items := make(map[string]ImageShieldedInstanceInitialStatePk)
	for k, item := range a {
		items[k] = *flattenImageShieldedInstanceInitialStatePk(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageShieldedInstanceInitialStatePkSlice flattens the contents of ImageShieldedInstanceInitialStatePk from a JSON
// response object.
func flattenImageShieldedInstanceInitialStatePkSlice(c *Client, i interface{}) []ImageShieldedInstanceInitialStatePk {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageShieldedInstanceInitialStatePk{}
	}

	if len(a) == 0 {
		return []ImageShieldedInstanceInitialStatePk{}
	}

	items := make([]ImageShieldedInstanceInitialStatePk, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageShieldedInstanceInitialStatePk(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageShieldedInstanceInitialStatePk expands an instance of ImageShieldedInstanceInitialStatePk into a JSON
// request object.
func expandImageShieldedInstanceInitialStatePk(c *Client, f *ImageShieldedInstanceInitialStatePk) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}
	if v := f.FileType; !dcl.IsEmptyValueIndirect(v) {
		m["fileType"] = v
	}

	return m, nil
}

// flattenImageShieldedInstanceInitialStatePk flattens an instance of ImageShieldedInstanceInitialStatePk from a JSON
// response object.
func flattenImageShieldedInstanceInitialStatePk(c *Client, i interface{}) *ImageShieldedInstanceInitialStatePk {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageShieldedInstanceInitialStatePk{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageShieldedInstanceInitialStatePk
	}
	r.Content = dcl.FlattenString(m["content"])
	r.FileType = flattenImageShieldedInstanceInitialStatePkFileTypeEnum(m["fileType"])

	return r
}

// expandImageShieldedInstanceInitialStateKekMap expands the contents of ImageShieldedInstanceInitialStateKek into a JSON
// request object.
func expandImageShieldedInstanceInitialStateKekMap(c *Client, f map[string]ImageShieldedInstanceInitialStateKek) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageShieldedInstanceInitialStateKek(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageShieldedInstanceInitialStateKekSlice expands the contents of ImageShieldedInstanceInitialStateKek into a JSON
// request object.
func expandImageShieldedInstanceInitialStateKekSlice(c *Client, f []ImageShieldedInstanceInitialStateKek) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageShieldedInstanceInitialStateKek(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageShieldedInstanceInitialStateKekMap flattens the contents of ImageShieldedInstanceInitialStateKek from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateKekMap(c *Client, i interface{}) map[string]ImageShieldedInstanceInitialStateKek {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageShieldedInstanceInitialStateKek{}
	}

	if len(a) == 0 {
		return map[string]ImageShieldedInstanceInitialStateKek{}
	}

	items := make(map[string]ImageShieldedInstanceInitialStateKek)
	for k, item := range a {
		items[k] = *flattenImageShieldedInstanceInitialStateKek(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateKekSlice flattens the contents of ImageShieldedInstanceInitialStateKek from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateKekSlice(c *Client, i interface{}) []ImageShieldedInstanceInitialStateKek {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageShieldedInstanceInitialStateKek{}
	}

	if len(a) == 0 {
		return []ImageShieldedInstanceInitialStateKek{}
	}

	items := make([]ImageShieldedInstanceInitialStateKek, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageShieldedInstanceInitialStateKek(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageShieldedInstanceInitialStateKek expands an instance of ImageShieldedInstanceInitialStateKek into a JSON
// request object.
func expandImageShieldedInstanceInitialStateKek(c *Client, f *ImageShieldedInstanceInitialStateKek) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}
	if v := f.FileType; !dcl.IsEmptyValueIndirect(v) {
		m["fileType"] = v
	}

	return m, nil
}

// flattenImageShieldedInstanceInitialStateKek flattens an instance of ImageShieldedInstanceInitialStateKek from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateKek(c *Client, i interface{}) *ImageShieldedInstanceInitialStateKek {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageShieldedInstanceInitialStateKek{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageShieldedInstanceInitialStateKek
	}
	r.Content = dcl.FlattenString(m["content"])
	r.FileType = flattenImageShieldedInstanceInitialStateKekFileTypeEnum(m["fileType"])

	return r
}

// expandImageShieldedInstanceInitialStateDbMap expands the contents of ImageShieldedInstanceInitialStateDb into a JSON
// request object.
func expandImageShieldedInstanceInitialStateDbMap(c *Client, f map[string]ImageShieldedInstanceInitialStateDb) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageShieldedInstanceInitialStateDb(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageShieldedInstanceInitialStateDbSlice expands the contents of ImageShieldedInstanceInitialStateDb into a JSON
// request object.
func expandImageShieldedInstanceInitialStateDbSlice(c *Client, f []ImageShieldedInstanceInitialStateDb) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageShieldedInstanceInitialStateDb(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageShieldedInstanceInitialStateDbMap flattens the contents of ImageShieldedInstanceInitialStateDb from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDbMap(c *Client, i interface{}) map[string]ImageShieldedInstanceInitialStateDb {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageShieldedInstanceInitialStateDb{}
	}

	if len(a) == 0 {
		return map[string]ImageShieldedInstanceInitialStateDb{}
	}

	items := make(map[string]ImageShieldedInstanceInitialStateDb)
	for k, item := range a {
		items[k] = *flattenImageShieldedInstanceInitialStateDb(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateDbSlice flattens the contents of ImageShieldedInstanceInitialStateDb from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDbSlice(c *Client, i interface{}) []ImageShieldedInstanceInitialStateDb {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageShieldedInstanceInitialStateDb{}
	}

	if len(a) == 0 {
		return []ImageShieldedInstanceInitialStateDb{}
	}

	items := make([]ImageShieldedInstanceInitialStateDb, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageShieldedInstanceInitialStateDb(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageShieldedInstanceInitialStateDb expands an instance of ImageShieldedInstanceInitialStateDb into a JSON
// request object.
func expandImageShieldedInstanceInitialStateDb(c *Client, f *ImageShieldedInstanceInitialStateDb) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}
	if v := f.FileType; !dcl.IsEmptyValueIndirect(v) {
		m["fileType"] = v
	}

	return m, nil
}

// flattenImageShieldedInstanceInitialStateDb flattens an instance of ImageShieldedInstanceInitialStateDb from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDb(c *Client, i interface{}) *ImageShieldedInstanceInitialStateDb {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageShieldedInstanceInitialStateDb{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageShieldedInstanceInitialStateDb
	}
	r.Content = dcl.FlattenString(m["content"])
	r.FileType = flattenImageShieldedInstanceInitialStateDbFileTypeEnum(m["fileType"])

	return r
}

// expandImageShieldedInstanceInitialStateDbxMap expands the contents of ImageShieldedInstanceInitialStateDbx into a JSON
// request object.
func expandImageShieldedInstanceInitialStateDbxMap(c *Client, f map[string]ImageShieldedInstanceInitialStateDbx) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageShieldedInstanceInitialStateDbx(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageShieldedInstanceInitialStateDbxSlice expands the contents of ImageShieldedInstanceInitialStateDbx into a JSON
// request object.
func expandImageShieldedInstanceInitialStateDbxSlice(c *Client, f []ImageShieldedInstanceInitialStateDbx) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageShieldedInstanceInitialStateDbx(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageShieldedInstanceInitialStateDbxMap flattens the contents of ImageShieldedInstanceInitialStateDbx from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDbxMap(c *Client, i interface{}) map[string]ImageShieldedInstanceInitialStateDbx {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageShieldedInstanceInitialStateDbx{}
	}

	if len(a) == 0 {
		return map[string]ImageShieldedInstanceInitialStateDbx{}
	}

	items := make(map[string]ImageShieldedInstanceInitialStateDbx)
	for k, item := range a {
		items[k] = *flattenImageShieldedInstanceInitialStateDbx(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateDbxSlice flattens the contents of ImageShieldedInstanceInitialStateDbx from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDbxSlice(c *Client, i interface{}) []ImageShieldedInstanceInitialStateDbx {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageShieldedInstanceInitialStateDbx{}
	}

	if len(a) == 0 {
		return []ImageShieldedInstanceInitialStateDbx{}
	}

	items := make([]ImageShieldedInstanceInitialStateDbx, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageShieldedInstanceInitialStateDbx(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageShieldedInstanceInitialStateDbx expands an instance of ImageShieldedInstanceInitialStateDbx into a JSON
// request object.
func expandImageShieldedInstanceInitialStateDbx(c *Client, f *ImageShieldedInstanceInitialStateDbx) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}
	if v := f.FileType; !dcl.IsEmptyValueIndirect(v) {
		m["fileType"] = v
	}

	return m, nil
}

// flattenImageShieldedInstanceInitialStateDbx flattens an instance of ImageShieldedInstanceInitialStateDbx from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDbx(c *Client, i interface{}) *ImageShieldedInstanceInitialStateDbx {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageShieldedInstanceInitialStateDbx{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageShieldedInstanceInitialStateDbx
	}
	r.Content = dcl.FlattenString(m["content"])
	r.FileType = flattenImageShieldedInstanceInitialStateDbxFileTypeEnum(m["fileType"])

	return r
}

// expandImageSourceDiskEncryptionKeyMap expands the contents of ImageSourceDiskEncryptionKey into a JSON
// request object.
func expandImageSourceDiskEncryptionKeyMap(c *Client, f map[string]ImageSourceDiskEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageSourceDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageSourceDiskEncryptionKeySlice expands the contents of ImageSourceDiskEncryptionKey into a JSON
// request object.
func expandImageSourceDiskEncryptionKeySlice(c *Client, f []ImageSourceDiskEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageSourceDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageSourceDiskEncryptionKeyMap flattens the contents of ImageSourceDiskEncryptionKey from a JSON
// response object.
func flattenImageSourceDiskEncryptionKeyMap(c *Client, i interface{}) map[string]ImageSourceDiskEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageSourceDiskEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]ImageSourceDiskEncryptionKey{}
	}

	items := make(map[string]ImageSourceDiskEncryptionKey)
	for k, item := range a {
		items[k] = *flattenImageSourceDiskEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageSourceDiskEncryptionKeySlice flattens the contents of ImageSourceDiskEncryptionKey from a JSON
// response object.
func flattenImageSourceDiskEncryptionKeySlice(c *Client, i interface{}) []ImageSourceDiskEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageSourceDiskEncryptionKey{}
	}

	if len(a) == 0 {
		return []ImageSourceDiskEncryptionKey{}
	}

	items := make([]ImageSourceDiskEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageSourceDiskEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageSourceDiskEncryptionKey expands an instance of ImageSourceDiskEncryptionKey into a JSON
// request object.
func expandImageSourceDiskEncryptionKey(c *Client, f *ImageSourceDiskEncryptionKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RawKey; !dcl.IsEmptyValueIndirect(v) {
		m["rawKey"] = v
	}
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}
	if v := f.Sha256; !dcl.IsEmptyValueIndirect(v) {
		m["sha256"] = v
	}
	if v := f.KmsKeyServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyServiceAccount"] = v
	}

	return m, nil
}

// flattenImageSourceDiskEncryptionKey flattens an instance of ImageSourceDiskEncryptionKey from a JSON
// response object.
func flattenImageSourceDiskEncryptionKey(c *Client, i interface{}) *ImageSourceDiskEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageSourceDiskEncryptionKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageSourceDiskEncryptionKey
	}
	r.RawKey = dcl.FlattenString(m["rawKey"])
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	r.Sha256 = dcl.FlattenString(m["sha256"])
	r.KmsKeyServiceAccount = dcl.FlattenString(m["kmsKeyServiceAccount"])

	return r
}

// expandImageSourceImageEncryptionKeyMap expands the contents of ImageSourceImageEncryptionKey into a JSON
// request object.
func expandImageSourceImageEncryptionKeyMap(c *Client, f map[string]ImageSourceImageEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageSourceImageEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageSourceImageEncryptionKeySlice expands the contents of ImageSourceImageEncryptionKey into a JSON
// request object.
func expandImageSourceImageEncryptionKeySlice(c *Client, f []ImageSourceImageEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageSourceImageEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageSourceImageEncryptionKeyMap flattens the contents of ImageSourceImageEncryptionKey from a JSON
// response object.
func flattenImageSourceImageEncryptionKeyMap(c *Client, i interface{}) map[string]ImageSourceImageEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageSourceImageEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]ImageSourceImageEncryptionKey{}
	}

	items := make(map[string]ImageSourceImageEncryptionKey)
	for k, item := range a {
		items[k] = *flattenImageSourceImageEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageSourceImageEncryptionKeySlice flattens the contents of ImageSourceImageEncryptionKey from a JSON
// response object.
func flattenImageSourceImageEncryptionKeySlice(c *Client, i interface{}) []ImageSourceImageEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageSourceImageEncryptionKey{}
	}

	if len(a) == 0 {
		return []ImageSourceImageEncryptionKey{}
	}

	items := make([]ImageSourceImageEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageSourceImageEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageSourceImageEncryptionKey expands an instance of ImageSourceImageEncryptionKey into a JSON
// request object.
func expandImageSourceImageEncryptionKey(c *Client, f *ImageSourceImageEncryptionKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RawKey; !dcl.IsEmptyValueIndirect(v) {
		m["rawKey"] = v
	}
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}
	if v := f.Sha256; !dcl.IsEmptyValueIndirect(v) {
		m["sha256"] = v
	}
	if v := f.KmsKeyServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyServiceAccount"] = v
	}

	return m, nil
}

// flattenImageSourceImageEncryptionKey flattens an instance of ImageSourceImageEncryptionKey from a JSON
// response object.
func flattenImageSourceImageEncryptionKey(c *Client, i interface{}) *ImageSourceImageEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageSourceImageEncryptionKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageSourceImageEncryptionKey
	}
	r.RawKey = dcl.FlattenString(m["rawKey"])
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	r.Sha256 = dcl.FlattenString(m["sha256"])
	r.KmsKeyServiceAccount = dcl.FlattenString(m["kmsKeyServiceAccount"])

	return r
}

// expandImageSourceSnapshotEncryptionKeyMap expands the contents of ImageSourceSnapshotEncryptionKey into a JSON
// request object.
func expandImageSourceSnapshotEncryptionKeyMap(c *Client, f map[string]ImageSourceSnapshotEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageSourceSnapshotEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageSourceSnapshotEncryptionKeySlice expands the contents of ImageSourceSnapshotEncryptionKey into a JSON
// request object.
func expandImageSourceSnapshotEncryptionKeySlice(c *Client, f []ImageSourceSnapshotEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageSourceSnapshotEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageSourceSnapshotEncryptionKeyMap flattens the contents of ImageSourceSnapshotEncryptionKey from a JSON
// response object.
func flattenImageSourceSnapshotEncryptionKeyMap(c *Client, i interface{}) map[string]ImageSourceSnapshotEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageSourceSnapshotEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]ImageSourceSnapshotEncryptionKey{}
	}

	items := make(map[string]ImageSourceSnapshotEncryptionKey)
	for k, item := range a {
		items[k] = *flattenImageSourceSnapshotEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageSourceSnapshotEncryptionKeySlice flattens the contents of ImageSourceSnapshotEncryptionKey from a JSON
// response object.
func flattenImageSourceSnapshotEncryptionKeySlice(c *Client, i interface{}) []ImageSourceSnapshotEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageSourceSnapshotEncryptionKey{}
	}

	if len(a) == 0 {
		return []ImageSourceSnapshotEncryptionKey{}
	}

	items := make([]ImageSourceSnapshotEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageSourceSnapshotEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageSourceSnapshotEncryptionKey expands an instance of ImageSourceSnapshotEncryptionKey into a JSON
// request object.
func expandImageSourceSnapshotEncryptionKey(c *Client, f *ImageSourceSnapshotEncryptionKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RawKey; !dcl.IsEmptyValueIndirect(v) {
		m["rawKey"] = v
	}
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}
	if v := f.Sha256; !dcl.IsEmptyValueIndirect(v) {
		m["sha256"] = v
	}
	if v := f.KmsKeyServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyServiceAccount"] = v
	}

	return m, nil
}

// flattenImageSourceSnapshotEncryptionKey flattens an instance of ImageSourceSnapshotEncryptionKey from a JSON
// response object.
func flattenImageSourceSnapshotEncryptionKey(c *Client, i interface{}) *ImageSourceSnapshotEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageSourceSnapshotEncryptionKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageSourceSnapshotEncryptionKey
	}
	r.RawKey = dcl.FlattenString(m["rawKey"])
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	r.Sha256 = dcl.FlattenString(m["sha256"])
	r.KmsKeyServiceAccount = dcl.FlattenString(m["kmsKeyServiceAccount"])

	return r
}

// expandImageDeprecatedMap expands the contents of ImageDeprecated into a JSON
// request object.
func expandImageDeprecatedMap(c *Client, f map[string]ImageDeprecated) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageDeprecated(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageDeprecatedSlice expands the contents of ImageDeprecated into a JSON
// request object.
func expandImageDeprecatedSlice(c *Client, f []ImageDeprecated) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageDeprecated(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageDeprecatedMap flattens the contents of ImageDeprecated from a JSON
// response object.
func flattenImageDeprecatedMap(c *Client, i interface{}) map[string]ImageDeprecated {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageDeprecated{}
	}

	if len(a) == 0 {
		return map[string]ImageDeprecated{}
	}

	items := make(map[string]ImageDeprecated)
	for k, item := range a {
		items[k] = *flattenImageDeprecated(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageDeprecatedSlice flattens the contents of ImageDeprecated from a JSON
// response object.
func flattenImageDeprecatedSlice(c *Client, i interface{}) []ImageDeprecated {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageDeprecated{}
	}

	if len(a) == 0 {
		return []ImageDeprecated{}
	}

	items := make([]ImageDeprecated, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageDeprecated(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageDeprecated expands an instance of ImageDeprecated into a JSON
// request object.
func expandImageDeprecated(c *Client, f *ImageDeprecated) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Replacement; !dcl.IsEmptyValueIndirect(v) {
		m["replacement"] = v
	}
	if v := f.Deprecated; !dcl.IsEmptyValueIndirect(v) {
		m["deprecated"] = v
	}
	if v := f.Obsolete; !dcl.IsEmptyValueIndirect(v) {
		m["obsolete"] = v
	}
	if v := f.Deleted; !dcl.IsEmptyValueIndirect(v) {
		m["deleted"] = v
	}

	return m, nil
}

// flattenImageDeprecated flattens an instance of ImageDeprecated from a JSON
// response object.
func flattenImageDeprecated(c *Client, i interface{}) *ImageDeprecated {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageDeprecated{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyImageDeprecated
	}
	r.State = flattenImageDeprecatedStateEnum(m["state"])
	r.Replacement = dcl.FlattenString(m["replacement"])
	r.Deprecated = dcl.FlattenString(m["deprecated"])
	r.Obsolete = dcl.FlattenString(m["obsolete"])
	r.Deleted = dcl.FlattenString(m["deleted"])

	return r
}

// flattenImageGuestOSFeatureTypeEnumMap flattens the contents of ImageGuestOSFeatureTypeEnum from a JSON
// response object.
func flattenImageGuestOSFeatureTypeEnumMap(c *Client, i interface{}) map[string]ImageGuestOSFeatureTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageGuestOSFeatureTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ImageGuestOSFeatureTypeEnum{}
	}

	items := make(map[string]ImageGuestOSFeatureTypeEnum)
	for k, item := range a {
		items[k] = *flattenImageGuestOSFeatureTypeEnum(item.(interface{}))
	}

	return items
}

// flattenImageGuestOSFeatureTypeEnumSlice flattens the contents of ImageGuestOSFeatureTypeEnum from a JSON
// response object.
func flattenImageGuestOSFeatureTypeEnumSlice(c *Client, i interface{}) []ImageGuestOSFeatureTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageGuestOSFeatureTypeEnum{}
	}

	if len(a) == 0 {
		return []ImageGuestOSFeatureTypeEnum{}
	}

	items := make([]ImageGuestOSFeatureTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageGuestOSFeatureTypeEnum(item.(interface{})))
	}

	return items
}

// flattenImageGuestOSFeatureTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ImageGuestOSFeatureTypeEnum with the same value as that string.
func flattenImageGuestOSFeatureTypeEnum(i interface{}) *ImageGuestOSFeatureTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ImageGuestOSFeatureTypeEnumRef("")
	}

	return ImageGuestOSFeatureTypeEnumRef(s)
}

// flattenImageRawDiskContainerTypeEnumMap flattens the contents of ImageRawDiskContainerTypeEnum from a JSON
// response object.
func flattenImageRawDiskContainerTypeEnumMap(c *Client, i interface{}) map[string]ImageRawDiskContainerTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageRawDiskContainerTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ImageRawDiskContainerTypeEnum{}
	}

	items := make(map[string]ImageRawDiskContainerTypeEnum)
	for k, item := range a {
		items[k] = *flattenImageRawDiskContainerTypeEnum(item.(interface{}))
	}

	return items
}

// flattenImageRawDiskContainerTypeEnumSlice flattens the contents of ImageRawDiskContainerTypeEnum from a JSON
// response object.
func flattenImageRawDiskContainerTypeEnumSlice(c *Client, i interface{}) []ImageRawDiskContainerTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageRawDiskContainerTypeEnum{}
	}

	if len(a) == 0 {
		return []ImageRawDiskContainerTypeEnum{}
	}

	items := make([]ImageRawDiskContainerTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageRawDiskContainerTypeEnum(item.(interface{})))
	}

	return items
}

// flattenImageRawDiskContainerTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ImageRawDiskContainerTypeEnum with the same value as that string.
func flattenImageRawDiskContainerTypeEnum(i interface{}) *ImageRawDiskContainerTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ImageRawDiskContainerTypeEnumRef("")
	}

	return ImageRawDiskContainerTypeEnumRef(s)
}

// flattenImageShieldedInstanceInitialStatePkFileTypeEnumMap flattens the contents of ImageShieldedInstanceInitialStatePkFileTypeEnum from a JSON
// response object.
func flattenImageShieldedInstanceInitialStatePkFileTypeEnumMap(c *Client, i interface{}) map[string]ImageShieldedInstanceInitialStatePkFileTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageShieldedInstanceInitialStatePkFileTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ImageShieldedInstanceInitialStatePkFileTypeEnum{}
	}

	items := make(map[string]ImageShieldedInstanceInitialStatePkFileTypeEnum)
	for k, item := range a {
		items[k] = *flattenImageShieldedInstanceInitialStatePkFileTypeEnum(item.(interface{}))
	}

	return items
}

// flattenImageShieldedInstanceInitialStatePkFileTypeEnumSlice flattens the contents of ImageShieldedInstanceInitialStatePkFileTypeEnum from a JSON
// response object.
func flattenImageShieldedInstanceInitialStatePkFileTypeEnumSlice(c *Client, i interface{}) []ImageShieldedInstanceInitialStatePkFileTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageShieldedInstanceInitialStatePkFileTypeEnum{}
	}

	if len(a) == 0 {
		return []ImageShieldedInstanceInitialStatePkFileTypeEnum{}
	}

	items := make([]ImageShieldedInstanceInitialStatePkFileTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageShieldedInstanceInitialStatePkFileTypeEnum(item.(interface{})))
	}

	return items
}

// flattenImageShieldedInstanceInitialStatePkFileTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ImageShieldedInstanceInitialStatePkFileTypeEnum with the same value as that string.
func flattenImageShieldedInstanceInitialStatePkFileTypeEnum(i interface{}) *ImageShieldedInstanceInitialStatePkFileTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ImageShieldedInstanceInitialStatePkFileTypeEnumRef("")
	}

	return ImageShieldedInstanceInitialStatePkFileTypeEnumRef(s)
}

// flattenImageShieldedInstanceInitialStateKekFileTypeEnumMap flattens the contents of ImageShieldedInstanceInitialStateKekFileTypeEnum from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateKekFileTypeEnumMap(c *Client, i interface{}) map[string]ImageShieldedInstanceInitialStateKekFileTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageShieldedInstanceInitialStateKekFileTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ImageShieldedInstanceInitialStateKekFileTypeEnum{}
	}

	items := make(map[string]ImageShieldedInstanceInitialStateKekFileTypeEnum)
	for k, item := range a {
		items[k] = *flattenImageShieldedInstanceInitialStateKekFileTypeEnum(item.(interface{}))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateKekFileTypeEnumSlice flattens the contents of ImageShieldedInstanceInitialStateKekFileTypeEnum from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateKekFileTypeEnumSlice(c *Client, i interface{}) []ImageShieldedInstanceInitialStateKekFileTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageShieldedInstanceInitialStateKekFileTypeEnum{}
	}

	if len(a) == 0 {
		return []ImageShieldedInstanceInitialStateKekFileTypeEnum{}
	}

	items := make([]ImageShieldedInstanceInitialStateKekFileTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageShieldedInstanceInitialStateKekFileTypeEnum(item.(interface{})))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateKekFileTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ImageShieldedInstanceInitialStateKekFileTypeEnum with the same value as that string.
func flattenImageShieldedInstanceInitialStateKekFileTypeEnum(i interface{}) *ImageShieldedInstanceInitialStateKekFileTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ImageShieldedInstanceInitialStateKekFileTypeEnumRef("")
	}

	return ImageShieldedInstanceInitialStateKekFileTypeEnumRef(s)
}

// flattenImageShieldedInstanceInitialStateDbFileTypeEnumMap flattens the contents of ImageShieldedInstanceInitialStateDbFileTypeEnum from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDbFileTypeEnumMap(c *Client, i interface{}) map[string]ImageShieldedInstanceInitialStateDbFileTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageShieldedInstanceInitialStateDbFileTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ImageShieldedInstanceInitialStateDbFileTypeEnum{}
	}

	items := make(map[string]ImageShieldedInstanceInitialStateDbFileTypeEnum)
	for k, item := range a {
		items[k] = *flattenImageShieldedInstanceInitialStateDbFileTypeEnum(item.(interface{}))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateDbFileTypeEnumSlice flattens the contents of ImageShieldedInstanceInitialStateDbFileTypeEnum from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDbFileTypeEnumSlice(c *Client, i interface{}) []ImageShieldedInstanceInitialStateDbFileTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageShieldedInstanceInitialStateDbFileTypeEnum{}
	}

	if len(a) == 0 {
		return []ImageShieldedInstanceInitialStateDbFileTypeEnum{}
	}

	items := make([]ImageShieldedInstanceInitialStateDbFileTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageShieldedInstanceInitialStateDbFileTypeEnum(item.(interface{})))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateDbFileTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ImageShieldedInstanceInitialStateDbFileTypeEnum with the same value as that string.
func flattenImageShieldedInstanceInitialStateDbFileTypeEnum(i interface{}) *ImageShieldedInstanceInitialStateDbFileTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ImageShieldedInstanceInitialStateDbFileTypeEnumRef("")
	}

	return ImageShieldedInstanceInitialStateDbFileTypeEnumRef(s)
}

// flattenImageShieldedInstanceInitialStateDbxFileTypeEnumMap flattens the contents of ImageShieldedInstanceInitialStateDbxFileTypeEnum from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDbxFileTypeEnumMap(c *Client, i interface{}) map[string]ImageShieldedInstanceInitialStateDbxFileTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageShieldedInstanceInitialStateDbxFileTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ImageShieldedInstanceInitialStateDbxFileTypeEnum{}
	}

	items := make(map[string]ImageShieldedInstanceInitialStateDbxFileTypeEnum)
	for k, item := range a {
		items[k] = *flattenImageShieldedInstanceInitialStateDbxFileTypeEnum(item.(interface{}))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateDbxFileTypeEnumSlice flattens the contents of ImageShieldedInstanceInitialStateDbxFileTypeEnum from a JSON
// response object.
func flattenImageShieldedInstanceInitialStateDbxFileTypeEnumSlice(c *Client, i interface{}) []ImageShieldedInstanceInitialStateDbxFileTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageShieldedInstanceInitialStateDbxFileTypeEnum{}
	}

	if len(a) == 0 {
		return []ImageShieldedInstanceInitialStateDbxFileTypeEnum{}
	}

	items := make([]ImageShieldedInstanceInitialStateDbxFileTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageShieldedInstanceInitialStateDbxFileTypeEnum(item.(interface{})))
	}

	return items
}

// flattenImageShieldedInstanceInitialStateDbxFileTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ImageShieldedInstanceInitialStateDbxFileTypeEnum with the same value as that string.
func flattenImageShieldedInstanceInitialStateDbxFileTypeEnum(i interface{}) *ImageShieldedInstanceInitialStateDbxFileTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ImageShieldedInstanceInitialStateDbxFileTypeEnumRef("")
	}

	return ImageShieldedInstanceInitialStateDbxFileTypeEnumRef(s)
}

// flattenImageSourceTypeEnumMap flattens the contents of ImageSourceTypeEnum from a JSON
// response object.
func flattenImageSourceTypeEnumMap(c *Client, i interface{}) map[string]ImageSourceTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageSourceTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ImageSourceTypeEnum{}
	}

	items := make(map[string]ImageSourceTypeEnum)
	for k, item := range a {
		items[k] = *flattenImageSourceTypeEnum(item.(interface{}))
	}

	return items
}

// flattenImageSourceTypeEnumSlice flattens the contents of ImageSourceTypeEnum from a JSON
// response object.
func flattenImageSourceTypeEnumSlice(c *Client, i interface{}) []ImageSourceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageSourceTypeEnum{}
	}

	if len(a) == 0 {
		return []ImageSourceTypeEnum{}
	}

	items := make([]ImageSourceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageSourceTypeEnum(item.(interface{})))
	}

	return items
}

// flattenImageSourceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ImageSourceTypeEnum with the same value as that string.
func flattenImageSourceTypeEnum(i interface{}) *ImageSourceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ImageSourceTypeEnumRef("")
	}

	return ImageSourceTypeEnumRef(s)
}

// flattenImageStatusEnumMap flattens the contents of ImageStatusEnum from a JSON
// response object.
func flattenImageStatusEnumMap(c *Client, i interface{}) map[string]ImageStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]ImageStatusEnum{}
	}

	items := make(map[string]ImageStatusEnum)
	for k, item := range a {
		items[k] = *flattenImageStatusEnum(item.(interface{}))
	}

	return items
}

// flattenImageStatusEnumSlice flattens the contents of ImageStatusEnum from a JSON
// response object.
func flattenImageStatusEnumSlice(c *Client, i interface{}) []ImageStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageStatusEnum{}
	}

	if len(a) == 0 {
		return []ImageStatusEnum{}
	}

	items := make([]ImageStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageStatusEnum(item.(interface{})))
	}

	return items
}

// flattenImageStatusEnum asserts that an interface is a string, and returns a
// pointer to a *ImageStatusEnum with the same value as that string.
func flattenImageStatusEnum(i interface{}) *ImageStatusEnum {
	s, ok := i.(string)
	if !ok {
		return ImageStatusEnumRef("")
	}

	return ImageStatusEnumRef(s)
}

// flattenImageDeprecatedStateEnumMap flattens the contents of ImageDeprecatedStateEnum from a JSON
// response object.
func flattenImageDeprecatedStateEnumMap(c *Client, i interface{}) map[string]ImageDeprecatedStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageDeprecatedStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ImageDeprecatedStateEnum{}
	}

	items := make(map[string]ImageDeprecatedStateEnum)
	for k, item := range a {
		items[k] = *flattenImageDeprecatedStateEnum(item.(interface{}))
	}

	return items
}

// flattenImageDeprecatedStateEnumSlice flattens the contents of ImageDeprecatedStateEnum from a JSON
// response object.
func flattenImageDeprecatedStateEnumSlice(c *Client, i interface{}) []ImageDeprecatedStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageDeprecatedStateEnum{}
	}

	if len(a) == 0 {
		return []ImageDeprecatedStateEnum{}
	}

	items := make([]ImageDeprecatedStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageDeprecatedStateEnum(item.(interface{})))
	}

	return items
}

// flattenImageDeprecatedStateEnum asserts that an interface is a string, and returns a
// pointer to a *ImageDeprecatedStateEnum with the same value as that string.
func flattenImageDeprecatedStateEnum(i interface{}) *ImageDeprecatedStateEnum {
	s, ok := i.(string)
	if !ok {
		return ImageDeprecatedStateEnumRef("")
	}

	return ImageDeprecatedStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Image) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalImage(b, c)
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

type imageDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         imageApiOperation
}

func convertFieldDiffsToImageDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]imageDiff, error) {
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
	var diffs []imageDiff
	// For each operation name, create a imageDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := imageDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToImageApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToImageApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (imageApiOperation, error) {
	switch opName {

	case "updateImageDeprecateOperation":
		return &updateImageDeprecateOperation{FieldDiffs: fieldDiffs}, nil

	case "updateImageSetLabelsOperation":
		return &updateImageSetLabelsOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
