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
package clouddeploy

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Target struct{}

func TargetToUnstructured(r *dclService.Target) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "clouddeploy",
			Version: "ga",
			Type:    "Target",
		},
		Object: make(map[string]interface{}),
	}
	if r.Annotations != nil {
		rAnnotations := make(map[string]interface{})
		for k, v := range r.Annotations {
			rAnnotations[k] = v
		}
		u.Object["annotations"] = rAnnotations
	}
	if r.AnthosCluster != nil && r.AnthosCluster != dclService.EmptyTargetAnthosCluster {
		rAnthosCluster := make(map[string]interface{})
		if r.AnthosCluster.Membership != nil {
			rAnthosCluster["membership"] = *r.AnthosCluster.Membership
		}
		u.Object["anthosCluster"] = rAnthosCluster
	}
	if r.AssociatedEntities != nil {
		rAssociatedEntities := make(map[string]interface{})
		for k, v := range r.AssociatedEntities {
			rAssociatedEntitiesMap := make(map[string]interface{})
			var vAnthosClusters []interface{}
			for _, vAnthosClustersVal := range v.AnthosClusters {
				vAnthosClustersObject := make(map[string]interface{})
				if vAnthosClustersVal.Membership != nil {
					vAnthosClustersObject["membership"] = *vAnthosClustersVal.Membership
				}
				vAnthosClusters = append(vAnthosClusters, vAnthosClustersObject)
			}
			rAssociatedEntitiesMap["anthosClusters"] = vAnthosClusters
			var vGkeClusters []interface{}
			for _, vGkeClustersVal := range v.GkeClusters {
				vGkeClustersObject := make(map[string]interface{})
				if vGkeClustersVal.Cluster != nil {
					vGkeClustersObject["cluster"] = *vGkeClustersVal.Cluster
				}
				if vGkeClustersVal.InternalIP != nil {
					vGkeClustersObject["internalIP"] = *vGkeClustersVal.InternalIP
				}
				if vGkeClustersVal.ProxyUrl != nil {
					vGkeClustersObject["proxyUrl"] = *vGkeClustersVal.ProxyUrl
				}
				vGkeClusters = append(vGkeClusters, vGkeClustersObject)
			}
			rAssociatedEntitiesMap["gkeClusters"] = vGkeClusters
			rAssociatedEntities[k] = rAssociatedEntitiesMap
		}
		u.Object["associatedEntities"] = rAssociatedEntities
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.CustomTarget != nil && r.CustomTarget != dclService.EmptyTargetCustomTarget {
		rCustomTarget := make(map[string]interface{})
		if r.CustomTarget.CustomTargetType != nil {
			rCustomTarget["customTargetType"] = *r.CustomTarget.CustomTargetType
		}
		u.Object["customTarget"] = rCustomTarget
	}
	if r.DeployParameters != nil {
		rDeployParameters := make(map[string]interface{})
		for k, v := range r.DeployParameters {
			rDeployParameters[k] = v
		}
		u.Object["deployParameters"] = rDeployParameters
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	var rExecutionConfigs []interface{}
	for _, rExecutionConfigsVal := range r.ExecutionConfigs {
		rExecutionConfigsObject := make(map[string]interface{})
		if rExecutionConfigsVal.ArtifactStorage != nil {
			rExecutionConfigsObject["artifactStorage"] = *rExecutionConfigsVal.ArtifactStorage
		}
		if rExecutionConfigsVal.ExecutionTimeout != nil {
			rExecutionConfigsObject["executionTimeout"] = *rExecutionConfigsVal.ExecutionTimeout
		}
		if rExecutionConfigsVal.ServiceAccount != nil {
			rExecutionConfigsObject["serviceAccount"] = *rExecutionConfigsVal.ServiceAccount
		}
		var rExecutionConfigsValUsages []interface{}
		for _, rExecutionConfigsValUsagesVal := range rExecutionConfigsVal.Usages {
			rExecutionConfigsValUsages = append(rExecutionConfigsValUsages, string(rExecutionConfigsValUsagesVal))
		}
		rExecutionConfigsObject["usages"] = rExecutionConfigsValUsages
		if rExecutionConfigsVal.Verbose != nil {
			rExecutionConfigsObject["verbose"] = *rExecutionConfigsVal.Verbose
		}
		if rExecutionConfigsVal.WorkerPool != nil {
			rExecutionConfigsObject["workerPool"] = *rExecutionConfigsVal.WorkerPool
		}
		rExecutionConfigs = append(rExecutionConfigs, rExecutionConfigsObject)
	}
	u.Object["executionConfigs"] = rExecutionConfigs
	if r.Gke != nil && r.Gke != dclService.EmptyTargetGke {
		rGke := make(map[string]interface{})
		if r.Gke.Cluster != nil {
			rGke["cluster"] = *r.Gke.Cluster
		}
		if r.Gke.DnsEndpoint != nil {
			rGke["dnsEndpoint"] = *r.Gke.DnsEndpoint
		}
		if r.Gke.InternalIP != nil {
			rGke["internalIP"] = *r.Gke.InternalIP
		}
		if r.Gke.ProxyUrl != nil {
			rGke["proxyUrl"] = *r.Gke.ProxyUrl
		}
		u.Object["gke"] = rGke
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
	if r.MultiTarget != nil && r.MultiTarget != dclService.EmptyTargetMultiTarget {
		rMultiTarget := make(map[string]interface{})
		var rMultiTargetTargetIds []interface{}
		for _, rMultiTargetTargetIdsVal := range r.MultiTarget.TargetIds {
			rMultiTargetTargetIds = append(rMultiTargetTargetIds, rMultiTargetTargetIdsVal)
		}
		rMultiTarget["targetIds"] = rMultiTargetTargetIds
		u.Object["multiTarget"] = rMultiTarget
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RequireApproval != nil {
		u.Object["requireApproval"] = *r.RequireApproval
	}
	if r.Run != nil && r.Run != dclService.EmptyTargetRun {
		rRun := make(map[string]interface{})
		if r.Run.Location != nil {
			rRun["location"] = *r.Run.Location
		}
		u.Object["run"] = rRun
	}
	if r.TargetId != nil {
		u.Object["targetId"] = *r.TargetId
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToTarget(u *unstructured.Resource) (*dclService.Target, error) {
	r := &dclService.Target{}
	if _, ok := u.Object["annotations"]; ok {
		if rAnnotations, ok := u.Object["annotations"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rAnnotations {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Annotations = m
		} else {
			return nil, fmt.Errorf("r.Annotations: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["anthosCluster"]; ok {
		if rAnthosCluster, ok := u.Object["anthosCluster"].(map[string]interface{}); ok {
			r.AnthosCluster = &dclService.TargetAnthosCluster{}
			if _, ok := rAnthosCluster["membership"]; ok {
				if s, ok := rAnthosCluster["membership"].(string); ok {
					r.AnthosCluster.Membership = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AnthosCluster.Membership: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AnthosCluster: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["associatedEntities"]; ok {
		if rAssociatedEntities, ok := u.Object["associatedEntities"].(map[string]interface{}); ok {
			m := make(map[string]dclService.TargetAssociatedEntities)
			for k, v := range rAssociatedEntities {
				if objval, ok := v.(map[string]interface{}); ok {
					var rAssociatedEntitiesObj dclService.TargetAssociatedEntities
					if _, ok := objval["anthosClusters"]; ok {
						if s, ok := objval["anthosClusters"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rAssociatedEntitiesObjAnthosClusters dclService.TargetAssociatedEntitiesAnthosClusters
									if _, ok := objval["membership"]; ok {
										if s, ok := objval["membership"].(string); ok {
											rAssociatedEntitiesObjAnthosClusters.Membership = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rAssociatedEntitiesObjAnthosClusters.Membership: expected string")
										}
									}
									rAssociatedEntitiesObj.AnthosClusters = append(rAssociatedEntitiesObj.AnthosClusters, rAssociatedEntitiesObjAnthosClusters)
								}
							}
						} else {
							return nil, fmt.Errorf("rAssociatedEntitiesObj.AnthosClusters: expected []interface{}")
						}
					}
					if _, ok := objval["gkeClusters"]; ok {
						if s, ok := objval["gkeClusters"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rAssociatedEntitiesObjGkeClusters dclService.TargetAssociatedEntitiesGkeClusters
									if _, ok := objval["cluster"]; ok {
										if s, ok := objval["cluster"].(string); ok {
											rAssociatedEntitiesObjGkeClusters.Cluster = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rAssociatedEntitiesObjGkeClusters.Cluster: expected string")
										}
									}
									if _, ok := objval["internalIP"]; ok {
										if b, ok := objval["internalIP"].(bool); ok {
											rAssociatedEntitiesObjGkeClusters.InternalIP = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rAssociatedEntitiesObjGkeClusters.InternalIP: expected bool")
										}
									}
									if _, ok := objval["proxyUrl"]; ok {
										if s, ok := objval["proxyUrl"].(string); ok {
											rAssociatedEntitiesObjGkeClusters.ProxyUrl = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rAssociatedEntitiesObjGkeClusters.ProxyUrl: expected string")
										}
									}
									rAssociatedEntitiesObj.GkeClusters = append(rAssociatedEntitiesObj.GkeClusters, rAssociatedEntitiesObjGkeClusters)
								}
							}
						} else {
							return nil, fmt.Errorf("rAssociatedEntitiesObj.GkeClusters: expected []interface{}")
						}
					}
					m[k] = rAssociatedEntitiesObj
				} else {
					return nil, fmt.Errorf("r.AssociatedEntities: expected map[string]interface{}")
				}
			}
			r.AssociatedEntities = m
		} else {
			return nil, fmt.Errorf("r.AssociatedEntities: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["customTarget"]; ok {
		if rCustomTarget, ok := u.Object["customTarget"].(map[string]interface{}); ok {
			r.CustomTarget = &dclService.TargetCustomTarget{}
			if _, ok := rCustomTarget["customTargetType"]; ok {
				if s, ok := rCustomTarget["customTargetType"].(string); ok {
					r.CustomTarget.CustomTargetType = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.CustomTarget.CustomTargetType: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.CustomTarget: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["deployParameters"]; ok {
		if rDeployParameters, ok := u.Object["deployParameters"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rDeployParameters {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.DeployParameters = m
		} else {
			return nil, fmt.Errorf("r.DeployParameters: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["executionConfigs"]; ok {
		if s, ok := u.Object["executionConfigs"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rExecutionConfigs dclService.TargetExecutionConfigs
					if _, ok := objval["artifactStorage"]; ok {
						if s, ok := objval["artifactStorage"].(string); ok {
							rExecutionConfigs.ArtifactStorage = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rExecutionConfigs.ArtifactStorage: expected string")
						}
					}
					if _, ok := objval["executionTimeout"]; ok {
						if s, ok := objval["executionTimeout"].(string); ok {
							rExecutionConfigs.ExecutionTimeout = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rExecutionConfigs.ExecutionTimeout: expected string")
						}
					}
					if _, ok := objval["serviceAccount"]; ok {
						if s, ok := objval["serviceAccount"].(string); ok {
							rExecutionConfigs.ServiceAccount = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rExecutionConfigs.ServiceAccount: expected string")
						}
					}
					if _, ok := objval["usages"]; ok {
						if s, ok := objval["usages"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rExecutionConfigs.Usages = append(rExecutionConfigs.Usages, dclService.TargetExecutionConfigsUsagesEnum(strval))
								}
							}
						} else {
							return nil, fmt.Errorf("rExecutionConfigs.Usages: expected []interface{}")
						}
					}
					if _, ok := objval["verbose"]; ok {
						if b, ok := objval["verbose"].(bool); ok {
							rExecutionConfigs.Verbose = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("rExecutionConfigs.Verbose: expected bool")
						}
					}
					if _, ok := objval["workerPool"]; ok {
						if s, ok := objval["workerPool"].(string); ok {
							rExecutionConfigs.WorkerPool = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rExecutionConfigs.WorkerPool: expected string")
						}
					}
					r.ExecutionConfigs = append(r.ExecutionConfigs, rExecutionConfigs)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ExecutionConfigs: expected []interface{}")
		}
	}
	if _, ok := u.Object["gke"]; ok {
		if rGke, ok := u.Object["gke"].(map[string]interface{}); ok {
			r.Gke = &dclService.TargetGke{}
			if _, ok := rGke["cluster"]; ok {
				if s, ok := rGke["cluster"].(string); ok {
					r.Gke.Cluster = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Gke.Cluster: expected string")
				}
			}
			if _, ok := rGke["dnsEndpoint"]; ok {
				if b, ok := rGke["dnsEndpoint"].(bool); ok {
					r.Gke.DnsEndpoint = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Gke.DnsEndpoint: expected bool")
				}
			}
			if _, ok := rGke["internalIP"]; ok {
				if b, ok := rGke["internalIP"].(bool); ok {
					r.Gke.InternalIP = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Gke.InternalIP: expected bool")
				}
			}
			if _, ok := rGke["proxyUrl"]; ok {
				if s, ok := rGke["proxyUrl"].(string); ok {
					r.Gke.ProxyUrl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Gke.ProxyUrl: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Gke: expected map[string]interface{}")
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
	if _, ok := u.Object["multiTarget"]; ok {
		if rMultiTarget, ok := u.Object["multiTarget"].(map[string]interface{}); ok {
			r.MultiTarget = &dclService.TargetMultiTarget{}
			if _, ok := rMultiTarget["targetIds"]; ok {
				if s, ok := rMultiTarget["targetIds"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.MultiTarget.TargetIds = append(r.MultiTarget.TargetIds, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MultiTarget.TargetIds: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MultiTarget: expected map[string]interface{}")
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
	if _, ok := u.Object["requireApproval"]; ok {
		if b, ok := u.Object["requireApproval"].(bool); ok {
			r.RequireApproval = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.RequireApproval: expected bool")
		}
	}
	if _, ok := u.Object["run"]; ok {
		if rRun, ok := u.Object["run"].(map[string]interface{}); ok {
			r.Run = &dclService.TargetRun{}
			if _, ok := rRun["location"]; ok {
				if s, ok := rRun["location"].(string); ok {
					r.Run.Location = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Run.Location: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Run: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["targetId"]; ok {
		if s, ok := u.Object["targetId"].(string); ok {
			r.TargetId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.TargetId: expected string")
		}
	}
	if _, ok := u.Object["uid"]; ok {
		if s, ok := u.Object["uid"].(string); ok {
			r.Uid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uid: expected string")
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

func GetTarget(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTarget(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetTarget(ctx, r)
	if err != nil {
		return nil, err
	}
	return TargetToUnstructured(r), nil
}

func ListTarget(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListTarget(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, TargetToUnstructured(r))
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

func ApplyTarget(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTarget(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTarget(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyTarget(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return TargetToUnstructured(r), nil
}

func TargetHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTarget(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTarget(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyTarget(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteTarget(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTarget(u)
	if err != nil {
		return err
	}
	return c.DeleteTarget(ctx, r)
}

func TargetID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToTarget(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Target) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"clouddeploy",
		"Target",
		"ga",
	}
}

func (r *Target) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Target) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Target) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Target) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Target) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Target) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Target) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetTarget(ctx, config, resource)
}

func (r *Target) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyTarget(ctx, config, resource, opts...)
}

func (r *Target) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return TargetHasDiff(ctx, config, resource, opts...)
}

func (r *Target) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteTarget(ctx, config, resource)
}

func (r *Target) ID(resource *unstructured.Resource) (string, error) {
	return TargetID(resource)
}

func init() {
	unstructured.Register(&Target{})
}
