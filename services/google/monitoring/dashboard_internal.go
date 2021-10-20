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
)

func (r *Dashboard) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"GridLayout", "MosaicLayout", "RowLayout", "ColumnLayout"}, r.GridLayout, r.MosaicLayout, r.RowLayout, r.ColumnLayout); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "displayName"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.GridLayout) {
		if err := r.GridLayout.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MosaicLayout) {
		if err := r.MosaicLayout.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RowLayout) {
		if err := r.RowLayout.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ColumnLayout) {
		if err := r.ColumnLayout.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardGridLayout) validate() error {
	return nil
}
func (r *DashboardMosaicLayout) validate() error {
	return nil
}
func (r *DashboardMosaicLayoutTiles) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Widget) {
		if err := r.Widget.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardRowLayout) validate() error {
	return nil
}
func (r *DashboardRowLayoutRows) validate() error {
	return nil
}
func (r *DashboardColumnLayout) validate() error {
	return nil
}
func (r *DashboardColumnLayoutColumns) validate() error {
	return nil
}
func (r *DashboardWidget) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"XyChart", "Scorecard", "Text", "Blank"}, r.XyChart, r.Scorecard, r.Text, r.Blank); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.XyChart) {
		if err := r.XyChart.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Scorecard) {
		if err := r.Scorecard.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Text) {
		if err := r.Text.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Blank) {
		if err := r.Blank.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetXyChart) validate() error {
	if err := dcl.Required(r, "dataSets"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.XAxis) {
		if err := r.XAxis.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.YAxis) {
		if err := r.YAxis.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ChartOptions) {
		if err := r.ChartOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetXyChartDataSets) validate() error {
	if err := dcl.Required(r, "timeSeriesQuery"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.TimeSeriesQuery) {
		if err := r.TimeSeriesQuery.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQuery) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TimeSeriesFilter) {
		if err := r.TimeSeriesFilter.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TimeSeriesFilterRatio) {
		if err := r.TimeSeriesFilterRatio.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) validate() error {
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Aggregation) {
		if err := r.Aggregation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecondaryAggregation) {
		if err := r.SecondaryAggregation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PickTimeSeriesFilter) {
		if err := r.PickTimeSeriesFilter.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Numerator) {
		if err := r.Numerator.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Denominator) {
		if err := r.Denominator.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecondaryAggregation) {
		if err := r.SecondaryAggregation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PickTimeSeriesFilter) {
		if err := r.PickTimeSeriesFilter.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) validate() error {
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Aggregation) {
		if err := r.Aggregation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) validate() error {
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Aggregation) {
		if err := r.Aggregation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartThresholds) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartXAxis) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartYAxis) validate() error {
	return nil
}
func (r *DashboardWidgetXyChartChartOptions) validate() error {
	return nil
}
func (r *DashboardWidgetScorecard) validate() error {
	if err := dcl.Required(r, "timeSeriesQuery"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.TimeSeriesQuery) {
		if err := r.TimeSeriesQuery.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.GaugeView) {
		if err := r.GaugeView.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SparkChartView) {
		if err := r.SparkChartView.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQuery) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TimeSeriesFilter) {
		if err := r.TimeSeriesFilter.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TimeSeriesFilterRatio) {
		if err := r.TimeSeriesFilterRatio.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) validate() error {
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Aggregation) {
		if err := r.Aggregation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecondaryAggregation) {
		if err := r.SecondaryAggregation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PickTimeSeriesFilter) {
		if err := r.PickTimeSeriesFilter.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) validate() error {
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Numerator) {
		if err := r.Numerator.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Denominator) {
		if err := r.Denominator.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecondaryAggregation) {
		if err := r.SecondaryAggregation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PickTimeSeriesFilter) {
		if err := r.PickTimeSeriesFilter.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) validate() error {
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Aggregation) {
		if err := r.Aggregation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) validate() error {
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Aggregation) {
		if err := r.Aggregation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) validate() error {
	return nil
}
func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) validate() error {
	return nil
}
func (r *DashboardWidgetScorecardGaugeView) validate() error {
	return nil
}
func (r *DashboardWidgetScorecardSparkChartView) validate() error {
	if err := dcl.Required(r, "sparkChartType"); err != nil {
		return err
	}
	return nil
}
func (r *DashboardWidgetScorecardThresholds) validate() error {
	return nil
}
func (r *DashboardWidgetText) validate() error {
	return nil
}
func (r *DashboardWidgetBlank) validate() error {
	return nil
}
func (r *Dashboard) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://monitoring.googleapis.com/v1/", params)
}

func (r *Dashboard) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/dashboards/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Dashboard) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/dashboards", nr.basePath(), userBasePath, params), nil

}

func (r *Dashboard) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/dashboards", nr.basePath(), userBasePath, params), nil

}

func (r *Dashboard) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/dashboards/{{name}}", nr.basePath(), userBasePath, params), nil
}

// dashboardApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type dashboardApiOperation interface {
	do(context.Context, *Dashboard, *Client) error
}

// newUpdateDashboardUpdateDashboardRequest creates a request for an
// Dashboard resource's UpdateDashboard update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDashboardUpdateDashboardRequest(ctx context.Context, f *Dashboard, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v, err := expandDashboardGridLayout(c, f.GridLayout); err != nil {
		return nil, fmt.Errorf("error expanding GridLayout into gridLayout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["gridLayout"] = v
	}
	if v, err := expandDashboardMosaicLayout(c, f.MosaicLayout); err != nil {
		return nil, fmt.Errorf("error expanding MosaicLayout into mosaicLayout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["mosaicLayout"] = v
	}
	if v, err := expandDashboardRowLayout(c, f.RowLayout); err != nil {
		return nil, fmt.Errorf("error expanding RowLayout into rowLayout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["rowLayout"] = v
	}
	if v, err := expandDashboardColumnLayout(c, f.ColumnLayout); err != nil {
		return nil, fmt.Errorf("error expanding ColumnLayout into columnLayout: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["columnLayout"] = v
	}
	b, err := c.getDashboardRaw(ctx, f)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	req["name"] = fmt.Sprintf("projects/%s/dashboards/%s", *f.Project, *f.Name)

	return req, nil
}

// marshalUpdateDashboardUpdateDashboardRequest converts the update into
// the final JSON request body.
func marshalUpdateDashboardUpdateDashboardRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateDashboardUpdateDashboardOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDashboardUpdateDashboardOperation) do(ctx context.Context, r *Dashboard, c *Client) error {
	_, err := c.GetDashboard(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateDashboard")
	if err != nil {
		return err
	}

	req, err := newUpdateDashboardUpdateDashboardRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateDashboardUpdateDashboardRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listDashboardRaw(ctx context.Context, r *Dashboard, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != DashboardMaxPage {
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

type listDashboardOperation struct {
	Dashboards []map[string]interface{} `json:"dashboards"`
	Token      string                   `json:"nextPageToken"`
}

func (c *Client) listDashboard(ctx context.Context, r *Dashboard, pageToken string, pageSize int32) ([]*Dashboard, string, error) {
	b, err := c.listDashboardRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listDashboardOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Dashboard
	for _, v := range m.Dashboards {
		res, err := unmarshalMapDashboard(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllDashboard(ctx context.Context, f func(*Dashboard) bool, resources []*Dashboard) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteDashboard(ctx, res)
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

type deleteDashboardOperation struct{}

func (op *deleteDashboardOperation) do(ctx context.Context, r *Dashboard, c *Client) error {
	r, err := c.GetDashboard(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Dashboard not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetDashboard checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Dashboard: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetDashboard(ctx, r)
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
type createDashboardOperation struct {
	response map[string]interface{}
}

func (op *createDashboardOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createDashboardOperation) do(ctx context.Context, r *Dashboard, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
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
	m["name"] = fmt.Sprintf("projects/%s/dashboards/%s", *normalized.Project, *normalized.Name)

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

	if _, err := c.GetDashboard(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDashboardRaw(ctx context.Context, r *Dashboard) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
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

func (c *Client) dashboardDiffsForRawDesired(ctx context.Context, rawDesired *Dashboard, opts ...dcl.ApplyOption) (initial, desired *Dashboard, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Dashboard
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Dashboard); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Dashboard, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetDashboard(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Dashboard resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Dashboard resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Dashboard resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeDashboardDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Dashboard: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Dashboard: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeDashboardInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Dashboard: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeDashboardDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Dashboard: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffDashboard(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeDashboardInitialState(rawInitial, rawDesired *Dashboard) (*Dashboard, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.GridLayout) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.MosaicLayout, rawInitial.RowLayout, rawInitial.ColumnLayout) {
			rawInitial.GridLayout = EmptyDashboardGridLayout
		}
	}

	if !dcl.IsZeroValue(rawInitial.MosaicLayout) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.GridLayout, rawInitial.RowLayout, rawInitial.ColumnLayout) {
			rawInitial.MosaicLayout = EmptyDashboardMosaicLayout
		}
	}

	if !dcl.IsZeroValue(rawInitial.RowLayout) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.GridLayout, rawInitial.MosaicLayout, rawInitial.ColumnLayout) {
			rawInitial.RowLayout = EmptyDashboardRowLayout
		}
	}

	if !dcl.IsZeroValue(rawInitial.ColumnLayout) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.GridLayout, rawInitial.MosaicLayout, rawInitial.RowLayout) {
			rawInitial.ColumnLayout = EmptyDashboardColumnLayout
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

func canonicalizeDashboardDesiredState(rawDesired, rawInitial *Dashboard, opts ...dcl.ApplyOption) (*Dashboard, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.GridLayout = canonicalizeDashboardGridLayout(rawDesired.GridLayout, nil, opts...)
		rawDesired.MosaicLayout = canonicalizeDashboardMosaicLayout(rawDesired.MosaicLayout, nil, opts...)
		rawDesired.RowLayout = canonicalizeDashboardRowLayout(rawDesired.RowLayout, nil, opts...)
		rawDesired.ColumnLayout = canonicalizeDashboardColumnLayout(rawDesired.ColumnLayout, nil, opts...)

		return rawDesired, nil
	}

	if rawDesired.GridLayout != nil || rawInitial.GridLayout != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.MosaicLayout, rawDesired.RowLayout, rawDesired.ColumnLayout) {
			rawDesired.GridLayout = nil
			rawInitial.GridLayout = nil
		}
	}

	if rawDesired.MosaicLayout != nil || rawInitial.MosaicLayout != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.GridLayout, rawDesired.RowLayout, rawDesired.ColumnLayout) {
			rawDesired.MosaicLayout = nil
			rawInitial.MosaicLayout = nil
		}
	}

	if rawDesired.RowLayout != nil || rawInitial.RowLayout != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.GridLayout, rawDesired.MosaicLayout, rawDesired.ColumnLayout) {
			rawDesired.RowLayout = nil
			rawInitial.RowLayout = nil
		}
	}

	if rawDesired.ColumnLayout != nil || rawInitial.ColumnLayout != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.GridLayout, rawDesired.MosaicLayout, rawDesired.RowLayout) {
			rawDesired.ColumnLayout = nil
			rawInitial.ColumnLayout = nil
		}
	}

	canonicalDesired := &Dashboard{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	canonicalDesired.GridLayout = canonicalizeDashboardGridLayout(rawDesired.GridLayout, rawInitial.GridLayout, opts...)
	canonicalDesired.MosaicLayout = canonicalizeDashboardMosaicLayout(rawDesired.MosaicLayout, rawInitial.MosaicLayout, opts...)
	canonicalDesired.RowLayout = canonicalizeDashboardRowLayout(rawDesired.RowLayout, rawInitial.RowLayout, opts...)
	canonicalDesired.ColumnLayout = canonicalizeDashboardColumnLayout(rawDesired.ColumnLayout, rawInitial.ColumnLayout, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeDashboardNewState(c *Client, rawNew, rawDesired *Dashboard) (*Dashboard, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsNotReturnedByServer(rawNew.DisplayName) && dcl.IsNotReturnedByServer(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.GridLayout) && dcl.IsNotReturnedByServer(rawDesired.GridLayout) {
		rawNew.GridLayout = rawDesired.GridLayout
	} else {
		rawNew.GridLayout = canonicalizeNewDashboardGridLayout(c, rawDesired.GridLayout, rawNew.GridLayout)
	}

	if dcl.IsNotReturnedByServer(rawNew.MosaicLayout) && dcl.IsNotReturnedByServer(rawDesired.MosaicLayout) {
		rawNew.MosaicLayout = rawDesired.MosaicLayout
	} else {
		rawNew.MosaicLayout = canonicalizeNewDashboardMosaicLayout(c, rawDesired.MosaicLayout, rawNew.MosaicLayout)
	}

	if dcl.IsNotReturnedByServer(rawNew.RowLayout) && dcl.IsNotReturnedByServer(rawDesired.RowLayout) {
		rawNew.RowLayout = rawDesired.RowLayout
	} else {
		rawNew.RowLayout = canonicalizeNewDashboardRowLayout(c, rawDesired.RowLayout, rawNew.RowLayout)
	}

	if dcl.IsNotReturnedByServer(rawNew.ColumnLayout) && dcl.IsNotReturnedByServer(rawDesired.ColumnLayout) {
		rawNew.ColumnLayout = rawDesired.ColumnLayout
	} else {
		rawNew.ColumnLayout = canonicalizeNewDashboardColumnLayout(c, rawDesired.ColumnLayout, rawNew.ColumnLayout)
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsNotReturnedByServer(rawNew.Etag) && dcl.IsNotReturnedByServer(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	return rawNew, nil
}

func canonicalizeDashboardGridLayout(des, initial *DashboardGridLayout, opts ...dcl.ApplyOption) *DashboardGridLayout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardGridLayout{}

	if dcl.IsZeroValue(des.Columns) {
		des.Columns = initial.Columns
	} else {
		cDes.Columns = des.Columns
	}
	if dcl.IsZeroValue(des.Widgets) {
		des.Widgets = initial.Widgets
	} else {
		cDes.Widgets = des.Widgets
	}

	return cDes
}

func canonicalizeDashboardGridLayoutSlice(des, initial []DashboardGridLayout, opts ...dcl.ApplyOption) []DashboardGridLayout {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardGridLayout, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardGridLayout(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardGridLayout, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardGridLayout(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardGridLayout(c *Client, des, nw *DashboardGridLayout) *DashboardGridLayout {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardGridLayout while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDashboardGridLayoutSet(c *Client, des, nw []DashboardGridLayout) []DashboardGridLayout {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardGridLayout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardGridLayoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardGridLayoutSlice(c *Client, des, nw []DashboardGridLayout) []DashboardGridLayout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardGridLayout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardGridLayout(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardMosaicLayout(des, initial *DashboardMosaicLayout, opts ...dcl.ApplyOption) *DashboardMosaicLayout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardMosaicLayout{}

	if dcl.IsZeroValue(des.Columns) {
		des.Columns = initial.Columns
	} else {
		cDes.Columns = des.Columns
	}
	cDes.Tiles = canonicalizeDashboardMosaicLayoutTilesSlice(des.Tiles, initial.Tiles, opts...)

	return cDes
}

func canonicalizeDashboardMosaicLayoutSlice(des, initial []DashboardMosaicLayout, opts ...dcl.ApplyOption) []DashboardMosaicLayout {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardMosaicLayout, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardMosaicLayout(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardMosaicLayout, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardMosaicLayout(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardMosaicLayout(c *Client, des, nw *DashboardMosaicLayout) *DashboardMosaicLayout {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardMosaicLayout while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Tiles = canonicalizeNewDashboardMosaicLayoutTilesSlice(c, des.Tiles, nw.Tiles)

	return nw
}

func canonicalizeNewDashboardMosaicLayoutSet(c *Client, des, nw []DashboardMosaicLayout) []DashboardMosaicLayout {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardMosaicLayout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardMosaicLayoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardMosaicLayoutSlice(c *Client, des, nw []DashboardMosaicLayout) []DashboardMosaicLayout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardMosaicLayout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardMosaicLayout(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardMosaicLayoutTiles(des, initial *DashboardMosaicLayoutTiles, opts ...dcl.ApplyOption) *DashboardMosaicLayoutTiles {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardMosaicLayoutTiles{}

	if dcl.IsZeroValue(des.XPos) {
		des.XPos = initial.XPos
	} else {
		cDes.XPos = des.XPos
	}
	if dcl.IsZeroValue(des.YPos) {
		des.YPos = initial.YPos
	} else {
		cDes.YPos = des.YPos
	}
	if dcl.IsZeroValue(des.Width) {
		des.Width = initial.Width
	} else {
		cDes.Width = des.Width
	}
	if dcl.IsZeroValue(des.Height) {
		des.Height = initial.Height
	} else {
		cDes.Height = des.Height
	}
	cDes.Widget = canonicalizeDashboardWidget(des.Widget, initial.Widget, opts...)

	return cDes
}

func canonicalizeDashboardMosaicLayoutTilesSlice(des, initial []DashboardMosaicLayoutTiles, opts ...dcl.ApplyOption) []DashboardMosaicLayoutTiles {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardMosaicLayoutTiles, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardMosaicLayoutTiles(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardMosaicLayoutTiles, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardMosaicLayoutTiles(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardMosaicLayoutTiles(c *Client, des, nw *DashboardMosaicLayoutTiles) *DashboardMosaicLayoutTiles {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardMosaicLayoutTiles while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Widget = canonicalizeNewDashboardWidget(c, des.Widget, nw.Widget)

	return nw
}

func canonicalizeNewDashboardMosaicLayoutTilesSet(c *Client, des, nw []DashboardMosaicLayoutTiles) []DashboardMosaicLayoutTiles {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardMosaicLayoutTiles
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardMosaicLayoutTilesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardMosaicLayoutTilesSlice(c *Client, des, nw []DashboardMosaicLayoutTiles) []DashboardMosaicLayoutTiles {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardMosaicLayoutTiles
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardMosaicLayoutTiles(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardRowLayout(des, initial *DashboardRowLayout, opts ...dcl.ApplyOption) *DashboardRowLayout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardRowLayout{}

	cDes.Rows = canonicalizeDashboardRowLayoutRowsSlice(des.Rows, initial.Rows, opts...)

	return cDes
}

func canonicalizeDashboardRowLayoutSlice(des, initial []DashboardRowLayout, opts ...dcl.ApplyOption) []DashboardRowLayout {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardRowLayout, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardRowLayout(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardRowLayout, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardRowLayout(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardRowLayout(c *Client, des, nw *DashboardRowLayout) *DashboardRowLayout {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardRowLayout while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Rows = canonicalizeNewDashboardRowLayoutRowsSlice(c, des.Rows, nw.Rows)

	return nw
}

func canonicalizeNewDashboardRowLayoutSet(c *Client, des, nw []DashboardRowLayout) []DashboardRowLayout {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardRowLayout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardRowLayoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardRowLayoutSlice(c *Client, des, nw []DashboardRowLayout) []DashboardRowLayout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardRowLayout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardRowLayout(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardRowLayoutRows(des, initial *DashboardRowLayoutRows, opts ...dcl.ApplyOption) *DashboardRowLayoutRows {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardRowLayoutRows{}

	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	} else {
		cDes.Weight = des.Weight
	}
	if dcl.IsZeroValue(des.Widgets) {
		des.Widgets = initial.Widgets
	} else {
		cDes.Widgets = des.Widgets
	}

	return cDes
}

func canonicalizeDashboardRowLayoutRowsSlice(des, initial []DashboardRowLayoutRows, opts ...dcl.ApplyOption) []DashboardRowLayoutRows {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardRowLayoutRows, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardRowLayoutRows(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardRowLayoutRows, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardRowLayoutRows(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardRowLayoutRows(c *Client, des, nw *DashboardRowLayoutRows) *DashboardRowLayoutRows {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardRowLayoutRows while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDashboardRowLayoutRowsSet(c *Client, des, nw []DashboardRowLayoutRows) []DashboardRowLayoutRows {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardRowLayoutRows
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardRowLayoutRowsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardRowLayoutRowsSlice(c *Client, des, nw []DashboardRowLayoutRows) []DashboardRowLayoutRows {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardRowLayoutRows
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardRowLayoutRows(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardColumnLayout(des, initial *DashboardColumnLayout, opts ...dcl.ApplyOption) *DashboardColumnLayout {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardColumnLayout{}

	cDes.Columns = canonicalizeDashboardColumnLayoutColumnsSlice(des.Columns, initial.Columns, opts...)

	return cDes
}

func canonicalizeDashboardColumnLayoutSlice(des, initial []DashboardColumnLayout, opts ...dcl.ApplyOption) []DashboardColumnLayout {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardColumnLayout, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardColumnLayout(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardColumnLayout, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardColumnLayout(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardColumnLayout(c *Client, des, nw *DashboardColumnLayout) *DashboardColumnLayout {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardColumnLayout while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Columns = canonicalizeNewDashboardColumnLayoutColumnsSlice(c, des.Columns, nw.Columns)

	return nw
}

func canonicalizeNewDashboardColumnLayoutSet(c *Client, des, nw []DashboardColumnLayout) []DashboardColumnLayout {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardColumnLayout
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardColumnLayoutNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardColumnLayoutSlice(c *Client, des, nw []DashboardColumnLayout) []DashboardColumnLayout {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardColumnLayout
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardColumnLayout(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardColumnLayoutColumns(des, initial *DashboardColumnLayoutColumns, opts ...dcl.ApplyOption) *DashboardColumnLayoutColumns {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardColumnLayoutColumns{}

	if dcl.IsZeroValue(des.Weight) {
		des.Weight = initial.Weight
	} else {
		cDes.Weight = des.Weight
	}
	cDes.Widgets = canonicalizeDashboardWidgetSlice(des.Widgets, initial.Widgets, opts...)

	return cDes
}

func canonicalizeDashboardColumnLayoutColumnsSlice(des, initial []DashboardColumnLayoutColumns, opts ...dcl.ApplyOption) []DashboardColumnLayoutColumns {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardColumnLayoutColumns, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardColumnLayoutColumns(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardColumnLayoutColumns, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardColumnLayoutColumns(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardColumnLayoutColumns(c *Client, des, nw *DashboardColumnLayoutColumns) *DashboardColumnLayoutColumns {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardColumnLayoutColumns while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Widgets = canonicalizeNewDashboardWidgetSlice(c, des.Widgets, nw.Widgets)

	return nw
}

func canonicalizeNewDashboardColumnLayoutColumnsSet(c *Client, des, nw []DashboardColumnLayoutColumns) []DashboardColumnLayoutColumns {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardColumnLayoutColumns
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardColumnLayoutColumnsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardColumnLayoutColumnsSlice(c *Client, des, nw []DashboardColumnLayoutColumns) []DashboardColumnLayoutColumns {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardColumnLayoutColumns
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardColumnLayoutColumns(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidget(des, initial *DashboardWidget, opts ...dcl.ApplyOption) *DashboardWidget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.XyChart != nil || (initial != nil && initial.XyChart != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Scorecard, des.Text, des.Blank) {
			des.XyChart = nil
			if initial != nil {
				initial.XyChart = nil
			}
		}
	}

	if des.Scorecard != nil || (initial != nil && initial.Scorecard != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.XyChart, des.Text, des.Blank) {
			des.Scorecard = nil
			if initial != nil {
				initial.Scorecard = nil
			}
		}
	}

	if des.Text != nil || (initial != nil && initial.Text != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.XyChart, des.Scorecard, des.Blank) {
			des.Text = nil
			if initial != nil {
				initial.Text = nil
			}
		}
	}

	if des.Blank != nil || (initial != nil && initial.Blank != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.XyChart, des.Scorecard, des.Text) {
			des.Blank = nil
			if initial != nil {
				initial.Blank = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidget{}

	if dcl.StringCanonicalize(des.Title, initial.Title) || dcl.IsZeroValue(des.Title) {
		cDes.Title = initial.Title
	} else {
		cDes.Title = des.Title
	}
	cDes.XyChart = canonicalizeDashboardWidgetXyChart(des.XyChart, initial.XyChart, opts...)
	cDes.Scorecard = canonicalizeDashboardWidgetScorecard(des.Scorecard, initial.Scorecard, opts...)
	cDes.Text = canonicalizeDashboardWidgetText(des.Text, initial.Text, opts...)
	cDes.Blank = canonicalizeDashboardWidgetBlank(des.Blank, initial.Blank, opts...)

	return cDes
}

func canonicalizeDashboardWidgetSlice(des, initial []DashboardWidget, opts ...dcl.ApplyOption) []DashboardWidget {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidget, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidget(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidget, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidget(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidget(c *Client, des, nw *DashboardWidget) *DashboardWidget {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidget while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Title, nw.Title) {
		nw.Title = des.Title
	}
	nw.XyChart = canonicalizeNewDashboardWidgetXyChart(c, des.XyChart, nw.XyChart)
	nw.Scorecard = canonicalizeNewDashboardWidgetScorecard(c, des.Scorecard, nw.Scorecard)
	nw.Text = canonicalizeNewDashboardWidgetText(c, des.Text, nw.Text)
	nw.Blank = canonicalizeNewDashboardWidgetBlank(c, des.Blank, nw.Blank)

	return nw
}

func canonicalizeNewDashboardWidgetSet(c *Client, des, nw []DashboardWidget) []DashboardWidget {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidget
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetSlice(c *Client, des, nw []DashboardWidget) []DashboardWidget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidget(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChart(des, initial *DashboardWidgetXyChart, opts ...dcl.ApplyOption) *DashboardWidgetXyChart {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChart{}

	cDes.DataSets = canonicalizeDashboardWidgetXyChartDataSetsSlice(des.DataSets, initial.DataSets, opts...)
	if dcl.StringCanonicalize(des.TimeshiftDuration, initial.TimeshiftDuration) || dcl.IsZeroValue(des.TimeshiftDuration) {
		cDes.TimeshiftDuration = initial.TimeshiftDuration
	} else {
		cDes.TimeshiftDuration = des.TimeshiftDuration
	}
	cDes.Thresholds = canonicalizeDashboardWidgetXyChartThresholdsSlice(des.Thresholds, initial.Thresholds, opts...)
	cDes.XAxis = canonicalizeDashboardWidgetXyChartXAxis(des.XAxis, initial.XAxis, opts...)
	cDes.YAxis = canonicalizeDashboardWidgetXyChartYAxis(des.YAxis, initial.YAxis, opts...)
	cDes.ChartOptions = canonicalizeDashboardWidgetXyChartChartOptions(des.ChartOptions, initial.ChartOptions, opts...)

	return cDes
}

func canonicalizeDashboardWidgetXyChartSlice(des, initial []DashboardWidgetXyChart, opts ...dcl.ApplyOption) []DashboardWidgetXyChart {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChart, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChart(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChart, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChart(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChart(c *Client, des, nw *DashboardWidgetXyChart) *DashboardWidgetXyChart {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChart while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.DataSets = canonicalizeNewDashboardWidgetXyChartDataSetsSlice(c, des.DataSets, nw.DataSets)
	if dcl.StringCanonicalize(des.TimeshiftDuration, nw.TimeshiftDuration) {
		nw.TimeshiftDuration = des.TimeshiftDuration
	}
	nw.Thresholds = canonicalizeNewDashboardWidgetXyChartThresholdsSlice(c, des.Thresholds, nw.Thresholds)
	nw.XAxis = canonicalizeNewDashboardWidgetXyChartXAxis(c, des.XAxis, nw.XAxis)
	nw.YAxis = canonicalizeNewDashboardWidgetXyChartYAxis(c, des.YAxis, nw.YAxis)
	nw.ChartOptions = canonicalizeNewDashboardWidgetXyChartChartOptions(c, des.ChartOptions, nw.ChartOptions)

	return nw
}

func canonicalizeNewDashboardWidgetXyChartSet(c *Client, des, nw []DashboardWidgetXyChart) []DashboardWidgetXyChart {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChart
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartSlice(c *Client, des, nw []DashboardWidgetXyChart) []DashboardWidgetXyChart {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChart
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChart(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSets(des, initial *DashboardWidgetXyChartDataSets, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSets{}

	cDes.TimeSeriesQuery = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQuery(des.TimeSeriesQuery, initial.TimeSeriesQuery, opts...)
	if dcl.IsZeroValue(des.PlotType) {
		des.PlotType = initial.PlotType
	} else {
		cDes.PlotType = des.PlotType
	}
	if dcl.StringCanonicalize(des.LegendTemplate, initial.LegendTemplate) || dcl.IsZeroValue(des.LegendTemplate) {
		cDes.LegendTemplate = initial.LegendTemplate
	} else {
		cDes.LegendTemplate = des.LegendTemplate
	}
	if dcl.StringCanonicalize(des.MinAlignmentPeriod, initial.MinAlignmentPeriod) || dcl.IsZeroValue(des.MinAlignmentPeriod) {
		cDes.MinAlignmentPeriod = initial.MinAlignmentPeriod
	} else {
		cDes.MinAlignmentPeriod = des.MinAlignmentPeriod
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsSlice(des, initial []DashboardWidgetXyChartDataSets, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSets {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSets, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSets(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSets, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSets(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSets(c *Client, des, nw *DashboardWidgetXyChartDataSets) *DashboardWidgetXyChartDataSets {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSets while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.TimeSeriesQuery = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQuery(c, des.TimeSeriesQuery, nw.TimeSeriesQuery)
	if dcl.StringCanonicalize(des.LegendTemplate, nw.LegendTemplate) {
		nw.LegendTemplate = des.LegendTemplate
	}
	if dcl.StringCanonicalize(des.MinAlignmentPeriod, nw.MinAlignmentPeriod) {
		nw.MinAlignmentPeriod = des.MinAlignmentPeriod
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsSet(c *Client, des, nw []DashboardWidgetXyChartDataSets) []DashboardWidgetXyChartDataSets {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsSlice(c *Client, des, nw []DashboardWidgetXyChartDataSets) []DashboardWidgetXyChartDataSets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSets(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQuery(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQuery, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQuery{}

	cDes.TimeSeriesFilter = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(des.TimeSeriesFilter, initial.TimeSeriesFilter, opts...)
	cDes.TimeSeriesFilterRatio = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(des.TimeSeriesFilterRatio, initial.TimeSeriesFilterRatio, opts...)
	if dcl.StringCanonicalize(des.TimeSeriesQueryLanguage, initial.TimeSeriesQueryLanguage) || dcl.IsZeroValue(des.TimeSeriesQueryLanguage) {
		cDes.TimeSeriesQueryLanguage = initial.TimeSeriesQueryLanguage
	} else {
		cDes.TimeSeriesQueryLanguage = des.TimeSeriesQueryLanguage
	}
	if dcl.StringCanonicalize(des.UnitOverride, initial.UnitOverride) || dcl.IsZeroValue(des.UnitOverride) {
		cDes.UnitOverride = initial.UnitOverride
	} else {
		cDes.UnitOverride = des.UnitOverride
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQuerySlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQuery, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQuery, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQuery(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQuery, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQuery(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQuery(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQuery) *DashboardWidgetXyChartDataSetsTimeSeriesQuery {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQuery while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.TimeSeriesFilter = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c, des.TimeSeriesFilter, nw.TimeSeriesFilter)
	nw.TimeSeriesFilterRatio = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c, des.TimeSeriesFilterRatio, nw.TimeSeriesFilterRatio)
	if dcl.StringCanonicalize(des.TimeSeriesQueryLanguage, nw.TimeSeriesQueryLanguage) {
		nw.TimeSeriesQueryLanguage = des.TimeSeriesQueryLanguage
	}
	if dcl.StringCanonicalize(des.UnitOverride, nw.UnitOverride) {
		nw.UnitOverride = des.UnitOverride
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQuerySet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQuery) []DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQuery
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQuerySlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQuery) []DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQuery
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQuery(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		cDes.Filter = initial.Filter
	} else {
		cDes.Filter = des.Filter
	}
	cDes.Aggregation = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(des.Aggregation, initial.Aggregation, opts...)
	cDes.SecondaryAggregation = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(des.SecondaryAggregation, initial.SecondaryAggregation, opts...)
	cDes.PickTimeSeriesFilter = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(des.PickTimeSeriesFilter, initial.PickTimeSeriesFilter, opts...)

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	nw.Aggregation = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c, des.Aggregation, nw.Aggregation)
	nw.SecondaryAggregation = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, des.SecondaryAggregation, nw.SecondaryAggregation)
	nw.PickTimeSeriesFilter = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, des.PickTimeSeriesFilter, nw.PickTimeSeriesFilter)

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}

	if dcl.IsZeroValue(des.RankingMethod) {
		des.RankingMethod = initial.RankingMethod
	} else {
		cDes.RankingMethod = des.RankingMethod
	}
	if dcl.IsZeroValue(des.NumTimeSeries) {
		des.NumTimeSeries = initial.NumTimeSeries
	} else {
		cDes.NumTimeSeries = des.NumTimeSeries
	}
	if dcl.IsZeroValue(des.Direction) {
		des.Direction = initial.Direction
	} else {
		cDes.Direction = des.Direction
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}

	cDes.Numerator = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(des.Numerator, initial.Numerator, opts...)
	cDes.Denominator = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(des.Denominator, initial.Denominator, opts...)
	cDes.SecondaryAggregation = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(des.SecondaryAggregation, initial.SecondaryAggregation, opts...)
	cDes.PickTimeSeriesFilter = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(des.PickTimeSeriesFilter, initial.PickTimeSeriesFilter, opts...)

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Numerator = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, des.Numerator, nw.Numerator)
	nw.Denominator = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, des.Denominator, nw.Denominator)
	nw.SecondaryAggregation = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, des.SecondaryAggregation, nw.SecondaryAggregation)
	nw.PickTimeSeriesFilter = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, des.PickTimeSeriesFilter, nw.PickTimeSeriesFilter)

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		cDes.Filter = initial.Filter
	} else {
		cDes.Filter = des.Filter
	}
	cDes.Aggregation = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(des.Aggregation, initial.Aggregation, opts...)

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	nw.Aggregation = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, des.Aggregation, nw.Aggregation)

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		cDes.Filter = initial.Filter
	} else {
		cDes.Filter = des.Filter
	}
	cDes.Aggregation = canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(des.Aggregation, initial.Aggregation, opts...)

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	nw.Aggregation = canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, des.Aggregation, nw.Aggregation)

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(des, initial *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, opts ...dcl.ApplyOption) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}

	if dcl.IsZeroValue(des.RankingMethod) {
		des.RankingMethod = initial.RankingMethod
	} else {
		cDes.RankingMethod = des.RankingMethod
	}
	if dcl.IsZeroValue(des.NumTimeSeries) {
		des.NumTimeSeries = initial.NumTimeSeries
	} else {
		cDes.NumTimeSeries = des.NumTimeSeries
	}
	if dcl.IsZeroValue(des.Direction) {
		des.Direction = initial.Direction
	} else {
		cDes.Direction = des.Direction
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice(des, initial []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, opts ...dcl.ApplyOption) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c *Client, des, nw *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSet(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice(c *Client, des, nw []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartThresholds(des, initial *DashboardWidgetXyChartThresholds, opts ...dcl.ApplyOption) *DashboardWidgetXyChartThresholds {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartThresholds{}

	if dcl.StringCanonicalize(des.Label, initial.Label) || dcl.IsZeroValue(des.Label) {
		cDes.Label = initial.Label
	} else {
		cDes.Label = des.Label
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}
	if dcl.IsZeroValue(des.Color) {
		des.Color = initial.Color
	} else {
		cDes.Color = des.Color
	}
	if dcl.IsZeroValue(des.Direction) {
		des.Direction = initial.Direction
	} else {
		cDes.Direction = des.Direction
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartThresholdsSlice(des, initial []DashboardWidgetXyChartThresholds, opts ...dcl.ApplyOption) []DashboardWidgetXyChartThresholds {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartThresholds, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartThresholds(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartThresholds, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartThresholds(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartThresholds(c *Client, des, nw *DashboardWidgetXyChartThresholds) *DashboardWidgetXyChartThresholds {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartThresholds while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Label, nw.Label) {
		nw.Label = des.Label
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartThresholdsSet(c *Client, des, nw []DashboardWidgetXyChartThresholds) []DashboardWidgetXyChartThresholds {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartThresholds
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartThresholdsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartThresholdsSlice(c *Client, des, nw []DashboardWidgetXyChartThresholds) []DashboardWidgetXyChartThresholds {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartThresholds
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartThresholds(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartXAxis(des, initial *DashboardWidgetXyChartXAxis, opts ...dcl.ApplyOption) *DashboardWidgetXyChartXAxis {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartXAxis{}

	if dcl.StringCanonicalize(des.Label, initial.Label) || dcl.IsZeroValue(des.Label) {
		cDes.Label = initial.Label
	} else {
		cDes.Label = des.Label
	}
	if dcl.IsZeroValue(des.Scale) {
		des.Scale = initial.Scale
	} else {
		cDes.Scale = des.Scale
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartXAxisSlice(des, initial []DashboardWidgetXyChartXAxis, opts ...dcl.ApplyOption) []DashboardWidgetXyChartXAxis {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartXAxis, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartXAxis(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartXAxis, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartXAxis(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartXAxis(c *Client, des, nw *DashboardWidgetXyChartXAxis) *DashboardWidgetXyChartXAxis {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartXAxis while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Label, nw.Label) {
		nw.Label = des.Label
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartXAxisSet(c *Client, des, nw []DashboardWidgetXyChartXAxis) []DashboardWidgetXyChartXAxis {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartXAxis
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartXAxisNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartXAxisSlice(c *Client, des, nw []DashboardWidgetXyChartXAxis) []DashboardWidgetXyChartXAxis {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartXAxis
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartXAxis(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartYAxis(des, initial *DashboardWidgetXyChartYAxis, opts ...dcl.ApplyOption) *DashboardWidgetXyChartYAxis {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartYAxis{}

	if dcl.StringCanonicalize(des.Label, initial.Label) || dcl.IsZeroValue(des.Label) {
		cDes.Label = initial.Label
	} else {
		cDes.Label = des.Label
	}
	if dcl.IsZeroValue(des.Scale) {
		des.Scale = initial.Scale
	} else {
		cDes.Scale = des.Scale
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartYAxisSlice(des, initial []DashboardWidgetXyChartYAxis, opts ...dcl.ApplyOption) []DashboardWidgetXyChartYAxis {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartYAxis, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartYAxis(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartYAxis, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartYAxis(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartYAxis(c *Client, des, nw *DashboardWidgetXyChartYAxis) *DashboardWidgetXyChartYAxis {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartYAxis while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Label, nw.Label) {
		nw.Label = des.Label
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartYAxisSet(c *Client, des, nw []DashboardWidgetXyChartYAxis) []DashboardWidgetXyChartYAxis {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartYAxis
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartYAxisNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartYAxisSlice(c *Client, des, nw []DashboardWidgetXyChartYAxis) []DashboardWidgetXyChartYAxis {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartYAxis
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartYAxis(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetXyChartChartOptions(des, initial *DashboardWidgetXyChartChartOptions, opts ...dcl.ApplyOption) *DashboardWidgetXyChartChartOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetXyChartChartOptions{}

	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	} else {
		cDes.Mode = des.Mode
	}

	return cDes
}

func canonicalizeDashboardWidgetXyChartChartOptionsSlice(des, initial []DashboardWidgetXyChartChartOptions, opts ...dcl.ApplyOption) []DashboardWidgetXyChartChartOptions {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetXyChartChartOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetXyChartChartOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetXyChartChartOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetXyChartChartOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetXyChartChartOptions(c *Client, des, nw *DashboardWidgetXyChartChartOptions) *DashboardWidgetXyChartChartOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetXyChartChartOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDashboardWidgetXyChartChartOptionsSet(c *Client, des, nw []DashboardWidgetXyChartChartOptions) []DashboardWidgetXyChartChartOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetXyChartChartOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetXyChartChartOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetXyChartChartOptionsSlice(c *Client, des, nw []DashboardWidgetXyChartChartOptions) []DashboardWidgetXyChartChartOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetXyChartChartOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetXyChartChartOptions(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecard(des, initial *DashboardWidgetScorecard, opts ...dcl.ApplyOption) *DashboardWidgetScorecard {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecard{}

	cDes.TimeSeriesQuery = canonicalizeDashboardWidgetScorecardTimeSeriesQuery(des.TimeSeriesQuery, initial.TimeSeriesQuery, opts...)
	cDes.GaugeView = canonicalizeDashboardWidgetScorecardGaugeView(des.GaugeView, initial.GaugeView, opts...)
	cDes.SparkChartView = canonicalizeDashboardWidgetScorecardSparkChartView(des.SparkChartView, initial.SparkChartView, opts...)
	cDes.Thresholds = canonicalizeDashboardWidgetScorecardThresholdsSlice(des.Thresholds, initial.Thresholds, opts...)

	return cDes
}

func canonicalizeDashboardWidgetScorecardSlice(des, initial []DashboardWidgetScorecard, opts ...dcl.ApplyOption) []DashboardWidgetScorecard {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecard, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecard(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecard, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecard(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecard(c *Client, des, nw *DashboardWidgetScorecard) *DashboardWidgetScorecard {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecard while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.TimeSeriesQuery = canonicalizeNewDashboardWidgetScorecardTimeSeriesQuery(c, des.TimeSeriesQuery, nw.TimeSeriesQuery)
	nw.GaugeView = canonicalizeNewDashboardWidgetScorecardGaugeView(c, des.GaugeView, nw.GaugeView)
	nw.SparkChartView = canonicalizeNewDashboardWidgetScorecardSparkChartView(c, des.SparkChartView, nw.SparkChartView)
	nw.Thresholds = canonicalizeNewDashboardWidgetScorecardThresholdsSlice(c, des.Thresholds, nw.Thresholds)

	return nw
}

func canonicalizeNewDashboardWidgetScorecardSet(c *Client, des, nw []DashboardWidgetScorecard) []DashboardWidgetScorecard {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecard
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardSlice(c *Client, des, nw []DashboardWidgetScorecard) []DashboardWidgetScorecard {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecard
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecard(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQuery(des, initial *DashboardWidgetScorecardTimeSeriesQuery, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQuery {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQuery{}

	cDes.TimeSeriesFilter = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(des.TimeSeriesFilter, initial.TimeSeriesFilter, opts...)
	cDes.TimeSeriesFilterRatio = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(des.TimeSeriesFilterRatio, initial.TimeSeriesFilterRatio, opts...)
	if dcl.StringCanonicalize(des.TimeSeriesQueryLanguage, initial.TimeSeriesQueryLanguage) || dcl.IsZeroValue(des.TimeSeriesQueryLanguage) {
		cDes.TimeSeriesQueryLanguage = initial.TimeSeriesQueryLanguage
	} else {
		cDes.TimeSeriesQueryLanguage = des.TimeSeriesQueryLanguage
	}
	if dcl.StringCanonicalize(des.UnitOverride, initial.UnitOverride) || dcl.IsZeroValue(des.UnitOverride) {
		cDes.UnitOverride = initial.UnitOverride
	} else {
		cDes.UnitOverride = des.UnitOverride
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQuerySlice(des, initial []DashboardWidgetScorecardTimeSeriesQuery, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQuery {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQuery, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQuery(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQuery, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQuery(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQuery(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQuery) *DashboardWidgetScorecardTimeSeriesQuery {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQuery while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.TimeSeriesFilter = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c, des.TimeSeriesFilter, nw.TimeSeriesFilter)
	nw.TimeSeriesFilterRatio = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c, des.TimeSeriesFilterRatio, nw.TimeSeriesFilterRatio)
	if dcl.StringCanonicalize(des.TimeSeriesQueryLanguage, nw.TimeSeriesQueryLanguage) {
		nw.TimeSeriesQueryLanguage = des.TimeSeriesQueryLanguage
	}
	if dcl.StringCanonicalize(des.UnitOverride, nw.UnitOverride) {
		nw.UnitOverride = des.UnitOverride
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQuerySet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQuery) []DashboardWidgetScorecardTimeSeriesQuery {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQuery
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQuerySlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQuery) []DashboardWidgetScorecardTimeSeriesQuery {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQuery
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQuery(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		cDes.Filter = initial.Filter
	} else {
		cDes.Filter = des.Filter
	}
	cDes.Aggregation = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(des.Aggregation, initial.Aggregation, opts...)
	cDes.SecondaryAggregation = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(des.SecondaryAggregation, initial.SecondaryAggregation, opts...)
	cDes.PickTimeSeriesFilter = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(des.PickTimeSeriesFilter, initial.PickTimeSeriesFilter, opts...)

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	nw.Aggregation = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c, des.Aggregation, nw.Aggregation)
	nw.SecondaryAggregation = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, des.SecondaryAggregation, nw.SecondaryAggregation)
	nw.PickTimeSeriesFilter = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, des.PickTimeSeriesFilter, nw.PickTimeSeriesFilter)

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}

	if dcl.IsZeroValue(des.RankingMethod) {
		des.RankingMethod = initial.RankingMethod
	} else {
		cDes.RankingMethod = des.RankingMethod
	}
	if dcl.IsZeroValue(des.NumTimeSeries) {
		des.NumTimeSeries = initial.NumTimeSeries
	} else {
		cDes.NumTimeSeries = des.NumTimeSeries
	}
	if dcl.IsZeroValue(des.Direction) {
		des.Direction = initial.Direction
	} else {
		cDes.Direction = des.Direction
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}

	cDes.Numerator = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(des.Numerator, initial.Numerator, opts...)
	cDes.Denominator = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(des.Denominator, initial.Denominator, opts...)
	cDes.SecondaryAggregation = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(des.SecondaryAggregation, initial.SecondaryAggregation, opts...)
	cDes.PickTimeSeriesFilter = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(des.PickTimeSeriesFilter, initial.PickTimeSeriesFilter, opts...)

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Numerator = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, des.Numerator, nw.Numerator)
	nw.Denominator = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, des.Denominator, nw.Denominator)
	nw.SecondaryAggregation = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, des.SecondaryAggregation, nw.SecondaryAggregation)
	nw.PickTimeSeriesFilter = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, des.PickTimeSeriesFilter, nw.PickTimeSeriesFilter)

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		cDes.Filter = initial.Filter
	} else {
		cDes.Filter = des.Filter
	}
	cDes.Aggregation = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(des.Aggregation, initial.Aggregation, opts...)

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	nw.Aggregation = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, des.Aggregation, nw.Aggregation)

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}

	if dcl.StringCanonicalize(des.Filter, initial.Filter) || dcl.IsZeroValue(des.Filter) {
		cDes.Filter = initial.Filter
	} else {
		cDes.Filter = des.Filter
	}
	cDes.Aggregation = canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(des.Aggregation, initial.Aggregation, opts...)

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Filter, nw.Filter) {
		nw.Filter = des.Filter
	}
	nw.Aggregation = canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, des.Aggregation, nw.Aggregation)

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}

	if dcl.StringCanonicalize(des.AlignmentPeriod, initial.AlignmentPeriod) || dcl.IsZeroValue(des.AlignmentPeriod) {
		cDes.AlignmentPeriod = initial.AlignmentPeriod
	} else {
		cDes.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.IsZeroValue(des.PerSeriesAligner) {
		des.PerSeriesAligner = initial.PerSeriesAligner
	} else {
		cDes.PerSeriesAligner = des.PerSeriesAligner
	}
	if dcl.IsZeroValue(des.CrossSeriesReducer) {
		des.CrossSeriesReducer = initial.CrossSeriesReducer
	} else {
		cDes.CrossSeriesReducer = des.CrossSeriesReducer
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, initial.GroupByFields) || dcl.IsZeroValue(des.GroupByFields) {
		cDes.GroupByFields = initial.GroupByFields
	} else {
		cDes.GroupByFields = des.GroupByFields
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AlignmentPeriod, nw.AlignmentPeriod) {
		nw.AlignmentPeriod = des.AlignmentPeriod
	}
	if dcl.StringArrayCanonicalize(des.GroupByFields, nw.GroupByFields) {
		nw.GroupByFields = des.GroupByFields
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(des, initial *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, opts ...dcl.ApplyOption) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}

	if dcl.IsZeroValue(des.RankingMethod) {
		des.RankingMethod = initial.RankingMethod
	} else {
		cDes.RankingMethod = des.RankingMethod
	}
	if dcl.IsZeroValue(des.NumTimeSeries) {
		des.NumTimeSeries = initial.NumTimeSeries
	} else {
		cDes.NumTimeSeries = des.NumTimeSeries
	}
	if dcl.IsZeroValue(des.Direction) {
		des.Direction = initial.Direction
	} else {
		cDes.Direction = des.Direction
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice(des, initial []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, opts ...dcl.ApplyOption) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c *Client, des, nw *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSet(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice(c *Client, des, nw []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardGaugeView(des, initial *DashboardWidgetScorecardGaugeView, opts ...dcl.ApplyOption) *DashboardWidgetScorecardGaugeView {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardGaugeView{}

	if dcl.IsZeroValue(des.LowerBound) {
		des.LowerBound = initial.LowerBound
	} else {
		cDes.LowerBound = des.LowerBound
	}
	if dcl.IsZeroValue(des.UpperBound) {
		des.UpperBound = initial.UpperBound
	} else {
		cDes.UpperBound = des.UpperBound
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardGaugeViewSlice(des, initial []DashboardWidgetScorecardGaugeView, opts ...dcl.ApplyOption) []DashboardWidgetScorecardGaugeView {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardGaugeView, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardGaugeView(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardGaugeView, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardGaugeView(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardGaugeView(c *Client, des, nw *DashboardWidgetScorecardGaugeView) *DashboardWidgetScorecardGaugeView {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardGaugeView while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardGaugeViewSet(c *Client, des, nw []DashboardWidgetScorecardGaugeView) []DashboardWidgetScorecardGaugeView {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardGaugeView
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardGaugeViewNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardGaugeViewSlice(c *Client, des, nw []DashboardWidgetScorecardGaugeView) []DashboardWidgetScorecardGaugeView {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardGaugeView
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardGaugeView(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardSparkChartView(des, initial *DashboardWidgetScorecardSparkChartView, opts ...dcl.ApplyOption) *DashboardWidgetScorecardSparkChartView {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardSparkChartView{}

	if dcl.IsZeroValue(des.SparkChartType) {
		des.SparkChartType = initial.SparkChartType
	} else {
		cDes.SparkChartType = des.SparkChartType
	}
	if dcl.StringCanonicalize(des.MinAlignmentPeriod, initial.MinAlignmentPeriod) || dcl.IsZeroValue(des.MinAlignmentPeriod) {
		cDes.MinAlignmentPeriod = initial.MinAlignmentPeriod
	} else {
		cDes.MinAlignmentPeriod = des.MinAlignmentPeriod
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardSparkChartViewSlice(des, initial []DashboardWidgetScorecardSparkChartView, opts ...dcl.ApplyOption) []DashboardWidgetScorecardSparkChartView {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardSparkChartView, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardSparkChartView(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardSparkChartView, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardSparkChartView(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardSparkChartView(c *Client, des, nw *DashboardWidgetScorecardSparkChartView) *DashboardWidgetScorecardSparkChartView {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardSparkChartView while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MinAlignmentPeriod, nw.MinAlignmentPeriod) {
		nw.MinAlignmentPeriod = des.MinAlignmentPeriod
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardSparkChartViewSet(c *Client, des, nw []DashboardWidgetScorecardSparkChartView) []DashboardWidgetScorecardSparkChartView {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardSparkChartView
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardSparkChartViewNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardSparkChartViewSlice(c *Client, des, nw []DashboardWidgetScorecardSparkChartView) []DashboardWidgetScorecardSparkChartView {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardSparkChartView
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardSparkChartView(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetScorecardThresholds(des, initial *DashboardWidgetScorecardThresholds, opts ...dcl.ApplyOption) *DashboardWidgetScorecardThresholds {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetScorecardThresholds{}

	if dcl.StringCanonicalize(des.Label, initial.Label) || dcl.IsZeroValue(des.Label) {
		cDes.Label = initial.Label
	} else {
		cDes.Label = des.Label
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}
	if dcl.IsZeroValue(des.Color) {
		des.Color = initial.Color
	} else {
		cDes.Color = des.Color
	}
	if dcl.IsZeroValue(des.Direction) {
		des.Direction = initial.Direction
	} else {
		cDes.Direction = des.Direction
	}

	return cDes
}

func canonicalizeDashboardWidgetScorecardThresholdsSlice(des, initial []DashboardWidgetScorecardThresholds, opts ...dcl.ApplyOption) []DashboardWidgetScorecardThresholds {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetScorecardThresholds, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetScorecardThresholds(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetScorecardThresholds, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetScorecardThresholds(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetScorecardThresholds(c *Client, des, nw *DashboardWidgetScorecardThresholds) *DashboardWidgetScorecardThresholds {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetScorecardThresholds while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Label, nw.Label) {
		nw.Label = des.Label
	}

	return nw
}

func canonicalizeNewDashboardWidgetScorecardThresholdsSet(c *Client, des, nw []DashboardWidgetScorecardThresholds) []DashboardWidgetScorecardThresholds {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetScorecardThresholds
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetScorecardThresholdsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetScorecardThresholdsSlice(c *Client, des, nw []DashboardWidgetScorecardThresholds) []DashboardWidgetScorecardThresholds {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetScorecardThresholds
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetScorecardThresholds(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetText(des, initial *DashboardWidgetText, opts ...dcl.ApplyOption) *DashboardWidgetText {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetText{}

	if dcl.StringCanonicalize(des.Content, initial.Content) || dcl.IsZeroValue(des.Content) {
		cDes.Content = initial.Content
	} else {
		cDes.Content = des.Content
	}
	if dcl.IsZeroValue(des.Format) {
		des.Format = initial.Format
	} else {
		cDes.Format = des.Format
	}

	return cDes
}

func canonicalizeDashboardWidgetTextSlice(des, initial []DashboardWidgetText, opts ...dcl.ApplyOption) []DashboardWidgetText {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetText, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetText(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetText, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetText(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetText(c *Client, des, nw *DashboardWidgetText) *DashboardWidgetText {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetText while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}

	return nw
}

func canonicalizeNewDashboardWidgetTextSet(c *Client, des, nw []DashboardWidgetText) []DashboardWidgetText {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetText
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetTextNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetTextSlice(c *Client, des, nw []DashboardWidgetText) []DashboardWidgetText {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetText
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetText(c, &d, &n))
	}

	return items
}

func canonicalizeDashboardWidgetBlank(des, initial *DashboardWidgetBlank, opts ...dcl.ApplyOption) *DashboardWidgetBlank {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &DashboardWidgetBlank{}

	return cDes
}

func canonicalizeDashboardWidgetBlankSlice(des, initial []DashboardWidgetBlank, opts ...dcl.ApplyOption) []DashboardWidgetBlank {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DashboardWidgetBlank, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDashboardWidgetBlank(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DashboardWidgetBlank, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDashboardWidgetBlank(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDashboardWidgetBlank(c *Client, des, nw *DashboardWidgetBlank) *DashboardWidgetBlank {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for DashboardWidgetBlank while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDashboardWidgetBlankSet(c *Client, des, nw []DashboardWidgetBlank) []DashboardWidgetBlank {
	if des == nil {
		return nw
	}
	var reorderedNew []DashboardWidgetBlank
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDashboardWidgetBlankNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewDashboardWidgetBlankSlice(c *Client, des, nw []DashboardWidgetBlank) []DashboardWidgetBlank {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DashboardWidgetBlank
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDashboardWidgetBlank(c, &d, &n))
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
func diffDashboard(c *Client, desired, actual *Dashboard, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GridLayout, actual.GridLayout, dcl.Info{ObjectFunction: compareDashboardGridLayoutNewStyle, EmptyObject: EmptyDashboardGridLayout, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GridLayout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MosaicLayout, actual.MosaicLayout, dcl.Info{ObjectFunction: compareDashboardMosaicLayoutNewStyle, EmptyObject: EmptyDashboardMosaicLayout, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("MosaicLayout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RowLayout, actual.RowLayout, dcl.Info{ObjectFunction: compareDashboardRowLayoutNewStyle, EmptyObject: EmptyDashboardRowLayout, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("RowLayout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ColumnLayout, actual.ColumnLayout, dcl.Info{ObjectFunction: compareDashboardColumnLayoutNewStyle, EmptyObject: EmptyDashboardColumnLayout, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("ColumnLayout")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareDashboardGridLayoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardGridLayout)
	if !ok {
		desiredNotPointer, ok := d.(DashboardGridLayout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardGridLayout or *DashboardGridLayout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardGridLayout)
	if !ok {
		actualNotPointer, ok := a.(DashboardGridLayout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardGridLayout", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Columns, actual.Columns, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Columns")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Widgets, actual.Widgets, dcl.Info{ObjectFunction: compareDashboardWidgetNewStyle, EmptyObject: EmptyDashboardWidget, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Widgets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardMosaicLayoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardMosaicLayout)
	if !ok {
		desiredNotPointer, ok := d.(DashboardMosaicLayout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardMosaicLayout or *DashboardMosaicLayout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardMosaicLayout)
	if !ok {
		actualNotPointer, ok := a.(DashboardMosaicLayout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardMosaicLayout", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Columns, actual.Columns, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Columns")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tiles, actual.Tiles, dcl.Info{ObjectFunction: compareDashboardMosaicLayoutTilesNewStyle, EmptyObject: EmptyDashboardMosaicLayoutTiles, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Tiles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardMosaicLayoutTilesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardMosaicLayoutTiles)
	if !ok {
		desiredNotPointer, ok := d.(DashboardMosaicLayoutTiles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardMosaicLayoutTiles or *DashboardMosaicLayoutTiles", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardMosaicLayoutTiles)
	if !ok {
		actualNotPointer, ok := a.(DashboardMosaicLayoutTiles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardMosaicLayoutTiles", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.XPos, actual.XPos, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("XPos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.YPos, actual.YPos, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("YPos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Width, actual.Width, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Width")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Height, actual.Height, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Height")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Widget, actual.Widget, dcl.Info{ObjectFunction: compareDashboardWidgetNewStyle, EmptyObject: EmptyDashboardWidget, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Widget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardRowLayoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardRowLayout)
	if !ok {
		desiredNotPointer, ok := d.(DashboardRowLayout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardRowLayout or *DashboardRowLayout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardRowLayout)
	if !ok {
		actualNotPointer, ok := a.(DashboardRowLayout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardRowLayout", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Rows, actual.Rows, dcl.Info{ObjectFunction: compareDashboardRowLayoutRowsNewStyle, EmptyObject: EmptyDashboardRowLayoutRows, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Rows")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardRowLayoutRowsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardRowLayoutRows)
	if !ok {
		desiredNotPointer, ok := d.(DashboardRowLayoutRows)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardRowLayoutRows or *DashboardRowLayoutRows", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardRowLayoutRows)
	if !ok {
		actualNotPointer, ok := a.(DashboardRowLayoutRows)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardRowLayoutRows", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Widgets, actual.Widgets, dcl.Info{ObjectFunction: compareDashboardWidgetNewStyle, EmptyObject: EmptyDashboardWidget, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Widgets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardColumnLayoutNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardColumnLayout)
	if !ok {
		desiredNotPointer, ok := d.(DashboardColumnLayout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardColumnLayout or *DashboardColumnLayout", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardColumnLayout)
	if !ok {
		actualNotPointer, ok := a.(DashboardColumnLayout)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardColumnLayout", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Columns, actual.Columns, dcl.Info{ObjectFunction: compareDashboardColumnLayoutColumnsNewStyle, EmptyObject: EmptyDashboardColumnLayoutColumns, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Columns")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardColumnLayoutColumnsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardColumnLayoutColumns)
	if !ok {
		desiredNotPointer, ok := d.(DashboardColumnLayoutColumns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardColumnLayoutColumns or *DashboardColumnLayoutColumns", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardColumnLayoutColumns)
	if !ok {
		actualNotPointer, ok := a.(DashboardColumnLayoutColumns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardColumnLayoutColumns", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Widgets, actual.Widgets, dcl.Info{ObjectFunction: compareDashboardWidgetNewStyle, EmptyObject: EmptyDashboardWidget, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Widgets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidget)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidget or *DashboardWidget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidget)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Title, actual.Title, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Title")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.XyChart, actual.XyChart, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartNewStyle, EmptyObject: EmptyDashboardWidgetXyChart, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("XyChart")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scorecard, actual.Scorecard, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardNewStyle, EmptyObject: EmptyDashboardWidgetScorecard, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Scorecard")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Text, actual.Text, dcl.Info{ObjectFunction: compareDashboardWidgetTextNewStyle, EmptyObject: EmptyDashboardWidgetText, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Text")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Blank, actual.Blank, dcl.Info{ObjectFunction: compareDashboardWidgetBlankNewStyle, EmptyObject: EmptyDashboardWidgetBlank, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Blank")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChart)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChart)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChart or *DashboardWidgetXyChart", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChart)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChart)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChart", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DataSets, actual.DataSets, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSets, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("DataSets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeshiftDuration, actual.TimeshiftDuration, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("TimeshiftDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Thresholds, actual.Thresholds, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartThresholdsNewStyle, EmptyObject: EmptyDashboardWidgetXyChartThresholds, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Thresholds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.XAxis, actual.XAxis, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartXAxisNewStyle, EmptyObject: EmptyDashboardWidgetXyChartXAxis, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("XAxis")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.YAxis, actual.YAxis, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartYAxisNewStyle, EmptyObject: EmptyDashboardWidgetXyChartYAxis, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("YAxis")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ChartOptions, actual.ChartOptions, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartChartOptionsNewStyle, EmptyObject: EmptyDashboardWidgetXyChartChartOptions, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("ChartOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSets)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSets or *DashboardWidgetXyChartDataSets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSets)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TimeSeriesQuery, actual.TimeSeriesQuery, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQuery, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("TimeSeriesQuery")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PlotType, actual.PlotType, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PlotType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LegendTemplate, actual.LegendTemplate, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("LegendTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinAlignmentPeriod, actual.MinAlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("MinAlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQuery)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQuery)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQuery or *DashboardWidgetXyChartDataSetsTimeSeriesQuery", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQuery)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQuery)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQuery", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TimeSeriesFilter, actual.TimeSeriesFilter, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("TimeSeriesFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeSeriesFilterRatio, actual.TimeSeriesFilterRatio, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("TimeSeriesFilterRatio")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeSeriesQueryLanguage, actual.TimeSeriesQueryLanguage, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("TimeSeriesQueryLanguage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UnitOverride, actual.UnitOverride, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("UnitOverride")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aggregation, actual.Aggregation, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Aggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecondaryAggregation, actual.SecondaryAggregation, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("SecondaryAggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PickTimeSeriesFilter, actual.PickTimeSeriesFilter, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PickTimeSeriesFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RankingMethod, actual.RankingMethod, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("RankingMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumTimeSeries, actual.NumTimeSeries, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("NumTimeSeries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Direction, actual.Direction, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Direction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Numerator, actual.Numerator, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Numerator")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Denominator, actual.Denominator, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Denominator")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecondaryAggregation, actual.SecondaryAggregation, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("SecondaryAggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PickTimeSeriesFilter, actual.PickTimeSeriesFilter, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PickTimeSeriesFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aggregation, actual.Aggregation, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Aggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aggregation, actual.Aggregation, dcl.Info{ObjectFunction: compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationNewStyle, EmptyObject: EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Aggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter or *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RankingMethod, actual.RankingMethod, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("RankingMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumTimeSeries, actual.NumTimeSeries, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("NumTimeSeries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Direction, actual.Direction, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Direction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartThresholdsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartThresholds)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartThresholds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartThresholds or *DashboardWidgetXyChartThresholds", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartThresholds)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartThresholds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartThresholds", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Label, actual.Label, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Label")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Color, actual.Color, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Color")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Direction, actual.Direction, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Direction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartXAxisNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartXAxis)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartXAxis)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartXAxis or *DashboardWidgetXyChartXAxis", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartXAxis)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartXAxis)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartXAxis", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Label, actual.Label, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Label")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scale, actual.Scale, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Scale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartYAxisNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartYAxis)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartYAxis)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartYAxis or *DashboardWidgetXyChartYAxis", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartYAxis)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartYAxis)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartYAxis", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Label, actual.Label, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Label")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scale, actual.Scale, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Scale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetXyChartChartOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetXyChartChartOptions)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetXyChartChartOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartChartOptions or *DashboardWidgetXyChartChartOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetXyChartChartOptions)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetXyChartChartOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetXyChartChartOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecard)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecard)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecard or *DashboardWidgetScorecard", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecard)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecard)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecard", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TimeSeriesQuery, actual.TimeSeriesQuery, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQuery, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("TimeSeriesQuery")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GaugeView, actual.GaugeView, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardGaugeViewNewStyle, EmptyObject: EmptyDashboardWidgetScorecardGaugeView, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GaugeView")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SparkChartView, actual.SparkChartView, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardSparkChartViewNewStyle, EmptyObject: EmptyDashboardWidgetScorecardSparkChartView, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("SparkChartView")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Thresholds, actual.Thresholds, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardThresholdsNewStyle, EmptyObject: EmptyDashboardWidgetScorecardThresholds, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Thresholds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQuery)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQuery)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQuery or *DashboardWidgetScorecardTimeSeriesQuery", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQuery)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQuery)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQuery", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TimeSeriesFilter, actual.TimeSeriesFilter, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("TimeSeriesFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeSeriesFilterRatio, actual.TimeSeriesFilterRatio, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("TimeSeriesFilterRatio")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeSeriesQueryLanguage, actual.TimeSeriesQueryLanguage, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("TimeSeriesQueryLanguage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UnitOverride, actual.UnitOverride, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("UnitOverride")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aggregation, actual.Aggregation, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Aggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecondaryAggregation, actual.SecondaryAggregation, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("SecondaryAggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PickTimeSeriesFilter, actual.PickTimeSeriesFilter, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PickTimeSeriesFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RankingMethod, actual.RankingMethod, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("RankingMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumTimeSeries, actual.NumTimeSeries, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("NumTimeSeries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Direction, actual.Direction, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Direction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Numerator, actual.Numerator, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Numerator")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Denominator, actual.Denominator, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Denominator")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecondaryAggregation, actual.SecondaryAggregation, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("SecondaryAggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PickTimeSeriesFilter, actual.PickTimeSeriesFilter, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PickTimeSeriesFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aggregation, actual.Aggregation, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Aggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aggregation, actual.Aggregation, dcl.Info{ObjectFunction: compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationNewStyle, EmptyObject: EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Aggregation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AlignmentPeriod, actual.AlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("AlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerSeriesAligner, actual.PerSeriesAligner, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("PerSeriesAligner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrossSeriesReducer, actual.CrossSeriesReducer, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("CrossSeriesReducer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByFields, actual.GroupByFields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("GroupByFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter or *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RankingMethod, actual.RankingMethod, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("RankingMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumTimeSeries, actual.NumTimeSeries, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("NumTimeSeries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Direction, actual.Direction, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Direction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardGaugeViewNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardGaugeView)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardGaugeView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardGaugeView or *DashboardWidgetScorecardGaugeView", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardGaugeView)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardGaugeView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardGaugeView", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LowerBound, actual.LowerBound, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("LowerBound")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpperBound, actual.UpperBound, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("UpperBound")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardSparkChartViewNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardSparkChartView)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardSparkChartView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardSparkChartView or *DashboardWidgetScorecardSparkChartView", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardSparkChartView)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardSparkChartView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardSparkChartView", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SparkChartType, actual.SparkChartType, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("SparkChartType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinAlignmentPeriod, actual.MinAlignmentPeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("MinAlignmentPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetScorecardThresholdsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetScorecardThresholds)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetScorecardThresholds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardThresholds or *DashboardWidgetScorecardThresholds", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetScorecardThresholds)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetScorecardThresholds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetScorecardThresholds", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Label, actual.Label, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Label")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Color, actual.Color, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Color")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Direction, actual.Direction, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Direction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetTextNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DashboardWidgetText)
	if !ok {
		desiredNotPointer, ok := d.(DashboardWidgetText)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetText or *DashboardWidgetText", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DashboardWidgetText)
	if !ok {
		actualNotPointer, ok := a.(DashboardWidgetText)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DashboardWidgetText", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Format, actual.Format, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDashboardUpdateDashboardOperation")}, fn.AddNest("Format")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDashboardWidgetBlankNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Dashboard) urlNormalized() *Dashboard {
	normalized := dcl.Copy(*r).(Dashboard)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	return &normalized
}

func (r *Dashboard) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateDashboard" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/dashboards/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Dashboard resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Dashboard) marshal(c *Client) ([]byte, error) {
	m, err := expandDashboard(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Dashboard: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalDashboard decodes JSON responses into the Dashboard resource schema.
func unmarshalDashboard(b []byte, c *Client) (*Dashboard, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapDashboard(m, c)
}

func unmarshalMapDashboard(m map[string]interface{}, c *Client) (*Dashboard, error) {

	flattened := flattenDashboard(c, m)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandDashboard expands Dashboard into a JSON request object.
func expandDashboard(c *Client, f *Dashboard) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v, err := expandDashboardGridLayout(c, f.GridLayout); err != nil {
		return nil, fmt.Errorf("error expanding GridLayout into gridLayout: %w", err)
	} else if v != nil {
		m["gridLayout"] = v
	}
	if v, err := expandDashboardMosaicLayout(c, f.MosaicLayout); err != nil {
		return nil, fmt.Errorf("error expanding MosaicLayout into mosaicLayout: %w", err)
	} else if v != nil {
		m["mosaicLayout"] = v
	}
	if v, err := expandDashboardRowLayout(c, f.RowLayout); err != nil {
		return nil, fmt.Errorf("error expanding RowLayout into rowLayout: %w", err)
	} else if v != nil {
		m["rowLayout"] = v
	}
	if v, err := expandDashboardColumnLayout(c, f.ColumnLayout); err != nil {
		return nil, fmt.Errorf("error expanding ColumnLayout into columnLayout: %w", err)
	} else if v != nil {
		m["columnLayout"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenDashboard flattens Dashboard from a JSON request object into the
// Dashboard type.
func flattenDashboard(c *Client, i interface{}) *Dashboard {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Dashboard{}
	res.Name = dcl.FlattenString(m["name"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.GridLayout = flattenDashboardGridLayout(c, m["gridLayout"])
	res.MosaicLayout = flattenDashboardMosaicLayout(c, m["mosaicLayout"])
	res.RowLayout = flattenDashboardRowLayout(c, m["rowLayout"])
	res.ColumnLayout = flattenDashboardColumnLayout(c, m["columnLayout"])
	res.Project = dcl.FlattenString(m["project"])
	res.Etag = dcl.FlattenString(m["etag"])

	return res
}

// expandDashboardGridLayoutMap expands the contents of DashboardGridLayout into a JSON
// request object.
func expandDashboardGridLayoutMap(c *Client, f map[string]DashboardGridLayout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardGridLayout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardGridLayoutSlice expands the contents of DashboardGridLayout into a JSON
// request object.
func expandDashboardGridLayoutSlice(c *Client, f []DashboardGridLayout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardGridLayout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardGridLayoutMap flattens the contents of DashboardGridLayout from a JSON
// response object.
func flattenDashboardGridLayoutMap(c *Client, i interface{}) map[string]DashboardGridLayout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardGridLayout{}
	}

	if len(a) == 0 {
		return map[string]DashboardGridLayout{}
	}

	items := make(map[string]DashboardGridLayout)
	for k, item := range a {
		items[k] = *flattenDashboardGridLayout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardGridLayoutSlice flattens the contents of DashboardGridLayout from a JSON
// response object.
func flattenDashboardGridLayoutSlice(c *Client, i interface{}) []DashboardGridLayout {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardGridLayout{}
	}

	if len(a) == 0 {
		return []DashboardGridLayout{}
	}

	items := make([]DashboardGridLayout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardGridLayout(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardGridLayout expands an instance of DashboardGridLayout into a JSON
// request object.
func expandDashboardGridLayout(c *Client, f *DashboardGridLayout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Columns; !dcl.IsEmptyValueIndirect(v) {
		m["columns"] = v
	}
	if v := f.Widgets; v != nil {
		m["widgets"] = v
	}

	return m, nil
}

// flattenDashboardGridLayout flattens an instance of DashboardGridLayout from a JSON
// response object.
func flattenDashboardGridLayout(c *Client, i interface{}) *DashboardGridLayout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardGridLayout{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardGridLayout
	}
	r.Columns = dcl.FlattenInteger(m["columns"])
	r.Widgets = flattenDashboardWidgetSlice(c, m["widgets"])

	return r
}

// expandDashboardMosaicLayoutMap expands the contents of DashboardMosaicLayout into a JSON
// request object.
func expandDashboardMosaicLayoutMap(c *Client, f map[string]DashboardMosaicLayout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardMosaicLayout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardMosaicLayoutSlice expands the contents of DashboardMosaicLayout into a JSON
// request object.
func expandDashboardMosaicLayoutSlice(c *Client, f []DashboardMosaicLayout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardMosaicLayout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardMosaicLayoutMap flattens the contents of DashboardMosaicLayout from a JSON
// response object.
func flattenDashboardMosaicLayoutMap(c *Client, i interface{}) map[string]DashboardMosaicLayout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardMosaicLayout{}
	}

	if len(a) == 0 {
		return map[string]DashboardMosaicLayout{}
	}

	items := make(map[string]DashboardMosaicLayout)
	for k, item := range a {
		items[k] = *flattenDashboardMosaicLayout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardMosaicLayoutSlice flattens the contents of DashboardMosaicLayout from a JSON
// response object.
func flattenDashboardMosaicLayoutSlice(c *Client, i interface{}) []DashboardMosaicLayout {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardMosaicLayout{}
	}

	if len(a) == 0 {
		return []DashboardMosaicLayout{}
	}

	items := make([]DashboardMosaicLayout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardMosaicLayout(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardMosaicLayout expands an instance of DashboardMosaicLayout into a JSON
// request object.
func expandDashboardMosaicLayout(c *Client, f *DashboardMosaicLayout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Columns; !dcl.IsEmptyValueIndirect(v) {
		m["columns"] = v
	}
	if v, err := expandDashboardMosaicLayoutTilesSlice(c, f.Tiles); err != nil {
		return nil, fmt.Errorf("error expanding Tiles into tiles: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tiles"] = v
	}

	return m, nil
}

// flattenDashboardMosaicLayout flattens an instance of DashboardMosaicLayout from a JSON
// response object.
func flattenDashboardMosaicLayout(c *Client, i interface{}) *DashboardMosaicLayout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardMosaicLayout{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardMosaicLayout
	}
	r.Columns = dcl.FlattenInteger(m["columns"])
	r.Tiles = flattenDashboardMosaicLayoutTilesSlice(c, m["tiles"])

	return r
}

// expandDashboardMosaicLayoutTilesMap expands the contents of DashboardMosaicLayoutTiles into a JSON
// request object.
func expandDashboardMosaicLayoutTilesMap(c *Client, f map[string]DashboardMosaicLayoutTiles) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardMosaicLayoutTiles(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardMosaicLayoutTilesSlice expands the contents of DashboardMosaicLayoutTiles into a JSON
// request object.
func expandDashboardMosaicLayoutTilesSlice(c *Client, f []DashboardMosaicLayoutTiles) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardMosaicLayoutTiles(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardMosaicLayoutTilesMap flattens the contents of DashboardMosaicLayoutTiles from a JSON
// response object.
func flattenDashboardMosaicLayoutTilesMap(c *Client, i interface{}) map[string]DashboardMosaicLayoutTiles {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardMosaicLayoutTiles{}
	}

	if len(a) == 0 {
		return map[string]DashboardMosaicLayoutTiles{}
	}

	items := make(map[string]DashboardMosaicLayoutTiles)
	for k, item := range a {
		items[k] = *flattenDashboardMosaicLayoutTiles(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardMosaicLayoutTilesSlice flattens the contents of DashboardMosaicLayoutTiles from a JSON
// response object.
func flattenDashboardMosaicLayoutTilesSlice(c *Client, i interface{}) []DashboardMosaicLayoutTiles {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardMosaicLayoutTiles{}
	}

	if len(a) == 0 {
		return []DashboardMosaicLayoutTiles{}
	}

	items := make([]DashboardMosaicLayoutTiles, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardMosaicLayoutTiles(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardMosaicLayoutTiles expands an instance of DashboardMosaicLayoutTiles into a JSON
// request object.
func expandDashboardMosaicLayoutTiles(c *Client, f *DashboardMosaicLayoutTiles) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.XPos; !dcl.IsEmptyValueIndirect(v) {
		m["xPos"] = v
	}
	if v := f.YPos; !dcl.IsEmptyValueIndirect(v) {
		m["yPos"] = v
	}
	if v := f.Width; !dcl.IsEmptyValueIndirect(v) {
		m["width"] = v
	}
	if v := f.Height; !dcl.IsEmptyValueIndirect(v) {
		m["height"] = v
	}
	if v, err := expandDashboardWidget(c, f.Widget); err != nil {
		return nil, fmt.Errorf("error expanding Widget into widget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["widget"] = v
	}

	return m, nil
}

// flattenDashboardMosaicLayoutTiles flattens an instance of DashboardMosaicLayoutTiles from a JSON
// response object.
func flattenDashboardMosaicLayoutTiles(c *Client, i interface{}) *DashboardMosaicLayoutTiles {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardMosaicLayoutTiles{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardMosaicLayoutTiles
	}
	r.XPos = dcl.FlattenInteger(m["xPos"])
	r.YPos = dcl.FlattenInteger(m["yPos"])
	r.Width = dcl.FlattenInteger(m["width"])
	r.Height = dcl.FlattenInteger(m["height"])
	r.Widget = flattenDashboardWidget(c, m["widget"])

	return r
}

// expandDashboardRowLayoutMap expands the contents of DashboardRowLayout into a JSON
// request object.
func expandDashboardRowLayoutMap(c *Client, f map[string]DashboardRowLayout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardRowLayout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardRowLayoutSlice expands the contents of DashboardRowLayout into a JSON
// request object.
func expandDashboardRowLayoutSlice(c *Client, f []DashboardRowLayout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardRowLayout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardRowLayoutMap flattens the contents of DashboardRowLayout from a JSON
// response object.
func flattenDashboardRowLayoutMap(c *Client, i interface{}) map[string]DashboardRowLayout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardRowLayout{}
	}

	if len(a) == 0 {
		return map[string]DashboardRowLayout{}
	}

	items := make(map[string]DashboardRowLayout)
	for k, item := range a {
		items[k] = *flattenDashboardRowLayout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardRowLayoutSlice flattens the contents of DashboardRowLayout from a JSON
// response object.
func flattenDashboardRowLayoutSlice(c *Client, i interface{}) []DashboardRowLayout {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardRowLayout{}
	}

	if len(a) == 0 {
		return []DashboardRowLayout{}
	}

	items := make([]DashboardRowLayout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardRowLayout(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardRowLayout expands an instance of DashboardRowLayout into a JSON
// request object.
func expandDashboardRowLayout(c *Client, f *DashboardRowLayout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDashboardRowLayoutRowsSlice(c, f.Rows); err != nil {
		return nil, fmt.Errorf("error expanding Rows into rows: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rows"] = v
	}

	return m, nil
}

// flattenDashboardRowLayout flattens an instance of DashboardRowLayout from a JSON
// response object.
func flattenDashboardRowLayout(c *Client, i interface{}) *DashboardRowLayout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardRowLayout{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardRowLayout
	}
	r.Rows = flattenDashboardRowLayoutRowsSlice(c, m["rows"])

	return r
}

// expandDashboardRowLayoutRowsMap expands the contents of DashboardRowLayoutRows into a JSON
// request object.
func expandDashboardRowLayoutRowsMap(c *Client, f map[string]DashboardRowLayoutRows) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardRowLayoutRows(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardRowLayoutRowsSlice expands the contents of DashboardRowLayoutRows into a JSON
// request object.
func expandDashboardRowLayoutRowsSlice(c *Client, f []DashboardRowLayoutRows) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardRowLayoutRows(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardRowLayoutRowsMap flattens the contents of DashboardRowLayoutRows from a JSON
// response object.
func flattenDashboardRowLayoutRowsMap(c *Client, i interface{}) map[string]DashboardRowLayoutRows {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardRowLayoutRows{}
	}

	if len(a) == 0 {
		return map[string]DashboardRowLayoutRows{}
	}

	items := make(map[string]DashboardRowLayoutRows)
	for k, item := range a {
		items[k] = *flattenDashboardRowLayoutRows(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardRowLayoutRowsSlice flattens the contents of DashboardRowLayoutRows from a JSON
// response object.
func flattenDashboardRowLayoutRowsSlice(c *Client, i interface{}) []DashboardRowLayoutRows {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardRowLayoutRows{}
	}

	if len(a) == 0 {
		return []DashboardRowLayoutRows{}
	}

	items := make([]DashboardRowLayoutRows, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardRowLayoutRows(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardRowLayoutRows expands an instance of DashboardRowLayoutRows into a JSON
// request object.
func expandDashboardRowLayoutRows(c *Client, f *DashboardRowLayoutRows) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Weight; !dcl.IsEmptyValueIndirect(v) {
		m["weight"] = v
	}
	if v := f.Widgets; v != nil {
		m["widgets"] = v
	}

	return m, nil
}

// flattenDashboardRowLayoutRows flattens an instance of DashboardRowLayoutRows from a JSON
// response object.
func flattenDashboardRowLayoutRows(c *Client, i interface{}) *DashboardRowLayoutRows {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardRowLayoutRows{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardRowLayoutRows
	}
	r.Weight = dcl.FlattenInteger(m["weight"])
	r.Widgets = flattenDashboardWidgetSlice(c, m["widgets"])

	return r
}

// expandDashboardColumnLayoutMap expands the contents of DashboardColumnLayout into a JSON
// request object.
func expandDashboardColumnLayoutMap(c *Client, f map[string]DashboardColumnLayout) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardColumnLayout(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardColumnLayoutSlice expands the contents of DashboardColumnLayout into a JSON
// request object.
func expandDashboardColumnLayoutSlice(c *Client, f []DashboardColumnLayout) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardColumnLayout(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardColumnLayoutMap flattens the contents of DashboardColumnLayout from a JSON
// response object.
func flattenDashboardColumnLayoutMap(c *Client, i interface{}) map[string]DashboardColumnLayout {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardColumnLayout{}
	}

	if len(a) == 0 {
		return map[string]DashboardColumnLayout{}
	}

	items := make(map[string]DashboardColumnLayout)
	for k, item := range a {
		items[k] = *flattenDashboardColumnLayout(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardColumnLayoutSlice flattens the contents of DashboardColumnLayout from a JSON
// response object.
func flattenDashboardColumnLayoutSlice(c *Client, i interface{}) []DashboardColumnLayout {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardColumnLayout{}
	}

	if len(a) == 0 {
		return []DashboardColumnLayout{}
	}

	items := make([]DashboardColumnLayout, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardColumnLayout(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardColumnLayout expands an instance of DashboardColumnLayout into a JSON
// request object.
func expandDashboardColumnLayout(c *Client, f *DashboardColumnLayout) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDashboardColumnLayoutColumnsSlice(c, f.Columns); err != nil {
		return nil, fmt.Errorf("error expanding Columns into columns: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["columns"] = v
	}

	return m, nil
}

// flattenDashboardColumnLayout flattens an instance of DashboardColumnLayout from a JSON
// response object.
func flattenDashboardColumnLayout(c *Client, i interface{}) *DashboardColumnLayout {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardColumnLayout{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardColumnLayout
	}
	r.Columns = flattenDashboardColumnLayoutColumnsSlice(c, m["columns"])

	return r
}

// expandDashboardColumnLayoutColumnsMap expands the contents of DashboardColumnLayoutColumns into a JSON
// request object.
func expandDashboardColumnLayoutColumnsMap(c *Client, f map[string]DashboardColumnLayoutColumns) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardColumnLayoutColumns(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardColumnLayoutColumnsSlice expands the contents of DashboardColumnLayoutColumns into a JSON
// request object.
func expandDashboardColumnLayoutColumnsSlice(c *Client, f []DashboardColumnLayoutColumns) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardColumnLayoutColumns(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardColumnLayoutColumnsMap flattens the contents of DashboardColumnLayoutColumns from a JSON
// response object.
func flattenDashboardColumnLayoutColumnsMap(c *Client, i interface{}) map[string]DashboardColumnLayoutColumns {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardColumnLayoutColumns{}
	}

	if len(a) == 0 {
		return map[string]DashboardColumnLayoutColumns{}
	}

	items := make(map[string]DashboardColumnLayoutColumns)
	for k, item := range a {
		items[k] = *flattenDashboardColumnLayoutColumns(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardColumnLayoutColumnsSlice flattens the contents of DashboardColumnLayoutColumns from a JSON
// response object.
func flattenDashboardColumnLayoutColumnsSlice(c *Client, i interface{}) []DashboardColumnLayoutColumns {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardColumnLayoutColumns{}
	}

	if len(a) == 0 {
		return []DashboardColumnLayoutColumns{}
	}

	items := make([]DashboardColumnLayoutColumns, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardColumnLayoutColumns(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardColumnLayoutColumns expands an instance of DashboardColumnLayoutColumns into a JSON
// request object.
func expandDashboardColumnLayoutColumns(c *Client, f *DashboardColumnLayoutColumns) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Weight; !dcl.IsEmptyValueIndirect(v) {
		m["weight"] = v
	}
	if v, err := expandDashboardWidgetSlice(c, f.Widgets); err != nil {
		return nil, fmt.Errorf("error expanding Widgets into widgets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["widgets"] = v
	}

	return m, nil
}

// flattenDashboardColumnLayoutColumns flattens an instance of DashboardColumnLayoutColumns from a JSON
// response object.
func flattenDashboardColumnLayoutColumns(c *Client, i interface{}) *DashboardColumnLayoutColumns {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardColumnLayoutColumns{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardColumnLayoutColumns
	}
	r.Weight = dcl.FlattenInteger(m["weight"])
	r.Widgets = flattenDashboardWidgetSlice(c, m["widgets"])

	return r
}

// expandDashboardWidgetMap expands the contents of DashboardWidget into a JSON
// request object.
func expandDashboardWidgetMap(c *Client, f map[string]DashboardWidget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidget(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetSlice expands the contents of DashboardWidget into a JSON
// request object.
func expandDashboardWidgetSlice(c *Client, f []DashboardWidget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidget(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetMap flattens the contents of DashboardWidget from a JSON
// response object.
func flattenDashboardWidgetMap(c *Client, i interface{}) map[string]DashboardWidget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidget{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidget{}
	}

	items := make(map[string]DashboardWidget)
	for k, item := range a {
		items[k] = *flattenDashboardWidget(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetSlice flattens the contents of DashboardWidget from a JSON
// response object.
func flattenDashboardWidgetSlice(c *Client, i interface{}) []DashboardWidget {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidget{}
	}

	if len(a) == 0 {
		return []DashboardWidget{}
	}

	items := make([]DashboardWidget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidget(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidget expands an instance of DashboardWidget into a JSON
// request object.
func expandDashboardWidget(c *Client, f *DashboardWidget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Title; !dcl.IsEmptyValueIndirect(v) {
		m["title"] = v
	}
	if v, err := expandDashboardWidgetXyChart(c, f.XyChart); err != nil {
		return nil, fmt.Errorf("error expanding XyChart into xyChart: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["xyChart"] = v
	}
	if v, err := expandDashboardWidgetScorecard(c, f.Scorecard); err != nil {
		return nil, fmt.Errorf("error expanding Scorecard into scorecard: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scorecard"] = v
	}
	if v, err := expandDashboardWidgetText(c, f.Text); err != nil {
		return nil, fmt.Errorf("error expanding Text into text: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["text"] = v
	}
	if v, err := expandDashboardWidgetBlank(c, f.Blank); err != nil {
		return nil, fmt.Errorf("error expanding Blank into blank: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["blank"] = v
	}

	return m, nil
}

// flattenDashboardWidget flattens an instance of DashboardWidget from a JSON
// response object.
func flattenDashboardWidget(c *Client, i interface{}) *DashboardWidget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidget{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidget
	}
	r.Title = dcl.FlattenString(m["title"])
	r.XyChart = flattenDashboardWidgetXyChart(c, m["xyChart"])
	r.Scorecard = flattenDashboardWidgetScorecard(c, m["scorecard"])
	r.Text = flattenDashboardWidgetText(c, m["text"])
	r.Blank = flattenDashboardWidgetBlank(c, m["blank"])

	return r
}

// expandDashboardWidgetXyChartMap expands the contents of DashboardWidgetXyChart into a JSON
// request object.
func expandDashboardWidgetXyChartMap(c *Client, f map[string]DashboardWidgetXyChart) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChart(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartSlice expands the contents of DashboardWidgetXyChart into a JSON
// request object.
func expandDashboardWidgetXyChartSlice(c *Client, f []DashboardWidgetXyChart) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChart(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartMap flattens the contents of DashboardWidgetXyChart from a JSON
// response object.
func flattenDashboardWidgetXyChartMap(c *Client, i interface{}) map[string]DashboardWidgetXyChart {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChart{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChart{}
	}

	items := make(map[string]DashboardWidgetXyChart)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChart(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartSlice flattens the contents of DashboardWidgetXyChart from a JSON
// response object.
func flattenDashboardWidgetXyChartSlice(c *Client, i interface{}) []DashboardWidgetXyChart {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChart{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChart{}
	}

	items := make([]DashboardWidgetXyChart, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChart(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChart expands an instance of DashboardWidgetXyChart into a JSON
// request object.
func expandDashboardWidgetXyChart(c *Client, f *DashboardWidgetXyChart) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDashboardWidgetXyChartDataSetsSlice(c, f.DataSets); err != nil {
		return nil, fmt.Errorf("error expanding DataSets into dataSets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dataSets"] = v
	}
	if v := f.TimeshiftDuration; !dcl.IsEmptyValueIndirect(v) {
		m["timeshiftDuration"] = v
	}
	if v, err := expandDashboardWidgetXyChartThresholdsSlice(c, f.Thresholds); err != nil {
		return nil, fmt.Errorf("error expanding Thresholds into thresholds: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["thresholds"] = v
	}
	if v, err := expandDashboardWidgetXyChartXAxis(c, f.XAxis); err != nil {
		return nil, fmt.Errorf("error expanding XAxis into xAxis: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["xAxis"] = v
	}
	if v, err := expandDashboardWidgetXyChartYAxis(c, f.YAxis); err != nil {
		return nil, fmt.Errorf("error expanding YAxis into yAxis: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["yAxis"] = v
	}
	if v, err := expandDashboardWidgetXyChartChartOptions(c, f.ChartOptions); err != nil {
		return nil, fmt.Errorf("error expanding ChartOptions into chartOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["chartOptions"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChart flattens an instance of DashboardWidgetXyChart from a JSON
// response object.
func flattenDashboardWidgetXyChart(c *Client, i interface{}) *DashboardWidgetXyChart {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChart{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChart
	}
	r.DataSets = flattenDashboardWidgetXyChartDataSetsSlice(c, m["dataSets"])
	r.TimeshiftDuration = dcl.FlattenString(m["timeshiftDuration"])
	r.Thresholds = flattenDashboardWidgetXyChartThresholdsSlice(c, m["thresholds"])
	r.XAxis = flattenDashboardWidgetXyChartXAxis(c, m["xAxis"])
	r.YAxis = flattenDashboardWidgetXyChartYAxis(c, m["yAxis"])
	r.ChartOptions = flattenDashboardWidgetXyChartChartOptions(c, m["chartOptions"])

	return r
}

// expandDashboardWidgetXyChartDataSetsMap expands the contents of DashboardWidgetXyChartDataSets into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsMap(c *Client, f map[string]DashboardWidgetXyChartDataSets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsSlice expands the contents of DashboardWidgetXyChartDataSets into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsSlice(c *Client, f []DashboardWidgetXyChartDataSets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsMap flattens the contents of DashboardWidgetXyChartDataSets from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSets{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSets{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSets)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsSlice flattens the contents of DashboardWidgetXyChartDataSets from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSets {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSets{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSets{}
	}

	items := make([]DashboardWidgetXyChartDataSets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSets(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSets expands an instance of DashboardWidgetXyChartDataSets into a JSON
// request object.
func expandDashboardWidgetXyChartDataSets(c *Client, f *DashboardWidgetXyChartDataSets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQuery(c, f.TimeSeriesQuery); err != nil {
		return nil, fmt.Errorf("error expanding TimeSeriesQuery into timeSeriesQuery: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeSeriesQuery"] = v
	}
	if v := f.PlotType; !dcl.IsEmptyValueIndirect(v) {
		m["plotType"] = v
	}
	if v := f.LegendTemplate; !dcl.IsEmptyValueIndirect(v) {
		m["legendTemplate"] = v
	}
	if v := f.MinAlignmentPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["minAlignmentPeriod"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSets flattens an instance of DashboardWidgetXyChartDataSets from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSets(c *Client, i interface{}) *DashboardWidgetXyChartDataSets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSets
	}
	r.TimeSeriesQuery = flattenDashboardWidgetXyChartDataSetsTimeSeriesQuery(c, m["timeSeriesQuery"])
	r.PlotType = flattenDashboardWidgetXyChartDataSetsPlotTypeEnum(m["plotType"])
	r.LegendTemplate = dcl.FlattenString(m["legendTemplate"])
	r.MinAlignmentPeriod = dcl.FlattenString(m["minAlignmentPeriod"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQuery into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQuery) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQuery(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQuerySlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQuery into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQuerySlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQuery) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQuery(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQuery from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQuery{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQuery{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQuery)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQuery(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQuerySlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQuery from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQuerySlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQuery{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQuery{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQuery, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQuery(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQuery expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQuery into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQuery(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQuery) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c, f.TimeSeriesFilter); err != nil {
		return nil, fmt.Errorf("error expanding TimeSeriesFilter into timeSeriesFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeSeriesFilter"] = v
	}
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c, f.TimeSeriesFilterRatio); err != nil {
		return nil, fmt.Errorf("error expanding TimeSeriesFilterRatio into timeSeriesFilterRatio: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeSeriesFilterRatio"] = v
	}
	if v := f.TimeSeriesQueryLanguage; !dcl.IsEmptyValueIndirect(v) {
		m["timeSeriesQueryLanguage"] = v
	}
	if v := f.UnitOverride; !dcl.IsEmptyValueIndirect(v) {
		m["unitOverride"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQuery flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQuery from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQuery(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQuery{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQuery
	}
	r.TimeSeriesFilter = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c, m["timeSeriesFilter"])
	r.TimeSeriesFilterRatio = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c, m["timeSeriesFilterRatio"])
	r.TimeSeriesQueryLanguage = dcl.FlattenString(m["timeSeriesQueryLanguage"])
	r.UnitOverride = dcl.FlattenString(m["unitOverride"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c, f.Aggregation); err != nil {
		return nil, fmt.Errorf("error expanding Aggregation into aggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aggregation"] = v
	}
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, f.SecondaryAggregation); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryAggregation into secondaryAggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secondaryAggregation"] = v
	}
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, f.PickTimeSeriesFilter); err != nil {
		return nil, fmt.Errorf("error expanding PickTimeSeriesFilter into pickTimeSeriesFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pickTimeSeriesFilter"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.Aggregation = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c, m["aggregation"])
	r.SecondaryAggregation = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, m["secondaryAggregation"])
	r.PickTimeSeriesFilter = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, m["pickTimeSeriesFilter"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RankingMethod; !dcl.IsEmptyValueIndirect(v) {
		m["rankingMethod"] = v
	}
	if v := f.NumTimeSeries; !dcl.IsEmptyValueIndirect(v) {
		m["numTimeSeries"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		m["direction"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	}
	r.RankingMethod = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(m["rankingMethod"])
	r.NumTimeSeries = dcl.FlattenInteger(m["numTimeSeries"])
	r.Direction = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(m["direction"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, f.Numerator); err != nil {
		return nil, fmt.Errorf("error expanding Numerator into numerator: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["numerator"] = v
	}
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, f.Denominator); err != nil {
		return nil, fmt.Errorf("error expanding Denominator into denominator: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["denominator"] = v
	}
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, f.SecondaryAggregation); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryAggregation into secondaryAggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secondaryAggregation"] = v
	}
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, f.PickTimeSeriesFilter); err != nil {
		return nil, fmt.Errorf("error expanding PickTimeSeriesFilter into pickTimeSeriesFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pickTimeSeriesFilter"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio
	}
	r.Numerator = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, m["numerator"])
	r.Denominator = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, m["denominator"])
	r.SecondaryAggregation = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, m["secondaryAggregation"])
	r.PickTimeSeriesFilter = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, m["pickTimeSeriesFilter"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, f.Aggregation); err != nil {
		return nil, fmt.Errorf("error expanding Aggregation into aggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aggregation"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.Aggregation = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, m["aggregation"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, f.Aggregation); err != nil {
		return nil, fmt.Errorf("error expanding Aggregation into aggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aggregation"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.Aggregation = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, m["aggregation"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterMap expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterMap(c *Client, f map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice expands the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice(c *Client, f []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter expands an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c *Client, f *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RankingMethod; !dcl.IsEmptyValueIndirect(v) {
		m["rankingMethod"] = v
	}
	if v := f.NumTimeSeries; !dcl.IsEmptyValueIndirect(v) {
		m["numTimeSeries"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		m["direction"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter flattens an instance of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c *Client, i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	}
	r.RankingMethod = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(m["rankingMethod"])
	r.NumTimeSeries = dcl.FlattenInteger(m["numTimeSeries"])
	r.Direction = flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(m["direction"])

	return r
}

// expandDashboardWidgetXyChartThresholdsMap expands the contents of DashboardWidgetXyChartThresholds into a JSON
// request object.
func expandDashboardWidgetXyChartThresholdsMap(c *Client, f map[string]DashboardWidgetXyChartThresholds) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartThresholds(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartThresholdsSlice expands the contents of DashboardWidgetXyChartThresholds into a JSON
// request object.
func expandDashboardWidgetXyChartThresholdsSlice(c *Client, f []DashboardWidgetXyChartThresholds) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartThresholds(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartThresholdsMap flattens the contents of DashboardWidgetXyChartThresholds from a JSON
// response object.
func flattenDashboardWidgetXyChartThresholdsMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartThresholds {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartThresholds{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartThresholds{}
	}

	items := make(map[string]DashboardWidgetXyChartThresholds)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartThresholds(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartThresholdsSlice flattens the contents of DashboardWidgetXyChartThresholds from a JSON
// response object.
func flattenDashboardWidgetXyChartThresholdsSlice(c *Client, i interface{}) []DashboardWidgetXyChartThresholds {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartThresholds{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartThresholds{}
	}

	items := make([]DashboardWidgetXyChartThresholds, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartThresholds(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartThresholds expands an instance of DashboardWidgetXyChartThresholds into a JSON
// request object.
func expandDashboardWidgetXyChartThresholds(c *Client, f *DashboardWidgetXyChartThresholds) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Label; !dcl.IsEmptyValueIndirect(v) {
		m["label"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.Color; !dcl.IsEmptyValueIndirect(v) {
		m["color"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		m["direction"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartThresholds flattens an instance of DashboardWidgetXyChartThresholds from a JSON
// response object.
func flattenDashboardWidgetXyChartThresholds(c *Client, i interface{}) *DashboardWidgetXyChartThresholds {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartThresholds{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartThresholds
	}
	r.Label = dcl.FlattenString(m["label"])
	r.Value = dcl.FlattenDouble(m["value"])
	r.Color = flattenDashboardWidgetXyChartThresholdsColorEnum(m["color"])
	r.Direction = flattenDashboardWidgetXyChartThresholdsDirectionEnum(m["direction"])

	return r
}

// expandDashboardWidgetXyChartXAxisMap expands the contents of DashboardWidgetXyChartXAxis into a JSON
// request object.
func expandDashboardWidgetXyChartXAxisMap(c *Client, f map[string]DashboardWidgetXyChartXAxis) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartXAxis(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartXAxisSlice expands the contents of DashboardWidgetXyChartXAxis into a JSON
// request object.
func expandDashboardWidgetXyChartXAxisSlice(c *Client, f []DashboardWidgetXyChartXAxis) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartXAxis(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartXAxisMap flattens the contents of DashboardWidgetXyChartXAxis from a JSON
// response object.
func flattenDashboardWidgetXyChartXAxisMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartXAxis {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartXAxis{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartXAxis{}
	}

	items := make(map[string]DashboardWidgetXyChartXAxis)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartXAxis(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartXAxisSlice flattens the contents of DashboardWidgetXyChartXAxis from a JSON
// response object.
func flattenDashboardWidgetXyChartXAxisSlice(c *Client, i interface{}) []DashboardWidgetXyChartXAxis {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartXAxis{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartXAxis{}
	}

	items := make([]DashboardWidgetXyChartXAxis, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartXAxis(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartXAxis expands an instance of DashboardWidgetXyChartXAxis into a JSON
// request object.
func expandDashboardWidgetXyChartXAxis(c *Client, f *DashboardWidgetXyChartXAxis) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Label; !dcl.IsEmptyValueIndirect(v) {
		m["label"] = v
	}
	if v := f.Scale; !dcl.IsEmptyValueIndirect(v) {
		m["scale"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartXAxis flattens an instance of DashboardWidgetXyChartXAxis from a JSON
// response object.
func flattenDashboardWidgetXyChartXAxis(c *Client, i interface{}) *DashboardWidgetXyChartXAxis {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartXAxis{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartXAxis
	}
	r.Label = dcl.FlattenString(m["label"])
	r.Scale = flattenDashboardWidgetXyChartXAxisScaleEnum(m["scale"])

	return r
}

// expandDashboardWidgetXyChartYAxisMap expands the contents of DashboardWidgetXyChartYAxis into a JSON
// request object.
func expandDashboardWidgetXyChartYAxisMap(c *Client, f map[string]DashboardWidgetXyChartYAxis) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartYAxis(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartYAxisSlice expands the contents of DashboardWidgetXyChartYAxis into a JSON
// request object.
func expandDashboardWidgetXyChartYAxisSlice(c *Client, f []DashboardWidgetXyChartYAxis) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartYAxis(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartYAxisMap flattens the contents of DashboardWidgetXyChartYAxis from a JSON
// response object.
func flattenDashboardWidgetXyChartYAxisMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartYAxis {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartYAxis{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartYAxis{}
	}

	items := make(map[string]DashboardWidgetXyChartYAxis)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartYAxis(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartYAxisSlice flattens the contents of DashboardWidgetXyChartYAxis from a JSON
// response object.
func flattenDashboardWidgetXyChartYAxisSlice(c *Client, i interface{}) []DashboardWidgetXyChartYAxis {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartYAxis{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartYAxis{}
	}

	items := make([]DashboardWidgetXyChartYAxis, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartYAxis(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartYAxis expands an instance of DashboardWidgetXyChartYAxis into a JSON
// request object.
func expandDashboardWidgetXyChartYAxis(c *Client, f *DashboardWidgetXyChartYAxis) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Label; !dcl.IsEmptyValueIndirect(v) {
		m["label"] = v
	}
	if v := f.Scale; !dcl.IsEmptyValueIndirect(v) {
		m["scale"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartYAxis flattens an instance of DashboardWidgetXyChartYAxis from a JSON
// response object.
func flattenDashboardWidgetXyChartYAxis(c *Client, i interface{}) *DashboardWidgetXyChartYAxis {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartYAxis{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartYAxis
	}
	r.Label = dcl.FlattenString(m["label"])
	r.Scale = flattenDashboardWidgetXyChartYAxisScaleEnum(m["scale"])

	return r
}

// expandDashboardWidgetXyChartChartOptionsMap expands the contents of DashboardWidgetXyChartChartOptions into a JSON
// request object.
func expandDashboardWidgetXyChartChartOptionsMap(c *Client, f map[string]DashboardWidgetXyChartChartOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetXyChartChartOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetXyChartChartOptionsSlice expands the contents of DashboardWidgetXyChartChartOptions into a JSON
// request object.
func expandDashboardWidgetXyChartChartOptionsSlice(c *Client, f []DashboardWidgetXyChartChartOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetXyChartChartOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetXyChartChartOptionsMap flattens the contents of DashboardWidgetXyChartChartOptions from a JSON
// response object.
func flattenDashboardWidgetXyChartChartOptionsMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartChartOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartChartOptions{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartChartOptions{}
	}

	items := make(map[string]DashboardWidgetXyChartChartOptions)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartChartOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartChartOptionsSlice flattens the contents of DashboardWidgetXyChartChartOptions from a JSON
// response object.
func flattenDashboardWidgetXyChartChartOptionsSlice(c *Client, i interface{}) []DashboardWidgetXyChartChartOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartChartOptions{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartChartOptions{}
	}

	items := make([]DashboardWidgetXyChartChartOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartChartOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetXyChartChartOptions expands an instance of DashboardWidgetXyChartChartOptions into a JSON
// request object.
func expandDashboardWidgetXyChartChartOptions(c *Client, f *DashboardWidgetXyChartChartOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenDashboardWidgetXyChartChartOptions flattens an instance of DashboardWidgetXyChartChartOptions from a JSON
// response object.
func flattenDashboardWidgetXyChartChartOptions(c *Client, i interface{}) *DashboardWidgetXyChartChartOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetXyChartChartOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetXyChartChartOptions
	}
	r.Mode = flattenDashboardWidgetXyChartChartOptionsModeEnum(m["mode"])

	return r
}

// expandDashboardWidgetScorecardMap expands the contents of DashboardWidgetScorecard into a JSON
// request object.
func expandDashboardWidgetScorecardMap(c *Client, f map[string]DashboardWidgetScorecard) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecard(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardSlice expands the contents of DashboardWidgetScorecard into a JSON
// request object.
func expandDashboardWidgetScorecardSlice(c *Client, f []DashboardWidgetScorecard) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecard(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardMap flattens the contents of DashboardWidgetScorecard from a JSON
// response object.
func flattenDashboardWidgetScorecardMap(c *Client, i interface{}) map[string]DashboardWidgetScorecard {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecard{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecard{}
	}

	items := make(map[string]DashboardWidgetScorecard)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecard(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardSlice flattens the contents of DashboardWidgetScorecard from a JSON
// response object.
func flattenDashboardWidgetScorecardSlice(c *Client, i interface{}) []DashboardWidgetScorecard {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecard{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecard{}
	}

	items := make([]DashboardWidgetScorecard, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecard(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecard expands an instance of DashboardWidgetScorecard into a JSON
// request object.
func expandDashboardWidgetScorecard(c *Client, f *DashboardWidgetScorecard) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDashboardWidgetScorecardTimeSeriesQuery(c, f.TimeSeriesQuery); err != nil {
		return nil, fmt.Errorf("error expanding TimeSeriesQuery into timeSeriesQuery: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeSeriesQuery"] = v
	}
	if v, err := expandDashboardWidgetScorecardGaugeView(c, f.GaugeView); err != nil {
		return nil, fmt.Errorf("error expanding GaugeView into gaugeView: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gaugeView"] = v
	}
	if v, err := expandDashboardWidgetScorecardSparkChartView(c, f.SparkChartView); err != nil {
		return nil, fmt.Errorf("error expanding SparkChartView into sparkChartView: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sparkChartView"] = v
	}
	if v, err := expandDashboardWidgetScorecardThresholdsSlice(c, f.Thresholds); err != nil {
		return nil, fmt.Errorf("error expanding Thresholds into thresholds: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["thresholds"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecard flattens an instance of DashboardWidgetScorecard from a JSON
// response object.
func flattenDashboardWidgetScorecard(c *Client, i interface{}) *DashboardWidgetScorecard {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecard{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecard
	}
	r.TimeSeriesQuery = flattenDashboardWidgetScorecardTimeSeriesQuery(c, m["timeSeriesQuery"])
	r.GaugeView = flattenDashboardWidgetScorecardGaugeView(c, m["gaugeView"])
	r.SparkChartView = flattenDashboardWidgetScorecardSparkChartView(c, m["sparkChartView"])
	r.Thresholds = flattenDashboardWidgetScorecardThresholdsSlice(c, m["thresholds"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryMap expands the contents of DashboardWidgetScorecardTimeSeriesQuery into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQuery) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQuery(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQuerySlice expands the contents of DashboardWidgetScorecardTimeSeriesQuery into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQuerySlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQuery) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQuery(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryMap flattens the contents of DashboardWidgetScorecardTimeSeriesQuery from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQuery {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQuery{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQuery{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQuery)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQuery(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQuerySlice flattens the contents of DashboardWidgetScorecardTimeSeriesQuery from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQuerySlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQuery {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQuery{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQuery{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQuery, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQuery(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQuery expands an instance of DashboardWidgetScorecardTimeSeriesQuery into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQuery(c *Client, f *DashboardWidgetScorecardTimeSeriesQuery) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c, f.TimeSeriesFilter); err != nil {
		return nil, fmt.Errorf("error expanding TimeSeriesFilter into timeSeriesFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeSeriesFilter"] = v
	}
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c, f.TimeSeriesFilterRatio); err != nil {
		return nil, fmt.Errorf("error expanding TimeSeriesFilterRatio into timeSeriesFilterRatio: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeSeriesFilterRatio"] = v
	}
	if v := f.TimeSeriesQueryLanguage; !dcl.IsEmptyValueIndirect(v) {
		m["timeSeriesQueryLanguage"] = v
	}
	if v := f.UnitOverride; !dcl.IsEmptyValueIndirect(v) {
		m["unitOverride"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQuery flattens an instance of DashboardWidgetScorecardTimeSeriesQuery from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQuery(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQuery {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQuery{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQuery
	}
	r.TimeSeriesFilter = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c, m["timeSeriesFilter"])
	r.TimeSeriesFilterRatio = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c, m["timeSeriesFilterRatio"])
	r.TimeSeriesQueryLanguage = dcl.FlattenString(m["timeSeriesQueryLanguage"])
	r.UnitOverride = dcl.FlattenString(m["unitOverride"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c, f.Aggregation); err != nil {
		return nil, fmt.Errorf("error expanding Aggregation into aggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aggregation"] = v
	}
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, f.SecondaryAggregation); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryAggregation into secondaryAggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secondaryAggregation"] = v
	}
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, f.PickTimeSeriesFilter); err != nil {
		return nil, fmt.Errorf("error expanding PickTimeSeriesFilter into pickTimeSeriesFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pickTimeSeriesFilter"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.Aggregation = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c, m["aggregation"])
	r.SecondaryAggregation = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, m["secondaryAggregation"])
	r.PickTimeSeriesFilter = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, m["pickTimeSeriesFilter"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RankingMethod; !dcl.IsEmptyValueIndirect(v) {
		m["rankingMethod"] = v
	}
	if v := f.NumTimeSeries; !dcl.IsEmptyValueIndirect(v) {
		m["numTimeSeries"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		m["direction"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	}
	r.RankingMethod = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(m["rankingMethod"])
	r.NumTimeSeries = dcl.FlattenInteger(m["numTimeSeries"])
	r.Direction = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(m["direction"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, f.Numerator); err != nil {
		return nil, fmt.Errorf("error expanding Numerator into numerator: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["numerator"] = v
	}
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, f.Denominator); err != nil {
		return nil, fmt.Errorf("error expanding Denominator into denominator: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["denominator"] = v
	}
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, f.SecondaryAggregation); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryAggregation into secondaryAggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secondaryAggregation"] = v
	}
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, f.PickTimeSeriesFilter); err != nil {
		return nil, fmt.Errorf("error expanding PickTimeSeriesFilter into pickTimeSeriesFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pickTimeSeriesFilter"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio
	}
	r.Numerator = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, m["numerator"])
	r.Denominator = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, m["denominator"])
	r.SecondaryAggregation = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, m["secondaryAggregation"])
	r.PickTimeSeriesFilter = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, m["pickTimeSeriesFilter"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, f.Aggregation); err != nil {
		return nil, fmt.Errorf("error expanding Aggregation into aggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aggregation"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.Aggregation = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, m["aggregation"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, f.Aggregation); err != nil {
		return nil, fmt.Errorf("error expanding Aggregation into aggregation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aggregation"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
	}
	r.Filter = dcl.FlattenString(m["filter"])
	r.Aggregation = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, m["aggregation"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) (map[string]interface{}, error) {
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
	if v := f.GroupByFields; v != nil {
		m["groupByFields"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
	}
	r.AlignmentPeriod = dcl.FlattenString(m["alignmentPeriod"])
	r.PerSeriesAligner = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(m["perSeriesAligner"])
	r.CrossSeriesReducer = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(m["crossSeriesReducer"])
	r.GroupByFields = dcl.FlattenStringSlice(m["groupByFields"])

	return r
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterMap expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterMap(c *Client, f map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice expands the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice(c *Client, f []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter expands an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter into a JSON
// request object.
func expandDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c *Client, f *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RankingMethod; !dcl.IsEmptyValueIndirect(v) {
		m["rankingMethod"] = v
	}
	if v := f.NumTimeSeries; !dcl.IsEmptyValueIndirect(v) {
		m["numTimeSeries"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		m["direction"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter flattens an instance of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(c *Client, i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	}
	r.RankingMethod = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(m["rankingMethod"])
	r.NumTimeSeries = dcl.FlattenInteger(m["numTimeSeries"])
	r.Direction = flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(m["direction"])

	return r
}

// expandDashboardWidgetScorecardGaugeViewMap expands the contents of DashboardWidgetScorecardGaugeView into a JSON
// request object.
func expandDashboardWidgetScorecardGaugeViewMap(c *Client, f map[string]DashboardWidgetScorecardGaugeView) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardGaugeView(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardGaugeViewSlice expands the contents of DashboardWidgetScorecardGaugeView into a JSON
// request object.
func expandDashboardWidgetScorecardGaugeViewSlice(c *Client, f []DashboardWidgetScorecardGaugeView) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardGaugeView(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardGaugeViewMap flattens the contents of DashboardWidgetScorecardGaugeView from a JSON
// response object.
func flattenDashboardWidgetScorecardGaugeViewMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardGaugeView {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardGaugeView{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardGaugeView{}
	}

	items := make(map[string]DashboardWidgetScorecardGaugeView)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardGaugeView(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardGaugeViewSlice flattens the contents of DashboardWidgetScorecardGaugeView from a JSON
// response object.
func flattenDashboardWidgetScorecardGaugeViewSlice(c *Client, i interface{}) []DashboardWidgetScorecardGaugeView {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardGaugeView{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardGaugeView{}
	}

	items := make([]DashboardWidgetScorecardGaugeView, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardGaugeView(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardGaugeView expands an instance of DashboardWidgetScorecardGaugeView into a JSON
// request object.
func expandDashboardWidgetScorecardGaugeView(c *Client, f *DashboardWidgetScorecardGaugeView) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.LowerBound; !dcl.IsEmptyValueIndirect(v) {
		m["lowerBound"] = v
	}
	if v := f.UpperBound; !dcl.IsEmptyValueIndirect(v) {
		m["upperBound"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardGaugeView flattens an instance of DashboardWidgetScorecardGaugeView from a JSON
// response object.
func flattenDashboardWidgetScorecardGaugeView(c *Client, i interface{}) *DashboardWidgetScorecardGaugeView {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardGaugeView{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardGaugeView
	}
	r.LowerBound = dcl.FlattenDouble(m["lowerBound"])
	r.UpperBound = dcl.FlattenDouble(m["upperBound"])

	return r
}

// expandDashboardWidgetScorecardSparkChartViewMap expands the contents of DashboardWidgetScorecardSparkChartView into a JSON
// request object.
func expandDashboardWidgetScorecardSparkChartViewMap(c *Client, f map[string]DashboardWidgetScorecardSparkChartView) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardSparkChartView(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardSparkChartViewSlice expands the contents of DashboardWidgetScorecardSparkChartView into a JSON
// request object.
func expandDashboardWidgetScorecardSparkChartViewSlice(c *Client, f []DashboardWidgetScorecardSparkChartView) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardSparkChartView(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardSparkChartViewMap flattens the contents of DashboardWidgetScorecardSparkChartView from a JSON
// response object.
func flattenDashboardWidgetScorecardSparkChartViewMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardSparkChartView {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardSparkChartView{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardSparkChartView{}
	}

	items := make(map[string]DashboardWidgetScorecardSparkChartView)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardSparkChartView(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardSparkChartViewSlice flattens the contents of DashboardWidgetScorecardSparkChartView from a JSON
// response object.
func flattenDashboardWidgetScorecardSparkChartViewSlice(c *Client, i interface{}) []DashboardWidgetScorecardSparkChartView {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardSparkChartView{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardSparkChartView{}
	}

	items := make([]DashboardWidgetScorecardSparkChartView, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardSparkChartView(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardSparkChartView expands an instance of DashboardWidgetScorecardSparkChartView into a JSON
// request object.
func expandDashboardWidgetScorecardSparkChartView(c *Client, f *DashboardWidgetScorecardSparkChartView) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SparkChartType; !dcl.IsEmptyValueIndirect(v) {
		m["sparkChartType"] = v
	}
	if v := f.MinAlignmentPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["minAlignmentPeriod"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardSparkChartView flattens an instance of DashboardWidgetScorecardSparkChartView from a JSON
// response object.
func flattenDashboardWidgetScorecardSparkChartView(c *Client, i interface{}) *DashboardWidgetScorecardSparkChartView {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardSparkChartView{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardSparkChartView
	}
	r.SparkChartType = flattenDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(m["sparkChartType"])
	r.MinAlignmentPeriod = dcl.FlattenString(m["minAlignmentPeriod"])

	return r
}

// expandDashboardWidgetScorecardThresholdsMap expands the contents of DashboardWidgetScorecardThresholds into a JSON
// request object.
func expandDashboardWidgetScorecardThresholdsMap(c *Client, f map[string]DashboardWidgetScorecardThresholds) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetScorecardThresholds(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetScorecardThresholdsSlice expands the contents of DashboardWidgetScorecardThresholds into a JSON
// request object.
func expandDashboardWidgetScorecardThresholdsSlice(c *Client, f []DashboardWidgetScorecardThresholds) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetScorecardThresholds(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetScorecardThresholdsMap flattens the contents of DashboardWidgetScorecardThresholds from a JSON
// response object.
func flattenDashboardWidgetScorecardThresholdsMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardThresholds {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardThresholds{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardThresholds{}
	}

	items := make(map[string]DashboardWidgetScorecardThresholds)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardThresholds(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardThresholdsSlice flattens the contents of DashboardWidgetScorecardThresholds from a JSON
// response object.
func flattenDashboardWidgetScorecardThresholdsSlice(c *Client, i interface{}) []DashboardWidgetScorecardThresholds {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardThresholds{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardThresholds{}
	}

	items := make([]DashboardWidgetScorecardThresholds, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardThresholds(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetScorecardThresholds expands an instance of DashboardWidgetScorecardThresholds into a JSON
// request object.
func expandDashboardWidgetScorecardThresholds(c *Client, f *DashboardWidgetScorecardThresholds) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Label; !dcl.IsEmptyValueIndirect(v) {
		m["label"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.Color; !dcl.IsEmptyValueIndirect(v) {
		m["color"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		m["direction"] = v
	}

	return m, nil
}

// flattenDashboardWidgetScorecardThresholds flattens an instance of DashboardWidgetScorecardThresholds from a JSON
// response object.
func flattenDashboardWidgetScorecardThresholds(c *Client, i interface{}) *DashboardWidgetScorecardThresholds {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetScorecardThresholds{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetScorecardThresholds
	}
	r.Label = dcl.FlattenString(m["label"])
	r.Value = dcl.FlattenDouble(m["value"])
	r.Color = flattenDashboardWidgetScorecardThresholdsColorEnum(m["color"])
	r.Direction = flattenDashboardWidgetScorecardThresholdsDirectionEnum(m["direction"])

	return r
}

// expandDashboardWidgetTextMap expands the contents of DashboardWidgetText into a JSON
// request object.
func expandDashboardWidgetTextMap(c *Client, f map[string]DashboardWidgetText) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetText(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetTextSlice expands the contents of DashboardWidgetText into a JSON
// request object.
func expandDashboardWidgetTextSlice(c *Client, f []DashboardWidgetText) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetText(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetTextMap flattens the contents of DashboardWidgetText from a JSON
// response object.
func flattenDashboardWidgetTextMap(c *Client, i interface{}) map[string]DashboardWidgetText {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetText{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetText{}
	}

	items := make(map[string]DashboardWidgetText)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetText(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetTextSlice flattens the contents of DashboardWidgetText from a JSON
// response object.
func flattenDashboardWidgetTextSlice(c *Client, i interface{}) []DashboardWidgetText {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetText{}
	}

	if len(a) == 0 {
		return []DashboardWidgetText{}
	}

	items := make([]DashboardWidgetText, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetText(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetText expands an instance of DashboardWidgetText into a JSON
// request object.
func expandDashboardWidgetText(c *Client, f *DashboardWidgetText) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}
	if v := f.Format; !dcl.IsEmptyValueIndirect(v) {
		m["format"] = v
	}

	return m, nil
}

// flattenDashboardWidgetText flattens an instance of DashboardWidgetText from a JSON
// response object.
func flattenDashboardWidgetText(c *Client, i interface{}) *DashboardWidgetText {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetText{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetText
	}
	r.Content = dcl.FlattenString(m["content"])
	r.Format = flattenDashboardWidgetTextFormatEnum(m["format"])

	return r
}

// expandDashboardWidgetBlankMap expands the contents of DashboardWidgetBlank into a JSON
// request object.
func expandDashboardWidgetBlankMap(c *Client, f map[string]DashboardWidgetBlank) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDashboardWidgetBlank(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDashboardWidgetBlankSlice expands the contents of DashboardWidgetBlank into a JSON
// request object.
func expandDashboardWidgetBlankSlice(c *Client, f []DashboardWidgetBlank) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDashboardWidgetBlank(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDashboardWidgetBlankMap flattens the contents of DashboardWidgetBlank from a JSON
// response object.
func flattenDashboardWidgetBlankMap(c *Client, i interface{}) map[string]DashboardWidgetBlank {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetBlank{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetBlank{}
	}

	items := make(map[string]DashboardWidgetBlank)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetBlank(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDashboardWidgetBlankSlice flattens the contents of DashboardWidgetBlank from a JSON
// response object.
func flattenDashboardWidgetBlankSlice(c *Client, i interface{}) []DashboardWidgetBlank {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetBlank{}
	}

	if len(a) == 0 {
		return []DashboardWidgetBlank{}
	}

	items := make([]DashboardWidgetBlank, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetBlank(c, item.(map[string]interface{})))
	}

	return items
}

// expandDashboardWidgetBlank expands an instance of DashboardWidgetBlank into a JSON
// request object.
func expandDashboardWidgetBlank(c *Client, f *DashboardWidgetBlank) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenDashboardWidgetBlank flattens an instance of DashboardWidgetBlank from a JSON
// response object.
func flattenDashboardWidgetBlank(c *Client, i interface{}) *DashboardWidgetBlank {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DashboardWidgetBlank{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDashboardWidgetBlank
	}

	return r
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumMap flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(i interface{}) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
}

// flattenDashboardWidgetXyChartDataSetsPlotTypeEnumMap flattens the contents of DashboardWidgetXyChartDataSetsPlotTypeEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsPlotTypeEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartDataSetsPlotTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartDataSetsPlotTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartDataSetsPlotTypeEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartDataSetsPlotTypeEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartDataSetsPlotTypeEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsPlotTypeEnumSlice flattens the contents of DashboardWidgetXyChartDataSetsPlotTypeEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartDataSetsPlotTypeEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartDataSetsPlotTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartDataSetsPlotTypeEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartDataSetsPlotTypeEnum{}
	}

	items := make([]DashboardWidgetXyChartDataSetsPlotTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartDataSetsPlotTypeEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartDataSetsPlotTypeEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartDataSetsPlotTypeEnum with the same value as that string.
func flattenDashboardWidgetXyChartDataSetsPlotTypeEnum(i interface{}) *DashboardWidgetXyChartDataSetsPlotTypeEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartDataSetsPlotTypeEnumRef("")
	}

	return DashboardWidgetXyChartDataSetsPlotTypeEnumRef(s)
}

// flattenDashboardWidgetXyChartThresholdsColorEnumMap flattens the contents of DashboardWidgetXyChartThresholdsColorEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartThresholdsColorEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartThresholdsColorEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartThresholdsColorEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartThresholdsColorEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartThresholdsColorEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartThresholdsColorEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartThresholdsColorEnumSlice flattens the contents of DashboardWidgetXyChartThresholdsColorEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartThresholdsColorEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartThresholdsColorEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartThresholdsColorEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartThresholdsColorEnum{}
	}

	items := make([]DashboardWidgetXyChartThresholdsColorEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartThresholdsColorEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartThresholdsColorEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartThresholdsColorEnum with the same value as that string.
func flattenDashboardWidgetXyChartThresholdsColorEnum(i interface{}) *DashboardWidgetXyChartThresholdsColorEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartThresholdsColorEnumRef("")
	}

	return DashboardWidgetXyChartThresholdsColorEnumRef(s)
}

// flattenDashboardWidgetXyChartThresholdsDirectionEnumMap flattens the contents of DashboardWidgetXyChartThresholdsDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartThresholdsDirectionEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartThresholdsDirectionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartThresholdsDirectionEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartThresholdsDirectionEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartThresholdsDirectionEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartThresholdsDirectionEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartThresholdsDirectionEnumSlice flattens the contents of DashboardWidgetXyChartThresholdsDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartThresholdsDirectionEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartThresholdsDirectionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartThresholdsDirectionEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartThresholdsDirectionEnum{}
	}

	items := make([]DashboardWidgetXyChartThresholdsDirectionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartThresholdsDirectionEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartThresholdsDirectionEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartThresholdsDirectionEnum with the same value as that string.
func flattenDashboardWidgetXyChartThresholdsDirectionEnum(i interface{}) *DashboardWidgetXyChartThresholdsDirectionEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartThresholdsDirectionEnumRef("")
	}

	return DashboardWidgetXyChartThresholdsDirectionEnumRef(s)
}

// flattenDashboardWidgetXyChartXAxisScaleEnumMap flattens the contents of DashboardWidgetXyChartXAxisScaleEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartXAxisScaleEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartXAxisScaleEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartXAxisScaleEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartXAxisScaleEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartXAxisScaleEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartXAxisScaleEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartXAxisScaleEnumSlice flattens the contents of DashboardWidgetXyChartXAxisScaleEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartXAxisScaleEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartXAxisScaleEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartXAxisScaleEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartXAxisScaleEnum{}
	}

	items := make([]DashboardWidgetXyChartXAxisScaleEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartXAxisScaleEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartXAxisScaleEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartXAxisScaleEnum with the same value as that string.
func flattenDashboardWidgetXyChartXAxisScaleEnum(i interface{}) *DashboardWidgetXyChartXAxisScaleEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartXAxisScaleEnumRef("")
	}

	return DashboardWidgetXyChartXAxisScaleEnumRef(s)
}

// flattenDashboardWidgetXyChartYAxisScaleEnumMap flattens the contents of DashboardWidgetXyChartYAxisScaleEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartYAxisScaleEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartYAxisScaleEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartYAxisScaleEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartYAxisScaleEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartYAxisScaleEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartYAxisScaleEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartYAxisScaleEnumSlice flattens the contents of DashboardWidgetXyChartYAxisScaleEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartYAxisScaleEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartYAxisScaleEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartYAxisScaleEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartYAxisScaleEnum{}
	}

	items := make([]DashboardWidgetXyChartYAxisScaleEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartYAxisScaleEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartYAxisScaleEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartYAxisScaleEnum with the same value as that string.
func flattenDashboardWidgetXyChartYAxisScaleEnum(i interface{}) *DashboardWidgetXyChartYAxisScaleEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartYAxisScaleEnumRef("")
	}

	return DashboardWidgetXyChartYAxisScaleEnumRef(s)
}

// flattenDashboardWidgetXyChartChartOptionsModeEnumMap flattens the contents of DashboardWidgetXyChartChartOptionsModeEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartChartOptionsModeEnumMap(c *Client, i interface{}) map[string]DashboardWidgetXyChartChartOptionsModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetXyChartChartOptionsModeEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetXyChartChartOptionsModeEnum{}
	}

	items := make(map[string]DashboardWidgetXyChartChartOptionsModeEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetXyChartChartOptionsModeEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetXyChartChartOptionsModeEnumSlice flattens the contents of DashboardWidgetXyChartChartOptionsModeEnum from a JSON
// response object.
func flattenDashboardWidgetXyChartChartOptionsModeEnumSlice(c *Client, i interface{}) []DashboardWidgetXyChartChartOptionsModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetXyChartChartOptionsModeEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetXyChartChartOptionsModeEnum{}
	}

	items := make([]DashboardWidgetXyChartChartOptionsModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetXyChartChartOptionsModeEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetXyChartChartOptionsModeEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetXyChartChartOptionsModeEnum with the same value as that string.
func flattenDashboardWidgetXyChartChartOptionsModeEnum(i interface{}) *DashboardWidgetXyChartChartOptionsModeEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetXyChartChartOptionsModeEnumRef("")
	}

	return DashboardWidgetXyChartChartOptionsModeEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumMap flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumSlice flattens the contents of DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum{}
	}

	items := make([]DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum with the same value as that string.
func flattenDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(i interface{}) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef("")
	}

	return DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
}

// flattenDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumMap flattens the contents of DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumSlice flattens the contents of DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum{}
	}

	items := make([]DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum with the same value as that string.
func flattenDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(i interface{}) *DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumRef("")
	}

	return DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumRef(s)
}

// flattenDashboardWidgetScorecardThresholdsColorEnumMap flattens the contents of DashboardWidgetScorecardThresholdsColorEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardThresholdsColorEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardThresholdsColorEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardThresholdsColorEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardThresholdsColorEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardThresholdsColorEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardThresholdsColorEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardThresholdsColorEnumSlice flattens the contents of DashboardWidgetScorecardThresholdsColorEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardThresholdsColorEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardThresholdsColorEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardThresholdsColorEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardThresholdsColorEnum{}
	}

	items := make([]DashboardWidgetScorecardThresholdsColorEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardThresholdsColorEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardThresholdsColorEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardThresholdsColorEnum with the same value as that string.
func flattenDashboardWidgetScorecardThresholdsColorEnum(i interface{}) *DashboardWidgetScorecardThresholdsColorEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardThresholdsColorEnumRef("")
	}

	return DashboardWidgetScorecardThresholdsColorEnumRef(s)
}

// flattenDashboardWidgetScorecardThresholdsDirectionEnumMap flattens the contents of DashboardWidgetScorecardThresholdsDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardThresholdsDirectionEnumMap(c *Client, i interface{}) map[string]DashboardWidgetScorecardThresholdsDirectionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetScorecardThresholdsDirectionEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetScorecardThresholdsDirectionEnum{}
	}

	items := make(map[string]DashboardWidgetScorecardThresholdsDirectionEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetScorecardThresholdsDirectionEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetScorecardThresholdsDirectionEnumSlice flattens the contents of DashboardWidgetScorecardThresholdsDirectionEnum from a JSON
// response object.
func flattenDashboardWidgetScorecardThresholdsDirectionEnumSlice(c *Client, i interface{}) []DashboardWidgetScorecardThresholdsDirectionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetScorecardThresholdsDirectionEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetScorecardThresholdsDirectionEnum{}
	}

	items := make([]DashboardWidgetScorecardThresholdsDirectionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetScorecardThresholdsDirectionEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetScorecardThresholdsDirectionEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetScorecardThresholdsDirectionEnum with the same value as that string.
func flattenDashboardWidgetScorecardThresholdsDirectionEnum(i interface{}) *DashboardWidgetScorecardThresholdsDirectionEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetScorecardThresholdsDirectionEnumRef("")
	}

	return DashboardWidgetScorecardThresholdsDirectionEnumRef(s)
}

// flattenDashboardWidgetTextFormatEnumMap flattens the contents of DashboardWidgetTextFormatEnum from a JSON
// response object.
func flattenDashboardWidgetTextFormatEnumMap(c *Client, i interface{}) map[string]DashboardWidgetTextFormatEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DashboardWidgetTextFormatEnum{}
	}

	if len(a) == 0 {
		return map[string]DashboardWidgetTextFormatEnum{}
	}

	items := make(map[string]DashboardWidgetTextFormatEnum)
	for k, item := range a {
		items[k] = *flattenDashboardWidgetTextFormatEnum(item.(interface{}))
	}

	return items
}

// flattenDashboardWidgetTextFormatEnumSlice flattens the contents of DashboardWidgetTextFormatEnum from a JSON
// response object.
func flattenDashboardWidgetTextFormatEnumSlice(c *Client, i interface{}) []DashboardWidgetTextFormatEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DashboardWidgetTextFormatEnum{}
	}

	if len(a) == 0 {
		return []DashboardWidgetTextFormatEnum{}
	}

	items := make([]DashboardWidgetTextFormatEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDashboardWidgetTextFormatEnum(item.(interface{})))
	}

	return items
}

// flattenDashboardWidgetTextFormatEnum asserts that an interface is a string, and returns a
// pointer to a *DashboardWidgetTextFormatEnum with the same value as that string.
func flattenDashboardWidgetTextFormatEnum(i interface{}) *DashboardWidgetTextFormatEnum {
	s, ok := i.(string)
	if !ok {
		return DashboardWidgetTextFormatEnumRef("")
	}

	return DashboardWidgetTextFormatEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Dashboard) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalDashboard(b, c)
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

type dashboardDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         dashboardApiOperation
}

func convertFieldDiffsToDashboardDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]dashboardDiff, error) {
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
	var diffs []dashboardDiff
	// For each operation name, create a dashboardDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := dashboardDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToDashboardApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToDashboardApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (dashboardApiOperation, error) {
	switch opName {

	case "updateDashboardUpdateDashboardOperation":
		return &updateDashboardUpdateDashboardOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractDashboardFields(r *Dashboard) error {
	vGridLayout := r.GridLayout
	if vGridLayout == nil {
		// note: explicitly not the empty object.
		vGridLayout = &DashboardGridLayout{}
	}
	if err := extractDashboardGridLayoutFields(r, vGridLayout); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGridLayout) {
		r.GridLayout = vGridLayout
	}
	vMosaicLayout := r.MosaicLayout
	if vMosaicLayout == nil {
		// note: explicitly not the empty object.
		vMosaicLayout = &DashboardMosaicLayout{}
	}
	if err := extractDashboardMosaicLayoutFields(r, vMosaicLayout); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMosaicLayout) {
		r.MosaicLayout = vMosaicLayout
	}
	vRowLayout := r.RowLayout
	if vRowLayout == nil {
		// note: explicitly not the empty object.
		vRowLayout = &DashboardRowLayout{}
	}
	if err := extractDashboardRowLayoutFields(r, vRowLayout); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vRowLayout) {
		r.RowLayout = vRowLayout
	}
	vColumnLayout := r.ColumnLayout
	if vColumnLayout == nil {
		// note: explicitly not the empty object.
		vColumnLayout = &DashboardColumnLayout{}
	}
	if err := extractDashboardColumnLayoutFields(r, vColumnLayout); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vColumnLayout) {
		r.ColumnLayout = vColumnLayout
	}
	return nil
}
func extractDashboardGridLayoutFields(r *Dashboard, o *DashboardGridLayout) error {
	return nil
}
func extractDashboardMosaicLayoutFields(r *Dashboard, o *DashboardMosaicLayout) error {
	return nil
}
func extractDashboardMosaicLayoutTilesFields(r *Dashboard, o *DashboardMosaicLayoutTiles) error {
	// *DashboardWidget is a reused type - that's not compatible with function extractors.

	return nil
}
func extractDashboardRowLayoutFields(r *Dashboard, o *DashboardRowLayout) error {
	return nil
}
func extractDashboardRowLayoutRowsFields(r *Dashboard, o *DashboardRowLayoutRows) error {
	return nil
}
func extractDashboardColumnLayoutFields(r *Dashboard, o *DashboardColumnLayout) error {
	return nil
}
func extractDashboardColumnLayoutColumnsFields(r *Dashboard, o *DashboardColumnLayoutColumns) error {
	return nil
}
func extractDashboardWidgetFields(r *Dashboard, o *DashboardWidget) error {
	vXyChart := o.XyChart
	if vXyChart == nil {
		// note: explicitly not the empty object.
		vXyChart = &DashboardWidgetXyChart{}
	}
	if err := extractDashboardWidgetXyChartFields(r, vXyChart); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vXyChart) {
		o.XyChart = vXyChart
	}
	vScorecard := o.Scorecard
	if vScorecard == nil {
		// note: explicitly not the empty object.
		vScorecard = &DashboardWidgetScorecard{}
	}
	if err := extractDashboardWidgetScorecardFields(r, vScorecard); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vScorecard) {
		o.Scorecard = vScorecard
	}
	vText := o.Text
	if vText == nil {
		// note: explicitly not the empty object.
		vText = &DashboardWidgetText{}
	}
	if err := extractDashboardWidgetTextFields(r, vText); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vText) {
		o.Text = vText
	}
	vBlank := o.Blank
	if vBlank == nil {
		// note: explicitly not the empty object.
		vBlank = &DashboardWidgetBlank{}
	}
	if err := extractDashboardWidgetBlankFields(r, vBlank); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vBlank) {
		o.Blank = vBlank
	}
	return nil
}
func extractDashboardWidgetXyChartFields(r *Dashboard, o *DashboardWidgetXyChart) error {
	vXAxis := o.XAxis
	if vXAxis == nil {
		// note: explicitly not the empty object.
		vXAxis = &DashboardWidgetXyChartXAxis{}
	}
	if err := extractDashboardWidgetXyChartXAxisFields(r, vXAxis); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vXAxis) {
		o.XAxis = vXAxis
	}
	vYAxis := o.YAxis
	if vYAxis == nil {
		// note: explicitly not the empty object.
		vYAxis = &DashboardWidgetXyChartYAxis{}
	}
	if err := extractDashboardWidgetXyChartYAxisFields(r, vYAxis); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vYAxis) {
		o.YAxis = vYAxis
	}
	vChartOptions := o.ChartOptions
	if vChartOptions == nil {
		// note: explicitly not the empty object.
		vChartOptions = &DashboardWidgetXyChartChartOptions{}
	}
	if err := extractDashboardWidgetXyChartChartOptionsFields(r, vChartOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vChartOptions) {
		o.ChartOptions = vChartOptions
	}
	return nil
}
func extractDashboardWidgetXyChartDataSetsFields(r *Dashboard, o *DashboardWidgetXyChartDataSets) error {
	vTimeSeriesQuery := o.TimeSeriesQuery
	if vTimeSeriesQuery == nil {
		// note: explicitly not the empty object.
		vTimeSeriesQuery = &DashboardWidgetXyChartDataSetsTimeSeriesQuery{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryFields(r, vTimeSeriesQuery); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesQuery) {
		o.TimeSeriesQuery = vTimeSeriesQuery
	}
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQuery) error {
	vTimeSeriesFilter := o.TimeSeriesFilter
	if vTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterFields(r, vTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesFilter) {
		o.TimeSeriesFilter = vTimeSeriesFilter
	}
	vTimeSeriesFilterRatio := o.TimeSeriesFilterRatio
	if vTimeSeriesFilterRatio == nil {
		// note: explicitly not the empty object.
		vTimeSeriesFilterRatio = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioFields(r, vTimeSeriesFilterRatio); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesFilterRatio) {
		o.TimeSeriesFilterRatio = vTimeSeriesFilterRatio
	}
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	vSecondaryAggregation := o.SecondaryAggregation
	if vSecondaryAggregation == nil {
		// note: explicitly not the empty object.
		vSecondaryAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationFields(r, vSecondaryAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryAggregation) {
		o.SecondaryAggregation = vSecondaryAggregation
	}
	vPickTimeSeriesFilter := o.PickTimeSeriesFilter
	if vPickTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vPickTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterFields(r, vPickTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPickTimeSeriesFilter) {
		o.PickTimeSeriesFilter = vPickTimeSeriesFilter
	}
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) error {
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) error {
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) error {
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) error {
	vNumerator := o.Numerator
	if vNumerator == nil {
		// note: explicitly not the empty object.
		vNumerator = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorFields(r, vNumerator); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNumerator) {
		o.Numerator = vNumerator
	}
	vDenominator := o.Denominator
	if vDenominator == nil {
		// note: explicitly not the empty object.
		vDenominator = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorFields(r, vDenominator); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDenominator) {
		o.Denominator = vDenominator
	}
	vSecondaryAggregation := o.SecondaryAggregation
	if vSecondaryAggregation == nil {
		// note: explicitly not the empty object.
		vSecondaryAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationFields(r, vSecondaryAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryAggregation) {
		o.SecondaryAggregation = vSecondaryAggregation
	}
	vPickTimeSeriesFilter := o.PickTimeSeriesFilter
	if vPickTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vPickTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterFields(r, vPickTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPickTimeSeriesFilter) {
		o.PickTimeSeriesFilter = vPickTimeSeriesFilter
	}
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) error {
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) error {
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) error {
	return nil
}
func extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) error {
	return nil
}
func extractDashboardWidgetXyChartThresholdsFields(r *Dashboard, o *DashboardWidgetXyChartThresholds) error {
	return nil
}
func extractDashboardWidgetXyChartXAxisFields(r *Dashboard, o *DashboardWidgetXyChartXAxis) error {
	return nil
}
func extractDashboardWidgetXyChartYAxisFields(r *Dashboard, o *DashboardWidgetXyChartYAxis) error {
	return nil
}
func extractDashboardWidgetXyChartChartOptionsFields(r *Dashboard, o *DashboardWidgetXyChartChartOptions) error {
	return nil
}
func extractDashboardWidgetScorecardFields(r *Dashboard, o *DashboardWidgetScorecard) error {
	vTimeSeriesQuery := o.TimeSeriesQuery
	if vTimeSeriesQuery == nil {
		// note: explicitly not the empty object.
		vTimeSeriesQuery = &DashboardWidgetScorecardTimeSeriesQuery{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryFields(r, vTimeSeriesQuery); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesQuery) {
		o.TimeSeriesQuery = vTimeSeriesQuery
	}
	vGaugeView := o.GaugeView
	if vGaugeView == nil {
		// note: explicitly not the empty object.
		vGaugeView = &DashboardWidgetScorecardGaugeView{}
	}
	if err := extractDashboardWidgetScorecardGaugeViewFields(r, vGaugeView); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGaugeView) {
		o.GaugeView = vGaugeView
	}
	vSparkChartView := o.SparkChartView
	if vSparkChartView == nil {
		// note: explicitly not the empty object.
		vSparkChartView = &DashboardWidgetScorecardSparkChartView{}
	}
	if err := extractDashboardWidgetScorecardSparkChartViewFields(r, vSparkChartView); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSparkChartView) {
		o.SparkChartView = vSparkChartView
	}
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQuery) error {
	vTimeSeriesFilter := o.TimeSeriesFilter
	if vTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterFields(r, vTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesFilter) {
		o.TimeSeriesFilter = vTimeSeriesFilter
	}
	vTimeSeriesFilterRatio := o.TimeSeriesFilterRatio
	if vTimeSeriesFilterRatio == nil {
		// note: explicitly not the empty object.
		vTimeSeriesFilterRatio = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioFields(r, vTimeSeriesFilterRatio); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesFilterRatio) {
		o.TimeSeriesFilterRatio = vTimeSeriesFilterRatio
	}
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	vSecondaryAggregation := o.SecondaryAggregation
	if vSecondaryAggregation == nil {
		// note: explicitly not the empty object.
		vSecondaryAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationFields(r, vSecondaryAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryAggregation) {
		o.SecondaryAggregation = vSecondaryAggregation
	}
	vPickTimeSeriesFilter := o.PickTimeSeriesFilter
	if vPickTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vPickTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterFields(r, vPickTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPickTimeSeriesFilter) {
		o.PickTimeSeriesFilter = vPickTimeSeriesFilter
	}
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) error {
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) error {
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) error {
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) error {
	vNumerator := o.Numerator
	if vNumerator == nil {
		// note: explicitly not the empty object.
		vNumerator = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorFields(r, vNumerator); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNumerator) {
		o.Numerator = vNumerator
	}
	vDenominator := o.Denominator
	if vDenominator == nil {
		// note: explicitly not the empty object.
		vDenominator = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorFields(r, vDenominator); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDenominator) {
		o.Denominator = vDenominator
	}
	vSecondaryAggregation := o.SecondaryAggregation
	if vSecondaryAggregation == nil {
		// note: explicitly not the empty object.
		vSecondaryAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationFields(r, vSecondaryAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryAggregation) {
		o.SecondaryAggregation = vSecondaryAggregation
	}
	vPickTimeSeriesFilter := o.PickTimeSeriesFilter
	if vPickTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vPickTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterFields(r, vPickTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPickTimeSeriesFilter) {
		o.PickTimeSeriesFilter = vPickTimeSeriesFilter
	}
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) error {
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) error {
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) error {
	return nil
}
func extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) error {
	return nil
}
func extractDashboardWidgetScorecardGaugeViewFields(r *Dashboard, o *DashboardWidgetScorecardGaugeView) error {
	return nil
}
func extractDashboardWidgetScorecardSparkChartViewFields(r *Dashboard, o *DashboardWidgetScorecardSparkChartView) error {
	return nil
}
func extractDashboardWidgetScorecardThresholdsFields(r *Dashboard, o *DashboardWidgetScorecardThresholds) error {
	return nil
}
func extractDashboardWidgetTextFields(r *Dashboard, o *DashboardWidgetText) error {
	return nil
}
func extractDashboardWidgetBlankFields(r *Dashboard, o *DashboardWidgetBlank) error {
	return nil
}

func postReadExtractDashboardFields(r *Dashboard) error {
	vGridLayout := r.GridLayout
	if vGridLayout == nil {
		// note: explicitly not the empty object.
		vGridLayout = &DashboardGridLayout{}
	}
	if err := postReadExtractDashboardGridLayoutFields(r, vGridLayout); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGridLayout) {
		r.GridLayout = vGridLayout
	}
	vMosaicLayout := r.MosaicLayout
	if vMosaicLayout == nil {
		// note: explicitly not the empty object.
		vMosaicLayout = &DashboardMosaicLayout{}
	}
	if err := postReadExtractDashboardMosaicLayoutFields(r, vMosaicLayout); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMosaicLayout) {
		r.MosaicLayout = vMosaicLayout
	}
	vRowLayout := r.RowLayout
	if vRowLayout == nil {
		// note: explicitly not the empty object.
		vRowLayout = &DashboardRowLayout{}
	}
	if err := postReadExtractDashboardRowLayoutFields(r, vRowLayout); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vRowLayout) {
		r.RowLayout = vRowLayout
	}
	vColumnLayout := r.ColumnLayout
	if vColumnLayout == nil {
		// note: explicitly not the empty object.
		vColumnLayout = &DashboardColumnLayout{}
	}
	if err := postReadExtractDashboardColumnLayoutFields(r, vColumnLayout); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vColumnLayout) {
		r.ColumnLayout = vColumnLayout
	}
	return nil
}
func postReadExtractDashboardGridLayoutFields(r *Dashboard, o *DashboardGridLayout) error {
	return nil
}
func postReadExtractDashboardMosaicLayoutFields(r *Dashboard, o *DashboardMosaicLayout) error {
	return nil
}
func postReadExtractDashboardMosaicLayoutTilesFields(r *Dashboard, o *DashboardMosaicLayoutTiles) error {
	// *DashboardWidget is a reused type - that's not compatible with function extractors.

	return nil
}
func postReadExtractDashboardRowLayoutFields(r *Dashboard, o *DashboardRowLayout) error {
	return nil
}
func postReadExtractDashboardRowLayoutRowsFields(r *Dashboard, o *DashboardRowLayoutRows) error {
	return nil
}
func postReadExtractDashboardColumnLayoutFields(r *Dashboard, o *DashboardColumnLayout) error {
	return nil
}
func postReadExtractDashboardColumnLayoutColumnsFields(r *Dashboard, o *DashboardColumnLayoutColumns) error {
	return nil
}
func postReadExtractDashboardWidgetFields(r *Dashboard, o *DashboardWidget) error {
	vXyChart := o.XyChart
	if vXyChart == nil {
		// note: explicitly not the empty object.
		vXyChart = &DashboardWidgetXyChart{}
	}
	if err := extractDashboardWidgetXyChartFields(r, vXyChart); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vXyChart) {
		o.XyChart = vXyChart
	}
	vScorecard := o.Scorecard
	if vScorecard == nil {
		// note: explicitly not the empty object.
		vScorecard = &DashboardWidgetScorecard{}
	}
	if err := extractDashboardWidgetScorecardFields(r, vScorecard); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vScorecard) {
		o.Scorecard = vScorecard
	}
	vText := o.Text
	if vText == nil {
		// note: explicitly not the empty object.
		vText = &DashboardWidgetText{}
	}
	if err := extractDashboardWidgetTextFields(r, vText); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vText) {
		o.Text = vText
	}
	vBlank := o.Blank
	if vBlank == nil {
		// note: explicitly not the empty object.
		vBlank = &DashboardWidgetBlank{}
	}
	if err := extractDashboardWidgetBlankFields(r, vBlank); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vBlank) {
		o.Blank = vBlank
	}
	return nil
}
func postReadExtractDashboardWidgetXyChartFields(r *Dashboard, o *DashboardWidgetXyChart) error {
	vXAxis := o.XAxis
	if vXAxis == nil {
		// note: explicitly not the empty object.
		vXAxis = &DashboardWidgetXyChartXAxis{}
	}
	if err := extractDashboardWidgetXyChartXAxisFields(r, vXAxis); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vXAxis) {
		o.XAxis = vXAxis
	}
	vYAxis := o.YAxis
	if vYAxis == nil {
		// note: explicitly not the empty object.
		vYAxis = &DashboardWidgetXyChartYAxis{}
	}
	if err := extractDashboardWidgetXyChartYAxisFields(r, vYAxis); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vYAxis) {
		o.YAxis = vYAxis
	}
	vChartOptions := o.ChartOptions
	if vChartOptions == nil {
		// note: explicitly not the empty object.
		vChartOptions = &DashboardWidgetXyChartChartOptions{}
	}
	if err := extractDashboardWidgetXyChartChartOptionsFields(r, vChartOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vChartOptions) {
		o.ChartOptions = vChartOptions
	}
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsFields(r *Dashboard, o *DashboardWidgetXyChartDataSets) error {
	vTimeSeriesQuery := o.TimeSeriesQuery
	if vTimeSeriesQuery == nil {
		// note: explicitly not the empty object.
		vTimeSeriesQuery = &DashboardWidgetXyChartDataSetsTimeSeriesQuery{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryFields(r, vTimeSeriesQuery); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesQuery) {
		o.TimeSeriesQuery = vTimeSeriesQuery
	}
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQuery) error {
	vTimeSeriesFilter := o.TimeSeriesFilter
	if vTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterFields(r, vTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesFilter) {
		o.TimeSeriesFilter = vTimeSeriesFilter
	}
	vTimeSeriesFilterRatio := o.TimeSeriesFilterRatio
	if vTimeSeriesFilterRatio == nil {
		// note: explicitly not the empty object.
		vTimeSeriesFilterRatio = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioFields(r, vTimeSeriesFilterRatio); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesFilterRatio) {
		o.TimeSeriesFilterRatio = vTimeSeriesFilterRatio
	}
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	vSecondaryAggregation := o.SecondaryAggregation
	if vSecondaryAggregation == nil {
		// note: explicitly not the empty object.
		vSecondaryAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationFields(r, vSecondaryAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryAggregation) {
		o.SecondaryAggregation = vSecondaryAggregation
	}
	vPickTimeSeriesFilter := o.PickTimeSeriesFilter
	if vPickTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vPickTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterFields(r, vPickTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPickTimeSeriesFilter) {
		o.PickTimeSeriesFilter = vPickTimeSeriesFilter
	}
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) error {
	vNumerator := o.Numerator
	if vNumerator == nil {
		// note: explicitly not the empty object.
		vNumerator = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorFields(r, vNumerator); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNumerator) {
		o.Numerator = vNumerator
	}
	vDenominator := o.Denominator
	if vDenominator == nil {
		// note: explicitly not the empty object.
		vDenominator = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorFields(r, vDenominator); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDenominator) {
		o.Denominator = vDenominator
	}
	vSecondaryAggregation := o.SecondaryAggregation
	if vSecondaryAggregation == nil {
		// note: explicitly not the empty object.
		vSecondaryAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationFields(r, vSecondaryAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryAggregation) {
		o.SecondaryAggregation = vSecondaryAggregation
	}
	vPickTimeSeriesFilter := o.PickTimeSeriesFilter
	if vPickTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vPickTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterFields(r, vPickTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPickTimeSeriesFilter) {
		o.PickTimeSeriesFilter = vPickTimeSeriesFilter
	}
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}
	if err := extractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartThresholdsFields(r *Dashboard, o *DashboardWidgetXyChartThresholds) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartXAxisFields(r *Dashboard, o *DashboardWidgetXyChartXAxis) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartYAxisFields(r *Dashboard, o *DashboardWidgetXyChartYAxis) error {
	return nil
}
func postReadExtractDashboardWidgetXyChartChartOptionsFields(r *Dashboard, o *DashboardWidgetXyChartChartOptions) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardFields(r *Dashboard, o *DashboardWidgetScorecard) error {
	vTimeSeriesQuery := o.TimeSeriesQuery
	if vTimeSeriesQuery == nil {
		// note: explicitly not the empty object.
		vTimeSeriesQuery = &DashboardWidgetScorecardTimeSeriesQuery{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryFields(r, vTimeSeriesQuery); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesQuery) {
		o.TimeSeriesQuery = vTimeSeriesQuery
	}
	vGaugeView := o.GaugeView
	if vGaugeView == nil {
		// note: explicitly not the empty object.
		vGaugeView = &DashboardWidgetScorecardGaugeView{}
	}
	if err := extractDashboardWidgetScorecardGaugeViewFields(r, vGaugeView); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGaugeView) {
		o.GaugeView = vGaugeView
	}
	vSparkChartView := o.SparkChartView
	if vSparkChartView == nil {
		// note: explicitly not the empty object.
		vSparkChartView = &DashboardWidgetScorecardSparkChartView{}
	}
	if err := extractDashboardWidgetScorecardSparkChartViewFields(r, vSparkChartView); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSparkChartView) {
		o.SparkChartView = vSparkChartView
	}
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQuery) error {
	vTimeSeriesFilter := o.TimeSeriesFilter
	if vTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterFields(r, vTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesFilter) {
		o.TimeSeriesFilter = vTimeSeriesFilter
	}
	vTimeSeriesFilterRatio := o.TimeSeriesFilterRatio
	if vTimeSeriesFilterRatio == nil {
		// note: explicitly not the empty object.
		vTimeSeriesFilterRatio = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioFields(r, vTimeSeriesFilterRatio); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimeSeriesFilterRatio) {
		o.TimeSeriesFilterRatio = vTimeSeriesFilterRatio
	}
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	vSecondaryAggregation := o.SecondaryAggregation
	if vSecondaryAggregation == nil {
		// note: explicitly not the empty object.
		vSecondaryAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationFields(r, vSecondaryAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryAggregation) {
		o.SecondaryAggregation = vSecondaryAggregation
	}
	vPickTimeSeriesFilter := o.PickTimeSeriesFilter
	if vPickTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vPickTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterFields(r, vPickTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPickTimeSeriesFilter) {
		o.PickTimeSeriesFilter = vPickTimeSeriesFilter
	}
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) error {
	vNumerator := o.Numerator
	if vNumerator == nil {
		// note: explicitly not the empty object.
		vNumerator = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorFields(r, vNumerator); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNumerator) {
		o.Numerator = vNumerator
	}
	vDenominator := o.Denominator
	if vDenominator == nil {
		// note: explicitly not the empty object.
		vDenominator = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorFields(r, vDenominator); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDenominator) {
		o.Denominator = vDenominator
	}
	vSecondaryAggregation := o.SecondaryAggregation
	if vSecondaryAggregation == nil {
		// note: explicitly not the empty object.
		vSecondaryAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationFields(r, vSecondaryAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSecondaryAggregation) {
		o.SecondaryAggregation = vSecondaryAggregation
	}
	vPickTimeSeriesFilter := o.PickTimeSeriesFilter
	if vPickTimeSeriesFilter == nil {
		// note: explicitly not the empty object.
		vPickTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterFields(r, vPickTimeSeriesFilter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPickTimeSeriesFilter) {
		o.PickTimeSeriesFilter = vPickTimeSeriesFilter
	}
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) error {
	vAggregation := o.Aggregation
	if vAggregation == nil {
		// note: explicitly not the empty object.
		vAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
	}
	if err := extractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationFields(r, vAggregation); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAggregation) {
		o.Aggregation = vAggregation
	}
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterFields(r *Dashboard, o *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardGaugeViewFields(r *Dashboard, o *DashboardWidgetScorecardGaugeView) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardSparkChartViewFields(r *Dashboard, o *DashboardWidgetScorecardSparkChartView) error {
	return nil
}
func postReadExtractDashboardWidgetScorecardThresholdsFields(r *Dashboard, o *DashboardWidgetScorecardThresholds) error {
	return nil
}
func postReadExtractDashboardWidgetTextFields(r *Dashboard, o *DashboardWidgetText) error {
	return nil
}
func postReadExtractDashboardWidgetBlankFields(r *Dashboard, o *DashboardWidgetBlank) error {
	return nil
}
