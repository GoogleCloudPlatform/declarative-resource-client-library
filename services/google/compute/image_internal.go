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
package compute

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mohae/deepcopy"
	"io/ioutil"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
	"reflect"
	"strings"
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
func (r *ImageGuestOsFeature) validate() error {
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
	return dcl.URL("projects/{{project}}/global/images/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func imageListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/images", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func imageCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/images", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func imageDeleteURL(userBasePath string, r *Image) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/images/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
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
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateImageDeprecateOperation) do(ctx context.Context, r *Image, c *Client) error {
	_, err := c.GetImage(ctx, r.urlNormalized())
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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

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
	b, err := c.getImageRaw(ctx, f.urlNormalized())
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
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateImageSetLabelsOperation) do(ctx context.Context, r *Image, c *Client) error {
	_, err := c.GetImage(ctx, r.urlNormalized())
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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "LIST", u, &bytes.Buffer{}, c.Config.Retry)
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
		res := flattenImage(c, v)
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

	_, err := c.GetImage(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Image not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetImage checking for existence. error: %v", err)
		return err
	}

	u, err := imageDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		return err
	}
	_, err = c.GetImage(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createImageOperation struct{}

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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetImage(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getImageRaw(ctx context.Context, r *Image) ([]byte, error) {

	u, err := imageGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
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

func (c *Client) imageDiffsForRawDesired(ctx context.Context, rawDesired *Image, opts ...dcl.ApplyOption) (initial, desired *Image, diffs []imageDiff, err error) {
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
	rawInitial, err := c.GetImage(ctx, fetchState.urlNormalized())
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Image); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected Image, got %T", sh)
		} else {
			_ = r
		}
	}

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
	if dcl.IsZeroValue(rawDesired.ArchiveSizeBytes) {
		rawDesired.ArchiveSizeBytes = rawInitial.ArchiveSizeBytes
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.DiskSizeGb) {
		rawDesired.DiskSizeGb = rawInitial.DiskSizeGb
	}
	if dcl.IsZeroValue(rawDesired.Family) {
		rawDesired.Family = rawInitial.Family
	}
	if dcl.IsZeroValue(rawDesired.GuestOsFeature) {
		rawDesired.GuestOsFeature = rawInitial.GuestOsFeature
	}
	rawDesired.ImageEncryptionKey = canonicalizeImageImageEncryptionKey(rawDesired.ImageEncryptionKey, rawInitial.ImageEncryptionKey, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.License) {
		rawDesired.License = rawInitial.License
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.RawDisk = canonicalizeImageRawDisk(rawDesired.RawDisk, rawInitial.RawDisk, opts...)
	rawDesired.ShieldedInstanceInitialState = canonicalizeImageShieldedInstanceInitialState(rawDesired.ShieldedInstanceInitialState, rawInitial.ShieldedInstanceInitialState, opts...)
	if dcl.IsZeroValue(rawDesired.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.IsZeroValue(rawDesired.SourceDisk) {
		rawDesired.SourceDisk = rawInitial.SourceDisk
	}
	rawDesired.SourceDiskEncryptionKey = canonicalizeImageSourceDiskEncryptionKey(rawDesired.SourceDiskEncryptionKey, rawInitial.SourceDiskEncryptionKey, opts...)
	if dcl.IsZeroValue(rawDesired.SourceDiskId) {
		rawDesired.SourceDiskId = rawInitial.SourceDiskId
	}
	if dcl.IsZeroValue(rawDesired.SourceImage) {
		rawDesired.SourceImage = rawInitial.SourceImage
	}
	rawDesired.SourceImageEncryptionKey = canonicalizeImageSourceImageEncryptionKey(rawDesired.SourceImageEncryptionKey, rawInitial.SourceImageEncryptionKey, opts...)
	if dcl.IsZeroValue(rawDesired.SourceImageId) {
		rawDesired.SourceImageId = rawInitial.SourceImageId
	}
	if dcl.IsZeroValue(rawDesired.SourceSnapshot) {
		rawDesired.SourceSnapshot = rawInitial.SourceSnapshot
	}
	rawDesired.SourceSnapshotEncryptionKey = canonicalizeImageSourceSnapshotEncryptionKey(rawDesired.SourceSnapshotEncryptionKey, rawInitial.SourceSnapshotEncryptionKey, opts...)
	if dcl.IsZeroValue(rawDesired.SourceSnapshotId) {
		rawDesired.SourceSnapshotId = rawInitial.SourceSnapshotId
	}
	if dcl.IsZeroValue(rawDesired.SourceType) {
		rawDesired.SourceType = rawInitial.SourceType
	}
	if dcl.IsZeroValue(rawDesired.Status) {
		rawDesired.Status = rawInitial.Status
	}
	if dcl.IsZeroValue(rawDesired.StorageLocation) {
		rawDesired.StorageLocation = rawInitial.StorageLocation
	}
	rawDesired.Deprecated = canonicalizeImageDeprecated(rawDesired.Deprecated, rawInitial.Deprecated, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeImageNewState(c *Client, rawNew, rawDesired *Image) (*Image, error) {

	if dcl.IsEmptyValueIndirect(rawNew.ArchiveSizeBytes) && dcl.IsEmptyValueIndirect(rawDesired.ArchiveSizeBytes) {
		rawNew.ArchiveSizeBytes = rawDesired.ArchiveSizeBytes
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskSizeGb) && dcl.IsEmptyValueIndirect(rawDesired.DiskSizeGb) {
		rawNew.DiskSizeGb = rawDesired.DiskSizeGb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Family) && dcl.IsEmptyValueIndirect(rawDesired.Family) {
		rawNew.Family = rawDesired.Family
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.GuestOsFeature) && dcl.IsEmptyValueIndirect(rawDesired.GuestOsFeature) {
		rawNew.GuestOsFeature = rawDesired.GuestOsFeature
	} else {
		rawNew.GuestOsFeature = canonicalizeNewImageGuestOsFeatureSet(c, rawDesired.GuestOsFeature, rawNew.GuestOsFeature)
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceDisk) && dcl.IsEmptyValueIndirect(rawDesired.SourceDisk) {
		rawNew.SourceDisk = rawDesired.SourceDisk
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceDiskEncryptionKey) && dcl.IsEmptyValueIndirect(rawDesired.SourceDiskEncryptionKey) {
		rawNew.SourceDiskEncryptionKey = rawDesired.SourceDiskEncryptionKey
	} else {
		rawNew.SourceDiskEncryptionKey = canonicalizeNewImageSourceDiskEncryptionKey(c, rawDesired.SourceDiskEncryptionKey, rawNew.SourceDiskEncryptionKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceDiskId) && dcl.IsEmptyValueIndirect(rawDesired.SourceDiskId) {
		rawNew.SourceDiskId = rawDesired.SourceDiskId
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceImage) && dcl.IsEmptyValueIndirect(rawDesired.SourceImage) {
		rawNew.SourceImage = rawDesired.SourceImage
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceImageEncryptionKey) && dcl.IsEmptyValueIndirect(rawDesired.SourceImageEncryptionKey) {
		rawNew.SourceImageEncryptionKey = rawDesired.SourceImageEncryptionKey
	} else {
		rawNew.SourceImageEncryptionKey = canonicalizeNewImageSourceImageEncryptionKey(c, rawDesired.SourceImageEncryptionKey, rawNew.SourceImageEncryptionKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceImageId) && dcl.IsEmptyValueIndirect(rawDesired.SourceImageId) {
		rawNew.SourceImageId = rawDesired.SourceImageId
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceSnapshot) && dcl.IsEmptyValueIndirect(rawDesired.SourceSnapshot) {
		rawNew.SourceSnapshot = rawDesired.SourceSnapshot
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceSnapshotEncryptionKey) && dcl.IsEmptyValueIndirect(rawDesired.SourceSnapshotEncryptionKey) {
		rawNew.SourceSnapshotEncryptionKey = rawDesired.SourceSnapshotEncryptionKey
	} else {
		rawNew.SourceSnapshotEncryptionKey = canonicalizeNewImageSourceSnapshotEncryptionKey(c, rawDesired.SourceSnapshotEncryptionKey, rawNew.SourceSnapshotEncryptionKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceSnapshotId) && dcl.IsEmptyValueIndirect(rawDesired.SourceSnapshotId) {
		rawNew.SourceSnapshotId = rawDesired.SourceSnapshotId
	} else {
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

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.NameToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	return rawNew, nil
}

func canonicalizeImageGuestOsFeature(des, initial *ImageGuestOsFeature, opts ...dcl.ApplyOption) *ImageGuestOsFeature {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewImageGuestOsFeature(c *Client, des, nw *ImageGuestOsFeature) *ImageGuestOsFeature {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewImageGuestOsFeatureSet(c *Client, des, nw []ImageGuestOsFeature) []ImageGuestOsFeature {
	if des == nil {
		return nw
	}
	var reorderedNew []ImageGuestOsFeature
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareImageGuestOsFeature(c, &d, &n) {
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

func canonicalizeImageImageEncryptionKey(des, initial *ImageImageEncryptionKey, opts ...dcl.ApplyOption) *ImageImageEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}
	if dcl.IsZeroValue(des.Sha256) {
		des.Sha256 = initial.Sha256
	}
	if dcl.IsZeroValue(des.KmsKeyServiceAccount) {
		des.KmsKeyServiceAccount = initial.KmsKeyServiceAccount
	}

	return des
}

func canonicalizeNewImageImageEncryptionKey(c *Client, des, nw *ImageImageEncryptionKey) *ImageImageEncryptionKey {
	if des == nil || nw == nil {
		return nw
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
			if !compareImageImageEncryptionKey(c, &d, &n) {
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Source) {
		des.Source = initial.Source
	}
	if dcl.IsZeroValue(des.Sha1Checksum) {
		des.Sha1Checksum = initial.Sha1Checksum
	}
	if dcl.IsZeroValue(des.ContainerType) {
		des.ContainerType = initial.ContainerType
	}

	return des
}

func canonicalizeNewImageRawDisk(c *Client, des, nw *ImageRawDisk) *ImageRawDisk {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ContainerType) {
		nw.ContainerType = ImageRawDiskContainerTypeEnumRef("TAR")
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
			if !compareImageRawDisk(c, &d, &n) {
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

func canonicalizeImageShieldedInstanceInitialState(des, initial *ImageShieldedInstanceInitialState, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.Pk = canonicalizeImageShieldedInstanceInitialStatePk(des.Pk, initial.Pk, opts...)
	if dcl.IsZeroValue(des.Kek) {
		des.Kek = initial.Kek
	}
	if dcl.IsZeroValue(des.Db) {
		des.Db = initial.Db
	}
	if dcl.IsZeroValue(des.Dbx) {
		des.Dbx = initial.Dbx
	}

	return des
}

func canonicalizeNewImageShieldedInstanceInitialState(c *Client, des, nw *ImageShieldedInstanceInitialState) *ImageShieldedInstanceInitialState {
	if des == nil || nw == nil {
		return nw
	}

	nw.Pk = canonicalizeNewImageShieldedInstanceInitialStatePk(c, des.Pk, nw.Pk)

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
			if !compareImageShieldedInstanceInitialState(c, &d, &n) {
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

func canonicalizeImageShieldedInstanceInitialStatePk(des, initial *ImageShieldedInstanceInitialStatePk, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialStatePk {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Content) {
		des.Content = initial.Content
	}
	if dcl.IsZeroValue(des.FileType) {
		des.FileType = initial.FileType
	}

	return des
}

func canonicalizeNewImageShieldedInstanceInitialStatePk(c *Client, des, nw *ImageShieldedInstanceInitialStatePk) *ImageShieldedInstanceInitialStatePk {
	if des == nil || nw == nil {
		return nw
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
			if !compareImageShieldedInstanceInitialStatePk(c, &d, &n) {
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

func canonicalizeImageShieldedInstanceInitialStateKek(des, initial *ImageShieldedInstanceInitialStateKek, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialStateKek {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Content) {
		des.Content = initial.Content
	}
	if dcl.IsZeroValue(des.FileType) {
		des.FileType = initial.FileType
	}

	return des
}

func canonicalizeNewImageShieldedInstanceInitialStateKek(c *Client, des, nw *ImageShieldedInstanceInitialStateKek) *ImageShieldedInstanceInitialStateKek {
	if des == nil || nw == nil {
		return nw
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
			if !compareImageShieldedInstanceInitialStateKek(c, &d, &n) {
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

func canonicalizeImageShieldedInstanceInitialStateDb(des, initial *ImageShieldedInstanceInitialStateDb, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialStateDb {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Content) {
		des.Content = initial.Content
	}
	if dcl.IsZeroValue(des.FileType) {
		des.FileType = initial.FileType
	}

	return des
}

func canonicalizeNewImageShieldedInstanceInitialStateDb(c *Client, des, nw *ImageShieldedInstanceInitialStateDb) *ImageShieldedInstanceInitialStateDb {
	if des == nil || nw == nil {
		return nw
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
			if !compareImageShieldedInstanceInitialStateDb(c, &d, &n) {
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

func canonicalizeImageShieldedInstanceInitialStateDbx(des, initial *ImageShieldedInstanceInitialStateDbx, opts ...dcl.ApplyOption) *ImageShieldedInstanceInitialStateDbx {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Content) {
		des.Content = initial.Content
	}
	if dcl.IsZeroValue(des.FileType) {
		des.FileType = initial.FileType
	}

	return des
}

func canonicalizeNewImageShieldedInstanceInitialStateDbx(c *Client, des, nw *ImageShieldedInstanceInitialStateDbx) *ImageShieldedInstanceInitialStateDbx {
	if des == nil || nw == nil {
		return nw
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
			if !compareImageShieldedInstanceInitialStateDbx(c, &d, &n) {
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

func canonicalizeImageSourceDiskEncryptionKey(des, initial *ImageSourceDiskEncryptionKey, opts ...dcl.ApplyOption) *ImageSourceDiskEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}
	if dcl.IsZeroValue(des.Sha256) {
		des.Sha256 = initial.Sha256
	}
	if dcl.IsZeroValue(des.KmsKeyServiceAccount) {
		des.KmsKeyServiceAccount = initial.KmsKeyServiceAccount
	}

	return des
}

func canonicalizeNewImageSourceDiskEncryptionKey(c *Client, des, nw *ImageSourceDiskEncryptionKey) *ImageSourceDiskEncryptionKey {
	if des == nil || nw == nil {
		return nw
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
			if !compareImageSourceDiskEncryptionKey(c, &d, &n) {
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

func canonicalizeImageSourceImageEncryptionKey(des, initial *ImageSourceImageEncryptionKey, opts ...dcl.ApplyOption) *ImageSourceImageEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}
	if dcl.IsZeroValue(des.Sha256) {
		des.Sha256 = initial.Sha256
	}
	if dcl.IsZeroValue(des.KmsKeyServiceAccount) {
		des.KmsKeyServiceAccount = initial.KmsKeyServiceAccount
	}

	return des
}

func canonicalizeNewImageSourceImageEncryptionKey(c *Client, des, nw *ImageSourceImageEncryptionKey) *ImageSourceImageEncryptionKey {
	if des == nil || nw == nil {
		return nw
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
			if !compareImageSourceImageEncryptionKey(c, &d, &n) {
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

func canonicalizeImageSourceSnapshotEncryptionKey(des, initial *ImageSourceSnapshotEncryptionKey, opts ...dcl.ApplyOption) *ImageSourceSnapshotEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}
	if dcl.IsZeroValue(des.Sha256) {
		des.Sha256 = initial.Sha256
	}
	if dcl.IsZeroValue(des.KmsKeyServiceAccount) {
		des.KmsKeyServiceAccount = initial.KmsKeyServiceAccount
	}

	return des
}

func canonicalizeNewImageSourceSnapshotEncryptionKey(c *Client, des, nw *ImageSourceSnapshotEncryptionKey) *ImageSourceSnapshotEncryptionKey {
	if des == nil || nw == nil {
		return nw
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
			if !compareImageSourceSnapshotEncryptionKey(c, &d, &n) {
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

func canonicalizeImageDeprecated(des, initial *ImageDeprecated, opts ...dcl.ApplyOption) *ImageDeprecated {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Image)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.IsZeroValue(des.Replacement) {
		des.Replacement = initial.Replacement
	}
	if dcl.IsZeroValue(des.Deprecated) {
		des.Deprecated = initial.Deprecated
	}
	if dcl.IsZeroValue(des.Obsolete) {
		des.Obsolete = initial.Obsolete
	}
	if dcl.IsZeroValue(des.Deleted) {
		des.Deleted = initial.Deleted
	}

	return des
}

func canonicalizeNewImageDeprecated(c *Client, des, nw *ImageDeprecated) *ImageDeprecated {
	if des == nil || nw == nil {
		return nw
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
			if !compareImageDeprecated(c, &d, &n) {
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

type imageDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         imageApiOperation
	// This is for reporting only.
	FieldName string
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffImage(c *Client, desired, actual *Image, opts ...dcl.ApplyOption) ([]imageDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []imageDiff
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %#v\nACTUAL: %#v", desired.Description, actual.Description)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.DiskSizeGb) && (dcl.IsZeroValue(actual.DiskSizeGb) || !reflect.DeepEqual(*desired.DiskSizeGb, *actual.DiskSizeGb)) {
		c.Config.Logger.Infof("Detected diff in DiskSizeGb.\nDESIRED: %#v\nACTUAL: %#v", desired.DiskSizeGb, actual.DiskSizeGb)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "DiskSizeGb",
		})
	}
	if !dcl.IsZeroValue(desired.Family) && (dcl.IsZeroValue(actual.Family) || !reflect.DeepEqual(*desired.Family, *actual.Family)) {
		c.Config.Logger.Infof("Detected diff in Family.\nDESIRED: %#v\nACTUAL: %#v", desired.Family, actual.Family)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "Family",
		})
	}
	if compareImageGuestOsFeatureSlice(c, desired.GuestOsFeature, actual.GuestOsFeature) {
		c.Config.Logger.Infof("Detected diff in GuestOsFeature.\nDESIRED: %#v\nACTUAL: %#v", desired.GuestOsFeature, actual.GuestOsFeature)
		toAdd, toRemove := compareImageGuestOsFeatureSets(c, desired.GuestOsFeature, actual.GuestOsFeature)
		if len(toAdd) > 0 {
			diffs = append(diffs, imageDiff{
				RequiresRecreate: true,
				FieldName:        "GuestOsFeature",
			})
		}
		if len(toRemove) > 0 {
			diffs = append(diffs, imageDiff{
				RequiresRecreate: true,
				FieldName:        "GuestOsFeature",
			})
		}
	}
	if compareImageImageEncryptionKey(c, desired.ImageEncryptionKey, actual.ImageEncryptionKey) {
		c.Config.Logger.Infof("Detected diff in ImageEncryptionKey.\nDESIRED: %#v\nACTUAL: %#v", desired.ImageEncryptionKey, actual.ImageEncryptionKey)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "ImageEncryptionKey",
		})
	}
	if !reflect.DeepEqual(desired.Labels, actual.Labels) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %#v\nACTUAL: %#v", desired.Labels, actual.Labels)

		diffs = append(diffs, imageDiff{
			UpdateOp:  &updateImageSetLabelsOperation{},
			FieldName: "Labels",
		})

	}
	if !dcl.SliceEquals(desired.License, actual.License) {
		c.Config.Logger.Infof("Detected diff in License.\nDESIRED: %#v\nACTUAL: %#v", desired.License, actual.License)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "License",
		})
	}
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %#v\nACTUAL: %#v", desired.Name, actual.Name)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if compareImageShieldedInstanceInitialState(c, desired.ShieldedInstanceInitialState, actual.ShieldedInstanceInitialState) {
		c.Config.Logger.Infof("Detected diff in ShieldedInstanceInitialState.\nDESIRED: %#v\nACTUAL: %#v", desired.ShieldedInstanceInitialState, actual.ShieldedInstanceInitialState)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "ShieldedInstanceInitialState",
		})
	}
	if !dcl.IsZeroValue(desired.SourceDisk) && (dcl.IsZeroValue(actual.SourceDisk) || !reflect.DeepEqual(*desired.SourceDisk, *actual.SourceDisk)) {
		c.Config.Logger.Infof("Detected diff in SourceDisk.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceDisk, actual.SourceDisk)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "SourceDisk",
		})
	}
	if compareImageSourceDiskEncryptionKey(c, desired.SourceDiskEncryptionKey, actual.SourceDiskEncryptionKey) {
		c.Config.Logger.Infof("Detected diff in SourceDiskEncryptionKey.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceDiskEncryptionKey, actual.SourceDiskEncryptionKey)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "SourceDiskEncryptionKey",
		})
	}
	if !dcl.IsZeroValue(desired.SourceImage) && (dcl.IsZeroValue(actual.SourceImage) || !reflect.DeepEqual(*desired.SourceImage, *actual.SourceImage)) {
		c.Config.Logger.Infof("Detected diff in SourceImage.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceImage, actual.SourceImage)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "SourceImage",
		})
	}
	if compareImageSourceImageEncryptionKey(c, desired.SourceImageEncryptionKey, actual.SourceImageEncryptionKey) {
		c.Config.Logger.Infof("Detected diff in SourceImageEncryptionKey.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceImageEncryptionKey, actual.SourceImageEncryptionKey)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "SourceImageEncryptionKey",
		})
	}
	if !dcl.IsZeroValue(desired.SourceImageId) && (dcl.IsZeroValue(actual.SourceImageId) || !reflect.DeepEqual(*desired.SourceImageId, *actual.SourceImageId)) {
		c.Config.Logger.Infof("Detected diff in SourceImageId.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceImageId, actual.SourceImageId)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "SourceImageId",
		})
	}
	if !dcl.IsZeroValue(desired.SourceSnapshot) && (dcl.IsZeroValue(actual.SourceSnapshot) || !reflect.DeepEqual(*desired.SourceSnapshot, *actual.SourceSnapshot)) {
		c.Config.Logger.Infof("Detected diff in SourceSnapshot.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceSnapshot, actual.SourceSnapshot)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "SourceSnapshot",
		})
	}
	if compareImageSourceSnapshotEncryptionKey(c, desired.SourceSnapshotEncryptionKey, actual.SourceSnapshotEncryptionKey) {
		c.Config.Logger.Infof("Detected diff in SourceSnapshotEncryptionKey.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceSnapshotEncryptionKey, actual.SourceSnapshotEncryptionKey)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "SourceSnapshotEncryptionKey",
		})
	}
	if !dcl.IsZeroValue(desired.SourceSnapshotId) && (dcl.IsZeroValue(actual.SourceSnapshotId) || !reflect.DeepEqual(*desired.SourceSnapshotId, *actual.SourceSnapshotId)) {
		c.Config.Logger.Infof("Detected diff in SourceSnapshotId.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceSnapshotId, actual.SourceSnapshotId)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "SourceSnapshotId",
		})
	}
	if !dcl.IsZeroValue(desired.SourceType) && (dcl.IsZeroValue(actual.SourceType) || !reflect.DeepEqual(*desired.SourceType, *actual.SourceType)) {
		c.Config.Logger.Infof("Detected diff in SourceType.\nDESIRED: %#v\nACTUAL: %#v", desired.SourceType, actual.SourceType)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "SourceType",
		})
	}
	if !dcl.SliceEquals(desired.StorageLocation, actual.StorageLocation) {
		c.Config.Logger.Infof("Detected diff in StorageLocation.\nDESIRED: %#v\nACTUAL: %#v", desired.StorageLocation, actual.StorageLocation)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "StorageLocation",
		})
	}
	if compareImageDeprecated(c, desired.Deprecated, actual.Deprecated) {
		c.Config.Logger.Infof("Detected diff in Deprecated.\nDESIRED: %#v\nACTUAL: %#v", desired.Deprecated, actual.Deprecated)

		diffs = append(diffs, imageDiff{
			UpdateOp:  &updateImageDeprecateOperation{},
			FieldName: "Deprecated",
		})

	}
	if !dcl.IsZeroValue(desired.Project) && !dcl.NameToSelfLink(desired.Project, actual.Project) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %#v\nACTUAL: %#v", desired.Project, actual.Project)
		diffs = append(diffs, imageDiff{
			RequiresRecreate: true,
			FieldName:        "Project",
		})
	}
	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []imageDiff
	for _, d := range diffs {
		// Two operations are considered identical if they have the same type.
		// The type of an operation is derived from the name of the update method.
		if !dcl.StringSliceContains(fmt.Sprintf("%T", d.UpdateOp), opTypes) {
			deduped = append(deduped, d)
			opTypes = append(opTypes, fmt.Sprintf("%T", d.UpdateOp))
		} else {
			c.Config.Logger.Infof("Omitting planned operation of type %T since once is already scheduled.", d.UpdateOp)
		}
	}

	return deduped, nil
}
func compareImageGuestOsFeatureSlice(c *Client, desired, actual []ImageGuestOsFeature) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageGuestOsFeature, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageGuestOsFeature(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageGuestOsFeature, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageGuestOsFeatureSets(c *Client, desired, actual []ImageGuestOsFeature) (toAdd, toRemove []ImageGuestOsFeature) {
	if actual == nil {
		return desired, nil
	}
	desiredHashes := make(map[string]ImageGuestOsFeature)
	for _, d := range desired {
		desiredHashes[d.HashCode()] = d
	}
	actualHashes := make(map[string]ImageGuestOsFeature)
	for _, a := range actual {
		actualHashes[a.HashCode()] = a
	}
	toAdd = make([]ImageGuestOsFeature, 0)
	toRemove = make([]ImageGuestOsFeature, 0)

	for k, v := range actualHashes {
		_, found := desiredHashes[k]
		if !found {
			// backup - search linearly for equivalent value.
			for _, des := range desiredHashes {
				if !compareImageGuestOsFeature(c, &des, &v) {
					found = true
					break
				}
			}
		}
		if !found {
			toRemove = append(toRemove, v)
		}
	}

	for k, v := range desiredHashes {
		_, found := actualHashes[k]
		if !found {
			for _, act := range actualHashes {
				if !compareImageGuestOsFeature(c, &v, &act) {
					found = true
					break
				}
			}
		}
		if !found {
			toAdd = append(toAdd, v)
		}
	}

	return toAdd, toRemove
}

func compareImageGuestOsFeature(c *Client, desired, actual *ImageGuestOsFeature) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) && !(dcl.IsEmptyValueIndirect(desired.Type) && dcl.IsZeroValue(actual.Type)) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}
func compareImageImageEncryptionKeySlice(c *Client, desired, actual []ImageImageEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageImageEncryptionKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageImageEncryptionKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageImageEncryptionKey, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageImageEncryptionKey(c *Client, desired, actual *ImageImageEncryptionKey) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RawKey == nil && desired.RawKey != nil && !dcl.IsEmptyValueIndirect(desired.RawKey) {
		c.Config.Logger.Infof("desired RawKey %s - but actually nil", dcl.SprintResource(desired.RawKey))
		return true
	}
	if !reflect.DeepEqual(desired.RawKey, actual.RawKey) && !dcl.IsZeroValue(desired.RawKey) && !(dcl.IsEmptyValueIndirect(desired.RawKey) && dcl.IsZeroValue(actual.RawKey)) {
		c.Config.Logger.Infof("Diff in RawKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RawKey), dcl.SprintResource(actual.RawKey))
		return true
	}
	if actual.KmsKeyName == nil && desired.KmsKeyName != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyName) {
		c.Config.Logger.Infof("desired KmsKeyName %s - but actually nil", dcl.SprintResource(desired.KmsKeyName))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyName, actual.KmsKeyName) && !dcl.IsZeroValue(desired.KmsKeyName) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyName) && dcl.IsZeroValue(actual.KmsKeyName)) {
		c.Config.Logger.Infof("Diff in KmsKeyName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyName), dcl.SprintResource(actual.KmsKeyName))
		return true
	}
	if actual.Sha256 == nil && desired.Sha256 != nil && !dcl.IsEmptyValueIndirect(desired.Sha256) {
		c.Config.Logger.Infof("desired Sha256 %s - but actually nil", dcl.SprintResource(desired.Sha256))
		return true
	}
	if !reflect.DeepEqual(desired.Sha256, actual.Sha256) && !dcl.IsZeroValue(desired.Sha256) && !(dcl.IsEmptyValueIndirect(desired.Sha256) && dcl.IsZeroValue(actual.Sha256)) {
		c.Config.Logger.Infof("Diff in Sha256. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sha256), dcl.SprintResource(actual.Sha256))
		return true
	}
	if actual.KmsKeyServiceAccount == nil && desired.KmsKeyServiceAccount != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyServiceAccount) {
		c.Config.Logger.Infof("desired KmsKeyServiceAccount %s - but actually nil", dcl.SprintResource(desired.KmsKeyServiceAccount))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyServiceAccount, actual.KmsKeyServiceAccount) && !dcl.IsZeroValue(desired.KmsKeyServiceAccount) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyServiceAccount) && dcl.IsZeroValue(actual.KmsKeyServiceAccount)) {
		c.Config.Logger.Infof("Diff in KmsKeyServiceAccount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyServiceAccount), dcl.SprintResource(actual.KmsKeyServiceAccount))
		return true
	}
	return false
}
func compareImageRawDiskSlice(c *Client, desired, actual []ImageRawDisk) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageRawDisk, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageRawDisk(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageRawDisk, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageRawDisk(c *Client, desired, actual *ImageRawDisk) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Source == nil && desired.Source != nil && !dcl.IsEmptyValueIndirect(desired.Source) {
		c.Config.Logger.Infof("desired Source %s - but actually nil", dcl.SprintResource(desired.Source))
		return true
	}
	if !reflect.DeepEqual(desired.Source, actual.Source) && !dcl.IsZeroValue(desired.Source) && !(dcl.IsEmptyValueIndirect(desired.Source) && dcl.IsZeroValue(actual.Source)) {
		c.Config.Logger.Infof("Diff in Source. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Source), dcl.SprintResource(actual.Source))
		return true
	}
	if actual.Sha1Checksum == nil && desired.Sha1Checksum != nil && !dcl.IsEmptyValueIndirect(desired.Sha1Checksum) {
		c.Config.Logger.Infof("desired Sha1Checksum %s - but actually nil", dcl.SprintResource(desired.Sha1Checksum))
		return true
	}
	if !reflect.DeepEqual(desired.Sha1Checksum, actual.Sha1Checksum) && !dcl.IsZeroValue(desired.Sha1Checksum) && !(dcl.IsEmptyValueIndirect(desired.Sha1Checksum) && dcl.IsZeroValue(actual.Sha1Checksum)) {
		c.Config.Logger.Infof("Diff in Sha1Checksum. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sha1Checksum), dcl.SprintResource(actual.Sha1Checksum))
		return true
	}
	if actual.ContainerType == nil && desired.ContainerType != nil && !dcl.IsEmptyValueIndirect(desired.ContainerType) {
		c.Config.Logger.Infof("desired ContainerType %s - but actually nil", dcl.SprintResource(desired.ContainerType))
		return true
	}
	if !reflect.DeepEqual(desired.ContainerType, actual.ContainerType) && !dcl.IsZeroValue(desired.ContainerType) && !(dcl.IsEmptyValueIndirect(desired.ContainerType) && dcl.IsZeroValue(actual.ContainerType)) {
		c.Config.Logger.Infof("Diff in ContainerType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContainerType), dcl.SprintResource(actual.ContainerType))
		return true
	}
	return false
}
func compareImageShieldedInstanceInitialStateSlice(c *Client, desired, actual []ImageShieldedInstanceInitialState) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageShieldedInstanceInitialState, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageShieldedInstanceInitialState(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageShieldedInstanceInitialState, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageShieldedInstanceInitialState(c *Client, desired, actual *ImageShieldedInstanceInitialState) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Pk == nil && desired.Pk != nil && !dcl.IsEmptyValueIndirect(desired.Pk) {
		c.Config.Logger.Infof("desired Pk %s - but actually nil", dcl.SprintResource(desired.Pk))
		return true
	}
	if compareImageShieldedInstanceInitialStatePk(c, desired.Pk, actual.Pk) && !dcl.IsZeroValue(desired.Pk) {
		c.Config.Logger.Infof("Diff in Pk. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Pk), dcl.SprintResource(actual.Pk))
		return true
	}
	if actual.Kek == nil && desired.Kek != nil && !dcl.IsEmptyValueIndirect(desired.Kek) {
		c.Config.Logger.Infof("desired Kek %s - but actually nil", dcl.SprintResource(desired.Kek))
		return true
	}
	if compareImageShieldedInstanceInitialStateKekSlice(c, desired.Kek, actual.Kek) && !dcl.IsZeroValue(desired.Kek) {
		c.Config.Logger.Infof("Diff in Kek. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kek), dcl.SprintResource(actual.Kek))
		return true
	}
	if actual.Db == nil && desired.Db != nil && !dcl.IsEmptyValueIndirect(desired.Db) {
		c.Config.Logger.Infof("desired Db %s - but actually nil", dcl.SprintResource(desired.Db))
		return true
	}
	if compareImageShieldedInstanceInitialStateDbSlice(c, desired.Db, actual.Db) && !dcl.IsZeroValue(desired.Db) {
		c.Config.Logger.Infof("Diff in Db. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Db), dcl.SprintResource(actual.Db))
		return true
	}
	if actual.Dbx == nil && desired.Dbx != nil && !dcl.IsEmptyValueIndirect(desired.Dbx) {
		c.Config.Logger.Infof("desired Dbx %s - but actually nil", dcl.SprintResource(desired.Dbx))
		return true
	}
	if compareImageShieldedInstanceInitialStateDbxSlice(c, desired.Dbx, actual.Dbx) && !dcl.IsZeroValue(desired.Dbx) {
		c.Config.Logger.Infof("Diff in Dbx. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Dbx), dcl.SprintResource(actual.Dbx))
		return true
	}
	return false
}
func compareImageShieldedInstanceInitialStatePkSlice(c *Client, desired, actual []ImageShieldedInstanceInitialStatePk) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageShieldedInstanceInitialStatePk, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageShieldedInstanceInitialStatePk(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageShieldedInstanceInitialStatePk, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageShieldedInstanceInitialStatePk(c *Client, desired, actual *ImageShieldedInstanceInitialStatePk) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Content == nil && desired.Content != nil && !dcl.IsEmptyValueIndirect(desired.Content) {
		c.Config.Logger.Infof("desired Content %s - but actually nil", dcl.SprintResource(desired.Content))
		return true
	}
	if !reflect.DeepEqual(desired.Content, actual.Content) && !dcl.IsZeroValue(desired.Content) && !(dcl.IsEmptyValueIndirect(desired.Content) && dcl.IsZeroValue(actual.Content)) {
		c.Config.Logger.Infof("Diff in Content. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Content), dcl.SprintResource(actual.Content))
		return true
	}
	if actual.FileType == nil && desired.FileType != nil && !dcl.IsEmptyValueIndirect(desired.FileType) {
		c.Config.Logger.Infof("desired FileType %s - but actually nil", dcl.SprintResource(desired.FileType))
		return true
	}
	if !reflect.DeepEqual(desired.FileType, actual.FileType) && !dcl.IsZeroValue(desired.FileType) && !(dcl.IsEmptyValueIndirect(desired.FileType) && dcl.IsZeroValue(actual.FileType)) {
		c.Config.Logger.Infof("Diff in FileType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileType), dcl.SprintResource(actual.FileType))
		return true
	}
	return false
}
func compareImageShieldedInstanceInitialStateKekSlice(c *Client, desired, actual []ImageShieldedInstanceInitialStateKek) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageShieldedInstanceInitialStateKek, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageShieldedInstanceInitialStateKek(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageShieldedInstanceInitialStateKek, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageShieldedInstanceInitialStateKek(c *Client, desired, actual *ImageShieldedInstanceInitialStateKek) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Content == nil && desired.Content != nil && !dcl.IsEmptyValueIndirect(desired.Content) {
		c.Config.Logger.Infof("desired Content %s - but actually nil", dcl.SprintResource(desired.Content))
		return true
	}
	if !reflect.DeepEqual(desired.Content, actual.Content) && !dcl.IsZeroValue(desired.Content) && !(dcl.IsEmptyValueIndirect(desired.Content) && dcl.IsZeroValue(actual.Content)) {
		c.Config.Logger.Infof("Diff in Content. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Content), dcl.SprintResource(actual.Content))
		return true
	}
	if actual.FileType == nil && desired.FileType != nil && !dcl.IsEmptyValueIndirect(desired.FileType) {
		c.Config.Logger.Infof("desired FileType %s - but actually nil", dcl.SprintResource(desired.FileType))
		return true
	}
	if !reflect.DeepEqual(desired.FileType, actual.FileType) && !dcl.IsZeroValue(desired.FileType) && !(dcl.IsEmptyValueIndirect(desired.FileType) && dcl.IsZeroValue(actual.FileType)) {
		c.Config.Logger.Infof("Diff in FileType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileType), dcl.SprintResource(actual.FileType))
		return true
	}
	return false
}
func compareImageShieldedInstanceInitialStateDbSlice(c *Client, desired, actual []ImageShieldedInstanceInitialStateDb) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageShieldedInstanceInitialStateDb, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageShieldedInstanceInitialStateDb(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageShieldedInstanceInitialStateDb, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageShieldedInstanceInitialStateDb(c *Client, desired, actual *ImageShieldedInstanceInitialStateDb) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Content == nil && desired.Content != nil && !dcl.IsEmptyValueIndirect(desired.Content) {
		c.Config.Logger.Infof("desired Content %s - but actually nil", dcl.SprintResource(desired.Content))
		return true
	}
	if !reflect.DeepEqual(desired.Content, actual.Content) && !dcl.IsZeroValue(desired.Content) && !(dcl.IsEmptyValueIndirect(desired.Content) && dcl.IsZeroValue(actual.Content)) {
		c.Config.Logger.Infof("Diff in Content. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Content), dcl.SprintResource(actual.Content))
		return true
	}
	if actual.FileType == nil && desired.FileType != nil && !dcl.IsEmptyValueIndirect(desired.FileType) {
		c.Config.Logger.Infof("desired FileType %s - but actually nil", dcl.SprintResource(desired.FileType))
		return true
	}
	if !reflect.DeepEqual(desired.FileType, actual.FileType) && !dcl.IsZeroValue(desired.FileType) && !(dcl.IsEmptyValueIndirect(desired.FileType) && dcl.IsZeroValue(actual.FileType)) {
		c.Config.Logger.Infof("Diff in FileType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileType), dcl.SprintResource(actual.FileType))
		return true
	}
	return false
}
func compareImageShieldedInstanceInitialStateDbxSlice(c *Client, desired, actual []ImageShieldedInstanceInitialStateDbx) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageShieldedInstanceInitialStateDbx, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageShieldedInstanceInitialStateDbx(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageShieldedInstanceInitialStateDbx, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageShieldedInstanceInitialStateDbx(c *Client, desired, actual *ImageShieldedInstanceInitialStateDbx) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Content == nil && desired.Content != nil && !dcl.IsEmptyValueIndirect(desired.Content) {
		c.Config.Logger.Infof("desired Content %s - but actually nil", dcl.SprintResource(desired.Content))
		return true
	}
	if !reflect.DeepEqual(desired.Content, actual.Content) && !dcl.IsZeroValue(desired.Content) && !(dcl.IsEmptyValueIndirect(desired.Content) && dcl.IsZeroValue(actual.Content)) {
		c.Config.Logger.Infof("Diff in Content. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Content), dcl.SprintResource(actual.Content))
		return true
	}
	if actual.FileType == nil && desired.FileType != nil && !dcl.IsEmptyValueIndirect(desired.FileType) {
		c.Config.Logger.Infof("desired FileType %s - but actually nil", dcl.SprintResource(desired.FileType))
		return true
	}
	if !reflect.DeepEqual(desired.FileType, actual.FileType) && !dcl.IsZeroValue(desired.FileType) && !(dcl.IsEmptyValueIndirect(desired.FileType) && dcl.IsZeroValue(actual.FileType)) {
		c.Config.Logger.Infof("Diff in FileType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileType), dcl.SprintResource(actual.FileType))
		return true
	}
	return false
}
func compareImageSourceDiskEncryptionKeySlice(c *Client, desired, actual []ImageSourceDiskEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageSourceDiskEncryptionKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageSourceDiskEncryptionKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageSourceDiskEncryptionKey, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageSourceDiskEncryptionKey(c *Client, desired, actual *ImageSourceDiskEncryptionKey) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RawKey == nil && desired.RawKey != nil && !dcl.IsEmptyValueIndirect(desired.RawKey) {
		c.Config.Logger.Infof("desired RawKey %s - but actually nil", dcl.SprintResource(desired.RawKey))
		return true
	}
	if !reflect.DeepEqual(desired.RawKey, actual.RawKey) && !dcl.IsZeroValue(desired.RawKey) && !(dcl.IsEmptyValueIndirect(desired.RawKey) && dcl.IsZeroValue(actual.RawKey)) {
		c.Config.Logger.Infof("Diff in RawKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RawKey), dcl.SprintResource(actual.RawKey))
		return true
	}
	if actual.KmsKeyName == nil && desired.KmsKeyName != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyName) {
		c.Config.Logger.Infof("desired KmsKeyName %s - but actually nil", dcl.SprintResource(desired.KmsKeyName))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyName, actual.KmsKeyName) && !dcl.IsZeroValue(desired.KmsKeyName) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyName) && dcl.IsZeroValue(actual.KmsKeyName)) {
		c.Config.Logger.Infof("Diff in KmsKeyName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyName), dcl.SprintResource(actual.KmsKeyName))
		return true
	}
	if actual.Sha256 == nil && desired.Sha256 != nil && !dcl.IsEmptyValueIndirect(desired.Sha256) {
		c.Config.Logger.Infof("desired Sha256 %s - but actually nil", dcl.SprintResource(desired.Sha256))
		return true
	}
	if !reflect.DeepEqual(desired.Sha256, actual.Sha256) && !dcl.IsZeroValue(desired.Sha256) && !(dcl.IsEmptyValueIndirect(desired.Sha256) && dcl.IsZeroValue(actual.Sha256)) {
		c.Config.Logger.Infof("Diff in Sha256. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sha256), dcl.SprintResource(actual.Sha256))
		return true
	}
	if actual.KmsKeyServiceAccount == nil && desired.KmsKeyServiceAccount != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyServiceAccount) {
		c.Config.Logger.Infof("desired KmsKeyServiceAccount %s - but actually nil", dcl.SprintResource(desired.KmsKeyServiceAccount))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyServiceAccount, actual.KmsKeyServiceAccount) && !dcl.IsZeroValue(desired.KmsKeyServiceAccount) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyServiceAccount) && dcl.IsZeroValue(actual.KmsKeyServiceAccount)) {
		c.Config.Logger.Infof("Diff in KmsKeyServiceAccount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyServiceAccount), dcl.SprintResource(actual.KmsKeyServiceAccount))
		return true
	}
	return false
}
func compareImageSourceImageEncryptionKeySlice(c *Client, desired, actual []ImageSourceImageEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageSourceImageEncryptionKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageSourceImageEncryptionKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageSourceImageEncryptionKey, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageSourceImageEncryptionKey(c *Client, desired, actual *ImageSourceImageEncryptionKey) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RawKey == nil && desired.RawKey != nil && !dcl.IsEmptyValueIndirect(desired.RawKey) {
		c.Config.Logger.Infof("desired RawKey %s - but actually nil", dcl.SprintResource(desired.RawKey))
		return true
	}
	if !reflect.DeepEqual(desired.RawKey, actual.RawKey) && !dcl.IsZeroValue(desired.RawKey) && !(dcl.IsEmptyValueIndirect(desired.RawKey) && dcl.IsZeroValue(actual.RawKey)) {
		c.Config.Logger.Infof("Diff in RawKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RawKey), dcl.SprintResource(actual.RawKey))
		return true
	}
	if actual.KmsKeyName == nil && desired.KmsKeyName != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyName) {
		c.Config.Logger.Infof("desired KmsKeyName %s - but actually nil", dcl.SprintResource(desired.KmsKeyName))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyName, actual.KmsKeyName) && !dcl.IsZeroValue(desired.KmsKeyName) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyName) && dcl.IsZeroValue(actual.KmsKeyName)) {
		c.Config.Logger.Infof("Diff in KmsKeyName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyName), dcl.SprintResource(actual.KmsKeyName))
		return true
	}
	if actual.Sha256 == nil && desired.Sha256 != nil && !dcl.IsEmptyValueIndirect(desired.Sha256) {
		c.Config.Logger.Infof("desired Sha256 %s - but actually nil", dcl.SprintResource(desired.Sha256))
		return true
	}
	if !reflect.DeepEqual(desired.Sha256, actual.Sha256) && !dcl.IsZeroValue(desired.Sha256) && !(dcl.IsEmptyValueIndirect(desired.Sha256) && dcl.IsZeroValue(actual.Sha256)) {
		c.Config.Logger.Infof("Diff in Sha256. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sha256), dcl.SprintResource(actual.Sha256))
		return true
	}
	if actual.KmsKeyServiceAccount == nil && desired.KmsKeyServiceAccount != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyServiceAccount) {
		c.Config.Logger.Infof("desired KmsKeyServiceAccount %s - but actually nil", dcl.SprintResource(desired.KmsKeyServiceAccount))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyServiceAccount, actual.KmsKeyServiceAccount) && !dcl.IsZeroValue(desired.KmsKeyServiceAccount) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyServiceAccount) && dcl.IsZeroValue(actual.KmsKeyServiceAccount)) {
		c.Config.Logger.Infof("Diff in KmsKeyServiceAccount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyServiceAccount), dcl.SprintResource(actual.KmsKeyServiceAccount))
		return true
	}
	return false
}
func compareImageSourceSnapshotEncryptionKeySlice(c *Client, desired, actual []ImageSourceSnapshotEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageSourceSnapshotEncryptionKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageSourceSnapshotEncryptionKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageSourceSnapshotEncryptionKey, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageSourceSnapshotEncryptionKey(c *Client, desired, actual *ImageSourceSnapshotEncryptionKey) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RawKey == nil && desired.RawKey != nil && !dcl.IsEmptyValueIndirect(desired.RawKey) {
		c.Config.Logger.Infof("desired RawKey %s - but actually nil", dcl.SprintResource(desired.RawKey))
		return true
	}
	if !reflect.DeepEqual(desired.RawKey, actual.RawKey) && !dcl.IsZeroValue(desired.RawKey) && !(dcl.IsEmptyValueIndirect(desired.RawKey) && dcl.IsZeroValue(actual.RawKey)) {
		c.Config.Logger.Infof("Diff in RawKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RawKey), dcl.SprintResource(actual.RawKey))
		return true
	}
	if actual.KmsKeyName == nil && desired.KmsKeyName != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyName) {
		c.Config.Logger.Infof("desired KmsKeyName %s - but actually nil", dcl.SprintResource(desired.KmsKeyName))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyName, actual.KmsKeyName) && !dcl.IsZeroValue(desired.KmsKeyName) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyName) && dcl.IsZeroValue(actual.KmsKeyName)) {
		c.Config.Logger.Infof("Diff in KmsKeyName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyName), dcl.SprintResource(actual.KmsKeyName))
		return true
	}
	if actual.Sha256 == nil && desired.Sha256 != nil && !dcl.IsEmptyValueIndirect(desired.Sha256) {
		c.Config.Logger.Infof("desired Sha256 %s - but actually nil", dcl.SprintResource(desired.Sha256))
		return true
	}
	if !reflect.DeepEqual(desired.Sha256, actual.Sha256) && !dcl.IsZeroValue(desired.Sha256) && !(dcl.IsEmptyValueIndirect(desired.Sha256) && dcl.IsZeroValue(actual.Sha256)) {
		c.Config.Logger.Infof("Diff in Sha256. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sha256), dcl.SprintResource(actual.Sha256))
		return true
	}
	if actual.KmsKeyServiceAccount == nil && desired.KmsKeyServiceAccount != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyServiceAccount) {
		c.Config.Logger.Infof("desired KmsKeyServiceAccount %s - but actually nil", dcl.SprintResource(desired.KmsKeyServiceAccount))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyServiceAccount, actual.KmsKeyServiceAccount) && !dcl.IsZeroValue(desired.KmsKeyServiceAccount) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyServiceAccount) && dcl.IsZeroValue(actual.KmsKeyServiceAccount)) {
		c.Config.Logger.Infof("Diff in KmsKeyServiceAccount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyServiceAccount), dcl.SprintResource(actual.KmsKeyServiceAccount))
		return true
	}
	return false
}
func compareImageDeprecatedSlice(c *Client, desired, actual []ImageDeprecated) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageDeprecated, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageDeprecated(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageDeprecated, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageDeprecated(c *Client, desired, actual *ImageDeprecated) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.State == nil && desired.State != nil && !dcl.IsEmptyValueIndirect(desired.State) {
		c.Config.Logger.Infof("desired State %s - but actually nil", dcl.SprintResource(desired.State))
		return true
	}
	if !reflect.DeepEqual(desired.State, actual.State) && !dcl.IsZeroValue(desired.State) && !(dcl.IsEmptyValueIndirect(desired.State) && dcl.IsZeroValue(actual.State)) {
		c.Config.Logger.Infof("Diff in State. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.State), dcl.SprintResource(actual.State))
		return true
	}
	if actual.Replacement == nil && desired.Replacement != nil && !dcl.IsEmptyValueIndirect(desired.Replacement) {
		c.Config.Logger.Infof("desired Replacement %s - but actually nil", dcl.SprintResource(desired.Replacement))
		return true
	}
	if !reflect.DeepEqual(desired.Replacement, actual.Replacement) && !dcl.IsZeroValue(desired.Replacement) && !(dcl.IsEmptyValueIndirect(desired.Replacement) && dcl.IsZeroValue(actual.Replacement)) {
		c.Config.Logger.Infof("Diff in Replacement. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Replacement), dcl.SprintResource(actual.Replacement))
		return true
	}
	if actual.Deprecated == nil && desired.Deprecated != nil && !dcl.IsEmptyValueIndirect(desired.Deprecated) {
		c.Config.Logger.Infof("desired Deprecated %s - but actually nil", dcl.SprintResource(desired.Deprecated))
		return true
	}
	if !reflect.DeepEqual(desired.Deprecated, actual.Deprecated) && !dcl.IsZeroValue(desired.Deprecated) && !(dcl.IsEmptyValueIndirect(desired.Deprecated) && dcl.IsZeroValue(actual.Deprecated)) {
		c.Config.Logger.Infof("Diff in Deprecated. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Deprecated), dcl.SprintResource(actual.Deprecated))
		return true
	}
	if actual.Obsolete == nil && desired.Obsolete != nil && !dcl.IsEmptyValueIndirect(desired.Obsolete) {
		c.Config.Logger.Infof("desired Obsolete %s - but actually nil", dcl.SprintResource(desired.Obsolete))
		return true
	}
	if !reflect.DeepEqual(desired.Obsolete, actual.Obsolete) && !dcl.IsZeroValue(desired.Obsolete) && !(dcl.IsEmptyValueIndirect(desired.Obsolete) && dcl.IsZeroValue(actual.Obsolete)) {
		c.Config.Logger.Infof("Diff in Obsolete. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Obsolete), dcl.SprintResource(actual.Obsolete))
		return true
	}
	if actual.Deleted == nil && desired.Deleted != nil && !dcl.IsEmptyValueIndirect(desired.Deleted) {
		c.Config.Logger.Infof("desired Deleted %s - but actually nil", dcl.SprintResource(desired.Deleted))
		return true
	}
	if !reflect.DeepEqual(desired.Deleted, actual.Deleted) && !dcl.IsZeroValue(desired.Deleted) && !(dcl.IsEmptyValueIndirect(desired.Deleted) && dcl.IsZeroValue(actual.Deleted)) {
		c.Config.Logger.Infof("Diff in Deleted. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Deleted), dcl.SprintResource(actual.Deleted))
		return true
	}
	return false
}
func compareImageGuestOsFeatureTypeEnumSlice(c *Client, desired, actual []ImageGuestOsFeatureTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageGuestOsFeatureTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageGuestOsFeatureTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageGuestOsFeatureTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageGuestOsFeatureTypeEnum(c *Client, desired, actual *ImageGuestOsFeatureTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareImageRawDiskContainerTypeEnumSlice(c *Client, desired, actual []ImageRawDiskContainerTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageRawDiskContainerTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageRawDiskContainerTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageRawDiskContainerTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageRawDiskContainerTypeEnum(c *Client, desired, actual *ImageRawDiskContainerTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareImageShieldedInstanceInitialStatePkFileTypeEnumSlice(c *Client, desired, actual []ImageShieldedInstanceInitialStatePkFileTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageShieldedInstanceInitialStatePkFileTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageShieldedInstanceInitialStatePkFileTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageShieldedInstanceInitialStatePkFileTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageShieldedInstanceInitialStatePkFileTypeEnum(c *Client, desired, actual *ImageShieldedInstanceInitialStatePkFileTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareImageShieldedInstanceInitialStateKekFileTypeEnumSlice(c *Client, desired, actual []ImageShieldedInstanceInitialStateKekFileTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageShieldedInstanceInitialStateKekFileTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageShieldedInstanceInitialStateKekFileTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageShieldedInstanceInitialStateKekFileTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageShieldedInstanceInitialStateKekFileTypeEnum(c *Client, desired, actual *ImageShieldedInstanceInitialStateKekFileTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareImageShieldedInstanceInitialStateDbFileTypeEnumSlice(c *Client, desired, actual []ImageShieldedInstanceInitialStateDbFileTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageShieldedInstanceInitialStateDbFileTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageShieldedInstanceInitialStateDbFileTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageShieldedInstanceInitialStateDbFileTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageShieldedInstanceInitialStateDbFileTypeEnum(c *Client, desired, actual *ImageShieldedInstanceInitialStateDbFileTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareImageShieldedInstanceInitialStateDbxFileTypeEnumSlice(c *Client, desired, actual []ImageShieldedInstanceInitialStateDbxFileTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageShieldedInstanceInitialStateDbxFileTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageShieldedInstanceInitialStateDbxFileTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageShieldedInstanceInitialStateDbxFileTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageShieldedInstanceInitialStateDbxFileTypeEnum(c *Client, desired, actual *ImageShieldedInstanceInitialStateDbxFileTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareImageSourceTypeEnumSlice(c *Client, desired, actual []ImageSourceTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageSourceTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageSourceTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageSourceTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageSourceTypeEnum(c *Client, desired, actual *ImageSourceTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareImageStatusEnumSlice(c *Client, desired, actual []ImageStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageStatusEnum(c *Client, desired, actual *ImageStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareImageDeprecatedStateEnumSlice(c *Client, desired, actual []ImageDeprecatedStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ImageDeprecatedStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareImageDeprecatedStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ImageDeprecatedStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareImageDeprecatedStateEnum(c *Client, desired, actual *ImageDeprecatedStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Image) urlNormalized() *Image {
	normalized := deepcopy.Copy(*r).(Image)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Image) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Image) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Image) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Image) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "deprecate" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/images/{{name}}/deprecate", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "setLabels" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/images/{{name}}/setLabels", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

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
	if v, err := expandImageGuestOsFeatureSlice(c, f.GuestOsFeature); err != nil {
		return nil, fmt.Errorf("error expanding GuestOsFeature into guestOsFeatures: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["guestOsFeatures"] = v
	}
	if v, err := expandImageImageEncryptionKey(c, f.ImageEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding ImageEncryptionKey into imageEncryptionKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["imageEncryptionKey"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.License; !dcl.IsEmptyValueIndirect(v) {
		m["licenses"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandImageRawDisk(c, f.RawDisk); err != nil {
		return nil, fmt.Errorf("error expanding RawDisk into rawDisk: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rawDisk"] = v
	}
	if v, err := expandImageShieldedInstanceInitialState(c, f.ShieldedInstanceInitialState); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedInstanceInitialState into shieldedInstanceInitialState: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	if v := f.StorageLocation; !dcl.IsEmptyValueIndirect(v) {
		m["storageLocations"] = v
	}
	if v, err := expandImageDeprecated(c, f.Deprecated); err != nil {
		return nil, fmt.Errorf("error expanding Deprecated into deprecated: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["deprecated"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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

	r := &Image{}
	r.ArchiveSizeBytes = dcl.FlattenInteger(m["archiveSizeBytes"])
	r.Description = dcl.FlattenString(m["description"])
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.Family = dcl.FlattenString(m["family"])
	r.GuestOsFeature = flattenImageGuestOsFeatureSlice(c, m["guestOsFeatures"])
	r.ImageEncryptionKey = flattenImageImageEncryptionKey(c, m["imageEncryptionKey"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.License = dcl.FlattenStringSlice(m["licenses"])
	r.Name = dcl.FlattenString(m["name"])
	r.RawDisk = flattenImageRawDisk(c, m["rawDisk"])
	r.ShieldedInstanceInitialState = flattenImageShieldedInstanceInitialState(c, m["shieldedInstanceInitialState"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.SourceDisk = dcl.FlattenString(m["sourceDisk"])
	r.SourceDiskEncryptionKey = flattenImageSourceDiskEncryptionKey(c, m["sourceDiskEncryptionKey"])
	r.SourceDiskId = dcl.FlattenString(m["sourceDiskId"])
	r.SourceImage = dcl.FlattenString(m["sourceImage"])
	r.SourceImageEncryptionKey = flattenImageSourceImageEncryptionKey(c, m["sourceImageEncryptionKey"])
	r.SourceImageId = dcl.FlattenString(m["sourceImageId"])
	r.SourceSnapshot = dcl.FlattenString(m["sourceSnapshot"])
	r.SourceSnapshotEncryptionKey = flattenImageSourceSnapshotEncryptionKey(c, m["sourceSnapshotEncryptionKey"])
	r.SourceSnapshotId = dcl.FlattenString(m["sourceSnapshotId"])
	r.SourceType = flattenImageSourceTypeEnum(m["sourceType"])
	r.Status = flattenImageStatusEnum(m["status"])
	r.StorageLocation = dcl.FlattenStringSlice(m["storageLocations"])
	r.Deprecated = flattenImageDeprecated(c, m["deprecated"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandImageGuestOsFeatureMap expands the contents of ImageGuestOsFeature into a JSON
// request object.
func expandImageGuestOsFeatureMap(c *Client, f map[string]ImageGuestOsFeature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandImageGuestOsFeature(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandImageGuestOsFeatureSlice expands the contents of ImageGuestOsFeature into a JSON
// request object.
func expandImageGuestOsFeatureSlice(c *Client, f []ImageGuestOsFeature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandImageGuestOsFeature(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenImageGuestOsFeatureMap flattens the contents of ImageGuestOsFeature from a JSON
// response object.
func flattenImageGuestOsFeatureMap(c *Client, i interface{}) map[string]ImageGuestOsFeature {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ImageGuestOsFeature{}
	}

	if len(a) == 0 {
		return map[string]ImageGuestOsFeature{}
	}

	items := make(map[string]ImageGuestOsFeature)
	for k, item := range a {
		items[k] = *flattenImageGuestOsFeature(c, item.(map[string]interface{}))
	}

	return items
}

// flattenImageGuestOsFeatureSlice flattens the contents of ImageGuestOsFeature from a JSON
// response object.
func flattenImageGuestOsFeatureSlice(c *Client, i interface{}) []ImageGuestOsFeature {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageGuestOsFeature{}
	}

	if len(a) == 0 {
		return []ImageGuestOsFeature{}
	}

	items := make([]ImageGuestOsFeature, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageGuestOsFeature(c, item.(map[string]interface{})))
	}

	return items
}

// expandImageGuestOsFeature expands an instance of ImageGuestOsFeature into a JSON
// request object.
func expandImageGuestOsFeature(c *Client, f *ImageGuestOsFeature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenImageGuestOsFeature flattens an instance of ImageGuestOsFeature from a JSON
// response object.
func flattenImageGuestOsFeature(c *Client, i interface{}) *ImageGuestOsFeature {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ImageGuestOsFeature{}
	r.Type = flattenImageGuestOsFeatureTypeEnum(m["type"])

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
	r.State = flattenImageDeprecatedStateEnum(m["state"])
	r.Replacement = dcl.FlattenString(m["replacement"])
	r.Deprecated = dcl.FlattenString(m["deprecated"])
	r.Obsolete = dcl.FlattenString(m["obsolete"])
	r.Deleted = dcl.FlattenString(m["deleted"])

	return r
}

// flattenImageGuestOsFeatureTypeEnumSlice flattens the contents of ImageGuestOsFeatureTypeEnum from a JSON
// response object.
func flattenImageGuestOsFeatureTypeEnumSlice(c *Client, i interface{}) []ImageGuestOsFeatureTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ImageGuestOsFeatureTypeEnum{}
	}

	if len(a) == 0 {
		return []ImageGuestOsFeatureTypeEnum{}
	}

	items := make([]ImageGuestOsFeatureTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenImageGuestOsFeatureTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenImageGuestOsFeatureTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ImageGuestOsFeatureTypeEnum with the same value as that string.
func flattenImageGuestOsFeatureTypeEnum(i interface{}) *ImageGuestOsFeatureTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ImageGuestOsFeatureTypeEnumRef("")
	}

	return ImageGuestOsFeatureTypeEnumRef(s)
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
		items = append(items, *flattenImageRawDiskContainerTypeEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenImageShieldedInstanceInitialStatePkFileTypeEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenImageShieldedInstanceInitialStateKekFileTypeEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenImageShieldedInstanceInitialStateDbFileTypeEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenImageShieldedInstanceInitialStateDbxFileTypeEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenImageSourceTypeEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenImageStatusEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenImageDeprecatedStateEnum(item.(map[string]interface{})))
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
