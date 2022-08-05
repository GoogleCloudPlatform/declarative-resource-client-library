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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Endpoint struct{}

func EndpointToUnstructured(r *dclService.Endpoint) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "vertexai",
			Version: "ga",
			Type:    "Endpoint",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	var rDeployedModels []interface{}
	for _, rDeployedModelsVal := range r.DeployedModels {
		rDeployedModelsObject := make(map[string]interface{})
		if rDeployedModelsVal.AutomaticResources != nil && rDeployedModelsVal.AutomaticResources != dclService.EmptyEndpointDeployedModelsAutomaticResources {
			rDeployedModelsValAutomaticResources := make(map[string]interface{})
			if rDeployedModelsVal.AutomaticResources.MaxReplicaCount != nil {
				rDeployedModelsValAutomaticResources["maxReplicaCount"] = *rDeployedModelsVal.AutomaticResources.MaxReplicaCount
			}
			if rDeployedModelsVal.AutomaticResources.MinReplicaCount != nil {
				rDeployedModelsValAutomaticResources["minReplicaCount"] = *rDeployedModelsVal.AutomaticResources.MinReplicaCount
			}
			rDeployedModelsObject["automaticResources"] = rDeployedModelsValAutomaticResources
		}
		if rDeployedModelsVal.CreateTime != nil {
			rDeployedModelsObject["createTime"] = *rDeployedModelsVal.CreateTime
		}
		if rDeployedModelsVal.DedicatedResources != nil && rDeployedModelsVal.DedicatedResources != dclService.EmptyEndpointDeployedModelsDedicatedResources {
			rDeployedModelsValDedicatedResources := make(map[string]interface{})
			var rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecs []interface{}
			for _, rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecsVal := range rDeployedModelsVal.DedicatedResources.AutoscalingMetricSpecs {
				rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecsObject := make(map[string]interface{})
				if rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecsVal.MetricName != nil {
					rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecsObject["metricName"] = *rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecsVal.MetricName
				}
				if rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecsVal.Target != nil {
					rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecsObject["target"] = *rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecsVal.Target
				}
				rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecs = append(rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecs, rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecsObject)
			}
			rDeployedModelsValDedicatedResources["autoscalingMetricSpecs"] = rDeployedModelsValDedicatedResourcesAutoscalingMetricSpecs
			if rDeployedModelsVal.DedicatedResources.MachineSpec != nil && rDeployedModelsVal.DedicatedResources.MachineSpec != dclService.EmptyEndpointDeployedModelsDedicatedResourcesMachineSpec {
				rDeployedModelsValDedicatedResourcesMachineSpec := make(map[string]interface{})
				if rDeployedModelsVal.DedicatedResources.MachineSpec.AcceleratorCount != nil {
					rDeployedModelsValDedicatedResourcesMachineSpec["acceleratorCount"] = *rDeployedModelsVal.DedicatedResources.MachineSpec.AcceleratorCount
				}
				if rDeployedModelsVal.DedicatedResources.MachineSpec.AcceleratorType != nil {
					rDeployedModelsValDedicatedResourcesMachineSpec["acceleratorType"] = string(*rDeployedModelsVal.DedicatedResources.MachineSpec.AcceleratorType)
				}
				if rDeployedModelsVal.DedicatedResources.MachineSpec.MachineType != nil {
					rDeployedModelsValDedicatedResourcesMachineSpec["machineType"] = *rDeployedModelsVal.DedicatedResources.MachineSpec.MachineType
				}
				rDeployedModelsValDedicatedResources["machineSpec"] = rDeployedModelsValDedicatedResourcesMachineSpec
			}
			if rDeployedModelsVal.DedicatedResources.MaxReplicaCount != nil {
				rDeployedModelsValDedicatedResources["maxReplicaCount"] = *rDeployedModelsVal.DedicatedResources.MaxReplicaCount
			}
			if rDeployedModelsVal.DedicatedResources.MinReplicaCount != nil {
				rDeployedModelsValDedicatedResources["minReplicaCount"] = *rDeployedModelsVal.DedicatedResources.MinReplicaCount
			}
			rDeployedModelsObject["dedicatedResources"] = rDeployedModelsValDedicatedResources
		}
		if rDeployedModelsVal.DisplayName != nil {
			rDeployedModelsObject["displayName"] = *rDeployedModelsVal.DisplayName
		}
		if rDeployedModelsVal.EnableAccessLogging != nil {
			rDeployedModelsObject["enableAccessLogging"] = *rDeployedModelsVal.EnableAccessLogging
		}
		if rDeployedModelsVal.Id != nil {
			rDeployedModelsObject["id"] = *rDeployedModelsVal.Id
		}
		if rDeployedModelsVal.Model != nil {
			rDeployedModelsObject["model"] = *rDeployedModelsVal.Model
		}
		if rDeployedModelsVal.ModelVersionId != nil {
			rDeployedModelsObject["modelVersionId"] = *rDeployedModelsVal.ModelVersionId
		}
		if rDeployedModelsVal.PrivateEndpoints != nil && rDeployedModelsVal.PrivateEndpoints != dclService.EmptyEndpointDeployedModelsPrivateEndpoints {
			rDeployedModelsValPrivateEndpoints := make(map[string]interface{})
			if rDeployedModelsVal.PrivateEndpoints.ExplainHttpUri != nil {
				rDeployedModelsValPrivateEndpoints["explainHttpUri"] = *rDeployedModelsVal.PrivateEndpoints.ExplainHttpUri
			}
			if rDeployedModelsVal.PrivateEndpoints.HealthHttpUri != nil {
				rDeployedModelsValPrivateEndpoints["healthHttpUri"] = *rDeployedModelsVal.PrivateEndpoints.HealthHttpUri
			}
			if rDeployedModelsVal.PrivateEndpoints.PredictHttpUri != nil {
				rDeployedModelsValPrivateEndpoints["predictHttpUri"] = *rDeployedModelsVal.PrivateEndpoints.PredictHttpUri
			}
			if rDeployedModelsVal.PrivateEndpoints.ServiceAttachment != nil {
				rDeployedModelsValPrivateEndpoints["serviceAttachment"] = *rDeployedModelsVal.PrivateEndpoints.ServiceAttachment
			}
			rDeployedModelsObject["privateEndpoints"] = rDeployedModelsValPrivateEndpoints
		}
		if rDeployedModelsVal.ServiceAccount != nil {
			rDeployedModelsObject["serviceAccount"] = *rDeployedModelsVal.ServiceAccount
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
	if r.EncryptionSpec != nil && r.EncryptionSpec != dclService.EmptyEndpointEncryptionSpec {
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
	if r.ModelDeploymentMonitoringJob != nil {
		u.Object["modelDeploymentMonitoringJob"] = *r.ModelDeploymentMonitoringJob
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Network != nil {
		u.Object["network"] = *r.Network
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToEndpoint(u *unstructured.Resource) (*dclService.Endpoint, error) {
	r := &dclService.Endpoint{}
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
					var rDeployedModels dclService.EndpointDeployedModels
					if _, ok := objval["automaticResources"]; ok {
						if rDeployedModelsAutomaticResources, ok := objval["automaticResources"].(map[string]interface{}); ok {
							rDeployedModels.AutomaticResources = &dclService.EndpointDeployedModelsAutomaticResources{}
							if _, ok := rDeployedModelsAutomaticResources["maxReplicaCount"]; ok {
								if i, ok := rDeployedModelsAutomaticResources["maxReplicaCount"].(int64); ok {
									rDeployedModels.AutomaticResources.MaxReplicaCount = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rDeployedModels.AutomaticResources.MaxReplicaCount: expected int64")
								}
							}
							if _, ok := rDeployedModelsAutomaticResources["minReplicaCount"]; ok {
								if i, ok := rDeployedModelsAutomaticResources["minReplicaCount"].(int64); ok {
									rDeployedModels.AutomaticResources.MinReplicaCount = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rDeployedModels.AutomaticResources.MinReplicaCount: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("rDeployedModels.AutomaticResources: expected map[string]interface{}")
						}
					}
					if _, ok := objval["createTime"]; ok {
						if s, ok := objval["createTime"].(string); ok {
							rDeployedModels.CreateTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDeployedModels.CreateTime: expected string")
						}
					}
					if _, ok := objval["dedicatedResources"]; ok {
						if rDeployedModelsDedicatedResources, ok := objval["dedicatedResources"].(map[string]interface{}); ok {
							rDeployedModels.DedicatedResources = &dclService.EndpointDeployedModelsDedicatedResources{}
							if _, ok := rDeployedModelsDedicatedResources["autoscalingMetricSpecs"]; ok {
								if s, ok := rDeployedModelsDedicatedResources["autoscalingMetricSpecs"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rDeployedModelsDedicatedResourcesAutoscalingMetricSpecs dclService.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs
											if _, ok := objval["metricName"]; ok {
												if s, ok := objval["metricName"].(string); ok {
													rDeployedModelsDedicatedResourcesAutoscalingMetricSpecs.MetricName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rDeployedModelsDedicatedResourcesAutoscalingMetricSpecs.MetricName: expected string")
												}
											}
											if _, ok := objval["target"]; ok {
												if i, ok := objval["target"].(int64); ok {
													rDeployedModelsDedicatedResourcesAutoscalingMetricSpecs.Target = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rDeployedModelsDedicatedResourcesAutoscalingMetricSpecs.Target: expected int64")
												}
											}
											rDeployedModels.DedicatedResources.AutoscalingMetricSpecs = append(rDeployedModels.DedicatedResources.AutoscalingMetricSpecs, rDeployedModelsDedicatedResourcesAutoscalingMetricSpecs)
										}
									}
								} else {
									return nil, fmt.Errorf("rDeployedModels.DedicatedResources.AutoscalingMetricSpecs: expected []interface{}")
								}
							}
							if _, ok := rDeployedModelsDedicatedResources["machineSpec"]; ok {
								if rDeployedModelsDedicatedResourcesMachineSpec, ok := rDeployedModelsDedicatedResources["machineSpec"].(map[string]interface{}); ok {
									rDeployedModels.DedicatedResources.MachineSpec = &dclService.EndpointDeployedModelsDedicatedResourcesMachineSpec{}
									if _, ok := rDeployedModelsDedicatedResourcesMachineSpec["acceleratorCount"]; ok {
										if i, ok := rDeployedModelsDedicatedResourcesMachineSpec["acceleratorCount"].(int64); ok {
											rDeployedModels.DedicatedResources.MachineSpec.AcceleratorCount = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rDeployedModels.DedicatedResources.MachineSpec.AcceleratorCount: expected int64")
										}
									}
									if _, ok := rDeployedModelsDedicatedResourcesMachineSpec["acceleratorType"]; ok {
										if s, ok := rDeployedModelsDedicatedResourcesMachineSpec["acceleratorType"].(string); ok {
											rDeployedModels.DedicatedResources.MachineSpec.AcceleratorType = dclService.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumRef(s)
										} else {
											return nil, fmt.Errorf("rDeployedModels.DedicatedResources.MachineSpec.AcceleratorType: expected string")
										}
									}
									if _, ok := rDeployedModelsDedicatedResourcesMachineSpec["machineType"]; ok {
										if s, ok := rDeployedModelsDedicatedResourcesMachineSpec["machineType"].(string); ok {
											rDeployedModels.DedicatedResources.MachineSpec.MachineType = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rDeployedModels.DedicatedResources.MachineSpec.MachineType: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rDeployedModels.DedicatedResources.MachineSpec: expected map[string]interface{}")
								}
							}
							if _, ok := rDeployedModelsDedicatedResources["maxReplicaCount"]; ok {
								if i, ok := rDeployedModelsDedicatedResources["maxReplicaCount"].(int64); ok {
									rDeployedModels.DedicatedResources.MaxReplicaCount = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rDeployedModels.DedicatedResources.MaxReplicaCount: expected int64")
								}
							}
							if _, ok := rDeployedModelsDedicatedResources["minReplicaCount"]; ok {
								if i, ok := rDeployedModelsDedicatedResources["minReplicaCount"].(int64); ok {
									rDeployedModels.DedicatedResources.MinReplicaCount = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rDeployedModels.DedicatedResources.MinReplicaCount: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("rDeployedModels.DedicatedResources: expected map[string]interface{}")
						}
					}
					if _, ok := objval["displayName"]; ok {
						if s, ok := objval["displayName"].(string); ok {
							rDeployedModels.DisplayName = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDeployedModels.DisplayName: expected string")
						}
					}
					if _, ok := objval["enableAccessLogging"]; ok {
						if b, ok := objval["enableAccessLogging"].(bool); ok {
							rDeployedModels.EnableAccessLogging = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("rDeployedModels.EnableAccessLogging: expected bool")
						}
					}
					if _, ok := objval["id"]; ok {
						if s, ok := objval["id"].(string); ok {
							rDeployedModels.Id = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDeployedModels.Id: expected string")
						}
					}
					if _, ok := objval["model"]; ok {
						if s, ok := objval["model"].(string); ok {
							rDeployedModels.Model = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDeployedModels.Model: expected string")
						}
					}
					if _, ok := objval["modelVersionId"]; ok {
						if s, ok := objval["modelVersionId"].(string); ok {
							rDeployedModels.ModelVersionId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDeployedModels.ModelVersionId: expected string")
						}
					}
					if _, ok := objval["privateEndpoints"]; ok {
						if rDeployedModelsPrivateEndpoints, ok := objval["privateEndpoints"].(map[string]interface{}); ok {
							rDeployedModels.PrivateEndpoints = &dclService.EndpointDeployedModelsPrivateEndpoints{}
							if _, ok := rDeployedModelsPrivateEndpoints["explainHttpUri"]; ok {
								if s, ok := rDeployedModelsPrivateEndpoints["explainHttpUri"].(string); ok {
									rDeployedModels.PrivateEndpoints.ExplainHttpUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDeployedModels.PrivateEndpoints.ExplainHttpUri: expected string")
								}
							}
							if _, ok := rDeployedModelsPrivateEndpoints["healthHttpUri"]; ok {
								if s, ok := rDeployedModelsPrivateEndpoints["healthHttpUri"].(string); ok {
									rDeployedModels.PrivateEndpoints.HealthHttpUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDeployedModels.PrivateEndpoints.HealthHttpUri: expected string")
								}
							}
							if _, ok := rDeployedModelsPrivateEndpoints["predictHttpUri"]; ok {
								if s, ok := rDeployedModelsPrivateEndpoints["predictHttpUri"].(string); ok {
									rDeployedModels.PrivateEndpoints.PredictHttpUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDeployedModels.PrivateEndpoints.PredictHttpUri: expected string")
								}
							}
							if _, ok := rDeployedModelsPrivateEndpoints["serviceAttachment"]; ok {
								if s, ok := rDeployedModelsPrivateEndpoints["serviceAttachment"].(string); ok {
									rDeployedModels.PrivateEndpoints.ServiceAttachment = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDeployedModels.PrivateEndpoints.ServiceAttachment: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rDeployedModels.PrivateEndpoints: expected map[string]interface{}")
						}
					}
					if _, ok := objval["serviceAccount"]; ok {
						if s, ok := objval["serviceAccount"].(string); ok {
							rDeployedModels.ServiceAccount = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDeployedModels.ServiceAccount: expected string")
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
			r.EncryptionSpec = &dclService.EndpointEncryptionSpec{}
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
	if _, ok := u.Object["modelDeploymentMonitoringJob"]; ok {
		if s, ok := u.Object["modelDeploymentMonitoringJob"].(string); ok {
			r.ModelDeploymentMonitoringJob = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ModelDeploymentMonitoringJob: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["network"]; ok {
		if s, ok := u.Object["network"].(string); ok {
			r.Network = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Network: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetEndpoint(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpoint(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetEndpoint(ctx, r)
	if err != nil {
		return nil, err
	}
	return EndpointToUnstructured(r), nil
}

func ListEndpoint(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListEndpoint(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, EndpointToUnstructured(r))
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

func ApplyEndpoint(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpoint(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEndpoint(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyEndpoint(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return EndpointToUnstructured(r), nil
}

func EndpointHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpoint(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEndpoint(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyEndpoint(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteEndpoint(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpoint(u)
	if err != nil {
		return err
	}
	return c.DeleteEndpoint(ctx, r)
}

func EndpointID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToEndpoint(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Endpoint) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"vertexai",
		"Endpoint",
		"ga",
	}
}

func (r *Endpoint) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Endpoint) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Endpoint) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Endpoint) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Endpoint) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Endpoint) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Endpoint) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetEndpoint(ctx, config, resource)
}

func (r *Endpoint) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyEndpoint(ctx, config, resource, opts...)
}

func (r *Endpoint) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return EndpointHasDiff(ctx, config, resource, opts...)
}

func (r *Endpoint) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteEndpoint(ctx, config, resource)
}

func (r *Endpoint) ID(resource *unstructured.Resource) (string, error) {
	return EndpointID(resource)
}

func init() {
	unstructured.Register(&Endpoint{})
}
