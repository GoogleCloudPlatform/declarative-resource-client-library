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
)

func (r *Routine) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "dataset"); err != nil {
		return err
	}
	if err := dcl.Required(r, "routineType"); err != nil {
		return err
	}
	if err := dcl.Required(r, "definitionBody"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ReturnType) {
		if err := r.ReturnType.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *RoutineArguments) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DataType) {
		if err := r.DataType.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *RoutineArgumentsDataType) validate() error {
	if err := dcl.Required(r, "typeKind"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"ArrayElementType", "StructType"}, r.ArrayElementType, r.StructType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ArrayElementType) {
		if err := r.ArrayElementType.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.StructType) {
		if err := r.StructType.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *RoutineArgumentsDataTypeStructType) validate() error {
	return nil
}
func (r *RoutineArgumentsDataTypeStructTypeFields) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Type) {
		if err := r.Type.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *Routine) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://bigquery.googleapis.com/bigquery/v2/", params)
}

func (r *Routine) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"dataset": dcl.ValueOrEmptyString(nr.Dataset),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/datasets/{{dataset}}/routines/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Routine) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"dataset": dcl.ValueOrEmptyString(nr.Dataset),
	}
	return dcl.URL("projects/{{project}}/datasets/{{dataset}}/routines", nr.basePath(), userBasePath, params), nil

}

func (r *Routine) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"dataset": dcl.ValueOrEmptyString(nr.Dataset),
	}
	return dcl.URL("projects/{{project}}/datasets/{{dataset}}/routines", nr.basePath(), userBasePath, params), nil

}

func (r *Routine) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"dataset": dcl.ValueOrEmptyString(nr.Dataset),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/datasets/{{dataset}}/routines/{{name}}", nr.basePath(), userBasePath, params), nil
}

// routineApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type routineApiOperation interface {
	do(context.Context, *Routine, *Client) error
}

// newUpdateRoutinePatchRoutineRequest creates a request for an
// Routine resource's PatchRoutine update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateRoutinePatchRoutineRequest(ctx context.Context, f *Routine, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Project; !dcl.IsEmptyValueIndirect(v) {
		req["project"] = v
	}
	if v := f.Dataset; !dcl.IsEmptyValueIndirect(v) {
		req["dataset"] = v
	}
	if v := f.RoutineType; !dcl.IsEmptyValueIndirect(v) {
		req["routineType"] = v
	}
	if v := f.Language; !dcl.IsEmptyValueIndirect(v) {
		req["language"] = v
	}
	if v, err := expandRoutineArgumentsSlice(c, f.Arguments, res); err != nil {
		return nil, fmt.Errorf("error expanding Arguments into arguments: %w", err)
	} else if v != nil {
		req["arguments"] = v
	}
	if v, err := expandRoutineArgumentsDataType(c, f.ReturnType, res); err != nil {
		return nil, fmt.Errorf("error expanding ReturnType into returnType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["returnType"] = v
	}
	if v := f.ImportedLibraries; v != nil {
		req["importedLibraries"] = v
	}
	if v := f.DefinitionBody; !dcl.IsEmptyValueIndirect(v) {
		req["definitionBody"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.DeterminismLevel; !dcl.IsEmptyValueIndirect(v) {
		req["determinismLevel"] = v
	}
	if v := f.StrictMode; !dcl.IsEmptyValueIndirect(v) {
		req["strictMode"] = v
	}
	b, err := c.getRoutineRaw(ctx, f)
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

// marshalUpdateRoutinePatchRoutineRequest converts the update into
// the final JSON request body.
func marshalUpdateRoutinePatchRoutineRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"routineReference", "routineId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{"routineReference", "projectId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"dataset"},
		[]string{"routineReference", "datasetId"},
	)
	return json.Marshal(m)
}

type updateRoutinePatchRoutineOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateRoutinePatchRoutineOperation) do(ctx context.Context, r *Routine, c *Client) error {
	_, err := c.GetRoutine(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchRoutine")
	if err != nil {
		return err
	}

	req, err := newUpdateRoutinePatchRoutineRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateRoutinePatchRoutineRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listRoutineRaw(ctx context.Context, r *Routine, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != RoutineMaxPage {
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

type listRoutineOperation struct {
	Routines []map[string]interface{} `json:"routines"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listRoutine(ctx context.Context, r *Routine, pageToken string, pageSize int32) ([]*Routine, string, error) {
	b, err := c.listRoutineRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listRoutineOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Routine
	for _, v := range m.Routines {
		res, err := unmarshalMapRoutine(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Dataset = r.Dataset
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllRoutine(ctx context.Context, f func(*Routine) bool, resources []*Routine) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteRoutine(ctx, res)
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

type deleteRoutineOperation struct{}

func (op *deleteRoutineOperation) do(ctx context.Context, r *Routine, c *Client) error {
	r, err := c.GetRoutine(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Routine not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetRoutine checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete Routine: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetRoutine(ctx, r)
		if dcl.IsNotFound(err) {
			return nil, nil
		}
		if retriesRemaining > 0 {
			retriesRemaining--
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, dcl.NotDeletedError{ExistingResource: r}
	}, c.Config.RetryProvider)
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createRoutineOperation struct {
	response map[string]interface{}
}

func (op *createRoutineOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createRoutineOperation) do(ctx context.Context, r *Routine, c *Client) error {
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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetRoutine(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getRoutineRaw(ctx context.Context, r *Routine) ([]byte, error) {

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

func (c *Client) routineDiffsForRawDesired(ctx context.Context, rawDesired *Routine, opts ...dcl.ApplyOption) (initial, desired *Routine, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Routine
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Routine); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Routine, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetRoutine(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Routine resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Routine resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Routine resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeRoutineDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Routine: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Routine: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractRoutineFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeRoutineInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Routine: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeRoutineDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Routine: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffRoutine(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeRoutineInitialState(rawInitial, rawDesired *Routine) (*Routine, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeRoutineDesiredState(rawDesired, rawInitial *Routine, opts ...dcl.ApplyOption) (*Routine, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ReturnType = canonicalizeRoutineArgumentsDataType(rawDesired.ReturnType, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Routine{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Project) || (dcl.IsEmptyValueIndirect(rawDesired.Project) && dcl.IsEmptyValueIndirect(rawInitial.Project)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.IsZeroValue(rawDesired.Dataset) || (dcl.IsEmptyValueIndirect(rawDesired.Dataset) && dcl.IsEmptyValueIndirect(rawInitial.Dataset)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Dataset = rawInitial.Dataset
	} else {
		canonicalDesired.Dataset = rawDesired.Dataset
	}
	if dcl.IsZeroValue(rawDesired.RoutineType) || (dcl.IsEmptyValueIndirect(rawDesired.RoutineType) && dcl.IsEmptyValueIndirect(rawInitial.RoutineType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.RoutineType = rawInitial.RoutineType
	} else {
		canonicalDesired.RoutineType = rawDesired.RoutineType
	}
	if dcl.IsZeroValue(rawDesired.Language) || (dcl.IsEmptyValueIndirect(rawDesired.Language) && dcl.IsEmptyValueIndirect(rawInitial.Language)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Language = rawInitial.Language
	} else {
		canonicalDesired.Language = rawDesired.Language
	}
	canonicalDesired.Arguments = canonicalizeRoutineArgumentsSlice(rawDesired.Arguments, rawInitial.Arguments, opts...)
	canonicalDesired.ReturnType = canonicalizeRoutineArgumentsDataType(rawDesired.ReturnType, rawInitial.ReturnType, opts...)
	if dcl.StringArrayCanonicalize(rawDesired.ImportedLibraries, rawInitial.ImportedLibraries) {
		canonicalDesired.ImportedLibraries = rawInitial.ImportedLibraries
	} else {
		canonicalDesired.ImportedLibraries = rawDesired.ImportedLibraries
	}
	if dcl.StringCanonicalize(rawDesired.DefinitionBody, rawInitial.DefinitionBody) {
		canonicalDesired.DefinitionBody = rawInitial.DefinitionBody
	} else {
		canonicalDesired.DefinitionBody = rawDesired.DefinitionBody
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.DeterminismLevel) || (dcl.IsEmptyValueIndirect(rawDesired.DeterminismLevel) && dcl.IsEmptyValueIndirect(rawInitial.DeterminismLevel)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.DeterminismLevel = rawInitial.DeterminismLevel
	} else {
		canonicalDesired.DeterminismLevel = rawDesired.DeterminismLevel
	}
	if dcl.BoolCanonicalize(rawDesired.StrictMode, rawInitial.StrictMode) {
		canonicalDesired.StrictMode = rawInitial.StrictMode
	} else {
		canonicalDesired.StrictMode = rawDesired.StrictMode
	}

	return canonicalDesired, nil
}

func canonicalizeRoutineNewState(c *Client, rawNew, rawDesired *Routine) (*Routine, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Dataset) && dcl.IsEmptyValueIndirect(rawDesired.Dataset) {
		rawNew.Dataset = rawDesired.Dataset
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RoutineType) && dcl.IsEmptyValueIndirect(rawDesired.RoutineType) {
		rawNew.RoutineType = rawDesired.RoutineType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTime) && dcl.IsEmptyValueIndirect(rawDesired.CreationTime) {
		rawNew.CreationTime = rawDesired.CreationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifiedTime) && dcl.IsEmptyValueIndirect(rawDesired.LastModifiedTime) {
		rawNew.LastModifiedTime = rawDesired.LastModifiedTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Language) && dcl.IsEmptyValueIndirect(rawDesired.Language) {
		rawNew.Language = rawDesired.Language
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Arguments) && dcl.IsEmptyValueIndirect(rawDesired.Arguments) {
		rawNew.Arguments = rawDesired.Arguments
	} else {
		rawNew.Arguments = canonicalizeNewRoutineArgumentsSlice(c, rawDesired.Arguments, rawNew.Arguments)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReturnType) && dcl.IsEmptyValueIndirect(rawDesired.ReturnType) {
		rawNew.ReturnType = rawDesired.ReturnType
	} else {
		rawNew.ReturnType = canonicalizeNewRoutineArgumentsDataType(c, rawDesired.ReturnType, rawNew.ReturnType)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ImportedLibraries) && dcl.IsEmptyValueIndirect(rawDesired.ImportedLibraries) {
		rawNew.ImportedLibraries = rawDesired.ImportedLibraries
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.ImportedLibraries, rawNew.ImportedLibraries) {
			rawNew.ImportedLibraries = rawDesired.ImportedLibraries
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefinitionBody) && dcl.IsEmptyValueIndirect(rawDesired.DefinitionBody) {
		rawNew.DefinitionBody = rawDesired.DefinitionBody
	} else {
		if dcl.StringCanonicalize(rawDesired.DefinitionBody, rawNew.DefinitionBody) {
			rawNew.DefinitionBody = rawDesired.DefinitionBody
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeterminismLevel) && dcl.IsEmptyValueIndirect(rawDesired.DeterminismLevel) {
		rawNew.DeterminismLevel = rawDesired.DeterminismLevel
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StrictMode) && dcl.IsEmptyValueIndirect(rawDesired.StrictMode) {
		rawNew.StrictMode = rawDesired.StrictMode
	} else {
		if dcl.BoolCanonicalize(rawDesired.StrictMode, rawNew.StrictMode) {
			rawNew.StrictMode = rawDesired.StrictMode
		}
	}

	return rawNew, nil
}

func canonicalizeRoutineArguments(des, initial *RoutineArguments, opts ...dcl.ApplyOption) *RoutineArguments {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &RoutineArguments{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.ArgumentKind) || (dcl.IsEmptyValueIndirect(des.ArgumentKind) && dcl.IsEmptyValueIndirect(initial.ArgumentKind)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ArgumentKind = initial.ArgumentKind
	} else {
		cDes.ArgumentKind = des.ArgumentKind
	}
	if dcl.IsZeroValue(des.Mode) || (dcl.IsEmptyValueIndirect(des.Mode) && dcl.IsEmptyValueIndirect(initial.Mode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Mode = initial.Mode
	} else {
		cDes.Mode = des.Mode
	}
	cDes.DataType = canonicalizeRoutineArgumentsDataType(des.DataType, initial.DataType, opts...)

	return cDes
}

func canonicalizeRoutineArgumentsSlice(des, initial []RoutineArguments, opts ...dcl.ApplyOption) []RoutineArguments {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]RoutineArguments, 0, len(des))
		for _, d := range des {
			cd := canonicalizeRoutineArguments(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]RoutineArguments, 0, len(des))
	for i, d := range des {
		cd := canonicalizeRoutineArguments(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewRoutineArguments(c *Client, des, nw *RoutineArguments) *RoutineArguments {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for RoutineArguments while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.DataType = canonicalizeNewRoutineArgumentsDataType(c, des.DataType, nw.DataType)

	return nw
}

func canonicalizeNewRoutineArgumentsSet(c *Client, des, nw []RoutineArguments) []RoutineArguments {
	if des == nil {
		return nw
	}
	var reorderedNew []RoutineArguments
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareRoutineArgumentsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewRoutineArgumentsSlice(c *Client, des, nw []RoutineArguments) []RoutineArguments {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []RoutineArguments
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRoutineArguments(c, &d, &n))
	}

	return items
}

func canonicalizeRoutineArgumentsDataType(des, initial *RoutineArgumentsDataType, opts ...dcl.ApplyOption) *RoutineArgumentsDataType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.ArrayElementType != nil || (initial != nil && initial.ArrayElementType != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.StructType) {
			des.ArrayElementType = nil
			if initial != nil {
				initial.ArrayElementType = nil
			}
		}
	}

	if des.StructType != nil || (initial != nil && initial.StructType != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ArrayElementType) {
			des.StructType = nil
			if initial != nil {
				initial.StructType = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &RoutineArgumentsDataType{}

	if dcl.IsZeroValue(des.TypeKind) || (dcl.IsEmptyValueIndirect(des.TypeKind) && dcl.IsEmptyValueIndirect(initial.TypeKind)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.TypeKind = initial.TypeKind
	} else {
		cDes.TypeKind = des.TypeKind
	}
	cDes.ArrayElementType = canonicalizeRoutineArgumentsDataType(des.ArrayElementType, initial.ArrayElementType, opts...)
	cDes.StructType = canonicalizeRoutineArgumentsDataTypeStructType(des.StructType, initial.StructType, opts...)

	return cDes
}

func canonicalizeRoutineArgumentsDataTypeSlice(des, initial []RoutineArgumentsDataType, opts ...dcl.ApplyOption) []RoutineArgumentsDataType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]RoutineArgumentsDataType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeRoutineArgumentsDataType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]RoutineArgumentsDataType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeRoutineArgumentsDataType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewRoutineArgumentsDataType(c *Client, des, nw *RoutineArgumentsDataType) *RoutineArgumentsDataType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for RoutineArgumentsDataType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.ArrayElementType = canonicalizeNewRoutineArgumentsDataType(c, des.ArrayElementType, nw.ArrayElementType)
	nw.StructType = canonicalizeNewRoutineArgumentsDataTypeStructType(c, des.StructType, nw.StructType)

	return nw
}

func canonicalizeNewRoutineArgumentsDataTypeSet(c *Client, des, nw []RoutineArgumentsDataType) []RoutineArgumentsDataType {
	if des == nil {
		return nw
	}
	var reorderedNew []RoutineArgumentsDataType
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareRoutineArgumentsDataTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewRoutineArgumentsDataTypeSlice(c *Client, des, nw []RoutineArgumentsDataType) []RoutineArgumentsDataType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []RoutineArgumentsDataType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRoutineArgumentsDataType(c, &d, &n))
	}

	return items
}

func canonicalizeRoutineArgumentsDataTypeStructType(des, initial *RoutineArgumentsDataTypeStructType, opts ...dcl.ApplyOption) *RoutineArgumentsDataTypeStructType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &RoutineArgumentsDataTypeStructType{}

	cDes.Fields = canonicalizeRoutineArgumentsDataTypeStructTypeFieldsSlice(des.Fields, initial.Fields, opts...)

	return cDes
}

func canonicalizeRoutineArgumentsDataTypeStructTypeSlice(des, initial []RoutineArgumentsDataTypeStructType, opts ...dcl.ApplyOption) []RoutineArgumentsDataTypeStructType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]RoutineArgumentsDataTypeStructType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeRoutineArgumentsDataTypeStructType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]RoutineArgumentsDataTypeStructType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeRoutineArgumentsDataTypeStructType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewRoutineArgumentsDataTypeStructType(c *Client, des, nw *RoutineArgumentsDataTypeStructType) *RoutineArgumentsDataTypeStructType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for RoutineArgumentsDataTypeStructType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Fields = canonicalizeNewRoutineArgumentsDataTypeStructTypeFieldsSlice(c, des.Fields, nw.Fields)

	return nw
}

func canonicalizeNewRoutineArgumentsDataTypeStructTypeSet(c *Client, des, nw []RoutineArgumentsDataTypeStructType) []RoutineArgumentsDataTypeStructType {
	if des == nil {
		return nw
	}
	var reorderedNew []RoutineArgumentsDataTypeStructType
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareRoutineArgumentsDataTypeStructTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewRoutineArgumentsDataTypeStructTypeSlice(c *Client, des, nw []RoutineArgumentsDataTypeStructType) []RoutineArgumentsDataTypeStructType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []RoutineArgumentsDataTypeStructType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRoutineArgumentsDataTypeStructType(c, &d, &n))
	}

	return items
}

func canonicalizeRoutineArgumentsDataTypeStructTypeFields(des, initial *RoutineArgumentsDataTypeStructTypeFields, opts ...dcl.ApplyOption) *RoutineArgumentsDataTypeStructTypeFields {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &RoutineArgumentsDataTypeStructTypeFields{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	cDes.Type = canonicalizeRoutineArgumentsDataType(des.Type, initial.Type, opts...)

	return cDes
}

func canonicalizeRoutineArgumentsDataTypeStructTypeFieldsSlice(des, initial []RoutineArgumentsDataTypeStructTypeFields, opts ...dcl.ApplyOption) []RoutineArgumentsDataTypeStructTypeFields {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]RoutineArgumentsDataTypeStructTypeFields, 0, len(des))
		for _, d := range des {
			cd := canonicalizeRoutineArgumentsDataTypeStructTypeFields(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]RoutineArgumentsDataTypeStructTypeFields, 0, len(des))
	for i, d := range des {
		cd := canonicalizeRoutineArgumentsDataTypeStructTypeFields(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewRoutineArgumentsDataTypeStructTypeFields(c *Client, des, nw *RoutineArgumentsDataTypeStructTypeFields) *RoutineArgumentsDataTypeStructTypeFields {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for RoutineArgumentsDataTypeStructTypeFields while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.Type = canonicalizeNewRoutineArgumentsDataType(c, des.Type, nw.Type)

	return nw
}

func canonicalizeNewRoutineArgumentsDataTypeStructTypeFieldsSet(c *Client, des, nw []RoutineArgumentsDataTypeStructTypeFields) []RoutineArgumentsDataTypeStructTypeFields {
	if des == nil {
		return nw
	}
	var reorderedNew []RoutineArgumentsDataTypeStructTypeFields
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareRoutineArgumentsDataTypeStructTypeFieldsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewRoutineArgumentsDataTypeStructTypeFieldsSlice(c *Client, des, nw []RoutineArgumentsDataTypeStructTypeFields) []RoutineArgumentsDataTypeStructTypeFields {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []RoutineArgumentsDataTypeStructTypeFields
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRoutineArgumentsDataTypeStructTypeFields(c, &d, &n))
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
func diffRoutine(c *Client, desired, actual *Routine, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dataset, actual.Dataset, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("Dataset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RoutineType, actual.RoutineType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("RoutineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreationTime, actual.CreationTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifiedTime, actual.LastModifiedTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifiedTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Language, actual.Language, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("Language")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Arguments, actual.Arguments, dcl.DiffInfo{ObjectFunction: compareRoutineArgumentsNewStyle, EmptyObject: EmptyRoutineArguments, OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("Arguments")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReturnType, actual.ReturnType, dcl.DiffInfo{ObjectFunction: compareRoutineArgumentsDataTypeNewStyle, EmptyObject: EmptyRoutineArgumentsDataType, OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("ReturnType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImportedLibraries, actual.ImportedLibraries, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("ImportedLibraries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefinitionBody, actual.DefinitionBody, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("DefinitionBody")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeterminismLevel, actual.DeterminismLevel, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("DeterminismLevel")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StrictMode, actual.StrictMode, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateRoutinePatchRoutineOperation")}, fn.AddNest("StrictMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareRoutineArgumentsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*RoutineArguments)
	if !ok {
		desiredNotPointer, ok := d.(RoutineArguments)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoutineArguments or *RoutineArguments", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*RoutineArguments)
	if !ok {
		actualNotPointer, ok := a.(RoutineArguments)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoutineArguments", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ArgumentKind, actual.ArgumentKind, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArgumentKind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DataType, actual.DataType, dcl.DiffInfo{ObjectFunction: compareRoutineArgumentsDataTypeNewStyle, EmptyObject: EmptyRoutineArgumentsDataType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DataType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareRoutineArgumentsDataTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*RoutineArgumentsDataType)
	if !ok {
		desiredNotPointer, ok := d.(RoutineArgumentsDataType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoutineArgumentsDataType or *RoutineArgumentsDataType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*RoutineArgumentsDataType)
	if !ok {
		actualNotPointer, ok := a.(RoutineArgumentsDataType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoutineArgumentsDataType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TypeKind, actual.TypeKind, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TypeKind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ArrayElementType, actual.ArrayElementType, dcl.DiffInfo{ObjectFunction: compareRoutineArgumentsDataTypeNewStyle, EmptyObject: EmptyRoutineArgumentsDataType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ArrayElementType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StructType, actual.StructType, dcl.DiffInfo{ObjectFunction: compareRoutineArgumentsDataTypeStructTypeNewStyle, EmptyObject: EmptyRoutineArgumentsDataTypeStructType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StructType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareRoutineArgumentsDataTypeStructTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*RoutineArgumentsDataTypeStructType)
	if !ok {
		desiredNotPointer, ok := d.(RoutineArgumentsDataTypeStructType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoutineArgumentsDataTypeStructType or *RoutineArgumentsDataTypeStructType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*RoutineArgumentsDataTypeStructType)
	if !ok {
		actualNotPointer, ok := a.(RoutineArgumentsDataTypeStructType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoutineArgumentsDataTypeStructType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fields, actual.Fields, dcl.DiffInfo{ObjectFunction: compareRoutineArgumentsDataTypeStructTypeFieldsNewStyle, EmptyObject: EmptyRoutineArgumentsDataTypeStructTypeFields, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Fields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareRoutineArgumentsDataTypeStructTypeFieldsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*RoutineArgumentsDataTypeStructTypeFields)
	if !ok {
		desiredNotPointer, ok := d.(RoutineArgumentsDataTypeStructTypeFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoutineArgumentsDataTypeStructTypeFields or *RoutineArgumentsDataTypeStructTypeFields", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*RoutineArgumentsDataTypeStructTypeFields)
	if !ok {
		actualNotPointer, ok := a.(RoutineArgumentsDataTypeStructTypeFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoutineArgumentsDataTypeStructTypeFields", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{ObjectFunction: compareRoutineArgumentsDataTypeNewStyle, EmptyObject: EmptyRoutineArgumentsDataType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
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
func (r *Routine) urlNormalized() *Routine {
	normalized := dcl.Copy(*r).(Routine)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Dataset = dcl.SelfLinkToName(r.Dataset)
	normalized.DefinitionBody = dcl.SelfLinkToName(r.DefinitionBody)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	return &normalized
}

func (r *Routine) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "PatchRoutine" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"dataset": dcl.ValueOrEmptyString(nr.Dataset),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/datasets/{{dataset}}/routines/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Routine resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Routine) marshal(c *Client) ([]byte, error) {
	m, err := expandRoutine(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Routine: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"routineReference", "routineId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{"routineReference", "projectId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"dataset"},
		[]string{"routineReference", "datasetId"},
	)

	return json.Marshal(m)
}

// unmarshalRoutine decodes JSON responses into the Routine resource schema.
func unmarshalRoutine(b []byte, c *Client, res *Routine) (*Routine, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapRoutine(m, c, res)
}

func unmarshalMapRoutine(m map[string]interface{}, c *Client, res *Routine) (*Routine, error) {
	dcl.MoveMapEntry(
		m,
		[]string{"routineReference", "routineId"},
		[]string{"name"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"routineReference", "projectId"},
		[]string{"project"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"routineReference", "datasetId"},
		[]string{"dataset"},
	)

	flattened := flattenRoutine(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandRoutine expands Routine into a JSON request object.
func expandRoutine(c *Client, f *Routine) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.Project; dcl.ValueShouldBeSent(v) {
		m["project"] = v
	}
	if v := f.Dataset; dcl.ValueShouldBeSent(v) {
		m["dataset"] = v
	}
	if v := f.RoutineType; dcl.ValueShouldBeSent(v) {
		m["routineType"] = v
	}
	if v := f.Language; dcl.ValueShouldBeSent(v) {
		m["language"] = v
	}
	if v, err := expandRoutineArgumentsSlice(c, f.Arguments, res); err != nil {
		return nil, fmt.Errorf("error expanding Arguments into arguments: %w", err)
	} else if v != nil {
		m["arguments"] = v
	}
	if v, err := expandRoutineArgumentsDataType(c, f.ReturnType, res); err != nil {
		return nil, fmt.Errorf("error expanding ReturnType into returnType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["returnType"] = v
	}
	if v := f.ImportedLibraries; v != nil {
		m["importedLibraries"] = v
	}
	if v := f.DefinitionBody; dcl.ValueShouldBeSent(v) {
		m["definitionBody"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.DeterminismLevel; dcl.ValueShouldBeSent(v) {
		m["determinismLevel"] = v
	}
	if v := f.StrictMode; dcl.ValueShouldBeSent(v) {
		m["strictMode"] = v
	}

	return m, nil
}

// flattenRoutine flattens Routine from a JSON request object into the
// Routine type.
func flattenRoutine(c *Client, i interface{}, res *Routine) *Routine {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Routine{}
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Dataset = dcl.FlattenString(m["dataset"])
	resultRes.RoutineType = flattenRoutineRoutineTypeEnum(m["routineType"])
	resultRes.CreationTime = dcl.FlattenInteger(m["creationTime"])
	resultRes.LastModifiedTime = dcl.FlattenInteger(m["lastModifiedTime"])
	resultRes.Language = flattenRoutineLanguageEnum(m["language"])
	resultRes.Arguments = flattenRoutineArgumentsSlice(c, m["arguments"], res)
	resultRes.ReturnType = flattenRoutineArgumentsDataType(c, m["returnType"], res)
	resultRes.ImportedLibraries = dcl.FlattenStringSlice(m["importedLibraries"])
	resultRes.DefinitionBody = dcl.FlattenString(m["definitionBody"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.DeterminismLevel = flattenRoutineDeterminismLevelEnum(m["determinismLevel"])
	resultRes.StrictMode = dcl.FlattenBool(m["strictMode"])

	return resultRes
}

// expandRoutineArgumentsMap expands the contents of RoutineArguments into a JSON
// request object.
func expandRoutineArgumentsMap(c *Client, f map[string]RoutineArguments, res *Routine) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRoutineArguments(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRoutineArgumentsSlice expands the contents of RoutineArguments into a JSON
// request object.
func expandRoutineArgumentsSlice(c *Client, f []RoutineArguments, res *Routine) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRoutineArguments(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRoutineArgumentsMap flattens the contents of RoutineArguments from a JSON
// response object.
func flattenRoutineArgumentsMap(c *Client, i interface{}, res *Routine) map[string]RoutineArguments {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineArguments{}
	}

	if len(a) == 0 {
		return map[string]RoutineArguments{}
	}

	items := make(map[string]RoutineArguments)
	for k, item := range a {
		items[k] = *flattenRoutineArguments(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenRoutineArgumentsSlice flattens the contents of RoutineArguments from a JSON
// response object.
func flattenRoutineArgumentsSlice(c *Client, i interface{}, res *Routine) []RoutineArguments {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineArguments{}
	}

	if len(a) == 0 {
		return []RoutineArguments{}
	}

	items := make([]RoutineArguments, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineArguments(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandRoutineArguments expands an instance of RoutineArguments into a JSON
// request object.
func expandRoutineArguments(c *Client, f *RoutineArguments, res *Routine) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.ArgumentKind; !dcl.IsEmptyValueIndirect(v) {
		m["argumentKind"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}
	if v, err := expandRoutineArgumentsDataType(c, f.DataType, res); err != nil {
		return nil, fmt.Errorf("error expanding DataType into dataType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dataType"] = v
	}

	return m, nil
}

// flattenRoutineArguments flattens an instance of RoutineArguments from a JSON
// response object.
func flattenRoutineArguments(c *Client, i interface{}, res *Routine) *RoutineArguments {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RoutineArguments{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyRoutineArguments
	}
	r.Name = dcl.FlattenString(m["name"])
	r.ArgumentKind = flattenRoutineArgumentsArgumentKindEnum(m["argumentKind"])
	r.Mode = flattenRoutineArgumentsModeEnum(m["mode"])
	r.DataType = flattenRoutineArgumentsDataType(c, m["dataType"], res)

	return r
}

// expandRoutineArgumentsDataTypeMap expands the contents of RoutineArgumentsDataType into a JSON
// request object.
func expandRoutineArgumentsDataTypeMap(c *Client, f map[string]RoutineArgumentsDataType, res *Routine) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRoutineArgumentsDataType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRoutineArgumentsDataTypeSlice expands the contents of RoutineArgumentsDataType into a JSON
// request object.
func expandRoutineArgumentsDataTypeSlice(c *Client, f []RoutineArgumentsDataType, res *Routine) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRoutineArgumentsDataType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRoutineArgumentsDataTypeMap flattens the contents of RoutineArgumentsDataType from a JSON
// response object.
func flattenRoutineArgumentsDataTypeMap(c *Client, i interface{}, res *Routine) map[string]RoutineArgumentsDataType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineArgumentsDataType{}
	}

	if len(a) == 0 {
		return map[string]RoutineArgumentsDataType{}
	}

	items := make(map[string]RoutineArgumentsDataType)
	for k, item := range a {
		items[k] = *flattenRoutineArgumentsDataType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenRoutineArgumentsDataTypeSlice flattens the contents of RoutineArgumentsDataType from a JSON
// response object.
func flattenRoutineArgumentsDataTypeSlice(c *Client, i interface{}, res *Routine) []RoutineArgumentsDataType {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineArgumentsDataType{}
	}

	if len(a) == 0 {
		return []RoutineArgumentsDataType{}
	}

	items := make([]RoutineArgumentsDataType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineArgumentsDataType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandRoutineArgumentsDataType expands an instance of RoutineArgumentsDataType into a JSON
// request object.
func expandRoutineArgumentsDataType(c *Client, f *RoutineArgumentsDataType, res *Routine) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TypeKind; !dcl.IsEmptyValueIndirect(v) {
		m["typeKind"] = v
	}
	if v, err := expandRoutineArgumentsDataType(c, f.ArrayElementType, res); err != nil {
		return nil, fmt.Errorf("error expanding ArrayElementType into arrayElementType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["arrayElementType"] = v
	}
	if v, err := expandRoutineArgumentsDataTypeStructType(c, f.StructType, res); err != nil {
		return nil, fmt.Errorf("error expanding StructType into structType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["structType"] = v
	}

	return m, nil
}

// flattenRoutineArgumentsDataType flattens an instance of RoutineArgumentsDataType from a JSON
// response object.
func flattenRoutineArgumentsDataType(c *Client, i interface{}, res *Routine) *RoutineArgumentsDataType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RoutineArgumentsDataType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyRoutineArgumentsDataType
	}
	r.TypeKind = flattenRoutineArgumentsDataTypeTypeKindEnum(m["typeKind"])
	r.ArrayElementType = flattenRoutineArgumentsDataType(c, m["arrayElementType"], res)
	r.StructType = flattenRoutineArgumentsDataTypeStructType(c, m["structType"], res)

	return r
}

// expandRoutineArgumentsDataTypeStructTypeMap expands the contents of RoutineArgumentsDataTypeStructType into a JSON
// request object.
func expandRoutineArgumentsDataTypeStructTypeMap(c *Client, f map[string]RoutineArgumentsDataTypeStructType, res *Routine) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRoutineArgumentsDataTypeStructType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRoutineArgumentsDataTypeStructTypeSlice expands the contents of RoutineArgumentsDataTypeStructType into a JSON
// request object.
func expandRoutineArgumentsDataTypeStructTypeSlice(c *Client, f []RoutineArgumentsDataTypeStructType, res *Routine) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRoutineArgumentsDataTypeStructType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRoutineArgumentsDataTypeStructTypeMap flattens the contents of RoutineArgumentsDataTypeStructType from a JSON
// response object.
func flattenRoutineArgumentsDataTypeStructTypeMap(c *Client, i interface{}, res *Routine) map[string]RoutineArgumentsDataTypeStructType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineArgumentsDataTypeStructType{}
	}

	if len(a) == 0 {
		return map[string]RoutineArgumentsDataTypeStructType{}
	}

	items := make(map[string]RoutineArgumentsDataTypeStructType)
	for k, item := range a {
		items[k] = *flattenRoutineArgumentsDataTypeStructType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenRoutineArgumentsDataTypeStructTypeSlice flattens the contents of RoutineArgumentsDataTypeStructType from a JSON
// response object.
func flattenRoutineArgumentsDataTypeStructTypeSlice(c *Client, i interface{}, res *Routine) []RoutineArgumentsDataTypeStructType {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineArgumentsDataTypeStructType{}
	}

	if len(a) == 0 {
		return []RoutineArgumentsDataTypeStructType{}
	}

	items := make([]RoutineArgumentsDataTypeStructType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineArgumentsDataTypeStructType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandRoutineArgumentsDataTypeStructType expands an instance of RoutineArgumentsDataTypeStructType into a JSON
// request object.
func expandRoutineArgumentsDataTypeStructType(c *Client, f *RoutineArgumentsDataTypeStructType, res *Routine) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandRoutineArgumentsDataTypeStructTypeFieldsSlice(c, f.Fields, res); err != nil {
		return nil, fmt.Errorf("error expanding Fields into fields: %w", err)
	} else if v != nil {
		m["fields"] = v
	}

	return m, nil
}

// flattenRoutineArgumentsDataTypeStructType flattens an instance of RoutineArgumentsDataTypeStructType from a JSON
// response object.
func flattenRoutineArgumentsDataTypeStructType(c *Client, i interface{}, res *Routine) *RoutineArgumentsDataTypeStructType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RoutineArgumentsDataTypeStructType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyRoutineArgumentsDataTypeStructType
	}
	r.Fields = flattenRoutineArgumentsDataTypeStructTypeFieldsSlice(c, m["fields"], res)

	return r
}

// expandRoutineArgumentsDataTypeStructTypeFieldsMap expands the contents of RoutineArgumentsDataTypeStructTypeFields into a JSON
// request object.
func expandRoutineArgumentsDataTypeStructTypeFieldsMap(c *Client, f map[string]RoutineArgumentsDataTypeStructTypeFields, res *Routine) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRoutineArgumentsDataTypeStructTypeFields(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRoutineArgumentsDataTypeStructTypeFieldsSlice expands the contents of RoutineArgumentsDataTypeStructTypeFields into a JSON
// request object.
func expandRoutineArgumentsDataTypeStructTypeFieldsSlice(c *Client, f []RoutineArgumentsDataTypeStructTypeFields, res *Routine) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRoutineArgumentsDataTypeStructTypeFields(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRoutineArgumentsDataTypeStructTypeFieldsMap flattens the contents of RoutineArgumentsDataTypeStructTypeFields from a JSON
// response object.
func flattenRoutineArgumentsDataTypeStructTypeFieldsMap(c *Client, i interface{}, res *Routine) map[string]RoutineArgumentsDataTypeStructTypeFields {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineArgumentsDataTypeStructTypeFields{}
	}

	if len(a) == 0 {
		return map[string]RoutineArgumentsDataTypeStructTypeFields{}
	}

	items := make(map[string]RoutineArgumentsDataTypeStructTypeFields)
	for k, item := range a {
		items[k] = *flattenRoutineArgumentsDataTypeStructTypeFields(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenRoutineArgumentsDataTypeStructTypeFieldsSlice flattens the contents of RoutineArgumentsDataTypeStructTypeFields from a JSON
// response object.
func flattenRoutineArgumentsDataTypeStructTypeFieldsSlice(c *Client, i interface{}, res *Routine) []RoutineArgumentsDataTypeStructTypeFields {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineArgumentsDataTypeStructTypeFields{}
	}

	if len(a) == 0 {
		return []RoutineArgumentsDataTypeStructTypeFields{}
	}

	items := make([]RoutineArgumentsDataTypeStructTypeFields, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineArgumentsDataTypeStructTypeFields(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandRoutineArgumentsDataTypeStructTypeFields expands an instance of RoutineArgumentsDataTypeStructTypeFields into a JSON
// request object.
func expandRoutineArgumentsDataTypeStructTypeFields(c *Client, f *RoutineArgumentsDataTypeStructTypeFields, res *Routine) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandRoutineArgumentsDataType(c, f.Type, res); err != nil {
		return nil, fmt.Errorf("error expanding Type into type: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenRoutineArgumentsDataTypeStructTypeFields flattens an instance of RoutineArgumentsDataTypeStructTypeFields from a JSON
// response object.
func flattenRoutineArgumentsDataTypeStructTypeFields(c *Client, i interface{}, res *Routine) *RoutineArgumentsDataTypeStructTypeFields {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RoutineArgumentsDataTypeStructTypeFields{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyRoutineArgumentsDataTypeStructTypeFields
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Type = flattenRoutineArgumentsDataType(c, m["type"], res)

	return r
}

// flattenRoutineRoutineTypeEnumMap flattens the contents of RoutineRoutineTypeEnum from a JSON
// response object.
func flattenRoutineRoutineTypeEnumMap(c *Client, i interface{}, res *Routine) map[string]RoutineRoutineTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineRoutineTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]RoutineRoutineTypeEnum{}
	}

	items := make(map[string]RoutineRoutineTypeEnum)
	for k, item := range a {
		items[k] = *flattenRoutineRoutineTypeEnum(item.(interface{}))
	}

	return items
}

// flattenRoutineRoutineTypeEnumSlice flattens the contents of RoutineRoutineTypeEnum from a JSON
// response object.
func flattenRoutineRoutineTypeEnumSlice(c *Client, i interface{}, res *Routine) []RoutineRoutineTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineRoutineTypeEnum{}
	}

	if len(a) == 0 {
		return []RoutineRoutineTypeEnum{}
	}

	items := make([]RoutineRoutineTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineRoutineTypeEnum(item.(interface{})))
	}

	return items
}

// flattenRoutineRoutineTypeEnum asserts that an interface is a string, and returns a
// pointer to a *RoutineRoutineTypeEnum with the same value as that string.
func flattenRoutineRoutineTypeEnum(i interface{}) *RoutineRoutineTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return RoutineRoutineTypeEnumRef(s)
}

// flattenRoutineLanguageEnumMap flattens the contents of RoutineLanguageEnum from a JSON
// response object.
func flattenRoutineLanguageEnumMap(c *Client, i interface{}, res *Routine) map[string]RoutineLanguageEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineLanguageEnum{}
	}

	if len(a) == 0 {
		return map[string]RoutineLanguageEnum{}
	}

	items := make(map[string]RoutineLanguageEnum)
	for k, item := range a {
		items[k] = *flattenRoutineLanguageEnum(item.(interface{}))
	}

	return items
}

// flattenRoutineLanguageEnumSlice flattens the contents of RoutineLanguageEnum from a JSON
// response object.
func flattenRoutineLanguageEnumSlice(c *Client, i interface{}, res *Routine) []RoutineLanguageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineLanguageEnum{}
	}

	if len(a) == 0 {
		return []RoutineLanguageEnum{}
	}

	items := make([]RoutineLanguageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineLanguageEnum(item.(interface{})))
	}

	return items
}

// flattenRoutineLanguageEnum asserts that an interface is a string, and returns a
// pointer to a *RoutineLanguageEnum with the same value as that string.
func flattenRoutineLanguageEnum(i interface{}) *RoutineLanguageEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return RoutineLanguageEnumRef(s)
}

// flattenRoutineArgumentsArgumentKindEnumMap flattens the contents of RoutineArgumentsArgumentKindEnum from a JSON
// response object.
func flattenRoutineArgumentsArgumentKindEnumMap(c *Client, i interface{}, res *Routine) map[string]RoutineArgumentsArgumentKindEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineArgumentsArgumentKindEnum{}
	}

	if len(a) == 0 {
		return map[string]RoutineArgumentsArgumentKindEnum{}
	}

	items := make(map[string]RoutineArgumentsArgumentKindEnum)
	for k, item := range a {
		items[k] = *flattenRoutineArgumentsArgumentKindEnum(item.(interface{}))
	}

	return items
}

// flattenRoutineArgumentsArgumentKindEnumSlice flattens the contents of RoutineArgumentsArgumentKindEnum from a JSON
// response object.
func flattenRoutineArgumentsArgumentKindEnumSlice(c *Client, i interface{}, res *Routine) []RoutineArgumentsArgumentKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineArgumentsArgumentKindEnum{}
	}

	if len(a) == 0 {
		return []RoutineArgumentsArgumentKindEnum{}
	}

	items := make([]RoutineArgumentsArgumentKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineArgumentsArgumentKindEnum(item.(interface{})))
	}

	return items
}

// flattenRoutineArgumentsArgumentKindEnum asserts that an interface is a string, and returns a
// pointer to a *RoutineArgumentsArgumentKindEnum with the same value as that string.
func flattenRoutineArgumentsArgumentKindEnum(i interface{}) *RoutineArgumentsArgumentKindEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return RoutineArgumentsArgumentKindEnumRef(s)
}

// flattenRoutineArgumentsModeEnumMap flattens the contents of RoutineArgumentsModeEnum from a JSON
// response object.
func flattenRoutineArgumentsModeEnumMap(c *Client, i interface{}, res *Routine) map[string]RoutineArgumentsModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineArgumentsModeEnum{}
	}

	if len(a) == 0 {
		return map[string]RoutineArgumentsModeEnum{}
	}

	items := make(map[string]RoutineArgumentsModeEnum)
	for k, item := range a {
		items[k] = *flattenRoutineArgumentsModeEnum(item.(interface{}))
	}

	return items
}

// flattenRoutineArgumentsModeEnumSlice flattens the contents of RoutineArgumentsModeEnum from a JSON
// response object.
func flattenRoutineArgumentsModeEnumSlice(c *Client, i interface{}, res *Routine) []RoutineArgumentsModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineArgumentsModeEnum{}
	}

	if len(a) == 0 {
		return []RoutineArgumentsModeEnum{}
	}

	items := make([]RoutineArgumentsModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineArgumentsModeEnum(item.(interface{})))
	}

	return items
}

// flattenRoutineArgumentsModeEnum asserts that an interface is a string, and returns a
// pointer to a *RoutineArgumentsModeEnum with the same value as that string.
func flattenRoutineArgumentsModeEnum(i interface{}) *RoutineArgumentsModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return RoutineArgumentsModeEnumRef(s)
}

// flattenRoutineArgumentsDataTypeTypeKindEnumMap flattens the contents of RoutineArgumentsDataTypeTypeKindEnum from a JSON
// response object.
func flattenRoutineArgumentsDataTypeTypeKindEnumMap(c *Client, i interface{}, res *Routine) map[string]RoutineArgumentsDataTypeTypeKindEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineArgumentsDataTypeTypeKindEnum{}
	}

	if len(a) == 0 {
		return map[string]RoutineArgumentsDataTypeTypeKindEnum{}
	}

	items := make(map[string]RoutineArgumentsDataTypeTypeKindEnum)
	for k, item := range a {
		items[k] = *flattenRoutineArgumentsDataTypeTypeKindEnum(item.(interface{}))
	}

	return items
}

// flattenRoutineArgumentsDataTypeTypeKindEnumSlice flattens the contents of RoutineArgumentsDataTypeTypeKindEnum from a JSON
// response object.
func flattenRoutineArgumentsDataTypeTypeKindEnumSlice(c *Client, i interface{}, res *Routine) []RoutineArgumentsDataTypeTypeKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineArgumentsDataTypeTypeKindEnum{}
	}

	if len(a) == 0 {
		return []RoutineArgumentsDataTypeTypeKindEnum{}
	}

	items := make([]RoutineArgumentsDataTypeTypeKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineArgumentsDataTypeTypeKindEnum(item.(interface{})))
	}

	return items
}

// flattenRoutineArgumentsDataTypeTypeKindEnum asserts that an interface is a string, and returns a
// pointer to a *RoutineArgumentsDataTypeTypeKindEnum with the same value as that string.
func flattenRoutineArgumentsDataTypeTypeKindEnum(i interface{}) *RoutineArgumentsDataTypeTypeKindEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return RoutineArgumentsDataTypeTypeKindEnumRef(s)
}

// flattenRoutineDeterminismLevelEnumMap flattens the contents of RoutineDeterminismLevelEnum from a JSON
// response object.
func flattenRoutineDeterminismLevelEnumMap(c *Client, i interface{}, res *Routine) map[string]RoutineDeterminismLevelEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoutineDeterminismLevelEnum{}
	}

	if len(a) == 0 {
		return map[string]RoutineDeterminismLevelEnum{}
	}

	items := make(map[string]RoutineDeterminismLevelEnum)
	for k, item := range a {
		items[k] = *flattenRoutineDeterminismLevelEnum(item.(interface{}))
	}

	return items
}

// flattenRoutineDeterminismLevelEnumSlice flattens the contents of RoutineDeterminismLevelEnum from a JSON
// response object.
func flattenRoutineDeterminismLevelEnumSlice(c *Client, i interface{}, res *Routine) []RoutineDeterminismLevelEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RoutineDeterminismLevelEnum{}
	}

	if len(a) == 0 {
		return []RoutineDeterminismLevelEnum{}
	}

	items := make([]RoutineDeterminismLevelEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoutineDeterminismLevelEnum(item.(interface{})))
	}

	return items
}

// flattenRoutineDeterminismLevelEnum asserts that an interface is a string, and returns a
// pointer to a *RoutineDeterminismLevelEnum with the same value as that string.
func flattenRoutineDeterminismLevelEnum(i interface{}) *RoutineDeterminismLevelEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return RoutineDeterminismLevelEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Routine) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalRoutine(b, c, r)
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
		if nr.Dataset == nil && ncr.Dataset == nil {
			c.Config.Logger.Info("Both Dataset fields null - considering equal.")
		} else if nr.Dataset == nil || ncr.Dataset == nil {
			c.Config.Logger.Info("Only one Dataset field is null - considering unequal.")
			return false
		} else if *nr.Dataset != *ncr.Dataset {
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

type routineDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         routineApiOperation
}

func convertFieldDiffsToRoutineDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]routineDiff, error) {
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
	var diffs []routineDiff
	// For each operation name, create a routineDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := routineDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToRoutineApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToRoutineApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (routineApiOperation, error) {
	switch opName {

	case "updateRoutinePatchRoutineOperation":
		return &updateRoutinePatchRoutineOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractRoutineFields(r *Routine) error {
	// *RoutineArgumentsDataType is a reused type - that's not compatible with function extractors.

	return nil
}
func extractRoutineArgumentsFields(r *Routine, o *RoutineArguments) error {
	// *RoutineArgumentsDataType is a reused type - that's not compatible with function extractors.

	return nil
}
func extractRoutineArgumentsDataTypeFields(r *Routine, o *RoutineArgumentsDataType) error {
	// *RoutineArgumentsDataType is a reused type - that's not compatible with function extractors.

	vStructType := o.StructType
	if vStructType == nil {
		// note: explicitly not the empty object.
		vStructType = &RoutineArgumentsDataTypeStructType{}
	}
	if err := extractRoutineArgumentsDataTypeStructTypeFields(r, vStructType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStructType) {
		o.StructType = vStructType
	}
	return nil
}
func extractRoutineArgumentsDataTypeStructTypeFields(r *Routine, o *RoutineArgumentsDataTypeStructType) error {
	return nil
}
func extractRoutineArgumentsDataTypeStructTypeFieldsFields(r *Routine, o *RoutineArgumentsDataTypeStructTypeFields) error {
	// *RoutineArgumentsDataType is a reused type - that's not compatible with function extractors.

	return nil
}

func postReadExtractRoutineFields(r *Routine) error {
	// *RoutineArgumentsDataType is a reused type - that's not compatible with function extractors.

	return nil
}
func postReadExtractRoutineArgumentsFields(r *Routine, o *RoutineArguments) error {
	// *RoutineArgumentsDataType is a reused type - that's not compatible with function extractors.

	return nil
}
func postReadExtractRoutineArgumentsDataTypeFields(r *Routine, o *RoutineArgumentsDataType) error {
	// *RoutineArgumentsDataType is a reused type - that's not compatible with function extractors.

	vStructType := o.StructType
	if vStructType == nil {
		// note: explicitly not the empty object.
		vStructType = &RoutineArgumentsDataTypeStructType{}
	}
	if err := extractRoutineArgumentsDataTypeStructTypeFields(r, vStructType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStructType) {
		o.StructType = vStructType
	}
	return nil
}
func postReadExtractRoutineArgumentsDataTypeStructTypeFields(r *Routine, o *RoutineArgumentsDataTypeStructType) error {
	return nil
}
func postReadExtractRoutineArgumentsDataTypeStructTypeFieldsFields(r *Routine, o *RoutineArgumentsDataTypeStructTypeFields) error {
	// *RoutineArgumentsDataType is a reused type - that's not compatible with function extractors.

	return nil
}
