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
package sourcerepo

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

func (r *Repo) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}
func (r *RepoPubsubConfigs) validate() error {
	if err := dcl.Required(r, "topic"); err != nil {
		return err
	}
	if err := dcl.Required(r, "messageFormat"); err != nil {
		return err
	}
	return nil
}

func repoGetURL(userBasePath string, r *Repo) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/repos/{{name}}", "https://sourcerepo.googleapis.com/v1/", userBasePath, params), nil
}

func repoListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/repos", "https://sourcerepo.googleapis.com/v1/", userBasePath, params), nil

}

func repoCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/repos", "https://sourcerepo.googleapis.com/v1/", userBasePath, params), nil

}

func repoDeleteURL(userBasePath string, r *Repo) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/repos/{{name}}", "https://sourcerepo.googleapis.com/v1/", userBasePath, params), nil
}

// repoApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type repoApiOperation interface {
	do(context.Context, *Repo, *Client) error
}

// newUpdateRepoUpdateRepoRequest creates a request for an
// Repo resource's UpdateRepo update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateRepoUpdateRepoRequest(ctx context.Context, f *Repo, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandRepoPubsubConfig(f, f.PubsubConfigs); err != nil {
		return nil, fmt.Errorf("error expanding PubsubConfigs into pubsubConfigs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["pubsubConfigs"] = v
	}
	return req, nil
}

// marshalUpdateRepoUpdateRepoRequest converts the update into
// the final JSON request body.
func marshalUpdateRepoUpdateRepoRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeRepoUpdateRequest(m)
	return json.Marshal(m)
}

type updateRepoUpdateRepoOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateRepoUpdateRepoOperation) do(ctx context.Context, r *Repo, c *Client) error {
	_, err := c.GetRepo(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateRepo")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"pubsubConfigs"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateRepoUpdateRepoRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateRepoUpdateRepoRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listRepoRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := repoListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != RepoMaxPage {
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

type listRepoOperation struct {
	Repos []map[string]interface{} `json:"repos"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listRepo(ctx context.Context, project, pageToken string, pageSize int32) ([]*Repo, string, error) {
	b, err := c.listRepoRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listRepoOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Repo
	for _, v := range m.Repos {
		res := flattenRepo(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllRepo(ctx context.Context, f func(*Repo) bool, resources []*Repo) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteRepo(ctx, res)
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

type deleteRepoOperation struct{}

func (op *deleteRepoOperation) do(ctx context.Context, r *Repo, c *Client) error {

	_, err := c.GetRepo(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Repo not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetRepo checking for existence. error: %v", err)
		return err
	}

	u, err := repoDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Repo: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetRepo(ctx, r.urlNormalized())
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
type createRepoOperation struct {
	response map[string]interface{}
}

func (op *createRepoOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createRepoOperation) do(ctx context.Context, r *Repo, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := repoCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetRepo(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getRepoRaw(ctx context.Context, r *Repo) ([]byte, error) {

	u, err := repoGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) repoDiffsForRawDesired(ctx context.Context, rawDesired *Repo, opts ...dcl.ApplyOption) (initial, desired *Repo, diffs []repoDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Repo
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Repo); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Repo, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetRepo(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Repo resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Repo resource: %v", err)
		}
		c.Config.Logger.Info("Found that Repo resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeRepoDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Repo: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Repo: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeRepoInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Repo: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeRepoDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Repo: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffRepo(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeRepoInitialState(rawInitial, rawDesired *Repo) (*Repo, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeRepoDesiredState(rawDesired, rawInitial *Repo, opts ...dcl.ApplyOption) (*Repo, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Size) {
		rawDesired.Size = rawInitial.Size
	}
	if dcl.StringCanonicalize(rawDesired.Url, rawInitial.Url) {
		rawDesired.Url = rawInitial.Url
	}
	if dcl.IsZeroValue(rawDesired.PubsubConfigs) {
		rawDesired.PubsubConfigs = rawInitial.PubsubConfigs
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeRepoNewState(c *Client, rawNew, rawDesired *Repo) (*Repo, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Size) && dcl.IsEmptyValueIndirect(rawDesired.Size) {
		rawNew.Size = rawDesired.Size
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Url) && dcl.IsEmptyValueIndirect(rawDesired.Url) {
		rawNew.Url = rawDesired.Url
	} else {
		if dcl.StringCanonicalize(rawDesired.Url, rawNew.Url) {
			rawNew.Url = rawDesired.Url
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PubsubConfigs) && dcl.IsEmptyValueIndirect(rawDesired.PubsubConfigs) {
		rawNew.PubsubConfigs = rawDesired.PubsubConfigs
	} else {
		rawNew.PubsubConfigs = canonicalizeNewRepoPubsubConfigsSet(c, rawDesired.PubsubConfigs, rawNew.PubsubConfigs)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeRepoPubsubConfigs(des, initial *RepoPubsubConfigs, opts ...dcl.ApplyOption) *RepoPubsubConfigs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Topic, initial.Topic) || dcl.IsZeroValue(des.Topic) {
		des.Topic = initial.Topic
	}
	if dcl.StringCanonicalize(des.MessageFormat, initial.MessageFormat) || dcl.IsZeroValue(des.MessageFormat) {
		des.MessageFormat = initial.MessageFormat
	}
	if dcl.StringCanonicalize(des.ServiceAccountEmail, initial.ServiceAccountEmail) || dcl.IsZeroValue(des.ServiceAccountEmail) {
		des.ServiceAccountEmail = initial.ServiceAccountEmail
	}

	return des
}

func canonicalizeNewRepoPubsubConfigs(c *Client, des, nw *RepoPubsubConfigs) *RepoPubsubConfigs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Topic, nw.Topic) {
		nw.Topic = des.Topic
	}
	if dcl.StringCanonicalize(des.MessageFormat, nw.MessageFormat) {
		nw.MessageFormat = des.MessageFormat
	}
	if dcl.StringCanonicalize(des.ServiceAccountEmail, nw.ServiceAccountEmail) {
		nw.ServiceAccountEmail = des.ServiceAccountEmail
	}

	return nw
}

func canonicalizeNewRepoPubsubConfigsSet(c *Client, des, nw []RepoPubsubConfigs) []RepoPubsubConfigs {
	if des == nil {
		return nw
	}
	var reorderedNew []RepoPubsubConfigs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRepoPubsubConfigs(c, &d, &n) {
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

func canonicalizeNewRepoPubsubConfigsSlice(c *Client, des, nw []RepoPubsubConfigs) []RepoPubsubConfigs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RepoPubsubConfigs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRepoPubsubConfigs(c, &d, &n))
	}

	return items
}

type repoDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         repoApiOperation
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
func diffRepo(c *Client, desired, actual *Repo, opts ...dcl.ApplyOption) ([]repoDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []repoDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "name"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, repoDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Size, actual.Size, dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: "", FieldName: "size"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, repoDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: "", FieldName: "url"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, repoDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "project"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, repoDiff{RequiresRecreate: true, Diffs: ds})
	}

	if compareRepoPubsubConfigsSlice(c, desired.PubsubConfigs, actual.PubsubConfigs) {
		c.Config.Logger.Infof("Detected diff in PubsubConfigs.\nDESIRED: %v\nACTUAL: %v", desired.PubsubConfigs, actual.PubsubConfigs)

		toAdd, toRemove := compareRepoPubsubConfigsSets(c, desired.PubsubConfigs, actual.PubsubConfigs)
		c.Config.Logger.Infof("diff in PubsubConfigs is a set field - recomparing with set logic.\nto add: %#v\nto remove: %#v", toAdd, toRemove)
		if len(toAdd) != 0 || len(toRemove) != 0 {
			c.Config.Logger.Info("diff in PubsubConfigs persists after set logic analysis.")
			diffs = append(diffs, repoDiff{
				UpdateOp:  &updateRepoUpdateRepoOperation{},
				FieldName: "PubsubConfigs",
			})
		}

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
	var deduped []repoDiff
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
func compareRepoPubsubConfigs(c *Client, desired, actual *RepoPubsubConfigs) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Topic, actual.Topic) && !dcl.IsZeroValue(desired.Topic) {
		c.Config.Logger.Infof("Diff in Topic.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Topic), dcl.SprintResource(actual.Topic))
		return true
	}
	if !dcl.StringCanonicalize(desired.MessageFormat, actual.MessageFormat) && !dcl.IsZeroValue(desired.MessageFormat) {
		c.Config.Logger.Infof("Diff in MessageFormat.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MessageFormat), dcl.SprintResource(actual.MessageFormat))
		return true
	}
	if !dcl.StringCanonicalize(desired.ServiceAccountEmail, actual.ServiceAccountEmail) && !dcl.IsZeroValue(desired.ServiceAccountEmail) {
		c.Config.Logger.Infof("Diff in ServiceAccountEmail.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccountEmail), dcl.SprintResource(actual.ServiceAccountEmail))
		return true
	}
	return false
}

func compareRepoPubsubConfigsSlice(c *Client, desired, actual []RepoPubsubConfigs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RepoPubsubConfigs, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRepoPubsubConfigs(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RepoPubsubConfigs, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRepoPubsubConfigsMap(c *Client, desired, actual map[string]RepoPubsubConfigs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RepoPubsubConfigs, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RepoPubsubConfigs, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRepoPubsubConfigs(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RepoPubsubConfigs, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRepoPubsubConfigsSets(c *Client, desired, actual []RepoPubsubConfigs) (toAdd, toRemove []RepoPubsubConfigs) {
	if actual == nil {
		return desired, nil
	}
	desiredHashes := make(map[string]RepoPubsubConfigs)
	for _, d := range desired {
		desiredHashes[d.HashCode()] = d
	}
	actualHashes := make(map[string]RepoPubsubConfigs)
	for _, a := range actual {
		actualHashes[a.HashCode()] = a
	}
	toAdd = make([]RepoPubsubConfigs, 0)
	toRemove = make([]RepoPubsubConfigs, 0)

	for k, v := range actualHashes {
		_, found := desiredHashes[k]
		if !found {
			// backup - search linearly for equivalent value.
			for _, des := range desiredHashes {
				if !compareRepoPubsubConfigs(c, &des, &v) {
					found = true
					break
				}
			}
		}
		if !found {
			toRemove = append(toRemove, v)
		}
	}

	for k, v := range desiredHashes {
		_, found := actualHashes[k]
		if !found {
			for _, act := range actualHashes {
				if !compareRepoPubsubConfigs(c, &v, &act) {
					found = true
					break
				}
			}
		}
		if !found {
			toAdd = append(toAdd, v)
		}
	}

	return toAdd, toRemove
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Repo) urlNormalized() *Repo {
	normalized := dcl.Copy(*r).(Repo)
	normalized.Name = dcl.SelfLinkToNameWithPattern(r.Name, "projects/%s/repos/%s")
	normalized.Url = dcl.SelfLinkToName(r.Url)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Repo) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Repo) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Repo) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Repo) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateRepo" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/repos/{{name}}", "https://sourcerepo.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Repo resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Repo) marshal(c *Client) ([]byte, error) {
	m, err := expandRepo(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Repo: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalRepo decodes JSON responses into the Repo resource schema.
func unmarshalRepo(b []byte, c *Client) (*Repo, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapRepo(m, c)
}

func unmarshalMapRepo(m map[string]interface{}, c *Client) (*Repo, error) {

	return flattenRepo(c, m), nil
}

// expandRepo expands Repo into a JSON request object.
func expandRepo(c *Client, f *Repo) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveFromPattern("projects/%s/repos/%s", f.Name, f.Project, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Size; !dcl.IsEmptyValueIndirect(v) {
		m["size"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v, err := expandRepoPubsubConfig(f, f.PubsubConfigs); err != nil {
		return nil, fmt.Errorf("error expanding PubsubConfigs into pubsubConfigs: %w", err)
	} else if v != nil {
		m["pubsubConfigs"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenRepo flattens Repo from a JSON request object into the
// Repo type.
func flattenRepo(c *Client, i interface{}) *Repo {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Repo{}
	r.Name = dcl.FlattenString(m["name"])
	r.Size = dcl.FlattenInteger(m["size"])
	r.Url = dcl.FlattenString(m["url"])
	r.PubsubConfigs = flattenRepoPubsubConfig(m["pubsubConfigs"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandRepoPubsubConfigsMap expands the contents of RepoPubsubConfigs into a JSON
// request object.
func expandRepoPubsubConfigsMap(c *Client, f map[string]RepoPubsubConfigs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRepoPubsubConfigs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRepoPubsubConfigsSlice expands the contents of RepoPubsubConfigs into a JSON
// request object.
func expandRepoPubsubConfigsSlice(c *Client, f []RepoPubsubConfigs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRepoPubsubConfigs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRepoPubsubConfigsMap flattens the contents of RepoPubsubConfigs from a JSON
// response object.
func flattenRepoPubsubConfigsMap(c *Client, i interface{}) map[string]RepoPubsubConfigs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RepoPubsubConfigs{}
	}

	if len(a) == 0 {
		return map[string]RepoPubsubConfigs{}
	}

	items := make(map[string]RepoPubsubConfigs)
	for k, item := range a {
		items[k] = *flattenRepoPubsubConfigs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRepoPubsubConfigsSlice flattens the contents of RepoPubsubConfigs from a JSON
// response object.
func flattenRepoPubsubConfigsSlice(c *Client, i interface{}) []RepoPubsubConfigs {
	a, ok := i.([]interface{})
	if !ok {
		return []RepoPubsubConfigs{}
	}

	if len(a) == 0 {
		return []RepoPubsubConfigs{}
	}

	items := make([]RepoPubsubConfigs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRepoPubsubConfigs(c, item.(map[string]interface{})))
	}

	return items
}

// expandRepoPubsubConfigs expands an instance of RepoPubsubConfigs into a JSON
// request object.
func expandRepoPubsubConfigs(c *Client, f *RepoPubsubConfigs) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Topic; !dcl.IsEmptyValueIndirect(v) {
		m["topic"] = v
	}
	if v := f.MessageFormat; !dcl.IsEmptyValueIndirect(v) {
		m["messageFormat"] = v
	}
	if v := f.ServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmail"] = v
	}

	return m, nil
}

// flattenRepoPubsubConfigs flattens an instance of RepoPubsubConfigs from a JSON
// response object.
func flattenRepoPubsubConfigs(c *Client, i interface{}) *RepoPubsubConfigs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RepoPubsubConfigs{}
	r.Topic = dcl.FlattenString(m["topic"])
	r.MessageFormat = dcl.FlattenString(m["messageFormat"])
	r.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Repo) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalRepo(b, c)
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
