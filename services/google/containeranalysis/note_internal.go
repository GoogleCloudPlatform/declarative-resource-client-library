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
package containeranalysis

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

func (r *Note) validate() error {

	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"Vulnerability", "Build", "Image", "Package", "Deployment", "Discovery", "Attestation"}, r.Vulnerability, r.Build, r.Image, r.Package, r.Deployment, r.Discovery, r.Attestation); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Vulnerability) {
		if err := r.Vulnerability.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Build) {
		if err := r.Build.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Image) {
		if err := r.Image.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Package) {
		if err := r.Package.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Discovery) {
		if err := r.Discovery.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Deployment) {
		if err := r.Deployment.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Attestation) {
		if err := r.Attestation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteRelatedUrl) validate() error {
	return nil
}
func (r *NoteVulnerability) validate() error {
	if !dcl.IsEmptyValueIndirect(r.CvssV3) {
		if err := r.CvssV3.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteVulnerabilityDetails) validate() error {
	if err := dcl.Required(r, "affectedCpeUri"); err != nil {
		return err
	}
	if err := dcl.Required(r, "affectedPackage"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.AffectedVersionStart) {
		if err := r.AffectedVersionStart.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AffectedVersionEnd) {
		if err := r.AffectedVersionEnd.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FixedVersion) {
		if err := r.FixedVersion.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteVulnerabilityDetailsAffectedVersionStart) validate() error {
	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteVulnerabilityDetailsAffectedVersionEnd) validate() error {
	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteVulnerabilityDetailsFixedVersion) validate() error {
	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteVulnerabilityCvssV3) validate() error {
	return nil
}
func (r *NoteVulnerabilityWindowsDetails) validate() error {
	if err := dcl.Required(r, "cpeUri"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "fixingKbs"); err != nil {
		return err
	}
	return nil
}
func (r *NoteVulnerabilityWindowsDetailsFixingKbs) validate() error {
	return nil
}
func (r *NoteBuild) validate() error {
	if err := dcl.Required(r, "builderVersion"); err != nil {
		return err
	}
	return nil
}
func (r *NoteImage) validate() error {
	if err := dcl.Required(r, "resourceUrl"); err != nil {
		return err
	}
	if err := dcl.Required(r, "fingerprint"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Fingerprint) {
		if err := r.Fingerprint.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteImageFingerprint) validate() error {
	if err := dcl.Required(r, "v1Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "v2Blob"); err != nil {
		return err
	}
	return nil
}
func (r *NotePackage) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	return nil
}
func (r *NotePackageDistribution) validate() error {
	if err := dcl.Required(r, "cpeUri"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.LatestVersion) {
		if err := r.LatestVersion.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NotePackageDistributionLatestVersion) validate() error {
	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteDiscovery) validate() error {
	if err := dcl.Required(r, "analysisKind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteDeployment) validate() error {
	if err := dcl.Required(r, "resourceUri"); err != nil {
		return err
	}
	return nil
}
func (r *NoteAttestation) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Hint) {
		if err := r.Hint.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteAttestationHint) validate() error {
	if err := dcl.Required(r, "humanReadableName"); err != nil {
		return err
	}
	return nil
}

func noteGetURL(userBasePath string, r *Note) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/notes/{{name}}", "https://containeranalysis.googleapis.com/v1/", userBasePath, params), nil
}

func noteListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/notes", "https://containeranalysis.googleapis.com/v1/", userBasePath, params), nil

}

func noteCreateURL(userBasePath, project, name string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"name":    name,
	}
	return dcl.URL("projects/{{project}}/notes?noteId={{name}}", "https://containeranalysis.googleapis.com/v1/", userBasePath, params), nil

}

func noteDeleteURL(userBasePath string, r *Note) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/notes/{{name}}", "https://containeranalysis.googleapis.com/v1/", userBasePath, params), nil
}

// noteApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type noteApiOperation interface {
	do(context.Context, *Note, *Client) error
}

// newUpdateNoteUpdateNoteRequest creates a request for an
// Note resource's UpdateNote update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNoteUpdateNoteRequest(ctx context.Context, f *Note, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.ShortDescription; !dcl.IsEmptyValueIndirect(v) {
		req["shortDescription"] = v
	}
	if v := f.LongDescription; !dcl.IsEmptyValueIndirect(v) {
		req["longDescription"] = v
	}
	if v, err := expandNoteRelatedUrlSlice(c, f.RelatedUrl); err != nil {
		return nil, fmt.Errorf("error expanding RelatedUrl into relatedUrl: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["relatedUrl"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		req["expirationTime"] = v
	}
	if v := f.RelatedNoteNames; !dcl.IsEmptyValueIndirect(v) {
		req["relatedNoteNames"] = v
	}
	if v, err := expandNoteVulnerability(c, f.Vulnerability); err != nil {
		return nil, fmt.Errorf("error expanding Vulnerability into vulnerability: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["vulnerability"] = v
	}
	if v, err := expandNoteBuild(c, f.Build); err != nil {
		return nil, fmt.Errorf("error expanding Build into build: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["build"] = v
	}
	if v, err := expandNoteImage(c, f.Image); err != nil {
		return nil, fmt.Errorf("error expanding Image into image: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["image"] = v
	}
	if v, err := expandNotePackage(c, f.Package); err != nil {
		return nil, fmt.Errorf("error expanding Package into package: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["package"] = v
	}
	if v, err := expandNoteDiscovery(c, f.Discovery); err != nil {
		return nil, fmt.Errorf("error expanding Discovery into discovery: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["discovery"] = v
	}
	if v, err := expandNoteDeployment(c, f.Deployment); err != nil {
		return nil, fmt.Errorf("error expanding Deployment into deployment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["deployment"] = v
	}
	if v, err := expandNoteAttestation(c, f.Attestation); err != nil {
		return nil, fmt.Errorf("error expanding Attestation into attestation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["attestation"] = v
	}
	return req, nil
}

// marshalUpdateNoteUpdateNoteRequest converts the update into
// the final JSON request body.
func marshalUpdateNoteUpdateNoteRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNoteUpdateNoteOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNoteUpdateNoteOperation) do(ctx context.Context, r *Note, c *Client) error {
	_, err := c.GetNote(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateNote")
	if err != nil {
		return err
	}

	req, err := newUpdateNoteUpdateNoteRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateNoteUpdateNoteRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listNoteRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := noteListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != NoteMaxPage {
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

type listNoteOperation struct {
	Notes []map[string]interface{} `json:"notes"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listNote(ctx context.Context, project, pageToken string, pageSize int32) ([]*Note, string, error) {
	b, err := c.listNoteRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listNoteOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Note
	for _, v := range m.Notes {
		res := flattenNote(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllNote(ctx context.Context, f func(*Note) bool, resources []*Note) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteNote(ctx, res)
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

type deleteNoteOperation struct{}

func (op *deleteNoteOperation) do(ctx context.Context, r *Note, c *Client) error {

	_, err := c.GetNote(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Note not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetNote checking for existence. error: %v", err)
		return err
	}

	u, err := noteDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Note: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetNote(ctx, r.urlNormalized())
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
type createNoteOperation struct {
	response map[string]interface{}
}

func (op *createNoteOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createNoteOperation) do(ctx context.Context, r *Note, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, name := r.createFields()
	u, err := noteCreateURL(c.Config.BasePath, project, name)

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

	if _, err := c.GetNote(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getNoteRaw(ctx context.Context, r *Note) ([]byte, error) {

	u, err := noteGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) noteDiffsForRawDesired(ctx context.Context, rawDesired *Note, opts ...dcl.ApplyOption) (initial, desired *Note, diffs []noteDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Note
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Note); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Note, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetNote(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Note resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Note resource: %v", err)
		}
		c.Config.Logger.Info("Found that Note resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeNoteDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Note: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Note: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeNoteInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Note: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeNoteDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Note: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffNote(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeNoteInitialState(rawInitial, rawDesired *Note) (*Note, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if dcl.IsZeroValue(rawInitial.Vulnerability) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Build, rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Vulnerability = EmptyNoteVulnerability
		}
	}

	if dcl.IsZeroValue(rawInitial.Build) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Build = EmptyNoteBuild
		}
	}

	if dcl.IsZeroValue(rawInitial.Image) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Image = EmptyNoteImage
		}
	}

	if dcl.IsZeroValue(rawInitial.Package) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Image, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Package = EmptyNotePackage
		}
	}

	if dcl.IsZeroValue(rawInitial.Deployment) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Image, rawInitial.Package, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Deployment = EmptyNoteDeployment
		}
	}

	if dcl.IsZeroValue(rawInitial.Discovery) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Attestation) {
			rawInitial.Discovery = EmptyNoteDiscovery
		}
	}

	if dcl.IsZeroValue(rawInitial.Attestation) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery) {
			rawInitial.Attestation = EmptyNoteAttestation
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeNoteDesiredState(rawDesired, rawInitial *Note, opts ...dcl.ApplyOption) (*Note, error) {

	if dcl.IsZeroValue(rawDesired.Vulnerability) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Build, rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			rawDesired.Vulnerability = EmptyNoteVulnerability
		}
	}

	if dcl.IsZeroValue(rawDesired.Build) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			rawDesired.Build = EmptyNoteBuild
		}
	}

	if dcl.IsZeroValue(rawDesired.Image) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			rawDesired.Image = EmptyNoteImage
		}
	}

	if dcl.IsZeroValue(rawDesired.Package) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Image, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			rawDesired.Package = EmptyNotePackage
		}
	}

	if dcl.IsZeroValue(rawDesired.Deployment) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Image, rawDesired.Package, rawDesired.Discovery, rawDesired.Attestation) {
			rawDesired.Deployment = EmptyNoteDeployment
		}
	}

	if dcl.IsZeroValue(rawDesired.Discovery) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Attestation) {
			rawDesired.Discovery = EmptyNoteDiscovery
		}
	}

	if dcl.IsZeroValue(rawDesired.Attestation) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery) {
			rawDesired.Attestation = EmptyNoteAttestation
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Vulnerability = canonicalizeNoteVulnerability(rawDesired.Vulnerability, nil, opts...)
		rawDesired.Build = canonicalizeNoteBuild(rawDesired.Build, nil, opts...)
		rawDesired.Image = canonicalizeNoteImage(rawDesired.Image, nil, opts...)
		rawDesired.Package = canonicalizeNotePackage(rawDesired.Package, nil, opts...)
		rawDesired.Discovery = canonicalizeNoteDiscovery(rawDesired.Discovery, nil, opts...)
		rawDesired.Deployment = canonicalizeNoteDeployment(rawDesired.Deployment, nil, opts...)
		rawDesired.Attestation = canonicalizeNoteAttestation(rawDesired.Attestation, nil, opts...)

		return rawDesired, nil
	}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.ShortDescription, rawInitial.ShortDescription) {
		rawDesired.ShortDescription = rawInitial.ShortDescription
	}
	if dcl.StringCanonicalize(rawDesired.LongDescription, rawInitial.LongDescription) {
		rawDesired.LongDescription = rawInitial.LongDescription
	}
	if dcl.IsZeroValue(rawDesired.RelatedUrl) {
		rawDesired.RelatedUrl = rawInitial.RelatedUrl
	}
	if dcl.IsZeroValue(rawDesired.ExpirationTime) {
		rawDesired.ExpirationTime = rawInitial.ExpirationTime
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.IsZeroValue(rawDesired.RelatedNoteNames) {
		rawDesired.RelatedNoteNames = rawInitial.RelatedNoteNames
	}
	rawDesired.Vulnerability = canonicalizeNoteVulnerability(rawDesired.Vulnerability, rawInitial.Vulnerability, opts...)
	rawDesired.Build = canonicalizeNoteBuild(rawDesired.Build, rawInitial.Build, opts...)
	rawDesired.Image = canonicalizeNoteImage(rawDesired.Image, rawInitial.Image, opts...)
	rawDesired.Package = canonicalizeNotePackage(rawDesired.Package, rawInitial.Package, opts...)
	rawDesired.Discovery = canonicalizeNoteDiscovery(rawDesired.Discovery, rawInitial.Discovery, opts...)
	rawDesired.Deployment = canonicalizeNoteDeployment(rawDesired.Deployment, rawInitial.Deployment, opts...)
	rawDesired.Attestation = canonicalizeNoteAttestation(rawDesired.Attestation, rawInitial.Attestation, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeNoteNewState(c *Client, rawNew, rawDesired *Note) (*Note, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.ShortDescription) && dcl.IsEmptyValueIndirect(rawDesired.ShortDescription) {
		rawNew.ShortDescription = rawDesired.ShortDescription
	} else {
		if dcl.StringCanonicalize(rawDesired.ShortDescription, rawNew.ShortDescription) {
			rawNew.ShortDescription = rawDesired.ShortDescription
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LongDescription) && dcl.IsEmptyValueIndirect(rawDesired.LongDescription) {
		rawNew.LongDescription = rawDesired.LongDescription
	} else {
		if dcl.StringCanonicalize(rawDesired.LongDescription, rawNew.LongDescription) {
			rawNew.LongDescription = rawDesired.LongDescription
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RelatedUrl) && dcl.IsEmptyValueIndirect(rawDesired.RelatedUrl) {
		rawNew.RelatedUrl = rawDesired.RelatedUrl
	} else {
		rawNew.RelatedUrl = canonicalizeNewNoteRelatedUrlSlice(c, rawDesired.RelatedUrl, rawNew.RelatedUrl)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpirationTime) && dcl.IsEmptyValueIndirect(rawDesired.ExpirationTime) {
		rawNew.ExpirationTime = rawDesired.ExpirationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RelatedNoteNames) && dcl.IsEmptyValueIndirect(rawDesired.RelatedNoteNames) {
		rawNew.RelatedNoteNames = rawDesired.RelatedNoteNames
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Vulnerability) && dcl.IsEmptyValueIndirect(rawDesired.Vulnerability) {
		rawNew.Vulnerability = rawDesired.Vulnerability
	} else {
		rawNew.Vulnerability = canonicalizeNewNoteVulnerability(c, rawDesired.Vulnerability, rawNew.Vulnerability)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Build) && dcl.IsEmptyValueIndirect(rawDesired.Build) {
		rawNew.Build = rawDesired.Build
	} else {
		rawNew.Build = canonicalizeNewNoteBuild(c, rawDesired.Build, rawNew.Build)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Image) && dcl.IsEmptyValueIndirect(rawDesired.Image) {
		rawNew.Image = rawDesired.Image
	} else {
		rawNew.Image = canonicalizeNewNoteImage(c, rawDesired.Image, rawNew.Image)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Package) && dcl.IsEmptyValueIndirect(rawDesired.Package) {
		rawNew.Package = rawDesired.Package
	} else {
		rawNew.Package = canonicalizeNewNotePackage(c, rawDesired.Package, rawNew.Package)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Discovery) && dcl.IsEmptyValueIndirect(rawDesired.Discovery) {
		rawNew.Discovery = rawDesired.Discovery
	} else {
		rawNew.Discovery = canonicalizeNewNoteDiscovery(c, rawDesired.Discovery, rawNew.Discovery)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Deployment) && dcl.IsEmptyValueIndirect(rawDesired.Deployment) {
		rawNew.Deployment = rawDesired.Deployment
	} else {
		rawNew.Deployment = canonicalizeNewNoteDeployment(c, rawDesired.Deployment, rawNew.Deployment)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Attestation) && dcl.IsEmptyValueIndirect(rawDesired.Attestation) {
		rawNew.Attestation = rawDesired.Attestation
	} else {
		rawNew.Attestation = canonicalizeNewNoteAttestation(c, rawDesired.Attestation, rawNew.Attestation)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeNoteRelatedUrl(des, initial *NoteRelatedUrl, opts ...dcl.ApplyOption) *NoteRelatedUrl {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}
	if dcl.StringCanonicalize(des.Label, initial.Label) || dcl.IsZeroValue(des.Label) {
		des.Label = initial.Label
	}

	return des
}

func canonicalizeNewNoteRelatedUrl(c *Client, des, nw *NoteRelatedUrl) *NoteRelatedUrl {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Url, nw.Url) || dcl.IsZeroValue(des.Url) {
		nw.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Label, nw.Label) || dcl.IsZeroValue(des.Label) {
		nw.Label = des.Label
	}

	return nw
}

func canonicalizeNewNoteRelatedUrlSet(c *Client, des, nw []NoteRelatedUrl) []NoteRelatedUrl {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteRelatedUrl
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteRelatedUrl(c, &d, &n) {
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

func canonicalizeNewNoteRelatedUrlSlice(c *Client, des, nw []NoteRelatedUrl) []NoteRelatedUrl {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteRelatedUrl
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteRelatedUrl(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerability(des, initial *NoteVulnerability, opts ...dcl.ApplyOption) *NoteVulnerability {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.CvssScore) {
		des.CvssScore = initial.CvssScore
	}
	if dcl.IsZeroValue(des.Severity) {
		des.Severity = initial.Severity
	}
	if dcl.IsZeroValue(des.Details) {
		des.Details = initial.Details
	}
	des.CvssV3 = canonicalizeNoteVulnerabilityCvssV3(des.CvssV3, initial.CvssV3, opts...)
	if dcl.IsZeroValue(des.WindowsDetails) {
		des.WindowsDetails = initial.WindowsDetails
	}
	if dcl.IsZeroValue(des.SourceUpdateTime) {
		des.SourceUpdateTime = initial.SourceUpdateTime
	}

	return des
}

func canonicalizeNewNoteVulnerability(c *Client, des, nw *NoteVulnerability) *NoteVulnerability {
	if des == nil || nw == nil {
		return nw
	}

	nw.Details = canonicalizeNewNoteVulnerabilityDetailsSlice(c, des.Details, nw.Details)
	nw.CvssV3 = canonicalizeNewNoteVulnerabilityCvssV3(c, des.CvssV3, nw.CvssV3)
	nw.WindowsDetails = canonicalizeNewNoteVulnerabilityWindowsDetailsSlice(c, des.WindowsDetails, nw.WindowsDetails)

	return nw
}

func canonicalizeNewNoteVulnerabilitySet(c *Client, des, nw []NoteVulnerability) []NoteVulnerability {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteVulnerability
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteVulnerability(c, &d, &n) {
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

func canonicalizeNewNoteVulnerabilitySlice(c *Client, des, nw []NoteVulnerability) []NoteVulnerability {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteVulnerability
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerability(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityDetails(des, initial *NoteVulnerabilityDetails, opts ...dcl.ApplyOption) *NoteVulnerabilityDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.SeverityName, initial.SeverityName) || dcl.IsZeroValue(des.SeverityName) {
		des.SeverityName = initial.SeverityName
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.StringCanonicalize(des.PackageType, initial.PackageType) || dcl.IsZeroValue(des.PackageType) {
		des.PackageType = initial.PackageType
	}
	if dcl.StringCanonicalize(des.AffectedCpeUri, initial.AffectedCpeUri) || dcl.IsZeroValue(des.AffectedCpeUri) {
		des.AffectedCpeUri = initial.AffectedCpeUri
	}
	if dcl.StringCanonicalize(des.AffectedPackage, initial.AffectedPackage) || dcl.IsZeroValue(des.AffectedPackage) {
		des.AffectedPackage = initial.AffectedPackage
	}
	des.AffectedVersionStart = canonicalizeNoteVulnerabilityDetailsAffectedVersionStart(des.AffectedVersionStart, initial.AffectedVersionStart, opts...)
	des.AffectedVersionEnd = canonicalizeNoteVulnerabilityDetailsAffectedVersionEnd(des.AffectedVersionEnd, initial.AffectedVersionEnd, opts...)
	if dcl.StringCanonicalize(des.FixedCpeUri, initial.FixedCpeUri) || dcl.IsZeroValue(des.FixedCpeUri) {
		des.FixedCpeUri = initial.FixedCpeUri
	}
	if dcl.StringCanonicalize(des.FixedPackage, initial.FixedPackage) || dcl.IsZeroValue(des.FixedPackage) {
		des.FixedPackage = initial.FixedPackage
	}
	des.FixedVersion = canonicalizeNoteVulnerabilityDetailsFixedVersion(des.FixedVersion, initial.FixedVersion, opts...)
	if dcl.BoolCanonicalize(des.IsObsolete, initial.IsObsolete) || dcl.IsZeroValue(des.IsObsolete) {
		des.IsObsolete = initial.IsObsolete
	}
	if dcl.IsZeroValue(des.SourceUpdateTime) {
		des.SourceUpdateTime = initial.SourceUpdateTime
	}

	return des
}

func canonicalizeNewNoteVulnerabilityDetails(c *Client, des, nw *NoteVulnerabilityDetails) *NoteVulnerabilityDetails {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.SeverityName, nw.SeverityName) || dcl.IsZeroValue(des.SeverityName) {
		nw.SeverityName = des.SeverityName
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) || dcl.IsZeroValue(des.Description) {
		nw.Description = des.Description
	}
	if dcl.StringCanonicalize(des.PackageType, nw.PackageType) || dcl.IsZeroValue(des.PackageType) {
		nw.PackageType = des.PackageType
	}
	if dcl.StringCanonicalize(des.AffectedCpeUri, nw.AffectedCpeUri) || dcl.IsZeroValue(des.AffectedCpeUri) {
		nw.AffectedCpeUri = des.AffectedCpeUri
	}
	if dcl.StringCanonicalize(des.AffectedPackage, nw.AffectedPackage) || dcl.IsZeroValue(des.AffectedPackage) {
		nw.AffectedPackage = des.AffectedPackage
	}
	nw.AffectedVersionStart = canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStart(c, des.AffectedVersionStart, nw.AffectedVersionStart)
	nw.AffectedVersionEnd = canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEnd(c, des.AffectedVersionEnd, nw.AffectedVersionEnd)
	if dcl.StringCanonicalize(des.FixedCpeUri, nw.FixedCpeUri) || dcl.IsZeroValue(des.FixedCpeUri) {
		nw.FixedCpeUri = des.FixedCpeUri
	}
	if dcl.StringCanonicalize(des.FixedPackage, nw.FixedPackage) || dcl.IsZeroValue(des.FixedPackage) {
		nw.FixedPackage = des.FixedPackage
	}
	nw.FixedVersion = canonicalizeNewNoteVulnerabilityDetailsFixedVersion(c, des.FixedVersion, nw.FixedVersion)
	if dcl.BoolCanonicalize(des.IsObsolete, nw.IsObsolete) || dcl.IsZeroValue(des.IsObsolete) {
		nw.IsObsolete = des.IsObsolete
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityDetailsSet(c *Client, des, nw []NoteVulnerabilityDetails) []NoteVulnerabilityDetails {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteVulnerabilityDetails
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteVulnerabilityDetails(c, &d, &n) {
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

func canonicalizeNewNoteVulnerabilityDetailsSlice(c *Client, des, nw []NoteVulnerabilityDetails) []NoteVulnerabilityDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteVulnerabilityDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityDetails(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityDetailsAffectedVersionStart(des, initial *NoteVulnerabilityDetailsAffectedVersionStart, opts ...dcl.ApplyOption) *NoteVulnerabilityDetailsAffectedVersionStart {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Epoch) {
		des.Epoch = initial.Epoch
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		des.Revision = initial.Revision
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.StringCanonicalize(des.FullName, initial.FullName) || dcl.IsZeroValue(des.FullName) {
		des.FullName = initial.FullName
	}

	return des
}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStart(c *Client, des, nw *NoteVulnerabilityDetailsAffectedVersionStart) *NoteVulnerabilityDetailsAffectedVersionStart {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, nw.Revision) || dcl.IsZeroValue(des.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) || dcl.IsZeroValue(des.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStartSet(c *Client, des, nw []NoteVulnerabilityDetailsAffectedVersionStart) []NoteVulnerabilityDetailsAffectedVersionStart {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteVulnerabilityDetailsAffectedVersionStart
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteVulnerabilityDetailsAffectedVersionStart(c, &d, &n) {
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

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStartSlice(c *Client, des, nw []NoteVulnerabilityDetailsAffectedVersionStart) []NoteVulnerabilityDetailsAffectedVersionStart {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteVulnerabilityDetailsAffectedVersionStart
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStart(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityDetailsAffectedVersionEnd(des, initial *NoteVulnerabilityDetailsAffectedVersionEnd, opts ...dcl.ApplyOption) *NoteVulnerabilityDetailsAffectedVersionEnd {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Epoch) {
		des.Epoch = initial.Epoch
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		des.Revision = initial.Revision
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.StringCanonicalize(des.FullName, initial.FullName) || dcl.IsZeroValue(des.FullName) {
		des.FullName = initial.FullName
	}

	return des
}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEnd(c *Client, des, nw *NoteVulnerabilityDetailsAffectedVersionEnd) *NoteVulnerabilityDetailsAffectedVersionEnd {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, nw.Revision) || dcl.IsZeroValue(des.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) || dcl.IsZeroValue(des.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEndSet(c *Client, des, nw []NoteVulnerabilityDetailsAffectedVersionEnd) []NoteVulnerabilityDetailsAffectedVersionEnd {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteVulnerabilityDetailsAffectedVersionEnd
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteVulnerabilityDetailsAffectedVersionEnd(c, &d, &n) {
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

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEndSlice(c *Client, des, nw []NoteVulnerabilityDetailsAffectedVersionEnd) []NoteVulnerabilityDetailsAffectedVersionEnd {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteVulnerabilityDetailsAffectedVersionEnd
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEnd(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityDetailsFixedVersion(des, initial *NoteVulnerabilityDetailsFixedVersion, opts ...dcl.ApplyOption) *NoteVulnerabilityDetailsFixedVersion {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Epoch) {
		des.Epoch = initial.Epoch
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		des.Revision = initial.Revision
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.StringCanonicalize(des.FullName, initial.FullName) || dcl.IsZeroValue(des.FullName) {
		des.FullName = initial.FullName
	}

	return des
}

func canonicalizeNewNoteVulnerabilityDetailsFixedVersion(c *Client, des, nw *NoteVulnerabilityDetailsFixedVersion) *NoteVulnerabilityDetailsFixedVersion {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, nw.Revision) || dcl.IsZeroValue(des.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) || dcl.IsZeroValue(des.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityDetailsFixedVersionSet(c *Client, des, nw []NoteVulnerabilityDetailsFixedVersion) []NoteVulnerabilityDetailsFixedVersion {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteVulnerabilityDetailsFixedVersion
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteVulnerabilityDetailsFixedVersion(c, &d, &n) {
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

func canonicalizeNewNoteVulnerabilityDetailsFixedVersionSlice(c *Client, des, nw []NoteVulnerabilityDetailsFixedVersion) []NoteVulnerabilityDetailsFixedVersion {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteVulnerabilityDetailsFixedVersion
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityDetailsFixedVersion(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityCvssV3(des, initial *NoteVulnerabilityCvssV3, opts ...dcl.ApplyOption) *NoteVulnerabilityCvssV3 {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.BaseScore) {
		des.BaseScore = initial.BaseScore
	}
	if dcl.IsZeroValue(des.ExploitabilityScore) {
		des.ExploitabilityScore = initial.ExploitabilityScore
	}
	if dcl.IsZeroValue(des.ImpactScore) {
		des.ImpactScore = initial.ImpactScore
	}
	if dcl.IsZeroValue(des.AttackVector) {
		des.AttackVector = initial.AttackVector
	}
	if dcl.IsZeroValue(des.AttackComplexity) {
		des.AttackComplexity = initial.AttackComplexity
	}
	if dcl.IsZeroValue(des.PrivilegesRequired) {
		des.PrivilegesRequired = initial.PrivilegesRequired
	}
	if dcl.IsZeroValue(des.UserInteraction) {
		des.UserInteraction = initial.UserInteraction
	}
	if dcl.IsZeroValue(des.Scope) {
		des.Scope = initial.Scope
	}
	if dcl.IsZeroValue(des.ConfidentialityImpact) {
		des.ConfidentialityImpact = initial.ConfidentialityImpact
	}
	if dcl.IsZeroValue(des.IntegrityImpact) {
		des.IntegrityImpact = initial.IntegrityImpact
	}
	if dcl.IsZeroValue(des.AvailabilityImpact) {
		des.AvailabilityImpact = initial.AvailabilityImpact
	}

	return des
}

func canonicalizeNewNoteVulnerabilityCvssV3(c *Client, des, nw *NoteVulnerabilityCvssV3) *NoteVulnerabilityCvssV3 {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityCvssV3Set(c *Client, des, nw []NoteVulnerabilityCvssV3) []NoteVulnerabilityCvssV3 {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteVulnerabilityCvssV3
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteVulnerabilityCvssV3(c, &d, &n) {
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

func canonicalizeNewNoteVulnerabilityCvssV3Slice(c *Client, des, nw []NoteVulnerabilityCvssV3) []NoteVulnerabilityCvssV3 {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteVulnerabilityCvssV3
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityCvssV3(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityWindowsDetails(des, initial *NoteVulnerabilityWindowsDetails, opts ...dcl.ApplyOption) *NoteVulnerabilityWindowsDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CpeUri, initial.CpeUri) || dcl.IsZeroValue(des.CpeUri) {
		des.CpeUri = initial.CpeUri
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.FixingKbs) {
		des.FixingKbs = initial.FixingKbs
	}

	return des
}

func canonicalizeNewNoteVulnerabilityWindowsDetails(c *Client, des, nw *NoteVulnerabilityWindowsDetails) *NoteVulnerabilityWindowsDetails {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CpeUri, nw.CpeUri) || dcl.IsZeroValue(des.CpeUri) {
		nw.CpeUri = des.CpeUri
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) || dcl.IsZeroValue(des.Description) {
		nw.Description = des.Description
	}
	nw.FixingKbs = canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbsSlice(c, des.FixingKbs, nw.FixingKbs)

	return nw
}

func canonicalizeNewNoteVulnerabilityWindowsDetailsSet(c *Client, des, nw []NoteVulnerabilityWindowsDetails) []NoteVulnerabilityWindowsDetails {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteVulnerabilityWindowsDetails
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteVulnerabilityWindowsDetails(c, &d, &n) {
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

func canonicalizeNewNoteVulnerabilityWindowsDetailsSlice(c *Client, des, nw []NoteVulnerabilityWindowsDetails) []NoteVulnerabilityWindowsDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteVulnerabilityWindowsDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityWindowsDetails(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityWindowsDetailsFixingKbs(des, initial *NoteVulnerabilityWindowsDetailsFixingKbs, opts ...dcl.ApplyOption) *NoteVulnerabilityWindowsDetailsFixingKbs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}

	return des
}

func canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbs(c *Client, des, nw *NoteVulnerabilityWindowsDetailsFixingKbs) *NoteVulnerabilityWindowsDetailsFixingKbs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Url, nw.Url) || dcl.IsZeroValue(des.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbsSet(c *Client, des, nw []NoteVulnerabilityWindowsDetailsFixingKbs) []NoteVulnerabilityWindowsDetailsFixingKbs {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteVulnerabilityWindowsDetailsFixingKbs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteVulnerabilityWindowsDetailsFixingKbs(c, &d, &n) {
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

func canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbsSlice(c *Client, des, nw []NoteVulnerabilityWindowsDetailsFixingKbs) []NoteVulnerabilityWindowsDetailsFixingKbs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteVulnerabilityWindowsDetailsFixingKbs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbs(c, &d, &n))
	}

	return items
}

func canonicalizeNoteBuild(des, initial *NoteBuild, opts ...dcl.ApplyOption) *NoteBuild {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.BuilderVersion, initial.BuilderVersion) || dcl.IsZeroValue(des.BuilderVersion) {
		des.BuilderVersion = initial.BuilderVersion
	}

	return des
}

func canonicalizeNewNoteBuild(c *Client, des, nw *NoteBuild) *NoteBuild {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.BuilderVersion, nw.BuilderVersion) || dcl.IsZeroValue(des.BuilderVersion) {
		nw.BuilderVersion = des.BuilderVersion
	}

	return nw
}

func canonicalizeNewNoteBuildSet(c *Client, des, nw []NoteBuild) []NoteBuild {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteBuild
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteBuild(c, &d, &n) {
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

func canonicalizeNewNoteBuildSlice(c *Client, des, nw []NoteBuild) []NoteBuild {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteBuild
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteBuild(c, &d, &n))
	}

	return items
}

func canonicalizeNoteImage(des, initial *NoteImage, opts ...dcl.ApplyOption) *NoteImage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ResourceUrl, initial.ResourceUrl) || dcl.IsZeroValue(des.ResourceUrl) {
		des.ResourceUrl = initial.ResourceUrl
	}
	des.Fingerprint = canonicalizeNoteImageFingerprint(des.Fingerprint, initial.Fingerprint, opts...)

	return des
}

func canonicalizeNewNoteImage(c *Client, des, nw *NoteImage) *NoteImage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ResourceUrl, nw.ResourceUrl) || dcl.IsZeroValue(des.ResourceUrl) {
		nw.ResourceUrl = des.ResourceUrl
	}
	nw.Fingerprint = canonicalizeNewNoteImageFingerprint(c, des.Fingerprint, nw.Fingerprint)

	return nw
}

func canonicalizeNewNoteImageSet(c *Client, des, nw []NoteImage) []NoteImage {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteImage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteImage(c, &d, &n) {
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

func canonicalizeNewNoteImageSlice(c *Client, des, nw []NoteImage) []NoteImage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteImage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteImage(c, &d, &n))
	}

	return items
}

func canonicalizeNoteImageFingerprint(des, initial *NoteImageFingerprint, opts ...dcl.ApplyOption) *NoteImageFingerprint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.V1Name, initial.V1Name) || dcl.IsZeroValue(des.V1Name) {
		des.V1Name = initial.V1Name
	}
	if dcl.IsZeroValue(des.V2Blob) {
		des.V2Blob = initial.V2Blob
	}
	if dcl.StringCanonicalize(des.V2Name, initial.V2Name) || dcl.IsZeroValue(des.V2Name) {
		des.V2Name = initial.V2Name
	}

	return des
}

func canonicalizeNewNoteImageFingerprint(c *Client, des, nw *NoteImageFingerprint) *NoteImageFingerprint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.V1Name, nw.V1Name) || dcl.IsZeroValue(des.V1Name) {
		nw.V1Name = des.V1Name
	}
	if dcl.StringCanonicalize(des.V2Name, nw.V2Name) || dcl.IsZeroValue(des.V2Name) {
		nw.V2Name = des.V2Name
	}

	return nw
}

func canonicalizeNewNoteImageFingerprintSet(c *Client, des, nw []NoteImageFingerprint) []NoteImageFingerprint {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteImageFingerprint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteImageFingerprint(c, &d, &n) {
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

func canonicalizeNewNoteImageFingerprintSlice(c *Client, des, nw []NoteImageFingerprint) []NoteImageFingerprint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteImageFingerprint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteImageFingerprint(c, &d, &n))
	}

	return items
}

func canonicalizeNotePackage(des, initial *NotePackage, opts ...dcl.ApplyOption) *NotePackage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Distribution) {
		des.Distribution = initial.Distribution
	}

	return des
}

func canonicalizeNewNotePackage(c *Client, des, nw *NotePackage) *NotePackage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	nw.Distribution = canonicalizeNewNotePackageDistributionSlice(c, des.Distribution, nw.Distribution)

	return nw
}

func canonicalizeNewNotePackageSet(c *Client, des, nw []NotePackage) []NotePackage {
	if des == nil {
		return nw
	}
	var reorderedNew []NotePackage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNotePackage(c, &d, &n) {
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

func canonicalizeNewNotePackageSlice(c *Client, des, nw []NotePackage) []NotePackage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NotePackage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNotePackage(c, &d, &n))
	}

	return items
}

func canonicalizeNotePackageDistribution(des, initial *NotePackageDistribution, opts ...dcl.ApplyOption) *NotePackageDistribution {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CpeUri, initial.CpeUri) || dcl.IsZeroValue(des.CpeUri) {
		des.CpeUri = initial.CpeUri
	}
	if dcl.IsZeroValue(des.Architecture) {
		des.Architecture = initial.Architecture
	}
	des.LatestVersion = canonicalizeNotePackageDistributionLatestVersion(des.LatestVersion, initial.LatestVersion, opts...)
	if dcl.StringCanonicalize(des.Maintainer, initial.Maintainer) || dcl.IsZeroValue(des.Maintainer) {
		des.Maintainer = initial.Maintainer
	}
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}

	return des
}

func canonicalizeNewNotePackageDistribution(c *Client, des, nw *NotePackageDistribution) *NotePackageDistribution {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CpeUri, nw.CpeUri) || dcl.IsZeroValue(des.CpeUri) {
		nw.CpeUri = des.CpeUri
	}
	nw.LatestVersion = canonicalizeNewNotePackageDistributionLatestVersion(c, des.LatestVersion, nw.LatestVersion)
	if dcl.StringCanonicalize(des.Maintainer, nw.Maintainer) || dcl.IsZeroValue(des.Maintainer) {
		nw.Maintainer = des.Maintainer
	}
	if dcl.StringCanonicalize(des.Url, nw.Url) || dcl.IsZeroValue(des.Url) {
		nw.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) || dcl.IsZeroValue(des.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewNotePackageDistributionSet(c *Client, des, nw []NotePackageDistribution) []NotePackageDistribution {
	if des == nil {
		return nw
	}
	var reorderedNew []NotePackageDistribution
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNotePackageDistribution(c, &d, &n) {
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

func canonicalizeNewNotePackageDistributionSlice(c *Client, des, nw []NotePackageDistribution) []NotePackageDistribution {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NotePackageDistribution
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNotePackageDistribution(c, &d, &n))
	}

	return items
}

func canonicalizeNotePackageDistributionLatestVersion(des, initial *NotePackageDistributionLatestVersion, opts ...dcl.ApplyOption) *NotePackageDistributionLatestVersion {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Epoch) {
		des.Epoch = initial.Epoch
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		des.Revision = initial.Revision
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.StringCanonicalize(des.FullName, initial.FullName) || dcl.IsZeroValue(des.FullName) {
		des.FullName = initial.FullName
	}

	return des
}

func canonicalizeNewNotePackageDistributionLatestVersion(c *Client, des, nw *NotePackageDistributionLatestVersion) *NotePackageDistributionLatestVersion {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, nw.Revision) || dcl.IsZeroValue(des.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) || dcl.IsZeroValue(des.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewNotePackageDistributionLatestVersionSet(c *Client, des, nw []NotePackageDistributionLatestVersion) []NotePackageDistributionLatestVersion {
	if des == nil {
		return nw
	}
	var reorderedNew []NotePackageDistributionLatestVersion
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNotePackageDistributionLatestVersion(c, &d, &n) {
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

func canonicalizeNewNotePackageDistributionLatestVersionSlice(c *Client, des, nw []NotePackageDistributionLatestVersion) []NotePackageDistributionLatestVersion {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NotePackageDistributionLatestVersion
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNotePackageDistributionLatestVersion(c, &d, &n))
	}

	return items
}

func canonicalizeNoteDiscovery(des, initial *NoteDiscovery, opts ...dcl.ApplyOption) *NoteDiscovery {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AnalysisKind) {
		des.AnalysisKind = initial.AnalysisKind
	}

	return des
}

func canonicalizeNewNoteDiscovery(c *Client, des, nw *NoteDiscovery) *NoteDiscovery {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewNoteDiscoverySet(c *Client, des, nw []NoteDiscovery) []NoteDiscovery {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteDiscovery
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteDiscovery(c, &d, &n) {
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

func canonicalizeNewNoteDiscoverySlice(c *Client, des, nw []NoteDiscovery) []NoteDiscovery {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteDiscovery
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteDiscovery(c, &d, &n))
	}

	return items
}

func canonicalizeNoteDeployment(des, initial *NoteDeployment, opts ...dcl.ApplyOption) *NoteDeployment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ResourceUri) {
		des.ResourceUri = initial.ResourceUri
	}

	return des
}

func canonicalizeNewNoteDeployment(c *Client, des, nw *NoteDeployment) *NoteDeployment {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewNoteDeploymentSet(c *Client, des, nw []NoteDeployment) []NoteDeployment {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteDeployment
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteDeployment(c, &d, &n) {
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

func canonicalizeNewNoteDeploymentSlice(c *Client, des, nw []NoteDeployment) []NoteDeployment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteDeployment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteDeployment(c, &d, &n))
	}

	return items
}

func canonicalizeNoteAttestation(des, initial *NoteAttestation, opts ...dcl.ApplyOption) *NoteAttestation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Hint = canonicalizeNoteAttestationHint(des.Hint, initial.Hint, opts...)

	return des
}

func canonicalizeNewNoteAttestation(c *Client, des, nw *NoteAttestation) *NoteAttestation {
	if des == nil || nw == nil {
		return nw
	}

	nw.Hint = canonicalizeNewNoteAttestationHint(c, des.Hint, nw.Hint)

	return nw
}

func canonicalizeNewNoteAttestationSet(c *Client, des, nw []NoteAttestation) []NoteAttestation {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteAttestation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteAttestation(c, &d, &n) {
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

func canonicalizeNewNoteAttestationSlice(c *Client, des, nw []NoteAttestation) []NoteAttestation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteAttestation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteAttestation(c, &d, &n))
	}

	return items
}

func canonicalizeNoteAttestationHint(des, initial *NoteAttestationHint, opts ...dcl.ApplyOption) *NoteAttestationHint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.HumanReadableName, initial.HumanReadableName) || dcl.IsZeroValue(des.HumanReadableName) {
		des.HumanReadableName = initial.HumanReadableName
	}

	return des
}

func canonicalizeNewNoteAttestationHint(c *Client, des, nw *NoteAttestationHint) *NoteAttestationHint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HumanReadableName, nw.HumanReadableName) || dcl.IsZeroValue(des.HumanReadableName) {
		nw.HumanReadableName = des.HumanReadableName
	}

	return nw
}

func canonicalizeNewNoteAttestationHintSet(c *Client, des, nw []NoteAttestationHint) []NoteAttestationHint {
	if des == nil {
		return nw
	}
	var reorderedNew []NoteAttestationHint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNoteAttestationHint(c, &d, &n) {
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

func canonicalizeNewNoteAttestationHintSlice(c *Client, des, nw []NoteAttestationHint) []NoteAttestationHint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NoteAttestationHint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteAttestationHint(c, &d, &n))
	}

	return items
}

type noteDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         noteApiOperation
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
func diffNote(c *Client, desired, actual *Note, opts ...dcl.ApplyOption) ([]noteDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []noteDiff
	if !dcl.IsZeroValue(desired.ShortDescription) && !dcl.StringCanonicalize(desired.ShortDescription, actual.ShortDescription) {
		c.Config.Logger.Infof("Detected diff in ShortDescription.\nDESIRED: %v\nACTUAL: %v", desired.ShortDescription, actual.ShortDescription)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "ShortDescription",
		})

	}
	if !dcl.IsZeroValue(desired.LongDescription) && !dcl.StringCanonicalize(desired.LongDescription, actual.LongDescription) {
		c.Config.Logger.Infof("Detected diff in LongDescription.\nDESIRED: %v\nACTUAL: %v", desired.LongDescription, actual.LongDescription)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "LongDescription",
		})

	}
	if compareNoteRelatedUrlSlice(c, desired.RelatedUrl, actual.RelatedUrl) {
		c.Config.Logger.Infof("Detected diff in RelatedUrl.\nDESIRED: %v\nACTUAL: %v", desired.RelatedUrl, actual.RelatedUrl)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "RelatedUrl",
		})

	}
	if !reflect.DeepEqual(desired.ExpirationTime, actual.ExpirationTime) {
		c.Config.Logger.Infof("Detected diff in ExpirationTime.\nDESIRED: %v\nACTUAL: %v", desired.ExpirationTime, actual.ExpirationTime)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "ExpirationTime",
		})

	}
	if !reflect.DeepEqual(desired.RelatedNoteNames, actual.RelatedNoteNames) {
		c.Config.Logger.Infof("Detected diff in RelatedNoteNames.\nDESIRED: %v\nACTUAL: %v", desired.RelatedNoteNames, actual.RelatedNoteNames)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "RelatedNoteNames",
		})

	}
	if compareNoteVulnerability(c, desired.Vulnerability, actual.Vulnerability) {
		c.Config.Logger.Infof("Detected diff in Vulnerability.\nDESIRED: %v\nACTUAL: %v", desired.Vulnerability, actual.Vulnerability)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "Vulnerability",
		})

	}
	if compareNoteBuild(c, desired.Build, actual.Build) {
		c.Config.Logger.Infof("Detected diff in Build.\nDESIRED: %v\nACTUAL: %v", desired.Build, actual.Build)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "Build",
		})

	}
	if compareNoteImage(c, desired.Image, actual.Image) {
		c.Config.Logger.Infof("Detected diff in Image.\nDESIRED: %v\nACTUAL: %v", desired.Image, actual.Image)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "Image",
		})

	}
	if compareNotePackage(c, desired.Package, actual.Package) {
		c.Config.Logger.Infof("Detected diff in Package.\nDESIRED: %v\nACTUAL: %v", desired.Package, actual.Package)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "Package",
		})

	}
	if compareNoteDiscovery(c, desired.Discovery, actual.Discovery) {
		c.Config.Logger.Infof("Detected diff in Discovery.\nDESIRED: %v\nACTUAL: %v", desired.Discovery, actual.Discovery)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "Discovery",
		})

	}
	if compareNoteDeployment(c, desired.Deployment, actual.Deployment) {
		c.Config.Logger.Infof("Detected diff in Deployment.\nDESIRED: %v\nACTUAL: %v", desired.Deployment, actual.Deployment)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "Deployment",
		})

	}
	if compareNoteAttestation(c, desired.Attestation, actual.Attestation) {
		c.Config.Logger.Infof("Detected diff in Attestation.\nDESIRED: %v\nACTUAL: %v", desired.Attestation, actual.Attestation)

		diffs = append(diffs, noteDiff{
			UpdateOp:  &updateNoteUpdateNoteOperation{},
			FieldName: "Attestation",
		})

	}
	diffs = analyzeNoteDiff(desired, actual, diffs)
	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []noteDiff
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
func compareNoteRelatedUrl(c *Client, desired, actual *NoteRelatedUrl) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Url == nil && desired.Url != nil && !dcl.IsEmptyValueIndirect(desired.Url) {
		c.Config.Logger.Infof("desired Url %s - but actually nil", dcl.SprintResource(desired.Url))
		return true
	}
	if !dcl.StringCanonicalize(desired.Url, actual.Url) && !dcl.IsZeroValue(desired.Url) {
		c.Config.Logger.Infof("Diff in Url. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Url), dcl.SprintResource(actual.Url))
		return true
	}
	if actual.Label == nil && desired.Label != nil && !dcl.IsEmptyValueIndirect(desired.Label) {
		c.Config.Logger.Infof("desired Label %s - but actually nil", dcl.SprintResource(desired.Label))
		return true
	}
	if !dcl.StringCanonicalize(desired.Label, actual.Label) && !dcl.IsZeroValue(desired.Label) {
		c.Config.Logger.Infof("Diff in Label. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Label), dcl.SprintResource(actual.Label))
		return true
	}
	return false
}

func compareNoteRelatedUrlSlice(c *Client, desired, actual []NoteRelatedUrl) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteRelatedUrl, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteRelatedUrl(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteRelatedUrl, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteRelatedUrlMap(c *Client, desired, actual map[string]NoteRelatedUrl) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteRelatedUrl, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteRelatedUrl, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteRelatedUrl(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteRelatedUrl, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteVulnerability(c *Client, desired, actual *NoteVulnerability) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.CvssScore == nil && desired.CvssScore != nil && !dcl.IsEmptyValueIndirect(desired.CvssScore) {
		c.Config.Logger.Infof("desired CvssScore %s - but actually nil", dcl.SprintResource(desired.CvssScore))
		return true
	}
	if !reflect.DeepEqual(desired.CvssScore, actual.CvssScore) && !dcl.IsZeroValue(desired.CvssScore) {
		c.Config.Logger.Infof("Diff in CvssScore. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CvssScore), dcl.SprintResource(actual.CvssScore))
		return true
	}
	if actual.Severity == nil && desired.Severity != nil && !dcl.IsEmptyValueIndirect(desired.Severity) {
		c.Config.Logger.Infof("desired Severity %s - but actually nil", dcl.SprintResource(desired.Severity))
		return true
	}
	if !reflect.DeepEqual(desired.Severity, actual.Severity) && !dcl.IsZeroValue(desired.Severity) {
		c.Config.Logger.Infof("Diff in Severity. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Severity), dcl.SprintResource(actual.Severity))
		return true
	}
	if actual.Details == nil && desired.Details != nil && !dcl.IsEmptyValueIndirect(desired.Details) {
		c.Config.Logger.Infof("desired Details %s - but actually nil", dcl.SprintResource(desired.Details))
		return true
	}
	if compareNoteVulnerabilityDetailsSlice(c, desired.Details, actual.Details) && !dcl.IsZeroValue(desired.Details) {
		c.Config.Logger.Infof("Diff in Details. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Details), dcl.SprintResource(actual.Details))
		return true
	}
	if actual.CvssV3 == nil && desired.CvssV3 != nil && !dcl.IsEmptyValueIndirect(desired.CvssV3) {
		c.Config.Logger.Infof("desired CvssV3 %s - but actually nil", dcl.SprintResource(desired.CvssV3))
		return true
	}
	if compareNoteVulnerabilityCvssV3(c, desired.CvssV3, actual.CvssV3) && !dcl.IsZeroValue(desired.CvssV3) {
		c.Config.Logger.Infof("Diff in CvssV3. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CvssV3), dcl.SprintResource(actual.CvssV3))
		return true
	}
	if actual.WindowsDetails == nil && desired.WindowsDetails != nil && !dcl.IsEmptyValueIndirect(desired.WindowsDetails) {
		c.Config.Logger.Infof("desired WindowsDetails %s - but actually nil", dcl.SprintResource(desired.WindowsDetails))
		return true
	}
	if compareNoteVulnerabilityWindowsDetailsSlice(c, desired.WindowsDetails, actual.WindowsDetails) && !dcl.IsZeroValue(desired.WindowsDetails) {
		c.Config.Logger.Infof("Diff in WindowsDetails. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WindowsDetails), dcl.SprintResource(actual.WindowsDetails))
		return true
	}
	if actual.SourceUpdateTime == nil && desired.SourceUpdateTime != nil && !dcl.IsEmptyValueIndirect(desired.SourceUpdateTime) {
		c.Config.Logger.Infof("desired SourceUpdateTime %s - but actually nil", dcl.SprintResource(desired.SourceUpdateTime))
		return true
	}
	if !reflect.DeepEqual(desired.SourceUpdateTime, actual.SourceUpdateTime) && !dcl.IsZeroValue(desired.SourceUpdateTime) {
		c.Config.Logger.Infof("Diff in SourceUpdateTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceUpdateTime), dcl.SprintResource(actual.SourceUpdateTime))
		return true
	}
	return false
}

func compareNoteVulnerabilitySlice(c *Client, desired, actual []NoteVulnerability) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerability, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerability(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerability, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityMap(c *Client, desired, actual map[string]NoteVulnerability) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerability, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteVulnerability, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteVulnerability(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteVulnerability, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetails(c *Client, desired, actual *NoteVulnerabilityDetails) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.SeverityName == nil && desired.SeverityName != nil && !dcl.IsEmptyValueIndirect(desired.SeverityName) {
		c.Config.Logger.Infof("desired SeverityName %s - but actually nil", dcl.SprintResource(desired.SeverityName))
		return true
	}
	if !dcl.StringCanonicalize(desired.SeverityName, actual.SeverityName) && !dcl.IsZeroValue(desired.SeverityName) {
		c.Config.Logger.Infof("Diff in SeverityName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SeverityName), dcl.SprintResource(actual.SeverityName))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !dcl.StringCanonicalize(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.PackageType == nil && desired.PackageType != nil && !dcl.IsEmptyValueIndirect(desired.PackageType) {
		c.Config.Logger.Infof("desired PackageType %s - but actually nil", dcl.SprintResource(desired.PackageType))
		return true
	}
	if !dcl.StringCanonicalize(desired.PackageType, actual.PackageType) && !dcl.IsZeroValue(desired.PackageType) {
		c.Config.Logger.Infof("Diff in PackageType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PackageType), dcl.SprintResource(actual.PackageType))
		return true
	}
	if actual.AffectedCpeUri == nil && desired.AffectedCpeUri != nil && !dcl.IsEmptyValueIndirect(desired.AffectedCpeUri) {
		c.Config.Logger.Infof("desired AffectedCpeUri %s - but actually nil", dcl.SprintResource(desired.AffectedCpeUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.AffectedCpeUri, actual.AffectedCpeUri) && !dcl.IsZeroValue(desired.AffectedCpeUri) {
		c.Config.Logger.Infof("Diff in AffectedCpeUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AffectedCpeUri), dcl.SprintResource(actual.AffectedCpeUri))
		return true
	}
	if actual.AffectedPackage == nil && desired.AffectedPackage != nil && !dcl.IsEmptyValueIndirect(desired.AffectedPackage) {
		c.Config.Logger.Infof("desired AffectedPackage %s - but actually nil", dcl.SprintResource(desired.AffectedPackage))
		return true
	}
	if !dcl.StringCanonicalize(desired.AffectedPackage, actual.AffectedPackage) && !dcl.IsZeroValue(desired.AffectedPackage) {
		c.Config.Logger.Infof("Diff in AffectedPackage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AffectedPackage), dcl.SprintResource(actual.AffectedPackage))
		return true
	}
	if actual.AffectedVersionStart == nil && desired.AffectedVersionStart != nil && !dcl.IsEmptyValueIndirect(desired.AffectedVersionStart) {
		c.Config.Logger.Infof("desired AffectedVersionStart %s - but actually nil", dcl.SprintResource(desired.AffectedVersionStart))
		return true
	}
	if compareNoteVulnerabilityDetailsAffectedVersionStart(c, desired.AffectedVersionStart, actual.AffectedVersionStart) && !dcl.IsZeroValue(desired.AffectedVersionStart) {
		c.Config.Logger.Infof("Diff in AffectedVersionStart. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AffectedVersionStart), dcl.SprintResource(actual.AffectedVersionStart))
		return true
	}
	if actual.AffectedVersionEnd == nil && desired.AffectedVersionEnd != nil && !dcl.IsEmptyValueIndirect(desired.AffectedVersionEnd) {
		c.Config.Logger.Infof("desired AffectedVersionEnd %s - but actually nil", dcl.SprintResource(desired.AffectedVersionEnd))
		return true
	}
	if compareNoteVulnerabilityDetailsAffectedVersionEnd(c, desired.AffectedVersionEnd, actual.AffectedVersionEnd) && !dcl.IsZeroValue(desired.AffectedVersionEnd) {
		c.Config.Logger.Infof("Diff in AffectedVersionEnd. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AffectedVersionEnd), dcl.SprintResource(actual.AffectedVersionEnd))
		return true
	}
	if actual.FixedCpeUri == nil && desired.FixedCpeUri != nil && !dcl.IsEmptyValueIndirect(desired.FixedCpeUri) {
		c.Config.Logger.Infof("desired FixedCpeUri %s - but actually nil", dcl.SprintResource(desired.FixedCpeUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.FixedCpeUri, actual.FixedCpeUri) && !dcl.IsZeroValue(desired.FixedCpeUri) {
		c.Config.Logger.Infof("Diff in FixedCpeUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FixedCpeUri), dcl.SprintResource(actual.FixedCpeUri))
		return true
	}
	if actual.FixedPackage == nil && desired.FixedPackage != nil && !dcl.IsEmptyValueIndirect(desired.FixedPackage) {
		c.Config.Logger.Infof("desired FixedPackage %s - but actually nil", dcl.SprintResource(desired.FixedPackage))
		return true
	}
	if !dcl.StringCanonicalize(desired.FixedPackage, actual.FixedPackage) && !dcl.IsZeroValue(desired.FixedPackage) {
		c.Config.Logger.Infof("Diff in FixedPackage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FixedPackage), dcl.SprintResource(actual.FixedPackage))
		return true
	}
	if actual.FixedVersion == nil && desired.FixedVersion != nil && !dcl.IsEmptyValueIndirect(desired.FixedVersion) {
		c.Config.Logger.Infof("desired FixedVersion %s - but actually nil", dcl.SprintResource(desired.FixedVersion))
		return true
	}
	if compareNoteVulnerabilityDetailsFixedVersion(c, desired.FixedVersion, actual.FixedVersion) && !dcl.IsZeroValue(desired.FixedVersion) {
		c.Config.Logger.Infof("Diff in FixedVersion. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FixedVersion), dcl.SprintResource(actual.FixedVersion))
		return true
	}
	if actual.IsObsolete == nil && desired.IsObsolete != nil && !dcl.IsEmptyValueIndirect(desired.IsObsolete) {
		c.Config.Logger.Infof("desired IsObsolete %s - but actually nil", dcl.SprintResource(desired.IsObsolete))
		return true
	}
	if !dcl.BoolCanonicalize(desired.IsObsolete, actual.IsObsolete) && !dcl.IsZeroValue(desired.IsObsolete) {
		c.Config.Logger.Infof("Diff in IsObsolete. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IsObsolete), dcl.SprintResource(actual.IsObsolete))
		return true
	}
	if actual.SourceUpdateTime == nil && desired.SourceUpdateTime != nil && !dcl.IsEmptyValueIndirect(desired.SourceUpdateTime) {
		c.Config.Logger.Infof("desired SourceUpdateTime %s - but actually nil", dcl.SprintResource(desired.SourceUpdateTime))
		return true
	}
	if !reflect.DeepEqual(desired.SourceUpdateTime, actual.SourceUpdateTime) && !dcl.IsZeroValue(desired.SourceUpdateTime) {
		c.Config.Logger.Infof("Diff in SourceUpdateTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceUpdateTime), dcl.SprintResource(actual.SourceUpdateTime))
		return true
	}
	return false
}

func compareNoteVulnerabilityDetailsSlice(c *Client, desired, actual []NoteVulnerabilityDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetails, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityDetails(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetails, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsMap(c *Client, desired, actual map[string]NoteVulnerabilityDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetails, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetails, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteVulnerabilityDetails(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetails, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsAffectedVersionStart(c *Client, desired, actual *NoteVulnerabilityDetailsAffectedVersionStart) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Epoch == nil && desired.Epoch != nil && !dcl.IsEmptyValueIndirect(desired.Epoch) {
		c.Config.Logger.Infof("desired Epoch %s - but actually nil", dcl.SprintResource(desired.Epoch))
		return true
	}
	if !reflect.DeepEqual(desired.Epoch, actual.Epoch) && !dcl.IsZeroValue(desired.Epoch) {
		c.Config.Logger.Infof("Diff in Epoch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Epoch), dcl.SprintResource(actual.Epoch))
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Revision == nil && desired.Revision != nil && !dcl.IsEmptyValueIndirect(desired.Revision) {
		c.Config.Logger.Infof("desired Revision %s - but actually nil", dcl.SprintResource(desired.Revision))
		return true
	}
	if !dcl.StringCanonicalize(desired.Revision, actual.Revision) && !dcl.IsZeroValue(desired.Revision) {
		c.Config.Logger.Infof("Diff in Revision. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Revision), dcl.SprintResource(actual.Revision))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.FullName == nil && desired.FullName != nil && !dcl.IsEmptyValueIndirect(desired.FullName) {
		c.Config.Logger.Infof("desired FullName %s - but actually nil", dcl.SprintResource(desired.FullName))
		return true
	}
	if !dcl.StringCanonicalize(desired.FullName, actual.FullName) && !dcl.IsZeroValue(desired.FullName) {
		c.Config.Logger.Infof("Diff in FullName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FullName), dcl.SprintResource(actual.FullName))
		return true
	}
	return false
}

func compareNoteVulnerabilityDetailsAffectedVersionStartSlice(c *Client, desired, actual []NoteVulnerabilityDetailsAffectedVersionStart) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetailsAffectedVersionStart, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityDetailsAffectedVersionStart(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsAffectedVersionStart, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsAffectedVersionStartMap(c *Client, desired, actual map[string]NoteVulnerabilityDetailsAffectedVersionStart) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetailsAffectedVersionStart, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsAffectedVersionStart, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteVulnerabilityDetailsAffectedVersionStart(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsAffectedVersionStart, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsAffectedVersionEnd(c *Client, desired, actual *NoteVulnerabilityDetailsAffectedVersionEnd) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Epoch == nil && desired.Epoch != nil && !dcl.IsEmptyValueIndirect(desired.Epoch) {
		c.Config.Logger.Infof("desired Epoch %s - but actually nil", dcl.SprintResource(desired.Epoch))
		return true
	}
	if !reflect.DeepEqual(desired.Epoch, actual.Epoch) && !dcl.IsZeroValue(desired.Epoch) {
		c.Config.Logger.Infof("Diff in Epoch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Epoch), dcl.SprintResource(actual.Epoch))
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Revision == nil && desired.Revision != nil && !dcl.IsEmptyValueIndirect(desired.Revision) {
		c.Config.Logger.Infof("desired Revision %s - but actually nil", dcl.SprintResource(desired.Revision))
		return true
	}
	if !dcl.StringCanonicalize(desired.Revision, actual.Revision) && !dcl.IsZeroValue(desired.Revision) {
		c.Config.Logger.Infof("Diff in Revision. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Revision), dcl.SprintResource(actual.Revision))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.FullName == nil && desired.FullName != nil && !dcl.IsEmptyValueIndirect(desired.FullName) {
		c.Config.Logger.Infof("desired FullName %s - but actually nil", dcl.SprintResource(desired.FullName))
		return true
	}
	if !dcl.StringCanonicalize(desired.FullName, actual.FullName) && !dcl.IsZeroValue(desired.FullName) {
		c.Config.Logger.Infof("Diff in FullName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FullName), dcl.SprintResource(actual.FullName))
		return true
	}
	return false
}

func compareNoteVulnerabilityDetailsAffectedVersionEndSlice(c *Client, desired, actual []NoteVulnerabilityDetailsAffectedVersionEnd) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetailsAffectedVersionEnd, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityDetailsAffectedVersionEnd(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsAffectedVersionEnd, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsAffectedVersionEndMap(c *Client, desired, actual map[string]NoteVulnerabilityDetailsAffectedVersionEnd) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetailsAffectedVersionEnd, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsAffectedVersionEnd, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteVulnerabilityDetailsAffectedVersionEnd(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsAffectedVersionEnd, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsFixedVersion(c *Client, desired, actual *NoteVulnerabilityDetailsFixedVersion) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Epoch == nil && desired.Epoch != nil && !dcl.IsEmptyValueIndirect(desired.Epoch) {
		c.Config.Logger.Infof("desired Epoch %s - but actually nil", dcl.SprintResource(desired.Epoch))
		return true
	}
	if !reflect.DeepEqual(desired.Epoch, actual.Epoch) && !dcl.IsZeroValue(desired.Epoch) {
		c.Config.Logger.Infof("Diff in Epoch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Epoch), dcl.SprintResource(actual.Epoch))
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Revision == nil && desired.Revision != nil && !dcl.IsEmptyValueIndirect(desired.Revision) {
		c.Config.Logger.Infof("desired Revision %s - but actually nil", dcl.SprintResource(desired.Revision))
		return true
	}
	if !dcl.StringCanonicalize(desired.Revision, actual.Revision) && !dcl.IsZeroValue(desired.Revision) {
		c.Config.Logger.Infof("Diff in Revision. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Revision), dcl.SprintResource(actual.Revision))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.FullName == nil && desired.FullName != nil && !dcl.IsEmptyValueIndirect(desired.FullName) {
		c.Config.Logger.Infof("desired FullName %s - but actually nil", dcl.SprintResource(desired.FullName))
		return true
	}
	if !dcl.StringCanonicalize(desired.FullName, actual.FullName) && !dcl.IsZeroValue(desired.FullName) {
		c.Config.Logger.Infof("Diff in FullName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FullName), dcl.SprintResource(actual.FullName))
		return true
	}
	return false
}

func compareNoteVulnerabilityDetailsFixedVersionSlice(c *Client, desired, actual []NoteVulnerabilityDetailsFixedVersion) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetailsFixedVersion, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityDetailsFixedVersion(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsFixedVersion, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsFixedVersionMap(c *Client, desired, actual map[string]NoteVulnerabilityDetailsFixedVersion) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetailsFixedVersion, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsFixedVersion, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteVulnerabilityDetailsFixedVersion(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsFixedVersion, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3(c *Client, desired, actual *NoteVulnerabilityCvssV3) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BaseScore == nil && desired.BaseScore != nil && !dcl.IsEmptyValueIndirect(desired.BaseScore) {
		c.Config.Logger.Infof("desired BaseScore %s - but actually nil", dcl.SprintResource(desired.BaseScore))
		return true
	}
	if !reflect.DeepEqual(desired.BaseScore, actual.BaseScore) && !dcl.IsZeroValue(desired.BaseScore) {
		c.Config.Logger.Infof("Diff in BaseScore. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BaseScore), dcl.SprintResource(actual.BaseScore))
		return true
	}
	if actual.ExploitabilityScore == nil && desired.ExploitabilityScore != nil && !dcl.IsEmptyValueIndirect(desired.ExploitabilityScore) {
		c.Config.Logger.Infof("desired ExploitabilityScore %s - but actually nil", dcl.SprintResource(desired.ExploitabilityScore))
		return true
	}
	if !reflect.DeepEqual(desired.ExploitabilityScore, actual.ExploitabilityScore) && !dcl.IsZeroValue(desired.ExploitabilityScore) {
		c.Config.Logger.Infof("Diff in ExploitabilityScore. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExploitabilityScore), dcl.SprintResource(actual.ExploitabilityScore))
		return true
	}
	if actual.ImpactScore == nil && desired.ImpactScore != nil && !dcl.IsEmptyValueIndirect(desired.ImpactScore) {
		c.Config.Logger.Infof("desired ImpactScore %s - but actually nil", dcl.SprintResource(desired.ImpactScore))
		return true
	}
	if !reflect.DeepEqual(desired.ImpactScore, actual.ImpactScore) && !dcl.IsZeroValue(desired.ImpactScore) {
		c.Config.Logger.Infof("Diff in ImpactScore. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ImpactScore), dcl.SprintResource(actual.ImpactScore))
		return true
	}
	if actual.AttackVector == nil && desired.AttackVector != nil && !dcl.IsEmptyValueIndirect(desired.AttackVector) {
		c.Config.Logger.Infof("desired AttackVector %s - but actually nil", dcl.SprintResource(desired.AttackVector))
		return true
	}
	if !reflect.DeepEqual(desired.AttackVector, actual.AttackVector) && !dcl.IsZeroValue(desired.AttackVector) {
		c.Config.Logger.Infof("Diff in AttackVector. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AttackVector), dcl.SprintResource(actual.AttackVector))
		return true
	}
	if actual.AttackComplexity == nil && desired.AttackComplexity != nil && !dcl.IsEmptyValueIndirect(desired.AttackComplexity) {
		c.Config.Logger.Infof("desired AttackComplexity %s - but actually nil", dcl.SprintResource(desired.AttackComplexity))
		return true
	}
	if !reflect.DeepEqual(desired.AttackComplexity, actual.AttackComplexity) && !dcl.IsZeroValue(desired.AttackComplexity) {
		c.Config.Logger.Infof("Diff in AttackComplexity. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AttackComplexity), dcl.SprintResource(actual.AttackComplexity))
		return true
	}
	if actual.PrivilegesRequired == nil && desired.PrivilegesRequired != nil && !dcl.IsEmptyValueIndirect(desired.PrivilegesRequired) {
		c.Config.Logger.Infof("desired PrivilegesRequired %s - but actually nil", dcl.SprintResource(desired.PrivilegesRequired))
		return true
	}
	if !reflect.DeepEqual(desired.PrivilegesRequired, actual.PrivilegesRequired) && !dcl.IsZeroValue(desired.PrivilegesRequired) {
		c.Config.Logger.Infof("Diff in PrivilegesRequired. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrivilegesRequired), dcl.SprintResource(actual.PrivilegesRequired))
		return true
	}
	if actual.UserInteraction == nil && desired.UserInteraction != nil && !dcl.IsEmptyValueIndirect(desired.UserInteraction) {
		c.Config.Logger.Infof("desired UserInteraction %s - but actually nil", dcl.SprintResource(desired.UserInteraction))
		return true
	}
	if !reflect.DeepEqual(desired.UserInteraction, actual.UserInteraction) && !dcl.IsZeroValue(desired.UserInteraction) {
		c.Config.Logger.Infof("Diff in UserInteraction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UserInteraction), dcl.SprintResource(actual.UserInteraction))
		return true
	}
	if actual.Scope == nil && desired.Scope != nil && !dcl.IsEmptyValueIndirect(desired.Scope) {
		c.Config.Logger.Infof("desired Scope %s - but actually nil", dcl.SprintResource(desired.Scope))
		return true
	}
	if !reflect.DeepEqual(desired.Scope, actual.Scope) && !dcl.IsZeroValue(desired.Scope) {
		c.Config.Logger.Infof("Diff in Scope. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scope), dcl.SprintResource(actual.Scope))
		return true
	}
	if actual.ConfidentialityImpact == nil && desired.ConfidentialityImpact != nil && !dcl.IsEmptyValueIndirect(desired.ConfidentialityImpact) {
		c.Config.Logger.Infof("desired ConfidentialityImpact %s - but actually nil", dcl.SprintResource(desired.ConfidentialityImpact))
		return true
	}
	if !reflect.DeepEqual(desired.ConfidentialityImpact, actual.ConfidentialityImpact) && !dcl.IsZeroValue(desired.ConfidentialityImpact) {
		c.Config.Logger.Infof("Diff in ConfidentialityImpact. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConfidentialityImpact), dcl.SprintResource(actual.ConfidentialityImpact))
		return true
	}
	if actual.IntegrityImpact == nil && desired.IntegrityImpact != nil && !dcl.IsEmptyValueIndirect(desired.IntegrityImpact) {
		c.Config.Logger.Infof("desired IntegrityImpact %s - but actually nil", dcl.SprintResource(desired.IntegrityImpact))
		return true
	}
	if !reflect.DeepEqual(desired.IntegrityImpact, actual.IntegrityImpact) && !dcl.IsZeroValue(desired.IntegrityImpact) {
		c.Config.Logger.Infof("Diff in IntegrityImpact. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IntegrityImpact), dcl.SprintResource(actual.IntegrityImpact))
		return true
	}
	if actual.AvailabilityImpact == nil && desired.AvailabilityImpact != nil && !dcl.IsEmptyValueIndirect(desired.AvailabilityImpact) {
		c.Config.Logger.Infof("desired AvailabilityImpact %s - but actually nil", dcl.SprintResource(desired.AvailabilityImpact))
		return true
	}
	if !reflect.DeepEqual(desired.AvailabilityImpact, actual.AvailabilityImpact) && !dcl.IsZeroValue(desired.AvailabilityImpact) {
		c.Config.Logger.Infof("Diff in AvailabilityImpact. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AvailabilityImpact), dcl.SprintResource(actual.AvailabilityImpact))
		return true
	}
	return false
}

func compareNoteVulnerabilityCvssV3Slice(c *Client, desired, actual []NoteVulnerabilityCvssV3) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityCvssV3(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3Map(c *Client, desired, actual map[string]NoteVulnerabilityCvssV3) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteVulnerabilityCvssV3(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityWindowsDetails(c *Client, desired, actual *NoteVulnerabilityWindowsDetails) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.CpeUri == nil && desired.CpeUri != nil && !dcl.IsEmptyValueIndirect(desired.CpeUri) {
		c.Config.Logger.Infof("desired CpeUri %s - but actually nil", dcl.SprintResource(desired.CpeUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.CpeUri, actual.CpeUri) && !dcl.IsZeroValue(desired.CpeUri) {
		c.Config.Logger.Infof("Diff in CpeUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CpeUri), dcl.SprintResource(actual.CpeUri))
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !dcl.StringCanonicalize(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.FixingKbs == nil && desired.FixingKbs != nil && !dcl.IsEmptyValueIndirect(desired.FixingKbs) {
		c.Config.Logger.Infof("desired FixingKbs %s - but actually nil", dcl.SprintResource(desired.FixingKbs))
		return true
	}
	if compareNoteVulnerabilityWindowsDetailsFixingKbsSlice(c, desired.FixingKbs, actual.FixingKbs) && !dcl.IsZeroValue(desired.FixingKbs) {
		c.Config.Logger.Infof("Diff in FixingKbs. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FixingKbs), dcl.SprintResource(actual.FixingKbs))
		return true
	}
	return false
}

func compareNoteVulnerabilityWindowsDetailsSlice(c *Client, desired, actual []NoteVulnerabilityWindowsDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityWindowsDetails, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityWindowsDetails(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityWindowsDetails, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityWindowsDetailsMap(c *Client, desired, actual map[string]NoteVulnerabilityWindowsDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityWindowsDetails, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityWindowsDetails, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteVulnerabilityWindowsDetails(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityWindowsDetails, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityWindowsDetailsFixingKbs(c *Client, desired, actual *NoteVulnerabilityWindowsDetailsFixingKbs) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Url == nil && desired.Url != nil && !dcl.IsEmptyValueIndirect(desired.Url) {
		c.Config.Logger.Infof("desired Url %s - but actually nil", dcl.SprintResource(desired.Url))
		return true
	}
	if !dcl.StringCanonicalize(desired.Url, actual.Url) && !dcl.IsZeroValue(desired.Url) {
		c.Config.Logger.Infof("Diff in Url. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Url), dcl.SprintResource(actual.Url))
		return true
	}
	return false
}

func compareNoteVulnerabilityWindowsDetailsFixingKbsSlice(c *Client, desired, actual []NoteVulnerabilityWindowsDetailsFixingKbs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityWindowsDetailsFixingKbs, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityWindowsDetailsFixingKbs(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityWindowsDetailsFixingKbs, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityWindowsDetailsFixingKbsMap(c *Client, desired, actual map[string]NoteVulnerabilityWindowsDetailsFixingKbs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityWindowsDetailsFixingKbs, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityWindowsDetailsFixingKbs, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteVulnerabilityWindowsDetailsFixingKbs(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityWindowsDetailsFixingKbs, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteBuild(c *Client, desired, actual *NoteBuild) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BuilderVersion == nil && desired.BuilderVersion != nil && !dcl.IsEmptyValueIndirect(desired.BuilderVersion) {
		c.Config.Logger.Infof("desired BuilderVersion %s - but actually nil", dcl.SprintResource(desired.BuilderVersion))
		return true
	}
	if !dcl.StringCanonicalize(desired.BuilderVersion, actual.BuilderVersion) && !dcl.IsZeroValue(desired.BuilderVersion) {
		c.Config.Logger.Infof("Diff in BuilderVersion. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BuilderVersion), dcl.SprintResource(actual.BuilderVersion))
		return true
	}
	return false
}

func compareNoteBuildSlice(c *Client, desired, actual []NoteBuild) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteBuild, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteBuild(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteBuild, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteBuildMap(c *Client, desired, actual map[string]NoteBuild) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteBuild, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteBuild, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteBuild(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteBuild, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteImage(c *Client, desired, actual *NoteImage) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ResourceUrl == nil && desired.ResourceUrl != nil && !dcl.IsEmptyValueIndirect(desired.ResourceUrl) {
		c.Config.Logger.Infof("desired ResourceUrl %s - but actually nil", dcl.SprintResource(desired.ResourceUrl))
		return true
	}
	if !dcl.StringCanonicalize(desired.ResourceUrl, actual.ResourceUrl) && !dcl.IsZeroValue(desired.ResourceUrl) {
		c.Config.Logger.Infof("Diff in ResourceUrl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceUrl), dcl.SprintResource(actual.ResourceUrl))
		return true
	}
	if actual.Fingerprint == nil && desired.Fingerprint != nil && !dcl.IsEmptyValueIndirect(desired.Fingerprint) {
		c.Config.Logger.Infof("desired Fingerprint %s - but actually nil", dcl.SprintResource(desired.Fingerprint))
		return true
	}
	if compareNoteImageFingerprint(c, desired.Fingerprint, actual.Fingerprint) && !dcl.IsZeroValue(desired.Fingerprint) {
		c.Config.Logger.Infof("Diff in Fingerprint. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Fingerprint), dcl.SprintResource(actual.Fingerprint))
		return true
	}
	return false
}

func compareNoteImageSlice(c *Client, desired, actual []NoteImage) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteImage, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteImage(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteImage, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteImageMap(c *Client, desired, actual map[string]NoteImage) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteImage, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteImage, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteImage(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteImage, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteImageFingerprint(c *Client, desired, actual *NoteImageFingerprint) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.V1Name == nil && desired.V1Name != nil && !dcl.IsEmptyValueIndirect(desired.V1Name) {
		c.Config.Logger.Infof("desired V1Name %s - but actually nil", dcl.SprintResource(desired.V1Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.V1Name, actual.V1Name) && !dcl.IsZeroValue(desired.V1Name) {
		c.Config.Logger.Infof("Diff in V1Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.V1Name), dcl.SprintResource(actual.V1Name))
		return true
	}
	if actual.V2Blob == nil && desired.V2Blob != nil && !dcl.IsEmptyValueIndirect(desired.V2Blob) {
		c.Config.Logger.Infof("desired V2Blob %s - but actually nil", dcl.SprintResource(desired.V2Blob))
		return true
	}
	if !dcl.StringSliceEquals(desired.V2Blob, actual.V2Blob) && !dcl.IsZeroValue(desired.V2Blob) {
		c.Config.Logger.Infof("Diff in V2Blob. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.V2Blob), dcl.SprintResource(actual.V2Blob))
		return true
	}
	return false
}

func compareNoteImageFingerprintSlice(c *Client, desired, actual []NoteImageFingerprint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteImageFingerprint, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteImageFingerprint(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteImageFingerprint, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteImageFingerprintMap(c *Client, desired, actual map[string]NoteImageFingerprint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteImageFingerprint, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteImageFingerprint, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteImageFingerprint(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteImageFingerprint, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNotePackage(c *Client, desired, actual *NotePackage) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Distribution == nil && desired.Distribution != nil && !dcl.IsEmptyValueIndirect(desired.Distribution) {
		c.Config.Logger.Infof("desired Distribution %s - but actually nil", dcl.SprintResource(desired.Distribution))
		return true
	}
	if compareNotePackageDistributionSlice(c, desired.Distribution, actual.Distribution) && !dcl.IsZeroValue(desired.Distribution) {
		c.Config.Logger.Infof("Diff in Distribution. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Distribution), dcl.SprintResource(actual.Distribution))
		return true
	}
	return false
}

func compareNotePackageSlice(c *Client, desired, actual []NotePackage) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NotePackage, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNotePackage(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NotePackage, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNotePackageMap(c *Client, desired, actual map[string]NotePackage) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NotePackage, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NotePackage, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNotePackage(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NotePackage, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNotePackageDistribution(c *Client, desired, actual *NotePackageDistribution) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.CpeUri == nil && desired.CpeUri != nil && !dcl.IsEmptyValueIndirect(desired.CpeUri) {
		c.Config.Logger.Infof("desired CpeUri %s - but actually nil", dcl.SprintResource(desired.CpeUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.CpeUri, actual.CpeUri) && !dcl.IsZeroValue(desired.CpeUri) {
		c.Config.Logger.Infof("Diff in CpeUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CpeUri), dcl.SprintResource(actual.CpeUri))
		return true
	}
	if actual.Architecture == nil && desired.Architecture != nil && !dcl.IsEmptyValueIndirect(desired.Architecture) {
		c.Config.Logger.Infof("desired Architecture %s - but actually nil", dcl.SprintResource(desired.Architecture))
		return true
	}
	if !reflect.DeepEqual(desired.Architecture, actual.Architecture) && !dcl.IsZeroValue(desired.Architecture) {
		c.Config.Logger.Infof("Diff in Architecture. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Architecture), dcl.SprintResource(actual.Architecture))
		return true
	}
	if actual.LatestVersion == nil && desired.LatestVersion != nil && !dcl.IsEmptyValueIndirect(desired.LatestVersion) {
		c.Config.Logger.Infof("desired LatestVersion %s - but actually nil", dcl.SprintResource(desired.LatestVersion))
		return true
	}
	if compareNotePackageDistributionLatestVersion(c, desired.LatestVersion, actual.LatestVersion) && !dcl.IsZeroValue(desired.LatestVersion) {
		c.Config.Logger.Infof("Diff in LatestVersion. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LatestVersion), dcl.SprintResource(actual.LatestVersion))
		return true
	}
	if actual.Maintainer == nil && desired.Maintainer != nil && !dcl.IsEmptyValueIndirect(desired.Maintainer) {
		c.Config.Logger.Infof("desired Maintainer %s - but actually nil", dcl.SprintResource(desired.Maintainer))
		return true
	}
	if !dcl.StringCanonicalize(desired.Maintainer, actual.Maintainer) && !dcl.IsZeroValue(desired.Maintainer) {
		c.Config.Logger.Infof("Diff in Maintainer. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Maintainer), dcl.SprintResource(actual.Maintainer))
		return true
	}
	if actual.Url == nil && desired.Url != nil && !dcl.IsEmptyValueIndirect(desired.Url) {
		c.Config.Logger.Infof("desired Url %s - but actually nil", dcl.SprintResource(desired.Url))
		return true
	}
	if !dcl.StringCanonicalize(desired.Url, actual.Url) && !dcl.IsZeroValue(desired.Url) {
		c.Config.Logger.Infof("Diff in Url. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Url), dcl.SprintResource(actual.Url))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !dcl.StringCanonicalize(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	return false
}

func compareNotePackageDistributionSlice(c *Client, desired, actual []NotePackageDistribution) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NotePackageDistribution, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNotePackageDistribution(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NotePackageDistribution, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNotePackageDistributionMap(c *Client, desired, actual map[string]NotePackageDistribution) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NotePackageDistribution, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NotePackageDistribution, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNotePackageDistribution(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NotePackageDistribution, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNotePackageDistributionLatestVersion(c *Client, desired, actual *NotePackageDistributionLatestVersion) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Epoch == nil && desired.Epoch != nil && !dcl.IsEmptyValueIndirect(desired.Epoch) {
		c.Config.Logger.Infof("desired Epoch %s - but actually nil", dcl.SprintResource(desired.Epoch))
		return true
	}
	if !reflect.DeepEqual(desired.Epoch, actual.Epoch) && !dcl.IsZeroValue(desired.Epoch) {
		c.Config.Logger.Infof("Diff in Epoch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Epoch), dcl.SprintResource(actual.Epoch))
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Revision == nil && desired.Revision != nil && !dcl.IsEmptyValueIndirect(desired.Revision) {
		c.Config.Logger.Infof("desired Revision %s - but actually nil", dcl.SprintResource(desired.Revision))
		return true
	}
	if !dcl.StringCanonicalize(desired.Revision, actual.Revision) && !dcl.IsZeroValue(desired.Revision) {
		c.Config.Logger.Infof("Diff in Revision. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Revision), dcl.SprintResource(actual.Revision))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.FullName == nil && desired.FullName != nil && !dcl.IsEmptyValueIndirect(desired.FullName) {
		c.Config.Logger.Infof("desired FullName %s - but actually nil", dcl.SprintResource(desired.FullName))
		return true
	}
	if !dcl.StringCanonicalize(desired.FullName, actual.FullName) && !dcl.IsZeroValue(desired.FullName) {
		c.Config.Logger.Infof("Diff in FullName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FullName), dcl.SprintResource(actual.FullName))
		return true
	}
	return false
}

func compareNotePackageDistributionLatestVersionSlice(c *Client, desired, actual []NotePackageDistributionLatestVersion) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NotePackageDistributionLatestVersion, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNotePackageDistributionLatestVersion(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NotePackageDistributionLatestVersion, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNotePackageDistributionLatestVersionMap(c *Client, desired, actual map[string]NotePackageDistributionLatestVersion) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NotePackageDistributionLatestVersion, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NotePackageDistributionLatestVersion, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNotePackageDistributionLatestVersion(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NotePackageDistributionLatestVersion, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteDiscovery(c *Client, desired, actual *NoteDiscovery) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AnalysisKind == nil && desired.AnalysisKind != nil && !dcl.IsEmptyValueIndirect(desired.AnalysisKind) {
		c.Config.Logger.Infof("desired AnalysisKind %s - but actually nil", dcl.SprintResource(desired.AnalysisKind))
		return true
	}
	if !reflect.DeepEqual(desired.AnalysisKind, actual.AnalysisKind) && !dcl.IsZeroValue(desired.AnalysisKind) {
		c.Config.Logger.Infof("Diff in AnalysisKind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AnalysisKind), dcl.SprintResource(actual.AnalysisKind))
		return true
	}
	return false
}

func compareNoteDiscoverySlice(c *Client, desired, actual []NoteDiscovery) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteDiscovery, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteDiscovery(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteDiscovery, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteDiscoveryMap(c *Client, desired, actual map[string]NoteDiscovery) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteDiscovery, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteDiscovery, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteDiscovery(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteDiscovery, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteDeployment(c *Client, desired, actual *NoteDeployment) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ResourceUri == nil && desired.ResourceUri != nil && !dcl.IsEmptyValueIndirect(desired.ResourceUri) {
		c.Config.Logger.Infof("desired ResourceUri %s - but actually nil", dcl.SprintResource(desired.ResourceUri))
		return true
	}
	if !dcl.StringSliceEquals(desired.ResourceUri, actual.ResourceUri) && !dcl.IsZeroValue(desired.ResourceUri) {
		c.Config.Logger.Infof("Diff in ResourceUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceUri), dcl.SprintResource(actual.ResourceUri))
		return true
	}
	return false
}

func compareNoteDeploymentSlice(c *Client, desired, actual []NoteDeployment) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteDeployment, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteDeployment(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteDeployment, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteDeploymentMap(c *Client, desired, actual map[string]NoteDeployment) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteDeployment, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteDeployment, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteDeployment(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteDeployment, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteAttestation(c *Client, desired, actual *NoteAttestation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Hint == nil && desired.Hint != nil && !dcl.IsEmptyValueIndirect(desired.Hint) {
		c.Config.Logger.Infof("desired Hint %s - but actually nil", dcl.SprintResource(desired.Hint))
		return true
	}
	if compareNoteAttestationHint(c, desired.Hint, actual.Hint) && !dcl.IsZeroValue(desired.Hint) {
		c.Config.Logger.Infof("Diff in Hint. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Hint), dcl.SprintResource(actual.Hint))
		return true
	}
	return false
}

func compareNoteAttestationSlice(c *Client, desired, actual []NoteAttestation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteAttestation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteAttestation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteAttestation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteAttestationMap(c *Client, desired, actual map[string]NoteAttestation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteAttestation, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteAttestation, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteAttestation(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteAttestation, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteAttestationHint(c *Client, desired, actual *NoteAttestationHint) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HumanReadableName == nil && desired.HumanReadableName != nil && !dcl.IsEmptyValueIndirect(desired.HumanReadableName) {
		c.Config.Logger.Infof("desired HumanReadableName %s - but actually nil", dcl.SprintResource(desired.HumanReadableName))
		return true
	}
	if !dcl.StringCanonicalize(desired.HumanReadableName, actual.HumanReadableName) && !dcl.IsZeroValue(desired.HumanReadableName) {
		c.Config.Logger.Infof("Diff in HumanReadableName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HumanReadableName), dcl.SprintResource(actual.HumanReadableName))
		return true
	}
	return false
}

func compareNoteAttestationHintSlice(c *Client, desired, actual []NoteAttestationHint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteAttestationHint, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteAttestationHint(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteAttestationHint, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteAttestationHintMap(c *Client, desired, actual map[string]NoteAttestationHint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteAttestationHint, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NoteAttestationHint, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNoteAttestationHint(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NoteAttestationHint, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilitySeverityEnumSlice(c *Client, desired, actual []NoteVulnerabilitySeverityEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilitySeverityEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilitySeverityEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilitySeverityEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilitySeverityEnum(c *Client, desired, actual *NoteVulnerabilitySeverityEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityDetailsAffectedVersionStartKindEnumSlice(c *Client, desired, actual []NoteVulnerabilityDetailsAffectedVersionStartKindEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetailsAffectedVersionStartKindEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityDetailsAffectedVersionStartKindEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsAffectedVersionStartKindEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsAffectedVersionStartKindEnum(c *Client, desired, actual *NoteVulnerabilityDetailsAffectedVersionStartKindEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityDetailsAffectedVersionEndKindEnumSlice(c *Client, desired, actual []NoteVulnerabilityDetailsAffectedVersionEndKindEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetailsAffectedVersionEndKindEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityDetailsAffectedVersionEndKindEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsAffectedVersionEndKindEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsAffectedVersionEndKindEnum(c *Client, desired, actual *NoteVulnerabilityDetailsAffectedVersionEndKindEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityDetailsFixedVersionKindEnumSlice(c *Client, desired, actual []NoteVulnerabilityDetailsFixedVersionKindEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityDetailsFixedVersionKindEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityDetailsFixedVersionKindEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityDetailsFixedVersionKindEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityDetailsFixedVersionKindEnum(c *Client, desired, actual *NoteVulnerabilityDetailsFixedVersionKindEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityCvssV3AttackVectorEnumSlice(c *Client, desired, actual []NoteVulnerabilityCvssV3AttackVectorEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3AttackVectorEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityCvssV3AttackVectorEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3AttackVectorEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3AttackVectorEnum(c *Client, desired, actual *NoteVulnerabilityCvssV3AttackVectorEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityCvssV3AttackComplexityEnumSlice(c *Client, desired, actual []NoteVulnerabilityCvssV3AttackComplexityEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3AttackComplexityEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityCvssV3AttackComplexityEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3AttackComplexityEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3AttackComplexityEnum(c *Client, desired, actual *NoteVulnerabilityCvssV3AttackComplexityEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityCvssV3PrivilegesRequiredEnumSlice(c *Client, desired, actual []NoteVulnerabilityCvssV3PrivilegesRequiredEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3PrivilegesRequiredEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityCvssV3PrivilegesRequiredEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3PrivilegesRequiredEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3PrivilegesRequiredEnum(c *Client, desired, actual *NoteVulnerabilityCvssV3PrivilegesRequiredEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityCvssV3UserInteractionEnumSlice(c *Client, desired, actual []NoteVulnerabilityCvssV3UserInteractionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3UserInteractionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityCvssV3UserInteractionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3UserInteractionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3UserInteractionEnum(c *Client, desired, actual *NoteVulnerabilityCvssV3UserInteractionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityCvssV3ScopeEnumSlice(c *Client, desired, actual []NoteVulnerabilityCvssV3ScopeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3ScopeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityCvssV3ScopeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3ScopeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3ScopeEnum(c *Client, desired, actual *NoteVulnerabilityCvssV3ScopeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityCvssV3ConfidentialityImpactEnumSlice(c *Client, desired, actual []NoteVulnerabilityCvssV3ConfidentialityImpactEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3ConfidentialityImpactEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityCvssV3ConfidentialityImpactEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3ConfidentialityImpactEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3ConfidentialityImpactEnum(c *Client, desired, actual *NoteVulnerabilityCvssV3ConfidentialityImpactEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityCvssV3IntegrityImpactEnumSlice(c *Client, desired, actual []NoteVulnerabilityCvssV3IntegrityImpactEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3IntegrityImpactEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityCvssV3IntegrityImpactEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3IntegrityImpactEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3IntegrityImpactEnum(c *Client, desired, actual *NoteVulnerabilityCvssV3IntegrityImpactEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteVulnerabilityCvssV3AvailabilityImpactEnumSlice(c *Client, desired, actual []NoteVulnerabilityCvssV3AvailabilityImpactEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteVulnerabilityCvssV3AvailabilityImpactEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteVulnerabilityCvssV3AvailabilityImpactEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteVulnerabilityCvssV3AvailabilityImpactEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteVulnerabilityCvssV3AvailabilityImpactEnum(c *Client, desired, actual *NoteVulnerabilityCvssV3AvailabilityImpactEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNotePackageDistributionArchitectureEnumSlice(c *Client, desired, actual []NotePackageDistributionArchitectureEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NotePackageDistributionArchitectureEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNotePackageDistributionArchitectureEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NotePackageDistributionArchitectureEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNotePackageDistributionArchitectureEnum(c *Client, desired, actual *NotePackageDistributionArchitectureEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNotePackageDistributionLatestVersionKindEnumSlice(c *Client, desired, actual []NotePackageDistributionLatestVersionKindEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NotePackageDistributionLatestVersionKindEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNotePackageDistributionLatestVersionKindEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NotePackageDistributionLatestVersionKindEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNotePackageDistributionLatestVersionKindEnum(c *Client, desired, actual *NotePackageDistributionLatestVersionKindEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNoteDiscoveryAnalysisKindEnumSlice(c *Client, desired, actual []NoteDiscoveryAnalysisKindEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NoteDiscoveryAnalysisKindEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNoteDiscoveryAnalysisKindEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NoteDiscoveryAnalysisKindEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNoteDiscoveryAnalysisKindEnum(c *Client, desired, actual *NoteDiscoveryAnalysisKindEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Note) urlNormalized() *Note {
	normalized := deepcopy.Copy(*r).(Note)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.ShortDescription = dcl.SelfLinkToName(r.ShortDescription)
	normalized.LongDescription = dcl.SelfLinkToName(r.LongDescription)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Note) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Note) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Note) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Note) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateNote" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/notes/{{name}}", "https://containeranalysis.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Note resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Note) marshal(c *Client) ([]byte, error) {
	m, err := expandNote(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Note: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalNote decodes JSON responses into the Note resource schema.
func unmarshalNote(b []byte, c *Client) (*Note, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapNote(m, c)
}

func unmarshalMapNote(m map[string]interface{}, c *Client) (*Note, error) {

	return flattenNote(c, m), nil
}

// expandNote expands Note into a JSON request object.
func expandNote(c *Client, f *Note) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.ShortDescription; !dcl.IsEmptyValueIndirect(v) {
		m["shortDescription"] = v
	}
	if v := f.LongDescription; !dcl.IsEmptyValueIndirect(v) {
		m["longDescription"] = v
	}
	if v, err := expandNoteRelatedUrlSlice(c, f.RelatedUrl); err != nil {
		return nil, fmt.Errorf("error expanding RelatedUrl into relatedUrl: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["relatedUrl"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		m["expirationTime"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.RelatedNoteNames; !dcl.IsEmptyValueIndirect(v) {
		m["relatedNoteNames"] = v
	}
	if v, err := expandNoteVulnerability(c, f.Vulnerability); err != nil {
		return nil, fmt.Errorf("error expanding Vulnerability into vulnerability: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["vulnerability"] = v
	}
	if v, err := expandNoteBuild(c, f.Build); err != nil {
		return nil, fmt.Errorf("error expanding Build into build: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["build"] = v
	}
	if v, err := expandNoteImage(c, f.Image); err != nil {
		return nil, fmt.Errorf("error expanding Image into image: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["image"] = v
	}
	if v, err := expandNotePackage(c, f.Package); err != nil {
		return nil, fmt.Errorf("error expanding Package into package: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["package"] = v
	}
	if v, err := expandNoteDiscovery(c, f.Discovery); err != nil {
		return nil, fmt.Errorf("error expanding Discovery into discovery: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["discovery"] = v
	}
	if v, err := expandNoteDeployment(c, f.Deployment); err != nil {
		return nil, fmt.Errorf("error expanding Deployment into deployment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["deployment"] = v
	}
	if v, err := expandNoteAttestation(c, f.Attestation); err != nil {
		return nil, fmt.Errorf("error expanding Attestation into attestation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["attestation"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenNote flattens Note from a JSON request object into the
// Note type.
func flattenNote(c *Client, i interface{}) *Note {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Note{}
	r.Name = dcl.FlattenString(m["name"])
	r.ShortDescription = dcl.FlattenString(m["shortDescription"])
	r.LongDescription = dcl.FlattenString(m["longDescription"])
	r.RelatedUrl = flattenNoteRelatedUrlSlice(c, m["relatedUrl"])
	r.ExpirationTime = dcl.FlattenString(m["expirationTime"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.RelatedNoteNames = dcl.FlattenStringSlice(m["relatedNoteNames"])
	r.Vulnerability = flattenNoteVulnerability(c, m["vulnerability"])
	r.Build = flattenNoteBuild(c, m["build"])
	r.Image = flattenNoteImage(c, m["image"])
	r.Package = flattenNotePackage(c, m["package"])
	r.Discovery = flattenNoteDiscovery(c, m["discovery"])
	r.Deployment = flattenNoteDeployment(c, m["deployment"])
	r.Attestation = flattenNoteAttestation(c, m["attestation"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandNoteRelatedUrlMap expands the contents of NoteRelatedUrl into a JSON
// request object.
func expandNoteRelatedUrlMap(c *Client, f map[string]NoteRelatedUrl) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteRelatedUrl(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteRelatedUrlSlice expands the contents of NoteRelatedUrl into a JSON
// request object.
func expandNoteRelatedUrlSlice(c *Client, f []NoteRelatedUrl) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteRelatedUrl(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteRelatedUrlMap flattens the contents of NoteRelatedUrl from a JSON
// response object.
func flattenNoteRelatedUrlMap(c *Client, i interface{}) map[string]NoteRelatedUrl {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteRelatedUrl{}
	}

	if len(a) == 0 {
		return map[string]NoteRelatedUrl{}
	}

	items := make(map[string]NoteRelatedUrl)
	for k, item := range a {
		items[k] = *flattenNoteRelatedUrl(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteRelatedUrlSlice flattens the contents of NoteRelatedUrl from a JSON
// response object.
func flattenNoteRelatedUrlSlice(c *Client, i interface{}) []NoteRelatedUrl {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteRelatedUrl{}
	}

	if len(a) == 0 {
		return []NoteRelatedUrl{}
	}

	items := make([]NoteRelatedUrl, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteRelatedUrl(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteRelatedUrl expands an instance of NoteRelatedUrl into a JSON
// request object.
func expandNoteRelatedUrl(c *Client, f *NoteRelatedUrl) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v := f.Label; !dcl.IsEmptyValueIndirect(v) {
		m["label"] = v
	}

	return m, nil
}

// flattenNoteRelatedUrl flattens an instance of NoteRelatedUrl from a JSON
// response object.
func flattenNoteRelatedUrl(c *Client, i interface{}) *NoteRelatedUrl {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteRelatedUrl{}
	r.Url = dcl.FlattenString(m["url"])
	r.Label = dcl.FlattenString(m["label"])

	return r
}

// expandNoteVulnerabilityMap expands the contents of NoteVulnerability into a JSON
// request object.
func expandNoteVulnerabilityMap(c *Client, f map[string]NoteVulnerability) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerability(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilitySlice expands the contents of NoteVulnerability into a JSON
// request object.
func expandNoteVulnerabilitySlice(c *Client, f []NoteVulnerability) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerability(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityMap flattens the contents of NoteVulnerability from a JSON
// response object.
func flattenNoteVulnerabilityMap(c *Client, i interface{}) map[string]NoteVulnerability {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerability{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerability{}
	}

	items := make(map[string]NoteVulnerability)
	for k, item := range a {
		items[k] = *flattenNoteVulnerability(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteVulnerabilitySlice flattens the contents of NoteVulnerability from a JSON
// response object.
func flattenNoteVulnerabilitySlice(c *Client, i interface{}) []NoteVulnerability {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerability{}
	}

	if len(a) == 0 {
		return []NoteVulnerability{}
	}

	items := make([]NoteVulnerability, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerability(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteVulnerability expands an instance of NoteVulnerability into a JSON
// request object.
func expandNoteVulnerability(c *Client, f *NoteVulnerability) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CvssScore; !dcl.IsEmptyValueIndirect(v) {
		m["cvssScore"] = v
	}
	if v := f.Severity; !dcl.IsEmptyValueIndirect(v) {
		m["severity"] = v
	}
	if v, err := expandNoteVulnerabilityDetailsSlice(c, f.Details); err != nil {
		return nil, fmt.Errorf("error expanding Details into details: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["details"] = v
	}
	if v, err := expandNoteVulnerabilityCvssV3(c, f.CvssV3); err != nil {
		return nil, fmt.Errorf("error expanding CvssV3 into cvssV3: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cvssV3"] = v
	}
	if v, err := expandNoteVulnerabilityWindowsDetailsSlice(c, f.WindowsDetails); err != nil {
		return nil, fmt.Errorf("error expanding WindowsDetails into windowsDetails: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["windowsDetails"] = v
	}
	if v := f.SourceUpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["sourceUpdateTime"] = v
	}

	return m, nil
}

// flattenNoteVulnerability flattens an instance of NoteVulnerability from a JSON
// response object.
func flattenNoteVulnerability(c *Client, i interface{}) *NoteVulnerability {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerability{}
	r.CvssScore = dcl.FlattenDouble(m["cvssScore"])
	r.Severity = flattenNoteVulnerabilitySeverityEnum(m["severity"])
	r.Details = flattenNoteVulnerabilityDetailsSlice(c, m["details"])
	r.CvssV3 = flattenNoteVulnerabilityCvssV3(c, m["cvssV3"])
	r.WindowsDetails = flattenNoteVulnerabilityWindowsDetailsSlice(c, m["windowsDetails"])
	r.SourceUpdateTime = dcl.FlattenString(m["sourceUpdateTime"])

	return r
}

// expandNoteVulnerabilityDetailsMap expands the contents of NoteVulnerabilityDetails into a JSON
// request object.
func expandNoteVulnerabilityDetailsMap(c *Client, f map[string]NoteVulnerabilityDetails) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityDetails(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityDetailsSlice expands the contents of NoteVulnerabilityDetails into a JSON
// request object.
func expandNoteVulnerabilityDetailsSlice(c *Client, f []NoteVulnerabilityDetails) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityDetails(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityDetailsMap flattens the contents of NoteVulnerabilityDetails from a JSON
// response object.
func flattenNoteVulnerabilityDetailsMap(c *Client, i interface{}) map[string]NoteVulnerabilityDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetails{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetails{}
	}

	items := make(map[string]NoteVulnerabilityDetails)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetails(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteVulnerabilityDetailsSlice flattens the contents of NoteVulnerabilityDetails from a JSON
// response object.
func flattenNoteVulnerabilityDetailsSlice(c *Client, i interface{}) []NoteVulnerabilityDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetails{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetails{}
	}

	items := make([]NoteVulnerabilityDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetails(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteVulnerabilityDetails expands an instance of NoteVulnerabilityDetails into a JSON
// request object.
func expandNoteVulnerabilityDetails(c *Client, f *NoteVulnerabilityDetails) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SeverityName; !dcl.IsEmptyValueIndirect(v) {
		m["severityName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.PackageType; !dcl.IsEmptyValueIndirect(v) {
		m["packageType"] = v
	}
	if v := f.AffectedCpeUri; !dcl.IsEmptyValueIndirect(v) {
		m["affectedCpeUri"] = v
	}
	if v := f.AffectedPackage; !dcl.IsEmptyValueIndirect(v) {
		m["affectedPackage"] = v
	}
	if v, err := expandNoteVulnerabilityDetailsAffectedVersionStart(c, f.AffectedVersionStart); err != nil {
		return nil, fmt.Errorf("error expanding AffectedVersionStart into affectedVersionStart: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["affectedVersionStart"] = v
	}
	if v, err := expandNoteVulnerabilityDetailsAffectedVersionEnd(c, f.AffectedVersionEnd); err != nil {
		return nil, fmt.Errorf("error expanding AffectedVersionEnd into affectedVersionEnd: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["affectedVersionEnd"] = v
	}
	if v := f.FixedCpeUri; !dcl.IsEmptyValueIndirect(v) {
		m["fixedCpeUri"] = v
	}
	if v := f.FixedPackage; !dcl.IsEmptyValueIndirect(v) {
		m["fixedPackage"] = v
	}
	if v, err := expandNoteVulnerabilityDetailsFixedVersion(c, f.FixedVersion); err != nil {
		return nil, fmt.Errorf("error expanding FixedVersion into fixedVersion: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fixedVersion"] = v
	}
	if v := f.IsObsolete; !dcl.IsEmptyValueIndirect(v) {
		m["isObsolete"] = v
	}
	if v := f.SourceUpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["sourceUpdateTime"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityDetails flattens an instance of NoteVulnerabilityDetails from a JSON
// response object.
func flattenNoteVulnerabilityDetails(c *Client, i interface{}) *NoteVulnerabilityDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityDetails{}
	r.SeverityName = dcl.FlattenString(m["severityName"])
	r.Description = dcl.FlattenString(m["description"])
	r.PackageType = dcl.FlattenString(m["packageType"])
	r.AffectedCpeUri = dcl.FlattenString(m["affectedCpeUri"])
	r.AffectedPackage = dcl.FlattenString(m["affectedPackage"])
	r.AffectedVersionStart = flattenNoteVulnerabilityDetailsAffectedVersionStart(c, m["affectedVersionStart"])
	r.AffectedVersionEnd = flattenNoteVulnerabilityDetailsAffectedVersionEnd(c, m["affectedVersionEnd"])
	r.FixedCpeUri = dcl.FlattenString(m["fixedCpeUri"])
	r.FixedPackage = dcl.FlattenString(m["fixedPackage"])
	r.FixedVersion = flattenNoteVulnerabilityDetailsFixedVersion(c, m["fixedVersion"])
	r.IsObsolete = dcl.FlattenBool(m["isObsolete"])
	r.SourceUpdateTime = dcl.FlattenString(m["sourceUpdateTime"])

	return r
}

// expandNoteVulnerabilityDetailsAffectedVersionStartMap expands the contents of NoteVulnerabilityDetailsAffectedVersionStart into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionStartMap(c *Client, f map[string]NoteVulnerabilityDetailsAffectedVersionStart) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityDetailsAffectedVersionStart(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityDetailsAffectedVersionStartSlice expands the contents of NoteVulnerabilityDetailsAffectedVersionStart into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionStartSlice(c *Client, f []NoteVulnerabilityDetailsAffectedVersionStart) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityDetailsAffectedVersionStart(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityDetailsAffectedVersionStartMap flattens the contents of NoteVulnerabilityDetailsAffectedVersionStart from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionStartMap(c *Client, i interface{}) map[string]NoteVulnerabilityDetailsAffectedVersionStart {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetailsAffectedVersionStart{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetailsAffectedVersionStart{}
	}

	items := make(map[string]NoteVulnerabilityDetailsAffectedVersionStart)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetailsAffectedVersionStart(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionStartSlice flattens the contents of NoteVulnerabilityDetailsAffectedVersionStart from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionStartSlice(c *Client, i interface{}) []NoteVulnerabilityDetailsAffectedVersionStart {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsAffectedVersionStart{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsAffectedVersionStart{}
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionStart, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsAffectedVersionStart(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteVulnerabilityDetailsAffectedVersionStart expands an instance of NoteVulnerabilityDetailsAffectedVersionStart into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionStart(c *Client, f *NoteVulnerabilityDetailsAffectedVersionStart) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Epoch; !dcl.IsEmptyValueIndirect(v) {
		m["epoch"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FullName; !dcl.IsEmptyValueIndirect(v) {
		m["fullName"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityDetailsAffectedVersionStart flattens an instance of NoteVulnerabilityDetailsAffectedVersionStart from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionStart(c *Client, i interface{}) *NoteVulnerabilityDetailsAffectedVersionStart {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityDetailsAffectedVersionStart{}
	r.Epoch = dcl.FlattenInteger(m["epoch"])
	r.Name = dcl.FlattenString(m["name"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Kind = flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnum(m["kind"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandNoteVulnerabilityDetailsAffectedVersionEndMap expands the contents of NoteVulnerabilityDetailsAffectedVersionEnd into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionEndMap(c *Client, f map[string]NoteVulnerabilityDetailsAffectedVersionEnd) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityDetailsAffectedVersionEnd(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityDetailsAffectedVersionEndSlice expands the contents of NoteVulnerabilityDetailsAffectedVersionEnd into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionEndSlice(c *Client, f []NoteVulnerabilityDetailsAffectedVersionEnd) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityDetailsAffectedVersionEnd(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityDetailsAffectedVersionEndMap flattens the contents of NoteVulnerabilityDetailsAffectedVersionEnd from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionEndMap(c *Client, i interface{}) map[string]NoteVulnerabilityDetailsAffectedVersionEnd {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetailsAffectedVersionEnd{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetailsAffectedVersionEnd{}
	}

	items := make(map[string]NoteVulnerabilityDetailsAffectedVersionEnd)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetailsAffectedVersionEnd(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionEndSlice flattens the contents of NoteVulnerabilityDetailsAffectedVersionEnd from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionEndSlice(c *Client, i interface{}) []NoteVulnerabilityDetailsAffectedVersionEnd {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsAffectedVersionEnd{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsAffectedVersionEnd{}
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionEnd, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsAffectedVersionEnd(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteVulnerabilityDetailsAffectedVersionEnd expands an instance of NoteVulnerabilityDetailsAffectedVersionEnd into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionEnd(c *Client, f *NoteVulnerabilityDetailsAffectedVersionEnd) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Epoch; !dcl.IsEmptyValueIndirect(v) {
		m["epoch"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FullName; !dcl.IsEmptyValueIndirect(v) {
		m["fullName"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityDetailsAffectedVersionEnd flattens an instance of NoteVulnerabilityDetailsAffectedVersionEnd from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionEnd(c *Client, i interface{}) *NoteVulnerabilityDetailsAffectedVersionEnd {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityDetailsAffectedVersionEnd{}
	r.Epoch = dcl.FlattenInteger(m["epoch"])
	r.Name = dcl.FlattenString(m["name"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Kind = flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnum(m["kind"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandNoteVulnerabilityDetailsFixedVersionMap expands the contents of NoteVulnerabilityDetailsFixedVersion into a JSON
// request object.
func expandNoteVulnerabilityDetailsFixedVersionMap(c *Client, f map[string]NoteVulnerabilityDetailsFixedVersion) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityDetailsFixedVersion(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityDetailsFixedVersionSlice expands the contents of NoteVulnerabilityDetailsFixedVersion into a JSON
// request object.
func expandNoteVulnerabilityDetailsFixedVersionSlice(c *Client, f []NoteVulnerabilityDetailsFixedVersion) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityDetailsFixedVersion(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityDetailsFixedVersionMap flattens the contents of NoteVulnerabilityDetailsFixedVersion from a JSON
// response object.
func flattenNoteVulnerabilityDetailsFixedVersionMap(c *Client, i interface{}) map[string]NoteVulnerabilityDetailsFixedVersion {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetailsFixedVersion{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetailsFixedVersion{}
	}

	items := make(map[string]NoteVulnerabilityDetailsFixedVersion)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetailsFixedVersion(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteVulnerabilityDetailsFixedVersionSlice flattens the contents of NoteVulnerabilityDetailsFixedVersion from a JSON
// response object.
func flattenNoteVulnerabilityDetailsFixedVersionSlice(c *Client, i interface{}) []NoteVulnerabilityDetailsFixedVersion {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsFixedVersion{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsFixedVersion{}
	}

	items := make([]NoteVulnerabilityDetailsFixedVersion, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsFixedVersion(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteVulnerabilityDetailsFixedVersion expands an instance of NoteVulnerabilityDetailsFixedVersion into a JSON
// request object.
func expandNoteVulnerabilityDetailsFixedVersion(c *Client, f *NoteVulnerabilityDetailsFixedVersion) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Epoch; !dcl.IsEmptyValueIndirect(v) {
		m["epoch"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FullName; !dcl.IsEmptyValueIndirect(v) {
		m["fullName"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityDetailsFixedVersion flattens an instance of NoteVulnerabilityDetailsFixedVersion from a JSON
// response object.
func flattenNoteVulnerabilityDetailsFixedVersion(c *Client, i interface{}) *NoteVulnerabilityDetailsFixedVersion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityDetailsFixedVersion{}
	r.Epoch = dcl.FlattenInteger(m["epoch"])
	r.Name = dcl.FlattenString(m["name"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Kind = flattenNoteVulnerabilityDetailsFixedVersionKindEnum(m["kind"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandNoteVulnerabilityCvssV3Map expands the contents of NoteVulnerabilityCvssV3 into a JSON
// request object.
func expandNoteVulnerabilityCvssV3Map(c *Client, f map[string]NoteVulnerabilityCvssV3) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityCvssV3(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityCvssV3Slice expands the contents of NoteVulnerabilityCvssV3 into a JSON
// request object.
func expandNoteVulnerabilityCvssV3Slice(c *Client, f []NoteVulnerabilityCvssV3) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityCvssV3(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityCvssV3Map flattens the contents of NoteVulnerabilityCvssV3 from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3Map(c *Client, i interface{}) map[string]NoteVulnerabilityCvssV3 {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3Slice flattens the contents of NoteVulnerabilityCvssV3 from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3Slice(c *Client, i interface{}) []NoteVulnerabilityCvssV3 {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3{}
	}

	items := make([]NoteVulnerabilityCvssV3, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteVulnerabilityCvssV3 expands an instance of NoteVulnerabilityCvssV3 into a JSON
// request object.
func expandNoteVulnerabilityCvssV3(c *Client, f *NoteVulnerabilityCvssV3) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BaseScore; !dcl.IsEmptyValueIndirect(v) {
		m["baseScore"] = v
	}
	if v := f.ExploitabilityScore; !dcl.IsEmptyValueIndirect(v) {
		m["exploitabilityScore"] = v
	}
	if v := f.ImpactScore; !dcl.IsEmptyValueIndirect(v) {
		m["impactScore"] = v
	}
	if v := f.AttackVector; !dcl.IsEmptyValueIndirect(v) {
		m["attackVector"] = v
	}
	if v := f.AttackComplexity; !dcl.IsEmptyValueIndirect(v) {
		m["attackComplexity"] = v
	}
	if v := f.PrivilegesRequired; !dcl.IsEmptyValueIndirect(v) {
		m["privilegesRequired"] = v
	}
	if v := f.UserInteraction; !dcl.IsEmptyValueIndirect(v) {
		m["userInteraction"] = v
	}
	if v := f.Scope; !dcl.IsEmptyValueIndirect(v) {
		m["scope"] = v
	}
	if v := f.ConfidentialityImpact; !dcl.IsEmptyValueIndirect(v) {
		m["confidentialityImpact"] = v
	}
	if v := f.IntegrityImpact; !dcl.IsEmptyValueIndirect(v) {
		m["integrityImpact"] = v
	}
	if v := f.AvailabilityImpact; !dcl.IsEmptyValueIndirect(v) {
		m["availabilityImpact"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityCvssV3 flattens an instance of NoteVulnerabilityCvssV3 from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3(c *Client, i interface{}) *NoteVulnerabilityCvssV3 {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityCvssV3{}
	r.BaseScore = dcl.FlattenDouble(m["baseScore"])
	r.ExploitabilityScore = dcl.FlattenDouble(m["exploitabilityScore"])
	r.ImpactScore = dcl.FlattenDouble(m["impactScore"])
	r.AttackVector = flattenNoteVulnerabilityCvssV3AttackVectorEnum(m["attackVector"])
	r.AttackComplexity = flattenNoteVulnerabilityCvssV3AttackComplexityEnum(m["attackComplexity"])
	r.PrivilegesRequired = flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnum(m["privilegesRequired"])
	r.UserInteraction = flattenNoteVulnerabilityCvssV3UserInteractionEnum(m["userInteraction"])
	r.Scope = flattenNoteVulnerabilityCvssV3ScopeEnum(m["scope"])
	r.ConfidentialityImpact = flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnum(m["confidentialityImpact"])
	r.IntegrityImpact = flattenNoteVulnerabilityCvssV3IntegrityImpactEnum(m["integrityImpact"])
	r.AvailabilityImpact = flattenNoteVulnerabilityCvssV3AvailabilityImpactEnum(m["availabilityImpact"])

	return r
}

// expandNoteVulnerabilityWindowsDetailsMap expands the contents of NoteVulnerabilityWindowsDetails into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsMap(c *Client, f map[string]NoteVulnerabilityWindowsDetails) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityWindowsDetails(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityWindowsDetailsSlice expands the contents of NoteVulnerabilityWindowsDetails into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsSlice(c *Client, f []NoteVulnerabilityWindowsDetails) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityWindowsDetails(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityWindowsDetailsMap flattens the contents of NoteVulnerabilityWindowsDetails from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsMap(c *Client, i interface{}) map[string]NoteVulnerabilityWindowsDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityWindowsDetails{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityWindowsDetails{}
	}

	items := make(map[string]NoteVulnerabilityWindowsDetails)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityWindowsDetails(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteVulnerabilityWindowsDetailsSlice flattens the contents of NoteVulnerabilityWindowsDetails from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsSlice(c *Client, i interface{}) []NoteVulnerabilityWindowsDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityWindowsDetails{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityWindowsDetails{}
	}

	items := make([]NoteVulnerabilityWindowsDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityWindowsDetails(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteVulnerabilityWindowsDetails expands an instance of NoteVulnerabilityWindowsDetails into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetails(c *Client, f *NoteVulnerabilityWindowsDetails) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CpeUri; !dcl.IsEmptyValueIndirect(v) {
		m["cpeUri"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandNoteVulnerabilityWindowsDetailsFixingKbsSlice(c, f.FixingKbs); err != nil {
		return nil, fmt.Errorf("error expanding FixingKbs into fixingKbs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fixingKbs"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityWindowsDetails flattens an instance of NoteVulnerabilityWindowsDetails from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetails(c *Client, i interface{}) *NoteVulnerabilityWindowsDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityWindowsDetails{}
	r.CpeUri = dcl.FlattenString(m["cpeUri"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.FixingKbs = flattenNoteVulnerabilityWindowsDetailsFixingKbsSlice(c, m["fixingKbs"])

	return r
}

// expandNoteVulnerabilityWindowsDetailsFixingKbsMap expands the contents of NoteVulnerabilityWindowsDetailsFixingKbs into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsFixingKbsMap(c *Client, f map[string]NoteVulnerabilityWindowsDetailsFixingKbs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityWindowsDetailsFixingKbs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityWindowsDetailsFixingKbsSlice expands the contents of NoteVulnerabilityWindowsDetailsFixingKbs into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsFixingKbsSlice(c *Client, f []NoteVulnerabilityWindowsDetailsFixingKbs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityWindowsDetailsFixingKbs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityWindowsDetailsFixingKbsMap flattens the contents of NoteVulnerabilityWindowsDetailsFixingKbs from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsFixingKbsMap(c *Client, i interface{}) map[string]NoteVulnerabilityWindowsDetailsFixingKbs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityWindowsDetailsFixingKbs{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityWindowsDetailsFixingKbs{}
	}

	items := make(map[string]NoteVulnerabilityWindowsDetailsFixingKbs)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityWindowsDetailsFixingKbs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteVulnerabilityWindowsDetailsFixingKbsSlice flattens the contents of NoteVulnerabilityWindowsDetailsFixingKbs from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsFixingKbsSlice(c *Client, i interface{}) []NoteVulnerabilityWindowsDetailsFixingKbs {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityWindowsDetailsFixingKbs{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityWindowsDetailsFixingKbs{}
	}

	items := make([]NoteVulnerabilityWindowsDetailsFixingKbs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityWindowsDetailsFixingKbs(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteVulnerabilityWindowsDetailsFixingKbs expands an instance of NoteVulnerabilityWindowsDetailsFixingKbs into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsFixingKbs(c *Client, f *NoteVulnerabilityWindowsDetailsFixingKbs) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityWindowsDetailsFixingKbs flattens an instance of NoteVulnerabilityWindowsDetailsFixingKbs from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsFixingKbs(c *Client, i interface{}) *NoteVulnerabilityWindowsDetailsFixingKbs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityWindowsDetailsFixingKbs{}
	r.Name = dcl.FlattenString(m["name"])
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandNoteBuildMap expands the contents of NoteBuild into a JSON
// request object.
func expandNoteBuildMap(c *Client, f map[string]NoteBuild) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteBuild(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteBuildSlice expands the contents of NoteBuild into a JSON
// request object.
func expandNoteBuildSlice(c *Client, f []NoteBuild) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteBuild(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteBuildMap flattens the contents of NoteBuild from a JSON
// response object.
func flattenNoteBuildMap(c *Client, i interface{}) map[string]NoteBuild {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteBuild{}
	}

	if len(a) == 0 {
		return map[string]NoteBuild{}
	}

	items := make(map[string]NoteBuild)
	for k, item := range a {
		items[k] = *flattenNoteBuild(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteBuildSlice flattens the contents of NoteBuild from a JSON
// response object.
func flattenNoteBuildSlice(c *Client, i interface{}) []NoteBuild {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteBuild{}
	}

	if len(a) == 0 {
		return []NoteBuild{}
	}

	items := make([]NoteBuild, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteBuild(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteBuild expands an instance of NoteBuild into a JSON
// request object.
func expandNoteBuild(c *Client, f *NoteBuild) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BuilderVersion; !dcl.IsEmptyValueIndirect(v) {
		m["builderVersion"] = v
	}

	return m, nil
}

// flattenNoteBuild flattens an instance of NoteBuild from a JSON
// response object.
func flattenNoteBuild(c *Client, i interface{}) *NoteBuild {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteBuild{}
	r.BuilderVersion = dcl.FlattenString(m["builderVersion"])

	return r
}

// expandNoteImageMap expands the contents of NoteImage into a JSON
// request object.
func expandNoteImageMap(c *Client, f map[string]NoteImage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteImage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteImageSlice expands the contents of NoteImage into a JSON
// request object.
func expandNoteImageSlice(c *Client, f []NoteImage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteImage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteImageMap flattens the contents of NoteImage from a JSON
// response object.
func flattenNoteImageMap(c *Client, i interface{}) map[string]NoteImage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteImage{}
	}

	if len(a) == 0 {
		return map[string]NoteImage{}
	}

	items := make(map[string]NoteImage)
	for k, item := range a {
		items[k] = *flattenNoteImage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteImageSlice flattens the contents of NoteImage from a JSON
// response object.
func flattenNoteImageSlice(c *Client, i interface{}) []NoteImage {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteImage{}
	}

	if len(a) == 0 {
		return []NoteImage{}
	}

	items := make([]NoteImage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteImage(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteImage expands an instance of NoteImage into a JSON
// request object.
func expandNoteImage(c *Client, f *NoteImage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceUrl; !dcl.IsEmptyValueIndirect(v) {
		m["resourceUrl"] = v
	}
	if v, err := expandNoteImageFingerprint(c, f.Fingerprint); err != nil {
		return nil, fmt.Errorf("error expanding Fingerprint into fingerprint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fingerprint"] = v
	}

	return m, nil
}

// flattenNoteImage flattens an instance of NoteImage from a JSON
// response object.
func flattenNoteImage(c *Client, i interface{}) *NoteImage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteImage{}
	r.ResourceUrl = dcl.FlattenString(m["resourceUrl"])
	r.Fingerprint = flattenNoteImageFingerprint(c, m["fingerprint"])

	return r
}

// expandNoteImageFingerprintMap expands the contents of NoteImageFingerprint into a JSON
// request object.
func expandNoteImageFingerprintMap(c *Client, f map[string]NoteImageFingerprint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteImageFingerprint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteImageFingerprintSlice expands the contents of NoteImageFingerprint into a JSON
// request object.
func expandNoteImageFingerprintSlice(c *Client, f []NoteImageFingerprint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteImageFingerprint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteImageFingerprintMap flattens the contents of NoteImageFingerprint from a JSON
// response object.
func flattenNoteImageFingerprintMap(c *Client, i interface{}) map[string]NoteImageFingerprint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteImageFingerprint{}
	}

	if len(a) == 0 {
		return map[string]NoteImageFingerprint{}
	}

	items := make(map[string]NoteImageFingerprint)
	for k, item := range a {
		items[k] = *flattenNoteImageFingerprint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteImageFingerprintSlice flattens the contents of NoteImageFingerprint from a JSON
// response object.
func flattenNoteImageFingerprintSlice(c *Client, i interface{}) []NoteImageFingerprint {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteImageFingerprint{}
	}

	if len(a) == 0 {
		return []NoteImageFingerprint{}
	}

	items := make([]NoteImageFingerprint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteImageFingerprint(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteImageFingerprint expands an instance of NoteImageFingerprint into a JSON
// request object.
func expandNoteImageFingerprint(c *Client, f *NoteImageFingerprint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.V1Name; !dcl.IsEmptyValueIndirect(v) {
		m["v1Name"] = v
	}
	if v := f.V2Blob; !dcl.IsEmptyValueIndirect(v) {
		m["v2Blob"] = v
	}
	if v := f.V2Name; !dcl.IsEmptyValueIndirect(v) {
		m["v2Name"] = v
	}

	return m, nil
}

// flattenNoteImageFingerprint flattens an instance of NoteImageFingerprint from a JSON
// response object.
func flattenNoteImageFingerprint(c *Client, i interface{}) *NoteImageFingerprint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteImageFingerprint{}
	r.V1Name = dcl.FlattenString(m["v1Name"])
	r.V2Blob = dcl.FlattenStringSlice(m["v2Blob"])
	r.V2Name = dcl.FlattenString(m["v2Name"])

	return r
}

// expandNotePackageMap expands the contents of NotePackage into a JSON
// request object.
func expandNotePackageMap(c *Client, f map[string]NotePackage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNotePackage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNotePackageSlice expands the contents of NotePackage into a JSON
// request object.
func expandNotePackageSlice(c *Client, f []NotePackage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNotePackage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNotePackageMap flattens the contents of NotePackage from a JSON
// response object.
func flattenNotePackageMap(c *Client, i interface{}) map[string]NotePackage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackage{}
	}

	if len(a) == 0 {
		return map[string]NotePackage{}
	}

	items := make(map[string]NotePackage)
	for k, item := range a {
		items[k] = *flattenNotePackage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNotePackageSlice flattens the contents of NotePackage from a JSON
// response object.
func flattenNotePackageSlice(c *Client, i interface{}) []NotePackage {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackage{}
	}

	if len(a) == 0 {
		return []NotePackage{}
	}

	items := make([]NotePackage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackage(c, item.(map[string]interface{})))
	}

	return items
}

// expandNotePackage expands an instance of NotePackage into a JSON
// request object.
func expandNotePackage(c *Client, f *NotePackage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandNotePackageDistributionSlice(c, f.Distribution); err != nil {
		return nil, fmt.Errorf("error expanding Distribution into distribution: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["distribution"] = v
	}

	return m, nil
}

// flattenNotePackage flattens an instance of NotePackage from a JSON
// response object.
func flattenNotePackage(c *Client, i interface{}) *NotePackage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NotePackage{}
	r.Name = dcl.FlattenString(m["name"])
	r.Distribution = flattenNotePackageDistributionSlice(c, m["distribution"])

	return r
}

// expandNotePackageDistributionMap expands the contents of NotePackageDistribution into a JSON
// request object.
func expandNotePackageDistributionMap(c *Client, f map[string]NotePackageDistribution) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNotePackageDistribution(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNotePackageDistributionSlice expands the contents of NotePackageDistribution into a JSON
// request object.
func expandNotePackageDistributionSlice(c *Client, f []NotePackageDistribution) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNotePackageDistribution(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNotePackageDistributionMap flattens the contents of NotePackageDistribution from a JSON
// response object.
func flattenNotePackageDistributionMap(c *Client, i interface{}) map[string]NotePackageDistribution {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistribution{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistribution{}
	}

	items := make(map[string]NotePackageDistribution)
	for k, item := range a {
		items[k] = *flattenNotePackageDistribution(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNotePackageDistributionSlice flattens the contents of NotePackageDistribution from a JSON
// response object.
func flattenNotePackageDistributionSlice(c *Client, i interface{}) []NotePackageDistribution {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistribution{}
	}

	if len(a) == 0 {
		return []NotePackageDistribution{}
	}

	items := make([]NotePackageDistribution, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistribution(c, item.(map[string]interface{})))
	}

	return items
}

// expandNotePackageDistribution expands an instance of NotePackageDistribution into a JSON
// request object.
func expandNotePackageDistribution(c *Client, f *NotePackageDistribution) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CpeUri; !dcl.IsEmptyValueIndirect(v) {
		m["cpeUri"] = v
	}
	if v := f.Architecture; !dcl.IsEmptyValueIndirect(v) {
		m["architecture"] = v
	}
	if v, err := expandNotePackageDistributionLatestVersion(c, f.LatestVersion); err != nil {
		return nil, fmt.Errorf("error expanding LatestVersion into latestVersion: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["latestVersion"] = v
	}
	if v := f.Maintainer; !dcl.IsEmptyValueIndirect(v) {
		m["maintainer"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenNotePackageDistribution flattens an instance of NotePackageDistribution from a JSON
// response object.
func flattenNotePackageDistribution(c *Client, i interface{}) *NotePackageDistribution {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NotePackageDistribution{}
	r.CpeUri = dcl.FlattenString(m["cpeUri"])
	r.Architecture = flattenNotePackageDistributionArchitectureEnum(m["architecture"])
	r.LatestVersion = flattenNotePackageDistributionLatestVersion(c, m["latestVersion"])
	r.Maintainer = dcl.FlattenString(m["maintainer"])
	r.Url = dcl.FlattenString(m["url"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandNotePackageDistributionLatestVersionMap expands the contents of NotePackageDistributionLatestVersion into a JSON
// request object.
func expandNotePackageDistributionLatestVersionMap(c *Client, f map[string]NotePackageDistributionLatestVersion) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNotePackageDistributionLatestVersion(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNotePackageDistributionLatestVersionSlice expands the contents of NotePackageDistributionLatestVersion into a JSON
// request object.
func expandNotePackageDistributionLatestVersionSlice(c *Client, f []NotePackageDistributionLatestVersion) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNotePackageDistributionLatestVersion(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNotePackageDistributionLatestVersionMap flattens the contents of NotePackageDistributionLatestVersion from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionMap(c *Client, i interface{}) map[string]NotePackageDistributionLatestVersion {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistributionLatestVersion{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistributionLatestVersion{}
	}

	items := make(map[string]NotePackageDistributionLatestVersion)
	for k, item := range a {
		items[k] = *flattenNotePackageDistributionLatestVersion(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNotePackageDistributionLatestVersionSlice flattens the contents of NotePackageDistributionLatestVersion from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionSlice(c *Client, i interface{}) []NotePackageDistributionLatestVersion {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistributionLatestVersion{}
	}

	if len(a) == 0 {
		return []NotePackageDistributionLatestVersion{}
	}

	items := make([]NotePackageDistributionLatestVersion, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistributionLatestVersion(c, item.(map[string]interface{})))
	}

	return items
}

// expandNotePackageDistributionLatestVersion expands an instance of NotePackageDistributionLatestVersion into a JSON
// request object.
func expandNotePackageDistributionLatestVersion(c *Client, f *NotePackageDistributionLatestVersion) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Epoch; !dcl.IsEmptyValueIndirect(v) {
		m["epoch"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FullName; !dcl.IsEmptyValueIndirect(v) {
		m["fullName"] = v
	}

	return m, nil
}

// flattenNotePackageDistributionLatestVersion flattens an instance of NotePackageDistributionLatestVersion from a JSON
// response object.
func flattenNotePackageDistributionLatestVersion(c *Client, i interface{}) *NotePackageDistributionLatestVersion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NotePackageDistributionLatestVersion{}
	r.Epoch = dcl.FlattenInteger(m["epoch"])
	r.Name = dcl.FlattenString(m["name"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Kind = flattenNotePackageDistributionLatestVersionKindEnum(m["kind"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandNoteDiscoveryMap expands the contents of NoteDiscovery into a JSON
// request object.
func expandNoteDiscoveryMap(c *Client, f map[string]NoteDiscovery) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteDiscovery(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteDiscoverySlice expands the contents of NoteDiscovery into a JSON
// request object.
func expandNoteDiscoverySlice(c *Client, f []NoteDiscovery) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteDiscovery(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteDiscoveryMap flattens the contents of NoteDiscovery from a JSON
// response object.
func flattenNoteDiscoveryMap(c *Client, i interface{}) map[string]NoteDiscovery {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteDiscovery{}
	}

	if len(a) == 0 {
		return map[string]NoteDiscovery{}
	}

	items := make(map[string]NoteDiscovery)
	for k, item := range a {
		items[k] = *flattenNoteDiscovery(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteDiscoverySlice flattens the contents of NoteDiscovery from a JSON
// response object.
func flattenNoteDiscoverySlice(c *Client, i interface{}) []NoteDiscovery {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteDiscovery{}
	}

	if len(a) == 0 {
		return []NoteDiscovery{}
	}

	items := make([]NoteDiscovery, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteDiscovery(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteDiscovery expands an instance of NoteDiscovery into a JSON
// request object.
func expandNoteDiscovery(c *Client, f *NoteDiscovery) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AnalysisKind; !dcl.IsEmptyValueIndirect(v) {
		m["analysisKind"] = v
	}

	return m, nil
}

// flattenNoteDiscovery flattens an instance of NoteDiscovery from a JSON
// response object.
func flattenNoteDiscovery(c *Client, i interface{}) *NoteDiscovery {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteDiscovery{}
	r.AnalysisKind = flattenNoteDiscoveryAnalysisKindEnum(m["analysisKind"])

	return r
}

// expandNoteDeploymentMap expands the contents of NoteDeployment into a JSON
// request object.
func expandNoteDeploymentMap(c *Client, f map[string]NoteDeployment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteDeployment(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteDeploymentSlice expands the contents of NoteDeployment into a JSON
// request object.
func expandNoteDeploymentSlice(c *Client, f []NoteDeployment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteDeployment(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteDeploymentMap flattens the contents of NoteDeployment from a JSON
// response object.
func flattenNoteDeploymentMap(c *Client, i interface{}) map[string]NoteDeployment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteDeployment{}
	}

	if len(a) == 0 {
		return map[string]NoteDeployment{}
	}

	items := make(map[string]NoteDeployment)
	for k, item := range a {
		items[k] = *flattenNoteDeployment(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteDeploymentSlice flattens the contents of NoteDeployment from a JSON
// response object.
func flattenNoteDeploymentSlice(c *Client, i interface{}) []NoteDeployment {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteDeployment{}
	}

	if len(a) == 0 {
		return []NoteDeployment{}
	}

	items := make([]NoteDeployment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteDeployment(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteDeployment expands an instance of NoteDeployment into a JSON
// request object.
func expandNoteDeployment(c *Client, f *NoteDeployment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceUri; !dcl.IsEmptyValueIndirect(v) {
		m["resourceUri"] = v
	}

	return m, nil
}

// flattenNoteDeployment flattens an instance of NoteDeployment from a JSON
// response object.
func flattenNoteDeployment(c *Client, i interface{}) *NoteDeployment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteDeployment{}
	r.ResourceUri = dcl.FlattenStringSlice(m["resourceUri"])

	return r
}

// expandNoteAttestationMap expands the contents of NoteAttestation into a JSON
// request object.
func expandNoteAttestationMap(c *Client, f map[string]NoteAttestation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteAttestation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteAttestationSlice expands the contents of NoteAttestation into a JSON
// request object.
func expandNoteAttestationSlice(c *Client, f []NoteAttestation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteAttestation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteAttestationMap flattens the contents of NoteAttestation from a JSON
// response object.
func flattenNoteAttestationMap(c *Client, i interface{}) map[string]NoteAttestation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteAttestation{}
	}

	if len(a) == 0 {
		return map[string]NoteAttestation{}
	}

	items := make(map[string]NoteAttestation)
	for k, item := range a {
		items[k] = *flattenNoteAttestation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteAttestationSlice flattens the contents of NoteAttestation from a JSON
// response object.
func flattenNoteAttestationSlice(c *Client, i interface{}) []NoteAttestation {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteAttestation{}
	}

	if len(a) == 0 {
		return []NoteAttestation{}
	}

	items := make([]NoteAttestation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteAttestation(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteAttestation expands an instance of NoteAttestation into a JSON
// request object.
func expandNoteAttestation(c *Client, f *NoteAttestation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandNoteAttestationHint(c, f.Hint); err != nil {
		return nil, fmt.Errorf("error expanding Hint into hint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hint"] = v
	}

	return m, nil
}

// flattenNoteAttestation flattens an instance of NoteAttestation from a JSON
// response object.
func flattenNoteAttestation(c *Client, i interface{}) *NoteAttestation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteAttestation{}
	r.Hint = flattenNoteAttestationHint(c, m["hint"])

	return r
}

// expandNoteAttestationHintMap expands the contents of NoteAttestationHint into a JSON
// request object.
func expandNoteAttestationHintMap(c *Client, f map[string]NoteAttestationHint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteAttestationHint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteAttestationHintSlice expands the contents of NoteAttestationHint into a JSON
// request object.
func expandNoteAttestationHintSlice(c *Client, f []NoteAttestationHint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteAttestationHint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteAttestationHintMap flattens the contents of NoteAttestationHint from a JSON
// response object.
func flattenNoteAttestationHintMap(c *Client, i interface{}) map[string]NoteAttestationHint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteAttestationHint{}
	}

	if len(a) == 0 {
		return map[string]NoteAttestationHint{}
	}

	items := make(map[string]NoteAttestationHint)
	for k, item := range a {
		items[k] = *flattenNoteAttestationHint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNoteAttestationHintSlice flattens the contents of NoteAttestationHint from a JSON
// response object.
func flattenNoteAttestationHintSlice(c *Client, i interface{}) []NoteAttestationHint {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteAttestationHint{}
	}

	if len(a) == 0 {
		return []NoteAttestationHint{}
	}

	items := make([]NoteAttestationHint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteAttestationHint(c, item.(map[string]interface{})))
	}

	return items
}

// expandNoteAttestationHint expands an instance of NoteAttestationHint into a JSON
// request object.
func expandNoteAttestationHint(c *Client, f *NoteAttestationHint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HumanReadableName; !dcl.IsEmptyValueIndirect(v) {
		m["humanReadableName"] = v
	}

	return m, nil
}

// flattenNoteAttestationHint flattens an instance of NoteAttestationHint from a JSON
// response object.
func flattenNoteAttestationHint(c *Client, i interface{}) *NoteAttestationHint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteAttestationHint{}
	r.HumanReadableName = dcl.FlattenString(m["humanReadableName"])

	return r
}

// flattenNoteVulnerabilitySeverityEnumSlice flattens the contents of NoteVulnerabilitySeverityEnum from a JSON
// response object.
func flattenNoteVulnerabilitySeverityEnumSlice(c *Client, i interface{}) []NoteVulnerabilitySeverityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilitySeverityEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilitySeverityEnum{}
	}

	items := make([]NoteVulnerabilitySeverityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilitySeverityEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilitySeverityEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilitySeverityEnum with the same value as that string.
func flattenNoteVulnerabilitySeverityEnum(i interface{}) *NoteVulnerabilitySeverityEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilitySeverityEnumRef("")
	}

	return NoteVulnerabilitySeverityEnumRef(s)
}

// flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnumSlice flattens the contents of NoteVulnerabilityDetailsAffectedVersionStartKindEnum from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnumSlice(c *Client, i interface{}) []NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsAffectedVersionStartKindEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsAffectedVersionStartKindEnum{}
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionStartKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityDetailsAffectedVersionStartKindEnum with the same value as that string.
func flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnum(i interface{}) *NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityDetailsAffectedVersionStartKindEnumRef("")
	}

	return NoteVulnerabilityDetailsAffectedVersionStartKindEnumRef(s)
}

// flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnumSlice flattens the contents of NoteVulnerabilityDetailsAffectedVersionEndKindEnum from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnumSlice(c *Client, i interface{}) []NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsAffectedVersionEndKindEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsAffectedVersionEndKindEnum{}
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionEndKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityDetailsAffectedVersionEndKindEnum with the same value as that string.
func flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnum(i interface{}) *NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityDetailsAffectedVersionEndKindEnumRef("")
	}

	return NoteVulnerabilityDetailsAffectedVersionEndKindEnumRef(s)
}

// flattenNoteVulnerabilityDetailsFixedVersionKindEnumSlice flattens the contents of NoteVulnerabilityDetailsFixedVersionKindEnum from a JSON
// response object.
func flattenNoteVulnerabilityDetailsFixedVersionKindEnumSlice(c *Client, i interface{}) []NoteVulnerabilityDetailsFixedVersionKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsFixedVersionKindEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsFixedVersionKindEnum{}
	}

	items := make([]NoteVulnerabilityDetailsFixedVersionKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsFixedVersionKindEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityDetailsFixedVersionKindEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityDetailsFixedVersionKindEnum with the same value as that string.
func flattenNoteVulnerabilityDetailsFixedVersionKindEnum(i interface{}) *NoteVulnerabilityDetailsFixedVersionKindEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityDetailsFixedVersionKindEnumRef("")
	}

	return NoteVulnerabilityDetailsFixedVersionKindEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3AttackVectorEnumSlice flattens the contents of NoteVulnerabilityCvssV3AttackVectorEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3AttackVectorEnumSlice(c *Client, i interface{}) []NoteVulnerabilityCvssV3AttackVectorEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3AttackVectorEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3AttackVectorEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3AttackVectorEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3AttackVectorEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3AttackVectorEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3AttackVectorEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3AttackVectorEnum(i interface{}) *NoteVulnerabilityCvssV3AttackVectorEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityCvssV3AttackVectorEnumRef("")
	}

	return NoteVulnerabilityCvssV3AttackVectorEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3AttackComplexityEnumSlice flattens the contents of NoteVulnerabilityCvssV3AttackComplexityEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3AttackComplexityEnumSlice(c *Client, i interface{}) []NoteVulnerabilityCvssV3AttackComplexityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3AttackComplexityEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3AttackComplexityEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3AttackComplexityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3AttackComplexityEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3AttackComplexityEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3AttackComplexityEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3AttackComplexityEnum(i interface{}) *NoteVulnerabilityCvssV3AttackComplexityEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityCvssV3AttackComplexityEnumRef("")
	}

	return NoteVulnerabilityCvssV3AttackComplexityEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnumSlice flattens the contents of NoteVulnerabilityCvssV3PrivilegesRequiredEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnumSlice(c *Client, i interface{}) []NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3PrivilegesRequiredEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3PrivilegesRequiredEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3PrivilegesRequiredEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3PrivilegesRequiredEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnum(i interface{}) *NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityCvssV3PrivilegesRequiredEnumRef("")
	}

	return NoteVulnerabilityCvssV3PrivilegesRequiredEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3UserInteractionEnumSlice flattens the contents of NoteVulnerabilityCvssV3UserInteractionEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3UserInteractionEnumSlice(c *Client, i interface{}) []NoteVulnerabilityCvssV3UserInteractionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3UserInteractionEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3UserInteractionEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3UserInteractionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3UserInteractionEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3UserInteractionEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3UserInteractionEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3UserInteractionEnum(i interface{}) *NoteVulnerabilityCvssV3UserInteractionEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityCvssV3UserInteractionEnumRef("")
	}

	return NoteVulnerabilityCvssV3UserInteractionEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3ScopeEnumSlice flattens the contents of NoteVulnerabilityCvssV3ScopeEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3ScopeEnumSlice(c *Client, i interface{}) []NoteVulnerabilityCvssV3ScopeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3ScopeEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3ScopeEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3ScopeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3ScopeEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3ScopeEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3ScopeEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3ScopeEnum(i interface{}) *NoteVulnerabilityCvssV3ScopeEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityCvssV3ScopeEnumRef("")
	}

	return NoteVulnerabilityCvssV3ScopeEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnumSlice flattens the contents of NoteVulnerabilityCvssV3ConfidentialityImpactEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnumSlice(c *Client, i interface{}) []NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3ConfidentialityImpactEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3ConfidentialityImpactEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3ConfidentialityImpactEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3ConfidentialityImpactEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnum(i interface{}) *NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityCvssV3ConfidentialityImpactEnumRef("")
	}

	return NoteVulnerabilityCvssV3ConfidentialityImpactEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3IntegrityImpactEnumSlice flattens the contents of NoteVulnerabilityCvssV3IntegrityImpactEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3IntegrityImpactEnumSlice(c *Client, i interface{}) []NoteVulnerabilityCvssV3IntegrityImpactEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3IntegrityImpactEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3IntegrityImpactEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3IntegrityImpactEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3IntegrityImpactEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3IntegrityImpactEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3IntegrityImpactEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3IntegrityImpactEnum(i interface{}) *NoteVulnerabilityCvssV3IntegrityImpactEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityCvssV3IntegrityImpactEnumRef("")
	}

	return NoteVulnerabilityCvssV3IntegrityImpactEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3AvailabilityImpactEnumSlice flattens the contents of NoteVulnerabilityCvssV3AvailabilityImpactEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3AvailabilityImpactEnumSlice(c *Client, i interface{}) []NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3AvailabilityImpactEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3AvailabilityImpactEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3AvailabilityImpactEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3AvailabilityImpactEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3AvailabilityImpactEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3AvailabilityImpactEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3AvailabilityImpactEnum(i interface{}) *NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	s, ok := i.(string)
	if !ok {
		return NoteVulnerabilityCvssV3AvailabilityImpactEnumRef("")
	}

	return NoteVulnerabilityCvssV3AvailabilityImpactEnumRef(s)
}

// flattenNotePackageDistributionArchitectureEnumSlice flattens the contents of NotePackageDistributionArchitectureEnum from a JSON
// response object.
func flattenNotePackageDistributionArchitectureEnumSlice(c *Client, i interface{}) []NotePackageDistributionArchitectureEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistributionArchitectureEnum{}
	}

	if len(a) == 0 {
		return []NotePackageDistributionArchitectureEnum{}
	}

	items := make([]NotePackageDistributionArchitectureEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistributionArchitectureEnum(item.(interface{})))
	}

	return items
}

// flattenNotePackageDistributionArchitectureEnum asserts that an interface is a string, and returns a
// pointer to a *NotePackageDistributionArchitectureEnum with the same value as that string.
func flattenNotePackageDistributionArchitectureEnum(i interface{}) *NotePackageDistributionArchitectureEnum {
	s, ok := i.(string)
	if !ok {
		return NotePackageDistributionArchitectureEnumRef("")
	}

	return NotePackageDistributionArchitectureEnumRef(s)
}

// flattenNotePackageDistributionLatestVersionKindEnumSlice flattens the contents of NotePackageDistributionLatestVersionKindEnum from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionKindEnumSlice(c *Client, i interface{}) []NotePackageDistributionLatestVersionKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistributionLatestVersionKindEnum{}
	}

	if len(a) == 0 {
		return []NotePackageDistributionLatestVersionKindEnum{}
	}

	items := make([]NotePackageDistributionLatestVersionKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistributionLatestVersionKindEnum(item.(interface{})))
	}

	return items
}

// flattenNotePackageDistributionLatestVersionKindEnum asserts that an interface is a string, and returns a
// pointer to a *NotePackageDistributionLatestVersionKindEnum with the same value as that string.
func flattenNotePackageDistributionLatestVersionKindEnum(i interface{}) *NotePackageDistributionLatestVersionKindEnum {
	s, ok := i.(string)
	if !ok {
		return NotePackageDistributionLatestVersionKindEnumRef("")
	}

	return NotePackageDistributionLatestVersionKindEnumRef(s)
}

// flattenNoteDiscoveryAnalysisKindEnumSlice flattens the contents of NoteDiscoveryAnalysisKindEnum from a JSON
// response object.
func flattenNoteDiscoveryAnalysisKindEnumSlice(c *Client, i interface{}) []NoteDiscoveryAnalysisKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteDiscoveryAnalysisKindEnum{}
	}

	if len(a) == 0 {
		return []NoteDiscoveryAnalysisKindEnum{}
	}

	items := make([]NoteDiscoveryAnalysisKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteDiscoveryAnalysisKindEnum(item.(interface{})))
	}

	return items
}

// flattenNoteDiscoveryAnalysisKindEnum asserts that an interface is a string, and returns a
// pointer to a *NoteDiscoveryAnalysisKindEnum with the same value as that string.
func flattenNoteDiscoveryAnalysisKindEnum(i interface{}) *NoteDiscoveryAnalysisKindEnum {
	s, ok := i.(string)
	if !ok {
		return NoteDiscoveryAnalysisKindEnumRef("")
	}

	return NoteDiscoveryAnalysisKindEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Note) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalNote(b, c)
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
