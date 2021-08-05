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

// ProtoToInstanceGroupManagerDistributionPolicyTargetShapeEnum converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(e computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum) *compute.InstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerDistributionPolicyTargetShapeEnum(n[len("ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyTypeEnum converts a InstanceGroupManagerUpdatePolicyTypeEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyTypeEnum(e computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum) *compute.InstanceGroupManagerUpdatePolicyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerUpdatePolicyTypeEnum(n[len("ComputeInstanceGroupManagerUpdatePolicyTypeEnum"):])
		return &e
	}
	return nil
}

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

// ProtoToInstanceGroupManagerUpdatePolicyReplacementMethodEnum converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(e computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum) *compute.InstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerUpdatePolicyReplacementMethodEnum(n[len("ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(e computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) *compute.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(n[len("ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerDistributionPolicy converts a InstanceGroupManagerDistributionPolicy resource from its proto representation.
func ProtoToComputeInstanceGroupManagerDistributionPolicy(p *computepb.ComputeInstanceGroupManagerDistributionPolicy) *compute.InstanceGroupManagerDistributionPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerDistributionPolicy{
		TargetShape: ProtoToComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(p.GetTargetShape()),
	}
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

// ProtoToInstanceGroupManagerVersions converts a InstanceGroupManagerVersions resource from its proto representation.
func ProtoToComputeInstanceGroupManagerVersions(p *computepb.ComputeInstanceGroupManagerVersions) *compute.InstanceGroupManagerVersions {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerVersions{
		Name:             dcl.StringOrNil(p.Name),
		InstanceTemplate: dcl.StringOrNil(p.InstanceTemplate),
		TargetSize:       ProtoToComputeInstanceGroupManagerFixedOrPercent(p.GetTargetSize()),
	}
	return obj
}

// ProtoToInstanceGroupManagerFixedOrPercent converts a InstanceGroupManagerFixedOrPercent resource from its proto representation.
func ProtoToComputeInstanceGroupManagerFixedOrPercent(p *computepb.ComputeInstanceGroupManagerFixedOrPercent) *compute.InstanceGroupManagerFixedOrPercent {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerFixedOrPercent{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToInstanceGroupManagerCurrentActions converts a InstanceGroupManagerCurrentActions resource from its proto representation.
func ProtoToComputeInstanceGroupManagerCurrentActions(p *computepb.ComputeInstanceGroupManagerCurrentActions) *compute.InstanceGroupManagerCurrentActions {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerCurrentActions{
		None:                   dcl.Int64OrNil(p.None),
		Creating:               dcl.Int64OrNil(p.Creating),
		CreatingWithoutRetries: dcl.Int64OrNil(p.CreatingWithoutRetries),
		Verifying:              dcl.Int64OrNil(p.Verifying),
		Recreating:             dcl.Int64OrNil(p.Recreating),
		Deleting:               dcl.Int64OrNil(p.Deleting),
		Abandoning:             dcl.Int64OrNil(p.Abandoning),
		Restarting:             dcl.Int64OrNil(p.Restarting),
		Refreshing:             dcl.Int64OrNil(p.Refreshing),
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
		Stateful:      ProtoToComputeInstanceGroupManagerStatusStateful(p.GetStateful()),
		Autoscaler:    dcl.StringOrNil(p.Autoscaler),
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

// ProtoToInstanceGroupManagerStatusStateful converts a InstanceGroupManagerStatusStateful resource from its proto representation.
func ProtoToComputeInstanceGroupManagerStatusStateful(p *computepb.ComputeInstanceGroupManagerStatusStateful) *compute.InstanceGroupManagerStatusStateful {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatusStateful{
		HasStatefulConfig:  dcl.Bool(p.HasStatefulConfig),
		PerInstanceConfigs: ProtoToComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs(p.GetPerInstanceConfigs()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStatefulPerInstanceConfigs converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs resource from its proto representation.
func ProtoToComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs(p *computepb.ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs) *compute.InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatusStatefulPerInstanceConfigs{
		AllEffective: dcl.Bool(p.AllEffective),
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
		Type:                       ProtoToComputeInstanceGroupManagerUpdatePolicyTypeEnum(p.GetType()),
		InstanceRedistributionType: ProtoToComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(p.GetInstanceRedistributionType()),
		MinimalAction:              ProtoToComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(p.GetMinimalAction()),
		MaxSurge:                   ProtoToComputeInstanceGroupManagerFixedOrPercent(p.GetMaxSurge()),
		MaxUnavailable:             ProtoToComputeInstanceGroupManagerFixedOrPercent(p.GetMaxUnavailable()),
		ReplacementMethod:          ProtoToComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(p.GetReplacementMethod()),
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

// ProtoToInstanceGroupManagerStatefulPolicy converts a InstanceGroupManagerStatefulPolicy resource from its proto representation.
func ProtoToComputeInstanceGroupManagerStatefulPolicy(p *computepb.ComputeInstanceGroupManagerStatefulPolicy) *compute.InstanceGroupManagerStatefulPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatefulPolicy{
		PreservedState: ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedState(p.GetPreservedState()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedState converts a InstanceGroupManagerStatefulPolicyPreservedState resource from its proto representation.
func ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedState(p *computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedState) *compute.InstanceGroupManagerStatefulPolicyPreservedState {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatefulPolicyPreservedState{}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisks converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks resource from its proto representation.
func ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks(p *computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks) *compute.InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatefulPolicyPreservedStateDisks{
		AutoDelete: ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManager converts a InstanceGroupManager resource from its proto representation.
func ProtoToInstanceGroupManager(p *computepb.ComputeInstanceGroupManager) *compute.InstanceGroupManager {
	obj := &compute.InstanceGroupManager{
		Id:                 dcl.Int64OrNil(p.Id),
		CreationTimestamp:  dcl.StringOrNil(p.CreationTimestamp),
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		Zone:               dcl.StringOrNil(p.Zone),
		Region:             dcl.StringOrNil(p.Region),
		DistributionPolicy: ProtoToComputeInstanceGroupManagerDistributionPolicy(p.GetDistributionPolicy()),
		InstanceTemplate:   dcl.StringOrNil(p.InstanceTemplate),
		InstanceGroup:      dcl.StringOrNil(p.InstanceGroup),
		BaseInstanceName:   dcl.StringOrNil(p.BaseInstanceName),
		Fingerprint:        dcl.StringOrNil(p.Fingerprint),
		CurrentActions:     ProtoToComputeInstanceGroupManagerCurrentActions(p.GetCurrentActions()),
		Status:             ProtoToComputeInstanceGroupManagerStatus(p.GetStatus()),
		TargetSize:         dcl.Int64OrNil(p.TargetSize),
		SelfLink:           dcl.StringOrNil(p.SelfLink),
		UpdatePolicy:       ProtoToComputeInstanceGroupManagerUpdatePolicy(p.GetUpdatePolicy()),
		StatefulPolicy:     ProtoToComputeInstanceGroupManagerStatefulPolicy(p.GetStatefulPolicy()),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetVersions() {
		obj.Versions = append(obj.Versions, *ProtoToComputeInstanceGroupManagerVersions(r))
	}
	for _, r := range p.GetTargetPools() {
		obj.TargetPools = append(obj.TargetPools, r)
	}
	for _, r := range p.GetAutoHealingPolicies() {
		obj.AutoHealingPolicies = append(obj.AutoHealingPolicies, *ProtoToComputeInstanceGroupManagerAutoHealingPolicies(r))
	}
	for _, r := range p.GetNamedPorts() {
		obj.NamedPorts = append(obj.NamedPorts, *ProtoToComputeInstanceGroupManagerNamedPorts(r))
	}
	return obj
}

// InstanceGroupManagerDistributionPolicyTargetShapeEnumToProto converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum to its proto representation.
func ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(e *compute.InstanceGroupManagerDistributionPolicyTargetShapeEnum) computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum_value["InstanceGroupManagerDistributionPolicyTargetShapeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
}

// InstanceGroupManagerUpdatePolicyTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyTypeEnum enum to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyTypeEnumToProto(e *compute.InstanceGroupManagerUpdatePolicyTypeEnum) computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum_value["InstanceGroupManagerUpdatePolicyTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum(0)
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

// InstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(e *compute.InstanceGroupManagerUpdatePolicyReplacementMethodEnum) computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum_value["InstanceGroupManagerUpdatePolicyReplacementMethodEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum to its proto representation.
func ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(e *compute.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
}

// InstanceGroupManagerDistributionPolicyToProto converts a InstanceGroupManagerDistributionPolicy resource to its proto representation.
func ComputeInstanceGroupManagerDistributionPolicyToProto(o *compute.InstanceGroupManagerDistributionPolicy) *computepb.ComputeInstanceGroupManagerDistributionPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerDistributionPolicy{
		TargetShape: ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(o.TargetShape),
	}
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

// InstanceGroupManagerVersionsToProto converts a InstanceGroupManagerVersions resource to its proto representation.
func ComputeInstanceGroupManagerVersionsToProto(o *compute.InstanceGroupManagerVersions) *computepb.ComputeInstanceGroupManagerVersions {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerVersions{
		Name:             dcl.ValueOrEmptyString(o.Name),
		InstanceTemplate: dcl.ValueOrEmptyString(o.InstanceTemplate),
		TargetSize:       ComputeInstanceGroupManagerFixedOrPercentToProto(o.TargetSize),
	}
	return p
}

// InstanceGroupManagerFixedOrPercentToProto converts a InstanceGroupManagerFixedOrPercent resource to its proto representation.
func ComputeInstanceGroupManagerFixedOrPercentToProto(o *compute.InstanceGroupManagerFixedOrPercent) *computepb.ComputeInstanceGroupManagerFixedOrPercent {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerFixedOrPercent{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// InstanceGroupManagerCurrentActionsToProto converts a InstanceGroupManagerCurrentActions resource to its proto representation.
func ComputeInstanceGroupManagerCurrentActionsToProto(o *compute.InstanceGroupManagerCurrentActions) *computepb.ComputeInstanceGroupManagerCurrentActions {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerCurrentActions{
		None:                   dcl.ValueOrEmptyInt64(o.None),
		Creating:               dcl.ValueOrEmptyInt64(o.Creating),
		CreatingWithoutRetries: dcl.ValueOrEmptyInt64(o.CreatingWithoutRetries),
		Verifying:              dcl.ValueOrEmptyInt64(o.Verifying),
		Recreating:             dcl.ValueOrEmptyInt64(o.Recreating),
		Deleting:               dcl.ValueOrEmptyInt64(o.Deleting),
		Abandoning:             dcl.ValueOrEmptyInt64(o.Abandoning),
		Restarting:             dcl.ValueOrEmptyInt64(o.Restarting),
		Refreshing:             dcl.ValueOrEmptyInt64(o.Refreshing),
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
		Stateful:      ComputeInstanceGroupManagerStatusStatefulToProto(o.Stateful),
		Autoscaler:    dcl.ValueOrEmptyString(o.Autoscaler),
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

// InstanceGroupManagerStatusStatefulToProto converts a InstanceGroupManagerStatusStateful resource to its proto representation.
func ComputeInstanceGroupManagerStatusStatefulToProto(o *compute.InstanceGroupManagerStatusStateful) *computepb.ComputeInstanceGroupManagerStatusStateful {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatusStateful{
		HasStatefulConfig:  dcl.ValueOrEmptyBool(o.HasStatefulConfig),
		PerInstanceConfigs: ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o.PerInstanceConfigs),
	}
	return p
}

// InstanceGroupManagerStatusStatefulPerInstanceConfigsToProto converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs resource to its proto representation.
func ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o *compute.InstanceGroupManagerStatusStatefulPerInstanceConfigs) *computepb.ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs{
		AllEffective: dcl.ValueOrEmptyBool(o.AllEffective),
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
		Type:                       ComputeInstanceGroupManagerUpdatePolicyTypeEnumToProto(o.Type),
		InstanceRedistributionType: ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(o.InstanceRedistributionType),
		MinimalAction:              ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(o.MinimalAction),
		MaxSurge:                   ComputeInstanceGroupManagerFixedOrPercentToProto(o.MaxSurge),
		MaxUnavailable:             ComputeInstanceGroupManagerFixedOrPercentToProto(o.MaxUnavailable),
		ReplacementMethod:          ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(o.ReplacementMethod),
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

// InstanceGroupManagerStatefulPolicyToProto converts a InstanceGroupManagerStatefulPolicy resource to its proto representation.
func ComputeInstanceGroupManagerStatefulPolicyToProto(o *compute.InstanceGroupManagerStatefulPolicy) *computepb.ComputeInstanceGroupManagerStatefulPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatefulPolicy{
		PreservedState: ComputeInstanceGroupManagerStatefulPolicyPreservedStateToProto(o.PreservedState),
	}
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateToProto converts a InstanceGroupManagerStatefulPolicyPreservedState resource to its proto representation.
func ComputeInstanceGroupManagerStatefulPolicyPreservedStateToProto(o *compute.InstanceGroupManagerStatefulPolicyPreservedState) *computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedState {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedState{}
	p.Disks = make(map[string]*computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks)
	for k, r := range o.Disks {
		p.Disks[k] = ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(&r)
	}
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks resource to its proto representation.
func ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(o *compute.InstanceGroupManagerStatefulPolicyPreservedStateDisks) *computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks{
		AutoDelete: ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(o.AutoDelete),
	}
	return p
}

// InstanceGroupManagerToProto converts a InstanceGroupManager resource to its proto representation.
func InstanceGroupManagerToProto(resource *compute.InstanceGroupManager) *computepb.ComputeInstanceGroupManager {
	p := &computepb.ComputeInstanceGroupManager{
		Id:                 dcl.ValueOrEmptyInt64(resource.Id),
		CreationTimestamp:  dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		Zone:               dcl.ValueOrEmptyString(resource.Zone),
		Region:             dcl.ValueOrEmptyString(resource.Region),
		DistributionPolicy: ComputeInstanceGroupManagerDistributionPolicyToProto(resource.DistributionPolicy),
		InstanceTemplate:   dcl.ValueOrEmptyString(resource.InstanceTemplate),
		InstanceGroup:      dcl.ValueOrEmptyString(resource.InstanceGroup),
		BaseInstanceName:   dcl.ValueOrEmptyString(resource.BaseInstanceName),
		Fingerprint:        dcl.ValueOrEmptyString(resource.Fingerprint),
		CurrentActions:     ComputeInstanceGroupManagerCurrentActionsToProto(resource.CurrentActions),
		Status:             ComputeInstanceGroupManagerStatusToProto(resource.Status),
		TargetSize:         dcl.ValueOrEmptyInt64(resource.TargetSize),
		SelfLink:           dcl.ValueOrEmptyString(resource.SelfLink),
		UpdatePolicy:       ComputeInstanceGroupManagerUpdatePolicyToProto(resource.UpdatePolicy),
		StatefulPolicy:     ComputeInstanceGroupManagerStatefulPolicyToProto(resource.StatefulPolicy),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Versions {
		p.Versions = append(p.Versions, ComputeInstanceGroupManagerVersionsToProto(&r))
	}
	for _, r := range resource.TargetPools {
		p.TargetPools = append(p.TargetPools, r)
	}
	for _, r := range resource.AutoHealingPolicies {
		p.AutoHealingPolicies = append(p.AutoHealingPolicies, ComputeInstanceGroupManagerAutoHealingPoliciesToProto(&r))
	}
	for _, r := range resource.NamedPorts {
		p.NamedPorts = append(p.NamedPorts, ComputeInstanceGroupManagerNamedPortsToProto(&r))
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

	resources, err := cl.ListInstanceGroupManager(ctx, ProtoToInstanceGroupManager(request.GetResource()))
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
