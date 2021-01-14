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
package dataproc

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
func (r *WorkflowTemplateLabels) validate() error {
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
func (r *WorkflowTemplatePlacementManagedClusterLabels) validate() error {
	return nil
}
func (r *WorkflowTemplatePlacementClusterSelector) validate() error {
	if err := dcl.Required(r, "clusterLabels"); err != nil {
		return err
	}
	return nil
}
func (r *WorkflowTemplatePlacementClusterSelectorClusterLabels) validate() error {
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
func (r *WorkflowTemplateJobsHadoopJobProperties) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsHadoopJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) validate() error {
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
func (r *WorkflowTemplateJobsSparkJobProperties) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) validate() error {
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
func (r *WorkflowTemplateJobsPysparkJobProperties) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsPysparkJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) validate() error {
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
func (r *WorkflowTemplateJobsHiveJobScriptVariables) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsHiveJobProperties) validate() error {
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
func (r *WorkflowTemplateJobsPigJobScriptVariables) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsPigJobProperties) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsPigJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) validate() error {
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
func (r *WorkflowTemplateJobsSparkRJobProperties) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkRJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) validate() error {
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
func (r *WorkflowTemplateJobsSparkSqlJobScriptVariables) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkSqlJobProperties) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkSqlJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) validate() error {
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
func (r *WorkflowTemplateJobsPrestoJobProperties) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsPrestoJobLoggingConfig) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) validate() error {
	return nil
}
func (r *WorkflowTemplateJobsLabels) validate() error {
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
	return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil
}

func workflowTemplateListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil

}

func workflowTemplateCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil

}

func workflowTemplateDeleteURL(userBasePath string, r *WorkflowTemplate) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}", "https://dataproc.googleapis.com/v1/", userBasePath, params), nil
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
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listWorkflowTemplateOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
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
	for _, v := range m.Items {
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
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return fmt.Errorf("failed to delete WorkflowTemplate: %w", err)
	}
	_, err = c.GetWorkflowTemplate(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createWorkflowTemplateOperation struct{}

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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	_ = o // We might not use resp- this will stop Go complaining

	if _, err := c.GetWorkflowTemplate(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getWorkflowTemplateRaw(ctx context.Context, r *WorkflowTemplate) ([]byte, error) {

	u, err := workflowTemplateGetURL(c.Config.BasePath, r.urlNormalized())
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*WorkflowTemplate); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected WorkflowTemplate, got %T", sh)
		} else {
			_ = r
		}
	}

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
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeWorkflowTemplateNewState(c *Client, rawNew, rawDesired *WorkflowTemplate) (*WorkflowTemplate, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.NameToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

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
	}

	if dcl.IsEmptyValueIndirect(rawNew.Parameters) && dcl.IsEmptyValueIndirect(rawDesired.Parameters) {
		rawNew.Parameters = rawDesired.Parameters
	} else {
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

func canonicalizeWorkflowTemplateLabels(des, initial *WorkflowTemplateLabels, opts ...dcl.ApplyOption) *WorkflowTemplateLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateLabels(c *Client, des, nw *WorkflowTemplateLabels) *WorkflowTemplateLabels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateLabelsSet(c *Client, des, nw []WorkflowTemplateLabels) []WorkflowTemplateLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateLabels(c, &d, &n) {
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

func canonicalizeWorkflowTemplatePlacement(des, initial *WorkflowTemplatePlacement, opts ...dcl.ApplyOption) *WorkflowTemplatePlacement {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplatePlacementManagedCluster(des, initial *WorkflowTemplatePlacementManagedCluster, opts ...dcl.ApplyOption) *WorkflowTemplatePlacementManagedCluster {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ClusterName) {
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

func canonicalizeWorkflowTemplatePlacementManagedClusterLabels(des, initial *WorkflowTemplatePlacementManagedClusterLabels, opts ...dcl.ApplyOption) *WorkflowTemplatePlacementManagedClusterLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplatePlacementManagedClusterLabels(c *Client, des, nw *WorkflowTemplatePlacementManagedClusterLabels) *WorkflowTemplatePlacementManagedClusterLabels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplatePlacementManagedClusterLabelsSet(c *Client, des, nw []WorkflowTemplatePlacementManagedClusterLabels) []WorkflowTemplatePlacementManagedClusterLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplatePlacementManagedClusterLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplatePlacementManagedClusterLabels(c, &d, &n) {
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

func canonicalizeWorkflowTemplatePlacementClusterSelector(des, initial *WorkflowTemplatePlacementClusterSelector, opts ...dcl.ApplyOption) *WorkflowTemplatePlacementClusterSelector {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Zone) {
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

func canonicalizeWorkflowTemplatePlacementClusterSelectorClusterLabels(des, initial *WorkflowTemplatePlacementClusterSelectorClusterLabels, opts ...dcl.ApplyOption) *WorkflowTemplatePlacementClusterSelectorClusterLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplatePlacementClusterSelectorClusterLabels(c *Client, des, nw *WorkflowTemplatePlacementClusterSelectorClusterLabels) *WorkflowTemplatePlacementClusterSelectorClusterLabels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplatePlacementClusterSelectorClusterLabelsSet(c *Client, des, nw []WorkflowTemplatePlacementClusterSelectorClusterLabels) []WorkflowTemplatePlacementClusterSelectorClusterLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplatePlacementClusterSelectorClusterLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplatePlacementClusterSelectorClusterLabels(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobs(des, initial *WorkflowTemplateJobs, opts ...dcl.ApplyOption) *WorkflowTemplateJobs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.StepId) {
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

func canonicalizeWorkflowTemplateJobsHadoopJob(des, initial *WorkflowTemplateJobsHadoopJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHadoopJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MainJarFileUri) {
		des.MainJarFileUri = initial.MainJarFileUri
	}
	if dcl.IsZeroValue(des.MainClass) {
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

func canonicalizeWorkflowTemplateJobsHadoopJobProperties(des, initial *WorkflowTemplateJobsHadoopJobProperties, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHadoopJobProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsHadoopJobProperties(c *Client, des, nw *WorkflowTemplateJobsHadoopJobProperties) *WorkflowTemplateJobsHadoopJobProperties {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsHadoopJobPropertiesSet(c *Client, des, nw []WorkflowTemplateJobsHadoopJobProperties) []WorkflowTemplateJobsHadoopJobProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsHadoopJobProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsHadoopJobProperties(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsHadoopJobLoggingConfig(des, initial *WorkflowTemplateJobsHadoopJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHadoopJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(des, initial *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c *Client, des, nw *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsSet(c *Client, des, nw []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsSparkJob(des, initial *WorkflowTemplateJobsSparkJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MainJarFileUri) {
		des.MainJarFileUri = initial.MainJarFileUri
	}
	if dcl.IsZeroValue(des.MainClass) {
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

func canonicalizeWorkflowTemplateJobsSparkJobProperties(des, initial *WorkflowTemplateJobsSparkJobProperties, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkJobProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsSparkJobProperties(c *Client, des, nw *WorkflowTemplateJobsSparkJobProperties) *WorkflowTemplateJobsSparkJobProperties {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkJobPropertiesSet(c *Client, des, nw []WorkflowTemplateJobsSparkJobProperties) []WorkflowTemplateJobsSparkJobProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkJobProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkJobProperties(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsSparkJobLoggingConfig(des, initial *WorkflowTemplateJobsSparkJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(des, initial *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c *Client, des, nw *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsSet(c *Client, des, nw []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsPysparkJob(des, initial *WorkflowTemplateJobsPysparkJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPysparkJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MainPythonFileUri) {
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

func canonicalizeWorkflowTemplateJobsPysparkJobProperties(des, initial *WorkflowTemplateJobsPysparkJobProperties, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPysparkJobProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsPysparkJobProperties(c *Client, des, nw *WorkflowTemplateJobsPysparkJobProperties) *WorkflowTemplateJobsPysparkJobProperties {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPysparkJobPropertiesSet(c *Client, des, nw []WorkflowTemplateJobsPysparkJobProperties) []WorkflowTemplateJobsPysparkJobProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPysparkJobProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPysparkJobProperties(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsPysparkJobLoggingConfig(des, initial *WorkflowTemplateJobsPysparkJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPysparkJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(des, initial *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c *Client, des, nw *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsSet(c *Client, des, nw []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsHiveJob(des, initial *WorkflowTemplateJobsHiveJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHiveJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.QueryFileUri) {
		des.QueryFileUri = initial.QueryFileUri
	}
	des.QueryList = canonicalizeWorkflowTemplateJobsHiveJobQueryList(des.QueryList, initial.QueryList, opts...)
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

func canonicalizeNewWorkflowTemplateJobsHiveJob(c *Client, des, nw *WorkflowTemplateJobsHiveJob) *WorkflowTemplateJobsHiveJob {
	if des == nil || nw == nil {
		return nw
	}

	nw.QueryList = canonicalizeNewWorkflowTemplateJobsHiveJobQueryList(c, des.QueryList, nw.QueryList)

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

func canonicalizeWorkflowTemplateJobsHiveJobQueryList(des, initial *WorkflowTemplateJobsHiveJobQueryList, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHiveJobQueryList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsHiveJobScriptVariables(des, initial *WorkflowTemplateJobsHiveJobScriptVariables, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHiveJobScriptVariables {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsHiveJobScriptVariables(c *Client, des, nw *WorkflowTemplateJobsHiveJobScriptVariables) *WorkflowTemplateJobsHiveJobScriptVariables {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsHiveJobScriptVariablesSet(c *Client, des, nw []WorkflowTemplateJobsHiveJobScriptVariables) []WorkflowTemplateJobsHiveJobScriptVariables {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsHiveJobScriptVariables
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsHiveJobScriptVariables(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsHiveJobProperties(des, initial *WorkflowTemplateJobsHiveJobProperties, opts ...dcl.ApplyOption) *WorkflowTemplateJobsHiveJobProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsHiveJobProperties(c *Client, des, nw *WorkflowTemplateJobsHiveJobProperties) *WorkflowTemplateJobsHiveJobProperties {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsHiveJobPropertiesSet(c *Client, des, nw []WorkflowTemplateJobsHiveJobProperties) []WorkflowTemplateJobsHiveJobProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsHiveJobProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsHiveJobProperties(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsPigJob(des, initial *WorkflowTemplateJobsPigJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPigJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.QueryFileUri) {
		des.QueryFileUri = initial.QueryFileUri
	}
	des.QueryList = canonicalizeWorkflowTemplateJobsPigJobQueryList(des.QueryList, initial.QueryList, opts...)
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
	des.LoggingConfig = canonicalizeWorkflowTemplateJobsPigJobLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateJobsPigJob(c *Client, des, nw *WorkflowTemplateJobsPigJob) *WorkflowTemplateJobsPigJob {
	if des == nil || nw == nil {
		return nw
	}

	nw.QueryList = canonicalizeNewWorkflowTemplateJobsPigJobQueryList(c, des.QueryList, nw.QueryList)
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

func canonicalizeWorkflowTemplateJobsPigJobQueryList(des, initial *WorkflowTemplateJobsPigJobQueryList, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPigJobQueryList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsPigJobScriptVariables(des, initial *WorkflowTemplateJobsPigJobScriptVariables, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPigJobScriptVariables {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsPigJobScriptVariables(c *Client, des, nw *WorkflowTemplateJobsPigJobScriptVariables) *WorkflowTemplateJobsPigJobScriptVariables {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPigJobScriptVariablesSet(c *Client, des, nw []WorkflowTemplateJobsPigJobScriptVariables) []WorkflowTemplateJobsPigJobScriptVariables {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPigJobScriptVariables
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPigJobScriptVariables(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsPigJobProperties(des, initial *WorkflowTemplateJobsPigJobProperties, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPigJobProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsPigJobProperties(c *Client, des, nw *WorkflowTemplateJobsPigJobProperties) *WorkflowTemplateJobsPigJobProperties {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPigJobPropertiesSet(c *Client, des, nw []WorkflowTemplateJobsPigJobProperties) []WorkflowTemplateJobsPigJobProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPigJobProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPigJobProperties(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsPigJobLoggingConfig(des, initial *WorkflowTemplateJobsPigJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPigJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(des, initial *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c *Client, des, nw *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsSet(c *Client, des, nw []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsSparkRJob(des, initial *WorkflowTemplateJobsSparkRJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkRJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MainRFileUri) {
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

func canonicalizeWorkflowTemplateJobsSparkRJobProperties(des, initial *WorkflowTemplateJobsSparkRJobProperties, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkRJobProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsSparkRJobProperties(c *Client, des, nw *WorkflowTemplateJobsSparkRJobProperties) *WorkflowTemplateJobsSparkRJobProperties {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkRJobPropertiesSet(c *Client, des, nw []WorkflowTemplateJobsSparkRJobProperties) []WorkflowTemplateJobsSparkRJobProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkRJobProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkRJobProperties(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsSparkRJobLoggingConfig(des, initial *WorkflowTemplateJobsSparkRJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkRJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(des, initial *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c *Client, des, nw *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsSet(c *Client, des, nw []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsSparkSqlJob(des, initial *WorkflowTemplateJobsSparkSqlJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkSqlJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.QueryFileUri) {
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

func canonicalizeWorkflowTemplateJobsSparkSqlJobQueryList(des, initial *WorkflowTemplateJobsSparkSqlJobQueryList, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkSqlJobQueryList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsSparkSqlJobScriptVariables(des, initial *WorkflowTemplateJobsSparkSqlJobScriptVariables, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkSqlJobScriptVariables {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobScriptVariables(c *Client, des, nw *WorkflowTemplateJobsSparkSqlJobScriptVariables) *WorkflowTemplateJobsSparkSqlJobScriptVariables {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobScriptVariablesSet(c *Client, des, nw []WorkflowTemplateJobsSparkSqlJobScriptVariables) []WorkflowTemplateJobsSparkSqlJobScriptVariables {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkSqlJobScriptVariables
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkSqlJobScriptVariables(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsSparkSqlJobProperties(des, initial *WorkflowTemplateJobsSparkSqlJobProperties, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkSqlJobProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobProperties(c *Client, des, nw *WorkflowTemplateJobsSparkSqlJobProperties) *WorkflowTemplateJobsSparkSqlJobProperties {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobPropertiesSet(c *Client, des, nw []WorkflowTemplateJobsSparkSqlJobProperties) []WorkflowTemplateJobsSparkSqlJobProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkSqlJobProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkSqlJobProperties(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsSparkSqlJobLoggingConfig(des, initial *WorkflowTemplateJobsSparkSqlJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(des, initial *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels, opts ...dcl.ApplyOption) *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c *Client, des, nw *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsSet(c *Client, des, nw []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsPrestoJob(des, initial *WorkflowTemplateJobsPrestoJob, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPrestoJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.QueryFileUri) {
		des.QueryFileUri = initial.QueryFileUri
	}
	des.QueryList = canonicalizeWorkflowTemplateJobsPrestoJobQueryList(des.QueryList, initial.QueryList, opts...)
	if dcl.IsZeroValue(des.ContinueOnFailure) {
		des.ContinueOnFailure = initial.ContinueOnFailure
	}
	if dcl.IsZeroValue(des.OutputFormat) {
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

	nw.QueryList = canonicalizeNewWorkflowTemplateJobsPrestoJobQueryList(c, des.QueryList, nw.QueryList)
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

func canonicalizeWorkflowTemplateJobsPrestoJobQueryList(des, initial *WorkflowTemplateJobsPrestoJobQueryList, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPrestoJobQueryList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsPrestoJobProperties(des, initial *WorkflowTemplateJobsPrestoJobProperties, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPrestoJobProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsPrestoJobProperties(c *Client, des, nw *WorkflowTemplateJobsPrestoJobProperties) *WorkflowTemplateJobsPrestoJobProperties {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPrestoJobPropertiesSet(c *Client, des, nw []WorkflowTemplateJobsPrestoJobProperties) []WorkflowTemplateJobsPrestoJobProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPrestoJobProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPrestoJobProperties(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsPrestoJobLoggingConfig(des, initial *WorkflowTemplateJobsPrestoJobLoggingConfig, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPrestoJobLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(des, initial *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels, opts ...dcl.ApplyOption) *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c *Client, des, nw *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsSet(c *Client, des, nw []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsLabels(des, initial *WorkflowTemplateJobsLabels, opts ...dcl.ApplyOption) *WorkflowTemplateJobsLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
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

func canonicalizeNewWorkflowTemplateJobsLabels(c *Client, des, nw *WorkflowTemplateJobsLabels) *WorkflowTemplateJobsLabels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewWorkflowTemplateJobsLabelsSet(c *Client, des, nw []WorkflowTemplateJobsLabels) []WorkflowTemplateJobsLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkflowTemplateJobsLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareWorkflowTemplateJobsLabels(c, &d, &n) {
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

func canonicalizeWorkflowTemplateJobsScheduling(des, initial *WorkflowTemplateJobsScheduling, opts ...dcl.ApplyOption) *WorkflowTemplateJobsScheduling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateParameters(des, initial *WorkflowTemplateParameters, opts ...dcl.ApplyOption) *WorkflowTemplateParameters {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Fields) {
		des.Fields = initial.Fields
	}
	if dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	des.Validation = canonicalizeWorkflowTemplateParametersValidation(des.Validation, initial.Validation, opts...)

	return des
}

func canonicalizeNewWorkflowTemplateParameters(c *Client, des, nw *WorkflowTemplateParameters) *WorkflowTemplateParameters {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeWorkflowTemplateParametersValidation(des, initial *WorkflowTemplateParametersValidation, opts ...dcl.ApplyOption) *WorkflowTemplateParametersValidation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateParametersValidationRegex(des, initial *WorkflowTemplateParametersValidationRegex, opts ...dcl.ApplyOption) *WorkflowTemplateParametersValidationRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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

func canonicalizeWorkflowTemplateParametersValidationValues(des, initial *WorkflowTemplateParametersValidationValues, opts ...dcl.ApplyOption) *WorkflowTemplateParametersValidationValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*WorkflowTemplate)
		_ = r
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
	if !dcl.IsZeroValue(desired.Name) && !dcl.NameToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %#v\nACTUAL: %#v", desired.Name, actual.Name)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Version) && (dcl.IsZeroValue(actual.Version) || !reflect.DeepEqual(*desired.Version, *actual.Version)) {
		c.Config.Logger.Infof("Detected diff in Version.\nDESIRED: %#v\nACTUAL: %#v", desired.Version, actual.Version)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Version",
		})
	}
	if compareWorkflowTemplateLabelsSlice(c, desired.Labels, actual.Labels) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %#v\nACTUAL: %#v", desired.Labels, actual.Labels)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Labels",
		})
	}
	if compareWorkflowTemplatePlacement(c, desired.Placement, actual.Placement) {
		c.Config.Logger.Infof("Detected diff in Placement.\nDESIRED: %#v\nACTUAL: %#v", desired.Placement, actual.Placement)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Placement",
		})
	}
	if compareWorkflowTemplateJobsSlice(c, desired.Jobs, actual.Jobs) {
		c.Config.Logger.Infof("Detected diff in Jobs.\nDESIRED: %#v\nACTUAL: %#v", desired.Jobs, actual.Jobs)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Jobs",
		})
	}
	if compareWorkflowTemplateParametersSlice(c, desired.Parameters, actual.Parameters) {
		c.Config.Logger.Infof("Detected diff in Parameters.\nDESIRED: %#v\nACTUAL: %#v", desired.Parameters, actual.Parameters)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Parameters",
		})
	}
	if !dcl.IsZeroValue(desired.Project) && !dcl.NameToSelfLink(desired.Project, actual.Project) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %#v\nACTUAL: %#v", desired.Project, actual.Project)
		diffs = append(diffs, workflowTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Project",
		})
	}
	if !dcl.IsZeroValue(desired.Location) && !dcl.NameToSelfLink(desired.Location, actual.Location) {
		c.Config.Logger.Infof("Detected diff in Location.\nDESIRED: %#v\nACTUAL: %#v", desired.Location, actual.Location)
		diffs = append(diffs, workflowTemplateDiff{
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
func compareWorkflowTemplateLabelsSlice(c *Client, desired, actual []WorkflowTemplateLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateLabels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateLabels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateLabels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateLabels(c *Client, desired, actual *WorkflowTemplateLabels) bool {
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
	if !reflect.DeepEqual(desired.ClusterName, actual.ClusterName) && !dcl.IsZeroValue(desired.ClusterName) && !(dcl.IsEmptyValueIndirect(desired.ClusterName) && dcl.IsZeroValue(actual.ClusterName)) {
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
	if compareWorkflowTemplatePlacementManagedClusterLabelsSlice(c, desired.Labels, actual.Labels) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	return false
}
func compareWorkflowTemplatePlacementManagedClusterLabelsSlice(c *Client, desired, actual []WorkflowTemplatePlacementManagedClusterLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplatePlacementManagedClusterLabels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplatePlacementManagedClusterLabels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacementManagedClusterLabels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplatePlacementManagedClusterLabels(c *Client, desired, actual *WorkflowTemplatePlacementManagedClusterLabels) bool {
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
	if !reflect.DeepEqual(desired.Zone, actual.Zone) && !dcl.IsZeroValue(desired.Zone) && !(dcl.IsEmptyValueIndirect(desired.Zone) && dcl.IsZeroValue(actual.Zone)) {
		c.Config.Logger.Infof("Diff in Zone. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Zone), dcl.SprintResource(actual.Zone))
		return true
	}
	if actual.ClusterLabels == nil && desired.ClusterLabels != nil && !dcl.IsEmptyValueIndirect(desired.ClusterLabels) {
		c.Config.Logger.Infof("desired ClusterLabels %s - but actually nil", dcl.SprintResource(desired.ClusterLabels))
		return true
	}
	if compareWorkflowTemplatePlacementClusterSelectorClusterLabelsSlice(c, desired.ClusterLabels, actual.ClusterLabels) && !dcl.IsZeroValue(desired.ClusterLabels) {
		c.Config.Logger.Infof("Diff in ClusterLabels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterLabels), dcl.SprintResource(actual.ClusterLabels))
		return true
	}
	return false
}
func compareWorkflowTemplatePlacementClusterSelectorClusterLabelsSlice(c *Client, desired, actual []WorkflowTemplatePlacementClusterSelectorClusterLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplatePlacementClusterSelectorClusterLabels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplatePlacementClusterSelectorClusterLabels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplatePlacementClusterSelectorClusterLabels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplatePlacementClusterSelectorClusterLabels(c *Client, desired, actual *WorkflowTemplatePlacementClusterSelectorClusterLabels) bool {
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
	if !reflect.DeepEqual(desired.StepId, actual.StepId) && !dcl.IsZeroValue(desired.StepId) && !(dcl.IsEmptyValueIndirect(desired.StepId) && dcl.IsZeroValue(actual.StepId)) {
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
	if compareWorkflowTemplateJobsLabelsSlice(c, desired.Labels, actual.Labels) && !dcl.IsZeroValue(desired.Labels) {
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
	if !dcl.SliceEquals(desired.PrerequisiteStepIds, actual.PrerequisiteStepIds) && !dcl.IsZeroValue(desired.PrerequisiteStepIds) {
		c.Config.Logger.Infof("Diff in PrerequisiteStepIds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrerequisiteStepIds), dcl.SprintResource(actual.PrerequisiteStepIds))
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
	if !reflect.DeepEqual(desired.MainJarFileUri, actual.MainJarFileUri) && !dcl.IsZeroValue(desired.MainJarFileUri) && !(dcl.IsEmptyValueIndirect(desired.MainJarFileUri) && dcl.IsZeroValue(actual.MainJarFileUri)) {
		c.Config.Logger.Infof("Diff in MainJarFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainJarFileUri), dcl.SprintResource(actual.MainJarFileUri))
		return true
	}
	if actual.MainClass == nil && desired.MainClass != nil && !dcl.IsEmptyValueIndirect(desired.MainClass) {
		c.Config.Logger.Infof("desired MainClass %s - but actually nil", dcl.SprintResource(desired.MainClass))
		return true
	}
	if !reflect.DeepEqual(desired.MainClass, actual.MainClass) && !dcl.IsZeroValue(desired.MainClass) && !(dcl.IsEmptyValueIndirect(desired.MainClass) && dcl.IsZeroValue(actual.MainClass)) {
		c.Config.Logger.Infof("Diff in MainClass. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainClass), dcl.SprintResource(actual.MainClass))
		return true
	}
	if actual.Args == nil && desired.Args != nil && !dcl.IsEmptyValueIndirect(desired.Args) {
		c.Config.Logger.Infof("desired Args %s - but actually nil", dcl.SprintResource(desired.Args))
		return true
	}
	if !dcl.SliceEquals(desired.Args, actual.Args) && !dcl.IsZeroValue(desired.Args) {
		c.Config.Logger.Infof("Diff in Args. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Args), dcl.SprintResource(actual.Args))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.SliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
		return true
	}
	if actual.FileUris == nil && desired.FileUris != nil && !dcl.IsEmptyValueIndirect(desired.FileUris) {
		c.Config.Logger.Infof("desired FileUris %s - but actually nil", dcl.SprintResource(desired.FileUris))
		return true
	}
	if !dcl.SliceEquals(desired.FileUris, actual.FileUris) && !dcl.IsZeroValue(desired.FileUris) {
		c.Config.Logger.Infof("Diff in FileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileUris), dcl.SprintResource(actual.FileUris))
		return true
	}
	if actual.ArchiveUris == nil && desired.ArchiveUris != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveUris) {
		c.Config.Logger.Infof("desired ArchiveUris %s - but actually nil", dcl.SprintResource(desired.ArchiveUris))
		return true
	}
	if !dcl.SliceEquals(desired.ArchiveUris, actual.ArchiveUris) && !dcl.IsZeroValue(desired.ArchiveUris) {
		c.Config.Logger.Infof("Diff in ArchiveUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveUris), dcl.SprintResource(actual.ArchiveUris))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if compareWorkflowTemplateJobsHadoopJobPropertiesSlice(c, desired.Properties, actual.Properties) && !dcl.IsZeroValue(desired.Properties) {
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
func compareWorkflowTemplateJobsHadoopJobPropertiesSlice(c *Client, desired, actual []WorkflowTemplateJobsHadoopJobProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHadoopJobProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsHadoopJobProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHadoopJobProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHadoopJobProperties(c *Client, desired, actual *WorkflowTemplateJobsHadoopJobProperties) bool {
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
	if compareWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsSlice(c, desired.DriverLogLevels, actual.DriverLogLevels) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsSlice(c *Client, desired, actual []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c *Client, desired, actual *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) bool {
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
	if !reflect.DeepEqual(desired.MainJarFileUri, actual.MainJarFileUri) && !dcl.IsZeroValue(desired.MainJarFileUri) && !(dcl.IsEmptyValueIndirect(desired.MainJarFileUri) && dcl.IsZeroValue(actual.MainJarFileUri)) {
		c.Config.Logger.Infof("Diff in MainJarFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainJarFileUri), dcl.SprintResource(actual.MainJarFileUri))
		return true
	}
	if actual.MainClass == nil && desired.MainClass != nil && !dcl.IsEmptyValueIndirect(desired.MainClass) {
		c.Config.Logger.Infof("desired MainClass %s - but actually nil", dcl.SprintResource(desired.MainClass))
		return true
	}
	if !reflect.DeepEqual(desired.MainClass, actual.MainClass) && !dcl.IsZeroValue(desired.MainClass) && !(dcl.IsEmptyValueIndirect(desired.MainClass) && dcl.IsZeroValue(actual.MainClass)) {
		c.Config.Logger.Infof("Diff in MainClass. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainClass), dcl.SprintResource(actual.MainClass))
		return true
	}
	if actual.Args == nil && desired.Args != nil && !dcl.IsEmptyValueIndirect(desired.Args) {
		c.Config.Logger.Infof("desired Args %s - but actually nil", dcl.SprintResource(desired.Args))
		return true
	}
	if !dcl.SliceEquals(desired.Args, actual.Args) && !dcl.IsZeroValue(desired.Args) {
		c.Config.Logger.Infof("Diff in Args. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Args), dcl.SprintResource(actual.Args))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.SliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
		return true
	}
	if actual.FileUris == nil && desired.FileUris != nil && !dcl.IsEmptyValueIndirect(desired.FileUris) {
		c.Config.Logger.Infof("desired FileUris %s - but actually nil", dcl.SprintResource(desired.FileUris))
		return true
	}
	if !dcl.SliceEquals(desired.FileUris, actual.FileUris) && !dcl.IsZeroValue(desired.FileUris) {
		c.Config.Logger.Infof("Diff in FileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileUris), dcl.SprintResource(actual.FileUris))
		return true
	}
	if actual.ArchiveUris == nil && desired.ArchiveUris != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveUris) {
		c.Config.Logger.Infof("desired ArchiveUris %s - but actually nil", dcl.SprintResource(desired.ArchiveUris))
		return true
	}
	if !dcl.SliceEquals(desired.ArchiveUris, actual.ArchiveUris) && !dcl.IsZeroValue(desired.ArchiveUris) {
		c.Config.Logger.Infof("Diff in ArchiveUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveUris), dcl.SprintResource(actual.ArchiveUris))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if compareWorkflowTemplateJobsSparkJobPropertiesSlice(c, desired.Properties, actual.Properties) && !dcl.IsZeroValue(desired.Properties) {
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
func compareWorkflowTemplateJobsSparkJobPropertiesSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkJobProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkJobProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkJobProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkJobProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkJobProperties(c *Client, desired, actual *WorkflowTemplateJobsSparkJobProperties) bool {
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
	if compareWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsSlice(c, desired.DriverLogLevels, actual.DriverLogLevels) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c *Client, desired, actual *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) bool {
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
	if !reflect.DeepEqual(desired.MainPythonFileUri, actual.MainPythonFileUri) && !dcl.IsZeroValue(desired.MainPythonFileUri) && !(dcl.IsEmptyValueIndirect(desired.MainPythonFileUri) && dcl.IsZeroValue(actual.MainPythonFileUri)) {
		c.Config.Logger.Infof("Diff in MainPythonFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainPythonFileUri), dcl.SprintResource(actual.MainPythonFileUri))
		return true
	}
	if actual.Args == nil && desired.Args != nil && !dcl.IsEmptyValueIndirect(desired.Args) {
		c.Config.Logger.Infof("desired Args %s - but actually nil", dcl.SprintResource(desired.Args))
		return true
	}
	if !dcl.SliceEquals(desired.Args, actual.Args) && !dcl.IsZeroValue(desired.Args) {
		c.Config.Logger.Infof("Diff in Args. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Args), dcl.SprintResource(actual.Args))
		return true
	}
	if actual.PythonFileUris == nil && desired.PythonFileUris != nil && !dcl.IsEmptyValueIndirect(desired.PythonFileUris) {
		c.Config.Logger.Infof("desired PythonFileUris %s - but actually nil", dcl.SprintResource(desired.PythonFileUris))
		return true
	}
	if !dcl.SliceEquals(desired.PythonFileUris, actual.PythonFileUris) && !dcl.IsZeroValue(desired.PythonFileUris) {
		c.Config.Logger.Infof("Diff in PythonFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PythonFileUris), dcl.SprintResource(actual.PythonFileUris))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.SliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
		return true
	}
	if actual.FileUris == nil && desired.FileUris != nil && !dcl.IsEmptyValueIndirect(desired.FileUris) {
		c.Config.Logger.Infof("desired FileUris %s - but actually nil", dcl.SprintResource(desired.FileUris))
		return true
	}
	if !dcl.SliceEquals(desired.FileUris, actual.FileUris) && !dcl.IsZeroValue(desired.FileUris) {
		c.Config.Logger.Infof("Diff in FileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileUris), dcl.SprintResource(actual.FileUris))
		return true
	}
	if actual.ArchiveUris == nil && desired.ArchiveUris != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveUris) {
		c.Config.Logger.Infof("desired ArchiveUris %s - but actually nil", dcl.SprintResource(desired.ArchiveUris))
		return true
	}
	if !dcl.SliceEquals(desired.ArchiveUris, actual.ArchiveUris) && !dcl.IsZeroValue(desired.ArchiveUris) {
		c.Config.Logger.Infof("Diff in ArchiveUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveUris), dcl.SprintResource(actual.ArchiveUris))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if compareWorkflowTemplateJobsPysparkJobPropertiesSlice(c, desired.Properties, actual.Properties) && !dcl.IsZeroValue(desired.Properties) {
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
func compareWorkflowTemplateJobsPysparkJobPropertiesSlice(c *Client, desired, actual []WorkflowTemplateJobsPysparkJobProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPysparkJobProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPysparkJobProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPysparkJobProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPysparkJobProperties(c *Client, desired, actual *WorkflowTemplateJobsPysparkJobProperties) bool {
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
	if compareWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsSlice(c, desired.DriverLogLevels, actual.DriverLogLevels) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsSlice(c *Client, desired, actual []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c *Client, desired, actual *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) bool {
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
	if !reflect.DeepEqual(desired.QueryFileUri, actual.QueryFileUri) && !dcl.IsZeroValue(desired.QueryFileUri) && !(dcl.IsEmptyValueIndirect(desired.QueryFileUri) && dcl.IsZeroValue(actual.QueryFileUri)) {
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
	if !reflect.DeepEqual(desired.ContinueOnFailure, actual.ContinueOnFailure) && !dcl.IsZeroValue(desired.ContinueOnFailure) && !(dcl.IsEmptyValueIndirect(desired.ContinueOnFailure) && dcl.IsZeroValue(actual.ContinueOnFailure)) {
		c.Config.Logger.Infof("Diff in ContinueOnFailure. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContinueOnFailure), dcl.SprintResource(actual.ContinueOnFailure))
		return true
	}
	if actual.ScriptVariables == nil && desired.ScriptVariables != nil && !dcl.IsEmptyValueIndirect(desired.ScriptVariables) {
		c.Config.Logger.Infof("desired ScriptVariables %s - but actually nil", dcl.SprintResource(desired.ScriptVariables))
		return true
	}
	if compareWorkflowTemplateJobsHiveJobScriptVariablesSlice(c, desired.ScriptVariables, actual.ScriptVariables) && !dcl.IsZeroValue(desired.ScriptVariables) {
		c.Config.Logger.Infof("Diff in ScriptVariables. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScriptVariables), dcl.SprintResource(actual.ScriptVariables))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if compareWorkflowTemplateJobsHiveJobPropertiesSlice(c, desired.Properties, actual.Properties) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.SliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
		c.Config.Logger.Infof("Diff in JarFileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.JarFileUris), dcl.SprintResource(actual.JarFileUris))
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
	if !dcl.SliceEquals(desired.Queries, actual.Queries) && !dcl.IsZeroValue(desired.Queries) {
		c.Config.Logger.Infof("Diff in Queries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Queries), dcl.SprintResource(actual.Queries))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsHiveJobScriptVariablesSlice(c *Client, desired, actual []WorkflowTemplateJobsHiveJobScriptVariables) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHiveJobScriptVariables, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsHiveJobScriptVariables(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHiveJobScriptVariables, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHiveJobScriptVariables(c *Client, desired, actual *WorkflowTemplateJobsHiveJobScriptVariables) bool {
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
func compareWorkflowTemplateJobsHiveJobPropertiesSlice(c *Client, desired, actual []WorkflowTemplateJobsHiveJobProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHiveJobProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsHiveJobProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHiveJobProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHiveJobProperties(c *Client, desired, actual *WorkflowTemplateJobsHiveJobProperties) bool {
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
	if !reflect.DeepEqual(desired.QueryFileUri, actual.QueryFileUri) && !dcl.IsZeroValue(desired.QueryFileUri) && !(dcl.IsEmptyValueIndirect(desired.QueryFileUri) && dcl.IsZeroValue(actual.QueryFileUri)) {
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
	if !reflect.DeepEqual(desired.ContinueOnFailure, actual.ContinueOnFailure) && !dcl.IsZeroValue(desired.ContinueOnFailure) && !(dcl.IsEmptyValueIndirect(desired.ContinueOnFailure) && dcl.IsZeroValue(actual.ContinueOnFailure)) {
		c.Config.Logger.Infof("Diff in ContinueOnFailure. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContinueOnFailure), dcl.SprintResource(actual.ContinueOnFailure))
		return true
	}
	if actual.ScriptVariables == nil && desired.ScriptVariables != nil && !dcl.IsEmptyValueIndirect(desired.ScriptVariables) {
		c.Config.Logger.Infof("desired ScriptVariables %s - but actually nil", dcl.SprintResource(desired.ScriptVariables))
		return true
	}
	if compareWorkflowTemplateJobsPigJobScriptVariablesSlice(c, desired.ScriptVariables, actual.ScriptVariables) && !dcl.IsZeroValue(desired.ScriptVariables) {
		c.Config.Logger.Infof("Diff in ScriptVariables. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScriptVariables), dcl.SprintResource(actual.ScriptVariables))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if compareWorkflowTemplateJobsPigJobPropertiesSlice(c, desired.Properties, actual.Properties) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.SliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
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
	if !dcl.SliceEquals(desired.Queries, actual.Queries) && !dcl.IsZeroValue(desired.Queries) {
		c.Config.Logger.Infof("Diff in Queries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Queries), dcl.SprintResource(actual.Queries))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsPigJobScriptVariablesSlice(c *Client, desired, actual []WorkflowTemplateJobsPigJobScriptVariables) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJobScriptVariables, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPigJobScriptVariables(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobScriptVariables, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJobScriptVariables(c *Client, desired, actual *WorkflowTemplateJobsPigJobScriptVariables) bool {
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
func compareWorkflowTemplateJobsPigJobPropertiesSlice(c *Client, desired, actual []WorkflowTemplateJobsPigJobProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJobProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPigJobProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJobProperties(c *Client, desired, actual *WorkflowTemplateJobsPigJobProperties) bool {
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
	if compareWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsSlice(c, desired.DriverLogLevels, actual.DriverLogLevels) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsSlice(c *Client, desired, actual []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c *Client, desired, actual *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) bool {
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
	if !reflect.DeepEqual(desired.MainRFileUri, actual.MainRFileUri) && !dcl.IsZeroValue(desired.MainRFileUri) && !(dcl.IsEmptyValueIndirect(desired.MainRFileUri) && dcl.IsZeroValue(actual.MainRFileUri)) {
		c.Config.Logger.Infof("Diff in MainRFileUri. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MainRFileUri), dcl.SprintResource(actual.MainRFileUri))
		return true
	}
	if actual.Args == nil && desired.Args != nil && !dcl.IsEmptyValueIndirect(desired.Args) {
		c.Config.Logger.Infof("desired Args %s - but actually nil", dcl.SprintResource(desired.Args))
		return true
	}
	if !dcl.SliceEquals(desired.Args, actual.Args) && !dcl.IsZeroValue(desired.Args) {
		c.Config.Logger.Infof("Diff in Args. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Args), dcl.SprintResource(actual.Args))
		return true
	}
	if actual.FileUris == nil && desired.FileUris != nil && !dcl.IsEmptyValueIndirect(desired.FileUris) {
		c.Config.Logger.Infof("desired FileUris %s - but actually nil", dcl.SprintResource(desired.FileUris))
		return true
	}
	if !dcl.SliceEquals(desired.FileUris, actual.FileUris) && !dcl.IsZeroValue(desired.FileUris) {
		c.Config.Logger.Infof("Diff in FileUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FileUris), dcl.SprintResource(actual.FileUris))
		return true
	}
	if actual.ArchiveUris == nil && desired.ArchiveUris != nil && !dcl.IsEmptyValueIndirect(desired.ArchiveUris) {
		c.Config.Logger.Infof("desired ArchiveUris %s - but actually nil", dcl.SprintResource(desired.ArchiveUris))
		return true
	}
	if !dcl.SliceEquals(desired.ArchiveUris, actual.ArchiveUris) && !dcl.IsZeroValue(desired.ArchiveUris) {
		c.Config.Logger.Infof("Diff in ArchiveUris. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ArchiveUris), dcl.SprintResource(actual.ArchiveUris))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if compareWorkflowTemplateJobsSparkRJobPropertiesSlice(c, desired.Properties, actual.Properties) && !dcl.IsZeroValue(desired.Properties) {
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
func compareWorkflowTemplateJobsSparkRJobPropertiesSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkRJobProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkRJobProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkRJobProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkRJobProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkRJobProperties(c *Client, desired, actual *WorkflowTemplateJobsSparkRJobProperties) bool {
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
	if compareWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsSlice(c, desired.DriverLogLevels, actual.DriverLogLevels) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c *Client, desired, actual *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) bool {
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
	if !reflect.DeepEqual(desired.QueryFileUri, actual.QueryFileUri) && !dcl.IsZeroValue(desired.QueryFileUri) && !(dcl.IsEmptyValueIndirect(desired.QueryFileUri) && dcl.IsZeroValue(actual.QueryFileUri)) {
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
	if compareWorkflowTemplateJobsSparkSqlJobScriptVariablesSlice(c, desired.ScriptVariables, actual.ScriptVariables) && !dcl.IsZeroValue(desired.ScriptVariables) {
		c.Config.Logger.Infof("Diff in ScriptVariables. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ScriptVariables), dcl.SprintResource(actual.ScriptVariables))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if compareWorkflowTemplateJobsSparkSqlJobPropertiesSlice(c, desired.Properties, actual.Properties) && !dcl.IsZeroValue(desired.Properties) {
		c.Config.Logger.Infof("Diff in Properties. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Properties), dcl.SprintResource(actual.Properties))
		return true
	}
	if actual.JarFileUris == nil && desired.JarFileUris != nil && !dcl.IsEmptyValueIndirect(desired.JarFileUris) {
		c.Config.Logger.Infof("desired JarFileUris %s - but actually nil", dcl.SprintResource(desired.JarFileUris))
		return true
	}
	if !dcl.SliceEquals(desired.JarFileUris, actual.JarFileUris) && !dcl.IsZeroValue(desired.JarFileUris) {
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
	if !dcl.SliceEquals(desired.Queries, actual.Queries) && !dcl.IsZeroValue(desired.Queries) {
		c.Config.Logger.Infof("Diff in Queries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Queries), dcl.SprintResource(actual.Queries))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsSparkSqlJobScriptVariablesSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkSqlJobScriptVariables) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJobScriptVariables, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkSqlJobScriptVariables(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobScriptVariables, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobScriptVariables(c *Client, desired, actual *WorkflowTemplateJobsSparkSqlJobScriptVariables) bool {
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
func compareWorkflowTemplateJobsSparkSqlJobPropertiesSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkSqlJobProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJobProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkSqlJobProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobProperties(c *Client, desired, actual *WorkflowTemplateJobsSparkSqlJobProperties) bool {
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
	if compareWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsSlice(c, desired.DriverLogLevels, actual.DriverLogLevels) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c *Client, desired, actual *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) bool {
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
	if !reflect.DeepEqual(desired.QueryFileUri, actual.QueryFileUri) && !dcl.IsZeroValue(desired.QueryFileUri) && !(dcl.IsEmptyValueIndirect(desired.QueryFileUri) && dcl.IsZeroValue(actual.QueryFileUri)) {
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
	if !reflect.DeepEqual(desired.ContinueOnFailure, actual.ContinueOnFailure) && !dcl.IsZeroValue(desired.ContinueOnFailure) && !(dcl.IsEmptyValueIndirect(desired.ContinueOnFailure) && dcl.IsZeroValue(actual.ContinueOnFailure)) {
		c.Config.Logger.Infof("Diff in ContinueOnFailure. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContinueOnFailure), dcl.SprintResource(actual.ContinueOnFailure))
		return true
	}
	if actual.OutputFormat == nil && desired.OutputFormat != nil && !dcl.IsEmptyValueIndirect(desired.OutputFormat) {
		c.Config.Logger.Infof("desired OutputFormat %s - but actually nil", dcl.SprintResource(desired.OutputFormat))
		return true
	}
	if !reflect.DeepEqual(desired.OutputFormat, actual.OutputFormat) && !dcl.IsZeroValue(desired.OutputFormat) && !(dcl.IsEmptyValueIndirect(desired.OutputFormat) && dcl.IsZeroValue(actual.OutputFormat)) {
		c.Config.Logger.Infof("Diff in OutputFormat. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OutputFormat), dcl.SprintResource(actual.OutputFormat))
		return true
	}
	if actual.ClientTags == nil && desired.ClientTags != nil && !dcl.IsEmptyValueIndirect(desired.ClientTags) {
		c.Config.Logger.Infof("desired ClientTags %s - but actually nil", dcl.SprintResource(desired.ClientTags))
		return true
	}
	if !dcl.SliceEquals(desired.ClientTags, actual.ClientTags) && !dcl.IsZeroValue(desired.ClientTags) {
		c.Config.Logger.Infof("Diff in ClientTags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientTags), dcl.SprintResource(actual.ClientTags))
		return true
	}
	if actual.Properties == nil && desired.Properties != nil && !dcl.IsEmptyValueIndirect(desired.Properties) {
		c.Config.Logger.Infof("desired Properties %s - but actually nil", dcl.SprintResource(desired.Properties))
		return true
	}
	if compareWorkflowTemplateJobsPrestoJobPropertiesSlice(c, desired.Properties, actual.Properties) && !dcl.IsZeroValue(desired.Properties) {
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
	if !dcl.SliceEquals(desired.Queries, actual.Queries) && !dcl.IsZeroValue(desired.Queries) {
		c.Config.Logger.Infof("Diff in Queries. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Queries), dcl.SprintResource(actual.Queries))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsPrestoJobPropertiesSlice(c *Client, desired, actual []WorkflowTemplateJobsPrestoJobProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPrestoJobProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPrestoJobProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJobProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJobProperties(c *Client, desired, actual *WorkflowTemplateJobsPrestoJobProperties) bool {
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
	if compareWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsSlice(c, desired.DriverLogLevels, actual.DriverLogLevels) && !dcl.IsZeroValue(desired.DriverLogLevels) {
		c.Config.Logger.Infof("Diff in DriverLogLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DriverLogLevels), dcl.SprintResource(actual.DriverLogLevels))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsSlice(c *Client, desired, actual []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c *Client, desired, actual *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) bool {
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
func compareWorkflowTemplateJobsLabelsSlice(c *Client, desired, actual []WorkflowTemplateJobsLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsLabels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsLabels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsLabels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsLabels(c *Client, desired, actual *WorkflowTemplateJobsLabels) bool {
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
	if !reflect.DeepEqual(desired.MaxFailuresPerHour, actual.MaxFailuresPerHour) && !dcl.IsZeroValue(desired.MaxFailuresPerHour) && !(dcl.IsEmptyValueIndirect(desired.MaxFailuresPerHour) && dcl.IsZeroValue(actual.MaxFailuresPerHour)) {
		c.Config.Logger.Infof("Diff in MaxFailuresPerHour. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxFailuresPerHour), dcl.SprintResource(actual.MaxFailuresPerHour))
		return true
	}
	if actual.MaxFailuresTotal == nil && desired.MaxFailuresTotal != nil && !dcl.IsEmptyValueIndirect(desired.MaxFailuresTotal) {
		c.Config.Logger.Infof("desired MaxFailuresTotal %s - but actually nil", dcl.SprintResource(desired.MaxFailuresTotal))
		return true
	}
	if !reflect.DeepEqual(desired.MaxFailuresTotal, actual.MaxFailuresTotal) && !dcl.IsZeroValue(desired.MaxFailuresTotal) && !(dcl.IsEmptyValueIndirect(desired.MaxFailuresTotal) && dcl.IsZeroValue(actual.MaxFailuresTotal)) {
		c.Config.Logger.Infof("Diff in MaxFailuresTotal. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxFailuresTotal), dcl.SprintResource(actual.MaxFailuresTotal))
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
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Fields == nil && desired.Fields != nil && !dcl.IsEmptyValueIndirect(desired.Fields) {
		c.Config.Logger.Infof("desired Fields %s - but actually nil", dcl.SprintResource(desired.Fields))
		return true
	}
	if !dcl.SliceEquals(desired.Fields, actual.Fields) && !dcl.IsZeroValue(desired.Fields) {
		c.Config.Logger.Infof("Diff in Fields. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Fields), dcl.SprintResource(actual.Fields))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !reflect.DeepEqual(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) && !(dcl.IsEmptyValueIndirect(desired.Description) && dcl.IsZeroValue(actual.Description)) {
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
	if !dcl.SliceEquals(desired.Regexes, actual.Regexes) && !dcl.IsZeroValue(desired.Regexes) {
		c.Config.Logger.Infof("Diff in Regexes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Regexes), dcl.SprintResource(actual.Regexes))
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
	if !dcl.SliceEquals(desired.Values, actual.Values) && !dcl.IsZeroValue(desired.Values) {
		c.Config.Logger.Infof("Diff in Values. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Values), dcl.SprintResource(actual.Values))
		return true
	}
	return false
}
func compareWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, desired, actual []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(c *Client, desired, actual *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(c *Client, desired, actual *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, desired, actual []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(c *Client, desired, actual *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, desired, actual []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(c *Client, desired, actual *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(c *Client, desired, actual *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, desired, actual []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(c *Client, desired, actual *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, desired, actual []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(c *Client, desired, actual *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *WorkflowTemplate) urlNormalized() *WorkflowTemplate {
	normalized := deepcopy.Copy(*r).(WorkflowTemplate)
	normalized.Name = dcl.SelfLinkToName(r.Name)
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
		return dcl.URL("projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}", "https://dataproc.googleapis.com/v1/", userBasePath, fields), nil

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
	if v, err := expandWorkflowTemplateLabelsSlice(c, f.Labels); err != nil {
		return nil, fmt.Errorf("error expanding Labels into labels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Labels = flattenWorkflowTemplateLabelsSlice(c, m["labels"])
	r.Placement = flattenWorkflowTemplatePlacement(c, m["placement"])
	r.Jobs = flattenWorkflowTemplateJobsSlice(c, m["jobs"])
	r.Parameters = flattenWorkflowTemplateParametersSlice(c, m["parameters"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandWorkflowTemplateLabelsMap expands the contents of WorkflowTemplateLabels into a JSON
// request object.
func expandWorkflowTemplateLabelsMap(c *Client, f map[string]WorkflowTemplateLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateLabelsSlice expands the contents of WorkflowTemplateLabels into a JSON
// request object.
func expandWorkflowTemplateLabelsSlice(c *Client, f []WorkflowTemplateLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateLabelsMap flattens the contents of WorkflowTemplateLabels from a JSON
// response object.
func flattenWorkflowTemplateLabelsMap(c *Client, i interface{}) map[string]WorkflowTemplateLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateLabels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateLabels{}
	}

	items := make(map[string]WorkflowTemplateLabels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateLabelsSlice flattens the contents of WorkflowTemplateLabels from a JSON
// response object.
func flattenWorkflowTemplateLabelsSlice(c *Client, i interface{}) []WorkflowTemplateLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateLabels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateLabels{}
	}

	items := make([]WorkflowTemplateLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateLabels expands an instance of WorkflowTemplateLabels into a JSON
// request object.
func expandWorkflowTemplateLabels(c *Client, f *WorkflowTemplateLabels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateLabels flattens an instance of WorkflowTemplateLabels from a JSON
// response object.
func flattenWorkflowTemplateLabels(c *Client, i interface{}) *WorkflowTemplateLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateLabels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplatePlacementManagedClusterLabelsSlice(c, f.Labels); err != nil {
		return nil, fmt.Errorf("error expanding Labels into labels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Labels = flattenWorkflowTemplatePlacementManagedClusterLabelsSlice(c, m["labels"])

	return r
}

// expandWorkflowTemplatePlacementManagedClusterLabelsMap expands the contents of WorkflowTemplatePlacementManagedClusterLabels into a JSON
// request object.
func expandWorkflowTemplatePlacementManagedClusterLabelsMap(c *Client, f map[string]WorkflowTemplatePlacementManagedClusterLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplatePlacementManagedClusterLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplatePlacementManagedClusterLabelsSlice expands the contents of WorkflowTemplatePlacementManagedClusterLabels into a JSON
// request object.
func expandWorkflowTemplatePlacementManagedClusterLabelsSlice(c *Client, f []WorkflowTemplatePlacementManagedClusterLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplatePlacementManagedClusterLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplatePlacementManagedClusterLabelsMap flattens the contents of WorkflowTemplatePlacementManagedClusterLabels from a JSON
// response object.
func flattenWorkflowTemplatePlacementManagedClusterLabelsMap(c *Client, i interface{}) map[string]WorkflowTemplatePlacementManagedClusterLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplatePlacementManagedClusterLabels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplatePlacementManagedClusterLabels{}
	}

	items := make(map[string]WorkflowTemplatePlacementManagedClusterLabels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplatePlacementManagedClusterLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplatePlacementManagedClusterLabelsSlice flattens the contents of WorkflowTemplatePlacementManagedClusterLabels from a JSON
// response object.
func flattenWorkflowTemplatePlacementManagedClusterLabelsSlice(c *Client, i interface{}) []WorkflowTemplatePlacementManagedClusterLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplatePlacementManagedClusterLabels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplatePlacementManagedClusterLabels{}
	}

	items := make([]WorkflowTemplatePlacementManagedClusterLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplatePlacementManagedClusterLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplatePlacementManagedClusterLabels expands an instance of WorkflowTemplatePlacementManagedClusterLabels into a JSON
// request object.
func expandWorkflowTemplatePlacementManagedClusterLabels(c *Client, f *WorkflowTemplatePlacementManagedClusterLabels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplatePlacementManagedClusterLabels flattens an instance of WorkflowTemplatePlacementManagedClusterLabels from a JSON
// response object.
func flattenWorkflowTemplatePlacementManagedClusterLabels(c *Client, i interface{}) *WorkflowTemplatePlacementManagedClusterLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplatePlacementManagedClusterLabels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplatePlacementClusterSelectorClusterLabelsSlice(c, f.ClusterLabels); err != nil {
		return nil, fmt.Errorf("error expanding ClusterLabels into clusterLabels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.ClusterLabels = flattenWorkflowTemplatePlacementClusterSelectorClusterLabelsSlice(c, m["clusterLabels"])

	return r
}

// expandWorkflowTemplatePlacementClusterSelectorClusterLabelsMap expands the contents of WorkflowTemplatePlacementClusterSelectorClusterLabels into a JSON
// request object.
func expandWorkflowTemplatePlacementClusterSelectorClusterLabelsMap(c *Client, f map[string]WorkflowTemplatePlacementClusterSelectorClusterLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplatePlacementClusterSelectorClusterLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplatePlacementClusterSelectorClusterLabelsSlice expands the contents of WorkflowTemplatePlacementClusterSelectorClusterLabels into a JSON
// request object.
func expandWorkflowTemplatePlacementClusterSelectorClusterLabelsSlice(c *Client, f []WorkflowTemplatePlacementClusterSelectorClusterLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplatePlacementClusterSelectorClusterLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplatePlacementClusterSelectorClusterLabelsMap flattens the contents of WorkflowTemplatePlacementClusterSelectorClusterLabels from a JSON
// response object.
func flattenWorkflowTemplatePlacementClusterSelectorClusterLabelsMap(c *Client, i interface{}) map[string]WorkflowTemplatePlacementClusterSelectorClusterLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplatePlacementClusterSelectorClusterLabels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplatePlacementClusterSelectorClusterLabels{}
	}

	items := make(map[string]WorkflowTemplatePlacementClusterSelectorClusterLabels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplatePlacementClusterSelectorClusterLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplatePlacementClusterSelectorClusterLabelsSlice flattens the contents of WorkflowTemplatePlacementClusterSelectorClusterLabels from a JSON
// response object.
func flattenWorkflowTemplatePlacementClusterSelectorClusterLabelsSlice(c *Client, i interface{}) []WorkflowTemplatePlacementClusterSelectorClusterLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplatePlacementClusterSelectorClusterLabels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplatePlacementClusterSelectorClusterLabels{}
	}

	items := make([]WorkflowTemplatePlacementClusterSelectorClusterLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplatePlacementClusterSelectorClusterLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplatePlacementClusterSelectorClusterLabels expands an instance of WorkflowTemplatePlacementClusterSelectorClusterLabels into a JSON
// request object.
func expandWorkflowTemplatePlacementClusterSelectorClusterLabels(c *Client, f *WorkflowTemplatePlacementClusterSelectorClusterLabels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplatePlacementClusterSelectorClusterLabels flattens an instance of WorkflowTemplatePlacementClusterSelectorClusterLabels from a JSON
// response object.
func flattenWorkflowTemplatePlacementClusterSelectorClusterLabels(c *Client, i interface{}) *WorkflowTemplatePlacementClusterSelectorClusterLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplatePlacementClusterSelectorClusterLabels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplateJobsLabelsSlice(c, f.Labels); err != nil {
		return nil, fmt.Errorf("error expanding Labels into labels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Labels = flattenWorkflowTemplateJobsLabelsSlice(c, m["labels"])
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
	if v, err := expandWorkflowTemplateJobsHadoopJobPropertiesSlice(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Properties = flattenWorkflowTemplateJobsHadoopJobPropertiesSlice(c, m["properties"])
	r.LoggingConfig = flattenWorkflowTemplateJobsHadoopJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsHadoopJobPropertiesMap expands the contents of WorkflowTemplateJobsHadoopJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobPropertiesMap(c *Client, f map[string]WorkflowTemplateJobsHadoopJobProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsHadoopJobProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsHadoopJobPropertiesSlice expands the contents of WorkflowTemplateJobsHadoopJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobPropertiesSlice(c *Client, f []WorkflowTemplateJobsHadoopJobProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsHadoopJobProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsHadoopJobPropertiesMap flattens the contents of WorkflowTemplateJobsHadoopJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobPropertiesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsHadoopJobProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsHadoopJobProperties{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsHadoopJobProperties{}
	}

	items := make(map[string]WorkflowTemplateJobsHadoopJobProperties)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsHadoopJobProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsHadoopJobPropertiesSlice flattens the contents of WorkflowTemplateJobsHadoopJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobPropertiesSlice(c *Client, i interface{}) []WorkflowTemplateJobsHadoopJobProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsHadoopJobProperties{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsHadoopJobProperties{}
	}

	items := make([]WorkflowTemplateJobsHadoopJobProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsHadoopJobProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsHadoopJobProperties expands an instance of WorkflowTemplateJobsHadoopJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobProperties(c *Client, f *WorkflowTemplateJobsHadoopJobProperties) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsHadoopJobProperties flattens an instance of WorkflowTemplateJobsHadoopJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobProperties(c *Client, i interface{}) *WorkflowTemplateJobsHadoopJobProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsHadoopJobProperties{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsSlice(c, f.DriverLogLevels); err != nil {
		return nil, fmt.Errorf("error expanding DriverLogLevels into driverLogLevels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.DriverLogLevels = flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsSlice(c, m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsMap expands the contents of WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsMap(c *Client, f map[string]WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsSlice expands the contents of WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsSlice(c *Client, f []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsMap flattens the contents of WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels{}
	}

	items := make(map[string]WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsSlice flattens the contents of WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsSlice(c *Client, i interface{}) []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels{}
	}

	items := make([]WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels expands an instance of WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c *Client, f *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels flattens an instance of WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(c *Client, i interface{}) *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(m["value"])

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
	if v, err := expandWorkflowTemplateJobsSparkJobPropertiesSlice(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Properties = flattenWorkflowTemplateJobsSparkJobPropertiesSlice(c, m["properties"])
	r.LoggingConfig = flattenWorkflowTemplateJobsSparkJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsSparkJobPropertiesMap expands the contents of WorkflowTemplateJobsSparkJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobPropertiesMap(c *Client, f map[string]WorkflowTemplateJobsSparkJobProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkJobProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkJobPropertiesSlice expands the contents of WorkflowTemplateJobsSparkJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobPropertiesSlice(c *Client, f []WorkflowTemplateJobsSparkJobProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkJobProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkJobPropertiesMap flattens the contents of WorkflowTemplateJobsSparkJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobPropertiesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkJobProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkJobProperties{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkJobProperties{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkJobProperties)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkJobProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkJobPropertiesSlice flattens the contents of WorkflowTemplateJobsSparkJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobPropertiesSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkJobProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkJobProperties{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkJobProperties{}
	}

	items := make([]WorkflowTemplateJobsSparkJobProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkJobProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkJobProperties expands an instance of WorkflowTemplateJobsSparkJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobProperties(c *Client, f *WorkflowTemplateJobsSparkJobProperties) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsSparkJobProperties flattens an instance of WorkflowTemplateJobsSparkJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobProperties(c *Client, i interface{}) *WorkflowTemplateJobsSparkJobProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkJobProperties{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsSlice(c, f.DriverLogLevels); err != nil {
		return nil, fmt.Errorf("error expanding DriverLogLevels into driverLogLevels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.DriverLogLevels = flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsSlice(c, m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsMap expands the contents of WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsMap(c *Client, f map[string]WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsSlice expands the contents of WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsSlice(c *Client, f []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsMap flattens the contents of WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsSlice flattens the contents of WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels{}
	}

	items := make([]WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels expands an instance of WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c *Client, f *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels flattens an instance of WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(c *Client, i interface{}) *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(m["value"])

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
	if v, err := expandWorkflowTemplateJobsPysparkJobPropertiesSlice(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Properties = flattenWorkflowTemplateJobsPysparkJobPropertiesSlice(c, m["properties"])
	r.LoggingConfig = flattenWorkflowTemplateJobsPysparkJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsPysparkJobPropertiesMap expands the contents of WorkflowTemplateJobsPysparkJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobPropertiesMap(c *Client, f map[string]WorkflowTemplateJobsPysparkJobProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPysparkJobProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPysparkJobPropertiesSlice expands the contents of WorkflowTemplateJobsPysparkJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobPropertiesSlice(c *Client, f []WorkflowTemplateJobsPysparkJobProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPysparkJobProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPysparkJobPropertiesMap flattens the contents of WorkflowTemplateJobsPysparkJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobPropertiesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPysparkJobProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPysparkJobProperties{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPysparkJobProperties{}
	}

	items := make(map[string]WorkflowTemplateJobsPysparkJobProperties)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPysparkJobProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPysparkJobPropertiesSlice flattens the contents of WorkflowTemplateJobsPysparkJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobPropertiesSlice(c *Client, i interface{}) []WorkflowTemplateJobsPysparkJobProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPysparkJobProperties{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPysparkJobProperties{}
	}

	items := make([]WorkflowTemplateJobsPysparkJobProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPysparkJobProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPysparkJobProperties expands an instance of WorkflowTemplateJobsPysparkJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobProperties(c *Client, f *WorkflowTemplateJobsPysparkJobProperties) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsPysparkJobProperties flattens an instance of WorkflowTemplateJobsPysparkJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobProperties(c *Client, i interface{}) *WorkflowTemplateJobsPysparkJobProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPysparkJobProperties{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsSlice(c, f.DriverLogLevels); err != nil {
		return nil, fmt.Errorf("error expanding DriverLogLevels into driverLogLevels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.DriverLogLevels = flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsSlice(c, m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsMap expands the contents of WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsMap(c *Client, f map[string]WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsSlice expands the contents of WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsSlice(c *Client, f []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsMap flattens the contents of WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels{}
	}

	items := make(map[string]WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsSlice flattens the contents of WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsSlice(c *Client, i interface{}) []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels{}
	}

	items := make([]WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels expands an instance of WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c *Client, f *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels flattens an instance of WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(c *Client, i interface{}) *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(m["value"])

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
	if v, err := expandWorkflowTemplateJobsHiveJobScriptVariablesSlice(c, f.ScriptVariables); err != nil {
		return nil, fmt.Errorf("error expanding ScriptVariables into scriptVariables: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scriptVariables"] = v
	}
	if v, err := expandWorkflowTemplateJobsHiveJobPropertiesSlice(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.ScriptVariables = flattenWorkflowTemplateJobsHiveJobScriptVariablesSlice(c, m["scriptVariables"])
	r.Properties = flattenWorkflowTemplateJobsHiveJobPropertiesSlice(c, m["properties"])
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

// expandWorkflowTemplateJobsHiveJobScriptVariablesMap expands the contents of WorkflowTemplateJobsHiveJobScriptVariables into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobScriptVariablesMap(c *Client, f map[string]WorkflowTemplateJobsHiveJobScriptVariables) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsHiveJobScriptVariables(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsHiveJobScriptVariablesSlice expands the contents of WorkflowTemplateJobsHiveJobScriptVariables into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobScriptVariablesSlice(c *Client, f []WorkflowTemplateJobsHiveJobScriptVariables) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsHiveJobScriptVariables(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsHiveJobScriptVariablesMap flattens the contents of WorkflowTemplateJobsHiveJobScriptVariables from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobScriptVariablesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsHiveJobScriptVariables {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsHiveJobScriptVariables{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsHiveJobScriptVariables{}
	}

	items := make(map[string]WorkflowTemplateJobsHiveJobScriptVariables)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsHiveJobScriptVariables(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsHiveJobScriptVariablesSlice flattens the contents of WorkflowTemplateJobsHiveJobScriptVariables from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobScriptVariablesSlice(c *Client, i interface{}) []WorkflowTemplateJobsHiveJobScriptVariables {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsHiveJobScriptVariables{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsHiveJobScriptVariables{}
	}

	items := make([]WorkflowTemplateJobsHiveJobScriptVariables, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsHiveJobScriptVariables(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsHiveJobScriptVariables expands an instance of WorkflowTemplateJobsHiveJobScriptVariables into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobScriptVariables(c *Client, f *WorkflowTemplateJobsHiveJobScriptVariables) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsHiveJobScriptVariables flattens an instance of WorkflowTemplateJobsHiveJobScriptVariables from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobScriptVariables(c *Client, i interface{}) *WorkflowTemplateJobsHiveJobScriptVariables {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsHiveJobScriptVariables{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandWorkflowTemplateJobsHiveJobPropertiesMap expands the contents of WorkflowTemplateJobsHiveJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobPropertiesMap(c *Client, f map[string]WorkflowTemplateJobsHiveJobProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsHiveJobProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsHiveJobPropertiesSlice expands the contents of WorkflowTemplateJobsHiveJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobPropertiesSlice(c *Client, f []WorkflowTemplateJobsHiveJobProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsHiveJobProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsHiveJobPropertiesMap flattens the contents of WorkflowTemplateJobsHiveJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobPropertiesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsHiveJobProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsHiveJobProperties{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsHiveJobProperties{}
	}

	items := make(map[string]WorkflowTemplateJobsHiveJobProperties)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsHiveJobProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsHiveJobPropertiesSlice flattens the contents of WorkflowTemplateJobsHiveJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobPropertiesSlice(c *Client, i interface{}) []WorkflowTemplateJobsHiveJobProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsHiveJobProperties{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsHiveJobProperties{}
	}

	items := make([]WorkflowTemplateJobsHiveJobProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsHiveJobProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsHiveJobProperties expands an instance of WorkflowTemplateJobsHiveJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsHiveJobProperties(c *Client, f *WorkflowTemplateJobsHiveJobProperties) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsHiveJobProperties flattens an instance of WorkflowTemplateJobsHiveJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsHiveJobProperties(c *Client, i interface{}) *WorkflowTemplateJobsHiveJobProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsHiveJobProperties{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplateJobsPigJobScriptVariablesSlice(c, f.ScriptVariables); err != nil {
		return nil, fmt.Errorf("error expanding ScriptVariables into scriptVariables: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scriptVariables"] = v
	}
	if v, err := expandWorkflowTemplateJobsPigJobPropertiesSlice(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.ScriptVariables = flattenWorkflowTemplateJobsPigJobScriptVariablesSlice(c, m["scriptVariables"])
	r.Properties = flattenWorkflowTemplateJobsPigJobPropertiesSlice(c, m["properties"])
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

// expandWorkflowTemplateJobsPigJobScriptVariablesMap expands the contents of WorkflowTemplateJobsPigJobScriptVariables into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobScriptVariablesMap(c *Client, f map[string]WorkflowTemplateJobsPigJobScriptVariables) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobScriptVariables(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPigJobScriptVariablesSlice expands the contents of WorkflowTemplateJobsPigJobScriptVariables into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobScriptVariablesSlice(c *Client, f []WorkflowTemplateJobsPigJobScriptVariables) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobScriptVariables(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPigJobScriptVariablesMap flattens the contents of WorkflowTemplateJobsPigJobScriptVariables from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobScriptVariablesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPigJobScriptVariables {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPigJobScriptVariables{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPigJobScriptVariables{}
	}

	items := make(map[string]WorkflowTemplateJobsPigJobScriptVariables)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPigJobScriptVariables(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPigJobScriptVariablesSlice flattens the contents of WorkflowTemplateJobsPigJobScriptVariables from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobScriptVariablesSlice(c *Client, i interface{}) []WorkflowTemplateJobsPigJobScriptVariables {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPigJobScriptVariables{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPigJobScriptVariables{}
	}

	items := make([]WorkflowTemplateJobsPigJobScriptVariables, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPigJobScriptVariables(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPigJobScriptVariables expands an instance of WorkflowTemplateJobsPigJobScriptVariables into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobScriptVariables(c *Client, f *WorkflowTemplateJobsPigJobScriptVariables) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsPigJobScriptVariables flattens an instance of WorkflowTemplateJobsPigJobScriptVariables from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobScriptVariables(c *Client, i interface{}) *WorkflowTemplateJobsPigJobScriptVariables {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPigJobScriptVariables{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandWorkflowTemplateJobsPigJobPropertiesMap expands the contents of WorkflowTemplateJobsPigJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobPropertiesMap(c *Client, f map[string]WorkflowTemplateJobsPigJobProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPigJobPropertiesSlice expands the contents of WorkflowTemplateJobsPigJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobPropertiesSlice(c *Client, f []WorkflowTemplateJobsPigJobProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPigJobPropertiesMap flattens the contents of WorkflowTemplateJobsPigJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobPropertiesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPigJobProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPigJobProperties{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPigJobProperties{}
	}

	items := make(map[string]WorkflowTemplateJobsPigJobProperties)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPigJobProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPigJobPropertiesSlice flattens the contents of WorkflowTemplateJobsPigJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobPropertiesSlice(c *Client, i interface{}) []WorkflowTemplateJobsPigJobProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPigJobProperties{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPigJobProperties{}
	}

	items := make([]WorkflowTemplateJobsPigJobProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPigJobProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPigJobProperties expands an instance of WorkflowTemplateJobsPigJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobProperties(c *Client, f *WorkflowTemplateJobsPigJobProperties) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsPigJobProperties flattens an instance of WorkflowTemplateJobsPigJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobProperties(c *Client, i interface{}) *WorkflowTemplateJobsPigJobProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPigJobProperties{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsSlice(c, f.DriverLogLevels); err != nil {
		return nil, fmt.Errorf("error expanding DriverLogLevels into driverLogLevels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.DriverLogLevels = flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsSlice(c, m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsMap expands the contents of WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsMap(c *Client, f map[string]WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsSlice expands the contents of WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsSlice(c *Client, f []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsMap flattens the contents of WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels{}
	}

	items := make(map[string]WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsSlice flattens the contents of WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsSlice(c *Client, i interface{}) []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels{}
	}

	items := make([]WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels expands an instance of WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c *Client, f *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels flattens an instance of WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(c *Client, i interface{}) *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(m["value"])

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
	if v, err := expandWorkflowTemplateJobsSparkRJobPropertiesSlice(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Properties = flattenWorkflowTemplateJobsSparkRJobPropertiesSlice(c, m["properties"])
	r.LoggingConfig = flattenWorkflowTemplateJobsSparkRJobLoggingConfig(c, m["loggingConfig"])

	return r
}

// expandWorkflowTemplateJobsSparkRJobPropertiesMap expands the contents of WorkflowTemplateJobsSparkRJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobPropertiesMap(c *Client, f map[string]WorkflowTemplateJobsSparkRJobProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkRJobProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkRJobPropertiesSlice expands the contents of WorkflowTemplateJobsSparkRJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobPropertiesSlice(c *Client, f []WorkflowTemplateJobsSparkRJobProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkRJobProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkRJobPropertiesMap flattens the contents of WorkflowTemplateJobsSparkRJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobPropertiesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkRJobProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkRJobProperties{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkRJobProperties{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkRJobProperties)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkRJobProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkRJobPropertiesSlice flattens the contents of WorkflowTemplateJobsSparkRJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobPropertiesSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkRJobProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkRJobProperties{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkRJobProperties{}
	}

	items := make([]WorkflowTemplateJobsSparkRJobProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkRJobProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkRJobProperties expands an instance of WorkflowTemplateJobsSparkRJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobProperties(c *Client, f *WorkflowTemplateJobsSparkRJobProperties) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsSparkRJobProperties flattens an instance of WorkflowTemplateJobsSparkRJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobProperties(c *Client, i interface{}) *WorkflowTemplateJobsSparkRJobProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkRJobProperties{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsSlice(c, f.DriverLogLevels); err != nil {
		return nil, fmt.Errorf("error expanding DriverLogLevels into driverLogLevels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.DriverLogLevels = flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsSlice(c, m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsMap expands the contents of WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsMap(c *Client, f map[string]WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsSlice expands the contents of WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsSlice(c *Client, f []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsMap flattens the contents of WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsSlice flattens the contents of WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels{}
	}

	items := make([]WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels expands an instance of WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c *Client, f *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels flattens an instance of WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(c *Client, i interface{}) *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(m["value"])

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
	if v, err := expandWorkflowTemplateJobsSparkSqlJobScriptVariablesSlice(c, f.ScriptVariables); err != nil {
		return nil, fmt.Errorf("error expanding ScriptVariables into scriptVariables: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scriptVariables"] = v
	}
	if v, err := expandWorkflowTemplateJobsSparkSqlJobPropertiesSlice(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.ScriptVariables = flattenWorkflowTemplateJobsSparkSqlJobScriptVariablesSlice(c, m["scriptVariables"])
	r.Properties = flattenWorkflowTemplateJobsSparkSqlJobPropertiesSlice(c, m["properties"])
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

// expandWorkflowTemplateJobsSparkSqlJobScriptVariablesMap expands the contents of WorkflowTemplateJobsSparkSqlJobScriptVariables into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobScriptVariablesMap(c *Client, f map[string]WorkflowTemplateJobsSparkSqlJobScriptVariables) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobScriptVariables(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkSqlJobScriptVariablesSlice expands the contents of WorkflowTemplateJobsSparkSqlJobScriptVariables into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobScriptVariablesSlice(c *Client, f []WorkflowTemplateJobsSparkSqlJobScriptVariables) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobScriptVariables(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkSqlJobScriptVariablesMap flattens the contents of WorkflowTemplateJobsSparkSqlJobScriptVariables from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobScriptVariablesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkSqlJobScriptVariables {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkSqlJobScriptVariables{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkSqlJobScriptVariables{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkSqlJobScriptVariables)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkSqlJobScriptVariables(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkSqlJobScriptVariablesSlice flattens the contents of WorkflowTemplateJobsSparkSqlJobScriptVariables from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobScriptVariablesSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkSqlJobScriptVariables {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkSqlJobScriptVariables{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkSqlJobScriptVariables{}
	}

	items := make([]WorkflowTemplateJobsSparkSqlJobScriptVariables, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkSqlJobScriptVariables(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkSqlJobScriptVariables expands an instance of WorkflowTemplateJobsSparkSqlJobScriptVariables into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobScriptVariables(c *Client, f *WorkflowTemplateJobsSparkSqlJobScriptVariables) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsSparkSqlJobScriptVariables flattens an instance of WorkflowTemplateJobsSparkSqlJobScriptVariables from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobScriptVariables(c *Client, i interface{}) *WorkflowTemplateJobsSparkSqlJobScriptVariables {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkSqlJobScriptVariables{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandWorkflowTemplateJobsSparkSqlJobPropertiesMap expands the contents of WorkflowTemplateJobsSparkSqlJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobPropertiesMap(c *Client, f map[string]WorkflowTemplateJobsSparkSqlJobProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkSqlJobPropertiesSlice expands the contents of WorkflowTemplateJobsSparkSqlJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobPropertiesSlice(c *Client, f []WorkflowTemplateJobsSparkSqlJobProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkSqlJobPropertiesMap flattens the contents of WorkflowTemplateJobsSparkSqlJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobPropertiesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkSqlJobProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkSqlJobProperties{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkSqlJobProperties{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkSqlJobProperties)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkSqlJobProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkSqlJobPropertiesSlice flattens the contents of WorkflowTemplateJobsSparkSqlJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobPropertiesSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkSqlJobProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkSqlJobProperties{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkSqlJobProperties{}
	}

	items := make([]WorkflowTemplateJobsSparkSqlJobProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkSqlJobProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkSqlJobProperties expands an instance of WorkflowTemplateJobsSparkSqlJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobProperties(c *Client, f *WorkflowTemplateJobsSparkSqlJobProperties) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsSparkSqlJobProperties flattens an instance of WorkflowTemplateJobsSparkSqlJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobProperties(c *Client, i interface{}) *WorkflowTemplateJobsSparkSqlJobProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkSqlJobProperties{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsSlice(c, f.DriverLogLevels); err != nil {
		return nil, fmt.Errorf("error expanding DriverLogLevels into driverLogLevels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.DriverLogLevels = flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsSlice(c, m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsMap expands the contents of WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsMap(c *Client, f map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsSlice expands the contents of WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsSlice(c *Client, f []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsMap flattens the contents of WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels{}
	}

	items := make(map[string]WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsSlice flattens the contents of WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels{}
	}

	items := make([]WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels expands an instance of WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c *Client, f *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels flattens an instance of WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(c *Client, i interface{}) *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(m["value"])

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
	if v, err := expandWorkflowTemplateJobsPrestoJobPropertiesSlice(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.Properties = flattenWorkflowTemplateJobsPrestoJobPropertiesSlice(c, m["properties"])
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

// expandWorkflowTemplateJobsPrestoJobPropertiesMap expands the contents of WorkflowTemplateJobsPrestoJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobPropertiesMap(c *Client, f map[string]WorkflowTemplateJobsPrestoJobProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJobProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPrestoJobPropertiesSlice expands the contents of WorkflowTemplateJobsPrestoJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobPropertiesSlice(c *Client, f []WorkflowTemplateJobsPrestoJobProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJobProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPrestoJobPropertiesMap flattens the contents of WorkflowTemplateJobsPrestoJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobPropertiesMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPrestoJobProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPrestoJobProperties{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPrestoJobProperties{}
	}

	items := make(map[string]WorkflowTemplateJobsPrestoJobProperties)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPrestoJobProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPrestoJobPropertiesSlice flattens the contents of WorkflowTemplateJobsPrestoJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobPropertiesSlice(c *Client, i interface{}) []WorkflowTemplateJobsPrestoJobProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPrestoJobProperties{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPrestoJobProperties{}
	}

	items := make([]WorkflowTemplateJobsPrestoJobProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPrestoJobProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPrestoJobProperties expands an instance of WorkflowTemplateJobsPrestoJobProperties into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobProperties(c *Client, f *WorkflowTemplateJobsPrestoJobProperties) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsPrestoJobProperties flattens an instance of WorkflowTemplateJobsPrestoJobProperties from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobProperties(c *Client, i interface{}) *WorkflowTemplateJobsPrestoJobProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPrestoJobProperties{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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
	if v, err := expandWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsSlice(c, f.DriverLogLevels); err != nil {
		return nil, fmt.Errorf("error expanding DriverLogLevels into driverLogLevels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	r.DriverLogLevels = flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsSlice(c, m["driverLogLevels"])

	return r
}

// expandWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsMap expands the contents of WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsMap(c *Client, f map[string]WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsSlice expands the contents of WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsSlice(c *Client, f []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsMap flattens the contents of WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels{}
	}

	items := make(map[string]WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsSlice flattens the contents of WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsSlice(c *Client, i interface{}) []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels{}
	}

	items := make([]WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels expands an instance of WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels into a JSON
// request object.
func expandWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c *Client, f *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels flattens an instance of WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(c *Client, i interface{}) *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(m["value"])

	return r
}

// expandWorkflowTemplateJobsLabelsMap expands the contents of WorkflowTemplateJobsLabels into a JSON
// request object.
func expandWorkflowTemplateJobsLabelsMap(c *Client, f map[string]WorkflowTemplateJobsLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkflowTemplateJobsLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkflowTemplateJobsLabelsSlice expands the contents of WorkflowTemplateJobsLabels into a JSON
// request object.
func expandWorkflowTemplateJobsLabelsSlice(c *Client, f []WorkflowTemplateJobsLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkflowTemplateJobsLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkflowTemplateJobsLabelsMap flattens the contents of WorkflowTemplateJobsLabels from a JSON
// response object.
func flattenWorkflowTemplateJobsLabelsMap(c *Client, i interface{}) map[string]WorkflowTemplateJobsLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkflowTemplateJobsLabels{}
	}

	if len(a) == 0 {
		return map[string]WorkflowTemplateJobsLabels{}
	}

	items := make(map[string]WorkflowTemplateJobsLabels)
	for k, item := range a {
		items[k] = *flattenWorkflowTemplateJobsLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkflowTemplateJobsLabelsSlice flattens the contents of WorkflowTemplateJobsLabels from a JSON
// response object.
func flattenWorkflowTemplateJobsLabelsSlice(c *Client, i interface{}) []WorkflowTemplateJobsLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsLabels{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsLabels{}
	}

	items := make([]WorkflowTemplateJobsLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkflowTemplateJobsLabels expands an instance of WorkflowTemplateJobsLabels into a JSON
// request object.
func expandWorkflowTemplateJobsLabels(c *Client, f *WorkflowTemplateJobsLabels) (map[string]interface{}, error) {
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

// flattenWorkflowTemplateJobsLabels flattens an instance of WorkflowTemplateJobsLabels from a JSON
// response object.
func flattenWorkflowTemplateJobsLabels(c *Client, i interface{}) *WorkflowTemplateJobsLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkflowTemplateJobsLabels{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

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

// flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumSlice flattens the contents of WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum from a JSON
// response object.
func flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, i interface{}) []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	items := make([]WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum asserts that an interface is a string, and returns a
// pointer to a *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum with the same value as that string.
func flattenWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(i interface{}) *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum {
	s, ok := i.(string)
	if !ok {
		return WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumRef("")
	}

	return WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumRef(s)
}

// flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumSlice flattens the contents of WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	items := make([]WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum asserts that an interface is a string, and returns a
// pointer to a *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum with the same value as that string.
func flattenWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(i interface{}) *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum {
	s, ok := i.(string)
	if !ok {
		return WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumRef("")
	}

	return WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumRef(s)
}

// flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumSlice flattens the contents of WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum from a JSON
// response object.
func flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, i interface{}) []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	items := make([]WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum asserts that an interface is a string, and returns a
// pointer to a *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum with the same value as that string.
func flattenWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(i interface{}) *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum {
	s, ok := i.(string)
	if !ok {
		return WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumRef("")
	}

	return WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumRef(s)
}

// flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumSlice flattens the contents of WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum from a JSON
// response object.
func flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, i interface{}) []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	items := make([]WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum asserts that an interface is a string, and returns a
// pointer to a *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum with the same value as that string.
func flattenWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(i interface{}) *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum {
	s, ok := i.(string)
	if !ok {
		return WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumRef("")
	}

	return WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumRef(s)
}

// flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumSlice flattens the contents of WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	items := make([]WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum asserts that an interface is a string, and returns a
// pointer to a *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum with the same value as that string.
func flattenWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(i interface{}) *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum {
	s, ok := i.(string)
	if !ok {
		return WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumRef("")
	}

	return WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumRef(s)
}

// flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumSlice flattens the contents of WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum from a JSON
// response object.
func flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, i interface{}) []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	items := make([]WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum asserts that an interface is a string, and returns a
// pointer to a *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum with the same value as that string.
func flattenWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(i interface{}) *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum {
	s, ok := i.(string)
	if !ok {
		return WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumRef("")
	}

	return WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumRef(s)
}

// flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumSlice flattens the contents of WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum from a JSON
// response object.
func flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumSlice(c *Client, i interface{}) []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	if len(a) == 0 {
		return []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum{}
	}

	items := make([]WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum asserts that an interface is a string, and returns a
// pointer to a *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum with the same value as that string.
func flattenWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(i interface{}) *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum {
	s, ok := i.(string)
	if !ok {
		return WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumRef("")
	}

	return WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumRef(s)
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
