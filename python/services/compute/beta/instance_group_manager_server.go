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

// ProtoToInstanceGroupManagerDistributionPolicyTargetShapeEnum converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(e betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum) *beta.InstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerDistributionPolicyTargetShapeEnum(n[len("ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyTypeEnum converts a InstanceGroupManagerUpdatePolicyTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum) *beta.InstanceGroupManagerUpdatePolicyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyTypeEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) *beta.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"):])
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
		e := beta.InstanceGroupManagerUpdatePolicyMinimalActionEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyReplacementMethodEnum converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum) *beta.InstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyReplacementMethodEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum converts a InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum) *beta.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(e betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) *beta.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(n[len("ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerFailoverActionEnum converts a InstanceGroupManagerFailoverActionEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerFailoverActionEnum(e betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum) *beta.InstanceGroupManagerFailoverActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerFailoverActionEnum(n[len("ComputeBetaInstanceGroupManagerFailoverActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerDistributionPolicy converts a InstanceGroupManagerDistributionPolicy resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerDistributionPolicy(p *betapb.ComputeBetaInstanceGroupManagerDistributionPolicy) *beta.InstanceGroupManagerDistributionPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerDistributionPolicy{
		TargetShape: ProtoToComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(p.GetTargetShape()),
	}
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

// ProtoToInstanceGroupManagerVersions converts a InstanceGroupManagerVersions resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerVersions(p *betapb.ComputeBetaInstanceGroupManagerVersions) *beta.InstanceGroupManagerVersions {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerVersions{
		Name:             dcl.StringOrNil(p.Name),
		InstanceTemplate: dcl.StringOrNil(p.InstanceTemplate),
		TargetSize:       ProtoToComputeBetaInstanceGroupManagerFixedOrPercent(p.GetTargetSize()),
	}
	return obj
}

// ProtoToInstanceGroupManagerFixedOrPercent converts a InstanceGroupManagerFixedOrPercent resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerFixedOrPercent(p *betapb.ComputeBetaInstanceGroupManagerFixedOrPercent) *beta.InstanceGroupManagerFixedOrPercent {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerFixedOrPercent{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToInstanceGroupManagerCurrentActions converts a InstanceGroupManagerCurrentActions resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerCurrentActions(p *betapb.ComputeBetaInstanceGroupManagerCurrentActions) *beta.InstanceGroupManagerCurrentActions {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerCurrentActions{
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
func ProtoToComputeBetaInstanceGroupManagerStatus(p *betapb.ComputeBetaInstanceGroupManagerStatus) *beta.InstanceGroupManagerStatus {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatus{
		IsStable:      dcl.Bool(p.IsStable),
		VersionTarget: ProtoToComputeBetaInstanceGroupManagerStatusVersionTarget(p.GetVersionTarget()),
		Stateful:      ProtoToComputeBetaInstanceGroupManagerStatusStateful(p.GetStateful()),
		Autoscaler:    dcl.StringOrNil(p.Autoscaler),
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

// ProtoToInstanceGroupManagerStatusStateful converts a InstanceGroupManagerStatusStateful resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatusStateful(p *betapb.ComputeBetaInstanceGroupManagerStatusStateful) *beta.InstanceGroupManagerStatusStateful {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatusStateful{
		HasStatefulConfig:  dcl.Bool(p.HasStatefulConfig),
		PerInstanceConfigs: ProtoToComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs(p.GetPerInstanceConfigs()),
		IsStateful:         dcl.Bool(p.IsStateful),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStatefulPerInstanceConfigs converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs(p *betapb.ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs) *beta.InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatusStatefulPerInstanceConfigs{
		AllEffective: dcl.Bool(p.AllEffective),
	}
	return obj
}

// ProtoToInstanceGroupManagerAutoHealingPolicies converts a InstanceGroupManagerAutoHealingPolicies resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerAutoHealingPolicies(p *betapb.ComputeBetaInstanceGroupManagerAutoHealingPolicies) *beta.InstanceGroupManagerAutoHealingPolicies {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerAutoHealingPolicies{
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
		Type:                        ProtoToComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(p.GetType()),
		InstanceRedistributionType:  ProtoToComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(p.GetInstanceRedistributionType()),
		MinimalAction:               ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(p.GetMinimalAction()),
		MaxSurge:                    ProtoToComputeBetaInstanceGroupManagerFixedOrPercent(p.GetMaxSurge()),
		MaxUnavailable:              ProtoToComputeBetaInstanceGroupManagerFixedOrPercent(p.GetMaxUnavailable()),
		ReplacementMethod:           ProtoToComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(p.GetReplacementMethod()),
		MostDisruptiveAllowedAction: ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(p.GetMostDisruptiveAllowedAction()),
		MinReadySec:                 dcl.Int64OrNil(p.MinReadySec),
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

// ProtoToInstanceGroupManagerStatefulPolicy converts a InstanceGroupManagerStatefulPolicy resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicy(p *betapb.ComputeBetaInstanceGroupManagerStatefulPolicy) *beta.InstanceGroupManagerStatefulPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatefulPolicy{
		PreservedState: ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedState(p.GetPreservedState()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedState converts a InstanceGroupManagerStatefulPolicyPreservedState resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedState(p *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedState) *beta.InstanceGroupManagerStatefulPolicyPreservedState {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatefulPolicyPreservedState{}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisks converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks resource from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks(p *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks) *beta.InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatefulPolicyPreservedStateDisks{
		AutoDelete: ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManager converts a InstanceGroupManager resource from its proto representation.
func ProtoToInstanceGroupManager(p *betapb.ComputeBetaInstanceGroupManager) *beta.InstanceGroupManager {
	obj := &beta.InstanceGroupManager{
		Id:                 dcl.Int64OrNil(p.Id),
		CreationTimestamp:  dcl.StringOrNil(p.CreationTimestamp),
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		Zone:               dcl.StringOrNil(p.Zone),
		Region:             dcl.StringOrNil(p.Region),
		DistributionPolicy: ProtoToComputeBetaInstanceGroupManagerDistributionPolicy(p.GetDistributionPolicy()),
		InstanceTemplate:   dcl.StringOrNil(p.InstanceTemplate),
		InstanceGroup:      dcl.StringOrNil(p.InstanceGroup),
		BaseInstanceName:   dcl.StringOrNil(p.BaseInstanceName),
		Fingerprint:        dcl.StringOrNil(p.Fingerprint),
		CurrentActions:     ProtoToComputeBetaInstanceGroupManagerCurrentActions(p.GetCurrentActions()),
		Status:             ProtoToComputeBetaInstanceGroupManagerStatus(p.GetStatus()),
		TargetSize:         dcl.Int64OrNil(p.TargetSize),
		SelfLink:           dcl.StringOrNil(p.SelfLink),
		UpdatePolicy:       ProtoToComputeBetaInstanceGroupManagerUpdatePolicy(p.GetUpdatePolicy()),
		StatefulPolicy:     ProtoToComputeBetaInstanceGroupManagerStatefulPolicy(p.GetStatefulPolicy()),
		ServiceAccount:     dcl.StringOrNil(p.ServiceAccount),
		FailoverAction:     ProtoToComputeBetaInstanceGroupManagerFailoverActionEnum(p.GetFailoverAction()),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetVersions() {
		obj.Versions = append(obj.Versions, *ProtoToComputeBetaInstanceGroupManagerVersions(r))
	}
	for _, r := range p.GetTargetPools() {
		obj.TargetPools = append(obj.TargetPools, r)
	}
	for _, r := range p.GetAutoHealingPolicies() {
		obj.AutoHealingPolicies = append(obj.AutoHealingPolicies, *ProtoToComputeBetaInstanceGroupManagerAutoHealingPolicies(r))
	}
	for _, r := range p.GetNamedPorts() {
		obj.NamedPorts = append(obj.NamedPorts, *ProtoToComputeBetaInstanceGroupManagerNamedPorts(r))
	}
	return obj
}

// InstanceGroupManagerDistributionPolicyTargetShapeEnumToProto converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(e *beta.InstanceGroupManagerDistributionPolicyTargetShapeEnum) betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum_value["InstanceGroupManagerDistributionPolicyTargetShapeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
}

// InstanceGroupManagerUpdatePolicyTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyTypeEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyTypeEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum_value["InstanceGroupManagerUpdatePolicyTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(0)
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

// InstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyReplacementMethodEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum_value["InstanceGroupManagerUpdatePolicyReplacementMethodEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
}

// InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum_value["InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(e *beta.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
}

// InstanceGroupManagerFailoverActionEnumToProto converts a InstanceGroupManagerFailoverActionEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerFailoverActionEnumToProto(e *beta.InstanceGroupManagerFailoverActionEnum) betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum_value["InstanceGroupManagerFailoverActionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum(0)
}

// InstanceGroupManagerDistributionPolicyToProto converts a InstanceGroupManagerDistributionPolicy resource to its proto representation.
func ComputeBetaInstanceGroupManagerDistributionPolicyToProto(o *beta.InstanceGroupManagerDistributionPolicy) *betapb.ComputeBetaInstanceGroupManagerDistributionPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerDistributionPolicy{
		TargetShape: ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(o.TargetShape),
	}
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

// InstanceGroupManagerVersionsToProto converts a InstanceGroupManagerVersions resource to its proto representation.
func ComputeBetaInstanceGroupManagerVersionsToProto(o *beta.InstanceGroupManagerVersions) *betapb.ComputeBetaInstanceGroupManagerVersions {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerVersions{
		Name:             dcl.ValueOrEmptyString(o.Name),
		InstanceTemplate: dcl.ValueOrEmptyString(o.InstanceTemplate),
		TargetSize:       ComputeBetaInstanceGroupManagerFixedOrPercentToProto(o.TargetSize),
	}
	return p
}

// InstanceGroupManagerFixedOrPercentToProto converts a InstanceGroupManagerFixedOrPercent resource to its proto representation.
func ComputeBetaInstanceGroupManagerFixedOrPercentToProto(o *beta.InstanceGroupManagerFixedOrPercent) *betapb.ComputeBetaInstanceGroupManagerFixedOrPercent {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerFixedOrPercent{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// InstanceGroupManagerCurrentActionsToProto converts a InstanceGroupManagerCurrentActions resource to its proto representation.
func ComputeBetaInstanceGroupManagerCurrentActionsToProto(o *beta.InstanceGroupManagerCurrentActions) *betapb.ComputeBetaInstanceGroupManagerCurrentActions {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerCurrentActions{
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
func ComputeBetaInstanceGroupManagerStatusToProto(o *beta.InstanceGroupManagerStatus) *betapb.ComputeBetaInstanceGroupManagerStatus {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatus{
		IsStable:      dcl.ValueOrEmptyBool(o.IsStable),
		VersionTarget: ComputeBetaInstanceGroupManagerStatusVersionTargetToProto(o.VersionTarget),
		Stateful:      ComputeBetaInstanceGroupManagerStatusStatefulToProto(o.Stateful),
		Autoscaler:    dcl.ValueOrEmptyString(o.Autoscaler),
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

// InstanceGroupManagerStatusStatefulToProto converts a InstanceGroupManagerStatusStateful resource to its proto representation.
func ComputeBetaInstanceGroupManagerStatusStatefulToProto(o *beta.InstanceGroupManagerStatusStateful) *betapb.ComputeBetaInstanceGroupManagerStatusStateful {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatusStateful{
		HasStatefulConfig:  dcl.ValueOrEmptyBool(o.HasStatefulConfig),
		PerInstanceConfigs: ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o.PerInstanceConfigs),
		IsStateful:         dcl.ValueOrEmptyBool(o.IsStateful),
	}
	return p
}

// InstanceGroupManagerStatusStatefulPerInstanceConfigsToProto converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs resource to its proto representation.
func ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o *beta.InstanceGroupManagerStatusStatefulPerInstanceConfigs) *betapb.ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs{
		AllEffective: dcl.ValueOrEmptyBool(o.AllEffective),
	}
	return p
}

// InstanceGroupManagerAutoHealingPoliciesToProto converts a InstanceGroupManagerAutoHealingPolicies resource to its proto representation.
func ComputeBetaInstanceGroupManagerAutoHealingPoliciesToProto(o *beta.InstanceGroupManagerAutoHealingPolicies) *betapb.ComputeBetaInstanceGroupManagerAutoHealingPolicies {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerAutoHealingPolicies{
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
		Type:                        ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnumToProto(o.Type),
		InstanceRedistributionType:  ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(o.InstanceRedistributionType),
		MinimalAction:               ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(o.MinimalAction),
		MaxSurge:                    ComputeBetaInstanceGroupManagerFixedOrPercentToProto(o.MaxSurge),
		MaxUnavailable:              ComputeBetaInstanceGroupManagerFixedOrPercentToProto(o.MaxUnavailable),
		ReplacementMethod:           ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(o.ReplacementMethod),
		MostDisruptiveAllowedAction: ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto(o.MostDisruptiveAllowedAction),
		MinReadySec:                 dcl.ValueOrEmptyInt64(o.MinReadySec),
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

// InstanceGroupManagerStatefulPolicyToProto converts a InstanceGroupManagerStatefulPolicy resource to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyToProto(o *beta.InstanceGroupManagerStatefulPolicy) *betapb.ComputeBetaInstanceGroupManagerStatefulPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatefulPolicy{
		PreservedState: ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateToProto(o.PreservedState),
	}
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateToProto converts a InstanceGroupManagerStatefulPolicyPreservedState resource to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateToProto(o *beta.InstanceGroupManagerStatefulPolicyPreservedState) *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedState {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedState{}
	p.Disks = make(map[string]*betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks)
	for k, r := range o.Disks {
		p.Disks[k] = ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(&r)
	}
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks resource to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(o *beta.InstanceGroupManagerStatefulPolicyPreservedStateDisks) *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks{
		AutoDelete: ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(o.AutoDelete),
	}
	return p
}

// InstanceGroupManagerToProto converts a InstanceGroupManager resource to its proto representation.
func InstanceGroupManagerToProto(resource *beta.InstanceGroupManager) *betapb.ComputeBetaInstanceGroupManager {
	p := &betapb.ComputeBetaInstanceGroupManager{
		Id:                 dcl.ValueOrEmptyInt64(resource.Id),
		CreationTimestamp:  dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		Zone:               dcl.ValueOrEmptyString(resource.Zone),
		Region:             dcl.ValueOrEmptyString(resource.Region),
		DistributionPolicy: ComputeBetaInstanceGroupManagerDistributionPolicyToProto(resource.DistributionPolicy),
		InstanceTemplate:   dcl.ValueOrEmptyString(resource.InstanceTemplate),
		InstanceGroup:      dcl.ValueOrEmptyString(resource.InstanceGroup),
		BaseInstanceName:   dcl.ValueOrEmptyString(resource.BaseInstanceName),
		Fingerprint:        dcl.ValueOrEmptyString(resource.Fingerprint),
		CurrentActions:     ComputeBetaInstanceGroupManagerCurrentActionsToProto(resource.CurrentActions),
		Status:             ComputeBetaInstanceGroupManagerStatusToProto(resource.Status),
		TargetSize:         dcl.ValueOrEmptyInt64(resource.TargetSize),
		SelfLink:           dcl.ValueOrEmptyString(resource.SelfLink),
		UpdatePolicy:       ComputeBetaInstanceGroupManagerUpdatePolicyToProto(resource.UpdatePolicy),
		StatefulPolicy:     ComputeBetaInstanceGroupManagerStatefulPolicyToProto(resource.StatefulPolicy),
		ServiceAccount:     dcl.ValueOrEmptyString(resource.ServiceAccount),
		FailoverAction:     ComputeBetaInstanceGroupManagerFailoverActionEnumToProto(resource.FailoverAction),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Versions {
		p.Versions = append(p.Versions, ComputeBetaInstanceGroupManagerVersionsToProto(&r))
	}
	for _, r := range resource.TargetPools {
		p.TargetPools = append(p.TargetPools, r)
	}
	for _, r := range resource.AutoHealingPolicies {
		p.AutoHealingPolicies = append(p.AutoHealingPolicies, ComputeBetaInstanceGroupManagerAutoHealingPoliciesToProto(&r))
	}
	for _, r := range resource.NamedPorts {
		p.NamedPorts = append(p.NamedPorts, ComputeBetaInstanceGroupManagerNamedPortsToProto(&r))
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

// ListComputeBetaInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManagerList() method.
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
