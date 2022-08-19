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

func DCLDatasetSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Healthcare/Dataset",
			Description: "The Healthcare Dataset resource",
			StructName:  "Dataset",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Dataset",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Dataset",
						Required:    true,
						Description: "A full instance of a Dataset",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Dataset",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Dataset",
						Required:    true,
						Description: "A full instance of a Dataset",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Dataset",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Dataset",
						Required:    true,
						Description: "A full instance of a Dataset",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Dataset",
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
				Description: "The function used to list information about many Dataset",
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
				"Dataset": &dcl.Component{
					Title:           "Dataset",
					ID:              "projects/{{project}}/locations/{{location}}/datasets/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.",
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
							"timeZone": &dcl.Property{
								Type:        "string",
								GoName:      "TimeZone",
								Description: "The default timezone used by this dataset. Must be a either a valid IANA time zone name such as \"America/New_York\" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7v2 messages, where no explicit timezone is specified.",
							},
						},
					},
				},
			},
		},
	}
}
