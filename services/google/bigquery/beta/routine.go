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

type Routine struct {
	Etag              *string                      `json:"etag"`
	Name              *string                      `json:"name"`
	Project           *string                      `json:"project"`
	Dataset           *string                      `json:"dataset"`
	RoutineType       *RoutineRoutineTypeEnum      `json:"routineType"`
	CreationTime      *int64                       `json:"creationTime"`
	LastModifiedTime  *int64                       `json:"lastModifiedTime"`
	Language          *RoutineLanguageEnum         `json:"language"`
	Arguments         []RoutineArguments           `json:"arguments"`
	ReturnType        *RoutineArgumentsDataType    `json:"returnType"`
	ImportedLibraries []string                     `json:"importedLibraries"`
	DefinitionBody    *string                      `json:"definitionBody"`
	Description       *string                      `json:"description"`
	DeterminismLevel  *RoutineDeterminismLevelEnum `json:"determinismLevel"`
	StrictMode        *bool                        `json:"strictMode"`
}

func (r *Routine) String() string {
	return dcl.SprintResource(r)
}

// The enum RoutineRoutineTypeEnum.
type RoutineRoutineTypeEnum string

// RoutineRoutineTypeEnumRef returns a *RoutineRoutineTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func RoutineRoutineTypeEnumRef(s string) *RoutineRoutineTypeEnum {
	v := RoutineRoutineTypeEnum(s)
	return &v
}

func (v RoutineRoutineTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ROUTINE_TYPE_UNSPECIFIED", "SCALAR_FUNCTION", "PROCEDURE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RoutineRoutineTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RoutineLanguageEnum.
type RoutineLanguageEnum string

// RoutineLanguageEnumRef returns a *RoutineLanguageEnum with the value of string s
// If the empty string is provided, nil is returned.
func RoutineLanguageEnumRef(s string) *RoutineLanguageEnum {
	v := RoutineLanguageEnum(s)
	return &v
}

func (v RoutineLanguageEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"LANGUAGE_UNSPECIFIED", "SQL", "JAVASCRIPT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RoutineLanguageEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RoutineArgumentsArgumentKindEnum.
type RoutineArgumentsArgumentKindEnum string

// RoutineArgumentsArgumentKindEnumRef returns a *RoutineArgumentsArgumentKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func RoutineArgumentsArgumentKindEnumRef(s string) *RoutineArgumentsArgumentKindEnum {
	v := RoutineArgumentsArgumentKindEnum(s)
	return &v
}

func (v RoutineArgumentsArgumentKindEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ARGUMENT_KIND_UNSPECIFIED", "FIXED_TYPE", "ANY_TYPE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RoutineArgumentsArgumentKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RoutineArgumentsModeEnum.
type RoutineArgumentsModeEnum string

// RoutineArgumentsModeEnumRef returns a *RoutineArgumentsModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func RoutineArgumentsModeEnumRef(s string) *RoutineArgumentsModeEnum {
	v := RoutineArgumentsModeEnum(s)
	return &v
}

func (v RoutineArgumentsModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"MODE_UNSPECIFIED", "IN", "OUT", "INOUT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RoutineArgumentsModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RoutineArgumentsDataTypeTypeKindEnum.
type RoutineArgumentsDataTypeTypeKindEnum string

// RoutineArgumentsDataTypeTypeKindEnumRef returns a *RoutineArgumentsDataTypeTypeKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func RoutineArgumentsDataTypeTypeKindEnumRef(s string) *RoutineArgumentsDataTypeTypeKindEnum {
	v := RoutineArgumentsDataTypeTypeKindEnum(s)
	return &v
}

func (v RoutineArgumentsDataTypeTypeKindEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"TYPE_KIND_UNSPECIFIED", "INT64", "BOOL", "FLOAT64", "STRING", "BYTES", "TIMESTAMP", "DATE", "TIME", "DATETIME", "INTERVAL", "GEOGRAPHY", "NUMERIC", "BIGNUMERIC", "JSON", "ARRAY", "STRUCT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RoutineArgumentsDataTypeTypeKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum RoutineDeterminismLevelEnum.
type RoutineDeterminismLevelEnum string

// RoutineDeterminismLevelEnumRef returns a *RoutineDeterminismLevelEnum with the value of string s
// If the empty string is provided, nil is returned.
func RoutineDeterminismLevelEnumRef(s string) *RoutineDeterminismLevelEnum {
	v := RoutineDeterminismLevelEnum(s)
	return &v
}

func (v RoutineDeterminismLevelEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"DETERMINISM_LEVEL_UNSPECIFIED", "DETERMINISTIC", "NOT_DETERMINISTIC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RoutineDeterminismLevelEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type RoutineArguments struct {
	empty        bool                              `json:"-"`
	Name         *string                           `json:"name"`
	ArgumentKind *RoutineArgumentsArgumentKindEnum `json:"argumentKind"`
	Mode         *RoutineArgumentsModeEnum         `json:"mode"`
	DataType     *RoutineArgumentsDataType         `json:"dataType"`
}

type jsonRoutineArguments RoutineArguments

func (r *RoutineArguments) UnmarshalJSON(data []byte) error {
	var res jsonRoutineArguments
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRoutineArguments
	} else {

		r.Name = res.Name

		r.ArgumentKind = res.ArgumentKind

		r.Mode = res.Mode

		r.DataType = res.DataType

	}
	return nil
}

// This object is used to assert a desired state where this RoutineArguments is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyRoutineArguments *RoutineArguments = &RoutineArguments{empty: true}

func (r *RoutineArguments) Empty() bool {
	return r.empty
}

func (r *RoutineArguments) String() string {
	return dcl.SprintResource(r)
}

func (r *RoutineArguments) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RoutineArgumentsDataType struct {
	empty            bool                                  `json:"-"`
	TypeKind         *RoutineArgumentsDataTypeTypeKindEnum `json:"typeKind"`
	ArrayElementType *RoutineArgumentsDataType             `json:"arrayElementType"`
	StructType       *RoutineArgumentsDataTypeStructType   `json:"structType"`
}

type jsonRoutineArgumentsDataType RoutineArgumentsDataType

func (r *RoutineArgumentsDataType) UnmarshalJSON(data []byte) error {
	var res jsonRoutineArgumentsDataType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRoutineArgumentsDataType
	} else {

		r.TypeKind = res.TypeKind

		r.ArrayElementType = res.ArrayElementType

		r.StructType = res.StructType

	}
	return nil
}

// This object is used to assert a desired state where this RoutineArgumentsDataType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyRoutineArgumentsDataType *RoutineArgumentsDataType = &RoutineArgumentsDataType{empty: true}

func (r *RoutineArgumentsDataType) Empty() bool {
	return r.empty
}

func (r *RoutineArgumentsDataType) String() string {
	return dcl.SprintResource(r)
}

func (r *RoutineArgumentsDataType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RoutineArgumentsDataTypeStructType struct {
	empty  bool                                       `json:"-"`
	Fields []RoutineArgumentsDataTypeStructTypeFields `json:"fields"`
}

type jsonRoutineArgumentsDataTypeStructType RoutineArgumentsDataTypeStructType

func (r *RoutineArgumentsDataTypeStructType) UnmarshalJSON(data []byte) error {
	var res jsonRoutineArgumentsDataTypeStructType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRoutineArgumentsDataTypeStructType
	} else {

		r.Fields = res.Fields

	}
	return nil
}

// This object is used to assert a desired state where this RoutineArgumentsDataTypeStructType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyRoutineArgumentsDataTypeStructType *RoutineArgumentsDataTypeStructType = &RoutineArgumentsDataTypeStructType{empty: true}

func (r *RoutineArgumentsDataTypeStructType) Empty() bool {
	return r.empty
}

func (r *RoutineArgumentsDataTypeStructType) String() string {
	return dcl.SprintResource(r)
}

func (r *RoutineArgumentsDataTypeStructType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RoutineArgumentsDataTypeStructTypeFields struct {
	empty bool                      `json:"-"`
	Name  *string                   `json:"name"`
	Type  *RoutineArgumentsDataType `json:"type"`
}

type jsonRoutineArgumentsDataTypeStructTypeFields RoutineArgumentsDataTypeStructTypeFields

func (r *RoutineArgumentsDataTypeStructTypeFields) UnmarshalJSON(data []byte) error {
	var res jsonRoutineArgumentsDataTypeStructTypeFields
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRoutineArgumentsDataTypeStructTypeFields
	} else {

		r.Name = res.Name

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this RoutineArgumentsDataTypeStructTypeFields is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyRoutineArgumentsDataTypeStructTypeFields *RoutineArgumentsDataTypeStructTypeFields = &RoutineArgumentsDataTypeStructTypeFields{empty: true}

func (r *RoutineArgumentsDataTypeStructTypeFields) Empty() bool {
	return r.empty
}

func (r *RoutineArgumentsDataTypeStructTypeFields) String() string {
	return dcl.SprintResource(r)
}

func (r *RoutineArgumentsDataTypeStructTypeFields) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Routine) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "bigquery",
		Type:    "Routine",
		Version: "beta",
	}
}

func (r *Routine) ID() (string, error) {
	if err := extractRoutineFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"etag":               dcl.ValueOrEmptyString(nr.Etag),
		"name":               dcl.ValueOrEmptyString(nr.Name),
		"project":            dcl.ValueOrEmptyString(nr.Project),
		"dataset":            dcl.ValueOrEmptyString(nr.Dataset),
		"routine_type":       dcl.ValueOrEmptyString(nr.RoutineType),
		"creation_time":      dcl.ValueOrEmptyString(nr.CreationTime),
		"last_modified_time": dcl.ValueOrEmptyString(nr.LastModifiedTime),
		"language":           dcl.ValueOrEmptyString(nr.Language),
		"arguments":          dcl.ValueOrEmptyString(nr.Arguments),
		"return_type":        dcl.ValueOrEmptyString(nr.ReturnType),
		"imported_libraries": dcl.ValueOrEmptyString(nr.ImportedLibraries),
		"definition_body":    dcl.ValueOrEmptyString(nr.DefinitionBody),
		"description":        dcl.ValueOrEmptyString(nr.Description),
		"determinism_level":  dcl.ValueOrEmptyString(nr.DeterminismLevel),
		"strict_mode":        dcl.ValueOrEmptyString(nr.StrictMode),
	}
	return dcl.Nprintf("projects/{{project}}/datasets/{{dataset}}/routines/{{name}}", params), nil
}

const RoutineMaxPage = -1

type RoutineList struct {
	Items []*Routine

	nextToken string

	pageSize int32

	resource *Routine
}

func (l *RoutineList) HasNext() bool {
	return l.nextToken != ""
}

func (l *RoutineList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listRoutine(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListRoutine(ctx context.Context, project, dataset string) (*RoutineList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListRoutineWithMaxResults(ctx, project, dataset, RoutineMaxPage)

}

func (c *Client) ListRoutineWithMaxResults(ctx context.Context, project, dataset string, pageSize int32) (*RoutineList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Routine{
		Project: &project,
		Dataset: &dataset,
	}
	items, token, err := c.listRoutine(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &RoutineList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetRoutine(ctx context.Context, r *Routine) (*Routine, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractRoutineFields(r)

	b, err := c.getRoutineRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalRoutine(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Dataset = r.Dataset
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeRoutineNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractRoutineFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteRoutine(ctx context.Context, r *Routine) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Routine resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Routine...")
	deleteOp := deleteRoutineOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllRoutine deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllRoutine(ctx context.Context, project, dataset string, filter func(*Routine) bool) error {
	listObj, err := c.ListRoutine(ctx, project, dataset)
	if err != nil {
		return err
	}

	err = c.deleteAllRoutine(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllRoutine(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyRoutine(ctx context.Context, rawDesired *Routine, opts ...dcl.ApplyOption) (*Routine, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Routine
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyRoutineHelper(c, ctx, rawDesired, opts...)
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

func applyRoutineHelper(c *Client, ctx context.Context, rawDesired *Routine, opts ...dcl.ApplyOption) (*Routine, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyRoutine...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractRoutineFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.routineDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToRoutineDiffs(c.Config, fieldDiffs, opts)
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
	var ops []routineApiOperation
	if create {
		ops = append(ops, &createRoutineOperation{})
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
	return applyRoutineDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyRoutineDiff(c *Client, ctx context.Context, desired *Routine, rawDesired *Routine, ops []routineApiOperation, opts ...dcl.ApplyOption) (*Routine, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetRoutine(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createRoutineOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapRoutine(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeRoutineNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeRoutineNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeRoutineDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractRoutineFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractRoutineFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffRoutine(c, newDesired, newState)
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
