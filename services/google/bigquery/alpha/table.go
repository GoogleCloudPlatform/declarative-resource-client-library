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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Table struct {
	Etag                      *string                         `json:"etag"`
	Id                        *string                         `json:"id"`
	SelfLink                  *string                         `json:"selfLink"`
	Name                      *string                         `json:"name"`
	Dataset                   *string                         `json:"dataset"`
	Project                   *string                         `json:"project"`
	FriendlyName              *string                         `json:"friendlyName"`
	Description               *string                         `json:"description"`
	Labels                    map[string]string               `json:"labels"`
	Model                     *TableModel                     `json:"model"`
	Schema                    *TableSchema                    `json:"schema"`
	TimePartitioning          *TableTimePartitioning          `json:"timePartitioning"`
	RangePartitioning         *TableRangePartitioning         `json:"rangePartitioning"`
	Clustering                *TableClustering                `json:"clustering"`
	RequirePartitionFilter    *bool                           `json:"requirePartitionFilter"`
	NumBytes                  *string                         `json:"numBytes"`
	NumPhysicalBytes          *string                         `json:"numPhysicalBytes"`
	NumLongTermBytes          *string                         `json:"numLongTermBytes"`
	NumRows                   *int64                          `json:"numRows"`
	CreationTime              *int64                          `json:"creationTime"`
	ExpirationTime            *int64                          `json:"expirationTime"`
	LastModifiedTime          *int64                          `json:"lastModifiedTime"`
	Type                      *string                         `json:"type"`
	View                      *TableView                      `json:"view"`
	MaterializedView          *TableMaterializedView          `json:"materializedView"`
	ExternalDataConfiguration *TableExternalDataConfiguration `json:"externalDataConfiguration"`
	Location                  *string                         `json:"location"`
	StreamingBuffer           *TableStreamingBuffer           `json:"streamingBuffer"`
	EncryptionConfiguration   *TableEncryptionConfiguration   `json:"encryptionConfiguration"`
	SnapshotDefinition        *TableSnapshotDefinition        `json:"snapshotDefinition"`
	DefaultCollation          *string                         `json:"defaultCollation"`
}

func (r *Table) String() string {
	return dcl.SprintResource(r)
}

// The enum TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum.
type TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum string

// TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumRef returns a *TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumRef(s string) *TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	v := TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(s)
	return &v
}

func (v TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"TEMPORAL_TYPES_VALUE_UNSPECIFIED", "TEMPORAL_TYPES_VALUE_REJECT", "TEMPORAL_TYPES_VALUE_CLAMP"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum.
type TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum string

// TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumRef returns a *TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumRef(s string) *TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	v := TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(s)
	return &v
}

func (v TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"NUMERIC_TYPE_VALUE_UNSPECIFIED", "NUMERIC_TYPE_VALUE_REJECT", "NUMERIC_TYPE_VALUE_ROUND"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum TableExternalDataConfigurationDecimalTargetTypesEnum.
type TableExternalDataConfigurationDecimalTargetTypesEnum string

// TableExternalDataConfigurationDecimalTargetTypesEnumRef returns a *TableExternalDataConfigurationDecimalTargetTypesEnum with the value of string s
// If the empty string is provided, nil is returned.
func TableExternalDataConfigurationDecimalTargetTypesEnumRef(s string) *TableExternalDataConfigurationDecimalTargetTypesEnum {
	v := TableExternalDataConfigurationDecimalTargetTypesEnum(s)
	return &v
}

func (v TableExternalDataConfigurationDecimalTargetTypesEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"DECIMAL_TARGET_TYPE_UNSPECIFIED", "NUMERIC", "BIGNUMERIC", "STRING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TableExternalDataConfigurationDecimalTargetTypesEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum TableExternalDataConfigurationJsonExtensionEnum.
type TableExternalDataConfigurationJsonExtensionEnum string

// TableExternalDataConfigurationJsonExtensionEnumRef returns a *TableExternalDataConfigurationJsonExtensionEnum with the value of string s
// If the empty string is provided, nil is returned.
func TableExternalDataConfigurationJsonExtensionEnumRef(s string) *TableExternalDataConfigurationJsonExtensionEnum {
	v := TableExternalDataConfigurationJsonExtensionEnum(s)
	return &v
}

func (v TableExternalDataConfigurationJsonExtensionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"JSON_EXTENSION_UNSPECIFIED", "GEOJSON"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TableExternalDataConfigurationJsonExtensionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type TableModel struct {
	empty        bool                     `json:"-"`
	ModelOptions *TableModelModelOptions  `json:"modelOptions"`
	TrainingRuns []TableModelTrainingRuns `json:"trainingRuns"`
}

type jsonTableModel TableModel

func (r *TableModel) UnmarshalJSON(data []byte) error {
	var res jsonTableModel
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableModel
	} else {

		r.ModelOptions = res.ModelOptions

		r.TrainingRuns = res.TrainingRuns

	}
	return nil
}

// This object is used to assert a desired state where this TableModel is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableModel *TableModel = &TableModel{empty: true}

func (r *TableModel) Empty() bool {
	return r.empty
}

func (r *TableModel) String() string {
	return dcl.SprintResource(r)
}

func (r *TableModel) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableModelModelOptions struct {
	empty     bool     `json:"-"`
	ModelType *string  `json:"modelType"`
	Labels    []string `json:"labels"`
	LossType  *string  `json:"lossType"`
}

type jsonTableModelModelOptions TableModelModelOptions

func (r *TableModelModelOptions) UnmarshalJSON(data []byte) error {
	var res jsonTableModelModelOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableModelModelOptions
	} else {

		r.ModelType = res.ModelType

		r.Labels = res.Labels

		r.LossType = res.LossType

	}
	return nil
}

// This object is used to assert a desired state where this TableModelModelOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableModelModelOptions *TableModelModelOptions = &TableModelModelOptions{empty: true}

func (r *TableModelModelOptions) Empty() bool {
	return r.empty
}

func (r *TableModelModelOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *TableModelModelOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableModelTrainingRuns struct {
	empty            bool                                     `json:"-"`
	State            *string                                  `json:"state"`
	StartTime        *string                                  `json:"startTime"`
	TrainingOptions  *TableModelTrainingRunsTrainingOptions   `json:"trainingOptions"`
	IterationResults []TableModelTrainingRunsIterationResults `json:"iterationResults"`
}

type jsonTableModelTrainingRuns TableModelTrainingRuns

func (r *TableModelTrainingRuns) UnmarshalJSON(data []byte) error {
	var res jsonTableModelTrainingRuns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableModelTrainingRuns
	} else {

		r.State = res.State

		r.StartTime = res.StartTime

		r.TrainingOptions = res.TrainingOptions

		r.IterationResults = res.IterationResults

	}
	return nil
}

// This object is used to assert a desired state where this TableModelTrainingRuns is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableModelTrainingRuns *TableModelTrainingRuns = &TableModelTrainingRuns{empty: true}

func (r *TableModelTrainingRuns) Empty() bool {
	return r.empty
}

func (r *TableModelTrainingRuns) String() string {
	return dcl.SprintResource(r)
}

func (r *TableModelTrainingRuns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableModelTrainingRunsTrainingOptions struct {
	empty                   bool     `json:"-"`
	MaxIteration            *int64   `json:"maxIteration"`
	LearnRate               *float64 `json:"learnRate"`
	L1Reg                   *float64 `json:"l1Reg"`
	L2Reg                   *float64 `json:"l2Reg"`
	MinRelProgress          *float64 `json:"minRelProgress"`
	WarmStart               *bool    `json:"warmStart"`
	EarlyStop               *bool    `json:"earlyStop"`
	LearnRateStrategy       *string  `json:"learnRateStrategy"`
	LineSearchInitLearnRate *float64 `json:"lineSearchInitLearnRate"`
}

type jsonTableModelTrainingRunsTrainingOptions TableModelTrainingRunsTrainingOptions

func (r *TableModelTrainingRunsTrainingOptions) UnmarshalJSON(data []byte) error {
	var res jsonTableModelTrainingRunsTrainingOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableModelTrainingRunsTrainingOptions
	} else {

		r.MaxIteration = res.MaxIteration

		r.LearnRate = res.LearnRate

		r.L1Reg = res.L1Reg

		r.L2Reg = res.L2Reg

		r.MinRelProgress = res.MinRelProgress

		r.WarmStart = res.WarmStart

		r.EarlyStop = res.EarlyStop

		r.LearnRateStrategy = res.LearnRateStrategy

		r.LineSearchInitLearnRate = res.LineSearchInitLearnRate

	}
	return nil
}

// This object is used to assert a desired state where this TableModelTrainingRunsTrainingOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableModelTrainingRunsTrainingOptions *TableModelTrainingRunsTrainingOptions = &TableModelTrainingRunsTrainingOptions{empty: true}

func (r *TableModelTrainingRunsTrainingOptions) Empty() bool {
	return r.empty
}

func (r *TableModelTrainingRunsTrainingOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *TableModelTrainingRunsTrainingOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableModelTrainingRunsIterationResults struct {
	empty        bool     `json:"-"`
	Index        *int64   `json:"index"`
	LearnRate    *float64 `json:"learnRate"`
	TrainingLoss *float64 `json:"trainingLoss"`
	EvalLoss     *float64 `json:"evalLoss"`
	DurationMs   *string  `json:"durationMs"`
}

type jsonTableModelTrainingRunsIterationResults TableModelTrainingRunsIterationResults

func (r *TableModelTrainingRunsIterationResults) UnmarshalJSON(data []byte) error {
	var res jsonTableModelTrainingRunsIterationResults
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableModelTrainingRunsIterationResults
	} else {

		r.Index = res.Index

		r.LearnRate = res.LearnRate

		r.TrainingLoss = res.TrainingLoss

		r.EvalLoss = res.EvalLoss

		r.DurationMs = res.DurationMs

	}
	return nil
}

// This object is used to assert a desired state where this TableModelTrainingRunsIterationResults is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableModelTrainingRunsIterationResults *TableModelTrainingRunsIterationResults = &TableModelTrainingRunsIterationResults{empty: true}

func (r *TableModelTrainingRunsIterationResults) Empty() bool {
	return r.empty
}

func (r *TableModelTrainingRunsIterationResults) String() string {
	return dcl.SprintResource(r)
}

func (r *TableModelTrainingRunsIterationResults) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableSchema struct {
	empty  bool                                         `json:"-"`
	Fields []TableGooglecloudbigqueryv2Tablefieldschema `json:"fields"`
}

type jsonTableSchema TableSchema

func (r *TableSchema) UnmarshalJSON(data []byte) error {
	var res jsonTableSchema
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableSchema
	} else {

		r.Fields = res.Fields

	}
	return nil
}

// This object is used to assert a desired state where this TableSchema is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableSchema *TableSchema = &TableSchema{empty: true}

func (r *TableSchema) Empty() bool {
	return r.empty
}

func (r *TableSchema) String() string {
	return dcl.SprintResource(r)
}

func (r *TableSchema) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableGooglecloudbigqueryv2Tablefieldschema struct {
	empty                  bool                                                  `json:"-"`
	Name                   *string                                               `json:"name"`
	Type                   *string                                               `json:"type"`
	Mode                   *string                                               `json:"mode"`
	Fields                 []TableGooglecloudbigqueryv2Tablefieldschema          `json:"fields"`
	Description            *string                                               `json:"description"`
	Categories             *TableGooglecloudbigqueryv2TablefieldschemaCategories `json:"categories"`
	PolicyTags             *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags `json:"policyTags"`
	NameAlternative        []string                                              `json:"nameAlternative"`
	MaxLength              *int64                                                `json:"maxLength"`
	Precision              *int64                                                `json:"precision"`
	Scale                  *int64                                                `json:"scale"`
	Collation              *string                                               `json:"collation"`
	DefaultValueExpression *string                                               `json:"defaultValueExpression"`
}

type jsonTableGooglecloudbigqueryv2Tablefieldschema TableGooglecloudbigqueryv2Tablefieldschema

func (r *TableGooglecloudbigqueryv2Tablefieldschema) UnmarshalJSON(data []byte) error {
	var res jsonTableGooglecloudbigqueryv2Tablefieldschema
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableGooglecloudbigqueryv2Tablefieldschema
	} else {

		r.Name = res.Name

		r.Type = res.Type

		r.Mode = res.Mode

		r.Fields = res.Fields

		r.Description = res.Description

		r.Categories = res.Categories

		r.PolicyTags = res.PolicyTags

		r.NameAlternative = res.NameAlternative

		r.MaxLength = res.MaxLength

		r.Precision = res.Precision

		r.Scale = res.Scale

		r.Collation = res.Collation

		r.DefaultValueExpression = res.DefaultValueExpression

	}
	return nil
}

// This object is used to assert a desired state where this TableGooglecloudbigqueryv2Tablefieldschema is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableGooglecloudbigqueryv2Tablefieldschema *TableGooglecloudbigqueryv2Tablefieldschema = &TableGooglecloudbigqueryv2Tablefieldschema{empty: true}

func (r *TableGooglecloudbigqueryv2Tablefieldschema) Empty() bool {
	return r.empty
}

func (r *TableGooglecloudbigqueryv2Tablefieldschema) String() string {
	return dcl.SprintResource(r)
}

func (r *TableGooglecloudbigqueryv2Tablefieldschema) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableGooglecloudbigqueryv2TablefieldschemaCategories struct {
	empty bool     `json:"-"`
	Names []string `json:"names"`
}

type jsonTableGooglecloudbigqueryv2TablefieldschemaCategories TableGooglecloudbigqueryv2TablefieldschemaCategories

func (r *TableGooglecloudbigqueryv2TablefieldschemaCategories) UnmarshalJSON(data []byte) error {
	var res jsonTableGooglecloudbigqueryv2TablefieldschemaCategories
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableGooglecloudbigqueryv2TablefieldschemaCategories
	} else {

		r.Names = res.Names

	}
	return nil
}

// This object is used to assert a desired state where this TableGooglecloudbigqueryv2TablefieldschemaCategories is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableGooglecloudbigqueryv2TablefieldschemaCategories *TableGooglecloudbigqueryv2TablefieldschemaCategories = &TableGooglecloudbigqueryv2TablefieldschemaCategories{empty: true}

func (r *TableGooglecloudbigqueryv2TablefieldschemaCategories) Empty() bool {
	return r.empty
}

func (r *TableGooglecloudbigqueryv2TablefieldschemaCategories) String() string {
	return dcl.SprintResource(r)
}

func (r *TableGooglecloudbigqueryv2TablefieldschemaCategories) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableGooglecloudbigqueryv2TablefieldschemaPolicyTags struct {
	empty bool     `json:"-"`
	Names []string `json:"names"`
}

type jsonTableGooglecloudbigqueryv2TablefieldschemaPolicyTags TableGooglecloudbigqueryv2TablefieldschemaPolicyTags

func (r *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) UnmarshalJSON(data []byte) error {
	var res jsonTableGooglecloudbigqueryv2TablefieldschemaPolicyTags
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableGooglecloudbigqueryv2TablefieldschemaPolicyTags
	} else {

		r.Names = res.Names

	}
	return nil
}

// This object is used to assert a desired state where this TableGooglecloudbigqueryv2TablefieldschemaPolicyTags is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableGooglecloudbigqueryv2TablefieldschemaPolicyTags *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags = &TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{empty: true}

func (r *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) Empty() bool {
	return r.empty
}

func (r *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) String() string {
	return dcl.SprintResource(r)
}

func (r *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableTimePartitioning struct {
	empty        bool    `json:"-"`
	Type         *string `json:"type"`
	ExpirationMs *string `json:"expirationMs"`
	Field        *string `json:"field"`
}

type jsonTableTimePartitioning TableTimePartitioning

func (r *TableTimePartitioning) UnmarshalJSON(data []byte) error {
	var res jsonTableTimePartitioning
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableTimePartitioning
	} else {

		r.Type = res.Type

		r.ExpirationMs = res.ExpirationMs

		r.Field = res.Field

	}
	return nil
}

// This object is used to assert a desired state where this TableTimePartitioning is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableTimePartitioning *TableTimePartitioning = &TableTimePartitioning{empty: true}

func (r *TableTimePartitioning) Empty() bool {
	return r.empty
}

func (r *TableTimePartitioning) String() string {
	return dcl.SprintResource(r)
}

func (r *TableTimePartitioning) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableRangePartitioning struct {
	empty bool                         `json:"-"`
	Field *string                      `json:"field"`
	Range *TableRangePartitioningRange `json:"range"`
}

type jsonTableRangePartitioning TableRangePartitioning

func (r *TableRangePartitioning) UnmarshalJSON(data []byte) error {
	var res jsonTableRangePartitioning
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableRangePartitioning
	} else {

		r.Field = res.Field

		r.Range = res.Range

	}
	return nil
}

// This object is used to assert a desired state where this TableRangePartitioning is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableRangePartitioning *TableRangePartitioning = &TableRangePartitioning{empty: true}

func (r *TableRangePartitioning) Empty() bool {
	return r.empty
}

func (r *TableRangePartitioning) String() string {
	return dcl.SprintResource(r)
}

func (r *TableRangePartitioning) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableRangePartitioningRange struct {
	empty    bool    `json:"-"`
	Start    *string `json:"start"`
	End      *string `json:"end"`
	Interval *string `json:"interval"`
}

type jsonTableRangePartitioningRange TableRangePartitioningRange

func (r *TableRangePartitioningRange) UnmarshalJSON(data []byte) error {
	var res jsonTableRangePartitioningRange
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableRangePartitioningRange
	} else {

		r.Start = res.Start

		r.End = res.End

		r.Interval = res.Interval

	}
	return nil
}

// This object is used to assert a desired state where this TableRangePartitioningRange is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableRangePartitioningRange *TableRangePartitioningRange = &TableRangePartitioningRange{empty: true}

func (r *TableRangePartitioningRange) Empty() bool {
	return r.empty
}

func (r *TableRangePartitioningRange) String() string {
	return dcl.SprintResource(r)
}

func (r *TableRangePartitioningRange) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableClustering struct {
	empty  bool     `json:"-"`
	Fields []string `json:"fields"`
}

type jsonTableClustering TableClustering

func (r *TableClustering) UnmarshalJSON(data []byte) error {
	var res jsonTableClustering
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableClustering
	} else {

		r.Fields = res.Fields

	}
	return nil
}

// This object is used to assert a desired state where this TableClustering is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableClustering *TableClustering = &TableClustering{empty: true}

func (r *TableClustering) Empty() bool {
	return r.empty
}

func (r *TableClustering) String() string {
	return dcl.SprintResource(r)
}

func (r *TableClustering) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableView struct {
	empty                        bool                                    `json:"-"`
	Query                        *string                                 `json:"query"`
	UserDefinedFunctionResources []TableViewUserDefinedFunctionResources `json:"userDefinedFunctionResources"`
	UseLegacySql                 *bool                                   `json:"useLegacySql"`
	UseExplicitColumnNames       *bool                                   `json:"useExplicitColumnNames"`
}

type jsonTableView TableView

func (r *TableView) UnmarshalJSON(data []byte) error {
	var res jsonTableView
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableView
	} else {

		r.Query = res.Query

		r.UserDefinedFunctionResources = res.UserDefinedFunctionResources

		r.UseLegacySql = res.UseLegacySql

		r.UseExplicitColumnNames = res.UseExplicitColumnNames

	}
	return nil
}

// This object is used to assert a desired state where this TableView is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableView *TableView = &TableView{empty: true}

func (r *TableView) Empty() bool {
	return r.empty
}

func (r *TableView) String() string {
	return dcl.SprintResource(r)
}

func (r *TableView) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableViewUserDefinedFunctionResources struct {
	empty                 bool     `json:"-"`
	ResourceUri           *string  `json:"resourceUri"`
	InlineCode            *string  `json:"inlineCode"`
	InlineCodeAlternative []string `json:"inlineCodeAlternative"`
}

type jsonTableViewUserDefinedFunctionResources TableViewUserDefinedFunctionResources

func (r *TableViewUserDefinedFunctionResources) UnmarshalJSON(data []byte) error {
	var res jsonTableViewUserDefinedFunctionResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableViewUserDefinedFunctionResources
	} else {

		r.ResourceUri = res.ResourceUri

		r.InlineCode = res.InlineCode

		r.InlineCodeAlternative = res.InlineCodeAlternative

	}
	return nil
}

// This object is used to assert a desired state where this TableViewUserDefinedFunctionResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableViewUserDefinedFunctionResources *TableViewUserDefinedFunctionResources = &TableViewUserDefinedFunctionResources{empty: true}

func (r *TableViewUserDefinedFunctionResources) Empty() bool {
	return r.empty
}

func (r *TableViewUserDefinedFunctionResources) String() string {
	return dcl.SprintResource(r)
}

func (r *TableViewUserDefinedFunctionResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableMaterializedView struct {
	empty             bool    `json:"-"`
	Query             *string `json:"query"`
	LastRefreshTime   *int64  `json:"lastRefreshTime"`
	EnableRefresh     *bool   `json:"enableRefresh"`
	RefreshIntervalMs *int64  `json:"refreshIntervalMs"`
}

type jsonTableMaterializedView TableMaterializedView

func (r *TableMaterializedView) UnmarshalJSON(data []byte) error {
	var res jsonTableMaterializedView
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableMaterializedView
	} else {

		r.Query = res.Query

		r.LastRefreshTime = res.LastRefreshTime

		r.EnableRefresh = res.EnableRefresh

		r.RefreshIntervalMs = res.RefreshIntervalMs

	}
	return nil
}

// This object is used to assert a desired state where this TableMaterializedView is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableMaterializedView *TableMaterializedView = &TableMaterializedView{empty: true}

func (r *TableMaterializedView) Empty() bool {
	return r.empty
}

func (r *TableMaterializedView) String() string {
	return dcl.SprintResource(r)
}

func (r *TableMaterializedView) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfiguration struct {
	empty                    bool                                                   `json:"-"`
	SourceUris               []string                                               `json:"sourceUris"`
	Schema                   *TableExternalDataConfigurationSchema                  `json:"schema"`
	SourceFormat             *string                                                `json:"sourceFormat"`
	MaxBadRecords            *int64                                                 `json:"maxBadRecords"`
	Autodetect               *bool                                                  `json:"autodetect"`
	IgnoreUnknownValues      *bool                                                  `json:"ignoreUnknownValues"`
	Compression              *string                                                `json:"compression"`
	CsvOptions               *TableExternalDataConfigurationCsvOptions              `json:"csvOptions"`
	BigtableOptions          *TableExternalDataConfigurationBigtableOptions         `json:"bigtableOptions"`
	GoogleSheetsOptions      *TableExternalDataConfigurationGoogleSheetsOptions     `json:"googleSheetsOptions"`
	MaxBadRecordsAlternative []int64                                                `json:"maxBadRecordsAlternative"`
	HivePartitioningOptions  *TableExternalDataConfigurationHivePartitioningOptions `json:"hivePartitioningOptions"`
	ConnectionId             *string                                                `json:"connectionId"`
	ValueConversionModes     *TableExternalDataConfigurationValueConversionModes    `json:"valueConversionModes"`
	DecimalTargetTypes       []TableExternalDataConfigurationDecimalTargetTypesEnum `json:"decimalTargetTypes"`
	AvroOptions              *TableExternalDataConfigurationAvroOptions             `json:"avroOptions"`
	JsonExtension            *TableExternalDataConfigurationJsonExtensionEnum       `json:"jsonExtension"`
	ParquetOptions           *TableExternalDataConfigurationParquetOptions          `json:"parquetOptions"`
}

type jsonTableExternalDataConfiguration TableExternalDataConfiguration

func (r *TableExternalDataConfiguration) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfiguration
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfiguration
	} else {

		r.SourceUris = res.SourceUris

		r.Schema = res.Schema

		r.SourceFormat = res.SourceFormat

		r.MaxBadRecords = res.MaxBadRecords

		r.Autodetect = res.Autodetect

		r.IgnoreUnknownValues = res.IgnoreUnknownValues

		r.Compression = res.Compression

		r.CsvOptions = res.CsvOptions

		r.BigtableOptions = res.BigtableOptions

		r.GoogleSheetsOptions = res.GoogleSheetsOptions

		r.MaxBadRecordsAlternative = res.MaxBadRecordsAlternative

		r.HivePartitioningOptions = res.HivePartitioningOptions

		r.ConnectionId = res.ConnectionId

		r.ValueConversionModes = res.ValueConversionModes

		r.DecimalTargetTypes = res.DecimalTargetTypes

		r.AvroOptions = res.AvroOptions

		r.JsonExtension = res.JsonExtension

		r.ParquetOptions = res.ParquetOptions

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfiguration is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfiguration *TableExternalDataConfiguration = &TableExternalDataConfiguration{empty: true}

func (r *TableExternalDataConfiguration) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationSchema struct {
	empty  bool                                         `json:"-"`
	Fields []TableGooglecloudbigqueryv2Tablefieldschema `json:"fields"`
}

type jsonTableExternalDataConfigurationSchema TableExternalDataConfigurationSchema

func (r *TableExternalDataConfigurationSchema) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationSchema
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationSchema
	} else {

		r.Fields = res.Fields

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationSchema is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationSchema *TableExternalDataConfigurationSchema = &TableExternalDataConfigurationSchema{empty: true}

func (r *TableExternalDataConfigurationSchema) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationSchema) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationSchema) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationCsvOptions struct {
	empty               bool    `json:"-"`
	FieldDelimiter      *string `json:"fieldDelimiter"`
	SkipLeadingRows     *string `json:"skipLeadingRows"`
	Quote               *string `json:"quote"`
	AllowQuotedNewlines *bool   `json:"allowQuotedNewlines"`
	AllowJaggedRows     *bool   `json:"allowJaggedRows"`
	Encoding            *string `json:"encoding"`
}

type jsonTableExternalDataConfigurationCsvOptions TableExternalDataConfigurationCsvOptions

func (r *TableExternalDataConfigurationCsvOptions) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationCsvOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationCsvOptions
	} else {

		r.FieldDelimiter = res.FieldDelimiter

		r.SkipLeadingRows = res.SkipLeadingRows

		r.Quote = res.Quote

		r.AllowQuotedNewlines = res.AllowQuotedNewlines

		r.AllowJaggedRows = res.AllowJaggedRows

		r.Encoding = res.Encoding

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationCsvOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationCsvOptions *TableExternalDataConfigurationCsvOptions = &TableExternalDataConfigurationCsvOptions{empty: true}

func (r *TableExternalDataConfigurationCsvOptions) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationCsvOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationCsvOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationBigtableOptions struct {
	empty                           bool                                                          `json:"-"`
	ColumnFamilies                  []TableExternalDataConfigurationBigtableOptionsColumnFamilies `json:"columnFamilies"`
	IgnoreUnspecifiedColumnFamilies *bool                                                         `json:"ignoreUnspecifiedColumnFamilies"`
	ReadRowkeyAsString              *bool                                                         `json:"readRowkeyAsString"`
}

type jsonTableExternalDataConfigurationBigtableOptions TableExternalDataConfigurationBigtableOptions

func (r *TableExternalDataConfigurationBigtableOptions) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationBigtableOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationBigtableOptions
	} else {

		r.ColumnFamilies = res.ColumnFamilies

		r.IgnoreUnspecifiedColumnFamilies = res.IgnoreUnspecifiedColumnFamilies

		r.ReadRowkeyAsString = res.ReadRowkeyAsString

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationBigtableOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationBigtableOptions *TableExternalDataConfigurationBigtableOptions = &TableExternalDataConfigurationBigtableOptions{empty: true}

func (r *TableExternalDataConfigurationBigtableOptions) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationBigtableOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationBigtableOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationBigtableOptionsColumnFamilies struct {
	empty          bool                                                                 `json:"-"`
	FamilyId       *string                                                              `json:"familyId"`
	Type           *string                                                              `json:"type"`
	Encoding       *string                                                              `json:"encoding"`
	Columns        []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns `json:"columns"`
	OnlyReadLatest *bool                                                                `json:"onlyReadLatest"`
}

type jsonTableExternalDataConfigurationBigtableOptionsColumnFamilies TableExternalDataConfigurationBigtableOptionsColumnFamilies

func (r *TableExternalDataConfigurationBigtableOptionsColumnFamilies) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationBigtableOptionsColumnFamilies
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationBigtableOptionsColumnFamilies
	} else {

		r.FamilyId = res.FamilyId

		r.Type = res.Type

		r.Encoding = res.Encoding

		r.Columns = res.Columns

		r.OnlyReadLatest = res.OnlyReadLatest

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationBigtableOptionsColumnFamilies is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationBigtableOptionsColumnFamilies *TableExternalDataConfigurationBigtableOptionsColumnFamilies = &TableExternalDataConfigurationBigtableOptionsColumnFamilies{empty: true}

func (r *TableExternalDataConfigurationBigtableOptionsColumnFamilies) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationBigtableOptionsColumnFamilies) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationBigtableOptionsColumnFamilies) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns struct {
	empty            bool    `json:"-"`
	QualifierEncoded *string `json:"qualifierEncoded"`
	QualifierString  *string `json:"qualifierString"`
	FieldName        *string `json:"fieldName"`
	Type             *string `json:"type"`
	Encoding         *string `json:"encoding"`
	OnlyReadLatest   *bool   `json:"onlyReadLatest"`
}

type jsonTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns

func (r *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns
	} else {

		r.QualifierEncoded = res.QualifierEncoded

		r.QualifierString = res.QualifierString

		r.FieldName = res.FieldName

		r.Type = res.Type

		r.Encoding = res.Encoding

		r.OnlyReadLatest = res.OnlyReadLatest

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns = &TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{empty: true}

func (r *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationGoogleSheetsOptions struct {
	empty           bool    `json:"-"`
	SkipLeadingRows *string `json:"skipLeadingRows"`
	Range           *string `json:"range"`
}

type jsonTableExternalDataConfigurationGoogleSheetsOptions TableExternalDataConfigurationGoogleSheetsOptions

func (r *TableExternalDataConfigurationGoogleSheetsOptions) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationGoogleSheetsOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationGoogleSheetsOptions
	} else {

		r.SkipLeadingRows = res.SkipLeadingRows

		r.Range = res.Range

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationGoogleSheetsOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationGoogleSheetsOptions *TableExternalDataConfigurationGoogleSheetsOptions = &TableExternalDataConfigurationGoogleSheetsOptions{empty: true}

func (r *TableExternalDataConfigurationGoogleSheetsOptions) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationGoogleSheetsOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationGoogleSheetsOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationHivePartitioningOptions struct {
	empty                  bool     `json:"-"`
	Mode                   *string  `json:"mode"`
	SourceUriPrefix        *string  `json:"sourceUriPrefix"`
	RequirePartitionFilter *bool    `json:"requirePartitionFilter"`
	Fields                 []string `json:"fields"`
}

type jsonTableExternalDataConfigurationHivePartitioningOptions TableExternalDataConfigurationHivePartitioningOptions

func (r *TableExternalDataConfigurationHivePartitioningOptions) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationHivePartitioningOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationHivePartitioningOptions
	} else {

		r.Mode = res.Mode

		r.SourceUriPrefix = res.SourceUriPrefix

		r.RequirePartitionFilter = res.RequirePartitionFilter

		r.Fields = res.Fields

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationHivePartitioningOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationHivePartitioningOptions *TableExternalDataConfigurationHivePartitioningOptions = &TableExternalDataConfigurationHivePartitioningOptions{empty: true}

func (r *TableExternalDataConfigurationHivePartitioningOptions) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationHivePartitioningOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationHivePartitioningOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationValueConversionModes struct {
	empty                                 bool                                                                                         `json:"-"`
	TemporalTypesOutOfRangeConversionMode *TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum `json:"temporalTypesOutOfRangeConversionMode"`
	NumericTypeOutOfRangeConversionMode   *TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum   `json:"numericTypeOutOfRangeConversionMode"`
}

type jsonTableExternalDataConfigurationValueConversionModes TableExternalDataConfigurationValueConversionModes

func (r *TableExternalDataConfigurationValueConversionModes) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationValueConversionModes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationValueConversionModes
	} else {

		r.TemporalTypesOutOfRangeConversionMode = res.TemporalTypesOutOfRangeConversionMode

		r.NumericTypeOutOfRangeConversionMode = res.NumericTypeOutOfRangeConversionMode

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationValueConversionModes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationValueConversionModes *TableExternalDataConfigurationValueConversionModes = &TableExternalDataConfigurationValueConversionModes{empty: true}

func (r *TableExternalDataConfigurationValueConversionModes) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationValueConversionModes) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationValueConversionModes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationAvroOptions struct {
	empty               bool  `json:"-"`
	UseAvroLogicalTypes *bool `json:"useAvroLogicalTypes"`
}

type jsonTableExternalDataConfigurationAvroOptions TableExternalDataConfigurationAvroOptions

func (r *TableExternalDataConfigurationAvroOptions) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationAvroOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationAvroOptions
	} else {

		r.UseAvroLogicalTypes = res.UseAvroLogicalTypes

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationAvroOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationAvroOptions *TableExternalDataConfigurationAvroOptions = &TableExternalDataConfigurationAvroOptions{empty: true}

func (r *TableExternalDataConfigurationAvroOptions) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationAvroOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationAvroOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableExternalDataConfigurationParquetOptions struct {
	empty               bool  `json:"-"`
	EnumAsString        *bool `json:"enumAsString"`
	EnableListInference *bool `json:"enableListInference"`
}

type jsonTableExternalDataConfigurationParquetOptions TableExternalDataConfigurationParquetOptions

func (r *TableExternalDataConfigurationParquetOptions) UnmarshalJSON(data []byte) error {
	var res jsonTableExternalDataConfigurationParquetOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableExternalDataConfigurationParquetOptions
	} else {

		r.EnumAsString = res.EnumAsString

		r.EnableListInference = res.EnableListInference

	}
	return nil
}

// This object is used to assert a desired state where this TableExternalDataConfigurationParquetOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableExternalDataConfigurationParquetOptions *TableExternalDataConfigurationParquetOptions = &TableExternalDataConfigurationParquetOptions{empty: true}

func (r *TableExternalDataConfigurationParquetOptions) Empty() bool {
	return r.empty
}

func (r *TableExternalDataConfigurationParquetOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *TableExternalDataConfigurationParquetOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableStreamingBuffer struct {
	empty           bool   `json:"-"`
	EstimatedBytes  *int64 `json:"estimatedBytes"`
	EstimatedRows   *int64 `json:"estimatedRows"`
	OldestEntryTime *int64 `json:"oldestEntryTime"`
}

type jsonTableStreamingBuffer TableStreamingBuffer

func (r *TableStreamingBuffer) UnmarshalJSON(data []byte) error {
	var res jsonTableStreamingBuffer
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableStreamingBuffer
	} else {

		r.EstimatedBytes = res.EstimatedBytes

		r.EstimatedRows = res.EstimatedRows

		r.OldestEntryTime = res.OldestEntryTime

	}
	return nil
}

// This object is used to assert a desired state where this TableStreamingBuffer is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableStreamingBuffer *TableStreamingBuffer = &TableStreamingBuffer{empty: true}

func (r *TableStreamingBuffer) Empty() bool {
	return r.empty
}

func (r *TableStreamingBuffer) String() string {
	return dcl.SprintResource(r)
}

func (r *TableStreamingBuffer) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableEncryptionConfiguration struct {
	empty      bool    `json:"-"`
	KmsKeyName *string `json:"kmsKeyName"`
}

type jsonTableEncryptionConfiguration TableEncryptionConfiguration

func (r *TableEncryptionConfiguration) UnmarshalJSON(data []byte) error {
	var res jsonTableEncryptionConfiguration
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableEncryptionConfiguration
	} else {

		r.KmsKeyName = res.KmsKeyName

	}
	return nil
}

// This object is used to assert a desired state where this TableEncryptionConfiguration is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableEncryptionConfiguration *TableEncryptionConfiguration = &TableEncryptionConfiguration{empty: true}

func (r *TableEncryptionConfiguration) Empty() bool {
	return r.empty
}

func (r *TableEncryptionConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *TableEncryptionConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TableSnapshotDefinition struct {
	empty        bool    `json:"-"`
	Table        *string `json:"table"`
	Dataset      *string `json:"dataset"`
	Project      *string `json:"project"`
	SnapshotTime *string `json:"snapshotTime"`
}

type jsonTableSnapshotDefinition TableSnapshotDefinition

func (r *TableSnapshotDefinition) UnmarshalJSON(data []byte) error {
	var res jsonTableSnapshotDefinition
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTableSnapshotDefinition
	} else {

		r.Table = res.Table

		r.Dataset = res.Dataset

		r.Project = res.Project

		r.SnapshotTime = res.SnapshotTime

	}
	return nil
}

// This object is used to assert a desired state where this TableSnapshotDefinition is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTableSnapshotDefinition *TableSnapshotDefinition = &TableSnapshotDefinition{empty: true}

func (r *TableSnapshotDefinition) Empty() bool {
	return r.empty
}

func (r *TableSnapshotDefinition) String() string {
	return dcl.SprintResource(r)
}

func (r *TableSnapshotDefinition) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Table) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "bigquery",
		Type:    "Table",
		Version: "alpha",
	}
}

func (r *Table) ID() (string, error) {
	if err := extractTableFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"etag":                      dcl.ValueOrEmptyString(nr.Etag),
		"id":                        dcl.ValueOrEmptyString(nr.Id),
		"selfLink":                  dcl.ValueOrEmptyString(nr.SelfLink),
		"name":                      dcl.ValueOrEmptyString(nr.Name),
		"dataset":                   dcl.ValueOrEmptyString(nr.Dataset),
		"project":                   dcl.ValueOrEmptyString(nr.Project),
		"friendlyName":              dcl.ValueOrEmptyString(nr.FriendlyName),
		"description":               dcl.ValueOrEmptyString(nr.Description),
		"labels":                    dcl.ValueOrEmptyString(nr.Labels),
		"model":                     dcl.ValueOrEmptyString(nr.Model),
		"schema":                    dcl.ValueOrEmptyString(nr.Schema),
		"timePartitioning":          dcl.ValueOrEmptyString(nr.TimePartitioning),
		"rangePartitioning":         dcl.ValueOrEmptyString(nr.RangePartitioning),
		"clustering":                dcl.ValueOrEmptyString(nr.Clustering),
		"requirePartitionFilter":    dcl.ValueOrEmptyString(nr.RequirePartitionFilter),
		"numBytes":                  dcl.ValueOrEmptyString(nr.NumBytes),
		"numPhysicalBytes":          dcl.ValueOrEmptyString(nr.NumPhysicalBytes),
		"numLongTermBytes":          dcl.ValueOrEmptyString(nr.NumLongTermBytes),
		"numRows":                   dcl.ValueOrEmptyString(nr.NumRows),
		"creationTime":              dcl.ValueOrEmptyString(nr.CreationTime),
		"expirationTime":            dcl.ValueOrEmptyString(nr.ExpirationTime),
		"lastModifiedTime":          dcl.ValueOrEmptyString(nr.LastModifiedTime),
		"type":                      dcl.ValueOrEmptyString(nr.Type),
		"view":                      dcl.ValueOrEmptyString(nr.View),
		"materializedView":          dcl.ValueOrEmptyString(nr.MaterializedView),
		"externalDataConfiguration": dcl.ValueOrEmptyString(nr.ExternalDataConfiguration),
		"location":                  dcl.ValueOrEmptyString(nr.Location),
		"streamingBuffer":           dcl.ValueOrEmptyString(nr.StreamingBuffer),
		"encryptionConfiguration":   dcl.ValueOrEmptyString(nr.EncryptionConfiguration),
		"snapshotDefinition":        dcl.ValueOrEmptyString(nr.SnapshotDefinition),
		"defaultCollation":          dcl.ValueOrEmptyString(nr.DefaultCollation),
	}
	return dcl.Nprintf("projects/{{project}}/datasets/{{dataset}}/tables/{{name}}", params), nil
}

const TableMaxPage = -1

type TableList struct {
	Items []*Table

	nextToken string

	pageSize int32

	resource *Table
}

func (l *TableList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TableList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTable(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTable(ctx context.Context, project, dataset string) (*TableList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTableWithMaxResults(ctx, project, dataset, TableMaxPage)

}

func (c *Client) ListTableWithMaxResults(ctx context.Context, project, dataset string, pageSize int32) (*TableList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Table{
		Project: &project,
		Dataset: &dataset,
	}
	items, token, err := c.listTable(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TableList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetTable(ctx context.Context, r *Table) (*Table, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractTableFields(r)

	b, err := c.getTableRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTable(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Dataset = r.Dataset
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTableNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractTableFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTable(ctx context.Context, r *Table) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Table resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Table...")
	deleteOp := deleteTableOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTable deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTable(ctx context.Context, project, dataset string, filter func(*Table) bool) error {
	listObj, err := c.ListTable(ctx, project, dataset)
	if err != nil {
		return err
	}

	err = c.deleteAllTable(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTable(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTable(ctx context.Context, rawDesired *Table, opts ...dcl.ApplyOption) (*Table, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Table
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTableHelper(c, ctx, rawDesired, opts...)
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

func applyTableHelper(c *Client, ctx context.Context, rawDesired *Table, opts ...dcl.ApplyOption) (*Table, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyTable...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractTableFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.tableDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTableDiffs(c.Config, fieldDiffs, opts)
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
	var ops []tableApiOperation
	if create {
		ops = append(ops, &createTableOperation{})
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
	return applyTableDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyTableDiff(c *Client, ctx context.Context, desired *Table, rawDesired *Table, ops []tableApiOperation, opts ...dcl.ApplyOption) (*Table, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetTable(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTableOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTable(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTableNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTableNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTableDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractTableFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractTableFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTable(c, newDesired, newState)
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
