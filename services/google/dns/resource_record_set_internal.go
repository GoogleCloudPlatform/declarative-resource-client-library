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
package dns

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *ResourceRecordSet) validate() error {

	if err := dcl.Required(r, "dnsName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "dnsType"); err != nil {
		return err
	}
	if err := dcl.Required(r, "ttl"); err != nil {
		return err
	}
	if err := dcl.Required(r, "target"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.ManagedZone, "ManagedZone"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func resourceRecordSetGetURL(userBasePath string, r *ResourceRecordSet) (string, error) {
	params := map[string]interface{}{
		"project":     dcl.ValueOrEmptyString(r.Project),
		"dnsName":     dcl.ValueOrEmptyString(r.DnsName),
		"dnsType":     dcl.ValueOrEmptyString(r.DnsType),
		"managedZone": dcl.ValueOrEmptyString(r.ManagedZone),
	}
	return dcl.URL("projects/{{project}}/managedZones/{{managedZone}}/rrsets?name={{dnsName}}&type={{dnsType}}", "https://www.googleapis.com/dns/v1/", userBasePath, params), nil
}

func resourceRecordSetListURL(userBasePath, project, managedZone string) (string, error) {
	params := map[string]interface{}{
		"project":     project,
		"managedZone": managedZone,
	}
	return dcl.URL("projects/{{project}}/managedZones/{{managedZone}}/rrsets", "https://www.googleapis.com/dns/v1/", userBasePath, params), nil

}

func resourceRecordSetCreateURL(userBasePath, project, managedZone string) (string, error) {
	params := map[string]interface{}{
		"project":     project,
		"managedZone": managedZone,
	}
	return dcl.URL("projects/{{project}}/managedZones/{{managedZone}}/changes", "https://www.googleapis.com/dns/v1/", userBasePath, params), nil

}

func resourceRecordSetDeleteURL(userBasePath string, r *ResourceRecordSet) (string, error) {
	params := map[string]interface{}{
		"project":     dcl.ValueOrEmptyString(r.Project),
		"dnsName":     dcl.ValueOrEmptyString(r.DnsName),
		"dnsType":     dcl.ValueOrEmptyString(r.DnsType),
		"managedZone": dcl.ValueOrEmptyString(r.ManagedZone),
	}
	return dcl.URL("projects/{{project}}/managedZones/{{managedZone}}/rrsets?name={{dnsName}}&type={{dnsType}}", "https://www.googleapis.com/dns/v1/", userBasePath, params), nil
}

// resourceRecordSetApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type resourceRecordSetApiOperation interface {
	do(context.Context, *ResourceRecordSet, *Client) error
}

// newUpdateResourceRecordSetUpdateRequest creates a request for an
// ResourceRecordSet resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateResourceRecordSetUpdateRequest(ctx context.Context, f *ResourceRecordSet, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DnsName; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.DnsType; !dcl.IsEmptyValueIndirect(v) {
		req["type"] = v
	}
	if v := f.Ttl; !dcl.IsEmptyValueIndirect(v) {
		req["ttl"] = v
	}
	if v := f.Target; !dcl.IsEmptyValueIndirect(v) {
		req["rrdatas"] = v
	}
	return req, nil
}

// marshalUpdateResourceRecordSetUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateResourceRecordSetUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateResourceRecordSetUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listResourceRecordSetRaw(ctx context.Context, project, managedZone, pageToken string, pageSize int32) ([]byte, error) {
	u, err := resourceRecordSetListURL(c.Config.BasePath, project, managedZone)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ResourceRecordSetMaxPage {
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

type listResourceRecordSetOperation struct {
	Rrsets []map[string]interface{} `json:"rrsets"`
	Token  string                   `json:"nextPageToken"`
}

func (c *Client) listResourceRecordSet(ctx context.Context, project, managedZone, pageToken string, pageSize int32) ([]*ResourceRecordSet, string, error) {
	b, err := c.listResourceRecordSetRaw(ctx, project, managedZone, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listResourceRecordSetOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ResourceRecordSet
	for _, v := range m.Rrsets {
		res, err := unmarshalMapResourceRecordSet(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.ManagedZone = &managedZone
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllResourceRecordSet(ctx context.Context, f func(*ResourceRecordSet) bool, resources []*ResourceRecordSet) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteResourceRecordSet(ctx, res)
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

type deleteResourceRecordSetOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createResourceRecordSetOperation struct {
	response map[string]interface{}
}

func (op *createResourceRecordSetOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) resourceRecordSetDiffsForRawDesired(ctx context.Context, rawDesired *ResourceRecordSet, opts ...dcl.ApplyOption) (initial, desired *ResourceRecordSet, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ResourceRecordSet
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ResourceRecordSet); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ResourceRecordSet, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetResourceRecordSet(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ResourceRecordSet resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ResourceRecordSet resource: %v", err)
		}
		c.Config.Logger.Info("Found that ResourceRecordSet resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeResourceRecordSetDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for ResourceRecordSet: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ResourceRecordSet: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeResourceRecordSetInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ResourceRecordSet: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeResourceRecordSetDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ResourceRecordSet: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffResourceRecordSet(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeResourceRecordSetInitialState(rawInitial, rawDesired *ResourceRecordSet) (*ResourceRecordSet, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeResourceRecordSetDesiredState(rawDesired, rawInitial *ResourceRecordSet, opts ...dcl.ApplyOption) (*ResourceRecordSet, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	if dcl.StringCanonicalize(rawDesired.DnsName, rawInitial.DnsName) {
		rawDesired.DnsName = rawInitial.DnsName
	}
	if dcl.StringCanonicalize(rawDesired.DnsType, rawInitial.DnsType) {
		rawDesired.DnsType = rawInitial.DnsType
	}
	if dcl.IsZeroValue(rawDesired.Ttl) {
		rawDesired.Ttl = rawInitial.Ttl
	}
	if dcl.QuoteAndCaseInsensitiveStringArray(rawDesired.Target, rawInitial.Target) {
		rawDesired.Target = rawInitial.Target
	}
	if dcl.NameToSelfLink(rawDesired.ManagedZone, rawInitial.ManagedZone) {
		rawDesired.ManagedZone = rawInitial.ManagedZone
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeResourceRecordSetNewState(c *Client, rawNew, rawDesired *ResourceRecordSet) (*ResourceRecordSet, error) {

	if dcl.IsEmptyValueIndirect(rawNew.DnsName) && dcl.IsEmptyValueIndirect(rawDesired.DnsName) {
		rawNew.DnsName = rawDesired.DnsName
	} else {
		if dcl.StringCanonicalize(rawDesired.DnsName, rawNew.DnsName) {
			rawNew.DnsName = rawDesired.DnsName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DnsType) && dcl.IsEmptyValueIndirect(rawDesired.DnsType) {
		rawNew.DnsType = rawDesired.DnsType
	} else {
		if dcl.StringCanonicalize(rawDesired.DnsType, rawNew.DnsType) {
			rawNew.DnsType = rawDesired.DnsType
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Ttl) && dcl.IsEmptyValueIndirect(rawDesired.Ttl) {
		rawNew.Ttl = rawDesired.Ttl
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Target) && dcl.IsEmptyValueIndirect(rawDesired.Target) {
		rawNew.Target = rawDesired.Target
	} else {
		if dcl.QuoteAndCaseInsensitiveStringArray(rawDesired.Target, rawNew.Target) {
			rawNew.Target = rawDesired.Target
		}
	}

	rawNew.ManagedZone = rawDesired.ManagedZone

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffResourceRecordSet(c *Client, desired, actual *ResourceRecordSet, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.DnsName, actual.DnsName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateResourceRecordSetUpdateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DnsType, actual.DnsType, dcl.Info{OperationSelector: dcl.TriggersOperation("updateResourceRecordSetUpdateOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ttl, actual.Ttl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateResourceRecordSetUpdateOperation")}, fn.AddNest("Ttl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Target, actual.Target, dcl.Info{OperationSelector: dcl.TriggersOperation("updateResourceRecordSetUpdateOperation")}, fn.AddNest("Rrdatas")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManagedZone, actual.ManagedZone, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateResourceRecordSetUpdateOperation")}, fn.AddNest("ManagedZone")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
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
func (r *ResourceRecordSet) urlNormalized() *ResourceRecordSet {
	normalized := dcl.Copy(*r).(ResourceRecordSet)
	normalized.DnsName = dcl.SelfLinkToName(r.DnsName)
	normalized.DnsType = dcl.SelfLinkToName(r.DnsType)
	normalized.ManagedZone = dcl.SelfLinkToName(r.ManagedZone)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *ResourceRecordSet) getFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.DnsName), dcl.ValueOrEmptyString(n.DnsType), dcl.ValueOrEmptyString(n.ManagedZone)
}

func (r *ResourceRecordSet) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.ManagedZone)
}

func (r *ResourceRecordSet) deleteFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.DnsName), dcl.ValueOrEmptyString(n.DnsType), dcl.ValueOrEmptyString(n.ManagedZone)
}

func (r *ResourceRecordSet) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project":     dcl.ValueOrEmptyString(n.Project),
			"dnsName":     dcl.ValueOrEmptyString(n.DnsName),
			"managedZone": dcl.ValueOrEmptyString(n.ManagedZone),
		}
		return dcl.URL("projects/{{project}}/managedZones/{{managedZone}}/rrsets?name={{dnsName}}", "https://www.googleapis.com/dns/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ResourceRecordSet resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ResourceRecordSet) marshal(c *Client) ([]byte, error) {
	m, err := expandResourceRecordSet(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ResourceRecordSet: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalResourceRecordSet decodes JSON responses into the ResourceRecordSet resource schema.
func unmarshalResourceRecordSet(b []byte, c *Client) (*ResourceRecordSet, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapResourceRecordSet(m, c)
}

func unmarshalMapResourceRecordSet(m map[string]interface{}, c *Client) (*ResourceRecordSet, error) {

	return flattenResourceRecordSet(c, m), nil
}

// expandResourceRecordSet expands ResourceRecordSet into a JSON request object.
func expandResourceRecordSet(c *Client, f *ResourceRecordSet) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.DnsName; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DnsType; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Ttl; !dcl.IsEmptyValueIndirect(v) {
		m["ttl"] = v
	}
	if v := f.Target; !dcl.IsEmptyValueIndirect(v) {
		m["rrdatas"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding ManagedZone into managedZone: %w", err)
	} else if v != nil {
		m["managedZone"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenResourceRecordSet flattens ResourceRecordSet from a JSON request object into the
// ResourceRecordSet type.
func flattenResourceRecordSet(c *Client, i interface{}) *ResourceRecordSet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &ResourceRecordSet{}
	res.DnsName = dcl.FlattenString(m["name"])
	res.DnsType = dcl.FlattenString(m["type"])
	res.Ttl = dcl.FlattenInteger(m["ttl"])
	res.Target = dcl.FlattenStringSlice(m["rrdatas"])
	res.ManagedZone = dcl.FlattenString(m["managedZone"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ResourceRecordSet) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalResourceRecordSet(b, c)
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
		if nr.DnsName == nil && ncr.DnsName == nil {
			c.Config.Logger.Info("Both DnsName fields null - considering equal.")
		} else if nr.DnsName == nil || ncr.DnsName == nil {
			c.Config.Logger.Info("Only one DnsName field is null - considering unequal.")
			return false
		} else if *nr.DnsName != *ncr.DnsName {
			return false
		}
		if nr.DnsType == nil && ncr.DnsType == nil {
			c.Config.Logger.Info("Both DnsType fields null - considering equal.")
		} else if nr.DnsType == nil || ncr.DnsType == nil {
			c.Config.Logger.Info("Only one DnsType field is null - considering unequal.")
			return false
		} else if *nr.DnsType != *ncr.DnsType {
			return false
		}
		if nr.ManagedZone == nil && ncr.ManagedZone == nil {
			c.Config.Logger.Info("Both ManagedZone fields null - considering equal.")
		} else if nr.ManagedZone == nil || ncr.ManagedZone == nil {
			c.Config.Logger.Info("Only one ManagedZone field is null - considering unequal.")
			return false
		} else if *nr.ManagedZone != *ncr.ManagedZone {
			return false
		}
		return true
	}
}

type resourceRecordSetDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         resourceRecordSetApiOperation
}

func convertFieldDiffToResourceRecordSetOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]resourceRecordSetDiff, error) {
	var diffs []resourceRecordSetDiff
	for _, op := range ops {
		diff := resourceRecordSetDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToresourceRecordSetApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToresourceRecordSetApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (resourceRecordSetApiOperation, error) {
	switch op {

	case "updateResourceRecordSetUpdateOperation":
		return &updateResourceRecordSetUpdateOperation{ApplyOptions: opts, Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
