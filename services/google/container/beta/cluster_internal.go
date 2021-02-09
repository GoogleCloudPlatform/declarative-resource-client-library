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
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Cluster) validate() error {

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
func (r *ClusterNodePools) validate() error {
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
	return nil
}
func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) validate() error {
	return nil
}
func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) validate() error {
	return nil
}
func (r *ClusterNetworkConfig) validate() error {
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

func clusterGetURL(userBasePath string, r *Cluster) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, params), nil
}

func clusterListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters", "https://container.googleapis.com/v1beta1/", userBasePath, params), nil

}

func clusterCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters", "https://container.googleapis.com/v1beta1/", userBasePath, params), nil

}

func clusterDeleteURL(userBasePath string, r *Cluster) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, params), nil
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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET")

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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET"); err != nil {
		return err
	}
	_, err = c.GetCluster(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createClusterOperation struct{}

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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://container.googleapis.com/v1beta1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetCluster(ctx, r.urlNormalized()); err != nil {
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
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
	return dcl.Do(ctx, op.operate, c.Config.Retry)
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Cluster); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected Cluster, got %T", sh)
		} else {
			_ = r
		}
	}

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

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.InitialNodeCount) {
		rawDesired.InitialNodeCount = rawInitial.InitialNodeCount
	}
	rawDesired.MasterAuth = canonicalizeClusterMasterAuth(rawDesired.MasterAuth, rawInitial.MasterAuth, opts...)
	if dcl.IsZeroValue(rawDesired.LoggingService) {
		rawDesired.LoggingService = rawInitial.LoggingService
	}
	if dcl.IsZeroValue(rawDesired.MonitoringService) {
		rawDesired.MonitoringService = rawInitial.MonitoringService
	}
	if dcl.IsZeroValue(rawDesired.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.IsZeroValue(rawDesired.ClusterIPv4Cidr) {
		rawDesired.ClusterIPv4Cidr = rawInitial.ClusterIPv4Cidr
	}
	rawDesired.AddonsConfig = canonicalizeClusterAddonsConfig(rawDesired.AddonsConfig, rawInitial.AddonsConfig, opts...)
	if dcl.IsZeroValue(rawDesired.Subnetwork) {
		rawDesired.Subnetwork = rawInitial.Subnetwork
	}
	if dcl.IsZeroValue(rawDesired.NodePools) {
		rawDesired.NodePools = rawInitial.NodePools
	}
	if dcl.IsZeroValue(rawDesired.Locations) {
		rawDesired.Locations = rawInitial.Locations
	}
	if dcl.IsZeroValue(rawDesired.EnableKubernetesAlpha) {
		rawDesired.EnableKubernetesAlpha = rawInitial.EnableKubernetesAlpha
	}
	if dcl.IsZeroValue(rawDesired.ResourceLabels) {
		rawDesired.ResourceLabels = rawInitial.ResourceLabels
	}
	if dcl.IsZeroValue(rawDesired.LabelFingerprint) {
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
	if dcl.IsZeroValue(rawDesired.Endpoint) {
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
	if dcl.IsZeroValue(rawDesired.StatusMessage) {
		rawDesired.StatusMessage = rawInitial.StatusMessage
	}
	if dcl.IsZeroValue(rawDesired.NodeIPv4CidrSize) {
		rawDesired.NodeIPv4CidrSize = rawInitial.NodeIPv4CidrSize
	}
	if dcl.IsZeroValue(rawDesired.ServicesIPv4Cidr) {
		rawDesired.ServicesIPv4Cidr = rawInitial.ServicesIPv4Cidr
	}
	if dcl.IsZeroValue(rawDesired.ExpireTime) {
		rawDesired.ExpireTime = rawInitial.ExpireTime
	}
	if dcl.IsZeroValue(rawDesired.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.IsZeroValue(rawDesired.EnableTPU) {
		rawDesired.EnableTPU = rawInitial.EnableTPU
	}
	if dcl.IsZeroValue(rawDesired.TPUIPv4CidrBlock) {
		rawDesired.TPUIPv4CidrBlock = rawInitial.TPUIPv4CidrBlock
	}
	if dcl.IsZeroValue(rawDesired.Conditions) {
		rawDesired.Conditions = rawInitial.Conditions
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeClusterNewState(c *Client, rawNew, rawDesired *Cluster) (*Cluster, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.MonitoringService) && dcl.IsEmptyValueIndirect(rawDesired.MonitoringService) {
		rawNew.MonitoringService = rawDesired.MonitoringService
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClusterIPv4Cidr) && dcl.IsEmptyValueIndirect(rawDesired.ClusterIPv4Cidr) {
		rawNew.ClusterIPv4Cidr = rawDesired.ClusterIPv4Cidr
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AddonsConfig) && dcl.IsEmptyValueIndirect(rawDesired.AddonsConfig) {
		rawNew.AddonsConfig = rawDesired.AddonsConfig
	} else {
		rawNew.AddonsConfig = canonicalizeNewClusterAddonsConfig(c, rawDesired.AddonsConfig, rawNew.AddonsConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Subnetwork) && dcl.IsEmptyValueIndirect(rawDesired.Subnetwork) {
		rawNew.Subnetwork = rawDesired.Subnetwork
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.NodePools) && dcl.IsEmptyValueIndirect(rawDesired.NodePools) {
		rawNew.NodePools = rawDesired.NodePools
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Locations) && dcl.IsEmptyValueIndirect(rawDesired.Locations) {
		rawNew.Locations = rawDesired.Locations
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableKubernetesAlpha) && dcl.IsEmptyValueIndirect(rawDesired.EnableKubernetesAlpha) {
		rawNew.EnableKubernetesAlpha = rawDesired.EnableKubernetesAlpha
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourceLabels) && dcl.IsEmptyValueIndirect(rawDesired.ResourceLabels) {
		rawNew.ResourceLabels = rawDesired.ResourceLabels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LabelFingerprint) && dcl.IsEmptyValueIndirect(rawDesired.LabelFingerprint) {
		rawNew.LabelFingerprint = rawDesired.LabelFingerprint
	} else {
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.NodeIPv4CidrSize) && dcl.IsEmptyValueIndirect(rawDesired.NodeIPv4CidrSize) {
		rawNew.NodeIPv4CidrSize = rawDesired.NodeIPv4CidrSize
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServicesIPv4Cidr) && dcl.IsEmptyValueIndirect(rawDesired.ServicesIPv4Cidr) {
		rawNew.ServicesIPv4Cidr = rawDesired.ServicesIPv4Cidr
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpireTime) && dcl.IsEmptyValueIndirect(rawDesired.ExpireTime) {
		rawNew.ExpireTime = rawDesired.ExpireTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Location) && dcl.IsEmptyValueIndirect(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableTPU) && dcl.IsEmptyValueIndirect(rawDesired.EnableTPU) {
		rawNew.EnableTPU = rawDesired.EnableTPU
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TPUIPv4CidrBlock) && dcl.IsEmptyValueIndirect(rawDesired.TPUIPv4CidrBlock) {
		rawNew.TPUIPv4CidrBlock = rawDesired.TPUIPv4CidrBlock
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Conditions) && dcl.IsEmptyValueIndirect(rawDesired.Conditions) {
		rawNew.Conditions = rawDesired.Conditions
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeClusterMasterAuth(des, initial *ClusterMasterAuth, opts ...dcl.ApplyOption) *ClusterMasterAuth {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Username) {
		des.Username = initial.Username
	}
	if dcl.IsZeroValue(des.Password) {
		des.Password = initial.Password
	}
	des.ClientCertificateConfig = canonicalizeClusterMasterAuthClientCertificateConfig(des.ClientCertificateConfig, initial.ClientCertificateConfig, opts...)
	if dcl.IsZeroValue(des.ClusterCaCertificate) {
		des.ClusterCaCertificate = initial.ClusterCaCertificate
	}
	if dcl.IsZeroValue(des.ClientCertificate) {
		des.ClientCertificate = initial.ClientCertificate
	}
	if dcl.IsZeroValue(des.ClientKey) {
		des.ClientKey = initial.ClientKey
	}

	return des
}

func canonicalizeNewClusterMasterAuth(c *Client, des, nw *ClusterMasterAuth) *ClusterMasterAuth {
	if des == nil || nw == nil {
		return nw
	}

	nw.ClientCertificateConfig = canonicalizeNewClusterMasterAuthClientCertificateConfig(c, des.ClientCertificateConfig, nw.ClientCertificateConfig)

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

func canonicalizeClusterMasterAuthClientCertificateConfig(des, initial *ClusterMasterAuthClientCertificateConfig, opts ...dcl.ApplyOption) *ClusterMasterAuthClientCertificateConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.IssueClientCertificate) {
		des.IssueClientCertificate = initial.IssueClientCertificate
	}

	return des
}

func canonicalizeNewClusterMasterAuthClientCertificateConfig(c *Client, des, nw *ClusterMasterAuthClientCertificateConfig) *ClusterMasterAuthClientCertificateConfig {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterAddonsConfig(des, initial *ClusterAddonsConfig, opts ...dcl.ApplyOption) *ClusterAddonsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.HttpLoadBalancing = canonicalizeClusterAddonsConfigHttpLoadBalancing(des.HttpLoadBalancing, initial.HttpLoadBalancing, opts...)
	des.HorizontalPodAutoscaling = canonicalizeClusterAddonsConfigHorizontalPodAutoscaling(des.HorizontalPodAutoscaling, initial.HorizontalPodAutoscaling, opts...)
	des.KubernetesDashboard = canonicalizeClusterAddonsConfigKubernetesDashboard(des.KubernetesDashboard, initial.KubernetesDashboard, opts...)
	des.NetworkPolicyConfig = canonicalizeClusterAddonsConfigNetworkPolicyConfig(des.NetworkPolicyConfig, initial.NetworkPolicyConfig, opts...)
	des.CloudRunConfig = canonicalizeClusterAddonsConfigCloudRunConfig(des.CloudRunConfig, initial.CloudRunConfig, opts...)

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

func canonicalizeClusterAddonsConfigHttpLoadBalancing(des, initial *ClusterAddonsConfigHttpLoadBalancing, opts ...dcl.ApplyOption) *ClusterAddonsConfigHttpLoadBalancing {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigHttpLoadBalancing(c *Client, des, nw *ClusterAddonsConfigHttpLoadBalancing) *ClusterAddonsConfigHttpLoadBalancing {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterAddonsConfigHorizontalPodAutoscaling(des, initial *ClusterAddonsConfigHorizontalPodAutoscaling, opts ...dcl.ApplyOption) *ClusterAddonsConfigHorizontalPodAutoscaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigHorizontalPodAutoscaling(c *Client, des, nw *ClusterAddonsConfigHorizontalPodAutoscaling) *ClusterAddonsConfigHorizontalPodAutoscaling {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterAddonsConfigKubernetesDashboard(des, initial *ClusterAddonsConfigKubernetesDashboard, opts ...dcl.ApplyOption) *ClusterAddonsConfigKubernetesDashboard {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigKubernetesDashboard(c *Client, des, nw *ClusterAddonsConfigKubernetesDashboard) *ClusterAddonsConfigKubernetesDashboard {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterAddonsConfigNetworkPolicyConfig(des, initial *ClusterAddonsConfigNetworkPolicyConfig, opts ...dcl.ApplyOption) *ClusterAddonsConfigNetworkPolicyConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigNetworkPolicyConfig(c *Client, des, nw *ClusterAddonsConfigNetworkPolicyConfig) *ClusterAddonsConfigNetworkPolicyConfig {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterAddonsConfigCloudRunConfig(des, initial *ClusterAddonsConfigCloudRunConfig, opts ...dcl.ApplyOption) *ClusterAddonsConfigCloudRunConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}

	return des
}

func canonicalizeNewClusterAddonsConfigCloudRunConfig(c *Client, des, nw *ClusterAddonsConfigCloudRunConfig) *ClusterAddonsConfigCloudRunConfig {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterNodePools(des, initial *ClusterNodePools, opts ...dcl.ApplyOption) *ClusterNodePools {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}

	return des
}

func canonicalizeNewClusterNodePools(c *Client, des, nw *ClusterNodePools) *ClusterNodePools {
	if des == nil || nw == nil {
		return nw
	}

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

func canonicalizeClusterLegacyAbac(des, initial *ClusterLegacyAbac, opts ...dcl.ApplyOption) *ClusterLegacyAbac {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterLegacyAbac(c *Client, des, nw *ClusterLegacyAbac) *ClusterLegacyAbac {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterNetworkPolicy(des, initial *ClusterNetworkPolicy, opts ...dcl.ApplyOption) *ClusterNetworkPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Provider) {
		des.Provider = initial.Provider
	}
	if dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterNetworkPolicy(c *Client, des, nw *ClusterNetworkPolicy) *ClusterNetworkPolicy {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterIPAllocationPolicy(des, initial *ClusterIPAllocationPolicy, opts ...dcl.ApplyOption) *ClusterIPAllocationPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.UseIPAliases) {
		des.UseIPAliases = initial.UseIPAliases
	}
	if dcl.IsZeroValue(des.CreateSubnetwork) {
		des.CreateSubnetwork = initial.CreateSubnetwork
	}
	if dcl.IsZeroValue(des.SubnetworkName) {
		des.SubnetworkName = initial.SubnetworkName
	}
	if dcl.IsZeroValue(des.ClusterSecondaryRangeName) {
		des.ClusterSecondaryRangeName = initial.ClusterSecondaryRangeName
	}
	if dcl.IsZeroValue(des.ServicesSecondaryRangeName) {
		des.ServicesSecondaryRangeName = initial.ServicesSecondaryRangeName
	}
	if dcl.IsZeroValue(des.ClusterIPv4CidrBlock) {
		des.ClusterIPv4CidrBlock = initial.ClusterIPv4CidrBlock
	}
	if dcl.IsZeroValue(des.NodeIPv4CidrBlock) {
		des.NodeIPv4CidrBlock = initial.NodeIPv4CidrBlock
	}
	if dcl.IsZeroValue(des.ServicesIPv4CidrBlock) {
		des.ServicesIPv4CidrBlock = initial.ServicesIPv4CidrBlock
	}
	if dcl.IsZeroValue(des.TPUIPv4CidrBlock) {
		des.TPUIPv4CidrBlock = initial.TPUIPv4CidrBlock
	}

	return des
}

func canonicalizeNewClusterIPAllocationPolicy(c *Client, des, nw *ClusterIPAllocationPolicy) *ClusterIPAllocationPolicy {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterMasterAuthorizedNetworksConfig(des, initial *ClusterMasterAuthorizedNetworksConfig, opts ...dcl.ApplyOption) *ClusterMasterAuthorizedNetworksConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Enabled) {
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

func canonicalizeClusterMasterAuthorizedNetworksConfigCidrBlocks(des, initial *ClusterMasterAuthorizedNetworksConfigCidrBlocks, opts ...dcl.ApplyOption) *ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DisplayName) {
		des.DisplayName = initial.DisplayName
	}
	if dcl.IsZeroValue(des.CidrBlock) {
		des.CidrBlock = initial.CidrBlock
	}

	return des
}

func canonicalizeNewClusterMasterAuthorizedNetworksConfigCidrBlocks(c *Client, des, nw *ClusterMasterAuthorizedNetworksConfigCidrBlocks) *ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterBinaryAuthorization(des, initial *ClusterBinaryAuthorization, opts ...dcl.ApplyOption) *ClusterBinaryAuthorization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterBinaryAuthorization(c *Client, des, nw *ClusterBinaryAuthorization) *ClusterBinaryAuthorization {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterAutoscaling(des, initial *ClusterAutoscaling, opts ...dcl.ApplyOption) *ClusterAutoscaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.EnableNodeAutoprovisioning) {
		des.EnableNodeAutoprovisioning = initial.EnableNodeAutoprovisioning
	}
	if dcl.IsZeroValue(des.ResourceLimits) {
		des.ResourceLimits = initial.ResourceLimits
	}
	des.AutoprovisioningNodePoolDefaults = canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaults(des.AutoprovisioningNodePoolDefaults, initial.AutoprovisioningNodePoolDefaults, opts...)

	return des
}

func canonicalizeNewClusterAutoscaling(c *Client, des, nw *ClusterAutoscaling) *ClusterAutoscaling {
	if des == nil || nw == nil {
		return nw
	}

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

func canonicalizeClusterAutoscalingResourceLimits(des, initial *ClusterAutoscalingResourceLimits, opts ...dcl.ApplyOption) *ClusterAutoscalingResourceLimits {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ResourceType) {
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

func canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaults(des, initial *ClusterAutoscalingAutoprovisioningNodePoolDefaults, opts ...dcl.ApplyOption) *ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.OAuthScopes) {
		des.OAuthScopes = initial.OAuthScopes
	}
	if dcl.IsZeroValue(des.ServiceAccount) {
		des.ServiceAccount = initial.ServiceAccount
	}
	des.UpgradeSettings = canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(des.UpgradeSettings, initial.UpgradeSettings, opts...)
	des.Management = canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(des.Management, initial.Management, opts...)

	return des
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaults(c *Client, des, nw *ClusterAutoscalingAutoprovisioningNodePoolDefaults) *ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if des == nil || nw == nil {
		return nw
	}

	nw.UpgradeSettings = canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, des.UpgradeSettings, nw.UpgradeSettings)
	nw.Management = canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, des.Management, nw.Management)

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

func canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(des, initial *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings, opts ...dcl.ApplyOption) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
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

func canonicalizeClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(des, initial *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement, opts ...dcl.ApplyOption) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AutoUpgrade) {
		des.AutoUpgrade = initial.AutoUpgrade
	}
	if dcl.IsZeroValue(des.AutoRepair) {
		des.AutoRepair = initial.AutoRepair
	}

	return des
}

func canonicalizeNewClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c *Client, des, nw *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if des == nil || nw == nil {
		return nw
	}

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

func canonicalizeClusterNetworkConfig(des, initial *ClusterNetworkConfig, opts ...dcl.ApplyOption) *ClusterNetworkConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Network) {
		des.Network = initial.Network
	}
	if dcl.IsZeroValue(des.Subnetwork) {
		des.Subnetwork = initial.Subnetwork
	}
	if dcl.IsZeroValue(des.EnableIntraNodeVisibility) {
		des.EnableIntraNodeVisibility = initial.EnableIntraNodeVisibility
	}

	return des
}

func canonicalizeNewClusterNetworkConfig(c *Client, des, nw *ClusterNetworkConfig) *ClusterNetworkConfig {
	if des == nil || nw == nil {
		return nw
	}

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

func canonicalizeClusterMaintenancePolicy(des, initial *ClusterMaintenancePolicy, opts ...dcl.ApplyOption) *ClusterMaintenancePolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.Window = canonicalizeClusterMaintenancePolicyWindow(des.Window, initial.Window, opts...)
	if dcl.IsZeroValue(des.ResourceVersion) {
		des.ResourceVersion = initial.ResourceVersion
	}

	return des
}

func canonicalizeNewClusterMaintenancePolicy(c *Client, des, nw *ClusterMaintenancePolicy) *ClusterMaintenancePolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.Window = canonicalizeNewClusterMaintenancePolicyWindow(c, des.Window, nw.Window)

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

func canonicalizeClusterMaintenancePolicyWindow(des, initial *ClusterMaintenancePolicyWindow, opts ...dcl.ApplyOption) *ClusterMaintenancePolicyWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.DailyMaintenanceWindow = canonicalizeClusterMaintenancePolicyWindowDailyMaintenanceWindow(des.DailyMaintenanceWindow, initial.DailyMaintenanceWindow, opts...)
	des.RecurringWindow = canonicalizeClusterMaintenancePolicyWindowRecurringWindow(des.RecurringWindow, initial.RecurringWindow, opts...)

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

func canonicalizeClusterMaintenancePolicyWindowDailyMaintenanceWindow(des, initial *ClusterMaintenancePolicyWindowDailyMaintenanceWindow, opts ...dcl.ApplyOption) *ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.StartTime) {
		des.StartTime = initial.StartTime
	}
	if dcl.IsZeroValue(des.Duration) {
		des.Duration = initial.Duration
	}

	return des
}

func canonicalizeNewClusterMaintenancePolicyWindowDailyMaintenanceWindow(c *Client, des, nw *ClusterMaintenancePolicyWindowDailyMaintenanceWindow) *ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterMaintenancePolicyWindowRecurringWindow(des, initial *ClusterMaintenancePolicyWindowRecurringWindow, opts ...dcl.ApplyOption) *ClusterMaintenancePolicyWindowRecurringWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.Window = canonicalizeClusterMaintenancePolicyWindowRecurringWindowWindow(des.Window, initial.Window, opts...)
	if dcl.IsZeroValue(des.Recurrence) {
		des.Recurrence = initial.Recurrence
	}

	return des
}

func canonicalizeNewClusterMaintenancePolicyWindowRecurringWindow(c *Client, des, nw *ClusterMaintenancePolicyWindowRecurringWindow) *ClusterMaintenancePolicyWindowRecurringWindow {
	if des == nil || nw == nil {
		return nw
	}

	nw.Window = canonicalizeNewClusterMaintenancePolicyWindowRecurringWindowWindow(c, des.Window, nw.Window)

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

func canonicalizeClusterMaintenancePolicyWindowRecurringWindowWindow(des, initial *ClusterMaintenancePolicyWindowRecurringWindowWindow, opts ...dcl.ApplyOption) *ClusterMaintenancePolicyWindowRecurringWindowWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
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

func canonicalizeClusterDefaultMaxPodsConstraint(des, initial *ClusterDefaultMaxPodsConstraint, opts ...dcl.ApplyOption) *ClusterDefaultMaxPodsConstraint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MaxPodsPerNode) {
		des.MaxPodsPerNode = initial.MaxPodsPerNode
	}

	return des
}

func canonicalizeNewClusterDefaultMaxPodsConstraint(c *Client, des, nw *ClusterDefaultMaxPodsConstraint) *ClusterDefaultMaxPodsConstraint {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterResourceUsageExportConfig(des, initial *ClusterResourceUsageExportConfig, opts ...dcl.ApplyOption) *ClusterResourceUsageExportConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.BigqueryDestination = canonicalizeClusterResourceUsageExportConfigBigqueryDestination(des.BigqueryDestination, initial.BigqueryDestination, opts...)
	if dcl.IsZeroValue(des.EnableNetworkEgressMonitoring) {
		des.EnableNetworkEgressMonitoring = initial.EnableNetworkEgressMonitoring
	}
	des.ConsumptionMeteringConfig = canonicalizeClusterResourceUsageExportConfigConsumptionMeteringConfig(des.ConsumptionMeteringConfig, initial.ConsumptionMeteringConfig, opts...)

	return des
}

func canonicalizeNewClusterResourceUsageExportConfig(c *Client, des, nw *ClusterResourceUsageExportConfig) *ClusterResourceUsageExportConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.BigqueryDestination = canonicalizeNewClusterResourceUsageExportConfigBigqueryDestination(c, des.BigqueryDestination, nw.BigqueryDestination)
	nw.ConsumptionMeteringConfig = canonicalizeNewClusterResourceUsageExportConfigConsumptionMeteringConfig(c, des.ConsumptionMeteringConfig, nw.ConsumptionMeteringConfig)

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

func canonicalizeClusterResourceUsageExportConfigBigqueryDestination(des, initial *ClusterResourceUsageExportConfigBigqueryDestination, opts ...dcl.ApplyOption) *ClusterResourceUsageExportConfigBigqueryDestination {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DatasetId) {
		des.DatasetId = initial.DatasetId
	}

	return des
}

func canonicalizeNewClusterResourceUsageExportConfigBigqueryDestination(c *Client, des, nw *ClusterResourceUsageExportConfigBigqueryDestination) *ClusterResourceUsageExportConfigBigqueryDestination {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterResourceUsageExportConfigConsumptionMeteringConfig(des, initial *ClusterResourceUsageExportConfigConsumptionMeteringConfig, opts ...dcl.ApplyOption) *ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterResourceUsageExportConfigConsumptionMeteringConfig(c *Client, des, nw *ClusterResourceUsageExportConfigConsumptionMeteringConfig) *ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterAuthenticatorGroupsConfig(des, initial *ClusterAuthenticatorGroupsConfig, opts ...dcl.ApplyOption) *ClusterAuthenticatorGroupsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.IsZeroValue(des.SecurityGroup) {
		des.SecurityGroup = initial.SecurityGroup
	}

	return des
}

func canonicalizeNewClusterAuthenticatorGroupsConfig(c *Client, des, nw *ClusterAuthenticatorGroupsConfig) *ClusterAuthenticatorGroupsConfig {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterPrivateClusterConfig(des, initial *ClusterPrivateClusterConfig, opts ...dcl.ApplyOption) *ClusterPrivateClusterConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.EnablePrivateNodes) {
		des.EnablePrivateNodes = initial.EnablePrivateNodes
	}
	if dcl.IsZeroValue(des.EnablePrivateEndpoint) {
		des.EnablePrivateEndpoint = initial.EnablePrivateEndpoint
	}
	if dcl.IsZeroValue(des.MasterIPv4CidrBlock) {
		des.MasterIPv4CidrBlock = initial.MasterIPv4CidrBlock
	}
	if dcl.IsZeroValue(des.PrivateEndpoint) {
		des.PrivateEndpoint = initial.PrivateEndpoint
	}
	if dcl.IsZeroValue(des.PublicEndpoint) {
		des.PublicEndpoint = initial.PublicEndpoint
	}
	if dcl.IsZeroValue(des.PeeringName) {
		des.PeeringName = initial.PeeringName
	}

	return des
}

func canonicalizeNewClusterPrivateClusterConfig(c *Client, des, nw *ClusterPrivateClusterConfig) *ClusterPrivateClusterConfig {
	if des == nil || nw == nil {
		return nw
	}

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

func canonicalizeClusterDatabaseEncryption(des, initial *ClusterDatabaseEncryption, opts ...dcl.ApplyOption) *ClusterDatabaseEncryption {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.IsZeroValue(des.KeyName) {
		des.KeyName = initial.KeyName
	}

	return des
}

func canonicalizeNewClusterDatabaseEncryption(c *Client, des, nw *ClusterDatabaseEncryption) *ClusterDatabaseEncryption {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterVerticalPodAutoscaling(des, initial *ClusterVerticalPodAutoscaling, opts ...dcl.ApplyOption) *ClusterVerticalPodAutoscaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterVerticalPodAutoscaling(c *Client, des, nw *ClusterVerticalPodAutoscaling) *ClusterVerticalPodAutoscaling {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterShieldedNodes(des, initial *ClusterShieldedNodes, opts ...dcl.ApplyOption) *ClusterShieldedNodes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}

	return des
}

func canonicalizeNewClusterShieldedNodes(c *Client, des, nw *ClusterShieldedNodes) *ClusterShieldedNodes {
	if des == nil || nw == nil {
		return nw
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

func canonicalizeClusterConditions(des, initial *ClusterConditions, opts ...dcl.ApplyOption) *ClusterConditions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Cluster)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Code) {
		des.Code = initial.Code
	}
	if dcl.IsZeroValue(des.Message) {
		des.Message = initial.Message
	}

	return des
}

func canonicalizeNewClusterConditions(c *Client, des, nw *ClusterConditions) *ClusterConditions {
	if des == nil || nw == nil {
		return nw
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

type clusterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         clusterApiOperation
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
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.InitialNodeCount) && (dcl.IsZeroValue(actual.InitialNodeCount) || !reflect.DeepEqual(*desired.InitialNodeCount, *actual.InitialNodeCount)) {
		c.Config.Logger.Infof("Detected diff in InitialNodeCount.\nDESIRED: %v\nACTUAL: %v", desired.InitialNodeCount, actual.InitialNodeCount)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "InitialNodeCount",
		})
	}
	if compareClusterMasterAuth(c, desired.MasterAuth, actual.MasterAuth) {
		c.Config.Logger.Infof("Detected diff in MasterAuth.\nDESIRED: %v\nACTUAL: %v", desired.MasterAuth, actual.MasterAuth)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "MasterAuth",
		})
	}
	if !dcl.IsZeroValue(desired.LoggingService) && (dcl.IsZeroValue(actual.LoggingService) || !reflect.DeepEqual(*desired.LoggingService, *actual.LoggingService)) {
		c.Config.Logger.Infof("Detected diff in LoggingService.\nDESIRED: %v\nACTUAL: %v", desired.LoggingService, actual.LoggingService)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateMonitoringAndLoggingServiceOperation{},
			FieldName: "LoggingService",
		})

	}
	if !dcl.IsZeroValue(desired.MonitoringService) && (dcl.IsZeroValue(actual.MonitoringService) || !reflect.DeepEqual(*desired.MonitoringService, *actual.MonitoringService)) {
		c.Config.Logger.Infof("Detected diff in MonitoringService.\nDESIRED: %v\nACTUAL: %v", desired.MonitoringService, actual.MonitoringService)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateMonitoringAndLoggingServiceOperation{},
			FieldName: "MonitoringService",
		})

	}
	if !dcl.IsZeroValue(desired.Network) && (dcl.IsZeroValue(actual.Network) || !reflect.DeepEqual(*desired.Network, *actual.Network)) {
		c.Config.Logger.Infof("Detected diff in Network.\nDESIRED: %v\nACTUAL: %v", desired.Network, actual.Network)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "Network",
		})
	}
	if !dcl.IsZeroValue(desired.ClusterIPv4Cidr) && (dcl.IsZeroValue(actual.ClusterIPv4Cidr) || !reflect.DeepEqual(*desired.ClusterIPv4Cidr, *actual.ClusterIPv4Cidr)) {
		c.Config.Logger.Infof("Detected diff in ClusterIPv4Cidr.\nDESIRED: %v\nACTUAL: %v", desired.ClusterIPv4Cidr, actual.ClusterIPv4Cidr)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "ClusterIPv4Cidr",
		})
	}
	if compareClusterAddonsConfig(c, desired.AddonsConfig, actual.AddonsConfig) {
		c.Config.Logger.Infof("Detected diff in AddonsConfig.\nDESIRED: %v\nACTUAL: %v", desired.AddonsConfig, actual.AddonsConfig)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateAddonsConfigOperation{},
			FieldName: "AddonsConfig",
		})

	}
	if !dcl.IsZeroValue(desired.Subnetwork) && (dcl.IsZeroValue(actual.Subnetwork) || !reflect.DeepEqual(*desired.Subnetwork, *actual.Subnetwork)) {
		c.Config.Logger.Infof("Detected diff in Subnetwork.\nDESIRED: %v\nACTUAL: %v", desired.Subnetwork, actual.Subnetwork)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "Subnetwork",
		})
	}
	if !dcl.SliceEquals(desired.Locations, actual.Locations) {
		c.Config.Logger.Infof("Detected diff in Locations.\nDESIRED: %v\nACTUAL: %v", desired.Locations, actual.Locations)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateLocationsOperation{},
			FieldName: "Locations",
		})

	}
	if !dcl.IsZeroValue(desired.EnableKubernetesAlpha) && (dcl.IsZeroValue(actual.EnableKubernetesAlpha) || !reflect.DeepEqual(*desired.EnableKubernetesAlpha, *actual.EnableKubernetesAlpha)) {
		c.Config.Logger.Infof("Detected diff in EnableKubernetesAlpha.\nDESIRED: %v\nACTUAL: %v", desired.EnableKubernetesAlpha, actual.EnableKubernetesAlpha)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "EnableKubernetesAlpha",
		})
	}
	if !reflect.DeepEqual(desired.ResourceLabels, actual.ResourceLabels) {
		c.Config.Logger.Infof("Detected diff in ResourceLabels.\nDESIRED: %v\nACTUAL: %v", desired.ResourceLabels, actual.ResourceLabels)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "ResourceLabels",
		})
	}
	if compareClusterLegacyAbac(c, desired.LegacyAbac, actual.LegacyAbac) {
		c.Config.Logger.Infof("Detected diff in LegacyAbac.\nDESIRED: %v\nACTUAL: %v", desired.LegacyAbac, actual.LegacyAbac)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateLegacyAbacOperation{},
			FieldName: "LegacyAbac",
		})

	}
	if compareClusterNetworkPolicy(c, desired.NetworkPolicy, actual.NetworkPolicy) {
		c.Config.Logger.Infof("Detected diff in NetworkPolicy.\nDESIRED: %v\nACTUAL: %v", desired.NetworkPolicy, actual.NetworkPolicy)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "NetworkPolicy",
		})
	}
	if compareClusterIPAllocationPolicy(c, desired.IPAllocationPolicy, actual.IPAllocationPolicy) {
		c.Config.Logger.Infof("Detected diff in IPAllocationPolicy.\nDESIRED: %v\nACTUAL: %v", desired.IPAllocationPolicy, actual.IPAllocationPolicy)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "IPAllocationPolicy",
		})
	}
	if compareClusterMasterAuthorizedNetworksConfig(c, desired.MasterAuthorizedNetworksConfig, actual.MasterAuthorizedNetworksConfig) {
		c.Config.Logger.Infof("Detected diff in MasterAuthorizedNetworksConfig.\nDESIRED: %v\nACTUAL: %v", desired.MasterAuthorizedNetworksConfig, actual.MasterAuthorizedNetworksConfig)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateMasterAuthorizedNetworksConfigOperation{},
			FieldName: "MasterAuthorizedNetworksConfig",
		})

	}
	if compareClusterBinaryAuthorization(c, desired.BinaryAuthorization, actual.BinaryAuthorization) {
		c.Config.Logger.Infof("Detected diff in BinaryAuthorization.\nDESIRED: %v\nACTUAL: %v", desired.BinaryAuthorization, actual.BinaryAuthorization)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateBinaryAuthorizationOperation{},
			FieldName: "BinaryAuthorization",
		})

	}
	if compareClusterAutoscaling(c, desired.Autoscaling, actual.Autoscaling) {
		c.Config.Logger.Infof("Detected diff in Autoscaling.\nDESIRED: %v\nACTUAL: %v", desired.Autoscaling, actual.Autoscaling)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "Autoscaling",
		})
	}
	if compareClusterNetworkConfig(c, desired.NetworkConfig, actual.NetworkConfig) {
		c.Config.Logger.Infof("Detected diff in NetworkConfig.\nDESIRED: %v\nACTUAL: %v", desired.NetworkConfig, actual.NetworkConfig)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "NetworkConfig",
		})
	}
	if compareClusterMaintenancePolicy(c, desired.MaintenancePolicy, actual.MaintenancePolicy) {
		c.Config.Logger.Infof("Detected diff in MaintenancePolicy.\nDESIRED: %v\nACTUAL: %v", desired.MaintenancePolicy, actual.MaintenancePolicy)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterSetMaintenancePolicyOperation{},
			FieldName: "MaintenancePolicy",
		})

	}
	if compareClusterDefaultMaxPodsConstraint(c, desired.DefaultMaxPodsConstraint, actual.DefaultMaxPodsConstraint) {
		c.Config.Logger.Infof("Detected diff in DefaultMaxPodsConstraint.\nDESIRED: %v\nACTUAL: %v", desired.DefaultMaxPodsConstraint, actual.DefaultMaxPodsConstraint)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "DefaultMaxPodsConstraint",
		})
	}
	if compareClusterResourceUsageExportConfig(c, desired.ResourceUsageExportConfig, actual.ResourceUsageExportConfig) {
		c.Config.Logger.Infof("Detected diff in ResourceUsageExportConfig.\nDESIRED: %v\nACTUAL: %v", desired.ResourceUsageExportConfig, actual.ResourceUsageExportConfig)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "ResourceUsageExportConfig",
		})
	}
	if compareClusterAuthenticatorGroupsConfig(c, desired.AuthenticatorGroupsConfig, actual.AuthenticatorGroupsConfig) {
		c.Config.Logger.Infof("Detected diff in AuthenticatorGroupsConfig.\nDESIRED: %v\nACTUAL: %v", desired.AuthenticatorGroupsConfig, actual.AuthenticatorGroupsConfig)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "AuthenticatorGroupsConfig",
		})
	}
	if compareClusterPrivateClusterConfig(c, desired.PrivateClusterConfig, actual.PrivateClusterConfig) {
		c.Config.Logger.Infof("Detected diff in PrivateClusterConfig.\nDESIRED: %v\nACTUAL: %v", desired.PrivateClusterConfig, actual.PrivateClusterConfig)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "PrivateClusterConfig",
		})
	}
	if compareClusterDatabaseEncryption(c, desired.DatabaseEncryption, actual.DatabaseEncryption) {
		c.Config.Logger.Infof("Detected diff in DatabaseEncryption.\nDESIRED: %v\nACTUAL: %v", desired.DatabaseEncryption, actual.DatabaseEncryption)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateDatabaseEncryptionOperation{},
			FieldName: "DatabaseEncryption",
		})

	}
	if compareClusterVerticalPodAutoscaling(c, desired.VerticalPodAutoscaling, actual.VerticalPodAutoscaling) {
		c.Config.Logger.Infof("Detected diff in VerticalPodAutoscaling.\nDESIRED: %v\nACTUAL: %v", desired.VerticalPodAutoscaling, actual.VerticalPodAutoscaling)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateVerticalPodAutoscalingOperation{},
			FieldName: "VerticalPodAutoscaling",
		})

	}
	if compareClusterShieldedNodes(c, desired.ShieldedNodes, actual.ShieldedNodes) {
		c.Config.Logger.Infof("Detected diff in ShieldedNodes.\nDESIRED: %v\nACTUAL: %v", desired.ShieldedNodes, actual.ShieldedNodes)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateShieldedNodesOperation{},
			FieldName: "ShieldedNodes",
		})

	}
	if !dcl.IsZeroValue(desired.MasterVersion) && !dcl.MatchingSemver(desired.MasterVersion, actual.MasterVersion) {
		c.Config.Logger.Infof("Detected diff in MasterVersion.\nDESIRED: %v\nACTUAL: %v", desired.MasterVersion, actual.MasterVersion)

		diffs = append(diffs, clusterDiff{
			UpdateOp:  &updateClusterUpdateMasterVersionOperation{},
			FieldName: "MasterVersion",
		})

	}
	if !dcl.IsZeroValue(desired.Location) && (dcl.IsZeroValue(actual.Location) || !reflect.DeepEqual(*desired.Location, *actual.Location)) {
		c.Config.Logger.Infof("Detected diff in Location.\nDESIRED: %v\nACTUAL: %v", desired.Location, actual.Location)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "Location",
		})
	}
	if !dcl.IsZeroValue(desired.EnableTPU) && (dcl.IsZeroValue(actual.EnableTPU) || !reflect.DeepEqual(*desired.EnableTPU, *actual.EnableTPU)) {
		c.Config.Logger.Infof("Detected diff in EnableTPU.\nDESIRED: %v\nACTUAL: %v", desired.EnableTPU, actual.EnableTPU)
		diffs = append(diffs, clusterDiff{
			RequiresRecreate: true,
			FieldName:        "EnableTPU",
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
func compareClusterMasterAuthSlice(c *Client, desired, actual []ClusterMasterAuth) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterMasterAuth, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterMasterAuth(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterMasterAuth, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMasterAuth(c *Client, desired, actual *ClusterMasterAuth) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Username == nil && desired.Username != nil && !dcl.IsEmptyValueIndirect(desired.Username) {
		c.Config.Logger.Infof("desired Username %s - but actually nil", dcl.SprintResource(desired.Username))
		return true
	}
	if !reflect.DeepEqual(desired.Username, actual.Username) && !dcl.IsZeroValue(desired.Username) && !(dcl.IsEmptyValueIndirect(desired.Username) && dcl.IsZeroValue(actual.Username)) {
		c.Config.Logger.Infof("Diff in Username. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Username), dcl.SprintResource(actual.Username))
		return true
	}
	if actual.Password == nil && desired.Password != nil && !dcl.IsEmptyValueIndirect(desired.Password) {
		c.Config.Logger.Infof("desired Password %s - but actually nil", dcl.SprintResource(desired.Password))
		return true
	}
	if !reflect.DeepEqual(desired.Password, actual.Password) && !dcl.IsZeroValue(desired.Password) && !(dcl.IsEmptyValueIndirect(desired.Password) && dcl.IsZeroValue(actual.Password)) {
		c.Config.Logger.Infof("Diff in Password. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Password), dcl.SprintResource(actual.Password))
		return true
	}
	if actual.ClientCertificateConfig == nil && desired.ClientCertificateConfig != nil && !dcl.IsEmptyValueIndirect(desired.ClientCertificateConfig) {
		c.Config.Logger.Infof("desired ClientCertificateConfig %s - but actually nil", dcl.SprintResource(desired.ClientCertificateConfig))
		return true
	}
	if compareClusterMasterAuthClientCertificateConfig(c, desired.ClientCertificateConfig, actual.ClientCertificateConfig) && !dcl.IsZeroValue(desired.ClientCertificateConfig) {
		c.Config.Logger.Infof("Diff in ClientCertificateConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientCertificateConfig), dcl.SprintResource(actual.ClientCertificateConfig))
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
			c.Config.Logger.Infof("Diff in ClusterMasterAuthClientCertificateConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthClientCertificateConfig(c *Client, desired, actual *ClusterMasterAuthClientCertificateConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.IssueClientCertificate == nil && desired.IssueClientCertificate != nil && !dcl.IsEmptyValueIndirect(desired.IssueClientCertificate) {
		c.Config.Logger.Infof("desired IssueClientCertificate %s - but actually nil", dcl.SprintResource(desired.IssueClientCertificate))
		return true
	}
	if !reflect.DeepEqual(desired.IssueClientCertificate, actual.IssueClientCertificate) && !dcl.IsZeroValue(desired.IssueClientCertificate) && !(dcl.IsEmptyValueIndirect(desired.IssueClientCertificate) && dcl.IsZeroValue(actual.IssueClientCertificate)) {
		c.Config.Logger.Infof("Diff in IssueClientCertificate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IssueClientCertificate), dcl.SprintResource(actual.IssueClientCertificate))
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
			c.Config.Logger.Infof("Diff in ClusterAddonsConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfig(c *Client, desired, actual *ClusterAddonsConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HttpLoadBalancing == nil && desired.HttpLoadBalancing != nil && !dcl.IsEmptyValueIndirect(desired.HttpLoadBalancing) {
		c.Config.Logger.Infof("desired HttpLoadBalancing %s - but actually nil", dcl.SprintResource(desired.HttpLoadBalancing))
		return true
	}
	if compareClusterAddonsConfigHttpLoadBalancing(c, desired.HttpLoadBalancing, actual.HttpLoadBalancing) && !dcl.IsZeroValue(desired.HttpLoadBalancing) {
		c.Config.Logger.Infof("Diff in HttpLoadBalancing. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HttpLoadBalancing), dcl.SprintResource(actual.HttpLoadBalancing))
		return true
	}
	if actual.HorizontalPodAutoscaling == nil && desired.HorizontalPodAutoscaling != nil && !dcl.IsEmptyValueIndirect(desired.HorizontalPodAutoscaling) {
		c.Config.Logger.Infof("desired HorizontalPodAutoscaling %s - but actually nil", dcl.SprintResource(desired.HorizontalPodAutoscaling))
		return true
	}
	if compareClusterAddonsConfigHorizontalPodAutoscaling(c, desired.HorizontalPodAutoscaling, actual.HorizontalPodAutoscaling) && !dcl.IsZeroValue(desired.HorizontalPodAutoscaling) {
		c.Config.Logger.Infof("Diff in HorizontalPodAutoscaling. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HorizontalPodAutoscaling), dcl.SprintResource(actual.HorizontalPodAutoscaling))
		return true
	}
	if actual.KubernetesDashboard == nil && desired.KubernetesDashboard != nil && !dcl.IsEmptyValueIndirect(desired.KubernetesDashboard) {
		c.Config.Logger.Infof("desired KubernetesDashboard %s - but actually nil", dcl.SprintResource(desired.KubernetesDashboard))
		return true
	}
	if compareClusterAddonsConfigKubernetesDashboard(c, desired.KubernetesDashboard, actual.KubernetesDashboard) && !dcl.IsZeroValue(desired.KubernetesDashboard) {
		c.Config.Logger.Infof("Diff in KubernetesDashboard. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KubernetesDashboard), dcl.SprintResource(actual.KubernetesDashboard))
		return true
	}
	if actual.NetworkPolicyConfig == nil && desired.NetworkPolicyConfig != nil && !dcl.IsEmptyValueIndirect(desired.NetworkPolicyConfig) {
		c.Config.Logger.Infof("desired NetworkPolicyConfig %s - but actually nil", dcl.SprintResource(desired.NetworkPolicyConfig))
		return true
	}
	if compareClusterAddonsConfigNetworkPolicyConfig(c, desired.NetworkPolicyConfig, actual.NetworkPolicyConfig) && !dcl.IsZeroValue(desired.NetworkPolicyConfig) {
		c.Config.Logger.Infof("Diff in NetworkPolicyConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NetworkPolicyConfig), dcl.SprintResource(actual.NetworkPolicyConfig))
		return true
	}
	if actual.CloudRunConfig == nil && desired.CloudRunConfig != nil && !dcl.IsEmptyValueIndirect(desired.CloudRunConfig) {
		c.Config.Logger.Infof("desired CloudRunConfig %s - but actually nil", dcl.SprintResource(desired.CloudRunConfig))
		return true
	}
	if compareClusterAddonsConfigCloudRunConfig(c, desired.CloudRunConfig, actual.CloudRunConfig) && !dcl.IsZeroValue(desired.CloudRunConfig) {
		c.Config.Logger.Infof("Diff in CloudRunConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CloudRunConfig), dcl.SprintResource(actual.CloudRunConfig))
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
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigHttpLoadBalancing, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigHttpLoadBalancing(c *Client, desired, actual *ClusterAddonsConfigHttpLoadBalancing) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Disabled == nil && desired.Disabled != nil && !dcl.IsEmptyValueIndirect(desired.Disabled) {
		c.Config.Logger.Infof("desired Disabled %s - but actually nil", dcl.SprintResource(desired.Disabled))
		return true
	}
	if !reflect.DeepEqual(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) && !(dcl.IsEmptyValueIndirect(desired.Disabled) && dcl.IsZeroValue(actual.Disabled)) {
		c.Config.Logger.Infof("Diff in Disabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
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
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigHorizontalPodAutoscaling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigHorizontalPodAutoscaling(c *Client, desired, actual *ClusterAddonsConfigHorizontalPodAutoscaling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Disabled == nil && desired.Disabled != nil && !dcl.IsEmptyValueIndirect(desired.Disabled) {
		c.Config.Logger.Infof("desired Disabled %s - but actually nil", dcl.SprintResource(desired.Disabled))
		return true
	}
	if !reflect.DeepEqual(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) && !(dcl.IsEmptyValueIndirect(desired.Disabled) && dcl.IsZeroValue(actual.Disabled)) {
		c.Config.Logger.Infof("Diff in Disabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
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
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigKubernetesDashboard, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigKubernetesDashboard(c *Client, desired, actual *ClusterAddonsConfigKubernetesDashboard) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Disabled == nil && desired.Disabled != nil && !dcl.IsEmptyValueIndirect(desired.Disabled) {
		c.Config.Logger.Infof("desired Disabled %s - but actually nil", dcl.SprintResource(desired.Disabled))
		return true
	}
	if !reflect.DeepEqual(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) && !(dcl.IsEmptyValueIndirect(desired.Disabled) && dcl.IsZeroValue(actual.Disabled)) {
		c.Config.Logger.Infof("Diff in Disabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
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
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigNetworkPolicyConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigNetworkPolicyConfig(c *Client, desired, actual *ClusterAddonsConfigNetworkPolicyConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Disabled == nil && desired.Disabled != nil && !dcl.IsEmptyValueIndirect(desired.Disabled) {
		c.Config.Logger.Infof("desired Disabled %s - but actually nil", dcl.SprintResource(desired.Disabled))
		return true
	}
	if !reflect.DeepEqual(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) && !(dcl.IsEmptyValueIndirect(desired.Disabled) && dcl.IsZeroValue(actual.Disabled)) {
		c.Config.Logger.Infof("Diff in Disabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
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
			c.Config.Logger.Infof("Diff in ClusterAddonsConfigCloudRunConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAddonsConfigCloudRunConfig(c *Client, desired, actual *ClusterAddonsConfigCloudRunConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Disabled == nil && desired.Disabled != nil && !dcl.IsEmptyValueIndirect(desired.Disabled) {
		c.Config.Logger.Infof("desired Disabled %s - but actually nil", dcl.SprintResource(desired.Disabled))
		return true
	}
	if !reflect.DeepEqual(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) && !(dcl.IsEmptyValueIndirect(desired.Disabled) && dcl.IsZeroValue(actual.Disabled)) {
		c.Config.Logger.Infof("Diff in Disabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
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
			c.Config.Logger.Infof("Diff in ClusterNodePools, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNodePools(c *Client, desired, actual *ClusterNodePools) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
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
			c.Config.Logger.Infof("Diff in ClusterLegacyAbac, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterLegacyAbac(c *Client, desired, actual *ClusterLegacyAbac) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) && !(dcl.IsEmptyValueIndirect(desired.Enabled) && dcl.IsZeroValue(actual.Enabled)) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
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
			c.Config.Logger.Infof("Diff in ClusterNetworkPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNetworkPolicy(c *Client, desired, actual *ClusterNetworkPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Provider == nil && desired.Provider != nil && !dcl.IsEmptyValueIndirect(desired.Provider) {
		c.Config.Logger.Infof("desired Provider %s - but actually nil", dcl.SprintResource(desired.Provider))
		return true
	}
	if !reflect.DeepEqual(desired.Provider, actual.Provider) && !dcl.IsZeroValue(desired.Provider) && !(dcl.IsEmptyValueIndirect(desired.Provider) && dcl.IsZeroValue(actual.Provider)) {
		c.Config.Logger.Infof("Diff in Provider. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Provider), dcl.SprintResource(actual.Provider))
		return true
	}
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) && !(dcl.IsEmptyValueIndirect(desired.Enabled) && dcl.IsZeroValue(actual.Enabled)) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
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
			c.Config.Logger.Infof("Diff in ClusterIPAllocationPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterIPAllocationPolicy(c *Client, desired, actual *ClusterIPAllocationPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.UseIPAliases == nil && desired.UseIPAliases != nil && !dcl.IsEmptyValueIndirect(desired.UseIPAliases) {
		c.Config.Logger.Infof("desired UseIPAliases %s - but actually nil", dcl.SprintResource(desired.UseIPAliases))
		return true
	}
	if !reflect.DeepEqual(desired.UseIPAliases, actual.UseIPAliases) && !dcl.IsZeroValue(desired.UseIPAliases) && !(dcl.IsEmptyValueIndirect(desired.UseIPAliases) && dcl.IsZeroValue(actual.UseIPAliases)) {
		c.Config.Logger.Infof("Diff in UseIPAliases. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UseIPAliases), dcl.SprintResource(actual.UseIPAliases))
		return true
	}
	if actual.CreateSubnetwork == nil && desired.CreateSubnetwork != nil && !dcl.IsEmptyValueIndirect(desired.CreateSubnetwork) {
		c.Config.Logger.Infof("desired CreateSubnetwork %s - but actually nil", dcl.SprintResource(desired.CreateSubnetwork))
		return true
	}
	if !reflect.DeepEqual(desired.CreateSubnetwork, actual.CreateSubnetwork) && !dcl.IsZeroValue(desired.CreateSubnetwork) && !(dcl.IsEmptyValueIndirect(desired.CreateSubnetwork) && dcl.IsZeroValue(actual.CreateSubnetwork)) {
		c.Config.Logger.Infof("Diff in CreateSubnetwork. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CreateSubnetwork), dcl.SprintResource(actual.CreateSubnetwork))
		return true
	}
	if actual.SubnetworkName == nil && desired.SubnetworkName != nil && !dcl.IsEmptyValueIndirect(desired.SubnetworkName) {
		c.Config.Logger.Infof("desired SubnetworkName %s - but actually nil", dcl.SprintResource(desired.SubnetworkName))
		return true
	}
	if !reflect.DeepEqual(desired.SubnetworkName, actual.SubnetworkName) && !dcl.IsZeroValue(desired.SubnetworkName) && !(dcl.IsEmptyValueIndirect(desired.SubnetworkName) && dcl.IsZeroValue(actual.SubnetworkName)) {
		c.Config.Logger.Infof("Diff in SubnetworkName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SubnetworkName), dcl.SprintResource(actual.SubnetworkName))
		return true
	}
	if actual.ClusterSecondaryRangeName == nil && desired.ClusterSecondaryRangeName != nil && !dcl.IsEmptyValueIndirect(desired.ClusterSecondaryRangeName) {
		c.Config.Logger.Infof("desired ClusterSecondaryRangeName %s - but actually nil", dcl.SprintResource(desired.ClusterSecondaryRangeName))
		return true
	}
	if !reflect.DeepEqual(desired.ClusterSecondaryRangeName, actual.ClusterSecondaryRangeName) && !dcl.IsZeroValue(desired.ClusterSecondaryRangeName) && !(dcl.IsEmptyValueIndirect(desired.ClusterSecondaryRangeName) && dcl.IsZeroValue(actual.ClusterSecondaryRangeName)) {
		c.Config.Logger.Infof("Diff in ClusterSecondaryRangeName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterSecondaryRangeName), dcl.SprintResource(actual.ClusterSecondaryRangeName))
		return true
	}
	if actual.ServicesSecondaryRangeName == nil && desired.ServicesSecondaryRangeName != nil && !dcl.IsEmptyValueIndirect(desired.ServicesSecondaryRangeName) {
		c.Config.Logger.Infof("desired ServicesSecondaryRangeName %s - but actually nil", dcl.SprintResource(desired.ServicesSecondaryRangeName))
		return true
	}
	if !reflect.DeepEqual(desired.ServicesSecondaryRangeName, actual.ServicesSecondaryRangeName) && !dcl.IsZeroValue(desired.ServicesSecondaryRangeName) && !(dcl.IsEmptyValueIndirect(desired.ServicesSecondaryRangeName) && dcl.IsZeroValue(actual.ServicesSecondaryRangeName)) {
		c.Config.Logger.Infof("Diff in ServicesSecondaryRangeName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServicesSecondaryRangeName), dcl.SprintResource(actual.ServicesSecondaryRangeName))
		return true
	}
	if actual.ClusterIPv4CidrBlock == nil && desired.ClusterIPv4CidrBlock != nil && !dcl.IsEmptyValueIndirect(desired.ClusterIPv4CidrBlock) {
		c.Config.Logger.Infof("desired ClusterIPv4CidrBlock %s - but actually nil", dcl.SprintResource(desired.ClusterIPv4CidrBlock))
		return true
	}
	if !reflect.DeepEqual(desired.ClusterIPv4CidrBlock, actual.ClusterIPv4CidrBlock) && !dcl.IsZeroValue(desired.ClusterIPv4CidrBlock) && !(dcl.IsEmptyValueIndirect(desired.ClusterIPv4CidrBlock) && dcl.IsZeroValue(actual.ClusterIPv4CidrBlock)) {
		c.Config.Logger.Infof("Diff in ClusterIPv4CidrBlock. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClusterIPv4CidrBlock), dcl.SprintResource(actual.ClusterIPv4CidrBlock))
		return true
	}
	if actual.NodeIPv4CidrBlock == nil && desired.NodeIPv4CidrBlock != nil && !dcl.IsEmptyValueIndirect(desired.NodeIPv4CidrBlock) {
		c.Config.Logger.Infof("desired NodeIPv4CidrBlock %s - but actually nil", dcl.SprintResource(desired.NodeIPv4CidrBlock))
		return true
	}
	if !reflect.DeepEqual(desired.NodeIPv4CidrBlock, actual.NodeIPv4CidrBlock) && !dcl.IsZeroValue(desired.NodeIPv4CidrBlock) && !(dcl.IsEmptyValueIndirect(desired.NodeIPv4CidrBlock) && dcl.IsZeroValue(actual.NodeIPv4CidrBlock)) {
		c.Config.Logger.Infof("Diff in NodeIPv4CidrBlock. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NodeIPv4CidrBlock), dcl.SprintResource(actual.NodeIPv4CidrBlock))
		return true
	}
	if actual.ServicesIPv4CidrBlock == nil && desired.ServicesIPv4CidrBlock != nil && !dcl.IsEmptyValueIndirect(desired.ServicesIPv4CidrBlock) {
		c.Config.Logger.Infof("desired ServicesIPv4CidrBlock %s - but actually nil", dcl.SprintResource(desired.ServicesIPv4CidrBlock))
		return true
	}
	if !reflect.DeepEqual(desired.ServicesIPv4CidrBlock, actual.ServicesIPv4CidrBlock) && !dcl.IsZeroValue(desired.ServicesIPv4CidrBlock) && !(dcl.IsEmptyValueIndirect(desired.ServicesIPv4CidrBlock) && dcl.IsZeroValue(actual.ServicesIPv4CidrBlock)) {
		c.Config.Logger.Infof("Diff in ServicesIPv4CidrBlock. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServicesIPv4CidrBlock), dcl.SprintResource(actual.ServicesIPv4CidrBlock))
		return true
	}
	if actual.TPUIPv4CidrBlock == nil && desired.TPUIPv4CidrBlock != nil && !dcl.IsEmptyValueIndirect(desired.TPUIPv4CidrBlock) {
		c.Config.Logger.Infof("desired TPUIPv4CidrBlock %s - but actually nil", dcl.SprintResource(desired.TPUIPv4CidrBlock))
		return true
	}
	if !reflect.DeepEqual(desired.TPUIPv4CidrBlock, actual.TPUIPv4CidrBlock) && !dcl.IsZeroValue(desired.TPUIPv4CidrBlock) && !(dcl.IsEmptyValueIndirect(desired.TPUIPv4CidrBlock) && dcl.IsZeroValue(actual.TPUIPv4CidrBlock)) {
		c.Config.Logger.Infof("Diff in TPUIPv4CidrBlock. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TPUIPv4CidrBlock), dcl.SprintResource(actual.TPUIPv4CidrBlock))
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
			c.Config.Logger.Infof("Diff in ClusterMasterAuthorizedNetworksConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthorizedNetworksConfig(c *Client, desired, actual *ClusterMasterAuthorizedNetworksConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) && !(dcl.IsEmptyValueIndirect(desired.Enabled) && dcl.IsZeroValue(actual.Enabled)) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	if actual.CidrBlocks == nil && desired.CidrBlocks != nil && !dcl.IsEmptyValueIndirect(desired.CidrBlocks) {
		c.Config.Logger.Infof("desired CidrBlocks %s - but actually nil", dcl.SprintResource(desired.CidrBlocks))
		return true
	}
	if compareClusterMasterAuthorizedNetworksConfigCidrBlocksSlice(c, desired.CidrBlocks, actual.CidrBlocks) && !dcl.IsZeroValue(desired.CidrBlocks) {
		c.Config.Logger.Infof("Diff in CidrBlocks. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CidrBlocks), dcl.SprintResource(actual.CidrBlocks))
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
			c.Config.Logger.Infof("Diff in ClusterMasterAuthorizedNetworksConfigCidrBlocks, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMasterAuthorizedNetworksConfigCidrBlocks(c *Client, desired, actual *ClusterMasterAuthorizedNetworksConfigCidrBlocks) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DisplayName == nil && desired.DisplayName != nil && !dcl.IsEmptyValueIndirect(desired.DisplayName) {
		c.Config.Logger.Infof("desired DisplayName %s - but actually nil", dcl.SprintResource(desired.DisplayName))
		return true
	}
	if !reflect.DeepEqual(desired.DisplayName, actual.DisplayName) && !dcl.IsZeroValue(desired.DisplayName) && !(dcl.IsEmptyValueIndirect(desired.DisplayName) && dcl.IsZeroValue(actual.DisplayName)) {
		c.Config.Logger.Infof("Diff in DisplayName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DisplayName), dcl.SprintResource(actual.DisplayName))
		return true
	}
	if actual.CidrBlock == nil && desired.CidrBlock != nil && !dcl.IsEmptyValueIndirect(desired.CidrBlock) {
		c.Config.Logger.Infof("desired CidrBlock %s - but actually nil", dcl.SprintResource(desired.CidrBlock))
		return true
	}
	if !reflect.DeepEqual(desired.CidrBlock, actual.CidrBlock) && !dcl.IsZeroValue(desired.CidrBlock) && !(dcl.IsEmptyValueIndirect(desired.CidrBlock) && dcl.IsZeroValue(actual.CidrBlock)) {
		c.Config.Logger.Infof("Diff in CidrBlock. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CidrBlock), dcl.SprintResource(actual.CidrBlock))
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
			c.Config.Logger.Infof("Diff in ClusterBinaryAuthorization, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterBinaryAuthorization(c *Client, desired, actual *ClusterBinaryAuthorization) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) && !(dcl.IsEmptyValueIndirect(desired.Enabled) && dcl.IsZeroValue(actual.Enabled)) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
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
			c.Config.Logger.Infof("Diff in ClusterAutoscaling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscaling(c *Client, desired, actual *ClusterAutoscaling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EnableNodeAutoprovisioning == nil && desired.EnableNodeAutoprovisioning != nil && !dcl.IsEmptyValueIndirect(desired.EnableNodeAutoprovisioning) {
		c.Config.Logger.Infof("desired EnableNodeAutoprovisioning %s - but actually nil", dcl.SprintResource(desired.EnableNodeAutoprovisioning))
		return true
	}
	if !reflect.DeepEqual(desired.EnableNodeAutoprovisioning, actual.EnableNodeAutoprovisioning) && !dcl.IsZeroValue(desired.EnableNodeAutoprovisioning) && !(dcl.IsEmptyValueIndirect(desired.EnableNodeAutoprovisioning) && dcl.IsZeroValue(actual.EnableNodeAutoprovisioning)) {
		c.Config.Logger.Infof("Diff in EnableNodeAutoprovisioning. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableNodeAutoprovisioning), dcl.SprintResource(actual.EnableNodeAutoprovisioning))
		return true
	}
	if actual.ResourceLimits == nil && desired.ResourceLimits != nil && !dcl.IsEmptyValueIndirect(desired.ResourceLimits) {
		c.Config.Logger.Infof("desired ResourceLimits %s - but actually nil", dcl.SprintResource(desired.ResourceLimits))
		return true
	}
	if compareClusterAutoscalingResourceLimitsSlice(c, desired.ResourceLimits, actual.ResourceLimits) && !dcl.IsZeroValue(desired.ResourceLimits) {
		c.Config.Logger.Infof("Diff in ResourceLimits. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceLimits), dcl.SprintResource(actual.ResourceLimits))
		return true
	}
	if actual.AutoprovisioningNodePoolDefaults == nil && desired.AutoprovisioningNodePoolDefaults != nil && !dcl.IsEmptyValueIndirect(desired.AutoprovisioningNodePoolDefaults) {
		c.Config.Logger.Infof("desired AutoprovisioningNodePoolDefaults %s - but actually nil", dcl.SprintResource(desired.AutoprovisioningNodePoolDefaults))
		return true
	}
	if compareClusterAutoscalingAutoprovisioningNodePoolDefaults(c, desired.AutoprovisioningNodePoolDefaults, actual.AutoprovisioningNodePoolDefaults) && !dcl.IsZeroValue(desired.AutoprovisioningNodePoolDefaults) {
		c.Config.Logger.Infof("Diff in AutoprovisioningNodePoolDefaults. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoprovisioningNodePoolDefaults), dcl.SprintResource(actual.AutoprovisioningNodePoolDefaults))
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
			c.Config.Logger.Infof("Diff in ClusterAutoscalingResourceLimits, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingResourceLimits(c *Client, desired, actual *ClusterAutoscalingResourceLimits) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ResourceType == nil && desired.ResourceType != nil && !dcl.IsEmptyValueIndirect(desired.ResourceType) {
		c.Config.Logger.Infof("desired ResourceType %s - but actually nil", dcl.SprintResource(desired.ResourceType))
		return true
	}
	if !reflect.DeepEqual(desired.ResourceType, actual.ResourceType) && !dcl.IsZeroValue(desired.ResourceType) && !(dcl.IsEmptyValueIndirect(desired.ResourceType) && dcl.IsZeroValue(actual.ResourceType)) {
		c.Config.Logger.Infof("Diff in ResourceType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceType), dcl.SprintResource(actual.ResourceType))
		return true
	}
	if actual.Minimum == nil && desired.Minimum != nil && !dcl.IsEmptyValueIndirect(desired.Minimum) {
		c.Config.Logger.Infof("desired Minimum %s - but actually nil", dcl.SprintResource(desired.Minimum))
		return true
	}
	if !reflect.DeepEqual(desired.Minimum, actual.Minimum) && !dcl.IsZeroValue(desired.Minimum) && !(dcl.IsEmptyValueIndirect(desired.Minimum) && dcl.IsZeroValue(actual.Minimum)) {
		c.Config.Logger.Infof("Diff in Minimum. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Minimum), dcl.SprintResource(actual.Minimum))
		return true
	}
	if actual.Maximum == nil && desired.Maximum != nil && !dcl.IsEmptyValueIndirect(desired.Maximum) {
		c.Config.Logger.Infof("desired Maximum %s - but actually nil", dcl.SprintResource(desired.Maximum))
		return true
	}
	if !reflect.DeepEqual(desired.Maximum, actual.Maximum) && !dcl.IsZeroValue(desired.Maximum) && !(dcl.IsEmptyValueIndirect(desired.Maximum) && dcl.IsZeroValue(actual.Maximum)) {
		c.Config.Logger.Infof("Diff in Maximum. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Maximum), dcl.SprintResource(actual.Maximum))
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
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaults, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaults(c *Client, desired, actual *ClusterAutoscalingAutoprovisioningNodePoolDefaults) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.OAuthScopes == nil && desired.OAuthScopes != nil && !dcl.IsEmptyValueIndirect(desired.OAuthScopes) {
		c.Config.Logger.Infof("desired OAuthScopes %s - but actually nil", dcl.SprintResource(desired.OAuthScopes))
		return true
	}
	if !dcl.SliceEquals(desired.OAuthScopes, actual.OAuthScopes) && !dcl.IsZeroValue(desired.OAuthScopes) {
		c.Config.Logger.Infof("Diff in OAuthScopes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OAuthScopes), dcl.SprintResource(actual.OAuthScopes))
		return true
	}
	if actual.ServiceAccount == nil && desired.ServiceAccount != nil && !dcl.IsEmptyValueIndirect(desired.ServiceAccount) {
		c.Config.Logger.Infof("desired ServiceAccount %s - but actually nil", dcl.SprintResource(desired.ServiceAccount))
		return true
	}
	if !reflect.DeepEqual(desired.ServiceAccount, actual.ServiceAccount) && !dcl.IsZeroValue(desired.ServiceAccount) && !(dcl.IsEmptyValueIndirect(desired.ServiceAccount) && dcl.IsZeroValue(actual.ServiceAccount)) {
		c.Config.Logger.Infof("Diff in ServiceAccount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccount), dcl.SprintResource(actual.ServiceAccount))
		return true
	}
	if actual.UpgradeSettings == nil && desired.UpgradeSettings != nil && !dcl.IsEmptyValueIndirect(desired.UpgradeSettings) {
		c.Config.Logger.Infof("desired UpgradeSettings %s - but actually nil", dcl.SprintResource(desired.UpgradeSettings))
		return true
	}
	if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c, desired.UpgradeSettings, actual.UpgradeSettings) && !dcl.IsZeroValue(desired.UpgradeSettings) {
		c.Config.Logger.Infof("Diff in UpgradeSettings. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpgradeSettings), dcl.SprintResource(actual.UpgradeSettings))
		return true
	}
	if actual.Management == nil && desired.Management != nil && !dcl.IsEmptyValueIndirect(desired.Management) {
		c.Config.Logger.Infof("desired Management %s - but actually nil", dcl.SprintResource(desired.Management))
		return true
	}
	if compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c, desired.Management, actual.Management) && !dcl.IsZeroValue(desired.Management) {
		c.Config.Logger.Infof("Diff in Management. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Management), dcl.SprintResource(actual.Management))
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
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(c *Client, desired, actual *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MaxSurge == nil && desired.MaxSurge != nil && !dcl.IsEmptyValueIndirect(desired.MaxSurge) {
		c.Config.Logger.Infof("desired MaxSurge %s - but actually nil", dcl.SprintResource(desired.MaxSurge))
		return true
	}
	if !reflect.DeepEqual(desired.MaxSurge, actual.MaxSurge) && !dcl.IsZeroValue(desired.MaxSurge) && !(dcl.IsEmptyValueIndirect(desired.MaxSurge) && dcl.IsZeroValue(actual.MaxSurge)) {
		c.Config.Logger.Infof("Diff in MaxSurge. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxSurge), dcl.SprintResource(actual.MaxSurge))
		return true
	}
	if actual.MaxUnavailable == nil && desired.MaxUnavailable != nil && !dcl.IsEmptyValueIndirect(desired.MaxUnavailable) {
		c.Config.Logger.Infof("desired MaxUnavailable %s - but actually nil", dcl.SprintResource(desired.MaxUnavailable))
		return true
	}
	if !reflect.DeepEqual(desired.MaxUnavailable, actual.MaxUnavailable) && !dcl.IsZeroValue(desired.MaxUnavailable) && !(dcl.IsEmptyValueIndirect(desired.MaxUnavailable) && dcl.IsZeroValue(actual.MaxUnavailable)) {
		c.Config.Logger.Infof("Diff in MaxUnavailable. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxUnavailable), dcl.SprintResource(actual.MaxUnavailable))
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
			c.Config.Logger.Infof("Diff in ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(c *Client, desired, actual *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AutoUpgrade == nil && desired.AutoUpgrade != nil && !dcl.IsEmptyValueIndirect(desired.AutoUpgrade) {
		c.Config.Logger.Infof("desired AutoUpgrade %s - but actually nil", dcl.SprintResource(desired.AutoUpgrade))
		return true
	}
	if !reflect.DeepEqual(desired.AutoUpgrade, actual.AutoUpgrade) && !dcl.IsZeroValue(desired.AutoUpgrade) && !(dcl.IsEmptyValueIndirect(desired.AutoUpgrade) && dcl.IsZeroValue(actual.AutoUpgrade)) {
		c.Config.Logger.Infof("Diff in AutoUpgrade. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoUpgrade), dcl.SprintResource(actual.AutoUpgrade))
		return true
	}
	if actual.AutoRepair == nil && desired.AutoRepair != nil && !dcl.IsEmptyValueIndirect(desired.AutoRepair) {
		c.Config.Logger.Infof("desired AutoRepair %s - but actually nil", dcl.SprintResource(desired.AutoRepair))
		return true
	}
	if !reflect.DeepEqual(desired.AutoRepair, actual.AutoRepair) && !dcl.IsZeroValue(desired.AutoRepair) && !(dcl.IsEmptyValueIndirect(desired.AutoRepair) && dcl.IsZeroValue(actual.AutoRepair)) {
		c.Config.Logger.Infof("Diff in AutoRepair. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoRepair), dcl.SprintResource(actual.AutoRepair))
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
			c.Config.Logger.Infof("Diff in ClusterNetworkConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNetworkConfig(c *Client, desired, actual *ClusterNetworkConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EnableIntraNodeVisibility == nil && desired.EnableIntraNodeVisibility != nil && !dcl.IsEmptyValueIndirect(desired.EnableIntraNodeVisibility) {
		c.Config.Logger.Infof("desired EnableIntraNodeVisibility %s - but actually nil", dcl.SprintResource(desired.EnableIntraNodeVisibility))
		return true
	}
	if !reflect.DeepEqual(desired.EnableIntraNodeVisibility, actual.EnableIntraNodeVisibility) && !dcl.IsZeroValue(desired.EnableIntraNodeVisibility) && !(dcl.IsEmptyValueIndirect(desired.EnableIntraNodeVisibility) && dcl.IsZeroValue(actual.EnableIntraNodeVisibility)) {
		c.Config.Logger.Infof("Diff in EnableIntraNodeVisibility. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableIntraNodeVisibility), dcl.SprintResource(actual.EnableIntraNodeVisibility))
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
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicy(c *Client, desired, actual *ClusterMaintenancePolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Window == nil && desired.Window != nil && !dcl.IsEmptyValueIndirect(desired.Window) {
		c.Config.Logger.Infof("desired Window %s - but actually nil", dcl.SprintResource(desired.Window))
		return true
	}
	if compareClusterMaintenancePolicyWindow(c, desired.Window, actual.Window) && !dcl.IsZeroValue(desired.Window) {
		c.Config.Logger.Infof("Diff in Window. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Window), dcl.SprintResource(actual.Window))
		return true
	}
	if actual.ResourceVersion == nil && desired.ResourceVersion != nil && !dcl.IsEmptyValueIndirect(desired.ResourceVersion) {
		c.Config.Logger.Infof("desired ResourceVersion %s - but actually nil", dcl.SprintResource(desired.ResourceVersion))
		return true
	}
	if !reflect.DeepEqual(desired.ResourceVersion, actual.ResourceVersion) && !dcl.IsZeroValue(desired.ResourceVersion) && !(dcl.IsEmptyValueIndirect(desired.ResourceVersion) && dcl.IsZeroValue(actual.ResourceVersion)) {
		c.Config.Logger.Infof("Diff in ResourceVersion. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceVersion), dcl.SprintResource(actual.ResourceVersion))
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
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindow, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindow(c *Client, desired, actual *ClusterMaintenancePolicyWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DailyMaintenanceWindow == nil && desired.DailyMaintenanceWindow != nil && !dcl.IsEmptyValueIndirect(desired.DailyMaintenanceWindow) {
		c.Config.Logger.Infof("desired DailyMaintenanceWindow %s - but actually nil", dcl.SprintResource(desired.DailyMaintenanceWindow))
		return true
	}
	if compareClusterMaintenancePolicyWindowDailyMaintenanceWindow(c, desired.DailyMaintenanceWindow, actual.DailyMaintenanceWindow) && !dcl.IsZeroValue(desired.DailyMaintenanceWindow) {
		c.Config.Logger.Infof("Diff in DailyMaintenanceWindow. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DailyMaintenanceWindow), dcl.SprintResource(actual.DailyMaintenanceWindow))
		return true
	}
	if actual.RecurringWindow == nil && desired.RecurringWindow != nil && !dcl.IsEmptyValueIndirect(desired.RecurringWindow) {
		c.Config.Logger.Infof("desired RecurringWindow %s - but actually nil", dcl.SprintResource(desired.RecurringWindow))
		return true
	}
	if compareClusterMaintenancePolicyWindowRecurringWindow(c, desired.RecurringWindow, actual.RecurringWindow) && !dcl.IsZeroValue(desired.RecurringWindow) {
		c.Config.Logger.Infof("Diff in RecurringWindow. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RecurringWindow), dcl.SprintResource(actual.RecurringWindow))
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
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowDailyMaintenanceWindow, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowDailyMaintenanceWindow(c *Client, desired, actual *ClusterMaintenancePolicyWindowDailyMaintenanceWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.StartTime == nil && desired.StartTime != nil && !dcl.IsEmptyValueIndirect(desired.StartTime) {
		c.Config.Logger.Infof("desired StartTime %s - but actually nil", dcl.SprintResource(desired.StartTime))
		return true
	}
	if !reflect.DeepEqual(desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) && !(dcl.IsEmptyValueIndirect(desired.StartTime) && dcl.IsZeroValue(actual.StartTime)) {
		c.Config.Logger.Infof("Diff in StartTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
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
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowRecurringWindow, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowRecurringWindow(c *Client, desired, actual *ClusterMaintenancePolicyWindowRecurringWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Window == nil && desired.Window != nil && !dcl.IsEmptyValueIndirect(desired.Window) {
		c.Config.Logger.Infof("desired Window %s - but actually nil", dcl.SprintResource(desired.Window))
		return true
	}
	if compareClusterMaintenancePolicyWindowRecurringWindowWindow(c, desired.Window, actual.Window) && !dcl.IsZeroValue(desired.Window) {
		c.Config.Logger.Infof("Diff in Window. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Window), dcl.SprintResource(actual.Window))
		return true
	}
	if actual.Recurrence == nil && desired.Recurrence != nil && !dcl.IsEmptyValueIndirect(desired.Recurrence) {
		c.Config.Logger.Infof("desired Recurrence %s - but actually nil", dcl.SprintResource(desired.Recurrence))
		return true
	}
	if !reflect.DeepEqual(desired.Recurrence, actual.Recurrence) && !dcl.IsZeroValue(desired.Recurrence) && !(dcl.IsEmptyValueIndirect(desired.Recurrence) && dcl.IsZeroValue(actual.Recurrence)) {
		c.Config.Logger.Infof("Diff in Recurrence. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Recurrence), dcl.SprintResource(actual.Recurrence))
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
			c.Config.Logger.Infof("Diff in ClusterMaintenancePolicyWindowRecurringWindowWindow, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterMaintenancePolicyWindowRecurringWindowWindow(c *Client, desired, actual *ClusterMaintenancePolicyWindowRecurringWindowWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.StartTime == nil && desired.StartTime != nil && !dcl.IsEmptyValueIndirect(desired.StartTime) {
		c.Config.Logger.Infof("desired StartTime %s - but actually nil", dcl.SprintResource(desired.StartTime))
		return true
	}
	if !reflect.DeepEqual(desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) && !(dcl.IsEmptyValueIndirect(desired.StartTime) && dcl.IsZeroValue(actual.StartTime)) {
		c.Config.Logger.Infof("Diff in StartTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if actual.EndTime == nil && desired.EndTime != nil && !dcl.IsEmptyValueIndirect(desired.EndTime) {
		c.Config.Logger.Infof("desired EndTime %s - but actually nil", dcl.SprintResource(desired.EndTime))
		return true
	}
	if !reflect.DeepEqual(desired.EndTime, actual.EndTime) && !dcl.IsZeroValue(desired.EndTime) && !(dcl.IsEmptyValueIndirect(desired.EndTime) && dcl.IsZeroValue(actual.EndTime)) {
		c.Config.Logger.Infof("Diff in EndTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EndTime), dcl.SprintResource(actual.EndTime))
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
			c.Config.Logger.Infof("Diff in ClusterDefaultMaxPodsConstraint, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterDefaultMaxPodsConstraint(c *Client, desired, actual *ClusterDefaultMaxPodsConstraint) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MaxPodsPerNode == nil && desired.MaxPodsPerNode != nil && !dcl.IsEmptyValueIndirect(desired.MaxPodsPerNode) {
		c.Config.Logger.Infof("desired MaxPodsPerNode %s - but actually nil", dcl.SprintResource(desired.MaxPodsPerNode))
		return true
	}
	if !reflect.DeepEqual(desired.MaxPodsPerNode, actual.MaxPodsPerNode) && !dcl.IsZeroValue(desired.MaxPodsPerNode) && !(dcl.IsEmptyValueIndirect(desired.MaxPodsPerNode) && dcl.IsZeroValue(actual.MaxPodsPerNode)) {
		c.Config.Logger.Infof("Diff in MaxPodsPerNode. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxPodsPerNode), dcl.SprintResource(actual.MaxPodsPerNode))
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
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterResourceUsageExportConfig(c *Client, desired, actual *ClusterResourceUsageExportConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.BigqueryDestination == nil && desired.BigqueryDestination != nil && !dcl.IsEmptyValueIndirect(desired.BigqueryDestination) {
		c.Config.Logger.Infof("desired BigqueryDestination %s - but actually nil", dcl.SprintResource(desired.BigqueryDestination))
		return true
	}
	if compareClusterResourceUsageExportConfigBigqueryDestination(c, desired.BigqueryDestination, actual.BigqueryDestination) && !dcl.IsZeroValue(desired.BigqueryDestination) {
		c.Config.Logger.Infof("Diff in BigqueryDestination. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BigqueryDestination), dcl.SprintResource(actual.BigqueryDestination))
		return true
	}
	if actual.EnableNetworkEgressMonitoring == nil && desired.EnableNetworkEgressMonitoring != nil && !dcl.IsEmptyValueIndirect(desired.EnableNetworkEgressMonitoring) {
		c.Config.Logger.Infof("desired EnableNetworkEgressMonitoring %s - but actually nil", dcl.SprintResource(desired.EnableNetworkEgressMonitoring))
		return true
	}
	if !reflect.DeepEqual(desired.EnableNetworkEgressMonitoring, actual.EnableNetworkEgressMonitoring) && !dcl.IsZeroValue(desired.EnableNetworkEgressMonitoring) && !(dcl.IsEmptyValueIndirect(desired.EnableNetworkEgressMonitoring) && dcl.IsZeroValue(actual.EnableNetworkEgressMonitoring)) {
		c.Config.Logger.Infof("Diff in EnableNetworkEgressMonitoring. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableNetworkEgressMonitoring), dcl.SprintResource(actual.EnableNetworkEgressMonitoring))
		return true
	}
	if actual.ConsumptionMeteringConfig == nil && desired.ConsumptionMeteringConfig != nil && !dcl.IsEmptyValueIndirect(desired.ConsumptionMeteringConfig) {
		c.Config.Logger.Infof("desired ConsumptionMeteringConfig %s - but actually nil", dcl.SprintResource(desired.ConsumptionMeteringConfig))
		return true
	}
	if compareClusterResourceUsageExportConfigConsumptionMeteringConfig(c, desired.ConsumptionMeteringConfig, actual.ConsumptionMeteringConfig) && !dcl.IsZeroValue(desired.ConsumptionMeteringConfig) {
		c.Config.Logger.Infof("Diff in ConsumptionMeteringConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConsumptionMeteringConfig), dcl.SprintResource(actual.ConsumptionMeteringConfig))
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
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfigBigqueryDestination, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterResourceUsageExportConfigBigqueryDestination(c *Client, desired, actual *ClusterResourceUsageExportConfigBigqueryDestination) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DatasetId == nil && desired.DatasetId != nil && !dcl.IsEmptyValueIndirect(desired.DatasetId) {
		c.Config.Logger.Infof("desired DatasetId %s - but actually nil", dcl.SprintResource(desired.DatasetId))
		return true
	}
	if !reflect.DeepEqual(desired.DatasetId, actual.DatasetId) && !dcl.IsZeroValue(desired.DatasetId) && !(dcl.IsEmptyValueIndirect(desired.DatasetId) && dcl.IsZeroValue(actual.DatasetId)) {
		c.Config.Logger.Infof("Diff in DatasetId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DatasetId), dcl.SprintResource(actual.DatasetId))
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
			c.Config.Logger.Infof("Diff in ClusterResourceUsageExportConfigConsumptionMeteringConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterResourceUsageExportConfigConsumptionMeteringConfig(c *Client, desired, actual *ClusterResourceUsageExportConfigConsumptionMeteringConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) && !(dcl.IsEmptyValueIndirect(desired.Enabled) && dcl.IsZeroValue(actual.Enabled)) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
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
			c.Config.Logger.Infof("Diff in ClusterAuthenticatorGroupsConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterAuthenticatorGroupsConfig(c *Client, desired, actual *ClusterAuthenticatorGroupsConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) && !(dcl.IsEmptyValueIndirect(desired.Enabled) && dcl.IsZeroValue(actual.Enabled)) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	if actual.SecurityGroup == nil && desired.SecurityGroup != nil && !dcl.IsEmptyValueIndirect(desired.SecurityGroup) {
		c.Config.Logger.Infof("desired SecurityGroup %s - but actually nil", dcl.SprintResource(desired.SecurityGroup))
		return true
	}
	if !reflect.DeepEqual(desired.SecurityGroup, actual.SecurityGroup) && !dcl.IsZeroValue(desired.SecurityGroup) && !(dcl.IsEmptyValueIndirect(desired.SecurityGroup) && dcl.IsZeroValue(actual.SecurityGroup)) {
		c.Config.Logger.Infof("Diff in SecurityGroup. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecurityGroup), dcl.SprintResource(actual.SecurityGroup))
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
			c.Config.Logger.Infof("Diff in ClusterPrivateClusterConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterPrivateClusterConfig(c *Client, desired, actual *ClusterPrivateClusterConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EnablePrivateNodes == nil && desired.EnablePrivateNodes != nil && !dcl.IsEmptyValueIndirect(desired.EnablePrivateNodes) {
		c.Config.Logger.Infof("desired EnablePrivateNodes %s - but actually nil", dcl.SprintResource(desired.EnablePrivateNodes))
		return true
	}
	if !reflect.DeepEqual(desired.EnablePrivateNodes, actual.EnablePrivateNodes) && !dcl.IsZeroValue(desired.EnablePrivateNodes) && !(dcl.IsEmptyValueIndirect(desired.EnablePrivateNodes) && dcl.IsZeroValue(actual.EnablePrivateNodes)) {
		c.Config.Logger.Infof("Diff in EnablePrivateNodes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnablePrivateNodes), dcl.SprintResource(actual.EnablePrivateNodes))
		return true
	}
	if actual.EnablePrivateEndpoint == nil && desired.EnablePrivateEndpoint != nil && !dcl.IsEmptyValueIndirect(desired.EnablePrivateEndpoint) {
		c.Config.Logger.Infof("desired EnablePrivateEndpoint %s - but actually nil", dcl.SprintResource(desired.EnablePrivateEndpoint))
		return true
	}
	if !reflect.DeepEqual(desired.EnablePrivateEndpoint, actual.EnablePrivateEndpoint) && !dcl.IsZeroValue(desired.EnablePrivateEndpoint) && !(dcl.IsEmptyValueIndirect(desired.EnablePrivateEndpoint) && dcl.IsZeroValue(actual.EnablePrivateEndpoint)) {
		c.Config.Logger.Infof("Diff in EnablePrivateEndpoint. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnablePrivateEndpoint), dcl.SprintResource(actual.EnablePrivateEndpoint))
		return true
	}
	if actual.MasterIPv4CidrBlock == nil && desired.MasterIPv4CidrBlock != nil && !dcl.IsEmptyValueIndirect(desired.MasterIPv4CidrBlock) {
		c.Config.Logger.Infof("desired MasterIPv4CidrBlock %s - but actually nil", dcl.SprintResource(desired.MasterIPv4CidrBlock))
		return true
	}
	if !reflect.DeepEqual(desired.MasterIPv4CidrBlock, actual.MasterIPv4CidrBlock) && !dcl.IsZeroValue(desired.MasterIPv4CidrBlock) && !(dcl.IsEmptyValueIndirect(desired.MasterIPv4CidrBlock) && dcl.IsZeroValue(actual.MasterIPv4CidrBlock)) {
		c.Config.Logger.Infof("Diff in MasterIPv4CidrBlock. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MasterIPv4CidrBlock), dcl.SprintResource(actual.MasterIPv4CidrBlock))
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
			c.Config.Logger.Infof("Diff in ClusterDatabaseEncryption, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterDatabaseEncryption(c *Client, desired, actual *ClusterDatabaseEncryption) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.KeyName == nil && desired.KeyName != nil && !dcl.IsEmptyValueIndirect(desired.KeyName) {
		c.Config.Logger.Infof("desired KeyName %s - but actually nil", dcl.SprintResource(desired.KeyName))
		return true
	}
	if !reflect.DeepEqual(desired.KeyName, actual.KeyName) && !dcl.IsZeroValue(desired.KeyName) && !(dcl.IsEmptyValueIndirect(desired.KeyName) && dcl.IsZeroValue(actual.KeyName)) {
		c.Config.Logger.Infof("Diff in KeyName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KeyName), dcl.SprintResource(actual.KeyName))
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
			c.Config.Logger.Infof("Diff in ClusterVerticalPodAutoscaling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterVerticalPodAutoscaling(c *Client, desired, actual *ClusterVerticalPodAutoscaling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) && !(dcl.IsEmptyValueIndirect(desired.Enabled) && dcl.IsZeroValue(actual.Enabled)) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
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
			c.Config.Logger.Infof("Diff in ClusterShieldedNodes, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterShieldedNodes(c *Client, desired, actual *ClusterShieldedNodes) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) && !(dcl.IsEmptyValueIndirect(desired.Enabled) && dcl.IsZeroValue(actual.Enabled)) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
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
			c.Config.Logger.Infof("Diff in ClusterConditions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterConditions(c *Client, desired, actual *ClusterConditions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Code == nil && desired.Code != nil && !dcl.IsEmptyValueIndirect(desired.Code) {
		c.Config.Logger.Infof("desired Code %s - but actually nil", dcl.SprintResource(desired.Code))
		return true
	}
	if !reflect.DeepEqual(desired.Code, actual.Code) && !dcl.IsZeroValue(desired.Code) && !(dcl.IsEmptyValueIndirect(desired.Code) && dcl.IsZeroValue(actual.Code)) {
		c.Config.Logger.Infof("Diff in Code. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Code), dcl.SprintResource(actual.Code))
		return true
	}
	if actual.Message == nil && desired.Message != nil && !dcl.IsEmptyValueIndirect(desired.Message) {
		c.Config.Logger.Infof("desired Message %s - but actually nil", dcl.SprintResource(desired.Message))
		return true
	}
	if !reflect.DeepEqual(desired.Message, actual.Message) && !dcl.IsZeroValue(desired.Message) && !(dcl.IsEmptyValueIndirect(desired.Message) && dcl.IsZeroValue(actual.Message)) {
		c.Config.Logger.Infof("Diff in Message. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Message), dcl.SprintResource(actual.Message))
		return true
	}
	return false
}
func compareClusterNetworkPolicyProviderEnumSlice(c *Client, desired, actual []ClusterNetworkPolicyProviderEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterNetworkPolicyProviderEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterNetworkPolicyProviderEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterNetworkPolicyProviderEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterNetworkPolicyProviderEnum(c *Client, desired, actual *ClusterNetworkPolicyProviderEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareClusterDatabaseEncryptionStateEnumSlice(c *Client, desired, actual []ClusterDatabaseEncryptionStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ClusterDatabaseEncryptionStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareClusterDatabaseEncryptionStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ClusterDatabaseEncryptionStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareClusterDatabaseEncryptionStateEnum(c *Client, desired, actual *ClusterDatabaseEncryptionStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Cluster) urlNormalized() *Cluster {
	normalized := deepcopy.Copy(*r).(Cluster)
	normalized.Project = dcl.SelfLinkToName(r.Project)
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
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}:setMaintenancePolicy", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateAddonsConfig" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateAutoscaling" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateBinaryAuthorization" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateDatabaseEncryption" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateLegacyAbac" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}:setLegacyAbac", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateLocations" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateMasterAuthorizedNetworksConfig" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateMasterVersion" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateMonitoringAndLoggingService" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateShieldedNodes" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	if updateName == "updateVerticalPodAutoscaling" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{name}}", "https://container.googleapis.com/v1beta1/", userBasePath, fields), nil

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
	} else if !dcl.IsEmptyValueIndirect(v) {
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
		m["clusterIPv4Cidr"] = v
	}
	if v, err := expandClusterAddonsConfig(c, f.AddonsConfig); err != nil {
		return nil, fmt.Errorf("error expanding AddonsConfig into addonsConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["addonsConfig"] = v
	}
	if v := f.Subnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["subnetwork"] = v
	}
	if v, err := expandClusterNodePoolsSlice(c, f.NodePools); err != nil {
		return nil, fmt.Errorf("error expanding NodePools into nodePools: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["legacyAbac"] = v
	}
	if v, err := expandClusterNetworkPolicy(c, f.NetworkPolicy); err != nil {
		return nil, fmt.Errorf("error expanding NetworkPolicy into networkPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["networkPolicy"] = v
	}
	if v, err := expandClusterIPAllocationPolicy(c, f.IPAllocationPolicy); err != nil {
		return nil, fmt.Errorf("error expanding IPAllocationPolicy into ipAllocationPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ipAllocationPolicy"] = v
	}
	if v, err := expandClusterMasterAuthorizedNetworksConfig(c, f.MasterAuthorizedNetworksConfig); err != nil {
		return nil, fmt.Errorf("error expanding MasterAuthorizedNetworksConfig into masterAuthorizedNetworksConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["masterAuthorizedNetworksConfig"] = v
	}
	if v, err := expandClusterBinaryAuthorization(c, f.BinaryAuthorization); err != nil {
		return nil, fmt.Errorf("error expanding BinaryAuthorization into binaryAuthorization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["binaryAuthorization"] = v
	}
	if v, err := expandClusterAutoscaling(c, f.Autoscaling); err != nil {
		return nil, fmt.Errorf("error expanding Autoscaling into autoscaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoscaling"] = v
	}
	if v, err := expandClusterNetworkConfig(c, f.NetworkConfig); err != nil {
		return nil, fmt.Errorf("error expanding NetworkConfig into networkConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["networkConfig"] = v
	}
	if v, err := expandClusterMaintenancePolicy(c, f.MaintenancePolicy); err != nil {
		return nil, fmt.Errorf("error expanding MaintenancePolicy into maintenancePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maintenancePolicy"] = v
	}
	if v, err := expandClusterDefaultMaxPodsConstraint(c, f.DefaultMaxPodsConstraint); err != nil {
		return nil, fmt.Errorf("error expanding DefaultMaxPodsConstraint into defaultMaxPodsConstraint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultMaxPodsConstraint"] = v
	}
	if v, err := expandClusterResourceUsageExportConfig(c, f.ResourceUsageExportConfig); err != nil {
		return nil, fmt.Errorf("error expanding ResourceUsageExportConfig into resourceUsageExportConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resourceUsageExportConfig"] = v
	}
	if v, err := expandClusterAuthenticatorGroupsConfig(c, f.AuthenticatorGroupsConfig); err != nil {
		return nil, fmt.Errorf("error expanding AuthenticatorGroupsConfig into authenticatorGroupsConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["authenticatorGroupsConfig"] = v
	}
	if v, err := expandClusterPrivateClusterConfig(c, f.PrivateClusterConfig); err != nil {
		return nil, fmt.Errorf("error expanding PrivateClusterConfig into privateClusterConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["privateClusterConfig"] = v
	}
	if v, err := expandClusterDatabaseEncryption(c, f.DatabaseEncryption); err != nil {
		return nil, fmt.Errorf("error expanding DatabaseEncryption into databaseEncryption: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["databaseEncryption"] = v
	}
	if v, err := expandClusterVerticalPodAutoscaling(c, f.VerticalPodAutoscaling); err != nil {
		return nil, fmt.Errorf("error expanding VerticalPodAutoscaling into verticalPodAutoscaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["verticalPodAutoscaling"] = v
	}
	if v, err := expandClusterShieldedNodes(c, f.ShieldedNodes); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedNodes into shieldedNodes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
		m["nodeIPv4CidrSize"] = v
	}
	if v := f.ServicesIPv4Cidr; !dcl.IsEmptyValueIndirect(v) {
		m["servicesIPv4Cidr"] = v
	}
	if v := f.ExpireTime; !dcl.IsEmptyValueIndirect(v) {
		m["expireTime"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.EnableTPU; !dcl.IsEmptyValueIndirect(v) {
		m["enableTPU"] = v
	}
	if v := f.TPUIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["tpuIPv4CidrBlock"] = v
	}
	if v, err := expandClusterConditionsSlice(c, f.Conditions); err != nil {
		return nil, fmt.Errorf("error expanding Conditions into conditions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditions"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
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
	r.ClusterIPv4Cidr = dcl.FlattenString(m["clusterIPv4Cidr"])
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
	r.NodeIPv4CidrSize = dcl.FlattenInteger(m["nodeIPv4CidrSize"])
	r.ServicesIPv4Cidr = dcl.FlattenString(m["servicesIPv4Cidr"])
	r.ExpireTime = dcl.FlattenString(m["expireTime"])
	r.Location = dcl.FlattenString(m["location"])
	r.EnableTPU = dcl.FlattenBool(m["enableTPU"])
	r.TPUIPv4CidrBlock = dcl.FlattenString(m["tpuIPv4CidrBlock"])
	r.Conditions = flattenClusterConditionsSlice(c, m["conditions"])
	r.Project = dcl.FlattenString(m["project"])

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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	m["disabled"] = f.Disabled

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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.UseIPAliases; !dcl.IsEmptyValueIndirect(v) {
		m["useIPAliases"] = v
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
		m["clusterIPv4CidrBlock"] = v
	}
	if v := f.NodeIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["nodeIPv4CidrBlock"] = v
	}
	if v := f.ServicesIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["servicesIPv4CidrBlock"] = v
	}
	if v := f.TPUIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["tpuIPv4CidrBlock"] = v
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
	r.UseIPAliases = dcl.FlattenBool(m["useIPAliases"])
	r.CreateSubnetwork = dcl.FlattenBool(m["createSubnetwork"])
	r.SubnetworkName = dcl.FlattenString(m["subnetworkName"])
	r.ClusterSecondaryRangeName = dcl.FlattenString(m["clusterSecondaryRangeName"])
	r.ServicesSecondaryRangeName = dcl.FlattenString(m["servicesSecondaryRangeName"])
	r.ClusterIPv4CidrBlock = dcl.FlattenString(m["clusterIPv4CidrBlock"])
	r.NodeIPv4CidrBlock = dcl.FlattenString(m["nodeIPv4CidrBlock"])
	r.ServicesIPv4CidrBlock = dcl.FlattenString(m["servicesIPv4CidrBlock"])
	r.TPUIPv4CidrBlock = dcl.FlattenString(m["tpuIPv4CidrBlock"])

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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AutoUpgrade; !dcl.IsEmptyValueIndirect(v) {
		m["autoUpgrade"] = v
	}
	if v := f.AutoRepair; !dcl.IsEmptyValueIndirect(v) {
		m["autoRepair"] = v
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.Subnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["subnetwork"] = v
	}
	if v := f.EnableIntraNodeVisibility; !dcl.IsEmptyValueIndirect(v) {
		m["enableIntraNodeVisibility"] = v
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnablePrivateNodes; !dcl.IsEmptyValueIndirect(v) {
		m["enablePrivateNodes"] = v
	}
	if v := f.EnablePrivateEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["enablePrivateEndpoint"] = v
	}
	if v := f.MasterIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["masterIPv4CidrBlock"] = v
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
	r.MasterIPv4CidrBlock = dcl.FlattenString(m["masterIPv4CidrBlock"])
	r.PrivateEndpoint = dcl.FlattenString(m["privateEndpoint"])
	r.PublicEndpoint = dcl.FlattenString(m["publicEndpoint"])
	r.PeeringName = dcl.FlattenString(m["peeringName"])

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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
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

	return r
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
		items = append(items, *flattenClusterNetworkPolicyProviderEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenClusterDatabaseEncryptionStateEnum(item.(map[string]interface{})))
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
