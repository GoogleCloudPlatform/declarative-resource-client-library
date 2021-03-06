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
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *BuildTrigger) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Filename", "Build"}, r.Filename, r.Build); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.TriggerTemplate) {
		if err := r.TriggerTemplate.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Github) {
		if err := r.Github.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Build) {
		if err := r.Build.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BuildTriggerTriggerTemplate) validate() error {
	return nil
}
func (r *BuildTriggerGithub) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PullRequest) {
		if err := r.PullRequest.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Push) {
		if err := r.Push.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BuildTriggerGithubPullRequest) validate() error {
	if err := dcl.Required(r, "branch"); err != nil {
		return err
	}
	return nil
}
func (r *BuildTriggerGithubPush) validate() error {
	return nil
}
func (r *BuildTriggerBuild) validate() error {
	if err := dcl.Required(r, "steps"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Source) {
		if err := r.Source.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BuildTriggerBuildSecrets) validate() error {
	if err := dcl.Required(r, "kmsKeyName"); err != nil {
		return err
	}
	return nil
}
func (r *BuildTriggerBuildSteps) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Timing) {
		if err := r.Timing.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PullTiming) {
		if err := r.PullTiming.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BuildTriggerBuildStepsVolumes) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "path"); err != nil {
		return err
	}
	return nil
}
func (r *BuildTriggerBuildStepsTiming) validate() error {
	return nil
}
func (r *BuildTriggerBuildStepsPullTiming) validate() error {
	return nil
}
func (r *BuildTriggerBuildSource) validate() error {
	if !dcl.IsEmptyValueIndirect(r.StorageSource) {
		if err := r.StorageSource.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RepoSource) {
		if err := r.RepoSource.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BuildTriggerBuildSourceStorageSource) validate() error {
	if err := dcl.Required(r, "bucket"); err != nil {
		return err
	}
	if err := dcl.Required(r, "object"); err != nil {
		return err
	}
	return nil
}
func (r *BuildTriggerBuildSourceRepoSource) validate() error {
	if err := dcl.Required(r, "repoName"); err != nil {
		return err
	}
	return nil
}

func buildTriggerGetURL(userBasePath string, r *BuildTrigger) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/triggers/{{name}}", "https://cloudbuild.googleapis.com/v1/", userBasePath, params), nil
}

func buildTriggerListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/triggers", "https://cloudbuild.googleapis.com/v1/", userBasePath, params), nil

}

func buildTriggerCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/triggers", "https://cloudbuild.googleapis.com/v1/", userBasePath, params), nil

}

func buildTriggerDeleteURL(userBasePath string, r *BuildTrigger) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/triggers/{{name}}", "https://cloudbuild.googleapis.com/v1/", userBasePath, params), nil
}

// buildTriggerApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type buildTriggerApiOperation interface {
	do(context.Context, *BuildTrigger, *Client) error
}

// newUpdateBuildTriggerUpdateBuildTriggerRequest creates a request for an
// BuildTrigger resource's UpdateBuildTrigger update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateBuildTriggerUpdateBuildTriggerRequest(ctx context.Context, f *BuildTrigger, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		req["tags"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		req["disabled"] = v
	}
	if v := f.Substitutions; !dcl.IsEmptyValueIndirect(v) {
		req["substitutions"] = v
	}
	if v := f.Filename; !dcl.IsEmptyValueIndirect(v) {
		req["filename"] = v
	}
	if v := f.IgnoredFiles; !dcl.IsEmptyValueIndirect(v) {
		req["ignoredFiles"] = v
	}
	if v := f.IncludedFiles; !dcl.IsEmptyValueIndirect(v) {
		req["includedFiles"] = v
	}
	if v, err := expandBuildTriggerTriggerTemplate(c, f.TriggerTemplate); err != nil {
		return nil, fmt.Errorf("error expanding TriggerTemplate into triggerTemplate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["triggerTemplate"] = v
	}
	if v, err := expandBuildTriggerGithub(c, f.Github); err != nil {
		return nil, fmt.Errorf("error expanding Github into github: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["github"] = v
	}
	if v, err := expandBuildTriggerBuild(c, f.Build); err != nil {
		return nil, fmt.Errorf("error expanding Build into build: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["build"] = v
	}
	return req, nil
}

// marshalUpdateBuildTriggerUpdateBuildTriggerRequest converts the update into
// the final JSON request body.
func marshalUpdateBuildTriggerUpdateBuildTriggerRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateBuildTriggerUpdateBuildTriggerOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateBuildTriggerUpdateBuildTriggerOperation) do(ctx context.Context, r *BuildTrigger, c *Client) error {
	_, err := c.GetBuildTrigger(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateBuildTrigger")
	if err != nil {
		return err
	}

	req, err := newUpdateBuildTriggerUpdateBuildTriggerRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateBuildTriggerUpdateBuildTriggerRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listBuildTriggerRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := buildTriggerListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != BuildTriggerMaxPage {
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

type listBuildTriggerOperation struct {
	Triggers []map[string]interface{} `json:"triggers"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listBuildTrigger(ctx context.Context, project, pageToken string, pageSize int32) ([]*BuildTrigger, string, error) {
	b, err := c.listBuildTriggerRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listBuildTriggerOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*BuildTrigger
	for _, v := range m.Triggers {
		res, err := unmarshalMapBuildTrigger(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllBuildTrigger(ctx context.Context, f func(*BuildTrigger) bool, resources []*BuildTrigger) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteBuildTrigger(ctx, res)
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

type deleteBuildTriggerOperation struct{}

func (op *deleteBuildTriggerOperation) do(ctx context.Context, r *BuildTrigger, c *Client) error {
	r, err := c.GetBuildTrigger(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("BuildTrigger not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetBuildTrigger checking for existence. error: %v", err)
		return err
	}

	u, err := buildTriggerDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete BuildTrigger: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetBuildTrigger(ctx, r.URLNormalized())
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
type createBuildTriggerOperation struct {
	response map[string]interface{}
}

func (op *createBuildTriggerOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createBuildTriggerOperation) do(ctx context.Context, r *BuildTrigger, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := buildTriggerCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetBuildTrigger(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getBuildTriggerRaw(ctx context.Context, r *BuildTrigger) ([]byte, error) {

	u, err := buildTriggerGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) buildTriggerDiffsForRawDesired(ctx context.Context, rawDesired *BuildTrigger, opts ...dcl.ApplyOption) (initial, desired *BuildTrigger, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *BuildTrigger
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*BuildTrigger); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected BuildTrigger, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetBuildTrigger(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a BuildTrigger resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve BuildTrigger resource: %v", err)
		}
		c.Config.Logger.Info("Found that BuildTrigger resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeBuildTriggerDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for BuildTrigger: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for BuildTrigger: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeBuildTriggerInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for BuildTrigger: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeBuildTriggerDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for BuildTrigger: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffBuildTrigger(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeBuildTriggerInitialState(rawInitial, rawDesired *BuildTrigger) (*BuildTrigger, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.Filename) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Build) {
			rawInitial.Filename = dcl.String("")
		}
	}

	if !dcl.IsZeroValue(rawInitial.Build) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Filename) {
			rawInitial.Build = EmptyBuildTriggerBuild
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

func canonicalizeBuildTriggerDesiredState(rawDesired, rawInitial *BuildTrigger, opts ...dcl.ApplyOption) (*BuildTrigger, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.TriggerTemplate = canonicalizeBuildTriggerTriggerTemplate(rawDesired.TriggerTemplate, nil, opts...)
		rawDesired.Github = canonicalizeBuildTriggerGithub(rawDesired.Github, nil, opts...)
		rawDesired.Build = canonicalizeBuildTriggerBuild(rawDesired.Build, nil, opts...)

		return rawDesired, nil
	}

	if rawDesired.Filename != nil || rawInitial.Filename != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Build) {
			rawDesired.Filename = nil
			rawInitial.Filename = nil
		}
	}

	if rawDesired.Build != nil || rawInitial.Build != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Filename) {
			rawDesired.Build = nil
			rawInitial.Build = nil
		}
	}

	canonicalDesired := &BuildTrigger{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Tags) {
		canonicalDesired.Tags = rawInitial.Tags
	} else {
		canonicalDesired.Tags = rawDesired.Tags
	}
	if dcl.BoolCanonicalize(rawDesired.Disabled, rawInitial.Disabled) {
		canonicalDesired.Disabled = rawInitial.Disabled
	} else {
		canonicalDesired.Disabled = rawDesired.Disabled
	}
	if dcl.IsZeroValue(rawDesired.Substitutions) {
		canonicalDesired.Substitutions = rawInitial.Substitutions
	} else {
		canonicalDesired.Substitutions = rawDesired.Substitutions
	}
	if dcl.StringCanonicalize(rawDesired.Filename, rawInitial.Filename) {
		canonicalDesired.Filename = rawInitial.Filename
	} else {
		canonicalDesired.Filename = rawDesired.Filename
	}
	if dcl.IsZeroValue(rawDesired.IgnoredFiles) {
		canonicalDesired.IgnoredFiles = rawInitial.IgnoredFiles
	} else {
		canonicalDesired.IgnoredFiles = rawDesired.IgnoredFiles
	}
	if dcl.IsZeroValue(rawDesired.IncludedFiles) {
		canonicalDesired.IncludedFiles = rawInitial.IncludedFiles
	} else {
		canonicalDesired.IncludedFiles = rawDesired.IncludedFiles
	}
	canonicalDesired.TriggerTemplate = canonicalizeBuildTriggerTriggerTemplate(rawDesired.TriggerTemplate, rawInitial.TriggerTemplate, opts...)
	canonicalDesired.Github = canonicalizeBuildTriggerGithub(rawDesired.Github, rawInitial.Github, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	canonicalDesired.Build = canonicalizeBuildTriggerBuild(rawDesired.Build, rawInitial.Build, opts...)

	return canonicalDesired, nil
}

func canonicalizeBuildTriggerNewState(c *Client, rawNew, rawDesired *BuildTrigger) (*BuildTrigger, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
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

	if dcl.IsEmptyValueIndirect(rawNew.Tags) && dcl.IsEmptyValueIndirect(rawDesired.Tags) {
		rawNew.Tags = rawDesired.Tags
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Substitutions) && dcl.IsEmptyValueIndirect(rawDesired.Substitutions) {
		rawNew.Substitutions = rawDesired.Substitutions
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Filename) && dcl.IsEmptyValueIndirect(rawDesired.Filename) {
		rawNew.Filename = rawDesired.Filename
	} else {
		if dcl.StringCanonicalize(rawDesired.Filename, rawNew.Filename) {
			rawNew.Filename = rawDesired.Filename
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IgnoredFiles) && dcl.IsEmptyValueIndirect(rawDesired.IgnoredFiles) {
		rawNew.IgnoredFiles = rawDesired.IgnoredFiles
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IncludedFiles) && dcl.IsEmptyValueIndirect(rawDesired.IncludedFiles) {
		rawNew.IncludedFiles = rawDesired.IncludedFiles
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TriggerTemplate) && dcl.IsEmptyValueIndirect(rawDesired.TriggerTemplate) {
		rawNew.TriggerTemplate = rawDesired.TriggerTemplate
	} else {
		rawNew.TriggerTemplate = canonicalizeNewBuildTriggerTriggerTemplate(c, rawDesired.TriggerTemplate, rawNew.TriggerTemplate)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Github) && dcl.IsEmptyValueIndirect(rawDesired.Github) {
		rawNew.Github = rawDesired.Github
	} else {
		rawNew.Github = canonicalizeNewBuildTriggerGithub(c, rawDesired.Github, rawNew.Github)
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Build) && dcl.IsEmptyValueIndirect(rawDesired.Build) {
		rawNew.Build = rawDesired.Build
	} else {
		rawNew.Build = canonicalizeNewBuildTriggerBuild(c, rawDesired.Build, rawNew.Build)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
		if dcl.StringCanonicalize(rawDesired.CreateTime, rawNew.CreateTime) {
			rawNew.CreateTime = rawDesired.CreateTime
		}
	}

	return rawNew, nil
}

func canonicalizeBuildTriggerTriggerTemplate(des, initial *BuildTriggerTriggerTemplate, opts ...dcl.ApplyOption) *BuildTriggerTriggerTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.RepoName) {
		des.RepoName = dcl.String("default")
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerTriggerTemplate{}

	if dcl.StringCanonicalize(des.ProjectId, initial.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		cDes.ProjectId = initial.ProjectId
	} else {
		cDes.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.RepoName, initial.RepoName) || dcl.IsZeroValue(des.RepoName) {
		cDes.RepoName = initial.RepoName
	} else {
		cDes.RepoName = des.RepoName
	}
	if dcl.StringCanonicalize(des.BranchName, initial.BranchName) || dcl.IsZeroValue(des.BranchName) {
		cDes.BranchName = initial.BranchName
	} else {
		cDes.BranchName = des.BranchName
	}
	if dcl.StringCanonicalize(des.TagName, initial.TagName) || dcl.IsZeroValue(des.TagName) {
		cDes.TagName = initial.TagName
	} else {
		cDes.TagName = des.TagName
	}
	if dcl.StringCanonicalize(des.CommitSha, initial.CommitSha) || dcl.IsZeroValue(des.CommitSha) {
		cDes.CommitSha = initial.CommitSha
	} else {
		cDes.CommitSha = des.CommitSha
	}
	if dcl.StringCanonicalize(des.Dir, initial.Dir) || dcl.IsZeroValue(des.Dir) {
		cDes.Dir = initial.Dir
	} else {
		cDes.Dir = des.Dir
	}
	if dcl.BoolCanonicalize(des.InvertRegex, initial.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
		cDes.InvertRegex = initial.InvertRegex
	} else {
		cDes.InvertRegex = des.InvertRegex
	}

	return cDes
}

func canonicalizeNewBuildTriggerTriggerTemplate(c *Client, des, nw *BuildTriggerTriggerTemplate) *BuildTriggerTriggerTemplate {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.RepoName) {
		nw.RepoName = dcl.String("default")
	}

	if dcl.StringCanonicalize(des.ProjectId, nw.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.RepoName, nw.RepoName) {
		nw.RepoName = des.RepoName
	}
	if dcl.StringCanonicalize(des.BranchName, nw.BranchName) {
		nw.BranchName = des.BranchName
	}
	if dcl.StringCanonicalize(des.TagName, nw.TagName) {
		nw.TagName = des.TagName
	}
	if dcl.StringCanonicalize(des.CommitSha, nw.CommitSha) {
		nw.CommitSha = des.CommitSha
	}
	if dcl.StringCanonicalize(des.Dir, nw.Dir) {
		nw.Dir = des.Dir
	}
	if dcl.BoolCanonicalize(des.InvertRegex, nw.InvertRegex) {
		nw.InvertRegex = des.InvertRegex
	}

	return nw
}

func canonicalizeNewBuildTriggerTriggerTemplateSet(c *Client, des, nw []BuildTriggerTriggerTemplate) []BuildTriggerTriggerTemplate {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerTriggerTemplate
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerTriggerTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerTriggerTemplateSlice(c *Client, des, nw []BuildTriggerTriggerTemplate) []BuildTriggerTriggerTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerTriggerTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerTriggerTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerGithub(des, initial *BuildTriggerGithub, opts ...dcl.ApplyOption) *BuildTriggerGithub {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerGithub{}

	if dcl.StringCanonicalize(des.Owner, initial.Owner) || dcl.IsZeroValue(des.Owner) {
		cDes.Owner = initial.Owner
	} else {
		cDes.Owner = des.Owner
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	cDes.PullRequest = canonicalizeBuildTriggerGithubPullRequest(des.PullRequest, initial.PullRequest, opts...)
	cDes.Push = canonicalizeBuildTriggerGithubPush(des.Push, initial.Push, opts...)

	return cDes
}

func canonicalizeNewBuildTriggerGithub(c *Client, des, nw *BuildTriggerGithub) *BuildTriggerGithub {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Owner, nw.Owner) {
		nw.Owner = des.Owner
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.PullRequest = canonicalizeNewBuildTriggerGithubPullRequest(c, des.PullRequest, nw.PullRequest)
	nw.Push = canonicalizeNewBuildTriggerGithubPush(c, des.Push, nw.Push)

	return nw
}

func canonicalizeNewBuildTriggerGithubSet(c *Client, des, nw []BuildTriggerGithub) []BuildTriggerGithub {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerGithub
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerGithubNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerGithubSlice(c *Client, des, nw []BuildTriggerGithub) []BuildTriggerGithub {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerGithub
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerGithub(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerGithubPullRequest(des, initial *BuildTriggerGithubPullRequest, opts ...dcl.ApplyOption) *BuildTriggerGithubPullRequest {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerGithubPullRequest{}

	if dcl.StringCanonicalize(des.Branch, initial.Branch) || dcl.IsZeroValue(des.Branch) {
		cDes.Branch = initial.Branch
	} else {
		cDes.Branch = des.Branch
	}
	if dcl.IsZeroValue(des.CommentControl) {
		des.CommentControl = initial.CommentControl
	} else {
		cDes.CommentControl = des.CommentControl
	}
	if dcl.BoolCanonicalize(des.InvertRegex, initial.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
		cDes.InvertRegex = initial.InvertRegex
	} else {
		cDes.InvertRegex = des.InvertRegex
	}

	return cDes
}

func canonicalizeNewBuildTriggerGithubPullRequest(c *Client, des, nw *BuildTriggerGithubPullRequest) *BuildTriggerGithubPullRequest {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Branch, nw.Branch) {
		nw.Branch = des.Branch
	}
	if dcl.IsZeroValue(nw.CommentControl) {
		nw.CommentControl = des.CommentControl
	}
	if dcl.BoolCanonicalize(des.InvertRegex, nw.InvertRegex) {
		nw.InvertRegex = des.InvertRegex
	}

	return nw
}

func canonicalizeNewBuildTriggerGithubPullRequestSet(c *Client, des, nw []BuildTriggerGithubPullRequest) []BuildTriggerGithubPullRequest {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerGithubPullRequest
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerGithubPullRequestNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerGithubPullRequestSlice(c *Client, des, nw []BuildTriggerGithubPullRequest) []BuildTriggerGithubPullRequest {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerGithubPullRequest
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerGithubPullRequest(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerGithubPush(des, initial *BuildTriggerGithubPush, opts ...dcl.ApplyOption) *BuildTriggerGithubPush {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerGithubPush{}

	if dcl.StringCanonicalize(des.Branch, initial.Branch) || dcl.IsZeroValue(des.Branch) {
		cDes.Branch = initial.Branch
	} else {
		cDes.Branch = des.Branch
	}
	if dcl.StringCanonicalize(des.Tag, initial.Tag) || dcl.IsZeroValue(des.Tag) {
		cDes.Tag = initial.Tag
	} else {
		cDes.Tag = des.Tag
	}
	if dcl.BoolCanonicalize(des.InvertRegex, initial.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
		cDes.InvertRegex = initial.InvertRegex
	} else {
		cDes.InvertRegex = des.InvertRegex
	}

	return cDes
}

func canonicalizeNewBuildTriggerGithubPush(c *Client, des, nw *BuildTriggerGithubPush) *BuildTriggerGithubPush {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Branch, nw.Branch) {
		nw.Branch = des.Branch
	}
	if dcl.StringCanonicalize(des.Tag, nw.Tag) {
		nw.Tag = des.Tag
	}
	if dcl.BoolCanonicalize(des.InvertRegex, nw.InvertRegex) {
		nw.InvertRegex = des.InvertRegex
	}

	return nw
}

func canonicalizeNewBuildTriggerGithubPushSet(c *Client, des, nw []BuildTriggerGithubPush) []BuildTriggerGithubPush {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerGithubPush
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerGithubPushNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerGithubPushSlice(c *Client, des, nw []BuildTriggerGithubPush) []BuildTriggerGithubPush {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerGithubPush
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerGithubPush(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerBuild(des, initial *BuildTriggerBuild, opts ...dcl.ApplyOption) *BuildTriggerBuild {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.Timeout) {
		des.Timeout = dcl.String("600s")
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerBuild{}

	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	} else {
		cDes.Tags = des.Tags
	}
	if dcl.IsZeroValue(des.Images) {
		des.Images = initial.Images
	} else {
		cDes.Images = des.Images
	}
	if dcl.IsZeroValue(des.Substitutions) {
		des.Substitutions = initial.Substitutions
	} else {
		cDes.Substitutions = des.Substitutions
	}
	if dcl.StringCanonicalize(des.QueueTtl, initial.QueueTtl) || dcl.IsZeroValue(des.QueueTtl) {
		cDes.QueueTtl = initial.QueueTtl
	} else {
		cDes.QueueTtl = des.QueueTtl
	}
	if dcl.StringCanonicalize(des.LogsBucket, initial.LogsBucket) || dcl.IsZeroValue(des.LogsBucket) {
		cDes.LogsBucket = initial.LogsBucket
	} else {
		cDes.LogsBucket = des.LogsBucket
	}
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		cDes.Timeout = initial.Timeout
	} else {
		cDes.Timeout = des.Timeout
	}
	if dcl.IsZeroValue(des.Secrets) {
		des.Secrets = initial.Secrets
	} else {
		cDes.Secrets = des.Secrets
	}
	if dcl.IsZeroValue(des.Steps) {
		des.Steps = initial.Steps
	} else {
		cDes.Steps = des.Steps
	}
	cDes.Source = canonicalizeBuildTriggerBuildSource(des.Source, initial.Source, opts...)

	return cDes
}

func canonicalizeNewBuildTriggerBuild(c *Client, des, nw *BuildTriggerBuild) *BuildTriggerBuild {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Timeout) {
		nw.Timeout = dcl.String("600s")
	}

	if dcl.IsZeroValue(nw.Tags) {
		nw.Tags = des.Tags
	}
	if dcl.IsZeroValue(nw.Images) {
		nw.Images = des.Images
	}
	if dcl.IsZeroValue(nw.Substitutions) {
		nw.Substitutions = des.Substitutions
	}
	if dcl.StringCanonicalize(des.QueueTtl, nw.QueueTtl) {
		nw.QueueTtl = des.QueueTtl
	}
	if dcl.StringCanonicalize(des.LogsBucket, nw.LogsBucket) {
		nw.LogsBucket = des.LogsBucket
	}
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}
	nw.Secrets = canonicalizeNewBuildTriggerBuildSecretsSlice(c, des.Secrets, nw.Secrets)
	nw.Steps = canonicalizeNewBuildTriggerBuildStepsSlice(c, des.Steps, nw.Steps)
	nw.Source = canonicalizeNewBuildTriggerBuildSource(c, des.Source, nw.Source)

	return nw
}

func canonicalizeNewBuildTriggerBuildSet(c *Client, des, nw []BuildTriggerBuild) []BuildTriggerBuild {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerBuild
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerBuildNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerBuildSlice(c *Client, des, nw []BuildTriggerBuild) []BuildTriggerBuild {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerBuild
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuild(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerBuildSecrets(des, initial *BuildTriggerBuildSecrets, opts ...dcl.ApplyOption) *BuildTriggerBuildSecrets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerBuildSecrets{}

	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}
	if dcl.IsZeroValue(des.SecretEnv) {
		des.SecretEnv = initial.SecretEnv
	} else {
		cDes.SecretEnv = des.SecretEnv
	}

	return cDes
}

func canonicalizeNewBuildTriggerBuildSecrets(c *Client, des, nw *BuildTriggerBuildSecrets) *BuildTriggerBuildSecrets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
	}
	if dcl.IsZeroValue(nw.SecretEnv) {
		nw.SecretEnv = des.SecretEnv
	}

	return nw
}

func canonicalizeNewBuildTriggerBuildSecretsSet(c *Client, des, nw []BuildTriggerBuildSecrets) []BuildTriggerBuildSecrets {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerBuildSecrets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerBuildSecretsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerBuildSecretsSlice(c *Client, des, nw []BuildTriggerBuildSecrets) []BuildTriggerBuildSecrets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerBuildSecrets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuildSecrets(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerBuildSteps(des, initial *BuildTriggerBuildSteps, opts ...dcl.ApplyOption) *BuildTriggerBuildSteps {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.Timeout) {
		des.Timeout = dcl.String("300s")
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerBuildSteps{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.Env) {
		des.Env = initial.Env
	} else {
		cDes.Env = des.Env
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	} else {
		cDes.Args = des.Args
	}
	if dcl.StringCanonicalize(des.Dir, initial.Dir) || dcl.IsZeroValue(des.Dir) {
		cDes.Dir = initial.Dir
	} else {
		cDes.Dir = des.Dir
	}
	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		cDes.Id = initial.Id
	} else {
		cDes.Id = des.Id
	}
	if dcl.IsZeroValue(des.WaitFor) {
		des.WaitFor = initial.WaitFor
	} else {
		cDes.WaitFor = des.WaitFor
	}
	if dcl.StringCanonicalize(des.Entrypoint, initial.Entrypoint) || dcl.IsZeroValue(des.Entrypoint) {
		cDes.Entrypoint = initial.Entrypoint
	} else {
		cDes.Entrypoint = des.Entrypoint
	}
	if dcl.IsZeroValue(des.SecretEnv) {
		des.SecretEnv = initial.SecretEnv
	} else {
		cDes.SecretEnv = des.SecretEnv
	}
	if dcl.IsZeroValue(des.Volumes) {
		des.Volumes = initial.Volumes
	} else {
		cDes.Volumes = des.Volumes
	}
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		cDes.Timeout = initial.Timeout
	} else {
		cDes.Timeout = des.Timeout
	}

	return cDes
}

func canonicalizeNewBuildTriggerBuildSteps(c *Client, des, nw *BuildTriggerBuildSteps) *BuildTriggerBuildSteps {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Timeout) {
		nw.Timeout = dcl.String("300s")
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.IsZeroValue(nw.Env) {
		nw.Env = des.Env
	}
	if dcl.IsZeroValue(nw.Args) {
		nw.Args = des.Args
	}
	if dcl.StringCanonicalize(des.Dir, nw.Dir) {
		nw.Dir = des.Dir
	}
	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.IsZeroValue(nw.WaitFor) {
		nw.WaitFor = des.WaitFor
	}
	if dcl.StringCanonicalize(des.Entrypoint, nw.Entrypoint) {
		nw.Entrypoint = des.Entrypoint
	}
	if dcl.IsZeroValue(nw.SecretEnv) {
		nw.SecretEnv = des.SecretEnv
	}
	nw.Volumes = canonicalizeNewBuildTriggerBuildStepsVolumesSlice(c, des.Volumes, nw.Volumes)
	nw.Timing = canonicalizeNewBuildTriggerBuildStepsTiming(c, des.Timing, nw.Timing)
	nw.PullTiming = canonicalizeNewBuildTriggerBuildStepsPullTiming(c, des.PullTiming, nw.PullTiming)
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}
	if dcl.IsZeroValue(nw.Status) {
		nw.Status = des.Status
	}

	return nw
}

func canonicalizeNewBuildTriggerBuildStepsSet(c *Client, des, nw []BuildTriggerBuildSteps) []BuildTriggerBuildSteps {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerBuildSteps
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerBuildStepsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerBuildStepsSlice(c *Client, des, nw []BuildTriggerBuildSteps) []BuildTriggerBuildSteps {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerBuildSteps
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuildSteps(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerBuildStepsVolumes(des, initial *BuildTriggerBuildStepsVolumes, opts ...dcl.ApplyOption) *BuildTriggerBuildStepsVolumes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerBuildStepsVolumes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}

	return cDes
}

func canonicalizeNewBuildTriggerBuildStepsVolumes(c *Client, des, nw *BuildTriggerBuildStepsVolumes) *BuildTriggerBuildStepsVolumes {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewBuildTriggerBuildStepsVolumesSet(c *Client, des, nw []BuildTriggerBuildStepsVolumes) []BuildTriggerBuildStepsVolumes {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerBuildStepsVolumes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerBuildStepsVolumesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerBuildStepsVolumesSlice(c *Client, des, nw []BuildTriggerBuildStepsVolumes) []BuildTriggerBuildStepsVolumes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerBuildStepsVolumes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuildStepsVolumes(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerBuildStepsTiming(des, initial *BuildTriggerBuildStepsTiming, opts ...dcl.ApplyOption) *BuildTriggerBuildStepsTiming {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerBuildStepsTiming{}

	if dcl.StringCanonicalize(des.StartTime, initial.StartTime) || dcl.IsZeroValue(des.StartTime) {
		cDes.StartTime = initial.StartTime
	} else {
		cDes.StartTime = des.StartTime
	}
	if dcl.StringCanonicalize(des.EndTime, initial.EndTime) || dcl.IsZeroValue(des.EndTime) {
		cDes.EndTime = initial.EndTime
	} else {
		cDes.EndTime = des.EndTime
	}

	return cDes
}

func canonicalizeNewBuildTriggerBuildStepsTiming(c *Client, des, nw *BuildTriggerBuildStepsTiming) *BuildTriggerBuildStepsTiming {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.StartTime, nw.StartTime) {
		nw.StartTime = des.StartTime
	}
	if dcl.StringCanonicalize(des.EndTime, nw.EndTime) {
		nw.EndTime = des.EndTime
	}

	return nw
}

func canonicalizeNewBuildTriggerBuildStepsTimingSet(c *Client, des, nw []BuildTriggerBuildStepsTiming) []BuildTriggerBuildStepsTiming {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerBuildStepsTiming
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerBuildStepsTimingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerBuildStepsTimingSlice(c *Client, des, nw []BuildTriggerBuildStepsTiming) []BuildTriggerBuildStepsTiming {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerBuildStepsTiming
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuildStepsTiming(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerBuildStepsPullTiming(des, initial *BuildTriggerBuildStepsPullTiming, opts ...dcl.ApplyOption) *BuildTriggerBuildStepsPullTiming {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerBuildStepsPullTiming{}

	if dcl.StringCanonicalize(des.StartTime, initial.StartTime) || dcl.IsZeroValue(des.StartTime) {
		cDes.StartTime = initial.StartTime
	} else {
		cDes.StartTime = des.StartTime
	}
	if dcl.StringCanonicalize(des.EndTime, initial.EndTime) || dcl.IsZeroValue(des.EndTime) {
		cDes.EndTime = initial.EndTime
	} else {
		cDes.EndTime = des.EndTime
	}

	return cDes
}

func canonicalizeNewBuildTriggerBuildStepsPullTiming(c *Client, des, nw *BuildTriggerBuildStepsPullTiming) *BuildTriggerBuildStepsPullTiming {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.StartTime, nw.StartTime) {
		nw.StartTime = des.StartTime
	}
	if dcl.StringCanonicalize(des.EndTime, nw.EndTime) {
		nw.EndTime = des.EndTime
	}

	return nw
}

func canonicalizeNewBuildTriggerBuildStepsPullTimingSet(c *Client, des, nw []BuildTriggerBuildStepsPullTiming) []BuildTriggerBuildStepsPullTiming {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerBuildStepsPullTiming
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerBuildStepsPullTimingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerBuildStepsPullTimingSlice(c *Client, des, nw []BuildTriggerBuildStepsPullTiming) []BuildTriggerBuildStepsPullTiming {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerBuildStepsPullTiming
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuildStepsPullTiming(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerBuildSource(des, initial *BuildTriggerBuildSource, opts ...dcl.ApplyOption) *BuildTriggerBuildSource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerBuildSource{}

	cDes.StorageSource = canonicalizeBuildTriggerBuildSourceStorageSource(des.StorageSource, initial.StorageSource, opts...)
	cDes.RepoSource = canonicalizeBuildTriggerBuildSourceRepoSource(des.RepoSource, initial.RepoSource, opts...)

	return cDes
}

func canonicalizeNewBuildTriggerBuildSource(c *Client, des, nw *BuildTriggerBuildSource) *BuildTriggerBuildSource {
	if des == nil || nw == nil {
		return nw
	}

	nw.StorageSource = canonicalizeNewBuildTriggerBuildSourceStorageSource(c, des.StorageSource, nw.StorageSource)
	nw.RepoSource = canonicalizeNewBuildTriggerBuildSourceRepoSource(c, des.RepoSource, nw.RepoSource)

	return nw
}

func canonicalizeNewBuildTriggerBuildSourceSet(c *Client, des, nw []BuildTriggerBuildSource) []BuildTriggerBuildSource {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerBuildSource
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerBuildSourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerBuildSourceSlice(c *Client, des, nw []BuildTriggerBuildSource) []BuildTriggerBuildSource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerBuildSource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuildSource(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerBuildSourceStorageSource(des, initial *BuildTriggerBuildSourceStorageSource, opts ...dcl.ApplyOption) *BuildTriggerBuildSourceStorageSource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerBuildSourceStorageSource{}

	if dcl.StringCanonicalize(des.Bucket, initial.Bucket) || dcl.IsZeroValue(des.Bucket) {
		cDes.Bucket = initial.Bucket
	} else {
		cDes.Bucket = des.Bucket
	}
	if dcl.StringCanonicalize(des.Object, initial.Object) || dcl.IsZeroValue(des.Object) {
		cDes.Object = initial.Object
	} else {
		cDes.Object = des.Object
	}
	if dcl.StringCanonicalize(des.Generation, initial.Generation) || dcl.IsZeroValue(des.Generation) {
		cDes.Generation = initial.Generation
	} else {
		cDes.Generation = des.Generation
	}

	return cDes
}

func canonicalizeNewBuildTriggerBuildSourceStorageSource(c *Client, des, nw *BuildTriggerBuildSourceStorageSource) *BuildTriggerBuildSourceStorageSource {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Bucket, nw.Bucket) {
		nw.Bucket = des.Bucket
	}
	if dcl.StringCanonicalize(des.Object, nw.Object) {
		nw.Object = des.Object
	}
	if dcl.StringCanonicalize(des.Generation, nw.Generation) {
		nw.Generation = des.Generation
	}

	return nw
}

func canonicalizeNewBuildTriggerBuildSourceStorageSourceSet(c *Client, des, nw []BuildTriggerBuildSourceStorageSource) []BuildTriggerBuildSourceStorageSource {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerBuildSourceStorageSource
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerBuildSourceStorageSourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerBuildSourceStorageSourceSlice(c *Client, des, nw []BuildTriggerBuildSourceStorageSource) []BuildTriggerBuildSourceStorageSource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerBuildSourceStorageSource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuildSourceStorageSource(c, &d, &n))
	}

	return items
}

func canonicalizeBuildTriggerBuildSourceRepoSource(des, initial *BuildTriggerBuildSourceRepoSource, opts ...dcl.ApplyOption) *BuildTriggerBuildSourceRepoSource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BuildTriggerBuildSourceRepoSource{}

	if dcl.StringCanonicalize(des.ProjectId, initial.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		cDes.ProjectId = initial.ProjectId
	} else {
		cDes.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.RepoName, initial.RepoName) || dcl.IsZeroValue(des.RepoName) {
		cDes.RepoName = initial.RepoName
	} else {
		cDes.RepoName = des.RepoName
	}
	if dcl.StringCanonicalize(des.BranchName, initial.BranchName) || dcl.IsZeroValue(des.BranchName) {
		cDes.BranchName = initial.BranchName
	} else {
		cDes.BranchName = des.BranchName
	}
	if dcl.StringCanonicalize(des.TagName, initial.TagName) || dcl.IsZeroValue(des.TagName) {
		cDes.TagName = initial.TagName
	} else {
		cDes.TagName = des.TagName
	}
	if dcl.StringCanonicalize(des.CommitSha, initial.CommitSha) || dcl.IsZeroValue(des.CommitSha) {
		cDes.CommitSha = initial.CommitSha
	} else {
		cDes.CommitSha = des.CommitSha
	}
	if dcl.StringCanonicalize(des.Dir, initial.Dir) || dcl.IsZeroValue(des.Dir) {
		cDes.Dir = initial.Dir
	} else {
		cDes.Dir = des.Dir
	}
	if dcl.BoolCanonicalize(des.InvertRegex, initial.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
		cDes.InvertRegex = initial.InvertRegex
	} else {
		cDes.InvertRegex = des.InvertRegex
	}
	if dcl.IsZeroValue(des.Substitutions) {
		des.Substitutions = initial.Substitutions
	} else {
		cDes.Substitutions = des.Substitutions
	}

	return cDes
}

func canonicalizeNewBuildTriggerBuildSourceRepoSource(c *Client, des, nw *BuildTriggerBuildSourceRepoSource) *BuildTriggerBuildSourceRepoSource {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ProjectId, nw.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.RepoName, nw.RepoName) {
		nw.RepoName = des.RepoName
	}
	if dcl.StringCanonicalize(des.BranchName, nw.BranchName) {
		nw.BranchName = des.BranchName
	}
	if dcl.StringCanonicalize(des.TagName, nw.TagName) {
		nw.TagName = des.TagName
	}
	if dcl.StringCanonicalize(des.CommitSha, nw.CommitSha) {
		nw.CommitSha = des.CommitSha
	}
	if dcl.StringCanonicalize(des.Dir, nw.Dir) {
		nw.Dir = des.Dir
	}
	if dcl.BoolCanonicalize(des.InvertRegex, nw.InvertRegex) {
		nw.InvertRegex = des.InvertRegex
	}
	if dcl.IsZeroValue(nw.Substitutions) {
		nw.Substitutions = des.Substitutions
	}

	return nw
}

func canonicalizeNewBuildTriggerBuildSourceRepoSourceSet(c *Client, des, nw []BuildTriggerBuildSourceRepoSource) []BuildTriggerBuildSourceRepoSource {
	if des == nil {
		return nw
	}
	var reorderedNew []BuildTriggerBuildSourceRepoSource
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBuildTriggerBuildSourceRepoSourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBuildTriggerBuildSourceRepoSourceSlice(c *Client, des, nw []BuildTriggerBuildSourceRepoSource) []BuildTriggerBuildSourceRepoSource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BuildTriggerBuildSourceRepoSource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuildSourceRepoSource(c, &d, &n))
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
func diffBuildTrigger(c *Client, desired, actual *BuildTrigger, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tags, actual.Tags, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Substitutions, actual.Substitutions, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Substitutions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Filename, actual.Filename, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Filename")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IgnoredFiles, actual.IgnoredFiles, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("IgnoredFiles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludedFiles, actual.IncludedFiles, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("IncludedFiles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TriggerTemplate, actual.TriggerTemplate, dcl.Info{ObjectFunction: compareBuildTriggerTriggerTemplateNewStyle, EmptyObject: EmptyBuildTriggerTriggerTemplate, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("TriggerTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Github, actual.Github, dcl.Info{ObjectFunction: compareBuildTriggerGithubNewStyle, EmptyObject: EmptyBuildTriggerGithub, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Github")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Build, actual.Build, dcl.Info{ObjectFunction: compareBuildTriggerBuildNewStyle, EmptyObject: EmptyBuildTriggerBuild, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Build")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareBuildTriggerTriggerTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerTriggerTemplate)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerTriggerTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerTriggerTemplate or *BuildTriggerTriggerTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerTriggerTemplate)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerTriggerTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerTriggerTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RepoName, actual.RepoName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("RepoName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BranchName, actual.BranchName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("BranchName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TagName, actual.TagName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("TagName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CommitSha, actual.CommitSha, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("CommitSha")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dir, actual.Dir, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Dir")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InvertRegex, actual.InvertRegex, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("InvertRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerGithubNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerGithub)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerGithub)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerGithub or *BuildTriggerGithub", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerGithub)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerGithub)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerGithub", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Owner, actual.Owner, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Owner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PullRequest, actual.PullRequest, dcl.Info{ObjectFunction: compareBuildTriggerGithubPullRequestNewStyle, EmptyObject: EmptyBuildTriggerGithubPullRequest, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("PullRequest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Push, actual.Push, dcl.Info{ObjectFunction: compareBuildTriggerGithubPushNewStyle, EmptyObject: EmptyBuildTriggerGithubPush, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Push")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerGithubPullRequestNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerGithubPullRequest)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerGithubPullRequest)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerGithubPullRequest or *BuildTriggerGithubPullRequest", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerGithubPullRequest)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerGithubPullRequest)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerGithubPullRequest", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Branch, actual.Branch, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Branch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CommentControl, actual.CommentControl, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("CommentControl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InvertRegex, actual.InvertRegex, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("InvertRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerGithubPushNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerGithubPush)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerGithubPush)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerGithubPush or *BuildTriggerGithubPush", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerGithubPush)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerGithubPush)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerGithubPush", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Branch, actual.Branch, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Branch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Tag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InvertRegex, actual.InvertRegex, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("InvertRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerBuildNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerBuild)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerBuild)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuild or *BuildTriggerBuild", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerBuild)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerBuild)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuild", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Tags, actual.Tags, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Images, actual.Images, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Images")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Substitutions, actual.Substitutions, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Substitutions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.QueueTtl, actual.QueueTtl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("QueueTtl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LogsBucket, actual.LogsBucket, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("LogsBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Secrets, actual.Secrets, dcl.Info{ObjectFunction: compareBuildTriggerBuildSecretsNewStyle, EmptyObject: EmptyBuildTriggerBuildSecrets, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Secrets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Steps, actual.Steps, dcl.Info{ObjectFunction: compareBuildTriggerBuildStepsNewStyle, EmptyObject: EmptyBuildTriggerBuildSteps, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Steps")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Source, actual.Source, dcl.Info{ObjectFunction: compareBuildTriggerBuildSourceNewStyle, EmptyObject: EmptyBuildTriggerBuildSource, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Source")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerBuildSecretsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerBuildSecrets)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerBuildSecrets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSecrets or *BuildTriggerBuildSecrets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerBuildSecrets)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerBuildSecrets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSecrets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecretEnv, actual.SecretEnv, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("SecretEnv")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerBuildStepsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerBuildSteps)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerBuildSteps)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSteps or *BuildTriggerBuildSteps", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerBuildSteps)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerBuildSteps)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSteps", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Env, actual.Env, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Env")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dir, actual.Dir, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Dir")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WaitFor, actual.WaitFor, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("WaitFor")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Entrypoint, actual.Entrypoint, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Entrypoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecretEnv, actual.SecretEnv, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("SecretEnv")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Volumes, actual.Volumes, dcl.Info{ObjectFunction: compareBuildTriggerBuildStepsVolumesNewStyle, EmptyObject: EmptyBuildTriggerBuildStepsVolumes, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Volumes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timing, actual.Timing, dcl.Info{OutputOnly: true, ObjectFunction: compareBuildTriggerBuildStepsTimingNewStyle, EmptyObject: EmptyBuildTriggerBuildStepsTiming, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Timing")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PullTiming, actual.PullTiming, dcl.Info{OutputOnly: true, ObjectFunction: compareBuildTriggerBuildStepsPullTimingNewStyle, EmptyObject: EmptyBuildTriggerBuildStepsPullTiming, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PullTiming")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerBuildStepsVolumesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerBuildStepsVolumes)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerBuildStepsVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildStepsVolumes or *BuildTriggerBuildStepsVolumes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerBuildStepsVolumes)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerBuildStepsVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildStepsVolumes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerBuildStepsTimingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerBuildStepsTiming)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerBuildStepsTiming)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildStepsTiming or *BuildTriggerBuildStepsTiming", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerBuildStepsTiming)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerBuildStepsTiming)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildStepsTiming", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StartTime, actual.StartTime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndTime, actual.EndTime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EndTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerBuildStepsPullTimingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerBuildStepsPullTiming)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerBuildStepsPullTiming)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildStepsPullTiming or *BuildTriggerBuildStepsPullTiming", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerBuildStepsPullTiming)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerBuildStepsPullTiming)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildStepsPullTiming", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StartTime, actual.StartTime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndTime, actual.EndTime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EndTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerBuildSourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerBuildSource)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerBuildSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSource or *BuildTriggerBuildSource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerBuildSource)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerBuildSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StorageSource, actual.StorageSource, dcl.Info{ObjectFunction: compareBuildTriggerBuildSourceStorageSourceNewStyle, EmptyObject: EmptyBuildTriggerBuildSourceStorageSource, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("StorageSource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RepoSource, actual.RepoSource, dcl.Info{ObjectFunction: compareBuildTriggerBuildSourceRepoSourceNewStyle, EmptyObject: EmptyBuildTriggerBuildSourceRepoSource, OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("RepoSource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerBuildSourceStorageSourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerBuildSourceStorageSource)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerBuildSourceStorageSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSourceStorageSource or *BuildTriggerBuildSourceStorageSource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerBuildSourceStorageSource)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerBuildSourceStorageSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSourceStorageSource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Bucket, actual.Bucket, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Bucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Object, actual.Object, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Object")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBuildTriggerBuildSourceRepoSourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BuildTriggerBuildSourceRepoSource)
	if !ok {
		desiredNotPointer, ok := d.(BuildTriggerBuildSourceRepoSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSourceRepoSource or *BuildTriggerBuildSourceRepoSource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BuildTriggerBuildSourceRepoSource)
	if !ok {
		actualNotPointer, ok := a.(BuildTriggerBuildSourceRepoSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BuildTriggerBuildSourceRepoSource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RepoName, actual.RepoName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("RepoName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BranchName, actual.BranchName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("BranchName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TagName, actual.TagName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("TagName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CommitSha, actual.CommitSha, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("CommitSha")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dir, actual.Dir, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Dir")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InvertRegex, actual.InvertRegex, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("InvertRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Substitutions, actual.Substitutions, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBuildTriggerUpdateBuildTriggerOperation")}, fn.AddNest("Substitutions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *BuildTrigger) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *BuildTrigger) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *BuildTrigger) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *BuildTrigger) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "UpdateBuildTrigger" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/triggers/{{name}}", "https://cloudbuild.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the BuildTrigger resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *BuildTrigger) marshal(c *Client) ([]byte, error) {
	m, err := expandBuildTrigger(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling BuildTrigger: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalBuildTrigger decodes JSON responses into the BuildTrigger resource schema.
func unmarshalBuildTrigger(b []byte, c *Client) (*BuildTrigger, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapBuildTrigger(m, c)
}

func unmarshalMapBuildTrigger(m map[string]interface{}, c *Client) (*BuildTrigger, error) {

	return flattenBuildTrigger(c, m), nil
}

// expandBuildTrigger expands BuildTrigger into a JSON request object.
func expandBuildTrigger(c *Client, f *BuildTrigger) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}
	if v := f.Substitutions; !dcl.IsEmptyValueIndirect(v) {
		m["substitutions"] = v
	}
	if v := f.Filename; !dcl.IsEmptyValueIndirect(v) {
		m["filename"] = v
	}
	if v := f.IgnoredFiles; !dcl.IsEmptyValueIndirect(v) {
		m["ignoredFiles"] = v
	}
	if v := f.IncludedFiles; !dcl.IsEmptyValueIndirect(v) {
		m["includedFiles"] = v
	}
	if v, err := expandBuildTriggerTriggerTemplate(c, f.TriggerTemplate); err != nil {
		return nil, fmt.Errorf("error expanding TriggerTemplate into triggerTemplate: %w", err)
	} else if v != nil {
		m["triggerTemplate"] = v
	}
	if v, err := expandBuildTriggerGithub(c, f.Github); err != nil {
		return nil, fmt.Errorf("error expanding Github into github: %w", err)
	} else if v != nil {
		m["github"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := expandBuildTriggerBuild(c, f.Build); err != nil {
		return nil, fmt.Errorf("error expanding Build into build: %w", err)
	} else if v != nil {
		m["build"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}

	return m, nil
}

// flattenBuildTrigger flattens BuildTrigger from a JSON request object into the
// BuildTrigger type.
func flattenBuildTrigger(c *Client, i interface{}) *BuildTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &BuildTrigger{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Tags = dcl.FlattenStringSlice(m["tags"])
	res.Disabled = dcl.FlattenBool(m["disabled"])
	res.Substitutions = dcl.FlattenKeyValuePairs(m["substitutions"])
	res.Filename = dcl.FlattenString(m["filename"])
	res.IgnoredFiles = dcl.FlattenStringSlice(m["ignoredFiles"])
	res.IncludedFiles = dcl.FlattenStringSlice(m["includedFiles"])
	res.TriggerTemplate = flattenBuildTriggerTriggerTemplate(c, m["triggerTemplate"])
	res.Github = flattenBuildTriggerGithub(c, m["github"])
	res.Project = dcl.FlattenString(m["project"])
	res.Build = flattenBuildTriggerBuild(c, m["build"])
	res.Id = dcl.FlattenString(m["id"])
	res.CreateTime = dcl.FlattenString(m["createTime"])

	return res
}

// expandBuildTriggerTriggerTemplateMap expands the contents of BuildTriggerTriggerTemplate into a JSON
// request object.
func expandBuildTriggerTriggerTemplateMap(c *Client, f map[string]BuildTriggerTriggerTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerTriggerTemplate(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerTriggerTemplateSlice expands the contents of BuildTriggerTriggerTemplate into a JSON
// request object.
func expandBuildTriggerTriggerTemplateSlice(c *Client, f []BuildTriggerTriggerTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerTriggerTemplate(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerTriggerTemplateMap flattens the contents of BuildTriggerTriggerTemplate from a JSON
// response object.
func flattenBuildTriggerTriggerTemplateMap(c *Client, i interface{}) map[string]BuildTriggerTriggerTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerTriggerTemplate{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerTriggerTemplate{}
	}

	items := make(map[string]BuildTriggerTriggerTemplate)
	for k, item := range a {
		items[k] = *flattenBuildTriggerTriggerTemplate(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerTriggerTemplateSlice flattens the contents of BuildTriggerTriggerTemplate from a JSON
// response object.
func flattenBuildTriggerTriggerTemplateSlice(c *Client, i interface{}) []BuildTriggerTriggerTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerTriggerTemplate{}
	}

	if len(a) == 0 {
		return []BuildTriggerTriggerTemplate{}
	}

	items := make([]BuildTriggerTriggerTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerTriggerTemplate(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerTriggerTemplate expands an instance of BuildTriggerTriggerTemplate into a JSON
// request object.
func expandBuildTriggerTriggerTemplate(c *Client, f *BuildTriggerTriggerTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.RepoName; !dcl.IsEmptyValueIndirect(v) {
		m["repoName"] = v
	}
	if v := f.BranchName; !dcl.IsEmptyValueIndirect(v) {
		m["branchName"] = v
	}
	if v := f.TagName; !dcl.IsEmptyValueIndirect(v) {
		m["tagName"] = v
	}
	if v := f.CommitSha; !dcl.IsEmptyValueIndirect(v) {
		m["commitSha"] = v
	}
	if v := f.Dir; !dcl.IsEmptyValueIndirect(v) {
		m["dir"] = v
	}
	if v := f.InvertRegex; !dcl.IsEmptyValueIndirect(v) {
		m["invertRegex"] = v
	}

	return m, nil
}

// flattenBuildTriggerTriggerTemplate flattens an instance of BuildTriggerTriggerTemplate from a JSON
// response object.
func flattenBuildTriggerTriggerTemplate(c *Client, i interface{}) *BuildTriggerTriggerTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerTriggerTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerTriggerTemplate
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.RepoName = dcl.FlattenString(m["repoName"])
	if dcl.IsEmptyValueIndirect(m["repoName"]) {
		c.Config.Logger.Info("Using default value for repoName.")
		r.RepoName = dcl.String("default")
	}
	r.BranchName = dcl.FlattenString(m["branchName"])
	r.TagName = dcl.FlattenString(m["tagName"])
	r.CommitSha = dcl.FlattenString(m["commitSha"])
	r.Dir = dcl.FlattenString(m["dir"])
	r.InvertRegex = dcl.FlattenBool(m["invertRegex"])

	return r
}

// expandBuildTriggerGithubMap expands the contents of BuildTriggerGithub into a JSON
// request object.
func expandBuildTriggerGithubMap(c *Client, f map[string]BuildTriggerGithub) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerGithub(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerGithubSlice expands the contents of BuildTriggerGithub into a JSON
// request object.
func expandBuildTriggerGithubSlice(c *Client, f []BuildTriggerGithub) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerGithub(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerGithubMap flattens the contents of BuildTriggerGithub from a JSON
// response object.
func flattenBuildTriggerGithubMap(c *Client, i interface{}) map[string]BuildTriggerGithub {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerGithub{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerGithub{}
	}

	items := make(map[string]BuildTriggerGithub)
	for k, item := range a {
		items[k] = *flattenBuildTriggerGithub(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerGithubSlice flattens the contents of BuildTriggerGithub from a JSON
// response object.
func flattenBuildTriggerGithubSlice(c *Client, i interface{}) []BuildTriggerGithub {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerGithub{}
	}

	if len(a) == 0 {
		return []BuildTriggerGithub{}
	}

	items := make([]BuildTriggerGithub, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerGithub(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerGithub expands an instance of BuildTriggerGithub into a JSON
// request object.
func expandBuildTriggerGithub(c *Client, f *BuildTriggerGithub) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Owner; !dcl.IsEmptyValueIndirect(v) {
		m["owner"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandBuildTriggerGithubPullRequest(c, f.PullRequest); err != nil {
		return nil, fmt.Errorf("error expanding PullRequest into pullRequest: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pullRequest"] = v
	}
	if v, err := expandBuildTriggerGithubPush(c, f.Push); err != nil {
		return nil, fmt.Errorf("error expanding Push into push: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["push"] = v
	}

	return m, nil
}

// flattenBuildTriggerGithub flattens an instance of BuildTriggerGithub from a JSON
// response object.
func flattenBuildTriggerGithub(c *Client, i interface{}) *BuildTriggerGithub {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerGithub{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerGithub
	}
	r.Owner = dcl.FlattenString(m["owner"])
	r.Name = dcl.FlattenString(m["name"])
	r.PullRequest = flattenBuildTriggerGithubPullRequest(c, m["pullRequest"])
	r.Push = flattenBuildTriggerGithubPush(c, m["push"])

	return r
}

// expandBuildTriggerGithubPullRequestMap expands the contents of BuildTriggerGithubPullRequest into a JSON
// request object.
func expandBuildTriggerGithubPullRequestMap(c *Client, f map[string]BuildTriggerGithubPullRequest) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerGithubPullRequest(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerGithubPullRequestSlice expands the contents of BuildTriggerGithubPullRequest into a JSON
// request object.
func expandBuildTriggerGithubPullRequestSlice(c *Client, f []BuildTriggerGithubPullRequest) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerGithubPullRequest(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerGithubPullRequestMap flattens the contents of BuildTriggerGithubPullRequest from a JSON
// response object.
func flattenBuildTriggerGithubPullRequestMap(c *Client, i interface{}) map[string]BuildTriggerGithubPullRequest {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerGithubPullRequest{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerGithubPullRequest{}
	}

	items := make(map[string]BuildTriggerGithubPullRequest)
	for k, item := range a {
		items[k] = *flattenBuildTriggerGithubPullRequest(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerGithubPullRequestSlice flattens the contents of BuildTriggerGithubPullRequest from a JSON
// response object.
func flattenBuildTriggerGithubPullRequestSlice(c *Client, i interface{}) []BuildTriggerGithubPullRequest {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerGithubPullRequest{}
	}

	if len(a) == 0 {
		return []BuildTriggerGithubPullRequest{}
	}

	items := make([]BuildTriggerGithubPullRequest, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerGithubPullRequest(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerGithubPullRequest expands an instance of BuildTriggerGithubPullRequest into a JSON
// request object.
func expandBuildTriggerGithubPullRequest(c *Client, f *BuildTriggerGithubPullRequest) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Branch; !dcl.IsEmptyValueIndirect(v) {
		m["branch"] = v
	}
	if v := f.CommentControl; !dcl.IsEmptyValueIndirect(v) {
		m["commentControl"] = v
	}
	if v := f.InvertRegex; !dcl.IsEmptyValueIndirect(v) {
		m["invertRegex"] = v
	}

	return m, nil
}

// flattenBuildTriggerGithubPullRequest flattens an instance of BuildTriggerGithubPullRequest from a JSON
// response object.
func flattenBuildTriggerGithubPullRequest(c *Client, i interface{}) *BuildTriggerGithubPullRequest {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerGithubPullRequest{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerGithubPullRequest
	}
	r.Branch = dcl.FlattenString(m["branch"])
	r.CommentControl = flattenBuildTriggerGithubPullRequestCommentControlEnum(m["commentControl"])
	r.InvertRegex = dcl.FlattenBool(m["invertRegex"])

	return r
}

// expandBuildTriggerGithubPushMap expands the contents of BuildTriggerGithubPush into a JSON
// request object.
func expandBuildTriggerGithubPushMap(c *Client, f map[string]BuildTriggerGithubPush) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerGithubPush(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerGithubPushSlice expands the contents of BuildTriggerGithubPush into a JSON
// request object.
func expandBuildTriggerGithubPushSlice(c *Client, f []BuildTriggerGithubPush) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerGithubPush(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerGithubPushMap flattens the contents of BuildTriggerGithubPush from a JSON
// response object.
func flattenBuildTriggerGithubPushMap(c *Client, i interface{}) map[string]BuildTriggerGithubPush {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerGithubPush{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerGithubPush{}
	}

	items := make(map[string]BuildTriggerGithubPush)
	for k, item := range a {
		items[k] = *flattenBuildTriggerGithubPush(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerGithubPushSlice flattens the contents of BuildTriggerGithubPush from a JSON
// response object.
func flattenBuildTriggerGithubPushSlice(c *Client, i interface{}) []BuildTriggerGithubPush {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerGithubPush{}
	}

	if len(a) == 0 {
		return []BuildTriggerGithubPush{}
	}

	items := make([]BuildTriggerGithubPush, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerGithubPush(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerGithubPush expands an instance of BuildTriggerGithubPush into a JSON
// request object.
func expandBuildTriggerGithubPush(c *Client, f *BuildTriggerGithubPush) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Branch; !dcl.IsEmptyValueIndirect(v) {
		m["branch"] = v
	}
	if v := f.Tag; !dcl.IsEmptyValueIndirect(v) {
		m["tag"] = v
	}
	if v := f.InvertRegex; !dcl.IsEmptyValueIndirect(v) {
		m["invertRegex"] = v
	}

	return m, nil
}

// flattenBuildTriggerGithubPush flattens an instance of BuildTriggerGithubPush from a JSON
// response object.
func flattenBuildTriggerGithubPush(c *Client, i interface{}) *BuildTriggerGithubPush {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerGithubPush{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerGithubPush
	}
	r.Branch = dcl.FlattenString(m["branch"])
	r.Tag = dcl.FlattenString(m["tag"])
	r.InvertRegex = dcl.FlattenBool(m["invertRegex"])

	return r
}

// expandBuildTriggerBuildMap expands the contents of BuildTriggerBuild into a JSON
// request object.
func expandBuildTriggerBuildMap(c *Client, f map[string]BuildTriggerBuild) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerBuild(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerBuildSlice expands the contents of BuildTriggerBuild into a JSON
// request object.
func expandBuildTriggerBuildSlice(c *Client, f []BuildTriggerBuild) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerBuild(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerBuildMap flattens the contents of BuildTriggerBuild from a JSON
// response object.
func flattenBuildTriggerBuildMap(c *Client, i interface{}) map[string]BuildTriggerBuild {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuild{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuild{}
	}

	items := make(map[string]BuildTriggerBuild)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuild(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerBuildSlice flattens the contents of BuildTriggerBuild from a JSON
// response object.
func flattenBuildTriggerBuildSlice(c *Client, i interface{}) []BuildTriggerBuild {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuild{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuild{}
	}

	items := make([]BuildTriggerBuild, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuild(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerBuild expands an instance of BuildTriggerBuild into a JSON
// request object.
func expandBuildTriggerBuild(c *Client, f *BuildTriggerBuild) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v := f.Images; !dcl.IsEmptyValueIndirect(v) {
		m["images"] = v
	}
	if v := f.Substitutions; !dcl.IsEmptyValueIndirect(v) {
		m["substitutions"] = v
	}
	if v := f.QueueTtl; !dcl.IsEmptyValueIndirect(v) {
		m["queueTtl"] = v
	}
	if v := f.LogsBucket; !dcl.IsEmptyValueIndirect(v) {
		m["logsBucket"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v, err := expandBuildTriggerBuildSecretsSlice(c, f.Secrets); err != nil {
		return nil, fmt.Errorf("error expanding Secrets into secrets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secrets"] = v
	}
	if v, err := expandBuildTriggerBuildStepsSlice(c, f.Steps); err != nil {
		return nil, fmt.Errorf("error expanding Steps into steps: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["steps"] = v
	}
	if v, err := expandBuildTriggerBuildSource(c, f.Source); err != nil {
		return nil, fmt.Errorf("error expanding Source into source: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["source"] = v
	}

	return m, nil
}

// flattenBuildTriggerBuild flattens an instance of BuildTriggerBuild from a JSON
// response object.
func flattenBuildTriggerBuild(c *Client, i interface{}) *BuildTriggerBuild {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerBuild{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerBuild
	}
	r.Tags = dcl.FlattenStringSlice(m["tags"])
	r.Images = dcl.FlattenStringSlice(m["images"])
	r.Substitutions = dcl.FlattenKeyValuePairs(m["substitutions"])
	r.QueueTtl = dcl.FlattenString(m["queueTtl"])
	r.LogsBucket = dcl.FlattenString(m["logsBucket"])
	r.Timeout = dcl.FlattenString(m["timeout"])
	if dcl.IsEmptyValueIndirect(m["timeout"]) {
		c.Config.Logger.Info("Using default value for timeout.")
		r.Timeout = dcl.String("600s")
	}
	r.Secrets = flattenBuildTriggerBuildSecretsSlice(c, m["secrets"])
	r.Steps = flattenBuildTriggerBuildStepsSlice(c, m["steps"])
	r.Source = flattenBuildTriggerBuildSource(c, m["source"])

	return r
}

// expandBuildTriggerBuildSecretsMap expands the contents of BuildTriggerBuildSecrets into a JSON
// request object.
func expandBuildTriggerBuildSecretsMap(c *Client, f map[string]BuildTriggerBuildSecrets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerBuildSecrets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerBuildSecretsSlice expands the contents of BuildTriggerBuildSecrets into a JSON
// request object.
func expandBuildTriggerBuildSecretsSlice(c *Client, f []BuildTriggerBuildSecrets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerBuildSecrets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerBuildSecretsMap flattens the contents of BuildTriggerBuildSecrets from a JSON
// response object.
func flattenBuildTriggerBuildSecretsMap(c *Client, i interface{}) map[string]BuildTriggerBuildSecrets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuildSecrets{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuildSecrets{}
	}

	items := make(map[string]BuildTriggerBuildSecrets)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuildSecrets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerBuildSecretsSlice flattens the contents of BuildTriggerBuildSecrets from a JSON
// response object.
func flattenBuildTriggerBuildSecretsSlice(c *Client, i interface{}) []BuildTriggerBuildSecrets {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuildSecrets{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuildSecrets{}
	}

	items := make([]BuildTriggerBuildSecrets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuildSecrets(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerBuildSecrets expands an instance of BuildTriggerBuildSecrets into a JSON
// request object.
func expandBuildTriggerBuildSecrets(c *Client, f *BuildTriggerBuildSecrets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}
	if v := f.SecretEnv; !dcl.IsEmptyValueIndirect(v) {
		m["secretEnv"] = v
	}

	return m, nil
}

// flattenBuildTriggerBuildSecrets flattens an instance of BuildTriggerBuildSecrets from a JSON
// response object.
func flattenBuildTriggerBuildSecrets(c *Client, i interface{}) *BuildTriggerBuildSecrets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerBuildSecrets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerBuildSecrets
	}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	r.SecretEnv = dcl.FlattenKeyValuePairs(m["secretEnv"])

	return r
}

// expandBuildTriggerBuildStepsMap expands the contents of BuildTriggerBuildSteps into a JSON
// request object.
func expandBuildTriggerBuildStepsMap(c *Client, f map[string]BuildTriggerBuildSteps) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerBuildSteps(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerBuildStepsSlice expands the contents of BuildTriggerBuildSteps into a JSON
// request object.
func expandBuildTriggerBuildStepsSlice(c *Client, f []BuildTriggerBuildSteps) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerBuildSteps(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerBuildStepsMap flattens the contents of BuildTriggerBuildSteps from a JSON
// response object.
func flattenBuildTriggerBuildStepsMap(c *Client, i interface{}) map[string]BuildTriggerBuildSteps {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuildSteps{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuildSteps{}
	}

	items := make(map[string]BuildTriggerBuildSteps)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuildSteps(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerBuildStepsSlice flattens the contents of BuildTriggerBuildSteps from a JSON
// response object.
func flattenBuildTriggerBuildStepsSlice(c *Client, i interface{}) []BuildTriggerBuildSteps {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuildSteps{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuildSteps{}
	}

	items := make([]BuildTriggerBuildSteps, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuildSteps(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerBuildSteps expands an instance of BuildTriggerBuildSteps into a JSON
// request object.
func expandBuildTriggerBuildSteps(c *Client, f *BuildTriggerBuildSteps) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Env; !dcl.IsEmptyValueIndirect(v) {
		m["env"] = v
	}
	if v := f.Args; !dcl.IsEmptyValueIndirect(v) {
		m["args"] = v
	}
	if v := f.Dir; !dcl.IsEmptyValueIndirect(v) {
		m["dir"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.WaitFor; !dcl.IsEmptyValueIndirect(v) {
		m["waitFor"] = v
	}
	if v := f.Entrypoint; !dcl.IsEmptyValueIndirect(v) {
		m["entrypoint"] = v
	}
	if v := f.SecretEnv; !dcl.IsEmptyValueIndirect(v) {
		m["secretEnv"] = v
	}
	if v, err := expandBuildTriggerBuildStepsVolumesSlice(c, f.Volumes); err != nil {
		return nil, fmt.Errorf("error expanding Volumes into volumes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["volumes"] = v
	}
	if v, err := expandBuildTriggerBuildStepsTiming(c, f.Timing); err != nil {
		return nil, fmt.Errorf("error expanding Timing into timing: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timing"] = v
	}
	if v, err := expandBuildTriggerBuildStepsPullTiming(c, f.PullTiming); err != nil {
		return nil, fmt.Errorf("error expanding PullTiming into pullTiming: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pullTiming"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}

	return m, nil
}

// flattenBuildTriggerBuildSteps flattens an instance of BuildTriggerBuildSteps from a JSON
// response object.
func flattenBuildTriggerBuildSteps(c *Client, i interface{}) *BuildTriggerBuildSteps {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerBuildSteps{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerBuildSteps
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Env = dcl.FlattenStringSlice(m["env"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.Dir = dcl.FlattenString(m["dir"])
	r.Id = dcl.FlattenString(m["id"])
	r.WaitFor = dcl.FlattenStringSlice(m["waitFor"])
	r.Entrypoint = dcl.FlattenString(m["entrypoint"])
	r.SecretEnv = dcl.FlattenStringSlice(m["secretEnv"])
	r.Volumes = flattenBuildTriggerBuildStepsVolumesSlice(c, m["volumes"])
	r.Timing = flattenBuildTriggerBuildStepsTiming(c, m["timing"])
	r.PullTiming = flattenBuildTriggerBuildStepsPullTiming(c, m["pullTiming"])
	r.Timeout = dcl.FlattenString(m["timeout"])
	if dcl.IsEmptyValueIndirect(m["timeout"]) {
		c.Config.Logger.Info("Using default value for timeout.")
		r.Timeout = dcl.String("300s")
	}
	r.Status = flattenBuildTriggerBuildStepsStatusEnum(m["status"])

	return r
}

// expandBuildTriggerBuildStepsVolumesMap expands the contents of BuildTriggerBuildStepsVolumes into a JSON
// request object.
func expandBuildTriggerBuildStepsVolumesMap(c *Client, f map[string]BuildTriggerBuildStepsVolumes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerBuildStepsVolumes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerBuildStepsVolumesSlice expands the contents of BuildTriggerBuildStepsVolumes into a JSON
// request object.
func expandBuildTriggerBuildStepsVolumesSlice(c *Client, f []BuildTriggerBuildStepsVolumes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerBuildStepsVolumes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerBuildStepsVolumesMap flattens the contents of BuildTriggerBuildStepsVolumes from a JSON
// response object.
func flattenBuildTriggerBuildStepsVolumesMap(c *Client, i interface{}) map[string]BuildTriggerBuildStepsVolumes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuildStepsVolumes{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuildStepsVolumes{}
	}

	items := make(map[string]BuildTriggerBuildStepsVolumes)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuildStepsVolumes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerBuildStepsVolumesSlice flattens the contents of BuildTriggerBuildStepsVolumes from a JSON
// response object.
func flattenBuildTriggerBuildStepsVolumesSlice(c *Client, i interface{}) []BuildTriggerBuildStepsVolumes {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuildStepsVolumes{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuildStepsVolumes{}
	}

	items := make([]BuildTriggerBuildStepsVolumes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuildStepsVolumes(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerBuildStepsVolumes expands an instance of BuildTriggerBuildStepsVolumes into a JSON
// request object.
func expandBuildTriggerBuildStepsVolumes(c *Client, f *BuildTriggerBuildStepsVolumes) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}

	return m, nil
}

// flattenBuildTriggerBuildStepsVolumes flattens an instance of BuildTriggerBuildStepsVolumes from a JSON
// response object.
func flattenBuildTriggerBuildStepsVolumes(c *Client, i interface{}) *BuildTriggerBuildStepsVolumes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerBuildStepsVolumes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerBuildStepsVolumes
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Path = dcl.FlattenString(m["path"])

	return r
}

// expandBuildTriggerBuildStepsTimingMap expands the contents of BuildTriggerBuildStepsTiming into a JSON
// request object.
func expandBuildTriggerBuildStepsTimingMap(c *Client, f map[string]BuildTriggerBuildStepsTiming) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerBuildStepsTiming(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerBuildStepsTimingSlice expands the contents of BuildTriggerBuildStepsTiming into a JSON
// request object.
func expandBuildTriggerBuildStepsTimingSlice(c *Client, f []BuildTriggerBuildStepsTiming) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerBuildStepsTiming(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerBuildStepsTimingMap flattens the contents of BuildTriggerBuildStepsTiming from a JSON
// response object.
func flattenBuildTriggerBuildStepsTimingMap(c *Client, i interface{}) map[string]BuildTriggerBuildStepsTiming {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuildStepsTiming{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuildStepsTiming{}
	}

	items := make(map[string]BuildTriggerBuildStepsTiming)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuildStepsTiming(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerBuildStepsTimingSlice flattens the contents of BuildTriggerBuildStepsTiming from a JSON
// response object.
func flattenBuildTriggerBuildStepsTimingSlice(c *Client, i interface{}) []BuildTriggerBuildStepsTiming {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuildStepsTiming{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuildStepsTiming{}
	}

	items := make([]BuildTriggerBuildStepsTiming, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuildStepsTiming(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerBuildStepsTiming expands an instance of BuildTriggerBuildStepsTiming into a JSON
// request object.
func expandBuildTriggerBuildStepsTiming(c *Client, f *BuildTriggerBuildStepsTiming) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.StartTime; !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.EndTime; !dcl.IsEmptyValueIndirect(v) {
		m["endTime"] = v
	}

	return m, nil
}

// flattenBuildTriggerBuildStepsTiming flattens an instance of BuildTriggerBuildStepsTiming from a JSON
// response object.
func flattenBuildTriggerBuildStepsTiming(c *Client, i interface{}) *BuildTriggerBuildStepsTiming {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerBuildStepsTiming{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerBuildStepsTiming
	}
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.EndTime = dcl.FlattenString(m["endTime"])

	return r
}

// expandBuildTriggerBuildStepsPullTimingMap expands the contents of BuildTriggerBuildStepsPullTiming into a JSON
// request object.
func expandBuildTriggerBuildStepsPullTimingMap(c *Client, f map[string]BuildTriggerBuildStepsPullTiming) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerBuildStepsPullTiming(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerBuildStepsPullTimingSlice expands the contents of BuildTriggerBuildStepsPullTiming into a JSON
// request object.
func expandBuildTriggerBuildStepsPullTimingSlice(c *Client, f []BuildTriggerBuildStepsPullTiming) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerBuildStepsPullTiming(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerBuildStepsPullTimingMap flattens the contents of BuildTriggerBuildStepsPullTiming from a JSON
// response object.
func flattenBuildTriggerBuildStepsPullTimingMap(c *Client, i interface{}) map[string]BuildTriggerBuildStepsPullTiming {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuildStepsPullTiming{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuildStepsPullTiming{}
	}

	items := make(map[string]BuildTriggerBuildStepsPullTiming)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuildStepsPullTiming(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerBuildStepsPullTimingSlice flattens the contents of BuildTriggerBuildStepsPullTiming from a JSON
// response object.
func flattenBuildTriggerBuildStepsPullTimingSlice(c *Client, i interface{}) []BuildTriggerBuildStepsPullTiming {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuildStepsPullTiming{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuildStepsPullTiming{}
	}

	items := make([]BuildTriggerBuildStepsPullTiming, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuildStepsPullTiming(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerBuildStepsPullTiming expands an instance of BuildTriggerBuildStepsPullTiming into a JSON
// request object.
func expandBuildTriggerBuildStepsPullTiming(c *Client, f *BuildTriggerBuildStepsPullTiming) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.StartTime; !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.EndTime; !dcl.IsEmptyValueIndirect(v) {
		m["endTime"] = v
	}

	return m, nil
}

// flattenBuildTriggerBuildStepsPullTiming flattens an instance of BuildTriggerBuildStepsPullTiming from a JSON
// response object.
func flattenBuildTriggerBuildStepsPullTiming(c *Client, i interface{}) *BuildTriggerBuildStepsPullTiming {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerBuildStepsPullTiming{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerBuildStepsPullTiming
	}
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.EndTime = dcl.FlattenString(m["endTime"])

	return r
}

// expandBuildTriggerBuildSourceMap expands the contents of BuildTriggerBuildSource into a JSON
// request object.
func expandBuildTriggerBuildSourceMap(c *Client, f map[string]BuildTriggerBuildSource) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerBuildSource(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerBuildSourceSlice expands the contents of BuildTriggerBuildSource into a JSON
// request object.
func expandBuildTriggerBuildSourceSlice(c *Client, f []BuildTriggerBuildSource) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerBuildSource(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerBuildSourceMap flattens the contents of BuildTriggerBuildSource from a JSON
// response object.
func flattenBuildTriggerBuildSourceMap(c *Client, i interface{}) map[string]BuildTriggerBuildSource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuildSource{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuildSource{}
	}

	items := make(map[string]BuildTriggerBuildSource)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuildSource(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerBuildSourceSlice flattens the contents of BuildTriggerBuildSource from a JSON
// response object.
func flattenBuildTriggerBuildSourceSlice(c *Client, i interface{}) []BuildTriggerBuildSource {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuildSource{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuildSource{}
	}

	items := make([]BuildTriggerBuildSource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuildSource(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerBuildSource expands an instance of BuildTriggerBuildSource into a JSON
// request object.
func expandBuildTriggerBuildSource(c *Client, f *BuildTriggerBuildSource) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandBuildTriggerBuildSourceStorageSource(c, f.StorageSource); err != nil {
		return nil, fmt.Errorf("error expanding StorageSource into storageSource: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["storageSource"] = v
	}
	if v, err := expandBuildTriggerBuildSourceRepoSource(c, f.RepoSource); err != nil {
		return nil, fmt.Errorf("error expanding RepoSource into repoSource: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["repoSource"] = v
	}

	return m, nil
}

// flattenBuildTriggerBuildSource flattens an instance of BuildTriggerBuildSource from a JSON
// response object.
func flattenBuildTriggerBuildSource(c *Client, i interface{}) *BuildTriggerBuildSource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerBuildSource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerBuildSource
	}
	r.StorageSource = flattenBuildTriggerBuildSourceStorageSource(c, m["storageSource"])
	r.RepoSource = flattenBuildTriggerBuildSourceRepoSource(c, m["repoSource"])

	return r
}

// expandBuildTriggerBuildSourceStorageSourceMap expands the contents of BuildTriggerBuildSourceStorageSource into a JSON
// request object.
func expandBuildTriggerBuildSourceStorageSourceMap(c *Client, f map[string]BuildTriggerBuildSourceStorageSource) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerBuildSourceStorageSource(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerBuildSourceStorageSourceSlice expands the contents of BuildTriggerBuildSourceStorageSource into a JSON
// request object.
func expandBuildTriggerBuildSourceStorageSourceSlice(c *Client, f []BuildTriggerBuildSourceStorageSource) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerBuildSourceStorageSource(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerBuildSourceStorageSourceMap flattens the contents of BuildTriggerBuildSourceStorageSource from a JSON
// response object.
func flattenBuildTriggerBuildSourceStorageSourceMap(c *Client, i interface{}) map[string]BuildTriggerBuildSourceStorageSource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuildSourceStorageSource{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuildSourceStorageSource{}
	}

	items := make(map[string]BuildTriggerBuildSourceStorageSource)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuildSourceStorageSource(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerBuildSourceStorageSourceSlice flattens the contents of BuildTriggerBuildSourceStorageSource from a JSON
// response object.
func flattenBuildTriggerBuildSourceStorageSourceSlice(c *Client, i interface{}) []BuildTriggerBuildSourceStorageSource {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuildSourceStorageSource{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuildSourceStorageSource{}
	}

	items := make([]BuildTriggerBuildSourceStorageSource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuildSourceStorageSource(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerBuildSourceStorageSource expands an instance of BuildTriggerBuildSourceStorageSource into a JSON
// request object.
func expandBuildTriggerBuildSourceStorageSource(c *Client, f *BuildTriggerBuildSourceStorageSource) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Bucket; !dcl.IsEmptyValueIndirect(v) {
		m["bucket"] = v
	}
	if v := f.Object; !dcl.IsEmptyValueIndirect(v) {
		m["object"] = v
	}
	if v := f.Generation; !dcl.IsEmptyValueIndirect(v) {
		m["generation"] = v
	}

	return m, nil
}

// flattenBuildTriggerBuildSourceStorageSource flattens an instance of BuildTriggerBuildSourceStorageSource from a JSON
// response object.
func flattenBuildTriggerBuildSourceStorageSource(c *Client, i interface{}) *BuildTriggerBuildSourceStorageSource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerBuildSourceStorageSource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerBuildSourceStorageSource
	}
	r.Bucket = dcl.FlattenString(m["bucket"])
	r.Object = dcl.FlattenString(m["object"])
	r.Generation = dcl.FlattenString(m["generation"])

	return r
}

// expandBuildTriggerBuildSourceRepoSourceMap expands the contents of BuildTriggerBuildSourceRepoSource into a JSON
// request object.
func expandBuildTriggerBuildSourceRepoSourceMap(c *Client, f map[string]BuildTriggerBuildSourceRepoSource) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBuildTriggerBuildSourceRepoSource(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBuildTriggerBuildSourceRepoSourceSlice expands the contents of BuildTriggerBuildSourceRepoSource into a JSON
// request object.
func expandBuildTriggerBuildSourceRepoSourceSlice(c *Client, f []BuildTriggerBuildSourceRepoSource) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBuildTriggerBuildSourceRepoSource(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBuildTriggerBuildSourceRepoSourceMap flattens the contents of BuildTriggerBuildSourceRepoSource from a JSON
// response object.
func flattenBuildTriggerBuildSourceRepoSourceMap(c *Client, i interface{}) map[string]BuildTriggerBuildSourceRepoSource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuildSourceRepoSource{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuildSourceRepoSource{}
	}

	items := make(map[string]BuildTriggerBuildSourceRepoSource)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuildSourceRepoSource(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBuildTriggerBuildSourceRepoSourceSlice flattens the contents of BuildTriggerBuildSourceRepoSource from a JSON
// response object.
func flattenBuildTriggerBuildSourceRepoSourceSlice(c *Client, i interface{}) []BuildTriggerBuildSourceRepoSource {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuildSourceRepoSource{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuildSourceRepoSource{}
	}

	items := make([]BuildTriggerBuildSourceRepoSource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuildSourceRepoSource(c, item.(map[string]interface{})))
	}

	return items
}

// expandBuildTriggerBuildSourceRepoSource expands an instance of BuildTriggerBuildSourceRepoSource into a JSON
// request object.
func expandBuildTriggerBuildSourceRepoSource(c *Client, f *BuildTriggerBuildSourceRepoSource) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.RepoName; !dcl.IsEmptyValueIndirect(v) {
		m["repoName"] = v
	}
	if v := f.BranchName; !dcl.IsEmptyValueIndirect(v) {
		m["branchName"] = v
	}
	if v := f.TagName; !dcl.IsEmptyValueIndirect(v) {
		m["tagName"] = v
	}
	if v := f.CommitSha; !dcl.IsEmptyValueIndirect(v) {
		m["commitSha"] = v
	}
	if v := f.Dir; !dcl.IsEmptyValueIndirect(v) {
		m["dir"] = v
	}
	if v := f.InvertRegex; !dcl.IsEmptyValueIndirect(v) {
		m["invertRegex"] = v
	}
	if v := f.Substitutions; !dcl.IsEmptyValueIndirect(v) {
		m["substitutions"] = v
	}

	return m, nil
}

// flattenBuildTriggerBuildSourceRepoSource flattens an instance of BuildTriggerBuildSourceRepoSource from a JSON
// response object.
func flattenBuildTriggerBuildSourceRepoSource(c *Client, i interface{}) *BuildTriggerBuildSourceRepoSource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BuildTriggerBuildSourceRepoSource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBuildTriggerBuildSourceRepoSource
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.RepoName = dcl.FlattenString(m["repoName"])
	r.BranchName = dcl.FlattenString(m["branchName"])
	r.TagName = dcl.FlattenString(m["tagName"])
	r.CommitSha = dcl.FlattenString(m["commitSha"])
	r.Dir = dcl.FlattenString(m["dir"])
	r.InvertRegex = dcl.FlattenBool(m["invertRegex"])
	r.Substitutions = dcl.FlattenKeyValuePairs(m["substitutions"])

	return r
}

// flattenBuildTriggerGithubPullRequestCommentControlEnumMap flattens the contents of BuildTriggerGithubPullRequestCommentControlEnum from a JSON
// response object.
func flattenBuildTriggerGithubPullRequestCommentControlEnumMap(c *Client, i interface{}) map[string]BuildTriggerGithubPullRequestCommentControlEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerGithubPullRequestCommentControlEnum{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerGithubPullRequestCommentControlEnum{}
	}

	items := make(map[string]BuildTriggerGithubPullRequestCommentControlEnum)
	for k, item := range a {
		items[k] = *flattenBuildTriggerGithubPullRequestCommentControlEnum(item.(interface{}))
	}

	return items
}

// flattenBuildTriggerGithubPullRequestCommentControlEnumSlice flattens the contents of BuildTriggerGithubPullRequestCommentControlEnum from a JSON
// response object.
func flattenBuildTriggerGithubPullRequestCommentControlEnumSlice(c *Client, i interface{}) []BuildTriggerGithubPullRequestCommentControlEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerGithubPullRequestCommentControlEnum{}
	}

	if len(a) == 0 {
		return []BuildTriggerGithubPullRequestCommentControlEnum{}
	}

	items := make([]BuildTriggerGithubPullRequestCommentControlEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerGithubPullRequestCommentControlEnum(item.(interface{})))
	}

	return items
}

// flattenBuildTriggerGithubPullRequestCommentControlEnum asserts that an interface is a string, and returns a
// pointer to a *BuildTriggerGithubPullRequestCommentControlEnum with the same value as that string.
func flattenBuildTriggerGithubPullRequestCommentControlEnum(i interface{}) *BuildTriggerGithubPullRequestCommentControlEnum {
	s, ok := i.(string)
	if !ok {
		return BuildTriggerGithubPullRequestCommentControlEnumRef("")
	}

	return BuildTriggerGithubPullRequestCommentControlEnumRef(s)
}

// flattenBuildTriggerBuildStepsStatusEnumMap flattens the contents of BuildTriggerBuildStepsStatusEnum from a JSON
// response object.
func flattenBuildTriggerBuildStepsStatusEnumMap(c *Client, i interface{}) map[string]BuildTriggerBuildStepsStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BuildTriggerBuildStepsStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]BuildTriggerBuildStepsStatusEnum{}
	}

	items := make(map[string]BuildTriggerBuildStepsStatusEnum)
	for k, item := range a {
		items[k] = *flattenBuildTriggerBuildStepsStatusEnum(item.(interface{}))
	}

	return items
}

// flattenBuildTriggerBuildStepsStatusEnumSlice flattens the contents of BuildTriggerBuildStepsStatusEnum from a JSON
// response object.
func flattenBuildTriggerBuildStepsStatusEnumSlice(c *Client, i interface{}) []BuildTriggerBuildStepsStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BuildTriggerBuildStepsStatusEnum{}
	}

	if len(a) == 0 {
		return []BuildTriggerBuildStepsStatusEnum{}
	}

	items := make([]BuildTriggerBuildStepsStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBuildTriggerBuildStepsStatusEnum(item.(interface{})))
	}

	return items
}

// flattenBuildTriggerBuildStepsStatusEnum asserts that an interface is a string, and returns a
// pointer to a *BuildTriggerBuildStepsStatusEnum with the same value as that string.
func flattenBuildTriggerBuildStepsStatusEnum(i interface{}) *BuildTriggerBuildStepsStatusEnum {
	s, ok := i.(string)
	if !ok {
		return BuildTriggerBuildStepsStatusEnumRef("")
	}

	return BuildTriggerBuildStepsStatusEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *BuildTrigger) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalBuildTrigger(b, c)
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

type buildTriggerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         buildTriggerApiOperation
}

func convertFieldDiffsToBuildTriggerDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]buildTriggerDiff, error) {
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
	var diffs []buildTriggerDiff
	// For each operation name, create a buildTriggerDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := buildTriggerDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToBuildTriggerApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToBuildTriggerApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (buildTriggerApiOperation, error) {
	switch opName {

	case "updateBuildTriggerUpdateBuildTriggerOperation":
		return &updateBuildTriggerUpdateBuildTriggerOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
