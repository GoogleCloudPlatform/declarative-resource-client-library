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
	"github.com/mohae/deepcopy"
	"io/ioutil"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"reflect"
	"strings"
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
	if !dcl.IsEmptyValueIndirect(r.AttemptDeadline) {
		if err := r.AttemptDeadline.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobLabels) validate() error {
	return nil
}
func (r *JobPubsubTarget) validate() error {
	if err := dcl.Required(r, "topicName"); err != nil {
		return err
	}
	return nil
}
func (r *JobPubsubTargetAttributes) validate() error {
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
func (r *JobAppEngineHttpTargetHeaders) validate() error {
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
func (r *JobHttpTargetHeaders) validate() error {
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
	if !dcl.IsEmptyValueIndirect(r.MaxRetryDuration) {
		if err := r.MaxRetryDuration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MinBackoffDuration) {
		if err := r.MinBackoffDuration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaxBackoffDuration) {
		if err := r.MaxBackoffDuration.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobRetryConfigMaxRetryDuration) validate() error {
	return nil
}
func (r *JobRetryConfigMinBackoffDuration) validate() error {
	return nil
}
func (r *JobRetryConfigMaxBackoffDuration) validate() error {
	return nil
}
func (r *JobAttemptDeadline) validate() error {
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
	if v, err := expandJobLabelsSlice(c, f.Labels); err != nil {
		return nil, fmt.Errorf("error expanding Labels into labels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
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
	if v, err := expandJobAttemptDeadline(c, f.AttemptDeadline); err != nil {
		return nil, fmt.Errorf("error expanding AttemptDeadline into attemptDeadline: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.Retry)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listJobOperation struct {
	Items []map[string]interface{} `json:"items"`
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
	for _, v := range m.Items {
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
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return fmt.Errorf("failed to delete Job: %w", err)
	}
	_, err = c.GetJob(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createJobOperation struct{}

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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	_ = o // We might not use resp- this will stop Go complaining

	if _, err := c.GetJob(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getJobRaw(ctx context.Context, r *Job) ([]byte, error) {

	u, err := jobGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Job); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected Job, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.PubsubTarget = canonicalizeJobPubsubTarget(rawDesired.PubsubTarget, nil, opts...)
		rawDesired.AppEngineHttpTarget = canonicalizeJobAppEngineHttpTarget(rawDesired.AppEngineHttpTarget, nil, opts...)
		rawDesired.HttpTarget = canonicalizeJobHttpTarget(rawDesired.HttpTarget, nil, opts...)
		rawDesired.Status = canonicalizeJobStatus(rawDesired.Status, nil, opts...)
		rawDesired.RetryConfig = canonicalizeJobRetryConfig(rawDesired.RetryConfig, nil, opts...)
		rawDesired.AttemptDeadline = canonicalizeJobAttemptDeadline(rawDesired.AttemptDeadline, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	rawDesired.PubsubTarget = canonicalizeJobPubsubTarget(rawDesired.PubsubTarget, rawInitial.PubsubTarget, opts...)
	rawDesired.AppEngineHttpTarget = canonicalizeJobAppEngineHttpTarget(rawDesired.AppEngineHttpTarget, rawInitial.AppEngineHttpTarget, opts...)
	rawDesired.HttpTarget = canonicalizeJobHttpTarget(rawDesired.HttpTarget, rawInitial.HttpTarget, opts...)
	if dcl.IsZeroValue(rawDesired.Schedule) {
		rawDesired.Schedule = rawInitial.Schedule
	}
	if dcl.IsZeroValue(rawDesired.TimeZone) {
		rawDesired.TimeZone = rawInitial.TimeZone
	}
	if dcl.IsZeroValue(rawDesired.UserUpdateTime) {
		rawDesired.UserUpdateTime = rawInitial.UserUpdateTime
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	rawDesired.Status = canonicalizeJobStatus(rawDesired.Status, rawInitial.Status, opts...)
	if dcl.IsZeroValue(rawDesired.TotalAttemptCount) {
		rawDesired.TotalAttemptCount = rawInitial.TotalAttemptCount
	}
	if dcl.IsZeroValue(rawDesired.FailedAttemptCount) {
		rawDesired.FailedAttemptCount = rawInitial.FailedAttemptCount
	}
	if dcl.IsZeroValue(rawDesired.TotalExecutionCount) {
		rawDesired.TotalExecutionCount = rawInitial.TotalExecutionCount
	}
	if dcl.IsZeroValue(rawDesired.FailedExecutionCount) {
		rawDesired.FailedExecutionCount = rawInitial.FailedExecutionCount
	}
	if dcl.IsZeroValue(rawDesired.View) {
		rawDesired.View = rawInitial.View
	}
	if dcl.IsZeroValue(rawDesired.ScheduleTime) {
		rawDesired.ScheduleTime = rawInitial.ScheduleTime
	}
	if dcl.IsZeroValue(rawDesired.LastAttemptTime) {
		rawDesired.LastAttemptTime = rawInitial.LastAttemptTime
	}
	rawDesired.RetryConfig = canonicalizeJobRetryConfig(rawDesired.RetryConfig, rawInitial.RetryConfig, opts...)
	rawDesired.AttemptDeadline = canonicalizeJobAttemptDeadline(rawDesired.AttemptDeadline, rawInitial.AttemptDeadline, opts...)
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeZone) && dcl.IsEmptyValueIndirect(rawDesired.TimeZone) {
		rawNew.TimeZone = rawDesired.TimeZone
	} else {
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

	if dcl.IsEmptyValueIndirect(rawNew.TotalAttemptCount) && dcl.IsEmptyValueIndirect(rawDesired.TotalAttemptCount) {
		rawNew.TotalAttemptCount = rawDesired.TotalAttemptCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.FailedAttemptCount) && dcl.IsEmptyValueIndirect(rawDesired.FailedAttemptCount) {
		rawNew.FailedAttemptCount = rawDesired.FailedAttemptCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TotalExecutionCount) && dcl.IsEmptyValueIndirect(rawDesired.TotalExecutionCount) {
		rawNew.TotalExecutionCount = rawDesired.TotalExecutionCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.FailedExecutionCount) && dcl.IsEmptyValueIndirect(rawDesired.FailedExecutionCount) {
		rawNew.FailedExecutionCount = rawDesired.FailedExecutionCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.View) && dcl.IsEmptyValueIndirect(rawDesired.View) {
		rawNew.View = rawDesired.View
	} else {
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
		rawNew.AttemptDeadline = canonicalizeNewJobAttemptDeadline(c, rawDesired.AttemptDeadline, rawNew.AttemptDeadline)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.NameToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Location) && dcl.IsEmptyValueIndirect(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
		if dcl.NameToSelfLink(rawDesired.Location, rawNew.Location) {
			rawNew.Location = rawDesired.Location
		}
	}

	return rawNew, nil
}

func canonicalizeJobLabels(des, initial *JobLabels, opts ...dcl.ApplyOption) *JobLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewJobLabels(c *Client, des, nw *JobLabels) *JobLabels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobLabelsSet(c *Client, des, nw []JobLabels) []JobLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []JobLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobLabels(c, &d, &n) {
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

func canonicalizeJobPubsubTarget(des, initial *JobPubsubTarget, opts ...dcl.ApplyOption) *JobPubsubTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.TopicName, initial.TopicName) || dcl.IsZeroValue(des.TopicName) {
		des.TopicName = initial.TopicName
	}
	if dcl.IsZeroValue(des.Data) {
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

	if dcl.NameToSelfLink(des.TopicName, nw.TopicName) || dcl.IsZeroValue(des.TopicName) {
		nw.TopicName = des.TopicName
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

func canonicalizeJobPubsubTargetAttributes(des, initial *JobPubsubTargetAttributes, opts ...dcl.ApplyOption) *JobPubsubTargetAttributes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewJobPubsubTargetAttributes(c *Client, des, nw *JobPubsubTargetAttributes) *JobPubsubTargetAttributes {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobPubsubTargetAttributesSet(c *Client, des, nw []JobPubsubTargetAttributes) []JobPubsubTargetAttributes {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPubsubTargetAttributes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPubsubTargetAttributes(c, &d, &n) {
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

func canonicalizeJobAppEngineHttpTarget(des, initial *JobAppEngineHttpTarget, opts ...dcl.ApplyOption) *JobAppEngineHttpTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HttpMethod) {
		des.HttpMethod = initial.HttpMethod
	}
	des.AppEngineRouting = canonicalizeJobAppEngineHttpTargetAppEngineRouting(des.AppEngineRouting, initial.AppEngineRouting, opts...)
	if dcl.IsZeroValue(des.RelativeUri) {
		des.RelativeUri = initial.RelativeUri
	}
	if dcl.IsZeroValue(des.Headers) {
		des.Headers = initial.Headers
	}
	if dcl.IsZeroValue(des.Body) {
		des.Body = initial.Body
	}

	return des
}

func canonicalizeNewJobAppEngineHttpTarget(c *Client, des, nw *JobAppEngineHttpTarget) *JobAppEngineHttpTarget {
	if des == nil || nw == nil {
		return nw
	}

	nw.AppEngineRouting = canonicalizeNewJobAppEngineHttpTargetAppEngineRouting(c, des.AppEngineRouting, nw.AppEngineRouting)

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

func canonicalizeJobAppEngineHttpTargetAppEngineRouting(des, initial *JobAppEngineHttpTargetAppEngineRouting, opts ...dcl.ApplyOption) *JobAppEngineHttpTargetAppEngineRouting {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Service) {
		des.Service = initial.Service
	}
	if dcl.IsZeroValue(des.Version) {
		des.Version = initial.Version
	}
	if dcl.IsZeroValue(des.Instance) {
		des.Instance = initial.Instance
	}
	if dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}

	return des
}

func canonicalizeNewJobAppEngineHttpTargetAppEngineRouting(c *Client, des, nw *JobAppEngineHttpTargetAppEngineRouting) *JobAppEngineHttpTargetAppEngineRouting {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeJobAppEngineHttpTargetHeaders(des, initial *JobAppEngineHttpTargetHeaders, opts ...dcl.ApplyOption) *JobAppEngineHttpTargetHeaders {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewJobAppEngineHttpTargetHeaders(c *Client, des, nw *JobAppEngineHttpTargetHeaders) *JobAppEngineHttpTargetHeaders {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobAppEngineHttpTargetHeadersSet(c *Client, des, nw []JobAppEngineHttpTargetHeaders) []JobAppEngineHttpTargetHeaders {
	if des == nil {
		return nw
	}
	var reorderedNew []JobAppEngineHttpTargetHeaders
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobAppEngineHttpTargetHeaders(c, &d, &n) {
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

func canonicalizeJobHttpTarget(des, initial *JobHttpTarget, opts ...dcl.ApplyOption) *JobHttpTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Uri) {
		des.Uri = initial.Uri
	}
	if dcl.IsZeroValue(des.HttpMethod) {
		des.HttpMethod = initial.HttpMethod
	}
	if dcl.IsZeroValue(des.Headers) {
		des.Headers = initial.Headers
	}
	if dcl.IsZeroValue(des.Body) {
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

func canonicalizeJobHttpTargetHeaders(des, initial *JobHttpTargetHeaders, opts ...dcl.ApplyOption) *JobHttpTargetHeaders {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewJobHttpTargetHeaders(c *Client, des, nw *JobHttpTargetHeaders) *JobHttpTargetHeaders {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobHttpTargetHeadersSet(c *Client, des, nw []JobHttpTargetHeaders) []JobHttpTargetHeaders {
	if des == nil {
		return nw
	}
	var reorderedNew []JobHttpTargetHeaders
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobHttpTargetHeaders(c, &d, &n) {
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

func canonicalizeJobHttpTargetOAuthToken(des, initial *JobHttpTargetOAuthToken, opts ...dcl.ApplyOption) *JobHttpTargetOAuthToken {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ServiceAccountEmail) {
		des.ServiceAccountEmail = initial.ServiceAccountEmail
	}
	if dcl.IsZeroValue(des.Scope) {
		des.Scope = initial.Scope
	}

	return des
}

func canonicalizeNewJobHttpTargetOAuthToken(c *Client, des, nw *JobHttpTargetOAuthToken) *JobHttpTargetOAuthToken {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeJobHttpTargetOidcToken(des, initial *JobHttpTargetOidcToken, opts ...dcl.ApplyOption) *JobHttpTargetOidcToken {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ServiceAccountEmail) {
		des.ServiceAccountEmail = initial.ServiceAccountEmail
	}
	if dcl.IsZeroValue(des.Audience) {
		des.Audience = initial.Audience
	}

	return des
}

func canonicalizeNewJobHttpTargetOidcToken(c *Client, des, nw *JobHttpTargetOidcToken) *JobHttpTargetOidcToken {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeJobStatus(des, initial *JobStatus, opts ...dcl.ApplyOption) *JobStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Code) {
		des.Code = initial.Code
	}
	if dcl.IsZeroValue(des.Message) {
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

func canonicalizeJobStatusDetails(des, initial *JobStatusDetails, opts ...dcl.ApplyOption) *JobStatusDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.TypeUrl) {
		des.TypeUrl = initial.TypeUrl
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewJobStatusDetails(c *Client, des, nw *JobStatusDetails) *JobStatusDetails {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeJobRetryConfig(des, initial *JobRetryConfig, opts ...dcl.ApplyOption) *JobRetryConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RetryCount) {
		des.RetryCount = initial.RetryCount
	}
	des.MaxRetryDuration = canonicalizeJobRetryConfigMaxRetryDuration(des.MaxRetryDuration, initial.MaxRetryDuration, opts...)
	des.MinBackoffDuration = canonicalizeJobRetryConfigMinBackoffDuration(des.MinBackoffDuration, initial.MinBackoffDuration, opts...)
	des.MaxBackoffDuration = canonicalizeJobRetryConfigMaxBackoffDuration(des.MaxBackoffDuration, initial.MaxBackoffDuration, opts...)
	if dcl.IsZeroValue(des.MaxDoublings) {
		des.MaxDoublings = initial.MaxDoublings
	}

	return des
}

func canonicalizeNewJobRetryConfig(c *Client, des, nw *JobRetryConfig) *JobRetryConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.MaxRetryDuration = canonicalizeNewJobRetryConfigMaxRetryDuration(c, des.MaxRetryDuration, nw.MaxRetryDuration)
	nw.MinBackoffDuration = canonicalizeNewJobRetryConfigMinBackoffDuration(c, des.MinBackoffDuration, nw.MinBackoffDuration)
	nw.MaxBackoffDuration = canonicalizeNewJobRetryConfigMaxBackoffDuration(c, des.MaxBackoffDuration, nw.MaxBackoffDuration)

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

func canonicalizeJobRetryConfigMaxRetryDuration(des, initial *JobRetryConfigMaxRetryDuration, opts ...dcl.ApplyOption) *JobRetryConfigMaxRetryDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewJobRetryConfigMaxRetryDuration(c *Client, des, nw *JobRetryConfigMaxRetryDuration) *JobRetryConfigMaxRetryDuration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobRetryConfigMaxRetryDurationSet(c *Client, des, nw []JobRetryConfigMaxRetryDuration) []JobRetryConfigMaxRetryDuration {
	if des == nil {
		return nw
	}
	var reorderedNew []JobRetryConfigMaxRetryDuration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobRetryConfigMaxRetryDuration(c, &d, &n) {
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

func canonicalizeJobRetryConfigMinBackoffDuration(des, initial *JobRetryConfigMinBackoffDuration, opts ...dcl.ApplyOption) *JobRetryConfigMinBackoffDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewJobRetryConfigMinBackoffDuration(c *Client, des, nw *JobRetryConfigMinBackoffDuration) *JobRetryConfigMinBackoffDuration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobRetryConfigMinBackoffDurationSet(c *Client, des, nw []JobRetryConfigMinBackoffDuration) []JobRetryConfigMinBackoffDuration {
	if des == nil {
		return nw
	}
	var reorderedNew []JobRetryConfigMinBackoffDuration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobRetryConfigMinBackoffDuration(c, &d, &n) {
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

func canonicalizeJobRetryConfigMaxBackoffDuration(des, initial *JobRetryConfigMaxBackoffDuration, opts ...dcl.ApplyOption) *JobRetryConfigMaxBackoffDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewJobRetryConfigMaxBackoffDuration(c *Client, des, nw *JobRetryConfigMaxBackoffDuration) *JobRetryConfigMaxBackoffDuration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobRetryConfigMaxBackoffDurationSet(c *Client, des, nw []JobRetryConfigMaxBackoffDuration) []JobRetryConfigMaxBackoffDuration {
	if des == nil {
		return nw
	}
	var reorderedNew []JobRetryConfigMaxBackoffDuration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobRetryConfigMaxBackoffDuration(c, &d, &n) {
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

func canonicalizeJobAttemptDeadline(des, initial *JobAttemptDeadline, opts ...dcl.ApplyOption) *JobAttemptDeadline {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Job)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewJobAttemptDeadline(c *Client, des, nw *JobAttemptDeadline) *JobAttemptDeadline {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobAttemptDeadlineSet(c *Client, des, nw []JobAttemptDeadline) []JobAttemptDeadline {
	if des == nil {
		return nw
	}
	var reorderedNew []JobAttemptDeadline
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobAttemptDeadline(c, &d, &n) {
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

type jobDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         jobApiOperation
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
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %#v\nACTUAL: %#v", desired.Name, actual.Name)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "Name",
		})

	}
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %#v\nACTUAL: %#v", desired.Description, actual.Description)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "Description",
		})

	}
	if compareJobLabelsSlice(c, desired.Labels, actual.Labels) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %#v\nACTUAL: %#v", desired.Labels, actual.Labels)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "Labels",
		})

	}
	if compareJobPubsubTarget(c, desired.PubsubTarget, actual.PubsubTarget) {
		c.Config.Logger.Infof("Detected diff in PubsubTarget.\nDESIRED: %#v\nACTUAL: %#v", desired.PubsubTarget, actual.PubsubTarget)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "PubsubTarget",
		})

	}
	if compareJobAppEngineHttpTarget(c, desired.AppEngineHttpTarget, actual.AppEngineHttpTarget) {
		c.Config.Logger.Infof("Detected diff in AppEngineHttpTarget.\nDESIRED: %#v\nACTUAL: %#v", desired.AppEngineHttpTarget, actual.AppEngineHttpTarget)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "AppEngineHttpTarget",
		})

	}
	if compareJobHttpTarget(c, desired.HttpTarget, actual.HttpTarget) {
		c.Config.Logger.Infof("Detected diff in HttpTarget.\nDESIRED: %#v\nACTUAL: %#v", desired.HttpTarget, actual.HttpTarget)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "HttpTarget",
		})

	}
	if !dcl.IsZeroValue(desired.Schedule) && (dcl.IsZeroValue(actual.Schedule) || !reflect.DeepEqual(*desired.Schedule, *actual.Schedule)) {
		c.Config.Logger.Infof("Detected diff in Schedule.\nDESIRED: %#v\nACTUAL: %#v", desired.Schedule, actual.Schedule)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "Schedule",
		})

	}
	if !dcl.IsZeroValue(desired.TimeZone) && (dcl.IsZeroValue(actual.TimeZone) || !reflect.DeepEqual(*desired.TimeZone, *actual.TimeZone)) {
		c.Config.Logger.Infof("Detected diff in TimeZone.\nDESIRED: %#v\nACTUAL: %#v", desired.TimeZone, actual.TimeZone)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "TimeZone",
		})

	}
	if compareJobRetryConfig(c, desired.RetryConfig, actual.RetryConfig) {
		c.Config.Logger.Infof("Detected diff in RetryConfig.\nDESIRED: %#v\nACTUAL: %#v", desired.RetryConfig, actual.RetryConfig)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "RetryConfig",
		})

	}
	if compareJobAttemptDeadline(c, desired.AttemptDeadline, actual.AttemptDeadline) {
		c.Config.Logger.Infof("Detected diff in AttemptDeadline.\nDESIRED: %#v\nACTUAL: %#v", desired.AttemptDeadline, actual.AttemptDeadline)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobUpdateJobOperation{},
			FieldName: "AttemptDeadline",
		})

	}
	if !dcl.IsZeroValue(desired.Project) && !dcl.NameToSelfLink(desired.Project, actual.Project) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %#v\nACTUAL: %#v", desired.Project, actual.Project)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "Project",
		})
	}
	if !dcl.IsZeroValue(desired.Location) && !dcl.NameToSelfLink(desired.Location, actual.Location) {
		c.Config.Logger.Infof("Detected diff in Location.\nDESIRED: %#v\nACTUAL: %#v", desired.Location, actual.Location)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "Location",
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
func compareJobLabelsSlice(c *Client, desired, actual []JobLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobLabels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobLabels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobLabels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobLabels(c *Client, desired, actual *JobLabels) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !reflect.DeepEqual(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) && !(dcl.IsEmptyValueIndirect(desired.Key) && dcl.IsZeroValue(actual.Key)) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
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
			c.Config.Logger.Infof("Diff in JobPubsubTarget, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPubsubTarget(c *Client, desired, actual *JobPubsubTarget) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.TopicName == nil && desired.TopicName != nil && !dcl.IsEmptyValueIndirect(desired.TopicName) {
		c.Config.Logger.Infof("desired TopicName %s - but actually nil", dcl.SprintResource(desired.TopicName))
		return true
	}
	if !dcl.NameToSelfLink(desired.TopicName, actual.TopicName) && !dcl.IsZeroValue(desired.TopicName) {
		c.Config.Logger.Infof("Diff in TopicName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TopicName), dcl.SprintResource(actual.TopicName))
		return true
	}
	if actual.Data == nil && desired.Data != nil && !dcl.IsEmptyValueIndirect(desired.Data) {
		c.Config.Logger.Infof("desired Data %s - but actually nil", dcl.SprintResource(desired.Data))
		return true
	}
	if !reflect.DeepEqual(desired.Data, actual.Data) && !dcl.IsZeroValue(desired.Data) && !(dcl.IsEmptyValueIndirect(desired.Data) && dcl.IsZeroValue(actual.Data)) {
		c.Config.Logger.Infof("Diff in Data. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Data), dcl.SprintResource(actual.Data))
		return true
	}
	if actual.Attributes == nil && desired.Attributes != nil && !dcl.IsEmptyValueIndirect(desired.Attributes) {
		c.Config.Logger.Infof("desired Attributes %s - but actually nil", dcl.SprintResource(desired.Attributes))
		return true
	}
	if compareJobPubsubTargetAttributesSlice(c, desired.Attributes, actual.Attributes) && !dcl.IsZeroValue(desired.Attributes) {
		c.Config.Logger.Infof("Diff in Attributes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Attributes), dcl.SprintResource(actual.Attributes))
		return true
	}
	return false
}
func compareJobPubsubTargetAttributesSlice(c *Client, desired, actual []JobPubsubTargetAttributes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPubsubTargetAttributes, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPubsubTargetAttributes(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPubsubTargetAttributes, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPubsubTargetAttributes(c *Client, desired, actual *JobPubsubTargetAttributes) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !reflect.DeepEqual(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) && !(dcl.IsEmptyValueIndirect(desired.Key) && dcl.IsZeroValue(actual.Key)) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
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
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTarget, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobAppEngineHttpTarget(c *Client, desired, actual *JobAppEngineHttpTarget) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HttpMethod == nil && desired.HttpMethod != nil && !dcl.IsEmptyValueIndirect(desired.HttpMethod) {
		c.Config.Logger.Infof("desired HttpMethod %s - but actually nil", dcl.SprintResource(desired.HttpMethod))
		return true
	}
	if !reflect.DeepEqual(desired.HttpMethod, actual.HttpMethod) && !dcl.IsZeroValue(desired.HttpMethod) && !(dcl.IsEmptyValueIndirect(desired.HttpMethod) && dcl.IsZeroValue(actual.HttpMethod)) {
		c.Config.Logger.Infof("Diff in HttpMethod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpMethod), dcl.SprintResource(actual.HttpMethod))
		return true
	}
	if actual.AppEngineRouting == nil && desired.AppEngineRouting != nil && !dcl.IsEmptyValueIndirect(desired.AppEngineRouting) {
		c.Config.Logger.Infof("desired AppEngineRouting %s - but actually nil", dcl.SprintResource(desired.AppEngineRouting))
		return true
	}
	if compareJobAppEngineHttpTargetAppEngineRouting(c, desired.AppEngineRouting, actual.AppEngineRouting) && !dcl.IsZeroValue(desired.AppEngineRouting) {
		c.Config.Logger.Infof("Diff in AppEngineRouting. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AppEngineRouting), dcl.SprintResource(actual.AppEngineRouting))
		return true
	}
	if actual.RelativeUri == nil && desired.RelativeUri != nil && !dcl.IsEmptyValueIndirect(desired.RelativeUri) {
		c.Config.Logger.Infof("desired RelativeUri %s - but actually nil", dcl.SprintResource(desired.RelativeUri))
		return true
	}
	if !reflect.DeepEqual(desired.RelativeUri, actual.RelativeUri) && !dcl.IsZeroValue(desired.RelativeUri) && !(dcl.IsEmptyValueIndirect(desired.RelativeUri) && dcl.IsZeroValue(actual.RelativeUri)) {
		c.Config.Logger.Infof("Diff in RelativeUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RelativeUri), dcl.SprintResource(actual.RelativeUri))
		return true
	}
	if actual.Headers == nil && desired.Headers != nil && !dcl.IsEmptyValueIndirect(desired.Headers) {
		c.Config.Logger.Infof("desired Headers %s - but actually nil", dcl.SprintResource(desired.Headers))
		return true
	}
	if compareJobAppEngineHttpTargetHeadersSlice(c, desired.Headers, actual.Headers) && !dcl.IsZeroValue(desired.Headers) {
		c.Config.Logger.Infof("Diff in Headers. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Headers), dcl.SprintResource(actual.Headers))
		return true
	}
	if actual.Body == nil && desired.Body != nil && !dcl.IsEmptyValueIndirect(desired.Body) {
		c.Config.Logger.Infof("desired Body %s - but actually nil", dcl.SprintResource(desired.Body))
		return true
	}
	if !reflect.DeepEqual(desired.Body, actual.Body) && !dcl.IsZeroValue(desired.Body) && !(dcl.IsEmptyValueIndirect(desired.Body) && dcl.IsZeroValue(actual.Body)) {
		c.Config.Logger.Infof("Diff in Body. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Body), dcl.SprintResource(actual.Body))
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
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTargetAppEngineRouting, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobAppEngineHttpTargetAppEngineRouting(c *Client, desired, actual *JobAppEngineHttpTargetAppEngineRouting) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Service == nil && desired.Service != nil && !dcl.IsEmptyValueIndirect(desired.Service) {
		c.Config.Logger.Infof("desired Service %s - but actually nil", dcl.SprintResource(desired.Service))
		return true
	}
	if !reflect.DeepEqual(desired.Service, actual.Service) && !dcl.IsZeroValue(desired.Service) && !(dcl.IsEmptyValueIndirect(desired.Service) && dcl.IsZeroValue(actual.Service)) {
		c.Config.Logger.Infof("Diff in Service. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Service), dcl.SprintResource(actual.Service))
		return true
	}
	if actual.Version == nil && desired.Version != nil && !dcl.IsEmptyValueIndirect(desired.Version) {
		c.Config.Logger.Infof("desired Version %s - but actually nil", dcl.SprintResource(desired.Version))
		return true
	}
	if !reflect.DeepEqual(desired.Version, actual.Version) && !dcl.IsZeroValue(desired.Version) && !(dcl.IsEmptyValueIndirect(desired.Version) && dcl.IsZeroValue(actual.Version)) {
		c.Config.Logger.Infof("Diff in Version. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Version), dcl.SprintResource(actual.Version))
		return true
	}
	if actual.Instance == nil && desired.Instance != nil && !dcl.IsEmptyValueIndirect(desired.Instance) {
		c.Config.Logger.Infof("desired Instance %s - but actually nil", dcl.SprintResource(desired.Instance))
		return true
	}
	if !reflect.DeepEqual(desired.Instance, actual.Instance) && !dcl.IsZeroValue(desired.Instance) && !(dcl.IsEmptyValueIndirect(desired.Instance) && dcl.IsZeroValue(actual.Instance)) {
		c.Config.Logger.Infof("Diff in Instance. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Instance), dcl.SprintResource(actual.Instance))
		return true
	}
	return false
}
func compareJobAppEngineHttpTargetHeadersSlice(c *Client, desired, actual []JobAppEngineHttpTargetHeaders) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobAppEngineHttpTargetHeaders, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobAppEngineHttpTargetHeaders(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTargetHeaders, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobAppEngineHttpTargetHeaders(c *Client, desired, actual *JobAppEngineHttpTargetHeaders) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !reflect.DeepEqual(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) && !(dcl.IsEmptyValueIndirect(desired.Key) && dcl.IsZeroValue(actual.Key)) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
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
			c.Config.Logger.Infof("Diff in JobHttpTarget, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHttpTarget(c *Client, desired, actual *JobHttpTarget) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Uri == nil && desired.Uri != nil && !dcl.IsEmptyValueIndirect(desired.Uri) {
		c.Config.Logger.Infof("desired Uri %s - but actually nil", dcl.SprintResource(desired.Uri))
		return true
	}
	if !reflect.DeepEqual(desired.Uri, actual.Uri) && !dcl.IsZeroValue(desired.Uri) && !(dcl.IsEmptyValueIndirect(desired.Uri) && dcl.IsZeroValue(actual.Uri)) {
		c.Config.Logger.Infof("Diff in Uri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Uri), dcl.SprintResource(actual.Uri))
		return true
	}
	if actual.HttpMethod == nil && desired.HttpMethod != nil && !dcl.IsEmptyValueIndirect(desired.HttpMethod) {
		c.Config.Logger.Infof("desired HttpMethod %s - but actually nil", dcl.SprintResource(desired.HttpMethod))
		return true
	}
	if !reflect.DeepEqual(desired.HttpMethod, actual.HttpMethod) && !dcl.IsZeroValue(desired.HttpMethod) && !(dcl.IsEmptyValueIndirect(desired.HttpMethod) && dcl.IsZeroValue(actual.HttpMethod)) {
		c.Config.Logger.Infof("Diff in HttpMethod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpMethod), dcl.SprintResource(actual.HttpMethod))
		return true
	}
	if actual.Headers == nil && desired.Headers != nil && !dcl.IsEmptyValueIndirect(desired.Headers) {
		c.Config.Logger.Infof("desired Headers %s - but actually nil", dcl.SprintResource(desired.Headers))
		return true
	}
	if compareJobHttpTargetHeadersSlice(c, desired.Headers, actual.Headers) && !dcl.IsZeroValue(desired.Headers) {
		c.Config.Logger.Infof("Diff in Headers. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Headers), dcl.SprintResource(actual.Headers))
		return true
	}
	if actual.Body == nil && desired.Body != nil && !dcl.IsEmptyValueIndirect(desired.Body) {
		c.Config.Logger.Infof("desired Body %s - but actually nil", dcl.SprintResource(desired.Body))
		return true
	}
	if !reflect.DeepEqual(desired.Body, actual.Body) && !dcl.IsZeroValue(desired.Body) && !(dcl.IsEmptyValueIndirect(desired.Body) && dcl.IsZeroValue(actual.Body)) {
		c.Config.Logger.Infof("Diff in Body. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Body), dcl.SprintResource(actual.Body))
		return true
	}
	if actual.OAuthToken == nil && desired.OAuthToken != nil && !dcl.IsEmptyValueIndirect(desired.OAuthToken) {
		c.Config.Logger.Infof("desired OAuthToken %s - but actually nil", dcl.SprintResource(desired.OAuthToken))
		return true
	}
	if compareJobHttpTargetOAuthToken(c, desired.OAuthToken, actual.OAuthToken) && !dcl.IsZeroValue(desired.OAuthToken) {
		c.Config.Logger.Infof("Diff in OAuthToken. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OAuthToken), dcl.SprintResource(actual.OAuthToken))
		return true
	}
	if actual.OidcToken == nil && desired.OidcToken != nil && !dcl.IsEmptyValueIndirect(desired.OidcToken) {
		c.Config.Logger.Infof("desired OidcToken %s - but actually nil", dcl.SprintResource(desired.OidcToken))
		return true
	}
	if compareJobHttpTargetOidcToken(c, desired.OidcToken, actual.OidcToken) && !dcl.IsZeroValue(desired.OidcToken) {
		c.Config.Logger.Infof("Diff in OidcToken. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OidcToken), dcl.SprintResource(actual.OidcToken))
		return true
	}
	return false
}
func compareJobHttpTargetHeadersSlice(c *Client, desired, actual []JobHttpTargetHeaders) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHttpTargetHeaders, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobHttpTargetHeaders(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobHttpTargetHeaders, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHttpTargetHeaders(c *Client, desired, actual *JobHttpTargetHeaders) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !reflect.DeepEqual(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) && !(dcl.IsEmptyValueIndirect(desired.Key) && dcl.IsZeroValue(actual.Key)) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
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
			c.Config.Logger.Infof("Diff in JobHttpTargetOAuthToken, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHttpTargetOAuthToken(c *Client, desired, actual *JobHttpTargetOAuthToken) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ServiceAccountEmail == nil && desired.ServiceAccountEmail != nil && !dcl.IsEmptyValueIndirect(desired.ServiceAccountEmail) {
		c.Config.Logger.Infof("desired ServiceAccountEmail %s - but actually nil", dcl.SprintResource(desired.ServiceAccountEmail))
		return true
	}
	if !reflect.DeepEqual(desired.ServiceAccountEmail, actual.ServiceAccountEmail) && !dcl.IsZeroValue(desired.ServiceAccountEmail) && !(dcl.IsEmptyValueIndirect(desired.ServiceAccountEmail) && dcl.IsZeroValue(actual.ServiceAccountEmail)) {
		c.Config.Logger.Infof("Diff in ServiceAccountEmail. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccountEmail), dcl.SprintResource(actual.ServiceAccountEmail))
		return true
	}
	if actual.Scope == nil && desired.Scope != nil && !dcl.IsEmptyValueIndirect(desired.Scope) {
		c.Config.Logger.Infof("desired Scope %s - but actually nil", dcl.SprintResource(desired.Scope))
		return true
	}
	if !reflect.DeepEqual(desired.Scope, actual.Scope) && !dcl.IsZeroValue(desired.Scope) && !(dcl.IsEmptyValueIndirect(desired.Scope) && dcl.IsZeroValue(actual.Scope)) {
		c.Config.Logger.Infof("Diff in Scope. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scope), dcl.SprintResource(actual.Scope))
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
			c.Config.Logger.Infof("Diff in JobHttpTargetOidcToken, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHttpTargetOidcToken(c *Client, desired, actual *JobHttpTargetOidcToken) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ServiceAccountEmail == nil && desired.ServiceAccountEmail != nil && !dcl.IsEmptyValueIndirect(desired.ServiceAccountEmail) {
		c.Config.Logger.Infof("desired ServiceAccountEmail %s - but actually nil", dcl.SprintResource(desired.ServiceAccountEmail))
		return true
	}
	if !reflect.DeepEqual(desired.ServiceAccountEmail, actual.ServiceAccountEmail) && !dcl.IsZeroValue(desired.ServiceAccountEmail) && !(dcl.IsEmptyValueIndirect(desired.ServiceAccountEmail) && dcl.IsZeroValue(actual.ServiceAccountEmail)) {
		c.Config.Logger.Infof("Diff in ServiceAccountEmail. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccountEmail), dcl.SprintResource(actual.ServiceAccountEmail))
		return true
	}
	if actual.Audience == nil && desired.Audience != nil && !dcl.IsEmptyValueIndirect(desired.Audience) {
		c.Config.Logger.Infof("desired Audience %s - but actually nil", dcl.SprintResource(desired.Audience))
		return true
	}
	if !reflect.DeepEqual(desired.Audience, actual.Audience) && !dcl.IsZeroValue(desired.Audience) && !(dcl.IsEmptyValueIndirect(desired.Audience) && dcl.IsZeroValue(actual.Audience)) {
		c.Config.Logger.Infof("Diff in Audience. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Audience), dcl.SprintResource(actual.Audience))
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
			c.Config.Logger.Infof("Diff in JobStatus, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStatus(c *Client, desired, actual *JobStatus) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Code == nil && desired.Code != nil && !dcl.IsEmptyValueIndirect(desired.Code) {
		c.Config.Logger.Infof("desired Code %s - but actually nil", dcl.SprintResource(desired.Code))
		return true
	}
	if !reflect.DeepEqual(desired.Code, actual.Code) && !dcl.IsZeroValue(desired.Code) && !(dcl.IsEmptyValueIndirect(desired.Code) && dcl.IsZeroValue(actual.Code)) {
		c.Config.Logger.Infof("Diff in Code. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Code), dcl.SprintResource(actual.Code))
		return true
	}
	if actual.Message == nil && desired.Message != nil && !dcl.IsEmptyValueIndirect(desired.Message) {
		c.Config.Logger.Infof("desired Message %s - but actually nil", dcl.SprintResource(desired.Message))
		return true
	}
	if !reflect.DeepEqual(desired.Message, actual.Message) && !dcl.IsZeroValue(desired.Message) && !(dcl.IsEmptyValueIndirect(desired.Message) && dcl.IsZeroValue(actual.Message)) {
		c.Config.Logger.Infof("Diff in Message. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Message), dcl.SprintResource(actual.Message))
		return true
	}
	if actual.Details == nil && desired.Details != nil && !dcl.IsEmptyValueIndirect(desired.Details) {
		c.Config.Logger.Infof("desired Details %s - but actually nil", dcl.SprintResource(desired.Details))
		return true
	}
	if compareJobStatusDetailsSlice(c, desired.Details, actual.Details) && !dcl.IsZeroValue(desired.Details) {
		c.Config.Logger.Infof("Diff in Details. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Details), dcl.SprintResource(actual.Details))
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
			c.Config.Logger.Infof("Diff in JobStatusDetails, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStatusDetails(c *Client, desired, actual *JobStatusDetails) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.TypeUrl == nil && desired.TypeUrl != nil && !dcl.IsEmptyValueIndirect(desired.TypeUrl) {
		c.Config.Logger.Infof("desired TypeUrl %s - but actually nil", dcl.SprintResource(desired.TypeUrl))
		return true
	}
	if !reflect.DeepEqual(desired.TypeUrl, actual.TypeUrl) && !dcl.IsZeroValue(desired.TypeUrl) && !(dcl.IsEmptyValueIndirect(desired.TypeUrl) && dcl.IsZeroValue(actual.TypeUrl)) {
		c.Config.Logger.Infof("Diff in TypeUrl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TypeUrl), dcl.SprintResource(actual.TypeUrl))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
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
			c.Config.Logger.Infof("Diff in JobRetryConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobRetryConfig(c *Client, desired, actual *JobRetryConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RetryCount == nil && desired.RetryCount != nil && !dcl.IsEmptyValueIndirect(desired.RetryCount) {
		c.Config.Logger.Infof("desired RetryCount %s - but actually nil", dcl.SprintResource(desired.RetryCount))
		return true
	}
	if !reflect.DeepEqual(desired.RetryCount, actual.RetryCount) && !dcl.IsZeroValue(desired.RetryCount) && !(dcl.IsEmptyValueIndirect(desired.RetryCount) && dcl.IsZeroValue(actual.RetryCount)) {
		c.Config.Logger.Infof("Diff in RetryCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetryCount), dcl.SprintResource(actual.RetryCount))
		return true
	}
	if actual.MaxRetryDuration == nil && desired.MaxRetryDuration != nil && !dcl.IsEmptyValueIndirect(desired.MaxRetryDuration) {
		c.Config.Logger.Infof("desired MaxRetryDuration %s - but actually nil", dcl.SprintResource(desired.MaxRetryDuration))
		return true
	}
	if compareJobRetryConfigMaxRetryDuration(c, desired.MaxRetryDuration, actual.MaxRetryDuration) && !dcl.IsZeroValue(desired.MaxRetryDuration) {
		c.Config.Logger.Infof("Diff in MaxRetryDuration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxRetryDuration), dcl.SprintResource(actual.MaxRetryDuration))
		return true
	}
	if actual.MinBackoffDuration == nil && desired.MinBackoffDuration != nil && !dcl.IsEmptyValueIndirect(desired.MinBackoffDuration) {
		c.Config.Logger.Infof("desired MinBackoffDuration %s - but actually nil", dcl.SprintResource(desired.MinBackoffDuration))
		return true
	}
	if compareJobRetryConfigMinBackoffDuration(c, desired.MinBackoffDuration, actual.MinBackoffDuration) && !dcl.IsZeroValue(desired.MinBackoffDuration) {
		c.Config.Logger.Infof("Diff in MinBackoffDuration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinBackoffDuration), dcl.SprintResource(actual.MinBackoffDuration))
		return true
	}
	if actual.MaxBackoffDuration == nil && desired.MaxBackoffDuration != nil && !dcl.IsEmptyValueIndirect(desired.MaxBackoffDuration) {
		c.Config.Logger.Infof("desired MaxBackoffDuration %s - but actually nil", dcl.SprintResource(desired.MaxBackoffDuration))
		return true
	}
	if compareJobRetryConfigMaxBackoffDuration(c, desired.MaxBackoffDuration, actual.MaxBackoffDuration) && !dcl.IsZeroValue(desired.MaxBackoffDuration) {
		c.Config.Logger.Infof("Diff in MaxBackoffDuration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxBackoffDuration), dcl.SprintResource(actual.MaxBackoffDuration))
		return true
	}
	if actual.MaxDoublings == nil && desired.MaxDoublings != nil && !dcl.IsEmptyValueIndirect(desired.MaxDoublings) {
		c.Config.Logger.Infof("desired MaxDoublings %s - but actually nil", dcl.SprintResource(desired.MaxDoublings))
		return true
	}
	if !reflect.DeepEqual(desired.MaxDoublings, actual.MaxDoublings) && !dcl.IsZeroValue(desired.MaxDoublings) && !(dcl.IsEmptyValueIndirect(desired.MaxDoublings) && dcl.IsZeroValue(actual.MaxDoublings)) {
		c.Config.Logger.Infof("Diff in MaxDoublings. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxDoublings), dcl.SprintResource(actual.MaxDoublings))
		return true
	}
	return false
}
func compareJobRetryConfigMaxRetryDurationSlice(c *Client, desired, actual []JobRetryConfigMaxRetryDuration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobRetryConfigMaxRetryDuration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobRetryConfigMaxRetryDuration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobRetryConfigMaxRetryDuration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobRetryConfigMaxRetryDuration(c *Client, desired, actual *JobRetryConfigMaxRetryDuration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareJobRetryConfigMinBackoffDurationSlice(c *Client, desired, actual []JobRetryConfigMinBackoffDuration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobRetryConfigMinBackoffDuration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobRetryConfigMinBackoffDuration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobRetryConfigMinBackoffDuration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobRetryConfigMinBackoffDuration(c *Client, desired, actual *JobRetryConfigMinBackoffDuration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareJobRetryConfigMaxBackoffDurationSlice(c *Client, desired, actual []JobRetryConfigMaxBackoffDuration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobRetryConfigMaxBackoffDuration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobRetryConfigMaxBackoffDuration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobRetryConfigMaxBackoffDuration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobRetryConfigMaxBackoffDuration(c *Client, desired, actual *JobRetryConfigMaxBackoffDuration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareJobAttemptDeadlineSlice(c *Client, desired, actual []JobAttemptDeadline) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobAttemptDeadline, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobAttemptDeadline(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobAttemptDeadline, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobAttemptDeadline(c *Client, desired, actual *JobAttemptDeadline) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
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
			c.Config.Logger.Infof("Diff in JobAppEngineHttpTargetHttpMethodEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
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
			c.Config.Logger.Infof("Diff in JobHttpTargetHttpMethodEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
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
			c.Config.Logger.Infof("Diff in JobStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStateEnum(c *Client, desired, actual *JobStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareJobViewEnumSlice(c *Client, desired, actual []JobViewEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobViewEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobViewEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobViewEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobViewEnum(c *Client, desired, actual *JobViewEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Job) urlNormalized() *Job {
	normalized := deepcopy.Copy(*r).(Job)
	normalized.Name = dcl.SelfLinkToName(r.Name)
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

	return flattenJob(c, m), nil
}

// expandJob expands Job into a JSON request object.
func expandJob(c *Client, f *Job) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/jobs/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandJobLabelsSlice(c, f.Labels); err != nil {
		return nil, fmt.Errorf("error expanding Labels into labels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandJobPubsubTarget(c, f.PubsubTarget); err != nil {
		return nil, fmt.Errorf("error expanding PubsubTarget into pubsubTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pubsubTarget"] = v
	}
	if v, err := expandJobAppEngineHttpTarget(c, f.AppEngineHttpTarget); err != nil {
		return nil, fmt.Errorf("error expanding AppEngineHttpTarget into appEngineHttpTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["appEngineHttpTarget"] = v
	}
	if v, err := expandJobHttpTarget(c, f.HttpTarget); err != nil {
		return nil, fmt.Errorf("error expanding HttpTarget into httpTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.TotalAttemptCount; !dcl.IsEmptyValueIndirect(v) {
		m["totalAttemptCount"] = v
	}
	if v := f.FailedAttemptCount; !dcl.IsEmptyValueIndirect(v) {
		m["failedAttemptCount"] = v
	}
	if v := f.TotalExecutionCount; !dcl.IsEmptyValueIndirect(v) {
		m["totalExecutionCount"] = v
	}
	if v := f.FailedExecutionCount; !dcl.IsEmptyValueIndirect(v) {
		m["failedExecutionCount"] = v
	}
	if v := f.View; !dcl.IsEmptyValueIndirect(v) {
		m["view"] = v
	}
	if v := f.ScheduleTime; !dcl.IsEmptyValueIndirect(v) {
		m["scheduleTime"] = v
	}
	if v := f.LastAttemptTime; !dcl.IsEmptyValueIndirect(v) {
		m["lastAttemptTime"] = v
	}
	if v, err := expandJobRetryConfig(c, f.RetryConfig); err != nil {
		return nil, fmt.Errorf("error expanding RetryConfig into retryConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["retryConfig"] = v
	}
	if v, err := expandJobAttemptDeadline(c, f.AttemptDeadline); err != nil {
		return nil, fmt.Errorf("error expanding AttemptDeadline into attemptDeadline: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["attemptDeadline"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Labels = flattenJobLabelsSlice(c, m["labels"])
	r.PubsubTarget = flattenJobPubsubTarget(c, m["pubsubTarget"])
	r.AppEngineHttpTarget = flattenJobAppEngineHttpTarget(c, m["appEngineHttpTarget"])
	r.HttpTarget = flattenJobHttpTarget(c, m["httpTarget"])
	r.Schedule = dcl.FlattenString(m["schedule"])
	r.TimeZone = dcl.FlattenString(m["timeZone"])
	r.UserUpdateTime = dcl.FlattenString(m["userUpdateTime"])
	r.State = flattenJobStateEnum(m["state"])
	r.Status = flattenJobStatus(c, m["status"])
	r.TotalAttemptCount = dcl.FlattenInteger(m["totalAttemptCount"])
	r.FailedAttemptCount = dcl.FlattenInteger(m["failedAttemptCount"])
	r.TotalExecutionCount = dcl.FlattenInteger(m["totalExecutionCount"])
	r.FailedExecutionCount = dcl.FlattenInteger(m["failedExecutionCount"])
	r.View = flattenJobViewEnum(m["view"])
	r.ScheduleTime = dcl.FlattenString(m["scheduleTime"])
	r.LastAttemptTime = dcl.FlattenString(m["lastAttemptTime"])
	r.RetryConfig = flattenJobRetryConfig(c, m["retryConfig"])
	r.AttemptDeadline = flattenJobAttemptDeadline(c, m["attemptDeadline"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandJobLabelsMap expands the contents of JobLabels into a JSON
// request object.
func expandJobLabelsMap(c *Client, f map[string]JobLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobLabelsSlice expands the contents of JobLabels into a JSON
// request object.
func expandJobLabelsSlice(c *Client, f []JobLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobLabelsMap flattens the contents of JobLabels from a JSON
// response object.
func flattenJobLabelsMap(c *Client, i interface{}) map[string]JobLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobLabels{}
	}

	if len(a) == 0 {
		return map[string]JobLabels{}
	}

	items := make(map[string]JobLabels)
	for k, item := range a {
		items[k] = *flattenJobLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobLabelsSlice flattens the contents of JobLabels from a JSON
// response object.
func flattenJobLabelsSlice(c *Client, i interface{}) []JobLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []JobLabels{}
	}

	if len(a) == 0 {
		return []JobLabels{}
	}

	items := make([]JobLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobLabels expands an instance of JobLabels into a JSON
// request object.
func expandJobLabels(c *Client, f *JobLabels) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenJobLabels flattens an instance of JobLabels from a JSON
// response object.
func flattenJobLabels(c *Client, i interface{}) *JobLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobLabels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TopicName; !dcl.IsEmptyValueIndirect(v) {
		m["topicName"] = v
	}
	if v := f.Data; !dcl.IsEmptyValueIndirect(v) {
		m["data"] = v
	}
	if v, err := expandJobPubsubTargetAttributesSlice(c, f.Attributes); err != nil {
		return nil, fmt.Errorf("error expanding Attributes into attributes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Attributes = flattenJobPubsubTargetAttributesSlice(c, m["attributes"])

	return r
}

// expandJobPubsubTargetAttributesMap expands the contents of JobPubsubTargetAttributes into a JSON
// request object.
func expandJobPubsubTargetAttributesMap(c *Client, f map[string]JobPubsubTargetAttributes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPubsubTargetAttributes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPubsubTargetAttributesSlice expands the contents of JobPubsubTargetAttributes into a JSON
// request object.
func expandJobPubsubTargetAttributesSlice(c *Client, f []JobPubsubTargetAttributes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPubsubTargetAttributes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPubsubTargetAttributesMap flattens the contents of JobPubsubTargetAttributes from a JSON
// response object.
func flattenJobPubsubTargetAttributesMap(c *Client, i interface{}) map[string]JobPubsubTargetAttributes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPubsubTargetAttributes{}
	}

	if len(a) == 0 {
		return map[string]JobPubsubTargetAttributes{}
	}

	items := make(map[string]JobPubsubTargetAttributes)
	for k, item := range a {
		items[k] = *flattenJobPubsubTargetAttributes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPubsubTargetAttributesSlice flattens the contents of JobPubsubTargetAttributes from a JSON
// response object.
func flattenJobPubsubTargetAttributesSlice(c *Client, i interface{}) []JobPubsubTargetAttributes {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPubsubTargetAttributes{}
	}

	if len(a) == 0 {
		return []JobPubsubTargetAttributes{}
	}

	items := make([]JobPubsubTargetAttributes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPubsubTargetAttributes(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPubsubTargetAttributes expands an instance of JobPubsubTargetAttributes into a JSON
// request object.
func expandJobPubsubTargetAttributes(c *Client, f *JobPubsubTargetAttributes) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenJobPubsubTargetAttributes flattens an instance of JobPubsubTargetAttributes from a JSON
// response object.
func flattenJobPubsubTargetAttributes(c *Client, i interface{}) *JobPubsubTargetAttributes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPubsubTargetAttributes{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if v, err := expandJobAppEngineHttpTargetHeadersSlice(c, f.Headers); err != nil {
		return nil, fmt.Errorf("error expanding Headers into headers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Headers = flattenJobAppEngineHttpTargetHeadersSlice(c, m["headers"])
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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

// expandJobAppEngineHttpTargetHeadersMap expands the contents of JobAppEngineHttpTargetHeaders into a JSON
// request object.
func expandJobAppEngineHttpTargetHeadersMap(c *Client, f map[string]JobAppEngineHttpTargetHeaders) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobAppEngineHttpTargetHeaders(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobAppEngineHttpTargetHeadersSlice expands the contents of JobAppEngineHttpTargetHeaders into a JSON
// request object.
func expandJobAppEngineHttpTargetHeadersSlice(c *Client, f []JobAppEngineHttpTargetHeaders) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobAppEngineHttpTargetHeaders(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobAppEngineHttpTargetHeadersMap flattens the contents of JobAppEngineHttpTargetHeaders from a JSON
// response object.
func flattenJobAppEngineHttpTargetHeadersMap(c *Client, i interface{}) map[string]JobAppEngineHttpTargetHeaders {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobAppEngineHttpTargetHeaders{}
	}

	if len(a) == 0 {
		return map[string]JobAppEngineHttpTargetHeaders{}
	}

	items := make(map[string]JobAppEngineHttpTargetHeaders)
	for k, item := range a {
		items[k] = *flattenJobAppEngineHttpTargetHeaders(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobAppEngineHttpTargetHeadersSlice flattens the contents of JobAppEngineHttpTargetHeaders from a JSON
// response object.
func flattenJobAppEngineHttpTargetHeadersSlice(c *Client, i interface{}) []JobAppEngineHttpTargetHeaders {
	a, ok := i.([]interface{})
	if !ok {
		return []JobAppEngineHttpTargetHeaders{}
	}

	if len(a) == 0 {
		return []JobAppEngineHttpTargetHeaders{}
	}

	items := make([]JobAppEngineHttpTargetHeaders, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobAppEngineHttpTargetHeaders(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobAppEngineHttpTargetHeaders expands an instance of JobAppEngineHttpTargetHeaders into a JSON
// request object.
func expandJobAppEngineHttpTargetHeaders(c *Client, f *JobAppEngineHttpTargetHeaders) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenJobAppEngineHttpTargetHeaders flattens an instance of JobAppEngineHttpTargetHeaders from a JSON
// response object.
func flattenJobAppEngineHttpTargetHeaders(c *Client, i interface{}) *JobAppEngineHttpTargetHeaders {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobAppEngineHttpTargetHeaders{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Uri; !dcl.IsEmptyValueIndirect(v) {
		m["uri"] = v
	}
	if v := f.HttpMethod; !dcl.IsEmptyValueIndirect(v) {
		m["httpMethod"] = v
	}
	if v, err := expandJobHttpTargetHeadersSlice(c, f.Headers); err != nil {
		return nil, fmt.Errorf("error expanding Headers into headers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Headers = flattenJobHttpTargetHeadersSlice(c, m["headers"])
	r.Body = dcl.FlattenString(m["body"])
	r.OAuthToken = flattenJobHttpTargetOAuthToken(c, m["oauthToken"])
	r.OidcToken = flattenJobHttpTargetOidcToken(c, m["oidcToken"])

	return r
}

// expandJobHttpTargetHeadersMap expands the contents of JobHttpTargetHeaders into a JSON
// request object.
func expandJobHttpTargetHeadersMap(c *Client, f map[string]JobHttpTargetHeaders) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHttpTargetHeaders(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHttpTargetHeadersSlice expands the contents of JobHttpTargetHeaders into a JSON
// request object.
func expandJobHttpTargetHeadersSlice(c *Client, f []JobHttpTargetHeaders) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHttpTargetHeaders(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHttpTargetHeadersMap flattens the contents of JobHttpTargetHeaders from a JSON
// response object.
func flattenJobHttpTargetHeadersMap(c *Client, i interface{}) map[string]JobHttpTargetHeaders {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHttpTargetHeaders{}
	}

	if len(a) == 0 {
		return map[string]JobHttpTargetHeaders{}
	}

	items := make(map[string]JobHttpTargetHeaders)
	for k, item := range a {
		items[k] = *flattenJobHttpTargetHeaders(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobHttpTargetHeadersSlice flattens the contents of JobHttpTargetHeaders from a JSON
// response object.
func flattenJobHttpTargetHeadersSlice(c *Client, i interface{}) []JobHttpTargetHeaders {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHttpTargetHeaders{}
	}

	if len(a) == 0 {
		return []JobHttpTargetHeaders{}
	}

	items := make([]JobHttpTargetHeaders, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHttpTargetHeaders(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobHttpTargetHeaders expands an instance of JobHttpTargetHeaders into a JSON
// request object.
func expandJobHttpTargetHeaders(c *Client, f *JobHttpTargetHeaders) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenJobHttpTargetHeaders flattens an instance of JobHttpTargetHeaders from a JSON
// response object.
func flattenJobHttpTargetHeaders(c *Client, i interface{}) *JobHttpTargetHeaders {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHttpTargetHeaders{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RetryCount; !dcl.IsEmptyValueIndirect(v) {
		m["retryCount"] = v
	}
	if v, err := expandJobRetryConfigMaxRetryDuration(c, f.MaxRetryDuration); err != nil {
		return nil, fmt.Errorf("error expanding MaxRetryDuration into maxRetryDuration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maxRetryDuration"] = v
	}
	if v, err := expandJobRetryConfigMinBackoffDuration(c, f.MinBackoffDuration); err != nil {
		return nil, fmt.Errorf("error expanding MinBackoffDuration into minBackoffDuration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["minBackoffDuration"] = v
	}
	if v, err := expandJobRetryConfigMaxBackoffDuration(c, f.MaxBackoffDuration); err != nil {
		return nil, fmt.Errorf("error expanding MaxBackoffDuration into maxBackoffDuration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.MaxRetryDuration = flattenJobRetryConfigMaxRetryDuration(c, m["maxRetryDuration"])
	r.MinBackoffDuration = flattenJobRetryConfigMinBackoffDuration(c, m["minBackoffDuration"])
	r.MaxBackoffDuration = flattenJobRetryConfigMaxBackoffDuration(c, m["maxBackoffDuration"])
	r.MaxDoublings = dcl.FlattenInteger(m["maxDoublings"])

	return r
}

// expandJobRetryConfigMaxRetryDurationMap expands the contents of JobRetryConfigMaxRetryDuration into a JSON
// request object.
func expandJobRetryConfigMaxRetryDurationMap(c *Client, f map[string]JobRetryConfigMaxRetryDuration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobRetryConfigMaxRetryDuration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobRetryConfigMaxRetryDurationSlice expands the contents of JobRetryConfigMaxRetryDuration into a JSON
// request object.
func expandJobRetryConfigMaxRetryDurationSlice(c *Client, f []JobRetryConfigMaxRetryDuration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobRetryConfigMaxRetryDuration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobRetryConfigMaxRetryDurationMap flattens the contents of JobRetryConfigMaxRetryDuration from a JSON
// response object.
func flattenJobRetryConfigMaxRetryDurationMap(c *Client, i interface{}) map[string]JobRetryConfigMaxRetryDuration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobRetryConfigMaxRetryDuration{}
	}

	if len(a) == 0 {
		return map[string]JobRetryConfigMaxRetryDuration{}
	}

	items := make(map[string]JobRetryConfigMaxRetryDuration)
	for k, item := range a {
		items[k] = *flattenJobRetryConfigMaxRetryDuration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobRetryConfigMaxRetryDurationSlice flattens the contents of JobRetryConfigMaxRetryDuration from a JSON
// response object.
func flattenJobRetryConfigMaxRetryDurationSlice(c *Client, i interface{}) []JobRetryConfigMaxRetryDuration {
	a, ok := i.([]interface{})
	if !ok {
		return []JobRetryConfigMaxRetryDuration{}
	}

	if len(a) == 0 {
		return []JobRetryConfigMaxRetryDuration{}
	}

	items := make([]JobRetryConfigMaxRetryDuration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobRetryConfigMaxRetryDuration(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobRetryConfigMaxRetryDuration expands an instance of JobRetryConfigMaxRetryDuration into a JSON
// request object.
func expandJobRetryConfigMaxRetryDuration(c *Client, f *JobRetryConfigMaxRetryDuration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenJobRetryConfigMaxRetryDuration flattens an instance of JobRetryConfigMaxRetryDuration from a JSON
// response object.
func flattenJobRetryConfigMaxRetryDuration(c *Client, i interface{}) *JobRetryConfigMaxRetryDuration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobRetryConfigMaxRetryDuration{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandJobRetryConfigMinBackoffDurationMap expands the contents of JobRetryConfigMinBackoffDuration into a JSON
// request object.
func expandJobRetryConfigMinBackoffDurationMap(c *Client, f map[string]JobRetryConfigMinBackoffDuration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobRetryConfigMinBackoffDuration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobRetryConfigMinBackoffDurationSlice expands the contents of JobRetryConfigMinBackoffDuration into a JSON
// request object.
func expandJobRetryConfigMinBackoffDurationSlice(c *Client, f []JobRetryConfigMinBackoffDuration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobRetryConfigMinBackoffDuration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobRetryConfigMinBackoffDurationMap flattens the contents of JobRetryConfigMinBackoffDuration from a JSON
// response object.
func flattenJobRetryConfigMinBackoffDurationMap(c *Client, i interface{}) map[string]JobRetryConfigMinBackoffDuration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobRetryConfigMinBackoffDuration{}
	}

	if len(a) == 0 {
		return map[string]JobRetryConfigMinBackoffDuration{}
	}

	items := make(map[string]JobRetryConfigMinBackoffDuration)
	for k, item := range a {
		items[k] = *flattenJobRetryConfigMinBackoffDuration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobRetryConfigMinBackoffDurationSlice flattens the contents of JobRetryConfigMinBackoffDuration from a JSON
// response object.
func flattenJobRetryConfigMinBackoffDurationSlice(c *Client, i interface{}) []JobRetryConfigMinBackoffDuration {
	a, ok := i.([]interface{})
	if !ok {
		return []JobRetryConfigMinBackoffDuration{}
	}

	if len(a) == 0 {
		return []JobRetryConfigMinBackoffDuration{}
	}

	items := make([]JobRetryConfigMinBackoffDuration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobRetryConfigMinBackoffDuration(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobRetryConfigMinBackoffDuration expands an instance of JobRetryConfigMinBackoffDuration into a JSON
// request object.
func expandJobRetryConfigMinBackoffDuration(c *Client, f *JobRetryConfigMinBackoffDuration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenJobRetryConfigMinBackoffDuration flattens an instance of JobRetryConfigMinBackoffDuration from a JSON
// response object.
func flattenJobRetryConfigMinBackoffDuration(c *Client, i interface{}) *JobRetryConfigMinBackoffDuration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobRetryConfigMinBackoffDuration{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandJobRetryConfigMaxBackoffDurationMap expands the contents of JobRetryConfigMaxBackoffDuration into a JSON
// request object.
func expandJobRetryConfigMaxBackoffDurationMap(c *Client, f map[string]JobRetryConfigMaxBackoffDuration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobRetryConfigMaxBackoffDuration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobRetryConfigMaxBackoffDurationSlice expands the contents of JobRetryConfigMaxBackoffDuration into a JSON
// request object.
func expandJobRetryConfigMaxBackoffDurationSlice(c *Client, f []JobRetryConfigMaxBackoffDuration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobRetryConfigMaxBackoffDuration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobRetryConfigMaxBackoffDurationMap flattens the contents of JobRetryConfigMaxBackoffDuration from a JSON
// response object.
func flattenJobRetryConfigMaxBackoffDurationMap(c *Client, i interface{}) map[string]JobRetryConfigMaxBackoffDuration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobRetryConfigMaxBackoffDuration{}
	}

	if len(a) == 0 {
		return map[string]JobRetryConfigMaxBackoffDuration{}
	}

	items := make(map[string]JobRetryConfigMaxBackoffDuration)
	for k, item := range a {
		items[k] = *flattenJobRetryConfigMaxBackoffDuration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobRetryConfigMaxBackoffDurationSlice flattens the contents of JobRetryConfigMaxBackoffDuration from a JSON
// response object.
func flattenJobRetryConfigMaxBackoffDurationSlice(c *Client, i interface{}) []JobRetryConfigMaxBackoffDuration {
	a, ok := i.([]interface{})
	if !ok {
		return []JobRetryConfigMaxBackoffDuration{}
	}

	if len(a) == 0 {
		return []JobRetryConfigMaxBackoffDuration{}
	}

	items := make([]JobRetryConfigMaxBackoffDuration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobRetryConfigMaxBackoffDuration(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobRetryConfigMaxBackoffDuration expands an instance of JobRetryConfigMaxBackoffDuration into a JSON
// request object.
func expandJobRetryConfigMaxBackoffDuration(c *Client, f *JobRetryConfigMaxBackoffDuration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenJobRetryConfigMaxBackoffDuration flattens an instance of JobRetryConfigMaxBackoffDuration from a JSON
// response object.
func flattenJobRetryConfigMaxBackoffDuration(c *Client, i interface{}) *JobRetryConfigMaxBackoffDuration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobRetryConfigMaxBackoffDuration{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandJobAttemptDeadlineMap expands the contents of JobAttemptDeadline into a JSON
// request object.
func expandJobAttemptDeadlineMap(c *Client, f map[string]JobAttemptDeadline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobAttemptDeadline(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobAttemptDeadlineSlice expands the contents of JobAttemptDeadline into a JSON
// request object.
func expandJobAttemptDeadlineSlice(c *Client, f []JobAttemptDeadline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobAttemptDeadline(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobAttemptDeadlineMap flattens the contents of JobAttemptDeadline from a JSON
// response object.
func flattenJobAttemptDeadlineMap(c *Client, i interface{}) map[string]JobAttemptDeadline {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobAttemptDeadline{}
	}

	if len(a) == 0 {
		return map[string]JobAttemptDeadline{}
	}

	items := make(map[string]JobAttemptDeadline)
	for k, item := range a {
		items[k] = *flattenJobAttemptDeadline(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobAttemptDeadlineSlice flattens the contents of JobAttemptDeadline from a JSON
// response object.
func flattenJobAttemptDeadlineSlice(c *Client, i interface{}) []JobAttemptDeadline {
	a, ok := i.([]interface{})
	if !ok {
		return []JobAttemptDeadline{}
	}

	if len(a) == 0 {
		return []JobAttemptDeadline{}
	}

	items := make([]JobAttemptDeadline, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobAttemptDeadline(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobAttemptDeadline expands an instance of JobAttemptDeadline into a JSON
// request object.
func expandJobAttemptDeadline(c *Client, f *JobAttemptDeadline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenJobAttemptDeadline flattens an instance of JobAttemptDeadline from a JSON
// response object.
func flattenJobAttemptDeadline(c *Client, i interface{}) *JobAttemptDeadline {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobAttemptDeadline{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

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
		items = append(items, *flattenJobAppEngineHttpTargetHttpMethodEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenJobHttpTargetHttpMethodEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenJobStateEnum(item.(map[string]interface{})))
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

// flattenJobViewEnumSlice flattens the contents of JobViewEnum from a JSON
// response object.
func flattenJobViewEnumSlice(c *Client, i interface{}) []JobViewEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobViewEnum{}
	}

	if len(a) == 0 {
		return []JobViewEnum{}
	}

	items := make([]JobViewEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobViewEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenJobViewEnum asserts that an interface is a string, and returns a
// pointer to a *JobViewEnum with the same value as that string.
func flattenJobViewEnum(i interface{}) *JobViewEnum {
	s, ok := i.(string)
	if !ok {
		return JobViewEnumRef("")
	}

	return JobViewEnumRef(s)
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
