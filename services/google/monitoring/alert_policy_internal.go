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
package monitoring

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *AlertPolicy) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Documentation) {
		if err := r.Documentation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Enabled) {
		if err := r.Enabled.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Validity) {
		if err := r.Validity.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CreationRecord) {
		if err := r.CreationRecord.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MutationRecord) {
		if err := r.MutationRecord.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.IncidentStrategy) {
		if err := r.IncidentStrategy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Metadata) {
		if err := r.Metadata.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyDocumentation) validate() error {
	return nil
}
func (r *AlertPolicyConditions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ConditionThreshold) {
		if err := r.ConditionThreshold.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConditionAbsent) {
		if err := r.ConditionAbsent.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConditionMatchedLog) {
		if err := r.ConditionMatchedLog.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConditionClusterOutlier) {
		if err := r.ConditionClusterOutlier.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConditionRate) {
		if err := r.ConditionRate.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConditionUpMon) {
		if err := r.ConditionUpMon.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConditionProcessCount) {
		if err := r.ConditionProcessCount.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConditionTimeSeriesQueryLanguage) {
		if err := r.ConditionTimeSeriesQueryLanguage.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConditionMonitoringQueryLanguage) {
		if err := r.ConditionMonitoringQueryLanguage.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionThreshold) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Trigger) {
		if err := r.Trigger.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdAggregations) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ReduceFractionLessThanParams) {
		if err := r.ReduceFractionLessThanParams.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReduceMakeDistributionParams) {
		if err := r.ReduceMakeDistributionParams.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BucketOptions) {
		if err := r.BucketOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExemplarSampling) {
		if err := r.ExemplarSampling.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LinearBuckets) {
		if err := r.LinearBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExponentialBuckets) {
		if err := r.ExponentialBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExplicitBuckets) {
		if err := r.ExplicitBuckets.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregations) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ReduceFractionLessThanParams) {
		if err := r.ReduceFractionLessThanParams.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReduceMakeDistributionParams) {
		if err := r.ReduceMakeDistributionParams.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BucketOptions) {
		if err := r.BucketOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExemplarSampling) {
		if err := r.ExemplarSampling.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LinearBuckets) {
		if err := r.LinearBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExponentialBuckets) {
		if err := r.ExponentialBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExplicitBuckets) {
		if err := r.ExplicitBuckets.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionThresholdTrigger) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionAbsent) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Duration) {
		if err := r.Duration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Trigger) {
		if err := r.Trigger.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentAggregations) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ReduceFractionLessThanParams) {
		if err := r.ReduceFractionLessThanParams.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReduceMakeDistributionParams) {
		if err := r.ReduceMakeDistributionParams.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BucketOptions) {
		if err := r.BucketOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExemplarSampling) {
		if err := r.ExemplarSampling.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LinearBuckets) {
		if err := r.LinearBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExponentialBuckets) {
		if err := r.ExponentialBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExplicitBuckets) {
		if err := r.ExplicitBuckets.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentDuration) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionAbsentTrigger) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionMatchedLog) validate() error {
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	return nil
}
func (r *AlertPolicyConditionsConditionClusterOutlier) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionRate) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TimeWindow) {
		if err := r.TimeWindow.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Trigger) {
		if err := r.Trigger.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionRateAggregations) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ReduceFractionLessThanParams) {
		if err := r.ReduceFractionLessThanParams.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReduceMakeDistributionParams) {
		if err := r.ReduceMakeDistributionParams.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BucketOptions) {
		if err := r.BucketOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExemplarSampling) {
		if err := r.ExemplarSampling.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LinearBuckets) {
		if err := r.LinearBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExponentialBuckets) {
		if err := r.ExponentialBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExplicitBuckets) {
		if err := r.ExplicitBuckets.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionRateTimeWindow) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionRateTrigger) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionUpMon) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Duration) {
		if err := r.Duration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Trigger) {
		if err := r.Trigger.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionUpMonDuration) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionUpMonTrigger) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionProcessCount) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Trigger) {
		if err := r.Trigger.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Duration) {
		if err := r.Duration.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionProcessCountTrigger) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionProcessCountDuration) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionTimeSeriesQueryLanguage) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionMonitoringQueryLanguage) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Duration) {
		if err := r.Duration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Trigger) {
		if err := r.Trigger.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) validate() error {
	return nil
}
func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) validate() error {
	return nil
}
func (r *AlertPolicyEnabled) validate() error {
	return nil
}
func (r *AlertPolicyValidity) validate() error {
	return nil
}
func (r *AlertPolicyValidityDetails) validate() error {
	return nil
}
func (r *AlertPolicyCreationRecord) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MutateTime) {
		if err := r.MutateTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyCreationRecordMutateTime) validate() error {
	return nil
}
func (r *AlertPolicyMutationRecord) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MutateTime) {
		if err := r.MutateTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AlertPolicyMutationRecordMutateTime) validate() error {
	return nil
}
func (r *AlertPolicyIncidentStrategy) validate() error {
	return nil
}
func (r *AlertPolicyMetadata) validate() error {
	return nil
}

func alertPolicyGetURL(userBasePath string, r *AlertPolicy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v3/projects/{{project}}/alertPolicies/{{name}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

func alertPolicyListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/alertPolicies", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func alertPolicyCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/alertPolicies", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func alertPolicyDeleteURL(userBasePath string, r *AlertPolicy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v3/projects/{{project}}/alertPolicies/{{name}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

// alertPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type alertPolicyApiOperation interface {
	do(context.Context, *AlertPolicy, *Client) error
}

// newUpdateAlertPolicyUpdateAlertPolicyRequest creates a request for an
// AlertPolicy resource's UpdateAlertPolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAlertPolicyUpdateAlertPolicyRequest(ctx context.Context, f *AlertPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateAlertPolicyUpdateAlertPolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateAlertPolicyUpdateAlertPolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAlertPolicyUpdateAlertPolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAlertPolicyUpdateAlertPolicyOperation) do(ctx context.Context, r *AlertPolicy, c *Client) error {
	_, err := c.GetAlertPolicy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateAlertPolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateAlertPolicyUpdateAlertPolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateAlertPolicyUpdateAlertPolicyRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://monitoring.googleapis.com/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listAlertPolicyRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := alertPolicyListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AlertPolicyMaxPage {
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

type listAlertPolicyOperation struct {
	AlertPolicies []map[string]interface{} `json:"alertPolicies"`
	Token         string                   `json:"nextPageToken"`
}

func (c *Client) listAlertPolicy(ctx context.Context, project, pageToken string, pageSize int32) ([]*AlertPolicy, string, error) {
	b, err := c.listAlertPolicyRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAlertPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AlertPolicy
	for _, v := range m.AlertPolicies {
		res, err := unmarshalMapAlertPolicy(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAlertPolicy(ctx context.Context, f func(*AlertPolicy) bool, resources []*AlertPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAlertPolicy(ctx, res)
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

type deleteAlertPolicyOperation struct{}

func (op *deleteAlertPolicyOperation) do(ctx context.Context, r *AlertPolicy, c *Client) error {

	_, err := c.GetAlertPolicy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("AlertPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAlertPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := alertPolicyDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete AlertPolicy: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAlertPolicy(ctx, r.urlNormalized())
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
type createAlertPolicyOperation struct {
	response map[string]interface{}
}

func (op *createAlertPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAlertPolicyOperation) do(ctx context.Context, r *AlertPolicy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := alertPolicyCreateURL(c.Config.BasePath, project)

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

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string, was %T", name)
	}
	r.Name = &name

	if _, err := c.GetAlertPolicy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAlertPolicyRaw(ctx context.Context, r *AlertPolicy) ([]byte, error) {

	u, err := alertPolicyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) alertPolicyDiffsForRawDesired(ctx context.Context, rawDesired *AlertPolicy, opts ...dcl.ApplyOption) (initial, desired *AlertPolicy, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AlertPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AlertPolicy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected AlertPolicy, got %T", sh)
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
		desired, err := canonicalizeAlertPolicyDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAlertPolicy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a AlertPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AlertPolicy resource: %v", err)
		}
		c.Config.Logger.Info("Found that AlertPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAlertPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for AlertPolicy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for AlertPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAlertPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for AlertPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAlertPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for AlertPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAlertPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAlertPolicyInitialState(rawInitial, rawDesired *AlertPolicy) (*AlertPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAlertPolicyDesiredState(rawDesired, rawInitial *AlertPolicy, opts ...dcl.ApplyOption) (*AlertPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Documentation = canonicalizeAlertPolicyDocumentation(rawDesired.Documentation, nil, opts...)
		rawDesired.Enabled = canonicalizeAlertPolicyEnabled(rawDesired.Enabled, nil, opts...)
		rawDesired.Validity = canonicalizeAlertPolicyValidity(rawDesired.Validity, nil, opts...)
		rawDesired.CreationRecord = canonicalizeAlertPolicyCreationRecord(rawDesired.CreationRecord, nil, opts...)
		rawDesired.MutationRecord = canonicalizeAlertPolicyMutationRecord(rawDesired.MutationRecord, nil, opts...)
		rawDesired.IncidentStrategy = canonicalizeAlertPolicyIncidentStrategy(rawDesired.IncidentStrategy, nil, opts...)
		rawDesired.Metadata = canonicalizeAlertPolicyMetadata(rawDesired.Metadata, nil, opts...)

		return rawDesired, nil
	}

	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	rawDesired.Documentation = canonicalizeAlertPolicyDocumentation(rawDesired.Documentation, rawInitial.Documentation, opts...)
	if dcl.IsZeroValue(rawDesired.UserLabels) {
		rawDesired.UserLabels = rawInitial.UserLabels
	}
	if dcl.IsZeroValue(rawDesired.Conditions) {
		rawDesired.Conditions = rawInitial.Conditions
	}
	if dcl.IsZeroValue(rawDesired.Combiner) {
		rawDesired.Combiner = rawInitial.Combiner
	}
	if dcl.BoolCanonicalize(rawDesired.Disabled, rawInitial.Disabled) {
		rawDesired.Disabled = rawInitial.Disabled
	}
	rawDesired.Enabled = canonicalizeAlertPolicyEnabled(rawDesired.Enabled, rawInitial.Enabled, opts...)
	rawDesired.Validity = canonicalizeAlertPolicyValidity(rawDesired.Validity, rawInitial.Validity, opts...)
	if dcl.IsZeroValue(rawDesired.NotificationChannels) {
		rawDesired.NotificationChannels = rawInitial.NotificationChannels
	}
	rawDesired.CreationRecord = canonicalizeAlertPolicyCreationRecord(rawDesired.CreationRecord, rawInitial.CreationRecord, opts...)
	rawDesired.MutationRecord = canonicalizeAlertPolicyMutationRecord(rawDesired.MutationRecord, rawInitial.MutationRecord, opts...)
	rawDesired.IncidentStrategy = canonicalizeAlertPolicyIncidentStrategy(rawDesired.IncidentStrategy, rawInitial.IncidentStrategy, opts...)
	rawDesired.Metadata = canonicalizeAlertPolicyMetadata(rawDesired.Metadata, rawInitial.Metadata, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeAlertPolicyNewState(c *Client, rawNew, rawDesired *AlertPolicy) (*AlertPolicy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Documentation) && dcl.IsEmptyValueIndirect(rawDesired.Documentation) {
		rawNew.Documentation = rawDesired.Documentation
	} else {
		rawNew.Documentation = canonicalizeNewAlertPolicyDocumentation(c, rawDesired.Documentation, rawNew.Documentation)
	}

	if dcl.IsEmptyValueIndirect(rawNew.UserLabels) && dcl.IsEmptyValueIndirect(rawDesired.UserLabels) {
		rawNew.UserLabels = rawDesired.UserLabels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Conditions) && dcl.IsEmptyValueIndirect(rawDesired.Conditions) {
		rawNew.Conditions = rawDesired.Conditions
	} else {
		rawNew.Conditions = canonicalizeNewAlertPolicyConditionsSlice(c, rawDesired.Conditions, rawNew.Conditions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Combiner) && dcl.IsEmptyValueIndirect(rawDesired.Combiner) {
		rawNew.Combiner = rawDesired.Combiner
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Enabled) && dcl.IsEmptyValueIndirect(rawDesired.Enabled) {
		rawNew.Enabled = rawDesired.Enabled
	} else {
		rawNew.Enabled = canonicalizeNewAlertPolicyEnabled(c, rawDesired.Enabled, rawNew.Enabled)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Validity) && dcl.IsEmptyValueIndirect(rawDesired.Validity) {
		rawNew.Validity = rawDesired.Validity
	} else {
		rawNew.Validity = canonicalizeNewAlertPolicyValidity(c, rawDesired.Validity, rawNew.Validity)
	}

	if dcl.IsEmptyValueIndirect(rawNew.NotificationChannels) && dcl.IsEmptyValueIndirect(rawDesired.NotificationChannels) {
		rawNew.NotificationChannels = rawDesired.NotificationChannels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationRecord) && dcl.IsEmptyValueIndirect(rawDesired.CreationRecord) {
		rawNew.CreationRecord = rawDesired.CreationRecord
	} else {
		rawNew.CreationRecord = canonicalizeNewAlertPolicyCreationRecord(c, rawDesired.CreationRecord, rawNew.CreationRecord)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MutationRecord) && dcl.IsEmptyValueIndirect(rawDesired.MutationRecord) {
		rawNew.MutationRecord = rawDesired.MutationRecord
	} else {
		rawNew.MutationRecord = canonicalizeNewAlertPolicyMutationRecord(c, rawDesired.MutationRecord, rawNew.MutationRecord)
	}

	if dcl.IsEmptyValueIndirect(rawNew.IncidentStrategy) && dcl.IsEmptyValueIndirect(rawDesired.IncidentStrategy) {
		rawNew.IncidentStrategy = rawDesired.IncidentStrategy
	} else {
		rawNew.IncidentStrategy = canonicalizeNewAlertPolicyIncidentStrategy(c, rawDesired.IncidentStrategy, rawNew.IncidentStrategy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Metadata) && dcl.IsEmptyValueIndirect(rawDesired.Metadata) {
		rawNew.Metadata = rawDesired.Metadata
	} else {
		rawNew.Metadata = canonicalizeNewAlertPolicyMetadata(c, rawDesired.Metadata, rawNew.Metadata)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeAlertPolicyDocumentation(des, initial *AlertPolicyDocumentation, opts ...dcl.ApplyOption) *AlertPolicyDocumentation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Content, initial.Content) || dcl.IsZeroValue(des.Content) {
		des.Content = initial.Content
	}
	if dcl.StringCanonicalize(des.MimeType, initial.MimeType) || dcl.IsZeroValue(des.MimeType) {
		des.MimeType = initial.MimeType
	}

	return des
}

func canonicalizeNewAlertPolicyDocumentation(c *Client, des, nw *AlertPolicyDocumentation) *AlertPolicyDocumentation {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}
	if dcl.StringCanonicalize(des.MimeType, nw.MimeType) {
		nw.MimeType = des.MimeType
	}

	return nw
}

func canonicalizeNewAlertPolicyDocumentationSet(c *Client, des, nw []AlertPolicyDocumentation) []AlertPolicyDocumentation {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyDocumentation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyDocumentationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyDocumentationSlice(c *Client, des, nw []AlertPolicyDocumentation) []AlertPolicyDocumentation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyDocumentation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyDocumentation(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditions(des, initial *AlertPolicyConditions, opts ...dcl.ApplyOption) *AlertPolicyConditions {
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
	if dcl.StringCanonicalize(des.DisplayName, initial.DisplayName) || dcl.IsZeroValue(des.DisplayName) {
		des.DisplayName = initial.DisplayName
	}
	if dcl.IsZeroValue(des.ResourceStateFilter) {
		des.ResourceStateFilter = initial.ResourceStateFilter
	}
	des.ConditionThreshold = canonicalizeAlertPolicyConditionsConditionThreshold(des.ConditionThreshold, initial.ConditionThreshold, opts...)
	des.ConditionAbsent = canonicalizeAlertPolicyConditionsConditionAbsent(des.ConditionAbsent, initial.ConditionAbsent, opts...)
	des.ConditionMatchedLog = canonicalizeAlertPolicyConditionsConditionMatchedLog(des.ConditionMatchedLog, initial.ConditionMatchedLog, opts...)
	des.ConditionClusterOutlier = canonicalizeAlertPolicyConditionsConditionClusterOutlier(des.ConditionClusterOutlier, initial.ConditionClusterOutlier, opts...)
	des.ConditionRate = canonicalizeAlertPolicyConditionsConditionRate(des.ConditionRate, initial.ConditionRate, opts...)
	des.ConditionUpMon = canonicalizeAlertPolicyConditionsConditionUpMon(des.ConditionUpMon, initial.ConditionUpMon, opts...)
	des.ConditionProcessCount = canonicalizeAlertPolicyConditionsConditionProcessCount(des.ConditionProcessCount, initial.ConditionProcessCount, opts...)
	des.ConditionTimeSeriesQueryLanguage = canonicalizeAlertPolicyConditionsConditionTimeSeriesQueryLanguage(des.ConditionTimeSeriesQueryLanguage, initial.ConditionTimeSeriesQueryLanguage, opts...)
	des.ConditionMonitoringQueryLanguage = canonicalizeAlertPolicyConditionsConditionMonitoringQueryLanguage(des.ConditionMonitoringQueryLanguage, initial.ConditionMonitoringQueryLanguage, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditions(c *Client, des, nw *AlertPolicyConditions) *AlertPolicyConditions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.DisplayName, nw.DisplayName) {
		nw.DisplayName = des.DisplayName
	}
	if dcl.IsZeroValue(nw.ResourceStateFilter) {
		nw.ResourceStateFilter = des.ResourceStateFilter
	}
	nw.ConditionThreshold = canonicalizeNewAlertPolicyConditionsConditionThreshold(c, des.ConditionThreshold, nw.ConditionThreshold)
	nw.ConditionAbsent = canonicalizeNewAlertPolicyConditionsConditionAbsent(c, des.ConditionAbsent, nw.ConditionAbsent)
	nw.ConditionMatchedLog = canonicalizeNewAlertPolicyConditionsConditionMatchedLog(c, des.ConditionMatchedLog, nw.ConditionMatchedLog)
	nw.ConditionClusterOutlier = canonicalizeNewAlertPolicyConditionsConditionClusterOutlier(c, des.ConditionClusterOutlier, nw.ConditionClusterOutlier)
	nw.ConditionRate = canonicalizeNewAlertPolicyConditionsConditionRate(c, des.ConditionRate, nw.ConditionRate)
	nw.ConditionUpMon = canonicalizeNewAlertPolicyConditionsConditionUpMon(c, des.ConditionUpMon, nw.ConditionUpMon)
	nw.ConditionProcessCount = canonicalizeNewAlertPolicyConditionsConditionProcessCount(c, des.ConditionProcessCount, nw.ConditionProcessCount)
	nw.ConditionTimeSeriesQueryLanguage = canonicalizeNewAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, des.ConditionTimeSeriesQueryLanguage, nw.ConditionTimeSeriesQueryLanguage)
	nw.ConditionMonitoringQueryLanguage = canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguage(c, des.ConditionMonitoringQueryLanguage, nw.ConditionMonitoringQueryLanguage)

	return nw
}

func canonicalizeNewAlertPolicyConditionsSet(c *Client, des, nw []AlertPolicyConditions) []AlertPolicyConditions {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsSlice(c *Client, des, nw []AlertPolicyConditions) []AlertPolicyConditions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditions(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThreshold(des, initial *AlertPolicyConditionsConditionThreshold, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThreshold {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}
	if dcl.IsZeroValue(des.Aggregations) {
		des.Aggregations = initial.Aggregations
	}
	if dcl.StringCanonicalize(des.DenominatorFilter, initial.DenominatorFilter) || dcl.IsZeroValue(des.DenominatorFilter) {
		des.DenominatorFilter = initial.DenominatorFilter
	}
	if dcl.IsZeroValue(des.DenominatorAggregations) {
		des.DenominatorAggregations = initial.DenominatorAggregations
	}
	if dcl.IsZeroValue(des.Comparison) {
		des.Comparison = initial.Comparison
	}
	if dcl.IsZeroValue(des.ThresholdValue) {
		des.ThresholdValue = initial.ThresholdValue
	}
	if dcl.StringCanonicalize(des.Duration, initial.Duration) || dcl.IsZeroValue(des.Duration) {
		des.Duration = initial.Duration
	}
	des.Trigger = canonicalizeAlertPolicyConditionsConditionThresholdTrigger(des.Trigger, initial.Trigger, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThreshold(c *Client, des, nw *AlertPolicyConditionsConditionThreshold) *AlertPolicyConditionsConditionThreshold {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	nw.Aggregations = canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsSlice(c, des.Aggregations, nw.Aggregations)
	if dcl.StringCanonicalize(des.DenominatorFilter, nw.DenominatorFilter) {
		nw.DenominatorFilter = des.DenominatorFilter
	}
	nw.DenominatorAggregations = canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice(c, des.DenominatorAggregations, nw.DenominatorAggregations)
	if dcl.IsZeroValue(nw.Comparison) {
		nw.Comparison = des.Comparison
	}
	if dcl.IsZeroValue(nw.ThresholdValue) {
		nw.ThresholdValue = des.ThresholdValue
	}
	if dcl.StringCanonicalize(des.Duration, nw.Duration) {
		nw.Duration = des.Duration
	}
	nw.Trigger = canonicalizeNewAlertPolicyConditionsConditionThresholdTrigger(c, des.Trigger, nw.Trigger)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdSet(c *Client, des, nw []AlertPolicyConditionsConditionThreshold) []AlertPolicyConditionsConditionThreshold {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThreshold
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdSlice(c *Client, des, nw []AlertPolicyConditionsConditionThreshold) []AlertPolicyConditionsConditionThreshold {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThreshold
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThreshold(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdAggregations(des, initial *AlertPolicyConditionsConditionThresholdAggregations, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		des.AlignmentPeriod = initial.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	}
	if dcl.IsZeroValue(des.GroupByFields) {
		des.GroupByFields = initial.GroupByFields
	}
	des.ReduceFractionLessThanParams = canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(des.ReduceFractionLessThanParams, initial.ReduceFractionLessThanParams, opts...)
	des.ReduceMakeDistributionParams = canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(des.ReduceMakeDistributionParams, initial.ReduceMakeDistributionParams, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregations(c *Client, des, nw *AlertPolicyConditionsConditionThresholdAggregations) *AlertPolicyConditionsConditionThresholdAggregations {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(nw.PerSeriesAligner) {
		nw.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(nw.CrossSeriesReducer) {
		nw.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.IsZeroValue(nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}
	nw.ReduceFractionLessThanParams = canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, des.ReduceFractionLessThanParams, nw.ReduceFractionLessThanParams)
	nw.ReduceMakeDistributionParams = canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, des.ReduceMakeDistributionParams, nw.ReduceMakeDistributionParams)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregations) []AlertPolicyConditionsConditionThresholdAggregations {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdAggregations
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdAggregationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregations) []AlertPolicyConditionsConditionThresholdAggregations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdAggregations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdAggregations(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Threshold) {
		des.Threshold = initial.Threshold
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c *Client, des, nw *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Threshold) {
		nw.Threshold = des.Threshold
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BucketOptions = canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(des.BucketOptions, initial.BucketOptions, opts...)
	des.ExemplarSampling = canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(des.ExemplarSampling, initial.ExemplarSampling, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c *Client, des, nw *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	if des == nil || nw == nil {
		return nw
	}

	nw.BucketOptions = canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, des.BucketOptions, nw.BucketOptions)
	nw.ExemplarSampling = canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, des.ExemplarSampling, nw.ExemplarSampling)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LinearBuckets = canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des.LinearBuckets, initial.LinearBuckets, opts...)
	des.ExponentialBuckets = canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des.ExponentialBuckets, initial.ExponentialBuckets, opts...)
	des.ExplicitBuckets = canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des.ExplicitBuckets, initial.ExplicitBuckets, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, des, nw *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil || nw == nil {
		return nw
	}

	nw.LinearBuckets = canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, des.LinearBuckets, nw.LinearBuckets)
	nw.ExponentialBuckets = canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, des.ExponentialBuckets, nw.ExponentialBuckets)
	nw.ExplicitBuckets = canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, des.ExplicitBuckets, nw.ExplicitBuckets)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.Width) {
		des.Width = initial.Width
	}
	if dcl.IsZeroValue(des.Offset) {
		des.Offset = initial.Offset
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, des, nw *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.Width) {
		nw.Width = des.Width
	}
	if dcl.IsZeroValue(nw.Offset) {
		nw.Offset = des.Offset
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.GrowthFactor) {
		des.GrowthFactor = initial.GrowthFactor
	}
	if dcl.IsZeroValue(des.Scale) {
		des.Scale = initial.Scale
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, des, nw *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.GrowthFactor) {
		nw.GrowthFactor = des.GrowthFactor
	}
	if dcl.IsZeroValue(nw.Scale) {
		nw.Scale = des.Scale
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Bounds) {
		des.Bounds = initial.Bounds
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, des, nw *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Bounds) {
		nw.Bounds = des.Bounds
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MinimumValue) {
		des.MinimumValue = initial.MinimumValue
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, des, nw *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.MinimumValue) {
		nw.MinimumValue = des.MinimumValue
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregations(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregations, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		des.AlignmentPeriod = initial.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	}
	if dcl.IsZeroValue(des.GroupByFields) {
		des.GroupByFields = initial.GroupByFields
	}
	des.ReduceFractionLessThanParams = canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(des.ReduceFractionLessThanParams, initial.ReduceFractionLessThanParams, opts...)
	des.ReduceMakeDistributionParams = canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(des.ReduceMakeDistributionParams, initial.ReduceMakeDistributionParams, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregations(c *Client, des, nw *AlertPolicyConditionsConditionThresholdDenominatorAggregations) *AlertPolicyConditionsConditionThresholdDenominatorAggregations {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(nw.PerSeriesAligner) {
		nw.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(nw.CrossSeriesReducer) {
		nw.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.IsZeroValue(nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}
	nw.ReduceFractionLessThanParams = canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, des.ReduceFractionLessThanParams, nw.ReduceFractionLessThanParams)
	nw.ReduceMakeDistributionParams = canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, des.ReduceMakeDistributionParams, nw.ReduceMakeDistributionParams)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregations) []AlertPolicyConditionsConditionThresholdDenominatorAggregations {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdDenominatorAggregations
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregations) []AlertPolicyConditionsConditionThresholdDenominatorAggregations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdDenominatorAggregations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregations(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Threshold) {
		des.Threshold = initial.Threshold
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c *Client, des, nw *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Threshold) {
		nw.Threshold = des.Threshold
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BucketOptions = canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(des.BucketOptions, initial.BucketOptions, opts...)
	des.ExemplarSampling = canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(des.ExemplarSampling, initial.ExemplarSampling, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c *Client, des, nw *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	if des == nil || nw == nil {
		return nw
	}

	nw.BucketOptions = canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, des.BucketOptions, nw.BucketOptions)
	nw.ExemplarSampling = canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, des.ExemplarSampling, nw.ExemplarSampling)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LinearBuckets = canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des.LinearBuckets, initial.LinearBuckets, opts...)
	des.ExponentialBuckets = canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des.ExponentialBuckets, initial.ExponentialBuckets, opts...)
	des.ExplicitBuckets = canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des.ExplicitBuckets, initial.ExplicitBuckets, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, des, nw *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil || nw == nil {
		return nw
	}

	nw.LinearBuckets = canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, des.LinearBuckets, nw.LinearBuckets)
	nw.ExponentialBuckets = canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, des.ExponentialBuckets, nw.ExponentialBuckets)
	nw.ExplicitBuckets = canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, des.ExplicitBuckets, nw.ExplicitBuckets)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.Width) {
		des.Width = initial.Width
	}
	if dcl.IsZeroValue(des.Offset) {
		des.Offset = initial.Offset
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, des, nw *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.Width) {
		nw.Width = des.Width
	}
	if dcl.IsZeroValue(nw.Offset) {
		nw.Offset = des.Offset
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.GrowthFactor) {
		des.GrowthFactor = initial.GrowthFactor
	}
	if dcl.IsZeroValue(des.Scale) {
		des.Scale = initial.Scale
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, des, nw *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.GrowthFactor) {
		nw.GrowthFactor = des.GrowthFactor
	}
	if dcl.IsZeroValue(nw.Scale) {
		nw.Scale = des.Scale
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Bounds) {
		des.Bounds = initial.Bounds
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, des, nw *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Bounds) {
		nw.Bounds = des.Bounds
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MinimumValue) {
		des.MinimumValue = initial.MinimumValue
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, des, nw *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.MinimumValue) {
		nw.MinimumValue = des.MinimumValue
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionThresholdTrigger(des, initial *AlertPolicyConditionsConditionThresholdTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Count) {
		des.Count = initial.Count
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdTrigger(c *Client, des, nw *AlertPolicyConditionsConditionThresholdTrigger) *AlertPolicyConditionsConditionThresholdTrigger {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Count) {
		nw.Count = des.Count
	}
	if dcl.IsZeroValue(nw.Percent) {
		nw.Percent = des.Percent
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionThresholdTriggerSet(c *Client, des, nw []AlertPolicyConditionsConditionThresholdTrigger) []AlertPolicyConditionsConditionThresholdTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionThresholdTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionThresholdTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionThresholdTriggerSlice(c *Client, des, nw []AlertPolicyConditionsConditionThresholdTrigger) []AlertPolicyConditionsConditionThresholdTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionThresholdTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionThresholdTrigger(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsent(des, initial *AlertPolicyConditionsConditionAbsent, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsent {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}
	if dcl.IsZeroValue(des.Aggregations) {
		des.Aggregations = initial.Aggregations
	}
	des.Duration = canonicalizeAlertPolicyConditionsConditionAbsentDuration(des.Duration, initial.Duration, opts...)
	des.Trigger = canonicalizeAlertPolicyConditionsConditionAbsentTrigger(des.Trigger, initial.Trigger, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsent(c *Client, des, nw *AlertPolicyConditionsConditionAbsent) *AlertPolicyConditionsConditionAbsent {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	nw.Aggregations = canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsSlice(c, des.Aggregations, nw.Aggregations)
	nw.Duration = canonicalizeNewAlertPolicyConditionsConditionAbsentDuration(c, des.Duration, nw.Duration)
	nw.Trigger = canonicalizeNewAlertPolicyConditionsConditionAbsentTrigger(c, des.Trigger, nw.Trigger)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsent) []AlertPolicyConditionsConditionAbsent {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsent
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsent) []AlertPolicyConditionsConditionAbsent {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsent
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsent(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentAggregations(des, initial *AlertPolicyConditionsConditionAbsentAggregations, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		des.AlignmentPeriod = initial.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	}
	if dcl.IsZeroValue(des.GroupByFields) {
		des.GroupByFields = initial.GroupByFields
	}
	des.ReduceFractionLessThanParams = canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(des.ReduceFractionLessThanParams, initial.ReduceFractionLessThanParams, opts...)
	des.ReduceMakeDistributionParams = canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(des.ReduceMakeDistributionParams, initial.ReduceMakeDistributionParams, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregations(c *Client, des, nw *AlertPolicyConditionsConditionAbsentAggregations) *AlertPolicyConditionsConditionAbsentAggregations {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(nw.PerSeriesAligner) {
		nw.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(nw.CrossSeriesReducer) {
		nw.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.IsZeroValue(nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}
	nw.ReduceFractionLessThanParams = canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, des.ReduceFractionLessThanParams, nw.ReduceFractionLessThanParams)
	nw.ReduceMakeDistributionParams = canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, des.ReduceMakeDistributionParams, nw.ReduceMakeDistributionParams)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregations) []AlertPolicyConditionsConditionAbsentAggregations {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentAggregations
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentAggregationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregations) []AlertPolicyConditionsConditionAbsentAggregations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentAggregations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentAggregations(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Threshold) {
		des.Threshold = initial.Threshold
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c *Client, des, nw *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Threshold) {
		nw.Threshold = des.Threshold
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BucketOptions = canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(des.BucketOptions, initial.BucketOptions, opts...)
	des.ExemplarSampling = canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(des.ExemplarSampling, initial.ExemplarSampling, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c *Client, des, nw *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	if des == nil || nw == nil {
		return nw
	}

	nw.BucketOptions = canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, des.BucketOptions, nw.BucketOptions)
	nw.ExemplarSampling = canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, des.ExemplarSampling, nw.ExemplarSampling)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LinearBuckets = canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des.LinearBuckets, initial.LinearBuckets, opts...)
	des.ExponentialBuckets = canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des.ExponentialBuckets, initial.ExponentialBuckets, opts...)
	des.ExplicitBuckets = canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des.ExplicitBuckets, initial.ExplicitBuckets, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, des, nw *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil || nw == nil {
		return nw
	}

	nw.LinearBuckets = canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, des.LinearBuckets, nw.LinearBuckets)
	nw.ExponentialBuckets = canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, des.ExponentialBuckets, nw.ExponentialBuckets)
	nw.ExplicitBuckets = canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, des.ExplicitBuckets, nw.ExplicitBuckets)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.Width) {
		des.Width = initial.Width
	}
	if dcl.IsZeroValue(des.Offset) {
		des.Offset = initial.Offset
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, des, nw *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.Width) {
		nw.Width = des.Width
	}
	if dcl.IsZeroValue(nw.Offset) {
		nw.Offset = des.Offset
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.GrowthFactor) {
		des.GrowthFactor = initial.GrowthFactor
	}
	if dcl.IsZeroValue(des.Scale) {
		des.Scale = initial.Scale
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, des, nw *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.GrowthFactor) {
		nw.GrowthFactor = des.GrowthFactor
	}
	if dcl.IsZeroValue(nw.Scale) {
		nw.Scale = des.Scale
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Bounds) {
		des.Bounds = initial.Bounds
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, des, nw *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Bounds) {
		nw.Bounds = des.Bounds
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MinimumValue) {
		des.MinimumValue = initial.MinimumValue
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, des, nw *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.MinimumValue) {
		nw.MinimumValue = des.MinimumValue
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentDuration(des, initial *AlertPolicyConditionsConditionAbsentDuration, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentDuration(c *Client, des, nw *AlertPolicyConditionsConditionAbsentDuration) *AlertPolicyConditionsConditionAbsentDuration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentDurationSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentDuration) []AlertPolicyConditionsConditionAbsentDuration {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentDuration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentDurationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentDurationSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentDuration) []AlertPolicyConditionsConditionAbsentDuration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentDuration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentDuration(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionAbsentTrigger(des, initial *AlertPolicyConditionsConditionAbsentTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Count) {
		des.Count = initial.Count
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentTrigger(c *Client, des, nw *AlertPolicyConditionsConditionAbsentTrigger) *AlertPolicyConditionsConditionAbsentTrigger {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Count) {
		nw.Count = des.Count
	}
	if dcl.IsZeroValue(nw.Percent) {
		nw.Percent = des.Percent
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionAbsentTriggerSet(c *Client, des, nw []AlertPolicyConditionsConditionAbsentTrigger) []AlertPolicyConditionsConditionAbsentTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionAbsentTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionAbsentTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentTriggerSlice(c *Client, des, nw []AlertPolicyConditionsConditionAbsentTrigger) []AlertPolicyConditionsConditionAbsentTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionAbsentTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionAbsentTrigger(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionMatchedLog(des, initial *AlertPolicyConditionsConditionMatchedLog, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionMatchedLog {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}
	if dcl.IsZeroValue(des.LabelExtractors) {
		des.LabelExtractors = initial.LabelExtractors
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionMatchedLog(c *Client, des, nw *AlertPolicyConditionsConditionMatchedLog) *AlertPolicyConditionsConditionMatchedLog {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	if dcl.IsZeroValue(nw.LabelExtractors) {
		nw.LabelExtractors = des.LabelExtractors
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionMatchedLogSet(c *Client, des, nw []AlertPolicyConditionsConditionMatchedLog) []AlertPolicyConditionsConditionMatchedLog {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionMatchedLog
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionMatchedLogNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionMatchedLogSlice(c *Client, des, nw []AlertPolicyConditionsConditionMatchedLog) []AlertPolicyConditionsConditionMatchedLog {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionMatchedLog
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionMatchedLog(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionClusterOutlier(des, initial *AlertPolicyConditionsConditionClusterOutlier, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionClusterOutlier {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionClusterOutlier(c *Client, des, nw *AlertPolicyConditionsConditionClusterOutlier) *AlertPolicyConditionsConditionClusterOutlier {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionClusterOutlierSet(c *Client, des, nw []AlertPolicyConditionsConditionClusterOutlier) []AlertPolicyConditionsConditionClusterOutlier {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionClusterOutlier
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionClusterOutlierNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionClusterOutlierSlice(c *Client, des, nw []AlertPolicyConditionsConditionClusterOutlier) []AlertPolicyConditionsConditionClusterOutlier {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionClusterOutlier
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionClusterOutlier(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRate(des, initial *AlertPolicyConditionsConditionRate, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}
	if dcl.IsZeroValue(des.Aggregations) {
		des.Aggregations = initial.Aggregations
	}
	if dcl.IsZeroValue(des.Comparison) {
		des.Comparison = initial.Comparison
	}
	if dcl.IsZeroValue(des.ThresholdValue) {
		des.ThresholdValue = initial.ThresholdValue
	}
	des.TimeWindow = canonicalizeAlertPolicyConditionsConditionRateTimeWindow(des.TimeWindow, initial.TimeWindow, opts...)
	des.Trigger = canonicalizeAlertPolicyConditionsConditionRateTrigger(des.Trigger, initial.Trigger, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRate(c *Client, des, nw *AlertPolicyConditionsConditionRate) *AlertPolicyConditionsConditionRate {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	nw.Aggregations = canonicalizeNewAlertPolicyConditionsConditionRateAggregationsSlice(c, des.Aggregations, nw.Aggregations)
	if dcl.IsZeroValue(nw.Comparison) {
		nw.Comparison = des.Comparison
	}
	if dcl.IsZeroValue(nw.ThresholdValue) {
		nw.ThresholdValue = des.ThresholdValue
	}
	nw.TimeWindow = canonicalizeNewAlertPolicyConditionsConditionRateTimeWindow(c, des.TimeWindow, nw.TimeWindow)
	nw.Trigger = canonicalizeNewAlertPolicyConditionsConditionRateTrigger(c, des.Trigger, nw.Trigger)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateSet(c *Client, des, nw []AlertPolicyConditionsConditionRate) []AlertPolicyConditionsConditionRate {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRate
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateSlice(c *Client, des, nw []AlertPolicyConditionsConditionRate) []AlertPolicyConditionsConditionRate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRate(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateAggregations(des, initial *AlertPolicyConditionsConditionRateAggregations, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		des.AlignmentPeriod = initial.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	}
	if dcl.IsZeroValue(des.GroupByFields) {
		des.GroupByFields = initial.GroupByFields
	}
	des.ReduceFractionLessThanParams = canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(des.ReduceFractionLessThanParams, initial.ReduceFractionLessThanParams, opts...)
	des.ReduceMakeDistributionParams = canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(des.ReduceMakeDistributionParams, initial.ReduceMakeDistributionParams, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregations(c *Client, des, nw *AlertPolicyConditionsConditionRateAggregations) *AlertPolicyConditionsConditionRateAggregations {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(nw.PerSeriesAligner) {
		nw.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(nw.CrossSeriesReducer) {
		nw.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.IsZeroValue(nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}
	nw.ReduceFractionLessThanParams = canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, des.ReduceFractionLessThanParams, nw.ReduceFractionLessThanParams)
	nw.ReduceMakeDistributionParams = canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, des.ReduceMakeDistributionParams, nw.ReduceMakeDistributionParams)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsSet(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregations) []AlertPolicyConditionsConditionRateAggregations {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateAggregations
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateAggregationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregations) []AlertPolicyConditionsConditionRateAggregations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateAggregations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateAggregations(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Threshold) {
		des.Threshold = initial.Threshold
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c *Client, des, nw *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Threshold) {
		nw.Threshold = des.Threshold
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsSet(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BucketOptions = canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(des.BucketOptions, initial.BucketOptions, opts...)
	des.ExemplarSampling = canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(des.ExemplarSampling, initial.ExemplarSampling, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c *Client, des, nw *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	if des == nil || nw == nil {
		return nw
	}

	nw.BucketOptions = canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, des.BucketOptions, nw.BucketOptions)
	nw.ExemplarSampling = canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, des.ExemplarSampling, nw.ExemplarSampling)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsSet(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LinearBuckets = canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des.LinearBuckets, initial.LinearBuckets, opts...)
	des.ExponentialBuckets = canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des.ExponentialBuckets, initial.ExponentialBuckets, opts...)
	des.ExplicitBuckets = canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des.ExplicitBuckets, initial.ExplicitBuckets, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, des, nw *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil || nw == nil {
		return nw
	}

	nw.LinearBuckets = canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, des.LinearBuckets, nw.LinearBuckets)
	nw.ExponentialBuckets = canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, des.ExponentialBuckets, nw.ExponentialBuckets)
	nw.ExplicitBuckets = canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, des.ExplicitBuckets, nw.ExplicitBuckets)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsSet(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.Width) {
		des.Width = initial.Width
	}
	if dcl.IsZeroValue(des.Offset) {
		des.Offset = initial.Offset
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, des, nw *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.Width) {
		nw.Width = des.Width
	}
	if dcl.IsZeroValue(nw.Offset) {
		nw.Offset = des.Offset
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.GrowthFactor) {
		des.GrowthFactor = initial.GrowthFactor
	}
	if dcl.IsZeroValue(des.Scale) {
		des.Scale = initial.Scale
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, des, nw *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.GrowthFactor) {
		nw.GrowthFactor = des.GrowthFactor
	}
	if dcl.IsZeroValue(nw.Scale) {
		nw.Scale = des.Scale
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Bounds) {
		des.Bounds = initial.Bounds
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, des, nw *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Bounds) {
		nw.Bounds = des.Bounds
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSet(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MinimumValue) {
		des.MinimumValue = initial.MinimumValue
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, des, nw *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.MinimumValue) {
		nw.MinimumValue = des.MinimumValue
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingSet(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateTimeWindow(des, initial *AlertPolicyConditionsConditionRateTimeWindow, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateTimeWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewAlertPolicyConditionsConditionRateTimeWindow(c *Client, des, nw *AlertPolicyConditionsConditionRateTimeWindow) *AlertPolicyConditionsConditionRateTimeWindow {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateTimeWindowSet(c *Client, des, nw []AlertPolicyConditionsConditionRateTimeWindow) []AlertPolicyConditionsConditionRateTimeWindow {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateTimeWindow
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateTimeWindowNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateTimeWindowSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateTimeWindow) []AlertPolicyConditionsConditionRateTimeWindow {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateTimeWindow
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateTimeWindow(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionRateTrigger(des, initial *AlertPolicyConditionsConditionRateTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Count) {
		des.Count = initial.Count
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionRateTrigger(c *Client, des, nw *AlertPolicyConditionsConditionRateTrigger) *AlertPolicyConditionsConditionRateTrigger {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Count) {
		nw.Count = des.Count
	}
	if dcl.IsZeroValue(nw.Percent) {
		nw.Percent = des.Percent
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionRateTriggerSet(c *Client, des, nw []AlertPolicyConditionsConditionRateTrigger) []AlertPolicyConditionsConditionRateTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionRateTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionRateTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionRateTriggerSlice(c *Client, des, nw []AlertPolicyConditionsConditionRateTrigger) []AlertPolicyConditionsConditionRateTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionRateTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionRateTrigger(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionUpMon(des, initial *AlertPolicyConditionsConditionUpMon, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionUpMon {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}
	if dcl.StringCanonicalize(des.EndpointId, initial.EndpointId) || dcl.IsZeroValue(des.EndpointId) {
		des.EndpointId = initial.EndpointId
	}
	if dcl.StringCanonicalize(des.CheckId, initial.CheckId) || dcl.IsZeroValue(des.CheckId) {
		des.CheckId = initial.CheckId
	}
	des.Duration = canonicalizeAlertPolicyConditionsConditionUpMonDuration(des.Duration, initial.Duration, opts...)
	des.Trigger = canonicalizeAlertPolicyConditionsConditionUpMonTrigger(des.Trigger, initial.Trigger, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionUpMon(c *Client, des, nw *AlertPolicyConditionsConditionUpMon) *AlertPolicyConditionsConditionUpMon {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	if dcl.StringCanonicalize(des.EndpointId, nw.EndpointId) {
		nw.EndpointId = des.EndpointId
	}
	if dcl.StringCanonicalize(des.CheckId, nw.CheckId) {
		nw.CheckId = des.CheckId
	}
	nw.Duration = canonicalizeNewAlertPolicyConditionsConditionUpMonDuration(c, des.Duration, nw.Duration)
	nw.Trigger = canonicalizeNewAlertPolicyConditionsConditionUpMonTrigger(c, des.Trigger, nw.Trigger)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionUpMonSet(c *Client, des, nw []AlertPolicyConditionsConditionUpMon) []AlertPolicyConditionsConditionUpMon {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionUpMon
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionUpMonNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionUpMonSlice(c *Client, des, nw []AlertPolicyConditionsConditionUpMon) []AlertPolicyConditionsConditionUpMon {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionUpMon
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionUpMon(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionUpMonDuration(des, initial *AlertPolicyConditionsConditionUpMonDuration, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionUpMonDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewAlertPolicyConditionsConditionUpMonDuration(c *Client, des, nw *AlertPolicyConditionsConditionUpMonDuration) *AlertPolicyConditionsConditionUpMonDuration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionUpMonDurationSet(c *Client, des, nw []AlertPolicyConditionsConditionUpMonDuration) []AlertPolicyConditionsConditionUpMonDuration {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionUpMonDuration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionUpMonDurationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionUpMonDurationSlice(c *Client, des, nw []AlertPolicyConditionsConditionUpMonDuration) []AlertPolicyConditionsConditionUpMonDuration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionUpMonDuration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionUpMonDuration(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionUpMonTrigger(des, initial *AlertPolicyConditionsConditionUpMonTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionUpMonTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Count) {
		des.Count = initial.Count
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionUpMonTrigger(c *Client, des, nw *AlertPolicyConditionsConditionUpMonTrigger) *AlertPolicyConditionsConditionUpMonTrigger {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Count) {
		nw.Count = des.Count
	}
	if dcl.IsZeroValue(nw.Percent) {
		nw.Percent = des.Percent
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionUpMonTriggerSet(c *Client, des, nw []AlertPolicyConditionsConditionUpMonTrigger) []AlertPolicyConditionsConditionUpMonTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionUpMonTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionUpMonTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionUpMonTriggerSlice(c *Client, des, nw []AlertPolicyConditionsConditionUpMonTrigger) []AlertPolicyConditionsConditionUpMonTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionUpMonTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionUpMonTrigger(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionProcessCount(des, initial *AlertPolicyConditionsConditionProcessCount, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionProcessCount {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Process, initial.Process) || dcl.IsZeroValue(des.Process) {
		des.Process = initial.Process
	}
	if dcl.StringCanonicalize(des.User, initial.User) || dcl.IsZeroValue(des.User) {
		des.User = initial.User
	}
	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}
	if dcl.IsZeroValue(des.Comparison) {
		des.Comparison = initial.Comparison
	}
	if dcl.IsZeroValue(des.ProcessCountThreshold) {
		des.ProcessCountThreshold = initial.ProcessCountThreshold
	}
	des.Trigger = canonicalizeAlertPolicyConditionsConditionProcessCountTrigger(des.Trigger, initial.Trigger, opts...)
	des.Duration = canonicalizeAlertPolicyConditionsConditionProcessCountDuration(des.Duration, initial.Duration, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionProcessCount(c *Client, des, nw *AlertPolicyConditionsConditionProcessCount) *AlertPolicyConditionsConditionProcessCount {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Process, nw.Process) {
		nw.Process = des.Process
	}
	if dcl.StringCanonicalize(des.User, nw.User) {
		nw.User = des.User
	}
	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	if dcl.IsZeroValue(nw.Comparison) {
		nw.Comparison = des.Comparison
	}
	if dcl.IsZeroValue(nw.ProcessCountThreshold) {
		nw.ProcessCountThreshold = des.ProcessCountThreshold
	}
	nw.Trigger = canonicalizeNewAlertPolicyConditionsConditionProcessCountTrigger(c, des.Trigger, nw.Trigger)
	nw.Duration = canonicalizeNewAlertPolicyConditionsConditionProcessCountDuration(c, des.Duration, nw.Duration)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionProcessCountSet(c *Client, des, nw []AlertPolicyConditionsConditionProcessCount) []AlertPolicyConditionsConditionProcessCount {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionProcessCount
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionProcessCountNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionProcessCountSlice(c *Client, des, nw []AlertPolicyConditionsConditionProcessCount) []AlertPolicyConditionsConditionProcessCount {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionProcessCount
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionProcessCount(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionProcessCountTrigger(des, initial *AlertPolicyConditionsConditionProcessCountTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionProcessCountTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Count) {
		des.Count = initial.Count
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionProcessCountTrigger(c *Client, des, nw *AlertPolicyConditionsConditionProcessCountTrigger) *AlertPolicyConditionsConditionProcessCountTrigger {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Count) {
		nw.Count = des.Count
	}
	if dcl.IsZeroValue(nw.Percent) {
		nw.Percent = des.Percent
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionProcessCountTriggerSet(c *Client, des, nw []AlertPolicyConditionsConditionProcessCountTrigger) []AlertPolicyConditionsConditionProcessCountTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionProcessCountTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionProcessCountTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionProcessCountTriggerSlice(c *Client, des, nw []AlertPolicyConditionsConditionProcessCountTrigger) []AlertPolicyConditionsConditionProcessCountTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionProcessCountTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionProcessCountTrigger(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionProcessCountDuration(des, initial *AlertPolicyConditionsConditionProcessCountDuration, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionProcessCountDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewAlertPolicyConditionsConditionProcessCountDuration(c *Client, des, nw *AlertPolicyConditionsConditionProcessCountDuration) *AlertPolicyConditionsConditionProcessCountDuration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionProcessCountDurationSet(c *Client, des, nw []AlertPolicyConditionsConditionProcessCountDuration) []AlertPolicyConditionsConditionProcessCountDuration {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionProcessCountDuration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionProcessCountDurationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionProcessCountDurationSlice(c *Client, des, nw []AlertPolicyConditionsConditionProcessCountDuration) []AlertPolicyConditionsConditionProcessCountDuration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionProcessCountDuration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionProcessCountDuration(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionTimeSeriesQueryLanguage(des, initial *AlertPolicyConditionsConditionTimeSeriesQueryLanguage, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Query, initial.Query) || dcl.IsZeroValue(des.Query) {
		des.Query = initial.Query
	}
	if dcl.StringCanonicalize(des.Summary, initial.Summary) || dcl.IsZeroValue(des.Summary) {
		des.Summary = initial.Summary
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c *Client, des, nw *AlertPolicyConditionsConditionTimeSeriesQueryLanguage) *AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Query, nw.Query) {
		nw.Query = des.Query
	}
	if dcl.StringCanonicalize(des.Summary, nw.Summary) {
		nw.Summary = des.Summary
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionTimeSeriesQueryLanguageSet(c *Client, des, nw []AlertPolicyConditionsConditionTimeSeriesQueryLanguage) []AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionTimeSeriesQueryLanguage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionTimeSeriesQueryLanguageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionTimeSeriesQueryLanguageSlice(c *Client, des, nw []AlertPolicyConditionsConditionTimeSeriesQueryLanguage) []AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionTimeSeriesQueryLanguage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionMonitoringQueryLanguage(des, initial *AlertPolicyConditionsConditionMonitoringQueryLanguage, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionMonitoringQueryLanguage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Query, initial.Query) || dcl.IsZeroValue(des.Query) {
		des.Query = initial.Query
	}
	des.Duration = canonicalizeAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(des.Duration, initial.Duration, opts...)
	des.Trigger = canonicalizeAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(des.Trigger, initial.Trigger, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguage(c *Client, des, nw *AlertPolicyConditionsConditionMonitoringQueryLanguage) *AlertPolicyConditionsConditionMonitoringQueryLanguage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Query, nw.Query) {
		nw.Query = des.Query
	}
	nw.Duration = canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, des.Duration, nw.Duration)
	nw.Trigger = canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, des.Trigger, nw.Trigger)

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageSet(c *Client, des, nw []AlertPolicyConditionsConditionMonitoringQueryLanguage) []AlertPolicyConditionsConditionMonitoringQueryLanguage {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionMonitoringQueryLanguage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionMonitoringQueryLanguageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageSlice(c *Client, des, nw []AlertPolicyConditionsConditionMonitoringQueryLanguage) []AlertPolicyConditionsConditionMonitoringQueryLanguage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionMonitoringQueryLanguage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguage(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(des, initial *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c *Client, des, nw *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageDurationSet(c *Client, des, nw []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionMonitoringQueryLanguageDurationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageDurationSlice(c *Client, des, nw []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(des, initial *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Count) {
		des.Count = initial.Count
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c *Client, des, nw *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Count) {
		nw.Count = des.Count
	}
	if dcl.IsZeroValue(nw.Percent) {
		nw.Percent = des.Percent
	}

	return nw
}

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerSet(c *Client, des, nw []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerSlice(c *Client, des, nw []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyEnabled(des, initial *AlertPolicyEnabled, opts ...dcl.ApplyOption) *AlertPolicyEnabled {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewAlertPolicyEnabled(c *Client, des, nw *AlertPolicyEnabled) *AlertPolicyEnabled {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewAlertPolicyEnabledSet(c *Client, des, nw []AlertPolicyEnabled) []AlertPolicyEnabled {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyEnabled
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyEnabledNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyEnabledSlice(c *Client, des, nw []AlertPolicyEnabled) []AlertPolicyEnabled {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyEnabled
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyEnabled(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyValidity(des, initial *AlertPolicyValidity, opts ...dcl.ApplyOption) *AlertPolicyValidity {
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

func canonicalizeNewAlertPolicyValidity(c *Client, des, nw *AlertPolicyValidity) *AlertPolicyValidity {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Code) {
		nw.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}
	nw.Details = canonicalizeNewAlertPolicyValidityDetailsSlice(c, des.Details, nw.Details)

	return nw
}

func canonicalizeNewAlertPolicyValiditySet(c *Client, des, nw []AlertPolicyValidity) []AlertPolicyValidity {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyValidity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyValidityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyValiditySlice(c *Client, des, nw []AlertPolicyValidity) []AlertPolicyValidity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyValidity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyValidity(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyValidityDetails(des, initial *AlertPolicyValidityDetails, opts ...dcl.ApplyOption) *AlertPolicyValidityDetails {
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

func canonicalizeNewAlertPolicyValidityDetails(c *Client, des, nw *AlertPolicyValidityDetails) *AlertPolicyValidityDetails {
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

func canonicalizeNewAlertPolicyValidityDetailsSet(c *Client, des, nw []AlertPolicyValidityDetails) []AlertPolicyValidityDetails {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyValidityDetails
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyValidityDetailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyValidityDetailsSlice(c *Client, des, nw []AlertPolicyValidityDetails) []AlertPolicyValidityDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyValidityDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyValidityDetails(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyCreationRecord(des, initial *AlertPolicyCreationRecord, opts ...dcl.ApplyOption) *AlertPolicyCreationRecord {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.MutateTime = canonicalizeAlertPolicyCreationRecordMutateTime(des.MutateTime, initial.MutateTime, opts...)
	if dcl.StringCanonicalize(des.MutatedBy, initial.MutatedBy) || dcl.IsZeroValue(des.MutatedBy) {
		des.MutatedBy = initial.MutatedBy
	}

	return des
}

func canonicalizeNewAlertPolicyCreationRecord(c *Client, des, nw *AlertPolicyCreationRecord) *AlertPolicyCreationRecord {
	if des == nil || nw == nil {
		return nw
	}

	nw.MutateTime = canonicalizeNewAlertPolicyCreationRecordMutateTime(c, des.MutateTime, nw.MutateTime)
	if dcl.StringCanonicalize(des.MutatedBy, nw.MutatedBy) {
		nw.MutatedBy = des.MutatedBy
	}

	return nw
}

func canonicalizeNewAlertPolicyCreationRecordSet(c *Client, des, nw []AlertPolicyCreationRecord) []AlertPolicyCreationRecord {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyCreationRecord
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyCreationRecordNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyCreationRecordSlice(c *Client, des, nw []AlertPolicyCreationRecord) []AlertPolicyCreationRecord {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyCreationRecord
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyCreationRecord(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyCreationRecordMutateTime(des, initial *AlertPolicyCreationRecordMutateTime, opts ...dcl.ApplyOption) *AlertPolicyCreationRecordMutateTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewAlertPolicyCreationRecordMutateTime(c *Client, des, nw *AlertPolicyCreationRecordMutateTime) *AlertPolicyCreationRecordMutateTime {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewAlertPolicyCreationRecordMutateTimeSet(c *Client, des, nw []AlertPolicyCreationRecordMutateTime) []AlertPolicyCreationRecordMutateTime {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyCreationRecordMutateTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyCreationRecordMutateTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyCreationRecordMutateTimeSlice(c *Client, des, nw []AlertPolicyCreationRecordMutateTime) []AlertPolicyCreationRecordMutateTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyCreationRecordMutateTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyCreationRecordMutateTime(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyMutationRecord(des, initial *AlertPolicyMutationRecord, opts ...dcl.ApplyOption) *AlertPolicyMutationRecord {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.MutateTime = canonicalizeAlertPolicyMutationRecordMutateTime(des.MutateTime, initial.MutateTime, opts...)
	if dcl.StringCanonicalize(des.MutatedBy, initial.MutatedBy) || dcl.IsZeroValue(des.MutatedBy) {
		des.MutatedBy = initial.MutatedBy
	}

	return des
}

func canonicalizeNewAlertPolicyMutationRecord(c *Client, des, nw *AlertPolicyMutationRecord) *AlertPolicyMutationRecord {
	if des == nil || nw == nil {
		return nw
	}

	nw.MutateTime = canonicalizeNewAlertPolicyMutationRecordMutateTime(c, des.MutateTime, nw.MutateTime)
	if dcl.StringCanonicalize(des.MutatedBy, nw.MutatedBy) {
		nw.MutatedBy = des.MutatedBy
	}

	return nw
}

func canonicalizeNewAlertPolicyMutationRecordSet(c *Client, des, nw []AlertPolicyMutationRecord) []AlertPolicyMutationRecord {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyMutationRecord
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyMutationRecordNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyMutationRecordSlice(c *Client, des, nw []AlertPolicyMutationRecord) []AlertPolicyMutationRecord {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyMutationRecord
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyMutationRecord(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyMutationRecordMutateTime(des, initial *AlertPolicyMutationRecordMutateTime, opts ...dcl.ApplyOption) *AlertPolicyMutationRecordMutateTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewAlertPolicyMutationRecordMutateTime(c *Client, des, nw *AlertPolicyMutationRecordMutateTime) *AlertPolicyMutationRecordMutateTime {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewAlertPolicyMutationRecordMutateTimeSet(c *Client, des, nw []AlertPolicyMutationRecordMutateTime) []AlertPolicyMutationRecordMutateTime {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyMutationRecordMutateTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyMutationRecordMutateTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyMutationRecordMutateTimeSlice(c *Client, des, nw []AlertPolicyMutationRecordMutateTime) []AlertPolicyMutationRecordMutateTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyMutationRecordMutateTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyMutationRecordMutateTime(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyIncidentStrategy(des, initial *AlertPolicyIncidentStrategy, opts ...dcl.ApplyOption) *AlertPolicyIncidentStrategy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewAlertPolicyIncidentStrategy(c *Client, des, nw *AlertPolicyIncidentStrategy) *AlertPolicyIncidentStrategy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
	}

	return nw
}

func canonicalizeNewAlertPolicyIncidentStrategySet(c *Client, des, nw []AlertPolicyIncidentStrategy) []AlertPolicyIncidentStrategy {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyIncidentStrategy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyIncidentStrategyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyIncidentStrategySlice(c *Client, des, nw []AlertPolicyIncidentStrategy) []AlertPolicyIncidentStrategy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyIncidentStrategy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyIncidentStrategy(c, &d, &n))
	}

	return items
}

func canonicalizeAlertPolicyMetadata(des, initial *AlertPolicyMetadata, opts ...dcl.ApplyOption) *AlertPolicyMetadata {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.SloNames) {
		des.SloNames = initial.SloNames
	}

	return des
}

func canonicalizeNewAlertPolicyMetadata(c *Client, des, nw *AlertPolicyMetadata) *AlertPolicyMetadata {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.SloNames) {
		nw.SloNames = des.SloNames
	}

	return nw
}

func canonicalizeNewAlertPolicyMetadataSet(c *Client, des, nw []AlertPolicyMetadata) []AlertPolicyMetadata {
	if des == nil {
		return nw
	}
	var reorderedNew []AlertPolicyMetadata
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareAlertPolicyMetadataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewAlertPolicyMetadataSlice(c *Client, des, nw []AlertPolicyMetadata) []AlertPolicyMetadata {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AlertPolicyMetadata
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAlertPolicyMetadata(c, &d, &n))
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
func diffAlertPolicy(c *Client, desired, actual *AlertPolicy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Documentation, actual.Documentation, dcl.Info{ObjectFunction: compareAlertPolicyDocumentationNewStyle, EmptyObject: EmptyAlertPolicyDocumentation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Documentation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserLabels, actual.UserLabels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UserLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Conditions, actual.Conditions, dcl.Info{ObjectFunction: compareAlertPolicyConditionsNewStyle, EmptyObject: EmptyAlertPolicyConditions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Conditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Combiner, actual.Combiner, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Combiner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{ObjectFunction: compareAlertPolicyEnabledNewStyle, EmptyObject: EmptyAlertPolicyEnabled, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Validity, actual.Validity, dcl.Info{ObjectFunction: compareAlertPolicyValidityNewStyle, EmptyObject: EmptyAlertPolicyValidity, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Validity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NotificationChannels, actual.NotificationChannels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NotificationChannels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreationRecord, actual.CreationRecord, dcl.Info{ObjectFunction: compareAlertPolicyCreationRecordNewStyle, EmptyObject: EmptyAlertPolicyCreationRecord, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationRecord")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MutationRecord, actual.MutationRecord, dcl.Info{ObjectFunction: compareAlertPolicyMutationRecordNewStyle, EmptyObject: EmptyAlertPolicyMutationRecord, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MutationRecord")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncidentStrategy, actual.IncidentStrategy, dcl.Info{ObjectFunction: compareAlertPolicyIncidentStrategyNewStyle, EmptyObject: EmptyAlertPolicyIncidentStrategy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IncidentStrategy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{ObjectFunction: compareAlertPolicyMetadataNewStyle, EmptyObject: EmptyAlertPolicyMetadata, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
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
func compareAlertPolicyDocumentationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyDocumentation)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyDocumentation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyDocumentation or *AlertPolicyDocumentation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyDocumentation)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyDocumentation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyDocumentation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MimeType, actual.MimeType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MimeType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditions)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditions or *AlertPolicyConditions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditions)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceStateFilter, actual.ResourceStateFilter, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceStateFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConditionThreshold, actual.ConditionThreshold, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThreshold, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConditionThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConditionAbsent, actual.ConditionAbsent, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsent, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConditionAbsent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConditionMatchedLog, actual.ConditionMatchedLog, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionMatchedLogNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionMatchedLog, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConditionMatchedLog")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConditionClusterOutlier, actual.ConditionClusterOutlier, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionClusterOutlierNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionClusterOutlier, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConditionClusterOutlier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConditionRate, actual.ConditionRate, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRate, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConditionRate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConditionUpMon, actual.ConditionUpMon, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionUpMonNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionUpMon, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConditionUpMon")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConditionProcessCount, actual.ConditionProcessCount, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionProcessCountNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionProcessCount, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConditionProcessCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConditionTimeSeriesQueryLanguage, actual.ConditionTimeSeriesQueryLanguage, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionTimeSeriesQueryLanguageNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionTimeSeriesQueryLanguage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConditionTimeSeriesQueryLanguage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConditionMonitoringQueryLanguage, actual.ConditionMonitoringQueryLanguage, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionMonitoringQueryLanguageNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionMonitoringQueryLanguage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConditionMonitoringQueryLanguage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThreshold)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThreshold)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThreshold or *AlertPolicyConditionsConditionThreshold", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThreshold)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThreshold)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThreshold", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aggregations, actual.Aggregations, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdAggregationsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdAggregations, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Aggregations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DenominatorFilter, actual.DenominatorFilter, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DenominatorFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DenominatorAggregations, actual.DenominatorAggregations, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregations, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DenominatorAggregations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Comparison, actual.Comparison, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Comparison")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ThresholdValue, actual.ThresholdValue, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ThresholdValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Duration, actual.Duration, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Duration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Trigger, actual.Trigger, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdTriggerNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdTrigger, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Trigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdAggregationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdAggregations)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdAggregations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregations or *AlertPolicyConditionsConditionThresholdAggregations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdAggregations)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdAggregations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReduceFractionLessThanParams, actual.ReduceFractionLessThanParams, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReduceFractionLessThanParams")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReduceMakeDistributionParams, actual.ReduceMakeDistributionParams, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReduceMakeDistributionParams")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams or *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Threshold, actual.Threshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Threshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams or *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BucketOptions, actual.BucketOptions, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BucketOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExemplarSampling, actual.ExemplarSampling, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExemplarSampling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions or *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LinearBuckets, actual.LinearBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LinearBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExponentialBuckets, actual.ExponentialBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExponentialBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExplicitBuckets, actual.ExplicitBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExplicitBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets or *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Width, actual.Width, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Width")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Offset, actual.Offset, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Offset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets or *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GrowthFactor, actual.GrowthFactor, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GrowthFactor")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scale, actual.Scale, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Scale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets or *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Bounds, actual.Bounds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Bounds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling or *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MinimumValue, actual.MinimumValue, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinimumValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdDenominatorAggregations)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdDenominatorAggregations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregations or *AlertPolicyConditionsConditionThresholdDenominatorAggregations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdDenominatorAggregations)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdDenominatorAggregations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReduceFractionLessThanParams, actual.ReduceFractionLessThanParams, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReduceFractionLessThanParams")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReduceMakeDistributionParams, actual.ReduceMakeDistributionParams, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReduceMakeDistributionParams")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams or *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Threshold, actual.Threshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Threshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams or *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BucketOptions, actual.BucketOptions, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BucketOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExemplarSampling, actual.ExemplarSampling, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExemplarSampling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions or *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LinearBuckets, actual.LinearBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LinearBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExponentialBuckets, actual.ExponentialBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExponentialBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExplicitBuckets, actual.ExplicitBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExplicitBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets or *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Width, actual.Width, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Width")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Offset, actual.Offset, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Offset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets or *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GrowthFactor, actual.GrowthFactor, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GrowthFactor")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scale, actual.Scale, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Scale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets or *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Bounds, actual.Bounds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Bounds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling or *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MinimumValue, actual.MinimumValue, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinimumValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionThresholdTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionThresholdTrigger)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionThresholdTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdTrigger or *AlertPolicyConditionsConditionThresholdTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionThresholdTrigger)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionThresholdTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionThresholdTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Count, actual.Count, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Count")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsent)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsent)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsent or *AlertPolicyConditionsConditionAbsent", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsent)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsent)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsent", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aggregations, actual.Aggregations, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentAggregationsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentAggregations, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Aggregations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Duration, actual.Duration, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentDurationNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentDuration, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Duration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Trigger, actual.Trigger, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentTriggerNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentTrigger, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Trigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentAggregationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentAggregations)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentAggregations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregations or *AlertPolicyConditionsConditionAbsentAggregations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentAggregations)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentAggregations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReduceFractionLessThanParams, actual.ReduceFractionLessThanParams, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReduceFractionLessThanParams")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReduceMakeDistributionParams, actual.ReduceMakeDistributionParams, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReduceMakeDistributionParams")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams or *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Threshold, actual.Threshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Threshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams or *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BucketOptions, actual.BucketOptions, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BucketOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExemplarSampling, actual.ExemplarSampling, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExemplarSampling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions or *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LinearBuckets, actual.LinearBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LinearBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExponentialBuckets, actual.ExponentialBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExponentialBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExplicitBuckets, actual.ExplicitBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExplicitBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets or *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Width, actual.Width, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Width")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Offset, actual.Offset, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Offset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets or *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GrowthFactor, actual.GrowthFactor, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GrowthFactor")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scale, actual.Scale, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Scale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets or *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Bounds, actual.Bounds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Bounds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling or *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MinimumValue, actual.MinimumValue, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinimumValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentDurationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentDuration)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentDuration or *AlertPolicyConditionsConditionAbsentDuration", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentDuration)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentDuration", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionAbsentTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionAbsentTrigger)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionAbsentTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentTrigger or *AlertPolicyConditionsConditionAbsentTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionAbsentTrigger)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionAbsentTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionAbsentTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Count, actual.Count, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Count")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionMatchedLogNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionMatchedLog)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionMatchedLog)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionMatchedLog or *AlertPolicyConditionsConditionMatchedLog", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionMatchedLog)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionMatchedLog)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionMatchedLog", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LabelExtractors, actual.LabelExtractors, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LabelExtractors")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionClusterOutlierNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionClusterOutlier)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionClusterOutlier)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionClusterOutlier or *AlertPolicyConditionsConditionClusterOutlier", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionClusterOutlier)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionClusterOutlier)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionClusterOutlier", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRate)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRate or *AlertPolicyConditionsConditionRate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRate)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aggregations, actual.Aggregations, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateAggregationsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateAggregations, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Aggregations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Comparison, actual.Comparison, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Comparison")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ThresholdValue, actual.ThresholdValue, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ThresholdValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeWindow, actual.TimeWindow, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateTimeWindowNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateTimeWindow, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TimeWindow")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Trigger, actual.Trigger, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateTriggerNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateTrigger, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Trigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateAggregationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateAggregations)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateAggregations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregations or *AlertPolicyConditionsConditionRateAggregations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateAggregations)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateAggregations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReduceFractionLessThanParams, actual.ReduceFractionLessThanParams, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReduceFractionLessThanParams")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReduceMakeDistributionParams, actual.ReduceMakeDistributionParams, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReduceMakeDistributionParams")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams or *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Threshold, actual.Threshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Threshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams or *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BucketOptions, actual.BucketOptions, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BucketOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExemplarSampling, actual.ExemplarSampling, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExemplarSampling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions or *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LinearBuckets, actual.LinearBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LinearBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExponentialBuckets, actual.ExponentialBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExponentialBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExplicitBuckets, actual.ExplicitBuckets, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExplicitBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets or *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Width, actual.Width, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Width")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Offset, actual.Offset, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Offset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets or *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GrowthFactor, actual.GrowthFactor, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GrowthFactor")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scale, actual.Scale, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Scale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets or *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Bounds, actual.Bounds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Bounds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling or *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MinimumValue, actual.MinimumValue, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinimumValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateTimeWindowNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateTimeWindow)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateTimeWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateTimeWindow or *AlertPolicyConditionsConditionRateTimeWindow", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateTimeWindow)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateTimeWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateTimeWindow", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionRateTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionRateTrigger)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionRateTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateTrigger or *AlertPolicyConditionsConditionRateTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionRateTrigger)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionRateTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionRateTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Count, actual.Count, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Count")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionUpMonNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionUpMon)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionUpMon)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionUpMon or *AlertPolicyConditionsConditionUpMon", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionUpMon)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionUpMon)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionUpMon", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndpointId, actual.EndpointId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EndpointId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CheckId, actual.CheckId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CheckId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Duration, actual.Duration, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionUpMonDurationNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionUpMonDuration, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Duration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Trigger, actual.Trigger, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionUpMonTriggerNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionUpMonTrigger, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Trigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionUpMonDurationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionUpMonDuration)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionUpMonDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionUpMonDuration or *AlertPolicyConditionsConditionUpMonDuration", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionUpMonDuration)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionUpMonDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionUpMonDuration", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionUpMonTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionUpMonTrigger)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionUpMonTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionUpMonTrigger or *AlertPolicyConditionsConditionUpMonTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionUpMonTrigger)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionUpMonTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionUpMonTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Count, actual.Count, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Count")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionProcessCountNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionProcessCount)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionProcessCount)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionProcessCount or *AlertPolicyConditionsConditionProcessCount", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionProcessCount)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionProcessCount)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionProcessCount", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Process, actual.Process, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Process")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.User, actual.User, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("User")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Comparison, actual.Comparison, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Comparison")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProcessCountThreshold, actual.ProcessCountThreshold, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProcessCountThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Trigger, actual.Trigger, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionProcessCountTriggerNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionProcessCountTrigger, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Trigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Duration, actual.Duration, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionProcessCountDurationNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionProcessCountDuration, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Duration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionProcessCountTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionProcessCountTrigger)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionProcessCountTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionProcessCountTrigger or *AlertPolicyConditionsConditionProcessCountTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionProcessCountTrigger)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionProcessCountTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionProcessCountTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Count, actual.Count, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Count")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionProcessCountDurationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionProcessCountDuration)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionProcessCountDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionProcessCountDuration or *AlertPolicyConditionsConditionProcessCountDuration", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionProcessCountDuration)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionProcessCountDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionProcessCountDuration", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionTimeSeriesQueryLanguageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionTimeSeriesQueryLanguage)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionTimeSeriesQueryLanguage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionTimeSeriesQueryLanguage or *AlertPolicyConditionsConditionTimeSeriesQueryLanguage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionTimeSeriesQueryLanguage)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionTimeSeriesQueryLanguage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionTimeSeriesQueryLanguage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Query, actual.Query, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Query")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Summary, actual.Summary, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Summary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionMonitoringQueryLanguageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionMonitoringQueryLanguage)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionMonitoringQueryLanguage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionMonitoringQueryLanguage or *AlertPolicyConditionsConditionMonitoringQueryLanguage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionMonitoringQueryLanguage)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionMonitoringQueryLanguage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionMonitoringQueryLanguage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Query, actual.Query, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Query")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Duration, actual.Duration, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionMonitoringQueryLanguageDurationNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionMonitoringQueryLanguageDuration, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Duration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Trigger, actual.Trigger, dcl.Info{ObjectFunction: compareAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerNewStyle, EmptyObject: EmptyAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Trigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionMonitoringQueryLanguageDurationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionMonitoringQueryLanguageDuration)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionMonitoringQueryLanguageDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionMonitoringQueryLanguageDuration or *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionMonitoringQueryLanguageDuration)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionMonitoringQueryLanguageDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionMonitoringQueryLanguageDuration", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger or *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Count, actual.Count, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Count")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyEnabledNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyEnabled)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyEnabled)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyEnabled or *AlertPolicyEnabled", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyEnabled)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyEnabled)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyEnabled", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyValidityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyValidity)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyValidity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyValidity or *AlertPolicyValidity", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyValidity)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyValidity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyValidity", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Details, actual.Details, dcl.Info{ObjectFunction: compareAlertPolicyValidityDetailsNewStyle, EmptyObject: EmptyAlertPolicyValidityDetails, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Details")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyValidityDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyValidityDetails)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyValidityDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyValidityDetails or *AlertPolicyValidityDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyValidityDetails)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyValidityDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyValidityDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TypeUrl, actual.TypeUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TypeUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyCreationRecordNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyCreationRecord)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyCreationRecord)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyCreationRecord or *AlertPolicyCreationRecord", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyCreationRecord)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyCreationRecord)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyCreationRecord", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MutateTime, actual.MutateTime, dcl.Info{ObjectFunction: compareAlertPolicyCreationRecordMutateTimeNewStyle, EmptyObject: EmptyAlertPolicyCreationRecordMutateTime, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MutateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MutatedBy, actual.MutatedBy, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MutatedBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyCreationRecordMutateTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyCreationRecordMutateTime)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyCreationRecordMutateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyCreationRecordMutateTime or *AlertPolicyCreationRecordMutateTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyCreationRecordMutateTime)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyCreationRecordMutateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyCreationRecordMutateTime", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyMutationRecordNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyMutationRecord)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyMutationRecord)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyMutationRecord or *AlertPolicyMutationRecord", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyMutationRecord)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyMutationRecord)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyMutationRecord", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MutateTime, actual.MutateTime, dcl.Info{ObjectFunction: compareAlertPolicyMutationRecordMutateTimeNewStyle, EmptyObject: EmptyAlertPolicyMutationRecordMutateTime, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MutateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MutatedBy, actual.MutatedBy, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MutatedBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyMutationRecordMutateTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyMutationRecordMutateTime)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyMutationRecordMutateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyMutationRecordMutateTime or *AlertPolicyMutationRecordMutateTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyMutationRecordMutateTime)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyMutationRecordMutateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyMutationRecordMutateTime", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyIncidentStrategyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyIncidentStrategy)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyIncidentStrategy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyIncidentStrategy or *AlertPolicyIncidentStrategy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyIncidentStrategy)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyIncidentStrategy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyIncidentStrategy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAlertPolicyMetadataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AlertPolicyMetadata)
	if !ok {
		desiredNotPointer, ok := d.(AlertPolicyMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyMetadata or *AlertPolicyMetadata", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AlertPolicyMetadata)
	if !ok {
		actualNotPointer, ok := a.(AlertPolicyMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AlertPolicyMetadata", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SloNames, actual.SloNames, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SloNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *AlertPolicy) urlNormalized() *AlertPolicy {
	normalized := dcl.Copy(*r).(AlertPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *AlertPolicy) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *AlertPolicy) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *AlertPolicy) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *AlertPolicy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateAlertPolicy" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("v3/projects/{{project}}/alertPolicies/{{name}}", "https://monitoring.googleapis.com/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AlertPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AlertPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandAlertPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AlertPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAlertPolicy decodes JSON responses into the AlertPolicy resource schema.
func unmarshalAlertPolicy(b []byte, c *Client) (*AlertPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAlertPolicy(m, c)
}

func unmarshalMapAlertPolicy(m map[string]interface{}, c *Client) (*AlertPolicy, error) {

	return flattenAlertPolicy(c, m), nil
}

// expandAlertPolicy expands AlertPolicy into a JSON request object.
func expandAlertPolicy(c *Client, f *AlertPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v, err := expandAlertPolicyDocumentation(c, f.Documentation); err != nil {
		return nil, fmt.Errorf("error expanding Documentation into documentation: %w", err)
	} else if v != nil {
		m["documentation"] = v
	}
	if v := f.UserLabels; !dcl.IsEmptyValueIndirect(v) {
		m["userLabels"] = v
	}
	if v, err := expandAlertPolicyConditionsSlice(c, f.Conditions); err != nil {
		return nil, fmt.Errorf("error expanding Conditions into conditions: %w", err)
	} else if v != nil {
		m["conditions"] = v
	}
	if v := f.Combiner; !dcl.IsEmptyValueIndirect(v) {
		m["combiner"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}
	if v, err := expandAlertPolicyEnabled(c, f.Enabled); err != nil {
		return nil, fmt.Errorf("error expanding Enabled into enabled: %w", err)
	} else if v != nil {
		m["enabled"] = v
	}
	if v, err := expandAlertPolicyValidity(c, f.Validity); err != nil {
		return nil, fmt.Errorf("error expanding Validity into validity: %w", err)
	} else if v != nil {
		m["validity"] = v
	}
	if v := f.NotificationChannels; !dcl.IsEmptyValueIndirect(v) {
		m["notificationChannels"] = v
	}
	if v, err := expandAlertPolicyCreationRecord(c, f.CreationRecord); err != nil {
		return nil, fmt.Errorf("error expanding CreationRecord into creationRecord: %w", err)
	} else if v != nil {
		m["creationRecord"] = v
	}
	if v, err := expandAlertPolicyMutationRecord(c, f.MutationRecord); err != nil {
		return nil, fmt.Errorf("error expanding MutationRecord into mutationRecord: %w", err)
	} else if v != nil {
		m["mutationRecord"] = v
	}
	if v, err := expandAlertPolicyIncidentStrategy(c, f.IncidentStrategy); err != nil {
		return nil, fmt.Errorf("error expanding IncidentStrategy into incidentStrategy: %w", err)
	} else if v != nil {
		m["incidentStrategy"] = v
	}
	if v, err := expandAlertPolicyMetadata(c, f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if v != nil {
		m["metadata"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenAlertPolicy flattens AlertPolicy from a JSON request object into the
// AlertPolicy type.
func flattenAlertPolicy(c *Client, i interface{}) *AlertPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &AlertPolicy{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.Documentation = flattenAlertPolicyDocumentation(c, m["documentation"])
	res.UserLabels = dcl.FlattenKeyValuePairs(m["userLabels"])
	res.Conditions = flattenAlertPolicyConditionsSlice(c, m["conditions"])
	res.Combiner = flattenAlertPolicyCombinerEnum(m["combiner"])
	res.Disabled = dcl.FlattenBool(m["disabled"])
	res.Enabled = flattenAlertPolicyEnabled(c, m["enabled"])
	res.Validity = flattenAlertPolicyValidity(c, m["validity"])
	res.NotificationChannels = dcl.FlattenStringSlice(m["notificationChannels"])
	res.CreationRecord = flattenAlertPolicyCreationRecord(c, m["creationRecord"])
	res.MutationRecord = flattenAlertPolicyMutationRecord(c, m["mutationRecord"])
	res.IncidentStrategy = flattenAlertPolicyIncidentStrategy(c, m["incidentStrategy"])
	res.Metadata = flattenAlertPolicyMetadata(c, m["metadata"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandAlertPolicyDocumentationMap expands the contents of AlertPolicyDocumentation into a JSON
// request object.
func expandAlertPolicyDocumentationMap(c *Client, f map[string]AlertPolicyDocumentation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyDocumentation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyDocumentationSlice expands the contents of AlertPolicyDocumentation into a JSON
// request object.
func expandAlertPolicyDocumentationSlice(c *Client, f []AlertPolicyDocumentation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyDocumentation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyDocumentationMap flattens the contents of AlertPolicyDocumentation from a JSON
// response object.
func flattenAlertPolicyDocumentationMap(c *Client, i interface{}) map[string]AlertPolicyDocumentation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyDocumentation{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyDocumentation{}
	}

	items := make(map[string]AlertPolicyDocumentation)
	for k, item := range a {
		items[k] = *flattenAlertPolicyDocumentation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyDocumentationSlice flattens the contents of AlertPolicyDocumentation from a JSON
// response object.
func flattenAlertPolicyDocumentationSlice(c *Client, i interface{}) []AlertPolicyDocumentation {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyDocumentation{}
	}

	if len(a) == 0 {
		return []AlertPolicyDocumentation{}
	}

	items := make([]AlertPolicyDocumentation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyDocumentation(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyDocumentation expands an instance of AlertPolicyDocumentation into a JSON
// request object.
func expandAlertPolicyDocumentation(c *Client, f *AlertPolicyDocumentation) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}
	if v := f.MimeType; !dcl.IsEmptyValueIndirect(v) {
		m["mimeType"] = v
	}

	return m, nil
}

// flattenAlertPolicyDocumentation flattens an instance of AlertPolicyDocumentation from a JSON
// response object.
func flattenAlertPolicyDocumentation(c *Client, i interface{}) *AlertPolicyDocumentation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyDocumentation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyDocumentation
	}
	r.Content = dcl.FlattenString(m["content"])
	r.MimeType = dcl.FlattenString(m["mimeType"])

	return r
}

// expandAlertPolicyConditionsMap expands the contents of AlertPolicyConditions into a JSON
// request object.
func expandAlertPolicyConditionsMap(c *Client, f map[string]AlertPolicyConditions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsSlice expands the contents of AlertPolicyConditions into a JSON
// request object.
func expandAlertPolicyConditionsSlice(c *Client, f []AlertPolicyConditions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsMap flattens the contents of AlertPolicyConditions from a JSON
// response object.
func flattenAlertPolicyConditionsMap(c *Client, i interface{}) map[string]AlertPolicyConditions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditions{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditions{}
	}

	items := make(map[string]AlertPolicyConditions)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsSlice flattens the contents of AlertPolicyConditions from a JSON
// response object.
func flattenAlertPolicyConditionsSlice(c *Client, i interface{}) []AlertPolicyConditions {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditions{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditions{}
	}

	items := make([]AlertPolicyConditions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditions(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditions expands an instance of AlertPolicyConditions into a JSON
// request object.
func expandAlertPolicyConditions(c *Client, f *AlertPolicyConditions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.ResourceStateFilter; !dcl.IsEmptyValueIndirect(v) {
		m["resourceStateFilter"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThreshold(c, f.ConditionThreshold); err != nil {
		return nil, fmt.Errorf("error expanding ConditionThreshold into conditionThreshold: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditionThreshold"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionAbsent(c, f.ConditionAbsent); err != nil {
		return nil, fmt.Errorf("error expanding ConditionAbsent into conditionAbsent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditionAbsent"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionMatchedLog(c, f.ConditionMatchedLog); err != nil {
		return nil, fmt.Errorf("error expanding ConditionMatchedLog into conditionMatchedLog: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditionMatchedLog"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionClusterOutlier(c, f.ConditionClusterOutlier); err != nil {
		return nil, fmt.Errorf("error expanding ConditionClusterOutlier into conditionClusterOutlier: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditionClusterOutlier"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionRate(c, f.ConditionRate); err != nil {
		return nil, fmt.Errorf("error expanding ConditionRate into conditionRate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditionRate"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionUpMon(c, f.ConditionUpMon); err != nil {
		return nil, fmt.Errorf("error expanding ConditionUpMon into conditionUpMon: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditionUpMon"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionProcessCount(c, f.ConditionProcessCount); err != nil {
		return nil, fmt.Errorf("error expanding ConditionProcessCount into conditionProcessCount: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditionProcessCount"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, f.ConditionTimeSeriesQueryLanguage); err != nil {
		return nil, fmt.Errorf("error expanding ConditionTimeSeriesQueryLanguage into conditionTimeSeriesQueryLanguage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditionTimeSeriesQueryLanguage"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionMonitoringQueryLanguage(c, f.ConditionMonitoringQueryLanguage); err != nil {
		return nil, fmt.Errorf("error expanding ConditionMonitoringQueryLanguage into conditionMonitoringQueryLanguage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditionMonitoringQueryLanguage"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditions flattens an instance of AlertPolicyConditions from a JSON
// response object.
func flattenAlertPolicyConditions(c *Client, i interface{}) *AlertPolicyConditions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditions
	}
	r.Name = dcl.FlattenString(m["name"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.ResourceStateFilter = flattenAlertPolicyConditionsResourceStateFilterEnum(m["resourceStateFilter"])
	r.ConditionThreshold = flattenAlertPolicyConditionsConditionThreshold(c, m["conditionThreshold"])
	r.ConditionAbsent = flattenAlertPolicyConditionsConditionAbsent(c, m["conditionAbsent"])
	r.ConditionMatchedLog = flattenAlertPolicyConditionsConditionMatchedLog(c, m["conditionMatchedLog"])
	r.ConditionClusterOutlier = flattenAlertPolicyConditionsConditionClusterOutlier(c, m["conditionClusterOutlier"])
	r.ConditionRate = flattenAlertPolicyConditionsConditionRate(c, m["conditionRate"])
	r.ConditionUpMon = flattenAlertPolicyConditionsConditionUpMon(c, m["conditionUpMon"])
	r.ConditionProcessCount = flattenAlertPolicyConditionsConditionProcessCount(c, m["conditionProcessCount"])
	r.ConditionTimeSeriesQueryLanguage = flattenAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, m["conditionTimeSeriesQueryLanguage"])
	r.ConditionMonitoringQueryLanguage = flattenAlertPolicyConditionsConditionMonitoringQueryLanguage(c, m["conditionMonitoringQueryLanguage"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdMap expands the contents of AlertPolicyConditionsConditionThreshold into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdMap(c *Client, f map[string]AlertPolicyConditionsConditionThreshold) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThreshold(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdSlice expands the contents of AlertPolicyConditionsConditionThreshold into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdSlice(c *Client, f []AlertPolicyConditionsConditionThreshold) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThreshold(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdMap flattens the contents of AlertPolicyConditionsConditionThreshold from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThreshold {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThreshold{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThreshold{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThreshold)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThreshold(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdSlice flattens the contents of AlertPolicyConditionsConditionThreshold from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThreshold {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThreshold{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThreshold{}
	}

	items := make([]AlertPolicyConditionsConditionThreshold, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThreshold(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThreshold expands an instance of AlertPolicyConditionsConditionThreshold into a JSON
// request object.
func expandAlertPolicyConditionsConditionThreshold(c *Client, f *AlertPolicyConditionsConditionThreshold) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdAggregationsSlice(c, f.Aggregations); err != nil {
		return nil, fmt.Errorf("error expanding Aggregations into aggregations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aggregations"] = v
	}
	if v := f.DenominatorFilter; !dcl.IsEmptyValueIndirect(v) {
		m["denominatorFilter"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice(c, f.DenominatorAggregations); err != nil {
		return nil, fmt.Errorf("error expanding DenominatorAggregations into denominatorAggregations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["denominatorAggregations"] = v
	}
	if v := f.Comparison; !dcl.IsEmptyValueIndirect(v) {
		m["comparison"] = v
	}
	if v := f.ThresholdValue; !dcl.IsEmptyValueIndirect(v) {
		m["thresholdValue"] = v
	}
	if v := f.Duration; !dcl.IsEmptyValueIndirect(v) {
		m["duration"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdTrigger(c, f.Trigger); err != nil {
		return nil, fmt.Errorf("error expanding Trigger into trigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["trigger"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThreshold flattens an instance of AlertPolicyConditionsConditionThreshold from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThreshold(c *Client, i interface{}) *AlertPolicyConditionsConditionThreshold {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThreshold{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThreshold
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.Aggregations = flattenAlertPolicyConditionsConditionThresholdAggregationsSlice(c, m["aggregations"])
	r.DenominatorFilter = dcl.FlattenString(m["denominatorFilter"])
	r.DenominatorAggregations = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice(c, m["denominatorAggregations"])
	r.Comparison = flattenAlertPolicyConditionsConditionThresholdComparisonEnum(m["comparison"])
	r.ThresholdValue = dcl.FlattenDouble(m["thresholdValue"])
	r.Duration = dcl.FlattenString(m["duration"])
	r.Trigger = flattenAlertPolicyConditionsConditionThresholdTrigger(c, m["trigger"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdAggregationsMap expands the contents of AlertPolicyConditionsConditionThresholdAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdAggregations) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregations(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdAggregationsSlice expands the contents of AlertPolicyConditionsConditionThresholdAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdAggregations) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregations(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsMap flattens the contents of AlertPolicyConditionsConditionThresholdAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdAggregations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdAggregations{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdAggregations{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdAggregations)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdAggregations(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregations {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregations{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregations{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregations(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdAggregations expands an instance of AlertPolicyConditionsConditionThresholdAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregations(c *Client, f *AlertPolicyConditionsConditionThresholdAggregations) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AlignmentPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["alignmentPeriod"] = v
	}
	if v := f.PerSeriesAligner; !dcl.IsEmptyValueIndirect(v) {
		m["perSeriesAligner"] = v
	}
	if v := f.CrossSeriesReducer; !dcl.IsEmptyValueIndirect(v) {
		m["crossSeriesReducer"] = v
	}
	if v := f.GroupByFields; !dcl.IsEmptyValueIndirect(v) {
		m["groupByFields"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, f.ReduceFractionLessThanParams); err != nil {
		return nil, fmt.Errorf("error expanding ReduceFractionLessThanParams into reduceFractionLessThanParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reduceFractionLessThanParams"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, f.ReduceMakeDistributionParams); err != nil {
		return nil, fmt.Errorf("error expanding ReduceMakeDistributionParams into reduceMakeDistributionParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reduceMakeDistributionParams"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregations flattens an instance of AlertPolicyConditionsConditionThresholdAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregations(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdAggregations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdAggregations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdAggregations
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])
	r.ReduceFractionLessThanParams = flattenAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, m["reduceFractionLessThanParams"])
	r.ReduceMakeDistributionParams = flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, m["reduceMakeDistributionParams"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsMap expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsSlice expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsMap flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams expands an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c *Client, f *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Threshold; !dcl.IsEmptyValueIndirect(v) {
		m["threshold"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams flattens an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams
	}
	r.Threshold = dcl.FlattenDouble(m["threshold"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsMap expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsSlice expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsMap flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams expands an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c *Client, f *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, f.BucketOptions); err != nil {
		return nil, fmt.Errorf("error expanding BucketOptions into bucketOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bucketOptions"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, f.ExemplarSampling); err != nil {
		return nil, fmt.Errorf("error expanding ExemplarSampling into exemplarSampling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exemplarSampling"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams flattens an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams
	}
	r.BucketOptions = flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, m["bucketOptions"])
	r.ExemplarSampling = flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, m["exemplarSampling"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsMap expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsSlice expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsMap flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions expands an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, f *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, f.LinearBuckets); err != nil {
		return nil, fmt.Errorf("error expanding LinearBuckets into linearBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["linearBuckets"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, f.ExponentialBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExponentialBuckets into exponentialBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exponentialBuckets"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, f.ExplicitBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExplicitBuckets into explicitBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["explicitBuckets"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions flattens an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions
	}
	r.LinearBuckets = flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, m["linearBuckets"])
	r.ExponentialBuckets = flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, m["exponentialBuckets"])
	r.ExplicitBuckets = flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, m["explicitBuckets"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets expands an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, f *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.Width; !dcl.IsEmptyValueIndirect(v) {
		m["width"] = v
	}
	if v := f.Offset; !dcl.IsEmptyValueIndirect(v) {
		m["offset"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets flattens an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.Width = dcl.FlattenDouble(m["width"])
	r.Offset = dcl.FlattenDouble(m["offset"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets expands an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, f *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.GrowthFactor; !dcl.IsEmptyValueIndirect(v) {
		m["growthFactor"] = v
	}
	if v := f.Scale; !dcl.IsEmptyValueIndirect(v) {
		m["scale"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets flattens an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.GrowthFactor = dcl.FlattenDouble(m["growthFactor"])
	r.Scale = dcl.FlattenDouble(m["scale"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets expands an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, f *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Bounds; !dcl.IsEmptyValueIndirect(v) {
		m["bounds"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets flattens an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	}
	r.Bounds = dcl.FlattenFloatSlice(m["bounds"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingMap expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingSlice expands the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, f []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingMap flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling expands an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, f *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinimumValue; !dcl.IsEmptyValueIndirect(v) {
		m["minimumValue"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling flattens an instance of AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling
	}
	r.MinimumValue = dcl.FlattenDouble(m["minimumValue"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsMap expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregations) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregations(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdDenominatorAggregations) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregations(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsMap flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregations{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregations{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregations)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregations(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregations {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregations{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregations{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregations(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregations expands an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregations(c *Client, f *AlertPolicyConditionsConditionThresholdDenominatorAggregations) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AlignmentPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["alignmentPeriod"] = v
	}
	if v := f.PerSeriesAligner; !dcl.IsEmptyValueIndirect(v) {
		m["perSeriesAligner"] = v
	}
	if v := f.CrossSeriesReducer; !dcl.IsEmptyValueIndirect(v) {
		m["crossSeriesReducer"] = v
	}
	if v := f.GroupByFields; !dcl.IsEmptyValueIndirect(v) {
		m["groupByFields"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, f.ReduceFractionLessThanParams); err != nil {
		return nil, fmt.Errorf("error expanding ReduceFractionLessThanParams into reduceFractionLessThanParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reduceFractionLessThanParams"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, f.ReduceMakeDistributionParams); err != nil {
		return nil, fmt.Errorf("error expanding ReduceMakeDistributionParams into reduceMakeDistributionParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reduceMakeDistributionParams"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregations flattens an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregations(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdDenominatorAggregations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregations
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])
	r.ReduceFractionLessThanParams = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, m["reduceFractionLessThanParams"])
	r.ReduceMakeDistributionParams = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, m["reduceMakeDistributionParams"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsMap expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsSlice expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsMap flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams expands an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c *Client, f *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Threshold; !dcl.IsEmptyValueIndirect(v) {
		m["threshold"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams flattens an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams
	}
	r.Threshold = dcl.FlattenDouble(m["threshold"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsMap expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsSlice expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsMap flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams expands an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c *Client, f *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, f.BucketOptions); err != nil {
		return nil, fmt.Errorf("error expanding BucketOptions into bucketOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bucketOptions"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, f.ExemplarSampling); err != nil {
		return nil, fmt.Errorf("error expanding ExemplarSampling into exemplarSampling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exemplarSampling"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams flattens an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams
	}
	r.BucketOptions = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, m["bucketOptions"])
	r.ExemplarSampling = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, m["exemplarSampling"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsMap expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsSlice expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsMap flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions expands an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, f *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, f.LinearBuckets); err != nil {
		return nil, fmt.Errorf("error expanding LinearBuckets into linearBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["linearBuckets"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, f.ExponentialBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExponentialBuckets into exponentialBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exponentialBuckets"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, f.ExplicitBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExplicitBuckets into explicitBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["explicitBuckets"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions flattens an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions
	}
	r.LinearBuckets = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, m["linearBuckets"])
	r.ExponentialBuckets = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, m["exponentialBuckets"])
	r.ExplicitBuckets = flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, m["explicitBuckets"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets expands an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, f *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.Width; !dcl.IsEmptyValueIndirect(v) {
		m["width"] = v
	}
	if v := f.Offset; !dcl.IsEmptyValueIndirect(v) {
		m["offset"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets flattens an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.Width = dcl.FlattenDouble(m["width"])
	r.Offset = dcl.FlattenDouble(m["offset"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets expands an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, f *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.GrowthFactor; !dcl.IsEmptyValueIndirect(v) {
		m["growthFactor"] = v
	}
	if v := f.Scale; !dcl.IsEmptyValueIndirect(v) {
		m["scale"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets flattens an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.GrowthFactor = dcl.FlattenDouble(m["growthFactor"])
	r.Scale = dcl.FlattenDouble(m["scale"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, f []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets expands an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, f *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Bounds; !dcl.IsEmptyValueIndirect(v) {
		m["bounds"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets flattens an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	}
	r.Bounds = dcl.FlattenFloatSlice(m["bounds"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingMap expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingSlice expands the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, f []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingMap flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling expands an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, f *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinimumValue; !dcl.IsEmptyValueIndirect(v) {
		m["minimumValue"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling flattens an instance of AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling
	}
	r.MinimumValue = dcl.FlattenDouble(m["minimumValue"])

	return r
}

// expandAlertPolicyConditionsConditionThresholdTriggerMap expands the contents of AlertPolicyConditionsConditionThresholdTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdTriggerMap(c *Client, f map[string]AlertPolicyConditionsConditionThresholdTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionThresholdTriggerSlice expands the contents of AlertPolicyConditionsConditionThresholdTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdTriggerSlice(c *Client, f []AlertPolicyConditionsConditionThresholdTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionThresholdTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionThresholdTriggerMap flattens the contents of AlertPolicyConditionsConditionThresholdTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdTriggerMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionThresholdTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionThresholdTrigger{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionThresholdTrigger{}
	}

	items := make(map[string]AlertPolicyConditionsConditionThresholdTrigger)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionThresholdTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdTriggerSlice flattens the contents of AlertPolicyConditionsConditionThresholdTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdTriggerSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdTrigger{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdTrigger{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionThresholdTrigger expands an instance of AlertPolicyConditionsConditionThresholdTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionThresholdTrigger(c *Client, f *AlertPolicyConditionsConditionThresholdTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Count; !dcl.IsEmptyValueIndirect(v) {
		m["count"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionThresholdTrigger flattens an instance of AlertPolicyConditionsConditionThresholdTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdTrigger(c *Client, i interface{}) *AlertPolicyConditionsConditionThresholdTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionThresholdTrigger{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionThresholdTrigger
	}
	r.Count = dcl.FlattenInteger(m["count"])
	r.Percent = dcl.FlattenDouble(m["percent"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentMap expands the contents of AlertPolicyConditionsConditionAbsent into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsent) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsent(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentSlice expands the contents of AlertPolicyConditionsConditionAbsent into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentSlice(c *Client, f []AlertPolicyConditionsConditionAbsent) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsent(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentMap flattens the contents of AlertPolicyConditionsConditionAbsent from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsent {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsent{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsent{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsent)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsent(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentSlice flattens the contents of AlertPolicyConditionsConditionAbsent from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsent {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsent{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsent{}
	}

	items := make([]AlertPolicyConditionsConditionAbsent, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsent(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsent expands an instance of AlertPolicyConditionsConditionAbsent into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsent(c *Client, f *AlertPolicyConditionsConditionAbsent) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionAbsentAggregationsSlice(c, f.Aggregations); err != nil {
		return nil, fmt.Errorf("error expanding Aggregations into aggregations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aggregations"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionAbsentDuration(c, f.Duration); err != nil {
		return nil, fmt.Errorf("error expanding Duration into duration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["duration"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionAbsentTrigger(c, f.Trigger); err != nil {
		return nil, fmt.Errorf("error expanding Trigger into trigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["trigger"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsent flattens an instance of AlertPolicyConditionsConditionAbsent from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsent(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsent {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsent{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsent
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.Aggregations = flattenAlertPolicyConditionsConditionAbsentAggregationsSlice(c, m["aggregations"])
	r.Duration = flattenAlertPolicyConditionsConditionAbsentDuration(c, m["duration"])
	r.Trigger = flattenAlertPolicyConditionsConditionAbsentTrigger(c, m["trigger"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentAggregationsMap expands the contents of AlertPolicyConditionsConditionAbsentAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentAggregations) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregations(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentAggregationsSlice expands the contents of AlertPolicyConditionsConditionAbsentAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsSlice(c *Client, f []AlertPolicyConditionsConditionAbsentAggregations) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregations(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsMap flattens the contents of AlertPolicyConditionsConditionAbsentAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentAggregations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentAggregations{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentAggregations{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentAggregations)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentAggregations(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregations {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregations{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregations{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregations(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentAggregations expands an instance of AlertPolicyConditionsConditionAbsentAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregations(c *Client, f *AlertPolicyConditionsConditionAbsentAggregations) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AlignmentPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["alignmentPeriod"] = v
	}
	if v := f.PerSeriesAligner; !dcl.IsEmptyValueIndirect(v) {
		m["perSeriesAligner"] = v
	}
	if v := f.CrossSeriesReducer; !dcl.IsEmptyValueIndirect(v) {
		m["crossSeriesReducer"] = v
	}
	if v := f.GroupByFields; !dcl.IsEmptyValueIndirect(v) {
		m["groupByFields"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, f.ReduceFractionLessThanParams); err != nil {
		return nil, fmt.Errorf("error expanding ReduceFractionLessThanParams into reduceFractionLessThanParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reduceFractionLessThanParams"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, f.ReduceMakeDistributionParams); err != nil {
		return nil, fmt.Errorf("error expanding ReduceMakeDistributionParams into reduceMakeDistributionParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reduceMakeDistributionParams"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregations flattens an instance of AlertPolicyConditionsConditionAbsentAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregations(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentAggregations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentAggregations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentAggregations
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])
	r.ReduceFractionLessThanParams = flattenAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, m["reduceFractionLessThanParams"])
	r.ReduceMakeDistributionParams = flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, m["reduceMakeDistributionParams"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsMap expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsSlice expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsSlice(c *Client, f []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsMap flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams expands an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c *Client, f *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Threshold; !dcl.IsEmptyValueIndirect(v) {
		m["threshold"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams flattens an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams
	}
	r.Threshold = dcl.FlattenDouble(m["threshold"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsMap expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsSlice expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsSlice(c *Client, f []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsMap flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams expands an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c *Client, f *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, f.BucketOptions); err != nil {
		return nil, fmt.Errorf("error expanding BucketOptions into bucketOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bucketOptions"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, f.ExemplarSampling); err != nil {
		return nil, fmt.Errorf("error expanding ExemplarSampling into exemplarSampling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exemplarSampling"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams flattens an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams
	}
	r.BucketOptions = flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, m["bucketOptions"])
	r.ExemplarSampling = flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, m["exemplarSampling"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsMap expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsSlice expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, f []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsMap flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions expands an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, f *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, f.LinearBuckets); err != nil {
		return nil, fmt.Errorf("error expanding LinearBuckets into linearBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["linearBuckets"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, f.ExponentialBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExponentialBuckets into exponentialBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exponentialBuckets"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, f.ExplicitBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExplicitBuckets into explicitBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["explicitBuckets"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions flattens an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions
	}
	r.LinearBuckets = flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, m["linearBuckets"])
	r.ExponentialBuckets = flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, m["exponentialBuckets"])
	r.ExplicitBuckets = flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, m["explicitBuckets"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, f []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets expands an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, f *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.Width; !dcl.IsEmptyValueIndirect(v) {
		m["width"] = v
	}
	if v := f.Offset; !dcl.IsEmptyValueIndirect(v) {
		m["offset"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets flattens an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.Width = dcl.FlattenDouble(m["width"])
	r.Offset = dcl.FlattenDouble(m["offset"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, f []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets expands an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, f *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.GrowthFactor; !dcl.IsEmptyValueIndirect(v) {
		m["growthFactor"] = v
	}
	if v := f.Scale; !dcl.IsEmptyValueIndirect(v) {
		m["scale"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets flattens an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.GrowthFactor = dcl.FlattenDouble(m["growthFactor"])
	r.Scale = dcl.FlattenDouble(m["scale"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, f []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets expands an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, f *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Bounds; !dcl.IsEmptyValueIndirect(v) {
		m["bounds"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets flattens an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	}
	r.Bounds = dcl.FlattenFloatSlice(m["bounds"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingMap expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingSlice expands the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, f []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingMap flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling expands an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, f *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinimumValue; !dcl.IsEmptyValueIndirect(v) {
		m["minimumValue"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling flattens an instance of AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling
	}
	r.MinimumValue = dcl.FlattenDouble(m["minimumValue"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentDurationMap expands the contents of AlertPolicyConditionsConditionAbsentDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentDurationMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentDuration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentDuration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentDurationSlice expands the contents of AlertPolicyConditionsConditionAbsentDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentDurationSlice(c *Client, f []AlertPolicyConditionsConditionAbsentDuration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentDuration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentDurationMap flattens the contents of AlertPolicyConditionsConditionAbsentDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentDurationMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentDuration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentDuration{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentDuration{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentDuration)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentDuration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentDurationSlice flattens the contents of AlertPolicyConditionsConditionAbsentDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentDurationSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentDuration {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentDuration{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentDuration{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentDuration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentDuration(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentDuration expands an instance of AlertPolicyConditionsConditionAbsentDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentDuration(c *Client, f *AlertPolicyConditionsConditionAbsentDuration) (map[string]interface{}, error) {
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

// flattenAlertPolicyConditionsConditionAbsentDuration flattens an instance of AlertPolicyConditionsConditionAbsentDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentDuration(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentDuration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentDuration{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentDuration
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandAlertPolicyConditionsConditionAbsentTriggerMap expands the contents of AlertPolicyConditionsConditionAbsentTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentTriggerMap(c *Client, f map[string]AlertPolicyConditionsConditionAbsentTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionAbsentTriggerSlice expands the contents of AlertPolicyConditionsConditionAbsentTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentTriggerSlice(c *Client, f []AlertPolicyConditionsConditionAbsentTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionAbsentTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionAbsentTriggerMap flattens the contents of AlertPolicyConditionsConditionAbsentTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentTriggerMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionAbsentTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionAbsentTrigger{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionAbsentTrigger{}
	}

	items := make(map[string]AlertPolicyConditionsConditionAbsentTrigger)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionAbsentTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentTriggerSlice flattens the contents of AlertPolicyConditionsConditionAbsentTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentTriggerSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentTrigger{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentTrigger{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionAbsentTrigger expands an instance of AlertPolicyConditionsConditionAbsentTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionAbsentTrigger(c *Client, f *AlertPolicyConditionsConditionAbsentTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Count; !dcl.IsEmptyValueIndirect(v) {
		m["count"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionAbsentTrigger flattens an instance of AlertPolicyConditionsConditionAbsentTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentTrigger(c *Client, i interface{}) *AlertPolicyConditionsConditionAbsentTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionAbsentTrigger{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionAbsentTrigger
	}
	r.Count = dcl.FlattenInteger(m["count"])
	r.Percent = dcl.FlattenDouble(m["percent"])

	return r
}

// expandAlertPolicyConditionsConditionMatchedLogMap expands the contents of AlertPolicyConditionsConditionMatchedLog into a JSON
// request object.
func expandAlertPolicyConditionsConditionMatchedLogMap(c *Client, f map[string]AlertPolicyConditionsConditionMatchedLog) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionMatchedLog(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionMatchedLogSlice expands the contents of AlertPolicyConditionsConditionMatchedLog into a JSON
// request object.
func expandAlertPolicyConditionsConditionMatchedLogSlice(c *Client, f []AlertPolicyConditionsConditionMatchedLog) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionMatchedLog(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionMatchedLogMap flattens the contents of AlertPolicyConditionsConditionMatchedLog from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMatchedLogMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionMatchedLog {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionMatchedLog{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionMatchedLog{}
	}

	items := make(map[string]AlertPolicyConditionsConditionMatchedLog)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionMatchedLog(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionMatchedLogSlice flattens the contents of AlertPolicyConditionsConditionMatchedLog from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMatchedLogSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionMatchedLog {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionMatchedLog{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionMatchedLog{}
	}

	items := make([]AlertPolicyConditionsConditionMatchedLog, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionMatchedLog(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionMatchedLog expands an instance of AlertPolicyConditionsConditionMatchedLog into a JSON
// request object.
func expandAlertPolicyConditionsConditionMatchedLog(c *Client, f *AlertPolicyConditionsConditionMatchedLog) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v := f.LabelExtractors; !dcl.IsEmptyValueIndirect(v) {
		m["labelExtractors"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionMatchedLog flattens an instance of AlertPolicyConditionsConditionMatchedLog from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMatchedLog(c *Client, i interface{}) *AlertPolicyConditionsConditionMatchedLog {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionMatchedLog{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionMatchedLog
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.LabelExtractors = dcl.FlattenKeyValuePairs(m["labelExtractors"])

	return r
}

// expandAlertPolicyConditionsConditionClusterOutlierMap expands the contents of AlertPolicyConditionsConditionClusterOutlier into a JSON
// request object.
func expandAlertPolicyConditionsConditionClusterOutlierMap(c *Client, f map[string]AlertPolicyConditionsConditionClusterOutlier) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionClusterOutlier(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionClusterOutlierSlice expands the contents of AlertPolicyConditionsConditionClusterOutlier into a JSON
// request object.
func expandAlertPolicyConditionsConditionClusterOutlierSlice(c *Client, f []AlertPolicyConditionsConditionClusterOutlier) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionClusterOutlier(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionClusterOutlierMap flattens the contents of AlertPolicyConditionsConditionClusterOutlier from a JSON
// response object.
func flattenAlertPolicyConditionsConditionClusterOutlierMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionClusterOutlier {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionClusterOutlier{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionClusterOutlier{}
	}

	items := make(map[string]AlertPolicyConditionsConditionClusterOutlier)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionClusterOutlier(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionClusterOutlierSlice flattens the contents of AlertPolicyConditionsConditionClusterOutlier from a JSON
// response object.
func flattenAlertPolicyConditionsConditionClusterOutlierSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionClusterOutlier {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionClusterOutlier{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionClusterOutlier{}
	}

	items := make([]AlertPolicyConditionsConditionClusterOutlier, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionClusterOutlier(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionClusterOutlier expands an instance of AlertPolicyConditionsConditionClusterOutlier into a JSON
// request object.
func expandAlertPolicyConditionsConditionClusterOutlier(c *Client, f *AlertPolicyConditionsConditionClusterOutlier) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionClusterOutlier flattens an instance of AlertPolicyConditionsConditionClusterOutlier from a JSON
// response object.
func flattenAlertPolicyConditionsConditionClusterOutlier(c *Client, i interface{}) *AlertPolicyConditionsConditionClusterOutlier {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionClusterOutlier{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionClusterOutlier
	}
	r.Filter = dcl.FlattenString(m["filter"])

	return r
}

// expandAlertPolicyConditionsConditionRateMap expands the contents of AlertPolicyConditionsConditionRate into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateMap(c *Client, f map[string]AlertPolicyConditionsConditionRate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRate(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateSlice expands the contents of AlertPolicyConditionsConditionRate into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateSlice(c *Client, f []AlertPolicyConditionsConditionRate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRate(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateMap flattens the contents of AlertPolicyConditionsConditionRate from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRate{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRate{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRate)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRate(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateSlice flattens the contents of AlertPolicyConditionsConditionRate from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRate {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRate{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRate{}
	}

	items := make([]AlertPolicyConditionsConditionRate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRate(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRate expands an instance of AlertPolicyConditionsConditionRate into a JSON
// request object.
func expandAlertPolicyConditionsConditionRate(c *Client, f *AlertPolicyConditionsConditionRate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionRateAggregationsSlice(c, f.Aggregations); err != nil {
		return nil, fmt.Errorf("error expanding Aggregations into aggregations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aggregations"] = v
	}
	if v := f.Comparison; !dcl.IsEmptyValueIndirect(v) {
		m["comparison"] = v
	}
	if v := f.ThresholdValue; !dcl.IsEmptyValueIndirect(v) {
		m["thresholdValue"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionRateTimeWindow(c, f.TimeWindow); err != nil {
		return nil, fmt.Errorf("error expanding TimeWindow into timeWindow: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeWindow"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionRateTrigger(c, f.Trigger); err != nil {
		return nil, fmt.Errorf("error expanding Trigger into trigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["trigger"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRate flattens an instance of AlertPolicyConditionsConditionRate from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRate(c *Client, i interface{}) *AlertPolicyConditionsConditionRate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRate
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.Aggregations = flattenAlertPolicyConditionsConditionRateAggregationsSlice(c, m["aggregations"])
	r.Comparison = flattenAlertPolicyConditionsConditionRateComparisonEnum(m["comparison"])
	r.ThresholdValue = dcl.FlattenDouble(m["thresholdValue"])
	r.TimeWindow = flattenAlertPolicyConditionsConditionRateTimeWindow(c, m["timeWindow"])
	r.Trigger = flattenAlertPolicyConditionsConditionRateTrigger(c, m["trigger"])

	return r
}

// expandAlertPolicyConditionsConditionRateAggregationsMap expands the contents of AlertPolicyConditionsConditionRateAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsMap(c *Client, f map[string]AlertPolicyConditionsConditionRateAggregations) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregations(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateAggregationsSlice expands the contents of AlertPolicyConditionsConditionRateAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsSlice(c *Client, f []AlertPolicyConditionsConditionRateAggregations) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregations(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsMap flattens the contents of AlertPolicyConditionsConditionRateAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateAggregations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateAggregations{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateAggregations{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateAggregations)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateAggregations(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsSlice flattens the contents of AlertPolicyConditionsConditionRateAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregations {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregations{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregations{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregations(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateAggregations expands an instance of AlertPolicyConditionsConditionRateAggregations into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregations(c *Client, f *AlertPolicyConditionsConditionRateAggregations) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AlignmentPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["alignmentPeriod"] = v
	}
	if v := f.PerSeriesAligner; !dcl.IsEmptyValueIndirect(v) {
		m["perSeriesAligner"] = v
	}
	if v := f.CrossSeriesReducer; !dcl.IsEmptyValueIndirect(v) {
		m["crossSeriesReducer"] = v
	}
	if v := f.GroupByFields; !dcl.IsEmptyValueIndirect(v) {
		m["groupByFields"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, f.ReduceFractionLessThanParams); err != nil {
		return nil, fmt.Errorf("error expanding ReduceFractionLessThanParams into reduceFractionLessThanParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reduceFractionLessThanParams"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, f.ReduceMakeDistributionParams); err != nil {
		return nil, fmt.Errorf("error expanding ReduceMakeDistributionParams into reduceMakeDistributionParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reduceMakeDistributionParams"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRateAggregations flattens an instance of AlertPolicyConditionsConditionRateAggregations from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregations(c *Client, i interface{}) *AlertPolicyConditionsConditionRateAggregations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateAggregations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateAggregations
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])
	r.ReduceFractionLessThanParams = flattenAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, m["reduceFractionLessThanParams"])
	r.ReduceMakeDistributionParams = flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, m["reduceMakeDistributionParams"])

	return r
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsMap expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsMap(c *Client, f map[string]AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsSlice expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsSlice(c *Client, f []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsMap flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsSlice flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams expands an instance of AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c *Client, f *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Threshold; !dcl.IsEmptyValueIndirect(v) {
		m["threshold"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams flattens an instance of AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c *Client, i interface{}) *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams
	}
	r.Threshold = dcl.FlattenDouble(m["threshold"])

	return r
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsMap expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsMap(c *Client, f map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsSlice expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsSlice(c *Client, f []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsMap flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsSlice flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams expands an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c *Client, f *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, f.BucketOptions); err != nil {
		return nil, fmt.Errorf("error expanding BucketOptions into bucketOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bucketOptions"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, f.ExemplarSampling); err != nil {
		return nil, fmt.Errorf("error expanding ExemplarSampling into exemplarSampling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exemplarSampling"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams flattens an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c *Client, i interface{}) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams
	}
	r.BucketOptions = flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, m["bucketOptions"])
	r.ExemplarSampling = flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, m["exemplarSampling"])

	return r
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsMap expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsMap(c *Client, f map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsSlice expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, f []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsMap flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsSlice flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions expands an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, f *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, f.LinearBuckets); err != nil {
		return nil, fmt.Errorf("error expanding LinearBuckets into linearBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["linearBuckets"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, f.ExponentialBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExponentialBuckets into exponentialBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exponentialBuckets"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, f.ExplicitBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExplicitBuckets into explicitBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["explicitBuckets"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions flattens an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, i interface{}) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions
	}
	r.LinearBuckets = flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, m["linearBuckets"])
	r.ExponentialBuckets = flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, m["exponentialBuckets"])
	r.ExplicitBuckets = flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, m["explicitBuckets"])

	return r
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, f []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets expands an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, f *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.Width; !dcl.IsEmptyValueIndirect(v) {
		m["width"] = v
	}
	if v := f.Offset; !dcl.IsEmptyValueIndirect(v) {
		m["offset"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets flattens an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.Width = dcl.FlattenDouble(m["width"])
	r.Offset = dcl.FlattenDouble(m["offset"])

	return r
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, f []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets expands an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, f *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.GrowthFactor; !dcl.IsEmptyValueIndirect(v) {
		m["growthFactor"] = v
	}
	if v := f.Scale; !dcl.IsEmptyValueIndirect(v) {
		m["scale"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets flattens an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.GrowthFactor = dcl.FlattenDouble(m["growthFactor"])
	r.Scale = dcl.FlattenDouble(m["scale"])

	return r
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap(c *Client, f map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, f []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets expands an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, f *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Bounds; !dcl.IsEmptyValueIndirect(v) {
		m["bounds"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets flattens an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, i interface{}) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	}
	r.Bounds = dcl.FlattenFloatSlice(m["bounds"])

	return r
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingMap expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingMap(c *Client, f map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingSlice expands the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, f []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingMap flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingSlice flattens the contents of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling expands an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, f *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinimumValue; !dcl.IsEmptyValueIndirect(v) {
		m["minimumValue"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling flattens an instance of AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, i interface{}) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling
	}
	r.MinimumValue = dcl.FlattenDouble(m["minimumValue"])

	return r
}

// expandAlertPolicyConditionsConditionRateTimeWindowMap expands the contents of AlertPolicyConditionsConditionRateTimeWindow into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateTimeWindowMap(c *Client, f map[string]AlertPolicyConditionsConditionRateTimeWindow) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateTimeWindow(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateTimeWindowSlice expands the contents of AlertPolicyConditionsConditionRateTimeWindow into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateTimeWindowSlice(c *Client, f []AlertPolicyConditionsConditionRateTimeWindow) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateTimeWindow(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateTimeWindowMap flattens the contents of AlertPolicyConditionsConditionRateTimeWindow from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateTimeWindowMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateTimeWindow {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateTimeWindow{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateTimeWindow{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateTimeWindow)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateTimeWindow(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateTimeWindowSlice flattens the contents of AlertPolicyConditionsConditionRateTimeWindow from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateTimeWindowSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateTimeWindow {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateTimeWindow{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateTimeWindow{}
	}

	items := make([]AlertPolicyConditionsConditionRateTimeWindow, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateTimeWindow(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateTimeWindow expands an instance of AlertPolicyConditionsConditionRateTimeWindow into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateTimeWindow(c *Client, f *AlertPolicyConditionsConditionRateTimeWindow) (map[string]interface{}, error) {
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

// flattenAlertPolicyConditionsConditionRateTimeWindow flattens an instance of AlertPolicyConditionsConditionRateTimeWindow from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateTimeWindow(c *Client, i interface{}) *AlertPolicyConditionsConditionRateTimeWindow {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateTimeWindow{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateTimeWindow
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandAlertPolicyConditionsConditionRateTriggerMap expands the contents of AlertPolicyConditionsConditionRateTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateTriggerMap(c *Client, f map[string]AlertPolicyConditionsConditionRateTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionRateTriggerSlice expands the contents of AlertPolicyConditionsConditionRateTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateTriggerSlice(c *Client, f []AlertPolicyConditionsConditionRateTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionRateTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionRateTriggerMap flattens the contents of AlertPolicyConditionsConditionRateTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateTriggerMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionRateTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionRateTrigger{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionRateTrigger{}
	}

	items := make(map[string]AlertPolicyConditionsConditionRateTrigger)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionRateTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateTriggerSlice flattens the contents of AlertPolicyConditionsConditionRateTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateTriggerSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateTrigger{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateTrigger{}
	}

	items := make([]AlertPolicyConditionsConditionRateTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionRateTrigger expands an instance of AlertPolicyConditionsConditionRateTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionRateTrigger(c *Client, f *AlertPolicyConditionsConditionRateTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Count; !dcl.IsEmptyValueIndirect(v) {
		m["count"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionRateTrigger flattens an instance of AlertPolicyConditionsConditionRateTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateTrigger(c *Client, i interface{}) *AlertPolicyConditionsConditionRateTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionRateTrigger{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionRateTrigger
	}
	r.Count = dcl.FlattenInteger(m["count"])
	r.Percent = dcl.FlattenDouble(m["percent"])

	return r
}

// expandAlertPolicyConditionsConditionUpMonMap expands the contents of AlertPolicyConditionsConditionUpMon into a JSON
// request object.
func expandAlertPolicyConditionsConditionUpMonMap(c *Client, f map[string]AlertPolicyConditionsConditionUpMon) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionUpMon(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionUpMonSlice expands the contents of AlertPolicyConditionsConditionUpMon into a JSON
// request object.
func expandAlertPolicyConditionsConditionUpMonSlice(c *Client, f []AlertPolicyConditionsConditionUpMon) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionUpMon(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionUpMonMap flattens the contents of AlertPolicyConditionsConditionUpMon from a JSON
// response object.
func flattenAlertPolicyConditionsConditionUpMonMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionUpMon {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionUpMon{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionUpMon{}
	}

	items := make(map[string]AlertPolicyConditionsConditionUpMon)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionUpMon(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionUpMonSlice flattens the contents of AlertPolicyConditionsConditionUpMon from a JSON
// response object.
func flattenAlertPolicyConditionsConditionUpMonSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionUpMon {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionUpMon{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionUpMon{}
	}

	items := make([]AlertPolicyConditionsConditionUpMon, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionUpMon(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionUpMon expands an instance of AlertPolicyConditionsConditionUpMon into a JSON
// request object.
func expandAlertPolicyConditionsConditionUpMon(c *Client, f *AlertPolicyConditionsConditionUpMon) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v := f.EndpointId; !dcl.IsEmptyValueIndirect(v) {
		m["endpointId"] = v
	}
	if v := f.CheckId; !dcl.IsEmptyValueIndirect(v) {
		m["checkId"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionUpMonDuration(c, f.Duration); err != nil {
		return nil, fmt.Errorf("error expanding Duration into duration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["duration"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionUpMonTrigger(c, f.Trigger); err != nil {
		return nil, fmt.Errorf("error expanding Trigger into trigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["trigger"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionUpMon flattens an instance of AlertPolicyConditionsConditionUpMon from a JSON
// response object.
func flattenAlertPolicyConditionsConditionUpMon(c *Client, i interface{}) *AlertPolicyConditionsConditionUpMon {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionUpMon{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionUpMon
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.EndpointId = dcl.FlattenString(m["endpointId"])
	r.CheckId = dcl.FlattenString(m["checkId"])
	r.Duration = flattenAlertPolicyConditionsConditionUpMonDuration(c, m["duration"])
	r.Trigger = flattenAlertPolicyConditionsConditionUpMonTrigger(c, m["trigger"])

	return r
}

// expandAlertPolicyConditionsConditionUpMonDurationMap expands the contents of AlertPolicyConditionsConditionUpMonDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionUpMonDurationMap(c *Client, f map[string]AlertPolicyConditionsConditionUpMonDuration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionUpMonDuration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionUpMonDurationSlice expands the contents of AlertPolicyConditionsConditionUpMonDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionUpMonDurationSlice(c *Client, f []AlertPolicyConditionsConditionUpMonDuration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionUpMonDuration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionUpMonDurationMap flattens the contents of AlertPolicyConditionsConditionUpMonDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionUpMonDurationMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionUpMonDuration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionUpMonDuration{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionUpMonDuration{}
	}

	items := make(map[string]AlertPolicyConditionsConditionUpMonDuration)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionUpMonDuration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionUpMonDurationSlice flattens the contents of AlertPolicyConditionsConditionUpMonDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionUpMonDurationSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionUpMonDuration {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionUpMonDuration{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionUpMonDuration{}
	}

	items := make([]AlertPolicyConditionsConditionUpMonDuration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionUpMonDuration(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionUpMonDuration expands an instance of AlertPolicyConditionsConditionUpMonDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionUpMonDuration(c *Client, f *AlertPolicyConditionsConditionUpMonDuration) (map[string]interface{}, error) {
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

// flattenAlertPolicyConditionsConditionUpMonDuration flattens an instance of AlertPolicyConditionsConditionUpMonDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionUpMonDuration(c *Client, i interface{}) *AlertPolicyConditionsConditionUpMonDuration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionUpMonDuration{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionUpMonDuration
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandAlertPolicyConditionsConditionUpMonTriggerMap expands the contents of AlertPolicyConditionsConditionUpMonTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionUpMonTriggerMap(c *Client, f map[string]AlertPolicyConditionsConditionUpMonTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionUpMonTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionUpMonTriggerSlice expands the contents of AlertPolicyConditionsConditionUpMonTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionUpMonTriggerSlice(c *Client, f []AlertPolicyConditionsConditionUpMonTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionUpMonTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionUpMonTriggerMap flattens the contents of AlertPolicyConditionsConditionUpMonTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionUpMonTriggerMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionUpMonTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionUpMonTrigger{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionUpMonTrigger{}
	}

	items := make(map[string]AlertPolicyConditionsConditionUpMonTrigger)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionUpMonTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionUpMonTriggerSlice flattens the contents of AlertPolicyConditionsConditionUpMonTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionUpMonTriggerSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionUpMonTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionUpMonTrigger{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionUpMonTrigger{}
	}

	items := make([]AlertPolicyConditionsConditionUpMonTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionUpMonTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionUpMonTrigger expands an instance of AlertPolicyConditionsConditionUpMonTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionUpMonTrigger(c *Client, f *AlertPolicyConditionsConditionUpMonTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Count; !dcl.IsEmptyValueIndirect(v) {
		m["count"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionUpMonTrigger flattens an instance of AlertPolicyConditionsConditionUpMonTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionUpMonTrigger(c *Client, i interface{}) *AlertPolicyConditionsConditionUpMonTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionUpMonTrigger{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionUpMonTrigger
	}
	r.Count = dcl.FlattenInteger(m["count"])
	r.Percent = dcl.FlattenDouble(m["percent"])

	return r
}

// expandAlertPolicyConditionsConditionProcessCountMap expands the contents of AlertPolicyConditionsConditionProcessCount into a JSON
// request object.
func expandAlertPolicyConditionsConditionProcessCountMap(c *Client, f map[string]AlertPolicyConditionsConditionProcessCount) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionProcessCount(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionProcessCountSlice expands the contents of AlertPolicyConditionsConditionProcessCount into a JSON
// request object.
func expandAlertPolicyConditionsConditionProcessCountSlice(c *Client, f []AlertPolicyConditionsConditionProcessCount) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionProcessCount(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionProcessCountMap flattens the contents of AlertPolicyConditionsConditionProcessCount from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCountMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionProcessCount {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionProcessCount{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionProcessCount{}
	}

	items := make(map[string]AlertPolicyConditionsConditionProcessCount)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionProcessCount(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionProcessCountSlice flattens the contents of AlertPolicyConditionsConditionProcessCount from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCountSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionProcessCount {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionProcessCount{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionProcessCount{}
	}

	items := make([]AlertPolicyConditionsConditionProcessCount, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionProcessCount(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionProcessCount expands an instance of AlertPolicyConditionsConditionProcessCount into a JSON
// request object.
func expandAlertPolicyConditionsConditionProcessCount(c *Client, f *AlertPolicyConditionsConditionProcessCount) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Process; !dcl.IsEmptyValueIndirect(v) {
		m["process"] = v
	}
	if v := f.User; !dcl.IsEmptyValueIndirect(v) {
		m["user"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v := f.Comparison; !dcl.IsEmptyValueIndirect(v) {
		m["comparison"] = v
	}
	if v := f.ProcessCountThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["processCountThreshold"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionProcessCountTrigger(c, f.Trigger); err != nil {
		return nil, fmt.Errorf("error expanding Trigger into trigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["trigger"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionProcessCountDuration(c, f.Duration); err != nil {
		return nil, fmt.Errorf("error expanding Duration into duration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["duration"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionProcessCount flattens an instance of AlertPolicyConditionsConditionProcessCount from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCount(c *Client, i interface{}) *AlertPolicyConditionsConditionProcessCount {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionProcessCount{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionProcessCount
	}
	r.Process = dcl.FlattenString(m["process"])
	r.User = dcl.FlattenString(m["user"])
	r.Filter = dcl.FlattenString(m["filter"])
	r.Comparison = flattenAlertPolicyConditionsConditionProcessCountComparisonEnum(m["comparison"])
	r.ProcessCountThreshold = dcl.FlattenInteger(m["processCountThreshold"])
	r.Trigger = flattenAlertPolicyConditionsConditionProcessCountTrigger(c, m["trigger"])
	r.Duration = flattenAlertPolicyConditionsConditionProcessCountDuration(c, m["duration"])

	return r
}

// expandAlertPolicyConditionsConditionProcessCountTriggerMap expands the contents of AlertPolicyConditionsConditionProcessCountTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionProcessCountTriggerMap(c *Client, f map[string]AlertPolicyConditionsConditionProcessCountTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionProcessCountTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionProcessCountTriggerSlice expands the contents of AlertPolicyConditionsConditionProcessCountTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionProcessCountTriggerSlice(c *Client, f []AlertPolicyConditionsConditionProcessCountTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionProcessCountTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionProcessCountTriggerMap flattens the contents of AlertPolicyConditionsConditionProcessCountTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCountTriggerMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionProcessCountTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionProcessCountTrigger{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionProcessCountTrigger{}
	}

	items := make(map[string]AlertPolicyConditionsConditionProcessCountTrigger)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionProcessCountTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionProcessCountTriggerSlice flattens the contents of AlertPolicyConditionsConditionProcessCountTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCountTriggerSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionProcessCountTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionProcessCountTrigger{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionProcessCountTrigger{}
	}

	items := make([]AlertPolicyConditionsConditionProcessCountTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionProcessCountTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionProcessCountTrigger expands an instance of AlertPolicyConditionsConditionProcessCountTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionProcessCountTrigger(c *Client, f *AlertPolicyConditionsConditionProcessCountTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Count; !dcl.IsEmptyValueIndirect(v) {
		m["count"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionProcessCountTrigger flattens an instance of AlertPolicyConditionsConditionProcessCountTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCountTrigger(c *Client, i interface{}) *AlertPolicyConditionsConditionProcessCountTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionProcessCountTrigger{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionProcessCountTrigger
	}
	r.Count = dcl.FlattenInteger(m["count"])
	r.Percent = dcl.FlattenDouble(m["percent"])

	return r
}

// expandAlertPolicyConditionsConditionProcessCountDurationMap expands the contents of AlertPolicyConditionsConditionProcessCountDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionProcessCountDurationMap(c *Client, f map[string]AlertPolicyConditionsConditionProcessCountDuration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionProcessCountDuration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionProcessCountDurationSlice expands the contents of AlertPolicyConditionsConditionProcessCountDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionProcessCountDurationSlice(c *Client, f []AlertPolicyConditionsConditionProcessCountDuration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionProcessCountDuration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionProcessCountDurationMap flattens the contents of AlertPolicyConditionsConditionProcessCountDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCountDurationMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionProcessCountDuration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionProcessCountDuration{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionProcessCountDuration{}
	}

	items := make(map[string]AlertPolicyConditionsConditionProcessCountDuration)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionProcessCountDuration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionProcessCountDurationSlice flattens the contents of AlertPolicyConditionsConditionProcessCountDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCountDurationSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionProcessCountDuration {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionProcessCountDuration{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionProcessCountDuration{}
	}

	items := make([]AlertPolicyConditionsConditionProcessCountDuration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionProcessCountDuration(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionProcessCountDuration expands an instance of AlertPolicyConditionsConditionProcessCountDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionProcessCountDuration(c *Client, f *AlertPolicyConditionsConditionProcessCountDuration) (map[string]interface{}, error) {
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

// flattenAlertPolicyConditionsConditionProcessCountDuration flattens an instance of AlertPolicyConditionsConditionProcessCountDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCountDuration(c *Client, i interface{}) *AlertPolicyConditionsConditionProcessCountDuration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionProcessCountDuration{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionProcessCountDuration
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandAlertPolicyConditionsConditionTimeSeriesQueryLanguageMap expands the contents of AlertPolicyConditionsConditionTimeSeriesQueryLanguage into a JSON
// request object.
func expandAlertPolicyConditionsConditionTimeSeriesQueryLanguageMap(c *Client, f map[string]AlertPolicyConditionsConditionTimeSeriesQueryLanguage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionTimeSeriesQueryLanguageSlice expands the contents of AlertPolicyConditionsConditionTimeSeriesQueryLanguage into a JSON
// request object.
func expandAlertPolicyConditionsConditionTimeSeriesQueryLanguageSlice(c *Client, f []AlertPolicyConditionsConditionTimeSeriesQueryLanguage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionTimeSeriesQueryLanguageMap flattens the contents of AlertPolicyConditionsConditionTimeSeriesQueryLanguage from a JSON
// response object.
func flattenAlertPolicyConditionsConditionTimeSeriesQueryLanguageMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionTimeSeriesQueryLanguage{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionTimeSeriesQueryLanguage{}
	}

	items := make(map[string]AlertPolicyConditionsConditionTimeSeriesQueryLanguage)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionTimeSeriesQueryLanguageSlice flattens the contents of AlertPolicyConditionsConditionTimeSeriesQueryLanguage from a JSON
// response object.
func flattenAlertPolicyConditionsConditionTimeSeriesQueryLanguageSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionTimeSeriesQueryLanguage{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionTimeSeriesQueryLanguage{}
	}

	items := make([]AlertPolicyConditionsConditionTimeSeriesQueryLanguage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionTimeSeriesQueryLanguage expands an instance of AlertPolicyConditionsConditionTimeSeriesQueryLanguage into a JSON
// request object.
func expandAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c *Client, f *AlertPolicyConditionsConditionTimeSeriesQueryLanguage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Query; !dcl.IsEmptyValueIndirect(v) {
		m["query"] = v
	}
	if v := f.Summary; !dcl.IsEmptyValueIndirect(v) {
		m["summary"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionTimeSeriesQueryLanguage flattens an instance of AlertPolicyConditionsConditionTimeSeriesQueryLanguage from a JSON
// response object.
func flattenAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c *Client, i interface{}) *AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionTimeSeriesQueryLanguage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionTimeSeriesQueryLanguage
	}
	r.Query = dcl.FlattenString(m["query"])
	r.Summary = dcl.FlattenString(m["summary"])

	return r
}

// expandAlertPolicyConditionsConditionMonitoringQueryLanguageMap expands the contents of AlertPolicyConditionsConditionMonitoringQueryLanguage into a JSON
// request object.
func expandAlertPolicyConditionsConditionMonitoringQueryLanguageMap(c *Client, f map[string]AlertPolicyConditionsConditionMonitoringQueryLanguage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionMonitoringQueryLanguage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionMonitoringQueryLanguageSlice expands the contents of AlertPolicyConditionsConditionMonitoringQueryLanguage into a JSON
// request object.
func expandAlertPolicyConditionsConditionMonitoringQueryLanguageSlice(c *Client, f []AlertPolicyConditionsConditionMonitoringQueryLanguage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionMonitoringQueryLanguage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionMonitoringQueryLanguageMap flattens the contents of AlertPolicyConditionsConditionMonitoringQueryLanguage from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMonitoringQueryLanguageMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionMonitoringQueryLanguage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionMonitoringQueryLanguage{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionMonitoringQueryLanguage{}
	}

	items := make(map[string]AlertPolicyConditionsConditionMonitoringQueryLanguage)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionMonitoringQueryLanguage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionMonitoringQueryLanguageSlice flattens the contents of AlertPolicyConditionsConditionMonitoringQueryLanguage from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMonitoringQueryLanguageSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionMonitoringQueryLanguage {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionMonitoringQueryLanguage{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionMonitoringQueryLanguage{}
	}

	items := make([]AlertPolicyConditionsConditionMonitoringQueryLanguage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionMonitoringQueryLanguage(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionMonitoringQueryLanguage expands an instance of AlertPolicyConditionsConditionMonitoringQueryLanguage into a JSON
// request object.
func expandAlertPolicyConditionsConditionMonitoringQueryLanguage(c *Client, f *AlertPolicyConditionsConditionMonitoringQueryLanguage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Query; !dcl.IsEmptyValueIndirect(v) {
		m["query"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, f.Duration); err != nil {
		return nil, fmt.Errorf("error expanding Duration into duration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["duration"] = v
	}
	if v, err := expandAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, f.Trigger); err != nil {
		return nil, fmt.Errorf("error expanding Trigger into trigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["trigger"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionMonitoringQueryLanguage flattens an instance of AlertPolicyConditionsConditionMonitoringQueryLanguage from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMonitoringQueryLanguage(c *Client, i interface{}) *AlertPolicyConditionsConditionMonitoringQueryLanguage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionMonitoringQueryLanguage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionMonitoringQueryLanguage
	}
	r.Query = dcl.FlattenString(m["query"])
	r.Duration = flattenAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, m["duration"])
	r.Trigger = flattenAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, m["trigger"])

	return r
}

// expandAlertPolicyConditionsConditionMonitoringQueryLanguageDurationMap expands the contents of AlertPolicyConditionsConditionMonitoringQueryLanguageDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionMonitoringQueryLanguageDurationMap(c *Client, f map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionMonitoringQueryLanguageDurationSlice expands the contents of AlertPolicyConditionsConditionMonitoringQueryLanguageDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionMonitoringQueryLanguageDurationSlice(c *Client, f []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionMonitoringQueryLanguageDurationMap flattens the contents of AlertPolicyConditionsConditionMonitoringQueryLanguageDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMonitoringQueryLanguageDurationMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageDuration{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageDuration{}
	}

	items := make(map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageDuration)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionMonitoringQueryLanguageDurationSlice flattens the contents of AlertPolicyConditionsConditionMonitoringQueryLanguageDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMonitoringQueryLanguageDurationSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration{}
	}

	items := make([]AlertPolicyConditionsConditionMonitoringQueryLanguageDuration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionMonitoringQueryLanguageDuration expands an instance of AlertPolicyConditionsConditionMonitoringQueryLanguageDuration into a JSON
// request object.
func expandAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c *Client, f *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) (map[string]interface{}, error) {
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

// flattenAlertPolicyConditionsConditionMonitoringQueryLanguageDuration flattens an instance of AlertPolicyConditionsConditionMonitoringQueryLanguageDuration from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c *Client, i interface{}) *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionMonitoringQueryLanguageDuration{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionMonitoringQueryLanguageDuration
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerMap expands the contents of AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerMap(c *Client, f map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerSlice expands the contents of AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerSlice(c *Client, f []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerMap flattens the contents of AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerMap(c *Client, i interface{}) map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger{}
	}

	items := make(map[string]AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger)
	for k, item := range a {
		items[k] = *flattenAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerSlice flattens the contents of AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger{}
	}

	items := make([]AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger expands an instance of AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger into a JSON
// request object.
func expandAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c *Client, f *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Count; !dcl.IsEmptyValueIndirect(v) {
		m["count"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}

	return m, nil
}

// flattenAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger flattens an instance of AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger from a JSON
// response object.
func flattenAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c *Client, i interface{}) *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger
	}
	r.Count = dcl.FlattenInteger(m["count"])
	r.Percent = dcl.FlattenDouble(m["percent"])

	return r
}

// expandAlertPolicyEnabledMap expands the contents of AlertPolicyEnabled into a JSON
// request object.
func expandAlertPolicyEnabledMap(c *Client, f map[string]AlertPolicyEnabled) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyEnabled(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyEnabledSlice expands the contents of AlertPolicyEnabled into a JSON
// request object.
func expandAlertPolicyEnabledSlice(c *Client, f []AlertPolicyEnabled) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyEnabled(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyEnabledMap flattens the contents of AlertPolicyEnabled from a JSON
// response object.
func flattenAlertPolicyEnabledMap(c *Client, i interface{}) map[string]AlertPolicyEnabled {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyEnabled{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyEnabled{}
	}

	items := make(map[string]AlertPolicyEnabled)
	for k, item := range a {
		items[k] = *flattenAlertPolicyEnabled(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyEnabledSlice flattens the contents of AlertPolicyEnabled from a JSON
// response object.
func flattenAlertPolicyEnabledSlice(c *Client, i interface{}) []AlertPolicyEnabled {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyEnabled{}
	}

	if len(a) == 0 {
		return []AlertPolicyEnabled{}
	}

	items := make([]AlertPolicyEnabled, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyEnabled(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyEnabled expands an instance of AlertPolicyEnabled into a JSON
// request object.
func expandAlertPolicyEnabled(c *Client, f *AlertPolicyEnabled) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenAlertPolicyEnabled flattens an instance of AlertPolicyEnabled from a JSON
// response object.
func flattenAlertPolicyEnabled(c *Client, i interface{}) *AlertPolicyEnabled {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyEnabled{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyEnabled
	}
	r.Value = dcl.FlattenBool(m["value"])

	return r
}

// expandAlertPolicyValidityMap expands the contents of AlertPolicyValidity into a JSON
// request object.
func expandAlertPolicyValidityMap(c *Client, f map[string]AlertPolicyValidity) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyValidity(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyValiditySlice expands the contents of AlertPolicyValidity into a JSON
// request object.
func expandAlertPolicyValiditySlice(c *Client, f []AlertPolicyValidity) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyValidity(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyValidityMap flattens the contents of AlertPolicyValidity from a JSON
// response object.
func flattenAlertPolicyValidityMap(c *Client, i interface{}) map[string]AlertPolicyValidity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyValidity{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyValidity{}
	}

	items := make(map[string]AlertPolicyValidity)
	for k, item := range a {
		items[k] = *flattenAlertPolicyValidity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyValiditySlice flattens the contents of AlertPolicyValidity from a JSON
// response object.
func flattenAlertPolicyValiditySlice(c *Client, i interface{}) []AlertPolicyValidity {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyValidity{}
	}

	if len(a) == 0 {
		return []AlertPolicyValidity{}
	}

	items := make([]AlertPolicyValidity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyValidity(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyValidity expands an instance of AlertPolicyValidity into a JSON
// request object.
func expandAlertPolicyValidity(c *Client, f *AlertPolicyValidity) (map[string]interface{}, error) {
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
	if v, err := expandAlertPolicyValidityDetailsSlice(c, f.Details); err != nil {
		return nil, fmt.Errorf("error expanding Details into details: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["details"] = v
	}

	return m, nil
}

// flattenAlertPolicyValidity flattens an instance of AlertPolicyValidity from a JSON
// response object.
func flattenAlertPolicyValidity(c *Client, i interface{}) *AlertPolicyValidity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyValidity{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyValidity
	}
	r.Code = dcl.FlattenInteger(m["code"])
	r.Message = dcl.FlattenString(m["message"])
	r.Details = flattenAlertPolicyValidityDetailsSlice(c, m["details"])

	return r
}

// expandAlertPolicyValidityDetailsMap expands the contents of AlertPolicyValidityDetails into a JSON
// request object.
func expandAlertPolicyValidityDetailsMap(c *Client, f map[string]AlertPolicyValidityDetails) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyValidityDetails(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyValidityDetailsSlice expands the contents of AlertPolicyValidityDetails into a JSON
// request object.
func expandAlertPolicyValidityDetailsSlice(c *Client, f []AlertPolicyValidityDetails) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyValidityDetails(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyValidityDetailsMap flattens the contents of AlertPolicyValidityDetails from a JSON
// response object.
func flattenAlertPolicyValidityDetailsMap(c *Client, i interface{}) map[string]AlertPolicyValidityDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyValidityDetails{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyValidityDetails{}
	}

	items := make(map[string]AlertPolicyValidityDetails)
	for k, item := range a {
		items[k] = *flattenAlertPolicyValidityDetails(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyValidityDetailsSlice flattens the contents of AlertPolicyValidityDetails from a JSON
// response object.
func flattenAlertPolicyValidityDetailsSlice(c *Client, i interface{}) []AlertPolicyValidityDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyValidityDetails{}
	}

	if len(a) == 0 {
		return []AlertPolicyValidityDetails{}
	}

	items := make([]AlertPolicyValidityDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyValidityDetails(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyValidityDetails expands an instance of AlertPolicyValidityDetails into a JSON
// request object.
func expandAlertPolicyValidityDetails(c *Client, f *AlertPolicyValidityDetails) (map[string]interface{}, error) {
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

// flattenAlertPolicyValidityDetails flattens an instance of AlertPolicyValidityDetails from a JSON
// response object.
func flattenAlertPolicyValidityDetails(c *Client, i interface{}) *AlertPolicyValidityDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyValidityDetails{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyValidityDetails
	}
	r.TypeUrl = dcl.FlattenString(m["typeUrl"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandAlertPolicyCreationRecordMap expands the contents of AlertPolicyCreationRecord into a JSON
// request object.
func expandAlertPolicyCreationRecordMap(c *Client, f map[string]AlertPolicyCreationRecord) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyCreationRecord(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyCreationRecordSlice expands the contents of AlertPolicyCreationRecord into a JSON
// request object.
func expandAlertPolicyCreationRecordSlice(c *Client, f []AlertPolicyCreationRecord) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyCreationRecord(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyCreationRecordMap flattens the contents of AlertPolicyCreationRecord from a JSON
// response object.
func flattenAlertPolicyCreationRecordMap(c *Client, i interface{}) map[string]AlertPolicyCreationRecord {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyCreationRecord{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyCreationRecord{}
	}

	items := make(map[string]AlertPolicyCreationRecord)
	for k, item := range a {
		items[k] = *flattenAlertPolicyCreationRecord(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyCreationRecordSlice flattens the contents of AlertPolicyCreationRecord from a JSON
// response object.
func flattenAlertPolicyCreationRecordSlice(c *Client, i interface{}) []AlertPolicyCreationRecord {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyCreationRecord{}
	}

	if len(a) == 0 {
		return []AlertPolicyCreationRecord{}
	}

	items := make([]AlertPolicyCreationRecord, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyCreationRecord(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyCreationRecord expands an instance of AlertPolicyCreationRecord into a JSON
// request object.
func expandAlertPolicyCreationRecord(c *Client, f *AlertPolicyCreationRecord) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyCreationRecordMutateTime(c, f.MutateTime); err != nil {
		return nil, fmt.Errorf("error expanding MutateTime into mutateTime: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mutateTime"] = v
	}
	if v := f.MutatedBy; !dcl.IsEmptyValueIndirect(v) {
		m["mutatedBy"] = v
	}

	return m, nil
}

// flattenAlertPolicyCreationRecord flattens an instance of AlertPolicyCreationRecord from a JSON
// response object.
func flattenAlertPolicyCreationRecord(c *Client, i interface{}) *AlertPolicyCreationRecord {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyCreationRecord{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyCreationRecord
	}
	r.MutateTime = flattenAlertPolicyCreationRecordMutateTime(c, m["mutateTime"])
	r.MutatedBy = dcl.FlattenString(m["mutatedBy"])

	return r
}

// expandAlertPolicyCreationRecordMutateTimeMap expands the contents of AlertPolicyCreationRecordMutateTime into a JSON
// request object.
func expandAlertPolicyCreationRecordMutateTimeMap(c *Client, f map[string]AlertPolicyCreationRecordMutateTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyCreationRecordMutateTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyCreationRecordMutateTimeSlice expands the contents of AlertPolicyCreationRecordMutateTime into a JSON
// request object.
func expandAlertPolicyCreationRecordMutateTimeSlice(c *Client, f []AlertPolicyCreationRecordMutateTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyCreationRecordMutateTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyCreationRecordMutateTimeMap flattens the contents of AlertPolicyCreationRecordMutateTime from a JSON
// response object.
func flattenAlertPolicyCreationRecordMutateTimeMap(c *Client, i interface{}) map[string]AlertPolicyCreationRecordMutateTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyCreationRecordMutateTime{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyCreationRecordMutateTime{}
	}

	items := make(map[string]AlertPolicyCreationRecordMutateTime)
	for k, item := range a {
		items[k] = *flattenAlertPolicyCreationRecordMutateTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyCreationRecordMutateTimeSlice flattens the contents of AlertPolicyCreationRecordMutateTime from a JSON
// response object.
func flattenAlertPolicyCreationRecordMutateTimeSlice(c *Client, i interface{}) []AlertPolicyCreationRecordMutateTime {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyCreationRecordMutateTime{}
	}

	if len(a) == 0 {
		return []AlertPolicyCreationRecordMutateTime{}
	}

	items := make([]AlertPolicyCreationRecordMutateTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyCreationRecordMutateTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyCreationRecordMutateTime expands an instance of AlertPolicyCreationRecordMutateTime into a JSON
// request object.
func expandAlertPolicyCreationRecordMutateTime(c *Client, f *AlertPolicyCreationRecordMutateTime) (map[string]interface{}, error) {
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

// flattenAlertPolicyCreationRecordMutateTime flattens an instance of AlertPolicyCreationRecordMutateTime from a JSON
// response object.
func flattenAlertPolicyCreationRecordMutateTime(c *Client, i interface{}) *AlertPolicyCreationRecordMutateTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyCreationRecordMutateTime{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyCreationRecordMutateTime
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandAlertPolicyMutationRecordMap expands the contents of AlertPolicyMutationRecord into a JSON
// request object.
func expandAlertPolicyMutationRecordMap(c *Client, f map[string]AlertPolicyMutationRecord) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyMutationRecord(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyMutationRecordSlice expands the contents of AlertPolicyMutationRecord into a JSON
// request object.
func expandAlertPolicyMutationRecordSlice(c *Client, f []AlertPolicyMutationRecord) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyMutationRecord(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyMutationRecordMap flattens the contents of AlertPolicyMutationRecord from a JSON
// response object.
func flattenAlertPolicyMutationRecordMap(c *Client, i interface{}) map[string]AlertPolicyMutationRecord {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyMutationRecord{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyMutationRecord{}
	}

	items := make(map[string]AlertPolicyMutationRecord)
	for k, item := range a {
		items[k] = *flattenAlertPolicyMutationRecord(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyMutationRecordSlice flattens the contents of AlertPolicyMutationRecord from a JSON
// response object.
func flattenAlertPolicyMutationRecordSlice(c *Client, i interface{}) []AlertPolicyMutationRecord {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyMutationRecord{}
	}

	if len(a) == 0 {
		return []AlertPolicyMutationRecord{}
	}

	items := make([]AlertPolicyMutationRecord, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyMutationRecord(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyMutationRecord expands an instance of AlertPolicyMutationRecord into a JSON
// request object.
func expandAlertPolicyMutationRecord(c *Client, f *AlertPolicyMutationRecord) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandAlertPolicyMutationRecordMutateTime(c, f.MutateTime); err != nil {
		return nil, fmt.Errorf("error expanding MutateTime into mutateTime: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mutateTime"] = v
	}
	if v := f.MutatedBy; !dcl.IsEmptyValueIndirect(v) {
		m["mutatedBy"] = v
	}

	return m, nil
}

// flattenAlertPolicyMutationRecord flattens an instance of AlertPolicyMutationRecord from a JSON
// response object.
func flattenAlertPolicyMutationRecord(c *Client, i interface{}) *AlertPolicyMutationRecord {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyMutationRecord{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyMutationRecord
	}
	r.MutateTime = flattenAlertPolicyMutationRecordMutateTime(c, m["mutateTime"])
	r.MutatedBy = dcl.FlattenString(m["mutatedBy"])

	return r
}

// expandAlertPolicyMutationRecordMutateTimeMap expands the contents of AlertPolicyMutationRecordMutateTime into a JSON
// request object.
func expandAlertPolicyMutationRecordMutateTimeMap(c *Client, f map[string]AlertPolicyMutationRecordMutateTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyMutationRecordMutateTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyMutationRecordMutateTimeSlice expands the contents of AlertPolicyMutationRecordMutateTime into a JSON
// request object.
func expandAlertPolicyMutationRecordMutateTimeSlice(c *Client, f []AlertPolicyMutationRecordMutateTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyMutationRecordMutateTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyMutationRecordMutateTimeMap flattens the contents of AlertPolicyMutationRecordMutateTime from a JSON
// response object.
func flattenAlertPolicyMutationRecordMutateTimeMap(c *Client, i interface{}) map[string]AlertPolicyMutationRecordMutateTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyMutationRecordMutateTime{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyMutationRecordMutateTime{}
	}

	items := make(map[string]AlertPolicyMutationRecordMutateTime)
	for k, item := range a {
		items[k] = *flattenAlertPolicyMutationRecordMutateTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyMutationRecordMutateTimeSlice flattens the contents of AlertPolicyMutationRecordMutateTime from a JSON
// response object.
func flattenAlertPolicyMutationRecordMutateTimeSlice(c *Client, i interface{}) []AlertPolicyMutationRecordMutateTime {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyMutationRecordMutateTime{}
	}

	if len(a) == 0 {
		return []AlertPolicyMutationRecordMutateTime{}
	}

	items := make([]AlertPolicyMutationRecordMutateTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyMutationRecordMutateTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyMutationRecordMutateTime expands an instance of AlertPolicyMutationRecordMutateTime into a JSON
// request object.
func expandAlertPolicyMutationRecordMutateTime(c *Client, f *AlertPolicyMutationRecordMutateTime) (map[string]interface{}, error) {
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

// flattenAlertPolicyMutationRecordMutateTime flattens an instance of AlertPolicyMutationRecordMutateTime from a JSON
// response object.
func flattenAlertPolicyMutationRecordMutateTime(c *Client, i interface{}) *AlertPolicyMutationRecordMutateTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyMutationRecordMutateTime{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyMutationRecordMutateTime
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandAlertPolicyIncidentStrategyMap expands the contents of AlertPolicyIncidentStrategy into a JSON
// request object.
func expandAlertPolicyIncidentStrategyMap(c *Client, f map[string]AlertPolicyIncidentStrategy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyIncidentStrategy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyIncidentStrategySlice expands the contents of AlertPolicyIncidentStrategy into a JSON
// request object.
func expandAlertPolicyIncidentStrategySlice(c *Client, f []AlertPolicyIncidentStrategy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyIncidentStrategy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyIncidentStrategyMap flattens the contents of AlertPolicyIncidentStrategy from a JSON
// response object.
func flattenAlertPolicyIncidentStrategyMap(c *Client, i interface{}) map[string]AlertPolicyIncidentStrategy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyIncidentStrategy{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyIncidentStrategy{}
	}

	items := make(map[string]AlertPolicyIncidentStrategy)
	for k, item := range a {
		items[k] = *flattenAlertPolicyIncidentStrategy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyIncidentStrategySlice flattens the contents of AlertPolicyIncidentStrategy from a JSON
// response object.
func flattenAlertPolicyIncidentStrategySlice(c *Client, i interface{}) []AlertPolicyIncidentStrategy {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyIncidentStrategy{}
	}

	if len(a) == 0 {
		return []AlertPolicyIncidentStrategy{}
	}

	items := make([]AlertPolicyIncidentStrategy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyIncidentStrategy(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyIncidentStrategy expands an instance of AlertPolicyIncidentStrategy into a JSON
// request object.
func expandAlertPolicyIncidentStrategy(c *Client, f *AlertPolicyIncidentStrategy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenAlertPolicyIncidentStrategy flattens an instance of AlertPolicyIncidentStrategy from a JSON
// response object.
func flattenAlertPolicyIncidentStrategy(c *Client, i interface{}) *AlertPolicyIncidentStrategy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyIncidentStrategy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyIncidentStrategy
	}
	r.Type = flattenAlertPolicyIncidentStrategyTypeEnum(m["type"])

	return r
}

// expandAlertPolicyMetadataMap expands the contents of AlertPolicyMetadata into a JSON
// request object.
func expandAlertPolicyMetadataMap(c *Client, f map[string]AlertPolicyMetadata) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAlertPolicyMetadata(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAlertPolicyMetadataSlice expands the contents of AlertPolicyMetadata into a JSON
// request object.
func expandAlertPolicyMetadataSlice(c *Client, f []AlertPolicyMetadata) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAlertPolicyMetadata(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAlertPolicyMetadataMap flattens the contents of AlertPolicyMetadata from a JSON
// response object.
func flattenAlertPolicyMetadataMap(c *Client, i interface{}) map[string]AlertPolicyMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AlertPolicyMetadata{}
	}

	if len(a) == 0 {
		return map[string]AlertPolicyMetadata{}
	}

	items := make(map[string]AlertPolicyMetadata)
	for k, item := range a {
		items[k] = *flattenAlertPolicyMetadata(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAlertPolicyMetadataSlice flattens the contents of AlertPolicyMetadata from a JSON
// response object.
func flattenAlertPolicyMetadataSlice(c *Client, i interface{}) []AlertPolicyMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyMetadata{}
	}

	if len(a) == 0 {
		return []AlertPolicyMetadata{}
	}

	items := make([]AlertPolicyMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyMetadata(c, item.(map[string]interface{})))
	}

	return items
}

// expandAlertPolicyMetadata expands an instance of AlertPolicyMetadata into a JSON
// request object.
func expandAlertPolicyMetadata(c *Client, f *AlertPolicyMetadata) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SloNames; !dcl.IsEmptyValueIndirect(v) {
		m["sloNames"] = v
	}

	return m, nil
}

// flattenAlertPolicyMetadata flattens an instance of AlertPolicyMetadata from a JSON
// response object.
func flattenAlertPolicyMetadata(c *Client, i interface{}) *AlertPolicyMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AlertPolicyMetadata{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAlertPolicyMetadata
	}
	r.SloNames = dcl.FlattenStringSlice(m["sloNames"])

	return r
}

// flattenAlertPolicyConditionsResourceStateFilterEnumSlice flattens the contents of AlertPolicyConditionsResourceStateFilterEnum from a JSON
// response object.
func flattenAlertPolicyConditionsResourceStateFilterEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsResourceStateFilterEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsResourceStateFilterEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsResourceStateFilterEnum{}
	}

	items := make([]AlertPolicyConditionsResourceStateFilterEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsResourceStateFilterEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsResourceStateFilterEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsResourceStateFilterEnum with the same value as that string.
func flattenAlertPolicyConditionsResourceStateFilterEnum(i interface{}) *AlertPolicyConditionsResourceStateFilterEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsResourceStateFilterEnumRef("")
	}

	return AlertPolicyConditionsResourceStateFilterEnumRef(s)
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(i interface{}) *AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumRef("")
	}

	return AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumRef(s)
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumSlice flattens the contents of AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(i interface{}) *AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumRef("")
	}

	return AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumRef(s)
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumRef("")
	}

	return AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumRef(s)
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumSlice flattens the contents of AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(i interface{}) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumRef("")
	}

	return AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumRef(s)
}

// flattenAlertPolicyConditionsConditionThresholdComparisonEnumSlice flattens the contents of AlertPolicyConditionsConditionThresholdComparisonEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionThresholdComparisonEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionThresholdComparisonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionThresholdComparisonEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionThresholdComparisonEnum{}
	}

	items := make([]AlertPolicyConditionsConditionThresholdComparisonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdComparisonEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionThresholdComparisonEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionThresholdComparisonEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionThresholdComparisonEnum(i interface{}) *AlertPolicyConditionsConditionThresholdComparisonEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionThresholdComparisonEnumRef("")
	}

	return AlertPolicyConditionsConditionThresholdComparisonEnumRef(s)
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(i interface{}) *AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumRef("")
	}

	return AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumRef(s)
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumSlice flattens the contents of AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum{}
	}

	items := make([]AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(i interface{}) *AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumRef("")
	}

	return AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumRef(s)
}

// flattenAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumSlice flattens the contents of AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(i interface{}) *AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumRef("")
	}

	return AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumRef(s)
}

// flattenAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumSlice flattens the contents of AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum{}
	}

	items := make([]AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(i interface{}) *AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumRef("")
	}

	return AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumRef(s)
}

// flattenAlertPolicyConditionsConditionRateComparisonEnumSlice flattens the contents of AlertPolicyConditionsConditionRateComparisonEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionRateComparisonEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionRateComparisonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionRateComparisonEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionRateComparisonEnum{}
	}

	items := make([]AlertPolicyConditionsConditionRateComparisonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionRateComparisonEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionRateComparisonEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionRateComparisonEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionRateComparisonEnum(i interface{}) *AlertPolicyConditionsConditionRateComparisonEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionRateComparisonEnumRef("")
	}

	return AlertPolicyConditionsConditionRateComparisonEnumRef(s)
}

// flattenAlertPolicyConditionsConditionProcessCountComparisonEnumSlice flattens the contents of AlertPolicyConditionsConditionProcessCountComparisonEnum from a JSON
// response object.
func flattenAlertPolicyConditionsConditionProcessCountComparisonEnumSlice(c *Client, i interface{}) []AlertPolicyConditionsConditionProcessCountComparisonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyConditionsConditionProcessCountComparisonEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyConditionsConditionProcessCountComparisonEnum{}
	}

	items := make([]AlertPolicyConditionsConditionProcessCountComparisonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyConditionsConditionProcessCountComparisonEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyConditionsConditionProcessCountComparisonEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyConditionsConditionProcessCountComparisonEnum with the same value as that string.
func flattenAlertPolicyConditionsConditionProcessCountComparisonEnum(i interface{}) *AlertPolicyConditionsConditionProcessCountComparisonEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyConditionsConditionProcessCountComparisonEnumRef("")
	}

	return AlertPolicyConditionsConditionProcessCountComparisonEnumRef(s)
}

// flattenAlertPolicyCombinerEnumSlice flattens the contents of AlertPolicyCombinerEnum from a JSON
// response object.
func flattenAlertPolicyCombinerEnumSlice(c *Client, i interface{}) []AlertPolicyCombinerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyCombinerEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyCombinerEnum{}
	}

	items := make([]AlertPolicyCombinerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyCombinerEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyCombinerEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyCombinerEnum with the same value as that string.
func flattenAlertPolicyCombinerEnum(i interface{}) *AlertPolicyCombinerEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyCombinerEnumRef("")
	}

	return AlertPolicyCombinerEnumRef(s)
}

// flattenAlertPolicyIncidentStrategyTypeEnumSlice flattens the contents of AlertPolicyIncidentStrategyTypeEnum from a JSON
// response object.
func flattenAlertPolicyIncidentStrategyTypeEnumSlice(c *Client, i interface{}) []AlertPolicyIncidentStrategyTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AlertPolicyIncidentStrategyTypeEnum{}
	}

	if len(a) == 0 {
		return []AlertPolicyIncidentStrategyTypeEnum{}
	}

	items := make([]AlertPolicyIncidentStrategyTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAlertPolicyIncidentStrategyTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAlertPolicyIncidentStrategyTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AlertPolicyIncidentStrategyTypeEnum with the same value as that string.
func flattenAlertPolicyIncidentStrategyTypeEnum(i interface{}) *AlertPolicyIncidentStrategyTypeEnum {
	s, ok := i.(string)
	if !ok {
		return AlertPolicyIncidentStrategyTypeEnumRef("")
	}

	return AlertPolicyIncidentStrategyTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AlertPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAlertPolicy(b, c)
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

type alertPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         alertPolicyApiOperation
}

func convertFieldDiffToAlertPolicyOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]alertPolicyDiff, error) {
	var diffs []alertPolicyDiff
	for _, op := range ops {
		diff := alertPolicyDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToalertPolicyApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToalertPolicyApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (alertPolicyApiOperation, error) {
	switch op {

	case "updateAlertPolicyUpdateAlertPolicyOperation":
		return &updateAlertPolicyUpdateAlertPolicyOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
