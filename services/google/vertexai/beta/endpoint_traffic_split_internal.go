// Copyright 2022 Google LLC. All Rights Reserved.
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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *EndpointTrafficSplit) validate() error {

	if err := dcl.RequiredParameter(r.Endpoint, "Endpoint"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.Required(r, "trafficSplit"); err != nil {
		return err
	}
	return nil
}
func (r *EndpointTrafficSplitTrafficSplit) validate() error {
	if err := dcl.Required(r, "deployedModelId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "trafficPercentage"); err != nil {
		return err
	}
	return nil
}
func (r *EndpointTrafficSplit) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-aiplatform.googleapis.com/v1beta1/", params)
}

func (r *EndpointTrafficSplit) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"endpoint": dcl.ValueOrEmptyString(nr.Endpoint),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}", nr.basePath(), userBasePath, params), nil
}

// endpointTrafficSplitApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type endpointTrafficSplitApiOperation interface {
	do(context.Context, *EndpointTrafficSplit, *Client) error
}

// newUpdateEndpointTrafficSplitUpdateEndpointTrafficSplitRequest creates a request for an
// EndpointTrafficSplit resource's UpdateEndpointTrafficSplit update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateEndpointTrafficSplitUpdateEndpointTrafficSplitRequest(ctx context.Context, f *EndpointTrafficSplit, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := endpointTrafficSplitSliceToMap(c, f.TrafficSplit, res); err != nil {
		return nil, fmt.Errorf("error expanding TrafficSplit into trafficSplit: %w", err)
	} else if v != nil {
		req["trafficSplit"] = v
	}
	b, err := c.getEndpointTrafficSplitRaw(ctx, f)
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
	return req, nil
}

// marshalUpdateEndpointTrafficSplitUpdateEndpointTrafficSplitRequest converts the update into
// the final JSON request body.
func marshalUpdateEndpointTrafficSplitUpdateEndpointTrafficSplitRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateEndpointTrafficSplitUpdateEndpointTrafficSplitOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateEndpointTrafficSplitUpdateEndpointTrafficSplitOperation) do(ctx context.Context, r *EndpointTrafficSplit, c *Client) error {
	_, err := c.GetEndpointTrafficSplit(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateEndpointTrafficSplit")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateEndpointTrafficSplitUpdateEndpointTrafficSplitRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateEndpointTrafficSplitUpdateEndpointTrafficSplitRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createEndpointTrafficSplitOperation struct {
	response map[string]interface{}
}

func (op *createEndpointTrafficSplitOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getEndpointTrafficSplitRaw(ctx context.Context, r *EndpointTrafficSplit) ([]byte, error) {

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

func (c *Client) endpointTrafficSplitDiffsForRawDesired(ctx context.Context, rawDesired *EndpointTrafficSplit, opts ...dcl.ApplyOption) (initial, desired *EndpointTrafficSplit, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *EndpointTrafficSplit
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*EndpointTrafficSplit); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected EndpointTrafficSplit, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEndpointTrafficSplit(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a EndpointTrafficSplit resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve EndpointTrafficSplit resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that EndpointTrafficSplit resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEndpointTrafficSplitDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for EndpointTrafficSplit: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for EndpointTrafficSplit: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractEndpointTrafficSplitFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEndpointTrafficSplitInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for EndpointTrafficSplit: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEndpointTrafficSplitDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for EndpointTrafficSplit: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEndpointTrafficSplit(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEndpointTrafficSplitInitialState(rawInitial, rawDesired *EndpointTrafficSplit) (*EndpointTrafficSplit, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEndpointTrafficSplitDesiredState(rawDesired, rawInitial *EndpointTrafficSplit, opts ...dcl.ApplyOption) (*EndpointTrafficSplit, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &EndpointTrafficSplit{}
	if dcl.NameToSelfLink(rawDesired.Endpoint, rawInitial.Endpoint) {
		canonicalDesired.Endpoint = rawInitial.Endpoint
	} else {
		canonicalDesired.Endpoint = rawDesired.Endpoint
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	canonicalDesired.TrafficSplit = canonicalizeEndpointTrafficSplitTrafficSplitSlice(rawDesired.TrafficSplit, rawInitial.TrafficSplit, opts...)

	return canonicalDesired, nil
}

func canonicalizeEndpointTrafficSplitNewState(c *Client, rawNew, rawDesired *EndpointTrafficSplit) (*EndpointTrafficSplit, error) {

	rawNew.Endpoint = rawDesired.Endpoint

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TrafficSplit) && dcl.IsEmptyValueIndirect(rawDesired.TrafficSplit) {
		rawNew.TrafficSplit = rawDesired.TrafficSplit
	} else {
		rawNew.TrafficSplit = canonicalizeNewEndpointTrafficSplitTrafficSplitSlice(c, rawDesired.TrafficSplit, rawNew.TrafficSplit)
	}

	return rawNew, nil
}

func canonicalizeEndpointTrafficSplitTrafficSplit(des, initial *EndpointTrafficSplitTrafficSplit, opts ...dcl.ApplyOption) *EndpointTrafficSplitTrafficSplit {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointTrafficSplitTrafficSplit{}

	if dcl.IsZeroValue(des.DeployedModelId) || (dcl.IsEmptyValueIndirect(des.DeployedModelId) && dcl.IsEmptyValueIndirect(initial.DeployedModelId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DeployedModelId = initial.DeployedModelId
	} else {
		cDes.DeployedModelId = des.DeployedModelId
	}
	if dcl.IsZeroValue(des.TrafficPercentage) || (dcl.IsEmptyValueIndirect(des.TrafficPercentage) && dcl.IsEmptyValueIndirect(initial.TrafficPercentage)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.TrafficPercentage = initial.TrafficPercentage
	} else {
		cDes.TrafficPercentage = des.TrafficPercentage
	}

	return cDes
}

func canonicalizeEndpointTrafficSplitTrafficSplitSlice(des, initial []EndpointTrafficSplitTrafficSplit, opts ...dcl.ApplyOption) []EndpointTrafficSplitTrafficSplit {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointTrafficSplitTrafficSplit, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointTrafficSplitTrafficSplit(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointTrafficSplitTrafficSplit, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointTrafficSplitTrafficSplit(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointTrafficSplitTrafficSplit(c *Client, des, nw *EndpointTrafficSplitTrafficSplit) *EndpointTrafficSplitTrafficSplit {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointTrafficSplitTrafficSplit while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewEndpointTrafficSplitTrafficSplitSet(c *Client, des, nw []EndpointTrafficSplitTrafficSplit) []EndpointTrafficSplitTrafficSplit {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointTrafficSplitTrafficSplit
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointTrafficSplitTrafficSplitNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewEndpointTrafficSplitTrafficSplitSlice(c *Client, des, nw []EndpointTrafficSplitTrafficSplit) []EndpointTrafficSplitTrafficSplit {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointTrafficSplitTrafficSplit
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointTrafficSplitTrafficSplit(c, &d, &n))
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
func diffEndpointTrafficSplit(c *Client, desired, actual *EndpointTrafficSplit, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Endpoint, actual.Endpoint, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Endpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrafficSplit, actual.TrafficSplit, dcl.DiffInfo{ObjectFunction: compareEndpointTrafficSplitTrafficSplitNewStyle, EmptyObject: EmptyEndpointTrafficSplitTrafficSplit, OperationSelector: dcl.TriggersOperation("updateEndpointTrafficSplitUpdateEndpointTrafficSplitOperation")}, fn.AddNest("TrafficSplit")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareEndpointTrafficSplitTrafficSplitNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointTrafficSplitTrafficSplit)
	if !ok {
		desiredNotPointer, ok := d.(EndpointTrafficSplitTrafficSplit)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointTrafficSplitTrafficSplit or *EndpointTrafficSplitTrafficSplit", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointTrafficSplitTrafficSplit)
	if !ok {
		actualNotPointer, ok := a.(EndpointTrafficSplitTrafficSplit)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointTrafficSplitTrafficSplit", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DeployedModelId, actual.DeployedModelId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeployedModelId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrafficPercentage, actual.TrafficPercentage, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TrafficPercentage")); len(ds) != 0 || err != nil {
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
func (r *EndpointTrafficSplit) urlNormalized() *EndpointTrafficSplit {
	normalized := dcl.Copy(*r).(EndpointTrafficSplit)
	normalized.Endpoint = dcl.SelfLinkToName(r.Endpoint)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	return &normalized
}

func (r *EndpointTrafficSplit) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateEndpointTrafficSplit" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"endpoint": dcl.ValueOrEmptyString(nr.Endpoint),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the EndpointTrafficSplit resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *EndpointTrafficSplit) marshal(c *Client) ([]byte, error) {
	m, err := expandEndpointTrafficSplit(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling EndpointTrafficSplit: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEndpointTrafficSplit decodes JSON responses into the EndpointTrafficSplit resource schema.
func unmarshalEndpointTrafficSplit(b []byte, c *Client, res *EndpointTrafficSplit) (*EndpointTrafficSplit, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEndpointTrafficSplit(m, c, res)
}

func unmarshalMapEndpointTrafficSplit(m map[string]interface{}, c *Client, res *EndpointTrafficSplit) (*EndpointTrafficSplit, error) {

	flattened := flattenEndpointTrafficSplit(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandEndpointTrafficSplit expands EndpointTrafficSplit into a JSON request object.
func expandEndpointTrafficSplit(c *Client, f *EndpointTrafficSplit) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Endpoint into endpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["endpoint"] = v
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
	if v, err := endpointTrafficSplitSliceToMap(c, f.TrafficSplit, res); err != nil {
		return nil, fmt.Errorf("error expanding TrafficSplit into trafficSplit: %w", err)
	} else if v != nil {
		m["trafficSplit"] = v
	}

	return m, nil
}

// flattenEndpointTrafficSplit flattens EndpointTrafficSplit from a JSON request object into the
// EndpointTrafficSplit type.
func flattenEndpointTrafficSplit(c *Client, i interface{}, res *EndpointTrafficSplit) *EndpointTrafficSplit {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &EndpointTrafficSplit{}
	resultRes.Endpoint = dcl.FlattenString(m["endpoint"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.TrafficSplit = endpointTrafficSplitMapToSlice(c, m["trafficSplit"], res)

	return resultRes
}

// expandEndpointTrafficSplitTrafficSplitMap expands the contents of EndpointTrafficSplitTrafficSplit into a JSON
// request object.
func expandEndpointTrafficSplitTrafficSplitMap(c *Client, f map[string]EndpointTrafficSplitTrafficSplit, res *EndpointTrafficSplit) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointTrafficSplitTrafficSplit(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointTrafficSplitTrafficSplitSlice expands the contents of EndpointTrafficSplitTrafficSplit into a JSON
// request object.
func expandEndpointTrafficSplitTrafficSplitSlice(c *Client, f []EndpointTrafficSplitTrafficSplit, res *EndpointTrafficSplit) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointTrafficSplitTrafficSplit(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointTrafficSplitTrafficSplitMap flattens the contents of EndpointTrafficSplitTrafficSplit from a JSON
// response object.
func flattenEndpointTrafficSplitTrafficSplitMap(c *Client, i interface{}, res *EndpointTrafficSplit) map[string]EndpointTrafficSplitTrafficSplit {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointTrafficSplitTrafficSplit{}
	}

	if len(a) == 0 {
		return map[string]EndpointTrafficSplitTrafficSplit{}
	}

	items := make(map[string]EndpointTrafficSplitTrafficSplit)
	for k, item := range a {
		items[k] = *flattenEndpointTrafficSplitTrafficSplit(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointTrafficSplitTrafficSplitSlice flattens the contents of EndpointTrafficSplitTrafficSplit from a JSON
// response object.
func flattenEndpointTrafficSplitTrafficSplitSlice(c *Client, i interface{}, res *EndpointTrafficSplit) []EndpointTrafficSplitTrafficSplit {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointTrafficSplitTrafficSplit{}
	}

	if len(a) == 0 {
		return []EndpointTrafficSplitTrafficSplit{}
	}

	items := make([]EndpointTrafficSplitTrafficSplit, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointTrafficSplitTrafficSplit(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointTrafficSplitTrafficSplit expands an instance of EndpointTrafficSplitTrafficSplit into a JSON
// request object.
func expandEndpointTrafficSplitTrafficSplit(c *Client, f *EndpointTrafficSplitTrafficSplit, res *EndpointTrafficSplit) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DeployedModelId; !dcl.IsEmptyValueIndirect(v) {
		m["deployedModelId"] = v
	}
	if v := f.TrafficPercentage; !dcl.IsEmptyValueIndirect(v) {
		m["trafficPercentage"] = v
	}

	return m, nil
}

// flattenEndpointTrafficSplitTrafficSplit flattens an instance of EndpointTrafficSplitTrafficSplit from a JSON
// response object.
func flattenEndpointTrafficSplitTrafficSplit(c *Client, i interface{}, res *EndpointTrafficSplit) *EndpointTrafficSplitTrafficSplit {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointTrafficSplitTrafficSplit{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointTrafficSplitTrafficSplit
	}
	r.DeployedModelId = dcl.FlattenString(m["deployedModelId"])
	r.TrafficPercentage = dcl.FlattenInteger(m["trafficPercentage"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *EndpointTrafficSplit) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEndpointTrafficSplit(b, c, r)
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
		if nr.Endpoint == nil && ncr.Endpoint == nil {
			c.Config.Logger.Info("Both Endpoint fields null - considering equal.")
		} else if nr.Endpoint == nil || ncr.Endpoint == nil {
			c.Config.Logger.Info("Only one Endpoint field is null - considering unequal.")
			return false
		} else if *nr.Endpoint != *ncr.Endpoint {
			return false
		}
		return true
	}
}

type endpointTrafficSplitDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         endpointTrafficSplitApiOperation
}

func convertFieldDiffsToEndpointTrafficSplitDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]endpointTrafficSplitDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []endpointTrafficSplitDiff
	// For each operation name, create a endpointTrafficSplitDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := endpointTrafficSplitDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEndpointTrafficSplitApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEndpointTrafficSplitApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (endpointTrafficSplitApiOperation, error) {
	switch opName {

	case "updateEndpointTrafficSplitUpdateEndpointTrafficSplitOperation":
		return &updateEndpointTrafficSplitUpdateEndpointTrafficSplitOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractEndpointTrafficSplitFields(r *EndpointTrafficSplit) error {
	return nil
}
func extractEndpointTrafficSplitTrafficSplitFields(r *EndpointTrafficSplit, o *EndpointTrafficSplitTrafficSplit) error {
	return nil
}

func postReadExtractEndpointTrafficSplitFields(r *EndpointTrafficSplit) error {
	return nil
}
func postReadExtractEndpointTrafficSplitTrafficSplitFields(r *EndpointTrafficSplit, o *EndpointTrafficSplitTrafficSplit) error {
	return nil
}
