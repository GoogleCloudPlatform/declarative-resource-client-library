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

func (r *MetadataSchema) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "schemaVersion"); err != nil {
		return err
	}
	if err := dcl.Required(r, "schema"); err != nil {
		return err
	}
	if err := dcl.Required(r, "schemaType"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.MetadataStore, "MetadataStore"); err != nil {
		return err
	}
	return nil
}
func (r *MetadataSchema) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-aiplatform.googleapis.com/v1beta1/", params)
}

func (r *MetadataSchema) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":       dcl.ValueOrEmptyString(nr.Project),
		"location":      dcl.ValueOrEmptyString(nr.Location),
		"metadataStore": dcl.ValueOrEmptyString(nr.MetadataStore),
		"name":          dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/metadataStores/{{metadataStore}}/metadataSchemas/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *MetadataSchema) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":       dcl.ValueOrEmptyString(nr.Project),
		"location":      dcl.ValueOrEmptyString(nr.Location),
		"metadataStore": dcl.ValueOrEmptyString(nr.MetadataStore),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/metadataStores/{{metadataStore}}/metadataSchemas", nr.basePath(), userBasePath, params), nil

}

func (r *MetadataSchema) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":       dcl.ValueOrEmptyString(nr.Project),
		"location":      dcl.ValueOrEmptyString(nr.Location),
		"metadataStore": dcl.ValueOrEmptyString(nr.MetadataStore),
		"name":          dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/metadataStores/{{metadataStore}}/metadataSchemas?metadataSchemaId={{name}}", nr.basePath(), userBasePath, params), nil

}

// metadataSchemaApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type metadataSchemaApiOperation interface {
	do(context.Context, *MetadataSchema, *Client) error
}

func (c *Client) listMetadataSchemaRaw(ctx context.Context, r *MetadataSchema, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != MetadataSchemaMaxPage {
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

type listMetadataSchemaOperation struct {
	MetadataSchemas []map[string]interface{} `json:"metadataSchemas"`
	Token           string                   `json:"nextPageToken"`
}

func (c *Client) listMetadataSchema(ctx context.Context, r *MetadataSchema, pageToken string, pageSize int32) ([]*MetadataSchema, string, error) {
	b, err := c.listMetadataSchemaRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listMetadataSchemaOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*MetadataSchema
	for _, v := range m.MetadataSchemas {
		res, err := unmarshalMapMetadataSchema(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		res.MetadataStore = r.MetadataStore
		l = append(l, res)
	}

	return l, m.Token, nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createMetadataSchemaOperation struct {
	response map[string]interface{}
}

func (op *createMetadataSchemaOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createMetadataSchemaOperation) do(ctx context.Context, r *MetadataSchema, c *Client) error {
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

	if _, err := c.GetMetadataSchema(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getMetadataSchemaRaw(ctx context.Context, r *MetadataSchema) ([]byte, error) {

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

func (c *Client) metadataSchemaDiffsForRawDesired(ctx context.Context, rawDesired *MetadataSchema, opts ...dcl.ApplyOption) (initial, desired *MetadataSchema, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *MetadataSchema
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*MetadataSchema); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected MetadataSchema, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetMetadataSchema(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a MetadataSchema resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve MetadataSchema resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that MetadataSchema resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeMetadataSchemaDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for MetadataSchema: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for MetadataSchema: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractMetadataSchemaFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeMetadataSchemaInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for MetadataSchema: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeMetadataSchemaDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for MetadataSchema: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffMetadataSchema(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeMetadataSchemaInitialState(rawInitial, rawDesired *MetadataSchema) (*MetadataSchema, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeMetadataSchemaDesiredState(rawDesired, rawInitial *MetadataSchema, opts ...dcl.ApplyOption) (*MetadataSchema, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &MetadataSchema{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.SchemaVersion, rawInitial.SchemaVersion) {
		canonicalDesired.SchemaVersion = rawInitial.SchemaVersion
	} else {
		canonicalDesired.SchemaVersion = rawDesired.SchemaVersion
	}
	if dcl.StringCanonicalize(rawDesired.Schema, rawInitial.Schema) {
		canonicalDesired.Schema = rawInitial.Schema
	} else {
		canonicalDesired.Schema = rawDesired.Schema
	}
	if dcl.IsZeroValue(rawDesired.SchemaType) || (dcl.IsEmptyValueIndirect(rawDesired.SchemaType) && dcl.IsEmptyValueIndirect(rawInitial.SchemaType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.SchemaType = rawInitial.SchemaType
	} else {
		canonicalDesired.SchemaType = rawDesired.SchemaType
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
	if dcl.NameToSelfLink(rawDesired.MetadataStore, rawInitial.MetadataStore) {
		canonicalDesired.MetadataStore = rawInitial.MetadataStore
	} else {
		canonicalDesired.MetadataStore = rawDesired.MetadataStore
	}

	return canonicalDesired, nil
}

func canonicalizeMetadataSchemaNewState(c *Client, rawNew, rawDesired *MetadataSchema) (*MetadataSchema, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SchemaVersion) && dcl.IsEmptyValueIndirect(rawDesired.SchemaVersion) {
		rawNew.SchemaVersion = rawDesired.SchemaVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.SchemaVersion, rawNew.SchemaVersion) {
			rawNew.SchemaVersion = rawDesired.SchemaVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Schema) && dcl.IsEmptyValueIndirect(rawDesired.Schema) {
		rawNew.Schema = rawDesired.Schema
	} else {
		if dcl.StringCanonicalize(rawDesired.Schema, rawNew.Schema) {
			rawNew.Schema = rawDesired.Schema
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SchemaType) && dcl.IsEmptyValueIndirect(rawDesired.SchemaType) {
		rawNew.SchemaType = rawDesired.SchemaType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.MetadataStore = rawDesired.MetadataStore

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffMetadataSchema(c *Client, desired, actual *MetadataSchema, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SchemaVersion, actual.SchemaVersion, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SchemaVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Schema, actual.Schema, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Schema")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SchemaType, actual.SchemaType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SchemaType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.MetadataStore, actual.MetadataStore, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetadataStore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *MetadataSchema) urlNormalized() *MetadataSchema {
	normalized := dcl.Copy(*r).(MetadataSchema)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.SchemaVersion = dcl.SelfLinkToName(r.SchemaVersion)
	normalized.Schema = dcl.SelfLinkToName(r.Schema)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.MetadataStore = dcl.SelfLinkToName(r.MetadataStore)
	return &normalized
}

func (r *MetadataSchema) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the MetadataSchema resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *MetadataSchema) marshal(c *Client) ([]byte, error) {
	m, err := expandMetadataSchema(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling MetadataSchema: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalMetadataSchema decodes JSON responses into the MetadataSchema resource schema.
func unmarshalMetadataSchema(b []byte, c *Client, res *MetadataSchema) (*MetadataSchema, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapMetadataSchema(m, c, res)
}

func unmarshalMapMetadataSchema(m map[string]interface{}, c *Client, res *MetadataSchema) (*MetadataSchema, error) {

	flattened := flattenMetadataSchema(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandMetadataSchema expands MetadataSchema into a JSON request object.
func expandMetadataSchema(c *Client, f *MetadataSchema) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/metadataStores/%s/metadataSchemas/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.MetadataStore), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.SchemaVersion; dcl.ValueShouldBeSent(v) {
		m["schemaVersion"] = v
	}
	if v := f.Schema; dcl.ValueShouldBeSent(v) {
		m["schema"] = v
	}
	if v := f.SchemaType; dcl.ValueShouldBeSent(v) {
		m["schemaType"] = v
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding MetadataStore into metadataStore: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadataStore"] = v
	}

	return m, nil
}

// flattenMetadataSchema flattens MetadataSchema from a JSON request object into the
// MetadataSchema type.
func flattenMetadataSchema(c *Client, i interface{}, res *MetadataSchema) *MetadataSchema {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &MetadataSchema{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.SchemaVersion = dcl.FlattenString(m["schemaVersion"])
	resultRes.Schema = dcl.FlattenString(m["schema"])
	resultRes.SchemaType = flattenMetadataSchemaSchemaTypeEnum(m["schemaType"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.MetadataStore = dcl.FlattenString(m["metadataStore"])

	return resultRes
}

// flattenMetadataSchemaSchemaTypeEnumMap flattens the contents of MetadataSchemaSchemaTypeEnum from a JSON
// response object.
func flattenMetadataSchemaSchemaTypeEnumMap(c *Client, i interface{}, res *MetadataSchema) map[string]MetadataSchemaSchemaTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MetadataSchemaSchemaTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]MetadataSchemaSchemaTypeEnum{}
	}

	items := make(map[string]MetadataSchemaSchemaTypeEnum)
	for k, item := range a {
		items[k] = *flattenMetadataSchemaSchemaTypeEnum(item.(interface{}))
	}

	return items
}

// flattenMetadataSchemaSchemaTypeEnumSlice flattens the contents of MetadataSchemaSchemaTypeEnum from a JSON
// response object.
func flattenMetadataSchemaSchemaTypeEnumSlice(c *Client, i interface{}, res *MetadataSchema) []MetadataSchemaSchemaTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MetadataSchemaSchemaTypeEnum{}
	}

	if len(a) == 0 {
		return []MetadataSchemaSchemaTypeEnum{}
	}

	items := make([]MetadataSchemaSchemaTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMetadataSchemaSchemaTypeEnum(item.(interface{})))
	}

	return items
}

// flattenMetadataSchemaSchemaTypeEnum asserts that an interface is a string, and returns a
// pointer to a *MetadataSchemaSchemaTypeEnum with the same value as that string.
func flattenMetadataSchemaSchemaTypeEnum(i interface{}) *MetadataSchemaSchemaTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return MetadataSchemaSchemaTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *MetadataSchema) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalMetadataSchema(b, c, r)
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
		if nr.MetadataStore == nil && ncr.MetadataStore == nil {
			c.Config.Logger.Info("Both MetadataStore fields null - considering equal.")
		} else if nr.MetadataStore == nil || ncr.MetadataStore == nil {
			c.Config.Logger.Info("Only one MetadataStore field is null - considering unequal.")
			return false
		} else if *nr.MetadataStore != *ncr.MetadataStore {
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

type metadataSchemaDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         metadataSchemaApiOperation
}

func convertFieldDiffsToMetadataSchemaDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]metadataSchemaDiff, error) {
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
	var diffs []metadataSchemaDiff
	// For each operation name, create a metadataSchemaDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := metadataSchemaDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToMetadataSchemaApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToMetadataSchemaApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (metadataSchemaApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractMetadataSchemaFields(r *MetadataSchema) error {
	return nil
}

func postReadExtractMetadataSchemaFields(r *MetadataSchema) error {
	return nil
}
