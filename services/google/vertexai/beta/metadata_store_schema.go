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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLMetadataStoreSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "VertexAI/MetadataStore",
			Description: "The VertexAI MetadataStore resource",
			StructName:  "MetadataStore",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a MetadataStore",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "MetadataStore",
						Required:    true,
						Description: "A full instance of a MetadataStore",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a MetadataStore",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "MetadataStore",
						Required:    true,
						Description: "A full instance of a MetadataStore",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a MetadataStore",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "MetadataStore",
						Required:    true,
						Description: "A full instance of a MetadataStore",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all MetadataStore",
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
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many MetadataStore",
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
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"MetadataStore": &dcl.Component{
					Title:           "MetadataStore",
					ID:              "projects/{{project}}/locations/{{location}}/metadataStores/{{name}}",
					ParentContainer: "project",
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Timestamp when this MetadataStore was created.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Description of the MetadataStore.",
								Immutable:   true,
							},
							"encryptionSpec": &dcl.Property{
								Type:        "object",
								GoName:      "EncryptionSpec",
								GoType:      "MetadataStoreEncryptionSpec",
								Description: "Customer-managed encryption key spec for a Metadata Store. If set, this Metadata Store and all sub-resources of this Metadata Store are secured using this key.",
								Immutable:   true,
								Required: []string{
									"kmsKeyName",
								},
								Properties: map[string]*dcl.Property{
									"kmsKeyName": &dcl.Property{
										Type:        "string",
										GoName:      "KmsKeyName",
										Description: "Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. The key needs to be in the same region as where the compute resource is created.",
										Immutable:   true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Cloudkms/CryptoKey",
												Field:    "name",
											},
										},
									},
								},
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
								Description: "The resource name of the MetadataStore instance.",
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
							"state": &dcl.Property{
								Type:        "object",
								GoName:      "State",
								GoType:      "MetadataStoreState",
								ReadOnly:    true,
								Description: "Output only. State information of the MetadataStore.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"diskUtilizationBytes": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "DiskUtilizationBytes",
										Description: "The disk utilization of the MetadataStore in bytes.",
										Immutable:   true,
									},
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Timestamp when this MetadataStore was last updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
