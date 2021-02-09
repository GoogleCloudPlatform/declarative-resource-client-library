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

	"github.com/mohae/deepcopy"
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
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
		res := flattenReservation(c, v)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
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
	_, err = c.GetReservation(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createReservationOperation struct{}

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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
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

	if _, err := c.GetReservation(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getReservationRaw(ctx context.Context, r *Reservation) ([]byte, error) {

	u, err := reservationGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) reservationDiffsForRawDesired(ctx context.Context, rawDesired *Reservation, opts ...dcl.ApplyOption) (initial, desired *Reservation, diffs []reservationDiff, err error) {
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Reservation); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected Reservation, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SpecificReservation = canonicalizeReservationSpecificReservation(rawDesired.SpecificReservation, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.IsZeroValue(rawDesired.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.IsZeroValue(rawDesired.Zone) {
		rawDesired.Zone = rawInitial.Zone
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.SpecificReservation = canonicalizeReservationSpecificReservation(rawDesired.SpecificReservation, rawInitial.SpecificReservation, opts...)
	if dcl.IsZeroValue(rawDesired.Commitment) {
		rawDesired.Commitment = rawInitial.Commitment
	}
	if dcl.IsZeroValue(rawDesired.SpecificReservationRequired) {
		rawDesired.SpecificReservationRequired = rawInitial.SpecificReservationRequired
	}
	if dcl.IsZeroValue(rawDesired.Status) {
		rawDesired.Status = rawInitial.Status
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.Zone) && dcl.IsEmptyValueIndirect(rawDesired.Zone) {
		rawNew.Zone = rawDesired.Zone
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SpecificReservation) && dcl.IsEmptyValueIndirect(rawDesired.SpecificReservation) {
		rawNew.SpecificReservation = rawDesired.SpecificReservation
	} else {
		rawNew.SpecificReservation = canonicalizeNewReservationSpecificReservation(c, rawDesired.SpecificReservation, rawNew.SpecificReservation)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Commitment) && dcl.IsEmptyValueIndirect(rawDesired.Commitment) {
		rawNew.Commitment = rawDesired.Commitment
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SpecificReservationRequired) && dcl.IsEmptyValueIndirect(rawDesired.SpecificReservationRequired) {
		rawNew.SpecificReservationRequired = rawDesired.SpecificReservationRequired
	} else {
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Reservation)
		_ = r
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
			if !compareReservationSpecificReservation(c, &d, &n) {
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

func canonicalizeReservationSpecificReservationInstanceProperties(des, initial *ReservationSpecificReservationInstanceProperties, opts ...dcl.ApplyOption) *ReservationSpecificReservationInstanceProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Reservation)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MachineType) {
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

	if canonicalizeReservationCPUPlatform(des.MinCpuPlatform, nw.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}

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
			if !compareReservationSpecificReservationInstanceProperties(c, &d, &n) {
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

func canonicalizeReservationSpecificReservationInstancePropertiesGuestAccelerators(des, initial *ReservationSpecificReservationInstancePropertiesGuestAccelerators, opts ...dcl.ApplyOption) *ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Reservation)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AcceleratorType) {
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
			if !compareReservationSpecificReservationInstancePropertiesGuestAccelerators(c, &d, &n) {
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

func canonicalizeReservationSpecificReservationInstancePropertiesLocalSsds(des, initial *ReservationSpecificReservationInstancePropertiesLocalSsds, opts ...dcl.ApplyOption) *ReservationSpecificReservationInstancePropertiesLocalSsds {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Reservation)
		_ = r
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
			if !compareReservationSpecificReservationInstancePropertiesLocalSsds(c, &d, &n) {
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

type reservationDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         reservationApiOperation
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
func diffReservation(c *Client, desired, actual *Reservation, opts ...dcl.ApplyOption) ([]reservationDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []reservationDiff
	if !dcl.IsZeroValue(desired.Id) && (dcl.IsZeroValue(actual.Id) || !reflect.DeepEqual(*desired.Id, *actual.Id)) {
		c.Config.Logger.Infof("Detected diff in Id.\nDESIRED: %v\nACTUAL: %v", desired.Id, actual.Id)
		diffs = append(diffs, reservationDiff{
			RequiresRecreate: true,
			FieldName:        "Id",
		})
	}
	if !dcl.IsZeroValue(desired.Zone) && (dcl.IsZeroValue(actual.Zone) || !reflect.DeepEqual(*desired.Zone, *actual.Zone)) {
		c.Config.Logger.Infof("Detected diff in Zone.\nDESIRED: %v\nACTUAL: %v", desired.Zone, actual.Zone)
		diffs = append(diffs, reservationDiff{
			RequiresRecreate: true,
			FieldName:        "Zone",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, reservationDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, reservationDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if compareReservationSpecificReservation(c, desired.SpecificReservation, actual.SpecificReservation) {
		c.Config.Logger.Infof("Detected diff in SpecificReservation.\nDESIRED: %v\nACTUAL: %v", desired.SpecificReservation, actual.SpecificReservation)
		diffs = append(diffs, reservationDiff{
			RequiresRecreate: true,
			FieldName:        "SpecificReservation",
		})
	}
	if !dcl.IsZeroValue(desired.Commitment) && (dcl.IsZeroValue(actual.Commitment) || !reflect.DeepEqual(*desired.Commitment, *actual.Commitment)) {
		c.Config.Logger.Infof("Detected diff in Commitment.\nDESIRED: %v\nACTUAL: %v", desired.Commitment, actual.Commitment)
		diffs = append(diffs, reservationDiff{
			RequiresRecreate: true,
			FieldName:        "Commitment",
		})
	}
	if !dcl.IsZeroValue(desired.SpecificReservationRequired) && (dcl.IsZeroValue(actual.SpecificReservationRequired) || !reflect.DeepEqual(*desired.SpecificReservationRequired, *actual.SpecificReservationRequired)) {
		c.Config.Logger.Infof("Detected diff in SpecificReservationRequired.\nDESIRED: %v\nACTUAL: %v", desired.SpecificReservationRequired, actual.SpecificReservationRequired)
		diffs = append(diffs, reservationDiff{
			RequiresRecreate: true,
			FieldName:        "SpecificReservationRequired",
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
	var deduped []reservationDiff
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
func compareReservationSpecificReservationSlice(c *Client, desired, actual []ReservationSpecificReservation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ReservationSpecificReservation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareReservationSpecificReservation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ReservationSpecificReservation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareReservationSpecificReservation(c *Client, desired, actual *ReservationSpecificReservation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.InstanceProperties == nil && desired.InstanceProperties != nil && !dcl.IsEmptyValueIndirect(desired.InstanceProperties) {
		c.Config.Logger.Infof("desired InstanceProperties %s - but actually nil", dcl.SprintResource(desired.InstanceProperties))
		return true
	}
	if compareReservationSpecificReservationInstanceProperties(c, desired.InstanceProperties, actual.InstanceProperties) && !dcl.IsZeroValue(desired.InstanceProperties) {
		c.Config.Logger.Infof("Diff in InstanceProperties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InstanceProperties), dcl.SprintResource(actual.InstanceProperties))
		return true
	}
	if actual.Count == nil && desired.Count != nil && !dcl.IsEmptyValueIndirect(desired.Count) {
		c.Config.Logger.Infof("desired Count %s - but actually nil", dcl.SprintResource(desired.Count))
		return true
	}
	if !reflect.DeepEqual(desired.Count, actual.Count) && !dcl.IsZeroValue(desired.Count) && !(dcl.IsEmptyValueIndirect(desired.Count) && dcl.IsZeroValue(actual.Count)) {
		c.Config.Logger.Infof("Diff in Count. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Count), dcl.SprintResource(actual.Count))
		return true
	}
	if actual.InUseCount == nil && desired.InUseCount != nil && !dcl.IsEmptyValueIndirect(desired.InUseCount) {
		c.Config.Logger.Infof("desired InUseCount %s - but actually nil", dcl.SprintResource(desired.InUseCount))
		return true
	}
	if !reflect.DeepEqual(desired.InUseCount, actual.InUseCount) && !dcl.IsZeroValue(desired.InUseCount) && !(dcl.IsEmptyValueIndirect(desired.InUseCount) && dcl.IsZeroValue(actual.InUseCount)) {
		c.Config.Logger.Infof("Diff in InUseCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InUseCount), dcl.SprintResource(actual.InUseCount))
		return true
	}
	return false
}
func compareReservationSpecificReservationInstancePropertiesSlice(c *Client, desired, actual []ReservationSpecificReservationInstanceProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ReservationSpecificReservationInstanceProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareReservationSpecificReservationInstanceProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ReservationSpecificReservationInstanceProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareReservationSpecificReservationInstanceProperties(c *Client, desired, actual *ReservationSpecificReservationInstanceProperties) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MachineType == nil && desired.MachineType != nil && !dcl.IsEmptyValueIndirect(desired.MachineType) {
		c.Config.Logger.Infof("desired MachineType %s - but actually nil", dcl.SprintResource(desired.MachineType))
		return true
	}
	if !reflect.DeepEqual(desired.MachineType, actual.MachineType) && !dcl.IsZeroValue(desired.MachineType) && !(dcl.IsEmptyValueIndirect(desired.MachineType) && dcl.IsZeroValue(actual.MachineType)) {
		c.Config.Logger.Infof("Diff in MachineType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MachineType), dcl.SprintResource(actual.MachineType))
		return true
	}
	if actual.GuestAccelerators == nil && desired.GuestAccelerators != nil && !dcl.IsEmptyValueIndirect(desired.GuestAccelerators) {
		c.Config.Logger.Infof("desired GuestAccelerators %s - but actually nil", dcl.SprintResource(desired.GuestAccelerators))
		return true
	}
	if compareReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice(c, desired.GuestAccelerators, actual.GuestAccelerators) && !dcl.IsZeroValue(desired.GuestAccelerators) {
		c.Config.Logger.Infof("Diff in GuestAccelerators. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GuestAccelerators), dcl.SprintResource(actual.GuestAccelerators))
		return true
	}
	if actual.MinCpuPlatform == nil && desired.MinCpuPlatform != nil && !dcl.IsEmptyValueIndirect(desired.MinCpuPlatform) {
		c.Config.Logger.Infof("desired MinCpuPlatform %s - but actually nil", dcl.SprintResource(desired.MinCpuPlatform))
		return true
	}
	if !canonicalizeReservationCPUPlatform(desired.MinCpuPlatform, actual.MinCpuPlatform) && !dcl.IsZeroValue(desired.MinCpuPlatform) {
		c.Config.Logger.Infof("Diff in MinCpuPlatform. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinCpuPlatform), dcl.SprintResource(actual.MinCpuPlatform))
		return true
	}
	if actual.LocalSsds == nil && desired.LocalSsds != nil && !dcl.IsEmptyValueIndirect(desired.LocalSsds) {
		c.Config.Logger.Infof("desired LocalSsds %s - but actually nil", dcl.SprintResource(desired.LocalSsds))
		return true
	}
	if compareReservationSpecificReservationInstancePropertiesLocalSsdsSlice(c, desired.LocalSsds, actual.LocalSsds) && !dcl.IsZeroValue(desired.LocalSsds) {
		c.Config.Logger.Infof("Diff in LocalSsds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalSsds), dcl.SprintResource(actual.LocalSsds))
		return true
	}
	return false
}
func compareReservationSpecificReservationInstancePropertiesGuestAcceleratorsSlice(c *Client, desired, actual []ReservationSpecificReservationInstancePropertiesGuestAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ReservationSpecificReservationInstancePropertiesGuestAccelerators, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareReservationSpecificReservationInstancePropertiesGuestAccelerators(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ReservationSpecificReservationInstancePropertiesGuestAccelerators, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareReservationSpecificReservationInstancePropertiesGuestAccelerators(c *Client, desired, actual *ReservationSpecificReservationInstancePropertiesGuestAccelerators) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AcceleratorType == nil && desired.AcceleratorType != nil && !dcl.IsEmptyValueIndirect(desired.AcceleratorType) {
		c.Config.Logger.Infof("desired AcceleratorType %s - but actually nil", dcl.SprintResource(desired.AcceleratorType))
		return true
	}
	if !reflect.DeepEqual(desired.AcceleratorType, actual.AcceleratorType) && !dcl.IsZeroValue(desired.AcceleratorType) && !(dcl.IsEmptyValueIndirect(desired.AcceleratorType) && dcl.IsZeroValue(actual.AcceleratorType)) {
		c.Config.Logger.Infof("Diff in AcceleratorType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorType), dcl.SprintResource(actual.AcceleratorType))
		return true
	}
	if actual.AcceleratorCount == nil && desired.AcceleratorCount != nil && !dcl.IsEmptyValueIndirect(desired.AcceleratorCount) {
		c.Config.Logger.Infof("desired AcceleratorCount %s - but actually nil", dcl.SprintResource(desired.AcceleratorCount))
		return true
	}
	if !reflect.DeepEqual(desired.AcceleratorCount, actual.AcceleratorCount) && !dcl.IsZeroValue(desired.AcceleratorCount) && !(dcl.IsEmptyValueIndirect(desired.AcceleratorCount) && dcl.IsZeroValue(actual.AcceleratorCount)) {
		c.Config.Logger.Infof("Diff in AcceleratorCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorCount), dcl.SprintResource(actual.AcceleratorCount))
		return true
	}
	return false
}
func compareReservationSpecificReservationInstancePropertiesLocalSsdsSlice(c *Client, desired, actual []ReservationSpecificReservationInstancePropertiesLocalSsds) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ReservationSpecificReservationInstancePropertiesLocalSsds, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareReservationSpecificReservationInstancePropertiesLocalSsds(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ReservationSpecificReservationInstancePropertiesLocalSsds, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareReservationSpecificReservationInstancePropertiesLocalSsds(c *Client, desired, actual *ReservationSpecificReservationInstancePropertiesLocalSsds) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DiskSizeGb == nil && desired.DiskSizeGb != nil && !dcl.IsEmptyValueIndirect(desired.DiskSizeGb) {
		c.Config.Logger.Infof("desired DiskSizeGb %s - but actually nil", dcl.SprintResource(desired.DiskSizeGb))
		return true
	}
	if !reflect.DeepEqual(desired.DiskSizeGb, actual.DiskSizeGb) && !dcl.IsZeroValue(desired.DiskSizeGb) && !(dcl.IsEmptyValueIndirect(desired.DiskSizeGb) && dcl.IsZeroValue(actual.DiskSizeGb)) {
		c.Config.Logger.Infof("Diff in DiskSizeGb. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskSizeGb), dcl.SprintResource(actual.DiskSizeGb))
		return true
	}
	if actual.Interface == nil && desired.Interface != nil && !dcl.IsEmptyValueIndirect(desired.Interface) {
		c.Config.Logger.Infof("desired Interface %s - but actually nil", dcl.SprintResource(desired.Interface))
		return true
	}
	if !reflect.DeepEqual(desired.Interface, actual.Interface) && !dcl.IsZeroValue(desired.Interface) && !(dcl.IsEmptyValueIndirect(desired.Interface) && dcl.IsZeroValue(actual.Interface)) {
		c.Config.Logger.Infof("Diff in Interface. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Interface), dcl.SprintResource(actual.Interface))
		return true
	}
	return false
}
func compareReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumSlice(c *Client, desired, actual []ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(c *Client, desired, actual *ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareReservationStatusEnumSlice(c *Client, desired, actual []ReservationStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ReservationStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareReservationStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ReservationStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareReservationStatusEnum(c *Client, desired, actual *ReservationStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Reservation) urlNormalized() *Reservation {
	normalized := deepcopy.Copy(*r).(Reservation)
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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

	r := &Reservation{}
	r.Id = dcl.FlattenInteger(m["id"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Zone = dcl.FlattenString(m["zone"])
	r.Description = dcl.FlattenString(m["description"])
	r.Name = dcl.FlattenString(m["name"])
	r.SpecificReservation = flattenReservationSpecificReservation(c, m["specificReservation"])
	r.Commitment = dcl.FlattenString(m["commitment"])
	r.SpecificReservationRequired = dcl.FlattenBool(m["specificReservationRequired"])
	r.Status = flattenReservationStatusEnum(m["status"])
	r.Project = dcl.FlattenString(m["project"])

	return r
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
		items = append(items, *flattenReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenReservationStatusEnum(item.(map[string]interface{})))
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
