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

func (r *Reservation) validate() error {

	if err := dcl.Required(r, "zone"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SpecificReservation) {
		if err := r.SpecificReservation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ReservationSpecificReservation) validate() error {
	if !dcl.IsEmptyValueIndirect(r.InstanceProperties) {
		if err := r.InstanceProperties.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ReservationSpecificReservationInstanceProperties) validate() error {
	return nil
}
func (r *ReservationSpecificReservationInstancePropertiesGuestAccelerators) validate() error {
	return nil
}
func (r *ReservationSpecificReservationInstancePropertiesLocalSsds) validate() error {
	return nil
}

func reservationGetURL(userBasePath string, r *Reservation) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"zone":    dcl.ValueOrEmptyString(r.Zone),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/reservations/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func reservationListURL(userBasePath, project, zone string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"zone":    zone,
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/reservations", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func reservationCreateURL(userBasePath, project, zone string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"zone":    zone,
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/reservations", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func reservationDeleteURL(userBasePath string, r *Reservation) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"zone":    dcl.ValueOrEmptyString(r.Zone),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/reservations/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// reservationApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type reservationApiOperation interface {
	do(context.Context, *Reservation, *Client) error
}

func (c *Client) listReservationRaw(ctx context.Context, project, zone, pageToken string, pageSize int32) ([]byte, error) {
	u, err := reservationListURL(c.Config.BasePath, project, zone)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ReservationMaxPage {
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

type listReservationOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listReservation(ctx context.Context, project, zone, pageToken string, pageSize int32) ([]*Reservation, string, error) {
	b, err := c.listReservationRaw(ctx, project, zone, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listReservationOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Reservation
	for _, v := range m.Items {
		res, err := unmarshalMapReservation(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Zone = &zone
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllReservation(ctx context.Context, f func(*Reservation) bool, resources []*Reservation) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteReservation(ctx, res)
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

type deleteReservationOperation struct{}

func (op *deleteReservationOperation) do(ctx context.Context, r *Reservation, c *Client) error {

	_, err := c.GetReservation(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Reservation not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetReservation checking for existence. error: %v", err)
		return err
	}

	u, err := reservationDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetReservation(ctx, r.urlNormalized())
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
type createReservationOperation struct {
	response map[string]interface{}
}

func (op *createReservationOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createReservationOperation) do(ctx context.Context, r *Reservation, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, zone := r.createFields()
	u, err := reservationCreateURL(c.Config.BasePath, project, zone)

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

	if _, err := c.GetReservation(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getReservationRaw(ctx context.Context, r *Reservation) ([]byte, error) {

	u, err := reservationGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) reservationDiffsForRawDesired(ctx context.Context, rawDesired *Reservation, opts ...dcl.ApplyOption) (initial, desired *Reservation, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Reservation
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Reservation); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Reservation, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetReservation(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Reservation resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Reservation resource: %v", err)
		}
		c.Config.Logger.Info("Found that Reservation resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeReservationDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Reservation: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Reservation: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeReservationInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Reservation: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeReservationDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Reservation: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffReservation(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeReservationInitialState(rawInitial, rawDesired *Reservation) (*Reservation, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeReservationDesiredState(rawDesired, rawInitial *Reservation, opts ...dcl.ApplyOption) (*Reservation, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SpecificReservation = canonicalizeReservationSpecificReservation(rawDesired.SpecificReservation, nil, opts...)

		return rawDesired, nil
	}

	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.Zone, rawInitial.Zone) {
		rawDesired.Zone = rawInitial.Zone
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.SpecificReservation = canonicalizeReservationSpecificReservation(rawDesired.SpecificReservation, rawInitial.SpecificReservation, opts...)
	if dcl.StringCanonicalize(rawDesired.Commitment, rawInitial.Commitment) {
		rawDesired.Commitment = rawInitial.Commitment
	}
	if dcl.BoolCanonicalize(rawDesired.SpecificReservationRequired, rawInitial.SpecificReservationRequired) {
		rawDesired.SpecificReservationRequired = rawInitial.SpecificReservationRequired
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeReservationNewState(c *Client, rawNew, rawDesired *Reservation) (*Reservation, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Zone) && dcl.IsEmptyValueIndirect(rawDesired.Zone) {
		rawNew.Zone = rawDesired.Zone
	} else {
		if dcl.StringCanonicalize(rawDesired.Zone, rawNew.Zone) {
			rawNew.Zone = rawDesired.Zone
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SpecificReservation) && dcl.IsEmptyValueIndirect(rawDesired.SpecificReservation) {
		rawNew.SpecificReservation = rawDesired.SpecificReservation
	} else {
		rawNew.SpecificReservation = canonicalizeNewReservationSpecificReservation(c, rawDesired.SpecificReservation, rawNew.SpecificReservation)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Commitment) && dcl.IsEmptyValueIndirect(rawDesired.Commitment) {
		rawNew.Commitment = rawDesired.Commitment
	} else {
		if dcl.StringCanonicalize(rawDesired.Commitment, rawNew.Commitment) {
			rawNew.Commitment = rawDesired.Commitment
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SpecificReservationRequired) && dcl.IsEmptyValueIndirect(rawDesired.SpecificReservationRequired) {
		rawNew.SpecificReservationRequired = rawDesired.SpecificReservationRequired
	} else {
		if dcl.BoolCanonicalize(rawDesired.SpecificReservationRequired, rawNew.SpecificReservationRequired) {
			rawNew.SpecificReservationRequired = rawDesired.SpecificReservationRequired
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeReservationSpecificReservation(des, initial *ReservationSpecificReservation, opts ...dcl.ApplyOption) *ReservationSpecificReservation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.InstanceProperties = canonicalizeReservationSpecificReservationInstanceProperties(des.InstanceProperties, initial.InstanceProperties, opts...)
	if dcl.IsZeroValue(des.Count) {
		des.Count = initial.Count
	}
	if dcl.IsZeroValue(des.InUseCount) {
		des.InUseCount = initial.InUseCount
	}

	return des
}

func canonicalizeNewReservationSpecificReservation(c *Client, des, nw *ReservationSpecificReservation) *ReservationSpecificReservation {
	if des == nil || nw == nil {
		return nw
	}

	nw.InstanceProperties = canonicalizeNewReservationSpecificReservationInstanceProperties(c, des.InstanceProperties, nw.InstanceProperties)
	if dcl.IsZeroValue(nw.Count) {
		nw.Count = des.Count
	}
	if dcl.IsZeroValue(nw.InUseCount) {
		nw.InUseCount = des.InUseCount
	}

	return nw
}

func canonicalizeNewReservationSpecificReservationSet(c *Client, des, nw []ReservationSpecificReservation) []ReservationSpecificReservation {
	if des == nil {
		return nw
	}
	var reorderedNew []ReservationSpecificReservation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareReservationSpecificReservationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewReservationSpecificReservationSlice(c *Client, des, nw []ReservationSpecificReservation) []ReservationSpecificReservation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ReservationSpecificReservation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewReservationSpecificReservation(c, &d, &n))
	}

	return items
}

func canonicalizeReservationSpecificReservationInstanceProperties(des, initial *ReservationSpecificReservationInstanceProperties, opts ...dcl.ApplyOption) *ReservationSpecificReservationInstanceProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		des.MachineType = initial.MachineType
	}
	if dcl.IsZeroValue(des.GuestAccelerators) {
		des.GuestAccelerators = initial.GuestAccelerators
	}
	if canonicalizeReservationCPUPlatform(des.MinCpuPlatform, initial.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
		des.MinCpuPlatform = initial.MinCpuPlatform
	}
	if dcl.IsZeroValue(des.LocalSsds) {
		des.LocalSsds = initial.LocalSsds
	}

	return des
}

func canonicalizeNewReservationSpecificReservationInstanceProperties(c *Client, des, nw *ReservationSpecificReservationInstanceProperties) *ReservationSpecificReservationInstanceProperties {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	nw.GuestAccelerators = canonicalizeNewReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice(c, des.GuestAccelerators, nw.GuestAccelerators)
	if canonicalizeReservationCPUPlatform(des.MinCpuPlatform, nw.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}
	nw.LocalSsds = canonicalizeNewReservationSpecificReservationInstancePropertiesLocalSsdsSlice(c, des.LocalSsds, nw.LocalSsds)

	return nw
}

func canonicalizeNewReservationSpecificReservationInstancePropertiesSet(c *Client, des, nw []ReservationSpecificReservationInstanceProperties) []ReservationSpecificReservationInstanceProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []ReservationSpecificReservationInstanceProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareReservationSpecificReservationInstancePropertiesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewReservationSpecificReservationInstancePropertiesSlice(c *Client, des, nw []ReservationSpecificReservationInstanceProperties) []ReservationSpecificReservationInstanceProperties {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ReservationSpecificReservationInstanceProperties
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewReservationSpecificReservationInstanceProperties(c, &d, &n))
	}

	return items
}

func canonicalizeReservationSpecificReservationInstancePropertiesGuestAccelerators(des, initial *ReservationSpecificReservationInstancePropertiesGuestAccelerators, opts ...dcl.ApplyOption) *ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.AcceleratorType, initial.AcceleratorType) || dcl.IsZeroValue(des.AcceleratorType) {
		des.AcceleratorType = initial.AcceleratorType
	}
	if dcl.IsZeroValue(des.AcceleratorCount) {
		des.AcceleratorCount = initial.AcceleratorCount
	}

	return des
}

func canonicalizeNewReservationSpecificReservationInstancePropertiesGuestAccelerators(c *Client, des, nw *ReservationSpecificReservationInstancePropertiesGuestAccelerators) *ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AcceleratorType, nw.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
	}
	if dcl.IsZeroValue(nw.AcceleratorCount) {
		nw.AcceleratorCount = des.AcceleratorCount
	}

	return nw
}

func canonicalizeNewReservationSpecificReservationInstancePropertiesGuestAcceleratorsSet(c *Client, des, nw []ReservationSpecificReservationInstancePropertiesGuestAccelerators) []ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []ReservationSpecificReservationInstancePropertiesGuestAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareReservationSpecificReservationInstancePropertiesGuestAcceleratorsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice(c *Client, des, nw []ReservationSpecificReservationInstancePropertiesGuestAccelerators) []ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ReservationSpecificReservationInstancePropertiesGuestAccelerators
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewReservationSpecificReservationInstancePropertiesGuestAccelerators(c, &d, &n))
	}

	return items
}

func canonicalizeReservationSpecificReservationInstancePropertiesLocalSsds(des, initial *ReservationSpecificReservationInstancePropertiesLocalSsds, opts ...dcl.ApplyOption) *ReservationSpecificReservationInstancePropertiesLocalSsds {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DiskSizeGb) {
		des.DiskSizeGb = initial.DiskSizeGb
	}
	if dcl.IsZeroValue(des.Interface) {
		des.Interface = initial.Interface
	}

	return des
}

func canonicalizeNewReservationSpecificReservationInstancePropertiesLocalSsds(c *Client, des, nw *ReservationSpecificReservationInstancePropertiesLocalSsds) *ReservationSpecificReservationInstancePropertiesLocalSsds {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.DiskSizeGb) {
		nw.DiskSizeGb = des.DiskSizeGb
	}
	if dcl.IsZeroValue(nw.Interface) {
		nw.Interface = des.Interface
	}

	return nw
}

func canonicalizeNewReservationSpecificReservationInstancePropertiesLocalSsdsSet(c *Client, des, nw []ReservationSpecificReservationInstancePropertiesLocalSsds) []ReservationSpecificReservationInstancePropertiesLocalSsds {
	if des == nil {
		return nw
	}
	var reorderedNew []ReservationSpecificReservationInstancePropertiesLocalSsds
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareReservationSpecificReservationInstancePropertiesLocalSsdsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewReservationSpecificReservationInstancePropertiesLocalSsdsSlice(c *Client, des, nw []ReservationSpecificReservationInstancePropertiesLocalSsds) []ReservationSpecificReservationInstancePropertiesLocalSsds {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ReservationSpecificReservationInstancePropertiesLocalSsds
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewReservationSpecificReservationInstancePropertiesLocalSsds(c, &d, &n))
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
func diffReservation(c *Client, desired, actual *Reservation, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zone")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SpecificReservation, actual.SpecificReservation, dcl.Info{ObjectFunction: compareReservationSpecificReservationNewStyle, EmptyObject: EmptyReservationSpecificReservation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SpecificReservation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Commitment, actual.Commitment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Commitment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SpecificReservationRequired, actual.SpecificReservationRequired, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SpecificReservationRequired")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareReservationSpecificReservationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ReservationSpecificReservation)
	if !ok {
		desiredNotPointer, ok := d.(ReservationSpecificReservation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ReservationSpecificReservation or *ReservationSpecificReservation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ReservationSpecificReservation)
	if !ok {
		actualNotPointer, ok := a.(ReservationSpecificReservation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ReservationSpecificReservation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InstanceProperties, actual.InstanceProperties, dcl.Info{ObjectFunction: compareReservationSpecificReservationInstancePropertiesNewStyle, EmptyObject: EmptyReservationSpecificReservationInstanceProperties, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceProperties")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Count, actual.Count, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Count")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InUseCount, actual.InUseCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InUseCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareReservationSpecificReservationInstancePropertiesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ReservationSpecificReservationInstanceProperties)
	if !ok {
		desiredNotPointer, ok := d.(ReservationSpecificReservationInstanceProperties)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ReservationSpecificReservationInstanceProperties or *ReservationSpecificReservationInstanceProperties", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ReservationSpecificReservationInstanceProperties)
	if !ok {
		actualNotPointer, ok := a.(ReservationSpecificReservationInstanceProperties)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ReservationSpecificReservationInstanceProperties", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GuestAccelerators, actual.GuestAccelerators, dcl.Info{ObjectFunction: compareReservationSpecificReservationInstancePropertiesGuestAcceleratorsNewStyle, EmptyObject: EmptyReservationSpecificReservationInstancePropertiesGuestAccelerators, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GuestAccelerators")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinCpuPlatform, actual.MinCpuPlatform, dcl.Info{CustomDiff: canonicalizeReservationCPUPlatform, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinCpuPlatform")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalSsds, actual.LocalSsds, dcl.Info{ObjectFunction: compareReservationSpecificReservationInstancePropertiesLocalSsdsNewStyle, EmptyObject: EmptyReservationSpecificReservationInstancePropertiesLocalSsds, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocalSsds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareReservationSpecificReservationInstancePropertiesGuestAcceleratorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ReservationSpecificReservationInstancePropertiesGuestAccelerators)
	if !ok {
		desiredNotPointer, ok := d.(ReservationSpecificReservationInstancePropertiesGuestAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ReservationSpecificReservationInstancePropertiesGuestAccelerators or *ReservationSpecificReservationInstancePropertiesGuestAccelerators", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ReservationSpecificReservationInstancePropertiesGuestAccelerators)
	if !ok {
		actualNotPointer, ok := a.(ReservationSpecificReservationInstancePropertiesGuestAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ReservationSpecificReservationInstancePropertiesGuestAccelerators", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AcceleratorType, actual.AcceleratorType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorCount, actual.AcceleratorCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareReservationSpecificReservationInstancePropertiesLocalSsdsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ReservationSpecificReservationInstancePropertiesLocalSsds)
	if !ok {
		desiredNotPointer, ok := d.(ReservationSpecificReservationInstancePropertiesLocalSsds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ReservationSpecificReservationInstancePropertiesLocalSsds or *ReservationSpecificReservationInstancePropertiesLocalSsds", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ReservationSpecificReservationInstancePropertiesLocalSsds)
	if !ok {
		actualNotPointer, ok := a.(ReservationSpecificReservationInstancePropertiesLocalSsds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ReservationSpecificReservationInstancePropertiesLocalSsds", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DiskSizeGb, actual.DiskSizeGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Interface, actual.Interface, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Interface")); len(ds) != 0 || err != nil {
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
func (r *Reservation) urlNormalized() *Reservation {
	normalized := dcl.Copy(*r).(Reservation)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Commitment = dcl.SelfLinkToName(r.Commitment)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Reservation) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Zone), dcl.ValueOrEmptyString(n.Name)
}

func (r *Reservation) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Zone)
}

func (r *Reservation) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Zone), dcl.ValueOrEmptyString(n.Name)
}

func (r *Reservation) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Reservation resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Reservation) marshal(c *Client) ([]byte, error) {
	m, err := expandReservation(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Reservation: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalReservation decodes JSON responses into the Reservation resource schema.
func unmarshalReservation(b []byte, c *Client) (*Reservation, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapReservation(m, c)
}

func unmarshalMapReservation(m map[string]interface{}, c *Client) (*Reservation, error) {

	return flattenReservation(c, m), nil
}

// expandReservation expands Reservation into a JSON request object.
func expandReservation(c *Client, f *Reservation) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandReservationSpecificReservation(c, f.SpecificReservation); err != nil {
		return nil, fmt.Errorf("error expanding SpecificReservation into specificReservation: %w", err)
	} else if v != nil {
		m["specificReservation"] = v
	}
	if v := f.Commitment; !dcl.IsEmptyValueIndirect(v) {
		m["commitment"] = v
	}
	if v := f.SpecificReservationRequired; !dcl.IsEmptyValueIndirect(v) {
		m["specificReservationRequired"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenReservation flattens Reservation from a JSON request object into the
// Reservation type.
func flattenReservation(c *Client, i interface{}) *Reservation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Reservation{}
	res.Id = dcl.FlattenInteger(m["id"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Zone = dcl.FlattenString(m["zone"])
	res.Description = dcl.FlattenString(m["description"])
	res.Name = dcl.FlattenString(m["name"])
	res.SpecificReservation = flattenReservationSpecificReservation(c, m["specificReservation"])
	res.Commitment = dcl.FlattenString(m["commitment"])
	res.SpecificReservationRequired = dcl.FlattenBool(m["specificReservationRequired"])
	res.Status = flattenReservationStatusEnum(m["status"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandReservationSpecificReservationMap expands the contents of ReservationSpecificReservation into a JSON
// request object.
func expandReservationSpecificReservationMap(c *Client, f map[string]ReservationSpecificReservation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandReservationSpecificReservation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandReservationSpecificReservationSlice expands the contents of ReservationSpecificReservation into a JSON
// request object.
func expandReservationSpecificReservationSlice(c *Client, f []ReservationSpecificReservation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandReservationSpecificReservation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenReservationSpecificReservationMap flattens the contents of ReservationSpecificReservation from a JSON
// response object.
func flattenReservationSpecificReservationMap(c *Client, i interface{}) map[string]ReservationSpecificReservation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ReservationSpecificReservation{}
	}

	if len(a) == 0 {
		return map[string]ReservationSpecificReservation{}
	}

	items := make(map[string]ReservationSpecificReservation)
	for k, item := range a {
		items[k] = *flattenReservationSpecificReservation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenReservationSpecificReservationSlice flattens the contents of ReservationSpecificReservation from a JSON
// response object.
func flattenReservationSpecificReservationSlice(c *Client, i interface{}) []ReservationSpecificReservation {
	a, ok := i.([]interface{})
	if !ok {
		return []ReservationSpecificReservation{}
	}

	if len(a) == 0 {
		return []ReservationSpecificReservation{}
	}

	items := make([]ReservationSpecificReservation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenReservationSpecificReservation(c, item.(map[string]interface{})))
	}

	return items
}

// expandReservationSpecificReservation expands an instance of ReservationSpecificReservation into a JSON
// request object.
func expandReservationSpecificReservation(c *Client, f *ReservationSpecificReservation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandReservationSpecificReservationInstanceProperties(c, f.InstanceProperties); err != nil {
		return nil, fmt.Errorf("error expanding InstanceProperties into instanceProperties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["instanceProperties"] = v
	}
	if v := f.Count; !dcl.IsEmptyValueIndirect(v) {
		m["count"] = v
	}
	if v := f.InUseCount; !dcl.IsEmptyValueIndirect(v) {
		m["inUseCount"] = v
	}

	return m, nil
}

// flattenReservationSpecificReservation flattens an instance of ReservationSpecificReservation from a JSON
// response object.
func flattenReservationSpecificReservation(c *Client, i interface{}) *ReservationSpecificReservation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ReservationSpecificReservation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyReservationSpecificReservation
	}
	r.InstanceProperties = flattenReservationSpecificReservationInstanceProperties(c, m["instanceProperties"])
	r.Count = dcl.FlattenInteger(m["count"])
	r.InUseCount = dcl.FlattenInteger(m["inUseCount"])

	return r
}

// expandReservationSpecificReservationInstancePropertiesMap expands the contents of ReservationSpecificReservationInstanceProperties into a JSON
// request object.
func expandReservationSpecificReservationInstancePropertiesMap(c *Client, f map[string]ReservationSpecificReservationInstanceProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandReservationSpecificReservationInstanceProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandReservationSpecificReservationInstancePropertiesSlice expands the contents of ReservationSpecificReservationInstanceProperties into a JSON
// request object.
func expandReservationSpecificReservationInstancePropertiesSlice(c *Client, f []ReservationSpecificReservationInstanceProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandReservationSpecificReservationInstanceProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenReservationSpecificReservationInstancePropertiesMap flattens the contents of ReservationSpecificReservationInstanceProperties from a JSON
// response object.
func flattenReservationSpecificReservationInstancePropertiesMap(c *Client, i interface{}) map[string]ReservationSpecificReservationInstanceProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ReservationSpecificReservationInstanceProperties{}
	}

	if len(a) == 0 {
		return map[string]ReservationSpecificReservationInstanceProperties{}
	}

	items := make(map[string]ReservationSpecificReservationInstanceProperties)
	for k, item := range a {
		items[k] = *flattenReservationSpecificReservationInstanceProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenReservationSpecificReservationInstancePropertiesSlice flattens the contents of ReservationSpecificReservationInstanceProperties from a JSON
// response object.
func flattenReservationSpecificReservationInstancePropertiesSlice(c *Client, i interface{}) []ReservationSpecificReservationInstanceProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []ReservationSpecificReservationInstanceProperties{}
	}

	if len(a) == 0 {
		return []ReservationSpecificReservationInstanceProperties{}
	}

	items := make([]ReservationSpecificReservationInstanceProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenReservationSpecificReservationInstanceProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandReservationSpecificReservationInstanceProperties expands an instance of ReservationSpecificReservationInstanceProperties into a JSON
// request object.
func expandReservationSpecificReservationInstanceProperties(c *Client, f *ReservationSpecificReservationInstanceProperties) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}
	if v, err := expandReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice(c, f.GuestAccelerators); err != nil {
		return nil, fmt.Errorf("error expanding GuestAccelerators into guestAccelerators: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["guestAccelerators"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}
	if v, err := expandReservationSpecificReservationInstancePropertiesLocalSsdsSlice(c, f.LocalSsds); err != nil {
		return nil, fmt.Errorf("error expanding LocalSsds into localSsds: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["localSsds"] = v
	}

	return m, nil
}

// flattenReservationSpecificReservationInstanceProperties flattens an instance of ReservationSpecificReservationInstanceProperties from a JSON
// response object.
func flattenReservationSpecificReservationInstanceProperties(c *Client, i interface{}) *ReservationSpecificReservationInstanceProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ReservationSpecificReservationInstanceProperties{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyReservationSpecificReservationInstanceProperties
	}
	r.MachineType = dcl.FlattenString(m["machineType"])
	r.GuestAccelerators = flattenReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice(c, m["guestAccelerators"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])
	r.LocalSsds = flattenReservationSpecificReservationInstancePropertiesLocalSsdsSlice(c, m["localSsds"])

	return r
}

// expandReservationSpecificReservationInstancePropertiesGuestAcceleratorsMap expands the contents of ReservationSpecificReservationInstancePropertiesGuestAccelerators into a JSON
// request object.
func expandReservationSpecificReservationInstancePropertiesGuestAcceleratorsMap(c *Client, f map[string]ReservationSpecificReservationInstancePropertiesGuestAccelerators) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandReservationSpecificReservationInstancePropertiesGuestAccelerators(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice expands the contents of ReservationSpecificReservationInstancePropertiesGuestAccelerators into a JSON
// request object.
func expandReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice(c *Client, f []ReservationSpecificReservationInstancePropertiesGuestAccelerators) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandReservationSpecificReservationInstancePropertiesGuestAccelerators(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenReservationSpecificReservationInstancePropertiesGuestAcceleratorsMap flattens the contents of ReservationSpecificReservationInstancePropertiesGuestAccelerators from a JSON
// response object.
func flattenReservationSpecificReservationInstancePropertiesGuestAcceleratorsMap(c *Client, i interface{}) map[string]ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ReservationSpecificReservationInstancePropertiesGuestAccelerators{}
	}

	if len(a) == 0 {
		return map[string]ReservationSpecificReservationInstancePropertiesGuestAccelerators{}
	}

	items := make(map[string]ReservationSpecificReservationInstancePropertiesGuestAccelerators)
	for k, item := range a {
		items[k] = *flattenReservationSpecificReservationInstancePropertiesGuestAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice flattens the contents of ReservationSpecificReservationInstancePropertiesGuestAccelerators from a JSON
// response object.
func flattenReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice(c *Client, i interface{}) []ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []ReservationSpecificReservationInstancePropertiesGuestAccelerators{}
	}

	if len(a) == 0 {
		return []ReservationSpecificReservationInstancePropertiesGuestAccelerators{}
	}

	items := make([]ReservationSpecificReservationInstancePropertiesGuestAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenReservationSpecificReservationInstancePropertiesGuestAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandReservationSpecificReservationInstancePropertiesGuestAccelerators expands an instance of ReservationSpecificReservationInstancePropertiesGuestAccelerators into a JSON
// request object.
func expandReservationSpecificReservationInstancePropertiesGuestAccelerators(c *Client, f *ReservationSpecificReservationInstancePropertiesGuestAccelerators) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorType"] = v
	}
	if v := f.AcceleratorCount; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorCount"] = v
	}

	return m, nil
}

// flattenReservationSpecificReservationInstancePropertiesGuestAccelerators flattens an instance of ReservationSpecificReservationInstancePropertiesGuestAccelerators from a JSON
// response object.
func flattenReservationSpecificReservationInstancePropertiesGuestAccelerators(c *Client, i interface{}) *ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ReservationSpecificReservationInstancePropertiesGuestAccelerators{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyReservationSpecificReservationInstancePropertiesGuestAccelerators
	}
	r.AcceleratorType = dcl.FlattenString(m["acceleratorType"])
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])

	return r
}

// expandReservationSpecificReservationInstancePropertiesLocalSsdsMap expands the contents of ReservationSpecificReservationInstancePropertiesLocalSsds into a JSON
// request object.
func expandReservationSpecificReservationInstancePropertiesLocalSsdsMap(c *Client, f map[string]ReservationSpecificReservationInstancePropertiesLocalSsds) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandReservationSpecificReservationInstancePropertiesLocalSsds(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandReservationSpecificReservationInstancePropertiesLocalSsdsSlice expands the contents of ReservationSpecificReservationInstancePropertiesLocalSsds into a JSON
// request object.
func expandReservationSpecificReservationInstancePropertiesLocalSsdsSlice(c *Client, f []ReservationSpecificReservationInstancePropertiesLocalSsds) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandReservationSpecificReservationInstancePropertiesLocalSsds(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenReservationSpecificReservationInstancePropertiesLocalSsdsMap flattens the contents of ReservationSpecificReservationInstancePropertiesLocalSsds from a JSON
// response object.
func flattenReservationSpecificReservationInstancePropertiesLocalSsdsMap(c *Client, i interface{}) map[string]ReservationSpecificReservationInstancePropertiesLocalSsds {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ReservationSpecificReservationInstancePropertiesLocalSsds{}
	}

	if len(a) == 0 {
		return map[string]ReservationSpecificReservationInstancePropertiesLocalSsds{}
	}

	items := make(map[string]ReservationSpecificReservationInstancePropertiesLocalSsds)
	for k, item := range a {
		items[k] = *flattenReservationSpecificReservationInstancePropertiesLocalSsds(c, item.(map[string]interface{}))
	}

	return items
}

// flattenReservationSpecificReservationInstancePropertiesLocalSsdsSlice flattens the contents of ReservationSpecificReservationInstancePropertiesLocalSsds from a JSON
// response object.
func flattenReservationSpecificReservationInstancePropertiesLocalSsdsSlice(c *Client, i interface{}) []ReservationSpecificReservationInstancePropertiesLocalSsds {
	a, ok := i.([]interface{})
	if !ok {
		return []ReservationSpecificReservationInstancePropertiesLocalSsds{}
	}

	if len(a) == 0 {
		return []ReservationSpecificReservationInstancePropertiesLocalSsds{}
	}

	items := make([]ReservationSpecificReservationInstancePropertiesLocalSsds, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenReservationSpecificReservationInstancePropertiesLocalSsds(c, item.(map[string]interface{})))
	}

	return items
}

// expandReservationSpecificReservationInstancePropertiesLocalSsds expands an instance of ReservationSpecificReservationInstancePropertiesLocalSsds into a JSON
// request object.
func expandReservationSpecificReservationInstancePropertiesLocalSsds(c *Client, f *ReservationSpecificReservationInstancePropertiesLocalSsds) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskSizeGb"] = v
	}
	if v := f.Interface; !dcl.IsEmptyValueIndirect(v) {
		m["interface"] = v
	}

	return m, nil
}

// flattenReservationSpecificReservationInstancePropertiesLocalSsds flattens an instance of ReservationSpecificReservationInstancePropertiesLocalSsds from a JSON
// response object.
func flattenReservationSpecificReservationInstancePropertiesLocalSsds(c *Client, i interface{}) *ReservationSpecificReservationInstancePropertiesLocalSsds {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ReservationSpecificReservationInstancePropertiesLocalSsds{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyReservationSpecificReservationInstancePropertiesLocalSsds
	}
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.Interface = flattenReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(m["interface"])

	return r
}

// flattenReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumSlice flattens the contents of ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum from a JSON
// response object.
func flattenReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumSlice(c *Client, i interface{}) []ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum{}
	}

	if len(a) == 0 {
		return []ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum{}
	}

	items := make([]ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(item.(interface{})))
	}

	return items
}

// flattenReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum asserts that an interface is a string, and returns a
// pointer to a *ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum with the same value as that string.
func flattenReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(i interface{}) *ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum {
	s, ok := i.(string)
	if !ok {
		return ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumRef("")
	}

	return ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumRef(s)
}

// flattenReservationStatusEnumSlice flattens the contents of ReservationStatusEnum from a JSON
// response object.
func flattenReservationStatusEnumSlice(c *Client, i interface{}) []ReservationStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ReservationStatusEnum{}
	}

	if len(a) == 0 {
		return []ReservationStatusEnum{}
	}

	items := make([]ReservationStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenReservationStatusEnum(item.(interface{})))
	}

	return items
}

// flattenReservationStatusEnum asserts that an interface is a string, and returns a
// pointer to a *ReservationStatusEnum with the same value as that string.
func flattenReservationStatusEnum(i interface{}) *ReservationStatusEnum {
	s, ok := i.(string)
	if !ok {
		return ReservationStatusEnumRef("")
	}

	return ReservationStatusEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Reservation) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalReservation(b, c)
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
		if nr.Zone == nil && ncr.Zone == nil {
			c.Config.Logger.Info("Both Zone fields null - considering equal.")
		} else if nr.Zone == nil || ncr.Zone == nil {
			c.Config.Logger.Info("Only one Zone field is null - considering unequal.")
			return false
		} else if *nr.Zone != *ncr.Zone {
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

type reservationDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         reservationApiOperation
}

func convertFieldDiffToReservationOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]reservationDiff, error) {
	var diffs []reservationDiff
	for _, op := range ops {
		diff := reservationDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToreservationApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToreservationApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (reservationApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
