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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLFhirStoreSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Healthcare/FhirStore",
			Description: "The Healthcare FhirStore resource",
			StructName:  "FhirStore",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a FhirStore",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "fhirStore",
						Required:    true,
						Description: "A full instance of a FhirStore",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a FhirStore",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "fhirStore",
						Required:    true,
						Description: "A full instance of a FhirStore",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a FhirStore",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "fhirStore",
						Required:    true,
						Description: "A full instance of a FhirStore",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all FhirStore",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
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
				Description: "The function used to list information about many FhirStore",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
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
				"FhirStore": &dcl.Component{
					Title:           "FhirStore",
					ID:              "projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/fhirStores/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"version",
							"project",
							"location",
							"dataset",
						},
						Properties: map[string]*dcl.Property{
							"complexDataTypeReferenceParsing": &dcl.Property{
								Type:        "string",
								GoName:      "ComplexDataTypeReferenceParsing",
								GoType:      "FhirStoreComplexDataTypeReferenceParsingEnum",
								Description: "Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources. Possible values: COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED, DISABLED, ENABLED",
								Enum: []string{
									"COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED",
									"DISABLED",
									"ENABLED",
								},
							},
							"dataset": &dcl.Property{
								Type:        "string",
								GoName:      "Dataset",
								Description: "The dataset for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Healthcare/Dataset",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"defaultSearchHandlingStrict": &dcl.Property{
								Type:        "boolean",
								GoName:      "DefaultSearchHandlingStrict",
								Description: "If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.",
							},
							"disableReferentialIntegrity": &dcl.Property{
								Type:        "boolean",
								GoName:      "DisableReferentialIntegrity",
								Description: "Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fail the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as Patient-everything, do not return all the results if broken references exist.",
								Immutable:   true,
							},
							"disableResourceVersioning": &dcl.Property{
								Type:        "boolean",
								GoName:      "DisableResourceVersioning",
								Description: "Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.",
								Immutable:   true,
							},
							"enableUpdateCreate": &dcl.Property{
								Type:        "boolean",
								GoName:      "EnableUpdateCreate",
								Description: "Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.",
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: p{Ll}p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [p{Ll}p{Lo}p{N}_-]{0,63} No more than 64 labels can be associated with a given store.",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Resource name of the FHIR store, of the form `projects/{project_id}/datasets/{dataset_id}/fhirStores/{fhir_store_id}`.",
								Immutable:   true,
							},
							"notificationConfig": &dcl.Property{
								Type:        "object",
								GoName:      "NotificationConfig",
								GoType:      "FhirStoreNotificationConfig",
								Description: "If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, \"action\":\"CreateResource\".",
								Properties: map[string]*dcl.Property{
									"pubsubTopic": &dcl.Property{
										Type:        "string",
										GoName:      "PubsubTopic",
										Description: "The [Pub/Sub](https://cloud.google.com/pubsub/docs/) topic that notifications of changes are published on. Supplied by the client. PubsubMessage.Data contains the resource name. PubsubMessage.MessageId is the ID of this message. It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message was published. Notifications are only sent if the topic is non-empty. [Topic names](https://cloud.google.com/pubsub/docs/overview#names) must be scoped to a project. Cloud Healthcare API service account must have publisher permissions on the given Pub/Sub topic. Not having adequate permissions causes the calls that send notifications to fail. If a notification can't be published to Pub/Sub, errors are logged to Cloud Logging (see [Viewing error logs in Cloud Logging](https://cloud.google.com/healthcare/docs/how-tos/logging)). If the number of errors exceeds a certain rate, some aren't submitted. Note that not all operations trigger notifications, see [Configuring Pub/Sub notifications](https://cloud.google.com/healthcare/docs/how-tos/pubsub) for specific details.",
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Pubsub/Topic",
												Field:    "name",
											},
										},
									},
								},
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"shardNum": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "ShardNum",
								Description: "Immutable. The number of shards to index/search for FHIR search. This is a trade-off between write and search performance. Larger shard number enables better parallelism, especially during bulk import, at the expense of potentially slower search due to more shards to scan. This is a workaround to the hotspot issue created by monotonically increasing search document ID. This value, if set, must be in the range of [0, 64]. Leaving it empty or setting it to 0 defaults to one shard. This field is immutable after store creation.",
								Immutable:   true,
							},
							"streamConfigs": &dcl.Property{
								Type:        "array",
								GoName:      "StreamConfigs",
								Description: "A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "FhirStoreStreamConfigs",
									Properties: map[string]*dcl.Property{
										"bigqueryDestination": &dcl.Property{
											Type:        "object",
											GoName:      "BigqueryDestination",
											GoType:      "FhirStoreStreamConfigsBigqueryDestination",
											Description: "The destination BigQuery structure that contains both the dataset location and corresponding schema config. The output is organized in one table per resource type. The server reuses the existing tables (if any) that are named after the resource types. For example, \"Patient\", \"Observation\". When there is no existing table for a given resource type, the server attempts to create one. When a table schema doesn't align with the schema config, either because of existing incompatible schema or out of band incompatible modification, the server does not stream in new data. BigQuery imposes a 1 MB limit on streaming insert row size, therefore any resource mutation that generates more than 1 MB of BigQuery data is not streamed. One resolution in this case is to delete the incompatible table and let the server recreate one, though the newly created table only contains data after the table recreation. Results are written to BigQuery tables according to the parameters in BigQueryDestination.WriteDisposition. Different versions of the same resource are distinguishable by the meta.versionId and meta.lastUpdated columns. The operation (CREATE/UPDATE/DELETE) that results in the new version is recorded in the meta.tag. The tables contain all historical resource versions since streaming was enabled. For query convenience, the server also creates one view per table of the same name containing only the current resource version. The streamed data in the BigQuery dataset is not guaranteed to be completely unique. The combination of the id and meta.versionId columns should ideally identify a single unique row. But in rare cases, duplicates may exist. At query time, users may use the SQL select statement to keep only one of the duplicate rows given an id and meta.versionId pair. Alternatively, the server created view mentioned above also filters out duplicates. If a resource mutation cannot be streamed to BigQuery, errors are logged to Cloud Logging. For more information, see [Viewing error logs in Cloud Logging](https://cloud.google.com/healthcare/docs/how-tos/logging)).",
											Properties: map[string]*dcl.Property{
												"datasetUri": &dcl.Property{
													Type:        "string",
													GoName:      "DatasetUri",
													Description: "BigQuery URI to an existing dataset, up to 2000 characters long, in the format `bq://projectId.bqDatasetId`.",
												},
												"schemaConfig": &dcl.Property{
													Type:        "object",
													GoName:      "SchemaConfig",
													GoType:      "FhirStoreStreamConfigsBigqueryDestinationSchemaConfig",
													Description: "The configuration for the exported BigQuery schema.",
													Properties: map[string]*dcl.Property{
														"recursiveStructureDepth": &dcl.Property{
															Type:          "integer",
															Format:        "int64",
															GoName:        "RecursiveStructureDepth",
															Description:   "The depth for all recursive structures in the output analytics schema. For example, `concept` in the CodeSystem resource is a recursive structure; when the depth is 2, the CodeSystem table will have a column called `concept.concept` but not `concept.concept.concept`. If not specified or set to 0, the server will use the default value 2. The maximum depth allowed is 5.",
															ServerDefault: true,
														},
														"schemaType": &dcl.Property{
															Type:        "string",
															GoName:      "SchemaType",
															GoType:      "FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum",
															Description: "Specifies the output schema type. Schema type is required. Possible values: SCHEMA_TYPE_UNSPECIFIED, LOSSLESS, ANALYTICS, ANALYTICS_V2",
															Enum: []string{
																"SCHEMA_TYPE_UNSPECIFIED",
																"LOSSLESS",
																"ANALYTICS",
																"ANALYTICS_V2",
															},
														},
													},
												},
											},
										},
										"resourceTypes": &dcl.Property{
											Type:        "array",
											GoName:      "ResourceTypes",
											Description: "Supply a FHIR resource type (such as \"Patient\" or \"Observation\"). See https://www.hl7.org/fhir/valueset-resource-types.html for a list of all FHIR resource types. The server treats an empty list as an intent to stream all the supported resource types in this FHIR store.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "string",
											},
										},
									},
								},
							},
							"validationConfig": &dcl.Property{
								Type:        "object",
								GoName:      "ValidationConfig",
								GoType:      "FhirStoreValidationConfig",
								Description: "Configuration for how to validate incoming FHIR resources against configured profiles.",
								Properties: map[string]*dcl.Property{
									"disableFhirpathValidation": &dcl.Property{
										Type:        "boolean",
										GoName:      "DisableFhirpathValidation",
										Description: "Whether to disable FHIRPath validation for incoming resources. Set this to true to disable checking incoming resources for conformance against FHIRPath requirement defined in the FHIR specification. This property only affects resource types that do not have profiles configured for them, any rules in enabled implementation guides will still be enforced.",
									},
									"disableProfileValidation": &dcl.Property{
										Type:        "boolean",
										GoName:      "DisableProfileValidation",
										Description: "Whether to disable profile validation for this FHIR store. Set this to true to disable checking incoming resources for conformance against structure definitions in this FHIR store.",
									},
									"disableReferenceTypeValidation": &dcl.Property{
										Type:        "boolean",
										GoName:      "DisableReferenceTypeValidation",
										Description: "Whether to disable reference type validation for incoming resources. Set this to true to disable checking incoming resources for conformance against reference type requirement defined in the FHIR specification. This property only affects resource types that do not have profiles configured for them, any rules in enabled implementation guides will still be enforced.",
									},
									"disableRequiredFieldValidation": &dcl.Property{
										Type:        "boolean",
										GoName:      "DisableRequiredFieldValidation",
										Description: "Whether to disable required fields validation for incoming resources. Set this to true to disable checking incoming resources for conformance against required fields requirement defined in the FHIR specification. This property only affects resource types that do not have profiles configured for them, any rules in enabled implementation guides will still be enforced.",
									},
									"enabledImplementationGuides": &dcl.Property{
										Type:        "array",
										GoName:      "EnabledImplementationGuides",
										Description: "A list of implementation guide URLs in this FHIR store that are used to configure the profiles to use for validation. For example, to use the US Core profiles for validation, set `enabled_implementation_guides` to `[\"http://hl7.org/fhir/us/core/ImplementationGuide/ig\"]`. If `enabled_implementation_guides` is empty or omitted, then incoming resources are only required to conform to the base FHIR profiles. Otherwise, a resource must conform to at least one profile listed in the `global` property of one of the enabled ImplementationGuides. The Cloud Healthcare API does not currently enforce all of the rules in a StructureDefinition. The following rules are supported: - min/max - minValue/maxValue - maxLength - type - fixed[x] - pattern[x] on simple types - slicing, when using \"value\" as the discriminator type When a URL cannot be resolved (for example, in a type assertion), the server does not return an error.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
								},
							},
							"version": &dcl.Property{
								Type:        "string",
								GoName:      "Version",
								GoType:      "FhirStoreVersionEnum",
								Description: "Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store. Possible values: VERSION_UNSPECIFIED, DSTU2, STU3, R4",
								Immutable:   true,
								Enum: []string{
									"VERSION_UNSPECIFIED",
									"DSTU2",
									"STU3",
									"R4",
								},
							},
						},
					},
				},
			},
		},
	}
}
