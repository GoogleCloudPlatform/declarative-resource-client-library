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
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *ModelDeployment) validate() error {

	if err := dcl.Required(r, "model"); err != nil {
		return err
	}
	if err := dcl.Required(r, "dedicatedResources"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Endpoint, "Endpoint"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DedicatedResources) {
		if err := r.DedicatedResources.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ModelDeploymentDedicatedResources) validate() error {
	if err := dcl.Required(r, "machineSpec"); err != nil {
		return err
	}
	if err := dcl.Required(r, "minReplicaCount"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MachineSpec) {
		if err := r.MachineSpec.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ModelDeploymentDedicatedResourcesMachineSpec) validate() error {
	if err := dcl.Required(r, "machineType"); err != nil {
		return err
	}
	return nil
}
func (r *ModelDeployment) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-aiplatform.googleapis.com/v1beta1/", params)
}

func (r *ModelDeployment) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"endpoint": dcl.ValueOrEmptyString(nr.Endpoint),
		"model":    dcl.ValueOrEmptyString(nr.Model),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}?fetchId=%s", nr.basePath(), userBasePath, params), nil
}

func (r *ModelDeployment) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"endpoint": dcl.ValueOrEmptyString(nr.Endpoint),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}", nr.basePath(), userBasePath, params), nil

}

func (r *ModelDeployment) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"endpoint": dcl.ValueOrEmptyString(nr.Endpoint),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}:deployModel", nr.basePath(), userBasePath, params), nil

}

func (r *ModelDeployment) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"endpoint": dcl.ValueOrEmptyString(nr.Endpoint),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}:undeployModel", nr.basePath(), userBasePath, params), nil
}

// modelDeploymentApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type modelDeploymentApiOperation interface {
	do(context.Context, *ModelDeployment, *Client) error
}

func (c *Client) listModelDeploymentRaw(ctx context.Context, r *ModelDeployment, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ModelDeploymentMaxPage {
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

type listModelDeploymentOperation struct {
	DeployedModels []map[string]interface{} `json:"deployedModels"`
	Token          string                   `json:"nextPageToken"`
}

func (c *Client) listModelDeployment(ctx context.Context, r *ModelDeployment, pageToken string, pageSize int32) ([]*ModelDeployment, string, error) {
	b, err := c.listModelDeploymentRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listModelDeploymentOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ModelDeployment
	for _, v := range m.DeployedModels {
		res, err := unmarshalMapModelDeployment(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		res.Endpoint = r.Endpoint
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllModelDeployment(ctx context.Context, f func(*ModelDeployment) bool, resources []*ModelDeployment) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteModelDeployment(ctx, res)
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

type deleteModelDeploymentOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createModelDeploymentOperation struct {
	response map[string]interface{}
}

func (op *createModelDeploymentOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createModelDeploymentOperation) do(ctx context.Context, r *ModelDeployment, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
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
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetModelDeployment(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getModelDeploymentRaw(ctx context.Context, r *ModelDeployment) ([]byte, error) {

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

	b, err = dcl.ExtractElementFromList(b, "deployedModels", r.matcher(c))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) modelDeploymentDiffsForRawDesired(ctx context.Context, rawDesired *ModelDeployment, opts ...dcl.ApplyOption) (initial, desired *ModelDeployment, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ModelDeployment
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ModelDeployment); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected ModelDeployment, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetModelDeployment(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a ModelDeployment resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ModelDeployment resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that ModelDeployment resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeModelDeploymentDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for ModelDeployment: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for ModelDeployment: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractModelDeploymentFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeModelDeploymentInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for ModelDeployment: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeModelDeploymentDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for ModelDeployment: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffModelDeployment(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeModelDeploymentInitialState(rawInitial, rawDesired *ModelDeployment) (*ModelDeployment, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeModelDeploymentDesiredState(rawDesired, rawInitial *ModelDeployment, opts ...dcl.ApplyOption) (*ModelDeployment, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DedicatedResources = canonicalizeModelDeploymentDedicatedResources(rawDesired.DedicatedResources, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &ModelDeployment{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Model, rawInitial.Model) {
		canonicalDesired.Model = rawInitial.Model
	} else {
		canonicalDesired.Model = rawDesired.Model
	}
	canonicalDesired.DedicatedResources = canonicalizeModelDeploymentDedicatedResources(rawDesired.DedicatedResources, rawInitial.DedicatedResources, opts...)
	if dcl.NameToSelfLink(rawDesired.Endpoint, rawInitial.Endpoint) {
		canonicalDesired.Endpoint = rawInitial.Endpoint
	} else {
		canonicalDesired.Endpoint = rawDesired.Endpoint
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeModelDeploymentNewState(c *Client, rawNew, rawDesired *ModelDeployment) (*ModelDeployment, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Model) && dcl.IsEmptyValueIndirect(rawDesired.Model) {
		rawNew.Model = rawDesired.Model
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Model, rawNew.Model) {
			rawNew.Model = rawDesired.Model
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DedicatedResources) && dcl.IsEmptyValueIndirect(rawDesired.DedicatedResources) {
		rawNew.DedicatedResources = rawDesired.DedicatedResources
	} else {
		rawNew.DedicatedResources = canonicalizeNewModelDeploymentDedicatedResources(c, rawDesired.DedicatedResources, rawNew.DedicatedResources)
	}

	rawNew.Endpoint = rawDesired.Endpoint

	rawNew.Location = rawDesired.Location

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeModelDeploymentDedicatedResources(des, initial *ModelDeploymentDedicatedResources, opts ...dcl.ApplyOption) *ModelDeploymentDedicatedResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ModelDeploymentDedicatedResources{}

	cDes.MachineSpec = canonicalizeModelDeploymentDedicatedResourcesMachineSpec(des.MachineSpec, initial.MachineSpec, opts...)
	if dcl.IsZeroValue(des.MinReplicaCount) || (dcl.IsEmptyValueIndirect(des.MinReplicaCount) && dcl.IsEmptyValueIndirect(initial.MinReplicaCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MinReplicaCount = initial.MinReplicaCount
	} else {
		cDes.MinReplicaCount = des.MinReplicaCount
	}
	if dcl.IsZeroValue(des.MaxReplicaCount) || (dcl.IsEmptyValueIndirect(des.MaxReplicaCount) && dcl.IsEmptyValueIndirect(initial.MaxReplicaCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxReplicaCount = initial.MaxReplicaCount
	} else {
		cDes.MaxReplicaCount = des.MaxReplicaCount
	}

	return cDes
}

func canonicalizeModelDeploymentDedicatedResourcesSlice(des, initial []ModelDeploymentDedicatedResources, opts ...dcl.ApplyOption) []ModelDeploymentDedicatedResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ModelDeploymentDedicatedResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeModelDeploymentDedicatedResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ModelDeploymentDedicatedResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeModelDeploymentDedicatedResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewModelDeploymentDedicatedResources(c *Client, des, nw *ModelDeploymentDedicatedResources) *ModelDeploymentDedicatedResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ModelDeploymentDedicatedResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MachineSpec = canonicalizeNewModelDeploymentDedicatedResourcesMachineSpec(c, des.MachineSpec, nw.MachineSpec)

	return nw
}

func canonicalizeNewModelDeploymentDedicatedResourcesSet(c *Client, des, nw []ModelDeploymentDedicatedResources) []ModelDeploymentDedicatedResources {
	if des == nil {
		return nw
	}
	var reorderedNew []ModelDeploymentDedicatedResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareModelDeploymentDedicatedResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewModelDeploymentDedicatedResourcesSlice(c *Client, des, nw []ModelDeploymentDedicatedResources) []ModelDeploymentDedicatedResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ModelDeploymentDedicatedResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewModelDeploymentDedicatedResources(c, &d, &n))
	}

	return items
}

func canonicalizeModelDeploymentDedicatedResourcesMachineSpec(des, initial *ModelDeploymentDedicatedResourcesMachineSpec, opts ...dcl.ApplyOption) *ModelDeploymentDedicatedResourcesMachineSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ModelDeploymentDedicatedResourcesMachineSpec{}

	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		cDes.MachineType = initial.MachineType
	} else {
		cDes.MachineType = des.MachineType
	}

	return cDes
}

func canonicalizeModelDeploymentDedicatedResourcesMachineSpecSlice(des, initial []ModelDeploymentDedicatedResourcesMachineSpec, opts ...dcl.ApplyOption) []ModelDeploymentDedicatedResourcesMachineSpec {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ModelDeploymentDedicatedResourcesMachineSpec, 0, len(des))
		for _, d := range des {
			cd := canonicalizeModelDeploymentDedicatedResourcesMachineSpec(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ModelDeploymentDedicatedResourcesMachineSpec, 0, len(des))
	for i, d := range des {
		cd := canonicalizeModelDeploymentDedicatedResourcesMachineSpec(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewModelDeploymentDedicatedResourcesMachineSpec(c *Client, des, nw *ModelDeploymentDedicatedResourcesMachineSpec) *ModelDeploymentDedicatedResourcesMachineSpec {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ModelDeploymentDedicatedResourcesMachineSpec while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}

	return nw
}

func canonicalizeNewModelDeploymentDedicatedResourcesMachineSpecSet(c *Client, des, nw []ModelDeploymentDedicatedResourcesMachineSpec) []ModelDeploymentDedicatedResourcesMachineSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []ModelDeploymentDedicatedResourcesMachineSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareModelDeploymentDedicatedResourcesMachineSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewModelDeploymentDedicatedResourcesMachineSpecSlice(c *Client, des, nw []ModelDeploymentDedicatedResourcesMachineSpec) []ModelDeploymentDedicatedResourcesMachineSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ModelDeploymentDedicatedResourcesMachineSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewModelDeploymentDedicatedResourcesMachineSpec(c, &d, &n))
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
func diffModelDeployment(c *Client, desired, actual *ModelDeployment, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Model, actual.Model, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Model")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DedicatedResources, actual.DedicatedResources, dcl.DiffInfo{ObjectFunction: compareModelDeploymentDedicatedResourcesNewStyle, EmptyObject: EmptyModelDeploymentDedicatedResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DedicatedResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Endpoint, actual.Endpoint, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Endpoint")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareModelDeploymentDedicatedResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ModelDeploymentDedicatedResources)
	if !ok {
		desiredNotPointer, ok := d.(ModelDeploymentDedicatedResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelDeploymentDedicatedResources or *ModelDeploymentDedicatedResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ModelDeploymentDedicatedResources)
	if !ok {
		actualNotPointer, ok := a.(ModelDeploymentDedicatedResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelDeploymentDedicatedResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineSpec, actual.MachineSpec, dcl.DiffInfo{ObjectFunction: compareModelDeploymentDedicatedResourcesMachineSpecNewStyle, EmptyObject: EmptyModelDeploymentDedicatedResourcesMachineSpec, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineSpec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinReplicaCount, actual.MinReplicaCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinReplicaCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxReplicaCount, actual.MaxReplicaCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxReplicaCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareModelDeploymentDedicatedResourcesMachineSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ModelDeploymentDedicatedResourcesMachineSpec)
	if !ok {
		desiredNotPointer, ok := d.(ModelDeploymentDedicatedResourcesMachineSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelDeploymentDedicatedResourcesMachineSpec or *ModelDeploymentDedicatedResourcesMachineSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ModelDeploymentDedicatedResourcesMachineSpec)
	if !ok {
		actualNotPointer, ok := a.(ModelDeploymentDedicatedResourcesMachineSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ModelDeploymentDedicatedResourcesMachineSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
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
func (r *ModelDeployment) urlNormalized() *ModelDeployment {
	normalized := dcl.Copy(*r).(ModelDeployment)
	normalized.Model = dcl.SelfLinkToName(r.Model)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.Endpoint = dcl.SelfLinkToName(r.Endpoint)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *ModelDeployment) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ModelDeployment resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ModelDeployment) marshal(c *Client) ([]byte, error) {
	m, err := expandModelDeployment(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ModelDeployment: %w", err)
	}
	m = encodeDeployModelRequest(m)

	return json.Marshal(m)
}

// unmarshalModelDeployment decodes JSON responses into the ModelDeployment resource schema.
func unmarshalModelDeployment(b []byte, c *Client, res *ModelDeployment) (*ModelDeployment, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapModelDeployment(m, c, res)
}

func unmarshalMapModelDeployment(m map[string]interface{}, c *Client, res *ModelDeployment) (*ModelDeployment, error) {

	flattened := flattenModelDeployment(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandModelDeployment expands ModelDeployment into a JSON request object.
func expandModelDeployment(c *Client, f *ModelDeployment) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.ExpandProjectIDsToNumbers(c.Config, f.Model); err != nil {
		return nil, fmt.Errorf("error expanding Model into model: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["model"] = v
	}
	if v, err := expandModelDeploymentDedicatedResources(c, f.DedicatedResources, res); err != nil {
		return nil, fmt.Errorf("error expanding DedicatedResources into dedicatedResources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dedicatedResources"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Endpoint into endpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["endpoint"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenModelDeployment flattens ModelDeployment from a JSON request object into the
// ModelDeployment type.
func flattenModelDeployment(c *Client, i interface{}, res *ModelDeployment) *ModelDeployment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &ModelDeployment{}
	resultRes.Model = dcl.FlattenProjectNumbersToIDs(c.Config, dcl.FlattenString(m["model"]))
	resultRes.Id = dcl.FlattenString(m["id"])
	resultRes.DedicatedResources = flattenModelDeploymentDedicatedResources(c, m["dedicatedResources"], res)
	resultRes.Endpoint = dcl.FlattenString(m["endpoint"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Project = dcl.FlattenString(m["project"])

	return resultRes
}

// expandModelDeploymentDedicatedResourcesMap expands the contents of ModelDeploymentDedicatedResources into a JSON
// request object.
func expandModelDeploymentDedicatedResourcesMap(c *Client, f map[string]ModelDeploymentDedicatedResources, res *ModelDeployment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandModelDeploymentDedicatedResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandModelDeploymentDedicatedResourcesSlice expands the contents of ModelDeploymentDedicatedResources into a JSON
// request object.
func expandModelDeploymentDedicatedResourcesSlice(c *Client, f []ModelDeploymentDedicatedResources, res *ModelDeployment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandModelDeploymentDedicatedResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenModelDeploymentDedicatedResourcesMap flattens the contents of ModelDeploymentDedicatedResources from a JSON
// response object.
func flattenModelDeploymentDedicatedResourcesMap(c *Client, i interface{}, res *ModelDeployment) map[string]ModelDeploymentDedicatedResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelDeploymentDedicatedResources{}
	}

	if len(a) == 0 {
		return map[string]ModelDeploymentDedicatedResources{}
	}

	items := make(map[string]ModelDeploymentDedicatedResources)
	for k, item := range a {
		items[k] = *flattenModelDeploymentDedicatedResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenModelDeploymentDedicatedResourcesSlice flattens the contents of ModelDeploymentDedicatedResources from a JSON
// response object.
func flattenModelDeploymentDedicatedResourcesSlice(c *Client, i interface{}, res *ModelDeployment) []ModelDeploymentDedicatedResources {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelDeploymentDedicatedResources{}
	}

	if len(a) == 0 {
		return []ModelDeploymentDedicatedResources{}
	}

	items := make([]ModelDeploymentDedicatedResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelDeploymentDedicatedResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandModelDeploymentDedicatedResources expands an instance of ModelDeploymentDedicatedResources into a JSON
// request object.
func expandModelDeploymentDedicatedResources(c *Client, f *ModelDeploymentDedicatedResources, res *ModelDeployment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandModelDeploymentDedicatedResourcesMachineSpec(c, f.MachineSpec, res); err != nil {
		return nil, fmt.Errorf("error expanding MachineSpec into machineSpec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["machineSpec"] = v
	}
	if v := f.MinReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["minReplicaCount"] = v
	}
	if v := f.MaxReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxReplicaCount"] = v
	}

	return m, nil
}

// flattenModelDeploymentDedicatedResources flattens an instance of ModelDeploymentDedicatedResources from a JSON
// response object.
func flattenModelDeploymentDedicatedResources(c *Client, i interface{}, res *ModelDeployment) *ModelDeploymentDedicatedResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ModelDeploymentDedicatedResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyModelDeploymentDedicatedResources
	}
	r.MachineSpec = flattenModelDeploymentDedicatedResourcesMachineSpec(c, m["machineSpec"], res)
	r.MinReplicaCount = dcl.FlattenInteger(m["minReplicaCount"])
	r.MaxReplicaCount = dcl.FlattenInteger(m["maxReplicaCount"])

	return r
}

// expandModelDeploymentDedicatedResourcesMachineSpecMap expands the contents of ModelDeploymentDedicatedResourcesMachineSpec into a JSON
// request object.
func expandModelDeploymentDedicatedResourcesMachineSpecMap(c *Client, f map[string]ModelDeploymentDedicatedResourcesMachineSpec, res *ModelDeployment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandModelDeploymentDedicatedResourcesMachineSpec(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandModelDeploymentDedicatedResourcesMachineSpecSlice expands the contents of ModelDeploymentDedicatedResourcesMachineSpec into a JSON
// request object.
func expandModelDeploymentDedicatedResourcesMachineSpecSlice(c *Client, f []ModelDeploymentDedicatedResourcesMachineSpec, res *ModelDeployment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandModelDeploymentDedicatedResourcesMachineSpec(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenModelDeploymentDedicatedResourcesMachineSpecMap flattens the contents of ModelDeploymentDedicatedResourcesMachineSpec from a JSON
// response object.
func flattenModelDeploymentDedicatedResourcesMachineSpecMap(c *Client, i interface{}, res *ModelDeployment) map[string]ModelDeploymentDedicatedResourcesMachineSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ModelDeploymentDedicatedResourcesMachineSpec{}
	}

	if len(a) == 0 {
		return map[string]ModelDeploymentDedicatedResourcesMachineSpec{}
	}

	items := make(map[string]ModelDeploymentDedicatedResourcesMachineSpec)
	for k, item := range a {
		items[k] = *flattenModelDeploymentDedicatedResourcesMachineSpec(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenModelDeploymentDedicatedResourcesMachineSpecSlice flattens the contents of ModelDeploymentDedicatedResourcesMachineSpec from a JSON
// response object.
func flattenModelDeploymentDedicatedResourcesMachineSpecSlice(c *Client, i interface{}, res *ModelDeployment) []ModelDeploymentDedicatedResourcesMachineSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []ModelDeploymentDedicatedResourcesMachineSpec{}
	}

	if len(a) == 0 {
		return []ModelDeploymentDedicatedResourcesMachineSpec{}
	}

	items := make([]ModelDeploymentDedicatedResourcesMachineSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenModelDeploymentDedicatedResourcesMachineSpec(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandModelDeploymentDedicatedResourcesMachineSpec expands an instance of ModelDeploymentDedicatedResourcesMachineSpec into a JSON
// request object.
func expandModelDeploymentDedicatedResourcesMachineSpec(c *Client, f *ModelDeploymentDedicatedResourcesMachineSpec, res *ModelDeployment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}

	return m, nil
}

// flattenModelDeploymentDedicatedResourcesMachineSpec flattens an instance of ModelDeploymentDedicatedResourcesMachineSpec from a JSON
// response object.
func flattenModelDeploymentDedicatedResourcesMachineSpec(c *Client, i interface{}, res *ModelDeployment) *ModelDeploymentDedicatedResourcesMachineSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ModelDeploymentDedicatedResourcesMachineSpec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyModelDeploymentDedicatedResourcesMachineSpec
	}
	r.MachineType = dcl.FlattenString(m["machineType"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ModelDeployment) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalModelDeployment(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Model == nil && ncr.Model == nil {
			c.Config.Logger.Info("Both Model fields null - considering equal.")
		} else if nr.Model == nil || ncr.Model == nil {
			c.Config.Logger.Info("Only one Model field is null - considering unequal.")
			return false
		} else if *nr.Model != *ncr.Model {
			return false
		}
		return true
	}
}

type modelDeploymentDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         modelDeploymentApiOperation
}

func convertFieldDiffsToModelDeploymentDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]modelDeploymentDiff, error) {
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
	var diffs []modelDeploymentDiff
	// For each operation name, create a modelDeploymentDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := modelDeploymentDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToModelDeploymentApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToModelDeploymentApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (modelDeploymentApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractModelDeploymentFields(r *ModelDeployment) error {
	vDedicatedResources := r.DedicatedResources
	if vDedicatedResources == nil {
		// note: explicitly not the empty object.
		vDedicatedResources = &ModelDeploymentDedicatedResources{}
	}
	if err := extractModelDeploymentDedicatedResourcesFields(r, vDedicatedResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDedicatedResources) {
		r.DedicatedResources = vDedicatedResources
	}
	vLocation, err := dcl.ValueFromRegexOnField("Location", r.Location, r.Endpoint, "projects/.*/locations/([a-z0-9A-Z-]*)/endpoints/.*")
	if err != nil {
		return err
	}
	r.Location = vLocation
	vProject, err := dcl.ValueFromRegexOnField("Project", r.Project, r.Endpoint, "projects/([a-z0-9A-Z-]*)/locations/.*/endpoints/.*")
	if err != nil {
		return err
	}
	r.Project = vProject
	return nil
}
func extractModelDeploymentDedicatedResourcesFields(r *ModelDeployment, o *ModelDeploymentDedicatedResources) error {
	vMachineSpec := o.MachineSpec
	if vMachineSpec == nil {
		// note: explicitly not the empty object.
		vMachineSpec = &ModelDeploymentDedicatedResourcesMachineSpec{}
	}
	if err := extractModelDeploymentDedicatedResourcesMachineSpecFields(r, vMachineSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMachineSpec) {
		o.MachineSpec = vMachineSpec
	}
	return nil
}
func extractModelDeploymentDedicatedResourcesMachineSpecFields(r *ModelDeployment, o *ModelDeploymentDedicatedResourcesMachineSpec) error {
	return nil
}

func postReadExtractModelDeploymentFields(r *ModelDeployment) error {
	vDedicatedResources := r.DedicatedResources
	if vDedicatedResources == nil {
		// note: explicitly not the empty object.
		vDedicatedResources = &ModelDeploymentDedicatedResources{}
	}
	if err := postReadExtractModelDeploymentDedicatedResourcesFields(r, vDedicatedResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDedicatedResources) {
		r.DedicatedResources = vDedicatedResources
	}
	return nil
}
func postReadExtractModelDeploymentDedicatedResourcesFields(r *ModelDeployment, o *ModelDeploymentDedicatedResources) error {
	vMachineSpec := o.MachineSpec
	if vMachineSpec == nil {
		// note: explicitly not the empty object.
		vMachineSpec = &ModelDeploymentDedicatedResourcesMachineSpec{}
	}
	if err := extractModelDeploymentDedicatedResourcesMachineSpecFields(r, vMachineSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMachineSpec) {
		o.MachineSpec = vMachineSpec
	}
	return nil
}
func postReadExtractModelDeploymentDedicatedResourcesMachineSpecFields(r *ModelDeployment, o *ModelDeploymentDedicatedResourcesMachineSpec) error {
	return nil
}
