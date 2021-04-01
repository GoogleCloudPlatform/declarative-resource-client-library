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
package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Index) validate() error {

	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "properties"); err != nil {
		return err
	}
	return nil
}
func (r *IndexProperties) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "direction"); err != nil {
		return err
	}
	return nil
}

func indexGetURL(userBasePath string, r *Index) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"indexId": dcl.ValueOrEmptyString(r.IndexId),
	}
	return dcl.URL("projects/{{project}}/indexes/{{indexId}}", "https://datastore.googleapis.com/v1/", userBasePath, params), nil
}

func indexListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/indexes", "https://datastore.googleapis.com/v1/", userBasePath, params), nil

}

func indexCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/indexes", "https://datastore.googleapis.com/v1/", userBasePath, params), nil

}

func indexDeleteURL(userBasePath string, r *Index) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"indexId": dcl.ValueOrEmptyString(r.IndexId),
	}
	return dcl.URL("projects/{{project}}/indexes/{{indexId}}", "https://datastore.googleapis.com/v1/", userBasePath, params), nil
}

// indexApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type indexApiOperation interface {
	do(context.Context, *Index, *Client) error
}

func (c *Client) listIndexRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := indexListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != IndexMaxPage {
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

type listIndexOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listIndex(ctx context.Context, project, pageToken string, pageSize int32) ([]*Index, string, error) {
	b, err := c.listIndexRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listIndexOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Index
	for _, v := range m.Items {
		res := flattenIndex(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllIndex(ctx context.Context, f func(*Index) bool, resources []*Index) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteIndex(ctx, res)
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

type deleteIndexOperation struct{}

func (op *deleteIndexOperation) do(ctx context.Context, r *Index, c *Client) error {

	_, err := c.GetIndex(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Index not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetIndex checking for existence. error: %v", err)
		return err
	}

	u, err := indexDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.DatastoreOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://datastore.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetIndex(ctx, r.urlNormalized())
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
type createIndexOperation struct {
	response map[string]interface{}
}

func (op *createIndexOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createIndexOperation) do(ctx context.Context, r *Index, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := indexCreateURL(c.Config.BasePath, project)

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
	var o operations.DatastoreOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://datastore.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include IndexId in URL substitution for initial GET request.
	indexId, ok := op.response["indexId"].(string)
	if !ok {
		return fmt.Errorf("expected indexId to be a string")
	}
	r.IndexId = &indexId

	if _, err := c.GetIndex(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getIndexRaw(ctx context.Context, r *Index) ([]byte, error) {
	if dcl.IsZeroValue(r.Ancestor) {
		r.Ancestor = IndexAncestorEnumRef("NONE")
	}

	u, err := indexGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) indexDiffsForRawDesired(ctx context.Context, rawDesired *Index, opts ...dcl.ApplyOption) (initial, desired *Index, diffs []indexDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Index
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Index); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Index, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.IndexId == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeIndexDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetIndex(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Index resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Index resource: %v", err)
		}
		c.Config.Logger.Info("Found that Index resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeIndexDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Index: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Index: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeIndexInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Index: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeIndexDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Index: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffIndex(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeIndexInitialState(rawInitial, rawDesired *Index) (*Index, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeIndexDesiredState(rawDesired, rawInitial *Index, opts ...dcl.ApplyOption) (*Index, error) {

	if dcl.IsZeroValue(rawDesired.Ancestor) {
		rawDesired.Ancestor = IndexAncestorEnumRef("NONE")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Ancestor) {
		rawDesired.Ancestor = rawInitial.Ancestor
	}
	if dcl.IsZeroValue(rawDesired.IndexId) {
		rawDesired.IndexId = rawInitial.IndexId
	}
	if dcl.StringCanonicalize(rawDesired.Kind, rawInitial.Kind) {
		rawDesired.Kind = rawInitial.Kind
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.IsZeroValue(rawDesired.Properties) {
		rawDesired.Properties = rawInitial.Properties
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}

	return rawDesired, nil
}

func canonicalizeIndexNewState(c *Client, rawNew, rawDesired *Index) (*Index, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Ancestor) && dcl.IsEmptyValueIndirect(rawDesired.Ancestor) {
		rawNew.Ancestor = rawDesired.Ancestor
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IndexId) && dcl.IsEmptyValueIndirect(rawDesired.IndexId) {
		rawNew.IndexId = rawDesired.IndexId
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Kind) && dcl.IsEmptyValueIndirect(rawDesired.Kind) {
		rawNew.Kind = rawDesired.Kind
	} else {
		if dcl.StringCanonicalize(rawDesired.Kind, rawNew.Kind) {
			rawNew.Kind = rawDesired.Kind
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Properties) && dcl.IsEmptyValueIndirect(rawDesired.Properties) {
		rawNew.Properties = rawDesired.Properties
	} else {
		rawNew.Properties = canonicalizeNewIndexPropertiesSlice(c, rawDesired.Properties, rawNew.Properties)
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	return rawNew, nil
}

func canonicalizeIndexProperties(des, initial *IndexProperties, opts ...dcl.ApplyOption) *IndexProperties {
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
	if dcl.IsZeroValue(des.Direction) {
		des.Direction = initial.Direction
	}

	return des
}

func canonicalizeNewIndexProperties(c *Client, des, nw *IndexProperties) *IndexProperties {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewIndexPropertiesSet(c *Client, des, nw []IndexProperties) []IndexProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []IndexProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareIndexProperties(c, &d, &n) {
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

func canonicalizeNewIndexPropertiesSlice(c *Client, des, nw []IndexProperties) []IndexProperties {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []IndexProperties
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewIndexProperties(c, &d, &n))
	}

	return items
}

type indexDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         indexApiOperation
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
func diffIndex(c *Client, desired, actual *Index, opts ...dcl.ApplyOption) ([]indexDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []indexDiff
	if !reflect.DeepEqual(desired.Ancestor, actual.Ancestor) {
		c.Config.Logger.Infof("Detected diff in Ancestor.\nDESIRED: %v\nACTUAL: %v", desired.Ancestor, actual.Ancestor)
		diffs = append(diffs, indexDiff{
			RequiresRecreate: true,
			FieldName:        "Ancestor",
		})
	}
	if !dcl.IsZeroValue(desired.Kind) && !dcl.StringCanonicalize(desired.Kind, actual.Kind) {
		c.Config.Logger.Infof("Detected diff in Kind.\nDESIRED: %v\nACTUAL: %v", desired.Kind, actual.Kind)
		diffs = append(diffs, indexDiff{
			RequiresRecreate: true,
			FieldName:        "Kind",
		})
	}
	if compareIndexPropertiesSlice(c, desired.Properties, actual.Properties) {
		c.Config.Logger.Infof("Detected diff in Properties.\nDESIRED: %v\nACTUAL: %v", desired.Properties, actual.Properties)
		diffs = append(diffs, indexDiff{
			RequiresRecreate: true,
			FieldName:        "Properties",
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
	var deduped []indexDiff
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
func compareIndexProperties(c *Client, desired, actual *IndexProperties) bool {
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
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Direction == nil && desired.Direction != nil && !dcl.IsEmptyValueIndirect(desired.Direction) {
		c.Config.Logger.Infof("desired Direction %s - but actually nil", dcl.SprintResource(desired.Direction))
		return true
	}
	if !reflect.DeepEqual(desired.Direction, actual.Direction) && !dcl.IsZeroValue(desired.Direction) {
		c.Config.Logger.Infof("Diff in Direction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Direction), dcl.SprintResource(actual.Direction))
		return true
	}
	return false
}

func compareIndexPropertiesSlice(c *Client, desired, actual []IndexProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in IndexProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareIndexProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in IndexProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareIndexPropertiesMap(c *Client, desired, actual map[string]IndexProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in IndexProperties, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in IndexProperties, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareIndexProperties(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in IndexProperties, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareIndexAncestorEnumSlice(c *Client, desired, actual []IndexAncestorEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in IndexAncestorEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareIndexAncestorEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in IndexAncestorEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareIndexAncestorEnum(c *Client, desired, actual *IndexAncestorEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareIndexPropertiesDirectionEnumSlice(c *Client, desired, actual []IndexPropertiesDirectionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in IndexPropertiesDirectionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareIndexPropertiesDirectionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in IndexPropertiesDirectionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareIndexPropertiesDirectionEnum(c *Client, desired, actual *IndexPropertiesDirectionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareIndexStateEnumSlice(c *Client, desired, actual []IndexStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in IndexStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareIndexStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in IndexStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareIndexStateEnum(c *Client, desired, actual *IndexStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Index) urlNormalized() *Index {
	normalized := deepcopy.Copy(*r).(Index)
	normalized.IndexId = dcl.SelfLinkToName(r.IndexId)
	normalized.Kind = dcl.SelfLinkToName(r.Kind)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Index) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.IndexId)
}

func (r *Index) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Index) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.IndexId)
}

func (r *Index) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Index resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Index) marshal(c *Client) ([]byte, error) {
	m, err := expandIndex(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Index: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalIndex decodes JSON responses into the Index resource schema.
func unmarshalIndex(b []byte, c *Client) (*Index, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapIndex(m, c)
}

func unmarshalMapIndex(m map[string]interface{}, c *Client) (*Index, error) {

	return flattenIndex(c, m), nil
}

// expandIndex expands Index into a JSON request object.
func expandIndex(c *Client, f *Index) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Ancestor; !dcl.IsEmptyValueIndirect(v) {
		m["ancestor"] = v
	}
	if v := f.IndexId; !dcl.IsEmptyValueIndirect(v) {
		m["indexId"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := expandIndexPropertiesSlice(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}

	return m, nil
}

// flattenIndex flattens Index from a JSON request object into the
// Index type.
func flattenIndex(c *Client, i interface{}) *Index {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Index{}
	r.Ancestor = flattenIndexAncestorEnum(m["ancestor"])
	if _, ok := m["ancestor"]; !ok {
		c.Config.Logger.Info("Using default value for ancestor")
		r.Ancestor = IndexAncestorEnumRef("NONE")
	}
	r.IndexId = dcl.SelfLinkToName(dcl.FlattenString(m["indexId"]))
	r.Kind = dcl.FlattenString(m["kind"])
	r.Project = dcl.FlattenString(m["project"])
	r.Properties = flattenIndexPropertiesSlice(c, m["properties"])
	r.State = flattenIndexStateEnum(m["state"])

	return r
}

// expandIndexPropertiesMap expands the contents of IndexProperties into a JSON
// request object.
func expandIndexPropertiesMap(c *Client, f map[string]IndexProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandIndexProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandIndexPropertiesSlice expands the contents of IndexProperties into a JSON
// request object.
func expandIndexPropertiesSlice(c *Client, f []IndexProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandIndexProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenIndexPropertiesMap flattens the contents of IndexProperties from a JSON
// response object.
func flattenIndexPropertiesMap(c *Client, i interface{}) map[string]IndexProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]IndexProperties{}
	}

	if len(a) == 0 {
		return map[string]IndexProperties{}
	}

	items := make(map[string]IndexProperties)
	for k, item := range a {
		items[k] = *flattenIndexProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenIndexPropertiesSlice flattens the contents of IndexProperties from a JSON
// response object.
func flattenIndexPropertiesSlice(c *Client, i interface{}) []IndexProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []IndexProperties{}
	}

	if len(a) == 0 {
		return []IndexProperties{}
	}

	items := make([]IndexProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenIndexProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandIndexProperties expands an instance of IndexProperties into a JSON
// request object.
func expandIndexProperties(c *Client, f *IndexProperties) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		m["direction"] = v
	}

	return m, nil
}

// flattenIndexProperties flattens an instance of IndexProperties from a JSON
// response object.
func flattenIndexProperties(c *Client, i interface{}) *IndexProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &IndexProperties{}
	r.Name = dcl.FlattenString(m["name"])
	r.Direction = flattenIndexPropertiesDirectionEnum(m["direction"])

	return r
}

// flattenIndexAncestorEnumSlice flattens the contents of IndexAncestorEnum from a JSON
// response object.
func flattenIndexAncestorEnumSlice(c *Client, i interface{}) []IndexAncestorEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []IndexAncestorEnum{}
	}

	if len(a) == 0 {
		return []IndexAncestorEnum{}
	}

	items := make([]IndexAncestorEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenIndexAncestorEnum(item.(interface{})))
	}

	return items
}

// flattenIndexAncestorEnum asserts that an interface is a string, and returns a
// pointer to a *IndexAncestorEnum with the same value as that string.
func flattenIndexAncestorEnum(i interface{}) *IndexAncestorEnum {
	s, ok := i.(string)
	if !ok {
		return IndexAncestorEnumRef("")
	}

	return IndexAncestorEnumRef(s)
}

// flattenIndexPropertiesDirectionEnumSlice flattens the contents of IndexPropertiesDirectionEnum from a JSON
// response object.
func flattenIndexPropertiesDirectionEnumSlice(c *Client, i interface{}) []IndexPropertiesDirectionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []IndexPropertiesDirectionEnum{}
	}

	if len(a) == 0 {
		return []IndexPropertiesDirectionEnum{}
	}

	items := make([]IndexPropertiesDirectionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenIndexPropertiesDirectionEnum(item.(interface{})))
	}

	return items
}

// flattenIndexPropertiesDirectionEnum asserts that an interface is a string, and returns a
// pointer to a *IndexPropertiesDirectionEnum with the same value as that string.
func flattenIndexPropertiesDirectionEnum(i interface{}) *IndexPropertiesDirectionEnum {
	s, ok := i.(string)
	if !ok {
		return IndexPropertiesDirectionEnumRef("")
	}

	return IndexPropertiesDirectionEnumRef(s)
}

// flattenIndexStateEnumSlice flattens the contents of IndexStateEnum from a JSON
// response object.
func flattenIndexStateEnumSlice(c *Client, i interface{}) []IndexStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []IndexStateEnum{}
	}

	if len(a) == 0 {
		return []IndexStateEnum{}
	}

	items := make([]IndexStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenIndexStateEnum(item.(interface{})))
	}

	return items
}

// flattenIndexStateEnum asserts that an interface is a string, and returns a
// pointer to a *IndexStateEnum with the same value as that string.
func flattenIndexStateEnum(i interface{}) *IndexStateEnum {
	s, ok := i.(string)
	if !ok {
		return IndexStateEnumRef("")
	}

	return IndexStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Index) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalIndex(b, c)
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
		if nr.IndexId == nil && ncr.IndexId == nil {
			c.Config.Logger.Info("Both IndexId fields null - considering equal.")
		} else if nr.IndexId == nil || ncr.IndexId == nil {
			c.Config.Logger.Info("Only one IndexId field is null - considering unequal.")
			return false
		} else if *nr.IndexId != *ncr.IndexId {
			return false
		}
		return true
	}
}
