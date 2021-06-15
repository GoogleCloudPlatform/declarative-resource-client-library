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
package apigee

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

func (r *Attachment) validate() error {

	if err := dcl.Required(r, "environment"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Envgroup, "Envgroup"); err != nil {
		return err
	}
	return nil
}

func attachmentGetURL(userBasePath string, r *Attachment) (string, error) {
	params := map[string]interface{}{
		"envgroup": dcl.ValueOrEmptyString(r.Envgroup),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("{{envgroup}}/attachments/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, params), nil
}

func attachmentListURL(userBasePath, envgroup string) (string, error) {
	params := map[string]interface{}{
		"envgroup": envgroup,
	}
	return dcl.URL("{{envgroup}}/attachments", "https://apigee.googleapis.com/v1/", userBasePath, params), nil

}

func attachmentCreateURL(userBasePath, envgroup string) (string, error) {
	params := map[string]interface{}{
		"envgroup": envgroup,
	}
	return dcl.URL("{{envgroup}}/attachments", "https://apigee.googleapis.com/v1/", userBasePath, params), nil

}

func attachmentDeleteURL(userBasePath string, r *Attachment) (string, error) {
	params := map[string]interface{}{
		"envgroup": dcl.ValueOrEmptyString(r.Envgroup),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("{{envgroup}}/attachments/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, params), nil
}

// attachmentApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type attachmentApiOperation interface {
	do(context.Context, *Attachment, *Client) error
}

func (c *Client) listAttachmentRaw(ctx context.Context, envgroup, pageToken string, pageSize int32) ([]byte, error) {
	u, err := attachmentListURL(c.Config.BasePath, envgroup)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AttachmentMaxPage {
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

type listAttachmentOperation struct {
	EnvironmentGroupAttachments []map[string]interface{} `json:"environmentGroupAttachments"`
	Token                       string                   `json:"nextPageToken"`
}

func (c *Client) listAttachment(ctx context.Context, envgroup, pageToken string, pageSize int32) ([]*Attachment, string, error) {
	b, err := c.listAttachmentRaw(ctx, envgroup, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAttachmentOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Attachment
	for _, v := range m.EnvironmentGroupAttachments {
		res, err := unmarshalMapAttachment(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Envgroup = &envgroup
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAttachment(ctx context.Context, f func(*Attachment) bool, resources []*Attachment) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAttachment(ctx, res)
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

type deleteAttachmentOperation struct{}

func (op *deleteAttachmentOperation) do(ctx context.Context, r *Attachment, c *Client) error {

	_, err := c.GetAttachment(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Attachment not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAttachment checking for existence. error: %v", err)
		return err
	}

	u, err := attachmentDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://apigee.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAttachment(ctx, r.urlNormalized())
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
type createAttachmentOperation struct {
	response map[string]interface{}
}

func (op *createAttachmentOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAttachmentOperation) do(ctx context.Context, r *Attachment, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	envgroup := r.createFields()
	u, err := attachmentCreateURL(c.Config.BasePath, envgroup)

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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://apigee.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string, was %T", name)
	}
	r.Name = &name

	if _, err := c.GetAttachment(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAttachmentRaw(ctx context.Context, r *Attachment) ([]byte, error) {

	u, err := attachmentGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) attachmentDiffsForRawDesired(ctx context.Context, rawDesired *Attachment, opts ...dcl.ApplyOption) (initial, desired *Attachment, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Attachment
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Attachment); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Attachment, got %T", sh)
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
		desired, err := canonicalizeAttachmentDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAttachment(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Attachment resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Attachment resource: %v", err)
		}
		c.Config.Logger.Info("Found that Attachment resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAttachmentDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Attachment: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Attachment: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAttachmentInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Attachment: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAttachmentDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Attachment: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAttachment(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAttachmentInitialState(rawInitial, rawDesired *Attachment) (*Attachment, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAttachmentDesiredState(rawDesired, rawInitial *Attachment, opts ...dcl.ApplyOption) (*Attachment, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Environment, rawInitial.Environment) {
		rawDesired.Environment = rawInitial.Environment
	}
	if dcl.NameToSelfLink(rawDesired.Envgroup, rawInitial.Envgroup) {
		rawDesired.Envgroup = rawInitial.Envgroup
	}

	return rawDesired, nil
}

func canonicalizeAttachmentNewState(c *Client, rawNew, rawDesired *Attachment) (*Attachment, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Environment) && dcl.IsEmptyValueIndirect(rawDesired.Environment) {
		rawNew.Environment = rawDesired.Environment
	} else {
		if dcl.NameToSelfLink(rawDesired.Environment, rawNew.Environment) {
			rawNew.Environment = rawDesired.Environment
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreatedAt) && dcl.IsEmptyValueIndirect(rawDesired.CreatedAt) {
		rawNew.CreatedAt = rawDesired.CreatedAt
	} else {
	}

	rawNew.Envgroup = rawDesired.Envgroup

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffAttachment(c *Client, desired, actual *Attachment, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Environment, actual.Environment, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Environment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreatedAt, actual.CreatedAt, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreatedAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Envgroup, actual.Envgroup, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Envgroup")); len(ds) != 0 || err != nil {
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
func (r *Attachment) urlNormalized() *Attachment {
	normalized := dcl.Copy(*r).(Attachment)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Environment = dcl.SelfLinkToName(r.Environment)
	normalized.Envgroup = r.Envgroup
	return &normalized
}

func (r *Attachment) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Envgroup), dcl.ValueOrEmptyString(n.Name)
}

func (r *Attachment) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Envgroup)
}

func (r *Attachment) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Envgroup), dcl.ValueOrEmptyString(n.Name)
}

func (r *Attachment) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Attachment resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Attachment) marshal(c *Client) ([]byte, error) {
	m, err := expandAttachment(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Attachment: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAttachment decodes JSON responses into the Attachment resource schema.
func unmarshalAttachment(b []byte, c *Client) (*Attachment, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAttachment(m, c)
}

func unmarshalMapAttachment(m map[string]interface{}, c *Client) (*Attachment, error) {

	return flattenAttachment(c, m), nil
}

// expandAttachment expands Attachment into a JSON request object.
func expandAttachment(c *Client, f *Attachment) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Environment; !dcl.IsEmptyValueIndirect(v) {
		m["environment"] = v
	}
	if v := f.CreatedAt; !dcl.IsEmptyValueIndirect(v) {
		m["createdAt"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Envgroup into envgroup: %w", err)
	} else if v != nil {
		m["envgroup"] = v
	}

	return m, nil
}

// flattenAttachment flattens Attachment from a JSON request object into the
// Attachment type.
func flattenAttachment(c *Client, i interface{}) *Attachment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Attachment{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.Environment = dcl.FlattenString(m["environment"])
	res.CreatedAt = dcl.FlattenInteger(m["createdAt"])
	res.Envgroup = dcl.FlattenString(m["envgroup"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Attachment) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAttachment(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Envgroup == nil && ncr.Envgroup == nil {
			c.Config.Logger.Info("Both Envgroup fields null - considering equal.")
		} else if nr.Envgroup == nil || ncr.Envgroup == nil {
			c.Config.Logger.Info("Only one Envgroup field is null - considering unequal.")
			return false
		} else if *nr.Envgroup != *ncr.Envgroup {
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

type attachmentDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         attachmentApiOperation
}

func convertFieldDiffToAttachmentOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]attachmentDiff, error) {
	var diffs []attachmentDiff
	for _, op := range ops {
		diff := attachmentDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToattachmentApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToattachmentApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (attachmentApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
