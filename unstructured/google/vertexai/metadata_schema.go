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

type MetadataSchema struct{}

func MetadataSchemaToUnstructured(r *dclService.MetadataSchema) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "vertexai",
			Version: "ga",
			Type:    "MetadataSchema",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.MetadataStore != nil {
		u.Object["metadataStore"] = *r.MetadataStore
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Schema != nil {
		u.Object["schema"] = *r.Schema
	}
	if r.SchemaType != nil {
		u.Object["schemaType"] = string(*r.SchemaType)
	}
	if r.SchemaVersion != nil {
		u.Object["schemaVersion"] = *r.SchemaVersion
	}
	return u
}

func UnstructuredToMetadataSchema(u *unstructured.Resource) (*dclService.MetadataSchema, error) {
	r := &dclService.MetadataSchema{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["metadataStore"]; ok {
		if s, ok := u.Object["metadataStore"].(string); ok {
			r.MetadataStore = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.MetadataStore: expected string")
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
	if _, ok := u.Object["schema"]; ok {
		if s, ok := u.Object["schema"].(string); ok {
			r.Schema = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Schema: expected string")
		}
	}
	if _, ok := u.Object["schemaType"]; ok {
		if s, ok := u.Object["schemaType"].(string); ok {
			r.SchemaType = dclService.MetadataSchemaSchemaTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.SchemaType: expected string")
		}
	}
	if _, ok := u.Object["schemaVersion"]; ok {
		if s, ok := u.Object["schemaVersion"].(string); ok {
			r.SchemaVersion = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SchemaVersion: expected string")
		}
	}
	return r, nil
}

func GetMetadataSchema(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetadataSchema(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetMetadataSchema(ctx, r)
	if err != nil {
		return nil, err
	}
	return MetadataSchemaToUnstructured(r), nil
}

func ListMetadataSchema(ctx context.Context, config *dcl.Config, project string, location string, metadatastore string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListMetadataSchema(ctx, project, location, metadatastore)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, MetadataSchemaToUnstructured(r))
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

func ApplyMetadataSchema(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetadataSchema(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMetadataSchema(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyMetadataSchema(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return MetadataSchemaToUnstructured(r), nil
}

func MetadataSchemaHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetadataSchema(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMetadataSchema(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyMetadataSchema(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteMetadataSchema(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func MetadataSchemaID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToMetadataSchema(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *MetadataSchema) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"vertexai",
		"MetadataSchema",
		"ga",
	}
}

func (r *MetadataSchema) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataSchema) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataSchema) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *MetadataSchema) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataSchema) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataSchema) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataSchema) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetMetadataSchema(ctx, config, resource)
}

func (r *MetadataSchema) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyMetadataSchema(ctx, config, resource, opts...)
}

func (r *MetadataSchema) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return MetadataSchemaHasDiff(ctx, config, resource, opts...)
}

func (r *MetadataSchema) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteMetadataSchema(ctx, config, resource)
}

func (r *MetadataSchema) ID(resource *unstructured.Resource) (string, error) {
	return MetadataSchemaID(resource)
}

func init() {
	unstructured.Register(&MetadataSchema{})
}
