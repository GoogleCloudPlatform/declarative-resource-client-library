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
package binaryauthorization

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Attestor) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.UserOwnedGrafeasNote) {
		if err := r.UserOwnedGrafeasNote.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AttestorUserOwnedGrafeasNote) validate() error {
	if err := dcl.Required(r, "noteReference"); err != nil {
		return err
	}
	return nil
}
func (r *AttestorUserOwnedGrafeasNotePublicKeys) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PkixPublicKey) {
		if err := r.PkixPublicKey.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) validate() error {
	return nil
}

func attestorGetURL(userBasePath string, r *Attestor) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/attestors/{{name}}", "https://binaryauthorization.googleapis.com/v1", userBasePath, params), nil
}

func attestorListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/attestors", "https://binaryauthorization.googleapis.com/v1", userBasePath, params), nil

}

func attestorCreateURL(userBasePath, project, name string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"name":    name,
	}
	return dcl.URL("projects/{{project}}/attestors?attestorId={{name}}", "https://binaryauthorization.googleapis.com/v1", userBasePath, params), nil

}

func attestorDeleteURL(userBasePath string, r *Attestor) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/attestors/{{name}}", "https://binaryauthorization.googleapis.com/v1", userBasePath, params), nil
}

// attestorApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type attestorApiOperation interface {
	do(context.Context, *Attestor, *Client) error
}

// newUpdateAttestorUpdateAttestorRequest creates a request for an
// Attestor resource's UpdateAttestor update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAttestorUpdateAttestorRequest(ctx context.Context, f *Attestor, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateAttestorUpdateAttestorRequest converts the update into
// the final JSON request body.
func marshalUpdateAttestorUpdateAttestorRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAttestorUpdateAttestorOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAttestorUpdateAttestorOperation) do(ctx context.Context, r *Attestor, c *Client) error {
	_, err := c.GetAttestor(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateAttestor")
	if err != nil {
		return err
	}

	req, err := newUpdateAttestorUpdateAttestorRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateAttestorUpdateAttestorRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://binaryauthorization.googleapis.com/v1", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listAttestorRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := attestorListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AttestorMaxPage {
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

type listAttestorOperation struct {
	Attestors []map[string]interface{} `json:"attestors"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listAttestor(ctx context.Context, project, pageToken string, pageSize int32) ([]*Attestor, string, error) {
	b, err := c.listAttestorRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAttestorOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Attestor
	for _, v := range m.Attestors {
		res := flattenAttestor(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAttestor(ctx context.Context, f func(*Attestor) bool, resources []*Attestor) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAttestor(ctx, res)
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

type deleteAttestorOperation struct{}

func (op *deleteAttestorOperation) do(ctx context.Context, r *Attestor, c *Client) error {

	_, err := c.GetAttestor(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Attestor not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAttestor checking for existence. error: %v", err)
		return err
	}

	u, err := attestorDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Attestor: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAttestor(ctx, r.urlNormalized())
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
type createAttestorOperation struct {
	response map[string]interface{}
}

func (op *createAttestorOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAttestorOperation) do(ctx context.Context, r *Attestor, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, name := r.createFields()
	u, err := attestorCreateURL(c.Config.BasePath, project, name)

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

	if _, err := c.GetAttestor(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAttestorRaw(ctx context.Context, r *Attestor) ([]byte, error) {

	u, err := attestorGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) attestorDiffsForRawDesired(ctx context.Context, rawDesired *Attestor, opts ...dcl.ApplyOption) (initial, desired *Attestor, diffs []attestorDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Attestor
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Attestor); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Attestor, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAttestor(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Attestor resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Attestor resource: %v", err)
		}
		c.Config.Logger.Info("Found that Attestor resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAttestorDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Attestor: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Attestor: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAttestorInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Attestor: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAttestorDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Attestor: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAttestor(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAttestorInitialState(rawInitial, rawDesired *Attestor) (*Attestor, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAttestorDesiredState(rawDesired, rawInitial *Attestor, opts ...dcl.ApplyOption) (*Attestor, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.UserOwnedGrafeasNote = canonicalizeAttestorUserOwnedGrafeasNote(rawDesired.UserOwnedGrafeasNote, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	rawDesired.UserOwnedGrafeasNote = canonicalizeAttestorUserOwnedGrafeasNote(rawDesired.UserOwnedGrafeasNote, rawInitial.UserOwnedGrafeasNote, opts...)
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeAttestorNewState(c *Client, rawNew, rawDesired *Attestor) (*Attestor, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.UserOwnedGrafeasNote) && dcl.IsEmptyValueIndirect(rawDesired.UserOwnedGrafeasNote) {
		rawNew.UserOwnedGrafeasNote = rawDesired.UserOwnedGrafeasNote
	} else {
		rawNew.UserOwnedGrafeasNote = canonicalizeNewAttestorUserOwnedGrafeasNote(c, rawDesired.UserOwnedGrafeasNote, rawNew.UserOwnedGrafeasNote)
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeAttestorUserOwnedGrafeasNote(des, initial *AttestorUserOwnedGrafeasNote, opts ...dcl.ApplyOption) *AttestorUserOwnedGrafeasNote {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.NoteReference, initial.NoteReference) || dcl.IsZeroValue(des.NoteReference) {
		des.NoteReference = initial.NoteReference
	}
	if dcl.IsZeroValue(des.PublicKeys) {
		des.PublicKeys = initial.PublicKeys
	}
	if dcl.StringCanonicalize(des.DelegationServiceAccountEmail, initial.DelegationServiceAccountEmail) || dcl.IsZeroValue(des.DelegationServiceAccountEmail) {
		des.DelegationServiceAccountEmail = initial.DelegationServiceAccountEmail
	}

	return des
}

func canonicalizeNewAttestorUserOwnedGrafeasNote(c *Client, des, nw *AttestorUserOwnedGrafeasNote) *AttestorUserOwnedGrafeasNote {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.NoteReference, nw.NoteReference) {
		nw.NoteReference = des.NoteReference
	}
	nw.PublicKeys = canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeysSlice(c, des.PublicKeys, nw.PublicKeys)
	if dcl.StringCanonicalize(des.DelegationServiceAccountEmail, nw.DelegationServiceAccountEmail) {
		nw.DelegationServiceAccountEmail = des.DelegationServiceAccountEmail
	}

	return nw
}

func canonicalizeNewAttestorUserOwnedGrafeasNoteSet(c *Client, des, nw []AttestorUserOwnedGrafeasNote) []AttestorUserOwnedGrafeasNote {
	if des == nil {
		return nw
	}
	var reorderedNew []AttestorUserOwnedGrafeasNote
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAttestorUserOwnedGrafeasNote(c, &d, &n) {
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

func canonicalizeNewAttestorUserOwnedGrafeasNoteSlice(c *Client, des, nw []AttestorUserOwnedGrafeasNote) []AttestorUserOwnedGrafeasNote {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AttestorUserOwnedGrafeasNote
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAttestorUserOwnedGrafeasNote(c, &d, &n))
	}

	return items
}

func canonicalizeAttestorUserOwnedGrafeasNotePublicKeys(des, initial *AttestorUserOwnedGrafeasNotePublicKeys, opts ...dcl.ApplyOption) *AttestorUserOwnedGrafeasNotePublicKeys {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Comment, initial.Comment) || dcl.IsZeroValue(des.Comment) {
		des.Comment = initial.Comment
	}
	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		des.Id = initial.Id
	}
	if dcl.StringCanonicalize(des.AsciiArmoredPgpPublicKey, initial.AsciiArmoredPgpPublicKey) || dcl.IsZeroValue(des.AsciiArmoredPgpPublicKey) {
		des.AsciiArmoredPgpPublicKey = initial.AsciiArmoredPgpPublicKey
	}
	des.PkixPublicKey = canonicalizeAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(des.PkixPublicKey, initial.PkixPublicKey, opts...)

	return des
}

func canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeys(c *Client, des, nw *AttestorUserOwnedGrafeasNotePublicKeys) *AttestorUserOwnedGrafeasNotePublicKeys {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Comment, nw.Comment) {
		nw.Comment = des.Comment
	}
	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.AsciiArmoredPgpPublicKey, nw.AsciiArmoredPgpPublicKey) {
		nw.AsciiArmoredPgpPublicKey = des.AsciiArmoredPgpPublicKey
	}
	nw.PkixPublicKey = canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, des.PkixPublicKey, nw.PkixPublicKey)

	return nw
}

func canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeysSet(c *Client, des, nw []AttestorUserOwnedGrafeasNotePublicKeys) []AttestorUserOwnedGrafeasNotePublicKeys {
	if des == nil {
		return nw
	}
	var reorderedNew []AttestorUserOwnedGrafeasNotePublicKeys
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAttestorUserOwnedGrafeasNotePublicKeys(c, &d, &n) {
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

func canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeysSlice(c *Client, des, nw []AttestorUserOwnedGrafeasNotePublicKeys) []AttestorUserOwnedGrafeasNotePublicKeys {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AttestorUserOwnedGrafeasNotePublicKeys
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeys(c, &d, &n))
	}

	return items
}

func canonicalizeAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(des, initial *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey, opts ...dcl.ApplyOption) *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.PublicKeyPem, initial.PublicKeyPem) || dcl.IsZeroValue(des.PublicKeyPem) {
		des.PublicKeyPem = initial.PublicKeyPem
	}
	if dcl.IsZeroValue(des.SignatureAlgorithm) {
		des.SignatureAlgorithm = initial.SignatureAlgorithm
	}

	return des
}

func canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c *Client, des, nw *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PublicKeyPem, nw.PublicKeyPem) {
		nw.PublicKeyPem = des.PublicKeyPem
	}

	return nw
}

func canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySet(c *Client, des, nw []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey {
	if des == nil {
		return nw
	}
	var reorderedNew []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, &d, &n) {
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

func canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySlice(c *Client, des, nw []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, &d, &n))
	}

	return items
}

type attestorDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         attestorApiOperation
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
func diffAttestor(c *Client, desired, actual *Attestor, opts ...dcl.ApplyOption) ([]attestorDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []attestorDiff

	var fn dcl.FieldName

	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, attestorDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Name",
		})
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, attestorDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Description",
		})
	}

	if ds, err := dcl.Diff(desired.UserOwnedGrafeasNote, actual.UserOwnedGrafeasNote, dcl.Info{ObjectFunction: compareAttestorUserOwnedGrafeasNoteNewStyle}, fn.AddNest("UserOwnedGrafeasNote")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, attestorDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "UserOwnedGrafeasNote",
		})
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, attestorDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "UpdateTime",
		})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, attestorDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Project",
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
	var deduped []attestorDiff
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
func compareAttestorUserOwnedGrafeasNoteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AttestorUserOwnedGrafeasNote)
	if !ok {
		desiredNotPointer, ok := d.(AttestorUserOwnedGrafeasNote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedGrafeasNote or *AttestorUserOwnedGrafeasNote", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AttestorUserOwnedGrafeasNote)
	if !ok {
		actualNotPointer, ok := a.(AttestorUserOwnedGrafeasNote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedGrafeasNote", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NoteReference, actual.NoteReference, dcl.Info{Type: "ReferenceType"}, fn.AddNest("NoteReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublicKeys, actual.PublicKeys, dcl.Info{ObjectFunction: compareAttestorUserOwnedGrafeasNotePublicKeysNewStyle}, fn.AddNest("PublicKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DelegationServiceAccountEmail, actual.DelegationServiceAccountEmail, dcl.Info{OutputOnly: true}, fn.AddNest("DelegationServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAttestorUserOwnedGrafeasNote(c *Client, desired, actual *AttestorUserOwnedGrafeasNote) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.NameToSelfLink(desired.NoteReference, actual.NoteReference) && !dcl.IsZeroValue(desired.NoteReference) {
		c.Config.Logger.Infof("Diff in NoteReference.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NoteReference), dcl.SprintResource(actual.NoteReference))
		return true
	}
	if compareAttestorUserOwnedGrafeasNotePublicKeysSlice(c, desired.PublicKeys, actual.PublicKeys) && !dcl.IsZeroValue(desired.PublicKeys) {
		c.Config.Logger.Infof("Diff in PublicKeys.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PublicKeys), dcl.SprintResource(actual.PublicKeys))
		return true
	}
	return false
}

func compareAttestorUserOwnedGrafeasNoteSlice(c *Client, desired, actual []AttestorUserOwnedGrafeasNote) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AttestorUserOwnedGrafeasNote, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAttestorUserOwnedGrafeasNote(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNote, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAttestorUserOwnedGrafeasNoteMap(c *Client, desired, actual map[string]AttestorUserOwnedGrafeasNote) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AttestorUserOwnedGrafeasNote, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNote, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAttestorUserOwnedGrafeasNote(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNote, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAttestorUserOwnedGrafeasNotePublicKeysNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AttestorUserOwnedGrafeasNotePublicKeys)
	if !ok {
		desiredNotPointer, ok := d.(AttestorUserOwnedGrafeasNotePublicKeys)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedGrafeasNotePublicKeys or *AttestorUserOwnedGrafeasNotePublicKeys", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AttestorUserOwnedGrafeasNotePublicKeys)
	if !ok {
		actualNotPointer, ok := a.(AttestorUserOwnedGrafeasNotePublicKeys)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedGrafeasNotePublicKeys", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Comment, actual.Comment, dcl.Info{}, fn.AddNest("Comment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AsciiArmoredPgpPublicKey, actual.AsciiArmoredPgpPublicKey, dcl.Info{}, fn.AddNest("AsciiArmoredPgpPublicKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PkixPublicKey, actual.PkixPublicKey, dcl.Info{ObjectFunction: compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyNewStyle}, fn.AddNest("PkixPublicKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAttestorUserOwnedGrafeasNotePublicKeys(c *Client, desired, actual *AttestorUserOwnedGrafeasNotePublicKeys) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Comment, actual.Comment) && !dcl.IsZeroValue(desired.Comment) {
		c.Config.Logger.Infof("Diff in Comment.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Comment), dcl.SprintResource(actual.Comment))
		return true
	}
	if !dcl.StringCanonicalize(desired.Id, actual.Id) && !dcl.IsZeroValue(desired.Id) {
		c.Config.Logger.Infof("Diff in Id.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Id), dcl.SprintResource(actual.Id))
		return true
	}
	if !dcl.StringCanonicalize(desired.AsciiArmoredPgpPublicKey, actual.AsciiArmoredPgpPublicKey) && !dcl.IsZeroValue(desired.AsciiArmoredPgpPublicKey) {
		c.Config.Logger.Infof("Diff in AsciiArmoredPgpPublicKey.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AsciiArmoredPgpPublicKey), dcl.SprintResource(actual.AsciiArmoredPgpPublicKey))
		return true
	}
	if compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, desired.PkixPublicKey, actual.PkixPublicKey) && !dcl.IsZeroValue(desired.PkixPublicKey) {
		c.Config.Logger.Infof("Diff in PkixPublicKey.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PkixPublicKey), dcl.SprintResource(actual.PkixPublicKey))
		return true
	}
	return false
}

func compareAttestorUserOwnedGrafeasNotePublicKeysSlice(c *Client, desired, actual []AttestorUserOwnedGrafeasNotePublicKeys) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AttestorUserOwnedGrafeasNotePublicKeys, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAttestorUserOwnedGrafeasNotePublicKeys(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNotePublicKeys, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAttestorUserOwnedGrafeasNotePublicKeysMap(c *Client, desired, actual map[string]AttestorUserOwnedGrafeasNotePublicKeys) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AttestorUserOwnedGrafeasNotePublicKeys, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNotePublicKeys, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAttestorUserOwnedGrafeasNotePublicKeys(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNotePublicKeys, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey)
	if !ok {
		desiredNotPointer, ok := d.(AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey or *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey)
	if !ok {
		actualNotPointer, ok := a.(AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PublicKeyPem, actual.PublicKeyPem, dcl.Info{}, fn.AddNest("PublicKeyPem")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SignatureAlgorithm, actual.SignatureAlgorithm, dcl.Info{Type: "EnumType"}, fn.AddNest("SignatureAlgorithm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c *Client, desired, actual *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.PublicKeyPem, actual.PublicKeyPem) && !dcl.IsZeroValue(desired.PublicKeyPem) {
		c.Config.Logger.Infof("Diff in PublicKeyPem.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PublicKeyPem), dcl.SprintResource(actual.PublicKeyPem))
		return true
	}
	if !reflect.DeepEqual(desired.SignatureAlgorithm, actual.SignatureAlgorithm) && !dcl.IsZeroValue(desired.SignatureAlgorithm) {
		c.Config.Logger.Infof("Diff in SignatureAlgorithm.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SignatureAlgorithm), dcl.SprintResource(actual.SignatureAlgorithm))
		return true
	}
	return false
}

func compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySlice(c *Client, desired, actual []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyMap(c *Client, desired, actual map[string]AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnumSlice(c *Client, desired, actual []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(c *Client, desired, actual *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Attestor) urlNormalized() *Attestor {
	normalized := dcl.Copy(*r).(Attestor)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Attestor) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Attestor) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Attestor) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Attestor) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateAttestor" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/attestors/{{name}}", "https://binaryauthorization.googleapis.com/v1", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Attestor resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Attestor) marshal(c *Client) ([]byte, error) {
	m, err := expandAttestor(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Attestor: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAttestor decodes JSON responses into the Attestor resource schema.
func unmarshalAttestor(b []byte, c *Client) (*Attestor, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAttestor(m, c)
}

func unmarshalMapAttestor(m map[string]interface{}, c *Client) (*Attestor, error) {

	return flattenAttestor(c, m), nil
}

// expandAttestor expands Attestor into a JSON request object.
func expandAttestor(c *Client, f *Attestor) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/attestors/%s", f.Name, f.Project, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandAttestorUserOwnedGrafeasNote(c, f.UserOwnedGrafeasNote); err != nil {
		return nil, fmt.Errorf("error expanding UserOwnedGrafeasNote into userOwnedGrafeasNote: %w", err)
	} else if v != nil {
		m["userOwnedGrafeasNote"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenAttestor flattens Attestor from a JSON request object into the
// Attestor type.
func flattenAttestor(c *Client, i interface{}) *Attestor {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Attestor{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.UserOwnedGrafeasNote = flattenAttestorUserOwnedGrafeasNote(c, m["userOwnedGrafeasNote"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandAttestorUserOwnedGrafeasNoteMap expands the contents of AttestorUserOwnedGrafeasNote into a JSON
// request object.
func expandAttestorUserOwnedGrafeasNoteMap(c *Client, f map[string]AttestorUserOwnedGrafeasNote) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAttestorUserOwnedGrafeasNote(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAttestorUserOwnedGrafeasNoteSlice expands the contents of AttestorUserOwnedGrafeasNote into a JSON
// request object.
func expandAttestorUserOwnedGrafeasNoteSlice(c *Client, f []AttestorUserOwnedGrafeasNote) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAttestorUserOwnedGrafeasNote(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAttestorUserOwnedGrafeasNoteMap flattens the contents of AttestorUserOwnedGrafeasNote from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNoteMap(c *Client, i interface{}) map[string]AttestorUserOwnedGrafeasNote {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedGrafeasNote{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedGrafeasNote{}
	}

	items := make(map[string]AttestorUserOwnedGrafeasNote)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedGrafeasNote(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAttestorUserOwnedGrafeasNoteSlice flattens the contents of AttestorUserOwnedGrafeasNote from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNoteSlice(c *Client, i interface{}) []AttestorUserOwnedGrafeasNote {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedGrafeasNote{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedGrafeasNote{}
	}

	items := make([]AttestorUserOwnedGrafeasNote, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedGrafeasNote(c, item.(map[string]interface{})))
	}

	return items
}

// expandAttestorUserOwnedGrafeasNote expands an instance of AttestorUserOwnedGrafeasNote into a JSON
// request object.
func expandAttestorUserOwnedGrafeasNote(c *Client, f *AttestorUserOwnedGrafeasNote) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.NoteReference; !dcl.IsEmptyValueIndirect(v) {
		m["noteReference"] = v
	}
	if v, err := expandAttestorUserOwnedGrafeasNotePublicKeysSlice(c, f.PublicKeys); err != nil {
		return nil, fmt.Errorf("error expanding PublicKeys into publicKeys: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["publicKeys"] = v
	}
	if v := f.DelegationServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["delegationServiceAccountEmail"] = v
	}

	return m, nil
}

// flattenAttestorUserOwnedGrafeasNote flattens an instance of AttestorUserOwnedGrafeasNote from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNote(c *Client, i interface{}) *AttestorUserOwnedGrafeasNote {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AttestorUserOwnedGrafeasNote{}
	r.NoteReference = dcl.FlattenString(m["noteReference"])
	r.PublicKeys = flattenAttestorUserOwnedGrafeasNotePublicKeysSlice(c, m["publicKeys"])
	r.DelegationServiceAccountEmail = dcl.FlattenString(m["delegationServiceAccountEmail"])

	return r
}

// expandAttestorUserOwnedGrafeasNotePublicKeysMap expands the contents of AttestorUserOwnedGrafeasNotePublicKeys into a JSON
// request object.
func expandAttestorUserOwnedGrafeasNotePublicKeysMap(c *Client, f map[string]AttestorUserOwnedGrafeasNotePublicKeys) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAttestorUserOwnedGrafeasNotePublicKeys(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAttestorUserOwnedGrafeasNotePublicKeysSlice expands the contents of AttestorUserOwnedGrafeasNotePublicKeys into a JSON
// request object.
func expandAttestorUserOwnedGrafeasNotePublicKeysSlice(c *Client, f []AttestorUserOwnedGrafeasNotePublicKeys) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAttestorUserOwnedGrafeasNotePublicKeys(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAttestorUserOwnedGrafeasNotePublicKeysMap flattens the contents of AttestorUserOwnedGrafeasNotePublicKeys from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNotePublicKeysMap(c *Client, i interface{}) map[string]AttestorUserOwnedGrafeasNotePublicKeys {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedGrafeasNotePublicKeys{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedGrafeasNotePublicKeys{}
	}

	items := make(map[string]AttestorUserOwnedGrafeasNotePublicKeys)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedGrafeasNotePublicKeys(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAttestorUserOwnedGrafeasNotePublicKeysSlice flattens the contents of AttestorUserOwnedGrafeasNotePublicKeys from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNotePublicKeysSlice(c *Client, i interface{}) []AttestorUserOwnedGrafeasNotePublicKeys {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedGrafeasNotePublicKeys{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedGrafeasNotePublicKeys{}
	}

	items := make([]AttestorUserOwnedGrafeasNotePublicKeys, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedGrafeasNotePublicKeys(c, item.(map[string]interface{})))
	}

	return items
}

// expandAttestorUserOwnedGrafeasNotePublicKeys expands an instance of AttestorUserOwnedGrafeasNotePublicKeys into a JSON
// request object.
func expandAttestorUserOwnedGrafeasNotePublicKeys(c *Client, f *AttestorUserOwnedGrafeasNotePublicKeys) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Comment; !dcl.IsEmptyValueIndirect(v) {
		m["comment"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.AsciiArmoredPgpPublicKey; !dcl.IsEmptyValueIndirect(v) {
		m["asciiArmoredPgpPublicKey"] = v
	}
	if v, err := expandAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, f.PkixPublicKey); err != nil {
		return nil, fmt.Errorf("error expanding PkixPublicKey into pkixPublicKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pkixPublicKey"] = v
	}

	return m, nil
}

// flattenAttestorUserOwnedGrafeasNotePublicKeys flattens an instance of AttestorUserOwnedGrafeasNotePublicKeys from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNotePublicKeys(c *Client, i interface{}) *AttestorUserOwnedGrafeasNotePublicKeys {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AttestorUserOwnedGrafeasNotePublicKeys{}
	r.Comment = dcl.FlattenString(m["comment"])
	r.Id = dcl.FlattenString(m["id"])
	r.AsciiArmoredPgpPublicKey = dcl.FlattenString(m["asciiArmoredPgpPublicKey"])
	r.PkixPublicKey = flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, m["pkixPublicKey"])

	return r
}

// expandAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyMap expands the contents of AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey into a JSON
// request object.
func expandAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyMap(c *Client, f map[string]AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySlice expands the contents of AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey into a JSON
// request object.
func expandAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySlice(c *Client, f []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyMap flattens the contents of AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyMap(c *Client, i interface{}) map[string]AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey{}
	}

	items := make(map[string]AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySlice flattens the contents of AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySlice(c *Client, i interface{}) []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey{}
	}

	items := make([]AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey expands an instance of AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey into a JSON
// request object.
func expandAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c *Client, f *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.PublicKeyPem; !dcl.IsEmptyValueIndirect(v) {
		m["publicKeyPem"] = v
	}
	if v := f.SignatureAlgorithm; !dcl.IsEmptyValueIndirect(v) {
		m["signatureAlgorithm"] = v
	}

	return m, nil
}

// flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey flattens an instance of AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(c *Client, i interface{}) *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey{}
	r.PublicKeyPem = dcl.FlattenString(m["publicKeyPem"])
	r.SignatureAlgorithm = flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(m["signatureAlgorithm"])

	return r
}

// flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnumSlice flattens the contents of AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum from a JSON
// response object.
func flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnumSlice(c *Client, i interface{}) []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum{}
	}

	items := make([]AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(item.(interface{})))
	}

	return items
}

// flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum asserts that an interface is a string, and returns a
// pointer to a *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum with the same value as that string.
func flattenAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(i interface{}) *AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	s, ok := i.(string)
	if !ok {
		return AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnumRef("")
	}

	return AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Attestor) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAttestor(b, c)
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
