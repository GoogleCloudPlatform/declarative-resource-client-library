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
package bigquery

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Routine struct{}

func RoutineToUnstructured(r *dclService.Routine) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "bigquery",
			Version: "ga",
			Type:    "Routine",
		},
		Object: make(map[string]interface{}),
	}
	var rArguments []interface{}
	for _, rArgumentsVal := range r.Arguments {
		rArgumentsObject := make(map[string]interface{})
		if rArgumentsVal.ArgumentKind != nil {
			rArgumentsObject["argumentKind"] = string(*rArgumentsVal.ArgumentKind)
		}
		if rArgumentsVal.DataType != nil {
			rArgumentsObject["dataType"] = RoutineArgumentsDataTypeToUnstructured(rArgumentsVal.DataType)
		}
		if rArgumentsVal.Mode != nil {
			rArgumentsObject["mode"] = string(*rArgumentsVal.Mode)
		}
		if rArgumentsVal.Name != nil {
			rArgumentsObject["name"] = *rArgumentsVal.Name
		}
		rArguments = append(rArguments, rArgumentsObject)
	}
	u.Object["arguments"] = rArguments
	if r.CreationTime != nil {
		u.Object["creationTime"] = *r.CreationTime
	}
	if r.Dataset != nil {
		u.Object["dataset"] = *r.Dataset
	}
	if r.DefinitionBody != nil {
		u.Object["definitionBody"] = *r.DefinitionBody
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DeterminismLevel != nil {
		u.Object["determinismLevel"] = string(*r.DeterminismLevel)
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	var rImportedLibraries []interface{}
	for _, rImportedLibrariesVal := range r.ImportedLibraries {
		rImportedLibraries = append(rImportedLibraries, rImportedLibrariesVal)
	}
	u.Object["importedLibraries"] = rImportedLibraries
	if r.Language != nil {
		u.Object["language"] = string(*r.Language)
	}
	if r.LastModifiedTime != nil {
		u.Object["lastModifiedTime"] = *r.LastModifiedTime
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ReturnType != nil {
		u.Object["returnType"] = RoutineArgumentsDataTypeToUnstructured(r.ReturnType)
	}
	if r.RoutineType != nil {
		u.Object["routineType"] = string(*r.RoutineType)
	}
	if r.StrictMode != nil {
		u.Object["strictMode"] = *r.StrictMode
	}
	return u
}

func RoutineArgumentsDataTypeToUnstructured(r *dclService.RoutineArgumentsDataType) map[string]interface{} {
	result := make(map[string]interface{})
	if r.ArrayElementType != nil {
		result["arrayElementType"] = RoutineArgumentsDataTypeToUnstructured(r.ArrayElementType)
	}
	if r.StructType != nil && r.StructType != dclService.EmptyRoutineArgumentsDataTypeStructType {
		rStructType := make(map[string]interface{})
		var rStructTypeFields []interface{}
		for _, rStructTypeFieldsVal := range r.StructType.Fields {
			rStructTypeFieldsObject := make(map[string]interface{})
			if rStructTypeFieldsVal.Name != nil {
				rStructTypeFieldsObject["name"] = *rStructTypeFieldsVal.Name
			}
			if rStructTypeFieldsVal.Type != nil {
				rStructTypeFieldsObject["type"] = RoutineArgumentsDataTypeToUnstructured(rStructTypeFieldsVal.Type)
			}
			rStructTypeFields = append(rStructTypeFields, rStructTypeFieldsObject)
		}
		rStructType["fields"] = rStructTypeFields
		result["structType"] = rStructType
	}
	if r.TypeKind != nil {
		result["typeKind"] = string(*r.TypeKind)
	}
	return result
}

func UnstructuredToRoutine(u *unstructured.Resource) (*dclService.Routine, error) {
	r := &dclService.Routine{}
	if _, ok := u.Object["arguments"]; ok {
		if s, ok := u.Object["arguments"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rArguments dclService.RoutineArguments
					if _, ok := objval["argumentKind"]; ok {
						if s, ok := objval["argumentKind"].(string); ok {
							rArguments.ArgumentKind = dclService.RoutineArgumentsArgumentKindEnumRef(s)
						} else {
							return nil, fmt.Errorf("rArguments.ArgumentKind: expected string")
						}
					}
					if _, ok := objval["dataType"]; ok {
						if rArgumentsDataType, ok := objval["dataType"].(map[string]interface{}); ok {
							var err error
							rArguments.DataType, err = UnstructuredToRoutineArgumentsDataType(rArgumentsDataType)
							if err != nil {
								return nil, err
							}
						} else {
							return nil, fmt.Errorf("rArguments.DataType: expected map[string]interface{}")
						}
					}
					if _, ok := objval["mode"]; ok {
						if s, ok := objval["mode"].(string); ok {
							rArguments.Mode = dclService.RoutineArgumentsModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rArguments.Mode: expected string")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rArguments.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rArguments.Name: expected string")
						}
					}
					r.Arguments = append(r.Arguments, rArguments)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Arguments: expected []interface{}")
		}
	}
	if _, ok := u.Object["creationTime"]; ok {
		if i, ok := u.Object["creationTime"].(int64); ok {
			r.CreationTime = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.CreationTime: expected int64")
		}
	}
	if _, ok := u.Object["dataset"]; ok {
		if s, ok := u.Object["dataset"].(string); ok {
			r.Dataset = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Dataset: expected string")
		}
	}
	if _, ok := u.Object["definitionBody"]; ok {
		if s, ok := u.Object["definitionBody"].(string); ok {
			r.DefinitionBody = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DefinitionBody: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["determinismLevel"]; ok {
		if s, ok := u.Object["determinismLevel"].(string); ok {
			r.DeterminismLevel = dclService.RoutineDeterminismLevelEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.DeterminismLevel: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["importedLibraries"]; ok {
		if s, ok := u.Object["importedLibraries"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.ImportedLibraries = append(r.ImportedLibraries, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ImportedLibraries: expected []interface{}")
		}
	}
	if _, ok := u.Object["language"]; ok {
		if s, ok := u.Object["language"].(string); ok {
			r.Language = dclService.RoutineLanguageEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Language: expected string")
		}
	}
	if _, ok := u.Object["lastModifiedTime"]; ok {
		if i, ok := u.Object["lastModifiedTime"].(int64); ok {
			r.LastModifiedTime = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.LastModifiedTime: expected int64")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["returnType"]; ok {
		if rReturnType, ok := u.Object["returnType"].(map[string]interface{}); ok {
			var err error
			r.ReturnType, err = UnstructuredToRoutineArgumentsDataType(rReturnType)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("r.ReturnType: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["routineType"]; ok {
		if s, ok := u.Object["routineType"].(string); ok {
			r.RoutineType = dclService.RoutineRoutineTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.RoutineType: expected string")
		}
	}
	if _, ok := u.Object["strictMode"]; ok {
		if b, ok := u.Object["strictMode"].(bool); ok {
			r.StrictMode = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.StrictMode: expected bool")
		}
	}
	return r, nil
}

func UnstructuredToRoutineArgumentsDataType(obj map[string]interface{}) (*dclService.RoutineArgumentsDataType, error) {
	r := &dclService.RoutineArgumentsDataType{}
	if _, ok := obj["arrayElementType"]; ok {
		if rArrayElementType, ok := obj["arrayElementType"].(map[string]interface{}); ok {
			var err error
			r.ArrayElementType, err = UnstructuredToRoutineArgumentsDataType(rArrayElementType)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("r.ArrayElementType: expected map[string]interface{}")
		}
	}
	if _, ok := obj["structType"]; ok {
		if rStructType, ok := obj["structType"].(map[string]interface{}); ok {
			r.StructType = &dclService.RoutineArgumentsDataTypeStructType{}
			if _, ok := rStructType["fields"]; ok {
				if s, ok := rStructType["fields"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rStructTypeFields dclService.RoutineArgumentsDataTypeStructTypeFields
							if _, ok := objval["name"]; ok {
								if s, ok := objval["name"].(string); ok {
									rStructTypeFields.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rStructTypeFields.Name: expected string")
								}
							}
							if _, ok := objval["type"]; ok {
								if rStructTypeFieldsType, ok := objval["type"].(map[string]interface{}); ok {
									var err error
									rStructTypeFields.Type, err = UnstructuredToRoutineArgumentsDataType(rStructTypeFieldsType)
									if err != nil {
										return nil, err
									}
								} else {
									return nil, fmt.Errorf("rStructTypeFields.Type: expected map[string]interface{}")
								}
							}
							r.StructType.Fields = append(r.StructType.Fields, rStructTypeFields)
						}
					}
				} else {
					return nil, fmt.Errorf("r.StructType.Fields: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.StructType: expected map[string]interface{}")
		}
	}
	if _, ok := obj["typeKind"]; ok {
		if s, ok := obj["typeKind"].(string); ok {
			r.TypeKind = dclService.RoutineArgumentsDataTypeTypeKindEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.TypeKind: expected string")
		}
	}
	return r, nil
}

func GetRoutine(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRoutine(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetRoutine(ctx, r)
	if err != nil {
		return nil, err
	}
	return RoutineToUnstructured(r), nil
}

func ListRoutine(ctx context.Context, config *dcl.Config, project string, dataset string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListRoutine(ctx, project, dataset)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, RoutineToUnstructured(r))
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

func ApplyRoutine(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRoutine(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRoutine(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyRoutine(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return RoutineToUnstructured(r), nil
}

func RoutineHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRoutine(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRoutine(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyRoutine(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteRoutine(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRoutine(u)
	if err != nil {
		return err
	}
	return c.DeleteRoutine(ctx, r)
}

func RoutineID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToRoutine(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Routine) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"bigquery",
		"Routine",
		"ga",
	}
}

func (r *Routine) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Routine) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Routine) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Routine) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Routine) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetRoutine(ctx, config, resource)
}

func (r *Routine) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyRoutine(ctx, config, resource, opts...)
}

func (r *Routine) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return RoutineHasDiff(ctx, config, resource, opts...)
}

func (r *Routine) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteRoutine(ctx, config, resource)
}

func (r *Routine) ID(resource *unstructured.Resource) (string, error) {
	return RoutineID(resource)
}

func init() {
	unstructured.Register(&Routine{})
}
