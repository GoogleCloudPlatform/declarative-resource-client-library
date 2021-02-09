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
	appenginepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/appengine/appengine_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/appengine"
)

// Server implements the gRPC interface for Version.
type VersionServer struct{}

// ProtoToVersionInboundServicesEnum converts a VersionInboundServicesEnum enum from its proto representation.
func ProtoToAppengineVersionInboundServicesEnum(e appenginepb.AppengineVersionInboundServicesEnum) *appengine.VersionInboundServicesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionInboundServicesEnum_name[int32(e)]; ok {
		e := appengine.VersionInboundServicesEnum(n[len("VersionInboundServicesEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionServingStatusEnum converts a VersionServingStatusEnum enum from its proto representation.
func ProtoToAppengineVersionServingStatusEnum(e appenginepb.AppengineVersionServingStatusEnum) *appengine.VersionServingStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionServingStatusEnum_name[int32(e)]; ok {
		e := appengine.VersionServingStatusEnum(n[len("VersionServingStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionHandlersSecurityLevelEnum converts a VersionHandlersSecurityLevelEnum enum from its proto representation.
func ProtoToAppengineVersionHandlersSecurityLevelEnum(e appenginepb.AppengineVersionHandlersSecurityLevelEnum) *appengine.VersionHandlersSecurityLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionHandlersSecurityLevelEnum_name[int32(e)]; ok {
		e := appengine.VersionHandlersSecurityLevelEnum(n[len("VersionHandlersSecurityLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionHandlersLoginEnum converts a VersionHandlersLoginEnum enum from its proto representation.
func ProtoToAppengineVersionHandlersLoginEnum(e appenginepb.AppengineVersionHandlersLoginEnum) *appengine.VersionHandlersLoginEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionHandlersLoginEnum_name[int32(e)]; ok {
		e := appengine.VersionHandlersLoginEnum(n[len("VersionHandlersLoginEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionHandlersAuthFailActionEnum converts a VersionHandlersAuthFailActionEnum enum from its proto representation.
func ProtoToAppengineVersionHandlersAuthFailActionEnum(e appenginepb.AppengineVersionHandlersAuthFailActionEnum) *appengine.VersionHandlersAuthFailActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionHandlersAuthFailActionEnum_name[int32(e)]; ok {
		e := appengine.VersionHandlersAuthFailActionEnum(n[len("VersionHandlersAuthFailActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionHandlersRedirectHttpResponseCodeEnum converts a VersionHandlersRedirectHttpResponseCodeEnum enum from its proto representation.
func ProtoToAppengineVersionHandlersRedirectHttpResponseCodeEnum(e appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum) *appengine.VersionHandlersRedirectHttpResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum_name[int32(e)]; ok {
		e := appengine.VersionHandlersRedirectHttpResponseCodeEnum(n[len("VersionHandlersRedirectHttpResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionErrorHandlersErrorCodeEnum converts a VersionErrorHandlersErrorCodeEnum enum from its proto representation.
func ProtoToAppengineVersionErrorHandlersErrorCodeEnum(e appenginepb.AppengineVersionErrorHandlersErrorCodeEnum) *appengine.VersionErrorHandlersErrorCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionErrorHandlersErrorCodeEnum_name[int32(e)]; ok {
		e := appengine.VersionErrorHandlersErrorCodeEnum(n[len("VersionErrorHandlersErrorCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionApiConfigAuthFailActionEnum converts a VersionApiConfigAuthFailActionEnum enum from its proto representation.
func ProtoToAppengineVersionApiConfigAuthFailActionEnum(e appenginepb.AppengineVersionApiConfigAuthFailActionEnum) *appengine.VersionApiConfigAuthFailActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionApiConfigAuthFailActionEnum_name[int32(e)]; ok {
		e := appengine.VersionApiConfigAuthFailActionEnum(n[len("VersionApiConfigAuthFailActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionApiConfigLoginEnum converts a VersionApiConfigLoginEnum enum from its proto representation.
func ProtoToAppengineVersionApiConfigLoginEnum(e appenginepb.AppengineVersionApiConfigLoginEnum) *appengine.VersionApiConfigLoginEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionApiConfigLoginEnum_name[int32(e)]; ok {
		e := appengine.VersionApiConfigLoginEnum(n[len("VersionApiConfigLoginEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionApiConfigSecurityLevelEnum converts a VersionApiConfigSecurityLevelEnum enum from its proto representation.
func ProtoToAppengineVersionApiConfigSecurityLevelEnum(e appenginepb.AppengineVersionApiConfigSecurityLevelEnum) *appengine.VersionApiConfigSecurityLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionApiConfigSecurityLevelEnum_name[int32(e)]; ok {
		e := appengine.VersionApiConfigSecurityLevelEnum(n[len("VersionApiConfigSecurityLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionVPCAccessConnectorEgressSettingEnum converts a VersionVPCAccessConnectorEgressSettingEnum enum from its proto representation.
func ProtoToAppengineVersionVPCAccessConnectorEgressSettingEnum(e appenginepb.AppengineVersionVPCAccessConnectorEgressSettingEnum) *appengine.VersionVPCAccessConnectorEgressSettingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionVPCAccessConnectorEgressSettingEnum_name[int32(e)]; ok {
		e := appengine.VersionVPCAccessConnectorEgressSettingEnum(n[len("VersionVPCAccessConnectorEgressSettingEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionNetworkSettingsIngressTrafficAllowedEnum converts a VersionNetworkSettingsIngressTrafficAllowedEnum enum from its proto representation.
func ProtoToAppengineVersionNetworkSettingsIngressTrafficAllowedEnum(e appenginepb.AppengineVersionNetworkSettingsIngressTrafficAllowedEnum) *appengine.VersionNetworkSettingsIngressTrafficAllowedEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionNetworkSettingsIngressTrafficAllowedEnum_name[int32(e)]; ok {
		e := appengine.VersionNetworkSettingsIngressTrafficAllowedEnum(n[len("VersionNetworkSettingsIngressTrafficAllowedEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionInstanceSpecPortsProtocolEnum converts a VersionInstanceSpecPortsProtocolEnum enum from its proto representation.
func ProtoToAppengineVersionInstanceSpecPortsProtocolEnum(e appenginepb.AppengineVersionInstanceSpecPortsProtocolEnum) *appengine.VersionInstanceSpecPortsProtocolEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionInstanceSpecPortsProtocolEnum_name[int32(e)]; ok {
		e := appengine.VersionInstanceSpecPortsProtocolEnum(n[len("VersionInstanceSpecPortsProtocolEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionAutomaticScaling converts a VersionAutomaticScaling resource from its proto representation.
func ProtoToAppengineVersionAutomaticScaling(p *appenginepb.AppengineVersionAutomaticScaling) *appengine.VersionAutomaticScaling {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScaling{
		CoolDownPeriod:            ProtoToAppengineVersionAutomaticScalingCoolDownPeriod(p.GetCoolDownPeriod()),
		CpuUtilization:            ProtoToAppengineVersionAutomaticScalingCpuUtilization(p.GetCpuUtilization()),
		MaxConcurrentRequests:     dcl.Int64OrNil(p.MaxConcurrentRequests),
		MaxIdleInstances:          dcl.Int64OrNil(p.MaxIdleInstances),
		MaxTotalInstances:         dcl.Int64OrNil(p.MaxTotalInstances),
		MaxPendingLatency:         ProtoToAppengineVersionAutomaticScalingMaxPendingLatency(p.GetMaxPendingLatency()),
		MinIdleInstances:          dcl.Int64OrNil(p.MinIdleInstances),
		MinTotalInstances:         dcl.Int64OrNil(p.MinTotalInstances),
		MinPendingLatency:         ProtoToAppengineVersionAutomaticScalingMinPendingLatency(p.GetMinPendingLatency()),
		RequestUtilization:        ProtoToAppengineVersionAutomaticScalingRequestUtilization(p.GetRequestUtilization()),
		DiskUtilization:           ProtoToAppengineVersionAutomaticScalingDiskUtilization(p.GetDiskUtilization()),
		NetworkUtilization:        ProtoToAppengineVersionAutomaticScalingNetworkUtilization(p.GetNetworkUtilization()),
		StandardSchedulerSettings: ProtoToAppengineVersionAutomaticScalingStandardSchedulerSettings(p.GetStandardSchedulerSettings()),
	}
	return obj
}

// ProtoToVersionAutomaticScalingCoolDownPeriod converts a VersionAutomaticScalingCoolDownPeriod resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingCoolDownPeriod(p *appenginepb.AppengineVersionAutomaticScalingCoolDownPeriod) *appengine.VersionAutomaticScalingCoolDownPeriod {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingCoolDownPeriod{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionAutomaticScalingCpuUtilization converts a VersionAutomaticScalingCpuUtilization resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingCpuUtilization(p *appenginepb.AppengineVersionAutomaticScalingCpuUtilization) *appengine.VersionAutomaticScalingCpuUtilization {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingCpuUtilization{
		AggregationWindowLength: ProtoToAppengineVersionAutomaticScalingCpuUtilizationAggregationWindowLength(p.GetAggregationWindowLength()),
		TargetUtilization:       dcl.Float64OrNil(p.TargetUtilization),
	}
	return obj
}

// ProtoToVersionAutomaticScalingCpuUtilizationAggregationWindowLength converts a VersionAutomaticScalingCpuUtilizationAggregationWindowLength resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingCpuUtilizationAggregationWindowLength(p *appenginepb.AppengineVersionAutomaticScalingCpuUtilizationAggregationWindowLength) *appengine.VersionAutomaticScalingCpuUtilizationAggregationWindowLength {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingCpuUtilizationAggregationWindowLength{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionAutomaticScalingMaxPendingLatency converts a VersionAutomaticScalingMaxPendingLatency resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingMaxPendingLatency(p *appenginepb.AppengineVersionAutomaticScalingMaxPendingLatency) *appengine.VersionAutomaticScalingMaxPendingLatency {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingMaxPendingLatency{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionAutomaticScalingMinPendingLatency converts a VersionAutomaticScalingMinPendingLatency resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingMinPendingLatency(p *appenginepb.AppengineVersionAutomaticScalingMinPendingLatency) *appengine.VersionAutomaticScalingMinPendingLatency {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingMinPendingLatency{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionAutomaticScalingRequestUtilization converts a VersionAutomaticScalingRequestUtilization resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingRequestUtilization(p *appenginepb.AppengineVersionAutomaticScalingRequestUtilization) *appengine.VersionAutomaticScalingRequestUtilization {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingRequestUtilization{
		TargetRequestCountPerSecond: dcl.Int64OrNil(p.TargetRequestCountPerSecond),
		TargetConcurrentRequests:    dcl.Int64OrNil(p.TargetConcurrentRequests),
	}
	return obj
}

// ProtoToVersionAutomaticScalingDiskUtilization converts a VersionAutomaticScalingDiskUtilization resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingDiskUtilization(p *appenginepb.AppengineVersionAutomaticScalingDiskUtilization) *appengine.VersionAutomaticScalingDiskUtilization {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingDiskUtilization{
		TargetWriteBytesPerSecond: dcl.Int64OrNil(p.TargetWriteBytesPerSecond),
		TargetWriteOpsPerSecond:   dcl.Int64OrNil(p.TargetWriteOpsPerSecond),
		TargetReadBytesPerSecond:  dcl.Int64OrNil(p.TargetReadBytesPerSecond),
		TargetReadOpsPerSecond:    dcl.Int64OrNil(p.TargetReadOpsPerSecond),
	}
	return obj
}

// ProtoToVersionAutomaticScalingNetworkUtilization converts a VersionAutomaticScalingNetworkUtilization resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingNetworkUtilization(p *appenginepb.AppengineVersionAutomaticScalingNetworkUtilization) *appengine.VersionAutomaticScalingNetworkUtilization {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingNetworkUtilization{
		TargetSentBytesPerSecond:       dcl.Int64OrNil(p.TargetSentBytesPerSecond),
		TargetSentPacketsPerSecond:     dcl.Int64OrNil(p.TargetSentPacketsPerSecond),
		TargetReceivedBytesPerSecond:   dcl.Int64OrNil(p.TargetReceivedBytesPerSecond),
		TargetReceivedPacketsPerSecond: dcl.Int64OrNil(p.TargetReceivedPacketsPerSecond),
	}
	return obj
}

// ProtoToVersionAutomaticScalingStandardSchedulerSettings converts a VersionAutomaticScalingStandardSchedulerSettings resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingStandardSchedulerSettings(p *appenginepb.AppengineVersionAutomaticScalingStandardSchedulerSettings) *appengine.VersionAutomaticScalingStandardSchedulerSettings {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingStandardSchedulerSettings{
		TargetCpuUtilization:        dcl.Float64OrNil(p.TargetCpuUtilization),
		TargetThroughputUtilization: dcl.Float64OrNil(p.TargetThroughputUtilization),
		MinInstances:                dcl.Int64OrNil(p.MinInstances),
		MaxInstances:                dcl.Int64OrNil(p.MaxInstances),
	}
	return obj
}

// ProtoToVersionBasicScaling converts a VersionBasicScaling resource from its proto representation.
func ProtoToAppengineVersionBasicScaling(p *appenginepb.AppengineVersionBasicScaling) *appengine.VersionBasicScaling {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionBasicScaling{
		IdleTimeout:  ProtoToAppengineVersionBasicScalingIdleTimeout(p.GetIdleTimeout()),
		MaxInstances: dcl.Int64OrNil(p.MaxInstances),
	}
	return obj
}

// ProtoToVersionBasicScalingIdleTimeout converts a VersionBasicScalingIdleTimeout resource from its proto representation.
func ProtoToAppengineVersionBasicScalingIdleTimeout(p *appenginepb.AppengineVersionBasicScalingIdleTimeout) *appengine.VersionBasicScalingIdleTimeout {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionBasicScalingIdleTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionManualScaling converts a VersionManualScaling resource from its proto representation.
func ProtoToAppengineVersionManualScaling(p *appenginepb.AppengineVersionManualScaling) *appengine.VersionManualScaling {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionManualScaling{
		Instances: dcl.Int64OrNil(p.Instances),
	}
	return obj
}

// ProtoToVersionJobScaling converts a VersionJobScaling resource from its proto representation.
func ProtoToAppengineVersionJobScaling(p *appenginepb.AppengineVersionJobScaling) *appengine.VersionJobScaling {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionJobScaling{
		Completions:               dcl.Int64OrNil(p.Completions),
		Parallelism:               dcl.Int64OrNil(p.Parallelism),
		JobDeadline:               ProtoToAppengineVersionJobScalingJobDeadline(p.GetJobDeadline()),
		InstanceRetries:           dcl.Int64OrNil(p.InstanceRetries),
		InstanceDeadline:          ProtoToAppengineVersionJobScalingInstanceDeadline(p.GetInstanceDeadline()),
		InstanceTerminationWindow: ProtoToAppengineVersionJobScalingInstanceTerminationWindow(p.GetInstanceTerminationWindow()),
	}
	return obj
}

// ProtoToVersionJobScalingJobDeadline converts a VersionJobScalingJobDeadline resource from its proto representation.
func ProtoToAppengineVersionJobScalingJobDeadline(p *appenginepb.AppengineVersionJobScalingJobDeadline) *appengine.VersionJobScalingJobDeadline {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionJobScalingJobDeadline{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionJobScalingInstanceDeadline converts a VersionJobScalingInstanceDeadline resource from its proto representation.
func ProtoToAppengineVersionJobScalingInstanceDeadline(p *appenginepb.AppengineVersionJobScalingInstanceDeadline) *appengine.VersionJobScalingInstanceDeadline {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionJobScalingInstanceDeadline{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionJobScalingInstanceTerminationWindow converts a VersionJobScalingInstanceTerminationWindow resource from its proto representation.
func ProtoToAppengineVersionJobScalingInstanceTerminationWindow(p *appenginepb.AppengineVersionJobScalingInstanceTerminationWindow) *appengine.VersionJobScalingInstanceTerminationWindow {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionJobScalingInstanceTerminationWindow{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionPoolScaling converts a VersionPoolScaling resource from its proto representation.
func ProtoToAppengineVersionPoolScaling(p *appenginepb.AppengineVersionPoolScaling) *appengine.VersionPoolScaling {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionPoolScaling{
		Replicas:       dcl.Int64OrNil(p.Replicas),
		MaxUnavailable: dcl.Int64OrNil(p.MaxUnavailable),
		MaxSurge:       dcl.Int64OrNil(p.MaxSurge),
	}
	return obj
}

// ProtoToVersionNetwork converts a VersionNetwork resource from its proto representation.
func ProtoToAppengineVersionNetwork(p *appenginepb.AppengineVersionNetwork) *appengine.VersionNetwork {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionNetwork{
		InstanceTag:     dcl.StringOrNil(p.InstanceTag),
		Name:            dcl.StringOrNil(p.Name),
		SubnetworkName:  dcl.StringOrNil(p.SubnetworkName),
		SessionAffinity: dcl.Bool(p.SessionAffinity),
	}
	for _, r := range p.GetForwardedPorts() {
		obj.ForwardedPorts = append(obj.ForwardedPorts, r)
	}
	return obj
}

// ProtoToVersionResources converts a VersionResources resource from its proto representation.
func ProtoToAppengineVersionResources(p *appenginepb.AppengineVersionResources) *appengine.VersionResources {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionResources{
		Cpu:             dcl.Float64OrNil(p.Cpu),
		DiskGb:          dcl.Float64OrNil(p.DiskGb),
		MemoryGb:        dcl.Float64OrNil(p.MemoryGb),
		KmsKeyReference: dcl.StringOrNil(p.KmsKeyReference),
	}
	for _, r := range p.GetVolumes() {
		obj.Volumes = append(obj.Volumes, *ProtoToAppengineVersionResourcesVolumes(r))
	}
	return obj
}

// ProtoToVersionResourcesVolumes converts a VersionResourcesVolumes resource from its proto representation.
func ProtoToAppengineVersionResourcesVolumes(p *appenginepb.AppengineVersionResourcesVolumes) *appengine.VersionResourcesVolumes {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionResourcesVolumes{
		Name:       dcl.StringOrNil(p.Name),
		VolumeType: dcl.StringOrNil(p.VolumeType),
		SizeGb:     dcl.Float64OrNil(p.SizeGb),
	}
	return obj
}

// ProtoToVersionBetaSettings converts a VersionBetaSettings resource from its proto representation.
func ProtoToAppengineVersionBetaSettings(p *appenginepb.AppengineVersionBetaSettings) *appengine.VersionBetaSettings {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionBetaSettings{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToVersionHandlers converts a VersionHandlers resource from its proto representation.
func ProtoToAppengineVersionHandlers(p *appenginepb.AppengineVersionHandlers) *appengine.VersionHandlers {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlers{
		UrlRegex:                 dcl.StringOrNil(p.UrlRegex),
		StaticFiles:              ProtoToAppengineVersionHandlersStaticFiles(p.GetStaticFiles()),
		Script:                   ProtoToAppengineVersionHandlersScript(p.GetScript()),
		ApiEndpoint:              ProtoToAppengineVersionHandlersApiEndpoint(p.GetApiEndpoint()),
		SecurityLevel:            ProtoToAppengineVersionHandlersSecurityLevelEnum(p.GetSecurityLevel()),
		Login:                    ProtoToAppengineVersionHandlersLoginEnum(p.GetLogin()),
		AuthFailAction:           ProtoToAppengineVersionHandlersAuthFailActionEnum(p.GetAuthFailAction()),
		RedirectHttpResponseCode: ProtoToAppengineVersionHandlersRedirectHttpResponseCodeEnum(p.GetRedirectHttpResponseCode()),
	}
	return obj
}

// ProtoToVersionHandlersStaticFiles converts a VersionHandlersStaticFiles resource from its proto representation.
func ProtoToAppengineVersionHandlersStaticFiles(p *appenginepb.AppengineVersionHandlersStaticFiles) *appengine.VersionHandlersStaticFiles {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlersStaticFiles{
		Path:                dcl.StringOrNil(p.Path),
		UploadPathRegex:     dcl.StringOrNil(p.UploadPathRegex),
		MimeType:            dcl.StringOrNil(p.MimeType),
		Expiration:          ProtoToAppengineVersionHandlersStaticFilesExpiration(p.GetExpiration()),
		RequireMatchingFile: dcl.Bool(p.RequireMatchingFile),
		ApplicationReadable: dcl.Bool(p.ApplicationReadable),
	}
	for _, r := range p.GetHttpHeaders() {
		obj.HttpHeaders = append(obj.HttpHeaders, *ProtoToAppengineVersionHandlersStaticFilesHttpHeaders(r))
	}
	return obj
}

// ProtoToVersionHandlersStaticFilesHttpHeaders converts a VersionHandlersStaticFilesHttpHeaders resource from its proto representation.
func ProtoToAppengineVersionHandlersStaticFilesHttpHeaders(p *appenginepb.AppengineVersionHandlersStaticFilesHttpHeaders) *appengine.VersionHandlersStaticFilesHttpHeaders {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlersStaticFilesHttpHeaders{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToVersionHandlersStaticFilesExpiration converts a VersionHandlersStaticFilesExpiration resource from its proto representation.
func ProtoToAppengineVersionHandlersStaticFilesExpiration(p *appenginepb.AppengineVersionHandlersStaticFilesExpiration) *appengine.VersionHandlersStaticFilesExpiration {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlersStaticFilesExpiration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionHandlersScript converts a VersionHandlersScript resource from its proto representation.
func ProtoToAppengineVersionHandlersScript(p *appenginepb.AppengineVersionHandlersScript) *appengine.VersionHandlersScript {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlersScript{
		ScriptPath: dcl.StringOrNil(p.ScriptPath),
	}
	return obj
}

// ProtoToVersionHandlersApiEndpoint converts a VersionHandlersApiEndpoint resource from its proto representation.
func ProtoToAppengineVersionHandlersApiEndpoint(p *appenginepb.AppengineVersionHandlersApiEndpoint) *appengine.VersionHandlersApiEndpoint {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlersApiEndpoint{
		ScriptPath: dcl.StringOrNil(p.ScriptPath),
	}
	return obj
}

// ProtoToVersionErrorHandlers converts a VersionErrorHandlers resource from its proto representation.
func ProtoToAppengineVersionErrorHandlers(p *appenginepb.AppengineVersionErrorHandlers) *appengine.VersionErrorHandlers {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionErrorHandlers{
		ErrorCode:  ProtoToAppengineVersionErrorHandlersErrorCodeEnum(p.GetErrorCode()),
		StaticFile: dcl.StringOrNil(p.StaticFile),
		MimeType:   dcl.StringOrNil(p.MimeType),
	}
	return obj
}

// ProtoToVersionLibraries converts a VersionLibraries resource from its proto representation.
func ProtoToAppengineVersionLibraries(p *appenginepb.AppengineVersionLibraries) *appengine.VersionLibraries {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionLibraries{
		Name:    dcl.StringOrNil(p.Name),
		Version: dcl.StringOrNil(p.Version),
	}
	return obj
}

// ProtoToVersionApiConfig converts a VersionApiConfig resource from its proto representation.
func ProtoToAppengineVersionApiConfig(p *appenginepb.AppengineVersionApiConfig) *appengine.VersionApiConfig {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionApiConfig{
		AuthFailAction: ProtoToAppengineVersionApiConfigAuthFailActionEnum(p.GetAuthFailAction()),
		Login:          ProtoToAppengineVersionApiConfigLoginEnum(p.GetLogin()),
		Script:         dcl.StringOrNil(p.Script),
		SecurityLevel:  ProtoToAppengineVersionApiConfigSecurityLevelEnum(p.GetSecurityLevel()),
		Url:            dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToVersionEnvVariables converts a VersionEnvVariables resource from its proto representation.
func ProtoToAppengineVersionEnvVariables(p *appenginepb.AppengineVersionEnvVariables) *appengine.VersionEnvVariables {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionEnvVariables{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToVersionBuildEnvVariables converts a VersionBuildEnvVariables resource from its proto representation.
func ProtoToAppengineVersionBuildEnvVariables(p *appenginepb.AppengineVersionBuildEnvVariables) *appengine.VersionBuildEnvVariables {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionBuildEnvVariables{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToVersionDefaultExpiration converts a VersionDefaultExpiration resource from its proto representation.
func ProtoToAppengineVersionDefaultExpiration(p *appenginepb.AppengineVersionDefaultExpiration) *appengine.VersionDefaultExpiration {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDefaultExpiration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionDeployment converts a VersionDeployment resource from its proto representation.
func ProtoToAppengineVersionDeployment(p *appenginepb.AppengineVersionDeployment) *appengine.VersionDeployment {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeployment{
		Container:         ProtoToAppengineVersionDeploymentContainer(p.GetContainer()),
		Zip:               ProtoToAppengineVersionDeploymentZip(p.GetZip()),
		CloudBuildOptions: ProtoToAppengineVersionDeploymentCloudBuildOptions(p.GetCloudBuildOptions()),
	}
	return obj
}

// ProtoToVersionDeploymentFiles converts a VersionDeploymentFiles resource from its proto representation.
func ProtoToAppengineVersionDeploymentFiles(p *appenginepb.AppengineVersionDeploymentFiles) *appengine.VersionDeploymentFiles {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeploymentFiles{
		SourceUrl: dcl.StringOrNil(p.SourceUrl),
		Sha1Sum:   dcl.StringOrNil(p.Sha1Sum),
		MimeType:  dcl.StringOrNil(p.MimeType),
	}
	return obj
}

// ProtoToVersionDeploymentContainer converts a VersionDeploymentContainer resource from its proto representation.
func ProtoToAppengineVersionDeploymentContainer(p *appenginepb.AppengineVersionDeploymentContainer) *appengine.VersionDeploymentContainer {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeploymentContainer{
		Image: dcl.StringOrNil(p.Image),
	}
	return obj
}

// ProtoToVersionDeploymentZip converts a VersionDeploymentZip resource from its proto representation.
func ProtoToAppengineVersionDeploymentZip(p *appenginepb.AppengineVersionDeploymentZip) *appengine.VersionDeploymentZip {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeploymentZip{
		SourceUrl:  dcl.StringOrNil(p.SourceUrl),
		FilesCount: dcl.Int64OrNil(p.FilesCount),
	}
	return obj
}

// ProtoToVersionDeploymentCloudBuildOptions converts a VersionDeploymentCloudBuildOptions resource from its proto representation.
func ProtoToAppengineVersionDeploymentCloudBuildOptions(p *appenginepb.AppengineVersionDeploymentCloudBuildOptions) *appengine.VersionDeploymentCloudBuildOptions {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeploymentCloudBuildOptions{
		AppYamlPath:       dcl.StringOrNil(p.AppYamlPath),
		CloudBuildTimeout: ProtoToAppengineVersionDeploymentCloudBuildOptionsCloudBuildTimeout(p.GetCloudBuildTimeout()),
	}
	return obj
}

// ProtoToVersionDeploymentCloudBuildOptionsCloudBuildTimeout converts a VersionDeploymentCloudBuildOptionsCloudBuildTimeout resource from its proto representation.
func ProtoToAppengineVersionDeploymentCloudBuildOptionsCloudBuildTimeout(p *appenginepb.AppengineVersionDeploymentCloudBuildOptionsCloudBuildTimeout) *appengine.VersionDeploymentCloudBuildOptionsCloudBuildTimeout {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeploymentCloudBuildOptionsCloudBuildTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionHealthCheck converts a VersionHealthCheck resource from its proto representation.
func ProtoToAppengineVersionHealthCheck(p *appenginepb.AppengineVersionHealthCheck) *appengine.VersionHealthCheck {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHealthCheck{
		DisableHealthCheck: dcl.Bool(p.DisableHealthCheck),
		Host:               dcl.StringOrNil(p.Host),
		HealthyThreshold:   dcl.Int64OrNil(p.HealthyThreshold),
		UnhealthyThreshold: dcl.Int64OrNil(p.UnhealthyThreshold),
		RestartThreshold:   dcl.Int64OrNil(p.RestartThreshold),
		CheckInterval:      ProtoToAppengineVersionHealthCheckCheckInterval(p.GetCheckInterval()),
		Timeout:            ProtoToAppengineVersionHealthCheckTimeout(p.GetTimeout()),
	}
	return obj
}

// ProtoToVersionHealthCheckCheckInterval converts a VersionHealthCheckCheckInterval resource from its proto representation.
func ProtoToAppengineVersionHealthCheckCheckInterval(p *appenginepb.AppengineVersionHealthCheckCheckInterval) *appengine.VersionHealthCheckCheckInterval {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHealthCheckCheckInterval{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionHealthCheckTimeout converts a VersionHealthCheckTimeout resource from its proto representation.
func ProtoToAppengineVersionHealthCheckTimeout(p *appenginepb.AppengineVersionHealthCheckTimeout) *appengine.VersionHealthCheckTimeout {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHealthCheckTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionReadinessCheck converts a VersionReadinessCheck resource from its proto representation.
func ProtoToAppengineVersionReadinessCheck(p *appenginepb.AppengineVersionReadinessCheck) *appengine.VersionReadinessCheck {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionReadinessCheck{
		Path:             dcl.StringOrNil(p.Path),
		Host:             dcl.StringOrNil(p.Host),
		FailureThreshold: dcl.Int64OrNil(p.FailureThreshold),
		SuccessThreshold: dcl.Int64OrNil(p.SuccessThreshold),
		CheckInterval:    ProtoToAppengineVersionReadinessCheckCheckInterval(p.GetCheckInterval()),
		Timeout:          ProtoToAppengineVersionReadinessCheckTimeout(p.GetTimeout()),
		AppStartTimeout:  ProtoToAppengineVersionReadinessCheckAppStartTimeout(p.GetAppStartTimeout()),
	}
	return obj
}

// ProtoToVersionReadinessCheckCheckInterval converts a VersionReadinessCheckCheckInterval resource from its proto representation.
func ProtoToAppengineVersionReadinessCheckCheckInterval(p *appenginepb.AppengineVersionReadinessCheckCheckInterval) *appengine.VersionReadinessCheckCheckInterval {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionReadinessCheckCheckInterval{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionReadinessCheckTimeout converts a VersionReadinessCheckTimeout resource from its proto representation.
func ProtoToAppengineVersionReadinessCheckTimeout(p *appenginepb.AppengineVersionReadinessCheckTimeout) *appengine.VersionReadinessCheckTimeout {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionReadinessCheckTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionReadinessCheckAppStartTimeout converts a VersionReadinessCheckAppStartTimeout resource from its proto representation.
func ProtoToAppengineVersionReadinessCheckAppStartTimeout(p *appenginepb.AppengineVersionReadinessCheckAppStartTimeout) *appengine.VersionReadinessCheckAppStartTimeout {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionReadinessCheckAppStartTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionLivenessCheck converts a VersionLivenessCheck resource from its proto representation.
func ProtoToAppengineVersionLivenessCheck(p *appenginepb.AppengineVersionLivenessCheck) *appengine.VersionLivenessCheck {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionLivenessCheck{
		Path:             dcl.StringOrNil(p.Path),
		Host:             dcl.StringOrNil(p.Host),
		FailureThreshold: dcl.Int64OrNil(p.FailureThreshold),
		SuccessThreshold: dcl.Int64OrNil(p.SuccessThreshold),
		CheckInterval:    ProtoToAppengineVersionLivenessCheckCheckInterval(p.GetCheckInterval()),
		Timeout:          ProtoToAppengineVersionLivenessCheckTimeout(p.GetTimeout()),
		InitialDelay:     ProtoToAppengineVersionLivenessCheckInitialDelay(p.GetInitialDelay()),
	}
	return obj
}

// ProtoToVersionLivenessCheckCheckInterval converts a VersionLivenessCheckCheckInterval resource from its proto representation.
func ProtoToAppengineVersionLivenessCheckCheckInterval(p *appenginepb.AppengineVersionLivenessCheckCheckInterval) *appengine.VersionLivenessCheckCheckInterval {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionLivenessCheckCheckInterval{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionLivenessCheckTimeout converts a VersionLivenessCheckTimeout resource from its proto representation.
func ProtoToAppengineVersionLivenessCheckTimeout(p *appenginepb.AppengineVersionLivenessCheckTimeout) *appengine.VersionLivenessCheckTimeout {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionLivenessCheckTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionLivenessCheckInitialDelay converts a VersionLivenessCheckInitialDelay resource from its proto representation.
func ProtoToAppengineVersionLivenessCheckInitialDelay(p *appenginepb.AppengineVersionLivenessCheckInitialDelay) *appengine.VersionLivenessCheckInitialDelay {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionLivenessCheckInitialDelay{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToVersionServiceAuthSpec converts a VersionServiceAuthSpec resource from its proto representation.
func ProtoToAppengineVersionServiceAuthSpec(p *appenginepb.AppengineVersionServiceAuthSpec) *appengine.VersionServiceAuthSpec {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionServiceAuthSpec{
		IamServiceName:       dcl.StringOrNil(p.IamServiceName),
		IamResourceName:      dcl.StringOrNil(p.IamResourceName),
		IamPolicyId:          dcl.StringOrNil(p.IamPolicyId),
		IamPolicyType:        dcl.StringOrNil(p.IamPolicyType),
		IamPermission:        dcl.StringOrNil(p.IamPermission),
		AcceptGcloudClientId: dcl.Bool(p.AcceptGcloudClientId),
		Clear:                dcl.Bool(p.Clear),
	}
	for _, r := range p.GetAudiences() {
		obj.Audiences = append(obj.Audiences, r)
	}
	return obj
}

// ProtoToVersionServiceCorsSpec converts a VersionServiceCorsSpec resource from its proto representation.
func ProtoToAppengineVersionServiceCorsSpec(p *appenginepb.AppengineVersionServiceCorsSpec) *appengine.VersionServiceCorsSpec {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionServiceCorsSpec{
		AllowCredential: dcl.Bool(p.AllowCredential),
		MaxAgeSeconds:   dcl.Int64OrNil(p.MaxAgeSeconds),
	}
	for _, r := range p.GetOrigin() {
		obj.Origin = append(obj.Origin, r)
	}
	for _, r := range p.GetMethod() {
		obj.Method = append(obj.Method, r)
	}
	for _, r := range p.GetHeader() {
		obj.Header = append(obj.Header, r)
	}
	for _, r := range p.GetExposedHeader() {
		obj.ExposedHeader = append(obj.ExposedHeader, r)
	}
	return obj
}

// ProtoToVersionEntrypoint converts a VersionEntrypoint resource from its proto representation.
func ProtoToAppengineVersionEntrypoint(p *appenginepb.AppengineVersionEntrypoint) *appengine.VersionEntrypoint {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionEntrypoint{
		Shell: dcl.StringOrNil(p.Shell),
	}
	return obj
}

// ProtoToVersionVPCAccessConnector converts a VersionVPCAccessConnector resource from its proto representation.
func ProtoToAppengineVersionVPCAccessConnector(p *appenginepb.AppengineVersionVPCAccessConnector) *appengine.VersionVPCAccessConnector {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionVPCAccessConnector{
		Name:          dcl.StringOrNil(p.Name),
		EgressSetting: ProtoToAppengineVersionVPCAccessConnectorEgressSettingEnum(p.GetEgressSetting()),
	}
	return obj
}

// ProtoToVersionNetworkSettings converts a VersionNetworkSettings resource from its proto representation.
func ProtoToAppengineVersionNetworkSettings(p *appenginepb.AppengineVersionNetworkSettings) *appengine.VersionNetworkSettings {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionNetworkSettings{
		IngressTrafficAllowed: ProtoToAppengineVersionNetworkSettingsIngressTrafficAllowedEnum(p.GetIngressTrafficAllowed()),
	}
	return obj
}

// ProtoToVersionInstanceSpec converts a VersionInstanceSpec resource from its proto representation.
func ProtoToAppengineVersionInstanceSpec(p *appenginepb.AppengineVersionInstanceSpec) *appengine.VersionInstanceSpec {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionInstanceSpec{}
	for _, r := range p.GetSandboxes() {
		obj.Sandboxes = append(obj.Sandboxes, *ProtoToAppengineVersionInstanceSpecSandboxes(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToAppengineVersionInstanceSpecPorts(r))
	}
	return obj
}

// ProtoToVersionInstanceSpecSandboxes converts a VersionInstanceSpecSandboxes resource from its proto representation.
func ProtoToAppengineVersionInstanceSpecSandboxes(p *appenginepb.AppengineVersionInstanceSpecSandboxes) *appengine.VersionInstanceSpecSandboxes {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionInstanceSpecSandboxes{
		Name: dcl.StringOrNil(p.Name),
	}
	for _, r := range p.GetContainers() {
		obj.Containers = append(obj.Containers, *ProtoToAppengineVersionInstanceSpecSandboxesContainers(r))
	}
	return obj
}

// ProtoToVersionInstanceSpecSandboxesContainers converts a VersionInstanceSpecSandboxesContainers resource from its proto representation.
func ProtoToAppengineVersionInstanceSpecSandboxesContainers(p *appenginepb.AppengineVersionInstanceSpecSandboxesContainers) *appengine.VersionInstanceSpecSandboxesContainers {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionInstanceSpecSandboxesContainers{}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToVersionInstanceSpecPorts converts a VersionInstanceSpecPorts resource from its proto representation.
func ProtoToAppengineVersionInstanceSpecPorts(p *appenginepb.AppengineVersionInstanceSpecPorts) *appengine.VersionInstanceSpecPorts {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionInstanceSpecPorts{
		Name:      dcl.StringOrNil(p.Name),
		Sandbox:   dcl.StringOrNil(p.Sandbox),
		Port:      dcl.Int64OrNil(p.Port),
		Protocol:  ProtoToAppengineVersionInstanceSpecPortsProtocolEnum(p.GetProtocol()),
		IsDefault: dcl.Bool(p.IsDefault),
	}
	return obj
}

// ProtoToVersion converts a Version resource from its proto representation.
func ProtoToVersion(p *appenginepb.AppengineVersion) *appengine.Version {
	obj := &appengine.Version{
		ConsumerName:              dcl.StringOrNil(p.ConsumerName),
		Name:                      dcl.StringOrNil(p.Name),
		AutomaticScaling:          ProtoToAppengineVersionAutomaticScaling(p.GetAutomaticScaling()),
		BasicScaling:              ProtoToAppengineVersionBasicScaling(p.GetBasicScaling()),
		ManualScaling:             ProtoToAppengineVersionManualScaling(p.GetManualScaling()),
		JobScaling:                ProtoToAppengineVersionJobScaling(p.GetJobScaling()),
		PoolScaling:               ProtoToAppengineVersionPoolScaling(p.GetPoolScaling()),
		InstanceClass:             dcl.StringOrNil(p.InstanceClass),
		Network:                   ProtoToAppengineVersionNetwork(p.GetNetwork()),
		Resources:                 ProtoToAppengineVersionResources(p.GetResources()),
		Runtime:                   dcl.StringOrNil(p.Runtime),
		RuntimeChannel:            dcl.StringOrNil(p.RuntimeChannel),
		Threadsafe:                dcl.Bool(p.Threadsafe),
		Vm:                        dcl.Bool(p.Vm),
		Env:                       dcl.StringOrNil(p.Env),
		ServingStatus:             ProtoToAppengineVersionServingStatusEnum(p.GetServingStatus()),
		CreatedBy:                 dcl.StringOrNil(p.CreatedBy),
		CreateTime:                dcl.StringOrNil(p.GetCreateTime()),
		DiskUsageBytes:            dcl.Int64OrNil(p.DiskUsageBytes),
		RuntimeApiVersion:         dcl.StringOrNil(p.RuntimeApiVersion),
		RuntimeMainExecutablePath: dcl.StringOrNil(p.RuntimeMainExecutablePath),
		ApiConfig:                 ProtoToAppengineVersionApiConfig(p.GetApiConfig()),
		DefaultExpiration:         ProtoToAppengineVersionDefaultExpiration(p.GetDefaultExpiration()),
		Deployment:                ProtoToAppengineVersionDeployment(p.GetDeployment()),
		HealthCheck:               ProtoToAppengineVersionHealthCheck(p.GetHealthCheck()),
		ReadinessCheck:            ProtoToAppengineVersionReadinessCheck(p.GetReadinessCheck()),
		LivenessCheck:             ProtoToAppengineVersionLivenessCheck(p.GetLivenessCheck()),
		NobuildFilesRegex:         dcl.StringOrNil(p.NobuildFilesRegex),
		VersionUrl:                dcl.StringOrNil(p.VersionUrl),
		ServiceAuthSpec:           ProtoToAppengineVersionServiceAuthSpec(p.GetServiceAuthSpec()),
		ServiceCorsSpec:           ProtoToAppengineVersionServiceCorsSpec(p.GetServiceCorsSpec()),
		RouteHash:                 dcl.StringOrNil(p.RouteHash),
		Entrypoint:                ProtoToAppengineVersionEntrypoint(p.GetEntrypoint()),
		VPCAccessConnector:        ProtoToAppengineVersionVPCAccessConnector(p.GetVpcAccessConnector()),
		NetworkSettings:           ProtoToAppengineVersionNetworkSettings(p.GetNetworkSettings()),
		InstanceSpec:              ProtoToAppengineVersionInstanceSpec(p.GetInstanceSpec()),
		App:                       dcl.StringOrNil(p.App),
		Service:                   dcl.StringOrNil(p.Service),
	}
	for _, r := range p.GetInboundServices() {
		obj.InboundServices = append(obj.InboundServices, *ProtoToAppengineVersionInboundServicesEnum(r))
	}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, r)
	}
	for _, r := range p.GetBetaSettings() {
		obj.BetaSettings = append(obj.BetaSettings, *ProtoToAppengineVersionBetaSettings(r))
	}
	for _, r := range p.GetHandlers() {
		obj.Handlers = append(obj.Handlers, *ProtoToAppengineVersionHandlers(r))
	}
	for _, r := range p.GetErrorHandlers() {
		obj.ErrorHandlers = append(obj.ErrorHandlers, *ProtoToAppengineVersionErrorHandlers(r))
	}
	for _, r := range p.GetLibraries() {
		obj.Libraries = append(obj.Libraries, *ProtoToAppengineVersionLibraries(r))
	}
	for _, r := range p.GetEnvVariables() {
		obj.EnvVariables = append(obj.EnvVariables, *ProtoToAppengineVersionEnvVariables(r))
	}
	for _, r := range p.GetBuildEnvVariables() {
		obj.BuildEnvVariables = append(obj.BuildEnvVariables, *ProtoToAppengineVersionBuildEnvVariables(r))
	}
	return obj
}

// VersionInboundServicesEnumToProto converts a VersionInboundServicesEnum enum to its proto representation.
func AppengineVersionInboundServicesEnumToProto(e *appengine.VersionInboundServicesEnum) appenginepb.AppengineVersionInboundServicesEnum {
	if e == nil {
		return appenginepb.AppengineVersionInboundServicesEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionInboundServicesEnum_value["VersionInboundServicesEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionInboundServicesEnum(v)
	}
	return appenginepb.AppengineVersionInboundServicesEnum(0)
}

// VersionServingStatusEnumToProto converts a VersionServingStatusEnum enum to its proto representation.
func AppengineVersionServingStatusEnumToProto(e *appengine.VersionServingStatusEnum) appenginepb.AppengineVersionServingStatusEnum {
	if e == nil {
		return appenginepb.AppengineVersionServingStatusEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionServingStatusEnum_value["VersionServingStatusEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionServingStatusEnum(v)
	}
	return appenginepb.AppengineVersionServingStatusEnum(0)
}

// VersionHandlersSecurityLevelEnumToProto converts a VersionHandlersSecurityLevelEnum enum to its proto representation.
func AppengineVersionHandlersSecurityLevelEnumToProto(e *appengine.VersionHandlersSecurityLevelEnum) appenginepb.AppengineVersionHandlersSecurityLevelEnum {
	if e == nil {
		return appenginepb.AppengineVersionHandlersSecurityLevelEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionHandlersSecurityLevelEnum_value["VersionHandlersSecurityLevelEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionHandlersSecurityLevelEnum(v)
	}
	return appenginepb.AppengineVersionHandlersSecurityLevelEnum(0)
}

// VersionHandlersLoginEnumToProto converts a VersionHandlersLoginEnum enum to its proto representation.
func AppengineVersionHandlersLoginEnumToProto(e *appengine.VersionHandlersLoginEnum) appenginepb.AppengineVersionHandlersLoginEnum {
	if e == nil {
		return appenginepb.AppengineVersionHandlersLoginEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionHandlersLoginEnum_value["VersionHandlersLoginEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionHandlersLoginEnum(v)
	}
	return appenginepb.AppengineVersionHandlersLoginEnum(0)
}

// VersionHandlersAuthFailActionEnumToProto converts a VersionHandlersAuthFailActionEnum enum to its proto representation.
func AppengineVersionHandlersAuthFailActionEnumToProto(e *appengine.VersionHandlersAuthFailActionEnum) appenginepb.AppengineVersionHandlersAuthFailActionEnum {
	if e == nil {
		return appenginepb.AppengineVersionHandlersAuthFailActionEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionHandlersAuthFailActionEnum_value["VersionHandlersAuthFailActionEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionHandlersAuthFailActionEnum(v)
	}
	return appenginepb.AppengineVersionHandlersAuthFailActionEnum(0)
}

// VersionHandlersRedirectHttpResponseCodeEnumToProto converts a VersionHandlersRedirectHttpResponseCodeEnum enum to its proto representation.
func AppengineVersionHandlersRedirectHttpResponseCodeEnumToProto(e *appengine.VersionHandlersRedirectHttpResponseCodeEnum) appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum {
	if e == nil {
		return appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum_value["VersionHandlersRedirectHttpResponseCodeEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum(v)
	}
	return appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum(0)
}

// VersionErrorHandlersErrorCodeEnumToProto converts a VersionErrorHandlersErrorCodeEnum enum to its proto representation.
func AppengineVersionErrorHandlersErrorCodeEnumToProto(e *appengine.VersionErrorHandlersErrorCodeEnum) appenginepb.AppengineVersionErrorHandlersErrorCodeEnum {
	if e == nil {
		return appenginepb.AppengineVersionErrorHandlersErrorCodeEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionErrorHandlersErrorCodeEnum_value["VersionErrorHandlersErrorCodeEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionErrorHandlersErrorCodeEnum(v)
	}
	return appenginepb.AppengineVersionErrorHandlersErrorCodeEnum(0)
}

// VersionApiConfigAuthFailActionEnumToProto converts a VersionApiConfigAuthFailActionEnum enum to its proto representation.
func AppengineVersionApiConfigAuthFailActionEnumToProto(e *appengine.VersionApiConfigAuthFailActionEnum) appenginepb.AppengineVersionApiConfigAuthFailActionEnum {
	if e == nil {
		return appenginepb.AppengineVersionApiConfigAuthFailActionEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionApiConfigAuthFailActionEnum_value["VersionApiConfigAuthFailActionEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionApiConfigAuthFailActionEnum(v)
	}
	return appenginepb.AppengineVersionApiConfigAuthFailActionEnum(0)
}

// VersionApiConfigLoginEnumToProto converts a VersionApiConfigLoginEnum enum to its proto representation.
func AppengineVersionApiConfigLoginEnumToProto(e *appengine.VersionApiConfigLoginEnum) appenginepb.AppengineVersionApiConfigLoginEnum {
	if e == nil {
		return appenginepb.AppengineVersionApiConfigLoginEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionApiConfigLoginEnum_value["VersionApiConfigLoginEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionApiConfigLoginEnum(v)
	}
	return appenginepb.AppengineVersionApiConfigLoginEnum(0)
}

// VersionApiConfigSecurityLevelEnumToProto converts a VersionApiConfigSecurityLevelEnum enum to its proto representation.
func AppengineVersionApiConfigSecurityLevelEnumToProto(e *appengine.VersionApiConfigSecurityLevelEnum) appenginepb.AppengineVersionApiConfigSecurityLevelEnum {
	if e == nil {
		return appenginepb.AppengineVersionApiConfigSecurityLevelEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionApiConfigSecurityLevelEnum_value["VersionApiConfigSecurityLevelEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionApiConfigSecurityLevelEnum(v)
	}
	return appenginepb.AppengineVersionApiConfigSecurityLevelEnum(0)
}

// VersionVPCAccessConnectorEgressSettingEnumToProto converts a VersionVPCAccessConnectorEgressSettingEnum enum to its proto representation.
func AppengineVersionVPCAccessConnectorEgressSettingEnumToProto(e *appengine.VersionVPCAccessConnectorEgressSettingEnum) appenginepb.AppengineVersionVPCAccessConnectorEgressSettingEnum {
	if e == nil {
		return appenginepb.AppengineVersionVPCAccessConnectorEgressSettingEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionVPCAccessConnectorEgressSettingEnum_value["VersionVPCAccessConnectorEgressSettingEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionVPCAccessConnectorEgressSettingEnum(v)
	}
	return appenginepb.AppengineVersionVPCAccessConnectorEgressSettingEnum(0)
}

// VersionNetworkSettingsIngressTrafficAllowedEnumToProto converts a VersionNetworkSettingsIngressTrafficAllowedEnum enum to its proto representation.
func AppengineVersionNetworkSettingsIngressTrafficAllowedEnumToProto(e *appengine.VersionNetworkSettingsIngressTrafficAllowedEnum) appenginepb.AppengineVersionNetworkSettingsIngressTrafficAllowedEnum {
	if e == nil {
		return appenginepb.AppengineVersionNetworkSettingsIngressTrafficAllowedEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionNetworkSettingsIngressTrafficAllowedEnum_value["VersionNetworkSettingsIngressTrafficAllowedEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionNetworkSettingsIngressTrafficAllowedEnum(v)
	}
	return appenginepb.AppengineVersionNetworkSettingsIngressTrafficAllowedEnum(0)
}

// VersionInstanceSpecPortsProtocolEnumToProto converts a VersionInstanceSpecPortsProtocolEnum enum to its proto representation.
func AppengineVersionInstanceSpecPortsProtocolEnumToProto(e *appengine.VersionInstanceSpecPortsProtocolEnum) appenginepb.AppengineVersionInstanceSpecPortsProtocolEnum {
	if e == nil {
		return appenginepb.AppengineVersionInstanceSpecPortsProtocolEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionInstanceSpecPortsProtocolEnum_value["VersionInstanceSpecPortsProtocolEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionInstanceSpecPortsProtocolEnum(v)
	}
	return appenginepb.AppengineVersionInstanceSpecPortsProtocolEnum(0)
}

// VersionAutomaticScalingToProto converts a VersionAutomaticScaling resource to its proto representation.
func AppengineVersionAutomaticScalingToProto(o *appengine.VersionAutomaticScaling) *appenginepb.AppengineVersionAutomaticScaling {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScaling{
		CoolDownPeriod:            AppengineVersionAutomaticScalingCoolDownPeriodToProto(o.CoolDownPeriod),
		CpuUtilization:            AppengineVersionAutomaticScalingCpuUtilizationToProto(o.CpuUtilization),
		MaxConcurrentRequests:     dcl.ValueOrEmptyInt64(o.MaxConcurrentRequests),
		MaxIdleInstances:          dcl.ValueOrEmptyInt64(o.MaxIdleInstances),
		MaxTotalInstances:         dcl.ValueOrEmptyInt64(o.MaxTotalInstances),
		MaxPendingLatency:         AppengineVersionAutomaticScalingMaxPendingLatencyToProto(o.MaxPendingLatency),
		MinIdleInstances:          dcl.ValueOrEmptyInt64(o.MinIdleInstances),
		MinTotalInstances:         dcl.ValueOrEmptyInt64(o.MinTotalInstances),
		MinPendingLatency:         AppengineVersionAutomaticScalingMinPendingLatencyToProto(o.MinPendingLatency),
		RequestUtilization:        AppengineVersionAutomaticScalingRequestUtilizationToProto(o.RequestUtilization),
		DiskUtilization:           AppengineVersionAutomaticScalingDiskUtilizationToProto(o.DiskUtilization),
		NetworkUtilization:        AppengineVersionAutomaticScalingNetworkUtilizationToProto(o.NetworkUtilization),
		StandardSchedulerSettings: AppengineVersionAutomaticScalingStandardSchedulerSettingsToProto(o.StandardSchedulerSettings),
	}
	return p
}

// VersionAutomaticScalingCoolDownPeriodToProto converts a VersionAutomaticScalingCoolDownPeriod resource to its proto representation.
func AppengineVersionAutomaticScalingCoolDownPeriodToProto(o *appengine.VersionAutomaticScalingCoolDownPeriod) *appenginepb.AppengineVersionAutomaticScalingCoolDownPeriod {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingCoolDownPeriod{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionAutomaticScalingCpuUtilizationToProto converts a VersionAutomaticScalingCpuUtilization resource to its proto representation.
func AppengineVersionAutomaticScalingCpuUtilizationToProto(o *appengine.VersionAutomaticScalingCpuUtilization) *appenginepb.AppengineVersionAutomaticScalingCpuUtilization {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingCpuUtilization{
		AggregationWindowLength: AppengineVersionAutomaticScalingCpuUtilizationAggregationWindowLengthToProto(o.AggregationWindowLength),
		TargetUtilization:       dcl.ValueOrEmptyDouble(o.TargetUtilization),
	}
	return p
}

// VersionAutomaticScalingCpuUtilizationAggregationWindowLengthToProto converts a VersionAutomaticScalingCpuUtilizationAggregationWindowLength resource to its proto representation.
func AppengineVersionAutomaticScalingCpuUtilizationAggregationWindowLengthToProto(o *appengine.VersionAutomaticScalingCpuUtilizationAggregationWindowLength) *appenginepb.AppengineVersionAutomaticScalingCpuUtilizationAggregationWindowLength {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingCpuUtilizationAggregationWindowLength{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionAutomaticScalingMaxPendingLatencyToProto converts a VersionAutomaticScalingMaxPendingLatency resource to its proto representation.
func AppengineVersionAutomaticScalingMaxPendingLatencyToProto(o *appengine.VersionAutomaticScalingMaxPendingLatency) *appenginepb.AppengineVersionAutomaticScalingMaxPendingLatency {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingMaxPendingLatency{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionAutomaticScalingMinPendingLatencyToProto converts a VersionAutomaticScalingMinPendingLatency resource to its proto representation.
func AppengineVersionAutomaticScalingMinPendingLatencyToProto(o *appengine.VersionAutomaticScalingMinPendingLatency) *appenginepb.AppengineVersionAutomaticScalingMinPendingLatency {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingMinPendingLatency{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionAutomaticScalingRequestUtilizationToProto converts a VersionAutomaticScalingRequestUtilization resource to its proto representation.
func AppengineVersionAutomaticScalingRequestUtilizationToProto(o *appengine.VersionAutomaticScalingRequestUtilization) *appenginepb.AppengineVersionAutomaticScalingRequestUtilization {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingRequestUtilization{
		TargetRequestCountPerSecond: dcl.ValueOrEmptyInt64(o.TargetRequestCountPerSecond),
		TargetConcurrentRequests:    dcl.ValueOrEmptyInt64(o.TargetConcurrentRequests),
	}
	return p
}

// VersionAutomaticScalingDiskUtilizationToProto converts a VersionAutomaticScalingDiskUtilization resource to its proto representation.
func AppengineVersionAutomaticScalingDiskUtilizationToProto(o *appengine.VersionAutomaticScalingDiskUtilization) *appenginepb.AppengineVersionAutomaticScalingDiskUtilization {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingDiskUtilization{
		TargetWriteBytesPerSecond: dcl.ValueOrEmptyInt64(o.TargetWriteBytesPerSecond),
		TargetWriteOpsPerSecond:   dcl.ValueOrEmptyInt64(o.TargetWriteOpsPerSecond),
		TargetReadBytesPerSecond:  dcl.ValueOrEmptyInt64(o.TargetReadBytesPerSecond),
		TargetReadOpsPerSecond:    dcl.ValueOrEmptyInt64(o.TargetReadOpsPerSecond),
	}
	return p
}

// VersionAutomaticScalingNetworkUtilizationToProto converts a VersionAutomaticScalingNetworkUtilization resource to its proto representation.
func AppengineVersionAutomaticScalingNetworkUtilizationToProto(o *appengine.VersionAutomaticScalingNetworkUtilization) *appenginepb.AppengineVersionAutomaticScalingNetworkUtilization {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingNetworkUtilization{
		TargetSentBytesPerSecond:       dcl.ValueOrEmptyInt64(o.TargetSentBytesPerSecond),
		TargetSentPacketsPerSecond:     dcl.ValueOrEmptyInt64(o.TargetSentPacketsPerSecond),
		TargetReceivedBytesPerSecond:   dcl.ValueOrEmptyInt64(o.TargetReceivedBytesPerSecond),
		TargetReceivedPacketsPerSecond: dcl.ValueOrEmptyInt64(o.TargetReceivedPacketsPerSecond),
	}
	return p
}

// VersionAutomaticScalingStandardSchedulerSettingsToProto converts a VersionAutomaticScalingStandardSchedulerSettings resource to its proto representation.
func AppengineVersionAutomaticScalingStandardSchedulerSettingsToProto(o *appengine.VersionAutomaticScalingStandardSchedulerSettings) *appenginepb.AppengineVersionAutomaticScalingStandardSchedulerSettings {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingStandardSchedulerSettings{
		TargetCpuUtilization:        dcl.ValueOrEmptyDouble(o.TargetCpuUtilization),
		TargetThroughputUtilization: dcl.ValueOrEmptyDouble(o.TargetThroughputUtilization),
		MinInstances:                dcl.ValueOrEmptyInt64(o.MinInstances),
		MaxInstances:                dcl.ValueOrEmptyInt64(o.MaxInstances),
	}
	return p
}

// VersionBasicScalingToProto converts a VersionBasicScaling resource to its proto representation.
func AppengineVersionBasicScalingToProto(o *appengine.VersionBasicScaling) *appenginepb.AppengineVersionBasicScaling {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionBasicScaling{
		IdleTimeout:  AppengineVersionBasicScalingIdleTimeoutToProto(o.IdleTimeout),
		MaxInstances: dcl.ValueOrEmptyInt64(o.MaxInstances),
	}
	return p
}

// VersionBasicScalingIdleTimeoutToProto converts a VersionBasicScalingIdleTimeout resource to its proto representation.
func AppengineVersionBasicScalingIdleTimeoutToProto(o *appengine.VersionBasicScalingIdleTimeout) *appenginepb.AppengineVersionBasicScalingIdleTimeout {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionBasicScalingIdleTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionManualScalingToProto converts a VersionManualScaling resource to its proto representation.
func AppengineVersionManualScalingToProto(o *appengine.VersionManualScaling) *appenginepb.AppengineVersionManualScaling {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionManualScaling{
		Instances: dcl.ValueOrEmptyInt64(o.Instances),
	}
	return p
}

// VersionJobScalingToProto converts a VersionJobScaling resource to its proto representation.
func AppengineVersionJobScalingToProto(o *appengine.VersionJobScaling) *appenginepb.AppengineVersionJobScaling {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionJobScaling{
		Completions:               dcl.ValueOrEmptyInt64(o.Completions),
		Parallelism:               dcl.ValueOrEmptyInt64(o.Parallelism),
		JobDeadline:               AppengineVersionJobScalingJobDeadlineToProto(o.JobDeadline),
		InstanceRetries:           dcl.ValueOrEmptyInt64(o.InstanceRetries),
		InstanceDeadline:          AppengineVersionJobScalingInstanceDeadlineToProto(o.InstanceDeadline),
		InstanceTerminationWindow: AppengineVersionJobScalingInstanceTerminationWindowToProto(o.InstanceTerminationWindow),
	}
	return p
}

// VersionJobScalingJobDeadlineToProto converts a VersionJobScalingJobDeadline resource to its proto representation.
func AppengineVersionJobScalingJobDeadlineToProto(o *appengine.VersionJobScalingJobDeadline) *appenginepb.AppengineVersionJobScalingJobDeadline {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionJobScalingJobDeadline{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionJobScalingInstanceDeadlineToProto converts a VersionJobScalingInstanceDeadline resource to its proto representation.
func AppengineVersionJobScalingInstanceDeadlineToProto(o *appengine.VersionJobScalingInstanceDeadline) *appenginepb.AppengineVersionJobScalingInstanceDeadline {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionJobScalingInstanceDeadline{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionJobScalingInstanceTerminationWindowToProto converts a VersionJobScalingInstanceTerminationWindow resource to its proto representation.
func AppengineVersionJobScalingInstanceTerminationWindowToProto(o *appengine.VersionJobScalingInstanceTerminationWindow) *appenginepb.AppengineVersionJobScalingInstanceTerminationWindow {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionJobScalingInstanceTerminationWindow{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionPoolScalingToProto converts a VersionPoolScaling resource to its proto representation.
func AppengineVersionPoolScalingToProto(o *appengine.VersionPoolScaling) *appenginepb.AppengineVersionPoolScaling {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionPoolScaling{
		Replicas:       dcl.ValueOrEmptyInt64(o.Replicas),
		MaxUnavailable: dcl.ValueOrEmptyInt64(o.MaxUnavailable),
		MaxSurge:       dcl.ValueOrEmptyInt64(o.MaxSurge),
	}
	return p
}

// VersionNetworkToProto converts a VersionNetwork resource to its proto representation.
func AppengineVersionNetworkToProto(o *appengine.VersionNetwork) *appenginepb.AppengineVersionNetwork {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionNetwork{
		InstanceTag:     dcl.ValueOrEmptyString(o.InstanceTag),
		Name:            dcl.ValueOrEmptyString(o.Name),
		SubnetworkName:  dcl.ValueOrEmptyString(o.SubnetworkName),
		SessionAffinity: dcl.ValueOrEmptyBool(o.SessionAffinity),
	}
	for _, r := range o.ForwardedPorts {
		p.ForwardedPorts = append(p.ForwardedPorts, r)
	}
	return p
}

// VersionResourcesToProto converts a VersionResources resource to its proto representation.
func AppengineVersionResourcesToProto(o *appengine.VersionResources) *appenginepb.AppengineVersionResources {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionResources{
		Cpu:             dcl.ValueOrEmptyDouble(o.Cpu),
		DiskGb:          dcl.ValueOrEmptyDouble(o.DiskGb),
		MemoryGb:        dcl.ValueOrEmptyDouble(o.MemoryGb),
		KmsKeyReference: dcl.ValueOrEmptyString(o.KmsKeyReference),
	}
	for _, r := range o.Volumes {
		p.Volumes = append(p.Volumes, AppengineVersionResourcesVolumesToProto(&r))
	}
	return p
}

// VersionResourcesVolumesToProto converts a VersionResourcesVolumes resource to its proto representation.
func AppengineVersionResourcesVolumesToProto(o *appengine.VersionResourcesVolumes) *appenginepb.AppengineVersionResourcesVolumes {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionResourcesVolumes{
		Name:       dcl.ValueOrEmptyString(o.Name),
		VolumeType: dcl.ValueOrEmptyString(o.VolumeType),
		SizeGb:     dcl.ValueOrEmptyDouble(o.SizeGb),
	}
	return p
}

// VersionBetaSettingsToProto converts a VersionBetaSettings resource to its proto representation.
func AppengineVersionBetaSettingsToProto(o *appengine.VersionBetaSettings) *appenginepb.AppengineVersionBetaSettings {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionBetaSettings{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// VersionHandlersToProto converts a VersionHandlers resource to its proto representation.
func AppengineVersionHandlersToProto(o *appengine.VersionHandlers) *appenginepb.AppengineVersionHandlers {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlers{
		UrlRegex:                 dcl.ValueOrEmptyString(o.UrlRegex),
		StaticFiles:              AppengineVersionHandlersStaticFilesToProto(o.StaticFiles),
		Script:                   AppengineVersionHandlersScriptToProto(o.Script),
		ApiEndpoint:              AppengineVersionHandlersApiEndpointToProto(o.ApiEndpoint),
		SecurityLevel:            AppengineVersionHandlersSecurityLevelEnumToProto(o.SecurityLevel),
		Login:                    AppengineVersionHandlersLoginEnumToProto(o.Login),
		AuthFailAction:           AppengineVersionHandlersAuthFailActionEnumToProto(o.AuthFailAction),
		RedirectHttpResponseCode: AppengineVersionHandlersRedirectHttpResponseCodeEnumToProto(o.RedirectHttpResponseCode),
	}
	return p
}

// VersionHandlersStaticFilesToProto converts a VersionHandlersStaticFiles resource to its proto representation.
func AppengineVersionHandlersStaticFilesToProto(o *appengine.VersionHandlersStaticFiles) *appenginepb.AppengineVersionHandlersStaticFiles {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlersStaticFiles{
		Path:                dcl.ValueOrEmptyString(o.Path),
		UploadPathRegex:     dcl.ValueOrEmptyString(o.UploadPathRegex),
		MimeType:            dcl.ValueOrEmptyString(o.MimeType),
		Expiration:          AppengineVersionHandlersStaticFilesExpirationToProto(o.Expiration),
		RequireMatchingFile: dcl.ValueOrEmptyBool(o.RequireMatchingFile),
		ApplicationReadable: dcl.ValueOrEmptyBool(o.ApplicationReadable),
	}
	for _, r := range o.HttpHeaders {
		p.HttpHeaders = append(p.HttpHeaders, AppengineVersionHandlersStaticFilesHttpHeadersToProto(&r))
	}
	return p
}

// VersionHandlersStaticFilesHttpHeadersToProto converts a VersionHandlersStaticFilesHttpHeaders resource to its proto representation.
func AppengineVersionHandlersStaticFilesHttpHeadersToProto(o *appengine.VersionHandlersStaticFilesHttpHeaders) *appenginepb.AppengineVersionHandlersStaticFilesHttpHeaders {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlersStaticFilesHttpHeaders{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// VersionHandlersStaticFilesExpirationToProto converts a VersionHandlersStaticFilesExpiration resource to its proto representation.
func AppengineVersionHandlersStaticFilesExpirationToProto(o *appengine.VersionHandlersStaticFilesExpiration) *appenginepb.AppengineVersionHandlersStaticFilesExpiration {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlersStaticFilesExpiration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionHandlersScriptToProto converts a VersionHandlersScript resource to its proto representation.
func AppengineVersionHandlersScriptToProto(o *appengine.VersionHandlersScript) *appenginepb.AppengineVersionHandlersScript {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlersScript{
		ScriptPath: dcl.ValueOrEmptyString(o.ScriptPath),
	}
	return p
}

// VersionHandlersApiEndpointToProto converts a VersionHandlersApiEndpoint resource to its proto representation.
func AppengineVersionHandlersApiEndpointToProto(o *appengine.VersionHandlersApiEndpoint) *appenginepb.AppengineVersionHandlersApiEndpoint {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlersApiEndpoint{
		ScriptPath: dcl.ValueOrEmptyString(o.ScriptPath),
	}
	return p
}

// VersionErrorHandlersToProto converts a VersionErrorHandlers resource to its proto representation.
func AppengineVersionErrorHandlersToProto(o *appengine.VersionErrorHandlers) *appenginepb.AppengineVersionErrorHandlers {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionErrorHandlers{
		ErrorCode:  AppengineVersionErrorHandlersErrorCodeEnumToProto(o.ErrorCode),
		StaticFile: dcl.ValueOrEmptyString(o.StaticFile),
		MimeType:   dcl.ValueOrEmptyString(o.MimeType),
	}
	return p
}

// VersionLibrariesToProto converts a VersionLibraries resource to its proto representation.
func AppengineVersionLibrariesToProto(o *appengine.VersionLibraries) *appenginepb.AppengineVersionLibraries {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionLibraries{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Version: dcl.ValueOrEmptyString(o.Version),
	}
	return p
}

// VersionApiConfigToProto converts a VersionApiConfig resource to its proto representation.
func AppengineVersionApiConfigToProto(o *appengine.VersionApiConfig) *appenginepb.AppengineVersionApiConfig {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionApiConfig{
		AuthFailAction: AppengineVersionApiConfigAuthFailActionEnumToProto(o.AuthFailAction),
		Login:          AppengineVersionApiConfigLoginEnumToProto(o.Login),
		Script:         dcl.ValueOrEmptyString(o.Script),
		SecurityLevel:  AppengineVersionApiConfigSecurityLevelEnumToProto(o.SecurityLevel),
		Url:            dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// VersionEnvVariablesToProto converts a VersionEnvVariables resource to its proto representation.
func AppengineVersionEnvVariablesToProto(o *appengine.VersionEnvVariables) *appenginepb.AppengineVersionEnvVariables {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionEnvVariables{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// VersionBuildEnvVariablesToProto converts a VersionBuildEnvVariables resource to its proto representation.
func AppengineVersionBuildEnvVariablesToProto(o *appengine.VersionBuildEnvVariables) *appenginepb.AppengineVersionBuildEnvVariables {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionBuildEnvVariables{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// VersionDefaultExpirationToProto converts a VersionDefaultExpiration resource to its proto representation.
func AppengineVersionDefaultExpirationToProto(o *appengine.VersionDefaultExpiration) *appenginepb.AppengineVersionDefaultExpiration {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDefaultExpiration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionDeploymentToProto converts a VersionDeployment resource to its proto representation.
func AppengineVersionDeploymentToProto(o *appengine.VersionDeployment) *appenginepb.AppengineVersionDeployment {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeployment{
		Container:         AppengineVersionDeploymentContainerToProto(o.Container),
		Zip:               AppengineVersionDeploymentZipToProto(o.Zip),
		CloudBuildOptions: AppengineVersionDeploymentCloudBuildOptionsToProto(o.CloudBuildOptions),
	}
	p.Files = make(map[string]*appenginepb.AppengineVersionDeploymentFiles)
	for k, r := range o.Files {
		p.Files[k] = AppengineVersionDeploymentFilesToProto(&r)
	}
	return p
}

// VersionDeploymentFilesToProto converts a VersionDeploymentFiles resource to its proto representation.
func AppengineVersionDeploymentFilesToProto(o *appengine.VersionDeploymentFiles) *appenginepb.AppengineVersionDeploymentFiles {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeploymentFiles{
		SourceUrl: dcl.ValueOrEmptyString(o.SourceUrl),
		Sha1Sum:   dcl.ValueOrEmptyString(o.Sha1Sum),
		MimeType:  dcl.ValueOrEmptyString(o.MimeType),
	}
	return p
}

// VersionDeploymentContainerToProto converts a VersionDeploymentContainer resource to its proto representation.
func AppengineVersionDeploymentContainerToProto(o *appengine.VersionDeploymentContainer) *appenginepb.AppengineVersionDeploymentContainer {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeploymentContainer{
		Image: dcl.ValueOrEmptyString(o.Image),
	}
	return p
}

// VersionDeploymentZipToProto converts a VersionDeploymentZip resource to its proto representation.
func AppengineVersionDeploymentZipToProto(o *appengine.VersionDeploymentZip) *appenginepb.AppengineVersionDeploymentZip {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeploymentZip{
		SourceUrl:  dcl.ValueOrEmptyString(o.SourceUrl),
		FilesCount: dcl.ValueOrEmptyInt64(o.FilesCount),
	}
	return p
}

// VersionDeploymentCloudBuildOptionsToProto converts a VersionDeploymentCloudBuildOptions resource to its proto representation.
func AppengineVersionDeploymentCloudBuildOptionsToProto(o *appengine.VersionDeploymentCloudBuildOptions) *appenginepb.AppengineVersionDeploymentCloudBuildOptions {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeploymentCloudBuildOptions{
		AppYamlPath:       dcl.ValueOrEmptyString(o.AppYamlPath),
		CloudBuildTimeout: AppengineVersionDeploymentCloudBuildOptionsCloudBuildTimeoutToProto(o.CloudBuildTimeout),
	}
	return p
}

// VersionDeploymentCloudBuildOptionsCloudBuildTimeoutToProto converts a VersionDeploymentCloudBuildOptionsCloudBuildTimeout resource to its proto representation.
func AppengineVersionDeploymentCloudBuildOptionsCloudBuildTimeoutToProto(o *appengine.VersionDeploymentCloudBuildOptionsCloudBuildTimeout) *appenginepb.AppengineVersionDeploymentCloudBuildOptionsCloudBuildTimeout {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeploymentCloudBuildOptionsCloudBuildTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionHealthCheckToProto converts a VersionHealthCheck resource to its proto representation.
func AppengineVersionHealthCheckToProto(o *appengine.VersionHealthCheck) *appenginepb.AppengineVersionHealthCheck {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHealthCheck{
		DisableHealthCheck: dcl.ValueOrEmptyBool(o.DisableHealthCheck),
		Host:               dcl.ValueOrEmptyString(o.Host),
		HealthyThreshold:   dcl.ValueOrEmptyInt64(o.HealthyThreshold),
		UnhealthyThreshold: dcl.ValueOrEmptyInt64(o.UnhealthyThreshold),
		RestartThreshold:   dcl.ValueOrEmptyInt64(o.RestartThreshold),
		CheckInterval:      AppengineVersionHealthCheckCheckIntervalToProto(o.CheckInterval),
		Timeout:            AppengineVersionHealthCheckTimeoutToProto(o.Timeout),
	}
	return p
}

// VersionHealthCheckCheckIntervalToProto converts a VersionHealthCheckCheckInterval resource to its proto representation.
func AppengineVersionHealthCheckCheckIntervalToProto(o *appengine.VersionHealthCheckCheckInterval) *appenginepb.AppengineVersionHealthCheckCheckInterval {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHealthCheckCheckInterval{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionHealthCheckTimeoutToProto converts a VersionHealthCheckTimeout resource to its proto representation.
func AppengineVersionHealthCheckTimeoutToProto(o *appengine.VersionHealthCheckTimeout) *appenginepb.AppengineVersionHealthCheckTimeout {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHealthCheckTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionReadinessCheckToProto converts a VersionReadinessCheck resource to its proto representation.
func AppengineVersionReadinessCheckToProto(o *appengine.VersionReadinessCheck) *appenginepb.AppengineVersionReadinessCheck {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionReadinessCheck{
		Path:             dcl.ValueOrEmptyString(o.Path),
		Host:             dcl.ValueOrEmptyString(o.Host),
		FailureThreshold: dcl.ValueOrEmptyInt64(o.FailureThreshold),
		SuccessThreshold: dcl.ValueOrEmptyInt64(o.SuccessThreshold),
		CheckInterval:    AppengineVersionReadinessCheckCheckIntervalToProto(o.CheckInterval),
		Timeout:          AppengineVersionReadinessCheckTimeoutToProto(o.Timeout),
		AppStartTimeout:  AppengineVersionReadinessCheckAppStartTimeoutToProto(o.AppStartTimeout),
	}
	return p
}

// VersionReadinessCheckCheckIntervalToProto converts a VersionReadinessCheckCheckInterval resource to its proto representation.
func AppengineVersionReadinessCheckCheckIntervalToProto(o *appengine.VersionReadinessCheckCheckInterval) *appenginepb.AppengineVersionReadinessCheckCheckInterval {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionReadinessCheckCheckInterval{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionReadinessCheckTimeoutToProto converts a VersionReadinessCheckTimeout resource to its proto representation.
func AppengineVersionReadinessCheckTimeoutToProto(o *appengine.VersionReadinessCheckTimeout) *appenginepb.AppengineVersionReadinessCheckTimeout {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionReadinessCheckTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionReadinessCheckAppStartTimeoutToProto converts a VersionReadinessCheckAppStartTimeout resource to its proto representation.
func AppengineVersionReadinessCheckAppStartTimeoutToProto(o *appengine.VersionReadinessCheckAppStartTimeout) *appenginepb.AppengineVersionReadinessCheckAppStartTimeout {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionReadinessCheckAppStartTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionLivenessCheckToProto converts a VersionLivenessCheck resource to its proto representation.
func AppengineVersionLivenessCheckToProto(o *appengine.VersionLivenessCheck) *appenginepb.AppengineVersionLivenessCheck {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionLivenessCheck{
		Path:             dcl.ValueOrEmptyString(o.Path),
		Host:             dcl.ValueOrEmptyString(o.Host),
		FailureThreshold: dcl.ValueOrEmptyInt64(o.FailureThreshold),
		SuccessThreshold: dcl.ValueOrEmptyInt64(o.SuccessThreshold),
		CheckInterval:    AppengineVersionLivenessCheckCheckIntervalToProto(o.CheckInterval),
		Timeout:          AppengineVersionLivenessCheckTimeoutToProto(o.Timeout),
		InitialDelay:     AppengineVersionLivenessCheckInitialDelayToProto(o.InitialDelay),
	}
	return p
}

// VersionLivenessCheckCheckIntervalToProto converts a VersionLivenessCheckCheckInterval resource to its proto representation.
func AppengineVersionLivenessCheckCheckIntervalToProto(o *appengine.VersionLivenessCheckCheckInterval) *appenginepb.AppengineVersionLivenessCheckCheckInterval {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionLivenessCheckCheckInterval{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionLivenessCheckTimeoutToProto converts a VersionLivenessCheckTimeout resource to its proto representation.
func AppengineVersionLivenessCheckTimeoutToProto(o *appengine.VersionLivenessCheckTimeout) *appenginepb.AppengineVersionLivenessCheckTimeout {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionLivenessCheckTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionLivenessCheckInitialDelayToProto converts a VersionLivenessCheckInitialDelay resource to its proto representation.
func AppengineVersionLivenessCheckInitialDelayToProto(o *appengine.VersionLivenessCheckInitialDelay) *appenginepb.AppengineVersionLivenessCheckInitialDelay {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionLivenessCheckInitialDelay{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// VersionServiceAuthSpecToProto converts a VersionServiceAuthSpec resource to its proto representation.
func AppengineVersionServiceAuthSpecToProto(o *appengine.VersionServiceAuthSpec) *appenginepb.AppengineVersionServiceAuthSpec {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionServiceAuthSpec{
		IamServiceName:       dcl.ValueOrEmptyString(o.IamServiceName),
		IamResourceName:      dcl.ValueOrEmptyString(o.IamResourceName),
		IamPolicyId:          dcl.ValueOrEmptyString(o.IamPolicyId),
		IamPolicyType:        dcl.ValueOrEmptyString(o.IamPolicyType),
		IamPermission:        dcl.ValueOrEmptyString(o.IamPermission),
		AcceptGcloudClientId: dcl.ValueOrEmptyBool(o.AcceptGcloudClientId),
		Clear:                dcl.ValueOrEmptyBool(o.Clear),
	}
	for _, r := range o.Audiences {
		p.Audiences = append(p.Audiences, r)
	}
	return p
}

// VersionServiceCorsSpecToProto converts a VersionServiceCorsSpec resource to its proto representation.
func AppengineVersionServiceCorsSpecToProto(o *appengine.VersionServiceCorsSpec) *appenginepb.AppengineVersionServiceCorsSpec {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionServiceCorsSpec{
		AllowCredential: dcl.ValueOrEmptyBool(o.AllowCredential),
		MaxAgeSeconds:   dcl.ValueOrEmptyInt64(o.MaxAgeSeconds),
	}
	for _, r := range o.Origin {
		p.Origin = append(p.Origin, r)
	}
	for _, r := range o.Method {
		p.Method = append(p.Method, r)
	}
	for _, r := range o.Header {
		p.Header = append(p.Header, r)
	}
	for _, r := range o.ExposedHeader {
		p.ExposedHeader = append(p.ExposedHeader, r)
	}
	return p
}

// VersionEntrypointToProto converts a VersionEntrypoint resource to its proto representation.
func AppengineVersionEntrypointToProto(o *appengine.VersionEntrypoint) *appenginepb.AppengineVersionEntrypoint {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionEntrypoint{
		Shell: dcl.ValueOrEmptyString(o.Shell),
	}
	return p
}

// VersionVPCAccessConnectorToProto converts a VersionVPCAccessConnector resource to its proto representation.
func AppengineVersionVPCAccessConnectorToProto(o *appengine.VersionVPCAccessConnector) *appenginepb.AppengineVersionVPCAccessConnector {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionVPCAccessConnector{
		Name:          dcl.ValueOrEmptyString(o.Name),
		EgressSetting: AppengineVersionVPCAccessConnectorEgressSettingEnumToProto(o.EgressSetting),
	}
	return p
}

// VersionNetworkSettingsToProto converts a VersionNetworkSettings resource to its proto representation.
func AppengineVersionNetworkSettingsToProto(o *appengine.VersionNetworkSettings) *appenginepb.AppengineVersionNetworkSettings {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionNetworkSettings{
		IngressTrafficAllowed: AppengineVersionNetworkSettingsIngressTrafficAllowedEnumToProto(o.IngressTrafficAllowed),
	}
	return p
}

// VersionInstanceSpecToProto converts a VersionInstanceSpec resource to its proto representation.
func AppengineVersionInstanceSpecToProto(o *appengine.VersionInstanceSpec) *appenginepb.AppengineVersionInstanceSpec {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionInstanceSpec{}
	for _, r := range o.Sandboxes {
		p.Sandboxes = append(p.Sandboxes, AppengineVersionInstanceSpecSandboxesToProto(&r))
	}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, AppengineVersionInstanceSpecPortsToProto(&r))
	}
	return p
}

// VersionInstanceSpecSandboxesToProto converts a VersionInstanceSpecSandboxes resource to its proto representation.
func AppengineVersionInstanceSpecSandboxesToProto(o *appengine.VersionInstanceSpecSandboxes) *appenginepb.AppengineVersionInstanceSpecSandboxes {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionInstanceSpecSandboxes{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	for _, r := range o.Containers {
		p.Containers = append(p.Containers, AppengineVersionInstanceSpecSandboxesContainersToProto(&r))
	}
	return p
}

// VersionInstanceSpecSandboxesContainersToProto converts a VersionInstanceSpecSandboxesContainers resource to its proto representation.
func AppengineVersionInstanceSpecSandboxesContainersToProto(o *appengine.VersionInstanceSpecSandboxesContainers) *appenginepb.AppengineVersionInstanceSpecSandboxesContainers {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionInstanceSpecSandboxesContainers{}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	return p
}

// VersionInstanceSpecPortsToProto converts a VersionInstanceSpecPorts resource to its proto representation.
func AppengineVersionInstanceSpecPortsToProto(o *appengine.VersionInstanceSpecPorts) *appenginepb.AppengineVersionInstanceSpecPorts {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionInstanceSpecPorts{
		Name:      dcl.ValueOrEmptyString(o.Name),
		Sandbox:   dcl.ValueOrEmptyString(o.Sandbox),
		Port:      dcl.ValueOrEmptyInt64(o.Port),
		Protocol:  AppengineVersionInstanceSpecPortsProtocolEnumToProto(o.Protocol),
		IsDefault: dcl.ValueOrEmptyBool(o.IsDefault),
	}
	return p
}

// VersionToProto converts a Version resource to its proto representation.
func VersionToProto(resource *appengine.Version) *appenginepb.AppengineVersion {
	p := &appenginepb.AppengineVersion{
		ConsumerName:              dcl.ValueOrEmptyString(resource.ConsumerName),
		Name:                      dcl.ValueOrEmptyString(resource.Name),
		AutomaticScaling:          AppengineVersionAutomaticScalingToProto(resource.AutomaticScaling),
		BasicScaling:              AppengineVersionBasicScalingToProto(resource.BasicScaling),
		ManualScaling:             AppengineVersionManualScalingToProto(resource.ManualScaling),
		JobScaling:                AppengineVersionJobScalingToProto(resource.JobScaling),
		PoolScaling:               AppengineVersionPoolScalingToProto(resource.PoolScaling),
		InstanceClass:             dcl.ValueOrEmptyString(resource.InstanceClass),
		Network:                   AppengineVersionNetworkToProto(resource.Network),
		Resources:                 AppengineVersionResourcesToProto(resource.Resources),
		Runtime:                   dcl.ValueOrEmptyString(resource.Runtime),
		RuntimeChannel:            dcl.ValueOrEmptyString(resource.RuntimeChannel),
		Threadsafe:                dcl.ValueOrEmptyBool(resource.Threadsafe),
		Vm:                        dcl.ValueOrEmptyBool(resource.Vm),
		Env:                       dcl.ValueOrEmptyString(resource.Env),
		ServingStatus:             AppengineVersionServingStatusEnumToProto(resource.ServingStatus),
		CreatedBy:                 dcl.ValueOrEmptyString(resource.CreatedBy),
		CreateTime:                dcl.ValueOrEmptyString(resource.CreateTime),
		DiskUsageBytes:            dcl.ValueOrEmptyInt64(resource.DiskUsageBytes),
		RuntimeApiVersion:         dcl.ValueOrEmptyString(resource.RuntimeApiVersion),
		RuntimeMainExecutablePath: dcl.ValueOrEmptyString(resource.RuntimeMainExecutablePath),
		ApiConfig:                 AppengineVersionApiConfigToProto(resource.ApiConfig),
		DefaultExpiration:         AppengineVersionDefaultExpirationToProto(resource.DefaultExpiration),
		Deployment:                AppengineVersionDeploymentToProto(resource.Deployment),
		HealthCheck:               AppengineVersionHealthCheckToProto(resource.HealthCheck),
		ReadinessCheck:            AppengineVersionReadinessCheckToProto(resource.ReadinessCheck),
		LivenessCheck:             AppengineVersionLivenessCheckToProto(resource.LivenessCheck),
		NobuildFilesRegex:         dcl.ValueOrEmptyString(resource.NobuildFilesRegex),
		VersionUrl:                dcl.ValueOrEmptyString(resource.VersionUrl),
		ServiceAuthSpec:           AppengineVersionServiceAuthSpecToProto(resource.ServiceAuthSpec),
		ServiceCorsSpec:           AppengineVersionServiceCorsSpecToProto(resource.ServiceCorsSpec),
		RouteHash:                 dcl.ValueOrEmptyString(resource.RouteHash),
		Entrypoint:                AppengineVersionEntrypointToProto(resource.Entrypoint),
		VpcAccessConnector:        AppengineVersionVPCAccessConnectorToProto(resource.VPCAccessConnector),
		NetworkSettings:           AppengineVersionNetworkSettingsToProto(resource.NetworkSettings),
		InstanceSpec:              AppengineVersionInstanceSpecToProto(resource.InstanceSpec),
		App:                       dcl.ValueOrEmptyString(resource.App),
		Service:                   dcl.ValueOrEmptyString(resource.Service),
	}
	for _, r := range resource.InboundServices {
		p.InboundServices = append(p.InboundServices, appenginepb.AppengineVersionInboundServicesEnum(appenginepb.AppengineVersionInboundServicesEnum_value[string(r)]))
	}
	for _, r := range resource.Zones {
		p.Zones = append(p.Zones, r)
	}
	for _, r := range resource.BetaSettings {
		p.BetaSettings = append(p.BetaSettings, AppengineVersionBetaSettingsToProto(&r))
	}
	for _, r := range resource.Handlers {
		p.Handlers = append(p.Handlers, AppengineVersionHandlersToProto(&r))
	}
	for _, r := range resource.ErrorHandlers {
		p.ErrorHandlers = append(p.ErrorHandlers, AppengineVersionErrorHandlersToProto(&r))
	}
	for _, r := range resource.Libraries {
		p.Libraries = append(p.Libraries, AppengineVersionLibrariesToProto(&r))
	}
	for _, r := range resource.EnvVariables {
		p.EnvVariables = append(p.EnvVariables, AppengineVersionEnvVariablesToProto(&r))
	}
	for _, r := range resource.BuildEnvVariables {
		p.BuildEnvVariables = append(p.BuildEnvVariables, AppengineVersionBuildEnvVariablesToProto(&r))
	}

	return p
}

// ApplyVersion handles the gRPC request by passing it to the underlying Version Apply() method.
func (s *VersionServer) applyVersion(ctx context.Context, c *appengine.Client, request *appenginepb.ApplyAppengineVersionRequest) (*appenginepb.AppengineVersion, error) {
	p := ProtoToVersion(request.GetResource())
	res, err := c.ApplyVersion(ctx, p)
	if err != nil {
		return nil, err
	}
	r := VersionToProto(res)
	return r, nil
}

// ApplyVersion handles the gRPC request by passing it to the underlying Version Apply() method.
func (s *VersionServer) ApplyAppengineVersion(ctx context.Context, request *appenginepb.ApplyAppengineVersionRequest) (*appenginepb.AppengineVersion, error) {
	cl, err := createConfigVersion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyVersion(ctx, cl, request)
}

// DeleteVersion handles the gRPC request by passing it to the underlying Version Delete() method.
func (s *VersionServer) DeleteAppengineVersion(ctx context.Context, request *appenginepb.DeleteAppengineVersionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigVersion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteVersion(ctx, ProtoToVersion(request.GetResource()))

}

// ListVersion handles the gRPC request by passing it to the underlying VersionList() method.
func (s *VersionServer) ListAppengineVersion(ctx context.Context, request *appenginepb.ListAppengineVersionRequest) (*appenginepb.ListAppengineVersionResponse, error) {
	cl, err := createConfigVersion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListVersion(ctx, request.App, request.Service)
	if err != nil {
		return nil, err
	}
	var protos []*appenginepb.AppengineVersion
	for _, r := range resources.Items {
		rp := VersionToProto(r)
		protos = append(protos, rp)
	}
	return &appenginepb.ListAppengineVersionResponse{Items: protos}, nil
}

func createConfigVersion(ctx context.Context, service_account_file string) (*appengine.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return appengine.NewClient(conf), nil
}
