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
package binaryauthorization

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type Policy struct{}

func PolicyToUnstructured(r *dclService.Policy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "binaryauthorization",
			Version: "alpha",
			Type:    "Policy",
		},
		Object: make(map[string]interface{}),
	}
	var rAdmissionWhitelistPatterns []interface{}
	for _, rAdmissionWhitelistPatternsVal := range r.AdmissionWhitelistPatterns {
		rAdmissionWhitelistPatternsObject := make(map[string]interface{})
		if rAdmissionWhitelistPatternsVal.NamePattern != nil {
			rAdmissionWhitelistPatternsObject["namePattern"] = *rAdmissionWhitelistPatternsVal.NamePattern
		}
		rAdmissionWhitelistPatterns = append(rAdmissionWhitelistPatterns, rAdmissionWhitelistPatternsObject)
	}
	u.Object["admissionWhitelistPatterns"] = rAdmissionWhitelistPatterns
	if r.ClusterAdmissionRules != nil {
		rClusterAdmissionRules := make(map[string]interface{})
		for k, v := range r.ClusterAdmissionRules {
			rClusterAdmissionRules[k] = PolicyAdmissionRuleToUnstructured(&v)
		}
		u.Object["clusterAdmissionRules"] = rClusterAdmissionRules
	}
	if r.DefaultAdmissionRule != nil {
		u.Object["defaultAdmissionRule"] = PolicyAdmissionRuleToUnstructured(r.DefaultAdmissionRule)
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.GlobalPolicyEvaluationMode != nil {
		u.Object["globalPolicyEvaluationMode"] = string(*r.GlobalPolicyEvaluationMode)
	}
	if r.IstioServiceIdentityAdmissionRules != nil {
		rIstioServiceIdentityAdmissionRules := make(map[string]interface{})
		for k, v := range r.IstioServiceIdentityAdmissionRules {
			rIstioServiceIdentityAdmissionRules[k] = PolicyAdmissionRuleToUnstructured(&v)
		}
		u.Object["istioServiceIdentityAdmissionRules"] = rIstioServiceIdentityAdmissionRules
	}
	if r.KubernetesNamespaceAdmissionRules != nil {
		rKubernetesNamespaceAdmissionRules := make(map[string]interface{})
		for k, v := range r.KubernetesNamespaceAdmissionRules {
			rKubernetesNamespaceAdmissionRules[k] = PolicyAdmissionRuleToUnstructured(&v)
		}
		u.Object["kubernetesNamespaceAdmissionRules"] = rKubernetesNamespaceAdmissionRules
	}
	if r.KubernetesServiceAccountAdmissionRules != nil {
		rKubernetesServiceAccountAdmissionRules := make(map[string]interface{})
		for k, v := range r.KubernetesServiceAccountAdmissionRules {
			rKubernetesServiceAccountAdmissionRules[k] = PolicyAdmissionRuleToUnstructured(&v)
		}
		u.Object["kubernetesServiceAccountAdmissionRules"] = rKubernetesServiceAccountAdmissionRules
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func PolicyAdmissionRuleToUnstructured(r *dclService.PolicyAdmissionRule) map[string]interface{} {
	result := make(map[string]interface{})
	if r.EnforcementMode != nil {
		result["enforcementMode"] = string(*r.EnforcementMode)
	}
	if r.EvaluationMode != nil {
		result["evaluationMode"] = string(*r.EvaluationMode)
	}
	var rRequireAttestationsBy []interface{}
	for _, rRequireAttestationsByVal := range r.RequireAttestationsBy {
		rRequireAttestationsBy = append(rRequireAttestationsBy, rRequireAttestationsByVal)
	}
	result["requireAttestationsBy"] = rRequireAttestationsBy
	return result
}

func UnstructuredToPolicy(u *unstructured.Resource) (*dclService.Policy, error) {
	r := &dclService.Policy{}
	if _, ok := u.Object["admissionWhitelistPatterns"]; ok {
		if s, ok := u.Object["admissionWhitelistPatterns"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rAdmissionWhitelistPatterns dclService.PolicyAdmissionWhitelistPatterns
					if _, ok := objval["namePattern"]; ok {
						if s, ok := objval["namePattern"].(string); ok {
							rAdmissionWhitelistPatterns.NamePattern = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAdmissionWhitelistPatterns.NamePattern: expected string")
						}
					}
					r.AdmissionWhitelistPatterns = append(r.AdmissionWhitelistPatterns, rAdmissionWhitelistPatterns)
				}
			}
		} else {
			return nil, fmt.Errorf("r.AdmissionWhitelistPatterns: expected []interface{}")
		}
	}
	if _, ok := u.Object["clusterAdmissionRules"]; ok {
		if rClusterAdmissionRules, ok := u.Object["clusterAdmissionRules"].(map[string]interface{}); ok {
			m := make(map[string]dclService.PolicyAdmissionRule)
			for k, v := range rClusterAdmissionRules {
				if mapVal, ok := v.(map[string]interface{}); ok {
					unstructuredObjval, err := UnstructuredToPolicyAdmissionRule(mapVal)
					if err != nil {
						return nil, err
					}
					m[k] = *unstructuredObjval
				} else {
					return nil, fmt.Errorf("r.ClusterAdmissionRules: expected map[string]interface{}")
				}
			}
			r.ClusterAdmissionRules = m
		} else {
			return nil, fmt.Errorf("r.ClusterAdmissionRules: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["defaultAdmissionRule"]; ok {
		if rDefaultAdmissionRule, ok := u.Object["defaultAdmissionRule"].(map[string]interface{}); ok {
			var err error
			r.DefaultAdmissionRule, err = UnstructuredToPolicyAdmissionRule(rDefaultAdmissionRule)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("r.DefaultAdmissionRule: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["globalPolicyEvaluationMode"]; ok {
		if s, ok := u.Object["globalPolicyEvaluationMode"].(string); ok {
			r.GlobalPolicyEvaluationMode = dclService.PolicyGlobalPolicyEvaluationModeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.GlobalPolicyEvaluationMode: expected string")
		}
	}
	if _, ok := u.Object["istioServiceIdentityAdmissionRules"]; ok {
		if rIstioServiceIdentityAdmissionRules, ok := u.Object["istioServiceIdentityAdmissionRules"].(map[string]interface{}); ok {
			m := make(map[string]dclService.PolicyAdmissionRule)
			for k, v := range rIstioServiceIdentityAdmissionRules {
				if mapVal, ok := v.(map[string]interface{}); ok {
					unstructuredObjval, err := UnstructuredToPolicyAdmissionRule(mapVal)
					if err != nil {
						return nil, err
					}
					m[k] = *unstructuredObjval
				} else {
					return nil, fmt.Errorf("r.IstioServiceIdentityAdmissionRules: expected map[string]interface{}")
				}
			}
			r.IstioServiceIdentityAdmissionRules = m
		} else {
			return nil, fmt.Errorf("r.IstioServiceIdentityAdmissionRules: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["kubernetesNamespaceAdmissionRules"]; ok {
		if rKubernetesNamespaceAdmissionRules, ok := u.Object["kubernetesNamespaceAdmissionRules"].(map[string]interface{}); ok {
			m := make(map[string]dclService.PolicyAdmissionRule)
			for k, v := range rKubernetesNamespaceAdmissionRules {
				if mapVal, ok := v.(map[string]interface{}); ok {
					unstructuredObjval, err := UnstructuredToPolicyAdmissionRule(mapVal)
					if err != nil {
						return nil, err
					}
					m[k] = *unstructuredObjval
				} else {
					return nil, fmt.Errorf("r.KubernetesNamespaceAdmissionRules: expected map[string]interface{}")
				}
			}
			r.KubernetesNamespaceAdmissionRules = m
		} else {
			return nil, fmt.Errorf("r.KubernetesNamespaceAdmissionRules: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["kubernetesServiceAccountAdmissionRules"]; ok {
		if rKubernetesServiceAccountAdmissionRules, ok := u.Object["kubernetesServiceAccountAdmissionRules"].(map[string]interface{}); ok {
			m := make(map[string]dclService.PolicyAdmissionRule)
			for k, v := range rKubernetesServiceAccountAdmissionRules {
				if mapVal, ok := v.(map[string]interface{}); ok {
					unstructuredObjval, err := UnstructuredToPolicyAdmissionRule(mapVal)
					if err != nil {
						return nil, err
					}
					m[k] = *unstructuredObjval
				} else {
					return nil, fmt.Errorf("r.KubernetesServiceAccountAdmissionRules: expected map[string]interface{}")
				}
			}
			r.KubernetesServiceAccountAdmissionRules = m
		} else {
			return nil, fmt.Errorf("r.KubernetesServiceAccountAdmissionRules: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
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

func UnstructuredToPolicyAdmissionRule(obj map[string]interface{}) (*dclService.PolicyAdmissionRule, error) {
	r := &dclService.PolicyAdmissionRule{}
	if _, ok := obj["enforcementMode"]; ok {
		if s, ok := obj["enforcementMode"].(string); ok {
			r.EnforcementMode = dclService.PolicyAdmissionRuleEnforcementModeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.EnforcementMode: expected string")
		}
	}
	if _, ok := obj["evaluationMode"]; ok {
		if s, ok := obj["evaluationMode"].(string); ok {
			r.EvaluationMode = dclService.PolicyAdmissionRuleEvaluationModeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.EvaluationMode: expected string")
		}
	}
	if _, ok := obj["requireAttestationsBy"]; ok {
		if s, ok := obj["requireAttestationsBy"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.RequireAttestationsBy = append(r.RequireAttestationsBy, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.RequireAttestationsBy: expected []interface{}")
		}
	}
	return r, nil
}

func GetPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return PolicyToUnstructured(r), nil
}

func ApplyPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return PolicyToUnstructured(r), nil
}

func PolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeletePolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func PolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Policy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"binaryauthorization",
		"Policy",
		"alpha",
	}
}

func SetPolicyPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func GetPolicyPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policy, err := iamClient.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func (r *Policy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyPolicy(ctx, config, resource, policy)
}

func (r *Policy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyPolicy(ctx, config, resource)
}

func (r *Policy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicy(ctx, config, resource)
}

func (r *Policy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyPolicy(ctx, config, resource, opts...)
}

func (r *Policy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return PolicyHasDiff(ctx, config, resource, opts...)
}

func (r *Policy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeletePolicy(ctx, config, resource)
}

func (r *Policy) ID(resource *unstructured.Resource) (string, error) {
	return PolicyID(resource)
}

func init() {
	unstructured.Register(&Policy{})
}
