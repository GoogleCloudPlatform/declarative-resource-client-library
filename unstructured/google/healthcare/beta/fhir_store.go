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
package healthcare

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/healthcare/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type FhirStore struct{}

func FhirStoreToUnstructured(r *dclService.FhirStore) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "healthcare",
			Version: "beta",
			Type:    "FhirStore",
		},
		Object: make(map[string]interface{}),
	}
	if r.ComplexDataTypeReferenceParsing != nil {
		u.Object["complexDataTypeReferenceParsing"] = string(*r.ComplexDataTypeReferenceParsing)
	}
	if r.Dataset != nil {
		u.Object["dataset"] = *r.Dataset
	}
	if r.DefaultSearchHandlingStrict != nil {
		u.Object["defaultSearchHandlingStrict"] = *r.DefaultSearchHandlingStrict
	}
	if r.DisableReferentialIntegrity != nil {
		u.Object["disableReferentialIntegrity"] = *r.DisableReferentialIntegrity
	}
	if r.DisableResourceVersioning != nil {
		u.Object["disableResourceVersioning"] = *r.DisableResourceVersioning
	}
	if r.EnableUpdateCreate != nil {
		u.Object["enableUpdateCreate"] = *r.EnableUpdateCreate
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
	if r.NotificationConfig != nil && r.NotificationConfig != dclService.EmptyFhirStoreNotificationConfig {
		rNotificationConfig := make(map[string]interface{})
		if r.NotificationConfig.PubsubTopic != nil {
			rNotificationConfig["pubsubTopic"] = *r.NotificationConfig.PubsubTopic
		}
		u.Object["notificationConfig"] = rNotificationConfig
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ShardNum != nil {
		u.Object["shardNum"] = *r.ShardNum
	}
	var rStreamConfigs []interface{}
	for _, rStreamConfigsVal := range r.StreamConfigs {
		rStreamConfigsObject := make(map[string]interface{})
		if rStreamConfigsVal.BigqueryDestination != nil && rStreamConfigsVal.BigqueryDestination != dclService.EmptyFhirStoreStreamConfigsBigqueryDestination {
			rStreamConfigsValBigqueryDestination := make(map[string]interface{})
			if rStreamConfigsVal.BigqueryDestination.DatasetUri != nil {
				rStreamConfigsValBigqueryDestination["datasetUri"] = *rStreamConfigsVal.BigqueryDestination.DatasetUri
			}
			if rStreamConfigsVal.BigqueryDestination.SchemaConfig != nil && rStreamConfigsVal.BigqueryDestination.SchemaConfig != dclService.EmptyFhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
				rStreamConfigsValBigqueryDestinationSchemaConfig := make(map[string]interface{})
				if rStreamConfigsVal.BigqueryDestination.SchemaConfig.RecursiveStructureDepth != nil {
					rStreamConfigsValBigqueryDestinationSchemaConfig["recursiveStructureDepth"] = *rStreamConfigsVal.BigqueryDestination.SchemaConfig.RecursiveStructureDepth
				}
				if rStreamConfigsVal.BigqueryDestination.SchemaConfig.SchemaType != nil {
					rStreamConfigsValBigqueryDestinationSchemaConfig["schemaType"] = string(*rStreamConfigsVal.BigqueryDestination.SchemaConfig.SchemaType)
				}
				rStreamConfigsValBigqueryDestination["schemaConfig"] = rStreamConfigsValBigqueryDestinationSchemaConfig
			}
			rStreamConfigsObject["bigqueryDestination"] = rStreamConfigsValBigqueryDestination
		}
		var rStreamConfigsValResourceTypes []interface{}
		for _, rStreamConfigsValResourceTypesVal := range rStreamConfigsVal.ResourceTypes {
			rStreamConfigsValResourceTypes = append(rStreamConfigsValResourceTypes, rStreamConfigsValResourceTypesVal)
		}
		rStreamConfigsObject["resourceTypes"] = rStreamConfigsValResourceTypes
		rStreamConfigs = append(rStreamConfigs, rStreamConfigsObject)
	}
	u.Object["streamConfigs"] = rStreamConfigs
	if r.ValidationConfig != nil && r.ValidationConfig != dclService.EmptyFhirStoreValidationConfig {
		rValidationConfig := make(map[string]interface{})
		if r.ValidationConfig.DisableFhirpathValidation != nil {
			rValidationConfig["disableFhirpathValidation"] = *r.ValidationConfig.DisableFhirpathValidation
		}
		if r.ValidationConfig.DisableProfileValidation != nil {
			rValidationConfig["disableProfileValidation"] = *r.ValidationConfig.DisableProfileValidation
		}
		if r.ValidationConfig.DisableReferenceTypeValidation != nil {
			rValidationConfig["disableReferenceTypeValidation"] = *r.ValidationConfig.DisableReferenceTypeValidation
		}
		if r.ValidationConfig.DisableRequiredFieldValidation != nil {
			rValidationConfig["disableRequiredFieldValidation"] = *r.ValidationConfig.DisableRequiredFieldValidation
		}
		var rValidationConfigEnabledImplementationGuides []interface{}
		for _, rValidationConfigEnabledImplementationGuidesVal := range r.ValidationConfig.EnabledImplementationGuides {
			rValidationConfigEnabledImplementationGuides = append(rValidationConfigEnabledImplementationGuides, rValidationConfigEnabledImplementationGuidesVal)
		}
		rValidationConfig["enabledImplementationGuides"] = rValidationConfigEnabledImplementationGuides
		u.Object["validationConfig"] = rValidationConfig
	}
	if r.Version != nil {
		u.Object["version"] = string(*r.Version)
	}
	return u
}

func UnstructuredToFhirStore(u *unstructured.Resource) (*dclService.FhirStore, error) {
	r := &dclService.FhirStore{}
	if _, ok := u.Object["complexDataTypeReferenceParsing"]; ok {
		if s, ok := u.Object["complexDataTypeReferenceParsing"].(string); ok {
			r.ComplexDataTypeReferenceParsing = dclService.FhirStoreComplexDataTypeReferenceParsingEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.ComplexDataTypeReferenceParsing: expected string")
		}
	}
	if _, ok := u.Object["dataset"]; ok {
		if s, ok := u.Object["dataset"].(string); ok {
			r.Dataset = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Dataset: expected string")
		}
	}
	if _, ok := u.Object["defaultSearchHandlingStrict"]; ok {
		if b, ok := u.Object["defaultSearchHandlingStrict"].(bool); ok {
			r.DefaultSearchHandlingStrict = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.DefaultSearchHandlingStrict: expected bool")
		}
	}
	if _, ok := u.Object["disableReferentialIntegrity"]; ok {
		if b, ok := u.Object["disableReferentialIntegrity"].(bool); ok {
			r.DisableReferentialIntegrity = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.DisableReferentialIntegrity: expected bool")
		}
	}
	if _, ok := u.Object["disableResourceVersioning"]; ok {
		if b, ok := u.Object["disableResourceVersioning"].(bool); ok {
			r.DisableResourceVersioning = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.DisableResourceVersioning: expected bool")
		}
	}
	if _, ok := u.Object["enableUpdateCreate"]; ok {
		if b, ok := u.Object["enableUpdateCreate"].(bool); ok {
			r.EnableUpdateCreate = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableUpdateCreate: expected bool")
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
	if _, ok := u.Object["notificationConfig"]; ok {
		if rNotificationConfig, ok := u.Object["notificationConfig"].(map[string]interface{}); ok {
			r.NotificationConfig = &dclService.FhirStoreNotificationConfig{}
			if _, ok := rNotificationConfig["pubsubTopic"]; ok {
				if s, ok := rNotificationConfig["pubsubTopic"].(string); ok {
					r.NotificationConfig.PubsubTopic = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.NotificationConfig.PubsubTopic: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.NotificationConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["shardNum"]; ok {
		if i, ok := u.Object["shardNum"].(int64); ok {
			r.ShardNum = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.ShardNum: expected int64")
		}
	}
	if _, ok := u.Object["streamConfigs"]; ok {
		if s, ok := u.Object["streamConfigs"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rStreamConfigs dclService.FhirStoreStreamConfigs
					if _, ok := objval["bigqueryDestination"]; ok {
						if rStreamConfigsBigqueryDestination, ok := objval["bigqueryDestination"].(map[string]interface{}); ok {
							rStreamConfigs.BigqueryDestination = &dclService.FhirStoreStreamConfigsBigqueryDestination{}
							if _, ok := rStreamConfigsBigqueryDestination["datasetUri"]; ok {
								if s, ok := rStreamConfigsBigqueryDestination["datasetUri"].(string); ok {
									rStreamConfigs.BigqueryDestination.DatasetUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rStreamConfigs.BigqueryDestination.DatasetUri: expected string")
								}
							}
							if _, ok := rStreamConfigsBigqueryDestination["schemaConfig"]; ok {
								if rStreamConfigsBigqueryDestinationSchemaConfig, ok := rStreamConfigsBigqueryDestination["schemaConfig"].(map[string]interface{}); ok {
									rStreamConfigs.BigqueryDestination.SchemaConfig = &dclService.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
									if _, ok := rStreamConfigsBigqueryDestinationSchemaConfig["recursiveStructureDepth"]; ok {
										if i, ok := rStreamConfigsBigqueryDestinationSchemaConfig["recursiveStructureDepth"].(int64); ok {
											rStreamConfigs.BigqueryDestination.SchemaConfig.RecursiveStructureDepth = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rStreamConfigs.BigqueryDestination.SchemaConfig.RecursiveStructureDepth: expected int64")
										}
									}
									if _, ok := rStreamConfigsBigqueryDestinationSchemaConfig["schemaType"]; ok {
										if s, ok := rStreamConfigsBigqueryDestinationSchemaConfig["schemaType"].(string); ok {
											rStreamConfigs.BigqueryDestination.SchemaConfig.SchemaType = dclService.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumRef(s)
										} else {
											return nil, fmt.Errorf("rStreamConfigs.BigqueryDestination.SchemaConfig.SchemaType: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rStreamConfigs.BigqueryDestination.SchemaConfig: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rStreamConfigs.BigqueryDestination: expected map[string]interface{}")
						}
					}
					if _, ok := objval["resourceTypes"]; ok {
						if s, ok := objval["resourceTypes"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rStreamConfigs.ResourceTypes = append(rStreamConfigs.ResourceTypes, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rStreamConfigs.ResourceTypes: expected []interface{}")
						}
					}
					r.StreamConfigs = append(r.StreamConfigs, rStreamConfigs)
				}
			}
		} else {
			return nil, fmt.Errorf("r.StreamConfigs: expected []interface{}")
		}
	}
	if _, ok := u.Object["validationConfig"]; ok {
		if rValidationConfig, ok := u.Object["validationConfig"].(map[string]interface{}); ok {
			r.ValidationConfig = &dclService.FhirStoreValidationConfig{}
			if _, ok := rValidationConfig["disableFhirpathValidation"]; ok {
				if b, ok := rValidationConfig["disableFhirpathValidation"].(bool); ok {
					r.ValidationConfig.DisableFhirpathValidation = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ValidationConfig.DisableFhirpathValidation: expected bool")
				}
			}
			if _, ok := rValidationConfig["disableProfileValidation"]; ok {
				if b, ok := rValidationConfig["disableProfileValidation"].(bool); ok {
					r.ValidationConfig.DisableProfileValidation = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ValidationConfig.DisableProfileValidation: expected bool")
				}
			}
			if _, ok := rValidationConfig["disableReferenceTypeValidation"]; ok {
				if b, ok := rValidationConfig["disableReferenceTypeValidation"].(bool); ok {
					r.ValidationConfig.DisableReferenceTypeValidation = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ValidationConfig.DisableReferenceTypeValidation: expected bool")
				}
			}
			if _, ok := rValidationConfig["disableRequiredFieldValidation"]; ok {
				if b, ok := rValidationConfig["disableRequiredFieldValidation"].(bool); ok {
					r.ValidationConfig.DisableRequiredFieldValidation = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ValidationConfig.DisableRequiredFieldValidation: expected bool")
				}
			}
			if _, ok := rValidationConfig["enabledImplementationGuides"]; ok {
				if s, ok := rValidationConfig["enabledImplementationGuides"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.ValidationConfig.EnabledImplementationGuides = append(r.ValidationConfig.EnabledImplementationGuides, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ValidationConfig.EnabledImplementationGuides: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ValidationConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["version"]; ok {
		if s, ok := u.Object["version"].(string); ok {
			r.Version = dclService.FhirStoreVersionEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Version: expected string")
		}
	}
	return r, nil
}

func GetFhirStore(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFhirStore(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetFhirStore(ctx, r)
	if err != nil {
		return nil, err
	}
	return FhirStoreToUnstructured(r), nil
}

func ListFhirStore(ctx context.Context, config *dcl.Config, project string, location string, dataset string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListFhirStore(ctx, project, location, dataset)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, FhirStoreToUnstructured(r))
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

func ApplyFhirStore(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFhirStore(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFhirStore(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyFhirStore(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return FhirStoreToUnstructured(r), nil
}

func FhirStoreHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFhirStore(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFhirStore(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyFhirStore(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteFhirStore(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFhirStore(u)
	if err != nil {
		return err
	}
	return c.DeleteFhirStore(ctx, r)
}

func FhirStoreID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToFhirStore(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *FhirStore) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"healthcare",
		"FhirStore",
		"beta",
	}
}

func (r *FhirStore) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FhirStore) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FhirStore) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *FhirStore) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FhirStore) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FhirStore) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FhirStore) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetFhirStore(ctx, config, resource)
}

func (r *FhirStore) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyFhirStore(ctx, config, resource, opts...)
}

func (r *FhirStore) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return FhirStoreHasDiff(ctx, config, resource, opts...)
}

func (r *FhirStore) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteFhirStore(ctx, config, resource)
}

func (r *FhirStore) ID(resource *unstructured.Resource) (string, error) {
	return FhirStoreID(resource)
}

func init() {
	unstructured.Register(&FhirStore{})
}
