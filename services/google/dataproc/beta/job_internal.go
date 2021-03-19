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
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Job) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"HadoopJob", "SparkJob", "PysparkJob", "HiveJob", "PigJob", "SparkRJob", "SparkSqlJob", "PrestoJob"}, r.HadoopJob, r.SparkJob, r.PysparkJob, r.HiveJob, r.PigJob, r.SparkRJob, r.SparkSqlJob, r.PrestoJob); err != nil {
		return err
	}
	if err := dcl.Required(r, "placement"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Region, "Region"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Reference) {
		if err := r.Reference.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Placement) {
		if err := r.Placement.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HadoopJob) {
		if err := r.HadoopJob.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SparkJob) {
		if err := r.SparkJob.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PysparkJob) {
		if err := r.PysparkJob.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HiveJob) {
		if err := r.HiveJob.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PigJob) {
		if err := r.PigJob.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SparkRJob) {
		if err := r.SparkRJob.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SparkSqlJob) {
		if err := r.SparkSqlJob.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PrestoJob) {
		if err := r.PrestoJob.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Scheduling) {
		if err := r.Scheduling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DriverRunner) {
		if err := r.DriverRunner.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobReference) validate() error {
	return nil
}
func (r *JobPlacement) validate() error {
	if err := dcl.Required(r, "clusterName"); err != nil {
		return err
	}
	return nil
}
func (r *JobHadoopJob) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobHadoopJobLoggingConfig) validate() error {
	return nil
}
func (r *JobSparkJob) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobSparkJobLoggingConfig) validate() error {
	return nil
}
func (r *JobPysparkJob) validate() error {
	if err := dcl.Required(r, "mainPythonFileUri"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobPysparkJobLoggingConfig) validate() error {
	return nil
}
func (r *JobHiveJob) validate() error {
	if !dcl.IsEmptyValueIndirect(r.QueryList) {
		if err := r.QueryList.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobHiveJobQueryList) validate() error {
	if err := dcl.Required(r, "queries"); err != nil {
		return err
	}
	return nil
}
func (r *JobPigJob) validate() error {
	if !dcl.IsEmptyValueIndirect(r.QueryList) {
		if err := r.QueryList.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobPigJobQueryList) validate() error {
	if err := dcl.Required(r, "queries"); err != nil {
		return err
	}
	return nil
}
func (r *JobPigJobLoggingConfig) validate() error {
	return nil
}
func (r *JobSparkRJob) validate() error {
	if err := dcl.Required(r, "mainRFileUri"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobSparkRJobLoggingConfig) validate() error {
	return nil
}
func (r *JobSparkSqlJob) validate() error {
	if !dcl.IsEmptyValueIndirect(r.QueryList) {
		if err := r.QueryList.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobSparkSqlJobQueryList) validate() error {
	if err := dcl.Required(r, "queries"); err != nil {
		return err
	}
	return nil
}
func (r *JobSparkSqlJobLoggingConfig) validate() error {
	return nil
}
func (r *JobPrestoJob) validate() error {
	if !dcl.IsEmptyValueIndirect(r.QueryList) {
		if err := r.QueryList.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobPrestoJobQueryList) validate() error {
	if err := dcl.Required(r, "queries"); err != nil {
		return err
	}
	return nil
}
func (r *JobPrestoJobLoggingConfig) validate() error {
	return nil
}
func (r *JobStatus) validate() error {
	return nil
}
func (r *JobStatusHistory) validate() error {
	return nil
}
func (r *JobYarnApplications) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "state"); err != nil {
		return err
	}
	if err := dcl.Required(r, "progress"); err != nil {
		return err
	}
	return nil
}
func (r *JobScheduling) validate() error {
	return nil
}
func (r *JobDriverRunner) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MasterDriverRunner) {
		if err := r.MasterDriverRunner.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.YarnDriverRunner) {
		if err := r.YarnDriverRunner.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobDriverRunnerMasterDriverRunner) validate() error {
	return nil
}
func (r *JobDriverRunnerYarnDriverRunner) validate() error {
	return nil
}

func jobGetURL(userBasePath string, r *Job) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/jobs/{{name}}", "https://dataproc.googleapis.com/v1beta2/", userBasePath, params), nil
}

func jobListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/jobs", "https://dataproc.googleapis.com/v1beta2/", userBasePath, params), nil

}

func jobCreateURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/jobs:submitAsOperation", "https://dataproc.googleapis.com/v1beta2/", userBasePath, params), nil

}

func jobDeleteURL(userBasePath string, r *Job) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/jobs/{{name}}", "https://dataproc.googleapis.com/v1beta2/", userBasePath, params), nil
}

// jobApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type jobApiOperation interface {
	do(context.Context, *Job, *Client) error
}

// newUpdateJobPatchRequest creates a request for an
// Job resource's patch update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateJobPatchRequest(ctx context.Context, f *Job, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	return req, nil
}

// marshalUpdateJobPatchRequest converts the update into
// the final JSON request body.
func marshalUpdateJobPatchRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateJobPatchOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateJobPatchOperation) do(ctx context.Context, r *Job, c *Client) error {
	_, err := c.GetJob(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "patch")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"labels"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateJobPatchRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateJobPatchRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listJobRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := jobListURL(c.Config.BasePath, project, region)
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

func (c *Client) listJob(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*Job, string, error) {
	b, err := c.listJobRaw(ctx, project, region, pageToken, pageSize)
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
		res.Region = &region
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
	_, err = c.GetJob(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
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

	project, region := r.createFields()
	u, err := jobCreateURL(c.Config.BasePath, project, region)

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
	if err := o.Wait(ctx, c.Config, "https://dataproc.googleapis.com/v1beta2/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["jobUuid"].(string)
	if !ok {
		return fmt.Errorf("expected jobUuid to be a string")
	}
	r.Name = &name

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

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeJobDesiredState(rawDesired, nil)
		return nil, desired, nil, err
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

	if dcl.IsZeroValue(rawInitial.HadoopJob) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.SparkJob, rawInitial.PysparkJob, rawInitial.HiveJob, rawInitial.PigJob, rawInitial.SparkRJob, rawInitial.SparkSqlJob, rawInitial.PrestoJob) {
			rawInitial.HadoopJob = EmptyJobHadoopJob
		}
	}

	if dcl.IsZeroValue(rawInitial.SparkJob) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.HadoopJob, rawInitial.PysparkJob, rawInitial.HiveJob, rawInitial.PigJob, rawInitial.SparkRJob, rawInitial.SparkSqlJob, rawInitial.PrestoJob) {
			rawInitial.SparkJob = EmptyJobSparkJob
		}
	}

	if dcl.IsZeroValue(rawInitial.PysparkJob) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.HadoopJob, rawInitial.SparkJob, rawInitial.HiveJob, rawInitial.PigJob, rawInitial.SparkRJob, rawInitial.SparkSqlJob, rawInitial.PrestoJob) {
			rawInitial.PysparkJob = EmptyJobPysparkJob
		}
	}

	if dcl.IsZeroValue(rawInitial.HiveJob) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.HadoopJob, rawInitial.SparkJob, rawInitial.PysparkJob, rawInitial.PigJob, rawInitial.SparkRJob, rawInitial.SparkSqlJob, rawInitial.PrestoJob) {
			rawInitial.HiveJob = EmptyJobHiveJob
		}
	}

	if dcl.IsZeroValue(rawInitial.PigJob) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.HadoopJob, rawInitial.SparkJob, rawInitial.PysparkJob, rawInitial.HiveJob, rawInitial.SparkRJob, rawInitial.SparkSqlJob, rawInitial.PrestoJob) {
			rawInitial.PigJob = EmptyJobPigJob
		}
	}

	if dcl.IsZeroValue(rawInitial.SparkRJob) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.HadoopJob, rawInitial.SparkJob, rawInitial.PysparkJob, rawInitial.HiveJob, rawInitial.PigJob, rawInitial.SparkSqlJob, rawInitial.PrestoJob) {
			rawInitial.SparkRJob = EmptyJobSparkRJob
		}
	}

	if dcl.IsZeroValue(rawInitial.SparkSqlJob) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.HadoopJob, rawInitial.SparkJob, rawInitial.PysparkJob, rawInitial.HiveJob, rawInitial.PigJob, rawInitial.SparkRJob, rawInitial.PrestoJob) {
			rawInitial.SparkSqlJob = EmptyJobSparkSqlJob
		}
	}

	if dcl.IsZeroValue(rawInitial.PrestoJob) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.HadoopJob, rawInitial.SparkJob, rawInitial.PysparkJob, rawInitial.HiveJob, rawInitial.PigJob, rawInitial.SparkRJob, rawInitial.SparkSqlJob) {
			rawInitial.PrestoJob = EmptyJobPrestoJob
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

func canonicalizeJobDesiredState(rawDesired, rawInitial *Job, opts ...dcl.ApplyOption) (*Job, error) {

	if dcl.IsZeroValue(rawDesired.HadoopJob) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.SparkJob, rawDesired.PysparkJob, rawDesired.HiveJob, rawDesired.PigJob, rawDesired.SparkRJob, rawDesired.SparkSqlJob, rawDesired.PrestoJob) {
			rawDesired.HadoopJob = EmptyJobHadoopJob
		}
	}

	if dcl.IsZeroValue(rawDesired.SparkJob) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HadoopJob, rawDesired.PysparkJob, rawDesired.HiveJob, rawDesired.PigJob, rawDesired.SparkRJob, rawDesired.SparkSqlJob, rawDesired.PrestoJob) {
			rawDesired.SparkJob = EmptyJobSparkJob
		}
	}

	if dcl.IsZeroValue(rawDesired.PysparkJob) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HadoopJob, rawDesired.SparkJob, rawDesired.HiveJob, rawDesired.PigJob, rawDesired.SparkRJob, rawDesired.SparkSqlJob, rawDesired.PrestoJob) {
			rawDesired.PysparkJob = EmptyJobPysparkJob
		}
	}

	if dcl.IsZeroValue(rawDesired.HiveJob) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HadoopJob, rawDesired.SparkJob, rawDesired.PysparkJob, rawDesired.PigJob, rawDesired.SparkRJob, rawDesired.SparkSqlJob, rawDesired.PrestoJob) {
			rawDesired.HiveJob = EmptyJobHiveJob
		}
	}

	if dcl.IsZeroValue(rawDesired.PigJob) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HadoopJob, rawDesired.SparkJob, rawDesired.PysparkJob, rawDesired.HiveJob, rawDesired.SparkRJob, rawDesired.SparkSqlJob, rawDesired.PrestoJob) {
			rawDesired.PigJob = EmptyJobPigJob
		}
	}

	if dcl.IsZeroValue(rawDesired.SparkRJob) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HadoopJob, rawDesired.SparkJob, rawDesired.PysparkJob, rawDesired.HiveJob, rawDesired.PigJob, rawDesired.SparkSqlJob, rawDesired.PrestoJob) {
			rawDesired.SparkRJob = EmptyJobSparkRJob
		}
	}

	if dcl.IsZeroValue(rawDesired.SparkSqlJob) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HadoopJob, rawDesired.SparkJob, rawDesired.PysparkJob, rawDesired.HiveJob, rawDesired.PigJob, rawDesired.SparkRJob, rawDesired.PrestoJob) {
			rawDesired.SparkSqlJob = EmptyJobSparkSqlJob
		}
	}

	if dcl.IsZeroValue(rawDesired.PrestoJob) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HadoopJob, rawDesired.SparkJob, rawDesired.PysparkJob, rawDesired.HiveJob, rawDesired.PigJob, rawDesired.SparkRJob, rawDesired.SparkSqlJob) {
			rawDesired.PrestoJob = EmptyJobPrestoJob
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Reference = canonicalizeJobReference(rawDesired.Reference, nil, opts...)
		rawDesired.Placement = canonicalizeJobPlacement(rawDesired.Placement, nil, opts...)
		rawDesired.HadoopJob = canonicalizeJobHadoopJob(rawDesired.HadoopJob, nil, opts...)
		rawDesired.SparkJob = canonicalizeJobSparkJob(rawDesired.SparkJob, nil, opts...)
		rawDesired.PysparkJob = canonicalizeJobPysparkJob(rawDesired.PysparkJob, nil, opts...)
		rawDesired.HiveJob = canonicalizeJobHiveJob(rawDesired.HiveJob, nil, opts...)
		rawDesired.PigJob = canonicalizeJobPigJob(rawDesired.PigJob, nil, opts...)
		rawDesired.SparkRJob = canonicalizeJobSparkRJob(rawDesired.SparkRJob, nil, opts...)
		rawDesired.SparkSqlJob = canonicalizeJobSparkSqlJob(rawDesired.SparkSqlJob, nil, opts...)
		rawDesired.PrestoJob = canonicalizeJobPrestoJob(rawDesired.PrestoJob, nil, opts...)
		rawDesired.Status = canonicalizeJobStatus(rawDesired.Status, nil, opts...)
		rawDesired.Scheduling = canonicalizeJobScheduling(rawDesired.Scheduling, nil, opts...)
		rawDesired.DriverRunner = canonicalizeJobDriverRunner(rawDesired.DriverRunner, nil, opts...)

		return rawDesired, nil
	}
	rawDesired.Reference = canonicalizeJobReference(rawDesired.Reference, rawInitial.Reference, opts...)
	rawDesired.Placement = canonicalizeJobPlacement(rawDesired.Placement, rawInitial.Placement, opts...)
	rawDesired.HadoopJob = canonicalizeJobHadoopJob(rawDesired.HadoopJob, rawInitial.HadoopJob, opts...)
	rawDesired.SparkJob = canonicalizeJobSparkJob(rawDesired.SparkJob, rawInitial.SparkJob, opts...)
	rawDesired.PysparkJob = canonicalizeJobPysparkJob(rawDesired.PysparkJob, rawInitial.PysparkJob, opts...)
	rawDesired.HiveJob = canonicalizeJobHiveJob(rawDesired.HiveJob, rawInitial.HiveJob, opts...)
	rawDesired.PigJob = canonicalizeJobPigJob(rawDesired.PigJob, rawInitial.PigJob, opts...)
	rawDesired.SparkRJob = canonicalizeJobSparkRJob(rawDesired.SparkRJob, rawInitial.SparkRJob, opts...)
	rawDesired.SparkSqlJob = canonicalizeJobSparkSqlJob(rawDesired.SparkSqlJob, rawInitial.SparkSqlJob, opts...)
	rawDesired.PrestoJob = canonicalizeJobPrestoJob(rawDesired.PrestoJob, rawInitial.PrestoJob, opts...)
	rawDesired.Status = canonicalizeJobStatus(rawDesired.Status, rawInitial.Status, opts...)
	if dcl.IsZeroValue(rawDesired.StatusHistory) {
		rawDesired.StatusHistory = rawInitial.StatusHistory
	}
	if dcl.IsZeroValue(rawDesired.YarnApplications) {
		rawDesired.YarnApplications = rawInitial.YarnApplications
	}
	if dcl.StringCanonicalize(rawDesired.SubmittedBy, rawInitial.SubmittedBy) {
		rawDesired.SubmittedBy = rawInitial.SubmittedBy
	}
	if dcl.StringCanonicalize(rawDesired.DriverInputResourceUri, rawInitial.DriverInputResourceUri) {
		rawDesired.DriverInputResourceUri = rawInitial.DriverInputResourceUri
	}
	if dcl.StringCanonicalize(rawDesired.DriverOutputResourceUri, rawInitial.DriverOutputResourceUri) {
		rawDesired.DriverOutputResourceUri = rawInitial.DriverOutputResourceUri
	}
	if dcl.StringCanonicalize(rawDesired.DriverControlFilesUri, rawInitial.DriverControlFilesUri) {
		rawDesired.DriverControlFilesUri = rawInitial.DriverControlFilesUri
	}
	if dcl.IsZeroValue(rawDesired.Interactive) {
		rawDesired.Interactive = rawInitial.Interactive
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	rawDesired.Scheduling = canonicalizeJobScheduling(rawDesired.Scheduling, rawInitial.Scheduling, opts...)
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Done) {
		rawDesired.Done = rawInitial.Done
	}
	rawDesired.DriverRunner = canonicalizeJobDriverRunner(rawDesired.DriverRunner, rawInitial.DriverRunner, opts...)
	if dcl.NameToSelfLink(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeJobNewState(c *Client, rawNew, rawDesired *Job) (*Job, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Reference) && dcl.IsEmptyValueIndirect(rawDesired.Reference) {
		rawNew.Reference = rawDesired.Reference
	} else {
		rawNew.Reference = canonicalizeNewJobReference(c, rawDesired.Reference, rawNew.Reference)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Placement) && dcl.IsEmptyValueIndirect(rawDesired.Placement) {
		rawNew.Placement = rawDesired.Placement
	} else {
		rawNew.Placement = canonicalizeNewJobPlacement(c, rawDesired.Placement, rawNew.Placement)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HadoopJob) && dcl.IsEmptyValueIndirect(rawDesired.HadoopJob) {
		rawNew.HadoopJob = rawDesired.HadoopJob
	} else {
		rawNew.HadoopJob = canonicalizeNewJobHadoopJob(c, rawDesired.HadoopJob, rawNew.HadoopJob)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SparkJob) && dcl.IsEmptyValueIndirect(rawDesired.SparkJob) {
		rawNew.SparkJob = rawDesired.SparkJob
	} else {
		rawNew.SparkJob = canonicalizeNewJobSparkJob(c, rawDesired.SparkJob, rawNew.SparkJob)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PysparkJob) && dcl.IsEmptyValueIndirect(rawDesired.PysparkJob) {
		rawNew.PysparkJob = rawDesired.PysparkJob
	} else {
		rawNew.PysparkJob = canonicalizeNewJobPysparkJob(c, rawDesired.PysparkJob, rawNew.PysparkJob)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HiveJob) && dcl.IsEmptyValueIndirect(rawDesired.HiveJob) {
		rawNew.HiveJob = rawDesired.HiveJob
	} else {
		rawNew.HiveJob = canonicalizeNewJobHiveJob(c, rawDesired.HiveJob, rawNew.HiveJob)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PigJob) && dcl.IsEmptyValueIndirect(rawDesired.PigJob) {
		rawNew.PigJob = rawDesired.PigJob
	} else {
		rawNew.PigJob = canonicalizeNewJobPigJob(c, rawDesired.PigJob, rawNew.PigJob)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SparkRJob) && dcl.IsEmptyValueIndirect(rawDesired.SparkRJob) {
		rawNew.SparkRJob = rawDesired.SparkRJob
	} else {
		rawNew.SparkRJob = canonicalizeNewJobSparkRJob(c, rawDesired.SparkRJob, rawNew.SparkRJob)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SparkSqlJob) && dcl.IsEmptyValueIndirect(rawDesired.SparkSqlJob) {
		rawNew.SparkSqlJob = rawDesired.SparkSqlJob
	} else {
		rawNew.SparkSqlJob = canonicalizeNewJobSparkSqlJob(c, rawDesired.SparkSqlJob, rawNew.SparkSqlJob)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PrestoJob) && dcl.IsEmptyValueIndirect(rawDesired.PrestoJob) {
		rawNew.PrestoJob = rawDesired.PrestoJob
	} else {
		rawNew.PrestoJob = canonicalizeNewJobPrestoJob(c, rawDesired.PrestoJob, rawNew.PrestoJob)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
		rawNew.Status = canonicalizeNewJobStatus(c, rawDesired.Status, rawNew.Status)
	}

	if dcl.IsEmptyValueIndirect(rawNew.StatusHistory) && dcl.IsEmptyValueIndirect(rawDesired.StatusHistory) {
		rawNew.StatusHistory = rawDesired.StatusHistory
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.YarnApplications) && dcl.IsEmptyValueIndirect(rawDesired.YarnApplications) {
		rawNew.YarnApplications = rawDesired.YarnApplications
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SubmittedBy) && dcl.IsEmptyValueIndirect(rawDesired.SubmittedBy) {
		rawNew.SubmittedBy = rawDesired.SubmittedBy
	} else {
		if dcl.StringCanonicalize(rawDesired.SubmittedBy, rawNew.SubmittedBy) {
			rawNew.SubmittedBy = rawDesired.SubmittedBy
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DriverInputResourceUri) && dcl.IsEmptyValueIndirect(rawDesired.DriverInputResourceUri) {
		rawNew.DriverInputResourceUri = rawDesired.DriverInputResourceUri
	} else {
		if dcl.StringCanonicalize(rawDesired.DriverInputResourceUri, rawNew.DriverInputResourceUri) {
			rawNew.DriverInputResourceUri = rawDesired.DriverInputResourceUri
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DriverOutputResourceUri) && dcl.IsEmptyValueIndirect(rawDesired.DriverOutputResourceUri) {
		rawNew.DriverOutputResourceUri = rawDesired.DriverOutputResourceUri
	} else {
		if dcl.StringCanonicalize(rawDesired.DriverOutputResourceUri, rawNew.DriverOutputResourceUri) {
			rawNew.DriverOutputResourceUri = rawDesired.DriverOutputResourceUri
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DriverControlFilesUri) && dcl.IsEmptyValueIndirect(rawDesired.DriverControlFilesUri) {
		rawNew.DriverControlFilesUri = rawDesired.DriverControlFilesUri
	} else {
		if dcl.StringCanonicalize(rawDesired.DriverControlFilesUri, rawNew.DriverControlFilesUri) {
			rawNew.DriverControlFilesUri = rawDesired.DriverControlFilesUri
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Interactive) && dcl.IsEmptyValueIndirect(rawDesired.Interactive) {
		rawNew.Interactive = rawDesired.Interactive
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Scheduling) && dcl.IsEmptyValueIndirect(rawDesired.Scheduling) {
		rawNew.Scheduling = rawDesired.Scheduling
	} else {
		rawNew.Scheduling = canonicalizeNewJobScheduling(c, rawDesired.Scheduling, rawNew.Scheduling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Done) && dcl.IsEmptyValueIndirect(rawDesired.Done) {
		rawNew.Done = rawDesired.Done
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DriverRunner) && dcl.IsEmptyValueIndirect(rawDesired.DriverRunner) {
		rawNew.DriverRunner = rawDesired.DriverRunner
	} else {
		rawNew.DriverRunner = canonicalizeNewJobDriverRunner(c, rawDesired.DriverRunner, rawNew.DriverRunner)
	}

	rawNew.Region = rawDesired.Region

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeJobReference(des, initial *JobReference, opts ...dcl.ApplyOption) *JobReference {
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
	if dcl.StringCanonicalize(des.JobId, initial.JobId) || dcl.IsZeroValue(des.JobId) {
		des.JobId = initial.JobId
	}

	return des
}

func canonicalizeNewJobReference(c *Client, des, nw *JobReference) *JobReference {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ProjectId, nw.ProjectId) || dcl.IsZeroValue(des.ProjectId) {
		nw.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.JobId, nw.JobId) || dcl.IsZeroValue(des.JobId) {
		nw.JobId = des.JobId
	}

	return nw
}

func canonicalizeNewJobReferenceSet(c *Client, des, nw []JobReference) []JobReference {
	if des == nil {
		return nw
	}
	var reorderedNew []JobReference
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobReference(c, &d, &n) {
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

func canonicalizeJobPlacement(des, initial *JobPlacement, opts ...dcl.ApplyOption) *JobPlacement {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ClusterName, initial.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
		des.ClusterName = initial.ClusterName
	}
	if dcl.StringCanonicalize(des.ClusterUuid, initial.ClusterUuid) || dcl.IsZeroValue(des.ClusterUuid) {
		des.ClusterUuid = initial.ClusterUuid
	}
	if dcl.IsZeroValue(des.ClusterLabels) {
		des.ClusterLabels = initial.ClusterLabels
	}

	return des
}

func canonicalizeNewJobPlacement(c *Client, des, nw *JobPlacement) *JobPlacement {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ClusterName, nw.ClusterName) || dcl.IsZeroValue(des.ClusterName) {
		nw.ClusterName = des.ClusterName
	}
	if dcl.StringCanonicalize(des.ClusterUuid, nw.ClusterUuid) || dcl.IsZeroValue(des.ClusterUuid) {
		nw.ClusterUuid = des.ClusterUuid
	}

	return nw
}

func canonicalizeNewJobPlacementSet(c *Client, des, nw []JobPlacement) []JobPlacement {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPlacement
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPlacement(c, &d, &n) {
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

func canonicalizeJobHadoopJob(des, initial *JobHadoopJob, opts ...dcl.ApplyOption) *JobHadoopJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MainJarFileUri, initial.MainJarFileUri) || dcl.IsZeroValue(des.MainJarFileUri) {
		des.MainJarFileUri = initial.MainJarFileUri
	}
	if dcl.StringCanonicalize(des.MainClass, initial.MainClass) || dcl.IsZeroValue(des.MainClass) {
		des.MainClass = initial.MainClass
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	}
	if dcl.IsZeroValue(des.JarFileUris) {
		des.JarFileUris = initial.JarFileUris
	}
	if dcl.IsZeroValue(des.FileUris) {
		des.FileUris = initial.FileUris
	}
	if dcl.IsZeroValue(des.ArchiveUris) {
		des.ArchiveUris = initial.ArchiveUris
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	des.LoggingConfig = canonicalizeJobHadoopJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewJobHadoopJob(c *Client, des, nw *JobHadoopJob) *JobHadoopJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MainJarFileUri, nw.MainJarFileUri) || dcl.IsZeroValue(des.MainJarFileUri) {
		nw.MainJarFileUri = des.MainJarFileUri
	}
	if dcl.StringCanonicalize(des.MainClass, nw.MainClass) || dcl.IsZeroValue(des.MainClass) {
		nw.MainClass = des.MainClass
	}
	nw.LoggingConfig = canonicalizeNewJobHadoopJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewJobHadoopJobSet(c *Client, des, nw []JobHadoopJob) []JobHadoopJob {
	if des == nil {
		return nw
	}
	var reorderedNew []JobHadoopJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobHadoopJob(c, &d, &n) {
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

func canonicalizeJobHadoopJobLoggingConfig(des, initial *JobHadoopJobLoggingConfig, opts ...dcl.ApplyOption) *JobHadoopJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DriverLogLevels) {
		des.DriverLogLevels = initial.DriverLogLevels
	}

	return des
}

func canonicalizeNewJobHadoopJobLoggingConfig(c *Client, des, nw *JobHadoopJobLoggingConfig) *JobHadoopJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobHadoopJobLoggingConfigSet(c *Client, des, nw []JobHadoopJobLoggingConfig) []JobHadoopJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []JobHadoopJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobHadoopJobLoggingConfig(c, &d, &n) {
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

func canonicalizeJobSparkJob(des, initial *JobSparkJob, opts ...dcl.ApplyOption) *JobSparkJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MainJarFileUri, initial.MainJarFileUri) || dcl.IsZeroValue(des.MainJarFileUri) {
		des.MainJarFileUri = initial.MainJarFileUri
	}
	if dcl.StringCanonicalize(des.MainClass, initial.MainClass) || dcl.IsZeroValue(des.MainClass) {
		des.MainClass = initial.MainClass
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	}
	if dcl.IsZeroValue(des.JarFileUris) {
		des.JarFileUris = initial.JarFileUris
	}
	if dcl.IsZeroValue(des.FileUris) {
		des.FileUris = initial.FileUris
	}
	if dcl.IsZeroValue(des.ArchiveUris) {
		des.ArchiveUris = initial.ArchiveUris
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	des.LoggingConfig = canonicalizeJobSparkJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewJobSparkJob(c *Client, des, nw *JobSparkJob) *JobSparkJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MainJarFileUri, nw.MainJarFileUri) || dcl.IsZeroValue(des.MainJarFileUri) {
		nw.MainJarFileUri = des.MainJarFileUri
	}
	if dcl.StringCanonicalize(des.MainClass, nw.MainClass) || dcl.IsZeroValue(des.MainClass) {
		nw.MainClass = des.MainClass
	}
	nw.LoggingConfig = canonicalizeNewJobSparkJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewJobSparkJobSet(c *Client, des, nw []JobSparkJob) []JobSparkJob {
	if des == nil {
		return nw
	}
	var reorderedNew []JobSparkJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobSparkJob(c, &d, &n) {
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

func canonicalizeJobSparkJobLoggingConfig(des, initial *JobSparkJobLoggingConfig, opts ...dcl.ApplyOption) *JobSparkJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DriverLogLevels) {
		des.DriverLogLevels = initial.DriverLogLevels
	}

	return des
}

func canonicalizeNewJobSparkJobLoggingConfig(c *Client, des, nw *JobSparkJobLoggingConfig) *JobSparkJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobSparkJobLoggingConfigSet(c *Client, des, nw []JobSparkJobLoggingConfig) []JobSparkJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []JobSparkJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobSparkJobLoggingConfig(c, &d, &n) {
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

func canonicalizeJobPysparkJob(des, initial *JobPysparkJob, opts ...dcl.ApplyOption) *JobPysparkJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MainPythonFileUri, initial.MainPythonFileUri) || dcl.IsZeroValue(des.MainPythonFileUri) {
		des.MainPythonFileUri = initial.MainPythonFileUri
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	}
	if dcl.IsZeroValue(des.PythonFileUris) {
		des.PythonFileUris = initial.PythonFileUris
	}
	if dcl.IsZeroValue(des.JarFileUris) {
		des.JarFileUris = initial.JarFileUris
	}
	if dcl.IsZeroValue(des.FileUris) {
		des.FileUris = initial.FileUris
	}
	if dcl.IsZeroValue(des.ArchiveUris) {
		des.ArchiveUris = initial.ArchiveUris
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	des.LoggingConfig = canonicalizeJobPysparkJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewJobPysparkJob(c *Client, des, nw *JobPysparkJob) *JobPysparkJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MainPythonFileUri, nw.MainPythonFileUri) || dcl.IsZeroValue(des.MainPythonFileUri) {
		nw.MainPythonFileUri = des.MainPythonFileUri
	}
	nw.LoggingConfig = canonicalizeNewJobPysparkJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewJobPysparkJobSet(c *Client, des, nw []JobPysparkJob) []JobPysparkJob {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPysparkJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPysparkJob(c, &d, &n) {
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

func canonicalizeJobPysparkJobLoggingConfig(des, initial *JobPysparkJobLoggingConfig, opts ...dcl.ApplyOption) *JobPysparkJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DriverLogLevels) {
		des.DriverLogLevels = initial.DriverLogLevels
	}

	return des
}

func canonicalizeNewJobPysparkJobLoggingConfig(c *Client, des, nw *JobPysparkJobLoggingConfig) *JobPysparkJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobPysparkJobLoggingConfigSet(c *Client, des, nw []JobPysparkJobLoggingConfig) []JobPysparkJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPysparkJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPysparkJobLoggingConfig(c, &d, &n) {
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

func canonicalizeJobHiveJob(des, initial *JobHiveJob, opts ...dcl.ApplyOption) *JobHiveJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.QueryFileUri, initial.QueryFileUri) || dcl.IsZeroValue(des.QueryFileUri) {
		des.QueryFileUri = initial.QueryFileUri
	}
	des.QueryList = canonicalizeJobHiveJobQueryList(des.QueryList, initial.QueryList, opts...)
	if dcl.IsZeroValue(des.ContinueOnFailure) {
		des.ContinueOnFailure = initial.ContinueOnFailure
	}
	if dcl.IsZeroValue(des.ScriptVariables) {
		des.ScriptVariables = initial.ScriptVariables
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	if dcl.IsZeroValue(des.JarFileUris) {
		des.JarFileUris = initial.JarFileUris
	}

	return des
}

func canonicalizeNewJobHiveJob(c *Client, des, nw *JobHiveJob) *JobHiveJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.QueryFileUri, nw.QueryFileUri) || dcl.IsZeroValue(des.QueryFileUri) {
		nw.QueryFileUri = des.QueryFileUri
	}
	nw.QueryList = canonicalizeNewJobHiveJobQueryList(c, des.QueryList, nw.QueryList)

	return nw
}

func canonicalizeNewJobHiveJobSet(c *Client, des, nw []JobHiveJob) []JobHiveJob {
	if des == nil {
		return nw
	}
	var reorderedNew []JobHiveJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobHiveJob(c, &d, &n) {
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

func canonicalizeJobHiveJobQueryList(des, initial *JobHiveJobQueryList, opts ...dcl.ApplyOption) *JobHiveJobQueryList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Queries) {
		des.Queries = initial.Queries
	}

	return des
}

func canonicalizeNewJobHiveJobQueryList(c *Client, des, nw *JobHiveJobQueryList) *JobHiveJobQueryList {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobHiveJobQueryListSet(c *Client, des, nw []JobHiveJobQueryList) []JobHiveJobQueryList {
	if des == nil {
		return nw
	}
	var reorderedNew []JobHiveJobQueryList
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobHiveJobQueryList(c, &d, &n) {
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

func canonicalizeJobPigJob(des, initial *JobPigJob, opts ...dcl.ApplyOption) *JobPigJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.QueryFileUri, initial.QueryFileUri) || dcl.IsZeroValue(des.QueryFileUri) {
		des.QueryFileUri = initial.QueryFileUri
	}
	des.QueryList = canonicalizeJobPigJobQueryList(des.QueryList, initial.QueryList, opts...)
	if dcl.IsZeroValue(des.ContinueOnFailure) {
		des.ContinueOnFailure = initial.ContinueOnFailure
	}
	if dcl.IsZeroValue(des.ScriptVariables) {
		des.ScriptVariables = initial.ScriptVariables
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	if dcl.IsZeroValue(des.JarFileUris) {
		des.JarFileUris = initial.JarFileUris
	}
	des.LoggingConfig = canonicalizeJobPigJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewJobPigJob(c *Client, des, nw *JobPigJob) *JobPigJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.QueryFileUri, nw.QueryFileUri) || dcl.IsZeroValue(des.QueryFileUri) {
		nw.QueryFileUri = des.QueryFileUri
	}
	nw.QueryList = canonicalizeNewJobPigJobQueryList(c, des.QueryList, nw.QueryList)
	nw.LoggingConfig = canonicalizeNewJobPigJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewJobPigJobSet(c *Client, des, nw []JobPigJob) []JobPigJob {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPigJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPigJob(c, &d, &n) {
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

func canonicalizeJobPigJobQueryList(des, initial *JobPigJobQueryList, opts ...dcl.ApplyOption) *JobPigJobQueryList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Queries) {
		des.Queries = initial.Queries
	}

	return des
}

func canonicalizeNewJobPigJobQueryList(c *Client, des, nw *JobPigJobQueryList) *JobPigJobQueryList {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobPigJobQueryListSet(c *Client, des, nw []JobPigJobQueryList) []JobPigJobQueryList {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPigJobQueryList
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPigJobQueryList(c, &d, &n) {
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

func canonicalizeJobPigJobLoggingConfig(des, initial *JobPigJobLoggingConfig, opts ...dcl.ApplyOption) *JobPigJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DriverLogLevels) {
		des.DriverLogLevels = initial.DriverLogLevels
	}

	return des
}

func canonicalizeNewJobPigJobLoggingConfig(c *Client, des, nw *JobPigJobLoggingConfig) *JobPigJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobPigJobLoggingConfigSet(c *Client, des, nw []JobPigJobLoggingConfig) []JobPigJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPigJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPigJobLoggingConfig(c, &d, &n) {
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

func canonicalizeJobSparkRJob(des, initial *JobSparkRJob, opts ...dcl.ApplyOption) *JobSparkRJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MainRFileUri, initial.MainRFileUri) || dcl.IsZeroValue(des.MainRFileUri) {
		des.MainRFileUri = initial.MainRFileUri
	}
	if dcl.IsZeroValue(des.Args) {
		des.Args = initial.Args
	}
	if dcl.IsZeroValue(des.FileUris) {
		des.FileUris = initial.FileUris
	}
	if dcl.IsZeroValue(des.ArchiveUris) {
		des.ArchiveUris = initial.ArchiveUris
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	des.LoggingConfig = canonicalizeJobSparkRJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewJobSparkRJob(c *Client, des, nw *JobSparkRJob) *JobSparkRJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MainRFileUri, nw.MainRFileUri) || dcl.IsZeroValue(des.MainRFileUri) {
		nw.MainRFileUri = des.MainRFileUri
	}
	nw.LoggingConfig = canonicalizeNewJobSparkRJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewJobSparkRJobSet(c *Client, des, nw []JobSparkRJob) []JobSparkRJob {
	if des == nil {
		return nw
	}
	var reorderedNew []JobSparkRJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobSparkRJob(c, &d, &n) {
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

func canonicalizeJobSparkRJobLoggingConfig(des, initial *JobSparkRJobLoggingConfig, opts ...dcl.ApplyOption) *JobSparkRJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DriverLogLevels) {
		des.DriverLogLevels = initial.DriverLogLevels
	}

	return des
}

func canonicalizeNewJobSparkRJobLoggingConfig(c *Client, des, nw *JobSparkRJobLoggingConfig) *JobSparkRJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobSparkRJobLoggingConfigSet(c *Client, des, nw []JobSparkRJobLoggingConfig) []JobSparkRJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []JobSparkRJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobSparkRJobLoggingConfig(c, &d, &n) {
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

func canonicalizeJobSparkSqlJob(des, initial *JobSparkSqlJob, opts ...dcl.ApplyOption) *JobSparkSqlJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.QueryFileUri, initial.QueryFileUri) || dcl.IsZeroValue(des.QueryFileUri) {
		des.QueryFileUri = initial.QueryFileUri
	}
	des.QueryList = canonicalizeJobSparkSqlJobQueryList(des.QueryList, initial.QueryList, opts...)
	if dcl.IsZeroValue(des.ScriptVariables) {
		des.ScriptVariables = initial.ScriptVariables
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	if dcl.IsZeroValue(des.JarFileUris) {
		des.JarFileUris = initial.JarFileUris
	}
	des.LoggingConfig = canonicalizeJobSparkSqlJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewJobSparkSqlJob(c *Client, des, nw *JobSparkSqlJob) *JobSparkSqlJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.QueryFileUri, nw.QueryFileUri) || dcl.IsZeroValue(des.QueryFileUri) {
		nw.QueryFileUri = des.QueryFileUri
	}
	nw.QueryList = canonicalizeNewJobSparkSqlJobQueryList(c, des.QueryList, nw.QueryList)
	nw.LoggingConfig = canonicalizeNewJobSparkSqlJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewJobSparkSqlJobSet(c *Client, des, nw []JobSparkSqlJob) []JobSparkSqlJob {
	if des == nil {
		return nw
	}
	var reorderedNew []JobSparkSqlJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobSparkSqlJob(c, &d, &n) {
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

func canonicalizeJobSparkSqlJobQueryList(des, initial *JobSparkSqlJobQueryList, opts ...dcl.ApplyOption) *JobSparkSqlJobQueryList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Queries) {
		des.Queries = initial.Queries
	}

	return des
}

func canonicalizeNewJobSparkSqlJobQueryList(c *Client, des, nw *JobSparkSqlJobQueryList) *JobSparkSqlJobQueryList {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobSparkSqlJobQueryListSet(c *Client, des, nw []JobSparkSqlJobQueryList) []JobSparkSqlJobQueryList {
	if des == nil {
		return nw
	}
	var reorderedNew []JobSparkSqlJobQueryList
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobSparkSqlJobQueryList(c, &d, &n) {
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

func canonicalizeJobSparkSqlJobLoggingConfig(des, initial *JobSparkSqlJobLoggingConfig, opts ...dcl.ApplyOption) *JobSparkSqlJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DriverLogLevels) {
		des.DriverLogLevels = initial.DriverLogLevels
	}

	return des
}

func canonicalizeNewJobSparkSqlJobLoggingConfig(c *Client, des, nw *JobSparkSqlJobLoggingConfig) *JobSparkSqlJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobSparkSqlJobLoggingConfigSet(c *Client, des, nw []JobSparkSqlJobLoggingConfig) []JobSparkSqlJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []JobSparkSqlJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobSparkSqlJobLoggingConfig(c, &d, &n) {
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

func canonicalizeJobPrestoJob(des, initial *JobPrestoJob, opts ...dcl.ApplyOption) *JobPrestoJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.QueryFileUri, initial.QueryFileUri) || dcl.IsZeroValue(des.QueryFileUri) {
		des.QueryFileUri = initial.QueryFileUri
	}
	des.QueryList = canonicalizeJobPrestoJobQueryList(des.QueryList, initial.QueryList, opts...)
	if dcl.IsZeroValue(des.ContinueOnFailure) {
		des.ContinueOnFailure = initial.ContinueOnFailure
	}
	if dcl.StringCanonicalize(des.OutputFormat, initial.OutputFormat) || dcl.IsZeroValue(des.OutputFormat) {
		des.OutputFormat = initial.OutputFormat
	}
	if dcl.IsZeroValue(des.ClientTags) {
		des.ClientTags = initial.ClientTags
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	des.LoggingConfig = canonicalizeJobPrestoJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewJobPrestoJob(c *Client, des, nw *JobPrestoJob) *JobPrestoJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.QueryFileUri, nw.QueryFileUri) || dcl.IsZeroValue(des.QueryFileUri) {
		nw.QueryFileUri = des.QueryFileUri
	}
	nw.QueryList = canonicalizeNewJobPrestoJobQueryList(c, des.QueryList, nw.QueryList)
	if dcl.StringCanonicalize(des.OutputFormat, nw.OutputFormat) || dcl.IsZeroValue(des.OutputFormat) {
		nw.OutputFormat = des.OutputFormat
	}
	nw.LoggingConfig = canonicalizeNewJobPrestoJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewJobPrestoJobSet(c *Client, des, nw []JobPrestoJob) []JobPrestoJob {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPrestoJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPrestoJob(c, &d, &n) {
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

func canonicalizeJobPrestoJobQueryList(des, initial *JobPrestoJobQueryList, opts ...dcl.ApplyOption) *JobPrestoJobQueryList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Queries) {
		des.Queries = initial.Queries
	}

	return des
}

func canonicalizeNewJobPrestoJobQueryList(c *Client, des, nw *JobPrestoJobQueryList) *JobPrestoJobQueryList {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobPrestoJobQueryListSet(c *Client, des, nw []JobPrestoJobQueryList) []JobPrestoJobQueryList {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPrestoJobQueryList
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPrestoJobQueryList(c, &d, &n) {
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

func canonicalizeJobPrestoJobLoggingConfig(des, initial *JobPrestoJobLoggingConfig, opts ...dcl.ApplyOption) *JobPrestoJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DriverLogLevels) {
		des.DriverLogLevels = initial.DriverLogLevels
	}

	return des
}

func canonicalizeNewJobPrestoJobLoggingConfig(c *Client, des, nw *JobPrestoJobLoggingConfig) *JobPrestoJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobPrestoJobLoggingConfigSet(c *Client, des, nw []JobPrestoJobLoggingConfig) []JobPrestoJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []JobPrestoJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobPrestoJobLoggingConfig(c, &d, &n) {
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

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.StringCanonicalize(des.Details, initial.Details) || dcl.IsZeroValue(des.Details) {
		des.Details = initial.Details
	}
	if dcl.IsZeroValue(des.StateStartTime) {
		des.StateStartTime = initial.StateStartTime
	}
	if dcl.IsZeroValue(des.Substate) {
		des.Substate = initial.Substate
	}

	return des
}

func canonicalizeNewJobStatus(c *Client, des, nw *JobStatus) *JobStatus {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Details, nw.Details) || dcl.IsZeroValue(des.Details) {
		nw.Details = des.Details
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

func canonicalizeJobStatusHistory(des, initial *JobStatusHistory, opts ...dcl.ApplyOption) *JobStatusHistory {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.StringCanonicalize(des.Details, initial.Details) || dcl.IsZeroValue(des.Details) {
		des.Details = initial.Details
	}
	if dcl.IsZeroValue(des.StateStartTime) {
		des.StateStartTime = initial.StateStartTime
	}
	if dcl.IsZeroValue(des.Substate) {
		des.Substate = initial.Substate
	}

	return des
}

func canonicalizeNewJobStatusHistory(c *Client, des, nw *JobStatusHistory) *JobStatusHistory {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Details, nw.Details) || dcl.IsZeroValue(des.Details) {
		nw.Details = des.Details
	}

	return nw
}

func canonicalizeNewJobStatusHistorySet(c *Client, des, nw []JobStatusHistory) []JobStatusHistory {
	if des == nil {
		return nw
	}
	var reorderedNew []JobStatusHistory
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobStatusHistory(c, &d, &n) {
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

func canonicalizeJobYarnApplications(des, initial *JobYarnApplications, opts ...dcl.ApplyOption) *JobYarnApplications {
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
	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.IsZeroValue(des.Progress) {
		des.Progress = initial.Progress
	}
	if dcl.StringCanonicalize(des.TrackingUrl, initial.TrackingUrl) || dcl.IsZeroValue(des.TrackingUrl) {
		des.TrackingUrl = initial.TrackingUrl
	}

	return des
}

func canonicalizeNewJobYarnApplications(c *Client, des, nw *JobYarnApplications) *JobYarnApplications {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.TrackingUrl, nw.TrackingUrl) || dcl.IsZeroValue(des.TrackingUrl) {
		nw.TrackingUrl = des.TrackingUrl
	}

	return nw
}

func canonicalizeNewJobYarnApplicationsSet(c *Client, des, nw []JobYarnApplications) []JobYarnApplications {
	if des == nil {
		return nw
	}
	var reorderedNew []JobYarnApplications
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobYarnApplications(c, &d, &n) {
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

func canonicalizeJobScheduling(des, initial *JobScheduling, opts ...dcl.ApplyOption) *JobScheduling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MaxFailuresPerHour) {
		des.MaxFailuresPerHour = initial.MaxFailuresPerHour
	}
	if dcl.IsZeroValue(des.MaxFailuresTotal) {
		des.MaxFailuresTotal = initial.MaxFailuresTotal
	}

	return des
}

func canonicalizeNewJobScheduling(c *Client, des, nw *JobScheduling) *JobScheduling {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobSchedulingSet(c *Client, des, nw []JobScheduling) []JobScheduling {
	if des == nil {
		return nw
	}
	var reorderedNew []JobScheduling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobScheduling(c, &d, &n) {
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

func canonicalizeJobDriverRunner(des, initial *JobDriverRunner, opts ...dcl.ApplyOption) *JobDriverRunner {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.MasterDriverRunner = canonicalizeJobDriverRunnerMasterDriverRunner(des.MasterDriverRunner, initial.MasterDriverRunner, opts...)
	des.YarnDriverRunner = canonicalizeJobDriverRunnerYarnDriverRunner(des.YarnDriverRunner, initial.YarnDriverRunner, opts...)

	return des
}

func canonicalizeNewJobDriverRunner(c *Client, des, nw *JobDriverRunner) *JobDriverRunner {
	if des == nil || nw == nil {
		return nw
	}

	nw.MasterDriverRunner = canonicalizeNewJobDriverRunnerMasterDriverRunner(c, des.MasterDriverRunner, nw.MasterDriverRunner)
	nw.YarnDriverRunner = canonicalizeNewJobDriverRunnerYarnDriverRunner(c, des.YarnDriverRunner, nw.YarnDriverRunner)

	return nw
}

func canonicalizeNewJobDriverRunnerSet(c *Client, des, nw []JobDriverRunner) []JobDriverRunner {
	if des == nil {
		return nw
	}
	var reorderedNew []JobDriverRunner
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobDriverRunner(c, &d, &n) {
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

func canonicalizeJobDriverRunnerMasterDriverRunner(des, initial *JobDriverRunnerMasterDriverRunner, opts ...dcl.ApplyOption) *JobDriverRunnerMasterDriverRunner {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	return des
}

func canonicalizeNewJobDriverRunnerMasterDriverRunner(c *Client, des, nw *JobDriverRunnerMasterDriverRunner) *JobDriverRunnerMasterDriverRunner {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobDriverRunnerMasterDriverRunnerSet(c *Client, des, nw []JobDriverRunnerMasterDriverRunner) []JobDriverRunnerMasterDriverRunner {
	if des == nil {
		return nw
	}
	var reorderedNew []JobDriverRunnerMasterDriverRunner
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobDriverRunnerMasterDriverRunner(c, &d, &n) {
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

func canonicalizeJobDriverRunnerYarnDriverRunner(des, initial *JobDriverRunnerYarnDriverRunner, opts ...dcl.ApplyOption) *JobDriverRunnerYarnDriverRunner {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MemoryMb) {
		des.MemoryMb = initial.MemoryMb
	}
	if dcl.IsZeroValue(des.Vcores) {
		des.Vcores = initial.Vcores
	}

	return des
}

func canonicalizeNewJobDriverRunnerYarnDriverRunner(c *Client, des, nw *JobDriverRunnerYarnDriverRunner) *JobDriverRunnerYarnDriverRunner {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewJobDriverRunnerYarnDriverRunnerSet(c *Client, des, nw []JobDriverRunnerYarnDriverRunner) []JobDriverRunnerYarnDriverRunner {
	if des == nil {
		return nw
	}
	var reorderedNew []JobDriverRunnerYarnDriverRunner
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareJobDriverRunnerYarnDriverRunner(c, &d, &n) {
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
	if compareJobReference(c, desired.Reference, actual.Reference) {
		c.Config.Logger.Infof("Detected diff in Reference.\nDESIRED: %v\nACTUAL: %v", desired.Reference, actual.Reference)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "Reference",
		})
	}
	if compareJobPlacement(c, desired.Placement, actual.Placement) {
		c.Config.Logger.Infof("Detected diff in Placement.\nDESIRED: %v\nACTUAL: %v", desired.Placement, actual.Placement)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "Placement",
		})
	}
	if compareJobHadoopJob(c, desired.HadoopJob, actual.HadoopJob) {
		c.Config.Logger.Infof("Detected diff in HadoopJob.\nDESIRED: %v\nACTUAL: %v", desired.HadoopJob, actual.HadoopJob)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "HadoopJob",
		})
	}
	if compareJobSparkJob(c, desired.SparkJob, actual.SparkJob) {
		c.Config.Logger.Infof("Detected diff in SparkJob.\nDESIRED: %v\nACTUAL: %v", desired.SparkJob, actual.SparkJob)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "SparkJob",
		})
	}
	if compareJobPysparkJob(c, desired.PysparkJob, actual.PysparkJob) {
		c.Config.Logger.Infof("Detected diff in PysparkJob.\nDESIRED: %v\nACTUAL: %v", desired.PysparkJob, actual.PysparkJob)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "PysparkJob",
		})
	}
	if compareJobHiveJob(c, desired.HiveJob, actual.HiveJob) {
		c.Config.Logger.Infof("Detected diff in HiveJob.\nDESIRED: %v\nACTUAL: %v", desired.HiveJob, actual.HiveJob)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "HiveJob",
		})
	}
	if compareJobPigJob(c, desired.PigJob, actual.PigJob) {
		c.Config.Logger.Infof("Detected diff in PigJob.\nDESIRED: %v\nACTUAL: %v", desired.PigJob, actual.PigJob)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "PigJob",
		})
	}
	if compareJobSparkRJob(c, desired.SparkRJob, actual.SparkRJob) {
		c.Config.Logger.Infof("Detected diff in SparkRJob.\nDESIRED: %v\nACTUAL: %v", desired.SparkRJob, actual.SparkRJob)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "SparkRJob",
		})
	}
	if compareJobSparkSqlJob(c, desired.SparkSqlJob, actual.SparkSqlJob) {
		c.Config.Logger.Infof("Detected diff in SparkSqlJob.\nDESIRED: %v\nACTUAL: %v", desired.SparkSqlJob, actual.SparkSqlJob)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "SparkSqlJob",
		})
	}
	if compareJobPrestoJob(c, desired.PrestoJob, actual.PrestoJob) {
		c.Config.Logger.Infof("Detected diff in PrestoJob.\nDESIRED: %v\nACTUAL: %v", desired.PrestoJob, actual.PrestoJob)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "PrestoJob",
		})
	}
	if !reflect.DeepEqual(desired.Interactive, actual.Interactive) {
		c.Config.Logger.Infof("Detected diff in Interactive.\nDESIRED: %v\nACTUAL: %v", desired.Interactive, actual.Interactive)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "Interactive",
		})
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, jobDiff{
			UpdateOp:  &updateJobPatchOperation{},
			FieldName: "Labels",
		})

	}
	if compareJobScheduling(c, desired.Scheduling, actual.Scheduling) {
		c.Config.Logger.Infof("Detected diff in Scheduling.\nDESIRED: %v\nACTUAL: %v", desired.Scheduling, actual.Scheduling)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "Scheduling",
		})
	}
	if compareJobDriverRunner(c, desired.DriverRunner, actual.DriverRunner) {
		c.Config.Logger.Infof("Detected diff in DriverRunner.\nDESIRED: %v\nACTUAL: %v", desired.DriverRunner, actual.DriverRunner)
		diffs = append(diffs, jobDiff{
			RequiresRecreate: true,
			FieldName:        "DriverRunner",
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
func compareJobReference(c *Client, desired, actual *JobReference) bool {
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
	if actual.JobId == nil && desired.JobId != nil && !dcl.IsEmptyValueIndirect(desired.JobId) {
		c.Config.Logger.Infof("desired JobId %s - but actually nil", dcl.SprintResource(desired.JobId))
		return true
	}
	if !dcl.StringCanonicalize(desired.JobId, actual.JobId) && !dcl.IsZeroValue(desired.JobId) {
		c.Config.Logger.Infof("Diff in JobId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JobId), dcl.SprintResource(actual.JobId))
		return true
	}
	return false
}

func compareJobReferenceSlice(c *Client, desired, actual []JobReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobReference, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobReference(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobReference, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobReferenceMap(c *Client, desired, actual map[string]JobReference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobReference, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobReference, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobReference(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobReference, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobPlacement(c *Client, desired, actual *JobPlacement) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ClusterName == nil && desired.ClusterName != nil && !dcl.IsEmptyValueIndirect(desired.ClusterName) {
		c.Config.Logger.Infof("desired ClusterName %s - but actually nil", dcl.SprintResource(desired.ClusterName))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClusterName, actual.ClusterName) && !dcl.IsZeroValue(desired.ClusterName) {
		c.Config.Logger.Infof("Diff in ClusterName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterName), dcl.SprintResource(actual.ClusterName))
		return true
	}
	if actual.ClusterLabels == nil && desired.ClusterLabels != nil && !dcl.IsEmptyValueIndirect(desired.ClusterLabels) {
		c.Config.Logger.Infof("desired ClusterLabels %s - but actually nil", dcl.SprintResource(desired.ClusterLabels))
		return true
	}
	if !dcl.MapEquals(desired.ClusterLabels, actual.ClusterLabels, []string(nil)) && !dcl.IsZeroValue(desired.ClusterLabels) {
		c.Config.Logger.Infof("Diff in ClusterLabels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterLabels), dcl.SprintResource(actual.ClusterLabels))
		return true
	}
	return false
}

func compareJobPlacementSlice(c *Client, desired, actual []JobPlacement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPlacement, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPlacement(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPlacement, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPlacementMap(c *Client, desired, actual map[string]JobPlacement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPlacement, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPlacement, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPlacement(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPlacement, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobHadoopJob(c *Client, desired, actual *JobHadoopJob) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MainJarFileUri == nil && desired.MainJarFileUri != nil && !dcl.IsEmptyValueIndirect(desired.MainJarFileUri) {
		c.Config.Logger.Infof("desired MainJarFileUri %s - but actually nil", dcl.SprintResource(desired.MainJarFileUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.MainJarFileUri, actual.MainJarFileUri) && !dcl.IsZeroValue(desired.MainJarFileUri) {
		c.Config.Logger.Infof("Diff in MainJarFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainJarFileUri), dcl.SprintResource(actual.MainJarFileUri))
		return true
	}
	if actual.MainClass == nil && desired.MainClass != nil && !dcl.IsEmptyValueIndirect(desired.MainClass) {
		c.Config.Logger.Infof("desired MainClass %s - but actually nil", dcl.SprintResource(desired.MainClass))
		return true
	}
	if !dcl.StringCanonicalize(desired.MainClass, actual.MainClass) && !dcl.IsZeroValue(desired.MainClass) {
		c.Config.Logger.Infof("Diff in MainClass. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainClass), dcl.SprintResource(actual.MainClass))
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
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
		return true
	}
	if actual.FileUris == nil && desired.FileUris != nil && !dcl.IsEmptyValueIndirect(desired.FileUris) {
		c.Config.Logger.Infof("desired FileUris %s - but actually nil", dcl.SprintResource(desired.FileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.FileUris, actual.FileUris) && !dcl.IsZeroValue(desired.FileUris) {
		c.Config.Logger.Infof("Diff in FileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileUris), dcl.SprintResource(actual.FileUris))
		return true
	}
	if actual.ArchiveUris == nil && desired.ArchiveUris != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveUris) {
		c.Config.Logger.Infof("desired ArchiveUris %s - but actually nil", dcl.SprintResource(desired.ArchiveUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.ArchiveUris, actual.ArchiveUris) && !dcl.IsZeroValue(desired.ArchiveUris) {
		c.Config.Logger.Infof("Diff in ArchiveUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveUris), dcl.SprintResource(actual.ArchiveUris))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if !dcl.MapEquals(desired.Properties, actual.Properties, []string(nil)) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.LoggingConfig == nil && desired.LoggingConfig != nil && !dcl.IsEmptyValueIndirect(desired.LoggingConfig) {
		c.Config.Logger.Infof("desired LoggingConfig %s - but actually nil", dcl.SprintResource(desired.LoggingConfig))
		return true
	}
	if compareJobHadoopJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareJobHadoopJobSlice(c *Client, desired, actual []JobHadoopJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHadoopJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobHadoopJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobHadoopJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHadoopJobMap(c *Client, desired, actual map[string]JobHadoopJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHadoopJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobHadoopJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobHadoopJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobHadoopJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobHadoopJobLoggingConfig(c *Client, desired, actual *JobHadoopJobLoggingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DriverLogLevels == nil && desired.DriverLogLevels != nil && !dcl.IsEmptyValueIndirect(desired.DriverLogLevels) {
		c.Config.Logger.Infof("desired DriverLogLevels %s - but actually nil", dcl.SprintResource(desired.DriverLogLevels))
		return true
	}
	if !dcl.MapEquals(desired.DriverLogLevels, actual.DriverLogLevels, []string(nil)) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}

func compareJobHadoopJobLoggingConfigSlice(c *Client, desired, actual []JobHadoopJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHadoopJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobHadoopJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobHadoopJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHadoopJobLoggingConfigMap(c *Client, desired, actual map[string]JobHadoopJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHadoopJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobHadoopJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobHadoopJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobHadoopJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobSparkJob(c *Client, desired, actual *JobSparkJob) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MainJarFileUri == nil && desired.MainJarFileUri != nil && !dcl.IsEmptyValueIndirect(desired.MainJarFileUri) {
		c.Config.Logger.Infof("desired MainJarFileUri %s - but actually nil", dcl.SprintResource(desired.MainJarFileUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.MainJarFileUri, actual.MainJarFileUri) && !dcl.IsZeroValue(desired.MainJarFileUri) {
		c.Config.Logger.Infof("Diff in MainJarFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainJarFileUri), dcl.SprintResource(actual.MainJarFileUri))
		return true
	}
	if actual.MainClass == nil && desired.MainClass != nil && !dcl.IsEmptyValueIndirect(desired.MainClass) {
		c.Config.Logger.Infof("desired MainClass %s - but actually nil", dcl.SprintResource(desired.MainClass))
		return true
	}
	if !dcl.StringCanonicalize(desired.MainClass, actual.MainClass) && !dcl.IsZeroValue(desired.MainClass) {
		c.Config.Logger.Infof("Diff in MainClass. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainClass), dcl.SprintResource(actual.MainClass))
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
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
		return true
	}
	if actual.FileUris == nil && desired.FileUris != nil && !dcl.IsEmptyValueIndirect(desired.FileUris) {
		c.Config.Logger.Infof("desired FileUris %s - but actually nil", dcl.SprintResource(desired.FileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.FileUris, actual.FileUris) && !dcl.IsZeroValue(desired.FileUris) {
		c.Config.Logger.Infof("Diff in FileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileUris), dcl.SprintResource(actual.FileUris))
		return true
	}
	if actual.ArchiveUris == nil && desired.ArchiveUris != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveUris) {
		c.Config.Logger.Infof("desired ArchiveUris %s - but actually nil", dcl.SprintResource(desired.ArchiveUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.ArchiveUris, actual.ArchiveUris) && !dcl.IsZeroValue(desired.ArchiveUris) {
		c.Config.Logger.Infof("Diff in ArchiveUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveUris), dcl.SprintResource(actual.ArchiveUris))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if !dcl.MapEquals(desired.Properties, actual.Properties, []string(nil)) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.LoggingConfig == nil && desired.LoggingConfig != nil && !dcl.IsEmptyValueIndirect(desired.LoggingConfig) {
		c.Config.Logger.Infof("desired LoggingConfig %s - but actually nil", dcl.SprintResource(desired.LoggingConfig))
		return true
	}
	if compareJobSparkJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareJobSparkJobSlice(c *Client, desired, actual []JobSparkJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobSparkJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobSparkJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobSparkJobMap(c *Client, desired, actual map[string]JobSparkJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobSparkJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobSparkJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobSparkJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobSparkJobLoggingConfig(c *Client, desired, actual *JobSparkJobLoggingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DriverLogLevels == nil && desired.DriverLogLevels != nil && !dcl.IsEmptyValueIndirect(desired.DriverLogLevels) {
		c.Config.Logger.Infof("desired DriverLogLevels %s - but actually nil", dcl.SprintResource(desired.DriverLogLevels))
		return true
	}
	if !dcl.MapEquals(desired.DriverLogLevels, actual.DriverLogLevels, []string(nil)) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}

func compareJobSparkJobLoggingConfigSlice(c *Client, desired, actual []JobSparkJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobSparkJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobSparkJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobSparkJobLoggingConfigMap(c *Client, desired, actual map[string]JobSparkJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobSparkJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobSparkJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobSparkJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobPysparkJob(c *Client, desired, actual *JobPysparkJob) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MainPythonFileUri == nil && desired.MainPythonFileUri != nil && !dcl.IsEmptyValueIndirect(desired.MainPythonFileUri) {
		c.Config.Logger.Infof("desired MainPythonFileUri %s - but actually nil", dcl.SprintResource(desired.MainPythonFileUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.MainPythonFileUri, actual.MainPythonFileUri) && !dcl.IsZeroValue(desired.MainPythonFileUri) {
		c.Config.Logger.Infof("Diff in MainPythonFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainPythonFileUri), dcl.SprintResource(actual.MainPythonFileUri))
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
	if actual.PythonFileUris == nil && desired.PythonFileUris != nil && !dcl.IsEmptyValueIndirect(desired.PythonFileUris) {
		c.Config.Logger.Infof("desired PythonFileUris %s - but actually nil", dcl.SprintResource(desired.PythonFileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.PythonFileUris, actual.PythonFileUris) && !dcl.IsZeroValue(desired.PythonFileUris) {
		c.Config.Logger.Infof("Diff in PythonFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PythonFileUris), dcl.SprintResource(actual.PythonFileUris))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
		return true
	}
	if actual.FileUris == nil && desired.FileUris != nil && !dcl.IsEmptyValueIndirect(desired.FileUris) {
		c.Config.Logger.Infof("desired FileUris %s - but actually nil", dcl.SprintResource(desired.FileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.FileUris, actual.FileUris) && !dcl.IsZeroValue(desired.FileUris) {
		c.Config.Logger.Infof("Diff in FileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileUris), dcl.SprintResource(actual.FileUris))
		return true
	}
	if actual.ArchiveUris == nil && desired.ArchiveUris != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveUris) {
		c.Config.Logger.Infof("desired ArchiveUris %s - but actually nil", dcl.SprintResource(desired.ArchiveUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.ArchiveUris, actual.ArchiveUris) && !dcl.IsZeroValue(desired.ArchiveUris) {
		c.Config.Logger.Infof("Diff in ArchiveUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveUris), dcl.SprintResource(actual.ArchiveUris))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if !dcl.MapEquals(desired.Properties, actual.Properties, []string(nil)) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.LoggingConfig == nil && desired.LoggingConfig != nil && !dcl.IsEmptyValueIndirect(desired.LoggingConfig) {
		c.Config.Logger.Infof("desired LoggingConfig %s - but actually nil", dcl.SprintResource(desired.LoggingConfig))
		return true
	}
	if compareJobPysparkJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareJobPysparkJobSlice(c *Client, desired, actual []JobPysparkJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPysparkJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPysparkJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPysparkJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPysparkJobMap(c *Client, desired, actual map[string]JobPysparkJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPysparkJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPysparkJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPysparkJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPysparkJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobPysparkJobLoggingConfig(c *Client, desired, actual *JobPysparkJobLoggingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DriverLogLevels == nil && desired.DriverLogLevels != nil && !dcl.IsEmptyValueIndirect(desired.DriverLogLevels) {
		c.Config.Logger.Infof("desired DriverLogLevels %s - but actually nil", dcl.SprintResource(desired.DriverLogLevels))
		return true
	}
	if !dcl.MapEquals(desired.DriverLogLevels, actual.DriverLogLevels, []string(nil)) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}

func compareJobPysparkJobLoggingConfigSlice(c *Client, desired, actual []JobPysparkJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPysparkJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPysparkJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPysparkJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPysparkJobLoggingConfigMap(c *Client, desired, actual map[string]JobPysparkJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPysparkJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPysparkJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPysparkJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPysparkJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobHiveJob(c *Client, desired, actual *JobHiveJob) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.QueryFileUri == nil && desired.QueryFileUri != nil && !dcl.IsEmptyValueIndirect(desired.QueryFileUri) {
		c.Config.Logger.Infof("desired QueryFileUri %s - but actually nil", dcl.SprintResource(desired.QueryFileUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.QueryFileUri, actual.QueryFileUri) && !dcl.IsZeroValue(desired.QueryFileUri) {
		c.Config.Logger.Infof("Diff in QueryFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryFileUri), dcl.SprintResource(actual.QueryFileUri))
		return true
	}
	if actual.QueryList == nil && desired.QueryList != nil && !dcl.IsEmptyValueIndirect(desired.QueryList) {
		c.Config.Logger.Infof("desired QueryList %s - but actually nil", dcl.SprintResource(desired.QueryList))
		return true
	}
	if compareJobHiveJobQueryList(c, desired.QueryList, actual.QueryList) && !dcl.IsZeroValue(desired.QueryList) {
		c.Config.Logger.Infof("Diff in QueryList. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryList), dcl.SprintResource(actual.QueryList))
		return true
	}
	if actual.ContinueOnFailure == nil && desired.ContinueOnFailure != nil && !dcl.IsEmptyValueIndirect(desired.ContinueOnFailure) {
		c.Config.Logger.Infof("desired ContinueOnFailure %s - but actually nil", dcl.SprintResource(desired.ContinueOnFailure))
		return true
	}
	if !reflect.DeepEqual(desired.ContinueOnFailure, actual.ContinueOnFailure) && !dcl.IsZeroValue(desired.ContinueOnFailure) {
		c.Config.Logger.Infof("Diff in ContinueOnFailure. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContinueOnFailure), dcl.SprintResource(actual.ContinueOnFailure))
		return true
	}
	if actual.ScriptVariables == nil && desired.ScriptVariables != nil && !dcl.IsEmptyValueIndirect(desired.ScriptVariables) {
		c.Config.Logger.Infof("desired ScriptVariables %s - but actually nil", dcl.SprintResource(desired.ScriptVariables))
		return true
	}
	if !dcl.MapEquals(desired.ScriptVariables, actual.ScriptVariables, []string(nil)) && !dcl.IsZeroValue(desired.ScriptVariables) {
		c.Config.Logger.Infof("Diff in ScriptVariables. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScriptVariables), dcl.SprintResource(actual.ScriptVariables))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if !dcl.MapEquals(desired.Properties, actual.Properties, []string(nil)) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
		return true
	}
	return false
}

func compareJobHiveJobSlice(c *Client, desired, actual []JobHiveJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHiveJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobHiveJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobHiveJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHiveJobMap(c *Client, desired, actual map[string]JobHiveJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHiveJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobHiveJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobHiveJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobHiveJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobHiveJobQueryList(c *Client, desired, actual *JobHiveJobQueryList) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Queries == nil && desired.Queries != nil && !dcl.IsEmptyValueIndirect(desired.Queries) {
		c.Config.Logger.Infof("desired Queries %s - but actually nil", dcl.SprintResource(desired.Queries))
		return true
	}
	if !dcl.StringSliceEquals(desired.Queries, actual.Queries) && !dcl.IsZeroValue(desired.Queries) {
		c.Config.Logger.Infof("Diff in Queries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Queries), dcl.SprintResource(actual.Queries))
		return true
	}
	return false
}

func compareJobHiveJobQueryListSlice(c *Client, desired, actual []JobHiveJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHiveJobQueryList, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobHiveJobQueryList(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobHiveJobQueryList, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobHiveJobQueryListMap(c *Client, desired, actual map[string]JobHiveJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobHiveJobQueryList, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobHiveJobQueryList, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobHiveJobQueryList(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobHiveJobQueryList, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobPigJob(c *Client, desired, actual *JobPigJob) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.QueryFileUri == nil && desired.QueryFileUri != nil && !dcl.IsEmptyValueIndirect(desired.QueryFileUri) {
		c.Config.Logger.Infof("desired QueryFileUri %s - but actually nil", dcl.SprintResource(desired.QueryFileUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.QueryFileUri, actual.QueryFileUri) && !dcl.IsZeroValue(desired.QueryFileUri) {
		c.Config.Logger.Infof("Diff in QueryFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryFileUri), dcl.SprintResource(actual.QueryFileUri))
		return true
	}
	if actual.QueryList == nil && desired.QueryList != nil && !dcl.IsEmptyValueIndirect(desired.QueryList) {
		c.Config.Logger.Infof("desired QueryList %s - but actually nil", dcl.SprintResource(desired.QueryList))
		return true
	}
	if compareJobPigJobQueryList(c, desired.QueryList, actual.QueryList) && !dcl.IsZeroValue(desired.QueryList) {
		c.Config.Logger.Infof("Diff in QueryList. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryList), dcl.SprintResource(actual.QueryList))
		return true
	}
	if actual.ContinueOnFailure == nil && desired.ContinueOnFailure != nil && !dcl.IsEmptyValueIndirect(desired.ContinueOnFailure) {
		c.Config.Logger.Infof("desired ContinueOnFailure %s - but actually nil", dcl.SprintResource(desired.ContinueOnFailure))
		return true
	}
	if !reflect.DeepEqual(desired.ContinueOnFailure, actual.ContinueOnFailure) && !dcl.IsZeroValue(desired.ContinueOnFailure) {
		c.Config.Logger.Infof("Diff in ContinueOnFailure. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContinueOnFailure), dcl.SprintResource(actual.ContinueOnFailure))
		return true
	}
	if actual.ScriptVariables == nil && desired.ScriptVariables != nil && !dcl.IsEmptyValueIndirect(desired.ScriptVariables) {
		c.Config.Logger.Infof("desired ScriptVariables %s - but actually nil", dcl.SprintResource(desired.ScriptVariables))
		return true
	}
	if !dcl.MapEquals(desired.ScriptVariables, actual.ScriptVariables, []string(nil)) && !dcl.IsZeroValue(desired.ScriptVariables) {
		c.Config.Logger.Infof("Diff in ScriptVariables. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScriptVariables), dcl.SprintResource(actual.ScriptVariables))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if !dcl.MapEquals(desired.Properties, actual.Properties, []string(nil)) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
		return true
	}
	if actual.LoggingConfig == nil && desired.LoggingConfig != nil && !dcl.IsEmptyValueIndirect(desired.LoggingConfig) {
		c.Config.Logger.Infof("desired LoggingConfig %s - but actually nil", dcl.SprintResource(desired.LoggingConfig))
		return true
	}
	if compareJobPigJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareJobPigJobSlice(c *Client, desired, actual []JobPigJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPigJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPigJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPigJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPigJobMap(c *Client, desired, actual map[string]JobPigJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPigJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPigJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPigJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPigJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobPigJobQueryList(c *Client, desired, actual *JobPigJobQueryList) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Queries == nil && desired.Queries != nil && !dcl.IsEmptyValueIndirect(desired.Queries) {
		c.Config.Logger.Infof("desired Queries %s - but actually nil", dcl.SprintResource(desired.Queries))
		return true
	}
	if !dcl.StringSliceEquals(desired.Queries, actual.Queries) && !dcl.IsZeroValue(desired.Queries) {
		c.Config.Logger.Infof("Diff in Queries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Queries), dcl.SprintResource(actual.Queries))
		return true
	}
	return false
}

func compareJobPigJobQueryListSlice(c *Client, desired, actual []JobPigJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPigJobQueryList, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPigJobQueryList(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPigJobQueryList, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPigJobQueryListMap(c *Client, desired, actual map[string]JobPigJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPigJobQueryList, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPigJobQueryList, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPigJobQueryList(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPigJobQueryList, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobPigJobLoggingConfig(c *Client, desired, actual *JobPigJobLoggingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DriverLogLevels == nil && desired.DriverLogLevels != nil && !dcl.IsEmptyValueIndirect(desired.DriverLogLevels) {
		c.Config.Logger.Infof("desired DriverLogLevels %s - but actually nil", dcl.SprintResource(desired.DriverLogLevels))
		return true
	}
	if !dcl.MapEquals(desired.DriverLogLevels, actual.DriverLogLevels, []string(nil)) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}

func compareJobPigJobLoggingConfigSlice(c *Client, desired, actual []JobPigJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPigJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPigJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPigJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPigJobLoggingConfigMap(c *Client, desired, actual map[string]JobPigJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPigJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPigJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPigJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPigJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobSparkRJob(c *Client, desired, actual *JobSparkRJob) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MainRFileUri == nil && desired.MainRFileUri != nil && !dcl.IsEmptyValueIndirect(desired.MainRFileUri) {
		c.Config.Logger.Infof("desired MainRFileUri %s - but actually nil", dcl.SprintResource(desired.MainRFileUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.MainRFileUri, actual.MainRFileUri) && !dcl.IsZeroValue(desired.MainRFileUri) {
		c.Config.Logger.Infof("Diff in MainRFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainRFileUri), dcl.SprintResource(actual.MainRFileUri))
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
	if actual.FileUris == nil && desired.FileUris != nil && !dcl.IsEmptyValueIndirect(desired.FileUris) {
		c.Config.Logger.Infof("desired FileUris %s - but actually nil", dcl.SprintResource(desired.FileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.FileUris, actual.FileUris) && !dcl.IsZeroValue(desired.FileUris) {
		c.Config.Logger.Infof("Diff in FileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileUris), dcl.SprintResource(actual.FileUris))
		return true
	}
	if actual.ArchiveUris == nil && desired.ArchiveUris != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveUris) {
		c.Config.Logger.Infof("desired ArchiveUris %s - but actually nil", dcl.SprintResource(desired.ArchiveUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.ArchiveUris, actual.ArchiveUris) && !dcl.IsZeroValue(desired.ArchiveUris) {
		c.Config.Logger.Infof("Diff in ArchiveUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveUris), dcl.SprintResource(actual.ArchiveUris))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if !dcl.MapEquals(desired.Properties, actual.Properties, []string(nil)) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.LoggingConfig == nil && desired.LoggingConfig != nil && !dcl.IsEmptyValueIndirect(desired.LoggingConfig) {
		c.Config.Logger.Infof("desired LoggingConfig %s - but actually nil", dcl.SprintResource(desired.LoggingConfig))
		return true
	}
	if compareJobSparkRJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareJobSparkRJobSlice(c *Client, desired, actual []JobSparkRJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkRJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobSparkRJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobSparkRJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobSparkRJobMap(c *Client, desired, actual map[string]JobSparkRJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkRJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobSparkRJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobSparkRJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobSparkRJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobSparkRJobLoggingConfig(c *Client, desired, actual *JobSparkRJobLoggingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DriverLogLevels == nil && desired.DriverLogLevels != nil && !dcl.IsEmptyValueIndirect(desired.DriverLogLevels) {
		c.Config.Logger.Infof("desired DriverLogLevels %s - but actually nil", dcl.SprintResource(desired.DriverLogLevels))
		return true
	}
	if !dcl.MapEquals(desired.DriverLogLevels, actual.DriverLogLevels, []string(nil)) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}

func compareJobSparkRJobLoggingConfigSlice(c *Client, desired, actual []JobSparkRJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkRJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobSparkRJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobSparkRJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobSparkRJobLoggingConfigMap(c *Client, desired, actual map[string]JobSparkRJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkRJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobSparkRJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobSparkRJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobSparkRJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobSparkSqlJob(c *Client, desired, actual *JobSparkSqlJob) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.QueryFileUri == nil && desired.QueryFileUri != nil && !dcl.IsEmptyValueIndirect(desired.QueryFileUri) {
		c.Config.Logger.Infof("desired QueryFileUri %s - but actually nil", dcl.SprintResource(desired.QueryFileUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.QueryFileUri, actual.QueryFileUri) && !dcl.IsZeroValue(desired.QueryFileUri) {
		c.Config.Logger.Infof("Diff in QueryFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryFileUri), dcl.SprintResource(actual.QueryFileUri))
		return true
	}
	if actual.QueryList == nil && desired.QueryList != nil && !dcl.IsEmptyValueIndirect(desired.QueryList) {
		c.Config.Logger.Infof("desired QueryList %s - but actually nil", dcl.SprintResource(desired.QueryList))
		return true
	}
	if compareJobSparkSqlJobQueryList(c, desired.QueryList, actual.QueryList) && !dcl.IsZeroValue(desired.QueryList) {
		c.Config.Logger.Infof("Diff in QueryList. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryList), dcl.SprintResource(actual.QueryList))
		return true
	}
	if actual.ScriptVariables == nil && desired.ScriptVariables != nil && !dcl.IsEmptyValueIndirect(desired.ScriptVariables) {
		c.Config.Logger.Infof("desired ScriptVariables %s - but actually nil", dcl.SprintResource(desired.ScriptVariables))
		return true
	}
	if !dcl.MapEquals(desired.ScriptVariables, actual.ScriptVariables, []string(nil)) && !dcl.IsZeroValue(desired.ScriptVariables) {
		c.Config.Logger.Infof("Diff in ScriptVariables. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScriptVariables), dcl.SprintResource(actual.ScriptVariables))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if !dcl.MapEquals(desired.Properties, actual.Properties, []string(nil)) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.StringSliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
		return true
	}
	if actual.LoggingConfig == nil && desired.LoggingConfig != nil && !dcl.IsEmptyValueIndirect(desired.LoggingConfig) {
		c.Config.Logger.Infof("desired LoggingConfig %s - but actually nil", dcl.SprintResource(desired.LoggingConfig))
		return true
	}
	if compareJobSparkSqlJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareJobSparkSqlJobSlice(c *Client, desired, actual []JobSparkSqlJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkSqlJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobSparkSqlJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobSparkSqlJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobSparkSqlJobMap(c *Client, desired, actual map[string]JobSparkSqlJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkSqlJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobSparkSqlJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobSparkSqlJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobSparkSqlJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobSparkSqlJobQueryList(c *Client, desired, actual *JobSparkSqlJobQueryList) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Queries == nil && desired.Queries != nil && !dcl.IsEmptyValueIndirect(desired.Queries) {
		c.Config.Logger.Infof("desired Queries %s - but actually nil", dcl.SprintResource(desired.Queries))
		return true
	}
	if !dcl.StringSliceEquals(desired.Queries, actual.Queries) && !dcl.IsZeroValue(desired.Queries) {
		c.Config.Logger.Infof("Diff in Queries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Queries), dcl.SprintResource(actual.Queries))
		return true
	}
	return false
}

func compareJobSparkSqlJobQueryListSlice(c *Client, desired, actual []JobSparkSqlJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkSqlJobQueryList, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobSparkSqlJobQueryList(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobSparkSqlJobQueryList, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobSparkSqlJobQueryListMap(c *Client, desired, actual map[string]JobSparkSqlJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkSqlJobQueryList, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobSparkSqlJobQueryList, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobSparkSqlJobQueryList(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobSparkSqlJobQueryList, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobSparkSqlJobLoggingConfig(c *Client, desired, actual *JobSparkSqlJobLoggingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DriverLogLevels == nil && desired.DriverLogLevels != nil && !dcl.IsEmptyValueIndirect(desired.DriverLogLevels) {
		c.Config.Logger.Infof("desired DriverLogLevels %s - but actually nil", dcl.SprintResource(desired.DriverLogLevels))
		return true
	}
	if !dcl.MapEquals(desired.DriverLogLevels, actual.DriverLogLevels, []string(nil)) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}

func compareJobSparkSqlJobLoggingConfigSlice(c *Client, desired, actual []JobSparkSqlJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkSqlJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobSparkSqlJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobSparkSqlJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobSparkSqlJobLoggingConfigMap(c *Client, desired, actual map[string]JobSparkSqlJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobSparkSqlJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobSparkSqlJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobSparkSqlJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobSparkSqlJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobPrestoJob(c *Client, desired, actual *JobPrestoJob) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.QueryFileUri == nil && desired.QueryFileUri != nil && !dcl.IsEmptyValueIndirect(desired.QueryFileUri) {
		c.Config.Logger.Infof("desired QueryFileUri %s - but actually nil", dcl.SprintResource(desired.QueryFileUri))
		return true
	}
	if !dcl.StringCanonicalize(desired.QueryFileUri, actual.QueryFileUri) && !dcl.IsZeroValue(desired.QueryFileUri) {
		c.Config.Logger.Infof("Diff in QueryFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryFileUri), dcl.SprintResource(actual.QueryFileUri))
		return true
	}
	if actual.QueryList == nil && desired.QueryList != nil && !dcl.IsEmptyValueIndirect(desired.QueryList) {
		c.Config.Logger.Infof("desired QueryList %s - but actually nil", dcl.SprintResource(desired.QueryList))
		return true
	}
	if compareJobPrestoJobQueryList(c, desired.QueryList, actual.QueryList) && !dcl.IsZeroValue(desired.QueryList) {
		c.Config.Logger.Infof("Diff in QueryList. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryList), dcl.SprintResource(actual.QueryList))
		return true
	}
	if actual.ContinueOnFailure == nil && desired.ContinueOnFailure != nil && !dcl.IsEmptyValueIndirect(desired.ContinueOnFailure) {
		c.Config.Logger.Infof("desired ContinueOnFailure %s - but actually nil", dcl.SprintResource(desired.ContinueOnFailure))
		return true
	}
	if !reflect.DeepEqual(desired.ContinueOnFailure, actual.ContinueOnFailure) && !dcl.IsZeroValue(desired.ContinueOnFailure) {
		c.Config.Logger.Infof("Diff in ContinueOnFailure. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContinueOnFailure), dcl.SprintResource(actual.ContinueOnFailure))
		return true
	}
	if actual.OutputFormat == nil && desired.OutputFormat != nil && !dcl.IsEmptyValueIndirect(desired.OutputFormat) {
		c.Config.Logger.Infof("desired OutputFormat %s - but actually nil", dcl.SprintResource(desired.OutputFormat))
		return true
	}
	if !dcl.StringCanonicalize(desired.OutputFormat, actual.OutputFormat) && !dcl.IsZeroValue(desired.OutputFormat) {
		c.Config.Logger.Infof("Diff in OutputFormat. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OutputFormat), dcl.SprintResource(actual.OutputFormat))
		return true
	}
	if actual.ClientTags == nil && desired.ClientTags != nil && !dcl.IsEmptyValueIndirect(desired.ClientTags) {
		c.Config.Logger.Infof("desired ClientTags %s - but actually nil", dcl.SprintResource(desired.ClientTags))
		return true
	}
	if !dcl.StringSliceEquals(desired.ClientTags, actual.ClientTags) && !dcl.IsZeroValue(desired.ClientTags) {
		c.Config.Logger.Infof("Diff in ClientTags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientTags), dcl.SprintResource(actual.ClientTags))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if !dcl.MapEquals(desired.Properties, actual.Properties, []string(nil)) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.LoggingConfig == nil && desired.LoggingConfig != nil && !dcl.IsEmptyValueIndirect(desired.LoggingConfig) {
		c.Config.Logger.Infof("desired LoggingConfig %s - but actually nil", dcl.SprintResource(desired.LoggingConfig))
		return true
	}
	if compareJobPrestoJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareJobPrestoJobSlice(c *Client, desired, actual []JobPrestoJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPrestoJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPrestoJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPrestoJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPrestoJobMap(c *Client, desired, actual map[string]JobPrestoJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPrestoJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPrestoJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPrestoJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPrestoJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobPrestoJobQueryList(c *Client, desired, actual *JobPrestoJobQueryList) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Queries == nil && desired.Queries != nil && !dcl.IsEmptyValueIndirect(desired.Queries) {
		c.Config.Logger.Infof("desired Queries %s - but actually nil", dcl.SprintResource(desired.Queries))
		return true
	}
	if !dcl.StringSliceEquals(desired.Queries, actual.Queries) && !dcl.IsZeroValue(desired.Queries) {
		c.Config.Logger.Infof("Diff in Queries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Queries), dcl.SprintResource(actual.Queries))
		return true
	}
	return false
}

func compareJobPrestoJobQueryListSlice(c *Client, desired, actual []JobPrestoJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPrestoJobQueryList, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPrestoJobQueryList(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPrestoJobQueryList, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPrestoJobQueryListMap(c *Client, desired, actual map[string]JobPrestoJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPrestoJobQueryList, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPrestoJobQueryList, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPrestoJobQueryList(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPrestoJobQueryList, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobPrestoJobLoggingConfig(c *Client, desired, actual *JobPrestoJobLoggingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DriverLogLevels == nil && desired.DriverLogLevels != nil && !dcl.IsEmptyValueIndirect(desired.DriverLogLevels) {
		c.Config.Logger.Infof("desired DriverLogLevels %s - but actually nil", dcl.SprintResource(desired.DriverLogLevels))
		return true
	}
	if !dcl.MapEquals(desired.DriverLogLevels, actual.DriverLogLevels, []string(nil)) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}

func compareJobPrestoJobLoggingConfigSlice(c *Client, desired, actual []JobPrestoJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPrestoJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobPrestoJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobPrestoJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobPrestoJobLoggingConfigMap(c *Client, desired, actual map[string]JobPrestoJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobPrestoJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobPrestoJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobPrestoJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobPrestoJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
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
			c.Config.Logger.Infof("Diff in JobStatus, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobStatusHistory(c *Client, desired, actual *JobStatusHistory) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}

func compareJobStatusHistorySlice(c *Client, desired, actual []JobStatusHistory) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatusHistory, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobStatusHistory(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobStatusHistory, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStatusHistoryMap(c *Client, desired, actual map[string]JobStatusHistory) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatusHistory, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobStatusHistory, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobStatusHistory(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobStatusHistory, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobYarnApplications(c *Client, desired, actual *JobYarnApplications) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}

func compareJobYarnApplicationsSlice(c *Client, desired, actual []JobYarnApplications) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobYarnApplications, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobYarnApplications(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobYarnApplications, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobYarnApplicationsMap(c *Client, desired, actual map[string]JobYarnApplications) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobYarnApplications, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobYarnApplications, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobYarnApplications(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobYarnApplications, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobScheduling(c *Client, desired, actual *JobScheduling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MaxFailuresPerHour == nil && desired.MaxFailuresPerHour != nil && !dcl.IsEmptyValueIndirect(desired.MaxFailuresPerHour) {
		c.Config.Logger.Infof("desired MaxFailuresPerHour %s - but actually nil", dcl.SprintResource(desired.MaxFailuresPerHour))
		return true
	}
	if !reflect.DeepEqual(desired.MaxFailuresPerHour, actual.MaxFailuresPerHour) && !dcl.IsZeroValue(desired.MaxFailuresPerHour) {
		c.Config.Logger.Infof("Diff in MaxFailuresPerHour. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxFailuresPerHour), dcl.SprintResource(actual.MaxFailuresPerHour))
		return true
	}
	if actual.MaxFailuresTotal == nil && desired.MaxFailuresTotal != nil && !dcl.IsEmptyValueIndirect(desired.MaxFailuresTotal) {
		c.Config.Logger.Infof("desired MaxFailuresTotal %s - but actually nil", dcl.SprintResource(desired.MaxFailuresTotal))
		return true
	}
	if !reflect.DeepEqual(desired.MaxFailuresTotal, actual.MaxFailuresTotal) && !dcl.IsZeroValue(desired.MaxFailuresTotal) {
		c.Config.Logger.Infof("Diff in MaxFailuresTotal. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxFailuresTotal), dcl.SprintResource(actual.MaxFailuresTotal))
		return true
	}
	return false
}

func compareJobSchedulingSlice(c *Client, desired, actual []JobScheduling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobScheduling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobScheduling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobScheduling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobSchedulingMap(c *Client, desired, actual map[string]JobScheduling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobScheduling, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobScheduling, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobScheduling(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobScheduling, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobDriverRunner(c *Client, desired, actual *JobDriverRunner) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MasterDriverRunner == nil && desired.MasterDriverRunner != nil && !dcl.IsEmptyValueIndirect(desired.MasterDriverRunner) {
		c.Config.Logger.Infof("desired MasterDriverRunner %s - but actually nil", dcl.SprintResource(desired.MasterDriverRunner))
		return true
	}
	if compareJobDriverRunnerMasterDriverRunner(c, desired.MasterDriverRunner, actual.MasterDriverRunner) && !dcl.IsZeroValue(desired.MasterDriverRunner) {
		c.Config.Logger.Infof("Diff in MasterDriverRunner. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MasterDriverRunner), dcl.SprintResource(actual.MasterDriverRunner))
		return true
	}
	if actual.YarnDriverRunner == nil && desired.YarnDriverRunner != nil && !dcl.IsEmptyValueIndirect(desired.YarnDriverRunner) {
		c.Config.Logger.Infof("desired YarnDriverRunner %s - but actually nil", dcl.SprintResource(desired.YarnDriverRunner))
		return true
	}
	if compareJobDriverRunnerYarnDriverRunner(c, desired.YarnDriverRunner, actual.YarnDriverRunner) && !dcl.IsZeroValue(desired.YarnDriverRunner) {
		c.Config.Logger.Infof("Diff in YarnDriverRunner. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.YarnDriverRunner), dcl.SprintResource(actual.YarnDriverRunner))
		return true
	}
	return false
}

func compareJobDriverRunnerSlice(c *Client, desired, actual []JobDriverRunner) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobDriverRunner, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobDriverRunner(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobDriverRunner, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobDriverRunnerMap(c *Client, desired, actual map[string]JobDriverRunner) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobDriverRunner, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobDriverRunner, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobDriverRunner(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobDriverRunner, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobDriverRunnerMasterDriverRunner(c *Client, desired, actual *JobDriverRunnerMasterDriverRunner) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}

func compareJobDriverRunnerMasterDriverRunnerSlice(c *Client, desired, actual []JobDriverRunnerMasterDriverRunner) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobDriverRunnerMasterDriverRunner, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobDriverRunnerMasterDriverRunner(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobDriverRunnerMasterDriverRunner, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobDriverRunnerMasterDriverRunnerMap(c *Client, desired, actual map[string]JobDriverRunnerMasterDriverRunner) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobDriverRunnerMasterDriverRunner, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobDriverRunnerMasterDriverRunner, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobDriverRunnerMasterDriverRunner(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobDriverRunnerMasterDriverRunner, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobDriverRunnerYarnDriverRunner(c *Client, desired, actual *JobDriverRunnerYarnDriverRunner) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MemoryMb == nil && desired.MemoryMb != nil && !dcl.IsEmptyValueIndirect(desired.MemoryMb) {
		c.Config.Logger.Infof("desired MemoryMb %s - but actually nil", dcl.SprintResource(desired.MemoryMb))
		return true
	}
	if !reflect.DeepEqual(desired.MemoryMb, actual.MemoryMb) && !dcl.IsZeroValue(desired.MemoryMb) {
		c.Config.Logger.Infof("Diff in MemoryMb. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MemoryMb), dcl.SprintResource(actual.MemoryMb))
		return true
	}
	if actual.Vcores == nil && desired.Vcores != nil && !dcl.IsEmptyValueIndirect(desired.Vcores) {
		c.Config.Logger.Infof("desired Vcores %s - but actually nil", dcl.SprintResource(desired.Vcores))
		return true
	}
	if !reflect.DeepEqual(desired.Vcores, actual.Vcores) && !dcl.IsZeroValue(desired.Vcores) {
		c.Config.Logger.Infof("Diff in Vcores. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Vcores), dcl.SprintResource(actual.Vcores))
		return true
	}
	return false
}

func compareJobDriverRunnerYarnDriverRunnerSlice(c *Client, desired, actual []JobDriverRunnerYarnDriverRunner) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobDriverRunnerYarnDriverRunner, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobDriverRunnerYarnDriverRunner(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobDriverRunnerYarnDriverRunner, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobDriverRunnerYarnDriverRunnerMap(c *Client, desired, actual map[string]JobDriverRunnerYarnDriverRunner) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobDriverRunnerYarnDriverRunner, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in JobDriverRunnerYarnDriverRunner, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareJobDriverRunnerYarnDriverRunner(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in JobDriverRunnerYarnDriverRunner, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareJobStatusStateEnumSlice(c *Client, desired, actual []JobStatusStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatusStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobStatusStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobStatusStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStatusStateEnum(c *Client, desired, actual *JobStatusStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareJobStatusSubstateEnumSlice(c *Client, desired, actual []JobStatusSubstateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatusSubstateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobStatusSubstateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobStatusSubstateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStatusSubstateEnum(c *Client, desired, actual *JobStatusSubstateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareJobStatusHistoryStateEnumSlice(c *Client, desired, actual []JobStatusHistoryStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatusHistoryStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobStatusHistoryStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobStatusHistoryStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStatusHistoryStateEnum(c *Client, desired, actual *JobStatusHistoryStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareJobStatusHistorySubstateEnumSlice(c *Client, desired, actual []JobStatusHistorySubstateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobStatusHistorySubstateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobStatusHistorySubstateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobStatusHistorySubstateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobStatusHistorySubstateEnum(c *Client, desired, actual *JobStatusHistorySubstateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareJobYarnApplicationsStateEnumSlice(c *Client, desired, actual []JobYarnApplicationsStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in JobYarnApplicationsStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareJobYarnApplicationsStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in JobYarnApplicationsStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareJobYarnApplicationsStateEnum(c *Client, desired, actual *JobYarnApplicationsStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Job) urlNormalized() *Job {
	normalized := deepcopy.Copy(*r).(Job)
	normalized.SubmittedBy = dcl.SelfLinkToName(r.SubmittedBy)
	normalized.DriverInputResourceUri = dcl.SelfLinkToName(r.DriverInputResourceUri)
	normalized.DriverOutputResourceUri = dcl.SelfLinkToName(r.DriverOutputResourceUri)
	normalized.DriverControlFilesUri = dcl.SelfLinkToName(r.DriverControlFilesUri)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Job) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *Job) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region)
}

func (r *Job) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *Job) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "patch" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/jobs/{{name}}", "https://dataproc.googleapis.com/v1beta2/", userBasePath, fields), nil

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
	m = encodeJobCreateRequest(m)

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
	if v, err := expandJobReference(c, f.Reference); err != nil {
		return nil, fmt.Errorf("error expanding Reference into reference: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reference"] = v
	}
	if v, err := expandJobPlacement(c, f.Placement); err != nil {
		return nil, fmt.Errorf("error expanding Placement into placement: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["placement"] = v
	}
	if v, err := expandJobHadoopJob(c, f.HadoopJob); err != nil {
		return nil, fmt.Errorf("error expanding HadoopJob into hadoopJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hadoopJob"] = v
	}
	if v, err := expandJobSparkJob(c, f.SparkJob); err != nil {
		return nil, fmt.Errorf("error expanding SparkJob into sparkJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sparkJob"] = v
	}
	if v, err := expandJobPysparkJob(c, f.PysparkJob); err != nil {
		return nil, fmt.Errorf("error expanding PysparkJob into pysparkJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pysparkJob"] = v
	}
	if v, err := expandJobHiveJob(c, f.HiveJob); err != nil {
		return nil, fmt.Errorf("error expanding HiveJob into hiveJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hiveJob"] = v
	}
	if v, err := expandJobPigJob(c, f.PigJob); err != nil {
		return nil, fmt.Errorf("error expanding PigJob into pigJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pigJob"] = v
	}
	if v, err := expandJobSparkRJob(c, f.SparkRJob); err != nil {
		return nil, fmt.Errorf("error expanding SparkRJob into sparkRJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sparkRJob"] = v
	}
	if v, err := expandJobSparkSqlJob(c, f.SparkSqlJob); err != nil {
		return nil, fmt.Errorf("error expanding SparkSqlJob into sparkSqlJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sparkSqlJob"] = v
	}
	if v, err := expandJobPrestoJob(c, f.PrestoJob); err != nil {
		return nil, fmt.Errorf("error expanding PrestoJob into prestoJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["prestoJob"] = v
	}
	if v, err := expandJobStatus(c, f.Status); err != nil {
		return nil, fmt.Errorf("error expanding Status into status: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v, err := expandJobStatusHistorySlice(c, f.StatusHistory); err != nil {
		return nil, fmt.Errorf("error expanding StatusHistory into statusHistory: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["statusHistory"] = v
	}
	if v, err := expandJobYarnApplicationsSlice(c, f.YarnApplications); err != nil {
		return nil, fmt.Errorf("error expanding YarnApplications into yarnApplications: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["yarnApplications"] = v
	}
	if v := f.SubmittedBy; !dcl.IsEmptyValueIndirect(v) {
		m["submittedBy"] = v
	}
	if v := f.DriverInputResourceUri; !dcl.IsEmptyValueIndirect(v) {
		m["driverInputResourceUri"] = v
	}
	if v := f.DriverOutputResourceUri; !dcl.IsEmptyValueIndirect(v) {
		m["driverOutputResourceUri"] = v
	}
	if v := f.DriverControlFilesUri; !dcl.IsEmptyValueIndirect(v) {
		m["driverControlFilesUri"] = v
	}
	if v := f.Interactive; !dcl.IsEmptyValueIndirect(v) {
		m["interactive"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandJobScheduling(c, f.Scheduling); err != nil {
		return nil, fmt.Errorf("error expanding Scheduling into scheduling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scheduling"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["jobUuid"] = v
	}
	if v := f.Done; !dcl.IsEmptyValueIndirect(v) {
		m["done"] = v
	}
	if v, err := expandJobDriverRunner(c, f.DriverRunner); err != nil {
		return nil, fmt.Errorf("error expanding DriverRunner into driverRunner: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["driverRunner"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Region into region: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
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
	r.Reference = flattenJobReference(c, m["reference"])
	r.Placement = flattenJobPlacement(c, m["placement"])
	r.HadoopJob = flattenJobHadoopJob(c, m["hadoopJob"])
	r.SparkJob = flattenJobSparkJob(c, m["sparkJob"])
	r.PysparkJob = flattenJobPysparkJob(c, m["pysparkJob"])
	r.HiveJob = flattenJobHiveJob(c, m["hiveJob"])
	r.PigJob = flattenJobPigJob(c, m["pigJob"])
	r.SparkRJob = flattenJobSparkRJob(c, m["sparkRJob"])
	r.SparkSqlJob = flattenJobSparkSqlJob(c, m["sparkSqlJob"])
	r.PrestoJob = flattenJobPrestoJob(c, m["prestoJob"])
	r.Status = flattenJobStatus(c, m["status"])
	r.StatusHistory = flattenJobStatusHistorySlice(c, m["statusHistory"])
	r.YarnApplications = flattenJobYarnApplicationsSlice(c, m["yarnApplications"])
	r.SubmittedBy = dcl.FlattenString(m["submittedBy"])
	r.DriverInputResourceUri = dcl.FlattenString(m["driverInputResourceUri"])
	r.DriverOutputResourceUri = dcl.FlattenString(m["driverOutputResourceUri"])
	r.DriverControlFilesUri = dcl.FlattenString(m["driverControlFilesUri"])
	r.Interactive = dcl.FlattenBool(m["interactive"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Scheduling = flattenJobScheduling(c, m["scheduling"])
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["jobUuid"]))
	r.Done = dcl.FlattenBool(m["done"])
	r.DriverRunner = flattenJobDriverRunner(c, m["driverRunner"])
	r.Region = dcl.FlattenString(m["region"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandJobReferenceMap expands the contents of JobReference into a JSON
// request object.
func expandJobReferenceMap(c *Client, f map[string]JobReference) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobReference(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobReferenceSlice expands the contents of JobReference into a JSON
// request object.
func expandJobReferenceSlice(c *Client, f []JobReference) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobReference(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobReferenceMap flattens the contents of JobReference from a JSON
// response object.
func flattenJobReferenceMap(c *Client, i interface{}) map[string]JobReference {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobReference{}
	}

	if len(a) == 0 {
		return map[string]JobReference{}
	}

	items := make(map[string]JobReference)
	for k, item := range a {
		items[k] = *flattenJobReference(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobReferenceSlice flattens the contents of JobReference from a JSON
// response object.
func flattenJobReferenceSlice(c *Client, i interface{}) []JobReference {
	a, ok := i.([]interface{})
	if !ok {
		return []JobReference{}
	}

	if len(a) == 0 {
		return []JobReference{}
	}

	items := make([]JobReference, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobReference(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobReference expands an instance of JobReference into a JSON
// request object.
func expandJobReference(c *Client, f *JobReference) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.JobId; !dcl.IsEmptyValueIndirect(v) {
		m["jobId"] = v
	}

	return m, nil
}

// flattenJobReference flattens an instance of JobReference from a JSON
// response object.
func flattenJobReference(c *Client, i interface{}) *JobReference {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobReference{}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.JobId = dcl.FlattenString(m["jobId"])

	return r
}

// expandJobPlacementMap expands the contents of JobPlacement into a JSON
// request object.
func expandJobPlacementMap(c *Client, f map[string]JobPlacement) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPlacement(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPlacementSlice expands the contents of JobPlacement into a JSON
// request object.
func expandJobPlacementSlice(c *Client, f []JobPlacement) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPlacement(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPlacementMap flattens the contents of JobPlacement from a JSON
// response object.
func flattenJobPlacementMap(c *Client, i interface{}) map[string]JobPlacement {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPlacement{}
	}

	if len(a) == 0 {
		return map[string]JobPlacement{}
	}

	items := make(map[string]JobPlacement)
	for k, item := range a {
		items[k] = *flattenJobPlacement(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPlacementSlice flattens the contents of JobPlacement from a JSON
// response object.
func flattenJobPlacementSlice(c *Client, i interface{}) []JobPlacement {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPlacement{}
	}

	if len(a) == 0 {
		return []JobPlacement{}
	}

	items := make([]JobPlacement, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPlacement(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPlacement expands an instance of JobPlacement into a JSON
// request object.
func expandJobPlacement(c *Client, f *JobPlacement) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ClusterName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterName"] = v
	}
	if v := f.ClusterUuid; !dcl.IsEmptyValueIndirect(v) {
		m["clusterUuid"] = v
	}
	if v := f.ClusterLabels; !dcl.IsEmptyValueIndirect(v) {
		m["clusterLabels"] = v
	}

	return m, nil
}

// flattenJobPlacement flattens an instance of JobPlacement from a JSON
// response object.
func flattenJobPlacement(c *Client, i interface{}) *JobPlacement {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPlacement{}
	r.ClusterName = dcl.FlattenString(m["clusterName"])
	r.ClusterUuid = dcl.FlattenString(m["clusterUuid"])
	r.ClusterLabels = dcl.FlattenKeyValuePairs(m["clusterLabels"])

	return r
}

// expandJobHadoopJobMap expands the contents of JobHadoopJob into a JSON
// request object.
func expandJobHadoopJobMap(c *Client, f map[string]JobHadoopJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHadoopJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHadoopJobSlice expands the contents of JobHadoopJob into a JSON
// request object.
func expandJobHadoopJobSlice(c *Client, f []JobHadoopJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHadoopJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHadoopJobMap flattens the contents of JobHadoopJob from a JSON
// response object.
func flattenJobHadoopJobMap(c *Client, i interface{}) map[string]JobHadoopJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHadoopJob{}
	}

	if len(a) == 0 {
		return map[string]JobHadoopJob{}
	}

	items := make(map[string]JobHadoopJob)
	for k, item := range a {
		items[k] = *flattenJobHadoopJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobHadoopJobSlice flattens the contents of JobHadoopJob from a JSON
// response object.
func flattenJobHadoopJobSlice(c *Client, i interface{}) []JobHadoopJob {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHadoopJob{}
	}

	if len(a) == 0 {
		return []JobHadoopJob{}
	}

	items := make([]JobHadoopJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHadoopJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobHadoopJob expands an instance of JobHadoopJob into a JSON
// request object.
func expandJobHadoopJob(c *Client, f *JobHadoopJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MainJarFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["mainJarFileUri"] = v
	}
	if v := f.MainClass; !dcl.IsEmptyValueIndirect(v) {
		m["mainClass"] = v
	}
	if v := f.Args; !dcl.IsEmptyValueIndirect(v) {
		m["args"] = v
	}
	if v := f.JarFileUris; !dcl.IsEmptyValueIndirect(v) {
		m["jarFileUris"] = v
	}
	if v := f.FileUris; !dcl.IsEmptyValueIndirect(v) {
		m["fileUris"] = v
	}
	if v := f.ArchiveUris; !dcl.IsEmptyValueIndirect(v) {
		m["archiveUris"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v, err := expandJobHadoopJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenJobHadoopJob flattens an instance of JobHadoopJob from a JSON
// response object.
func flattenJobHadoopJob(c *Client, i interface{}) *JobHadoopJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHadoopJob{}
	r.MainJarFileUri = dcl.FlattenString(m["mainJarFileUri"])
	r.MainClass = dcl.FlattenString(m["mainClass"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.FileUris = dcl.FlattenStringSlice(m["fileUris"])
	r.ArchiveUris = dcl.FlattenStringSlice(m["archiveUris"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenJobHadoopJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandJobHadoopJobLoggingConfigMap expands the contents of JobHadoopJobLoggingConfig into a JSON
// request object.
func expandJobHadoopJobLoggingConfigMap(c *Client, f map[string]JobHadoopJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHadoopJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHadoopJobLoggingConfigSlice expands the contents of JobHadoopJobLoggingConfig into a JSON
// request object.
func expandJobHadoopJobLoggingConfigSlice(c *Client, f []JobHadoopJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHadoopJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHadoopJobLoggingConfigMap flattens the contents of JobHadoopJobLoggingConfig from a JSON
// response object.
func flattenJobHadoopJobLoggingConfigMap(c *Client, i interface{}) map[string]JobHadoopJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHadoopJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]JobHadoopJobLoggingConfig{}
	}

	items := make(map[string]JobHadoopJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenJobHadoopJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobHadoopJobLoggingConfigSlice flattens the contents of JobHadoopJobLoggingConfig from a JSON
// response object.
func flattenJobHadoopJobLoggingConfigSlice(c *Client, i interface{}) []JobHadoopJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHadoopJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []JobHadoopJobLoggingConfig{}
	}

	items := make([]JobHadoopJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHadoopJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobHadoopJobLoggingConfig expands an instance of JobHadoopJobLoggingConfig into a JSON
// request object.
func expandJobHadoopJobLoggingConfig(c *Client, f *JobHadoopJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenJobHadoopJobLoggingConfig flattens an instance of JobHadoopJobLoggingConfig from a JSON
// response object.
func flattenJobHadoopJobLoggingConfig(c *Client, i interface{}) *JobHadoopJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHadoopJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandJobSparkJobMap expands the contents of JobSparkJob into a JSON
// request object.
func expandJobSparkJobMap(c *Client, f map[string]JobSparkJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobSparkJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobSparkJobSlice expands the contents of JobSparkJob into a JSON
// request object.
func expandJobSparkJobSlice(c *Client, f []JobSparkJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobSparkJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobSparkJobMap flattens the contents of JobSparkJob from a JSON
// response object.
func flattenJobSparkJobMap(c *Client, i interface{}) map[string]JobSparkJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobSparkJob{}
	}

	if len(a) == 0 {
		return map[string]JobSparkJob{}
	}

	items := make(map[string]JobSparkJob)
	for k, item := range a {
		items[k] = *flattenJobSparkJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobSparkJobSlice flattens the contents of JobSparkJob from a JSON
// response object.
func flattenJobSparkJobSlice(c *Client, i interface{}) []JobSparkJob {
	a, ok := i.([]interface{})
	if !ok {
		return []JobSparkJob{}
	}

	if len(a) == 0 {
		return []JobSparkJob{}
	}

	items := make([]JobSparkJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobSparkJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobSparkJob expands an instance of JobSparkJob into a JSON
// request object.
func expandJobSparkJob(c *Client, f *JobSparkJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MainJarFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["mainJarFileUri"] = v
	}
	if v := f.MainClass; !dcl.IsEmptyValueIndirect(v) {
		m["mainClass"] = v
	}
	if v := f.Args; !dcl.IsEmptyValueIndirect(v) {
		m["args"] = v
	}
	if v := f.JarFileUris; !dcl.IsEmptyValueIndirect(v) {
		m["jarFileUris"] = v
	}
	if v := f.FileUris; !dcl.IsEmptyValueIndirect(v) {
		m["fileUris"] = v
	}
	if v := f.ArchiveUris; !dcl.IsEmptyValueIndirect(v) {
		m["archiveUris"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v, err := expandJobSparkJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenJobSparkJob flattens an instance of JobSparkJob from a JSON
// response object.
func flattenJobSparkJob(c *Client, i interface{}) *JobSparkJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobSparkJob{}
	r.MainJarFileUri = dcl.FlattenString(m["mainJarFileUri"])
	r.MainClass = dcl.FlattenString(m["mainClass"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.FileUris = dcl.FlattenStringSlice(m["fileUris"])
	r.ArchiveUris = dcl.FlattenStringSlice(m["archiveUris"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenJobSparkJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandJobSparkJobLoggingConfigMap expands the contents of JobSparkJobLoggingConfig into a JSON
// request object.
func expandJobSparkJobLoggingConfigMap(c *Client, f map[string]JobSparkJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobSparkJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobSparkJobLoggingConfigSlice expands the contents of JobSparkJobLoggingConfig into a JSON
// request object.
func expandJobSparkJobLoggingConfigSlice(c *Client, f []JobSparkJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobSparkJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobSparkJobLoggingConfigMap flattens the contents of JobSparkJobLoggingConfig from a JSON
// response object.
func flattenJobSparkJobLoggingConfigMap(c *Client, i interface{}) map[string]JobSparkJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobSparkJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]JobSparkJobLoggingConfig{}
	}

	items := make(map[string]JobSparkJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenJobSparkJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobSparkJobLoggingConfigSlice flattens the contents of JobSparkJobLoggingConfig from a JSON
// response object.
func flattenJobSparkJobLoggingConfigSlice(c *Client, i interface{}) []JobSparkJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobSparkJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []JobSparkJobLoggingConfig{}
	}

	items := make([]JobSparkJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobSparkJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobSparkJobLoggingConfig expands an instance of JobSparkJobLoggingConfig into a JSON
// request object.
func expandJobSparkJobLoggingConfig(c *Client, f *JobSparkJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenJobSparkJobLoggingConfig flattens an instance of JobSparkJobLoggingConfig from a JSON
// response object.
func flattenJobSparkJobLoggingConfig(c *Client, i interface{}) *JobSparkJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobSparkJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandJobPysparkJobMap expands the contents of JobPysparkJob into a JSON
// request object.
func expandJobPysparkJobMap(c *Client, f map[string]JobPysparkJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPysparkJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPysparkJobSlice expands the contents of JobPysparkJob into a JSON
// request object.
func expandJobPysparkJobSlice(c *Client, f []JobPysparkJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPysparkJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPysparkJobMap flattens the contents of JobPysparkJob from a JSON
// response object.
func flattenJobPysparkJobMap(c *Client, i interface{}) map[string]JobPysparkJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPysparkJob{}
	}

	if len(a) == 0 {
		return map[string]JobPysparkJob{}
	}

	items := make(map[string]JobPysparkJob)
	for k, item := range a {
		items[k] = *flattenJobPysparkJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPysparkJobSlice flattens the contents of JobPysparkJob from a JSON
// response object.
func flattenJobPysparkJobSlice(c *Client, i interface{}) []JobPysparkJob {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPysparkJob{}
	}

	if len(a) == 0 {
		return []JobPysparkJob{}
	}

	items := make([]JobPysparkJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPysparkJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPysparkJob expands an instance of JobPysparkJob into a JSON
// request object.
func expandJobPysparkJob(c *Client, f *JobPysparkJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MainPythonFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["mainPythonFileUri"] = v
	}
	if v := f.Args; !dcl.IsEmptyValueIndirect(v) {
		m["args"] = v
	}
	if v := f.PythonFileUris; !dcl.IsEmptyValueIndirect(v) {
		m["pythonFileUris"] = v
	}
	if v := f.JarFileUris; !dcl.IsEmptyValueIndirect(v) {
		m["jarFileUris"] = v
	}
	if v := f.FileUris; !dcl.IsEmptyValueIndirect(v) {
		m["fileUris"] = v
	}
	if v := f.ArchiveUris; !dcl.IsEmptyValueIndirect(v) {
		m["archiveUris"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v, err := expandJobPysparkJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenJobPysparkJob flattens an instance of JobPysparkJob from a JSON
// response object.
func flattenJobPysparkJob(c *Client, i interface{}) *JobPysparkJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPysparkJob{}
	r.MainPythonFileUri = dcl.FlattenString(m["mainPythonFileUri"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.PythonFileUris = dcl.FlattenStringSlice(m["pythonFileUris"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.FileUris = dcl.FlattenStringSlice(m["fileUris"])
	r.ArchiveUris = dcl.FlattenStringSlice(m["archiveUris"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenJobPysparkJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandJobPysparkJobLoggingConfigMap expands the contents of JobPysparkJobLoggingConfig into a JSON
// request object.
func expandJobPysparkJobLoggingConfigMap(c *Client, f map[string]JobPysparkJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPysparkJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPysparkJobLoggingConfigSlice expands the contents of JobPysparkJobLoggingConfig into a JSON
// request object.
func expandJobPysparkJobLoggingConfigSlice(c *Client, f []JobPysparkJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPysparkJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPysparkJobLoggingConfigMap flattens the contents of JobPysparkJobLoggingConfig from a JSON
// response object.
func flattenJobPysparkJobLoggingConfigMap(c *Client, i interface{}) map[string]JobPysparkJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPysparkJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]JobPysparkJobLoggingConfig{}
	}

	items := make(map[string]JobPysparkJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenJobPysparkJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPysparkJobLoggingConfigSlice flattens the contents of JobPysparkJobLoggingConfig from a JSON
// response object.
func flattenJobPysparkJobLoggingConfigSlice(c *Client, i interface{}) []JobPysparkJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPysparkJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []JobPysparkJobLoggingConfig{}
	}

	items := make([]JobPysparkJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPysparkJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPysparkJobLoggingConfig expands an instance of JobPysparkJobLoggingConfig into a JSON
// request object.
func expandJobPysparkJobLoggingConfig(c *Client, f *JobPysparkJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenJobPysparkJobLoggingConfig flattens an instance of JobPysparkJobLoggingConfig from a JSON
// response object.
func flattenJobPysparkJobLoggingConfig(c *Client, i interface{}) *JobPysparkJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPysparkJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandJobHiveJobMap expands the contents of JobHiveJob into a JSON
// request object.
func expandJobHiveJobMap(c *Client, f map[string]JobHiveJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHiveJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHiveJobSlice expands the contents of JobHiveJob into a JSON
// request object.
func expandJobHiveJobSlice(c *Client, f []JobHiveJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHiveJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHiveJobMap flattens the contents of JobHiveJob from a JSON
// response object.
func flattenJobHiveJobMap(c *Client, i interface{}) map[string]JobHiveJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHiveJob{}
	}

	if len(a) == 0 {
		return map[string]JobHiveJob{}
	}

	items := make(map[string]JobHiveJob)
	for k, item := range a {
		items[k] = *flattenJobHiveJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobHiveJobSlice flattens the contents of JobHiveJob from a JSON
// response object.
func flattenJobHiveJobSlice(c *Client, i interface{}) []JobHiveJob {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHiveJob{}
	}

	if len(a) == 0 {
		return []JobHiveJob{}
	}

	items := make([]JobHiveJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHiveJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobHiveJob expands an instance of JobHiveJob into a JSON
// request object.
func expandJobHiveJob(c *Client, f *JobHiveJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["queryFileUri"] = v
	}
	if v, err := expandJobHiveJobQueryList(c, f.QueryList); err != nil {
		return nil, fmt.Errorf("error expanding QueryList into queryList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["queryList"] = v
	}
	if v := f.ContinueOnFailure; !dcl.IsEmptyValueIndirect(v) {
		m["continueOnFailure"] = v
	}
	if v := f.ScriptVariables; !dcl.IsEmptyValueIndirect(v) {
		m["scriptVariables"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v := f.JarFileUris; !dcl.IsEmptyValueIndirect(v) {
		m["jarFileUris"] = v
	}

	return m, nil
}

// flattenJobHiveJob flattens an instance of JobHiveJob from a JSON
// response object.
func flattenJobHiveJob(c *Client, i interface{}) *JobHiveJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHiveJob{}
	r.QueryFileUri = dcl.FlattenString(m["queryFileUri"])
	r.QueryList = flattenJobHiveJobQueryList(c, m["queryList"])
	r.ContinueOnFailure = dcl.FlattenBool(m["continueOnFailure"])
	r.ScriptVariables = dcl.FlattenKeyValuePairs(m["scriptVariables"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])

	return r
}

// expandJobHiveJobQueryListMap expands the contents of JobHiveJobQueryList into a JSON
// request object.
func expandJobHiveJobQueryListMap(c *Client, f map[string]JobHiveJobQueryList) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHiveJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHiveJobQueryListSlice expands the contents of JobHiveJobQueryList into a JSON
// request object.
func expandJobHiveJobQueryListSlice(c *Client, f []JobHiveJobQueryList) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHiveJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHiveJobQueryListMap flattens the contents of JobHiveJobQueryList from a JSON
// response object.
func flattenJobHiveJobQueryListMap(c *Client, i interface{}) map[string]JobHiveJobQueryList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHiveJobQueryList{}
	}

	if len(a) == 0 {
		return map[string]JobHiveJobQueryList{}
	}

	items := make(map[string]JobHiveJobQueryList)
	for k, item := range a {
		items[k] = *flattenJobHiveJobQueryList(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobHiveJobQueryListSlice flattens the contents of JobHiveJobQueryList from a JSON
// response object.
func flattenJobHiveJobQueryListSlice(c *Client, i interface{}) []JobHiveJobQueryList {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHiveJobQueryList{}
	}

	if len(a) == 0 {
		return []JobHiveJobQueryList{}
	}

	items := make([]JobHiveJobQueryList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHiveJobQueryList(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobHiveJobQueryList expands an instance of JobHiveJobQueryList into a JSON
// request object.
func expandJobHiveJobQueryList(c *Client, f *JobHiveJobQueryList) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Queries; !dcl.IsEmptyValueIndirect(v) {
		m["queries"] = v
	}

	return m, nil
}

// flattenJobHiveJobQueryList flattens an instance of JobHiveJobQueryList from a JSON
// response object.
func flattenJobHiveJobQueryList(c *Client, i interface{}) *JobHiveJobQueryList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHiveJobQueryList{}
	r.Queries = dcl.FlattenStringSlice(m["queries"])

	return r
}

// expandJobPigJobMap expands the contents of JobPigJob into a JSON
// request object.
func expandJobPigJobMap(c *Client, f map[string]JobPigJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPigJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPigJobSlice expands the contents of JobPigJob into a JSON
// request object.
func expandJobPigJobSlice(c *Client, f []JobPigJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPigJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPigJobMap flattens the contents of JobPigJob from a JSON
// response object.
func flattenJobPigJobMap(c *Client, i interface{}) map[string]JobPigJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPigJob{}
	}

	if len(a) == 0 {
		return map[string]JobPigJob{}
	}

	items := make(map[string]JobPigJob)
	for k, item := range a {
		items[k] = *flattenJobPigJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPigJobSlice flattens the contents of JobPigJob from a JSON
// response object.
func flattenJobPigJobSlice(c *Client, i interface{}) []JobPigJob {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPigJob{}
	}

	if len(a) == 0 {
		return []JobPigJob{}
	}

	items := make([]JobPigJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPigJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPigJob expands an instance of JobPigJob into a JSON
// request object.
func expandJobPigJob(c *Client, f *JobPigJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["queryFileUri"] = v
	}
	if v, err := expandJobPigJobQueryList(c, f.QueryList); err != nil {
		return nil, fmt.Errorf("error expanding QueryList into queryList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["queryList"] = v
	}
	if v := f.ContinueOnFailure; !dcl.IsEmptyValueIndirect(v) {
		m["continueOnFailure"] = v
	}
	if v := f.ScriptVariables; !dcl.IsEmptyValueIndirect(v) {
		m["scriptVariables"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v := f.JarFileUris; !dcl.IsEmptyValueIndirect(v) {
		m["jarFileUris"] = v
	}
	if v, err := expandJobPigJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenJobPigJob flattens an instance of JobPigJob from a JSON
// response object.
func flattenJobPigJob(c *Client, i interface{}) *JobPigJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPigJob{}
	r.QueryFileUri = dcl.FlattenString(m["queryFileUri"])
	r.QueryList = flattenJobPigJobQueryList(c, m["queryList"])
	r.ContinueOnFailure = dcl.FlattenBool(m["continueOnFailure"])
	r.ScriptVariables = dcl.FlattenKeyValuePairs(m["scriptVariables"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.LoggingConfig = flattenJobPigJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandJobPigJobQueryListMap expands the contents of JobPigJobQueryList into a JSON
// request object.
func expandJobPigJobQueryListMap(c *Client, f map[string]JobPigJobQueryList) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPigJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPigJobQueryListSlice expands the contents of JobPigJobQueryList into a JSON
// request object.
func expandJobPigJobQueryListSlice(c *Client, f []JobPigJobQueryList) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPigJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPigJobQueryListMap flattens the contents of JobPigJobQueryList from a JSON
// response object.
func flattenJobPigJobQueryListMap(c *Client, i interface{}) map[string]JobPigJobQueryList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPigJobQueryList{}
	}

	if len(a) == 0 {
		return map[string]JobPigJobQueryList{}
	}

	items := make(map[string]JobPigJobQueryList)
	for k, item := range a {
		items[k] = *flattenJobPigJobQueryList(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPigJobQueryListSlice flattens the contents of JobPigJobQueryList from a JSON
// response object.
func flattenJobPigJobQueryListSlice(c *Client, i interface{}) []JobPigJobQueryList {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPigJobQueryList{}
	}

	if len(a) == 0 {
		return []JobPigJobQueryList{}
	}

	items := make([]JobPigJobQueryList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPigJobQueryList(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPigJobQueryList expands an instance of JobPigJobQueryList into a JSON
// request object.
func expandJobPigJobQueryList(c *Client, f *JobPigJobQueryList) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Queries; !dcl.IsEmptyValueIndirect(v) {
		m["queries"] = v
	}

	return m, nil
}

// flattenJobPigJobQueryList flattens an instance of JobPigJobQueryList from a JSON
// response object.
func flattenJobPigJobQueryList(c *Client, i interface{}) *JobPigJobQueryList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPigJobQueryList{}
	r.Queries = dcl.FlattenStringSlice(m["queries"])

	return r
}

// expandJobPigJobLoggingConfigMap expands the contents of JobPigJobLoggingConfig into a JSON
// request object.
func expandJobPigJobLoggingConfigMap(c *Client, f map[string]JobPigJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPigJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPigJobLoggingConfigSlice expands the contents of JobPigJobLoggingConfig into a JSON
// request object.
func expandJobPigJobLoggingConfigSlice(c *Client, f []JobPigJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPigJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPigJobLoggingConfigMap flattens the contents of JobPigJobLoggingConfig from a JSON
// response object.
func flattenJobPigJobLoggingConfigMap(c *Client, i interface{}) map[string]JobPigJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPigJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]JobPigJobLoggingConfig{}
	}

	items := make(map[string]JobPigJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenJobPigJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPigJobLoggingConfigSlice flattens the contents of JobPigJobLoggingConfig from a JSON
// response object.
func flattenJobPigJobLoggingConfigSlice(c *Client, i interface{}) []JobPigJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPigJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []JobPigJobLoggingConfig{}
	}

	items := make([]JobPigJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPigJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPigJobLoggingConfig expands an instance of JobPigJobLoggingConfig into a JSON
// request object.
func expandJobPigJobLoggingConfig(c *Client, f *JobPigJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenJobPigJobLoggingConfig flattens an instance of JobPigJobLoggingConfig from a JSON
// response object.
func flattenJobPigJobLoggingConfig(c *Client, i interface{}) *JobPigJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPigJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandJobSparkRJobMap expands the contents of JobSparkRJob into a JSON
// request object.
func expandJobSparkRJobMap(c *Client, f map[string]JobSparkRJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobSparkRJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobSparkRJobSlice expands the contents of JobSparkRJob into a JSON
// request object.
func expandJobSparkRJobSlice(c *Client, f []JobSparkRJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobSparkRJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobSparkRJobMap flattens the contents of JobSparkRJob from a JSON
// response object.
func flattenJobSparkRJobMap(c *Client, i interface{}) map[string]JobSparkRJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobSparkRJob{}
	}

	if len(a) == 0 {
		return map[string]JobSparkRJob{}
	}

	items := make(map[string]JobSparkRJob)
	for k, item := range a {
		items[k] = *flattenJobSparkRJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobSparkRJobSlice flattens the contents of JobSparkRJob from a JSON
// response object.
func flattenJobSparkRJobSlice(c *Client, i interface{}) []JobSparkRJob {
	a, ok := i.([]interface{})
	if !ok {
		return []JobSparkRJob{}
	}

	if len(a) == 0 {
		return []JobSparkRJob{}
	}

	items := make([]JobSparkRJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobSparkRJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobSparkRJob expands an instance of JobSparkRJob into a JSON
// request object.
func expandJobSparkRJob(c *Client, f *JobSparkRJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MainRFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["mainRFileUri"] = v
	}
	if v := f.Args; !dcl.IsEmptyValueIndirect(v) {
		m["args"] = v
	}
	if v := f.FileUris; !dcl.IsEmptyValueIndirect(v) {
		m["fileUris"] = v
	}
	if v := f.ArchiveUris; !dcl.IsEmptyValueIndirect(v) {
		m["archiveUris"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v, err := expandJobSparkRJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenJobSparkRJob flattens an instance of JobSparkRJob from a JSON
// response object.
func flattenJobSparkRJob(c *Client, i interface{}) *JobSparkRJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobSparkRJob{}
	r.MainRFileUri = dcl.FlattenString(m["mainRFileUri"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.FileUris = dcl.FlattenStringSlice(m["fileUris"])
	r.ArchiveUris = dcl.FlattenStringSlice(m["archiveUris"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenJobSparkRJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandJobSparkRJobLoggingConfigMap expands the contents of JobSparkRJobLoggingConfig into a JSON
// request object.
func expandJobSparkRJobLoggingConfigMap(c *Client, f map[string]JobSparkRJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobSparkRJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobSparkRJobLoggingConfigSlice expands the contents of JobSparkRJobLoggingConfig into a JSON
// request object.
func expandJobSparkRJobLoggingConfigSlice(c *Client, f []JobSparkRJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobSparkRJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobSparkRJobLoggingConfigMap flattens the contents of JobSparkRJobLoggingConfig from a JSON
// response object.
func flattenJobSparkRJobLoggingConfigMap(c *Client, i interface{}) map[string]JobSparkRJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobSparkRJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]JobSparkRJobLoggingConfig{}
	}

	items := make(map[string]JobSparkRJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenJobSparkRJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobSparkRJobLoggingConfigSlice flattens the contents of JobSparkRJobLoggingConfig from a JSON
// response object.
func flattenJobSparkRJobLoggingConfigSlice(c *Client, i interface{}) []JobSparkRJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobSparkRJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []JobSparkRJobLoggingConfig{}
	}

	items := make([]JobSparkRJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobSparkRJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobSparkRJobLoggingConfig expands an instance of JobSparkRJobLoggingConfig into a JSON
// request object.
func expandJobSparkRJobLoggingConfig(c *Client, f *JobSparkRJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenJobSparkRJobLoggingConfig flattens an instance of JobSparkRJobLoggingConfig from a JSON
// response object.
func flattenJobSparkRJobLoggingConfig(c *Client, i interface{}) *JobSparkRJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobSparkRJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandJobSparkSqlJobMap expands the contents of JobSparkSqlJob into a JSON
// request object.
func expandJobSparkSqlJobMap(c *Client, f map[string]JobSparkSqlJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobSparkSqlJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobSparkSqlJobSlice expands the contents of JobSparkSqlJob into a JSON
// request object.
func expandJobSparkSqlJobSlice(c *Client, f []JobSparkSqlJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobSparkSqlJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobSparkSqlJobMap flattens the contents of JobSparkSqlJob from a JSON
// response object.
func flattenJobSparkSqlJobMap(c *Client, i interface{}) map[string]JobSparkSqlJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobSparkSqlJob{}
	}

	if len(a) == 0 {
		return map[string]JobSparkSqlJob{}
	}

	items := make(map[string]JobSparkSqlJob)
	for k, item := range a {
		items[k] = *flattenJobSparkSqlJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobSparkSqlJobSlice flattens the contents of JobSparkSqlJob from a JSON
// response object.
func flattenJobSparkSqlJobSlice(c *Client, i interface{}) []JobSparkSqlJob {
	a, ok := i.([]interface{})
	if !ok {
		return []JobSparkSqlJob{}
	}

	if len(a) == 0 {
		return []JobSparkSqlJob{}
	}

	items := make([]JobSparkSqlJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobSparkSqlJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobSparkSqlJob expands an instance of JobSparkSqlJob into a JSON
// request object.
func expandJobSparkSqlJob(c *Client, f *JobSparkSqlJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["queryFileUri"] = v
	}
	if v, err := expandJobSparkSqlJobQueryList(c, f.QueryList); err != nil {
		return nil, fmt.Errorf("error expanding QueryList into queryList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["queryList"] = v
	}
	if v := f.ScriptVariables; !dcl.IsEmptyValueIndirect(v) {
		m["scriptVariables"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v := f.JarFileUris; !dcl.IsEmptyValueIndirect(v) {
		m["jarFileUris"] = v
	}
	if v, err := expandJobSparkSqlJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenJobSparkSqlJob flattens an instance of JobSparkSqlJob from a JSON
// response object.
func flattenJobSparkSqlJob(c *Client, i interface{}) *JobSparkSqlJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobSparkSqlJob{}
	r.QueryFileUri = dcl.FlattenString(m["queryFileUri"])
	r.QueryList = flattenJobSparkSqlJobQueryList(c, m["queryList"])
	r.ScriptVariables = dcl.FlattenKeyValuePairs(m["scriptVariables"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.LoggingConfig = flattenJobSparkSqlJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandJobSparkSqlJobQueryListMap expands the contents of JobSparkSqlJobQueryList into a JSON
// request object.
func expandJobSparkSqlJobQueryListMap(c *Client, f map[string]JobSparkSqlJobQueryList) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobSparkSqlJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobSparkSqlJobQueryListSlice expands the contents of JobSparkSqlJobQueryList into a JSON
// request object.
func expandJobSparkSqlJobQueryListSlice(c *Client, f []JobSparkSqlJobQueryList) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobSparkSqlJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobSparkSqlJobQueryListMap flattens the contents of JobSparkSqlJobQueryList from a JSON
// response object.
func flattenJobSparkSqlJobQueryListMap(c *Client, i interface{}) map[string]JobSparkSqlJobQueryList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobSparkSqlJobQueryList{}
	}

	if len(a) == 0 {
		return map[string]JobSparkSqlJobQueryList{}
	}

	items := make(map[string]JobSparkSqlJobQueryList)
	for k, item := range a {
		items[k] = *flattenJobSparkSqlJobQueryList(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobSparkSqlJobQueryListSlice flattens the contents of JobSparkSqlJobQueryList from a JSON
// response object.
func flattenJobSparkSqlJobQueryListSlice(c *Client, i interface{}) []JobSparkSqlJobQueryList {
	a, ok := i.([]interface{})
	if !ok {
		return []JobSparkSqlJobQueryList{}
	}

	if len(a) == 0 {
		return []JobSparkSqlJobQueryList{}
	}

	items := make([]JobSparkSqlJobQueryList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobSparkSqlJobQueryList(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobSparkSqlJobQueryList expands an instance of JobSparkSqlJobQueryList into a JSON
// request object.
func expandJobSparkSqlJobQueryList(c *Client, f *JobSparkSqlJobQueryList) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Queries; !dcl.IsEmptyValueIndirect(v) {
		m["queries"] = v
	}

	return m, nil
}

// flattenJobSparkSqlJobQueryList flattens an instance of JobSparkSqlJobQueryList from a JSON
// response object.
func flattenJobSparkSqlJobQueryList(c *Client, i interface{}) *JobSparkSqlJobQueryList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobSparkSqlJobQueryList{}
	r.Queries = dcl.FlattenStringSlice(m["queries"])

	return r
}

// expandJobSparkSqlJobLoggingConfigMap expands the contents of JobSparkSqlJobLoggingConfig into a JSON
// request object.
func expandJobSparkSqlJobLoggingConfigMap(c *Client, f map[string]JobSparkSqlJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobSparkSqlJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobSparkSqlJobLoggingConfigSlice expands the contents of JobSparkSqlJobLoggingConfig into a JSON
// request object.
func expandJobSparkSqlJobLoggingConfigSlice(c *Client, f []JobSparkSqlJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobSparkSqlJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobSparkSqlJobLoggingConfigMap flattens the contents of JobSparkSqlJobLoggingConfig from a JSON
// response object.
func flattenJobSparkSqlJobLoggingConfigMap(c *Client, i interface{}) map[string]JobSparkSqlJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobSparkSqlJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]JobSparkSqlJobLoggingConfig{}
	}

	items := make(map[string]JobSparkSqlJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenJobSparkSqlJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobSparkSqlJobLoggingConfigSlice flattens the contents of JobSparkSqlJobLoggingConfig from a JSON
// response object.
func flattenJobSparkSqlJobLoggingConfigSlice(c *Client, i interface{}) []JobSparkSqlJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobSparkSqlJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []JobSparkSqlJobLoggingConfig{}
	}

	items := make([]JobSparkSqlJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobSparkSqlJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobSparkSqlJobLoggingConfig expands an instance of JobSparkSqlJobLoggingConfig into a JSON
// request object.
func expandJobSparkSqlJobLoggingConfig(c *Client, f *JobSparkSqlJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenJobSparkSqlJobLoggingConfig flattens an instance of JobSparkSqlJobLoggingConfig from a JSON
// response object.
func flattenJobSparkSqlJobLoggingConfig(c *Client, i interface{}) *JobSparkSqlJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobSparkSqlJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandJobPrestoJobMap expands the contents of JobPrestoJob into a JSON
// request object.
func expandJobPrestoJobMap(c *Client, f map[string]JobPrestoJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPrestoJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPrestoJobSlice expands the contents of JobPrestoJob into a JSON
// request object.
func expandJobPrestoJobSlice(c *Client, f []JobPrestoJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPrestoJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPrestoJobMap flattens the contents of JobPrestoJob from a JSON
// response object.
func flattenJobPrestoJobMap(c *Client, i interface{}) map[string]JobPrestoJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPrestoJob{}
	}

	if len(a) == 0 {
		return map[string]JobPrestoJob{}
	}

	items := make(map[string]JobPrestoJob)
	for k, item := range a {
		items[k] = *flattenJobPrestoJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPrestoJobSlice flattens the contents of JobPrestoJob from a JSON
// response object.
func flattenJobPrestoJobSlice(c *Client, i interface{}) []JobPrestoJob {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPrestoJob{}
	}

	if len(a) == 0 {
		return []JobPrestoJob{}
	}

	items := make([]JobPrestoJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPrestoJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPrestoJob expands an instance of JobPrestoJob into a JSON
// request object.
func expandJobPrestoJob(c *Client, f *JobPrestoJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["queryFileUri"] = v
	}
	if v, err := expandJobPrestoJobQueryList(c, f.QueryList); err != nil {
		return nil, fmt.Errorf("error expanding QueryList into queryList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["queryList"] = v
	}
	if v := f.ContinueOnFailure; !dcl.IsEmptyValueIndirect(v) {
		m["continueOnFailure"] = v
	}
	if v := f.OutputFormat; !dcl.IsEmptyValueIndirect(v) {
		m["outputFormat"] = v
	}
	if v := f.ClientTags; !dcl.IsEmptyValueIndirect(v) {
		m["clientTags"] = v
	}
	if v := f.Properties; !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v, err := expandJobPrestoJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenJobPrestoJob flattens an instance of JobPrestoJob from a JSON
// response object.
func flattenJobPrestoJob(c *Client, i interface{}) *JobPrestoJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPrestoJob{}
	r.QueryFileUri = dcl.FlattenString(m["queryFileUri"])
	r.QueryList = flattenJobPrestoJobQueryList(c, m["queryList"])
	r.ContinueOnFailure = dcl.FlattenBool(m["continueOnFailure"])
	r.OutputFormat = dcl.FlattenString(m["outputFormat"])
	r.ClientTags = dcl.FlattenStringSlice(m["clientTags"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenJobPrestoJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandJobPrestoJobQueryListMap expands the contents of JobPrestoJobQueryList into a JSON
// request object.
func expandJobPrestoJobQueryListMap(c *Client, f map[string]JobPrestoJobQueryList) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPrestoJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPrestoJobQueryListSlice expands the contents of JobPrestoJobQueryList into a JSON
// request object.
func expandJobPrestoJobQueryListSlice(c *Client, f []JobPrestoJobQueryList) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPrestoJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPrestoJobQueryListMap flattens the contents of JobPrestoJobQueryList from a JSON
// response object.
func flattenJobPrestoJobQueryListMap(c *Client, i interface{}) map[string]JobPrestoJobQueryList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPrestoJobQueryList{}
	}

	if len(a) == 0 {
		return map[string]JobPrestoJobQueryList{}
	}

	items := make(map[string]JobPrestoJobQueryList)
	for k, item := range a {
		items[k] = *flattenJobPrestoJobQueryList(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPrestoJobQueryListSlice flattens the contents of JobPrestoJobQueryList from a JSON
// response object.
func flattenJobPrestoJobQueryListSlice(c *Client, i interface{}) []JobPrestoJobQueryList {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPrestoJobQueryList{}
	}

	if len(a) == 0 {
		return []JobPrestoJobQueryList{}
	}

	items := make([]JobPrestoJobQueryList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPrestoJobQueryList(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPrestoJobQueryList expands an instance of JobPrestoJobQueryList into a JSON
// request object.
func expandJobPrestoJobQueryList(c *Client, f *JobPrestoJobQueryList) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Queries; !dcl.IsEmptyValueIndirect(v) {
		m["queries"] = v
	}

	return m, nil
}

// flattenJobPrestoJobQueryList flattens an instance of JobPrestoJobQueryList from a JSON
// response object.
func flattenJobPrestoJobQueryList(c *Client, i interface{}) *JobPrestoJobQueryList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPrestoJobQueryList{}
	r.Queries = dcl.FlattenStringSlice(m["queries"])

	return r
}

// expandJobPrestoJobLoggingConfigMap expands the contents of JobPrestoJobLoggingConfig into a JSON
// request object.
func expandJobPrestoJobLoggingConfigMap(c *Client, f map[string]JobPrestoJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPrestoJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPrestoJobLoggingConfigSlice expands the contents of JobPrestoJobLoggingConfig into a JSON
// request object.
func expandJobPrestoJobLoggingConfigSlice(c *Client, f []JobPrestoJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPrestoJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPrestoJobLoggingConfigMap flattens the contents of JobPrestoJobLoggingConfig from a JSON
// response object.
func flattenJobPrestoJobLoggingConfigMap(c *Client, i interface{}) map[string]JobPrestoJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPrestoJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]JobPrestoJobLoggingConfig{}
	}

	items := make(map[string]JobPrestoJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenJobPrestoJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobPrestoJobLoggingConfigSlice flattens the contents of JobPrestoJobLoggingConfig from a JSON
// response object.
func flattenJobPrestoJobLoggingConfigSlice(c *Client, i interface{}) []JobPrestoJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPrestoJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []JobPrestoJobLoggingConfig{}
	}

	items := make([]JobPrestoJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPrestoJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobPrestoJobLoggingConfig expands an instance of JobPrestoJobLoggingConfig into a JSON
// request object.
func expandJobPrestoJobLoggingConfig(c *Client, f *JobPrestoJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenJobPrestoJobLoggingConfig flattens an instance of JobPrestoJobLoggingConfig from a JSON
// response object.
func flattenJobPrestoJobLoggingConfig(c *Client, i interface{}) *JobPrestoJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPrestoJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

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
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Details; !dcl.IsEmptyValueIndirect(v) {
		m["details"] = v
	}
	if v := f.StateStartTime; !dcl.IsEmptyValueIndirect(v) {
		m["stateStartTime"] = v
	}
	if v := f.Substate; !dcl.IsEmptyValueIndirect(v) {
		m["substate"] = v
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
	r.State = flattenJobStatusStateEnum(m["state"])
	r.Details = dcl.FlattenString(m["details"])
	r.StateStartTime = dcl.FlattenString(m["stateStartTime"])
	r.Substate = flattenJobStatusSubstateEnum(m["substate"])

	return r
}

// expandJobStatusHistoryMap expands the contents of JobStatusHistory into a JSON
// request object.
func expandJobStatusHistoryMap(c *Client, f map[string]JobStatusHistory) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobStatusHistory(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobStatusHistorySlice expands the contents of JobStatusHistory into a JSON
// request object.
func expandJobStatusHistorySlice(c *Client, f []JobStatusHistory) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobStatusHistory(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobStatusHistoryMap flattens the contents of JobStatusHistory from a JSON
// response object.
func flattenJobStatusHistoryMap(c *Client, i interface{}) map[string]JobStatusHistory {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobStatusHistory{}
	}

	if len(a) == 0 {
		return map[string]JobStatusHistory{}
	}

	items := make(map[string]JobStatusHistory)
	for k, item := range a {
		items[k] = *flattenJobStatusHistory(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobStatusHistorySlice flattens the contents of JobStatusHistory from a JSON
// response object.
func flattenJobStatusHistorySlice(c *Client, i interface{}) []JobStatusHistory {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStatusHistory{}
	}

	if len(a) == 0 {
		return []JobStatusHistory{}
	}

	items := make([]JobStatusHistory, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStatusHistory(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobStatusHistory expands an instance of JobStatusHistory into a JSON
// request object.
func expandJobStatusHistory(c *Client, f *JobStatusHistory) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Details; !dcl.IsEmptyValueIndirect(v) {
		m["details"] = v
	}
	if v := f.StateStartTime; !dcl.IsEmptyValueIndirect(v) {
		m["stateStartTime"] = v
	}
	if v := f.Substate; !dcl.IsEmptyValueIndirect(v) {
		m["substate"] = v
	}

	return m, nil
}

// flattenJobStatusHistory flattens an instance of JobStatusHistory from a JSON
// response object.
func flattenJobStatusHistory(c *Client, i interface{}) *JobStatusHistory {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobStatusHistory{}
	r.State = flattenJobStatusHistoryStateEnum(m["state"])
	r.Details = dcl.FlattenString(m["details"])
	r.StateStartTime = dcl.FlattenString(m["stateStartTime"])
	r.Substate = flattenJobStatusHistorySubstateEnum(m["substate"])

	return r
}

// expandJobYarnApplicationsMap expands the contents of JobYarnApplications into a JSON
// request object.
func expandJobYarnApplicationsMap(c *Client, f map[string]JobYarnApplications) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobYarnApplications(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobYarnApplicationsSlice expands the contents of JobYarnApplications into a JSON
// request object.
func expandJobYarnApplicationsSlice(c *Client, f []JobYarnApplications) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobYarnApplications(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobYarnApplicationsMap flattens the contents of JobYarnApplications from a JSON
// response object.
func flattenJobYarnApplicationsMap(c *Client, i interface{}) map[string]JobYarnApplications {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobYarnApplications{}
	}

	if len(a) == 0 {
		return map[string]JobYarnApplications{}
	}

	items := make(map[string]JobYarnApplications)
	for k, item := range a {
		items[k] = *flattenJobYarnApplications(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobYarnApplicationsSlice flattens the contents of JobYarnApplications from a JSON
// response object.
func flattenJobYarnApplicationsSlice(c *Client, i interface{}) []JobYarnApplications {
	a, ok := i.([]interface{})
	if !ok {
		return []JobYarnApplications{}
	}

	if len(a) == 0 {
		return []JobYarnApplications{}
	}

	items := make([]JobYarnApplications, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobYarnApplications(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobYarnApplications expands an instance of JobYarnApplications into a JSON
// request object.
func expandJobYarnApplications(c *Client, f *JobYarnApplications) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Progress; !dcl.IsEmptyValueIndirect(v) {
		m["progress"] = v
	}
	if v := f.TrackingUrl; !dcl.IsEmptyValueIndirect(v) {
		m["trackingUrl"] = v
	}

	return m, nil
}

// flattenJobYarnApplications flattens an instance of JobYarnApplications from a JSON
// response object.
func flattenJobYarnApplications(c *Client, i interface{}) *JobYarnApplications {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobYarnApplications{}
	r.Name = dcl.FlattenString(m["name"])
	r.State = flattenJobYarnApplicationsStateEnum(m["state"])
	r.Progress = dcl.FlattenDouble(m["progress"])
	r.TrackingUrl = dcl.FlattenString(m["trackingUrl"])

	return r
}

// expandJobSchedulingMap expands the contents of JobScheduling into a JSON
// request object.
func expandJobSchedulingMap(c *Client, f map[string]JobScheduling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobScheduling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobSchedulingSlice expands the contents of JobScheduling into a JSON
// request object.
func expandJobSchedulingSlice(c *Client, f []JobScheduling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobScheduling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobSchedulingMap flattens the contents of JobScheduling from a JSON
// response object.
func flattenJobSchedulingMap(c *Client, i interface{}) map[string]JobScheduling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobScheduling{}
	}

	if len(a) == 0 {
		return map[string]JobScheduling{}
	}

	items := make(map[string]JobScheduling)
	for k, item := range a {
		items[k] = *flattenJobScheduling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobSchedulingSlice flattens the contents of JobScheduling from a JSON
// response object.
func flattenJobSchedulingSlice(c *Client, i interface{}) []JobScheduling {
	a, ok := i.([]interface{})
	if !ok {
		return []JobScheduling{}
	}

	if len(a) == 0 {
		return []JobScheduling{}
	}

	items := make([]JobScheduling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobScheduling(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobScheduling expands an instance of JobScheduling into a JSON
// request object.
func expandJobScheduling(c *Client, f *JobScheduling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxFailuresPerHour; !dcl.IsEmptyValueIndirect(v) {
		m["maxFailuresPerHour"] = v
	}
	if v := f.MaxFailuresTotal; !dcl.IsEmptyValueIndirect(v) {
		m["maxFailuresTotal"] = v
	}

	return m, nil
}

// flattenJobScheduling flattens an instance of JobScheduling from a JSON
// response object.
func flattenJobScheduling(c *Client, i interface{}) *JobScheduling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobScheduling{}
	r.MaxFailuresPerHour = dcl.FlattenInteger(m["maxFailuresPerHour"])
	r.MaxFailuresTotal = dcl.FlattenInteger(m["maxFailuresTotal"])

	return r
}

// expandJobDriverRunnerMap expands the contents of JobDriverRunner into a JSON
// request object.
func expandJobDriverRunnerMap(c *Client, f map[string]JobDriverRunner) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobDriverRunner(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobDriverRunnerSlice expands the contents of JobDriverRunner into a JSON
// request object.
func expandJobDriverRunnerSlice(c *Client, f []JobDriverRunner) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobDriverRunner(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobDriverRunnerMap flattens the contents of JobDriverRunner from a JSON
// response object.
func flattenJobDriverRunnerMap(c *Client, i interface{}) map[string]JobDriverRunner {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobDriverRunner{}
	}

	if len(a) == 0 {
		return map[string]JobDriverRunner{}
	}

	items := make(map[string]JobDriverRunner)
	for k, item := range a {
		items[k] = *flattenJobDriverRunner(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobDriverRunnerSlice flattens the contents of JobDriverRunner from a JSON
// response object.
func flattenJobDriverRunnerSlice(c *Client, i interface{}) []JobDriverRunner {
	a, ok := i.([]interface{})
	if !ok {
		return []JobDriverRunner{}
	}

	if len(a) == 0 {
		return []JobDriverRunner{}
	}

	items := make([]JobDriverRunner, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobDriverRunner(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobDriverRunner expands an instance of JobDriverRunner into a JSON
// request object.
func expandJobDriverRunner(c *Client, f *JobDriverRunner) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobDriverRunnerMasterDriverRunner(c, f.MasterDriverRunner); err != nil {
		return nil, fmt.Errorf("error expanding MasterDriverRunner into masterDriverRunner: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["masterDriverRunner"] = v
	}
	if v, err := expandJobDriverRunnerYarnDriverRunner(c, f.YarnDriverRunner); err != nil {
		return nil, fmt.Errorf("error expanding YarnDriverRunner into yarnDriverRunner: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["yarnDriverRunner"] = v
	}

	return m, nil
}

// flattenJobDriverRunner flattens an instance of JobDriverRunner from a JSON
// response object.
func flattenJobDriverRunner(c *Client, i interface{}) *JobDriverRunner {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobDriverRunner{}
	r.MasterDriverRunner = flattenJobDriverRunnerMasterDriverRunner(c, m["masterDriverRunner"])
	r.YarnDriverRunner = flattenJobDriverRunnerYarnDriverRunner(c, m["yarnDriverRunner"])

	return r
}

// expandJobDriverRunnerMasterDriverRunnerMap expands the contents of JobDriverRunnerMasterDriverRunner into a JSON
// request object.
func expandJobDriverRunnerMasterDriverRunnerMap(c *Client, f map[string]JobDriverRunnerMasterDriverRunner) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobDriverRunnerMasterDriverRunner(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobDriverRunnerMasterDriverRunnerSlice expands the contents of JobDriverRunnerMasterDriverRunner into a JSON
// request object.
func expandJobDriverRunnerMasterDriverRunnerSlice(c *Client, f []JobDriverRunnerMasterDriverRunner) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobDriverRunnerMasterDriverRunner(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobDriverRunnerMasterDriverRunnerMap flattens the contents of JobDriverRunnerMasterDriverRunner from a JSON
// response object.
func flattenJobDriverRunnerMasterDriverRunnerMap(c *Client, i interface{}) map[string]JobDriverRunnerMasterDriverRunner {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobDriverRunnerMasterDriverRunner{}
	}

	if len(a) == 0 {
		return map[string]JobDriverRunnerMasterDriverRunner{}
	}

	items := make(map[string]JobDriverRunnerMasterDriverRunner)
	for k, item := range a {
		items[k] = *flattenJobDriverRunnerMasterDriverRunner(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobDriverRunnerMasterDriverRunnerSlice flattens the contents of JobDriverRunnerMasterDriverRunner from a JSON
// response object.
func flattenJobDriverRunnerMasterDriverRunnerSlice(c *Client, i interface{}) []JobDriverRunnerMasterDriverRunner {
	a, ok := i.([]interface{})
	if !ok {
		return []JobDriverRunnerMasterDriverRunner{}
	}

	if len(a) == 0 {
		return []JobDriverRunnerMasterDriverRunner{}
	}

	items := make([]JobDriverRunnerMasterDriverRunner, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobDriverRunnerMasterDriverRunner(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobDriverRunnerMasterDriverRunner expands an instance of JobDriverRunnerMasterDriverRunner into a JSON
// request object.
func expandJobDriverRunnerMasterDriverRunner(c *Client, f *JobDriverRunnerMasterDriverRunner) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenJobDriverRunnerMasterDriverRunner flattens an instance of JobDriverRunnerMasterDriverRunner from a JSON
// response object.
func flattenJobDriverRunnerMasterDriverRunner(c *Client, i interface{}) *JobDriverRunnerMasterDriverRunner {

	r := &JobDriverRunnerMasterDriverRunner{}

	return r
}

// expandJobDriverRunnerYarnDriverRunnerMap expands the contents of JobDriverRunnerYarnDriverRunner into a JSON
// request object.
func expandJobDriverRunnerYarnDriverRunnerMap(c *Client, f map[string]JobDriverRunnerYarnDriverRunner) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobDriverRunnerYarnDriverRunner(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobDriverRunnerYarnDriverRunnerSlice expands the contents of JobDriverRunnerYarnDriverRunner into a JSON
// request object.
func expandJobDriverRunnerYarnDriverRunnerSlice(c *Client, f []JobDriverRunnerYarnDriverRunner) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobDriverRunnerYarnDriverRunner(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobDriverRunnerYarnDriverRunnerMap flattens the contents of JobDriverRunnerYarnDriverRunner from a JSON
// response object.
func flattenJobDriverRunnerYarnDriverRunnerMap(c *Client, i interface{}) map[string]JobDriverRunnerYarnDriverRunner {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobDriverRunnerYarnDriverRunner{}
	}

	if len(a) == 0 {
		return map[string]JobDriverRunnerYarnDriverRunner{}
	}

	items := make(map[string]JobDriverRunnerYarnDriverRunner)
	for k, item := range a {
		items[k] = *flattenJobDriverRunnerYarnDriverRunner(c, item.(map[string]interface{}))
	}

	return items
}

// flattenJobDriverRunnerYarnDriverRunnerSlice flattens the contents of JobDriverRunnerYarnDriverRunner from a JSON
// response object.
func flattenJobDriverRunnerYarnDriverRunnerSlice(c *Client, i interface{}) []JobDriverRunnerYarnDriverRunner {
	a, ok := i.([]interface{})
	if !ok {
		return []JobDriverRunnerYarnDriverRunner{}
	}

	if len(a) == 0 {
		return []JobDriverRunnerYarnDriverRunner{}
	}

	items := make([]JobDriverRunnerYarnDriverRunner, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobDriverRunnerYarnDriverRunner(c, item.(map[string]interface{})))
	}

	return items
}

// expandJobDriverRunnerYarnDriverRunner expands an instance of JobDriverRunnerYarnDriverRunner into a JSON
// request object.
func expandJobDriverRunnerYarnDriverRunner(c *Client, f *JobDriverRunnerYarnDriverRunner) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MemoryMb; !dcl.IsEmptyValueIndirect(v) {
		m["memoryMb"] = v
	}
	if v := f.Vcores; !dcl.IsEmptyValueIndirect(v) {
		m["vcores"] = v
	}

	return m, nil
}

// flattenJobDriverRunnerYarnDriverRunner flattens an instance of JobDriverRunnerYarnDriverRunner from a JSON
// response object.
func flattenJobDriverRunnerYarnDriverRunner(c *Client, i interface{}) *JobDriverRunnerYarnDriverRunner {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobDriverRunnerYarnDriverRunner{}
	r.MemoryMb = dcl.FlattenInteger(m["memoryMb"])
	r.Vcores = dcl.FlattenInteger(m["vcores"])

	return r
}

// flattenJobStatusStateEnumSlice flattens the contents of JobStatusStateEnum from a JSON
// response object.
func flattenJobStatusStateEnumSlice(c *Client, i interface{}) []JobStatusStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStatusStateEnum{}
	}

	if len(a) == 0 {
		return []JobStatusStateEnum{}
	}

	items := make([]JobStatusStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStatusStateEnum(item.(interface{})))
	}

	return items
}

// flattenJobStatusStateEnum asserts that an interface is a string, and returns a
// pointer to a *JobStatusStateEnum with the same value as that string.
func flattenJobStatusStateEnum(i interface{}) *JobStatusStateEnum {
	s, ok := i.(string)
	if !ok {
		return JobStatusStateEnumRef("")
	}

	return JobStatusStateEnumRef(s)
}

// flattenJobStatusSubstateEnumSlice flattens the contents of JobStatusSubstateEnum from a JSON
// response object.
func flattenJobStatusSubstateEnumSlice(c *Client, i interface{}) []JobStatusSubstateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStatusSubstateEnum{}
	}

	if len(a) == 0 {
		return []JobStatusSubstateEnum{}
	}

	items := make([]JobStatusSubstateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStatusSubstateEnum(item.(interface{})))
	}

	return items
}

// flattenJobStatusSubstateEnum asserts that an interface is a string, and returns a
// pointer to a *JobStatusSubstateEnum with the same value as that string.
func flattenJobStatusSubstateEnum(i interface{}) *JobStatusSubstateEnum {
	s, ok := i.(string)
	if !ok {
		return JobStatusSubstateEnumRef("")
	}

	return JobStatusSubstateEnumRef(s)
}

// flattenJobStatusHistoryStateEnumSlice flattens the contents of JobStatusHistoryStateEnum from a JSON
// response object.
func flattenJobStatusHistoryStateEnumSlice(c *Client, i interface{}) []JobStatusHistoryStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStatusHistoryStateEnum{}
	}

	if len(a) == 0 {
		return []JobStatusHistoryStateEnum{}
	}

	items := make([]JobStatusHistoryStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStatusHistoryStateEnum(item.(interface{})))
	}

	return items
}

// flattenJobStatusHistoryStateEnum asserts that an interface is a string, and returns a
// pointer to a *JobStatusHistoryStateEnum with the same value as that string.
func flattenJobStatusHistoryStateEnum(i interface{}) *JobStatusHistoryStateEnum {
	s, ok := i.(string)
	if !ok {
		return JobStatusHistoryStateEnumRef("")
	}

	return JobStatusHistoryStateEnumRef(s)
}

// flattenJobStatusHistorySubstateEnumSlice flattens the contents of JobStatusHistorySubstateEnum from a JSON
// response object.
func flattenJobStatusHistorySubstateEnumSlice(c *Client, i interface{}) []JobStatusHistorySubstateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStatusHistorySubstateEnum{}
	}

	if len(a) == 0 {
		return []JobStatusHistorySubstateEnum{}
	}

	items := make([]JobStatusHistorySubstateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStatusHistorySubstateEnum(item.(interface{})))
	}

	return items
}

// flattenJobStatusHistorySubstateEnum asserts that an interface is a string, and returns a
// pointer to a *JobStatusHistorySubstateEnum with the same value as that string.
func flattenJobStatusHistorySubstateEnum(i interface{}) *JobStatusHistorySubstateEnum {
	s, ok := i.(string)
	if !ok {
		return JobStatusHistorySubstateEnumRef("")
	}

	return JobStatusHistorySubstateEnumRef(s)
}

// flattenJobYarnApplicationsStateEnumSlice flattens the contents of JobYarnApplicationsStateEnum from a JSON
// response object.
func flattenJobYarnApplicationsStateEnumSlice(c *Client, i interface{}) []JobYarnApplicationsStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobYarnApplicationsStateEnum{}
	}

	if len(a) == 0 {
		return []JobYarnApplicationsStateEnum{}
	}

	items := make([]JobYarnApplicationsStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobYarnApplicationsStateEnum(item.(interface{})))
	}

	return items
}

// flattenJobYarnApplicationsStateEnum asserts that an interface is a string, and returns a
// pointer to a *JobYarnApplicationsStateEnum with the same value as that string.
func flattenJobYarnApplicationsStateEnum(i interface{}) *JobYarnApplicationsStateEnum {
	s, ok := i.(string)
	if !ok {
		return JobYarnApplicationsStateEnumRef("")
	}

	return JobYarnApplicationsStateEnumRef(s)
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
		if nr.Region == nil && ncr.Region == nil {
			c.Config.Logger.Info("Both Region fields null - considering equal.")
		} else if nr.Region == nil || ncr.Region == nil {
			c.Config.Logger.Info("Only one Region field is null - considering unequal.")
			return false
		} else if *nr.Region != *ncr.Region {
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
