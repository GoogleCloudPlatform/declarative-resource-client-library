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
package cloudbuild

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
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateBuildTriggerUpdateBuildTriggerOperation) do(ctx context.Context, r *BuildTrigger, c *Client) error {
	_, err := c.GetBuildTrigger(ctx, r.urlNormalized())
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
		res := flattenBuildTrigger(c, v)
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

	_, err := c.GetBuildTrigger(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("BuildTrigger not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetBuildTrigger checking for existence. error: %v", err)
		return err
	}

	u, err := buildTriggerDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetBuildTrigger(ctx, r.urlNormalized())
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

	if _, err := c.GetBuildTrigger(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getBuildTriggerRaw(ctx context.Context, r *BuildTrigger) ([]byte, error) {

	u, err := buildTriggerGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) buildTriggerDiffsForRawDesired(ctx context.Context, rawDesired *BuildTrigger, opts ...dcl.ApplyOption) (initial, desired *BuildTrigger, diffs []buildTriggerDiff, err error) {
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
	rawInitial, err := c.GetBuildTrigger(ctx, fetchState.urlNormalized())
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

	if dcl.IsZeroValue(rawInitial.Filename) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Build) {
			rawInitial.Filename = dcl.String("")
		}
	}

	if dcl.IsZeroValue(rawInitial.Build) {
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

	if dcl.IsZeroValue(rawDesired.Filename) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Build) {
			rawDesired.Filename = dcl.String("")
		}
	}

	if dcl.IsZeroValue(rawDesired.Build) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Filename) {
			rawDesired.Build = EmptyBuildTriggerBuild
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.TriggerTemplate = canonicalizeBuildTriggerTriggerTemplate(rawDesired.TriggerTemplate, nil, opts...)
		rawDesired.Github = canonicalizeBuildTriggerGithub(rawDesired.Github, nil, opts...)
		rawDesired.Build = canonicalizeBuildTriggerBuild(rawDesired.Build, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Tags) {
		rawDesired.Tags = rawInitial.Tags
	}
	if dcl.BoolCanonicalize(rawDesired.Disabled, rawInitial.Disabled) {
		rawDesired.Disabled = rawInitial.Disabled
	}
	if dcl.IsZeroValue(rawDesired.Substitutions) {
		rawDesired.Substitutions = rawInitial.Substitutions
	}
	if dcl.StringCanonicalize(rawDesired.Filename, rawInitial.Filename) {
		rawDesired.Filename = rawInitial.Filename
	}
	if dcl.IsZeroValue(rawDesired.IgnoredFiles) {
		rawDesired.IgnoredFiles = rawInitial.IgnoredFiles
	}
	if dcl.IsZeroValue(rawDesired.IncludedFiles) {
		rawDesired.IncludedFiles = rawInitial.IncludedFiles
	}
	rawDesired.TriggerTemplate = canonicalizeBuildTriggerTriggerTemplate(rawDesired.TriggerTemplate, rawInitial.TriggerTemplate, opts...)
	rawDesired.Github = canonicalizeBuildTriggerGithub(rawDesired.Github, rawInitial.Github, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	rawDesired.Build = canonicalizeBuildTriggerBuild(rawDesired.Build, rawInitial.Build, opts...)
	if dcl.StringCanonicalize(rawDesired.Id, rawInitial.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.CreateTime, rawInitial.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}

	return rawDesired, nil
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

	if dcl.StringCanonicalize(des.ProjectId, initial.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		des.ProjectId = initial.ProjectId
	}
	if dcl.StringCanonicalize(des.RepoName, initial.RepoName) || dcl.IsZeroValue(des.RepoName) {
		des.RepoName = initial.RepoName
	}
	if dcl.StringCanonicalize(des.BranchName, initial.BranchName) || dcl.IsZeroValue(des.BranchName) {
		des.BranchName = initial.BranchName
	}
	if dcl.StringCanonicalize(des.TagName, initial.TagName) || dcl.IsZeroValue(des.TagName) {
		des.TagName = initial.TagName
	}
	if dcl.StringCanonicalize(des.CommitSha, initial.CommitSha) || dcl.IsZeroValue(des.CommitSha) {
		des.CommitSha = initial.CommitSha
	}
	if dcl.StringCanonicalize(des.Dir, initial.Dir) || dcl.IsZeroValue(des.Dir) {
		des.Dir = initial.Dir
	}
	if dcl.BoolCanonicalize(des.InvertRegex, initial.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
		des.InvertRegex = initial.InvertRegex
	}

	return des
}

func canonicalizeNewBuildTriggerTriggerTemplate(c *Client, des, nw *BuildTriggerTriggerTemplate) *BuildTriggerTriggerTemplate {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.RepoName) {
		nw.RepoName = dcl.String("default")
	}

	if dcl.StringCanonicalize(des.ProjectId, nw.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.RepoName, nw.RepoName) || dcl.IsZeroValue(des.RepoName) {
		nw.RepoName = des.RepoName
	}
	if dcl.StringCanonicalize(des.BranchName, nw.BranchName) || dcl.IsZeroValue(des.BranchName) {
		nw.BranchName = des.BranchName
	}
	if dcl.StringCanonicalize(des.TagName, nw.TagName) || dcl.IsZeroValue(des.TagName) {
		nw.TagName = des.TagName
	}
	if dcl.StringCanonicalize(des.CommitSha, nw.CommitSha) || dcl.IsZeroValue(des.CommitSha) {
		nw.CommitSha = des.CommitSha
	}
	if dcl.StringCanonicalize(des.Dir, nw.Dir) || dcl.IsZeroValue(des.Dir) {
		nw.Dir = des.Dir
	}
	if dcl.BoolCanonicalize(des.InvertRegex, nw.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
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
			if !compareBuildTriggerTriggerTemplate(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.Owner, initial.Owner) || dcl.IsZeroValue(des.Owner) {
		des.Owner = initial.Owner
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	des.PullRequest = canonicalizeBuildTriggerGithubPullRequest(des.PullRequest, initial.PullRequest, opts...)
	des.Push = canonicalizeBuildTriggerGithubPush(des.Push, initial.Push, opts...)

	return des
}

func canonicalizeNewBuildTriggerGithub(c *Client, des, nw *BuildTriggerGithub) *BuildTriggerGithub {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Owner, nw.Owner) || dcl.IsZeroValue(des.Owner) {
		nw.Owner = des.Owner
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
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
			if !compareBuildTriggerGithub(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.Branch, initial.Branch) || dcl.IsZeroValue(des.Branch) {
		des.Branch = initial.Branch
	}
	if dcl.IsZeroValue(des.CommentControl) {
		des.CommentControl = initial.CommentControl
	}
	if dcl.BoolCanonicalize(des.InvertRegex, initial.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
		des.InvertRegex = initial.InvertRegex
	}

	return des
}

func canonicalizeNewBuildTriggerGithubPullRequest(c *Client, des, nw *BuildTriggerGithubPullRequest) *BuildTriggerGithubPullRequest {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Branch, nw.Branch) || dcl.IsZeroValue(des.Branch) {
		nw.Branch = des.Branch
	}
	if dcl.BoolCanonicalize(des.InvertRegex, nw.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
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
			if !compareBuildTriggerGithubPullRequest(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.Branch, initial.Branch) || dcl.IsZeroValue(des.Branch) {
		des.Branch = initial.Branch
	}
	if dcl.StringCanonicalize(des.Tag, initial.Tag) || dcl.IsZeroValue(des.Tag) {
		des.Tag = initial.Tag
	}
	if dcl.BoolCanonicalize(des.InvertRegex, initial.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
		des.InvertRegex = initial.InvertRegex
	}

	return des
}

func canonicalizeNewBuildTriggerGithubPush(c *Client, des, nw *BuildTriggerGithubPush) *BuildTriggerGithubPush {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Branch, nw.Branch) || dcl.IsZeroValue(des.Branch) {
		nw.Branch = des.Branch
	}
	if dcl.StringCanonicalize(des.Tag, nw.Tag) || dcl.IsZeroValue(des.Tag) {
		nw.Tag = des.Tag
	}
	if dcl.BoolCanonicalize(des.InvertRegex, nw.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
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
			if !compareBuildTriggerGithubPush(c, &d, &n) {
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
		return des
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

	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	}
	if dcl.IsZeroValue(des.Images) {
		des.Images = initial.Images
	}
	if dcl.IsZeroValue(des.Substitutions) {
		des.Substitutions = initial.Substitutions
	}
	if dcl.StringCanonicalize(des.QueueTtl, initial.QueueTtl) || dcl.IsZeroValue(des.QueueTtl) {
		des.QueueTtl = initial.QueueTtl
	}
	if dcl.StringCanonicalize(des.LogsBucket, initial.LogsBucket) || dcl.IsZeroValue(des.LogsBucket) {
		des.LogsBucket = initial.LogsBucket
	}
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		des.Timeout = initial.Timeout
	}
	if dcl.IsZeroValue(des.Secrets) {
		des.Secrets = initial.Secrets
	}
	if dcl.IsZeroValue(des.Steps) {
		des.Steps = initial.Steps
	}
	des.Source = canonicalizeBuildTriggerBuildSource(des.Source, initial.Source, opts...)

	return des
}

func canonicalizeNewBuildTriggerBuild(c *Client, des, nw *BuildTriggerBuild) *BuildTriggerBuild {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Timeout) {
		nw.Timeout = dcl.String("600s")
	}

	if dcl.StringCanonicalize(des.QueueTtl, nw.QueueTtl) || dcl.IsZeroValue(des.QueueTtl) {
		nw.QueueTtl = des.QueueTtl
	}
	if dcl.StringCanonicalize(des.LogsBucket, nw.LogsBucket) || dcl.IsZeroValue(des.LogsBucket) {
		nw.LogsBucket = des.LogsBucket
	}
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) || dcl.IsZeroValue(des.Timeout) {
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
			if !compareBuildTriggerBuild(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}
	if dcl.IsZeroValue(des.SecretEnv) {
		des.SecretEnv = initial.SecretEnv
	}

	return des
}

func canonicalizeNewBuildTriggerBuildSecrets(c *Client, des, nw *BuildTriggerBuildSecrets) *BuildTriggerBuildSecrets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
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
			if !compareBuildTriggerBuildSecrets(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Env) {
		des.Env = initial.Env
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	}
	if dcl.StringCanonicalize(des.Dir, initial.Dir) || dcl.IsZeroValue(des.Dir) {
		des.Dir = initial.Dir
	}
	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		des.Id = initial.Id
	}
	if dcl.IsZeroValue(des.WaitFor) {
		des.WaitFor = initial.WaitFor
	}
	if dcl.StringCanonicalize(des.Entrypoint, initial.Entrypoint) || dcl.IsZeroValue(des.Entrypoint) {
		des.Entrypoint = initial.Entrypoint
	}
	if dcl.IsZeroValue(des.SecretEnv) {
		des.SecretEnv = initial.SecretEnv
	}
	if dcl.IsZeroValue(des.Volumes) {
		des.Volumes = initial.Volumes
	}
	des.Timing = canonicalizeBuildTriggerBuildStepsTiming(des.Timing, initial.Timing, opts...)
	des.PullTiming = canonicalizeBuildTriggerBuildStepsPullTiming(des.PullTiming, initial.PullTiming, opts...)
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		des.Timeout = initial.Timeout
	}
	if dcl.IsZeroValue(des.Status) {
		des.Status = initial.Status
	}

	return des
}

func canonicalizeNewBuildTriggerBuildSteps(c *Client, des, nw *BuildTriggerBuildSteps) *BuildTriggerBuildSteps {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Timeout) {
		nw.Timeout = dcl.String("300s")
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Dir, nw.Dir) || dcl.IsZeroValue(des.Dir) {
		nw.Dir = des.Dir
	}
	if dcl.StringCanonicalize(des.Id, nw.Id) || dcl.IsZeroValue(des.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Entrypoint, nw.Entrypoint) || dcl.IsZeroValue(des.Entrypoint) {
		nw.Entrypoint = des.Entrypoint
	}
	nw.Volumes = canonicalizeNewBuildTriggerBuildStepsVolumesSlice(c, des.Volumes, nw.Volumes)
	nw.Timing = canonicalizeNewBuildTriggerBuildStepsTiming(c, des.Timing, nw.Timing)
	nw.PullTiming = canonicalizeNewBuildTriggerBuildStepsPullTiming(c, des.PullTiming, nw.PullTiming)
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) || dcl.IsZeroValue(des.Timeout) {
		nw.Timeout = des.Timeout
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
			if !compareBuildTriggerBuildSteps(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}

	return des
}

func canonicalizeNewBuildTriggerBuildStepsVolumes(c *Client, des, nw *BuildTriggerBuildStepsVolumes) *BuildTriggerBuildStepsVolumes {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) || dcl.IsZeroValue(des.Path) {
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
			if !compareBuildTriggerBuildStepsVolumes(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.StartTime, initial.StartTime) || dcl.IsZeroValue(des.StartTime) {
		des.StartTime = initial.StartTime
	}
	if dcl.StringCanonicalize(des.EndTime, initial.EndTime) || dcl.IsZeroValue(des.EndTime) {
		des.EndTime = initial.EndTime
	}

	return des
}

func canonicalizeNewBuildTriggerBuildStepsTiming(c *Client, des, nw *BuildTriggerBuildStepsTiming) *BuildTriggerBuildStepsTiming {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.StartTime, nw.StartTime) || dcl.IsZeroValue(des.StartTime) {
		nw.StartTime = des.StartTime
	}
	if dcl.StringCanonicalize(des.EndTime, nw.EndTime) || dcl.IsZeroValue(des.EndTime) {
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
			if !compareBuildTriggerBuildStepsTiming(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.StartTime, initial.StartTime) || dcl.IsZeroValue(des.StartTime) {
		des.StartTime = initial.StartTime
	}
	if dcl.StringCanonicalize(des.EndTime, initial.EndTime) || dcl.IsZeroValue(des.EndTime) {
		des.EndTime = initial.EndTime
	}

	return des
}

func canonicalizeNewBuildTriggerBuildStepsPullTiming(c *Client, des, nw *BuildTriggerBuildStepsPullTiming) *BuildTriggerBuildStepsPullTiming {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.StartTime, nw.StartTime) || dcl.IsZeroValue(des.StartTime) {
		nw.StartTime = des.StartTime
	}
	if dcl.StringCanonicalize(des.EndTime, nw.EndTime) || dcl.IsZeroValue(des.EndTime) {
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
			if !compareBuildTriggerBuildStepsPullTiming(c, &d, &n) {
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
		return des
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

	des.StorageSource = canonicalizeBuildTriggerBuildSourceStorageSource(des.StorageSource, initial.StorageSource, opts...)
	des.RepoSource = canonicalizeBuildTriggerBuildSourceRepoSource(des.RepoSource, initial.RepoSource, opts...)

	return des
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
			if !compareBuildTriggerBuildSource(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.Bucket, initial.Bucket) || dcl.IsZeroValue(des.Bucket) {
		des.Bucket = initial.Bucket
	}
	if dcl.StringCanonicalize(des.Object, initial.Object) || dcl.IsZeroValue(des.Object) {
		des.Object = initial.Object
	}
	if dcl.StringCanonicalize(des.Generation, initial.Generation) || dcl.IsZeroValue(des.Generation) {
		des.Generation = initial.Generation
	}

	return des
}

func canonicalizeNewBuildTriggerBuildSourceStorageSource(c *Client, des, nw *BuildTriggerBuildSourceStorageSource) *BuildTriggerBuildSourceStorageSource {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Bucket, nw.Bucket) || dcl.IsZeroValue(des.Bucket) {
		nw.Bucket = des.Bucket
	}
	if dcl.StringCanonicalize(des.Object, nw.Object) || dcl.IsZeroValue(des.Object) {
		nw.Object = des.Object
	}
	if dcl.StringCanonicalize(des.Generation, nw.Generation) || dcl.IsZeroValue(des.Generation) {
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
			if !compareBuildTriggerBuildSourceStorageSource(c, &d, &n) {
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
		return des
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

	if dcl.StringCanonicalize(des.ProjectId, initial.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		des.ProjectId = initial.ProjectId
	}
	if dcl.StringCanonicalize(des.RepoName, initial.RepoName) || dcl.IsZeroValue(des.RepoName) {
		des.RepoName = initial.RepoName
	}
	if dcl.StringCanonicalize(des.BranchName, initial.BranchName) || dcl.IsZeroValue(des.BranchName) {
		des.BranchName = initial.BranchName
	}
	if dcl.StringCanonicalize(des.TagName, initial.TagName) || dcl.IsZeroValue(des.TagName) {
		des.TagName = initial.TagName
	}
	if dcl.StringCanonicalize(des.CommitSha, initial.CommitSha) || dcl.IsZeroValue(des.CommitSha) {
		des.CommitSha = initial.CommitSha
	}
	if dcl.StringCanonicalize(des.Dir, initial.Dir) || dcl.IsZeroValue(des.Dir) {
		des.Dir = initial.Dir
	}
	if dcl.BoolCanonicalize(des.InvertRegex, initial.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
		des.InvertRegex = initial.InvertRegex
	}
	if dcl.IsZeroValue(des.Substitutions) {
		des.Substitutions = initial.Substitutions
	}

	return des
}

func canonicalizeNewBuildTriggerBuildSourceRepoSource(c *Client, des, nw *BuildTriggerBuildSourceRepoSource) *BuildTriggerBuildSourceRepoSource {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ProjectId, nw.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.RepoName, nw.RepoName) || dcl.IsZeroValue(des.RepoName) {
		nw.RepoName = des.RepoName
	}
	if dcl.StringCanonicalize(des.BranchName, nw.BranchName) || dcl.IsZeroValue(des.BranchName) {
		nw.BranchName = des.BranchName
	}
	if dcl.StringCanonicalize(des.TagName, nw.TagName) || dcl.IsZeroValue(des.TagName) {
		nw.TagName = des.TagName
	}
	if dcl.StringCanonicalize(des.CommitSha, nw.CommitSha) || dcl.IsZeroValue(des.CommitSha) {
		nw.CommitSha = des.CommitSha
	}
	if dcl.StringCanonicalize(des.Dir, nw.Dir) || dcl.IsZeroValue(des.Dir) {
		nw.Dir = des.Dir
	}
	if dcl.BoolCanonicalize(des.InvertRegex, nw.InvertRegex) || dcl.IsZeroValue(des.InvertRegex) {
		nw.InvertRegex = des.InvertRegex
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
			if !compareBuildTriggerBuildSourceRepoSource(c, &d, &n) {
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
		return des
	}

	var items []BuildTriggerBuildSourceRepoSource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBuildTriggerBuildSourceRepoSource(c, &d, &n))
	}

	return items
}

type buildTriggerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         buildTriggerApiOperation
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
func diffBuildTrigger(c *Client, desired, actual *BuildTrigger, opts ...dcl.ApplyOption) ([]buildTriggerDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []buildTriggerDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "Name",
		})

	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "Description",
		})

	}
	if !dcl.StringSliceEquals(desired.Tags, actual.Tags) {
		c.Config.Logger.Infof("Detected diff in Tags.\nDESIRED: %v\nACTUAL: %v", desired.Tags, actual.Tags)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "Tags",
		})

	}
	if !dcl.IsZeroValue(desired.Disabled) && !dcl.BoolCanonicalize(desired.Disabled, actual.Disabled) {
		c.Config.Logger.Infof("Detected diff in Disabled.\nDESIRED: %v\nACTUAL: %v", desired.Disabled, actual.Disabled)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "Disabled",
		})

	}
	if !dcl.MapEquals(desired.Substitutions, actual.Substitutions, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Substitutions.\nDESIRED: %v\nACTUAL: %v", desired.Substitutions, actual.Substitutions)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "Substitutions",
		})

	}
	if !dcl.IsZeroValue(desired.Filename) && !dcl.StringCanonicalize(desired.Filename, actual.Filename) {
		c.Config.Logger.Infof("Detected diff in Filename.\nDESIRED: %v\nACTUAL: %v", desired.Filename, actual.Filename)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "Filename",
		})

	}
	if !dcl.StringSliceEquals(desired.IgnoredFiles, actual.IgnoredFiles) {
		c.Config.Logger.Infof("Detected diff in IgnoredFiles.\nDESIRED: %v\nACTUAL: %v", desired.IgnoredFiles, actual.IgnoredFiles)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "IgnoredFiles",
		})

	}
	if !dcl.StringSliceEquals(desired.IncludedFiles, actual.IncludedFiles) {
		c.Config.Logger.Infof("Detected diff in IncludedFiles.\nDESIRED: %v\nACTUAL: %v", desired.IncludedFiles, actual.IncludedFiles)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "IncludedFiles",
		})

	}
	if compareBuildTriggerTriggerTemplate(c, desired.TriggerTemplate, actual.TriggerTemplate) {
		c.Config.Logger.Infof("Detected diff in TriggerTemplate.\nDESIRED: %v\nACTUAL: %v", desired.TriggerTemplate, actual.TriggerTemplate)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "TriggerTemplate",
		})

	}
	if compareBuildTriggerGithub(c, desired.Github, actual.Github) {
		c.Config.Logger.Infof("Detected diff in Github.\nDESIRED: %v\nACTUAL: %v", desired.Github, actual.Github)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "Github",
		})

	}
	if compareBuildTriggerBuild(c, desired.Build, actual.Build) {
		c.Config.Logger.Infof("Detected diff in Build.\nDESIRED: %v\nACTUAL: %v", desired.Build, actual.Build)

		diffs = append(diffs, buildTriggerDiff{
			UpdateOp:  &updateBuildTriggerUpdateBuildTriggerOperation{},
			FieldName: "Build",
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
	var deduped []buildTriggerDiff
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
func compareBuildTriggerTriggerTemplate(c *Client, desired, actual *BuildTriggerTriggerTemplate) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ProjectId == nil && desired.ProjectId != nil && !dcl.IsEmptyValueIndirect(desired.ProjectId) {
		c.Config.Logger.Infof("desired ProjectId %s - but actually nil", dcl.SprintResource(desired.ProjectId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ProjectId, actual.ProjectId) && !dcl.IsZeroValue(desired.ProjectId) {
		c.Config.Logger.Infof("Diff in ProjectId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ProjectId), dcl.SprintResource(actual.ProjectId))
		return true
	}
	if actual.RepoName == nil && desired.RepoName != nil && !dcl.IsEmptyValueIndirect(desired.RepoName) {
		c.Config.Logger.Infof("desired RepoName %s - but actually nil", dcl.SprintResource(desired.RepoName))
		return true
	}
	if !dcl.StringCanonicalize(desired.RepoName, actual.RepoName) && !dcl.IsZeroValue(desired.RepoName) {
		c.Config.Logger.Infof("Diff in RepoName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RepoName), dcl.SprintResource(actual.RepoName))
		return true
	}
	if actual.BranchName == nil && desired.BranchName != nil && !dcl.IsEmptyValueIndirect(desired.BranchName) {
		c.Config.Logger.Infof("desired BranchName %s - but actually nil", dcl.SprintResource(desired.BranchName))
		return true
	}
	if !dcl.StringCanonicalize(desired.BranchName, actual.BranchName) && !dcl.IsZeroValue(desired.BranchName) {
		c.Config.Logger.Infof("Diff in BranchName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BranchName), dcl.SprintResource(actual.BranchName))
		return true
	}
	if actual.TagName == nil && desired.TagName != nil && !dcl.IsEmptyValueIndirect(desired.TagName) {
		c.Config.Logger.Infof("desired TagName %s - but actually nil", dcl.SprintResource(desired.TagName))
		return true
	}
	if !dcl.StringCanonicalize(desired.TagName, actual.TagName) && !dcl.IsZeroValue(desired.TagName) {
		c.Config.Logger.Infof("Diff in TagName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TagName), dcl.SprintResource(actual.TagName))
		return true
	}
	if actual.CommitSha == nil && desired.CommitSha != nil && !dcl.IsEmptyValueIndirect(desired.CommitSha) {
		c.Config.Logger.Infof("desired CommitSha %s - but actually nil", dcl.SprintResource(desired.CommitSha))
		return true
	}
	if !dcl.StringCanonicalize(desired.CommitSha, actual.CommitSha) && !dcl.IsZeroValue(desired.CommitSha) {
		c.Config.Logger.Infof("Diff in CommitSha. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CommitSha), dcl.SprintResource(actual.CommitSha))
		return true
	}
	if actual.Dir == nil && desired.Dir != nil && !dcl.IsEmptyValueIndirect(desired.Dir) {
		c.Config.Logger.Infof("desired Dir %s - but actually nil", dcl.SprintResource(desired.Dir))
		return true
	}
	if !dcl.StringCanonicalize(desired.Dir, actual.Dir) && !dcl.IsZeroValue(desired.Dir) {
		c.Config.Logger.Infof("Diff in Dir. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Dir), dcl.SprintResource(actual.Dir))
		return true
	}
	if actual.InvertRegex == nil && desired.InvertRegex != nil && !dcl.IsEmptyValueIndirect(desired.InvertRegex) {
		c.Config.Logger.Infof("desired InvertRegex %s - but actually nil", dcl.SprintResource(desired.InvertRegex))
		return true
	}
	if !dcl.BoolCanonicalize(desired.InvertRegex, actual.InvertRegex) && !dcl.IsZeroValue(desired.InvertRegex) {
		c.Config.Logger.Infof("Diff in InvertRegex. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InvertRegex), dcl.SprintResource(actual.InvertRegex))
		return true
	}
	return false
}

func compareBuildTriggerTriggerTemplateSlice(c *Client, desired, actual []BuildTriggerTriggerTemplate) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerTriggerTemplate, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerTriggerTemplate(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerTriggerTemplate, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerTriggerTemplateMap(c *Client, desired, actual map[string]BuildTriggerTriggerTemplate) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerTriggerTemplate, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerTriggerTemplate, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerTriggerTemplate(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerTriggerTemplate, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerGithub(c *Client, desired, actual *BuildTriggerGithub) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Owner == nil && desired.Owner != nil && !dcl.IsEmptyValueIndirect(desired.Owner) {
		c.Config.Logger.Infof("desired Owner %s - but actually nil", dcl.SprintResource(desired.Owner))
		return true
	}
	if !dcl.StringCanonicalize(desired.Owner, actual.Owner) && !dcl.IsZeroValue(desired.Owner) {
		c.Config.Logger.Infof("Diff in Owner. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Owner), dcl.SprintResource(actual.Owner))
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
	if actual.PullRequest == nil && desired.PullRequest != nil && !dcl.IsEmptyValueIndirect(desired.PullRequest) {
		c.Config.Logger.Infof("desired PullRequest %s - but actually nil", dcl.SprintResource(desired.PullRequest))
		return true
	}
	if compareBuildTriggerGithubPullRequest(c, desired.PullRequest, actual.PullRequest) && !dcl.IsZeroValue(desired.PullRequest) {
		c.Config.Logger.Infof("Diff in PullRequest. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PullRequest), dcl.SprintResource(actual.PullRequest))
		return true
	}
	if actual.Push == nil && desired.Push != nil && !dcl.IsEmptyValueIndirect(desired.Push) {
		c.Config.Logger.Infof("desired Push %s - but actually nil", dcl.SprintResource(desired.Push))
		return true
	}
	if compareBuildTriggerGithubPush(c, desired.Push, actual.Push) && !dcl.IsZeroValue(desired.Push) {
		c.Config.Logger.Infof("Diff in Push. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Push), dcl.SprintResource(actual.Push))
		return true
	}
	return false
}

func compareBuildTriggerGithubSlice(c *Client, desired, actual []BuildTriggerGithub) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerGithub, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerGithub(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerGithub, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerGithubMap(c *Client, desired, actual map[string]BuildTriggerGithub) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerGithub, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerGithub, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerGithub(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerGithub, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerGithubPullRequest(c *Client, desired, actual *BuildTriggerGithubPullRequest) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Branch == nil && desired.Branch != nil && !dcl.IsEmptyValueIndirect(desired.Branch) {
		c.Config.Logger.Infof("desired Branch %s - but actually nil", dcl.SprintResource(desired.Branch))
		return true
	}
	if !dcl.StringCanonicalize(desired.Branch, actual.Branch) && !dcl.IsZeroValue(desired.Branch) {
		c.Config.Logger.Infof("Diff in Branch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Branch), dcl.SprintResource(actual.Branch))
		return true
	}
	if actual.CommentControl == nil && desired.CommentControl != nil && !dcl.IsEmptyValueIndirect(desired.CommentControl) {
		c.Config.Logger.Infof("desired CommentControl %s - but actually nil", dcl.SprintResource(desired.CommentControl))
		return true
	}
	if !reflect.DeepEqual(desired.CommentControl, actual.CommentControl) && !dcl.IsZeroValue(desired.CommentControl) {
		c.Config.Logger.Infof("Diff in CommentControl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CommentControl), dcl.SprintResource(actual.CommentControl))
		return true
	}
	if actual.InvertRegex == nil && desired.InvertRegex != nil && !dcl.IsEmptyValueIndirect(desired.InvertRegex) {
		c.Config.Logger.Infof("desired InvertRegex %s - but actually nil", dcl.SprintResource(desired.InvertRegex))
		return true
	}
	if !dcl.BoolCanonicalize(desired.InvertRegex, actual.InvertRegex) && !dcl.IsZeroValue(desired.InvertRegex) {
		c.Config.Logger.Infof("Diff in InvertRegex. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InvertRegex), dcl.SprintResource(actual.InvertRegex))
		return true
	}
	return false
}

func compareBuildTriggerGithubPullRequestSlice(c *Client, desired, actual []BuildTriggerGithubPullRequest) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerGithubPullRequest, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerGithubPullRequest(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerGithubPullRequest, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerGithubPullRequestMap(c *Client, desired, actual map[string]BuildTriggerGithubPullRequest) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerGithubPullRequest, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerGithubPullRequest, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerGithubPullRequest(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerGithubPullRequest, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerGithubPush(c *Client, desired, actual *BuildTriggerGithubPush) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Branch == nil && desired.Branch != nil && !dcl.IsEmptyValueIndirect(desired.Branch) {
		c.Config.Logger.Infof("desired Branch %s - but actually nil", dcl.SprintResource(desired.Branch))
		return true
	}
	if !dcl.StringCanonicalize(desired.Branch, actual.Branch) && !dcl.IsZeroValue(desired.Branch) {
		c.Config.Logger.Infof("Diff in Branch. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Branch), dcl.SprintResource(actual.Branch))
		return true
	}
	if actual.Tag == nil && desired.Tag != nil && !dcl.IsEmptyValueIndirect(desired.Tag) {
		c.Config.Logger.Infof("desired Tag %s - but actually nil", dcl.SprintResource(desired.Tag))
		return true
	}
	if !dcl.StringCanonicalize(desired.Tag, actual.Tag) && !dcl.IsZeroValue(desired.Tag) {
		c.Config.Logger.Infof("Diff in Tag. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tag), dcl.SprintResource(actual.Tag))
		return true
	}
	if actual.InvertRegex == nil && desired.InvertRegex != nil && !dcl.IsEmptyValueIndirect(desired.InvertRegex) {
		c.Config.Logger.Infof("desired InvertRegex %s - but actually nil", dcl.SprintResource(desired.InvertRegex))
		return true
	}
	if !dcl.BoolCanonicalize(desired.InvertRegex, actual.InvertRegex) && !dcl.IsZeroValue(desired.InvertRegex) {
		c.Config.Logger.Infof("Diff in InvertRegex. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InvertRegex), dcl.SprintResource(actual.InvertRegex))
		return true
	}
	return false
}

func compareBuildTriggerGithubPushSlice(c *Client, desired, actual []BuildTriggerGithubPush) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerGithubPush, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerGithubPush(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerGithubPush, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerGithubPushMap(c *Client, desired, actual map[string]BuildTriggerGithubPush) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerGithubPush, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerGithubPush, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerGithubPush(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerGithubPush, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuild(c *Client, desired, actual *BuildTriggerBuild) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Tags == nil && desired.Tags != nil && !dcl.IsEmptyValueIndirect(desired.Tags) {
		c.Config.Logger.Infof("desired Tags %s - but actually nil", dcl.SprintResource(desired.Tags))
		return true
	}
	if !dcl.StringSliceEquals(desired.Tags, actual.Tags) && !dcl.IsZeroValue(desired.Tags) {
		c.Config.Logger.Infof("Diff in Tags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tags), dcl.SprintResource(actual.Tags))
		return true
	}
	if actual.Images == nil && desired.Images != nil && !dcl.IsEmptyValueIndirect(desired.Images) {
		c.Config.Logger.Infof("desired Images %s - but actually nil", dcl.SprintResource(desired.Images))
		return true
	}
	if !dcl.StringSliceEquals(desired.Images, actual.Images) && !dcl.IsZeroValue(desired.Images) {
		c.Config.Logger.Infof("Diff in Images. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Images), dcl.SprintResource(actual.Images))
		return true
	}
	if actual.Substitutions == nil && desired.Substitutions != nil && !dcl.IsEmptyValueIndirect(desired.Substitutions) {
		c.Config.Logger.Infof("desired Substitutions %s - but actually nil", dcl.SprintResource(desired.Substitutions))
		return true
	}
	if !dcl.MapEquals(desired.Substitutions, actual.Substitutions, []string(nil)) && !dcl.IsZeroValue(desired.Substitutions) {
		c.Config.Logger.Infof("Diff in Substitutions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Substitutions), dcl.SprintResource(actual.Substitutions))
		return true
	}
	if actual.QueueTtl == nil && desired.QueueTtl != nil && !dcl.IsEmptyValueIndirect(desired.QueueTtl) {
		c.Config.Logger.Infof("desired QueueTtl %s - but actually nil", dcl.SprintResource(desired.QueueTtl))
		return true
	}
	if !dcl.StringCanonicalize(desired.QueueTtl, actual.QueueTtl) && !dcl.IsZeroValue(desired.QueueTtl) {
		c.Config.Logger.Infof("Diff in QueueTtl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueueTtl), dcl.SprintResource(actual.QueueTtl))
		return true
	}
	if actual.LogsBucket == nil && desired.LogsBucket != nil && !dcl.IsEmptyValueIndirect(desired.LogsBucket) {
		c.Config.Logger.Infof("desired LogsBucket %s - but actually nil", dcl.SprintResource(desired.LogsBucket))
		return true
	}
	if !dcl.StringCanonicalize(desired.LogsBucket, actual.LogsBucket) && !dcl.IsZeroValue(desired.LogsBucket) {
		c.Config.Logger.Infof("Diff in LogsBucket. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LogsBucket), dcl.SprintResource(actual.LogsBucket))
		return true
	}
	if actual.Timeout == nil && desired.Timeout != nil && !dcl.IsEmptyValueIndirect(desired.Timeout) {
		c.Config.Logger.Infof("desired Timeout %s - but actually nil", dcl.SprintResource(desired.Timeout))
		return true
	}
	if !dcl.StringCanonicalize(desired.Timeout, actual.Timeout) && !dcl.IsZeroValue(desired.Timeout) {
		c.Config.Logger.Infof("Diff in Timeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Timeout), dcl.SprintResource(actual.Timeout))
		return true
	}
	if actual.Secrets == nil && desired.Secrets != nil && !dcl.IsEmptyValueIndirect(desired.Secrets) {
		c.Config.Logger.Infof("desired Secrets %s - but actually nil", dcl.SprintResource(desired.Secrets))
		return true
	}
	if compareBuildTriggerBuildSecretsSlice(c, desired.Secrets, actual.Secrets) && !dcl.IsZeroValue(desired.Secrets) {
		c.Config.Logger.Infof("Diff in Secrets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Secrets), dcl.SprintResource(actual.Secrets))
		return true
	}
	if actual.Steps == nil && desired.Steps != nil && !dcl.IsEmptyValueIndirect(desired.Steps) {
		c.Config.Logger.Infof("desired Steps %s - but actually nil", dcl.SprintResource(desired.Steps))
		return true
	}
	if compareBuildTriggerBuildStepsSlice(c, desired.Steps, actual.Steps) && !dcl.IsZeroValue(desired.Steps) {
		c.Config.Logger.Infof("Diff in Steps. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Steps), dcl.SprintResource(actual.Steps))
		return true
	}
	if actual.Source == nil && desired.Source != nil && !dcl.IsEmptyValueIndirect(desired.Source) {
		c.Config.Logger.Infof("desired Source %s - but actually nil", dcl.SprintResource(desired.Source))
		return true
	}
	if compareBuildTriggerBuildSource(c, desired.Source, actual.Source) && !dcl.IsZeroValue(desired.Source) {
		c.Config.Logger.Infof("Diff in Source. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Source), dcl.SprintResource(actual.Source))
		return true
	}
	return false
}

func compareBuildTriggerBuildSlice(c *Client, desired, actual []BuildTriggerBuild) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuild, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuild(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuild, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildMap(c *Client, desired, actual map[string]BuildTriggerBuild) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuild, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerBuild, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerBuild(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuild, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildSecrets(c *Client, desired, actual *BuildTriggerBuildSecrets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.KmsKeyName == nil && desired.KmsKeyName != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyName) {
		c.Config.Logger.Infof("desired KmsKeyName %s - but actually nil", dcl.SprintResource(desired.KmsKeyName))
		return true
	}
	if !dcl.StringCanonicalize(desired.KmsKeyName, actual.KmsKeyName) && !dcl.IsZeroValue(desired.KmsKeyName) {
		c.Config.Logger.Infof("Diff in KmsKeyName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyName), dcl.SprintResource(actual.KmsKeyName))
		return true
	}
	if actual.SecretEnv == nil && desired.SecretEnv != nil && !dcl.IsEmptyValueIndirect(desired.SecretEnv) {
		c.Config.Logger.Infof("desired SecretEnv %s - but actually nil", dcl.SprintResource(desired.SecretEnv))
		return true
	}
	if !dcl.MapEquals(desired.SecretEnv, actual.SecretEnv, []string(nil)) && !dcl.IsZeroValue(desired.SecretEnv) {
		c.Config.Logger.Infof("Diff in SecretEnv. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecretEnv), dcl.SprintResource(actual.SecretEnv))
		return true
	}
	return false
}

func compareBuildTriggerBuildSecretsSlice(c *Client, desired, actual []BuildTriggerBuildSecrets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSecrets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuildSecrets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSecrets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildSecretsMap(c *Client, desired, actual map[string]BuildTriggerBuildSecrets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSecrets, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSecrets, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerBuildSecrets(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSecrets, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildSteps(c *Client, desired, actual *BuildTriggerBuildSteps) bool {
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
	if actual.Env == nil && desired.Env != nil && !dcl.IsEmptyValueIndirect(desired.Env) {
		c.Config.Logger.Infof("desired Env %s - but actually nil", dcl.SprintResource(desired.Env))
		return true
	}
	if !dcl.StringSliceEquals(desired.Env, actual.Env) && !dcl.IsZeroValue(desired.Env) {
		c.Config.Logger.Infof("Diff in Env. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Env), dcl.SprintResource(actual.Env))
		return true
	}
	if actual.Args == nil && desired.Args != nil && !dcl.IsEmptyValueIndirect(desired.Args) {
		c.Config.Logger.Infof("desired Args %s - but actually nil", dcl.SprintResource(desired.Args))
		return true
	}
	if !dcl.StringSliceEquals(desired.Args, actual.Args) && !dcl.IsZeroValue(desired.Args) {
		c.Config.Logger.Infof("Diff in Args. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Args), dcl.SprintResource(actual.Args))
		return true
	}
	if actual.Dir == nil && desired.Dir != nil && !dcl.IsEmptyValueIndirect(desired.Dir) {
		c.Config.Logger.Infof("desired Dir %s - but actually nil", dcl.SprintResource(desired.Dir))
		return true
	}
	if !dcl.StringCanonicalize(desired.Dir, actual.Dir) && !dcl.IsZeroValue(desired.Dir) {
		c.Config.Logger.Infof("Diff in Dir. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Dir), dcl.SprintResource(actual.Dir))
		return true
	}
	if actual.Id == nil && desired.Id != nil && !dcl.IsEmptyValueIndirect(desired.Id) {
		c.Config.Logger.Infof("desired Id %s - but actually nil", dcl.SprintResource(desired.Id))
		return true
	}
	if !dcl.StringCanonicalize(desired.Id, actual.Id) && !dcl.IsZeroValue(desired.Id) {
		c.Config.Logger.Infof("Diff in Id. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Id), dcl.SprintResource(actual.Id))
		return true
	}
	if actual.WaitFor == nil && desired.WaitFor != nil && !dcl.IsEmptyValueIndirect(desired.WaitFor) {
		c.Config.Logger.Infof("desired WaitFor %s - but actually nil", dcl.SprintResource(desired.WaitFor))
		return true
	}
	if !dcl.StringSliceEquals(desired.WaitFor, actual.WaitFor) && !dcl.IsZeroValue(desired.WaitFor) {
		c.Config.Logger.Infof("Diff in WaitFor. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WaitFor), dcl.SprintResource(actual.WaitFor))
		return true
	}
	if actual.Entrypoint == nil && desired.Entrypoint != nil && !dcl.IsEmptyValueIndirect(desired.Entrypoint) {
		c.Config.Logger.Infof("desired Entrypoint %s - but actually nil", dcl.SprintResource(desired.Entrypoint))
		return true
	}
	if !dcl.StringCanonicalize(desired.Entrypoint, actual.Entrypoint) && !dcl.IsZeroValue(desired.Entrypoint) {
		c.Config.Logger.Infof("Diff in Entrypoint. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Entrypoint), dcl.SprintResource(actual.Entrypoint))
		return true
	}
	if actual.SecretEnv == nil && desired.SecretEnv != nil && !dcl.IsEmptyValueIndirect(desired.SecretEnv) {
		c.Config.Logger.Infof("desired SecretEnv %s - but actually nil", dcl.SprintResource(desired.SecretEnv))
		return true
	}
	if !dcl.StringSliceEquals(desired.SecretEnv, actual.SecretEnv) && !dcl.IsZeroValue(desired.SecretEnv) {
		c.Config.Logger.Infof("Diff in SecretEnv. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecretEnv), dcl.SprintResource(actual.SecretEnv))
		return true
	}
	if actual.Volumes == nil && desired.Volumes != nil && !dcl.IsEmptyValueIndirect(desired.Volumes) {
		c.Config.Logger.Infof("desired Volumes %s - but actually nil", dcl.SprintResource(desired.Volumes))
		return true
	}
	if compareBuildTriggerBuildStepsVolumesSlice(c, desired.Volumes, actual.Volumes) && !dcl.IsZeroValue(desired.Volumes) {
		c.Config.Logger.Infof("Diff in Volumes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Volumes), dcl.SprintResource(actual.Volumes))
		return true
	}
	if actual.Timeout == nil && desired.Timeout != nil && !dcl.IsEmptyValueIndirect(desired.Timeout) {
		c.Config.Logger.Infof("desired Timeout %s - but actually nil", dcl.SprintResource(desired.Timeout))
		return true
	}
	if !dcl.StringCanonicalize(desired.Timeout, actual.Timeout) && !dcl.IsZeroValue(desired.Timeout) {
		c.Config.Logger.Infof("Diff in Timeout. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Timeout), dcl.SprintResource(actual.Timeout))
		return true
	}
	return false
}

func compareBuildTriggerBuildStepsSlice(c *Client, desired, actual []BuildTriggerBuildSteps) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSteps, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuildSteps(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSteps, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildStepsMap(c *Client, desired, actual map[string]BuildTriggerBuildSteps) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSteps, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSteps, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerBuildSteps(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSteps, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildStepsVolumes(c *Client, desired, actual *BuildTriggerBuildStepsVolumes) bool {
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
	if actual.Path == nil && desired.Path != nil && !dcl.IsEmptyValueIndirect(desired.Path) {
		c.Config.Logger.Infof("desired Path %s - but actually nil", dcl.SprintResource(desired.Path))
		return true
	}
	if !dcl.StringCanonicalize(desired.Path, actual.Path) && !dcl.IsZeroValue(desired.Path) {
		c.Config.Logger.Infof("Diff in Path. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Path), dcl.SprintResource(actual.Path))
		return true
	}
	return false
}

func compareBuildTriggerBuildStepsVolumesSlice(c *Client, desired, actual []BuildTriggerBuildStepsVolumes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildStepsVolumes, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuildStepsVolumes(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsVolumes, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildStepsVolumesMap(c *Client, desired, actual map[string]BuildTriggerBuildStepsVolumes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildStepsVolumes, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsVolumes, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerBuildStepsVolumes(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsVolumes, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildStepsTiming(c *Client, desired, actual *BuildTriggerBuildStepsTiming) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.StartTime == nil && desired.StartTime != nil && !dcl.IsEmptyValueIndirect(desired.StartTime) {
		c.Config.Logger.Infof("desired StartTime %s - but actually nil", dcl.SprintResource(desired.StartTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) {
		c.Config.Logger.Infof("Diff in StartTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if actual.EndTime == nil && desired.EndTime != nil && !dcl.IsEmptyValueIndirect(desired.EndTime) {
		c.Config.Logger.Infof("desired EndTime %s - but actually nil", dcl.SprintResource(desired.EndTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.EndTime, actual.EndTime) && !dcl.IsZeroValue(desired.EndTime) {
		c.Config.Logger.Infof("Diff in EndTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EndTime), dcl.SprintResource(actual.EndTime))
		return true
	}
	return false
}

func compareBuildTriggerBuildStepsTimingSlice(c *Client, desired, actual []BuildTriggerBuildStepsTiming) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildStepsTiming, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuildStepsTiming(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsTiming, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildStepsTimingMap(c *Client, desired, actual map[string]BuildTriggerBuildStepsTiming) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildStepsTiming, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsTiming, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerBuildStepsTiming(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsTiming, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildStepsPullTiming(c *Client, desired, actual *BuildTriggerBuildStepsPullTiming) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.StartTime == nil && desired.StartTime != nil && !dcl.IsEmptyValueIndirect(desired.StartTime) {
		c.Config.Logger.Infof("desired StartTime %s - but actually nil", dcl.SprintResource(desired.StartTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) {
		c.Config.Logger.Infof("Diff in StartTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if actual.EndTime == nil && desired.EndTime != nil && !dcl.IsEmptyValueIndirect(desired.EndTime) {
		c.Config.Logger.Infof("desired EndTime %s - but actually nil", dcl.SprintResource(desired.EndTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.EndTime, actual.EndTime) && !dcl.IsZeroValue(desired.EndTime) {
		c.Config.Logger.Infof("Diff in EndTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EndTime), dcl.SprintResource(actual.EndTime))
		return true
	}
	return false
}

func compareBuildTriggerBuildStepsPullTimingSlice(c *Client, desired, actual []BuildTriggerBuildStepsPullTiming) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildStepsPullTiming, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuildStepsPullTiming(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsPullTiming, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildStepsPullTimingMap(c *Client, desired, actual map[string]BuildTriggerBuildStepsPullTiming) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildStepsPullTiming, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsPullTiming, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerBuildStepsPullTiming(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsPullTiming, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildSource(c *Client, desired, actual *BuildTriggerBuildSource) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.StorageSource == nil && desired.StorageSource != nil && !dcl.IsEmptyValueIndirect(desired.StorageSource) {
		c.Config.Logger.Infof("desired StorageSource %s - but actually nil", dcl.SprintResource(desired.StorageSource))
		return true
	}
	if compareBuildTriggerBuildSourceStorageSource(c, desired.StorageSource, actual.StorageSource) && !dcl.IsZeroValue(desired.StorageSource) {
		c.Config.Logger.Infof("Diff in StorageSource. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StorageSource), dcl.SprintResource(actual.StorageSource))
		return true
	}
	if actual.RepoSource == nil && desired.RepoSource != nil && !dcl.IsEmptyValueIndirect(desired.RepoSource) {
		c.Config.Logger.Infof("desired RepoSource %s - but actually nil", dcl.SprintResource(desired.RepoSource))
		return true
	}
	if compareBuildTriggerBuildSourceRepoSource(c, desired.RepoSource, actual.RepoSource) && !dcl.IsZeroValue(desired.RepoSource) {
		c.Config.Logger.Infof("Diff in RepoSource. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RepoSource), dcl.SprintResource(actual.RepoSource))
		return true
	}
	return false
}

func compareBuildTriggerBuildSourceSlice(c *Client, desired, actual []BuildTriggerBuildSource) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSource, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuildSource(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSource, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildSourceMap(c *Client, desired, actual map[string]BuildTriggerBuildSource) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSource, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSource, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerBuildSource(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSource, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildSourceStorageSource(c *Client, desired, actual *BuildTriggerBuildSourceStorageSource) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Bucket == nil && desired.Bucket != nil && !dcl.IsEmptyValueIndirect(desired.Bucket) {
		c.Config.Logger.Infof("desired Bucket %s - but actually nil", dcl.SprintResource(desired.Bucket))
		return true
	}
	if !dcl.StringCanonicalize(desired.Bucket, actual.Bucket) && !dcl.IsZeroValue(desired.Bucket) {
		c.Config.Logger.Infof("Diff in Bucket. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Bucket), dcl.SprintResource(actual.Bucket))
		return true
	}
	if actual.Object == nil && desired.Object != nil && !dcl.IsEmptyValueIndirect(desired.Object) {
		c.Config.Logger.Infof("desired Object %s - but actually nil", dcl.SprintResource(desired.Object))
		return true
	}
	if !dcl.StringCanonicalize(desired.Object, actual.Object) && !dcl.IsZeroValue(desired.Object) {
		c.Config.Logger.Infof("Diff in Object. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Object), dcl.SprintResource(actual.Object))
		return true
	}
	if actual.Generation == nil && desired.Generation != nil && !dcl.IsEmptyValueIndirect(desired.Generation) {
		c.Config.Logger.Infof("desired Generation %s - but actually nil", dcl.SprintResource(desired.Generation))
		return true
	}
	if !dcl.StringCanonicalize(desired.Generation, actual.Generation) && !dcl.IsZeroValue(desired.Generation) {
		c.Config.Logger.Infof("Diff in Generation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Generation), dcl.SprintResource(actual.Generation))
		return true
	}
	return false
}

func compareBuildTriggerBuildSourceStorageSourceSlice(c *Client, desired, actual []BuildTriggerBuildSourceStorageSource) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSourceStorageSource, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuildSourceStorageSource(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSourceStorageSource, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildSourceStorageSourceMap(c *Client, desired, actual map[string]BuildTriggerBuildSourceStorageSource) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSourceStorageSource, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSourceStorageSource, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerBuildSourceStorageSource(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSourceStorageSource, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildSourceRepoSource(c *Client, desired, actual *BuildTriggerBuildSourceRepoSource) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ProjectId == nil && desired.ProjectId != nil && !dcl.IsEmptyValueIndirect(desired.ProjectId) {
		c.Config.Logger.Infof("desired ProjectId %s - but actually nil", dcl.SprintResource(desired.ProjectId))
		return true
	}
	if !dcl.StringCanonicalize(desired.ProjectId, actual.ProjectId) && !dcl.IsZeroValue(desired.ProjectId) {
		c.Config.Logger.Infof("Diff in ProjectId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ProjectId), dcl.SprintResource(actual.ProjectId))
		return true
	}
	if actual.RepoName == nil && desired.RepoName != nil && !dcl.IsEmptyValueIndirect(desired.RepoName) {
		c.Config.Logger.Infof("desired RepoName %s - but actually nil", dcl.SprintResource(desired.RepoName))
		return true
	}
	if !dcl.StringCanonicalize(desired.RepoName, actual.RepoName) && !dcl.IsZeroValue(desired.RepoName) {
		c.Config.Logger.Infof("Diff in RepoName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RepoName), dcl.SprintResource(actual.RepoName))
		return true
	}
	if actual.BranchName == nil && desired.BranchName != nil && !dcl.IsEmptyValueIndirect(desired.BranchName) {
		c.Config.Logger.Infof("desired BranchName %s - but actually nil", dcl.SprintResource(desired.BranchName))
		return true
	}
	if !dcl.StringCanonicalize(desired.BranchName, actual.BranchName) && !dcl.IsZeroValue(desired.BranchName) {
		c.Config.Logger.Infof("Diff in BranchName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BranchName), dcl.SprintResource(actual.BranchName))
		return true
	}
	if actual.TagName == nil && desired.TagName != nil && !dcl.IsEmptyValueIndirect(desired.TagName) {
		c.Config.Logger.Infof("desired TagName %s - but actually nil", dcl.SprintResource(desired.TagName))
		return true
	}
	if !dcl.StringCanonicalize(desired.TagName, actual.TagName) && !dcl.IsZeroValue(desired.TagName) {
		c.Config.Logger.Infof("Diff in TagName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TagName), dcl.SprintResource(actual.TagName))
		return true
	}
	if actual.CommitSha == nil && desired.CommitSha != nil && !dcl.IsEmptyValueIndirect(desired.CommitSha) {
		c.Config.Logger.Infof("desired CommitSha %s - but actually nil", dcl.SprintResource(desired.CommitSha))
		return true
	}
	if !dcl.StringCanonicalize(desired.CommitSha, actual.CommitSha) && !dcl.IsZeroValue(desired.CommitSha) {
		c.Config.Logger.Infof("Diff in CommitSha. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CommitSha), dcl.SprintResource(actual.CommitSha))
		return true
	}
	if actual.Dir == nil && desired.Dir != nil && !dcl.IsEmptyValueIndirect(desired.Dir) {
		c.Config.Logger.Infof("desired Dir %s - but actually nil", dcl.SprintResource(desired.Dir))
		return true
	}
	if !dcl.StringCanonicalize(desired.Dir, actual.Dir) && !dcl.IsZeroValue(desired.Dir) {
		c.Config.Logger.Infof("Diff in Dir. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Dir), dcl.SprintResource(actual.Dir))
		return true
	}
	if actual.InvertRegex == nil && desired.InvertRegex != nil && !dcl.IsEmptyValueIndirect(desired.InvertRegex) {
		c.Config.Logger.Infof("desired InvertRegex %s - but actually nil", dcl.SprintResource(desired.InvertRegex))
		return true
	}
	if !dcl.BoolCanonicalize(desired.InvertRegex, actual.InvertRegex) && !dcl.IsZeroValue(desired.InvertRegex) {
		c.Config.Logger.Infof("Diff in InvertRegex. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InvertRegex), dcl.SprintResource(actual.InvertRegex))
		return true
	}
	if actual.Substitutions == nil && desired.Substitutions != nil && !dcl.IsEmptyValueIndirect(desired.Substitutions) {
		c.Config.Logger.Infof("desired Substitutions %s - but actually nil", dcl.SprintResource(desired.Substitutions))
		return true
	}
	if !dcl.MapEquals(desired.Substitutions, actual.Substitutions, []string(nil)) && !dcl.IsZeroValue(desired.Substitutions) {
		c.Config.Logger.Infof("Diff in Substitutions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Substitutions), dcl.SprintResource(actual.Substitutions))
		return true
	}
	return false
}

func compareBuildTriggerBuildSourceRepoSourceSlice(c *Client, desired, actual []BuildTriggerBuildSourceRepoSource) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSourceRepoSource, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuildSourceRepoSource(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSourceRepoSource, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildSourceRepoSourceMap(c *Client, desired, actual map[string]BuildTriggerBuildSourceRepoSource) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildSourceRepoSource, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSourceRepoSource, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareBuildTriggerBuildSourceRepoSource(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildSourceRepoSource, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareBuildTriggerGithubPullRequestCommentControlEnumSlice(c *Client, desired, actual []BuildTriggerGithubPullRequestCommentControlEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerGithubPullRequestCommentControlEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerGithubPullRequestCommentControlEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerGithubPullRequestCommentControlEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerGithubPullRequestCommentControlEnum(c *Client, desired, actual *BuildTriggerGithubPullRequestCommentControlEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareBuildTriggerBuildStepsStatusEnumSlice(c *Client, desired, actual []BuildTriggerBuildStepsStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in BuildTriggerBuildStepsStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareBuildTriggerBuildStepsStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in BuildTriggerBuildStepsStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareBuildTriggerBuildStepsStatusEnum(c *Client, desired, actual *BuildTriggerBuildStepsStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *BuildTrigger) urlNormalized() *BuildTrigger {
	normalized := deepcopy.Copy(*r).(BuildTrigger)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Filename = dcl.SelfLinkToName(r.Filename)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.CreateTime = dcl.SelfLinkToName(r.CreateTime)
	return &normalized
}

func (r *BuildTrigger) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *BuildTrigger) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *BuildTrigger) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *BuildTrigger) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
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
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["triggerTemplate"] = v
	}
	if v, err := expandBuildTriggerGithub(c, f.Github); err != nil {
		return nil, fmt.Errorf("error expanding Github into github: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["github"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := expandBuildTriggerBuild(c, f.Build); err != nil {
		return nil, fmt.Errorf("error expanding Build into build: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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

	r := &BuildTrigger{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.Tags = dcl.FlattenStringSlice(m["tags"])
	r.Disabled = dcl.FlattenBool(m["disabled"])
	r.Substitutions = dcl.FlattenKeyValuePairs(m["substitutions"])
	r.Filename = dcl.FlattenString(m["filename"])
	r.IgnoredFiles = dcl.FlattenStringSlice(m["ignoredFiles"])
	r.IncludedFiles = dcl.FlattenStringSlice(m["includedFiles"])
	r.TriggerTemplate = flattenBuildTriggerTriggerTemplate(c, m["triggerTemplate"])
	r.Github = flattenBuildTriggerGithub(c, m["github"])
	r.Project = dcl.FlattenString(m["project"])
	r.Build = flattenBuildTriggerBuild(c, m["build"])
	r.Id = dcl.FlattenString(m["id"])
	r.CreateTime = dcl.FlattenString(m["createTime"])

	return r
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
