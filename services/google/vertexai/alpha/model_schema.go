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

func DCLModelSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "VertexAI/Model",
			Description: "The VertexAI Model resource",
			StructName:  "Model",
			HasCreate:   true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Model",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Model",
						Required:    true,
						Description: "A full instance of a Model",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Model",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Model",
						Required:    true,
						Description: "A full instance of a Model",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Model",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Model",
						Required:    true,
						Description: "A full instance of a Model",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Model",
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
				Description: "The function used to list information about many Model",
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
				"Model": &dcl.Component{
					Title:           "Model",
					ID:              "projects/{{project}}/locations/{{location}}/models/{{name}}",
					UsesStateHint:   true,
					ParentContainer: "project",
					LabelsField:     "labels",
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"displayName",
							"containerSpec",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"artifactUri": &dcl.Property{
								Type:        "string",
								GoName:      "ArtifactUri",
								Description: "Immutable. The path to the directory containing the Model artifact and any of its supporting files. Not present for AutoML Models.",
								Immutable:   true,
							},
							"containerSpec": &dcl.Property{
								Type:        "object",
								GoName:      "ContainerSpec",
								GoType:      "ModelContainerSpec",
								Description: "Input only. The specification of the container that is to be used when deploying this Model. The specification is ingested upon ModelService.UploadModel, and all binaries it contains are copied and stored internally by Vertex AI. Not present for AutoML Models.",
								Immutable:   true,
								Unreadable:  true,
								Required: []string{
									"imageUri",
								},
								Properties: map[string]*dcl.Property{
									"acceleratorRequirements": &dcl.Property{
										Type:        "array",
										GoName:      "AcceleratorRequirements",
										Description: "Immutable. Accelerators required to run the container. This changes how containers are started. For example, GPU requirements are used to set the `resources` field in the Kubernetes [Container v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core) spec, as described on the [Schedule GPUs](https://kubernetes.io/docs/tasks/manage-gpus/scheduling-gpus/) page of the Kubernetes documentation. Currently, this field is only used when deploying to EdgeDevices.",
										Immutable:   true,
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "ModelContainerSpecAcceleratorRequirements",
											Properties: map[string]*dcl.Property{
												"count": &dcl.Property{
													Type:        "integer",
													Format:      "int64",
													GoName:      "Count",
													Description: "Number of accelerators of the specified type needed to run the container.",
													Immutable:   true,
												},
												"type": &dcl.Property{
													Type:        "string",
													GoName:      "Type",
													GoType:      "ModelContainerSpecAcceleratorRequirementsTypeEnum",
													Description: "Type of the accelerator needed to run the container. Possible values: ACCELERATOR_TYPE_UNSPECIFIED, CORAL_EDGE_TPU, NVIDIA_GPU, AMD_GPU",
													Immutable:   true,
													Enum: []string{
														"ACCELERATOR_TYPE_UNSPECIFIED",
														"CORAL_EDGE_TPU",
														"NVIDIA_GPU",
														"AMD_GPU",
													},
												},
											},
										},
									},
									"args": &dcl.Property{
										Type:        "array",
										GoName:      "Args",
										Description: "Immutable. Specifies arguments for the command that runs when the container starts. This overrides the container's [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd). Specify this field as an array of executable and arguments, similar to a Docker `CMD`'s \"default parameters\" form. If you don't specify this field but do specify the command field, then the command from the `command` field runs without any additional arguments. See the [Kubernetes documentation about how the `command` and `args` fields interact with a container's `ENTRYPOINT` and `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes). If you don't specify this field and don't specify the `command` field, then the container's [`ENTRYPOINT`](https://docs.docker.com/engine/reference/builder/#cmd) and `CMD` determine what runs based on their default behavior. See the Docker documentation about [how `CMD` and `ENTRYPOINT` interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). In this field, you can reference [environment variables set by Vertex AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables) and environment variables set in the env field. You cannot reference environment variables set in the Docker image. In order for environment variables to be expanded, reference them by using the following syntax: `$(VARIABLE_NAME)` Note that this differs from Bash variable expansion, which does not use parentheses. If a variable cannot be resolved, the reference in the input string is used unchanged. To avoid variable expansion, you can escape this syntax with `$$`; for example: `$$(VARIABLE_NAME)` This field corresponds to the `args` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
										Immutable:   true,
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
									"command": &dcl.Property{
										Type:        "array",
										GoName:      "Command",
										Description: "Immutable. Specifies the command that runs when the container starts. This overrides the container's [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint). Specify this field as an array of executable and arguments, similar to a Docker `ENTRYPOINT`'s \"exec\" form, not its \"shell\" form. If you do not specify this field, then the container's `ENTRYPOINT` runs, in conjunction with the args field or the container's [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd), if either exists. If this field is not specified and the container does not have an `ENTRYPOINT`, then refer to the Docker documentation about [how `CMD` and `ENTRYPOINT` interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). If you specify this field, then you can also specify the `args` field to provide additional arguments for this command. However, if you specify this field, then the container's `CMD` is ignored. See the [Kubernetes documentation about how the `command` and `args` fields interact with a container's `ENTRYPOINT` and `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes). In this field, you can reference [environment variables set by Vertex AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables) and environment variables set in the env field. You cannot reference environment variables set in the Docker image. In order for environment variables to be expanded, reference them by using the following syntax: `$(VARIABLE_NAME)` Note that this differs from Bash variable expansion, which does not use parentheses. If a variable cannot be resolved, the reference in the input string is used unchanged. To avoid variable expansion, you can escape this syntax with `$$`; for example: `$$(VARIABLE_NAME)` This field corresponds to the `command` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
										Immutable:   true,
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
									"env": &dcl.Property{
										Type:        "array",
										GoName:      "Env",
										Description: "Immutable. List of environment variables to set in the container. After the container starts running, code running in the container can read these environment variables. Additionally, the command and args fields can reference these variables. Later entries in this list can also reference earlier entries. For example, the following example sets the variable `VAR_2` to have the value `foo bar`: ```json [ { \"name\": \"VAR_1\", \"value\": \"foo\" }, { \"name\": \"VAR_2\", \"value\": \"$(VAR_1) bar\" } ] ``` If you switch the order of the variables in the example, then the expansion does not occur. This field corresponds to the `env` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
										Immutable:   true,
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "ModelContainerSpecEnv",
											Required: []string{
												"name",
												"value",
											},
											Properties: map[string]*dcl.Property{
												"name": &dcl.Property{
													Type:        "string",
													GoName:      "Name",
													Description: "Required. Name of the environment variable. Must be a valid C identifier.",
													Immutable:   true,
												},
												"value": &dcl.Property{
													Type:        "string",
													GoName:      "Value",
													Description: "Required. Variables that reference a $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not.",
													Immutable:   true,
												},
											},
										},
									},
									"healthRoute": &dcl.Property{
										Type:        "string",
										GoName:      "HealthRoute",
										Description: "Immutable. HTTP path on the container to send health checks to. Vertex AI intermittently sends GET requests to this path on the container's IP address and port to check that the container is healthy. Read more about [health checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#health). For example, if you set this field to `/bar`, then Vertex AI intermittently sends a GET request to the `/bar` path on the port of your container specified by the first value of this `ModelContainerSpec`'s ports field. If you don't specify this field, it defaults to the following value when you deploy this Model to an Endpoint: `/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict` The placeholders in this value are replaced as follows: * ENDPOINT: The last segment (following `endpoints/`)of the Endpoint.name][] field of the Endpoint where this Model has been deployed. (Vertex AI makes this value available to your container code as the [`AIP_ENDPOINT_ID` environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).) * DEPLOYED_MODEL: DeployedModel.id of the `DeployedModel`. (Vertex AI makes this value available to your container code as the [`AIP_DEPLOYED_MODEL_ID` environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)",
										Immutable:   true,
									},
									"imageUri": &dcl.Property{
										Type:        "string",
										GoName:      "ImageUri",
										Description: "Required. Immutable. URI of the Docker image to be used as the custom container for serving predictions. This URI must identify an image in Artifact Registry or Container Registry. Learn more about the [container publishing requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#publishing), including permissions requirements for the Vertex AI Service Agent. The container image is ingested upon ModelService.UploadModel, stored internally, and this original path is afterwards not used. To learn about the requirements for the Docker image itself, see [Custom container requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#). You can use the URI to one of Vertex AI's [pre-built container images for prediction](https://cloud.google.com/vertex-ai/docs/predictions/pre-built-containers) in this field.",
										Immutable:   true,
									},
									"ports": &dcl.Property{
										Type:        "array",
										GoName:      "Ports",
										Description: "Immutable. List of ports to expose from the container. Vertex AI sends any prediction requests that it receives to the first port on this list. Vertex AI also sends [liveness and health checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#liveness) to this port. If you do not specify this field, it defaults to following value: ```json [ { \"containerPort\": 8080 } ] ``` Vertex AI does not use ports other than the first one listed. This field corresponds to the `ports` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
										Immutable:   true,
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "ModelContainerSpecPorts",
											Properties: map[string]*dcl.Property{
												"containerPort": &dcl.Property{
													Type:        "integer",
													Format:      "int64",
													GoName:      "ContainerPort",
													Description: "The number of the port to expose on the pod's IP address. Must be a valid port number, between 1 and 65535 inclusive.",
													Immutable:   true,
												},
											},
										},
									},
									"predictRoute": &dcl.Property{
										Type:        "string",
										GoName:      "PredictRoute",
										Description: "Immutable. HTTP path on the container to send prediction requests to. Vertex AI forwards requests sent using projects.locations.endpoints.predict to this path on the container's IP address and port. Vertex AI then returns the container's response in the API response. For example, if you set this field to `/foo`, then when Vertex AI receives a prediction request, it forwards the request body in a POST request to the `/foo` path on the port of your container specified by the first value of this `ModelContainerSpec`'s ports field. If you don't specify this field, it defaults to the following value when you deploy this Model to an Endpoint: `/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict` The placeholders in this value are replaced as follows: * ENDPOINT: The last segment (following `endpoints/`)of the Endpoint.name][] field of the Endpoint where this Model has been deployed. (Vertex AI makes this value available to your container code as the [`AIP_ENDPOINT_ID` environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).) * DEPLOYED_MODEL: DeployedModel.id of the `DeployedModel`. (Vertex AI makes this value available to your container code as the [`AIP_DEPLOYED_MODEL_ID` environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)",
										Immutable:   true,
									},
								},
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Timestamp when this Model was uploaded into Vertex AI.",
								Immutable:   true,
							},
							"deployedModels": &dcl.Property{
								Type:        "array",
								GoName:      "DeployedModels",
								ReadOnly:    true,
								Description: "Output only. The pointers to DeployedModels created from this Model. Note that Model could have been deployed to Endpoints in different Locations.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "ModelDeployedModels",
									Properties: map[string]*dcl.Property{
										"deployedModelId": &dcl.Property{
											Type:        "string",
											GoName:      "DeployedModelId",
											Description: "Immutable. An ID of a DeployedModel in the above Endpoint.",
											Immutable:   true,
										},
										"endpoint": &dcl.Property{
											Type:        "string",
											GoName:      "Endpoint",
											Description: "Immutable. A resource name of an Endpoint.",
											Immutable:   true,
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Aiplatform/Endpoint",
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
								Description: "The description of the Model.",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Required. The display name of the Model. The name can be up to 128 characters long and can be consist of any UTF-8 characters.",
							},
							"encryptionSpec": &dcl.Property{
								Type:        "object",
								GoName:      "EncryptionSpec",
								GoType:      "ModelEncryptionSpec",
								Description: "Customer-managed encryption key spec for a Model. If set, this Model and all sub-resources of this Model will be secured by this key.",
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
								Description: "The labels with user-defined metadata to organize your Models. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "The resource name of the Model.",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"originalModelInfo": &dcl.Property{
								Type:        "object",
								GoName:      "OriginalModelInfo",
								GoType:      "ModelOriginalModelInfo",
								ReadOnly:    true,
								Description: "Output only. If this Model is a copy of another Model, this contains info about the original.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"model": &dcl.Property{
										Type:        "string",
										GoName:      "Model",
										ReadOnly:    true,
										Description: "Output only. The resource name of the Model this Model is a copy of, including the revision. Format: `projects/{project}/locations/{location}/models/{model_id}@{version_id}`",
										Immutable:   true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Aiplatform/Model",
												Field:    "selfLink",
											},
										},
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
							"supportedDeploymentResourcesTypes": &dcl.Property{
								Type:        "array",
								GoName:      "SupportedDeploymentResourcesTypes",
								ReadOnly:    true,
								Description: "Output only. When this Model is deployed, its prediction resources are described by the `prediction_resources` field of the Endpoint.deployed_models object. Because not all Models support all resource configuration types, the configuration types this Model supports are listed here. If no configuration types are listed, the Model cannot be deployed to an Endpoint and does not support online predictions (PredictionService.Predict or PredictionService.Explain). Such a Model can serve predictions by using a BatchPredictionJob, if it has at least one entry each in supported_input_storage_formats and supported_output_storage_formats.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "ModelSupportedDeploymentResourcesTypesEnum",
									Enum: []string{
										"DEPLOYMENT_RESOURCES_TYPE_UNSPECIFIED",
										"DEDICATED_RESOURCES",
										"AUTOMATIC_RESOURCES",
									},
								},
							},
							"supportedExportFormats": &dcl.Property{
								Type:        "array",
								GoName:      "SupportedExportFormats",
								ReadOnly:    true,
								Description: "Output only. The formats in which this Model may be exported. If empty, this Model is not available for export.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "ModelSupportedExportFormats",
									Properties: map[string]*dcl.Property{
										"exportableContents": &dcl.Property{
											Type:        "array",
											GoName:      "ExportableContents",
											ReadOnly:    true,
											Description: "Output only. The content of this Model that may be exported.",
											Immutable:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "ModelSupportedExportFormatsExportableContentsEnum",
												Enum: []string{
													"EXPORTABLE_CONTENT_UNSPECIFIED",
													"ARTIFACT",
													"IMAGE",
												},
											},
										},
										"id": &dcl.Property{
											Type:        "string",
											GoName:      "Id",
											ReadOnly:    true,
											Description: "Output only. The ID of the export format. The possible format IDs are: * `tflite` Used for Android mobile devices. * `edgetpu-tflite` Used for [Edge TPU](https://cloud.google.com/edge-tpu/) devices. * `tf-saved-model` A tensorflow model in SavedModel format. * `tf-js` A [TensorFlow.js](https://www.tensorflow.org/js) model that can be used in the browser and in Node.js using JavaScript. * `core-ml` Used for iOS mobile devices. * `custom-trained` A Model that was uploaded or trained by custom code.",
											Immutable:   true,
										},
									},
								},
							},
							"supportedInputStorageFormats": &dcl.Property{
								Type:        "array",
								GoName:      "SupportedInputStorageFormats",
								ReadOnly:    true,
								Description: "Output only. The formats this Model supports in BatchPredictionJob.input_config. If PredictSchemata.instance_schema_uri exists, the instances should be given as per that schema. The possible formats are: * `jsonl` The JSON Lines format, where each instance is a single line. Uses GcsSource. * `csv` The CSV format, where each instance is a single comma-separated line. The first line in the file is the header, containing comma-separated field names. Uses GcsSource. * `tf-record` The TFRecord format, where each instance is a single record in tfrecord syntax. Uses GcsSource. * `tf-record-gzip` Similar to `tf-record`, but the file is gzipped. Uses GcsSource. * `bigquery` Each instance is a single row in BigQuery. Uses BigQuerySource. * `file-list` Each line of the file is the location of an instance to process, uses `gcs_source` field of the InputConfig object. If this Model doesn't support any of these formats it means it cannot be used with a BatchPredictionJob. However, if it has supported_deployment_resources_types, it could serve online predictions by using PredictionService.Predict or PredictionService.Explain. TODO(rsurowka): Give a link describing how OpenAPI schema instances are expressed in JSONL and BigQuery. TODO(rsurowka): Should we provide a schema for TFRecord? Or maybe say that at least for now TFRecord input is not supported via schemata (that would also simplify giving them back as part of predictions). TODO(rsurowka): Define CSV format (decide how much we want to support). E.g. no nesting? Or no arrays, or no nested arrays? E.g. https://json-csv.com/ seems to be able to do pretty advanced conversions, but we may decide to make it relatively simple for now.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"supportedOutputStorageFormats": &dcl.Property{
								Type:        "array",
								GoName:      "SupportedOutputStorageFormats",
								ReadOnly:    true,
								Description: "Output only. The formats this Model supports in BatchPredictionJob.output_config. If both PredictSchemata.instance_schema_uri and PredictSchemata.prediction_schema_uri exist, the predictions are returned together with their instances. In other words, the prediction has the original instance data first, followed by the actual prediction content (as per the schema). The possible formats are: * `jsonl` The JSON Lines format, where each prediction is a single line. Uses GcsDestination. * `csv` The CSV format, where each prediction is a single comma-separated line. The first line in the file is the header, containing comma-separated field names. Uses GcsDestination. * `bigquery` Each prediction is a single row in a BigQuery table, uses BigQueryDestination . If this Model doesn't support any of these formats it means it cannot be used with a BatchPredictionJob. However, if it has supported_deployment_resources_types, it could serve online predictions by using PredictionService.Predict or PredictionService.Explain. TODO(rsurowka): Analogous TODOs as for instances field above.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"trainingPipeline": &dcl.Property{
								Type:        "string",
								GoName:      "TrainingPipeline",
								ReadOnly:    true,
								Description: "Output only. The resource name of the TrainingPipeline that uploaded this Model, if any.",
								Immutable:   true,
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Timestamp when this Model was most recently updated.",
								Immutable:   true,
							},
							"versionAliases": &dcl.Property{
								Type:          "array",
								GoName:        "VersionAliases",
								Description:   "User provided version aliases so that a model version can be referenced via alias (i.e. projects/{project}/locations/{location}/models/{model_id}@{version_alias} instead of auto-generated version id (i.e. projects/{project}/locations/{location}/models/{model_id}@{version_id}). The format is a-z{0,126}[a-z0-9] to distinguish from version_id. A default version alias will be created for the first version of the model, and there must be exactly one default version alias for a model.",
								Immutable:     true,
								ServerDefault: true,
								SendEmpty:     true,
								ListType:      "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"versionCreateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "VersionCreateTime",
								ReadOnly:    true,
								Description: "Output only. Timestamp when this version was created.",
								Immutable:   true,
							},
							"versionDescription": &dcl.Property{
								Type:        "string",
								GoName:      "VersionDescription",
								Description: "The description of this version.",
								Immutable:   true,
							},
							"versionId": &dcl.Property{
								Type:        "string",
								GoName:      "VersionId",
								ReadOnly:    true,
								Description: "Output only. Immutable. The version ID of the model. A new version is committed when a new model version is uploaded or trained under an existing model id. It is an auto-incrementing decimal number in string representation.",
								Immutable:   true,
							},
							"versionUpdateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "VersionUpdateTime",
								ReadOnly:    true,
								Description: "Output only. Timestamp when this version was most recently updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
