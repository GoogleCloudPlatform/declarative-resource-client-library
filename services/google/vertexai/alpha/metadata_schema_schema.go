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

func DCLMetadataSchemaSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "VertexAI/MetadataSchema",
			Description: "The VertexAI MetadataSchema resource",
			StructName:  "MetadataSchema",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a MetadataSchema",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "MetadataSchema",
						Required:    true,
						Description: "A full instance of a MetadataSchema",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a MetadataSchema",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "MetadataSchema",
						Required:    true,
						Description: "A full instance of a MetadataSchema",
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many MetadataSchema",
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
						Name:     "metadatastore",
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
				"MetadataSchema": &dcl.Component{
					Title:           "MetadataSchema",
					ID:              "projects/{{project}}/locations/{{location}}/metadataStores/{{metadata_store}}/metadataSchemas/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"schemaVersion",
							"schema",
							"schemaType",
							"project",
							"location",
							"metadataStore",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Timestamp when this MetadataSchema was created.",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"metadataStore": &dcl.Property{
								Type:        "string",
								GoName:      "MetadataStore",
								Description: "The metadata store for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Vertex/MetadataStore",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Output only. The resource name of the MetadataSchema.",
								Immutable:   true,
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
							"schema": &dcl.Property{
								Type:        "string",
								GoName:      "Schema",
								Description: "Required. The raw YAML string representation of the MetadataSchema. The combination of [MetadataSchema.version] and the schema name given by `title` in [MetadataSchema.schema] must be unique within a MetadataStore. The schema is defined as an OpenAPI 3.0.2 [MetadataSchema Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schemaObject)",
								Immutable:   true,
							},
							"schemaType": &dcl.Property{
								Type:        "string",
								GoName:      "SchemaType",
								GoType:      "MetadataSchemaSchemaTypeEnum",
								Description: "The type of the MetadataSchema. This is a property that identifies which metadata types will use the MetadataSchema. Possible values: METADATA_SCHEMA_TYPE_UNSPECIFIED, ARTIFACT_TYPE, EXECUTION_TYPE, CONTEXT_TYPE",
								Immutable:   true,
								Enum: []string{
									"METADATA_SCHEMA_TYPE_UNSPECIFIED",
									"ARTIFACT_TYPE",
									"EXECUTION_TYPE",
									"CONTEXT_TYPE",
								},
							},
							"schemaVersion": &dcl.Property{
								Type:        "string",
								GoName:      "SchemaVersion",
								Description: "The version of the MetadataSchema. The version's format must match the following regular expression: `^[0-9]+.+.+$`, which would allow to order/compare different versions. Example: 1.0.0, 1.0.1, etc.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
