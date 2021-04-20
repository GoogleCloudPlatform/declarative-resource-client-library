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
package serviceusage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Service) validate() error {

	return nil
}

func serviceGetURL(userBasePath string, r *Service) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/services/{{name}}", "https://serviceusage.googleapis.com/v1/", userBasePath, params), nil
}

func serviceListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/services", "https://serviceusage.googleapis.com/v1/", userBasePath, params), nil

}

func serviceCreateURL(userBasePath, project, name string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"name":    name,
	}
	return dcl.URL("projects/{{project}}/services/{{name}}:enable", "https://serviceusage.googleapis.com/v1/", userBasePath, params), nil

}

func serviceDeleteURL(userBasePath string, r *Service) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/services/{{name}}:disable", "https://serviceusage.googleapis.com/v1/", userBasePath, params), nil
}

// serviceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type serviceApiOperation interface {
	do(context.Context, *Service, *Client) error
}

func (c *Client) listServiceRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := serviceListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ServiceMaxPage {
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

type listServiceOperation struct {
	Services []map[string]interface{} `json:"services"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listService(ctx context.Context, project, pageToken string, pageSize int32) ([]*Service, string, error) {
	b, err := c.listServiceRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listServiceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Service
	for _, v := range m.Services {
		res := flattenService(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllService(ctx context.Context, f func(*Service) bool, resources []*Service) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteService(ctx, res)
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

type deleteServiceOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createServiceOperation struct {
	response map[string]interface{}
}

func (op *createServiceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getServiceRaw(ctx context.Context, r *Service) ([]byte, error) {

	u, err := serviceGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) serviceDiffsForRawDesired(ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (initial, desired *Service, diffs []serviceDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Service
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Service); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Service, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// Simulate the resource not existing because create operation should be called anyway.
	rawInitial := &Service{}
	desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial)
	return nil, desired, nil, err

	c.Config.Logger.Infof("Found initial state for Service: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Service: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeServiceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Service: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Service: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffService(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeServiceInitialState(rawInitial, rawDesired *Service) (*Service, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeServiceDesiredState(rawDesired, rawInitial *Service, opts ...dcl.ApplyOption) (*Service, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeServiceNewState(c *Client, rawNew, rawDesired *Service) (*Service, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	return rawNew, nil
}

type serviceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         serviceApiOperation
	Diffs            []*dcl.FieldDiff
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
func diffService(c *Client, desired, actual *Service, opts ...dcl.ApplyOption) ([]serviceDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []serviceDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "name"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "EnumType", FieldName: "state"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: "ReferenceType", FieldName: "project"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, serviceDiff{RequiresRecreate: true, Diffs: ds})
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
	var deduped []serviceDiff
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
func compareServiceStateEnumSlice(c *Client, desired, actual []ServiceStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ServiceStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareServiceStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ServiceStateEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareServiceStateEnum(c *Client, desired, actual *ServiceStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Service) urlNormalized() *Service {
	normalized := dcl.Copy(*r).(Service)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Service) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Service) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Service) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Service) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Service resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Service) marshal(c *Client) ([]byte, error) {
	m, err := expandService(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Service: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalService decodes JSON responses into the Service resource schema.
func unmarshalService(b []byte, c *Client) (*Service, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapService(m, c)
}

func unmarshalMapService(m map[string]interface{}, c *Client) (*Service, error) {

	return flattenService(c, m), nil
}

// expandService expands Service into a JSON request object.
func expandService(c *Client, f *Service) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := dcl.DeriveField("projects/%s", f.Project, f.Project); err != nil {
		return nil, fmt.Errorf("error expanding Project into parent: %w", err)
	} else if v != nil {
		m["parent"] = v
	}

	return m, nil
}

// flattenService flattens Service from a JSON request object into the
// Service type.
func flattenService(c *Client, i interface{}) *Service {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Service{}
	r.Name = dcl.FlattenString(m["name"])
	r.State = flattenServiceStateEnum(m["state"])
	r.Project = dcl.FlattenString(m["parent"])

	return r
}

// flattenServiceStateEnumSlice flattens the contents of ServiceStateEnum from a JSON
// response object.
func flattenServiceStateEnumSlice(c *Client, i interface{}) []ServiceStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceStateEnum{}
	}

	if len(a) == 0 {
		return []ServiceStateEnum{}
	}

	items := make([]ServiceStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceStateEnum(item.(interface{})))
	}

	return items
}

// flattenServiceStateEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceStateEnum with the same value as that string.
func flattenServiceStateEnum(i interface{}) *ServiceStateEnum {
	s, ok := i.(string)
	if !ok {
		return ServiceStateEnumRef("")
	}

	return ServiceStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Service) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalService(b, c)
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
