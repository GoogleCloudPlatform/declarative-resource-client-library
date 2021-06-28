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
package bigqueryreservation

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

func (r *Reservation) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}

func reservationGetURL(userBasePath string, r *Reservation) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/reservations/{{name}}", "https://bigqueryreservation.googleapis.com/v1/", userBasePath, params), nil
}

func reservationListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/reservations", "https://bigqueryreservation.googleapis.com/v1/", userBasePath, params), nil

}

func reservationCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/reservations?reservationId={{name}}", "https://bigqueryreservation.googleapis.com/v1/", userBasePath, params), nil

}

func reservationDeleteURL(userBasePath string, r *Reservation) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/reservations/{{name}}", "https://bigqueryreservation.googleapis.com/v1/", userBasePath, params), nil
}

// reservationApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type reservationApiOperation interface {
	do(context.Context, *Reservation, *Client) error
}

// newUpdateReservationUpdateReservationRequest creates a request for an
// Reservation resource's UpdateReservation update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateReservationUpdateReservationRequest(ctx context.Context, f *Reservation, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.SlotCapacity; !dcl.IsEmptyValueIndirect(v) {
		req["slotCapacity"] = v
	}
	if v := f.IgnoreIdleSlots; !dcl.IsEmptyValueIndirect(v) {
		req["ignoreIdleSlots"] = v
	}
	return req, nil
}

// marshalUpdateReservationUpdateReservationRequest converts the update into
// the final JSON request body.
func marshalUpdateReservationUpdateReservationRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateReservationUpdateReservationOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateReservationUpdateReservationOperation) do(ctx context.Context, r *Reservation, c *Client) error {
	_, err := c.GetReservation(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateReservation")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.Diffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateReservationUpdateReservationRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateReservationUpdateReservationRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listReservationRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := reservationListURL(c.Config.BasePath, project, location)
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
	Reservations []map[string]interface{} `json:"reservations"`
	Token        string                   `json:"nextPageToken"`
}

func (c *Client) listReservation(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Reservation, string, error) {
	b, err := c.listReservationRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listReservationOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Reservation
	for _, v := range m.Reservations {
		res, err := unmarshalMapReservation(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
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
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Reservation: %w", err)
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

	project, location, name := r.createFields()
	u, err := reservationCreateURL(c.Config.BasePath, project, location, name)

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
	fmt.Printf("newDiffs: %v\n", diffs)
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

		return rawDesired, nil
	}

	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.SlotCapacity) {
		rawDesired.SlotCapacity = rawInitial.SlotCapacity
	}
	if dcl.BoolCanonicalize(rawDesired.IgnoreIdleSlots, rawInitial.IgnoreIdleSlots) {
		rawDesired.IgnoreIdleSlots = rawInitial.IgnoreIdleSlots
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeReservationNewState(c *Client, rawNew, rawDesired *Reservation) (*Reservation, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.SlotCapacity) && dcl.IsEmptyValueIndirect(rawDesired.SlotCapacity) {
		rawNew.SlotCapacity = rawDesired.SlotCapacity
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IgnoreIdleSlots) && dcl.IsEmptyValueIndirect(rawDesired.IgnoreIdleSlots) {
		rawNew.IgnoreIdleSlots = rawDesired.IgnoreIdleSlots
	} else {
		if dcl.BoolCanonicalize(rawDesired.IgnoreIdleSlots, rawNew.IgnoreIdleSlots) {
			rawNew.IgnoreIdleSlots = rawDesired.IgnoreIdleSlots
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTime) && dcl.IsEmptyValueIndirect(rawDesired.CreationTime) {
		rawNew.CreationTime = rawDesired.CreationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
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
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SlotCapacity, actual.SlotCapacity, dcl.Info{OperationSelector: dcl.TriggersOperation("updateReservationUpdateReservationOperation")}, fn.AddNest("SlotCapacity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IgnoreIdleSlots, actual.IgnoreIdleSlots, dcl.Info{OperationSelector: dcl.TriggersOperation("updateReservationUpdateReservationOperation")}, fn.AddNest("IgnoreIdleSlots")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreationTime, actual.CreationTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Reservation) urlNormalized() *Reservation {
	normalized := dcl.Copy(*r).(Reservation)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Reservation) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Reservation) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Reservation) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Reservation) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateReservation" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/reservations/{{name}}", "https://bigqueryreservation.googleapis.com/v1/", userBasePath, fields), nil

	}
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.SlotCapacity; !dcl.IsEmptyValueIndirect(v) {
		m["slotCapacity"] = v
	}
	if v := f.IgnoreIdleSlots; !dcl.IsEmptyValueIndirect(v) {
		m["ignoreIdleSlots"] = v
	}
	if v := f.CreationTime; !dcl.IsEmptyValueIndirect(v) {
		m["creationTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
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
	res.Name = dcl.FlattenString(m["name"])
	res.SlotCapacity = dcl.FlattenInteger(m["slotCapacity"])
	res.IgnoreIdleSlots = dcl.FlattenBool(m["ignoreIdleSlots"])
	res.CreationTime = dcl.FlattenString(m["creationTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
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

	case "updateReservationUpdateReservationOperation":
		return &updateReservationUpdateReservationOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
