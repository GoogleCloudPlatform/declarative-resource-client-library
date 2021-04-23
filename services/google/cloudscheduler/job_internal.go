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
package cloudscheduler

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
)

func (r *Job) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.PubsubTarget) {
		if err := r.PubsubTarget.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AppEngineHttpTarget) {
		if err := r.AppEngineHttpTarget.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpTarget) {
		if err := r.HttpTarget.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RetryConfig) {
		if err := r.RetryConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobPubsubTarget) validate() error {
	if err := dcl.Required(r, "topicName"); err != nil {
		return err
	}
	return nil
}
func (r *JobAppEngineHttpTarget) validate() error {
	if !dcl.IsEmptyValueIndirect(r.AppEngineRouting) {
		if err := r.AppEngineRouting.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobAppEngineHttpTargetAppEngineRouting) validate() error {
	return nil
}
func (r *JobHttpTarget) validate() error {
	if err := dcl.Required(r, "uri"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.OAuthToken) {
		if err := r.OAuthToken.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.OidcToken) {
		if err := r.OidcToken.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobHttpTargetOAuthToken) validate() error {
	return nil
}
func (r *JobHttpTargetOidcToken) validate() error {
	return nil
}
func (r *JobStatus) validate() error {
	return nil
}
func (r *JobStatusDetails) validate() error {
	return nil
}
func (r *JobRetryConfig) validate() error {
	return nil
}

func jobGetURL(userBasePath string, r *Job) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs/{{name}}", "https://cloudscheduler.googleapis.com/v1/", userBasePath, params), nil
}

func jobListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs", "https://cloudscheduler.googleapis.com/v1/", userBasePath, params), nil

}

func jobCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs", "https://cloudscheduler.googleapis.com/v1/", userBasePath, params), nil

}

func jobDeleteURL(userBasePath string, r *Job) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs/{{name}}", "https://cloudscheduler.googleapis.com/v1/", userBasePath, params), nil
}

// jobApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type jobApiOperation interface {
	do(context.Context, *Job, *Client) error
}

// newUpdateJobUpdateJobRequest creates a request for an
// Job resource's UpdateJob update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateJobUpdateJobRequest(ctx context.Context, f *Job, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/locations/%s/jobs/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandJobPubsubTarget(c, f.PubsubTarget); err != nil {
		return nil, fmt.Errorf("error expanding PubsubTarget into pubsubTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["pubsubTarget"] = v
	}
	if v, err := expandJobAppEngineHttpTarget(c, f.AppEngineHttpTarget); err != nil {
		return nil, fmt.Errorf("error expanding AppEngineHttpTarget into appEngineHttpTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["appEngineHttpTarget"] = v
	}
	if v, err := expandJobHttpTarget(c, f.HttpTarget); err != nil {
		return nil, fmt.Errorf("error expanding HttpTarget into httpTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["httpTarget"] = v
	}
	if v := f.Schedule; !dcl.IsEmptyValueIndirect(v) {
		req["schedule"] = v
	}
	if v := f.TimeZone; !dcl.IsEmptyValueIndirect(v) {
		req["timeZone"] = v
	}
	if v, err := expandJobRetryConfig(c, f.RetryConfig); err != nil {
		return nil, fmt.Errorf("error expanding RetryConfig into retryConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["retryConfig"] = v
	}
	if v := f.AttemptDeadline; !dcl.IsEmptyValueIndirect(v) {
		req["attemptDeadline"] = v
	}
	return req, nil
}

// marshalUpdateJobUpdateJobRequest converts the update into
// the final JSON request body.
func marshalUpdateJobUpdateJobRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateJobUpdateJobOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateJobUpdateJobOperation) do(ctx context.Context, r *Job, c *Client) error {
	_, err := c.GetJob(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateJob")
	if err != nil {
		return err
	}

	req, err := newUpdateJobUpdateJobRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateJobUpdateJobRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listJobRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := jobListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != JobMaxPage {
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

type listJobOperation struct {
	Jobs  []map[string]interface{} `json:"jobs"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listJob(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Job, string, error) {
	b, err := c.listJobRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listJobOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Job
	for _, v := range m.Jobs {
		res := flattenJob(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllJob(ctx context.Context, f func(*Job) bool, resources []*Job) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteJob(ctx, res)
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

type deleteJobOperation struct{}

func (op *deleteJobOperation) do(ctx context.Context, r *Job, c *Client) error {

	_, err := c.GetJob(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Job not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetJob checking for existence. error: %v", err)
		return err
	}

	u, err := jobDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Job: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetJob(ctx, r.urlNormalized())
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
type createJobOperation struct {
	response map[string]interface{}
}

func (op *createJobOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createJobOperation) do(ctx context.Context, r *Job, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := jobCreateURL(c.Config.BasePath, project, location)

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

	if _, err := c.GetJob(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getJobRaw(ctx context.Context, r *Job) ([]byte, error) {

	u, err := jobGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) jobDiffsForRawDesired(ctx context.Context, rawDesired *Job, opts ...dcl.ApplyOption) (initial, desired *Job, diffs []jobDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Job
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Job); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Job, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetJob(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Job resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Job resource: %v", err)
		}
		c.Config.Logger.Info("Found that Job resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeJobDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Job: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Job: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeJobInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Job: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeJobDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Job: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffJob(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeJobInitialState(rawInitial, rawDesired *Job) (*Job, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeJobDesiredState(rawDesired, rawInitial *Job, opts ...dcl.ApplyOption) (*Job, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.PubsubTarget = canonicalizeJobPubsubTarget(rawDesired.PubsubTarget, nil, opts...)
		rawDesired.AppEngineHttpTarget = canonicalizeJobAppEngineHttpTarget(rawDesired.AppEngineHttpTarget, nil, opts...)
		rawDesired.HttpTarget = canonicalizeJobHttpTarget(rawDesired.HttpTarget, nil, opts...)
		rawDesired.Status = canonicalizeJobStatus(rawDesired.Status, nil, opts...)
		rawDesired.RetryConfig = canonicalizeJobRetryConfig(rawDesired.RetryConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	rawDesired.PubsubTarget = canonicalizeJobPubsubTarget(rawDesired.PubsubTarget, rawInitial.PubsubTarget, opts...)
	rawDesired.AppEngineHttpTarget = canonicalizeJobAppEngineHttpTarget(rawDesired.AppEngineHttpTarget, rawInitial.AppEngineHttpTarget, opts...)
	rawDesired.HttpTarget = canonicalizeJobHttpTarget(rawDesired.HttpTarget, rawInitial.HttpTarget, opts...)
	if dcl.StringCanonicalize(rawDesired.Schedule, rawInitial.Schedule) {
		rawDesired.Schedule = rawInitial.Schedule
	}
	if dcl.StringCanonicalize(rawDesired.TimeZone, rawInitial.TimeZone) {
		rawDesired.TimeZone = rawInitial.TimeZone
	}
	if dcl.IsZeroValue(rawDesired.UserUpdateTime) {
		rawDesired.UserUpdateTime = rawInitial.UserUpdateTime
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	rawDesired.Status = canonicalizeJobStatus(rawDesired.Status, rawInitial.Status, opts...)
	if dcl.IsZeroValue(rawDesired.ScheduleTime) {
		rawDesired.ScheduleTime = rawInitial.ScheduleTime
	}
	if dcl.IsZeroValue(rawDesired.LastAttemptTime) {
		rawDesired.LastAttemptTime = rawInitial.LastAttemptTime
	}
	rawDesired.RetryConfig = canonicalizeJobRetryConfig(rawDesired.RetryConfig, rawInitial.RetryConfig, opts...)
	if dcl.StringCanonicalize(rawDesired.AttemptDeadline, rawInitial.AttemptDeadline) {
		rawDesired.AttemptDeadline = rawInitial.AttemptDeadline
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeJobNewState(c *Client, rawNew, rawDesired *Job) (*Job, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.PubsubTarget) && dcl.IsEmptyValueIndirect(rawDesired.PubsubTarget) {
		rawNew.PubsubTarget = rawDesired.PubsubTarget
	} else {
		rawNew.PubsubTarget = canonicalizeNewJobPubsubTarget(c, rawDesired.PubsubTarget, rawNew.PubsubTarget)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AppEngineHttpTarget) && dcl.IsEmptyValueIndirect(rawDesired.AppEngineHttpTarget) {
		rawNew.AppEngineHttpTarget = rawDesired.AppEngineHttpTarget
	} else {
		rawNew.AppEngineHttpTarget = canonicalizeNewJobAppEngineHttpTarget(c, rawDesired.AppEngineHttpTarget, rawNew.AppEngineHttpTarget)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpTarget) && dcl.IsEmptyValueIndirect(rawDesired.HttpTarget) {
		rawNew.HttpTarget = rawDesired.HttpTarget
	} else {
		rawNew.HttpTarget = canonicalizeNewJobHttpTarget(c, rawDesired.HttpTarget, rawNew.HttpTarget)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Schedule) && dcl.IsEmptyValueIndirect(rawDesired.Schedule) {
		rawNew.Schedule = rawDesired.Schedule
	} else {
		if dcl.StringCanonicalize(rawDesired.Schedule, rawNew.Schedule) {
			rawNew.Schedule = rawDesired.Schedule
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeZone) && dcl.IsEmptyValueIndirect(rawDesired.TimeZone) {
		rawNew.TimeZone = rawDesired.TimeZone
	} else {
		if dcl.StringCanonicalize(rawDesired.TimeZone, rawNew.TimeZone) {
			rawNew.TimeZone = rawDesired.TimeZone
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.UserUpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UserUpdateTime) {
		rawNew.UserUpdateTime = rawDesired.UserUpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
		rawNew.Status = canonicalizeNewJobStatus(c, rawDesired.Status, rawNew.Status)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ScheduleTime) && dcl.IsEmptyValueIndirect(rawDesired.ScheduleTime) {
		rawNew.ScheduleTime = rawDesired.ScheduleTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastAttemptTime) && dcl.IsEmptyValueIndirect(rawDesired.LastAttemptTime) {
		rawNew.LastAttemptTime = rawDesired.LastAttemptTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RetryConfig) && dcl.IsEmptyValueIndirect(rawDesired.RetryConfig) {
		rawNew.RetryConfig = rawDesired.RetryConfig
	} else {
		rawNew.RetryConfig = canonicalizeNewJobRetryConfig(c, rawDesired.RetryConfig, rawNew.RetryConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AttemptDeadline) && dcl.IsEmptyValueIndirect(rawDesired.AttemptDeadline) {
		rawNew.AttemptDeadline = rawDesired.AttemptDeadline
	} else {
		if dcl.StringCanonicalize(rawDesired.AttemptDeadline, rawNew.AttemptDeadline) {
			rawNew.AttemptDeadline = rawDesired.AttemptDeadline
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeJobPubsubTarget(des, initial *JobPubsubTarget, opts ...dcl.ApplyOption) *JobPubsubTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.TopicName, initial.TopicName) || dcl.IsZeroValue(des.TopicName) {
		des.TopicName = initial.TopicName
	}
	if dcl.StringCanonicalize(des.Data, initial.Data) || dcl.IsZeroValue(des.Data) {
		des.Data = initial.Data
	}
	if dcl.IsZeroValue(des.Attributes) {
		des.Attributes = initial.Attributes
	}

	return des
}

func canonicalizeNewJobPubsubTarget(c *Client, des, nw *JobPubsubTarget) *JobPubsubTarget {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.TopicName, nw.TopicName) {
		nw.TopicName = des.TopicName
	}
	if dcl.StringCanonicalize(des.Data, nw.Data) {
		nw.Data = des.Data
	}

	return nw
}

func canonicalizeNewJobPubsubTargetSet(c *Client, des, nw []JobPubsubTarget) []JobPubsubTarget {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPubsubTarget
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPubsubTarget(c, &d, &n) {
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

func canonicalizeNewJobPubsubTargetSlice(c *Client, des, nw []JobPubsubTarget) []JobPubsubTarget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []JobPubsubTarget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobPubsubTarget(c, &d, &n))
	}

	return items
}

func canonicalizeJobAppEngineHttpTarget(des, initial *JobAppEngineHttpTarget, opts ...dcl.ApplyOption) *JobAppEngineHttpTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HttpMethod) {
		des.HttpMethod = initial.HttpMethod
	}
	des.AppEngineRouting = canonicalizeJobAppEngineHttpTargetAppEngineRouting(des.AppEngineRouting, initial.AppEngineRouting, opts...)
	if dcl.StringCanonicalize(des.RelativeUri, initial.RelativeUri) || dcl.IsZeroValue(des.RelativeUri) {
		des.RelativeUri = initial.RelativeUri
	}
	if dcl.IsZeroValue(des.Headers) {
		des.Headers = initial.Headers
	}
	if dcl.StringCanonicalize(des.Body, initial.Body) || dcl.IsZeroValue(des.Body) {
		des.Body = initial.Body
	}

	return des
}

func canonicalizeNewJobAppEngineHttpTarget(c *Client, des, nw *JobAppEngineHttpTarget) *JobAppEngineHttpTarget {
	if des == nil || nw == nil {
		return nw
	}

	nw.AppEngineRouting = canonicalizeNewJobAppEngineHttpTargetAppEngineRouting(c, des.AppEngineRouting, nw.AppEngineRouting)
	if dcl.StringCanonicalize(des.RelativeUri, nw.RelativeUri) {
		nw.RelativeUri = des.RelativeUri
	}
	if dcl.StringCanonicalize(des.Body, nw.Body) {
		nw.Body = des.Body
	}

	return nw
}

func canonicalizeNewJobAppEngineHttpTargetSet(c *Client, des, nw []JobAppEngineHttpTarget) []JobAppEngineHttpTarget {
	if des == nil {
		return nw
	}
	var reorderedNew []JobAppEngineHttpTarget
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobAppEngineHttpTarget(c, &d, &n) {
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

func canonicalizeNewJobAppEngineHttpTargetSlice(c *Client, des, nw []JobAppEngineHttpTarget) []JobAppEngineHttpTarget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []JobAppEngineHttpTarget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobAppEngineHttpTarget(c, &d, &n))
	}

	return items
}

func canonicalizeJobAppEngineHttpTargetAppEngineRouting(des, initial *JobAppEngineHttpTargetAppEngineRouting, opts ...dcl.ApplyOption) *JobAppEngineHttpTargetAppEngineRouting {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		des.Service = initial.Service
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		des.Version = initial.Version
	}
	if dcl.StringCanonicalize(des.Instance, initial.Instance) || dcl.IsZeroValue(des.Instance) {
		des.Instance = initial.Instance
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}

	return des
}

func canonicalizeNewJobAppEngineHttpTargetAppEngineRouting(c *Client, des, nw *JobAppEngineHttpTargetAppEngineRouting) *JobAppEngineHttpTargetAppEngineRouting {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}
	if dcl.StringCanonicalize(des.Instance, nw.Instance) {
		nw.Instance = des.Instance
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}

	return nw
}

func canonicalizeNewJobAppEngineHttpTargetAppEngineRoutingSet(c *Client, des, nw []JobAppEngineHttpTargetAppEngineRouting) []JobAppEngineHttpTargetAppEngineRouting {
	if des == nil {
		return nw
	}
	var reorderedNew []JobAppEngineHttpTargetAppEngineRouting
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobAppEngineHttpTargetAppEngineRouting(c, &d, &n) {
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

func canonicalizeNewJobAppEngineHttpTargetAppEngineRoutingSlice(c *Client, des, nw []JobAppEngineHttpTargetAppEngineRouting) []JobAppEngineHttpTargetAppEngineRouting {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []JobAppEngineHttpTargetAppEngineRouting
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobAppEngineHttpTargetAppEngineRouting(c, &d, &n))
	}

	return items
}

func canonicalizeJobHttpTarget(des, initial *JobHttpTarget, opts ...dcl.ApplyOption) *JobHttpTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Uri, initial.Uri) || dcl.IsZeroValue(des.Uri) {
		des.Uri = initial.Uri
	}
	if dcl.IsZeroValue(des.HttpMethod) {
		des.HttpMethod = initial.HttpMethod
	}
	if dcl.IsZeroValue(des.Headers) {
		des.Headers = initial.Headers
	}
	if dcl.StringCanonicalize(des.Body, initial.Body) || dcl.IsZeroValue(des.Body) {
		des.Body = initial.Body
	}
	des.OAuthToken = canonicalizeJobHttpTargetOAuthToken(des.OAuthToken, initial.OAuthToken, opts...)
	des.OidcToken = canonicalizeJobHttpTargetOidcToken(des.OidcToken, initial.OidcToken, opts...)

	return des
}

func canonicalizeNewJobHttpTarget(c *Client, des, nw *JobHttpTarget) *JobHttpTarget {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Uri, nw.Uri) {
		nw.Uri = des.Uri
	}
	if dcl.StringCanonicalize(des.Body, nw.Body) {
		nw.Body = des.Body
	}
	nw.OAuthToken = canonicalizeNewJobHttpTargetOAuthToken(c, des.OAuthToken, nw.OAuthToken)
	nw.OidcToken = canonicalizeNewJobHttpTargetOidcToken(c, des.OidcToken, nw.OidcToken)

	return nw
}

func canonicalizeNewJobHttpTargetSet(c *Client, des, nw []JobHttpTarget) []JobHttpTarget {
	if des == nil {
		return nw
	}
	var reorderedNew []JobHttpTarget
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobHttpTarget(c, &d, &n) {
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

func canonicalizeNewJobHttpTargetSlice(c *Client, des, nw []JobHttpTarget) []JobHttpTarget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []JobHttpTarget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobHttpTarget(c, &d, &n))
	}

	return items
}

func canonicalizeJobHttpTargetOAuthToken(des, initial *JobHttpTargetOAuthToken, opts ...dcl.ApplyOption) *JobHttpTargetOAuthToken {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.ServiceAccountEmail, initial.ServiceAccountEmail) || dcl.IsZeroValue(des.ServiceAccountEmail) {
		des.ServiceAccountEmail = initial.ServiceAccountEmail
	}
	if dcl.StringCanonicalize(des.Scope, initial.Scope) || dcl.IsZeroValue(des.Scope) {
		des.Scope = initial.Scope
	}

	return des
}

func canonicalizeNewJobHttpTargetOAuthToken(c *Client, des, nw *JobHttpTargetOAuthToken) *JobHttpTargetOAuthToken {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.ServiceAccountEmail, nw.ServiceAccountEmail) {
		nw.ServiceAccountEmail = des.ServiceAccountEmail
	}
	if dcl.StringCanonicalize(des.Scope, nw.Scope) {
		nw.Scope = des.Scope
	}

	return nw
}

func canonicalizeNewJobHttpTargetOAuthTokenSet(c *Client, des, nw []JobHttpTargetOAuthToken) []JobHttpTargetOAuthToken {
	if des == nil {
		return nw
	}
	var reorderedNew []JobHttpTargetOAuthToken
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobHttpTargetOAuthToken(c, &d, &n) {
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

func canonicalizeNewJobHttpTargetOAuthTokenSlice(c *Client, des, nw []JobHttpTargetOAuthToken) []JobHttpTargetOAuthToken {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []JobHttpTargetOAuthToken
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobHttpTargetOAuthToken(c, &d, &n))
	}

	return items
}

func canonicalizeJobHttpTargetOidcToken(des, initial *JobHttpTargetOidcToken, opts ...dcl.ApplyOption) *JobHttpTargetOidcToken {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.ServiceAccountEmail, initial.ServiceAccountEmail) || dcl.IsZeroValue(des.ServiceAccountEmail) {
		des.ServiceAccountEmail = initial.ServiceAccountEmail
	}
	if dcl.StringCanonicalize(des.Audience, initial.Audience) || dcl.IsZeroValue(des.Audience) {
		des.Audience = initial.Audience
	}

	return des
}

func canonicalizeNewJobHttpTargetOidcToken(c *Client, des, nw *JobHttpTargetOidcToken) *JobHttpTargetOidcToken {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.ServiceAccountEmail, nw.ServiceAccountEmail) {
		nw.ServiceAccountEmail = des.ServiceAccountEmail
	}
	if dcl.StringCanonicalize(des.Audience, nw.Audience) {
		nw.Audience = des.Audience
	}

	return nw
}

func canonicalizeNewJobHttpTargetOidcTokenSet(c *Client, des, nw []JobHttpTargetOidcToken) []JobHttpTargetOidcToken {
	if des == nil {
		return nw
	}
	var reorderedNew []JobHttpTargetOidcToken
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobHttpTargetOidcToken(c, &d, &n) {
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

func canonicalizeNewJobHttpTargetOidcTokenSlice(c *Client, des, nw []JobHttpTargetOidcToken) []JobHttpTargetOidcToken {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []JobHttpTargetOidcToken
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobHttpTargetOidcToken(c, &d, &n))
	}

	return items
}

func canonicalizeJobStatus(des, initial *JobStatus, opts ...dcl.ApplyOption) *JobStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Code) {
		des.Code = initial.Code
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		des.Message = initial.Message
	}
	if dcl.IsZeroValue(des.Details) {
		des.Details = initial.Details
	}

	return des
}

func canonicalizeNewJobStatus(c *Client, des, nw *JobStatus) *JobStatus {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}
	nw.Details = canonicalizeNewJobStatusDetailsSlice(c, des.Details, nw.Details)

	return nw
}

func canonicalizeNewJobStatusSet(c *Client, des, nw []JobStatus) []JobStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []JobStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobStatus(c, &d, &n) {
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

func canonicalizeNewJobStatusSlice(c *Client, des, nw []JobStatus) []JobStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []JobStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobStatus(c, &d, &n))
	}

	return items
}

func canonicalizeJobStatusDetails(des, initial *JobStatusDetails, opts ...dcl.ApplyOption) *JobStatusDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.TypeUrl, initial.TypeUrl) || dcl.IsZeroValue(des.TypeUrl) {
		des.TypeUrl = initial.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewJobStatusDetails(c *Client, des, nw *JobStatusDetails) *JobStatusDetails {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.TypeUrl, nw.TypeUrl) {
		nw.TypeUrl = des.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewJobStatusDetailsSet(c *Client, des, nw []JobStatusDetails) []JobStatusDetails {
	if des == nil {
		return nw
	}
	var reorderedNew []JobStatusDetails
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobStatusDetails(c, &d, &n) {
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

func canonicalizeNewJobStatusDetailsSlice(c *Client, des, nw []JobStatusDetails) []JobStatusDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []JobStatusDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobStatusDetails(c, &d, &n))
	}

	return items
}

func canonicalizeJobRetryConfig(des, initial *JobRetryConfig, opts ...dcl.ApplyOption) *JobRetryConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RetryCount) {
		des.RetryCount = initial.RetryCount
	}
	if dcl.StringCanonicalize(des.MaxRetryDuration, initial.MaxRetryDuration) || dcl.IsZeroValue(des.MaxRetryDuration) {
		des.MaxRetryDuration = initial.MaxRetryDuration
	}
	if dcl.StringCanonicalize(des.MinBackoffDuration, initial.MinBackoffDuration) || dcl.IsZeroValue(des.MinBackoffDuration) {
		des.MinBackoffDuration = initial.MinBackoffDuration
	}
	if dcl.StringCanonicalize(des.MaxBackoffDuration, initial.MaxBackoffDuration) || dcl.IsZeroValue(des.MaxBackoffDuration) {
		des.MaxBackoffDuration = initial.MaxBackoffDuration
	}
	if dcl.IsZeroValue(des.MaxDoublings) {
		des.MaxDoublings = initial.MaxDoublings
	}

	return des
}

func canonicalizeNewJobRetryConfig(c *Client, des, nw *JobRetryConfig) *JobRetryConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MaxRetryDuration, nw.MaxRetryDuration) {
		nw.MaxRetryDuration = des.MaxRetryDuration
	}
	if dcl.StringCanonicalize(des.MinBackoffDuration, nw.MinBackoffDuration) {
		nw.MinBackoffDuration = des.MinBackoffDuration
	}
	if dcl.StringCanonicalize(des.MaxBackoffDuration, nw.MaxBackoffDuration) {
		nw.MaxBackoffDuration = des.MaxBackoffDuration
	}

	return nw
}

func canonicalizeNewJobRetryConfigSet(c *Client, des, nw []JobRetryConfig) []JobRetryConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []JobRetryConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobRetryConfig(c, &d, &n) {
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

func canonicalizeNewJobRetryConfigSlice(c *Client, des, nw []JobRetryConfig) []JobRetryConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []JobRetryConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobRetryConfig(c, &d, &n))
	}

	return items
}

type jobDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         jobApiOperation
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
func diffJob(c *Client, desired, actual *Job, opts ...dcl.ApplyOption) ([]jobDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []jobDiff

	var fn dcl.FieldName

	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{
			UpdateOp: &updateJobUpdateJobOperation{}, Diffs: ds,
			FieldName: "Name",
		})
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{
			UpdateOp: &updateJobUpdateJobOperation{}, Diffs: ds,
			FieldName: "Description",
		})
	}

	if ds, err := dcl.Diff(desired.PubsubTarget, actual.PubsubTarget, dcl.Info{ObjectFunction: compareJobPubsubTargetNewStyle}, fn.AddNest("PubsubTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{
			UpdateOp: &updateJobUpdateJobOperation{}, Diffs: ds,
			FieldName: "PubsubTarget",
		})
	}

	if ds, err := dcl.Diff(desired.AppEngineHttpTarget, actual.AppEngineHttpTarget, dcl.Info{ObjectFunction: compareJobAppEngineHttpTargetNewStyle}, fn.AddNest("AppEngineHttpTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{
			UpdateOp: &updateJobUpdateJobOperation{}, Diffs: ds,
			FieldName: "AppEngineHttpTarget",
		})
	}

	if ds, err := dcl.Diff(desired.HttpTarget, actual.HttpTarget, dcl.Info{ObjectFunction: compareJobHttpTargetNewStyle}, fn.AddNest("HttpTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{
			UpdateOp: &updateJobUpdateJobOperation{}, Diffs: ds,
			FieldName: "HttpTarget",
		})
	}

	if ds, err := dcl.Diff(desired.Schedule, actual.Schedule, dcl.Info{}, fn.AddNest("Schedule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{
			UpdateOp: &updateJobUpdateJobOperation{}, Diffs: ds,
			FieldName: "Schedule",
		})
	}

	if ds, err := dcl.Diff(desired.TimeZone, actual.TimeZone, dcl.Info{}, fn.AddNest("TimeZone")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{
			UpdateOp: &updateJobUpdateJobOperation{}, Diffs: ds,
			FieldName: "TimeZone",
		})
	}

	if ds, err := dcl.Diff(desired.UserUpdateTime, actual.UserUpdateTime, dcl.Info{OutputOnly: true}, fn.AddNest("UserUpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "UserUpdateTime",
		})
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType"}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "State",
		})
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, ObjectFunction: compareJobStatusNewStyle}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Status",
		})
	}

	if ds, err := dcl.Diff(desired.ScheduleTime, actual.ScheduleTime, dcl.Info{OutputOnly: true}, fn.AddNest("ScheduleTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ScheduleTime",
		})
	}

	if ds, err := dcl.Diff(desired.LastAttemptTime, actual.LastAttemptTime, dcl.Info{OutputOnly: true}, fn.AddNest("LastAttemptTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "LastAttemptTime",
		})
	}

	if ds, err := dcl.Diff(desired.RetryConfig, actual.RetryConfig, dcl.Info{ObjectFunction: compareJobRetryConfigNewStyle}, fn.AddNest("RetryConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{
			UpdateOp: &updateJobUpdateJobOperation{}, Diffs: ds,
			FieldName: "RetryConfig",
		})
	}

	if ds, err := dcl.Diff(desired.AttemptDeadline, actual.AttemptDeadline, dcl.Info{}, fn.AddNest("AttemptDeadline")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{
			UpdateOp: &updateJobUpdateJobOperation{}, Diffs: ds,
			FieldName: "AttemptDeadline",
		})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Project",
		})
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, jobDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Location",
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
	var deduped []jobDiff
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
func compareJobPubsubTargetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobPubsubTarget)
	if !ok {
		desiredNotPointer, ok := d.(JobPubsubTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobPubsubTarget or *JobPubsubTarget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobPubsubTarget)
	if !ok {
		actualNotPointer, ok := a.(JobPubsubTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobPubsubTarget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TopicName, actual.TopicName, dcl.Info{Type: "ReferenceType"}, fn.AddNest("TopicName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Data, actual.Data, dcl.Info{}, fn.AddNest("Data")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Attributes, actual.Attributes, dcl.Info{}, fn.AddNest("Attributes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobPubsubTarget(c *Client, desired, actual *JobPubsubTarget) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.NameToSelfLink(desired.TopicName, actual.TopicName) && !dcl.IsZeroValue(desired.TopicName) {
		c.Config.Logger.Infof("Diff in TopicName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TopicName), dcl.SprintResource(actual.TopicName))
		return true
	}
	if !dcl.StringCanonicalize(desired.Data, actual.Data) && !dcl.IsZeroValue(desired.Data) {
		c.Config.Logger.Infof("Diff in Data.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Data), dcl.SprintResource(actual.Data))
		return true
	}
	if !dcl.MapEquals(desired.Attributes, actual.Attributes, []string(nil)) && !dcl.IsZeroValue(desired.Attributes) {
		c.Config.Logger.Infof("Diff in Attributes.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Attributes), dcl.SprintResource(actual.Attributes))
		return true
	}
	return false
}

func compareJobPubsubTargetSlice(c *Client, desired, actual []JobPubsubTarget) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPubsubTarget, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPubsubTarget(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPubsubTarget, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPubsubTargetMap(c *Client, desired, actual map[string]JobPubsubTarget) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPubsubTarget, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPubsubTarget, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPubsubTarget(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPubsubTarget, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobAppEngineHttpTargetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobAppEngineHttpTarget)
	if !ok {
		desiredNotPointer, ok := d.(JobAppEngineHttpTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobAppEngineHttpTarget or *JobAppEngineHttpTarget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobAppEngineHttpTarget)
	if !ok {
		actualNotPointer, ok := a.(JobAppEngineHttpTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobAppEngineHttpTarget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpMethod, actual.HttpMethod, dcl.Info{Type: "EnumType"}, fn.AddNest("HttpMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AppEngineRouting, actual.AppEngineRouting, dcl.Info{ObjectFunction: compareJobAppEngineHttpTargetAppEngineRoutingNewStyle}, fn.AddNest("AppEngineRouting")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RelativeUri, actual.RelativeUri, dcl.Info{}, fn.AddNest("RelativeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Headers, actual.Headers, dcl.Info{}, fn.AddNest("Headers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Body, actual.Body, dcl.Info{}, fn.AddNest("Body")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobAppEngineHttpTarget(c *Client, desired, actual *JobAppEngineHttpTarget) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.HttpMethod, actual.HttpMethod) && !dcl.IsZeroValue(desired.HttpMethod) {
		c.Config.Logger.Infof("Diff in HttpMethod.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpMethod), dcl.SprintResource(actual.HttpMethod))
		return true
	}
	if compareJobAppEngineHttpTargetAppEngineRouting(c, desired.AppEngineRouting, actual.AppEngineRouting) && !dcl.IsZeroValue(desired.AppEngineRouting) {
		c.Config.Logger.Infof("Diff in AppEngineRouting.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AppEngineRouting), dcl.SprintResource(actual.AppEngineRouting))
		return true
	}
	if !dcl.StringCanonicalize(desired.RelativeUri, actual.RelativeUri) && !dcl.IsZeroValue(desired.RelativeUri) {
		c.Config.Logger.Infof("Diff in RelativeUri.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RelativeUri), dcl.SprintResource(actual.RelativeUri))
		return true
	}
	if !dcl.MapEquals(desired.Headers, actual.Headers, []string(nil)) && !dcl.IsZeroValue(desired.Headers) {
		c.Config.Logger.Infof("Diff in Headers.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Headers), dcl.SprintResource(actual.Headers))
		return true
	}
	if !dcl.StringCanonicalize(desired.Body, actual.Body) && !dcl.IsZeroValue(desired.Body) {
		c.Config.Logger.Infof("Diff in Body.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Body), dcl.SprintResource(actual.Body))
		return true
	}
	return false
}

func compareJobAppEngineHttpTargetSlice(c *Client, desired, actual []JobAppEngineHttpTarget) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobAppEngineHttpTarget, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobAppEngineHttpTarget(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTarget, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobAppEngineHttpTargetMap(c *Client, desired, actual map[string]JobAppEngineHttpTarget) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobAppEngineHttpTarget, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTarget, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobAppEngineHttpTarget(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTarget, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobAppEngineHttpTargetAppEngineRoutingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobAppEngineHttpTargetAppEngineRouting)
	if !ok {
		desiredNotPointer, ok := d.(JobAppEngineHttpTargetAppEngineRouting)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobAppEngineHttpTargetAppEngineRouting or *JobAppEngineHttpTargetAppEngineRouting", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobAppEngineHttpTargetAppEngineRouting)
	if !ok {
		actualNotPointer, ok := a.(JobAppEngineHttpTargetAppEngineRouting)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobAppEngineHttpTargetAppEngineRouting", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Instance, actual.Instance, dcl.Info{}, fn.AddNest("Instance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OutputOnly: true}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobAppEngineHttpTargetAppEngineRouting(c *Client, desired, actual *JobAppEngineHttpTargetAppEngineRouting) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Service, actual.Service) && !dcl.IsZeroValue(desired.Service) {
		c.Config.Logger.Infof("Diff in Service.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Service), dcl.SprintResource(actual.Service))
		return true
	}
	if !dcl.StringCanonicalize(desired.Version, actual.Version) && !dcl.IsZeroValue(desired.Version) {
		c.Config.Logger.Infof("Diff in Version.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Version), dcl.SprintResource(actual.Version))
		return true
	}
	if !dcl.StringCanonicalize(desired.Instance, actual.Instance) && !dcl.IsZeroValue(desired.Instance) {
		c.Config.Logger.Infof("Diff in Instance.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Instance), dcl.SprintResource(actual.Instance))
		return true
	}
	return false
}

func compareJobAppEngineHttpTargetAppEngineRoutingSlice(c *Client, desired, actual []JobAppEngineHttpTargetAppEngineRouting) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobAppEngineHttpTargetAppEngineRouting, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobAppEngineHttpTargetAppEngineRouting(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTargetAppEngineRouting, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobAppEngineHttpTargetAppEngineRoutingMap(c *Client, desired, actual map[string]JobAppEngineHttpTargetAppEngineRouting) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobAppEngineHttpTargetAppEngineRouting, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTargetAppEngineRouting, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobAppEngineHttpTargetAppEngineRouting(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTargetAppEngineRouting, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobHttpTargetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobHttpTarget)
	if !ok {
		desiredNotPointer, ok := d.(JobHttpTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTarget or *JobHttpTarget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobHttpTarget)
	if !ok {
		actualNotPointer, ok := a.(JobHttpTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTarget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.Info{}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpMethod, actual.HttpMethod, dcl.Info{Type: "EnumType"}, fn.AddNest("HttpMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Headers, actual.Headers, dcl.Info{}, fn.AddNest("Headers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Body, actual.Body, dcl.Info{}, fn.AddNest("Body")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuthToken, actual.OAuthToken, dcl.Info{ObjectFunction: compareJobHttpTargetOAuthTokenNewStyle}, fn.AddNest("OAuthToken")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OidcToken, actual.OidcToken, dcl.Info{ObjectFunction: compareJobHttpTargetOidcTokenNewStyle}, fn.AddNest("OidcToken")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobHttpTarget(c *Client, desired, actual *JobHttpTarget) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Uri, actual.Uri) && !dcl.IsZeroValue(desired.Uri) {
		c.Config.Logger.Infof("Diff in Uri.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Uri), dcl.SprintResource(actual.Uri))
		return true
	}
	if !reflect.DeepEqual(desired.HttpMethod, actual.HttpMethod) && !dcl.IsZeroValue(desired.HttpMethod) {
		c.Config.Logger.Infof("Diff in HttpMethod.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpMethod), dcl.SprintResource(actual.HttpMethod))
		return true
	}
	if !dcl.MapEquals(desired.Headers, actual.Headers, []string(nil)) && !dcl.IsZeroValue(desired.Headers) {
		c.Config.Logger.Infof("Diff in Headers.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Headers), dcl.SprintResource(actual.Headers))
		return true
	}
	if !dcl.StringCanonicalize(desired.Body, actual.Body) && !dcl.IsZeroValue(desired.Body) {
		c.Config.Logger.Infof("Diff in Body.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Body), dcl.SprintResource(actual.Body))
		return true
	}
	if compareJobHttpTargetOAuthToken(c, desired.OAuthToken, actual.OAuthToken) && !dcl.IsZeroValue(desired.OAuthToken) {
		c.Config.Logger.Infof("Diff in OAuthToken.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OAuthToken), dcl.SprintResource(actual.OAuthToken))
		return true
	}
	if compareJobHttpTargetOidcToken(c, desired.OidcToken, actual.OidcToken) && !dcl.IsZeroValue(desired.OidcToken) {
		c.Config.Logger.Infof("Diff in OidcToken.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OidcToken), dcl.SprintResource(actual.OidcToken))
		return true
	}
	return false
}

func compareJobHttpTargetSlice(c *Client, desired, actual []JobHttpTarget) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHttpTarget, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobHttpTarget(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobHttpTarget, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHttpTargetMap(c *Client, desired, actual map[string]JobHttpTarget) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHttpTarget, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobHttpTarget, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobHttpTarget(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobHttpTarget, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobHttpTargetOAuthTokenNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobHttpTargetOAuthToken)
	if !ok {
		desiredNotPointer, ok := d.(JobHttpTargetOAuthToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTargetOAuthToken or *JobHttpTargetOAuthToken", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobHttpTargetOAuthToken)
	if !ok {
		actualNotPointer, ok := a.(JobHttpTargetOAuthToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTargetOAuthToken", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServiceAccountEmail, actual.ServiceAccountEmail, dcl.Info{Type: "ReferenceType"}, fn.AddNest("ServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scope, actual.Scope, dcl.Info{}, fn.AddNest("Scope")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobHttpTargetOAuthToken(c *Client, desired, actual *JobHttpTargetOAuthToken) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.NameToSelfLink(desired.ServiceAccountEmail, actual.ServiceAccountEmail) && !dcl.IsZeroValue(desired.ServiceAccountEmail) {
		c.Config.Logger.Infof("Diff in ServiceAccountEmail.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccountEmail), dcl.SprintResource(actual.ServiceAccountEmail))
		return true
	}
	if !dcl.StringCanonicalize(desired.Scope, actual.Scope) && !dcl.IsZeroValue(desired.Scope) {
		c.Config.Logger.Infof("Diff in Scope.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scope), dcl.SprintResource(actual.Scope))
		return true
	}
	return false
}

func compareJobHttpTargetOAuthTokenSlice(c *Client, desired, actual []JobHttpTargetOAuthToken) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHttpTargetOAuthToken, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobHttpTargetOAuthToken(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobHttpTargetOAuthToken, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHttpTargetOAuthTokenMap(c *Client, desired, actual map[string]JobHttpTargetOAuthToken) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHttpTargetOAuthToken, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobHttpTargetOAuthToken, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobHttpTargetOAuthToken(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobHttpTargetOAuthToken, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobHttpTargetOidcTokenNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobHttpTargetOidcToken)
	if !ok {
		desiredNotPointer, ok := d.(JobHttpTargetOidcToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTargetOidcToken or *JobHttpTargetOidcToken", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobHttpTargetOidcToken)
	if !ok {
		actualNotPointer, ok := a.(JobHttpTargetOidcToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTargetOidcToken", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServiceAccountEmail, actual.ServiceAccountEmail, dcl.Info{Type: "ReferenceType"}, fn.AddNest("ServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Audience, actual.Audience, dcl.Info{}, fn.AddNest("Audience")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobHttpTargetOidcToken(c *Client, desired, actual *JobHttpTargetOidcToken) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.NameToSelfLink(desired.ServiceAccountEmail, actual.ServiceAccountEmail) && !dcl.IsZeroValue(desired.ServiceAccountEmail) {
		c.Config.Logger.Infof("Diff in ServiceAccountEmail.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccountEmail), dcl.SprintResource(actual.ServiceAccountEmail))
		return true
	}
	if !dcl.StringCanonicalize(desired.Audience, actual.Audience) && !dcl.IsZeroValue(desired.Audience) {
		c.Config.Logger.Infof("Diff in Audience.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Audience), dcl.SprintResource(actual.Audience))
		return true
	}
	return false
}

func compareJobHttpTargetOidcTokenSlice(c *Client, desired, actual []JobHttpTargetOidcToken) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHttpTargetOidcToken, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobHttpTargetOidcToken(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobHttpTargetOidcToken, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHttpTargetOidcTokenMap(c *Client, desired, actual map[string]JobHttpTargetOidcToken) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHttpTargetOidcToken, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobHttpTargetOidcToken, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobHttpTargetOidcToken(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobHttpTargetOidcToken, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobStatus)
	if !ok {
		desiredNotPointer, ok := d.(JobStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobStatus or *JobStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobStatus)
	if !ok {
		actualNotPointer, ok := a.(JobStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Details, actual.Details, dcl.Info{ObjectFunction: compareJobStatusDetailsNewStyle}, fn.AddNest("Details")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobStatus(c *Client, desired, actual *JobStatus) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Code, actual.Code) && !dcl.IsZeroValue(desired.Code) {
		c.Config.Logger.Infof("Diff in Code.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Code), dcl.SprintResource(actual.Code))
		return true
	}
	if !dcl.StringCanonicalize(desired.Message, actual.Message) && !dcl.IsZeroValue(desired.Message) {
		c.Config.Logger.Infof("Diff in Message.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Message), dcl.SprintResource(actual.Message))
		return true
	}
	if compareJobStatusDetailsSlice(c, desired.Details, actual.Details) && !dcl.IsZeroValue(desired.Details) {
		c.Config.Logger.Infof("Diff in Details.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Details), dcl.SprintResource(actual.Details))
		return true
	}
	return false
}

func compareJobStatusSlice(c *Client, desired, actual []JobStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatus, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobStatus(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobStatus, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStatusMap(c *Client, desired, actual map[string]JobStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatus, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobStatus, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobStatus(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobStatus, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobStatusDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobStatusDetails)
	if !ok {
		desiredNotPointer, ok := d.(JobStatusDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobStatusDetails or *JobStatusDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobStatusDetails)
	if !ok {
		actualNotPointer, ok := a.(JobStatusDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobStatusDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TypeUrl, actual.TypeUrl, dcl.Info{}, fn.AddNest("TypeUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobStatusDetails(c *Client, desired, actual *JobStatusDetails) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.TypeUrl, actual.TypeUrl) && !dcl.IsZeroValue(desired.TypeUrl) {
		c.Config.Logger.Infof("Diff in TypeUrl.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TypeUrl), dcl.SprintResource(actual.TypeUrl))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareJobStatusDetailsSlice(c *Client, desired, actual []JobStatusDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatusDetails, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobStatusDetails(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobStatusDetails, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStatusDetailsMap(c *Client, desired, actual map[string]JobStatusDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatusDetails, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobStatusDetails, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobStatusDetails(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobStatusDetails, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobRetryConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobRetryConfig)
	if !ok {
		desiredNotPointer, ok := d.(JobRetryConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobRetryConfig or *JobRetryConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobRetryConfig)
	if !ok {
		actualNotPointer, ok := a.(JobRetryConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobRetryConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RetryCount, actual.RetryCount, dcl.Info{}, fn.AddNest("RetryCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxRetryDuration, actual.MaxRetryDuration, dcl.Info{}, fn.AddNest("MaxRetryDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinBackoffDuration, actual.MinBackoffDuration, dcl.Info{}, fn.AddNest("MinBackoffDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxBackoffDuration, actual.MaxBackoffDuration, dcl.Info{}, fn.AddNest("MaxBackoffDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxDoublings, actual.MaxDoublings, dcl.Info{}, fn.AddNest("MaxDoublings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobRetryConfig(c *Client, desired, actual *JobRetryConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.RetryCount, actual.RetryCount) && !dcl.IsZeroValue(desired.RetryCount) {
		c.Config.Logger.Infof("Diff in RetryCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetryCount), dcl.SprintResource(actual.RetryCount))
		return true
	}
	if !dcl.StringCanonicalize(desired.MaxRetryDuration, actual.MaxRetryDuration) && !dcl.IsZeroValue(desired.MaxRetryDuration) {
		c.Config.Logger.Infof("Diff in MaxRetryDuration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxRetryDuration), dcl.SprintResource(actual.MaxRetryDuration))
		return true
	}
	if !dcl.StringCanonicalize(desired.MinBackoffDuration, actual.MinBackoffDuration) && !dcl.IsZeroValue(desired.MinBackoffDuration) {
		c.Config.Logger.Infof("Diff in MinBackoffDuration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinBackoffDuration), dcl.SprintResource(actual.MinBackoffDuration))
		return true
	}
	if !dcl.StringCanonicalize(desired.MaxBackoffDuration, actual.MaxBackoffDuration) && !dcl.IsZeroValue(desired.MaxBackoffDuration) {
		c.Config.Logger.Infof("Diff in MaxBackoffDuration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxBackoffDuration), dcl.SprintResource(actual.MaxBackoffDuration))
		return true
	}
	if !reflect.DeepEqual(desired.MaxDoublings, actual.MaxDoublings) && !dcl.IsZeroValue(desired.MaxDoublings) {
		c.Config.Logger.Infof("Diff in MaxDoublings.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxDoublings), dcl.SprintResource(actual.MaxDoublings))
		return true
	}
	return false
}

func compareJobRetryConfigSlice(c *Client, desired, actual []JobRetryConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobRetryConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobRetryConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobRetryConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobRetryConfigMap(c *Client, desired, actual map[string]JobRetryConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobRetryConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobRetryConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobRetryConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobRetryConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobAppEngineHttpTargetHttpMethodEnumSlice(c *Client, desired, actual []JobAppEngineHttpTargetHttpMethodEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobAppEngineHttpTargetHttpMethodEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobAppEngineHttpTargetHttpMethodEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTargetHttpMethodEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobAppEngineHttpTargetHttpMethodEnum(c *Client, desired, actual *JobAppEngineHttpTargetHttpMethodEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareJobHttpTargetHttpMethodEnumSlice(c *Client, desired, actual []JobHttpTargetHttpMethodEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHttpTargetHttpMethodEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobHttpTargetHttpMethodEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobHttpTargetHttpMethodEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHttpTargetHttpMethodEnum(c *Client, desired, actual *JobHttpTargetHttpMethodEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareJobStateEnumSlice(c *Client, desired, actual []JobStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobStateEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStateEnum(c *Client, desired, actual *JobStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Job) urlNormalized() *Job {
	normalized := dcl.Copy(*r).(Job)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Schedule = dcl.SelfLinkToName(r.Schedule)
	normalized.TimeZone = dcl.SelfLinkToName(r.TimeZone)
	normalized.AttemptDeadline = dcl.SelfLinkToName(r.AttemptDeadline)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Job) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Job) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *Job) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Job) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateJob" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/jobs/{{name}}", "https://cloudscheduler.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Job resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Job) marshal(c *Client) ([]byte, error) {
	m, err := expandJob(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Job: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalJob decodes JSON responses into the Job resource schema.
func unmarshalJob(b []byte, c *Client) (*Job, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapJob(m, c)
}

func unmarshalMapJob(m map[string]interface{}, c *Client) (*Job, error) {

	return flattenJob(c, m), nil
}

// expandJob expands Job into a JSON request object.
func expandJob(c *Client, f *Job) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/jobs/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandJobPubsubTarget(c, f.PubsubTarget); err != nil {
		return nil, fmt.Errorf("error expanding PubsubTarget into pubsubTarget: %w", err)
	} else if v != nil {
		m["pubsubTarget"] = v
	}
	if v, err := expandJobAppEngineHttpTarget(c, f.AppEngineHttpTarget); err != nil {
		return nil, fmt.Errorf("error expanding AppEngineHttpTarget into appEngineHttpTarget: %w", err)
	} else if v != nil {
		m["appEngineHttpTarget"] = v
	}
	if v, err := expandJobHttpTarget(c, f.HttpTarget); err != nil {
		return nil, fmt.Errorf("error expanding HttpTarget into httpTarget: %w", err)
	} else if v != nil {
		m["httpTarget"] = v
	}
	if v := f.Schedule; !dcl.IsEmptyValueIndirect(v) {
		m["schedule"] = v
	}
	if v := f.TimeZone; !dcl.IsEmptyValueIndirect(v) {
		m["timeZone"] = v
	}
	if v := f.UserUpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["userUpdateTime"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := expandJobStatus(c, f.Status); err != nil {
		return nil, fmt.Errorf("error expanding Status into status: %w", err)
	} else if v != nil {
		m["status"] = v
	}
	if v := f.ScheduleTime; !dcl.IsEmptyValueIndirect(v) {
		m["scheduleTime"] = v
	}
	if v := f.LastAttemptTime; !dcl.IsEmptyValueIndirect(v) {
		m["lastAttemptTime"] = v
	}
	if v, err := expandJobRetryConfig(c, f.RetryConfig); err != nil {
		return nil, fmt.Errorf("error expanding RetryConfig into retryConfig: %w", err)
	} else if v != nil {
		m["retryConfig"] = v
	}
	if v := f.AttemptDeadline; !dcl.IsEmptyValueIndirect(v) {
		m["attemptDeadline"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenJob flattens Job from a JSON request object into the
// Job type.
func flattenJob(c *Client, i interface{}) *Job {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Job{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.PubsubTarget = flattenJobPubsubTarget(c, m["pubsubTarget"])
	r.AppEngineHttpTarget = flattenJobAppEngineHttpTarget(c, m["appEngineHttpTarget"])
	r.HttpTarget = flattenJobHttpTarget(c, m["httpTarget"])
	r.Schedule = dcl.FlattenString(m["schedule"])
	r.TimeZone = dcl.FlattenString(m["timeZone"])
	r.UserUpdateTime = dcl.FlattenString(m["userUpdateTime"])
	r.State = flattenJobStateEnum(m["state"])
	r.Status = flattenJobStatus(c, m["status"])
	r.ScheduleTime = dcl.FlattenString(m["scheduleTime"])
	r.LastAttemptTime = dcl.FlattenString(m["lastAttemptTime"])
	r.RetryConfig = flattenJobRetryConfig(c, m["retryConfig"])
	r.AttemptDeadline = dcl.FlattenString(m["attemptDeadline"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandJobPubsubTargetMap expands the contents of JobPubsubTarget into a JSON
// request object.
func expandJobPubsubTargetMap(c *Client, f map[string]JobPubsubTarget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPubsubTarget(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPubsubTargetSlice expands the contents of JobPubsubTarget into a JSON
// request object.
func expandJobPubsubTargetSlice(c *Client, f []JobPubsubTarget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPubsubTarget(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPubsubTargetMap flattens the contents of JobPubsubTarget from a JSON
// response object.
func flattenJobPubsubTargetMap(c *Client, i interface{}) map[string]JobPubsubTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPubsubTarget{}
	}

	if len(a) == 0 {
		return map[string]JobPubsubTarget{}
	}

	items := make(map[string]JobPubsubTarget)
	for k, item := range a {
		items[k] = *flattenJobPubsubTarget(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPubsubTargetSlice flattens the contents of JobPubsubTarget from a JSON
// response object.
func flattenJobPubsubTargetSlice(c *Client, i interface{}) []JobPubsubTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPubsubTarget{}
	}

	if len(a) == 0 {
		return []JobPubsubTarget{}
	}

	items := make([]JobPubsubTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPubsubTarget(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPubsubTarget expands an instance of JobPubsubTarget into a JSON
// request object.
func expandJobPubsubTarget(c *Client, f *JobPubsubTarget) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.TopicName; !dcl.IsEmptyValueIndirect(v) {
		m["topicName"] = v
	}
	if v := f.Data; !dcl.IsEmptyValueIndirect(v) {
		m["data"] = v
	}
	if v := f.Attributes; !dcl.IsEmptyValueIndirect(v) {
		m["attributes"] = v
	}

	return m, nil
}

// flattenJobPubsubTarget flattens an instance of JobPubsubTarget from a JSON
// response object.
func flattenJobPubsubTarget(c *Client, i interface{}) *JobPubsubTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPubsubTarget{}
	r.TopicName = dcl.FlattenString(m["topicName"])
	r.Data = dcl.FlattenString(m["data"])
	r.Attributes = dcl.FlattenKeyValuePairs(m["attributes"])

	return r
}

// expandJobAppEngineHttpTargetMap expands the contents of JobAppEngineHttpTarget into a JSON
// request object.
func expandJobAppEngineHttpTargetMap(c *Client, f map[string]JobAppEngineHttpTarget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobAppEngineHttpTarget(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobAppEngineHttpTargetSlice expands the contents of JobAppEngineHttpTarget into a JSON
// request object.
func expandJobAppEngineHttpTargetSlice(c *Client, f []JobAppEngineHttpTarget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobAppEngineHttpTarget(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobAppEngineHttpTargetMap flattens the contents of JobAppEngineHttpTarget from a JSON
// response object.
func flattenJobAppEngineHttpTargetMap(c *Client, i interface{}) map[string]JobAppEngineHttpTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobAppEngineHttpTarget{}
	}

	if len(a) == 0 {
		return map[string]JobAppEngineHttpTarget{}
	}

	items := make(map[string]JobAppEngineHttpTarget)
	for k, item := range a {
		items[k] = *flattenJobAppEngineHttpTarget(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobAppEngineHttpTargetSlice flattens the contents of JobAppEngineHttpTarget from a JSON
// response object.
func flattenJobAppEngineHttpTargetSlice(c *Client, i interface{}) []JobAppEngineHttpTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []JobAppEngineHttpTarget{}
	}

	if len(a) == 0 {
		return []JobAppEngineHttpTarget{}
	}

	items := make([]JobAppEngineHttpTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobAppEngineHttpTarget(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobAppEngineHttpTarget expands an instance of JobAppEngineHttpTarget into a JSON
// request object.
func expandJobAppEngineHttpTarget(c *Client, f *JobAppEngineHttpTarget) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.HttpMethod; !dcl.IsEmptyValueIndirect(v) {
		m["httpMethod"] = v
	}
	if v, err := expandJobAppEngineHttpTargetAppEngineRouting(c, f.AppEngineRouting); err != nil {
		return nil, fmt.Errorf("error expanding AppEngineRouting into appEngineRouting: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["appEngineRouting"] = v
	}
	if v := f.RelativeUri; !dcl.IsEmptyValueIndirect(v) {
		m["relativeUri"] = v
	}
	if v := f.Headers; !dcl.IsEmptyValueIndirect(v) {
		m["headers"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}

	return m, nil
}

// flattenJobAppEngineHttpTarget flattens an instance of JobAppEngineHttpTarget from a JSON
// response object.
func flattenJobAppEngineHttpTarget(c *Client, i interface{}) *JobAppEngineHttpTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobAppEngineHttpTarget{}
	r.HttpMethod = flattenJobAppEngineHttpTargetHttpMethodEnum(m["httpMethod"])
	r.AppEngineRouting = flattenJobAppEngineHttpTargetAppEngineRouting(c, m["appEngineRouting"])
	r.RelativeUri = dcl.FlattenString(m["relativeUri"])
	r.Headers = dcl.FlattenKeyValuePairs(m["headers"])
	r.Body = dcl.FlattenString(m["body"])

	return r
}

// expandJobAppEngineHttpTargetAppEngineRoutingMap expands the contents of JobAppEngineHttpTargetAppEngineRouting into a JSON
// request object.
func expandJobAppEngineHttpTargetAppEngineRoutingMap(c *Client, f map[string]JobAppEngineHttpTargetAppEngineRouting) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobAppEngineHttpTargetAppEngineRouting(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobAppEngineHttpTargetAppEngineRoutingSlice expands the contents of JobAppEngineHttpTargetAppEngineRouting into a JSON
// request object.
func expandJobAppEngineHttpTargetAppEngineRoutingSlice(c *Client, f []JobAppEngineHttpTargetAppEngineRouting) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobAppEngineHttpTargetAppEngineRouting(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobAppEngineHttpTargetAppEngineRoutingMap flattens the contents of JobAppEngineHttpTargetAppEngineRouting from a JSON
// response object.
func flattenJobAppEngineHttpTargetAppEngineRoutingMap(c *Client, i interface{}) map[string]JobAppEngineHttpTargetAppEngineRouting {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobAppEngineHttpTargetAppEngineRouting{}
	}

	if len(a) == 0 {
		return map[string]JobAppEngineHttpTargetAppEngineRouting{}
	}

	items := make(map[string]JobAppEngineHttpTargetAppEngineRouting)
	for k, item := range a {
		items[k] = *flattenJobAppEngineHttpTargetAppEngineRouting(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobAppEngineHttpTargetAppEngineRoutingSlice flattens the contents of JobAppEngineHttpTargetAppEngineRouting from a JSON
// response object.
func flattenJobAppEngineHttpTargetAppEngineRoutingSlice(c *Client, i interface{}) []JobAppEngineHttpTargetAppEngineRouting {
	a, ok := i.([]interface{})
	if !ok {
		return []JobAppEngineHttpTargetAppEngineRouting{}
	}

	if len(a) == 0 {
		return []JobAppEngineHttpTargetAppEngineRouting{}
	}

	items := make([]JobAppEngineHttpTargetAppEngineRouting, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobAppEngineHttpTargetAppEngineRouting(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobAppEngineHttpTargetAppEngineRouting expands an instance of JobAppEngineHttpTargetAppEngineRouting into a JSON
// request object.
func expandJobAppEngineHttpTargetAppEngineRouting(c *Client, f *JobAppEngineHttpTargetAppEngineRouting) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.Instance; !dcl.IsEmptyValueIndirect(v) {
		m["instance"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}

	return m, nil
}

// flattenJobAppEngineHttpTargetAppEngineRouting flattens an instance of JobAppEngineHttpTargetAppEngineRouting from a JSON
// response object.
func flattenJobAppEngineHttpTargetAppEngineRouting(c *Client, i interface{}) *JobAppEngineHttpTargetAppEngineRouting {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobAppEngineHttpTargetAppEngineRouting{}
	r.Service = dcl.FlattenString(m["service"])
	r.Version = dcl.FlattenString(m["version"])
	r.Instance = dcl.FlattenString(m["instance"])
	r.Host = dcl.FlattenString(m["host"])

	return r
}

// expandJobHttpTargetMap expands the contents of JobHttpTarget into a JSON
// request object.
func expandJobHttpTargetMap(c *Client, f map[string]JobHttpTarget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHttpTarget(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHttpTargetSlice expands the contents of JobHttpTarget into a JSON
// request object.
func expandJobHttpTargetSlice(c *Client, f []JobHttpTarget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHttpTarget(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHttpTargetMap flattens the contents of JobHttpTarget from a JSON
// response object.
func flattenJobHttpTargetMap(c *Client, i interface{}) map[string]JobHttpTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHttpTarget{}
	}

	if len(a) == 0 {
		return map[string]JobHttpTarget{}
	}

	items := make(map[string]JobHttpTarget)
	for k, item := range a {
		items[k] = *flattenJobHttpTarget(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobHttpTargetSlice flattens the contents of JobHttpTarget from a JSON
// response object.
func flattenJobHttpTargetSlice(c *Client, i interface{}) []JobHttpTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHttpTarget{}
	}

	if len(a) == 0 {
		return []JobHttpTarget{}
	}

	items := make([]JobHttpTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHttpTarget(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobHttpTarget expands an instance of JobHttpTarget into a JSON
// request object.
func expandJobHttpTarget(c *Client, f *JobHttpTarget) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Uri; !dcl.IsEmptyValueIndirect(v) {
		m["uri"] = v
	}
	if v := f.HttpMethod; !dcl.IsEmptyValueIndirect(v) {
		m["httpMethod"] = v
	}
	if v := f.Headers; !dcl.IsEmptyValueIndirect(v) {
		m["headers"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}
	if v, err := expandJobHttpTargetOAuthToken(c, f.OAuthToken); err != nil {
		return nil, fmt.Errorf("error expanding OAuthToken into oauthToken: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["oauthToken"] = v
	}
	if v, err := expandJobHttpTargetOidcToken(c, f.OidcToken); err != nil {
		return nil, fmt.Errorf("error expanding OidcToken into oidcToken: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["oidcToken"] = v
	}

	return m, nil
}

// flattenJobHttpTarget flattens an instance of JobHttpTarget from a JSON
// response object.
func flattenJobHttpTarget(c *Client, i interface{}) *JobHttpTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHttpTarget{}
	r.Uri = dcl.FlattenString(m["uri"])
	r.HttpMethod = flattenJobHttpTargetHttpMethodEnum(m["httpMethod"])
	r.Headers = dcl.FlattenKeyValuePairs(m["headers"])
	r.Body = dcl.FlattenString(m["body"])
	r.OAuthToken = flattenJobHttpTargetOAuthToken(c, m["oauthToken"])
	r.OidcToken = flattenJobHttpTargetOidcToken(c, m["oidcToken"])

	return r
}

// expandJobHttpTargetOAuthTokenMap expands the contents of JobHttpTargetOAuthToken into a JSON
// request object.
func expandJobHttpTargetOAuthTokenMap(c *Client, f map[string]JobHttpTargetOAuthToken) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHttpTargetOAuthToken(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHttpTargetOAuthTokenSlice expands the contents of JobHttpTargetOAuthToken into a JSON
// request object.
func expandJobHttpTargetOAuthTokenSlice(c *Client, f []JobHttpTargetOAuthToken) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHttpTargetOAuthToken(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHttpTargetOAuthTokenMap flattens the contents of JobHttpTargetOAuthToken from a JSON
// response object.
func flattenJobHttpTargetOAuthTokenMap(c *Client, i interface{}) map[string]JobHttpTargetOAuthToken {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHttpTargetOAuthToken{}
	}

	if len(a) == 0 {
		return map[string]JobHttpTargetOAuthToken{}
	}

	items := make(map[string]JobHttpTargetOAuthToken)
	for k, item := range a {
		items[k] = *flattenJobHttpTargetOAuthToken(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobHttpTargetOAuthTokenSlice flattens the contents of JobHttpTargetOAuthToken from a JSON
// response object.
func flattenJobHttpTargetOAuthTokenSlice(c *Client, i interface{}) []JobHttpTargetOAuthToken {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHttpTargetOAuthToken{}
	}

	if len(a) == 0 {
		return []JobHttpTargetOAuthToken{}
	}

	items := make([]JobHttpTargetOAuthToken, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHttpTargetOAuthToken(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobHttpTargetOAuthToken expands an instance of JobHttpTargetOAuthToken into a JSON
// request object.
func expandJobHttpTargetOAuthToken(c *Client, f *JobHttpTargetOAuthToken) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmail"] = v
	}
	if v := f.Scope; !dcl.IsEmptyValueIndirect(v) {
		m["scope"] = v
	}

	return m, nil
}

// flattenJobHttpTargetOAuthToken flattens an instance of JobHttpTargetOAuthToken from a JSON
// response object.
func flattenJobHttpTargetOAuthToken(c *Client, i interface{}) *JobHttpTargetOAuthToken {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHttpTargetOAuthToken{}
	r.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])
	r.Scope = dcl.FlattenString(m["scope"])

	return r
}

// expandJobHttpTargetOidcTokenMap expands the contents of JobHttpTargetOidcToken into a JSON
// request object.
func expandJobHttpTargetOidcTokenMap(c *Client, f map[string]JobHttpTargetOidcToken) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHttpTargetOidcToken(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHttpTargetOidcTokenSlice expands the contents of JobHttpTargetOidcToken into a JSON
// request object.
func expandJobHttpTargetOidcTokenSlice(c *Client, f []JobHttpTargetOidcToken) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHttpTargetOidcToken(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHttpTargetOidcTokenMap flattens the contents of JobHttpTargetOidcToken from a JSON
// response object.
func flattenJobHttpTargetOidcTokenMap(c *Client, i interface{}) map[string]JobHttpTargetOidcToken {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHttpTargetOidcToken{}
	}

	if len(a) == 0 {
		return map[string]JobHttpTargetOidcToken{}
	}

	items := make(map[string]JobHttpTargetOidcToken)
	for k, item := range a {
		items[k] = *flattenJobHttpTargetOidcToken(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobHttpTargetOidcTokenSlice flattens the contents of JobHttpTargetOidcToken from a JSON
// response object.
func flattenJobHttpTargetOidcTokenSlice(c *Client, i interface{}) []JobHttpTargetOidcToken {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHttpTargetOidcToken{}
	}

	if len(a) == 0 {
		return []JobHttpTargetOidcToken{}
	}

	items := make([]JobHttpTargetOidcToken, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHttpTargetOidcToken(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobHttpTargetOidcToken expands an instance of JobHttpTargetOidcToken into a JSON
// request object.
func expandJobHttpTargetOidcToken(c *Client, f *JobHttpTargetOidcToken) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmail"] = v
	}
	if v := f.Audience; !dcl.IsEmptyValueIndirect(v) {
		m["audience"] = v
	}

	return m, nil
}

// flattenJobHttpTargetOidcToken flattens an instance of JobHttpTargetOidcToken from a JSON
// response object.
func flattenJobHttpTargetOidcToken(c *Client, i interface{}) *JobHttpTargetOidcToken {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHttpTargetOidcToken{}
	r.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])
	r.Audience = dcl.FlattenString(m["audience"])

	return r
}

// expandJobStatusMap expands the contents of JobStatus into a JSON
// request object.
func expandJobStatusMap(c *Client, f map[string]JobStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobStatus(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobStatusSlice expands the contents of JobStatus into a JSON
// request object.
func expandJobStatusSlice(c *Client, f []JobStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobStatus(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobStatusMap flattens the contents of JobStatus from a JSON
// response object.
func flattenJobStatusMap(c *Client, i interface{}) map[string]JobStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobStatus{}
	}

	if len(a) == 0 {
		return map[string]JobStatus{}
	}

	items := make(map[string]JobStatus)
	for k, item := range a {
		items[k] = *flattenJobStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobStatusSlice flattens the contents of JobStatus from a JSON
// response object.
func flattenJobStatusSlice(c *Client, i interface{}) []JobStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStatus{}
	}

	if len(a) == 0 {
		return []JobStatus{}
	}

	items := make([]JobStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobStatus expands an instance of JobStatus into a JSON
// request object.
func expandJobStatus(c *Client, f *JobStatus) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v, err := expandJobStatusDetailsSlice(c, f.Details); err != nil {
		return nil, fmt.Errorf("error expanding Details into details: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["details"] = v
	}

	return m, nil
}

// flattenJobStatus flattens an instance of JobStatus from a JSON
// response object.
func flattenJobStatus(c *Client, i interface{}) *JobStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobStatus{}
	r.Code = dcl.FlattenInteger(m["code"])
	r.Message = dcl.FlattenString(m["message"])
	r.Details = flattenJobStatusDetailsSlice(c, m["details"])

	return r
}

// expandJobStatusDetailsMap expands the contents of JobStatusDetails into a JSON
// request object.
func expandJobStatusDetailsMap(c *Client, f map[string]JobStatusDetails) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobStatusDetails(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobStatusDetailsSlice expands the contents of JobStatusDetails into a JSON
// request object.
func expandJobStatusDetailsSlice(c *Client, f []JobStatusDetails) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobStatusDetails(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobStatusDetailsMap flattens the contents of JobStatusDetails from a JSON
// response object.
func flattenJobStatusDetailsMap(c *Client, i interface{}) map[string]JobStatusDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobStatusDetails{}
	}

	if len(a) == 0 {
		return map[string]JobStatusDetails{}
	}

	items := make(map[string]JobStatusDetails)
	for k, item := range a {
		items[k] = *flattenJobStatusDetails(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobStatusDetailsSlice flattens the contents of JobStatusDetails from a JSON
// response object.
func flattenJobStatusDetailsSlice(c *Client, i interface{}) []JobStatusDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStatusDetails{}
	}

	if len(a) == 0 {
		return []JobStatusDetails{}
	}

	items := make([]JobStatusDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStatusDetails(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobStatusDetails expands an instance of JobStatusDetails into a JSON
// request object.
func expandJobStatusDetails(c *Client, f *JobStatusDetails) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.TypeUrl; !dcl.IsEmptyValueIndirect(v) {
		m["typeUrl"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenJobStatusDetails flattens an instance of JobStatusDetails from a JSON
// response object.
func flattenJobStatusDetails(c *Client, i interface{}) *JobStatusDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobStatusDetails{}
	r.TypeUrl = dcl.FlattenString(m["typeUrl"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandJobRetryConfigMap expands the contents of JobRetryConfig into a JSON
// request object.
func expandJobRetryConfigMap(c *Client, f map[string]JobRetryConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobRetryConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobRetryConfigSlice expands the contents of JobRetryConfig into a JSON
// request object.
func expandJobRetryConfigSlice(c *Client, f []JobRetryConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobRetryConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobRetryConfigMap flattens the contents of JobRetryConfig from a JSON
// response object.
func flattenJobRetryConfigMap(c *Client, i interface{}) map[string]JobRetryConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobRetryConfig{}
	}

	if len(a) == 0 {
		return map[string]JobRetryConfig{}
	}

	items := make(map[string]JobRetryConfig)
	for k, item := range a {
		items[k] = *flattenJobRetryConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobRetryConfigSlice flattens the contents of JobRetryConfig from a JSON
// response object.
func flattenJobRetryConfigSlice(c *Client, i interface{}) []JobRetryConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobRetryConfig{}
	}

	if len(a) == 0 {
		return []JobRetryConfig{}
	}

	items := make([]JobRetryConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobRetryConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobRetryConfig expands an instance of JobRetryConfig into a JSON
// request object.
func expandJobRetryConfig(c *Client, f *JobRetryConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.RetryCount; !dcl.IsEmptyValueIndirect(v) {
		m["retryCount"] = v
	}
	if v := f.MaxRetryDuration; !dcl.IsEmptyValueIndirect(v) {
		m["maxRetryDuration"] = v
	}
	if v := f.MinBackoffDuration; !dcl.IsEmptyValueIndirect(v) {
		m["minBackoffDuration"] = v
	}
	if v := f.MaxBackoffDuration; !dcl.IsEmptyValueIndirect(v) {
		m["maxBackoffDuration"] = v
	}
	if v := f.MaxDoublings; !dcl.IsEmptyValueIndirect(v) {
		m["maxDoublings"] = v
	}

	return m, nil
}

// flattenJobRetryConfig flattens an instance of JobRetryConfig from a JSON
// response object.
func flattenJobRetryConfig(c *Client, i interface{}) *JobRetryConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobRetryConfig{}
	r.RetryCount = dcl.FlattenInteger(m["retryCount"])
	r.MaxRetryDuration = dcl.FlattenString(m["maxRetryDuration"])
	r.MinBackoffDuration = dcl.FlattenString(m["minBackoffDuration"])
	r.MaxBackoffDuration = dcl.FlattenString(m["maxBackoffDuration"])
	r.MaxDoublings = dcl.FlattenInteger(m["maxDoublings"])

	return r
}

// flattenJobAppEngineHttpTargetHttpMethodEnumSlice flattens the contents of JobAppEngineHttpTargetHttpMethodEnum from a JSON
// response object.
func flattenJobAppEngineHttpTargetHttpMethodEnumSlice(c *Client, i interface{}) []JobAppEngineHttpTargetHttpMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobAppEngineHttpTargetHttpMethodEnum{}
	}

	if len(a) == 0 {
		return []JobAppEngineHttpTargetHttpMethodEnum{}
	}

	items := make([]JobAppEngineHttpTargetHttpMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobAppEngineHttpTargetHttpMethodEnum(item.(interface{})))
	}

	return items
}

// flattenJobAppEngineHttpTargetHttpMethodEnum asserts that an interface is a string, and returns a
// pointer to a *JobAppEngineHttpTargetHttpMethodEnum with the same value as that string.
func flattenJobAppEngineHttpTargetHttpMethodEnum(i interface{}) *JobAppEngineHttpTargetHttpMethodEnum {
	s, ok := i.(string)
	if !ok {
		return JobAppEngineHttpTargetHttpMethodEnumRef("")
	}

	return JobAppEngineHttpTargetHttpMethodEnumRef(s)
}

// flattenJobHttpTargetHttpMethodEnumSlice flattens the contents of JobHttpTargetHttpMethodEnum from a JSON
// response object.
func flattenJobHttpTargetHttpMethodEnumSlice(c *Client, i interface{}) []JobHttpTargetHttpMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHttpTargetHttpMethodEnum{}
	}

	if len(a) == 0 {
		return []JobHttpTargetHttpMethodEnum{}
	}

	items := make([]JobHttpTargetHttpMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHttpTargetHttpMethodEnum(item.(interface{})))
	}

	return items
}

// flattenJobHttpTargetHttpMethodEnum asserts that an interface is a string, and returns a
// pointer to a *JobHttpTargetHttpMethodEnum with the same value as that string.
func flattenJobHttpTargetHttpMethodEnum(i interface{}) *JobHttpTargetHttpMethodEnum {
	s, ok := i.(string)
	if !ok {
		return JobHttpTargetHttpMethodEnumRef("")
	}

	return JobHttpTargetHttpMethodEnumRef(s)
}

// flattenJobStateEnumSlice flattens the contents of JobStateEnum from a JSON
// response object.
func flattenJobStateEnumSlice(c *Client, i interface{}) []JobStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStateEnum{}
	}

	if len(a) == 0 {
		return []JobStateEnum{}
	}

	items := make([]JobStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStateEnum(item.(interface{})))
	}

	return items
}

// flattenJobStateEnum asserts that an interface is a string, and returns a
// pointer to a *JobStateEnum with the same value as that string.
func flattenJobStateEnum(i interface{}) *JobStateEnum {
	s, ok := i.(string)
	if !ok {
		return JobStateEnumRef("")
	}

	return JobStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Job) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalJob(b, c)
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
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
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
