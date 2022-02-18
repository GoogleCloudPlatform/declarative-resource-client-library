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
package servicemanagement

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/servicemanagement/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type ManagedService struct{}

func ManagedServiceToUnstructured(r *dclService.ManagedService) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "servicemanagement",
			Version: "beta",
			Type:    "ManagedService",
		},
		Object: make(map[string]interface{}),
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	return u
}

func UnstructuredToManagedService(u *unstructured.Resource) (*dclService.ManagedService, error) {
	r := &dclService.ManagedService{}
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
	return r, nil
}

func GetManagedService(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToManagedService(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetManagedService(ctx, r)
	if err != nil {
		return nil, err
	}
	return ManagedServiceToUnstructured(r), nil
}

func ListManagedService(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListManagedService(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ManagedServiceToUnstructured(r))
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

func ApplyManagedService(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToManagedService(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToManagedService(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyManagedService(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ManagedServiceToUnstructured(r), nil
}

func ManagedServiceHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToManagedService(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToManagedService(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyManagedService(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteManagedService(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToManagedService(u)
	if err != nil {
		return err
	}
	return c.DeleteManagedService(ctx, r)
}

func ManagedServiceID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToManagedService(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *ManagedService) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"servicemanagement",
		"ManagedService",
		"beta",
	}
}

func (r *ManagedService) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ManagedService) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ManagedService) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ManagedService) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ManagedService) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetManagedService(ctx, config, resource)
}

func (r *ManagedService) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyManagedService(ctx, config, resource, opts...)
}

func (r *ManagedService) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ManagedServiceHasDiff(ctx, config, resource, opts...)
}

func (r *ManagedService) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteManagedService(ctx, config, resource)
}

func (r *ManagedService) ID(resource *unstructured.Resource) (string, error) {
	return ManagedServiceID(resource)
}

func init() {
	unstructured.Register(&ManagedService{})
}
