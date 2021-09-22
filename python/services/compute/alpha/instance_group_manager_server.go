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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/alpha/compute_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
)

// Server implements the gRPC interface for InstanceGroupManager.
type InstanceGroupManagerServer struct{}

// ProtoToInstanceGroupManagerDistributionPolicyTargetShapeEnum converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(e alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum) *alpha.InstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerDistributionPolicyTargetShapeEnum(n[len("ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyTypeEnum converts a InstanceGroupManagerUpdatePolicyTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum) *alpha.InstanceGroupManagerUpdatePolicyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyTypeEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) *alpha.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMinimalActionEnum converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum) *alpha.InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyMinimalActionEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyReplacementMethodEnum converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum) *alpha.InstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyReplacementMethodEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum converts a InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum) *alpha.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(e alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) *alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(n[len("ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerFailoverActionEnum converts a InstanceGroupManagerFailoverActionEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerFailoverActionEnum(e alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum) *alpha.InstanceGroupManagerFailoverActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerFailoverActionEnum(n[len("ComputeAlphaInstanceGroupManagerFailoverActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerDistributionPolicy converts a InstanceGroupManagerDistributionPolicy resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerDistributionPolicy(p *alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicy) *alpha.InstanceGroupManagerDistributionPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerDistributionPolicy{
		TargetShape: ProtoToComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(p.GetTargetShape()),
	}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, *ProtoToComputeAlphaInstanceGroupManagerDistributionPolicyZones(r))
	}
	return obj
}

// ProtoToInstanceGroupManagerDistributionPolicyZones converts a InstanceGroupManagerDistributionPolicyZones resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerDistributionPolicyZones(p *alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyZones) *alpha.InstanceGroupManagerDistributionPolicyZones {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerDistributionPolicyZones{
		Zone: dcl.StringOrNil(p.Zone),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersions converts a InstanceGroupManagerVersions resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerVersions(p *alphapb.ComputeAlphaInstanceGroupManagerVersions) *alpha.InstanceGroupManagerVersions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerVersions{
		Name:             dcl.StringOrNil(p.Name),
		InstanceTemplate: dcl.StringOrNil(p.InstanceTemplate),
		TargetSize:       ProtoToComputeAlphaInstanceGroupManagerFixedOrPercent(p.GetTargetSize()),
	}
	return obj
}

// ProtoToInstanceGroupManagerFixedOrPercent converts a InstanceGroupManagerFixedOrPercent resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerFixedOrPercent(p *alphapb.ComputeAlphaInstanceGroupManagerFixedOrPercent) *alpha.InstanceGroupManagerFixedOrPercent {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerFixedOrPercent{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToInstanceGroupManagerCurrentActions converts a InstanceGroupManagerCurrentActions resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerCurrentActions(p *alphapb.ComputeAlphaInstanceGroupManagerCurrentActions) *alpha.InstanceGroupManagerCurrentActions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerCurrentActions{
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
func ProtoToComputeAlphaInstanceGroupManagerStatus(p *alphapb.ComputeAlphaInstanceGroupManagerStatus) *alpha.InstanceGroupManagerStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatus{
		IsStable:      dcl.Bool(p.IsStable),
		VersionTarget: ProtoToComputeAlphaInstanceGroupManagerStatusVersionTarget(p.GetVersionTarget()),
		Stateful:      ProtoToComputeAlphaInstanceGroupManagerStatusStateful(p.GetStateful()),
		Autoscaler:    dcl.StringOrNil(p.Autoscaler),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusVersionTarget converts a InstanceGroupManagerStatusVersionTarget resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatusVersionTarget(p *alphapb.ComputeAlphaInstanceGroupManagerStatusVersionTarget) *alpha.InstanceGroupManagerStatusVersionTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatusVersionTarget{
		IsReached: dcl.Bool(p.IsReached),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStateful converts a InstanceGroupManagerStatusStateful resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatusStateful(p *alphapb.ComputeAlphaInstanceGroupManagerStatusStateful) *alpha.InstanceGroupManagerStatusStateful {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatusStateful{
		HasStatefulConfig:  dcl.Bool(p.HasStatefulConfig),
		PerInstanceConfigs: ProtoToComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs(p.GetPerInstanceConfigs()),
		IsStateful:         dcl.Bool(p.IsStateful),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStatefulPerInstanceConfigs converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs(p *alphapb.ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs) *alpha.InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatusStatefulPerInstanceConfigs{
		AllEffective: dcl.Bool(p.AllEffective),
	}
	return obj
}

// ProtoToInstanceGroupManagerAutoHealingPolicies converts a InstanceGroupManagerAutoHealingPolicies resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerAutoHealingPolicies(p *alphapb.ComputeAlphaInstanceGroupManagerAutoHealingPolicies) *alpha.InstanceGroupManagerAutoHealingPolicies {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerAutoHealingPolicies{
		HealthCheck:     dcl.StringOrNil(p.HealthCheck),
		InitialDelaySec: dcl.Int64OrNil(p.InitialDelaySec),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicy converts a InstanceGroupManagerUpdatePolicy resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicy(p *alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicy) *alpha.InstanceGroupManagerUpdatePolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerUpdatePolicy{
		Type:                        ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(p.GetType()),
		InstanceRedistributionType:  ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(p.GetInstanceRedistributionType()),
		MinimalAction:               ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(p.GetMinimalAction()),
		MaxSurge:                    ProtoToComputeAlphaInstanceGroupManagerFixedOrPercent(p.GetMaxSurge()),
		MaxUnavailable:              ProtoToComputeAlphaInstanceGroupManagerFixedOrPercent(p.GetMaxUnavailable()),
		ReplacementMethod:           ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(p.GetReplacementMethod()),
		MostDisruptiveAllowedAction: ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(p.GetMostDisruptiveAllowedAction()),
		MinReadySec:                 dcl.Int64OrNil(p.MinReadySec),
	}
	return obj
}

// ProtoToInstanceGroupManagerNamedPorts converts a InstanceGroupManagerNamedPorts resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerNamedPorts(p *alphapb.ComputeAlphaInstanceGroupManagerNamedPorts) *alpha.InstanceGroupManagerNamedPorts {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerNamedPorts{
		Name: dcl.StringOrNil(p.Name),
		Port: dcl.Int64OrNil(p.Port),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicy converts a InstanceGroupManagerStatefulPolicy resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicy(p *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicy) *alpha.InstanceGroupManagerStatefulPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatefulPolicy{
		PreservedState: ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState(p.GetPreservedState()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedState converts a InstanceGroupManagerStatefulPolicyPreservedState resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState(p *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState) *alpha.InstanceGroupManagerStatefulPolicyPreservedState {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatefulPolicyPreservedState{}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisks converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks resource from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks(p *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks) *alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisks{
		AutoDelete: ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManager converts a InstanceGroupManager resource from its proto representation.
func ProtoToInstanceGroupManager(p *alphapb.ComputeAlphaInstanceGroupManager) *alpha.InstanceGroupManager {
	obj := &alpha.InstanceGroupManager{
		Id:                 dcl.Int64OrNil(p.Id),
		CreationTimestamp:  dcl.StringOrNil(p.CreationTimestamp),
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		Zone:               dcl.StringOrNil(p.Zone),
		Region:             dcl.StringOrNil(p.Region),
		DistributionPolicy: ProtoToComputeAlphaInstanceGroupManagerDistributionPolicy(p.GetDistributionPolicy()),
		InstanceTemplate:   dcl.StringOrNil(p.InstanceTemplate),
		InstanceGroup:      dcl.StringOrNil(p.InstanceGroup),
		BaseInstanceName:   dcl.StringOrNil(p.BaseInstanceName),
		Fingerprint:        dcl.StringOrNil(p.Fingerprint),
		CurrentActions:     ProtoToComputeAlphaInstanceGroupManagerCurrentActions(p.GetCurrentActions()),
		Status:             ProtoToComputeAlphaInstanceGroupManagerStatus(p.GetStatus()),
		TargetSize:         dcl.Int64OrNil(p.TargetSize),
		SelfLink:           dcl.StringOrNil(p.SelfLink),
		UpdatePolicy:       ProtoToComputeAlphaInstanceGroupManagerUpdatePolicy(p.GetUpdatePolicy()),
		StatefulPolicy:     ProtoToComputeAlphaInstanceGroupManagerStatefulPolicy(p.GetStatefulPolicy()),
		ServiceAccount:     dcl.StringOrNil(p.ServiceAccount),
		FailoverAction:     ProtoToComputeAlphaInstanceGroupManagerFailoverActionEnum(p.GetFailoverAction()),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetVersions() {
		obj.Versions = append(obj.Versions, *ProtoToComputeAlphaInstanceGroupManagerVersions(r))
	}
	for _, r := range p.GetTargetPools() {
		obj.TargetPools = append(obj.TargetPools, r)
	}
	for _, r := range p.GetAutoHealingPolicies() {
		obj.AutoHealingPolicies = append(obj.AutoHealingPolicies, *ProtoToComputeAlphaInstanceGroupManagerAutoHealingPolicies(r))
	}
	for _, r := range p.GetNamedPorts() {
		obj.NamedPorts = append(obj.NamedPorts, *ProtoToComputeAlphaInstanceGroupManagerNamedPorts(r))
	}
	return obj
}

// InstanceGroupManagerDistributionPolicyTargetShapeEnumToProto converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(e *alpha.InstanceGroupManagerDistributionPolicyTargetShapeEnum) alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum_value["InstanceGroupManagerDistributionPolicyTargetShapeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
}

// InstanceGroupManagerUpdatePolicyTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyTypeEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyTypeEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum_value["InstanceGroupManagerUpdatePolicyTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_value["InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyMinimalActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyMinimalActionEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum_value["InstanceGroupManagerUpdatePolicyMinimalActionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
}

// InstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyReplacementMethodEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum_value["InstanceGroupManagerUpdatePolicyReplacementMethodEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
}

// InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum_value["InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(e *alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
}

// InstanceGroupManagerFailoverActionEnumToProto converts a InstanceGroupManagerFailoverActionEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerFailoverActionEnumToProto(e *alpha.InstanceGroupManagerFailoverActionEnum) alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum_value["InstanceGroupManagerFailoverActionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum(0)
}

// InstanceGroupManagerDistributionPolicyToProto converts a InstanceGroupManagerDistributionPolicy resource to its proto representation.
func ComputeAlphaInstanceGroupManagerDistributionPolicyToProto(o *alpha.InstanceGroupManagerDistributionPolicy) *alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicy{
		TargetShape: ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(o.TargetShape),
	}
	for _, r := range o.Zones {
		p.Zones = append(p.Zones, ComputeAlphaInstanceGroupManagerDistributionPolicyZonesToProto(&r))
	}
	return p
}

// InstanceGroupManagerDistributionPolicyZonesToProto converts a InstanceGroupManagerDistributionPolicyZones resource to its proto representation.
func ComputeAlphaInstanceGroupManagerDistributionPolicyZonesToProto(o *alpha.InstanceGroupManagerDistributionPolicyZones) *alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyZones {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyZones{
		Zone: dcl.ValueOrEmptyString(o.Zone),
	}
	return p
}

// InstanceGroupManagerVersionsToProto converts a InstanceGroupManagerVersions resource to its proto representation.
func ComputeAlphaInstanceGroupManagerVersionsToProto(o *alpha.InstanceGroupManagerVersions) *alphapb.ComputeAlphaInstanceGroupManagerVersions {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerVersions{
		Name:             dcl.ValueOrEmptyString(o.Name),
		InstanceTemplate: dcl.ValueOrEmptyString(o.InstanceTemplate),
		TargetSize:       ComputeAlphaInstanceGroupManagerFixedOrPercentToProto(o.TargetSize),
	}
	return p
}

// InstanceGroupManagerFixedOrPercentToProto converts a InstanceGroupManagerFixedOrPercent resource to its proto representation.
func ComputeAlphaInstanceGroupManagerFixedOrPercentToProto(o *alpha.InstanceGroupManagerFixedOrPercent) *alphapb.ComputeAlphaInstanceGroupManagerFixedOrPercent {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerFixedOrPercent{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// InstanceGroupManagerCurrentActionsToProto converts a InstanceGroupManagerCurrentActions resource to its proto representation.
func ComputeAlphaInstanceGroupManagerCurrentActionsToProto(o *alpha.InstanceGroupManagerCurrentActions) *alphapb.ComputeAlphaInstanceGroupManagerCurrentActions {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerCurrentActions{
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
func ComputeAlphaInstanceGroupManagerStatusToProto(o *alpha.InstanceGroupManagerStatus) *alphapb.ComputeAlphaInstanceGroupManagerStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatus{
		IsStable:      dcl.ValueOrEmptyBool(o.IsStable),
		VersionTarget: ComputeAlphaInstanceGroupManagerStatusVersionTargetToProto(o.VersionTarget),
		Stateful:      ComputeAlphaInstanceGroupManagerStatusStatefulToProto(o.Stateful),
		Autoscaler:    dcl.ValueOrEmptyString(o.Autoscaler),
	}
	return p
}

// InstanceGroupManagerStatusVersionTargetToProto converts a InstanceGroupManagerStatusVersionTarget resource to its proto representation.
func ComputeAlphaInstanceGroupManagerStatusVersionTargetToProto(o *alpha.InstanceGroupManagerStatusVersionTarget) *alphapb.ComputeAlphaInstanceGroupManagerStatusVersionTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatusVersionTarget{
		IsReached: dcl.ValueOrEmptyBool(o.IsReached),
	}
	return p
}

// InstanceGroupManagerStatusStatefulToProto converts a InstanceGroupManagerStatusStateful resource to its proto representation.
func ComputeAlphaInstanceGroupManagerStatusStatefulToProto(o *alpha.InstanceGroupManagerStatusStateful) *alphapb.ComputeAlphaInstanceGroupManagerStatusStateful {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatusStateful{
		HasStatefulConfig:  dcl.ValueOrEmptyBool(o.HasStatefulConfig),
		PerInstanceConfigs: ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o.PerInstanceConfigs),
		IsStateful:         dcl.ValueOrEmptyBool(o.IsStateful),
	}
	return p
}

// InstanceGroupManagerStatusStatefulPerInstanceConfigsToProto converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs resource to its proto representation.
func ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o *alpha.InstanceGroupManagerStatusStatefulPerInstanceConfigs) *alphapb.ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs{
		AllEffective: dcl.ValueOrEmptyBool(o.AllEffective),
	}
	return p
}

// InstanceGroupManagerAutoHealingPoliciesToProto converts a InstanceGroupManagerAutoHealingPolicies resource to its proto representation.
func ComputeAlphaInstanceGroupManagerAutoHealingPoliciesToProto(o *alpha.InstanceGroupManagerAutoHealingPolicies) *alphapb.ComputeAlphaInstanceGroupManagerAutoHealingPolicies {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerAutoHealingPolicies{
		HealthCheck:     dcl.ValueOrEmptyString(o.HealthCheck),
		InitialDelaySec: dcl.ValueOrEmptyInt64(o.InitialDelaySec),
	}
	return p
}

// InstanceGroupManagerUpdatePolicyToProto converts a InstanceGroupManagerUpdatePolicy resource to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyToProto(o *alpha.InstanceGroupManagerUpdatePolicy) *alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicy{
		Type:                        ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnumToProto(o.Type),
		InstanceRedistributionType:  ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(o.InstanceRedistributionType),
		MinimalAction:               ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(o.MinimalAction),
		MaxSurge:                    ComputeAlphaInstanceGroupManagerFixedOrPercentToProto(o.MaxSurge),
		MaxUnavailable:              ComputeAlphaInstanceGroupManagerFixedOrPercentToProto(o.MaxUnavailable),
		ReplacementMethod:           ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(o.ReplacementMethod),
		MostDisruptiveAllowedAction: ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto(o.MostDisruptiveAllowedAction),
		MinReadySec:                 dcl.ValueOrEmptyInt64(o.MinReadySec),
	}
	return p
}

// InstanceGroupManagerNamedPortsToProto converts a InstanceGroupManagerNamedPorts resource to its proto representation.
func ComputeAlphaInstanceGroupManagerNamedPortsToProto(o *alpha.InstanceGroupManagerNamedPorts) *alphapb.ComputeAlphaInstanceGroupManagerNamedPorts {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerNamedPorts{
		Name: dcl.ValueOrEmptyString(o.Name),
		Port: dcl.ValueOrEmptyInt64(o.Port),
	}
	return p
}

// InstanceGroupManagerStatefulPolicyToProto converts a InstanceGroupManagerStatefulPolicy resource to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyToProto(o *alpha.InstanceGroupManagerStatefulPolicy) *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicy{
		PreservedState: ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateToProto(o.PreservedState),
	}
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateToProto converts a InstanceGroupManagerStatefulPolicyPreservedState resource to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateToProto(o *alpha.InstanceGroupManagerStatefulPolicyPreservedState) *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState{}
	p.Disks = make(map[string]*alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks)
	for k, r := range o.Disks {
		p.Disks[k] = ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(&r)
	}
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks resource to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(o *alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisks) *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks{
		AutoDelete: ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(o.AutoDelete),
	}
	return p
}

// InstanceGroupManagerToProto converts a InstanceGroupManager resource to its proto representation.
func InstanceGroupManagerToProto(resource *alpha.InstanceGroupManager) *alphapb.ComputeAlphaInstanceGroupManager {
	p := &alphapb.ComputeAlphaInstanceGroupManager{
		Id:                 dcl.ValueOrEmptyInt64(resource.Id),
		CreationTimestamp:  dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		Zone:               dcl.ValueOrEmptyString(resource.Zone),
		Region:             dcl.ValueOrEmptyString(resource.Region),
		DistributionPolicy: ComputeAlphaInstanceGroupManagerDistributionPolicyToProto(resource.DistributionPolicy),
		InstanceTemplate:   dcl.ValueOrEmptyString(resource.InstanceTemplate),
		InstanceGroup:      dcl.ValueOrEmptyString(resource.InstanceGroup),
		BaseInstanceName:   dcl.ValueOrEmptyString(resource.BaseInstanceName),
		Fingerprint:        dcl.ValueOrEmptyString(resource.Fingerprint),
		CurrentActions:     ComputeAlphaInstanceGroupManagerCurrentActionsToProto(resource.CurrentActions),
		Status:             ComputeAlphaInstanceGroupManagerStatusToProto(resource.Status),
		TargetSize:         dcl.ValueOrEmptyInt64(resource.TargetSize),
		SelfLink:           dcl.ValueOrEmptyString(resource.SelfLink),
		UpdatePolicy:       ComputeAlphaInstanceGroupManagerUpdatePolicyToProto(resource.UpdatePolicy),
		StatefulPolicy:     ComputeAlphaInstanceGroupManagerStatefulPolicyToProto(resource.StatefulPolicy),
		ServiceAccount:     dcl.ValueOrEmptyString(resource.ServiceAccount),
		FailoverAction:     ComputeAlphaInstanceGroupManagerFailoverActionEnumToProto(resource.FailoverAction),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Versions {
		p.Versions = append(p.Versions, ComputeAlphaInstanceGroupManagerVersionsToProto(&r))
	}
	for _, r := range resource.TargetPools {
		p.TargetPools = append(p.TargetPools, r)
	}
	for _, r := range resource.AutoHealingPolicies {
		p.AutoHealingPolicies = append(p.AutoHealingPolicies, ComputeAlphaInstanceGroupManagerAutoHealingPoliciesToProto(&r))
	}
	for _, r := range resource.NamedPorts {
		p.NamedPorts = append(p.NamedPorts, ComputeAlphaInstanceGroupManagerNamedPortsToProto(&r))
	}

	return p
}

// ApplyInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) applyInstanceGroupManager(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaInstanceGroupManagerRequest) (*alphapb.ComputeAlphaInstanceGroupManager, error) {
	p := ProtoToInstanceGroupManager(request.GetResource())
	res, err := c.ApplyInstanceGroupManager(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceGroupManagerToProto(res)
	return r, nil
}

// ApplyInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) ApplyComputeAlphaInstanceGroupManager(ctx context.Context, request *alphapb.ApplyComputeAlphaInstanceGroupManagerRequest) (*alphapb.ComputeAlphaInstanceGroupManager, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstanceGroupManager(ctx, cl, request)
}

// DeleteInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Delete() method.
func (s *InstanceGroupManagerServer) DeleteComputeAlphaInstanceGroupManager(ctx context.Context, request *alphapb.DeleteComputeAlphaInstanceGroupManagerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstanceGroupManager(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstanceGroupManager(ctx, ProtoToInstanceGroupManager(request.GetResource()))

}

// ListComputeAlphaInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManagerList() method.
func (s *InstanceGroupManagerServer) ListComputeAlphaInstanceGroupManager(ctx context.Context, request *alphapb.ListComputeAlphaInstanceGroupManagerRequest) (*alphapb.ListComputeAlphaInstanceGroupManagerResponse, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstanceGroupManager(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaInstanceGroupManager
	for _, r := range resources.Items {
		rp := InstanceGroupManagerToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListComputeAlphaInstanceGroupManagerResponse{Items: protos}, nil
}

func createConfigInstanceGroupManager(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
