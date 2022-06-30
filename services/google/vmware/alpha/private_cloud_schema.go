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

func DCLPrivateCloudSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Vmware/PrivateCloud",
			Description: "The Vmware PrivateCloud resource",
			StructName:  "PrivateCloud",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a PrivateCloud",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "PrivateCloud",
						Required:    true,
						Description: "A full instance of a PrivateCloud",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a PrivateCloud",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "PrivateCloud",
						Required:    true,
						Description: "A full instance of a PrivateCloud",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a PrivateCloud",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "PrivateCloud",
						Required:    true,
						Description: "A full instance of a PrivateCloud",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all PrivateCloud",
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
				Description: "The function used to list information about many PrivateCloud",
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
				"PrivateCloud": &dcl.Component{
					Title:           "PrivateCloud",
					ID:              "projects/{{project}}/locations/{{location}}/privateClouds/{{name}}",
					UsesStateHint:   true,
					ParentContainer: "project",
					HasCreate:       true,
					HasIAM:          true,
					ApplyTimeout:    9600,
					DeleteTimeout:   4800,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"networkConfig",
							"managementCluster",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"conditions": &dcl.Property{
								Type:        "array",
								GoName:      "Conditions",
								ReadOnly:    true,
								Description: "Output only. The conditions that caused the current private cloud state. For example, private cloud provisioning failure description.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "PrivateCloudConditions",
									Properties: map[string]*dcl.Property{
										"code": &dcl.Property{
											Type:        "string",
											GoName:      "Code",
											ReadOnly:    true,
											Description: "Output only. Machine-readable representation of the condition.",
											Immutable:   true,
										},
										"message": &dcl.Property{
											Type:        "string",
											GoName:      "Message",
											ReadOnly:    true,
											Description: "Output only. Human-readable description of the condition.",
											Immutable:   true,
										},
									},
								},
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Creation time of this resource in RFC3339 text format.",
								Immutable:   true,
							},
							"deleteTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "DeleteTime",
								ReadOnly:    true,
								Description: "Output only. Time the resource was marked as deleted, in RFC3339 text format.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "User-provided description for this private cloud.",
							},
							"expireTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "ExpireTime",
								ReadOnly:    true,
								Description: "Output only. Planned deletion time of this resource in RFC3339 text format.",
								Immutable:   true,
							},
							"hcx": &dcl.Property{
								Type:        "object",
								GoName:      "Hcx",
								GoType:      "PrivateCloudHcx",
								ReadOnly:    true,
								Description: "Output only. HCX appliance.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"externalIP": &dcl.Property{
										Type:        "string",
										GoName:      "ExternalIP",
										Description: "External IP address of the appliance.",
										Immutable:   true,
									},
									"fdqn": &dcl.Property{
										Type:        "string",
										GoName:      "Fdqn",
										Description: "Fully qualified domain name of the appliance.",
										Immutable:   true,
									},
									"internalIP": &dcl.Property{
										Type:        "string",
										GoName:      "InternalIP",
										Description: "Internal IP address of the appliance.",
										Immutable:   true,
									},
									"version": &dcl.Property{
										Type:        "string",
										GoName:      "Version",
										Description: "Version of the appliance.",
										Immutable:   true,
									},
								},
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"managementCluster": &dcl.Property{
								Type:        "object",
								GoName:      "ManagementCluster",
								GoType:      "PrivateCloudManagementCluster",
								Description: "Input only. The management cluster for this private cloud. This parameter is required during creation of private cloud to provide details for the default cluster.",
								Unreadable:  true,
								Required: []string{
									"clusterId",
									"nodeTypeId",
									"nodeCount",
								},
								Properties: map[string]*dcl.Property{
									"clusterId": &dcl.Property{
										Type:        "string",
										GoName:      "ClusterId",
										Description: "Required. The user-provided identifier of the new `Cluster`.",
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
										Immutable:   true,
									},
								},
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The resource name of this private cloud. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud`",
								Immutable:   true,
							},
							"networkConfig": &dcl.Property{
								Type:        "object",
								GoName:      "NetworkConfig",
								GoType:      "PrivateCloudNetworkConfig",
								Description: "Required. Network configuration.",
								Immutable:   true,
								Required: []string{
									"network",
									"managementCidr",
								},
								Properties: map[string]*dcl.Property{
									"managementCidr": &dcl.Property{
										Type:        "string",
										GoName:      "ManagementCidr",
										Description: "Required. Management CIDR used by VMWare management appliances.",
										Immutable:   true,
									},
									"network": &dcl.Property{
										Type:        "string",
										GoName:      "Network",
										Description: "Required. The relative resource name of the consumer VPC network this private cloud is attached to. Specify the name in the following form: `projects/{project}/global/networks/{network_id}` where `{project}` can either be a project number or a project ID.",
										Immutable:   true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Compute/Network",
												Field:    "name",
											},
										},
									},
									"serviceNetwork": &dcl.Property{
										Type:        "string",
										GoName:      "ServiceNetwork",
										ReadOnly:    true,
										Description: "Output only. The relative resource name of the service VPC network this private cloud is attached to. The name is specified in the following form: `projects/{service_project_number}/global/networks/{network_id}`.",
										Immutable:   true,
									},
								},
							},
							"nsx": &dcl.Property{
								Type:        "object",
								GoName:      "Nsx",
								GoType:      "PrivateCloudNsx",
								ReadOnly:    true,
								Description: "Output only. NSX appliance.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"externalIP": &dcl.Property{
										Type:        "string",
										GoName:      "ExternalIP",
										Description: "External IP address of the appliance.",
										Immutable:   true,
									},
									"fdqn": &dcl.Property{
										Type:        "string",
										GoName:      "Fdqn",
										Description: "Fully qualified domain name of the appliance.",
										Immutable:   true,
									},
									"internalIP": &dcl.Property{
										Type:        "string",
										GoName:      "InternalIP",
										Description: "Internal IP address of the appliance.",
										Immutable:   true,
									},
									"version": &dcl.Property{
										Type:        "string",
										GoName:      "Version",
										Description: "Version of the appliance.",
										Immutable:   true,
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
								GoType:      "PrivateCloudStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the resource. Possible values: ACTIVE, CREATING, UPDATING, FAILED, DELETED",
								Immutable:   true,
								Enum: []string{
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
							"vcenter": &dcl.Property{
								Type:        "object",
								GoName:      "Vcenter",
								GoType:      "PrivateCloudVcenter",
								ReadOnly:    true,
								Description: "Output only. Vcenter appliance.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"externalIP": &dcl.Property{
										Type:        "string",
										GoName:      "ExternalIP",
										Description: "External IP address of the appliance.",
										Immutable:   true,
									},
									"fdqn": &dcl.Property{
										Type:        "string",
										GoName:      "Fdqn",
										Description: "Fully qualified domain name of the appliance.",
										Immutable:   true,
									},
									"internalIP": &dcl.Property{
										Type:        "string",
										GoName:      "InternalIP",
										Description: "Internal IP address of the appliance.",
										Immutable:   true,
									},
									"version": &dcl.Property{
										Type:        "string",
										GoName:      "Version",
										Description: "Version of the appliance.",
										Immutable:   true,
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
