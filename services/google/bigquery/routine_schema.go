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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLRoutineSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Bigquery/Routine",
			Description: "The Bigquery Routine resource",
			StructName:  "Routine",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Routine",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Routine",
						Required:    true,
						Description: "A full instance of a Routine",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Routine",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Routine",
						Required:    true,
						Description: "A full instance of a Routine",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Routine",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Routine",
						Required:    true,
						Description: "A full instance of a Routine",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Routine",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "dataset",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Routine",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "dataset",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"ArgumentsDataType": &dcl.Component{
					SchemaProperty: dcl.Property{
						Type:        "object",
						GoName:      "DataType",
						GoType:      "RoutineArgumentsDataType",
						Description: "Required unless argument_kind = ANY_TYPE.",
						Immutable:   true,
						Required: []string{
							"typeKind",
						},
						Properties: map[string]*dcl.Property{
							"arrayElementType": &dcl.Property{
								Ref:       "#/components/schemas/ArgumentsDataType",
								GoName:    "ArrayElementType",
								Immutable: true,
								Conflicts: []string{
									"structType",
								},
							},
							"structType": &dcl.Property{
								Type:        "object",
								GoName:      "StructType",
								GoType:      "RoutineArgumentsDataTypeStructType",
								Description: "The fields of this struct, in order, if type_kind = \"STRUCT\".",
								Immutable:   true,
								Conflicts: []string{
									"arrayElementType",
								},
								Properties: map[string]*dcl.Property{
									"fields": &dcl.Property{
										Type:      "array",
										GoName:    "Fields",
										Immutable: true,
										SendEmpty: true,
										ListType:  "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "RoutineArgumentsDataTypeStructTypeFields",
											Properties: map[string]*dcl.Property{
												"name": &dcl.Property{
													Type:        "string",
													GoName:      "Name",
													Description: "Optional. The name of this field. Can be absent for struct fields.",
													Immutable:   true,
												},
												"type": &dcl.Property{
													Ref:       "#/components/schemas/ArgumentsDataType",
													GoName:    "Type",
													Immutable: true,
												},
											},
										},
									},
								},
							},
							"typeKind": &dcl.Property{
								Type:        "string",
								GoName:      "TypeKind",
								GoType:      "RoutineArgumentsDataTypeTypeKindEnum",
								Description: "Required. The top level type of this field. Can be any standard SQL data type (e.g., \"INT64\", \"DATE\", \"ARRAY\"). Possible values: TYPE_KIND_UNSPECIFIED, INT64, BOOL, FLOAT64, STRING, BYTES, TIMESTAMP, DATE, TIME, DATETIME, INTERVAL, GEOGRAPHY, NUMERIC, BIGNUMERIC, JSON, ARRAY, STRUCT",
								Immutable:   true,
								Enum: []string{
									"TYPE_KIND_UNSPECIFIED",
									"INT64",
									"BOOL",
									"FLOAT64",
									"STRING",
									"BYTES",
									"TIMESTAMP",
									"DATE",
									"TIME",
									"DATETIME",
									"INTERVAL",
									"GEOGRAPHY",
									"NUMERIC",
									"BIGNUMERIC",
									"JSON",
									"ARRAY",
									"STRUCT",
								},
							},
						},
					},
				},
				"Routine": &dcl.Component{
					Title:           "Routine",
					ID:              "projects/{{project}}/datasets/{{dataset}}/routines/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"dataset",
							"routineType",
							"definitionBody",
						},
						Properties: map[string]*dcl.Property{
							"arguments": &dcl.Property{
								Type:        "array",
								GoName:      "Arguments",
								Description: "Optional.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "RoutineArguments",
									Properties: map[string]*dcl.Property{
										"argumentKind": &dcl.Property{
											Type:        "string",
											GoName:      "ArgumentKind",
											GoType:      "RoutineArgumentsArgumentKindEnum",
											Description: "Optional. Defaults to FIXED_TYPE. Possible values: ARGUMENT_KIND_UNSPECIFIED, FIXED_TYPE, ANY_TYPE",
											Enum: []string{
												"ARGUMENT_KIND_UNSPECIFIED",
												"FIXED_TYPE",
												"ANY_TYPE",
											},
										},
										"dataType": &dcl.Property{
											Ref:       "#/components/schemas/ArgumentsDataType",
											GoName:    "DataType",
											Immutable: true,
										},
										"mode": &dcl.Property{
											Type:        "string",
											GoName:      "Mode",
											GoType:      "RoutineArgumentsModeEnum",
											Description: "Optional. Specifies whether the argument is input or output. Can be set for procedures only. Possible values: MODE_UNSPECIFIED, IN, OUT, INOUT",
											Enum: []string{
												"MODE_UNSPECIFIED",
												"IN",
												"OUT",
												"INOUT",
											},
										},
										"name": &dcl.Property{
											Type:        "string",
											GoName:      "Name",
											Description: "Optional. The name of this argument. Can be absent for function return argument.",
										},
									},
								},
							},
							"creationTime": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "CreationTime",
								ReadOnly:    true,
								Description: "Output only. The time when this routine was created, in milliseconds since the epoch.",
								Immutable:   true,
							},
							"dataset": &dcl.Property{
								Type:        "string",
								GoName:      "Dataset",
								Description: "Required. The ID of the dataset containing this routine.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Bigquery/Dataset",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"definitionBody": &dcl.Property{
								Type:        "string",
								GoName:      "DefinitionBody",
								Description: "Required. The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring inside (but excluding) the parentheses. For example, for the function created with the following statement: `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, \"\n\", y))` The definition_body is `concat(x, \"\n\", y)` (\n is not replaced with linebreak). If language=JAVASCRIPT, it is the evaluated string in the AS clause. For example, for the function created with the following statement: `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return \"\n\";\n'` The definition_body is `return \"\n\";\n` Note that both \n are replaced with linebreaks.",
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. The description of the routine if defined.",
							},
							"determinismLevel": &dcl.Property{
								Type:        "string",
								GoName:      "DeterminismLevel",
								GoType:      "RoutineDeterminismLevelEnum",
								Description: "Optional. The determinism level of the JavaScript UDF if defined. Possible values: DETERMINISM_LEVEL_UNSPECIFIED, DETERMINISTIC, NOT_DETERMINISTIC",
								Enum: []string{
									"DETERMINISM_LEVEL_UNSPECIFIED",
									"DETERMINISTIC",
									"NOT_DETERMINISTIC",
								},
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "Output only. A hash of this resource.",
								Immutable:   true,
							},
							"importedLibraries": &dcl.Property{
								Type:        "array",
								GoName:      "ImportedLibraries",
								Description: "Optional. If language = \"JAVASCRIPT\", this field stores the path of the imported JAVASCRIPT libraries.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"language": &dcl.Property{
								Type:          "string",
								GoName:        "Language",
								GoType:        "RoutineLanguageEnum",
								Description:   "Optional. Defaults to \"SQL\". Possible values: LANGUAGE_UNSPECIFIED, SQL, JAVASCRIPT",
								ServerDefault: true,
								Enum: []string{
									"LANGUAGE_UNSPECIFIED",
									"SQL",
									"JAVASCRIPT",
								},
							},
							"lastModifiedTime": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "LastModifiedTime",
								ReadOnly:    true,
								Description: "Output only. The time when this routine was last modified, in milliseconds since the epoch.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Required. The ID of the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.",
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "Required. The ID of the project containing this routine.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"returnType": &dcl.Property{
								Ref:    "#/components/schemas/ArgumentsDataType",
								GoName: "ReturnType",
							},
							"routineType": &dcl.Property{
								Type:        "string",
								GoName:      "RoutineType",
								GoType:      "RoutineRoutineTypeEnum",
								Description: "Required. The type of routine. Possible values: ROUTINE_TYPE_UNSPECIFIED, SCALAR_FUNCTION, PROCEDURE",
								Enum: []string{
									"ROUTINE_TYPE_UNSPECIFIED",
									"SCALAR_FUNCTION",
									"PROCEDURE",
								},
							},
							"strictMode": &dcl.Property{
								Type:        "boolean",
								GoName:      "StrictMode",
								Description: "Optional. Can be set for procedures only. If true (default), the definition body will be validated in the creation and the updates of the procedure. For procedures with an argument of ANY TYPE, the definition body validtion is not supported at creation/update time, and thus this field must be set to false explicitly.",
							},
						},
					},
				},
			},
		},
	}
}
