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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type FhirStore struct {
	Name                            *string                                       `json:"name"`
	EnableUpdateCreate              *bool                                         `json:"enableUpdateCreate"`
	NotificationConfig              *FhirStoreNotificationConfig                  `json:"notificationConfig"`
	DisableReferentialIntegrity     *bool                                         `json:"disableReferentialIntegrity"`
	ShardNum                        *int64                                        `json:"shardNum"`
	DisableResourceVersioning       *bool                                         `json:"disableResourceVersioning"`
	Labels                          map[string]string                             `json:"labels"`
	Version                         *FhirStoreVersionEnum                         `json:"version"`
	StreamConfigs                   []FhirStoreStreamConfigs                      `json:"streamConfigs"`
	ValidationConfig                *FhirStoreValidationConfig                    `json:"validationConfig"`
	DefaultSearchHandlingStrict     *bool                                         `json:"defaultSearchHandlingStrict"`
	ComplexDataTypeReferenceParsing *FhirStoreComplexDataTypeReferenceParsingEnum `json:"complexDataTypeReferenceParsing"`
	Project                         *string                                       `json:"project"`
	Location                        *string                                       `json:"location"`
	Dataset                         *string                                       `json:"dataset"`
}

func (r *FhirStore) String() string {
	return dcl.SprintResource(r)
}

// The enum FhirStoreVersionEnum.
type FhirStoreVersionEnum string

// FhirStoreVersionEnumRef returns a *FhirStoreVersionEnum with the value of string s
// If the empty string is provided, nil is returned.
func FhirStoreVersionEnumRef(s string) *FhirStoreVersionEnum {
	v := FhirStoreVersionEnum(s)
	return &v
}

func (v FhirStoreVersionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"VERSION_UNSPECIFIED", "DSTU2", "STU3", "R4"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FhirStoreVersionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum.
type FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum string

// FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumRef returns a *FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumRef(s string) *FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	v := FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(s)
	return &v
}

func (v FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SCHEMA_TYPE_UNSPECIFIED", "LOSSLESS", "ANALYTICS", "ANALYTICS_V2"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FhirStoreComplexDataTypeReferenceParsingEnum.
type FhirStoreComplexDataTypeReferenceParsingEnum string

// FhirStoreComplexDataTypeReferenceParsingEnumRef returns a *FhirStoreComplexDataTypeReferenceParsingEnum with the value of string s
// If the empty string is provided, nil is returned.
func FhirStoreComplexDataTypeReferenceParsingEnumRef(s string) *FhirStoreComplexDataTypeReferenceParsingEnum {
	v := FhirStoreComplexDataTypeReferenceParsingEnum(s)
	return &v
}

func (v FhirStoreComplexDataTypeReferenceParsingEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED", "DISABLED", "ENABLED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FhirStoreComplexDataTypeReferenceParsingEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type FhirStoreNotificationConfig struct {
	empty       bool    `json:"-"`
	PubsubTopic *string `json:"pubsubTopic"`
}

type jsonFhirStoreNotificationConfig FhirStoreNotificationConfig

func (r *FhirStoreNotificationConfig) UnmarshalJSON(data []byte) error {
	var res jsonFhirStoreNotificationConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFhirStoreNotificationConfig
	} else {

		r.PubsubTopic = res.PubsubTopic

	}
	return nil
}

// This object is used to assert a desired state where this FhirStoreNotificationConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFhirStoreNotificationConfig *FhirStoreNotificationConfig = &FhirStoreNotificationConfig{empty: true}

func (r *FhirStoreNotificationConfig) Empty() bool {
	return r.empty
}

func (r *FhirStoreNotificationConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *FhirStoreNotificationConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FhirStoreStreamConfigs struct {
	empty               bool                                       `json:"-"`
	ResourceTypes       []string                                   `json:"resourceTypes"`
	BigqueryDestination *FhirStoreStreamConfigsBigqueryDestination `json:"bigqueryDestination"`
}

type jsonFhirStoreStreamConfigs FhirStoreStreamConfigs

func (r *FhirStoreStreamConfigs) UnmarshalJSON(data []byte) error {
	var res jsonFhirStoreStreamConfigs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFhirStoreStreamConfigs
	} else {

		r.ResourceTypes = res.ResourceTypes

		r.BigqueryDestination = res.BigqueryDestination

	}
	return nil
}

// This object is used to assert a desired state where this FhirStoreStreamConfigs is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFhirStoreStreamConfigs *FhirStoreStreamConfigs = &FhirStoreStreamConfigs{empty: true}

func (r *FhirStoreStreamConfigs) Empty() bool {
	return r.empty
}

func (r *FhirStoreStreamConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *FhirStoreStreamConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FhirStoreStreamConfigsBigqueryDestination struct {
	empty        bool                                                   `json:"-"`
	DatasetUri   *string                                                `json:"datasetUri"`
	SchemaConfig *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig `json:"schemaConfig"`
}

type jsonFhirStoreStreamConfigsBigqueryDestination FhirStoreStreamConfigsBigqueryDestination

func (r *FhirStoreStreamConfigsBigqueryDestination) UnmarshalJSON(data []byte) error {
	var res jsonFhirStoreStreamConfigsBigqueryDestination
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFhirStoreStreamConfigsBigqueryDestination
	} else {

		r.DatasetUri = res.DatasetUri

		r.SchemaConfig = res.SchemaConfig

	}
	return nil
}

// This object is used to assert a desired state where this FhirStoreStreamConfigsBigqueryDestination is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFhirStoreStreamConfigsBigqueryDestination *FhirStoreStreamConfigsBigqueryDestination = &FhirStoreStreamConfigsBigqueryDestination{empty: true}

func (r *FhirStoreStreamConfigsBigqueryDestination) Empty() bool {
	return r.empty
}

func (r *FhirStoreStreamConfigsBigqueryDestination) String() string {
	return dcl.SprintResource(r)
}

func (r *FhirStoreStreamConfigsBigqueryDestination) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FhirStoreStreamConfigsBigqueryDestinationSchemaConfig struct {
	empty                   bool                                                                 `json:"-"`
	SchemaType              *FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum `json:"schemaType"`
	RecursiveStructureDepth *int64                                                               `json:"recursiveStructureDepth"`
}

type jsonFhirStoreStreamConfigsBigqueryDestinationSchemaConfig FhirStoreStreamConfigsBigqueryDestinationSchemaConfig

func (r *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) UnmarshalJSON(data []byte) error {
	var res jsonFhirStoreStreamConfigsBigqueryDestinationSchemaConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFhirStoreStreamConfigsBigqueryDestinationSchemaConfig
	} else {

		r.SchemaType = res.SchemaType

		r.RecursiveStructureDepth = res.RecursiveStructureDepth

	}
	return nil
}

// This object is used to assert a desired state where this FhirStoreStreamConfigsBigqueryDestinationSchemaConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFhirStoreStreamConfigsBigqueryDestinationSchemaConfig *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig = &FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{empty: true}

func (r *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) Empty() bool {
	return r.empty
}

func (r *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FhirStoreValidationConfig struct {
	empty                          bool     `json:"-"`
	DisableProfileValidation       *bool    `json:"disableProfileValidation"`
	EnabledImplementationGuides    []string `json:"enabledImplementationGuides"`
	DisableRequiredFieldValidation *bool    `json:"disableRequiredFieldValidation"`
	DisableReferenceTypeValidation *bool    `json:"disableReferenceTypeValidation"`
	DisableFhirpathValidation      *bool    `json:"disableFhirpathValidation"`
}

type jsonFhirStoreValidationConfig FhirStoreValidationConfig

func (r *FhirStoreValidationConfig) UnmarshalJSON(data []byte) error {
	var res jsonFhirStoreValidationConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFhirStoreValidationConfig
	} else {

		r.DisableProfileValidation = res.DisableProfileValidation

		r.EnabledImplementationGuides = res.EnabledImplementationGuides

		r.DisableRequiredFieldValidation = res.DisableRequiredFieldValidation

		r.DisableReferenceTypeValidation = res.DisableReferenceTypeValidation

		r.DisableFhirpathValidation = res.DisableFhirpathValidation

	}
	return nil
}

// This object is used to assert a desired state where this FhirStoreValidationConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFhirStoreValidationConfig *FhirStoreValidationConfig = &FhirStoreValidationConfig{empty: true}

func (r *FhirStoreValidationConfig) Empty() bool {
	return r.empty
}

func (r *FhirStoreValidationConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *FhirStoreValidationConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *FhirStore) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "healthcare",
		Type:    "FhirStore",
		Version: "beta",
	}
}

func (r *FhirStore) ID() (string, error) {
	if err := extractFhirStoreFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                                dcl.ValueOrEmptyString(nr.Name),
		"enable_update_create":                dcl.ValueOrEmptyString(nr.EnableUpdateCreate),
		"notification_config":                 dcl.ValueOrEmptyString(nr.NotificationConfig),
		"disable_referential_integrity":       dcl.ValueOrEmptyString(nr.DisableReferentialIntegrity),
		"shard_num":                           dcl.ValueOrEmptyString(nr.ShardNum),
		"disable_resource_versioning":         dcl.ValueOrEmptyString(nr.DisableResourceVersioning),
		"labels":                              dcl.ValueOrEmptyString(nr.Labels),
		"version":                             dcl.ValueOrEmptyString(nr.Version),
		"stream_configs":                      dcl.ValueOrEmptyString(nr.StreamConfigs),
		"validation_config":                   dcl.ValueOrEmptyString(nr.ValidationConfig),
		"default_search_handling_strict":      dcl.ValueOrEmptyString(nr.DefaultSearchHandlingStrict),
		"complex_data_type_reference_parsing": dcl.ValueOrEmptyString(nr.ComplexDataTypeReferenceParsing),
		"project":                             dcl.ValueOrEmptyString(nr.Project),
		"location":                            dcl.ValueOrEmptyString(nr.Location),
		"dataset":                             dcl.ValueOrEmptyString(nr.Dataset),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/fhirStores/{{name}}", params), nil
}

const FhirStoreMaxPage = -1

type FhirStoreList struct {
	Items []*FhirStore

	nextToken string

	pageSize int32

	resource *FhirStore
}

func (l *FhirStoreList) HasNext() bool {
	return l.nextToken != ""
}

func (l *FhirStoreList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listFhirStore(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListFhirStore(ctx context.Context, project, location, dataset string) (*FhirStoreList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListFhirStoreWithMaxResults(ctx, project, location, dataset, FhirStoreMaxPage)

}

func (c *Client) ListFhirStoreWithMaxResults(ctx context.Context, project, location, dataset string, pageSize int32) (*FhirStoreList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &FhirStore{
		Project:  &project,
		Location: &location,
		Dataset:  &dataset,
	}
	items, token, err := c.listFhirStore(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &FhirStoreList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetFhirStore(ctx context.Context, r *FhirStore) (*FhirStore, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractFhirStoreFields(r)

	b, err := c.getFhirStoreRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalFhirStore(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Dataset = r.Dataset
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFhirStoreNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractFhirStoreFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteFhirStore(ctx context.Context, r *FhirStore) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("FhirStore resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting FhirStore...")
	deleteOp := deleteFhirStoreOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllFhirStore deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllFhirStore(ctx context.Context, project, location, dataset string, filter func(*FhirStore) bool) error {
	listObj, err := c.ListFhirStore(ctx, project, location, dataset)
	if err != nil {
		return err
	}

	err = c.deleteAllFhirStore(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllFhirStore(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyFhirStore(ctx context.Context, rawDesired *FhirStore, opts ...dcl.ApplyOption) (*FhirStore, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *FhirStore
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyFhirStoreHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyFhirStoreHelper(c *Client, ctx context.Context, rawDesired *FhirStore, opts ...dcl.ApplyOption) (*FhirStore, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyFhirStore...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractFhirStoreFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.fhirStoreDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToFhirStoreDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []fhirStoreApiOperation
	if create {
		ops = append(ops, &createFhirStoreOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyFhirStoreDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyFhirStoreDiff(c *Client, ctx context.Context, desired *FhirStore, rawDesired *FhirStore, ops []fhirStoreApiOperation, opts ...dcl.ApplyOption) (*FhirStore, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetFhirStore(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createFhirStoreOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapFhirStore(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeFhirStoreNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFhirStoreNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFhirStoreDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractFhirStoreFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractFhirStoreFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFhirStore(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
