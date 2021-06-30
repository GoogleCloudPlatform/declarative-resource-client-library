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
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Attestor) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.UserOwnedDrydockNote) {
		if err := r.UserOwnedDrydockNote.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AttestorUserOwnedDrydockNote) validate() error {
	if err := dcl.Required(r, "noteReference"); err != nil {
		return err
	}
	return nil
}
func (r *AttestorUserOwnedDrydockNotePublicKeys) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PkixPublicKey) {
		if err := r.PkixPublicKey.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) validate() error {
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

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandAttestorUserOwnedDrydockNote(c, f.UserOwnedDrydockNote); err != nil {
		return nil, fmt.Errorf("error expanding UserOwnedDrydockNote into userOwnedGrafeasNote: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["userOwnedGrafeasNote"] = v
	}
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
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAttestorUpdateAttestorOperation) do(ctx context.Context, r *Attestor, c *Client) error {
	_, err := c.GetAttestor(ctx, r.URLNormalized())
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
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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
		res, err := unmarshalMapAttestor(v, c)
		if err != nil {
			return nil, m.Token, err
		}
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
	r, err := c.GetAttestor(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Attestor not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAttestor checking for existence. error: %v", err)
		return err
	}

	u, err := attestorDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetAttestor(ctx, r.URLNormalized())
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

	if _, err := c.GetAttestor(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAttestorRaw(ctx context.Context, r *Attestor) ([]byte, error) {

	u, err := attestorGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) attestorDiffsForRawDesired(ctx context.Context, rawDesired *Attestor, opts ...dcl.ApplyOption) (initial, desired *Attestor, diffs []*dcl.FieldDiff, err error) {
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
	rawInitial, err := c.GetAttestor(ctx, fetchState.URLNormalized())
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
		rawDesired.UserOwnedDrydockNote = canonicalizeAttestorUserOwnedDrydockNote(rawDesired.UserOwnedDrydockNote, nil, opts...)

		return rawDesired, nil
	}

	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	rawDesired.UserOwnedDrydockNote = canonicalizeAttestorUserOwnedDrydockNote(rawDesired.UserOwnedDrydockNote, rawInitial.UserOwnedDrydockNote, opts...)
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

	if dcl.IsEmptyValueIndirect(rawNew.UserOwnedDrydockNote) && dcl.IsEmptyValueIndirect(rawDesired.UserOwnedDrydockNote) {
		rawNew.UserOwnedDrydockNote = rawDesired.UserOwnedDrydockNote
	} else {
		rawNew.UserOwnedDrydockNote = canonicalizeNewAttestorUserOwnedDrydockNote(c, rawDesired.UserOwnedDrydockNote, rawNew.UserOwnedDrydockNote)
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeAttestorUserOwnedDrydockNote(des, initial *AttestorUserOwnedDrydockNote, opts ...dcl.ApplyOption) *AttestorUserOwnedDrydockNote {
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

	return des
}

func canonicalizeNewAttestorUserOwnedDrydockNote(c *Client, des, nw *AttestorUserOwnedDrydockNote) *AttestorUserOwnedDrydockNote {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.NoteReference, nw.NoteReference) {
		nw.NoteReference = des.NoteReference
	}
	nw.PublicKeys = canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysSlice(c, des.PublicKeys, nw.PublicKeys)
	if dcl.StringCanonicalize(des.DelegationServiceAccountEmail, nw.DelegationServiceAccountEmail) {
		nw.DelegationServiceAccountEmail = des.DelegationServiceAccountEmail
	}

	return nw
}

func canonicalizeNewAttestorUserOwnedDrydockNoteSet(c *Client, des, nw []AttestorUserOwnedDrydockNote) []AttestorUserOwnedDrydockNote {
	if des == nil {
		return nw
	}
	var reorderedNew []AttestorUserOwnedDrydockNote
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAttestorUserOwnedDrydockNoteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAttestorUserOwnedDrydockNoteSlice(c *Client, des, nw []AttestorUserOwnedDrydockNote) []AttestorUserOwnedDrydockNote {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AttestorUserOwnedDrydockNote
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAttestorUserOwnedDrydockNote(c, &d, &n))
	}

	return items
}

func canonicalizeAttestorUserOwnedDrydockNotePublicKeys(des, initial *AttestorUserOwnedDrydockNotePublicKeys, opts ...dcl.ApplyOption) *AttestorUserOwnedDrydockNotePublicKeys {
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
	des.PkixPublicKey = canonicalizeAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(des.PkixPublicKey, initial.PkixPublicKey, opts...)

	return des
}

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeys(c *Client, des, nw *AttestorUserOwnedDrydockNotePublicKeys) *AttestorUserOwnedDrydockNotePublicKeys {
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
	nw.PkixPublicKey = canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, des.PkixPublicKey, nw.PkixPublicKey)

	return nw
}

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysSet(c *Client, des, nw []AttestorUserOwnedDrydockNotePublicKeys) []AttestorUserOwnedDrydockNotePublicKeys {
	if des == nil {
		return nw
	}
	var reorderedNew []AttestorUserOwnedDrydockNotePublicKeys
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAttestorUserOwnedDrydockNotePublicKeysNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysSlice(c *Client, des, nw []AttestorUserOwnedDrydockNotePublicKeys) []AttestorUserOwnedDrydockNotePublicKeys {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AttestorUserOwnedDrydockNotePublicKeys
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAttestorUserOwnedDrydockNotePublicKeys(c, &d, &n))
	}

	return items
}

func canonicalizeAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(des, initial *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, opts ...dcl.ApplyOption) *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
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

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c *Client, des, nw *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.PublicKeyPem, nw.PublicKeyPem) {
		nw.PublicKeyPem = des.PublicKeyPem
	}
	if dcl.IsZeroValue(nw.SignatureAlgorithm) {
		nw.SignatureAlgorithm = des.SignatureAlgorithm
	}

	return nw
}

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySet(c *Client, des, nw []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if des == nil {
		return nw
	}
	var reorderedNew []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice(c *Client, des, nw []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffAttestor(c *Client, desired, actual *Attestor, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserOwnedDrydockNote, actual.UserOwnedDrydockNote, dcl.Info{ObjectFunction: compareAttestorUserOwnedDrydockNoteNewStyle, EmptyObject: EmptyAttestorUserOwnedDrydockNote, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UserOwnedGrafeasNote")); len(ds) != 0 || err != nil {
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

	return newDiffs, nil
}
func compareAttestorUserOwnedDrydockNoteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AttestorUserOwnedDrydockNote)
	if !ok {
		desiredNotPointer, ok := d.(AttestorUserOwnedDrydockNote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNote or *AttestorUserOwnedDrydockNote", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AttestorUserOwnedDrydockNote)
	if !ok {
		actualNotPointer, ok := a.(AttestorUserOwnedDrydockNote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNote", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NoteReference, actual.NoteReference, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NoteReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublicKeys, actual.PublicKeys, dcl.Info{ObjectFunction: compareAttestorUserOwnedDrydockNotePublicKeysNewStyle, EmptyObject: EmptyAttestorUserOwnedDrydockNotePublicKeys, OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("PublicKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DelegationServiceAccountEmail, actual.DelegationServiceAccountEmail, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DelegationServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAttestorUserOwnedDrydockNotePublicKeysNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AttestorUserOwnedDrydockNotePublicKeys)
	if !ok {
		desiredNotPointer, ok := d.(AttestorUserOwnedDrydockNotePublicKeys)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNotePublicKeys or *AttestorUserOwnedDrydockNotePublicKeys", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AttestorUserOwnedDrydockNotePublicKeys)
	if !ok {
		actualNotPointer, ok := a.(AttestorUserOwnedDrydockNotePublicKeys)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNotePublicKeys", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Comment, actual.Comment, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("Comment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AsciiArmoredPgpPublicKey, actual.AsciiArmoredPgpPublicKey, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("AsciiArmoredPgpPublicKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PkixPublicKey, actual.PkixPublicKey, dcl.Info{ObjectFunction: compareAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyNewStyle, EmptyObject: EmptyAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("PkixPublicKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
	if !ok {
		desiredNotPointer, ok := d.(AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey or *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
	if !ok {
		actualNotPointer, ok := a.(AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PublicKeyPem, actual.PublicKeyPem, dcl.Info{OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("PublicKeyPem")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SignatureAlgorithm, actual.SignatureAlgorithm, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("SignatureAlgorithm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Attestor) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Attestor) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Attestor) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Attestor) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
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
	if v, err := expandAttestorUserOwnedDrydockNote(c, f.UserOwnedDrydockNote); err != nil {
		return nil, fmt.Errorf("error expanding UserOwnedDrydockNote into userOwnedGrafeasNote: %w", err)
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

	res := &Attestor{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.UserOwnedDrydockNote = flattenAttestorUserOwnedDrydockNote(c, m["userOwnedGrafeasNote"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandAttestorUserOwnedDrydockNoteMap expands the contents of AttestorUserOwnedDrydockNote into a JSON
// request object.
func expandAttestorUserOwnedDrydockNoteMap(c *Client, f map[string]AttestorUserOwnedDrydockNote) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAttestorUserOwnedDrydockNote(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAttestorUserOwnedDrydockNoteSlice expands the contents of AttestorUserOwnedDrydockNote into a JSON
// request object.
func expandAttestorUserOwnedDrydockNoteSlice(c *Client, f []AttestorUserOwnedDrydockNote) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAttestorUserOwnedDrydockNote(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAttestorUserOwnedDrydockNoteMap flattens the contents of AttestorUserOwnedDrydockNote from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNoteMap(c *Client, i interface{}) map[string]AttestorUserOwnedDrydockNote {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedDrydockNote{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedDrydockNote{}
	}

	items := make(map[string]AttestorUserOwnedDrydockNote)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedDrydockNote(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAttestorUserOwnedDrydockNoteSlice flattens the contents of AttestorUserOwnedDrydockNote from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNoteSlice(c *Client, i interface{}) []AttestorUserOwnedDrydockNote {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedDrydockNote{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedDrydockNote{}
	}

	items := make([]AttestorUserOwnedDrydockNote, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedDrydockNote(c, item.(map[string]interface{})))
	}

	return items
}

// expandAttestorUserOwnedDrydockNote expands an instance of AttestorUserOwnedDrydockNote into a JSON
// request object.
func expandAttestorUserOwnedDrydockNote(c *Client, f *AttestorUserOwnedDrydockNote) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NoteReference; !dcl.IsEmptyValueIndirect(v) {
		m["noteReference"] = v
	}
	if v, err := expandAttestorUserOwnedDrydockNotePublicKeysSlice(c, f.PublicKeys); err != nil {
		return nil, fmt.Errorf("error expanding PublicKeys into publicKeys: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["publicKeys"] = v
	}
	if v := f.DelegationServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["delegationServiceAccountEmail"] = v
	}

	return m, nil
}

// flattenAttestorUserOwnedDrydockNote flattens an instance of AttestorUserOwnedDrydockNote from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNote(c *Client, i interface{}) *AttestorUserOwnedDrydockNote {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AttestorUserOwnedDrydockNote{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAttestorUserOwnedDrydockNote
	}
	r.NoteReference = dcl.FlattenString(m["noteReference"])
	r.PublicKeys = flattenAttestorUserOwnedDrydockNotePublicKeysSlice(c, m["publicKeys"])
	r.DelegationServiceAccountEmail = dcl.FlattenString(m["delegationServiceAccountEmail"])

	return r
}

// expandAttestorUserOwnedDrydockNotePublicKeysMap expands the contents of AttestorUserOwnedDrydockNotePublicKeys into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysMap(c *Client, f map[string]AttestorUserOwnedDrydockNotePublicKeys) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAttestorUserOwnedDrydockNotePublicKeys(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAttestorUserOwnedDrydockNotePublicKeysSlice expands the contents of AttestorUserOwnedDrydockNotePublicKeys into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysSlice(c *Client, f []AttestorUserOwnedDrydockNotePublicKeys) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAttestorUserOwnedDrydockNotePublicKeys(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAttestorUserOwnedDrydockNotePublicKeysMap flattens the contents of AttestorUserOwnedDrydockNotePublicKeys from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysMap(c *Client, i interface{}) map[string]AttestorUserOwnedDrydockNotePublicKeys {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedDrydockNotePublicKeys{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedDrydockNotePublicKeys{}
	}

	items := make(map[string]AttestorUserOwnedDrydockNotePublicKeys)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedDrydockNotePublicKeys(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAttestorUserOwnedDrydockNotePublicKeysSlice flattens the contents of AttestorUserOwnedDrydockNotePublicKeys from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysSlice(c *Client, i interface{}) []AttestorUserOwnedDrydockNotePublicKeys {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedDrydockNotePublicKeys{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedDrydockNotePublicKeys{}
	}

	items := make([]AttestorUserOwnedDrydockNotePublicKeys, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedDrydockNotePublicKeys(c, item.(map[string]interface{})))
	}

	return items
}

// expandAttestorUserOwnedDrydockNotePublicKeys expands an instance of AttestorUserOwnedDrydockNotePublicKeys into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeys(c *Client, f *AttestorUserOwnedDrydockNotePublicKeys) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Comment; !dcl.IsEmptyValueIndirect(v) {
		m["comment"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.AsciiArmoredPgpPublicKey; !dcl.IsEmptyValueIndirect(v) {
		m["asciiArmoredPgpPublicKey"] = v
	}
	if v, err := expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, f.PkixPublicKey); err != nil {
		return nil, fmt.Errorf("error expanding PkixPublicKey into pkixPublicKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pkixPublicKey"] = v
	}

	return m, nil
}

// flattenAttestorUserOwnedDrydockNotePublicKeys flattens an instance of AttestorUserOwnedDrydockNotePublicKeys from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeys(c *Client, i interface{}) *AttestorUserOwnedDrydockNotePublicKeys {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AttestorUserOwnedDrydockNotePublicKeys{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAttestorUserOwnedDrydockNotePublicKeys
	}
	r.Comment = dcl.FlattenString(m["comment"])
	r.Id = dcl.FlattenString(m["id"])
	r.AsciiArmoredPgpPublicKey = dcl.FlattenString(m["asciiArmoredPgpPublicKey"])
	r.PkixPublicKey = flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, m["pkixPublicKey"])

	return r
}

// expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyMap expands the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyMap(c *Client, f map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice expands the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice(c *Client, f []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyMap flattens the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyMap(c *Client, i interface{}) map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}

	items := make(map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice flattens the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice(c *Client, i interface{}) []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}

	items := make([]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey expands an instance of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c *Client, f *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PublicKeyPem; !dcl.IsEmptyValueIndirect(v) {
		m["publicKeyPem"] = v
	}
	if v := f.SignatureAlgorithm; !dcl.IsEmptyValueIndirect(v) {
		m["signatureAlgorithm"] = v
	}

	return m, nil
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey flattens an instance of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c *Client, i interface{}) *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey
	}
	r.PublicKeyPem = dcl.FlattenString(m["publicKeyPem"])
	r.SignatureAlgorithm = flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(m["signatureAlgorithm"])

	return r
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumSlice flattens the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumSlice(c *Client, i interface{}) []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum{}
	}

	items := make([]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(item.(interface{})))
	}

	return items
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum asserts that an interface is a string, and returns a
// pointer to a *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum with the same value as that string.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(i interface{}) *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	s, ok := i.(string)
	if !ok {
		return AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumRef("")
	}

	return AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumRef(s)
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

type attestorDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         attestorApiOperation
}

func convertFieldDiffToAttestorOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]attestorDiff, error) {
	var diffs []attestorDiff
	for _, op := range ops {
		diff := attestorDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToattestorApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToattestorApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (attestorApiOperation, error) {
	switch op {

	case "updateAttestorUpdateAttestorOperation":
		return &updateAttestorUpdateAttestorOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
