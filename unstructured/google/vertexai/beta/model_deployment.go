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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type ModelDeployment struct{}

func ModelDeploymentToUnstructured(r *dclService.ModelDeployment) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "vertexai",
			Version: "beta",
			Type:    "ModelDeployment",
		},
		Object: make(map[string]interface{}),
	}
	if r.DedicatedResources != nil && r.DedicatedResources != dclService.EmptyModelDeploymentDedicatedResources {
		rDedicatedResources := make(map[string]interface{})
		if r.DedicatedResources.MachineSpec != nil && r.DedicatedResources.MachineSpec != dclService.EmptyModelDeploymentDedicatedResourcesMachineSpec {
			rDedicatedResourcesMachineSpec := make(map[string]interface{})
			if r.DedicatedResources.MachineSpec.MachineType != nil {
				rDedicatedResourcesMachineSpec["machineType"] = *r.DedicatedResources.MachineSpec.MachineType
			}
			rDedicatedResources["machineSpec"] = rDedicatedResourcesMachineSpec
		}
		if r.DedicatedResources.MaxReplicaCount != nil {
			rDedicatedResources["maxReplicaCount"] = *r.DedicatedResources.MaxReplicaCount
		}
		if r.DedicatedResources.MinReplicaCount != nil {
			rDedicatedResources["minReplicaCount"] = *r.DedicatedResources.MinReplicaCount
		}
		u.Object["dedicatedResources"] = rDedicatedResources
	}
	if r.Endpoint != nil {
		u.Object["endpoint"] = *r.Endpoint
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Model != nil {
		u.Object["model"] = *r.Model
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	return u
}

func UnstructuredToModelDeployment(u *unstructured.Resource) (*dclService.ModelDeployment, error) {
	r := &dclService.ModelDeployment{}
	if _, ok := u.Object["dedicatedResources"]; ok {
		if rDedicatedResources, ok := u.Object["dedicatedResources"].(map[string]interface{}); ok {
			r.DedicatedResources = &dclService.ModelDeploymentDedicatedResources{}
			if _, ok := rDedicatedResources["machineSpec"]; ok {
				if rDedicatedResourcesMachineSpec, ok := rDedicatedResources["machineSpec"].(map[string]interface{}); ok {
					r.DedicatedResources.MachineSpec = &dclService.ModelDeploymentDedicatedResourcesMachineSpec{}
					if _, ok := rDedicatedResourcesMachineSpec["machineType"]; ok {
						if s, ok := rDedicatedResourcesMachineSpec["machineType"].(string); ok {
							r.DedicatedResources.MachineSpec.MachineType = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.DedicatedResources.MachineSpec.MachineType: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.DedicatedResources.MachineSpec: expected map[string]interface{}")
				}
			}
			if _, ok := rDedicatedResources["maxReplicaCount"]; ok {
				if i, ok := rDedicatedResources["maxReplicaCount"].(int64); ok {
					r.DedicatedResources.MaxReplicaCount = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.DedicatedResources.MaxReplicaCount: expected int64")
				}
			}
			if _, ok := rDedicatedResources["minReplicaCount"]; ok {
				if i, ok := rDedicatedResources["minReplicaCount"].(int64); ok {
					r.DedicatedResources.MinReplicaCount = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.DedicatedResources.MinReplicaCount: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DedicatedResources: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["endpoint"]; ok {
		if s, ok := u.Object["endpoint"].(string); ok {
			r.Endpoint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Endpoint: expected string")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if s, ok := u.Object["id"].(string); ok {
			r.Id = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Id: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["model"]; ok {
		if s, ok := u.Object["model"].(string); ok {
			r.Model = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Model: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	return r, nil
}

func GetModelDeployment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToModelDeployment(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetModelDeployment(ctx, r)
	if err != nil {
		return nil, err
	}
	return ModelDeploymentToUnstructured(r), nil
}

func ListModelDeployment(ctx context.Context, config *dcl.Config, project string, location string, endpoint string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListModelDeployment(ctx, project, location, endpoint)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ModelDeploymentToUnstructured(r))
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

func ApplyModelDeployment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToModelDeployment(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToModelDeployment(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyModelDeployment(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ModelDeploymentToUnstructured(r), nil
}

func ModelDeploymentHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToModelDeployment(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToModelDeployment(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyModelDeployment(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteModelDeployment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToModelDeployment(u)
	if err != nil {
		return err
	}
	return c.DeleteModelDeployment(ctx, r)
}

func ModelDeploymentID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToModelDeployment(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *ModelDeployment) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"vertexai",
		"ModelDeployment",
		"beta",
	}
}

func (r *ModelDeployment) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ModelDeployment) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ModelDeployment) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *ModelDeployment) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ModelDeployment) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ModelDeployment) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ModelDeployment) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetModelDeployment(ctx, config, resource)
}

func (r *ModelDeployment) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyModelDeployment(ctx, config, resource, opts...)
}

func (r *ModelDeployment) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ModelDeploymentHasDiff(ctx, config, resource, opts...)
}

func (r *ModelDeployment) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteModelDeployment(ctx, config, resource)
}

func (r *ModelDeployment) ID(resource *unstructured.Resource) (string, error) {
	return ModelDeploymentID(resource)
}

func init() {
	unstructured.Register(&ModelDeployment{})
}
