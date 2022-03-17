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
package bigqueryreservation

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type CapacityCommitment struct{}

func CapacityCommitmentToUnstructured(r *dclService.CapacityCommitment) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "bigqueryreservation",
			Version: "beta",
			Type:    "CapacityCommitment",
		},
		Object: make(map[string]interface{}),
	}
	if r.CommitmentEndTime != nil {
		u.Object["commitmentEndTime"] = *r.CommitmentEndTime
	}
	if r.CommitmentStartTime != nil {
		u.Object["commitmentStartTime"] = *r.CommitmentStartTime
	}
	if r.FailureStatus != nil && r.FailureStatus != dclService.EmptyCapacityCommitmentFailureStatus {
		rFailureStatus := make(map[string]interface{})
		if r.FailureStatus.Code != nil {
			rFailureStatus["code"] = *r.FailureStatus.Code
		}
		var rFailureStatusDetails []interface{}
		for _, rFailureStatusDetailsVal := range r.FailureStatus.Details {
			rFailureStatusDetailsObject := make(map[string]interface{})
			if rFailureStatusDetailsVal.TypeUrl != nil {
				rFailureStatusDetailsObject["typeUrl"] = *rFailureStatusDetailsVal.TypeUrl
			}
			if rFailureStatusDetailsVal.Value != nil {
				rFailureStatusDetailsObject["value"] = *rFailureStatusDetailsVal.Value
			}
			rFailureStatusDetails = append(rFailureStatusDetails, rFailureStatusDetailsObject)
		}
		rFailureStatus["details"] = rFailureStatusDetails
		if r.FailureStatus.Message != nil {
			rFailureStatus["message"] = *r.FailureStatus.Message
		}
		u.Object["failureStatus"] = rFailureStatus
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Plan != nil {
		u.Object["plan"] = string(*r.Plan)
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RenewalPlan != nil {
		u.Object["renewalPlan"] = string(*r.RenewalPlan)
	}
	if r.SlotCount != nil {
		u.Object["slotCount"] = *r.SlotCount
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	return u
}

func UnstructuredToCapacityCommitment(u *unstructured.Resource) (*dclService.CapacityCommitment, error) {
	r := &dclService.CapacityCommitment{}
	if _, ok := u.Object["commitmentEndTime"]; ok {
		if s, ok := u.Object["commitmentEndTime"].(string); ok {
			r.CommitmentEndTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CommitmentEndTime: expected string")
		}
	}
	if _, ok := u.Object["commitmentStartTime"]; ok {
		if s, ok := u.Object["commitmentStartTime"].(string); ok {
			r.CommitmentStartTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CommitmentStartTime: expected string")
		}
	}
	if _, ok := u.Object["failureStatus"]; ok {
		if rFailureStatus, ok := u.Object["failureStatus"].(map[string]interface{}); ok {
			r.FailureStatus = &dclService.CapacityCommitmentFailureStatus{}
			if _, ok := rFailureStatus["code"]; ok {
				if i, ok := rFailureStatus["code"].(int64); ok {
					r.FailureStatus.Code = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.FailureStatus.Code: expected int64")
				}
			}
			if _, ok := rFailureStatus["details"]; ok {
				if s, ok := rFailureStatus["details"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rFailureStatusDetails dclService.CapacityCommitmentFailureStatusDetails
							if _, ok := objval["typeUrl"]; ok {
								if s, ok := objval["typeUrl"].(string); ok {
									rFailureStatusDetails.TypeUrl = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rFailureStatusDetails.TypeUrl: expected string")
								}
							}
							if _, ok := objval["value"]; ok {
								if s, ok := objval["value"].(string); ok {
									rFailureStatusDetails.Value = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rFailureStatusDetails.Value: expected string")
								}
							}
							r.FailureStatus.Details = append(r.FailureStatus.Details, rFailureStatusDetails)
						}
					}
				} else {
					return nil, fmt.Errorf("r.FailureStatus.Details: expected []interface{}")
				}
			}
			if _, ok := rFailureStatus["message"]; ok {
				if s, ok := rFailureStatus["message"].(string); ok {
					r.FailureStatus.Message = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.FailureStatus.Message: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.FailureStatus: expected map[string]interface{}")
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
	if _, ok := u.Object["plan"]; ok {
		if s, ok := u.Object["plan"].(string); ok {
			r.Plan = dclService.CapacityCommitmentPlanEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Plan: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["renewalPlan"]; ok {
		if s, ok := u.Object["renewalPlan"].(string); ok {
			r.RenewalPlan = dclService.CapacityCommitmentRenewalPlanEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.RenewalPlan: expected string")
		}
	}
	if _, ok := u.Object["slotCount"]; ok {
		if i, ok := u.Object["slotCount"].(int64); ok {
			r.SlotCount = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.SlotCount: expected int64")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.CapacityCommitmentStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	return r, nil
}

func GetCapacityCommitment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCapacityCommitment(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetCapacityCommitment(ctx, r)
	if err != nil {
		return nil, err
	}
	return CapacityCommitmentToUnstructured(r), nil
}

func ListCapacityCommitment(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListCapacityCommitment(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, CapacityCommitmentToUnstructured(r))
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

func ApplyCapacityCommitment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCapacityCommitment(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCapacityCommitment(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyCapacityCommitment(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return CapacityCommitmentToUnstructured(r), nil
}

func CapacityCommitmentHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCapacityCommitment(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCapacityCommitment(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyCapacityCommitment(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteCapacityCommitment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCapacityCommitment(u)
	if err != nil {
		return err
	}
	return c.DeleteCapacityCommitment(ctx, r)
}

func CapacityCommitmentID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToCapacityCommitment(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *CapacityCommitment) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"bigqueryreservation",
		"CapacityCommitment",
		"beta",
	}
}

func (r *CapacityCommitment) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CapacityCommitment) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CapacityCommitment) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CapacityCommitment) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CapacityCommitment) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CapacityCommitment) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetCapacityCommitment(ctx, config, resource)
}

func (r *CapacityCommitment) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyCapacityCommitment(ctx, config, resource, opts...)
}

func (r *CapacityCommitment) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return CapacityCommitmentHasDiff(ctx, config, resource, opts...)
}

func (r *CapacityCommitment) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteCapacityCommitment(ctx, config, resource)
}

func (r *CapacityCommitment) ID(resource *unstructured.Resource) (string, error) {
	return CapacityCommitmentID(resource)
}

func init() {
	unstructured.Register(&CapacityCommitment{})
}
