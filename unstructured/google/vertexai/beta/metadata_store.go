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

type MetadataStore struct{}

func MetadataStoreToUnstructured(r *dclService.MetadataStore) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "vertexai",
			Version: "beta",
			Type:    "MetadataStore",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.EncryptionSpec != nil && r.EncryptionSpec != dclService.EmptyMetadataStoreEncryptionSpec {
		rEncryptionSpec := make(map[string]interface{})
		if r.EncryptionSpec.KmsKeyName != nil {
			rEncryptionSpec["kmsKeyName"] = *r.EncryptionSpec.KmsKeyName
		}
		u.Object["encryptionSpec"] = rEncryptionSpec
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.State != nil && r.State != dclService.EmptyMetadataStoreState {
		rState := make(map[string]interface{})
		if r.State.DiskUtilizationBytes != nil {
			rState["diskUtilizationBytes"] = *r.State.DiskUtilizationBytes
		}
		u.Object["state"] = rState
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToMetadataStore(u *unstructured.Resource) (*dclService.MetadataStore, error) {
	r := &dclService.MetadataStore{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["encryptionSpec"]; ok {
		if rEncryptionSpec, ok := u.Object["encryptionSpec"].(map[string]interface{}); ok {
			r.EncryptionSpec = &dclService.MetadataStoreEncryptionSpec{}
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
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if rState, ok := u.Object["state"].(map[string]interface{}); ok {
			r.State = &dclService.MetadataStoreState{}
			if _, ok := rState["diskUtilizationBytes"]; ok {
				if i, ok := rState["diskUtilizationBytes"].(int64); ok {
					r.State.DiskUtilizationBytes = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.State.DiskUtilizationBytes: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.State: expected map[string]interface{}")
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

func GetMetadataStore(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetadataStore(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetMetadataStore(ctx, r)
	if err != nil {
		return nil, err
	}
	return MetadataStoreToUnstructured(r), nil
}

func ListMetadataStore(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListMetadataStore(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, MetadataStoreToUnstructured(r))
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

func ApplyMetadataStore(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetadataStore(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMetadataStore(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyMetadataStore(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return MetadataStoreToUnstructured(r), nil
}

func MetadataStoreHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetadataStore(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMetadataStore(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyMetadataStore(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteMetadataStore(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetadataStore(u)
	if err != nil {
		return err
	}
	return c.DeleteMetadataStore(ctx, r)
}

func MetadataStoreID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToMetadataStore(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *MetadataStore) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"vertexai",
		"MetadataStore",
		"beta",
	}
}

func (r *MetadataStore) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataStore) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataStore) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *MetadataStore) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataStore) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataStore) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetadataStore) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetMetadataStore(ctx, config, resource)
}

func (r *MetadataStore) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyMetadataStore(ctx, config, resource, opts...)
}

func (r *MetadataStore) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return MetadataStoreHasDiff(ctx, config, resource, opts...)
}

func (r *MetadataStore) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteMetadataStore(ctx, config, resource)
}

func (r *MetadataStore) ID(resource *unstructured.Resource) (string, error) {
	return MetadataStoreID(resource)
}

func init() {
	unstructured.Register(&MetadataStore{})
}
