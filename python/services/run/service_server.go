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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	runpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/run/run_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run"
)

// Server implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceMetadata converts a ServiceMetadata resource from its proto representation.
func ProtoToRunServiceMetadata(p *runpb.RunServiceMetadata) *run.ServiceMetadata {
	if p == nil {
		return nil
	}
	obj := &run.ServiceMetadata{
		Name:                       dcl.StringOrNil(p.Name),
		GenerateName:               dcl.StringOrNil(p.GenerateName),
		Namespace:                  dcl.StringOrNil(p.Namespace),
		SelfLink:                   dcl.StringOrNil(p.SelfLink),
		Uid:                        dcl.StringOrNil(p.Uid),
		ResourceVersion:            dcl.StringOrNil(p.ResourceVersion),
		Generation:                 dcl.Int64OrNil(p.Generation),
		CreateTime:                 ProtoToRunServiceMetadataCreateTime(p.GetCreateTime()),
		DeleteTime:                 ProtoToRunServiceMetadataDeleteTime(p.GetDeleteTime()),
		DeletionGracePeriodSeconds: dcl.Int64OrNil(p.DeletionGracePeriodSeconds),
		ClusterName:                dcl.StringOrNil(p.ClusterName),
	}
	for _, r := range p.GetOwnerReferences() {
		obj.OwnerReferences = append(obj.OwnerReferences, *ProtoToRunServiceMetadataOwnerReferences(r))
	}
	for _, r := range p.GetFinalizers() {
		obj.Finalizers = append(obj.Finalizers, r)
	}
	return obj
}

// ProtoToServiceMetadataCreateTime converts a ServiceMetadataCreateTime resource from its proto representation.
func ProtoToRunServiceMetadataCreateTime(p *runpb.RunServiceMetadataCreateTime) *run.ServiceMetadataCreateTime {
	if p == nil {
		return nil
	}
	obj := &run.ServiceMetadataCreateTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToServiceMetadataOwnerReferences converts a ServiceMetadataOwnerReferences resource from its proto representation.
func ProtoToRunServiceMetadataOwnerReferences(p *runpb.RunServiceMetadataOwnerReferences) *run.ServiceMetadataOwnerReferences {
	if p == nil {
		return nil
	}
	obj := &run.ServiceMetadataOwnerReferences{
		ApiVersion:         dcl.StringOrNil(p.ApiVersion),
		Kind:               dcl.StringOrNil(p.Kind),
		Name:               dcl.StringOrNil(p.Name),
		Uid:                dcl.StringOrNil(p.Uid),
		Controller:         dcl.Bool(p.Controller),
		BlockOwnerDeletion: dcl.Bool(p.BlockOwnerDeletion),
	}
	return obj
}

// ProtoToServiceMetadataDeleteTime converts a ServiceMetadataDeleteTime resource from its proto representation.
func ProtoToRunServiceMetadataDeleteTime(p *runpb.RunServiceMetadataDeleteTime) *run.ServiceMetadataDeleteTime {
	if p == nil {
		return nil
	}
	obj := &run.ServiceMetadataDeleteTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToServiceSpec converts a ServiceSpec resource from its proto representation.
func ProtoToRunServiceSpec(p *runpb.RunServiceSpec) *run.ServiceSpec {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpec{
		Template: ProtoToRunServiceSpecTemplate(p.GetTemplate()),
	}
	for _, r := range p.GetTraffic() {
		obj.Traffic = append(obj.Traffic, *ProtoToRunServiceSpecTraffic(r))
	}
	return obj
}

// ProtoToServiceSpecTemplate converts a ServiceSpecTemplate resource from its proto representation.
func ProtoToRunServiceSpecTemplate(p *runpb.RunServiceSpecTemplate) *run.ServiceSpecTemplate {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplate{
		Metadata: ProtoToRunServiceSpecTemplateMetadata(p.GetMetadata()),
		Spec:     ProtoToRunServiceSpecTemplateSpec(p.GetSpec()),
	}
	return obj
}

// ProtoToServiceSpecTemplateMetadata converts a ServiceSpecTemplateMetadata resource from its proto representation.
func ProtoToRunServiceSpecTemplateMetadata(p *runpb.RunServiceSpecTemplateMetadata) *run.ServiceSpecTemplateMetadata {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateMetadata{
		Name:                       dcl.StringOrNil(p.Name),
		GenerateName:               dcl.StringOrNil(p.GenerateName),
		Namespace:                  dcl.StringOrNil(p.Namespace),
		SelfLink:                   dcl.StringOrNil(p.SelfLink),
		Uid:                        dcl.StringOrNil(p.Uid),
		ResourceVersion:            dcl.StringOrNil(p.ResourceVersion),
		Generation:                 dcl.Int64OrNil(p.Generation),
		CreateTime:                 ProtoToRunServiceSpecTemplateMetadataCreateTime(p.GetCreateTime()),
		DeleteTime:                 ProtoToRunServiceSpecTemplateMetadataDeleteTime(p.GetDeleteTime()),
		DeletionGracePeriodSeconds: dcl.Int64OrNil(p.DeletionGracePeriodSeconds),
		ClusterName:                dcl.StringOrNil(p.ClusterName),
	}
	for _, r := range p.GetOwnerReferences() {
		obj.OwnerReferences = append(obj.OwnerReferences, *ProtoToRunServiceSpecTemplateMetadataOwnerReferences(r))
	}
	for _, r := range p.GetFinalizers() {
		obj.Finalizers = append(obj.Finalizers, r)
	}
	return obj
}

// ProtoToServiceSpecTemplateMetadataCreateTime converts a ServiceSpecTemplateMetadataCreateTime resource from its proto representation.
func ProtoToRunServiceSpecTemplateMetadataCreateTime(p *runpb.RunServiceSpecTemplateMetadataCreateTime) *run.ServiceSpecTemplateMetadataCreateTime {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateMetadataCreateTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToServiceSpecTemplateMetadataOwnerReferences converts a ServiceSpecTemplateMetadataOwnerReferences resource from its proto representation.
func ProtoToRunServiceSpecTemplateMetadataOwnerReferences(p *runpb.RunServiceSpecTemplateMetadataOwnerReferences) *run.ServiceSpecTemplateMetadataOwnerReferences {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateMetadataOwnerReferences{
		ApiVersion:         dcl.StringOrNil(p.ApiVersion),
		Kind:               dcl.StringOrNil(p.Kind),
		Name:               dcl.StringOrNil(p.Name),
		Uid:                dcl.StringOrNil(p.Uid),
		Controller:         dcl.Bool(p.Controller),
		BlockOwnerDeletion: dcl.Bool(p.BlockOwnerDeletion),
	}
	return obj
}

// ProtoToServiceSpecTemplateMetadataDeleteTime converts a ServiceSpecTemplateMetadataDeleteTime resource from its proto representation.
func ProtoToRunServiceSpecTemplateMetadataDeleteTime(p *runpb.RunServiceSpecTemplateMetadataDeleteTime) *run.ServiceSpecTemplateMetadataDeleteTime {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateMetadataDeleteTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpec converts a ServiceSpecTemplateSpec resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpec(p *runpb.RunServiceSpecTemplateSpec) *run.ServiceSpecTemplateSpec {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpec{
		ContainerConcurrency: dcl.Int64OrNil(p.ContainerConcurrency),
		TimeoutSeconds:       dcl.Int64OrNil(p.TimeoutSeconds),
		ServiceAccountName:   dcl.StringOrNil(p.ServiceAccountName),
	}
	for _, r := range p.GetContainers() {
		obj.Containers = append(obj.Containers, *ProtoToRunServiceSpecTemplateSpecContainers(r))
	}
	for _, r := range p.GetVolumes() {
		obj.Volumes = append(obj.Volumes, *ProtoToRunServiceSpecTemplateSpecVolumes(r))
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainers converts a ServiceSpecTemplateSpecContainers resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainers(p *runpb.RunServiceSpecTemplateSpecContainers) *run.ServiceSpecTemplateSpecContainers {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainers{
		Name:                     dcl.StringOrNil(p.Name),
		Image:                    dcl.StringOrNil(p.Image),
		Resources:                ProtoToRunServiceSpecTemplateSpecContainersResources(p.GetResources()),
		WorkingDir:               dcl.StringOrNil(p.WorkingDir),
		LivenessProbe:            ProtoToRunServiceSpecTemplateSpecContainersLivenessProbe(p.GetLivenessProbe()),
		ReadinessProbe:           ProtoToRunServiceSpecTemplateSpecContainersReadinessProbe(p.GetReadinessProbe()),
		TerminationMessagePath:   dcl.StringOrNil(p.TerminationMessagePath),
		TerminationMessagePolicy: dcl.StringOrNil(p.TerminationMessagePolicy),
		ImagePullPolicy:          dcl.StringOrNil(p.ImagePullPolicy),
		SecurityContext:          ProtoToRunServiceSpecTemplateSpecContainersSecurityContext(p.GetSecurityContext()),
	}
	for _, r := range p.GetCommand() {
		obj.Command = append(obj.Command, r)
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetEnv() {
		obj.Env = append(obj.Env, *ProtoToRunServiceSpecTemplateSpecContainersEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToRunServiceSpecTemplateSpecContainersPorts(r))
	}
	for _, r := range p.GetEnvFrom() {
		obj.EnvFrom = append(obj.EnvFrom, *ProtoToRunServiceSpecTemplateSpecContainersEnvFrom(r))
	}
	for _, r := range p.GetVolumeMounts() {
		obj.VolumeMounts = append(obj.VolumeMounts, *ProtoToRunServiceSpecTemplateSpecContainersVolumeMounts(r))
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnv converts a ServiceSpecTemplateSpecContainersEnv resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnv(p *runpb.RunServiceSpecTemplateSpecContainersEnv) *run.ServiceSpecTemplateSpecContainersEnv {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnv{
		Name:      dcl.StringOrNil(p.Name),
		Value:     dcl.StringOrNil(p.Value),
		ValueFrom: ProtoToRunServiceSpecTemplateSpecContainersEnvValueFrom(p.GetValueFrom()),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvValueFrom converts a ServiceSpecTemplateSpecContainersEnvValueFrom resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvValueFrom(p *runpb.RunServiceSpecTemplateSpecContainersEnvValueFrom) *run.ServiceSpecTemplateSpecContainersEnvValueFrom {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvValueFrom{
		ConfigMapKeyRef: ProtoToRunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(p.GetConfigMapKeyRef()),
		SecretKeyRef:    ProtoToRunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(p.GetSecretKeyRef()),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef converts a ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(p *runpb.RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) *run.ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef{
		LocalObjectReference: ProtoToRunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(p.GetLocalObjectReference()),
		Key:                  dcl.StringOrNil(p.Key),
		Optional:             dcl.Bool(p.Optional),
		Name:                 dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference converts a ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(p *runpb.RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) *run.ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef converts a ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(p *runpb.RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) *run.ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef{
		LocalObjectReference: ProtoToRunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(p.GetLocalObjectReference()),
		Key:                  dcl.StringOrNil(p.Key),
		Optional:             dcl.Bool(p.Optional),
		Name:                 dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference converts a ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(p *runpb.RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) *run.ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersResources converts a ServiceSpecTemplateSpecContainersResources resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersResources(p *runpb.RunServiceSpecTemplateSpecContainersResources) *run.ServiceSpecTemplateSpecContainersResources {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersResources{}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersPorts converts a ServiceSpecTemplateSpecContainersPorts resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersPorts(p *runpb.RunServiceSpecTemplateSpecContainersPorts) *run.ServiceSpecTemplateSpecContainersPorts {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersPorts{
		Name:          dcl.StringOrNil(p.Name),
		ContainerPort: dcl.Int64OrNil(p.ContainerPort),
		Protocol:      dcl.StringOrNil(p.Protocol),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvFrom converts a ServiceSpecTemplateSpecContainersEnvFrom resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvFrom(p *runpb.RunServiceSpecTemplateSpecContainersEnvFrom) *run.ServiceSpecTemplateSpecContainersEnvFrom {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvFrom{
		Prefix:       dcl.StringOrNil(p.Prefix),
		ConfigMapRef: ProtoToRunServiceSpecTemplateSpecContainersEnvFromConfigMapRef(p.GetConfigMapRef()),
		SecretRef:    ProtoToRunServiceSpecTemplateSpecContainersEnvFromSecretRef(p.GetSecretRef()),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvFromConfigMapRef converts a ServiceSpecTemplateSpecContainersEnvFromConfigMapRef resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvFromConfigMapRef(p *runpb.RunServiceSpecTemplateSpecContainersEnvFromConfigMapRef) *run.ServiceSpecTemplateSpecContainersEnvFromConfigMapRef {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvFromConfigMapRef{
		LocalObjectReference: ProtoToRunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(p.GetLocalObjectReference()),
		Optional:             dcl.Bool(p.Optional),
		Name:                 dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference converts a ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(p *runpb.RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) *run.ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvFromSecretRef converts a ServiceSpecTemplateSpecContainersEnvFromSecretRef resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvFromSecretRef(p *runpb.RunServiceSpecTemplateSpecContainersEnvFromSecretRef) *run.ServiceSpecTemplateSpecContainersEnvFromSecretRef {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvFromSecretRef{
		LocalObjectReference: ProtoToRunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(p.GetLocalObjectReference()),
		Optional:             dcl.Bool(p.Optional),
		Name:                 dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference converts a ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(p *runpb.RunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) *run.ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersVolumeMounts converts a ServiceSpecTemplateSpecContainersVolumeMounts resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersVolumeMounts(p *runpb.RunServiceSpecTemplateSpecContainersVolumeMounts) *run.ServiceSpecTemplateSpecContainersVolumeMounts {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersVolumeMounts{
		Name:      dcl.StringOrNil(p.Name),
		ReadOnly:  dcl.Bool(p.ReadOnly),
		MountPath: dcl.StringOrNil(p.MountPath),
		SubPath:   dcl.StringOrNil(p.SubPath),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersLivenessProbe converts a ServiceSpecTemplateSpecContainersLivenessProbe resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersLivenessProbe(p *runpb.RunServiceSpecTemplateSpecContainersLivenessProbe) *run.ServiceSpecTemplateSpecContainersLivenessProbe {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersLivenessProbe{
		InitialDelaySeconds: dcl.Int64OrNil(p.InitialDelaySeconds),
		TimeoutSeconds:      dcl.Int64OrNil(p.TimeoutSeconds),
		PeriodSeconds:       dcl.Int64OrNil(p.PeriodSeconds),
		SuccessThreshold:    dcl.Int64OrNil(p.SuccessThreshold),
		FailureThreshold:    dcl.Int64OrNil(p.FailureThreshold),
		Exec:                ProtoToRunServiceSpecTemplateSpecContainersLivenessProbeExec(p.GetExec()),
		HttpGet:             ProtoToRunServiceSpecTemplateSpecContainersLivenessProbeHttpGet(p.GetHttpGet()),
		TcpSocket:           ProtoToRunServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(p.GetTcpSocket()),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersLivenessProbeExec converts a ServiceSpecTemplateSpecContainersLivenessProbeExec resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersLivenessProbeExec(p *runpb.RunServiceSpecTemplateSpecContainersLivenessProbeExec) *run.ServiceSpecTemplateSpecContainersLivenessProbeExec {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersLivenessProbeExec{
		Command: dcl.StringOrNil(p.Command),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersLivenessProbeHttpGet converts a ServiceSpecTemplateSpecContainersLivenessProbeHttpGet resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersLivenessProbeHttpGet(p *runpb.RunServiceSpecTemplateSpecContainersLivenessProbeHttpGet) *run.ServiceSpecTemplateSpecContainersLivenessProbeHttpGet {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersLivenessProbeHttpGet{
		Path:   dcl.StringOrNil(p.Path),
		Host:   dcl.StringOrNil(p.Host),
		Scheme: dcl.StringOrNil(p.Scheme),
	}
	for _, r := range p.GetHttpHeaders() {
		obj.HttpHeaders = append(obj.HttpHeaders, *ProtoToRunServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(r))
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders converts a ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(p *runpb.RunServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) *run.ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersLivenessProbeTcpSocket converts a ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(p *runpb.RunServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) *run.ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket{
		Port: dcl.Int64OrNil(p.Port),
		Host: dcl.StringOrNil(p.Host),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersReadinessProbe converts a ServiceSpecTemplateSpecContainersReadinessProbe resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersReadinessProbe(p *runpb.RunServiceSpecTemplateSpecContainersReadinessProbe) *run.ServiceSpecTemplateSpecContainersReadinessProbe {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersReadinessProbe{
		InitialDelaySeconds: dcl.Int64OrNil(p.InitialDelaySeconds),
		TimeoutSeconds:      dcl.Int64OrNil(p.TimeoutSeconds),
		PeriodSeconds:       dcl.Int64OrNil(p.PeriodSeconds),
		SuccessThreshold:    dcl.Int64OrNil(p.SuccessThreshold),
		FailureThreshold:    dcl.Int64OrNil(p.FailureThreshold),
		Exec:                ProtoToRunServiceSpecTemplateSpecContainersReadinessProbeExec(p.GetExec()),
		HttpGet:             ProtoToRunServiceSpecTemplateSpecContainersReadinessProbeHttpGet(p.GetHttpGet()),
		TcpSocket:           ProtoToRunServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(p.GetTcpSocket()),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersReadinessProbeExec converts a ServiceSpecTemplateSpecContainersReadinessProbeExec resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersReadinessProbeExec(p *runpb.RunServiceSpecTemplateSpecContainersReadinessProbeExec) *run.ServiceSpecTemplateSpecContainersReadinessProbeExec {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersReadinessProbeExec{
		Command: dcl.StringOrNil(p.Command),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersReadinessProbeHttpGet converts a ServiceSpecTemplateSpecContainersReadinessProbeHttpGet resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersReadinessProbeHttpGet(p *runpb.RunServiceSpecTemplateSpecContainersReadinessProbeHttpGet) *run.ServiceSpecTemplateSpecContainersReadinessProbeHttpGet {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersReadinessProbeHttpGet{
		Path:   dcl.StringOrNil(p.Path),
		Host:   dcl.StringOrNil(p.Host),
		Scheme: dcl.StringOrNil(p.Scheme),
	}
	for _, r := range p.GetHttpHeaders() {
		obj.HttpHeaders = append(obj.HttpHeaders, *ProtoToRunServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(r))
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders converts a ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(p *runpb.RunServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) *run.ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersReadinessProbeTcpSocket converts a ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(p *runpb.RunServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) *run.ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket{
		Port: dcl.Int64OrNil(p.Port),
		Host: dcl.StringOrNil(p.Host),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecContainersSecurityContext converts a ServiceSpecTemplateSpecContainersSecurityContext resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecContainersSecurityContext(p *runpb.RunServiceSpecTemplateSpecContainersSecurityContext) *run.ServiceSpecTemplateSpecContainersSecurityContext {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecContainersSecurityContext{
		RunAsUser: dcl.Int64OrNil(p.RunAsUser),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecVolumes converts a ServiceSpecTemplateSpecVolumes resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecVolumes(p *runpb.RunServiceSpecTemplateSpecVolumes) *run.ServiceSpecTemplateSpecVolumes {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecVolumes{
		Name:      dcl.StringOrNil(p.Name),
		Secret:    ProtoToRunServiceSpecTemplateSpecVolumesSecret(p.GetSecret()),
		ConfigMap: ProtoToRunServiceSpecTemplateSpecVolumesConfigMap(p.GetConfigMap()),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecVolumesSecret converts a ServiceSpecTemplateSpecVolumesSecret resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecVolumesSecret(p *runpb.RunServiceSpecTemplateSpecVolumesSecret) *run.ServiceSpecTemplateSpecVolumesSecret {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecVolumesSecret{
		SecretName:  dcl.StringOrNil(p.SecretName),
		DefaultMode: dcl.Int64OrNil(p.DefaultMode),
		Optional:    dcl.Bool(p.Optional),
	}
	for _, r := range p.GetItems() {
		obj.Items = append(obj.Items, *ProtoToRunServiceSpecTemplateSpecVolumesSecretItems(r))
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecVolumesSecretItems converts a ServiceSpecTemplateSpecVolumesSecretItems resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecVolumesSecretItems(p *runpb.RunServiceSpecTemplateSpecVolumesSecretItems) *run.ServiceSpecTemplateSpecVolumesSecretItems {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecVolumesSecretItems{
		Key:  dcl.StringOrNil(p.Key),
		Path: dcl.StringOrNil(p.Path),
		Mode: dcl.Int64OrNil(p.Mode),
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecVolumesConfigMap converts a ServiceSpecTemplateSpecVolumesConfigMap resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecVolumesConfigMap(p *runpb.RunServiceSpecTemplateSpecVolumesConfigMap) *run.ServiceSpecTemplateSpecVolumesConfigMap {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecVolumesConfigMap{
		Name:        dcl.StringOrNil(p.Name),
		DefaultMode: dcl.Int64OrNil(p.DefaultMode),
		Optional:    dcl.Bool(p.Optional),
	}
	for _, r := range p.GetItems() {
		obj.Items = append(obj.Items, *ProtoToRunServiceSpecTemplateSpecVolumesConfigMapItems(r))
	}
	return obj
}

// ProtoToServiceSpecTemplateSpecVolumesConfigMapItems converts a ServiceSpecTemplateSpecVolumesConfigMapItems resource from its proto representation.
func ProtoToRunServiceSpecTemplateSpecVolumesConfigMapItems(p *runpb.RunServiceSpecTemplateSpecVolumesConfigMapItems) *run.ServiceSpecTemplateSpecVolumesConfigMapItems {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTemplateSpecVolumesConfigMapItems{
		Key:  dcl.StringOrNil(p.Key),
		Path: dcl.StringOrNil(p.Path),
		Mode: dcl.Int64OrNil(p.Mode),
	}
	return obj
}

// ProtoToServiceSpecTraffic converts a ServiceSpecTraffic resource from its proto representation.
func ProtoToRunServiceSpecTraffic(p *runpb.RunServiceSpecTraffic) *run.ServiceSpecTraffic {
	if p == nil {
		return nil
	}
	obj := &run.ServiceSpecTraffic{
		ConfigurationName: dcl.StringOrNil(p.ConfigurationName),
		RevisionName:      dcl.StringOrNil(p.RevisionName),
		Percent:           dcl.Int64OrNil(p.Percent),
		Tag:               dcl.StringOrNil(p.Tag),
		LatestRevision:    dcl.Bool(p.LatestRevision),
		Url:               dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToServiceStatus converts a ServiceStatus resource from its proto representation.
func ProtoToRunServiceStatus(p *runpb.RunServiceStatus) *run.ServiceStatus {
	if p == nil {
		return nil
	}
	obj := &run.ServiceStatus{
		ObservedGeneration:        dcl.Int64OrNil(p.ObservedGeneration),
		LatestReadyRevisionName:   dcl.StringOrNil(p.LatestReadyRevisionName),
		LatestCreatedRevisionName: dcl.StringOrNil(p.LatestCreatedRevisionName),
		Url:                       dcl.StringOrNil(p.Url),
		Address:                   ProtoToRunServiceStatusAddress(p.GetAddress()),
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToRunServiceStatusConditions(r))
	}
	for _, r := range p.GetTraffic() {
		obj.Traffic = append(obj.Traffic, *ProtoToRunServiceStatusTraffic(r))
	}
	return obj
}

// ProtoToServiceStatusConditions converts a ServiceStatusConditions resource from its proto representation.
func ProtoToRunServiceStatusConditions(p *runpb.RunServiceStatusConditions) *run.ServiceStatusConditions {
	if p == nil {
		return nil
	}
	obj := &run.ServiceStatusConditions{
		Type:               dcl.StringOrNil(p.Type),
		Status:             dcl.StringOrNil(p.Status),
		Reason:             dcl.StringOrNil(p.Reason),
		Message:            dcl.StringOrNil(p.Message),
		LastTransitionTime: ProtoToRunServiceStatusConditionsLastTransitionTime(p.GetLastTransitionTime()),
		Severity:           dcl.StringOrNil(p.Severity),
	}
	return obj
}

// ProtoToServiceStatusConditionsLastTransitionTime converts a ServiceStatusConditionsLastTransitionTime resource from its proto representation.
func ProtoToRunServiceStatusConditionsLastTransitionTime(p *runpb.RunServiceStatusConditionsLastTransitionTime) *run.ServiceStatusConditionsLastTransitionTime {
	if p == nil {
		return nil
	}
	obj := &run.ServiceStatusConditionsLastTransitionTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToServiceStatusTraffic converts a ServiceStatusTraffic resource from its proto representation.
func ProtoToRunServiceStatusTraffic(p *runpb.RunServiceStatusTraffic) *run.ServiceStatusTraffic {
	if p == nil {
		return nil
	}
	obj := &run.ServiceStatusTraffic{
		ConfigurationName: dcl.StringOrNil(p.ConfigurationName),
		RevisionName:      dcl.StringOrNil(p.RevisionName),
		Percent:           dcl.Int64OrNil(p.Percent),
		Tag:               dcl.StringOrNil(p.Tag),
		LatestRevision:    dcl.Bool(p.LatestRevision),
		Url:               dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToServiceStatusAddress converts a ServiceStatusAddress resource from its proto representation.
func ProtoToRunServiceStatusAddress(p *runpb.RunServiceStatusAddress) *run.ServiceStatusAddress {
	if p == nil {
		return nil
	}
	obj := &run.ServiceStatusAddress{
		Url: dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *runpb.RunService) *run.Service {
	obj := &run.Service{
		ApiVersion: dcl.StringOrNil(p.ApiVersion),
		Kind:       dcl.StringOrNil(p.Kind),
		Metadata:   ProtoToRunServiceMetadata(p.GetMetadata()),
		Spec:       ProtoToRunServiceSpec(p.GetSpec()),
		Status:     ProtoToRunServiceStatus(p.GetStatus()),
		Project:    dcl.StringOrNil(p.Project),
		Location:   dcl.StringOrNil(p.Location),
		Name:       dcl.StringOrNil(p.Name),
	}
	return obj
}

// ServiceMetadataToProto converts a ServiceMetadata resource to its proto representation.
func RunServiceMetadataToProto(o *run.ServiceMetadata) *runpb.RunServiceMetadata {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceMetadata{
		Name:                       dcl.ValueOrEmptyString(o.Name),
		GenerateName:               dcl.ValueOrEmptyString(o.GenerateName),
		Namespace:                  dcl.ValueOrEmptyString(o.Namespace),
		SelfLink:                   dcl.ValueOrEmptyString(o.SelfLink),
		Uid:                        dcl.ValueOrEmptyString(o.Uid),
		ResourceVersion:            dcl.ValueOrEmptyString(o.ResourceVersion),
		Generation:                 dcl.ValueOrEmptyInt64(o.Generation),
		CreateTime:                 RunServiceMetadataCreateTimeToProto(o.CreateTime),
		DeleteTime:                 RunServiceMetadataDeleteTimeToProto(o.DeleteTime),
		DeletionGracePeriodSeconds: dcl.ValueOrEmptyInt64(o.DeletionGracePeriodSeconds),
		ClusterName:                dcl.ValueOrEmptyString(o.ClusterName),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	p.Annotations = make(map[string]string)
	for k, r := range o.Annotations {
		p.Annotations[k] = r
	}
	for _, r := range o.OwnerReferences {
		p.OwnerReferences = append(p.OwnerReferences, RunServiceMetadataOwnerReferencesToProto(&r))
	}
	for _, r := range o.Finalizers {
		p.Finalizers = append(p.Finalizers, r)
	}
	return p
}

// ServiceMetadataCreateTimeToProto converts a ServiceMetadataCreateTime resource to its proto representation.
func RunServiceMetadataCreateTimeToProto(o *run.ServiceMetadataCreateTime) *runpb.RunServiceMetadataCreateTime {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceMetadataCreateTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// ServiceMetadataOwnerReferencesToProto converts a ServiceMetadataOwnerReferences resource to its proto representation.
func RunServiceMetadataOwnerReferencesToProto(o *run.ServiceMetadataOwnerReferences) *runpb.RunServiceMetadataOwnerReferences {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceMetadataOwnerReferences{
		ApiVersion:         dcl.ValueOrEmptyString(o.ApiVersion),
		Kind:               dcl.ValueOrEmptyString(o.Kind),
		Name:               dcl.ValueOrEmptyString(o.Name),
		Uid:                dcl.ValueOrEmptyString(o.Uid),
		Controller:         dcl.ValueOrEmptyBool(o.Controller),
		BlockOwnerDeletion: dcl.ValueOrEmptyBool(o.BlockOwnerDeletion),
	}
	return p
}

// ServiceMetadataDeleteTimeToProto converts a ServiceMetadataDeleteTime resource to its proto representation.
func RunServiceMetadataDeleteTimeToProto(o *run.ServiceMetadataDeleteTime) *runpb.RunServiceMetadataDeleteTime {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceMetadataDeleteTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// ServiceSpecToProto converts a ServiceSpec resource to its proto representation.
func RunServiceSpecToProto(o *run.ServiceSpec) *runpb.RunServiceSpec {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpec{
		Template: RunServiceSpecTemplateToProto(o.Template),
	}
	for _, r := range o.Traffic {
		p.Traffic = append(p.Traffic, RunServiceSpecTrafficToProto(&r))
	}
	return p
}

// ServiceSpecTemplateToProto converts a ServiceSpecTemplate resource to its proto representation.
func RunServiceSpecTemplateToProto(o *run.ServiceSpecTemplate) *runpb.RunServiceSpecTemplate {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplate{
		Metadata: RunServiceSpecTemplateMetadataToProto(o.Metadata),
		Spec:     RunServiceSpecTemplateSpecToProto(o.Spec),
	}
	return p
}

// ServiceSpecTemplateMetadataToProto converts a ServiceSpecTemplateMetadata resource to its proto representation.
func RunServiceSpecTemplateMetadataToProto(o *run.ServiceSpecTemplateMetadata) *runpb.RunServiceSpecTemplateMetadata {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateMetadata{
		Name:                       dcl.ValueOrEmptyString(o.Name),
		GenerateName:               dcl.ValueOrEmptyString(o.GenerateName),
		Namespace:                  dcl.ValueOrEmptyString(o.Namespace),
		SelfLink:                   dcl.ValueOrEmptyString(o.SelfLink),
		Uid:                        dcl.ValueOrEmptyString(o.Uid),
		ResourceVersion:            dcl.ValueOrEmptyString(o.ResourceVersion),
		Generation:                 dcl.ValueOrEmptyInt64(o.Generation),
		CreateTime:                 RunServiceSpecTemplateMetadataCreateTimeToProto(o.CreateTime),
		DeleteTime:                 RunServiceSpecTemplateMetadataDeleteTimeToProto(o.DeleteTime),
		DeletionGracePeriodSeconds: dcl.ValueOrEmptyInt64(o.DeletionGracePeriodSeconds),
		ClusterName:                dcl.ValueOrEmptyString(o.ClusterName),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	p.Annotations = make(map[string]string)
	for k, r := range o.Annotations {
		p.Annotations[k] = r
	}
	for _, r := range o.OwnerReferences {
		p.OwnerReferences = append(p.OwnerReferences, RunServiceSpecTemplateMetadataOwnerReferencesToProto(&r))
	}
	for _, r := range o.Finalizers {
		p.Finalizers = append(p.Finalizers, r)
	}
	return p
}

// ServiceSpecTemplateMetadataCreateTimeToProto converts a ServiceSpecTemplateMetadataCreateTime resource to its proto representation.
func RunServiceSpecTemplateMetadataCreateTimeToProto(o *run.ServiceSpecTemplateMetadataCreateTime) *runpb.RunServiceSpecTemplateMetadataCreateTime {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateMetadataCreateTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// ServiceSpecTemplateMetadataOwnerReferencesToProto converts a ServiceSpecTemplateMetadataOwnerReferences resource to its proto representation.
func RunServiceSpecTemplateMetadataOwnerReferencesToProto(o *run.ServiceSpecTemplateMetadataOwnerReferences) *runpb.RunServiceSpecTemplateMetadataOwnerReferences {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateMetadataOwnerReferences{
		ApiVersion:         dcl.ValueOrEmptyString(o.ApiVersion),
		Kind:               dcl.ValueOrEmptyString(o.Kind),
		Name:               dcl.ValueOrEmptyString(o.Name),
		Uid:                dcl.ValueOrEmptyString(o.Uid),
		Controller:         dcl.ValueOrEmptyBool(o.Controller),
		BlockOwnerDeletion: dcl.ValueOrEmptyBool(o.BlockOwnerDeletion),
	}
	return p
}

// ServiceSpecTemplateMetadataDeleteTimeToProto converts a ServiceSpecTemplateMetadataDeleteTime resource to its proto representation.
func RunServiceSpecTemplateMetadataDeleteTimeToProto(o *run.ServiceSpecTemplateMetadataDeleteTime) *runpb.RunServiceSpecTemplateMetadataDeleteTime {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateMetadataDeleteTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// ServiceSpecTemplateSpecToProto converts a ServiceSpecTemplateSpec resource to its proto representation.
func RunServiceSpecTemplateSpecToProto(o *run.ServiceSpecTemplateSpec) *runpb.RunServiceSpecTemplateSpec {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpec{
		ContainerConcurrency: dcl.ValueOrEmptyInt64(o.ContainerConcurrency),
		TimeoutSeconds:       dcl.ValueOrEmptyInt64(o.TimeoutSeconds),
		ServiceAccountName:   dcl.ValueOrEmptyString(o.ServiceAccountName),
	}
	for _, r := range o.Containers {
		p.Containers = append(p.Containers, RunServiceSpecTemplateSpecContainersToProto(&r))
	}
	for _, r := range o.Volumes {
		p.Volumes = append(p.Volumes, RunServiceSpecTemplateSpecVolumesToProto(&r))
	}
	return p
}

// ServiceSpecTemplateSpecContainersToProto converts a ServiceSpecTemplateSpecContainers resource to its proto representation.
func RunServiceSpecTemplateSpecContainersToProto(o *run.ServiceSpecTemplateSpecContainers) *runpb.RunServiceSpecTemplateSpecContainers {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainers{
		Name:                     dcl.ValueOrEmptyString(o.Name),
		Image:                    dcl.ValueOrEmptyString(o.Image),
		Resources:                RunServiceSpecTemplateSpecContainersResourcesToProto(o.Resources),
		WorkingDir:               dcl.ValueOrEmptyString(o.WorkingDir),
		LivenessProbe:            RunServiceSpecTemplateSpecContainersLivenessProbeToProto(o.LivenessProbe),
		ReadinessProbe:           RunServiceSpecTemplateSpecContainersReadinessProbeToProto(o.ReadinessProbe),
		TerminationMessagePath:   dcl.ValueOrEmptyString(o.TerminationMessagePath),
		TerminationMessagePolicy: dcl.ValueOrEmptyString(o.TerminationMessagePolicy),
		ImagePullPolicy:          dcl.ValueOrEmptyString(o.ImagePullPolicy),
		SecurityContext:          RunServiceSpecTemplateSpecContainersSecurityContextToProto(o.SecurityContext),
	}
	for _, r := range o.Command {
		p.Command = append(p.Command, r)
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.Env {
		p.Env = append(p.Env, RunServiceSpecTemplateSpecContainersEnvToProto(&r))
	}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, RunServiceSpecTemplateSpecContainersPortsToProto(&r))
	}
	for _, r := range o.EnvFrom {
		p.EnvFrom = append(p.EnvFrom, RunServiceSpecTemplateSpecContainersEnvFromToProto(&r))
	}
	for _, r := range o.VolumeMounts {
		p.VolumeMounts = append(p.VolumeMounts, RunServiceSpecTemplateSpecContainersVolumeMountsToProto(&r))
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvToProto converts a ServiceSpecTemplateSpecContainersEnv resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvToProto(o *run.ServiceSpecTemplateSpecContainersEnv) *runpb.RunServiceSpecTemplateSpecContainersEnv {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnv{
		Name:      dcl.ValueOrEmptyString(o.Name),
		Value:     dcl.ValueOrEmptyString(o.Value),
		ValueFrom: RunServiceSpecTemplateSpecContainersEnvValueFromToProto(o.ValueFrom),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvValueFromToProto converts a ServiceSpecTemplateSpecContainersEnvValueFrom resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvValueFromToProto(o *run.ServiceSpecTemplateSpecContainersEnvValueFrom) *runpb.RunServiceSpecTemplateSpecContainersEnvValueFrom {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvValueFrom{
		ConfigMapKeyRef: RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefToProto(o.ConfigMapKeyRef),
		SecretKeyRef:    RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefToProto(o.SecretKeyRef),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefToProto converts a ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefToProto(o *run.ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) *runpb.RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef{
		LocalObjectReference: RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceToProto(o.LocalObjectReference),
		Key:                  dcl.ValueOrEmptyString(o.Key),
		Optional:             dcl.ValueOrEmptyBool(o.Optional),
		Name:                 dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceToProto converts a ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceToProto(o *run.ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) *runpb.RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefToProto converts a ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefToProto(o *run.ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) *runpb.RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef{
		LocalObjectReference: RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceToProto(o.LocalObjectReference),
		Key:                  dcl.ValueOrEmptyString(o.Key),
		Optional:             dcl.ValueOrEmptyBool(o.Optional),
		Name:                 dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceToProto converts a ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceToProto(o *run.ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) *runpb.RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ServiceSpecTemplateSpecContainersResourcesToProto converts a ServiceSpecTemplateSpecContainersResources resource to its proto representation.
func RunServiceSpecTemplateSpecContainersResourcesToProto(o *run.ServiceSpecTemplateSpecContainersResources) *runpb.RunServiceSpecTemplateSpecContainersResources {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersResources{}
	p.Limits = make(map[string]string)
	for k, r := range o.Limits {
		p.Limits[k] = r
	}
	p.Requests = make(map[string]string)
	for k, r := range o.Requests {
		p.Requests[k] = r
	}
	return p
}

// ServiceSpecTemplateSpecContainersPortsToProto converts a ServiceSpecTemplateSpecContainersPorts resource to its proto representation.
func RunServiceSpecTemplateSpecContainersPortsToProto(o *run.ServiceSpecTemplateSpecContainersPorts) *runpb.RunServiceSpecTemplateSpecContainersPorts {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersPorts{
		Name:          dcl.ValueOrEmptyString(o.Name),
		ContainerPort: dcl.ValueOrEmptyInt64(o.ContainerPort),
		Protocol:      dcl.ValueOrEmptyString(o.Protocol),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvFromToProto converts a ServiceSpecTemplateSpecContainersEnvFrom resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvFromToProto(o *run.ServiceSpecTemplateSpecContainersEnvFrom) *runpb.RunServiceSpecTemplateSpecContainersEnvFrom {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvFrom{
		Prefix:       dcl.ValueOrEmptyString(o.Prefix),
		ConfigMapRef: RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefToProto(o.ConfigMapRef),
		SecretRef:    RunServiceSpecTemplateSpecContainersEnvFromSecretRefToProto(o.SecretRef),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvFromConfigMapRefToProto converts a ServiceSpecTemplateSpecContainersEnvFromConfigMapRef resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefToProto(o *run.ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) *runpb.RunServiceSpecTemplateSpecContainersEnvFromConfigMapRef {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvFromConfigMapRef{
		LocalObjectReference: RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceToProto(o.LocalObjectReference),
		Optional:             dcl.ValueOrEmptyBool(o.Optional),
		Name:                 dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceToProto converts a ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceToProto(o *run.ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) *runpb.RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvFromSecretRefToProto converts a ServiceSpecTemplateSpecContainersEnvFromSecretRef resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvFromSecretRefToProto(o *run.ServiceSpecTemplateSpecContainersEnvFromSecretRef) *runpb.RunServiceSpecTemplateSpecContainersEnvFromSecretRef {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvFromSecretRef{
		LocalObjectReference: RunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceToProto(o.LocalObjectReference),
		Optional:             dcl.ValueOrEmptyBool(o.Optional),
		Name:                 dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceToProto converts a ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference resource to its proto representation.
func RunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceToProto(o *run.ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) *runpb.RunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ServiceSpecTemplateSpecContainersVolumeMountsToProto converts a ServiceSpecTemplateSpecContainersVolumeMounts resource to its proto representation.
func RunServiceSpecTemplateSpecContainersVolumeMountsToProto(o *run.ServiceSpecTemplateSpecContainersVolumeMounts) *runpb.RunServiceSpecTemplateSpecContainersVolumeMounts {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersVolumeMounts{
		Name:      dcl.ValueOrEmptyString(o.Name),
		ReadOnly:  dcl.ValueOrEmptyBool(o.ReadOnly),
		MountPath: dcl.ValueOrEmptyString(o.MountPath),
		SubPath:   dcl.ValueOrEmptyString(o.SubPath),
	}
	return p
}

// ServiceSpecTemplateSpecContainersLivenessProbeToProto converts a ServiceSpecTemplateSpecContainersLivenessProbe resource to its proto representation.
func RunServiceSpecTemplateSpecContainersLivenessProbeToProto(o *run.ServiceSpecTemplateSpecContainersLivenessProbe) *runpb.RunServiceSpecTemplateSpecContainersLivenessProbe {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersLivenessProbe{
		InitialDelaySeconds: dcl.ValueOrEmptyInt64(o.InitialDelaySeconds),
		TimeoutSeconds:      dcl.ValueOrEmptyInt64(o.TimeoutSeconds),
		PeriodSeconds:       dcl.ValueOrEmptyInt64(o.PeriodSeconds),
		SuccessThreshold:    dcl.ValueOrEmptyInt64(o.SuccessThreshold),
		FailureThreshold:    dcl.ValueOrEmptyInt64(o.FailureThreshold),
		Exec:                RunServiceSpecTemplateSpecContainersLivenessProbeExecToProto(o.Exec),
		HttpGet:             RunServiceSpecTemplateSpecContainersLivenessProbeHttpGetToProto(o.HttpGet),
		TcpSocket:           RunServiceSpecTemplateSpecContainersLivenessProbeTcpSocketToProto(o.TcpSocket),
	}
	return p
}

// ServiceSpecTemplateSpecContainersLivenessProbeExecToProto converts a ServiceSpecTemplateSpecContainersLivenessProbeExec resource to its proto representation.
func RunServiceSpecTemplateSpecContainersLivenessProbeExecToProto(o *run.ServiceSpecTemplateSpecContainersLivenessProbeExec) *runpb.RunServiceSpecTemplateSpecContainersLivenessProbeExec {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersLivenessProbeExec{
		Command: dcl.ValueOrEmptyString(o.Command),
	}
	return p
}

// ServiceSpecTemplateSpecContainersLivenessProbeHttpGetToProto converts a ServiceSpecTemplateSpecContainersLivenessProbeHttpGet resource to its proto representation.
func RunServiceSpecTemplateSpecContainersLivenessProbeHttpGetToProto(o *run.ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) *runpb.RunServiceSpecTemplateSpecContainersLivenessProbeHttpGet {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersLivenessProbeHttpGet{
		Path:   dcl.ValueOrEmptyString(o.Path),
		Host:   dcl.ValueOrEmptyString(o.Host),
		Scheme: dcl.ValueOrEmptyString(o.Scheme),
	}
	for _, r := range o.HttpHeaders {
		p.HttpHeaders = append(p.HttpHeaders, RunServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersToProto(&r))
	}
	return p
}

// ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersToProto converts a ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders resource to its proto representation.
func RunServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersToProto(o *run.ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) *runpb.RunServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// ServiceSpecTemplateSpecContainersLivenessProbeTcpSocketToProto converts a ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket resource to its proto representation.
func RunServiceSpecTemplateSpecContainersLivenessProbeTcpSocketToProto(o *run.ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) *runpb.RunServiceSpecTemplateSpecContainersLivenessProbeTcpSocket {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersLivenessProbeTcpSocket{
		Port: dcl.ValueOrEmptyInt64(o.Port),
		Host: dcl.ValueOrEmptyString(o.Host),
	}
	return p
}

// ServiceSpecTemplateSpecContainersReadinessProbeToProto converts a ServiceSpecTemplateSpecContainersReadinessProbe resource to its proto representation.
func RunServiceSpecTemplateSpecContainersReadinessProbeToProto(o *run.ServiceSpecTemplateSpecContainersReadinessProbe) *runpb.RunServiceSpecTemplateSpecContainersReadinessProbe {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersReadinessProbe{
		InitialDelaySeconds: dcl.ValueOrEmptyInt64(o.InitialDelaySeconds),
		TimeoutSeconds:      dcl.ValueOrEmptyInt64(o.TimeoutSeconds),
		PeriodSeconds:       dcl.ValueOrEmptyInt64(o.PeriodSeconds),
		SuccessThreshold:    dcl.ValueOrEmptyInt64(o.SuccessThreshold),
		FailureThreshold:    dcl.ValueOrEmptyInt64(o.FailureThreshold),
		Exec:                RunServiceSpecTemplateSpecContainersReadinessProbeExecToProto(o.Exec),
		HttpGet:             RunServiceSpecTemplateSpecContainersReadinessProbeHttpGetToProto(o.HttpGet),
		TcpSocket:           RunServiceSpecTemplateSpecContainersReadinessProbeTcpSocketToProto(o.TcpSocket),
	}
	return p
}

// ServiceSpecTemplateSpecContainersReadinessProbeExecToProto converts a ServiceSpecTemplateSpecContainersReadinessProbeExec resource to its proto representation.
func RunServiceSpecTemplateSpecContainersReadinessProbeExecToProto(o *run.ServiceSpecTemplateSpecContainersReadinessProbeExec) *runpb.RunServiceSpecTemplateSpecContainersReadinessProbeExec {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersReadinessProbeExec{
		Command: dcl.ValueOrEmptyString(o.Command),
	}
	return p
}

// ServiceSpecTemplateSpecContainersReadinessProbeHttpGetToProto converts a ServiceSpecTemplateSpecContainersReadinessProbeHttpGet resource to its proto representation.
func RunServiceSpecTemplateSpecContainersReadinessProbeHttpGetToProto(o *run.ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) *runpb.RunServiceSpecTemplateSpecContainersReadinessProbeHttpGet {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersReadinessProbeHttpGet{
		Path:   dcl.ValueOrEmptyString(o.Path),
		Host:   dcl.ValueOrEmptyString(o.Host),
		Scheme: dcl.ValueOrEmptyString(o.Scheme),
	}
	for _, r := range o.HttpHeaders {
		p.HttpHeaders = append(p.HttpHeaders, RunServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersToProto(&r))
	}
	return p
}

// ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersToProto converts a ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders resource to its proto representation.
func RunServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersToProto(o *run.ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) *runpb.RunServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// ServiceSpecTemplateSpecContainersReadinessProbeTcpSocketToProto converts a ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket resource to its proto representation.
func RunServiceSpecTemplateSpecContainersReadinessProbeTcpSocketToProto(o *run.ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) *runpb.RunServiceSpecTemplateSpecContainersReadinessProbeTcpSocket {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersReadinessProbeTcpSocket{
		Port: dcl.ValueOrEmptyInt64(o.Port),
		Host: dcl.ValueOrEmptyString(o.Host),
	}
	return p
}

// ServiceSpecTemplateSpecContainersSecurityContextToProto converts a ServiceSpecTemplateSpecContainersSecurityContext resource to its proto representation.
func RunServiceSpecTemplateSpecContainersSecurityContextToProto(o *run.ServiceSpecTemplateSpecContainersSecurityContext) *runpb.RunServiceSpecTemplateSpecContainersSecurityContext {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecContainersSecurityContext{
		RunAsUser: dcl.ValueOrEmptyInt64(o.RunAsUser),
	}
	return p
}

// ServiceSpecTemplateSpecVolumesToProto converts a ServiceSpecTemplateSpecVolumes resource to its proto representation.
func RunServiceSpecTemplateSpecVolumesToProto(o *run.ServiceSpecTemplateSpecVolumes) *runpb.RunServiceSpecTemplateSpecVolumes {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecVolumes{
		Name:      dcl.ValueOrEmptyString(o.Name),
		Secret:    RunServiceSpecTemplateSpecVolumesSecretToProto(o.Secret),
		ConfigMap: RunServiceSpecTemplateSpecVolumesConfigMapToProto(o.ConfigMap),
	}
	return p
}

// ServiceSpecTemplateSpecVolumesSecretToProto converts a ServiceSpecTemplateSpecVolumesSecret resource to its proto representation.
func RunServiceSpecTemplateSpecVolumesSecretToProto(o *run.ServiceSpecTemplateSpecVolumesSecret) *runpb.RunServiceSpecTemplateSpecVolumesSecret {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecVolumesSecret{
		SecretName:  dcl.ValueOrEmptyString(o.SecretName),
		DefaultMode: dcl.ValueOrEmptyInt64(o.DefaultMode),
		Optional:    dcl.ValueOrEmptyBool(o.Optional),
	}
	for _, r := range o.Items {
		p.Items = append(p.Items, RunServiceSpecTemplateSpecVolumesSecretItemsToProto(&r))
	}
	return p
}

// ServiceSpecTemplateSpecVolumesSecretItemsToProto converts a ServiceSpecTemplateSpecVolumesSecretItems resource to its proto representation.
func RunServiceSpecTemplateSpecVolumesSecretItemsToProto(o *run.ServiceSpecTemplateSpecVolumesSecretItems) *runpb.RunServiceSpecTemplateSpecVolumesSecretItems {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecVolumesSecretItems{
		Key:  dcl.ValueOrEmptyString(o.Key),
		Path: dcl.ValueOrEmptyString(o.Path),
		Mode: dcl.ValueOrEmptyInt64(o.Mode),
	}
	return p
}

// ServiceSpecTemplateSpecVolumesConfigMapToProto converts a ServiceSpecTemplateSpecVolumesConfigMap resource to its proto representation.
func RunServiceSpecTemplateSpecVolumesConfigMapToProto(o *run.ServiceSpecTemplateSpecVolumesConfigMap) *runpb.RunServiceSpecTemplateSpecVolumesConfigMap {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecVolumesConfigMap{
		Name:        dcl.ValueOrEmptyString(o.Name),
		DefaultMode: dcl.ValueOrEmptyInt64(o.DefaultMode),
		Optional:    dcl.ValueOrEmptyBool(o.Optional),
	}
	for _, r := range o.Items {
		p.Items = append(p.Items, RunServiceSpecTemplateSpecVolumesConfigMapItemsToProto(&r))
	}
	return p
}

// ServiceSpecTemplateSpecVolumesConfigMapItemsToProto converts a ServiceSpecTemplateSpecVolumesConfigMapItems resource to its proto representation.
func RunServiceSpecTemplateSpecVolumesConfigMapItemsToProto(o *run.ServiceSpecTemplateSpecVolumesConfigMapItems) *runpb.RunServiceSpecTemplateSpecVolumesConfigMapItems {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTemplateSpecVolumesConfigMapItems{
		Key:  dcl.ValueOrEmptyString(o.Key),
		Path: dcl.ValueOrEmptyString(o.Path),
		Mode: dcl.ValueOrEmptyInt64(o.Mode),
	}
	return p
}

// ServiceSpecTrafficToProto converts a ServiceSpecTraffic resource to its proto representation.
func RunServiceSpecTrafficToProto(o *run.ServiceSpecTraffic) *runpb.RunServiceSpecTraffic {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceSpecTraffic{
		ConfigurationName: dcl.ValueOrEmptyString(o.ConfigurationName),
		RevisionName:      dcl.ValueOrEmptyString(o.RevisionName),
		Percent:           dcl.ValueOrEmptyInt64(o.Percent),
		Tag:               dcl.ValueOrEmptyString(o.Tag),
		LatestRevision:    dcl.ValueOrEmptyBool(o.LatestRevision),
		Url:               dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// ServiceStatusToProto converts a ServiceStatus resource to its proto representation.
func RunServiceStatusToProto(o *run.ServiceStatus) *runpb.RunServiceStatus {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceStatus{
		ObservedGeneration:        dcl.ValueOrEmptyInt64(o.ObservedGeneration),
		LatestReadyRevisionName:   dcl.ValueOrEmptyString(o.LatestReadyRevisionName),
		LatestCreatedRevisionName: dcl.ValueOrEmptyString(o.LatestCreatedRevisionName),
		Url:                       dcl.ValueOrEmptyString(o.Url),
		Address:                   RunServiceStatusAddressToProto(o.Address),
	}
	for _, r := range o.Conditions {
		p.Conditions = append(p.Conditions, RunServiceStatusConditionsToProto(&r))
	}
	for _, r := range o.Traffic {
		p.Traffic = append(p.Traffic, RunServiceStatusTrafficToProto(&r))
	}
	return p
}

// ServiceStatusConditionsToProto converts a ServiceStatusConditions resource to its proto representation.
func RunServiceStatusConditionsToProto(o *run.ServiceStatusConditions) *runpb.RunServiceStatusConditions {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceStatusConditions{
		Type:               dcl.ValueOrEmptyString(o.Type),
		Status:             dcl.ValueOrEmptyString(o.Status),
		Reason:             dcl.ValueOrEmptyString(o.Reason),
		Message:            dcl.ValueOrEmptyString(o.Message),
		LastTransitionTime: RunServiceStatusConditionsLastTransitionTimeToProto(o.LastTransitionTime),
		Severity:           dcl.ValueOrEmptyString(o.Severity),
	}
	return p
}

// ServiceStatusConditionsLastTransitionTimeToProto converts a ServiceStatusConditionsLastTransitionTime resource to its proto representation.
func RunServiceStatusConditionsLastTransitionTimeToProto(o *run.ServiceStatusConditionsLastTransitionTime) *runpb.RunServiceStatusConditionsLastTransitionTime {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceStatusConditionsLastTransitionTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// ServiceStatusTrafficToProto converts a ServiceStatusTraffic resource to its proto representation.
func RunServiceStatusTrafficToProto(o *run.ServiceStatusTraffic) *runpb.RunServiceStatusTraffic {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceStatusTraffic{
		ConfigurationName: dcl.ValueOrEmptyString(o.ConfigurationName),
		RevisionName:      dcl.ValueOrEmptyString(o.RevisionName),
		Percent:           dcl.ValueOrEmptyInt64(o.Percent),
		Tag:               dcl.ValueOrEmptyString(o.Tag),
		LatestRevision:    dcl.ValueOrEmptyBool(o.LatestRevision),
		Url:               dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// ServiceStatusAddressToProto converts a ServiceStatusAddress resource to its proto representation.
func RunServiceStatusAddressToProto(o *run.ServiceStatusAddress) *runpb.RunServiceStatusAddress {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceStatusAddress{
		Url: dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *run.Service) *runpb.RunService {
	p := &runpb.RunService{
		ApiVersion: dcl.ValueOrEmptyString(resource.ApiVersion),
		Kind:       dcl.ValueOrEmptyString(resource.Kind),
		Metadata:   RunServiceMetadataToProto(resource.Metadata),
		Spec:       RunServiceSpecToProto(resource.Spec),
		Status:     RunServiceStatusToProto(resource.Status),
		Project:    dcl.ValueOrEmptyString(resource.Project),
		Location:   dcl.ValueOrEmptyString(resource.Location),
		Name:       dcl.ValueOrEmptyString(resource.Name),
	}

	return p
}

// ApplyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) applyService(ctx context.Context, c *run.Client, request *runpb.ApplyRunServiceRequest) (*runpb.RunService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// ApplyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyRunService(ctx context.Context, request *runpb.ApplyRunServiceRequest) (*runpb.RunService, error) {
	cl, err := createConfigService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteRunService(ctx context.Context, request *runpb.DeleteRunServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListRunService(ctx context.Context, request *runpb.ListRunServiceRequest) (*runpb.ListRunServiceResponse, error) {
	cl, err := createConfigService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*runpb.RunService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	return &runpb.ListRunServiceResponse{Items: protos}, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*run.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return run.NewClient(conf), nil
}
