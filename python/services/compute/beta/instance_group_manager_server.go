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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for InstanceGroupManager.
type InstanceGroupManagerServer struct{}

// ProtoToInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) *beta.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(n[len("InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMinimalActionEnum converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum) *beta.InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyMinimalActionEnum(n[len("InstanceGroupManagerUpdatePolicyMinimalActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerDistributionPolicy converts a InstanceGroupManagerDistributionPolicy resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerDistributionPolicy(p *betapb.ComputeBetaInstanceGroupManagerDistributionPolicy) *beta.InstanceGroupManagerDistributionPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerDistributionPolicy{}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, *ProtoToComputeBetaInstanceGroupManagerDistributionPolicyZones(r))
	}
	return obj
}

// ProtoToInstanceGroupManagerDistributionPolicyZones converts a InstanceGroupManagerDistributionPolicyZones resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerDistributionPolicyZones(p *betapb.ComputeBetaInstanceGroupManagerDistributionPolicyZones) *beta.InstanceGroupManagerDistributionPolicyZones {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerDistributionPolicyZones{
		Zone: dcl.StringOrNil(p.Zone),
	}
	return obj
}

// ProtoToInstanceGroupManagerCurrentActions converts a InstanceGroupManagerCurrentActions resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerCurrentActions(p *betapb.ComputeBetaInstanceGroupManagerCurrentActions) *beta.InstanceGroupManagerCurrentActions {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerCurrentActions{
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
func ProtoToComputeBetaInstanceGroupManagerVersions(p *betapb.ComputeBetaInstanceGroupManagerVersions) *beta.InstanceGroupManagerVersions {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerVersions{
		Name:             dcl.StringOrNil(p.Name),
		InstanceTemplate: dcl.StringOrNil(p.InstanceTemplate),
		TargetSize:       ProtoToComputeBetaInstanceGroupManagerVersionsTargetSize(p.GetTargetSize()),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersionsTargetSize converts a InstanceGroupManagerVersionsTargetSize resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerVersionsTargetSize(p *betapb.ComputeBetaInstanceGroupManagerVersionsTargetSize) *beta.InstanceGroupManagerVersionsTargetSize {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerVersionsTargetSize{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToInstanceGroupManagerNamedPorts converts a InstanceGroupManagerNamedPorts resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerNamedPorts(p *betapb.ComputeBetaInstanceGroupManagerNamedPorts) *beta.InstanceGroupManagerNamedPorts {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerNamedPorts{
		Name: dcl.StringOrNil(p.Name),
		Port: dcl.Int64OrNil(p.Port),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatus converts a InstanceGroupManagerStatus resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatus(p *betapb.ComputeBetaInstanceGroupManagerStatus) *beta.InstanceGroupManagerStatus {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatus{
		IsStable:      dcl.Bool(p.IsStable),
		VersionTarget: ProtoToComputeBetaInstanceGroupManagerStatusVersionTarget(p.GetVersionTarget()),
		Autoscalar:    dcl.StringOrNil(p.Autoscalar),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusVersionTarget converts a InstanceGroupManagerStatusVersionTarget resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatusVersionTarget(p *betapb.ComputeBetaInstanceGroupManagerStatusVersionTarget) *beta.InstanceGroupManagerStatusVersionTarget {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatusVersionTarget{
		IsReached: dcl.Bool(p.IsReached),
	}
	return obj
}

// ProtoToInstanceGroupManagerAutohealingPolicies converts a InstanceGroupManagerAutohealingPolicies resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerAutohealingPolicies(p *betapb.ComputeBetaInstanceGroupManagerAutohealingPolicies) *beta.InstanceGroupManagerAutohealingPolicies {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerAutohealingPolicies{
		HealthCheck:     dcl.StringOrNil(p.HealthCheck),
		InitialDelaySec: dcl.Int64OrNil(p.InitialDelaySec),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicy converts a InstanceGroupManagerUpdatePolicy resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicy(p *betapb.ComputeBetaInstanceGroupManagerUpdatePolicy) *beta.InstanceGroupManagerUpdatePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerUpdatePolicy{
		InstanceRedistributionType: ProtoToComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(p.GetInstanceRedistributionType()),
		MinimalAction:              ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(p.GetMinimalAction()),
		MaxSurge:                   ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge(p.GetMaxSurge()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxSurge converts a InstanceGroupManagerUpdatePolicyMaxSurge resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge(p *betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge) *beta.InstanceGroupManagerUpdatePolicyMaxSurge {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerUpdatePolicyMaxSurge{
		Fixed:          dcl.Int64OrNil(p.Fixed),
		Percent:        dcl.Int64OrNil(p.Percent),
		Calculated:     dcl.Int64OrNil(p.Calculated),
		MaxUnavailable: ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(p.GetMaxUnavailable()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable converts a InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(p *betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) *beta.InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToInstanceGroupManager converts a InstanceGroupManager resource from its proto representation.
func ProtoToInstanceGroupManager(p *betapb.ComputeBetaInstanceGroupManager) *beta.InstanceGroupManager {
	obj := &beta.InstanceGroupManager{
		BaseInstanceName:   dcl.StringOrNil(p.BaseInstanceName),
		CreationTimestamp:  dcl.StringOrNil(p.GetCreationTimestamp()),
		DistributionPolicy: ProtoToComputeBetaInstanceGroupManagerDistributionPolicy(p.GetDistributionPolicy()),
		CurrentActions:     ProtoToComputeBetaInstanceGroupManagerCurrentActions(p.GetCurrentActions()),
		Description:        dcl.StringOrNil(p.Description),
		Id:                 dcl.Int64OrNil(p.Id),
		InstanceGroup:      dcl.StringOrNil(p.InstanceGroup),
		InstanceTemplate:   dcl.StringOrNil(p.InstanceTemplate),
		Name:               dcl.StringOrNil(p.Name),
		Status:             ProtoToComputeBetaInstanceGroupManagerStatus(p.GetStatus()),
		UpdatePolicy:       ProtoToComputeBetaInstanceGroupManagerUpdatePolicy(p.GetUpdatePolicy()),
		TargetSize:         dcl.Int64OrNil(p.TargetSize),
		Zone:               dcl.StringOrNil(p.Zone),
		Region:             dcl.StringOrNil(p.Region),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetVersions() {
		obj.Versions = append(obj.Versions, *ProtoToComputeBetaInstanceGroupManagerVersions(r))
	}
	for _, r := range p.GetNamedPorts() {
		obj.NamedPorts = append(obj.NamedPorts, *ProtoToComputeBetaInstanceGroupManagerNamedPorts(r))
	}
	for _, r := range p.GetTargetPools() {
		obj.TargetPools = append(obj.TargetPools, r)
	}
	for _, r := range p.GetAutohealingPolicies() {
		obj.AutohealingPolicies = append(obj.AutohealingPolicies, *ProtoToComputeBetaInstanceGroupManagerAutohealingPolicies(r))
	}
	return obj
}

// InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_value["InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyMinimalActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyMinimalActionEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum_value["InstanceGroupManagerUpdatePolicyMinimalActionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
}

// InstanceGroupManagerDistributionPolicyToProto converts a InstanceGroupManagerDistributionPolicy resource to its proto representation.
func ComputeBetaInstanceGroupManagerDistributionPolicyToProto(o *beta.InstanceGroupManagerDistributionPolicy) *betapb.ComputeBetaInstanceGroupManagerDistributionPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerDistributionPolicy{}
	for _, r := range o.Zones {
		p.Zones = append(p.Zones, ComputeBetaInstanceGroupManagerDistributionPolicyZonesToProto(&r))
	}
	return p
}

// InstanceGroupManagerDistributionPolicyZonesToProto converts a InstanceGroupManagerDistributionPolicyZones resource to its proto representation.
func ComputeBetaInstanceGroupManagerDistributionPolicyZonesToProto(o *beta.InstanceGroupManagerDistributionPolicyZones) *betapb.ComputeBetaInstanceGroupManagerDistributionPolicyZones {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerDistributionPolicyZones{
		Zone: dcl.ValueOrEmptyString(o.Zone),
	}
	return p
}

// InstanceGroupManagerCurrentActionsToProto converts a InstanceGroupManagerCurrentActions resource to its proto representation.
func ComputeBetaInstanceGroupManagerCurrentActionsToProto(o *beta.InstanceGroupManagerCurrentActions) *betapb.ComputeBetaInstanceGroupManagerCurrentActions {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerCurrentActions{
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
func ComputeBetaInstanceGroupManagerVersionsToProto(o *beta.InstanceGroupManagerVersions) *betapb.ComputeBetaInstanceGroupManagerVersions {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerVersions{
		Name:             dcl.ValueOrEmptyString(o.Name),
		InstanceTemplate: dcl.ValueOrEmptyString(o.InstanceTemplate),
		TargetSize:       ComputeBetaInstanceGroupManagerVersionsTargetSizeToProto(o.TargetSize),
	}
	return p
}

// InstanceGroupManagerVersionsTargetSizeToProto converts a InstanceGroupManagerVersionsTargetSize resource to its proto representation.
func ComputeBetaInstanceGroupManagerVersionsTargetSizeToProto(o *beta.InstanceGroupManagerVersionsTargetSize) *betapb.ComputeBetaInstanceGroupManagerVersionsTargetSize {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerVersionsTargetSize{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// InstanceGroupManagerNamedPortsToProto converts a InstanceGroupManagerNamedPorts resource to its proto representation.
func ComputeBetaInstanceGroupManagerNamedPortsToProto(o *beta.InstanceGroupManagerNamedPorts) *betapb.ComputeBetaInstanceGroupManagerNamedPorts {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerNamedPorts{
		Name: dcl.ValueOrEmptyString(o.Name),
		Port: dcl.ValueOrEmptyInt64(o.Port),
	}
	return p
}

// InstanceGroupManagerStatusToProto converts a InstanceGroupManagerStatus resource to its proto representation.
func ComputeBetaInstanceGroupManagerStatusToProto(o *beta.InstanceGroupManagerStatus) *betapb.ComputeBetaInstanceGroupManagerStatus {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatus{
		IsStable:      dcl.ValueOrEmptyBool(o.IsStable),
		VersionTarget: ComputeBetaInstanceGroupManagerStatusVersionTargetToProto(o.VersionTarget),
		Autoscalar:    dcl.ValueOrEmptyString(o.Autoscalar),
	}
	return p
}

// InstanceGroupManagerStatusVersionTargetToProto converts a InstanceGroupManagerStatusVersionTarget resource to its proto representation.
func ComputeBetaInstanceGroupManagerStatusVersionTargetToProto(o *beta.InstanceGroupManagerStatusVersionTarget) *betapb.ComputeBetaInstanceGroupManagerStatusVersionTarget {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatusVersionTarget{
		IsReached: dcl.ValueOrEmptyBool(o.IsReached),
	}
	return p
}

// InstanceGroupManagerAutohealingPoliciesToProto converts a InstanceGroupManagerAutohealingPolicies resource to its proto representation.
func ComputeBetaInstanceGroupManagerAutohealingPoliciesToProto(o *beta.InstanceGroupManagerAutohealingPolicies) *betapb.ComputeBetaInstanceGroupManagerAutohealingPolicies {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerAutohealingPolicies{
		HealthCheck:     dcl.ValueOrEmptyString(o.HealthCheck),
		InitialDelaySec: dcl.ValueOrEmptyInt64(o.InitialDelaySec),
	}
	return p
}

// InstanceGroupManagerUpdatePolicyToProto converts a InstanceGroupManagerUpdatePolicy resource to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyToProto(o *beta.InstanceGroupManagerUpdatePolicy) *betapb.ComputeBetaInstanceGroupManagerUpdatePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerUpdatePolicy{
		InstanceRedistributionType: ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(o.InstanceRedistributionType),
		MinimalAction:              ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(o.MinimalAction),
		MaxSurge:                   ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o.MaxSurge),
	}
	return p
}

// InstanceGroupManagerUpdatePolicyMaxSurgeToProto converts a InstanceGroupManagerUpdatePolicyMaxSurge resource to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o *beta.InstanceGroupManagerUpdatePolicyMaxSurge) *betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge{
		Fixed:          dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:        dcl.ValueOrEmptyInt64(o.Percent),
		Calculated:     dcl.ValueOrEmptyInt64(o.Calculated),
		MaxUnavailable: ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableToProto(o.MaxUnavailable),
	}
	return p
}

// InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableToProto converts a InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable resource to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableToProto(o *beta.InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) *betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// InstanceGroupManagerToProto converts a InstanceGroupManager resource to its proto representation.
func InstanceGroupManagerToProto(resource *beta.InstanceGroupManager) *betapb.ComputeBetaInstanceGroupManager {
	p := &betapb.ComputeBetaInstanceGroupManager{
		BaseInstanceName:   dcl.ValueOrEmptyString(resource.BaseInstanceName),
		CreationTimestamp:  dcl.ValueOrEmptyString(resource.CreationTimestamp),
		DistributionPolicy: ComputeBetaInstanceGroupManagerDistributionPolicyToProto(resource.DistributionPolicy),
		CurrentActions:     ComputeBetaInstanceGroupManagerCurrentActionsToProto(resource.CurrentActions),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		Id:                 dcl.ValueOrEmptyInt64(resource.Id),
		InstanceGroup:      dcl.ValueOrEmptyString(resource.InstanceGroup),
		InstanceTemplate:   dcl.ValueOrEmptyString(resource.InstanceTemplate),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Status:             ComputeBetaInstanceGroupManagerStatusToProto(resource.Status),
		UpdatePolicy:       ComputeBetaInstanceGroupManagerUpdatePolicyToProto(resource.UpdatePolicy),
		TargetSize:         dcl.ValueOrEmptyInt64(resource.TargetSize),
		Zone:               dcl.ValueOrEmptyString(resource.Zone),
		Region:             dcl.ValueOrEmptyString(resource.Region),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Versions {
		p.Versions = append(p.Versions, ComputeBetaInstanceGroupManagerVersionsToProto(&r))
	}
	for _, r := range resource.NamedPorts {
		p.NamedPorts = append(p.NamedPorts, ComputeBetaInstanceGroupManagerNamedPortsToProto(&r))
	}
	for _, r := range resource.TargetPools {
		p.TargetPools = append(p.TargetPools, r)
	}
	for _, r := range resource.AutohealingPolicies {
		p.AutohealingPolicies = append(p.AutohealingPolicies, ComputeBetaInstanceGroupManagerAutohealingPoliciesToProto(&r))
	}

	return p
}

// ApplyInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) applyInstanceGroupManager(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaInstanceGroupManagerRequest) (*betapb.ComputeBetaInstanceGroupManager, error) {
	p := ProtoToInstanceGroupManager(request.GetResource())
	res, err := c.ApplyInstanceGroupManager(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceGroupManagerToProto(res)
	return r, nil
}

// ApplyInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) ApplyComputeBetaInstanceGroupManager(ctx context.Context, request *betapb.ApplyComputeBetaInstanceGroupManagerRequest) (*betapb.ComputeBetaInstanceGroupManager, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstanceGroupManager(ctx, cl, request)
}

// DeleteInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Delete() method.
func (s *InstanceGroupManagerServer) DeleteComputeBetaInstanceGroupManager(ctx context.Context, request *betapb.DeleteComputeBetaInstanceGroupManagerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstanceGroupManager(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstanceGroupManager(ctx, ProtoToInstanceGroupManager(request.GetResource()))

}

// ListInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManagerList() method.
func (s *InstanceGroupManagerServer) ListComputeBetaInstanceGroupManager(ctx context.Context, request *betapb.ListComputeBetaInstanceGroupManagerRequest) (*betapb.ListComputeBetaInstanceGroupManagerResponse, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstanceGroupManager(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaInstanceGroupManager
	for _, r := range resources.Items {
		rp := InstanceGroupManagerToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaInstanceGroupManagerResponse{Items: protos}, nil
}

func createConfigInstanceGroupManager(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
