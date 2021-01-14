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
	"github.com/mohae/deepcopy"
	"io/ioutil"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
	"reflect"
	"strings"
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
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.MonitoringOperation
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listAlertPolicyOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
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
	for _, v := range m.Items {
		res := flattenAlertPolicy(c, v)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.MonitoringOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://monitoring.googleapis.com/", "GET"); err != nil {
		return err
	}
	_, err = c.GetAlertPolicy(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createAlertPolicyOperation struct{}

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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.MonitoringOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://monitoring.googleapis.com/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	r.Name, err = o.FetchName()
	if err != nil {
		return fmt.Errorf("error trying to retrieve Name: %w", err)
	}

	if _, err := c.GetAlertPolicy(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getAlertPolicyRaw(ctx context.Context, r *AlertPolicy) ([]byte, error) {

	u, err := alertPolicyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) alertPolicyDiffsForRawDesired(ctx context.Context, rawDesired *AlertPolicy, opts ...dcl.ApplyOption) (initial, desired *AlertPolicy, diffs []alertPolicyDiff, err error) {
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AlertPolicy); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected AlertPolicy, got %T", sh)
		} else {
			_ = r
		}
	}

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
	if dcl.IsZeroValue(rawDesired.DisplayName) {
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
	if dcl.IsZeroValue(rawDesired.Disabled) {
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.Combiner) && dcl.IsEmptyValueIndirect(rawDesired.Combiner) {
		rawNew.Combiner = rawDesired.Combiner
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
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

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.NameToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	return rawNew, nil
}

func canonicalizeAlertPolicyDocumentation(des, initial *AlertPolicyDocumentation, opts ...dcl.ApplyOption) *AlertPolicyDocumentation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Content) {
		des.Content = initial.Content
	}
	if dcl.IsZeroValue(des.MimeType) {
		des.MimeType = initial.MimeType
	}

	return des
}

func canonicalizeNewAlertPolicyDocumentation(c *Client, des, nw *AlertPolicyDocumentation) *AlertPolicyDocumentation {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyDocumentation(c, &d, &n) {
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

func canonicalizeAlertPolicyConditions(des, initial *AlertPolicyConditions, opts ...dcl.ApplyOption) *AlertPolicyConditions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.DisplayName) {
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
			if !compareAlertPolicyConditions(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThreshold(des, initial *AlertPolicyConditionsConditionThreshold, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThreshold {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}
	if dcl.IsZeroValue(des.Aggregations) {
		des.Aggregations = initial.Aggregations
	}
	if dcl.IsZeroValue(des.DenominatorFilter) {
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
	if dcl.IsZeroValue(des.Duration) {
		des.Duration = initial.Duration
	}
	des.Trigger = canonicalizeAlertPolicyConditionsConditionThresholdTrigger(des.Trigger, initial.Trigger, opts...)

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionThreshold(c *Client, des, nw *AlertPolicyConditionsConditionThreshold) *AlertPolicyConditionsConditionThreshold {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyConditionsConditionThreshold(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdAggregations(des, initial *AlertPolicyConditionsConditionThresholdAggregations, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AlignmentPeriod) {
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
			if !compareAlertPolicyConditionsConditionThresholdAggregations(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(des, initial *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregations(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregations, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AlignmentPeriod) {
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
			if !compareAlertPolicyConditionsConditionThresholdDenominatorAggregations(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(des, initial *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionThresholdTrigger(des, initial *AlertPolicyConditionsConditionThresholdTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionThresholdTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionThresholdTrigger(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsent(des, initial *AlertPolicyConditionsConditionAbsent, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsent {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Filter) {
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
			if !compareAlertPolicyConditionsConditionAbsent(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentAggregations(des, initial *AlertPolicyConditionsConditionAbsentAggregations, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AlignmentPeriod) {
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
			if !compareAlertPolicyConditionsConditionAbsentAggregations(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(des, initial *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentDuration(des, initial *AlertPolicyConditionsConditionAbsentDuration, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
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

func canonicalizeNewAlertPolicyConditionsConditionAbsentDuration(c *Client, des, nw *AlertPolicyConditionsConditionAbsentDuration) *AlertPolicyConditionsConditionAbsentDuration {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyConditionsConditionAbsentDuration(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionAbsentTrigger(des, initial *AlertPolicyConditionsConditionAbsentTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionAbsentTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionAbsentTrigger(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionMatchedLog(des, initial *AlertPolicyConditionsConditionMatchedLog, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionMatchedLog {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Filter) {
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
			if !compareAlertPolicyConditionsConditionMatchedLog(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionClusterOutlier(des, initial *AlertPolicyConditionsConditionClusterOutlier, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionClusterOutlier {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionClusterOutlier(c *Client, des, nw *AlertPolicyConditionsConditionClusterOutlier) *AlertPolicyConditionsConditionClusterOutlier {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyConditionsConditionClusterOutlier(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRate(des, initial *AlertPolicyConditionsConditionRate, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Filter) {
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
			if !compareAlertPolicyConditionsConditionRate(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateAggregations(des, initial *AlertPolicyConditionsConditionRateAggregations, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AlignmentPeriod) {
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
			if !compareAlertPolicyConditionsConditionRateAggregations(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(des, initial *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateTimeWindow(des, initial *AlertPolicyConditionsConditionRateTimeWindow, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateTimeWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
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

func canonicalizeNewAlertPolicyConditionsConditionRateTimeWindow(c *Client, des, nw *AlertPolicyConditionsConditionRateTimeWindow) *AlertPolicyConditionsConditionRateTimeWindow {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyConditionsConditionRateTimeWindow(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionRateTrigger(des, initial *AlertPolicyConditionsConditionRateTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionRateTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionRateTrigger(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionUpMon(des, initial *AlertPolicyConditionsConditionUpMon, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionUpMon {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}
	if dcl.IsZeroValue(des.EndpointId) {
		des.EndpointId = initial.EndpointId
	}
	if dcl.IsZeroValue(des.CheckId) {
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
			if !compareAlertPolicyConditionsConditionUpMon(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionUpMonDuration(des, initial *AlertPolicyConditionsConditionUpMonDuration, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionUpMonDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
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

func canonicalizeNewAlertPolicyConditionsConditionUpMonDuration(c *Client, des, nw *AlertPolicyConditionsConditionUpMonDuration) *AlertPolicyConditionsConditionUpMonDuration {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyConditionsConditionUpMonDuration(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionUpMonTrigger(des, initial *AlertPolicyConditionsConditionUpMonTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionUpMonTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionUpMonTrigger(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionProcessCount(des, initial *AlertPolicyConditionsConditionProcessCount, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionProcessCount {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Process) {
		des.Process = initial.Process
	}
	if dcl.IsZeroValue(des.User) {
		des.User = initial.User
	}
	if dcl.IsZeroValue(des.Filter) {
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
			if !compareAlertPolicyConditionsConditionProcessCount(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionProcessCountTrigger(des, initial *AlertPolicyConditionsConditionProcessCountTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionProcessCountTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionProcessCountTrigger(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionProcessCountDuration(des, initial *AlertPolicyConditionsConditionProcessCountDuration, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionProcessCountDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
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

func canonicalizeNewAlertPolicyConditionsConditionProcessCountDuration(c *Client, des, nw *AlertPolicyConditionsConditionProcessCountDuration) *AlertPolicyConditionsConditionProcessCountDuration {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyConditionsConditionProcessCountDuration(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionTimeSeriesQueryLanguage(des, initial *AlertPolicyConditionsConditionTimeSeriesQueryLanguage, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Query) {
		des.Query = initial.Query
	}
	if dcl.IsZeroValue(des.Summary) {
		des.Summary = initial.Summary
	}

	return des
}

func canonicalizeNewAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c *Client, des, nw *AlertPolicyConditionsConditionTimeSeriesQueryLanguage) *AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionMonitoringQueryLanguage(des, initial *AlertPolicyConditionsConditionMonitoringQueryLanguage, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionMonitoringQueryLanguage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Query) {
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
			if !compareAlertPolicyConditionsConditionMonitoringQueryLanguage(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(des, initial *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
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

func canonicalizeNewAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c *Client, des, nw *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, &d, &n) {
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

func canonicalizeAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(des, initial *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger, opts ...dcl.ApplyOption) *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, &d, &n) {
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

func canonicalizeAlertPolicyEnabled(des, initial *AlertPolicyEnabled, opts ...dcl.ApplyOption) *AlertPolicyEnabled {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewAlertPolicyEnabled(c *Client, des, nw *AlertPolicyEnabled) *AlertPolicyEnabled {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyEnabled(c, &d, &n) {
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

func canonicalizeAlertPolicyValidity(des, initial *AlertPolicyValidity, opts ...dcl.ApplyOption) *AlertPolicyValidity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
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

func canonicalizeNewAlertPolicyValidity(c *Client, des, nw *AlertPolicyValidity) *AlertPolicyValidity {
	if des == nil || nw == nil {
		return nw
	}

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
			if !compareAlertPolicyValidity(c, &d, &n) {
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

func canonicalizeAlertPolicyValidityDetails(des, initial *AlertPolicyValidityDetails, opts ...dcl.ApplyOption) *AlertPolicyValidityDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
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

func canonicalizeNewAlertPolicyValidityDetails(c *Client, des, nw *AlertPolicyValidityDetails) *AlertPolicyValidityDetails {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyValidityDetails(c, &d, &n) {
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

func canonicalizeAlertPolicyCreationRecord(des, initial *AlertPolicyCreationRecord, opts ...dcl.ApplyOption) *AlertPolicyCreationRecord {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.MutateTime = canonicalizeAlertPolicyCreationRecordMutateTime(des.MutateTime, initial.MutateTime, opts...)
	if dcl.IsZeroValue(des.MutatedBy) {
		des.MutatedBy = initial.MutatedBy
	}

	return des
}

func canonicalizeNewAlertPolicyCreationRecord(c *Client, des, nw *AlertPolicyCreationRecord) *AlertPolicyCreationRecord {
	if des == nil || nw == nil {
		return nw
	}

	nw.MutateTime = canonicalizeNewAlertPolicyCreationRecordMutateTime(c, des.MutateTime, nw.MutateTime)

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
			if !compareAlertPolicyCreationRecord(c, &d, &n) {
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

func canonicalizeAlertPolicyCreationRecordMutateTime(des, initial *AlertPolicyCreationRecordMutateTime, opts ...dcl.ApplyOption) *AlertPolicyCreationRecordMutateTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
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

func canonicalizeNewAlertPolicyCreationRecordMutateTime(c *Client, des, nw *AlertPolicyCreationRecordMutateTime) *AlertPolicyCreationRecordMutateTime {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyCreationRecordMutateTime(c, &d, &n) {
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

func canonicalizeAlertPolicyMutationRecord(des, initial *AlertPolicyMutationRecord, opts ...dcl.ApplyOption) *AlertPolicyMutationRecord {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.MutateTime = canonicalizeAlertPolicyMutationRecordMutateTime(des.MutateTime, initial.MutateTime, opts...)
	if dcl.IsZeroValue(des.MutatedBy) {
		des.MutatedBy = initial.MutatedBy
	}

	return des
}

func canonicalizeNewAlertPolicyMutationRecord(c *Client, des, nw *AlertPolicyMutationRecord) *AlertPolicyMutationRecord {
	if des == nil || nw == nil {
		return nw
	}

	nw.MutateTime = canonicalizeNewAlertPolicyMutationRecordMutateTime(c, des.MutateTime, nw.MutateTime)

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
			if !compareAlertPolicyMutationRecord(c, &d, &n) {
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

func canonicalizeAlertPolicyMutationRecordMutateTime(des, initial *AlertPolicyMutationRecordMutateTime, opts ...dcl.ApplyOption) *AlertPolicyMutationRecordMutateTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
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

func canonicalizeNewAlertPolicyMutationRecordMutateTime(c *Client, des, nw *AlertPolicyMutationRecordMutateTime) *AlertPolicyMutationRecordMutateTime {
	if des == nil || nw == nil {
		return nw
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
			if !compareAlertPolicyMutationRecordMutateTime(c, &d, &n) {
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

func canonicalizeAlertPolicyIncidentStrategy(des, initial *AlertPolicyIncidentStrategy, opts ...dcl.ApplyOption) *AlertPolicyIncidentStrategy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyIncidentStrategy(c, &d, &n) {
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

func canonicalizeAlertPolicyMetadata(des, initial *AlertPolicyMetadata, opts ...dcl.ApplyOption) *AlertPolicyMetadata {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*AlertPolicy)
		_ = r
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
			if !compareAlertPolicyMetadata(c, &d, &n) {
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

type alertPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         alertPolicyApiOperation
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
func diffAlertPolicy(c *Client, desired, actual *AlertPolicy, opts ...dcl.ApplyOption) ([]alertPolicyDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []alertPolicyDiff
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %#v\nACTUAL: %#v", desired.Name, actual.Name)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.DisplayName) && (dcl.IsZeroValue(actual.DisplayName) || !reflect.DeepEqual(*desired.DisplayName, *actual.DisplayName)) {
		c.Config.Logger.Infof("Detected diff in DisplayName.\nDESIRED: %#v\nACTUAL: %#v", desired.DisplayName, actual.DisplayName)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "DisplayName",
		})
	}
	if compareAlertPolicyDocumentation(c, desired.Documentation, actual.Documentation) {
		c.Config.Logger.Infof("Detected diff in Documentation.\nDESIRED: %#v\nACTUAL: %#v", desired.Documentation, actual.Documentation)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Documentation",
		})
	}
	if !reflect.DeepEqual(desired.UserLabels, actual.UserLabels) {
		c.Config.Logger.Infof("Detected diff in UserLabels.\nDESIRED: %#v\nACTUAL: %#v", desired.UserLabels, actual.UserLabels)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "UserLabels",
		})
	}
	if compareAlertPolicyConditionsSlice(c, desired.Conditions, actual.Conditions) {
		c.Config.Logger.Infof("Detected diff in Conditions.\nDESIRED: %#v\nACTUAL: %#v", desired.Conditions, actual.Conditions)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Conditions",
		})
	}
	if !dcl.IsZeroValue(desired.Combiner) && (dcl.IsZeroValue(actual.Combiner) || !reflect.DeepEqual(*desired.Combiner, *actual.Combiner)) {
		c.Config.Logger.Infof("Detected diff in Combiner.\nDESIRED: %#v\nACTUAL: %#v", desired.Combiner, actual.Combiner)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Combiner",
		})
	}
	if !dcl.IsZeroValue(desired.Disabled) && (dcl.IsZeroValue(actual.Disabled) || !reflect.DeepEqual(*desired.Disabled, *actual.Disabled)) {
		c.Config.Logger.Infof("Detected diff in Disabled.\nDESIRED: %#v\nACTUAL: %#v", desired.Disabled, actual.Disabled)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Disabled",
		})
	}
	if compareAlertPolicyEnabled(c, desired.Enabled, actual.Enabled) {
		c.Config.Logger.Infof("Detected diff in Enabled.\nDESIRED: %#v\nACTUAL: %#v", desired.Enabled, actual.Enabled)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Enabled",
		})
	}
	if compareAlertPolicyValidity(c, desired.Validity, actual.Validity) {
		c.Config.Logger.Infof("Detected diff in Validity.\nDESIRED: %#v\nACTUAL: %#v", desired.Validity, actual.Validity)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Validity",
		})
	}
	if !dcl.SliceEquals(desired.NotificationChannels, actual.NotificationChannels) {
		c.Config.Logger.Infof("Detected diff in NotificationChannels.\nDESIRED: %#v\nACTUAL: %#v", desired.NotificationChannels, actual.NotificationChannels)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "NotificationChannels",
		})
	}
	if compareAlertPolicyCreationRecord(c, desired.CreationRecord, actual.CreationRecord) {
		c.Config.Logger.Infof("Detected diff in CreationRecord.\nDESIRED: %#v\nACTUAL: %#v", desired.CreationRecord, actual.CreationRecord)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "CreationRecord",
		})
	}
	if compareAlertPolicyMutationRecord(c, desired.MutationRecord, actual.MutationRecord) {
		c.Config.Logger.Infof("Detected diff in MutationRecord.\nDESIRED: %#v\nACTUAL: %#v", desired.MutationRecord, actual.MutationRecord)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "MutationRecord",
		})
	}
	if compareAlertPolicyIncidentStrategy(c, desired.IncidentStrategy, actual.IncidentStrategy) {
		c.Config.Logger.Infof("Detected diff in IncidentStrategy.\nDESIRED: %#v\nACTUAL: %#v", desired.IncidentStrategy, actual.IncidentStrategy)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "IncidentStrategy",
		})
	}
	if compareAlertPolicyMetadata(c, desired.Metadata, actual.Metadata) {
		c.Config.Logger.Infof("Detected diff in Metadata.\nDESIRED: %#v\nACTUAL: %#v", desired.Metadata, actual.Metadata)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Metadata",
		})
	}
	if !dcl.IsZeroValue(desired.Project) && !dcl.NameToSelfLink(desired.Project, actual.Project) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %#v\nACTUAL: %#v", desired.Project, actual.Project)
		diffs = append(diffs, alertPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Project",
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
	var deduped []alertPolicyDiff
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
func compareAlertPolicyDocumentationSlice(c *Client, desired, actual []AlertPolicyDocumentation) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyDocumentation, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyDocumentation(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyDocumentation, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyDocumentation(c *Client, desired, actual *AlertPolicyDocumentation) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Content == nil && desired.Content != nil && !dcl.IsEmptyValueIndirect(desired.Content) {
		c.Config.Logger.Infof("desired Content %s - but actually nil", dcl.SprintResource(desired.Content))
		return true
	}
	if !reflect.DeepEqual(desired.Content, actual.Content) && !dcl.IsZeroValue(desired.Content) && !(dcl.IsEmptyValueIndirect(desired.Content) && dcl.IsZeroValue(actual.Content)) {
		c.Config.Logger.Infof("Diff in Content. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Content), dcl.SprintResource(actual.Content))
		return true
	}
	if actual.MimeType == nil && desired.MimeType != nil && !dcl.IsEmptyValueIndirect(desired.MimeType) {
		c.Config.Logger.Infof("desired MimeType %s - but actually nil", dcl.SprintResource(desired.MimeType))
		return true
	}
	if !reflect.DeepEqual(desired.MimeType, actual.MimeType) && !dcl.IsZeroValue(desired.MimeType) && !(dcl.IsEmptyValueIndirect(desired.MimeType) && dcl.IsZeroValue(actual.MimeType)) {
		c.Config.Logger.Infof("Diff in MimeType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MimeType), dcl.SprintResource(actual.MimeType))
		return true
	}
	return false
}
func compareAlertPolicyConditionsSlice(c *Client, desired, actual []AlertPolicyConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditions(c *Client, desired, actual *AlertPolicyConditions) bool {
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
	if actual.DisplayName == nil && desired.DisplayName != nil && !dcl.IsEmptyValueIndirect(desired.DisplayName) {
		c.Config.Logger.Infof("desired DisplayName %s - but actually nil", dcl.SprintResource(desired.DisplayName))
		return true
	}
	if !reflect.DeepEqual(desired.DisplayName, actual.DisplayName) && !dcl.IsZeroValue(desired.DisplayName) && !(dcl.IsEmptyValueIndirect(desired.DisplayName) && dcl.IsZeroValue(actual.DisplayName)) {
		c.Config.Logger.Infof("Diff in DisplayName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DisplayName), dcl.SprintResource(actual.DisplayName))
		return true
	}
	if actual.ResourceStateFilter == nil && desired.ResourceStateFilter != nil && !dcl.IsEmptyValueIndirect(desired.ResourceStateFilter) {
		c.Config.Logger.Infof("desired ResourceStateFilter %s - but actually nil", dcl.SprintResource(desired.ResourceStateFilter))
		return true
	}
	if !reflect.DeepEqual(desired.ResourceStateFilter, actual.ResourceStateFilter) && !dcl.IsZeroValue(desired.ResourceStateFilter) && !(dcl.IsEmptyValueIndirect(desired.ResourceStateFilter) && dcl.IsZeroValue(actual.ResourceStateFilter)) {
		c.Config.Logger.Infof("Diff in ResourceStateFilter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceStateFilter), dcl.SprintResource(actual.ResourceStateFilter))
		return true
	}
	if actual.ConditionThreshold == nil && desired.ConditionThreshold != nil && !dcl.IsEmptyValueIndirect(desired.ConditionThreshold) {
		c.Config.Logger.Infof("desired ConditionThreshold %s - but actually nil", dcl.SprintResource(desired.ConditionThreshold))
		return true
	}
	if compareAlertPolicyConditionsConditionThreshold(c, desired.ConditionThreshold, actual.ConditionThreshold) && !dcl.IsZeroValue(desired.ConditionThreshold) {
		c.Config.Logger.Infof("Diff in ConditionThreshold. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConditionThreshold), dcl.SprintResource(actual.ConditionThreshold))
		return true
	}
	if actual.ConditionAbsent == nil && desired.ConditionAbsent != nil && !dcl.IsEmptyValueIndirect(desired.ConditionAbsent) {
		c.Config.Logger.Infof("desired ConditionAbsent %s - but actually nil", dcl.SprintResource(desired.ConditionAbsent))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsent(c, desired.ConditionAbsent, actual.ConditionAbsent) && !dcl.IsZeroValue(desired.ConditionAbsent) {
		c.Config.Logger.Infof("Diff in ConditionAbsent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConditionAbsent), dcl.SprintResource(actual.ConditionAbsent))
		return true
	}
	if actual.ConditionMatchedLog == nil && desired.ConditionMatchedLog != nil && !dcl.IsEmptyValueIndirect(desired.ConditionMatchedLog) {
		c.Config.Logger.Infof("desired ConditionMatchedLog %s - but actually nil", dcl.SprintResource(desired.ConditionMatchedLog))
		return true
	}
	if compareAlertPolicyConditionsConditionMatchedLog(c, desired.ConditionMatchedLog, actual.ConditionMatchedLog) && !dcl.IsZeroValue(desired.ConditionMatchedLog) {
		c.Config.Logger.Infof("Diff in ConditionMatchedLog. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConditionMatchedLog), dcl.SprintResource(actual.ConditionMatchedLog))
		return true
	}
	if actual.ConditionClusterOutlier == nil && desired.ConditionClusterOutlier != nil && !dcl.IsEmptyValueIndirect(desired.ConditionClusterOutlier) {
		c.Config.Logger.Infof("desired ConditionClusterOutlier %s - but actually nil", dcl.SprintResource(desired.ConditionClusterOutlier))
		return true
	}
	if compareAlertPolicyConditionsConditionClusterOutlier(c, desired.ConditionClusterOutlier, actual.ConditionClusterOutlier) && !dcl.IsZeroValue(desired.ConditionClusterOutlier) {
		c.Config.Logger.Infof("Diff in ConditionClusterOutlier. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConditionClusterOutlier), dcl.SprintResource(actual.ConditionClusterOutlier))
		return true
	}
	if actual.ConditionRate == nil && desired.ConditionRate != nil && !dcl.IsEmptyValueIndirect(desired.ConditionRate) {
		c.Config.Logger.Infof("desired ConditionRate %s - but actually nil", dcl.SprintResource(desired.ConditionRate))
		return true
	}
	if compareAlertPolicyConditionsConditionRate(c, desired.ConditionRate, actual.ConditionRate) && !dcl.IsZeroValue(desired.ConditionRate) {
		c.Config.Logger.Infof("Diff in ConditionRate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConditionRate), dcl.SprintResource(actual.ConditionRate))
		return true
	}
	if actual.ConditionUpMon == nil && desired.ConditionUpMon != nil && !dcl.IsEmptyValueIndirect(desired.ConditionUpMon) {
		c.Config.Logger.Infof("desired ConditionUpMon %s - but actually nil", dcl.SprintResource(desired.ConditionUpMon))
		return true
	}
	if compareAlertPolicyConditionsConditionUpMon(c, desired.ConditionUpMon, actual.ConditionUpMon) && !dcl.IsZeroValue(desired.ConditionUpMon) {
		c.Config.Logger.Infof("Diff in ConditionUpMon. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConditionUpMon), dcl.SprintResource(actual.ConditionUpMon))
		return true
	}
	if actual.ConditionProcessCount == nil && desired.ConditionProcessCount != nil && !dcl.IsEmptyValueIndirect(desired.ConditionProcessCount) {
		c.Config.Logger.Infof("desired ConditionProcessCount %s - but actually nil", dcl.SprintResource(desired.ConditionProcessCount))
		return true
	}
	if compareAlertPolicyConditionsConditionProcessCount(c, desired.ConditionProcessCount, actual.ConditionProcessCount) && !dcl.IsZeroValue(desired.ConditionProcessCount) {
		c.Config.Logger.Infof("Diff in ConditionProcessCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConditionProcessCount), dcl.SprintResource(actual.ConditionProcessCount))
		return true
	}
	if actual.ConditionTimeSeriesQueryLanguage == nil && desired.ConditionTimeSeriesQueryLanguage != nil && !dcl.IsEmptyValueIndirect(desired.ConditionTimeSeriesQueryLanguage) {
		c.Config.Logger.Infof("desired ConditionTimeSeriesQueryLanguage %s - but actually nil", dcl.SprintResource(desired.ConditionTimeSeriesQueryLanguage))
		return true
	}
	if compareAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, desired.ConditionTimeSeriesQueryLanguage, actual.ConditionTimeSeriesQueryLanguage) && !dcl.IsZeroValue(desired.ConditionTimeSeriesQueryLanguage) {
		c.Config.Logger.Infof("Diff in ConditionTimeSeriesQueryLanguage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConditionTimeSeriesQueryLanguage), dcl.SprintResource(actual.ConditionTimeSeriesQueryLanguage))
		return true
	}
	if actual.ConditionMonitoringQueryLanguage == nil && desired.ConditionMonitoringQueryLanguage != nil && !dcl.IsEmptyValueIndirect(desired.ConditionMonitoringQueryLanguage) {
		c.Config.Logger.Infof("desired ConditionMonitoringQueryLanguage %s - but actually nil", dcl.SprintResource(desired.ConditionMonitoringQueryLanguage))
		return true
	}
	if compareAlertPolicyConditionsConditionMonitoringQueryLanguage(c, desired.ConditionMonitoringQueryLanguage, actual.ConditionMonitoringQueryLanguage) && !dcl.IsZeroValue(desired.ConditionMonitoringQueryLanguage) {
		c.Config.Logger.Infof("Diff in ConditionMonitoringQueryLanguage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConditionMonitoringQueryLanguage), dcl.SprintResource(actual.ConditionMonitoringQueryLanguage))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThreshold) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThreshold, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThreshold(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThreshold, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThreshold(c *Client, desired, actual *AlertPolicyConditionsConditionThreshold) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Filter == nil && desired.Filter != nil && !dcl.IsEmptyValueIndirect(desired.Filter) {
		c.Config.Logger.Infof("desired Filter %s - but actually nil", dcl.SprintResource(desired.Filter))
		return true
	}
	if !reflect.DeepEqual(desired.Filter, actual.Filter) && !dcl.IsZeroValue(desired.Filter) && !(dcl.IsEmptyValueIndirect(desired.Filter) && dcl.IsZeroValue(actual.Filter)) {
		c.Config.Logger.Infof("Diff in Filter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Filter), dcl.SprintResource(actual.Filter))
		return true
	}
	if actual.Aggregations == nil && desired.Aggregations != nil && !dcl.IsEmptyValueIndirect(desired.Aggregations) {
		c.Config.Logger.Infof("desired Aggregations %s - but actually nil", dcl.SprintResource(desired.Aggregations))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdAggregationsSlice(c, desired.Aggregations, actual.Aggregations) && !dcl.IsZeroValue(desired.Aggregations) {
		c.Config.Logger.Infof("Diff in Aggregations. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Aggregations), dcl.SprintResource(actual.Aggregations))
		return true
	}
	if actual.DenominatorFilter == nil && desired.DenominatorFilter != nil && !dcl.IsEmptyValueIndirect(desired.DenominatorFilter) {
		c.Config.Logger.Infof("desired DenominatorFilter %s - but actually nil", dcl.SprintResource(desired.DenominatorFilter))
		return true
	}
	if !reflect.DeepEqual(desired.DenominatorFilter, actual.DenominatorFilter) && !dcl.IsZeroValue(desired.DenominatorFilter) && !(dcl.IsEmptyValueIndirect(desired.DenominatorFilter) && dcl.IsZeroValue(actual.DenominatorFilter)) {
		c.Config.Logger.Infof("Diff in DenominatorFilter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DenominatorFilter), dcl.SprintResource(actual.DenominatorFilter))
		return true
	}
	if actual.DenominatorAggregations == nil && desired.DenominatorAggregations != nil && !dcl.IsEmptyValueIndirect(desired.DenominatorAggregations) {
		c.Config.Logger.Infof("desired DenominatorAggregations %s - but actually nil", dcl.SprintResource(desired.DenominatorAggregations))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice(c, desired.DenominatorAggregations, actual.DenominatorAggregations) && !dcl.IsZeroValue(desired.DenominatorAggregations) {
		c.Config.Logger.Infof("Diff in DenominatorAggregations. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DenominatorAggregations), dcl.SprintResource(actual.DenominatorAggregations))
		return true
	}
	if actual.Comparison == nil && desired.Comparison != nil && !dcl.IsEmptyValueIndirect(desired.Comparison) {
		c.Config.Logger.Infof("desired Comparison %s - but actually nil", dcl.SprintResource(desired.Comparison))
		return true
	}
	if !reflect.DeepEqual(desired.Comparison, actual.Comparison) && !dcl.IsZeroValue(desired.Comparison) && !(dcl.IsEmptyValueIndirect(desired.Comparison) && dcl.IsZeroValue(actual.Comparison)) {
		c.Config.Logger.Infof("Diff in Comparison. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Comparison), dcl.SprintResource(actual.Comparison))
		return true
	}
	if actual.ThresholdValue == nil && desired.ThresholdValue != nil && !dcl.IsEmptyValueIndirect(desired.ThresholdValue) {
		c.Config.Logger.Infof("desired ThresholdValue %s - but actually nil", dcl.SprintResource(desired.ThresholdValue))
		return true
	}
	if !reflect.DeepEqual(desired.ThresholdValue, actual.ThresholdValue) && !dcl.IsZeroValue(desired.ThresholdValue) && !(dcl.IsEmptyValueIndirect(desired.ThresholdValue) && dcl.IsZeroValue(actual.ThresholdValue)) {
		c.Config.Logger.Infof("Diff in ThresholdValue. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ThresholdValue), dcl.SprintResource(actual.ThresholdValue))
		return true
	}
	if actual.Duration == nil && desired.Duration != nil && !dcl.IsEmptyValueIndirect(desired.Duration) {
		c.Config.Logger.Infof("desired Duration %s - but actually nil", dcl.SprintResource(desired.Duration))
		return true
	}
	if !reflect.DeepEqual(desired.Duration, actual.Duration) && !dcl.IsZeroValue(desired.Duration) && !(dcl.IsEmptyValueIndirect(desired.Duration) && dcl.IsZeroValue(actual.Duration)) {
		c.Config.Logger.Infof("Diff in Duration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Duration), dcl.SprintResource(actual.Duration))
		return true
	}
	if actual.Trigger == nil && desired.Trigger != nil && !dcl.IsEmptyValueIndirect(desired.Trigger) {
		c.Config.Logger.Infof("desired Trigger %s - but actually nil", dcl.SprintResource(desired.Trigger))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdTrigger(c, desired.Trigger, actual.Trigger) && !dcl.IsZeroValue(desired.Trigger) {
		c.Config.Logger.Infof("Diff in Trigger. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Trigger), dcl.SprintResource(actual.Trigger))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdAggregationsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregations) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregations, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregations(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregations, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregations(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregations) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AlignmentPeriod == nil && desired.AlignmentPeriod != nil && !dcl.IsEmptyValueIndirect(desired.AlignmentPeriod) {
		c.Config.Logger.Infof("desired AlignmentPeriod %s - but actually nil", dcl.SprintResource(desired.AlignmentPeriod))
		return true
	}
	if !reflect.DeepEqual(desired.AlignmentPeriod, actual.AlignmentPeriod) && !dcl.IsZeroValue(desired.AlignmentPeriod) && !(dcl.IsEmptyValueIndirect(desired.AlignmentPeriod) && dcl.IsZeroValue(actual.AlignmentPeriod)) {
		c.Config.Logger.Infof("Diff in AlignmentPeriod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AlignmentPeriod), dcl.SprintResource(actual.AlignmentPeriod))
		return true
	}
	if actual.PerSeriesAligner == nil && desired.PerSeriesAligner != nil && !dcl.IsEmptyValueIndirect(desired.PerSeriesAligner) {
		c.Config.Logger.Infof("desired PerSeriesAligner %s - but actually nil", dcl.SprintResource(desired.PerSeriesAligner))
		return true
	}
	if !reflect.DeepEqual(desired.PerSeriesAligner, actual.PerSeriesAligner) && !dcl.IsZeroValue(desired.PerSeriesAligner) && !(dcl.IsEmptyValueIndirect(desired.PerSeriesAligner) && dcl.IsZeroValue(actual.PerSeriesAligner)) {
		c.Config.Logger.Infof("Diff in PerSeriesAligner. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PerSeriesAligner), dcl.SprintResource(actual.PerSeriesAligner))
		return true
	}
	if actual.CrossSeriesReducer == nil && desired.CrossSeriesReducer != nil && !dcl.IsEmptyValueIndirect(desired.CrossSeriesReducer) {
		c.Config.Logger.Infof("desired CrossSeriesReducer %s - but actually nil", dcl.SprintResource(desired.CrossSeriesReducer))
		return true
	}
	if !reflect.DeepEqual(desired.CrossSeriesReducer, actual.CrossSeriesReducer) && !dcl.IsZeroValue(desired.CrossSeriesReducer) && !(dcl.IsEmptyValueIndirect(desired.CrossSeriesReducer) && dcl.IsZeroValue(actual.CrossSeriesReducer)) {
		c.Config.Logger.Infof("Diff in CrossSeriesReducer. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrossSeriesReducer), dcl.SprintResource(actual.CrossSeriesReducer))
		return true
	}
	if actual.GroupByFields == nil && desired.GroupByFields != nil && !dcl.IsEmptyValueIndirect(desired.GroupByFields) {
		c.Config.Logger.Infof("desired GroupByFields %s - but actually nil", dcl.SprintResource(desired.GroupByFields))
		return true
	}
	if !dcl.SliceEquals(desired.GroupByFields, actual.GroupByFields) && !dcl.IsZeroValue(desired.GroupByFields) {
		c.Config.Logger.Infof("Diff in GroupByFields. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GroupByFields), dcl.SprintResource(actual.GroupByFields))
		return true
	}
	if actual.ReduceFractionLessThanParams == nil && desired.ReduceFractionLessThanParams != nil && !dcl.IsEmptyValueIndirect(desired.ReduceFractionLessThanParams) {
		c.Config.Logger.Infof("desired ReduceFractionLessThanParams %s - but actually nil", dcl.SprintResource(desired.ReduceFractionLessThanParams))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, desired.ReduceFractionLessThanParams, actual.ReduceFractionLessThanParams) && !dcl.IsZeroValue(desired.ReduceFractionLessThanParams) {
		c.Config.Logger.Infof("Diff in ReduceFractionLessThanParams. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReduceFractionLessThanParams), dcl.SprintResource(actual.ReduceFractionLessThanParams))
		return true
	}
	if actual.ReduceMakeDistributionParams == nil && desired.ReduceMakeDistributionParams != nil && !dcl.IsEmptyValueIndirect(desired.ReduceMakeDistributionParams) {
		c.Config.Logger.Infof("desired ReduceMakeDistributionParams %s - but actually nil", dcl.SprintResource(desired.ReduceMakeDistributionParams))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, desired.ReduceMakeDistributionParams, actual.ReduceMakeDistributionParams) && !dcl.IsZeroValue(desired.ReduceMakeDistributionParams) {
		c.Config.Logger.Infof("Diff in ReduceMakeDistributionParams. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReduceMakeDistributionParams), dcl.SprintResource(actual.ReduceMakeDistributionParams))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Threshold == nil && desired.Threshold != nil && !dcl.IsEmptyValueIndirect(desired.Threshold) {
		c.Config.Logger.Infof("desired Threshold %s - but actually nil", dcl.SprintResource(desired.Threshold))
		return true
	}
	if !reflect.DeepEqual(desired.Threshold, actual.Threshold) && !dcl.IsZeroValue(desired.Threshold) && !(dcl.IsEmptyValueIndirect(desired.Threshold) && dcl.IsZeroValue(actual.Threshold)) {
		c.Config.Logger.Infof("Diff in Threshold. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Threshold), dcl.SprintResource(actual.Threshold))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BucketOptions == nil && desired.BucketOptions != nil && !dcl.IsEmptyValueIndirect(desired.BucketOptions) {
		c.Config.Logger.Infof("desired BucketOptions %s - but actually nil", dcl.SprintResource(desired.BucketOptions))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, desired.BucketOptions, actual.BucketOptions) && !dcl.IsZeroValue(desired.BucketOptions) {
		c.Config.Logger.Infof("Diff in BucketOptions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BucketOptions), dcl.SprintResource(actual.BucketOptions))
		return true
	}
	if actual.ExemplarSampling == nil && desired.ExemplarSampling != nil && !dcl.IsEmptyValueIndirect(desired.ExemplarSampling) {
		c.Config.Logger.Infof("desired ExemplarSampling %s - but actually nil", dcl.SprintResource(desired.ExemplarSampling))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, desired.ExemplarSampling, actual.ExemplarSampling) && !dcl.IsZeroValue(desired.ExemplarSampling) {
		c.Config.Logger.Infof("Diff in ExemplarSampling. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExemplarSampling), dcl.SprintResource(actual.ExemplarSampling))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.LinearBuckets == nil && desired.LinearBuckets != nil && !dcl.IsEmptyValueIndirect(desired.LinearBuckets) {
		c.Config.Logger.Infof("desired LinearBuckets %s - but actually nil", dcl.SprintResource(desired.LinearBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, desired.LinearBuckets, actual.LinearBuckets) && !dcl.IsZeroValue(desired.LinearBuckets) {
		c.Config.Logger.Infof("Diff in LinearBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LinearBuckets), dcl.SprintResource(actual.LinearBuckets))
		return true
	}
	if actual.ExponentialBuckets == nil && desired.ExponentialBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("desired ExponentialBuckets %s - but actually nil", dcl.SprintResource(desired.ExponentialBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, desired.ExponentialBuckets, actual.ExponentialBuckets) && !dcl.IsZeroValue(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("Diff in ExponentialBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExponentialBuckets), dcl.SprintResource(actual.ExponentialBuckets))
		return true
	}
	if actual.ExplicitBuckets == nil && desired.ExplicitBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("desired ExplicitBuckets %s - but actually nil", dcl.SprintResource(desired.ExplicitBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, desired.ExplicitBuckets, actual.ExplicitBuckets) && !dcl.IsZeroValue(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("Diff in ExplicitBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExplicitBuckets), dcl.SprintResource(actual.ExplicitBuckets))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NumFiniteBuckets == nil && desired.NumFiniteBuckets != nil && !dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) {
		c.Config.Logger.Infof("desired NumFiniteBuckets %s - but actually nil", dcl.SprintResource(desired.NumFiniteBuckets))
		return true
	}
	if !reflect.DeepEqual(desired.NumFiniteBuckets, actual.NumFiniteBuckets) && !dcl.IsZeroValue(desired.NumFiniteBuckets) && !(dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) && dcl.IsZeroValue(actual.NumFiniteBuckets)) {
		c.Config.Logger.Infof("Diff in NumFiniteBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumFiniteBuckets), dcl.SprintResource(actual.NumFiniteBuckets))
		return true
	}
	if actual.Width == nil && desired.Width != nil && !dcl.IsEmptyValueIndirect(desired.Width) {
		c.Config.Logger.Infof("desired Width %s - but actually nil", dcl.SprintResource(desired.Width))
		return true
	}
	if !reflect.DeepEqual(desired.Width, actual.Width) && !dcl.IsZeroValue(desired.Width) && !(dcl.IsEmptyValueIndirect(desired.Width) && dcl.IsZeroValue(actual.Width)) {
		c.Config.Logger.Infof("Diff in Width. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Width), dcl.SprintResource(actual.Width))
		return true
	}
	if actual.Offset == nil && desired.Offset != nil && !dcl.IsEmptyValueIndirect(desired.Offset) {
		c.Config.Logger.Infof("desired Offset %s - but actually nil", dcl.SprintResource(desired.Offset))
		return true
	}
	if !reflect.DeepEqual(desired.Offset, actual.Offset) && !dcl.IsZeroValue(desired.Offset) && !(dcl.IsEmptyValueIndirect(desired.Offset) && dcl.IsZeroValue(actual.Offset)) {
		c.Config.Logger.Infof("Diff in Offset. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Offset), dcl.SprintResource(actual.Offset))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NumFiniteBuckets == nil && desired.NumFiniteBuckets != nil && !dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) {
		c.Config.Logger.Infof("desired NumFiniteBuckets %s - but actually nil", dcl.SprintResource(desired.NumFiniteBuckets))
		return true
	}
	if !reflect.DeepEqual(desired.NumFiniteBuckets, actual.NumFiniteBuckets) && !dcl.IsZeroValue(desired.NumFiniteBuckets) && !(dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) && dcl.IsZeroValue(actual.NumFiniteBuckets)) {
		c.Config.Logger.Infof("Diff in NumFiniteBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumFiniteBuckets), dcl.SprintResource(actual.NumFiniteBuckets))
		return true
	}
	if actual.GrowthFactor == nil && desired.GrowthFactor != nil && !dcl.IsEmptyValueIndirect(desired.GrowthFactor) {
		c.Config.Logger.Infof("desired GrowthFactor %s - but actually nil", dcl.SprintResource(desired.GrowthFactor))
		return true
	}
	if !reflect.DeepEqual(desired.GrowthFactor, actual.GrowthFactor) && !dcl.IsZeroValue(desired.GrowthFactor) && !(dcl.IsEmptyValueIndirect(desired.GrowthFactor) && dcl.IsZeroValue(actual.GrowthFactor)) {
		c.Config.Logger.Infof("Diff in GrowthFactor. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GrowthFactor), dcl.SprintResource(actual.GrowthFactor))
		return true
	}
	if actual.Scale == nil && desired.Scale != nil && !dcl.IsEmptyValueIndirect(desired.Scale) {
		c.Config.Logger.Infof("desired Scale %s - but actually nil", dcl.SprintResource(desired.Scale))
		return true
	}
	if !reflect.DeepEqual(desired.Scale, actual.Scale) && !dcl.IsZeroValue(desired.Scale) && !(dcl.IsEmptyValueIndirect(desired.Scale) && dcl.IsZeroValue(actual.Scale)) {
		c.Config.Logger.Infof("Diff in Scale. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scale), dcl.SprintResource(actual.Scale))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Bounds == nil && desired.Bounds != nil && !dcl.IsEmptyValueIndirect(desired.Bounds) {
		c.Config.Logger.Infof("desired Bounds %s - but actually nil", dcl.SprintResource(desired.Bounds))
		return true
	}
	if !dcl.FloatSliceEquals(desired.Bounds, actual.Bounds) && !dcl.IsZeroValue(desired.Bounds) {
		c.Config.Logger.Infof("Diff in Bounds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Bounds), dcl.SprintResource(actual.Bounds))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MinimumValue == nil && desired.MinimumValue != nil && !dcl.IsEmptyValueIndirect(desired.MinimumValue) {
		c.Config.Logger.Infof("desired MinimumValue %s - but actually nil", dcl.SprintResource(desired.MinimumValue))
		return true
	}
	if !reflect.DeepEqual(desired.MinimumValue, actual.MinimumValue) && !dcl.IsZeroValue(desired.MinimumValue) && !(dcl.IsEmptyValueIndirect(desired.MinimumValue) && dcl.IsZeroValue(actual.MinimumValue)) {
		c.Config.Logger.Infof("Diff in MinimumValue. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinimumValue), dcl.SprintResource(actual.MinimumValue))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregations) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregations, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregations(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregations, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregations(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregations) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AlignmentPeriod == nil && desired.AlignmentPeriod != nil && !dcl.IsEmptyValueIndirect(desired.AlignmentPeriod) {
		c.Config.Logger.Infof("desired AlignmentPeriod %s - but actually nil", dcl.SprintResource(desired.AlignmentPeriod))
		return true
	}
	if !reflect.DeepEqual(desired.AlignmentPeriod, actual.AlignmentPeriod) && !dcl.IsZeroValue(desired.AlignmentPeriod) && !(dcl.IsEmptyValueIndirect(desired.AlignmentPeriod) && dcl.IsZeroValue(actual.AlignmentPeriod)) {
		c.Config.Logger.Infof("Diff in AlignmentPeriod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AlignmentPeriod), dcl.SprintResource(actual.AlignmentPeriod))
		return true
	}
	if actual.PerSeriesAligner == nil && desired.PerSeriesAligner != nil && !dcl.IsEmptyValueIndirect(desired.PerSeriesAligner) {
		c.Config.Logger.Infof("desired PerSeriesAligner %s - but actually nil", dcl.SprintResource(desired.PerSeriesAligner))
		return true
	}
	if !reflect.DeepEqual(desired.PerSeriesAligner, actual.PerSeriesAligner) && !dcl.IsZeroValue(desired.PerSeriesAligner) && !(dcl.IsEmptyValueIndirect(desired.PerSeriesAligner) && dcl.IsZeroValue(actual.PerSeriesAligner)) {
		c.Config.Logger.Infof("Diff in PerSeriesAligner. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PerSeriesAligner), dcl.SprintResource(actual.PerSeriesAligner))
		return true
	}
	if actual.CrossSeriesReducer == nil && desired.CrossSeriesReducer != nil && !dcl.IsEmptyValueIndirect(desired.CrossSeriesReducer) {
		c.Config.Logger.Infof("desired CrossSeriesReducer %s - but actually nil", dcl.SprintResource(desired.CrossSeriesReducer))
		return true
	}
	if !reflect.DeepEqual(desired.CrossSeriesReducer, actual.CrossSeriesReducer) && !dcl.IsZeroValue(desired.CrossSeriesReducer) && !(dcl.IsEmptyValueIndirect(desired.CrossSeriesReducer) && dcl.IsZeroValue(actual.CrossSeriesReducer)) {
		c.Config.Logger.Infof("Diff in CrossSeriesReducer. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrossSeriesReducer), dcl.SprintResource(actual.CrossSeriesReducer))
		return true
	}
	if actual.GroupByFields == nil && desired.GroupByFields != nil && !dcl.IsEmptyValueIndirect(desired.GroupByFields) {
		c.Config.Logger.Infof("desired GroupByFields %s - but actually nil", dcl.SprintResource(desired.GroupByFields))
		return true
	}
	if !dcl.SliceEquals(desired.GroupByFields, actual.GroupByFields) && !dcl.IsZeroValue(desired.GroupByFields) {
		c.Config.Logger.Infof("Diff in GroupByFields. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GroupByFields), dcl.SprintResource(actual.GroupByFields))
		return true
	}
	if actual.ReduceFractionLessThanParams == nil && desired.ReduceFractionLessThanParams != nil && !dcl.IsEmptyValueIndirect(desired.ReduceFractionLessThanParams) {
		c.Config.Logger.Infof("desired ReduceFractionLessThanParams %s - but actually nil", dcl.SprintResource(desired.ReduceFractionLessThanParams))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, desired.ReduceFractionLessThanParams, actual.ReduceFractionLessThanParams) && !dcl.IsZeroValue(desired.ReduceFractionLessThanParams) {
		c.Config.Logger.Infof("Diff in ReduceFractionLessThanParams. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReduceFractionLessThanParams), dcl.SprintResource(actual.ReduceFractionLessThanParams))
		return true
	}
	if actual.ReduceMakeDistributionParams == nil && desired.ReduceMakeDistributionParams != nil && !dcl.IsEmptyValueIndirect(desired.ReduceMakeDistributionParams) {
		c.Config.Logger.Infof("desired ReduceMakeDistributionParams %s - but actually nil", dcl.SprintResource(desired.ReduceMakeDistributionParams))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, desired.ReduceMakeDistributionParams, actual.ReduceMakeDistributionParams) && !dcl.IsZeroValue(desired.ReduceMakeDistributionParams) {
		c.Config.Logger.Infof("Diff in ReduceMakeDistributionParams. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReduceMakeDistributionParams), dcl.SprintResource(actual.ReduceMakeDistributionParams))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Threshold == nil && desired.Threshold != nil && !dcl.IsEmptyValueIndirect(desired.Threshold) {
		c.Config.Logger.Infof("desired Threshold %s - but actually nil", dcl.SprintResource(desired.Threshold))
		return true
	}
	if !reflect.DeepEqual(desired.Threshold, actual.Threshold) && !dcl.IsZeroValue(desired.Threshold) && !(dcl.IsEmptyValueIndirect(desired.Threshold) && dcl.IsZeroValue(actual.Threshold)) {
		c.Config.Logger.Infof("Diff in Threshold. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Threshold), dcl.SprintResource(actual.Threshold))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BucketOptions == nil && desired.BucketOptions != nil && !dcl.IsEmptyValueIndirect(desired.BucketOptions) {
		c.Config.Logger.Infof("desired BucketOptions %s - but actually nil", dcl.SprintResource(desired.BucketOptions))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, desired.BucketOptions, actual.BucketOptions) && !dcl.IsZeroValue(desired.BucketOptions) {
		c.Config.Logger.Infof("Diff in BucketOptions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BucketOptions), dcl.SprintResource(actual.BucketOptions))
		return true
	}
	if actual.ExemplarSampling == nil && desired.ExemplarSampling != nil && !dcl.IsEmptyValueIndirect(desired.ExemplarSampling) {
		c.Config.Logger.Infof("desired ExemplarSampling %s - but actually nil", dcl.SprintResource(desired.ExemplarSampling))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, desired.ExemplarSampling, actual.ExemplarSampling) && !dcl.IsZeroValue(desired.ExemplarSampling) {
		c.Config.Logger.Infof("Diff in ExemplarSampling. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExemplarSampling), dcl.SprintResource(actual.ExemplarSampling))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.LinearBuckets == nil && desired.LinearBuckets != nil && !dcl.IsEmptyValueIndirect(desired.LinearBuckets) {
		c.Config.Logger.Infof("desired LinearBuckets %s - but actually nil", dcl.SprintResource(desired.LinearBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, desired.LinearBuckets, actual.LinearBuckets) && !dcl.IsZeroValue(desired.LinearBuckets) {
		c.Config.Logger.Infof("Diff in LinearBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LinearBuckets), dcl.SprintResource(actual.LinearBuckets))
		return true
	}
	if actual.ExponentialBuckets == nil && desired.ExponentialBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("desired ExponentialBuckets %s - but actually nil", dcl.SprintResource(desired.ExponentialBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, desired.ExponentialBuckets, actual.ExponentialBuckets) && !dcl.IsZeroValue(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("Diff in ExponentialBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExponentialBuckets), dcl.SprintResource(actual.ExponentialBuckets))
		return true
	}
	if actual.ExplicitBuckets == nil && desired.ExplicitBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("desired ExplicitBuckets %s - but actually nil", dcl.SprintResource(desired.ExplicitBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, desired.ExplicitBuckets, actual.ExplicitBuckets) && !dcl.IsZeroValue(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("Diff in ExplicitBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExplicitBuckets), dcl.SprintResource(actual.ExplicitBuckets))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NumFiniteBuckets == nil && desired.NumFiniteBuckets != nil && !dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) {
		c.Config.Logger.Infof("desired NumFiniteBuckets %s - but actually nil", dcl.SprintResource(desired.NumFiniteBuckets))
		return true
	}
	if !reflect.DeepEqual(desired.NumFiniteBuckets, actual.NumFiniteBuckets) && !dcl.IsZeroValue(desired.NumFiniteBuckets) && !(dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) && dcl.IsZeroValue(actual.NumFiniteBuckets)) {
		c.Config.Logger.Infof("Diff in NumFiniteBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumFiniteBuckets), dcl.SprintResource(actual.NumFiniteBuckets))
		return true
	}
	if actual.Width == nil && desired.Width != nil && !dcl.IsEmptyValueIndirect(desired.Width) {
		c.Config.Logger.Infof("desired Width %s - but actually nil", dcl.SprintResource(desired.Width))
		return true
	}
	if !reflect.DeepEqual(desired.Width, actual.Width) && !dcl.IsZeroValue(desired.Width) && !(dcl.IsEmptyValueIndirect(desired.Width) && dcl.IsZeroValue(actual.Width)) {
		c.Config.Logger.Infof("Diff in Width. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Width), dcl.SprintResource(actual.Width))
		return true
	}
	if actual.Offset == nil && desired.Offset != nil && !dcl.IsEmptyValueIndirect(desired.Offset) {
		c.Config.Logger.Infof("desired Offset %s - but actually nil", dcl.SprintResource(desired.Offset))
		return true
	}
	if !reflect.DeepEqual(desired.Offset, actual.Offset) && !dcl.IsZeroValue(desired.Offset) && !(dcl.IsEmptyValueIndirect(desired.Offset) && dcl.IsZeroValue(actual.Offset)) {
		c.Config.Logger.Infof("Diff in Offset. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Offset), dcl.SprintResource(actual.Offset))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NumFiniteBuckets == nil && desired.NumFiniteBuckets != nil && !dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) {
		c.Config.Logger.Infof("desired NumFiniteBuckets %s - but actually nil", dcl.SprintResource(desired.NumFiniteBuckets))
		return true
	}
	if !reflect.DeepEqual(desired.NumFiniteBuckets, actual.NumFiniteBuckets) && !dcl.IsZeroValue(desired.NumFiniteBuckets) && !(dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) && dcl.IsZeroValue(actual.NumFiniteBuckets)) {
		c.Config.Logger.Infof("Diff in NumFiniteBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumFiniteBuckets), dcl.SprintResource(actual.NumFiniteBuckets))
		return true
	}
	if actual.GrowthFactor == nil && desired.GrowthFactor != nil && !dcl.IsEmptyValueIndirect(desired.GrowthFactor) {
		c.Config.Logger.Infof("desired GrowthFactor %s - but actually nil", dcl.SprintResource(desired.GrowthFactor))
		return true
	}
	if !reflect.DeepEqual(desired.GrowthFactor, actual.GrowthFactor) && !dcl.IsZeroValue(desired.GrowthFactor) && !(dcl.IsEmptyValueIndirect(desired.GrowthFactor) && dcl.IsZeroValue(actual.GrowthFactor)) {
		c.Config.Logger.Infof("Diff in GrowthFactor. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GrowthFactor), dcl.SprintResource(actual.GrowthFactor))
		return true
	}
	if actual.Scale == nil && desired.Scale != nil && !dcl.IsEmptyValueIndirect(desired.Scale) {
		c.Config.Logger.Infof("desired Scale %s - but actually nil", dcl.SprintResource(desired.Scale))
		return true
	}
	if !reflect.DeepEqual(desired.Scale, actual.Scale) && !dcl.IsZeroValue(desired.Scale) && !(dcl.IsEmptyValueIndirect(desired.Scale) && dcl.IsZeroValue(actual.Scale)) {
		c.Config.Logger.Infof("Diff in Scale. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scale), dcl.SprintResource(actual.Scale))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Bounds == nil && desired.Bounds != nil && !dcl.IsEmptyValueIndirect(desired.Bounds) {
		c.Config.Logger.Infof("desired Bounds %s - but actually nil", dcl.SprintResource(desired.Bounds))
		return true
	}
	if !dcl.FloatSliceEquals(desired.Bounds, actual.Bounds) && !dcl.IsZeroValue(desired.Bounds) {
		c.Config.Logger.Infof("Diff in Bounds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Bounds), dcl.SprintResource(actual.Bounds))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MinimumValue == nil && desired.MinimumValue != nil && !dcl.IsEmptyValueIndirect(desired.MinimumValue) {
		c.Config.Logger.Infof("desired MinimumValue %s - but actually nil", dcl.SprintResource(desired.MinimumValue))
		return true
	}
	if !reflect.DeepEqual(desired.MinimumValue, actual.MinimumValue) && !dcl.IsZeroValue(desired.MinimumValue) && !(dcl.IsEmptyValueIndirect(desired.MinimumValue) && dcl.IsZeroValue(actual.MinimumValue)) {
		c.Config.Logger.Infof("Diff in MinimumValue. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinimumValue), dcl.SprintResource(actual.MinimumValue))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionThresholdTriggerSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdTrigger) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdTrigger, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdTrigger(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdTrigger, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdTrigger(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdTrigger) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Count == nil && desired.Count != nil && !dcl.IsEmptyValueIndirect(desired.Count) {
		c.Config.Logger.Infof("desired Count %s - but actually nil", dcl.SprintResource(desired.Count))
		return true
	}
	if !reflect.DeepEqual(desired.Count, actual.Count) && !dcl.IsZeroValue(desired.Count) && !(dcl.IsEmptyValueIndirect(desired.Count) && dcl.IsZeroValue(actual.Count)) {
		c.Config.Logger.Infof("Diff in Count. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Count), dcl.SprintResource(actual.Count))
		return true
	}
	if actual.Percent == nil && desired.Percent != nil && !dcl.IsEmptyValueIndirect(desired.Percent) {
		c.Config.Logger.Infof("desired Percent %s - but actually nil", dcl.SprintResource(desired.Percent))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) && !(dcl.IsEmptyValueIndirect(desired.Percent) && dcl.IsZeroValue(actual.Percent)) {
		c.Config.Logger.Infof("Diff in Percent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsent) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsent, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsent(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsent, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsent(c *Client, desired, actual *AlertPolicyConditionsConditionAbsent) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Filter == nil && desired.Filter != nil && !dcl.IsEmptyValueIndirect(desired.Filter) {
		c.Config.Logger.Infof("desired Filter %s - but actually nil", dcl.SprintResource(desired.Filter))
		return true
	}
	if !reflect.DeepEqual(desired.Filter, actual.Filter) && !dcl.IsZeroValue(desired.Filter) && !(dcl.IsEmptyValueIndirect(desired.Filter) && dcl.IsZeroValue(actual.Filter)) {
		c.Config.Logger.Infof("Diff in Filter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Filter), dcl.SprintResource(actual.Filter))
		return true
	}
	if actual.Aggregations == nil && desired.Aggregations != nil && !dcl.IsEmptyValueIndirect(desired.Aggregations) {
		c.Config.Logger.Infof("desired Aggregations %s - but actually nil", dcl.SprintResource(desired.Aggregations))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentAggregationsSlice(c, desired.Aggregations, actual.Aggregations) && !dcl.IsZeroValue(desired.Aggregations) {
		c.Config.Logger.Infof("Diff in Aggregations. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Aggregations), dcl.SprintResource(actual.Aggregations))
		return true
	}
	if actual.Duration == nil && desired.Duration != nil && !dcl.IsEmptyValueIndirect(desired.Duration) {
		c.Config.Logger.Infof("desired Duration %s - but actually nil", dcl.SprintResource(desired.Duration))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentDuration(c, desired.Duration, actual.Duration) && !dcl.IsZeroValue(desired.Duration) {
		c.Config.Logger.Infof("Diff in Duration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Duration), dcl.SprintResource(actual.Duration))
		return true
	}
	if actual.Trigger == nil && desired.Trigger != nil && !dcl.IsEmptyValueIndirect(desired.Trigger) {
		c.Config.Logger.Infof("desired Trigger %s - but actually nil", dcl.SprintResource(desired.Trigger))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentTrigger(c, desired.Trigger, actual.Trigger) && !dcl.IsZeroValue(desired.Trigger) {
		c.Config.Logger.Infof("Diff in Trigger. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Trigger), dcl.SprintResource(actual.Trigger))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentAggregationsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregations) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregations, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregations(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregations, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregations(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregations) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AlignmentPeriod == nil && desired.AlignmentPeriod != nil && !dcl.IsEmptyValueIndirect(desired.AlignmentPeriod) {
		c.Config.Logger.Infof("desired AlignmentPeriod %s - but actually nil", dcl.SprintResource(desired.AlignmentPeriod))
		return true
	}
	if !reflect.DeepEqual(desired.AlignmentPeriod, actual.AlignmentPeriod) && !dcl.IsZeroValue(desired.AlignmentPeriod) && !(dcl.IsEmptyValueIndirect(desired.AlignmentPeriod) && dcl.IsZeroValue(actual.AlignmentPeriod)) {
		c.Config.Logger.Infof("Diff in AlignmentPeriod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AlignmentPeriod), dcl.SprintResource(actual.AlignmentPeriod))
		return true
	}
	if actual.PerSeriesAligner == nil && desired.PerSeriesAligner != nil && !dcl.IsEmptyValueIndirect(desired.PerSeriesAligner) {
		c.Config.Logger.Infof("desired PerSeriesAligner %s - but actually nil", dcl.SprintResource(desired.PerSeriesAligner))
		return true
	}
	if !reflect.DeepEqual(desired.PerSeriesAligner, actual.PerSeriesAligner) && !dcl.IsZeroValue(desired.PerSeriesAligner) && !(dcl.IsEmptyValueIndirect(desired.PerSeriesAligner) && dcl.IsZeroValue(actual.PerSeriesAligner)) {
		c.Config.Logger.Infof("Diff in PerSeriesAligner. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PerSeriesAligner), dcl.SprintResource(actual.PerSeriesAligner))
		return true
	}
	if actual.CrossSeriesReducer == nil && desired.CrossSeriesReducer != nil && !dcl.IsEmptyValueIndirect(desired.CrossSeriesReducer) {
		c.Config.Logger.Infof("desired CrossSeriesReducer %s - but actually nil", dcl.SprintResource(desired.CrossSeriesReducer))
		return true
	}
	if !reflect.DeepEqual(desired.CrossSeriesReducer, actual.CrossSeriesReducer) && !dcl.IsZeroValue(desired.CrossSeriesReducer) && !(dcl.IsEmptyValueIndirect(desired.CrossSeriesReducer) && dcl.IsZeroValue(actual.CrossSeriesReducer)) {
		c.Config.Logger.Infof("Diff in CrossSeriesReducer. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrossSeriesReducer), dcl.SprintResource(actual.CrossSeriesReducer))
		return true
	}
	if actual.GroupByFields == nil && desired.GroupByFields != nil && !dcl.IsEmptyValueIndirect(desired.GroupByFields) {
		c.Config.Logger.Infof("desired GroupByFields %s - but actually nil", dcl.SprintResource(desired.GroupByFields))
		return true
	}
	if !dcl.SliceEquals(desired.GroupByFields, actual.GroupByFields) && !dcl.IsZeroValue(desired.GroupByFields) {
		c.Config.Logger.Infof("Diff in GroupByFields. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GroupByFields), dcl.SprintResource(actual.GroupByFields))
		return true
	}
	if actual.ReduceFractionLessThanParams == nil && desired.ReduceFractionLessThanParams != nil && !dcl.IsEmptyValueIndirect(desired.ReduceFractionLessThanParams) {
		c.Config.Logger.Infof("desired ReduceFractionLessThanParams %s - but actually nil", dcl.SprintResource(desired.ReduceFractionLessThanParams))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, desired.ReduceFractionLessThanParams, actual.ReduceFractionLessThanParams) && !dcl.IsZeroValue(desired.ReduceFractionLessThanParams) {
		c.Config.Logger.Infof("Diff in ReduceFractionLessThanParams. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReduceFractionLessThanParams), dcl.SprintResource(actual.ReduceFractionLessThanParams))
		return true
	}
	if actual.ReduceMakeDistributionParams == nil && desired.ReduceMakeDistributionParams != nil && !dcl.IsEmptyValueIndirect(desired.ReduceMakeDistributionParams) {
		c.Config.Logger.Infof("desired ReduceMakeDistributionParams %s - but actually nil", dcl.SprintResource(desired.ReduceMakeDistributionParams))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, desired.ReduceMakeDistributionParams, actual.ReduceMakeDistributionParams) && !dcl.IsZeroValue(desired.ReduceMakeDistributionParams) {
		c.Config.Logger.Infof("Diff in ReduceMakeDistributionParams. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReduceMakeDistributionParams), dcl.SprintResource(actual.ReduceMakeDistributionParams))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Threshold == nil && desired.Threshold != nil && !dcl.IsEmptyValueIndirect(desired.Threshold) {
		c.Config.Logger.Infof("desired Threshold %s - but actually nil", dcl.SprintResource(desired.Threshold))
		return true
	}
	if !reflect.DeepEqual(desired.Threshold, actual.Threshold) && !dcl.IsZeroValue(desired.Threshold) && !(dcl.IsEmptyValueIndirect(desired.Threshold) && dcl.IsZeroValue(actual.Threshold)) {
		c.Config.Logger.Infof("Diff in Threshold. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Threshold), dcl.SprintResource(actual.Threshold))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BucketOptions == nil && desired.BucketOptions != nil && !dcl.IsEmptyValueIndirect(desired.BucketOptions) {
		c.Config.Logger.Infof("desired BucketOptions %s - but actually nil", dcl.SprintResource(desired.BucketOptions))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, desired.BucketOptions, actual.BucketOptions) && !dcl.IsZeroValue(desired.BucketOptions) {
		c.Config.Logger.Infof("Diff in BucketOptions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BucketOptions), dcl.SprintResource(actual.BucketOptions))
		return true
	}
	if actual.ExemplarSampling == nil && desired.ExemplarSampling != nil && !dcl.IsEmptyValueIndirect(desired.ExemplarSampling) {
		c.Config.Logger.Infof("desired ExemplarSampling %s - but actually nil", dcl.SprintResource(desired.ExemplarSampling))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, desired.ExemplarSampling, actual.ExemplarSampling) && !dcl.IsZeroValue(desired.ExemplarSampling) {
		c.Config.Logger.Infof("Diff in ExemplarSampling. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExemplarSampling), dcl.SprintResource(actual.ExemplarSampling))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.LinearBuckets == nil && desired.LinearBuckets != nil && !dcl.IsEmptyValueIndirect(desired.LinearBuckets) {
		c.Config.Logger.Infof("desired LinearBuckets %s - but actually nil", dcl.SprintResource(desired.LinearBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, desired.LinearBuckets, actual.LinearBuckets) && !dcl.IsZeroValue(desired.LinearBuckets) {
		c.Config.Logger.Infof("Diff in LinearBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LinearBuckets), dcl.SprintResource(actual.LinearBuckets))
		return true
	}
	if actual.ExponentialBuckets == nil && desired.ExponentialBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("desired ExponentialBuckets %s - but actually nil", dcl.SprintResource(desired.ExponentialBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, desired.ExponentialBuckets, actual.ExponentialBuckets) && !dcl.IsZeroValue(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("Diff in ExponentialBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExponentialBuckets), dcl.SprintResource(actual.ExponentialBuckets))
		return true
	}
	if actual.ExplicitBuckets == nil && desired.ExplicitBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("desired ExplicitBuckets %s - but actually nil", dcl.SprintResource(desired.ExplicitBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, desired.ExplicitBuckets, actual.ExplicitBuckets) && !dcl.IsZeroValue(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("Diff in ExplicitBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExplicitBuckets), dcl.SprintResource(actual.ExplicitBuckets))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NumFiniteBuckets == nil && desired.NumFiniteBuckets != nil && !dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) {
		c.Config.Logger.Infof("desired NumFiniteBuckets %s - but actually nil", dcl.SprintResource(desired.NumFiniteBuckets))
		return true
	}
	if !reflect.DeepEqual(desired.NumFiniteBuckets, actual.NumFiniteBuckets) && !dcl.IsZeroValue(desired.NumFiniteBuckets) && !(dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) && dcl.IsZeroValue(actual.NumFiniteBuckets)) {
		c.Config.Logger.Infof("Diff in NumFiniteBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumFiniteBuckets), dcl.SprintResource(actual.NumFiniteBuckets))
		return true
	}
	if actual.Width == nil && desired.Width != nil && !dcl.IsEmptyValueIndirect(desired.Width) {
		c.Config.Logger.Infof("desired Width %s - but actually nil", dcl.SprintResource(desired.Width))
		return true
	}
	if !reflect.DeepEqual(desired.Width, actual.Width) && !dcl.IsZeroValue(desired.Width) && !(dcl.IsEmptyValueIndirect(desired.Width) && dcl.IsZeroValue(actual.Width)) {
		c.Config.Logger.Infof("Diff in Width. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Width), dcl.SprintResource(actual.Width))
		return true
	}
	if actual.Offset == nil && desired.Offset != nil && !dcl.IsEmptyValueIndirect(desired.Offset) {
		c.Config.Logger.Infof("desired Offset %s - but actually nil", dcl.SprintResource(desired.Offset))
		return true
	}
	if !reflect.DeepEqual(desired.Offset, actual.Offset) && !dcl.IsZeroValue(desired.Offset) && !(dcl.IsEmptyValueIndirect(desired.Offset) && dcl.IsZeroValue(actual.Offset)) {
		c.Config.Logger.Infof("Diff in Offset. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Offset), dcl.SprintResource(actual.Offset))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NumFiniteBuckets == nil && desired.NumFiniteBuckets != nil && !dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) {
		c.Config.Logger.Infof("desired NumFiniteBuckets %s - but actually nil", dcl.SprintResource(desired.NumFiniteBuckets))
		return true
	}
	if !reflect.DeepEqual(desired.NumFiniteBuckets, actual.NumFiniteBuckets) && !dcl.IsZeroValue(desired.NumFiniteBuckets) && !(dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) && dcl.IsZeroValue(actual.NumFiniteBuckets)) {
		c.Config.Logger.Infof("Diff in NumFiniteBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumFiniteBuckets), dcl.SprintResource(actual.NumFiniteBuckets))
		return true
	}
	if actual.GrowthFactor == nil && desired.GrowthFactor != nil && !dcl.IsEmptyValueIndirect(desired.GrowthFactor) {
		c.Config.Logger.Infof("desired GrowthFactor %s - but actually nil", dcl.SprintResource(desired.GrowthFactor))
		return true
	}
	if !reflect.DeepEqual(desired.GrowthFactor, actual.GrowthFactor) && !dcl.IsZeroValue(desired.GrowthFactor) && !(dcl.IsEmptyValueIndirect(desired.GrowthFactor) && dcl.IsZeroValue(actual.GrowthFactor)) {
		c.Config.Logger.Infof("Diff in GrowthFactor. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GrowthFactor), dcl.SprintResource(actual.GrowthFactor))
		return true
	}
	if actual.Scale == nil && desired.Scale != nil && !dcl.IsEmptyValueIndirect(desired.Scale) {
		c.Config.Logger.Infof("desired Scale %s - but actually nil", dcl.SprintResource(desired.Scale))
		return true
	}
	if !reflect.DeepEqual(desired.Scale, actual.Scale) && !dcl.IsZeroValue(desired.Scale) && !(dcl.IsEmptyValueIndirect(desired.Scale) && dcl.IsZeroValue(actual.Scale)) {
		c.Config.Logger.Infof("Diff in Scale. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scale), dcl.SprintResource(actual.Scale))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Bounds == nil && desired.Bounds != nil && !dcl.IsEmptyValueIndirect(desired.Bounds) {
		c.Config.Logger.Infof("desired Bounds %s - but actually nil", dcl.SprintResource(desired.Bounds))
		return true
	}
	if !dcl.FloatSliceEquals(desired.Bounds, actual.Bounds) && !dcl.IsZeroValue(desired.Bounds) {
		c.Config.Logger.Infof("Diff in Bounds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Bounds), dcl.SprintResource(actual.Bounds))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MinimumValue == nil && desired.MinimumValue != nil && !dcl.IsEmptyValueIndirect(desired.MinimumValue) {
		c.Config.Logger.Infof("desired MinimumValue %s - but actually nil", dcl.SprintResource(desired.MinimumValue))
		return true
	}
	if !reflect.DeepEqual(desired.MinimumValue, actual.MinimumValue) && !dcl.IsZeroValue(desired.MinimumValue) && !(dcl.IsEmptyValueIndirect(desired.MinimumValue) && dcl.IsZeroValue(actual.MinimumValue)) {
		c.Config.Logger.Infof("Diff in MinimumValue. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinimumValue), dcl.SprintResource(actual.MinimumValue))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionAbsentDurationSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentDuration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentDuration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentDuration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentDuration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentDuration(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentDuration) bool {
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
func compareAlertPolicyConditionsConditionAbsentTriggerSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentTrigger) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentTrigger, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentTrigger(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentTrigger, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentTrigger(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentTrigger) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Count == nil && desired.Count != nil && !dcl.IsEmptyValueIndirect(desired.Count) {
		c.Config.Logger.Infof("desired Count %s - but actually nil", dcl.SprintResource(desired.Count))
		return true
	}
	if !reflect.DeepEqual(desired.Count, actual.Count) && !dcl.IsZeroValue(desired.Count) && !(dcl.IsEmptyValueIndirect(desired.Count) && dcl.IsZeroValue(actual.Count)) {
		c.Config.Logger.Infof("Diff in Count. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Count), dcl.SprintResource(actual.Count))
		return true
	}
	if actual.Percent == nil && desired.Percent != nil && !dcl.IsEmptyValueIndirect(desired.Percent) {
		c.Config.Logger.Infof("desired Percent %s - but actually nil", dcl.SprintResource(desired.Percent))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) && !(dcl.IsEmptyValueIndirect(desired.Percent) && dcl.IsZeroValue(actual.Percent)) {
		c.Config.Logger.Infof("Diff in Percent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionMatchedLogSlice(c *Client, desired, actual []AlertPolicyConditionsConditionMatchedLog) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionMatchedLog, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionMatchedLog(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionMatchedLog, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionMatchedLog(c *Client, desired, actual *AlertPolicyConditionsConditionMatchedLog) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Filter == nil && desired.Filter != nil && !dcl.IsEmptyValueIndirect(desired.Filter) {
		c.Config.Logger.Infof("desired Filter %s - but actually nil", dcl.SprintResource(desired.Filter))
		return true
	}
	if !reflect.DeepEqual(desired.Filter, actual.Filter) && !dcl.IsZeroValue(desired.Filter) && !(dcl.IsEmptyValueIndirect(desired.Filter) && dcl.IsZeroValue(actual.Filter)) {
		c.Config.Logger.Infof("Diff in Filter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Filter), dcl.SprintResource(actual.Filter))
		return true
	}
	if actual.LabelExtractors == nil && desired.LabelExtractors != nil && !dcl.IsEmptyValueIndirect(desired.LabelExtractors) {
		c.Config.Logger.Infof("desired LabelExtractors %s - but actually nil", dcl.SprintResource(desired.LabelExtractors))
		return true
	}
	if !reflect.DeepEqual(desired.LabelExtractors, actual.LabelExtractors) && !dcl.IsZeroValue(desired.LabelExtractors) {
		c.Config.Logger.Infof("Diff in LabelExtractors. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LabelExtractors), dcl.SprintResource(actual.LabelExtractors))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionClusterOutlierSlice(c *Client, desired, actual []AlertPolicyConditionsConditionClusterOutlier) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionClusterOutlier, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionClusterOutlier(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionClusterOutlier, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionClusterOutlier(c *Client, desired, actual *AlertPolicyConditionsConditionClusterOutlier) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Filter == nil && desired.Filter != nil && !dcl.IsEmptyValueIndirect(desired.Filter) {
		c.Config.Logger.Infof("desired Filter %s - but actually nil", dcl.SprintResource(desired.Filter))
		return true
	}
	if !reflect.DeepEqual(desired.Filter, actual.Filter) && !dcl.IsZeroValue(desired.Filter) && !(dcl.IsEmptyValueIndirect(desired.Filter) && dcl.IsZeroValue(actual.Filter)) {
		c.Config.Logger.Infof("Diff in Filter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Filter), dcl.SprintResource(actual.Filter))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRate) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRate, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRate(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRate, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRate(c *Client, desired, actual *AlertPolicyConditionsConditionRate) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Filter == nil && desired.Filter != nil && !dcl.IsEmptyValueIndirect(desired.Filter) {
		c.Config.Logger.Infof("desired Filter %s - but actually nil", dcl.SprintResource(desired.Filter))
		return true
	}
	if !reflect.DeepEqual(desired.Filter, actual.Filter) && !dcl.IsZeroValue(desired.Filter) && !(dcl.IsEmptyValueIndirect(desired.Filter) && dcl.IsZeroValue(actual.Filter)) {
		c.Config.Logger.Infof("Diff in Filter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Filter), dcl.SprintResource(actual.Filter))
		return true
	}
	if actual.Aggregations == nil && desired.Aggregations != nil && !dcl.IsEmptyValueIndirect(desired.Aggregations) {
		c.Config.Logger.Infof("desired Aggregations %s - but actually nil", dcl.SprintResource(desired.Aggregations))
		return true
	}
	if compareAlertPolicyConditionsConditionRateAggregationsSlice(c, desired.Aggregations, actual.Aggregations) && !dcl.IsZeroValue(desired.Aggregations) {
		c.Config.Logger.Infof("Diff in Aggregations. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Aggregations), dcl.SprintResource(actual.Aggregations))
		return true
	}
	if actual.Comparison == nil && desired.Comparison != nil && !dcl.IsEmptyValueIndirect(desired.Comparison) {
		c.Config.Logger.Infof("desired Comparison %s - but actually nil", dcl.SprintResource(desired.Comparison))
		return true
	}
	if !reflect.DeepEqual(desired.Comparison, actual.Comparison) && !dcl.IsZeroValue(desired.Comparison) && !(dcl.IsEmptyValueIndirect(desired.Comparison) && dcl.IsZeroValue(actual.Comparison)) {
		c.Config.Logger.Infof("Diff in Comparison. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Comparison), dcl.SprintResource(actual.Comparison))
		return true
	}
	if actual.ThresholdValue == nil && desired.ThresholdValue != nil && !dcl.IsEmptyValueIndirect(desired.ThresholdValue) {
		c.Config.Logger.Infof("desired ThresholdValue %s - but actually nil", dcl.SprintResource(desired.ThresholdValue))
		return true
	}
	if !reflect.DeepEqual(desired.ThresholdValue, actual.ThresholdValue) && !dcl.IsZeroValue(desired.ThresholdValue) && !(dcl.IsEmptyValueIndirect(desired.ThresholdValue) && dcl.IsZeroValue(actual.ThresholdValue)) {
		c.Config.Logger.Infof("Diff in ThresholdValue. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ThresholdValue), dcl.SprintResource(actual.ThresholdValue))
		return true
	}
	if actual.TimeWindow == nil && desired.TimeWindow != nil && !dcl.IsEmptyValueIndirect(desired.TimeWindow) {
		c.Config.Logger.Infof("desired TimeWindow %s - but actually nil", dcl.SprintResource(desired.TimeWindow))
		return true
	}
	if compareAlertPolicyConditionsConditionRateTimeWindow(c, desired.TimeWindow, actual.TimeWindow) && !dcl.IsZeroValue(desired.TimeWindow) {
		c.Config.Logger.Infof("Diff in TimeWindow. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TimeWindow), dcl.SprintResource(actual.TimeWindow))
		return true
	}
	if actual.Trigger == nil && desired.Trigger != nil && !dcl.IsEmptyValueIndirect(desired.Trigger) {
		c.Config.Logger.Infof("desired Trigger %s - but actually nil", dcl.SprintResource(desired.Trigger))
		return true
	}
	if compareAlertPolicyConditionsConditionRateTrigger(c, desired.Trigger, actual.Trigger) && !dcl.IsZeroValue(desired.Trigger) {
		c.Config.Logger.Infof("Diff in Trigger. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Trigger), dcl.SprintResource(actual.Trigger))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateAggregationsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregations) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregations, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregations(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregations, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregations(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregations) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AlignmentPeriod == nil && desired.AlignmentPeriod != nil && !dcl.IsEmptyValueIndirect(desired.AlignmentPeriod) {
		c.Config.Logger.Infof("desired AlignmentPeriod %s - but actually nil", dcl.SprintResource(desired.AlignmentPeriod))
		return true
	}
	if !reflect.DeepEqual(desired.AlignmentPeriod, actual.AlignmentPeriod) && !dcl.IsZeroValue(desired.AlignmentPeriod) && !(dcl.IsEmptyValueIndirect(desired.AlignmentPeriod) && dcl.IsZeroValue(actual.AlignmentPeriod)) {
		c.Config.Logger.Infof("Diff in AlignmentPeriod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AlignmentPeriod), dcl.SprintResource(actual.AlignmentPeriod))
		return true
	}
	if actual.PerSeriesAligner == nil && desired.PerSeriesAligner != nil && !dcl.IsEmptyValueIndirect(desired.PerSeriesAligner) {
		c.Config.Logger.Infof("desired PerSeriesAligner %s - but actually nil", dcl.SprintResource(desired.PerSeriesAligner))
		return true
	}
	if !reflect.DeepEqual(desired.PerSeriesAligner, actual.PerSeriesAligner) && !dcl.IsZeroValue(desired.PerSeriesAligner) && !(dcl.IsEmptyValueIndirect(desired.PerSeriesAligner) && dcl.IsZeroValue(actual.PerSeriesAligner)) {
		c.Config.Logger.Infof("Diff in PerSeriesAligner. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PerSeriesAligner), dcl.SprintResource(actual.PerSeriesAligner))
		return true
	}
	if actual.CrossSeriesReducer == nil && desired.CrossSeriesReducer != nil && !dcl.IsEmptyValueIndirect(desired.CrossSeriesReducer) {
		c.Config.Logger.Infof("desired CrossSeriesReducer %s - but actually nil", dcl.SprintResource(desired.CrossSeriesReducer))
		return true
	}
	if !reflect.DeepEqual(desired.CrossSeriesReducer, actual.CrossSeriesReducer) && !dcl.IsZeroValue(desired.CrossSeriesReducer) && !(dcl.IsEmptyValueIndirect(desired.CrossSeriesReducer) && dcl.IsZeroValue(actual.CrossSeriesReducer)) {
		c.Config.Logger.Infof("Diff in CrossSeriesReducer. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrossSeriesReducer), dcl.SprintResource(actual.CrossSeriesReducer))
		return true
	}
	if actual.GroupByFields == nil && desired.GroupByFields != nil && !dcl.IsEmptyValueIndirect(desired.GroupByFields) {
		c.Config.Logger.Infof("desired GroupByFields %s - but actually nil", dcl.SprintResource(desired.GroupByFields))
		return true
	}
	if !dcl.SliceEquals(desired.GroupByFields, actual.GroupByFields) && !dcl.IsZeroValue(desired.GroupByFields) {
		c.Config.Logger.Infof("Diff in GroupByFields. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GroupByFields), dcl.SprintResource(actual.GroupByFields))
		return true
	}
	if actual.ReduceFractionLessThanParams == nil && desired.ReduceFractionLessThanParams != nil && !dcl.IsEmptyValueIndirect(desired.ReduceFractionLessThanParams) {
		c.Config.Logger.Infof("desired ReduceFractionLessThanParams %s - but actually nil", dcl.SprintResource(desired.ReduceFractionLessThanParams))
		return true
	}
	if compareAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, desired.ReduceFractionLessThanParams, actual.ReduceFractionLessThanParams) && !dcl.IsZeroValue(desired.ReduceFractionLessThanParams) {
		c.Config.Logger.Infof("Diff in ReduceFractionLessThanParams. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReduceFractionLessThanParams), dcl.SprintResource(actual.ReduceFractionLessThanParams))
		return true
	}
	if actual.ReduceMakeDistributionParams == nil && desired.ReduceMakeDistributionParams != nil && !dcl.IsEmptyValueIndirect(desired.ReduceMakeDistributionParams) {
		c.Config.Logger.Infof("desired ReduceMakeDistributionParams %s - but actually nil", dcl.SprintResource(desired.ReduceMakeDistributionParams))
		return true
	}
	if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, desired.ReduceMakeDistributionParams, actual.ReduceMakeDistributionParams) && !dcl.IsZeroValue(desired.ReduceMakeDistributionParams) {
		c.Config.Logger.Infof("Diff in ReduceMakeDistributionParams. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReduceMakeDistributionParams), dcl.SprintResource(actual.ReduceMakeDistributionParams))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Threshold == nil && desired.Threshold != nil && !dcl.IsEmptyValueIndirect(desired.Threshold) {
		c.Config.Logger.Infof("desired Threshold %s - but actually nil", dcl.SprintResource(desired.Threshold))
		return true
	}
	if !reflect.DeepEqual(desired.Threshold, actual.Threshold) && !dcl.IsZeroValue(desired.Threshold) && !(dcl.IsEmptyValueIndirect(desired.Threshold) && dcl.IsZeroValue(actual.Threshold)) {
		c.Config.Logger.Infof("Diff in Threshold. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Threshold), dcl.SprintResource(actual.Threshold))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BucketOptions == nil && desired.BucketOptions != nil && !dcl.IsEmptyValueIndirect(desired.BucketOptions) {
		c.Config.Logger.Infof("desired BucketOptions %s - but actually nil", dcl.SprintResource(desired.BucketOptions))
		return true
	}
	if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, desired.BucketOptions, actual.BucketOptions) && !dcl.IsZeroValue(desired.BucketOptions) {
		c.Config.Logger.Infof("Diff in BucketOptions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BucketOptions), dcl.SprintResource(actual.BucketOptions))
		return true
	}
	if actual.ExemplarSampling == nil && desired.ExemplarSampling != nil && !dcl.IsEmptyValueIndirect(desired.ExemplarSampling) {
		c.Config.Logger.Infof("desired ExemplarSampling %s - but actually nil", dcl.SprintResource(desired.ExemplarSampling))
		return true
	}
	if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, desired.ExemplarSampling, actual.ExemplarSampling) && !dcl.IsZeroValue(desired.ExemplarSampling) {
		c.Config.Logger.Infof("Diff in ExemplarSampling. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExemplarSampling), dcl.SprintResource(actual.ExemplarSampling))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.LinearBuckets == nil && desired.LinearBuckets != nil && !dcl.IsEmptyValueIndirect(desired.LinearBuckets) {
		c.Config.Logger.Infof("desired LinearBuckets %s - but actually nil", dcl.SprintResource(desired.LinearBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, desired.LinearBuckets, actual.LinearBuckets) && !dcl.IsZeroValue(desired.LinearBuckets) {
		c.Config.Logger.Infof("Diff in LinearBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LinearBuckets), dcl.SprintResource(actual.LinearBuckets))
		return true
	}
	if actual.ExponentialBuckets == nil && desired.ExponentialBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("desired ExponentialBuckets %s - but actually nil", dcl.SprintResource(desired.ExponentialBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, desired.ExponentialBuckets, actual.ExponentialBuckets) && !dcl.IsZeroValue(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("Diff in ExponentialBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExponentialBuckets), dcl.SprintResource(actual.ExponentialBuckets))
		return true
	}
	if actual.ExplicitBuckets == nil && desired.ExplicitBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("desired ExplicitBuckets %s - but actually nil", dcl.SprintResource(desired.ExplicitBuckets))
		return true
	}
	if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, desired.ExplicitBuckets, actual.ExplicitBuckets) && !dcl.IsZeroValue(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("Diff in ExplicitBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExplicitBuckets), dcl.SprintResource(actual.ExplicitBuckets))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NumFiniteBuckets == nil && desired.NumFiniteBuckets != nil && !dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) {
		c.Config.Logger.Infof("desired NumFiniteBuckets %s - but actually nil", dcl.SprintResource(desired.NumFiniteBuckets))
		return true
	}
	if !reflect.DeepEqual(desired.NumFiniteBuckets, actual.NumFiniteBuckets) && !dcl.IsZeroValue(desired.NumFiniteBuckets) && !(dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) && dcl.IsZeroValue(actual.NumFiniteBuckets)) {
		c.Config.Logger.Infof("Diff in NumFiniteBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumFiniteBuckets), dcl.SprintResource(actual.NumFiniteBuckets))
		return true
	}
	if actual.Width == nil && desired.Width != nil && !dcl.IsEmptyValueIndirect(desired.Width) {
		c.Config.Logger.Infof("desired Width %s - but actually nil", dcl.SprintResource(desired.Width))
		return true
	}
	if !reflect.DeepEqual(desired.Width, actual.Width) && !dcl.IsZeroValue(desired.Width) && !(dcl.IsEmptyValueIndirect(desired.Width) && dcl.IsZeroValue(actual.Width)) {
		c.Config.Logger.Infof("Diff in Width. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Width), dcl.SprintResource(actual.Width))
		return true
	}
	if actual.Offset == nil && desired.Offset != nil && !dcl.IsEmptyValueIndirect(desired.Offset) {
		c.Config.Logger.Infof("desired Offset %s - but actually nil", dcl.SprintResource(desired.Offset))
		return true
	}
	if !reflect.DeepEqual(desired.Offset, actual.Offset) && !dcl.IsZeroValue(desired.Offset) && !(dcl.IsEmptyValueIndirect(desired.Offset) && dcl.IsZeroValue(actual.Offset)) {
		c.Config.Logger.Infof("Diff in Offset. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Offset), dcl.SprintResource(actual.Offset))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.NumFiniteBuckets == nil && desired.NumFiniteBuckets != nil && !dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) {
		c.Config.Logger.Infof("desired NumFiniteBuckets %s - but actually nil", dcl.SprintResource(desired.NumFiniteBuckets))
		return true
	}
	if !reflect.DeepEqual(desired.NumFiniteBuckets, actual.NumFiniteBuckets) && !dcl.IsZeroValue(desired.NumFiniteBuckets) && !(dcl.IsEmptyValueIndirect(desired.NumFiniteBuckets) && dcl.IsZeroValue(actual.NumFiniteBuckets)) {
		c.Config.Logger.Infof("Diff in NumFiniteBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NumFiniteBuckets), dcl.SprintResource(actual.NumFiniteBuckets))
		return true
	}
	if actual.GrowthFactor == nil && desired.GrowthFactor != nil && !dcl.IsEmptyValueIndirect(desired.GrowthFactor) {
		c.Config.Logger.Infof("desired GrowthFactor %s - but actually nil", dcl.SprintResource(desired.GrowthFactor))
		return true
	}
	if !reflect.DeepEqual(desired.GrowthFactor, actual.GrowthFactor) && !dcl.IsZeroValue(desired.GrowthFactor) && !(dcl.IsEmptyValueIndirect(desired.GrowthFactor) && dcl.IsZeroValue(actual.GrowthFactor)) {
		c.Config.Logger.Infof("Diff in GrowthFactor. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GrowthFactor), dcl.SprintResource(actual.GrowthFactor))
		return true
	}
	if actual.Scale == nil && desired.Scale != nil && !dcl.IsEmptyValueIndirect(desired.Scale) {
		c.Config.Logger.Infof("desired Scale %s - but actually nil", dcl.SprintResource(desired.Scale))
		return true
	}
	if !reflect.DeepEqual(desired.Scale, actual.Scale) && !dcl.IsZeroValue(desired.Scale) && !(dcl.IsEmptyValueIndirect(desired.Scale) && dcl.IsZeroValue(actual.Scale)) {
		c.Config.Logger.Infof("Diff in Scale. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scale), dcl.SprintResource(actual.Scale))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Bounds == nil && desired.Bounds != nil && !dcl.IsEmptyValueIndirect(desired.Bounds) {
		c.Config.Logger.Infof("desired Bounds %s - but actually nil", dcl.SprintResource(desired.Bounds))
		return true
	}
	if !dcl.FloatSliceEquals(desired.Bounds, actual.Bounds) && !dcl.IsZeroValue(desired.Bounds) {
		c.Config.Logger.Infof("Diff in Bounds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Bounds), dcl.SprintResource(actual.Bounds))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MinimumValue == nil && desired.MinimumValue != nil && !dcl.IsEmptyValueIndirect(desired.MinimumValue) {
		c.Config.Logger.Infof("desired MinimumValue %s - but actually nil", dcl.SprintResource(desired.MinimumValue))
		return true
	}
	if !reflect.DeepEqual(desired.MinimumValue, actual.MinimumValue) && !dcl.IsZeroValue(desired.MinimumValue) && !(dcl.IsEmptyValueIndirect(desired.MinimumValue) && dcl.IsZeroValue(actual.MinimumValue)) {
		c.Config.Logger.Infof("Diff in MinimumValue. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinimumValue), dcl.SprintResource(actual.MinimumValue))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionRateTimeWindowSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateTimeWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateTimeWindow, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateTimeWindow(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateTimeWindow, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateTimeWindow(c *Client, desired, actual *AlertPolicyConditionsConditionRateTimeWindow) bool {
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
func compareAlertPolicyConditionsConditionRateTriggerSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateTrigger) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateTrigger, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateTrigger(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateTrigger, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateTrigger(c *Client, desired, actual *AlertPolicyConditionsConditionRateTrigger) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Count == nil && desired.Count != nil && !dcl.IsEmptyValueIndirect(desired.Count) {
		c.Config.Logger.Infof("desired Count %s - but actually nil", dcl.SprintResource(desired.Count))
		return true
	}
	if !reflect.DeepEqual(desired.Count, actual.Count) && !dcl.IsZeroValue(desired.Count) && !(dcl.IsEmptyValueIndirect(desired.Count) && dcl.IsZeroValue(actual.Count)) {
		c.Config.Logger.Infof("Diff in Count. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Count), dcl.SprintResource(actual.Count))
		return true
	}
	if actual.Percent == nil && desired.Percent != nil && !dcl.IsEmptyValueIndirect(desired.Percent) {
		c.Config.Logger.Infof("desired Percent %s - but actually nil", dcl.SprintResource(desired.Percent))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) && !(dcl.IsEmptyValueIndirect(desired.Percent) && dcl.IsZeroValue(actual.Percent)) {
		c.Config.Logger.Infof("Diff in Percent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionUpMonSlice(c *Client, desired, actual []AlertPolicyConditionsConditionUpMon) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionUpMon, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionUpMon(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionUpMon, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionUpMon(c *Client, desired, actual *AlertPolicyConditionsConditionUpMon) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Filter == nil && desired.Filter != nil && !dcl.IsEmptyValueIndirect(desired.Filter) {
		c.Config.Logger.Infof("desired Filter %s - but actually nil", dcl.SprintResource(desired.Filter))
		return true
	}
	if !reflect.DeepEqual(desired.Filter, actual.Filter) && !dcl.IsZeroValue(desired.Filter) && !(dcl.IsEmptyValueIndirect(desired.Filter) && dcl.IsZeroValue(actual.Filter)) {
		c.Config.Logger.Infof("Diff in Filter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Filter), dcl.SprintResource(actual.Filter))
		return true
	}
	if actual.EndpointId == nil && desired.EndpointId != nil && !dcl.IsEmptyValueIndirect(desired.EndpointId) {
		c.Config.Logger.Infof("desired EndpointId %s - but actually nil", dcl.SprintResource(desired.EndpointId))
		return true
	}
	if !reflect.DeepEqual(desired.EndpointId, actual.EndpointId) && !dcl.IsZeroValue(desired.EndpointId) && !(dcl.IsEmptyValueIndirect(desired.EndpointId) && dcl.IsZeroValue(actual.EndpointId)) {
		c.Config.Logger.Infof("Diff in EndpointId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EndpointId), dcl.SprintResource(actual.EndpointId))
		return true
	}
	if actual.CheckId == nil && desired.CheckId != nil && !dcl.IsEmptyValueIndirect(desired.CheckId) {
		c.Config.Logger.Infof("desired CheckId %s - but actually nil", dcl.SprintResource(desired.CheckId))
		return true
	}
	if !reflect.DeepEqual(desired.CheckId, actual.CheckId) && !dcl.IsZeroValue(desired.CheckId) && !(dcl.IsEmptyValueIndirect(desired.CheckId) && dcl.IsZeroValue(actual.CheckId)) {
		c.Config.Logger.Infof("Diff in CheckId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CheckId), dcl.SprintResource(actual.CheckId))
		return true
	}
	if actual.Duration == nil && desired.Duration != nil && !dcl.IsEmptyValueIndirect(desired.Duration) {
		c.Config.Logger.Infof("desired Duration %s - but actually nil", dcl.SprintResource(desired.Duration))
		return true
	}
	if compareAlertPolicyConditionsConditionUpMonDuration(c, desired.Duration, actual.Duration) && !dcl.IsZeroValue(desired.Duration) {
		c.Config.Logger.Infof("Diff in Duration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Duration), dcl.SprintResource(actual.Duration))
		return true
	}
	if actual.Trigger == nil && desired.Trigger != nil && !dcl.IsEmptyValueIndirect(desired.Trigger) {
		c.Config.Logger.Infof("desired Trigger %s - but actually nil", dcl.SprintResource(desired.Trigger))
		return true
	}
	if compareAlertPolicyConditionsConditionUpMonTrigger(c, desired.Trigger, actual.Trigger) && !dcl.IsZeroValue(desired.Trigger) {
		c.Config.Logger.Infof("Diff in Trigger. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Trigger), dcl.SprintResource(actual.Trigger))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionUpMonDurationSlice(c *Client, desired, actual []AlertPolicyConditionsConditionUpMonDuration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionUpMonDuration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionUpMonDuration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionUpMonDuration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionUpMonDuration(c *Client, desired, actual *AlertPolicyConditionsConditionUpMonDuration) bool {
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
func compareAlertPolicyConditionsConditionUpMonTriggerSlice(c *Client, desired, actual []AlertPolicyConditionsConditionUpMonTrigger) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionUpMonTrigger, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionUpMonTrigger(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionUpMonTrigger, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionUpMonTrigger(c *Client, desired, actual *AlertPolicyConditionsConditionUpMonTrigger) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Count == nil && desired.Count != nil && !dcl.IsEmptyValueIndirect(desired.Count) {
		c.Config.Logger.Infof("desired Count %s - but actually nil", dcl.SprintResource(desired.Count))
		return true
	}
	if !reflect.DeepEqual(desired.Count, actual.Count) && !dcl.IsZeroValue(desired.Count) && !(dcl.IsEmptyValueIndirect(desired.Count) && dcl.IsZeroValue(actual.Count)) {
		c.Config.Logger.Infof("Diff in Count. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Count), dcl.SprintResource(actual.Count))
		return true
	}
	if actual.Percent == nil && desired.Percent != nil && !dcl.IsEmptyValueIndirect(desired.Percent) {
		c.Config.Logger.Infof("desired Percent %s - but actually nil", dcl.SprintResource(desired.Percent))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) && !(dcl.IsEmptyValueIndirect(desired.Percent) && dcl.IsZeroValue(actual.Percent)) {
		c.Config.Logger.Infof("Diff in Percent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionProcessCountSlice(c *Client, desired, actual []AlertPolicyConditionsConditionProcessCount) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionProcessCount, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionProcessCount(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionProcessCount, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionProcessCount(c *Client, desired, actual *AlertPolicyConditionsConditionProcessCount) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Process == nil && desired.Process != nil && !dcl.IsEmptyValueIndirect(desired.Process) {
		c.Config.Logger.Infof("desired Process %s - but actually nil", dcl.SprintResource(desired.Process))
		return true
	}
	if !reflect.DeepEqual(desired.Process, actual.Process) && !dcl.IsZeroValue(desired.Process) && !(dcl.IsEmptyValueIndirect(desired.Process) && dcl.IsZeroValue(actual.Process)) {
		c.Config.Logger.Infof("Diff in Process. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Process), dcl.SprintResource(actual.Process))
		return true
	}
	if actual.User == nil && desired.User != nil && !dcl.IsEmptyValueIndirect(desired.User) {
		c.Config.Logger.Infof("desired User %s - but actually nil", dcl.SprintResource(desired.User))
		return true
	}
	if !reflect.DeepEqual(desired.User, actual.User) && !dcl.IsZeroValue(desired.User) && !(dcl.IsEmptyValueIndirect(desired.User) && dcl.IsZeroValue(actual.User)) {
		c.Config.Logger.Infof("Diff in User. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.User), dcl.SprintResource(actual.User))
		return true
	}
	if actual.Filter == nil && desired.Filter != nil && !dcl.IsEmptyValueIndirect(desired.Filter) {
		c.Config.Logger.Infof("desired Filter %s - but actually nil", dcl.SprintResource(desired.Filter))
		return true
	}
	if !reflect.DeepEqual(desired.Filter, actual.Filter) && !dcl.IsZeroValue(desired.Filter) && !(dcl.IsEmptyValueIndirect(desired.Filter) && dcl.IsZeroValue(actual.Filter)) {
		c.Config.Logger.Infof("Diff in Filter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Filter), dcl.SprintResource(actual.Filter))
		return true
	}
	if actual.Comparison == nil && desired.Comparison != nil && !dcl.IsEmptyValueIndirect(desired.Comparison) {
		c.Config.Logger.Infof("desired Comparison %s - but actually nil", dcl.SprintResource(desired.Comparison))
		return true
	}
	if !reflect.DeepEqual(desired.Comparison, actual.Comparison) && !dcl.IsZeroValue(desired.Comparison) && !(dcl.IsEmptyValueIndirect(desired.Comparison) && dcl.IsZeroValue(actual.Comparison)) {
		c.Config.Logger.Infof("Diff in Comparison. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Comparison), dcl.SprintResource(actual.Comparison))
		return true
	}
	if actual.ProcessCountThreshold == nil && desired.ProcessCountThreshold != nil && !dcl.IsEmptyValueIndirect(desired.ProcessCountThreshold) {
		c.Config.Logger.Infof("desired ProcessCountThreshold %s - but actually nil", dcl.SprintResource(desired.ProcessCountThreshold))
		return true
	}
	if !reflect.DeepEqual(desired.ProcessCountThreshold, actual.ProcessCountThreshold) && !dcl.IsZeroValue(desired.ProcessCountThreshold) && !(dcl.IsEmptyValueIndirect(desired.ProcessCountThreshold) && dcl.IsZeroValue(actual.ProcessCountThreshold)) {
		c.Config.Logger.Infof("Diff in ProcessCountThreshold. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ProcessCountThreshold), dcl.SprintResource(actual.ProcessCountThreshold))
		return true
	}
	if actual.Trigger == nil && desired.Trigger != nil && !dcl.IsEmptyValueIndirect(desired.Trigger) {
		c.Config.Logger.Infof("desired Trigger %s - but actually nil", dcl.SprintResource(desired.Trigger))
		return true
	}
	if compareAlertPolicyConditionsConditionProcessCountTrigger(c, desired.Trigger, actual.Trigger) && !dcl.IsZeroValue(desired.Trigger) {
		c.Config.Logger.Infof("Diff in Trigger. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Trigger), dcl.SprintResource(actual.Trigger))
		return true
	}
	if actual.Duration == nil && desired.Duration != nil && !dcl.IsEmptyValueIndirect(desired.Duration) {
		c.Config.Logger.Infof("desired Duration %s - but actually nil", dcl.SprintResource(desired.Duration))
		return true
	}
	if compareAlertPolicyConditionsConditionProcessCountDuration(c, desired.Duration, actual.Duration) && !dcl.IsZeroValue(desired.Duration) {
		c.Config.Logger.Infof("Diff in Duration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Duration), dcl.SprintResource(actual.Duration))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionProcessCountTriggerSlice(c *Client, desired, actual []AlertPolicyConditionsConditionProcessCountTrigger) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionProcessCountTrigger, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionProcessCountTrigger(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionProcessCountTrigger, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionProcessCountTrigger(c *Client, desired, actual *AlertPolicyConditionsConditionProcessCountTrigger) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Count == nil && desired.Count != nil && !dcl.IsEmptyValueIndirect(desired.Count) {
		c.Config.Logger.Infof("desired Count %s - but actually nil", dcl.SprintResource(desired.Count))
		return true
	}
	if !reflect.DeepEqual(desired.Count, actual.Count) && !dcl.IsZeroValue(desired.Count) && !(dcl.IsEmptyValueIndirect(desired.Count) && dcl.IsZeroValue(actual.Count)) {
		c.Config.Logger.Infof("Diff in Count. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Count), dcl.SprintResource(actual.Count))
		return true
	}
	if actual.Percent == nil && desired.Percent != nil && !dcl.IsEmptyValueIndirect(desired.Percent) {
		c.Config.Logger.Infof("desired Percent %s - but actually nil", dcl.SprintResource(desired.Percent))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) && !(dcl.IsEmptyValueIndirect(desired.Percent) && dcl.IsZeroValue(actual.Percent)) {
		c.Config.Logger.Infof("Diff in Percent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionProcessCountDurationSlice(c *Client, desired, actual []AlertPolicyConditionsConditionProcessCountDuration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionProcessCountDuration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionProcessCountDuration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionProcessCountDuration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionProcessCountDuration(c *Client, desired, actual *AlertPolicyConditionsConditionProcessCountDuration) bool {
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
func compareAlertPolicyConditionsConditionTimeSeriesQueryLanguageSlice(c *Client, desired, actual []AlertPolicyConditionsConditionTimeSeriesQueryLanguage) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionTimeSeriesQueryLanguage, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionTimeSeriesQueryLanguage, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionTimeSeriesQueryLanguage(c *Client, desired, actual *AlertPolicyConditionsConditionTimeSeriesQueryLanguage) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Query == nil && desired.Query != nil && !dcl.IsEmptyValueIndirect(desired.Query) {
		c.Config.Logger.Infof("desired Query %s - but actually nil", dcl.SprintResource(desired.Query))
		return true
	}
	if !reflect.DeepEqual(desired.Query, actual.Query) && !dcl.IsZeroValue(desired.Query) && !(dcl.IsEmptyValueIndirect(desired.Query) && dcl.IsZeroValue(actual.Query)) {
		c.Config.Logger.Infof("Diff in Query. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Query), dcl.SprintResource(actual.Query))
		return true
	}
	if actual.Summary == nil && desired.Summary != nil && !dcl.IsEmptyValueIndirect(desired.Summary) {
		c.Config.Logger.Infof("desired Summary %s - but actually nil", dcl.SprintResource(desired.Summary))
		return true
	}
	if !reflect.DeepEqual(desired.Summary, actual.Summary) && !dcl.IsZeroValue(desired.Summary) && !(dcl.IsEmptyValueIndirect(desired.Summary) && dcl.IsZeroValue(actual.Summary)) {
		c.Config.Logger.Infof("Diff in Summary. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Summary), dcl.SprintResource(actual.Summary))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionMonitoringQueryLanguageSlice(c *Client, desired, actual []AlertPolicyConditionsConditionMonitoringQueryLanguage) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionMonitoringQueryLanguage, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionMonitoringQueryLanguage(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionMonitoringQueryLanguage, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionMonitoringQueryLanguage(c *Client, desired, actual *AlertPolicyConditionsConditionMonitoringQueryLanguage) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Query == nil && desired.Query != nil && !dcl.IsEmptyValueIndirect(desired.Query) {
		c.Config.Logger.Infof("desired Query %s - but actually nil", dcl.SprintResource(desired.Query))
		return true
	}
	if !reflect.DeepEqual(desired.Query, actual.Query) && !dcl.IsZeroValue(desired.Query) && !(dcl.IsEmptyValueIndirect(desired.Query) && dcl.IsZeroValue(actual.Query)) {
		c.Config.Logger.Infof("Diff in Query. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Query), dcl.SprintResource(actual.Query))
		return true
	}
	if actual.Duration == nil && desired.Duration != nil && !dcl.IsEmptyValueIndirect(desired.Duration) {
		c.Config.Logger.Infof("desired Duration %s - but actually nil", dcl.SprintResource(desired.Duration))
		return true
	}
	if compareAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, desired.Duration, actual.Duration) && !dcl.IsZeroValue(desired.Duration) {
		c.Config.Logger.Infof("Diff in Duration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Duration), dcl.SprintResource(actual.Duration))
		return true
	}
	if actual.Trigger == nil && desired.Trigger != nil && !dcl.IsEmptyValueIndirect(desired.Trigger) {
		c.Config.Logger.Infof("desired Trigger %s - but actually nil", dcl.SprintResource(desired.Trigger))
		return true
	}
	if compareAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, desired.Trigger, actual.Trigger) && !dcl.IsZeroValue(desired.Trigger) {
		c.Config.Logger.Infof("Diff in Trigger. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Trigger), dcl.SprintResource(actual.Trigger))
		return true
	}
	return false
}
func compareAlertPolicyConditionsConditionMonitoringQueryLanguageDurationSlice(c *Client, desired, actual []AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionMonitoringQueryLanguageDuration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionMonitoringQueryLanguageDuration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(c *Client, desired, actual *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) bool {
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
func compareAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerSlice(c *Client, desired, actual []AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(c *Client, desired, actual *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Count == nil && desired.Count != nil && !dcl.IsEmptyValueIndirect(desired.Count) {
		c.Config.Logger.Infof("desired Count %s - but actually nil", dcl.SprintResource(desired.Count))
		return true
	}
	if !reflect.DeepEqual(desired.Count, actual.Count) && !dcl.IsZeroValue(desired.Count) && !(dcl.IsEmptyValueIndirect(desired.Count) && dcl.IsZeroValue(actual.Count)) {
		c.Config.Logger.Infof("Diff in Count. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Count), dcl.SprintResource(actual.Count))
		return true
	}
	if actual.Percent == nil && desired.Percent != nil && !dcl.IsEmptyValueIndirect(desired.Percent) {
		c.Config.Logger.Infof("desired Percent %s - but actually nil", dcl.SprintResource(desired.Percent))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) && !(dcl.IsEmptyValueIndirect(desired.Percent) && dcl.IsZeroValue(actual.Percent)) {
		c.Config.Logger.Infof("Diff in Percent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	return false
}
func compareAlertPolicyEnabledSlice(c *Client, desired, actual []AlertPolicyEnabled) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyEnabled, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyEnabled(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyEnabled, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyEnabled(c *Client, desired, actual *AlertPolicyEnabled) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
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
func compareAlertPolicyValiditySlice(c *Client, desired, actual []AlertPolicyValidity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyValidity, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyValidity(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyValidity, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyValidity(c *Client, desired, actual *AlertPolicyValidity) bool {
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
	if compareAlertPolicyValidityDetailsSlice(c, desired.Details, actual.Details) && !dcl.IsZeroValue(desired.Details) {
		c.Config.Logger.Infof("Diff in Details. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Details), dcl.SprintResource(actual.Details))
		return true
	}
	return false
}
func compareAlertPolicyValidityDetailsSlice(c *Client, desired, actual []AlertPolicyValidityDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyValidityDetails, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyValidityDetails(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyValidityDetails, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyValidityDetails(c *Client, desired, actual *AlertPolicyValidityDetails) bool {
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
func compareAlertPolicyCreationRecordSlice(c *Client, desired, actual []AlertPolicyCreationRecord) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyCreationRecord, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyCreationRecord(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyCreationRecord, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyCreationRecord(c *Client, desired, actual *AlertPolicyCreationRecord) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MutateTime == nil && desired.MutateTime != nil && !dcl.IsEmptyValueIndirect(desired.MutateTime) {
		c.Config.Logger.Infof("desired MutateTime %s - but actually nil", dcl.SprintResource(desired.MutateTime))
		return true
	}
	if compareAlertPolicyCreationRecordMutateTime(c, desired.MutateTime, actual.MutateTime) && !dcl.IsZeroValue(desired.MutateTime) {
		c.Config.Logger.Infof("Diff in MutateTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MutateTime), dcl.SprintResource(actual.MutateTime))
		return true
	}
	if actual.MutatedBy == nil && desired.MutatedBy != nil && !dcl.IsEmptyValueIndirect(desired.MutatedBy) {
		c.Config.Logger.Infof("desired MutatedBy %s - but actually nil", dcl.SprintResource(desired.MutatedBy))
		return true
	}
	if !reflect.DeepEqual(desired.MutatedBy, actual.MutatedBy) && !dcl.IsZeroValue(desired.MutatedBy) && !(dcl.IsEmptyValueIndirect(desired.MutatedBy) && dcl.IsZeroValue(actual.MutatedBy)) {
		c.Config.Logger.Infof("Diff in MutatedBy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MutatedBy), dcl.SprintResource(actual.MutatedBy))
		return true
	}
	return false
}
func compareAlertPolicyCreationRecordMutateTimeSlice(c *Client, desired, actual []AlertPolicyCreationRecordMutateTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyCreationRecordMutateTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyCreationRecordMutateTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyCreationRecordMutateTime, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyCreationRecordMutateTime(c *Client, desired, actual *AlertPolicyCreationRecordMutateTime) bool {
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
func compareAlertPolicyMutationRecordSlice(c *Client, desired, actual []AlertPolicyMutationRecord) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyMutationRecord, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyMutationRecord(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyMutationRecord, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyMutationRecord(c *Client, desired, actual *AlertPolicyMutationRecord) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MutateTime == nil && desired.MutateTime != nil && !dcl.IsEmptyValueIndirect(desired.MutateTime) {
		c.Config.Logger.Infof("desired MutateTime %s - but actually nil", dcl.SprintResource(desired.MutateTime))
		return true
	}
	if compareAlertPolicyMutationRecordMutateTime(c, desired.MutateTime, actual.MutateTime) && !dcl.IsZeroValue(desired.MutateTime) {
		c.Config.Logger.Infof("Diff in MutateTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MutateTime), dcl.SprintResource(actual.MutateTime))
		return true
	}
	if actual.MutatedBy == nil && desired.MutatedBy != nil && !dcl.IsEmptyValueIndirect(desired.MutatedBy) {
		c.Config.Logger.Infof("desired MutatedBy %s - but actually nil", dcl.SprintResource(desired.MutatedBy))
		return true
	}
	if !reflect.DeepEqual(desired.MutatedBy, actual.MutatedBy) && !dcl.IsZeroValue(desired.MutatedBy) && !(dcl.IsEmptyValueIndirect(desired.MutatedBy) && dcl.IsZeroValue(actual.MutatedBy)) {
		c.Config.Logger.Infof("Diff in MutatedBy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MutatedBy), dcl.SprintResource(actual.MutatedBy))
		return true
	}
	return false
}
func compareAlertPolicyMutationRecordMutateTimeSlice(c *Client, desired, actual []AlertPolicyMutationRecordMutateTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyMutationRecordMutateTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyMutationRecordMutateTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyMutationRecordMutateTime, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyMutationRecordMutateTime(c *Client, desired, actual *AlertPolicyMutationRecordMutateTime) bool {
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
func compareAlertPolicyIncidentStrategySlice(c *Client, desired, actual []AlertPolicyIncidentStrategy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyIncidentStrategy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyIncidentStrategy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyIncidentStrategy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyIncidentStrategy(c *Client, desired, actual *AlertPolicyIncidentStrategy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) && !(dcl.IsEmptyValueIndirect(desired.Type) && dcl.IsZeroValue(actual.Type)) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}
func compareAlertPolicyMetadataSlice(c *Client, desired, actual []AlertPolicyMetadata) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyMetadata, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyMetadata(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyMetadata, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyMetadata(c *Client, desired, actual *AlertPolicyMetadata) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.SloNames == nil && desired.SloNames != nil && !dcl.IsEmptyValueIndirect(desired.SloNames) {
		c.Config.Logger.Infof("desired SloNames %s - but actually nil", dcl.SprintResource(desired.SloNames))
		return true
	}
	if !dcl.SliceEquals(desired.SloNames, actual.SloNames) && !dcl.IsZeroValue(desired.SloNames) {
		c.Config.Logger.Infof("Diff in SloNames. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SloNames), dcl.SprintResource(actual.SloNames))
		return true
	}
	return false
}
func compareAlertPolicyConditionsResourceStateFilterEnumSlice(c *Client, desired, actual []AlertPolicyConditionsResourceStateFilterEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsResourceStateFilterEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsResourceStateFilterEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsResourceStateFilterEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsResourceStateFilterEnum(c *Client, desired, actual *AlertPolicyConditionsResourceStateFilterEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionThresholdComparisonEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionThresholdComparisonEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionThresholdComparisonEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionThresholdComparisonEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionThresholdComparisonEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionThresholdComparisonEnum(c *Client, desired, actual *AlertPolicyConditionsConditionThresholdComparisonEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(c *Client, desired, actual *AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(c *Client, desired, actual *AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionRateComparisonEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionRateComparisonEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionRateComparisonEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionRateComparisonEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionRateComparisonEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionRateComparisonEnum(c *Client, desired, actual *AlertPolicyConditionsConditionRateComparisonEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyConditionsConditionProcessCountComparisonEnumSlice(c *Client, desired, actual []AlertPolicyConditionsConditionProcessCountComparisonEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyConditionsConditionProcessCountComparisonEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyConditionsConditionProcessCountComparisonEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyConditionsConditionProcessCountComparisonEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyConditionsConditionProcessCountComparisonEnum(c *Client, desired, actual *AlertPolicyConditionsConditionProcessCountComparisonEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyCombinerEnumSlice(c *Client, desired, actual []AlertPolicyCombinerEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyCombinerEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyCombinerEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyCombinerEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyCombinerEnum(c *Client, desired, actual *AlertPolicyCombinerEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAlertPolicyIncidentStrategyTypeEnumSlice(c *Client, desired, actual []AlertPolicyIncidentStrategyTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AlertPolicyIncidentStrategyTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAlertPolicyIncidentStrategyTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AlertPolicyIncidentStrategyTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAlertPolicyIncidentStrategyTypeEnum(c *Client, desired, actual *AlertPolicyIncidentStrategyTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *AlertPolicy) urlNormalized() *AlertPolicy {
	normalized := deepcopy.Copy(*r).(AlertPolicy)
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
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["documentation"] = v
	}
	if v := f.UserLabels; !dcl.IsEmptyValueIndirect(v) {
		m["userLabels"] = v
	}
	if v, err := expandAlertPolicyConditionsSlice(c, f.Conditions); err != nil {
		return nil, fmt.Errorf("error expanding Conditions into conditions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v, err := expandAlertPolicyValidity(c, f.Validity); err != nil {
		return nil, fmt.Errorf("error expanding Validity into validity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["validity"] = v
	}
	if v := f.NotificationChannels; !dcl.IsEmptyValueIndirect(v) {
		m["notificationChannels"] = v
	}
	if v, err := expandAlertPolicyCreationRecord(c, f.CreationRecord); err != nil {
		return nil, fmt.Errorf("error expanding CreationRecord into creationRecord: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["creationRecord"] = v
	}
	if v, err := expandAlertPolicyMutationRecord(c, f.MutationRecord); err != nil {
		return nil, fmt.Errorf("error expanding MutationRecord into mutationRecord: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mutationRecord"] = v
	}
	if v, err := expandAlertPolicyIncidentStrategy(c, f.IncidentStrategy); err != nil {
		return nil, fmt.Errorf("error expanding IncidentStrategy into incidentStrategy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["incidentStrategy"] = v
	}
	if v, err := expandAlertPolicyMetadata(c, f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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

	r := &AlertPolicy{}
	r.Name = dcl.FlattenSecretValue(m["name"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.Documentation = flattenAlertPolicyDocumentation(c, m["documentation"])
	r.UserLabels = dcl.FlattenKeyValuePairs(m["userLabels"])
	r.Conditions = flattenAlertPolicyConditionsSlice(c, m["conditions"])
	r.Combiner = flattenAlertPolicyCombinerEnum(m["combiner"])
	r.Disabled = dcl.FlattenBool(m["disabled"])
	r.Enabled = flattenAlertPolicyEnabled(c, m["enabled"])
	r.Validity = flattenAlertPolicyValidity(c, m["validity"])
	r.NotificationChannels = dcl.FlattenStringSlice(m["notificationChannels"])
	r.CreationRecord = flattenAlertPolicyCreationRecord(c, m["creationRecord"])
	r.MutationRecord = flattenAlertPolicyMutationRecord(c, m["mutationRecord"])
	r.IncidentStrategy = flattenAlertPolicyIncidentStrategy(c, m["incidentStrategy"])
	r.Metadata = flattenAlertPolicyMetadata(c, m["metadata"])
	r.Project = dcl.FlattenString(m["project"])

	return r
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
		items = append(items, *flattenAlertPolicyConditionsResourceStateFilterEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionThresholdComparisonEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionRateComparisonEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyConditionsConditionProcessCountComparisonEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyCombinerEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenAlertPolicyIncidentStrategyTypeEnum(item.(map[string]interface{})))
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
