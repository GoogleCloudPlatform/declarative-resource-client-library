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
package vertexai

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLEndpointTrafficSplitSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "VertexAI/EndpointTrafficSplit",
			Description: "The VertexAI EndpointTrafficSplit resource",
			StructName:  "EndpointTrafficSplit",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a EndpointTrafficSplit",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "EndpointTrafficSplit",
						Required:    true,
						Description: "A full instance of a EndpointTrafficSplit",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a EndpointTrafficSplit",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "EndpointTrafficSplit",
						Required:    true,
						Description: "A full instance of a EndpointTrafficSplit",
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"EndpointTrafficSplit": &dcl.Component{
					Title:           "EndpointTrafficSplit",
					ID:              "projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}",
					ParentContainer: "project",
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"endpoint",
							"project",
							"location",
							"trafficSplit",
						},
						Properties: map[string]*dcl.Property{
							"endpoint": &dcl.Property{
								Type:        "string",
								GoName:      "Endpoint",
								Description: "The endpoint for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Vertexai/Endpoint",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "Used to perform consistent read-modify-write updates. If not set, a blind \"overwrite\" update happens.",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
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
							"trafficSplit": &dcl.Property{
								Type:        "array",
								GoName:      "TrafficSplit",
								Description: "A map from a DeployedModel's ID to the percentage of this Endpoint's traffic that should be forwarded to that DeployedModel. If a DeployedModel's ID is not listed in this map, then it receives no traffic. The traffic percentage values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "EndpointTrafficSplitTrafficSplit",
									Required: []string{
										"deployedModelId",
										"trafficPercentage",
									},
									Properties: map[string]*dcl.Property{
										"deployedModelId": &dcl.Property{
											Type:        "string",
											GoName:      "DeployedModelId",
											Description: "A deployed model's id.",
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Vertexai/ModelDeployment",
													Field:    "id",
												},
											},
										},
										"trafficPercentage": &dcl.Property{
											Type:        "integer",
											Format:      "int64",
											GoName:      "TrafficPercentage",
											Description: "The percentage of this Endpoint's traffic that should be forwarded to the DeployedModel.",
										},
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
