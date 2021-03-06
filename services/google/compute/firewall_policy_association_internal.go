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
package compute

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *FirewallPolicyAssociation) validate() error {

	if err := dcl.Required(r, "firewallPolicy"); err != nil {
		return err
	}
	return nil
}

func firewallPolicyAssociationGetURL(userBasePath string, r *FirewallPolicyAssociation) (string, error) {
	params := map[string]interface{}{
		"firewallPolicy": dcl.ValueOrEmptyString(r.FirewallPolicy),
		"name":           dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("locations/global/firewallPolicies/{{firewallPolicy}}/getAssociation?name={{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func firewallPolicyAssociationListURL(userBasePath, firewallPolicy string) (string, error) {
	params := map[string]interface{}{
		"firewallPolicy": firewallPolicy,
	}
	return dcl.URL("locations/global/firewallPolicies/{{firewallPolicy}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func firewallPolicyAssociationCreateURL(userBasePath, firewallPolicy string) (string, error) {
	params := map[string]interface{}{
		"firewallPolicy": firewallPolicy,
	}
	return dcl.URL("locations/global/firewallPolicies/{{firewallPolicy}}/addAssociation", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func firewallPolicyAssociationDeleteURL(userBasePath string, r *FirewallPolicyAssociation) (string, error) {
	params := map[string]interface{}{
		"firewallPolicy": dcl.ValueOrEmptyString(r.FirewallPolicy),
		"name":           dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("locations/global/firewallPolicies/{{firewallPolicy}}/removeAssociation?name={{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// firewallPolicyAssociationApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type firewallPolicyAssociationApiOperation interface {
	do(context.Context, *FirewallPolicyAssociation, *Client) error
}

func (c *Client) listFirewallPolicyAssociationRaw(ctx context.Context, firewallPolicy, pageToken string, pageSize int32) ([]byte, error) {
	u, err := firewallPolicyAssociationListURL(c.Config.BasePath, firewallPolicy)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FirewallPolicyAssociationMaxPage {
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

type listFirewallPolicyAssociationOperation struct {
	Associations []map[string]interface{} `json:"associations"`
	Token        string                   `json:"nextPageToken"`
}

func (c *Client) listFirewallPolicyAssociation(ctx context.Context, firewallPolicy, pageToken string, pageSize int32) ([]*FirewallPolicyAssociation, string, error) {
	b, err := c.listFirewallPolicyAssociationRaw(ctx, firewallPolicy, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFirewallPolicyAssociationOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*FirewallPolicyAssociation
	for _, v := range m.Associations {
		res, err := unmarshalMapFirewallPolicyAssociation(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.FirewallPolicy = &firewallPolicy
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFirewallPolicyAssociation(ctx context.Context, f func(*FirewallPolicyAssociation) bool, resources []*FirewallPolicyAssociation) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFirewallPolicyAssociation(ctx, res)
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

type deleteFirewallPolicyAssociationOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createFirewallPolicyAssociationOperation struct {
	response map[string]interface{}
}

func (op *createFirewallPolicyAssociationOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getFirewallPolicyAssociationRaw(ctx context.Context, r *FirewallPolicyAssociation) ([]byte, error) {

	u, err := firewallPolicyAssociationGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) firewallPolicyAssociationDiffsForRawDesired(ctx context.Context, rawDesired *FirewallPolicyAssociation, opts ...dcl.ApplyOption) (initial, desired *FirewallPolicyAssociation, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *FirewallPolicyAssociation
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*FirewallPolicyAssociation); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected FirewallPolicyAssociation, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFirewallPolicyAssociation(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 400) {
			c.Config.Logger.Warningf("Failed to retrieve whether a FirewallPolicyAssociation resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve FirewallPolicyAssociation resource: %v", err)
		}
		c.Config.Logger.Info("Found that FirewallPolicyAssociation resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFirewallPolicyAssociationDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for FirewallPolicyAssociation: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for FirewallPolicyAssociation: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFirewallPolicyAssociationInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for FirewallPolicyAssociation: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFirewallPolicyAssociationDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for FirewallPolicyAssociation: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFirewallPolicyAssociation(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFirewallPolicyAssociationInitialState(rawInitial, rawDesired *FirewallPolicyAssociation) (*FirewallPolicyAssociation, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFirewallPolicyAssociationDesiredState(rawDesired, rawInitial *FirewallPolicyAssociation, opts ...dcl.ApplyOption) (*FirewallPolicyAssociation, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &FirewallPolicyAssociation{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.NameToSelfLink(rawDesired.AttachmentTarget, rawInitial.AttachmentTarget) {
		canonicalDesired.AttachmentTarget = rawInitial.AttachmentTarget
	} else {
		canonicalDesired.AttachmentTarget = rawDesired.AttachmentTarget
	}
	if dcl.NameToSelfLink(rawDesired.FirewallPolicy, rawInitial.FirewallPolicy) {
		canonicalDesired.FirewallPolicy = rawInitial.FirewallPolicy
	} else {
		canonicalDesired.FirewallPolicy = rawDesired.FirewallPolicy
	}

	return canonicalDesired, nil
}

func canonicalizeFirewallPolicyAssociationNewState(c *Client, rawNew, rawDesired *FirewallPolicyAssociation) (*FirewallPolicyAssociation, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AttachmentTarget) && dcl.IsEmptyValueIndirect(rawDesired.AttachmentTarget) {
		rawNew.AttachmentTarget = rawDesired.AttachmentTarget
	} else {
		if dcl.NameToSelfLink(rawDesired.AttachmentTarget, rawNew.AttachmentTarget) {
			rawNew.AttachmentTarget = rawDesired.AttachmentTarget
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.FirewallPolicy) && dcl.IsEmptyValueIndirect(rawDesired.FirewallPolicy) {
		rawNew.FirewallPolicy = rawDesired.FirewallPolicy
	} else {
		if dcl.NameToSelfLink(rawDesired.FirewallPolicy, rawNew.FirewallPolicy) {
			rawNew.FirewallPolicy = rawDesired.FirewallPolicy
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ShortName) && dcl.IsEmptyValueIndirect(rawDesired.ShortName) {
		rawNew.ShortName = rawDesired.ShortName
	} else {
		if dcl.StringCanonicalize(rawDesired.ShortName, rawNew.ShortName) {
			rawNew.ShortName = rawDesired.ShortName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffFirewallPolicyAssociation(c *Client, desired, actual *FirewallPolicyAssociation, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.AttachmentTarget, actual.AttachmentTarget, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AttachmentTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FirewallPolicy, actual.FirewallPolicy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FirewallPolicyId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ShortName, actual.ShortName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ShortName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *FirewallPolicyAssociation) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.FirewallPolicy), dcl.ValueOrEmptyString(n.Name)
}

func (r *FirewallPolicyAssociation) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.FirewallPolicy)
}

func (r *FirewallPolicyAssociation) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.FirewallPolicy), dcl.ValueOrEmptyString(n.Name)
}

func (r *FirewallPolicyAssociation) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the FirewallPolicyAssociation resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *FirewallPolicyAssociation) marshal(c *Client) ([]byte, error) {
	m, err := expandFirewallPolicyAssociation(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling FirewallPolicyAssociation: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFirewallPolicyAssociation decodes JSON responses into the FirewallPolicyAssociation resource schema.
func unmarshalFirewallPolicyAssociation(b []byte, c *Client) (*FirewallPolicyAssociation, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFirewallPolicyAssociation(m, c)
}

func unmarshalMapFirewallPolicyAssociation(m map[string]interface{}, c *Client) (*FirewallPolicyAssociation, error) {

	return flattenFirewallPolicyAssociation(c, m), nil
}

// expandFirewallPolicyAssociation expands FirewallPolicyAssociation into a JSON request object.
func expandFirewallPolicyAssociation(c *Client, f *FirewallPolicyAssociation) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.AttachmentTarget; !dcl.IsEmptyValueIndirect(v) {
		m["attachmentTarget"] = v
	}
	if v := f.FirewallPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["firewallPolicyId"] = v
	}
	if v := f.ShortName; !dcl.IsEmptyValueIndirect(v) {
		m["shortName"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}

	return m, nil
}

// flattenFirewallPolicyAssociation flattens FirewallPolicyAssociation from a JSON request object into the
// FirewallPolicyAssociation type.
func flattenFirewallPolicyAssociation(c *Client, i interface{}) *FirewallPolicyAssociation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &FirewallPolicyAssociation{}
	res.Name = dcl.FlattenString(m["name"])
	res.AttachmentTarget = dcl.FlattenString(m["attachmentTarget"])
	res.FirewallPolicy = dcl.FlattenString(m["firewallPolicyId"])
	res.ShortName = dcl.FlattenString(m["shortName"])
	res.DisplayName = dcl.FlattenString(m["displayName"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *FirewallPolicyAssociation) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFirewallPolicyAssociation(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.FirewallPolicy == nil && ncr.FirewallPolicy == nil {
			c.Config.Logger.Info("Both FirewallPolicy fields null - considering equal.")
		} else if nr.FirewallPolicy == nil || ncr.FirewallPolicy == nil {
			c.Config.Logger.Info("Only one FirewallPolicy field is null - considering unequal.")
			return false
		} else if *nr.FirewallPolicy != *ncr.FirewallPolicy {
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

type firewallPolicyAssociationDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         firewallPolicyAssociationApiOperation
}

func convertFieldDiffsToFirewallPolicyAssociationDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]firewallPolicyAssociationDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []firewallPolicyAssociationDiff
	// For each operation name, create a firewallPolicyAssociationDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := firewallPolicyAssociationDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToFirewallPolicyAssociationApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToFirewallPolicyAssociationApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (firewallPolicyAssociationApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
