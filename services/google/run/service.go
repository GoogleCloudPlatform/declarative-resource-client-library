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
	"encoding/json"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Service struct {
	Name       *string          `json:"name"`
	ApiVersion *string          `json:"apiVersion"`
	Kind       *string          `json:"kind"`
	Metadata   *ServiceMetadata `json:"metadata"`
	Spec       *ServiceSpec     `json:"spec"`
	Status     *ServiceStatus   `json:"status"`
	Project    *string          `json:"project"`
	Location   *string          `json:"location"`
}

func (r *Service) String() string {
	return dcl.SprintResource(r)
}

type ServiceMetadata struct {
	empty                      bool                             `json:"-"`
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

type jsonServiceMetadata ServiceMetadata

func (r *ServiceMetadata) UnmarshalJSON(data []byte) error {
	var res jsonServiceMetadata
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceMetadata
	} else {

		r.GenerateName = res.GenerateName

		r.Namespace = res.Namespace

		r.SelfLink = res.SelfLink

		r.Uid = res.Uid

		r.ResourceVersion = res.ResourceVersion

		r.Generation = res.Generation

		r.CreateTime = res.CreateTime

		r.Labels = res.Labels

		r.Annotations = res.Annotations

		r.OwnerReferences = res.OwnerReferences

		r.DeleteTime = res.DeleteTime

		r.DeletionGracePeriodSeconds = res.DeletionGracePeriodSeconds

		r.Finalizers = res.Finalizers

		r.ClusterName = res.ClusterName

	}
	return nil
}

// This object is used to assert a desired state where this ServiceMetadata is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceMetadata *ServiceMetadata = &ServiceMetadata{empty: true}

func (r *ServiceMetadata) Empty() bool {
	return r.empty
}

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

type jsonServiceMetadataCreateTime ServiceMetadataCreateTime

func (r *ServiceMetadataCreateTime) UnmarshalJSON(data []byte) error {
	var res jsonServiceMetadataCreateTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceMetadataCreateTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this ServiceMetadataCreateTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceMetadataCreateTime *ServiceMetadataCreateTime = &ServiceMetadataCreateTime{empty: true}

func (r *ServiceMetadataCreateTime) Empty() bool {
	return r.empty
}

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

type jsonServiceMetadataOwnerReferences ServiceMetadataOwnerReferences

func (r *ServiceMetadataOwnerReferences) UnmarshalJSON(data []byte) error {
	var res jsonServiceMetadataOwnerReferences
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceMetadataOwnerReferences
	} else {

		r.ApiVersion = res.ApiVersion

		r.Kind = res.Kind

		r.Name = res.Name

		r.Uid = res.Uid

		r.Controller = res.Controller

		r.BlockOwnerDeletion = res.BlockOwnerDeletion

	}
	return nil
}

// This object is used to assert a desired state where this ServiceMetadataOwnerReferences is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceMetadataOwnerReferences *ServiceMetadataOwnerReferences = &ServiceMetadataOwnerReferences{empty: true}

func (r *ServiceMetadataOwnerReferences) Empty() bool {
	return r.empty
}

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

type jsonServiceMetadataDeleteTime ServiceMetadataDeleteTime

func (r *ServiceMetadataDeleteTime) UnmarshalJSON(data []byte) error {
	var res jsonServiceMetadataDeleteTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceMetadataDeleteTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this ServiceMetadataDeleteTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceMetadataDeleteTime *ServiceMetadataDeleteTime = &ServiceMetadataDeleteTime{empty: true}

func (r *ServiceMetadataDeleteTime) Empty() bool {
	return r.empty
}

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

type jsonServiceSpec ServiceSpec

func (r *ServiceSpec) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpec
	} else {

		r.Template = res.Template

		r.Traffic = res.Traffic

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpec *ServiceSpec = &ServiceSpec{empty: true}

func (r *ServiceSpec) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplate ServiceSpecTemplate

func (r *ServiceSpecTemplate) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplate
	} else {

		r.Metadata = res.Metadata

		r.Spec = res.Spec

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplate *ServiceSpecTemplate = &ServiceSpecTemplate{empty: true}

func (r *ServiceSpecTemplate) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateMetadata ServiceSpecTemplateMetadata

func (r *ServiceSpecTemplateMetadata) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateMetadata
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateMetadata
	} else {

		r.Name = res.Name

		r.GenerateName = res.GenerateName

		r.Namespace = res.Namespace

		r.SelfLink = res.SelfLink

		r.Uid = res.Uid

		r.ResourceVersion = res.ResourceVersion

		r.Generation = res.Generation

		r.CreateTime = res.CreateTime

		r.Labels = res.Labels

		r.Annotations = res.Annotations

		r.OwnerReferences = res.OwnerReferences

		r.DeleteTime = res.DeleteTime

		r.DeletionGracePeriodSeconds = res.DeletionGracePeriodSeconds

		r.Finalizers = res.Finalizers

		r.ClusterName = res.ClusterName

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateMetadata is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateMetadata *ServiceSpecTemplateMetadata = &ServiceSpecTemplateMetadata{empty: true}

func (r *ServiceSpecTemplateMetadata) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateMetadataCreateTime ServiceSpecTemplateMetadataCreateTime

func (r *ServiceSpecTemplateMetadataCreateTime) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateMetadataCreateTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateMetadataCreateTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateMetadataCreateTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateMetadataCreateTime *ServiceSpecTemplateMetadataCreateTime = &ServiceSpecTemplateMetadataCreateTime{empty: true}

func (r *ServiceSpecTemplateMetadataCreateTime) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateMetadataOwnerReferences ServiceSpecTemplateMetadataOwnerReferences

func (r *ServiceSpecTemplateMetadataOwnerReferences) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateMetadataOwnerReferences
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateMetadataOwnerReferences
	} else {

		r.ApiVersion = res.ApiVersion

		r.Kind = res.Kind

		r.Name = res.Name

		r.Uid = res.Uid

		r.Controller = res.Controller

		r.BlockOwnerDeletion = res.BlockOwnerDeletion

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateMetadataOwnerReferences is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateMetadataOwnerReferences *ServiceSpecTemplateMetadataOwnerReferences = &ServiceSpecTemplateMetadataOwnerReferences{empty: true}

func (r *ServiceSpecTemplateMetadataOwnerReferences) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateMetadataDeleteTime ServiceSpecTemplateMetadataDeleteTime

func (r *ServiceSpecTemplateMetadataDeleteTime) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateMetadataDeleteTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateMetadataDeleteTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateMetadataDeleteTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateMetadataDeleteTime *ServiceSpecTemplateMetadataDeleteTime = &ServiceSpecTemplateMetadataDeleteTime{empty: true}

func (r *ServiceSpecTemplateMetadataDeleteTime) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpec ServiceSpecTemplateSpec

func (r *ServiceSpecTemplateSpec) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpec
	} else {

		r.ContainerConcurrency = res.ContainerConcurrency

		r.TimeoutSeconds = res.TimeoutSeconds

		r.ServiceAccountName = res.ServiceAccountName

		r.Containers = res.Containers

		r.Volumes = res.Volumes

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpec *ServiceSpecTemplateSpec = &ServiceSpecTemplateSpec{empty: true}

func (r *ServiceSpecTemplateSpec) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainers ServiceSpecTemplateSpecContainers

func (r *ServiceSpecTemplateSpecContainers) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainers
	} else {

		r.Name = res.Name

		r.Image = res.Image

		r.Command = res.Command

		r.Args = res.Args

		r.Env = res.Env

		r.Resources = res.Resources

		r.WorkingDir = res.WorkingDir

		r.Ports = res.Ports

		r.EnvFrom = res.EnvFrom

		r.VolumeMounts = res.VolumeMounts

		r.LivenessProbe = res.LivenessProbe

		r.ReadinessProbe = res.ReadinessProbe

		r.TerminationMessagePath = res.TerminationMessagePath

		r.TerminationMessagePolicy = res.TerminationMessagePolicy

		r.ImagePullPolicy = res.ImagePullPolicy

		r.SecurityContext = res.SecurityContext

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainers is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainers *ServiceSpecTemplateSpecContainers = &ServiceSpecTemplateSpecContainers{empty: true}

func (r *ServiceSpecTemplateSpecContainers) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnv ServiceSpecTemplateSpecContainersEnv

func (r *ServiceSpecTemplateSpecContainersEnv) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnv
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnv
	} else {

		r.Name = res.Name

		r.Value = res.Value

		r.ValueFrom = res.ValueFrom

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnv is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnv *ServiceSpecTemplateSpecContainersEnv = &ServiceSpecTemplateSpecContainersEnv{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnv) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvValueFrom ServiceSpecTemplateSpecContainersEnvValueFrom

func (r *ServiceSpecTemplateSpecContainersEnvValueFrom) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvValueFrom
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvValueFrom
	} else {

		r.ConfigMapKeyRef = res.ConfigMapKeyRef

		r.SecretKeyRef = res.SecretKeyRef

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFrom is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFrom *ServiceSpecTemplateSpecContainersEnvValueFrom = &ServiceSpecTemplateSpecContainersEnvValueFrom{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFrom) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef

func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef
	} else {

		r.LocalObjectReference = res.LocalObjectReference

		r.Key = res.Key

		r.Optional = res.Optional

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef = &ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference

func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference = &ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef

func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef
	} else {

		r.LocalObjectReference = res.LocalObjectReference

		r.Key = res.Key

		r.Optional = res.Optional

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef = &ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference

func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference = &ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersResources ServiceSpecTemplateSpecContainersResources

func (r *ServiceSpecTemplateSpecContainersResources) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersResources
	} else {

		r.Limits = res.Limits

		r.Requests = res.Requests

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersResources is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersResources *ServiceSpecTemplateSpecContainersResources = &ServiceSpecTemplateSpecContainersResources{empty: true}

func (r *ServiceSpecTemplateSpecContainersResources) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersPorts ServiceSpecTemplateSpecContainersPorts

func (r *ServiceSpecTemplateSpecContainersPorts) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersPorts
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersPorts
	} else {

		r.Name = res.Name

		r.ContainerPort = res.ContainerPort

		r.Protocol = res.Protocol

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersPorts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersPorts *ServiceSpecTemplateSpecContainersPorts = &ServiceSpecTemplateSpecContainersPorts{empty: true}

func (r *ServiceSpecTemplateSpecContainersPorts) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvFrom ServiceSpecTemplateSpecContainersEnvFrom

func (r *ServiceSpecTemplateSpecContainersEnvFrom) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvFrom
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvFrom
	} else {

		r.Prefix = res.Prefix

		r.ConfigMapRef = res.ConfigMapRef

		r.SecretRef = res.SecretRef

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFrom is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFrom *ServiceSpecTemplateSpecContainersEnvFrom = &ServiceSpecTemplateSpecContainersEnvFrom{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFrom) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvFromConfigMapRef ServiceSpecTemplateSpecContainersEnvFromConfigMapRef

func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvFromConfigMapRef
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRef
	} else {

		r.LocalObjectReference = res.LocalObjectReference

		r.Optional = res.Optional

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFromConfigMapRef is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRef *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef = &ServiceSpecTemplateSpecContainersEnvFromConfigMapRef{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRef) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference

func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference = &ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvFromSecretRef ServiceSpecTemplateSpecContainersEnvFromSecretRef

func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRef) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvFromSecretRef
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvFromSecretRef
	} else {

		r.LocalObjectReference = res.LocalObjectReference

		r.Optional = res.Optional

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFromSecretRef is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFromSecretRef *ServiceSpecTemplateSpecContainersEnvFromSecretRef = &ServiceSpecTemplateSpecContainersEnvFromSecretRef{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRef) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference

func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference = &ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference{empty: true}

func (r *ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersVolumeMounts ServiceSpecTemplateSpecContainersVolumeMounts

func (r *ServiceSpecTemplateSpecContainersVolumeMounts) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersVolumeMounts
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersVolumeMounts
	} else {

		r.Name = res.Name

		r.ReadOnly = res.ReadOnly

		r.MountPath = res.MountPath

		r.SubPath = res.SubPath

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersVolumeMounts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersVolumeMounts *ServiceSpecTemplateSpecContainersVolumeMounts = &ServiceSpecTemplateSpecContainersVolumeMounts{empty: true}

func (r *ServiceSpecTemplateSpecContainersVolumeMounts) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersLivenessProbe ServiceSpecTemplateSpecContainersLivenessProbe

func (r *ServiceSpecTemplateSpecContainersLivenessProbe) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersLivenessProbe
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersLivenessProbe
	} else {

		r.InitialDelaySeconds = res.InitialDelaySeconds

		r.TimeoutSeconds = res.TimeoutSeconds

		r.PeriodSeconds = res.PeriodSeconds

		r.SuccessThreshold = res.SuccessThreshold

		r.FailureThreshold = res.FailureThreshold

		r.Exec = res.Exec

		r.HttpGet = res.HttpGet

		r.TcpSocket = res.TcpSocket

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbe *ServiceSpecTemplateSpecContainersLivenessProbe = &ServiceSpecTemplateSpecContainersLivenessProbe{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbe) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersLivenessProbeExec ServiceSpecTemplateSpecContainersLivenessProbeExec

func (r *ServiceSpecTemplateSpecContainersLivenessProbeExec) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersLivenessProbeExec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersLivenessProbeExec
	} else {

		r.Command = res.Command

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbeExec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbeExec *ServiceSpecTemplateSpecContainersLivenessProbeExec = &ServiceSpecTemplateSpecContainersLivenessProbeExec{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeExec) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersLivenessProbeHttpGet ServiceSpecTemplateSpecContainersLivenessProbeHttpGet

func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersLivenessProbeHttpGet
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGet
	} else {

		r.Path = res.Path

		r.Host = res.Host

		r.Scheme = res.Scheme

		r.HttpHeaders = res.HttpHeaders

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbeHttpGet is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGet *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet = &ServiceSpecTemplateSpecContainersLivenessProbeHttpGet{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGet) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders

func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders
	} else {

		r.Name = res.Name

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders = &ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersLivenessProbeTcpSocket ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket

func (r *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersLivenessProbeTcpSocket
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersLivenessProbeTcpSocket
	} else {

		r.Port = res.Port

		r.Host = res.Host

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersLivenessProbeTcpSocket *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket = &ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket{empty: true}

func (r *ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersReadinessProbe ServiceSpecTemplateSpecContainersReadinessProbe

func (r *ServiceSpecTemplateSpecContainersReadinessProbe) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersReadinessProbe
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersReadinessProbe
	} else {

		r.InitialDelaySeconds = res.InitialDelaySeconds

		r.TimeoutSeconds = res.TimeoutSeconds

		r.PeriodSeconds = res.PeriodSeconds

		r.SuccessThreshold = res.SuccessThreshold

		r.FailureThreshold = res.FailureThreshold

		r.Exec = res.Exec

		r.HttpGet = res.HttpGet

		r.TcpSocket = res.TcpSocket

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbe *ServiceSpecTemplateSpecContainersReadinessProbe = &ServiceSpecTemplateSpecContainersReadinessProbe{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbe) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersReadinessProbeExec ServiceSpecTemplateSpecContainersReadinessProbeExec

func (r *ServiceSpecTemplateSpecContainersReadinessProbeExec) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersReadinessProbeExec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersReadinessProbeExec
	} else {

		r.Command = res.Command

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbeExec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbeExec *ServiceSpecTemplateSpecContainersReadinessProbeExec = &ServiceSpecTemplateSpecContainersReadinessProbeExec{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeExec) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersReadinessProbeHttpGet ServiceSpecTemplateSpecContainersReadinessProbeHttpGet

func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersReadinessProbeHttpGet
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGet
	} else {

		r.Path = res.Path

		r.Host = res.Host

		r.Scheme = res.Scheme

		r.HttpHeaders = res.HttpHeaders

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbeHttpGet is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGet *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet = &ServiceSpecTemplateSpecContainersReadinessProbeHttpGet{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGet) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders

func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders
	} else {

		r.Name = res.Name

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders = &ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersReadinessProbeTcpSocket ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket

func (r *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersReadinessProbeTcpSocket
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersReadinessProbeTcpSocket
	} else {

		r.Port = res.Port

		r.Host = res.Host

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersReadinessProbeTcpSocket *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket = &ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket{empty: true}

func (r *ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecContainersSecurityContext ServiceSpecTemplateSpecContainersSecurityContext

func (r *ServiceSpecTemplateSpecContainersSecurityContext) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecContainersSecurityContext
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecContainersSecurityContext
	} else {

		r.RunAsUser = res.RunAsUser

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecContainersSecurityContext is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecContainersSecurityContext *ServiceSpecTemplateSpecContainersSecurityContext = &ServiceSpecTemplateSpecContainersSecurityContext{empty: true}

func (r *ServiceSpecTemplateSpecContainersSecurityContext) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecVolumes ServiceSpecTemplateSpecVolumes

func (r *ServiceSpecTemplateSpecVolumes) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecVolumes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecVolumes
	} else {

		r.Name = res.Name

		r.Secret = res.Secret

		r.ConfigMap = res.ConfigMap

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumes *ServiceSpecTemplateSpecVolumes = &ServiceSpecTemplateSpecVolumes{empty: true}

func (r *ServiceSpecTemplateSpecVolumes) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecVolumesSecret ServiceSpecTemplateSpecVolumesSecret

func (r *ServiceSpecTemplateSpecVolumesSecret) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecVolumesSecret
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecVolumesSecret
	} else {

		r.SecretName = res.SecretName

		r.Items = res.Items

		r.DefaultMode = res.DefaultMode

		r.Optional = res.Optional

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumesSecret is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumesSecret *ServiceSpecTemplateSpecVolumesSecret = &ServiceSpecTemplateSpecVolumesSecret{empty: true}

func (r *ServiceSpecTemplateSpecVolumesSecret) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecVolumesSecretItems ServiceSpecTemplateSpecVolumesSecretItems

func (r *ServiceSpecTemplateSpecVolumesSecretItems) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecVolumesSecretItems
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecVolumesSecretItems
	} else {

		r.Key = res.Key

		r.Path = res.Path

		r.Mode = res.Mode

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumesSecretItems is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumesSecretItems *ServiceSpecTemplateSpecVolumesSecretItems = &ServiceSpecTemplateSpecVolumesSecretItems{empty: true}

func (r *ServiceSpecTemplateSpecVolumesSecretItems) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecVolumesConfigMap ServiceSpecTemplateSpecVolumesConfigMap

func (r *ServiceSpecTemplateSpecVolumesConfigMap) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecVolumesConfigMap
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecVolumesConfigMap
	} else {

		r.Name = res.Name

		r.Items = res.Items

		r.DefaultMode = res.DefaultMode

		r.Optional = res.Optional

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumesConfigMap is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumesConfigMap *ServiceSpecTemplateSpecVolumesConfigMap = &ServiceSpecTemplateSpecVolumesConfigMap{empty: true}

func (r *ServiceSpecTemplateSpecVolumesConfigMap) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTemplateSpecVolumesConfigMapItems ServiceSpecTemplateSpecVolumesConfigMapItems

func (r *ServiceSpecTemplateSpecVolumesConfigMapItems) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTemplateSpecVolumesConfigMapItems
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTemplateSpecVolumesConfigMapItems
	} else {

		r.Key = res.Key

		r.Path = res.Path

		r.Mode = res.Mode

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTemplateSpecVolumesConfigMapItems is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTemplateSpecVolumesConfigMapItems *ServiceSpecTemplateSpecVolumesConfigMapItems = &ServiceSpecTemplateSpecVolumesConfigMapItems{empty: true}

func (r *ServiceSpecTemplateSpecVolumesConfigMapItems) Empty() bool {
	return r.empty
}

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

type jsonServiceSpecTraffic ServiceSpecTraffic

func (r *ServiceSpecTraffic) UnmarshalJSON(data []byte) error {
	var res jsonServiceSpecTraffic
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceSpecTraffic
	} else {

		r.ConfigurationName = res.ConfigurationName

		r.RevisionName = res.RevisionName

		r.Percent = res.Percent

		r.Tag = res.Tag

		r.LatestRevision = res.LatestRevision

		r.Url = res.Url

	}
	return nil
}

// This object is used to assert a desired state where this ServiceSpecTraffic is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceSpecTraffic *ServiceSpecTraffic = &ServiceSpecTraffic{empty: true}

func (r *ServiceSpecTraffic) Empty() bool {
	return r.empty
}

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

type jsonServiceStatus ServiceStatus

func (r *ServiceStatus) UnmarshalJSON(data []byte) error {
	var res jsonServiceStatus
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceStatus
	} else {

		r.ObservedGeneration = res.ObservedGeneration

		r.Conditions = res.Conditions

		r.LatestReadyRevisionName = res.LatestReadyRevisionName

		r.LatestCreatedRevisionName = res.LatestCreatedRevisionName

		r.Traffic = res.Traffic

		r.Url = res.Url

		r.Address = res.Address

	}
	return nil
}

// This object is used to assert a desired state where this ServiceStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatus *ServiceStatus = &ServiceStatus{empty: true}

func (r *ServiceStatus) Empty() bool {
	return r.empty
}

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

type jsonServiceStatusConditions ServiceStatusConditions

func (r *ServiceStatusConditions) UnmarshalJSON(data []byte) error {
	var res jsonServiceStatusConditions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceStatusConditions
	} else {

		r.Type = res.Type

		r.Status = res.Status

		r.Reason = res.Reason

		r.Message = res.Message

		r.LastTransitionTime = res.LastTransitionTime

		r.Severity = res.Severity

	}
	return nil
}

// This object is used to assert a desired state where this ServiceStatusConditions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatusConditions *ServiceStatusConditions = &ServiceStatusConditions{empty: true}

func (r *ServiceStatusConditions) Empty() bool {
	return r.empty
}

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

type jsonServiceStatusConditionsLastTransitionTime ServiceStatusConditionsLastTransitionTime

func (r *ServiceStatusConditionsLastTransitionTime) UnmarshalJSON(data []byte) error {
	var res jsonServiceStatusConditionsLastTransitionTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceStatusConditionsLastTransitionTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this ServiceStatusConditionsLastTransitionTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatusConditionsLastTransitionTime *ServiceStatusConditionsLastTransitionTime = &ServiceStatusConditionsLastTransitionTime{empty: true}

func (r *ServiceStatusConditionsLastTransitionTime) Empty() bool {
	return r.empty
}

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

type jsonServiceStatusTraffic ServiceStatusTraffic

func (r *ServiceStatusTraffic) UnmarshalJSON(data []byte) error {
	var res jsonServiceStatusTraffic
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceStatusTraffic
	} else {

		r.ConfigurationName = res.ConfigurationName

		r.RevisionName = res.RevisionName

		r.Percent = res.Percent

		r.Tag = res.Tag

		r.LatestRevision = res.LatestRevision

		r.Url = res.Url

	}
	return nil
}

// This object is used to assert a desired state where this ServiceStatusTraffic is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatusTraffic *ServiceStatusTraffic = &ServiceStatusTraffic{empty: true}

func (r *ServiceStatusTraffic) Empty() bool {
	return r.empty
}

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

type jsonServiceStatusAddress ServiceStatusAddress

func (r *ServiceStatusAddress) UnmarshalJSON(data []byte) error {
	var res jsonServiceStatusAddress
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceStatusAddress
	} else {

		r.Url = res.Url

	}
	return nil
}

// This object is used to assert a desired state where this ServiceStatusAddress is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceStatusAddress *ServiceStatusAddress = &ServiceStatusAddress{empty: true}

func (r *ServiceStatusAddress) Empty() bool {
	return r.empty
}

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
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
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
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListServiceWithMaxResults(ctx, project, location, ServiceMaxPage)

}

func (c *Client) ListServiceWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ServiceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
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
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
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

	var resultNewState *Service
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyServiceHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyServiceHelper(c *Client, ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (*Service, error) {
	c.Config.Logger.Info("Beginning ApplyService...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.serviceDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	for _, fd := range fieldDiffs {
		fmt.Printf("fd: %+v\n", fd)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToServiceOp(opStrings, fieldDiffs, opts)
	fmt.Printf("diffs: %+v, opStrings: %v\n", diffs, opStrings)
	if err != nil {
		return nil, err
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
