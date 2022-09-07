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

func DCLEndpointSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "VertexAI/Endpoint",
			Description: "The VertexAI Endpoint resource",
			StructName:  "Endpoint",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Endpoint",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "endpoint",
						Required:    true,
						Description: "A full instance of a Endpoint",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Endpoint",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "endpoint",
						Required:    true,
						Description: "A full instance of a Endpoint",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Endpoint",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "endpoint",
						Required:    true,
						Description: "A full instance of a Endpoint",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Endpoint",
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
				Description: "The function used to list information about many Endpoint",
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
				"Endpoint": &dcl.Component{
					Title:           "Endpoint",
					ID:              "projects/{{project}}/locations/{{location}}/endpoints/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"displayName",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Timestamp when this Endpoint was created.",
								Immutable:   true,
							},
							"deployedModels": &dcl.Property{
								Type:        "array",
								GoName:      "DeployedModels",
								ReadOnly:    true,
								Description: "Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "EndpointDeployedModels",
									Properties: map[string]*dcl.Property{
										"automaticResources": &dcl.Property{
											Type:        "object",
											GoName:      "AutomaticResources",
											GoType:      "EndpointDeployedModelsAutomaticResources",
											Description: "A description of resources that to large degree are decided by Vertex AI, and require only a modest additional configuration.",
											Immutable:   true,
											Conflicts: []string{
												"dedicatedResources",
											},
											Properties: map[string]*dcl.Property{
												"maxReplicaCount": &dcl.Property{
													Type:        "integer",
													Format:      "int64",
													GoName:      "MaxReplicaCount",
													Description: "The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, a no upper bound for scaling under heavy traffic will be assume, though Vertex AI may be unable to scale beyond certain replica number.",
													Immutable:   true,
												},
												"minReplicaCount": &dcl.Property{
													Type:        "integer",
													Format:      "int64",
													GoName:      "MinReplicaCount",
													Description: "The minimum number of replicas this DeployedModel will be always deployed on. If traffic against it increases, it may dynamically be deployed onto more replicas up to max_replica_count, and as traffic decreases, some of these extra replicas may be freed. If the requested value is too large, the deployment will error.",
													Immutable:   true,
												},
											},
										},
										"createTime": &dcl.Property{
											Type:        "string",
											Format:      "date-time",
											GoName:      "CreateTime",
											ReadOnly:    true,
											Description: "Output only. Timestamp when the DeployedModel was created.",
											Immutable:   true,
										},
										"dedicatedResources": &dcl.Property{
											Type:        "object",
											GoName:      "DedicatedResources",
											GoType:      "EndpointDeployedModelsDedicatedResources",
											Description: "A description of resources that are dedicated to the DeployedModel, and that need a higher degree of manual configuration.",
											Immutable:   true,
											Conflicts: []string{
												"automaticResources",
											},
											Properties: map[string]*dcl.Property{
												"autoscalingMetricSpecs": &dcl.Property{
													Type:        "array",
													GoName:      "AutoscalingMetricSpecs",
													Description: "The metric specifications that overrides a resource utilization metric (CPU utilization, accelerator's duty cycle, and so on) target value (default to 60 if not set). At most one entry is allowed per metric. If machine_spec.accelerator_count is above 0, the autoscaling will be based on both CPU utilization and accelerator's duty cycle metrics and scale up when either metrics exceeds its target value while scale down if both metrics are under their target value. The default target value is 60 for both metrics. If machine_spec.accelerator_count is 0, the autoscaling will be based on CPU utilization metric only with default target value 60 if not explicitly set. For example, in the case of Online Prediction, if you want to override target CPU utilization to 80, you should set autoscaling_metric_specs.metric_name to `aiplatform.googleapis.com/prediction/online/cpu/utilization` and autoscaling_metric_specs.target to `80`.",
													Immutable:   true,
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "object",
														GoType: "EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs",
														Properties: map[string]*dcl.Property{
															"metricName": &dcl.Property{
																Type:        "string",
																GoName:      "MetricName",
																Description: "The resource metric name. Supported metrics: * For Online Prediction: * `aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle` * `aiplatform.googleapis.com/prediction/online/cpu/utilization`",
																Immutable:   true,
															},
															"target": &dcl.Property{
																Type:        "integer",
																Format:      "int64",
																GoName:      "Target",
																Description: "The target resource utilization in percentage (1% - 100%) for the given metric; once the real usage deviates from the target by a certain percentage, the machine replicas change. The default value is 60 (representing 60%) if not provided.",
																Immutable:   true,
															},
														},
													},
												},
												"machineSpec": &dcl.Property{
													Type:        "object",
													GoName:      "MachineSpec",
													GoType:      "EndpointDeployedModelsDedicatedResourcesMachineSpec",
													Description: "The specification of a single machine used by the prediction.",
													Immutable:   true,
													Properties: map[string]*dcl.Property{
														"acceleratorCount": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "AcceleratorCount",
															Description: "The number of accelerators to attach to the machine.",
															Immutable:   true,
														},
														"acceleratorType": &dcl.Property{
															Type:        "string",
															GoName:      "AcceleratorType",
															GoType:      "EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum",
															Description: "The type of accelerator(s) that may be attached to the machine as per accelerator_count. Possible values: ACCELERATOR_TYPE_UNSPECIFIED, NVIDIA_TESLA_K80, NVIDIA_TESLA_P100, NVIDIA_TESLA_V100, NVIDIA_TESLA_P4, NVIDIA_TESLA_T4, NVIDIA_TESLA_A100, TPU_V2, TPU_V3",
															Immutable:   true,
															Enum: []string{
																"ACCELERATOR_TYPE_UNSPECIFIED",
																"NVIDIA_TESLA_K80",
																"NVIDIA_TESLA_P100",
																"NVIDIA_TESLA_V100",
																"NVIDIA_TESLA_P4",
																"NVIDIA_TESLA_T4",
																"NVIDIA_TESLA_A100",
																"TPU_V2",
																"TPU_V3",
															},
														},
														"machineType": &dcl.Property{
															Type:        "string",
															GoName:      "MachineType",
															Description: "The type of the machine. See the [list of machine types supported for prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types) See the [list of machine types supported for custom training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types). For DeployedModel this field is optional, and the default value is `n1-standard-2`. For BatchPredictionJob or as part of WorkerPoolSpec this field is required. TODO(rsurowka): Try to better unify the required vs optional.",
															Immutable:   true,
														},
													},
												},
												"maxReplicaCount": &dcl.Property{
													Type:        "integer",
													Format:      "int64",
													GoName:      "MaxReplicaCount",
													Description: "The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, will use min_replica_count as the default value. The value of this field impacts the charge against Vertex CPU and GPU quotas. Specifically, you will be charged for max_replica_count * number of cores in the selected machine type) and (max_replica_count * number of GPUs per replica in the selected machine type).",
													Immutable:   true,
												},
												"minReplicaCount": &dcl.Property{
													Type:        "integer",
													Format:      "int64",
													GoName:      "MinReplicaCount",
													Description: "The minimum number of machine replicas this DeployedModel will be always deployed on. This value must be greater than or equal to 1. If traffic against the DeployedModel increases, it may dynamically be deployed onto more replicas, and as traffic decreases, some of these extra replicas may be freed.",
													Immutable:   true,
												},
											},
										},
										"displayName": &dcl.Property{
											Type:        "string",
											GoName:      "DisplayName",
											Description: "The display name of the DeployedModel. If not provided upon creation, the Model's display_name is used.",
											Immutable:   true,
										},
										"enableAccessLogging": &dcl.Property{
											Type:        "boolean",
											GoName:      "EnableAccessLogging",
											Description: "These logs are like standard server access logs, containing information like timestamp and latency for each prediction request. Note that Stackdriver logs may incur a cost, especially if your project receives prediction requests at a high queries per second rate (QPS). Estimate your costs before enabling this option.",
											Immutable:   true,
										},
										"enableContainerLogging": &dcl.Property{
											Type:        "boolean",
											GoName:      "EnableContainerLogging",
											Description: "If true, the container of the DeployedModel instances will send `stderr` and `stdout` streams to Stackdriver Logging. Only supported for custom-trained Models and AutoML Tabular Models.",
											Immutable:   true,
										},
										"id": &dcl.Property{
											Type:        "string",
											GoName:      "Id",
											Description: "The ID of the DeployedModel. If not provided upon deployment, Vertex AI will generate a value for this ID. This value should be 1-10 characters, and valid characters are /[0-9]/.",
											Immutable:   true,
										},
										"model": &dcl.Property{
											Type:        "string",
											GoName:      "Model",
											Description: "The name of the Model that this is the deployment of. Note that the Model may be in a different location than the DeployedModel's Endpoint.",
											Immutable:   true,
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Aiplatform/Model",
													Field:    "selfLink",
												},
											},
										},
										"modelVersionId": &dcl.Property{
											Type:        "string",
											GoName:      "ModelVersionId",
											ReadOnly:    true,
											Description: "Output only. The version ID of the model that is deployed.",
											Immutable:   true,
										},
										"privateEndpoints": &dcl.Property{
											Type:        "object",
											GoName:      "PrivateEndpoints",
											GoType:      "EndpointDeployedModelsPrivateEndpoints",
											ReadOnly:    true,
											Description: "Output only. Provide paths for users to send predict/explain/health requests directly to the deployed model services running on Cloud via private services access. This field is populated if network is configured.",
											Immutable:   true,
											Properties: map[string]*dcl.Property{
												"explainHttpUri": &dcl.Property{
													Type:        "string",
													GoName:      "ExplainHttpUri",
													ReadOnly:    true,
													Description: "Output only. Http(s) path to send explain requests.",
													Immutable:   true,
												},
												"healthHttpUri": &dcl.Property{
													Type:        "string",
													GoName:      "HealthHttpUri",
													ReadOnly:    true,
													Description: "Output only. Http(s) path to send health check requests.",
													Immutable:   true,
												},
												"predictHttpUri": &dcl.Property{
													Type:        "string",
													GoName:      "PredictHttpUri",
													ReadOnly:    true,
													Description: "Output only. Http(s) path to send prediction requests.",
													Immutable:   true,
												},
												"serviceAttachment": &dcl.Property{
													Type:        "string",
													GoName:      "ServiceAttachment",
													ReadOnly:    true,
													Description: "Output only. The name of the service attachment resource. Populated if private service connect is enabled.",
													Immutable:   true,
												},
											},
										},
										"serviceAccount": &dcl.Property{
											Type:        "string",
											GoName:      "ServiceAccount",
											Description: "The service account that the DeployedModel's container runs as. Specify the email address of the service account. If this service account is not specified, the container runs as a service account that doesn't have access to the resource project. Users deploying the Model must have the `iam.serviceAccounts.actAs` permission on this service account.",
											Immutable:   true,
										},
										"sharedResources": &dcl.Property{
											Type:        "string",
											GoName:      "SharedResources",
											Description: "The resource name of the shared DeploymentResourcePool to deploy on. Format: projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}",
											Immutable:   true,
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Aiplatform/DeploymentResourcePool",
													Field:    "selfLink",
												},
											},
										},
									},
								},
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "The description of the Endpoint.",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Required. The display name of the Endpoint. The name can be up to 128 characters long and can be consist of any UTF-8 characters.",
							},
							"encryptionSpec": &dcl.Property{
								Type:        "object",
								GoName:      "EncryptionSpec",
								GoType:      "EndpointEncryptionSpec",
								Description: "Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.",
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
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "Used to perform consistent read-modify-write updates. If not set, a blind \"overwrite\" update happens.",
								Immutable:   true,
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"modelDeploymentMonitoringJob": &dcl.Property{
								Type:        "string",
								GoName:      "ModelDeploymentMonitoringJob",
								ReadOnly:    true,
								Description: "Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Output only. The resource name of the Endpoint.",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"network": &dcl.Property{
								Type:        "string",
								GoName:      "Network",
								Description: "The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Compute/Network",
										Field:    "selfLink",
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
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Timestamp when this Endpoint was last updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
