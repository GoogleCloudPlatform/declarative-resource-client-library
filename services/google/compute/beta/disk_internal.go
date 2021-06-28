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

func (r *Disk) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DiskEncryptionKey) {
		if err := r.DiskEncryptionKey.validate(); err != nil {
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
	return nil
}
func (r *DiskGuestOsFeature) validate() error {
	return nil
}
func (r *DiskEncryptionKey) validate() error {
	return nil
}
func (r *DiskGuestOsFeatures) validate() error {
	return nil
}

func diskGetURL(userBasePath string, r *Disk) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/disks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/disks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Get URL found")

}

func diskListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/disks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/disks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid List URL found")

}

func diskCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/disks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/disks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Create URL found")

}

func diskDeleteURL(userBasePath string, r *Disk) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/disks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/disks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Delete URL found")

}

// diskApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type diskApiOperation interface {
	do(context.Context, *Disk, *Client) error
}

// newUpdateDiskResizeRequest creates a request for an
// Disk resource's resize update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDiskResizeRequest(ctx context.Context, f *Disk, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.SizeGb; !dcl.IsEmptyValueIndirect(v) {
		req["sizeGb"] = v
	}
	return req, nil
}

// marshalUpdateDiskResizeRequest converts the update into
// the final JSON request body.
func marshalUpdateDiskResizeRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateDiskResizeOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDiskResizeOperation) do(ctx context.Context, r *Disk, c *Client) error {
	_, err := c.GetDisk(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "resize")
	if err != nil {
		return err
	}

	req, err := newUpdateDiskResizeRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateDiskResizeRequest(c, req)
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

// newUpdateDiskSetLabelsRequest creates a request for an
// Disk resource's setLabels update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDiskSetLabelsRequest(ctx context.Context, f *Disk, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.LabelFingerprint; !dcl.IsEmptyValueIndirect(v) {
		req["labelFingerprint"] = v
	}
	b, err := c.getDiskRaw(ctx, f.urlNormalized())
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

// marshalUpdateDiskSetLabelsRequest converts the update into
// the final JSON request body.
func marshalUpdateDiskSetLabelsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateDiskSetLabelsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDiskSetLabelsOperation) do(ctx context.Context, r *Disk, c *Client) error {
	_, err := c.GetDisk(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setLabels")
	if err != nil {
		return err
	}

	req, err := newUpdateDiskSetLabelsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateDiskSetLabelsRequest(c, req)
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

func (c *Client) listDiskRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := diskListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != DiskMaxPage {
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

type listDiskOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listDisk(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Disk, string, error) {
	b, err := c.listDiskRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listDiskOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Disk
	for _, v := range m.Items {
		res, err := unmarshalMapDisk(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllDisk(ctx context.Context, f func(*Disk) bool, resources []*Disk) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteDisk(ctx, res)
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

type deleteDiskOperation struct{}

func (op *deleteDiskOperation) do(ctx context.Context, r *Disk, c *Client) error {

	_, err := c.GetDisk(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Disk not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetDisk checking for existence. error: %v", err)
		return err
	}

	u, err := diskDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetDisk(ctx, r.urlNormalized())
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
type createDiskOperation struct {
	response map[string]interface{}
}

func (op *createDiskOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createDiskOperation) do(ctx context.Context, r *Disk, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := diskCreateURL(c.Config.BasePath, project, location)

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

	if _, err := c.GetDisk(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDiskRaw(ctx context.Context, r *Disk) ([]byte, error) {

	u, err := diskGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) diskDiffsForRawDesired(ctx context.Context, rawDesired *Disk, opts ...dcl.ApplyOption) (initial, desired *Disk, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Disk
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Disk); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Disk, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetDisk(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Disk resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Disk resource: %v", err)
		}
		c.Config.Logger.Info("Found that Disk resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeDiskDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Disk: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Disk: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeDiskInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Disk: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeDiskDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Disk: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffDisk(c, desired, initial, opts...)
	fmt.Printf("newDiffs: %v\n", diffs)
	return initial, desired, diffs, err
}

func canonicalizeDiskInitialState(rawInitial, rawDesired *Disk) (*Disk, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeDiskDesiredState(rawDesired, rawInitial *Disk, opts ...dcl.ApplyOption) (*Disk, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DiskEncryptionKey = canonicalizeDiskEncryptionKey(rawDesired.DiskEncryptionKey, nil, opts...)
		rawDesired.SourceImageEncryptionKey = canonicalizeDiskEncryptionKey(rawDesired.SourceImageEncryptionKey, nil, opts...)
		rawDesired.SourceSnapshotEncryptionKey = canonicalizeDiskEncryptionKey(rawDesired.SourceSnapshotEncryptionKey, nil, opts...)

		return rawDesired, nil
	}

	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	rawDesired.DiskEncryptionKey = canonicalizeDiskEncryptionKey(rawDesired.DiskEncryptionKey, rawInitial.DiskEncryptionKey, opts...)
	if dcl.IsZeroValue(rawDesired.GuestOsFeature) {
		rawDesired.GuestOsFeature = rawInitial.GuestOsFeature
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.License) {
		rawDesired.License = rawInitial.License
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.IsZeroValue(rawDesired.ReplicaZones) {
		rawDesired.ReplicaZones = rawInitial.ReplicaZones
	}
	if dcl.IsZeroValue(rawDesired.ResourcePolicy) {
		rawDesired.ResourcePolicy = rawInitial.ResourcePolicy
	}
	if dcl.IsZeroValue(rawDesired.SizeGb) {
		rawDesired.SizeGb = rawInitial.SizeGb
	}
	if dcl.NameToSelfLink(rawDesired.SourceImage, rawInitial.SourceImage) {
		rawDesired.SourceImage = rawInitial.SourceImage
	}
	rawDesired.SourceImageEncryptionKey = canonicalizeDiskEncryptionKey(rawDesired.SourceImageEncryptionKey, rawInitial.SourceImageEncryptionKey, opts...)
	if dcl.StringCanonicalize(rawDesired.SourceSnapshot, rawInitial.SourceSnapshot) {
		rawDesired.SourceSnapshot = rawInitial.SourceSnapshot
	}
	rawDesired.SourceSnapshotEncryptionKey = canonicalizeDiskEncryptionKey(rawDesired.SourceSnapshotEncryptionKey, rawInitial.SourceSnapshotEncryptionKey, opts...)
	if dcl.NameToSelfLink(rawDesired.Type, rawInitial.Type) {
		rawDesired.Type = rawInitial.Type
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.Options, rawInitial.Options) {
		rawDesired.Options = rawInitial.Options
	}
	if dcl.IsZeroValue(rawDesired.Licenses) {
		rawDesired.Licenses = rawInitial.Licenses
	}
	if dcl.IsZeroValue(rawDesired.GuestOsFeatures) {
		rawDesired.GuestOsFeatures = rawInitial.GuestOsFeatures
	}
	if dcl.IsZeroValue(rawDesired.LicenseCodes) {
		rawDesired.LicenseCodes = rawInitial.LicenseCodes
	}
	if dcl.IsZeroValue(rawDesired.PhysicalBlockSizeBytes) {
		rawDesired.PhysicalBlockSizeBytes = rawInitial.PhysicalBlockSizeBytes
	}
	if dcl.IsZeroValue(rawDesired.ResourcePolicies) {
		rawDesired.ResourcePolicies = rawInitial.ResourcePolicies
	}
	if dcl.StringCanonicalize(rawDesired.SourceDisk, rawInitial.SourceDisk) {
		rawDesired.SourceDisk = rawInitial.SourceDisk
	}
	if dcl.StringCanonicalize(rawDesired.SourceDiskId, rawInitial.SourceDiskId) {
		rawDesired.SourceDiskId = rawInitial.SourceDiskId
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeDiskNewState(c *Client, rawNew, rawDesired *Disk) (*Disk, error) {

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskEncryptionKey) && dcl.IsEmptyValueIndirect(rawDesired.DiskEncryptionKey) {
		rawNew.DiskEncryptionKey = rawDesired.DiskEncryptionKey
	} else {
		rawNew.DiskEncryptionKey = canonicalizeNewDiskEncryptionKey(c, rawDesired.DiskEncryptionKey, rawNew.DiskEncryptionKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.GuestOsFeature) && dcl.IsEmptyValueIndirect(rawDesired.GuestOsFeature) {
		rawNew.GuestOsFeature = rawDesired.GuestOsFeature
	} else {
		rawNew.GuestOsFeature = canonicalizeNewDiskGuestOsFeatureSlice(c, rawDesired.GuestOsFeature, rawNew.GuestOsFeature)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LabelFingerprint) && dcl.IsEmptyValueIndirect(rawDesired.LabelFingerprint) {
		rawNew.LabelFingerprint = rawDesired.LabelFingerprint
	} else {
		if dcl.StringCanonicalize(rawDesired.LabelFingerprint, rawNew.LabelFingerprint) {
			rawNew.LabelFingerprint = rawDesired.LabelFingerprint
		}
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

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReplicaZones) && dcl.IsEmptyValueIndirect(rawDesired.ReplicaZones) {
		rawNew.ReplicaZones = rawDesired.ReplicaZones
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourcePolicy) && dcl.IsEmptyValueIndirect(rawDesired.ResourcePolicy) {
		rawNew.ResourcePolicy = rawDesired.ResourcePolicy
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SizeGb) && dcl.IsEmptyValueIndirect(rawDesired.SizeGb) {
		rawNew.SizeGb = rawDesired.SizeGb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceImage) && dcl.IsEmptyValueIndirect(rawDesired.SourceImage) {
		rawNew.SourceImage = rawDesired.SourceImage
	} else {
		if dcl.NameToSelfLink(rawDesired.SourceImage, rawNew.SourceImage) {
			rawNew.SourceImage = rawDesired.SourceImage
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceImageEncryptionKey) && dcl.IsEmptyValueIndirect(rawDesired.SourceImageEncryptionKey) {
		rawNew.SourceImageEncryptionKey = rawDesired.SourceImageEncryptionKey
	} else {
		rawNew.SourceImageEncryptionKey = canonicalizeNewDiskEncryptionKey(c, rawDesired.SourceImageEncryptionKey, rawNew.SourceImageEncryptionKey)
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
		rawNew.SourceSnapshotEncryptionKey = canonicalizeNewDiskEncryptionKey(c, rawDesired.SourceSnapshotEncryptionKey, rawNew.SourceSnapshotEncryptionKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceSnapshotId) && dcl.IsEmptyValueIndirect(rawDesired.SourceSnapshotId) {
		rawNew.SourceSnapshotId = rawDesired.SourceSnapshotId
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceSnapshotId, rawNew.SourceSnapshotId) {
			rawNew.SourceSnapshotId = rawDesired.SourceSnapshotId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
		if dcl.NameToSelfLink(rawDesired.Type, rawNew.Type) {
			rawNew.Type = rawDesired.Type
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Zone) && dcl.IsEmptyValueIndirect(rawDesired.Zone) {
		rawNew.Zone = rawDesired.Zone
	} else {
		if dcl.StringCanonicalize(rawDesired.Zone, rawNew.Zone) {
			rawNew.Zone = rawDesired.Zone
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Options) && dcl.IsEmptyValueIndirect(rawDesired.Options) {
		rawNew.Options = rawDesired.Options
	} else {
		if dcl.StringCanonicalize(rawDesired.Options, rawNew.Options) {
			rawNew.Options = rawDesired.Options
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Licenses) && dcl.IsEmptyValueIndirect(rawDesired.Licenses) {
		rawNew.Licenses = rawDesired.Licenses
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.GuestOsFeatures) && dcl.IsEmptyValueIndirect(rawDesired.GuestOsFeatures) {
		rawNew.GuestOsFeatures = rawDesired.GuestOsFeatures
	} else {
		rawNew.GuestOsFeatures = canonicalizeNewDiskGuestOsFeaturesSlice(c, rawDesired.GuestOsFeatures, rawNew.GuestOsFeatures)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastAttachTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.LastAttachTimestamp) {
		rawNew.LastAttachTimestamp = rawDesired.LastAttachTimestamp
	} else {
		if dcl.StringCanonicalize(rawDesired.LastAttachTimestamp, rawNew.LastAttachTimestamp) {
			rawNew.LastAttachTimestamp = rawDesired.LastAttachTimestamp
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastDetachTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.LastDetachTimestamp) {
		rawNew.LastDetachTimestamp = rawDesired.LastDetachTimestamp
	} else {
		if dcl.StringCanonicalize(rawDesired.LastDetachTimestamp, rawNew.LastDetachTimestamp) {
			rawNew.LastDetachTimestamp = rawDesired.LastDetachTimestamp
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Users) && dcl.IsEmptyValueIndirect(rawDesired.Users) {
		rawNew.Users = rawDesired.Users
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LicenseCodes) && dcl.IsEmptyValueIndirect(rawDesired.LicenseCodes) {
		rawNew.LicenseCodes = rawDesired.LicenseCodes
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PhysicalBlockSizeBytes) && dcl.IsEmptyValueIndirect(rawDesired.PhysicalBlockSizeBytes) {
		rawNew.PhysicalBlockSizeBytes = rawDesired.PhysicalBlockSizeBytes
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourcePolicies) && dcl.IsEmptyValueIndirect(rawDesired.ResourcePolicies) {
		rawNew.ResourcePolicies = rawDesired.ResourcePolicies
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceDisk) && dcl.IsEmptyValueIndirect(rawDesired.SourceDisk) {
		rawNew.SourceDisk = rawDesired.SourceDisk
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceDisk, rawNew.SourceDisk) {
			rawNew.SourceDisk = rawDesired.SourceDisk
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceDiskId) && dcl.IsEmptyValueIndirect(rawDesired.SourceDiskId) {
		rawNew.SourceDiskId = rawDesired.SourceDiskId
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceDiskId, rawNew.SourceDiskId) {
			rawNew.SourceDiskId = rawDesired.SourceDiskId
		}
	}

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeDiskGuestOsFeature(des, initial *DiskGuestOsFeature, opts ...dcl.ApplyOption) *DiskGuestOsFeature {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.IsZeroValue(des.TypeAlt) {
		des.TypeAlt = initial.TypeAlt
	}

	return des
}

func canonicalizeNewDiskGuestOsFeature(c *Client, des, nw *DiskGuestOsFeature) *DiskGuestOsFeature {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
	}
	if dcl.IsZeroValue(nw.TypeAlt) {
		nw.TypeAlt = des.TypeAlt
	}

	return nw
}

func canonicalizeNewDiskGuestOsFeatureSet(c *Client, des, nw []DiskGuestOsFeature) []DiskGuestOsFeature {
	if des == nil {
		return nw
	}
	var reorderedNew []DiskGuestOsFeature
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDiskGuestOsFeatureNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDiskGuestOsFeatureSlice(c *Client, des, nw []DiskGuestOsFeature) []DiskGuestOsFeature {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DiskGuestOsFeature
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDiskGuestOsFeature(c, &d, &n))
	}

	return items
}

func canonicalizeDiskEncryptionKey(des, initial *DiskEncryptionKey, opts ...dcl.ApplyOption) *DiskEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.RawKey, initial.RawKey) || dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}
	if dcl.StringCanonicalize(des.KmsKeyServiceAccount, initial.KmsKeyServiceAccount) || dcl.IsZeroValue(des.KmsKeyServiceAccount) {
		des.KmsKeyServiceAccount = initial.KmsKeyServiceAccount
	}

	return des
}

func canonicalizeNewDiskEncryptionKey(c *Client, des, nw *DiskEncryptionKey) *DiskEncryptionKey {
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

func canonicalizeNewDiskEncryptionKeySet(c *Client, des, nw []DiskEncryptionKey) []DiskEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []DiskEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDiskEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDiskEncryptionKeySlice(c *Client, des, nw []DiskEncryptionKey) []DiskEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DiskEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDiskEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeDiskGuestOsFeatures(des, initial *DiskGuestOsFeatures, opts ...dcl.ApplyOption) *DiskGuestOsFeatures {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.IsZeroValue(des.TypeAlts) {
		des.TypeAlts = initial.TypeAlts
	}

	return des
}

func canonicalizeNewDiskGuestOsFeatures(c *Client, des, nw *DiskGuestOsFeatures) *DiskGuestOsFeatures {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
	}
	if dcl.IsZeroValue(nw.TypeAlts) {
		nw.TypeAlts = des.TypeAlts
	}

	return nw
}

func canonicalizeNewDiskGuestOsFeaturesSet(c *Client, des, nw []DiskGuestOsFeatures) []DiskGuestOsFeatures {
	if des == nil {
		return nw
	}
	var reorderedNew []DiskGuestOsFeatures
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDiskGuestOsFeaturesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDiskGuestOsFeaturesSlice(c *Client, des, nw []DiskGuestOsFeatures) []DiskGuestOsFeatures {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DiskGuestOsFeatures
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDiskGuestOsFeatures(c, &d, &n))
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
func diffDisk(c *Client, desired, actual *Disk, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.DiskEncryptionKey, actual.DiskEncryptionKey, dcl.Info{ObjectFunction: compareDiskEncryptionKeyNewStyle, EmptyObject: EmptyDiskEncryptionKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GuestOsFeature, actual.GuestOsFeature, dcl.Info{ObjectFunction: compareDiskGuestOsFeatureNewStyle, EmptyObject: EmptyDiskGuestOsFeature, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GuestOsFeature")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDiskSetLabelsOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LabelFingerprint, actual.LabelFingerprint, dcl.Info{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateDiskSetLabelsOperation")}, fn.AddNest("LabelFingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.License, actual.License, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("License")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReplicaZones, actual.ReplicaZones, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReplicaZones")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourcePolicy, actual.ResourcePolicy, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourcePolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SizeGb, actual.SizeGb, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDiskResizeOperation")}, fn.AddNest("SizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceImage, actual.SourceImage, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceImage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceImageEncryptionKey, actual.SourceImageEncryptionKey, dcl.Info{ObjectFunction: compareDiskEncryptionKeyNewStyle, EmptyObject: EmptyDiskEncryptionKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceImageEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceImageId, actual.SourceImageId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceImageId")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.SourceSnapshotEncryptionKey, actual.SourceSnapshotEncryptionKey, dcl.Info{ObjectFunction: compareDiskEncryptionKeyNewStyle, EmptyObject: EmptyDiskEncryptionKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceSnapshotEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceSnapshotId, actual.SourceSnapshotId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceSnapshotId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zone")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Options, actual.Options, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Options")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Licenses, actual.Licenses, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Licenses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GuestOsFeatures, actual.GuestOsFeatures, dcl.Info{ObjectFunction: compareDiskGuestOsFeaturesNewStyle, EmptyObject: EmptyDiskGuestOsFeatures, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GuestOsFeatures")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastAttachTimestamp, actual.LastAttachTimestamp, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastAttachTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastDetachTimestamp, actual.LastDetachTimestamp, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastDetachTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Users, actual.Users, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Users")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LicenseCodes, actual.LicenseCodes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LicenseCodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PhysicalBlockSizeBytes, actual.PhysicalBlockSizeBytes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PhysicalBlockSizeBytes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourcePolicies, actual.ResourcePolicies, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourcePolicies")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.SourceDiskId, actual.SourceDiskId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceDiskId")); len(ds) != 0 || err != nil {
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

	return newDiffs, nil
}
func compareDiskGuestOsFeatureNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DiskGuestOsFeature)
	if !ok {
		desiredNotPointer, ok := d.(DiskGuestOsFeature)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DiskGuestOsFeature or *DiskGuestOsFeature", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DiskGuestOsFeature)
	if !ok {
		actualNotPointer, ok := a.(DiskGuestOsFeature)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DiskGuestOsFeature", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TypeAlt, actual.TypeAlt, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TypeAlt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDiskEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DiskEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(DiskEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DiskEncryptionKey or *DiskEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DiskEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(DiskEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DiskEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{Ignore: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Sha256, actual.Sha256, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256")); len(ds) != 0 || err != nil {
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

func compareDiskGuestOsFeaturesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DiskGuestOsFeatures)
	if !ok {
		desiredNotPointer, ok := d.(DiskGuestOsFeatures)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DiskGuestOsFeatures or *DiskGuestOsFeatures", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DiskGuestOsFeatures)
	if !ok {
		actualNotPointer, ok := a.(DiskGuestOsFeatures)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DiskGuestOsFeatures", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TypeAlts, actual.TypeAlts, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TypeAlts")); len(ds) != 0 || err != nil {
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
func (r *Disk) urlNormalized() *Disk {
	normalized := dcl.Copy(*r).(Disk)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.LabelFingerprint = dcl.SelfLinkToName(r.LabelFingerprint)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.SourceImage = dcl.SelfLinkToName(r.SourceImage)
	normalized.SourceImageId = dcl.SelfLinkToName(r.SourceImageId)
	normalized.SourceSnapshot = dcl.SelfLinkToName(r.SourceSnapshot)
	normalized.SourceSnapshotId = dcl.SelfLinkToName(r.SourceSnapshotId)
	normalized.Type = dcl.SelfLinkToName(r.Type)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Options = dcl.SelfLinkToName(r.Options)
	normalized.LastAttachTimestamp = dcl.SelfLinkToName(r.LastAttachTimestamp)
	normalized.LastDetachTimestamp = dcl.SelfLinkToName(r.LastDetachTimestamp)
	normalized.SourceDisk = dcl.SelfLinkToName(r.SourceDisk)
	normalized.SourceDiskId = dcl.SelfLinkToName(r.SourceDiskId)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Disk) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Disk) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *Disk) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Disk) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "resize" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		if dcl.IsRegion(r.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/disks/{{name}}/resize", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil
		}

		if dcl.IsZone(r.Location) {
			return dcl.URL("projects/{{project}}/zones/{{location}}/disks/{{name}}/resize", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil
		}

		return "", fmt.Errorf("No valid update URL for %s", updateName)

	}
	if updateName == "setLabels" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		if dcl.IsRegion(r.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/disks/{{name}}/setLabels", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil
		}

		if dcl.IsZone(r.Location) {
			return dcl.URL("projects/{{project}}/zones/{{location}}/disks/{{name}}/setLabels", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil
		}

		return "", fmt.Errorf("No valid update URL for %s", updateName)

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Disk resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Disk) marshal(c *Client) ([]byte, error) {
	m, err := expandDisk(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Disk: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalDisk decodes JSON responses into the Disk resource schema.
func unmarshalDisk(b []byte, c *Client) (*Disk, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapDisk(m, c)
}

func unmarshalMapDisk(m map[string]interface{}, c *Client) (*Disk, error) {

	return flattenDisk(c, m), nil
}

// expandDisk expands Disk into a JSON request object.
func expandDisk(c *Client, f *Disk) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandDiskEncryptionKey(c, f.DiskEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding DiskEncryptionKey into diskEncryptionKey: %w", err)
	} else if v != nil {
		m["diskEncryptionKey"] = v
	}
	if v, err := expandDiskGuestOsFeatureSlice(c, f.GuestOsFeature); err != nil {
		return nil, fmt.Errorf("error expanding GuestOsFeature into guestOsFeature: %w", err)
	} else if v != nil {
		m["guestOsFeature"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.LabelFingerprint; !dcl.IsEmptyValueIndirect(v) {
		m["labelFingerprint"] = v
	}
	if v := f.License; !dcl.IsEmptyValueIndirect(v) {
		m["license"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v := f.ReplicaZones; !dcl.IsEmptyValueIndirect(v) {
		m["replicaZones"] = v
	}
	if v := f.ResourcePolicy; !dcl.IsEmptyValueIndirect(v) {
		m["resourcePolicy"] = v
	}
	if v := f.SizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["sizeGb"] = v
	}
	if v := f.SourceImage; !dcl.IsEmptyValueIndirect(v) {
		m["sourceImage"] = v
	}
	if v, err := expandDiskEncryptionKey(c, f.SourceImageEncryptionKey); err != nil {
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
	if v, err := expandDiskEncryptionKey(c, f.SourceSnapshotEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding SourceSnapshotEncryptionKey into sourceSnapshotEncryptionKey: %w", err)
	} else if v != nil {
		m["sourceSnapshotEncryptionKey"] = v
	}
	if v := f.SourceSnapshotId; !dcl.IsEmptyValueIndirect(v) {
		m["sourceSnapshotId"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.Options; !dcl.IsEmptyValueIndirect(v) {
		m["options"] = v
	}
	if v := f.Licenses; !dcl.IsEmptyValueIndirect(v) {
		m["licenses"] = v
	}
	if v, err := expandDiskGuestOsFeaturesSlice(c, f.GuestOsFeatures); err != nil {
		return nil, fmt.Errorf("error expanding GuestOsFeatures into guestOsFeatures: %w", err)
	} else if v != nil {
		m["guestOsFeatures"] = v
	}
	if v := f.LastAttachTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["lastAttachTimestamp"] = v
	}
	if v := f.LastDetachTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["lastDetachTimestamp"] = v
	}
	if v := f.Users; !dcl.IsEmptyValueIndirect(v) {
		m["users"] = v
	}
	if v := f.LicenseCodes; !dcl.IsEmptyValueIndirect(v) {
		m["licenseCodes"] = v
	}
	if v := f.PhysicalBlockSizeBytes; !dcl.IsEmptyValueIndirect(v) {
		m["physicalBlockSizeBytes"] = v
	}
	if v := f.ResourcePolicies; !dcl.IsEmptyValueIndirect(v) {
		m["resourcePolicies"] = v
	}
	if v := f.SourceDisk; !dcl.IsEmptyValueIndirect(v) {
		m["sourceDisk"] = v
	}
	if v := f.SourceDiskId; !dcl.IsEmptyValueIndirect(v) {
		m["sourceDiskId"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenDisk flattens Disk from a JSON request object into the
// Disk type.
func flattenDisk(c *Client, i interface{}) *Disk {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Disk{}
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Description = dcl.FlattenString(m["description"])
	res.DiskEncryptionKey = flattenDiskEncryptionKey(c, m["diskEncryptionKey"])
	res.GuestOsFeature = flattenDiskGuestOsFeatureSlice(c, m["guestOsFeature"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.LabelFingerprint = dcl.FlattenString(m["labelFingerprint"])
	res.License = dcl.FlattenStringSlice(m["license"])
	res.Name = dcl.FlattenString(m["name"])
	res.Region = dcl.FlattenString(m["region"])
	res.ReplicaZones = dcl.FlattenStringSlice(m["replicaZones"])
	res.ResourcePolicy = dcl.FlattenStringSlice(m["resourcePolicy"])
	res.SizeGb = dcl.FlattenInteger(m["sizeGb"])
	res.SourceImage = dcl.FlattenString(m["sourceImage"])
	res.SourceImageEncryptionKey = flattenDiskEncryptionKey(c, m["sourceImageEncryptionKey"])
	res.SourceImageId = dcl.FlattenString(m["sourceImageId"])
	res.SourceSnapshot = dcl.FlattenString(m["sourceSnapshot"])
	res.SourceSnapshotEncryptionKey = flattenDiskEncryptionKey(c, m["sourceSnapshotEncryptionKey"])
	res.SourceSnapshotId = dcl.FlattenString(m["sourceSnapshotId"])
	res.Type = dcl.FlattenString(m["type"])
	res.Zone = dcl.FlattenString(m["zone"])
	res.Project = dcl.FlattenString(m["project"])
	res.Id = dcl.FlattenInteger(m["id"])
	res.Status = flattenDiskStatusEnum(m["status"])
	res.Options = dcl.FlattenString(m["options"])
	res.Licenses = dcl.FlattenStringSlice(m["licenses"])
	res.GuestOsFeatures = flattenDiskGuestOsFeaturesSlice(c, m["guestOsFeatures"])
	res.LastAttachTimestamp = dcl.FlattenString(m["lastAttachTimestamp"])
	res.LastDetachTimestamp = dcl.FlattenString(m["lastDetachTimestamp"])
	res.Users = dcl.FlattenStringSlice(m["users"])
	res.LicenseCodes = dcl.FlattenIntSlice(m["licenseCodes"])
	res.PhysicalBlockSizeBytes = dcl.FlattenInteger(m["physicalBlockSizeBytes"])
	res.ResourcePolicies = dcl.FlattenStringSlice(m["resourcePolicies"])
	res.SourceDisk = dcl.FlattenString(m["sourceDisk"])
	res.SourceDiskId = dcl.FlattenString(m["sourceDiskId"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandDiskGuestOsFeatureMap expands the contents of DiskGuestOsFeature into a JSON
// request object.
func expandDiskGuestOsFeatureMap(c *Client, f map[string]DiskGuestOsFeature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDiskGuestOsFeature(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDiskGuestOsFeatureSlice expands the contents of DiskGuestOsFeature into a JSON
// request object.
func expandDiskGuestOsFeatureSlice(c *Client, f []DiskGuestOsFeature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDiskGuestOsFeature(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDiskGuestOsFeatureMap flattens the contents of DiskGuestOsFeature from a JSON
// response object.
func flattenDiskGuestOsFeatureMap(c *Client, i interface{}) map[string]DiskGuestOsFeature {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DiskGuestOsFeature{}
	}

	if len(a) == 0 {
		return map[string]DiskGuestOsFeature{}
	}

	items := make(map[string]DiskGuestOsFeature)
	for k, item := range a {
		items[k] = *flattenDiskGuestOsFeature(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDiskGuestOsFeatureSlice flattens the contents of DiskGuestOsFeature from a JSON
// response object.
func flattenDiskGuestOsFeatureSlice(c *Client, i interface{}) []DiskGuestOsFeature {
	a, ok := i.([]interface{})
	if !ok {
		return []DiskGuestOsFeature{}
	}

	if len(a) == 0 {
		return []DiskGuestOsFeature{}
	}

	items := make([]DiskGuestOsFeature, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDiskGuestOsFeature(c, item.(map[string]interface{})))
	}

	return items
}

// expandDiskGuestOsFeature expands an instance of DiskGuestOsFeature into a JSON
// request object.
func expandDiskGuestOsFeature(c *Client, f *DiskGuestOsFeature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.TypeAlt; !dcl.IsEmptyValueIndirect(v) {
		m["typeAlt"] = v
	}

	return m, nil
}

// flattenDiskGuestOsFeature flattens an instance of DiskGuestOsFeature from a JSON
// response object.
func flattenDiskGuestOsFeature(c *Client, i interface{}) *DiskGuestOsFeature {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DiskGuestOsFeature{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDiskGuestOsFeature
	}
	r.Type = flattenDiskGuestOsFeatureTypeEnum(m["type"])
	r.TypeAlt = flattenDiskGuestOsFeatureTypeAltEnumSlice(c, m["typeAlt"])

	return r
}

// expandDiskEncryptionKeyMap expands the contents of DiskEncryptionKey into a JSON
// request object.
func expandDiskEncryptionKeyMap(c *Client, f map[string]DiskEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDiskEncryptionKeySlice expands the contents of DiskEncryptionKey into a JSON
// request object.
func expandDiskEncryptionKeySlice(c *Client, f []DiskEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDiskEncryptionKeyMap flattens the contents of DiskEncryptionKey from a JSON
// response object.
func flattenDiskEncryptionKeyMap(c *Client, i interface{}) map[string]DiskEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DiskEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]DiskEncryptionKey{}
	}

	items := make(map[string]DiskEncryptionKey)
	for k, item := range a {
		items[k] = *flattenDiskEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDiskEncryptionKeySlice flattens the contents of DiskEncryptionKey from a JSON
// response object.
func flattenDiskEncryptionKeySlice(c *Client, i interface{}) []DiskEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []DiskEncryptionKey{}
	}

	if len(a) == 0 {
		return []DiskEncryptionKey{}
	}

	items := make([]DiskEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDiskEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandDiskEncryptionKey expands an instance of DiskEncryptionKey into a JSON
// request object.
func expandDiskEncryptionKey(c *Client, f *DiskEncryptionKey) (map[string]interface{}, error) {
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

// flattenDiskEncryptionKey flattens an instance of DiskEncryptionKey from a JSON
// response object.
func flattenDiskEncryptionKey(c *Client, i interface{}) *DiskEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DiskEncryptionKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDiskEncryptionKey
	}
	r.RawKey = dcl.FlattenSecretValue(m["rawKey"])
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	r.Sha256 = dcl.FlattenString(m["sha256"])
	r.KmsKeyServiceAccount = dcl.FlattenString(m["kmsKeyServiceAccount"])

	return r
}

// expandDiskGuestOsFeaturesMap expands the contents of DiskGuestOsFeatures into a JSON
// request object.
func expandDiskGuestOsFeaturesMap(c *Client, f map[string]DiskGuestOsFeatures) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDiskGuestOsFeatures(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDiskGuestOsFeaturesSlice expands the contents of DiskGuestOsFeatures into a JSON
// request object.
func expandDiskGuestOsFeaturesSlice(c *Client, f []DiskGuestOsFeatures) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDiskGuestOsFeatures(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDiskGuestOsFeaturesMap flattens the contents of DiskGuestOsFeatures from a JSON
// response object.
func flattenDiskGuestOsFeaturesMap(c *Client, i interface{}) map[string]DiskGuestOsFeatures {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DiskGuestOsFeatures{}
	}

	if len(a) == 0 {
		return map[string]DiskGuestOsFeatures{}
	}

	items := make(map[string]DiskGuestOsFeatures)
	for k, item := range a {
		items[k] = *flattenDiskGuestOsFeatures(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDiskGuestOsFeaturesSlice flattens the contents of DiskGuestOsFeatures from a JSON
// response object.
func flattenDiskGuestOsFeaturesSlice(c *Client, i interface{}) []DiskGuestOsFeatures {
	a, ok := i.([]interface{})
	if !ok {
		return []DiskGuestOsFeatures{}
	}

	if len(a) == 0 {
		return []DiskGuestOsFeatures{}
	}

	items := make([]DiskGuestOsFeatures, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDiskGuestOsFeatures(c, item.(map[string]interface{})))
	}

	return items
}

// expandDiskGuestOsFeatures expands an instance of DiskGuestOsFeatures into a JSON
// request object.
func expandDiskGuestOsFeatures(c *Client, f *DiskGuestOsFeatures) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.TypeAlts; !dcl.IsEmptyValueIndirect(v) {
		m["typeAlts"] = v
	}

	return m, nil
}

// flattenDiskGuestOsFeatures flattens an instance of DiskGuestOsFeatures from a JSON
// response object.
func flattenDiskGuestOsFeatures(c *Client, i interface{}) *DiskGuestOsFeatures {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DiskGuestOsFeatures{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDiskGuestOsFeatures
	}
	r.Type = flattenDiskGuestOsFeaturesTypeEnum(m["type"])
	r.TypeAlts = flattenDiskGuestOsFeaturesTypeAltsEnumSlice(c, m["typeAlts"])

	return r
}

// flattenDiskGuestOsFeatureTypeEnumSlice flattens the contents of DiskGuestOsFeatureTypeEnum from a JSON
// response object.
func flattenDiskGuestOsFeatureTypeEnumSlice(c *Client, i interface{}) []DiskGuestOsFeatureTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DiskGuestOsFeatureTypeEnum{}
	}

	if len(a) == 0 {
		return []DiskGuestOsFeatureTypeEnum{}
	}

	items := make([]DiskGuestOsFeatureTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDiskGuestOsFeatureTypeEnum(item.(interface{})))
	}

	return items
}

// flattenDiskGuestOsFeatureTypeEnum asserts that an interface is a string, and returns a
// pointer to a *DiskGuestOsFeatureTypeEnum with the same value as that string.
func flattenDiskGuestOsFeatureTypeEnum(i interface{}) *DiskGuestOsFeatureTypeEnum {
	s, ok := i.(string)
	if !ok {
		return DiskGuestOsFeatureTypeEnumRef("")
	}

	return DiskGuestOsFeatureTypeEnumRef(s)
}

// flattenDiskGuestOsFeatureTypeAltEnumSlice flattens the contents of DiskGuestOsFeatureTypeAltEnum from a JSON
// response object.
func flattenDiskGuestOsFeatureTypeAltEnumSlice(c *Client, i interface{}) []DiskGuestOsFeatureTypeAltEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DiskGuestOsFeatureTypeAltEnum{}
	}

	if len(a) == 0 {
		return []DiskGuestOsFeatureTypeAltEnum{}
	}

	items := make([]DiskGuestOsFeatureTypeAltEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDiskGuestOsFeatureTypeAltEnum(item.(interface{})))
	}

	return items
}

// flattenDiskGuestOsFeatureTypeAltEnum asserts that an interface is a string, and returns a
// pointer to a *DiskGuestOsFeatureTypeAltEnum with the same value as that string.
func flattenDiskGuestOsFeatureTypeAltEnum(i interface{}) *DiskGuestOsFeatureTypeAltEnum {
	s, ok := i.(string)
	if !ok {
		return DiskGuestOsFeatureTypeAltEnumRef("")
	}

	return DiskGuestOsFeatureTypeAltEnumRef(s)
}

// flattenDiskStatusEnumSlice flattens the contents of DiskStatusEnum from a JSON
// response object.
func flattenDiskStatusEnumSlice(c *Client, i interface{}) []DiskStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DiskStatusEnum{}
	}

	if len(a) == 0 {
		return []DiskStatusEnum{}
	}

	items := make([]DiskStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDiskStatusEnum(item.(interface{})))
	}

	return items
}

// flattenDiskStatusEnum asserts that an interface is a string, and returns a
// pointer to a *DiskStatusEnum with the same value as that string.
func flattenDiskStatusEnum(i interface{}) *DiskStatusEnum {
	s, ok := i.(string)
	if !ok {
		return DiskStatusEnumRef("")
	}

	return DiskStatusEnumRef(s)
}

// flattenDiskGuestOsFeaturesTypeEnumSlice flattens the contents of DiskGuestOsFeaturesTypeEnum from a JSON
// response object.
func flattenDiskGuestOsFeaturesTypeEnumSlice(c *Client, i interface{}) []DiskGuestOsFeaturesTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DiskGuestOsFeaturesTypeEnum{}
	}

	if len(a) == 0 {
		return []DiskGuestOsFeaturesTypeEnum{}
	}

	items := make([]DiskGuestOsFeaturesTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDiskGuestOsFeaturesTypeEnum(item.(interface{})))
	}

	return items
}

// flattenDiskGuestOsFeaturesTypeEnum asserts that an interface is a string, and returns a
// pointer to a *DiskGuestOsFeaturesTypeEnum with the same value as that string.
func flattenDiskGuestOsFeaturesTypeEnum(i interface{}) *DiskGuestOsFeaturesTypeEnum {
	s, ok := i.(string)
	if !ok {
		return DiskGuestOsFeaturesTypeEnumRef("")
	}

	return DiskGuestOsFeaturesTypeEnumRef(s)
}

// flattenDiskGuestOsFeaturesTypeAltsEnumSlice flattens the contents of DiskGuestOsFeaturesTypeAltsEnum from a JSON
// response object.
func flattenDiskGuestOsFeaturesTypeAltsEnumSlice(c *Client, i interface{}) []DiskGuestOsFeaturesTypeAltsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DiskGuestOsFeaturesTypeAltsEnum{}
	}

	if len(a) == 0 {
		return []DiskGuestOsFeaturesTypeAltsEnum{}
	}

	items := make([]DiskGuestOsFeaturesTypeAltsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDiskGuestOsFeaturesTypeAltsEnum(item.(interface{})))
	}

	return items
}

// flattenDiskGuestOsFeaturesTypeAltsEnum asserts that an interface is a string, and returns a
// pointer to a *DiskGuestOsFeaturesTypeAltsEnum with the same value as that string.
func flattenDiskGuestOsFeaturesTypeAltsEnum(i interface{}) *DiskGuestOsFeaturesTypeAltsEnum {
	s, ok := i.(string)
	if !ok {
		return DiskGuestOsFeaturesTypeAltsEnumRef("")
	}

	return DiskGuestOsFeaturesTypeAltsEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Disk) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalDisk(b, c)
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
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
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

type diskDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         diskApiOperation
}

func convertFieldDiffToDiskOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]diskDiff, error) {
	var diffs []diskDiff
	for _, op := range ops {
		diff := diskDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTodiskApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTodiskApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (diskApiOperation, error) {
	switch op {

	case "updateDiskResizeOperation":
		return &updateDiskResizeOperation{Diffs: diffs}, nil

	case "updateDiskSetLabelsOperation":
		return &updateDiskSetLabelsOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
