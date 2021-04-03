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
package servicenetworking

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Connection) validate() error {

	if err := dcl.Required(r, "network"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "reservedPeeringRanges"); err != nil {
		return err
	}
	return nil
}

func connectionGetURL(userBasePath string, r *Connection) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"network": dcl.ValueOrEmptyString(r.Network),
		"service": dcl.ValueOrEmptyString(r.Service),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("services/{{service}}/connections?network=projects%2F{{project}}%2Fglobal%2Fnetworks%2F{{network}}&fetchId=%s", "https://servicenetworking.googleapis.com/v1/", userBasePath, params), nil
}

func connectionListURL(userBasePath, project, network, service string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"network": network,
		"service": service,
	}
	return dcl.URL("services/{{service}}/connections?network=projects%2F{{project}}%2Fglobal%2Fnetworks%2F{{network}}", "https://servicenetworking.googleapis.com/v1/", userBasePath, params), nil

}

func connectionCreateURL(userBasePath, service string) (string, error) {
	params := map[string]interface{}{
		"service": service,
	}
	return dcl.URL("services/{{service}}/connections/-?force=true", "https://servicenetworking.googleapis.com/v1/", userBasePath, params), nil

}

func connectionDeleteURL(userBasePath string, r *Connection) (string, error) {
	params := map[string]interface{}{
		"service": dcl.ValueOrEmptyString(r.Service),
	}
	return dcl.URL("services/{{service}}/connections/-?force=true", "https://servicenetworking.googleapis.com/v1/", userBasePath, params), nil
}

// connectionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type connectionApiOperation interface {
	do(context.Context, *Connection, *Client) error
}

// newUpdateConnectionPatchRequest creates a request for an
// Connection resource's patch update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateConnectionPatchRequest(ctx context.Context, f *Connection, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/global/networks/%s", f.Network, f.Project, f.Network); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["network"] = v
	}
	if v := f.ReservedPeeringRanges; !dcl.IsEmptyValueIndirect(v) {
		req["reservedPeeringRanges"] = v
	}
	if v, err := dcl.DeriveField("services/%s", f.Service, f.Service); err != nil {
		return nil, fmt.Errorf("error expanding Service into service: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["service"] = v
	}
	return req, nil
}

// marshalUpdateConnectionPatchRequest converts the update into
// the final JSON request body.
func marshalUpdateConnectionPatchRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateConnectionPatchOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateConnectionPatchOperation) do(ctx context.Context, r *Connection, c *Client) error {
	_, err := c.GetConnection(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "patch")
	if err != nil {
		return err
	}

	req, err := newUpdateConnectionPatchRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateConnectionPatchRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://servicenetworking.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listConnectionRaw(ctx context.Context, project, network, service, pageToken string, pageSize int32) ([]byte, error) {
	u, err := connectionListURL(c.Config.BasePath, project, network, service)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ConnectionMaxPage {
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

type listConnectionOperation struct {
	Connections []map[string]interface{} `json:"connections"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listConnection(ctx context.Context, project, network, service, pageToken string, pageSize int32) ([]*Connection, string, error) {
	b, err := c.listConnectionRaw(ctx, project, network, service, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listConnectionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Connection
	for _, v := range m.Connections {
		res := flattenConnection(c, v)
		res.Project = &project
		res.Network = &network
		res.Service = &service
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllConnection(ctx context.Context, f func(*Connection) bool, resources []*Connection) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteConnection(ctx, res)
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

type deleteConnectionOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createConnectionOperation struct {
	response map[string]interface{}
}

func (op *createConnectionOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createConnectionOperation) do(ctx context.Context, r *Connection, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	service := r.createFields()
	u, err := connectionCreateURL(c.Config.BasePath, service)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://servicenetworking.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["peering"].(string)
	if !ok {
		return fmt.Errorf("expected peering to be a string")
	}
	r.Name = &name

	if _, err := c.GetConnection(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getConnectionRaw(ctx context.Context, r *Connection) ([]byte, error) {
	if dcl.IsZeroValue(r.Service) {
		r.Service = dcl.String("services/servicenetworking.googleapis.com")
	}

	u, err := connectionGetURL(c.Config.BasePath, r.urlNormalized())
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

	b, err = dcl.ExtractElementFromList(b, "connections", r.matcher(c))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) connectionDiffsForRawDesired(ctx context.Context, rawDesired *Connection, opts ...dcl.ApplyOption) (initial, desired *Connection, diffs []connectionDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Connection
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Connection); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Connection, got %T", sh)
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
		desired, err := canonicalizeConnectionDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetConnection(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Connection resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Connection resource: %v", err)
		}
		c.Config.Logger.Info("Found that Connection resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeConnectionDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Connection: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Connection: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeConnectionInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Connection: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeConnectionDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Connection: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffConnection(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeConnectionInitialState(rawInitial, rawDesired *Connection) (*Connection, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeConnectionDesiredState(rawDesired, rawInitial *Connection, opts ...dcl.ApplyOption) (*Connection, error) {

	if dcl.IsZeroValue(rawDesired.Service) {
		rawDesired.Service = dcl.String("services/servicenetworking.googleapis.com")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.ReservedPeeringRanges) {
		rawDesired.ReservedPeeringRanges = rawInitial.ReservedPeeringRanges
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Service, rawInitial.Service) {
		rawDesired.Service = rawInitial.Service
	}

	return rawDesired, nil
}

func canonicalizeConnectionNewState(c *Client, rawNew, rawDesired *Connection) (*Connection, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReservedPeeringRanges) && dcl.IsEmptyValueIndirect(rawDesired.ReservedPeeringRanges) {
		rawNew.ReservedPeeringRanges = rawDesired.ReservedPeeringRanges
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Service) && dcl.IsEmptyValueIndirect(rawDesired.Service) {
		rawNew.Service = rawDesired.Service
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Service, rawNew.Service) {
			rawNew.Service = rawDesired.Service
		}
	}

	return rawNew, nil
}

type connectionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         connectionApiOperation
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
func diffConnection(c *Client, desired, actual *Connection, opts ...dcl.ApplyOption) ([]connectionDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []connectionDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.Network, actual.Network, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "ReferenceType"}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, connectionDiff{
			UpdateOp: &updateConnectionPatchOperation{}, FieldName: "Network",
		})
	}

	if d, err := dcl.Diff(desired.Service, actual.Service, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, connectionDiff{
			UpdateOp: &updateConnectionPatchOperation{}, FieldName: "Service",
		})
	}

	if !dcl.IsZeroValue(desired.Network) && !dcl.PartialSelfLinkToSelfLink(desired.Network, actual.Network) {
		c.Config.Logger.Infof("Detected diff in Network.\nDESIRED: %v\nACTUAL: %v", desired.Network, actual.Network)

		diffs = append(diffs, connectionDiff{
			UpdateOp:  &updateConnectionPatchOperation{},
			FieldName: "Network",
		})

	}
	if !dcl.StringSliceEqualsWithSelfLink(desired.ReservedPeeringRanges, actual.ReservedPeeringRanges) {
		c.Config.Logger.Infof("Detected diff in ReservedPeeringRanges.\nDESIRED: %v\nACTUAL: %v", desired.ReservedPeeringRanges, actual.ReservedPeeringRanges)

		diffs = append(diffs, connectionDiff{
			UpdateOp:  &updateConnectionPatchOperation{},
			FieldName: "ReservedPeeringRanges",
		})

	}
	if !dcl.IsZeroValue(desired.Service) && !dcl.PartialSelfLinkToSelfLink(desired.Service, actual.Service) {
		c.Config.Logger.Infof("Detected diff in Service.\nDESIRED: %v\nACTUAL: %v", desired.Service, actual.Service)

		diffs = append(diffs, connectionDiff{
			UpdateOp:  &updateConnectionPatchOperation{},
			FieldName: "Service",
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
	var deduped []connectionDiff
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Connection) urlNormalized() *Connection {
	normalized := deepcopy.Copy(*r).(Connection)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Service = dcl.SelfLinkToName(r.Service)
	return &normalized
}

func (r *Connection) getFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Network), dcl.ValueOrEmptyString(n.Service), dcl.ValueOrEmptyString(n.Name)
}

func (r *Connection) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Service)
}

func (r *Connection) deleteFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Service)
}

func (r *Connection) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "patch" {
		fields := map[string]interface{}{
			"service": dcl.ValueOrEmptyString(n.Service),
		}
		return dcl.URL("services/{{service}}/connections/-?force=true", "https://servicenetworking.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Connection resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Connection) marshal(c *Client) ([]byte, error) {
	m, err := expandConnection(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Connection: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalConnection decodes JSON responses into the Connection resource schema.
func unmarshalConnection(b []byte, c *Client) (*Connection, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapConnection(m, c)
}

func unmarshalMapConnection(m map[string]interface{}, c *Client) (*Connection, error) {

	return flattenConnection(c, m), nil
}

// expandConnection expands Connection into a JSON request object.
func expandConnection(c *Client, f *Connection) (map[string]interface{}, error) {
	m := make(map[string]interface{})
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
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["peering"] = v
	}
	if v := f.ReservedPeeringRanges; !dcl.IsEmptyValueIndirect(v) {
		m["reservedPeeringRanges"] = v
	}
	if v, err := dcl.DeriveField("services/%s", f.Service, f.Service); err != nil {
		return nil, fmt.Errorf("error expanding Service into service: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}

	return m, nil
}

// flattenConnection flattens Connection from a JSON request object into the
// Connection type.
func flattenConnection(c *Client, i interface{}) *Connection {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Connection{}
	r.Network = dcl.FlattenString(m["network"])
	r.Project = dcl.FlattenString(m["project"])
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["peering"]))
	r.ReservedPeeringRanges = dcl.FlattenStringSlice(m["reservedPeeringRanges"])
	r.Service = dcl.FlattenString(m["service"])
	if _, ok := m["service"]; !ok {
		c.Config.Logger.Info("Using default value for service")
		r.Service = dcl.String("services/servicenetworking.googleapis.com")
	}

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Connection) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalConnection(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Network == nil && ncr.Network == nil {
			c.Config.Logger.Info("Both Network fields null - considering equal.")
		} else if nr.Network == nil || ncr.Network == nil {
			c.Config.Logger.Info("Only one Network field is null - considering unequal.")
			return false
		} else if *nr.Network != *ncr.Network {
			return false
		}
		if nr.Service == nil && ncr.Service == nil {
			c.Config.Logger.Info("Both Service fields null - considering equal.")
		} else if nr.Service == nil || ncr.Service == nil {
			c.Config.Logger.Info("Only one Service field is null - considering unequal.")
			return false
		} else if *nr.Service != *ncr.Service {
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
