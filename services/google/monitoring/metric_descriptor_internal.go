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

func (r *MetricDescriptor) validate() error {

	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	if err := dcl.Required(r, "metricKind"); err != nil {
		return err
	}
	if err := dcl.Required(r, "valueType"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Metadata) {
		if err := r.Metadata.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MetricDescriptorDescriptorLabels) validate() error {
	return nil
}
func (r *MetricDescriptorMetadata) validate() error {
	return nil
}

func metricDescriptorGetURL(userBasePath string, r *MetricDescriptor) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"type":    dcl.ValueOrEmptyString(r.Type),
	}
	return dcl.URL("v3/projects/{{project}}/metricDescriptors/{{type}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

func metricDescriptorListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/metricDescriptors", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func metricDescriptorCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/metricDescriptors", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func metricDescriptorDeleteURL(userBasePath string, r *MetricDescriptor) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"type":    dcl.ValueOrEmptyString(r.Type),
	}
	return dcl.URL("v3/projects/{{project}}/metricDescriptors/{{type}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

// metricDescriptorApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type metricDescriptorApiOperation interface {
	do(context.Context, *MetricDescriptor, *Client) error
}

func (c *Client) listMetricDescriptorRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := metricDescriptorListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != MetricDescriptorMaxPage {
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

type listMetricDescriptorOperation struct {
	MetricDescriptors []map[string]interface{} `json:"metricDescriptors"`
	Token             string                   `json:"nextPageToken"`
}

func (c *Client) listMetricDescriptor(ctx context.Context, project, pageToken string, pageSize int32) ([]*MetricDescriptor, string, error) {
	b, err := c.listMetricDescriptorRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listMetricDescriptorOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*MetricDescriptor
	for _, v := range m.MetricDescriptors {
		res := flattenMetricDescriptor(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllMetricDescriptor(ctx context.Context, f func(*MetricDescriptor) bool, resources []*MetricDescriptor) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteMetricDescriptor(ctx, res)
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

type deleteMetricDescriptorOperation struct{}

func (op *deleteMetricDescriptorOperation) do(ctx context.Context, r *MetricDescriptor, c *Client) error {

	_, err := c.GetMetricDescriptor(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("MetricDescriptor not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetMetricDescriptor checking for existence. error: %v", err)
		return err
	}

	u, err := metricDescriptorDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete MetricDescriptor: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetMetricDescriptor(ctx, r.urlNormalized())
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
type createMetricDescriptorOperation struct {
	response map[string]interface{}
}

func (op *createMetricDescriptorOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createMetricDescriptorOperation) do(ctx context.Context, r *MetricDescriptor, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := metricDescriptorCreateURL(c.Config.BasePath, project)

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

	// Poll for the MetricDescriptor resource to be created. MetricDescriptor resources are eventually consistent but do not support operations
	// so we must repeatedly poll to check for their creation.
	err = dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		u, err := metricDescriptorGetURL(c.Config.BasePath, r.urlNormalized())
		if err != nil {
			return nil, err
		}
		getResp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, nil)
		if err != nil {
			// If the error is a transient server error (e.g., 500) or not found (i.e., the resource has not yet been created),
			// continue retrying until the transient error is resolved, the resource is created, or we time out.
			if dcl.IsRetryableRequestError(c.Config, err, true) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		getResp.Response.Body.Close()
		return getResp, nil
	}, c.Config.RetryProvider)

	if _, err := c.GetMetricDescriptor(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getMetricDescriptorRaw(ctx context.Context, r *MetricDescriptor) ([]byte, error) {

	u, err := metricDescriptorGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) metricDescriptorDiffsForRawDesired(ctx context.Context, rawDesired *MetricDescriptor, opts ...dcl.ApplyOption) (initial, desired *MetricDescriptor, diffs []metricDescriptorDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *MetricDescriptor
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*MetricDescriptor); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected MetricDescriptor, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetMetricDescriptor(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a MetricDescriptor resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve MetricDescriptor resource: %v", err)
		}
		c.Config.Logger.Info("Found that MetricDescriptor resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeMetricDescriptorDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for MetricDescriptor: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for MetricDescriptor: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeMetricDescriptorInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for MetricDescriptor: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeMetricDescriptorDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for MetricDescriptor: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffMetricDescriptor(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeMetricDescriptorInitialState(rawInitial, rawDesired *MetricDescriptor) (*MetricDescriptor, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeMetricDescriptorDesiredState(rawDesired, rawInitial *MetricDescriptor, opts ...dcl.ApplyOption) (*MetricDescriptor, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Metadata = canonicalizeMetricDescriptorMetadata(rawDesired.Metadata, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Type, rawInitial.Type) {
		rawDesired.Type = rawInitial.Type
	}
	if dcl.IsZeroValue(rawDesired.DescriptorLabels) {
		rawDesired.DescriptorLabels = rawInitial.DescriptorLabels
	}
	if dcl.IsZeroValue(rawDesired.MetricKind) {
		rawDesired.MetricKind = rawInitial.MetricKind
	}
	if canonicalizeMetricDescriptorValueType(rawDesired.ValueType, rawInitial.ValueType) {
		rawDesired.ValueType = rawInitial.ValueType
	}
	if dcl.StringCanonicalize(rawDesired.Unit, rawInitial.Unit) {
		rawDesired.Unit = rawInitial.Unit
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	rawDesired.Metadata = canonicalizeMetricDescriptorMetadata(rawDesired.Metadata, rawInitial.Metadata, opts...)
	if dcl.IsZeroValue(rawDesired.LaunchStage) {
		rawDesired.LaunchStage = rawInitial.LaunchStage
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeMetricDescriptorNewState(c *Client, rawNew, rawDesired *MetricDescriptor) (*MetricDescriptor, error) {

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
		if dcl.StringCanonicalize(rawDesired.Type, rawNew.Type) {
			rawNew.Type = rawDesired.Type
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DescriptorLabels) && dcl.IsEmptyValueIndirect(rawDesired.DescriptorLabels) {
		rawNew.DescriptorLabels = rawDesired.DescriptorLabels
	} else {
		rawNew.DescriptorLabels = canonicalizeNewMetricDescriptorDescriptorLabelsSet(c, rawDesired.DescriptorLabels, rawNew.DescriptorLabels)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MetricKind) && dcl.IsEmptyValueIndirect(rawDesired.MetricKind) {
		rawNew.MetricKind = rawDesired.MetricKind
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ValueType) && dcl.IsEmptyValueIndirect(rawDesired.ValueType) {
		rawNew.ValueType = rawDesired.ValueType
	} else {
		if canonicalizeMetricDescriptorValueType(rawDesired.ValueType, rawNew.ValueType) {
			rawNew.ValueType = rawDesired.ValueType
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Unit) && dcl.IsEmptyValueIndirect(rawDesired.Unit) {
		rawNew.Unit = rawDesired.Unit
	} else {
		if dcl.StringCanonicalize(rawDesired.Unit, rawNew.Unit) {
			rawNew.Unit = rawDesired.Unit
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	rawNew.Metadata = rawDesired.Metadata

	rawNew.LaunchStage = rawDesired.LaunchStage

	if dcl.IsEmptyValueIndirect(rawNew.MonitoredResourceTypes) && dcl.IsEmptyValueIndirect(rawDesired.MonitoredResourceTypes) {
		rawNew.MonitoredResourceTypes = rawDesired.MonitoredResourceTypes
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeMetricDescriptorDescriptorLabels(des, initial *MetricDescriptorDescriptorLabels, opts ...dcl.ApplyOption) *MetricDescriptorDescriptorLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if canonicalizeMetricDescriptorDescriptorLabelsValueType(des.ValueType, initial.ValueType) || dcl.IsZeroValue(des.ValueType) {
		des.ValueType = initial.ValueType
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}

	return des
}

func canonicalizeNewMetricDescriptorDescriptorLabels(c *Client, des, nw *MetricDescriptorDescriptorLabels) *MetricDescriptorDescriptorLabels {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if canonicalizeMetricDescriptorDescriptorLabelsValueType(des.ValueType, nw.ValueType) {
		nw.ValueType = des.ValueType
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewMetricDescriptorDescriptorLabelsSet(c *Client, des, nw []MetricDescriptorDescriptorLabels) []MetricDescriptorDescriptorLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []MetricDescriptorDescriptorLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMetricDescriptorDescriptorLabelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewMetricDescriptorDescriptorLabelsSlice(c *Client, des, nw []MetricDescriptorDescriptorLabels) []MetricDescriptorDescriptorLabels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MetricDescriptorDescriptorLabels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMetricDescriptorDescriptorLabels(c, &d, &n))
	}

	return items
}

func canonicalizeMetricDescriptorMetadata(des, initial *MetricDescriptorMetadata, opts ...dcl.ApplyOption) *MetricDescriptorMetadata {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.LaunchStage) {
		des.LaunchStage = initial.LaunchStage
	}
	if dcl.StringCanonicalize(des.SamplePeriod, initial.SamplePeriod) || dcl.IsZeroValue(des.SamplePeriod) {
		des.SamplePeriod = initial.SamplePeriod
	}
	if dcl.StringCanonicalize(des.IngestDelay, initial.IngestDelay) || dcl.IsZeroValue(des.IngestDelay) {
		des.IngestDelay = initial.IngestDelay
	}

	return des
}

func canonicalizeNewMetricDescriptorMetadata(c *Client, des, nw *MetricDescriptorMetadata) *MetricDescriptorMetadata {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.LaunchStage) {
		nw.LaunchStage = des.LaunchStage
	}
	if dcl.StringCanonicalize(des.SamplePeriod, nw.SamplePeriod) {
		nw.SamplePeriod = des.SamplePeriod
	}
	if dcl.StringCanonicalize(des.IngestDelay, nw.IngestDelay) {
		nw.IngestDelay = des.IngestDelay
	}

	return nw
}

func canonicalizeNewMetricDescriptorMetadataSet(c *Client, des, nw []MetricDescriptorMetadata) []MetricDescriptorMetadata {
	if des == nil {
		return nw
	}
	var reorderedNew []MetricDescriptorMetadata
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareMetricDescriptorMetadataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewMetricDescriptorMetadataSlice(c *Client, des, nw []MetricDescriptorMetadata) []MetricDescriptorMetadata {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MetricDescriptorMetadata
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMetricDescriptorMetadata(c, &d, &n))
	}

	return items
}

type metricDescriptorDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         metricDescriptorApiOperation
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
func diffMetricDescriptor(c *Client, desired, actual *MetricDescriptor, opts ...dcl.ApplyOption) ([]metricDescriptorDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []metricDescriptorDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.DescriptorLabels, actual.DescriptorLabels, dcl.Info{Type: "Set", ObjectFunction: compareMetricDescriptorDescriptorLabelsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DescriptorLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.MetricKind, actual.MetricKind, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetricKind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.ValueType, actual.ValueType, dcl.Info{Type: "EnumType", CustomDiff: canonicalizeMetricDescriptorValueType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ValueType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Unit, actual.Unit, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Unit")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{Ignore: true, ObjectFunction: compareMetricDescriptorMetadataNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.LaunchStage, actual.LaunchStage, dcl.Info{Ignore: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LaunchStage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.MonitoredResourceTypes, actual.MonitoredResourceTypes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MonitoredResourceTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToMetricDescriptorDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
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
	var deduped []metricDescriptorDiff
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
func compareMetricDescriptorDescriptorLabelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MetricDescriptorDescriptorLabels)
	if !ok {
		desiredNotPointer, ok := d.(MetricDescriptorDescriptorLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MetricDescriptorDescriptorLabels or *MetricDescriptorDescriptorLabels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MetricDescriptorDescriptorLabels)
	if !ok {
		actualNotPointer, ok := a.(MetricDescriptorDescriptorLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MetricDescriptorDescriptorLabels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueType, actual.ValueType, dcl.Info{Type: "EnumType", CustomDiff: canonicalizeMetricDescriptorDescriptorLabelsValueType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ValueType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMetricDescriptorMetadataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MetricDescriptorMetadata)
	if !ok {
		desiredNotPointer, ok := d.(MetricDescriptorMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MetricDescriptorMetadata or *MetricDescriptorMetadata", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MetricDescriptorMetadata)
	if !ok {
		actualNotPointer, ok := a.(MetricDescriptorMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MetricDescriptorMetadata", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LaunchStage, actual.LaunchStage, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LaunchStage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SamplePeriod, actual.SamplePeriod, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SamplePeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IngestDelay, actual.IngestDelay, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IngestDelay")); len(ds) != 0 || err != nil {
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
func (r *MetricDescriptor) urlNormalized() *MetricDescriptor {
	normalized := dcl.Copy(*r).(MetricDescriptor)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Type = r.Type
	normalized.Unit = dcl.SelfLinkToName(r.Unit)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *MetricDescriptor) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Type)
}

func (r *MetricDescriptor) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *MetricDescriptor) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Type)
}

func (r *MetricDescriptor) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the MetricDescriptor resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *MetricDescriptor) marshal(c *Client) ([]byte, error) {
	m, err := expandMetricDescriptor(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling MetricDescriptor: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalMetricDescriptor decodes JSON responses into the MetricDescriptor resource schema.
func unmarshalMetricDescriptor(b []byte, c *Client) (*MetricDescriptor, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapMetricDescriptor(m, c)
}

func unmarshalMapMetricDescriptor(m map[string]interface{}, c *Client) (*MetricDescriptor, error) {

	return flattenMetricDescriptor(c, m), nil
}

// expandMetricDescriptor expands MetricDescriptor into a JSON request object.
func expandMetricDescriptor(c *Client, f *MetricDescriptor) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v, err := expandMetricDescriptorDescriptorLabelsSlice(c, f.DescriptorLabels); err != nil {
		return nil, fmt.Errorf("error expanding DescriptorLabels into labels: %w", err)
	} else if v != nil {
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
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v, err := expandMetricDescriptorMetadata(c, f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if v != nil {
		m["metadata"] = v
	}
	if v := f.LaunchStage; !dcl.IsEmptyValueIndirect(v) {
		m["launchStage"] = v
	}
	if v := f.MonitoredResourceTypes; !dcl.IsEmptyValueIndirect(v) {
		m["monitoredResourceTypes"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenMetricDescriptor flattens MetricDescriptor from a JSON request object into the
// MetricDescriptor type.
func flattenMetricDescriptor(c *Client, i interface{}) *MetricDescriptor {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &MetricDescriptor{}
	res.SelfLink = dcl.FlattenString(m["name"])
	res.Type = dcl.FlattenString(m["type"])
	res.DescriptorLabels = flattenMetricDescriptorDescriptorLabelsSlice(c, m["labels"])
	res.MetricKind = flattenMetricDescriptorMetricKindEnum(m["metricKind"])
	res.ValueType = flattenMetricDescriptorValueTypeEnum(m["valueType"])
	res.Unit = dcl.FlattenString(m["unit"])
	res.Description = dcl.FlattenString(m["description"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.Metadata = flattenMetricDescriptorMetadata(c, m["metadata"])
	res.LaunchStage = flattenMetricDescriptorLaunchStageEnum(m["launchStage"])
	res.MonitoredResourceTypes = dcl.FlattenStringSlice(m["monitoredResourceTypes"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandMetricDescriptorDescriptorLabelsMap expands the contents of MetricDescriptorDescriptorLabels into a JSON
// request object.
func expandMetricDescriptorDescriptorLabelsMap(c *Client, f map[string]MetricDescriptorDescriptorLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMetricDescriptorDescriptorLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMetricDescriptorDescriptorLabelsSlice expands the contents of MetricDescriptorDescriptorLabels into a JSON
// request object.
func expandMetricDescriptorDescriptorLabelsSlice(c *Client, f []MetricDescriptorDescriptorLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMetricDescriptorDescriptorLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMetricDescriptorDescriptorLabelsMap flattens the contents of MetricDescriptorDescriptorLabels from a JSON
// response object.
func flattenMetricDescriptorDescriptorLabelsMap(c *Client, i interface{}) map[string]MetricDescriptorDescriptorLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MetricDescriptorDescriptorLabels{}
	}

	if len(a) == 0 {
		return map[string]MetricDescriptorDescriptorLabels{}
	}

	items := make(map[string]MetricDescriptorDescriptorLabels)
	for k, item := range a {
		items[k] = *flattenMetricDescriptorDescriptorLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMetricDescriptorDescriptorLabelsSlice flattens the contents of MetricDescriptorDescriptorLabels from a JSON
// response object.
func flattenMetricDescriptorDescriptorLabelsSlice(c *Client, i interface{}) []MetricDescriptorDescriptorLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []MetricDescriptorDescriptorLabels{}
	}

	if len(a) == 0 {
		return []MetricDescriptorDescriptorLabels{}
	}

	items := make([]MetricDescriptorDescriptorLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetricDescriptorDescriptorLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandMetricDescriptorDescriptorLabels expands an instance of MetricDescriptorDescriptorLabels into a JSON
// request object.
func expandMetricDescriptorDescriptorLabels(c *Client, f *MetricDescriptorDescriptorLabels) (map[string]interface{}, error) {
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

// flattenMetricDescriptorDescriptorLabels flattens an instance of MetricDescriptorDescriptorLabels from a JSON
// response object.
func flattenMetricDescriptorDescriptorLabels(c *Client, i interface{}) *MetricDescriptorDescriptorLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MetricDescriptorDescriptorLabels{}
	r.Key = dcl.FlattenString(m["key"])
	r.ValueType = flattenMetricDescriptorDescriptorLabelsValueTypeEnum(m["valueType"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandMetricDescriptorMetadataMap expands the contents of MetricDescriptorMetadata into a JSON
// request object.
func expandMetricDescriptorMetadataMap(c *Client, f map[string]MetricDescriptorMetadata) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMetricDescriptorMetadata(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMetricDescriptorMetadataSlice expands the contents of MetricDescriptorMetadata into a JSON
// request object.
func expandMetricDescriptorMetadataSlice(c *Client, f []MetricDescriptorMetadata) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMetricDescriptorMetadata(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMetricDescriptorMetadataMap flattens the contents of MetricDescriptorMetadata from a JSON
// response object.
func flattenMetricDescriptorMetadataMap(c *Client, i interface{}) map[string]MetricDescriptorMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MetricDescriptorMetadata{}
	}

	if len(a) == 0 {
		return map[string]MetricDescriptorMetadata{}
	}

	items := make(map[string]MetricDescriptorMetadata)
	for k, item := range a {
		items[k] = *flattenMetricDescriptorMetadata(c, item.(map[string]interface{}))
	}

	return items
}

// flattenMetricDescriptorMetadataSlice flattens the contents of MetricDescriptorMetadata from a JSON
// response object.
func flattenMetricDescriptorMetadataSlice(c *Client, i interface{}) []MetricDescriptorMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []MetricDescriptorMetadata{}
	}

	if len(a) == 0 {
		return []MetricDescriptorMetadata{}
	}

	items := make([]MetricDescriptorMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetricDescriptorMetadata(c, item.(map[string]interface{})))
	}

	return items
}

// expandMetricDescriptorMetadata expands an instance of MetricDescriptorMetadata into a JSON
// request object.
func expandMetricDescriptorMetadata(c *Client, f *MetricDescriptorMetadata) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.LaunchStage; !dcl.IsEmptyValueIndirect(v) {
		m["launchStage"] = v
	}
	if v := f.SamplePeriod; !dcl.IsEmptyValueIndirect(v) {
		m["samplePeriod"] = v
	}
	if v := f.IngestDelay; !dcl.IsEmptyValueIndirect(v) {
		m["ingestDelay"] = v
	}

	return m, nil
}

// flattenMetricDescriptorMetadata flattens an instance of MetricDescriptorMetadata from a JSON
// response object.
func flattenMetricDescriptorMetadata(c *Client, i interface{}) *MetricDescriptorMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MetricDescriptorMetadata{}
	r.LaunchStage = flattenMetricDescriptorMetadataLaunchStageEnum(m["launchStage"])
	r.SamplePeriod = dcl.FlattenString(m["samplePeriod"])
	r.IngestDelay = dcl.FlattenString(m["ingestDelay"])

	return r
}

// flattenMetricDescriptorDescriptorLabelsValueTypeEnumSlice flattens the contents of MetricDescriptorDescriptorLabelsValueTypeEnum from a JSON
// response object.
func flattenMetricDescriptorDescriptorLabelsValueTypeEnumSlice(c *Client, i interface{}) []MetricDescriptorDescriptorLabelsValueTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MetricDescriptorDescriptorLabelsValueTypeEnum{}
	}

	if len(a) == 0 {
		return []MetricDescriptorDescriptorLabelsValueTypeEnum{}
	}

	items := make([]MetricDescriptorDescriptorLabelsValueTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetricDescriptorDescriptorLabelsValueTypeEnum(item.(interface{})))
	}

	return items
}

// flattenMetricDescriptorDescriptorLabelsValueTypeEnum asserts that an interface is a string, and returns a
// pointer to a *MetricDescriptorDescriptorLabelsValueTypeEnum with the same value as that string.
func flattenMetricDescriptorDescriptorLabelsValueTypeEnum(i interface{}) *MetricDescriptorDescriptorLabelsValueTypeEnum {
	s, ok := i.(string)
	if !ok {
		return MetricDescriptorDescriptorLabelsValueTypeEnumRef("")
	}

	return MetricDescriptorDescriptorLabelsValueTypeEnumRef(s)
}

// flattenMetricDescriptorMetricKindEnumSlice flattens the contents of MetricDescriptorMetricKindEnum from a JSON
// response object.
func flattenMetricDescriptorMetricKindEnumSlice(c *Client, i interface{}) []MetricDescriptorMetricKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MetricDescriptorMetricKindEnum{}
	}

	if len(a) == 0 {
		return []MetricDescriptorMetricKindEnum{}
	}

	items := make([]MetricDescriptorMetricKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetricDescriptorMetricKindEnum(item.(interface{})))
	}

	return items
}

// flattenMetricDescriptorMetricKindEnum asserts that an interface is a string, and returns a
// pointer to a *MetricDescriptorMetricKindEnum with the same value as that string.
func flattenMetricDescriptorMetricKindEnum(i interface{}) *MetricDescriptorMetricKindEnum {
	s, ok := i.(string)
	if !ok {
		return MetricDescriptorMetricKindEnumRef("")
	}

	return MetricDescriptorMetricKindEnumRef(s)
}

// flattenMetricDescriptorValueTypeEnumSlice flattens the contents of MetricDescriptorValueTypeEnum from a JSON
// response object.
func flattenMetricDescriptorValueTypeEnumSlice(c *Client, i interface{}) []MetricDescriptorValueTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MetricDescriptorValueTypeEnum{}
	}

	if len(a) == 0 {
		return []MetricDescriptorValueTypeEnum{}
	}

	items := make([]MetricDescriptorValueTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetricDescriptorValueTypeEnum(item.(interface{})))
	}

	return items
}

// flattenMetricDescriptorValueTypeEnum asserts that an interface is a string, and returns a
// pointer to a *MetricDescriptorValueTypeEnum with the same value as that string.
func flattenMetricDescriptorValueTypeEnum(i interface{}) *MetricDescriptorValueTypeEnum {
	s, ok := i.(string)
	if !ok {
		return MetricDescriptorValueTypeEnumRef("")
	}

	return MetricDescriptorValueTypeEnumRef(s)
}

// flattenMetricDescriptorMetadataLaunchStageEnumSlice flattens the contents of MetricDescriptorMetadataLaunchStageEnum from a JSON
// response object.
func flattenMetricDescriptorMetadataLaunchStageEnumSlice(c *Client, i interface{}) []MetricDescriptorMetadataLaunchStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MetricDescriptorMetadataLaunchStageEnum{}
	}

	if len(a) == 0 {
		return []MetricDescriptorMetadataLaunchStageEnum{}
	}

	items := make([]MetricDescriptorMetadataLaunchStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetricDescriptorMetadataLaunchStageEnum(item.(interface{})))
	}

	return items
}

// flattenMetricDescriptorMetadataLaunchStageEnum asserts that an interface is a string, and returns a
// pointer to a *MetricDescriptorMetadataLaunchStageEnum with the same value as that string.
func flattenMetricDescriptorMetadataLaunchStageEnum(i interface{}) *MetricDescriptorMetadataLaunchStageEnum {
	s, ok := i.(string)
	if !ok {
		return MetricDescriptorMetadataLaunchStageEnumRef("")
	}

	return MetricDescriptorMetadataLaunchStageEnumRef(s)
}

// flattenMetricDescriptorLaunchStageEnumSlice flattens the contents of MetricDescriptorLaunchStageEnum from a JSON
// response object.
func flattenMetricDescriptorLaunchStageEnumSlice(c *Client, i interface{}) []MetricDescriptorLaunchStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MetricDescriptorLaunchStageEnum{}
	}

	if len(a) == 0 {
		return []MetricDescriptorLaunchStageEnum{}
	}

	items := make([]MetricDescriptorLaunchStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetricDescriptorLaunchStageEnum(item.(interface{})))
	}

	return items
}

// flattenMetricDescriptorLaunchStageEnum asserts that an interface is a string, and returns a
// pointer to a *MetricDescriptorLaunchStageEnum with the same value as that string.
func flattenMetricDescriptorLaunchStageEnum(i interface{}) *MetricDescriptorLaunchStageEnum {
	s, ok := i.(string)
	if !ok {
		return MetricDescriptorLaunchStageEnumRef("")
	}

	return MetricDescriptorLaunchStageEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *MetricDescriptor) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalMetricDescriptor(b, c)
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
		if nr.Type == nil && ncr.Type == nil {
			c.Config.Logger.Info("Both Type fields null - considering equal.")
		} else if nr.Type == nil || ncr.Type == nil {
			c.Config.Logger.Info("Only one Type field is null - considering unequal.")
			return false
		} else if *nr.Type != *ncr.Type {
			return false
		}
		return true
	}
}

func convertFieldDiffToMetricDescriptorDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]metricDescriptorDiff, error) {
	var diffs []metricDescriptorDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := metricDescriptorDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameTometricDescriptorApiOperation(op, opts...)
				if err != nil {
					return nil, err
				}
				diff.UpdateOp = op
			}
			diffs = append(diffs, diff)
		}
	}
	return diffs, nil
}

func convertOpNameTometricDescriptorApiOperation(op string, opts ...dcl.ApplyOption) (metricDescriptorApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
