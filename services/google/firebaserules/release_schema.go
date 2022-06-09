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
package firebaserules

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLReleaseSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Firebaserules/Release",
			Description: "The Firebaserules Release resource",
			StructName:  "Release",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Release",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Release",
						Required:    true,
						Description: "A full instance of a Release",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Release",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Release",
						Required:    true,
						Description: "A full instance of a Release",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Release",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Release",
						Required:    true,
						Description: "A full instance of a Release",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Release",
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
				Description: "The function used to list information about many Release",
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
				"Release": &dcl.Component{
					Title:           "Release",
					ID:              "projects/{{project}}/releases/{{name}}",
					ParentContainer: "project",
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"rulesetName",
							"project",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Time the release was created.",
								Immutable:   true,
							},
							"disabled": &dcl.Property{
								Type:        "boolean",
								GoName:      "Disabled",
								ReadOnly:    true,
								Description: "Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated from this Release.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                "string",
								GoName:              "Name",
								Description:         "Format: `projects/{project_id}/releases/{release_id}`",
								Immutable:           true,
								ForwardSlashAllowed: true,
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
							"rulesetName": &dcl.Property{
								Type:        "string",
								GoName:      "RulesetName",
								Description: "Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Firebaserules/Ruleset",
										Field:    "name",
									},
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Time the release was updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
