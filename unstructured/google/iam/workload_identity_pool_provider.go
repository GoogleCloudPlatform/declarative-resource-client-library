// Copyright 2025 Google LLC. All Rights Reserved.
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
package iam

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type WorkloadIdentityPoolProvider struct{}

func WorkloadIdentityPoolProviderToUnstructured(r *dclService.WorkloadIdentityPoolProvider) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "iam",
			Version: "ga",
			Type:    "WorkloadIdentityPoolProvider",
		},
		Object: make(map[string]interface{}),
	}
	if r.AttributeCondition != nil {
		u.Object["attributeCondition"] = *r.AttributeCondition
	}
	if r.AttributeMapping != nil {
		rAttributeMapping := make(map[string]interface{})
		for k, v := range r.AttributeMapping {
			rAttributeMapping[k] = v
		}
		u.Object["attributeMapping"] = rAttributeMapping
	}
	if r.Aws != nil && r.Aws != dclService.EmptyWorkloadIdentityPoolProviderAws {
		rAws := make(map[string]interface{})
		if r.Aws.AccountId != nil {
			rAws["accountId"] = *r.Aws.AccountId
		}
		var rAwsStsUri []interface{}
		for _, rAwsStsUriVal := range r.Aws.StsUri {
			rAwsStsUri = append(rAwsStsUri, rAwsStsUriVal)
		}
		rAws["stsUri"] = rAwsStsUri
		u.Object["aws"] = rAws
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Disabled != nil {
		u.Object["disabled"] = *r.Disabled
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Oidc != nil && r.Oidc != dclService.EmptyWorkloadIdentityPoolProviderOidc {
		rOidc := make(map[string]interface{})
		var rOidcAllowedAudiences []interface{}
		for _, rOidcAllowedAudiencesVal := range r.Oidc.AllowedAudiences {
			rOidcAllowedAudiences = append(rOidcAllowedAudiences, rOidcAllowedAudiencesVal)
		}
		rOidc["allowedAudiences"] = rOidcAllowedAudiences
		if r.Oidc.IssuerUri != nil {
			rOidc["issuerUri"] = *r.Oidc.IssuerUri
		}
		u.Object["oidc"] = rOidc
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.WorkloadIdentityPool != nil {
		u.Object["workloadIdentityPool"] = *r.WorkloadIdentityPool
	}
	return u
}

func UnstructuredToWorkloadIdentityPoolProvider(u *unstructured.Resource) (*dclService.WorkloadIdentityPoolProvider, error) {
	r := &dclService.WorkloadIdentityPoolProvider{}
	if _, ok := u.Object["attributeCondition"]; ok {
		if s, ok := u.Object["attributeCondition"].(string); ok {
			r.AttributeCondition = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AttributeCondition: expected string")
		}
	}
	if _, ok := u.Object["attributeMapping"]; ok {
		if rAttributeMapping, ok := u.Object["attributeMapping"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rAttributeMapping {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.AttributeMapping = m
		} else {
			return nil, fmt.Errorf("r.AttributeMapping: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["aws"]; ok {
		if rAws, ok := u.Object["aws"].(map[string]interface{}); ok {
			r.Aws = &dclService.WorkloadIdentityPoolProviderAws{}
			if _, ok := rAws["accountId"]; ok {
				if s, ok := rAws["accountId"].(string); ok {
					r.Aws.AccountId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Aws.AccountId: expected string")
				}
			}
			if _, ok := rAws["stsUri"]; ok {
				if s, ok := rAws["stsUri"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Aws.StsUri = append(r.Aws.StsUri, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Aws.StsUri: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Aws: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["disabled"]; ok {
		if b, ok := u.Object["disabled"].(bool); ok {
			r.Disabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Disabled: expected bool")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
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
	if _, ok := u.Object["oidc"]; ok {
		if rOidc, ok := u.Object["oidc"].(map[string]interface{}); ok {
			r.Oidc = &dclService.WorkloadIdentityPoolProviderOidc{}
			if _, ok := rOidc["allowedAudiences"]; ok {
				if s, ok := rOidc["allowedAudiences"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Oidc.AllowedAudiences = append(r.Oidc.AllowedAudiences, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Oidc.AllowedAudiences: expected []interface{}")
				}
			}
			if _, ok := rOidc["issuerUri"]; ok {
				if s, ok := rOidc["issuerUri"].(string); ok {
					r.Oidc.IssuerUri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Oidc.IssuerUri: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Oidc: expected map[string]interface{}")
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
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.WorkloadIdentityPoolProviderStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["workloadIdentityPool"]; ok {
		if s, ok := u.Object["workloadIdentityPool"].(string); ok {
			r.WorkloadIdentityPool = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.WorkloadIdentityPool: expected string")
		}
	}
	return r, nil
}

func GetWorkloadIdentityPoolProvider(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkloadIdentityPoolProvider(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetWorkloadIdentityPoolProvider(ctx, r)
	if err != nil {
		return nil, err
	}
	return WorkloadIdentityPoolProviderToUnstructured(r), nil
}

func ListWorkloadIdentityPoolProvider(ctx context.Context, config *dcl.Config, project string, location string, workloadIdentityPool string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListWorkloadIdentityPoolProvider(ctx, project, location, workloadIdentityPool)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, WorkloadIdentityPoolProviderToUnstructured(r))
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

func ApplyWorkloadIdentityPoolProvider(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkloadIdentityPoolProvider(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkloadIdentityPoolProvider(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyWorkloadIdentityPoolProvider(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return WorkloadIdentityPoolProviderToUnstructured(r), nil
}

func WorkloadIdentityPoolProviderHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkloadIdentityPoolProvider(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkloadIdentityPoolProvider(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyWorkloadIdentityPoolProvider(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteWorkloadIdentityPoolProvider(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkloadIdentityPoolProvider(u)
	if err != nil {
		return err
	}
	return c.DeleteWorkloadIdentityPoolProvider(ctx, r)
}

func WorkloadIdentityPoolProviderID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToWorkloadIdentityPoolProvider(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *WorkloadIdentityPoolProvider) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"iam",
		"WorkloadIdentityPoolProvider",
		"ga",
	}
}

func (r *WorkloadIdentityPoolProvider) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkloadIdentityPoolProvider) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkloadIdentityPoolProvider) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *WorkloadIdentityPoolProvider) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkloadIdentityPoolProvider) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkloadIdentityPoolProvider) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkloadIdentityPoolProvider) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetWorkloadIdentityPoolProvider(ctx, config, resource)
}

func (r *WorkloadIdentityPoolProvider) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyWorkloadIdentityPoolProvider(ctx, config, resource, opts...)
}

func (r *WorkloadIdentityPoolProvider) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return WorkloadIdentityPoolProviderHasDiff(ctx, config, resource, opts...)
}

func (r *WorkloadIdentityPoolProvider) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteWorkloadIdentityPoolProvider(ctx, config, resource)
}

func (r *WorkloadIdentityPoolProvider) ID(resource *unstructured.Resource) (string, error) {
	return WorkloadIdentityPoolProviderID(resource)
}

func init() {
	unstructured.Register(&WorkloadIdentityPoolProvider{})
}
