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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *NetworkEndpoint) validate() error {

	return nil
}

func networkEndpointGetURL(userBasePath string, r *NetworkEndpoint) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"group":    dcl.ValueOrEmptyString(r.Group),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/networkEndpointGroups/{{group}}/listNetworkEndpoints", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/networkEndpointGroups/{{group}}/listNetworkEndpoints", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Get URL found")

}

func networkEndpointListURL(userBasePath, project, location, group string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"group":    group,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/networkEndpointGroups/{{group}}/listNetworkEndpoints", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/networkEndpointGroups/{{group}}/listNetworkEndpoints", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid List URL found")

}

func networkEndpointCreateURL(userBasePath, project, location, group string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"group":    group,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/networkEndpointGroups/{{group}}/attachNetworkEndpoints", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/networkEndpointGroups/{{group}}/attachNetworkEndpoints", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Create URL found")

}

func networkEndpointDeleteURL(userBasePath string, r *NetworkEndpoint) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"group":    dcl.ValueOrEmptyString(r.Group),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/networkEndpointGroups/{{group}}/detachNetworkEndpoints", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/networkEndpointGroups/{{group}}/detachNetworkEndpoints", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Delete URL found")

}

// networkEndpointApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type networkEndpointApiOperation interface {
	do(context.Context, *NetworkEndpoint, *Client) error
}

func (c *Client) listNetworkEndpointRaw(ctx context.Context, project, location, group, pageToken string, pageSize int32) ([]byte, error) {
	u, err := networkEndpointListURL(c.Config.BasePath, project, location, group)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != NetworkEndpointMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listNetworkEndpointOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listNetworkEndpoint(ctx context.Context, project, location, group, pageToken string, pageSize int32) ([]*NetworkEndpoint, string, error) {
	b, err := c.listNetworkEndpointRaw(ctx, project, location, group, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listNetworkEndpointOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*NetworkEndpoint
	for _, v := range m.Items {
		res := flattenNetworkEndpoint(c, v)
		res.Project = &project
		res.Location = &location
		res.Group = &group
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllNetworkEndpoint(ctx context.Context, f func(*NetworkEndpoint) bool, resources []*NetworkEndpoint) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteNetworkEndpoint(ctx, res)
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

type deleteNetworkEndpointOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createNetworkEndpointOperation struct {
	response map[string]interface{}
}

func (op *createNetworkEndpointOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createNetworkEndpointOperation) do(ctx context.Context, r *NetworkEndpoint, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, group := r.createFields()
	u, err := networkEndpointCreateURL(c.Config.BasePath, project, location, group)

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

	if _, err := c.GetNetworkEndpoint(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getNetworkEndpointRaw(ctx context.Context, r *NetworkEndpoint) ([]byte, error) {

	u, err := networkEndpointGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, &bytes.Buffer{}, c.Config.RetryProvider)
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

func (c *Client) networkEndpointDiffsForRawDesired(ctx context.Context, rawDesired *NetworkEndpoint, opts ...dcl.ApplyOption) (initial, desired *NetworkEndpoint, diffs []networkEndpointDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *NetworkEndpoint
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*NetworkEndpoint); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected NetworkEndpoint, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetNetworkEndpoint(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a NetworkEndpoint resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve NetworkEndpoint resource: %v", err)
		}
		c.Config.Logger.Info("Found that NetworkEndpoint resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeNetworkEndpointDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for NetworkEndpoint: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for NetworkEndpoint: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeNetworkEndpointInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for NetworkEndpoint: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeNetworkEndpointDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for NetworkEndpoint: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffNetworkEndpoint(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeNetworkEndpointInitialState(rawInitial, rawDesired *NetworkEndpoint) (*NetworkEndpoint, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeNetworkEndpointDesiredState(rawDesired, rawInitial *NetworkEndpoint, opts ...dcl.ApplyOption) (*NetworkEndpoint, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Port) {
		rawDesired.Port = rawInitial.Port
	}
	if dcl.StringCanonicalize(rawDesired.IPAddress, rawInitial.IPAddress) {
		rawDesired.IPAddress = rawInitial.IPAddress
	}
	if dcl.StringCanonicalize(rawDesired.Fqdn, rawInitial.Fqdn) {
		rawDesired.Fqdn = rawInitial.Fqdn
	}
	if dcl.NameToSelfLink(rawDesired.Instance, rawInitial.Instance) {
		rawDesired.Instance = rawInitial.Instance
	}
	if dcl.IsZeroValue(rawDesired.Annotations) {
		rawDesired.Annotations = rawInitial.Annotations
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.NameToSelfLink(rawDesired.Group, rawInitial.Group) {
		rawDesired.Group = rawInitial.Group
	}

	return rawDesired, nil
}

func canonicalizeNetworkEndpointNewState(c *Client, rawNew, rawDesired *NetworkEndpoint) (*NetworkEndpoint, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Port) && dcl.IsEmptyValueIndirect(rawDesired.Port) {
		rawNew.Port = rawDesired.Port
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPAddress) && dcl.IsEmptyValueIndirect(rawDesired.IPAddress) {
		rawNew.IPAddress = rawDesired.IPAddress
	} else {
		if dcl.StringCanonicalize(rawDesired.IPAddress, rawNew.IPAddress) {
			rawNew.IPAddress = rawDesired.IPAddress
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Fqdn) && dcl.IsEmptyValueIndirect(rawDesired.Fqdn) {
		rawNew.Fqdn = rawDesired.Fqdn
	} else {
		if dcl.StringCanonicalize(rawDesired.Fqdn, rawNew.Fqdn) {
			rawNew.Fqdn = rawDesired.Fqdn
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Instance) && dcl.IsEmptyValueIndirect(rawDesired.Instance) {
		rawNew.Instance = rawDesired.Instance
	} else {
		if dcl.NameToSelfLink(rawDesired.Instance, rawNew.Instance) {
			rawNew.Instance = rawDesired.Instance
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Annotations) && dcl.IsEmptyValueIndirect(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.Group = rawDesired.Group

	return rawNew, nil
}

type networkEndpointDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         networkEndpointApiOperation
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
func diffNetworkEndpoint(c *Client, desired, actual *NetworkEndpoint, opts ...dcl.ApplyOption) ([]networkEndpointDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []networkEndpointDiff

	var fn dcl.FieldName

	// New style diffs.
	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Port",
		})
	}

	if ds, err := dcl.Diff(desired.IPAddress, actual.IPAddress, dcl.Info{}, fn.AddNest("IPAddress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "IPAddress",
		})
	}

	if ds, err := dcl.Diff(desired.Fqdn, actual.Fqdn, dcl.Info{}, fn.AddNest("Fqdn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Fqdn",
		})
	}

	if ds, err := dcl.Diff(desired.Instance, actual.Instance, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Instance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Instance",
		})
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.Info{}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Annotations",
		})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Project",
		})
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Location",
		})
	}

	if ds, err := dcl.Diff(desired.Group, actual.Group, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Group")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Group",
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
	var deduped []networkEndpointDiff
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
func (r *NetworkEndpoint) urlNormalized() *NetworkEndpoint {
	normalized := dcl.Copy(*r).(NetworkEndpoint)
	normalized.IPAddress = dcl.SelfLinkToName(r.IPAddress)
	normalized.Fqdn = dcl.SelfLinkToName(r.Fqdn)
	normalized.Instance = dcl.SelfLinkToName(r.Instance)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Group = dcl.SelfLinkToName(r.Group)
	return &normalized
}

func (r *NetworkEndpoint) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Group)
}

func (r *NetworkEndpoint) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Group)
}

func (r *NetworkEndpoint) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Group)
}

func (r *NetworkEndpoint) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the NetworkEndpoint resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *NetworkEndpoint) marshal(c *Client) ([]byte, error) {
	m, err := expandNetworkEndpoint(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling NetworkEndpoint: %w", err)
	}
	m = encodeNetworkEndpointRequest(m)

	return json.Marshal(m)
}

// unmarshalNetworkEndpoint decodes JSON responses into the NetworkEndpoint resource schema.
func unmarshalNetworkEndpoint(b []byte, c *Client) (*NetworkEndpoint, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapNetworkEndpoint(m, c)
}

func unmarshalMapNetworkEndpoint(m map[string]interface{}, c *Client) (*NetworkEndpoint, error) {

	return flattenNetworkEndpoint(c, m), nil
}

// expandNetworkEndpoint expands NetworkEndpoint into a JSON request object.
func expandNetworkEndpoint(c *Client, f *NetworkEndpoint) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.IPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddress"] = v
	}
	if v := f.Fqdn; !dcl.IsEmptyValueIndirect(v) {
		m["fqdn"] = v
	}
	if v := f.Instance; !dcl.IsEmptyValueIndirect(v) {
		m["instance"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Group into group: %w", err)
	} else if v != nil {
		m["group"] = v
	}

	return m, nil
}

// flattenNetworkEndpoint flattens NetworkEndpoint from a JSON request object into the
// NetworkEndpoint type.
func flattenNetworkEndpoint(c *Client, i interface{}) *NetworkEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &NetworkEndpoint{}
	r.Port = dcl.FlattenInteger(m["port"])
	r.IPAddress = dcl.FlattenString(m["ipAddress"])
	r.Fqdn = dcl.FlattenString(m["fqdn"])
	r.Instance = dcl.FlattenString(m["instance"])
	r.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])
	r.Group = dcl.FlattenString(m["group"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *NetworkEndpoint) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalNetworkEndpoint(b, c)
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
		if nr.Group == nil && ncr.Group == nil {
			c.Config.Logger.Info("Both Group fields null - considering equal.")
		} else if nr.Group == nil || ncr.Group == nil {
			c.Config.Logger.Info("Only one Group field is null - considering unequal.")
			return false
		} else if *nr.Group != *ncr.Group {
			return false
		}
		return true
	}
}
