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
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Model struct{}

func ModelToUnstructured(r *dclService.Model) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "vertexai",
			Version: "alpha",
			Type:    "Model",
		},
		Object: make(map[string]interface{}),
	}
	if r.ArtifactUri != nil {
		u.Object["artifactUri"] = *r.ArtifactUri
	}
	if r.ContainerSpec != nil && r.ContainerSpec != dclService.EmptyModelContainerSpec {
		rContainerSpec := make(map[string]interface{})
		var rContainerSpecAcceleratorRequirements []interface{}
		for _, rContainerSpecAcceleratorRequirementsVal := range r.ContainerSpec.AcceleratorRequirements {
			rContainerSpecAcceleratorRequirementsObject := make(map[string]interface{})
			if rContainerSpecAcceleratorRequirementsVal.Count != nil {
				rContainerSpecAcceleratorRequirementsObject["count"] = *rContainerSpecAcceleratorRequirementsVal.Count
			}
			if rContainerSpecAcceleratorRequirementsVal.Type != nil {
				rContainerSpecAcceleratorRequirementsObject["type"] = string(*rContainerSpecAcceleratorRequirementsVal.Type)
			}
			rContainerSpecAcceleratorRequirements = append(rContainerSpecAcceleratorRequirements, rContainerSpecAcceleratorRequirementsObject)
		}
		rContainerSpec["acceleratorRequirements"] = rContainerSpecAcceleratorRequirements
		var rContainerSpecArgs []interface{}
		for _, rContainerSpecArgsVal := range r.ContainerSpec.Args {
			rContainerSpecArgs = append(rContainerSpecArgs, rContainerSpecArgsVal)
		}
		rContainerSpec["args"] = rContainerSpecArgs
		var rContainerSpecCommand []interface{}
		for _, rContainerSpecCommandVal := range r.ContainerSpec.Command {
			rContainerSpecCommand = append(rContainerSpecCommand, rContainerSpecCommandVal)
		}
		rContainerSpec["command"] = rContainerSpecCommand
		var rContainerSpecEnv []interface{}
		for _, rContainerSpecEnvVal := range r.ContainerSpec.Env {
			rContainerSpecEnvObject := make(map[string]interface{})
			if rContainerSpecEnvVal.Name != nil {
				rContainerSpecEnvObject["name"] = *rContainerSpecEnvVal.Name
			}
			if rContainerSpecEnvVal.Value != nil {
				rContainerSpecEnvObject["value"] = *rContainerSpecEnvVal.Value
			}
			rContainerSpecEnv = append(rContainerSpecEnv, rContainerSpecEnvObject)
		}
		rContainerSpec["env"] = rContainerSpecEnv
		if r.ContainerSpec.HealthRoute != nil {
			rContainerSpec["healthRoute"] = *r.ContainerSpec.HealthRoute
		}
		if r.ContainerSpec.ImageUri != nil {
			rContainerSpec["imageUri"] = *r.ContainerSpec.ImageUri
		}
		var rContainerSpecPorts []interface{}
		for _, rContainerSpecPortsVal := range r.ContainerSpec.Ports {
			rContainerSpecPortsObject := make(map[string]interface{})
			if rContainerSpecPortsVal.ContainerPort != nil {
				rContainerSpecPortsObject["containerPort"] = *rContainerSpecPortsVal.ContainerPort
			}
			rContainerSpecPorts = append(rContainerSpecPorts, rContainerSpecPortsObject)
		}
		rContainerSpec["ports"] = rContainerSpecPorts
		if r.ContainerSpec.PredictRoute != nil {
			rContainerSpec["predictRoute"] = *r.ContainerSpec.PredictRoute
		}
		u.Object["containerSpec"] = rContainerSpec
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	var rDeployedModels []interface{}
	for _, rDeployedModelsVal := range r.DeployedModels {
		rDeployedModelsObject := make(map[string]interface{})
		if rDeployedModelsVal.DeployedModelId != nil {
			rDeployedModelsObject["deployedModelId"] = *rDeployedModelsVal.DeployedModelId
		}
		if rDeployedModelsVal.Endpoint != nil {
			rDeployedModelsObject["endpoint"] = *rDeployedModelsVal.Endpoint
		}
		rDeployedModels = append(rDeployedModels, rDeployedModelsObject)
	}
	u.Object["deployedModels"] = rDeployedModels
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.EncryptionSpec != nil && r.EncryptionSpec != dclService.EmptyModelEncryptionSpec {
		rEncryptionSpec := make(map[string]interface{})
		if r.EncryptionSpec.KmsKeyName != nil {
			rEncryptionSpec["kmsKeyName"] = *r.EncryptionSpec.KmsKeyName
		}
		u.Object["encryptionSpec"] = rEncryptionSpec
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.OriginalModelInfo != nil && r.OriginalModelInfo != dclService.EmptyModelOriginalModelInfo {
		rOriginalModelInfo := make(map[string]interface{})
		if r.OriginalModelInfo.Model != nil {
			rOriginalModelInfo["model"] = *r.OriginalModelInfo.Model
		}
		u.Object["originalModelInfo"] = rOriginalModelInfo
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	var rSupportedDeploymentResourcesTypes []interface{}
	for _, rSupportedDeploymentResourcesTypesVal := range r.SupportedDeploymentResourcesTypes {
		rSupportedDeploymentResourcesTypes = append(rSupportedDeploymentResourcesTypes, string(rSupportedDeploymentResourcesTypesVal))
	}
	u.Object["supportedDeploymentResourcesTypes"] = rSupportedDeploymentResourcesTypes
	var rSupportedExportFormats []interface{}
	for _, rSupportedExportFormatsVal := range r.SupportedExportFormats {
		rSupportedExportFormatsObject := make(map[string]interface{})
		var rSupportedExportFormatsValExportableContents []interface{}
		for _, rSupportedExportFormatsValExportableContentsVal := range rSupportedExportFormatsVal.ExportableContents {
			rSupportedExportFormatsValExportableContents = append(rSupportedExportFormatsValExportableContents, string(rSupportedExportFormatsValExportableContentsVal))
		}
		rSupportedExportFormatsObject["exportableContents"] = rSupportedExportFormatsValExportableContents
		if rSupportedExportFormatsVal.Id != nil {
			rSupportedExportFormatsObject["id"] = *rSupportedExportFormatsVal.Id
		}
		rSupportedExportFormats = append(rSupportedExportFormats, rSupportedExportFormatsObject)
	}
	u.Object["supportedExportFormats"] = rSupportedExportFormats
	var rSupportedInputStorageFormats []interface{}
	for _, rSupportedInputStorageFormatsVal := range r.SupportedInputStorageFormats {
		rSupportedInputStorageFormats = append(rSupportedInputStorageFormats, rSupportedInputStorageFormatsVal)
	}
	u.Object["supportedInputStorageFormats"] = rSupportedInputStorageFormats
	var rSupportedOutputStorageFormats []interface{}
	for _, rSupportedOutputStorageFormatsVal := range r.SupportedOutputStorageFormats {
		rSupportedOutputStorageFormats = append(rSupportedOutputStorageFormats, rSupportedOutputStorageFormatsVal)
	}
	u.Object["supportedOutputStorageFormats"] = rSupportedOutputStorageFormats
	if r.TrainingPipeline != nil {
		u.Object["trainingPipeline"] = *r.TrainingPipeline
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	var rVersionAliases []interface{}
	for _, rVersionAliasesVal := range r.VersionAliases {
		rVersionAliases = append(rVersionAliases, rVersionAliasesVal)
	}
	u.Object["versionAliases"] = rVersionAliases
	if r.VersionCreateTime != nil {
		u.Object["versionCreateTime"] = *r.VersionCreateTime
	}
	if r.VersionDescription != nil {
		u.Object["versionDescription"] = *r.VersionDescription
	}
	if r.VersionId != nil {
		u.Object["versionId"] = *r.VersionId
	}
	if r.VersionUpdateTime != nil {
		u.Object["versionUpdateTime"] = *r.VersionUpdateTime
	}
	return u
}

func UnstructuredToModel(u *unstructured.Resource) (*dclService.Model, error) {
	r := &dclService.Model{}
	if _, ok := u.Object["artifactUri"]; ok {
		if s, ok := u.Object["artifactUri"].(string); ok {
			r.ArtifactUri = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ArtifactUri: expected string")
		}
	}
	if _, ok := u.Object["containerSpec"]; ok {
		if rContainerSpec, ok := u.Object["containerSpec"].(map[string]interface{}); ok {
			r.ContainerSpec = &dclService.ModelContainerSpec{}
			if _, ok := rContainerSpec["acceleratorRequirements"]; ok {
				if s, ok := rContainerSpec["acceleratorRequirements"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rContainerSpecAcceleratorRequirements dclService.ModelContainerSpecAcceleratorRequirements
							if _, ok := objval["count"]; ok {
								if i, ok := objval["count"].(int64); ok {
									rContainerSpecAcceleratorRequirements.Count = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rContainerSpecAcceleratorRequirements.Count: expected int64")
								}
							}
							if _, ok := objval["type"]; ok {
								if s, ok := objval["type"].(string); ok {
									rContainerSpecAcceleratorRequirements.Type = dclService.ModelContainerSpecAcceleratorRequirementsTypeEnumRef(s)
								} else {
									return nil, fmt.Errorf("rContainerSpecAcceleratorRequirements.Type: expected string")
								}
							}
							r.ContainerSpec.AcceleratorRequirements = append(r.ContainerSpec.AcceleratorRequirements, rContainerSpecAcceleratorRequirements)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ContainerSpec.AcceleratorRequirements: expected []interface{}")
				}
			}
			if _, ok := rContainerSpec["args"]; ok {
				if s, ok := rContainerSpec["args"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.ContainerSpec.Args = append(r.ContainerSpec.Args, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ContainerSpec.Args: expected []interface{}")
				}
			}
			if _, ok := rContainerSpec["command"]; ok {
				if s, ok := rContainerSpec["command"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.ContainerSpec.Command = append(r.ContainerSpec.Command, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ContainerSpec.Command: expected []interface{}")
				}
			}
			if _, ok := rContainerSpec["env"]; ok {
				if s, ok := rContainerSpec["env"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rContainerSpecEnv dclService.ModelContainerSpecEnv
							if _, ok := objval["name"]; ok {
								if s, ok := objval["name"].(string); ok {
									rContainerSpecEnv.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rContainerSpecEnv.Name: expected string")
								}
							}
							if _, ok := objval["value"]; ok {
								if s, ok := objval["value"].(string); ok {
									rContainerSpecEnv.Value = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rContainerSpecEnv.Value: expected string")
								}
							}
							r.ContainerSpec.Env = append(r.ContainerSpec.Env, rContainerSpecEnv)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ContainerSpec.Env: expected []interface{}")
				}
			}
			if _, ok := rContainerSpec["healthRoute"]; ok {
				if s, ok := rContainerSpec["healthRoute"].(string); ok {
					r.ContainerSpec.HealthRoute = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ContainerSpec.HealthRoute: expected string")
				}
			}
			if _, ok := rContainerSpec["imageUri"]; ok {
				if s, ok := rContainerSpec["imageUri"].(string); ok {
					r.ContainerSpec.ImageUri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ContainerSpec.ImageUri: expected string")
				}
			}
			if _, ok := rContainerSpec["ports"]; ok {
				if s, ok := rContainerSpec["ports"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rContainerSpecPorts dclService.ModelContainerSpecPorts
							if _, ok := objval["containerPort"]; ok {
								if i, ok := objval["containerPort"].(int64); ok {
									rContainerSpecPorts.ContainerPort = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rContainerSpecPorts.ContainerPort: expected int64")
								}
							}
							r.ContainerSpec.Ports = append(r.ContainerSpec.Ports, rContainerSpecPorts)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ContainerSpec.Ports: expected []interface{}")
				}
			}
			if _, ok := rContainerSpec["predictRoute"]; ok {
				if s, ok := rContainerSpec["predictRoute"].(string); ok {
					r.ContainerSpec.PredictRoute = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ContainerSpec.PredictRoute: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ContainerSpec: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["deployedModels"]; ok {
		if s, ok := u.Object["deployedModels"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rDeployedModels dclService.ModelDeployedModels
					if _, ok := objval["deployedModelId"]; ok {
						if s, ok := objval["deployedModelId"].(string); ok {
							rDeployedModels.DeployedModelId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDeployedModels.DeployedModelId: expected string")
						}
					}
					if _, ok := objval["endpoint"]; ok {
						if s, ok := objval["endpoint"].(string); ok {
							rDeployedModels.Endpoint = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDeployedModels.Endpoint: expected string")
						}
					}
					r.DeployedModels = append(r.DeployedModels, rDeployedModels)
				}
			}
		} else {
			return nil, fmt.Errorf("r.DeployedModels: expected []interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["encryptionSpec"]; ok {
		if rEncryptionSpec, ok := u.Object["encryptionSpec"].(map[string]interface{}); ok {
			r.EncryptionSpec = &dclService.ModelEncryptionSpec{}
			if _, ok := rEncryptionSpec["kmsKeyName"]; ok {
				if s, ok := rEncryptionSpec["kmsKeyName"].(string); ok {
					r.EncryptionSpec.KmsKeyName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.EncryptionSpec.KmsKeyName: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.EncryptionSpec: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["originalModelInfo"]; ok {
		if rOriginalModelInfo, ok := u.Object["originalModelInfo"].(map[string]interface{}); ok {
			r.OriginalModelInfo = &dclService.ModelOriginalModelInfo{}
			if _, ok := rOriginalModelInfo["model"]; ok {
				if s, ok := rOriginalModelInfo["model"].(string); ok {
					r.OriginalModelInfo.Model = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.OriginalModelInfo.Model: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.OriginalModelInfo: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["supportedDeploymentResourcesTypes"]; ok {
		if s, ok := u.Object["supportedDeploymentResourcesTypes"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.SupportedDeploymentResourcesTypes = append(r.SupportedDeploymentResourcesTypes, dclService.ModelSupportedDeploymentResourcesTypesEnum(strval))
				}
			}
		} else {
			return nil, fmt.Errorf("r.SupportedDeploymentResourcesTypes: expected []interface{}")
		}
	}
	if _, ok := u.Object["supportedExportFormats"]; ok {
		if s, ok := u.Object["supportedExportFormats"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rSupportedExportFormats dclService.ModelSupportedExportFormats
					if _, ok := objval["exportableContents"]; ok {
						if s, ok := objval["exportableContents"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rSupportedExportFormats.ExportableContents = append(rSupportedExportFormats.ExportableContents, dclService.ModelSupportedExportFormatsExportableContentsEnum(strval))
								}
							}
						} else {
							return nil, fmt.Errorf("rSupportedExportFormats.ExportableContents: expected []interface{}")
						}
					}
					if _, ok := objval["id"]; ok {
						if s, ok := objval["id"].(string); ok {
							rSupportedExportFormats.Id = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rSupportedExportFormats.Id: expected string")
						}
					}
					r.SupportedExportFormats = append(r.SupportedExportFormats, rSupportedExportFormats)
				}
			}
		} else {
			return nil, fmt.Errorf("r.SupportedExportFormats: expected []interface{}")
		}
	}
	if _, ok := u.Object["supportedInputStorageFormats"]; ok {
		if s, ok := u.Object["supportedInputStorageFormats"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.SupportedInputStorageFormats = append(r.SupportedInputStorageFormats, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.SupportedInputStorageFormats: expected []interface{}")
		}
	}
	if _, ok := u.Object["supportedOutputStorageFormats"]; ok {
		if s, ok := u.Object["supportedOutputStorageFormats"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.SupportedOutputStorageFormats = append(r.SupportedOutputStorageFormats, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.SupportedOutputStorageFormats: expected []interface{}")
		}
	}
	if _, ok := u.Object["trainingPipeline"]; ok {
		if s, ok := u.Object["trainingPipeline"].(string); ok {
			r.TrainingPipeline = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.TrainingPipeline: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["versionAliases"]; ok {
		if s, ok := u.Object["versionAliases"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.VersionAliases = append(r.VersionAliases, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.VersionAliases: expected []interface{}")
		}
	}
	if _, ok := u.Object["versionCreateTime"]; ok {
		if s, ok := u.Object["versionCreateTime"].(string); ok {
			r.VersionCreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.VersionCreateTime: expected string")
		}
	}
	if _, ok := u.Object["versionDescription"]; ok {
		if s, ok := u.Object["versionDescription"].(string); ok {
			r.VersionDescription = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.VersionDescription: expected string")
		}
	}
	if _, ok := u.Object["versionId"]; ok {
		if s, ok := u.Object["versionId"].(string); ok {
			r.VersionId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.VersionId: expected string")
		}
	}
	if _, ok := u.Object["versionUpdateTime"]; ok {
		if s, ok := u.Object["versionUpdateTime"].(string); ok {
			r.VersionUpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.VersionUpdateTime: expected string")
		}
	}
	return r, nil
}

func GetModel(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToModel(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetModel(ctx, r)
	if err != nil {
		return nil, err
	}
	return ModelToUnstructured(r), nil
}

func ListModel(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListModel(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ModelToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyModel(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToModel(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToModel(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyModel(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ModelToUnstructured(r), nil
}

func ModelHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToModel(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToModel(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyModel(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteModel(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToModel(u)
	if err != nil {
		return err
	}
	return c.DeleteModel(ctx, r)
}

func ModelID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToModel(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Model) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"vertexai",
		"Model",
		"alpha",
	}
}

func (r *Model) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Model) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Model) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Model) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Model) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Model) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Model) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetModel(ctx, config, resource)
}

func (r *Model) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyModel(ctx, config, resource, opts...)
}

func (r *Model) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ModelHasDiff(ctx, config, resource, opts...)
}

func (r *Model) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteModel(ctx, config, resource)
}

func (r *Model) ID(resource *unstructured.Resource) (string, error) {
	return ModelID(resource)
}

func init() {
	unstructured.Register(&Model{})
}
