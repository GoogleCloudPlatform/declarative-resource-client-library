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
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Address) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func addressGetURL(userBasePath string, r *Address) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/addresses/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/addresses/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func addressListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/addresses", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/addresses", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func addressCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/addresses", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/addresses", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func addressDeleteURL(userBasePath string, r *Address) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/addresses/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/addresses/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// addressApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type addressApiOperation interface {
	do(context.Context, *Address, *Client) error
}

// newUpdateAddressSetLabelsRequest creates a request for an
// Address resource's setLabels update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAddressSetLabelsRequest(ctx context.Context, f *Address, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.LabelFingerprint; !dcl.IsEmptyValueIndirect(v) {
		req["labelFingerprint"] = v
	}
	return req, nil
}

// marshalUpdateAddressSetLabelsRequest converts the update into
// the final JSON request body.
func marshalUpdateAddressSetLabelsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAddressSetLabelsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAddressSetLabelsOperation) do(ctx context.Context, r *Address, c *Client) error {
	_, err := c.GetAddress(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setLabels")
	if err != nil {
		return err
	}

	req, err := newUpdateAddressSetLabelsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateAddressSetLabelsRequest(c, req)
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

func (c *Client) listAddressRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := addressListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AddressMaxPage {
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

type listAddressOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listAddress(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Address, string, error) {
	b, err := c.listAddressRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAddressOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Address
	for _, v := range m.Items {
		res := flattenAddress(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAddress(ctx context.Context, f func(*Address) bool, resources []*Address) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAddress(ctx, res)
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

type deleteAddressOperation struct{}

func (op *deleteAddressOperation) do(ctx context.Context, r *Address, c *Client) error {

	_, err := c.GetAddress(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Address not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAddress checking for existence. error: %v", err)
		return err
	}

	u, err := addressDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetAddress(ctx, r.urlNormalized())
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
type createAddressOperation struct {
	response map[string]interface{}
}

func (op *createAddressOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAddressOperation) do(ctx context.Context, r *Address, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := addressCreateURL(c.Config.BasePath, project, location)

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

	if _, err := c.GetAddress(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAddressRaw(ctx context.Context, r *Address) ([]byte, error) {
	if dcl.IsZeroValue(r.AddressType) {
		r.AddressType = AddressAddressTypeEnumRef("EXTERNAL")
	}

	u, err := addressGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) addressDiffsForRawDesired(ctx context.Context, rawDesired *Address, opts ...dcl.ApplyOption) (initial, desired *Address, diffs []addressDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Address
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Address); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Address, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAddress(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Address resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Address resource: %v", err)
		}
		c.Config.Logger.Info("Found that Address resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAddressDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Address: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Address: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAddressInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Address: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAddressDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Address: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAddress(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAddressInitialState(rawInitial, rawDesired *Address) (*Address, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAddressDesiredState(rawDesired, rawInitial *Address, opts ...dcl.ApplyOption) (*Address, error) {

	if dcl.IsZeroValue(rawDesired.AddressType) {
		rawDesired.AddressType = AddressAddressTypeEnumRef("EXTERNAL")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.Address, rawInitial.Address) {
		rawDesired.Address = rawInitial.Address
	}
	if dcl.IsZeroValue(rawDesired.PrefixLength) {
		rawDesired.PrefixLength = rawInitial.PrefixLength
	}
	if dcl.IsZeroValue(rawDesired.Status) {
		rawDesired.Status = rawInitial.Status
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.IsZeroValue(rawDesired.NetworkTier) {
		rawDesired.NetworkTier = rawInitial.NetworkTier
	}
	if dcl.IsZeroValue(rawDesired.IPVersion) {
		rawDesired.IPVersion = rawInitial.IPVersion
	}
	if dcl.IsZeroValue(rawDesired.AddressType) {
		rawDesired.AddressType = rawInitial.AddressType
	}
	if dcl.IsZeroValue(rawDesired.Purpose) {
		rawDesired.Purpose = rawInitial.Purpose
	}
	if dcl.NameToSelfLink(rawDesired.Subnetwork, rawInitial.Subnetwork) {
		rawDesired.Subnetwork = rawInitial.Subnetwork
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.IsZeroValue(rawDesired.CreationTimestamp) {
		rawDesired.CreationTimestamp = rawInitial.CreationTimestamp
	}
	if dcl.IsZeroValue(rawDesired.Users) {
		rawDesired.Users = rawInitial.Users
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.StringCanonicalize(rawDesired.LabelFingerprint, rawInitial.LabelFingerprint) {
		rawDesired.LabelFingerprint = rawInitial.LabelFingerprint
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeAddressNewState(c *Client, rawNew, rawDesired *Address) (*Address, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

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

	if dcl.IsEmptyValueIndirect(rawNew.Address) && dcl.IsEmptyValueIndirect(rawDesired.Address) {
		rawNew.Address = rawDesired.Address
	} else {
		if dcl.StringCanonicalize(rawDesired.Address, rawNew.Address) {
			rawNew.Address = rawDesired.Address
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PrefixLength) && dcl.IsEmptyValueIndirect(rawDesired.PrefixLength) {
		rawNew.PrefixLength = rawDesired.PrefixLength
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NetworkTier) && dcl.IsEmptyValueIndirect(rawDesired.NetworkTier) {
		rawNew.NetworkTier = rawDesired.NetworkTier
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPVersion) && dcl.IsEmptyValueIndirect(rawDesired.IPVersion) {
		rawNew.IPVersion = rawDesired.IPVersion
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AddressType) && dcl.IsEmptyValueIndirect(rawDesired.AddressType) {
		rawNew.AddressType = rawDesired.AddressType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Purpose) && dcl.IsEmptyValueIndirect(rawDesired.Purpose) {
		rawNew.Purpose = rawDesired.Purpose
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Subnetwork) && dcl.IsEmptyValueIndirect(rawDesired.Subnetwork) {
		rawNew.Subnetwork = rawDesired.Subnetwork
	} else {
		if dcl.NameToSelfLink(rawDesired.Subnetwork, rawNew.Subnetwork) {
			rawNew.Subnetwork = rawDesired.Subnetwork
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Users) && dcl.IsEmptyValueIndirect(rawDesired.Users) {
		rawNew.Users = rawDesired.Users
	} else {
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

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

type addressDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         addressApiOperation
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
func diffAddress(c *Client, desired, actual *Address, opts ...dcl.ApplyOption) ([]addressDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []addressDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.Id, actual.Id, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "Id"})
	}

	if d, err := dcl.Diff(desired.Name, actual.Name, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "Name"})
	}

	if d, err := dcl.Diff(desired.Description, actual.Description, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "Description"})
	}

	if d, err := dcl.Diff(desired.Address, actual.Address, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "Address"})
	}

	if d, err := dcl.Diff(desired.PrefixLength, actual.PrefixLength, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "PrefixLength"})
	}

	if d, err := dcl.Diff(desired.Region, actual.Region, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "Region"})
	}

	if d, err := dcl.Diff(desired.SelfLink, actual.SelfLink, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "SelfLink"})
	}

	if d, err := dcl.Diff(desired.Subnetwork, actual.Subnetwork, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "ReferenceType"}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "Subnetwork"})
	}

	if d, err := dcl.Diff(desired.Network, actual.Network, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "ReferenceType"}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "Network"})
	}

	if d, err := dcl.Diff(desired.CreationTimestamp, actual.CreationTimestamp, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{RequiresRecreate: true, FieldName: "CreationTimestamp"})
	}

	if d, err := dcl.Diff(desired.Labels, actual.Labels, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{
			UpdateOp: &updateAddressSetLabelsOperation{}, FieldName: "Labels",
		})
	}

	if d, err := dcl.Diff(desired.LabelFingerprint, actual.LabelFingerprint, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, addressDiff{
			UpdateOp: &updateAddressSetLabelsOperation{}, FieldName: "LabelFingerprint",
		})
	}

	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.Address) && !dcl.StringCanonicalize(desired.Address, actual.Address) {
		c.Config.Logger.Infof("Detected diff in Address.\nDESIRED: %v\nACTUAL: %v", desired.Address, actual.Address)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "Address",
		})
	}
	if !reflect.DeepEqual(desired.PrefixLength, actual.PrefixLength) {
		c.Config.Logger.Infof("Detected diff in PrefixLength.\nDESIRED: %v\nACTUAL: %v", desired.PrefixLength, actual.PrefixLength)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "PrefixLength",
		})
	}
	if !dcl.IsZeroValue(desired.Region) && !dcl.StringCanonicalize(desired.Region, actual.Region) {
		c.Config.Logger.Infof("Detected diff in Region.\nDESIRED: %v\nACTUAL: %v", desired.Region, actual.Region)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "Region",
		})
	}
	if !reflect.DeepEqual(desired.NetworkTier, actual.NetworkTier) {
		c.Config.Logger.Infof("Detected diff in NetworkTier.\nDESIRED: %v\nACTUAL: %v", desired.NetworkTier, actual.NetworkTier)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "NetworkTier",
		})
	}
	if !reflect.DeepEqual(desired.IPVersion, actual.IPVersion) {
		c.Config.Logger.Infof("Detected diff in IPVersion.\nDESIRED: %v\nACTUAL: %v", desired.IPVersion, actual.IPVersion)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "IPVersion",
		})
	}
	if !reflect.DeepEqual(desired.AddressType, actual.AddressType) {
		c.Config.Logger.Infof("Detected diff in AddressType.\nDESIRED: %v\nACTUAL: %v", desired.AddressType, actual.AddressType)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "AddressType",
		})
	}
	if !reflect.DeepEqual(desired.Purpose, actual.Purpose) {
		c.Config.Logger.Infof("Detected diff in Purpose.\nDESIRED: %v\nACTUAL: %v", desired.Purpose, actual.Purpose)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "Purpose",
		})
	}
	if !dcl.IsZeroValue(desired.Subnetwork) && !dcl.NameToSelfLink(desired.Subnetwork, actual.Subnetwork) {
		c.Config.Logger.Infof("Detected diff in Subnetwork.\nDESIRED: %v\nACTUAL: %v", desired.Subnetwork, actual.Subnetwork)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "Subnetwork",
		})
	}
	if !dcl.IsZeroValue(desired.Network) && !dcl.PartialSelfLinkToSelfLink(desired.Network, actual.Network) {
		c.Config.Logger.Infof("Detected diff in Network.\nDESIRED: %v\nACTUAL: %v", desired.Network, actual.Network)
		diffs = append(diffs, addressDiff{
			RequiresRecreate: true,
			FieldName:        "Network",
		})
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, addressDiff{
			UpdateOp:  &updateAddressSetLabelsOperation{},
			FieldName: "Labels",
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
	var deduped []addressDiff
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
func compareAddressStatusEnumSlice(c *Client, desired, actual []AddressStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AddressStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAddressStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AddressStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAddressStatusEnum(c *Client, desired, actual *AddressStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAddressNetworkTierEnumSlice(c *Client, desired, actual []AddressNetworkTierEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AddressNetworkTierEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAddressNetworkTierEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AddressNetworkTierEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAddressNetworkTierEnum(c *Client, desired, actual *AddressNetworkTierEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAddressIPVersionEnumSlice(c *Client, desired, actual []AddressIPVersionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AddressIPVersionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAddressIPVersionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AddressIPVersionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAddressIPVersionEnum(c *Client, desired, actual *AddressIPVersionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAddressAddressTypeEnumSlice(c *Client, desired, actual []AddressAddressTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AddressAddressTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAddressAddressTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AddressAddressTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAddressAddressTypeEnum(c *Client, desired, actual *AddressAddressTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAddressPurposeEnumSlice(c *Client, desired, actual []AddressPurposeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AddressPurposeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAddressPurposeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AddressPurposeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAddressPurposeEnum(c *Client, desired, actual *AddressPurposeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Address) urlNormalized() *Address {
	normalized := deepcopy.Copy(*r).(Address)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Address = dcl.SelfLinkToName(r.Address)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Subnetwork = dcl.SelfLinkToName(r.Subnetwork)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.LabelFingerprint = dcl.SelfLinkToName(r.LabelFingerprint)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Address) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Address) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *Address) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Address) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "setLabels" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		if dcl.IsRegion(r.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/addresses/{{name}}/setLabels", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil
		}

		return dcl.URL("projects/{{project}}/global/addresses/{{name}}/setLabels", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Address resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Address) marshal(c *Client) ([]byte, error) {
	m, err := expandAddress(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Address: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAddress decodes JSON responses into the Address resource schema.
func unmarshalAddress(b []byte, c *Client) (*Address, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAddress(m, c)
}

func unmarshalMapAddress(m map[string]interface{}, c *Client) (*Address, error) {

	return flattenAddress(c, m), nil
}

// expandAddress expands Address into a JSON request object.
func expandAddress(c *Client, f *Address) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Address; !dcl.IsEmptyValueIndirect(v) {
		m["address"] = v
	}
	if v := f.PrefixLength; !dcl.IsEmptyValueIndirect(v) {
		m["prefixLength"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.NetworkTier; !dcl.IsEmptyValueIndirect(v) {
		m["networkTier"] = v
	}
	if v := f.IPVersion; !dcl.IsEmptyValueIndirect(v) {
		m["ipVersion"] = v
	}
	if v := f.AddressType; !dcl.IsEmptyValueIndirect(v) {
		m["addressType"] = v
	}
	if v := f.Purpose; !dcl.IsEmptyValueIndirect(v) {
		m["purpose"] = v
	}
	if v := f.Subnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["subnetwork"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/global/networks/%s", f.Network, f.Project, f.Network); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v := f.Users; !dcl.IsEmptyValueIndirect(v) {
		m["users"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.LabelFingerprint; !dcl.IsEmptyValueIndirect(v) {
		m["labelFingerprint"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenAddress flattens Address from a JSON request object into the
// Address type.
func flattenAddress(c *Client, i interface{}) *Address {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Address{}
	r.Id = dcl.FlattenInteger(m["id"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.Address = dcl.FlattenString(m["address"])
	r.PrefixLength = dcl.FlattenInteger(m["prefixLength"])
	r.Status = flattenAddressStatusEnum(m["status"])
	r.Region = dcl.FlattenString(m["region"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.NetworkTier = flattenAddressNetworkTierEnum(m["networkTier"])
	r.IPVersion = flattenAddressIPVersionEnum(m["ipVersion"])
	r.AddressType = flattenAddressAddressTypeEnum(m["addressType"])
	if _, ok := m["addressType"]; !ok {
		c.Config.Logger.Info("Using default value for addressType")
		r.AddressType = AddressAddressTypeEnumRef("EXTERNAL")
	}
	r.Purpose = flattenAddressPurposeEnum(m["purpose"])
	r.Subnetwork = dcl.FlattenString(m["subnetwork"])
	r.Network = dcl.FlattenString(m["network"])
	r.Project = dcl.FlattenString(m["project"])
	r.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	r.Users = dcl.FlattenStringSlice(m["users"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.LabelFingerprint = dcl.FlattenString(m["labelFingerprint"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// flattenAddressStatusEnumSlice flattens the contents of AddressStatusEnum from a JSON
// response object.
func flattenAddressStatusEnumSlice(c *Client, i interface{}) []AddressStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AddressStatusEnum{}
	}

	if len(a) == 0 {
		return []AddressStatusEnum{}
	}

	items := make([]AddressStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAddressStatusEnum(item.(interface{})))
	}

	return items
}

// flattenAddressStatusEnum asserts that an interface is a string, and returns a
// pointer to a *AddressStatusEnum with the same value as that string.
func flattenAddressStatusEnum(i interface{}) *AddressStatusEnum {
	s, ok := i.(string)
	if !ok {
		return AddressStatusEnumRef("")
	}

	return AddressStatusEnumRef(s)
}

// flattenAddressNetworkTierEnumSlice flattens the contents of AddressNetworkTierEnum from a JSON
// response object.
func flattenAddressNetworkTierEnumSlice(c *Client, i interface{}) []AddressNetworkTierEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AddressNetworkTierEnum{}
	}

	if len(a) == 0 {
		return []AddressNetworkTierEnum{}
	}

	items := make([]AddressNetworkTierEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAddressNetworkTierEnum(item.(interface{})))
	}

	return items
}

// flattenAddressNetworkTierEnum asserts that an interface is a string, and returns a
// pointer to a *AddressNetworkTierEnum with the same value as that string.
func flattenAddressNetworkTierEnum(i interface{}) *AddressNetworkTierEnum {
	s, ok := i.(string)
	if !ok {
		return AddressNetworkTierEnumRef("")
	}

	return AddressNetworkTierEnumRef(s)
}

// flattenAddressIPVersionEnumSlice flattens the contents of AddressIPVersionEnum from a JSON
// response object.
func flattenAddressIPVersionEnumSlice(c *Client, i interface{}) []AddressIPVersionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AddressIPVersionEnum{}
	}

	if len(a) == 0 {
		return []AddressIPVersionEnum{}
	}

	items := make([]AddressIPVersionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAddressIPVersionEnum(item.(interface{})))
	}

	return items
}

// flattenAddressIPVersionEnum asserts that an interface is a string, and returns a
// pointer to a *AddressIPVersionEnum with the same value as that string.
func flattenAddressIPVersionEnum(i interface{}) *AddressIPVersionEnum {
	s, ok := i.(string)
	if !ok {
		return AddressIPVersionEnumRef("")
	}

	return AddressIPVersionEnumRef(s)
}

// flattenAddressAddressTypeEnumSlice flattens the contents of AddressAddressTypeEnum from a JSON
// response object.
func flattenAddressAddressTypeEnumSlice(c *Client, i interface{}) []AddressAddressTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AddressAddressTypeEnum{}
	}

	if len(a) == 0 {
		return []AddressAddressTypeEnum{}
	}

	items := make([]AddressAddressTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAddressAddressTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAddressAddressTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AddressAddressTypeEnum with the same value as that string.
func flattenAddressAddressTypeEnum(i interface{}) *AddressAddressTypeEnum {
	s, ok := i.(string)
	if !ok {
		return AddressAddressTypeEnumRef("")
	}

	return AddressAddressTypeEnumRef(s)
}

// flattenAddressPurposeEnumSlice flattens the contents of AddressPurposeEnum from a JSON
// response object.
func flattenAddressPurposeEnumSlice(c *Client, i interface{}) []AddressPurposeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AddressPurposeEnum{}
	}

	if len(a) == 0 {
		return []AddressPurposeEnum{}
	}

	items := make([]AddressPurposeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAddressPurposeEnum(item.(interface{})))
	}

	return items
}

// flattenAddressPurposeEnum asserts that an interface is a string, and returns a
// pointer to a *AddressPurposeEnum with the same value as that string.
func flattenAddressPurposeEnum(i interface{}) *AddressPurposeEnum {
	s, ok := i.(string)
	if !ok {
		return AddressPurposeEnumRef("")
	}

	return AddressPurposeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Address) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAddress(b, c)
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
