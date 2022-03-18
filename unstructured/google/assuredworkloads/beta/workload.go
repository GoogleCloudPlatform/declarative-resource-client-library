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
package assuredworkloads

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Workload struct{}

func WorkloadToUnstructured(r *dclService.Workload) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "assuredworkloads",
			Version: "beta",
			Type:    "Workload",
		},
		Object: make(map[string]interface{}),
	}
	if r.BillingAccount != nil {
		u.Object["billingAccount"] = *r.BillingAccount
	}
	if r.ComplianceRegime != nil {
		u.Object["complianceRegime"] = string(*r.ComplianceRegime)
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.KmsSettings != nil && r.KmsSettings != dclService.EmptyWorkloadKmsSettings {
		rKmsSettings := make(map[string]interface{})
		if r.KmsSettings.NextRotationTime != nil {
			rKmsSettings["nextRotationTime"] = *r.KmsSettings.NextRotationTime
		}
		if r.KmsSettings.RotationPeriod != nil {
			rKmsSettings["rotationPeriod"] = *r.KmsSettings.RotationPeriod
		}
		u.Object["kmsSettings"] = rKmsSettings
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
	if r.Organization != nil {
		u.Object["organization"] = *r.Organization
	}
	if r.ProvisionedResourcesParent != nil {
		u.Object["provisionedResourcesParent"] = *r.ProvisionedResourcesParent
	}
	var rResourceSettings []interface{}
	for _, rResourceSettingsVal := range r.ResourceSettings {
		rResourceSettingsObject := make(map[string]interface{})
		if rResourceSettingsVal.ResourceId != nil {
			rResourceSettingsObject["resourceId"] = *rResourceSettingsVal.ResourceId
		}
		if rResourceSettingsVal.ResourceType != nil {
			rResourceSettingsObject["resourceType"] = string(*rResourceSettingsVal.ResourceType)
		}
		rResourceSettings = append(rResourceSettings, rResourceSettingsObject)
	}
	u.Object["resourceSettings"] = rResourceSettings
	var rResources []interface{}
	for _, rResourcesVal := range r.Resources {
		rResourcesObject := make(map[string]interface{})
		if rResourcesVal.ResourceId != nil {
			rResourcesObject["resourceId"] = *rResourcesVal.ResourceId
		}
		if rResourcesVal.ResourceType != nil {
			rResourcesObject["resourceType"] = string(*rResourcesVal.ResourceType)
		}
		rResources = append(rResources, rResourcesObject)
	}
	u.Object["resources"] = rResources
	return u
}

func UnstructuredToWorkload(u *unstructured.Resource) (*dclService.Workload, error) {
	r := &dclService.Workload{}
	if _, ok := u.Object["billingAccount"]; ok {
		if s, ok := u.Object["billingAccount"].(string); ok {
			r.BillingAccount = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.BillingAccount: expected string")
		}
	}
	if _, ok := u.Object["complianceRegime"]; ok {
		if s, ok := u.Object["complianceRegime"].(string); ok {
			r.ComplianceRegime = dclService.WorkloadComplianceRegimeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.ComplianceRegime: expected string")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["kmsSettings"]; ok {
		if rKmsSettings, ok := u.Object["kmsSettings"].(map[string]interface{}); ok {
			r.KmsSettings = &dclService.WorkloadKmsSettings{}
			if _, ok := rKmsSettings["nextRotationTime"]; ok {
				if s, ok := rKmsSettings["nextRotationTime"].(string); ok {
					r.KmsSettings.NextRotationTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.KmsSettings.NextRotationTime: expected string")
				}
			}
			if _, ok := rKmsSettings["rotationPeriod"]; ok {
				if s, ok := rKmsSettings["rotationPeriod"].(string); ok {
					r.KmsSettings.RotationPeriod = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.KmsSettings.RotationPeriod: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.KmsSettings: expected map[string]interface{}")
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
	if _, ok := u.Object["organization"]; ok {
		if s, ok := u.Object["organization"].(string); ok {
			r.Organization = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Organization: expected string")
		}
	}
	if _, ok := u.Object["provisionedResourcesParent"]; ok {
		if s, ok := u.Object["provisionedResourcesParent"].(string); ok {
			r.ProvisionedResourcesParent = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ProvisionedResourcesParent: expected string")
		}
	}
	if _, ok := u.Object["resourceSettings"]; ok {
		if s, ok := u.Object["resourceSettings"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rResourceSettings dclService.WorkloadResourceSettings
					if _, ok := objval["resourceId"]; ok {
						if s, ok := objval["resourceId"].(string); ok {
							rResourceSettings.ResourceId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rResourceSettings.ResourceId: expected string")
						}
					}
					if _, ok := objval["resourceType"]; ok {
						if s, ok := objval["resourceType"].(string); ok {
							rResourceSettings.ResourceType = dclService.WorkloadResourceSettingsResourceTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rResourceSettings.ResourceType: expected string")
						}
					}
					r.ResourceSettings = append(r.ResourceSettings, rResourceSettings)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ResourceSettings: expected []interface{}")
		}
	}
	if _, ok := u.Object["resources"]; ok {
		if s, ok := u.Object["resources"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rResources dclService.WorkloadResources
					if _, ok := objval["resourceId"]; ok {
						if i, ok := objval["resourceId"].(int64); ok {
							rResources.ResourceId = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rResources.ResourceId: expected int64")
						}
					}
					if _, ok := objval["resourceType"]; ok {
						if s, ok := objval["resourceType"].(string); ok {
							rResources.ResourceType = dclService.WorkloadResourcesResourceTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rResources.ResourceType: expected string")
						}
					}
					r.Resources = append(r.Resources, rResources)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Resources: expected []interface{}")
		}
	}
	return r, nil
}

func GetWorkload(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetWorkload(ctx, r)
	if err != nil {
		return nil, err
	}
	return WorkloadToUnstructured(r), nil
}

func ListWorkload(ctx context.Context, config *dcl.Config, organization string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListWorkload(ctx, organization, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, WorkloadToUnstructured(r))
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

func ApplyWorkload(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkload(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyWorkload(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return WorkloadToUnstructured(r), nil
}

func WorkloadHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkload(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyWorkload(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteWorkload(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return err
	}
	return c.DeleteWorkload(ctx, r)
}

func WorkloadID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Workload) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"assuredworkloads",
		"Workload",
		"beta",
	}
}

func (r *Workload) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Workload) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetWorkload(ctx, config, resource)
}

func (r *Workload) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyWorkload(ctx, config, resource, opts...)
}

func (r *Workload) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return WorkloadHasDiff(ctx, config, resource, opts...)
}

func (r *Workload) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteWorkload(ctx, config, resource)
}

func (r *Workload) ID(resource *unstructured.Resource) (string, error) {
	return WorkloadID(resource)
}

func init() {
	unstructured.Register(&Workload{})
}
