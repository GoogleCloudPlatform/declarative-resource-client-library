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
package eventarc

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Trigger struct{}

func TriggerToUnstructured(r *dclService.Trigger) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "eventarc",
			Version: "ga",
			Type:    "Trigger",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Destination != nil && r.Destination != dclService.EmptyTriggerDestination {
		rDestination := make(map[string]interface{})
		if r.Destination.CloudFunction != nil {
			rDestination["cloudFunction"] = *r.Destination.CloudFunction
		}
		if r.Destination.CloudRun != nil && r.Destination.CloudRun != dclService.EmptyTriggerDestinationCloudRun {
			rDestinationCloudRun := make(map[string]interface{})
			if r.Destination.CloudRun.Path != nil {
				rDestinationCloudRun["path"] = *r.Destination.CloudRun.Path
			}
			if r.Destination.CloudRun.Region != nil {
				rDestinationCloudRun["region"] = *r.Destination.CloudRun.Region
			}
			if r.Destination.CloudRun.Service != nil {
				rDestinationCloudRun["service"] = *r.Destination.CloudRun.Service
			}
			rDestination["cloudRun"] = rDestinationCloudRun
		}
		if r.Destination.Gke != nil && r.Destination.Gke != dclService.EmptyTriggerDestinationGke {
			rDestinationGke := make(map[string]interface{})
			if r.Destination.Gke.Cluster != nil {
				rDestinationGke["cluster"] = *r.Destination.Gke.Cluster
			}
			if r.Destination.Gke.Location != nil {
				rDestinationGke["location"] = *r.Destination.Gke.Location
			}
			if r.Destination.Gke.Namespace != nil {
				rDestinationGke["namespace"] = *r.Destination.Gke.Namespace
			}
			if r.Destination.Gke.Path != nil {
				rDestinationGke["path"] = *r.Destination.Gke.Path
			}
			if r.Destination.Gke.Service != nil {
				rDestinationGke["service"] = *r.Destination.Gke.Service
			}
			rDestination["gke"] = rDestinationGke
		}
		if r.Destination.Workflow != nil {
			rDestination["workflow"] = *r.Destination.Workflow
		}
		u.Object["destination"] = rDestination
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	var rEventFilters []interface{}
	for _, rEventFiltersVal := range r.EventFilters {
		rEventFiltersObject := make(map[string]interface{})
		if rEventFiltersVal.Attribute != nil {
			rEventFiltersObject["attribute"] = *rEventFiltersVal.Attribute
		}
		if rEventFiltersVal.Operator != nil {
			rEventFiltersObject["operator"] = *rEventFiltersVal.Operator
		}
		if rEventFiltersVal.Value != nil {
			rEventFiltersObject["value"] = *rEventFiltersVal.Value
		}
		rEventFilters = append(rEventFilters, rEventFiltersObject)
	}
	u.Object["eventFilters"] = rEventFilters
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
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ServiceAccount != nil {
		u.Object["serviceAccount"] = *r.ServiceAccount
	}
	if r.Transport != nil && r.Transport != dclService.EmptyTriggerTransport {
		rTransport := make(map[string]interface{})
		if r.Transport.Pubsub != nil && r.Transport.Pubsub != dclService.EmptyTriggerTransportPubsub {
			rTransportPubsub := make(map[string]interface{})
			if r.Transport.Pubsub.Subscription != nil {
				rTransportPubsub["subscription"] = *r.Transport.Pubsub.Subscription
			}
			if r.Transport.Pubsub.Topic != nil {
				rTransportPubsub["topic"] = *r.Transport.Pubsub.Topic
			}
			rTransport["pubsub"] = rTransportPubsub
		}
		u.Object["transport"] = rTransport
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToTrigger(u *unstructured.Resource) (*dclService.Trigger, error) {
	r := &dclService.Trigger{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["destination"]; ok {
		if rDestination, ok := u.Object["destination"].(map[string]interface{}); ok {
			r.Destination = &dclService.TriggerDestination{}
			if _, ok := rDestination["cloudFunction"]; ok {
				if s, ok := rDestination["cloudFunction"].(string); ok {
					r.Destination.CloudFunction = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Destination.CloudFunction: expected string")
				}
			}
			if _, ok := rDestination["cloudRun"]; ok {
				if rDestinationCloudRun, ok := rDestination["cloudRun"].(map[string]interface{}); ok {
					r.Destination.CloudRun = &dclService.TriggerDestinationCloudRun{}
					if _, ok := rDestinationCloudRun["path"]; ok {
						if s, ok := rDestinationCloudRun["path"].(string); ok {
							r.Destination.CloudRun.Path = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Destination.CloudRun.Path: expected string")
						}
					}
					if _, ok := rDestinationCloudRun["region"]; ok {
						if s, ok := rDestinationCloudRun["region"].(string); ok {
							r.Destination.CloudRun.Region = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Destination.CloudRun.Region: expected string")
						}
					}
					if _, ok := rDestinationCloudRun["service"]; ok {
						if s, ok := rDestinationCloudRun["service"].(string); ok {
							r.Destination.CloudRun.Service = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Destination.CloudRun.Service: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Destination.CloudRun: expected map[string]interface{}")
				}
			}
			if _, ok := rDestination["gke"]; ok {
				if rDestinationGke, ok := rDestination["gke"].(map[string]interface{}); ok {
					r.Destination.Gke = &dclService.TriggerDestinationGke{}
					if _, ok := rDestinationGke["cluster"]; ok {
						if s, ok := rDestinationGke["cluster"].(string); ok {
							r.Destination.Gke.Cluster = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Destination.Gke.Cluster: expected string")
						}
					}
					if _, ok := rDestinationGke["location"]; ok {
						if s, ok := rDestinationGke["location"].(string); ok {
							r.Destination.Gke.Location = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Destination.Gke.Location: expected string")
						}
					}
					if _, ok := rDestinationGke["namespace"]; ok {
						if s, ok := rDestinationGke["namespace"].(string); ok {
							r.Destination.Gke.Namespace = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Destination.Gke.Namespace: expected string")
						}
					}
					if _, ok := rDestinationGke["path"]; ok {
						if s, ok := rDestinationGke["path"].(string); ok {
							r.Destination.Gke.Path = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Destination.Gke.Path: expected string")
						}
					}
					if _, ok := rDestinationGke["service"]; ok {
						if s, ok := rDestinationGke["service"].(string); ok {
							r.Destination.Gke.Service = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Destination.Gke.Service: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Destination.Gke: expected map[string]interface{}")
				}
			}
			if _, ok := rDestination["workflow"]; ok {
				if s, ok := rDestination["workflow"].(string); ok {
					r.Destination.Workflow = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Destination.Workflow: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Destination: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["eventFilters"]; ok {
		if s, ok := u.Object["eventFilters"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rEventFilters dclService.TriggerEventFilters
					if _, ok := objval["attribute"]; ok {
						if s, ok := objval["attribute"].(string); ok {
							rEventFilters.Attribute = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rEventFilters.Attribute: expected string")
						}
					}
					if _, ok := objval["operator"]; ok {
						if s, ok := objval["operator"].(string); ok {
							rEventFilters.Operator = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rEventFilters.Operator: expected string")
						}
					}
					if _, ok := objval["value"]; ok {
						if s, ok := objval["value"].(string); ok {
							rEventFilters.Value = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rEventFilters.Value: expected string")
						}
					}
					r.EventFilters = append(r.EventFilters, rEventFilters)
				}
			}
		} else {
			return nil, fmt.Errorf("r.EventFilters: expected []interface{}")
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
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["serviceAccount"]; ok {
		if s, ok := u.Object["serviceAccount"].(string); ok {
			r.ServiceAccount = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ServiceAccount: expected string")
		}
	}
	if _, ok := u.Object["transport"]; ok {
		if rTransport, ok := u.Object["transport"].(map[string]interface{}); ok {
			r.Transport = &dclService.TriggerTransport{}
			if _, ok := rTransport["pubsub"]; ok {
				if rTransportPubsub, ok := rTransport["pubsub"].(map[string]interface{}); ok {
					r.Transport.Pubsub = &dclService.TriggerTransportPubsub{}
					if _, ok := rTransportPubsub["subscription"]; ok {
						if s, ok := rTransportPubsub["subscription"].(string); ok {
							r.Transport.Pubsub.Subscription = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Transport.Pubsub.Subscription: expected string")
						}
					}
					if _, ok := rTransportPubsub["topic"]; ok {
						if s, ok := rTransportPubsub["topic"].(string); ok {
							r.Transport.Pubsub.Topic = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Transport.Pubsub.Topic: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Transport.Pubsub: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Transport: expected map[string]interface{}")
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

func GetTrigger(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTrigger(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetTrigger(ctx, r)
	if err != nil {
		return nil, err
	}
	return TriggerToUnstructured(r), nil
}

func ListTrigger(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListTrigger(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, TriggerToUnstructured(r))
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

func ApplyTrigger(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTrigger(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTrigger(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyTrigger(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return TriggerToUnstructured(r), nil
}

func TriggerHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTrigger(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTrigger(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyTrigger(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteTrigger(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTrigger(u)
	if err != nil {
		return err
	}
	return c.DeleteTrigger(ctx, r)
}

func TriggerID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToTrigger(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Trigger) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"eventarc",
		"Trigger",
		"ga",
	}
}

func (r *Trigger) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Trigger) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Trigger) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Trigger) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Trigger) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Trigger) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Trigger) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetTrigger(ctx, config, resource)
}

func (r *Trigger) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyTrigger(ctx, config, resource, opts...)
}

func (r *Trigger) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return TriggerHasDiff(ctx, config, resource, opts...)
}

func (r *Trigger) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteTrigger(ctx, config, resource)
}

func (r *Trigger) ID(resource *unstructured.Resource) (string, error) {
	return TriggerID(resource)
}

func init() {
	unstructured.Register(&Trigger{})
}
