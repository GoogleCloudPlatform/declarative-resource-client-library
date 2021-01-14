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
package logging

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

func (r *LogMetric) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MetricDescriptor) {
		if err := r.MetricDescriptor.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BucketOptions) {
		if err := r.BucketOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *LogMetricMetricDescriptor) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Metadata) {
		if err := r.Metadata.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *LogMetricMetricDescriptorLabels) validate() error {
	return nil
}
func (r *LogMetricMetricDescriptorMetadata) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SamplePeriod) {
		if err := r.SamplePeriod.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.IngestDelay) {
		if err := r.IngestDelay.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *LogMetricMetricDescriptorMetadataSamplePeriod) validate() error {
	return nil
}
func (r *LogMetricMetricDescriptorMetadataIngestDelay) validate() error {
	return nil
}
func (r *LogMetricBucketOptions) validate() error {
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
func (r *LogMetricBucketOptionsLinearBuckets) validate() error {
	return nil
}
func (r *LogMetricBucketOptionsExponentialBuckets) validate() error {
	return nil
}
func (r *LogMetricBucketOptionsExplicitBuckets) validate() error {
	return nil
}

func logMetricGetURL(userBasePath string, r *LogMetric) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/metrics/{{name}}", "https://logging.googleapis.com/v2/", userBasePath, params), nil
}

func logMetricListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/metrics", "https://logging.googleapis.com/v2/", userBasePath, params), nil

}

func logMetricCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/metrics", "https://logging.googleapis.com/v2/", userBasePath, params), nil

}

func logMetricDeleteURL(userBasePath string, r *LogMetric) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/metrics/{{name}}", "https://logging.googleapis.com/v2/", userBasePath, params), nil
}

// logMetricApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type logMetricApiOperation interface {
	do(context.Context, *LogMetric, *Client) error
}

// newUpdateLogMetricUpdateRequest creates a request for an
// LogMetric resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateLogMetricUpdateRequest(ctx context.Context, f *LogMetric, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		req["filter"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		req["disabled"] = v
	}
	if v, err := expandLogMetricMetricDescriptor(c, f.MetricDescriptor); err != nil {
		return nil, fmt.Errorf("error expanding MetricDescriptor into metricDescriptor: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["metricDescriptor"] = v
	}
	if v := f.ValueExtractor; !dcl.IsEmptyValueIndirect(v) {
		req["valueExtractor"] = v
	}
	if v := f.LabelExtractors; !dcl.IsEmptyValueIndirect(v) {
		req["labelExtractors"] = v
	}
	if v, err := expandLogMetricBucketOptions(c, f.BucketOptions); err != nil {
		return nil, fmt.Errorf("error expanding BucketOptions into bucketOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["bucketOptions"] = v
	}
	if v := f.Resolution; !dcl.IsEmptyValueIndirect(v) {
		req["resolution"] = v
	}
	return req, nil
}

// marshalUpdateLogMetricUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateLogMetricUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateLogMetricUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateLogMetricUpdateOperation) do(ctx context.Context, r *LogMetric, c *Client) error {
	_, err := c.GetLogMetric(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateLogMetricUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateLogMetricUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listLogMetricRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := logMetricListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != LogMetricMaxPage {
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

type listLogMetricOperation struct {
	Metrics []map[string]interface{} `json:"metrics"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listLogMetric(ctx context.Context, project, pageToken string, pageSize int32) ([]*LogMetric, string, error) {
	b, err := c.listLogMetricRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listLogMetricOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*LogMetric
	for _, v := range m.Metrics {
		res := flattenLogMetric(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllLogMetric(ctx context.Context, f func(*LogMetric) bool, resources []*LogMetric) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteLogMetric(ctx, res)
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

type deleteLogMetricOperation struct{}

func (op *deleteLogMetricOperation) do(ctx context.Context, r *LogMetric, c *Client) error {

	_, err := c.GetLogMetric(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("LogMetric not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetLogMetric checking for existence. error: %v", err)
		return err
	}

	u, err := logMetricDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return fmt.Errorf("failed to delete LogMetric: %w", err)
	}
	_, err = c.GetLogMetric(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createLogMetricOperation struct{}

func (op *createLogMetricOperation) do(ctx context.Context, r *LogMetric, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := logMetricCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetLogMetric(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getLogMetricRaw(ctx context.Context, r *LogMetric) ([]byte, error) {

	u, err := logMetricGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) logMetricDiffsForRawDesired(ctx context.Context, rawDesired *LogMetric, opts ...dcl.ApplyOption) (initial, desired *LogMetric, diffs []logMetricDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *LogMetric
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*LogMetric); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected LogMetric, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetLogMetric(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a LogMetric resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve LogMetric resource: %v", err)
		}

		c.Config.Logger.Info("Found that LogMetric resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeLogMetricDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for LogMetric: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for LogMetric: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeLogMetricInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for LogMetric: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeLogMetricDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for LogMetric: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffLogMetric(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeLogMetricInitialState(rawInitial, rawDesired *LogMetric) (*LogMetric, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeLogMetricDesiredState(rawDesired, rawInitial *LogMetric, opts ...dcl.ApplyOption) (*LogMetric, error) {

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*LogMetric); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected LogMetric, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MetricDescriptor = canonicalizeLogMetricMetricDescriptor(rawDesired.MetricDescriptor, nil, opts...)
		rawDesired.BucketOptions = canonicalizeLogMetricBucketOptions(rawDesired.BucketOptions, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Filter) {
		rawDesired.Filter = rawInitial.Filter
	}
	if dcl.IsZeroValue(rawDesired.Disabled) {
		rawDesired.Disabled = rawInitial.Disabled
	}
	rawDesired.MetricDescriptor = canonicalizeLogMetricMetricDescriptor(rawDesired.MetricDescriptor, rawInitial.MetricDescriptor, opts...)
	if dcl.IsZeroValue(rawDesired.ValueExtractor) {
		rawDesired.ValueExtractor = rawInitial.ValueExtractor
	}
	if dcl.IsZeroValue(rawDesired.LabelExtractors) {
		rawDesired.LabelExtractors = rawInitial.LabelExtractors
	}
	rawDesired.BucketOptions = canonicalizeLogMetricBucketOptions(rawDesired.BucketOptions, rawInitial.BucketOptions, opts...)
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.IsZeroValue(rawDesired.Resolution) {
		rawDesired.Resolution = rawInitial.Resolution
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeLogMetricNewState(c *Client, rawNew, rawDesired *LogMetric) (*LogMetric, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Filter) && dcl.IsEmptyValueIndirect(rawDesired.Filter) {
		rawNew.Filter = rawDesired.Filter
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MetricDescriptor) && dcl.IsEmptyValueIndirect(rawDesired.MetricDescriptor) {
		rawNew.MetricDescriptor = rawDesired.MetricDescriptor
	} else {
		rawNew.MetricDescriptor = canonicalizeNewLogMetricMetricDescriptor(c, rawDesired.MetricDescriptor, rawNew.MetricDescriptor)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ValueExtractor) && dcl.IsEmptyValueIndirect(rawDesired.ValueExtractor) {
		rawNew.ValueExtractor = rawDesired.ValueExtractor
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LabelExtractors) && dcl.IsEmptyValueIndirect(rawDesired.LabelExtractors) {
		rawNew.LabelExtractors = rawDesired.LabelExtractors
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BucketOptions) && dcl.IsEmptyValueIndirect(rawDesired.BucketOptions) {
		rawNew.BucketOptions = rawDesired.BucketOptions
	} else {
		rawNew.BucketOptions = canonicalizeNewLogMetricBucketOptions(c, rawDesired.BucketOptions, rawNew.BucketOptions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Resolution) && dcl.IsEmptyValueIndirect(rawDesired.Resolution) {
		rawNew.Resolution = rawDesired.Resolution
	} else {
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

func canonicalizeLogMetricMetricDescriptor(des, initial *LogMetricMetricDescriptor, opts ...dcl.ApplyOption) *LogMetricMetricDescriptor {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*LogMetric)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.MetricKind) {
		des.MetricKind = initial.MetricKind
	}
	if dcl.IsZeroValue(des.ValueType) {
		des.ValueType = initial.ValueType
	}
	if dcl.IsZeroValue(des.Unit) {
		des.Unit = initial.Unit
	}
	if dcl.IsZeroValue(des.DisplayName) {
		des.DisplayName = initial.DisplayName
	}
	des.Metadata = canonicalizeLogMetricMetricDescriptorMetadata(des.Metadata, initial.Metadata, opts...)
	if dcl.IsZeroValue(des.LaunchStage) {
		des.LaunchStage = initial.LaunchStage
	}
	if dcl.IsZeroValue(des.MonitoredResourceTypes) {
		des.MonitoredResourceTypes = initial.MonitoredResourceTypes
	}

	return des
}

func canonicalizeNewLogMetricMetricDescriptor(c *Client, des, nw *LogMetricMetricDescriptor) *LogMetricMetricDescriptor {
	if des == nil || nw == nil {
		return nw
	}

	nw.Metadata = canonicalizeNewLogMetricMetricDescriptorMetadata(c, des.Metadata, nw.Metadata)

	return nw
}

func canonicalizeNewLogMetricMetricDescriptorSet(c *Client, des, nw []LogMetricMetricDescriptor) []LogMetricMetricDescriptor {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricMetricDescriptor
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareLogMetricMetricDescriptor(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeLogMetricMetricDescriptorLabels(des, initial *LogMetricMetricDescriptorLabels, opts ...dcl.ApplyOption) *LogMetricMetricDescriptorLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*LogMetric)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.ValueType) {
		des.ValueType = initial.ValueType
	}
	if dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}

	return des
}

func canonicalizeNewLogMetricMetricDescriptorLabels(c *Client, des, nw *LogMetricMetricDescriptorLabels) *LogMetricMetricDescriptorLabels {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewLogMetricMetricDescriptorLabelsSet(c *Client, des, nw []LogMetricMetricDescriptorLabels) []LogMetricMetricDescriptorLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricMetricDescriptorLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareLogMetricMetricDescriptorLabels(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeLogMetricMetricDescriptorMetadata(des, initial *LogMetricMetricDescriptorMetadata, opts ...dcl.ApplyOption) *LogMetricMetricDescriptorMetadata {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*LogMetric)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.LaunchStage) {
		des.LaunchStage = initial.LaunchStage
	}
	des.SamplePeriod = canonicalizeLogMetricMetricDescriptorMetadataSamplePeriod(des.SamplePeriod, initial.SamplePeriod, opts...)
	des.IngestDelay = canonicalizeLogMetricMetricDescriptorMetadataIngestDelay(des.IngestDelay, initial.IngestDelay, opts...)

	return des
}

func canonicalizeNewLogMetricMetricDescriptorMetadata(c *Client, des, nw *LogMetricMetricDescriptorMetadata) *LogMetricMetricDescriptorMetadata {
	if des == nil || nw == nil {
		return nw
	}

	nw.SamplePeriod = canonicalizeNewLogMetricMetricDescriptorMetadataSamplePeriod(c, des.SamplePeriod, nw.SamplePeriod)
	nw.IngestDelay = canonicalizeNewLogMetricMetricDescriptorMetadataIngestDelay(c, des.IngestDelay, nw.IngestDelay)

	return nw
}

func canonicalizeNewLogMetricMetricDescriptorMetadataSet(c *Client, des, nw []LogMetricMetricDescriptorMetadata) []LogMetricMetricDescriptorMetadata {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricMetricDescriptorMetadata
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareLogMetricMetricDescriptorMetadata(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeLogMetricMetricDescriptorMetadataSamplePeriod(des, initial *LogMetricMetricDescriptorMetadataSamplePeriod, opts ...dcl.ApplyOption) *LogMetricMetricDescriptorMetadataSamplePeriod {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*LogMetric)
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

func canonicalizeNewLogMetricMetricDescriptorMetadataSamplePeriod(c *Client, des, nw *LogMetricMetricDescriptorMetadataSamplePeriod) *LogMetricMetricDescriptorMetadataSamplePeriod {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewLogMetricMetricDescriptorMetadataSamplePeriodSet(c *Client, des, nw []LogMetricMetricDescriptorMetadataSamplePeriod) []LogMetricMetricDescriptorMetadataSamplePeriod {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricMetricDescriptorMetadataSamplePeriod
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareLogMetricMetricDescriptorMetadataSamplePeriod(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeLogMetricMetricDescriptorMetadataIngestDelay(des, initial *LogMetricMetricDescriptorMetadataIngestDelay, opts ...dcl.ApplyOption) *LogMetricMetricDescriptorMetadataIngestDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*LogMetric)
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

func canonicalizeNewLogMetricMetricDescriptorMetadataIngestDelay(c *Client, des, nw *LogMetricMetricDescriptorMetadataIngestDelay) *LogMetricMetricDescriptorMetadataIngestDelay {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewLogMetricMetricDescriptorMetadataIngestDelaySet(c *Client, des, nw []LogMetricMetricDescriptorMetadataIngestDelay) []LogMetricMetricDescriptorMetadataIngestDelay {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricMetricDescriptorMetadataIngestDelay
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareLogMetricMetricDescriptorMetadataIngestDelay(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeLogMetricBucketOptions(des, initial *LogMetricBucketOptions, opts ...dcl.ApplyOption) *LogMetricBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*LogMetric)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.LinearBuckets = canonicalizeLogMetricBucketOptionsLinearBuckets(des.LinearBuckets, initial.LinearBuckets, opts...)
	des.ExponentialBuckets = canonicalizeLogMetricBucketOptionsExponentialBuckets(des.ExponentialBuckets, initial.ExponentialBuckets, opts...)
	des.ExplicitBuckets = canonicalizeLogMetricBucketOptionsExplicitBuckets(des.ExplicitBuckets, initial.ExplicitBuckets, opts...)

	return des
}

func canonicalizeNewLogMetricBucketOptions(c *Client, des, nw *LogMetricBucketOptions) *LogMetricBucketOptions {
	if des == nil || nw == nil {
		return nw
	}

	nw.LinearBuckets = canonicalizeNewLogMetricBucketOptionsLinearBuckets(c, des.LinearBuckets, nw.LinearBuckets)
	nw.ExponentialBuckets = canonicalizeNewLogMetricBucketOptionsExponentialBuckets(c, des.ExponentialBuckets, nw.ExponentialBuckets)
	nw.ExplicitBuckets = canonicalizeNewLogMetricBucketOptionsExplicitBuckets(c, des.ExplicitBuckets, nw.ExplicitBuckets)

	return nw
}

func canonicalizeNewLogMetricBucketOptionsSet(c *Client, des, nw []LogMetricBucketOptions) []LogMetricBucketOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricBucketOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareLogMetricBucketOptions(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeLogMetricBucketOptionsLinearBuckets(des, initial *LogMetricBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *LogMetricBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*LogMetric)
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

func canonicalizeNewLogMetricBucketOptionsLinearBuckets(c *Client, des, nw *LogMetricBucketOptionsLinearBuckets) *LogMetricBucketOptionsLinearBuckets {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewLogMetricBucketOptionsLinearBucketsSet(c *Client, des, nw []LogMetricBucketOptionsLinearBuckets) []LogMetricBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricBucketOptionsLinearBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareLogMetricBucketOptionsLinearBuckets(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeLogMetricBucketOptionsExponentialBuckets(des, initial *LogMetricBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *LogMetricBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*LogMetric)
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

func canonicalizeNewLogMetricBucketOptionsExponentialBuckets(c *Client, des, nw *LogMetricBucketOptionsExponentialBuckets) *LogMetricBucketOptionsExponentialBuckets {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewLogMetricBucketOptionsExponentialBucketsSet(c *Client, des, nw []LogMetricBucketOptionsExponentialBuckets) []LogMetricBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricBucketOptionsExponentialBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareLogMetricBucketOptionsExponentialBuckets(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeLogMetricBucketOptionsExplicitBuckets(des, initial *LogMetricBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *LogMetricBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*LogMetric)
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

func canonicalizeNewLogMetricBucketOptionsExplicitBuckets(c *Client, des, nw *LogMetricBucketOptionsExplicitBuckets) *LogMetricBucketOptionsExplicitBuckets {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewLogMetricBucketOptionsExplicitBucketsSet(c *Client, des, nw []LogMetricBucketOptionsExplicitBuckets) []LogMetricBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricBucketOptionsExplicitBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareLogMetricBucketOptionsExplicitBuckets(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

type logMetricDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         logMetricApiOperation
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
func diffLogMetric(c *Client, desired, actual *LogMetric, opts ...dcl.ApplyOption) ([]logMetricDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []logMetricDiff
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %#v\nACTUAL: %#v", desired.Name, actual.Name)
		diffs = append(diffs, logMetricDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %#v\nACTUAL: %#v", desired.Description, actual.Description)

		diffs = append(diffs, logMetricDiff{
			UpdateOp:  &updateLogMetricUpdateOperation{},
			FieldName: "Description",
		})

	}
	if !dcl.IsZeroValue(desired.Filter) && (dcl.IsZeroValue(actual.Filter) || !reflect.DeepEqual(*desired.Filter, *actual.Filter)) {
		c.Config.Logger.Infof("Detected diff in Filter.\nDESIRED: %#v\nACTUAL: %#v", desired.Filter, actual.Filter)

		diffs = append(diffs, logMetricDiff{
			UpdateOp:  &updateLogMetricUpdateOperation{},
			FieldName: "Filter",
		})

	}
	if !dcl.IsZeroValue(desired.Disabled) && (dcl.IsZeroValue(actual.Disabled) || !reflect.DeepEqual(*desired.Disabled, *actual.Disabled)) {
		c.Config.Logger.Infof("Detected diff in Disabled.\nDESIRED: %#v\nACTUAL: %#v", desired.Disabled, actual.Disabled)

		diffs = append(diffs, logMetricDiff{
			UpdateOp:  &updateLogMetricUpdateOperation{},
			FieldName: "Disabled",
		})

	}
	if compareLogMetricMetricDescriptor(c, desired.MetricDescriptor, actual.MetricDescriptor) {
		c.Config.Logger.Infof("Detected diff in MetricDescriptor.\nDESIRED: %#v\nACTUAL: %#v", desired.MetricDescriptor, actual.MetricDescriptor)

		diffs = append(diffs, logMetricDiff{
			UpdateOp:  &updateLogMetricUpdateOperation{},
			FieldName: "MetricDescriptor",
		})

	}
	if !dcl.IsZeroValue(desired.ValueExtractor) && (dcl.IsZeroValue(actual.ValueExtractor) || !reflect.DeepEqual(*desired.ValueExtractor, *actual.ValueExtractor)) {
		c.Config.Logger.Infof("Detected diff in ValueExtractor.\nDESIRED: %#v\nACTUAL: %#v", desired.ValueExtractor, actual.ValueExtractor)

		diffs = append(diffs, logMetricDiff{
			UpdateOp:  &updateLogMetricUpdateOperation{},
			FieldName: "ValueExtractor",
		})

	}
	if !reflect.DeepEqual(desired.LabelExtractors, actual.LabelExtractors) {
		c.Config.Logger.Infof("Detected diff in LabelExtractors.\nDESIRED: %#v\nACTUAL: %#v", desired.LabelExtractors, actual.LabelExtractors)

		diffs = append(diffs, logMetricDiff{
			UpdateOp:  &updateLogMetricUpdateOperation{},
			FieldName: "LabelExtractors",
		})

	}
	if compareLogMetricBucketOptions(c, desired.BucketOptions, actual.BucketOptions) {
		c.Config.Logger.Infof("Detected diff in BucketOptions.\nDESIRED: %#v\nACTUAL: %#v", desired.BucketOptions, actual.BucketOptions)

		diffs = append(diffs, logMetricDiff{
			UpdateOp:  &updateLogMetricUpdateOperation{},
			FieldName: "BucketOptions",
		})

	}
	if !dcl.IsZeroValue(desired.Resolution) && (dcl.IsZeroValue(actual.Resolution) || !reflect.DeepEqual(*desired.Resolution, *actual.Resolution)) {
		c.Config.Logger.Infof("Detected diff in Resolution.\nDESIRED: %#v\nACTUAL: %#v", desired.Resolution, actual.Resolution)

		diffs = append(diffs, logMetricDiff{
			UpdateOp:  &updateLogMetricUpdateOperation{},
			FieldName: "Resolution",
		})

	}
	if !dcl.IsZeroValue(desired.Project) && !dcl.NameToSelfLink(desired.Project, actual.Project) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %#v\nACTUAL: %#v", desired.Project, actual.Project)
		diffs = append(diffs, logMetricDiff{
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
	var deduped []logMetricDiff
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
func compareLogMetricMetricDescriptorSlice(c *Client, desired, actual []LogMetricMetricDescriptor) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptor, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptor(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptor, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptor(c *Client, desired, actual *LogMetricMetricDescriptor) bool {
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
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) && !(dcl.IsEmptyValueIndirect(desired.Type) && dcl.IsZeroValue(actual.Type)) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	if actual.Labels == nil && desired.Labels != nil && !dcl.IsEmptyValueIndirect(desired.Labels) {
		c.Config.Logger.Infof("desired Labels %s - but actually nil", dcl.SprintResource(desired.Labels))
		return true
	}
	if compareLogMetricMetricDescriptorLabelsSlice(c, desired.Labels, actual.Labels) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	if actual.MetricKind == nil && desired.MetricKind != nil && !dcl.IsEmptyValueIndirect(desired.MetricKind) {
		c.Config.Logger.Infof("desired MetricKind %s - but actually nil", dcl.SprintResource(desired.MetricKind))
		return true
	}
	if !reflect.DeepEqual(desired.MetricKind, actual.MetricKind) && !dcl.IsZeroValue(desired.MetricKind) && !(dcl.IsEmptyValueIndirect(desired.MetricKind) && dcl.IsZeroValue(actual.MetricKind)) {
		c.Config.Logger.Infof("Diff in MetricKind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MetricKind), dcl.SprintResource(actual.MetricKind))
		return true
	}
	if actual.ValueType == nil && desired.ValueType != nil && !dcl.IsEmptyValueIndirect(desired.ValueType) {
		c.Config.Logger.Infof("desired ValueType %s - but actually nil", dcl.SprintResource(desired.ValueType))
		return true
	}
	if !reflect.DeepEqual(desired.ValueType, actual.ValueType) && !dcl.IsZeroValue(desired.ValueType) && !(dcl.IsEmptyValueIndirect(desired.ValueType) && dcl.IsZeroValue(actual.ValueType)) {
		c.Config.Logger.Infof("Diff in ValueType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ValueType), dcl.SprintResource(actual.ValueType))
		return true
	}
	if actual.Unit == nil && desired.Unit != nil && !dcl.IsEmptyValueIndirect(desired.Unit) {
		c.Config.Logger.Infof("desired Unit %s - but actually nil", dcl.SprintResource(desired.Unit))
		return true
	}
	if !reflect.DeepEqual(desired.Unit, actual.Unit) && !dcl.IsZeroValue(desired.Unit) && !(dcl.IsEmptyValueIndirect(desired.Unit) && dcl.IsZeroValue(actual.Unit)) {
		c.Config.Logger.Infof("Diff in Unit. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Unit), dcl.SprintResource(actual.Unit))
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
	if actual.Metadata == nil && desired.Metadata != nil && !dcl.IsEmptyValueIndirect(desired.Metadata) {
		c.Config.Logger.Infof("desired Metadata %s - but actually nil", dcl.SprintResource(desired.Metadata))
		return true
	}
	if compareLogMetricMetricDescriptorMetadata(c, desired.Metadata, actual.Metadata) && !dcl.IsZeroValue(desired.Metadata) {
		c.Config.Logger.Infof("Diff in Metadata. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Metadata), dcl.SprintResource(actual.Metadata))
		return true
	}
	if actual.LaunchStage == nil && desired.LaunchStage != nil && !dcl.IsEmptyValueIndirect(desired.LaunchStage) {
		c.Config.Logger.Infof("desired LaunchStage %s - but actually nil", dcl.SprintResource(desired.LaunchStage))
		return true
	}
	if !reflect.DeepEqual(desired.LaunchStage, actual.LaunchStage) && !dcl.IsZeroValue(desired.LaunchStage) && !(dcl.IsEmptyValueIndirect(desired.LaunchStage) && dcl.IsZeroValue(actual.LaunchStage)) {
		c.Config.Logger.Infof("Diff in LaunchStage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LaunchStage), dcl.SprintResource(actual.LaunchStage))
		return true
	}
	if actual.MonitoredResourceTypes == nil && desired.MonitoredResourceTypes != nil && !dcl.IsEmptyValueIndirect(desired.MonitoredResourceTypes) {
		c.Config.Logger.Infof("desired MonitoredResourceTypes %s - but actually nil", dcl.SprintResource(desired.MonitoredResourceTypes))
		return true
	}
	if !dcl.SliceEquals(desired.MonitoredResourceTypes, actual.MonitoredResourceTypes) && !dcl.IsZeroValue(desired.MonitoredResourceTypes) {
		c.Config.Logger.Infof("Diff in MonitoredResourceTypes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MonitoredResourceTypes), dcl.SprintResource(actual.MonitoredResourceTypes))
		return true
	}
	return false
}
func compareLogMetricMetricDescriptorLabelsSlice(c *Client, desired, actual []LogMetricMetricDescriptorLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptorLabels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptorLabels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptorLabels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptorLabels(c *Client, desired, actual *LogMetricMetricDescriptorLabels) bool {
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
	if actual.ValueType == nil && desired.ValueType != nil && !dcl.IsEmptyValueIndirect(desired.ValueType) {
		c.Config.Logger.Infof("desired ValueType %s - but actually nil", dcl.SprintResource(desired.ValueType))
		return true
	}
	if !reflect.DeepEqual(desired.ValueType, actual.ValueType) && !dcl.IsZeroValue(desired.ValueType) && !(dcl.IsEmptyValueIndirect(desired.ValueType) && dcl.IsZeroValue(actual.ValueType)) {
		c.Config.Logger.Infof("Diff in ValueType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ValueType), dcl.SprintResource(actual.ValueType))
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
	return false
}
func compareLogMetricMetricDescriptorMetadataSlice(c *Client, desired, actual []LogMetricMetricDescriptorMetadata) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptorMetadata, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptorMetadata(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptorMetadata, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptorMetadata(c *Client, desired, actual *LogMetricMetricDescriptorMetadata) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.LaunchStage == nil && desired.LaunchStage != nil && !dcl.IsEmptyValueIndirect(desired.LaunchStage) {
		c.Config.Logger.Infof("desired LaunchStage %s - but actually nil", dcl.SprintResource(desired.LaunchStage))
		return true
	}
	if !reflect.DeepEqual(desired.LaunchStage, actual.LaunchStage) && !dcl.IsZeroValue(desired.LaunchStage) && !(dcl.IsEmptyValueIndirect(desired.LaunchStage) && dcl.IsZeroValue(actual.LaunchStage)) {
		c.Config.Logger.Infof("Diff in LaunchStage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LaunchStage), dcl.SprintResource(actual.LaunchStage))
		return true
	}
	if actual.SamplePeriod == nil && desired.SamplePeriod != nil && !dcl.IsEmptyValueIndirect(desired.SamplePeriod) {
		c.Config.Logger.Infof("desired SamplePeriod %s - but actually nil", dcl.SprintResource(desired.SamplePeriod))
		return true
	}
	if compareLogMetricMetricDescriptorMetadataSamplePeriod(c, desired.SamplePeriod, actual.SamplePeriod) && !dcl.IsZeroValue(desired.SamplePeriod) {
		c.Config.Logger.Infof("Diff in SamplePeriod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SamplePeriod), dcl.SprintResource(actual.SamplePeriod))
		return true
	}
	if actual.IngestDelay == nil && desired.IngestDelay != nil && !dcl.IsEmptyValueIndirect(desired.IngestDelay) {
		c.Config.Logger.Infof("desired IngestDelay %s - but actually nil", dcl.SprintResource(desired.IngestDelay))
		return true
	}
	if compareLogMetricMetricDescriptorMetadataIngestDelay(c, desired.IngestDelay, actual.IngestDelay) && !dcl.IsZeroValue(desired.IngestDelay) {
		c.Config.Logger.Infof("Diff in IngestDelay. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IngestDelay), dcl.SprintResource(actual.IngestDelay))
		return true
	}
	return false
}
func compareLogMetricMetricDescriptorMetadataSamplePeriodSlice(c *Client, desired, actual []LogMetricMetricDescriptorMetadataSamplePeriod) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptorMetadataSamplePeriod, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptorMetadataSamplePeriod(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptorMetadataSamplePeriod, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptorMetadataSamplePeriod(c *Client, desired, actual *LogMetricMetricDescriptorMetadataSamplePeriod) bool {
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
func compareLogMetricMetricDescriptorMetadataIngestDelaySlice(c *Client, desired, actual []LogMetricMetricDescriptorMetadataIngestDelay) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptorMetadataIngestDelay, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptorMetadataIngestDelay(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptorMetadataIngestDelay, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptorMetadataIngestDelay(c *Client, desired, actual *LogMetricMetricDescriptorMetadataIngestDelay) bool {
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
func compareLogMetricBucketOptionsSlice(c *Client, desired, actual []LogMetricBucketOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricBucketOptions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricBucketOptions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricBucketOptions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricBucketOptions(c *Client, desired, actual *LogMetricBucketOptions) bool {
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
	if compareLogMetricBucketOptionsLinearBuckets(c, desired.LinearBuckets, actual.LinearBuckets) && !dcl.IsZeroValue(desired.LinearBuckets) {
		c.Config.Logger.Infof("Diff in LinearBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LinearBuckets), dcl.SprintResource(actual.LinearBuckets))
		return true
	}
	if actual.ExponentialBuckets == nil && desired.ExponentialBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("desired ExponentialBuckets %s - but actually nil", dcl.SprintResource(desired.ExponentialBuckets))
		return true
	}
	if compareLogMetricBucketOptionsExponentialBuckets(c, desired.ExponentialBuckets, actual.ExponentialBuckets) && !dcl.IsZeroValue(desired.ExponentialBuckets) {
		c.Config.Logger.Infof("Diff in ExponentialBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExponentialBuckets), dcl.SprintResource(actual.ExponentialBuckets))
		return true
	}
	if actual.ExplicitBuckets == nil && desired.ExplicitBuckets != nil && !dcl.IsEmptyValueIndirect(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("desired ExplicitBuckets %s - but actually nil", dcl.SprintResource(desired.ExplicitBuckets))
		return true
	}
	if compareLogMetricBucketOptionsExplicitBuckets(c, desired.ExplicitBuckets, actual.ExplicitBuckets) && !dcl.IsZeroValue(desired.ExplicitBuckets) {
		c.Config.Logger.Infof("Diff in ExplicitBuckets. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExplicitBuckets), dcl.SprintResource(actual.ExplicitBuckets))
		return true
	}
	return false
}
func compareLogMetricBucketOptionsLinearBucketsSlice(c *Client, desired, actual []LogMetricBucketOptionsLinearBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricBucketOptionsLinearBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricBucketOptionsLinearBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricBucketOptionsLinearBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricBucketOptionsLinearBuckets(c *Client, desired, actual *LogMetricBucketOptionsLinearBuckets) bool {
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
func compareLogMetricBucketOptionsExponentialBucketsSlice(c *Client, desired, actual []LogMetricBucketOptionsExponentialBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricBucketOptionsExponentialBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricBucketOptionsExponentialBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricBucketOptionsExponentialBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricBucketOptionsExponentialBuckets(c *Client, desired, actual *LogMetricBucketOptionsExponentialBuckets) bool {
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
func compareLogMetricBucketOptionsExplicitBucketsSlice(c *Client, desired, actual []LogMetricBucketOptionsExplicitBuckets) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricBucketOptionsExplicitBuckets, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricBucketOptionsExplicitBuckets(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricBucketOptionsExplicitBuckets, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricBucketOptionsExplicitBuckets(c *Client, desired, actual *LogMetricBucketOptionsExplicitBuckets) bool {
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
func compareLogMetricMetricDescriptorLabelsValueTypeEnumSlice(c *Client, desired, actual []LogMetricMetricDescriptorLabelsValueTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptorLabelsValueTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptorLabelsValueTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptorLabelsValueTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptorLabelsValueTypeEnum(c *Client, desired, actual *LogMetricMetricDescriptorLabelsValueTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareLogMetricMetricDescriptorMetricKindEnumSlice(c *Client, desired, actual []LogMetricMetricDescriptorMetricKindEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptorMetricKindEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptorMetricKindEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptorMetricKindEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptorMetricKindEnum(c *Client, desired, actual *LogMetricMetricDescriptorMetricKindEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareLogMetricMetricDescriptorValueTypeEnumSlice(c *Client, desired, actual []LogMetricMetricDescriptorValueTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptorValueTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptorValueTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptorValueTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptorValueTypeEnum(c *Client, desired, actual *LogMetricMetricDescriptorValueTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareLogMetricMetricDescriptorMetadataLaunchStageEnumSlice(c *Client, desired, actual []LogMetricMetricDescriptorMetadataLaunchStageEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptorMetadataLaunchStageEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptorMetadataLaunchStageEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptorMetadataLaunchStageEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptorMetadataLaunchStageEnum(c *Client, desired, actual *LogMetricMetricDescriptorMetadataLaunchStageEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareLogMetricMetricDescriptorLaunchStageEnumSlice(c *Client, desired, actual []LogMetricMetricDescriptorLaunchStageEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in LogMetricMetricDescriptorLaunchStageEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareLogMetricMetricDescriptorLaunchStageEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in LogMetricMetricDescriptorLaunchStageEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareLogMetricMetricDescriptorLaunchStageEnum(c *Client, desired, actual *LogMetricMetricDescriptorLaunchStageEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *LogMetric) urlNormalized() *LogMetric {
	normalized := deepcopy.Copy(*r).(LogMetric)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *LogMetric) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *LogMetric) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *LogMetric) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *LogMetric) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/metrics/{{name}}", "https://logging.googleapis.com/v2/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the LogMetric resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *LogMetric) marshal(c *Client) ([]byte, error) {
	m, err := expandLogMetric(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling LogMetric: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalLogMetric decodes JSON responses into the LogMetric resource schema.
func unmarshalLogMetric(b []byte, c *Client) (*LogMetric, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return flattenLogMetric(c, m), nil
}

// expandLogMetric expands LogMetric into a JSON request object.
func expandLogMetric(c *Client, f *LogMetric) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}
	if v, err := expandLogMetricMetricDescriptor(c, f.MetricDescriptor); err != nil {
		return nil, fmt.Errorf("error expanding MetricDescriptor into metricDescriptor: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metricDescriptor"] = v
	}
	if v := f.ValueExtractor; !dcl.IsEmptyValueIndirect(v) {
		m["valueExtractor"] = v
	}
	if v := f.LabelExtractors; !dcl.IsEmptyValueIndirect(v) {
		m["labelExtractors"] = v
	}
	if v, err := expandLogMetricBucketOptions(c, f.BucketOptions); err != nil {
		return nil, fmt.Errorf("error expanding BucketOptions into bucketOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bucketOptions"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Resolution; !dcl.IsEmptyValueIndirect(v) {
		m["resolution"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenLogMetric flattens LogMetric from a JSON request object into the
// LogMetric type.
func flattenLogMetric(c *Client, i interface{}) *LogMetric {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &LogMetric{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.Filter = dcl.FlattenString(m["filter"])
	r.Disabled = dcl.FlattenBool(m["disabled"])
	r.MetricDescriptor = flattenLogMetricMetricDescriptor(c, m["metricDescriptor"])
	r.ValueExtractor = dcl.FlattenString(m["valueExtractor"])
	r.LabelExtractors = dcl.FlattenKeyValuePairs(m["labelExtractors"])
	r.BucketOptions = flattenLogMetricBucketOptions(c, m["bucketOptions"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.Resolution = dcl.FlattenString(m["resolution"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandLogMetricMetricDescriptorMap expands the contents of LogMetricMetricDescriptor into a JSON
// request object.
func expandLogMetricMetricDescriptorMap(c *Client, f map[string]LogMetricMetricDescriptor) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricMetricDescriptor(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricMetricDescriptorSlice expands the contents of LogMetricMetricDescriptor into a JSON
// request object.
func expandLogMetricMetricDescriptorSlice(c *Client, f []LogMetricMetricDescriptor) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricMetricDescriptor(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricMetricDescriptorMap flattens the contents of LogMetricMetricDescriptor from a JSON
// response object.
func flattenLogMetricMetricDescriptorMap(c *Client, i interface{}) map[string]LogMetricMetricDescriptor {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricMetricDescriptor{}
	}

	if len(a) == 0 {
		return map[string]LogMetricMetricDescriptor{}
	}

	items := make(map[string]LogMetricMetricDescriptor)
	for k, item := range a {
		items[k] = *flattenLogMetricMetricDescriptor(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricMetricDescriptorSlice flattens the contents of LogMetricMetricDescriptor from a JSON
// response object.
func flattenLogMetricMetricDescriptorSlice(c *Client, i interface{}) []LogMetricMetricDescriptor {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptor{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptor{}
	}

	items := make([]LogMetricMetricDescriptor, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptor(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricMetricDescriptor expands an instance of LogMetricMetricDescriptor into a JSON
// request object.
func expandLogMetricMetricDescriptor(c *Client, f *LogMetricMetricDescriptor) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v, err := expandLogMetricMetricDescriptorLabelsSlice(c, f.Labels); err != nil {
		return nil, fmt.Errorf("error expanding Labels into labels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.MetricKind; !dcl.IsEmptyValueIndirect(v) {
		m["metricKind"] = v
	}
	if v := f.ValueType; !dcl.IsEmptyValueIndirect(v) {
		m["valueType"] = v
	}
	if v := f.Unit; !dcl.IsEmptyValueIndirect(v) {
		m["unit"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v, err := expandLogMetricMetricDescriptorMetadata(c, f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v := f.LaunchStage; !dcl.IsEmptyValueIndirect(v) {
		m["launchStage"] = v
	}
	if v := f.MonitoredResourceTypes; !dcl.IsEmptyValueIndirect(v) {
		m["monitoredResourceTypes"] = v
	}

	return m, nil
}

// flattenLogMetricMetricDescriptor flattens an instance of LogMetricMetricDescriptor from a JSON
// response object.
func flattenLogMetricMetricDescriptor(c *Client, i interface{}) *LogMetricMetricDescriptor {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricMetricDescriptor{}
	r.Name = dcl.FlattenString(m["name"])
	r.Type = dcl.FlattenString(m["type"])
	r.Labels = flattenLogMetricMetricDescriptorLabelsSlice(c, m["labels"])
	r.MetricKind = flattenLogMetricMetricDescriptorMetricKindEnum(m["metricKind"])
	r.ValueType = flattenLogMetricMetricDescriptorValueTypeEnum(m["valueType"])
	r.Unit = dcl.FlattenString(m["unit"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.Metadata = flattenLogMetricMetricDescriptorMetadata(c, m["metadata"])
	r.LaunchStage = flattenLogMetricMetricDescriptorLaunchStageEnum(m["launchStage"])
	r.MonitoredResourceTypes = dcl.FlattenStringSlice(m["monitoredResourceTypes"])

	return r
}

// expandLogMetricMetricDescriptorLabelsMap expands the contents of LogMetricMetricDescriptorLabels into a JSON
// request object.
func expandLogMetricMetricDescriptorLabelsMap(c *Client, f map[string]LogMetricMetricDescriptorLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricMetricDescriptorLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricMetricDescriptorLabelsSlice expands the contents of LogMetricMetricDescriptorLabels into a JSON
// request object.
func expandLogMetricMetricDescriptorLabelsSlice(c *Client, f []LogMetricMetricDescriptorLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricMetricDescriptorLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricMetricDescriptorLabelsMap flattens the contents of LogMetricMetricDescriptorLabels from a JSON
// response object.
func flattenLogMetricMetricDescriptorLabelsMap(c *Client, i interface{}) map[string]LogMetricMetricDescriptorLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricMetricDescriptorLabels{}
	}

	if len(a) == 0 {
		return map[string]LogMetricMetricDescriptorLabels{}
	}

	items := make(map[string]LogMetricMetricDescriptorLabels)
	for k, item := range a {
		items[k] = *flattenLogMetricMetricDescriptorLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricMetricDescriptorLabelsSlice flattens the contents of LogMetricMetricDescriptorLabels from a JSON
// response object.
func flattenLogMetricMetricDescriptorLabelsSlice(c *Client, i interface{}) []LogMetricMetricDescriptorLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorLabels{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorLabels{}
	}

	items := make([]LogMetricMetricDescriptorLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricMetricDescriptorLabels expands an instance of LogMetricMetricDescriptorLabels into a JSON
// request object.
func expandLogMetricMetricDescriptorLabels(c *Client, f *LogMetricMetricDescriptorLabels) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.ValueType; !dcl.IsEmptyValueIndirect(v) {
		m["valueType"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenLogMetricMetricDescriptorLabels flattens an instance of LogMetricMetricDescriptorLabels from a JSON
// response object.
func flattenLogMetricMetricDescriptorLabels(c *Client, i interface{}) *LogMetricMetricDescriptorLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricMetricDescriptorLabels{}
	r.Key = dcl.FlattenString(m["key"])
	r.ValueType = flattenLogMetricMetricDescriptorLabelsValueTypeEnum(m["valueType"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandLogMetricMetricDescriptorMetadataMap expands the contents of LogMetricMetricDescriptorMetadata into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataMap(c *Client, f map[string]LogMetricMetricDescriptorMetadata) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricMetricDescriptorMetadata(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricMetricDescriptorMetadataSlice expands the contents of LogMetricMetricDescriptorMetadata into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataSlice(c *Client, f []LogMetricMetricDescriptorMetadata) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricMetricDescriptorMetadata(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricMetricDescriptorMetadataMap flattens the contents of LogMetricMetricDescriptorMetadata from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataMap(c *Client, i interface{}) map[string]LogMetricMetricDescriptorMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricMetricDescriptorMetadata{}
	}

	if len(a) == 0 {
		return map[string]LogMetricMetricDescriptorMetadata{}
	}

	items := make(map[string]LogMetricMetricDescriptorMetadata)
	for k, item := range a {
		items[k] = *flattenLogMetricMetricDescriptorMetadata(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricMetricDescriptorMetadataSlice flattens the contents of LogMetricMetricDescriptorMetadata from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataSlice(c *Client, i interface{}) []LogMetricMetricDescriptorMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorMetadata{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorMetadata{}
	}

	items := make([]LogMetricMetricDescriptorMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorMetadata(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricMetricDescriptorMetadata expands an instance of LogMetricMetricDescriptorMetadata into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadata(c *Client, f *LogMetricMetricDescriptorMetadata) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.LaunchStage; !dcl.IsEmptyValueIndirect(v) {
		m["launchStage"] = v
	}
	if v, err := expandLogMetricMetricDescriptorMetadataSamplePeriod(c, f.SamplePeriod); err != nil {
		return nil, fmt.Errorf("error expanding SamplePeriod into samplePeriod: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["samplePeriod"] = v
	}
	if v, err := expandLogMetricMetricDescriptorMetadataIngestDelay(c, f.IngestDelay); err != nil {
		return nil, fmt.Errorf("error expanding IngestDelay into ingestDelay: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ingestDelay"] = v
	}

	return m, nil
}

// flattenLogMetricMetricDescriptorMetadata flattens an instance of LogMetricMetricDescriptorMetadata from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadata(c *Client, i interface{}) *LogMetricMetricDescriptorMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricMetricDescriptorMetadata{}
	r.LaunchStage = flattenLogMetricMetricDescriptorMetadataLaunchStageEnum(m["launchStage"])
	r.SamplePeriod = flattenLogMetricMetricDescriptorMetadataSamplePeriod(c, m["samplePeriod"])
	r.IngestDelay = flattenLogMetricMetricDescriptorMetadataIngestDelay(c, m["ingestDelay"])

	return r
}

// expandLogMetricMetricDescriptorMetadataSamplePeriodMap expands the contents of LogMetricMetricDescriptorMetadataSamplePeriod into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataSamplePeriodMap(c *Client, f map[string]LogMetricMetricDescriptorMetadataSamplePeriod) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricMetricDescriptorMetadataSamplePeriod(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricMetricDescriptorMetadataSamplePeriodSlice expands the contents of LogMetricMetricDescriptorMetadataSamplePeriod into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataSamplePeriodSlice(c *Client, f []LogMetricMetricDescriptorMetadataSamplePeriod) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricMetricDescriptorMetadataSamplePeriod(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricMetricDescriptorMetadataSamplePeriodMap flattens the contents of LogMetricMetricDescriptorMetadataSamplePeriod from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataSamplePeriodMap(c *Client, i interface{}) map[string]LogMetricMetricDescriptorMetadataSamplePeriod {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricMetricDescriptorMetadataSamplePeriod{}
	}

	if len(a) == 0 {
		return map[string]LogMetricMetricDescriptorMetadataSamplePeriod{}
	}

	items := make(map[string]LogMetricMetricDescriptorMetadataSamplePeriod)
	for k, item := range a {
		items[k] = *flattenLogMetricMetricDescriptorMetadataSamplePeriod(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricMetricDescriptorMetadataSamplePeriodSlice flattens the contents of LogMetricMetricDescriptorMetadataSamplePeriod from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataSamplePeriodSlice(c *Client, i interface{}) []LogMetricMetricDescriptorMetadataSamplePeriod {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorMetadataSamplePeriod{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorMetadataSamplePeriod{}
	}

	items := make([]LogMetricMetricDescriptorMetadataSamplePeriod, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorMetadataSamplePeriod(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricMetricDescriptorMetadataSamplePeriod expands an instance of LogMetricMetricDescriptorMetadataSamplePeriod into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataSamplePeriod(c *Client, f *LogMetricMetricDescriptorMetadataSamplePeriod) (map[string]interface{}, error) {
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

// flattenLogMetricMetricDescriptorMetadataSamplePeriod flattens an instance of LogMetricMetricDescriptorMetadataSamplePeriod from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataSamplePeriod(c *Client, i interface{}) *LogMetricMetricDescriptorMetadataSamplePeriod {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricMetricDescriptorMetadataSamplePeriod{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandLogMetricMetricDescriptorMetadataIngestDelayMap expands the contents of LogMetricMetricDescriptorMetadataIngestDelay into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataIngestDelayMap(c *Client, f map[string]LogMetricMetricDescriptorMetadataIngestDelay) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricMetricDescriptorMetadataIngestDelay(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricMetricDescriptorMetadataIngestDelaySlice expands the contents of LogMetricMetricDescriptorMetadataIngestDelay into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataIngestDelaySlice(c *Client, f []LogMetricMetricDescriptorMetadataIngestDelay) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricMetricDescriptorMetadataIngestDelay(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricMetricDescriptorMetadataIngestDelayMap flattens the contents of LogMetricMetricDescriptorMetadataIngestDelay from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataIngestDelayMap(c *Client, i interface{}) map[string]LogMetricMetricDescriptorMetadataIngestDelay {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricMetricDescriptorMetadataIngestDelay{}
	}

	if len(a) == 0 {
		return map[string]LogMetricMetricDescriptorMetadataIngestDelay{}
	}

	items := make(map[string]LogMetricMetricDescriptorMetadataIngestDelay)
	for k, item := range a {
		items[k] = *flattenLogMetricMetricDescriptorMetadataIngestDelay(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricMetricDescriptorMetadataIngestDelaySlice flattens the contents of LogMetricMetricDescriptorMetadataIngestDelay from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataIngestDelaySlice(c *Client, i interface{}) []LogMetricMetricDescriptorMetadataIngestDelay {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorMetadataIngestDelay{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorMetadataIngestDelay{}
	}

	items := make([]LogMetricMetricDescriptorMetadataIngestDelay, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorMetadataIngestDelay(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricMetricDescriptorMetadataIngestDelay expands an instance of LogMetricMetricDescriptorMetadataIngestDelay into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataIngestDelay(c *Client, f *LogMetricMetricDescriptorMetadataIngestDelay) (map[string]interface{}, error) {
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

// flattenLogMetricMetricDescriptorMetadataIngestDelay flattens an instance of LogMetricMetricDescriptorMetadataIngestDelay from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataIngestDelay(c *Client, i interface{}) *LogMetricMetricDescriptorMetadataIngestDelay {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricMetricDescriptorMetadataIngestDelay{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandLogMetricBucketOptionsMap expands the contents of LogMetricBucketOptions into a JSON
// request object.
func expandLogMetricBucketOptionsMap(c *Client, f map[string]LogMetricBucketOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricBucketOptionsSlice expands the contents of LogMetricBucketOptions into a JSON
// request object.
func expandLogMetricBucketOptionsSlice(c *Client, f []LogMetricBucketOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricBucketOptionsMap flattens the contents of LogMetricBucketOptions from a JSON
// response object.
func flattenLogMetricBucketOptionsMap(c *Client, i interface{}) map[string]LogMetricBucketOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricBucketOptions{}
	}

	if len(a) == 0 {
		return map[string]LogMetricBucketOptions{}
	}

	items := make(map[string]LogMetricBucketOptions)
	for k, item := range a {
		items[k] = *flattenLogMetricBucketOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricBucketOptionsSlice flattens the contents of LogMetricBucketOptions from a JSON
// response object.
func flattenLogMetricBucketOptionsSlice(c *Client, i interface{}) []LogMetricBucketOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricBucketOptions{}
	}

	if len(a) == 0 {
		return []LogMetricBucketOptions{}
	}

	items := make([]LogMetricBucketOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricBucketOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricBucketOptions expands an instance of LogMetricBucketOptions into a JSON
// request object.
func expandLogMetricBucketOptions(c *Client, f *LogMetricBucketOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandLogMetricBucketOptionsLinearBuckets(c, f.LinearBuckets); err != nil {
		return nil, fmt.Errorf("error expanding LinearBuckets into linearBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["linearBuckets"] = v
	}
	if v, err := expandLogMetricBucketOptionsExponentialBuckets(c, f.ExponentialBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExponentialBuckets into exponentialBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exponentialBuckets"] = v
	}
	if v, err := expandLogMetricBucketOptionsExplicitBuckets(c, f.ExplicitBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExplicitBuckets into explicitBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["explicitBuckets"] = v
	}

	return m, nil
}

// flattenLogMetricBucketOptions flattens an instance of LogMetricBucketOptions from a JSON
// response object.
func flattenLogMetricBucketOptions(c *Client, i interface{}) *LogMetricBucketOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricBucketOptions{}
	r.LinearBuckets = flattenLogMetricBucketOptionsLinearBuckets(c, m["linearBuckets"])
	r.ExponentialBuckets = flattenLogMetricBucketOptionsExponentialBuckets(c, m["exponentialBuckets"])
	r.ExplicitBuckets = flattenLogMetricBucketOptionsExplicitBuckets(c, m["explicitBuckets"])

	return r
}

// expandLogMetricBucketOptionsLinearBucketsMap expands the contents of LogMetricBucketOptionsLinearBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsLinearBucketsMap(c *Client, f map[string]LogMetricBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricBucketOptionsLinearBucketsSlice expands the contents of LogMetricBucketOptionsLinearBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsLinearBucketsSlice(c *Client, f []LogMetricBucketOptionsLinearBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricBucketOptionsLinearBucketsMap flattens the contents of LogMetricBucketOptionsLinearBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsLinearBucketsMap(c *Client, i interface{}) map[string]LogMetricBucketOptionsLinearBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return map[string]LogMetricBucketOptionsLinearBuckets{}
	}

	items := make(map[string]LogMetricBucketOptionsLinearBuckets)
	for k, item := range a {
		items[k] = *flattenLogMetricBucketOptionsLinearBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricBucketOptionsLinearBucketsSlice flattens the contents of LogMetricBucketOptionsLinearBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsLinearBucketsSlice(c *Client, i interface{}) []LogMetricBucketOptionsLinearBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return []LogMetricBucketOptionsLinearBuckets{}
	}

	items := make([]LogMetricBucketOptionsLinearBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricBucketOptionsLinearBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricBucketOptionsLinearBuckets expands an instance of LogMetricBucketOptionsLinearBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsLinearBuckets(c *Client, f *LogMetricBucketOptionsLinearBuckets) (map[string]interface{}, error) {
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

// flattenLogMetricBucketOptionsLinearBuckets flattens an instance of LogMetricBucketOptionsLinearBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsLinearBuckets(c *Client, i interface{}) *LogMetricBucketOptionsLinearBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricBucketOptionsLinearBuckets{}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.Width = dcl.FlattenDouble(m["width"])
	r.Offset = dcl.FlattenDouble(m["offset"])

	return r
}

// expandLogMetricBucketOptionsExponentialBucketsMap expands the contents of LogMetricBucketOptionsExponentialBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExponentialBucketsMap(c *Client, f map[string]LogMetricBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricBucketOptionsExponentialBucketsSlice expands the contents of LogMetricBucketOptionsExponentialBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExponentialBucketsSlice(c *Client, f []LogMetricBucketOptionsExponentialBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricBucketOptionsExponentialBucketsMap flattens the contents of LogMetricBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExponentialBucketsMap(c *Client, i interface{}) map[string]LogMetricBucketOptionsExponentialBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return map[string]LogMetricBucketOptionsExponentialBuckets{}
	}

	items := make(map[string]LogMetricBucketOptionsExponentialBuckets)
	for k, item := range a {
		items[k] = *flattenLogMetricBucketOptionsExponentialBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricBucketOptionsExponentialBucketsSlice flattens the contents of LogMetricBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExponentialBucketsSlice(c *Client, i interface{}) []LogMetricBucketOptionsExponentialBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return []LogMetricBucketOptionsExponentialBuckets{}
	}

	items := make([]LogMetricBucketOptionsExponentialBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricBucketOptionsExponentialBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricBucketOptionsExponentialBuckets expands an instance of LogMetricBucketOptionsExponentialBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExponentialBuckets(c *Client, f *LogMetricBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
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

// flattenLogMetricBucketOptionsExponentialBuckets flattens an instance of LogMetricBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExponentialBuckets(c *Client, i interface{}) *LogMetricBucketOptionsExponentialBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricBucketOptionsExponentialBuckets{}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.GrowthFactor = dcl.FlattenDouble(m["growthFactor"])
	r.Scale = dcl.FlattenDouble(m["scale"])

	return r
}

// expandLogMetricBucketOptionsExplicitBucketsMap expands the contents of LogMetricBucketOptionsExplicitBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExplicitBucketsMap(c *Client, f map[string]LogMetricBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricBucketOptionsExplicitBucketsSlice expands the contents of LogMetricBucketOptionsExplicitBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExplicitBucketsSlice(c *Client, f []LogMetricBucketOptionsExplicitBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricBucketOptionsExplicitBucketsMap flattens the contents of LogMetricBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExplicitBucketsMap(c *Client, i interface{}) map[string]LogMetricBucketOptionsExplicitBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return map[string]LogMetricBucketOptionsExplicitBuckets{}
	}

	items := make(map[string]LogMetricBucketOptionsExplicitBuckets)
	for k, item := range a {
		items[k] = *flattenLogMetricBucketOptionsExplicitBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricBucketOptionsExplicitBucketsSlice flattens the contents of LogMetricBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExplicitBucketsSlice(c *Client, i interface{}) []LogMetricBucketOptionsExplicitBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return []LogMetricBucketOptionsExplicitBuckets{}
	}

	items := make([]LogMetricBucketOptionsExplicitBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricBucketOptionsExplicitBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricBucketOptionsExplicitBuckets expands an instance of LogMetricBucketOptionsExplicitBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExplicitBuckets(c *Client, f *LogMetricBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Bounds; !dcl.IsEmptyValueIndirect(v) {
		m["bounds"] = v
	}

	return m, nil
}

// flattenLogMetricBucketOptionsExplicitBuckets flattens an instance of LogMetricBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExplicitBuckets(c *Client, i interface{}) *LogMetricBucketOptionsExplicitBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricBucketOptionsExplicitBuckets{}
	r.Bounds = dcl.FlattenFloatSlice(m["bounds"])

	return r
}

// flattenLogMetricMetricDescriptorLabelsValueTypeEnumSlice flattens the contents of LogMetricMetricDescriptorLabelsValueTypeEnum from a JSON
// response object.
func flattenLogMetricMetricDescriptorLabelsValueTypeEnumSlice(c *Client, i interface{}) []LogMetricMetricDescriptorLabelsValueTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorLabelsValueTypeEnum{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorLabelsValueTypeEnum{}
	}

	items := make([]LogMetricMetricDescriptorLabelsValueTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorLabelsValueTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenLogMetricMetricDescriptorLabelsValueTypeEnum asserts that an interface is a string, and returns a
// pointer to a *LogMetricMetricDescriptorLabelsValueTypeEnum with the same value as that string.
func flattenLogMetricMetricDescriptorLabelsValueTypeEnum(i interface{}) *LogMetricMetricDescriptorLabelsValueTypeEnum {
	s, ok := i.(string)
	if !ok {
		return LogMetricMetricDescriptorLabelsValueTypeEnumRef("")
	}

	return LogMetricMetricDescriptorLabelsValueTypeEnumRef(s)
}

// flattenLogMetricMetricDescriptorMetricKindEnumSlice flattens the contents of LogMetricMetricDescriptorMetricKindEnum from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetricKindEnumSlice(c *Client, i interface{}) []LogMetricMetricDescriptorMetricKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorMetricKindEnum{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorMetricKindEnum{}
	}

	items := make([]LogMetricMetricDescriptorMetricKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorMetricKindEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenLogMetricMetricDescriptorMetricKindEnum asserts that an interface is a string, and returns a
// pointer to a *LogMetricMetricDescriptorMetricKindEnum with the same value as that string.
func flattenLogMetricMetricDescriptorMetricKindEnum(i interface{}) *LogMetricMetricDescriptorMetricKindEnum {
	s, ok := i.(string)
	if !ok {
		return LogMetricMetricDescriptorMetricKindEnumRef("")
	}

	return LogMetricMetricDescriptorMetricKindEnumRef(s)
}

// flattenLogMetricMetricDescriptorValueTypeEnumSlice flattens the contents of LogMetricMetricDescriptorValueTypeEnum from a JSON
// response object.
func flattenLogMetricMetricDescriptorValueTypeEnumSlice(c *Client, i interface{}) []LogMetricMetricDescriptorValueTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorValueTypeEnum{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorValueTypeEnum{}
	}

	items := make([]LogMetricMetricDescriptorValueTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorValueTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenLogMetricMetricDescriptorValueTypeEnum asserts that an interface is a string, and returns a
// pointer to a *LogMetricMetricDescriptorValueTypeEnum with the same value as that string.
func flattenLogMetricMetricDescriptorValueTypeEnum(i interface{}) *LogMetricMetricDescriptorValueTypeEnum {
	s, ok := i.(string)
	if !ok {
		return LogMetricMetricDescriptorValueTypeEnumRef("")
	}

	return LogMetricMetricDescriptorValueTypeEnumRef(s)
}

// flattenLogMetricMetricDescriptorMetadataLaunchStageEnumSlice flattens the contents of LogMetricMetricDescriptorMetadataLaunchStageEnum from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataLaunchStageEnumSlice(c *Client, i interface{}) []LogMetricMetricDescriptorMetadataLaunchStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorMetadataLaunchStageEnum{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorMetadataLaunchStageEnum{}
	}

	items := make([]LogMetricMetricDescriptorMetadataLaunchStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorMetadataLaunchStageEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenLogMetricMetricDescriptorMetadataLaunchStageEnum asserts that an interface is a string, and returns a
// pointer to a *LogMetricMetricDescriptorMetadataLaunchStageEnum with the same value as that string.
func flattenLogMetricMetricDescriptorMetadataLaunchStageEnum(i interface{}) *LogMetricMetricDescriptorMetadataLaunchStageEnum {
	s, ok := i.(string)
	if !ok {
		return LogMetricMetricDescriptorMetadataLaunchStageEnumRef("")
	}

	return LogMetricMetricDescriptorMetadataLaunchStageEnumRef(s)
}

// flattenLogMetricMetricDescriptorLaunchStageEnumSlice flattens the contents of LogMetricMetricDescriptorLaunchStageEnum from a JSON
// response object.
func flattenLogMetricMetricDescriptorLaunchStageEnumSlice(c *Client, i interface{}) []LogMetricMetricDescriptorLaunchStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorLaunchStageEnum{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorLaunchStageEnum{}
	}

	items := make([]LogMetricMetricDescriptorLaunchStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorLaunchStageEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenLogMetricMetricDescriptorLaunchStageEnum asserts that an interface is a string, and returns a
// pointer to a *LogMetricMetricDescriptorLaunchStageEnum with the same value as that string.
func flattenLogMetricMetricDescriptorLaunchStageEnum(i interface{}) *LogMetricMetricDescriptorLaunchStageEnum {
	s, ok := i.(string)
	if !ok {
		return LogMetricMetricDescriptorLaunchStageEnumRef("")
	}

	return LogMetricMetricDescriptorLaunchStageEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *LogMetric) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalLogMetric(b, c)
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
