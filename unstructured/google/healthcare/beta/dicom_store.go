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

type DicomStore struct{}

func DicomStoreToUnstructured(r *dclService.DicomStore) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "healthcare",
			Version: "beta",
			Type:    "DicomStore",
		},
		Object: make(map[string]interface{}),
	}
	if r.Dataset != nil {
		u.Object["dataset"] = *r.Dataset
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
	if r.NotificationConfig != nil && r.NotificationConfig != dclService.EmptyDicomStoreNotificationConfig {
		rNotificationConfig := make(map[string]interface{})
		if r.NotificationConfig.PubsubTopic != nil {
			rNotificationConfig["pubsubTopic"] = *r.NotificationConfig.PubsubTopic
		}
		u.Object["notificationConfig"] = rNotificationConfig
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	return u
}

func UnstructuredToDicomStore(u *unstructured.Resource) (*dclService.DicomStore, error) {
	r := &dclService.DicomStore{}
	if _, ok := u.Object["dataset"]; ok {
		if s, ok := u.Object["dataset"].(string); ok {
			r.Dataset = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Dataset: expected string")
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
			r.NotificationConfig = &dclService.DicomStoreNotificationConfig{}
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
	return r, nil
}

func GetDicomStore(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDicomStore(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetDicomStore(ctx, r)
	if err != nil {
		return nil, err
	}
	return DicomStoreToUnstructured(r), nil
}

func ListDicomStore(ctx context.Context, config *dcl.Config, project string, location string, dataset string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListDicomStore(ctx, project, location, dataset)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, DicomStoreToUnstructured(r))
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

func ApplyDicomStore(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDicomStore(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDicomStore(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyDicomStore(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return DicomStoreToUnstructured(r), nil
}

func DicomStoreHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDicomStore(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDicomStore(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyDicomStore(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteDicomStore(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDicomStore(u)
	if err != nil {
		return err
	}
	return c.DeleteDicomStore(ctx, r)
}

func DicomStoreID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToDicomStore(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *DicomStore) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"healthcare",
		"DicomStore",
		"beta",
	}
}

func (r *DicomStore) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DicomStore) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DicomStore) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *DicomStore) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DicomStore) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DicomStore) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DicomStore) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetDicomStore(ctx, config, resource)
}

func (r *DicomStore) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyDicomStore(ctx, config, resource, opts...)
}

func (r *DicomStore) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return DicomStoreHasDiff(ctx, config, resource, opts...)
}

func (r *DicomStore) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteDicomStore(ctx, config, resource)
}

func (r *DicomStore) ID(resource *unstructured.Resource) (string, error) {
	return DicomStoreID(resource)
}

func init() {
	unstructured.Register(&DicomStore{})
}
