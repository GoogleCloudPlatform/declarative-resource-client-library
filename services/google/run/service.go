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
package run

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Service struct {
	ApiVersion *string          `json:"apiVersion"`
	Kind       *string          `json:"kind"`
	Metadata   *ServiceMetadata `json:"metadata"`
	Spec       *ServiceSpec     `json:"spec"`
	Status     *ServiceStatus   `json:"status"`
	Project    *string          `json:"project"`
	Location   *string          `json:"location"`
	Name       *string          `json:"name"`
}

func (r *Service) String() string {
	return dcl.SprintResource(r)
}

type ServiceMetadata struct {
	empty                      bool                             `json:"-"`
	Name                       *string                          `json:"name"`
	GenerateName               *string                          `json:"generateName"`
	Namespace                  *string                          `json:"namespace"`
	SelfLink                   *string                          `json:"selfLink"`
	Uid                        *string                          `json:"uid"`
	ResourceVersion            *string                          `json:"resourceVersion"`
	Generation                 *int64                           `json:"generation"`
	CreateTime                 *ServiceMetadataCreateTime       `json:"createTime"`
	Labels                     map[string]string                `json:"labels"`
	Annotations                map[string]string                `json:"annotations"`
	OwnerReferences            []ServiceMetadataOwnerReferences `json:"ownerReferences"`
	DeleteTime                 *ServiceMetadataDeleteTime       `json:"deleteTime"`
	DeletionGracePeriodSeconds *int64                           `json:"deletionGracePeriodSeconds"`
	Finalizers                 []string                         `json:"finalizers"`
	ClusterName                *string                          `json:"clusterName"`
}

// This object is used to assert a desired state where this ServiceMetadata is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceMetadata *ServiceMetadata = &ServiceMetadata{empty: true}

func (r *ServiceMetadata) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceMetadata) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceMetadataCreateTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this ServiceMetadataCreateTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceMetadataCreateTime *ServiceMetadataCreateTime = &ServiceMetadataCreateTime{empty: true}

func (r *ServiceMetadataCreateTime) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceMetadataCreateTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceMetadataOwnerReferences struct {
	empty              bool    `json:"-"`
	ApiVersion         *string `json:"apiVersion"`
	Kind               *string `json:"kind"`
	Name               *string `json:"name"`
	Uid                *string `json:"uid"`
	Controller         *bool   `json:"controller"`
	BlockOwnerDeletion *bool   `json:"blockOwnerDeletion"`
}

// This object is used to assert a desired state where this ServiceMetadataOwnerReferences is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceMetadataOwnerReferences *ServiceMetadataOwnerReferences = &ServiceMetadataOwnerReferences{empty: true}

func (r *ServiceMetadataOwnerReferences) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceMetadataOwnerReferences) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceMetadataDeleteTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this ServiceMetadataDeleteTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceMetadataDeleteTime *ServiceMetadataDeleteTime = &ServiceMetadataDeleteTime{empty: true}

func (r *ServiceMetadataDeleteTime) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceMetadataDeleteTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpec struct {
	empty    bool                 `json:"-"`
	Template *ServiceSpecTemplate `json:"template"`
	Traffic  []ServiceSpecTraffic `json:"traffic"`
}

// This object is used to assert a desired state where this ServiceSpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpec *ServiceSpec = &ServiceSpec{empty: true}

func (r *ServiceSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplate struct {
	empty    bool                         `json:"-"`
	Metadata *ServiceSpecTemplateMetadata `json:"metadata"`
	Spec     *ServiceSpecTemplateSpec     `json:"spec"`
}

// This object is used to assert a desired state where this ServiceSpecTemplate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplate *ServiceSpecTemplate = &ServiceSpecTemplate{empty: true}

func (r *ServiceSpecTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateMetadata struct {
	empty                      bool                                         `json:"-"`
	Name                       *string                                      `json:"name"`
	GenerateName               *string                                      `json:"generateName"`
	Namespace                  *string                                      `json:"namespace"`
	SelfLink                   *string                                      `json:"selfLink"`
	Uid                        *string                                      `json:"uid"`
	ResourceVersion            *string                                      `json:"resourceVersion"`
	Generation                 *int64                                       `json:"generation"`
	CreateTime                 *ServiceSpecTemplateMetadataCreateTime       `json:"createTime"`
	Labels                     map[string]string                            `json:"labels"`
	Annotations                map[string]string                            `json:"annotations"`
	OwnerReferences            []ServiceSpecTemplateMetadataOwnerReferences `json:"ownerReferences"`
	DeleteTime                 *ServiceSpecTemplateMetadataDeleteTime       `json:"deleteTime"`
	DeletionGracePeriodSeconds *int64                                       `json:"deletionGracePeriodSeconds"`
	Finalizers                 []string                                     `json:"finalizers"`
	ClusterName                *string                                      `json:"clusterName"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateMetadata is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateMetadata *ServiceSpecTemplateMetadata = &ServiceSpecTemplateMetadata{empty: true}

func (r *ServiceSpecTemplateMetadata) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateMetadata) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateMetadataCreateTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateMetadataCreateTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateMetadataCreateTime *ServiceSpecTemplateMetadataCreateTime = &ServiceSpecTemplateMetadataCreateTime{empty: true}

func (r *ServiceSpecTemplateMetadataCreateTime) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateMetadataCreateTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateMetadataOwnerReferences struct {
	empty              bool    `json:"-"`
	ApiVersion         *string `json:"apiVersion"`
	Kind               *string `json:"kind"`
	Name               *string `json:"name"`
	Uid                *string `json:"uid"`
	Controller         *bool   `json:"controller"`
	BlockOwnerDeletion *bool   `json:"blockOwnerDeletion"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateMetadataOwnerReferences is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateMetadataOwnerReferences *ServiceSpecTemplateMetadataOwnerReferences = &ServiceSpecTemplateMetadataOwnerReferences{empty: true}

func (r *ServiceSpecTemplateMetadataOwnerReferences) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateMetadataOwnerReferences) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateMetadataDeleteTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateMetadataDeleteTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateMetadataDeleteTime *ServiceSpecTemplateMetadataDeleteTime = &ServiceSpecTemplateMetadataDeleteTime{empty: true}

func (r *ServiceSpecTemplateMetadataDeleteTime) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateMetadataDeleteTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpec struct {
	empty                bool                                `json:"-"`
	ContainerConcurrency *int64                              `json:"containerConcurrency"`
	TimeoutSeconds       *int64                              `json:"timeoutSeconds"`
	ServiceAccountName   *string                             `json:"serviceAccountName"`
	Containers           []ServiceSpecTemplateSpecContainers `json:"containers"`
	Volumes              []ServiceSpecTemplateSpecVolumes    `json:"volumes"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpec *ServiceSpecTemplateSpec = &ServiceSpecTemplateSpec{empty: true}

func (r *ServiceSpecTemplateSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainers struct {
	empty                    bool                                              `json:"-"`
	Name                     *string                                           `json:"name"`
	Image                    *string                                           `json:"image"`
	Command                  []string                                          `json:"command"`
	Args                     []string                                          `json:"args"`
	Env                      []ServiceSpecTemplateSpecContainersEnv            `json:"env"`
	Resources                *ServiceSpecTemplateSpecContainersResources       `json:"resources"`
	WorkingDir               *string                                           `json:"workingDir"`
	Ports                    []ServiceSpecTemplateSpecContainersPorts          `json:"ports"`
	EnvFrom                  []ServiceSpecTemplateSpecContainersEnvFrom        `json:"envFrom"`
	VolumeMounts             []ServiceSpecTemplateSpecContainersVolumeMounts   `json:"volumeMounts"`
	LivenessProbe            *ServiceSpecTemplateSpecContainersLivenessProbe   `json:"livenessProbe"`
	ReadinessProbe           *ServiceSpecTemplateSpecContainersReadinessProbe  `json:"readinessProbe"`
	TerminationMessagePath   *string                                           `json:"terminationMessagePath"`
	TerminationMessagePolicy *string                                           `json:"terminationMessagePolicy"`
	ImagePullPolicy          *string                                           `json:"imagePullPolicy"`
	SecurityContext          *ServiceSpecTemplateSpecContainersSecurityContext `json:"securityContext"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainers *ServiceSpecTemplateSpecContainers = &ServiceSpecTemplateSpecContainers{empty: true}

func (r *ServiceSpecTemplateSpecContainers) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnv struct {
	empty     bool                                           `json:"-"`
	Name      *string                                        `json:"name"`
	Value     *string                                        `json:"value"`
	ValueFrom *ServiceSpecTemplateSpecContainersEnvValueFrom `json:"valueFrom"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnv is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnv *ServiceSpecTemplateSpecContainersEnv = &ServiceSpecTemplateSpecContainersEnv{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnv) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnv) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvValueFrom struct {
	empty           bool                                                          `json:"-"`
	ConfigMapKeyRef *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef `json:"configMapKeyRef"`
	SecretKeyRef    *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef    `json:"secretKeyRef"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFrom is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFrom *ServiceSpecTemplateSpecContainersEnvValueFrom = &ServiceSpecTemplateSpecContainersEnvValueFrom{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFrom) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvValueFrom) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef struct {
	empty                bool                                                                              `json:"-"`
	LocalObjectReference *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference `json:"localObjectReference"`
	Key                  *string                                                                           `json:"key"`
	Optional             *bool                                                                             `json:"optional"`
	Name                 *string                                                                           `json:"name"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef = &ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference = &ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef struct {
	empty                bool                                                                           `json:"-"`
	LocalObjectReference *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference `json:"localObjectReference"`
	Key                  *string                                                                        `json:"key"`
	Optional             *bool                                                                          `json:"optional"`
	Name                 *string                                                                        `json:"name"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef = &ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference = &ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersResources struct {
	empty    bool              `json:"-"`
	Limits   map[string]string `json:"limits"`
	Requests map[string]string `json:"requests"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersResources is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersResources *ServiceSpecTemplateSpecContainersResources = &ServiceSpecTemplateSpecContainersResources{empty: true}

func (r *ServiceSpecTemplateSpecContainersResources) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersPorts struct {
	empty         bool    `json:"-"`
	Name          *string `json:"name"`
	ContainerPort *int64  `json:"containerPort"`
	Protocol      *string `json:"protocol"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersPorts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersPorts *ServiceSpecTemplateSpecContainersPorts = &ServiceSpecTemplateSpecContainersPorts{empty: true}

func (r *ServiceSpecTemplateSpecContainersPorts) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersPorts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvFrom struct {
	empty        bool                                                  `json:"-"`
	Prefix       *string                                               `json:"prefix"`
	ConfigMapRef *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef `json:"configMapRef"`
	SecretRef    *ServiceSpecTemplateSpecContainersEnvFromSecretRef    `json:"secretRef"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFrom is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFrom *ServiceSpecTemplateSpecContainersEnvFrom = &ServiceSpecTemplateSpecContainersEnvFrom{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFrom) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvFrom) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvFromConfigMapRef struct {
	empty                bool                                                                      `json:"-"`
	LocalObjectReference *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference `json:"localObjectReference"`
	Optional             *bool                                                                     `json:"optional"`
	Name                 *string                                                                   `json:"name"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFromConfigMapRef is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRef *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef = &ServiceSpecTemplateSpecContainersEnvFromConfigMapRef{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference = &ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvFromSecretRef struct {
	empty                bool                                                                   `json:"-"`
	LocalObjectReference *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference `json:"localObjectReference"`
	Optional             *bool                                                                  `json:"optional"`
	Name                 *string                                                                `json:"name"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFromSecretRef is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFromSecretRef *ServiceSpecTemplateSpecContainersEnvFromSecretRef = &ServiceSpecTemplateSpecContainersEnvFromSecretRef{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRef) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRef) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference = &ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersVolumeMounts struct {
	empty     bool    `json:"-"`
	Name      *string `json:"name"`
	ReadOnly  *bool   `json:"readOnly"`
	MountPath *string `json:"mountPath"`
	SubPath   *string `json:"subPath"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersVolumeMounts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersVolumeMounts *ServiceSpecTemplateSpecContainersVolumeMounts = &ServiceSpecTemplateSpecContainersVolumeMounts{empty: true}

func (r *ServiceSpecTemplateSpecContainersVolumeMounts) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersVolumeMounts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersLivenessProbe struct {
	empty               bool                                                     `json:"-"`
	InitialDelaySeconds *int64                                                   `json:"initialDelaySeconds"`
	TimeoutSeconds      *int64                                                   `json:"timeoutSeconds"`
	PeriodSeconds       *int64                                                   `json:"periodSeconds"`
	SuccessThreshold    *int64                                                   `json:"successThreshold"`
	FailureThreshold    *int64                                                   `json:"failureThreshold"`
	Exec                *ServiceSpecTemplateSpecContainersLivenessProbeExec      `json:"exec"`
	HttpGet             *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet   `json:"httpGet"`
	TcpSocket           *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket `json:"tcpSocket"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbe *ServiceSpecTemplateSpecContainersLivenessProbe = &ServiceSpecTemplateSpecContainersLivenessProbe{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbe) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersLivenessProbe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersLivenessProbeExec struct {
	empty   bool    `json:"-"`
	Command *string `json:"command"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbeExec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbeExec *ServiceSpecTemplateSpecContainersLivenessProbeExec = &ServiceSpecTemplateSpecContainersLivenessProbeExec{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeExec) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeExec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersLivenessProbeHttpGet struct {
	empty       bool                                                               `json:"-"`
	Path        *string                                                            `json:"path"`
	Host        *string                                                            `json:"host"`
	Scheme      *string                                                            `json:"scheme"`
	HttpHeaders []ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders `json:"httpHeaders"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbeHttpGet is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGet *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet = &ServiceSpecTemplateSpecContainersLivenessProbeHttpGet{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders = &ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket struct {
	empty bool    `json:"-"`
	Port  *int64  `json:"port"`
	Host  *string `json:"host"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbeTcpSocket *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket = &ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersReadinessProbe struct {
	empty               bool                                                      `json:"-"`
	InitialDelaySeconds *int64                                                    `json:"initialDelaySeconds"`
	TimeoutSeconds      *int64                                                    `json:"timeoutSeconds"`
	PeriodSeconds       *int64                                                    `json:"periodSeconds"`
	SuccessThreshold    *int64                                                    `json:"successThreshold"`
	FailureThreshold    *int64                                                    `json:"failureThreshold"`
	Exec                *ServiceSpecTemplateSpecContainersReadinessProbeExec      `json:"exec"`
	HttpGet             *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet   `json:"httpGet"`
	TcpSocket           *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket `json:"tcpSocket"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbe *ServiceSpecTemplateSpecContainersReadinessProbe = &ServiceSpecTemplateSpecContainersReadinessProbe{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbe) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersReadinessProbe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersReadinessProbeExec struct {
	empty   bool    `json:"-"`
	Command *string `json:"command"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbeExec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbeExec *ServiceSpecTemplateSpecContainersReadinessProbeExec = &ServiceSpecTemplateSpecContainersReadinessProbeExec{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeExec) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeExec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersReadinessProbeHttpGet struct {
	empty       bool                                                                `json:"-"`
	Path        *string                                                             `json:"path"`
	Host        *string                                                             `json:"host"`
	Scheme      *string                                                             `json:"scheme"`
	HttpHeaders []ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders `json:"httpHeaders"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbeHttpGet is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGet *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet = &ServiceSpecTemplateSpecContainersReadinessProbeHttpGet{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders = &ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket struct {
	empty bool    `json:"-"`
	Port  *int64  `json:"port"`
	Host  *string `json:"host"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbeTcpSocket *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket = &ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecContainersSecurityContext struct {
	empty     bool   `json:"-"`
	RunAsUser *int64 `json:"runAsUser"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersSecurityContext is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersSecurityContext *ServiceSpecTemplateSpecContainersSecurityContext = &ServiceSpecTemplateSpecContainersSecurityContext{empty: true}

func (r *ServiceSpecTemplateSpecContainersSecurityContext) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecContainersSecurityContext) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecVolumes struct {
	empty     bool                                     `json:"-"`
	Name      *string                                  `json:"name"`
	Secret    *ServiceSpecTemplateSpecVolumesSecret    `json:"secret"`
	ConfigMap *ServiceSpecTemplateSpecVolumesConfigMap `json:"configMap"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumes *ServiceSpecTemplateSpecVolumes = &ServiceSpecTemplateSpecVolumes{empty: true}

func (r *ServiceSpecTemplateSpecVolumes) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecVolumes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecVolumesSecret struct {
	empty       bool                                        `json:"-"`
	SecretName  *string                                     `json:"secretName"`
	Items       []ServiceSpecTemplateSpecVolumesSecretItems `json:"items"`
	DefaultMode *int64                                      `json:"defaultMode"`
	Optional    *bool                                       `json:"optional"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumesSecret is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumesSecret *ServiceSpecTemplateSpecVolumesSecret = &ServiceSpecTemplateSpecVolumesSecret{empty: true}

func (r *ServiceSpecTemplateSpecVolumesSecret) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecVolumesSecret) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecVolumesSecretItems struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Path  *string `json:"path"`
	Mode  *int64  `json:"mode"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumesSecretItems is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumesSecretItems *ServiceSpecTemplateSpecVolumesSecretItems = &ServiceSpecTemplateSpecVolumesSecretItems{empty: true}

func (r *ServiceSpecTemplateSpecVolumesSecretItems) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecVolumesSecretItems) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecVolumesConfigMap struct {
	empty       bool                                           `json:"-"`
	Name        *string                                        `json:"name"`
	Items       []ServiceSpecTemplateSpecVolumesConfigMapItems `json:"items"`
	DefaultMode *int64                                         `json:"defaultMode"`
	Optional    *bool                                          `json:"optional"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumesConfigMap is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumesConfigMap *ServiceSpecTemplateSpecVolumesConfigMap = &ServiceSpecTemplateSpecVolumesConfigMap{empty: true}

func (r *ServiceSpecTemplateSpecVolumesConfigMap) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecVolumesConfigMap) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTemplateSpecVolumesConfigMapItems struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Path  *string `json:"path"`
	Mode  *int64  `json:"mode"`
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumesConfigMapItems is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumesConfigMapItems *ServiceSpecTemplateSpecVolumesConfigMapItems = &ServiceSpecTemplateSpecVolumesConfigMapItems{empty: true}

func (r *ServiceSpecTemplateSpecVolumesConfigMapItems) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTemplateSpecVolumesConfigMapItems) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceSpecTraffic struct {
	empty             bool    `json:"-"`
	ConfigurationName *string `json:"configurationName"`
	RevisionName      *string `json:"revisionName"`
	Percent           *int64  `json:"percent"`
	Tag               *string `json:"tag"`
	LatestRevision    *bool   `json:"latestRevision"`
	Url               *string `json:"url"`
}

// This object is used to assert a desired state where this ServiceSpecTraffic is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTraffic *ServiceSpecTraffic = &ServiceSpecTraffic{empty: true}

func (r *ServiceSpecTraffic) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceSpecTraffic) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceStatus struct {
	empty                     bool                      `json:"-"`
	ObservedGeneration        *int64                    `json:"observedGeneration"`
	Conditions                []ServiceStatusConditions `json:"conditions"`
	LatestReadyRevisionName   *string                   `json:"latestReadyRevisionName"`
	LatestCreatedRevisionName *string                   `json:"latestCreatedRevisionName"`
	Traffic                   []ServiceStatusTraffic    `json:"traffic"`
	Url                       *string                   `json:"url"`
	Address                   *ServiceStatusAddress     `json:"address"`
}

// This object is used to assert a desired state where this ServiceStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatus *ServiceStatus = &ServiceStatus{empty: true}

func (r *ServiceStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceStatusConditions struct {
	empty              bool                                       `json:"-"`
	Type               *string                                    `json:"type"`
	Status             *string                                    `json:"status"`
	Reason             *string                                    `json:"reason"`
	Message            *string                                    `json:"message"`
	LastTransitionTime *ServiceStatusConditionsLastTransitionTime `json:"lastTransitionTime"`
	Severity           *string                                    `json:"severity"`
}

// This object is used to assert a desired state where this ServiceStatusConditions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatusConditions *ServiceStatusConditions = &ServiceStatusConditions{empty: true}

func (r *ServiceStatusConditions) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceStatusConditions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceStatusConditionsLastTransitionTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this ServiceStatusConditionsLastTransitionTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatusConditionsLastTransitionTime *ServiceStatusConditionsLastTransitionTime = &ServiceStatusConditionsLastTransitionTime{empty: true}

func (r *ServiceStatusConditionsLastTransitionTime) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceStatusConditionsLastTransitionTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceStatusTraffic struct {
	empty             bool    `json:"-"`
	ConfigurationName *string `json:"configurationName"`
	RevisionName      *string `json:"revisionName"`
	Percent           *int64  `json:"percent"`
	Tag               *string `json:"tag"`
	LatestRevision    *bool   `json:"latestRevision"`
	Url               *string `json:"url"`
}

// This object is used to assert a desired state where this ServiceStatusTraffic is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatusTraffic *ServiceStatusTraffic = &ServiceStatusTraffic{empty: true}

func (r *ServiceStatusTraffic) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceStatusTraffic) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceStatusAddress struct {
	empty bool    `json:"-"`
	Url   *string `json:"url"`
}

// This object is used to assert a desired state where this ServiceStatusAddress is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatusAddress *ServiceStatusAddress = &ServiceStatusAddress{empty: true}

func (r *ServiceStatusAddress) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceStatusAddress) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Service) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "run",
		Type:    "Service",
		Version: "run",
	}
}

const ServiceMaxPage = -1

type ServiceList struct {
	Items []*Service

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *ServiceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ServiceList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listService(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListService(ctx context.Context, project, location string) (*ServiceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListServiceWithMaxResults(ctx, project, location, ServiceMaxPage)

}

func (c *Client) ListServiceWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ServiceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listService(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ServiceList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) DeleteService(ctx context.Context, r *Service) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Service resource is nil")
	}
	c.Config.Logger.Info("Deleting Service...")
	deleteOp := deleteServiceOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllService deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllService(ctx context.Context, project, location string, filter func(*Service) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListService(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllService(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllService(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyService(ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (*Service, error) {
	c.Config.Logger.Info("Beginning ApplyService...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.serviceDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []serviceApiOperation
	if create {
		ops = append(ops, &createServiceOperation{})
	} else if recreate {

		ops = append(ops, &deleteServiceOperation{})

		ops = append(ops, &createServiceOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeServiceDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %#v", op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %#v: %v", op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %#v", op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetService(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createServiceOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapService(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeServiceNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeServiceNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeServiceDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffService(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
