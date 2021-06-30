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
package iap

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

func (r *IdentityAwareProxyClient) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Brand, "Brand"); err != nil {
		return err
	}
	return nil
}

func identityAwareProxyClientGetURL(userBasePath string, r *IdentityAwareProxyClient) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"brand":   dcl.ValueOrEmptyString(r.Brand),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/brands/{{brand}}/identityAwareProxyClients/{{name}}", "https://iap.googleapis.com/v1/", userBasePath, params), nil
}

func identityAwareProxyClientListURL(userBasePath, project, brand string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"brand":   brand,
	}
	return dcl.URL("projects/{{project}}/brands/{{brand}}/identityAwareProxyClients", "https://iap.googleapis.com/v1/", userBasePath, params), nil

}

func identityAwareProxyClientCreateURL(userBasePath, project, brand string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"brand":   brand,
	}
	return dcl.URL("projects/{{project}}/brands/{{brand}}/identityAwareProxyClients", "https://iap.googleapis.com/v1/", userBasePath, params), nil

}

func identityAwareProxyClientDeleteURL(userBasePath string, r *IdentityAwareProxyClient) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"brand":   dcl.ValueOrEmptyString(r.Brand),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/brands/{{brand}}/identityAwareProxyClients/{{name}}", "https://iap.googleapis.com/v1/", userBasePath, params), nil
}

// identityAwareProxyClientApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type identityAwareProxyClientApiOperation interface {
	do(context.Context, *IdentityAwareProxyClient, *Client) error
}

func (c *Client) listIdentityAwareProxyClientRaw(ctx context.Context, project, brand, pageToken string, pageSize int32) ([]byte, error) {
	u, err := identityAwareProxyClientListURL(c.Config.BasePath, project, brand)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != IdentityAwareProxyClientMaxPage {
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

type listIdentityAwareProxyClientOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listIdentityAwareProxyClient(ctx context.Context, project, brand, pageToken string, pageSize int32) ([]*IdentityAwareProxyClient, string, error) {
	b, err := c.listIdentityAwareProxyClientRaw(ctx, project, brand, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listIdentityAwareProxyClientOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*IdentityAwareProxyClient
	for _, v := range m.Items {
		res, err := unmarshalMapIdentityAwareProxyClient(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Brand = &brand
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllIdentityAwareProxyClient(ctx context.Context, f func(*IdentityAwareProxyClient) bool, resources []*IdentityAwareProxyClient) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteIdentityAwareProxyClient(ctx, res)
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

type deleteIdentityAwareProxyClientOperation struct{}

func (op *deleteIdentityAwareProxyClientOperation) do(ctx context.Context, r *IdentityAwareProxyClient, c *Client) error {
	r, err := c.GetIdentityAwareProxyClient(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("IdentityAwareProxyClient not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetIdentityAwareProxyClient checking for existence. error: %v", err)
		return err
	}

	u, err := identityAwareProxyClientDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete IdentityAwareProxyClient: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetIdentityAwareProxyClient(ctx, r.URLNormalized())
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
type createIdentityAwareProxyClientOperation struct {
	response map[string]interface{}
}

func (op *createIdentityAwareProxyClientOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createIdentityAwareProxyClientOperation) do(ctx context.Context, r *IdentityAwareProxyClient, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, brand := r.createFields()
	u, err := identityAwareProxyClientCreateURL(c.Config.BasePath, project, brand)

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
		return fmt.Errorf("expected name to be a string, was %T", name)
	}
	r.Name = &name

	if _, err := c.GetIdentityAwareProxyClient(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getIdentityAwareProxyClientRaw(ctx context.Context, r *IdentityAwareProxyClient) ([]byte, error) {

	u, err := identityAwareProxyClientGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) identityAwareProxyClientDiffsForRawDesired(ctx context.Context, rawDesired *IdentityAwareProxyClient, opts ...dcl.ApplyOption) (initial, desired *IdentityAwareProxyClient, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *IdentityAwareProxyClient
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*IdentityAwareProxyClient); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected IdentityAwareProxyClient, got %T", sh)
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
		desired, err := canonicalizeIdentityAwareProxyClientDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetIdentityAwareProxyClient(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a IdentityAwareProxyClient resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve IdentityAwareProxyClient resource: %v", err)
		}
		c.Config.Logger.Info("Found that IdentityAwareProxyClient resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeIdentityAwareProxyClientDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for IdentityAwareProxyClient: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for IdentityAwareProxyClient: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeIdentityAwareProxyClientInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for IdentityAwareProxyClient: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeIdentityAwareProxyClientDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for IdentityAwareProxyClient: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffIdentityAwareProxyClient(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeIdentityAwareProxyClientInitialState(rawInitial, rawDesired *IdentityAwareProxyClient) (*IdentityAwareProxyClient, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeIdentityAwareProxyClientDesiredState(rawDesired, rawInitial *IdentityAwareProxyClient, opts ...dcl.ApplyOption) (*IdentityAwareProxyClient, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Brand, rawInitial.Brand) {
		rawDesired.Brand = rawInitial.Brand
	}

	return rawDesired, nil
}

func canonicalizeIdentityAwareProxyClientNewState(c *Client, rawNew, rawDesired *IdentityAwareProxyClient) (*IdentityAwareProxyClient, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Secret) && dcl.IsEmptyValueIndirect(rawDesired.Secret) {
		rawNew.Secret = rawDesired.Secret
	} else {
		if dcl.StringCanonicalize(rawDesired.Secret, rawNew.Secret) {
			rawNew.Secret = rawDesired.Secret
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Brand = rawDesired.Brand

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffIdentityAwareProxyClient(c *Client, desired, actual *IdentityAwareProxyClient, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Brand, actual.Brand, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Brand")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *IdentityAwareProxyClient) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Brand), dcl.ValueOrEmptyString(n.Name)
}

func (r *IdentityAwareProxyClient) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Brand)
}

func (r *IdentityAwareProxyClient) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Brand), dcl.ValueOrEmptyString(n.Name)
}

func (r *IdentityAwareProxyClient) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the IdentityAwareProxyClient resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *IdentityAwareProxyClient) marshal(c *Client) ([]byte, error) {
	m, err := expandIdentityAwareProxyClient(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling IdentityAwareProxyClient: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalIdentityAwareProxyClient decodes JSON responses into the IdentityAwareProxyClient resource schema.
func unmarshalIdentityAwareProxyClient(b []byte, c *Client) (*IdentityAwareProxyClient, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapIdentityAwareProxyClient(m, c)
}

func unmarshalMapIdentityAwareProxyClient(m map[string]interface{}, c *Client) (*IdentityAwareProxyClient, error) {

	return flattenIdentityAwareProxyClient(c, m), nil
}

// expandIdentityAwareProxyClient expands IdentityAwareProxyClient into a JSON request object.
func expandIdentityAwareProxyClient(c *Client, f *IdentityAwareProxyClient) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/brands/%s/identityAwareProxyClients/%s", f.Name, f.Project, f.Brand, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Secret; !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Brand into brand: %w", err)
	} else if v != nil {
		m["brand"] = v
	}

	return m, nil
}

// flattenIdentityAwareProxyClient flattens IdentityAwareProxyClient from a JSON request object into the
// IdentityAwareProxyClient type.
func flattenIdentityAwareProxyClient(c *Client, i interface{}) *IdentityAwareProxyClient {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &IdentityAwareProxyClient{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.Secret = dcl.FlattenString(m["secret"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.Project = dcl.FlattenString(m["project"])
	res.Brand = dcl.FlattenString(m["brand"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *IdentityAwareProxyClient) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalIdentityAwareProxyClient(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Brand == nil && ncr.Brand == nil {
			c.Config.Logger.Info("Both Brand fields null - considering equal.")
		} else if nr.Brand == nil || ncr.Brand == nil {
			c.Config.Logger.Info("Only one Brand field is null - considering unequal.")
			return false
		} else if *nr.Brand != *ncr.Brand {
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

type identityAwareProxyClientDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         identityAwareProxyClientApiOperation
}

func convertFieldDiffToIdentityAwareProxyClientOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]identityAwareProxyClientDiff, error) {
	var diffs []identityAwareProxyClientDiff
	for _, op := range ops {
		diff := identityAwareProxyClientDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToidentityAwareProxyClientApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToidentityAwareProxyClientApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (identityAwareProxyClientApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
