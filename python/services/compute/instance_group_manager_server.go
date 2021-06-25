// Copyright 2021 Google LLC. All Rights Reserved.
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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for InstanceGroupManager.
type InstanceGroupManagerServer struct{}

// ProtoToInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(e computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) *compute.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(n[len("ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMinimalActionEnum converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(e computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum) *compute.InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerUpdatePolicyMinimalActionEnum(n[len("ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerDistributionPolicy converts a InstanceGroupManagerDistributionPolicy resource from its proto representation.
func ProtoToComputeInstanceGroupManagerDistributionPolicy(p *computepb.ComputeInstanceGroupManagerDistributionPolicy) *compute.InstanceGroupManagerDistributionPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerDistributionPolicy{}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, *ProtoToComputeInstanceGroupManagerDistributionPolicyZones(r))
	}
	return obj
}

// ProtoToInstanceGroupManagerDistributionPolicyZones converts a InstanceGroupManagerDistributionPolicyZones resource from its proto representation.
func ProtoToComputeInstanceGroupManagerDistributionPolicyZones(p *computepb.ComputeInstanceGroupManagerDistributionPolicyZones) *compute.InstanceGroupManagerDistributionPolicyZones {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerDistributionPolicyZones{
		Zone: dcl.StringOrNil(p.Zone),
	}
	return obj
}

// ProtoToInstanceGroupManagerCurrentActions converts a InstanceGroupManagerCurrentActions resource from its proto representation.
func ProtoToComputeInstanceGroupManagerCurrentActions(p *computepb.ComputeInstanceGroupManagerCurrentActions) *compute.InstanceGroupManagerCurrentActions {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerCurrentActions{
		Abandoning:             dcl.Int64OrNil(p.Abandoning),
		Creating:               dcl.Int64OrNil(p.Creating),
		CreatingWithoutRetries: dcl.Int64OrNil(p.CreatingWithoutRetries),
		Deleting:               dcl.Int64OrNil(p.Deleting),
		None:                   dcl.Int64OrNil(p.None),
		Recreating:             dcl.Int64OrNil(p.Recreating),
		Refreshing:             dcl.Int64OrNil(p.Refreshing),
		Restarting:             dcl.Int64OrNil(p.Restarting),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersions converts a InstanceGroupManagerVersions resource from its proto representation.
func ProtoToComputeInstanceGroupManagerVersions(p *computepb.ComputeInstanceGroupManagerVersions) *compute.InstanceGroupManagerVersions {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerVersions{
		Name:             dcl.StringOrNil(p.Name),
		InstanceTemplate: dcl.StringOrNil(p.InstanceTemplate),
		TargetSize:       ProtoToComputeInstanceGroupManagerVersionsTargetSize(p.GetTargetSize()),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersionsTargetSize converts a InstanceGroupManagerVersionsTargetSize resource from its proto representation.
func ProtoToComputeInstanceGroupManagerVersionsTargetSize(p *computepb.ComputeInstanceGroupManagerVersionsTargetSize) *compute.InstanceGroupManagerVersionsTargetSize {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerVersionsTargetSize{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToInstanceGroupManagerNamedPorts converts a InstanceGroupManagerNamedPorts resource from its proto representation.
func ProtoToComputeInstanceGroupManagerNamedPorts(p *computepb.ComputeInstanceGroupManagerNamedPorts) *compute.InstanceGroupManagerNamedPorts {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerNamedPorts{
		Name: dcl.StringOrNil(p.Name),
		Port: dcl.Int64OrNil(p.Port),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatus converts a InstanceGroupManagerStatus resource from its proto representation.
func ProtoToComputeInstanceGroupManagerStatus(p *computepb.ComputeInstanceGroupManagerStatus) *compute.InstanceGroupManagerStatus {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatus{
		IsStable:      dcl.Bool(p.IsStable),
		VersionTarget: ProtoToComputeInstanceGroupManagerStatusVersionTarget(p.GetVersionTarget()),
		Autoscalar:    dcl.StringOrNil(p.Autoscalar),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusVersionTarget converts a InstanceGroupManagerStatusVersionTarget resource from its proto representation.
func ProtoToComputeInstanceGroupManagerStatusVersionTarget(p *computepb.ComputeInstanceGroupManagerStatusVersionTarget) *compute.InstanceGroupManagerStatusVersionTarget {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatusVersionTarget{
		IsReached: dcl.Bool(p.IsReached),
	}
	return obj
}

// ProtoToInstanceGroupManagerAutoHealingPolicies converts a InstanceGroupManagerAutoHealingPolicies resource from its proto representation.
func ProtoToComputeInstanceGroupManagerAutoHealingPolicies(p *computepb.ComputeInstanceGroupManagerAutoHealingPolicies) *compute.InstanceGroupManagerAutoHealingPolicies {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerAutoHealingPolicies{
		HealthCheck:     dcl.StringOrNil(p.HealthCheck),
		InitialDelaySec: dcl.Int64OrNil(p.InitialDelaySec),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicy converts a InstanceGroupManagerUpdatePolicy resource from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicy(p *computepb.ComputeInstanceGroupManagerUpdatePolicy) *compute.InstanceGroupManagerUpdatePolicy {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerUpdatePolicy{
		InstanceRedistributionType: ProtoToComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(p.GetInstanceRedistributionType()),
		MinimalAction:              ProtoToComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(p.GetMinimalAction()),
		MaxSurge:                   ProtoToComputeInstanceGroupManagerUpdatePolicyMaxSurge(p.GetMaxSurge()),
		MaxUnavailable:             ProtoToComputeInstanceGroupManagerUpdatePolicyMaxUnavailable(p.GetMaxUnavailable()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxSurge converts a InstanceGroupManagerUpdatePolicyMaxSurge resource from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyMaxSurge(p *computepb.ComputeInstanceGroupManagerUpdatePolicyMaxSurge) *compute.InstanceGroupManagerUpdatePolicyMaxSurge {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerUpdatePolicyMaxSurge{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxUnavailable converts a InstanceGroupManagerUpdatePolicyMaxUnavailable resource from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyMaxUnavailable(p *computepb.ComputeInstanceGroupManagerUpdatePolicyMaxUnavailable) *compute.InstanceGroupManagerUpdatePolicyMaxUnavailable {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerUpdatePolicyMaxUnavailable{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToInstanceGroupManager converts a InstanceGroupManager resource from its proto representation.
func ProtoToInstanceGroupManager(p *computepb.ComputeInstanceGroupManager) *compute.InstanceGroupManager {
	obj := &compute.InstanceGroupManager{
		BaseInstanceName:   dcl.StringOrNil(p.BaseInstanceName),
		CreationTimestamp:  dcl.StringOrNil(p.GetCreationTimestamp()),
		DistributionPolicy: ProtoToComputeInstanceGroupManagerDistributionPolicy(p.GetDistributionPolicy()),
		CurrentActions:     ProtoToComputeInstanceGroupManagerCurrentActions(p.GetCurrentActions()),
		Description:        dcl.StringOrNil(p.Description),
		Id:                 dcl.Int64OrNil(p.Id),
		InstanceGroup:      dcl.StringOrNil(p.InstanceGroup),
		InstanceTemplate:   dcl.StringOrNil(p.InstanceTemplate),
		Name:               dcl.StringOrNil(p.Name),
		Status:             ProtoToComputeInstanceGroupManagerStatus(p.GetStatus()),
		UpdatePolicy:       ProtoToComputeInstanceGroupManagerUpdatePolicy(p.GetUpdatePolicy()),
		TargetSize:         dcl.Int64OrNil(p.TargetSize),
		Zone:               dcl.StringOrNil(p.Zone),
		Region:             dcl.StringOrNil(p.Region),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetVersions() {
		obj.Versions = append(obj.Versions, *ProtoToComputeInstanceGroupManagerVersions(r))
	}
	for _, r := range p.GetNamedPorts() {
		obj.NamedPorts = append(obj.NamedPorts, *ProtoToComputeInstanceGroupManagerNamedPorts(r))
	}
	for _, r := range p.GetTargetPools() {
		obj.TargetPools = append(obj.TargetPools, r)
	}
	for _, r := range p.GetAutoHealingPolicies() {
		obj.AutoHealingPolicies = append(obj.AutoHealingPolicies, *ProtoToComputeInstanceGroupManagerAutoHealingPolicies(r))
	}
	return obj
}

// InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(e *compute.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_value["InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyMinimalActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(e *compute.InstanceGroupManagerUpdatePolicyMinimalActionEnum) computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum_value["InstanceGroupManagerUpdatePolicyMinimalActionEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
}

// InstanceGroupManagerDistributionPolicyToProto converts a InstanceGroupManagerDistributionPolicy resource to its proto representation.
func ComputeInstanceGroupManagerDistributionPolicyToProto(o *compute.InstanceGroupManagerDistributionPolicy) *computepb.ComputeInstanceGroupManagerDistributionPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerDistributionPolicy{}
	for _, r := range o.Zones {
		p.Zones = append(p.Zones, ComputeInstanceGroupManagerDistributionPolicyZonesToProto(&r))
	}
	return p
}

// InstanceGroupManagerDistributionPolicyZonesToProto converts a InstanceGroupManagerDistributionPolicyZones resource to its proto representation.
func ComputeInstanceGroupManagerDistributionPolicyZonesToProto(o *compute.InstanceGroupManagerDistributionPolicyZones) *computepb.ComputeInstanceGroupManagerDistributionPolicyZones {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerDistributionPolicyZones{
		Zone: dcl.ValueOrEmptyString(o.Zone),
	}
	return p
}

// InstanceGroupManagerCurrentActionsToProto converts a InstanceGroupManagerCurrentActions resource to its proto representation.
func ComputeInstanceGroupManagerCurrentActionsToProto(o *compute.InstanceGroupManagerCurrentActions) *computepb.ComputeInstanceGroupManagerCurrentActions {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerCurrentActions{
		Abandoning:             dcl.ValueOrEmptyInt64(o.Abandoning),
		Creating:               dcl.ValueOrEmptyInt64(o.Creating),
		CreatingWithoutRetries: dcl.ValueOrEmptyInt64(o.CreatingWithoutRetries),
		Deleting:               dcl.ValueOrEmptyInt64(o.Deleting),
		None:                   dcl.ValueOrEmptyInt64(o.None),
		Recreating:             dcl.ValueOrEmptyInt64(o.Recreating),
		Refreshing:             dcl.ValueOrEmptyInt64(o.Refreshing),
		Restarting:             dcl.ValueOrEmptyInt64(o.Restarting),
	}
	return p
}

// InstanceGroupManagerVersionsToProto converts a InstanceGroupManagerVersions resource to its proto representation.
func ComputeInstanceGroupManagerVersionsToProto(o *compute.InstanceGroupManagerVersions) *computepb.ComputeInstanceGroupManagerVersions {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerVersions{
		Name:             dcl.ValueOrEmptyString(o.Name),
		InstanceTemplate: dcl.ValueOrEmptyString(o.InstanceTemplate),
		TargetSize:       ComputeInstanceGroupManagerVersionsTargetSizeToProto(o.TargetSize),
	}
	return p
}

// InstanceGroupManagerVersionsTargetSizeToProto converts a InstanceGroupManagerVersionsTargetSize resource to its proto representation.
func ComputeInstanceGroupManagerVersionsTargetSizeToProto(o *compute.InstanceGroupManagerVersionsTargetSize) *computepb.ComputeInstanceGroupManagerVersionsTargetSize {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerVersionsTargetSize{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// InstanceGroupManagerNamedPortsToProto converts a InstanceGroupManagerNamedPorts resource to its proto representation.
func ComputeInstanceGroupManagerNamedPortsToProto(o *compute.InstanceGroupManagerNamedPorts) *computepb.ComputeInstanceGroupManagerNamedPorts {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerNamedPorts{
		Name: dcl.ValueOrEmptyString(o.Name),
		Port: dcl.ValueOrEmptyInt64(o.Port),
	}
	return p
}

// InstanceGroupManagerStatusToProto converts a InstanceGroupManagerStatus resource to its proto representation.
func ComputeInstanceGroupManagerStatusToProto(o *compute.InstanceGroupManagerStatus) *computepb.ComputeInstanceGroupManagerStatus {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatus{
		IsStable:      dcl.ValueOrEmptyBool(o.IsStable),
		VersionTarget: ComputeInstanceGroupManagerStatusVersionTargetToProto(o.VersionTarget),
		Autoscalar:    dcl.ValueOrEmptyString(o.Autoscalar),
	}
	return p
}

// InstanceGroupManagerStatusVersionTargetToProto converts a InstanceGroupManagerStatusVersionTarget resource to its proto representation.
func ComputeInstanceGroupManagerStatusVersionTargetToProto(o *compute.InstanceGroupManagerStatusVersionTarget) *computepb.ComputeInstanceGroupManagerStatusVersionTarget {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatusVersionTarget{
		IsReached: dcl.ValueOrEmptyBool(o.IsReached),
	}
	return p
}

// InstanceGroupManagerAutoHealingPoliciesToProto converts a InstanceGroupManagerAutoHealingPolicies resource to its proto representation.
func ComputeInstanceGroupManagerAutoHealingPoliciesToProto(o *compute.InstanceGroupManagerAutoHealingPolicies) *computepb.ComputeInstanceGroupManagerAutoHealingPolicies {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerAutoHealingPolicies{
		HealthCheck:     dcl.ValueOrEmptyString(o.HealthCheck),
		InitialDelaySec: dcl.ValueOrEmptyInt64(o.InitialDelaySec),
	}
	return p
}

// InstanceGroupManagerUpdatePolicyToProto converts a InstanceGroupManagerUpdatePolicy resource to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyToProto(o *compute.InstanceGroupManagerUpdatePolicy) *computepb.ComputeInstanceGroupManagerUpdatePolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerUpdatePolicy{
		InstanceRedistributionType: ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(o.InstanceRedistributionType),
		MinimalAction:              ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(o.MinimalAction),
		MaxSurge:                   ComputeInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o.MaxSurge),
		MaxUnavailable:             ComputeInstanceGroupManagerUpdatePolicyMaxUnavailableToProto(o.MaxUnavailable),
	}
	return p
}

// InstanceGroupManagerUpdatePolicyMaxSurgeToProto converts a InstanceGroupManagerUpdatePolicyMaxSurge resource to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o *compute.InstanceGroupManagerUpdatePolicyMaxSurge) *computepb.ComputeInstanceGroupManagerUpdatePolicyMaxSurge {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerUpdatePolicyMaxSurge{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// InstanceGroupManagerUpdatePolicyMaxUnavailableToProto converts a InstanceGroupManagerUpdatePolicyMaxUnavailable resource to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyMaxUnavailableToProto(o *compute.InstanceGroupManagerUpdatePolicyMaxUnavailable) *computepb.ComputeInstanceGroupManagerUpdatePolicyMaxUnavailable {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerUpdatePolicyMaxUnavailable{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// InstanceGroupManagerToProto converts a InstanceGroupManager resource to its proto representation.
func InstanceGroupManagerToProto(resource *compute.InstanceGroupManager) *computepb.ComputeInstanceGroupManager {
	p := &computepb.ComputeInstanceGroupManager{
		BaseInstanceName:   dcl.ValueOrEmptyString(resource.BaseInstanceName),
		CreationTimestamp:  dcl.ValueOrEmptyString(resource.CreationTimestamp),
		DistributionPolicy: ComputeInstanceGroupManagerDistributionPolicyToProto(resource.DistributionPolicy),
		CurrentActions:     ComputeInstanceGroupManagerCurrentActionsToProto(resource.CurrentActions),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		Id:                 dcl.ValueOrEmptyInt64(resource.Id),
		InstanceGroup:      dcl.ValueOrEmptyString(resource.InstanceGroup),
		InstanceTemplate:   dcl.ValueOrEmptyString(resource.InstanceTemplate),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Status:             ComputeInstanceGroupManagerStatusToProto(resource.Status),
		UpdatePolicy:       ComputeInstanceGroupManagerUpdatePolicyToProto(resource.UpdatePolicy),
		TargetSize:         dcl.ValueOrEmptyInt64(resource.TargetSize),
		Zone:               dcl.ValueOrEmptyString(resource.Zone),
		Region:             dcl.ValueOrEmptyString(resource.Region),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Versions {
		p.Versions = append(p.Versions, ComputeInstanceGroupManagerVersionsToProto(&r))
	}
	for _, r := range resource.NamedPorts {
		p.NamedPorts = append(p.NamedPorts, ComputeInstanceGroupManagerNamedPortsToProto(&r))
	}
	for _, r := range resource.TargetPools {
		p.TargetPools = append(p.TargetPools, r)
	}
	for _, r := range resource.AutoHealingPolicies {
		p.AutoHealingPolicies = append(p.AutoHealingPolicies, ComputeInstanceGroupManagerAutoHealingPoliciesToProto(&r))
	}

	return p
}

// ApplyInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) applyInstanceGroupManager(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeInstanceGroupManagerRequest) (*computepb.ComputeInstanceGroupManager, error) {
	p := ProtoToInstanceGroupManager(request.GetResource())
	res, err := c.ApplyInstanceGroupManager(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceGroupManagerToProto(res)
	return r, nil
}

// ApplyInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) ApplyComputeInstanceGroupManager(ctx context.Context, request *computepb.ApplyComputeInstanceGroupManagerRequest) (*computepb.ComputeInstanceGroupManager, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstanceGroupManager(ctx, cl, request)
}

// DeleteInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Delete() method.
func (s *InstanceGroupManagerServer) DeleteComputeInstanceGroupManager(ctx context.Context, request *computepb.DeleteComputeInstanceGroupManagerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstanceGroupManager(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstanceGroupManager(ctx, ProtoToInstanceGroupManager(request.GetResource()))

}

// ListComputeInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManagerList() method.
func (s *InstanceGroupManagerServer) ListComputeInstanceGroupManager(ctx context.Context, request *computepb.ListComputeInstanceGroupManagerRequest) (*computepb.ListComputeInstanceGroupManagerResponse, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstanceGroupManager(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeInstanceGroupManager
	for _, r := range resources.Items {
		rp := InstanceGroupManagerToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeInstanceGroupManagerResponse{Items: protos}, nil
}

func createConfigInstanceGroupManager(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
