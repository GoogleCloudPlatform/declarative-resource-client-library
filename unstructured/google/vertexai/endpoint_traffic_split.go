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

type EndpointTrafficSplit struct{}

func EndpointTrafficSplitToUnstructured(r *dclService.EndpointTrafficSplit) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "vertexai",
			Version: "ga",
			Type:    "EndpointTrafficSplit",
		},
		Object: make(map[string]interface{}),
	}
	if r.Endpoint != nil {
		u.Object["endpoint"] = *r.Endpoint
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	var rTrafficSplit []interface{}
	for _, rTrafficSplitVal := range r.TrafficSplit {
		rTrafficSplitObject := make(map[string]interface{})
		if rTrafficSplitVal.DeployedModelId != nil {
			rTrafficSplitObject["deployedModelId"] = *rTrafficSplitVal.DeployedModelId
		}
		if rTrafficSplitVal.TrafficPercentage != nil {
			rTrafficSplitObject["trafficPercentage"] = *rTrafficSplitVal.TrafficPercentage
		}
		rTrafficSplit = append(rTrafficSplit, rTrafficSplitObject)
	}
	u.Object["trafficSplit"] = rTrafficSplit
	return u
}

func UnstructuredToEndpointTrafficSplit(u *unstructured.Resource) (*dclService.EndpointTrafficSplit, error) {
	r := &dclService.EndpointTrafficSplit{}
	if _, ok := u.Object["endpoint"]; ok {
		if s, ok := u.Object["endpoint"].(string); ok {
			r.Endpoint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Endpoint: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["trafficSplit"]; ok {
		if s, ok := u.Object["trafficSplit"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rTrafficSplit dclService.EndpointTrafficSplitTrafficSplit
					if _, ok := objval["deployedModelId"]; ok {
						if s, ok := objval["deployedModelId"].(string); ok {
							rTrafficSplit.DeployedModelId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rTrafficSplit.DeployedModelId: expected string")
						}
					}
					if _, ok := objval["trafficPercentage"]; ok {
						if i, ok := objval["trafficPercentage"].(int64); ok {
							rTrafficSplit.TrafficPercentage = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rTrafficSplit.TrafficPercentage: expected int64")
						}
					}
					r.TrafficSplit = append(r.TrafficSplit, rTrafficSplit)
				}
			}
		} else {
			return nil, fmt.Errorf("r.TrafficSplit: expected []interface{}")
		}
	}
	return r, nil
}

func GetEndpointTrafficSplit(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpointTrafficSplit(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetEndpointTrafficSplit(ctx, r)
	if err != nil {
		return nil, err
	}
	return EndpointTrafficSplitToUnstructured(r), nil
}

func ApplyEndpointTrafficSplit(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpointTrafficSplit(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEndpointTrafficSplit(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyEndpointTrafficSplit(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return EndpointTrafficSplitToUnstructured(r), nil
}

func EndpointTrafficSplitHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpointTrafficSplit(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEndpointTrafficSplit(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyEndpointTrafficSplit(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteEndpointTrafficSplit(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpointTrafficSplit(u)
	if err != nil {
		return err
	}
	return c.DeleteEndpointTrafficSplit(ctx, r)
}

func EndpointTrafficSplitID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToEndpointTrafficSplit(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *EndpointTrafficSplit) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"vertexai",
		"EndpointTrafficSplit",
		"ga",
	}
}

func (r *EndpointTrafficSplit) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointTrafficSplit) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointTrafficSplit) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *EndpointTrafficSplit) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointTrafficSplit) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointTrafficSplit) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointTrafficSplit) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetEndpointTrafficSplit(ctx, config, resource)
}

func (r *EndpointTrafficSplit) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyEndpointTrafficSplit(ctx, config, resource, opts...)
}

func (r *EndpointTrafficSplit) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return EndpointTrafficSplitHasDiff(ctx, config, resource, opts...)
}

func (r *EndpointTrafficSplit) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteEndpointTrafficSplit(ctx, config, resource)
}

func (r *EndpointTrafficSplit) ID(resource *unstructured.Resource) (string, error) {
	return EndpointTrafficSplitID(resource)
}

func init() {
	unstructured.Register(&EndpointTrafficSplit{})
}
