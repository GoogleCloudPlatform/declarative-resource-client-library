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

func DCLManagedServiceSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Servicemanagement/ManagedService",
			Description: "The Servicemanagement ManagedService resource",
			StructName:  "ManagedService",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a ManagedService",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "managedService",
						Required:    true,
						Description: "A full instance of a ManagedService",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a ManagedService",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "managedService",
						Required:    true,
						Description: "A full instance of a ManagedService",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a ManagedService",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "managedService",
						Required:    true,
						Description: "A full instance of a ManagedService",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all ManagedService",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many ManagedService",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
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
				"ManagedService": &dcl.Component{
					Title:           "ManagedService",
					ID:              "services/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
						},
						Properties: map[string]*dcl.Property{
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The name of the service. See the overview for naming requirements.",
								Immutable:   true,
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "ID of the project that produces and owns this service.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
