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
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
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
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://monitoring.googleapis.com/", "GET"); err != nil {
		return err
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
		u, err := metricDescriptorGetURL(c.Config.BasePath, r)
		if err != nil {
			return nil, err
		}
		getResp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, nil)
		if err != nil {
			// If the error is a transient server error (e.g., 500) or not found (i.e., the resource has not yet been created),
			// continue retrying until the transient error is resolved, the resource is created, or we time out.
			if dcl.IsRetryableRequestError(c.Config, err, true) {
				return nil, dcl.OperationNotDone{Err: err}
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
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
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
	if dcl.IsZeroValue(rawDesired.ValueType) {
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
	if dcl.IsZeroValue(rawDesired.MonitoredResourceTypes) {
		rawDesired.MonitoredResourceTypes = rawInitial.MonitoredResourceTypes
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

	if dcl.IsEmptyValueIndirect(rawNew.Metadata) && dcl.IsEmptyValueIndirect(rawDesired.Metadata) {
		rawNew.Metadata = rawDesired.Metadata
	} else {
		rawNew.Metadata = canonicalizeNewMetricDescriptorMetadata(c, rawDesired.Metadata, rawNew.Metadata)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LaunchStage) && dcl.IsEmptyValueIndirect(rawDesired.LaunchStage) {
		rawNew.LaunchStage = rawDesired.LaunchStage
	} else {
	}

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
	if dcl.IsZeroValue(des.ValueType) {
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

	if dcl.StringCanonicalize(des.Key, nw.Key) || dcl.IsZeroValue(des.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) || dcl.IsZeroValue(des.Description) {
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
			if !compareMetricDescriptorDescriptorLabels(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
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

	if dcl.StringCanonicalize(des.SamplePeriod, nw.SamplePeriod) || dcl.IsZeroValue(des.SamplePeriod) {
		nw.SamplePeriod = des.SamplePeriod
	}
	if dcl.StringCanonicalize(des.IngestDelay, nw.IngestDelay) || dcl.IsZeroValue(des.IngestDelay) {
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
			if !compareMetricDescriptorMetadata(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

type metricDescriptorDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         metricDescriptorApiOperation
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
	if !dcl.IsZeroValue(desired.Type) && !dcl.StringCanonicalize(desired.Type, actual.Type) {
		c.Config.Logger.Infof("Detected diff in Type.\nDESIRED: %v\nACTUAL: %v", desired.Type, actual.Type)
		diffs = append(diffs, metricDescriptorDiff{
			RequiresRecreate: true,
			FieldName:        "Type",
		})
	}
	if compareMetricDescriptorDescriptorLabelsSlice(c, desired.DescriptorLabels, actual.DescriptorLabels) {
		c.Config.Logger.Infof("Detected diff in DescriptorLabels.\nDESIRED: %v\nACTUAL: %v", desired.DescriptorLabels, actual.DescriptorLabels)
		toAdd, toRemove := compareMetricDescriptorDescriptorLabelsSets(c, desired.DescriptorLabels, actual.DescriptorLabels)
		if len(toAdd) > 0 {
			diffs = append(diffs, metricDescriptorDiff{
				RequiresRecreate: true,
				FieldName:        "DescriptorLabels",
			})
		}
		if len(toRemove) > 0 {
			diffs = append(diffs, metricDescriptorDiff{
				RequiresRecreate: true,
				FieldName:        "DescriptorLabels",
			})
		}
	}
	if !reflect.DeepEqual(desired.MetricKind, actual.MetricKind) {
		c.Config.Logger.Infof("Detected diff in MetricKind.\nDESIRED: %v\nACTUAL: %v", desired.MetricKind, actual.MetricKind)
		diffs = append(diffs, metricDescriptorDiff{
			RequiresRecreate: true,
			FieldName:        "MetricKind",
		})
	}
	if !reflect.DeepEqual(desired.ValueType, actual.ValueType) {
		c.Config.Logger.Infof("Detected diff in ValueType.\nDESIRED: %v\nACTUAL: %v", desired.ValueType, actual.ValueType)
		diffs = append(diffs, metricDescriptorDiff{
			RequiresRecreate: true,
			FieldName:        "ValueType",
		})
	}
	if !dcl.IsZeroValue(desired.Unit) && !dcl.StringCanonicalize(desired.Unit, actual.Unit) {
		c.Config.Logger.Infof("Detected diff in Unit.\nDESIRED: %v\nACTUAL: %v", desired.Unit, actual.Unit)
		diffs = append(diffs, metricDescriptorDiff{
			RequiresRecreate: true,
			FieldName:        "Unit",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, metricDescriptorDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.DisplayName) && !dcl.StringCanonicalize(desired.DisplayName, actual.DisplayName) {
		c.Config.Logger.Infof("Detected diff in DisplayName.\nDESIRED: %v\nACTUAL: %v", desired.DisplayName, actual.DisplayName)
		diffs = append(diffs, metricDescriptorDiff{
			RequiresRecreate: true,
			FieldName:        "DisplayName",
		})
	}
	if compareMetricDescriptorMetadata(c, desired.Metadata, actual.Metadata) {
		c.Config.Logger.Infof("Detected diff in Metadata.\nDESIRED: %v\nACTUAL: %v", desired.Metadata, actual.Metadata)
		diffs = append(diffs, metricDescriptorDiff{
			RequiresRecreate: true,
			FieldName:        "Metadata",
		})
	}
	if !reflect.DeepEqual(desired.LaunchStage, actual.LaunchStage) {
		c.Config.Logger.Infof("Detected diff in LaunchStage.\nDESIRED: %v\nACTUAL: %v", desired.LaunchStage, actual.LaunchStage)
		diffs = append(diffs, metricDescriptorDiff{
			RequiresRecreate: true,
			FieldName:        "LaunchStage",
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
func compareMetricDescriptorDescriptorLabels(c *Client, desired, actual *MetricDescriptorDescriptorLabels) bool {
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
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.ValueType == nil && desired.ValueType != nil && !dcl.IsEmptyValueIndirect(desired.ValueType) {
		c.Config.Logger.Infof("desired ValueType %s - but actually nil", dcl.SprintResource(desired.ValueType))
		return true
	}
	if !reflect.DeepEqual(desired.ValueType, actual.ValueType) && !dcl.IsZeroValue(desired.ValueType) {
		c.Config.Logger.Infof("Diff in ValueType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ValueType), dcl.SprintResource(actual.ValueType))
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
	return false
}

func compareMetricDescriptorDescriptorLabelsSlice(c *Client, desired, actual []MetricDescriptorDescriptorLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in MetricDescriptorDescriptorLabels, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareMetricDescriptorDescriptorLabels(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in MetricDescriptorDescriptorLabels, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareMetricDescriptorDescriptorLabelsMap(c *Client, desired, actual map[string]MetricDescriptorDescriptorLabels) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in MetricDescriptorDescriptorLabels, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in MetricDescriptorDescriptorLabels, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareMetricDescriptorDescriptorLabels(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in MetricDescriptorDescriptorLabels, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareMetricDescriptorDescriptorLabelsSets(c *Client, desired, actual []MetricDescriptorDescriptorLabels) (toAdd, toRemove []MetricDescriptorDescriptorLabels) {
	if actual == nil {
		return desired, nil
	}
	desiredHashes := make(map[string]MetricDescriptorDescriptorLabels)
	for _, d := range desired {
		desiredHashes[d.HashCode()] = d
	}
	actualHashes := make(map[string]MetricDescriptorDescriptorLabels)
	for _, a := range actual {
		actualHashes[a.HashCode()] = a
	}
	toAdd = make([]MetricDescriptorDescriptorLabels, 0)
	toRemove = make([]MetricDescriptorDescriptorLabels, 0)

	for k, v := range actualHashes {
		_, found := desiredHashes[k]
		if !found {
			// backup - search linearly for equivalent value.
			for _, des := range desiredHashes {
				if !compareMetricDescriptorDescriptorLabels(c, &des, &v) {
					found = true
					break
				}
			}
		}
		if !found {
			toRemove = append(toRemove, v)
		}
	}

	for k, v := range desiredHashes {
		_, found := actualHashes[k]
		if !found {
			for _, act := range actualHashes {
				if !compareMetricDescriptorDescriptorLabels(c, &v, &act) {
					found = true
					break
				}
			}
		}
		if !found {
			toAdd = append(toAdd, v)
		}
	}

	return toAdd, toRemove
}

func compareMetricDescriptorMetadata(c *Client, desired, actual *MetricDescriptorMetadata) bool {
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
	if !reflect.DeepEqual(desired.LaunchStage, actual.LaunchStage) && !dcl.IsZeroValue(desired.LaunchStage) {
		c.Config.Logger.Infof("Diff in LaunchStage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LaunchStage), dcl.SprintResource(actual.LaunchStage))
		return true
	}
	if actual.SamplePeriod == nil && desired.SamplePeriod != nil && !dcl.IsEmptyValueIndirect(desired.SamplePeriod) {
		c.Config.Logger.Infof("desired SamplePeriod %s - but actually nil", dcl.SprintResource(desired.SamplePeriod))
		return true
	}
	if !dcl.StringCanonicalize(desired.SamplePeriod, actual.SamplePeriod) && !dcl.IsZeroValue(desired.SamplePeriod) {
		c.Config.Logger.Infof("Diff in SamplePeriod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SamplePeriod), dcl.SprintResource(actual.SamplePeriod))
		return true
	}
	if actual.IngestDelay == nil && desired.IngestDelay != nil && !dcl.IsEmptyValueIndirect(desired.IngestDelay) {
		c.Config.Logger.Infof("desired IngestDelay %s - but actually nil", dcl.SprintResource(desired.IngestDelay))
		return true
	}
	if !dcl.StringCanonicalize(desired.IngestDelay, actual.IngestDelay) && !dcl.IsZeroValue(desired.IngestDelay) {
		c.Config.Logger.Infof("Diff in IngestDelay. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IngestDelay), dcl.SprintResource(actual.IngestDelay))
		return true
	}
	return false
}

func compareMetricDescriptorMetadataSlice(c *Client, desired, actual []MetricDescriptorMetadata) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in MetricDescriptorMetadata, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareMetricDescriptorMetadata(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in MetricDescriptorMetadata, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareMetricDescriptorMetadataMap(c *Client, desired, actual map[string]MetricDescriptorMetadata) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in MetricDescriptorMetadata, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in MetricDescriptorMetadata, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareMetricDescriptorMetadata(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in MetricDescriptorMetadata, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareMetricDescriptorDescriptorLabelsValueTypeEnumSlice(c *Client, desired, actual []MetricDescriptorDescriptorLabelsValueTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in MetricDescriptorDescriptorLabelsValueTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareMetricDescriptorDescriptorLabelsValueTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in MetricDescriptorDescriptorLabelsValueTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareMetricDescriptorDescriptorLabelsValueTypeEnum(c *Client, desired, actual *MetricDescriptorDescriptorLabelsValueTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareMetricDescriptorMetricKindEnumSlice(c *Client, desired, actual []MetricDescriptorMetricKindEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in MetricDescriptorMetricKindEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareMetricDescriptorMetricKindEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in MetricDescriptorMetricKindEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareMetricDescriptorMetricKindEnum(c *Client, desired, actual *MetricDescriptorMetricKindEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareMetricDescriptorValueTypeEnumSlice(c *Client, desired, actual []MetricDescriptorValueTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in MetricDescriptorValueTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareMetricDescriptorValueTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in MetricDescriptorValueTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareMetricDescriptorValueTypeEnum(c *Client, desired, actual *MetricDescriptorValueTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareMetricDescriptorMetadataLaunchStageEnumSlice(c *Client, desired, actual []MetricDescriptorMetadataLaunchStageEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in MetricDescriptorMetadataLaunchStageEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareMetricDescriptorMetadataLaunchStageEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in MetricDescriptorMetadataLaunchStageEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareMetricDescriptorMetadataLaunchStageEnum(c *Client, desired, actual *MetricDescriptorMetadataLaunchStageEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareMetricDescriptorLaunchStageEnumSlice(c *Client, desired, actual []MetricDescriptorLaunchStageEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in MetricDescriptorLaunchStageEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareMetricDescriptorLaunchStageEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in MetricDescriptorLaunchStageEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareMetricDescriptorLaunchStageEnum(c *Client, desired, actual *MetricDescriptorLaunchStageEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *MetricDescriptor) urlNormalized() *MetricDescriptor {
	normalized := deepcopy.Copy(*r).(MetricDescriptor)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Type = dcl.SelfLinkToName(r.Type)
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
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v, err := expandMetricDescriptorMetadata(c, f.Metadata); err != nil {
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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

	r := &MetricDescriptor{}
	r.SelfLink = dcl.FlattenString(m["name"])
	r.Type = dcl.FlattenString(m["type"])
	r.DescriptorLabels = flattenMetricDescriptorDescriptorLabelsSlice(c, m["labels"])
	r.MetricKind = flattenMetricDescriptorMetricKindEnum(m["metricKind"])
	r.ValueType = flattenMetricDescriptorValueTypeEnum(m["valueType"])
	r.Unit = dcl.FlattenString(m["unit"])
	r.Description = dcl.FlattenString(m["description"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.Metadata = flattenMetricDescriptorMetadata(c, m["metadata"])
	r.LaunchStage = flattenMetricDescriptorLaunchStageEnum(m["launchStage"])
	r.MonitoredResourceTypes = dcl.FlattenStringSlice(m["monitoredResourceTypes"])
	r.Project = dcl.FlattenString(m["project"])

	return r
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
