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
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *WorkflowTemplate) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "placement"); err != nil {
		return err
	}
	if err := dcl.Required(r, "jobs"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Placement) {
		if err := r.Placement.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkflowTemplatePlacement) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ManagedCluster) {
		if err := r.ManagedCluster.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ClusterSelector) {
		if err := r.ClusterSelector.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkflowTemplatePlacementManagedCluster) validate() error {
	if err := dcl.Required(r, "clusterName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "config"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Config) {
		if err := r.Config.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkflowTemplatePlacementClusterSelector) validate() error {
	if err := dcl.Required(r, "clusterLabels"); err != nil {
		return err
	}
	return nil
}
func (r *WorkflowTemplateJobs) validate() error {
	if err := dcl.Required(r, "stepId"); err != nil {
		return err
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
	if !dcl.IsEmptyValueIndirect(r.Scheduling) {
		if err := r.Scheduling.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkflowTemplateJobsHadoopJob) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkflowTemplateJobsHadoopJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkJob) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkflowTemplateJobsSparkJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsPysparkJob) validate() error {
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
func (r *WorkflowTemplateJobsPysparkJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsHiveJob) validate() error {
	if !dcl.IsEmptyValueIndirect(r.QueryList) {
		if err := r.QueryList.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkflowTemplateJobsHiveJobQueryList) validate() error {
	if err := dcl.Required(r, "queries"); err != nil {
		return err
	}
	return nil
}
func (r *WorkflowTemplateJobsPigJob) validate() error {
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
func (r *WorkflowTemplateJobsPigJobQueryList) validate() error {
	if err := dcl.Required(r, "queries"); err != nil {
		return err
	}
	return nil
}
func (r *WorkflowTemplateJobsPigJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkRJob) validate() error {
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
func (r *WorkflowTemplateJobsSparkRJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkSqlJob) validate() error {
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
func (r *WorkflowTemplateJobsSparkSqlJobQueryList) validate() error {
	if err := dcl.Required(r, "queries"); err != nil {
		return err
	}
	return nil
}
func (r *WorkflowTemplateJobsSparkSqlJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsPrestoJob) validate() error {
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
func (r *WorkflowTemplateJobsPrestoJobQueryList) validate() error {
	if err := dcl.Required(r, "queries"); err != nil {
		return err
	}
	return nil
}
func (r *WorkflowTemplateJobsPrestoJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsScheduling) validate() error {
	return nil
}
func (r *WorkflowTemplateParameters) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "fields"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Validation) {
		if err := r.Validation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkflowTemplateParametersValidation) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Regex) {
		if err := r.Regex.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Values) {
		if err := r.Values.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkflowTemplateParametersValidationRegex) validate() error {
	if err := dcl.Required(r, "regexes"); err != nil {
		return err
	}
	return nil
}
func (r *WorkflowTemplateParametersValidationValues) validate() error {
	if err := dcl.Required(r, "values"); err != nil {
		return err
	}
	return nil
}

func workflowTemplateGetURL(userBasePath string, r *WorkflowTemplate) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}", "https://dataproc.googleapis.com/v1beta2/", userBasePath, params), nil
}

func workflowTemplateListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates", "https://dataproc.googleapis.com/v1beta2/", userBasePath, params), nil

}

func workflowTemplateCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates", "https://dataproc.googleapis.com/v1beta2/", userBasePath, params), nil

}

func workflowTemplateDeleteURL(userBasePath string, r *WorkflowTemplate) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}", "https://dataproc.googleapis.com/v1beta2/", userBasePath, params), nil
}

// workflowTemplateApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type workflowTemplateApiOperation interface {
	do(context.Context, *WorkflowTemplate, *Client) error
}

// newUpdateWorkflowTemplateUpdateWorkflowTemplateRequest creates a request for an
// WorkflowTemplate resource's UpdateWorkflowTemplate update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateWorkflowTemplateUpdateWorkflowTemplateRequest(ctx context.Context, f *WorkflowTemplate, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateWorkflowTemplateUpdateWorkflowTemplateRequest converts the update into
// the final JSON request body.
func marshalUpdateWorkflowTemplateUpdateWorkflowTemplateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateWorkflowTemplateUpdateWorkflowTemplateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateWorkflowTemplateUpdateWorkflowTemplateOperation) do(ctx context.Context, r *WorkflowTemplate, c *Client) error {
	_, err := c.GetWorkflowTemplate(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateWorkflowTemplate")
	if err != nil {
		return err
	}

	req, err := newUpdateWorkflowTemplateUpdateWorkflowTemplateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateWorkflowTemplateUpdateWorkflowTemplateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listWorkflowTemplateRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := workflowTemplateListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != WorkflowTemplateMaxPage {
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

type listWorkflowTemplateOperation struct {
	Templates []map[string]interface{} `json:"templates"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listWorkflowTemplate(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*WorkflowTemplate, string, error) {
	b, err := c.listWorkflowTemplateRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listWorkflowTemplateOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*WorkflowTemplate
	for _, v := range m.Templates {
		res := flattenWorkflowTemplate(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllWorkflowTemplate(ctx context.Context, f func(*WorkflowTemplate) bool, resources []*WorkflowTemplate) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteWorkflowTemplate(ctx, res)
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

type deleteWorkflowTemplateOperation struct{}

func (op *deleteWorkflowTemplateOperation) do(ctx context.Context, r *WorkflowTemplate, c *Client) error {

	_, err := c.GetWorkflowTemplate(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("WorkflowTemplate not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetWorkflowTemplate checking for existence. error: %v", err)
		return err
	}

	u, err := workflowTemplateDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete WorkflowTemplate: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetWorkflowTemplate(ctx, r.urlNormalized())
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
type createWorkflowTemplateOperation struct {
	response map[string]interface{}
}

func (op *createWorkflowTemplateOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createWorkflowTemplateOperation) do(ctx context.Context, r *WorkflowTemplate, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := workflowTemplateCreateURL(c.Config.BasePath, project, location)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(req, &m); err != nil {
		return err
	}
	normalized := r.urlNormalized()
	m["id"] = fmt.Sprintf("%s", *normalized.Name)

	req, err = json.Marshal(m)
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

	if _, err := c.GetWorkflowTemplate(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getWorkflowTemplateRaw(ctx context.Context, r *WorkflowTemplate) ([]byte, error) {

	u, err := workflowTemplateGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) workflowTemplateDiffsForRawDesired(ctx context.Context, rawDesired *WorkflowTemplate, opts ...dcl.ApplyOption) (initial, desired *WorkflowTemplate, diffs []workflowTemplateDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *WorkflowTemplate
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*WorkflowTemplate); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected WorkflowTemplate, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetWorkflowTemplate(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a WorkflowTemplate resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve WorkflowTemplate resource: %v", err)
		}
		c.Config.Logger.Info("Found that WorkflowTemplate resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeWorkflowTemplateDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for WorkflowTemplate: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for WorkflowTemplate: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeWorkflowTemplateInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for WorkflowTemplate: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeWorkflowTemplateDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for WorkflowTemplate: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffWorkflowTemplate(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeWorkflowTemplateInitialState(rawInitial, rawDesired *WorkflowTemplate) (*WorkflowTemplate, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeWorkflowTemplateDesiredState(rawDesired, rawInitial *WorkflowTemplate, opts ...dcl.ApplyOption) (*WorkflowTemplate, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Placement = canonicalizeWorkflowTemplatePlacement(rawDesired.Placement, nil, opts...)

		return rawDesired, nil
	}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Version) {
		rawDesired.Version = rawInitial.Version
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	rawDesired.Placement = canonicalizeWorkflowTemplatePlacement(rawDesired.Placement, rawInitial.Placement, opts...)
	if dcl.IsZeroValue(rawDesired.Jobs) {
		rawDesired.Jobs = rawInitial.Jobs
	}
	if dcl.IsZeroValue(rawDesired.Parameters) {
		rawDesired.Parameters = rawInitial.Parameters
	}
	if dcl.StringCanonicalize(rawDesired.DagTimeout, rawInitial.DagTimeout) {
		rawDesired.DagTimeout = rawInitial.DagTimeout
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeWorkflowTemplateNewState(c *Client, rawNew, rawDesired *WorkflowTemplate) (*WorkflowTemplate, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Version) && dcl.IsEmptyValueIndirect(rawDesired.Version) {
		rawNew.Version = rawDesired.Version
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

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Placement) && dcl.IsEmptyValueIndirect(rawDesired.Placement) {
		rawNew.Placement = rawDesired.Placement
	} else {
		rawNew.Placement = canonicalizeNewWorkflowTemplatePlacement(c, rawDesired.Placement, rawNew.Placement)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Jobs) && dcl.IsEmptyValueIndirect(rawDesired.Jobs) {
		rawNew.Jobs = rawDesired.Jobs
	} else {
		rawNew.Jobs = canonicalizeNewWorkflowTemplateJobsSlice(c, rawDesired.Jobs, rawNew.Jobs)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Parameters) && dcl.IsEmptyValueIndirect(rawDesired.Parameters) {
		rawNew.Parameters = rawDesired.Parameters
	} else {
		rawNew.Parameters = canonicalizeNewWorkflowTemplateParametersSlice(c, rawDesired.Parameters, rawNew.Parameters)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DagTimeout) && dcl.IsEmptyValueIndirect(rawDesired.DagTimeout) {
		rawNew.DagTimeout = rawDesired.DagTimeout
	} else {
		if dcl.StringCanonicalize(rawDesired.DagTimeout, rawNew.DagTimeout) {
			rawNew.DagTimeout = rawDesired.DagTimeout
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeWorkflowTemplatePlacement(des, initial *WorkflowTemplatePlacement, opts ...dcl.ApplyOption) *WorkflowTemplatePlacement {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.ManagedCluster = canonicalizeWorkflowTemplatePlacementManagedCluster(des.ManagedCluster, initial.ManagedCluster, opts...)
	des.ClusterSelector = canonicalizeWorkflowTemplatePlacementClusterSelector(des.ClusterSelector, initial.ClusterSelector, opts...)

	return des
}

func canonicalizeNewWorkflowTemplatePlacement(c *Client, des, nw *WorkflowTemplatePlacement) *WorkflowTemplatePlacement {
	if des == nil || nw == nil {
		return nw
	}

	nw.ManagedCluster = canonicalizeNewWorkflowTemplatePlacementManagedCluster(c, des.ManagedCluster, nw.ManagedCluster)
	nw.ClusterSelector = canonicalizeNewWorkflowTemplatePlacementClusterSelector(c, des.ClusterSelector, nw.ClusterSelector)

	return nw
}

func canonicalizeNewWorkflowTemplatePlacementSet(c *Client, des, nw []WorkflowTemplatePlacement) []WorkflowTemplatePlacement {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplatePlacement
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplatePlacement(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplatePlacementSlice(c *Client, des, nw []WorkflowTemplatePlacement) []WorkflowTemplatePlacement {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplatePlacement
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplatePlacement(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplatePlacementManagedCluster(des, initial *WorkflowTemplatePlacementManagedCluster, opts ...dcl.ApplyOption) *WorkflowTemplatePlacementManagedCluster {
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
	des.Config = canonicalizeClusterClusterConfig(des.Config, initial.Config, opts...)
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}

	return des
}

func canonicalizeNewWorkflowTemplatePlacementManagedCluster(c *Client, des, nw *WorkflowTemplatePlacementManagedCluster) *WorkflowTemplatePlacementManagedCluster {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ClusterName, nw.ClusterName) {
		nw.ClusterName = des.ClusterName
	}
	nw.Config = canonicalizeNewClusterClusterConfig(c, des.Config, nw.Config)

	return nw
}

func canonicalizeNewWorkflowTemplatePlacementManagedClusterSet(c *Client, des, nw []WorkflowTemplatePlacementManagedCluster) []WorkflowTemplatePlacementManagedCluster {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplatePlacementManagedCluster
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplatePlacementManagedCluster(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplatePlacementManagedClusterSlice(c *Client, des, nw []WorkflowTemplatePlacementManagedCluster) []WorkflowTemplatePlacementManagedCluster {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplatePlacementManagedCluster
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplatePlacementManagedCluster(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplatePlacementClusterSelector(des, initial *WorkflowTemplatePlacementClusterSelector, opts ...dcl.ApplyOption) *WorkflowTemplatePlacementClusterSelector {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Zone, initial.Zone) || dcl.IsZeroValue(des.Zone) {
		des.Zone = initial.Zone
	}
	if dcl.IsZeroValue(des.ClusterLabels) {
		des.ClusterLabels = initial.ClusterLabels
	}

	return des
}

func canonicalizeNewWorkflowTemplatePlacementClusterSelector(c *Client, des, nw *WorkflowTemplatePlacementClusterSelector) *WorkflowTemplatePlacementClusterSelector {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Zone, nw.Zone) {
		nw.Zone = des.Zone
	}

	return nw
}

func canonicalizeNewWorkflowTemplatePlacementClusterSelectorSet(c *Client, des, nw []WorkflowTemplatePlacementClusterSelector) []WorkflowTemplatePlacementClusterSelector {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplatePlacementClusterSelector
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplatePlacementClusterSelector(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplatePlacementClusterSelectorSlice(c *Client, des, nw []WorkflowTemplatePlacementClusterSelector) []WorkflowTemplatePlacementClusterSelector {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplatePlacementClusterSelector
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplatePlacementClusterSelector(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobs(des, initial *WorkflowTemplateJobs, opts ...dcl.ApplyOption) *WorkflowTemplateJobs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.StepId, initial.StepId) || dcl.IsZeroValue(des.StepId) {
		des.StepId = initial.StepId
	}
	des.HadoopJob = canonicalizeWorkflowTemplateJobsHadoopJob(des.HadoopJob, initial.HadoopJob, opts...)
	des.SparkJob = canonicalizeWorkflowTemplateJobsSparkJob(des.SparkJob, initial.SparkJob, opts...)
	des.PysparkJob = canonicalizeWorkflowTemplateJobsPysparkJob(des.PysparkJob, initial.PysparkJob, opts...)
	des.HiveJob = canonicalizeWorkflowTemplateJobsHiveJob(des.HiveJob, initial.HiveJob, opts...)
	des.PigJob = canonicalizeWorkflowTemplateJobsPigJob(des.PigJob, initial.PigJob, opts...)
	des.SparkRJob = canonicalizeWorkflowTemplateJobsSparkRJob(des.SparkRJob, initial.SparkRJob, opts...)
	des.SparkSqlJob = canonicalizeWorkflowTemplateJobsSparkSqlJob(des.SparkSqlJob, initial.SparkSqlJob, opts...)
	des.PrestoJob = canonicalizeWorkflowTemplateJobsPrestoJob(des.PrestoJob, initial.PrestoJob, opts...)
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	des.Scheduling = canonicalizeWorkflowTemplateJobsScheduling(des.Scheduling, initial.Scheduling, opts...)
	if dcl.IsZeroValue(des.PrerequisiteStepIds) {
		des.PrerequisiteStepIds = initial.PrerequisiteStepIds
	}

	return des
}

func canonicalizeNewWorkflowTemplateJobs(c *Client, des, nw *WorkflowTemplateJobs) *WorkflowTemplateJobs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.StepId, nw.StepId) {
		nw.StepId = des.StepId
	}
	nw.HadoopJob = canonicalizeNewWorkflowTemplateJobsHadoopJob(c, des.HadoopJob, nw.HadoopJob)
	nw.SparkJob = canonicalizeNewWorkflowTemplateJobsSparkJob(c, des.SparkJob, nw.SparkJob)
	nw.PysparkJob = canonicalizeNewWorkflowTemplateJobsPysparkJob(c, des.PysparkJob, nw.PysparkJob)
	nw.HiveJob = canonicalizeNewWorkflowTemplateJobsHiveJob(c, des.HiveJob, nw.HiveJob)
	nw.PigJob = canonicalizeNewWorkflowTemplateJobsPigJob(c, des.PigJob, nw.PigJob)
	nw.SparkRJob = canonicalizeNewWorkflowTemplateJobsSparkRJob(c, des.SparkRJob, nw.SparkRJob)
	nw.SparkSqlJob = canonicalizeNewWorkflowTemplateJobsSparkSqlJob(c, des.SparkSqlJob, nw.SparkSqlJob)
	nw.PrestoJob = canonicalizeNewWorkflowTemplateJobsPrestoJob(c, des.PrestoJob, nw.PrestoJob)
	nw.Scheduling = canonicalizeNewWorkflowTemplateJobsScheduling(c, des.Scheduling, nw.Scheduling)

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSet(c *Client, des, nw []WorkflowTemplateJobs) []WorkflowTemplateJobs {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobs(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsSlice(c *Client, des, nw []WorkflowTemplateJobs) []WorkflowTemplateJobs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobs(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsHadoopJob(des, initial *WorkflowTemplateJobsHadoopJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHadoopJob {
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
	des.LoggingConfig = canonicalizeWorkflowTemplateJobsHadoopJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateJobsHadoopJob(c *Client, des, nw *WorkflowTemplateJobsHadoopJob) *WorkflowTemplateJobsHadoopJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MainJarFileUri, nw.MainJarFileUri) {
		nw.MainJarFileUri = des.MainJarFileUri
	}
	if dcl.StringCanonicalize(des.MainClass, nw.MainClass) {
		nw.MainClass = des.MainClass
	}
	nw.LoggingConfig = canonicalizeNewWorkflowTemplateJobsHadoopJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewWorkflowTemplateJobsHadoopJobSet(c *Client, des, nw []WorkflowTemplateJobsHadoopJob) []WorkflowTemplateJobsHadoopJob {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsHadoopJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsHadoopJob(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsHadoopJobSlice(c *Client, des, nw []WorkflowTemplateJobsHadoopJob) []WorkflowTemplateJobsHadoopJob {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsHadoopJob
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsHadoopJob(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsHadoopJobLoggingConfig(des, initial *WorkflowTemplateJobsHadoopJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHadoopJobLoggingConfig {
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

func canonicalizeNewWorkflowTemplateJobsHadoopJobLoggingConfig(c *Client, des, nw *WorkflowTemplateJobsHadoopJobLoggingConfig) *WorkflowTemplateJobsHadoopJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsHadoopJobLoggingConfigSet(c *Client, des, nw []WorkflowTemplateJobsHadoopJobLoggingConfig) []WorkflowTemplateJobsHadoopJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsHadoopJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsHadoopJobLoggingConfig(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsHadoopJobLoggingConfigSlice(c *Client, des, nw []WorkflowTemplateJobsHadoopJobLoggingConfig) []WorkflowTemplateJobsHadoopJobLoggingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsHadoopJobLoggingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsHadoopJobLoggingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsSparkJob(des, initial *WorkflowTemplateJobsSparkJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkJob {
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
	des.LoggingConfig = canonicalizeWorkflowTemplateJobsSparkJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateJobsSparkJob(c *Client, des, nw *WorkflowTemplateJobsSparkJob) *WorkflowTemplateJobsSparkJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MainJarFileUri, nw.MainJarFileUri) {
		nw.MainJarFileUri = des.MainJarFileUri
	}
	if dcl.StringCanonicalize(des.MainClass, nw.MainClass) {
		nw.MainClass = des.MainClass
	}
	nw.LoggingConfig = canonicalizeNewWorkflowTemplateJobsSparkJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkJobSet(c *Client, des, nw []WorkflowTemplateJobsSparkJob) []WorkflowTemplateJobsSparkJob {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkJob(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsSparkJobSlice(c *Client, des, nw []WorkflowTemplateJobsSparkJob) []WorkflowTemplateJobsSparkJob {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsSparkJob
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsSparkJob(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsSparkJobLoggingConfig(des, initial *WorkflowTemplateJobsSparkJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkJobLoggingConfig {
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

func canonicalizeNewWorkflowTemplateJobsSparkJobLoggingConfig(c *Client, des, nw *WorkflowTemplateJobsSparkJobLoggingConfig) *WorkflowTemplateJobsSparkJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkJobLoggingConfigSet(c *Client, des, nw []WorkflowTemplateJobsSparkJobLoggingConfig) []WorkflowTemplateJobsSparkJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkJobLoggingConfig(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsSparkJobLoggingConfigSlice(c *Client, des, nw []WorkflowTemplateJobsSparkJobLoggingConfig) []WorkflowTemplateJobsSparkJobLoggingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsSparkJobLoggingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsSparkJobLoggingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsPysparkJob(des, initial *WorkflowTemplateJobsPysparkJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPysparkJob {
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
	des.LoggingConfig = canonicalizeWorkflowTemplateJobsPysparkJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateJobsPysparkJob(c *Client, des, nw *WorkflowTemplateJobsPysparkJob) *WorkflowTemplateJobsPysparkJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MainPythonFileUri, nw.MainPythonFileUri) {
		nw.MainPythonFileUri = des.MainPythonFileUri
	}
	nw.LoggingConfig = canonicalizeNewWorkflowTemplateJobsPysparkJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPysparkJobSet(c *Client, des, nw []WorkflowTemplateJobsPysparkJob) []WorkflowTemplateJobsPysparkJob {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPysparkJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPysparkJob(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsPysparkJobSlice(c *Client, des, nw []WorkflowTemplateJobsPysparkJob) []WorkflowTemplateJobsPysparkJob {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsPysparkJob
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsPysparkJob(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsPysparkJobLoggingConfig(des, initial *WorkflowTemplateJobsPysparkJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPysparkJobLoggingConfig {
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

func canonicalizeNewWorkflowTemplateJobsPysparkJobLoggingConfig(c *Client, des, nw *WorkflowTemplateJobsPysparkJobLoggingConfig) *WorkflowTemplateJobsPysparkJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPysparkJobLoggingConfigSet(c *Client, des, nw []WorkflowTemplateJobsPysparkJobLoggingConfig) []WorkflowTemplateJobsPysparkJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPysparkJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPysparkJobLoggingConfig(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsPysparkJobLoggingConfigSlice(c *Client, des, nw []WorkflowTemplateJobsPysparkJobLoggingConfig) []WorkflowTemplateJobsPysparkJobLoggingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsPysparkJobLoggingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsPysparkJobLoggingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsHiveJob(des, initial *WorkflowTemplateJobsHiveJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHiveJob {
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
	des.QueryList = canonicalizeWorkflowTemplateJobsHiveJobQueryList(des.QueryList, initial.QueryList, opts...)
	if dcl.BoolCanonicalize(des.ContinueOnFailure, initial.ContinueOnFailure) || dcl.IsZeroValue(des.ContinueOnFailure) {
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

func canonicalizeNewWorkflowTemplateJobsHiveJob(c *Client, des, nw *WorkflowTemplateJobsHiveJob) *WorkflowTemplateJobsHiveJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.QueryFileUri, nw.QueryFileUri) {
		nw.QueryFileUri = des.QueryFileUri
	}
	nw.QueryList = canonicalizeNewWorkflowTemplateJobsHiveJobQueryList(c, des.QueryList, nw.QueryList)
	if dcl.BoolCanonicalize(des.ContinueOnFailure, nw.ContinueOnFailure) {
		nw.ContinueOnFailure = des.ContinueOnFailure
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsHiveJobSet(c *Client, des, nw []WorkflowTemplateJobsHiveJob) []WorkflowTemplateJobsHiveJob {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsHiveJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsHiveJob(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsHiveJobSlice(c *Client, des, nw []WorkflowTemplateJobsHiveJob) []WorkflowTemplateJobsHiveJob {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsHiveJob
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsHiveJob(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsHiveJobQueryList(des, initial *WorkflowTemplateJobsHiveJobQueryList, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHiveJobQueryList {
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

func canonicalizeNewWorkflowTemplateJobsHiveJobQueryList(c *Client, des, nw *WorkflowTemplateJobsHiveJobQueryList) *WorkflowTemplateJobsHiveJobQueryList {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsHiveJobQueryListSet(c *Client, des, nw []WorkflowTemplateJobsHiveJobQueryList) []WorkflowTemplateJobsHiveJobQueryList {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsHiveJobQueryList
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsHiveJobQueryList(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsHiveJobQueryListSlice(c *Client, des, nw []WorkflowTemplateJobsHiveJobQueryList) []WorkflowTemplateJobsHiveJobQueryList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsHiveJobQueryList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsHiveJobQueryList(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsPigJob(des, initial *WorkflowTemplateJobsPigJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPigJob {
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
	des.QueryList = canonicalizeWorkflowTemplateJobsPigJobQueryList(des.QueryList, initial.QueryList, opts...)
	if dcl.BoolCanonicalize(des.ContinueOnFailure, initial.ContinueOnFailure) || dcl.IsZeroValue(des.ContinueOnFailure) {
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
	des.LoggingConfig = canonicalizeWorkflowTemplateJobsPigJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateJobsPigJob(c *Client, des, nw *WorkflowTemplateJobsPigJob) *WorkflowTemplateJobsPigJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.QueryFileUri, nw.QueryFileUri) {
		nw.QueryFileUri = des.QueryFileUri
	}
	nw.QueryList = canonicalizeNewWorkflowTemplateJobsPigJobQueryList(c, des.QueryList, nw.QueryList)
	if dcl.BoolCanonicalize(des.ContinueOnFailure, nw.ContinueOnFailure) {
		nw.ContinueOnFailure = des.ContinueOnFailure
	}
	nw.LoggingConfig = canonicalizeNewWorkflowTemplateJobsPigJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPigJobSet(c *Client, des, nw []WorkflowTemplateJobsPigJob) []WorkflowTemplateJobsPigJob {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPigJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPigJob(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsPigJobSlice(c *Client, des, nw []WorkflowTemplateJobsPigJob) []WorkflowTemplateJobsPigJob {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsPigJob
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsPigJob(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsPigJobQueryList(des, initial *WorkflowTemplateJobsPigJobQueryList, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPigJobQueryList {
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

func canonicalizeNewWorkflowTemplateJobsPigJobQueryList(c *Client, des, nw *WorkflowTemplateJobsPigJobQueryList) *WorkflowTemplateJobsPigJobQueryList {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPigJobQueryListSet(c *Client, des, nw []WorkflowTemplateJobsPigJobQueryList) []WorkflowTemplateJobsPigJobQueryList {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPigJobQueryList
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPigJobQueryList(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsPigJobQueryListSlice(c *Client, des, nw []WorkflowTemplateJobsPigJobQueryList) []WorkflowTemplateJobsPigJobQueryList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsPigJobQueryList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsPigJobQueryList(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsPigJobLoggingConfig(des, initial *WorkflowTemplateJobsPigJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPigJobLoggingConfig {
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

func canonicalizeNewWorkflowTemplateJobsPigJobLoggingConfig(c *Client, des, nw *WorkflowTemplateJobsPigJobLoggingConfig) *WorkflowTemplateJobsPigJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPigJobLoggingConfigSet(c *Client, des, nw []WorkflowTemplateJobsPigJobLoggingConfig) []WorkflowTemplateJobsPigJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPigJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPigJobLoggingConfig(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsPigJobLoggingConfigSlice(c *Client, des, nw []WorkflowTemplateJobsPigJobLoggingConfig) []WorkflowTemplateJobsPigJobLoggingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsPigJobLoggingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsPigJobLoggingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsSparkRJob(des, initial *WorkflowTemplateJobsSparkRJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkRJob {
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
	des.LoggingConfig = canonicalizeWorkflowTemplateJobsSparkRJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateJobsSparkRJob(c *Client, des, nw *WorkflowTemplateJobsSparkRJob) *WorkflowTemplateJobsSparkRJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MainRFileUri, nw.MainRFileUri) {
		nw.MainRFileUri = des.MainRFileUri
	}
	nw.LoggingConfig = canonicalizeNewWorkflowTemplateJobsSparkRJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkRJobSet(c *Client, des, nw []WorkflowTemplateJobsSparkRJob) []WorkflowTemplateJobsSparkRJob {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkRJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkRJob(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsSparkRJobSlice(c *Client, des, nw []WorkflowTemplateJobsSparkRJob) []WorkflowTemplateJobsSparkRJob {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsSparkRJob
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsSparkRJob(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsSparkRJobLoggingConfig(des, initial *WorkflowTemplateJobsSparkRJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkRJobLoggingConfig {
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

func canonicalizeNewWorkflowTemplateJobsSparkRJobLoggingConfig(c *Client, des, nw *WorkflowTemplateJobsSparkRJobLoggingConfig) *WorkflowTemplateJobsSparkRJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkRJobLoggingConfigSet(c *Client, des, nw []WorkflowTemplateJobsSparkRJobLoggingConfig) []WorkflowTemplateJobsSparkRJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkRJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkRJobLoggingConfig(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsSparkRJobLoggingConfigSlice(c *Client, des, nw []WorkflowTemplateJobsSparkRJobLoggingConfig) []WorkflowTemplateJobsSparkRJobLoggingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsSparkRJobLoggingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsSparkRJobLoggingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsSparkSqlJob(des, initial *WorkflowTemplateJobsSparkSqlJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkSqlJob {
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
	des.QueryList = canonicalizeWorkflowTemplateJobsSparkSqlJobQueryList(des.QueryList, initial.QueryList, opts...)
	if dcl.IsZeroValue(des.ScriptVariables) {
		des.ScriptVariables = initial.ScriptVariables
	}
	if dcl.IsZeroValue(des.Properties) {
		des.Properties = initial.Properties
	}
	if dcl.IsZeroValue(des.JarFileUris) {
		des.JarFileUris = initial.JarFileUris
	}
	des.LoggingConfig = canonicalizeWorkflowTemplateJobsSparkSqlJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateJobsSparkSqlJob(c *Client, des, nw *WorkflowTemplateJobsSparkSqlJob) *WorkflowTemplateJobsSparkSqlJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.QueryFileUri, nw.QueryFileUri) {
		nw.QueryFileUri = des.QueryFileUri
	}
	nw.QueryList = canonicalizeNewWorkflowTemplateJobsSparkSqlJobQueryList(c, des.QueryList, nw.QueryList)
	nw.LoggingConfig = canonicalizeNewWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobSet(c *Client, des, nw []WorkflowTemplateJobsSparkSqlJob) []WorkflowTemplateJobsSparkSqlJob {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkSqlJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkSqlJob(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobSlice(c *Client, des, nw []WorkflowTemplateJobsSparkSqlJob) []WorkflowTemplateJobsSparkSqlJob {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsSparkSqlJob
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsSparkSqlJob(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsSparkSqlJobQueryList(des, initial *WorkflowTemplateJobsSparkSqlJobQueryList, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkSqlJobQueryList {
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

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobQueryList(c *Client, des, nw *WorkflowTemplateJobsSparkSqlJobQueryList) *WorkflowTemplateJobsSparkSqlJobQueryList {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobQueryListSet(c *Client, des, nw []WorkflowTemplateJobsSparkSqlJobQueryList) []WorkflowTemplateJobsSparkSqlJobQueryList {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkSqlJobQueryList
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkSqlJobQueryList(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobQueryListSlice(c *Client, des, nw []WorkflowTemplateJobsSparkSqlJobQueryList) []WorkflowTemplateJobsSparkSqlJobQueryList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsSparkSqlJobQueryList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsSparkSqlJobQueryList(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsSparkSqlJobLoggingConfig(des, initial *WorkflowTemplateJobsSparkSqlJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkSqlJobLoggingConfig {
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

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobLoggingConfig(c *Client, des, nw *WorkflowTemplateJobsSparkSqlJobLoggingConfig) *WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobLoggingConfigSet(c *Client, des, nw []WorkflowTemplateJobsSparkSqlJobLoggingConfig) []WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkSqlJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobLoggingConfigSlice(c *Client, des, nw []WorkflowTemplateJobsSparkSqlJobLoggingConfig) []WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsSparkSqlJobLoggingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsPrestoJob(des, initial *WorkflowTemplateJobsPrestoJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPrestoJob {
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
	des.QueryList = canonicalizeWorkflowTemplateJobsPrestoJobQueryList(des.QueryList, initial.QueryList, opts...)
	if dcl.BoolCanonicalize(des.ContinueOnFailure, initial.ContinueOnFailure) || dcl.IsZeroValue(des.ContinueOnFailure) {
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
	des.LoggingConfig = canonicalizeWorkflowTemplateJobsPrestoJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateJobsPrestoJob(c *Client, des, nw *WorkflowTemplateJobsPrestoJob) *WorkflowTemplateJobsPrestoJob {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.QueryFileUri, nw.QueryFileUri) {
		nw.QueryFileUri = des.QueryFileUri
	}
	nw.QueryList = canonicalizeNewWorkflowTemplateJobsPrestoJobQueryList(c, des.QueryList, nw.QueryList)
	if dcl.BoolCanonicalize(des.ContinueOnFailure, nw.ContinueOnFailure) {
		nw.ContinueOnFailure = des.ContinueOnFailure
	}
	if dcl.StringCanonicalize(des.OutputFormat, nw.OutputFormat) {
		nw.OutputFormat = des.OutputFormat
	}
	nw.LoggingConfig = canonicalizeNewWorkflowTemplateJobsPrestoJobLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPrestoJobSet(c *Client, des, nw []WorkflowTemplateJobsPrestoJob) []WorkflowTemplateJobsPrestoJob {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPrestoJob
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPrestoJob(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsPrestoJobSlice(c *Client, des, nw []WorkflowTemplateJobsPrestoJob) []WorkflowTemplateJobsPrestoJob {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsPrestoJob
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsPrestoJob(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsPrestoJobQueryList(des, initial *WorkflowTemplateJobsPrestoJobQueryList, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPrestoJobQueryList {
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

func canonicalizeNewWorkflowTemplateJobsPrestoJobQueryList(c *Client, des, nw *WorkflowTemplateJobsPrestoJobQueryList) *WorkflowTemplateJobsPrestoJobQueryList {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPrestoJobQueryListSet(c *Client, des, nw []WorkflowTemplateJobsPrestoJobQueryList) []WorkflowTemplateJobsPrestoJobQueryList {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPrestoJobQueryList
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPrestoJobQueryList(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsPrestoJobQueryListSlice(c *Client, des, nw []WorkflowTemplateJobsPrestoJobQueryList) []WorkflowTemplateJobsPrestoJobQueryList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsPrestoJobQueryList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsPrestoJobQueryList(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsPrestoJobLoggingConfig(des, initial *WorkflowTemplateJobsPrestoJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPrestoJobLoggingConfig {
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

func canonicalizeNewWorkflowTemplateJobsPrestoJobLoggingConfig(c *Client, des, nw *WorkflowTemplateJobsPrestoJobLoggingConfig) *WorkflowTemplateJobsPrestoJobLoggingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPrestoJobLoggingConfigSet(c *Client, des, nw []WorkflowTemplateJobsPrestoJobLoggingConfig) []WorkflowTemplateJobsPrestoJobLoggingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPrestoJobLoggingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPrestoJobLoggingConfig(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsPrestoJobLoggingConfigSlice(c *Client, des, nw []WorkflowTemplateJobsPrestoJobLoggingConfig) []WorkflowTemplateJobsPrestoJobLoggingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsPrestoJobLoggingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsPrestoJobLoggingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateJobsScheduling(des, initial *WorkflowTemplateJobsScheduling, opts ...dcl.ApplyOption) *WorkflowTemplateJobsScheduling {
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

func canonicalizeNewWorkflowTemplateJobsScheduling(c *Client, des, nw *WorkflowTemplateJobsScheduling) *WorkflowTemplateJobsScheduling {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSchedulingSet(c *Client, des, nw []WorkflowTemplateJobsScheduling) []WorkflowTemplateJobsScheduling {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsScheduling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsScheduling(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateJobsSchedulingSlice(c *Client, des, nw []WorkflowTemplateJobsScheduling) []WorkflowTemplateJobsScheduling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateJobsScheduling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateJobsScheduling(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateParameters(des, initial *WorkflowTemplateParameters, opts ...dcl.ApplyOption) *WorkflowTemplateParameters {
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
	if dcl.IsZeroValue(des.Fields) {
		des.Fields = initial.Fields
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	des.Validation = canonicalizeWorkflowTemplateParametersValidation(des.Validation, initial.Validation, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateParameters(c *Client, des, nw *WorkflowTemplateParameters) *WorkflowTemplateParameters {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	nw.Validation = canonicalizeNewWorkflowTemplateParametersValidation(c, des.Validation, nw.Validation)

	return nw
}

func canonicalizeNewWorkflowTemplateParametersSet(c *Client, des, nw []WorkflowTemplateParameters) []WorkflowTemplateParameters {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateParameters
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateParameters(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateParametersSlice(c *Client, des, nw []WorkflowTemplateParameters) []WorkflowTemplateParameters {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateParameters
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateParameters(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateParametersValidation(des, initial *WorkflowTemplateParametersValidation, opts ...dcl.ApplyOption) *WorkflowTemplateParametersValidation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Regex = canonicalizeWorkflowTemplateParametersValidationRegex(des.Regex, initial.Regex, opts...)
	des.Values = canonicalizeWorkflowTemplateParametersValidationValues(des.Values, initial.Values, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateParametersValidation(c *Client, des, nw *WorkflowTemplateParametersValidation) *WorkflowTemplateParametersValidation {
	if des == nil || nw == nil {
		return nw
	}

	nw.Regex = canonicalizeNewWorkflowTemplateParametersValidationRegex(c, des.Regex, nw.Regex)
	nw.Values = canonicalizeNewWorkflowTemplateParametersValidationValues(c, des.Values, nw.Values)

	return nw
}

func canonicalizeNewWorkflowTemplateParametersValidationSet(c *Client, des, nw []WorkflowTemplateParametersValidation) []WorkflowTemplateParametersValidation {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateParametersValidation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateParametersValidation(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateParametersValidationSlice(c *Client, des, nw []WorkflowTemplateParametersValidation) []WorkflowTemplateParametersValidation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateParametersValidation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateParametersValidation(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateParametersValidationRegex(des, initial *WorkflowTemplateParametersValidationRegex, opts ...dcl.ApplyOption) *WorkflowTemplateParametersValidationRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Regexes) {
		des.Regexes = initial.Regexes
	}

	return des
}

func canonicalizeNewWorkflowTemplateParametersValidationRegex(c *Client, des, nw *WorkflowTemplateParametersValidationRegex) *WorkflowTemplateParametersValidationRegex {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateParametersValidationRegexSet(c *Client, des, nw []WorkflowTemplateParametersValidationRegex) []WorkflowTemplateParametersValidationRegex {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateParametersValidationRegex
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateParametersValidationRegex(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateParametersValidationRegexSlice(c *Client, des, nw []WorkflowTemplateParametersValidationRegex) []WorkflowTemplateParametersValidationRegex {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateParametersValidationRegex
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateParametersValidationRegex(c, &d, &n))
	}

	return items
}

func canonicalizeWorkflowTemplateParametersValidationValues(des, initial *WorkflowTemplateParametersValidationValues, opts ...dcl.ApplyOption) *WorkflowTemplateParametersValidationValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Values) {
		des.Values = initial.Values
	}

	return des
}

func canonicalizeNewWorkflowTemplateParametersValidationValues(c *Client, des, nw *WorkflowTemplateParametersValidationValues) *WorkflowTemplateParametersValidationValues {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateParametersValidationValuesSet(c *Client, des, nw []WorkflowTemplateParametersValidationValues) []WorkflowTemplateParametersValidationValues {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateParametersValidationValues
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateParametersValidationValues(c, &d, &n) {
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

func canonicalizeNewWorkflowTemplateParametersValidationValuesSlice(c *Client, des, nw []WorkflowTemplateParametersValidationValues) []WorkflowTemplateParametersValidationValues {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []WorkflowTemplateParametersValidationValues
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkflowTemplateParametersValidationValues(c, &d, &n))
	}

	return items
}

type workflowTemplateDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         workflowTemplateApiOperation
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
func diffWorkflowTemplate(c *Client, desired, actual *WorkflowTemplate, opts ...dcl.ApplyOption) ([]workflowTemplateDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []workflowTemplateDiff
	if !reflect.DeepEqual(desired.Version, actual.Version) {
		c.Config.Logger.Infof("Detected diff in Version.\nDESIRED: %v\nACTUAL: %v", desired.Version, actual.Version)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Version",
		})
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Labels",
		})
	}
	if compareWorkflowTemplatePlacement(c, desired.Placement, actual.Placement) {
		c.Config.Logger.Infof("Detected diff in Placement.\nDESIRED: %v\nACTUAL: %v", desired.Placement, actual.Placement)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Placement",
		})
	}
	if compareWorkflowTemplateJobsSlice(c, desired.Jobs, actual.Jobs) {
		c.Config.Logger.Infof("Detected diff in Jobs.\nDESIRED: %v\nACTUAL: %v", desired.Jobs, actual.Jobs)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Jobs",
		})
	}
	if compareWorkflowTemplateParametersSlice(c, desired.Parameters, actual.Parameters) {
		c.Config.Logger.Infof("Detected diff in Parameters.\nDESIRED: %v\nACTUAL: %v", desired.Parameters, actual.Parameters)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Parameters",
		})
	}
	if !dcl.IsZeroValue(desired.DagTimeout) && !dcl.StringCanonicalize(desired.DagTimeout, actual.DagTimeout) {
		c.Config.Logger.Infof("Detected diff in DagTimeout.\nDESIRED: %v\nACTUAL: %v", desired.DagTimeout, actual.DagTimeout)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "DagTimeout",
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
	var deduped []workflowTemplateDiff
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
func compareWorkflowTemplatePlacement(c *Client, desired, actual *WorkflowTemplatePlacement) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ManagedCluster == nil && desired.ManagedCluster != nil && !dcl.IsEmptyValueIndirect(desired.ManagedCluster) {
		c.Config.Logger.Infof("desired ManagedCluster %s - but actually nil", dcl.SprintResource(desired.ManagedCluster))
		return true
	}
	if compareWorkflowTemplatePlacementManagedCluster(c, desired.ManagedCluster, actual.ManagedCluster) && !dcl.IsZeroValue(desired.ManagedCluster) {
		c.Config.Logger.Infof("Diff in ManagedCluster. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ManagedCluster), dcl.SprintResource(actual.ManagedCluster))
		return true
	}
	if actual.ClusterSelector == nil && desired.ClusterSelector != nil && !dcl.IsEmptyValueIndirect(desired.ClusterSelector) {
		c.Config.Logger.Infof("desired ClusterSelector %s - but actually nil", dcl.SprintResource(desired.ClusterSelector))
		return true
	}
	if compareWorkflowTemplatePlacementClusterSelector(c, desired.ClusterSelector, actual.ClusterSelector) && !dcl.IsZeroValue(desired.ClusterSelector) {
		c.Config.Logger.Infof("Diff in ClusterSelector. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterSelector), dcl.SprintResource(actual.ClusterSelector))
		return true
	}
	return false
}

func compareWorkflowTemplatePlacementSlice(c *Client, desired, actual []WorkflowTemplatePlacement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplatePlacement, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplatePlacement(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacement, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplatePlacementMap(c *Client, desired, actual map[string]WorkflowTemplatePlacement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplatePlacement, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacement, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplatePlacement(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacement, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplatePlacementManagedCluster(c *Client, desired, actual *WorkflowTemplatePlacementManagedCluster) bool {
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
	if actual.Config == nil && desired.Config != nil && !dcl.IsEmptyValueIndirect(desired.Config) {
		c.Config.Logger.Infof("desired Config %s - but actually nil", dcl.SprintResource(desired.Config))
		return true
	}
	if compareClusterClusterConfig(c, desired.Config, actual.Config) && !dcl.IsZeroValue(desired.Config) {
		c.Config.Logger.Infof("Diff in Config. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Config), dcl.SprintResource(actual.Config))
		return true
	}
	if actual.Labels == nil && desired.Labels != nil && !dcl.IsEmptyValueIndirect(desired.Labels) {
		c.Config.Logger.Infof("desired Labels %s - but actually nil", dcl.SprintResource(desired.Labels))
		return true
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	return false
}

func compareWorkflowTemplatePlacementManagedClusterSlice(c *Client, desired, actual []WorkflowTemplatePlacementManagedCluster) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplatePlacementManagedCluster, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplatePlacementManagedCluster(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacementManagedCluster, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplatePlacementManagedClusterMap(c *Client, desired, actual map[string]WorkflowTemplatePlacementManagedCluster) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplatePlacementManagedCluster, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacementManagedCluster, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplatePlacementManagedCluster(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacementManagedCluster, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplatePlacementClusterSelector(c *Client, desired, actual *WorkflowTemplatePlacementClusterSelector) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Zone == nil && desired.Zone != nil && !dcl.IsEmptyValueIndirect(desired.Zone) {
		c.Config.Logger.Infof("desired Zone %s - but actually nil", dcl.SprintResource(desired.Zone))
		return true
	}
	if !dcl.StringCanonicalize(desired.Zone, actual.Zone) && !dcl.IsZeroValue(desired.Zone) {
		c.Config.Logger.Infof("Diff in Zone. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Zone), dcl.SprintResource(actual.Zone))
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

func compareWorkflowTemplatePlacementClusterSelectorSlice(c *Client, desired, actual []WorkflowTemplatePlacementClusterSelector) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplatePlacementClusterSelector, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplatePlacementClusterSelector(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacementClusterSelector, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplatePlacementClusterSelectorMap(c *Client, desired, actual map[string]WorkflowTemplatePlacementClusterSelector) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplatePlacementClusterSelector, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacementClusterSelector, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplatePlacementClusterSelector(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacementClusterSelector, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobs(c *Client, desired, actual *WorkflowTemplateJobs) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.StepId == nil && desired.StepId != nil && !dcl.IsEmptyValueIndirect(desired.StepId) {
		c.Config.Logger.Infof("desired StepId %s - but actually nil", dcl.SprintResource(desired.StepId))
		return true
	}
	if !dcl.StringCanonicalize(desired.StepId, actual.StepId) && !dcl.IsZeroValue(desired.StepId) {
		c.Config.Logger.Infof("Diff in StepId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StepId), dcl.SprintResource(actual.StepId))
		return true
	}
	if actual.HadoopJob == nil && desired.HadoopJob != nil && !dcl.IsEmptyValueIndirect(desired.HadoopJob) {
		c.Config.Logger.Infof("desired HadoopJob %s - but actually nil", dcl.SprintResource(desired.HadoopJob))
		return true
	}
	if compareWorkflowTemplateJobsHadoopJob(c, desired.HadoopJob, actual.HadoopJob) && !dcl.IsZeroValue(desired.HadoopJob) {
		c.Config.Logger.Infof("Diff in HadoopJob. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HadoopJob), dcl.SprintResource(actual.HadoopJob))
		return true
	}
	if actual.SparkJob == nil && desired.SparkJob != nil && !dcl.IsEmptyValueIndirect(desired.SparkJob) {
		c.Config.Logger.Infof("desired SparkJob %s - but actually nil", dcl.SprintResource(desired.SparkJob))
		return true
	}
	if compareWorkflowTemplateJobsSparkJob(c, desired.SparkJob, actual.SparkJob) && !dcl.IsZeroValue(desired.SparkJob) {
		c.Config.Logger.Infof("Diff in SparkJob. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SparkJob), dcl.SprintResource(actual.SparkJob))
		return true
	}
	if actual.PysparkJob == nil && desired.PysparkJob != nil && !dcl.IsEmptyValueIndirect(desired.PysparkJob) {
		c.Config.Logger.Infof("desired PysparkJob %s - but actually nil", dcl.SprintResource(desired.PysparkJob))
		return true
	}
	if compareWorkflowTemplateJobsPysparkJob(c, desired.PysparkJob, actual.PysparkJob) && !dcl.IsZeroValue(desired.PysparkJob) {
		c.Config.Logger.Infof("Diff in PysparkJob. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PysparkJob), dcl.SprintResource(actual.PysparkJob))
		return true
	}
	if actual.HiveJob == nil && desired.HiveJob != nil && !dcl.IsEmptyValueIndirect(desired.HiveJob) {
		c.Config.Logger.Infof("desired HiveJob %s - but actually nil", dcl.SprintResource(desired.HiveJob))
		return true
	}
	if compareWorkflowTemplateJobsHiveJob(c, desired.HiveJob, actual.HiveJob) && !dcl.IsZeroValue(desired.HiveJob) {
		c.Config.Logger.Infof("Diff in HiveJob. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HiveJob), dcl.SprintResource(actual.HiveJob))
		return true
	}
	if actual.PigJob == nil && desired.PigJob != nil && !dcl.IsEmptyValueIndirect(desired.PigJob) {
		c.Config.Logger.Infof("desired PigJob %s - but actually nil", dcl.SprintResource(desired.PigJob))
		return true
	}
	if compareWorkflowTemplateJobsPigJob(c, desired.PigJob, actual.PigJob) && !dcl.IsZeroValue(desired.PigJob) {
		c.Config.Logger.Infof("Diff in PigJob. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PigJob), dcl.SprintResource(actual.PigJob))
		return true
	}
	if actual.SparkRJob == nil && desired.SparkRJob != nil && !dcl.IsEmptyValueIndirect(desired.SparkRJob) {
		c.Config.Logger.Infof("desired SparkRJob %s - but actually nil", dcl.SprintResource(desired.SparkRJob))
		return true
	}
	if compareWorkflowTemplateJobsSparkRJob(c, desired.SparkRJob, actual.SparkRJob) && !dcl.IsZeroValue(desired.SparkRJob) {
		c.Config.Logger.Infof("Diff in SparkRJob. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SparkRJob), dcl.SprintResource(actual.SparkRJob))
		return true
	}
	if actual.SparkSqlJob == nil && desired.SparkSqlJob != nil && !dcl.IsEmptyValueIndirect(desired.SparkSqlJob) {
		c.Config.Logger.Infof("desired SparkSqlJob %s - but actually nil", dcl.SprintResource(desired.SparkSqlJob))
		return true
	}
	if compareWorkflowTemplateJobsSparkSqlJob(c, desired.SparkSqlJob, actual.SparkSqlJob) && !dcl.IsZeroValue(desired.SparkSqlJob) {
		c.Config.Logger.Infof("Diff in SparkSqlJob. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SparkSqlJob), dcl.SprintResource(actual.SparkSqlJob))
		return true
	}
	if actual.PrestoJob == nil && desired.PrestoJob != nil && !dcl.IsEmptyValueIndirect(desired.PrestoJob) {
		c.Config.Logger.Infof("desired PrestoJob %s - but actually nil", dcl.SprintResource(desired.PrestoJob))
		return true
	}
	if compareWorkflowTemplateJobsPrestoJob(c, desired.PrestoJob, actual.PrestoJob) && !dcl.IsZeroValue(desired.PrestoJob) {
		c.Config.Logger.Infof("Diff in PrestoJob. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrestoJob), dcl.SprintResource(actual.PrestoJob))
		return true
	}
	if actual.Labels == nil && desired.Labels != nil && !dcl.IsEmptyValueIndirect(desired.Labels) {
		c.Config.Logger.Infof("desired Labels %s - but actually nil", dcl.SprintResource(desired.Labels))
		return true
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	if actual.Scheduling == nil && desired.Scheduling != nil && !dcl.IsEmptyValueIndirect(desired.Scheduling) {
		c.Config.Logger.Infof("desired Scheduling %s - but actually nil", dcl.SprintResource(desired.Scheduling))
		return true
	}
	if compareWorkflowTemplateJobsScheduling(c, desired.Scheduling, actual.Scheduling) && !dcl.IsZeroValue(desired.Scheduling) {
		c.Config.Logger.Infof("Diff in Scheduling. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scheduling), dcl.SprintResource(actual.Scheduling))
		return true
	}
	if actual.PrerequisiteStepIds == nil && desired.PrerequisiteStepIds != nil && !dcl.IsEmptyValueIndirect(desired.PrerequisiteStepIds) {
		c.Config.Logger.Infof("desired PrerequisiteStepIds %s - but actually nil", dcl.SprintResource(desired.PrerequisiteStepIds))
		return true
	}
	if !dcl.StringSliceEquals(desired.PrerequisiteStepIds, actual.PrerequisiteStepIds) && !dcl.IsZeroValue(desired.PrerequisiteStepIds) {
		c.Config.Logger.Infof("Diff in PrerequisiteStepIds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrerequisiteStepIds), dcl.SprintResource(actual.PrerequisiteStepIds))
		return true
	}
	return false
}

func compareWorkflowTemplateJobsSlice(c *Client, desired, actual []WorkflowTemplateJobs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobs, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobs(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobs, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsMap(c *Client, desired, actual map[string]WorkflowTemplateJobs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobs, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobs, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobs(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobs, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHadoopJob(c *Client, desired, actual *WorkflowTemplateJobsHadoopJob) bool {
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
	if compareWorkflowTemplateJobsHadoopJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareWorkflowTemplateJobsHadoopJobSlice(c *Client, desired, actual []WorkflowTemplateJobsHadoopJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHadoopJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsHadoopJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHadoopJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHadoopJobMap(c *Client, desired, actual map[string]WorkflowTemplateJobsHadoopJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHadoopJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHadoopJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsHadoopJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHadoopJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHadoopJobLoggingConfig(c *Client, desired, actual *WorkflowTemplateJobsHadoopJobLoggingConfig) bool {
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

func compareWorkflowTemplateJobsHadoopJobLoggingConfigSlice(c *Client, desired, actual []WorkflowTemplateJobsHadoopJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHadoopJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsHadoopJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHadoopJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHadoopJobLoggingConfigMap(c *Client, desired, actual map[string]WorkflowTemplateJobsHadoopJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHadoopJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHadoopJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsHadoopJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHadoopJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkJob(c *Client, desired, actual *WorkflowTemplateJobsSparkJob) bool {
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
	if compareWorkflowTemplateJobsSparkJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareWorkflowTemplateJobsSparkJobSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkJobMap(c *Client, desired, actual map[string]WorkflowTemplateJobsSparkJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsSparkJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkJobLoggingConfig(c *Client, desired, actual *WorkflowTemplateJobsSparkJobLoggingConfig) bool {
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

func compareWorkflowTemplateJobsSparkJobLoggingConfigSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkJobLoggingConfigMap(c *Client, desired, actual map[string]WorkflowTemplateJobsSparkJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsSparkJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPysparkJob(c *Client, desired, actual *WorkflowTemplateJobsPysparkJob) bool {
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
	if compareWorkflowTemplateJobsPysparkJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareWorkflowTemplateJobsPysparkJobSlice(c *Client, desired, actual []WorkflowTemplateJobsPysparkJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPysparkJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPysparkJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPysparkJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPysparkJobMap(c *Client, desired, actual map[string]WorkflowTemplateJobsPysparkJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPysparkJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPysparkJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsPysparkJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPysparkJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPysparkJobLoggingConfig(c *Client, desired, actual *WorkflowTemplateJobsPysparkJobLoggingConfig) bool {
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

func compareWorkflowTemplateJobsPysparkJobLoggingConfigSlice(c *Client, desired, actual []WorkflowTemplateJobsPysparkJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPysparkJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPysparkJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPysparkJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPysparkJobLoggingConfigMap(c *Client, desired, actual map[string]WorkflowTemplateJobsPysparkJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPysparkJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPysparkJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsPysparkJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPysparkJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHiveJob(c *Client, desired, actual *WorkflowTemplateJobsHiveJob) bool {
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
	if compareWorkflowTemplateJobsHiveJobQueryList(c, desired.QueryList, actual.QueryList) && !dcl.IsZeroValue(desired.QueryList) {
		c.Config.Logger.Infof("Diff in QueryList. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryList), dcl.SprintResource(actual.QueryList))
		return true
	}
	if actual.ContinueOnFailure == nil && desired.ContinueOnFailure != nil && !dcl.IsEmptyValueIndirect(desired.ContinueOnFailure) {
		c.Config.Logger.Infof("desired ContinueOnFailure %s - but actually nil", dcl.SprintResource(desired.ContinueOnFailure))
		return true
	}
	if !dcl.BoolCanonicalize(desired.ContinueOnFailure, actual.ContinueOnFailure) && !dcl.IsZeroValue(desired.ContinueOnFailure) {
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

func compareWorkflowTemplateJobsHiveJobSlice(c *Client, desired, actual []WorkflowTemplateJobsHiveJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHiveJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsHiveJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHiveJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHiveJobMap(c *Client, desired, actual map[string]WorkflowTemplateJobsHiveJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHiveJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHiveJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsHiveJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHiveJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHiveJobQueryList(c *Client, desired, actual *WorkflowTemplateJobsHiveJobQueryList) bool {
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

func compareWorkflowTemplateJobsHiveJobQueryListSlice(c *Client, desired, actual []WorkflowTemplateJobsHiveJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHiveJobQueryList, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsHiveJobQueryList(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHiveJobQueryList, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHiveJobQueryListMap(c *Client, desired, actual map[string]WorkflowTemplateJobsHiveJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHiveJobQueryList, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHiveJobQueryList, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsHiveJobQueryList(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHiveJobQueryList, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJob(c *Client, desired, actual *WorkflowTemplateJobsPigJob) bool {
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
	if compareWorkflowTemplateJobsPigJobQueryList(c, desired.QueryList, actual.QueryList) && !dcl.IsZeroValue(desired.QueryList) {
		c.Config.Logger.Infof("Diff in QueryList. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryList), dcl.SprintResource(actual.QueryList))
		return true
	}
	if actual.ContinueOnFailure == nil && desired.ContinueOnFailure != nil && !dcl.IsEmptyValueIndirect(desired.ContinueOnFailure) {
		c.Config.Logger.Infof("desired ContinueOnFailure %s - but actually nil", dcl.SprintResource(desired.ContinueOnFailure))
		return true
	}
	if !dcl.BoolCanonicalize(desired.ContinueOnFailure, actual.ContinueOnFailure) && !dcl.IsZeroValue(desired.ContinueOnFailure) {
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
	if compareWorkflowTemplateJobsPigJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareWorkflowTemplateJobsPigJobSlice(c *Client, desired, actual []WorkflowTemplateJobsPigJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPigJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJobMap(c *Client, desired, actual map[string]WorkflowTemplateJobsPigJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsPigJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJobQueryList(c *Client, desired, actual *WorkflowTemplateJobsPigJobQueryList) bool {
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

func compareWorkflowTemplateJobsPigJobQueryListSlice(c *Client, desired, actual []WorkflowTemplateJobsPigJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJobQueryList, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPigJobQueryList(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobQueryList, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJobQueryListMap(c *Client, desired, actual map[string]WorkflowTemplateJobsPigJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJobQueryList, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobQueryList, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsPigJobQueryList(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobQueryList, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJobLoggingConfig(c *Client, desired, actual *WorkflowTemplateJobsPigJobLoggingConfig) bool {
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

func compareWorkflowTemplateJobsPigJobLoggingConfigSlice(c *Client, desired, actual []WorkflowTemplateJobsPigJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPigJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJobLoggingConfigMap(c *Client, desired, actual map[string]WorkflowTemplateJobsPigJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsPigJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkRJob(c *Client, desired, actual *WorkflowTemplateJobsSparkRJob) bool {
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
	if compareWorkflowTemplateJobsSparkRJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareWorkflowTemplateJobsSparkRJobSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkRJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkRJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkRJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkRJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkRJobMap(c *Client, desired, actual map[string]WorkflowTemplateJobsSparkRJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkRJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkRJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsSparkRJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkRJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkRJobLoggingConfig(c *Client, desired, actual *WorkflowTemplateJobsSparkRJobLoggingConfig) bool {
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

func compareWorkflowTemplateJobsSparkRJobLoggingConfigSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkRJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkRJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkRJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkRJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkRJobLoggingConfigMap(c *Client, desired, actual map[string]WorkflowTemplateJobsSparkRJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkRJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkRJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsSparkRJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkRJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJob(c *Client, desired, actual *WorkflowTemplateJobsSparkSqlJob) bool {
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
	if compareWorkflowTemplateJobsSparkSqlJobQueryList(c, desired.QueryList, actual.QueryList) && !dcl.IsZeroValue(desired.QueryList) {
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
	if compareWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkSqlJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkSqlJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobMap(c *Client, desired, actual map[string]WorkflowTemplateJobsSparkSqlJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsSparkSqlJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobQueryList(c *Client, desired, actual *WorkflowTemplateJobsSparkSqlJobQueryList) bool {
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

func compareWorkflowTemplateJobsSparkSqlJobQueryListSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkSqlJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJobQueryList, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkSqlJobQueryList(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobQueryList, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobQueryListMap(c *Client, desired, actual map[string]WorkflowTemplateJobsSparkSqlJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJobQueryList, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobQueryList, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsSparkSqlJobQueryList(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobQueryList, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobLoggingConfig(c *Client, desired, actual *WorkflowTemplateJobsSparkSqlJobLoggingConfig) bool {
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

func compareWorkflowTemplateJobsSparkSqlJobLoggingConfigSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkSqlJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobLoggingConfigMap(c *Client, desired, actual map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJob(c *Client, desired, actual *WorkflowTemplateJobsPrestoJob) bool {
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
	if compareWorkflowTemplateJobsPrestoJobQueryList(c, desired.QueryList, actual.QueryList) && !dcl.IsZeroValue(desired.QueryList) {
		c.Config.Logger.Infof("Diff in QueryList. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryList), dcl.SprintResource(actual.QueryList))
		return true
	}
	if actual.ContinueOnFailure == nil && desired.ContinueOnFailure != nil && !dcl.IsEmptyValueIndirect(desired.ContinueOnFailure) {
		c.Config.Logger.Infof("desired ContinueOnFailure %s - but actually nil", dcl.SprintResource(desired.ContinueOnFailure))
		return true
	}
	if !dcl.BoolCanonicalize(desired.ContinueOnFailure, actual.ContinueOnFailure) && !dcl.IsZeroValue(desired.ContinueOnFailure) {
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
	if compareWorkflowTemplateJobsPrestoJobLoggingConfig(c, desired.LoggingConfig, actual.LoggingConfig) && !dcl.IsZeroValue(desired.LoggingConfig) {
		c.Config.Logger.Infof("Diff in LoggingConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoggingConfig), dcl.SprintResource(actual.LoggingConfig))
		return true
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJobSlice(c *Client, desired, actual []WorkflowTemplateJobsPrestoJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPrestoJob, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPrestoJob(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJob, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJobMap(c *Client, desired, actual map[string]WorkflowTemplateJobsPrestoJob) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPrestoJob, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJob, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsPrestoJob(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJob, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJobQueryList(c *Client, desired, actual *WorkflowTemplateJobsPrestoJobQueryList) bool {
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

func compareWorkflowTemplateJobsPrestoJobQueryListSlice(c *Client, desired, actual []WorkflowTemplateJobsPrestoJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPrestoJobQueryList, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPrestoJobQueryList(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJobQueryList, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJobQueryListMap(c *Client, desired, actual map[string]WorkflowTemplateJobsPrestoJobQueryList) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPrestoJobQueryList, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJobQueryList, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsPrestoJobQueryList(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJobQueryList, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJobLoggingConfig(c *Client, desired, actual *WorkflowTemplateJobsPrestoJobLoggingConfig) bool {
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

func compareWorkflowTemplateJobsPrestoJobLoggingConfigSlice(c *Client, desired, actual []WorkflowTemplateJobsPrestoJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPrestoJobLoggingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPrestoJobLoggingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJobLoggingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJobLoggingConfigMap(c *Client, desired, actual map[string]WorkflowTemplateJobsPrestoJobLoggingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPrestoJobLoggingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJobLoggingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsPrestoJobLoggingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJobLoggingConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsScheduling(c *Client, desired, actual *WorkflowTemplateJobsScheduling) bool {
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

func compareWorkflowTemplateJobsSchedulingSlice(c *Client, desired, actual []WorkflowTemplateJobsScheduling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsScheduling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsScheduling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsScheduling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSchedulingMap(c *Client, desired, actual map[string]WorkflowTemplateJobsScheduling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsScheduling, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsScheduling, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateJobsScheduling(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsScheduling, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateParameters(c *Client, desired, actual *WorkflowTemplateParameters) bool {
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
	if actual.Fields == nil && desired.Fields != nil && !dcl.IsEmptyValueIndirect(desired.Fields) {
		c.Config.Logger.Infof("desired Fields %s - but actually nil", dcl.SprintResource(desired.Fields))
		return true
	}
	if !dcl.StringSliceEquals(desired.Fields, actual.Fields) && !dcl.IsZeroValue(desired.Fields) {
		c.Config.Logger.Infof("Diff in Fields. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Fields), dcl.SprintResource(actual.Fields))
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
	if actual.Validation == nil && desired.Validation != nil && !dcl.IsEmptyValueIndirect(desired.Validation) {
		c.Config.Logger.Infof("desired Validation %s - but actually nil", dcl.SprintResource(desired.Validation))
		return true
	}
	if compareWorkflowTemplateParametersValidation(c, desired.Validation, actual.Validation) && !dcl.IsZeroValue(desired.Validation) {
		c.Config.Logger.Infof("Diff in Validation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Validation), dcl.SprintResource(actual.Validation))
		return true
	}
	return false
}

func compareWorkflowTemplateParametersSlice(c *Client, desired, actual []WorkflowTemplateParameters) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateParameters, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateParameters(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParameters, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateParametersMap(c *Client, desired, actual map[string]WorkflowTemplateParameters) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateParameters, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParameters, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateParameters(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParameters, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateParametersValidation(c *Client, desired, actual *WorkflowTemplateParametersValidation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Regex == nil && desired.Regex != nil && !dcl.IsEmptyValueIndirect(desired.Regex) {
		c.Config.Logger.Infof("desired Regex %s - but actually nil", dcl.SprintResource(desired.Regex))
		return true
	}
	if compareWorkflowTemplateParametersValidationRegex(c, desired.Regex, actual.Regex) && !dcl.IsZeroValue(desired.Regex) {
		c.Config.Logger.Infof("Diff in Regex. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Regex), dcl.SprintResource(actual.Regex))
		return true
	}
	if actual.Values == nil && desired.Values != nil && !dcl.IsEmptyValueIndirect(desired.Values) {
		c.Config.Logger.Infof("desired Values %s - but actually nil", dcl.SprintResource(desired.Values))
		return true
	}
	if compareWorkflowTemplateParametersValidationValues(c, desired.Values, actual.Values) && !dcl.IsZeroValue(desired.Values) {
		c.Config.Logger.Infof("Diff in Values. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Values), dcl.SprintResource(actual.Values))
		return true
	}
	return false
}

func compareWorkflowTemplateParametersValidationSlice(c *Client, desired, actual []WorkflowTemplateParametersValidation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateParametersValidation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateParametersValidation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParametersValidation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateParametersValidationMap(c *Client, desired, actual map[string]WorkflowTemplateParametersValidation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateParametersValidation, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParametersValidation, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateParametersValidation(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParametersValidation, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateParametersValidationRegex(c *Client, desired, actual *WorkflowTemplateParametersValidationRegex) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Regexes == nil && desired.Regexes != nil && !dcl.IsEmptyValueIndirect(desired.Regexes) {
		c.Config.Logger.Infof("desired Regexes %s - but actually nil", dcl.SprintResource(desired.Regexes))
		return true
	}
	if !dcl.StringSliceEquals(desired.Regexes, actual.Regexes) && !dcl.IsZeroValue(desired.Regexes) {
		c.Config.Logger.Infof("Diff in Regexes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Regexes), dcl.SprintResource(actual.Regexes))
		return true
	}
	return false
}

func compareWorkflowTemplateParametersValidationRegexSlice(c *Client, desired, actual []WorkflowTemplateParametersValidationRegex) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateParametersValidationRegex, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateParametersValidationRegex(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParametersValidationRegex, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateParametersValidationRegexMap(c *Client, desired, actual map[string]WorkflowTemplateParametersValidationRegex) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateParametersValidationRegex, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParametersValidationRegex, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateParametersValidationRegex(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParametersValidationRegex, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateParametersValidationValues(c *Client, desired, actual *WorkflowTemplateParametersValidationValues) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Values == nil && desired.Values != nil && !dcl.IsEmptyValueIndirect(desired.Values) {
		c.Config.Logger.Infof("desired Values %s - but actually nil", dcl.SprintResource(desired.Values))
		return true
	}
	if !dcl.StringSliceEquals(desired.Values, actual.Values) && !dcl.IsZeroValue(desired.Values) {
		c.Config.Logger.Infof("Diff in Values. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Values), dcl.SprintResource(actual.Values))
		return true
	}
	return false
}

func compareWorkflowTemplateParametersValidationValuesSlice(c *Client, desired, actual []WorkflowTemplateParametersValidationValues) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateParametersValidationValues, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateParametersValidationValues(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParametersValidationValues, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateParametersValidationValuesMap(c *Client, desired, actual map[string]WorkflowTemplateParametersValidationValues) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateParametersValidationValues, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParametersValidationValues, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareWorkflowTemplateParametersValidationValues(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateParametersValidationValues, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *WorkflowTemplate) urlNormalized() *WorkflowTemplate {
	normalized := deepcopy.Copy(*r).(WorkflowTemplate)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DagTimeout = dcl.SelfLinkToName(r.DagTimeout)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *WorkflowTemplate) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *WorkflowTemplate) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *WorkflowTemplate) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *WorkflowTemplate) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateWorkflowTemplate" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}", "https://dataproc.googleapis.com/v1beta2/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the WorkflowTemplate resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *WorkflowTemplate) marshal(c *Client) ([]byte, error) {
	m, err := expandWorkflowTemplate(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling WorkflowTemplate: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalWorkflowTemplate decodes JSON responses into the WorkflowTemplate resource schema.
func unmarshalWorkflowTemplate(b []byte, c *Client) (*WorkflowTemplate, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapWorkflowTemplate(m, c)
}

func unmarshalMapWorkflowTemplate(m map[string]interface{}, c *Client) (*WorkflowTemplate, error) {

	return flattenWorkflowTemplate(c, m), nil
}

// expandWorkflowTemplate expands WorkflowTemplate into a JSON request object.
func expandWorkflowTemplate(c *Client, f *WorkflowTemplate) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandWorkflowTemplatePlacement(c, f.Placement); err != nil {
		return nil, fmt.Errorf("error expanding Placement into placement: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["placement"] = v
	}
	if v, err := expandWorkflowTemplateJobsSlice(c, f.Jobs); err != nil {
		return nil, fmt.Errorf("error expanding Jobs into jobs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["jobs"] = v
	}
	if v, err := expandWorkflowTemplateParametersSlice(c, f.Parameters); err != nil {
		return nil, fmt.Errorf("error expanding Parameters into parameters: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parameters"] = v
	}
	if v := f.DagTimeout; !dcl.IsEmptyValueIndirect(v) {
		m["dagTimeout"] = v
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

// flattenWorkflowTemplate flattens WorkflowTemplate from a JSON request object into the
// WorkflowTemplate type.
func flattenWorkflowTemplate(c *Client, i interface{}) *WorkflowTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &WorkflowTemplate{}
	r.Name = dcl.FlattenString(m["name"])
	r.Version = dcl.FlattenInteger(m["version"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Placement = flattenWorkflowTemplatePlacement(c, m["placement"])
	r.Jobs = flattenWorkflowTemplateJobsSlice(c, m["jobs"])
	r.Parameters = flattenWorkflowTemplateParametersSlice(c, m["parameters"])
	r.DagTimeout = dcl.FlattenString(m["dagTimeout"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandWorkflowTemplatePlacementMap expands the contents of WorkflowTemplatePlacement into a JSON
// request object.
func expandWorkflowTemplatePlacementMap(c *Client, f map[string]WorkflowTemplatePlacement) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplatePlacement(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplatePlacementSlice expands the contents of WorkflowTemplatePlacement into a JSON
// request object.
func expandWorkflowTemplatePlacementSlice(c *Client, f []WorkflowTemplatePlacement) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplatePlacement(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplatePlacementMap flattens the contents of WorkflowTemplatePlacement from a JSON
// response object.
func flattenWorkflowTemplatePlacementMap(c *Client, i interface{}) map[string]WorkflowTemplatePlacement {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplatePlacement{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplatePlacement{}
	}

	items := make(map[string]WorkflowTemplatePlacement)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplatePlacement(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplatePlacementSlice flattens the contents of WorkflowTemplatePlacement from a JSON
// response object.
func flattenWorkflowTemplatePlacementSlice(c *Client, i interface{}) []WorkflowTemplatePlacement {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplatePlacement{}
	}

	if len(a) == 0 {
		return []WorkflowTemplatePlacement{}
	}

	items := make([]WorkflowTemplatePlacement, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplatePlacement(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplatePlacement expands an instance of WorkflowTemplatePlacement into a JSON
// request object.
func expandWorkflowTemplatePlacement(c *Client, f *WorkflowTemplatePlacement) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandWorkflowTemplatePlacementManagedCluster(c, f.ManagedCluster); err != nil {
		return nil, fmt.Errorf("error expanding ManagedCluster into managedCluster: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["managedCluster"] = v
	}
	if v, err := expandWorkflowTemplatePlacementClusterSelector(c, f.ClusterSelector); err != nil {
		return nil, fmt.Errorf("error expanding ClusterSelector into clusterSelector: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["clusterSelector"] = v
	}

	return m, nil
}

// flattenWorkflowTemplatePlacement flattens an instance of WorkflowTemplatePlacement from a JSON
// response object.
func flattenWorkflowTemplatePlacement(c *Client, i interface{}) *WorkflowTemplatePlacement {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplatePlacement{}
	r.ManagedCluster = flattenWorkflowTemplatePlacementManagedCluster(c, m["managedCluster"])
	r.ClusterSelector = flattenWorkflowTemplatePlacementClusterSelector(c, m["clusterSelector"])

	return r
}

// expandWorkflowTemplatePlacementManagedClusterMap expands the contents of WorkflowTemplatePlacementManagedCluster into a JSON
// request object.
func expandWorkflowTemplatePlacementManagedClusterMap(c *Client, f map[string]WorkflowTemplatePlacementManagedCluster) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplatePlacementManagedCluster(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplatePlacementManagedClusterSlice expands the contents of WorkflowTemplatePlacementManagedCluster into a JSON
// request object.
func expandWorkflowTemplatePlacementManagedClusterSlice(c *Client, f []WorkflowTemplatePlacementManagedCluster) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplatePlacementManagedCluster(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplatePlacementManagedClusterMap flattens the contents of WorkflowTemplatePlacementManagedCluster from a JSON
// response object.
func flattenWorkflowTemplatePlacementManagedClusterMap(c *Client, i interface{}) map[string]WorkflowTemplatePlacementManagedCluster {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplatePlacementManagedCluster{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplatePlacementManagedCluster{}
	}

	items := make(map[string]WorkflowTemplatePlacementManagedCluster)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplatePlacementManagedCluster(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplatePlacementManagedClusterSlice flattens the contents of WorkflowTemplatePlacementManagedCluster from a JSON
// response object.
func flattenWorkflowTemplatePlacementManagedClusterSlice(c *Client, i interface{}) []WorkflowTemplatePlacementManagedCluster {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplatePlacementManagedCluster{}
	}

	if len(a) == 0 {
		return []WorkflowTemplatePlacementManagedCluster{}
	}

	items := make([]WorkflowTemplatePlacementManagedCluster, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplatePlacementManagedCluster(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplatePlacementManagedCluster expands an instance of WorkflowTemplatePlacementManagedCluster into a JSON
// request object.
func expandWorkflowTemplatePlacementManagedCluster(c *Client, f *WorkflowTemplatePlacementManagedCluster) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ClusterName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterName"] = v
	}
	if v, err := expandClusterClusterConfig(c, f.Config); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["config"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}

	return m, nil
}

// flattenWorkflowTemplatePlacementManagedCluster flattens an instance of WorkflowTemplatePlacementManagedCluster from a JSON
// response object.
func flattenWorkflowTemplatePlacementManagedCluster(c *Client, i interface{}) *WorkflowTemplatePlacementManagedCluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplatePlacementManagedCluster{}
	r.ClusterName = dcl.FlattenString(m["clusterName"])
	r.Config = flattenClusterClusterConfig(c, m["config"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])

	return r
}

// expandWorkflowTemplatePlacementClusterSelectorMap expands the contents of WorkflowTemplatePlacementClusterSelector into a JSON
// request object.
func expandWorkflowTemplatePlacementClusterSelectorMap(c *Client, f map[string]WorkflowTemplatePlacementClusterSelector) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplatePlacementClusterSelector(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplatePlacementClusterSelectorSlice expands the contents of WorkflowTemplatePlacementClusterSelector into a JSON
// request object.
func expandWorkflowTemplatePlacementClusterSelectorSlice(c *Client, f []WorkflowTemplatePlacementClusterSelector) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplatePlacementClusterSelector(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplatePlacementClusterSelectorMap flattens the contents of WorkflowTemplatePlacementClusterSelector from a JSON
// response object.
func flattenWorkflowTemplatePlacementClusterSelectorMap(c *Client, i interface{}) map[string]WorkflowTemplatePlacementClusterSelector {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplatePlacementClusterSelector{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplatePlacementClusterSelector{}
	}

	items := make(map[string]WorkflowTemplatePlacementClusterSelector)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplatePlacementClusterSelector(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplatePlacementClusterSelectorSlice flattens the contents of WorkflowTemplatePlacementClusterSelector from a JSON
// response object.
func flattenWorkflowTemplatePlacementClusterSelectorSlice(c *Client, i interface{}) []WorkflowTemplatePlacementClusterSelector {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplatePlacementClusterSelector{}
	}

	if len(a) == 0 {
		return []WorkflowTemplatePlacementClusterSelector{}
	}

	items := make([]WorkflowTemplatePlacementClusterSelector, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplatePlacementClusterSelector(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplatePlacementClusterSelector expands an instance of WorkflowTemplatePlacementClusterSelector into a JSON
// request object.
func expandWorkflowTemplatePlacementClusterSelector(c *Client, f *WorkflowTemplatePlacementClusterSelector) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v := f.ClusterLabels; !dcl.IsEmptyValueIndirect(v) {
		m["clusterLabels"] = v
	}

	return m, nil
}

// flattenWorkflowTemplatePlacementClusterSelector flattens an instance of WorkflowTemplatePlacementClusterSelector from a JSON
// response object.
func flattenWorkflowTemplatePlacementClusterSelector(c *Client, i interface{}) *WorkflowTemplatePlacementClusterSelector {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplatePlacementClusterSelector{}
	r.Zone = dcl.FlattenString(m["zone"])
	r.ClusterLabels = dcl.FlattenKeyValuePairs(m["clusterLabels"])

	return r
}

// expandWorkflowTemplateJobsMap expands the contents of WorkflowTemplateJobs into a JSON
// request object.
func expandWorkflowTemplateJobsMap(c *Client, f map[string]WorkflowTemplateJobs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSlice expands the contents of WorkflowTemplateJobs into a JSON
// request object.
func expandWorkflowTemplateJobsSlice(c *Client, f []WorkflowTemplateJobs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsMap flattens the contents of WorkflowTemplateJobs from a JSON
// response object.
func flattenWorkflowTemplateJobsMap(c *Client, i interface{}) map[string]WorkflowTemplateJobs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobs{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobs{}
	}

	items := make(map[string]WorkflowTemplateJobs)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSlice flattens the contents of WorkflowTemplateJobs from a JSON
// response object.
func flattenWorkflowTemplateJobsSlice(c *Client, i interface{}) []WorkflowTemplateJobs {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobs{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobs{}
	}

	items := make([]WorkflowTemplateJobs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobs(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobs expands an instance of WorkflowTemplateJobs into a JSON
// request object.
func expandWorkflowTemplateJobs(c *Client, f *WorkflowTemplateJobs) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.StepId; !dcl.IsEmptyValueIndirect(v) {
		m["stepId"] = v
	}
	if v, err := expandWorkflowTemplateJobsHadoopJob(c, f.HadoopJob); err != nil {
		return nil, fmt.Errorf("error expanding HadoopJob into hadoopJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hadoopJob"] = v
	}
	if v, err := expandWorkflowTemplateJobsSparkJob(c, f.SparkJob); err != nil {
		return nil, fmt.Errorf("error expanding SparkJob into sparkJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sparkJob"] = v
	}
	if v, err := expandWorkflowTemplateJobsPysparkJob(c, f.PysparkJob); err != nil {
		return nil, fmt.Errorf("error expanding PysparkJob into pysparkJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pysparkJob"] = v
	}
	if v, err := expandWorkflowTemplateJobsHiveJob(c, f.HiveJob); err != nil {
		return nil, fmt.Errorf("error expanding HiveJob into hiveJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hiveJob"] = v
	}
	if v, err := expandWorkflowTemplateJobsPigJob(c, f.PigJob); err != nil {
		return nil, fmt.Errorf("error expanding PigJob into pigJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pigJob"] = v
	}
	if v, err := expandWorkflowTemplateJobsSparkRJob(c, f.SparkRJob); err != nil {
		return nil, fmt.Errorf("error expanding SparkRJob into sparkRJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sparkRJob"] = v
	}
	if v, err := expandWorkflowTemplateJobsSparkSqlJob(c, f.SparkSqlJob); err != nil {
		return nil, fmt.Errorf("error expanding SparkSqlJob into sparkSqlJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sparkSqlJob"] = v
	}
	if v, err := expandWorkflowTemplateJobsPrestoJob(c, f.PrestoJob); err != nil {
		return nil, fmt.Errorf("error expanding PrestoJob into prestoJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["prestoJob"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandWorkflowTemplateJobsScheduling(c, f.Scheduling); err != nil {
		return nil, fmt.Errorf("error expanding Scheduling into scheduling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scheduling"] = v
	}
	if v := f.PrerequisiteStepIds; !dcl.IsEmptyValueIndirect(v) {
		m["prerequisiteStepIds"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobs flattens an instance of WorkflowTemplateJobs from a JSON
// response object.
func flattenWorkflowTemplateJobs(c *Client, i interface{}) *WorkflowTemplateJobs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobs{}
	r.StepId = dcl.FlattenString(m["stepId"])
	r.HadoopJob = flattenWorkflowTemplateJobsHadoopJob(c, m["hadoopJob"])
	r.SparkJob = flattenWorkflowTemplateJobsSparkJob(c, m["sparkJob"])
	r.PysparkJob = flattenWorkflowTemplateJobsPysparkJob(c, m["pysparkJob"])
	r.HiveJob = flattenWorkflowTemplateJobsHiveJob(c, m["hiveJob"])
	r.PigJob = flattenWorkflowTemplateJobsPigJob(c, m["pigJob"])
	r.SparkRJob = flattenWorkflowTemplateJobsSparkRJob(c, m["sparkRJob"])
	r.SparkSqlJob = flattenWorkflowTemplateJobsSparkSqlJob(c, m["sparkSqlJob"])
	r.PrestoJob = flattenWorkflowTemplateJobsPrestoJob(c, m["prestoJob"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Scheduling = flattenWorkflowTemplateJobsScheduling(c, m["scheduling"])
	r.PrerequisiteStepIds = dcl.FlattenStringSlice(m["prerequisiteStepIds"])

	return r
}

// expandWorkflowTemplateJobsHadoopJobMap expands the contents of WorkflowTemplateJobsHadoopJob into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobMap(c *Client, f map[string]WorkflowTemplateJobsHadoopJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsHadoopJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsHadoopJobSlice expands the contents of WorkflowTemplateJobsHadoopJob into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobSlice(c *Client, f []WorkflowTemplateJobsHadoopJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsHadoopJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsHadoopJobMap flattens the contents of WorkflowTemplateJobsHadoopJob from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsHadoopJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsHadoopJob{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsHadoopJob{}
	}

	items := make(map[string]WorkflowTemplateJobsHadoopJob)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsHadoopJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsHadoopJobSlice flattens the contents of WorkflowTemplateJobsHadoopJob from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobSlice(c *Client, i interface{}) []WorkflowTemplateJobsHadoopJob {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsHadoopJob{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsHadoopJob{}
	}

	items := make([]WorkflowTemplateJobsHadoopJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsHadoopJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsHadoopJob expands an instance of WorkflowTemplateJobsHadoopJob into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJob(c *Client, f *WorkflowTemplateJobsHadoopJob) (map[string]interface{}, error) {
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
	if v, err := expandWorkflowTemplateJobsHadoopJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsHadoopJob flattens an instance of WorkflowTemplateJobsHadoopJob from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJob(c *Client, i interface{}) *WorkflowTemplateJobsHadoopJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsHadoopJob{}
	r.MainJarFileUri = dcl.FlattenString(m["mainJarFileUri"])
	r.MainClass = dcl.FlattenString(m["mainClass"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.FileUris = dcl.FlattenStringSlice(m["fileUris"])
	r.ArchiveUris = dcl.FlattenStringSlice(m["archiveUris"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenWorkflowTemplateJobsHadoopJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsHadoopJobLoggingConfigMap expands the contents of WorkflowTemplateJobsHadoopJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobLoggingConfigMap(c *Client, f map[string]WorkflowTemplateJobsHadoopJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsHadoopJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsHadoopJobLoggingConfigSlice expands the contents of WorkflowTemplateJobsHadoopJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobLoggingConfigSlice(c *Client, f []WorkflowTemplateJobsHadoopJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsHadoopJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsHadoopJobLoggingConfigMap flattens the contents of WorkflowTemplateJobsHadoopJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobLoggingConfigMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsHadoopJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsHadoopJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsHadoopJobLoggingConfig{}
	}

	items := make(map[string]WorkflowTemplateJobsHadoopJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsHadoopJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsHadoopJobLoggingConfigSlice flattens the contents of WorkflowTemplateJobsHadoopJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobLoggingConfigSlice(c *Client, i interface{}) []WorkflowTemplateJobsHadoopJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsHadoopJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsHadoopJobLoggingConfig{}
	}

	items := make([]WorkflowTemplateJobsHadoopJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsHadoopJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsHadoopJobLoggingConfig expands an instance of WorkflowTemplateJobsHadoopJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobLoggingConfig(c *Client, f *WorkflowTemplateJobsHadoopJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsHadoopJobLoggingConfig flattens an instance of WorkflowTemplateJobsHadoopJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobLoggingConfig(c *Client, i interface{}) *WorkflowTemplateJobsHadoopJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsHadoopJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsSparkJobMap expands the contents of WorkflowTemplateJobsSparkJob into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobMap(c *Client, f map[string]WorkflowTemplateJobsSparkJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkJobSlice expands the contents of WorkflowTemplateJobsSparkJob into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobSlice(c *Client, f []WorkflowTemplateJobsSparkJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkJobMap flattens the contents of WorkflowTemplateJobsSparkJob from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkJob{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkJob{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkJob)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkJobSlice flattens the contents of WorkflowTemplateJobsSparkJob from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkJob {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkJob{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkJob{}
	}

	items := make([]WorkflowTemplateJobsSparkJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkJob expands an instance of WorkflowTemplateJobsSparkJob into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJob(c *Client, f *WorkflowTemplateJobsSparkJob) (map[string]interface{}, error) {
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
	if v, err := expandWorkflowTemplateJobsSparkJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsSparkJob flattens an instance of WorkflowTemplateJobsSparkJob from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJob(c *Client, i interface{}) *WorkflowTemplateJobsSparkJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkJob{}
	r.MainJarFileUri = dcl.FlattenString(m["mainJarFileUri"])
	r.MainClass = dcl.FlattenString(m["mainClass"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.FileUris = dcl.FlattenStringSlice(m["fileUris"])
	r.ArchiveUris = dcl.FlattenStringSlice(m["archiveUris"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenWorkflowTemplateJobsSparkJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsSparkJobLoggingConfigMap expands the contents of WorkflowTemplateJobsSparkJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobLoggingConfigMap(c *Client, f map[string]WorkflowTemplateJobsSparkJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkJobLoggingConfigSlice expands the contents of WorkflowTemplateJobsSparkJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobLoggingConfigSlice(c *Client, f []WorkflowTemplateJobsSparkJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkJobLoggingConfigMap flattens the contents of WorkflowTemplateJobsSparkJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobLoggingConfigMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkJobLoggingConfig{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkJobLoggingConfigSlice flattens the contents of WorkflowTemplateJobsSparkJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobLoggingConfigSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkJobLoggingConfig{}
	}

	items := make([]WorkflowTemplateJobsSparkJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkJobLoggingConfig expands an instance of WorkflowTemplateJobsSparkJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobLoggingConfig(c *Client, f *WorkflowTemplateJobsSparkJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsSparkJobLoggingConfig flattens an instance of WorkflowTemplateJobsSparkJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobLoggingConfig(c *Client, i interface{}) *WorkflowTemplateJobsSparkJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsPysparkJobMap expands the contents of WorkflowTemplateJobsPysparkJob into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobMap(c *Client, f map[string]WorkflowTemplateJobsPysparkJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPysparkJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPysparkJobSlice expands the contents of WorkflowTemplateJobsPysparkJob into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobSlice(c *Client, f []WorkflowTemplateJobsPysparkJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPysparkJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPysparkJobMap flattens the contents of WorkflowTemplateJobsPysparkJob from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPysparkJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPysparkJob{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPysparkJob{}
	}

	items := make(map[string]WorkflowTemplateJobsPysparkJob)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPysparkJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPysparkJobSlice flattens the contents of WorkflowTemplateJobsPysparkJob from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobSlice(c *Client, i interface{}) []WorkflowTemplateJobsPysparkJob {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPysparkJob{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPysparkJob{}
	}

	items := make([]WorkflowTemplateJobsPysparkJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPysparkJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPysparkJob expands an instance of WorkflowTemplateJobsPysparkJob into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJob(c *Client, f *WorkflowTemplateJobsPysparkJob) (map[string]interface{}, error) {
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
	if v, err := expandWorkflowTemplateJobsPysparkJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsPysparkJob flattens an instance of WorkflowTemplateJobsPysparkJob from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJob(c *Client, i interface{}) *WorkflowTemplateJobsPysparkJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPysparkJob{}
	r.MainPythonFileUri = dcl.FlattenString(m["mainPythonFileUri"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.PythonFileUris = dcl.FlattenStringSlice(m["pythonFileUris"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.FileUris = dcl.FlattenStringSlice(m["fileUris"])
	r.ArchiveUris = dcl.FlattenStringSlice(m["archiveUris"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenWorkflowTemplateJobsPysparkJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsPysparkJobLoggingConfigMap expands the contents of WorkflowTemplateJobsPysparkJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobLoggingConfigMap(c *Client, f map[string]WorkflowTemplateJobsPysparkJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPysparkJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPysparkJobLoggingConfigSlice expands the contents of WorkflowTemplateJobsPysparkJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobLoggingConfigSlice(c *Client, f []WorkflowTemplateJobsPysparkJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPysparkJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPysparkJobLoggingConfigMap flattens the contents of WorkflowTemplateJobsPysparkJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobLoggingConfigMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPysparkJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPysparkJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPysparkJobLoggingConfig{}
	}

	items := make(map[string]WorkflowTemplateJobsPysparkJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPysparkJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPysparkJobLoggingConfigSlice flattens the contents of WorkflowTemplateJobsPysparkJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobLoggingConfigSlice(c *Client, i interface{}) []WorkflowTemplateJobsPysparkJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPysparkJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPysparkJobLoggingConfig{}
	}

	items := make([]WorkflowTemplateJobsPysparkJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPysparkJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPysparkJobLoggingConfig expands an instance of WorkflowTemplateJobsPysparkJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobLoggingConfig(c *Client, f *WorkflowTemplateJobsPysparkJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsPysparkJobLoggingConfig flattens an instance of WorkflowTemplateJobsPysparkJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobLoggingConfig(c *Client, i interface{}) *WorkflowTemplateJobsPysparkJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPysparkJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsHiveJobMap expands the contents of WorkflowTemplateJobsHiveJob into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobMap(c *Client, f map[string]WorkflowTemplateJobsHiveJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsHiveJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsHiveJobSlice expands the contents of WorkflowTemplateJobsHiveJob into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobSlice(c *Client, f []WorkflowTemplateJobsHiveJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsHiveJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsHiveJobMap flattens the contents of WorkflowTemplateJobsHiveJob from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsHiveJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsHiveJob{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsHiveJob{}
	}

	items := make(map[string]WorkflowTemplateJobsHiveJob)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsHiveJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsHiveJobSlice flattens the contents of WorkflowTemplateJobsHiveJob from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobSlice(c *Client, i interface{}) []WorkflowTemplateJobsHiveJob {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsHiveJob{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsHiveJob{}
	}

	items := make([]WorkflowTemplateJobsHiveJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsHiveJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsHiveJob expands an instance of WorkflowTemplateJobsHiveJob into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJob(c *Client, f *WorkflowTemplateJobsHiveJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["queryFileUri"] = v
	}
	if v, err := expandWorkflowTemplateJobsHiveJobQueryList(c, f.QueryList); err != nil {
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

// flattenWorkflowTemplateJobsHiveJob flattens an instance of WorkflowTemplateJobsHiveJob from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJob(c *Client, i interface{}) *WorkflowTemplateJobsHiveJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsHiveJob{}
	r.QueryFileUri = dcl.FlattenString(m["queryFileUri"])
	r.QueryList = flattenWorkflowTemplateJobsHiveJobQueryList(c, m["queryList"])
	r.ContinueOnFailure = dcl.FlattenBool(m["continueOnFailure"])
	r.ScriptVariables = dcl.FlattenKeyValuePairs(m["scriptVariables"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])

	return r
}

// expandWorkflowTemplateJobsHiveJobQueryListMap expands the contents of WorkflowTemplateJobsHiveJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobQueryListMap(c *Client, f map[string]WorkflowTemplateJobsHiveJobQueryList) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsHiveJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsHiveJobQueryListSlice expands the contents of WorkflowTemplateJobsHiveJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobQueryListSlice(c *Client, f []WorkflowTemplateJobsHiveJobQueryList) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsHiveJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsHiveJobQueryListMap flattens the contents of WorkflowTemplateJobsHiveJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobQueryListMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsHiveJobQueryList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsHiveJobQueryList{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsHiveJobQueryList{}
	}

	items := make(map[string]WorkflowTemplateJobsHiveJobQueryList)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsHiveJobQueryList(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsHiveJobQueryListSlice flattens the contents of WorkflowTemplateJobsHiveJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobQueryListSlice(c *Client, i interface{}) []WorkflowTemplateJobsHiveJobQueryList {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsHiveJobQueryList{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsHiveJobQueryList{}
	}

	items := make([]WorkflowTemplateJobsHiveJobQueryList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsHiveJobQueryList(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsHiveJobQueryList expands an instance of WorkflowTemplateJobsHiveJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobQueryList(c *Client, f *WorkflowTemplateJobsHiveJobQueryList) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Queries; !dcl.IsEmptyValueIndirect(v) {
		m["queries"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsHiveJobQueryList flattens an instance of WorkflowTemplateJobsHiveJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobQueryList(c *Client, i interface{}) *WorkflowTemplateJobsHiveJobQueryList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsHiveJobQueryList{}
	r.Queries = dcl.FlattenStringSlice(m["queries"])

	return r
}

// expandWorkflowTemplateJobsPigJobMap expands the contents of WorkflowTemplateJobsPigJob into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobMap(c *Client, f map[string]WorkflowTemplateJobsPigJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPigJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPigJobSlice expands the contents of WorkflowTemplateJobsPigJob into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobSlice(c *Client, f []WorkflowTemplateJobsPigJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPigJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPigJobMap flattens the contents of WorkflowTemplateJobsPigJob from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPigJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPigJob{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPigJob{}
	}

	items := make(map[string]WorkflowTemplateJobsPigJob)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPigJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPigJobSlice flattens the contents of WorkflowTemplateJobsPigJob from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobSlice(c *Client, i interface{}) []WorkflowTemplateJobsPigJob {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPigJob{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPigJob{}
	}

	items := make([]WorkflowTemplateJobsPigJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPigJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPigJob expands an instance of WorkflowTemplateJobsPigJob into a JSON
// request object.
func expandWorkflowTemplateJobsPigJob(c *Client, f *WorkflowTemplateJobsPigJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["queryFileUri"] = v
	}
	if v, err := expandWorkflowTemplateJobsPigJobQueryList(c, f.QueryList); err != nil {
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
	if v, err := expandWorkflowTemplateJobsPigJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsPigJob flattens an instance of WorkflowTemplateJobsPigJob from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJob(c *Client, i interface{}) *WorkflowTemplateJobsPigJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPigJob{}
	r.QueryFileUri = dcl.FlattenString(m["queryFileUri"])
	r.QueryList = flattenWorkflowTemplateJobsPigJobQueryList(c, m["queryList"])
	r.ContinueOnFailure = dcl.FlattenBool(m["continueOnFailure"])
	r.ScriptVariables = dcl.FlattenKeyValuePairs(m["scriptVariables"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.LoggingConfig = flattenWorkflowTemplateJobsPigJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsPigJobQueryListMap expands the contents of WorkflowTemplateJobsPigJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobQueryListMap(c *Client, f map[string]WorkflowTemplateJobsPigJobQueryList) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPigJobQueryListSlice expands the contents of WorkflowTemplateJobsPigJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobQueryListSlice(c *Client, f []WorkflowTemplateJobsPigJobQueryList) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPigJobQueryListMap flattens the contents of WorkflowTemplateJobsPigJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobQueryListMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPigJobQueryList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPigJobQueryList{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPigJobQueryList{}
	}

	items := make(map[string]WorkflowTemplateJobsPigJobQueryList)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPigJobQueryList(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPigJobQueryListSlice flattens the contents of WorkflowTemplateJobsPigJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobQueryListSlice(c *Client, i interface{}) []WorkflowTemplateJobsPigJobQueryList {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPigJobQueryList{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPigJobQueryList{}
	}

	items := make([]WorkflowTemplateJobsPigJobQueryList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPigJobQueryList(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPigJobQueryList expands an instance of WorkflowTemplateJobsPigJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobQueryList(c *Client, f *WorkflowTemplateJobsPigJobQueryList) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Queries; !dcl.IsEmptyValueIndirect(v) {
		m["queries"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsPigJobQueryList flattens an instance of WorkflowTemplateJobsPigJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobQueryList(c *Client, i interface{}) *WorkflowTemplateJobsPigJobQueryList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPigJobQueryList{}
	r.Queries = dcl.FlattenStringSlice(m["queries"])

	return r
}

// expandWorkflowTemplateJobsPigJobLoggingConfigMap expands the contents of WorkflowTemplateJobsPigJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobLoggingConfigMap(c *Client, f map[string]WorkflowTemplateJobsPigJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPigJobLoggingConfigSlice expands the contents of WorkflowTemplateJobsPigJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobLoggingConfigSlice(c *Client, f []WorkflowTemplateJobsPigJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPigJobLoggingConfigMap flattens the contents of WorkflowTemplateJobsPigJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobLoggingConfigMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPigJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPigJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPigJobLoggingConfig{}
	}

	items := make(map[string]WorkflowTemplateJobsPigJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPigJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPigJobLoggingConfigSlice flattens the contents of WorkflowTemplateJobsPigJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobLoggingConfigSlice(c *Client, i interface{}) []WorkflowTemplateJobsPigJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPigJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPigJobLoggingConfig{}
	}

	items := make([]WorkflowTemplateJobsPigJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPigJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPigJobLoggingConfig expands an instance of WorkflowTemplateJobsPigJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobLoggingConfig(c *Client, f *WorkflowTemplateJobsPigJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsPigJobLoggingConfig flattens an instance of WorkflowTemplateJobsPigJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobLoggingConfig(c *Client, i interface{}) *WorkflowTemplateJobsPigJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPigJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsSparkRJobMap expands the contents of WorkflowTemplateJobsSparkRJob into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobMap(c *Client, f map[string]WorkflowTemplateJobsSparkRJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkRJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkRJobSlice expands the contents of WorkflowTemplateJobsSparkRJob into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobSlice(c *Client, f []WorkflowTemplateJobsSparkRJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkRJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkRJobMap flattens the contents of WorkflowTemplateJobsSparkRJob from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkRJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkRJob{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkRJob{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkRJob)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkRJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkRJobSlice flattens the contents of WorkflowTemplateJobsSparkRJob from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkRJob {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkRJob{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkRJob{}
	}

	items := make([]WorkflowTemplateJobsSparkRJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkRJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkRJob expands an instance of WorkflowTemplateJobsSparkRJob into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJob(c *Client, f *WorkflowTemplateJobsSparkRJob) (map[string]interface{}, error) {
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
	if v, err := expandWorkflowTemplateJobsSparkRJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsSparkRJob flattens an instance of WorkflowTemplateJobsSparkRJob from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJob(c *Client, i interface{}) *WorkflowTemplateJobsSparkRJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkRJob{}
	r.MainRFileUri = dcl.FlattenString(m["mainRFileUri"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.FileUris = dcl.FlattenStringSlice(m["fileUris"])
	r.ArchiveUris = dcl.FlattenStringSlice(m["archiveUris"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenWorkflowTemplateJobsSparkRJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsSparkRJobLoggingConfigMap expands the contents of WorkflowTemplateJobsSparkRJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobLoggingConfigMap(c *Client, f map[string]WorkflowTemplateJobsSparkRJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkRJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkRJobLoggingConfigSlice expands the contents of WorkflowTemplateJobsSparkRJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobLoggingConfigSlice(c *Client, f []WorkflowTemplateJobsSparkRJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkRJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkRJobLoggingConfigMap flattens the contents of WorkflowTemplateJobsSparkRJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobLoggingConfigMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkRJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkRJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkRJobLoggingConfig{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkRJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkRJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkRJobLoggingConfigSlice flattens the contents of WorkflowTemplateJobsSparkRJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobLoggingConfigSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkRJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkRJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkRJobLoggingConfig{}
	}

	items := make([]WorkflowTemplateJobsSparkRJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkRJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkRJobLoggingConfig expands an instance of WorkflowTemplateJobsSparkRJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobLoggingConfig(c *Client, f *WorkflowTemplateJobsSparkRJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsSparkRJobLoggingConfig flattens an instance of WorkflowTemplateJobsSparkRJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobLoggingConfig(c *Client, i interface{}) *WorkflowTemplateJobsSparkRJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkRJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsSparkSqlJobMap expands the contents of WorkflowTemplateJobsSparkSqlJob into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobMap(c *Client, f map[string]WorkflowTemplateJobsSparkSqlJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkSqlJobSlice expands the contents of WorkflowTemplateJobsSparkSqlJob into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobSlice(c *Client, f []WorkflowTemplateJobsSparkSqlJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkSqlJobMap flattens the contents of WorkflowTemplateJobsSparkSqlJob from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkSqlJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkSqlJob{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkSqlJob{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkSqlJob)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkSqlJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkSqlJobSlice flattens the contents of WorkflowTemplateJobsSparkSqlJob from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkSqlJob {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkSqlJob{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkSqlJob{}
	}

	items := make([]WorkflowTemplateJobsSparkSqlJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkSqlJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkSqlJob expands an instance of WorkflowTemplateJobsSparkSqlJob into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJob(c *Client, f *WorkflowTemplateJobsSparkSqlJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["queryFileUri"] = v
	}
	if v, err := expandWorkflowTemplateJobsSparkSqlJobQueryList(c, f.QueryList); err != nil {
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
	if v, err := expandWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsSparkSqlJob flattens an instance of WorkflowTemplateJobsSparkSqlJob from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJob(c *Client, i interface{}) *WorkflowTemplateJobsSparkSqlJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkSqlJob{}
	r.QueryFileUri = dcl.FlattenString(m["queryFileUri"])
	r.QueryList = flattenWorkflowTemplateJobsSparkSqlJobQueryList(c, m["queryList"])
	r.ScriptVariables = dcl.FlattenKeyValuePairs(m["scriptVariables"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.JarFileUris = dcl.FlattenStringSlice(m["jarFileUris"])
	r.LoggingConfig = flattenWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsSparkSqlJobQueryListMap expands the contents of WorkflowTemplateJobsSparkSqlJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobQueryListMap(c *Client, f map[string]WorkflowTemplateJobsSparkSqlJobQueryList) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkSqlJobQueryListSlice expands the contents of WorkflowTemplateJobsSparkSqlJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobQueryListSlice(c *Client, f []WorkflowTemplateJobsSparkSqlJobQueryList) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkSqlJobQueryListMap flattens the contents of WorkflowTemplateJobsSparkSqlJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobQueryListMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkSqlJobQueryList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkSqlJobQueryList{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkSqlJobQueryList{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkSqlJobQueryList)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkSqlJobQueryList(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkSqlJobQueryListSlice flattens the contents of WorkflowTemplateJobsSparkSqlJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobQueryListSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkSqlJobQueryList {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkSqlJobQueryList{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkSqlJobQueryList{}
	}

	items := make([]WorkflowTemplateJobsSparkSqlJobQueryList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkSqlJobQueryList(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkSqlJobQueryList expands an instance of WorkflowTemplateJobsSparkSqlJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobQueryList(c *Client, f *WorkflowTemplateJobsSparkSqlJobQueryList) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Queries; !dcl.IsEmptyValueIndirect(v) {
		m["queries"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsSparkSqlJobQueryList flattens an instance of WorkflowTemplateJobsSparkSqlJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobQueryList(c *Client, i interface{}) *WorkflowTemplateJobsSparkSqlJobQueryList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkSqlJobQueryList{}
	r.Queries = dcl.FlattenStringSlice(m["queries"])

	return r
}

// expandWorkflowTemplateJobsSparkSqlJobLoggingConfigMap expands the contents of WorkflowTemplateJobsSparkSqlJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobLoggingConfigMap(c *Client, f map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkSqlJobLoggingConfigSlice expands the contents of WorkflowTemplateJobsSparkSqlJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobLoggingConfigSlice(c *Client, f []WorkflowTemplateJobsSparkSqlJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigMap flattens the contents of WorkflowTemplateJobsSparkSqlJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigSlice flattens the contents of WorkflowTemplateJobsSparkSqlJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	}

	items := make([]WorkflowTemplateJobsSparkSqlJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkSqlJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkSqlJobLoggingConfig expands an instance of WorkflowTemplateJobsSparkSqlJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobLoggingConfig(c *Client, f *WorkflowTemplateJobsSparkSqlJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsSparkSqlJobLoggingConfig flattens an instance of WorkflowTemplateJobsSparkSqlJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobLoggingConfig(c *Client, i interface{}) *WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsPrestoJobMap expands the contents of WorkflowTemplateJobsPrestoJob into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobMap(c *Client, f map[string]WorkflowTemplateJobsPrestoJob) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJob(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPrestoJobSlice expands the contents of WorkflowTemplateJobsPrestoJob into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobSlice(c *Client, f []WorkflowTemplateJobsPrestoJob) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJob(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPrestoJobMap flattens the contents of WorkflowTemplateJobsPrestoJob from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPrestoJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPrestoJob{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPrestoJob{}
	}

	items := make(map[string]WorkflowTemplateJobsPrestoJob)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPrestoJob(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPrestoJobSlice flattens the contents of WorkflowTemplateJobsPrestoJob from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobSlice(c *Client, i interface{}) []WorkflowTemplateJobsPrestoJob {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPrestoJob{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPrestoJob{}
	}

	items := make([]WorkflowTemplateJobsPrestoJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPrestoJob(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPrestoJob expands an instance of WorkflowTemplateJobsPrestoJob into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJob(c *Client, f *WorkflowTemplateJobsPrestoJob) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryFileUri; !dcl.IsEmptyValueIndirect(v) {
		m["queryFileUri"] = v
	}
	if v, err := expandWorkflowTemplateJobsPrestoJobQueryList(c, f.QueryList); err != nil {
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
	if v, err := expandWorkflowTemplateJobsPrestoJobLoggingConfig(c, f.LoggingConfig); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsPrestoJob flattens an instance of WorkflowTemplateJobsPrestoJob from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJob(c *Client, i interface{}) *WorkflowTemplateJobsPrestoJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPrestoJob{}
	r.QueryFileUri = dcl.FlattenString(m["queryFileUri"])
	r.QueryList = flattenWorkflowTemplateJobsPrestoJobQueryList(c, m["queryList"])
	r.ContinueOnFailure = dcl.FlattenBool(m["continueOnFailure"])
	r.OutputFormat = dcl.FlattenString(m["outputFormat"])
	r.ClientTags = dcl.FlattenStringSlice(m["clientTags"])
	r.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	r.LoggingConfig = flattenWorkflowTemplateJobsPrestoJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsPrestoJobQueryListMap expands the contents of WorkflowTemplateJobsPrestoJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobQueryListMap(c *Client, f map[string]WorkflowTemplateJobsPrestoJobQueryList) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPrestoJobQueryListSlice expands the contents of WorkflowTemplateJobsPrestoJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobQueryListSlice(c *Client, f []WorkflowTemplateJobsPrestoJobQueryList) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJobQueryList(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPrestoJobQueryListMap flattens the contents of WorkflowTemplateJobsPrestoJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobQueryListMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPrestoJobQueryList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPrestoJobQueryList{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPrestoJobQueryList{}
	}

	items := make(map[string]WorkflowTemplateJobsPrestoJobQueryList)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPrestoJobQueryList(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPrestoJobQueryListSlice flattens the contents of WorkflowTemplateJobsPrestoJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobQueryListSlice(c *Client, i interface{}) []WorkflowTemplateJobsPrestoJobQueryList {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPrestoJobQueryList{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPrestoJobQueryList{}
	}

	items := make([]WorkflowTemplateJobsPrestoJobQueryList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPrestoJobQueryList(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPrestoJobQueryList expands an instance of WorkflowTemplateJobsPrestoJobQueryList into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobQueryList(c *Client, f *WorkflowTemplateJobsPrestoJobQueryList) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Queries; !dcl.IsEmptyValueIndirect(v) {
		m["queries"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsPrestoJobQueryList flattens an instance of WorkflowTemplateJobsPrestoJobQueryList from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobQueryList(c *Client, i interface{}) *WorkflowTemplateJobsPrestoJobQueryList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPrestoJobQueryList{}
	r.Queries = dcl.FlattenStringSlice(m["queries"])

	return r
}

// expandWorkflowTemplateJobsPrestoJobLoggingConfigMap expands the contents of WorkflowTemplateJobsPrestoJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobLoggingConfigMap(c *Client, f map[string]WorkflowTemplateJobsPrestoJobLoggingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPrestoJobLoggingConfigSlice expands the contents of WorkflowTemplateJobsPrestoJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobLoggingConfigSlice(c *Client, f []WorkflowTemplateJobsPrestoJobLoggingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJobLoggingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPrestoJobLoggingConfigMap flattens the contents of WorkflowTemplateJobsPrestoJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobLoggingConfigMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPrestoJobLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPrestoJobLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPrestoJobLoggingConfig{}
	}

	items := make(map[string]WorkflowTemplateJobsPrestoJobLoggingConfig)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPrestoJobLoggingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPrestoJobLoggingConfigSlice flattens the contents of WorkflowTemplateJobsPrestoJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobLoggingConfigSlice(c *Client, i interface{}) []WorkflowTemplateJobsPrestoJobLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPrestoJobLoggingConfig{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPrestoJobLoggingConfig{}
	}

	items := make([]WorkflowTemplateJobsPrestoJobLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPrestoJobLoggingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPrestoJobLoggingConfig expands an instance of WorkflowTemplateJobsPrestoJobLoggingConfig into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobLoggingConfig(c *Client, f *WorkflowTemplateJobsPrestoJobLoggingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DriverLogLevels; !dcl.IsEmptyValueIndirect(v) {
		m["driverLogLevels"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateJobsPrestoJobLoggingConfig flattens an instance of WorkflowTemplateJobsPrestoJobLoggingConfig from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobLoggingConfig(c *Client, i interface{}) *WorkflowTemplateJobsPrestoJobLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPrestoJobLoggingConfig{}
	r.DriverLogLevels = dcl.FlattenKeyValuePairs(m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsSchedulingMap expands the contents of WorkflowTemplateJobsScheduling into a JSON
// request object.
func expandWorkflowTemplateJobsSchedulingMap(c *Client, f map[string]WorkflowTemplateJobsScheduling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsScheduling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSchedulingSlice expands the contents of WorkflowTemplateJobsScheduling into a JSON
// request object.
func expandWorkflowTemplateJobsSchedulingSlice(c *Client, f []WorkflowTemplateJobsScheduling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsScheduling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSchedulingMap flattens the contents of WorkflowTemplateJobsScheduling from a JSON
// response object.
func flattenWorkflowTemplateJobsSchedulingMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsScheduling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsScheduling{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsScheduling{}
	}

	items := make(map[string]WorkflowTemplateJobsScheduling)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsScheduling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSchedulingSlice flattens the contents of WorkflowTemplateJobsScheduling from a JSON
// response object.
func flattenWorkflowTemplateJobsSchedulingSlice(c *Client, i interface{}) []WorkflowTemplateJobsScheduling {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsScheduling{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsScheduling{}
	}

	items := make([]WorkflowTemplateJobsScheduling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsScheduling(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsScheduling expands an instance of WorkflowTemplateJobsScheduling into a JSON
// request object.
func expandWorkflowTemplateJobsScheduling(c *Client, f *WorkflowTemplateJobsScheduling) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsScheduling flattens an instance of WorkflowTemplateJobsScheduling from a JSON
// response object.
func flattenWorkflowTemplateJobsScheduling(c *Client, i interface{}) *WorkflowTemplateJobsScheduling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsScheduling{}
	r.MaxFailuresPerHour = dcl.FlattenInteger(m["maxFailuresPerHour"])
	r.MaxFailuresTotal = dcl.FlattenInteger(m["maxFailuresTotal"])

	return r
}

// expandWorkflowTemplateParametersMap expands the contents of WorkflowTemplateParameters into a JSON
// request object.
func expandWorkflowTemplateParametersMap(c *Client, f map[string]WorkflowTemplateParameters) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateParameters(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateParametersSlice expands the contents of WorkflowTemplateParameters into a JSON
// request object.
func expandWorkflowTemplateParametersSlice(c *Client, f []WorkflowTemplateParameters) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateParameters(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateParametersMap flattens the contents of WorkflowTemplateParameters from a JSON
// response object.
func flattenWorkflowTemplateParametersMap(c *Client, i interface{}) map[string]WorkflowTemplateParameters {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateParameters{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateParameters{}
	}

	items := make(map[string]WorkflowTemplateParameters)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateParameters(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateParametersSlice flattens the contents of WorkflowTemplateParameters from a JSON
// response object.
func flattenWorkflowTemplateParametersSlice(c *Client, i interface{}) []WorkflowTemplateParameters {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateParameters{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateParameters{}
	}

	items := make([]WorkflowTemplateParameters, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateParameters(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateParameters expands an instance of WorkflowTemplateParameters into a JSON
// request object.
func expandWorkflowTemplateParameters(c *Client, f *WorkflowTemplateParameters) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Fields; !dcl.IsEmptyValueIndirect(v) {
		m["fields"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandWorkflowTemplateParametersValidation(c, f.Validation); err != nil {
		return nil, fmt.Errorf("error expanding Validation into validation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["validation"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateParameters flattens an instance of WorkflowTemplateParameters from a JSON
// response object.
func flattenWorkflowTemplateParameters(c *Client, i interface{}) *WorkflowTemplateParameters {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateParameters{}
	r.Name = dcl.FlattenString(m["name"])
	r.Fields = dcl.FlattenStringSlice(m["fields"])
	r.Description = dcl.FlattenString(m["description"])
	r.Validation = flattenWorkflowTemplateParametersValidation(c, m["validation"])

	return r
}

// expandWorkflowTemplateParametersValidationMap expands the contents of WorkflowTemplateParametersValidation into a JSON
// request object.
func expandWorkflowTemplateParametersValidationMap(c *Client, f map[string]WorkflowTemplateParametersValidation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateParametersValidation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateParametersValidationSlice expands the contents of WorkflowTemplateParametersValidation into a JSON
// request object.
func expandWorkflowTemplateParametersValidationSlice(c *Client, f []WorkflowTemplateParametersValidation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateParametersValidation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateParametersValidationMap flattens the contents of WorkflowTemplateParametersValidation from a JSON
// response object.
func flattenWorkflowTemplateParametersValidationMap(c *Client, i interface{}) map[string]WorkflowTemplateParametersValidation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateParametersValidation{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateParametersValidation{}
	}

	items := make(map[string]WorkflowTemplateParametersValidation)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateParametersValidation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateParametersValidationSlice flattens the contents of WorkflowTemplateParametersValidation from a JSON
// response object.
func flattenWorkflowTemplateParametersValidationSlice(c *Client, i interface{}) []WorkflowTemplateParametersValidation {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateParametersValidation{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateParametersValidation{}
	}

	items := make([]WorkflowTemplateParametersValidation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateParametersValidation(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateParametersValidation expands an instance of WorkflowTemplateParametersValidation into a JSON
// request object.
func expandWorkflowTemplateParametersValidation(c *Client, f *WorkflowTemplateParametersValidation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandWorkflowTemplateParametersValidationRegex(c, f.Regex); err != nil {
		return nil, fmt.Errorf("error expanding Regex into regex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["regex"] = v
	}
	if v, err := expandWorkflowTemplateParametersValidationValues(c, f.Values); err != nil {
		return nil, fmt.Errorf("error expanding Values into values: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["values"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateParametersValidation flattens an instance of WorkflowTemplateParametersValidation from a JSON
// response object.
func flattenWorkflowTemplateParametersValidation(c *Client, i interface{}) *WorkflowTemplateParametersValidation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateParametersValidation{}
	r.Regex = flattenWorkflowTemplateParametersValidationRegex(c, m["regex"])
	r.Values = flattenWorkflowTemplateParametersValidationValues(c, m["values"])

	return r
}

// expandWorkflowTemplateParametersValidationRegexMap expands the contents of WorkflowTemplateParametersValidationRegex into a JSON
// request object.
func expandWorkflowTemplateParametersValidationRegexMap(c *Client, f map[string]WorkflowTemplateParametersValidationRegex) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateParametersValidationRegex(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateParametersValidationRegexSlice expands the contents of WorkflowTemplateParametersValidationRegex into a JSON
// request object.
func expandWorkflowTemplateParametersValidationRegexSlice(c *Client, f []WorkflowTemplateParametersValidationRegex) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateParametersValidationRegex(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateParametersValidationRegexMap flattens the contents of WorkflowTemplateParametersValidationRegex from a JSON
// response object.
func flattenWorkflowTemplateParametersValidationRegexMap(c *Client, i interface{}) map[string]WorkflowTemplateParametersValidationRegex {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateParametersValidationRegex{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateParametersValidationRegex{}
	}

	items := make(map[string]WorkflowTemplateParametersValidationRegex)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateParametersValidationRegex(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateParametersValidationRegexSlice flattens the contents of WorkflowTemplateParametersValidationRegex from a JSON
// response object.
func flattenWorkflowTemplateParametersValidationRegexSlice(c *Client, i interface{}) []WorkflowTemplateParametersValidationRegex {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateParametersValidationRegex{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateParametersValidationRegex{}
	}

	items := make([]WorkflowTemplateParametersValidationRegex, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateParametersValidationRegex(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateParametersValidationRegex expands an instance of WorkflowTemplateParametersValidationRegex into a JSON
// request object.
func expandWorkflowTemplateParametersValidationRegex(c *Client, f *WorkflowTemplateParametersValidationRegex) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Regexes; !dcl.IsEmptyValueIndirect(v) {
		m["regexes"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateParametersValidationRegex flattens an instance of WorkflowTemplateParametersValidationRegex from a JSON
// response object.
func flattenWorkflowTemplateParametersValidationRegex(c *Client, i interface{}) *WorkflowTemplateParametersValidationRegex {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateParametersValidationRegex{}
	r.Regexes = dcl.FlattenStringSlice(m["regexes"])

	return r
}

// expandWorkflowTemplateParametersValidationValuesMap expands the contents of WorkflowTemplateParametersValidationValues into a JSON
// request object.
func expandWorkflowTemplateParametersValidationValuesMap(c *Client, f map[string]WorkflowTemplateParametersValidationValues) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateParametersValidationValues(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateParametersValidationValuesSlice expands the contents of WorkflowTemplateParametersValidationValues into a JSON
// request object.
func expandWorkflowTemplateParametersValidationValuesSlice(c *Client, f []WorkflowTemplateParametersValidationValues) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateParametersValidationValues(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateParametersValidationValuesMap flattens the contents of WorkflowTemplateParametersValidationValues from a JSON
// response object.
func flattenWorkflowTemplateParametersValidationValuesMap(c *Client, i interface{}) map[string]WorkflowTemplateParametersValidationValues {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateParametersValidationValues{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateParametersValidationValues{}
	}

	items := make(map[string]WorkflowTemplateParametersValidationValues)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateParametersValidationValues(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateParametersValidationValuesSlice flattens the contents of WorkflowTemplateParametersValidationValues from a JSON
// response object.
func flattenWorkflowTemplateParametersValidationValuesSlice(c *Client, i interface{}) []WorkflowTemplateParametersValidationValues {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateParametersValidationValues{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateParametersValidationValues{}
	}

	items := make([]WorkflowTemplateParametersValidationValues, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateParametersValidationValues(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateParametersValidationValues expands an instance of WorkflowTemplateParametersValidationValues into a JSON
// request object.
func expandWorkflowTemplateParametersValidationValues(c *Client, f *WorkflowTemplateParametersValidationValues) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Values; !dcl.IsEmptyValueIndirect(v) {
		m["values"] = v
	}

	return m, nil
}

// flattenWorkflowTemplateParametersValidationValues flattens an instance of WorkflowTemplateParametersValidationValues from a JSON
// response object.
func flattenWorkflowTemplateParametersValidationValues(c *Client, i interface{}) *WorkflowTemplateParametersValidationValues {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateParametersValidationValues{}
	r.Values = dcl.FlattenStringSlice(m["values"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *WorkflowTemplate) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalWorkflowTemplate(b, c)
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
