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
package container

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Cluster) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MasterAuth) {
		if err := r.MasterAuth.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AddonsConfig) {
		if err := r.AddonsConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LegacyAbac) {
		if err := r.LegacyAbac.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.NetworkPolicy) {
		if err := r.NetworkPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.IPAllocationPolicy) {
		if err := r.IPAllocationPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MasterAuthorizedNetworksConfig) {
		if err := r.MasterAuthorizedNetworksConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BinaryAuthorization) {
		if err := r.BinaryAuthorization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Autoscaling) {
		if err := r.Autoscaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.NetworkConfig) {
		if err := r.NetworkConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaintenancePolicy) {
		if err := r.MaintenancePolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DefaultMaxPodsConstraint) {
		if err := r.DefaultMaxPodsConstraint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ResourceUsageExportConfig) {
		if err := r.ResourceUsageExportConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AuthenticatorGroupsConfig) {
		if err := r.AuthenticatorGroupsConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PrivateClusterConfig) {
		if err := r.PrivateClusterConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DatabaseEncryption) {
		if err := r.DatabaseEncryption.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.VerticalPodAutoscaling) {
		if err := r.VerticalPodAutoscaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ShieldedNodes) {
		if err := r.ShieldedNodes.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Autopilot) {
		if err := r.Autopilot.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.NodeConfig) {
		if err := r.NodeConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReleaseChannel) {
		if err := r.ReleaseChannel.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.WorkloadIdentityConfig) {
		if err := r.WorkloadIdentityConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.NotificationConfig) {
		if err := r.NotificationConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConfidentialNodes) {
		if err := r.ConfidentialNodes.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterMasterAuth) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ClientCertificateConfig) {
		if err := r.ClientCertificateConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterMasterAuthClientCertificateConfig) validate() error {
	return nil
}
func (r *ClusterAddonsConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.HttpLoadBalancing) {
		if err := r.HttpLoadBalancing.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HorizontalPodAutoscaling) {
		if err := r.HorizontalPodAutoscaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.KubernetesDashboard) {
		if err := r.KubernetesDashboard.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.NetworkPolicyConfig) {
		if err := r.NetworkPolicyConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudRunConfig) {
		if err := r.CloudRunConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DnsCacheConfig) {
		if err := r.DnsCacheConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConfigConnectorConfig) {
		if err := r.ConfigConnectorConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.GcePersistentDiskCsiDriverConfig) {
		if err := r.GcePersistentDiskCsiDriverConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterAddonsConfigHttpLoadBalancing) validate() error {
	return nil
}
func (r *ClusterAddonsConfigHorizontalPodAutoscaling) validate() error {
	return nil
}
func (r *ClusterAddonsConfigKubernetesDashboard) validate() error {
	return nil
}
func (r *ClusterAddonsConfigNetworkPolicyConfig) validate() error {
	return nil
}
func (r *ClusterAddonsConfigCloudRunConfig) validate() error {
	return nil
}
func (r *ClusterAddonsConfigDnsCacheConfig) validate() error {
	return nil
}
func (r *ClusterAddonsConfigConfigConnectorConfig) validate() error {
	return nil
}
func (r *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) validate() error {
	return nil
}
func (r *ClusterNodePools) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Config) {
		if err := r.Config.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Autoscaling) {
		if err := r.Autoscaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Management) {
		if err := r.Management.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaxPodsConstraint) {
		if err := r.MaxPodsConstraint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.UpgradeSettings) {
		if err := r.UpgradeSettings.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterNodePoolsConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.WorkloadMetadataConfig) {
		if err := r.WorkloadMetadataConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SandboxConfig) {
		if err := r.SandboxConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReservationAffinity) {
		if err := r.ReservationAffinity.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ShieldedInstanceConfig) {
		if err := r.ShieldedInstanceConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LinuxNodeConfig) {
		if err := r.LinuxNodeConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.KubeletConfig) {
		if err := r.KubeletConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterNodePoolsConfigAccelerators) validate() error {
	return nil
}
func (r *ClusterNodePoolsConfigWorkloadMetadataConfig) validate() error {
	return nil
}
func (r *ClusterNodePoolsConfigTaints) validate() error {
	return nil
}
func (r *ClusterNodePoolsConfigSandboxConfig) validate() error {
	return nil
}
func (r *ClusterNodePoolsConfigReservationAffinity) validate() error {
	return nil
}
func (r *ClusterNodePoolsConfigShieldedInstanceConfig) validate() error {
	return nil
}
func (r *ClusterNodePoolsConfigLinuxNodeConfig) validate() error {
	return nil
}
func (r *ClusterNodePoolsConfigKubeletConfig) validate() error {
	return nil
}
func (r *ClusterNodePoolsAutoscaling) validate() error {
	return nil
}
func (r *ClusterNodePoolsManagement) validate() error {
	if !dcl.IsEmptyValueIndirect(r.UpgradeOptions) {
		if err := r.UpgradeOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterNodePoolsManagementUpgradeOptions) validate() error {
	return nil
}
func (r *ClusterNodePoolsMaxPodsConstraint) validate() error {
	return nil
}
func (r *ClusterNodePoolsConditions) validate() error {
	return nil
}
func (r *ClusterNodePoolsUpgradeSettings) validate() error {
	return nil
}
func (r *ClusterLegacyAbac) validate() error {
	return nil
}
func (r *ClusterNetworkPolicy) validate() error {
	return nil
}
func (r *ClusterIPAllocationPolicy) validate() error {
	return nil
}
func (r *ClusterMasterAuthorizedNetworksConfig) validate() error {
	return nil
}
func (r *ClusterMasterAuthorizedNetworksConfigCidrBlocks) validate() error {
	return nil
}
func (r *ClusterBinaryAuthorization) validate() error {
	return nil
}
func (r *ClusterAutoscaling) validate() error {
	if !dcl.IsEmptyValueIndirect(r.AutoprovisioningNodePoolDefaults) {
		if err := r.AutoprovisioningNodePoolDefaults.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterAutoscalingResourceLimits) validate() error {
	return nil
}
func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaults) validate() error {
	if !dcl.IsEmptyValueIndirect(r.UpgradeSettings) {
		if err := r.UpgradeSettings.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Management) {
		if err := r.Management.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ShieldedInstanceConfig) {
		if err := r.ShieldedInstanceConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) validate() error {
	return nil
}
func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) validate() error {
	if !dcl.IsEmptyValueIndirect(r.UpgradeOptions) {
		if err := r.UpgradeOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) validate() error {
	return nil
}
func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) validate() error {
	return nil
}
func (r *ClusterNetworkConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DefaultSnatStatus) {
		if err := r.DefaultSnatStatus.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterNetworkConfigDefaultSnatStatus) validate() error {
	return nil
}
func (r *ClusterMaintenancePolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Window) {
		if err := r.Window.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterMaintenancePolicyWindow) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DailyMaintenanceWindow) {
		if err := r.DailyMaintenanceWindow.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RecurringWindow) {
		if err := r.RecurringWindow.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterMaintenancePolicyWindowDailyMaintenanceWindow) validate() error {
	return nil
}
func (r *ClusterMaintenancePolicyWindowRecurringWindow) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Window) {
		if err := r.Window.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterMaintenancePolicyWindowRecurringWindowWindow) validate() error {
	return nil
}
func (r *ClusterDefaultMaxPodsConstraint) validate() error {
	return nil
}
func (r *ClusterResourceUsageExportConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BigqueryDestination) {
		if err := r.BigqueryDestination.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConsumptionMeteringConfig) {
		if err := r.ConsumptionMeteringConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterResourceUsageExportConfigBigqueryDestination) validate() error {
	return nil
}
func (r *ClusterResourceUsageExportConfigConsumptionMeteringConfig) validate() error {
	return nil
}
func (r *ClusterAuthenticatorGroupsConfig) validate() error {
	return nil
}
func (r *ClusterPrivateClusterConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MasterGlobalAccessConfig) {
		if err := r.MasterGlobalAccessConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterPrivateClusterConfigMasterGlobalAccessConfig) validate() error {
	return nil
}
func (r *ClusterDatabaseEncryption) validate() error {
	return nil
}
func (r *ClusterVerticalPodAutoscaling) validate() error {
	return nil
}
func (r *ClusterShieldedNodes) validate() error {
	return nil
}
func (r *ClusterConditions) validate() error {
	return nil
}
func (r *ClusterAutopilot) validate() error {
	return nil
}
func (r *ClusterNodeConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.WorkloadMetadataConfig) {
		if err := r.WorkloadMetadataConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SandboxConfig) {
		if err := r.SandboxConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReservationAffinity) {
		if err := r.ReservationAffinity.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ShieldedInstanceConfig) {
		if err := r.ShieldedInstanceConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LinuxNodeConfig) {
		if err := r.LinuxNodeConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.KubeletConfig) {
		if err := r.KubeletConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterNodeConfigAccelerators) validate() error {
	return nil
}
func (r *ClusterNodeConfigWorkloadMetadataConfig) validate() error {
	return nil
}
func (r *ClusterNodeConfigTaints) validate() error {
	return nil
}
func (r *ClusterNodeConfigSandboxConfig) validate() error {
	return nil
}
func (r *ClusterNodeConfigReservationAffinity) validate() error {
	return nil
}
func (r *ClusterNodeConfigShieldedInstanceConfig) validate() error {
	return nil
}
func (r *ClusterNodeConfigLinuxNodeConfig) validate() error {
	return nil
}
func (r *ClusterNodeConfigKubeletConfig) validate() error {
	return nil
}
func (r *ClusterReleaseChannel) validate() error {
	return nil
}
func (r *ClusterWorkloadIdentityConfig) validate() error {
	return nil
}
func (r *ClusterNotificationConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Pubsub) {
		if err := r.Pubsub.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterNotificationConfigPubsub) validate() error {
	return nil
}
func (r *ClusterConfidentialNodes) validate() error {
	return nil
}

func clusterGetURL(userBasePath string, r *Cluster) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, params), nil
}

func clusterListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters", "https://container.googleapis.com/v1/", userBasePath, params), nil

}

func clusterCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters", "https://container.googleapis.com/v1/", userBasePath, params), nil

}

func clusterDeleteURL(userBasePath string, r *Cluster) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, params), nil
}

// clusterApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type clusterApiOperation interface {
	do(context.Context, *Cluster, *Client) error
}

// newUpdateClusterSetMaintenancePolicyRequest creates a request for an
// Cluster resource's setMaintenancePolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterSetMaintenancePolicyRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandClusterMaintenancePolicy(c, f.MaintenancePolicy); err != nil {
		return nil, fmt.Errorf("error expanding MaintenancePolicy into maintenancePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["maintenancePolicy"] = v
	}
	return req, nil
}

// marshalUpdateClusterSetMaintenancePolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterSetMaintenancePolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateClusterSetMaintenancePolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterSetMaintenancePolicyOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setMaintenancePolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterSetMaintenancePolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterSetMaintenancePolicyRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateAddonsConfigRequest creates a request for an
// Cluster resource's updateAddonsConfig update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateAddonsConfigRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandClusterAddonsConfig(c, f.AddonsConfig); err != nil {
		return nil, fmt.Errorf("error expanding AddonsConfig into desiredAddonsConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["desiredAddonsConfig"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateAddonsConfigRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateAddonsConfigRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateAddonsConfigOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateAddonsConfigOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateAddonsConfig")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateAddonsConfigRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateAddonsConfigRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateAutoscalingRequest creates a request for an
// Cluster resource's updateAutoscaling update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateAutoscalingRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateClusterUpdateAutoscalingRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateAutoscalingRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateAutoscalingOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateAutoscalingOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateAutoscaling")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateAutoscalingRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateAutoscalingRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateBinaryAuthorizationRequest creates a request for an
// Cluster resource's updateBinaryAuthorization update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateBinaryAuthorizationRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandClusterBinaryAuthorization(c, f.BinaryAuthorization); err != nil {
		return nil, fmt.Errorf("error expanding BinaryAuthorization into desiredBinaryAuthorization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["desiredBinaryAuthorization"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateBinaryAuthorizationRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateBinaryAuthorizationRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateBinaryAuthorizationOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateBinaryAuthorizationOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateBinaryAuthorization")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateBinaryAuthorizationRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateBinaryAuthorizationRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateDatabaseEncryptionRequest creates a request for an
// Cluster resource's updateDatabaseEncryption update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateDatabaseEncryptionRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandClusterDatabaseEncryption(c, f.DatabaseEncryption); err != nil {
		return nil, fmt.Errorf("error expanding DatabaseEncryption into desiredDatabaseEncryption: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["desiredDatabaseEncryption"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateDatabaseEncryptionRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateDatabaseEncryptionRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateDatabaseEncryptionOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateDatabaseEncryptionOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateDatabaseEncryption")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateDatabaseEncryptionRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateDatabaseEncryptionRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateLegacyAbacRequest creates a request for an
// Cluster resource's updateLegacyAbac update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateLegacyAbacRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandClusterLegacyAbac(c, f.LegacyAbac); err != nil {
		return nil, fmt.Errorf("error expanding LegacyAbac into legacyAbac: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["legacyAbac"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateLegacyAbacRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateLegacyAbacRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterLegacyAbacUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateLegacyAbacOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateLegacyAbacOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateLegacyAbac")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateLegacyAbacRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateLegacyAbacRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateLocationsRequest creates a request for an
// Cluster resource's updateLocations update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateLocationsRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Locations; !dcl.IsEmptyValueIndirect(v) {
		req["desiredLocations"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateLocationsRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateLocationsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateLocationsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateLocationsOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateLocations")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateLocationsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateLocationsRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateMasterAuthorizedNetworksConfigRequest creates a request for an
// Cluster resource's updateMasterAuthorizedNetworksConfig update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateMasterAuthorizedNetworksConfigRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandClusterMasterAuthorizedNetworksConfig(c, f.MasterAuthorizedNetworksConfig); err != nil {
		return nil, fmt.Errorf("error expanding MasterAuthorizedNetworksConfig into desiredMasterAuthorizedNetworksConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["desiredMasterAuthorizedNetworksConfig"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateMasterAuthorizedNetworksConfigRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateMasterAuthorizedNetworksConfigRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateMasterAuthorizedNetworksConfigOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateMasterAuthorizedNetworksConfigOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateMasterAuthorizedNetworksConfig")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateMasterAuthorizedNetworksConfigRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateMasterAuthorizedNetworksConfigRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateMasterVersionRequest creates a request for an
// Cluster resource's updateMasterVersion update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateMasterVersionRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.MasterVersion; !dcl.IsEmptyValueIndirect(v) {
		req["desiredMasterVersion"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateMasterVersionRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateMasterVersionRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateMasterVersionOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateMasterVersionOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateMasterVersion")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateMasterVersionRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateMasterVersionRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateMonitoringAndLoggingServiceRequest creates a request for an
// Cluster resource's updateMonitoringAndLoggingService update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateMonitoringAndLoggingServiceRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.LoggingService; !dcl.IsEmptyValueIndirect(v) {
		req["desiredLoggingService"] = v
	}
	if v := f.MonitoringService; !dcl.IsEmptyValueIndirect(v) {
		req["desiredMonitoringService"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateMonitoringAndLoggingServiceRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateMonitoringAndLoggingServiceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateMonitoringAndLoggingServiceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateMonitoringAndLoggingServiceOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateMonitoringAndLoggingService")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateMonitoringAndLoggingServiceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateMonitoringAndLoggingServiceRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateShieldedNodesRequest creates a request for an
// Cluster resource's updateShieldedNodes update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateShieldedNodesRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandClusterShieldedNodes(c, f.ShieldedNodes); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedNodes into desiredShieldedNodes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["desiredShieldedNodes"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateShieldedNodesRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateShieldedNodesRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateShieldedNodesOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateShieldedNodesOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateShieldedNodes")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateShieldedNodesRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateShieldedNodesRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateVerticalPodAutoscalingRequest creates a request for an
// Cluster resource's updateVerticalPodAutoscaling update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateVerticalPodAutoscalingRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandClusterVerticalPodAutoscaling(c, f.VerticalPodAutoscaling); err != nil {
		return nil, fmt.Errorf("error expanding VerticalPodAutoscaling into desiredVerticalPodAutoscaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["desiredVerticalPodAutoscaling"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateVerticalPodAutoscalingRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateVerticalPodAutoscalingRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateVerticalPodAutoscalingOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateVerticalPodAutoscalingOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateVerticalPodAutoscaling")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateVerticalPodAutoscalingRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateVerticalPodAutoscalingRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateClusterUpdateWorkloadIdentityConfigRequest creates a request for an
// Cluster resource's updateWorkloadIdentityConfig update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateClusterUpdateWorkloadIdentityConfigRequest(ctx context.Context, f *Cluster, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandClusterWorkloadIdentityConfig(c, f.WorkloadIdentityConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkloadIdentityConfig into desiredWorkloadIdentityConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["desiredWorkloadIdentityConfig"] = v
	}
	return req, nil
}

// marshalUpdateClusterUpdateWorkloadIdentityConfigRequest converts the update into
// the final JSON request body.
func marshalUpdateClusterUpdateWorkloadIdentityConfigRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeClusterUpdateRequest(m)
	return json.Marshal(m)
}

type updateClusterUpdateWorkloadIdentityConfigOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateClusterUpdateWorkloadIdentityConfigOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	_, err := c.GetCluster(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateWorkloadIdentityConfig")
	if err != nil {
		return err
	}

	req, err := newUpdateClusterUpdateWorkloadIdentityConfigRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateClusterUpdateWorkloadIdentityConfigRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

func (c *Client) listClusterRaw(ctx context.Context, project, location, pageToken string) ([]byte, error) {
	u, err := clusterListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listClusterOperation struct {
	Clusters []map[string]interface{} `json:"clusters"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listCluster(ctx context.Context, project, location, pageToken string) ([]*Cluster, string, error) {
	b, err := c.listClusterRaw(ctx, project, location, pageToken)
	if err != nil {
		return nil, "", err
	}

	var m listClusterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Cluster
	for _, v := range m.Clusters {
		res := flattenCluster(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllCluster(ctx context.Context, f func(*Cluster) bool, resources []*Cluster) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteCluster(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteClusterOperation struct{}

func (op *deleteClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {

	_, err := c.GetCluster(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Cluster not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetCluster checking for existence. error: %v", err)
		return err
	}

	if err := clusterWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := clusterDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetCluster(ctx, r.urlNormalized())
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createClusterOperation struct {
	response map[string]interface{}
}

func (op *createClusterOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := clusterCreateURL(c.Config.BasePath, project, location)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetCluster(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	err = r.DeleteDefaultNodePool(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) getClusterRaw(ctx context.Context, r *Cluster) ([]byte, error) {

	u, err := clusterGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) clusterDiffsForRawDesired(ctx context.Context, rawDesired *Cluster, opts ...dcl.ApplyOption) (initial, desired *Cluster, diffs []clusterDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Cluster
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Cluster); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Cluster, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetCluster(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Cluster resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Cluster resource: %v", err)
		}
		c.Config.Logger.Info("Found that Cluster resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Cluster: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Cluster: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeClusterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Cluster: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Cluster: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffCluster(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func clusterCheckForErrorStateAndDelete(ctx context.Context, r *Cluster, c *Client) error {
	if dcl.FindStringInArray(*r.Status, []string{"STOPPING", "ERROR", "DEGRADED", "RUNNING_WITH_ERROR"}) {
		err := c.DeleteCluster(ctx, r.urlNormalized())
		if err != nil {
			return fmt.Errorf("Cluster was is in error state %s, delete attempted with error: %v", *r.StatusMessage, err)
		} else {
			return fmt.Errorf("Cluster was is in error state %s, delete succeeded.", *r.StatusMessage)
		}
	}
	return nil
}

type clusterWaitOperation struct {
	Client   *Client
	Resource *Cluster
}

func (op *clusterWaitOperation) operate(ctx context.Context) (*dcl.RetryDetails, error) {
	current, err := op.Client.GetCluster(ctx, op.Resource.urlNormalized())
	if err != nil {
		return nil, err
	}
	if dcl.FindStringInArray(*current.Status, []string{"PROVISIONING", "RECONCILING"}) {
		return nil, dcl.OperationNotDone{}
	}
	return &dcl.RetryDetails{}, nil
}

func clusterWaitForRestingState(ctx context.Context, r *Cluster, c *Client) error {
	op := clusterWaitOperation{
		Client:   c,
		Resource: r,
	}
	return dcl.Do(ctx, op.operate, c.Config.RetryProvider)
}

func canonicalizeClusterInitialState(rawInitial, rawDesired *Cluster) (*Cluster, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeClusterDesiredState(rawDesired, rawInitial *Cluster, opts ...dcl.ApplyOption) (*Cluster, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MasterAuth = canonicalizeClusterMasterAuth(rawDesired.MasterAuth, nil, opts...)
		rawDesired.AddonsConfig = canonicalizeClusterAddonsConfig(rawDesired.AddonsConfig, nil, opts...)
		rawDesired.LegacyAbac = canonicalizeClusterLegacyAbac(rawDesired.LegacyAbac, nil, opts...)
		rawDesired.NetworkPolicy = canonicalizeClusterNetworkPolicy(rawDesired.NetworkPolicy, nil, opts...)
		rawDesired.IPAllocationPolicy = canonicalizeClusterIPAllocationPolicy(rawDesired.IPAllocationPolicy, nil, opts...)
		rawDesired.MasterAuthorizedNetworksConfig = canonicalizeClusterMasterAuthorizedNetworksConfig(rawDesired.MasterAuthorizedNetworksConfig, nil, opts...)
		rawDesired.BinaryAuthorization = canonicalizeClusterBinaryAuthorization(rawDesired.BinaryAuthorization, nil, opts...)
		rawDesired.Autoscaling = canonicalizeClusterAutoscaling(rawDesired.Autoscaling, nil, opts...)
		rawDesired.NetworkConfig = canonicalizeClusterNetworkConfig(rawDesired.NetworkConfig, nil, opts...)
		rawDesired.MaintenancePolicy = canonicalizeClusterMaintenancePolicy(rawDesired.MaintenancePolicy, nil, opts...)
		rawDesired.DefaultMaxPodsConstraint = canonicalizeClusterDefaultMaxPodsConstraint(rawDesired.DefaultMaxPodsConstraint, nil, opts...)
		rawDesired.ResourceUsageExportConfig = canonicalizeClusterResourceUsageExportConfig(rawDesired.ResourceUsageExportConfig, nil, opts...)
		rawDesired.AuthenticatorGroupsConfig = canonicalizeClusterAuthenticatorGroupsConfig(rawDesired.AuthenticatorGroupsConfig, nil, opts...)
		rawDesired.PrivateClusterConfig = canonicalizeClusterPrivateClusterConfig(rawDesired.PrivateClusterConfig, nil, opts...)
		rawDesired.DatabaseEncryption = canonicalizeClusterDatabaseEncryption(rawDesired.DatabaseEncryption, nil, opts...)
		rawDesired.VerticalPodAutoscaling = canonicalizeClusterVerticalPodAutoscaling(rawDesired.VerticalPodAutoscaling, nil, opts...)
		rawDesired.ShieldedNodes = canonicalizeClusterShieldedNodes(rawDesired.ShieldedNodes, nil, opts...)
		rawDesired.Autopilot = canonicalizeClusterAutopilot(rawDesired.Autopilot, nil, opts...)
		rawDesired.NodeConfig = canonicalizeClusterNodeConfig(rawDesired.NodeConfig, nil, opts...)
		rawDesired.ReleaseChannel = canonicalizeClusterReleaseChannel(rawDesired.ReleaseChannel, nil, opts...)
		rawDesired.WorkloadIdentityConfig = canonicalizeClusterWorkloadIdentityConfig(rawDesired.WorkloadIdentityConfig, nil, opts...)
		rawDesired.NotificationConfig = canonicalizeClusterNotificationConfig(rawDesired.NotificationConfig, nil, opts...)
		rawDesired.ConfidentialNodes = canonicalizeClusterConfidentialNodes(rawDesired.ConfidentialNodes, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.InitialNodeCount) {
		rawDesired.InitialNodeCount = rawInitial.InitialNodeCount
	}
	rawDesired.MasterAuth = canonicalizeClusterMasterAuth(rawDesired.MasterAuth, rawInitial.MasterAuth, opts...)
	if dcl.StringCanonicalize(rawDesired.LoggingService, rawInitial.LoggingService) {
		rawDesired.LoggingService = rawInitial.LoggingService
	}
	if dcl.StringCanonicalize(rawDesired.MonitoringService, rawInitial.MonitoringService) {
		rawDesired.MonitoringService = rawInitial.MonitoringService
	}
	if dcl.StringCanonicalize(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.StringCanonicalize(rawDesired.ClusterIPv4Cidr, rawInitial.ClusterIPv4Cidr) {
		rawDesired.ClusterIPv4Cidr = rawInitial.ClusterIPv4Cidr
	}
	rawDesired.AddonsConfig = canonicalizeClusterAddonsConfig(rawDesired.AddonsConfig, rawInitial.AddonsConfig, opts...)
	if dcl.StringCanonicalize(rawDesired.Subnetwork, rawInitial.Subnetwork) {
		rawDesired.Subnetwork = rawInitial.Subnetwork
	}
	if dcl.IsZeroValue(rawDesired.NodePools) {
		rawDesired.NodePools = rawInitial.NodePools
	}
	if dcl.IsZeroValue(rawDesired.Locations) {
		rawDesired.Locations = rawInitial.Locations
	}
	if dcl.BoolCanonicalize(rawDesired.EnableKubernetesAlpha, rawInitial.EnableKubernetesAlpha) {
		rawDesired.EnableKubernetesAlpha = rawInitial.EnableKubernetesAlpha
	}
	if dcl.IsZeroValue(rawDesired.ResourceLabels) {
		rawDesired.ResourceLabels = rawInitial.ResourceLabels
	}
	if dcl.StringCanonicalize(rawDesired.LabelFingerprint, rawInitial.LabelFingerprint) {
		rawDesired.LabelFingerprint = rawInitial.LabelFingerprint
	}
	rawDesired.LegacyAbac = canonicalizeClusterLegacyAbac(rawDesired.LegacyAbac, rawInitial.LegacyAbac, opts...)
	rawDesired.NetworkPolicy = canonicalizeClusterNetworkPolicy(rawDesired.NetworkPolicy, rawInitial.NetworkPolicy, opts...)
	rawDesired.IPAllocationPolicy = canonicalizeClusterIPAllocationPolicy(rawDesired.IPAllocationPolicy, rawInitial.IPAllocationPolicy, opts...)
	rawDesired.MasterAuthorizedNetworksConfig = canonicalizeClusterMasterAuthorizedNetworksConfig(rawDesired.MasterAuthorizedNetworksConfig, rawInitial.MasterAuthorizedNetworksConfig, opts...)
	rawDesired.BinaryAuthorization = canonicalizeClusterBinaryAuthorization(rawDesired.BinaryAuthorization, rawInitial.BinaryAuthorization, opts...)
	rawDesired.Autoscaling = canonicalizeClusterAutoscaling(rawDesired.Autoscaling, rawInitial.Autoscaling, opts...)
	rawDesired.NetworkConfig = canonicalizeClusterNetworkConfig(rawDesired.NetworkConfig, rawInitial.NetworkConfig, opts...)
	rawDesired.MaintenancePolicy = canonicalizeClusterMaintenancePolicy(rawDesired.MaintenancePolicy, rawInitial.MaintenancePolicy, opts...)
	rawDesired.DefaultMaxPodsConstraint = canonicalizeClusterDefaultMaxPodsConstraint(rawDesired.DefaultMaxPodsConstraint, rawInitial.DefaultMaxPodsConstraint, opts...)
	rawDesired.ResourceUsageExportConfig = canonicalizeClusterResourceUsageExportConfig(rawDesired.ResourceUsageExportConfig, rawInitial.ResourceUsageExportConfig, opts...)
	rawDesired.AuthenticatorGroupsConfig = canonicalizeClusterAuthenticatorGroupsConfig(rawDesired.AuthenticatorGroupsConfig, rawInitial.AuthenticatorGroupsConfig, opts...)
	rawDesired.PrivateClusterConfig = canonicalizeClusterPrivateClusterConfig(rawDesired.PrivateClusterConfig, rawInitial.PrivateClusterConfig, opts...)
	rawDesired.DatabaseEncryption = canonicalizeClusterDatabaseEncryption(rawDesired.DatabaseEncryption, rawInitial.DatabaseEncryption, opts...)
	rawDesired.VerticalPodAutoscaling = canonicalizeClusterVerticalPodAutoscaling(rawDesired.VerticalPodAutoscaling, rawInitial.VerticalPodAutoscaling, opts...)
	rawDesired.ShieldedNodes = canonicalizeClusterShieldedNodes(rawDesired.ShieldedNodes, rawInitial.ShieldedNodes, opts...)
	if dcl.StringCanonicalize(rawDesired.Endpoint, rawInitial.Endpoint) {
		rawDesired.Endpoint = rawInitial.Endpoint
	}
	if dcl.MatchingSemver(rawDesired.MasterVersion, rawInitial.MasterVersion) {
		rawDesired.MasterVersion = rawInitial.MasterVersion
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.Status) {
		rawDesired.Status = rawInitial.Status
	}
	if dcl.StringCanonicalize(rawDesired.StatusMessage, rawInitial.StatusMessage) {
		rawDesired.StatusMessage = rawInitial.StatusMessage
	}
	if dcl.IsZeroValue(rawDesired.NodeIPv4CidrSize) {
		rawDesired.NodeIPv4CidrSize = rawInitial.NodeIPv4CidrSize
	}
	if dcl.StringCanonicalize(rawDesired.ServicesIPv4Cidr, rawInitial.ServicesIPv4Cidr) {
		rawDesired.ServicesIPv4Cidr = rawInitial.ServicesIPv4Cidr
	}
	if dcl.IsZeroValue(rawDesired.ExpireTime) {
		rawDesired.ExpireTime = rawInitial.ExpireTime
	}
	if dcl.StringCanonicalize(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.BoolCanonicalize(rawDesired.EnableTPU, rawInitial.EnableTPU) {
		rawDesired.EnableTPU = rawInitial.EnableTPU
	}
	if dcl.StringCanonicalize(rawDesired.TPUIPv4CidrBlock, rawInitial.TPUIPv4CidrBlock) {
		rawDesired.TPUIPv4CidrBlock = rawInitial.TPUIPv4CidrBlock
	}
	if dcl.IsZeroValue(rawDesired.Conditions) {
		rawDesired.Conditions = rawInitial.Conditions
	}
	rawDesired.Autopilot = canonicalizeClusterAutopilot(rawDesired.Autopilot, rawInitial.Autopilot, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	rawDesired.NodeConfig = canonicalizeClusterNodeConfig(rawDesired.NodeConfig, rawInitial.NodeConfig, opts...)
	rawDesired.ReleaseChannel = canonicalizeClusterReleaseChannel(rawDesired.ReleaseChannel, rawInitial.ReleaseChannel, opts...)
	rawDesired.WorkloadIdentityConfig = canonicalizeClusterWorkloadIdentityConfig(rawDesired.WorkloadIdentityConfig, rawInitial.WorkloadIdentityConfig, opts...)
	rawDesired.NotificationConfig = canonicalizeClusterNotificationConfig(rawDesired.NotificationConfig, rawInitial.NotificationConfig, opts...)
	rawDesired.ConfidentialNodes = canonicalizeClusterConfidentialNodes(rawDesired.ConfidentialNodes, rawInitial.ConfidentialNodes, opts...)
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.StringCanonicalize(rawDesired.Zone, rawInitial.Zone) {
		rawDesired.Zone = rawInitial.Zone
	}
	if dcl.StringCanonicalize(rawDesired.InitialClusterVersion, rawInitial.InitialClusterVersion) {
		rawDesired.InitialClusterVersion = rawInitial.InitialClusterVersion
	}
	if dcl.StringCanonicalize(rawDesired.CurrentMasterVersion, rawInitial.CurrentMasterVersion) {
		rawDesired.CurrentMasterVersion = rawInitial.CurrentMasterVersion
	}
	if dcl.StringCanonicalize(rawDesired.CurrentNodeVersion, rawInitial.CurrentNodeVersion) {
		rawDesired.CurrentNodeVersion = rawInitial.CurrentNodeVersion
	}
	if dcl.IsZeroValue(rawDesired.InstanceGroupUrls) {
		rawDesired.InstanceGroupUrls = rawInitial.InstanceGroupUrls
	}
	if dcl.IsZeroValue(rawDesired.CurrentNodeCount) {
		rawDesired.CurrentNodeCount = rawInitial.CurrentNodeCount
	}
	if dcl.StringCanonicalize(rawDesired.Id, rawInitial.Id) {
		rawDesired.Id = rawInitial.Id
	}

	return rawDesired, nil
}

func canonicalizeClusterNewState(c *Client, rawNew, rawDesired *Cluster) (*Cluster, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.InitialNodeCount) && dcl.IsEmptyValueIndirect(rawDesired.InitialNodeCount) {
		rawNew.InitialNodeCount = rawDesired.InitialNodeCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MasterAuth) && dcl.IsEmptyValueIndirect(rawDesired.MasterAuth) {
		rawNew.MasterAuth = rawDesired.MasterAuth
	} else {
		rawNew.MasterAuth = canonicalizeNewClusterMasterAuth(c, rawDesired.MasterAuth, rawNew.MasterAuth)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LoggingService) && dcl.IsEmptyValueIndirect(rawDesired.LoggingService) {
		rawNew.LoggingService = rawDesired.LoggingService
	} else {
		if dcl.StringCanonicalize(rawDesired.LoggingService, rawNew.LoggingService) {
			rawNew.LoggingService = rawDesired.LoggingService
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MonitoringService) && dcl.IsEmptyValueIndirect(rawDesired.MonitoringService) {
		rawNew.MonitoringService = rawDesired.MonitoringService
	} else {
		if dcl.StringCanonicalize(rawDesired.MonitoringService, rawNew.MonitoringService) {
			rawNew.MonitoringService = rawDesired.MonitoringService
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.StringCanonicalize(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClusterIPv4Cidr) && dcl.IsEmptyValueIndirect(rawDesired.ClusterIPv4Cidr) {
		rawNew.ClusterIPv4Cidr = rawDesired.ClusterIPv4Cidr
	} else {
		if dcl.StringCanonicalize(rawDesired.ClusterIPv4Cidr, rawNew.ClusterIPv4Cidr) {
			rawNew.ClusterIPv4Cidr = rawDesired.ClusterIPv4Cidr
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AddonsConfig) && dcl.IsEmptyValueIndirect(rawDesired.AddonsConfig) {
		rawNew.AddonsConfig = rawDesired.AddonsConfig
	} else {
		rawNew.AddonsConfig = canonicalizeNewClusterAddonsConfig(c, rawDesired.AddonsConfig, rawNew.AddonsConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Subnetwork) && dcl.IsEmptyValueIndirect(rawDesired.Subnetwork) {
		rawNew.Subnetwork = rawDesired.Subnetwork
	} else {
		if dcl.StringCanonicalize(rawDesired.Subnetwork, rawNew.Subnetwork) {
			rawNew.Subnetwork = rawDesired.Subnetwork
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NodePools) && dcl.IsEmptyValueIndirect(rawDesired.NodePools) {
		rawNew.NodePools = rawDesired.NodePools
	} else {
		rawNew.NodePools = canonicalizeNewClusterNodePoolsSlice(c, rawDesired.NodePools, rawNew.NodePools)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Locations) && dcl.IsEmptyValueIndirect(rawDesired.Locations) {
		rawNew.Locations = rawDesired.Locations
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableKubernetesAlpha) && dcl.IsEmptyValueIndirect(rawDesired.EnableKubernetesAlpha) {
		rawNew.EnableKubernetesAlpha = rawDesired.EnableKubernetesAlpha
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableKubernetesAlpha, rawNew.EnableKubernetesAlpha) {
			rawNew.EnableKubernetesAlpha = rawDesired.EnableKubernetesAlpha
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourceLabels) && dcl.IsEmptyValueIndirect(rawDesired.ResourceLabels) {
		rawNew.ResourceLabels = rawDesired.ResourceLabels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LabelFingerprint) && dcl.IsEmptyValueIndirect(rawDesired.LabelFingerprint) {
		rawNew.LabelFingerprint = rawDesired.LabelFingerprint
	} else {
		if dcl.StringCanonicalize(rawDesired.LabelFingerprint, rawNew.LabelFingerprint) {
			rawNew.LabelFingerprint = rawDesired.LabelFingerprint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LegacyAbac) && dcl.IsEmptyValueIndirect(rawDesired.LegacyAbac) {
		rawNew.LegacyAbac = rawDesired.LegacyAbac
	} else {
		rawNew.LegacyAbac = canonicalizeNewClusterLegacyAbac(c, rawDesired.LegacyAbac, rawNew.LegacyAbac)
	}

	if dcl.IsEmptyValueIndirect(rawNew.NetworkPolicy) && dcl.IsEmptyValueIndirect(rawDesired.NetworkPolicy) {
		rawNew.NetworkPolicy = rawDesired.NetworkPolicy
	} else {
		rawNew.NetworkPolicy = canonicalizeNewClusterNetworkPolicy(c, rawDesired.NetworkPolicy, rawNew.NetworkPolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPAllocationPolicy) && dcl.IsEmptyValueIndirect(rawDesired.IPAllocationPolicy) {
		rawNew.IPAllocationPolicy = rawDesired.IPAllocationPolicy
	} else {
		rawNew.IPAllocationPolicy = canonicalizeNewClusterIPAllocationPolicy(c, rawDesired.IPAllocationPolicy, rawNew.IPAllocationPolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MasterAuthorizedNetworksConfig) && dcl.IsEmptyValueIndirect(rawDesired.MasterAuthorizedNetworksConfig) {
		rawNew.MasterAuthorizedNetworksConfig = rawDesired.MasterAuthorizedNetworksConfig
	} else {
		rawNew.MasterAuthorizedNetworksConfig = canonicalizeNewClusterMasterAuthorizedNetworksConfig(c, rawDesired.MasterAuthorizedNetworksConfig, rawNew.MasterAuthorizedNetworksConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.BinaryAuthorization) && dcl.IsEmptyValueIndirect(rawDesired.BinaryAuthorization) {
		rawNew.BinaryAuthorization = rawDesired.BinaryAuthorization
	} else {
		rawNew.BinaryAuthorization = canonicalizeNewClusterBinaryAuthorization(c, rawDesired.BinaryAuthorization, rawNew.BinaryAuthorization)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Autoscaling) && dcl.IsEmptyValueIndirect(rawDesired.Autoscaling) {
		rawNew.Autoscaling = rawDesired.Autoscaling
	} else {
		rawNew.Autoscaling = canonicalizeNewClusterAutoscaling(c, rawDesired.Autoscaling, rawNew.Autoscaling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.NetworkConfig) && dcl.IsEmptyValueIndirect(rawDesired.NetworkConfig) {
		rawNew.NetworkConfig = rawDesired.NetworkConfig
	} else {
		rawNew.NetworkConfig = canonicalizeNewClusterNetworkConfig(c, rawDesired.NetworkConfig, rawNew.NetworkConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaintenancePolicy) && dcl.IsEmptyValueIndirect(rawDesired.MaintenancePolicy) {
		rawNew.MaintenancePolicy = rawDesired.MaintenancePolicy
	} else {
		rawNew.MaintenancePolicy = canonicalizeNewClusterMaintenancePolicy(c, rawDesired.MaintenancePolicy, rawNew.MaintenancePolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultMaxPodsConstraint) && dcl.IsEmptyValueIndirect(rawDesired.DefaultMaxPodsConstraint) {
		rawNew.DefaultMaxPodsConstraint = rawDesired.DefaultMaxPodsConstraint
	} else {
		rawNew.DefaultMaxPodsConstraint = canonicalizeNewClusterDefaultMaxPodsConstraint(c, rawDesired.DefaultMaxPodsConstraint, rawNew.DefaultMaxPodsConstraint)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourceUsageExportConfig) && dcl.IsEmptyValueIndirect(rawDesired.ResourceUsageExportConfig) {
		rawNew.ResourceUsageExportConfig = rawDesired.ResourceUsageExportConfig
	} else {
		rawNew.ResourceUsageExportConfig = canonicalizeNewClusterResourceUsageExportConfig(c, rawDesired.ResourceUsageExportConfig, rawNew.ResourceUsageExportConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AuthenticatorGroupsConfig) && dcl.IsEmptyValueIndirect(rawDesired.AuthenticatorGroupsConfig) {
		rawNew.AuthenticatorGroupsConfig = rawDesired.AuthenticatorGroupsConfig
	} else {
		rawNew.AuthenticatorGroupsConfig = canonicalizeNewClusterAuthenticatorGroupsConfig(c, rawDesired.AuthenticatorGroupsConfig, rawNew.AuthenticatorGroupsConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PrivateClusterConfig) && dcl.IsEmptyValueIndirect(rawDesired.PrivateClusterConfig) {
		rawNew.PrivateClusterConfig = rawDesired.PrivateClusterConfig
	} else {
		rawNew.PrivateClusterConfig = canonicalizeNewClusterPrivateClusterConfig(c, rawDesired.PrivateClusterConfig, rawNew.PrivateClusterConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DatabaseEncryption) && dcl.IsEmptyValueIndirect(rawDesired.DatabaseEncryption) {
		rawNew.DatabaseEncryption = rawDesired.DatabaseEncryption
	} else {
		rawNew.DatabaseEncryption = canonicalizeNewClusterDatabaseEncryption(c, rawDesired.DatabaseEncryption, rawNew.DatabaseEncryption)
	}

	if dcl.IsEmptyValueIndirect(rawNew.VerticalPodAutoscaling) && dcl.IsEmptyValueIndirect(rawDesired.VerticalPodAutoscaling) {
		rawNew.VerticalPodAutoscaling = rawDesired.VerticalPodAutoscaling
	} else {
		rawNew.VerticalPodAutoscaling = canonicalizeNewClusterVerticalPodAutoscaling(c, rawDesired.VerticalPodAutoscaling, rawNew.VerticalPodAutoscaling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ShieldedNodes) && dcl.IsEmptyValueIndirect(rawDesired.ShieldedNodes) {
		rawNew.ShieldedNodes = rawDesired.ShieldedNodes
	} else {
		rawNew.ShieldedNodes = canonicalizeNewClusterShieldedNodes(c, rawDesired.ShieldedNodes, rawNew.ShieldedNodes)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Endpoint) && dcl.IsEmptyValueIndirect(rawDesired.Endpoint) {
		rawNew.Endpoint = rawDesired.Endpoint
	} else {
		if dcl.StringCanonicalize(rawDesired.Endpoint, rawNew.Endpoint) {
			rawNew.Endpoint = rawDesired.Endpoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MasterVersion) && dcl.IsEmptyValueIndirect(rawDesired.MasterVersion) {
		rawNew.MasterVersion = rawDesired.MasterVersion
	} else {
		if dcl.MatchingSemver(rawDesired.MasterVersion, rawNew.MasterVersion) {
			rawNew.MasterVersion = rawDesired.MasterVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StatusMessage) && dcl.IsEmptyValueIndirect(rawDesired.StatusMessage) {
		rawNew.StatusMessage = rawDesired.StatusMessage
	} else {
		if dcl.StringCanonicalize(rawDesired.StatusMessage, rawNew.StatusMessage) {
			rawNew.StatusMessage = rawDesired.StatusMessage
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NodeIPv4CidrSize) && dcl.IsEmptyValueIndirect(rawDesired.NodeIPv4CidrSize) {
		rawNew.NodeIPv4CidrSize = rawDesired.NodeIPv4CidrSize
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServicesIPv4Cidr) && dcl.IsEmptyValueIndirect(rawDesired.ServicesIPv4Cidr) {
		rawNew.ServicesIPv4Cidr = rawDesired.ServicesIPv4Cidr
	} else {
		if dcl.StringCanonicalize(rawDesired.ServicesIPv4Cidr, rawNew.ServicesIPv4Cidr) {
			rawNew.ServicesIPv4Cidr = rawDesired.ServicesIPv4Cidr
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpireTime) && dcl.IsEmptyValueIndirect(rawDesired.ExpireTime) {
		rawNew.ExpireTime = rawDesired.ExpireTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Location) && dcl.IsEmptyValueIndirect(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
		if dcl.StringCanonicalize(rawDesired.Location, rawNew.Location) {
			rawNew.Location = rawDesired.Location
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableTPU) && dcl.IsEmptyValueIndirect(rawDesired.EnableTPU) {
		rawNew.EnableTPU = rawDesired.EnableTPU
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableTPU, rawNew.EnableTPU) {
			rawNew.EnableTPU = rawDesired.EnableTPU
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TPUIPv4CidrBlock) && dcl.IsEmptyValueIndirect(rawDesired.TPUIPv4CidrBlock) {
		rawNew.TPUIPv4CidrBlock = rawDesired.TPUIPv4CidrBlock
	} else {
		if dcl.StringCanonicalize(rawDesired.TPUIPv4CidrBlock, rawNew.TPUIPv4CidrBlock) {
			rawNew.TPUIPv4CidrBlock = rawDesired.TPUIPv4CidrBlock
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Conditions) && dcl.IsEmptyValueIndirect(rawDesired.Conditions) {
		rawNew.Conditions = rawDesired.Conditions
	} else {
		rawNew.Conditions = canonicalizeNewClusterConditionsSlice(c, rawDesired.Conditions, rawNew.Conditions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Autopilot) && dcl.IsEmptyValueIndirect(rawDesired.Autopilot) {
		rawNew.Autopilot = rawDesired.Autopilot
	} else {
		rawNew.Autopilot = canonicalizeNewClusterAutopilot(c, rawDesired.Autopilot, rawNew.Autopilot)
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.NodeConfig) && dcl.IsEmptyValueIndirect(rawDesired.NodeConfig) {
		rawNew.NodeConfig = rawDesired.NodeConfig
	} else {
		rawNew.NodeConfig = canonicalizeNewClusterNodeConfig(c, rawDesired.NodeConfig, rawNew.NodeConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReleaseChannel) && dcl.IsEmptyValueIndirect(rawDesired.ReleaseChannel) {
		rawNew.ReleaseChannel = rawDesired.ReleaseChannel
	} else {
		rawNew.ReleaseChannel = canonicalizeNewClusterReleaseChannel(c, rawDesired.ReleaseChannel, rawNew.ReleaseChannel)
	}

	if dcl.IsEmptyValueIndirect(rawNew.WorkloadIdentityConfig) && dcl.IsEmptyValueIndirect(rawDesired.WorkloadIdentityConfig) {
		rawNew.WorkloadIdentityConfig = rawDesired.WorkloadIdentityConfig
	} else {
		rawNew.WorkloadIdentityConfig = canonicalizeNewClusterWorkloadIdentityConfig(c, rawDesired.WorkloadIdentityConfig, rawNew.WorkloadIdentityConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.NotificationConfig) && dcl.IsEmptyValueIndirect(rawDesired.NotificationConfig) {
		rawNew.NotificationConfig = rawDesired.NotificationConfig
	} else {
		rawNew.NotificationConfig = canonicalizeNewClusterNotificationConfig(c, rawDesired.NotificationConfig, rawNew.NotificationConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ConfidentialNodes) && dcl.IsEmptyValueIndirect(rawDesired.ConfidentialNodes) {
		rawNew.ConfidentialNodes = rawDesired.ConfidentialNodes
	} else {
		rawNew.ConfidentialNodes = canonicalizeNewClusterConfidentialNodes(c, rawDesired.ConfidentialNodes, rawNew.ConfidentialNodes)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Zone) && dcl.IsEmptyValueIndirect(rawDesired.Zone) {
		rawNew.Zone = rawDesired.Zone
	} else {
		if dcl.StringCanonicalize(rawDesired.Zone, rawNew.Zone) {
			rawNew.Zone = rawDesired.Zone
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.InitialClusterVersion) && dcl.IsEmptyValueIndirect(rawDesired.InitialClusterVersion) {
		rawNew.InitialClusterVersion = rawDesired.InitialClusterVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.InitialClusterVersion, rawNew.InitialClusterVersion) {
			rawNew.InitialClusterVersion = rawDesired.InitialClusterVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CurrentMasterVersion) && dcl.IsEmptyValueIndirect(rawDesired.CurrentMasterVersion) {
		rawNew.CurrentMasterVersion = rawDesired.CurrentMasterVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.CurrentMasterVersion, rawNew.CurrentMasterVersion) {
			rawNew.CurrentMasterVersion = rawDesired.CurrentMasterVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CurrentNodeVersion) && dcl.IsEmptyValueIndirect(rawDesired.CurrentNodeVersion) {
		rawNew.CurrentNodeVersion = rawDesired.CurrentNodeVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.CurrentNodeVersion, rawNew.CurrentNodeVersion) {
			rawNew.CurrentNodeVersion = rawDesired.CurrentNodeVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.InstanceGroupUrls) && dcl.IsEmptyValueIndirect(rawDesired.InstanceGroupUrls) {
		rawNew.InstanceGroupUrls = rawDesired.InstanceGroupUrls
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CurrentNodeCount) && dcl.IsEmptyValueIndirect(rawDesired.CurrentNodeCount) {
		rawNew.CurrentNodeCount = rawDesired.CurrentNodeCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	return rawNew, nil
}

func canonicalizeClusterMasterAuth(des, initial *ClusterMasterAuth, opts ...dcl.ApplyOption) *ClusterMasterAuth {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Username, initial.Username) || dcl.IsZeroValue(des.Username) {
		des.Username = initial.Username
	}
	if dcl.StringCanonicalize(des.Password, initial.Password) || dcl.IsZeroValue(des.Password) {
		des.Password = initial.Password
	}
	des.ClientCertificateConfig = canonicalizeClusterMasterAuthClientCertificateConfig(des.ClientCertificateConfig, initial.ClientCertificateConfig, opts...)
	if dcl.StringCanonicalize(des.ClusterCaCertificate, initial.ClusterCaCertificate) || dcl.IsZeroValue(des.ClusterCaCertificate) {
		des.ClusterCaCertificate = initial.ClusterCaCertificate
	}
	if dcl.StringCanonicalize(des.ClientCertificate, initial.ClientCertificate) || dcl.IsZeroValue(des.ClientCertificate) {
		des.ClientCertificate = initial.ClientCertificate
	}
	if dcl.StringCanonicalize(des.ClientKey, initial.ClientKey) || dcl.IsZeroValue(des.ClientKey) {
		des.ClientKey = initial.ClientKey
	}

	return des
}

func canonicalizeNewClusterMasterAuth(c *Client, des, nw *ClusterMasterAuth) *ClusterMasterAuth {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Username, nw.Username) {
		nw.Username = des.Username
	}
	if dcl.StringCanonicalize(des.Password, nw.Password) {
		nw.Password = des.Password
	}
	nw.ClientCertificateConfig = canonicalizeNewClusterMasterAuthClientCertificateConfig(c, des.ClientCertificateConfig, nw.ClientCertificateConfig)
	if dcl.StringCanonicalize(des.ClusterCaCertificate, nw.ClusterCaCertificate) {
		nw.ClusterCaCertificate = des.ClusterCaCertificate
	}
	if dcl.StringCanonicalize(des.ClientCertificate, nw.ClientCertificate) {
		nw.ClientCertificate = des.ClientCertificate
	}
	if dcl.StringCanonicalize(des.ClientKey, nw.ClientKey) {
		nw.ClientKey = des.ClientKey
	}

	return nw
}

func canonicalizeNewClusterMasterAuthSet(c *Client, des, nw []ClusterMasterAuth) []ClusterMasterAuth {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMasterAuth
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterMasterAuth(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterMasterAuthSlice(c *Client, des, nw []ClusterMasterAuth) []ClusterMasterAuth {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterMasterAuth
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMasterAuth(c, &d, &n))
	}

	return items
}

func canonicalizeClusterMasterAuthClientCertificateConfig(des, initial *ClusterMasterAuthClientCertificateConfig, opts ...dcl.ApplyOption) *ClusterMasterAuthClientCertificateConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.IssueClientCertificate, initial.IssueClientCertificate) || dcl.IsZeroValue(des.IssueClientCertificate) {
		des.IssueClientCertificate = initial.IssueClientCertificate
	}

	return des
}

func canonicalizeNewClusterMasterAuthClientCertificateConfig(c *Client, des, nw *ClusterMasterAuthClientCertificateConfig) *ClusterMasterAuthClientCertificateConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.IssueClientCertificate, nw.IssueClientCertificate) {
		nw.IssueClientCertificate = des.IssueClientCertificate
	}

	return nw
}

func canonicalizeNewClusterMasterAuthClientCertificateConfigSet(c *Client, des, nw []ClusterMasterAuthClientCertificateConfig) []ClusterMasterAuthClientCertificateConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMasterAuthClientCertificateConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterMasterAuthClientCertificateConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterMasterAuthClientCertificateConfigSlice(c *Client, des, nw []ClusterMasterAuthClientCertificateConfig) []ClusterMasterAuthClientCertificateConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterMasterAuthClientCertificateConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMasterAuthClientCertificateConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAddonsConfig(des, initial *ClusterAddonsConfig, opts ...dcl.ApplyOption) *ClusterAddonsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.HttpLoadBalancing = canonicalizeClusterAddonsConfigHttpLoadBalancing(des.HttpLoadBalancing, initial.HttpLoadBalancing, opts...)
	des.HorizontalPodAutoscaling = canonicalizeClusterAddonsConfigHorizontalPodAutoscaling(des.HorizontalPodAutoscaling, initial.HorizontalPodAutoscaling, opts...)
	des.KubernetesDashboard = canonicalizeClusterAddonsConfigKubernetesDashboard(des.KubernetesDashboard, initial.KubernetesDashboard, opts...)
	des.NetworkPolicyConfig = canonicalizeClusterAddonsConfigNetworkPolicyConfig(des.NetworkPolicyConfig, initial.NetworkPolicyConfig, opts...)
	des.CloudRunConfig = canonicalizeClusterAddonsConfigCloudRunConfig(des.CloudRunConfig, initial.CloudRunConfig, opts...)
	des.DnsCacheConfig = canonicalizeClusterAddonsConfigDnsCacheConfig(des.DnsCacheConfig, initial.DnsCacheConfig, opts...)
	des.ConfigConnectorConfig = canonicalizeClusterAddonsConfigConfigConnectorConfig(des.ConfigConnectorConfig, initial.ConfigConnectorConfig, opts...)
	des.GcePersistentDiskCsiDriverConfig = canonicalizeClusterAddonsConfigGcePersistentDiskCsiDriverConfig(des.GcePersistentDiskCsiDriverConfig, initial.GcePersistentDiskCsiDriverConfig, opts...)

	return des
}

func canonicalizeNewClusterAddonsConfig(c *Client, des, nw *ClusterAddonsConfig) *ClusterAddonsConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.HttpLoadBalancing = canonicalizeNewClusterAddonsConfigHttpLoadBalancing(c, des.HttpLoadBalancing, nw.HttpLoadBalancing)
	nw.HorizontalPodAutoscaling = canonicalizeNewClusterAddonsConfigHorizontalPodAutoscaling(c, des.HorizontalPodAutoscaling, nw.HorizontalPodAutoscaling)
	nw.KubernetesDashboard = canonicalizeNewClusterAddonsConfigKubernetesDashboard(c, des.KubernetesDashboard, nw.KubernetesDashboard)
	nw.NetworkPolicyConfig = canonicalizeNewClusterAddonsConfigNetworkPolicyConfig(c, des.NetworkPolicyConfig, nw.NetworkPolicyConfig)
	nw.CloudRunConfig = canonicalizeNewClusterAddonsConfigCloudRunConfig(c, des.CloudRunConfig, nw.CloudRunConfig)
	nw.DnsCacheConfig = canonicalizeNewClusterAddonsConfigDnsCacheConfig(c, des.DnsCacheConfig, nw.DnsCacheConfig)
	nw.ConfigConnectorConfig = canonicalizeNewClusterAddonsConfigConfigConnectorConfig(c, des.ConfigConnectorConfig, nw.ConfigConnectorConfig)
	nw.GcePersistentDiskCsiDriverConfig = canonicalizeNewClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, des.GcePersistentDiskCsiDriverConfig, nw.GcePersistentDiskCsiDriverConfig)

	return nw
}

func canonicalizeNewClusterAddonsConfigSet(c *Client, des, nw []ClusterAddonsConfig) []ClusterAddonsConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAddonsConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAddonsConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAddonsConfigSlice(c *Client, des, nw []ClusterAddonsConfig) []ClusterAddonsConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAddonsConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAddonsConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAddonsConfigHttpLoadBalancing(des, initial *ClusterAddonsConfigHttpLoadBalancing, opts ...dcl.ApplyOption) *ClusterAddonsConfigHttpLoadBalancing {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigHttpLoadBalancing(c *Client, des, nw *ClusterAddonsConfigHttpLoadBalancing) *ClusterAddonsConfigHttpLoadBalancing {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
	}

	return nw
}

func canonicalizeNewClusterAddonsConfigHttpLoadBalancingSet(c *Client, des, nw []ClusterAddonsConfigHttpLoadBalancing) []ClusterAddonsConfigHttpLoadBalancing {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAddonsConfigHttpLoadBalancing
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAddonsConfigHttpLoadBalancing(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAddonsConfigHttpLoadBalancingSlice(c *Client, des, nw []ClusterAddonsConfigHttpLoadBalancing) []ClusterAddonsConfigHttpLoadBalancing {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAddonsConfigHttpLoadBalancing
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAddonsConfigHttpLoadBalancing(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAddonsConfigHorizontalPodAutoscaling(des, initial *ClusterAddonsConfigHorizontalPodAutoscaling, opts ...dcl.ApplyOption) *ClusterAddonsConfigHorizontalPodAutoscaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigHorizontalPodAutoscaling(c *Client, des, nw *ClusterAddonsConfigHorizontalPodAutoscaling) *ClusterAddonsConfigHorizontalPodAutoscaling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
	}

	return nw
}

func canonicalizeNewClusterAddonsConfigHorizontalPodAutoscalingSet(c *Client, des, nw []ClusterAddonsConfigHorizontalPodAutoscaling) []ClusterAddonsConfigHorizontalPodAutoscaling {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAddonsConfigHorizontalPodAutoscaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAddonsConfigHorizontalPodAutoscaling(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAddonsConfigHorizontalPodAutoscalingSlice(c *Client, des, nw []ClusterAddonsConfigHorizontalPodAutoscaling) []ClusterAddonsConfigHorizontalPodAutoscaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAddonsConfigHorizontalPodAutoscaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAddonsConfigHorizontalPodAutoscaling(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAddonsConfigKubernetesDashboard(des, initial *ClusterAddonsConfigKubernetesDashboard, opts ...dcl.ApplyOption) *ClusterAddonsConfigKubernetesDashboard {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigKubernetesDashboard(c *Client, des, nw *ClusterAddonsConfigKubernetesDashboard) *ClusterAddonsConfigKubernetesDashboard {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
	}

	return nw
}

func canonicalizeNewClusterAddonsConfigKubernetesDashboardSet(c *Client, des, nw []ClusterAddonsConfigKubernetesDashboard) []ClusterAddonsConfigKubernetesDashboard {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAddonsConfigKubernetesDashboard
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAddonsConfigKubernetesDashboard(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAddonsConfigKubernetesDashboardSlice(c *Client, des, nw []ClusterAddonsConfigKubernetesDashboard) []ClusterAddonsConfigKubernetesDashboard {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAddonsConfigKubernetesDashboard
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAddonsConfigKubernetesDashboard(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAddonsConfigNetworkPolicyConfig(des, initial *ClusterAddonsConfigNetworkPolicyConfig, opts ...dcl.ApplyOption) *ClusterAddonsConfigNetworkPolicyConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigNetworkPolicyConfig(c *Client, des, nw *ClusterAddonsConfigNetworkPolicyConfig) *ClusterAddonsConfigNetworkPolicyConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
	}

	return nw
}

func canonicalizeNewClusterAddonsConfigNetworkPolicyConfigSet(c *Client, des, nw []ClusterAddonsConfigNetworkPolicyConfig) []ClusterAddonsConfigNetworkPolicyConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAddonsConfigNetworkPolicyConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAddonsConfigNetworkPolicyConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAddonsConfigNetworkPolicyConfigSlice(c *Client, des, nw []ClusterAddonsConfigNetworkPolicyConfig) []ClusterAddonsConfigNetworkPolicyConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAddonsConfigNetworkPolicyConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAddonsConfigNetworkPolicyConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAddonsConfigCloudRunConfig(des, initial *ClusterAddonsConfigCloudRunConfig, opts ...dcl.ApplyOption) *ClusterAddonsConfigCloudRunConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}
	if dcl.IsZeroValue(des.LoadBalancerType) {
		des.LoadBalancerType = initial.LoadBalancerType
	}

	return des
}

func canonicalizeNewClusterAddonsConfigCloudRunConfig(c *Client, des, nw *ClusterAddonsConfigCloudRunConfig) *ClusterAddonsConfigCloudRunConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
	}

	return nw
}

func canonicalizeNewClusterAddonsConfigCloudRunConfigSet(c *Client, des, nw []ClusterAddonsConfigCloudRunConfig) []ClusterAddonsConfigCloudRunConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAddonsConfigCloudRunConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAddonsConfigCloudRunConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAddonsConfigCloudRunConfigSlice(c *Client, des, nw []ClusterAddonsConfigCloudRunConfig) []ClusterAddonsConfigCloudRunConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAddonsConfigCloudRunConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAddonsConfigCloudRunConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAddonsConfigDnsCacheConfig(des, initial *ClusterAddonsConfigDnsCacheConfig, opts ...dcl.ApplyOption) *ClusterAddonsConfigDnsCacheConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigDnsCacheConfig(c *Client, des, nw *ClusterAddonsConfigDnsCacheConfig) *ClusterAddonsConfigDnsCacheConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterAddonsConfigDnsCacheConfigSet(c *Client, des, nw []ClusterAddonsConfigDnsCacheConfig) []ClusterAddonsConfigDnsCacheConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAddonsConfigDnsCacheConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAddonsConfigDnsCacheConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAddonsConfigDnsCacheConfigSlice(c *Client, des, nw []ClusterAddonsConfigDnsCacheConfig) []ClusterAddonsConfigDnsCacheConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAddonsConfigDnsCacheConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAddonsConfigDnsCacheConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAddonsConfigConfigConnectorConfig(des, initial *ClusterAddonsConfigConfigConnectorConfig, opts ...dcl.ApplyOption) *ClusterAddonsConfigConfigConnectorConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigConfigConnectorConfig(c *Client, des, nw *ClusterAddonsConfigConfigConnectorConfig) *ClusterAddonsConfigConfigConnectorConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterAddonsConfigConfigConnectorConfigSet(c *Client, des, nw []ClusterAddonsConfigConfigConnectorConfig) []ClusterAddonsConfigConfigConnectorConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAddonsConfigConfigConnectorConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAddonsConfigConfigConnectorConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAddonsConfigConfigConnectorConfigSlice(c *Client, des, nw []ClusterAddonsConfigConfigConnectorConfig) []ClusterAddonsConfigConfigConnectorConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAddonsConfigConfigConnectorConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAddonsConfigConfigConnectorConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAddonsConfigGcePersistentDiskCsiDriverConfig(des, initial *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig, opts ...dcl.ApplyOption) *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c *Client, des, nw *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterAddonsConfigGcePersistentDiskCsiDriverConfigSet(c *Client, des, nw []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAddonsConfigGcePersistentDiskCsiDriverConfigSlice(c *Client, des, nw []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePools(des, initial *ClusterNodePools, opts ...dcl.ApplyOption) *ClusterNodePools {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	des.Config = canonicalizeClusterNodePoolsConfig(des.Config, initial.Config, opts...)
	if dcl.IsZeroValue(des.InitialNodeCount) {
		des.InitialNodeCount = initial.InitialNodeCount
	}
	if dcl.IsZeroValue(des.Locations) {
		des.Locations = initial.Locations
	}
	if dcl.StringCanonicalize(des.SelfLink, initial.SelfLink) || dcl.IsZeroValue(des.SelfLink) {
		des.SelfLink = initial.SelfLink
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		des.Version = initial.Version
	}
	if dcl.IsZeroValue(des.InstanceGroupUrls) {
		des.InstanceGroupUrls = initial.InstanceGroupUrls
	}
	if dcl.IsZeroValue(des.Status) {
		des.Status = initial.Status
	}
	if dcl.StringCanonicalize(des.StatusMessage, initial.StatusMessage) || dcl.IsZeroValue(des.StatusMessage) {
		des.StatusMessage = initial.StatusMessage
	}
	des.Autoscaling = canonicalizeClusterNodePoolsAutoscaling(des.Autoscaling, initial.Autoscaling, opts...)
	des.Management = canonicalizeClusterNodePoolsManagement(des.Management, initial.Management, opts...)
	des.MaxPodsConstraint = canonicalizeClusterNodePoolsMaxPodsConstraint(des.MaxPodsConstraint, initial.MaxPodsConstraint, opts...)
	if dcl.IsZeroValue(des.Conditions) {
		des.Conditions = initial.Conditions
	}
	if dcl.IsZeroValue(des.PodIPv4CidrSize) {
		des.PodIPv4CidrSize = initial.PodIPv4CidrSize
	}
	des.UpgradeSettings = canonicalizeClusterNodePoolsUpgradeSettings(des.UpgradeSettings, initial.UpgradeSettings, opts...)

	return des
}

func canonicalizeNewClusterNodePools(c *Client, des, nw *ClusterNodePools) *ClusterNodePools {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.Config = canonicalizeNewClusterNodePoolsConfig(c, des.Config, nw.Config)
	if dcl.StringCanonicalize(des.SelfLink, nw.SelfLink) {
		nw.SelfLink = des.SelfLink
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}
	if dcl.StringCanonicalize(des.StatusMessage, nw.StatusMessage) {
		nw.StatusMessage = des.StatusMessage
	}
	nw.Autoscaling = canonicalizeNewClusterNodePoolsAutoscaling(c, des.Autoscaling, nw.Autoscaling)
	nw.Management = canonicalizeNewClusterNodePoolsManagement(c, des.Management, nw.Management)
	nw.MaxPodsConstraint = canonicalizeNewClusterNodePoolsMaxPodsConstraint(c, des.MaxPodsConstraint, nw.MaxPodsConstraint)
	nw.Conditions = canonicalizeNewClusterNodePoolsConditionsSlice(c, des.Conditions, nw.Conditions)
	nw.UpgradeSettings = canonicalizeNewClusterNodePoolsUpgradeSettings(c, des.UpgradeSettings, nw.UpgradeSettings)

	return nw
}

func canonicalizeNewClusterNodePoolsSet(c *Client, des, nw []ClusterNodePools) []ClusterNodePools {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePools
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePools(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsSlice(c *Client, des, nw []ClusterNodePools) []ClusterNodePools {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePools
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePools(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConfig(des, initial *ClusterNodePoolsConfig, opts ...dcl.ApplyOption) *ClusterNodePoolsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		des.MachineType = initial.MachineType
	}
	if dcl.IsZeroValue(des.DiskSizeGb) {
		des.DiskSizeGb = initial.DiskSizeGb
	}
	if dcl.IsZeroValue(des.OAuthScopes) {
		des.OAuthScopes = initial.OAuthScopes
	}
	if dcl.StringCanonicalize(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		des.ServiceAccount = initial.ServiceAccount
	}
	if dcl.IsZeroValue(des.Metadata) {
		des.Metadata = initial.Metadata
	}
	if dcl.StringCanonicalize(des.ImageType, initial.ImageType) || dcl.IsZeroValue(des.ImageType) {
		des.ImageType = initial.ImageType
	}
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.LocalSsdCount) {
		des.LocalSsdCount = initial.LocalSsdCount
	}
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	}
	if dcl.BoolCanonicalize(des.Preemptible, initial.Preemptible) || dcl.IsZeroValue(des.Preemptible) {
		des.Preemptible = initial.Preemptible
	}
	if dcl.IsZeroValue(des.Accelerators) {
		des.Accelerators = initial.Accelerators
	}
	if dcl.StringCanonicalize(des.DiskType, initial.DiskType) || dcl.IsZeroValue(des.DiskType) {
		des.DiskType = initial.DiskType
	}
	if dcl.StringCanonicalize(des.MinCpuPlatform, initial.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
		des.MinCpuPlatform = initial.MinCpuPlatform
	}
	des.WorkloadMetadataConfig = canonicalizeClusterNodePoolsConfigWorkloadMetadataConfig(des.WorkloadMetadataConfig, initial.WorkloadMetadataConfig, opts...)
	if dcl.IsZeroValue(des.Taints) {
		des.Taints = initial.Taints
	}
	des.SandboxConfig = canonicalizeClusterNodePoolsConfigSandboxConfig(des.SandboxConfig, initial.SandboxConfig, opts...)
	if dcl.StringCanonicalize(des.NodeGroup, initial.NodeGroup) || dcl.IsZeroValue(des.NodeGroup) {
		des.NodeGroup = initial.NodeGroup
	}
	des.ReservationAffinity = canonicalizeClusterNodePoolsConfigReservationAffinity(des.ReservationAffinity, initial.ReservationAffinity, opts...)
	des.ShieldedInstanceConfig = canonicalizeClusterNodePoolsConfigShieldedInstanceConfig(des.ShieldedInstanceConfig, initial.ShieldedInstanceConfig, opts...)
	des.LinuxNodeConfig = canonicalizeClusterNodePoolsConfigLinuxNodeConfig(des.LinuxNodeConfig, initial.LinuxNodeConfig, opts...)
	des.KubeletConfig = canonicalizeClusterNodePoolsConfigKubeletConfig(des.KubeletConfig, initial.KubeletConfig, opts...)
	if dcl.StringCanonicalize(des.BootDiskKmsKey, initial.BootDiskKmsKey) || dcl.IsZeroValue(des.BootDiskKmsKey) {
		des.BootDiskKmsKey = initial.BootDiskKmsKey
	}

	return des
}

func canonicalizeNewClusterNodePoolsConfig(c *Client, des, nw *ClusterNodePoolsConfig) *ClusterNodePoolsConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	if dcl.StringCanonicalize(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	if dcl.StringCanonicalize(des.ImageType, nw.ImageType) {
		nw.ImageType = des.ImageType
	}
	if dcl.BoolCanonicalize(des.Preemptible, nw.Preemptible) {
		nw.Preemptible = des.Preemptible
	}
	nw.Accelerators = canonicalizeNewClusterNodePoolsConfigAcceleratorsSlice(c, des.Accelerators, nw.Accelerators)
	if dcl.StringCanonicalize(des.DiskType, nw.DiskType) {
		nw.DiskType = des.DiskType
	}
	if dcl.StringCanonicalize(des.MinCpuPlatform, nw.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}
	nw.WorkloadMetadataConfig = canonicalizeNewClusterNodePoolsConfigWorkloadMetadataConfig(c, des.WorkloadMetadataConfig, nw.WorkloadMetadataConfig)
	nw.Taints = canonicalizeNewClusterNodePoolsConfigTaintsSlice(c, des.Taints, nw.Taints)
	nw.SandboxConfig = canonicalizeNewClusterNodePoolsConfigSandboxConfig(c, des.SandboxConfig, nw.SandboxConfig)
	if dcl.StringCanonicalize(des.NodeGroup, nw.NodeGroup) {
		nw.NodeGroup = des.NodeGroup
	}
	nw.ReservationAffinity = canonicalizeNewClusterNodePoolsConfigReservationAffinity(c, des.ReservationAffinity, nw.ReservationAffinity)
	nw.ShieldedInstanceConfig = canonicalizeNewClusterNodePoolsConfigShieldedInstanceConfig(c, des.ShieldedInstanceConfig, nw.ShieldedInstanceConfig)
	nw.LinuxNodeConfig = canonicalizeNewClusterNodePoolsConfigLinuxNodeConfig(c, des.LinuxNodeConfig, nw.LinuxNodeConfig)
	nw.KubeletConfig = canonicalizeNewClusterNodePoolsConfigKubeletConfig(c, des.KubeletConfig, nw.KubeletConfig)
	if dcl.StringCanonicalize(des.BootDiskKmsKey, nw.BootDiskKmsKey) {
		nw.BootDiskKmsKey = des.BootDiskKmsKey
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConfigSet(c *Client, des, nw []ClusterNodePoolsConfig) []ClusterNodePoolsConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConfigSlice(c *Client, des, nw []ClusterNodePoolsConfig) []ClusterNodePoolsConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConfigAccelerators(des, initial *ClusterNodePoolsConfigAccelerators, opts ...dcl.ApplyOption) *ClusterNodePoolsConfigAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AcceleratorCount) {
		des.AcceleratorCount = initial.AcceleratorCount
	}
	if dcl.StringCanonicalize(des.AcceleratorType, initial.AcceleratorType) || dcl.IsZeroValue(des.AcceleratorType) {
		des.AcceleratorType = initial.AcceleratorType
	}

	return des
}

func canonicalizeNewClusterNodePoolsConfigAccelerators(c *Client, des, nw *ClusterNodePoolsConfigAccelerators) *ClusterNodePoolsConfigAccelerators {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AcceleratorType, nw.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConfigAcceleratorsSet(c *Client, des, nw []ClusterNodePoolsConfigAccelerators) []ClusterNodePoolsConfigAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConfigAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConfigAccelerators(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConfigAcceleratorsSlice(c *Client, des, nw []ClusterNodePoolsConfigAccelerators) []ClusterNodePoolsConfigAccelerators {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConfigAccelerators
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConfigAccelerators(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConfigWorkloadMetadataConfig(des, initial *ClusterNodePoolsConfigWorkloadMetadataConfig, opts ...dcl.ApplyOption) *ClusterNodePoolsConfigWorkloadMetadataConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	}

	return des
}

func canonicalizeNewClusterNodePoolsConfigWorkloadMetadataConfig(c *Client, des, nw *ClusterNodePoolsConfigWorkloadMetadataConfig) *ClusterNodePoolsConfigWorkloadMetadataConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConfigWorkloadMetadataConfigSet(c *Client, des, nw []ClusterNodePoolsConfigWorkloadMetadataConfig) []ClusterNodePoolsConfigWorkloadMetadataConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConfigWorkloadMetadataConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConfigWorkloadMetadataConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConfigWorkloadMetadataConfigSlice(c *Client, des, nw []ClusterNodePoolsConfigWorkloadMetadataConfig) []ClusterNodePoolsConfigWorkloadMetadataConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConfigWorkloadMetadataConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConfigWorkloadMetadataConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConfigTaints(des, initial *ClusterNodePoolsConfigTaints, opts ...dcl.ApplyOption) *ClusterNodePoolsConfigTaints {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}
	if dcl.IsZeroValue(des.Effect) {
		des.Effect = initial.Effect
	}

	return des
}

func canonicalizeNewClusterNodePoolsConfigTaints(c *Client, des, nw *ClusterNodePoolsConfigTaints) *ClusterNodePoolsConfigTaints {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConfigTaintsSet(c *Client, des, nw []ClusterNodePoolsConfigTaints) []ClusterNodePoolsConfigTaints {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConfigTaints
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConfigTaints(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConfigTaintsSlice(c *Client, des, nw []ClusterNodePoolsConfigTaints) []ClusterNodePoolsConfigTaints {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConfigTaints
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConfigTaints(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConfigSandboxConfig(des, initial *ClusterNodePoolsConfigSandboxConfig, opts ...dcl.ApplyOption) *ClusterNodePoolsConfigSandboxConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewClusterNodePoolsConfigSandboxConfig(c *Client, des, nw *ClusterNodePoolsConfigSandboxConfig) *ClusterNodePoolsConfigSandboxConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConfigSandboxConfigSet(c *Client, des, nw []ClusterNodePoolsConfigSandboxConfig) []ClusterNodePoolsConfigSandboxConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConfigSandboxConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConfigSandboxConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConfigSandboxConfigSlice(c *Client, des, nw []ClusterNodePoolsConfigSandboxConfig) []ClusterNodePoolsConfigSandboxConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConfigSandboxConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConfigSandboxConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConfigReservationAffinity(des, initial *ClusterNodePoolsConfigReservationAffinity, opts ...dcl.ApplyOption) *ClusterNodePoolsConfigReservationAffinity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ConsumeReservationType) {
		des.ConsumeReservationType = initial.ConsumeReservationType
	}
	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Values) {
		des.Values = initial.Values
	}

	return des
}

func canonicalizeNewClusterNodePoolsConfigReservationAffinity(c *Client, des, nw *ClusterNodePoolsConfigReservationAffinity) *ClusterNodePoolsConfigReservationAffinity {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConfigReservationAffinitySet(c *Client, des, nw []ClusterNodePoolsConfigReservationAffinity) []ClusterNodePoolsConfigReservationAffinity {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConfigReservationAffinity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConfigReservationAffinity(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConfigReservationAffinitySlice(c *Client, des, nw []ClusterNodePoolsConfigReservationAffinity) []ClusterNodePoolsConfigReservationAffinity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConfigReservationAffinity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConfigReservationAffinity(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConfigShieldedInstanceConfig(des, initial *ClusterNodePoolsConfigShieldedInstanceConfig, opts ...dcl.ApplyOption) *ClusterNodePoolsConfigShieldedInstanceConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.EnableSecureBoot, initial.EnableSecureBoot) || dcl.IsZeroValue(des.EnableSecureBoot) {
		des.EnableSecureBoot = initial.EnableSecureBoot
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, initial.EnableIntegrityMonitoring) || dcl.IsZeroValue(des.EnableIntegrityMonitoring) {
		des.EnableIntegrityMonitoring = initial.EnableIntegrityMonitoring
	}

	return des
}

func canonicalizeNewClusterNodePoolsConfigShieldedInstanceConfig(c *Client, des, nw *ClusterNodePoolsConfigShieldedInstanceConfig) *ClusterNodePoolsConfigShieldedInstanceConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnableSecureBoot, nw.EnableSecureBoot) {
		nw.EnableSecureBoot = des.EnableSecureBoot
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, nw.EnableIntegrityMonitoring) {
		nw.EnableIntegrityMonitoring = des.EnableIntegrityMonitoring
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConfigShieldedInstanceConfigSet(c *Client, des, nw []ClusterNodePoolsConfigShieldedInstanceConfig) []ClusterNodePoolsConfigShieldedInstanceConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConfigShieldedInstanceConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConfigShieldedInstanceConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConfigShieldedInstanceConfigSlice(c *Client, des, nw []ClusterNodePoolsConfigShieldedInstanceConfig) []ClusterNodePoolsConfigShieldedInstanceConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConfigShieldedInstanceConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConfigShieldedInstanceConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConfigLinuxNodeConfig(des, initial *ClusterNodePoolsConfigLinuxNodeConfig, opts ...dcl.ApplyOption) *ClusterNodePoolsConfigLinuxNodeConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Sysctls) {
		des.Sysctls = initial.Sysctls
	}

	return des
}

func canonicalizeNewClusterNodePoolsConfigLinuxNodeConfig(c *Client, des, nw *ClusterNodePoolsConfigLinuxNodeConfig) *ClusterNodePoolsConfigLinuxNodeConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConfigLinuxNodeConfigSet(c *Client, des, nw []ClusterNodePoolsConfigLinuxNodeConfig) []ClusterNodePoolsConfigLinuxNodeConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConfigLinuxNodeConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConfigLinuxNodeConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConfigLinuxNodeConfigSlice(c *Client, des, nw []ClusterNodePoolsConfigLinuxNodeConfig) []ClusterNodePoolsConfigLinuxNodeConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConfigLinuxNodeConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConfigLinuxNodeConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConfigKubeletConfig(des, initial *ClusterNodePoolsConfigKubeletConfig, opts ...dcl.ApplyOption) *ClusterNodePoolsConfigKubeletConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CpuManagerPolicy, initial.CpuManagerPolicy) || dcl.IsZeroValue(des.CpuManagerPolicy) {
		des.CpuManagerPolicy = initial.CpuManagerPolicy
	}
	if dcl.BoolCanonicalize(des.CpuCfsQuota, initial.CpuCfsQuota) || dcl.IsZeroValue(des.CpuCfsQuota) {
		des.CpuCfsQuota = initial.CpuCfsQuota
	}
	if dcl.StringCanonicalize(des.CpuCfsQuotaPeriod, initial.CpuCfsQuotaPeriod) || dcl.IsZeroValue(des.CpuCfsQuotaPeriod) {
		des.CpuCfsQuotaPeriod = initial.CpuCfsQuotaPeriod
	}

	return des
}

func canonicalizeNewClusterNodePoolsConfigKubeletConfig(c *Client, des, nw *ClusterNodePoolsConfigKubeletConfig) *ClusterNodePoolsConfigKubeletConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CpuManagerPolicy, nw.CpuManagerPolicy) {
		nw.CpuManagerPolicy = des.CpuManagerPolicy
	}
	if dcl.BoolCanonicalize(des.CpuCfsQuota, nw.CpuCfsQuota) {
		nw.CpuCfsQuota = des.CpuCfsQuota
	}
	if dcl.StringCanonicalize(des.CpuCfsQuotaPeriod, nw.CpuCfsQuotaPeriod) {
		nw.CpuCfsQuotaPeriod = des.CpuCfsQuotaPeriod
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConfigKubeletConfigSet(c *Client, des, nw []ClusterNodePoolsConfigKubeletConfig) []ClusterNodePoolsConfigKubeletConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConfigKubeletConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConfigKubeletConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConfigKubeletConfigSlice(c *Client, des, nw []ClusterNodePoolsConfigKubeletConfig) []ClusterNodePoolsConfigKubeletConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConfigKubeletConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConfigKubeletConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsAutoscaling(des, initial *ClusterNodePoolsAutoscaling, opts ...dcl.ApplyOption) *ClusterNodePoolsAutoscaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.IsZeroValue(des.MinNodeCount) {
		des.MinNodeCount = initial.MinNodeCount
	}
	if dcl.IsZeroValue(des.MaxNodeCount) {
		des.MaxNodeCount = initial.MaxNodeCount
	}
	if dcl.BoolCanonicalize(des.Autoprovisioned, initial.Autoprovisioned) || dcl.IsZeroValue(des.Autoprovisioned) {
		des.Autoprovisioned = initial.Autoprovisioned
	}

	return des
}

func canonicalizeNewClusterNodePoolsAutoscaling(c *Client, des, nw *ClusterNodePoolsAutoscaling) *ClusterNodePoolsAutoscaling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.BoolCanonicalize(des.Autoprovisioned, nw.Autoprovisioned) {
		nw.Autoprovisioned = des.Autoprovisioned
	}

	return nw
}

func canonicalizeNewClusterNodePoolsAutoscalingSet(c *Client, des, nw []ClusterNodePoolsAutoscaling) []ClusterNodePoolsAutoscaling {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsAutoscaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsAutoscaling(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsAutoscalingSlice(c *Client, des, nw []ClusterNodePoolsAutoscaling) []ClusterNodePoolsAutoscaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsAutoscaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsAutoscaling(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsManagement(des, initial *ClusterNodePoolsManagement, opts ...dcl.ApplyOption) *ClusterNodePoolsManagement {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.AutoUpgrade, initial.AutoUpgrade) || dcl.IsZeroValue(des.AutoUpgrade) {
		des.AutoUpgrade = initial.AutoUpgrade
	}
	if dcl.BoolCanonicalize(des.AutoRepair, initial.AutoRepair) || dcl.IsZeroValue(des.AutoRepair) {
		des.AutoRepair = initial.AutoRepair
	}
	des.UpgradeOptions = canonicalizeClusterNodePoolsManagementUpgradeOptions(des.UpgradeOptions, initial.UpgradeOptions, opts...)

	return des
}

func canonicalizeNewClusterNodePoolsManagement(c *Client, des, nw *ClusterNodePoolsManagement) *ClusterNodePoolsManagement {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.AutoUpgrade, nw.AutoUpgrade) {
		nw.AutoUpgrade = des.AutoUpgrade
	}
	if dcl.BoolCanonicalize(des.AutoRepair, nw.AutoRepair) {
		nw.AutoRepair = des.AutoRepair
	}
	nw.UpgradeOptions = canonicalizeNewClusterNodePoolsManagementUpgradeOptions(c, des.UpgradeOptions, nw.UpgradeOptions)

	return nw
}

func canonicalizeNewClusterNodePoolsManagementSet(c *Client, des, nw []ClusterNodePoolsManagement) []ClusterNodePoolsManagement {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsManagement
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsManagement(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsManagementSlice(c *Client, des, nw []ClusterNodePoolsManagement) []ClusterNodePoolsManagement {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsManagement
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsManagement(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsManagementUpgradeOptions(des, initial *ClusterNodePoolsManagementUpgradeOptions, opts ...dcl.ApplyOption) *ClusterNodePoolsManagementUpgradeOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.AutoUpgradeStartTime, initial.AutoUpgradeStartTime) || dcl.IsZeroValue(des.AutoUpgradeStartTime) {
		des.AutoUpgradeStartTime = initial.AutoUpgradeStartTime
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}

	return des
}

func canonicalizeNewClusterNodePoolsManagementUpgradeOptions(c *Client, des, nw *ClusterNodePoolsManagementUpgradeOptions) *ClusterNodePoolsManagementUpgradeOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AutoUpgradeStartTime, nw.AutoUpgradeStartTime) {
		nw.AutoUpgradeStartTime = des.AutoUpgradeStartTime
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewClusterNodePoolsManagementUpgradeOptionsSet(c *Client, des, nw []ClusterNodePoolsManagementUpgradeOptions) []ClusterNodePoolsManagementUpgradeOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsManagementUpgradeOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsManagementUpgradeOptions(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsManagementUpgradeOptionsSlice(c *Client, des, nw []ClusterNodePoolsManagementUpgradeOptions) []ClusterNodePoolsManagementUpgradeOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsManagementUpgradeOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsManagementUpgradeOptions(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsMaxPodsConstraint(des, initial *ClusterNodePoolsMaxPodsConstraint, opts ...dcl.ApplyOption) *ClusterNodePoolsMaxPodsConstraint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MaxPodsPerNode) {
		des.MaxPodsPerNode = initial.MaxPodsPerNode
	}

	return des
}

func canonicalizeNewClusterNodePoolsMaxPodsConstraint(c *Client, des, nw *ClusterNodePoolsMaxPodsConstraint) *ClusterNodePoolsMaxPodsConstraint {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterNodePoolsMaxPodsConstraintSet(c *Client, des, nw []ClusterNodePoolsMaxPodsConstraint) []ClusterNodePoolsMaxPodsConstraint {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsMaxPodsConstraint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsMaxPodsConstraint(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsMaxPodsConstraintSlice(c *Client, des, nw []ClusterNodePoolsMaxPodsConstraint) []ClusterNodePoolsMaxPodsConstraint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsMaxPodsConstraint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsMaxPodsConstraint(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsConditions(des, initial *ClusterNodePoolsConditions, opts ...dcl.ApplyOption) *ClusterNodePoolsConditions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Code) {
		des.Code = initial.Code
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		des.Message = initial.Message
	}
	if dcl.IsZeroValue(des.CanonicalCode) {
		des.CanonicalCode = initial.CanonicalCode
	}

	return des
}

func canonicalizeNewClusterNodePoolsConditions(c *Client, des, nw *ClusterNodePoolsConditions) *ClusterNodePoolsConditions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}

	return nw
}

func canonicalizeNewClusterNodePoolsConditionsSet(c *Client, des, nw []ClusterNodePoolsConditions) []ClusterNodePoolsConditions {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsConditions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsConditions(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsConditionsSlice(c *Client, des, nw []ClusterNodePoolsConditions) []ClusterNodePoolsConditions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsConditions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsConditions(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodePoolsUpgradeSettings(des, initial *ClusterNodePoolsUpgradeSettings, opts ...dcl.ApplyOption) *ClusterNodePoolsUpgradeSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MaxSurge) {
		des.MaxSurge = initial.MaxSurge
	}
	if dcl.IsZeroValue(des.MaxUnavailable) {
		des.MaxUnavailable = initial.MaxUnavailable
	}

	return des
}

func canonicalizeNewClusterNodePoolsUpgradeSettings(c *Client, des, nw *ClusterNodePoolsUpgradeSettings) *ClusterNodePoolsUpgradeSettings {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterNodePoolsUpgradeSettingsSet(c *Client, des, nw []ClusterNodePoolsUpgradeSettings) []ClusterNodePoolsUpgradeSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodePoolsUpgradeSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodePoolsUpgradeSettings(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodePoolsUpgradeSettingsSlice(c *Client, des, nw []ClusterNodePoolsUpgradeSettings) []ClusterNodePoolsUpgradeSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodePoolsUpgradeSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodePoolsUpgradeSettings(c, &d, &n))
	}

	return items
}

func canonicalizeClusterLegacyAbac(des, initial *ClusterLegacyAbac, opts ...dcl.ApplyOption) *ClusterLegacyAbac {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterLegacyAbac(c *Client, des, nw *ClusterLegacyAbac) *ClusterLegacyAbac {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterLegacyAbacSet(c *Client, des, nw []ClusterLegacyAbac) []ClusterLegacyAbac {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterLegacyAbac
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterLegacyAbac(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterLegacyAbacSlice(c *Client, des, nw []ClusterLegacyAbac) []ClusterLegacyAbac {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterLegacyAbac
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterLegacyAbac(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNetworkPolicy(des, initial *ClusterNetworkPolicy, opts ...dcl.ApplyOption) *ClusterNetworkPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Provider) {
		des.Provider = initial.Provider
	}
	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterNetworkPolicy(c *Client, des, nw *ClusterNetworkPolicy) *ClusterNetworkPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterNetworkPolicySet(c *Client, des, nw []ClusterNetworkPolicy) []ClusterNetworkPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNetworkPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNetworkPolicy(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNetworkPolicySlice(c *Client, des, nw []ClusterNetworkPolicy) []ClusterNetworkPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNetworkPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNetworkPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeClusterIPAllocationPolicy(des, initial *ClusterIPAllocationPolicy, opts ...dcl.ApplyOption) *ClusterIPAllocationPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.UseIPAliases, initial.UseIPAliases) || dcl.IsZeroValue(des.UseIPAliases) {
		des.UseIPAliases = initial.UseIPAliases
	}
	if dcl.BoolCanonicalize(des.CreateSubnetwork, initial.CreateSubnetwork) || dcl.IsZeroValue(des.CreateSubnetwork) {
		des.CreateSubnetwork = initial.CreateSubnetwork
	}
	if dcl.StringCanonicalize(des.SubnetworkName, initial.SubnetworkName) || dcl.IsZeroValue(des.SubnetworkName) {
		des.SubnetworkName = initial.SubnetworkName
	}
	if dcl.StringCanonicalize(des.ClusterSecondaryRangeName, initial.ClusterSecondaryRangeName) || dcl.IsZeroValue(des.ClusterSecondaryRangeName) {
		des.ClusterSecondaryRangeName = initial.ClusterSecondaryRangeName
	}
	if dcl.StringCanonicalize(des.ServicesSecondaryRangeName, initial.ServicesSecondaryRangeName) || dcl.IsZeroValue(des.ServicesSecondaryRangeName) {
		des.ServicesSecondaryRangeName = initial.ServicesSecondaryRangeName
	}
	if dcl.StringCanonicalize(des.ClusterIPv4CidrBlock, initial.ClusterIPv4CidrBlock) || dcl.IsZeroValue(des.ClusterIPv4CidrBlock) {
		des.ClusterIPv4CidrBlock = initial.ClusterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.NodeIPv4CidrBlock, initial.NodeIPv4CidrBlock) || dcl.IsZeroValue(des.NodeIPv4CidrBlock) {
		des.NodeIPv4CidrBlock = initial.NodeIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesIPv4CidrBlock, initial.ServicesIPv4CidrBlock) || dcl.IsZeroValue(des.ServicesIPv4CidrBlock) {
		des.ServicesIPv4CidrBlock = initial.ServicesIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.TPUIPv4CidrBlock, initial.TPUIPv4CidrBlock) || dcl.IsZeroValue(des.TPUIPv4CidrBlock) {
		des.TPUIPv4CidrBlock = initial.TPUIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ClusterIPv4Cidr, initial.ClusterIPv4Cidr) || dcl.IsZeroValue(des.ClusterIPv4Cidr) {
		des.ClusterIPv4Cidr = initial.ClusterIPv4Cidr
	}
	if dcl.StringCanonicalize(des.NodeIPv4Cidr, initial.NodeIPv4Cidr) || dcl.IsZeroValue(des.NodeIPv4Cidr) {
		des.NodeIPv4Cidr = initial.NodeIPv4Cidr
	}
	if dcl.StringCanonicalize(des.ServicesIPv4Cidr, initial.ServicesIPv4Cidr) || dcl.IsZeroValue(des.ServicesIPv4Cidr) {
		des.ServicesIPv4Cidr = initial.ServicesIPv4Cidr
	}
	if dcl.BoolCanonicalize(des.UseRoutes, initial.UseRoutes) || dcl.IsZeroValue(des.UseRoutes) {
		des.UseRoutes = initial.UseRoutes
	}

	return des
}

func canonicalizeNewClusterIPAllocationPolicy(c *Client, des, nw *ClusterIPAllocationPolicy) *ClusterIPAllocationPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.UseIPAliases, nw.UseIPAliases) {
		nw.UseIPAliases = des.UseIPAliases
	}
	if dcl.BoolCanonicalize(des.CreateSubnetwork, nw.CreateSubnetwork) {
		nw.CreateSubnetwork = des.CreateSubnetwork
	}
	if dcl.StringCanonicalize(des.SubnetworkName, nw.SubnetworkName) {
		nw.SubnetworkName = des.SubnetworkName
	}
	if dcl.StringCanonicalize(des.ClusterSecondaryRangeName, nw.ClusterSecondaryRangeName) {
		nw.ClusterSecondaryRangeName = des.ClusterSecondaryRangeName
	}
	if dcl.StringCanonicalize(des.ServicesSecondaryRangeName, nw.ServicesSecondaryRangeName) {
		nw.ServicesSecondaryRangeName = des.ServicesSecondaryRangeName
	}
	if dcl.StringCanonicalize(des.ClusterIPv4CidrBlock, nw.ClusterIPv4CidrBlock) {
		nw.ClusterIPv4CidrBlock = des.ClusterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.NodeIPv4CidrBlock, nw.NodeIPv4CidrBlock) {
		nw.NodeIPv4CidrBlock = des.NodeIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesIPv4CidrBlock, nw.ServicesIPv4CidrBlock) {
		nw.ServicesIPv4CidrBlock = des.ServicesIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.TPUIPv4CidrBlock, nw.TPUIPv4CidrBlock) {
		nw.TPUIPv4CidrBlock = des.TPUIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ClusterIPv4Cidr, nw.ClusterIPv4Cidr) {
		nw.ClusterIPv4Cidr = des.ClusterIPv4Cidr
	}
	if dcl.StringCanonicalize(des.NodeIPv4Cidr, nw.NodeIPv4Cidr) {
		nw.NodeIPv4Cidr = des.NodeIPv4Cidr
	}
	if dcl.StringCanonicalize(des.ServicesIPv4Cidr, nw.ServicesIPv4Cidr) {
		nw.ServicesIPv4Cidr = des.ServicesIPv4Cidr
	}
	if dcl.BoolCanonicalize(des.UseRoutes, nw.UseRoutes) {
		nw.UseRoutes = des.UseRoutes
	}

	return nw
}

func canonicalizeNewClusterIPAllocationPolicySet(c *Client, des, nw []ClusterIPAllocationPolicy) []ClusterIPAllocationPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterIPAllocationPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterIPAllocationPolicy(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterIPAllocationPolicySlice(c *Client, des, nw []ClusterIPAllocationPolicy) []ClusterIPAllocationPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterIPAllocationPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterIPAllocationPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeClusterMasterAuthorizedNetworksConfig(des, initial *ClusterMasterAuthorizedNetworksConfig, opts ...dcl.ApplyOption) *ClusterMasterAuthorizedNetworksConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.IsZeroValue(des.CidrBlocks) {
		des.CidrBlocks = initial.CidrBlocks
	}

	return des
}

func canonicalizeNewClusterMasterAuthorizedNetworksConfig(c *Client, des, nw *ClusterMasterAuthorizedNetworksConfig) *ClusterMasterAuthorizedNetworksConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	nw.CidrBlocks = canonicalizeNewClusterMasterAuthorizedNetworksConfigCidrBlocksSlice(c, des.CidrBlocks, nw.CidrBlocks)

	return nw
}

func canonicalizeNewClusterMasterAuthorizedNetworksConfigSet(c *Client, des, nw []ClusterMasterAuthorizedNetworksConfig) []ClusterMasterAuthorizedNetworksConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMasterAuthorizedNetworksConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterMasterAuthorizedNetworksConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterMasterAuthorizedNetworksConfigSlice(c *Client, des, nw []ClusterMasterAuthorizedNetworksConfig) []ClusterMasterAuthorizedNetworksConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterMasterAuthorizedNetworksConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMasterAuthorizedNetworksConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterMasterAuthorizedNetworksConfigCidrBlocks(des, initial *ClusterMasterAuthorizedNetworksConfigCidrBlocks, opts ...dcl.ApplyOption) *ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.DisplayName, initial.DisplayName) || dcl.IsZeroValue(des.DisplayName) {
		des.DisplayName = initial.DisplayName
	}
	if dcl.StringCanonicalize(des.CidrBlock, initial.CidrBlock) || dcl.IsZeroValue(des.CidrBlock) {
		des.CidrBlock = initial.CidrBlock
	}

	return des
}

func canonicalizeNewClusterMasterAuthorizedNetworksConfigCidrBlocks(c *Client, des, nw *ClusterMasterAuthorizedNetworksConfigCidrBlocks) *ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.DisplayName, nw.DisplayName) {
		nw.DisplayName = des.DisplayName
	}
	if dcl.StringCanonicalize(des.CidrBlock, nw.CidrBlock) {
		nw.CidrBlock = des.CidrBlock
	}

	return nw
}

func canonicalizeNewClusterMasterAuthorizedNetworksConfigCidrBlocksSet(c *Client, des, nw []ClusterMasterAuthorizedNetworksConfigCidrBlocks) []ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMasterAuthorizedNetworksConfigCidrBlocks
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterMasterAuthorizedNetworksConfigCidrBlocks(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterMasterAuthorizedNetworksConfigCidrBlocksSlice(c *Client, des, nw []ClusterMasterAuthorizedNetworksConfigCidrBlocks) []ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterMasterAuthorizedNetworksConfigCidrBlocks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMasterAuthorizedNetworksConfigCidrBlocks(c, &d, &n))
	}

	return items
}

func canonicalizeClusterBinaryAuthorization(des, initial *ClusterBinaryAuthorization, opts ...dcl.ApplyOption) *ClusterBinaryAuthorization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterBinaryAuthorization(c *Client, des, nw *ClusterBinaryAuthorization) *ClusterBinaryAuthorization {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterBinaryAuthorizationSet(c *Client, des, nw []ClusterBinaryAuthorization) []ClusterBinaryAuthorization {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterBinaryAuthorization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterBinaryAuthorization(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterBinaryAuthorizationSlice(c *Client, des, nw []ClusterBinaryAuthorization) []ClusterBinaryAuthorization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterBinaryAuthorization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterBinaryAuthorization(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAutoscaling(des, initial *ClusterAutoscaling, opts ...dcl.ApplyOption) *ClusterAutoscaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.EnableNodeAutoprovisioning, initial.EnableNodeAutoprovisioning) || dcl.IsZeroValue(des.EnableNodeAutoprovisioning) {
		des.EnableNodeAutoprovisioning = initial.EnableNodeAutoprovisioning
	}
	if dcl.IsZeroValue(des.ResourceLimits) {
		des.ResourceLimits = initial.ResourceLimits
	}
	des.AutoprovisioningNodePoolDefaults = canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaults(des.AutoprovisioningNodePoolDefaults, initial.AutoprovisioningNodePoolDefaults, opts...)
	if dcl.IsZeroValue(des.AutoprovisioningLocations) {
		des.AutoprovisioningLocations = initial.AutoprovisioningLocations
	}

	return des
}

func canonicalizeNewClusterAutoscaling(c *Client, des, nw *ClusterAutoscaling) *ClusterAutoscaling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnableNodeAutoprovisioning, nw.EnableNodeAutoprovisioning) {
		nw.EnableNodeAutoprovisioning = des.EnableNodeAutoprovisioning
	}
	nw.ResourceLimits = canonicalizeNewClusterAutoscalingResourceLimitsSlice(c, des.ResourceLimits, nw.ResourceLimits)
	nw.AutoprovisioningNodePoolDefaults = canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaults(c, des.AutoprovisioningNodePoolDefaults, nw.AutoprovisioningNodePoolDefaults)

	return nw
}

func canonicalizeNewClusterAutoscalingSet(c *Client, des, nw []ClusterAutoscaling) []ClusterAutoscaling {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAutoscaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAutoscaling(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAutoscalingSlice(c *Client, des, nw []ClusterAutoscaling) []ClusterAutoscaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAutoscaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAutoscaling(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAutoscalingResourceLimits(des, initial *ClusterAutoscalingResourceLimits, opts ...dcl.ApplyOption) *ClusterAutoscalingResourceLimits {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ResourceType, initial.ResourceType) || dcl.IsZeroValue(des.ResourceType) {
		des.ResourceType = initial.ResourceType
	}
	if dcl.IsZeroValue(des.Minimum) {
		des.Minimum = initial.Minimum
	}
	if dcl.IsZeroValue(des.Maximum) {
		des.Maximum = initial.Maximum
	}

	return des
}

func canonicalizeNewClusterAutoscalingResourceLimits(c *Client, des, nw *ClusterAutoscalingResourceLimits) *ClusterAutoscalingResourceLimits {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ResourceType, nw.ResourceType) {
		nw.ResourceType = des.ResourceType
	}

	return nw
}

func canonicalizeNewClusterAutoscalingResourceLimitsSet(c *Client, des, nw []ClusterAutoscalingResourceLimits) []ClusterAutoscalingResourceLimits {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAutoscalingResourceLimits
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAutoscalingResourceLimits(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAutoscalingResourceLimitsSlice(c *Client, des, nw []ClusterAutoscalingResourceLimits) []ClusterAutoscalingResourceLimits {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAutoscalingResourceLimits
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAutoscalingResourceLimits(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaults(des, initial *ClusterAutoscalingAutoprovisioningNodePoolDefaults, opts ...dcl.ApplyOption) *ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.OAuthScopes) {
		des.OAuthScopes = initial.OAuthScopes
	}
	if dcl.StringCanonicalize(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		des.ServiceAccount = initial.ServiceAccount
	}
	des.UpgradeSettings = canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(des.UpgradeSettings, initial.UpgradeSettings, opts...)
	des.Management = canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(des.Management, initial.Management, opts...)
	if dcl.StringCanonicalize(des.MinCpuPlatform, initial.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
		des.MinCpuPlatform = initial.MinCpuPlatform
	}
	if dcl.IsZeroValue(des.DiskSizeGb) {
		des.DiskSizeGb = initial.DiskSizeGb
	}
	if dcl.StringCanonicalize(des.DiskType, initial.DiskType) || dcl.IsZeroValue(des.DiskType) {
		des.DiskType = initial.DiskType
	}
	des.ShieldedInstanceConfig = canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(des.ShieldedInstanceConfig, initial.ShieldedInstanceConfig, opts...)
	if dcl.StringCanonicalize(des.BootDiskKmsKey, initial.BootDiskKmsKey) || dcl.IsZeroValue(des.BootDiskKmsKey) {
		des.BootDiskKmsKey = initial.BootDiskKmsKey
	}

	return des
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaults(c *Client, des, nw *ClusterAutoscalingAutoprovisioningNodePoolDefaults) *ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	nw.UpgradeSettings = canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, des.UpgradeSettings, nw.UpgradeSettings)
	nw.Management = canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, des.Management, nw.Management)
	if dcl.StringCanonicalize(des.MinCpuPlatform, nw.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}
	if dcl.StringCanonicalize(des.DiskType, nw.DiskType) {
		nw.DiskType = des.DiskType
	}
	nw.ShieldedInstanceConfig = canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, des.ShieldedInstanceConfig, nw.ShieldedInstanceConfig)
	if dcl.StringCanonicalize(des.BootDiskKmsKey, nw.BootDiskKmsKey) {
		nw.BootDiskKmsKey = des.BootDiskKmsKey
	}

	return nw
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsSet(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaults) []ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAutoscalingAutoprovisioningNodePoolDefaults
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAutoscalingAutoprovisioningNodePoolDefaults(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsSlice(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaults) []ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAutoscalingAutoprovisioningNodePoolDefaults
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaults(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(des, initial *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings, opts ...dcl.ApplyOption) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MaxSurge) {
		des.MaxSurge = initial.MaxSurge
	}
	if dcl.IsZeroValue(des.MaxUnavailable) {
		des.MaxUnavailable = initial.MaxUnavailable
	}

	return des
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c *Client, des, nw *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsSet(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsSlice(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(des, initial *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement, opts ...dcl.ApplyOption) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.AutoUpgrade, initial.AutoUpgrade) || dcl.IsZeroValue(des.AutoUpgrade) {
		des.AutoUpgrade = initial.AutoUpgrade
	}
	if dcl.BoolCanonicalize(des.AutoRepair, initial.AutoRepair) || dcl.IsZeroValue(des.AutoRepair) {
		des.AutoRepair = initial.AutoRepair
	}
	des.UpgradeOptions = canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(des.UpgradeOptions, initial.UpgradeOptions, opts...)

	return des
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c *Client, des, nw *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.AutoUpgrade, nw.AutoUpgrade) {
		nw.AutoUpgrade = des.AutoUpgrade
	}
	if dcl.BoolCanonicalize(des.AutoRepair, nw.AutoRepair) {
		nw.AutoRepair = des.AutoRepair
	}
	nw.UpgradeOptions = canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, des.UpgradeOptions, nw.UpgradeOptions)

	return nw
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementSet(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementSlice(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(des, initial *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions, opts ...dcl.ApplyOption) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.AutoUpgradeStartTime, initial.AutoUpgradeStartTime) || dcl.IsZeroValue(des.AutoUpgradeStartTime) {
		des.AutoUpgradeStartTime = initial.AutoUpgradeStartTime
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}

	return des
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c *Client, des, nw *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AutoUpgradeStartTime, nw.AutoUpgradeStartTime) {
		nw.AutoUpgradeStartTime = des.AutoUpgradeStartTime
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsSet(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsSlice(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(des, initial *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig, opts ...dcl.ApplyOption) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.EnableSecureBoot, initial.EnableSecureBoot) || dcl.IsZeroValue(des.EnableSecureBoot) {
		des.EnableSecureBoot = initial.EnableSecureBoot
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, initial.EnableIntegrityMonitoring) || dcl.IsZeroValue(des.EnableIntegrityMonitoring) {
		des.EnableIntegrityMonitoring = initial.EnableIntegrityMonitoring
	}

	return des
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c *Client, des, nw *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnableSecureBoot, nw.EnableSecureBoot) {
		nw.EnableSecureBoot = des.EnableSecureBoot
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, nw.EnableIntegrityMonitoring) {
		nw.EnableIntegrityMonitoring = des.EnableIntegrityMonitoring
	}

	return nw
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigSet(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigSlice(c *Client, des, nw []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNetworkConfig(des, initial *ClusterNetworkConfig, opts ...dcl.ApplyOption) *ClusterNetworkConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Network, initial.Network) || dcl.IsZeroValue(des.Network) {
		des.Network = initial.Network
	}
	if dcl.StringCanonicalize(des.Subnetwork, initial.Subnetwork) || dcl.IsZeroValue(des.Subnetwork) {
		des.Subnetwork = initial.Subnetwork
	}
	if dcl.BoolCanonicalize(des.EnableIntraNodeVisibility, initial.EnableIntraNodeVisibility) || dcl.IsZeroValue(des.EnableIntraNodeVisibility) {
		des.EnableIntraNodeVisibility = initial.EnableIntraNodeVisibility
	}
	des.DefaultSnatStatus = canonicalizeClusterNetworkConfigDefaultSnatStatus(des.DefaultSnatStatus, initial.DefaultSnatStatus, opts...)
	if dcl.IsZeroValue(des.PrivateIPv6GoogleAccess) {
		des.PrivateIPv6GoogleAccess = initial.PrivateIPv6GoogleAccess
	}

	return des
}

func canonicalizeNewClusterNetworkConfig(c *Client, des, nw *ClusterNetworkConfig) *ClusterNetworkConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Network, nw.Network) {
		nw.Network = des.Network
	}
	if dcl.StringCanonicalize(des.Subnetwork, nw.Subnetwork) {
		nw.Subnetwork = des.Subnetwork
	}
	if dcl.BoolCanonicalize(des.EnableIntraNodeVisibility, nw.EnableIntraNodeVisibility) {
		nw.EnableIntraNodeVisibility = des.EnableIntraNodeVisibility
	}
	nw.DefaultSnatStatus = canonicalizeNewClusterNetworkConfigDefaultSnatStatus(c, des.DefaultSnatStatus, nw.DefaultSnatStatus)

	return nw
}

func canonicalizeNewClusterNetworkConfigSet(c *Client, des, nw []ClusterNetworkConfig) []ClusterNetworkConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNetworkConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNetworkConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNetworkConfigSlice(c *Client, des, nw []ClusterNetworkConfig) []ClusterNetworkConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNetworkConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNetworkConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNetworkConfigDefaultSnatStatus(des, initial *ClusterNetworkConfigDefaultSnatStatus, opts ...dcl.ApplyOption) *ClusterNetworkConfigDefaultSnatStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterNetworkConfigDefaultSnatStatus(c *Client, des, nw *ClusterNetworkConfigDefaultSnatStatus) *ClusterNetworkConfigDefaultSnatStatus {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
	}

	return nw
}

func canonicalizeNewClusterNetworkConfigDefaultSnatStatusSet(c *Client, des, nw []ClusterNetworkConfigDefaultSnatStatus) []ClusterNetworkConfigDefaultSnatStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNetworkConfigDefaultSnatStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNetworkConfigDefaultSnatStatus(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNetworkConfigDefaultSnatStatusSlice(c *Client, des, nw []ClusterNetworkConfigDefaultSnatStatus) []ClusterNetworkConfigDefaultSnatStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNetworkConfigDefaultSnatStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNetworkConfigDefaultSnatStatus(c, &d, &n))
	}

	return items
}

func canonicalizeClusterMaintenancePolicy(des, initial *ClusterMaintenancePolicy, opts ...dcl.ApplyOption) *ClusterMaintenancePolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Window = canonicalizeClusterMaintenancePolicyWindow(des.Window, initial.Window, opts...)
	if dcl.StringCanonicalize(des.ResourceVersion, initial.ResourceVersion) || dcl.IsZeroValue(des.ResourceVersion) {
		des.ResourceVersion = initial.ResourceVersion
	}

	return des
}

func canonicalizeNewClusterMaintenancePolicy(c *Client, des, nw *ClusterMaintenancePolicy) *ClusterMaintenancePolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.Window = canonicalizeNewClusterMaintenancePolicyWindow(c, des.Window, nw.Window)
	if dcl.StringCanonicalize(des.ResourceVersion, nw.ResourceVersion) {
		nw.ResourceVersion = des.ResourceVersion
	}

	return nw
}

func canonicalizeNewClusterMaintenancePolicySet(c *Client, des, nw []ClusterMaintenancePolicy) []ClusterMaintenancePolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMaintenancePolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterMaintenancePolicy(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterMaintenancePolicySlice(c *Client, des, nw []ClusterMaintenancePolicy) []ClusterMaintenancePolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterMaintenancePolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMaintenancePolicy(c, &d, &n))
	}

	return items
}

func canonicalizeClusterMaintenancePolicyWindow(des, initial *ClusterMaintenancePolicyWindow, opts ...dcl.ApplyOption) *ClusterMaintenancePolicyWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.DailyMaintenanceWindow = canonicalizeClusterMaintenancePolicyWindowDailyMaintenanceWindow(des.DailyMaintenanceWindow, initial.DailyMaintenanceWindow, opts...)
	des.RecurringWindow = canonicalizeClusterMaintenancePolicyWindowRecurringWindow(des.RecurringWindow, initial.RecurringWindow, opts...)
	if dcl.IsZeroValue(des.MaintenanceExclusions) {
		des.MaintenanceExclusions = initial.MaintenanceExclusions
	}

	return des
}

func canonicalizeNewClusterMaintenancePolicyWindow(c *Client, des, nw *ClusterMaintenancePolicyWindow) *ClusterMaintenancePolicyWindow {
	if des == nil || nw == nil {
		return nw
	}

	nw.DailyMaintenanceWindow = canonicalizeNewClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, des.DailyMaintenanceWindow, nw.DailyMaintenanceWindow)
	nw.RecurringWindow = canonicalizeNewClusterMaintenancePolicyWindowRecurringWindow(c, des.RecurringWindow, nw.RecurringWindow)

	return nw
}

func canonicalizeNewClusterMaintenancePolicyWindowSet(c *Client, des, nw []ClusterMaintenancePolicyWindow) []ClusterMaintenancePolicyWindow {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMaintenancePolicyWindow
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterMaintenancePolicyWindow(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterMaintenancePolicyWindowSlice(c *Client, des, nw []ClusterMaintenancePolicyWindow) []ClusterMaintenancePolicyWindow {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterMaintenancePolicyWindow
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMaintenancePolicyWindow(c, &d, &n))
	}

	return items
}

func canonicalizeClusterMaintenancePolicyWindowDailyMaintenanceWindow(des, initial *ClusterMaintenancePolicyWindowDailyMaintenanceWindow, opts ...dcl.ApplyOption) *ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.StartTime) {
		des.StartTime = initial.StartTime
	}
	if dcl.StringCanonicalize(des.Duration, initial.Duration) || dcl.IsZeroValue(des.Duration) {
		des.Duration = initial.Duration
	}

	return des
}

func canonicalizeNewClusterMaintenancePolicyWindowDailyMaintenanceWindow(c *Client, des, nw *ClusterMaintenancePolicyWindowDailyMaintenanceWindow) *ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Duration, nw.Duration) {
		nw.Duration = des.Duration
	}

	return nw
}

func canonicalizeNewClusterMaintenancePolicyWindowDailyMaintenanceWindowSet(c *Client, des, nw []ClusterMaintenancePolicyWindowDailyMaintenanceWindow) []ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMaintenancePolicyWindowDailyMaintenanceWindow
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterMaintenancePolicyWindowDailyMaintenanceWindowSlice(c *Client, des, nw []ClusterMaintenancePolicyWindowDailyMaintenanceWindow) []ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterMaintenancePolicyWindowDailyMaintenanceWindow
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, &d, &n))
	}

	return items
}

func canonicalizeClusterMaintenancePolicyWindowRecurringWindow(des, initial *ClusterMaintenancePolicyWindowRecurringWindow, opts ...dcl.ApplyOption) *ClusterMaintenancePolicyWindowRecurringWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Window = canonicalizeClusterMaintenancePolicyWindowRecurringWindowWindow(des.Window, initial.Window, opts...)
	if dcl.StringCanonicalize(des.Recurrence, initial.Recurrence) || dcl.IsZeroValue(des.Recurrence) {
		des.Recurrence = initial.Recurrence
	}

	return des
}

func canonicalizeNewClusterMaintenancePolicyWindowRecurringWindow(c *Client, des, nw *ClusterMaintenancePolicyWindowRecurringWindow) *ClusterMaintenancePolicyWindowRecurringWindow {
	if des == nil || nw == nil {
		return nw
	}

	nw.Window = canonicalizeNewClusterMaintenancePolicyWindowRecurringWindowWindow(c, des.Window, nw.Window)
	if dcl.StringCanonicalize(des.Recurrence, nw.Recurrence) {
		nw.Recurrence = des.Recurrence
	}

	return nw
}

func canonicalizeNewClusterMaintenancePolicyWindowRecurringWindowSet(c *Client, des, nw []ClusterMaintenancePolicyWindowRecurringWindow) []ClusterMaintenancePolicyWindowRecurringWindow {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMaintenancePolicyWindowRecurringWindow
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterMaintenancePolicyWindowRecurringWindow(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterMaintenancePolicyWindowRecurringWindowSlice(c *Client, des, nw []ClusterMaintenancePolicyWindowRecurringWindow) []ClusterMaintenancePolicyWindowRecurringWindow {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterMaintenancePolicyWindowRecurringWindow
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMaintenancePolicyWindowRecurringWindow(c, &d, &n))
	}

	return items
}

func canonicalizeClusterMaintenancePolicyWindowRecurringWindowWindow(des, initial *ClusterMaintenancePolicyWindowRecurringWindowWindow, opts ...dcl.ApplyOption) *ClusterMaintenancePolicyWindowRecurringWindowWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.StartTime) {
		des.StartTime = initial.StartTime
	}
	if dcl.IsZeroValue(des.EndTime) {
		des.EndTime = initial.EndTime
	}

	return des
}

func canonicalizeNewClusterMaintenancePolicyWindowRecurringWindowWindow(c *Client, des, nw *ClusterMaintenancePolicyWindowRecurringWindowWindow) *ClusterMaintenancePolicyWindowRecurringWindowWindow {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterMaintenancePolicyWindowRecurringWindowWindowSet(c *Client, des, nw []ClusterMaintenancePolicyWindowRecurringWindowWindow) []ClusterMaintenancePolicyWindowRecurringWindowWindow {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterMaintenancePolicyWindowRecurringWindowWindow
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterMaintenancePolicyWindowRecurringWindowWindow(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterMaintenancePolicyWindowRecurringWindowWindowSlice(c *Client, des, nw []ClusterMaintenancePolicyWindowRecurringWindowWindow) []ClusterMaintenancePolicyWindowRecurringWindowWindow {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterMaintenancePolicyWindowRecurringWindowWindow
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterMaintenancePolicyWindowRecurringWindowWindow(c, &d, &n))
	}

	return items
}

func canonicalizeClusterDefaultMaxPodsConstraint(des, initial *ClusterDefaultMaxPodsConstraint, opts ...dcl.ApplyOption) *ClusterDefaultMaxPodsConstraint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MaxPodsPerNode, initial.MaxPodsPerNode) || dcl.IsZeroValue(des.MaxPodsPerNode) {
		des.MaxPodsPerNode = initial.MaxPodsPerNode
	}

	return des
}

func canonicalizeNewClusterDefaultMaxPodsConstraint(c *Client, des, nw *ClusterDefaultMaxPodsConstraint) *ClusterDefaultMaxPodsConstraint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MaxPodsPerNode, nw.MaxPodsPerNode) {
		nw.MaxPodsPerNode = des.MaxPodsPerNode
	}

	return nw
}

func canonicalizeNewClusterDefaultMaxPodsConstraintSet(c *Client, des, nw []ClusterDefaultMaxPodsConstraint) []ClusterDefaultMaxPodsConstraint {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterDefaultMaxPodsConstraint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterDefaultMaxPodsConstraint(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterDefaultMaxPodsConstraintSlice(c *Client, des, nw []ClusterDefaultMaxPodsConstraint) []ClusterDefaultMaxPodsConstraint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterDefaultMaxPodsConstraint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterDefaultMaxPodsConstraint(c, &d, &n))
	}

	return items
}

func canonicalizeClusterResourceUsageExportConfig(des, initial *ClusterResourceUsageExportConfig, opts ...dcl.ApplyOption) *ClusterResourceUsageExportConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BigqueryDestination = canonicalizeClusterResourceUsageExportConfigBigqueryDestination(des.BigqueryDestination, initial.BigqueryDestination, opts...)
	if dcl.BoolCanonicalize(des.EnableNetworkEgressMonitoring, initial.EnableNetworkEgressMonitoring) || dcl.IsZeroValue(des.EnableNetworkEgressMonitoring) {
		des.EnableNetworkEgressMonitoring = initial.EnableNetworkEgressMonitoring
	}
	des.ConsumptionMeteringConfig = canonicalizeClusterResourceUsageExportConfigConsumptionMeteringConfig(des.ConsumptionMeteringConfig, initial.ConsumptionMeteringConfig, opts...)
	if dcl.BoolCanonicalize(des.EnableNetworkEgressMetering, initial.EnableNetworkEgressMetering) || dcl.IsZeroValue(des.EnableNetworkEgressMetering) {
		des.EnableNetworkEgressMetering = initial.EnableNetworkEgressMetering
	}

	return des
}

func canonicalizeNewClusterResourceUsageExportConfig(c *Client, des, nw *ClusterResourceUsageExportConfig) *ClusterResourceUsageExportConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.BigqueryDestination = canonicalizeNewClusterResourceUsageExportConfigBigqueryDestination(c, des.BigqueryDestination, nw.BigqueryDestination)
	if dcl.BoolCanonicalize(des.EnableNetworkEgressMonitoring, nw.EnableNetworkEgressMonitoring) {
		nw.EnableNetworkEgressMonitoring = des.EnableNetworkEgressMonitoring
	}
	nw.ConsumptionMeteringConfig = canonicalizeNewClusterResourceUsageExportConfigConsumptionMeteringConfig(c, des.ConsumptionMeteringConfig, nw.ConsumptionMeteringConfig)
	if dcl.BoolCanonicalize(des.EnableNetworkEgressMetering, nw.EnableNetworkEgressMetering) {
		nw.EnableNetworkEgressMetering = des.EnableNetworkEgressMetering
	}

	return nw
}

func canonicalizeNewClusterResourceUsageExportConfigSet(c *Client, des, nw []ClusterResourceUsageExportConfig) []ClusterResourceUsageExportConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterResourceUsageExportConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterResourceUsageExportConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterResourceUsageExportConfigSlice(c *Client, des, nw []ClusterResourceUsageExportConfig) []ClusterResourceUsageExportConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterResourceUsageExportConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterResourceUsageExportConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterResourceUsageExportConfigBigqueryDestination(des, initial *ClusterResourceUsageExportConfigBigqueryDestination, opts ...dcl.ApplyOption) *ClusterResourceUsageExportConfigBigqueryDestination {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.DatasetId, initial.DatasetId) || dcl.IsZeroValue(des.DatasetId) {
		des.DatasetId = initial.DatasetId
	}

	return des
}

func canonicalizeNewClusterResourceUsageExportConfigBigqueryDestination(c *Client, des, nw *ClusterResourceUsageExportConfigBigqueryDestination) *ClusterResourceUsageExportConfigBigqueryDestination {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.DatasetId, nw.DatasetId) {
		nw.DatasetId = des.DatasetId
	}

	return nw
}

func canonicalizeNewClusterResourceUsageExportConfigBigqueryDestinationSet(c *Client, des, nw []ClusterResourceUsageExportConfigBigqueryDestination) []ClusterResourceUsageExportConfigBigqueryDestination {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterResourceUsageExportConfigBigqueryDestination
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterResourceUsageExportConfigBigqueryDestination(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterResourceUsageExportConfigBigqueryDestinationSlice(c *Client, des, nw []ClusterResourceUsageExportConfigBigqueryDestination) []ClusterResourceUsageExportConfigBigqueryDestination {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterResourceUsageExportConfigBigqueryDestination
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterResourceUsageExportConfigBigqueryDestination(c, &d, &n))
	}

	return items
}

func canonicalizeClusterResourceUsageExportConfigConsumptionMeteringConfig(des, initial *ClusterResourceUsageExportConfigConsumptionMeteringConfig, opts ...dcl.ApplyOption) *ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterResourceUsageExportConfigConsumptionMeteringConfig(c *Client, des, nw *ClusterResourceUsageExportConfigConsumptionMeteringConfig) *ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterResourceUsageExportConfigConsumptionMeteringConfigSet(c *Client, des, nw []ClusterResourceUsageExportConfigConsumptionMeteringConfig) []ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterResourceUsageExportConfigConsumptionMeteringConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterResourceUsageExportConfigConsumptionMeteringConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterResourceUsageExportConfigConsumptionMeteringConfigSlice(c *Client, des, nw []ClusterResourceUsageExportConfigConsumptionMeteringConfig) []ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterResourceUsageExportConfigConsumptionMeteringConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterResourceUsageExportConfigConsumptionMeteringConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAuthenticatorGroupsConfig(des, initial *ClusterAuthenticatorGroupsConfig, opts ...dcl.ApplyOption) *ClusterAuthenticatorGroupsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.StringCanonicalize(des.SecurityGroup, initial.SecurityGroup) || dcl.IsZeroValue(des.SecurityGroup) {
		des.SecurityGroup = initial.SecurityGroup
	}

	return des
}

func canonicalizeNewClusterAuthenticatorGroupsConfig(c *Client, des, nw *ClusterAuthenticatorGroupsConfig) *ClusterAuthenticatorGroupsConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.StringCanonicalize(des.SecurityGroup, nw.SecurityGroup) {
		nw.SecurityGroup = des.SecurityGroup
	}

	return nw
}

func canonicalizeNewClusterAuthenticatorGroupsConfigSet(c *Client, des, nw []ClusterAuthenticatorGroupsConfig) []ClusterAuthenticatorGroupsConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAuthenticatorGroupsConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAuthenticatorGroupsConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAuthenticatorGroupsConfigSlice(c *Client, des, nw []ClusterAuthenticatorGroupsConfig) []ClusterAuthenticatorGroupsConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAuthenticatorGroupsConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAuthenticatorGroupsConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterPrivateClusterConfig(des, initial *ClusterPrivateClusterConfig, opts ...dcl.ApplyOption) *ClusterPrivateClusterConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.EnablePrivateNodes, initial.EnablePrivateNodes) || dcl.IsZeroValue(des.EnablePrivateNodes) {
		des.EnablePrivateNodes = initial.EnablePrivateNodes
	}
	if dcl.BoolCanonicalize(des.EnablePrivateEndpoint, initial.EnablePrivateEndpoint) || dcl.IsZeroValue(des.EnablePrivateEndpoint) {
		des.EnablePrivateEndpoint = initial.EnablePrivateEndpoint
	}
	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, initial.MasterIPv4CidrBlock) || dcl.IsZeroValue(des.MasterIPv4CidrBlock) {
		des.MasterIPv4CidrBlock = initial.MasterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.PrivateEndpoint, initial.PrivateEndpoint) || dcl.IsZeroValue(des.PrivateEndpoint) {
		des.PrivateEndpoint = initial.PrivateEndpoint
	}
	if dcl.StringCanonicalize(des.PublicEndpoint, initial.PublicEndpoint) || dcl.IsZeroValue(des.PublicEndpoint) {
		des.PublicEndpoint = initial.PublicEndpoint
	}
	if dcl.StringCanonicalize(des.PeeringName, initial.PeeringName) || dcl.IsZeroValue(des.PeeringName) {
		des.PeeringName = initial.PeeringName
	}
	des.MasterGlobalAccessConfig = canonicalizeClusterPrivateClusterConfigMasterGlobalAccessConfig(des.MasterGlobalAccessConfig, initial.MasterGlobalAccessConfig, opts...)

	return des
}

func canonicalizeNewClusterPrivateClusterConfig(c *Client, des, nw *ClusterPrivateClusterConfig) *ClusterPrivateClusterConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnablePrivateNodes, nw.EnablePrivateNodes) {
		nw.EnablePrivateNodes = des.EnablePrivateNodes
	}
	if dcl.BoolCanonicalize(des.EnablePrivateEndpoint, nw.EnablePrivateEndpoint) {
		nw.EnablePrivateEndpoint = des.EnablePrivateEndpoint
	}
	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, nw.MasterIPv4CidrBlock) {
		nw.MasterIPv4CidrBlock = des.MasterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.PrivateEndpoint, nw.PrivateEndpoint) {
		nw.PrivateEndpoint = des.PrivateEndpoint
	}
	if dcl.StringCanonicalize(des.PublicEndpoint, nw.PublicEndpoint) {
		nw.PublicEndpoint = des.PublicEndpoint
	}
	if dcl.StringCanonicalize(des.PeeringName, nw.PeeringName) {
		nw.PeeringName = des.PeeringName
	}
	nw.MasterGlobalAccessConfig = canonicalizeNewClusterPrivateClusterConfigMasterGlobalAccessConfig(c, des.MasterGlobalAccessConfig, nw.MasterGlobalAccessConfig)

	return nw
}

func canonicalizeNewClusterPrivateClusterConfigSet(c *Client, des, nw []ClusterPrivateClusterConfig) []ClusterPrivateClusterConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterPrivateClusterConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterPrivateClusterConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterPrivateClusterConfigSlice(c *Client, des, nw []ClusterPrivateClusterConfig) []ClusterPrivateClusterConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterPrivateClusterConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterPrivateClusterConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterPrivateClusterConfigMasterGlobalAccessConfig(des, initial *ClusterPrivateClusterConfigMasterGlobalAccessConfig, opts ...dcl.ApplyOption) *ClusterPrivateClusterConfigMasterGlobalAccessConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterPrivateClusterConfigMasterGlobalAccessConfig(c *Client, des, nw *ClusterPrivateClusterConfigMasterGlobalAccessConfig) *ClusterPrivateClusterConfigMasterGlobalAccessConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterPrivateClusterConfigMasterGlobalAccessConfigSet(c *Client, des, nw []ClusterPrivateClusterConfigMasterGlobalAccessConfig) []ClusterPrivateClusterConfigMasterGlobalAccessConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterPrivateClusterConfigMasterGlobalAccessConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterPrivateClusterConfigMasterGlobalAccessConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterPrivateClusterConfigMasterGlobalAccessConfigSlice(c *Client, des, nw []ClusterPrivateClusterConfigMasterGlobalAccessConfig) []ClusterPrivateClusterConfigMasterGlobalAccessConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterPrivateClusterConfigMasterGlobalAccessConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterPrivateClusterConfigMasterGlobalAccessConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterDatabaseEncryption(des, initial *ClusterDatabaseEncryption, opts ...dcl.ApplyOption) *ClusterDatabaseEncryption {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.StringCanonicalize(des.KeyName, initial.KeyName) || dcl.IsZeroValue(des.KeyName) {
		des.KeyName = initial.KeyName
	}

	return des
}

func canonicalizeNewClusterDatabaseEncryption(c *Client, des, nw *ClusterDatabaseEncryption) *ClusterDatabaseEncryption {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KeyName, nw.KeyName) {
		nw.KeyName = des.KeyName
	}

	return nw
}

func canonicalizeNewClusterDatabaseEncryptionSet(c *Client, des, nw []ClusterDatabaseEncryption) []ClusterDatabaseEncryption {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterDatabaseEncryption
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterDatabaseEncryption(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterDatabaseEncryptionSlice(c *Client, des, nw []ClusterDatabaseEncryption) []ClusterDatabaseEncryption {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterDatabaseEncryption
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterDatabaseEncryption(c, &d, &n))
	}

	return items
}

func canonicalizeClusterVerticalPodAutoscaling(des, initial *ClusterVerticalPodAutoscaling, opts ...dcl.ApplyOption) *ClusterVerticalPodAutoscaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterVerticalPodAutoscaling(c *Client, des, nw *ClusterVerticalPodAutoscaling) *ClusterVerticalPodAutoscaling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterVerticalPodAutoscalingSet(c *Client, des, nw []ClusterVerticalPodAutoscaling) []ClusterVerticalPodAutoscaling {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterVerticalPodAutoscaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterVerticalPodAutoscaling(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterVerticalPodAutoscalingSlice(c *Client, des, nw []ClusterVerticalPodAutoscaling) []ClusterVerticalPodAutoscaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterVerticalPodAutoscaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterVerticalPodAutoscaling(c, &d, &n))
	}

	return items
}

func canonicalizeClusterShieldedNodes(des, initial *ClusterShieldedNodes, opts ...dcl.ApplyOption) *ClusterShieldedNodes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterShieldedNodes(c *Client, des, nw *ClusterShieldedNodes) *ClusterShieldedNodes {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterShieldedNodesSet(c *Client, des, nw []ClusterShieldedNodes) []ClusterShieldedNodes {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterShieldedNodes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterShieldedNodes(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterShieldedNodesSlice(c *Client, des, nw []ClusterShieldedNodes) []ClusterShieldedNodes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterShieldedNodes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterShieldedNodes(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConditions(des, initial *ClusterConditions, opts ...dcl.ApplyOption) *ClusterConditions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Code, initial.Code) || dcl.IsZeroValue(des.Code) {
		des.Code = initial.Code
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		des.Message = initial.Message
	}
	if dcl.IsZeroValue(des.CanonicalCode) {
		des.CanonicalCode = initial.CanonicalCode
	}

	return des
}

func canonicalizeNewClusterConditions(c *Client, des, nw *ClusterConditions) *ClusterConditions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Code, nw.Code) {
		nw.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}

	return nw
}

func canonicalizeNewClusterConditionsSet(c *Client, des, nw []ClusterConditions) []ClusterConditions {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConditions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterConditions(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterConditionsSlice(c *Client, des, nw []ClusterConditions) []ClusterConditions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterConditions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConditions(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAutopilot(des, initial *ClusterAutopilot, opts ...dcl.ApplyOption) *ClusterAutopilot {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterAutopilot(c *Client, des, nw *ClusterAutopilot) *ClusterAutopilot {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterAutopilotSet(c *Client, des, nw []ClusterAutopilot) []ClusterAutopilot {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAutopilot
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterAutopilot(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterAutopilotSlice(c *Client, des, nw []ClusterAutopilot) []ClusterAutopilot {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterAutopilot
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAutopilot(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodeConfig(des, initial *ClusterNodeConfig, opts ...dcl.ApplyOption) *ClusterNodeConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		des.MachineType = initial.MachineType
	}
	if dcl.IsZeroValue(des.DiskSizeGb) {
		des.DiskSizeGb = initial.DiskSizeGb
	}
	if dcl.IsZeroValue(des.OAuthScopes) {
		des.OAuthScopes = initial.OAuthScopes
	}
	if dcl.StringCanonicalize(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		des.ServiceAccount = initial.ServiceAccount
	}
	if dcl.IsZeroValue(des.Metadata) {
		des.Metadata = initial.Metadata
	}
	if dcl.StringCanonicalize(des.ImageType, initial.ImageType) || dcl.IsZeroValue(des.ImageType) {
		des.ImageType = initial.ImageType
	}
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.LocalSsdCount) {
		des.LocalSsdCount = initial.LocalSsdCount
	}
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	}
	if dcl.BoolCanonicalize(des.Preemptible, initial.Preemptible) || dcl.IsZeroValue(des.Preemptible) {
		des.Preemptible = initial.Preemptible
	}
	if dcl.IsZeroValue(des.Accelerators) {
		des.Accelerators = initial.Accelerators
	}
	if dcl.StringCanonicalize(des.DiskType, initial.DiskType) || dcl.IsZeroValue(des.DiskType) {
		des.DiskType = initial.DiskType
	}
	if dcl.StringCanonicalize(des.MinCpuPlatform, initial.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
		des.MinCpuPlatform = initial.MinCpuPlatform
	}
	des.WorkloadMetadataConfig = canonicalizeClusterNodeConfigWorkloadMetadataConfig(des.WorkloadMetadataConfig, initial.WorkloadMetadataConfig, opts...)
	if dcl.IsZeroValue(des.Taints) {
		des.Taints = initial.Taints
	}
	des.SandboxConfig = canonicalizeClusterNodeConfigSandboxConfig(des.SandboxConfig, initial.SandboxConfig, opts...)
	if dcl.StringCanonicalize(des.NodeGroup, initial.NodeGroup) || dcl.IsZeroValue(des.NodeGroup) {
		des.NodeGroup = initial.NodeGroup
	}
	des.ReservationAffinity = canonicalizeClusterNodeConfigReservationAffinity(des.ReservationAffinity, initial.ReservationAffinity, opts...)
	des.ShieldedInstanceConfig = canonicalizeClusterNodeConfigShieldedInstanceConfig(des.ShieldedInstanceConfig, initial.ShieldedInstanceConfig, opts...)
	des.LinuxNodeConfig = canonicalizeClusterNodeConfigLinuxNodeConfig(des.LinuxNodeConfig, initial.LinuxNodeConfig, opts...)
	des.KubeletConfig = canonicalizeClusterNodeConfigKubeletConfig(des.KubeletConfig, initial.KubeletConfig, opts...)
	if dcl.StringCanonicalize(des.BootDiskKmsKey, initial.BootDiskKmsKey) || dcl.IsZeroValue(des.BootDiskKmsKey) {
		des.BootDiskKmsKey = initial.BootDiskKmsKey
	}

	return des
}

func canonicalizeNewClusterNodeConfig(c *Client, des, nw *ClusterNodeConfig) *ClusterNodeConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	if dcl.StringCanonicalize(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	if dcl.StringCanonicalize(des.ImageType, nw.ImageType) {
		nw.ImageType = des.ImageType
	}
	if dcl.BoolCanonicalize(des.Preemptible, nw.Preemptible) {
		nw.Preemptible = des.Preemptible
	}
	nw.Accelerators = canonicalizeNewClusterNodeConfigAcceleratorsSlice(c, des.Accelerators, nw.Accelerators)
	if dcl.StringCanonicalize(des.DiskType, nw.DiskType) {
		nw.DiskType = des.DiskType
	}
	if dcl.StringCanonicalize(des.MinCpuPlatform, nw.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}
	nw.WorkloadMetadataConfig = canonicalizeNewClusterNodeConfigWorkloadMetadataConfig(c, des.WorkloadMetadataConfig, nw.WorkloadMetadataConfig)
	nw.Taints = canonicalizeNewClusterNodeConfigTaintsSlice(c, des.Taints, nw.Taints)
	nw.SandboxConfig = canonicalizeNewClusterNodeConfigSandboxConfig(c, des.SandboxConfig, nw.SandboxConfig)
	if dcl.StringCanonicalize(des.NodeGroup, nw.NodeGroup) {
		nw.NodeGroup = des.NodeGroup
	}
	nw.ReservationAffinity = canonicalizeNewClusterNodeConfigReservationAffinity(c, des.ReservationAffinity, nw.ReservationAffinity)
	nw.ShieldedInstanceConfig = canonicalizeNewClusterNodeConfigShieldedInstanceConfig(c, des.ShieldedInstanceConfig, nw.ShieldedInstanceConfig)
	nw.LinuxNodeConfig = canonicalizeNewClusterNodeConfigLinuxNodeConfig(c, des.LinuxNodeConfig, nw.LinuxNodeConfig)
	nw.KubeletConfig = canonicalizeNewClusterNodeConfigKubeletConfig(c, des.KubeletConfig, nw.KubeletConfig)
	if dcl.StringCanonicalize(des.BootDiskKmsKey, nw.BootDiskKmsKey) {
		nw.BootDiskKmsKey = des.BootDiskKmsKey
	}

	return nw
}

func canonicalizeNewClusterNodeConfigSet(c *Client, des, nw []ClusterNodeConfig) []ClusterNodeConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodeConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodeConfigSlice(c *Client, des, nw []ClusterNodeConfig) []ClusterNodeConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodeConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodeConfigAccelerators(des, initial *ClusterNodeConfigAccelerators, opts ...dcl.ApplyOption) *ClusterNodeConfigAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AcceleratorCount) {
		des.AcceleratorCount = initial.AcceleratorCount
	}
	if dcl.StringCanonicalize(des.AcceleratorType, initial.AcceleratorType) || dcl.IsZeroValue(des.AcceleratorType) {
		des.AcceleratorType = initial.AcceleratorType
	}

	return des
}

func canonicalizeNewClusterNodeConfigAccelerators(c *Client, des, nw *ClusterNodeConfigAccelerators) *ClusterNodeConfigAccelerators {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AcceleratorType, nw.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
	}

	return nw
}

func canonicalizeNewClusterNodeConfigAcceleratorsSet(c *Client, des, nw []ClusterNodeConfigAccelerators) []ClusterNodeConfigAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeConfigAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodeConfigAccelerators(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodeConfigAcceleratorsSlice(c *Client, des, nw []ClusterNodeConfigAccelerators) []ClusterNodeConfigAccelerators {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodeConfigAccelerators
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeConfigAccelerators(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodeConfigWorkloadMetadataConfig(des, initial *ClusterNodeConfigWorkloadMetadataConfig, opts ...dcl.ApplyOption) *ClusterNodeConfigWorkloadMetadataConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	}

	return des
}

func canonicalizeNewClusterNodeConfigWorkloadMetadataConfig(c *Client, des, nw *ClusterNodeConfigWorkloadMetadataConfig) *ClusterNodeConfigWorkloadMetadataConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterNodeConfigWorkloadMetadataConfigSet(c *Client, des, nw []ClusterNodeConfigWorkloadMetadataConfig) []ClusterNodeConfigWorkloadMetadataConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeConfigWorkloadMetadataConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodeConfigWorkloadMetadataConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodeConfigWorkloadMetadataConfigSlice(c *Client, des, nw []ClusterNodeConfigWorkloadMetadataConfig) []ClusterNodeConfigWorkloadMetadataConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodeConfigWorkloadMetadataConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeConfigWorkloadMetadataConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodeConfigTaints(des, initial *ClusterNodeConfigTaints, opts ...dcl.ApplyOption) *ClusterNodeConfigTaints {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}
	if dcl.IsZeroValue(des.Effect) {
		des.Effect = initial.Effect
	}

	return des
}

func canonicalizeNewClusterNodeConfigTaints(c *Client, des, nw *ClusterNodeConfigTaints) *ClusterNodeConfigTaints {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewClusterNodeConfigTaintsSet(c *Client, des, nw []ClusterNodeConfigTaints) []ClusterNodeConfigTaints {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeConfigTaints
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodeConfigTaints(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodeConfigTaintsSlice(c *Client, des, nw []ClusterNodeConfigTaints) []ClusterNodeConfigTaints {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodeConfigTaints
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeConfigTaints(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodeConfigSandboxConfig(des, initial *ClusterNodeConfigSandboxConfig, opts ...dcl.ApplyOption) *ClusterNodeConfigSandboxConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewClusterNodeConfigSandboxConfig(c *Client, des, nw *ClusterNodeConfigSandboxConfig) *ClusterNodeConfigSandboxConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterNodeConfigSandboxConfigSet(c *Client, des, nw []ClusterNodeConfigSandboxConfig) []ClusterNodeConfigSandboxConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeConfigSandboxConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodeConfigSandboxConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodeConfigSandboxConfigSlice(c *Client, des, nw []ClusterNodeConfigSandboxConfig) []ClusterNodeConfigSandboxConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodeConfigSandboxConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeConfigSandboxConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodeConfigReservationAffinity(des, initial *ClusterNodeConfigReservationAffinity, opts ...dcl.ApplyOption) *ClusterNodeConfigReservationAffinity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ConsumeReservationType) {
		des.ConsumeReservationType = initial.ConsumeReservationType
	}
	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Values) {
		des.Values = initial.Values
	}

	return des
}

func canonicalizeNewClusterNodeConfigReservationAffinity(c *Client, des, nw *ClusterNodeConfigReservationAffinity) *ClusterNodeConfigReservationAffinity {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}

	return nw
}

func canonicalizeNewClusterNodeConfigReservationAffinitySet(c *Client, des, nw []ClusterNodeConfigReservationAffinity) []ClusterNodeConfigReservationAffinity {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeConfigReservationAffinity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodeConfigReservationAffinity(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodeConfigReservationAffinitySlice(c *Client, des, nw []ClusterNodeConfigReservationAffinity) []ClusterNodeConfigReservationAffinity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodeConfigReservationAffinity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeConfigReservationAffinity(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodeConfigShieldedInstanceConfig(des, initial *ClusterNodeConfigShieldedInstanceConfig, opts ...dcl.ApplyOption) *ClusterNodeConfigShieldedInstanceConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.EnableSecureBoot, initial.EnableSecureBoot) || dcl.IsZeroValue(des.EnableSecureBoot) {
		des.EnableSecureBoot = initial.EnableSecureBoot
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, initial.EnableIntegrityMonitoring) || dcl.IsZeroValue(des.EnableIntegrityMonitoring) {
		des.EnableIntegrityMonitoring = initial.EnableIntegrityMonitoring
	}

	return des
}

func canonicalizeNewClusterNodeConfigShieldedInstanceConfig(c *Client, des, nw *ClusterNodeConfigShieldedInstanceConfig) *ClusterNodeConfigShieldedInstanceConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnableSecureBoot, nw.EnableSecureBoot) {
		nw.EnableSecureBoot = des.EnableSecureBoot
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, nw.EnableIntegrityMonitoring) {
		nw.EnableIntegrityMonitoring = des.EnableIntegrityMonitoring
	}

	return nw
}

func canonicalizeNewClusterNodeConfigShieldedInstanceConfigSet(c *Client, des, nw []ClusterNodeConfigShieldedInstanceConfig) []ClusterNodeConfigShieldedInstanceConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeConfigShieldedInstanceConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodeConfigShieldedInstanceConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodeConfigShieldedInstanceConfigSlice(c *Client, des, nw []ClusterNodeConfigShieldedInstanceConfig) []ClusterNodeConfigShieldedInstanceConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodeConfigShieldedInstanceConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeConfigShieldedInstanceConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodeConfigLinuxNodeConfig(des, initial *ClusterNodeConfigLinuxNodeConfig, opts ...dcl.ApplyOption) *ClusterNodeConfigLinuxNodeConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Sysctls) {
		des.Sysctls = initial.Sysctls
	}

	return des
}

func canonicalizeNewClusterNodeConfigLinuxNodeConfig(c *Client, des, nw *ClusterNodeConfigLinuxNodeConfig) *ClusterNodeConfigLinuxNodeConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterNodeConfigLinuxNodeConfigSet(c *Client, des, nw []ClusterNodeConfigLinuxNodeConfig) []ClusterNodeConfigLinuxNodeConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeConfigLinuxNodeConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodeConfigLinuxNodeConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodeConfigLinuxNodeConfigSlice(c *Client, des, nw []ClusterNodeConfigLinuxNodeConfig) []ClusterNodeConfigLinuxNodeConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodeConfigLinuxNodeConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeConfigLinuxNodeConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNodeConfigKubeletConfig(des, initial *ClusterNodeConfigKubeletConfig, opts ...dcl.ApplyOption) *ClusterNodeConfigKubeletConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CpuManagerPolicy, initial.CpuManagerPolicy) || dcl.IsZeroValue(des.CpuManagerPolicy) {
		des.CpuManagerPolicy = initial.CpuManagerPolicy
	}
	if dcl.BoolCanonicalize(des.CpuCfsQuota, initial.CpuCfsQuota) || dcl.IsZeroValue(des.CpuCfsQuota) {
		des.CpuCfsQuota = initial.CpuCfsQuota
	}
	if dcl.StringCanonicalize(des.CpuCfsQuotaPeriod, initial.CpuCfsQuotaPeriod) || dcl.IsZeroValue(des.CpuCfsQuotaPeriod) {
		des.CpuCfsQuotaPeriod = initial.CpuCfsQuotaPeriod
	}

	return des
}

func canonicalizeNewClusterNodeConfigKubeletConfig(c *Client, des, nw *ClusterNodeConfigKubeletConfig) *ClusterNodeConfigKubeletConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CpuManagerPolicy, nw.CpuManagerPolicy) {
		nw.CpuManagerPolicy = des.CpuManagerPolicy
	}
	if dcl.BoolCanonicalize(des.CpuCfsQuota, nw.CpuCfsQuota) {
		nw.CpuCfsQuota = des.CpuCfsQuota
	}
	if dcl.StringCanonicalize(des.CpuCfsQuotaPeriod, nw.CpuCfsQuotaPeriod) {
		nw.CpuCfsQuotaPeriod = des.CpuCfsQuotaPeriod
	}

	return nw
}

func canonicalizeNewClusterNodeConfigKubeletConfigSet(c *Client, des, nw []ClusterNodeConfigKubeletConfig) []ClusterNodeConfigKubeletConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNodeConfigKubeletConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNodeConfigKubeletConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNodeConfigKubeletConfigSlice(c *Client, des, nw []ClusterNodeConfigKubeletConfig) []ClusterNodeConfigKubeletConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNodeConfigKubeletConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNodeConfigKubeletConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterReleaseChannel(des, initial *ClusterReleaseChannel, opts ...dcl.ApplyOption) *ClusterReleaseChannel {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Channel) {
		des.Channel = initial.Channel
	}

	return des
}

func canonicalizeNewClusterReleaseChannel(c *Client, des, nw *ClusterReleaseChannel) *ClusterReleaseChannel {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewClusterReleaseChannelSet(c *Client, des, nw []ClusterReleaseChannel) []ClusterReleaseChannel {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterReleaseChannel
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterReleaseChannel(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterReleaseChannelSlice(c *Client, des, nw []ClusterReleaseChannel) []ClusterReleaseChannel {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterReleaseChannel
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterReleaseChannel(c, &d, &n))
	}

	return items
}

func canonicalizeClusterWorkloadIdentityConfig(des, initial *ClusterWorkloadIdentityConfig, opts ...dcl.ApplyOption) *ClusterWorkloadIdentityConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.WorkloadPool, initial.WorkloadPool) || dcl.IsZeroValue(des.WorkloadPool) {
		des.WorkloadPool = initial.WorkloadPool
	}

	return des
}

func canonicalizeNewClusterWorkloadIdentityConfig(c *Client, des, nw *ClusterWorkloadIdentityConfig) *ClusterWorkloadIdentityConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.WorkloadPool, nw.WorkloadPool) {
		nw.WorkloadPool = des.WorkloadPool
	}

	return nw
}

func canonicalizeNewClusterWorkloadIdentityConfigSet(c *Client, des, nw []ClusterWorkloadIdentityConfig) []ClusterWorkloadIdentityConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterWorkloadIdentityConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterWorkloadIdentityConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterWorkloadIdentityConfigSlice(c *Client, des, nw []ClusterWorkloadIdentityConfig) []ClusterWorkloadIdentityConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterWorkloadIdentityConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterWorkloadIdentityConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNotificationConfig(des, initial *ClusterNotificationConfig, opts ...dcl.ApplyOption) *ClusterNotificationConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Pubsub = canonicalizeClusterNotificationConfigPubsub(des.Pubsub, initial.Pubsub, opts...)

	return des
}

func canonicalizeNewClusterNotificationConfig(c *Client, des, nw *ClusterNotificationConfig) *ClusterNotificationConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.Pubsub = canonicalizeNewClusterNotificationConfigPubsub(c, des.Pubsub, nw.Pubsub)

	return nw
}

func canonicalizeNewClusterNotificationConfigSet(c *Client, des, nw []ClusterNotificationConfig) []ClusterNotificationConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNotificationConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNotificationConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNotificationConfigSlice(c *Client, des, nw []ClusterNotificationConfig) []ClusterNotificationConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNotificationConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNotificationConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterNotificationConfigPubsub(des, initial *ClusterNotificationConfigPubsub, opts ...dcl.ApplyOption) *ClusterNotificationConfigPubsub {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.NameToSelfLink(des.Topic, initial.Topic) || dcl.IsZeroValue(des.Topic) {
		des.Topic = initial.Topic
	}

	return des
}

func canonicalizeNewClusterNotificationConfigPubsub(c *Client, des, nw *ClusterNotificationConfigPubsub) *ClusterNotificationConfigPubsub {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.NameToSelfLink(des.Topic, nw.Topic) {
		nw.Topic = des.Topic
	}

	return nw
}

func canonicalizeNewClusterNotificationConfigPubsubSet(c *Client, des, nw []ClusterNotificationConfigPubsub) []ClusterNotificationConfigPubsub {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNotificationConfigPubsub
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterNotificationConfigPubsub(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterNotificationConfigPubsubSlice(c *Client, des, nw []ClusterNotificationConfigPubsub) []ClusterNotificationConfigPubsub {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterNotificationConfigPubsub
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNotificationConfigPubsub(c, &d, &n))
	}

	return items
}

func canonicalizeClusterConfidentialNodes(des, initial *ClusterConfidentialNodes, opts ...dcl.ApplyOption) *ClusterConfidentialNodes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterConfidentialNodes(c *Client, des, nw *ClusterConfidentialNodes) *ClusterConfidentialNodes {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewClusterConfidentialNodesSet(c *Client, des, nw []ClusterConfidentialNodes) []ClusterConfidentialNodes {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterConfidentialNodes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareClusterConfidentialNodes(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewClusterConfidentialNodesSlice(c *Client, des, nw []ClusterConfidentialNodes) []ClusterConfidentialNodes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ClusterConfidentialNodes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterConfidentialNodes(c, &d, &n))
	}

	return items
}

type clusterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         clusterApiOperation
	Diffs            []*dcl.FieldDiff
	// This is for reporting only.
	FieldName string
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffCluster(c *Client, desired, actual *Cluster, opts ...dcl.ApplyOption) ([]clusterDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []clusterDiff

	var fn dcl.FieldName

	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Name",
		})
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Description",
		})
	}

	if ds, err := dcl.Diff(desired.InitialNodeCount, actual.InitialNodeCount, dcl.Info{}, fn.AddNest("InitialNodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "InitialNodeCount",
		})
	}

	if ds, err := dcl.Diff(desired.MasterAuth, actual.MasterAuth, dcl.Info{ObjectFunction: compareClusterMasterAuthNewStyle}, fn.AddNest("MasterAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "MasterAuth",
		})
	}

	if ds, err := dcl.Diff(desired.LoggingService, actual.LoggingService, dcl.Info{}, fn.AddNest("LoggingService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateMonitoringAndLoggingServiceOperation{}, Diffs: ds,
			FieldName: "LoggingService",
		})
	}

	if ds, err := dcl.Diff(desired.MonitoringService, actual.MonitoringService, dcl.Info{}, fn.AddNest("MonitoringService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateMonitoringAndLoggingServiceOperation{}, Diffs: ds,
			FieldName: "MonitoringService",
		})
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Network",
		})
	}

	if ds, err := dcl.Diff(desired.ClusterIPv4Cidr, actual.ClusterIPv4Cidr, dcl.Info{}, fn.AddNest("ClusterIPv4Cidr")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ClusterIPv4Cidr",
		})
	}

	if ds, err := dcl.Diff(desired.AddonsConfig, actual.AddonsConfig, dcl.Info{ObjectFunction: compareClusterAddonsConfigNewStyle}, fn.AddNest("AddonsConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateAddonsConfigOperation{}, Diffs: ds,
			FieldName: "AddonsConfig",
		})
	}

	if ds, err := dcl.Diff(desired.Subnetwork, actual.Subnetwork, dcl.Info{}, fn.AddNest("Subnetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Subnetwork",
		})
	}

	if ds, err := dcl.Diff(desired.NodePools, actual.NodePools, dcl.Info{ObjectFunction: compareClusterNodePoolsNewStyle}, fn.AddNest("NodePools")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "NodePools",
		})
	}

	if ds, err := dcl.Diff(desired.Locations, actual.Locations, dcl.Info{}, fn.AddNest("Locations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateLocationsOperation{}, Diffs: ds,
			FieldName: "Locations",
		})
	}

	if ds, err := dcl.Diff(desired.EnableKubernetesAlpha, actual.EnableKubernetesAlpha, dcl.Info{}, fn.AddNest("EnableKubernetesAlpha")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "EnableKubernetesAlpha",
		})
	}

	if ds, err := dcl.Diff(desired.ResourceLabels, actual.ResourceLabels, dcl.Info{}, fn.AddNest("ResourceLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ResourceLabels",
		})
	}

	if ds, err := dcl.Diff(desired.LabelFingerprint, actual.LabelFingerprint, dcl.Info{}, fn.AddNest("LabelFingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "LabelFingerprint",
		})
	}

	if ds, err := dcl.Diff(desired.LegacyAbac, actual.LegacyAbac, dcl.Info{ObjectFunction: compareClusterLegacyAbacNewStyle}, fn.AddNest("LegacyAbac")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateLegacyAbacOperation{}, Diffs: ds,
			FieldName: "LegacyAbac",
		})
	}

	if ds, err := dcl.Diff(desired.NetworkPolicy, actual.NetworkPolicy, dcl.Info{ObjectFunction: compareClusterNetworkPolicyNewStyle}, fn.AddNest("NetworkPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "NetworkPolicy",
		})
	}

	if ds, err := dcl.Diff(desired.IPAllocationPolicy, actual.IPAllocationPolicy, dcl.Info{ObjectFunction: compareClusterIPAllocationPolicyNewStyle}, fn.AddNest("IPAllocationPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "IPAllocationPolicy",
		})
	}

	if ds, err := dcl.Diff(desired.MasterAuthorizedNetworksConfig, actual.MasterAuthorizedNetworksConfig, dcl.Info{ObjectFunction: compareClusterMasterAuthorizedNetworksConfigNewStyle}, fn.AddNest("MasterAuthorizedNetworksConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateMasterAuthorizedNetworksConfigOperation{}, Diffs: ds,
			FieldName: "MasterAuthorizedNetworksConfig",
		})
	}

	if ds, err := dcl.Diff(desired.BinaryAuthorization, actual.BinaryAuthorization, dcl.Info{ObjectFunction: compareClusterBinaryAuthorizationNewStyle}, fn.AddNest("BinaryAuthorization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateBinaryAuthorizationOperation{}, Diffs: ds,
			FieldName: "BinaryAuthorization",
		})
	}

	if ds, err := dcl.Diff(desired.Autoscaling, actual.Autoscaling, dcl.Info{ObjectFunction: compareClusterAutoscalingNewStyle}, fn.AddNest("Autoscaling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Autoscaling",
		})
	}

	if ds, err := dcl.Diff(desired.NetworkConfig, actual.NetworkConfig, dcl.Info{ObjectFunction: compareClusterNetworkConfigNewStyle}, fn.AddNest("NetworkConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "NetworkConfig",
		})
	}

	if ds, err := dcl.Diff(desired.MaintenancePolicy, actual.MaintenancePolicy, dcl.Info{ObjectFunction: compareClusterMaintenancePolicyNewStyle}, fn.AddNest("MaintenancePolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterSetMaintenancePolicyOperation{}, Diffs: ds,
			FieldName: "MaintenancePolicy",
		})
	}

	if ds, err := dcl.Diff(desired.DefaultMaxPodsConstraint, actual.DefaultMaxPodsConstraint, dcl.Info{ObjectFunction: compareClusterDefaultMaxPodsConstraintNewStyle}, fn.AddNest("DefaultMaxPodsConstraint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "DefaultMaxPodsConstraint",
		})
	}

	if ds, err := dcl.Diff(desired.ResourceUsageExportConfig, actual.ResourceUsageExportConfig, dcl.Info{ObjectFunction: compareClusterResourceUsageExportConfigNewStyle}, fn.AddNest("ResourceUsageExportConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ResourceUsageExportConfig",
		})
	}

	if ds, err := dcl.Diff(desired.AuthenticatorGroupsConfig, actual.AuthenticatorGroupsConfig, dcl.Info{ObjectFunction: compareClusterAuthenticatorGroupsConfigNewStyle}, fn.AddNest("AuthenticatorGroupsConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "AuthenticatorGroupsConfig",
		})
	}

	if ds, err := dcl.Diff(desired.PrivateClusterConfig, actual.PrivateClusterConfig, dcl.Info{ObjectFunction: compareClusterPrivateClusterConfigNewStyle}, fn.AddNest("PrivateClusterConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "PrivateClusterConfig",
		})
	}

	if ds, err := dcl.Diff(desired.DatabaseEncryption, actual.DatabaseEncryption, dcl.Info{ObjectFunction: compareClusterDatabaseEncryptionNewStyle}, fn.AddNest("DatabaseEncryption")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateDatabaseEncryptionOperation{}, Diffs: ds,
			FieldName: "DatabaseEncryption",
		})
	}

	if ds, err := dcl.Diff(desired.VerticalPodAutoscaling, actual.VerticalPodAutoscaling, dcl.Info{ObjectFunction: compareClusterVerticalPodAutoscalingNewStyle}, fn.AddNest("VerticalPodAutoscaling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateVerticalPodAutoscalingOperation{}, Diffs: ds,
			FieldName: "VerticalPodAutoscaling",
		})
	}

	if ds, err := dcl.Diff(desired.ShieldedNodes, actual.ShieldedNodes, dcl.Info{ObjectFunction: compareClusterShieldedNodesNewStyle}, fn.AddNest("ShieldedNodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateShieldedNodesOperation{}, Diffs: ds,
			FieldName: "ShieldedNodes",
		})
	}

	if ds, err := dcl.Diff(desired.Endpoint, actual.Endpoint, dcl.Info{}, fn.AddNest("Endpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Endpoint",
		})
	}

	if ds, err := dcl.Diff(desired.MasterVersion, actual.MasterVersion, dcl.Info{}, fn.AddNest("MasterVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateMasterVersionOperation{}, Diffs: ds,
			FieldName: "MasterVersion",
		})
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "CreateTime",
		})
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Status",
		})
	}

	if ds, err := dcl.Diff(desired.StatusMessage, actual.StatusMessage, dcl.Info{}, fn.AddNest("StatusMessage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "StatusMessage",
		})
	}

	if ds, err := dcl.Diff(desired.NodeIPv4CidrSize, actual.NodeIPv4CidrSize, dcl.Info{}, fn.AddNest("NodeIPv4CidrSize")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "NodeIPv4CidrSize",
		})
	}

	if ds, err := dcl.Diff(desired.ServicesIPv4Cidr, actual.ServicesIPv4Cidr, dcl.Info{}, fn.AddNest("ServicesIPv4Cidr")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ServicesIPv4Cidr",
		})
	}

	if ds, err := dcl.Diff(desired.ExpireTime, actual.ExpireTime, dcl.Info{}, fn.AddNest("ExpireTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ExpireTime",
		})
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Location",
		})
	}

	if ds, err := dcl.Diff(desired.EnableTPU, actual.EnableTPU, dcl.Info{}, fn.AddNest("EnableTPU")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "EnableTPU",
		})
	}

	if ds, err := dcl.Diff(desired.TPUIPv4CidrBlock, actual.TPUIPv4CidrBlock, dcl.Info{}, fn.AddNest("TPUIPv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "TPUIPv4CidrBlock",
		})
	}

	if ds, err := dcl.Diff(desired.Conditions, actual.Conditions, dcl.Info{ObjectFunction: compareClusterConditionsNewStyle}, fn.AddNest("Conditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Conditions",
		})
	}

	if ds, err := dcl.Diff(desired.Autopilot, actual.Autopilot, dcl.Info{ObjectFunction: compareClusterAutopilotNewStyle}, fn.AddNest("Autopilot")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Autopilot",
		})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Project",
		})
	}

	if ds, err := dcl.Diff(desired.NodeConfig, actual.NodeConfig, dcl.Info{ObjectFunction: compareClusterNodeConfigNewStyle}, fn.AddNest("NodeConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "NodeConfig",
		})
	}

	if ds, err := dcl.Diff(desired.ReleaseChannel, actual.ReleaseChannel, dcl.Info{ObjectFunction: compareClusterReleaseChannelNewStyle}, fn.AddNest("ReleaseChannel")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ReleaseChannel",
		})
	}

	if ds, err := dcl.Diff(desired.WorkloadIdentityConfig, actual.WorkloadIdentityConfig, dcl.Info{ObjectFunction: compareClusterWorkloadIdentityConfigNewStyle}, fn.AddNest("WorkloadIdentityConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{
			UpdateOp: &updateClusterUpdateWorkloadIdentityConfigOperation{}, Diffs: ds,
			FieldName: "WorkloadIdentityConfig",
		})
	}

	if ds, err := dcl.Diff(desired.NotificationConfig, actual.NotificationConfig, dcl.Info{ObjectFunction: compareClusterNotificationConfigNewStyle}, fn.AddNest("NotificationConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "NotificationConfig",
		})
	}

	if ds, err := dcl.Diff(desired.ConfidentialNodes, actual.ConfidentialNodes, dcl.Info{ObjectFunction: compareClusterConfidentialNodesNewStyle}, fn.AddNest("ConfidentialNodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ConfidentialNodes",
		})
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "SelfLink",
		})
	}

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.Info{}, fn.AddNest("Zone")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Zone",
		})
	}

	if ds, err := dcl.Diff(desired.InitialClusterVersion, actual.InitialClusterVersion, dcl.Info{}, fn.AddNest("InitialClusterVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "InitialClusterVersion",
		})
	}

	if ds, err := dcl.Diff(desired.CurrentMasterVersion, actual.CurrentMasterVersion, dcl.Info{}, fn.AddNest("CurrentMasterVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "CurrentMasterVersion",
		})
	}

	if ds, err := dcl.Diff(desired.CurrentNodeVersion, actual.CurrentNodeVersion, dcl.Info{}, fn.AddNest("CurrentNodeVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "CurrentNodeVersion",
		})
	}

	if ds, err := dcl.Diff(desired.InstanceGroupUrls, actual.InstanceGroupUrls, dcl.Info{}, fn.AddNest("InstanceGroupUrls")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "InstanceGroupUrls",
		})
	}

	if ds, err := dcl.Diff(desired.CurrentNodeCount, actual.CurrentNodeCount, dcl.Info{}, fn.AddNest("CurrentNodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "CurrentNodeCount",
		})
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, clusterDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Id",
		})
	}

	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []clusterDiff
	for _, d := range diffs {
		// Two operations are considered identical if they have the same type.
		// The type of an operation is derived from the name of the update method.
		if !dcl.StringSliceContains(fmt.Sprintf("%T", d.UpdateOp), opTypes) {
			deduped = append(deduped, d)
			opTypes = append(opTypes, fmt.Sprintf("%T", d.UpdateOp))
		} else {
			c.Config.Logger.Infof("Omitting planned operation of type %T since once is already scheduled.", d.UpdateOp)
		}
	}

	return deduped, nil
}
func compareClusterMasterAuthNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMasterAuth)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMasterAuth)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMasterAuth or *ClusterMasterAuth", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMasterAuth)
	if !ok {
		actualNotPointer, ok := a.(ClusterMasterAuth)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMasterAuth", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Username, actual.Username, dcl.Info{}, fn.AddNest("Username")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Password, actual.Password, dcl.Info{}, fn.AddNest("Password")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientCertificateConfig, actual.ClientCertificateConfig, dcl.Info{ObjectFunction: compareClusterMasterAuthClientCertificateConfigNewStyle}, fn.AddNest("ClientCertificateConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterCaCertificate, actual.ClusterCaCertificate, dcl.Info{}, fn.AddNest("ClusterCaCertificate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientCertificate, actual.ClientCertificate, dcl.Info{}, fn.AddNest("ClientCertificate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientKey, actual.ClientKey, dcl.Info{}, fn.AddNest("ClientKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMasterAuth(c *Client, desired, actual *ClusterMasterAuth) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Username, actual.Username) && !dcl.IsZeroValue(desired.Username) {
		c.Config.Logger.Infof("Diff in Username.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Username), dcl.SprintResource(actual.Username))
		return true
	}
	if !dcl.StringCanonicalize(desired.Password, actual.Password) && !dcl.IsZeroValue(desired.Password) {
		c.Config.Logger.Infof("Diff in Password.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Password), dcl.SprintResource(actual.Password))
		return true
	}
	if compareClusterMasterAuthClientCertificateConfig(c, desired.ClientCertificateConfig, actual.ClientCertificateConfig) && !dcl.IsZeroValue(desired.ClientCertificateConfig) {
		c.Config.Logger.Infof("Diff in ClientCertificateConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientCertificateConfig), dcl.SprintResource(actual.ClientCertificateConfig))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClusterCaCertificate, actual.ClusterCaCertificate) && !dcl.IsZeroValue(desired.ClusterCaCertificate) {
		c.Config.Logger.Infof("Diff in ClusterCaCertificate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterCaCertificate), dcl.SprintResource(actual.ClusterCaCertificate))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClientCertificate, actual.ClientCertificate) && !dcl.IsZeroValue(desired.ClientCertificate) {
		c.Config.Logger.Infof("Diff in ClientCertificate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientCertificate), dcl.SprintResource(actual.ClientCertificate))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClientKey, actual.ClientKey) && !dcl.IsZeroValue(desired.ClientKey) {
		c.Config.Logger.Infof("Diff in ClientKey.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientKey), dcl.SprintResource(actual.ClientKey))
		return true
	}
	return false
}

func compareClusterMasterAuthSlice(c *Client, desired, actual []ClusterMasterAuth) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMasterAuth, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMasterAuth(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMasterAuth, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthMap(c *Client, desired, actual map[string]ClusterMasterAuth) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMasterAuth, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterMasterAuth, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterMasterAuth(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterMasterAuth, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthClientCertificateConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMasterAuthClientCertificateConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMasterAuthClientCertificateConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMasterAuthClientCertificateConfig or *ClusterMasterAuthClientCertificateConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMasterAuthClientCertificateConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterMasterAuthClientCertificateConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMasterAuthClientCertificateConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IssueClientCertificate, actual.IssueClientCertificate, dcl.Info{}, fn.AddNest("IssueClientCertificate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMasterAuthClientCertificateConfig(c *Client, desired, actual *ClusterMasterAuthClientCertificateConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.IssueClientCertificate, actual.IssueClientCertificate) && !dcl.IsZeroValue(desired.IssueClientCertificate) {
		c.Config.Logger.Infof("Diff in IssueClientCertificate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IssueClientCertificate), dcl.SprintResource(actual.IssueClientCertificate))
		return true
	}
	return false
}

func compareClusterMasterAuthClientCertificateConfigSlice(c *Client, desired, actual []ClusterMasterAuthClientCertificateConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMasterAuthClientCertificateConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMasterAuthClientCertificateConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMasterAuthClientCertificateConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthClientCertificateConfigMap(c *Client, desired, actual map[string]ClusterMasterAuthClientCertificateConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMasterAuthClientCertificateConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterMasterAuthClientCertificateConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterMasterAuthClientCertificateConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterMasterAuthClientCertificateConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAddonsConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAddonsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfig or *ClusterAddonsConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAddonsConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterAddonsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpLoadBalancing, actual.HttpLoadBalancing, dcl.Info{ObjectFunction: compareClusterAddonsConfigHttpLoadBalancingNewStyle}, fn.AddNest("HttpLoadBalancing")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HorizontalPodAutoscaling, actual.HorizontalPodAutoscaling, dcl.Info{ObjectFunction: compareClusterAddonsConfigHorizontalPodAutoscalingNewStyle}, fn.AddNest("HorizontalPodAutoscaling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubernetesDashboard, actual.KubernetesDashboard, dcl.Info{ObjectFunction: compareClusterAddonsConfigKubernetesDashboardNewStyle}, fn.AddNest("KubernetesDashboard")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NetworkPolicyConfig, actual.NetworkPolicyConfig, dcl.Info{ObjectFunction: compareClusterAddonsConfigNetworkPolicyConfigNewStyle}, fn.AddNest("NetworkPolicyConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudRunConfig, actual.CloudRunConfig, dcl.Info{ObjectFunction: compareClusterAddonsConfigCloudRunConfigNewStyle}, fn.AddNest("CloudRunConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DnsCacheConfig, actual.DnsCacheConfig, dcl.Info{ObjectFunction: compareClusterAddonsConfigDnsCacheConfigNewStyle}, fn.AddNest("DnsCacheConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConfigConnectorConfig, actual.ConfigConnectorConfig, dcl.Info{ObjectFunction: compareClusterAddonsConfigConfigConnectorConfigNewStyle}, fn.AddNest("ConfigConnectorConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GcePersistentDiskCsiDriverConfig, actual.GcePersistentDiskCsiDriverConfig, dcl.Info{ObjectFunction: compareClusterAddonsConfigGcePersistentDiskCsiDriverConfigNewStyle}, fn.AddNest("GcePersistentDiskCsiDriverConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAddonsConfig(c *Client, desired, actual *ClusterAddonsConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareClusterAddonsConfigHttpLoadBalancing(c, desired.HttpLoadBalancing, actual.HttpLoadBalancing) && !dcl.IsZeroValue(desired.HttpLoadBalancing) {
		c.Config.Logger.Infof("Diff in HttpLoadBalancing.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpLoadBalancing), dcl.SprintResource(actual.HttpLoadBalancing))
		return true
	}
	if compareClusterAddonsConfigHorizontalPodAutoscaling(c, desired.HorizontalPodAutoscaling, actual.HorizontalPodAutoscaling) && !dcl.IsZeroValue(desired.HorizontalPodAutoscaling) {
		c.Config.Logger.Infof("Diff in HorizontalPodAutoscaling.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HorizontalPodAutoscaling), dcl.SprintResource(actual.HorizontalPodAutoscaling))
		return true
	}
	if compareClusterAddonsConfigKubernetesDashboard(c, desired.KubernetesDashboard, actual.KubernetesDashboard) && !dcl.IsZeroValue(desired.KubernetesDashboard) {
		c.Config.Logger.Infof("Diff in KubernetesDashboard.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KubernetesDashboard), dcl.SprintResource(actual.KubernetesDashboard))
		return true
	}
	if compareClusterAddonsConfigNetworkPolicyConfig(c, desired.NetworkPolicyConfig, actual.NetworkPolicyConfig) && !dcl.IsZeroValue(desired.NetworkPolicyConfig) {
		c.Config.Logger.Infof("Diff in NetworkPolicyConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NetworkPolicyConfig), dcl.SprintResource(actual.NetworkPolicyConfig))
		return true
	}
	if compareClusterAddonsConfigCloudRunConfig(c, desired.CloudRunConfig, actual.CloudRunConfig) && !dcl.IsZeroValue(desired.CloudRunConfig) {
		c.Config.Logger.Infof("Diff in CloudRunConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CloudRunConfig), dcl.SprintResource(actual.CloudRunConfig))
		return true
	}
	if compareClusterAddonsConfigDnsCacheConfig(c, desired.DnsCacheConfig, actual.DnsCacheConfig) && !dcl.IsZeroValue(desired.DnsCacheConfig) {
		c.Config.Logger.Infof("Diff in DnsCacheConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DnsCacheConfig), dcl.SprintResource(actual.DnsCacheConfig))
		return true
	}
	if compareClusterAddonsConfigConfigConnectorConfig(c, desired.ConfigConnectorConfig, actual.ConfigConnectorConfig) && !dcl.IsZeroValue(desired.ConfigConnectorConfig) {
		c.Config.Logger.Infof("Diff in ConfigConnectorConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConfigConnectorConfig), dcl.SprintResource(actual.ConfigConnectorConfig))
		return true
	}
	if compareClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, desired.GcePersistentDiskCsiDriverConfig, actual.GcePersistentDiskCsiDriverConfig) && !dcl.IsZeroValue(desired.GcePersistentDiskCsiDriverConfig) {
		c.Config.Logger.Infof("Diff in GcePersistentDiskCsiDriverConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GcePersistentDiskCsiDriverConfig), dcl.SprintResource(actual.GcePersistentDiskCsiDriverConfig))
		return true
	}
	return false
}

func compareClusterAddonsConfigSlice(c *Client, desired, actual []ClusterAddonsConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigMap(c *Client, desired, actual map[string]ClusterAddonsConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAddonsConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigHttpLoadBalancingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAddonsConfigHttpLoadBalancing)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAddonsConfigHttpLoadBalancing)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigHttpLoadBalancing or *ClusterAddonsConfigHttpLoadBalancing", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAddonsConfigHttpLoadBalancing)
	if !ok {
		actualNotPointer, ok := a.(ClusterAddonsConfigHttpLoadBalancing)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigHttpLoadBalancing", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAddonsConfigHttpLoadBalancing(c *Client, desired, actual *ClusterAddonsConfigHttpLoadBalancing) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) {
		c.Config.Logger.Infof("Diff in Disabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	return false
}

func compareClusterAddonsConfigHttpLoadBalancingSlice(c *Client, desired, actual []ClusterAddonsConfigHttpLoadBalancing) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigHttpLoadBalancing, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfigHttpLoadBalancing(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigHttpLoadBalancing, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigHttpLoadBalancingMap(c *Client, desired, actual map[string]ClusterAddonsConfigHttpLoadBalancing) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigHttpLoadBalancing, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigHttpLoadBalancing, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAddonsConfigHttpLoadBalancing(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigHttpLoadBalancing, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigHorizontalPodAutoscalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAddonsConfigHorizontalPodAutoscaling)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAddonsConfigHorizontalPodAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigHorizontalPodAutoscaling or *ClusterAddonsConfigHorizontalPodAutoscaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAddonsConfigHorizontalPodAutoscaling)
	if !ok {
		actualNotPointer, ok := a.(ClusterAddonsConfigHorizontalPodAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigHorizontalPodAutoscaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAddonsConfigHorizontalPodAutoscaling(c *Client, desired, actual *ClusterAddonsConfigHorizontalPodAutoscaling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) {
		c.Config.Logger.Infof("Diff in Disabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	return false
}

func compareClusterAddonsConfigHorizontalPodAutoscalingSlice(c *Client, desired, actual []ClusterAddonsConfigHorizontalPodAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigHorizontalPodAutoscaling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfigHorizontalPodAutoscaling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigHorizontalPodAutoscaling, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigHorizontalPodAutoscalingMap(c *Client, desired, actual map[string]ClusterAddonsConfigHorizontalPodAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigHorizontalPodAutoscaling, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigHorizontalPodAutoscaling, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAddonsConfigHorizontalPodAutoscaling(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigHorizontalPodAutoscaling, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigKubernetesDashboardNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAddonsConfigKubernetesDashboard)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAddonsConfigKubernetesDashboard)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigKubernetesDashboard or *ClusterAddonsConfigKubernetesDashboard", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAddonsConfigKubernetesDashboard)
	if !ok {
		actualNotPointer, ok := a.(ClusterAddonsConfigKubernetesDashboard)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigKubernetesDashboard", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAddonsConfigKubernetesDashboard(c *Client, desired, actual *ClusterAddonsConfigKubernetesDashboard) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) {
		c.Config.Logger.Infof("Diff in Disabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	return false
}

func compareClusterAddonsConfigKubernetesDashboardSlice(c *Client, desired, actual []ClusterAddonsConfigKubernetesDashboard) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigKubernetesDashboard, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfigKubernetesDashboard(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigKubernetesDashboard, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigKubernetesDashboardMap(c *Client, desired, actual map[string]ClusterAddonsConfigKubernetesDashboard) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigKubernetesDashboard, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigKubernetesDashboard, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAddonsConfigKubernetesDashboard(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigKubernetesDashboard, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigNetworkPolicyConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAddonsConfigNetworkPolicyConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAddonsConfigNetworkPolicyConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigNetworkPolicyConfig or *ClusterAddonsConfigNetworkPolicyConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAddonsConfigNetworkPolicyConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterAddonsConfigNetworkPolicyConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigNetworkPolicyConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAddonsConfigNetworkPolicyConfig(c *Client, desired, actual *ClusterAddonsConfigNetworkPolicyConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) {
		c.Config.Logger.Infof("Diff in Disabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	return false
}

func compareClusterAddonsConfigNetworkPolicyConfigSlice(c *Client, desired, actual []ClusterAddonsConfigNetworkPolicyConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigNetworkPolicyConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfigNetworkPolicyConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigNetworkPolicyConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigNetworkPolicyConfigMap(c *Client, desired, actual map[string]ClusterAddonsConfigNetworkPolicyConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigNetworkPolicyConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigNetworkPolicyConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAddonsConfigNetworkPolicyConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigNetworkPolicyConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigCloudRunConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAddonsConfigCloudRunConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAddonsConfigCloudRunConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigCloudRunConfig or *ClusterAddonsConfigCloudRunConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAddonsConfigCloudRunConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterAddonsConfigCloudRunConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigCloudRunConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LoadBalancerType, actual.LoadBalancerType, dcl.Info{Type: "EnumType"}, fn.AddNest("LoadBalancerType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAddonsConfigCloudRunConfig(c *Client, desired, actual *ClusterAddonsConfigCloudRunConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) {
		c.Config.Logger.Infof("Diff in Disabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	if !reflect.DeepEqual(desired.LoadBalancerType, actual.LoadBalancerType) && !dcl.IsZeroValue(desired.LoadBalancerType) {
		c.Config.Logger.Infof("Diff in LoadBalancerType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LoadBalancerType), dcl.SprintResource(actual.LoadBalancerType))
		return true
	}
	return false
}

func compareClusterAddonsConfigCloudRunConfigSlice(c *Client, desired, actual []ClusterAddonsConfigCloudRunConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigCloudRunConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfigCloudRunConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigCloudRunConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigCloudRunConfigMap(c *Client, desired, actual map[string]ClusterAddonsConfigCloudRunConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigCloudRunConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigCloudRunConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAddonsConfigCloudRunConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigCloudRunConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigDnsCacheConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAddonsConfigDnsCacheConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAddonsConfigDnsCacheConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigDnsCacheConfig or *ClusterAddonsConfigDnsCacheConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAddonsConfigDnsCacheConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterAddonsConfigDnsCacheConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigDnsCacheConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAddonsConfigDnsCacheConfig(c *Client, desired, actual *ClusterAddonsConfigDnsCacheConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterAddonsConfigDnsCacheConfigSlice(c *Client, desired, actual []ClusterAddonsConfigDnsCacheConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigDnsCacheConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfigDnsCacheConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigDnsCacheConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigDnsCacheConfigMap(c *Client, desired, actual map[string]ClusterAddonsConfigDnsCacheConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigDnsCacheConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigDnsCacheConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAddonsConfigDnsCacheConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigDnsCacheConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigConfigConnectorConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAddonsConfigConfigConnectorConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAddonsConfigConfigConnectorConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigConfigConnectorConfig or *ClusterAddonsConfigConfigConnectorConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAddonsConfigConfigConnectorConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterAddonsConfigConfigConnectorConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigConfigConnectorConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAddonsConfigConfigConnectorConfig(c *Client, desired, actual *ClusterAddonsConfigConfigConnectorConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterAddonsConfigConfigConnectorConfigSlice(c *Client, desired, actual []ClusterAddonsConfigConfigConnectorConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigConfigConnectorConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfigConfigConnectorConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigConfigConnectorConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigConfigConnectorConfigMap(c *Client, desired, actual map[string]ClusterAddonsConfigConfigConnectorConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigConfigConnectorConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigConfigConnectorConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAddonsConfigConfigConnectorConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigConfigConnectorConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigGcePersistentDiskCsiDriverConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAddonsConfigGcePersistentDiskCsiDriverConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAddonsConfigGcePersistentDiskCsiDriverConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigGcePersistentDiskCsiDriverConfig or *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAddonsConfigGcePersistentDiskCsiDriverConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterAddonsConfigGcePersistentDiskCsiDriverConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAddonsConfigGcePersistentDiskCsiDriverConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c *Client, desired, actual *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterAddonsConfigGcePersistentDiskCsiDriverConfigSlice(c *Client, desired, actual []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigGcePersistentDiskCsiDriverConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigGcePersistentDiskCsiDriverConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigGcePersistentDiskCsiDriverConfigMap(c *Client, desired, actual map[string]ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigGcePersistentDiskCsiDriverConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigGcePersistentDiskCsiDriverConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigGcePersistentDiskCsiDriverConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePools)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePools)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePools or *ClusterNodePools", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePools)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePools)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePools", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Config, actual.Config, dcl.Info{ObjectFunction: compareClusterNodePoolsConfigNewStyle}, fn.AddNest("Config")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InitialNodeCount, actual.InitialNodeCount, dcl.Info{}, fn.AddNest("InitialNodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Locations, actual.Locations, dcl.Info{}, fn.AddNest("Locations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceGroupUrls, actual.InstanceGroupUrls, dcl.Info{}, fn.AddNest("InstanceGroupUrls")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{Type: "EnumType"}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StatusMessage, actual.StatusMessage, dcl.Info{}, fn.AddNest("StatusMessage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Autoscaling, actual.Autoscaling, dcl.Info{ObjectFunction: compareClusterNodePoolsAutoscalingNewStyle}, fn.AddNest("Autoscaling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Management, actual.Management, dcl.Info{ObjectFunction: compareClusterNodePoolsManagementNewStyle}, fn.AddNest("Management")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxPodsConstraint, actual.MaxPodsConstraint, dcl.Info{ObjectFunction: compareClusterNodePoolsMaxPodsConstraintNewStyle}, fn.AddNest("MaxPodsConstraint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Conditions, actual.Conditions, dcl.Info{ObjectFunction: compareClusterNodePoolsConditionsNewStyle}, fn.AddNest("Conditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PodIPv4CidrSize, actual.PodIPv4CidrSize, dcl.Info{}, fn.AddNest("PodIPv4CidrSize")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpgradeSettings, actual.UpgradeSettings, dcl.Info{ObjectFunction: compareClusterNodePoolsUpgradeSettingsNewStyle}, fn.AddNest("UpgradeSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePools(c *Client, desired, actual *ClusterNodePools) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if compareClusterNodePoolsConfig(c, desired.Config, actual.Config) && !dcl.IsZeroValue(desired.Config) {
		c.Config.Logger.Infof("Diff in Config.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Config), dcl.SprintResource(actual.Config))
		return true
	}
	if !reflect.DeepEqual(desired.InitialNodeCount, actual.InitialNodeCount) && !dcl.IsZeroValue(desired.InitialNodeCount) {
		c.Config.Logger.Infof("Diff in InitialNodeCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InitialNodeCount), dcl.SprintResource(actual.InitialNodeCount))
		return true
	}
	if !dcl.StringSliceEquals(desired.Locations, actual.Locations) && !dcl.IsZeroValue(desired.Locations) {
		c.Config.Logger.Infof("Diff in Locations.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Locations), dcl.SprintResource(actual.Locations))
		return true
	}
	if !dcl.StringCanonicalize(desired.SelfLink, actual.SelfLink) && !dcl.IsZeroValue(desired.SelfLink) {
		c.Config.Logger.Infof("Diff in SelfLink.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SelfLink), dcl.SprintResource(actual.SelfLink))
		return true
	}
	if !dcl.StringCanonicalize(desired.Version, actual.Version) && !dcl.IsZeroValue(desired.Version) {
		c.Config.Logger.Infof("Diff in Version.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Version), dcl.SprintResource(actual.Version))
		return true
	}
	if !dcl.StringSliceEquals(desired.InstanceGroupUrls, actual.InstanceGroupUrls) && !dcl.IsZeroValue(desired.InstanceGroupUrls) {
		c.Config.Logger.Infof("Diff in InstanceGroupUrls.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InstanceGroupUrls), dcl.SprintResource(actual.InstanceGroupUrls))
		return true
	}
	if !reflect.DeepEqual(desired.Status, actual.Status) && !dcl.IsZeroValue(desired.Status) {
		c.Config.Logger.Infof("Diff in Status.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Status), dcl.SprintResource(actual.Status))
		return true
	}
	if !dcl.StringCanonicalize(desired.StatusMessage, actual.StatusMessage) && !dcl.IsZeroValue(desired.StatusMessage) {
		c.Config.Logger.Infof("Diff in StatusMessage.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StatusMessage), dcl.SprintResource(actual.StatusMessage))
		return true
	}
	if compareClusterNodePoolsAutoscaling(c, desired.Autoscaling, actual.Autoscaling) && !dcl.IsZeroValue(desired.Autoscaling) {
		c.Config.Logger.Infof("Diff in Autoscaling.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Autoscaling), dcl.SprintResource(actual.Autoscaling))
		return true
	}
	if compareClusterNodePoolsManagement(c, desired.Management, actual.Management) && !dcl.IsZeroValue(desired.Management) {
		c.Config.Logger.Infof("Diff in Management.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Management), dcl.SprintResource(actual.Management))
		return true
	}
	if compareClusterNodePoolsMaxPodsConstraint(c, desired.MaxPodsConstraint, actual.MaxPodsConstraint) && !dcl.IsZeroValue(desired.MaxPodsConstraint) {
		c.Config.Logger.Infof("Diff in MaxPodsConstraint.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxPodsConstraint), dcl.SprintResource(actual.MaxPodsConstraint))
		return true
	}
	if compareClusterNodePoolsConditionsSlice(c, desired.Conditions, actual.Conditions) && !dcl.IsZeroValue(desired.Conditions) {
		c.Config.Logger.Infof("Diff in Conditions.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Conditions), dcl.SprintResource(actual.Conditions))
		return true
	}
	if !reflect.DeepEqual(desired.PodIPv4CidrSize, actual.PodIPv4CidrSize) && !dcl.IsZeroValue(desired.PodIPv4CidrSize) {
		c.Config.Logger.Infof("Diff in PodIPv4CidrSize.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PodIPv4CidrSize), dcl.SprintResource(actual.PodIPv4CidrSize))
		return true
	}
	if compareClusterNodePoolsUpgradeSettings(c, desired.UpgradeSettings, actual.UpgradeSettings) && !dcl.IsZeroValue(desired.UpgradeSettings) {
		c.Config.Logger.Infof("Diff in UpgradeSettings.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpgradeSettings), dcl.SprintResource(actual.UpgradeSettings))
		return true
	}
	return false
}

func compareClusterNodePoolsSlice(c *Client, desired, actual []ClusterNodePools) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePools, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePools(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePools, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsMap(c *Client, desired, actual map[string]ClusterNodePools) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePools, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePools, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePools(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePools, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfig or *ClusterNodePoolsConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskSizeGb, actual.DiskSizeGb, dcl.Info{}, fn.AddNest("DiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuthScopes, actual.OAuthScopes, dcl.Info{}, fn.AddNest("OAuthScopes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.Info{}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImageType, actual.ImageType, dcl.Info{}, fn.AddNest("ImageType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalSsdCount, actual.LocalSsdCount, dcl.Info{}, fn.AddNest("LocalSsdCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tags, actual.Tags, dcl.Info{}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Preemptible, actual.Preemptible, dcl.Info{}, fn.AddNest("Preemptible")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Accelerators, actual.Accelerators, dcl.Info{ObjectFunction: compareClusterNodePoolsConfigAcceleratorsNewStyle}, fn.AddNest("Accelerators")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskType, actual.DiskType, dcl.Info{}, fn.AddNest("DiskType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinCpuPlatform, actual.MinCpuPlatform, dcl.Info{}, fn.AddNest("MinCpuPlatform")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkloadMetadataConfig, actual.WorkloadMetadataConfig, dcl.Info{ObjectFunction: compareClusterNodePoolsConfigWorkloadMetadataConfigNewStyle}, fn.AddNest("WorkloadMetadataConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Taints, actual.Taints, dcl.Info{ObjectFunction: compareClusterNodePoolsConfigTaintsNewStyle}, fn.AddNest("Taints")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SandboxConfig, actual.SandboxConfig, dcl.Info{ObjectFunction: compareClusterNodePoolsConfigSandboxConfigNewStyle}, fn.AddNest("SandboxConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeGroup, actual.NodeGroup, dcl.Info{}, fn.AddNest("NodeGroup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReservationAffinity, actual.ReservationAffinity, dcl.Info{ObjectFunction: compareClusterNodePoolsConfigReservationAffinityNewStyle}, fn.AddNest("ReservationAffinity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig, dcl.Info{ObjectFunction: compareClusterNodePoolsConfigShieldedInstanceConfigNewStyle}, fn.AddNest("ShieldedInstanceConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LinuxNodeConfig, actual.LinuxNodeConfig, dcl.Info{ObjectFunction: compareClusterNodePoolsConfigLinuxNodeConfigNewStyle}, fn.AddNest("LinuxNodeConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubeletConfig, actual.KubeletConfig, dcl.Info{ObjectFunction: compareClusterNodePoolsConfigKubeletConfigNewStyle}, fn.AddNest("KubeletConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BootDiskKmsKey, actual.BootDiskKmsKey, dcl.Info{}, fn.AddNest("BootDiskKmsKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConfig(c *Client, desired, actual *ClusterNodePoolsConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.MachineType, actual.MachineType) && !dcl.IsZeroValue(desired.MachineType) {
		c.Config.Logger.Infof("Diff in MachineType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MachineType), dcl.SprintResource(actual.MachineType))
		return true
	}
	if !reflect.DeepEqual(desired.DiskSizeGb, actual.DiskSizeGb) && !dcl.IsZeroValue(desired.DiskSizeGb) {
		c.Config.Logger.Infof("Diff in DiskSizeGb.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskSizeGb), dcl.SprintResource(actual.DiskSizeGb))
		return true
	}
	if !dcl.StringSliceEquals(desired.OAuthScopes, actual.OAuthScopes) && !dcl.IsZeroValue(desired.OAuthScopes) {
		c.Config.Logger.Infof("Diff in OAuthScopes.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OAuthScopes), dcl.SprintResource(actual.OAuthScopes))
		return true
	}
	if !dcl.StringCanonicalize(desired.ServiceAccount, actual.ServiceAccount) && !dcl.IsZeroValue(desired.ServiceAccount) {
		c.Config.Logger.Infof("Diff in ServiceAccount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccount), dcl.SprintResource(actual.ServiceAccount))
		return true
	}
	if !dcl.MapEquals(desired.Metadata, actual.Metadata, []string(nil)) && !dcl.IsZeroValue(desired.Metadata) {
		c.Config.Logger.Infof("Diff in Metadata.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Metadata), dcl.SprintResource(actual.Metadata))
		return true
	}
	if !dcl.StringCanonicalize(desired.ImageType, actual.ImageType) && !dcl.IsZeroValue(desired.ImageType) {
		c.Config.Logger.Infof("Diff in ImageType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ImageType), dcl.SprintResource(actual.ImageType))
		return true
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	if !reflect.DeepEqual(desired.LocalSsdCount, actual.LocalSsdCount) && !dcl.IsZeroValue(desired.LocalSsdCount) {
		c.Config.Logger.Infof("Diff in LocalSsdCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalSsdCount), dcl.SprintResource(actual.LocalSsdCount))
		return true
	}
	if !dcl.StringSliceEquals(desired.Tags, actual.Tags) && !dcl.IsZeroValue(desired.Tags) {
		c.Config.Logger.Infof("Diff in Tags.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tags), dcl.SprintResource(actual.Tags))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Preemptible, actual.Preemptible) && !dcl.IsZeroValue(desired.Preemptible) {
		c.Config.Logger.Infof("Diff in Preemptible.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Preemptible), dcl.SprintResource(actual.Preemptible))
		return true
	}
	if compareClusterNodePoolsConfigAcceleratorsSlice(c, desired.Accelerators, actual.Accelerators) && !dcl.IsZeroValue(desired.Accelerators) {
		c.Config.Logger.Infof("Diff in Accelerators.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Accelerators), dcl.SprintResource(actual.Accelerators))
		return true
	}
	if !dcl.StringCanonicalize(desired.DiskType, actual.DiskType) && !dcl.IsZeroValue(desired.DiskType) {
		c.Config.Logger.Infof("Diff in DiskType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskType), dcl.SprintResource(actual.DiskType))
		return true
	}
	if !dcl.StringCanonicalize(desired.MinCpuPlatform, actual.MinCpuPlatform) && !dcl.IsZeroValue(desired.MinCpuPlatform) {
		c.Config.Logger.Infof("Diff in MinCpuPlatform.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinCpuPlatform), dcl.SprintResource(actual.MinCpuPlatform))
		return true
	}
	if compareClusterNodePoolsConfigWorkloadMetadataConfig(c, desired.WorkloadMetadataConfig, actual.WorkloadMetadataConfig) && !dcl.IsZeroValue(desired.WorkloadMetadataConfig) {
		c.Config.Logger.Infof("Diff in WorkloadMetadataConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WorkloadMetadataConfig), dcl.SprintResource(actual.WorkloadMetadataConfig))
		return true
	}
	if compareClusterNodePoolsConfigTaintsSlice(c, desired.Taints, actual.Taints) && !dcl.IsZeroValue(desired.Taints) {
		c.Config.Logger.Infof("Diff in Taints.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Taints), dcl.SprintResource(actual.Taints))
		return true
	}
	if compareClusterNodePoolsConfigSandboxConfig(c, desired.SandboxConfig, actual.SandboxConfig) && !dcl.IsZeroValue(desired.SandboxConfig) {
		c.Config.Logger.Infof("Diff in SandboxConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SandboxConfig), dcl.SprintResource(actual.SandboxConfig))
		return true
	}
	if !dcl.StringCanonicalize(desired.NodeGroup, actual.NodeGroup) && !dcl.IsZeroValue(desired.NodeGroup) {
		c.Config.Logger.Infof("Diff in NodeGroup.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NodeGroup), dcl.SprintResource(actual.NodeGroup))
		return true
	}
	if compareClusterNodePoolsConfigReservationAffinity(c, desired.ReservationAffinity, actual.ReservationAffinity) && !dcl.IsZeroValue(desired.ReservationAffinity) {
		c.Config.Logger.Infof("Diff in ReservationAffinity.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReservationAffinity), dcl.SprintResource(actual.ReservationAffinity))
		return true
	}
	if compareClusterNodePoolsConfigShieldedInstanceConfig(c, desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig) && !dcl.IsZeroValue(desired.ShieldedInstanceConfig) {
		c.Config.Logger.Infof("Diff in ShieldedInstanceConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ShieldedInstanceConfig), dcl.SprintResource(actual.ShieldedInstanceConfig))
		return true
	}
	if compareClusterNodePoolsConfigLinuxNodeConfig(c, desired.LinuxNodeConfig, actual.LinuxNodeConfig) && !dcl.IsZeroValue(desired.LinuxNodeConfig) {
		c.Config.Logger.Infof("Diff in LinuxNodeConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LinuxNodeConfig), dcl.SprintResource(actual.LinuxNodeConfig))
		return true
	}
	if compareClusterNodePoolsConfigKubeletConfig(c, desired.KubeletConfig, actual.KubeletConfig) && !dcl.IsZeroValue(desired.KubeletConfig) {
		c.Config.Logger.Infof("Diff in KubeletConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KubeletConfig), dcl.SprintResource(actual.KubeletConfig))
		return true
	}
	if !dcl.StringCanonicalize(desired.BootDiskKmsKey, actual.BootDiskKmsKey) && !dcl.IsZeroValue(desired.BootDiskKmsKey) {
		c.Config.Logger.Infof("Diff in BootDiskKmsKey.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BootDiskKmsKey), dcl.SprintResource(actual.BootDiskKmsKey))
		return true
	}
	return false
}

func compareClusterNodePoolsConfigSlice(c *Client, desired, actual []ClusterNodePoolsConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigMap(c *Client, desired, actual map[string]ClusterNodePoolsConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigAcceleratorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConfigAccelerators)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigAccelerators or *ClusterNodePoolsConfigAccelerators", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConfigAccelerators)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigAccelerators", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AcceleratorCount, actual.AcceleratorCount, dcl.Info{}, fn.AddNest("AcceleratorCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorType, actual.AcceleratorType, dcl.Info{}, fn.AddNest("AcceleratorType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConfigAccelerators(c *Client, desired, actual *ClusterNodePoolsConfigAccelerators) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.AcceleratorCount, actual.AcceleratorCount) && !dcl.IsZeroValue(desired.AcceleratorCount) {
		c.Config.Logger.Infof("Diff in AcceleratorCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorCount), dcl.SprintResource(actual.AcceleratorCount))
		return true
	}
	if !dcl.StringCanonicalize(desired.AcceleratorType, actual.AcceleratorType) && !dcl.IsZeroValue(desired.AcceleratorType) {
		c.Config.Logger.Infof("Diff in AcceleratorType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorType), dcl.SprintResource(actual.AcceleratorType))
		return true
	}
	return false
}

func compareClusterNodePoolsConfigAcceleratorsSlice(c *Client, desired, actual []ClusterNodePoolsConfigAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigAccelerators, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigAccelerators(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigAccelerators, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigAcceleratorsMap(c *Client, desired, actual map[string]ClusterNodePoolsConfigAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigAccelerators, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigAccelerators, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConfigAccelerators(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigAccelerators, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigWorkloadMetadataConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConfigWorkloadMetadataConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConfigWorkloadMetadataConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigWorkloadMetadataConfig or *ClusterNodePoolsConfigWorkloadMetadataConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConfigWorkloadMetadataConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConfigWorkloadMetadataConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigWorkloadMetadataConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{Type: "EnumType"}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConfigWorkloadMetadataConfig(c *Client, desired, actual *ClusterNodePoolsConfigWorkloadMetadataConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Mode, actual.Mode) && !dcl.IsZeroValue(desired.Mode) {
		c.Config.Logger.Infof("Diff in Mode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Mode), dcl.SprintResource(actual.Mode))
		return true
	}
	return false
}

func compareClusterNodePoolsConfigWorkloadMetadataConfigSlice(c *Client, desired, actual []ClusterNodePoolsConfigWorkloadMetadataConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigWorkloadMetadataConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigWorkloadMetadataConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigWorkloadMetadataConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigWorkloadMetadataConfigMap(c *Client, desired, actual map[string]ClusterNodePoolsConfigWorkloadMetadataConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigWorkloadMetadataConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigWorkloadMetadataConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConfigWorkloadMetadataConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigWorkloadMetadataConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigTaintsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConfigTaints)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConfigTaints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigTaints or *ClusterNodePoolsConfigTaints", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConfigTaints)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConfigTaints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigTaints", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Effect, actual.Effect, dcl.Info{Type: "EnumType"}, fn.AddNest("Effect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConfigTaints(c *Client, desired, actual *ClusterNodePoolsConfigTaints) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Effect, actual.Effect) && !dcl.IsZeroValue(desired.Effect) {
		c.Config.Logger.Infof("Diff in Effect.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Effect), dcl.SprintResource(actual.Effect))
		return true
	}
	return false
}

func compareClusterNodePoolsConfigTaintsSlice(c *Client, desired, actual []ClusterNodePoolsConfigTaints) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigTaints, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigTaints(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigTaints, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigTaintsMap(c *Client, desired, actual map[string]ClusterNodePoolsConfigTaints) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigTaints, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigTaints, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConfigTaints(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigTaints, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigSandboxConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConfigSandboxConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConfigSandboxConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigSandboxConfig or *ClusterNodePoolsConfigSandboxConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConfigSandboxConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConfigSandboxConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigSandboxConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType"}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConfigSandboxConfig(c *Client, desired, actual *ClusterNodePoolsConfigSandboxConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}

func compareClusterNodePoolsConfigSandboxConfigSlice(c *Client, desired, actual []ClusterNodePoolsConfigSandboxConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigSandboxConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigSandboxConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigSandboxConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigSandboxConfigMap(c *Client, desired, actual map[string]ClusterNodePoolsConfigSandboxConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigSandboxConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigSandboxConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConfigSandboxConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigSandboxConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigReservationAffinityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConfigReservationAffinity)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConfigReservationAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigReservationAffinity or *ClusterNodePoolsConfigReservationAffinity", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConfigReservationAffinity)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConfigReservationAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigReservationAffinity", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConsumeReservationType, actual.ConsumeReservationType, dcl.Info{Type: "EnumType"}, fn.AddNest("ConsumeReservationType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Values, actual.Values, dcl.Info{}, fn.AddNest("Values")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConfigReservationAffinity(c *Client, desired, actual *ClusterNodePoolsConfigReservationAffinity) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.ConsumeReservationType, actual.ConsumeReservationType) && !dcl.IsZeroValue(desired.ConsumeReservationType) {
		c.Config.Logger.Infof("Diff in ConsumeReservationType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConsumeReservationType), dcl.SprintResource(actual.ConsumeReservationType))
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if !dcl.StringSliceEquals(desired.Values, actual.Values) && !dcl.IsZeroValue(desired.Values) {
		c.Config.Logger.Infof("Diff in Values.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Values), dcl.SprintResource(actual.Values))
		return true
	}
	return false
}

func compareClusterNodePoolsConfigReservationAffinitySlice(c *Client, desired, actual []ClusterNodePoolsConfigReservationAffinity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigReservationAffinity, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigReservationAffinity(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigReservationAffinity, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigReservationAffinityMap(c *Client, desired, actual map[string]ClusterNodePoolsConfigReservationAffinity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigReservationAffinity, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigReservationAffinity, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConfigReservationAffinity(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigReservationAffinity, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigShieldedInstanceConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConfigShieldedInstanceConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConfigShieldedInstanceConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigShieldedInstanceConfig or *ClusterNodePoolsConfigShieldedInstanceConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConfigShieldedInstanceConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConfigShieldedInstanceConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigShieldedInstanceConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnableSecureBoot, actual.EnableSecureBoot, dcl.Info{}, fn.AddNest("EnableSecureBoot")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring, dcl.Info{}, fn.AddNest("EnableIntegrityMonitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConfigShieldedInstanceConfig(c *Client, desired, actual *ClusterNodePoolsConfigShieldedInstanceConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableSecureBoot, actual.EnableSecureBoot) && !dcl.IsZeroValue(desired.EnableSecureBoot) {
		c.Config.Logger.Infof("Diff in EnableSecureBoot.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableSecureBoot), dcl.SprintResource(actual.EnableSecureBoot))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring) && !dcl.IsZeroValue(desired.EnableIntegrityMonitoring) {
		c.Config.Logger.Infof("Diff in EnableIntegrityMonitoring.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableIntegrityMonitoring), dcl.SprintResource(actual.EnableIntegrityMonitoring))
		return true
	}
	return false
}

func compareClusterNodePoolsConfigShieldedInstanceConfigSlice(c *Client, desired, actual []ClusterNodePoolsConfigShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigShieldedInstanceConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigShieldedInstanceConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigShieldedInstanceConfigMap(c *Client, desired, actual map[string]ClusterNodePoolsConfigShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigShieldedInstanceConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConfigShieldedInstanceConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigShieldedInstanceConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigLinuxNodeConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConfigLinuxNodeConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConfigLinuxNodeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigLinuxNodeConfig or *ClusterNodePoolsConfigLinuxNodeConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConfigLinuxNodeConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConfigLinuxNodeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigLinuxNodeConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Sysctls, actual.Sysctls, dcl.Info{}, fn.AddNest("Sysctls")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConfigLinuxNodeConfig(c *Client, desired, actual *ClusterNodePoolsConfigLinuxNodeConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.MapEquals(desired.Sysctls, actual.Sysctls, []string(nil)) && !dcl.IsZeroValue(desired.Sysctls) {
		c.Config.Logger.Infof("Diff in Sysctls.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sysctls), dcl.SprintResource(actual.Sysctls))
		return true
	}
	return false
}

func compareClusterNodePoolsConfigLinuxNodeConfigSlice(c *Client, desired, actual []ClusterNodePoolsConfigLinuxNodeConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigLinuxNodeConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigLinuxNodeConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigLinuxNodeConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigLinuxNodeConfigMap(c *Client, desired, actual map[string]ClusterNodePoolsConfigLinuxNodeConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigLinuxNodeConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigLinuxNodeConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConfigLinuxNodeConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigLinuxNodeConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigKubeletConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConfigKubeletConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConfigKubeletConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigKubeletConfig or *ClusterNodePoolsConfigKubeletConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConfigKubeletConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConfigKubeletConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConfigKubeletConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CpuManagerPolicy, actual.CpuManagerPolicy, dcl.Info{}, fn.AddNest("CpuManagerPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CpuCfsQuota, actual.CpuCfsQuota, dcl.Info{}, fn.AddNest("CpuCfsQuota")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CpuCfsQuotaPeriod, actual.CpuCfsQuotaPeriod, dcl.Info{}, fn.AddNest("CpuCfsQuotaPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConfigKubeletConfig(c *Client, desired, actual *ClusterNodePoolsConfigKubeletConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.CpuManagerPolicy, actual.CpuManagerPolicy) && !dcl.IsZeroValue(desired.CpuManagerPolicy) {
		c.Config.Logger.Infof("Diff in CpuManagerPolicy.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CpuManagerPolicy), dcl.SprintResource(actual.CpuManagerPolicy))
		return true
	}
	if !dcl.BoolCanonicalize(desired.CpuCfsQuota, actual.CpuCfsQuota) && !dcl.IsZeroValue(desired.CpuCfsQuota) {
		c.Config.Logger.Infof("Diff in CpuCfsQuota.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CpuCfsQuota), dcl.SprintResource(actual.CpuCfsQuota))
		return true
	}
	if !dcl.StringCanonicalize(desired.CpuCfsQuotaPeriod, actual.CpuCfsQuotaPeriod) && !dcl.IsZeroValue(desired.CpuCfsQuotaPeriod) {
		c.Config.Logger.Infof("Diff in CpuCfsQuotaPeriod.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CpuCfsQuotaPeriod), dcl.SprintResource(actual.CpuCfsQuotaPeriod))
		return true
	}
	return false
}

func compareClusterNodePoolsConfigKubeletConfigSlice(c *Client, desired, actual []ClusterNodePoolsConfigKubeletConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigKubeletConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigKubeletConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigKubeletConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigKubeletConfigMap(c *Client, desired, actual map[string]ClusterNodePoolsConfigKubeletConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigKubeletConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigKubeletConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConfigKubeletConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigKubeletConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsAutoscalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsAutoscaling)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsAutoscaling or *ClusterNodePoolsAutoscaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsAutoscaling)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsAutoscaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinNodeCount, actual.MinNodeCount, dcl.Info{}, fn.AddNest("MinNodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxNodeCount, actual.MaxNodeCount, dcl.Info{}, fn.AddNest("MaxNodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Autoprovisioned, actual.Autoprovisioned, dcl.Info{}, fn.AddNest("Autoprovisioned")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsAutoscaling(c *Client, desired, actual *ClusterNodePoolsAutoscaling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.MinNodeCount, actual.MinNodeCount) && !dcl.IsZeroValue(desired.MinNodeCount) {
		c.Config.Logger.Infof("Diff in MinNodeCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinNodeCount), dcl.SprintResource(actual.MinNodeCount))
		return true
	}
	if !reflect.DeepEqual(desired.MaxNodeCount, actual.MaxNodeCount) && !dcl.IsZeroValue(desired.MaxNodeCount) {
		c.Config.Logger.Infof("Diff in MaxNodeCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxNodeCount), dcl.SprintResource(actual.MaxNodeCount))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Autoprovisioned, actual.Autoprovisioned) && !dcl.IsZeroValue(desired.Autoprovisioned) {
		c.Config.Logger.Infof("Diff in Autoprovisioned.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Autoprovisioned), dcl.SprintResource(actual.Autoprovisioned))
		return true
	}
	return false
}

func compareClusterNodePoolsAutoscalingSlice(c *Client, desired, actual []ClusterNodePoolsAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsAutoscaling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsAutoscaling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsAutoscaling, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsAutoscalingMap(c *Client, desired, actual map[string]ClusterNodePoolsAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsAutoscaling, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsAutoscaling, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsAutoscaling(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsAutoscaling, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsManagementNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsManagement)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsManagement)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsManagement or *ClusterNodePoolsManagement", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsManagement)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsManagement)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsManagement", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AutoUpgrade, actual.AutoUpgrade, dcl.Info{}, fn.AddNest("AutoUpgrade")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutoRepair, actual.AutoRepair, dcl.Info{}, fn.AddNest("AutoRepair")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpgradeOptions, actual.UpgradeOptions, dcl.Info{ObjectFunction: compareClusterNodePoolsManagementUpgradeOptionsNewStyle}, fn.AddNest("UpgradeOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsManagement(c *Client, desired, actual *ClusterNodePoolsManagement) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.AutoUpgrade, actual.AutoUpgrade) && !dcl.IsZeroValue(desired.AutoUpgrade) {
		c.Config.Logger.Infof("Diff in AutoUpgrade.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoUpgrade), dcl.SprintResource(actual.AutoUpgrade))
		return true
	}
	if !dcl.BoolCanonicalize(desired.AutoRepair, actual.AutoRepair) && !dcl.IsZeroValue(desired.AutoRepair) {
		c.Config.Logger.Infof("Diff in AutoRepair.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoRepair), dcl.SprintResource(actual.AutoRepair))
		return true
	}
	if compareClusterNodePoolsManagementUpgradeOptions(c, desired.UpgradeOptions, actual.UpgradeOptions) && !dcl.IsZeroValue(desired.UpgradeOptions) {
		c.Config.Logger.Infof("Diff in UpgradeOptions.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpgradeOptions), dcl.SprintResource(actual.UpgradeOptions))
		return true
	}
	return false
}

func compareClusterNodePoolsManagementSlice(c *Client, desired, actual []ClusterNodePoolsManagement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsManagement, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsManagement(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsManagement, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsManagementMap(c *Client, desired, actual map[string]ClusterNodePoolsManagement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsManagement, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsManagement, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsManagement(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsManagement, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsManagementUpgradeOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsManagementUpgradeOptions)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsManagementUpgradeOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsManagementUpgradeOptions or *ClusterNodePoolsManagementUpgradeOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsManagementUpgradeOptions)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsManagementUpgradeOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsManagementUpgradeOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AutoUpgradeStartTime, actual.AutoUpgradeStartTime, dcl.Info{}, fn.AddNest("AutoUpgradeStartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsManagementUpgradeOptions(c *Client, desired, actual *ClusterNodePoolsManagementUpgradeOptions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.AutoUpgradeStartTime, actual.AutoUpgradeStartTime) && !dcl.IsZeroValue(desired.AutoUpgradeStartTime) {
		c.Config.Logger.Infof("Diff in AutoUpgradeStartTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoUpgradeStartTime), dcl.SprintResource(actual.AutoUpgradeStartTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) {
		c.Config.Logger.Infof("Diff in Description.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	return false
}

func compareClusterNodePoolsManagementUpgradeOptionsSlice(c *Client, desired, actual []ClusterNodePoolsManagementUpgradeOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsManagementUpgradeOptions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsManagementUpgradeOptions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsManagementUpgradeOptions, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsManagementUpgradeOptionsMap(c *Client, desired, actual map[string]ClusterNodePoolsManagementUpgradeOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsManagementUpgradeOptions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsManagementUpgradeOptions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsManagementUpgradeOptions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsManagementUpgradeOptions, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsMaxPodsConstraintNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsMaxPodsConstraint)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsMaxPodsConstraint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsMaxPodsConstraint or *ClusterNodePoolsMaxPodsConstraint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsMaxPodsConstraint)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsMaxPodsConstraint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsMaxPodsConstraint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxPodsPerNode, actual.MaxPodsPerNode, dcl.Info{}, fn.AddNest("MaxPodsPerNode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsMaxPodsConstraint(c *Client, desired, actual *ClusterNodePoolsMaxPodsConstraint) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.MaxPodsPerNode, actual.MaxPodsPerNode) && !dcl.IsZeroValue(desired.MaxPodsPerNode) {
		c.Config.Logger.Infof("Diff in MaxPodsPerNode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxPodsPerNode), dcl.SprintResource(actual.MaxPodsPerNode))
		return true
	}
	return false
}

func compareClusterNodePoolsMaxPodsConstraintSlice(c *Client, desired, actual []ClusterNodePoolsMaxPodsConstraint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsMaxPodsConstraint, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsMaxPodsConstraint(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsMaxPodsConstraint, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsMaxPodsConstraintMap(c *Client, desired, actual map[string]ClusterNodePoolsMaxPodsConstraint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsMaxPodsConstraint, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsMaxPodsConstraint, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsMaxPodsConstraint(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsMaxPodsConstraint, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConditionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsConditions)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConditions or *ClusterNodePoolsConditions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsConditions)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsConditions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{Type: "EnumType"}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CanonicalCode, actual.CanonicalCode, dcl.Info{Type: "EnumType"}, fn.AddNest("CanonicalCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsConditions(c *Client, desired, actual *ClusterNodePoolsConditions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Code, actual.Code) && !dcl.IsZeroValue(desired.Code) {
		c.Config.Logger.Infof("Diff in Code.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Code), dcl.SprintResource(actual.Code))
		return true
	}
	if !dcl.StringCanonicalize(desired.Message, actual.Message) && !dcl.IsZeroValue(desired.Message) {
		c.Config.Logger.Infof("Diff in Message.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Message), dcl.SprintResource(actual.Message))
		return true
	}
	if !reflect.DeepEqual(desired.CanonicalCode, actual.CanonicalCode) && !dcl.IsZeroValue(desired.CanonicalCode) {
		c.Config.Logger.Infof("Diff in CanonicalCode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CanonicalCode), dcl.SprintResource(actual.CanonicalCode))
		return true
	}
	return false
}

func compareClusterNodePoolsConditionsSlice(c *Client, desired, actual []ClusterNodePoolsConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConditions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConditions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConditions, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConditionsMap(c *Client, desired, actual map[string]ClusterNodePoolsConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConditions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConditions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsConditions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConditions, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsUpgradeSettingsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodePoolsUpgradeSettings)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodePoolsUpgradeSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsUpgradeSettings or *ClusterNodePoolsUpgradeSettings", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodePoolsUpgradeSettings)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodePoolsUpgradeSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodePoolsUpgradeSettings", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxSurge, actual.MaxSurge, dcl.Info{}, fn.AddNest("MaxSurge")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxUnavailable, actual.MaxUnavailable, dcl.Info{}, fn.AddNest("MaxUnavailable")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodePoolsUpgradeSettings(c *Client, desired, actual *ClusterNodePoolsUpgradeSettings) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.MaxSurge, actual.MaxSurge) && !dcl.IsZeroValue(desired.MaxSurge) {
		c.Config.Logger.Infof("Diff in MaxSurge.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxSurge), dcl.SprintResource(actual.MaxSurge))
		return true
	}
	if !reflect.DeepEqual(desired.MaxUnavailable, actual.MaxUnavailable) && !dcl.IsZeroValue(desired.MaxUnavailable) {
		c.Config.Logger.Infof("Diff in MaxUnavailable.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxUnavailable), dcl.SprintResource(actual.MaxUnavailable))
		return true
	}
	return false
}

func compareClusterNodePoolsUpgradeSettingsSlice(c *Client, desired, actual []ClusterNodePoolsUpgradeSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsUpgradeSettings, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsUpgradeSettings(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsUpgradeSettings, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsUpgradeSettingsMap(c *Client, desired, actual map[string]ClusterNodePoolsUpgradeSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsUpgradeSettings, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsUpgradeSettings, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodePoolsUpgradeSettings(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsUpgradeSettings, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterLegacyAbacNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterLegacyAbac)
	if !ok {
		desiredNotPointer, ok := d.(ClusterLegacyAbac)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterLegacyAbac or *ClusterLegacyAbac", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterLegacyAbac)
	if !ok {
		actualNotPointer, ok := a.(ClusterLegacyAbac)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterLegacyAbac", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterLegacyAbac(c *Client, desired, actual *ClusterLegacyAbac) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterLegacyAbacSlice(c *Client, desired, actual []ClusterLegacyAbac) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterLegacyAbac, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterLegacyAbac(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterLegacyAbac, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterLegacyAbacMap(c *Client, desired, actual map[string]ClusterLegacyAbac) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterLegacyAbac, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterLegacyAbac, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterLegacyAbac(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterLegacyAbac, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNetworkPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNetworkPolicy)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNetworkPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNetworkPolicy or *ClusterNetworkPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNetworkPolicy)
	if !ok {
		actualNotPointer, ok := a.(ClusterNetworkPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNetworkPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Provider, actual.Provider, dcl.Info{Type: "EnumType"}, fn.AddNest("Provider")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNetworkPolicy(c *Client, desired, actual *ClusterNetworkPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Provider, actual.Provider) && !dcl.IsZeroValue(desired.Provider) {
		c.Config.Logger.Infof("Diff in Provider.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Provider), dcl.SprintResource(actual.Provider))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterNetworkPolicySlice(c *Client, desired, actual []ClusterNetworkPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNetworkPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNetworkPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNetworkPolicy, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNetworkPolicyMap(c *Client, desired, actual map[string]ClusterNetworkPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNetworkPolicy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNetworkPolicy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNetworkPolicy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNetworkPolicy, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterIPAllocationPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterIPAllocationPolicy)
	if !ok {
		desiredNotPointer, ok := d.(ClusterIPAllocationPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterIPAllocationPolicy or *ClusterIPAllocationPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterIPAllocationPolicy)
	if !ok {
		actualNotPointer, ok := a.(ClusterIPAllocationPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterIPAllocationPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UseIPAliases, actual.UseIPAliases, dcl.Info{}, fn.AddNest("UseIPAliases")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateSubnetwork, actual.CreateSubnetwork, dcl.Info{}, fn.AddNest("CreateSubnetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubnetworkName, actual.SubnetworkName, dcl.Info{}, fn.AddNest("SubnetworkName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterSecondaryRangeName, actual.ClusterSecondaryRangeName, dcl.Info{}, fn.AddNest("ClusterSecondaryRangeName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesSecondaryRangeName, actual.ServicesSecondaryRangeName, dcl.Info{}, fn.AddNest("ServicesSecondaryRangeName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterIPv4CidrBlock, actual.ClusterIPv4CidrBlock, dcl.Info{}, fn.AddNest("ClusterIPv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeIPv4CidrBlock, actual.NodeIPv4CidrBlock, dcl.Info{}, fn.AddNest("NodeIPv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesIPv4CidrBlock, actual.ServicesIPv4CidrBlock, dcl.Info{}, fn.AddNest("ServicesIPv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TPUIPv4CidrBlock, actual.TPUIPv4CidrBlock, dcl.Info{}, fn.AddNest("TPUIPv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterIPv4Cidr, actual.ClusterIPv4Cidr, dcl.Info{}, fn.AddNest("ClusterIPv4Cidr")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeIPv4Cidr, actual.NodeIPv4Cidr, dcl.Info{}, fn.AddNest("NodeIPv4Cidr")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesIPv4Cidr, actual.ServicesIPv4Cidr, dcl.Info{}, fn.AddNest("ServicesIPv4Cidr")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UseRoutes, actual.UseRoutes, dcl.Info{}, fn.AddNest("UseRoutes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterIPAllocationPolicy(c *Client, desired, actual *ClusterIPAllocationPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.UseIPAliases, actual.UseIPAliases) && !dcl.IsZeroValue(desired.UseIPAliases) {
		c.Config.Logger.Infof("Diff in UseIPAliases.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UseIPAliases), dcl.SprintResource(actual.UseIPAliases))
		return true
	}
	if !dcl.BoolCanonicalize(desired.CreateSubnetwork, actual.CreateSubnetwork) && !dcl.IsZeroValue(desired.CreateSubnetwork) {
		c.Config.Logger.Infof("Diff in CreateSubnetwork.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CreateSubnetwork), dcl.SprintResource(actual.CreateSubnetwork))
		return true
	}
	if !dcl.StringCanonicalize(desired.SubnetworkName, actual.SubnetworkName) && !dcl.IsZeroValue(desired.SubnetworkName) {
		c.Config.Logger.Infof("Diff in SubnetworkName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SubnetworkName), dcl.SprintResource(actual.SubnetworkName))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClusterSecondaryRangeName, actual.ClusterSecondaryRangeName) && !dcl.IsZeroValue(desired.ClusterSecondaryRangeName) {
		c.Config.Logger.Infof("Diff in ClusterSecondaryRangeName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterSecondaryRangeName), dcl.SprintResource(actual.ClusterSecondaryRangeName))
		return true
	}
	if !dcl.StringCanonicalize(desired.ServicesSecondaryRangeName, actual.ServicesSecondaryRangeName) && !dcl.IsZeroValue(desired.ServicesSecondaryRangeName) {
		c.Config.Logger.Infof("Diff in ServicesSecondaryRangeName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServicesSecondaryRangeName), dcl.SprintResource(actual.ServicesSecondaryRangeName))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClusterIPv4CidrBlock, actual.ClusterIPv4CidrBlock) && !dcl.IsZeroValue(desired.ClusterIPv4CidrBlock) {
		c.Config.Logger.Infof("Diff in ClusterIPv4CidrBlock.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterIPv4CidrBlock), dcl.SprintResource(actual.ClusterIPv4CidrBlock))
		return true
	}
	if !dcl.StringCanonicalize(desired.NodeIPv4CidrBlock, actual.NodeIPv4CidrBlock) && !dcl.IsZeroValue(desired.NodeIPv4CidrBlock) {
		c.Config.Logger.Infof("Diff in NodeIPv4CidrBlock.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NodeIPv4CidrBlock), dcl.SprintResource(actual.NodeIPv4CidrBlock))
		return true
	}
	if !dcl.StringCanonicalize(desired.ServicesIPv4CidrBlock, actual.ServicesIPv4CidrBlock) && !dcl.IsZeroValue(desired.ServicesIPv4CidrBlock) {
		c.Config.Logger.Infof("Diff in ServicesIPv4CidrBlock.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServicesIPv4CidrBlock), dcl.SprintResource(actual.ServicesIPv4CidrBlock))
		return true
	}
	if !dcl.StringCanonicalize(desired.TPUIPv4CidrBlock, actual.TPUIPv4CidrBlock) && !dcl.IsZeroValue(desired.TPUIPv4CidrBlock) {
		c.Config.Logger.Infof("Diff in TPUIPv4CidrBlock.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TPUIPv4CidrBlock), dcl.SprintResource(actual.TPUIPv4CidrBlock))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClusterIPv4Cidr, actual.ClusterIPv4Cidr) && !dcl.IsZeroValue(desired.ClusterIPv4Cidr) {
		c.Config.Logger.Infof("Diff in ClusterIPv4Cidr.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterIPv4Cidr), dcl.SprintResource(actual.ClusterIPv4Cidr))
		return true
	}
	if !dcl.StringCanonicalize(desired.NodeIPv4Cidr, actual.NodeIPv4Cidr) && !dcl.IsZeroValue(desired.NodeIPv4Cidr) {
		c.Config.Logger.Infof("Diff in NodeIPv4Cidr.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NodeIPv4Cidr), dcl.SprintResource(actual.NodeIPv4Cidr))
		return true
	}
	if !dcl.StringCanonicalize(desired.ServicesIPv4Cidr, actual.ServicesIPv4Cidr) && !dcl.IsZeroValue(desired.ServicesIPv4Cidr) {
		c.Config.Logger.Infof("Diff in ServicesIPv4Cidr.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServicesIPv4Cidr), dcl.SprintResource(actual.ServicesIPv4Cidr))
		return true
	}
	if !dcl.BoolCanonicalize(desired.UseRoutes, actual.UseRoutes) && !dcl.IsZeroValue(desired.UseRoutes) {
		c.Config.Logger.Infof("Diff in UseRoutes.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UseRoutes), dcl.SprintResource(actual.UseRoutes))
		return true
	}
	return false
}

func compareClusterIPAllocationPolicySlice(c *Client, desired, actual []ClusterIPAllocationPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterIPAllocationPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterIPAllocationPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterIPAllocationPolicy, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterIPAllocationPolicyMap(c *Client, desired, actual map[string]ClusterIPAllocationPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterIPAllocationPolicy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterIPAllocationPolicy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterIPAllocationPolicy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterIPAllocationPolicy, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthorizedNetworksConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMasterAuthorizedNetworksConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMasterAuthorizedNetworksConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMasterAuthorizedNetworksConfig or *ClusterMasterAuthorizedNetworksConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMasterAuthorizedNetworksConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterMasterAuthorizedNetworksConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMasterAuthorizedNetworksConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CidrBlocks, actual.CidrBlocks, dcl.Info{ObjectFunction: compareClusterMasterAuthorizedNetworksConfigCidrBlocksNewStyle}, fn.AddNest("CidrBlocks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMasterAuthorizedNetworksConfig(c *Client, desired, actual *ClusterMasterAuthorizedNetworksConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	if compareClusterMasterAuthorizedNetworksConfigCidrBlocksSlice(c, desired.CidrBlocks, actual.CidrBlocks) && !dcl.IsZeroValue(desired.CidrBlocks) {
		c.Config.Logger.Infof("Diff in CidrBlocks.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CidrBlocks), dcl.SprintResource(actual.CidrBlocks))
		return true
	}
	return false
}

func compareClusterMasterAuthorizedNetworksConfigSlice(c *Client, desired, actual []ClusterMasterAuthorizedNetworksConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMasterAuthorizedNetworksConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMasterAuthorizedNetworksConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMasterAuthorizedNetworksConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthorizedNetworksConfigMap(c *Client, desired, actual map[string]ClusterMasterAuthorizedNetworksConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMasterAuthorizedNetworksConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterMasterAuthorizedNetworksConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterMasterAuthorizedNetworksConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterMasterAuthorizedNetworksConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthorizedNetworksConfigCidrBlocksNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMasterAuthorizedNetworksConfigCidrBlocks)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMasterAuthorizedNetworksConfigCidrBlocks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMasterAuthorizedNetworksConfigCidrBlocks or *ClusterMasterAuthorizedNetworksConfigCidrBlocks", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMasterAuthorizedNetworksConfigCidrBlocks)
	if !ok {
		actualNotPointer, ok := a.(ClusterMasterAuthorizedNetworksConfigCidrBlocks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMasterAuthorizedNetworksConfigCidrBlocks", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CidrBlock, actual.CidrBlock, dcl.Info{}, fn.AddNest("CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMasterAuthorizedNetworksConfigCidrBlocks(c *Client, desired, actual *ClusterMasterAuthorizedNetworksConfigCidrBlocks) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.DisplayName, actual.DisplayName) && !dcl.IsZeroValue(desired.DisplayName) {
		c.Config.Logger.Infof("Diff in DisplayName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DisplayName), dcl.SprintResource(actual.DisplayName))
		return true
	}
	if !dcl.StringCanonicalize(desired.CidrBlock, actual.CidrBlock) && !dcl.IsZeroValue(desired.CidrBlock) {
		c.Config.Logger.Infof("Diff in CidrBlock.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CidrBlock), dcl.SprintResource(actual.CidrBlock))
		return true
	}
	return false
}

func compareClusterMasterAuthorizedNetworksConfigCidrBlocksSlice(c *Client, desired, actual []ClusterMasterAuthorizedNetworksConfigCidrBlocks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMasterAuthorizedNetworksConfigCidrBlocks, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMasterAuthorizedNetworksConfigCidrBlocks(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMasterAuthorizedNetworksConfigCidrBlocks, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthorizedNetworksConfigCidrBlocksMap(c *Client, desired, actual map[string]ClusterMasterAuthorizedNetworksConfigCidrBlocks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMasterAuthorizedNetworksConfigCidrBlocks, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterMasterAuthorizedNetworksConfigCidrBlocks, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterMasterAuthorizedNetworksConfigCidrBlocks(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterMasterAuthorizedNetworksConfigCidrBlocks, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterBinaryAuthorizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterBinaryAuthorization)
	if !ok {
		desiredNotPointer, ok := d.(ClusterBinaryAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterBinaryAuthorization or *ClusterBinaryAuthorization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterBinaryAuthorization)
	if !ok {
		actualNotPointer, ok := a.(ClusterBinaryAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterBinaryAuthorization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterBinaryAuthorization(c *Client, desired, actual *ClusterBinaryAuthorization) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterBinaryAuthorizationSlice(c *Client, desired, actual []ClusterBinaryAuthorization) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterBinaryAuthorization, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterBinaryAuthorization(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterBinaryAuthorization, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterBinaryAuthorizationMap(c *Client, desired, actual map[string]ClusterBinaryAuthorization) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterBinaryAuthorization, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterBinaryAuthorization, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterBinaryAuthorization(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterBinaryAuthorization, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAutoscaling)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscaling or *ClusterAutoscaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAutoscaling)
	if !ok {
		actualNotPointer, ok := a.(ClusterAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnableNodeAutoprovisioning, actual.EnableNodeAutoprovisioning, dcl.Info{}, fn.AddNest("EnableNodeAutoprovisioning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceLimits, actual.ResourceLimits, dcl.Info{ObjectFunction: compareClusterAutoscalingResourceLimitsNewStyle}, fn.AddNest("ResourceLimits")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutoprovisioningNodePoolDefaults, actual.AutoprovisioningNodePoolDefaults, dcl.Info{ObjectFunction: compareClusterAutoscalingAutoprovisioningNodePoolDefaultsNewStyle}, fn.AddNest("AutoprovisioningNodePoolDefaults")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutoprovisioningLocations, actual.AutoprovisioningLocations, dcl.Info{}, fn.AddNest("AutoprovisioningLocations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAutoscaling(c *Client, desired, actual *ClusterAutoscaling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableNodeAutoprovisioning, actual.EnableNodeAutoprovisioning) && !dcl.IsZeroValue(desired.EnableNodeAutoprovisioning) {
		c.Config.Logger.Infof("Diff in EnableNodeAutoprovisioning.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableNodeAutoprovisioning), dcl.SprintResource(actual.EnableNodeAutoprovisioning))
		return true
	}
	if compareClusterAutoscalingResourceLimitsSlice(c, desired.ResourceLimits, actual.ResourceLimits) && !dcl.IsZeroValue(desired.ResourceLimits) {
		c.Config.Logger.Infof("Diff in ResourceLimits.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceLimits), dcl.SprintResource(actual.ResourceLimits))
		return true
	}
	if compareClusterAutoscalingAutoprovisioningNodePoolDefaults(c, desired.AutoprovisioningNodePoolDefaults, actual.AutoprovisioningNodePoolDefaults) && !dcl.IsZeroValue(desired.AutoprovisioningNodePoolDefaults) {
		c.Config.Logger.Infof("Diff in AutoprovisioningNodePoolDefaults.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoprovisioningNodePoolDefaults), dcl.SprintResource(actual.AutoprovisioningNodePoolDefaults))
		return true
	}
	if !dcl.StringSliceEquals(desired.AutoprovisioningLocations, actual.AutoprovisioningLocations) && !dcl.IsZeroValue(desired.AutoprovisioningLocations) {
		c.Config.Logger.Infof("Diff in AutoprovisioningLocations.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoprovisioningLocations), dcl.SprintResource(actual.AutoprovisioningLocations))
		return true
	}
	return false
}

func compareClusterAutoscalingSlice(c *Client, desired, actual []ClusterAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscaling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAutoscaling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAutoscaling, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingMap(c *Client, desired, actual map[string]ClusterAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscaling, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAutoscaling, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAutoscaling(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAutoscaling, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingResourceLimitsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAutoscalingResourceLimits)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAutoscalingResourceLimits)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingResourceLimits or *ClusterAutoscalingResourceLimits", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAutoscalingResourceLimits)
	if !ok {
		actualNotPointer, ok := a.(ClusterAutoscalingResourceLimits)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingResourceLimits", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceType, actual.ResourceType, dcl.Info{}, fn.AddNest("ResourceType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Minimum, actual.Minimum, dcl.Info{}, fn.AddNest("Minimum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Maximum, actual.Maximum, dcl.Info{}, fn.AddNest("Maximum")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAutoscalingResourceLimits(c *Client, desired, actual *ClusterAutoscalingResourceLimits) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.ResourceType, actual.ResourceType) && !dcl.IsZeroValue(desired.ResourceType) {
		c.Config.Logger.Infof("Diff in ResourceType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceType), dcl.SprintResource(actual.ResourceType))
		return true
	}
	if !reflect.DeepEqual(desired.Minimum, actual.Minimum) && !dcl.IsZeroValue(desired.Minimum) {
		c.Config.Logger.Infof("Diff in Minimum.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Minimum), dcl.SprintResource(actual.Minimum))
		return true
	}
	if !reflect.DeepEqual(desired.Maximum, actual.Maximum) && !dcl.IsZeroValue(desired.Maximum) {
		c.Config.Logger.Infof("Diff in Maximum.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Maximum), dcl.SprintResource(actual.Maximum))
		return true
	}
	return false
}

func compareClusterAutoscalingResourceLimitsSlice(c *Client, desired, actual []ClusterAutoscalingResourceLimits) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingResourceLimits, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAutoscalingResourceLimits(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingResourceLimits, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingResourceLimitsMap(c *Client, desired, actual map[string]ClusterAutoscalingResourceLimits) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingResourceLimits, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingResourceLimits, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAutoscalingResourceLimits(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingResourceLimits, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAutoscalingAutoprovisioningNodePoolDefaults)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAutoscalingAutoprovisioningNodePoolDefaults)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaults or *ClusterAutoscalingAutoprovisioningNodePoolDefaults", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAutoscalingAutoprovisioningNodePoolDefaults)
	if !ok {
		actualNotPointer, ok := a.(ClusterAutoscalingAutoprovisioningNodePoolDefaults)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaults", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.OAuthScopes, actual.OAuthScopes, dcl.Info{}, fn.AddNest("OAuthScopes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.Info{}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpgradeSettings, actual.UpgradeSettings, dcl.Info{ObjectFunction: compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsNewStyle}, fn.AddNest("UpgradeSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Management, actual.Management, dcl.Info{ObjectFunction: compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementNewStyle}, fn.AddNest("Management")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinCpuPlatform, actual.MinCpuPlatform, dcl.Info{}, fn.AddNest("MinCpuPlatform")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskSizeGb, actual.DiskSizeGb, dcl.Info{}, fn.AddNest("DiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskType, actual.DiskType, dcl.Info{}, fn.AddNest("DiskType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig, dcl.Info{ObjectFunction: compareClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigNewStyle}, fn.AddNest("ShieldedInstanceConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BootDiskKmsKey, actual.BootDiskKmsKey, dcl.Info{}, fn.AddNest("BootDiskKmsKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaults(c *Client, desired, actual *ClusterAutoscalingAutoprovisioningNodePoolDefaults) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringSliceEquals(desired.OAuthScopes, actual.OAuthScopes) && !dcl.IsZeroValue(desired.OAuthScopes) {
		c.Config.Logger.Infof("Diff in OAuthScopes.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OAuthScopes), dcl.SprintResource(actual.OAuthScopes))
		return true
	}
	if !dcl.StringCanonicalize(desired.ServiceAccount, actual.ServiceAccount) && !dcl.IsZeroValue(desired.ServiceAccount) {
		c.Config.Logger.Infof("Diff in ServiceAccount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccount), dcl.SprintResource(actual.ServiceAccount))
		return true
	}
	if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, desired.UpgradeSettings, actual.UpgradeSettings) && !dcl.IsZeroValue(desired.UpgradeSettings) {
		c.Config.Logger.Infof("Diff in UpgradeSettings.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpgradeSettings), dcl.SprintResource(actual.UpgradeSettings))
		return true
	}
	if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, desired.Management, actual.Management) && !dcl.IsZeroValue(desired.Management) {
		c.Config.Logger.Infof("Diff in Management.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Management), dcl.SprintResource(actual.Management))
		return true
	}
	if !dcl.StringCanonicalize(desired.MinCpuPlatform, actual.MinCpuPlatform) && !dcl.IsZeroValue(desired.MinCpuPlatform) {
		c.Config.Logger.Infof("Diff in MinCpuPlatform.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinCpuPlatform), dcl.SprintResource(actual.MinCpuPlatform))
		return true
	}
	if !reflect.DeepEqual(desired.DiskSizeGb, actual.DiskSizeGb) && !dcl.IsZeroValue(desired.DiskSizeGb) {
		c.Config.Logger.Infof("Diff in DiskSizeGb.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskSizeGb), dcl.SprintResource(actual.DiskSizeGb))
		return true
	}
	if !dcl.StringCanonicalize(desired.DiskType, actual.DiskType) && !dcl.IsZeroValue(desired.DiskType) {
		c.Config.Logger.Infof("Diff in DiskType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskType), dcl.SprintResource(actual.DiskType))
		return true
	}
	if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig) && !dcl.IsZeroValue(desired.ShieldedInstanceConfig) {
		c.Config.Logger.Infof("Diff in ShieldedInstanceConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ShieldedInstanceConfig), dcl.SprintResource(actual.ShieldedInstanceConfig))
		return true
	}
	if !dcl.StringCanonicalize(desired.BootDiskKmsKey, actual.BootDiskKmsKey) && !dcl.IsZeroValue(desired.BootDiskKmsKey) {
		c.Config.Logger.Infof("Diff in BootDiskKmsKey.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BootDiskKmsKey), dcl.SprintResource(actual.BootDiskKmsKey))
		return true
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsSlice(c *Client, desired, actual []ClusterAutoscalingAutoprovisioningNodePoolDefaults) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaults, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaults(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaults, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsMap(c *Client, desired, actual map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaults) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaults, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaults, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaults(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaults, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings or *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings)
	if !ok {
		actualNotPointer, ok := a.(ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxSurge, actual.MaxSurge, dcl.Info{}, fn.AddNest("MaxSurge")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxUnavailable, actual.MaxUnavailable, dcl.Info{}, fn.AddNest("MaxUnavailable")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c *Client, desired, actual *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.MaxSurge, actual.MaxSurge) && !dcl.IsZeroValue(desired.MaxSurge) {
		c.Config.Logger.Infof("Diff in MaxSurge.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxSurge), dcl.SprintResource(actual.MaxSurge))
		return true
	}
	if !reflect.DeepEqual(desired.MaxUnavailable, actual.MaxUnavailable) && !dcl.IsZeroValue(desired.MaxUnavailable) {
		c.Config.Logger.Infof("Diff in MaxUnavailable.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxUnavailable), dcl.SprintResource(actual.MaxUnavailable))
		return true
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsSlice(c *Client, desired, actual []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsMap(c *Client, desired, actual map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement or *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement)
	if !ok {
		actualNotPointer, ok := a.(ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AutoUpgrade, actual.AutoUpgrade, dcl.Info{}, fn.AddNest("AutoUpgrade")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AutoRepair, actual.AutoRepair, dcl.Info{}, fn.AddNest("AutoRepair")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpgradeOptions, actual.UpgradeOptions, dcl.Info{ObjectFunction: compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsNewStyle}, fn.AddNest("UpgradeOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c *Client, desired, actual *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.AutoUpgrade, actual.AutoUpgrade) && !dcl.IsZeroValue(desired.AutoUpgrade) {
		c.Config.Logger.Infof("Diff in AutoUpgrade.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoUpgrade), dcl.SprintResource(actual.AutoUpgrade))
		return true
	}
	if !dcl.BoolCanonicalize(desired.AutoRepair, actual.AutoRepair) && !dcl.IsZeroValue(desired.AutoRepair) {
		c.Config.Logger.Infof("Diff in AutoRepair.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoRepair), dcl.SprintResource(actual.AutoRepair))
		return true
	}
	if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, desired.UpgradeOptions, actual.UpgradeOptions) && !dcl.IsZeroValue(desired.UpgradeOptions) {
		c.Config.Logger.Infof("Diff in UpgradeOptions.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpgradeOptions), dcl.SprintResource(actual.UpgradeOptions))
		return true
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementSlice(c *Client, desired, actual []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementMap(c *Client, desired, actual map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions or *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions)
	if !ok {
		actualNotPointer, ok := a.(ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AutoUpgradeStartTime, actual.AutoUpgradeStartTime, dcl.Info{}, fn.AddNest("AutoUpgradeStartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c *Client, desired, actual *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.AutoUpgradeStartTime, actual.AutoUpgradeStartTime) && !dcl.IsZeroValue(desired.AutoUpgradeStartTime) {
		c.Config.Logger.Infof("Diff in AutoUpgradeStartTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoUpgradeStartTime), dcl.SprintResource(actual.AutoUpgradeStartTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) {
		c.Config.Logger.Infof("Diff in Description.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsSlice(c *Client, desired, actual []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsMap(c *Client, desired, actual map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig or *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnableSecureBoot, actual.EnableSecureBoot, dcl.Info{}, fn.AddNest("EnableSecureBoot")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring, dcl.Info{}, fn.AddNest("EnableIntegrityMonitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c *Client, desired, actual *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableSecureBoot, actual.EnableSecureBoot) && !dcl.IsZeroValue(desired.EnableSecureBoot) {
		c.Config.Logger.Infof("Diff in EnableSecureBoot.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableSecureBoot), dcl.SprintResource(actual.EnableSecureBoot))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring) && !dcl.IsZeroValue(desired.EnableIntegrityMonitoring) {
		c.Config.Logger.Infof("Diff in EnableIntegrityMonitoring.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableIntegrityMonitoring), dcl.SprintResource(actual.EnableIntegrityMonitoring))
		return true
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigSlice(c *Client, desired, actual []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigMap(c *Client, desired, actual map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNetworkConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNetworkConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNetworkConfig or *ClusterNetworkConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNetworkConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNetworkConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{OutputOnly: true}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subnetwork, actual.Subnetwork, dcl.Info{OutputOnly: true}, fn.AddNest("Subnetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableIntraNodeVisibility, actual.EnableIntraNodeVisibility, dcl.Info{}, fn.AddNest("EnableIntraNodeVisibility")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultSnatStatus, actual.DefaultSnatStatus, dcl.Info{ObjectFunction: compareClusterNetworkConfigDefaultSnatStatusNewStyle}, fn.AddNest("DefaultSnatStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateIPv6GoogleAccess, actual.PrivateIPv6GoogleAccess, dcl.Info{Type: "EnumType"}, fn.AddNest("PrivateIPv6GoogleAccess")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNetworkConfig(c *Client, desired, actual *ClusterNetworkConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableIntraNodeVisibility, actual.EnableIntraNodeVisibility) && !dcl.IsZeroValue(desired.EnableIntraNodeVisibility) {
		c.Config.Logger.Infof("Diff in EnableIntraNodeVisibility.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableIntraNodeVisibility), dcl.SprintResource(actual.EnableIntraNodeVisibility))
		return true
	}
	if compareClusterNetworkConfigDefaultSnatStatus(c, desired.DefaultSnatStatus, actual.DefaultSnatStatus) && !dcl.IsZeroValue(desired.DefaultSnatStatus) {
		c.Config.Logger.Infof("Diff in DefaultSnatStatus.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DefaultSnatStatus), dcl.SprintResource(actual.DefaultSnatStatus))
		return true
	}
	if !reflect.DeepEqual(desired.PrivateIPv6GoogleAccess, actual.PrivateIPv6GoogleAccess) && !dcl.IsZeroValue(desired.PrivateIPv6GoogleAccess) {
		c.Config.Logger.Infof("Diff in PrivateIPv6GoogleAccess.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrivateIPv6GoogleAccess), dcl.SprintResource(actual.PrivateIPv6GoogleAccess))
		return true
	}
	return false
}

func compareClusterNetworkConfigSlice(c *Client, desired, actual []ClusterNetworkConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNetworkConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNetworkConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNetworkConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNetworkConfigMap(c *Client, desired, actual map[string]ClusterNetworkConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNetworkConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNetworkConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNetworkConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNetworkConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNetworkConfigDefaultSnatStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNetworkConfigDefaultSnatStatus)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNetworkConfigDefaultSnatStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNetworkConfigDefaultSnatStatus or *ClusterNetworkConfigDefaultSnatStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNetworkConfigDefaultSnatStatus)
	if !ok {
		actualNotPointer, ok := a.(ClusterNetworkConfigDefaultSnatStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNetworkConfigDefaultSnatStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNetworkConfigDefaultSnatStatus(c *Client, desired, actual *ClusterNetworkConfigDefaultSnatStatus) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) {
		c.Config.Logger.Infof("Diff in Disabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	return false
}

func compareClusterNetworkConfigDefaultSnatStatusSlice(c *Client, desired, actual []ClusterNetworkConfigDefaultSnatStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNetworkConfigDefaultSnatStatus, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNetworkConfigDefaultSnatStatus(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNetworkConfigDefaultSnatStatus, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNetworkConfigDefaultSnatStatusMap(c *Client, desired, actual map[string]ClusterNetworkConfigDefaultSnatStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNetworkConfigDefaultSnatStatus, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNetworkConfigDefaultSnatStatus, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNetworkConfigDefaultSnatStatus(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNetworkConfigDefaultSnatStatus, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMaintenancePolicy)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMaintenancePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicy or *ClusterMaintenancePolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMaintenancePolicy)
	if !ok {
		actualNotPointer, ok := a.(ClusterMaintenancePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Window, actual.Window, dcl.Info{ObjectFunction: compareClusterMaintenancePolicyWindowNewStyle}, fn.AddNest("Window")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceVersion, actual.ResourceVersion, dcl.Info{}, fn.AddNest("ResourceVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMaintenancePolicy(c *Client, desired, actual *ClusterMaintenancePolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareClusterMaintenancePolicyWindow(c, desired.Window, actual.Window) && !dcl.IsZeroValue(desired.Window) {
		c.Config.Logger.Infof("Diff in Window.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Window), dcl.SprintResource(actual.Window))
		return true
	}
	if !dcl.StringCanonicalize(desired.ResourceVersion, actual.ResourceVersion) && !dcl.IsZeroValue(desired.ResourceVersion) {
		c.Config.Logger.Infof("Diff in ResourceVersion.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceVersion), dcl.SprintResource(actual.ResourceVersion))
		return true
	}
	return false
}

func compareClusterMaintenancePolicySlice(c *Client, desired, actual []ClusterMaintenancePolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMaintenancePolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicy, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyMap(c *Client, desired, actual map[string]ClusterMaintenancePolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterMaintenancePolicy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicy, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMaintenancePolicyWindow)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMaintenancePolicyWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicyWindow or *ClusterMaintenancePolicyWindow", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMaintenancePolicyWindow)
	if !ok {
		actualNotPointer, ok := a.(ClusterMaintenancePolicyWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicyWindow", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DailyMaintenanceWindow, actual.DailyMaintenanceWindow, dcl.Info{ObjectFunction: compareClusterMaintenancePolicyWindowDailyMaintenanceWindowNewStyle}, fn.AddNest("DailyMaintenanceWindow")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RecurringWindow, actual.RecurringWindow, dcl.Info{ObjectFunction: compareClusterMaintenancePolicyWindowRecurringWindowNewStyle}, fn.AddNest("RecurringWindow")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaintenanceExclusions, actual.MaintenanceExclusions, dcl.Info{}, fn.AddNest("MaintenanceExclusions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMaintenancePolicyWindow(c *Client, desired, actual *ClusterMaintenancePolicyWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, desired.DailyMaintenanceWindow, actual.DailyMaintenanceWindow) && !dcl.IsZeroValue(desired.DailyMaintenanceWindow) {
		c.Config.Logger.Infof("Diff in DailyMaintenanceWindow.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DailyMaintenanceWindow), dcl.SprintResource(actual.DailyMaintenanceWindow))
		return true
	}
	if compareClusterMaintenancePolicyWindowRecurringWindow(c, desired.RecurringWindow, actual.RecurringWindow) && !dcl.IsZeroValue(desired.RecurringWindow) {
		c.Config.Logger.Infof("Diff in RecurringWindow.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RecurringWindow), dcl.SprintResource(actual.RecurringWindow))
		return true
	}
	if !dcl.MapEquals(desired.MaintenanceExclusions, actual.MaintenanceExclusions, []string(nil)) && !dcl.IsZeroValue(desired.MaintenanceExclusions) {
		c.Config.Logger.Infof("Diff in MaintenanceExclusions.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaintenanceExclusions), dcl.SprintResource(actual.MaintenanceExclusions))
		return true
	}
	return false
}

func compareClusterMaintenancePolicyWindowSlice(c *Client, desired, actual []ClusterMaintenancePolicyWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicyWindow, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMaintenancePolicyWindow(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindow, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowMap(c *Client, desired, actual map[string]ClusterMaintenancePolicyWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicyWindow, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindow, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterMaintenancePolicyWindow(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindow, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowDailyMaintenanceWindowNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMaintenancePolicyWindowDailyMaintenanceWindow)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMaintenancePolicyWindowDailyMaintenanceWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicyWindowDailyMaintenanceWindow or *ClusterMaintenancePolicyWindowDailyMaintenanceWindow", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMaintenancePolicyWindowDailyMaintenanceWindow)
	if !ok {
		actualNotPointer, ok := a.(ClusterMaintenancePolicyWindowDailyMaintenanceWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicyWindowDailyMaintenanceWindow", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StartTime, actual.StartTime, dcl.Info{}, fn.AddNest("StartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Duration, actual.Duration, dcl.Info{}, fn.AddNest("Duration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMaintenancePolicyWindowDailyMaintenanceWindow(c *Client, desired, actual *ClusterMaintenancePolicyWindowDailyMaintenanceWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) {
		c.Config.Logger.Infof("Diff in StartTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.Duration, actual.Duration) && !dcl.IsZeroValue(desired.Duration) {
		c.Config.Logger.Infof("Diff in Duration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Duration), dcl.SprintResource(actual.Duration))
		return true
	}
	return false
}

func compareClusterMaintenancePolicyWindowDailyMaintenanceWindowSlice(c *Client, desired, actual []ClusterMaintenancePolicyWindowDailyMaintenanceWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicyWindowDailyMaintenanceWindow, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowDailyMaintenanceWindow, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowDailyMaintenanceWindowMap(c *Client, desired, actual map[string]ClusterMaintenancePolicyWindowDailyMaintenanceWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicyWindowDailyMaintenanceWindow, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowDailyMaintenanceWindow, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowDailyMaintenanceWindow, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowRecurringWindowNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMaintenancePolicyWindowRecurringWindow)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMaintenancePolicyWindowRecurringWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicyWindowRecurringWindow or *ClusterMaintenancePolicyWindowRecurringWindow", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMaintenancePolicyWindowRecurringWindow)
	if !ok {
		actualNotPointer, ok := a.(ClusterMaintenancePolicyWindowRecurringWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicyWindowRecurringWindow", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Window, actual.Window, dcl.Info{ObjectFunction: compareClusterMaintenancePolicyWindowRecurringWindowWindowNewStyle}, fn.AddNest("Window")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Recurrence, actual.Recurrence, dcl.Info{}, fn.AddNest("Recurrence")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMaintenancePolicyWindowRecurringWindow(c *Client, desired, actual *ClusterMaintenancePolicyWindowRecurringWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareClusterMaintenancePolicyWindowRecurringWindowWindow(c, desired.Window, actual.Window) && !dcl.IsZeroValue(desired.Window) {
		c.Config.Logger.Infof("Diff in Window.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Window), dcl.SprintResource(actual.Window))
		return true
	}
	if !dcl.StringCanonicalize(desired.Recurrence, actual.Recurrence) && !dcl.IsZeroValue(desired.Recurrence) {
		c.Config.Logger.Infof("Diff in Recurrence.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Recurrence), dcl.SprintResource(actual.Recurrence))
		return true
	}
	return false
}

func compareClusterMaintenancePolicyWindowRecurringWindowSlice(c *Client, desired, actual []ClusterMaintenancePolicyWindowRecurringWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicyWindowRecurringWindow, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMaintenancePolicyWindowRecurringWindow(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowRecurringWindow, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowRecurringWindowMap(c *Client, desired, actual map[string]ClusterMaintenancePolicyWindowRecurringWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicyWindowRecurringWindow, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowRecurringWindow, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterMaintenancePolicyWindowRecurringWindow(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowRecurringWindow, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowRecurringWindowWindowNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterMaintenancePolicyWindowRecurringWindowWindow)
	if !ok {
		desiredNotPointer, ok := d.(ClusterMaintenancePolicyWindowRecurringWindowWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicyWindowRecurringWindowWindow or *ClusterMaintenancePolicyWindowRecurringWindowWindow", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterMaintenancePolicyWindowRecurringWindowWindow)
	if !ok {
		actualNotPointer, ok := a.(ClusterMaintenancePolicyWindowRecurringWindowWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterMaintenancePolicyWindowRecurringWindowWindow", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StartTime, actual.StartTime, dcl.Info{}, fn.AddNest("StartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndTime, actual.EndTime, dcl.Info{}, fn.AddNest("EndTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterMaintenancePolicyWindowRecurringWindowWindow(c *Client, desired, actual *ClusterMaintenancePolicyWindowRecurringWindowWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) {
		c.Config.Logger.Infof("Diff in StartTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if !reflect.DeepEqual(desired.EndTime, actual.EndTime) && !dcl.IsZeroValue(desired.EndTime) {
		c.Config.Logger.Infof("Diff in EndTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EndTime), dcl.SprintResource(actual.EndTime))
		return true
	}
	return false
}

func compareClusterMaintenancePolicyWindowRecurringWindowWindowSlice(c *Client, desired, actual []ClusterMaintenancePolicyWindowRecurringWindowWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicyWindowRecurringWindowWindow, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMaintenancePolicyWindowRecurringWindowWindow(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowRecurringWindowWindow, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowRecurringWindowWindowMap(c *Client, desired, actual map[string]ClusterMaintenancePolicyWindowRecurringWindowWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMaintenancePolicyWindowRecurringWindowWindow, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowRecurringWindowWindow, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterMaintenancePolicyWindowRecurringWindowWindow(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowRecurringWindowWindow, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterDefaultMaxPodsConstraintNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterDefaultMaxPodsConstraint)
	if !ok {
		desiredNotPointer, ok := d.(ClusterDefaultMaxPodsConstraint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterDefaultMaxPodsConstraint or *ClusterDefaultMaxPodsConstraint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterDefaultMaxPodsConstraint)
	if !ok {
		actualNotPointer, ok := a.(ClusterDefaultMaxPodsConstraint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterDefaultMaxPodsConstraint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxPodsPerNode, actual.MaxPodsPerNode, dcl.Info{}, fn.AddNest("MaxPodsPerNode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterDefaultMaxPodsConstraint(c *Client, desired, actual *ClusterDefaultMaxPodsConstraint) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.MaxPodsPerNode, actual.MaxPodsPerNode) && !dcl.IsZeroValue(desired.MaxPodsPerNode) {
		c.Config.Logger.Infof("Diff in MaxPodsPerNode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxPodsPerNode), dcl.SprintResource(actual.MaxPodsPerNode))
		return true
	}
	return false
}

func compareClusterDefaultMaxPodsConstraintSlice(c *Client, desired, actual []ClusterDefaultMaxPodsConstraint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterDefaultMaxPodsConstraint, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterDefaultMaxPodsConstraint(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterDefaultMaxPodsConstraint, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterDefaultMaxPodsConstraintMap(c *Client, desired, actual map[string]ClusterDefaultMaxPodsConstraint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterDefaultMaxPodsConstraint, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterDefaultMaxPodsConstraint, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterDefaultMaxPodsConstraint(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterDefaultMaxPodsConstraint, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterResourceUsageExportConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterResourceUsageExportConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterResourceUsageExportConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterResourceUsageExportConfig or *ClusterResourceUsageExportConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterResourceUsageExportConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterResourceUsageExportConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterResourceUsageExportConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BigqueryDestination, actual.BigqueryDestination, dcl.Info{ObjectFunction: compareClusterResourceUsageExportConfigBigqueryDestinationNewStyle}, fn.AddNest("BigqueryDestination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableNetworkEgressMonitoring, actual.EnableNetworkEgressMonitoring, dcl.Info{}, fn.AddNest("EnableNetworkEgressMonitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConsumptionMeteringConfig, actual.ConsumptionMeteringConfig, dcl.Info{ObjectFunction: compareClusterResourceUsageExportConfigConsumptionMeteringConfigNewStyle}, fn.AddNest("ConsumptionMeteringConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableNetworkEgressMetering, actual.EnableNetworkEgressMetering, dcl.Info{}, fn.AddNest("EnableNetworkEgressMetering")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterResourceUsageExportConfig(c *Client, desired, actual *ClusterResourceUsageExportConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareClusterResourceUsageExportConfigBigqueryDestination(c, desired.BigqueryDestination, actual.BigqueryDestination) && !dcl.IsZeroValue(desired.BigqueryDestination) {
		c.Config.Logger.Infof("Diff in BigqueryDestination.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BigqueryDestination), dcl.SprintResource(actual.BigqueryDestination))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableNetworkEgressMonitoring, actual.EnableNetworkEgressMonitoring) && !dcl.IsZeroValue(desired.EnableNetworkEgressMonitoring) {
		c.Config.Logger.Infof("Diff in EnableNetworkEgressMonitoring.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableNetworkEgressMonitoring), dcl.SprintResource(actual.EnableNetworkEgressMonitoring))
		return true
	}
	if compareClusterResourceUsageExportConfigConsumptionMeteringConfig(c, desired.ConsumptionMeteringConfig, actual.ConsumptionMeteringConfig) && !dcl.IsZeroValue(desired.ConsumptionMeteringConfig) {
		c.Config.Logger.Infof("Diff in ConsumptionMeteringConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConsumptionMeteringConfig), dcl.SprintResource(actual.ConsumptionMeteringConfig))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableNetworkEgressMetering, actual.EnableNetworkEgressMetering) && !dcl.IsZeroValue(desired.EnableNetworkEgressMetering) {
		c.Config.Logger.Infof("Diff in EnableNetworkEgressMetering.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableNetworkEgressMetering), dcl.SprintResource(actual.EnableNetworkEgressMetering))
		return true
	}
	return false
}

func compareClusterResourceUsageExportConfigSlice(c *Client, desired, actual []ClusterResourceUsageExportConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterResourceUsageExportConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterResourceUsageExportConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterResourceUsageExportConfigMap(c *Client, desired, actual map[string]ClusterResourceUsageExportConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterResourceUsageExportConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterResourceUsageExportConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterResourceUsageExportConfigBigqueryDestinationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterResourceUsageExportConfigBigqueryDestination)
	if !ok {
		desiredNotPointer, ok := d.(ClusterResourceUsageExportConfigBigqueryDestination)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterResourceUsageExportConfigBigqueryDestination or *ClusterResourceUsageExportConfigBigqueryDestination", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterResourceUsageExportConfigBigqueryDestination)
	if !ok {
		actualNotPointer, ok := a.(ClusterResourceUsageExportConfigBigqueryDestination)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterResourceUsageExportConfigBigqueryDestination", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DatasetId, actual.DatasetId, dcl.Info{}, fn.AddNest("DatasetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterResourceUsageExportConfigBigqueryDestination(c *Client, desired, actual *ClusterResourceUsageExportConfigBigqueryDestination) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.DatasetId, actual.DatasetId) && !dcl.IsZeroValue(desired.DatasetId) {
		c.Config.Logger.Infof("Diff in DatasetId.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DatasetId), dcl.SprintResource(actual.DatasetId))
		return true
	}
	return false
}

func compareClusterResourceUsageExportConfigBigqueryDestinationSlice(c *Client, desired, actual []ClusterResourceUsageExportConfigBigqueryDestination) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterResourceUsageExportConfigBigqueryDestination, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterResourceUsageExportConfigBigqueryDestination(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfigBigqueryDestination, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterResourceUsageExportConfigBigqueryDestinationMap(c *Client, desired, actual map[string]ClusterResourceUsageExportConfigBigqueryDestination) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterResourceUsageExportConfigBigqueryDestination, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfigBigqueryDestination, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterResourceUsageExportConfigBigqueryDestination(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfigBigqueryDestination, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterResourceUsageExportConfigConsumptionMeteringConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterResourceUsageExportConfigConsumptionMeteringConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterResourceUsageExportConfigConsumptionMeteringConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterResourceUsageExportConfigConsumptionMeteringConfig or *ClusterResourceUsageExportConfigConsumptionMeteringConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterResourceUsageExportConfigConsumptionMeteringConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterResourceUsageExportConfigConsumptionMeteringConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterResourceUsageExportConfigConsumptionMeteringConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterResourceUsageExportConfigConsumptionMeteringConfig(c *Client, desired, actual *ClusterResourceUsageExportConfigConsumptionMeteringConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterResourceUsageExportConfigConsumptionMeteringConfigSlice(c *Client, desired, actual []ClusterResourceUsageExportConfigConsumptionMeteringConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterResourceUsageExportConfigConsumptionMeteringConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterResourceUsageExportConfigConsumptionMeteringConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfigConsumptionMeteringConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterResourceUsageExportConfigConsumptionMeteringConfigMap(c *Client, desired, actual map[string]ClusterResourceUsageExportConfigConsumptionMeteringConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterResourceUsageExportConfigConsumptionMeteringConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfigConsumptionMeteringConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterResourceUsageExportConfigConsumptionMeteringConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfigConsumptionMeteringConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAuthenticatorGroupsConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAuthenticatorGroupsConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAuthenticatorGroupsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAuthenticatorGroupsConfig or *ClusterAuthenticatorGroupsConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAuthenticatorGroupsConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterAuthenticatorGroupsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAuthenticatorGroupsConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityGroup, actual.SecurityGroup, dcl.Info{}, fn.AddNest("SecurityGroup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAuthenticatorGroupsConfig(c *Client, desired, actual *ClusterAuthenticatorGroupsConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	if !dcl.StringCanonicalize(desired.SecurityGroup, actual.SecurityGroup) && !dcl.IsZeroValue(desired.SecurityGroup) {
		c.Config.Logger.Infof("Diff in SecurityGroup.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecurityGroup), dcl.SprintResource(actual.SecurityGroup))
		return true
	}
	return false
}

func compareClusterAuthenticatorGroupsConfigSlice(c *Client, desired, actual []ClusterAuthenticatorGroupsConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAuthenticatorGroupsConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAuthenticatorGroupsConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAuthenticatorGroupsConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAuthenticatorGroupsConfigMap(c *Client, desired, actual map[string]ClusterAuthenticatorGroupsConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAuthenticatorGroupsConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAuthenticatorGroupsConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAuthenticatorGroupsConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAuthenticatorGroupsConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterPrivateClusterConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterPrivateClusterConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterPrivateClusterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterPrivateClusterConfig or *ClusterPrivateClusterConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterPrivateClusterConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterPrivateClusterConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterPrivateClusterConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnablePrivateNodes, actual.EnablePrivateNodes, dcl.Info{}, fn.AddNest("EnablePrivateNodes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnablePrivateEndpoint, actual.EnablePrivateEndpoint, dcl.Info{}, fn.AddNest("EnablePrivateEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MasterIPv4CidrBlock, actual.MasterIPv4CidrBlock, dcl.Info{}, fn.AddNest("MasterIPv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateEndpoint, actual.PrivateEndpoint, dcl.Info{OutputOnly: true}, fn.AddNest("PrivateEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublicEndpoint, actual.PublicEndpoint, dcl.Info{OutputOnly: true}, fn.AddNest("PublicEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PeeringName, actual.PeeringName, dcl.Info{OutputOnly: true}, fn.AddNest("PeeringName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MasterGlobalAccessConfig, actual.MasterGlobalAccessConfig, dcl.Info{ObjectFunction: compareClusterPrivateClusterConfigMasterGlobalAccessConfigNewStyle}, fn.AddNest("MasterGlobalAccessConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterPrivateClusterConfig(c *Client, desired, actual *ClusterPrivateClusterConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnablePrivateNodes, actual.EnablePrivateNodes) && !dcl.IsZeroValue(desired.EnablePrivateNodes) {
		c.Config.Logger.Infof("Diff in EnablePrivateNodes.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnablePrivateNodes), dcl.SprintResource(actual.EnablePrivateNodes))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnablePrivateEndpoint, actual.EnablePrivateEndpoint) && !dcl.IsZeroValue(desired.EnablePrivateEndpoint) {
		c.Config.Logger.Infof("Diff in EnablePrivateEndpoint.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnablePrivateEndpoint), dcl.SprintResource(actual.EnablePrivateEndpoint))
		return true
	}
	if !dcl.StringCanonicalize(desired.MasterIPv4CidrBlock, actual.MasterIPv4CidrBlock) && !dcl.IsZeroValue(desired.MasterIPv4CidrBlock) {
		c.Config.Logger.Infof("Diff in MasterIPv4CidrBlock.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MasterIPv4CidrBlock), dcl.SprintResource(actual.MasterIPv4CidrBlock))
		return true
	}
	if compareClusterPrivateClusterConfigMasterGlobalAccessConfig(c, desired.MasterGlobalAccessConfig, actual.MasterGlobalAccessConfig) && !dcl.IsZeroValue(desired.MasterGlobalAccessConfig) {
		c.Config.Logger.Infof("Diff in MasterGlobalAccessConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MasterGlobalAccessConfig), dcl.SprintResource(actual.MasterGlobalAccessConfig))
		return true
	}
	return false
}

func compareClusterPrivateClusterConfigSlice(c *Client, desired, actual []ClusterPrivateClusterConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterPrivateClusterConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterPrivateClusterConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterPrivateClusterConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterPrivateClusterConfigMap(c *Client, desired, actual map[string]ClusterPrivateClusterConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterPrivateClusterConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterPrivateClusterConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterPrivateClusterConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterPrivateClusterConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterPrivateClusterConfigMasterGlobalAccessConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterPrivateClusterConfigMasterGlobalAccessConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterPrivateClusterConfigMasterGlobalAccessConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterPrivateClusterConfigMasterGlobalAccessConfig or *ClusterPrivateClusterConfigMasterGlobalAccessConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterPrivateClusterConfigMasterGlobalAccessConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterPrivateClusterConfigMasterGlobalAccessConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterPrivateClusterConfigMasterGlobalAccessConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterPrivateClusterConfigMasterGlobalAccessConfig(c *Client, desired, actual *ClusterPrivateClusterConfigMasterGlobalAccessConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterPrivateClusterConfigMasterGlobalAccessConfigSlice(c *Client, desired, actual []ClusterPrivateClusterConfigMasterGlobalAccessConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterPrivateClusterConfigMasterGlobalAccessConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterPrivateClusterConfigMasterGlobalAccessConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterPrivateClusterConfigMasterGlobalAccessConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterPrivateClusterConfigMasterGlobalAccessConfigMap(c *Client, desired, actual map[string]ClusterPrivateClusterConfigMasterGlobalAccessConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterPrivateClusterConfigMasterGlobalAccessConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterPrivateClusterConfigMasterGlobalAccessConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterPrivateClusterConfigMasterGlobalAccessConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterPrivateClusterConfigMasterGlobalAccessConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterDatabaseEncryptionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterDatabaseEncryption)
	if !ok {
		desiredNotPointer, ok := d.(ClusterDatabaseEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterDatabaseEncryption or *ClusterDatabaseEncryption", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterDatabaseEncryption)
	if !ok {
		actualNotPointer, ok := a.(ClusterDatabaseEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterDatabaseEncryption", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{Type: "EnumType"}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyName, actual.KeyName, dcl.Info{}, fn.AddNest("KeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterDatabaseEncryption(c *Client, desired, actual *ClusterDatabaseEncryption) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.State, actual.State) && !dcl.IsZeroValue(desired.State) {
		c.Config.Logger.Infof("Diff in State.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.State), dcl.SprintResource(actual.State))
		return true
	}
	if !dcl.StringCanonicalize(desired.KeyName, actual.KeyName) && !dcl.IsZeroValue(desired.KeyName) {
		c.Config.Logger.Infof("Diff in KeyName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KeyName), dcl.SprintResource(actual.KeyName))
		return true
	}
	return false
}

func compareClusterDatabaseEncryptionSlice(c *Client, desired, actual []ClusterDatabaseEncryption) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterDatabaseEncryption, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterDatabaseEncryption(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterDatabaseEncryption, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterDatabaseEncryptionMap(c *Client, desired, actual map[string]ClusterDatabaseEncryption) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterDatabaseEncryption, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterDatabaseEncryption, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterDatabaseEncryption(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterDatabaseEncryption, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterVerticalPodAutoscalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterVerticalPodAutoscaling)
	if !ok {
		desiredNotPointer, ok := d.(ClusterVerticalPodAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterVerticalPodAutoscaling or *ClusterVerticalPodAutoscaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterVerticalPodAutoscaling)
	if !ok {
		actualNotPointer, ok := a.(ClusterVerticalPodAutoscaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterVerticalPodAutoscaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterVerticalPodAutoscaling(c *Client, desired, actual *ClusterVerticalPodAutoscaling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterVerticalPodAutoscalingSlice(c *Client, desired, actual []ClusterVerticalPodAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterVerticalPodAutoscaling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterVerticalPodAutoscaling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterVerticalPodAutoscaling, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterVerticalPodAutoscalingMap(c *Client, desired, actual map[string]ClusterVerticalPodAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterVerticalPodAutoscaling, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterVerticalPodAutoscaling, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterVerticalPodAutoscaling(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterVerticalPodAutoscaling, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterShieldedNodesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterShieldedNodes)
	if !ok {
		desiredNotPointer, ok := d.(ClusterShieldedNodes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterShieldedNodes or *ClusterShieldedNodes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterShieldedNodes)
	if !ok {
		actualNotPointer, ok := a.(ClusterShieldedNodes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterShieldedNodes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterShieldedNodes(c *Client, desired, actual *ClusterShieldedNodes) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterShieldedNodesSlice(c *Client, desired, actual []ClusterShieldedNodes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterShieldedNodes, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterShieldedNodes(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterShieldedNodes, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterShieldedNodesMap(c *Client, desired, actual map[string]ClusterShieldedNodes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterShieldedNodes, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterShieldedNodes, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterShieldedNodes(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterShieldedNodes, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterConditionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConditions)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConditions or *ClusterConditions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConditions)
	if !ok {
		actualNotPointer, ok := a.(ClusterConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConditions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CanonicalCode, actual.CanonicalCode, dcl.Info{Type: "EnumType"}, fn.AddNest("CanonicalCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConditions(c *Client, desired, actual *ClusterConditions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Code, actual.Code) && !dcl.IsZeroValue(desired.Code) {
		c.Config.Logger.Infof("Diff in Code.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Code), dcl.SprintResource(actual.Code))
		return true
	}
	if !dcl.StringCanonicalize(desired.Message, actual.Message) && !dcl.IsZeroValue(desired.Message) {
		c.Config.Logger.Infof("Diff in Message.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Message), dcl.SprintResource(actual.Message))
		return true
	}
	if !reflect.DeepEqual(desired.CanonicalCode, actual.CanonicalCode) && !dcl.IsZeroValue(desired.CanonicalCode) {
		c.Config.Logger.Infof("Diff in CanonicalCode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CanonicalCode), dcl.SprintResource(actual.CanonicalCode))
		return true
	}
	return false
}

func compareClusterConditionsSlice(c *Client, desired, actual []ClusterConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterConditions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterConditions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterConditions, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterConditionsMap(c *Client, desired, actual map[string]ClusterConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterConditions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterConditions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterConditions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterConditions, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAutopilotNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAutopilot)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAutopilot)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutopilot or *ClusterAutopilot", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAutopilot)
	if !ok {
		actualNotPointer, ok := a.(ClusterAutopilot)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAutopilot", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAutopilot(c *Client, desired, actual *ClusterAutopilot) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterAutopilotSlice(c *Client, desired, actual []ClusterAutopilot) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutopilot, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAutopilot(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAutopilot, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutopilotMap(c *Client, desired, actual map[string]ClusterAutopilot) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAutopilot, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterAutopilot, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterAutopilot(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterAutopilot, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfig or *ClusterNodeConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskSizeGb, actual.DiskSizeGb, dcl.Info{}, fn.AddNest("DiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuthScopes, actual.OAuthScopes, dcl.Info{}, fn.AddNest("OAuthScopes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.Info{}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImageType, actual.ImageType, dcl.Info{}, fn.AddNest("ImageType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalSsdCount, actual.LocalSsdCount, dcl.Info{}, fn.AddNest("LocalSsdCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tags, actual.Tags, dcl.Info{}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Preemptible, actual.Preemptible, dcl.Info{}, fn.AddNest("Preemptible")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Accelerators, actual.Accelerators, dcl.Info{ObjectFunction: compareClusterNodeConfigAcceleratorsNewStyle}, fn.AddNest("Accelerators")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskType, actual.DiskType, dcl.Info{}, fn.AddNest("DiskType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinCpuPlatform, actual.MinCpuPlatform, dcl.Info{}, fn.AddNest("MinCpuPlatform")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkloadMetadataConfig, actual.WorkloadMetadataConfig, dcl.Info{ObjectFunction: compareClusterNodeConfigWorkloadMetadataConfigNewStyle}, fn.AddNest("WorkloadMetadataConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Taints, actual.Taints, dcl.Info{ObjectFunction: compareClusterNodeConfigTaintsNewStyle}, fn.AddNest("Taints")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SandboxConfig, actual.SandboxConfig, dcl.Info{ObjectFunction: compareClusterNodeConfigSandboxConfigNewStyle}, fn.AddNest("SandboxConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeGroup, actual.NodeGroup, dcl.Info{}, fn.AddNest("NodeGroup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReservationAffinity, actual.ReservationAffinity, dcl.Info{ObjectFunction: compareClusterNodeConfigReservationAffinityNewStyle}, fn.AddNest("ReservationAffinity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig, dcl.Info{ObjectFunction: compareClusterNodeConfigShieldedInstanceConfigNewStyle}, fn.AddNest("ShieldedInstanceConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LinuxNodeConfig, actual.LinuxNodeConfig, dcl.Info{ObjectFunction: compareClusterNodeConfigLinuxNodeConfigNewStyle}, fn.AddNest("LinuxNodeConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubeletConfig, actual.KubeletConfig, dcl.Info{ObjectFunction: compareClusterNodeConfigKubeletConfigNewStyle}, fn.AddNest("KubeletConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BootDiskKmsKey, actual.BootDiskKmsKey, dcl.Info{}, fn.AddNest("BootDiskKmsKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodeConfig(c *Client, desired, actual *ClusterNodeConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.MachineType, actual.MachineType) && !dcl.IsZeroValue(desired.MachineType) {
		c.Config.Logger.Infof("Diff in MachineType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MachineType), dcl.SprintResource(actual.MachineType))
		return true
	}
	if !reflect.DeepEqual(desired.DiskSizeGb, actual.DiskSizeGb) && !dcl.IsZeroValue(desired.DiskSizeGb) {
		c.Config.Logger.Infof("Diff in DiskSizeGb.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskSizeGb), dcl.SprintResource(actual.DiskSizeGb))
		return true
	}
	if !dcl.StringSliceEquals(desired.OAuthScopes, actual.OAuthScopes) && !dcl.IsZeroValue(desired.OAuthScopes) {
		c.Config.Logger.Infof("Diff in OAuthScopes.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OAuthScopes), dcl.SprintResource(actual.OAuthScopes))
		return true
	}
	if !dcl.StringCanonicalize(desired.ServiceAccount, actual.ServiceAccount) && !dcl.IsZeroValue(desired.ServiceAccount) {
		c.Config.Logger.Infof("Diff in ServiceAccount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccount), dcl.SprintResource(actual.ServiceAccount))
		return true
	}
	if !dcl.MapEquals(desired.Metadata, actual.Metadata, []string(nil)) && !dcl.IsZeroValue(desired.Metadata) {
		c.Config.Logger.Infof("Diff in Metadata.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Metadata), dcl.SprintResource(actual.Metadata))
		return true
	}
	if !dcl.StringCanonicalize(desired.ImageType, actual.ImageType) && !dcl.IsZeroValue(desired.ImageType) {
		c.Config.Logger.Infof("Diff in ImageType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ImageType), dcl.SprintResource(actual.ImageType))
		return true
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	if !reflect.DeepEqual(desired.LocalSsdCount, actual.LocalSsdCount) && !dcl.IsZeroValue(desired.LocalSsdCount) {
		c.Config.Logger.Infof("Diff in LocalSsdCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalSsdCount), dcl.SprintResource(actual.LocalSsdCount))
		return true
	}
	if !dcl.StringSliceEquals(desired.Tags, actual.Tags) && !dcl.IsZeroValue(desired.Tags) {
		c.Config.Logger.Infof("Diff in Tags.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tags), dcl.SprintResource(actual.Tags))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Preemptible, actual.Preemptible) && !dcl.IsZeroValue(desired.Preemptible) {
		c.Config.Logger.Infof("Diff in Preemptible.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Preemptible), dcl.SprintResource(actual.Preemptible))
		return true
	}
	if compareClusterNodeConfigAcceleratorsSlice(c, desired.Accelerators, actual.Accelerators) && !dcl.IsZeroValue(desired.Accelerators) {
		c.Config.Logger.Infof("Diff in Accelerators.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Accelerators), dcl.SprintResource(actual.Accelerators))
		return true
	}
	if !dcl.StringCanonicalize(desired.DiskType, actual.DiskType) && !dcl.IsZeroValue(desired.DiskType) {
		c.Config.Logger.Infof("Diff in DiskType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskType), dcl.SprintResource(actual.DiskType))
		return true
	}
	if !dcl.StringCanonicalize(desired.MinCpuPlatform, actual.MinCpuPlatform) && !dcl.IsZeroValue(desired.MinCpuPlatform) {
		c.Config.Logger.Infof("Diff in MinCpuPlatform.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinCpuPlatform), dcl.SprintResource(actual.MinCpuPlatform))
		return true
	}
	if compareClusterNodeConfigWorkloadMetadataConfig(c, desired.WorkloadMetadataConfig, actual.WorkloadMetadataConfig) && !dcl.IsZeroValue(desired.WorkloadMetadataConfig) {
		c.Config.Logger.Infof("Diff in WorkloadMetadataConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WorkloadMetadataConfig), dcl.SprintResource(actual.WorkloadMetadataConfig))
		return true
	}
	if compareClusterNodeConfigTaintsSlice(c, desired.Taints, actual.Taints) && !dcl.IsZeroValue(desired.Taints) {
		c.Config.Logger.Infof("Diff in Taints.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Taints), dcl.SprintResource(actual.Taints))
		return true
	}
	if compareClusterNodeConfigSandboxConfig(c, desired.SandboxConfig, actual.SandboxConfig) && !dcl.IsZeroValue(desired.SandboxConfig) {
		c.Config.Logger.Infof("Diff in SandboxConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SandboxConfig), dcl.SprintResource(actual.SandboxConfig))
		return true
	}
	if !dcl.StringCanonicalize(desired.NodeGroup, actual.NodeGroup) && !dcl.IsZeroValue(desired.NodeGroup) {
		c.Config.Logger.Infof("Diff in NodeGroup.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NodeGroup), dcl.SprintResource(actual.NodeGroup))
		return true
	}
	if compareClusterNodeConfigReservationAffinity(c, desired.ReservationAffinity, actual.ReservationAffinity) && !dcl.IsZeroValue(desired.ReservationAffinity) {
		c.Config.Logger.Infof("Diff in ReservationAffinity.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReservationAffinity), dcl.SprintResource(actual.ReservationAffinity))
		return true
	}
	if compareClusterNodeConfigShieldedInstanceConfig(c, desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig) && !dcl.IsZeroValue(desired.ShieldedInstanceConfig) {
		c.Config.Logger.Infof("Diff in ShieldedInstanceConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ShieldedInstanceConfig), dcl.SprintResource(actual.ShieldedInstanceConfig))
		return true
	}
	if compareClusterNodeConfigLinuxNodeConfig(c, desired.LinuxNodeConfig, actual.LinuxNodeConfig) && !dcl.IsZeroValue(desired.LinuxNodeConfig) {
		c.Config.Logger.Infof("Diff in LinuxNodeConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LinuxNodeConfig), dcl.SprintResource(actual.LinuxNodeConfig))
		return true
	}
	if compareClusterNodeConfigKubeletConfig(c, desired.KubeletConfig, actual.KubeletConfig) && !dcl.IsZeroValue(desired.KubeletConfig) {
		c.Config.Logger.Infof("Diff in KubeletConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KubeletConfig), dcl.SprintResource(actual.KubeletConfig))
		return true
	}
	if !dcl.StringCanonicalize(desired.BootDiskKmsKey, actual.BootDiskKmsKey) && !dcl.IsZeroValue(desired.BootDiskKmsKey) {
		c.Config.Logger.Infof("Diff in BootDiskKmsKey.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BootDiskKmsKey), dcl.SprintResource(actual.BootDiskKmsKey))
		return true
	}
	return false
}

func compareClusterNodeConfigSlice(c *Client, desired, actual []ClusterNodeConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigMap(c *Client, desired, actual map[string]ClusterNodeConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodeConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodeConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigAcceleratorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeConfigAccelerators)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigAccelerators or *ClusterNodeConfigAccelerators", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeConfigAccelerators)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeConfigAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigAccelerators", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AcceleratorCount, actual.AcceleratorCount, dcl.Info{}, fn.AddNest("AcceleratorCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorType, actual.AcceleratorType, dcl.Info{}, fn.AddNest("AcceleratorType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodeConfigAccelerators(c *Client, desired, actual *ClusterNodeConfigAccelerators) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.AcceleratorCount, actual.AcceleratorCount) && !dcl.IsZeroValue(desired.AcceleratorCount) {
		c.Config.Logger.Infof("Diff in AcceleratorCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorCount), dcl.SprintResource(actual.AcceleratorCount))
		return true
	}
	if !dcl.StringCanonicalize(desired.AcceleratorType, actual.AcceleratorType) && !dcl.IsZeroValue(desired.AcceleratorType) {
		c.Config.Logger.Infof("Diff in AcceleratorType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorType), dcl.SprintResource(actual.AcceleratorType))
		return true
	}
	return false
}

func compareClusterNodeConfigAcceleratorsSlice(c *Client, desired, actual []ClusterNodeConfigAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigAccelerators, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigAccelerators(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigAccelerators, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigAcceleratorsMap(c *Client, desired, actual map[string]ClusterNodeConfigAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigAccelerators, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigAccelerators, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodeConfigAccelerators(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigAccelerators, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigWorkloadMetadataConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeConfigWorkloadMetadataConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeConfigWorkloadMetadataConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigWorkloadMetadataConfig or *ClusterNodeConfigWorkloadMetadataConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeConfigWorkloadMetadataConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeConfigWorkloadMetadataConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigWorkloadMetadataConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{Type: "EnumType"}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodeConfigWorkloadMetadataConfig(c *Client, desired, actual *ClusterNodeConfigWorkloadMetadataConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Mode, actual.Mode) && !dcl.IsZeroValue(desired.Mode) {
		c.Config.Logger.Infof("Diff in Mode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Mode), dcl.SprintResource(actual.Mode))
		return true
	}
	return false
}

func compareClusterNodeConfigWorkloadMetadataConfigSlice(c *Client, desired, actual []ClusterNodeConfigWorkloadMetadataConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigWorkloadMetadataConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigWorkloadMetadataConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigWorkloadMetadataConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigWorkloadMetadataConfigMap(c *Client, desired, actual map[string]ClusterNodeConfigWorkloadMetadataConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigWorkloadMetadataConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigWorkloadMetadataConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodeConfigWorkloadMetadataConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigWorkloadMetadataConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigTaintsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeConfigTaints)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeConfigTaints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigTaints or *ClusterNodeConfigTaints", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeConfigTaints)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeConfigTaints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigTaints", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Effect, actual.Effect, dcl.Info{Type: "EnumType"}, fn.AddNest("Effect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodeConfigTaints(c *Client, desired, actual *ClusterNodeConfigTaints) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Effect, actual.Effect) && !dcl.IsZeroValue(desired.Effect) {
		c.Config.Logger.Infof("Diff in Effect.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Effect), dcl.SprintResource(actual.Effect))
		return true
	}
	return false
}

func compareClusterNodeConfigTaintsSlice(c *Client, desired, actual []ClusterNodeConfigTaints) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigTaints, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigTaints(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigTaints, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigTaintsMap(c *Client, desired, actual map[string]ClusterNodeConfigTaints) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigTaints, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigTaints, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodeConfigTaints(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigTaints, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigSandboxConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeConfigSandboxConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeConfigSandboxConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigSandboxConfig or *ClusterNodeConfigSandboxConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeConfigSandboxConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeConfigSandboxConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigSandboxConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType"}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodeConfigSandboxConfig(c *Client, desired, actual *ClusterNodeConfigSandboxConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}

func compareClusterNodeConfigSandboxConfigSlice(c *Client, desired, actual []ClusterNodeConfigSandboxConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigSandboxConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigSandboxConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigSandboxConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigSandboxConfigMap(c *Client, desired, actual map[string]ClusterNodeConfigSandboxConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigSandboxConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigSandboxConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodeConfigSandboxConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigSandboxConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigReservationAffinityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeConfigReservationAffinity)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeConfigReservationAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigReservationAffinity or *ClusterNodeConfigReservationAffinity", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeConfigReservationAffinity)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeConfigReservationAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigReservationAffinity", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConsumeReservationType, actual.ConsumeReservationType, dcl.Info{Type: "EnumType"}, fn.AddNest("ConsumeReservationType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Values, actual.Values, dcl.Info{}, fn.AddNest("Values")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodeConfigReservationAffinity(c *Client, desired, actual *ClusterNodeConfigReservationAffinity) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.ConsumeReservationType, actual.ConsumeReservationType) && !dcl.IsZeroValue(desired.ConsumeReservationType) {
		c.Config.Logger.Infof("Diff in ConsumeReservationType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConsumeReservationType), dcl.SprintResource(actual.ConsumeReservationType))
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if !dcl.StringSliceEquals(desired.Values, actual.Values) && !dcl.IsZeroValue(desired.Values) {
		c.Config.Logger.Infof("Diff in Values.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Values), dcl.SprintResource(actual.Values))
		return true
	}
	return false
}

func compareClusterNodeConfigReservationAffinitySlice(c *Client, desired, actual []ClusterNodeConfigReservationAffinity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigReservationAffinity, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigReservationAffinity(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigReservationAffinity, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigReservationAffinityMap(c *Client, desired, actual map[string]ClusterNodeConfigReservationAffinity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigReservationAffinity, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigReservationAffinity, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodeConfigReservationAffinity(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigReservationAffinity, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigShieldedInstanceConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeConfigShieldedInstanceConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeConfigShieldedInstanceConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigShieldedInstanceConfig or *ClusterNodeConfigShieldedInstanceConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeConfigShieldedInstanceConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeConfigShieldedInstanceConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigShieldedInstanceConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnableSecureBoot, actual.EnableSecureBoot, dcl.Info{}, fn.AddNest("EnableSecureBoot")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring, dcl.Info{}, fn.AddNest("EnableIntegrityMonitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodeConfigShieldedInstanceConfig(c *Client, desired, actual *ClusterNodeConfigShieldedInstanceConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableSecureBoot, actual.EnableSecureBoot) && !dcl.IsZeroValue(desired.EnableSecureBoot) {
		c.Config.Logger.Infof("Diff in EnableSecureBoot.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableSecureBoot), dcl.SprintResource(actual.EnableSecureBoot))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring) && !dcl.IsZeroValue(desired.EnableIntegrityMonitoring) {
		c.Config.Logger.Infof("Diff in EnableIntegrityMonitoring.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableIntegrityMonitoring), dcl.SprintResource(actual.EnableIntegrityMonitoring))
		return true
	}
	return false
}

func compareClusterNodeConfigShieldedInstanceConfigSlice(c *Client, desired, actual []ClusterNodeConfigShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigShieldedInstanceConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigShieldedInstanceConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigShieldedInstanceConfigMap(c *Client, desired, actual map[string]ClusterNodeConfigShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigShieldedInstanceConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodeConfigShieldedInstanceConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigShieldedInstanceConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigLinuxNodeConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeConfigLinuxNodeConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeConfigLinuxNodeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigLinuxNodeConfig or *ClusterNodeConfigLinuxNodeConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeConfigLinuxNodeConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeConfigLinuxNodeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigLinuxNodeConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Sysctls, actual.Sysctls, dcl.Info{}, fn.AddNest("Sysctls")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodeConfigLinuxNodeConfig(c *Client, desired, actual *ClusterNodeConfigLinuxNodeConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.MapEquals(desired.Sysctls, actual.Sysctls, []string(nil)) && !dcl.IsZeroValue(desired.Sysctls) {
		c.Config.Logger.Infof("Diff in Sysctls.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sysctls), dcl.SprintResource(actual.Sysctls))
		return true
	}
	return false
}

func compareClusterNodeConfigLinuxNodeConfigSlice(c *Client, desired, actual []ClusterNodeConfigLinuxNodeConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigLinuxNodeConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigLinuxNodeConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigLinuxNodeConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigLinuxNodeConfigMap(c *Client, desired, actual map[string]ClusterNodeConfigLinuxNodeConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigLinuxNodeConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigLinuxNodeConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodeConfigLinuxNodeConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigLinuxNodeConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigKubeletConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNodeConfigKubeletConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNodeConfigKubeletConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigKubeletConfig or *ClusterNodeConfigKubeletConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNodeConfigKubeletConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNodeConfigKubeletConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNodeConfigKubeletConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CpuManagerPolicy, actual.CpuManagerPolicy, dcl.Info{}, fn.AddNest("CpuManagerPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CpuCfsQuota, actual.CpuCfsQuota, dcl.Info{}, fn.AddNest("CpuCfsQuota")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CpuCfsQuotaPeriod, actual.CpuCfsQuotaPeriod, dcl.Info{}, fn.AddNest("CpuCfsQuotaPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNodeConfigKubeletConfig(c *Client, desired, actual *ClusterNodeConfigKubeletConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.CpuManagerPolicy, actual.CpuManagerPolicy) && !dcl.IsZeroValue(desired.CpuManagerPolicy) {
		c.Config.Logger.Infof("Diff in CpuManagerPolicy.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CpuManagerPolicy), dcl.SprintResource(actual.CpuManagerPolicy))
		return true
	}
	if !dcl.BoolCanonicalize(desired.CpuCfsQuota, actual.CpuCfsQuota) && !dcl.IsZeroValue(desired.CpuCfsQuota) {
		c.Config.Logger.Infof("Diff in CpuCfsQuota.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CpuCfsQuota), dcl.SprintResource(actual.CpuCfsQuota))
		return true
	}
	if !dcl.StringCanonicalize(desired.CpuCfsQuotaPeriod, actual.CpuCfsQuotaPeriod) && !dcl.IsZeroValue(desired.CpuCfsQuotaPeriod) {
		c.Config.Logger.Infof("Diff in CpuCfsQuotaPeriod.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CpuCfsQuotaPeriod), dcl.SprintResource(actual.CpuCfsQuotaPeriod))
		return true
	}
	return false
}

func compareClusterNodeConfigKubeletConfigSlice(c *Client, desired, actual []ClusterNodeConfigKubeletConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigKubeletConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigKubeletConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigKubeletConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigKubeletConfigMap(c *Client, desired, actual map[string]ClusterNodeConfigKubeletConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigKubeletConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigKubeletConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNodeConfigKubeletConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigKubeletConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterReleaseChannelNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterReleaseChannel)
	if !ok {
		desiredNotPointer, ok := d.(ClusterReleaseChannel)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterReleaseChannel or *ClusterReleaseChannel", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterReleaseChannel)
	if !ok {
		actualNotPointer, ok := a.(ClusterReleaseChannel)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterReleaseChannel", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Channel, actual.Channel, dcl.Info{Type: "EnumType"}, fn.AddNest("Channel")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterReleaseChannel(c *Client, desired, actual *ClusterReleaseChannel) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Channel, actual.Channel) && !dcl.IsZeroValue(desired.Channel) {
		c.Config.Logger.Infof("Diff in Channel.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Channel), dcl.SprintResource(actual.Channel))
		return true
	}
	return false
}

func compareClusterReleaseChannelSlice(c *Client, desired, actual []ClusterReleaseChannel) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterReleaseChannel, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterReleaseChannel(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterReleaseChannel, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterReleaseChannelMap(c *Client, desired, actual map[string]ClusterReleaseChannel) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterReleaseChannel, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterReleaseChannel, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterReleaseChannel(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterReleaseChannel, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterWorkloadIdentityConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterWorkloadIdentityConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterWorkloadIdentityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterWorkloadIdentityConfig or *ClusterWorkloadIdentityConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterWorkloadIdentityConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterWorkloadIdentityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterWorkloadIdentityConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WorkloadPool, actual.WorkloadPool, dcl.Info{}, fn.AddNest("WorkloadPool")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterWorkloadIdentityConfig(c *Client, desired, actual *ClusterWorkloadIdentityConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.WorkloadPool, actual.WorkloadPool) && !dcl.IsZeroValue(desired.WorkloadPool) {
		c.Config.Logger.Infof("Diff in WorkloadPool.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WorkloadPool), dcl.SprintResource(actual.WorkloadPool))
		return true
	}
	return false
}

func compareClusterWorkloadIdentityConfigSlice(c *Client, desired, actual []ClusterWorkloadIdentityConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterWorkloadIdentityConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterWorkloadIdentityConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterWorkloadIdentityConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterWorkloadIdentityConfigMap(c *Client, desired, actual map[string]ClusterWorkloadIdentityConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterWorkloadIdentityConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterWorkloadIdentityConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterWorkloadIdentityConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterWorkloadIdentityConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNotificationConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNotificationConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNotificationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNotificationConfig or *ClusterNotificationConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNotificationConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterNotificationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNotificationConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pubsub, actual.Pubsub, dcl.Info{ObjectFunction: compareClusterNotificationConfigPubsubNewStyle}, fn.AddNest("Pubsub")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNotificationConfig(c *Client, desired, actual *ClusterNotificationConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareClusterNotificationConfigPubsub(c, desired.Pubsub, actual.Pubsub) && !dcl.IsZeroValue(desired.Pubsub) {
		c.Config.Logger.Infof("Diff in Pubsub.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Pubsub), dcl.SprintResource(actual.Pubsub))
		return true
	}
	return false
}

func compareClusterNotificationConfigSlice(c *Client, desired, actual []ClusterNotificationConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNotificationConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNotificationConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNotificationConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNotificationConfigMap(c *Client, desired, actual map[string]ClusterNotificationConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNotificationConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNotificationConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNotificationConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNotificationConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterNotificationConfigPubsubNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNotificationConfigPubsub)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNotificationConfigPubsub)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNotificationConfigPubsub or *ClusterNotificationConfigPubsub", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNotificationConfigPubsub)
	if !ok {
		actualNotPointer, ok := a.(ClusterNotificationConfigPubsub)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNotificationConfigPubsub", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Topic, actual.Topic, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Topic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterNotificationConfigPubsub(c *Client, desired, actual *ClusterNotificationConfigPubsub) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	if !dcl.NameToSelfLink(desired.Topic, actual.Topic) && !dcl.IsZeroValue(desired.Topic) {
		c.Config.Logger.Infof("Diff in Topic.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Topic), dcl.SprintResource(actual.Topic))
		return true
	}
	return false
}

func compareClusterNotificationConfigPubsubSlice(c *Client, desired, actual []ClusterNotificationConfigPubsub) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNotificationConfigPubsub, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNotificationConfigPubsub(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNotificationConfigPubsub, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNotificationConfigPubsubMap(c *Client, desired, actual map[string]ClusterNotificationConfigPubsub) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNotificationConfigPubsub, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterNotificationConfigPubsub, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterNotificationConfigPubsub(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterNotificationConfigPubsub, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterConfidentialNodesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterConfidentialNodes)
	if !ok {
		desiredNotPointer, ok := d.(ClusterConfidentialNodes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfidentialNodes or *ClusterConfidentialNodes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterConfidentialNodes)
	if !ok {
		actualNotPointer, ok := a.(ClusterConfidentialNodes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterConfidentialNodes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterConfidentialNodes(c *Client, desired, actual *ClusterConfidentialNodes) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	return false
}

func compareClusterConfidentialNodesSlice(c *Client, desired, actual []ClusterConfidentialNodes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterConfidentialNodes, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterConfidentialNodes(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterConfidentialNodes, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterConfidentialNodesMap(c *Client, desired, actual map[string]ClusterConfidentialNodes) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterConfidentialNodes, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ClusterConfidentialNodes, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareClusterConfidentialNodes(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ClusterConfidentialNodes, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumSlice(c *Client, desired, actual []ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(c *Client, desired, actual *ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodePoolsConfigWorkloadMetadataConfigModeEnumSlice(c *Client, desired, actual []ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(c *Client, desired, actual *ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodePoolsConfigTaintsEffectEnumSlice(c *Client, desired, actual []ClusterNodePoolsConfigTaintsEffectEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigTaintsEffectEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigTaintsEffectEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigTaintsEffectEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigTaintsEffectEnum(c *Client, desired, actual *ClusterNodePoolsConfigTaintsEffectEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodePoolsConfigSandboxConfigTypeEnumSlice(c *Client, desired, actual []ClusterNodePoolsConfigSandboxConfigTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigSandboxConfigTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigSandboxConfigTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigSandboxConfigTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigSandboxConfigTypeEnum(c *Client, desired, actual *ClusterNodePoolsConfigSandboxConfigTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumSlice(c *Client, desired, actual []ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(c *Client, desired, actual *ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodePoolsStatusEnumSlice(c *Client, desired, actual []ClusterNodePoolsStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsStatusEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsStatusEnum(c *Client, desired, actual *ClusterNodePoolsStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodePoolsConditionsCodeEnumSlice(c *Client, desired, actual []ClusterNodePoolsConditionsCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConditionsCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConditionsCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConditionsCodeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConditionsCodeEnum(c *Client, desired, actual *ClusterNodePoolsConditionsCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodePoolsConditionsCanonicalCodeEnumSlice(c *Client, desired, actual []ClusterNodePoolsConditionsCanonicalCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodePoolsConditionsCanonicalCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodePoolsConditionsCanonicalCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodePoolsConditionsCanonicalCodeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePoolsConditionsCanonicalCodeEnum(c *Client, desired, actual *ClusterNodePoolsConditionsCanonicalCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNetworkPolicyProviderEnumSlice(c *Client, desired, actual []ClusterNetworkPolicyProviderEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNetworkPolicyProviderEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNetworkPolicyProviderEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNetworkPolicyProviderEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNetworkPolicyProviderEnum(c *Client, desired, actual *ClusterNetworkPolicyProviderEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNetworkConfigPrivateIPv6GoogleAccessEnumSlice(c *Client, desired, actual []ClusterNetworkConfigPrivateIPv6GoogleAccessEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNetworkConfigPrivateIPv6GoogleAccessEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNetworkConfigPrivateIPv6GoogleAccessEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNetworkConfigPrivateIPv6GoogleAccessEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNetworkConfigPrivateIPv6GoogleAccessEnum(c *Client, desired, actual *ClusterNetworkConfigPrivateIPv6GoogleAccessEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterDatabaseEncryptionStateEnumSlice(c *Client, desired, actual []ClusterDatabaseEncryptionStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterDatabaseEncryptionStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterDatabaseEncryptionStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterDatabaseEncryptionStateEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterDatabaseEncryptionStateEnum(c *Client, desired, actual *ClusterDatabaseEncryptionStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterConditionsCanonicalCodeEnumSlice(c *Client, desired, actual []ClusterConditionsCanonicalCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterConditionsCanonicalCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterConditionsCanonicalCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterConditionsCanonicalCodeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterConditionsCanonicalCodeEnum(c *Client, desired, actual *ClusterConditionsCanonicalCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodeConfigWorkloadMetadataConfigModeEnumSlice(c *Client, desired, actual []ClusterNodeConfigWorkloadMetadataConfigModeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigWorkloadMetadataConfigModeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigWorkloadMetadataConfigModeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigWorkloadMetadataConfigModeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigWorkloadMetadataConfigModeEnum(c *Client, desired, actual *ClusterNodeConfigWorkloadMetadataConfigModeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodeConfigTaintsEffectEnumSlice(c *Client, desired, actual []ClusterNodeConfigTaintsEffectEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigTaintsEffectEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigTaintsEffectEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigTaintsEffectEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigTaintsEffectEnum(c *Client, desired, actual *ClusterNodeConfigTaintsEffectEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodeConfigSandboxConfigTypeEnumSlice(c *Client, desired, actual []ClusterNodeConfigSandboxConfigTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigSandboxConfigTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigSandboxConfigTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigSandboxConfigTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigSandboxConfigTypeEnum(c *Client, desired, actual *ClusterNodeConfigSandboxConfigTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterNodeConfigReservationAffinityConsumeReservationTypeEnumSlice(c *Client, desired, actual []ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(c *Client, desired, actual *ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterReleaseChannelChannelEnumSlice(c *Client, desired, actual []ClusterReleaseChannelChannelEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterReleaseChannelChannelEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterReleaseChannelChannelEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterReleaseChannelChannelEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterReleaseChannelChannelEnum(c *Client, desired, actual *ClusterReleaseChannelChannelEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Cluster) urlNormalized() *Cluster {
	normalized := dcl.Copy(*r).(Cluster)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.LoggingService = dcl.SelfLinkToName(r.LoggingService)
	normalized.MonitoringService = dcl.SelfLinkToName(r.MonitoringService)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.ClusterIPv4Cidr = dcl.SelfLinkToName(r.ClusterIPv4Cidr)
	normalized.Subnetwork = dcl.SelfLinkToName(r.Subnetwork)
	normalized.LabelFingerprint = dcl.SelfLinkToName(r.LabelFingerprint)
	normalized.Endpoint = dcl.SelfLinkToName(r.Endpoint)
	normalized.MasterVersion = dcl.SelfLinkToName(r.MasterVersion)
	normalized.StatusMessage = dcl.SelfLinkToName(r.StatusMessage)
	normalized.ServicesIPv4Cidr = dcl.SelfLinkToName(r.ServicesIPv4Cidr)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.TPUIPv4CidrBlock = dcl.SelfLinkToName(r.TPUIPv4CidrBlock)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.InitialClusterVersion = dcl.SelfLinkToName(r.InitialClusterVersion)
	normalized.CurrentMasterVersion = dcl.SelfLinkToName(r.CurrentMasterVersion)
	normalized.CurrentNodeVersion = dcl.SelfLinkToName(r.CurrentNodeVersion)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	return &normalized
}

func (r *Cluster) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Cluster) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *Cluster) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Cluster) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "setMaintenancePolicy" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}:setMaintenancePolicy", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateAddonsConfig" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateAutoscaling" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateBinaryAuthorization" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateDatabaseEncryption" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateLegacyAbac" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}:setLegacyAbac", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateLocations" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateMasterAuthorizedNetworksConfig" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateMasterVersion" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateMonitoringAndLoggingService" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateShieldedNodes" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateVerticalPodAutoscaling" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "updateWorkloadIdentityConfig" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Cluster resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Cluster) marshal(c *Client) ([]byte, error) {
	m, err := expandCluster(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Cluster: %w", err)
	}
	m = EncodeClusterCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalCluster decodes JSON responses into the Cluster resource schema.
func unmarshalCluster(b []byte, c *Client) (*Cluster, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapCluster(m, c)
}

func unmarshalMapCluster(m map[string]interface{}, c *Client) (*Cluster, error) {

	return flattenCluster(c, m), nil
}

// expandCluster expands Cluster into a JSON request object.
func expandCluster(c *Client, f *Cluster) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.InitialNodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["initialNodeCount"] = v
	}
	if v, err := expandClusterMasterAuth(c, f.MasterAuth); err != nil {
		return nil, fmt.Errorf("error expanding MasterAuth into masterAuth: %w", err)
	} else if v != nil {
		m["masterAuth"] = v
	}
	if v := f.LoggingService; !dcl.IsEmptyValueIndirect(v) {
		m["loggingService"] = v
	}
	if v := f.MonitoringService; !dcl.IsEmptyValueIndirect(v) {
		m["monitoringService"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.ClusterIPv4Cidr; !dcl.IsEmptyValueIndirect(v) {
		m["clusterIpv4Cidr"] = v
	}
	if v, err := expandClusterAddonsConfig(c, f.AddonsConfig); err != nil {
		return nil, fmt.Errorf("error expanding AddonsConfig into addonsConfig: %w", err)
	} else if v != nil {
		m["addonsConfig"] = v
	}
	if v := f.Subnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["subnetwork"] = v
	}
	if v, err := expandClusterNodePoolsSlice(c, f.NodePools); err != nil {
		return nil, fmt.Errorf("error expanding NodePools into nodePools: %w", err)
	} else if v != nil {
		m["nodePools"] = v
	}
	if v := f.Locations; !dcl.IsEmptyValueIndirect(v) {
		m["locations"] = v
	}
	if v := f.EnableKubernetesAlpha; !dcl.IsEmptyValueIndirect(v) {
		m["enableKubernetesAlpha"] = v
	}
	if v := f.ResourceLabels; !dcl.IsEmptyValueIndirect(v) {
		m["resourceLabels"] = v
	}
	if v := f.LabelFingerprint; !dcl.IsEmptyValueIndirect(v) {
		m["labelFingerprint"] = v
	}
	if v, err := expandClusterLegacyAbac(c, f.LegacyAbac); err != nil {
		return nil, fmt.Errorf("error expanding LegacyAbac into legacyAbac: %w", err)
	} else if v != nil {
		m["legacyAbac"] = v
	}
	if v, err := expandClusterNetworkPolicy(c, f.NetworkPolicy); err != nil {
		return nil, fmt.Errorf("error expanding NetworkPolicy into networkPolicy: %w", err)
	} else if v != nil {
		m["networkPolicy"] = v
	}
	if v, err := expandClusterIPAllocationPolicy(c, f.IPAllocationPolicy); err != nil {
		return nil, fmt.Errorf("error expanding IPAllocationPolicy into ipAllocationPolicy: %w", err)
	} else if v != nil {
		m["ipAllocationPolicy"] = v
	}
	if v, err := expandClusterMasterAuthorizedNetworksConfig(c, f.MasterAuthorizedNetworksConfig); err != nil {
		return nil, fmt.Errorf("error expanding MasterAuthorizedNetworksConfig into masterAuthorizedNetworksConfig: %w", err)
	} else if v != nil {
		m["masterAuthorizedNetworksConfig"] = v
	}
	if v, err := expandClusterBinaryAuthorization(c, f.BinaryAuthorization); err != nil {
		return nil, fmt.Errorf("error expanding BinaryAuthorization into binaryAuthorization: %w", err)
	} else if v != nil {
		m["binaryAuthorization"] = v
	}
	if v, err := expandClusterAutoscaling(c, f.Autoscaling); err != nil {
		return nil, fmt.Errorf("error expanding Autoscaling into autoscaling: %w", err)
	} else if v != nil {
		m["autoscaling"] = v
	}
	if v, err := expandClusterNetworkConfig(c, f.NetworkConfig); err != nil {
		return nil, fmt.Errorf("error expanding NetworkConfig into networkConfig: %w", err)
	} else if v != nil {
		m["networkConfig"] = v
	}
	if v, err := expandClusterMaintenancePolicy(c, f.MaintenancePolicy); err != nil {
		return nil, fmt.Errorf("error expanding MaintenancePolicy into maintenancePolicy: %w", err)
	} else if v != nil {
		m["maintenancePolicy"] = v
	}
	if v, err := expandClusterDefaultMaxPodsConstraint(c, f.DefaultMaxPodsConstraint); err != nil {
		return nil, fmt.Errorf("error expanding DefaultMaxPodsConstraint into defaultMaxPodsConstraint: %w", err)
	} else if v != nil {
		m["defaultMaxPodsConstraint"] = v
	}
	if v, err := expandClusterResourceUsageExportConfig(c, f.ResourceUsageExportConfig); err != nil {
		return nil, fmt.Errorf("error expanding ResourceUsageExportConfig into resourceUsageExportConfig: %w", err)
	} else if v != nil {
		m["resourceUsageExportConfig"] = v
	}
	if v, err := expandClusterAuthenticatorGroupsConfig(c, f.AuthenticatorGroupsConfig); err != nil {
		return nil, fmt.Errorf("error expanding AuthenticatorGroupsConfig into authenticatorGroupsConfig: %w", err)
	} else if v != nil {
		m["authenticatorGroupsConfig"] = v
	}
	if v, err := expandClusterPrivateClusterConfig(c, f.PrivateClusterConfig); err != nil {
		return nil, fmt.Errorf("error expanding PrivateClusterConfig into privateClusterConfig: %w", err)
	} else if v != nil {
		m["privateClusterConfig"] = v
	}
	if v, err := expandClusterDatabaseEncryption(c, f.DatabaseEncryption); err != nil {
		return nil, fmt.Errorf("error expanding DatabaseEncryption into databaseEncryption: %w", err)
	} else if v != nil {
		m["databaseEncryption"] = v
	}
	if v, err := expandClusterVerticalPodAutoscaling(c, f.VerticalPodAutoscaling); err != nil {
		return nil, fmt.Errorf("error expanding VerticalPodAutoscaling into verticalPodAutoscaling: %w", err)
	} else if v != nil {
		m["verticalPodAutoscaling"] = v
	}
	if v, err := expandClusterShieldedNodes(c, f.ShieldedNodes); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedNodes into shieldedNodes: %w", err)
	} else if v != nil {
		m["shieldedNodes"] = v
	}
	if v := f.Endpoint; !dcl.IsEmptyValueIndirect(v) {
		m["endpoint"] = v
	}
	if v := f.MasterVersion; !dcl.IsEmptyValueIndirect(v) {
		m["currentMasterVersion"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.StatusMessage; !dcl.IsEmptyValueIndirect(v) {
		m["statusMessage"] = v
	}
	if v := f.NodeIPv4CidrSize; !dcl.IsEmptyValueIndirect(v) {
		m["nodeIpv4CidrSize"] = v
	}
	if v := f.ServicesIPv4Cidr; !dcl.IsEmptyValueIndirect(v) {
		m["servicesIpv4Cidr"] = v
	}
	if v := f.ExpireTime; !dcl.IsEmptyValueIndirect(v) {
		m["expireTime"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.EnableTPU; !dcl.IsEmptyValueIndirect(v) {
		m["enableTpu"] = v
	}
	if v := f.TPUIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["tpuIpv4CidrBlock"] = v
	}
	if v, err := expandClusterConditionsSlice(c, f.Conditions); err != nil {
		return nil, fmt.Errorf("error expanding Conditions into conditions: %w", err)
	} else if v != nil {
		m["conditions"] = v
	}
	if v, err := expandClusterAutopilot(c, f.Autopilot); err != nil {
		return nil, fmt.Errorf("error expanding Autopilot into autopilot: %w", err)
	} else if v != nil {
		m["autopilot"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := expandClusterNodeConfig(c, f.NodeConfig); err != nil {
		return nil, fmt.Errorf("error expanding NodeConfig into nodeConfig: %w", err)
	} else if v != nil {
		m["nodeConfig"] = v
	}
	if v, err := expandClusterReleaseChannel(c, f.ReleaseChannel); err != nil {
		return nil, fmt.Errorf("error expanding ReleaseChannel into releaseChannel: %w", err)
	} else if v != nil {
		m["releaseChannel"] = v
	}
	if v, err := expandClusterWorkloadIdentityConfig(c, f.WorkloadIdentityConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkloadIdentityConfig into workloadIdentityConfig: %w", err)
	} else if v != nil {
		m["workloadIdentityConfig"] = v
	}
	if v, err := expandClusterNotificationConfig(c, f.NotificationConfig); err != nil {
		return nil, fmt.Errorf("error expanding NotificationConfig into notificationConfig: %w", err)
	} else if v != nil {
		m["notificationConfig"] = v
	}
	if v, err := expandClusterConfidentialNodes(c, f.ConfidentialNodes); err != nil {
		return nil, fmt.Errorf("error expanding ConfidentialNodes into confidentialNodes: %w", err)
	} else if v != nil {
		m["confidentialNodes"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v := f.InitialClusterVersion; !dcl.IsEmptyValueIndirect(v) {
		m["initialClusterVersion"] = v
	}
	if v := f.CurrentMasterVersion; !dcl.IsEmptyValueIndirect(v) {
		m["currentMasterVersion"] = v
	}
	if v := f.CurrentNodeVersion; !dcl.IsEmptyValueIndirect(v) {
		m["currentNodeVersion"] = v
	}
	if v := f.InstanceGroupUrls; !dcl.IsEmptyValueIndirect(v) {
		m["instanceGroupUrls"] = v
	}
	if v := f.CurrentNodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["currentNodeCount"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}

	return m, nil
}

// flattenCluster flattens Cluster from a JSON request object into the
// Cluster type.
func flattenCluster(c *Client, i interface{}) *Cluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Cluster{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.InitialNodeCount = dcl.FlattenInteger(m["initialNodeCount"])
	r.MasterAuth = flattenClusterMasterAuth(c, m["masterAuth"])
	r.LoggingService = dcl.FlattenString(m["loggingService"])
	r.MonitoringService = dcl.FlattenString(m["monitoringService"])
	r.Network = dcl.FlattenString(m["network"])
	r.ClusterIPv4Cidr = dcl.FlattenString(m["clusterIpv4Cidr"])
	r.AddonsConfig = flattenClusterAddonsConfig(c, m["addonsConfig"])
	r.Subnetwork = dcl.FlattenString(m["subnetwork"])
	r.NodePools = flattenClusterNodePoolsSlice(c, m["nodePools"])
	r.Locations = dcl.FlattenStringSlice(m["locations"])
	r.EnableKubernetesAlpha = dcl.FlattenBool(m["enableKubernetesAlpha"])
	r.ResourceLabels = dcl.FlattenKeyValuePairs(m["resourceLabels"])
	r.LabelFingerprint = dcl.FlattenString(m["labelFingerprint"])
	r.LegacyAbac = flattenClusterLegacyAbac(c, m["legacyAbac"])
	r.NetworkPolicy = flattenClusterNetworkPolicy(c, m["networkPolicy"])
	r.IPAllocationPolicy = flattenClusterIPAllocationPolicy(c, m["ipAllocationPolicy"])
	r.MasterAuthorizedNetworksConfig = flattenClusterMasterAuthorizedNetworksConfig(c, m["masterAuthorizedNetworksConfig"])
	r.BinaryAuthorization = flattenClusterBinaryAuthorization(c, m["binaryAuthorization"])
	r.Autoscaling = flattenClusterAutoscaling(c, m["autoscaling"])
	r.NetworkConfig = flattenClusterNetworkConfig(c, m["networkConfig"])
	r.MaintenancePolicy = flattenClusterMaintenancePolicy(c, m["maintenancePolicy"])
	r.DefaultMaxPodsConstraint = flattenClusterDefaultMaxPodsConstraint(c, m["defaultMaxPodsConstraint"])
	r.ResourceUsageExportConfig = flattenClusterResourceUsageExportConfig(c, m["resourceUsageExportConfig"])
	r.AuthenticatorGroupsConfig = flattenClusterAuthenticatorGroupsConfig(c, m["authenticatorGroupsConfig"])
	r.PrivateClusterConfig = flattenClusterPrivateClusterConfig(c, m["privateClusterConfig"])
	r.DatabaseEncryption = flattenClusterDatabaseEncryption(c, m["databaseEncryption"])
	r.VerticalPodAutoscaling = flattenClusterVerticalPodAutoscaling(c, m["verticalPodAutoscaling"])
	r.ShieldedNodes = flattenClusterShieldedNodes(c, m["shieldedNodes"])
	r.Endpoint = dcl.FlattenString(m["endpoint"])
	r.MasterVersion = dcl.FlattenString(m["currentMasterVersion"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.Status = dcl.FlattenString(m["status"])
	r.StatusMessage = dcl.FlattenString(m["statusMessage"])
	r.NodeIPv4CidrSize = dcl.FlattenInteger(m["nodeIpv4CidrSize"])
	r.ServicesIPv4Cidr = dcl.FlattenString(m["servicesIpv4Cidr"])
	r.ExpireTime = dcl.FlattenString(m["expireTime"])
	r.Location = dcl.FlattenString(m["location"])
	r.EnableTPU = dcl.FlattenBool(m["enableTpu"])
	r.TPUIPv4CidrBlock = dcl.FlattenString(m["tpuIpv4CidrBlock"])
	r.Conditions = flattenClusterConditionsSlice(c, m["conditions"])
	r.Autopilot = flattenClusterAutopilot(c, m["autopilot"])
	r.Project = dcl.FlattenString(m["project"])
	r.NodeConfig = flattenClusterNodeConfig(c, m["nodeConfig"])
	r.ReleaseChannel = flattenClusterReleaseChannel(c, m["releaseChannel"])
	r.WorkloadIdentityConfig = flattenClusterWorkloadIdentityConfig(c, m["workloadIdentityConfig"])
	r.NotificationConfig = flattenClusterNotificationConfig(c, m["notificationConfig"])
	r.ConfidentialNodes = flattenClusterConfidentialNodes(c, m["confidentialNodes"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Zone = dcl.FlattenString(m["zone"])
	r.InitialClusterVersion = dcl.FlattenString(m["initialClusterVersion"])
	r.CurrentMasterVersion = dcl.FlattenString(m["currentMasterVersion"])
	r.CurrentNodeVersion = dcl.FlattenString(m["currentNodeVersion"])
	r.InstanceGroupUrls = dcl.FlattenStringSlice(m["instanceGroupUrls"])
	r.CurrentNodeCount = dcl.FlattenInteger(m["currentNodeCount"])
	r.Id = dcl.FlattenString(m["id"])

	return r
}

// expandClusterMasterAuthMap expands the contents of ClusterMasterAuth into a JSON
// request object.
func expandClusterMasterAuthMap(c *Client, f map[string]ClusterMasterAuth) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMasterAuth(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMasterAuthSlice expands the contents of ClusterMasterAuth into a JSON
// request object.
func expandClusterMasterAuthSlice(c *Client, f []ClusterMasterAuth) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMasterAuth(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMasterAuthMap flattens the contents of ClusterMasterAuth from a JSON
// response object.
func flattenClusterMasterAuthMap(c *Client, i interface{}) map[string]ClusterMasterAuth {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMasterAuth{}
	}

	if len(a) == 0 {
		return map[string]ClusterMasterAuth{}
	}

	items := make(map[string]ClusterMasterAuth)
	for k, item := range a {
		items[k] = *flattenClusterMasterAuth(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMasterAuthSlice flattens the contents of ClusterMasterAuth from a JSON
// response object.
func flattenClusterMasterAuthSlice(c *Client, i interface{}) []ClusterMasterAuth {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMasterAuth{}
	}

	if len(a) == 0 {
		return []ClusterMasterAuth{}
	}

	items := make([]ClusterMasterAuth, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMasterAuth(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMasterAuth expands an instance of ClusterMasterAuth into a JSON
// request object.
func expandClusterMasterAuth(c *Client, f *ClusterMasterAuth) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}
	if v, err := expandClusterMasterAuthClientCertificateConfig(c, f.ClientCertificateConfig); err != nil {
		return nil, fmt.Errorf("error expanding ClientCertificateConfig into clientCertificateConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["clientCertificateConfig"] = v
	}
	if v := f.ClusterCaCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["clusterCaCertificate"] = v
	}
	if v := f.ClientCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["clientCertificate"] = v
	}
	if v := f.ClientKey; !dcl.IsEmptyValueIndirect(v) {
		m["clientKey"] = v
	}

	return m, nil
}

// flattenClusterMasterAuth flattens an instance of ClusterMasterAuth from a JSON
// response object.
func flattenClusterMasterAuth(c *Client, i interface{}) *ClusterMasterAuth {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMasterAuth{}
	r.Username = dcl.FlattenString(m["username"])
	r.Password = dcl.FlattenString(m["password"])
	r.ClientCertificateConfig = flattenClusterMasterAuthClientCertificateConfig(c, m["clientCertificateConfig"])
	r.ClusterCaCertificate = dcl.FlattenString(m["clusterCaCertificate"])
	r.ClientCertificate = dcl.FlattenString(m["clientCertificate"])
	r.ClientKey = dcl.FlattenString(m["clientKey"])

	return r
}

// expandClusterMasterAuthClientCertificateConfigMap expands the contents of ClusterMasterAuthClientCertificateConfig into a JSON
// request object.
func expandClusterMasterAuthClientCertificateConfigMap(c *Client, f map[string]ClusterMasterAuthClientCertificateConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMasterAuthClientCertificateConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMasterAuthClientCertificateConfigSlice expands the contents of ClusterMasterAuthClientCertificateConfig into a JSON
// request object.
func expandClusterMasterAuthClientCertificateConfigSlice(c *Client, f []ClusterMasterAuthClientCertificateConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMasterAuthClientCertificateConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMasterAuthClientCertificateConfigMap flattens the contents of ClusterMasterAuthClientCertificateConfig from a JSON
// response object.
func flattenClusterMasterAuthClientCertificateConfigMap(c *Client, i interface{}) map[string]ClusterMasterAuthClientCertificateConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMasterAuthClientCertificateConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterMasterAuthClientCertificateConfig{}
	}

	items := make(map[string]ClusterMasterAuthClientCertificateConfig)
	for k, item := range a {
		items[k] = *flattenClusterMasterAuthClientCertificateConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMasterAuthClientCertificateConfigSlice flattens the contents of ClusterMasterAuthClientCertificateConfig from a JSON
// response object.
func flattenClusterMasterAuthClientCertificateConfigSlice(c *Client, i interface{}) []ClusterMasterAuthClientCertificateConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMasterAuthClientCertificateConfig{}
	}

	if len(a) == 0 {
		return []ClusterMasterAuthClientCertificateConfig{}
	}

	items := make([]ClusterMasterAuthClientCertificateConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMasterAuthClientCertificateConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMasterAuthClientCertificateConfig expands an instance of ClusterMasterAuthClientCertificateConfig into a JSON
// request object.
func expandClusterMasterAuthClientCertificateConfig(c *Client, f *ClusterMasterAuthClientCertificateConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.IssueClientCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["issueClientCertificate"] = v
	}

	return m, nil
}

// flattenClusterMasterAuthClientCertificateConfig flattens an instance of ClusterMasterAuthClientCertificateConfig from a JSON
// response object.
func flattenClusterMasterAuthClientCertificateConfig(c *Client, i interface{}) *ClusterMasterAuthClientCertificateConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMasterAuthClientCertificateConfig{}
	r.IssueClientCertificate = dcl.FlattenBool(m["issueClientCertificate"])

	return r
}

// expandClusterAddonsConfigMap expands the contents of ClusterAddonsConfig into a JSON
// request object.
func expandClusterAddonsConfigMap(c *Client, f map[string]ClusterAddonsConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAddonsConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAddonsConfigSlice expands the contents of ClusterAddonsConfig into a JSON
// request object.
func expandClusterAddonsConfigSlice(c *Client, f []ClusterAddonsConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAddonsConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAddonsConfigMap flattens the contents of ClusterAddonsConfig from a JSON
// response object.
func flattenClusterAddonsConfigMap(c *Client, i interface{}) map[string]ClusterAddonsConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAddonsConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterAddonsConfig{}
	}

	items := make(map[string]ClusterAddonsConfig)
	for k, item := range a {
		items[k] = *flattenClusterAddonsConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAddonsConfigSlice flattens the contents of ClusterAddonsConfig from a JSON
// response object.
func flattenClusterAddonsConfigSlice(c *Client, i interface{}) []ClusterAddonsConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfig{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfig{}
	}

	items := make([]ClusterAddonsConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAddonsConfig expands an instance of ClusterAddonsConfig into a JSON
// request object.
func expandClusterAddonsConfig(c *Client, f *ClusterAddonsConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandClusterAddonsConfigHttpLoadBalancing(c, f.HttpLoadBalancing); err != nil {
		return nil, fmt.Errorf("error expanding HttpLoadBalancing into httpLoadBalancing: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpLoadBalancing"] = v
	}
	if v, err := expandClusterAddonsConfigHorizontalPodAutoscaling(c, f.HorizontalPodAutoscaling); err != nil {
		return nil, fmt.Errorf("error expanding HorizontalPodAutoscaling into horizontalPodAutoscaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["horizontalPodAutoscaling"] = v
	}
	if v, err := expandClusterAddonsConfigKubernetesDashboard(c, f.KubernetesDashboard); err != nil {
		return nil, fmt.Errorf("error expanding KubernetesDashboard into kubernetesDashboard: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kubernetesDashboard"] = v
	}
	if v, err := expandClusterAddonsConfigNetworkPolicyConfig(c, f.NetworkPolicyConfig); err != nil {
		return nil, fmt.Errorf("error expanding NetworkPolicyConfig into networkPolicyConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["networkPolicyConfig"] = v
	}
	if v, err := expandClusterAddonsConfigCloudRunConfig(c, f.CloudRunConfig); err != nil {
		return nil, fmt.Errorf("error expanding CloudRunConfig into cloudRunConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudRunConfig"] = v
	}
	if v, err := expandClusterAddonsConfigDnsCacheConfig(c, f.DnsCacheConfig); err != nil {
		return nil, fmt.Errorf("error expanding DnsCacheConfig into dnsCacheConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dnsCacheConfig"] = v
	}
	if v, err := expandClusterAddonsConfigConfigConnectorConfig(c, f.ConfigConnectorConfig); err != nil {
		return nil, fmt.Errorf("error expanding ConfigConnectorConfig into configConnectorConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["configConnectorConfig"] = v
	}
	if v, err := expandClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, f.GcePersistentDiskCsiDriverConfig); err != nil {
		return nil, fmt.Errorf("error expanding GcePersistentDiskCsiDriverConfig into gcePersistentDiskCsiDriverConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gcePersistentDiskCsiDriverConfig"] = v
	}

	return m, nil
}

// flattenClusterAddonsConfig flattens an instance of ClusterAddonsConfig from a JSON
// response object.
func flattenClusterAddonsConfig(c *Client, i interface{}) *ClusterAddonsConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAddonsConfig{}
	r.HttpLoadBalancing = flattenClusterAddonsConfigHttpLoadBalancing(c, m["httpLoadBalancing"])
	r.HorizontalPodAutoscaling = flattenClusterAddonsConfigHorizontalPodAutoscaling(c, m["horizontalPodAutoscaling"])
	r.KubernetesDashboard = flattenClusterAddonsConfigKubernetesDashboard(c, m["kubernetesDashboard"])
	r.NetworkPolicyConfig = flattenClusterAddonsConfigNetworkPolicyConfig(c, m["networkPolicyConfig"])
	r.CloudRunConfig = flattenClusterAddonsConfigCloudRunConfig(c, m["cloudRunConfig"])
	r.DnsCacheConfig = flattenClusterAddonsConfigDnsCacheConfig(c, m["dnsCacheConfig"])
	r.ConfigConnectorConfig = flattenClusterAddonsConfigConfigConnectorConfig(c, m["configConnectorConfig"])
	r.GcePersistentDiskCsiDriverConfig = flattenClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, m["gcePersistentDiskCsiDriverConfig"])

	return r
}

// expandClusterAddonsConfigHttpLoadBalancingMap expands the contents of ClusterAddonsConfigHttpLoadBalancing into a JSON
// request object.
func expandClusterAddonsConfigHttpLoadBalancingMap(c *Client, f map[string]ClusterAddonsConfigHttpLoadBalancing) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAddonsConfigHttpLoadBalancing(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAddonsConfigHttpLoadBalancingSlice expands the contents of ClusterAddonsConfigHttpLoadBalancing into a JSON
// request object.
func expandClusterAddonsConfigHttpLoadBalancingSlice(c *Client, f []ClusterAddonsConfigHttpLoadBalancing) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAddonsConfigHttpLoadBalancing(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAddonsConfigHttpLoadBalancingMap flattens the contents of ClusterAddonsConfigHttpLoadBalancing from a JSON
// response object.
func flattenClusterAddonsConfigHttpLoadBalancingMap(c *Client, i interface{}) map[string]ClusterAddonsConfigHttpLoadBalancing {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAddonsConfigHttpLoadBalancing{}
	}

	if len(a) == 0 {
		return map[string]ClusterAddonsConfigHttpLoadBalancing{}
	}

	items := make(map[string]ClusterAddonsConfigHttpLoadBalancing)
	for k, item := range a {
		items[k] = *flattenClusterAddonsConfigHttpLoadBalancing(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAddonsConfigHttpLoadBalancingSlice flattens the contents of ClusterAddonsConfigHttpLoadBalancing from a JSON
// response object.
func flattenClusterAddonsConfigHttpLoadBalancingSlice(c *Client, i interface{}) []ClusterAddonsConfigHttpLoadBalancing {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfigHttpLoadBalancing{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfigHttpLoadBalancing{}
	}

	items := make([]ClusterAddonsConfigHttpLoadBalancing, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfigHttpLoadBalancing(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAddonsConfigHttpLoadBalancing expands an instance of ClusterAddonsConfigHttpLoadBalancing into a JSON
// request object.
func expandClusterAddonsConfigHttpLoadBalancing(c *Client, f *ClusterAddonsConfigHttpLoadBalancing) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	m["disabled"] = f.Disabled

	return m, nil
}

// flattenClusterAddonsConfigHttpLoadBalancing flattens an instance of ClusterAddonsConfigHttpLoadBalancing from a JSON
// response object.
func flattenClusterAddonsConfigHttpLoadBalancing(c *Client, i interface{}) *ClusterAddonsConfigHttpLoadBalancing {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAddonsConfigHttpLoadBalancing{}
	r.Disabled = dcl.FlattenBool(m["disabled"])

	return r
}

// expandClusterAddonsConfigHorizontalPodAutoscalingMap expands the contents of ClusterAddonsConfigHorizontalPodAutoscaling into a JSON
// request object.
func expandClusterAddonsConfigHorizontalPodAutoscalingMap(c *Client, f map[string]ClusterAddonsConfigHorizontalPodAutoscaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAddonsConfigHorizontalPodAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAddonsConfigHorizontalPodAutoscalingSlice expands the contents of ClusterAddonsConfigHorizontalPodAutoscaling into a JSON
// request object.
func expandClusterAddonsConfigHorizontalPodAutoscalingSlice(c *Client, f []ClusterAddonsConfigHorizontalPodAutoscaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAddonsConfigHorizontalPodAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAddonsConfigHorizontalPodAutoscalingMap flattens the contents of ClusterAddonsConfigHorizontalPodAutoscaling from a JSON
// response object.
func flattenClusterAddonsConfigHorizontalPodAutoscalingMap(c *Client, i interface{}) map[string]ClusterAddonsConfigHorizontalPodAutoscaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAddonsConfigHorizontalPodAutoscaling{}
	}

	if len(a) == 0 {
		return map[string]ClusterAddonsConfigHorizontalPodAutoscaling{}
	}

	items := make(map[string]ClusterAddonsConfigHorizontalPodAutoscaling)
	for k, item := range a {
		items[k] = *flattenClusterAddonsConfigHorizontalPodAutoscaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAddonsConfigHorizontalPodAutoscalingSlice flattens the contents of ClusterAddonsConfigHorizontalPodAutoscaling from a JSON
// response object.
func flattenClusterAddonsConfigHorizontalPodAutoscalingSlice(c *Client, i interface{}) []ClusterAddonsConfigHorizontalPodAutoscaling {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfigHorizontalPodAutoscaling{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfigHorizontalPodAutoscaling{}
	}

	items := make([]ClusterAddonsConfigHorizontalPodAutoscaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfigHorizontalPodAutoscaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAddonsConfigHorizontalPodAutoscaling expands an instance of ClusterAddonsConfigHorizontalPodAutoscaling into a JSON
// request object.
func expandClusterAddonsConfigHorizontalPodAutoscaling(c *Client, f *ClusterAddonsConfigHorizontalPodAutoscaling) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	m["disabled"] = f.Disabled

	return m, nil
}

// flattenClusterAddonsConfigHorizontalPodAutoscaling flattens an instance of ClusterAddonsConfigHorizontalPodAutoscaling from a JSON
// response object.
func flattenClusterAddonsConfigHorizontalPodAutoscaling(c *Client, i interface{}) *ClusterAddonsConfigHorizontalPodAutoscaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAddonsConfigHorizontalPodAutoscaling{}
	r.Disabled = dcl.FlattenBool(m["disabled"])

	return r
}

// expandClusterAddonsConfigKubernetesDashboardMap expands the contents of ClusterAddonsConfigKubernetesDashboard into a JSON
// request object.
func expandClusterAddonsConfigKubernetesDashboardMap(c *Client, f map[string]ClusterAddonsConfigKubernetesDashboard) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAddonsConfigKubernetesDashboard(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAddonsConfigKubernetesDashboardSlice expands the contents of ClusterAddonsConfigKubernetesDashboard into a JSON
// request object.
func expandClusterAddonsConfigKubernetesDashboardSlice(c *Client, f []ClusterAddonsConfigKubernetesDashboard) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAddonsConfigKubernetesDashboard(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAddonsConfigKubernetesDashboardMap flattens the contents of ClusterAddonsConfigKubernetesDashboard from a JSON
// response object.
func flattenClusterAddonsConfigKubernetesDashboardMap(c *Client, i interface{}) map[string]ClusterAddonsConfigKubernetesDashboard {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAddonsConfigKubernetesDashboard{}
	}

	if len(a) == 0 {
		return map[string]ClusterAddonsConfigKubernetesDashboard{}
	}

	items := make(map[string]ClusterAddonsConfigKubernetesDashboard)
	for k, item := range a {
		items[k] = *flattenClusterAddonsConfigKubernetesDashboard(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAddonsConfigKubernetesDashboardSlice flattens the contents of ClusterAddonsConfigKubernetesDashboard from a JSON
// response object.
func flattenClusterAddonsConfigKubernetesDashboardSlice(c *Client, i interface{}) []ClusterAddonsConfigKubernetesDashboard {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfigKubernetesDashboard{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfigKubernetesDashboard{}
	}

	items := make([]ClusterAddonsConfigKubernetesDashboard, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfigKubernetesDashboard(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAddonsConfigKubernetesDashboard expands an instance of ClusterAddonsConfigKubernetesDashboard into a JSON
// request object.
func expandClusterAddonsConfigKubernetesDashboard(c *Client, f *ClusterAddonsConfigKubernetesDashboard) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	m["disabled"] = f.Disabled

	return m, nil
}

// flattenClusterAddonsConfigKubernetesDashboard flattens an instance of ClusterAddonsConfigKubernetesDashboard from a JSON
// response object.
func flattenClusterAddonsConfigKubernetesDashboard(c *Client, i interface{}) *ClusterAddonsConfigKubernetesDashboard {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAddonsConfigKubernetesDashboard{}
	r.Disabled = dcl.FlattenBool(m["disabled"])

	return r
}

// expandClusterAddonsConfigNetworkPolicyConfigMap expands the contents of ClusterAddonsConfigNetworkPolicyConfig into a JSON
// request object.
func expandClusterAddonsConfigNetworkPolicyConfigMap(c *Client, f map[string]ClusterAddonsConfigNetworkPolicyConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAddonsConfigNetworkPolicyConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAddonsConfigNetworkPolicyConfigSlice expands the contents of ClusterAddonsConfigNetworkPolicyConfig into a JSON
// request object.
func expandClusterAddonsConfigNetworkPolicyConfigSlice(c *Client, f []ClusterAddonsConfigNetworkPolicyConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAddonsConfigNetworkPolicyConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAddonsConfigNetworkPolicyConfigMap flattens the contents of ClusterAddonsConfigNetworkPolicyConfig from a JSON
// response object.
func flattenClusterAddonsConfigNetworkPolicyConfigMap(c *Client, i interface{}) map[string]ClusterAddonsConfigNetworkPolicyConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAddonsConfigNetworkPolicyConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterAddonsConfigNetworkPolicyConfig{}
	}

	items := make(map[string]ClusterAddonsConfigNetworkPolicyConfig)
	for k, item := range a {
		items[k] = *flattenClusterAddonsConfigNetworkPolicyConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAddonsConfigNetworkPolicyConfigSlice flattens the contents of ClusterAddonsConfigNetworkPolicyConfig from a JSON
// response object.
func flattenClusterAddonsConfigNetworkPolicyConfigSlice(c *Client, i interface{}) []ClusterAddonsConfigNetworkPolicyConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfigNetworkPolicyConfig{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfigNetworkPolicyConfig{}
	}

	items := make([]ClusterAddonsConfigNetworkPolicyConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfigNetworkPolicyConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAddonsConfigNetworkPolicyConfig expands an instance of ClusterAddonsConfigNetworkPolicyConfig into a JSON
// request object.
func expandClusterAddonsConfigNetworkPolicyConfig(c *Client, f *ClusterAddonsConfigNetworkPolicyConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	m["disabled"] = f.Disabled

	return m, nil
}

// flattenClusterAddonsConfigNetworkPolicyConfig flattens an instance of ClusterAddonsConfigNetworkPolicyConfig from a JSON
// response object.
func flattenClusterAddonsConfigNetworkPolicyConfig(c *Client, i interface{}) *ClusterAddonsConfigNetworkPolicyConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAddonsConfigNetworkPolicyConfig{}
	r.Disabled = dcl.FlattenBool(m["disabled"])

	return r
}

// expandClusterAddonsConfigCloudRunConfigMap expands the contents of ClusterAddonsConfigCloudRunConfig into a JSON
// request object.
func expandClusterAddonsConfigCloudRunConfigMap(c *Client, f map[string]ClusterAddonsConfigCloudRunConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAddonsConfigCloudRunConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAddonsConfigCloudRunConfigSlice expands the contents of ClusterAddonsConfigCloudRunConfig into a JSON
// request object.
func expandClusterAddonsConfigCloudRunConfigSlice(c *Client, f []ClusterAddonsConfigCloudRunConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAddonsConfigCloudRunConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAddonsConfigCloudRunConfigMap flattens the contents of ClusterAddonsConfigCloudRunConfig from a JSON
// response object.
func flattenClusterAddonsConfigCloudRunConfigMap(c *Client, i interface{}) map[string]ClusterAddonsConfigCloudRunConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAddonsConfigCloudRunConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterAddonsConfigCloudRunConfig{}
	}

	items := make(map[string]ClusterAddonsConfigCloudRunConfig)
	for k, item := range a {
		items[k] = *flattenClusterAddonsConfigCloudRunConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAddonsConfigCloudRunConfigSlice flattens the contents of ClusterAddonsConfigCloudRunConfig from a JSON
// response object.
func flattenClusterAddonsConfigCloudRunConfigSlice(c *Client, i interface{}) []ClusterAddonsConfigCloudRunConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfigCloudRunConfig{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfigCloudRunConfig{}
	}

	items := make([]ClusterAddonsConfigCloudRunConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfigCloudRunConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAddonsConfigCloudRunConfig expands an instance of ClusterAddonsConfigCloudRunConfig into a JSON
// request object.
func expandClusterAddonsConfigCloudRunConfig(c *Client, f *ClusterAddonsConfigCloudRunConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	m["disabled"] = f.Disabled
	if v := f.LoadBalancerType; !dcl.IsEmptyValueIndirect(v) {
		m["loadBalancerType"] = v
	}

	return m, nil
}

// flattenClusterAddonsConfigCloudRunConfig flattens an instance of ClusterAddonsConfigCloudRunConfig from a JSON
// response object.
func flattenClusterAddonsConfigCloudRunConfig(c *Client, i interface{}) *ClusterAddonsConfigCloudRunConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAddonsConfigCloudRunConfig{}
	r.Disabled = dcl.FlattenBool(m["disabled"])
	r.LoadBalancerType = flattenClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(m["loadBalancerType"])

	return r
}

// expandClusterAddonsConfigDnsCacheConfigMap expands the contents of ClusterAddonsConfigDnsCacheConfig into a JSON
// request object.
func expandClusterAddonsConfigDnsCacheConfigMap(c *Client, f map[string]ClusterAddonsConfigDnsCacheConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAddonsConfigDnsCacheConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAddonsConfigDnsCacheConfigSlice expands the contents of ClusterAddonsConfigDnsCacheConfig into a JSON
// request object.
func expandClusterAddonsConfigDnsCacheConfigSlice(c *Client, f []ClusterAddonsConfigDnsCacheConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAddonsConfigDnsCacheConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAddonsConfigDnsCacheConfigMap flattens the contents of ClusterAddonsConfigDnsCacheConfig from a JSON
// response object.
func flattenClusterAddonsConfigDnsCacheConfigMap(c *Client, i interface{}) map[string]ClusterAddonsConfigDnsCacheConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAddonsConfigDnsCacheConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterAddonsConfigDnsCacheConfig{}
	}

	items := make(map[string]ClusterAddonsConfigDnsCacheConfig)
	for k, item := range a {
		items[k] = *flattenClusterAddonsConfigDnsCacheConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAddonsConfigDnsCacheConfigSlice flattens the contents of ClusterAddonsConfigDnsCacheConfig from a JSON
// response object.
func flattenClusterAddonsConfigDnsCacheConfigSlice(c *Client, i interface{}) []ClusterAddonsConfigDnsCacheConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfigDnsCacheConfig{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfigDnsCacheConfig{}
	}

	items := make([]ClusterAddonsConfigDnsCacheConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfigDnsCacheConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAddonsConfigDnsCacheConfig expands an instance of ClusterAddonsConfigDnsCacheConfig into a JSON
// request object.
func expandClusterAddonsConfigDnsCacheConfig(c *Client, f *ClusterAddonsConfigDnsCacheConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterAddonsConfigDnsCacheConfig flattens an instance of ClusterAddonsConfigDnsCacheConfig from a JSON
// response object.
func flattenClusterAddonsConfigDnsCacheConfig(c *Client, i interface{}) *ClusterAddonsConfigDnsCacheConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAddonsConfigDnsCacheConfig{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterAddonsConfigConfigConnectorConfigMap expands the contents of ClusterAddonsConfigConfigConnectorConfig into a JSON
// request object.
func expandClusterAddonsConfigConfigConnectorConfigMap(c *Client, f map[string]ClusterAddonsConfigConfigConnectorConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAddonsConfigConfigConnectorConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAddonsConfigConfigConnectorConfigSlice expands the contents of ClusterAddonsConfigConfigConnectorConfig into a JSON
// request object.
func expandClusterAddonsConfigConfigConnectorConfigSlice(c *Client, f []ClusterAddonsConfigConfigConnectorConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAddonsConfigConfigConnectorConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAddonsConfigConfigConnectorConfigMap flattens the contents of ClusterAddonsConfigConfigConnectorConfig from a JSON
// response object.
func flattenClusterAddonsConfigConfigConnectorConfigMap(c *Client, i interface{}) map[string]ClusterAddonsConfigConfigConnectorConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAddonsConfigConfigConnectorConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterAddonsConfigConfigConnectorConfig{}
	}

	items := make(map[string]ClusterAddonsConfigConfigConnectorConfig)
	for k, item := range a {
		items[k] = *flattenClusterAddonsConfigConfigConnectorConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAddonsConfigConfigConnectorConfigSlice flattens the contents of ClusterAddonsConfigConfigConnectorConfig from a JSON
// response object.
func flattenClusterAddonsConfigConfigConnectorConfigSlice(c *Client, i interface{}) []ClusterAddonsConfigConfigConnectorConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfigConfigConnectorConfig{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfigConfigConnectorConfig{}
	}

	items := make([]ClusterAddonsConfigConfigConnectorConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfigConfigConnectorConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAddonsConfigConfigConnectorConfig expands an instance of ClusterAddonsConfigConfigConnectorConfig into a JSON
// request object.
func expandClusterAddonsConfigConfigConnectorConfig(c *Client, f *ClusterAddonsConfigConfigConnectorConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterAddonsConfigConfigConnectorConfig flattens an instance of ClusterAddonsConfigConfigConnectorConfig from a JSON
// response object.
func flattenClusterAddonsConfigConfigConnectorConfig(c *Client, i interface{}) *ClusterAddonsConfigConfigConnectorConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAddonsConfigConfigConnectorConfig{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterAddonsConfigGcePersistentDiskCsiDriverConfigMap expands the contents of ClusterAddonsConfigGcePersistentDiskCsiDriverConfig into a JSON
// request object.
func expandClusterAddonsConfigGcePersistentDiskCsiDriverConfigMap(c *Client, f map[string]ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAddonsConfigGcePersistentDiskCsiDriverConfigSlice expands the contents of ClusterAddonsConfigGcePersistentDiskCsiDriverConfig into a JSON
// request object.
func expandClusterAddonsConfigGcePersistentDiskCsiDriverConfigSlice(c *Client, f []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAddonsConfigGcePersistentDiskCsiDriverConfigMap flattens the contents of ClusterAddonsConfigGcePersistentDiskCsiDriverConfig from a JSON
// response object.
func flattenClusterAddonsConfigGcePersistentDiskCsiDriverConfigMap(c *Client, i interface{}) map[string]ClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAddonsConfigGcePersistentDiskCsiDriverConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterAddonsConfigGcePersistentDiskCsiDriverConfig{}
	}

	items := make(map[string]ClusterAddonsConfigGcePersistentDiskCsiDriverConfig)
	for k, item := range a {
		items[k] = *flattenClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAddonsConfigGcePersistentDiskCsiDriverConfigSlice flattens the contents of ClusterAddonsConfigGcePersistentDiskCsiDriverConfig from a JSON
// response object.
func flattenClusterAddonsConfigGcePersistentDiskCsiDriverConfigSlice(c *Client, i interface{}) []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfigGcePersistentDiskCsiDriverConfig{}
	}

	items := make([]ClusterAddonsConfigGcePersistentDiskCsiDriverConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAddonsConfigGcePersistentDiskCsiDriverConfig expands an instance of ClusterAddonsConfigGcePersistentDiskCsiDriverConfig into a JSON
// request object.
func expandClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c *Client, f *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterAddonsConfigGcePersistentDiskCsiDriverConfig flattens an instance of ClusterAddonsConfigGcePersistentDiskCsiDriverConfig from a JSON
// response object.
func flattenClusterAddonsConfigGcePersistentDiskCsiDriverConfig(c *Client, i interface{}) *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAddonsConfigGcePersistentDiskCsiDriverConfig{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterNodePoolsMap expands the contents of ClusterNodePools into a JSON
// request object.
func expandClusterNodePoolsMap(c *Client, f map[string]ClusterNodePools) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePools(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsSlice expands the contents of ClusterNodePools into a JSON
// request object.
func expandClusterNodePoolsSlice(c *Client, f []ClusterNodePools) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePools(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsMap flattens the contents of ClusterNodePools from a JSON
// response object.
func flattenClusterNodePoolsMap(c *Client, i interface{}) map[string]ClusterNodePools {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePools{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePools{}
	}

	items := make(map[string]ClusterNodePools)
	for k, item := range a {
		items[k] = *flattenClusterNodePools(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsSlice flattens the contents of ClusterNodePools from a JSON
// response object.
func flattenClusterNodePoolsSlice(c *Client, i interface{}) []ClusterNodePools {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePools{}
	}

	if len(a) == 0 {
		return []ClusterNodePools{}
	}

	items := make([]ClusterNodePools, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePools(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePools expands an instance of ClusterNodePools into a JSON
// request object.
func expandClusterNodePools(c *Client, f *ClusterNodePools) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandClusterNodePoolsConfig(c, f.Config); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["config"] = v
	}
	if v := f.InitialNodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["initialNodeCount"] = v
	}
	if v := f.Locations; !dcl.IsEmptyValueIndirect(v) {
		m["locations"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.InstanceGroupUrls; !dcl.IsEmptyValueIndirect(v) {
		m["instanceGroupUrls"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.StatusMessage; !dcl.IsEmptyValueIndirect(v) {
		m["statusMessage"] = v
	}
	if v, err := expandClusterNodePoolsAutoscaling(c, f.Autoscaling); err != nil {
		return nil, fmt.Errorf("error expanding Autoscaling into autoscaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoscaling"] = v
	}
	if v, err := expandClusterNodePoolsManagement(c, f.Management); err != nil {
		return nil, fmt.Errorf("error expanding Management into management: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["management"] = v
	}
	if v, err := expandClusterNodePoolsMaxPodsConstraint(c, f.MaxPodsConstraint); err != nil {
		return nil, fmt.Errorf("error expanding MaxPodsConstraint into maxPodsConstraint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maxPodsConstraint"] = v
	}
	if v, err := expandClusterNodePoolsConditionsSlice(c, f.Conditions); err != nil {
		return nil, fmt.Errorf("error expanding Conditions into conditions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditions"] = v
	}
	if v := f.PodIPv4CidrSize; !dcl.IsEmptyValueIndirect(v) {
		m["podIpv4CidrSize"] = v
	}
	if v, err := expandClusterNodePoolsUpgradeSettings(c, f.UpgradeSettings); err != nil {
		return nil, fmt.Errorf("error expanding UpgradeSettings into upgradeSettings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["upgradeSettings"] = v
	}

	return m, nil
}

// flattenClusterNodePools flattens an instance of ClusterNodePools from a JSON
// response object.
func flattenClusterNodePools(c *Client, i interface{}) *ClusterNodePools {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePools{}
	r.Name = dcl.FlattenString(m["name"])
	r.Config = flattenClusterNodePoolsConfig(c, m["config"])
	r.InitialNodeCount = dcl.FlattenInteger(m["initialNodeCount"])
	r.Locations = dcl.FlattenStringSlice(m["locations"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Version = dcl.FlattenString(m["version"])
	r.InstanceGroupUrls = dcl.FlattenStringSlice(m["instanceGroupUrls"])
	r.Status = flattenClusterNodePoolsStatusEnum(m["status"])
	r.StatusMessage = dcl.FlattenString(m["statusMessage"])
	r.Autoscaling = flattenClusterNodePoolsAutoscaling(c, m["autoscaling"])
	r.Management = flattenClusterNodePoolsManagement(c, m["management"])
	r.MaxPodsConstraint = flattenClusterNodePoolsMaxPodsConstraint(c, m["maxPodsConstraint"])
	r.Conditions = flattenClusterNodePoolsConditionsSlice(c, m["conditions"])
	r.PodIPv4CidrSize = dcl.FlattenInteger(m["podIpv4CidrSize"])
	r.UpgradeSettings = flattenClusterNodePoolsUpgradeSettings(c, m["upgradeSettings"])

	return r
}

// expandClusterNodePoolsConfigMap expands the contents of ClusterNodePoolsConfig into a JSON
// request object.
func expandClusterNodePoolsConfigMap(c *Client, f map[string]ClusterNodePoolsConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConfigSlice expands the contents of ClusterNodePoolsConfig into a JSON
// request object.
func expandClusterNodePoolsConfigSlice(c *Client, f []ClusterNodePoolsConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConfigMap flattens the contents of ClusterNodePoolsConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigMap(c *Client, i interface{}) map[string]ClusterNodePoolsConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConfig{}
	}

	items := make(map[string]ClusterNodePoolsConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConfigSlice flattens the contents of ClusterNodePoolsConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigSlice(c *Client, i interface{}) []ClusterNodePoolsConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfig{}
	}

	items := make([]ClusterNodePoolsConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConfig expands an instance of ClusterNodePoolsConfig into a JSON
// request object.
func expandClusterNodePoolsConfig(c *Client, f *ClusterNodePoolsConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}
	if v := f.DiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskSizeGb"] = v
	}
	if v := f.OAuthScopes; !dcl.IsEmptyValueIndirect(v) {
		m["oauthScopes"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v := f.Metadata; !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v := f.ImageType; !dcl.IsEmptyValueIndirect(v) {
		m["imageType"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.LocalSsdCount; !dcl.IsEmptyValueIndirect(v) {
		m["localSsdCount"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v := f.Preemptible; !dcl.IsEmptyValueIndirect(v) {
		m["preemptible"] = v
	}
	if v, err := expandClusterNodePoolsConfigAcceleratorsSlice(c, f.Accelerators); err != nil {
		return nil, fmt.Errorf("error expanding Accelerators into accelerators: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["accelerators"] = v
	}
	if v := f.DiskType; !dcl.IsEmptyValueIndirect(v) {
		m["diskType"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}
	if v, err := expandClusterNodePoolsConfigWorkloadMetadataConfig(c, f.WorkloadMetadataConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkloadMetadataConfig into workloadMetadataConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["workloadMetadataConfig"] = v
	}
	if v, err := expandClusterNodePoolsConfigTaintsSlice(c, f.Taints); err != nil {
		return nil, fmt.Errorf("error expanding Taints into taints: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["taints"] = v
	}
	if v, err := expandClusterNodePoolsConfigSandboxConfig(c, f.SandboxConfig); err != nil {
		return nil, fmt.Errorf("error expanding SandboxConfig into sandboxConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sandboxConfig"] = v
	}
	if v := f.NodeGroup; !dcl.IsEmptyValueIndirect(v) {
		m["nodeGroup"] = v
	}
	if v, err := expandClusterNodePoolsConfigReservationAffinity(c, f.ReservationAffinity); err != nil {
		return nil, fmt.Errorf("error expanding ReservationAffinity into reservationAffinity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reservationAffinity"] = v
	}
	if v, err := expandClusterNodePoolsConfigShieldedInstanceConfig(c, f.ShieldedInstanceConfig); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedInstanceConfig into shieldedInstanceConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["shieldedInstanceConfig"] = v
	}
	if v, err := expandClusterNodePoolsConfigLinuxNodeConfig(c, f.LinuxNodeConfig); err != nil {
		return nil, fmt.Errorf("error expanding LinuxNodeConfig into linuxNodeConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["linuxNodeConfig"] = v
	}
	if v, err := expandClusterNodePoolsConfigKubeletConfig(c, f.KubeletConfig); err != nil {
		return nil, fmt.Errorf("error expanding KubeletConfig into kubeletConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kubeletConfig"] = v
	}
	if v := f.BootDiskKmsKey; !dcl.IsEmptyValueIndirect(v) {
		m["bootDiskKmsKey"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConfig flattens an instance of ClusterNodePoolsConfig from a JSON
// response object.
func flattenClusterNodePoolsConfig(c *Client, i interface{}) *ClusterNodePoolsConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConfig{}
	r.MachineType = dcl.FlattenString(m["machineType"])
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.OAuthScopes = dcl.FlattenStringSlice(m["oauthScopes"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.Metadata = dcl.FlattenKeyValuePairs(m["metadata"])
	r.ImageType = dcl.FlattenString(m["imageType"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.LocalSsdCount = dcl.FlattenInteger(m["localSsdCount"])
	r.Tags = dcl.FlattenStringSlice(m["tags"])
	r.Preemptible = dcl.FlattenBool(m["preemptible"])
	r.Accelerators = flattenClusterNodePoolsConfigAcceleratorsSlice(c, m["accelerators"])
	r.DiskType = dcl.FlattenString(m["diskType"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])
	r.WorkloadMetadataConfig = flattenClusterNodePoolsConfigWorkloadMetadataConfig(c, m["workloadMetadataConfig"])
	r.Taints = flattenClusterNodePoolsConfigTaintsSlice(c, m["taints"])
	r.SandboxConfig = flattenClusterNodePoolsConfigSandboxConfig(c, m["sandboxConfig"])
	r.NodeGroup = dcl.FlattenString(m["nodeGroup"])
	r.ReservationAffinity = flattenClusterNodePoolsConfigReservationAffinity(c, m["reservationAffinity"])
	r.ShieldedInstanceConfig = flattenClusterNodePoolsConfigShieldedInstanceConfig(c, m["shieldedInstanceConfig"])
	r.LinuxNodeConfig = flattenClusterNodePoolsConfigLinuxNodeConfig(c, m["linuxNodeConfig"])
	r.KubeletConfig = flattenClusterNodePoolsConfigKubeletConfig(c, m["kubeletConfig"])
	r.BootDiskKmsKey = dcl.FlattenString(m["bootDiskKmsKey"])

	return r
}

// expandClusterNodePoolsConfigAcceleratorsMap expands the contents of ClusterNodePoolsConfigAccelerators into a JSON
// request object.
func expandClusterNodePoolsConfigAcceleratorsMap(c *Client, f map[string]ClusterNodePoolsConfigAccelerators) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConfigAccelerators(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConfigAcceleratorsSlice expands the contents of ClusterNodePoolsConfigAccelerators into a JSON
// request object.
func expandClusterNodePoolsConfigAcceleratorsSlice(c *Client, f []ClusterNodePoolsConfigAccelerators) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConfigAccelerators(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConfigAcceleratorsMap flattens the contents of ClusterNodePoolsConfigAccelerators from a JSON
// response object.
func flattenClusterNodePoolsConfigAcceleratorsMap(c *Client, i interface{}) map[string]ClusterNodePoolsConfigAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConfigAccelerators{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConfigAccelerators{}
	}

	items := make(map[string]ClusterNodePoolsConfigAccelerators)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConfigAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConfigAcceleratorsSlice flattens the contents of ClusterNodePoolsConfigAccelerators from a JSON
// response object.
func flattenClusterNodePoolsConfigAcceleratorsSlice(c *Client, i interface{}) []ClusterNodePoolsConfigAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigAccelerators{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigAccelerators{}
	}

	items := make([]ClusterNodePoolsConfigAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConfigAccelerators expands an instance of ClusterNodePoolsConfigAccelerators into a JSON
// request object.
func expandClusterNodePoolsConfigAccelerators(c *Client, f *ClusterNodePoolsConfigAccelerators) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.AcceleratorCount; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorCount"] = v
	}
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorType"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConfigAccelerators flattens an instance of ClusterNodePoolsConfigAccelerators from a JSON
// response object.
func flattenClusterNodePoolsConfigAccelerators(c *Client, i interface{}) *ClusterNodePoolsConfigAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConfigAccelerators{}
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])
	r.AcceleratorType = dcl.FlattenString(m["acceleratorType"])

	return r
}

// expandClusterNodePoolsConfigWorkloadMetadataConfigMap expands the contents of ClusterNodePoolsConfigWorkloadMetadataConfig into a JSON
// request object.
func expandClusterNodePoolsConfigWorkloadMetadataConfigMap(c *Client, f map[string]ClusterNodePoolsConfigWorkloadMetadataConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConfigWorkloadMetadataConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConfigWorkloadMetadataConfigSlice expands the contents of ClusterNodePoolsConfigWorkloadMetadataConfig into a JSON
// request object.
func expandClusterNodePoolsConfigWorkloadMetadataConfigSlice(c *Client, f []ClusterNodePoolsConfigWorkloadMetadataConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConfigWorkloadMetadataConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConfigWorkloadMetadataConfigMap flattens the contents of ClusterNodePoolsConfigWorkloadMetadataConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigWorkloadMetadataConfigMap(c *Client, i interface{}) map[string]ClusterNodePoolsConfigWorkloadMetadataConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConfigWorkloadMetadataConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConfigWorkloadMetadataConfig{}
	}

	items := make(map[string]ClusterNodePoolsConfigWorkloadMetadataConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConfigWorkloadMetadataConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConfigWorkloadMetadataConfigSlice flattens the contents of ClusterNodePoolsConfigWorkloadMetadataConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigWorkloadMetadataConfigSlice(c *Client, i interface{}) []ClusterNodePoolsConfigWorkloadMetadataConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigWorkloadMetadataConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigWorkloadMetadataConfig{}
	}

	items := make([]ClusterNodePoolsConfigWorkloadMetadataConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigWorkloadMetadataConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConfigWorkloadMetadataConfig expands an instance of ClusterNodePoolsConfigWorkloadMetadataConfig into a JSON
// request object.
func expandClusterNodePoolsConfigWorkloadMetadataConfig(c *Client, f *ClusterNodePoolsConfigWorkloadMetadataConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConfigWorkloadMetadataConfig flattens an instance of ClusterNodePoolsConfigWorkloadMetadataConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigWorkloadMetadataConfig(c *Client, i interface{}) *ClusterNodePoolsConfigWorkloadMetadataConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConfigWorkloadMetadataConfig{}
	r.Mode = flattenClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(m["mode"])

	return r
}

// expandClusterNodePoolsConfigTaintsMap expands the contents of ClusterNodePoolsConfigTaints into a JSON
// request object.
func expandClusterNodePoolsConfigTaintsMap(c *Client, f map[string]ClusterNodePoolsConfigTaints) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConfigTaints(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConfigTaintsSlice expands the contents of ClusterNodePoolsConfigTaints into a JSON
// request object.
func expandClusterNodePoolsConfigTaintsSlice(c *Client, f []ClusterNodePoolsConfigTaints) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConfigTaints(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConfigTaintsMap flattens the contents of ClusterNodePoolsConfigTaints from a JSON
// response object.
func flattenClusterNodePoolsConfigTaintsMap(c *Client, i interface{}) map[string]ClusterNodePoolsConfigTaints {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConfigTaints{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConfigTaints{}
	}

	items := make(map[string]ClusterNodePoolsConfigTaints)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConfigTaints(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConfigTaintsSlice flattens the contents of ClusterNodePoolsConfigTaints from a JSON
// response object.
func flattenClusterNodePoolsConfigTaintsSlice(c *Client, i interface{}) []ClusterNodePoolsConfigTaints {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigTaints{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigTaints{}
	}

	items := make([]ClusterNodePoolsConfigTaints, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigTaints(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConfigTaints expands an instance of ClusterNodePoolsConfigTaints into a JSON
// request object.
func expandClusterNodePoolsConfigTaints(c *Client, f *ClusterNodePoolsConfigTaints) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.Effect; !dcl.IsEmptyValueIndirect(v) {
		m["effect"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConfigTaints flattens an instance of ClusterNodePoolsConfigTaints from a JSON
// response object.
func flattenClusterNodePoolsConfigTaints(c *Client, i interface{}) *ClusterNodePoolsConfigTaints {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConfigTaints{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])
	r.Effect = flattenClusterNodePoolsConfigTaintsEffectEnum(m["effect"])

	return r
}

// expandClusterNodePoolsConfigSandboxConfigMap expands the contents of ClusterNodePoolsConfigSandboxConfig into a JSON
// request object.
func expandClusterNodePoolsConfigSandboxConfigMap(c *Client, f map[string]ClusterNodePoolsConfigSandboxConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConfigSandboxConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConfigSandboxConfigSlice expands the contents of ClusterNodePoolsConfigSandboxConfig into a JSON
// request object.
func expandClusterNodePoolsConfigSandboxConfigSlice(c *Client, f []ClusterNodePoolsConfigSandboxConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConfigSandboxConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConfigSandboxConfigMap flattens the contents of ClusterNodePoolsConfigSandboxConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigSandboxConfigMap(c *Client, i interface{}) map[string]ClusterNodePoolsConfigSandboxConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConfigSandboxConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConfigSandboxConfig{}
	}

	items := make(map[string]ClusterNodePoolsConfigSandboxConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConfigSandboxConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConfigSandboxConfigSlice flattens the contents of ClusterNodePoolsConfigSandboxConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigSandboxConfigSlice(c *Client, i interface{}) []ClusterNodePoolsConfigSandboxConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigSandboxConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigSandboxConfig{}
	}

	items := make([]ClusterNodePoolsConfigSandboxConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigSandboxConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConfigSandboxConfig expands an instance of ClusterNodePoolsConfigSandboxConfig into a JSON
// request object.
func expandClusterNodePoolsConfigSandboxConfig(c *Client, f *ClusterNodePoolsConfigSandboxConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConfigSandboxConfig flattens an instance of ClusterNodePoolsConfigSandboxConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigSandboxConfig(c *Client, i interface{}) *ClusterNodePoolsConfigSandboxConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConfigSandboxConfig{}
	r.Type = flattenClusterNodePoolsConfigSandboxConfigTypeEnum(m["type"])

	return r
}

// expandClusterNodePoolsConfigReservationAffinityMap expands the contents of ClusterNodePoolsConfigReservationAffinity into a JSON
// request object.
func expandClusterNodePoolsConfigReservationAffinityMap(c *Client, f map[string]ClusterNodePoolsConfigReservationAffinity) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConfigReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConfigReservationAffinitySlice expands the contents of ClusterNodePoolsConfigReservationAffinity into a JSON
// request object.
func expandClusterNodePoolsConfigReservationAffinitySlice(c *Client, f []ClusterNodePoolsConfigReservationAffinity) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConfigReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConfigReservationAffinityMap flattens the contents of ClusterNodePoolsConfigReservationAffinity from a JSON
// response object.
func flattenClusterNodePoolsConfigReservationAffinityMap(c *Client, i interface{}) map[string]ClusterNodePoolsConfigReservationAffinity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConfigReservationAffinity{}
	}

	items := make(map[string]ClusterNodePoolsConfigReservationAffinity)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConfigReservationAffinity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConfigReservationAffinitySlice flattens the contents of ClusterNodePoolsConfigReservationAffinity from a JSON
// response object.
func flattenClusterNodePoolsConfigReservationAffinitySlice(c *Client, i interface{}) []ClusterNodePoolsConfigReservationAffinity {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigReservationAffinity{}
	}

	items := make([]ClusterNodePoolsConfigReservationAffinity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigReservationAffinity(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConfigReservationAffinity expands an instance of ClusterNodePoolsConfigReservationAffinity into a JSON
// request object.
func expandClusterNodePoolsConfigReservationAffinity(c *Client, f *ClusterNodePoolsConfigReservationAffinity) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ConsumeReservationType; !dcl.IsEmptyValueIndirect(v) {
		m["consumeReservationType"] = v
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Values; !dcl.IsEmptyValueIndirect(v) {
		m["values"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConfigReservationAffinity flattens an instance of ClusterNodePoolsConfigReservationAffinity from a JSON
// response object.
func flattenClusterNodePoolsConfigReservationAffinity(c *Client, i interface{}) *ClusterNodePoolsConfigReservationAffinity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConfigReservationAffinity{}
	r.ConsumeReservationType = flattenClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(m["consumeReservationType"])
	r.Key = dcl.FlattenString(m["key"])
	r.Values = dcl.FlattenStringSlice(m["values"])

	return r
}

// expandClusterNodePoolsConfigShieldedInstanceConfigMap expands the contents of ClusterNodePoolsConfigShieldedInstanceConfig into a JSON
// request object.
func expandClusterNodePoolsConfigShieldedInstanceConfigMap(c *Client, f map[string]ClusterNodePoolsConfigShieldedInstanceConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConfigShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConfigShieldedInstanceConfigSlice expands the contents of ClusterNodePoolsConfigShieldedInstanceConfig into a JSON
// request object.
func expandClusterNodePoolsConfigShieldedInstanceConfigSlice(c *Client, f []ClusterNodePoolsConfigShieldedInstanceConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConfigShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConfigShieldedInstanceConfigMap flattens the contents of ClusterNodePoolsConfigShieldedInstanceConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigShieldedInstanceConfigMap(c *Client, i interface{}) map[string]ClusterNodePoolsConfigShieldedInstanceConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConfigShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConfigShieldedInstanceConfig{}
	}

	items := make(map[string]ClusterNodePoolsConfigShieldedInstanceConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConfigShieldedInstanceConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConfigShieldedInstanceConfigSlice flattens the contents of ClusterNodePoolsConfigShieldedInstanceConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigShieldedInstanceConfigSlice(c *Client, i interface{}) []ClusterNodePoolsConfigShieldedInstanceConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigShieldedInstanceConfig{}
	}

	items := make([]ClusterNodePoolsConfigShieldedInstanceConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigShieldedInstanceConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConfigShieldedInstanceConfig expands an instance of ClusterNodePoolsConfigShieldedInstanceConfig into a JSON
// request object.
func expandClusterNodePoolsConfigShieldedInstanceConfig(c *Client, f *ClusterNodePoolsConfigShieldedInstanceConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.EnableSecureBoot; !dcl.IsEmptyValueIndirect(v) {
		m["enableSecureBoot"] = v
	}
	if v := f.EnableIntegrityMonitoring; !dcl.IsEmptyValueIndirect(v) {
		m["enableIntegrityMonitoring"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConfigShieldedInstanceConfig flattens an instance of ClusterNodePoolsConfigShieldedInstanceConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigShieldedInstanceConfig(c *Client, i interface{}) *ClusterNodePoolsConfigShieldedInstanceConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConfigShieldedInstanceConfig{}
	r.EnableSecureBoot = dcl.FlattenBool(m["enableSecureBoot"])
	r.EnableIntegrityMonitoring = dcl.FlattenBool(m["enableIntegrityMonitoring"])

	return r
}

// expandClusterNodePoolsConfigLinuxNodeConfigMap expands the contents of ClusterNodePoolsConfigLinuxNodeConfig into a JSON
// request object.
func expandClusterNodePoolsConfigLinuxNodeConfigMap(c *Client, f map[string]ClusterNodePoolsConfigLinuxNodeConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConfigLinuxNodeConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConfigLinuxNodeConfigSlice expands the contents of ClusterNodePoolsConfigLinuxNodeConfig into a JSON
// request object.
func expandClusterNodePoolsConfigLinuxNodeConfigSlice(c *Client, f []ClusterNodePoolsConfigLinuxNodeConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConfigLinuxNodeConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConfigLinuxNodeConfigMap flattens the contents of ClusterNodePoolsConfigLinuxNodeConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigLinuxNodeConfigMap(c *Client, i interface{}) map[string]ClusterNodePoolsConfigLinuxNodeConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConfigLinuxNodeConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConfigLinuxNodeConfig{}
	}

	items := make(map[string]ClusterNodePoolsConfigLinuxNodeConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConfigLinuxNodeConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConfigLinuxNodeConfigSlice flattens the contents of ClusterNodePoolsConfigLinuxNodeConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigLinuxNodeConfigSlice(c *Client, i interface{}) []ClusterNodePoolsConfigLinuxNodeConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigLinuxNodeConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigLinuxNodeConfig{}
	}

	items := make([]ClusterNodePoolsConfigLinuxNodeConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigLinuxNodeConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConfigLinuxNodeConfig expands an instance of ClusterNodePoolsConfigLinuxNodeConfig into a JSON
// request object.
func expandClusterNodePoolsConfigLinuxNodeConfig(c *Client, f *ClusterNodePoolsConfigLinuxNodeConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Sysctls; !dcl.IsEmptyValueIndirect(v) {
		m["sysctls"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConfigLinuxNodeConfig flattens an instance of ClusterNodePoolsConfigLinuxNodeConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigLinuxNodeConfig(c *Client, i interface{}) *ClusterNodePoolsConfigLinuxNodeConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConfigLinuxNodeConfig{}
	r.Sysctls = dcl.FlattenKeyValuePairs(m["sysctls"])

	return r
}

// expandClusterNodePoolsConfigKubeletConfigMap expands the contents of ClusterNodePoolsConfigKubeletConfig into a JSON
// request object.
func expandClusterNodePoolsConfigKubeletConfigMap(c *Client, f map[string]ClusterNodePoolsConfigKubeletConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConfigKubeletConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConfigKubeletConfigSlice expands the contents of ClusterNodePoolsConfigKubeletConfig into a JSON
// request object.
func expandClusterNodePoolsConfigKubeletConfigSlice(c *Client, f []ClusterNodePoolsConfigKubeletConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConfigKubeletConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConfigKubeletConfigMap flattens the contents of ClusterNodePoolsConfigKubeletConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigKubeletConfigMap(c *Client, i interface{}) map[string]ClusterNodePoolsConfigKubeletConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConfigKubeletConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConfigKubeletConfig{}
	}

	items := make(map[string]ClusterNodePoolsConfigKubeletConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConfigKubeletConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConfigKubeletConfigSlice flattens the contents of ClusterNodePoolsConfigKubeletConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigKubeletConfigSlice(c *Client, i interface{}) []ClusterNodePoolsConfigKubeletConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigKubeletConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigKubeletConfig{}
	}

	items := make([]ClusterNodePoolsConfigKubeletConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigKubeletConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConfigKubeletConfig expands an instance of ClusterNodePoolsConfigKubeletConfig into a JSON
// request object.
func expandClusterNodePoolsConfigKubeletConfig(c *Client, f *ClusterNodePoolsConfigKubeletConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.CpuManagerPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["cpuManagerPolicy"] = v
	}
	if v := f.CpuCfsQuota; !dcl.IsEmptyValueIndirect(v) {
		m["cpuCfsQuota"] = v
	}
	if v := f.CpuCfsQuotaPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["cpuCfsQuotaPeriod"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConfigKubeletConfig flattens an instance of ClusterNodePoolsConfigKubeletConfig from a JSON
// response object.
func flattenClusterNodePoolsConfigKubeletConfig(c *Client, i interface{}) *ClusterNodePoolsConfigKubeletConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConfigKubeletConfig{}
	r.CpuManagerPolicy = dcl.FlattenString(m["cpuManagerPolicy"])
	r.CpuCfsQuota = dcl.FlattenBool(m["cpuCfsQuota"])
	r.CpuCfsQuotaPeriod = dcl.FlattenString(m["cpuCfsQuotaPeriod"])

	return r
}

// expandClusterNodePoolsAutoscalingMap expands the contents of ClusterNodePoolsAutoscaling into a JSON
// request object.
func expandClusterNodePoolsAutoscalingMap(c *Client, f map[string]ClusterNodePoolsAutoscaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsAutoscalingSlice expands the contents of ClusterNodePoolsAutoscaling into a JSON
// request object.
func expandClusterNodePoolsAutoscalingSlice(c *Client, f []ClusterNodePoolsAutoscaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsAutoscalingMap flattens the contents of ClusterNodePoolsAutoscaling from a JSON
// response object.
func flattenClusterNodePoolsAutoscalingMap(c *Client, i interface{}) map[string]ClusterNodePoolsAutoscaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsAutoscaling{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsAutoscaling{}
	}

	items := make(map[string]ClusterNodePoolsAutoscaling)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsAutoscaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsAutoscalingSlice flattens the contents of ClusterNodePoolsAutoscaling from a JSON
// response object.
func flattenClusterNodePoolsAutoscalingSlice(c *Client, i interface{}) []ClusterNodePoolsAutoscaling {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsAutoscaling{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsAutoscaling{}
	}

	items := make([]ClusterNodePoolsAutoscaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsAutoscaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsAutoscaling expands an instance of ClusterNodePoolsAutoscaling into a JSON
// request object.
func expandClusterNodePoolsAutoscaling(c *Client, f *ClusterNodePoolsAutoscaling) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.MinNodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["minNodeCount"] = v
	}
	if v := f.MaxNodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxNodeCount"] = v
	}
	if v := f.Autoprovisioned; !dcl.IsEmptyValueIndirect(v) {
		m["autoprovisioned"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsAutoscaling flattens an instance of ClusterNodePoolsAutoscaling from a JSON
// response object.
func flattenClusterNodePoolsAutoscaling(c *Client, i interface{}) *ClusterNodePoolsAutoscaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsAutoscaling{}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.MinNodeCount = dcl.FlattenInteger(m["minNodeCount"])
	r.MaxNodeCount = dcl.FlattenInteger(m["maxNodeCount"])
	r.Autoprovisioned = dcl.FlattenBool(m["autoprovisioned"])

	return r
}

// expandClusterNodePoolsManagementMap expands the contents of ClusterNodePoolsManagement into a JSON
// request object.
func expandClusterNodePoolsManagementMap(c *Client, f map[string]ClusterNodePoolsManagement) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsManagement(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsManagementSlice expands the contents of ClusterNodePoolsManagement into a JSON
// request object.
func expandClusterNodePoolsManagementSlice(c *Client, f []ClusterNodePoolsManagement) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsManagement(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsManagementMap flattens the contents of ClusterNodePoolsManagement from a JSON
// response object.
func flattenClusterNodePoolsManagementMap(c *Client, i interface{}) map[string]ClusterNodePoolsManagement {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsManagement{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsManagement{}
	}

	items := make(map[string]ClusterNodePoolsManagement)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsManagement(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsManagementSlice flattens the contents of ClusterNodePoolsManagement from a JSON
// response object.
func flattenClusterNodePoolsManagementSlice(c *Client, i interface{}) []ClusterNodePoolsManagement {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsManagement{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsManagement{}
	}

	items := make([]ClusterNodePoolsManagement, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsManagement(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsManagement expands an instance of ClusterNodePoolsManagement into a JSON
// request object.
func expandClusterNodePoolsManagement(c *Client, f *ClusterNodePoolsManagement) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.AutoUpgrade; !dcl.IsEmptyValueIndirect(v) {
		m["autoUpgrade"] = v
	}
	if v := f.AutoRepair; !dcl.IsEmptyValueIndirect(v) {
		m["autoRepair"] = v
	}
	if v, err := expandClusterNodePoolsManagementUpgradeOptions(c, f.UpgradeOptions); err != nil {
		return nil, fmt.Errorf("error expanding UpgradeOptions into upgradeOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["upgradeOptions"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsManagement flattens an instance of ClusterNodePoolsManagement from a JSON
// response object.
func flattenClusterNodePoolsManagement(c *Client, i interface{}) *ClusterNodePoolsManagement {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsManagement{}
	r.AutoUpgrade = dcl.FlattenBool(m["autoUpgrade"])
	r.AutoRepair = dcl.FlattenBool(m["autoRepair"])
	r.UpgradeOptions = flattenClusterNodePoolsManagementUpgradeOptions(c, m["upgradeOptions"])

	return r
}

// expandClusterNodePoolsManagementUpgradeOptionsMap expands the contents of ClusterNodePoolsManagementUpgradeOptions into a JSON
// request object.
func expandClusterNodePoolsManagementUpgradeOptionsMap(c *Client, f map[string]ClusterNodePoolsManagementUpgradeOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsManagementUpgradeOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsManagementUpgradeOptionsSlice expands the contents of ClusterNodePoolsManagementUpgradeOptions into a JSON
// request object.
func expandClusterNodePoolsManagementUpgradeOptionsSlice(c *Client, f []ClusterNodePoolsManagementUpgradeOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsManagementUpgradeOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsManagementUpgradeOptionsMap flattens the contents of ClusterNodePoolsManagementUpgradeOptions from a JSON
// response object.
func flattenClusterNodePoolsManagementUpgradeOptionsMap(c *Client, i interface{}) map[string]ClusterNodePoolsManagementUpgradeOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsManagementUpgradeOptions{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsManagementUpgradeOptions{}
	}

	items := make(map[string]ClusterNodePoolsManagementUpgradeOptions)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsManagementUpgradeOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsManagementUpgradeOptionsSlice flattens the contents of ClusterNodePoolsManagementUpgradeOptions from a JSON
// response object.
func flattenClusterNodePoolsManagementUpgradeOptionsSlice(c *Client, i interface{}) []ClusterNodePoolsManagementUpgradeOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsManagementUpgradeOptions{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsManagementUpgradeOptions{}
	}

	items := make([]ClusterNodePoolsManagementUpgradeOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsManagementUpgradeOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsManagementUpgradeOptions expands an instance of ClusterNodePoolsManagementUpgradeOptions into a JSON
// request object.
func expandClusterNodePoolsManagementUpgradeOptions(c *Client, f *ClusterNodePoolsManagementUpgradeOptions) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.AutoUpgradeStartTime; !dcl.IsEmptyValueIndirect(v) {
		m["autoUpgradeStartTime"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsManagementUpgradeOptions flattens an instance of ClusterNodePoolsManagementUpgradeOptions from a JSON
// response object.
func flattenClusterNodePoolsManagementUpgradeOptions(c *Client, i interface{}) *ClusterNodePoolsManagementUpgradeOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsManagementUpgradeOptions{}
	r.AutoUpgradeStartTime = dcl.FlattenString(m["autoUpgradeStartTime"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandClusterNodePoolsMaxPodsConstraintMap expands the contents of ClusterNodePoolsMaxPodsConstraint into a JSON
// request object.
func expandClusterNodePoolsMaxPodsConstraintMap(c *Client, f map[string]ClusterNodePoolsMaxPodsConstraint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsMaxPodsConstraintSlice expands the contents of ClusterNodePoolsMaxPodsConstraint into a JSON
// request object.
func expandClusterNodePoolsMaxPodsConstraintSlice(c *Client, f []ClusterNodePoolsMaxPodsConstraint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsMaxPodsConstraintMap flattens the contents of ClusterNodePoolsMaxPodsConstraint from a JSON
// response object.
func flattenClusterNodePoolsMaxPodsConstraintMap(c *Client, i interface{}) map[string]ClusterNodePoolsMaxPodsConstraint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsMaxPodsConstraint{}
	}

	items := make(map[string]ClusterNodePoolsMaxPodsConstraint)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsMaxPodsConstraint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsMaxPodsConstraintSlice flattens the contents of ClusterNodePoolsMaxPodsConstraint from a JSON
// response object.
func flattenClusterNodePoolsMaxPodsConstraintSlice(c *Client, i interface{}) []ClusterNodePoolsMaxPodsConstraint {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsMaxPodsConstraint{}
	}

	items := make([]ClusterNodePoolsMaxPodsConstraint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsMaxPodsConstraint(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsMaxPodsConstraint expands an instance of ClusterNodePoolsMaxPodsConstraint into a JSON
// request object.
func expandClusterNodePoolsMaxPodsConstraint(c *Client, f *ClusterNodePoolsMaxPodsConstraint) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.MaxPodsPerNode; !dcl.IsEmptyValueIndirect(v) {
		m["maxPodsPerNode"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsMaxPodsConstraint flattens an instance of ClusterNodePoolsMaxPodsConstraint from a JSON
// response object.
func flattenClusterNodePoolsMaxPodsConstraint(c *Client, i interface{}) *ClusterNodePoolsMaxPodsConstraint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsMaxPodsConstraint{}
	r.MaxPodsPerNode = dcl.FlattenInteger(m["maxPodsPerNode"])

	return r
}

// expandClusterNodePoolsConditionsMap expands the contents of ClusterNodePoolsConditions into a JSON
// request object.
func expandClusterNodePoolsConditionsMap(c *Client, f map[string]ClusterNodePoolsConditions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsConditions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsConditionsSlice expands the contents of ClusterNodePoolsConditions into a JSON
// request object.
func expandClusterNodePoolsConditionsSlice(c *Client, f []ClusterNodePoolsConditions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsConditions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsConditionsMap flattens the contents of ClusterNodePoolsConditions from a JSON
// response object.
func flattenClusterNodePoolsConditionsMap(c *Client, i interface{}) map[string]ClusterNodePoolsConditions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsConditions{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsConditions{}
	}

	items := make(map[string]ClusterNodePoolsConditions)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsConditions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsConditionsSlice flattens the contents of ClusterNodePoolsConditions from a JSON
// response object.
func flattenClusterNodePoolsConditionsSlice(c *Client, i interface{}) []ClusterNodePoolsConditions {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConditions{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConditions{}
	}

	items := make([]ClusterNodePoolsConditions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConditions(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsConditions expands an instance of ClusterNodePoolsConditions into a JSON
// request object.
func expandClusterNodePoolsConditions(c *Client, f *ClusterNodePoolsConditions) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v := f.CanonicalCode; !dcl.IsEmptyValueIndirect(v) {
		m["canonicalCode"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsConditions flattens an instance of ClusterNodePoolsConditions from a JSON
// response object.
func flattenClusterNodePoolsConditions(c *Client, i interface{}) *ClusterNodePoolsConditions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsConditions{}
	r.Code = flattenClusterNodePoolsConditionsCodeEnum(m["code"])
	r.Message = dcl.FlattenString(m["message"])
	r.CanonicalCode = flattenClusterNodePoolsConditionsCanonicalCodeEnum(m["canonicalCode"])

	return r
}

// expandClusterNodePoolsUpgradeSettingsMap expands the contents of ClusterNodePoolsUpgradeSettings into a JSON
// request object.
func expandClusterNodePoolsUpgradeSettingsMap(c *Client, f map[string]ClusterNodePoolsUpgradeSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodePoolsUpgradeSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodePoolsUpgradeSettingsSlice expands the contents of ClusterNodePoolsUpgradeSettings into a JSON
// request object.
func expandClusterNodePoolsUpgradeSettingsSlice(c *Client, f []ClusterNodePoolsUpgradeSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodePoolsUpgradeSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodePoolsUpgradeSettingsMap flattens the contents of ClusterNodePoolsUpgradeSettings from a JSON
// response object.
func flattenClusterNodePoolsUpgradeSettingsMap(c *Client, i interface{}) map[string]ClusterNodePoolsUpgradeSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodePoolsUpgradeSettings{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodePoolsUpgradeSettings{}
	}

	items := make(map[string]ClusterNodePoolsUpgradeSettings)
	for k, item := range a {
		items[k] = *flattenClusterNodePoolsUpgradeSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodePoolsUpgradeSettingsSlice flattens the contents of ClusterNodePoolsUpgradeSettings from a JSON
// response object.
func flattenClusterNodePoolsUpgradeSettingsSlice(c *Client, i interface{}) []ClusterNodePoolsUpgradeSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsUpgradeSettings{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsUpgradeSettings{}
	}

	items := make([]ClusterNodePoolsUpgradeSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsUpgradeSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodePoolsUpgradeSettings expands an instance of ClusterNodePoolsUpgradeSettings into a JSON
// request object.
func expandClusterNodePoolsUpgradeSettings(c *Client, f *ClusterNodePoolsUpgradeSettings) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.MaxSurge; !dcl.IsEmptyValueIndirect(v) {
		m["maxSurge"] = v
	}
	if v := f.MaxUnavailable; !dcl.IsEmptyValueIndirect(v) {
		m["maxUnavailable"] = v
	}

	return m, nil
}

// flattenClusterNodePoolsUpgradeSettings flattens an instance of ClusterNodePoolsUpgradeSettings from a JSON
// response object.
func flattenClusterNodePoolsUpgradeSettings(c *Client, i interface{}) *ClusterNodePoolsUpgradeSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodePoolsUpgradeSettings{}
	r.MaxSurge = dcl.FlattenInteger(m["maxSurge"])
	r.MaxUnavailable = dcl.FlattenInteger(m["maxUnavailable"])

	return r
}

// expandClusterLegacyAbacMap expands the contents of ClusterLegacyAbac into a JSON
// request object.
func expandClusterLegacyAbacMap(c *Client, f map[string]ClusterLegacyAbac) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterLegacyAbac(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterLegacyAbacSlice expands the contents of ClusterLegacyAbac into a JSON
// request object.
func expandClusterLegacyAbacSlice(c *Client, f []ClusterLegacyAbac) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterLegacyAbac(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterLegacyAbacMap flattens the contents of ClusterLegacyAbac from a JSON
// response object.
func flattenClusterLegacyAbacMap(c *Client, i interface{}) map[string]ClusterLegacyAbac {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterLegacyAbac{}
	}

	if len(a) == 0 {
		return map[string]ClusterLegacyAbac{}
	}

	items := make(map[string]ClusterLegacyAbac)
	for k, item := range a {
		items[k] = *flattenClusterLegacyAbac(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterLegacyAbacSlice flattens the contents of ClusterLegacyAbac from a JSON
// response object.
func flattenClusterLegacyAbacSlice(c *Client, i interface{}) []ClusterLegacyAbac {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterLegacyAbac{}
	}

	if len(a) == 0 {
		return []ClusterLegacyAbac{}
	}

	items := make([]ClusterLegacyAbac, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterLegacyAbac(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterLegacyAbac expands an instance of ClusterLegacyAbac into a JSON
// request object.
func expandClusterLegacyAbac(c *Client, f *ClusterLegacyAbac) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterLegacyAbac flattens an instance of ClusterLegacyAbac from a JSON
// response object.
func flattenClusterLegacyAbac(c *Client, i interface{}) *ClusterLegacyAbac {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterLegacyAbac{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterNetworkPolicyMap expands the contents of ClusterNetworkPolicy into a JSON
// request object.
func expandClusterNetworkPolicyMap(c *Client, f map[string]ClusterNetworkPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNetworkPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNetworkPolicySlice expands the contents of ClusterNetworkPolicy into a JSON
// request object.
func expandClusterNetworkPolicySlice(c *Client, f []ClusterNetworkPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNetworkPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNetworkPolicyMap flattens the contents of ClusterNetworkPolicy from a JSON
// response object.
func flattenClusterNetworkPolicyMap(c *Client, i interface{}) map[string]ClusterNetworkPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNetworkPolicy{}
	}

	if len(a) == 0 {
		return map[string]ClusterNetworkPolicy{}
	}

	items := make(map[string]ClusterNetworkPolicy)
	for k, item := range a {
		items[k] = *flattenClusterNetworkPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNetworkPolicySlice flattens the contents of ClusterNetworkPolicy from a JSON
// response object.
func flattenClusterNetworkPolicySlice(c *Client, i interface{}) []ClusterNetworkPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNetworkPolicy{}
	}

	if len(a) == 0 {
		return []ClusterNetworkPolicy{}
	}

	items := make([]ClusterNetworkPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNetworkPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNetworkPolicy expands an instance of ClusterNetworkPolicy into a JSON
// request object.
func expandClusterNetworkPolicy(c *Client, f *ClusterNetworkPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Provider; !dcl.IsEmptyValueIndirect(v) {
		m["provider"] = v
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterNetworkPolicy flattens an instance of ClusterNetworkPolicy from a JSON
// response object.
func flattenClusterNetworkPolicy(c *Client, i interface{}) *ClusterNetworkPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNetworkPolicy{}
	r.Provider = flattenClusterNetworkPolicyProviderEnum(m["provider"])
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterIPAllocationPolicyMap expands the contents of ClusterIPAllocationPolicy into a JSON
// request object.
func expandClusterIPAllocationPolicyMap(c *Client, f map[string]ClusterIPAllocationPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterIPAllocationPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterIPAllocationPolicySlice expands the contents of ClusterIPAllocationPolicy into a JSON
// request object.
func expandClusterIPAllocationPolicySlice(c *Client, f []ClusterIPAllocationPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterIPAllocationPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterIPAllocationPolicyMap flattens the contents of ClusterIPAllocationPolicy from a JSON
// response object.
func flattenClusterIPAllocationPolicyMap(c *Client, i interface{}) map[string]ClusterIPAllocationPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterIPAllocationPolicy{}
	}

	if len(a) == 0 {
		return map[string]ClusterIPAllocationPolicy{}
	}

	items := make(map[string]ClusterIPAllocationPolicy)
	for k, item := range a {
		items[k] = *flattenClusterIPAllocationPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterIPAllocationPolicySlice flattens the contents of ClusterIPAllocationPolicy from a JSON
// response object.
func flattenClusterIPAllocationPolicySlice(c *Client, i interface{}) []ClusterIPAllocationPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterIPAllocationPolicy{}
	}

	if len(a) == 0 {
		return []ClusterIPAllocationPolicy{}
	}

	items := make([]ClusterIPAllocationPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterIPAllocationPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterIPAllocationPolicy expands an instance of ClusterIPAllocationPolicy into a JSON
// request object.
func expandClusterIPAllocationPolicy(c *Client, f *ClusterIPAllocationPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.UseIPAliases; !dcl.IsEmptyValueIndirect(v) {
		m["useIpAliases"] = v
	}
	if v := f.CreateSubnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["createSubnetwork"] = v
	}
	if v := f.SubnetworkName; !dcl.IsEmptyValueIndirect(v) {
		m["subnetworkName"] = v
	}
	if v := f.ClusterSecondaryRangeName; !dcl.IsEmptyValueIndirect(v) {
		m["clusterSecondaryRangeName"] = v
	}
	if v := f.ServicesSecondaryRangeName; !dcl.IsEmptyValueIndirect(v) {
		m["servicesSecondaryRangeName"] = v
	}
	if v := f.ClusterIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["clusterIpv4CidrBlock"] = v
	}
	if v := f.NodeIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["nodeIpv4CidrBlock"] = v
	}
	if v := f.ServicesIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["servicesIpv4CidrBlock"] = v
	}
	if v := f.TPUIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["tpuIpv4CidrBlock"] = v
	}
	if v := f.ClusterIPv4Cidr; !dcl.IsEmptyValueIndirect(v) {
		m["clusterIpv4Cidr"] = v
	}
	if v := f.NodeIPv4Cidr; !dcl.IsEmptyValueIndirect(v) {
		m["nodeIpv4Cidr"] = v
	}
	if v := f.ServicesIPv4Cidr; !dcl.IsEmptyValueIndirect(v) {
		m["servicesIpv4Cidr"] = v
	}
	if v := f.UseRoutes; !dcl.IsEmptyValueIndirect(v) {
		m["useRoutes"] = v
	}

	return m, nil
}

// flattenClusterIPAllocationPolicy flattens an instance of ClusterIPAllocationPolicy from a JSON
// response object.
func flattenClusterIPAllocationPolicy(c *Client, i interface{}) *ClusterIPAllocationPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterIPAllocationPolicy{}
	r.UseIPAliases = dcl.FlattenBool(m["useIpAliases"])
	r.CreateSubnetwork = dcl.FlattenBool(m["createSubnetwork"])
	r.SubnetworkName = dcl.FlattenString(m["subnetworkName"])
	r.ClusterSecondaryRangeName = dcl.FlattenString(m["clusterSecondaryRangeName"])
	r.ServicesSecondaryRangeName = dcl.FlattenString(m["servicesSecondaryRangeName"])
	r.ClusterIPv4CidrBlock = dcl.FlattenString(m["clusterIpv4CidrBlock"])
	r.NodeIPv4CidrBlock = dcl.FlattenString(m["nodeIpv4CidrBlock"])
	r.ServicesIPv4CidrBlock = dcl.FlattenString(m["servicesIpv4CidrBlock"])
	r.TPUIPv4CidrBlock = dcl.FlattenString(m["tpuIpv4CidrBlock"])
	r.ClusterIPv4Cidr = dcl.FlattenString(m["clusterIpv4Cidr"])
	r.NodeIPv4Cidr = dcl.FlattenString(m["nodeIpv4Cidr"])
	r.ServicesIPv4Cidr = dcl.FlattenString(m["servicesIpv4Cidr"])
	r.UseRoutes = dcl.FlattenBool(m["useRoutes"])

	return r
}

// expandClusterMasterAuthorizedNetworksConfigMap expands the contents of ClusterMasterAuthorizedNetworksConfig into a JSON
// request object.
func expandClusterMasterAuthorizedNetworksConfigMap(c *Client, f map[string]ClusterMasterAuthorizedNetworksConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMasterAuthorizedNetworksConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMasterAuthorizedNetworksConfigSlice expands the contents of ClusterMasterAuthorizedNetworksConfig into a JSON
// request object.
func expandClusterMasterAuthorizedNetworksConfigSlice(c *Client, f []ClusterMasterAuthorizedNetworksConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMasterAuthorizedNetworksConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMasterAuthorizedNetworksConfigMap flattens the contents of ClusterMasterAuthorizedNetworksConfig from a JSON
// response object.
func flattenClusterMasterAuthorizedNetworksConfigMap(c *Client, i interface{}) map[string]ClusterMasterAuthorizedNetworksConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMasterAuthorizedNetworksConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterMasterAuthorizedNetworksConfig{}
	}

	items := make(map[string]ClusterMasterAuthorizedNetworksConfig)
	for k, item := range a {
		items[k] = *flattenClusterMasterAuthorizedNetworksConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMasterAuthorizedNetworksConfigSlice flattens the contents of ClusterMasterAuthorizedNetworksConfig from a JSON
// response object.
func flattenClusterMasterAuthorizedNetworksConfigSlice(c *Client, i interface{}) []ClusterMasterAuthorizedNetworksConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMasterAuthorizedNetworksConfig{}
	}

	if len(a) == 0 {
		return []ClusterMasterAuthorizedNetworksConfig{}
	}

	items := make([]ClusterMasterAuthorizedNetworksConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMasterAuthorizedNetworksConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMasterAuthorizedNetworksConfig expands an instance of ClusterMasterAuthorizedNetworksConfig into a JSON
// request object.
func expandClusterMasterAuthorizedNetworksConfig(c *Client, f *ClusterMasterAuthorizedNetworksConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v, err := expandClusterMasterAuthorizedNetworksConfigCidrBlocksSlice(c, f.CidrBlocks); err != nil {
		return nil, fmt.Errorf("error expanding CidrBlocks into cidrBlocks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cidrBlocks"] = v
	}

	return m, nil
}

// flattenClusterMasterAuthorizedNetworksConfig flattens an instance of ClusterMasterAuthorizedNetworksConfig from a JSON
// response object.
func flattenClusterMasterAuthorizedNetworksConfig(c *Client, i interface{}) *ClusterMasterAuthorizedNetworksConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMasterAuthorizedNetworksConfig{}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.CidrBlocks = flattenClusterMasterAuthorizedNetworksConfigCidrBlocksSlice(c, m["cidrBlocks"])

	return r
}

// expandClusterMasterAuthorizedNetworksConfigCidrBlocksMap expands the contents of ClusterMasterAuthorizedNetworksConfigCidrBlocks into a JSON
// request object.
func expandClusterMasterAuthorizedNetworksConfigCidrBlocksMap(c *Client, f map[string]ClusterMasterAuthorizedNetworksConfigCidrBlocks) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMasterAuthorizedNetworksConfigCidrBlocks(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMasterAuthorizedNetworksConfigCidrBlocksSlice expands the contents of ClusterMasterAuthorizedNetworksConfigCidrBlocks into a JSON
// request object.
func expandClusterMasterAuthorizedNetworksConfigCidrBlocksSlice(c *Client, f []ClusterMasterAuthorizedNetworksConfigCidrBlocks) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMasterAuthorizedNetworksConfigCidrBlocks(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMasterAuthorizedNetworksConfigCidrBlocksMap flattens the contents of ClusterMasterAuthorizedNetworksConfigCidrBlocks from a JSON
// response object.
func flattenClusterMasterAuthorizedNetworksConfigCidrBlocksMap(c *Client, i interface{}) map[string]ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMasterAuthorizedNetworksConfigCidrBlocks{}
	}

	if len(a) == 0 {
		return map[string]ClusterMasterAuthorizedNetworksConfigCidrBlocks{}
	}

	items := make(map[string]ClusterMasterAuthorizedNetworksConfigCidrBlocks)
	for k, item := range a {
		items[k] = *flattenClusterMasterAuthorizedNetworksConfigCidrBlocks(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMasterAuthorizedNetworksConfigCidrBlocksSlice flattens the contents of ClusterMasterAuthorizedNetworksConfigCidrBlocks from a JSON
// response object.
func flattenClusterMasterAuthorizedNetworksConfigCidrBlocksSlice(c *Client, i interface{}) []ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMasterAuthorizedNetworksConfigCidrBlocks{}
	}

	if len(a) == 0 {
		return []ClusterMasterAuthorizedNetworksConfigCidrBlocks{}
	}

	items := make([]ClusterMasterAuthorizedNetworksConfigCidrBlocks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMasterAuthorizedNetworksConfigCidrBlocks(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMasterAuthorizedNetworksConfigCidrBlocks expands an instance of ClusterMasterAuthorizedNetworksConfigCidrBlocks into a JSON
// request object.
func expandClusterMasterAuthorizedNetworksConfigCidrBlocks(c *Client, f *ClusterMasterAuthorizedNetworksConfigCidrBlocks) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["cidrBlock"] = v
	}

	return m, nil
}

// flattenClusterMasterAuthorizedNetworksConfigCidrBlocks flattens an instance of ClusterMasterAuthorizedNetworksConfigCidrBlocks from a JSON
// response object.
func flattenClusterMasterAuthorizedNetworksConfigCidrBlocks(c *Client, i interface{}) *ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMasterAuthorizedNetworksConfigCidrBlocks{}
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.CidrBlock = dcl.FlattenString(m["cidrBlock"])

	return r
}

// expandClusterBinaryAuthorizationMap expands the contents of ClusterBinaryAuthorization into a JSON
// request object.
func expandClusterBinaryAuthorizationMap(c *Client, f map[string]ClusterBinaryAuthorization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterBinaryAuthorization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterBinaryAuthorizationSlice expands the contents of ClusterBinaryAuthorization into a JSON
// request object.
func expandClusterBinaryAuthorizationSlice(c *Client, f []ClusterBinaryAuthorization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterBinaryAuthorization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterBinaryAuthorizationMap flattens the contents of ClusterBinaryAuthorization from a JSON
// response object.
func flattenClusterBinaryAuthorizationMap(c *Client, i interface{}) map[string]ClusterBinaryAuthorization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterBinaryAuthorization{}
	}

	if len(a) == 0 {
		return map[string]ClusterBinaryAuthorization{}
	}

	items := make(map[string]ClusterBinaryAuthorization)
	for k, item := range a {
		items[k] = *flattenClusterBinaryAuthorization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterBinaryAuthorizationSlice flattens the contents of ClusterBinaryAuthorization from a JSON
// response object.
func flattenClusterBinaryAuthorizationSlice(c *Client, i interface{}) []ClusterBinaryAuthorization {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterBinaryAuthorization{}
	}

	if len(a) == 0 {
		return []ClusterBinaryAuthorization{}
	}

	items := make([]ClusterBinaryAuthorization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterBinaryAuthorization(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterBinaryAuthorization expands an instance of ClusterBinaryAuthorization into a JSON
// request object.
func expandClusterBinaryAuthorization(c *Client, f *ClusterBinaryAuthorization) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterBinaryAuthorization flattens an instance of ClusterBinaryAuthorization from a JSON
// response object.
func flattenClusterBinaryAuthorization(c *Client, i interface{}) *ClusterBinaryAuthorization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterBinaryAuthorization{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterAutoscalingMap expands the contents of ClusterAutoscaling into a JSON
// request object.
func expandClusterAutoscalingMap(c *Client, f map[string]ClusterAutoscaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAutoscalingSlice expands the contents of ClusterAutoscaling into a JSON
// request object.
func expandClusterAutoscalingSlice(c *Client, f []ClusterAutoscaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAutoscalingMap flattens the contents of ClusterAutoscaling from a JSON
// response object.
func flattenClusterAutoscalingMap(c *Client, i interface{}) map[string]ClusterAutoscaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAutoscaling{}
	}

	if len(a) == 0 {
		return map[string]ClusterAutoscaling{}
	}

	items := make(map[string]ClusterAutoscaling)
	for k, item := range a {
		items[k] = *flattenClusterAutoscaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAutoscalingSlice flattens the contents of ClusterAutoscaling from a JSON
// response object.
func flattenClusterAutoscalingSlice(c *Client, i interface{}) []ClusterAutoscaling {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAutoscaling{}
	}

	if len(a) == 0 {
		return []ClusterAutoscaling{}
	}

	items := make([]ClusterAutoscaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAutoscaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAutoscaling expands an instance of ClusterAutoscaling into a JSON
// request object.
func expandClusterAutoscaling(c *Client, f *ClusterAutoscaling) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.EnableNodeAutoprovisioning; !dcl.IsEmptyValueIndirect(v) {
		m["enableNodeAutoprovisioning"] = v
	}
	if v, err := expandClusterAutoscalingResourceLimitsSlice(c, f.ResourceLimits); err != nil {
		return nil, fmt.Errorf("error expanding ResourceLimits into resourceLimits: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resourceLimits"] = v
	}
	if v, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaults(c, f.AutoprovisioningNodePoolDefaults); err != nil {
		return nil, fmt.Errorf("error expanding AutoprovisioningNodePoolDefaults into autoprovisioningNodePoolDefaults: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoprovisioningNodePoolDefaults"] = v
	}
	if v := f.AutoprovisioningLocations; !dcl.IsEmptyValueIndirect(v) {
		m["autoprovisioningLocations"] = v
	}

	return m, nil
}

// flattenClusterAutoscaling flattens an instance of ClusterAutoscaling from a JSON
// response object.
func flattenClusterAutoscaling(c *Client, i interface{}) *ClusterAutoscaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAutoscaling{}
	r.EnableNodeAutoprovisioning = dcl.FlattenBool(m["enableNodeAutoprovisioning"])
	r.ResourceLimits = flattenClusterAutoscalingResourceLimitsSlice(c, m["resourceLimits"])
	r.AutoprovisioningNodePoolDefaults = flattenClusterAutoscalingAutoprovisioningNodePoolDefaults(c, m["autoprovisioningNodePoolDefaults"])
	r.AutoprovisioningLocations = dcl.FlattenStringSlice(m["autoprovisioningLocations"])

	return r
}

// expandClusterAutoscalingResourceLimitsMap expands the contents of ClusterAutoscalingResourceLimits into a JSON
// request object.
func expandClusterAutoscalingResourceLimitsMap(c *Client, f map[string]ClusterAutoscalingResourceLimits) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAutoscalingResourceLimits(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAutoscalingResourceLimitsSlice expands the contents of ClusterAutoscalingResourceLimits into a JSON
// request object.
func expandClusterAutoscalingResourceLimitsSlice(c *Client, f []ClusterAutoscalingResourceLimits) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAutoscalingResourceLimits(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAutoscalingResourceLimitsMap flattens the contents of ClusterAutoscalingResourceLimits from a JSON
// response object.
func flattenClusterAutoscalingResourceLimitsMap(c *Client, i interface{}) map[string]ClusterAutoscalingResourceLimits {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAutoscalingResourceLimits{}
	}

	if len(a) == 0 {
		return map[string]ClusterAutoscalingResourceLimits{}
	}

	items := make(map[string]ClusterAutoscalingResourceLimits)
	for k, item := range a {
		items[k] = *flattenClusterAutoscalingResourceLimits(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAutoscalingResourceLimitsSlice flattens the contents of ClusterAutoscalingResourceLimits from a JSON
// response object.
func flattenClusterAutoscalingResourceLimitsSlice(c *Client, i interface{}) []ClusterAutoscalingResourceLimits {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAutoscalingResourceLimits{}
	}

	if len(a) == 0 {
		return []ClusterAutoscalingResourceLimits{}
	}

	items := make([]ClusterAutoscalingResourceLimits, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAutoscalingResourceLimits(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAutoscalingResourceLimits expands an instance of ClusterAutoscalingResourceLimits into a JSON
// request object.
func expandClusterAutoscalingResourceLimits(c *Client, f *ClusterAutoscalingResourceLimits) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ResourceType; !dcl.IsEmptyValueIndirect(v) {
		m["resourceType"] = v
	}
	if v := f.Minimum; !dcl.IsEmptyValueIndirect(v) {
		m["minimum"] = v
	}
	if v := f.Maximum; !dcl.IsEmptyValueIndirect(v) {
		m["maximum"] = v
	}

	return m, nil
}

// flattenClusterAutoscalingResourceLimits flattens an instance of ClusterAutoscalingResourceLimits from a JSON
// response object.
func flattenClusterAutoscalingResourceLimits(c *Client, i interface{}) *ClusterAutoscalingResourceLimits {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAutoscalingResourceLimits{}
	r.ResourceType = dcl.FlattenString(m["resourceType"])
	r.Minimum = dcl.FlattenInteger(m["minimum"])
	r.Maximum = dcl.FlattenInteger(m["maximum"])

	return r
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsMap expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaults into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsMap(c *Client, f map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaults) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaults(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsSlice expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaults into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsSlice(c *Client, f []ClusterAutoscalingAutoprovisioningNodePoolDefaults) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaults(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsMap flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaults from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsMap(c *Client, i interface{}) map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaults{}
	}

	if len(a) == 0 {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaults{}
	}

	items := make(map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaults)
	for k, item := range a {
		items[k] = *flattenClusterAutoscalingAutoprovisioningNodePoolDefaults(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsSlice flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaults from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsSlice(c *Client, i interface{}) []ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaults{}
	}

	if len(a) == 0 {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaults{}
	}

	items := make([]ClusterAutoscalingAutoprovisioningNodePoolDefaults, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAutoscalingAutoprovisioningNodePoolDefaults(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaults expands an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaults into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaults(c *Client, f *ClusterAutoscalingAutoprovisioningNodePoolDefaults) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.OAuthScopes; !dcl.IsEmptyValueIndirect(v) {
		m["oauthScopes"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, f.UpgradeSettings); err != nil {
		return nil, fmt.Errorf("error expanding UpgradeSettings into upgradeSettings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["upgradeSettings"] = v
	}
	if v, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, f.Management); err != nil {
		return nil, fmt.Errorf("error expanding Management into management: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["management"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}
	if v := f.DiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskSizeGb"] = v
	}
	if v := f.DiskType; !dcl.IsEmptyValueIndirect(v) {
		m["diskType"] = v
	}
	if v, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, f.ShieldedInstanceConfig); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedInstanceConfig into shieldedInstanceConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["shieldedInstanceConfig"] = v
	}
	if v := f.BootDiskKmsKey; !dcl.IsEmptyValueIndirect(v) {
		m["bootDiskKmsKey"] = v
	}

	return m, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaults flattens an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaults from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaults(c *Client, i interface{}) *ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAutoscalingAutoprovisioningNodePoolDefaults{}
	r.OAuthScopes = dcl.FlattenStringSlice(m["oauthScopes"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.UpgradeSettings = flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, m["upgradeSettings"])
	r.Management = flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, m["management"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.DiskType = dcl.FlattenString(m["diskType"])
	r.ShieldedInstanceConfig = flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, m["shieldedInstanceConfig"])
	r.BootDiskKmsKey = dcl.FlattenString(m["bootDiskKmsKey"])

	return r
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsMap expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsMap(c *Client, f map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsSlice expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsSlice(c *Client, f []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsMap flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsMap(c *Client, i interface{}) map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{}
	}

	if len(a) == 0 {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{}
	}

	items := make(map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings)
	for k, item := range a {
		items[k] = *flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsSlice flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsSlice(c *Client, i interface{}) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{}
	}

	if len(a) == 0 {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{}
	}

	items := make([]ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings expands an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c *Client, f *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.MaxSurge; !dcl.IsEmptyValueIndirect(v) {
		m["maxSurge"] = v
	}
	if v := f.MaxUnavailable; !dcl.IsEmptyValueIndirect(v) {
		m["maxUnavailable"] = v
	}

	return m, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings flattens an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c *Client, i interface{}) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{}
	r.MaxSurge = dcl.FlattenInteger(m["maxSurge"])
	r.MaxUnavailable = dcl.FlattenInteger(m["maxUnavailable"])

	return r
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementMap expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementMap(c *Client, f map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementSlice expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementSlice(c *Client, f []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementMap flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementMap(c *Client, i interface{}) map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{}
	}

	if len(a) == 0 {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{}
	}

	items := make(map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement)
	for k, item := range a {
		items[k] = *flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementSlice flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementSlice(c *Client, i interface{}) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{}
	}

	if len(a) == 0 {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{}
	}

	items := make([]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement expands an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c *Client, f *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.AutoUpgrade; !dcl.IsEmptyValueIndirect(v) {
		m["autoUpgrade"] = v
	}
	if v := f.AutoRepair; !dcl.IsEmptyValueIndirect(v) {
		m["autoRepair"] = v
	}
	if v, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, f.UpgradeOptions); err != nil {
		return nil, fmt.Errorf("error expanding UpgradeOptions into upgradeOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["upgradeOptions"] = v
	}

	return m, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement flattens an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c *Client, i interface{}) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{}
	r.AutoUpgrade = dcl.FlattenBool(m["autoUpgrade"])
	r.AutoRepair = dcl.FlattenBool(m["autoRepair"])
	r.UpgradeOptions = flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, m["upgradeOptions"])

	return r
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsMap expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsMap(c *Client, f map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsSlice expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsSlice(c *Client, f []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsMap flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsMap(c *Client, i interface{}) map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{}
	}

	if len(a) == 0 {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{}
	}

	items := make(map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions)
	for k, item := range a {
		items[k] = *flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsSlice flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsSlice(c *Client, i interface{}) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{}
	}

	if len(a) == 0 {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{}
	}

	items := make([]ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions expands an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c *Client, f *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.AutoUpgradeStartTime; !dcl.IsEmptyValueIndirect(v) {
		m["autoUpgradeStartTime"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions flattens an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(c *Client, i interface{}) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{}
	r.AutoUpgradeStartTime = dcl.FlattenString(m["autoUpgradeStartTime"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigMap expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigMap(c *Client, f map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigSlice expands the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigSlice(c *Client, f []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigMap flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigMap(c *Client, i interface{}) map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{}
	}

	items := make(map[string]ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig)
	for k, item := range a {
		items[k] = *flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigSlice flattens the contents of ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigSlice(c *Client, i interface{}) []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return []ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{}
	}

	items := make([]ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig expands an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig into a JSON
// request object.
func expandClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c *Client, f *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.EnableSecureBoot; !dcl.IsEmptyValueIndirect(v) {
		m["enableSecureBoot"] = v
	}
	if v := f.EnableIntegrityMonitoring; !dcl.IsEmptyValueIndirect(v) {
		m["enableIntegrityMonitoring"] = v
	}

	return m, nil
}

// flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig flattens an instance of ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig from a JSON
// response object.
func flattenClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(c *Client, i interface{}) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{}
	r.EnableSecureBoot = dcl.FlattenBool(m["enableSecureBoot"])
	r.EnableIntegrityMonitoring = dcl.FlattenBool(m["enableIntegrityMonitoring"])

	return r
}

// expandClusterNetworkConfigMap expands the contents of ClusterNetworkConfig into a JSON
// request object.
func expandClusterNetworkConfigMap(c *Client, f map[string]ClusterNetworkConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNetworkConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNetworkConfigSlice expands the contents of ClusterNetworkConfig into a JSON
// request object.
func expandClusterNetworkConfigSlice(c *Client, f []ClusterNetworkConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNetworkConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNetworkConfigMap flattens the contents of ClusterNetworkConfig from a JSON
// response object.
func flattenClusterNetworkConfigMap(c *Client, i interface{}) map[string]ClusterNetworkConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNetworkConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNetworkConfig{}
	}

	items := make(map[string]ClusterNetworkConfig)
	for k, item := range a {
		items[k] = *flattenClusterNetworkConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNetworkConfigSlice flattens the contents of ClusterNetworkConfig from a JSON
// response object.
func flattenClusterNetworkConfigSlice(c *Client, i interface{}) []ClusterNetworkConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNetworkConfig{}
	}

	if len(a) == 0 {
		return []ClusterNetworkConfig{}
	}

	items := make([]ClusterNetworkConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNetworkConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNetworkConfig expands an instance of ClusterNetworkConfig into a JSON
// request object.
func expandClusterNetworkConfig(c *Client, f *ClusterNetworkConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.Subnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["subnetwork"] = v
	}
	if v := f.EnableIntraNodeVisibility; !dcl.IsEmptyValueIndirect(v) {
		m["enableIntraNodeVisibility"] = v
	}
	if v, err := expandClusterNetworkConfigDefaultSnatStatus(c, f.DefaultSnatStatus); err != nil {
		return nil, fmt.Errorf("error expanding DefaultSnatStatus into defaultSnatStatus: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultSnatStatus"] = v
	}
	if v := f.PrivateIPv6GoogleAccess; !dcl.IsEmptyValueIndirect(v) {
		m["privateIpv6GoogleAccess"] = v
	}

	return m, nil
}

// flattenClusterNetworkConfig flattens an instance of ClusterNetworkConfig from a JSON
// response object.
func flattenClusterNetworkConfig(c *Client, i interface{}) *ClusterNetworkConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNetworkConfig{}
	r.Network = dcl.FlattenString(m["network"])
	r.Subnetwork = dcl.FlattenString(m["subnetwork"])
	r.EnableIntraNodeVisibility = dcl.FlattenBool(m["enableIntraNodeVisibility"])
	r.DefaultSnatStatus = flattenClusterNetworkConfigDefaultSnatStatus(c, m["defaultSnatStatus"])
	r.PrivateIPv6GoogleAccess = flattenClusterNetworkConfigPrivateIPv6GoogleAccessEnum(m["privateIpv6GoogleAccess"])

	return r
}

// expandClusterNetworkConfigDefaultSnatStatusMap expands the contents of ClusterNetworkConfigDefaultSnatStatus into a JSON
// request object.
func expandClusterNetworkConfigDefaultSnatStatusMap(c *Client, f map[string]ClusterNetworkConfigDefaultSnatStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNetworkConfigDefaultSnatStatus(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNetworkConfigDefaultSnatStatusSlice expands the contents of ClusterNetworkConfigDefaultSnatStatus into a JSON
// request object.
func expandClusterNetworkConfigDefaultSnatStatusSlice(c *Client, f []ClusterNetworkConfigDefaultSnatStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNetworkConfigDefaultSnatStatus(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNetworkConfigDefaultSnatStatusMap flattens the contents of ClusterNetworkConfigDefaultSnatStatus from a JSON
// response object.
func flattenClusterNetworkConfigDefaultSnatStatusMap(c *Client, i interface{}) map[string]ClusterNetworkConfigDefaultSnatStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNetworkConfigDefaultSnatStatus{}
	}

	if len(a) == 0 {
		return map[string]ClusterNetworkConfigDefaultSnatStatus{}
	}

	items := make(map[string]ClusterNetworkConfigDefaultSnatStatus)
	for k, item := range a {
		items[k] = *flattenClusterNetworkConfigDefaultSnatStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNetworkConfigDefaultSnatStatusSlice flattens the contents of ClusterNetworkConfigDefaultSnatStatus from a JSON
// response object.
func flattenClusterNetworkConfigDefaultSnatStatusSlice(c *Client, i interface{}) []ClusterNetworkConfigDefaultSnatStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNetworkConfigDefaultSnatStatus{}
	}

	if len(a) == 0 {
		return []ClusterNetworkConfigDefaultSnatStatus{}
	}

	items := make([]ClusterNetworkConfigDefaultSnatStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNetworkConfigDefaultSnatStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNetworkConfigDefaultSnatStatus expands an instance of ClusterNetworkConfigDefaultSnatStatus into a JSON
// request object.
func expandClusterNetworkConfigDefaultSnatStatus(c *Client, f *ClusterNetworkConfigDefaultSnatStatus) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}

	return m, nil
}

// flattenClusterNetworkConfigDefaultSnatStatus flattens an instance of ClusterNetworkConfigDefaultSnatStatus from a JSON
// response object.
func flattenClusterNetworkConfigDefaultSnatStatus(c *Client, i interface{}) *ClusterNetworkConfigDefaultSnatStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNetworkConfigDefaultSnatStatus{}
	r.Disabled = dcl.FlattenBool(m["disabled"])

	return r
}

// expandClusterMaintenancePolicyMap expands the contents of ClusterMaintenancePolicy into a JSON
// request object.
func expandClusterMaintenancePolicyMap(c *Client, f map[string]ClusterMaintenancePolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMaintenancePolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMaintenancePolicySlice expands the contents of ClusterMaintenancePolicy into a JSON
// request object.
func expandClusterMaintenancePolicySlice(c *Client, f []ClusterMaintenancePolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMaintenancePolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMaintenancePolicyMap flattens the contents of ClusterMaintenancePolicy from a JSON
// response object.
func flattenClusterMaintenancePolicyMap(c *Client, i interface{}) map[string]ClusterMaintenancePolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMaintenancePolicy{}
	}

	if len(a) == 0 {
		return map[string]ClusterMaintenancePolicy{}
	}

	items := make(map[string]ClusterMaintenancePolicy)
	for k, item := range a {
		items[k] = *flattenClusterMaintenancePolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMaintenancePolicySlice flattens the contents of ClusterMaintenancePolicy from a JSON
// response object.
func flattenClusterMaintenancePolicySlice(c *Client, i interface{}) []ClusterMaintenancePolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMaintenancePolicy{}
	}

	if len(a) == 0 {
		return []ClusterMaintenancePolicy{}
	}

	items := make([]ClusterMaintenancePolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMaintenancePolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMaintenancePolicy expands an instance of ClusterMaintenancePolicy into a JSON
// request object.
func expandClusterMaintenancePolicy(c *Client, f *ClusterMaintenancePolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandClusterMaintenancePolicyWindow(c, f.Window); err != nil {
		return nil, fmt.Errorf("error expanding Window into window: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["window"] = v
	}
	if v := f.ResourceVersion; !dcl.IsEmptyValueIndirect(v) {
		m["resourceVersion"] = v
	}

	return m, nil
}

// flattenClusterMaintenancePolicy flattens an instance of ClusterMaintenancePolicy from a JSON
// response object.
func flattenClusterMaintenancePolicy(c *Client, i interface{}) *ClusterMaintenancePolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMaintenancePolicy{}
	r.Window = flattenClusterMaintenancePolicyWindow(c, m["window"])
	r.ResourceVersion = dcl.FlattenString(m["resourceVersion"])

	return r
}

// expandClusterMaintenancePolicyWindowMap expands the contents of ClusterMaintenancePolicyWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowMap(c *Client, f map[string]ClusterMaintenancePolicyWindow) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMaintenancePolicyWindow(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMaintenancePolicyWindowSlice expands the contents of ClusterMaintenancePolicyWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowSlice(c *Client, f []ClusterMaintenancePolicyWindow) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMaintenancePolicyWindow(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMaintenancePolicyWindowMap flattens the contents of ClusterMaintenancePolicyWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowMap(c *Client, i interface{}) map[string]ClusterMaintenancePolicyWindow {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMaintenancePolicyWindow{}
	}

	if len(a) == 0 {
		return map[string]ClusterMaintenancePolicyWindow{}
	}

	items := make(map[string]ClusterMaintenancePolicyWindow)
	for k, item := range a {
		items[k] = *flattenClusterMaintenancePolicyWindow(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMaintenancePolicyWindowSlice flattens the contents of ClusterMaintenancePolicyWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowSlice(c *Client, i interface{}) []ClusterMaintenancePolicyWindow {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMaintenancePolicyWindow{}
	}

	if len(a) == 0 {
		return []ClusterMaintenancePolicyWindow{}
	}

	items := make([]ClusterMaintenancePolicyWindow, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMaintenancePolicyWindow(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMaintenancePolicyWindow expands an instance of ClusterMaintenancePolicyWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindow(c *Client, f *ClusterMaintenancePolicyWindow) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, f.DailyMaintenanceWindow); err != nil {
		return nil, fmt.Errorf("error expanding DailyMaintenanceWindow into dailyMaintenanceWindow: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dailyMaintenanceWindow"] = v
	}
	if v, err := expandClusterMaintenancePolicyWindowRecurringWindow(c, f.RecurringWindow); err != nil {
		return nil, fmt.Errorf("error expanding RecurringWindow into recurringWindow: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["recurringWindow"] = v
	}
	if v := f.MaintenanceExclusions; !dcl.IsEmptyValueIndirect(v) {
		m["maintenanceExclusions"] = v
	}

	return m, nil
}

// flattenClusterMaintenancePolicyWindow flattens an instance of ClusterMaintenancePolicyWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindow(c *Client, i interface{}) *ClusterMaintenancePolicyWindow {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMaintenancePolicyWindow{}
	r.DailyMaintenanceWindow = flattenClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, m["dailyMaintenanceWindow"])
	r.RecurringWindow = flattenClusterMaintenancePolicyWindowRecurringWindow(c, m["recurringWindow"])
	r.MaintenanceExclusions = dcl.FlattenKeyValuePairs(m["maintenanceExclusions"])

	return r
}

// expandClusterMaintenancePolicyWindowDailyMaintenanceWindowMap expands the contents of ClusterMaintenancePolicyWindowDailyMaintenanceWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowDailyMaintenanceWindowMap(c *Client, f map[string]ClusterMaintenancePolicyWindowDailyMaintenanceWindow) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMaintenancePolicyWindowDailyMaintenanceWindowSlice expands the contents of ClusterMaintenancePolicyWindowDailyMaintenanceWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowDailyMaintenanceWindowSlice(c *Client, f []ClusterMaintenancePolicyWindowDailyMaintenanceWindow) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMaintenancePolicyWindowDailyMaintenanceWindowMap flattens the contents of ClusterMaintenancePolicyWindowDailyMaintenanceWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowDailyMaintenanceWindowMap(c *Client, i interface{}) map[string]ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMaintenancePolicyWindowDailyMaintenanceWindow{}
	}

	if len(a) == 0 {
		return map[string]ClusterMaintenancePolicyWindowDailyMaintenanceWindow{}
	}

	items := make(map[string]ClusterMaintenancePolicyWindowDailyMaintenanceWindow)
	for k, item := range a {
		items[k] = *flattenClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMaintenancePolicyWindowDailyMaintenanceWindowSlice flattens the contents of ClusterMaintenancePolicyWindowDailyMaintenanceWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowDailyMaintenanceWindowSlice(c *Client, i interface{}) []ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMaintenancePolicyWindowDailyMaintenanceWindow{}
	}

	if len(a) == 0 {
		return []ClusterMaintenancePolicyWindowDailyMaintenanceWindow{}
	}

	items := make([]ClusterMaintenancePolicyWindowDailyMaintenanceWindow, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMaintenancePolicyWindowDailyMaintenanceWindow expands an instance of ClusterMaintenancePolicyWindowDailyMaintenanceWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowDailyMaintenanceWindow(c *Client, f *ClusterMaintenancePolicyWindowDailyMaintenanceWindow) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.StartTime; !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.Duration; !dcl.IsEmptyValueIndirect(v) {
		m["duration"] = v
	}

	return m, nil
}

// flattenClusterMaintenancePolicyWindowDailyMaintenanceWindow flattens an instance of ClusterMaintenancePolicyWindowDailyMaintenanceWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowDailyMaintenanceWindow(c *Client, i interface{}) *ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMaintenancePolicyWindowDailyMaintenanceWindow{}
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.Duration = dcl.FlattenString(m["duration"])

	return r
}

// expandClusterMaintenancePolicyWindowRecurringWindowMap expands the contents of ClusterMaintenancePolicyWindowRecurringWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowRecurringWindowMap(c *Client, f map[string]ClusterMaintenancePolicyWindowRecurringWindow) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMaintenancePolicyWindowRecurringWindow(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMaintenancePolicyWindowRecurringWindowSlice expands the contents of ClusterMaintenancePolicyWindowRecurringWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowRecurringWindowSlice(c *Client, f []ClusterMaintenancePolicyWindowRecurringWindow) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMaintenancePolicyWindowRecurringWindow(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMaintenancePolicyWindowRecurringWindowMap flattens the contents of ClusterMaintenancePolicyWindowRecurringWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowRecurringWindowMap(c *Client, i interface{}) map[string]ClusterMaintenancePolicyWindowRecurringWindow {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMaintenancePolicyWindowRecurringWindow{}
	}

	if len(a) == 0 {
		return map[string]ClusterMaintenancePolicyWindowRecurringWindow{}
	}

	items := make(map[string]ClusterMaintenancePolicyWindowRecurringWindow)
	for k, item := range a {
		items[k] = *flattenClusterMaintenancePolicyWindowRecurringWindow(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMaintenancePolicyWindowRecurringWindowSlice flattens the contents of ClusterMaintenancePolicyWindowRecurringWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowRecurringWindowSlice(c *Client, i interface{}) []ClusterMaintenancePolicyWindowRecurringWindow {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMaintenancePolicyWindowRecurringWindow{}
	}

	if len(a) == 0 {
		return []ClusterMaintenancePolicyWindowRecurringWindow{}
	}

	items := make([]ClusterMaintenancePolicyWindowRecurringWindow, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMaintenancePolicyWindowRecurringWindow(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMaintenancePolicyWindowRecurringWindow expands an instance of ClusterMaintenancePolicyWindowRecurringWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowRecurringWindow(c *Client, f *ClusterMaintenancePolicyWindowRecurringWindow) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandClusterMaintenancePolicyWindowRecurringWindowWindow(c, f.Window); err != nil {
		return nil, fmt.Errorf("error expanding Window into window: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["window"] = v
	}
	if v := f.Recurrence; !dcl.IsEmptyValueIndirect(v) {
		m["recurrence"] = v
	}

	return m, nil
}

// flattenClusterMaintenancePolicyWindowRecurringWindow flattens an instance of ClusterMaintenancePolicyWindowRecurringWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowRecurringWindow(c *Client, i interface{}) *ClusterMaintenancePolicyWindowRecurringWindow {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMaintenancePolicyWindowRecurringWindow{}
	r.Window = flattenClusterMaintenancePolicyWindowRecurringWindowWindow(c, m["window"])
	r.Recurrence = dcl.FlattenString(m["recurrence"])

	return r
}

// expandClusterMaintenancePolicyWindowRecurringWindowWindowMap expands the contents of ClusterMaintenancePolicyWindowRecurringWindowWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowRecurringWindowWindowMap(c *Client, f map[string]ClusterMaintenancePolicyWindowRecurringWindowWindow) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterMaintenancePolicyWindowRecurringWindowWindow(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterMaintenancePolicyWindowRecurringWindowWindowSlice expands the contents of ClusterMaintenancePolicyWindowRecurringWindowWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowRecurringWindowWindowSlice(c *Client, f []ClusterMaintenancePolicyWindowRecurringWindowWindow) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterMaintenancePolicyWindowRecurringWindowWindow(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterMaintenancePolicyWindowRecurringWindowWindowMap flattens the contents of ClusterMaintenancePolicyWindowRecurringWindowWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowRecurringWindowWindowMap(c *Client, i interface{}) map[string]ClusterMaintenancePolicyWindowRecurringWindowWindow {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterMaintenancePolicyWindowRecurringWindowWindow{}
	}

	if len(a) == 0 {
		return map[string]ClusterMaintenancePolicyWindowRecurringWindowWindow{}
	}

	items := make(map[string]ClusterMaintenancePolicyWindowRecurringWindowWindow)
	for k, item := range a {
		items[k] = *flattenClusterMaintenancePolicyWindowRecurringWindowWindow(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterMaintenancePolicyWindowRecurringWindowWindowSlice flattens the contents of ClusterMaintenancePolicyWindowRecurringWindowWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowRecurringWindowWindowSlice(c *Client, i interface{}) []ClusterMaintenancePolicyWindowRecurringWindowWindow {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterMaintenancePolicyWindowRecurringWindowWindow{}
	}

	if len(a) == 0 {
		return []ClusterMaintenancePolicyWindowRecurringWindowWindow{}
	}

	items := make([]ClusterMaintenancePolicyWindowRecurringWindowWindow, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterMaintenancePolicyWindowRecurringWindowWindow(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterMaintenancePolicyWindowRecurringWindowWindow expands an instance of ClusterMaintenancePolicyWindowRecurringWindowWindow into a JSON
// request object.
func expandClusterMaintenancePolicyWindowRecurringWindowWindow(c *Client, f *ClusterMaintenancePolicyWindowRecurringWindowWindow) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.StartTime; !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.EndTime; !dcl.IsEmptyValueIndirect(v) {
		m["endTime"] = v
	}

	return m, nil
}

// flattenClusterMaintenancePolicyWindowRecurringWindowWindow flattens an instance of ClusterMaintenancePolicyWindowRecurringWindowWindow from a JSON
// response object.
func flattenClusterMaintenancePolicyWindowRecurringWindowWindow(c *Client, i interface{}) *ClusterMaintenancePolicyWindowRecurringWindowWindow {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterMaintenancePolicyWindowRecurringWindowWindow{}
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.EndTime = dcl.FlattenString(m["endTime"])

	return r
}

// expandClusterDefaultMaxPodsConstraintMap expands the contents of ClusterDefaultMaxPodsConstraint into a JSON
// request object.
func expandClusterDefaultMaxPodsConstraintMap(c *Client, f map[string]ClusterDefaultMaxPodsConstraint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterDefaultMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterDefaultMaxPodsConstraintSlice expands the contents of ClusterDefaultMaxPodsConstraint into a JSON
// request object.
func expandClusterDefaultMaxPodsConstraintSlice(c *Client, f []ClusterDefaultMaxPodsConstraint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterDefaultMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterDefaultMaxPodsConstraintMap flattens the contents of ClusterDefaultMaxPodsConstraint from a JSON
// response object.
func flattenClusterDefaultMaxPodsConstraintMap(c *Client, i interface{}) map[string]ClusterDefaultMaxPodsConstraint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterDefaultMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return map[string]ClusterDefaultMaxPodsConstraint{}
	}

	items := make(map[string]ClusterDefaultMaxPodsConstraint)
	for k, item := range a {
		items[k] = *flattenClusterDefaultMaxPodsConstraint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterDefaultMaxPodsConstraintSlice flattens the contents of ClusterDefaultMaxPodsConstraint from a JSON
// response object.
func flattenClusterDefaultMaxPodsConstraintSlice(c *Client, i interface{}) []ClusterDefaultMaxPodsConstraint {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterDefaultMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return []ClusterDefaultMaxPodsConstraint{}
	}

	items := make([]ClusterDefaultMaxPodsConstraint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterDefaultMaxPodsConstraint(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterDefaultMaxPodsConstraint expands an instance of ClusterDefaultMaxPodsConstraint into a JSON
// request object.
func expandClusterDefaultMaxPodsConstraint(c *Client, f *ClusterDefaultMaxPodsConstraint) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.MaxPodsPerNode; !dcl.IsEmptyValueIndirect(v) {
		m["maxPodsPerNode"] = v
	}

	return m, nil
}

// flattenClusterDefaultMaxPodsConstraint flattens an instance of ClusterDefaultMaxPodsConstraint from a JSON
// response object.
func flattenClusterDefaultMaxPodsConstraint(c *Client, i interface{}) *ClusterDefaultMaxPodsConstraint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterDefaultMaxPodsConstraint{}
	r.MaxPodsPerNode = dcl.FlattenString(m["maxPodsPerNode"])

	return r
}

// expandClusterResourceUsageExportConfigMap expands the contents of ClusterResourceUsageExportConfig into a JSON
// request object.
func expandClusterResourceUsageExportConfigMap(c *Client, f map[string]ClusterResourceUsageExportConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterResourceUsageExportConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterResourceUsageExportConfigSlice expands the contents of ClusterResourceUsageExportConfig into a JSON
// request object.
func expandClusterResourceUsageExportConfigSlice(c *Client, f []ClusterResourceUsageExportConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterResourceUsageExportConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterResourceUsageExportConfigMap flattens the contents of ClusterResourceUsageExportConfig from a JSON
// response object.
func flattenClusterResourceUsageExportConfigMap(c *Client, i interface{}) map[string]ClusterResourceUsageExportConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterResourceUsageExportConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterResourceUsageExportConfig{}
	}

	items := make(map[string]ClusterResourceUsageExportConfig)
	for k, item := range a {
		items[k] = *flattenClusterResourceUsageExportConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterResourceUsageExportConfigSlice flattens the contents of ClusterResourceUsageExportConfig from a JSON
// response object.
func flattenClusterResourceUsageExportConfigSlice(c *Client, i interface{}) []ClusterResourceUsageExportConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterResourceUsageExportConfig{}
	}

	if len(a) == 0 {
		return []ClusterResourceUsageExportConfig{}
	}

	items := make([]ClusterResourceUsageExportConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterResourceUsageExportConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterResourceUsageExportConfig expands an instance of ClusterResourceUsageExportConfig into a JSON
// request object.
func expandClusterResourceUsageExportConfig(c *Client, f *ClusterResourceUsageExportConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandClusterResourceUsageExportConfigBigqueryDestination(c, f.BigqueryDestination); err != nil {
		return nil, fmt.Errorf("error expanding BigqueryDestination into bigqueryDestination: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bigqueryDestination"] = v
	}
	if v := f.EnableNetworkEgressMonitoring; !dcl.IsEmptyValueIndirect(v) {
		m["enableNetworkEgressMonitoring"] = v
	}
	if v, err := expandClusterResourceUsageExportConfigConsumptionMeteringConfig(c, f.ConsumptionMeteringConfig); err != nil {
		return nil, fmt.Errorf("error expanding ConsumptionMeteringConfig into consumptionMeteringConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["consumptionMeteringConfig"] = v
	}
	if v := f.EnableNetworkEgressMetering; !dcl.IsEmptyValueIndirect(v) {
		m["enableNetworkEgressMetering"] = v
	}

	return m, nil
}

// flattenClusterResourceUsageExportConfig flattens an instance of ClusterResourceUsageExportConfig from a JSON
// response object.
func flattenClusterResourceUsageExportConfig(c *Client, i interface{}) *ClusterResourceUsageExportConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterResourceUsageExportConfig{}
	r.BigqueryDestination = flattenClusterResourceUsageExportConfigBigqueryDestination(c, m["bigqueryDestination"])
	r.EnableNetworkEgressMonitoring = dcl.FlattenBool(m["enableNetworkEgressMonitoring"])
	r.ConsumptionMeteringConfig = flattenClusterResourceUsageExportConfigConsumptionMeteringConfig(c, m["consumptionMeteringConfig"])
	r.EnableNetworkEgressMetering = dcl.FlattenBool(m["enableNetworkEgressMetering"])

	return r
}

// expandClusterResourceUsageExportConfigBigqueryDestinationMap expands the contents of ClusterResourceUsageExportConfigBigqueryDestination into a JSON
// request object.
func expandClusterResourceUsageExportConfigBigqueryDestinationMap(c *Client, f map[string]ClusterResourceUsageExportConfigBigqueryDestination) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterResourceUsageExportConfigBigqueryDestination(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterResourceUsageExportConfigBigqueryDestinationSlice expands the contents of ClusterResourceUsageExportConfigBigqueryDestination into a JSON
// request object.
func expandClusterResourceUsageExportConfigBigqueryDestinationSlice(c *Client, f []ClusterResourceUsageExportConfigBigqueryDestination) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterResourceUsageExportConfigBigqueryDestination(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterResourceUsageExportConfigBigqueryDestinationMap flattens the contents of ClusterResourceUsageExportConfigBigqueryDestination from a JSON
// response object.
func flattenClusterResourceUsageExportConfigBigqueryDestinationMap(c *Client, i interface{}) map[string]ClusterResourceUsageExportConfigBigqueryDestination {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterResourceUsageExportConfigBigqueryDestination{}
	}

	if len(a) == 0 {
		return map[string]ClusterResourceUsageExportConfigBigqueryDestination{}
	}

	items := make(map[string]ClusterResourceUsageExportConfigBigqueryDestination)
	for k, item := range a {
		items[k] = *flattenClusterResourceUsageExportConfigBigqueryDestination(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterResourceUsageExportConfigBigqueryDestinationSlice flattens the contents of ClusterResourceUsageExportConfigBigqueryDestination from a JSON
// response object.
func flattenClusterResourceUsageExportConfigBigqueryDestinationSlice(c *Client, i interface{}) []ClusterResourceUsageExportConfigBigqueryDestination {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterResourceUsageExportConfigBigqueryDestination{}
	}

	if len(a) == 0 {
		return []ClusterResourceUsageExportConfigBigqueryDestination{}
	}

	items := make([]ClusterResourceUsageExportConfigBigqueryDestination, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterResourceUsageExportConfigBigqueryDestination(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterResourceUsageExportConfigBigqueryDestination expands an instance of ClusterResourceUsageExportConfigBigqueryDestination into a JSON
// request object.
func expandClusterResourceUsageExportConfigBigqueryDestination(c *Client, f *ClusterResourceUsageExportConfigBigqueryDestination) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.DatasetId; !dcl.IsEmptyValueIndirect(v) {
		m["datasetId"] = v
	}

	return m, nil
}

// flattenClusterResourceUsageExportConfigBigqueryDestination flattens an instance of ClusterResourceUsageExportConfigBigqueryDestination from a JSON
// response object.
func flattenClusterResourceUsageExportConfigBigqueryDestination(c *Client, i interface{}) *ClusterResourceUsageExportConfigBigqueryDestination {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterResourceUsageExportConfigBigqueryDestination{}
	r.DatasetId = dcl.FlattenString(m["datasetId"])

	return r
}

// expandClusterResourceUsageExportConfigConsumptionMeteringConfigMap expands the contents of ClusterResourceUsageExportConfigConsumptionMeteringConfig into a JSON
// request object.
func expandClusterResourceUsageExportConfigConsumptionMeteringConfigMap(c *Client, f map[string]ClusterResourceUsageExportConfigConsumptionMeteringConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterResourceUsageExportConfigConsumptionMeteringConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterResourceUsageExportConfigConsumptionMeteringConfigSlice expands the contents of ClusterResourceUsageExportConfigConsumptionMeteringConfig into a JSON
// request object.
func expandClusterResourceUsageExportConfigConsumptionMeteringConfigSlice(c *Client, f []ClusterResourceUsageExportConfigConsumptionMeteringConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterResourceUsageExportConfigConsumptionMeteringConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterResourceUsageExportConfigConsumptionMeteringConfigMap flattens the contents of ClusterResourceUsageExportConfigConsumptionMeteringConfig from a JSON
// response object.
func flattenClusterResourceUsageExportConfigConsumptionMeteringConfigMap(c *Client, i interface{}) map[string]ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterResourceUsageExportConfigConsumptionMeteringConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterResourceUsageExportConfigConsumptionMeteringConfig{}
	}

	items := make(map[string]ClusterResourceUsageExportConfigConsumptionMeteringConfig)
	for k, item := range a {
		items[k] = *flattenClusterResourceUsageExportConfigConsumptionMeteringConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterResourceUsageExportConfigConsumptionMeteringConfigSlice flattens the contents of ClusterResourceUsageExportConfigConsumptionMeteringConfig from a JSON
// response object.
func flattenClusterResourceUsageExportConfigConsumptionMeteringConfigSlice(c *Client, i interface{}) []ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterResourceUsageExportConfigConsumptionMeteringConfig{}
	}

	if len(a) == 0 {
		return []ClusterResourceUsageExportConfigConsumptionMeteringConfig{}
	}

	items := make([]ClusterResourceUsageExportConfigConsumptionMeteringConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterResourceUsageExportConfigConsumptionMeteringConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterResourceUsageExportConfigConsumptionMeteringConfig expands an instance of ClusterResourceUsageExportConfigConsumptionMeteringConfig into a JSON
// request object.
func expandClusterResourceUsageExportConfigConsumptionMeteringConfig(c *Client, f *ClusterResourceUsageExportConfigConsumptionMeteringConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterResourceUsageExportConfigConsumptionMeteringConfig flattens an instance of ClusterResourceUsageExportConfigConsumptionMeteringConfig from a JSON
// response object.
func flattenClusterResourceUsageExportConfigConsumptionMeteringConfig(c *Client, i interface{}) *ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterResourceUsageExportConfigConsumptionMeteringConfig{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterAuthenticatorGroupsConfigMap expands the contents of ClusterAuthenticatorGroupsConfig into a JSON
// request object.
func expandClusterAuthenticatorGroupsConfigMap(c *Client, f map[string]ClusterAuthenticatorGroupsConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAuthenticatorGroupsConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAuthenticatorGroupsConfigSlice expands the contents of ClusterAuthenticatorGroupsConfig into a JSON
// request object.
func expandClusterAuthenticatorGroupsConfigSlice(c *Client, f []ClusterAuthenticatorGroupsConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAuthenticatorGroupsConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAuthenticatorGroupsConfigMap flattens the contents of ClusterAuthenticatorGroupsConfig from a JSON
// response object.
func flattenClusterAuthenticatorGroupsConfigMap(c *Client, i interface{}) map[string]ClusterAuthenticatorGroupsConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAuthenticatorGroupsConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterAuthenticatorGroupsConfig{}
	}

	items := make(map[string]ClusterAuthenticatorGroupsConfig)
	for k, item := range a {
		items[k] = *flattenClusterAuthenticatorGroupsConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAuthenticatorGroupsConfigSlice flattens the contents of ClusterAuthenticatorGroupsConfig from a JSON
// response object.
func flattenClusterAuthenticatorGroupsConfigSlice(c *Client, i interface{}) []ClusterAuthenticatorGroupsConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAuthenticatorGroupsConfig{}
	}

	if len(a) == 0 {
		return []ClusterAuthenticatorGroupsConfig{}
	}

	items := make([]ClusterAuthenticatorGroupsConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAuthenticatorGroupsConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAuthenticatorGroupsConfig expands an instance of ClusterAuthenticatorGroupsConfig into a JSON
// request object.
func expandClusterAuthenticatorGroupsConfig(c *Client, f *ClusterAuthenticatorGroupsConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.SecurityGroup; !dcl.IsEmptyValueIndirect(v) {
		m["securityGroup"] = v
	}

	return m, nil
}

// flattenClusterAuthenticatorGroupsConfig flattens an instance of ClusterAuthenticatorGroupsConfig from a JSON
// response object.
func flattenClusterAuthenticatorGroupsConfig(c *Client, i interface{}) *ClusterAuthenticatorGroupsConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAuthenticatorGroupsConfig{}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.SecurityGroup = dcl.FlattenString(m["securityGroup"])

	return r
}

// expandClusterPrivateClusterConfigMap expands the contents of ClusterPrivateClusterConfig into a JSON
// request object.
func expandClusterPrivateClusterConfigMap(c *Client, f map[string]ClusterPrivateClusterConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterPrivateClusterConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterPrivateClusterConfigSlice expands the contents of ClusterPrivateClusterConfig into a JSON
// request object.
func expandClusterPrivateClusterConfigSlice(c *Client, f []ClusterPrivateClusterConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterPrivateClusterConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterPrivateClusterConfigMap flattens the contents of ClusterPrivateClusterConfig from a JSON
// response object.
func flattenClusterPrivateClusterConfigMap(c *Client, i interface{}) map[string]ClusterPrivateClusterConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterPrivateClusterConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterPrivateClusterConfig{}
	}

	items := make(map[string]ClusterPrivateClusterConfig)
	for k, item := range a {
		items[k] = *flattenClusterPrivateClusterConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterPrivateClusterConfigSlice flattens the contents of ClusterPrivateClusterConfig from a JSON
// response object.
func flattenClusterPrivateClusterConfigSlice(c *Client, i interface{}) []ClusterPrivateClusterConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterPrivateClusterConfig{}
	}

	if len(a) == 0 {
		return []ClusterPrivateClusterConfig{}
	}

	items := make([]ClusterPrivateClusterConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterPrivateClusterConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterPrivateClusterConfig expands an instance of ClusterPrivateClusterConfig into a JSON
// request object.
func expandClusterPrivateClusterConfig(c *Client, f *ClusterPrivateClusterConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.EnablePrivateNodes; !dcl.IsEmptyValueIndirect(v) {
		m["enablePrivateNodes"] = v
	}
	if v := f.EnablePrivateEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["enablePrivateEndpoint"] = v
	}
	if v := f.MasterIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["masterIpv4CidrBlock"] = v
	}
	if v := f.PrivateEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["privateEndpoint"] = v
	}
	if v := f.PublicEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["publicEndpoint"] = v
	}
	if v := f.PeeringName; !dcl.IsEmptyValueIndirect(v) {
		m["peeringName"] = v
	}
	if v, err := expandClusterPrivateClusterConfigMasterGlobalAccessConfig(c, f.MasterGlobalAccessConfig); err != nil {
		return nil, fmt.Errorf("error expanding MasterGlobalAccessConfig into masterGlobalAccessConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["masterGlobalAccessConfig"] = v
	}

	return m, nil
}

// flattenClusterPrivateClusterConfig flattens an instance of ClusterPrivateClusterConfig from a JSON
// response object.
func flattenClusterPrivateClusterConfig(c *Client, i interface{}) *ClusterPrivateClusterConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterPrivateClusterConfig{}
	r.EnablePrivateNodes = dcl.FlattenBool(m["enablePrivateNodes"])
	r.EnablePrivateEndpoint = dcl.FlattenBool(m["enablePrivateEndpoint"])
	r.MasterIPv4CidrBlock = dcl.FlattenString(m["masterIpv4CidrBlock"])
	r.PrivateEndpoint = dcl.FlattenString(m["privateEndpoint"])
	r.PublicEndpoint = dcl.FlattenString(m["publicEndpoint"])
	r.PeeringName = dcl.FlattenString(m["peeringName"])
	r.MasterGlobalAccessConfig = flattenClusterPrivateClusterConfigMasterGlobalAccessConfig(c, m["masterGlobalAccessConfig"])

	return r
}

// expandClusterPrivateClusterConfigMasterGlobalAccessConfigMap expands the contents of ClusterPrivateClusterConfigMasterGlobalAccessConfig into a JSON
// request object.
func expandClusterPrivateClusterConfigMasterGlobalAccessConfigMap(c *Client, f map[string]ClusterPrivateClusterConfigMasterGlobalAccessConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterPrivateClusterConfigMasterGlobalAccessConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterPrivateClusterConfigMasterGlobalAccessConfigSlice expands the contents of ClusterPrivateClusterConfigMasterGlobalAccessConfig into a JSON
// request object.
func expandClusterPrivateClusterConfigMasterGlobalAccessConfigSlice(c *Client, f []ClusterPrivateClusterConfigMasterGlobalAccessConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterPrivateClusterConfigMasterGlobalAccessConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterPrivateClusterConfigMasterGlobalAccessConfigMap flattens the contents of ClusterPrivateClusterConfigMasterGlobalAccessConfig from a JSON
// response object.
func flattenClusterPrivateClusterConfigMasterGlobalAccessConfigMap(c *Client, i interface{}) map[string]ClusterPrivateClusterConfigMasterGlobalAccessConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterPrivateClusterConfigMasterGlobalAccessConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterPrivateClusterConfigMasterGlobalAccessConfig{}
	}

	items := make(map[string]ClusterPrivateClusterConfigMasterGlobalAccessConfig)
	for k, item := range a {
		items[k] = *flattenClusterPrivateClusterConfigMasterGlobalAccessConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterPrivateClusterConfigMasterGlobalAccessConfigSlice flattens the contents of ClusterPrivateClusterConfigMasterGlobalAccessConfig from a JSON
// response object.
func flattenClusterPrivateClusterConfigMasterGlobalAccessConfigSlice(c *Client, i interface{}) []ClusterPrivateClusterConfigMasterGlobalAccessConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterPrivateClusterConfigMasterGlobalAccessConfig{}
	}

	if len(a) == 0 {
		return []ClusterPrivateClusterConfigMasterGlobalAccessConfig{}
	}

	items := make([]ClusterPrivateClusterConfigMasterGlobalAccessConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterPrivateClusterConfigMasterGlobalAccessConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterPrivateClusterConfigMasterGlobalAccessConfig expands an instance of ClusterPrivateClusterConfigMasterGlobalAccessConfig into a JSON
// request object.
func expandClusterPrivateClusterConfigMasterGlobalAccessConfig(c *Client, f *ClusterPrivateClusterConfigMasterGlobalAccessConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterPrivateClusterConfigMasterGlobalAccessConfig flattens an instance of ClusterPrivateClusterConfigMasterGlobalAccessConfig from a JSON
// response object.
func flattenClusterPrivateClusterConfigMasterGlobalAccessConfig(c *Client, i interface{}) *ClusterPrivateClusterConfigMasterGlobalAccessConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterPrivateClusterConfigMasterGlobalAccessConfig{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterDatabaseEncryptionMap expands the contents of ClusterDatabaseEncryption into a JSON
// request object.
func expandClusterDatabaseEncryptionMap(c *Client, f map[string]ClusterDatabaseEncryption) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterDatabaseEncryption(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterDatabaseEncryptionSlice expands the contents of ClusterDatabaseEncryption into a JSON
// request object.
func expandClusterDatabaseEncryptionSlice(c *Client, f []ClusterDatabaseEncryption) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterDatabaseEncryption(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterDatabaseEncryptionMap flattens the contents of ClusterDatabaseEncryption from a JSON
// response object.
func flattenClusterDatabaseEncryptionMap(c *Client, i interface{}) map[string]ClusterDatabaseEncryption {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterDatabaseEncryption{}
	}

	if len(a) == 0 {
		return map[string]ClusterDatabaseEncryption{}
	}

	items := make(map[string]ClusterDatabaseEncryption)
	for k, item := range a {
		items[k] = *flattenClusterDatabaseEncryption(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterDatabaseEncryptionSlice flattens the contents of ClusterDatabaseEncryption from a JSON
// response object.
func flattenClusterDatabaseEncryptionSlice(c *Client, i interface{}) []ClusterDatabaseEncryption {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterDatabaseEncryption{}
	}

	if len(a) == 0 {
		return []ClusterDatabaseEncryption{}
	}

	items := make([]ClusterDatabaseEncryption, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterDatabaseEncryption(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterDatabaseEncryption expands an instance of ClusterDatabaseEncryption into a JSON
// request object.
func expandClusterDatabaseEncryption(c *Client, f *ClusterDatabaseEncryption) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.KeyName; !dcl.IsEmptyValueIndirect(v) {
		m["keyName"] = v
	}

	return m, nil
}

// flattenClusterDatabaseEncryption flattens an instance of ClusterDatabaseEncryption from a JSON
// response object.
func flattenClusterDatabaseEncryption(c *Client, i interface{}) *ClusterDatabaseEncryption {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterDatabaseEncryption{}
	r.State = flattenClusterDatabaseEncryptionStateEnum(m["state"])
	r.KeyName = dcl.FlattenString(m["keyName"])

	return r
}

// expandClusterVerticalPodAutoscalingMap expands the contents of ClusterVerticalPodAutoscaling into a JSON
// request object.
func expandClusterVerticalPodAutoscalingMap(c *Client, f map[string]ClusterVerticalPodAutoscaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterVerticalPodAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterVerticalPodAutoscalingSlice expands the contents of ClusterVerticalPodAutoscaling into a JSON
// request object.
func expandClusterVerticalPodAutoscalingSlice(c *Client, f []ClusterVerticalPodAutoscaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterVerticalPodAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterVerticalPodAutoscalingMap flattens the contents of ClusterVerticalPodAutoscaling from a JSON
// response object.
func flattenClusterVerticalPodAutoscalingMap(c *Client, i interface{}) map[string]ClusterVerticalPodAutoscaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterVerticalPodAutoscaling{}
	}

	if len(a) == 0 {
		return map[string]ClusterVerticalPodAutoscaling{}
	}

	items := make(map[string]ClusterVerticalPodAutoscaling)
	for k, item := range a {
		items[k] = *flattenClusterVerticalPodAutoscaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterVerticalPodAutoscalingSlice flattens the contents of ClusterVerticalPodAutoscaling from a JSON
// response object.
func flattenClusterVerticalPodAutoscalingSlice(c *Client, i interface{}) []ClusterVerticalPodAutoscaling {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterVerticalPodAutoscaling{}
	}

	if len(a) == 0 {
		return []ClusterVerticalPodAutoscaling{}
	}

	items := make([]ClusterVerticalPodAutoscaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterVerticalPodAutoscaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterVerticalPodAutoscaling expands an instance of ClusterVerticalPodAutoscaling into a JSON
// request object.
func expandClusterVerticalPodAutoscaling(c *Client, f *ClusterVerticalPodAutoscaling) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterVerticalPodAutoscaling flattens an instance of ClusterVerticalPodAutoscaling from a JSON
// response object.
func flattenClusterVerticalPodAutoscaling(c *Client, i interface{}) *ClusterVerticalPodAutoscaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterVerticalPodAutoscaling{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterShieldedNodesMap expands the contents of ClusterShieldedNodes into a JSON
// request object.
func expandClusterShieldedNodesMap(c *Client, f map[string]ClusterShieldedNodes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterShieldedNodes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterShieldedNodesSlice expands the contents of ClusterShieldedNodes into a JSON
// request object.
func expandClusterShieldedNodesSlice(c *Client, f []ClusterShieldedNodes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterShieldedNodes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterShieldedNodesMap flattens the contents of ClusterShieldedNodes from a JSON
// response object.
func flattenClusterShieldedNodesMap(c *Client, i interface{}) map[string]ClusterShieldedNodes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterShieldedNodes{}
	}

	if len(a) == 0 {
		return map[string]ClusterShieldedNodes{}
	}

	items := make(map[string]ClusterShieldedNodes)
	for k, item := range a {
		items[k] = *flattenClusterShieldedNodes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterShieldedNodesSlice flattens the contents of ClusterShieldedNodes from a JSON
// response object.
func flattenClusterShieldedNodesSlice(c *Client, i interface{}) []ClusterShieldedNodes {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterShieldedNodes{}
	}

	if len(a) == 0 {
		return []ClusterShieldedNodes{}
	}

	items := make([]ClusterShieldedNodes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterShieldedNodes(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterShieldedNodes expands an instance of ClusterShieldedNodes into a JSON
// request object.
func expandClusterShieldedNodes(c *Client, f *ClusterShieldedNodes) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterShieldedNodes flattens an instance of ClusterShieldedNodes from a JSON
// response object.
func flattenClusterShieldedNodes(c *Client, i interface{}) *ClusterShieldedNodes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterShieldedNodes{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterConditionsMap expands the contents of ClusterConditions into a JSON
// request object.
func expandClusterConditionsMap(c *Client, f map[string]ClusterConditions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConditions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConditionsSlice expands the contents of ClusterConditions into a JSON
// request object.
func expandClusterConditionsSlice(c *Client, f []ClusterConditions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConditions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConditionsMap flattens the contents of ClusterConditions from a JSON
// response object.
func flattenClusterConditionsMap(c *Client, i interface{}) map[string]ClusterConditions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConditions{}
	}

	if len(a) == 0 {
		return map[string]ClusterConditions{}
	}

	items := make(map[string]ClusterConditions)
	for k, item := range a {
		items[k] = *flattenClusterConditions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConditionsSlice flattens the contents of ClusterConditions from a JSON
// response object.
func flattenClusterConditionsSlice(c *Client, i interface{}) []ClusterConditions {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConditions{}
	}

	if len(a) == 0 {
		return []ClusterConditions{}
	}

	items := make([]ClusterConditions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConditions(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConditions expands an instance of ClusterConditions into a JSON
// request object.
func expandClusterConditions(c *Client, f *ClusterConditions) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v := f.CanonicalCode; !dcl.IsEmptyValueIndirect(v) {
		m["canonicalCode"] = v
	}

	return m, nil
}

// flattenClusterConditions flattens an instance of ClusterConditions from a JSON
// response object.
func flattenClusterConditions(c *Client, i interface{}) *ClusterConditions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConditions{}
	r.Code = dcl.FlattenString(m["code"])
	r.Message = dcl.FlattenString(m["message"])
	r.CanonicalCode = flattenClusterConditionsCanonicalCodeEnum(m["canonicalCode"])

	return r
}

// expandClusterAutopilotMap expands the contents of ClusterAutopilot into a JSON
// request object.
func expandClusterAutopilotMap(c *Client, f map[string]ClusterAutopilot) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAutopilot(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAutopilotSlice expands the contents of ClusterAutopilot into a JSON
// request object.
func expandClusterAutopilotSlice(c *Client, f []ClusterAutopilot) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAutopilot(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAutopilotMap flattens the contents of ClusterAutopilot from a JSON
// response object.
func flattenClusterAutopilotMap(c *Client, i interface{}) map[string]ClusterAutopilot {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAutopilot{}
	}

	if len(a) == 0 {
		return map[string]ClusterAutopilot{}
	}

	items := make(map[string]ClusterAutopilot)
	for k, item := range a {
		items[k] = *flattenClusterAutopilot(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAutopilotSlice flattens the contents of ClusterAutopilot from a JSON
// response object.
func flattenClusterAutopilotSlice(c *Client, i interface{}) []ClusterAutopilot {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAutopilot{}
	}

	if len(a) == 0 {
		return []ClusterAutopilot{}
	}

	items := make([]ClusterAutopilot, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAutopilot(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAutopilot expands an instance of ClusterAutopilot into a JSON
// request object.
func expandClusterAutopilot(c *Client, f *ClusterAutopilot) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterAutopilot flattens an instance of ClusterAutopilot from a JSON
// response object.
func flattenClusterAutopilot(c *Client, i interface{}) *ClusterAutopilot {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAutopilot{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandClusterNodeConfigMap expands the contents of ClusterNodeConfig into a JSON
// request object.
func expandClusterNodeConfigMap(c *Client, f map[string]ClusterNodeConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeConfigSlice expands the contents of ClusterNodeConfig into a JSON
// request object.
func expandClusterNodeConfigSlice(c *Client, f []ClusterNodeConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeConfigMap flattens the contents of ClusterNodeConfig from a JSON
// response object.
func flattenClusterNodeConfigMap(c *Client, i interface{}) map[string]ClusterNodeConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeConfig{}
	}

	items := make(map[string]ClusterNodeConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodeConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodeConfigSlice flattens the contents of ClusterNodeConfig from a JSON
// response object.
func flattenClusterNodeConfigSlice(c *Client, i interface{}) []ClusterNodeConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfig{}
	}

	items := make([]ClusterNodeConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodeConfig expands an instance of ClusterNodeConfig into a JSON
// request object.
func expandClusterNodeConfig(c *Client, f *ClusterNodeConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}
	if v := f.DiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskSizeGb"] = v
	}
	if v := f.OAuthScopes; !dcl.IsEmptyValueIndirect(v) {
		m["oauthScopes"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v := f.Metadata; !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v := f.ImageType; !dcl.IsEmptyValueIndirect(v) {
		m["imageType"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.LocalSsdCount; !dcl.IsEmptyValueIndirect(v) {
		m["localSsdCount"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v := f.Preemptible; !dcl.IsEmptyValueIndirect(v) {
		m["preemptible"] = v
	}
	if v, err := expandClusterNodeConfigAcceleratorsSlice(c, f.Accelerators); err != nil {
		return nil, fmt.Errorf("error expanding Accelerators into accelerators: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["accelerators"] = v
	}
	if v := f.DiskType; !dcl.IsEmptyValueIndirect(v) {
		m["diskType"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}
	if v, err := expandClusterNodeConfigWorkloadMetadataConfig(c, f.WorkloadMetadataConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkloadMetadataConfig into workloadMetadataConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["workloadMetadataConfig"] = v
	}
	if v, err := expandClusterNodeConfigTaintsSlice(c, f.Taints); err != nil {
		return nil, fmt.Errorf("error expanding Taints into taints: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["taints"] = v
	}
	if v, err := expandClusterNodeConfigSandboxConfig(c, f.SandboxConfig); err != nil {
		return nil, fmt.Errorf("error expanding SandboxConfig into sandboxConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sandboxConfig"] = v
	}
	if v := f.NodeGroup; !dcl.IsEmptyValueIndirect(v) {
		m["nodeGroup"] = v
	}
	if v, err := expandClusterNodeConfigReservationAffinity(c, f.ReservationAffinity); err != nil {
		return nil, fmt.Errorf("error expanding ReservationAffinity into reservationAffinity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reservationAffinity"] = v
	}
	if v, err := expandClusterNodeConfigShieldedInstanceConfig(c, f.ShieldedInstanceConfig); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedInstanceConfig into shieldedInstanceConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["shieldedInstanceConfig"] = v
	}
	if v, err := expandClusterNodeConfigLinuxNodeConfig(c, f.LinuxNodeConfig); err != nil {
		return nil, fmt.Errorf("error expanding LinuxNodeConfig into linuxNodeConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["linuxNodeConfig"] = v
	}
	if v, err := expandClusterNodeConfigKubeletConfig(c, f.KubeletConfig); err != nil {
		return nil, fmt.Errorf("error expanding KubeletConfig into kubeletConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kubeletConfig"] = v
	}
	if v := f.BootDiskKmsKey; !dcl.IsEmptyValueIndirect(v) {
		m["bootDiskKmsKey"] = v
	}

	return m, nil
}

// flattenClusterNodeConfig flattens an instance of ClusterNodeConfig from a JSON
// response object.
func flattenClusterNodeConfig(c *Client, i interface{}) *ClusterNodeConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeConfig{}
	r.MachineType = dcl.FlattenString(m["machineType"])
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.OAuthScopes = dcl.FlattenStringSlice(m["oauthScopes"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.Metadata = dcl.FlattenKeyValuePairs(m["metadata"])
	r.ImageType = dcl.FlattenString(m["imageType"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.LocalSsdCount = dcl.FlattenInteger(m["localSsdCount"])
	r.Tags = dcl.FlattenStringSlice(m["tags"])
	r.Preemptible = dcl.FlattenBool(m["preemptible"])
	r.Accelerators = flattenClusterNodeConfigAcceleratorsSlice(c, m["accelerators"])
	r.DiskType = dcl.FlattenString(m["diskType"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])
	r.WorkloadMetadataConfig = flattenClusterNodeConfigWorkloadMetadataConfig(c, m["workloadMetadataConfig"])
	r.Taints = flattenClusterNodeConfigTaintsSlice(c, m["taints"])
	r.SandboxConfig = flattenClusterNodeConfigSandboxConfig(c, m["sandboxConfig"])
	r.NodeGroup = dcl.FlattenString(m["nodeGroup"])
	r.ReservationAffinity = flattenClusterNodeConfigReservationAffinity(c, m["reservationAffinity"])
	r.ShieldedInstanceConfig = flattenClusterNodeConfigShieldedInstanceConfig(c, m["shieldedInstanceConfig"])
	r.LinuxNodeConfig = flattenClusterNodeConfigLinuxNodeConfig(c, m["linuxNodeConfig"])
	r.KubeletConfig = flattenClusterNodeConfigKubeletConfig(c, m["kubeletConfig"])
	r.BootDiskKmsKey = dcl.FlattenString(m["bootDiskKmsKey"])

	return r
}

// expandClusterNodeConfigAcceleratorsMap expands the contents of ClusterNodeConfigAccelerators into a JSON
// request object.
func expandClusterNodeConfigAcceleratorsMap(c *Client, f map[string]ClusterNodeConfigAccelerators) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeConfigAccelerators(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeConfigAcceleratorsSlice expands the contents of ClusterNodeConfigAccelerators into a JSON
// request object.
func expandClusterNodeConfigAcceleratorsSlice(c *Client, f []ClusterNodeConfigAccelerators) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeConfigAccelerators(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeConfigAcceleratorsMap flattens the contents of ClusterNodeConfigAccelerators from a JSON
// response object.
func flattenClusterNodeConfigAcceleratorsMap(c *Client, i interface{}) map[string]ClusterNodeConfigAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeConfigAccelerators{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeConfigAccelerators{}
	}

	items := make(map[string]ClusterNodeConfigAccelerators)
	for k, item := range a {
		items[k] = *flattenClusterNodeConfigAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodeConfigAcceleratorsSlice flattens the contents of ClusterNodeConfigAccelerators from a JSON
// response object.
func flattenClusterNodeConfigAcceleratorsSlice(c *Client, i interface{}) []ClusterNodeConfigAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigAccelerators{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigAccelerators{}
	}

	items := make([]ClusterNodeConfigAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodeConfigAccelerators expands an instance of ClusterNodeConfigAccelerators into a JSON
// request object.
func expandClusterNodeConfigAccelerators(c *Client, f *ClusterNodeConfigAccelerators) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.AcceleratorCount; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorCount"] = v
	}
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorType"] = v
	}

	return m, nil
}

// flattenClusterNodeConfigAccelerators flattens an instance of ClusterNodeConfigAccelerators from a JSON
// response object.
func flattenClusterNodeConfigAccelerators(c *Client, i interface{}) *ClusterNodeConfigAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeConfigAccelerators{}
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])
	r.AcceleratorType = dcl.FlattenString(m["acceleratorType"])

	return r
}

// expandClusterNodeConfigWorkloadMetadataConfigMap expands the contents of ClusterNodeConfigWorkloadMetadataConfig into a JSON
// request object.
func expandClusterNodeConfigWorkloadMetadataConfigMap(c *Client, f map[string]ClusterNodeConfigWorkloadMetadataConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeConfigWorkloadMetadataConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeConfigWorkloadMetadataConfigSlice expands the contents of ClusterNodeConfigWorkloadMetadataConfig into a JSON
// request object.
func expandClusterNodeConfigWorkloadMetadataConfigSlice(c *Client, f []ClusterNodeConfigWorkloadMetadataConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeConfigWorkloadMetadataConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeConfigWorkloadMetadataConfigMap flattens the contents of ClusterNodeConfigWorkloadMetadataConfig from a JSON
// response object.
func flattenClusterNodeConfigWorkloadMetadataConfigMap(c *Client, i interface{}) map[string]ClusterNodeConfigWorkloadMetadataConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeConfigWorkloadMetadataConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeConfigWorkloadMetadataConfig{}
	}

	items := make(map[string]ClusterNodeConfigWorkloadMetadataConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodeConfigWorkloadMetadataConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodeConfigWorkloadMetadataConfigSlice flattens the contents of ClusterNodeConfigWorkloadMetadataConfig from a JSON
// response object.
func flattenClusterNodeConfigWorkloadMetadataConfigSlice(c *Client, i interface{}) []ClusterNodeConfigWorkloadMetadataConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigWorkloadMetadataConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigWorkloadMetadataConfig{}
	}

	items := make([]ClusterNodeConfigWorkloadMetadataConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigWorkloadMetadataConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodeConfigWorkloadMetadataConfig expands an instance of ClusterNodeConfigWorkloadMetadataConfig into a JSON
// request object.
func expandClusterNodeConfigWorkloadMetadataConfig(c *Client, f *ClusterNodeConfigWorkloadMetadataConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenClusterNodeConfigWorkloadMetadataConfig flattens an instance of ClusterNodeConfigWorkloadMetadataConfig from a JSON
// response object.
func flattenClusterNodeConfigWorkloadMetadataConfig(c *Client, i interface{}) *ClusterNodeConfigWorkloadMetadataConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeConfigWorkloadMetadataConfig{}
	r.Mode = flattenClusterNodeConfigWorkloadMetadataConfigModeEnum(m["mode"])

	return r
}

// expandClusterNodeConfigTaintsMap expands the contents of ClusterNodeConfigTaints into a JSON
// request object.
func expandClusterNodeConfigTaintsMap(c *Client, f map[string]ClusterNodeConfigTaints) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeConfigTaints(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeConfigTaintsSlice expands the contents of ClusterNodeConfigTaints into a JSON
// request object.
func expandClusterNodeConfigTaintsSlice(c *Client, f []ClusterNodeConfigTaints) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeConfigTaints(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeConfigTaintsMap flattens the contents of ClusterNodeConfigTaints from a JSON
// response object.
func flattenClusterNodeConfigTaintsMap(c *Client, i interface{}) map[string]ClusterNodeConfigTaints {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeConfigTaints{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeConfigTaints{}
	}

	items := make(map[string]ClusterNodeConfigTaints)
	for k, item := range a {
		items[k] = *flattenClusterNodeConfigTaints(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodeConfigTaintsSlice flattens the contents of ClusterNodeConfigTaints from a JSON
// response object.
func flattenClusterNodeConfigTaintsSlice(c *Client, i interface{}) []ClusterNodeConfigTaints {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigTaints{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigTaints{}
	}

	items := make([]ClusterNodeConfigTaints, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigTaints(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodeConfigTaints expands an instance of ClusterNodeConfigTaints into a JSON
// request object.
func expandClusterNodeConfigTaints(c *Client, f *ClusterNodeConfigTaints) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.Effect; !dcl.IsEmptyValueIndirect(v) {
		m["effect"] = v
	}

	return m, nil
}

// flattenClusterNodeConfigTaints flattens an instance of ClusterNodeConfigTaints from a JSON
// response object.
func flattenClusterNodeConfigTaints(c *Client, i interface{}) *ClusterNodeConfigTaints {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeConfigTaints{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])
	r.Effect = flattenClusterNodeConfigTaintsEffectEnum(m["effect"])

	return r
}

// expandClusterNodeConfigSandboxConfigMap expands the contents of ClusterNodeConfigSandboxConfig into a JSON
// request object.
func expandClusterNodeConfigSandboxConfigMap(c *Client, f map[string]ClusterNodeConfigSandboxConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeConfigSandboxConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeConfigSandboxConfigSlice expands the contents of ClusterNodeConfigSandboxConfig into a JSON
// request object.
func expandClusterNodeConfigSandboxConfigSlice(c *Client, f []ClusterNodeConfigSandboxConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeConfigSandboxConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeConfigSandboxConfigMap flattens the contents of ClusterNodeConfigSandboxConfig from a JSON
// response object.
func flattenClusterNodeConfigSandboxConfigMap(c *Client, i interface{}) map[string]ClusterNodeConfigSandboxConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeConfigSandboxConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeConfigSandboxConfig{}
	}

	items := make(map[string]ClusterNodeConfigSandboxConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodeConfigSandboxConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodeConfigSandboxConfigSlice flattens the contents of ClusterNodeConfigSandboxConfig from a JSON
// response object.
func flattenClusterNodeConfigSandboxConfigSlice(c *Client, i interface{}) []ClusterNodeConfigSandboxConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigSandboxConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigSandboxConfig{}
	}

	items := make([]ClusterNodeConfigSandboxConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigSandboxConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodeConfigSandboxConfig expands an instance of ClusterNodeConfigSandboxConfig into a JSON
// request object.
func expandClusterNodeConfigSandboxConfig(c *Client, f *ClusterNodeConfigSandboxConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenClusterNodeConfigSandboxConfig flattens an instance of ClusterNodeConfigSandboxConfig from a JSON
// response object.
func flattenClusterNodeConfigSandboxConfig(c *Client, i interface{}) *ClusterNodeConfigSandboxConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeConfigSandboxConfig{}
	r.Type = flattenClusterNodeConfigSandboxConfigTypeEnum(m["type"])

	return r
}

// expandClusterNodeConfigReservationAffinityMap expands the contents of ClusterNodeConfigReservationAffinity into a JSON
// request object.
func expandClusterNodeConfigReservationAffinityMap(c *Client, f map[string]ClusterNodeConfigReservationAffinity) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeConfigReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeConfigReservationAffinitySlice expands the contents of ClusterNodeConfigReservationAffinity into a JSON
// request object.
func expandClusterNodeConfigReservationAffinitySlice(c *Client, f []ClusterNodeConfigReservationAffinity) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeConfigReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeConfigReservationAffinityMap flattens the contents of ClusterNodeConfigReservationAffinity from a JSON
// response object.
func flattenClusterNodeConfigReservationAffinityMap(c *Client, i interface{}) map[string]ClusterNodeConfigReservationAffinity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeConfigReservationAffinity{}
	}

	items := make(map[string]ClusterNodeConfigReservationAffinity)
	for k, item := range a {
		items[k] = *flattenClusterNodeConfigReservationAffinity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodeConfigReservationAffinitySlice flattens the contents of ClusterNodeConfigReservationAffinity from a JSON
// response object.
func flattenClusterNodeConfigReservationAffinitySlice(c *Client, i interface{}) []ClusterNodeConfigReservationAffinity {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigReservationAffinity{}
	}

	items := make([]ClusterNodeConfigReservationAffinity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigReservationAffinity(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodeConfigReservationAffinity expands an instance of ClusterNodeConfigReservationAffinity into a JSON
// request object.
func expandClusterNodeConfigReservationAffinity(c *Client, f *ClusterNodeConfigReservationAffinity) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.ConsumeReservationType; !dcl.IsEmptyValueIndirect(v) {
		m["consumeReservationType"] = v
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Values; !dcl.IsEmptyValueIndirect(v) {
		m["values"] = v
	}

	return m, nil
}

// flattenClusterNodeConfigReservationAffinity flattens an instance of ClusterNodeConfigReservationAffinity from a JSON
// response object.
func flattenClusterNodeConfigReservationAffinity(c *Client, i interface{}) *ClusterNodeConfigReservationAffinity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeConfigReservationAffinity{}
	r.ConsumeReservationType = flattenClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(m["consumeReservationType"])
	r.Key = dcl.FlattenString(m["key"])
	r.Values = dcl.FlattenStringSlice(m["values"])

	return r
}

// expandClusterNodeConfigShieldedInstanceConfigMap expands the contents of ClusterNodeConfigShieldedInstanceConfig into a JSON
// request object.
func expandClusterNodeConfigShieldedInstanceConfigMap(c *Client, f map[string]ClusterNodeConfigShieldedInstanceConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeConfigShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeConfigShieldedInstanceConfigSlice expands the contents of ClusterNodeConfigShieldedInstanceConfig into a JSON
// request object.
func expandClusterNodeConfigShieldedInstanceConfigSlice(c *Client, f []ClusterNodeConfigShieldedInstanceConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeConfigShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeConfigShieldedInstanceConfigMap flattens the contents of ClusterNodeConfigShieldedInstanceConfig from a JSON
// response object.
func flattenClusterNodeConfigShieldedInstanceConfigMap(c *Client, i interface{}) map[string]ClusterNodeConfigShieldedInstanceConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeConfigShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeConfigShieldedInstanceConfig{}
	}

	items := make(map[string]ClusterNodeConfigShieldedInstanceConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodeConfigShieldedInstanceConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodeConfigShieldedInstanceConfigSlice flattens the contents of ClusterNodeConfigShieldedInstanceConfig from a JSON
// response object.
func flattenClusterNodeConfigShieldedInstanceConfigSlice(c *Client, i interface{}) []ClusterNodeConfigShieldedInstanceConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigShieldedInstanceConfig{}
	}

	items := make([]ClusterNodeConfigShieldedInstanceConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigShieldedInstanceConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodeConfigShieldedInstanceConfig expands an instance of ClusterNodeConfigShieldedInstanceConfig into a JSON
// request object.
func expandClusterNodeConfigShieldedInstanceConfig(c *Client, f *ClusterNodeConfigShieldedInstanceConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.EnableSecureBoot; !dcl.IsEmptyValueIndirect(v) {
		m["enableSecureBoot"] = v
	}
	if v := f.EnableIntegrityMonitoring; !dcl.IsEmptyValueIndirect(v) {
		m["enableIntegrityMonitoring"] = v
	}

	return m, nil
}

// flattenClusterNodeConfigShieldedInstanceConfig flattens an instance of ClusterNodeConfigShieldedInstanceConfig from a JSON
// response object.
func flattenClusterNodeConfigShieldedInstanceConfig(c *Client, i interface{}) *ClusterNodeConfigShieldedInstanceConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeConfigShieldedInstanceConfig{}
	r.EnableSecureBoot = dcl.FlattenBool(m["enableSecureBoot"])
	r.EnableIntegrityMonitoring = dcl.FlattenBool(m["enableIntegrityMonitoring"])

	return r
}

// expandClusterNodeConfigLinuxNodeConfigMap expands the contents of ClusterNodeConfigLinuxNodeConfig into a JSON
// request object.
func expandClusterNodeConfigLinuxNodeConfigMap(c *Client, f map[string]ClusterNodeConfigLinuxNodeConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeConfigLinuxNodeConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeConfigLinuxNodeConfigSlice expands the contents of ClusterNodeConfigLinuxNodeConfig into a JSON
// request object.
func expandClusterNodeConfigLinuxNodeConfigSlice(c *Client, f []ClusterNodeConfigLinuxNodeConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeConfigLinuxNodeConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeConfigLinuxNodeConfigMap flattens the contents of ClusterNodeConfigLinuxNodeConfig from a JSON
// response object.
func flattenClusterNodeConfigLinuxNodeConfigMap(c *Client, i interface{}) map[string]ClusterNodeConfigLinuxNodeConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeConfigLinuxNodeConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeConfigLinuxNodeConfig{}
	}

	items := make(map[string]ClusterNodeConfigLinuxNodeConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodeConfigLinuxNodeConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodeConfigLinuxNodeConfigSlice flattens the contents of ClusterNodeConfigLinuxNodeConfig from a JSON
// response object.
func flattenClusterNodeConfigLinuxNodeConfigSlice(c *Client, i interface{}) []ClusterNodeConfigLinuxNodeConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigLinuxNodeConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigLinuxNodeConfig{}
	}

	items := make([]ClusterNodeConfigLinuxNodeConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigLinuxNodeConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodeConfigLinuxNodeConfig expands an instance of ClusterNodeConfigLinuxNodeConfig into a JSON
// request object.
func expandClusterNodeConfigLinuxNodeConfig(c *Client, f *ClusterNodeConfigLinuxNodeConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Sysctls; !dcl.IsEmptyValueIndirect(v) {
		m["sysctls"] = v
	}

	return m, nil
}

// flattenClusterNodeConfigLinuxNodeConfig flattens an instance of ClusterNodeConfigLinuxNodeConfig from a JSON
// response object.
func flattenClusterNodeConfigLinuxNodeConfig(c *Client, i interface{}) *ClusterNodeConfigLinuxNodeConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeConfigLinuxNodeConfig{}
	r.Sysctls = dcl.FlattenKeyValuePairs(m["sysctls"])

	return r
}

// expandClusterNodeConfigKubeletConfigMap expands the contents of ClusterNodeConfigKubeletConfig into a JSON
// request object.
func expandClusterNodeConfigKubeletConfigMap(c *Client, f map[string]ClusterNodeConfigKubeletConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNodeConfigKubeletConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNodeConfigKubeletConfigSlice expands the contents of ClusterNodeConfigKubeletConfig into a JSON
// request object.
func expandClusterNodeConfigKubeletConfigSlice(c *Client, f []ClusterNodeConfigKubeletConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNodeConfigKubeletConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNodeConfigKubeletConfigMap flattens the contents of ClusterNodeConfigKubeletConfig from a JSON
// response object.
func flattenClusterNodeConfigKubeletConfigMap(c *Client, i interface{}) map[string]ClusterNodeConfigKubeletConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNodeConfigKubeletConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNodeConfigKubeletConfig{}
	}

	items := make(map[string]ClusterNodeConfigKubeletConfig)
	for k, item := range a {
		items[k] = *flattenClusterNodeConfigKubeletConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNodeConfigKubeletConfigSlice flattens the contents of ClusterNodeConfigKubeletConfig from a JSON
// response object.
func flattenClusterNodeConfigKubeletConfigSlice(c *Client, i interface{}) []ClusterNodeConfigKubeletConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigKubeletConfig{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigKubeletConfig{}
	}

	items := make([]ClusterNodeConfigKubeletConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigKubeletConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNodeConfigKubeletConfig expands an instance of ClusterNodeConfigKubeletConfig into a JSON
// request object.
func expandClusterNodeConfigKubeletConfig(c *Client, f *ClusterNodeConfigKubeletConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.CpuManagerPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["cpuManagerPolicy"] = v
	}
	if v := f.CpuCfsQuota; !dcl.IsEmptyValueIndirect(v) {
		m["cpuCfsQuota"] = v
	}
	if v := f.CpuCfsQuotaPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["cpuCfsQuotaPeriod"] = v
	}

	return m, nil
}

// flattenClusterNodeConfigKubeletConfig flattens an instance of ClusterNodeConfigKubeletConfig from a JSON
// response object.
func flattenClusterNodeConfigKubeletConfig(c *Client, i interface{}) *ClusterNodeConfigKubeletConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNodeConfigKubeletConfig{}
	r.CpuManagerPolicy = dcl.FlattenString(m["cpuManagerPolicy"])
	r.CpuCfsQuota = dcl.FlattenBool(m["cpuCfsQuota"])
	r.CpuCfsQuotaPeriod = dcl.FlattenString(m["cpuCfsQuotaPeriod"])

	return r
}

// expandClusterReleaseChannelMap expands the contents of ClusterReleaseChannel into a JSON
// request object.
func expandClusterReleaseChannelMap(c *Client, f map[string]ClusterReleaseChannel) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterReleaseChannel(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterReleaseChannelSlice expands the contents of ClusterReleaseChannel into a JSON
// request object.
func expandClusterReleaseChannelSlice(c *Client, f []ClusterReleaseChannel) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterReleaseChannel(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterReleaseChannelMap flattens the contents of ClusterReleaseChannel from a JSON
// response object.
func flattenClusterReleaseChannelMap(c *Client, i interface{}) map[string]ClusterReleaseChannel {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterReleaseChannel{}
	}

	if len(a) == 0 {
		return map[string]ClusterReleaseChannel{}
	}

	items := make(map[string]ClusterReleaseChannel)
	for k, item := range a {
		items[k] = *flattenClusterReleaseChannel(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterReleaseChannelSlice flattens the contents of ClusterReleaseChannel from a JSON
// response object.
func flattenClusterReleaseChannelSlice(c *Client, i interface{}) []ClusterReleaseChannel {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterReleaseChannel{}
	}

	if len(a) == 0 {
		return []ClusterReleaseChannel{}
	}

	items := make([]ClusterReleaseChannel, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterReleaseChannel(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterReleaseChannel expands an instance of ClusterReleaseChannel into a JSON
// request object.
func expandClusterReleaseChannel(c *Client, f *ClusterReleaseChannel) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Channel; !dcl.IsEmptyValueIndirect(v) {
		m["channel"] = v
	}

	return m, nil
}

// flattenClusterReleaseChannel flattens an instance of ClusterReleaseChannel from a JSON
// response object.
func flattenClusterReleaseChannel(c *Client, i interface{}) *ClusterReleaseChannel {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterReleaseChannel{}
	r.Channel = flattenClusterReleaseChannelChannelEnum(m["channel"])

	return r
}

// expandClusterWorkloadIdentityConfigMap expands the contents of ClusterWorkloadIdentityConfig into a JSON
// request object.
func expandClusterWorkloadIdentityConfigMap(c *Client, f map[string]ClusterWorkloadIdentityConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterWorkloadIdentityConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterWorkloadIdentityConfigSlice expands the contents of ClusterWorkloadIdentityConfig into a JSON
// request object.
func expandClusterWorkloadIdentityConfigSlice(c *Client, f []ClusterWorkloadIdentityConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterWorkloadIdentityConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterWorkloadIdentityConfigMap flattens the contents of ClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenClusterWorkloadIdentityConfigMap(c *Client, i interface{}) map[string]ClusterWorkloadIdentityConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterWorkloadIdentityConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterWorkloadIdentityConfig{}
	}

	items := make(map[string]ClusterWorkloadIdentityConfig)
	for k, item := range a {
		items[k] = *flattenClusterWorkloadIdentityConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterWorkloadIdentityConfigSlice flattens the contents of ClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenClusterWorkloadIdentityConfigSlice(c *Client, i interface{}) []ClusterWorkloadIdentityConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterWorkloadIdentityConfig{}
	}

	if len(a) == 0 {
		return []ClusterWorkloadIdentityConfig{}
	}

	items := make([]ClusterWorkloadIdentityConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterWorkloadIdentityConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterWorkloadIdentityConfig expands an instance of ClusterWorkloadIdentityConfig into a JSON
// request object.
func expandClusterWorkloadIdentityConfig(c *Client, f *ClusterWorkloadIdentityConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.WorkloadPool; !dcl.IsEmptyValueIndirect(v) {
		m["workloadPool"] = v
	}

	return m, nil
}

// flattenClusterWorkloadIdentityConfig flattens an instance of ClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenClusterWorkloadIdentityConfig(c *Client, i interface{}) *ClusterWorkloadIdentityConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterWorkloadIdentityConfig{}
	r.WorkloadPool = dcl.FlattenString(m["workloadPool"])

	return r
}

// expandClusterNotificationConfigMap expands the contents of ClusterNotificationConfig into a JSON
// request object.
func expandClusterNotificationConfigMap(c *Client, f map[string]ClusterNotificationConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNotificationConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNotificationConfigSlice expands the contents of ClusterNotificationConfig into a JSON
// request object.
func expandClusterNotificationConfigSlice(c *Client, f []ClusterNotificationConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNotificationConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNotificationConfigMap flattens the contents of ClusterNotificationConfig from a JSON
// response object.
func flattenClusterNotificationConfigMap(c *Client, i interface{}) map[string]ClusterNotificationConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNotificationConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterNotificationConfig{}
	}

	items := make(map[string]ClusterNotificationConfig)
	for k, item := range a {
		items[k] = *flattenClusterNotificationConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNotificationConfigSlice flattens the contents of ClusterNotificationConfig from a JSON
// response object.
func flattenClusterNotificationConfigSlice(c *Client, i interface{}) []ClusterNotificationConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNotificationConfig{}
	}

	if len(a) == 0 {
		return []ClusterNotificationConfig{}
	}

	items := make([]ClusterNotificationConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNotificationConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNotificationConfig expands an instance of ClusterNotificationConfig into a JSON
// request object.
func expandClusterNotificationConfig(c *Client, f *ClusterNotificationConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandClusterNotificationConfigPubsub(c, f.Pubsub); err != nil {
		return nil, fmt.Errorf("error expanding Pubsub into pubsub: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pubsub"] = v
	}

	return m, nil
}

// flattenClusterNotificationConfig flattens an instance of ClusterNotificationConfig from a JSON
// response object.
func flattenClusterNotificationConfig(c *Client, i interface{}) *ClusterNotificationConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNotificationConfig{}
	r.Pubsub = flattenClusterNotificationConfigPubsub(c, m["pubsub"])

	return r
}

// expandClusterNotificationConfigPubsubMap expands the contents of ClusterNotificationConfigPubsub into a JSON
// request object.
func expandClusterNotificationConfigPubsubMap(c *Client, f map[string]ClusterNotificationConfigPubsub) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNotificationConfigPubsub(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNotificationConfigPubsubSlice expands the contents of ClusterNotificationConfigPubsub into a JSON
// request object.
func expandClusterNotificationConfigPubsubSlice(c *Client, f []ClusterNotificationConfigPubsub) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNotificationConfigPubsub(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNotificationConfigPubsubMap flattens the contents of ClusterNotificationConfigPubsub from a JSON
// response object.
func flattenClusterNotificationConfigPubsubMap(c *Client, i interface{}) map[string]ClusterNotificationConfigPubsub {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNotificationConfigPubsub{}
	}

	if len(a) == 0 {
		return map[string]ClusterNotificationConfigPubsub{}
	}

	items := make(map[string]ClusterNotificationConfigPubsub)
	for k, item := range a {
		items[k] = *flattenClusterNotificationConfigPubsub(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNotificationConfigPubsubSlice flattens the contents of ClusterNotificationConfigPubsub from a JSON
// response object.
func flattenClusterNotificationConfigPubsubSlice(c *Client, i interface{}) []ClusterNotificationConfigPubsub {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNotificationConfigPubsub{}
	}

	if len(a) == 0 {
		return []ClusterNotificationConfigPubsub{}
	}

	items := make([]ClusterNotificationConfigPubsub, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNotificationConfigPubsub(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNotificationConfigPubsub expands an instance of ClusterNotificationConfigPubsub into a JSON
// request object.
func expandClusterNotificationConfigPubsub(c *Client, f *ClusterNotificationConfigPubsub) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.Topic; !dcl.IsEmptyValueIndirect(v) {
		m["topic"] = v
	}

	return m, nil
}

// flattenClusterNotificationConfigPubsub flattens an instance of ClusterNotificationConfigPubsub from a JSON
// response object.
func flattenClusterNotificationConfigPubsub(c *Client, i interface{}) *ClusterNotificationConfigPubsub {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNotificationConfigPubsub{}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.Topic = dcl.FlattenString(m["topic"])

	return r
}

// expandClusterConfidentialNodesMap expands the contents of ClusterConfidentialNodes into a JSON
// request object.
func expandClusterConfidentialNodesMap(c *Client, f map[string]ClusterConfidentialNodes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterConfidentialNodes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterConfidentialNodesSlice expands the contents of ClusterConfidentialNodes into a JSON
// request object.
func expandClusterConfidentialNodesSlice(c *Client, f []ClusterConfidentialNodes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterConfidentialNodes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterConfidentialNodesMap flattens the contents of ClusterConfidentialNodes from a JSON
// response object.
func flattenClusterConfidentialNodesMap(c *Client, i interface{}) map[string]ClusterConfidentialNodes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterConfidentialNodes{}
	}

	if len(a) == 0 {
		return map[string]ClusterConfidentialNodes{}
	}

	items := make(map[string]ClusterConfidentialNodes)
	for k, item := range a {
		items[k] = *flattenClusterConfidentialNodes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterConfidentialNodesSlice flattens the contents of ClusterConfidentialNodes from a JSON
// response object.
func flattenClusterConfidentialNodesSlice(c *Client, i interface{}) []ClusterConfidentialNodes {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConfidentialNodes{}
	}

	if len(a) == 0 {
		return []ClusterConfidentialNodes{}
	}

	items := make([]ClusterConfidentialNodes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConfidentialNodes(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterConfidentialNodes expands an instance of ClusterConfidentialNodes into a JSON
// request object.
func expandClusterConfidentialNodes(c *Client, f *ClusterConfidentialNodes) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenClusterConfidentialNodes flattens an instance of ClusterConfidentialNodes from a JSON
// response object.
func flattenClusterConfidentialNodes(c *Client, i interface{}) *ClusterConfidentialNodes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterConfidentialNodes{}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// flattenClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumSlice flattens the contents of ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum from a JSON
// response object.
func flattenClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumSlice(c *Client, i interface{}) []ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum{}
	}

	if len(a) == 0 {
		return []ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum{}
	}

	items := make([]ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum with the same value as that string.
func flattenClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(i interface{}) *ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumRef("")
	}

	return ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumRef(s)
}

// flattenClusterNodePoolsConfigWorkloadMetadataConfigModeEnumSlice flattens the contents of ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum from a JSON
// response object.
func flattenClusterNodePoolsConfigWorkloadMetadataConfigModeEnumSlice(c *Client, i interface{}) []ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum{}
	}

	items := make([]ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodePoolsConfigWorkloadMetadataConfigModeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum with the same value as that string.
func flattenClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(i interface{}) *ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodePoolsConfigWorkloadMetadataConfigModeEnumRef("")
	}

	return ClusterNodePoolsConfigWorkloadMetadataConfigModeEnumRef(s)
}

// flattenClusterNodePoolsConfigTaintsEffectEnumSlice flattens the contents of ClusterNodePoolsConfigTaintsEffectEnum from a JSON
// response object.
func flattenClusterNodePoolsConfigTaintsEffectEnumSlice(c *Client, i interface{}) []ClusterNodePoolsConfigTaintsEffectEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigTaintsEffectEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigTaintsEffectEnum{}
	}

	items := make([]ClusterNodePoolsConfigTaintsEffectEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigTaintsEffectEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodePoolsConfigTaintsEffectEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodePoolsConfigTaintsEffectEnum with the same value as that string.
func flattenClusterNodePoolsConfigTaintsEffectEnum(i interface{}) *ClusterNodePoolsConfigTaintsEffectEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodePoolsConfigTaintsEffectEnumRef("")
	}

	return ClusterNodePoolsConfigTaintsEffectEnumRef(s)
}

// flattenClusterNodePoolsConfigSandboxConfigTypeEnumSlice flattens the contents of ClusterNodePoolsConfigSandboxConfigTypeEnum from a JSON
// response object.
func flattenClusterNodePoolsConfigSandboxConfigTypeEnumSlice(c *Client, i interface{}) []ClusterNodePoolsConfigSandboxConfigTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigSandboxConfigTypeEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigSandboxConfigTypeEnum{}
	}

	items := make([]ClusterNodePoolsConfigSandboxConfigTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigSandboxConfigTypeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodePoolsConfigSandboxConfigTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodePoolsConfigSandboxConfigTypeEnum with the same value as that string.
func flattenClusterNodePoolsConfigSandboxConfigTypeEnum(i interface{}) *ClusterNodePoolsConfigSandboxConfigTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodePoolsConfigSandboxConfigTypeEnumRef("")
	}

	return ClusterNodePoolsConfigSandboxConfigTypeEnumRef(s)
}

// flattenClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumSlice flattens the contents of ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum from a JSON
// response object.
func flattenClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumSlice(c *Client, i interface{}) []ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	items := make([]ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum with the same value as that string.
func flattenClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(i interface{}) *ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumRef("")
	}

	return ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumRef(s)
}

// flattenClusterNodePoolsStatusEnumSlice flattens the contents of ClusterNodePoolsStatusEnum from a JSON
// response object.
func flattenClusterNodePoolsStatusEnumSlice(c *Client, i interface{}) []ClusterNodePoolsStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsStatusEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsStatusEnum{}
	}

	items := make([]ClusterNodePoolsStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsStatusEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodePoolsStatusEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodePoolsStatusEnum with the same value as that string.
func flattenClusterNodePoolsStatusEnum(i interface{}) *ClusterNodePoolsStatusEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodePoolsStatusEnumRef("")
	}

	return ClusterNodePoolsStatusEnumRef(s)
}

// flattenClusterNodePoolsConditionsCodeEnumSlice flattens the contents of ClusterNodePoolsConditionsCodeEnum from a JSON
// response object.
func flattenClusterNodePoolsConditionsCodeEnumSlice(c *Client, i interface{}) []ClusterNodePoolsConditionsCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConditionsCodeEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConditionsCodeEnum{}
	}

	items := make([]ClusterNodePoolsConditionsCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConditionsCodeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodePoolsConditionsCodeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodePoolsConditionsCodeEnum with the same value as that string.
func flattenClusterNodePoolsConditionsCodeEnum(i interface{}) *ClusterNodePoolsConditionsCodeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodePoolsConditionsCodeEnumRef("")
	}

	return ClusterNodePoolsConditionsCodeEnumRef(s)
}

// flattenClusterNodePoolsConditionsCanonicalCodeEnumSlice flattens the contents of ClusterNodePoolsConditionsCanonicalCodeEnum from a JSON
// response object.
func flattenClusterNodePoolsConditionsCanonicalCodeEnumSlice(c *Client, i interface{}) []ClusterNodePoolsConditionsCanonicalCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodePoolsConditionsCanonicalCodeEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodePoolsConditionsCanonicalCodeEnum{}
	}

	items := make([]ClusterNodePoolsConditionsCanonicalCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodePoolsConditionsCanonicalCodeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodePoolsConditionsCanonicalCodeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodePoolsConditionsCanonicalCodeEnum with the same value as that string.
func flattenClusterNodePoolsConditionsCanonicalCodeEnum(i interface{}) *ClusterNodePoolsConditionsCanonicalCodeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodePoolsConditionsCanonicalCodeEnumRef("")
	}

	return ClusterNodePoolsConditionsCanonicalCodeEnumRef(s)
}

// flattenClusterNetworkPolicyProviderEnumSlice flattens the contents of ClusterNetworkPolicyProviderEnum from a JSON
// response object.
func flattenClusterNetworkPolicyProviderEnumSlice(c *Client, i interface{}) []ClusterNetworkPolicyProviderEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNetworkPolicyProviderEnum{}
	}

	if len(a) == 0 {
		return []ClusterNetworkPolicyProviderEnum{}
	}

	items := make([]ClusterNetworkPolicyProviderEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNetworkPolicyProviderEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNetworkPolicyProviderEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNetworkPolicyProviderEnum with the same value as that string.
func flattenClusterNetworkPolicyProviderEnum(i interface{}) *ClusterNetworkPolicyProviderEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNetworkPolicyProviderEnumRef("")
	}

	return ClusterNetworkPolicyProviderEnumRef(s)
}

// flattenClusterNetworkConfigPrivateIPv6GoogleAccessEnumSlice flattens the contents of ClusterNetworkConfigPrivateIPv6GoogleAccessEnum from a JSON
// response object.
func flattenClusterNetworkConfigPrivateIPv6GoogleAccessEnumSlice(c *Client, i interface{}) []ClusterNetworkConfigPrivateIPv6GoogleAccessEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNetworkConfigPrivateIPv6GoogleAccessEnum{}
	}

	if len(a) == 0 {
		return []ClusterNetworkConfigPrivateIPv6GoogleAccessEnum{}
	}

	items := make([]ClusterNetworkConfigPrivateIPv6GoogleAccessEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNetworkConfigPrivateIPv6GoogleAccessEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNetworkConfigPrivateIPv6GoogleAccessEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNetworkConfigPrivateIPv6GoogleAccessEnum with the same value as that string.
func flattenClusterNetworkConfigPrivateIPv6GoogleAccessEnum(i interface{}) *ClusterNetworkConfigPrivateIPv6GoogleAccessEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNetworkConfigPrivateIPv6GoogleAccessEnumRef("")
	}

	return ClusterNetworkConfigPrivateIPv6GoogleAccessEnumRef(s)
}

// flattenClusterDatabaseEncryptionStateEnumSlice flattens the contents of ClusterDatabaseEncryptionStateEnum from a JSON
// response object.
func flattenClusterDatabaseEncryptionStateEnumSlice(c *Client, i interface{}) []ClusterDatabaseEncryptionStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterDatabaseEncryptionStateEnum{}
	}

	if len(a) == 0 {
		return []ClusterDatabaseEncryptionStateEnum{}
	}

	items := make([]ClusterDatabaseEncryptionStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterDatabaseEncryptionStateEnum(item.(interface{})))
	}

	return items
}

// flattenClusterDatabaseEncryptionStateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterDatabaseEncryptionStateEnum with the same value as that string.
func flattenClusterDatabaseEncryptionStateEnum(i interface{}) *ClusterDatabaseEncryptionStateEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterDatabaseEncryptionStateEnumRef("")
	}

	return ClusterDatabaseEncryptionStateEnumRef(s)
}

// flattenClusterConditionsCanonicalCodeEnumSlice flattens the contents of ClusterConditionsCanonicalCodeEnum from a JSON
// response object.
func flattenClusterConditionsCanonicalCodeEnumSlice(c *Client, i interface{}) []ClusterConditionsCanonicalCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterConditionsCanonicalCodeEnum{}
	}

	if len(a) == 0 {
		return []ClusterConditionsCanonicalCodeEnum{}
	}

	items := make([]ClusterConditionsCanonicalCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterConditionsCanonicalCodeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterConditionsCanonicalCodeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterConditionsCanonicalCodeEnum with the same value as that string.
func flattenClusterConditionsCanonicalCodeEnum(i interface{}) *ClusterConditionsCanonicalCodeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterConditionsCanonicalCodeEnumRef("")
	}

	return ClusterConditionsCanonicalCodeEnumRef(s)
}

// flattenClusterNodeConfigWorkloadMetadataConfigModeEnumSlice flattens the contents of ClusterNodeConfigWorkloadMetadataConfigModeEnum from a JSON
// response object.
func flattenClusterNodeConfigWorkloadMetadataConfigModeEnumSlice(c *Client, i interface{}) []ClusterNodeConfigWorkloadMetadataConfigModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigWorkloadMetadataConfigModeEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigWorkloadMetadataConfigModeEnum{}
	}

	items := make([]ClusterNodeConfigWorkloadMetadataConfigModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigWorkloadMetadataConfigModeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodeConfigWorkloadMetadataConfigModeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodeConfigWorkloadMetadataConfigModeEnum with the same value as that string.
func flattenClusterNodeConfigWorkloadMetadataConfigModeEnum(i interface{}) *ClusterNodeConfigWorkloadMetadataConfigModeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodeConfigWorkloadMetadataConfigModeEnumRef("")
	}

	return ClusterNodeConfigWorkloadMetadataConfigModeEnumRef(s)
}

// flattenClusterNodeConfigTaintsEffectEnumSlice flattens the contents of ClusterNodeConfigTaintsEffectEnum from a JSON
// response object.
func flattenClusterNodeConfigTaintsEffectEnumSlice(c *Client, i interface{}) []ClusterNodeConfigTaintsEffectEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigTaintsEffectEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigTaintsEffectEnum{}
	}

	items := make([]ClusterNodeConfigTaintsEffectEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigTaintsEffectEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodeConfigTaintsEffectEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodeConfigTaintsEffectEnum with the same value as that string.
func flattenClusterNodeConfigTaintsEffectEnum(i interface{}) *ClusterNodeConfigTaintsEffectEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodeConfigTaintsEffectEnumRef("")
	}

	return ClusterNodeConfigTaintsEffectEnumRef(s)
}

// flattenClusterNodeConfigSandboxConfigTypeEnumSlice flattens the contents of ClusterNodeConfigSandboxConfigTypeEnum from a JSON
// response object.
func flattenClusterNodeConfigSandboxConfigTypeEnumSlice(c *Client, i interface{}) []ClusterNodeConfigSandboxConfigTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigSandboxConfigTypeEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigSandboxConfigTypeEnum{}
	}

	items := make([]ClusterNodeConfigSandboxConfigTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigSandboxConfigTypeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodeConfigSandboxConfigTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodeConfigSandboxConfigTypeEnum with the same value as that string.
func flattenClusterNodeConfigSandboxConfigTypeEnum(i interface{}) *ClusterNodeConfigSandboxConfigTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodeConfigSandboxConfigTypeEnumRef("")
	}

	return ClusterNodeConfigSandboxConfigTypeEnumRef(s)
}

// flattenClusterNodeConfigReservationAffinityConsumeReservationTypeEnumSlice flattens the contents of ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum from a JSON
// response object.
func flattenClusterNodeConfigReservationAffinityConsumeReservationTypeEnumSlice(c *Client, i interface{}) []ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	if len(a) == 0 {
		return []ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	items := make([]ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(item.(interface{})))
	}

	return items
}

// flattenClusterNodeConfigReservationAffinityConsumeReservationTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum with the same value as that string.
func flattenClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(i interface{}) *ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterNodeConfigReservationAffinityConsumeReservationTypeEnumRef("")
	}

	return ClusterNodeConfigReservationAffinityConsumeReservationTypeEnumRef(s)
}

// flattenClusterReleaseChannelChannelEnumSlice flattens the contents of ClusterReleaseChannelChannelEnum from a JSON
// response object.
func flattenClusterReleaseChannelChannelEnumSlice(c *Client, i interface{}) []ClusterReleaseChannelChannelEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterReleaseChannelChannelEnum{}
	}

	if len(a) == 0 {
		return []ClusterReleaseChannelChannelEnum{}
	}

	items := make([]ClusterReleaseChannelChannelEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterReleaseChannelChannelEnum(item.(interface{})))
	}

	return items
}

// flattenClusterReleaseChannelChannelEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterReleaseChannelChannelEnum with the same value as that string.
func flattenClusterReleaseChannelChannelEnum(i interface{}) *ClusterReleaseChannelChannelEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterReleaseChannelChannelEnumRef("")
	}

	return ClusterReleaseChannelChannelEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Cluster) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalCluster(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}
