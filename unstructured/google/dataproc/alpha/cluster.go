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
package dataproc

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type Cluster struct{}

func ClusterToUnstructured(r *dclService.Cluster) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dataproc",
			Version: "alpha",
			Type:    "Cluster",
		},
		Object: make(map[string]interface{}),
	}
	if r.ClusterUuid != nil {
		u.Object["clusterUuid"] = *r.ClusterUuid
	}
	if r.Config != nil {
		u.Object["config"] = ClusterClusterConfigToUnstructured(r.Config)
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Metrics != nil && r.Metrics != dclService.EmptyClusterMetrics {
		rMetrics := make(map[string]interface{})
		if r.Metrics.HdfsMetrics != nil {
			rMetricsHdfsMetrics := make(map[string]interface{})
			for k, v := range r.Metrics.HdfsMetrics {
				rMetricsHdfsMetrics[k] = v
			}
			rMetrics["hdfsMetrics"] = rMetricsHdfsMetrics
		}
		if r.Metrics.YarnMetrics != nil {
			rMetricsYarnMetrics := make(map[string]interface{})
			for k, v := range r.Metrics.YarnMetrics {
				rMetricsYarnMetrics[k] = v
			}
			rMetrics["yarnMetrics"] = rMetricsYarnMetrics
		}
		u.Object["metrics"] = rMetrics
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Status != nil && r.Status != dclService.EmptyClusterStatus {
		rStatus := make(map[string]interface{})
		if r.Status.Detail != nil {
			rStatus["detail"] = *r.Status.Detail
		}
		if r.Status.State != nil {
			rStatus["state"] = string(*r.Status.State)
		}
		if r.Status.StateStartTime != nil {
			rStatus["stateStartTime"] = *r.Status.StateStartTime
		}
		if r.Status.Substate != nil {
			rStatus["substate"] = string(*r.Status.Substate)
		}
		u.Object["status"] = rStatus
	}
	var rStatusHistory []interface{}
	for _, rStatusHistoryVal := range r.StatusHistory {
		rStatusHistoryObject := make(map[string]interface{})
		if rStatusHistoryVal.Detail != nil {
			rStatusHistoryObject["detail"] = *rStatusHistoryVal.Detail
		}
		if rStatusHistoryVal.State != nil {
			rStatusHistoryObject["state"] = string(*rStatusHistoryVal.State)
		}
		if rStatusHistoryVal.StateStartTime != nil {
			rStatusHistoryObject["stateStartTime"] = *rStatusHistoryVal.StateStartTime
		}
		if rStatusHistoryVal.Substate != nil {
			rStatusHistoryObject["substate"] = string(*rStatusHistoryVal.Substate)
		}
		rStatusHistory = append(rStatusHistory, rStatusHistoryObject)
	}
	u.Object["statusHistory"] = rStatusHistory
	return u
}

func ClusterInstanceGroupConfigToUnstructured(r *dclService.ClusterInstanceGroupConfig) map[string]interface{} {
	result := make(map[string]interface{})
	var rAccelerators []interface{}
	for _, rAcceleratorsVal := range r.Accelerators {
		rAcceleratorsObject := make(map[string]interface{})
		if rAcceleratorsVal.AcceleratorCount != nil {
			rAcceleratorsObject["acceleratorCount"] = *rAcceleratorsVal.AcceleratorCount
		}
		if rAcceleratorsVal.AcceleratorType != nil {
			rAcceleratorsObject["acceleratorType"] = *rAcceleratorsVal.AcceleratorType
		}
		rAccelerators = append(rAccelerators, rAcceleratorsObject)
	}
	result["accelerators"] = rAccelerators
	if r.DiskConfig != nil && r.DiskConfig != dclService.EmptyClusterInstanceGroupConfigDiskConfig {
		rDiskConfig := make(map[string]interface{})
		if r.DiskConfig.BootDiskSizeGb != nil {
			rDiskConfig["bootDiskSizeGb"] = *r.DiskConfig.BootDiskSizeGb
		}
		if r.DiskConfig.BootDiskType != nil {
			rDiskConfig["bootDiskType"] = *r.DiskConfig.BootDiskType
		}
		if r.DiskConfig.NumLocalSsds != nil {
			rDiskConfig["numLocalSsds"] = *r.DiskConfig.NumLocalSsds
		}
		result["diskConfig"] = rDiskConfig
	}
	if r.Image != nil {
		result["image"] = *r.Image
	}
	var rInstanceNames []interface{}
	for _, rInstanceNamesVal := range r.InstanceNames {
		rInstanceNames = append(rInstanceNames, rInstanceNamesVal)
	}
	result["instanceNames"] = rInstanceNames
	if r.IsPreemptible != nil {
		result["isPreemptible"] = *r.IsPreemptible
	}
	if r.MachineType != nil {
		result["machineType"] = *r.MachineType
	}
	if r.ManagedGroupConfig != nil && r.ManagedGroupConfig != dclService.EmptyClusterInstanceGroupConfigManagedGroupConfig {
		rManagedGroupConfig := make(map[string]interface{})
		if r.ManagedGroupConfig.InstanceGroupManagerName != nil {
			rManagedGroupConfig["instanceGroupManagerName"] = *r.ManagedGroupConfig.InstanceGroupManagerName
		}
		if r.ManagedGroupConfig.InstanceTemplateName != nil {
			rManagedGroupConfig["instanceTemplateName"] = *r.ManagedGroupConfig.InstanceTemplateName
		}
		result["managedGroupConfig"] = rManagedGroupConfig
	}
	if r.MinCpuPlatform != nil {
		result["minCpuPlatform"] = *r.MinCpuPlatform
	}
	if r.NumInstances != nil {
		result["numInstances"] = *r.NumInstances
	}
	if r.Preemptibility != nil {
		result["preemptibility"] = string(*r.Preemptibility)
	}
	return result
}

func ClusterClusterConfigToUnstructured(r *dclService.ClusterClusterConfig) map[string]interface{} {
	result := make(map[string]interface{})
	if r.AutoscalingConfig != nil && r.AutoscalingConfig != dclService.EmptyClusterClusterConfigAutoscalingConfig {
		rAutoscalingConfig := make(map[string]interface{})
		if r.AutoscalingConfig.Policy != nil {
			rAutoscalingConfig["policy"] = *r.AutoscalingConfig.Policy
		}
		result["autoscalingConfig"] = rAutoscalingConfig
	}
	if r.EncryptionConfig != nil && r.EncryptionConfig != dclService.EmptyClusterClusterConfigEncryptionConfig {
		rEncryptionConfig := make(map[string]interface{})
		if r.EncryptionConfig.GcePdKmsKeyName != nil {
			rEncryptionConfig["gcePdKmsKeyName"] = *r.EncryptionConfig.GcePdKmsKeyName
		}
		result["encryptionConfig"] = rEncryptionConfig
	}
	if r.EndpointConfig != nil && r.EndpointConfig != dclService.EmptyClusterClusterConfigEndpointConfig {
		rEndpointConfig := make(map[string]interface{})
		if r.EndpointConfig.EnableHttpPortAccess != nil {
			rEndpointConfig["enableHttpPortAccess"] = *r.EndpointConfig.EnableHttpPortAccess
		}
		if r.EndpointConfig.HttpPorts != nil {
			rEndpointConfigHttpPorts := make(map[string]interface{})
			for k, v := range r.EndpointConfig.HttpPorts {
				rEndpointConfigHttpPorts[k] = v
			}
			rEndpointConfig["httpPorts"] = rEndpointConfigHttpPorts
		}
		result["endpointConfig"] = rEndpointConfig
	}
	if r.GceClusterConfig != nil && r.GceClusterConfig != dclService.EmptyClusterClusterConfigGceClusterConfig {
		rGceClusterConfig := make(map[string]interface{})
		if r.GceClusterConfig.InternalIPOnly != nil {
			rGceClusterConfig["internalIPOnly"] = *r.GceClusterConfig.InternalIPOnly
		}
		if r.GceClusterConfig.Metadata != nil {
			rGceClusterConfigMetadata := make(map[string]interface{})
			for k, v := range r.GceClusterConfig.Metadata {
				rGceClusterConfigMetadata[k] = v
			}
			rGceClusterConfig["metadata"] = rGceClusterConfigMetadata
		}
		if r.GceClusterConfig.Network != nil {
			rGceClusterConfig["network"] = *r.GceClusterConfig.Network
		}
		if r.GceClusterConfig.NodeGroupAffinity != nil && r.GceClusterConfig.NodeGroupAffinity != dclService.EmptyClusterClusterConfigGceClusterConfigNodeGroupAffinity {
			rGceClusterConfigNodeGroupAffinity := make(map[string]interface{})
			if r.GceClusterConfig.NodeGroupAffinity.NodeGroup != nil {
				rGceClusterConfigNodeGroupAffinity["nodeGroup"] = *r.GceClusterConfig.NodeGroupAffinity.NodeGroup
			}
			rGceClusterConfig["nodeGroupAffinity"] = rGceClusterConfigNodeGroupAffinity
		}
		if r.GceClusterConfig.PrivateIPv6GoogleAccess != nil {
			rGceClusterConfig["privateIPv6GoogleAccess"] = string(*r.GceClusterConfig.PrivateIPv6GoogleAccess)
		}
		if r.GceClusterConfig.ReservationAffinity != nil && r.GceClusterConfig.ReservationAffinity != dclService.EmptyClusterClusterConfigGceClusterConfigReservationAffinity {
			rGceClusterConfigReservationAffinity := make(map[string]interface{})
			if r.GceClusterConfig.ReservationAffinity.ConsumeReservationType != nil {
				rGceClusterConfigReservationAffinity["consumeReservationType"] = string(*r.GceClusterConfig.ReservationAffinity.ConsumeReservationType)
			}
			if r.GceClusterConfig.ReservationAffinity.Key != nil {
				rGceClusterConfigReservationAffinity["key"] = *r.GceClusterConfig.ReservationAffinity.Key
			}
			var rGceClusterConfigReservationAffinityValues []interface{}
			for _, rGceClusterConfigReservationAffinityValuesVal := range r.GceClusterConfig.ReservationAffinity.Values {
				rGceClusterConfigReservationAffinityValues = append(rGceClusterConfigReservationAffinityValues, rGceClusterConfigReservationAffinityValuesVal)
			}
			rGceClusterConfigReservationAffinity["values"] = rGceClusterConfigReservationAffinityValues
			rGceClusterConfig["reservationAffinity"] = rGceClusterConfigReservationAffinity
		}
		if r.GceClusterConfig.ServiceAccount != nil {
			rGceClusterConfig["serviceAccount"] = *r.GceClusterConfig.ServiceAccount
		}
		var rGceClusterConfigServiceAccountScopes []interface{}
		for _, rGceClusterConfigServiceAccountScopesVal := range r.GceClusterConfig.ServiceAccountScopes {
			rGceClusterConfigServiceAccountScopes = append(rGceClusterConfigServiceAccountScopes, rGceClusterConfigServiceAccountScopesVal)
		}
		rGceClusterConfig["serviceAccountScopes"] = rGceClusterConfigServiceAccountScopes
		if r.GceClusterConfig.Subnetwork != nil {
			rGceClusterConfig["subnetwork"] = *r.GceClusterConfig.Subnetwork
		}
		var rGceClusterConfigTags []interface{}
		for _, rGceClusterConfigTagsVal := range r.GceClusterConfig.Tags {
			rGceClusterConfigTags = append(rGceClusterConfigTags, rGceClusterConfigTagsVal)
		}
		rGceClusterConfig["tags"] = rGceClusterConfigTags
		if r.GceClusterConfig.Zone != nil {
			rGceClusterConfig["zone"] = *r.GceClusterConfig.Zone
		}
		result["gceClusterConfig"] = rGceClusterConfig
	}
	if r.GkeClusterConfig != nil && r.GkeClusterConfig != dclService.EmptyClusterClusterConfigGkeClusterConfig {
		rGkeClusterConfig := make(map[string]interface{})
		if r.GkeClusterConfig.NamespacedGkeDeploymentTarget != nil && r.GkeClusterConfig.NamespacedGkeDeploymentTarget != dclService.EmptyClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
			rGkeClusterConfigNamespacedGkeDeploymentTarget := make(map[string]interface{})
			if r.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace != nil {
				rGkeClusterConfigNamespacedGkeDeploymentTarget["clusterNamespace"] = *r.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace
			}
			if r.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster != nil {
				rGkeClusterConfigNamespacedGkeDeploymentTarget["targetGkeCluster"] = *r.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster
			}
			rGkeClusterConfig["namespacedGkeDeploymentTarget"] = rGkeClusterConfigNamespacedGkeDeploymentTarget
		}
		result["gkeClusterConfig"] = rGkeClusterConfig
	}
	var rInitializationActions []interface{}
	for _, rInitializationActionsVal := range r.InitializationActions {
		rInitializationActionsObject := make(map[string]interface{})
		if rInitializationActionsVal.ExecutableFile != nil {
			rInitializationActionsObject["executableFile"] = *rInitializationActionsVal.ExecutableFile
		}
		if rInitializationActionsVal.ExecutionTimeout != nil {
			rInitializationActionsObject["executionTimeout"] = *rInitializationActionsVal.ExecutionTimeout
		}
		rInitializationActions = append(rInitializationActions, rInitializationActionsObject)
	}
	result["initializationActions"] = rInitializationActions
	if r.LifecycleConfig != nil && r.LifecycleConfig != dclService.EmptyClusterClusterConfigLifecycleConfig {
		rLifecycleConfig := make(map[string]interface{})
		if r.LifecycleConfig.AutoDeleteTime != nil {
			rLifecycleConfig["autoDeleteTime"] = *r.LifecycleConfig.AutoDeleteTime
		}
		if r.LifecycleConfig.AutoDeleteTtl != nil {
			rLifecycleConfig["autoDeleteTtl"] = *r.LifecycleConfig.AutoDeleteTtl
		}
		if r.LifecycleConfig.IdleDeleteTtl != nil {
			rLifecycleConfig["idleDeleteTtl"] = *r.LifecycleConfig.IdleDeleteTtl
		}
		if r.LifecycleConfig.IdleStartTime != nil {
			rLifecycleConfig["idleStartTime"] = *r.LifecycleConfig.IdleStartTime
		}
		result["lifecycleConfig"] = rLifecycleConfig
	}
	if r.MasterConfig != nil {
		result["masterConfig"] = ClusterInstanceGroupConfigToUnstructured(r.MasterConfig)
	}
	if r.MetastoreConfig != nil && r.MetastoreConfig != dclService.EmptyClusterClusterConfigMetastoreConfig {
		rMetastoreConfig := make(map[string]interface{})
		if r.MetastoreConfig.DataprocMetastoreService != nil {
			rMetastoreConfig["dataprocMetastoreService"] = *r.MetastoreConfig.DataprocMetastoreService
		}
		result["metastoreConfig"] = rMetastoreConfig
	}
	if r.SecondaryWorkerConfig != nil {
		result["secondaryWorkerConfig"] = ClusterInstanceGroupConfigToUnstructured(r.SecondaryWorkerConfig)
	}
	if r.SecurityConfig != nil && r.SecurityConfig != dclService.EmptyClusterClusterConfigSecurityConfig {
		rSecurityConfig := make(map[string]interface{})
		if r.SecurityConfig.KerberosConfig != nil && r.SecurityConfig.KerberosConfig != dclService.EmptyClusterClusterConfigSecurityConfigKerberosConfig {
			rSecurityConfigKerberosConfig := make(map[string]interface{})
			if r.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer != nil {
				rSecurityConfigKerberosConfig["crossRealmTrustAdminServer"] = *r.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer
			}
			if r.SecurityConfig.KerberosConfig.CrossRealmTrustKdc != nil {
				rSecurityConfigKerberosConfig["crossRealmTrustKdc"] = *r.SecurityConfig.KerberosConfig.CrossRealmTrustKdc
			}
			if r.SecurityConfig.KerberosConfig.CrossRealmTrustRealm != nil {
				rSecurityConfigKerberosConfig["crossRealmTrustRealm"] = *r.SecurityConfig.KerberosConfig.CrossRealmTrustRealm
			}
			if r.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword != nil {
				rSecurityConfigKerberosConfig["crossRealmTrustSharedPassword"] = *r.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword
			}
			if r.SecurityConfig.KerberosConfig.EnableKerberos != nil {
				rSecurityConfigKerberosConfig["enableKerberos"] = *r.SecurityConfig.KerberosConfig.EnableKerberos
			}
			if r.SecurityConfig.KerberosConfig.KdcDbKey != nil {
				rSecurityConfigKerberosConfig["kdcDbKey"] = *r.SecurityConfig.KerberosConfig.KdcDbKey
			}
			if r.SecurityConfig.KerberosConfig.KeyPassword != nil {
				rSecurityConfigKerberosConfig["keyPassword"] = *r.SecurityConfig.KerberosConfig.KeyPassword
			}
			if r.SecurityConfig.KerberosConfig.Keystore != nil {
				rSecurityConfigKerberosConfig["keystore"] = *r.SecurityConfig.KerberosConfig.Keystore
			}
			if r.SecurityConfig.KerberosConfig.KeystorePassword != nil {
				rSecurityConfigKerberosConfig["keystorePassword"] = *r.SecurityConfig.KerberosConfig.KeystorePassword
			}
			if r.SecurityConfig.KerberosConfig.KmsKey != nil {
				rSecurityConfigKerberosConfig["kmsKey"] = *r.SecurityConfig.KerberosConfig.KmsKey
			}
			if r.SecurityConfig.KerberosConfig.Realm != nil {
				rSecurityConfigKerberosConfig["realm"] = *r.SecurityConfig.KerberosConfig.Realm
			}
			if r.SecurityConfig.KerberosConfig.RootPrincipalPassword != nil {
				rSecurityConfigKerberosConfig["rootPrincipalPassword"] = *r.SecurityConfig.KerberosConfig.RootPrincipalPassword
			}
			if r.SecurityConfig.KerberosConfig.TgtLifetimeHours != nil {
				rSecurityConfigKerberosConfig["tgtLifetimeHours"] = *r.SecurityConfig.KerberosConfig.TgtLifetimeHours
			}
			if r.SecurityConfig.KerberosConfig.Truststore != nil {
				rSecurityConfigKerberosConfig["truststore"] = *r.SecurityConfig.KerberosConfig.Truststore
			}
			if r.SecurityConfig.KerberosConfig.TruststorePassword != nil {
				rSecurityConfigKerberosConfig["truststorePassword"] = *r.SecurityConfig.KerberosConfig.TruststorePassword
			}
			rSecurityConfig["kerberosConfig"] = rSecurityConfigKerberosConfig
		}
		result["securityConfig"] = rSecurityConfig
	}
	if r.SoftwareConfig != nil && r.SoftwareConfig != dclService.EmptyClusterClusterConfigSoftwareConfig {
		rSoftwareConfig := make(map[string]interface{})
		if r.SoftwareConfig.ImageVersion != nil {
			rSoftwareConfig["imageVersion"] = *r.SoftwareConfig.ImageVersion
		}
		var rSoftwareConfigOptionalComponents []interface{}
		for _, rSoftwareConfigOptionalComponentsVal := range r.SoftwareConfig.OptionalComponents {
			rSoftwareConfigOptionalComponents = append(rSoftwareConfigOptionalComponents, string(rSoftwareConfigOptionalComponentsVal))
		}
		rSoftwareConfig["optionalComponents"] = rSoftwareConfigOptionalComponents
		if r.SoftwareConfig.Properties != nil {
			rSoftwareConfigProperties := make(map[string]interface{})
			for k, v := range r.SoftwareConfig.Properties {
				rSoftwareConfigProperties[k] = v
			}
			rSoftwareConfig["properties"] = rSoftwareConfigProperties
		}
		result["softwareConfig"] = rSoftwareConfig
	}
	if r.StagingBucket != nil {
		result["stagingBucket"] = *r.StagingBucket
	}
	if r.TempBucket != nil {
		result["tempBucket"] = *r.TempBucket
	}
	if r.WorkerConfig != nil {
		result["workerConfig"] = ClusterInstanceGroupConfigToUnstructured(r.WorkerConfig)
	}
	return result
}

func UnstructuredToCluster(u *unstructured.Resource) (*dclService.Cluster, error) {
	r := &dclService.Cluster{}
	if _, ok := u.Object["clusterUuid"]; ok {
		if s, ok := u.Object["clusterUuid"].(string); ok {
			r.ClusterUuid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ClusterUuid: expected string")
		}
	}
	if _, ok := u.Object["config"]; ok {
		if rConfig, ok := u.Object["config"].(map[string]interface{}); ok {
			var err error
			r.Config, err = UnstructuredToClusterClusterConfig(rConfig)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("r.Config: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["metrics"]; ok {
		if rMetrics, ok := u.Object["metrics"].(map[string]interface{}); ok {
			r.Metrics = &dclService.ClusterMetrics{}
			if _, ok := rMetrics["hdfsMetrics"]; ok {
				if rMetricsHdfsMetrics, ok := rMetrics["hdfsMetrics"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rMetricsHdfsMetrics {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.Metrics.HdfsMetrics = m
				} else {
					return nil, fmt.Errorf("r.Metrics.HdfsMetrics: expected map[string]interface{}")
				}
			}
			if _, ok := rMetrics["yarnMetrics"]; ok {
				if rMetricsYarnMetrics, ok := rMetrics["yarnMetrics"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rMetricsYarnMetrics {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.Metrics.YarnMetrics = m
				} else {
					return nil, fmt.Errorf("r.Metrics.YarnMetrics: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Metrics: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["status"]; ok {
		if rStatus, ok := u.Object["status"].(map[string]interface{}); ok {
			r.Status = &dclService.ClusterStatus{}
			if _, ok := rStatus["detail"]; ok {
				if s, ok := rStatus["detail"].(string); ok {
					r.Status.Detail = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Status.Detail: expected string")
				}
			}
			if _, ok := rStatus["state"]; ok {
				if s, ok := rStatus["state"].(string); ok {
					r.Status.State = dclService.ClusterStatusStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Status.State: expected string")
				}
			}
			if _, ok := rStatus["stateStartTime"]; ok {
				if s, ok := rStatus["stateStartTime"].(string); ok {
					r.Status.StateStartTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Status.StateStartTime: expected string")
				}
			}
			if _, ok := rStatus["substate"]; ok {
				if s, ok := rStatus["substate"].(string); ok {
					r.Status.Substate = dclService.ClusterStatusSubstateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Status.Substate: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Status: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["statusHistory"]; ok {
		if s, ok := u.Object["statusHistory"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rStatusHistory dclService.ClusterStatusHistory
					if _, ok := objval["detail"]; ok {
						if s, ok := objval["detail"].(string); ok {
							rStatusHistory.Detail = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rStatusHistory.Detail: expected string")
						}
					}
					if _, ok := objval["state"]; ok {
						if s, ok := objval["state"].(string); ok {
							rStatusHistory.State = dclService.ClusterStatusHistoryStateEnumRef(s)
						} else {
							return nil, fmt.Errorf("rStatusHistory.State: expected string")
						}
					}
					if _, ok := objval["stateStartTime"]; ok {
						if s, ok := objval["stateStartTime"].(string); ok {
							rStatusHistory.StateStartTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rStatusHistory.StateStartTime: expected string")
						}
					}
					if _, ok := objval["substate"]; ok {
						if s, ok := objval["substate"].(string); ok {
							rStatusHistory.Substate = dclService.ClusterStatusHistorySubstateEnumRef(s)
						} else {
							return nil, fmt.Errorf("rStatusHistory.Substate: expected string")
						}
					}
					r.StatusHistory = append(r.StatusHistory, rStatusHistory)
				}
			}
		} else {
			return nil, fmt.Errorf("r.StatusHistory: expected []interface{}")
		}
	}
	return r, nil
}

func UnstructuredToClusterClusterConfig(obj map[string]interface{}) (*dclService.ClusterClusterConfig, error) {
	r := &dclService.ClusterClusterConfig{}
	if _, ok := obj["autoscalingConfig"]; ok {
		if rAutoscalingConfig, ok := obj["autoscalingConfig"].(map[string]interface{}); ok {
			r.AutoscalingConfig = &dclService.ClusterClusterConfigAutoscalingConfig{}
			if _, ok := rAutoscalingConfig["policy"]; ok {
				if s, ok := rAutoscalingConfig["policy"].(string); ok {
					r.AutoscalingConfig.Policy = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AutoscalingConfig.Policy: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AutoscalingConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["encryptionConfig"]; ok {
		if rEncryptionConfig, ok := obj["encryptionConfig"].(map[string]interface{}); ok {
			r.EncryptionConfig = &dclService.ClusterClusterConfigEncryptionConfig{}
			if _, ok := rEncryptionConfig["gcePdKmsKeyName"]; ok {
				if s, ok := rEncryptionConfig["gcePdKmsKeyName"].(string); ok {
					r.EncryptionConfig.GcePdKmsKeyName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.EncryptionConfig.GcePdKmsKeyName: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.EncryptionConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["endpointConfig"]; ok {
		if rEndpointConfig, ok := obj["endpointConfig"].(map[string]interface{}); ok {
			r.EndpointConfig = &dclService.ClusterClusterConfigEndpointConfig{}
			if _, ok := rEndpointConfig["enableHttpPortAccess"]; ok {
				if b, ok := rEndpointConfig["enableHttpPortAccess"].(bool); ok {
					r.EndpointConfig.EnableHttpPortAccess = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.EndpointConfig.EnableHttpPortAccess: expected bool")
				}
			}
			if _, ok := rEndpointConfig["httpPorts"]; ok {
				if rEndpointConfigHttpPorts, ok := rEndpointConfig["httpPorts"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rEndpointConfigHttpPorts {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.EndpointConfig.HttpPorts = m
				} else {
					return nil, fmt.Errorf("r.EndpointConfig.HttpPorts: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.EndpointConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["gceClusterConfig"]; ok {
		if rGceClusterConfig, ok := obj["gceClusterConfig"].(map[string]interface{}); ok {
			r.GceClusterConfig = &dclService.ClusterClusterConfigGceClusterConfig{}
			if _, ok := rGceClusterConfig["internalIPOnly"]; ok {
				if b, ok := rGceClusterConfig["internalIPOnly"].(bool); ok {
					r.GceClusterConfig.InternalIPOnly = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.InternalIPOnly: expected bool")
				}
			}
			if _, ok := rGceClusterConfig["metadata"]; ok {
				if rGceClusterConfigMetadata, ok := rGceClusterConfig["metadata"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rGceClusterConfigMetadata {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.GceClusterConfig.Metadata = m
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.Metadata: expected map[string]interface{}")
				}
			}
			if _, ok := rGceClusterConfig["network"]; ok {
				if s, ok := rGceClusterConfig["network"].(string); ok {
					r.GceClusterConfig.Network = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.Network: expected string")
				}
			}
			if _, ok := rGceClusterConfig["nodeGroupAffinity"]; ok {
				if rGceClusterConfigNodeGroupAffinity, ok := rGceClusterConfig["nodeGroupAffinity"].(map[string]interface{}); ok {
					r.GceClusterConfig.NodeGroupAffinity = &dclService.ClusterClusterConfigGceClusterConfigNodeGroupAffinity{}
					if _, ok := rGceClusterConfigNodeGroupAffinity["nodeGroup"]; ok {
						if s, ok := rGceClusterConfigNodeGroupAffinity["nodeGroup"].(string); ok {
							r.GceClusterConfig.NodeGroupAffinity.NodeGroup = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GceClusterConfig.NodeGroupAffinity.NodeGroup: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.NodeGroupAffinity: expected map[string]interface{}")
				}
			}
			if _, ok := rGceClusterConfig["privateIPv6GoogleAccess"]; ok {
				if s, ok := rGceClusterConfig["privateIPv6GoogleAccess"].(string); ok {
					r.GceClusterConfig.PrivateIPv6GoogleAccess = dclService.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.PrivateIPv6GoogleAccess: expected string")
				}
			}
			if _, ok := rGceClusterConfig["reservationAffinity"]; ok {
				if rGceClusterConfigReservationAffinity, ok := rGceClusterConfig["reservationAffinity"].(map[string]interface{}); ok {
					r.GceClusterConfig.ReservationAffinity = &dclService.ClusterClusterConfigGceClusterConfigReservationAffinity{}
					if _, ok := rGceClusterConfigReservationAffinity["consumeReservationType"]; ok {
						if s, ok := rGceClusterConfigReservationAffinity["consumeReservationType"].(string); ok {
							r.GceClusterConfig.ReservationAffinity.ConsumeReservationType = dclService.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.GceClusterConfig.ReservationAffinity.ConsumeReservationType: expected string")
						}
					}
					if _, ok := rGceClusterConfigReservationAffinity["key"]; ok {
						if s, ok := rGceClusterConfigReservationAffinity["key"].(string); ok {
							r.GceClusterConfig.ReservationAffinity.Key = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GceClusterConfig.ReservationAffinity.Key: expected string")
						}
					}
					if _, ok := rGceClusterConfigReservationAffinity["values"]; ok {
						if s, ok := rGceClusterConfigReservationAffinity["values"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.GceClusterConfig.ReservationAffinity.Values = append(r.GceClusterConfig.ReservationAffinity.Values, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.GceClusterConfig.ReservationAffinity.Values: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.ReservationAffinity: expected map[string]interface{}")
				}
			}
			if _, ok := rGceClusterConfig["serviceAccount"]; ok {
				if s, ok := rGceClusterConfig["serviceAccount"].(string); ok {
					r.GceClusterConfig.ServiceAccount = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.ServiceAccount: expected string")
				}
			}
			if _, ok := rGceClusterConfig["serviceAccountScopes"]; ok {
				if s, ok := rGceClusterConfig["serviceAccountScopes"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.GceClusterConfig.ServiceAccountScopes = append(r.GceClusterConfig.ServiceAccountScopes, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.ServiceAccountScopes: expected []interface{}")
				}
			}
			if _, ok := rGceClusterConfig["subnetwork"]; ok {
				if s, ok := rGceClusterConfig["subnetwork"].(string); ok {
					r.GceClusterConfig.Subnetwork = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.Subnetwork: expected string")
				}
			}
			if _, ok := rGceClusterConfig["tags"]; ok {
				if s, ok := rGceClusterConfig["tags"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.GceClusterConfig.Tags = append(r.GceClusterConfig.Tags, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.Tags: expected []interface{}")
				}
			}
			if _, ok := rGceClusterConfig["zone"]; ok {
				if s, ok := rGceClusterConfig["zone"].(string); ok {
					r.GceClusterConfig.Zone = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GceClusterConfig.Zone: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.GceClusterConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["gkeClusterConfig"]; ok {
		if rGkeClusterConfig, ok := obj["gkeClusterConfig"].(map[string]interface{}); ok {
			r.GkeClusterConfig = &dclService.ClusterClusterConfigGkeClusterConfig{}
			if _, ok := rGkeClusterConfig["namespacedGkeDeploymentTarget"]; ok {
				if rGkeClusterConfigNamespacedGkeDeploymentTarget, ok := rGkeClusterConfig["namespacedGkeDeploymentTarget"].(map[string]interface{}); ok {
					r.GkeClusterConfig.NamespacedGkeDeploymentTarget = &dclService.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{}
					if _, ok := rGkeClusterConfigNamespacedGkeDeploymentTarget["clusterNamespace"]; ok {
						if s, ok := rGkeClusterConfigNamespacedGkeDeploymentTarget["clusterNamespace"].(string); ok {
							r.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace: expected string")
						}
					}
					if _, ok := rGkeClusterConfigNamespacedGkeDeploymentTarget["targetGkeCluster"]; ok {
						if s, ok := rGkeClusterConfigNamespacedGkeDeploymentTarget["targetGkeCluster"].(string); ok {
							r.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.GkeClusterConfig.NamespacedGkeDeploymentTarget: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.GkeClusterConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["initializationActions"]; ok {
		if s, ok := obj["initializationActions"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rInitializationActions dclService.ClusterClusterConfigInitializationActions
					if _, ok := objval["executableFile"]; ok {
						if s, ok := objval["executableFile"].(string); ok {
							rInitializationActions.ExecutableFile = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rInitializationActions.ExecutableFile: expected string")
						}
					}
					if _, ok := objval["executionTimeout"]; ok {
						if s, ok := objval["executionTimeout"].(string); ok {
							rInitializationActions.ExecutionTimeout = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rInitializationActions.ExecutionTimeout: expected string")
						}
					}
					r.InitializationActions = append(r.InitializationActions, rInitializationActions)
				}
			}
		} else {
			return nil, fmt.Errorf("r.InitializationActions: expected []interface{}")
		}
	}
	if _, ok := obj["lifecycleConfig"]; ok {
		if rLifecycleConfig, ok := obj["lifecycleConfig"].(map[string]interface{}); ok {
			r.LifecycleConfig = &dclService.ClusterClusterConfigLifecycleConfig{}
			if _, ok := rLifecycleConfig["autoDeleteTime"]; ok {
				if s, ok := rLifecycleConfig["autoDeleteTime"].(string); ok {
					r.LifecycleConfig.AutoDeleteTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LifecycleConfig.AutoDeleteTime: expected string")
				}
			}
			if _, ok := rLifecycleConfig["autoDeleteTtl"]; ok {
				if s, ok := rLifecycleConfig["autoDeleteTtl"].(string); ok {
					r.LifecycleConfig.AutoDeleteTtl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LifecycleConfig.AutoDeleteTtl: expected string")
				}
			}
			if _, ok := rLifecycleConfig["idleDeleteTtl"]; ok {
				if s, ok := rLifecycleConfig["idleDeleteTtl"].(string); ok {
					r.LifecycleConfig.IdleDeleteTtl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LifecycleConfig.IdleDeleteTtl: expected string")
				}
			}
			if _, ok := rLifecycleConfig["idleStartTime"]; ok {
				if s, ok := rLifecycleConfig["idleStartTime"].(string); ok {
					r.LifecycleConfig.IdleStartTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LifecycleConfig.IdleStartTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LifecycleConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["masterConfig"]; ok {
		if rMasterConfig, ok := obj["masterConfig"].(map[string]interface{}); ok {
			var err error
			r.MasterConfig, err = UnstructuredToClusterInstanceGroupConfig(rMasterConfig)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("r.MasterConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["metastoreConfig"]; ok {
		if rMetastoreConfig, ok := obj["metastoreConfig"].(map[string]interface{}); ok {
			r.MetastoreConfig = &dclService.ClusterClusterConfigMetastoreConfig{}
			if _, ok := rMetastoreConfig["dataprocMetastoreService"]; ok {
				if s, ok := rMetastoreConfig["dataprocMetastoreService"].(string); ok {
					r.MetastoreConfig.DataprocMetastoreService = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MetastoreConfig.DataprocMetastoreService: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MetastoreConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["secondaryWorkerConfig"]; ok {
		if rSecondaryWorkerConfig, ok := obj["secondaryWorkerConfig"].(map[string]interface{}); ok {
			var err error
			r.SecondaryWorkerConfig, err = UnstructuredToClusterInstanceGroupConfig(rSecondaryWorkerConfig)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("r.SecondaryWorkerConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["securityConfig"]; ok {
		if rSecurityConfig, ok := obj["securityConfig"].(map[string]interface{}); ok {
			r.SecurityConfig = &dclService.ClusterClusterConfigSecurityConfig{}
			if _, ok := rSecurityConfig["kerberosConfig"]; ok {
				if rSecurityConfigKerberosConfig, ok := rSecurityConfig["kerberosConfig"].(map[string]interface{}); ok {
					r.SecurityConfig.KerberosConfig = &dclService.ClusterClusterConfigSecurityConfigKerberosConfig{}
					if _, ok := rSecurityConfigKerberosConfig["crossRealmTrustAdminServer"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["crossRealmTrustAdminServer"].(string); ok {
							r.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["crossRealmTrustKdc"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["crossRealmTrustKdc"].(string); ok {
							r.SecurityConfig.KerberosConfig.CrossRealmTrustKdc = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.CrossRealmTrustKdc: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["crossRealmTrustRealm"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["crossRealmTrustRealm"].(string); ok {
							r.SecurityConfig.KerberosConfig.CrossRealmTrustRealm = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.CrossRealmTrustRealm: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["crossRealmTrustSharedPassword"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["crossRealmTrustSharedPassword"].(string); ok {
							r.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["enableKerberos"]; ok {
						if b, ok := rSecurityConfigKerberosConfig["enableKerberos"].(bool); ok {
							r.SecurityConfig.KerberosConfig.EnableKerberos = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.EnableKerberos: expected bool")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["kdcDbKey"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["kdcDbKey"].(string); ok {
							r.SecurityConfig.KerberosConfig.KdcDbKey = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.KdcDbKey: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["keyPassword"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["keyPassword"].(string); ok {
							r.SecurityConfig.KerberosConfig.KeyPassword = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.KeyPassword: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["keystore"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["keystore"].(string); ok {
							r.SecurityConfig.KerberosConfig.Keystore = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.Keystore: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["keystorePassword"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["keystorePassword"].(string); ok {
							r.SecurityConfig.KerberosConfig.KeystorePassword = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.KeystorePassword: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["kmsKey"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["kmsKey"].(string); ok {
							r.SecurityConfig.KerberosConfig.KmsKey = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.KmsKey: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["realm"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["realm"].(string); ok {
							r.SecurityConfig.KerberosConfig.Realm = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.Realm: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["rootPrincipalPassword"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["rootPrincipalPassword"].(string); ok {
							r.SecurityConfig.KerberosConfig.RootPrincipalPassword = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.RootPrincipalPassword: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["tgtLifetimeHours"]; ok {
						if i, ok := rSecurityConfigKerberosConfig["tgtLifetimeHours"].(int64); ok {
							r.SecurityConfig.KerberosConfig.TgtLifetimeHours = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.TgtLifetimeHours: expected int64")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["truststore"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["truststore"].(string); ok {
							r.SecurityConfig.KerberosConfig.Truststore = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.Truststore: expected string")
						}
					}
					if _, ok := rSecurityConfigKerberosConfig["truststorePassword"]; ok {
						if s, ok := rSecurityConfigKerberosConfig["truststorePassword"].(string); ok {
							r.SecurityConfig.KerberosConfig.TruststorePassword = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig.TruststorePassword: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.SecurityConfig.KerberosConfig: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SecurityConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["softwareConfig"]; ok {
		if rSoftwareConfig, ok := obj["softwareConfig"].(map[string]interface{}); ok {
			r.SoftwareConfig = &dclService.ClusterClusterConfigSoftwareConfig{}
			if _, ok := rSoftwareConfig["imageVersion"]; ok {
				if s, ok := rSoftwareConfig["imageVersion"].(string); ok {
					r.SoftwareConfig.ImageVersion = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SoftwareConfig.ImageVersion: expected string")
				}
			}
			if _, ok := rSoftwareConfig["optionalComponents"]; ok {
				if s, ok := rSoftwareConfig["optionalComponents"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.SoftwareConfig.OptionalComponents = append(r.SoftwareConfig.OptionalComponents, dclService.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum(strval))
						}
					}
				} else {
					return nil, fmt.Errorf("r.SoftwareConfig.OptionalComponents: expected []interface{}")
				}
			}
			if _, ok := rSoftwareConfig["properties"]; ok {
				if rSoftwareConfigProperties, ok := rSoftwareConfig["properties"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rSoftwareConfigProperties {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.SoftwareConfig.Properties = m
				} else {
					return nil, fmt.Errorf("r.SoftwareConfig.Properties: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SoftwareConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["stagingBucket"]; ok {
		if s, ok := obj["stagingBucket"].(string); ok {
			r.StagingBucket = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.StagingBucket: expected string")
		}
	}
	if _, ok := obj["tempBucket"]; ok {
		if s, ok := obj["tempBucket"].(string); ok {
			r.TempBucket = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.TempBucket: expected string")
		}
	}
	if _, ok := obj["workerConfig"]; ok {
		if rWorkerConfig, ok := obj["workerConfig"].(map[string]interface{}); ok {
			var err error
			r.WorkerConfig, err = UnstructuredToClusterInstanceGroupConfig(rWorkerConfig)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("r.WorkerConfig: expected map[string]interface{}")
		}
	}
	return r, nil
}

func UnstructuredToClusterInstanceGroupConfig(obj map[string]interface{}) (*dclService.ClusterInstanceGroupConfig, error) {
	r := &dclService.ClusterInstanceGroupConfig{}
	if _, ok := obj["accelerators"]; ok {
		if s, ok := obj["accelerators"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rAccelerators dclService.ClusterInstanceGroupConfigAccelerators
					if _, ok := objval["acceleratorCount"]; ok {
						if i, ok := objval["acceleratorCount"].(int64); ok {
							rAccelerators.AcceleratorCount = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rAccelerators.AcceleratorCount: expected int64")
						}
					}
					if _, ok := objval["acceleratorType"]; ok {
						if s, ok := objval["acceleratorType"].(string); ok {
							rAccelerators.AcceleratorType = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAccelerators.AcceleratorType: expected string")
						}
					}
					r.Accelerators = append(r.Accelerators, rAccelerators)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Accelerators: expected []interface{}")
		}
	}
	if _, ok := obj["diskConfig"]; ok {
		if rDiskConfig, ok := obj["diskConfig"].(map[string]interface{}); ok {
			r.DiskConfig = &dclService.ClusterInstanceGroupConfigDiskConfig{}
			if _, ok := rDiskConfig["bootDiskSizeGb"]; ok {
				if i, ok := rDiskConfig["bootDiskSizeGb"].(int64); ok {
					r.DiskConfig.BootDiskSizeGb = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.DiskConfig.BootDiskSizeGb: expected int64")
				}
			}
			if _, ok := rDiskConfig["bootDiskType"]; ok {
				if s, ok := rDiskConfig["bootDiskType"].(string); ok {
					r.DiskConfig.BootDiskType = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DiskConfig.BootDiskType: expected string")
				}
			}
			if _, ok := rDiskConfig["numLocalSsds"]; ok {
				if i, ok := rDiskConfig["numLocalSsds"].(int64); ok {
					r.DiskConfig.NumLocalSsds = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.DiskConfig.NumLocalSsds: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DiskConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["image"]; ok {
		if s, ok := obj["image"].(string); ok {
			r.Image = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Image: expected string")
		}
	}
	if _, ok := obj["instanceNames"]; ok {
		if s, ok := obj["instanceNames"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.InstanceNames = append(r.InstanceNames, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.InstanceNames: expected []interface{}")
		}
	}
	if _, ok := obj["isPreemptible"]; ok {
		if b, ok := obj["isPreemptible"].(bool); ok {
			r.IsPreemptible = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.IsPreemptible: expected bool")
		}
	}
	if _, ok := obj["machineType"]; ok {
		if s, ok := obj["machineType"].(string); ok {
			r.MachineType = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.MachineType: expected string")
		}
	}
	if _, ok := obj["managedGroupConfig"]; ok {
		if rManagedGroupConfig, ok := obj["managedGroupConfig"].(map[string]interface{}); ok {
			r.ManagedGroupConfig = &dclService.ClusterInstanceGroupConfigManagedGroupConfig{}
			if _, ok := rManagedGroupConfig["instanceGroupManagerName"]; ok {
				if s, ok := rManagedGroupConfig["instanceGroupManagerName"].(string); ok {
					r.ManagedGroupConfig.InstanceGroupManagerName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ManagedGroupConfig.InstanceGroupManagerName: expected string")
				}
			}
			if _, ok := rManagedGroupConfig["instanceTemplateName"]; ok {
				if s, ok := rManagedGroupConfig["instanceTemplateName"].(string); ok {
					r.ManagedGroupConfig.InstanceTemplateName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ManagedGroupConfig.InstanceTemplateName: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ManagedGroupConfig: expected map[string]interface{}")
		}
	}
	if _, ok := obj["minCpuPlatform"]; ok {
		if s, ok := obj["minCpuPlatform"].(string); ok {
			r.MinCpuPlatform = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.MinCpuPlatform: expected string")
		}
	}
	if _, ok := obj["numInstances"]; ok {
		if i, ok := obj["numInstances"].(int64); ok {
			r.NumInstances = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.NumInstances: expected int64")
		}
	}
	if _, ok := obj["preemptibility"]; ok {
		if s, ok := obj["preemptibility"].(string); ok {
			r.Preemptibility = dclService.ClusterInstanceGroupConfigPreemptibilityEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Preemptibility: expected string")
		}
	}
	return r, nil
}

func GetCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetCluster(ctx, r)
	if err != nil {
		return nil, err
	}
	return ClusterToUnstructured(r), nil
}

func ListCluster(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListCluster(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ClusterToUnstructured(r))
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

func ApplyCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCluster(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyCluster(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ClusterToUnstructured(r), nil
}

func ClusterHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCluster(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyCluster(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return err
	}
	return c.DeleteCluster(ctx, r)
}

func ClusterID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Cluster) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dataproc",
		"Cluster",
		"alpha",
	}
}

func SetPolicyCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func GetPolicyCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policy, err := iamClient.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func (r *Cluster) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyCluster(ctx, config, resource, policy)
}

func (r *Cluster) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyCluster(ctx, config, resource)
}

func (r *Cluster) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetCluster(ctx, config, resource)
}

func (r *Cluster) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyCluster(ctx, config, resource, opts...)
}

func (r *Cluster) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ClusterHasDiff(ctx, config, resource, opts...)
}

func (r *Cluster) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteCluster(ctx, config, resource)
}

func (r *Cluster) ID(resource *unstructured.Resource) (string, error) {
	return ClusterID(resource)
}

func init() {
	unstructured.Register(&Cluster{})
}
