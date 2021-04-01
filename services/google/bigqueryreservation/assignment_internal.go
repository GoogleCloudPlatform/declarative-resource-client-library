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
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Assignment) validate() error {

	if err := dcl.Required(r, "assignee"); err != nil {
		return err
	}
	if err := dcl.Required(r, "jobType"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Reservation, "Reservation"); err != nil {
		return err
	}
	return nil
}

func assignmentGetURL(userBasePath string, r *Assignment) (string, error) {
	params := map[string]interface{}{
		"project":     dcl.ValueOrEmptyString(r.Project),
		"location":    dcl.ValueOrEmptyString(r.Location),
		"reservation": dcl.ValueOrEmptyString(r.Reservation),
		"assignee":    dcl.ValueOrEmptyString(r.Assignee),
		"jobType":     dcl.ValueOrEmptyString(r.JobType),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments?fetchId=%s%s", "https://bigqueryreservation.googleapis.com/v1/", userBasePath, params), nil
}

func assignmentListURL(userBasePath, project, location, reservation string) (string, error) {
	params := map[string]interface{}{
		"project":     project,
		"location":    location,
		"reservation": reservation,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments", "https://bigqueryreservation.googleapis.com/v1/", userBasePath, params), nil

}

func assignmentCreateURL(userBasePath, project, location, reservation string) (string, error) {
	params := map[string]interface{}{
		"project":     project,
		"location":    location,
		"reservation": reservation,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments", "https://bigqueryreservation.googleapis.com/v1/", userBasePath, params), nil

}

func assignmentDeleteURL(userBasePath string, r *Assignment) (string, error) {
	params := map[string]interface{}{
		"project":     dcl.ValueOrEmptyString(r.Project),
		"location":    dcl.ValueOrEmptyString(r.Location),
		"reservation": dcl.ValueOrEmptyString(r.Reservation),
		"name":        dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}", "https://bigqueryreservation.googleapis.com/v1/", userBasePath, params), nil
}

// assignmentApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type assignmentApiOperation interface {
	do(context.Context, *Assignment, *Client) error
}

func (c *Client) listAssignmentRaw(ctx context.Context, project, location, reservation, pageToken string, pageSize int32) ([]byte, error) {
	u, err := assignmentListURL(c.Config.BasePath, project, location, reservation)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AssignmentMaxPage {
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

type listAssignmentOperation struct {
	Assignments []map[string]interface{} `json:"assignments"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listAssignment(ctx context.Context, project, location, reservation, pageToken string, pageSize int32) ([]*Assignment, string, error) {
	b, err := c.listAssignmentRaw(ctx, project, location, reservation, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAssignmentOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Assignment
	for _, v := range m.Assignments {
		res := flattenAssignment(c, v)
		res.Project = &project
		res.Location = &location
		res.Reservation = &reservation
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAssignment(ctx context.Context, f func(*Assignment) bool, resources []*Assignment) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAssignment(ctx, res)
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

type deleteAssignmentOperation struct{}

func (op *deleteAssignmentOperation) do(ctx context.Context, r *Assignment, c *Client) error {

	_, err := c.GetAssignment(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Assignment not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAssignment checking for existence. error: %v", err)
		return err
	}

	u, err := assignmentDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Assignment: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAssignment(ctx, r.urlNormalized())
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
type createAssignmentOperation struct {
	response map[string]interface{}
}

func (op *createAssignmentOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAssignmentOperation) do(ctx context.Context, r *Assignment, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, reservation := r.createFields()
	u, err := assignmentCreateURL(c.Config.BasePath, project, location, reservation)

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

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string")
	}
	r.Name = &name

	if _, err := c.GetAssignment(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAssignmentRaw(ctx context.Context, r *Assignment) ([]byte, error) {

	u, err := assignmentGetURL(c.Config.BasePath, r.urlNormalized())
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

	b, err = dcl.ExtractElementFromList(b, "assignments", r.matcher(c))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) assignmentDiffsForRawDesired(ctx context.Context, rawDesired *Assignment, opts ...dcl.ApplyOption) (initial, desired *Assignment, diffs []assignmentDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Assignment
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Assignment); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Assignment, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeAssignmentDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAssignment(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Assignment resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Assignment resource: %v", err)
		}
		c.Config.Logger.Info("Found that Assignment resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAssignmentDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Assignment: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Assignment: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAssignmentInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Assignment: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAssignmentDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Assignment: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAssignment(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAssignmentInitialState(rawInitial, rawDesired *Assignment) (*Assignment, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAssignmentDesiredState(rawDesired, rawInitial *Assignment, opts ...dcl.ApplyOption) (*Assignment, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Assignee, rawInitial.Assignee) {
		rawDesired.Assignee = rawInitial.Assignee
	}
	if dcl.IsZeroValue(rawDesired.JobType) {
		rawDesired.JobType = rawInitial.JobType
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.NameToSelfLink(rawDesired.Reservation, rawInitial.Reservation) {
		rawDesired.Reservation = rawInitial.Reservation
	}

	return rawDesired, nil
}

func canonicalizeAssignmentNewState(c *Client, rawNew, rawDesired *Assignment) (*Assignment, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Assignee) && dcl.IsEmptyValueIndirect(rawDesired.Assignee) {
		rawNew.Assignee = rawDesired.Assignee
	} else {
		if dcl.StringCanonicalize(rawDesired.Assignee, rawNew.Assignee) {
			rawNew.Assignee = rawDesired.Assignee
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.JobType) && dcl.IsEmptyValueIndirect(rawDesired.JobType) {
		rawNew.JobType = rawDesired.JobType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.Reservation = rawDesired.Reservation

	return rawNew, nil
}

type assignmentDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         assignmentApiOperation
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
func diffAssignment(c *Client, desired, actual *Assignment, opts ...dcl.ApplyOption) ([]assignmentDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []assignmentDiff
	if !dcl.IsZeroValue(desired.Assignee) && !dcl.StringCanonicalize(desired.Assignee, actual.Assignee) {
		c.Config.Logger.Infof("Detected diff in Assignee.\nDESIRED: %v\nACTUAL: %v", desired.Assignee, actual.Assignee)
		diffs = append(diffs, assignmentDiff{
			RequiresRecreate: true,
			FieldName:        "Assignee",
		})
	}
	if !reflect.DeepEqual(desired.JobType, actual.JobType) {
		c.Config.Logger.Infof("Detected diff in JobType.\nDESIRED: %v\nACTUAL: %v", desired.JobType, actual.JobType)
		diffs = append(diffs, assignmentDiff{
			RequiresRecreate: true,
			FieldName:        "JobType",
		})
	}
	if !reflect.DeepEqual(desired.State, actual.State) {
		c.Config.Logger.Infof("Detected diff in State.\nDESIRED: %v\nACTUAL: %v", desired.State, actual.State)
		diffs = append(diffs, assignmentDiff{
			RequiresRecreate: true,
			FieldName:        "State",
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
	var deduped []assignmentDiff
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
func compareAssignmentJobTypeEnumSlice(c *Client, desired, actual []AssignmentJobTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AssignmentJobTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAssignmentJobTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AssignmentJobTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAssignmentJobTypeEnum(c *Client, desired, actual *AssignmentJobTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAssignmentStateEnumSlice(c *Client, desired, actual []AssignmentStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AssignmentStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAssignmentStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AssignmentStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAssignmentStateEnum(c *Client, desired, actual *AssignmentStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Assignment) urlNormalized() *Assignment {
	normalized := deepcopy.Copy(*r).(Assignment)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Assignee = dcl.SelfLinkToName(r.Assignee)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Reservation = dcl.SelfLinkToName(r.Reservation)
	return &normalized
}

func (r *Assignment) getFields() (string, string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Reservation), dcl.ValueOrEmptyString(n.Assignee), dcl.ValueOrEmptyString(n.JobType)
}

func (r *Assignment) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Reservation)
}

func (r *Assignment) deleteFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Reservation), dcl.ValueOrEmptyString(n.Name)
}

func (r *Assignment) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Assignment resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Assignment) marshal(c *Client) ([]byte, error) {
	m, err := expandAssignment(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Assignment: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAssignment decodes JSON responses into the Assignment resource schema.
func unmarshalAssignment(b []byte, c *Client) (*Assignment, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAssignment(m, c)
}

func unmarshalMapAssignment(m map[string]interface{}, c *Client) (*Assignment, error) {

	return flattenAssignment(c, m), nil
}

// expandAssignment expands Assignment into a JSON request object.
func expandAssignment(c *Client, f *Assignment) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Assignee; !dcl.IsEmptyValueIndirect(v) {
		m["assignee"] = v
	}
	if v := f.JobType; !dcl.IsEmptyValueIndirect(v) {
		m["jobType"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Reservation into reservation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reservation"] = v
	}

	return m, nil
}

// flattenAssignment flattens Assignment from a JSON request object into the
// Assignment type.
func flattenAssignment(c *Client, i interface{}) *Assignment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Assignment{}
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	r.Assignee = dcl.FlattenString(m["assignee"])
	r.JobType = flattenAssignmentJobTypeEnum(m["jobType"])
	r.State = flattenAssignmentStateEnum(m["state"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])
	r.Reservation = dcl.FlattenString(m["reservation"])

	return r
}

// flattenAssignmentJobTypeEnumSlice flattens the contents of AssignmentJobTypeEnum from a JSON
// response object.
func flattenAssignmentJobTypeEnumSlice(c *Client, i interface{}) []AssignmentJobTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AssignmentJobTypeEnum{}
	}

	if len(a) == 0 {
		return []AssignmentJobTypeEnum{}
	}

	items := make([]AssignmentJobTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAssignmentJobTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAssignmentJobTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AssignmentJobTypeEnum with the same value as that string.
func flattenAssignmentJobTypeEnum(i interface{}) *AssignmentJobTypeEnum {
	s, ok := i.(string)
	if !ok {
		return AssignmentJobTypeEnumRef("")
	}

	return AssignmentJobTypeEnumRef(s)
}

// flattenAssignmentStateEnumSlice flattens the contents of AssignmentStateEnum from a JSON
// response object.
func flattenAssignmentStateEnumSlice(c *Client, i interface{}) []AssignmentStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AssignmentStateEnum{}
	}

	if len(a) == 0 {
		return []AssignmentStateEnum{}
	}

	items := make([]AssignmentStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAssignmentStateEnum(item.(interface{})))
	}

	return items
}

// flattenAssignmentStateEnum asserts that an interface is a string, and returns a
// pointer to a *AssignmentStateEnum with the same value as that string.
func flattenAssignmentStateEnum(i interface{}) *AssignmentStateEnum {
	s, ok := i.(string)
	if !ok {
		return AssignmentStateEnumRef("")
	}

	return AssignmentStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Assignment) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAssignment(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Assignee == nil && ncr.Assignee == nil {
			c.Config.Logger.Info("Both Assignee fields null - considering equal.")
		} else if nr.Assignee == nil || ncr.Assignee == nil {
			c.Config.Logger.Info("Only one Assignee field is null - considering unequal.")
			return false
		} else if *nr.Assignee != *ncr.Assignee {
			return false
		}
		if nr.JobType == nil && ncr.JobType == nil {
			c.Config.Logger.Info("Both JobType fields null - considering equal.")
		} else if nr.JobType == nil || ncr.JobType == nil {
			c.Config.Logger.Info("Only one JobType field is null - considering unequal.")
			return false
		} else if *nr.JobType != *ncr.JobType {
			return false
		}
		return true
	}
}
