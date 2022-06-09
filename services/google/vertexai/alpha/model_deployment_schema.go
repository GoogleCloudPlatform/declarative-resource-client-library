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

func DCLModelDeploymentSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "VertexAI/ModelDeployment",
			Description: "The VertexAI ModelDeployment resource",
			StructName:  "ModelDeployment",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a ModelDeployment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "ModelDeployment",
						Required:    true,
						Description: "A full instance of a ModelDeployment",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a ModelDeployment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "ModelDeployment",
						Required:    true,
						Description: "A full instance of a ModelDeployment",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a ModelDeployment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "ModelDeployment",
						Required:    true,
						Description: "A full instance of a ModelDeployment",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all ModelDeployment",
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
						Name:     "endpoint",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many ModelDeployment",
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
						Name:     "endpoint",
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
				"ModelDeployment": &dcl.Component{
					Title:           "ModelDeployment",
					ID:              "models/{{model}}",
					ParentContainer: "project",
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"model",
							"dedicatedResources",
							"endpoint",
						},
						Properties: map[string]*dcl.Property{
							"dedicatedResources": &dcl.Property{
								Type:        "object",
								GoName:      "DedicatedResources",
								GoType:      "ModelDeploymentDedicatedResources",
								Description: "A description of resources that are dedicated to the DeployedModel, and that need a higher degree of manual configuration.",
								Immutable:   true,
								Required: []string{
									"machineSpec",
									"minReplicaCount",
								},
								Properties: map[string]*dcl.Property{
									"machineSpec": &dcl.Property{
										Type:        "object",
										GoName:      "MachineSpec",
										GoType:      "ModelDeploymentDedicatedResourcesMachineSpec",
										Description: "Required. Immutable. The specification of a single machine used by the prediction.",
										Immutable:   true,
										Required: []string{
											"machineType",
										},
										Properties: map[string]*dcl.Property{
											"machineType": &dcl.Property{
												Type:        "string",
												GoName:      "MachineType",
												Description: "Immutable. The type of the machine. See the [list of machine types supported for prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types) See the [list of machine types supported for custom training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types). For DeployedModel this field is optional, and the default value is `n1-standard-2`. For BatchPredictionJob or as part of WorkerPoolSpec this field is required. TODO(rsurowka): Try to better unify the required vs optional.",
												Immutable:   true,
											},
										},
									},
									"maxReplicaCount": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "MaxReplicaCount",
										Description: "Immutable. The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, will use min_replica_count as the default value. The value of this field impacts the charge against Vertex CPU and GPU quotas. Specifically, you will be charged for max_replica_count * number of cores in the selected machine type) and (max_replica_count * number of GPUs per replica in the selected machine type).",
										Immutable:   true,
									},
									"minReplicaCount": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "MinReplicaCount",
										Description: "Required. Immutable. The minimum number of machine replicas this DeployedModel will be always deployed on. This value must be greater than or equal to 1. If traffic against the DeployedModel increases, it may dynamically be deployed onto more replicas, and as traffic decreases, some of these extra replicas may be freed.",
										Immutable:   true,
									},
								},
							},
							"endpoint": &dcl.Property{
								Type:        "string",
								GoName:      "Endpoint",
								Description: "The name of the endpoint to deploy to",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Vertexai/Endpoint",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"id": &dcl.Property{
								Type:        "string",
								GoName:      "Id",
								ReadOnly:    true,
								Description: "The deployed ID of the model in the endpoint",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:           "string",
								GoName:         "Location",
								Description:    "The location of the endpoint",
								Immutable:      true,
								ExtractIfEmpty: true,
							},
							"model": &dcl.Property{
								Type:        "string",
								GoName:      "Model",
								Description: "The name of the model to deploy",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Vertexai/Model",
										Field:    "name",
									},
								},
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project of the endpoint",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
								ExtractIfEmpty: true,
							},
						},
					},
				},
			},
		},
	}
}
