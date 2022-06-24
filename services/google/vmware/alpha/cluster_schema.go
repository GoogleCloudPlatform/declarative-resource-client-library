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

func DCLClusterSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Vmware/Cluster",
			Description: "The Vmware Cluster resource",
			StructName:  "Cluster",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Cluster",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Cluster",
						Required:    true,
						Description: "A full instance of a Cluster",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Cluster",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Cluster",
						Required:    true,
						Description: "A full instance of a Cluster",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Cluster",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Cluster",
						Required:    true,
						Description: "A full instance of a Cluster",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Cluster",
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
						Name:     "privatecloud",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Cluster",
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
						Name:     "privatecloud",
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
				"Cluster": &dcl.Component{
					Title:           "Cluster",
					ID:              "projects/{{project}}/locations/{{location}}/privateClouds/{{private_cloud}}/clusters/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"nodeTypeId",
							"nodeCount",
							"project",
							"location",
							"privateCloud",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Creation time of this resource in RFC3339 text format.",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"management": &dcl.Property{
								Type:        "boolean",
								GoName:      "Management",
								ReadOnly:    true,
								Description: "Output only. True if the cluster is a management cluster; false otherwise. There can only be one management cluster in a private cloud and it has to be the first one.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The resource name of this cluster. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud/clusters/my-cluster`",
								Immutable:   true,
							},
							"nodeCount": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "NodeCount",
								Description: "Required. Number of nodes in this cluster.",
							},
							"nodeTypeId": &dcl.Property{
								Type:        "string",
								GoName:      "NodeTypeId",
								Description: "Required. The canonical identifier of node types (`NodeType`) in this cluster. For example: standard-72.",
							},
							"privateCloud": &dcl.Property{
								Type:        "string",
								GoName:      "PrivateCloud",
								Description: "The privateCloud for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Vmware/PrivateCloud",
										Field:    "name",
										Parent:   true,
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
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "ClusterStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the resource. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, UPDATING, FAILED, DELETED",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"ACTIVE",
									"CREATING",
									"UPDATING",
									"FAILED",
									"DELETED",
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Last update time of this resource in RFC3339 text format.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
