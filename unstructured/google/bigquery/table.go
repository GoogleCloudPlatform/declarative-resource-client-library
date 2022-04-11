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
package bigquery

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Table struct{}

func TableToUnstructured(r *dclService.Table) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "bigquery",
			Version: "ga",
			Type:    "Table",
		},
		Object: make(map[string]interface{}),
	}
	if r.Clustering != nil && r.Clustering != dclService.EmptyTableClustering {
		rClustering := make(map[string]interface{})
		var rClusteringFields []interface{}
		for _, rClusteringFieldsVal := range r.Clustering.Fields {
			rClusteringFields = append(rClusteringFields, rClusteringFieldsVal)
		}
		rClustering["fields"] = rClusteringFields
		u.Object["clustering"] = rClustering
	}
	if r.CreationTime != nil {
		u.Object["creationTime"] = *r.CreationTime
	}
	if r.Dataset != nil {
		u.Object["dataset"] = *r.Dataset
	}
	if r.DefaultCollation != nil {
		u.Object["defaultCollation"] = *r.DefaultCollation
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.EncryptionConfiguration != nil && r.EncryptionConfiguration != dclService.EmptyTableEncryptionConfiguration {
		rEncryptionConfiguration := make(map[string]interface{})
		if r.EncryptionConfiguration.KmsKeyName != nil {
			rEncryptionConfiguration["kmsKeyName"] = *r.EncryptionConfiguration.KmsKeyName
		}
		u.Object["encryptionConfiguration"] = rEncryptionConfiguration
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.ExpirationTime != nil {
		u.Object["expirationTime"] = *r.ExpirationTime
	}
	if r.ExternalDataConfiguration != nil && r.ExternalDataConfiguration != dclService.EmptyTableExternalDataConfiguration {
		rExternalDataConfiguration := make(map[string]interface{})
		if r.ExternalDataConfiguration.Autodetect != nil {
			rExternalDataConfiguration["autodetect"] = *r.ExternalDataConfiguration.Autodetect
		}
		if r.ExternalDataConfiguration.AvroOptions != nil && r.ExternalDataConfiguration.AvroOptions != dclService.EmptyTableExternalDataConfigurationAvroOptions {
			rExternalDataConfigurationAvroOptions := make(map[string]interface{})
			if r.ExternalDataConfiguration.AvroOptions.UseAvroLogicalTypes != nil {
				rExternalDataConfigurationAvroOptions["useAvroLogicalTypes"] = *r.ExternalDataConfiguration.AvroOptions.UseAvroLogicalTypes
			}
			rExternalDataConfiguration["avroOptions"] = rExternalDataConfigurationAvroOptions
		}
		if r.ExternalDataConfiguration.BigtableOptions != nil && r.ExternalDataConfiguration.BigtableOptions != dclService.EmptyTableExternalDataConfigurationBigtableOptions {
			rExternalDataConfigurationBigtableOptions := make(map[string]interface{})
			var rExternalDataConfigurationBigtableOptionsColumnFamilies []interface{}
			for _, rExternalDataConfigurationBigtableOptionsColumnFamiliesVal := range r.ExternalDataConfiguration.BigtableOptions.ColumnFamilies {
				rExternalDataConfigurationBigtableOptionsColumnFamiliesObject := make(map[string]interface{})
				var rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumns []interface{}
				for _, rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal := range rExternalDataConfigurationBigtableOptionsColumnFamiliesVal.Columns {
					rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsObject := make(map[string]interface{})
					if rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.Encoding != nil {
						rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsObject["encoding"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.Encoding
					}
					if rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.FieldName != nil {
						rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsObject["fieldName"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.FieldName
					}
					if rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.OnlyReadLatest != nil {
						rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsObject["onlyReadLatest"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.OnlyReadLatest
					}
					if rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.QualifierEncoded != nil {
						rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsObject["qualifierEncoded"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.QualifierEncoded
					}
					if rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.QualifierString != nil {
						rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsObject["qualifierString"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.QualifierString
					}
					if rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.Type != nil {
						rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsObject["type"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsVal.Type
					}
					rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumns = append(rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumns, rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumnsObject)
				}
				rExternalDataConfigurationBigtableOptionsColumnFamiliesObject["columns"] = rExternalDataConfigurationBigtableOptionsColumnFamiliesValColumns
				if rExternalDataConfigurationBigtableOptionsColumnFamiliesVal.Encoding != nil {
					rExternalDataConfigurationBigtableOptionsColumnFamiliesObject["encoding"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesVal.Encoding
				}
				if rExternalDataConfigurationBigtableOptionsColumnFamiliesVal.FamilyId != nil {
					rExternalDataConfigurationBigtableOptionsColumnFamiliesObject["familyId"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesVal.FamilyId
				}
				if rExternalDataConfigurationBigtableOptionsColumnFamiliesVal.OnlyReadLatest != nil {
					rExternalDataConfigurationBigtableOptionsColumnFamiliesObject["onlyReadLatest"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesVal.OnlyReadLatest
				}
				if rExternalDataConfigurationBigtableOptionsColumnFamiliesVal.Type != nil {
					rExternalDataConfigurationBigtableOptionsColumnFamiliesObject["type"] = *rExternalDataConfigurationBigtableOptionsColumnFamiliesVal.Type
				}
				rExternalDataConfigurationBigtableOptionsColumnFamilies = append(rExternalDataConfigurationBigtableOptionsColumnFamilies, rExternalDataConfigurationBigtableOptionsColumnFamiliesObject)
			}
			rExternalDataConfigurationBigtableOptions["columnFamilies"] = rExternalDataConfigurationBigtableOptionsColumnFamilies
			if r.ExternalDataConfiguration.BigtableOptions.IgnoreUnspecifiedColumnFamilies != nil {
				rExternalDataConfigurationBigtableOptions["ignoreUnspecifiedColumnFamilies"] = *r.ExternalDataConfiguration.BigtableOptions.IgnoreUnspecifiedColumnFamilies
			}
			if r.ExternalDataConfiguration.BigtableOptions.ReadRowkeyAsString != nil {
				rExternalDataConfigurationBigtableOptions["readRowkeyAsString"] = *r.ExternalDataConfiguration.BigtableOptions.ReadRowkeyAsString
			}
			rExternalDataConfiguration["bigtableOptions"] = rExternalDataConfigurationBigtableOptions
		}
		if r.ExternalDataConfiguration.Compression != nil {
			rExternalDataConfiguration["compression"] = *r.ExternalDataConfiguration.Compression
		}
		if r.ExternalDataConfiguration.ConnectionId != nil {
			rExternalDataConfiguration["connectionId"] = *r.ExternalDataConfiguration.ConnectionId
		}
		if r.ExternalDataConfiguration.CsvOptions != nil && r.ExternalDataConfiguration.CsvOptions != dclService.EmptyTableExternalDataConfigurationCsvOptions {
			rExternalDataConfigurationCsvOptions := make(map[string]interface{})
			if r.ExternalDataConfiguration.CsvOptions.AllowJaggedRows != nil {
				rExternalDataConfigurationCsvOptions["allowJaggedRows"] = *r.ExternalDataConfiguration.CsvOptions.AllowJaggedRows
			}
			if r.ExternalDataConfiguration.CsvOptions.AllowQuotedNewlines != nil {
				rExternalDataConfigurationCsvOptions["allowQuotedNewlines"] = *r.ExternalDataConfiguration.CsvOptions.AllowQuotedNewlines
			}
			if r.ExternalDataConfiguration.CsvOptions.Encoding != nil {
				rExternalDataConfigurationCsvOptions["encoding"] = *r.ExternalDataConfiguration.CsvOptions.Encoding
			}
			if r.ExternalDataConfiguration.CsvOptions.FieldDelimiter != nil {
				rExternalDataConfigurationCsvOptions["fieldDelimiter"] = *r.ExternalDataConfiguration.CsvOptions.FieldDelimiter
			}
			if r.ExternalDataConfiguration.CsvOptions.Quote != nil {
				rExternalDataConfigurationCsvOptions["quote"] = *r.ExternalDataConfiguration.CsvOptions.Quote
			}
			if r.ExternalDataConfiguration.CsvOptions.SkipLeadingRows != nil {
				rExternalDataConfigurationCsvOptions["skipLeadingRows"] = *r.ExternalDataConfiguration.CsvOptions.SkipLeadingRows
			}
			rExternalDataConfiguration["csvOptions"] = rExternalDataConfigurationCsvOptions
		}
		var rExternalDataConfigurationDecimalTargetTypes []interface{}
		for _, rExternalDataConfigurationDecimalTargetTypesVal := range r.ExternalDataConfiguration.DecimalTargetTypes {
			rExternalDataConfigurationDecimalTargetTypes = append(rExternalDataConfigurationDecimalTargetTypes, string(rExternalDataConfigurationDecimalTargetTypesVal))
		}
		rExternalDataConfiguration["decimalTargetTypes"] = rExternalDataConfigurationDecimalTargetTypes
		if r.ExternalDataConfiguration.GoogleSheetsOptions != nil && r.ExternalDataConfiguration.GoogleSheetsOptions != dclService.EmptyTableExternalDataConfigurationGoogleSheetsOptions {
			rExternalDataConfigurationGoogleSheetsOptions := make(map[string]interface{})
			if r.ExternalDataConfiguration.GoogleSheetsOptions.Range != nil {
				rExternalDataConfigurationGoogleSheetsOptions["range"] = *r.ExternalDataConfiguration.GoogleSheetsOptions.Range
			}
			if r.ExternalDataConfiguration.GoogleSheetsOptions.SkipLeadingRows != nil {
				rExternalDataConfigurationGoogleSheetsOptions["skipLeadingRows"] = *r.ExternalDataConfiguration.GoogleSheetsOptions.SkipLeadingRows
			}
			rExternalDataConfiguration["googleSheetsOptions"] = rExternalDataConfigurationGoogleSheetsOptions
		}
		if r.ExternalDataConfiguration.HivePartitioningOptions != nil && r.ExternalDataConfiguration.HivePartitioningOptions != dclService.EmptyTableExternalDataConfigurationHivePartitioningOptions {
			rExternalDataConfigurationHivePartitioningOptions := make(map[string]interface{})
			var rExternalDataConfigurationHivePartitioningOptionsFields []interface{}
			for _, rExternalDataConfigurationHivePartitioningOptionsFieldsVal := range r.ExternalDataConfiguration.HivePartitioningOptions.Fields {
				rExternalDataConfigurationHivePartitioningOptionsFields = append(rExternalDataConfigurationHivePartitioningOptionsFields, rExternalDataConfigurationHivePartitioningOptionsFieldsVal)
			}
			rExternalDataConfigurationHivePartitioningOptions["fields"] = rExternalDataConfigurationHivePartitioningOptionsFields
			if r.ExternalDataConfiguration.HivePartitioningOptions.Mode != nil {
				rExternalDataConfigurationHivePartitioningOptions["mode"] = *r.ExternalDataConfiguration.HivePartitioningOptions.Mode
			}
			if r.ExternalDataConfiguration.HivePartitioningOptions.RequirePartitionFilter != nil {
				rExternalDataConfigurationHivePartitioningOptions["requirePartitionFilter"] = *r.ExternalDataConfiguration.HivePartitioningOptions.RequirePartitionFilter
			}
			if r.ExternalDataConfiguration.HivePartitioningOptions.SourceUriPrefix != nil {
				rExternalDataConfigurationHivePartitioningOptions["sourceUriPrefix"] = *r.ExternalDataConfiguration.HivePartitioningOptions.SourceUriPrefix
			}
			rExternalDataConfiguration["hivePartitioningOptions"] = rExternalDataConfigurationHivePartitioningOptions
		}
		if r.ExternalDataConfiguration.IgnoreUnknownValues != nil {
			rExternalDataConfiguration["ignoreUnknownValues"] = *r.ExternalDataConfiguration.IgnoreUnknownValues
		}
		if r.ExternalDataConfiguration.JsonExtension != nil {
			rExternalDataConfiguration["jsonExtension"] = string(*r.ExternalDataConfiguration.JsonExtension)
		}
		if r.ExternalDataConfiguration.MaxBadRecords != nil {
			rExternalDataConfiguration["maxBadRecords"] = *r.ExternalDataConfiguration.MaxBadRecords
		}
		var rExternalDataConfigurationMaxBadRecordsAlternative []interface{}
		for _, rExternalDataConfigurationMaxBadRecordsAlternativeVal := range r.ExternalDataConfiguration.MaxBadRecordsAlternative {
			rExternalDataConfigurationMaxBadRecordsAlternative = append(rExternalDataConfigurationMaxBadRecordsAlternative, rExternalDataConfigurationMaxBadRecordsAlternativeVal)
		}
		rExternalDataConfiguration["maxBadRecordsAlternative"] = rExternalDataConfigurationMaxBadRecordsAlternative
		if r.ExternalDataConfiguration.ParquetOptions != nil && r.ExternalDataConfiguration.ParquetOptions != dclService.EmptyTableExternalDataConfigurationParquetOptions {
			rExternalDataConfigurationParquetOptions := make(map[string]interface{})
			if r.ExternalDataConfiguration.ParquetOptions.EnableListInference != nil {
				rExternalDataConfigurationParquetOptions["enableListInference"] = *r.ExternalDataConfiguration.ParquetOptions.EnableListInference
			}
			if r.ExternalDataConfiguration.ParquetOptions.EnumAsString != nil {
				rExternalDataConfigurationParquetOptions["enumAsString"] = *r.ExternalDataConfiguration.ParquetOptions.EnumAsString
			}
			rExternalDataConfiguration["parquetOptions"] = rExternalDataConfigurationParquetOptions
		}
		if r.ExternalDataConfiguration.Schema != nil && r.ExternalDataConfiguration.Schema != dclService.EmptyTableExternalDataConfigurationSchema {
			rExternalDataConfigurationSchema := make(map[string]interface{})
			var rExternalDataConfigurationSchemaFields []interface{}
			for _, rExternalDataConfigurationSchemaFieldsVal := range r.ExternalDataConfiguration.Schema.Fields {
				rExternalDataConfigurationSchemaFields = append(rExternalDataConfigurationSchemaFields, TableGooglecloudbigqueryv2TablefieldschemaToUnstructured(&rExternalDataConfigurationSchemaFieldsVal))
			}
			rExternalDataConfigurationSchema["fields"] = rExternalDataConfigurationSchemaFields
			rExternalDataConfiguration["schema"] = rExternalDataConfigurationSchema
		}
		if r.ExternalDataConfiguration.SourceFormat != nil {
			rExternalDataConfiguration["sourceFormat"] = *r.ExternalDataConfiguration.SourceFormat
		}
		var rExternalDataConfigurationSourceUris []interface{}
		for _, rExternalDataConfigurationSourceUrisVal := range r.ExternalDataConfiguration.SourceUris {
			rExternalDataConfigurationSourceUris = append(rExternalDataConfigurationSourceUris, rExternalDataConfigurationSourceUrisVal)
		}
		rExternalDataConfiguration["sourceUris"] = rExternalDataConfigurationSourceUris
		if r.ExternalDataConfiguration.ValueConversionModes != nil && r.ExternalDataConfiguration.ValueConversionModes != dclService.EmptyTableExternalDataConfigurationValueConversionModes {
			rExternalDataConfigurationValueConversionModes := make(map[string]interface{})
			if r.ExternalDataConfiguration.ValueConversionModes.NumericTypeOutOfRangeConversionMode != nil {
				rExternalDataConfigurationValueConversionModes["numericTypeOutOfRangeConversionMode"] = string(*r.ExternalDataConfiguration.ValueConversionModes.NumericTypeOutOfRangeConversionMode)
			}
			if r.ExternalDataConfiguration.ValueConversionModes.TemporalTypesOutOfRangeConversionMode != nil {
				rExternalDataConfigurationValueConversionModes["temporalTypesOutOfRangeConversionMode"] = string(*r.ExternalDataConfiguration.ValueConversionModes.TemporalTypesOutOfRangeConversionMode)
			}
			rExternalDataConfiguration["valueConversionModes"] = rExternalDataConfigurationValueConversionModes
		}
		u.Object["externalDataConfiguration"] = rExternalDataConfiguration
	}
	if r.FriendlyName != nil {
		u.Object["friendlyName"] = *r.FriendlyName
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.LastModifiedTime != nil {
		u.Object["lastModifiedTime"] = *r.LastModifiedTime
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.MaterializedView != nil && r.MaterializedView != dclService.EmptyTableMaterializedView {
		rMaterializedView := make(map[string]interface{})
		if r.MaterializedView.EnableRefresh != nil {
			rMaterializedView["enableRefresh"] = *r.MaterializedView.EnableRefresh
		}
		if r.MaterializedView.LastRefreshTime != nil {
			rMaterializedView["lastRefreshTime"] = *r.MaterializedView.LastRefreshTime
		}
		if r.MaterializedView.Query != nil {
			rMaterializedView["query"] = *r.MaterializedView.Query
		}
		if r.MaterializedView.RefreshIntervalMs != nil {
			rMaterializedView["refreshIntervalMs"] = *r.MaterializedView.RefreshIntervalMs
		}
		u.Object["materializedView"] = rMaterializedView
	}
	if r.Model != nil && r.Model != dclService.EmptyTableModel {
		rModel := make(map[string]interface{})
		if r.Model.ModelOptions != nil && r.Model.ModelOptions != dclService.EmptyTableModelModelOptions {
			rModelModelOptions := make(map[string]interface{})
			var rModelModelOptionsLabels []interface{}
			for _, rModelModelOptionsLabelsVal := range r.Model.ModelOptions.Labels {
				rModelModelOptionsLabels = append(rModelModelOptionsLabels, rModelModelOptionsLabelsVal)
			}
			rModelModelOptions["labels"] = rModelModelOptionsLabels
			if r.Model.ModelOptions.LossType != nil {
				rModelModelOptions["lossType"] = *r.Model.ModelOptions.LossType
			}
			if r.Model.ModelOptions.ModelType != nil {
				rModelModelOptions["modelType"] = *r.Model.ModelOptions.ModelType
			}
			rModel["modelOptions"] = rModelModelOptions
		}
		var rModelTrainingRuns []interface{}
		for _, rModelTrainingRunsVal := range r.Model.TrainingRuns {
			rModelTrainingRunsObject := make(map[string]interface{})
			var rModelTrainingRunsValIterationResults []interface{}
			for _, rModelTrainingRunsValIterationResultsVal := range rModelTrainingRunsVal.IterationResults {
				rModelTrainingRunsValIterationResultsObject := make(map[string]interface{})
				if rModelTrainingRunsValIterationResultsVal.DurationMs != nil {
					rModelTrainingRunsValIterationResultsObject["durationMs"] = *rModelTrainingRunsValIterationResultsVal.DurationMs
				}
				if rModelTrainingRunsValIterationResultsVal.EvalLoss != nil {
					rModelTrainingRunsValIterationResultsObject["evalLoss"] = *rModelTrainingRunsValIterationResultsVal.EvalLoss
				}
				if rModelTrainingRunsValIterationResultsVal.Index != nil {
					rModelTrainingRunsValIterationResultsObject["index"] = *rModelTrainingRunsValIterationResultsVal.Index
				}
				if rModelTrainingRunsValIterationResultsVal.LearnRate != nil {
					rModelTrainingRunsValIterationResultsObject["learnRate"] = *rModelTrainingRunsValIterationResultsVal.LearnRate
				}
				if rModelTrainingRunsValIterationResultsVal.TrainingLoss != nil {
					rModelTrainingRunsValIterationResultsObject["trainingLoss"] = *rModelTrainingRunsValIterationResultsVal.TrainingLoss
				}
				rModelTrainingRunsValIterationResults = append(rModelTrainingRunsValIterationResults, rModelTrainingRunsValIterationResultsObject)
			}
			rModelTrainingRunsObject["iterationResults"] = rModelTrainingRunsValIterationResults
			if rModelTrainingRunsVal.StartTime != nil {
				rModelTrainingRunsObject["startTime"] = *rModelTrainingRunsVal.StartTime
			}
			if rModelTrainingRunsVal.State != nil {
				rModelTrainingRunsObject["state"] = *rModelTrainingRunsVal.State
			}
			if rModelTrainingRunsVal.TrainingOptions != nil && rModelTrainingRunsVal.TrainingOptions != dclService.EmptyTableModelTrainingRunsTrainingOptions {
				rModelTrainingRunsValTrainingOptions := make(map[string]interface{})
				if rModelTrainingRunsVal.TrainingOptions.EarlyStop != nil {
					rModelTrainingRunsValTrainingOptions["earlyStop"] = *rModelTrainingRunsVal.TrainingOptions.EarlyStop
				}
				if rModelTrainingRunsVal.TrainingOptions.L1Reg != nil {
					rModelTrainingRunsValTrainingOptions["l1Reg"] = *rModelTrainingRunsVal.TrainingOptions.L1Reg
				}
				if rModelTrainingRunsVal.TrainingOptions.L2Reg != nil {
					rModelTrainingRunsValTrainingOptions["l2Reg"] = *rModelTrainingRunsVal.TrainingOptions.L2Reg
				}
				if rModelTrainingRunsVal.TrainingOptions.LearnRate != nil {
					rModelTrainingRunsValTrainingOptions["learnRate"] = *rModelTrainingRunsVal.TrainingOptions.LearnRate
				}
				if rModelTrainingRunsVal.TrainingOptions.LearnRateStrategy != nil {
					rModelTrainingRunsValTrainingOptions["learnRateStrategy"] = *rModelTrainingRunsVal.TrainingOptions.LearnRateStrategy
				}
				if rModelTrainingRunsVal.TrainingOptions.LineSearchInitLearnRate != nil {
					rModelTrainingRunsValTrainingOptions["lineSearchInitLearnRate"] = *rModelTrainingRunsVal.TrainingOptions.LineSearchInitLearnRate
				}
				if rModelTrainingRunsVal.TrainingOptions.MaxIteration != nil {
					rModelTrainingRunsValTrainingOptions["maxIteration"] = *rModelTrainingRunsVal.TrainingOptions.MaxIteration
				}
				if rModelTrainingRunsVal.TrainingOptions.MinRelProgress != nil {
					rModelTrainingRunsValTrainingOptions["minRelProgress"] = *rModelTrainingRunsVal.TrainingOptions.MinRelProgress
				}
				if rModelTrainingRunsVal.TrainingOptions.WarmStart != nil {
					rModelTrainingRunsValTrainingOptions["warmStart"] = *rModelTrainingRunsVal.TrainingOptions.WarmStart
				}
				rModelTrainingRunsObject["trainingOptions"] = rModelTrainingRunsValTrainingOptions
			}
			rModelTrainingRuns = append(rModelTrainingRuns, rModelTrainingRunsObject)
		}
		rModel["trainingRuns"] = rModelTrainingRuns
		u.Object["model"] = rModel
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.NumBytes != nil {
		u.Object["numBytes"] = *r.NumBytes
	}
	if r.NumLongTermBytes != nil {
		u.Object["numLongTermBytes"] = *r.NumLongTermBytes
	}
	if r.NumPhysicalBytes != nil {
		u.Object["numPhysicalBytes"] = *r.NumPhysicalBytes
	}
	if r.NumRows != nil {
		u.Object["numRows"] = *r.NumRows
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RangePartitioning != nil && r.RangePartitioning != dclService.EmptyTableRangePartitioning {
		rRangePartitioning := make(map[string]interface{})
		if r.RangePartitioning.Field != nil {
			rRangePartitioning["field"] = *r.RangePartitioning.Field
		}
		if r.RangePartitioning.Range != nil && r.RangePartitioning.Range != dclService.EmptyTableRangePartitioningRange {
			rRangePartitioningRange := make(map[string]interface{})
			if r.RangePartitioning.Range.End != nil {
				rRangePartitioningRange["end"] = *r.RangePartitioning.Range.End
			}
			if r.RangePartitioning.Range.Interval != nil {
				rRangePartitioningRange["interval"] = *r.RangePartitioning.Range.Interval
			}
			if r.RangePartitioning.Range.Start != nil {
				rRangePartitioningRange["start"] = *r.RangePartitioning.Range.Start
			}
			rRangePartitioning["range"] = rRangePartitioningRange
		}
		u.Object["rangePartitioning"] = rRangePartitioning
	}
	if r.RequirePartitionFilter != nil {
		u.Object["requirePartitionFilter"] = *r.RequirePartitionFilter
	}
	if r.Schema != nil && r.Schema != dclService.EmptyTableSchema {
		rSchema := make(map[string]interface{})
		var rSchemaFields []interface{}
		for _, rSchemaFieldsVal := range r.Schema.Fields {
			rSchemaFieldsObject := make(map[string]interface{})
			if rSchemaFieldsVal.Categories != nil && rSchemaFieldsVal.Categories != dclService.EmptyTableGooglecloudbigqueryv2TablefieldschemaCategories {
				rSchemaFieldsValCategories := make(map[string]interface{})
				var rSchemaFieldsValCategoriesNames []interface{}
				for _, rSchemaFieldsValCategoriesNamesVal := range rSchemaFieldsVal.Categories.Names {
					rSchemaFieldsValCategoriesNames = append(rSchemaFieldsValCategoriesNames, rSchemaFieldsValCategoriesNamesVal)
				}
				rSchemaFieldsValCategories["names"] = rSchemaFieldsValCategoriesNames
				rSchemaFieldsObject["categories"] = rSchemaFieldsValCategories
			}
			if rSchemaFieldsVal.Collation != nil {
				rSchemaFieldsObject["collation"] = *rSchemaFieldsVal.Collation
			}
			if rSchemaFieldsVal.DefaultValueExpression != nil {
				rSchemaFieldsObject["defaultValueExpression"] = *rSchemaFieldsVal.DefaultValueExpression
			}
			if rSchemaFieldsVal.Description != nil {
				rSchemaFieldsObject["description"] = *rSchemaFieldsVal.Description
			}
			var rSchemaFieldsValFields []interface{}
			for _, rSchemaFieldsValFieldsVal := range rSchemaFieldsVal.Fields {
				rSchemaFieldsValFields = append(rSchemaFieldsValFields, TableGooglecloudbigqueryv2TablefieldschemaToUnstructured(&rSchemaFieldsValFieldsVal))
			}
			rSchemaFieldsObject["fields"] = rSchemaFieldsValFields
			if rSchemaFieldsVal.MaxLength != nil {
				rSchemaFieldsObject["maxLength"] = *rSchemaFieldsVal.MaxLength
			}
			if rSchemaFieldsVal.Mode != nil {
				rSchemaFieldsObject["mode"] = *rSchemaFieldsVal.Mode
			}
			if rSchemaFieldsVal.Name != nil {
				rSchemaFieldsObject["name"] = *rSchemaFieldsVal.Name
			}
			var rSchemaFieldsValNameAlternative []interface{}
			for _, rSchemaFieldsValNameAlternativeVal := range rSchemaFieldsVal.NameAlternative {
				rSchemaFieldsValNameAlternative = append(rSchemaFieldsValNameAlternative, rSchemaFieldsValNameAlternativeVal)
			}
			rSchemaFieldsObject["nameAlternative"] = rSchemaFieldsValNameAlternative
			if rSchemaFieldsVal.PolicyTags != nil && rSchemaFieldsVal.PolicyTags != dclService.EmptyTableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
				rSchemaFieldsValPolicyTags := make(map[string]interface{})
				var rSchemaFieldsValPolicyTagsNames []interface{}
				for _, rSchemaFieldsValPolicyTagsNamesVal := range rSchemaFieldsVal.PolicyTags.Names {
					rSchemaFieldsValPolicyTagsNames = append(rSchemaFieldsValPolicyTagsNames, rSchemaFieldsValPolicyTagsNamesVal)
				}
				rSchemaFieldsValPolicyTags["names"] = rSchemaFieldsValPolicyTagsNames
				rSchemaFieldsObject["policyTags"] = rSchemaFieldsValPolicyTags
			}
			if rSchemaFieldsVal.Precision != nil {
				rSchemaFieldsObject["precision"] = *rSchemaFieldsVal.Precision
			}
			if rSchemaFieldsVal.Scale != nil {
				rSchemaFieldsObject["scale"] = *rSchemaFieldsVal.Scale
			}
			if rSchemaFieldsVal.Type != nil {
				rSchemaFieldsObject["type"] = *rSchemaFieldsVal.Type
			}
			rSchemaFields = append(rSchemaFields, rSchemaFieldsObject)
		}
		rSchema["fields"] = rSchemaFields
		u.Object["schema"] = rSchema
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.SnapshotDefinition != nil && r.SnapshotDefinition != dclService.EmptyTableSnapshotDefinition {
		rSnapshotDefinition := make(map[string]interface{})
		if r.SnapshotDefinition.Dataset != nil {
			rSnapshotDefinition["dataset"] = *r.SnapshotDefinition.Dataset
		}
		if r.SnapshotDefinition.Project != nil {
			rSnapshotDefinition["project"] = *r.SnapshotDefinition.Project
		}
		if r.SnapshotDefinition.SnapshotTime != nil {
			rSnapshotDefinition["snapshotTime"] = *r.SnapshotDefinition.SnapshotTime
		}
		if r.SnapshotDefinition.Table != nil {
			rSnapshotDefinition["table"] = *r.SnapshotDefinition.Table
		}
		u.Object["snapshotDefinition"] = rSnapshotDefinition
	}
	if r.StreamingBuffer != nil && r.StreamingBuffer != dclService.EmptyTableStreamingBuffer {
		rStreamingBuffer := make(map[string]interface{})
		if r.StreamingBuffer.EstimatedBytes != nil {
			rStreamingBuffer["estimatedBytes"] = *r.StreamingBuffer.EstimatedBytes
		}
		if r.StreamingBuffer.EstimatedRows != nil {
			rStreamingBuffer["estimatedRows"] = *r.StreamingBuffer.EstimatedRows
		}
		if r.StreamingBuffer.OldestEntryTime != nil {
			rStreamingBuffer["oldestEntryTime"] = *r.StreamingBuffer.OldestEntryTime
		}
		u.Object["streamingBuffer"] = rStreamingBuffer
	}
	if r.TimePartitioning != nil && r.TimePartitioning != dclService.EmptyTableTimePartitioning {
		rTimePartitioning := make(map[string]interface{})
		if r.TimePartitioning.ExpirationMs != nil {
			rTimePartitioning["expirationMs"] = *r.TimePartitioning.ExpirationMs
		}
		if r.TimePartitioning.Field != nil {
			rTimePartitioning["field"] = *r.TimePartitioning.Field
		}
		if r.TimePartitioning.Type != nil {
			rTimePartitioning["type"] = *r.TimePartitioning.Type
		}
		u.Object["timePartitioning"] = rTimePartitioning
	}
	if r.Type != nil {
		u.Object["type"] = *r.Type
	}
	if r.View != nil && r.View != dclService.EmptyTableView {
		rView := make(map[string]interface{})
		if r.View.Query != nil {
			rView["query"] = *r.View.Query
		}
		if r.View.UseExplicitColumnNames != nil {
			rView["useExplicitColumnNames"] = *r.View.UseExplicitColumnNames
		}
		if r.View.UseLegacySql != nil {
			rView["useLegacySql"] = *r.View.UseLegacySql
		}
		var rViewUserDefinedFunctionResources []interface{}
		for _, rViewUserDefinedFunctionResourcesVal := range r.View.UserDefinedFunctionResources {
			rViewUserDefinedFunctionResourcesObject := make(map[string]interface{})
			if rViewUserDefinedFunctionResourcesVal.InlineCode != nil {
				rViewUserDefinedFunctionResourcesObject["inlineCode"] = *rViewUserDefinedFunctionResourcesVal.InlineCode
			}
			var rViewUserDefinedFunctionResourcesValInlineCodeAlternative []interface{}
			for _, rViewUserDefinedFunctionResourcesValInlineCodeAlternativeVal := range rViewUserDefinedFunctionResourcesVal.InlineCodeAlternative {
				rViewUserDefinedFunctionResourcesValInlineCodeAlternative = append(rViewUserDefinedFunctionResourcesValInlineCodeAlternative, rViewUserDefinedFunctionResourcesValInlineCodeAlternativeVal)
			}
			rViewUserDefinedFunctionResourcesObject["inlineCodeAlternative"] = rViewUserDefinedFunctionResourcesValInlineCodeAlternative
			if rViewUserDefinedFunctionResourcesVal.ResourceUri != nil {
				rViewUserDefinedFunctionResourcesObject["resourceUri"] = *rViewUserDefinedFunctionResourcesVal.ResourceUri
			}
			rViewUserDefinedFunctionResources = append(rViewUserDefinedFunctionResources, rViewUserDefinedFunctionResourcesObject)
		}
		rView["userDefinedFunctionResources"] = rViewUserDefinedFunctionResources
		u.Object["view"] = rView
	}
	return u
}

func TableGooglecloudbigqueryv2TablefieldschemaToUnstructured(r *dclService.TableGooglecloudbigqueryv2Tablefieldschema) map[string]interface{} {
	result := make(map[string]interface{})
	if r.Categories != nil && r.Categories != dclService.EmptyTableGooglecloudbigqueryv2TablefieldschemaCategories {
		rCategories := make(map[string]interface{})
		var rCategoriesNames []interface{}
		for _, rCategoriesNamesVal := range r.Categories.Names {
			rCategoriesNames = append(rCategoriesNames, rCategoriesNamesVal)
		}
		rCategories["names"] = rCategoriesNames
		result["categories"] = rCategories
	}
	if r.Collation != nil {
		result["collation"] = *r.Collation
	}
	if r.DefaultValueExpression != nil {
		result["defaultValueExpression"] = *r.DefaultValueExpression
	}
	if r.Description != nil {
		result["description"] = *r.Description
	}
	var rFields []interface{}
	for _, rFieldsVal := range r.Fields {
		rFields = append(rFields, TableGooglecloudbigqueryv2TablefieldschemaToUnstructured(&rFieldsVal))
	}
	result["fields"] = rFields
	if r.MaxLength != nil {
		result["maxLength"] = *r.MaxLength
	}
	if r.Mode != nil {
		result["mode"] = *r.Mode
	}
	if r.Name != nil {
		result["name"] = *r.Name
	}
	var rNameAlternative []interface{}
	for _, rNameAlternativeVal := range r.NameAlternative {
		rNameAlternative = append(rNameAlternative, rNameAlternativeVal)
	}
	result["nameAlternative"] = rNameAlternative
	if r.PolicyTags != nil && r.PolicyTags != dclService.EmptyTableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
		rPolicyTags := make(map[string]interface{})
		var rPolicyTagsNames []interface{}
		for _, rPolicyTagsNamesVal := range r.PolicyTags.Names {
			rPolicyTagsNames = append(rPolicyTagsNames, rPolicyTagsNamesVal)
		}
		rPolicyTags["names"] = rPolicyTagsNames
		result["policyTags"] = rPolicyTags
	}
	if r.Precision != nil {
		result["precision"] = *r.Precision
	}
	if r.Scale != nil {
		result["scale"] = *r.Scale
	}
	if r.Type != nil {
		result["type"] = *r.Type
	}
	return result
}

func UnstructuredToTable(u *unstructured.Resource) (*dclService.Table, error) {
	r := &dclService.Table{}
	if _, ok := u.Object["clustering"]; ok {
		if rClustering, ok := u.Object["clustering"].(map[string]interface{}); ok {
			r.Clustering = &dclService.TableClustering{}
			if _, ok := rClustering["fields"]; ok {
				if s, ok := rClustering["fields"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Clustering.Fields = append(r.Clustering.Fields, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Clustering.Fields: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Clustering: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["creationTime"]; ok {
		if i, ok := u.Object["creationTime"].(int64); ok {
			r.CreationTime = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.CreationTime: expected int64")
		}
	}
	if _, ok := u.Object["dataset"]; ok {
		if s, ok := u.Object["dataset"].(string); ok {
			r.Dataset = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Dataset: expected string")
		}
	}
	if _, ok := u.Object["defaultCollation"]; ok {
		if s, ok := u.Object["defaultCollation"].(string); ok {
			r.DefaultCollation = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DefaultCollation: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["encryptionConfiguration"]; ok {
		if rEncryptionConfiguration, ok := u.Object["encryptionConfiguration"].(map[string]interface{}); ok {
			r.EncryptionConfiguration = &dclService.TableEncryptionConfiguration{}
			if _, ok := rEncryptionConfiguration["kmsKeyName"]; ok {
				if s, ok := rEncryptionConfiguration["kmsKeyName"].(string); ok {
					r.EncryptionConfiguration.KmsKeyName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.EncryptionConfiguration.KmsKeyName: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.EncryptionConfiguration: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["expirationTime"]; ok {
		if i, ok := u.Object["expirationTime"].(int64); ok {
			r.ExpirationTime = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.ExpirationTime: expected int64")
		}
	}
	if _, ok := u.Object["externalDataConfiguration"]; ok {
		if rExternalDataConfiguration, ok := u.Object["externalDataConfiguration"].(map[string]interface{}); ok {
			r.ExternalDataConfiguration = &dclService.TableExternalDataConfiguration{}
			if _, ok := rExternalDataConfiguration["autodetect"]; ok {
				if b, ok := rExternalDataConfiguration["autodetect"].(bool); ok {
					r.ExternalDataConfiguration.Autodetect = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.Autodetect: expected bool")
				}
			}
			if _, ok := rExternalDataConfiguration["avroOptions"]; ok {
				if rExternalDataConfigurationAvroOptions, ok := rExternalDataConfiguration["avroOptions"].(map[string]interface{}); ok {
					r.ExternalDataConfiguration.AvroOptions = &dclService.TableExternalDataConfigurationAvroOptions{}
					if _, ok := rExternalDataConfigurationAvroOptions["useAvroLogicalTypes"]; ok {
						if b, ok := rExternalDataConfigurationAvroOptions["useAvroLogicalTypes"].(bool); ok {
							r.ExternalDataConfiguration.AvroOptions.UseAvroLogicalTypes = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.AvroOptions.UseAvroLogicalTypes: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.AvroOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["bigtableOptions"]; ok {
				if rExternalDataConfigurationBigtableOptions, ok := rExternalDataConfiguration["bigtableOptions"].(map[string]interface{}); ok {
					r.ExternalDataConfiguration.BigtableOptions = &dclService.TableExternalDataConfigurationBigtableOptions{}
					if _, ok := rExternalDataConfigurationBigtableOptions["columnFamilies"]; ok {
						if s, ok := rExternalDataConfigurationBigtableOptions["columnFamilies"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rExternalDataConfigurationBigtableOptionsColumnFamilies dclService.TableExternalDataConfigurationBigtableOptionsColumnFamilies
									if _, ok := objval["columns"]; ok {
										if s, ok := objval["columns"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns dclService.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns
													if _, ok := objval["encoding"]; ok {
														if s, ok := objval["encoding"].(string); ok {
															rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.Encoding = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.Encoding: expected string")
														}
													}
													if _, ok := objval["fieldName"]; ok {
														if s, ok := objval["fieldName"].(string); ok {
															rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.FieldName = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.FieldName: expected string")
														}
													}
													if _, ok := objval["onlyReadLatest"]; ok {
														if b, ok := objval["onlyReadLatest"].(bool); ok {
															rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.OnlyReadLatest = dcl.Bool(b)
														} else {
															return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.OnlyReadLatest: expected bool")
														}
													}
													if _, ok := objval["qualifierEncoded"]; ok {
														if s, ok := objval["qualifierEncoded"].(string); ok {
															rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.QualifierEncoded = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.QualifierEncoded: expected string")
														}
													}
													if _, ok := objval["qualifierString"]; ok {
														if s, ok := objval["qualifierString"].(string); ok {
															rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.QualifierString = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.QualifierString: expected string")
														}
													}
													if _, ok := objval["type"]; ok {
														if s, ok := objval["type"].(string); ok {
															rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.Type = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.Type: expected string")
														}
													}
													rExternalDataConfigurationBigtableOptionsColumnFamilies.Columns = append(rExternalDataConfigurationBigtableOptionsColumnFamilies.Columns, rExternalDataConfigurationBigtableOptionsColumnFamiliesColumns)
												}
											}
										} else {
											return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamilies.Columns: expected []interface{}")
										}
									}
									if _, ok := objval["encoding"]; ok {
										if s, ok := objval["encoding"].(string); ok {
											rExternalDataConfigurationBigtableOptionsColumnFamilies.Encoding = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamilies.Encoding: expected string")
										}
									}
									if _, ok := objval["familyId"]; ok {
										if s, ok := objval["familyId"].(string); ok {
											rExternalDataConfigurationBigtableOptionsColumnFamilies.FamilyId = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamilies.FamilyId: expected string")
										}
									}
									if _, ok := objval["onlyReadLatest"]; ok {
										if b, ok := objval["onlyReadLatest"].(bool); ok {
											rExternalDataConfigurationBigtableOptionsColumnFamilies.OnlyReadLatest = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamilies.OnlyReadLatest: expected bool")
										}
									}
									if _, ok := objval["type"]; ok {
										if s, ok := objval["type"].(string); ok {
											rExternalDataConfigurationBigtableOptionsColumnFamilies.Type = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rExternalDataConfigurationBigtableOptionsColumnFamilies.Type: expected string")
										}
									}
									r.ExternalDataConfiguration.BigtableOptions.ColumnFamilies = append(r.ExternalDataConfiguration.BigtableOptions.ColumnFamilies, rExternalDataConfigurationBigtableOptionsColumnFamilies)
								}
							}
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.BigtableOptions.ColumnFamilies: expected []interface{}")
						}
					}
					if _, ok := rExternalDataConfigurationBigtableOptions["ignoreUnspecifiedColumnFamilies"]; ok {
						if b, ok := rExternalDataConfigurationBigtableOptions["ignoreUnspecifiedColumnFamilies"].(bool); ok {
							r.ExternalDataConfiguration.BigtableOptions.IgnoreUnspecifiedColumnFamilies = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.BigtableOptions.IgnoreUnspecifiedColumnFamilies: expected bool")
						}
					}
					if _, ok := rExternalDataConfigurationBigtableOptions["readRowkeyAsString"]; ok {
						if b, ok := rExternalDataConfigurationBigtableOptions["readRowkeyAsString"].(bool); ok {
							r.ExternalDataConfiguration.BigtableOptions.ReadRowkeyAsString = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.BigtableOptions.ReadRowkeyAsString: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.BigtableOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["compression"]; ok {
				if s, ok := rExternalDataConfiguration["compression"].(string); ok {
					r.ExternalDataConfiguration.Compression = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.Compression: expected string")
				}
			}
			if _, ok := rExternalDataConfiguration["connectionId"]; ok {
				if s, ok := rExternalDataConfiguration["connectionId"].(string); ok {
					r.ExternalDataConfiguration.ConnectionId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.ConnectionId: expected string")
				}
			}
			if _, ok := rExternalDataConfiguration["csvOptions"]; ok {
				if rExternalDataConfigurationCsvOptions, ok := rExternalDataConfiguration["csvOptions"].(map[string]interface{}); ok {
					r.ExternalDataConfiguration.CsvOptions = &dclService.TableExternalDataConfigurationCsvOptions{}
					if _, ok := rExternalDataConfigurationCsvOptions["allowJaggedRows"]; ok {
						if b, ok := rExternalDataConfigurationCsvOptions["allowJaggedRows"].(bool); ok {
							r.ExternalDataConfiguration.CsvOptions.AllowJaggedRows = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.CsvOptions.AllowJaggedRows: expected bool")
						}
					}
					if _, ok := rExternalDataConfigurationCsvOptions["allowQuotedNewlines"]; ok {
						if b, ok := rExternalDataConfigurationCsvOptions["allowQuotedNewlines"].(bool); ok {
							r.ExternalDataConfiguration.CsvOptions.AllowQuotedNewlines = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.CsvOptions.AllowQuotedNewlines: expected bool")
						}
					}
					if _, ok := rExternalDataConfigurationCsvOptions["encoding"]; ok {
						if s, ok := rExternalDataConfigurationCsvOptions["encoding"].(string); ok {
							r.ExternalDataConfiguration.CsvOptions.Encoding = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.CsvOptions.Encoding: expected string")
						}
					}
					if _, ok := rExternalDataConfigurationCsvOptions["fieldDelimiter"]; ok {
						if s, ok := rExternalDataConfigurationCsvOptions["fieldDelimiter"].(string); ok {
							r.ExternalDataConfiguration.CsvOptions.FieldDelimiter = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.CsvOptions.FieldDelimiter: expected string")
						}
					}
					if _, ok := rExternalDataConfigurationCsvOptions["quote"]; ok {
						if s, ok := rExternalDataConfigurationCsvOptions["quote"].(string); ok {
							r.ExternalDataConfiguration.CsvOptions.Quote = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.CsvOptions.Quote: expected string")
						}
					}
					if _, ok := rExternalDataConfigurationCsvOptions["skipLeadingRows"]; ok {
						if s, ok := rExternalDataConfigurationCsvOptions["skipLeadingRows"].(string); ok {
							r.ExternalDataConfiguration.CsvOptions.SkipLeadingRows = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.CsvOptions.SkipLeadingRows: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.CsvOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["decimalTargetTypes"]; ok {
				if s, ok := rExternalDataConfiguration["decimalTargetTypes"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.ExternalDataConfiguration.DecimalTargetTypes = append(r.ExternalDataConfiguration.DecimalTargetTypes, dclService.TableExternalDataConfigurationDecimalTargetTypesEnum(strval))
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.DecimalTargetTypes: expected []interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["googleSheetsOptions"]; ok {
				if rExternalDataConfigurationGoogleSheetsOptions, ok := rExternalDataConfiguration["googleSheetsOptions"].(map[string]interface{}); ok {
					r.ExternalDataConfiguration.GoogleSheetsOptions = &dclService.TableExternalDataConfigurationGoogleSheetsOptions{}
					if _, ok := rExternalDataConfigurationGoogleSheetsOptions["range"]; ok {
						if s, ok := rExternalDataConfigurationGoogleSheetsOptions["range"].(string); ok {
							r.ExternalDataConfiguration.GoogleSheetsOptions.Range = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.GoogleSheetsOptions.Range: expected string")
						}
					}
					if _, ok := rExternalDataConfigurationGoogleSheetsOptions["skipLeadingRows"]; ok {
						if s, ok := rExternalDataConfigurationGoogleSheetsOptions["skipLeadingRows"].(string); ok {
							r.ExternalDataConfiguration.GoogleSheetsOptions.SkipLeadingRows = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.GoogleSheetsOptions.SkipLeadingRows: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.GoogleSheetsOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["hivePartitioningOptions"]; ok {
				if rExternalDataConfigurationHivePartitioningOptions, ok := rExternalDataConfiguration["hivePartitioningOptions"].(map[string]interface{}); ok {
					r.ExternalDataConfiguration.HivePartitioningOptions = &dclService.TableExternalDataConfigurationHivePartitioningOptions{}
					if _, ok := rExternalDataConfigurationHivePartitioningOptions["fields"]; ok {
						if s, ok := rExternalDataConfigurationHivePartitioningOptions["fields"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.ExternalDataConfiguration.HivePartitioningOptions.Fields = append(r.ExternalDataConfiguration.HivePartitioningOptions.Fields, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.HivePartitioningOptions.Fields: expected []interface{}")
						}
					}
					if _, ok := rExternalDataConfigurationHivePartitioningOptions["mode"]; ok {
						if s, ok := rExternalDataConfigurationHivePartitioningOptions["mode"].(string); ok {
							r.ExternalDataConfiguration.HivePartitioningOptions.Mode = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.HivePartitioningOptions.Mode: expected string")
						}
					}
					if _, ok := rExternalDataConfigurationHivePartitioningOptions["requirePartitionFilter"]; ok {
						if b, ok := rExternalDataConfigurationHivePartitioningOptions["requirePartitionFilter"].(bool); ok {
							r.ExternalDataConfiguration.HivePartitioningOptions.RequirePartitionFilter = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.HivePartitioningOptions.RequirePartitionFilter: expected bool")
						}
					}
					if _, ok := rExternalDataConfigurationHivePartitioningOptions["sourceUriPrefix"]; ok {
						if s, ok := rExternalDataConfigurationHivePartitioningOptions["sourceUriPrefix"].(string); ok {
							r.ExternalDataConfiguration.HivePartitioningOptions.SourceUriPrefix = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.HivePartitioningOptions.SourceUriPrefix: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.HivePartitioningOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["ignoreUnknownValues"]; ok {
				if b, ok := rExternalDataConfiguration["ignoreUnknownValues"].(bool); ok {
					r.ExternalDataConfiguration.IgnoreUnknownValues = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.IgnoreUnknownValues: expected bool")
				}
			}
			if _, ok := rExternalDataConfiguration["jsonExtension"]; ok {
				if s, ok := rExternalDataConfiguration["jsonExtension"].(string); ok {
					r.ExternalDataConfiguration.JsonExtension = dclService.TableExternalDataConfigurationJsonExtensionEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.JsonExtension: expected string")
				}
			}
			if _, ok := rExternalDataConfiguration["maxBadRecords"]; ok {
				if i, ok := rExternalDataConfiguration["maxBadRecords"].(int64); ok {
					r.ExternalDataConfiguration.MaxBadRecords = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.MaxBadRecords: expected int64")
				}
			}
			if _, ok := rExternalDataConfiguration["maxBadRecordsAlternative"]; ok {
				if s, ok := rExternalDataConfiguration["maxBadRecordsAlternative"].([]interface{}); ok {
					for _, ss := range s {
						if intval, ok := ss.(int64); ok {
							r.ExternalDataConfiguration.MaxBadRecordsAlternative = append(r.ExternalDataConfiguration.MaxBadRecordsAlternative, intval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.MaxBadRecordsAlternative: expected []interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["parquetOptions"]; ok {
				if rExternalDataConfigurationParquetOptions, ok := rExternalDataConfiguration["parquetOptions"].(map[string]interface{}); ok {
					r.ExternalDataConfiguration.ParquetOptions = &dclService.TableExternalDataConfigurationParquetOptions{}
					if _, ok := rExternalDataConfigurationParquetOptions["enableListInference"]; ok {
						if b, ok := rExternalDataConfigurationParquetOptions["enableListInference"].(bool); ok {
							r.ExternalDataConfiguration.ParquetOptions.EnableListInference = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.ParquetOptions.EnableListInference: expected bool")
						}
					}
					if _, ok := rExternalDataConfigurationParquetOptions["enumAsString"]; ok {
						if b, ok := rExternalDataConfigurationParquetOptions["enumAsString"].(bool); ok {
							r.ExternalDataConfiguration.ParquetOptions.EnumAsString = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.ParquetOptions.EnumAsString: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.ParquetOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["schema"]; ok {
				if rExternalDataConfigurationSchema, ok := rExternalDataConfiguration["schema"].(map[string]interface{}); ok {
					r.ExternalDataConfiguration.Schema = &dclService.TableExternalDataConfigurationSchema{}
					if _, ok := rExternalDataConfigurationSchema["fields"]; ok {
						if s, ok := rExternalDataConfigurationSchema["fields"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									unstructuredObjval, err := UnstructuredToTableGooglecloudbigqueryv2Tablefieldschema(objval)
									if err != nil {
										return nil, err
									}
									r.ExternalDataConfiguration.Schema.Fields = append(r.ExternalDataConfiguration.Schema.Fields, *unstructuredObjval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.Schema.Fields: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.Schema: expected map[string]interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["sourceFormat"]; ok {
				if s, ok := rExternalDataConfiguration["sourceFormat"].(string); ok {
					r.ExternalDataConfiguration.SourceFormat = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.SourceFormat: expected string")
				}
			}
			if _, ok := rExternalDataConfiguration["sourceUris"]; ok {
				if s, ok := rExternalDataConfiguration["sourceUris"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.ExternalDataConfiguration.SourceUris = append(r.ExternalDataConfiguration.SourceUris, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.SourceUris: expected []interface{}")
				}
			}
			if _, ok := rExternalDataConfiguration["valueConversionModes"]; ok {
				if rExternalDataConfigurationValueConversionModes, ok := rExternalDataConfiguration["valueConversionModes"].(map[string]interface{}); ok {
					r.ExternalDataConfiguration.ValueConversionModes = &dclService.TableExternalDataConfigurationValueConversionModes{}
					if _, ok := rExternalDataConfigurationValueConversionModes["numericTypeOutOfRangeConversionMode"]; ok {
						if s, ok := rExternalDataConfigurationValueConversionModes["numericTypeOutOfRangeConversionMode"].(string); ok {
							r.ExternalDataConfiguration.ValueConversionModes.NumericTypeOutOfRangeConversionMode = dclService.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.ValueConversionModes.NumericTypeOutOfRangeConversionMode: expected string")
						}
					}
					if _, ok := rExternalDataConfigurationValueConversionModes["temporalTypesOutOfRangeConversionMode"]; ok {
						if s, ok := rExternalDataConfigurationValueConversionModes["temporalTypesOutOfRangeConversionMode"].(string); ok {
							r.ExternalDataConfiguration.ValueConversionModes.TemporalTypesOutOfRangeConversionMode = dclService.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.ExternalDataConfiguration.ValueConversionModes.TemporalTypesOutOfRangeConversionMode: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ExternalDataConfiguration.ValueConversionModes: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ExternalDataConfiguration: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["friendlyName"]; ok {
		if s, ok := u.Object["friendlyName"].(string); ok {
			r.FriendlyName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.FriendlyName: expected string")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if s, ok := u.Object["id"].(string); ok {
			r.Id = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Id: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["lastModifiedTime"]; ok {
		if i, ok := u.Object["lastModifiedTime"].(int64); ok {
			r.LastModifiedTime = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.LastModifiedTime: expected int64")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["materializedView"]; ok {
		if rMaterializedView, ok := u.Object["materializedView"].(map[string]interface{}); ok {
			r.MaterializedView = &dclService.TableMaterializedView{}
			if _, ok := rMaterializedView["enableRefresh"]; ok {
				if b, ok := rMaterializedView["enableRefresh"].(bool); ok {
					r.MaterializedView.EnableRefresh = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.MaterializedView.EnableRefresh: expected bool")
				}
			}
			if _, ok := rMaterializedView["lastRefreshTime"]; ok {
				if i, ok := rMaterializedView["lastRefreshTime"].(int64); ok {
					r.MaterializedView.LastRefreshTime = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.MaterializedView.LastRefreshTime: expected int64")
				}
			}
			if _, ok := rMaterializedView["query"]; ok {
				if s, ok := rMaterializedView["query"].(string); ok {
					r.MaterializedView.Query = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MaterializedView.Query: expected string")
				}
			}
			if _, ok := rMaterializedView["refreshIntervalMs"]; ok {
				if i, ok := rMaterializedView["refreshIntervalMs"].(int64); ok {
					r.MaterializedView.RefreshIntervalMs = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.MaterializedView.RefreshIntervalMs: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MaterializedView: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["model"]; ok {
		if rModel, ok := u.Object["model"].(map[string]interface{}); ok {
			r.Model = &dclService.TableModel{}
			if _, ok := rModel["modelOptions"]; ok {
				if rModelModelOptions, ok := rModel["modelOptions"].(map[string]interface{}); ok {
					r.Model.ModelOptions = &dclService.TableModelModelOptions{}
					if _, ok := rModelModelOptions["labels"]; ok {
						if s, ok := rModelModelOptions["labels"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Model.ModelOptions.Labels = append(r.Model.ModelOptions.Labels, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Model.ModelOptions.Labels: expected []interface{}")
						}
					}
					if _, ok := rModelModelOptions["lossType"]; ok {
						if s, ok := rModelModelOptions["lossType"].(string); ok {
							r.Model.ModelOptions.LossType = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Model.ModelOptions.LossType: expected string")
						}
					}
					if _, ok := rModelModelOptions["modelType"]; ok {
						if s, ok := rModelModelOptions["modelType"].(string); ok {
							r.Model.ModelOptions.ModelType = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Model.ModelOptions.ModelType: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Model.ModelOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rModel["trainingRuns"]; ok {
				if s, ok := rModel["trainingRuns"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rModelTrainingRuns dclService.TableModelTrainingRuns
							if _, ok := objval["iterationResults"]; ok {
								if s, ok := objval["iterationResults"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rModelTrainingRunsIterationResults dclService.TableModelTrainingRunsIterationResults
											if _, ok := objval["durationMs"]; ok {
												if s, ok := objval["durationMs"].(string); ok {
													rModelTrainingRunsIterationResults.DurationMs = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rModelTrainingRunsIterationResults.DurationMs: expected string")
												}
											}
											if _, ok := objval["evalLoss"]; ok {
												if f, ok := objval["evalLoss"].(float64); ok {
													rModelTrainingRunsIterationResults.EvalLoss = dcl.Float64(f)
												} else {
													return nil, fmt.Errorf("rModelTrainingRunsIterationResults.EvalLoss: expected float64")
												}
											}
											if _, ok := objval["index"]; ok {
												if i, ok := objval["index"].(int64); ok {
													rModelTrainingRunsIterationResults.Index = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rModelTrainingRunsIterationResults.Index: expected int64")
												}
											}
											if _, ok := objval["learnRate"]; ok {
												if f, ok := objval["learnRate"].(float64); ok {
													rModelTrainingRunsIterationResults.LearnRate = dcl.Float64(f)
												} else {
													return nil, fmt.Errorf("rModelTrainingRunsIterationResults.LearnRate: expected float64")
												}
											}
											if _, ok := objval["trainingLoss"]; ok {
												if f, ok := objval["trainingLoss"].(float64); ok {
													rModelTrainingRunsIterationResults.TrainingLoss = dcl.Float64(f)
												} else {
													return nil, fmt.Errorf("rModelTrainingRunsIterationResults.TrainingLoss: expected float64")
												}
											}
											rModelTrainingRuns.IterationResults = append(rModelTrainingRuns.IterationResults, rModelTrainingRunsIterationResults)
										}
									}
								} else {
									return nil, fmt.Errorf("rModelTrainingRuns.IterationResults: expected []interface{}")
								}
							}
							if _, ok := objval["startTime"]; ok {
								if s, ok := objval["startTime"].(string); ok {
									rModelTrainingRuns.StartTime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rModelTrainingRuns.StartTime: expected string")
								}
							}
							if _, ok := objval["state"]; ok {
								if s, ok := objval["state"].(string); ok {
									rModelTrainingRuns.State = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rModelTrainingRuns.State: expected string")
								}
							}
							if _, ok := objval["trainingOptions"]; ok {
								if rModelTrainingRunsTrainingOptions, ok := objval["trainingOptions"].(map[string]interface{}); ok {
									rModelTrainingRuns.TrainingOptions = &dclService.TableModelTrainingRunsTrainingOptions{}
									if _, ok := rModelTrainingRunsTrainingOptions["earlyStop"]; ok {
										if b, ok := rModelTrainingRunsTrainingOptions["earlyStop"].(bool); ok {
											rModelTrainingRuns.TrainingOptions.EarlyStop = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions.EarlyStop: expected bool")
										}
									}
									if _, ok := rModelTrainingRunsTrainingOptions["l1Reg"]; ok {
										if f, ok := rModelTrainingRunsTrainingOptions["l1Reg"].(float64); ok {
											rModelTrainingRuns.TrainingOptions.L1Reg = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions.L1Reg: expected float64")
										}
									}
									if _, ok := rModelTrainingRunsTrainingOptions["l2Reg"]; ok {
										if f, ok := rModelTrainingRunsTrainingOptions["l2Reg"].(float64); ok {
											rModelTrainingRuns.TrainingOptions.L2Reg = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions.L2Reg: expected float64")
										}
									}
									if _, ok := rModelTrainingRunsTrainingOptions["learnRate"]; ok {
										if f, ok := rModelTrainingRunsTrainingOptions["learnRate"].(float64); ok {
											rModelTrainingRuns.TrainingOptions.LearnRate = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions.LearnRate: expected float64")
										}
									}
									if _, ok := rModelTrainingRunsTrainingOptions["learnRateStrategy"]; ok {
										if s, ok := rModelTrainingRunsTrainingOptions["learnRateStrategy"].(string); ok {
											rModelTrainingRuns.TrainingOptions.LearnRateStrategy = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions.LearnRateStrategy: expected string")
										}
									}
									if _, ok := rModelTrainingRunsTrainingOptions["lineSearchInitLearnRate"]; ok {
										if f, ok := rModelTrainingRunsTrainingOptions["lineSearchInitLearnRate"].(float64); ok {
											rModelTrainingRuns.TrainingOptions.LineSearchInitLearnRate = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions.LineSearchInitLearnRate: expected float64")
										}
									}
									if _, ok := rModelTrainingRunsTrainingOptions["maxIteration"]; ok {
										if i, ok := rModelTrainingRunsTrainingOptions["maxIteration"].(int64); ok {
											rModelTrainingRuns.TrainingOptions.MaxIteration = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions.MaxIteration: expected int64")
										}
									}
									if _, ok := rModelTrainingRunsTrainingOptions["minRelProgress"]; ok {
										if f, ok := rModelTrainingRunsTrainingOptions["minRelProgress"].(float64); ok {
											rModelTrainingRuns.TrainingOptions.MinRelProgress = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions.MinRelProgress: expected float64")
										}
									}
									if _, ok := rModelTrainingRunsTrainingOptions["warmStart"]; ok {
										if b, ok := rModelTrainingRunsTrainingOptions["warmStart"].(bool); ok {
											rModelTrainingRuns.TrainingOptions.WarmStart = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions.WarmStart: expected bool")
										}
									}
								} else {
									return nil, fmt.Errorf("rModelTrainingRuns.TrainingOptions: expected map[string]interface{}")
								}
							}
							r.Model.TrainingRuns = append(r.Model.TrainingRuns, rModelTrainingRuns)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Model.TrainingRuns: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Model: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["numBytes"]; ok {
		if s, ok := u.Object["numBytes"].(string); ok {
			r.NumBytes = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NumBytes: expected string")
		}
	}
	if _, ok := u.Object["numLongTermBytes"]; ok {
		if s, ok := u.Object["numLongTermBytes"].(string); ok {
			r.NumLongTermBytes = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NumLongTermBytes: expected string")
		}
	}
	if _, ok := u.Object["numPhysicalBytes"]; ok {
		if s, ok := u.Object["numPhysicalBytes"].(string); ok {
			r.NumPhysicalBytes = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NumPhysicalBytes: expected string")
		}
	}
	if _, ok := u.Object["numRows"]; ok {
		if i, ok := u.Object["numRows"].(int64); ok {
			r.NumRows = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.NumRows: expected int64")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["rangePartitioning"]; ok {
		if rRangePartitioning, ok := u.Object["rangePartitioning"].(map[string]interface{}); ok {
			r.RangePartitioning = &dclService.TableRangePartitioning{}
			if _, ok := rRangePartitioning["field"]; ok {
				if s, ok := rRangePartitioning["field"].(string); ok {
					r.RangePartitioning.Field = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.RangePartitioning.Field: expected string")
				}
			}
			if _, ok := rRangePartitioning["range"]; ok {
				if rRangePartitioningRange, ok := rRangePartitioning["range"].(map[string]interface{}); ok {
					r.RangePartitioning.Range = &dclService.TableRangePartitioningRange{}
					if _, ok := rRangePartitioningRange["end"]; ok {
						if s, ok := rRangePartitioningRange["end"].(string); ok {
							r.RangePartitioning.Range.End = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.RangePartitioning.Range.End: expected string")
						}
					}
					if _, ok := rRangePartitioningRange["interval"]; ok {
						if s, ok := rRangePartitioningRange["interval"].(string); ok {
							r.RangePartitioning.Range.Interval = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.RangePartitioning.Range.Interval: expected string")
						}
					}
					if _, ok := rRangePartitioningRange["start"]; ok {
						if s, ok := rRangePartitioningRange["start"].(string); ok {
							r.RangePartitioning.Range.Start = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.RangePartitioning.Range.Start: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.RangePartitioning.Range: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.RangePartitioning: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["requirePartitionFilter"]; ok {
		if b, ok := u.Object["requirePartitionFilter"].(bool); ok {
			r.RequirePartitionFilter = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.RequirePartitionFilter: expected bool")
		}
	}
	if _, ok := u.Object["schema"]; ok {
		if rSchema, ok := u.Object["schema"].(map[string]interface{}); ok {
			r.Schema = &dclService.TableSchema{}
			if _, ok := rSchema["fields"]; ok {
				if s, ok := rSchema["fields"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rSchemaFields dclService.TableGooglecloudbigqueryv2Tablefieldschema
							if _, ok := objval["categories"]; ok {
								if rSchemaFieldsCategories, ok := objval["categories"].(map[string]interface{}); ok {
									rSchemaFields.Categories = &dclService.TableGooglecloudbigqueryv2TablefieldschemaCategories{}
									if _, ok := rSchemaFieldsCategories["names"]; ok {
										if s, ok := rSchemaFieldsCategories["names"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rSchemaFields.Categories.Names = append(rSchemaFields.Categories.Names, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rSchemaFields.Categories.Names: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rSchemaFields.Categories: expected map[string]interface{}")
								}
							}
							if _, ok := objval["collation"]; ok {
								if s, ok := objval["collation"].(string); ok {
									rSchemaFields.Collation = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSchemaFields.Collation: expected string")
								}
							}
							if _, ok := objval["defaultValueExpression"]; ok {
								if s, ok := objval["defaultValueExpression"].(string); ok {
									rSchemaFields.DefaultValueExpression = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSchemaFields.DefaultValueExpression: expected string")
								}
							}
							if _, ok := objval["description"]; ok {
								if s, ok := objval["description"].(string); ok {
									rSchemaFields.Description = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSchemaFields.Description: expected string")
								}
							}
							if _, ok := objval["fields"]; ok {
								if s, ok := objval["fields"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											unstructuredObjval, err := UnstructuredToTableGooglecloudbigqueryv2Tablefieldschema(objval)
											if err != nil {
												return nil, err
											}
											rSchemaFields.Fields = append(rSchemaFields.Fields, *unstructuredObjval)
										}
									}
								} else {
									return nil, fmt.Errorf("rSchemaFields.Fields: expected []interface{}")
								}
							}
							if _, ok := objval["maxLength"]; ok {
								if i, ok := objval["maxLength"].(int64); ok {
									rSchemaFields.MaxLength = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rSchemaFields.MaxLength: expected int64")
								}
							}
							if _, ok := objval["mode"]; ok {
								if s, ok := objval["mode"].(string); ok {
									rSchemaFields.Mode = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSchemaFields.Mode: expected string")
								}
							}
							if _, ok := objval["name"]; ok {
								if s, ok := objval["name"].(string); ok {
									rSchemaFields.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSchemaFields.Name: expected string")
								}
							}
							if _, ok := objval["nameAlternative"]; ok {
								if s, ok := objval["nameAlternative"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rSchemaFields.NameAlternative = append(rSchemaFields.NameAlternative, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rSchemaFields.NameAlternative: expected []interface{}")
								}
							}
							if _, ok := objval["policyTags"]; ok {
								if rSchemaFieldsPolicyTags, ok := objval["policyTags"].(map[string]interface{}); ok {
									rSchemaFields.PolicyTags = &dclService.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
									if _, ok := rSchemaFieldsPolicyTags["names"]; ok {
										if s, ok := rSchemaFieldsPolicyTags["names"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rSchemaFields.PolicyTags.Names = append(rSchemaFields.PolicyTags.Names, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rSchemaFields.PolicyTags.Names: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rSchemaFields.PolicyTags: expected map[string]interface{}")
								}
							}
							if _, ok := objval["precision"]; ok {
								if i, ok := objval["precision"].(int64); ok {
									rSchemaFields.Precision = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rSchemaFields.Precision: expected int64")
								}
							}
							if _, ok := objval["scale"]; ok {
								if i, ok := objval["scale"].(int64); ok {
									rSchemaFields.Scale = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rSchemaFields.Scale: expected int64")
								}
							}
							if _, ok := objval["type"]; ok {
								if s, ok := objval["type"].(string); ok {
									rSchemaFields.Type = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSchemaFields.Type: expected string")
								}
							}
							r.Schema.Fields = append(r.Schema.Fields, rSchemaFields)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Schema.Fields: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Schema: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["snapshotDefinition"]; ok {
		if rSnapshotDefinition, ok := u.Object["snapshotDefinition"].(map[string]interface{}); ok {
			r.SnapshotDefinition = &dclService.TableSnapshotDefinition{}
			if _, ok := rSnapshotDefinition["dataset"]; ok {
				if s, ok := rSnapshotDefinition["dataset"].(string); ok {
					r.SnapshotDefinition.Dataset = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SnapshotDefinition.Dataset: expected string")
				}
			}
			if _, ok := rSnapshotDefinition["project"]; ok {
				if s, ok := rSnapshotDefinition["project"].(string); ok {
					r.SnapshotDefinition.Project = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SnapshotDefinition.Project: expected string")
				}
			}
			if _, ok := rSnapshotDefinition["snapshotTime"]; ok {
				if s, ok := rSnapshotDefinition["snapshotTime"].(string); ok {
					r.SnapshotDefinition.SnapshotTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SnapshotDefinition.SnapshotTime: expected string")
				}
			}
			if _, ok := rSnapshotDefinition["table"]; ok {
				if s, ok := rSnapshotDefinition["table"].(string); ok {
					r.SnapshotDefinition.Table = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SnapshotDefinition.Table: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SnapshotDefinition: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["streamingBuffer"]; ok {
		if rStreamingBuffer, ok := u.Object["streamingBuffer"].(map[string]interface{}); ok {
			r.StreamingBuffer = &dclService.TableStreamingBuffer{}
			if _, ok := rStreamingBuffer["estimatedBytes"]; ok {
				if i, ok := rStreamingBuffer["estimatedBytes"].(int64); ok {
					r.StreamingBuffer.EstimatedBytes = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.StreamingBuffer.EstimatedBytes: expected int64")
				}
			}
			if _, ok := rStreamingBuffer["estimatedRows"]; ok {
				if i, ok := rStreamingBuffer["estimatedRows"].(int64); ok {
					r.StreamingBuffer.EstimatedRows = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.StreamingBuffer.EstimatedRows: expected int64")
				}
			}
			if _, ok := rStreamingBuffer["oldestEntryTime"]; ok {
				if i, ok := rStreamingBuffer["oldestEntryTime"].(int64); ok {
					r.StreamingBuffer.OldestEntryTime = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.StreamingBuffer.OldestEntryTime: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.StreamingBuffer: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["timePartitioning"]; ok {
		if rTimePartitioning, ok := u.Object["timePartitioning"].(map[string]interface{}); ok {
			r.TimePartitioning = &dclService.TableTimePartitioning{}
			if _, ok := rTimePartitioning["expirationMs"]; ok {
				if s, ok := rTimePartitioning["expirationMs"].(string); ok {
					r.TimePartitioning.ExpirationMs = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.TimePartitioning.ExpirationMs: expected string")
				}
			}
			if _, ok := rTimePartitioning["field"]; ok {
				if s, ok := rTimePartitioning["field"].(string); ok {
					r.TimePartitioning.Field = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.TimePartitioning.Field: expected string")
				}
			}
			if _, ok := rTimePartitioning["type"]; ok {
				if s, ok := rTimePartitioning["type"].(string); ok {
					r.TimePartitioning.Type = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.TimePartitioning.Type: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.TimePartitioning: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
		}
	}
	if _, ok := u.Object["view"]; ok {
		if rView, ok := u.Object["view"].(map[string]interface{}); ok {
			r.View = &dclService.TableView{}
			if _, ok := rView["query"]; ok {
				if s, ok := rView["query"].(string); ok {
					r.View.Query = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.View.Query: expected string")
				}
			}
			if _, ok := rView["useExplicitColumnNames"]; ok {
				if b, ok := rView["useExplicitColumnNames"].(bool); ok {
					r.View.UseExplicitColumnNames = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.View.UseExplicitColumnNames: expected bool")
				}
			}
			if _, ok := rView["useLegacySql"]; ok {
				if b, ok := rView["useLegacySql"].(bool); ok {
					r.View.UseLegacySql = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.View.UseLegacySql: expected bool")
				}
			}
			if _, ok := rView["userDefinedFunctionResources"]; ok {
				if s, ok := rView["userDefinedFunctionResources"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rViewUserDefinedFunctionResources dclService.TableViewUserDefinedFunctionResources
							if _, ok := objval["inlineCode"]; ok {
								if s, ok := objval["inlineCode"].(string); ok {
									rViewUserDefinedFunctionResources.InlineCode = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rViewUserDefinedFunctionResources.InlineCode: expected string")
								}
							}
							if _, ok := objval["inlineCodeAlternative"]; ok {
								if s, ok := objval["inlineCodeAlternative"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rViewUserDefinedFunctionResources.InlineCodeAlternative = append(rViewUserDefinedFunctionResources.InlineCodeAlternative, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rViewUserDefinedFunctionResources.InlineCodeAlternative: expected []interface{}")
								}
							}
							if _, ok := objval["resourceUri"]; ok {
								if s, ok := objval["resourceUri"].(string); ok {
									rViewUserDefinedFunctionResources.ResourceUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rViewUserDefinedFunctionResources.ResourceUri: expected string")
								}
							}
							r.View.UserDefinedFunctionResources = append(r.View.UserDefinedFunctionResources, rViewUserDefinedFunctionResources)
						}
					}
				} else {
					return nil, fmt.Errorf("r.View.UserDefinedFunctionResources: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.View: expected map[string]interface{}")
		}
	}
	return r, nil
}

func UnstructuredToTableGooglecloudbigqueryv2Tablefieldschema(obj map[string]interface{}) (*dclService.TableGooglecloudbigqueryv2Tablefieldschema, error) {
	r := &dclService.TableGooglecloudbigqueryv2Tablefieldschema{}
	if _, ok := obj["categories"]; ok {
		if rCategories, ok := obj["categories"].(map[string]interface{}); ok {
			r.Categories = &dclService.TableGooglecloudbigqueryv2TablefieldschemaCategories{}
			if _, ok := rCategories["names"]; ok {
				if s, ok := rCategories["names"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Categories.Names = append(r.Categories.Names, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Categories.Names: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Categories: expected map[string]interface{}")
		}
	}
	if _, ok := obj["collation"]; ok {
		if s, ok := obj["collation"].(string); ok {
			r.Collation = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Collation: expected string")
		}
	}
	if _, ok := obj["defaultValueExpression"]; ok {
		if s, ok := obj["defaultValueExpression"].(string); ok {
			r.DefaultValueExpression = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DefaultValueExpression: expected string")
		}
	}
	if _, ok := obj["description"]; ok {
		if s, ok := obj["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := obj["fields"]; ok {
		if s, ok := obj["fields"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					unstructuredObjval, err := UnstructuredToTableGooglecloudbigqueryv2Tablefieldschema(objval)
					if err != nil {
						return nil, err
					}
					r.Fields = append(r.Fields, *unstructuredObjval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Fields: expected []interface{}")
		}
	}
	if _, ok := obj["maxLength"]; ok {
		if i, ok := obj["maxLength"].(int64); ok {
			r.MaxLength = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.MaxLength: expected int64")
		}
	}
	if _, ok := obj["mode"]; ok {
		if s, ok := obj["mode"].(string); ok {
			r.Mode = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Mode: expected string")
		}
	}
	if _, ok := obj["name"]; ok {
		if s, ok := obj["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := obj["nameAlternative"]; ok {
		if s, ok := obj["nameAlternative"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.NameAlternative = append(r.NameAlternative, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.NameAlternative: expected []interface{}")
		}
	}
	if _, ok := obj["policyTags"]; ok {
		if rPolicyTags, ok := obj["policyTags"].(map[string]interface{}); ok {
			r.PolicyTags = &dclService.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
			if _, ok := rPolicyTags["names"]; ok {
				if s, ok := rPolicyTags["names"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.PolicyTags.Names = append(r.PolicyTags.Names, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.PolicyTags.Names: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PolicyTags: expected map[string]interface{}")
		}
	}
	if _, ok := obj["precision"]; ok {
		if i, ok := obj["precision"].(int64); ok {
			r.Precision = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Precision: expected int64")
		}
	}
	if _, ok := obj["scale"]; ok {
		if i, ok := obj["scale"].(int64); ok {
			r.Scale = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Scale: expected int64")
		}
	}
	if _, ok := obj["type"]; ok {
		if s, ok := obj["type"].(string); ok {
			r.Type = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
		}
	}
	return r, nil
}

func GetTable(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTable(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetTable(ctx, r)
	if err != nil {
		return nil, err
	}
	return TableToUnstructured(r), nil
}

func ListTable(ctx context.Context, config *dcl.Config, project string, dataset string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListTable(ctx, project, dataset)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, TableToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyTable(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTable(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTable(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyTable(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return TableToUnstructured(r), nil
}

func TableHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTable(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTable(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyTable(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteTable(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTable(u)
	if err != nil {
		return err
	}
	return c.DeleteTable(ctx, r)
}

func TableID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToTable(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Table) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"bigquery",
		"Table",
		"ga",
	}
}

func (r *Table) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Table) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Table) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Table) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Table) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Table) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Table) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetTable(ctx, config, resource)
}

func (r *Table) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyTable(ctx, config, resource, opts...)
}

func (r *Table) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return TableHasDiff(ctx, config, resource, opts...)
}

func (r *Table) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteTable(ctx, config, resource)
}

func (r *Table) ID(resource *unstructured.Resource) (string, error) {
	return TableID(resource)
}

func init() {
	unstructured.Register(&Table{})
}
