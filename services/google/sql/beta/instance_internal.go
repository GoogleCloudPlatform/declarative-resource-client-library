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

func (r *Instance) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MaxDiskSize) {
		if err := r.MaxDiskSize.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CurrentDiskSize) {
		if err := r.CurrentDiskSize.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DiskEncryptionConfiguration) {
		if err := r.DiskEncryptionConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FailoverReplica) {
		if err := r.FailoverReplica.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MasterInstance) {
		if err := r.MasterInstance.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReplicaConfiguration) {
		if err := r.ReplicaConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ScheduledMaintenance) {
		if err := r.ScheduledMaintenance.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Settings) {
		if err := r.Settings.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ServerCaCert) {
		if err := r.ServerCaCert.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.OnPremisesConfiguration) {
		if err := r.OnPremisesConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DiskEncryptionStatus) {
		if err := r.DiskEncryptionStatus.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceMaxDiskSize) validate() error {
	return nil
}
func (r *InstanceCurrentDiskSize) validate() error {
	return nil
}
func (r *InstanceDiskEncryptionConfiguration) validate() error {
	return nil
}
func (r *InstanceFailoverReplica) validate() error {
	if !dcl.IsEmptyValueIndirect(r.FailoverInstance) {
		if err := r.FailoverInstance.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceFailoverReplicaFailoverInstance) validate() error {
	return nil
}
func (r *InstanceIPAddresses) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TimeToRetire) {
		if err := r.TimeToRetire.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceIPAddressesTimeToRetire) validate() error {
	return nil
}
func (r *InstanceMasterInstance) validate() error {
	return nil
}
func (r *InstanceReplicaConfiguration) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MysqlReplicaConfiguration) {
		if err := r.MysqlReplicaConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReplicaPoolConfiguration) {
		if err := r.ReplicaPoolConfiguration.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceReplicaConfigurationMysqlReplicaConfiguration) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MasterHeartbeatPeriod) {
		if err := r.MasterHeartbeatPeriod.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) validate() error {
	return nil
}
func (r *InstanceReplicaConfigurationReplicaPoolConfiguration) validate() error {
	if !dcl.IsEmptyValueIndirect(r.StaticPoolConfiguration) {
		if err := r.StaticPoolConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AutoscalingPoolConfiguration) {
		if err := r.AutoscalingPoolConfiguration.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) validate() error {
	return nil
}
func (r *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) validate() error {
	return nil
}
func (r *InstanceScheduledMaintenance) validate() error {
	if !dcl.IsEmptyValueIndirect(r.StartTime) {
		if err := r.StartTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceScheduledMaintenanceStartTime) validate() error {
	return nil
}
func (r *InstanceSettings) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SettingsVersion) {
		if err := r.SettingsVersion.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.StorageAutoResizeLimit) {
		if err := r.StorageAutoResizeLimit.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.IPConfiguration) {
		if err := r.IPConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LocationPreference) {
		if err := r.LocationPreference.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaintenanceWindow) {
		if err := r.MaintenanceWindow.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BackupConfiguration) {
		if err := r.BackupConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DataDiskSizeGb) {
		if err := r.DataDiskSizeGb.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ActiveDirectoryConfig) {
		if err := r.ActiveDirectoryConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.InsightsConfig) {
		if err := r.InsightsConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceSettingsSettingsVersion) validate() error {
	return nil
}
func (r *InstanceSettingsStorageAutoResizeLimit) validate() error {
	return nil
}
func (r *InstanceSettingsIPConfiguration) validate() error {
	return nil
}
func (r *InstanceSettingsIPConfigurationAuthorizedNetworks) validate() error {
	return nil
}
func (r *InstanceSettingsLocationPreference) validate() error {
	return nil
}
func (r *InstanceSettingsDatabaseFlags) validate() error {
	return nil
}
func (r *InstanceSettingsMaintenanceWindow) validate() error {
	return nil
}
func (r *InstanceSettingsBackupConfiguration) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BackupRetentionSettings) {
		if err := r.BackupRetentionSettings.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceSettingsBackupConfigurationBackupRetentionSettings) validate() error {
	return nil
}
func (r *InstanceSettingsDataDiskSizeGb) validate() error {
	return nil
}
func (r *InstanceSettingsActiveDirectoryConfig) validate() error {
	return nil
}
func (r *InstanceSettingsDenyMaintenancePeriods) validate() error {
	return nil
}
func (r *InstanceSettingsInsightsConfig) validate() error {
	return nil
}
func (r *InstanceReplicaInstances) validate() error {
	return nil
}
func (r *InstanceServerCaCert) validate() error {
	return nil
}
func (r *InstanceOnPremisesConfiguration) validate() error {
	return nil
}
func (r *InstanceDiskEncryptionStatus) validate() error {
	return nil
}

func instanceGetURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

func instanceListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/instances/", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func instanceCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/instances", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func instanceDeleteURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

// instanceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceApiOperation interface {
	do(context.Context, *Instance, *Client) error
}

// newUpdateInstanceUpdateRequest creates a request for an
// Instance resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceUpdateRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		req["state"] = v
	}
	if v, err := expandInstanceReplicaInstancesSlice(c, f.ReplicaInstances); err != nil {
		return nil, fmt.Errorf("error expanding ReplicaInstances into replicaInstances: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["replicaInstances"] = v
	}
	if v, err := expandInstanceServerCaCert(c, f.ServerCaCert); err != nil {
		return nil, fmt.Errorf("error expanding ServerCaCert into serverCaCert: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["serverCaCert"] = v
	}
	if v := f.IPv6Address; !dcl.IsEmptyValueIndirect(v) {
		req["ipv6Address"] = v
	}
	if v := f.ServiceAccountEmailAddress; !dcl.IsEmptyValueIndirect(v) {
		req["serviceAccountEmailAddress"] = v
	}
	if v, err := expandInstanceOnPremisesConfiguration(c, f.OnPremisesConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding OnPremisesConfiguration into onPremisesConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["onPremisesConfiguration"] = v
	}
	if v := f.SuspensionReason; !dcl.IsEmptyValueIndirect(v) {
		req["suspensionReason"] = v
	}
	if v, err := expandInstanceDiskEncryptionStatus(c, f.DiskEncryptionStatus); err != nil {
		return nil, fmt.Errorf("error expanding DiskEncryptionStatus into diskEncryptionStatus: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["diskEncryptionStatus"] = v
	}
	if v := f.InstanceUid; !dcl.IsEmptyValueIndirect(v) {
		req["instanceUid"] = v
	}
	return req, nil
}

// marshalUpdateInstanceUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceUpdateOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInstanceRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := instanceListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InstanceMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
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

type listInstanceOperation struct {
	Instances []map[string]interface{} `json:"instances"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listInstance(ctx context.Context, project, pageToken string, pageSize int32) ([]*Instance, string, error) {
	b, err := c.listInstanceRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInstanceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Instance
	for _, v := range m.Instances {
		res := flattenInstance(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInstance(ctx context.Context, f func(*Instance) bool, resources []*Instance) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInstance(ctx, res)
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

type deleteInstanceOperation struct{}

func (op *deleteInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {

	_, err := c.GetInstance(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Instance not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetInstance checking for existence. error: %v", err)
		return err
	}

	u, err := instanceDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		return err
	}
	_, err = c.GetInstance(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createInstanceOperation struct{}

func (op *createInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := instanceCreateURL(c.Config.BasePath, project)

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
	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetInstance(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getInstanceRaw(ctx context.Context, r *Instance) ([]byte, error) {

	u, err := instanceGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) instanceDiffsForRawDesired(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (initial, desired *Instance, diffs []instanceDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Instance
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Instance); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Instance, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInstance(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Instance resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Instance resource: %v", err)
		}
		c.Config.Logger.Info("Found that Instance resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Instance: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Instance: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInstanceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Instance: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Instance: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInstance(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInstanceInitialState(rawInitial, rawDesired *Instance) (*Instance, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInstanceDesiredState(rawDesired, rawInitial *Instance, opts ...dcl.ApplyOption) (*Instance, error) {

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Instance); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected Instance, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MaxDiskSize = canonicalizeInstanceMaxDiskSize(rawDesired.MaxDiskSize, nil, opts...)
		rawDesired.CurrentDiskSize = canonicalizeInstanceCurrentDiskSize(rawDesired.CurrentDiskSize, nil, opts...)
		rawDesired.DiskEncryptionConfiguration = canonicalizeInstanceDiskEncryptionConfiguration(rawDesired.DiskEncryptionConfiguration, nil, opts...)
		rawDesired.FailoverReplica = canonicalizeInstanceFailoverReplica(rawDesired.FailoverReplica, nil, opts...)
		rawDesired.MasterInstance = canonicalizeInstanceMasterInstance(rawDesired.MasterInstance, nil, opts...)
		rawDesired.ReplicaConfiguration = canonicalizeInstanceReplicaConfiguration(rawDesired.ReplicaConfiguration, nil, opts...)
		rawDesired.ScheduledMaintenance = canonicalizeInstanceScheduledMaintenance(rawDesired.ScheduledMaintenance, nil, opts...)
		rawDesired.Settings = canonicalizeInstanceSettings(rawDesired.Settings, nil, opts...)
		rawDesired.ServerCaCert = canonicalizeInstanceServerCaCert(rawDesired.ServerCaCert, nil, opts...)
		rawDesired.OnPremisesConfiguration = canonicalizeInstanceOnPremisesConfiguration(rawDesired.OnPremisesConfiguration, nil, opts...)
		rawDesired.DiskEncryptionStatus = canonicalizeInstanceDiskEncryptionStatus(rawDesired.DiskEncryptionStatus, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.BackendType) {
		rawDesired.BackendType = rawInitial.BackendType
	}
	if dcl.IsZeroValue(rawDesired.ConnectionName) {
		rawDesired.ConnectionName = rawInitial.ConnectionName
	}
	if dcl.IsZeroValue(rawDesired.DatabaseVersion) {
		rawDesired.DatabaseVersion = rawInitial.DatabaseVersion
	}
	if dcl.IsZeroValue(rawDesired.Etag) {
		rawDesired.Etag = rawInitial.Etag
	}
	if dcl.IsZeroValue(rawDesired.GceZone) {
		rawDesired.GceZone = rawInitial.GceZone
	}
	if dcl.IsZeroValue(rawDesired.InstanceType) {
		rawDesired.InstanceType = rawInitial.InstanceType
	}
	if dcl.IsZeroValue(rawDesired.MasterInstanceName) {
		rawDesired.MasterInstanceName = rawInitial.MasterInstanceName
	}
	rawDesired.MaxDiskSize = canonicalizeInstanceMaxDiskSize(rawDesired.MaxDiskSize, rawInitial.MaxDiskSize, opts...)
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.IsZeroValue(rawDesired.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.IsZeroValue(rawDesired.RootPassword) {
		rawDesired.RootPassword = rawInitial.RootPassword
	}
	rawDesired.CurrentDiskSize = canonicalizeInstanceCurrentDiskSize(rawDesired.CurrentDiskSize, rawInitial.CurrentDiskSize, opts...)
	rawDesired.DiskEncryptionConfiguration = canonicalizeInstanceDiskEncryptionConfiguration(rawDesired.DiskEncryptionConfiguration, rawInitial.DiskEncryptionConfiguration, opts...)
	rawDesired.FailoverReplica = canonicalizeInstanceFailoverReplica(rawDesired.FailoverReplica, rawInitial.FailoverReplica, opts...)
	if dcl.IsZeroValue(rawDesired.IPAddresses) {
		rawDesired.IPAddresses = rawInitial.IPAddresses
	}
	rawDesired.MasterInstance = canonicalizeInstanceMasterInstance(rawDesired.MasterInstance, rawInitial.MasterInstance, opts...)
	rawDesired.ReplicaConfiguration = canonicalizeInstanceReplicaConfiguration(rawDesired.ReplicaConfiguration, rawInitial.ReplicaConfiguration, opts...)
	rawDesired.ScheduledMaintenance = canonicalizeInstanceScheduledMaintenance(rawDesired.ScheduledMaintenance, rawInitial.ScheduledMaintenance, opts...)
	rawDesired.Settings = canonicalizeInstanceSettings(rawDesired.Settings, rawInitial.Settings, opts...)
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	if dcl.IsZeroValue(rawDesired.ReplicaInstances) {
		rawDesired.ReplicaInstances = rawInitial.ReplicaInstances
	}
	rawDesired.ServerCaCert = canonicalizeInstanceServerCaCert(rawDesired.ServerCaCert, rawInitial.ServerCaCert, opts...)
	if dcl.IsZeroValue(rawDesired.IPv6Address) {
		rawDesired.IPv6Address = rawInitial.IPv6Address
	}
	if dcl.IsZeroValue(rawDesired.ServiceAccountEmailAddress) {
		rawDesired.ServiceAccountEmailAddress = rawInitial.ServiceAccountEmailAddress
	}
	rawDesired.OnPremisesConfiguration = canonicalizeInstanceOnPremisesConfiguration(rawDesired.OnPremisesConfiguration, rawInitial.OnPremisesConfiguration, opts...)
	if dcl.IsZeroValue(rawDesired.SuspensionReason) {
		rawDesired.SuspensionReason = rawInitial.SuspensionReason
	}
	rawDesired.DiskEncryptionStatus = canonicalizeInstanceDiskEncryptionStatus(rawDesired.DiskEncryptionStatus, rawInitial.DiskEncryptionStatus, opts...)
	if dcl.IsZeroValue(rawDesired.InstanceUid) {
		rawDesired.InstanceUid = rawInitial.InstanceUid
	}

	return rawDesired, nil
}

func canonicalizeInstanceNewState(c *Client, rawNew, rawDesired *Instance) (*Instance, error) {

	if dcl.IsEmptyValueIndirect(rawNew.BackendType) && dcl.IsEmptyValueIndirect(rawDesired.BackendType) {
		rawNew.BackendType = rawDesired.BackendType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ConnectionName) && dcl.IsEmptyValueIndirect(rawDesired.ConnectionName) {
		rawNew.ConnectionName = rawDesired.ConnectionName
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DatabaseVersion) && dcl.IsEmptyValueIndirect(rawDesired.DatabaseVersion) {
		rawNew.DatabaseVersion = rawDesired.DatabaseVersion
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.GceZone) && dcl.IsEmptyValueIndirect(rawDesired.GceZone) {
		rawNew.GceZone = rawDesired.GceZone
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.InstanceType) && dcl.IsEmptyValueIndirect(rawDesired.InstanceType) {
		rawNew.InstanceType = rawDesired.InstanceType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MasterInstanceName) && dcl.IsEmptyValueIndirect(rawDesired.MasterInstanceName) {
		rawNew.MasterInstanceName = rawDesired.MasterInstanceName
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxDiskSize) && dcl.IsEmptyValueIndirect(rawDesired.MaxDiskSize) {
		rawNew.MaxDiskSize = rawDesired.MaxDiskSize
	} else {
		rawNew.MaxDiskSize = canonicalizeNewInstanceMaxDiskSize(c, rawDesired.MaxDiskSize, rawNew.MaxDiskSize)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RootPassword) && dcl.IsEmptyValueIndirect(rawDesired.RootPassword) {
		rawNew.RootPassword = rawDesired.RootPassword
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CurrentDiskSize) && dcl.IsEmptyValueIndirect(rawDesired.CurrentDiskSize) {
		rawNew.CurrentDiskSize = rawDesired.CurrentDiskSize
	} else {
		rawNew.CurrentDiskSize = canonicalizeNewInstanceCurrentDiskSize(c, rawDesired.CurrentDiskSize, rawNew.CurrentDiskSize)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskEncryptionConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.DiskEncryptionConfiguration) {
		rawNew.DiskEncryptionConfiguration = rawDesired.DiskEncryptionConfiguration
	} else {
		rawNew.DiskEncryptionConfiguration = canonicalizeNewInstanceDiskEncryptionConfiguration(c, rawDesired.DiskEncryptionConfiguration, rawNew.DiskEncryptionConfiguration)
	}

	if dcl.IsEmptyValueIndirect(rawNew.FailoverReplica) && dcl.IsEmptyValueIndirect(rawDesired.FailoverReplica) {
		rawNew.FailoverReplica = rawDesired.FailoverReplica
	} else {
		rawNew.FailoverReplica = canonicalizeNewInstanceFailoverReplica(c, rawDesired.FailoverReplica, rawNew.FailoverReplica)
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPAddresses) && dcl.IsEmptyValueIndirect(rawDesired.IPAddresses) {
		rawNew.IPAddresses = rawDesired.IPAddresses
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MasterInstance) && dcl.IsEmptyValueIndirect(rawDesired.MasterInstance) {
		rawNew.MasterInstance = rawDesired.MasterInstance
	} else {
		rawNew.MasterInstance = canonicalizeNewInstanceMasterInstance(c, rawDesired.MasterInstance, rawNew.MasterInstance)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReplicaConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.ReplicaConfiguration) {
		rawNew.ReplicaConfiguration = rawDesired.ReplicaConfiguration
	} else {
		rawNew.ReplicaConfiguration = canonicalizeNewInstanceReplicaConfiguration(c, rawDesired.ReplicaConfiguration, rawNew.ReplicaConfiguration)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ScheduledMaintenance) && dcl.IsEmptyValueIndirect(rawDesired.ScheduledMaintenance) {
		rawNew.ScheduledMaintenance = rawDesired.ScheduledMaintenance
	} else {
		rawNew.ScheduledMaintenance = canonicalizeNewInstanceScheduledMaintenance(c, rawDesired.ScheduledMaintenance, rawNew.ScheduledMaintenance)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Settings) && dcl.IsEmptyValueIndirect(rawDesired.Settings) {
		rawNew.Settings = rawDesired.Settings
	} else {
		rawNew.Settings = canonicalizeNewInstanceSettings(c, rawDesired.Settings, rawNew.Settings)
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReplicaInstances) && dcl.IsEmptyValueIndirect(rawDesired.ReplicaInstances) {
		rawNew.ReplicaInstances = rawDesired.ReplicaInstances
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServerCaCert) && dcl.IsEmptyValueIndirect(rawDesired.ServerCaCert) {
		rawNew.ServerCaCert = rawDesired.ServerCaCert
	} else {
		rawNew.ServerCaCert = canonicalizeNewInstanceServerCaCert(c, rawDesired.ServerCaCert, rawNew.ServerCaCert)
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPv6Address) && dcl.IsEmptyValueIndirect(rawDesired.IPv6Address) {
		rawNew.IPv6Address = rawDesired.IPv6Address
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccountEmailAddress) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccountEmailAddress) {
		rawNew.ServiceAccountEmailAddress = rawDesired.ServiceAccountEmailAddress
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.OnPremisesConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.OnPremisesConfiguration) {
		rawNew.OnPremisesConfiguration = rawDesired.OnPremisesConfiguration
	} else {
		rawNew.OnPremisesConfiguration = canonicalizeNewInstanceOnPremisesConfiguration(c, rawDesired.OnPremisesConfiguration, rawNew.OnPremisesConfiguration)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SuspensionReason) && dcl.IsEmptyValueIndirect(rawDesired.SuspensionReason) {
		rawNew.SuspensionReason = rawDesired.SuspensionReason
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskEncryptionStatus) && dcl.IsEmptyValueIndirect(rawDesired.DiskEncryptionStatus) {
		rawNew.DiskEncryptionStatus = rawDesired.DiskEncryptionStatus
	} else {
		rawNew.DiskEncryptionStatus = canonicalizeNewInstanceDiskEncryptionStatus(c, rawDesired.DiskEncryptionStatus, rawNew.DiskEncryptionStatus)
	}

	if dcl.IsEmptyValueIndirect(rawNew.InstanceUid) && dcl.IsEmptyValueIndirect(rawDesired.InstanceUid) {
		rawNew.InstanceUid = rawDesired.InstanceUid
	} else {
	}

	return rawNew, nil
}

func canonicalizeInstanceMaxDiskSize(des, initial *InstanceMaxDiskSize, opts ...dcl.ApplyOption) *InstanceMaxDiskSize {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceMaxDiskSize(c *Client, des, nw *InstanceMaxDiskSize) *InstanceMaxDiskSize {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceMaxDiskSizeSet(c *Client, des, nw []InstanceMaxDiskSize) []InstanceMaxDiskSize {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceMaxDiskSize
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceMaxDiskSize(c, &d, &n) {
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

func canonicalizeInstanceCurrentDiskSize(des, initial *InstanceCurrentDiskSize, opts ...dcl.ApplyOption) *InstanceCurrentDiskSize {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceCurrentDiskSize(c *Client, des, nw *InstanceCurrentDiskSize) *InstanceCurrentDiskSize {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceCurrentDiskSizeSet(c *Client, des, nw []InstanceCurrentDiskSize) []InstanceCurrentDiskSize {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceCurrentDiskSize
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceCurrentDiskSize(c, &d, &n) {
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

func canonicalizeInstanceDiskEncryptionConfiguration(des, initial *InstanceDiskEncryptionConfiguration, opts ...dcl.ApplyOption) *InstanceDiskEncryptionConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceDiskEncryptionConfiguration(c *Client, des, nw *InstanceDiskEncryptionConfiguration) *InstanceDiskEncryptionConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceDiskEncryptionConfigurationSet(c *Client, des, nw []InstanceDiskEncryptionConfiguration) []InstanceDiskEncryptionConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceDiskEncryptionConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceDiskEncryptionConfiguration(c, &d, &n) {
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

func canonicalizeInstanceFailoverReplica(des, initial *InstanceFailoverReplica, opts ...dcl.ApplyOption) *InstanceFailoverReplica {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Available) {
		des.Available = initial.Available
	}
	des.FailoverInstance = canonicalizeInstanceFailoverReplicaFailoverInstance(des.FailoverInstance, initial.FailoverInstance, opts...)

	return des
}

func canonicalizeNewInstanceFailoverReplica(c *Client, des, nw *InstanceFailoverReplica) *InstanceFailoverReplica {
	if des == nil || nw == nil {
		return nw
	}

	nw.FailoverInstance = canonicalizeNewInstanceFailoverReplicaFailoverInstance(c, des.FailoverInstance, nw.FailoverInstance)

	return nw
}

func canonicalizeNewInstanceFailoverReplicaSet(c *Client, des, nw []InstanceFailoverReplica) []InstanceFailoverReplica {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceFailoverReplica
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceFailoverReplica(c, &d, &n) {
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

func canonicalizeInstanceFailoverReplicaFailoverInstance(des, initial *InstanceFailoverReplicaFailoverInstance, opts ...dcl.ApplyOption) *InstanceFailoverReplicaFailoverInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Region) {
		des.Region = initial.Region
	}

	return des
}

func canonicalizeNewInstanceFailoverReplicaFailoverInstance(c *Client, des, nw *InstanceFailoverReplicaFailoverInstance) *InstanceFailoverReplicaFailoverInstance {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceFailoverReplicaFailoverInstanceSet(c *Client, des, nw []InstanceFailoverReplicaFailoverInstance) []InstanceFailoverReplicaFailoverInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceFailoverReplicaFailoverInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceFailoverReplicaFailoverInstance(c, &d, &n) {
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

func canonicalizeInstanceIPAddresses(des, initial *InstanceIPAddresses, opts ...dcl.ApplyOption) *InstanceIPAddresses {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.IsZeroValue(des.IPAddress) {
		des.IPAddress = initial.IPAddress
	}
	des.TimeToRetire = canonicalizeInstanceIPAddressesTimeToRetire(des.TimeToRetire, initial.TimeToRetire, opts...)

	return des
}

func canonicalizeNewInstanceIPAddresses(c *Client, des, nw *InstanceIPAddresses) *InstanceIPAddresses {
	if des == nil || nw == nil {
		return nw
	}

	nw.TimeToRetire = canonicalizeNewInstanceIPAddressesTimeToRetire(c, des.TimeToRetire, nw.TimeToRetire)

	return nw
}

func canonicalizeNewInstanceIPAddressesSet(c *Client, des, nw []InstanceIPAddresses) []InstanceIPAddresses {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceIPAddresses
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceIPAddresses(c, &d, &n) {
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

func canonicalizeInstanceIPAddressesTimeToRetire(des, initial *InstanceIPAddressesTimeToRetire, opts ...dcl.ApplyOption) *InstanceIPAddressesTimeToRetire {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewInstanceIPAddressesTimeToRetire(c *Client, des, nw *InstanceIPAddressesTimeToRetire) *InstanceIPAddressesTimeToRetire {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceIPAddressesTimeToRetireSet(c *Client, des, nw []InstanceIPAddressesTimeToRetire) []InstanceIPAddressesTimeToRetire {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceIPAddressesTimeToRetire
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceIPAddressesTimeToRetire(c, &d, &n) {
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

func canonicalizeInstanceMasterInstance(des, initial *InstanceMasterInstance, opts ...dcl.ApplyOption) *InstanceMasterInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Region) {
		des.Region = initial.Region
	}

	return des
}

func canonicalizeNewInstanceMasterInstance(c *Client, des, nw *InstanceMasterInstance) *InstanceMasterInstance {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceMasterInstanceSet(c *Client, des, nw []InstanceMasterInstance) []InstanceMasterInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceMasterInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceMasterInstance(c, &d, &n) {
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

func canonicalizeInstanceReplicaConfiguration(des, initial *InstanceReplicaConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	des.MysqlReplicaConfiguration = canonicalizeInstanceReplicaConfigurationMysqlReplicaConfiguration(des.MysqlReplicaConfiguration, initial.MysqlReplicaConfiguration, opts...)
	if dcl.IsZeroValue(des.FailoverTarget) {
		des.FailoverTarget = initial.FailoverTarget
	}
	des.ReplicaPoolConfiguration = canonicalizeInstanceReplicaConfigurationReplicaPoolConfiguration(des.ReplicaPoolConfiguration, initial.ReplicaPoolConfiguration, opts...)

	return des
}

func canonicalizeNewInstanceReplicaConfiguration(c *Client, des, nw *InstanceReplicaConfiguration) *InstanceReplicaConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	nw.MysqlReplicaConfiguration = canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfiguration(c, des.MysqlReplicaConfiguration, nw.MysqlReplicaConfiguration)
	nw.ReplicaPoolConfiguration = canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfiguration(c, des.ReplicaPoolConfiguration, nw.ReplicaPoolConfiguration)

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationSet(c *Client, des, nw []InstanceReplicaConfiguration) []InstanceReplicaConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfiguration(c, &d, &n) {
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

func canonicalizeInstanceReplicaConfigurationMysqlReplicaConfiguration(des, initial *InstanceReplicaConfigurationMysqlReplicaConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DumpFilePath) {
		des.DumpFilePath = initial.DumpFilePath
	}
	if dcl.IsZeroValue(des.Username) {
		des.Username = initial.Username
	}
	if dcl.IsZeroValue(des.Password) {
		des.Password = initial.Password
	}
	if dcl.IsZeroValue(des.ConnectRetryInterval) {
		des.ConnectRetryInterval = initial.ConnectRetryInterval
	}
	des.MasterHeartbeatPeriod = canonicalizeInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(des.MasterHeartbeatPeriod, initial.MasterHeartbeatPeriod, opts...)
	if dcl.IsZeroValue(des.CaCertificate) {
		des.CaCertificate = initial.CaCertificate
	}
	if dcl.IsZeroValue(des.ClientCertificate) {
		des.ClientCertificate = initial.ClientCertificate
	}
	if dcl.IsZeroValue(des.ClientKey) {
		des.ClientKey = initial.ClientKey
	}
	if dcl.IsZeroValue(des.SslCipher) {
		des.SslCipher = initial.SslCipher
	}
	if dcl.IsZeroValue(des.VerifyServerCertificate) {
		des.VerifyServerCertificate = initial.VerifyServerCertificate
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfiguration(c *Client, des, nw *InstanceReplicaConfigurationMysqlReplicaConfiguration) *InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	nw.MasterHeartbeatPeriod = canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, des.MasterHeartbeatPeriod, nw.MasterHeartbeatPeriod)

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationSet(c *Client, des, nw []InstanceReplicaConfigurationMysqlReplicaConfiguration) []InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationMysqlReplicaConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &d, &n) {
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

func canonicalizeInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(des, initial *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c *Client, des, nw *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSet(c *Client, des, nw []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &d, &n) {
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

func canonicalizeInstanceReplicaConfigurationReplicaPoolConfiguration(des, initial *InstanceReplicaConfigurationReplicaPoolConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationReplicaPoolConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	des.StaticPoolConfiguration = canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(des.StaticPoolConfiguration, initial.StaticPoolConfiguration, opts...)
	des.AutoscalingPoolConfiguration = canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(des.AutoscalingPoolConfiguration, initial.AutoscalingPoolConfiguration, opts...)
	if dcl.IsZeroValue(des.ReplicaCount) {
		des.ReplicaCount = initial.ReplicaCount
	}
	if dcl.IsZeroValue(des.ExposeReplicaIP) {
		des.ExposeReplicaIP = initial.ExposeReplicaIP
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfiguration(c *Client, des, nw *InstanceReplicaConfigurationReplicaPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	nw.StaticPoolConfiguration = canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, des.StaticPoolConfiguration, nw.StaticPoolConfiguration)
	nw.AutoscalingPoolConfiguration = canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, des.AutoscalingPoolConfiguration, nw.AutoscalingPoolConfiguration)

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationSet(c *Client, des, nw []InstanceReplicaConfigurationReplicaPoolConfiguration) []InstanceReplicaConfigurationReplicaPoolConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationReplicaPoolConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationReplicaPoolConfiguration(c, &d, &n) {
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

func canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(des, initial *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.ReplicaCount) {
		des.ReplicaCount = initial.ReplicaCount
	}
	if dcl.IsZeroValue(des.ExposeReplicaIP) {
		des.ExposeReplicaIP = initial.ExposeReplicaIP
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c *Client, des, nw *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSet(c *Client, des, nw []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &d, &n) {
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

func canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(des, initial *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.MinReplicaCount) {
		des.MinReplicaCount = initial.MinReplicaCount
	}
	if dcl.IsZeroValue(des.MaxReplicaCount) {
		des.MaxReplicaCount = initial.MaxReplicaCount
	}
	if dcl.IsZeroValue(des.TargetCpuUtil) {
		des.TargetCpuUtil = initial.TargetCpuUtil
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c *Client, des, nw *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSet(c *Client, des, nw []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &d, &n) {
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

func canonicalizeInstanceScheduledMaintenance(des, initial *InstanceScheduledMaintenance, opts ...dcl.ApplyOption) *InstanceScheduledMaintenance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.StartTime = canonicalizeInstanceScheduledMaintenanceStartTime(des.StartTime, initial.StartTime, opts...)
	if dcl.IsZeroValue(des.CanDefer) {
		des.CanDefer = initial.CanDefer
	}
	if dcl.IsZeroValue(des.CanReschedule) {
		des.CanReschedule = initial.CanReschedule
	}

	return des
}

func canonicalizeNewInstanceScheduledMaintenance(c *Client, des, nw *InstanceScheduledMaintenance) *InstanceScheduledMaintenance {
	if des == nil || nw == nil {
		return nw
	}

	nw.StartTime = canonicalizeNewInstanceScheduledMaintenanceStartTime(c, des.StartTime, nw.StartTime)

	return nw
}

func canonicalizeNewInstanceScheduledMaintenanceSet(c *Client, des, nw []InstanceScheduledMaintenance) []InstanceScheduledMaintenance {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceScheduledMaintenance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceScheduledMaintenance(c, &d, &n) {
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

func canonicalizeInstanceScheduledMaintenanceStartTime(des, initial *InstanceScheduledMaintenanceStartTime, opts ...dcl.ApplyOption) *InstanceScheduledMaintenanceStartTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewInstanceScheduledMaintenanceStartTime(c *Client, des, nw *InstanceScheduledMaintenanceStartTime) *InstanceScheduledMaintenanceStartTime {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceScheduledMaintenanceStartTimeSet(c *Client, des, nw []InstanceScheduledMaintenanceStartTime) []InstanceScheduledMaintenanceStartTime {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceScheduledMaintenanceStartTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceScheduledMaintenanceStartTime(c, &d, &n) {
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

func canonicalizeInstanceSettings(des, initial *InstanceSettings, opts ...dcl.ApplyOption) *InstanceSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AuthorizedGaeApplications) {
		des.AuthorizedGaeApplications = initial.AuthorizedGaeApplications
	}
	if dcl.IsZeroValue(des.Tier) {
		des.Tier = initial.Tier
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.AvailabilityType) {
		des.AvailabilityType = initial.AvailabilityType
	}
	if dcl.IsZeroValue(des.PricingPlan) {
		des.PricingPlan = initial.PricingPlan
	}
	if dcl.IsZeroValue(des.ReplicationType) {
		des.ReplicationType = initial.ReplicationType
	}
	if dcl.IsZeroValue(des.ActivationPolicy) {
		des.ActivationPolicy = initial.ActivationPolicy
	}
	if dcl.IsZeroValue(des.StorageAutoResize) {
		des.StorageAutoResize = initial.StorageAutoResize
	}
	if dcl.IsZeroValue(des.DataDiskType) {
		des.DataDiskType = initial.DataDiskType
	}
	if dcl.IsZeroValue(des.DatabaseReplicationEnabled) {
		des.DatabaseReplicationEnabled = initial.DatabaseReplicationEnabled
	}
	if dcl.IsZeroValue(des.CrashSafeReplicationEnabled) {
		des.CrashSafeReplicationEnabled = initial.CrashSafeReplicationEnabled
	}
	des.SettingsVersion = canonicalizeInstanceSettingsSettingsVersion(des.SettingsVersion, initial.SettingsVersion, opts...)
	if dcl.IsZeroValue(des.UserLabels) {
		des.UserLabels = initial.UserLabels
	}
	des.StorageAutoResizeLimit = canonicalizeInstanceSettingsStorageAutoResizeLimit(des.StorageAutoResizeLimit, initial.StorageAutoResizeLimit, opts...)
	des.IPConfiguration = canonicalizeInstanceSettingsIPConfiguration(des.IPConfiguration, initial.IPConfiguration, opts...)
	des.LocationPreference = canonicalizeInstanceSettingsLocationPreference(des.LocationPreference, initial.LocationPreference, opts...)
	if dcl.IsZeroValue(des.DatabaseFlags) {
		des.DatabaseFlags = initial.DatabaseFlags
	}
	des.MaintenanceWindow = canonicalizeInstanceSettingsMaintenanceWindow(des.MaintenanceWindow, initial.MaintenanceWindow, opts...)
	des.BackupConfiguration = canonicalizeInstanceSettingsBackupConfiguration(des.BackupConfiguration, initial.BackupConfiguration, opts...)
	des.DataDiskSizeGb = canonicalizeInstanceSettingsDataDiskSizeGb(des.DataDiskSizeGb, initial.DataDiskSizeGb, opts...)
	des.ActiveDirectoryConfig = canonicalizeInstanceSettingsActiveDirectoryConfig(des.ActiveDirectoryConfig, initial.ActiveDirectoryConfig, opts...)
	if dcl.IsZeroValue(des.Collation) {
		des.Collation = initial.Collation
	}
	if dcl.IsZeroValue(des.DenyMaintenancePeriods) {
		des.DenyMaintenancePeriods = initial.DenyMaintenancePeriods
	}
	des.InsightsConfig = canonicalizeInstanceSettingsInsightsConfig(des.InsightsConfig, initial.InsightsConfig, opts...)

	return des
}

func canonicalizeNewInstanceSettings(c *Client, des, nw *InstanceSettings) *InstanceSettings {
	if des == nil || nw == nil {
		return nw
	}

	nw.SettingsVersion = canonicalizeNewInstanceSettingsSettingsVersion(c, des.SettingsVersion, nw.SettingsVersion)
	nw.StorageAutoResizeLimit = canonicalizeNewInstanceSettingsStorageAutoResizeLimit(c, des.StorageAutoResizeLimit, nw.StorageAutoResizeLimit)
	nw.IPConfiguration = canonicalizeNewInstanceSettingsIPConfiguration(c, des.IPConfiguration, nw.IPConfiguration)
	nw.LocationPreference = canonicalizeNewInstanceSettingsLocationPreference(c, des.LocationPreference, nw.LocationPreference)
	nw.MaintenanceWindow = canonicalizeNewInstanceSettingsMaintenanceWindow(c, des.MaintenanceWindow, nw.MaintenanceWindow)
	nw.BackupConfiguration = canonicalizeNewInstanceSettingsBackupConfiguration(c, des.BackupConfiguration, nw.BackupConfiguration)
	nw.DataDiskSizeGb = canonicalizeNewInstanceSettingsDataDiskSizeGb(c, des.DataDiskSizeGb, nw.DataDiskSizeGb)
	nw.ActiveDirectoryConfig = canonicalizeNewInstanceSettingsActiveDirectoryConfig(c, des.ActiveDirectoryConfig, nw.ActiveDirectoryConfig)
	nw.InsightsConfig = canonicalizeNewInstanceSettingsInsightsConfig(c, des.InsightsConfig, nw.InsightsConfig)

	return nw
}

func canonicalizeNewInstanceSettingsSet(c *Client, des, nw []InstanceSettings) []InstanceSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettings(c, &d, &n) {
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

func canonicalizeInstanceSettingsSettingsVersion(des, initial *InstanceSettingsSettingsVersion, opts ...dcl.ApplyOption) *InstanceSettingsSettingsVersion {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceSettingsSettingsVersion(c *Client, des, nw *InstanceSettingsSettingsVersion) *InstanceSettingsSettingsVersion {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsSettingsVersionSet(c *Client, des, nw []InstanceSettingsSettingsVersion) []InstanceSettingsSettingsVersion {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsSettingsVersion
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsSettingsVersion(c, &d, &n) {
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

func canonicalizeInstanceSettingsStorageAutoResizeLimit(des, initial *InstanceSettingsStorageAutoResizeLimit, opts ...dcl.ApplyOption) *InstanceSettingsStorageAutoResizeLimit {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceSettingsStorageAutoResizeLimit(c *Client, des, nw *InstanceSettingsStorageAutoResizeLimit) *InstanceSettingsStorageAutoResizeLimit {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsStorageAutoResizeLimitSet(c *Client, des, nw []InstanceSettingsStorageAutoResizeLimit) []InstanceSettingsStorageAutoResizeLimit {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsStorageAutoResizeLimit
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsStorageAutoResizeLimit(c, &d, &n) {
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

func canonicalizeInstanceSettingsIPConfiguration(des, initial *InstanceSettingsIPConfiguration, opts ...dcl.ApplyOption) *InstanceSettingsIPConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.IPv4Enabled) {
		des.IPv4Enabled = initial.IPv4Enabled
	}
	if dcl.NameToSelfLink(des.PrivateNetwork, initial.PrivateNetwork) || dcl.IsZeroValue(des.PrivateNetwork) {
		des.PrivateNetwork = initial.PrivateNetwork
	}
	if dcl.IsZeroValue(des.RequireSsl) {
		des.RequireSsl = initial.RequireSsl
	}
	if dcl.IsZeroValue(des.AuthorizedNetworks) {
		des.AuthorizedNetworks = initial.AuthorizedNetworks
	}

	return des
}

func canonicalizeNewInstanceSettingsIPConfiguration(c *Client, des, nw *InstanceSettingsIPConfiguration) *InstanceSettingsIPConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.PrivateNetwork, nw.PrivateNetwork) || dcl.IsZeroValue(des.PrivateNetwork) {
		nw.PrivateNetwork = des.PrivateNetwork
	}

	return nw
}

func canonicalizeNewInstanceSettingsIPConfigurationSet(c *Client, des, nw []InstanceSettingsIPConfiguration) []InstanceSettingsIPConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsIPConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsIPConfiguration(c, &d, &n) {
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

func canonicalizeInstanceSettingsIPConfigurationAuthorizedNetworks(des, initial *InstanceSettingsIPConfigurationAuthorizedNetworks, opts ...dcl.ApplyOption) *InstanceSettingsIPConfigurationAuthorizedNetworks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}
	if dcl.IsZeroValue(des.ExpirationTime) {
		des.ExpirationTime = initial.ExpirationTime
	}
	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceSettingsIPConfigurationAuthorizedNetworks(c *Client, des, nw *InstanceSettingsIPConfigurationAuthorizedNetworks) *InstanceSettingsIPConfigurationAuthorizedNetworks {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsIPConfigurationAuthorizedNetworksSet(c *Client, des, nw []InstanceSettingsIPConfigurationAuthorizedNetworks) []InstanceSettingsIPConfigurationAuthorizedNetworks {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsIPConfigurationAuthorizedNetworks
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsIPConfigurationAuthorizedNetworks(c, &d, &n) {
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

func canonicalizeInstanceSettingsLocationPreference(des, initial *InstanceSettingsLocationPreference, opts ...dcl.ApplyOption) *InstanceSettingsLocationPreference {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Zone) {
		des.Zone = initial.Zone
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceSettingsLocationPreference(c *Client, des, nw *InstanceSettingsLocationPreference) *InstanceSettingsLocationPreference {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsLocationPreferenceSet(c *Client, des, nw []InstanceSettingsLocationPreference) []InstanceSettingsLocationPreference {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsLocationPreference
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsLocationPreference(c, &d, &n) {
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

func canonicalizeInstanceSettingsDatabaseFlags(des, initial *InstanceSettingsDatabaseFlags, opts ...dcl.ApplyOption) *InstanceSettingsDatabaseFlags {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceSettingsDatabaseFlags(c *Client, des, nw *InstanceSettingsDatabaseFlags) *InstanceSettingsDatabaseFlags {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsDatabaseFlagsSet(c *Client, des, nw []InstanceSettingsDatabaseFlags) []InstanceSettingsDatabaseFlags {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsDatabaseFlags
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsDatabaseFlags(c, &d, &n) {
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

func canonicalizeInstanceSettingsMaintenanceWindow(des, initial *InstanceSettingsMaintenanceWindow, opts ...dcl.ApplyOption) *InstanceSettingsMaintenanceWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Hour) {
		des.Hour = initial.Hour
	}
	if dcl.IsZeroValue(des.Day) {
		des.Day = initial.Day
	}
	if dcl.IsZeroValue(des.UpdateTrack) {
		des.UpdateTrack = initial.UpdateTrack
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceSettingsMaintenanceWindow(c *Client, des, nw *InstanceSettingsMaintenanceWindow) *InstanceSettingsMaintenanceWindow {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsMaintenanceWindowSet(c *Client, des, nw []InstanceSettingsMaintenanceWindow) []InstanceSettingsMaintenanceWindow {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsMaintenanceWindow
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsMaintenanceWindow(c, &d, &n) {
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

func canonicalizeInstanceSettingsBackupConfiguration(des, initial *InstanceSettingsBackupConfiguration, opts ...dcl.ApplyOption) *InstanceSettingsBackupConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.StartTime) {
		des.StartTime = initial.StartTime
	}
	if dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.BinaryLogEnabled) {
		des.BinaryLogEnabled = initial.BinaryLogEnabled
	}
	if dcl.IsZeroValue(des.Location) {
		des.Location = initial.Location
	}
	des.BackupRetentionSettings = canonicalizeInstanceSettingsBackupConfigurationBackupRetentionSettings(des.BackupRetentionSettings, initial.BackupRetentionSettings, opts...)
	if dcl.IsZeroValue(des.TransactionLogRetentionDays) {
		des.TransactionLogRetentionDays = initial.TransactionLogRetentionDays
	}

	return des
}

func canonicalizeNewInstanceSettingsBackupConfiguration(c *Client, des, nw *InstanceSettingsBackupConfiguration) *InstanceSettingsBackupConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	nw.BackupRetentionSettings = canonicalizeNewInstanceSettingsBackupConfigurationBackupRetentionSettings(c, des.BackupRetentionSettings, nw.BackupRetentionSettings)

	return nw
}

func canonicalizeNewInstanceSettingsBackupConfigurationSet(c *Client, des, nw []InstanceSettingsBackupConfiguration) []InstanceSettingsBackupConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsBackupConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsBackupConfiguration(c, &d, &n) {
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

func canonicalizeInstanceSettingsBackupConfigurationBackupRetentionSettings(des, initial *InstanceSettingsBackupConfigurationBackupRetentionSettings, opts ...dcl.ApplyOption) *InstanceSettingsBackupConfigurationBackupRetentionSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RetentionUnit) {
		des.RetentionUnit = initial.RetentionUnit
	}
	if dcl.IsZeroValue(des.RetainedBackups) {
		des.RetainedBackups = initial.RetainedBackups
	}

	return des
}

func canonicalizeNewInstanceSettingsBackupConfigurationBackupRetentionSettings(c *Client, des, nw *InstanceSettingsBackupConfigurationBackupRetentionSettings) *InstanceSettingsBackupConfigurationBackupRetentionSettings {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsBackupConfigurationBackupRetentionSettingsSet(c *Client, des, nw []InstanceSettingsBackupConfigurationBackupRetentionSettings) []InstanceSettingsBackupConfigurationBackupRetentionSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsBackupConfigurationBackupRetentionSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &d, &n) {
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

func canonicalizeInstanceSettingsDataDiskSizeGb(des, initial *InstanceSettingsDataDiskSizeGb, opts ...dcl.ApplyOption) *InstanceSettingsDataDiskSizeGb {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceSettingsDataDiskSizeGb(c *Client, des, nw *InstanceSettingsDataDiskSizeGb) *InstanceSettingsDataDiskSizeGb {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsDataDiskSizeGbSet(c *Client, des, nw []InstanceSettingsDataDiskSizeGb) []InstanceSettingsDataDiskSizeGb {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsDataDiskSizeGb
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsDataDiskSizeGb(c, &d, &n) {
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

func canonicalizeInstanceSettingsActiveDirectoryConfig(des, initial *InstanceSettingsActiveDirectoryConfig, opts ...dcl.ApplyOption) *InstanceSettingsActiveDirectoryConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.Domain) {
		des.Domain = initial.Domain
	}

	return des
}

func canonicalizeNewInstanceSettingsActiveDirectoryConfig(c *Client, des, nw *InstanceSettingsActiveDirectoryConfig) *InstanceSettingsActiveDirectoryConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsActiveDirectoryConfigSet(c *Client, des, nw []InstanceSettingsActiveDirectoryConfig) []InstanceSettingsActiveDirectoryConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsActiveDirectoryConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsActiveDirectoryConfig(c, &d, &n) {
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

func canonicalizeInstanceSettingsDenyMaintenancePeriods(des, initial *InstanceSettingsDenyMaintenancePeriods, opts ...dcl.ApplyOption) *InstanceSettingsDenyMaintenancePeriods {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.StartDate) {
		des.StartDate = initial.StartDate
	}
	if dcl.IsZeroValue(des.EndDate) {
		des.EndDate = initial.EndDate
	}
	if dcl.IsZeroValue(des.Time) {
		des.Time = initial.Time
	}

	return des
}

func canonicalizeNewInstanceSettingsDenyMaintenancePeriods(c *Client, des, nw *InstanceSettingsDenyMaintenancePeriods) *InstanceSettingsDenyMaintenancePeriods {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsDenyMaintenancePeriodsSet(c *Client, des, nw []InstanceSettingsDenyMaintenancePeriods) []InstanceSettingsDenyMaintenancePeriods {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsDenyMaintenancePeriods
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsDenyMaintenancePeriods(c, &d, &n) {
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

func canonicalizeInstanceSettingsInsightsConfig(des, initial *InstanceSettingsInsightsConfig, opts ...dcl.ApplyOption) *InstanceSettingsInsightsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.QueryInsightsEnabled) {
		des.QueryInsightsEnabled = initial.QueryInsightsEnabled
	}
	if dcl.IsZeroValue(des.RecordClientAddress) {
		des.RecordClientAddress = initial.RecordClientAddress
	}
	if dcl.IsZeroValue(des.RecordApplicationTags) {
		des.RecordApplicationTags = initial.RecordApplicationTags
	}
	if dcl.IsZeroValue(des.QueryStringLength) {
		des.QueryStringLength = initial.QueryStringLength
	}

	return des
}

func canonicalizeNewInstanceSettingsInsightsConfig(c *Client, des, nw *InstanceSettingsInsightsConfig) *InstanceSettingsInsightsConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsInsightsConfigSet(c *Client, des, nw []InstanceSettingsInsightsConfig) []InstanceSettingsInsightsConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsInsightsConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsInsightsConfig(c, &d, &n) {
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

func canonicalizeInstanceReplicaInstances(des, initial *InstanceReplicaInstances, opts ...dcl.ApplyOption) *InstanceReplicaInstances {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.IsZeroValue(des.Region) {
		des.Region = initial.Region
	}

	return des
}

func canonicalizeNewInstanceReplicaInstances(c *Client, des, nw *InstanceReplicaInstances) *InstanceReplicaInstances {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceReplicaInstancesSet(c *Client, des, nw []InstanceReplicaInstances) []InstanceReplicaInstances {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaInstances
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaInstances(c, &d, &n) {
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

func canonicalizeInstanceServerCaCert(des, initial *InstanceServerCaCert, opts ...dcl.ApplyOption) *InstanceServerCaCert {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.CertSerialNumber) {
		des.CertSerialNumber = initial.CertSerialNumber
	}
	if dcl.IsZeroValue(des.Cert) {
		des.Cert = initial.Cert
	}
	if dcl.IsZeroValue(des.CreateTime) {
		des.CreateTime = initial.CreateTime
	}
	if dcl.IsZeroValue(des.CommonName) {
		des.CommonName = initial.CommonName
	}
	if dcl.IsZeroValue(des.ExpirationTime) {
		des.ExpirationTime = initial.ExpirationTime
	}
	if dcl.IsZeroValue(des.Sha1Fingerprint) {
		des.Sha1Fingerprint = initial.Sha1Fingerprint
	}
	if dcl.IsZeroValue(des.Instance) {
		des.Instance = initial.Instance
	}

	return des
}

func canonicalizeNewInstanceServerCaCert(c *Client, des, nw *InstanceServerCaCert) *InstanceServerCaCert {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceServerCaCertSet(c *Client, des, nw []InstanceServerCaCert) []InstanceServerCaCert {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceServerCaCert
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceServerCaCert(c, &d, &n) {
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

func canonicalizeInstanceOnPremisesConfiguration(des, initial *InstanceOnPremisesConfiguration, opts ...dcl.ApplyOption) *InstanceOnPremisesConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.HostPort) {
		des.HostPort = initial.HostPort
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.Username) {
		des.Username = initial.Username
	}
	if dcl.IsZeroValue(des.Password) {
		des.Password = initial.Password
	}
	if dcl.IsZeroValue(des.CaCertificate) {
		des.CaCertificate = initial.CaCertificate
	}
	if dcl.IsZeroValue(des.ClientCertificate) {
		des.ClientCertificate = initial.ClientCertificate
	}
	if dcl.IsZeroValue(des.ClientKey) {
		des.ClientKey = initial.ClientKey
	}
	if dcl.IsZeroValue(des.DumpFilePath) {
		des.DumpFilePath = initial.DumpFilePath
	}
	if dcl.IsZeroValue(des.Database) {
		des.Database = initial.Database
	}
	if dcl.IsZeroValue(des.ReplicatedDatabases) {
		des.ReplicatedDatabases = initial.ReplicatedDatabases
	}

	return des
}

func canonicalizeNewInstanceOnPremisesConfiguration(c *Client, des, nw *InstanceOnPremisesConfiguration) *InstanceOnPremisesConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceOnPremisesConfigurationSet(c *Client, des, nw []InstanceOnPremisesConfiguration) []InstanceOnPremisesConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceOnPremisesConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceOnPremisesConfiguration(c, &d, &n) {
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

func canonicalizeInstanceDiskEncryptionStatus(des, initial *InstanceDiskEncryptionStatus, opts ...dcl.ApplyOption) *InstanceDiskEncryptionStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Instance)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.KmsKeyVersionName) {
		des.KmsKeyVersionName = initial.KmsKeyVersionName
	}
	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceDiskEncryptionStatus(c *Client, des, nw *InstanceDiskEncryptionStatus) *InstanceDiskEncryptionStatus {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceDiskEncryptionStatusSet(c *Client, des, nw []InstanceDiskEncryptionStatus) []InstanceDiskEncryptionStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceDiskEncryptionStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceDiskEncryptionStatus(c, &d, &n) {
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

type instanceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceApiOperation
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
func diffInstance(c *Client, desired, actual *Instance, opts ...dcl.ApplyOption) ([]instanceDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []instanceDiff
	if !dcl.IsZeroValue(desired.BackendType) && (dcl.IsZeroValue(actual.BackendType) || !reflect.DeepEqual(*desired.BackendType, *actual.BackendType)) {
		c.Config.Logger.Infof("Detected diff in BackendType.\nDESIRED: %v\nACTUAL: %v", desired.BackendType, actual.BackendType)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "BackendType",
		})
	}
	if !dcl.IsZeroValue(desired.ConnectionName) && (dcl.IsZeroValue(actual.ConnectionName) || !reflect.DeepEqual(*desired.ConnectionName, *actual.ConnectionName)) {
		c.Config.Logger.Infof("Detected diff in ConnectionName.\nDESIRED: %v\nACTUAL: %v", desired.ConnectionName, actual.ConnectionName)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "ConnectionName",
		})
	}
	if !dcl.IsZeroValue(desired.DatabaseVersion) && (dcl.IsZeroValue(actual.DatabaseVersion) || !reflect.DeepEqual(*desired.DatabaseVersion, *actual.DatabaseVersion)) {
		c.Config.Logger.Infof("Detected diff in DatabaseVersion.\nDESIRED: %v\nACTUAL: %v", desired.DatabaseVersion, actual.DatabaseVersion)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "DatabaseVersion",
		})
	}
	if !dcl.IsZeroValue(desired.Etag) && (dcl.IsZeroValue(actual.Etag) || !reflect.DeepEqual(*desired.Etag, *actual.Etag)) {
		c.Config.Logger.Infof("Detected diff in Etag.\nDESIRED: %v\nACTUAL: %v", desired.Etag, actual.Etag)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Etag",
		})
	}
	if !dcl.IsZeroValue(desired.GceZone) && (dcl.IsZeroValue(actual.GceZone) || !reflect.DeepEqual(*desired.GceZone, *actual.GceZone)) {
		c.Config.Logger.Infof("Detected diff in GceZone.\nDESIRED: %v\nACTUAL: %v", desired.GceZone, actual.GceZone)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "GceZone",
		})
	}
	if !dcl.IsZeroValue(desired.InstanceType) && (dcl.IsZeroValue(actual.InstanceType) || !reflect.DeepEqual(*desired.InstanceType, *actual.InstanceType)) {
		c.Config.Logger.Infof("Detected diff in InstanceType.\nDESIRED: %v\nACTUAL: %v", desired.InstanceType, actual.InstanceType)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "InstanceType",
		})
	}
	if !dcl.IsZeroValue(desired.MasterInstanceName) && (dcl.IsZeroValue(actual.MasterInstanceName) || !reflect.DeepEqual(*desired.MasterInstanceName, *actual.MasterInstanceName)) {
		c.Config.Logger.Infof("Detected diff in MasterInstanceName.\nDESIRED: %v\nACTUAL: %v", desired.MasterInstanceName, actual.MasterInstanceName)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "MasterInstanceName",
		})
	}
	if compareInstanceMaxDiskSize(c, desired.MaxDiskSize, actual.MaxDiskSize) {
		c.Config.Logger.Infof("Detected diff in MaxDiskSize.\nDESIRED: %v\nACTUAL: %v", desired.MaxDiskSize, actual.MaxDiskSize)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "MaxDiskSize",
		})
	}
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Region) && (dcl.IsZeroValue(actual.Region) || !reflect.DeepEqual(*desired.Region, *actual.Region)) {
		c.Config.Logger.Infof("Detected diff in Region.\nDESIRED: %v\nACTUAL: %v", desired.Region, actual.Region)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Region",
		})
	}
	if !dcl.IsZeroValue(desired.RootPassword) && (dcl.IsZeroValue(actual.RootPassword) || !reflect.DeepEqual(*desired.RootPassword, *actual.RootPassword)) {
		c.Config.Logger.Infof("Detected diff in RootPassword.\nDESIRED: %v\nACTUAL: %v", desired.RootPassword, actual.RootPassword)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "RootPassword",
		})
	}
	if compareInstanceCurrentDiskSize(c, desired.CurrentDiskSize, actual.CurrentDiskSize) {
		c.Config.Logger.Infof("Detected diff in CurrentDiskSize.\nDESIRED: %v\nACTUAL: %v", desired.CurrentDiskSize, actual.CurrentDiskSize)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "CurrentDiskSize",
		})
	}
	if compareInstanceDiskEncryptionConfiguration(c, desired.DiskEncryptionConfiguration, actual.DiskEncryptionConfiguration) {
		c.Config.Logger.Infof("Detected diff in DiskEncryptionConfiguration.\nDESIRED: %v\nACTUAL: %v", desired.DiskEncryptionConfiguration, actual.DiskEncryptionConfiguration)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "DiskEncryptionConfiguration",
		})
	}
	if compareInstanceFailoverReplica(c, desired.FailoverReplica, actual.FailoverReplica) {
		c.Config.Logger.Infof("Detected diff in FailoverReplica.\nDESIRED: %v\nACTUAL: %v", desired.FailoverReplica, actual.FailoverReplica)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "FailoverReplica",
		})
	}
	if compareInstanceIPAddressesSlice(c, desired.IPAddresses, actual.IPAddresses) {
		c.Config.Logger.Infof("Detected diff in IPAddresses.\nDESIRED: %v\nACTUAL: %v", desired.IPAddresses, actual.IPAddresses)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "IPAddresses",
		})
	}
	if compareInstanceMasterInstance(c, desired.MasterInstance, actual.MasterInstance) {
		c.Config.Logger.Infof("Detected diff in MasterInstance.\nDESIRED: %v\nACTUAL: %v", desired.MasterInstance, actual.MasterInstance)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "MasterInstance",
		})
	}
	if compareInstanceReplicaConfiguration(c, desired.ReplicaConfiguration, actual.ReplicaConfiguration) {
		c.Config.Logger.Infof("Detected diff in ReplicaConfiguration.\nDESIRED: %v\nACTUAL: %v", desired.ReplicaConfiguration, actual.ReplicaConfiguration)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "ReplicaConfiguration",
		})
	}
	if compareInstanceScheduledMaintenance(c, desired.ScheduledMaintenance, actual.ScheduledMaintenance) {
		c.Config.Logger.Infof("Detected diff in ScheduledMaintenance.\nDESIRED: %v\nACTUAL: %v", desired.ScheduledMaintenance, actual.ScheduledMaintenance)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "ScheduledMaintenance",
		})
	}
	if compareInstanceSettings(c, desired.Settings, actual.Settings) {
		c.Config.Logger.Infof("Detected diff in Settings.\nDESIRED: %v\nACTUAL: %v", desired.Settings, actual.Settings)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Settings",
		})
	}
	if !dcl.IsZeroValue(desired.State) && (dcl.IsZeroValue(actual.State) || !reflect.DeepEqual(*desired.State, *actual.State)) {
		c.Config.Logger.Infof("Detected diff in State.\nDESIRED: %v\nACTUAL: %v", desired.State, actual.State)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "State",
		})

	}
	if compareInstanceReplicaInstancesSlice(c, desired.ReplicaInstances, actual.ReplicaInstances) {
		c.Config.Logger.Infof("Detected diff in ReplicaInstances.\nDESIRED: %v\nACTUAL: %v", desired.ReplicaInstances, actual.ReplicaInstances)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "ReplicaInstances",
		})

	}
	if compareInstanceServerCaCert(c, desired.ServerCaCert, actual.ServerCaCert) {
		c.Config.Logger.Infof("Detected diff in ServerCaCert.\nDESIRED: %v\nACTUAL: %v", desired.ServerCaCert, actual.ServerCaCert)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "ServerCaCert",
		})

	}
	if !dcl.IsZeroValue(desired.IPv6Address) && (dcl.IsZeroValue(actual.IPv6Address) || !reflect.DeepEqual(*desired.IPv6Address, *actual.IPv6Address)) {
		c.Config.Logger.Infof("Detected diff in IPv6Address.\nDESIRED: %v\nACTUAL: %v", desired.IPv6Address, actual.IPv6Address)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "IPv6Address",
		})

	}
	if !dcl.IsZeroValue(desired.ServiceAccountEmailAddress) && (dcl.IsZeroValue(actual.ServiceAccountEmailAddress) || !reflect.DeepEqual(*desired.ServiceAccountEmailAddress, *actual.ServiceAccountEmailAddress)) {
		c.Config.Logger.Infof("Detected diff in ServiceAccountEmailAddress.\nDESIRED: %v\nACTUAL: %v", desired.ServiceAccountEmailAddress, actual.ServiceAccountEmailAddress)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "ServiceAccountEmailAddress",
		})

	}
	if compareInstanceOnPremisesConfiguration(c, desired.OnPremisesConfiguration, actual.OnPremisesConfiguration) {
		c.Config.Logger.Infof("Detected diff in OnPremisesConfiguration.\nDESIRED: %v\nACTUAL: %v", desired.OnPremisesConfiguration, actual.OnPremisesConfiguration)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "OnPremisesConfiguration",
		})

	}
	if !dcl.SliceEquals(desired.SuspensionReason, actual.SuspensionReason) {
		c.Config.Logger.Infof("Detected diff in SuspensionReason.\nDESIRED: %v\nACTUAL: %v", desired.SuspensionReason, actual.SuspensionReason)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "SuspensionReason",
		})

	}
	if compareInstanceDiskEncryptionStatus(c, desired.DiskEncryptionStatus, actual.DiskEncryptionStatus) {
		c.Config.Logger.Infof("Detected diff in DiskEncryptionStatus.\nDESIRED: %v\nACTUAL: %v", desired.DiskEncryptionStatus, actual.DiskEncryptionStatus)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "DiskEncryptionStatus",
		})

	}
	if !dcl.IsZeroValue(desired.InstanceUid) && (dcl.IsZeroValue(actual.InstanceUid) || !reflect.DeepEqual(*desired.InstanceUid, *actual.InstanceUid)) {
		c.Config.Logger.Infof("Detected diff in InstanceUid.\nDESIRED: %v\nACTUAL: %v", desired.InstanceUid, actual.InstanceUid)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "InstanceUid",
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
	var deduped []instanceDiff
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
func compareInstanceMaxDiskSizeSlice(c *Client, desired, actual []InstanceMaxDiskSize) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaxDiskSize, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceMaxDiskSize(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceMaxDiskSize, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceMaxDiskSize(c *Client, desired, actual *InstanceMaxDiskSize) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}
func compareInstanceCurrentDiskSizeSlice(c *Client, desired, actual []InstanceCurrentDiskSize) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceCurrentDiskSize, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceCurrentDiskSize(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceCurrentDiskSize, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceCurrentDiskSize(c *Client, desired, actual *InstanceCurrentDiskSize) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}
func compareInstanceDiskEncryptionConfigurationSlice(c *Client, desired, actual []InstanceDiskEncryptionConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDiskEncryptionConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDiskEncryptionConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDiskEncryptionConfiguration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDiskEncryptionConfiguration(c *Client, desired, actual *InstanceDiskEncryptionConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.KmsKeyName == nil && desired.KmsKeyName != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyName) {
		c.Config.Logger.Infof("desired KmsKeyName %s - but actually nil", dcl.SprintResource(desired.KmsKeyName))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyName, actual.KmsKeyName) && !dcl.IsZeroValue(desired.KmsKeyName) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyName) && dcl.IsZeroValue(actual.KmsKeyName)) {
		c.Config.Logger.Infof("Diff in KmsKeyName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyName), dcl.SprintResource(actual.KmsKeyName))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}
func compareInstanceFailoverReplicaSlice(c *Client, desired, actual []InstanceFailoverReplica) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceFailoverReplica, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceFailoverReplica(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceFailoverReplica, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceFailoverReplica(c *Client, desired, actual *InstanceFailoverReplica) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Available == nil && desired.Available != nil && !dcl.IsEmptyValueIndirect(desired.Available) {
		c.Config.Logger.Infof("desired Available %s - but actually nil", dcl.SprintResource(desired.Available))
		return true
	}
	if !reflect.DeepEqual(desired.Available, actual.Available) && !dcl.IsZeroValue(desired.Available) && !(dcl.IsEmptyValueIndirect(desired.Available) && dcl.IsZeroValue(actual.Available)) {
		c.Config.Logger.Infof("Diff in Available. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Available), dcl.SprintResource(actual.Available))
		return true
	}
	if actual.FailoverInstance == nil && desired.FailoverInstance != nil && !dcl.IsEmptyValueIndirect(desired.FailoverInstance) {
		c.Config.Logger.Infof("desired FailoverInstance %s - but actually nil", dcl.SprintResource(desired.FailoverInstance))
		return true
	}
	if compareInstanceFailoverReplicaFailoverInstance(c, desired.FailoverInstance, actual.FailoverInstance) && !dcl.IsZeroValue(desired.FailoverInstance) {
		c.Config.Logger.Infof("Diff in FailoverInstance. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FailoverInstance), dcl.SprintResource(actual.FailoverInstance))
		return true
	}
	return false
}
func compareInstanceFailoverReplicaFailoverInstanceSlice(c *Client, desired, actual []InstanceFailoverReplicaFailoverInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceFailoverReplicaFailoverInstance, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceFailoverReplicaFailoverInstance(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceFailoverReplicaFailoverInstance, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceFailoverReplicaFailoverInstance(c *Client, desired, actual *InstanceFailoverReplicaFailoverInstance) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Region == nil && desired.Region != nil && !dcl.IsEmptyValueIndirect(desired.Region) {
		c.Config.Logger.Infof("desired Region %s - but actually nil", dcl.SprintResource(desired.Region))
		return true
	}
	if !reflect.DeepEqual(desired.Region, actual.Region) && !dcl.IsZeroValue(desired.Region) && !(dcl.IsEmptyValueIndirect(desired.Region) && dcl.IsZeroValue(actual.Region)) {
		c.Config.Logger.Infof("Diff in Region. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Region), dcl.SprintResource(actual.Region))
		return true
	}
	return false
}
func compareInstanceIPAddressesSlice(c *Client, desired, actual []InstanceIPAddresses) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceIPAddresses, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceIPAddresses(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceIPAddresses, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceIPAddresses(c *Client, desired, actual *InstanceIPAddresses) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) && !(dcl.IsEmptyValueIndirect(desired.Type) && dcl.IsZeroValue(actual.Type)) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	if actual.IPAddress == nil && desired.IPAddress != nil && !dcl.IsEmptyValueIndirect(desired.IPAddress) {
		c.Config.Logger.Infof("desired IPAddress %s - but actually nil", dcl.SprintResource(desired.IPAddress))
		return true
	}
	if !reflect.DeepEqual(desired.IPAddress, actual.IPAddress) && !dcl.IsZeroValue(desired.IPAddress) && !(dcl.IsEmptyValueIndirect(desired.IPAddress) && dcl.IsZeroValue(actual.IPAddress)) {
		c.Config.Logger.Infof("Diff in IPAddress. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPAddress), dcl.SprintResource(actual.IPAddress))
		return true
	}
	if actual.TimeToRetire == nil && desired.TimeToRetire != nil && !dcl.IsEmptyValueIndirect(desired.TimeToRetire) {
		c.Config.Logger.Infof("desired TimeToRetire %s - but actually nil", dcl.SprintResource(desired.TimeToRetire))
		return true
	}
	if compareInstanceIPAddressesTimeToRetire(c, desired.TimeToRetire, actual.TimeToRetire) && !dcl.IsZeroValue(desired.TimeToRetire) {
		c.Config.Logger.Infof("Diff in TimeToRetire. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TimeToRetire), dcl.SprintResource(actual.TimeToRetire))
		return true
	}
	return false
}
func compareInstanceIPAddressesTimeToRetireSlice(c *Client, desired, actual []InstanceIPAddressesTimeToRetire) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceIPAddressesTimeToRetire, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceIPAddressesTimeToRetire(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceIPAddressesTimeToRetire, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceIPAddressesTimeToRetire(c *Client, desired, actual *InstanceIPAddressesTimeToRetire) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareInstanceMasterInstanceSlice(c *Client, desired, actual []InstanceMasterInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMasterInstance, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceMasterInstance(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceMasterInstance, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceMasterInstance(c *Client, desired, actual *InstanceMasterInstance) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Region == nil && desired.Region != nil && !dcl.IsEmptyValueIndirect(desired.Region) {
		c.Config.Logger.Infof("desired Region %s - but actually nil", dcl.SprintResource(desired.Region))
		return true
	}
	if !reflect.DeepEqual(desired.Region, actual.Region) && !dcl.IsZeroValue(desired.Region) && !(dcl.IsEmptyValueIndirect(desired.Region) && dcl.IsZeroValue(actual.Region)) {
		c.Config.Logger.Infof("Diff in Region. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Region), dcl.SprintResource(actual.Region))
		return true
	}
	return false
}
func compareInstanceReplicaConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfiguration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfiguration(c *Client, desired, actual *InstanceReplicaConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.MysqlReplicaConfiguration == nil && desired.MysqlReplicaConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.MysqlReplicaConfiguration) {
		c.Config.Logger.Infof("desired MysqlReplicaConfiguration %s - but actually nil", dcl.SprintResource(desired.MysqlReplicaConfiguration))
		return true
	}
	if compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c, desired.MysqlReplicaConfiguration, actual.MysqlReplicaConfiguration) && !dcl.IsZeroValue(desired.MysqlReplicaConfiguration) {
		c.Config.Logger.Infof("Diff in MysqlReplicaConfiguration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MysqlReplicaConfiguration), dcl.SprintResource(actual.MysqlReplicaConfiguration))
		return true
	}
	if actual.FailoverTarget == nil && desired.FailoverTarget != nil && !dcl.IsEmptyValueIndirect(desired.FailoverTarget) {
		c.Config.Logger.Infof("desired FailoverTarget %s - but actually nil", dcl.SprintResource(desired.FailoverTarget))
		return true
	}
	if !reflect.DeepEqual(desired.FailoverTarget, actual.FailoverTarget) && !dcl.IsZeroValue(desired.FailoverTarget) && !(dcl.IsEmptyValueIndirect(desired.FailoverTarget) && dcl.IsZeroValue(actual.FailoverTarget)) {
		c.Config.Logger.Infof("Diff in FailoverTarget. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FailoverTarget), dcl.SprintResource(actual.FailoverTarget))
		return true
	}
	if actual.ReplicaPoolConfiguration == nil && desired.ReplicaPoolConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.ReplicaPoolConfiguration) {
		c.Config.Logger.Infof("desired ReplicaPoolConfiguration %s - but actually nil", dcl.SprintResource(desired.ReplicaPoolConfiguration))
		return true
	}
	if compareInstanceReplicaConfigurationReplicaPoolConfiguration(c, desired.ReplicaPoolConfiguration, actual.ReplicaPoolConfiguration) && !dcl.IsZeroValue(desired.ReplicaPoolConfiguration) {
		c.Config.Logger.Infof("Diff in ReplicaPoolConfiguration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicaPoolConfiguration), dcl.SprintResource(actual.ReplicaPoolConfiguration))
		return true
	}
	return false
}
func compareInstanceReplicaConfigurationMysqlReplicaConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfigurationMysqlReplicaConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationMysqlReplicaConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationMysqlReplicaConfiguration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c *Client, desired, actual *InstanceReplicaConfigurationMysqlReplicaConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DumpFilePath == nil && desired.DumpFilePath != nil && !dcl.IsEmptyValueIndirect(desired.DumpFilePath) {
		c.Config.Logger.Infof("desired DumpFilePath %s - but actually nil", dcl.SprintResource(desired.DumpFilePath))
		return true
	}
	if !reflect.DeepEqual(desired.DumpFilePath, actual.DumpFilePath) && !dcl.IsZeroValue(desired.DumpFilePath) && !(dcl.IsEmptyValueIndirect(desired.DumpFilePath) && dcl.IsZeroValue(actual.DumpFilePath)) {
		c.Config.Logger.Infof("Diff in DumpFilePath. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DumpFilePath), dcl.SprintResource(actual.DumpFilePath))
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
	if actual.ConnectRetryInterval == nil && desired.ConnectRetryInterval != nil && !dcl.IsEmptyValueIndirect(desired.ConnectRetryInterval) {
		c.Config.Logger.Infof("desired ConnectRetryInterval %s - but actually nil", dcl.SprintResource(desired.ConnectRetryInterval))
		return true
	}
	if !reflect.DeepEqual(desired.ConnectRetryInterval, actual.ConnectRetryInterval) && !dcl.IsZeroValue(desired.ConnectRetryInterval) && !(dcl.IsEmptyValueIndirect(desired.ConnectRetryInterval) && dcl.IsZeroValue(actual.ConnectRetryInterval)) {
		c.Config.Logger.Infof("Diff in ConnectRetryInterval. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConnectRetryInterval), dcl.SprintResource(actual.ConnectRetryInterval))
		return true
	}
	if actual.MasterHeartbeatPeriod == nil && desired.MasterHeartbeatPeriod != nil && !dcl.IsEmptyValueIndirect(desired.MasterHeartbeatPeriod) {
		c.Config.Logger.Infof("desired MasterHeartbeatPeriod %s - but actually nil", dcl.SprintResource(desired.MasterHeartbeatPeriod))
		return true
	}
	if compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, desired.MasterHeartbeatPeriod, actual.MasterHeartbeatPeriod) && !dcl.IsZeroValue(desired.MasterHeartbeatPeriod) {
		c.Config.Logger.Infof("Diff in MasterHeartbeatPeriod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MasterHeartbeatPeriod), dcl.SprintResource(actual.MasterHeartbeatPeriod))
		return true
	}
	if actual.CaCertificate == nil && desired.CaCertificate != nil && !dcl.IsEmptyValueIndirect(desired.CaCertificate) {
		c.Config.Logger.Infof("desired CaCertificate %s - but actually nil", dcl.SprintResource(desired.CaCertificate))
		return true
	}
	if !reflect.DeepEqual(desired.CaCertificate, actual.CaCertificate) && !dcl.IsZeroValue(desired.CaCertificate) && !(dcl.IsEmptyValueIndirect(desired.CaCertificate) && dcl.IsZeroValue(actual.CaCertificate)) {
		c.Config.Logger.Infof("Diff in CaCertificate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CaCertificate), dcl.SprintResource(actual.CaCertificate))
		return true
	}
	if actual.ClientCertificate == nil && desired.ClientCertificate != nil && !dcl.IsEmptyValueIndirect(desired.ClientCertificate) {
		c.Config.Logger.Infof("desired ClientCertificate %s - but actually nil", dcl.SprintResource(desired.ClientCertificate))
		return true
	}
	if !reflect.DeepEqual(desired.ClientCertificate, actual.ClientCertificate) && !dcl.IsZeroValue(desired.ClientCertificate) && !(dcl.IsEmptyValueIndirect(desired.ClientCertificate) && dcl.IsZeroValue(actual.ClientCertificate)) {
		c.Config.Logger.Infof("Diff in ClientCertificate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientCertificate), dcl.SprintResource(actual.ClientCertificate))
		return true
	}
	if actual.ClientKey == nil && desired.ClientKey != nil && !dcl.IsEmptyValueIndirect(desired.ClientKey) {
		c.Config.Logger.Infof("desired ClientKey %s - but actually nil", dcl.SprintResource(desired.ClientKey))
		return true
	}
	if !reflect.DeepEqual(desired.ClientKey, actual.ClientKey) && !dcl.IsZeroValue(desired.ClientKey) && !(dcl.IsEmptyValueIndirect(desired.ClientKey) && dcl.IsZeroValue(actual.ClientKey)) {
		c.Config.Logger.Infof("Diff in ClientKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientKey), dcl.SprintResource(actual.ClientKey))
		return true
	}
	if actual.SslCipher == nil && desired.SslCipher != nil && !dcl.IsEmptyValueIndirect(desired.SslCipher) {
		c.Config.Logger.Infof("desired SslCipher %s - but actually nil", dcl.SprintResource(desired.SslCipher))
		return true
	}
	if !reflect.DeepEqual(desired.SslCipher, actual.SslCipher) && !dcl.IsZeroValue(desired.SslCipher) && !(dcl.IsEmptyValueIndirect(desired.SslCipher) && dcl.IsZeroValue(actual.SslCipher)) {
		c.Config.Logger.Infof("Diff in SslCipher. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SslCipher), dcl.SprintResource(actual.SslCipher))
		return true
	}
	if actual.VerifyServerCertificate == nil && desired.VerifyServerCertificate != nil && !dcl.IsEmptyValueIndirect(desired.VerifyServerCertificate) {
		c.Config.Logger.Infof("desired VerifyServerCertificate %s - but actually nil", dcl.SprintResource(desired.VerifyServerCertificate))
		return true
	}
	if !reflect.DeepEqual(desired.VerifyServerCertificate, actual.VerifyServerCertificate) && !dcl.IsZeroValue(desired.VerifyServerCertificate) && !(dcl.IsEmptyValueIndirect(desired.VerifyServerCertificate) && dcl.IsZeroValue(actual.VerifyServerCertificate)) {
		c.Config.Logger.Infof("Diff in VerifyServerCertificate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.VerifyServerCertificate), dcl.SprintResource(actual.VerifyServerCertificate))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}
func compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice(c *Client, desired, actual []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c *Client, desired, actual *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}
func compareInstanceReplicaConfigurationReplicaPoolConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfigurationReplicaPoolConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationReplicaPoolConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationReplicaPoolConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfiguration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfiguration(c *Client, desired, actual *InstanceReplicaConfigurationReplicaPoolConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.StaticPoolConfiguration == nil && desired.StaticPoolConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.StaticPoolConfiguration) {
		c.Config.Logger.Infof("desired StaticPoolConfiguration %s - but actually nil", dcl.SprintResource(desired.StaticPoolConfiguration))
		return true
	}
	if compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, desired.StaticPoolConfiguration, actual.StaticPoolConfiguration) && !dcl.IsZeroValue(desired.StaticPoolConfiguration) {
		c.Config.Logger.Infof("Diff in StaticPoolConfiguration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StaticPoolConfiguration), dcl.SprintResource(actual.StaticPoolConfiguration))
		return true
	}
	if actual.AutoscalingPoolConfiguration == nil && desired.AutoscalingPoolConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.AutoscalingPoolConfiguration) {
		c.Config.Logger.Infof("desired AutoscalingPoolConfiguration %s - but actually nil", dcl.SprintResource(desired.AutoscalingPoolConfiguration))
		return true
	}
	if compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, desired.AutoscalingPoolConfiguration, actual.AutoscalingPoolConfiguration) && !dcl.IsZeroValue(desired.AutoscalingPoolConfiguration) {
		c.Config.Logger.Infof("Diff in AutoscalingPoolConfiguration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoscalingPoolConfiguration), dcl.SprintResource(actual.AutoscalingPoolConfiguration))
		return true
	}
	if actual.ReplicaCount == nil && desired.ReplicaCount != nil && !dcl.IsEmptyValueIndirect(desired.ReplicaCount) {
		c.Config.Logger.Infof("desired ReplicaCount %s - but actually nil", dcl.SprintResource(desired.ReplicaCount))
		return true
	}
	if !reflect.DeepEqual(desired.ReplicaCount, actual.ReplicaCount) && !dcl.IsZeroValue(desired.ReplicaCount) && !(dcl.IsEmptyValueIndirect(desired.ReplicaCount) && dcl.IsZeroValue(actual.ReplicaCount)) {
		c.Config.Logger.Infof("Diff in ReplicaCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicaCount), dcl.SprintResource(actual.ReplicaCount))
		return true
	}
	if actual.ExposeReplicaIP == nil && desired.ExposeReplicaIP != nil && !dcl.IsEmptyValueIndirect(desired.ExposeReplicaIP) {
		c.Config.Logger.Infof("desired ExposeReplicaIP %s - but actually nil", dcl.SprintResource(desired.ExposeReplicaIP))
		return true
	}
	if !reflect.DeepEqual(desired.ExposeReplicaIP, actual.ExposeReplicaIP) && !dcl.IsZeroValue(desired.ExposeReplicaIP) && !(dcl.IsEmptyValueIndirect(desired.ExposeReplicaIP) && dcl.IsZeroValue(actual.ExposeReplicaIP)) {
		c.Config.Logger.Infof("Diff in ExposeReplicaIP. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExposeReplicaIP), dcl.SprintResource(actual.ExposeReplicaIP))
		return true
	}
	return false
}
func compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c *Client, desired, actual *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.ReplicaCount == nil && desired.ReplicaCount != nil && !dcl.IsEmptyValueIndirect(desired.ReplicaCount) {
		c.Config.Logger.Infof("desired ReplicaCount %s - but actually nil", dcl.SprintResource(desired.ReplicaCount))
		return true
	}
	if !reflect.DeepEqual(desired.ReplicaCount, actual.ReplicaCount) && !dcl.IsZeroValue(desired.ReplicaCount) && !(dcl.IsEmptyValueIndirect(desired.ReplicaCount) && dcl.IsZeroValue(actual.ReplicaCount)) {
		c.Config.Logger.Infof("Diff in ReplicaCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicaCount), dcl.SprintResource(actual.ReplicaCount))
		return true
	}
	if actual.ExposeReplicaIP == nil && desired.ExposeReplicaIP != nil && !dcl.IsEmptyValueIndirect(desired.ExposeReplicaIP) {
		c.Config.Logger.Infof("desired ExposeReplicaIP %s - but actually nil", dcl.SprintResource(desired.ExposeReplicaIP))
		return true
	}
	if !reflect.DeepEqual(desired.ExposeReplicaIP, actual.ExposeReplicaIP) && !dcl.IsZeroValue(desired.ExposeReplicaIP) && !(dcl.IsEmptyValueIndirect(desired.ExposeReplicaIP) && dcl.IsZeroValue(actual.ExposeReplicaIP)) {
		c.Config.Logger.Infof("Diff in ExposeReplicaIP. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExposeReplicaIP), dcl.SprintResource(actual.ExposeReplicaIP))
		return true
	}
	return false
}
func compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c *Client, desired, actual *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.MinReplicaCount == nil && desired.MinReplicaCount != nil && !dcl.IsEmptyValueIndirect(desired.MinReplicaCount) {
		c.Config.Logger.Infof("desired MinReplicaCount %s - but actually nil", dcl.SprintResource(desired.MinReplicaCount))
		return true
	}
	if !reflect.DeepEqual(desired.MinReplicaCount, actual.MinReplicaCount) && !dcl.IsZeroValue(desired.MinReplicaCount) && !(dcl.IsEmptyValueIndirect(desired.MinReplicaCount) && dcl.IsZeroValue(actual.MinReplicaCount)) {
		c.Config.Logger.Infof("Diff in MinReplicaCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinReplicaCount), dcl.SprintResource(actual.MinReplicaCount))
		return true
	}
	if actual.MaxReplicaCount == nil && desired.MaxReplicaCount != nil && !dcl.IsEmptyValueIndirect(desired.MaxReplicaCount) {
		c.Config.Logger.Infof("desired MaxReplicaCount %s - but actually nil", dcl.SprintResource(desired.MaxReplicaCount))
		return true
	}
	if !reflect.DeepEqual(desired.MaxReplicaCount, actual.MaxReplicaCount) && !dcl.IsZeroValue(desired.MaxReplicaCount) && !(dcl.IsEmptyValueIndirect(desired.MaxReplicaCount) && dcl.IsZeroValue(actual.MaxReplicaCount)) {
		c.Config.Logger.Infof("Diff in MaxReplicaCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxReplicaCount), dcl.SprintResource(actual.MaxReplicaCount))
		return true
	}
	if actual.TargetCpuUtil == nil && desired.TargetCpuUtil != nil && !dcl.IsEmptyValueIndirect(desired.TargetCpuUtil) {
		c.Config.Logger.Infof("desired TargetCpuUtil %s - but actually nil", dcl.SprintResource(desired.TargetCpuUtil))
		return true
	}
	if !reflect.DeepEqual(desired.TargetCpuUtil, actual.TargetCpuUtil) && !dcl.IsZeroValue(desired.TargetCpuUtil) && !(dcl.IsEmptyValueIndirect(desired.TargetCpuUtil) && dcl.IsZeroValue(actual.TargetCpuUtil)) {
		c.Config.Logger.Infof("Diff in TargetCpuUtil. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TargetCpuUtil), dcl.SprintResource(actual.TargetCpuUtil))
		return true
	}
	return false
}
func compareInstanceScheduledMaintenanceSlice(c *Client, desired, actual []InstanceScheduledMaintenance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceScheduledMaintenance, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceScheduledMaintenance(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceScheduledMaintenance, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceScheduledMaintenance(c *Client, desired, actual *InstanceScheduledMaintenance) bool {
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
	if compareInstanceScheduledMaintenanceStartTime(c, desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) {
		c.Config.Logger.Infof("Diff in StartTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if actual.CanDefer == nil && desired.CanDefer != nil && !dcl.IsEmptyValueIndirect(desired.CanDefer) {
		c.Config.Logger.Infof("desired CanDefer %s - but actually nil", dcl.SprintResource(desired.CanDefer))
		return true
	}
	if !reflect.DeepEqual(desired.CanDefer, actual.CanDefer) && !dcl.IsZeroValue(desired.CanDefer) && !(dcl.IsEmptyValueIndirect(desired.CanDefer) && dcl.IsZeroValue(actual.CanDefer)) {
		c.Config.Logger.Infof("Diff in CanDefer. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CanDefer), dcl.SprintResource(actual.CanDefer))
		return true
	}
	if actual.CanReschedule == nil && desired.CanReschedule != nil && !dcl.IsEmptyValueIndirect(desired.CanReschedule) {
		c.Config.Logger.Infof("desired CanReschedule %s - but actually nil", dcl.SprintResource(desired.CanReschedule))
		return true
	}
	if !reflect.DeepEqual(desired.CanReschedule, actual.CanReschedule) && !dcl.IsZeroValue(desired.CanReschedule) && !(dcl.IsEmptyValueIndirect(desired.CanReschedule) && dcl.IsZeroValue(actual.CanReschedule)) {
		c.Config.Logger.Infof("Diff in CanReschedule. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CanReschedule), dcl.SprintResource(actual.CanReschedule))
		return true
	}
	return false
}
func compareInstanceScheduledMaintenanceStartTimeSlice(c *Client, desired, actual []InstanceScheduledMaintenanceStartTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceScheduledMaintenanceStartTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceScheduledMaintenanceStartTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceScheduledMaintenanceStartTime, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceScheduledMaintenanceStartTime(c *Client, desired, actual *InstanceScheduledMaintenanceStartTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareInstanceSettingsSlice(c *Client, desired, actual []InstanceSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettings, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettings(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettings, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettings(c *Client, desired, actual *InstanceSettings) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AuthorizedGaeApplications == nil && desired.AuthorizedGaeApplications != nil && !dcl.IsEmptyValueIndirect(desired.AuthorizedGaeApplications) {
		c.Config.Logger.Infof("desired AuthorizedGaeApplications %s - but actually nil", dcl.SprintResource(desired.AuthorizedGaeApplications))
		return true
	}
	if !dcl.SliceEquals(desired.AuthorizedGaeApplications, actual.AuthorizedGaeApplications) && !dcl.IsZeroValue(desired.AuthorizedGaeApplications) {
		c.Config.Logger.Infof("Diff in AuthorizedGaeApplications. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AuthorizedGaeApplications), dcl.SprintResource(actual.AuthorizedGaeApplications))
		return true
	}
	if actual.Tier == nil && desired.Tier != nil && !dcl.IsEmptyValueIndirect(desired.Tier) {
		c.Config.Logger.Infof("desired Tier %s - but actually nil", dcl.SprintResource(desired.Tier))
		return true
	}
	if !reflect.DeepEqual(desired.Tier, actual.Tier) && !dcl.IsZeroValue(desired.Tier) && !(dcl.IsEmptyValueIndirect(desired.Tier) && dcl.IsZeroValue(actual.Tier)) {
		c.Config.Logger.Infof("Diff in Tier. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tier), dcl.SprintResource(actual.Tier))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.AvailabilityType == nil && desired.AvailabilityType != nil && !dcl.IsEmptyValueIndirect(desired.AvailabilityType) {
		c.Config.Logger.Infof("desired AvailabilityType %s - but actually nil", dcl.SprintResource(desired.AvailabilityType))
		return true
	}
	if !reflect.DeepEqual(desired.AvailabilityType, actual.AvailabilityType) && !dcl.IsZeroValue(desired.AvailabilityType) && !(dcl.IsEmptyValueIndirect(desired.AvailabilityType) && dcl.IsZeroValue(actual.AvailabilityType)) {
		c.Config.Logger.Infof("Diff in AvailabilityType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AvailabilityType), dcl.SprintResource(actual.AvailabilityType))
		return true
	}
	if actual.PricingPlan == nil && desired.PricingPlan != nil && !dcl.IsEmptyValueIndirect(desired.PricingPlan) {
		c.Config.Logger.Infof("desired PricingPlan %s - but actually nil", dcl.SprintResource(desired.PricingPlan))
		return true
	}
	if !reflect.DeepEqual(desired.PricingPlan, actual.PricingPlan) && !dcl.IsZeroValue(desired.PricingPlan) && !(dcl.IsEmptyValueIndirect(desired.PricingPlan) && dcl.IsZeroValue(actual.PricingPlan)) {
		c.Config.Logger.Infof("Diff in PricingPlan. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PricingPlan), dcl.SprintResource(actual.PricingPlan))
		return true
	}
	if actual.ReplicationType == nil && desired.ReplicationType != nil && !dcl.IsEmptyValueIndirect(desired.ReplicationType) {
		c.Config.Logger.Infof("desired ReplicationType %s - but actually nil", dcl.SprintResource(desired.ReplicationType))
		return true
	}
	if !reflect.DeepEqual(desired.ReplicationType, actual.ReplicationType) && !dcl.IsZeroValue(desired.ReplicationType) && !(dcl.IsEmptyValueIndirect(desired.ReplicationType) && dcl.IsZeroValue(actual.ReplicationType)) {
		c.Config.Logger.Infof("Diff in ReplicationType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicationType), dcl.SprintResource(actual.ReplicationType))
		return true
	}
	if actual.ActivationPolicy == nil && desired.ActivationPolicy != nil && !dcl.IsEmptyValueIndirect(desired.ActivationPolicy) {
		c.Config.Logger.Infof("desired ActivationPolicy %s - but actually nil", dcl.SprintResource(desired.ActivationPolicy))
		return true
	}
	if !reflect.DeepEqual(desired.ActivationPolicy, actual.ActivationPolicy) && !dcl.IsZeroValue(desired.ActivationPolicy) && !(dcl.IsEmptyValueIndirect(desired.ActivationPolicy) && dcl.IsZeroValue(actual.ActivationPolicy)) {
		c.Config.Logger.Infof("Diff in ActivationPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ActivationPolicy), dcl.SprintResource(actual.ActivationPolicy))
		return true
	}
	if actual.StorageAutoResize == nil && desired.StorageAutoResize != nil && !dcl.IsEmptyValueIndirect(desired.StorageAutoResize) {
		c.Config.Logger.Infof("desired StorageAutoResize %s - but actually nil", dcl.SprintResource(desired.StorageAutoResize))
		return true
	}
	if !reflect.DeepEqual(desired.StorageAutoResize, actual.StorageAutoResize) && !dcl.IsZeroValue(desired.StorageAutoResize) && !(dcl.IsEmptyValueIndirect(desired.StorageAutoResize) && dcl.IsZeroValue(actual.StorageAutoResize)) {
		c.Config.Logger.Infof("Diff in StorageAutoResize. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StorageAutoResize), dcl.SprintResource(actual.StorageAutoResize))
		return true
	}
	if actual.DataDiskType == nil && desired.DataDiskType != nil && !dcl.IsEmptyValueIndirect(desired.DataDiskType) {
		c.Config.Logger.Infof("desired DataDiskType %s - but actually nil", dcl.SprintResource(desired.DataDiskType))
		return true
	}
	if !reflect.DeepEqual(desired.DataDiskType, actual.DataDiskType) && !dcl.IsZeroValue(desired.DataDiskType) && !(dcl.IsEmptyValueIndirect(desired.DataDiskType) && dcl.IsZeroValue(actual.DataDiskType)) {
		c.Config.Logger.Infof("Diff in DataDiskType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DataDiskType), dcl.SprintResource(actual.DataDiskType))
		return true
	}
	if actual.DatabaseReplicationEnabled == nil && desired.DatabaseReplicationEnabled != nil && !dcl.IsEmptyValueIndirect(desired.DatabaseReplicationEnabled) {
		c.Config.Logger.Infof("desired DatabaseReplicationEnabled %s - but actually nil", dcl.SprintResource(desired.DatabaseReplicationEnabled))
		return true
	}
	if !reflect.DeepEqual(desired.DatabaseReplicationEnabled, actual.DatabaseReplicationEnabled) && !dcl.IsZeroValue(desired.DatabaseReplicationEnabled) && !(dcl.IsEmptyValueIndirect(desired.DatabaseReplicationEnabled) && dcl.IsZeroValue(actual.DatabaseReplicationEnabled)) {
		c.Config.Logger.Infof("Diff in DatabaseReplicationEnabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DatabaseReplicationEnabled), dcl.SprintResource(actual.DatabaseReplicationEnabled))
		return true
	}
	if actual.CrashSafeReplicationEnabled == nil && desired.CrashSafeReplicationEnabled != nil && !dcl.IsEmptyValueIndirect(desired.CrashSafeReplicationEnabled) {
		c.Config.Logger.Infof("desired CrashSafeReplicationEnabled %s - but actually nil", dcl.SprintResource(desired.CrashSafeReplicationEnabled))
		return true
	}
	if !reflect.DeepEqual(desired.CrashSafeReplicationEnabled, actual.CrashSafeReplicationEnabled) && !dcl.IsZeroValue(desired.CrashSafeReplicationEnabled) && !(dcl.IsEmptyValueIndirect(desired.CrashSafeReplicationEnabled) && dcl.IsZeroValue(actual.CrashSafeReplicationEnabled)) {
		c.Config.Logger.Infof("Diff in CrashSafeReplicationEnabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrashSafeReplicationEnabled), dcl.SprintResource(actual.CrashSafeReplicationEnabled))
		return true
	}
	if actual.SettingsVersion == nil && desired.SettingsVersion != nil && !dcl.IsEmptyValueIndirect(desired.SettingsVersion) {
		c.Config.Logger.Infof("desired SettingsVersion %s - but actually nil", dcl.SprintResource(desired.SettingsVersion))
		return true
	}
	if compareInstanceSettingsSettingsVersion(c, desired.SettingsVersion, actual.SettingsVersion) && !dcl.IsZeroValue(desired.SettingsVersion) {
		c.Config.Logger.Infof("Diff in SettingsVersion. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SettingsVersion), dcl.SprintResource(actual.SettingsVersion))
		return true
	}
	if actual.UserLabels == nil && desired.UserLabels != nil && !dcl.IsEmptyValueIndirect(desired.UserLabels) {
		c.Config.Logger.Infof("desired UserLabels %s - but actually nil", dcl.SprintResource(desired.UserLabels))
		return true
	}
	if !reflect.DeepEqual(desired.UserLabels, actual.UserLabels) && !dcl.IsZeroValue(desired.UserLabels) {
		c.Config.Logger.Infof("Diff in UserLabels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UserLabels), dcl.SprintResource(actual.UserLabels))
		return true
	}
	if actual.StorageAutoResizeLimit == nil && desired.StorageAutoResizeLimit != nil && !dcl.IsEmptyValueIndirect(desired.StorageAutoResizeLimit) {
		c.Config.Logger.Infof("desired StorageAutoResizeLimit %s - but actually nil", dcl.SprintResource(desired.StorageAutoResizeLimit))
		return true
	}
	if compareInstanceSettingsStorageAutoResizeLimit(c, desired.StorageAutoResizeLimit, actual.StorageAutoResizeLimit) && !dcl.IsZeroValue(desired.StorageAutoResizeLimit) {
		c.Config.Logger.Infof("Diff in StorageAutoResizeLimit. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StorageAutoResizeLimit), dcl.SprintResource(actual.StorageAutoResizeLimit))
		return true
	}
	if actual.IPConfiguration == nil && desired.IPConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.IPConfiguration) {
		c.Config.Logger.Infof("desired IPConfiguration %s - but actually nil", dcl.SprintResource(desired.IPConfiguration))
		return true
	}
	if compareInstanceSettingsIPConfiguration(c, desired.IPConfiguration, actual.IPConfiguration) && !dcl.IsZeroValue(desired.IPConfiguration) {
		c.Config.Logger.Infof("Diff in IPConfiguration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPConfiguration), dcl.SprintResource(actual.IPConfiguration))
		return true
	}
	if actual.LocationPreference == nil && desired.LocationPreference != nil && !dcl.IsEmptyValueIndirect(desired.LocationPreference) {
		c.Config.Logger.Infof("desired LocationPreference %s - but actually nil", dcl.SprintResource(desired.LocationPreference))
		return true
	}
	if compareInstanceSettingsLocationPreference(c, desired.LocationPreference, actual.LocationPreference) && !dcl.IsZeroValue(desired.LocationPreference) {
		c.Config.Logger.Infof("Diff in LocationPreference. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocationPreference), dcl.SprintResource(actual.LocationPreference))
		return true
	}
	if actual.DatabaseFlags == nil && desired.DatabaseFlags != nil && !dcl.IsEmptyValueIndirect(desired.DatabaseFlags) {
		c.Config.Logger.Infof("desired DatabaseFlags %s - but actually nil", dcl.SprintResource(desired.DatabaseFlags))
		return true
	}
	if compareInstanceSettingsDatabaseFlagsSlice(c, desired.DatabaseFlags, actual.DatabaseFlags) && !dcl.IsZeroValue(desired.DatabaseFlags) {
		c.Config.Logger.Infof("Diff in DatabaseFlags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DatabaseFlags), dcl.SprintResource(actual.DatabaseFlags))
		return true
	}
	if actual.MaintenanceWindow == nil && desired.MaintenanceWindow != nil && !dcl.IsEmptyValueIndirect(desired.MaintenanceWindow) {
		c.Config.Logger.Infof("desired MaintenanceWindow %s - but actually nil", dcl.SprintResource(desired.MaintenanceWindow))
		return true
	}
	if compareInstanceSettingsMaintenanceWindow(c, desired.MaintenanceWindow, actual.MaintenanceWindow) && !dcl.IsZeroValue(desired.MaintenanceWindow) {
		c.Config.Logger.Infof("Diff in MaintenanceWindow. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaintenanceWindow), dcl.SprintResource(actual.MaintenanceWindow))
		return true
	}
	if actual.BackupConfiguration == nil && desired.BackupConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.BackupConfiguration) {
		c.Config.Logger.Infof("desired BackupConfiguration %s - but actually nil", dcl.SprintResource(desired.BackupConfiguration))
		return true
	}
	if compareInstanceSettingsBackupConfiguration(c, desired.BackupConfiguration, actual.BackupConfiguration) && !dcl.IsZeroValue(desired.BackupConfiguration) {
		c.Config.Logger.Infof("Diff in BackupConfiguration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackupConfiguration), dcl.SprintResource(actual.BackupConfiguration))
		return true
	}
	if actual.DataDiskSizeGb == nil && desired.DataDiskSizeGb != nil && !dcl.IsEmptyValueIndirect(desired.DataDiskSizeGb) {
		c.Config.Logger.Infof("desired DataDiskSizeGb %s - but actually nil", dcl.SprintResource(desired.DataDiskSizeGb))
		return true
	}
	if compareInstanceSettingsDataDiskSizeGb(c, desired.DataDiskSizeGb, actual.DataDiskSizeGb) && !dcl.IsZeroValue(desired.DataDiskSizeGb) {
		c.Config.Logger.Infof("Diff in DataDiskSizeGb. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DataDiskSizeGb), dcl.SprintResource(actual.DataDiskSizeGb))
		return true
	}
	if actual.ActiveDirectoryConfig == nil && desired.ActiveDirectoryConfig != nil && !dcl.IsEmptyValueIndirect(desired.ActiveDirectoryConfig) {
		c.Config.Logger.Infof("desired ActiveDirectoryConfig %s - but actually nil", dcl.SprintResource(desired.ActiveDirectoryConfig))
		return true
	}
	if compareInstanceSettingsActiveDirectoryConfig(c, desired.ActiveDirectoryConfig, actual.ActiveDirectoryConfig) && !dcl.IsZeroValue(desired.ActiveDirectoryConfig) {
		c.Config.Logger.Infof("Diff in ActiveDirectoryConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ActiveDirectoryConfig), dcl.SprintResource(actual.ActiveDirectoryConfig))
		return true
	}
	if actual.Collation == nil && desired.Collation != nil && !dcl.IsEmptyValueIndirect(desired.Collation) {
		c.Config.Logger.Infof("desired Collation %s - but actually nil", dcl.SprintResource(desired.Collation))
		return true
	}
	if !reflect.DeepEqual(desired.Collation, actual.Collation) && !dcl.IsZeroValue(desired.Collation) && !(dcl.IsEmptyValueIndirect(desired.Collation) && dcl.IsZeroValue(actual.Collation)) {
		c.Config.Logger.Infof("Diff in Collation. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Collation), dcl.SprintResource(actual.Collation))
		return true
	}
	if actual.DenyMaintenancePeriods == nil && desired.DenyMaintenancePeriods != nil && !dcl.IsEmptyValueIndirect(desired.DenyMaintenancePeriods) {
		c.Config.Logger.Infof("desired DenyMaintenancePeriods %s - but actually nil", dcl.SprintResource(desired.DenyMaintenancePeriods))
		return true
	}
	if compareInstanceSettingsDenyMaintenancePeriodsSlice(c, desired.DenyMaintenancePeriods, actual.DenyMaintenancePeriods) && !dcl.IsZeroValue(desired.DenyMaintenancePeriods) {
		c.Config.Logger.Infof("Diff in DenyMaintenancePeriods. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DenyMaintenancePeriods), dcl.SprintResource(actual.DenyMaintenancePeriods))
		return true
	}
	if actual.InsightsConfig == nil && desired.InsightsConfig != nil && !dcl.IsEmptyValueIndirect(desired.InsightsConfig) {
		c.Config.Logger.Infof("desired InsightsConfig %s - but actually nil", dcl.SprintResource(desired.InsightsConfig))
		return true
	}
	if compareInstanceSettingsInsightsConfig(c, desired.InsightsConfig, actual.InsightsConfig) && !dcl.IsZeroValue(desired.InsightsConfig) {
		c.Config.Logger.Infof("Diff in InsightsConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InsightsConfig), dcl.SprintResource(actual.InsightsConfig))
		return true
	}
	return false
}
func compareInstanceSettingsSettingsVersionSlice(c *Client, desired, actual []InstanceSettingsSettingsVersion) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsSettingsVersion, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsSettingsVersion(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsSettingsVersion, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsSettingsVersion(c *Client, desired, actual *InstanceSettingsSettingsVersion) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}
func compareInstanceSettingsStorageAutoResizeLimitSlice(c *Client, desired, actual []InstanceSettingsStorageAutoResizeLimit) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsStorageAutoResizeLimit, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsStorageAutoResizeLimit(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsStorageAutoResizeLimit, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsStorageAutoResizeLimit(c *Client, desired, actual *InstanceSettingsStorageAutoResizeLimit) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}
func compareInstanceSettingsIPConfigurationSlice(c *Client, desired, actual []InstanceSettingsIPConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsIPConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsIPConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsIPConfiguration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsIPConfiguration(c *Client, desired, actual *InstanceSettingsIPConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.IPv4Enabled == nil && desired.IPv4Enabled != nil && !dcl.IsEmptyValueIndirect(desired.IPv4Enabled) {
		c.Config.Logger.Infof("desired IPv4Enabled %s - but actually nil", dcl.SprintResource(desired.IPv4Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.IPv4Enabled, actual.IPv4Enabled) && !dcl.IsZeroValue(desired.IPv4Enabled) && !(dcl.IsEmptyValueIndirect(desired.IPv4Enabled) && dcl.IsZeroValue(actual.IPv4Enabled)) {
		c.Config.Logger.Infof("Diff in IPv4Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPv4Enabled), dcl.SprintResource(actual.IPv4Enabled))
		return true
	}
	if actual.PrivateNetwork == nil && desired.PrivateNetwork != nil && !dcl.IsEmptyValueIndirect(desired.PrivateNetwork) {
		c.Config.Logger.Infof("desired PrivateNetwork %s - but actually nil", dcl.SprintResource(desired.PrivateNetwork))
		return true
	}
	if !dcl.NameToSelfLink(desired.PrivateNetwork, actual.PrivateNetwork) && !dcl.IsZeroValue(desired.PrivateNetwork) {
		c.Config.Logger.Infof("Diff in PrivateNetwork. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrivateNetwork), dcl.SprintResource(actual.PrivateNetwork))
		return true
	}
	if actual.RequireSsl == nil && desired.RequireSsl != nil && !dcl.IsEmptyValueIndirect(desired.RequireSsl) {
		c.Config.Logger.Infof("desired RequireSsl %s - but actually nil", dcl.SprintResource(desired.RequireSsl))
		return true
	}
	if !reflect.DeepEqual(desired.RequireSsl, actual.RequireSsl) && !dcl.IsZeroValue(desired.RequireSsl) && !(dcl.IsEmptyValueIndirect(desired.RequireSsl) && dcl.IsZeroValue(actual.RequireSsl)) {
		c.Config.Logger.Infof("Diff in RequireSsl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequireSsl), dcl.SprintResource(actual.RequireSsl))
		return true
	}
	if actual.AuthorizedNetworks == nil && desired.AuthorizedNetworks != nil && !dcl.IsEmptyValueIndirect(desired.AuthorizedNetworks) {
		c.Config.Logger.Infof("desired AuthorizedNetworks %s - but actually nil", dcl.SprintResource(desired.AuthorizedNetworks))
		return true
	}
	if compareInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c, desired.AuthorizedNetworks, actual.AuthorizedNetworks) && !dcl.IsZeroValue(desired.AuthorizedNetworks) {
		c.Config.Logger.Infof("Diff in AuthorizedNetworks. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AuthorizedNetworks), dcl.SprintResource(actual.AuthorizedNetworks))
		return true
	}
	return false
}
func compareInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c *Client, desired, actual []InstanceSettingsIPConfigurationAuthorizedNetworks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsIPConfigurationAuthorizedNetworks, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsIPConfigurationAuthorizedNetworks(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsIPConfigurationAuthorizedNetworks, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsIPConfigurationAuthorizedNetworks(c *Client, desired, actual *InstanceSettingsIPConfigurationAuthorizedNetworks) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	if actual.ExpirationTime == nil && desired.ExpirationTime != nil && !dcl.IsEmptyValueIndirect(desired.ExpirationTime) {
		c.Config.Logger.Infof("desired ExpirationTime %s - but actually nil", dcl.SprintResource(desired.ExpirationTime))
		return true
	}
	if !reflect.DeepEqual(desired.ExpirationTime, actual.ExpirationTime) && !dcl.IsZeroValue(desired.ExpirationTime) && !(dcl.IsEmptyValueIndirect(desired.ExpirationTime) && dcl.IsZeroValue(actual.ExpirationTime)) {
		c.Config.Logger.Infof("Diff in ExpirationTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExpirationTime), dcl.SprintResource(actual.ExpirationTime))
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}
func compareInstanceSettingsLocationPreferenceSlice(c *Client, desired, actual []InstanceSettingsLocationPreference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsLocationPreference, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsLocationPreference(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsLocationPreference, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsLocationPreference(c *Client, desired, actual *InstanceSettingsLocationPreference) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Zone == nil && desired.Zone != nil && !dcl.IsEmptyValueIndirect(desired.Zone) {
		c.Config.Logger.Infof("desired Zone %s - but actually nil", dcl.SprintResource(desired.Zone))
		return true
	}
	if !reflect.DeepEqual(desired.Zone, actual.Zone) && !dcl.IsZeroValue(desired.Zone) && !(dcl.IsEmptyValueIndirect(desired.Zone) && dcl.IsZeroValue(actual.Zone)) {
		c.Config.Logger.Infof("Diff in Zone. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Zone), dcl.SprintResource(actual.Zone))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}
func compareInstanceSettingsDatabaseFlagsSlice(c *Client, desired, actual []InstanceSettingsDatabaseFlags) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDatabaseFlags, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsDatabaseFlags(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDatabaseFlags, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDatabaseFlags(c *Client, desired, actual *InstanceSettingsDatabaseFlags) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}
func compareInstanceSettingsMaintenanceWindowSlice(c *Client, desired, actual []InstanceSettingsMaintenanceWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsMaintenanceWindow, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsMaintenanceWindow(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsMaintenanceWindow, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsMaintenanceWindow(c *Client, desired, actual *InstanceSettingsMaintenanceWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Hour == nil && desired.Hour != nil && !dcl.IsEmptyValueIndirect(desired.Hour) {
		c.Config.Logger.Infof("desired Hour %s - but actually nil", dcl.SprintResource(desired.Hour))
		return true
	}
	if !reflect.DeepEqual(desired.Hour, actual.Hour) && !dcl.IsZeroValue(desired.Hour) && !(dcl.IsEmptyValueIndirect(desired.Hour) && dcl.IsZeroValue(actual.Hour)) {
		c.Config.Logger.Infof("Diff in Hour. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Hour), dcl.SprintResource(actual.Hour))
		return true
	}
	if actual.Day == nil && desired.Day != nil && !dcl.IsEmptyValueIndirect(desired.Day) {
		c.Config.Logger.Infof("desired Day %s - but actually nil", dcl.SprintResource(desired.Day))
		return true
	}
	if !reflect.DeepEqual(desired.Day, actual.Day) && !dcl.IsZeroValue(desired.Day) && !(dcl.IsEmptyValueIndirect(desired.Day) && dcl.IsZeroValue(actual.Day)) {
		c.Config.Logger.Infof("Diff in Day. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Day), dcl.SprintResource(actual.Day))
		return true
	}
	if actual.UpdateTrack == nil && desired.UpdateTrack != nil && !dcl.IsEmptyValueIndirect(desired.UpdateTrack) {
		c.Config.Logger.Infof("desired UpdateTrack %s - but actually nil", dcl.SprintResource(desired.UpdateTrack))
		return true
	}
	if !reflect.DeepEqual(desired.UpdateTrack, actual.UpdateTrack) && !dcl.IsZeroValue(desired.UpdateTrack) && !(dcl.IsEmptyValueIndirect(desired.UpdateTrack) && dcl.IsZeroValue(actual.UpdateTrack)) {
		c.Config.Logger.Infof("Diff in UpdateTrack. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpdateTrack), dcl.SprintResource(actual.UpdateTrack))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}
func compareInstanceSettingsBackupConfigurationSlice(c *Client, desired, actual []InstanceSettingsBackupConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsBackupConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsBackupConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfiguration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsBackupConfiguration(c *Client, desired, actual *InstanceSettingsBackupConfiguration) bool {
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
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !reflect.DeepEqual(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) && !(dcl.IsEmptyValueIndirect(desired.Enabled) && dcl.IsZeroValue(actual.Enabled)) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.BinaryLogEnabled == nil && desired.BinaryLogEnabled != nil && !dcl.IsEmptyValueIndirect(desired.BinaryLogEnabled) {
		c.Config.Logger.Infof("desired BinaryLogEnabled %s - but actually nil", dcl.SprintResource(desired.BinaryLogEnabled))
		return true
	}
	if !reflect.DeepEqual(desired.BinaryLogEnabled, actual.BinaryLogEnabled) && !dcl.IsZeroValue(desired.BinaryLogEnabled) && !(dcl.IsEmptyValueIndirect(desired.BinaryLogEnabled) && dcl.IsZeroValue(actual.BinaryLogEnabled)) {
		c.Config.Logger.Infof("Diff in BinaryLogEnabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BinaryLogEnabled), dcl.SprintResource(actual.BinaryLogEnabled))
		return true
	}
	if actual.Location == nil && desired.Location != nil && !dcl.IsEmptyValueIndirect(desired.Location) {
		c.Config.Logger.Infof("desired Location %s - but actually nil", dcl.SprintResource(desired.Location))
		return true
	}
	if !reflect.DeepEqual(desired.Location, actual.Location) && !dcl.IsZeroValue(desired.Location) && !(dcl.IsEmptyValueIndirect(desired.Location) && dcl.IsZeroValue(actual.Location)) {
		c.Config.Logger.Infof("Diff in Location. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Location), dcl.SprintResource(actual.Location))
		return true
	}
	if actual.BackupRetentionSettings == nil && desired.BackupRetentionSettings != nil && !dcl.IsEmptyValueIndirect(desired.BackupRetentionSettings) {
		c.Config.Logger.Infof("desired BackupRetentionSettings %s - but actually nil", dcl.SprintResource(desired.BackupRetentionSettings))
		return true
	}
	if compareInstanceSettingsBackupConfigurationBackupRetentionSettings(c, desired.BackupRetentionSettings, actual.BackupRetentionSettings) && !dcl.IsZeroValue(desired.BackupRetentionSettings) {
		c.Config.Logger.Infof("Diff in BackupRetentionSettings. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackupRetentionSettings), dcl.SprintResource(actual.BackupRetentionSettings))
		return true
	}
	if actual.TransactionLogRetentionDays == nil && desired.TransactionLogRetentionDays != nil && !dcl.IsEmptyValueIndirect(desired.TransactionLogRetentionDays) {
		c.Config.Logger.Infof("desired TransactionLogRetentionDays %s - but actually nil", dcl.SprintResource(desired.TransactionLogRetentionDays))
		return true
	}
	if !reflect.DeepEqual(desired.TransactionLogRetentionDays, actual.TransactionLogRetentionDays) && !dcl.IsZeroValue(desired.TransactionLogRetentionDays) && !(dcl.IsEmptyValueIndirect(desired.TransactionLogRetentionDays) && dcl.IsZeroValue(actual.TransactionLogRetentionDays)) {
		c.Config.Logger.Infof("Diff in TransactionLogRetentionDays. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TransactionLogRetentionDays), dcl.SprintResource(actual.TransactionLogRetentionDays))
		return true
	}
	return false
}
func compareInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice(c *Client, desired, actual []InstanceSettingsBackupConfigurationBackupRetentionSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettings, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettings, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsBackupConfigurationBackupRetentionSettings(c *Client, desired, actual *InstanceSettingsBackupConfigurationBackupRetentionSettings) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RetentionUnit == nil && desired.RetentionUnit != nil && !dcl.IsEmptyValueIndirect(desired.RetentionUnit) {
		c.Config.Logger.Infof("desired RetentionUnit %s - but actually nil", dcl.SprintResource(desired.RetentionUnit))
		return true
	}
	if !reflect.DeepEqual(desired.RetentionUnit, actual.RetentionUnit) && !dcl.IsZeroValue(desired.RetentionUnit) && !(dcl.IsEmptyValueIndirect(desired.RetentionUnit) && dcl.IsZeroValue(actual.RetentionUnit)) {
		c.Config.Logger.Infof("Diff in RetentionUnit. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetentionUnit), dcl.SprintResource(actual.RetentionUnit))
		return true
	}
	if actual.RetainedBackups == nil && desired.RetainedBackups != nil && !dcl.IsEmptyValueIndirect(desired.RetainedBackups) {
		c.Config.Logger.Infof("desired RetainedBackups %s - but actually nil", dcl.SprintResource(desired.RetainedBackups))
		return true
	}
	if !reflect.DeepEqual(desired.RetainedBackups, actual.RetainedBackups) && !dcl.IsZeroValue(desired.RetainedBackups) && !(dcl.IsEmptyValueIndirect(desired.RetainedBackups) && dcl.IsZeroValue(actual.RetainedBackups)) {
		c.Config.Logger.Infof("Diff in RetainedBackups. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetainedBackups), dcl.SprintResource(actual.RetainedBackups))
		return true
	}
	return false
}
func compareInstanceSettingsDataDiskSizeGbSlice(c *Client, desired, actual []InstanceSettingsDataDiskSizeGb) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDataDiskSizeGb, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsDataDiskSizeGb(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDataDiskSizeGb, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDataDiskSizeGb(c *Client, desired, actual *InstanceSettingsDataDiskSizeGb) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) && !(dcl.IsEmptyValueIndirect(desired.Value) && dcl.IsZeroValue(actual.Value)) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}
func compareInstanceSettingsActiveDirectoryConfigSlice(c *Client, desired, actual []InstanceSettingsActiveDirectoryConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsActiveDirectoryConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsActiveDirectoryConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsActiveDirectoryConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsActiveDirectoryConfig(c *Client, desired, actual *InstanceSettingsActiveDirectoryConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.Domain == nil && desired.Domain != nil && !dcl.IsEmptyValueIndirect(desired.Domain) {
		c.Config.Logger.Infof("desired Domain %s - but actually nil", dcl.SprintResource(desired.Domain))
		return true
	}
	if !reflect.DeepEqual(desired.Domain, actual.Domain) && !dcl.IsZeroValue(desired.Domain) && !(dcl.IsEmptyValueIndirect(desired.Domain) && dcl.IsZeroValue(actual.Domain)) {
		c.Config.Logger.Infof("Diff in Domain. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Domain), dcl.SprintResource(actual.Domain))
		return true
	}
	return false
}
func compareInstanceSettingsDenyMaintenancePeriodsSlice(c *Client, desired, actual []InstanceSettingsDenyMaintenancePeriods) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDenyMaintenancePeriods, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsDenyMaintenancePeriods(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDenyMaintenancePeriods, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDenyMaintenancePeriods(c *Client, desired, actual *InstanceSettingsDenyMaintenancePeriods) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.StartDate == nil && desired.StartDate != nil && !dcl.IsEmptyValueIndirect(desired.StartDate) {
		c.Config.Logger.Infof("desired StartDate %s - but actually nil", dcl.SprintResource(desired.StartDate))
		return true
	}
	if !reflect.DeepEqual(desired.StartDate, actual.StartDate) && !dcl.IsZeroValue(desired.StartDate) && !(dcl.IsEmptyValueIndirect(desired.StartDate) && dcl.IsZeroValue(actual.StartDate)) {
		c.Config.Logger.Infof("Diff in StartDate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartDate), dcl.SprintResource(actual.StartDate))
		return true
	}
	if actual.EndDate == nil && desired.EndDate != nil && !dcl.IsEmptyValueIndirect(desired.EndDate) {
		c.Config.Logger.Infof("desired EndDate %s - but actually nil", dcl.SprintResource(desired.EndDate))
		return true
	}
	if !reflect.DeepEqual(desired.EndDate, actual.EndDate) && !dcl.IsZeroValue(desired.EndDate) && !(dcl.IsEmptyValueIndirect(desired.EndDate) && dcl.IsZeroValue(actual.EndDate)) {
		c.Config.Logger.Infof("Diff in EndDate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EndDate), dcl.SprintResource(actual.EndDate))
		return true
	}
	if actual.Time == nil && desired.Time != nil && !dcl.IsEmptyValueIndirect(desired.Time) {
		c.Config.Logger.Infof("desired Time %s - but actually nil", dcl.SprintResource(desired.Time))
		return true
	}
	if !reflect.DeepEqual(desired.Time, actual.Time) && !dcl.IsZeroValue(desired.Time) && !(dcl.IsEmptyValueIndirect(desired.Time) && dcl.IsZeroValue(actual.Time)) {
		c.Config.Logger.Infof("Diff in Time. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Time), dcl.SprintResource(actual.Time))
		return true
	}
	return false
}
func compareInstanceSettingsInsightsConfigSlice(c *Client, desired, actual []InstanceSettingsInsightsConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsInsightsConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsInsightsConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsInsightsConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsInsightsConfig(c *Client, desired, actual *InstanceSettingsInsightsConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.QueryInsightsEnabled == nil && desired.QueryInsightsEnabled != nil && !dcl.IsEmptyValueIndirect(desired.QueryInsightsEnabled) {
		c.Config.Logger.Infof("desired QueryInsightsEnabled %s - but actually nil", dcl.SprintResource(desired.QueryInsightsEnabled))
		return true
	}
	if !reflect.DeepEqual(desired.QueryInsightsEnabled, actual.QueryInsightsEnabled) && !dcl.IsZeroValue(desired.QueryInsightsEnabled) && !(dcl.IsEmptyValueIndirect(desired.QueryInsightsEnabled) && dcl.IsZeroValue(actual.QueryInsightsEnabled)) {
		c.Config.Logger.Infof("Diff in QueryInsightsEnabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryInsightsEnabled), dcl.SprintResource(actual.QueryInsightsEnabled))
		return true
	}
	if actual.RecordClientAddress == nil && desired.RecordClientAddress != nil && !dcl.IsEmptyValueIndirect(desired.RecordClientAddress) {
		c.Config.Logger.Infof("desired RecordClientAddress %s - but actually nil", dcl.SprintResource(desired.RecordClientAddress))
		return true
	}
	if !reflect.DeepEqual(desired.RecordClientAddress, actual.RecordClientAddress) && !dcl.IsZeroValue(desired.RecordClientAddress) && !(dcl.IsEmptyValueIndirect(desired.RecordClientAddress) && dcl.IsZeroValue(actual.RecordClientAddress)) {
		c.Config.Logger.Infof("Diff in RecordClientAddress. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RecordClientAddress), dcl.SprintResource(actual.RecordClientAddress))
		return true
	}
	if actual.RecordApplicationTags == nil && desired.RecordApplicationTags != nil && !dcl.IsEmptyValueIndirect(desired.RecordApplicationTags) {
		c.Config.Logger.Infof("desired RecordApplicationTags %s - but actually nil", dcl.SprintResource(desired.RecordApplicationTags))
		return true
	}
	if !reflect.DeepEqual(desired.RecordApplicationTags, actual.RecordApplicationTags) && !dcl.IsZeroValue(desired.RecordApplicationTags) && !(dcl.IsEmptyValueIndirect(desired.RecordApplicationTags) && dcl.IsZeroValue(actual.RecordApplicationTags)) {
		c.Config.Logger.Infof("Diff in RecordApplicationTags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RecordApplicationTags), dcl.SprintResource(actual.RecordApplicationTags))
		return true
	}
	if actual.QueryStringLength == nil && desired.QueryStringLength != nil && !dcl.IsEmptyValueIndirect(desired.QueryStringLength) {
		c.Config.Logger.Infof("desired QueryStringLength %s - but actually nil", dcl.SprintResource(desired.QueryStringLength))
		return true
	}
	if !reflect.DeepEqual(desired.QueryStringLength, actual.QueryStringLength) && !dcl.IsZeroValue(desired.QueryStringLength) && !(dcl.IsEmptyValueIndirect(desired.QueryStringLength) && dcl.IsZeroValue(actual.QueryStringLength)) {
		c.Config.Logger.Infof("Diff in QueryStringLength. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryStringLength), dcl.SprintResource(actual.QueryStringLength))
		return true
	}
	return false
}
func compareInstanceReplicaInstancesSlice(c *Client, desired, actual []InstanceReplicaInstances) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaInstances, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaInstances(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaInstances, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaInstances(c *Client, desired, actual *InstanceReplicaInstances) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) && !(dcl.IsEmptyValueIndirect(desired.Name) && dcl.IsZeroValue(actual.Name)) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Region == nil && desired.Region != nil && !dcl.IsEmptyValueIndirect(desired.Region) {
		c.Config.Logger.Infof("desired Region %s - but actually nil", dcl.SprintResource(desired.Region))
		return true
	}
	if !reflect.DeepEqual(desired.Region, actual.Region) && !dcl.IsZeroValue(desired.Region) && !(dcl.IsEmptyValueIndirect(desired.Region) && dcl.IsZeroValue(actual.Region)) {
		c.Config.Logger.Infof("Diff in Region. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Region), dcl.SprintResource(actual.Region))
		return true
	}
	return false
}
func compareInstanceServerCaCertSlice(c *Client, desired, actual []InstanceServerCaCert) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceServerCaCert, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceServerCaCert(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceServerCaCert, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceServerCaCert(c *Client, desired, actual *InstanceServerCaCert) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.CertSerialNumber == nil && desired.CertSerialNumber != nil && !dcl.IsEmptyValueIndirect(desired.CertSerialNumber) {
		c.Config.Logger.Infof("desired CertSerialNumber %s - but actually nil", dcl.SprintResource(desired.CertSerialNumber))
		return true
	}
	if !reflect.DeepEqual(desired.CertSerialNumber, actual.CertSerialNumber) && !dcl.IsZeroValue(desired.CertSerialNumber) && !(dcl.IsEmptyValueIndirect(desired.CertSerialNumber) && dcl.IsZeroValue(actual.CertSerialNumber)) {
		c.Config.Logger.Infof("Diff in CertSerialNumber. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CertSerialNumber), dcl.SprintResource(actual.CertSerialNumber))
		return true
	}
	if actual.Cert == nil && desired.Cert != nil && !dcl.IsEmptyValueIndirect(desired.Cert) {
		c.Config.Logger.Infof("desired Cert %s - but actually nil", dcl.SprintResource(desired.Cert))
		return true
	}
	if !reflect.DeepEqual(desired.Cert, actual.Cert) && !dcl.IsZeroValue(desired.Cert) && !(dcl.IsEmptyValueIndirect(desired.Cert) && dcl.IsZeroValue(actual.Cert)) {
		c.Config.Logger.Infof("Diff in Cert. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Cert), dcl.SprintResource(actual.Cert))
		return true
	}
	if actual.CreateTime == nil && desired.CreateTime != nil && !dcl.IsEmptyValueIndirect(desired.CreateTime) {
		c.Config.Logger.Infof("desired CreateTime %s - but actually nil", dcl.SprintResource(desired.CreateTime))
		return true
	}
	if !reflect.DeepEqual(desired.CreateTime, actual.CreateTime) && !dcl.IsZeroValue(desired.CreateTime) && !(dcl.IsEmptyValueIndirect(desired.CreateTime) && dcl.IsZeroValue(actual.CreateTime)) {
		c.Config.Logger.Infof("Diff in CreateTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CreateTime), dcl.SprintResource(actual.CreateTime))
		return true
	}
	if actual.CommonName == nil && desired.CommonName != nil && !dcl.IsEmptyValueIndirect(desired.CommonName) {
		c.Config.Logger.Infof("desired CommonName %s - but actually nil", dcl.SprintResource(desired.CommonName))
		return true
	}
	if !reflect.DeepEqual(desired.CommonName, actual.CommonName) && !dcl.IsZeroValue(desired.CommonName) && !(dcl.IsEmptyValueIndirect(desired.CommonName) && dcl.IsZeroValue(actual.CommonName)) {
		c.Config.Logger.Infof("Diff in CommonName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CommonName), dcl.SprintResource(actual.CommonName))
		return true
	}
	if actual.ExpirationTime == nil && desired.ExpirationTime != nil && !dcl.IsEmptyValueIndirect(desired.ExpirationTime) {
		c.Config.Logger.Infof("desired ExpirationTime %s - but actually nil", dcl.SprintResource(desired.ExpirationTime))
		return true
	}
	if !reflect.DeepEqual(desired.ExpirationTime, actual.ExpirationTime) && !dcl.IsZeroValue(desired.ExpirationTime) && !(dcl.IsEmptyValueIndirect(desired.ExpirationTime) && dcl.IsZeroValue(actual.ExpirationTime)) {
		c.Config.Logger.Infof("Diff in ExpirationTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExpirationTime), dcl.SprintResource(actual.ExpirationTime))
		return true
	}
	if actual.Sha1Fingerprint == nil && desired.Sha1Fingerprint != nil && !dcl.IsEmptyValueIndirect(desired.Sha1Fingerprint) {
		c.Config.Logger.Infof("desired Sha1Fingerprint %s - but actually nil", dcl.SprintResource(desired.Sha1Fingerprint))
		return true
	}
	if !reflect.DeepEqual(desired.Sha1Fingerprint, actual.Sha1Fingerprint) && !dcl.IsZeroValue(desired.Sha1Fingerprint) && !(dcl.IsEmptyValueIndirect(desired.Sha1Fingerprint) && dcl.IsZeroValue(actual.Sha1Fingerprint)) {
		c.Config.Logger.Infof("Diff in Sha1Fingerprint. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sha1Fingerprint), dcl.SprintResource(actual.Sha1Fingerprint))
		return true
	}
	if actual.Instance == nil && desired.Instance != nil && !dcl.IsEmptyValueIndirect(desired.Instance) {
		c.Config.Logger.Infof("desired Instance %s - but actually nil", dcl.SprintResource(desired.Instance))
		return true
	}
	if !reflect.DeepEqual(desired.Instance, actual.Instance) && !dcl.IsZeroValue(desired.Instance) && !(dcl.IsEmptyValueIndirect(desired.Instance) && dcl.IsZeroValue(actual.Instance)) {
		c.Config.Logger.Infof("Diff in Instance. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Instance), dcl.SprintResource(actual.Instance))
		return true
	}
	return false
}
func compareInstanceOnPremisesConfigurationSlice(c *Client, desired, actual []InstanceOnPremisesConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceOnPremisesConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceOnPremisesConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceOnPremisesConfiguration, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceOnPremisesConfiguration(c *Client, desired, actual *InstanceOnPremisesConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HostPort == nil && desired.HostPort != nil && !dcl.IsEmptyValueIndirect(desired.HostPort) {
		c.Config.Logger.Infof("desired HostPort %s - but actually nil", dcl.SprintResource(desired.HostPort))
		return true
	}
	if !reflect.DeepEqual(desired.HostPort, actual.HostPort) && !dcl.IsZeroValue(desired.HostPort) && !(dcl.IsEmptyValueIndirect(desired.HostPort) && dcl.IsZeroValue(actual.HostPort)) {
		c.Config.Logger.Infof("Diff in HostPort. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HostPort), dcl.SprintResource(actual.HostPort))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
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
	if actual.CaCertificate == nil && desired.CaCertificate != nil && !dcl.IsEmptyValueIndirect(desired.CaCertificate) {
		c.Config.Logger.Infof("desired CaCertificate %s - but actually nil", dcl.SprintResource(desired.CaCertificate))
		return true
	}
	if !reflect.DeepEqual(desired.CaCertificate, actual.CaCertificate) && !dcl.IsZeroValue(desired.CaCertificate) && !(dcl.IsEmptyValueIndirect(desired.CaCertificate) && dcl.IsZeroValue(actual.CaCertificate)) {
		c.Config.Logger.Infof("Diff in CaCertificate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CaCertificate), dcl.SprintResource(actual.CaCertificate))
		return true
	}
	if actual.ClientCertificate == nil && desired.ClientCertificate != nil && !dcl.IsEmptyValueIndirect(desired.ClientCertificate) {
		c.Config.Logger.Infof("desired ClientCertificate %s - but actually nil", dcl.SprintResource(desired.ClientCertificate))
		return true
	}
	if !reflect.DeepEqual(desired.ClientCertificate, actual.ClientCertificate) && !dcl.IsZeroValue(desired.ClientCertificate) && !(dcl.IsEmptyValueIndirect(desired.ClientCertificate) && dcl.IsZeroValue(actual.ClientCertificate)) {
		c.Config.Logger.Infof("Diff in ClientCertificate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientCertificate), dcl.SprintResource(actual.ClientCertificate))
		return true
	}
	if actual.ClientKey == nil && desired.ClientKey != nil && !dcl.IsEmptyValueIndirect(desired.ClientKey) {
		c.Config.Logger.Infof("desired ClientKey %s - but actually nil", dcl.SprintResource(desired.ClientKey))
		return true
	}
	if !reflect.DeepEqual(desired.ClientKey, actual.ClientKey) && !dcl.IsZeroValue(desired.ClientKey) && !(dcl.IsEmptyValueIndirect(desired.ClientKey) && dcl.IsZeroValue(actual.ClientKey)) {
		c.Config.Logger.Infof("Diff in ClientKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientKey), dcl.SprintResource(actual.ClientKey))
		return true
	}
	if actual.DumpFilePath == nil && desired.DumpFilePath != nil && !dcl.IsEmptyValueIndirect(desired.DumpFilePath) {
		c.Config.Logger.Infof("desired DumpFilePath %s - but actually nil", dcl.SprintResource(desired.DumpFilePath))
		return true
	}
	if !reflect.DeepEqual(desired.DumpFilePath, actual.DumpFilePath) && !dcl.IsZeroValue(desired.DumpFilePath) && !(dcl.IsEmptyValueIndirect(desired.DumpFilePath) && dcl.IsZeroValue(actual.DumpFilePath)) {
		c.Config.Logger.Infof("Diff in DumpFilePath. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DumpFilePath), dcl.SprintResource(actual.DumpFilePath))
		return true
	}
	if actual.Database == nil && desired.Database != nil && !dcl.IsEmptyValueIndirect(desired.Database) {
		c.Config.Logger.Infof("desired Database %s - but actually nil", dcl.SprintResource(desired.Database))
		return true
	}
	if !reflect.DeepEqual(desired.Database, actual.Database) && !dcl.IsZeroValue(desired.Database) && !(dcl.IsEmptyValueIndirect(desired.Database) && dcl.IsZeroValue(actual.Database)) {
		c.Config.Logger.Infof("Diff in Database. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Database), dcl.SprintResource(actual.Database))
		return true
	}
	if actual.ReplicatedDatabases == nil && desired.ReplicatedDatabases != nil && !dcl.IsEmptyValueIndirect(desired.ReplicatedDatabases) {
		c.Config.Logger.Infof("desired ReplicatedDatabases %s - but actually nil", dcl.SprintResource(desired.ReplicatedDatabases))
		return true
	}
	if !dcl.SliceEquals(desired.ReplicatedDatabases, actual.ReplicatedDatabases) && !dcl.IsZeroValue(desired.ReplicatedDatabases) {
		c.Config.Logger.Infof("Diff in ReplicatedDatabases. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicatedDatabases), dcl.SprintResource(actual.ReplicatedDatabases))
		return true
	}
	return false
}
func compareInstanceDiskEncryptionStatusSlice(c *Client, desired, actual []InstanceDiskEncryptionStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDiskEncryptionStatus, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDiskEncryptionStatus(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDiskEncryptionStatus, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDiskEncryptionStatus(c *Client, desired, actual *InstanceDiskEncryptionStatus) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.KmsKeyVersionName == nil && desired.KmsKeyVersionName != nil && !dcl.IsEmptyValueIndirect(desired.KmsKeyVersionName) {
		c.Config.Logger.Infof("desired KmsKeyVersionName %s - but actually nil", dcl.SprintResource(desired.KmsKeyVersionName))
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyVersionName, actual.KmsKeyVersionName) && !dcl.IsZeroValue(desired.KmsKeyVersionName) && !(dcl.IsEmptyValueIndirect(desired.KmsKeyVersionName) && dcl.IsZeroValue(actual.KmsKeyVersionName)) {
		c.Config.Logger.Infof("Diff in KmsKeyVersionName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyVersionName), dcl.SprintResource(actual.KmsKeyVersionName))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		c.Config.Logger.Infof("desired Kind %s - but actually nil", dcl.SprintResource(desired.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) && !(dcl.IsEmptyValueIndirect(desired.Kind) && dcl.IsZeroValue(actual.Kind)) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}
func compareInstanceBackendTypeEnumSlice(c *Client, desired, actual []InstanceBackendTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceBackendTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceBackendTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceBackendTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceBackendTypeEnum(c *Client, desired, actual *InstanceBackendTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceDatabaseVersionEnumSlice(c *Client, desired, actual []InstanceDatabaseVersionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDatabaseVersionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDatabaseVersionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDatabaseVersionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDatabaseVersionEnum(c *Client, desired, actual *InstanceDatabaseVersionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceInstanceTypeEnumSlice(c *Client, desired, actual []InstanceInstanceTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceInstanceTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceInstanceTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceInstanceTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceInstanceTypeEnum(c *Client, desired, actual *InstanceInstanceTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceIPAddressesTypeEnumSlice(c *Client, desired, actual []InstanceIPAddressesTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceIPAddressesTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceIPAddressesTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceIPAddressesTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceIPAddressesTypeEnum(c *Client, desired, actual *InstanceIPAddressesTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsAvailabilityTypeEnumSlice(c *Client, desired, actual []InstanceSettingsAvailabilityTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsAvailabilityTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsAvailabilityTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsAvailabilityTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsAvailabilityTypeEnum(c *Client, desired, actual *InstanceSettingsAvailabilityTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsPricingPlanEnumSlice(c *Client, desired, actual []InstanceSettingsPricingPlanEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsPricingPlanEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsPricingPlanEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsPricingPlanEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsPricingPlanEnum(c *Client, desired, actual *InstanceSettingsPricingPlanEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsReplicationTypeEnumSlice(c *Client, desired, actual []InstanceSettingsReplicationTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsReplicationTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsReplicationTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsReplicationTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsReplicationTypeEnum(c *Client, desired, actual *InstanceSettingsReplicationTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsActivationPolicyEnumSlice(c *Client, desired, actual []InstanceSettingsActivationPolicyEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsActivationPolicyEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsActivationPolicyEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsActivationPolicyEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsActivationPolicyEnum(c *Client, desired, actual *InstanceSettingsActivationPolicyEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsDataDiskTypeEnumSlice(c *Client, desired, actual []InstanceSettingsDataDiskTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDataDiskTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsDataDiskTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDataDiskTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDataDiskTypeEnum(c *Client, desired, actual *InstanceSettingsDataDiskTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsMaintenanceWindowUpdateTrackEnumSlice(c *Client, desired, actual []InstanceSettingsMaintenanceWindowUpdateTrackEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsMaintenanceWindowUpdateTrackEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsMaintenanceWindowUpdateTrackEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsMaintenanceWindowUpdateTrackEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsMaintenanceWindowUpdateTrackEnum(c *Client, desired, actual *InstanceSettingsMaintenanceWindowUpdateTrackEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumSlice(c *Client, desired, actual []InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(c *Client, desired, actual *InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Instance) urlNormalized() *Instance {
	normalized := deepcopy.Copy(*r).(Instance)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Instance) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Instance) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/instances/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Instance resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Instance) marshal(c *Client) ([]byte, error) {
	m, err := expandInstance(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Instance: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalInstance decodes JSON responses into the Instance resource schema.
func unmarshalInstance(b []byte, c *Client) (*Instance, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return flattenInstance(c, m), nil
}

// expandInstance expands Instance into a JSON request object.
func expandInstance(c *Client, f *Instance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.BackendType; !dcl.IsEmptyValueIndirect(v) {
		m["backendType"] = v
	}
	if v := f.ConnectionName; !dcl.IsEmptyValueIndirect(v) {
		m["connectionName"] = v
	}
	if v := f.DatabaseVersion; !dcl.IsEmptyValueIndirect(v) {
		m["databaseVersion"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v := f.GceZone; !dcl.IsEmptyValueIndirect(v) {
		m["gceZone"] = v
	}
	if v := f.InstanceType; !dcl.IsEmptyValueIndirect(v) {
		m["instanceType"] = v
	}
	if v := f.MasterInstanceName; !dcl.IsEmptyValueIndirect(v) {
		m["masterInstanceName"] = v
	}
	if v, err := expandInstanceMaxDiskSize(c, f.MaxDiskSize); err != nil {
		return nil, fmt.Errorf("error expanding MaxDiskSize into maxDiskSize: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maxDiskSize"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v := f.RootPassword; !dcl.IsEmptyValueIndirect(v) {
		m["rootPassword"] = v
	}
	if v, err := expandInstanceCurrentDiskSize(c, f.CurrentDiskSize); err != nil {
		return nil, fmt.Errorf("error expanding CurrentDiskSize into currentDiskSize: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["currentDiskSize"] = v
	}
	if v, err := expandInstanceDiskEncryptionConfiguration(c, f.DiskEncryptionConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding DiskEncryptionConfiguration into diskEncryptionConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["diskEncryptionConfiguration"] = v
	}
	if v, err := expandInstanceFailoverReplica(c, f.FailoverReplica); err != nil {
		return nil, fmt.Errorf("error expanding FailoverReplica into failoverReplica: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["failoverReplica"] = v
	}
	if v, err := expandInstanceIPAddressesSlice(c, f.IPAddresses); err != nil {
		return nil, fmt.Errorf("error expanding IPAddresses into ipAddresses: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ipAddresses"] = v
	}
	if v, err := expandInstanceMasterInstance(c, f.MasterInstance); err != nil {
		return nil, fmt.Errorf("error expanding MasterInstance into masterInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["masterInstance"] = v
	}
	if v, err := expandInstanceReplicaConfiguration(c, f.ReplicaConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding ReplicaConfiguration into replicaConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["replicaConfiguration"] = v
	}
	if v, err := expandInstanceScheduledMaintenance(c, f.ScheduledMaintenance); err != nil {
		return nil, fmt.Errorf("error expanding ScheduledMaintenance into scheduledMaintenance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scheduledMaintenance"] = v
	}
	if v, err := expandInstanceSettings(c, f.Settings); err != nil {
		return nil, fmt.Errorf("error expanding Settings into settings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["settings"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := expandInstanceReplicaInstancesSlice(c, f.ReplicaInstances); err != nil {
		return nil, fmt.Errorf("error expanding ReplicaInstances into replicaInstances: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["replicaInstances"] = v
	}
	if v, err := expandInstanceServerCaCert(c, f.ServerCaCert); err != nil {
		return nil, fmt.Errorf("error expanding ServerCaCert into serverCaCert: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["serverCaCert"] = v
	}
	if v := f.IPv6Address; !dcl.IsEmptyValueIndirect(v) {
		m["ipv6Address"] = v
	}
	if v := f.ServiceAccountEmailAddress; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmailAddress"] = v
	}
	if v, err := expandInstanceOnPremisesConfiguration(c, f.OnPremisesConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding OnPremisesConfiguration into onPremisesConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["onPremisesConfiguration"] = v
	}
	if v := f.SuspensionReason; !dcl.IsEmptyValueIndirect(v) {
		m["suspensionReason"] = v
	}
	if v, err := expandInstanceDiskEncryptionStatus(c, f.DiskEncryptionStatus); err != nil {
		return nil, fmt.Errorf("error expanding DiskEncryptionStatus into diskEncryptionStatus: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["diskEncryptionStatus"] = v
	}
	if v := f.InstanceUid; !dcl.IsEmptyValueIndirect(v) {
		m["instanceUid"] = v
	}

	return m, nil
}

// flattenInstance flattens Instance from a JSON request object into the
// Instance type.
func flattenInstance(c *Client, i interface{}) *Instance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Instance{}
	r.BackendType = flattenInstanceBackendTypeEnum(m["backendType"])
	r.ConnectionName = dcl.FlattenString(m["connectionName"])
	r.DatabaseVersion = flattenInstanceDatabaseVersionEnum(m["databaseVersion"])
	r.Etag = dcl.FlattenString(m["etag"])
	r.GceZone = dcl.FlattenString(m["gceZone"])
	r.InstanceType = flattenInstanceInstanceTypeEnum(m["instanceType"])
	r.MasterInstanceName = dcl.FlattenString(m["masterInstanceName"])
	r.MaxDiskSize = flattenInstanceMaxDiskSize(c, m["maxDiskSize"])
	r.Name = dcl.FlattenString(m["name"])
	r.Project = dcl.FlattenString(m["project"])
	r.Region = dcl.FlattenString(m["region"])
	r.RootPassword = dcl.FlattenString(m["rootPassword"])
	r.CurrentDiskSize = flattenInstanceCurrentDiskSize(c, m["currentDiskSize"])
	r.DiskEncryptionConfiguration = flattenInstanceDiskEncryptionConfiguration(c, m["diskEncryptionConfiguration"])
	r.FailoverReplica = flattenInstanceFailoverReplica(c, m["failoverReplica"])
	r.IPAddresses = flattenInstanceIPAddressesSlice(c, m["ipAddresses"])
	r.MasterInstance = flattenInstanceMasterInstance(c, m["masterInstance"])
	r.ReplicaConfiguration = flattenInstanceReplicaConfiguration(c, m["replicaConfiguration"])
	r.ScheduledMaintenance = flattenInstanceScheduledMaintenance(c, m["scheduledMaintenance"])
	r.Settings = flattenInstanceSettings(c, m["settings"])
	r.State = dcl.FlattenString(m["state"])
	r.ReplicaInstances = flattenInstanceReplicaInstancesSlice(c, m["replicaInstances"])
	r.ServerCaCert = flattenInstanceServerCaCert(c, m["serverCaCert"])
	r.IPv6Address = dcl.FlattenString(m["ipv6Address"])
	r.ServiceAccountEmailAddress = dcl.FlattenString(m["serviceAccountEmailAddress"])
	r.OnPremisesConfiguration = flattenInstanceOnPremisesConfiguration(c, m["onPremisesConfiguration"])
	r.SuspensionReason = dcl.FlattenStringSlice(m["suspensionReason"])
	r.DiskEncryptionStatus = flattenInstanceDiskEncryptionStatus(c, m["diskEncryptionStatus"])
	r.InstanceUid = dcl.FlattenString(m["instanceUid"])

	return r
}

// expandInstanceMaxDiskSizeMap expands the contents of InstanceMaxDiskSize into a JSON
// request object.
func expandInstanceMaxDiskSizeMap(c *Client, f map[string]InstanceMaxDiskSize) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceMaxDiskSize(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceMaxDiskSizeSlice expands the contents of InstanceMaxDiskSize into a JSON
// request object.
func expandInstanceMaxDiskSizeSlice(c *Client, f []InstanceMaxDiskSize) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceMaxDiskSize(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceMaxDiskSizeMap flattens the contents of InstanceMaxDiskSize from a JSON
// response object.
func flattenInstanceMaxDiskSizeMap(c *Client, i interface{}) map[string]InstanceMaxDiskSize {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceMaxDiskSize{}
	}

	if len(a) == 0 {
		return map[string]InstanceMaxDiskSize{}
	}

	items := make(map[string]InstanceMaxDiskSize)
	for k, item := range a {
		items[k] = *flattenInstanceMaxDiskSize(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceMaxDiskSizeSlice flattens the contents of InstanceMaxDiskSize from a JSON
// response object.
func flattenInstanceMaxDiskSizeSlice(c *Client, i interface{}) []InstanceMaxDiskSize {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceMaxDiskSize{}
	}

	if len(a) == 0 {
		return []InstanceMaxDiskSize{}
	}

	items := make([]InstanceMaxDiskSize, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceMaxDiskSize(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceMaxDiskSize expands an instance of InstanceMaxDiskSize into a JSON
// request object.
func expandInstanceMaxDiskSize(c *Client, f *InstanceMaxDiskSize) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceMaxDiskSize flattens an instance of InstanceMaxDiskSize from a JSON
// response object.
func flattenInstanceMaxDiskSize(c *Client, i interface{}) *InstanceMaxDiskSize {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceMaxDiskSize{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceCurrentDiskSizeMap expands the contents of InstanceCurrentDiskSize into a JSON
// request object.
func expandInstanceCurrentDiskSizeMap(c *Client, f map[string]InstanceCurrentDiskSize) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceCurrentDiskSize(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceCurrentDiskSizeSlice expands the contents of InstanceCurrentDiskSize into a JSON
// request object.
func expandInstanceCurrentDiskSizeSlice(c *Client, f []InstanceCurrentDiskSize) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceCurrentDiskSize(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceCurrentDiskSizeMap flattens the contents of InstanceCurrentDiskSize from a JSON
// response object.
func flattenInstanceCurrentDiskSizeMap(c *Client, i interface{}) map[string]InstanceCurrentDiskSize {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceCurrentDiskSize{}
	}

	if len(a) == 0 {
		return map[string]InstanceCurrentDiskSize{}
	}

	items := make(map[string]InstanceCurrentDiskSize)
	for k, item := range a {
		items[k] = *flattenInstanceCurrentDiskSize(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceCurrentDiskSizeSlice flattens the contents of InstanceCurrentDiskSize from a JSON
// response object.
func flattenInstanceCurrentDiskSizeSlice(c *Client, i interface{}) []InstanceCurrentDiskSize {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceCurrentDiskSize{}
	}

	if len(a) == 0 {
		return []InstanceCurrentDiskSize{}
	}

	items := make([]InstanceCurrentDiskSize, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceCurrentDiskSize(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceCurrentDiskSize expands an instance of InstanceCurrentDiskSize into a JSON
// request object.
func expandInstanceCurrentDiskSize(c *Client, f *InstanceCurrentDiskSize) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceCurrentDiskSize flattens an instance of InstanceCurrentDiskSize from a JSON
// response object.
func flattenInstanceCurrentDiskSize(c *Client, i interface{}) *InstanceCurrentDiskSize {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceCurrentDiskSize{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceDiskEncryptionConfigurationMap expands the contents of InstanceDiskEncryptionConfiguration into a JSON
// request object.
func expandInstanceDiskEncryptionConfigurationMap(c *Client, f map[string]InstanceDiskEncryptionConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceDiskEncryptionConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceDiskEncryptionConfigurationSlice expands the contents of InstanceDiskEncryptionConfiguration into a JSON
// request object.
func expandInstanceDiskEncryptionConfigurationSlice(c *Client, f []InstanceDiskEncryptionConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceDiskEncryptionConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceDiskEncryptionConfigurationMap flattens the contents of InstanceDiskEncryptionConfiguration from a JSON
// response object.
func flattenInstanceDiskEncryptionConfigurationMap(c *Client, i interface{}) map[string]InstanceDiskEncryptionConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceDiskEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceDiskEncryptionConfiguration{}
	}

	items := make(map[string]InstanceDiskEncryptionConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceDiskEncryptionConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceDiskEncryptionConfigurationSlice flattens the contents of InstanceDiskEncryptionConfiguration from a JSON
// response object.
func flattenInstanceDiskEncryptionConfigurationSlice(c *Client, i interface{}) []InstanceDiskEncryptionConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDiskEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceDiskEncryptionConfiguration{}
	}

	items := make([]InstanceDiskEncryptionConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDiskEncryptionConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceDiskEncryptionConfiguration expands an instance of InstanceDiskEncryptionConfiguration into a JSON
// request object.
func expandInstanceDiskEncryptionConfiguration(c *Client, f *InstanceDiskEncryptionConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceDiskEncryptionConfiguration flattens an instance of InstanceDiskEncryptionConfiguration from a JSON
// response object.
func flattenInstanceDiskEncryptionConfiguration(c *Client, i interface{}) *InstanceDiskEncryptionConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceDiskEncryptionConfiguration{}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceFailoverReplicaMap expands the contents of InstanceFailoverReplica into a JSON
// request object.
func expandInstanceFailoverReplicaMap(c *Client, f map[string]InstanceFailoverReplica) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceFailoverReplica(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceFailoverReplicaSlice expands the contents of InstanceFailoverReplica into a JSON
// request object.
func expandInstanceFailoverReplicaSlice(c *Client, f []InstanceFailoverReplica) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceFailoverReplica(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceFailoverReplicaMap flattens the contents of InstanceFailoverReplica from a JSON
// response object.
func flattenInstanceFailoverReplicaMap(c *Client, i interface{}) map[string]InstanceFailoverReplica {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFailoverReplica{}
	}

	if len(a) == 0 {
		return map[string]InstanceFailoverReplica{}
	}

	items := make(map[string]InstanceFailoverReplica)
	for k, item := range a {
		items[k] = *flattenInstanceFailoverReplica(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceFailoverReplicaSlice flattens the contents of InstanceFailoverReplica from a JSON
// response object.
func flattenInstanceFailoverReplicaSlice(c *Client, i interface{}) []InstanceFailoverReplica {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFailoverReplica{}
	}

	if len(a) == 0 {
		return []InstanceFailoverReplica{}
	}

	items := make([]InstanceFailoverReplica, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFailoverReplica(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceFailoverReplica expands an instance of InstanceFailoverReplica into a JSON
// request object.
func expandInstanceFailoverReplica(c *Client, f *InstanceFailoverReplica) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Available; !dcl.IsEmptyValueIndirect(v) {
		m["available"] = v
	}
	if v, err := expandInstanceFailoverReplicaFailoverInstance(c, f.FailoverInstance); err != nil {
		return nil, fmt.Errorf("error expanding FailoverInstance into failoverInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["failoverInstance"] = v
	}

	return m, nil
}

// flattenInstanceFailoverReplica flattens an instance of InstanceFailoverReplica from a JSON
// response object.
func flattenInstanceFailoverReplica(c *Client, i interface{}) *InstanceFailoverReplica {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceFailoverReplica{}
	r.Name = dcl.FlattenString(m["name"])
	r.Available = dcl.FlattenBool(m["available"])
	r.FailoverInstance = flattenInstanceFailoverReplicaFailoverInstance(c, m["failoverInstance"])

	return r
}

// expandInstanceFailoverReplicaFailoverInstanceMap expands the contents of InstanceFailoverReplicaFailoverInstance into a JSON
// request object.
func expandInstanceFailoverReplicaFailoverInstanceMap(c *Client, f map[string]InstanceFailoverReplicaFailoverInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceFailoverReplicaFailoverInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceFailoverReplicaFailoverInstanceSlice expands the contents of InstanceFailoverReplicaFailoverInstance into a JSON
// request object.
func expandInstanceFailoverReplicaFailoverInstanceSlice(c *Client, f []InstanceFailoverReplicaFailoverInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceFailoverReplicaFailoverInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceFailoverReplicaFailoverInstanceMap flattens the contents of InstanceFailoverReplicaFailoverInstance from a JSON
// response object.
func flattenInstanceFailoverReplicaFailoverInstanceMap(c *Client, i interface{}) map[string]InstanceFailoverReplicaFailoverInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFailoverReplicaFailoverInstance{}
	}

	if len(a) == 0 {
		return map[string]InstanceFailoverReplicaFailoverInstance{}
	}

	items := make(map[string]InstanceFailoverReplicaFailoverInstance)
	for k, item := range a {
		items[k] = *flattenInstanceFailoverReplicaFailoverInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceFailoverReplicaFailoverInstanceSlice flattens the contents of InstanceFailoverReplicaFailoverInstance from a JSON
// response object.
func flattenInstanceFailoverReplicaFailoverInstanceSlice(c *Client, i interface{}) []InstanceFailoverReplicaFailoverInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFailoverReplicaFailoverInstance{}
	}

	if len(a) == 0 {
		return []InstanceFailoverReplicaFailoverInstance{}
	}

	items := make([]InstanceFailoverReplicaFailoverInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFailoverReplicaFailoverInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceFailoverReplicaFailoverInstance expands an instance of InstanceFailoverReplicaFailoverInstance into a JSON
// request object.
func expandInstanceFailoverReplicaFailoverInstance(c *Client, f *InstanceFailoverReplicaFailoverInstance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}

	return m, nil
}

// flattenInstanceFailoverReplicaFailoverInstance flattens an instance of InstanceFailoverReplicaFailoverInstance from a JSON
// response object.
func flattenInstanceFailoverReplicaFailoverInstance(c *Client, i interface{}) *InstanceFailoverReplicaFailoverInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceFailoverReplicaFailoverInstance{}
	r.Name = dcl.FlattenString(m["name"])
	r.Region = dcl.FlattenString(m["region"])

	return r
}

// expandInstanceIPAddressesMap expands the contents of InstanceIPAddresses into a JSON
// request object.
func expandInstanceIPAddressesMap(c *Client, f map[string]InstanceIPAddresses) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceIPAddresses(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceIPAddressesSlice expands the contents of InstanceIPAddresses into a JSON
// request object.
func expandInstanceIPAddressesSlice(c *Client, f []InstanceIPAddresses) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceIPAddresses(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceIPAddressesMap flattens the contents of InstanceIPAddresses from a JSON
// response object.
func flattenInstanceIPAddressesMap(c *Client, i interface{}) map[string]InstanceIPAddresses {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceIPAddresses{}
	}

	if len(a) == 0 {
		return map[string]InstanceIPAddresses{}
	}

	items := make(map[string]InstanceIPAddresses)
	for k, item := range a {
		items[k] = *flattenInstanceIPAddresses(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceIPAddressesSlice flattens the contents of InstanceIPAddresses from a JSON
// response object.
func flattenInstanceIPAddressesSlice(c *Client, i interface{}) []InstanceIPAddresses {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceIPAddresses{}
	}

	if len(a) == 0 {
		return []InstanceIPAddresses{}
	}

	items := make([]InstanceIPAddresses, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceIPAddresses(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceIPAddresses expands an instance of InstanceIPAddresses into a JSON
// request object.
func expandInstanceIPAddresses(c *Client, f *InstanceIPAddresses) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.IPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddress"] = v
	}
	if v, err := expandInstanceIPAddressesTimeToRetire(c, f.TimeToRetire); err != nil {
		return nil, fmt.Errorf("error expanding TimeToRetire into timeToRetire: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeToRetire"] = v
	}

	return m, nil
}

// flattenInstanceIPAddresses flattens an instance of InstanceIPAddresses from a JSON
// response object.
func flattenInstanceIPAddresses(c *Client, i interface{}) *InstanceIPAddresses {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceIPAddresses{}
	r.Type = flattenInstanceIPAddressesTypeEnum(m["type"])
	r.IPAddress = dcl.FlattenString(m["ipAddress"])
	r.TimeToRetire = flattenInstanceIPAddressesTimeToRetire(c, m["timeToRetire"])

	return r
}

// expandInstanceIPAddressesTimeToRetireMap expands the contents of InstanceIPAddressesTimeToRetire into a JSON
// request object.
func expandInstanceIPAddressesTimeToRetireMap(c *Client, f map[string]InstanceIPAddressesTimeToRetire) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceIPAddressesTimeToRetire(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceIPAddressesTimeToRetireSlice expands the contents of InstanceIPAddressesTimeToRetire into a JSON
// request object.
func expandInstanceIPAddressesTimeToRetireSlice(c *Client, f []InstanceIPAddressesTimeToRetire) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceIPAddressesTimeToRetire(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceIPAddressesTimeToRetireMap flattens the contents of InstanceIPAddressesTimeToRetire from a JSON
// response object.
func flattenInstanceIPAddressesTimeToRetireMap(c *Client, i interface{}) map[string]InstanceIPAddressesTimeToRetire {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceIPAddressesTimeToRetire{}
	}

	if len(a) == 0 {
		return map[string]InstanceIPAddressesTimeToRetire{}
	}

	items := make(map[string]InstanceIPAddressesTimeToRetire)
	for k, item := range a {
		items[k] = *flattenInstanceIPAddressesTimeToRetire(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceIPAddressesTimeToRetireSlice flattens the contents of InstanceIPAddressesTimeToRetire from a JSON
// response object.
func flattenInstanceIPAddressesTimeToRetireSlice(c *Client, i interface{}) []InstanceIPAddressesTimeToRetire {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceIPAddressesTimeToRetire{}
	}

	if len(a) == 0 {
		return []InstanceIPAddressesTimeToRetire{}
	}

	items := make([]InstanceIPAddressesTimeToRetire, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceIPAddressesTimeToRetire(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceIPAddressesTimeToRetire expands an instance of InstanceIPAddressesTimeToRetire into a JSON
// request object.
func expandInstanceIPAddressesTimeToRetire(c *Client, f *InstanceIPAddressesTimeToRetire) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenInstanceIPAddressesTimeToRetire flattens an instance of InstanceIPAddressesTimeToRetire from a JSON
// response object.
func flattenInstanceIPAddressesTimeToRetire(c *Client, i interface{}) *InstanceIPAddressesTimeToRetire {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceIPAddressesTimeToRetire{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandInstanceMasterInstanceMap expands the contents of InstanceMasterInstance into a JSON
// request object.
func expandInstanceMasterInstanceMap(c *Client, f map[string]InstanceMasterInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceMasterInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceMasterInstanceSlice expands the contents of InstanceMasterInstance into a JSON
// request object.
func expandInstanceMasterInstanceSlice(c *Client, f []InstanceMasterInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceMasterInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceMasterInstanceMap flattens the contents of InstanceMasterInstance from a JSON
// response object.
func flattenInstanceMasterInstanceMap(c *Client, i interface{}) map[string]InstanceMasterInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceMasterInstance{}
	}

	if len(a) == 0 {
		return map[string]InstanceMasterInstance{}
	}

	items := make(map[string]InstanceMasterInstance)
	for k, item := range a {
		items[k] = *flattenInstanceMasterInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceMasterInstanceSlice flattens the contents of InstanceMasterInstance from a JSON
// response object.
func flattenInstanceMasterInstanceSlice(c *Client, i interface{}) []InstanceMasterInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceMasterInstance{}
	}

	if len(a) == 0 {
		return []InstanceMasterInstance{}
	}

	items := make([]InstanceMasterInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceMasterInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceMasterInstance expands an instance of InstanceMasterInstance into a JSON
// request object.
func expandInstanceMasterInstance(c *Client, f *InstanceMasterInstance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}

	return m, nil
}

// flattenInstanceMasterInstance flattens an instance of InstanceMasterInstance from a JSON
// response object.
func flattenInstanceMasterInstance(c *Client, i interface{}) *InstanceMasterInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceMasterInstance{}
	r.Name = dcl.FlattenString(m["name"])
	r.Region = dcl.FlattenString(m["region"])

	return r
}

// expandInstanceReplicaConfigurationMap expands the contents of InstanceReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationMap(c *Client, f map[string]InstanceReplicaConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationSlice expands the contents of InstanceReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationSlice(c *Client, f []InstanceReplicaConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationMap flattens the contents of InstanceReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationSlice flattens the contents of InstanceReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfiguration{}
	}

	items := make([]InstanceReplicaConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfiguration expands an instance of InstanceReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfiguration(c *Client, f *InstanceReplicaConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v, err := expandInstanceReplicaConfigurationMysqlReplicaConfiguration(c, f.MysqlReplicaConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding MysqlReplicaConfiguration into mysqlReplicaConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mysqlReplicaConfiguration"] = v
	}
	if v := f.FailoverTarget; !dcl.IsEmptyValueIndirect(v) {
		m["failoverTarget"] = v
	}
	if v, err := expandInstanceReplicaConfigurationReplicaPoolConfiguration(c, f.ReplicaPoolConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding ReplicaPoolConfiguration into replicaPoolConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["replicaPoolConfiguration"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfiguration flattens an instance of InstanceReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfiguration(c *Client, i interface{}) *InstanceReplicaConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfiguration{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.MysqlReplicaConfiguration = flattenInstanceReplicaConfigurationMysqlReplicaConfiguration(c, m["mysqlReplicaConfiguration"])
	r.FailoverTarget = dcl.FlattenBool(m["failoverTarget"])
	r.ReplicaPoolConfiguration = flattenInstanceReplicaConfigurationReplicaPoolConfiguration(c, m["replicaPoolConfiguration"])

	return r
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationMap expands the contents of InstanceReplicaConfigurationMysqlReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationMap(c *Client, f map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationSlice expands the contents of InstanceReplicaConfigurationMysqlReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationSlice(c *Client, f []InstanceReplicaConfigurationMysqlReplicaConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMap flattens the contents of InstanceReplicaConfigurationMysqlReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationMysqlReplicaConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationSlice flattens the contents of InstanceReplicaConfigurationMysqlReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfigurationMysqlReplicaConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	}

	items := make([]InstanceReplicaConfigurationMysqlReplicaConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationMysqlReplicaConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationMysqlReplicaConfiguration expands an instance of InstanceReplicaConfigurationMysqlReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfiguration(c *Client, f *InstanceReplicaConfigurationMysqlReplicaConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DumpFilePath; !dcl.IsEmptyValueIndirect(v) {
		m["dumpFilePath"] = v
	}
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}
	if v := f.ConnectRetryInterval; !dcl.IsEmptyValueIndirect(v) {
		m["connectRetryInterval"] = v
	}
	if v, err := expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, f.MasterHeartbeatPeriod); err != nil {
		return nil, fmt.Errorf("error expanding MasterHeartbeatPeriod into masterHeartbeatPeriod: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["masterHeartbeatPeriod"] = v
	}
	if v := f.CaCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["caCertificate"] = v
	}
	if v := f.ClientCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["clientCertificate"] = v
	}
	if v := f.ClientKey; !dcl.IsEmptyValueIndirect(v) {
		m["clientKey"] = v
	}
	if v := f.SslCipher; !dcl.IsEmptyValueIndirect(v) {
		m["sslCipher"] = v
	}
	if v := f.VerifyServerCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["verifyServerCertificate"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfiguration flattens an instance of InstanceReplicaConfigurationMysqlReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfiguration(c *Client, i interface{}) *InstanceReplicaConfigurationMysqlReplicaConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	r.DumpFilePath = dcl.FlattenString(m["dumpFilePath"])
	r.Username = dcl.FlattenString(m["username"])
	r.Password = dcl.FlattenString(m["password"])
	r.ConnectRetryInterval = dcl.FlattenInteger(m["connectRetryInterval"])
	r.MasterHeartbeatPeriod = flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, m["masterHeartbeatPeriod"])
	r.CaCertificate = dcl.FlattenString(m["caCertificate"])
	r.ClientCertificate = dcl.FlattenString(m["clientCertificate"])
	r.ClientKey = dcl.FlattenString(m["clientKey"])
	r.SslCipher = dcl.FlattenString(m["sslCipher"])
	r.VerifyServerCertificate = dcl.FlattenBool(m["verifyServerCertificate"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodMap expands the contents of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodMap(c *Client, f map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice expands the contents of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice(c *Client, f []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodMap flattens the contents of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	}

	items := make(map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice flattens the contents of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice(c *Client, i interface{}) []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	}

	items := make([]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod expands an instance of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c *Client, f *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod flattens an instance of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c *Client, i interface{}) *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationMap expands the contents of InstanceReplicaConfigurationReplicaPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationMap(c *Client, f map[string]InstanceReplicaConfigurationReplicaPoolConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationSlice expands the contents of InstanceReplicaConfigurationReplicaPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationSlice(c *Client, f []InstanceReplicaConfigurationReplicaPoolConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationMap flattens the contents of InstanceReplicaConfigurationReplicaPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationReplicaPoolConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfigurationReplicaPoolConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationReplicaPoolConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationSlice flattens the contents of InstanceReplicaConfigurationReplicaPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfigurationReplicaPoolConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationReplicaPoolConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationReplicaPoolConfiguration{}
	}

	items := make([]InstanceReplicaConfigurationReplicaPoolConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationReplicaPoolConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationReplicaPoolConfiguration expands an instance of InstanceReplicaConfigurationReplicaPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfiguration(c *Client, f *InstanceReplicaConfigurationReplicaPoolConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, f.StaticPoolConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding StaticPoolConfiguration into staticPoolConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["staticPoolConfiguration"] = v
	}
	if v, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, f.AutoscalingPoolConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding AutoscalingPoolConfiguration into autoscalingPoolConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoscalingPoolConfiguration"] = v
	}
	if v := f.ReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["replicaCount"] = v
	}
	if v := f.ExposeReplicaIP; !dcl.IsEmptyValueIndirect(v) {
		m["exposeReplicaIp"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfiguration flattens an instance of InstanceReplicaConfigurationReplicaPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfiguration(c *Client, i interface{}) *InstanceReplicaConfigurationReplicaPoolConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationReplicaPoolConfiguration{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.StaticPoolConfiguration = flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, m["staticPoolConfiguration"])
	r.AutoscalingPoolConfiguration = flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, m["autoscalingPoolConfiguration"])
	r.ReplicaCount = dcl.FlattenInteger(m["replicaCount"])
	r.ExposeReplicaIP = dcl.FlattenBool(m["exposeReplicaIp"])

	return r
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationMap expands the contents of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationMap(c *Client, f map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice expands the contents of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice(c *Client, f []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationMap flattens the contents of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice flattens the contents of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	}

	items := make([]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration expands an instance of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c *Client, f *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.ReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["replicaCount"] = v
	}
	if v := f.ExposeReplicaIP; !dcl.IsEmptyValueIndirect(v) {
		m["exposeReplicaIP"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration flattens an instance of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c *Client, i interface{}) *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.ReplicaCount = dcl.FlattenInteger(m["replicaCount"])
	r.ExposeReplicaIP = dcl.FlattenBool(m["exposeReplicaIP"])

	return r
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationMap expands the contents of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationMap(c *Client, f map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice expands the contents of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice(c *Client, f []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationMap flattens the contents of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice flattens the contents of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	}

	items := make([]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration expands an instance of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c *Client, f *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.MinReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["minReplicaCount"] = v
	}
	if v := f.MaxReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxReplicaCount"] = v
	}
	if v := f.TargetCpuUtil; !dcl.IsEmptyValueIndirect(v) {
		m["targetCpuUtil"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration flattens an instance of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c *Client, i interface{}) *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.MinReplicaCount = dcl.FlattenInteger(m["minReplicaCount"])
	r.MaxReplicaCount = dcl.FlattenInteger(m["maxReplicaCount"])
	r.TargetCpuUtil = dcl.FlattenDouble(m["targetCpuUtil"])

	return r
}

// expandInstanceScheduledMaintenanceMap expands the contents of InstanceScheduledMaintenance into a JSON
// request object.
func expandInstanceScheduledMaintenanceMap(c *Client, f map[string]InstanceScheduledMaintenance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceScheduledMaintenance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceScheduledMaintenanceSlice expands the contents of InstanceScheduledMaintenance into a JSON
// request object.
func expandInstanceScheduledMaintenanceSlice(c *Client, f []InstanceScheduledMaintenance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceScheduledMaintenance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceScheduledMaintenanceMap flattens the contents of InstanceScheduledMaintenance from a JSON
// response object.
func flattenInstanceScheduledMaintenanceMap(c *Client, i interface{}) map[string]InstanceScheduledMaintenance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceScheduledMaintenance{}
	}

	if len(a) == 0 {
		return map[string]InstanceScheduledMaintenance{}
	}

	items := make(map[string]InstanceScheduledMaintenance)
	for k, item := range a {
		items[k] = *flattenInstanceScheduledMaintenance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceScheduledMaintenanceSlice flattens the contents of InstanceScheduledMaintenance from a JSON
// response object.
func flattenInstanceScheduledMaintenanceSlice(c *Client, i interface{}) []InstanceScheduledMaintenance {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceScheduledMaintenance{}
	}

	if len(a) == 0 {
		return []InstanceScheduledMaintenance{}
	}

	items := make([]InstanceScheduledMaintenance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceScheduledMaintenance(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceScheduledMaintenance expands an instance of InstanceScheduledMaintenance into a JSON
// request object.
func expandInstanceScheduledMaintenance(c *Client, f *InstanceScheduledMaintenance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInstanceScheduledMaintenanceStartTime(c, f.StartTime); err != nil {
		return nil, fmt.Errorf("error expanding StartTime into startTime: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.CanDefer; !dcl.IsEmptyValueIndirect(v) {
		m["canDefer"] = v
	}
	if v := f.CanReschedule; !dcl.IsEmptyValueIndirect(v) {
		m["canReschedule"] = v
	}

	return m, nil
}

// flattenInstanceScheduledMaintenance flattens an instance of InstanceScheduledMaintenance from a JSON
// response object.
func flattenInstanceScheduledMaintenance(c *Client, i interface{}) *InstanceScheduledMaintenance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceScheduledMaintenance{}
	r.StartTime = flattenInstanceScheduledMaintenanceStartTime(c, m["startTime"])
	r.CanDefer = dcl.FlattenBool(m["canDefer"])
	r.CanReschedule = dcl.FlattenBool(m["canReschedule"])

	return r
}

// expandInstanceScheduledMaintenanceStartTimeMap expands the contents of InstanceScheduledMaintenanceStartTime into a JSON
// request object.
func expandInstanceScheduledMaintenanceStartTimeMap(c *Client, f map[string]InstanceScheduledMaintenanceStartTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceScheduledMaintenanceStartTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceScheduledMaintenanceStartTimeSlice expands the contents of InstanceScheduledMaintenanceStartTime into a JSON
// request object.
func expandInstanceScheduledMaintenanceStartTimeSlice(c *Client, f []InstanceScheduledMaintenanceStartTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceScheduledMaintenanceStartTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceScheduledMaintenanceStartTimeMap flattens the contents of InstanceScheduledMaintenanceStartTime from a JSON
// response object.
func flattenInstanceScheduledMaintenanceStartTimeMap(c *Client, i interface{}) map[string]InstanceScheduledMaintenanceStartTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceScheduledMaintenanceStartTime{}
	}

	if len(a) == 0 {
		return map[string]InstanceScheduledMaintenanceStartTime{}
	}

	items := make(map[string]InstanceScheduledMaintenanceStartTime)
	for k, item := range a {
		items[k] = *flattenInstanceScheduledMaintenanceStartTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceScheduledMaintenanceStartTimeSlice flattens the contents of InstanceScheduledMaintenanceStartTime from a JSON
// response object.
func flattenInstanceScheduledMaintenanceStartTimeSlice(c *Client, i interface{}) []InstanceScheduledMaintenanceStartTime {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceScheduledMaintenanceStartTime{}
	}

	if len(a) == 0 {
		return []InstanceScheduledMaintenanceStartTime{}
	}

	items := make([]InstanceScheduledMaintenanceStartTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceScheduledMaintenanceStartTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceScheduledMaintenanceStartTime expands an instance of InstanceScheduledMaintenanceStartTime into a JSON
// request object.
func expandInstanceScheduledMaintenanceStartTime(c *Client, f *InstanceScheduledMaintenanceStartTime) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenInstanceScheduledMaintenanceStartTime flattens an instance of InstanceScheduledMaintenanceStartTime from a JSON
// response object.
func flattenInstanceScheduledMaintenanceStartTime(c *Client, i interface{}) *InstanceScheduledMaintenanceStartTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceScheduledMaintenanceStartTime{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandInstanceSettingsMap expands the contents of InstanceSettings into a JSON
// request object.
func expandInstanceSettingsMap(c *Client, f map[string]InstanceSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsSlice expands the contents of InstanceSettings into a JSON
// request object.
func expandInstanceSettingsSlice(c *Client, f []InstanceSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsMap flattens the contents of InstanceSettings from a JSON
// response object.
func flattenInstanceSettingsMap(c *Client, i interface{}) map[string]InstanceSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettings{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettings{}
	}

	items := make(map[string]InstanceSettings)
	for k, item := range a {
		items[k] = *flattenInstanceSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsSlice flattens the contents of InstanceSettings from a JSON
// response object.
func flattenInstanceSettingsSlice(c *Client, i interface{}) []InstanceSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettings{}
	}

	if len(a) == 0 {
		return []InstanceSettings{}
	}

	items := make([]InstanceSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettings expands an instance of InstanceSettings into a JSON
// request object.
func expandInstanceSettings(c *Client, f *InstanceSettings) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AuthorizedGaeApplications; !dcl.IsEmptyValueIndirect(v) {
		m["authorizedGaeApplications"] = v
	}
	if v := f.Tier; !dcl.IsEmptyValueIndirect(v) {
		m["tier"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.AvailabilityType; !dcl.IsEmptyValueIndirect(v) {
		m["availabilityType"] = v
	}
	if v := f.PricingPlan; !dcl.IsEmptyValueIndirect(v) {
		m["pricingPlan"] = v
	}
	if v := f.ReplicationType; !dcl.IsEmptyValueIndirect(v) {
		m["replicationType"] = v
	}
	if v := f.ActivationPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["activationPolicy"] = v
	}
	if v := f.StorageAutoResize; !dcl.IsEmptyValueIndirect(v) {
		m["storageAutoResize"] = v
	}
	if v := f.DataDiskType; !dcl.IsEmptyValueIndirect(v) {
		m["dataDiskType"] = v
	}
	if v := f.DatabaseReplicationEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["databaseReplicationEnabled"] = v
	}
	if v := f.CrashSafeReplicationEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["crashSafeReplicationEnabled"] = v
	}
	if v, err := expandInstanceSettingsSettingsVersion(c, f.SettingsVersion); err != nil {
		return nil, fmt.Errorf("error expanding SettingsVersion into settingsVersion: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["settingsVersion"] = v
	}
	if v := f.UserLabels; !dcl.IsEmptyValueIndirect(v) {
		m["userLabels"] = v
	}
	if v, err := expandInstanceSettingsStorageAutoResizeLimit(c, f.StorageAutoResizeLimit); err != nil {
		return nil, fmt.Errorf("error expanding StorageAutoResizeLimit into storageAutoResizeLimit: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["storageAutoResizeLimit"] = v
	}
	if v, err := expandInstanceSettingsIPConfiguration(c, f.IPConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding IPConfiguration into ipConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ipConfiguration"] = v
	}
	if v, err := expandInstanceSettingsLocationPreference(c, f.LocationPreference); err != nil {
		return nil, fmt.Errorf("error expanding LocationPreference into locationPreference: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["locationPreference"] = v
	}
	if v, err := expandInstanceSettingsDatabaseFlagsSlice(c, f.DatabaseFlags); err != nil {
		return nil, fmt.Errorf("error expanding DatabaseFlags into databaseFlags: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["databaseFlags"] = v
	}
	if v, err := expandInstanceSettingsMaintenanceWindow(c, f.MaintenanceWindow); err != nil {
		return nil, fmt.Errorf("error expanding MaintenanceWindow into maintenanceWindow: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maintenanceWindow"] = v
	}
	if v, err := expandInstanceSettingsBackupConfiguration(c, f.BackupConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding BackupConfiguration into backupConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["backupConfiguration"] = v
	}
	if v, err := expandInstanceSettingsDataDiskSizeGb(c, f.DataDiskSizeGb); err != nil {
		return nil, fmt.Errorf("error expanding DataDiskSizeGb into dataDiskSizeGb: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dataDiskSizeGb"] = v
	}
	if v, err := expandInstanceSettingsActiveDirectoryConfig(c, f.ActiveDirectoryConfig); err != nil {
		return nil, fmt.Errorf("error expanding ActiveDirectoryConfig into activeDirectoryConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["activeDirectoryConfig"] = v
	}
	if v := f.Collation; !dcl.IsEmptyValueIndirect(v) {
		m["collation"] = v
	}
	if v, err := expandInstanceSettingsDenyMaintenancePeriodsSlice(c, f.DenyMaintenancePeriods); err != nil {
		return nil, fmt.Errorf("error expanding DenyMaintenancePeriods into denyMaintenancePeriods: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["denyMaintenancePeriods"] = v
	}
	if v, err := expandInstanceSettingsInsightsConfig(c, f.InsightsConfig); err != nil {
		return nil, fmt.Errorf("error expanding InsightsConfig into insightsConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["insightsConfig"] = v
	}

	return m, nil
}

// flattenInstanceSettings flattens an instance of InstanceSettings from a JSON
// response object.
func flattenInstanceSettings(c *Client, i interface{}) *InstanceSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettings{}
	r.AuthorizedGaeApplications = dcl.FlattenStringSlice(m["authorizedGaeApplications"])
	r.Tier = dcl.FlattenString(m["tier"])
	r.Kind = dcl.FlattenString(m["kind"])
	r.AvailabilityType = flattenInstanceSettingsAvailabilityTypeEnum(m["availabilityType"])
	r.PricingPlan = flattenInstanceSettingsPricingPlanEnum(m["pricingPlan"])
	r.ReplicationType = flattenInstanceSettingsReplicationTypeEnum(m["replicationType"])
	r.ActivationPolicy = flattenInstanceSettingsActivationPolicyEnum(m["activationPolicy"])
	r.StorageAutoResize = dcl.FlattenBool(m["storageAutoResize"])
	r.DataDiskType = flattenInstanceSettingsDataDiskTypeEnum(m["dataDiskType"])
	r.DatabaseReplicationEnabled = dcl.FlattenBool(m["databaseReplicationEnabled"])
	r.CrashSafeReplicationEnabled = dcl.FlattenBool(m["crashSafeReplicationEnabled"])
	r.SettingsVersion = flattenInstanceSettingsSettingsVersion(c, m["settingsVersion"])
	r.UserLabels = dcl.FlattenKeyValuePairs(m["userLabels"])
	r.StorageAutoResizeLimit = flattenInstanceSettingsStorageAutoResizeLimit(c, m["storageAutoResizeLimit"])
	r.IPConfiguration = flattenInstanceSettingsIPConfiguration(c, m["ipConfiguration"])
	r.LocationPreference = flattenInstanceSettingsLocationPreference(c, m["locationPreference"])
	r.DatabaseFlags = flattenInstanceSettingsDatabaseFlagsSlice(c, m["databaseFlags"])
	r.MaintenanceWindow = flattenInstanceSettingsMaintenanceWindow(c, m["maintenanceWindow"])
	r.BackupConfiguration = flattenInstanceSettingsBackupConfiguration(c, m["backupConfiguration"])
	r.DataDiskSizeGb = flattenInstanceSettingsDataDiskSizeGb(c, m["dataDiskSizeGb"])
	r.ActiveDirectoryConfig = flattenInstanceSettingsActiveDirectoryConfig(c, m["activeDirectoryConfig"])
	r.Collation = dcl.FlattenString(m["collation"])
	r.DenyMaintenancePeriods = flattenInstanceSettingsDenyMaintenancePeriodsSlice(c, m["denyMaintenancePeriods"])
	r.InsightsConfig = flattenInstanceSettingsInsightsConfig(c, m["insightsConfig"])

	return r
}

// expandInstanceSettingsSettingsVersionMap expands the contents of InstanceSettingsSettingsVersion into a JSON
// request object.
func expandInstanceSettingsSettingsVersionMap(c *Client, f map[string]InstanceSettingsSettingsVersion) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsSettingsVersion(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsSettingsVersionSlice expands the contents of InstanceSettingsSettingsVersion into a JSON
// request object.
func expandInstanceSettingsSettingsVersionSlice(c *Client, f []InstanceSettingsSettingsVersion) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsSettingsVersion(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsSettingsVersionMap flattens the contents of InstanceSettingsSettingsVersion from a JSON
// response object.
func flattenInstanceSettingsSettingsVersionMap(c *Client, i interface{}) map[string]InstanceSettingsSettingsVersion {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsSettingsVersion{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsSettingsVersion{}
	}

	items := make(map[string]InstanceSettingsSettingsVersion)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsSettingsVersion(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsSettingsVersionSlice flattens the contents of InstanceSettingsSettingsVersion from a JSON
// response object.
func flattenInstanceSettingsSettingsVersionSlice(c *Client, i interface{}) []InstanceSettingsSettingsVersion {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsSettingsVersion{}
	}

	if len(a) == 0 {
		return []InstanceSettingsSettingsVersion{}
	}

	items := make([]InstanceSettingsSettingsVersion, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsSettingsVersion(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsSettingsVersion expands an instance of InstanceSettingsSettingsVersion into a JSON
// request object.
func expandInstanceSettingsSettingsVersion(c *Client, f *InstanceSettingsSettingsVersion) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceSettingsSettingsVersion flattens an instance of InstanceSettingsSettingsVersion from a JSON
// response object.
func flattenInstanceSettingsSettingsVersion(c *Client, i interface{}) *InstanceSettingsSettingsVersion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsSettingsVersion{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceSettingsStorageAutoResizeLimitMap expands the contents of InstanceSettingsStorageAutoResizeLimit into a JSON
// request object.
func expandInstanceSettingsStorageAutoResizeLimitMap(c *Client, f map[string]InstanceSettingsStorageAutoResizeLimit) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsStorageAutoResizeLimit(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsStorageAutoResizeLimitSlice expands the contents of InstanceSettingsStorageAutoResizeLimit into a JSON
// request object.
func expandInstanceSettingsStorageAutoResizeLimitSlice(c *Client, f []InstanceSettingsStorageAutoResizeLimit) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsStorageAutoResizeLimit(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsStorageAutoResizeLimitMap flattens the contents of InstanceSettingsStorageAutoResizeLimit from a JSON
// response object.
func flattenInstanceSettingsStorageAutoResizeLimitMap(c *Client, i interface{}) map[string]InstanceSettingsStorageAutoResizeLimit {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsStorageAutoResizeLimit{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsStorageAutoResizeLimit{}
	}

	items := make(map[string]InstanceSettingsStorageAutoResizeLimit)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsStorageAutoResizeLimit(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsStorageAutoResizeLimitSlice flattens the contents of InstanceSettingsStorageAutoResizeLimit from a JSON
// response object.
func flattenInstanceSettingsStorageAutoResizeLimitSlice(c *Client, i interface{}) []InstanceSettingsStorageAutoResizeLimit {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsStorageAutoResizeLimit{}
	}

	if len(a) == 0 {
		return []InstanceSettingsStorageAutoResizeLimit{}
	}

	items := make([]InstanceSettingsStorageAutoResizeLimit, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsStorageAutoResizeLimit(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsStorageAutoResizeLimit expands an instance of InstanceSettingsStorageAutoResizeLimit into a JSON
// request object.
func expandInstanceSettingsStorageAutoResizeLimit(c *Client, f *InstanceSettingsStorageAutoResizeLimit) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceSettingsStorageAutoResizeLimit flattens an instance of InstanceSettingsStorageAutoResizeLimit from a JSON
// response object.
func flattenInstanceSettingsStorageAutoResizeLimit(c *Client, i interface{}) *InstanceSettingsStorageAutoResizeLimit {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsStorageAutoResizeLimit{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceSettingsIPConfigurationMap expands the contents of InstanceSettingsIPConfiguration into a JSON
// request object.
func expandInstanceSettingsIPConfigurationMap(c *Client, f map[string]InstanceSettingsIPConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsIPConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsIPConfigurationSlice expands the contents of InstanceSettingsIPConfiguration into a JSON
// request object.
func expandInstanceSettingsIPConfigurationSlice(c *Client, f []InstanceSettingsIPConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsIPConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsIPConfigurationMap flattens the contents of InstanceSettingsIPConfiguration from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationMap(c *Client, i interface{}) map[string]InstanceSettingsIPConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsIPConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsIPConfiguration{}
	}

	items := make(map[string]InstanceSettingsIPConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsIPConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsIPConfigurationSlice flattens the contents of InstanceSettingsIPConfiguration from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationSlice(c *Client, i interface{}) []InstanceSettingsIPConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsIPConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceSettingsIPConfiguration{}
	}

	items := make([]InstanceSettingsIPConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsIPConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsIPConfiguration expands an instance of InstanceSettingsIPConfiguration into a JSON
// request object.
func expandInstanceSettingsIPConfiguration(c *Client, f *InstanceSettingsIPConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IPv4Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["ipv4Enabled"] = v
	}
	if v := f.PrivateNetwork; !dcl.IsEmptyValueIndirect(v) {
		m["privateNetwork"] = v
	}
	if v := f.RequireSsl; !dcl.IsEmptyValueIndirect(v) {
		m["requireSsl"] = v
	}
	if v, err := expandInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c, f.AuthorizedNetworks); err != nil {
		return nil, fmt.Errorf("error expanding AuthorizedNetworks into authorizedNetworks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["authorizedNetworks"] = v
	}

	return m, nil
}

// flattenInstanceSettingsIPConfiguration flattens an instance of InstanceSettingsIPConfiguration from a JSON
// response object.
func flattenInstanceSettingsIPConfiguration(c *Client, i interface{}) *InstanceSettingsIPConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsIPConfiguration{}
	r.IPv4Enabled = dcl.FlattenBool(m["ipv4Enabled"])
	r.PrivateNetwork = dcl.FlattenString(m["privateNetwork"])
	r.RequireSsl = dcl.FlattenBool(m["requireSsl"])
	r.AuthorizedNetworks = flattenInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c, m["authorizedNetworks"])

	return r
}

// expandInstanceSettingsIPConfigurationAuthorizedNetworksMap expands the contents of InstanceSettingsIPConfigurationAuthorizedNetworks into a JSON
// request object.
func expandInstanceSettingsIPConfigurationAuthorizedNetworksMap(c *Client, f map[string]InstanceSettingsIPConfigurationAuthorizedNetworks) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsIPConfigurationAuthorizedNetworks(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsIPConfigurationAuthorizedNetworksSlice expands the contents of InstanceSettingsIPConfigurationAuthorizedNetworks into a JSON
// request object.
func expandInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c *Client, f []InstanceSettingsIPConfigurationAuthorizedNetworks) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsIPConfigurationAuthorizedNetworks(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsIPConfigurationAuthorizedNetworksMap flattens the contents of InstanceSettingsIPConfigurationAuthorizedNetworks from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationAuthorizedNetworksMap(c *Client, i interface{}) map[string]InstanceSettingsIPConfigurationAuthorizedNetworks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsIPConfigurationAuthorizedNetworks{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsIPConfigurationAuthorizedNetworks{}
	}

	items := make(map[string]InstanceSettingsIPConfigurationAuthorizedNetworks)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsIPConfigurationAuthorizedNetworks(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsIPConfigurationAuthorizedNetworksSlice flattens the contents of InstanceSettingsIPConfigurationAuthorizedNetworks from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c *Client, i interface{}) []InstanceSettingsIPConfigurationAuthorizedNetworks {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsIPConfigurationAuthorizedNetworks{}
	}

	if len(a) == 0 {
		return []InstanceSettingsIPConfigurationAuthorizedNetworks{}
	}

	items := make([]InstanceSettingsIPConfigurationAuthorizedNetworks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsIPConfigurationAuthorizedNetworks(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsIPConfigurationAuthorizedNetworks expands an instance of InstanceSettingsIPConfigurationAuthorizedNetworks into a JSON
// request object.
func expandInstanceSettingsIPConfigurationAuthorizedNetworks(c *Client, f *InstanceSettingsIPConfigurationAuthorizedNetworks) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		m["expirationTime"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceSettingsIPConfigurationAuthorizedNetworks flattens an instance of InstanceSettingsIPConfigurationAuthorizedNetworks from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationAuthorizedNetworks(c *Client, i interface{}) *InstanceSettingsIPConfigurationAuthorizedNetworks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsIPConfigurationAuthorizedNetworks{}
	r.Value = dcl.FlattenString(m["value"])
	r.ExpirationTime = dcl.FlattenString(m["expirationTime"])
	r.Name = dcl.FlattenString(m["name"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceSettingsLocationPreferenceMap expands the contents of InstanceSettingsLocationPreference into a JSON
// request object.
func expandInstanceSettingsLocationPreferenceMap(c *Client, f map[string]InstanceSettingsLocationPreference) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsLocationPreference(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsLocationPreferenceSlice expands the contents of InstanceSettingsLocationPreference into a JSON
// request object.
func expandInstanceSettingsLocationPreferenceSlice(c *Client, f []InstanceSettingsLocationPreference) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsLocationPreference(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsLocationPreferenceMap flattens the contents of InstanceSettingsLocationPreference from a JSON
// response object.
func flattenInstanceSettingsLocationPreferenceMap(c *Client, i interface{}) map[string]InstanceSettingsLocationPreference {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsLocationPreference{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsLocationPreference{}
	}

	items := make(map[string]InstanceSettingsLocationPreference)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsLocationPreference(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsLocationPreferenceSlice flattens the contents of InstanceSettingsLocationPreference from a JSON
// response object.
func flattenInstanceSettingsLocationPreferenceSlice(c *Client, i interface{}) []InstanceSettingsLocationPreference {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsLocationPreference{}
	}

	if len(a) == 0 {
		return []InstanceSettingsLocationPreference{}
	}

	items := make([]InstanceSettingsLocationPreference, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsLocationPreference(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsLocationPreference expands an instance of InstanceSettingsLocationPreference into a JSON
// request object.
func expandInstanceSettingsLocationPreference(c *Client, f *InstanceSettingsLocationPreference) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceSettingsLocationPreference flattens an instance of InstanceSettingsLocationPreference from a JSON
// response object.
func flattenInstanceSettingsLocationPreference(c *Client, i interface{}) *InstanceSettingsLocationPreference {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsLocationPreference{}
	r.Zone = dcl.FlattenString(m["zone"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceSettingsDatabaseFlagsMap expands the contents of InstanceSettingsDatabaseFlags into a JSON
// request object.
func expandInstanceSettingsDatabaseFlagsMap(c *Client, f map[string]InstanceSettingsDatabaseFlags) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsDatabaseFlags(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsDatabaseFlagsSlice expands the contents of InstanceSettingsDatabaseFlags into a JSON
// request object.
func expandInstanceSettingsDatabaseFlagsSlice(c *Client, f []InstanceSettingsDatabaseFlags) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsDatabaseFlags(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsDatabaseFlagsMap flattens the contents of InstanceSettingsDatabaseFlags from a JSON
// response object.
func flattenInstanceSettingsDatabaseFlagsMap(c *Client, i interface{}) map[string]InstanceSettingsDatabaseFlags {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsDatabaseFlags{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsDatabaseFlags{}
	}

	items := make(map[string]InstanceSettingsDatabaseFlags)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsDatabaseFlags(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsDatabaseFlagsSlice flattens the contents of InstanceSettingsDatabaseFlags from a JSON
// response object.
func flattenInstanceSettingsDatabaseFlagsSlice(c *Client, i interface{}) []InstanceSettingsDatabaseFlags {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsDatabaseFlags{}
	}

	if len(a) == 0 {
		return []InstanceSettingsDatabaseFlags{}
	}

	items := make([]InstanceSettingsDatabaseFlags, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsDatabaseFlags(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsDatabaseFlags expands an instance of InstanceSettingsDatabaseFlags into a JSON
// request object.
func expandInstanceSettingsDatabaseFlags(c *Client, f *InstanceSettingsDatabaseFlags) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceSettingsDatabaseFlags flattens an instance of InstanceSettingsDatabaseFlags from a JSON
// response object.
func flattenInstanceSettingsDatabaseFlags(c *Client, i interface{}) *InstanceSettingsDatabaseFlags {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsDatabaseFlags{}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandInstanceSettingsMaintenanceWindowMap expands the contents of InstanceSettingsMaintenanceWindow into a JSON
// request object.
func expandInstanceSettingsMaintenanceWindowMap(c *Client, f map[string]InstanceSettingsMaintenanceWindow) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsMaintenanceWindow(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsMaintenanceWindowSlice expands the contents of InstanceSettingsMaintenanceWindow into a JSON
// request object.
func expandInstanceSettingsMaintenanceWindowSlice(c *Client, f []InstanceSettingsMaintenanceWindow) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsMaintenanceWindow(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsMaintenanceWindowMap flattens the contents of InstanceSettingsMaintenanceWindow from a JSON
// response object.
func flattenInstanceSettingsMaintenanceWindowMap(c *Client, i interface{}) map[string]InstanceSettingsMaintenanceWindow {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsMaintenanceWindow{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsMaintenanceWindow{}
	}

	items := make(map[string]InstanceSettingsMaintenanceWindow)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsMaintenanceWindow(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsMaintenanceWindowSlice flattens the contents of InstanceSettingsMaintenanceWindow from a JSON
// response object.
func flattenInstanceSettingsMaintenanceWindowSlice(c *Client, i interface{}) []InstanceSettingsMaintenanceWindow {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsMaintenanceWindow{}
	}

	if len(a) == 0 {
		return []InstanceSettingsMaintenanceWindow{}
	}

	items := make([]InstanceSettingsMaintenanceWindow, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsMaintenanceWindow(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsMaintenanceWindow expands an instance of InstanceSettingsMaintenanceWindow into a JSON
// request object.
func expandInstanceSettingsMaintenanceWindow(c *Client, f *InstanceSettingsMaintenanceWindow) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Hour; !dcl.IsEmptyValueIndirect(v) {
		m["hour"] = v
	}
	if v := f.Day; !dcl.IsEmptyValueIndirect(v) {
		m["day"] = v
	}
	if v := f.UpdateTrack; !dcl.IsEmptyValueIndirect(v) {
		m["updateTrack"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceSettingsMaintenanceWindow flattens an instance of InstanceSettingsMaintenanceWindow from a JSON
// response object.
func flattenInstanceSettingsMaintenanceWindow(c *Client, i interface{}) *InstanceSettingsMaintenanceWindow {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsMaintenanceWindow{}
	r.Hour = dcl.FlattenInteger(m["hour"])
	r.Day = dcl.FlattenInteger(m["day"])
	r.UpdateTrack = flattenInstanceSettingsMaintenanceWindowUpdateTrackEnum(m["updateTrack"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceSettingsBackupConfigurationMap expands the contents of InstanceSettingsBackupConfiguration into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationMap(c *Client, f map[string]InstanceSettingsBackupConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsBackupConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsBackupConfigurationSlice expands the contents of InstanceSettingsBackupConfiguration into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationSlice(c *Client, f []InstanceSettingsBackupConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsBackupConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsBackupConfigurationMap flattens the contents of InstanceSettingsBackupConfiguration from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationMap(c *Client, i interface{}) map[string]InstanceSettingsBackupConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsBackupConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsBackupConfiguration{}
	}

	items := make(map[string]InstanceSettingsBackupConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsBackupConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsBackupConfigurationSlice flattens the contents of InstanceSettingsBackupConfiguration from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationSlice(c *Client, i interface{}) []InstanceSettingsBackupConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsBackupConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceSettingsBackupConfiguration{}
	}

	items := make([]InstanceSettingsBackupConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsBackupConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsBackupConfiguration expands an instance of InstanceSettingsBackupConfiguration into a JSON
// request object.
func expandInstanceSettingsBackupConfiguration(c *Client, f *InstanceSettingsBackupConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.StartTime; !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.BinaryLogEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["binaryLogEnabled"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v, err := expandInstanceSettingsBackupConfigurationBackupRetentionSettings(c, f.BackupRetentionSettings); err != nil {
		return nil, fmt.Errorf("error expanding BackupRetentionSettings into backupRetentionSettings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["backupRetentionSettings"] = v
	}
	if v := f.TransactionLogRetentionDays; !dcl.IsEmptyValueIndirect(v) {
		m["transactionLogRetentionDays"] = v
	}

	return m, nil
}

// flattenInstanceSettingsBackupConfiguration flattens an instance of InstanceSettingsBackupConfiguration from a JSON
// response object.
func flattenInstanceSettingsBackupConfiguration(c *Client, i interface{}) *InstanceSettingsBackupConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsBackupConfiguration{}
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.Kind = dcl.FlattenString(m["kind"])
	r.BinaryLogEnabled = dcl.FlattenBool(m["binaryLogEnabled"])
	r.Location = dcl.FlattenString(m["location"])
	r.BackupRetentionSettings = flattenInstanceSettingsBackupConfigurationBackupRetentionSettings(c, m["backupRetentionSettings"])
	r.TransactionLogRetentionDays = dcl.FlattenInteger(m["transactionLogRetentionDays"])

	return r
}

// expandInstanceSettingsBackupConfigurationBackupRetentionSettingsMap expands the contents of InstanceSettingsBackupConfigurationBackupRetentionSettings into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationBackupRetentionSettingsMap(c *Client, f map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice expands the contents of InstanceSettingsBackupConfigurationBackupRetentionSettings into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice(c *Client, f []InstanceSettingsBackupConfigurationBackupRetentionSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsMap flattens the contents of InstanceSettingsBackupConfigurationBackupRetentionSettings from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsMap(c *Client, i interface{}) map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	}

	items := make(map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsBackupConfigurationBackupRetentionSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice flattens the contents of InstanceSettingsBackupConfigurationBackupRetentionSettings from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice(c *Client, i interface{}) []InstanceSettingsBackupConfigurationBackupRetentionSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	}

	if len(a) == 0 {
		return []InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	}

	items := make([]InstanceSettingsBackupConfigurationBackupRetentionSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsBackupConfigurationBackupRetentionSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsBackupConfigurationBackupRetentionSettings expands an instance of InstanceSettingsBackupConfigurationBackupRetentionSettings into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationBackupRetentionSettings(c *Client, f *InstanceSettingsBackupConfigurationBackupRetentionSettings) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RetentionUnit; !dcl.IsEmptyValueIndirect(v) {
		m["retentionUnit"] = v
	}
	if v := f.RetainedBackups; !dcl.IsEmptyValueIndirect(v) {
		m["retainedBackups"] = v
	}

	return m, nil
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettings flattens an instance of InstanceSettingsBackupConfigurationBackupRetentionSettings from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettings(c *Client, i interface{}) *InstanceSettingsBackupConfigurationBackupRetentionSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	r.RetentionUnit = flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(m["retentionUnit"])
	r.RetainedBackups = dcl.FlattenInteger(m["retainedBackups"])

	return r
}

// expandInstanceSettingsDataDiskSizeGbMap expands the contents of InstanceSettingsDataDiskSizeGb into a JSON
// request object.
func expandInstanceSettingsDataDiskSizeGbMap(c *Client, f map[string]InstanceSettingsDataDiskSizeGb) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsDataDiskSizeGb(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsDataDiskSizeGbSlice expands the contents of InstanceSettingsDataDiskSizeGb into a JSON
// request object.
func expandInstanceSettingsDataDiskSizeGbSlice(c *Client, f []InstanceSettingsDataDiskSizeGb) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsDataDiskSizeGb(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsDataDiskSizeGbMap flattens the contents of InstanceSettingsDataDiskSizeGb from a JSON
// response object.
func flattenInstanceSettingsDataDiskSizeGbMap(c *Client, i interface{}) map[string]InstanceSettingsDataDiskSizeGb {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsDataDiskSizeGb{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsDataDiskSizeGb{}
	}

	items := make(map[string]InstanceSettingsDataDiskSizeGb)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsDataDiskSizeGb(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsDataDiskSizeGbSlice flattens the contents of InstanceSettingsDataDiskSizeGb from a JSON
// response object.
func flattenInstanceSettingsDataDiskSizeGbSlice(c *Client, i interface{}) []InstanceSettingsDataDiskSizeGb {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsDataDiskSizeGb{}
	}

	if len(a) == 0 {
		return []InstanceSettingsDataDiskSizeGb{}
	}

	items := make([]InstanceSettingsDataDiskSizeGb, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsDataDiskSizeGb(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsDataDiskSizeGb expands an instance of InstanceSettingsDataDiskSizeGb into a JSON
// request object.
func expandInstanceSettingsDataDiskSizeGb(c *Client, f *InstanceSettingsDataDiskSizeGb) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceSettingsDataDiskSizeGb flattens an instance of InstanceSettingsDataDiskSizeGb from a JSON
// response object.
func flattenInstanceSettingsDataDiskSizeGb(c *Client, i interface{}) *InstanceSettingsDataDiskSizeGb {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsDataDiskSizeGb{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceSettingsActiveDirectoryConfigMap expands the contents of InstanceSettingsActiveDirectoryConfig into a JSON
// request object.
func expandInstanceSettingsActiveDirectoryConfigMap(c *Client, f map[string]InstanceSettingsActiveDirectoryConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsActiveDirectoryConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsActiveDirectoryConfigSlice expands the contents of InstanceSettingsActiveDirectoryConfig into a JSON
// request object.
func expandInstanceSettingsActiveDirectoryConfigSlice(c *Client, f []InstanceSettingsActiveDirectoryConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsActiveDirectoryConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsActiveDirectoryConfigMap flattens the contents of InstanceSettingsActiveDirectoryConfig from a JSON
// response object.
func flattenInstanceSettingsActiveDirectoryConfigMap(c *Client, i interface{}) map[string]InstanceSettingsActiveDirectoryConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsActiveDirectoryConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsActiveDirectoryConfig{}
	}

	items := make(map[string]InstanceSettingsActiveDirectoryConfig)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsActiveDirectoryConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsActiveDirectoryConfigSlice flattens the contents of InstanceSettingsActiveDirectoryConfig from a JSON
// response object.
func flattenInstanceSettingsActiveDirectoryConfigSlice(c *Client, i interface{}) []InstanceSettingsActiveDirectoryConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsActiveDirectoryConfig{}
	}

	if len(a) == 0 {
		return []InstanceSettingsActiveDirectoryConfig{}
	}

	items := make([]InstanceSettingsActiveDirectoryConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsActiveDirectoryConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsActiveDirectoryConfig expands an instance of InstanceSettingsActiveDirectoryConfig into a JSON
// request object.
func expandInstanceSettingsActiveDirectoryConfig(c *Client, f *InstanceSettingsActiveDirectoryConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.Domain; !dcl.IsEmptyValueIndirect(v) {
		m["domain"] = v
	}

	return m, nil
}

// flattenInstanceSettingsActiveDirectoryConfig flattens an instance of InstanceSettingsActiveDirectoryConfig from a JSON
// response object.
func flattenInstanceSettingsActiveDirectoryConfig(c *Client, i interface{}) *InstanceSettingsActiveDirectoryConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsActiveDirectoryConfig{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.Domain = dcl.FlattenString(m["domain"])

	return r
}

// expandInstanceSettingsDenyMaintenancePeriodsMap expands the contents of InstanceSettingsDenyMaintenancePeriods into a JSON
// request object.
func expandInstanceSettingsDenyMaintenancePeriodsMap(c *Client, f map[string]InstanceSettingsDenyMaintenancePeriods) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsDenyMaintenancePeriods(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsDenyMaintenancePeriodsSlice expands the contents of InstanceSettingsDenyMaintenancePeriods into a JSON
// request object.
func expandInstanceSettingsDenyMaintenancePeriodsSlice(c *Client, f []InstanceSettingsDenyMaintenancePeriods) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsDenyMaintenancePeriods(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsDenyMaintenancePeriodsMap flattens the contents of InstanceSettingsDenyMaintenancePeriods from a JSON
// response object.
func flattenInstanceSettingsDenyMaintenancePeriodsMap(c *Client, i interface{}) map[string]InstanceSettingsDenyMaintenancePeriods {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsDenyMaintenancePeriods{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsDenyMaintenancePeriods{}
	}

	items := make(map[string]InstanceSettingsDenyMaintenancePeriods)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsDenyMaintenancePeriods(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsDenyMaintenancePeriodsSlice flattens the contents of InstanceSettingsDenyMaintenancePeriods from a JSON
// response object.
func flattenInstanceSettingsDenyMaintenancePeriodsSlice(c *Client, i interface{}) []InstanceSettingsDenyMaintenancePeriods {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsDenyMaintenancePeriods{}
	}

	if len(a) == 0 {
		return []InstanceSettingsDenyMaintenancePeriods{}
	}

	items := make([]InstanceSettingsDenyMaintenancePeriods, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsDenyMaintenancePeriods(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsDenyMaintenancePeriods expands an instance of InstanceSettingsDenyMaintenancePeriods into a JSON
// request object.
func expandInstanceSettingsDenyMaintenancePeriods(c *Client, f *InstanceSettingsDenyMaintenancePeriods) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.StartDate; !dcl.IsEmptyValueIndirect(v) {
		m["startDate"] = v
	}
	if v := f.EndDate; !dcl.IsEmptyValueIndirect(v) {
		m["endDate"] = v
	}
	if v := f.Time; !dcl.IsEmptyValueIndirect(v) {
		m["time"] = v
	}

	return m, nil
}

// flattenInstanceSettingsDenyMaintenancePeriods flattens an instance of InstanceSettingsDenyMaintenancePeriods from a JSON
// response object.
func flattenInstanceSettingsDenyMaintenancePeriods(c *Client, i interface{}) *InstanceSettingsDenyMaintenancePeriods {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsDenyMaintenancePeriods{}
	r.StartDate = dcl.FlattenString(m["startDate"])
	r.EndDate = dcl.FlattenString(m["endDate"])
	r.Time = dcl.FlattenString(m["time"])

	return r
}

// expandInstanceSettingsInsightsConfigMap expands the contents of InstanceSettingsInsightsConfig into a JSON
// request object.
func expandInstanceSettingsInsightsConfigMap(c *Client, f map[string]InstanceSettingsInsightsConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsInsightsConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsInsightsConfigSlice expands the contents of InstanceSettingsInsightsConfig into a JSON
// request object.
func expandInstanceSettingsInsightsConfigSlice(c *Client, f []InstanceSettingsInsightsConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsInsightsConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsInsightsConfigMap flattens the contents of InstanceSettingsInsightsConfig from a JSON
// response object.
func flattenInstanceSettingsInsightsConfigMap(c *Client, i interface{}) map[string]InstanceSettingsInsightsConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsInsightsConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsInsightsConfig{}
	}

	items := make(map[string]InstanceSettingsInsightsConfig)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsInsightsConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsInsightsConfigSlice flattens the contents of InstanceSettingsInsightsConfig from a JSON
// response object.
func flattenInstanceSettingsInsightsConfigSlice(c *Client, i interface{}) []InstanceSettingsInsightsConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsInsightsConfig{}
	}

	if len(a) == 0 {
		return []InstanceSettingsInsightsConfig{}
	}

	items := make([]InstanceSettingsInsightsConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsInsightsConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsInsightsConfig expands an instance of InstanceSettingsInsightsConfig into a JSON
// request object.
func expandInstanceSettingsInsightsConfig(c *Client, f *InstanceSettingsInsightsConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryInsightsEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["queryInsightsEnabled"] = v
	}
	if v := f.RecordClientAddress; !dcl.IsEmptyValueIndirect(v) {
		m["recordClientAddress"] = v
	}
	if v := f.RecordApplicationTags; !dcl.IsEmptyValueIndirect(v) {
		m["recordApplicationTags"] = v
	}
	if v := f.QueryStringLength; !dcl.IsEmptyValueIndirect(v) {
		m["queryStringLength"] = v
	}

	return m, nil
}

// flattenInstanceSettingsInsightsConfig flattens an instance of InstanceSettingsInsightsConfig from a JSON
// response object.
func flattenInstanceSettingsInsightsConfig(c *Client, i interface{}) *InstanceSettingsInsightsConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsInsightsConfig{}
	r.QueryInsightsEnabled = dcl.FlattenBool(m["queryInsightsEnabled"])
	r.RecordClientAddress = dcl.FlattenBool(m["recordClientAddress"])
	r.RecordApplicationTags = dcl.FlattenBool(m["recordApplicationTags"])
	r.QueryStringLength = dcl.FlattenInteger(m["queryStringLength"])

	return r
}

// expandInstanceReplicaInstancesMap expands the contents of InstanceReplicaInstances into a JSON
// request object.
func expandInstanceReplicaInstancesMap(c *Client, f map[string]InstanceReplicaInstances) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaInstances(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaInstancesSlice expands the contents of InstanceReplicaInstances into a JSON
// request object.
func expandInstanceReplicaInstancesSlice(c *Client, f []InstanceReplicaInstances) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaInstances(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaInstancesMap flattens the contents of InstanceReplicaInstances from a JSON
// response object.
func flattenInstanceReplicaInstancesMap(c *Client, i interface{}) map[string]InstanceReplicaInstances {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaInstances{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaInstances{}
	}

	items := make(map[string]InstanceReplicaInstances)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaInstances(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaInstancesSlice flattens the contents of InstanceReplicaInstances from a JSON
// response object.
func flattenInstanceReplicaInstancesSlice(c *Client, i interface{}) []InstanceReplicaInstances {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaInstances{}
	}

	if len(a) == 0 {
		return []InstanceReplicaInstances{}
	}

	items := make([]InstanceReplicaInstances, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaInstances(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaInstances expands an instance of InstanceReplicaInstances into a JSON
// request object.
func expandInstanceReplicaInstances(c *Client, f *InstanceReplicaInstances) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}

	return m, nil
}

// flattenInstanceReplicaInstances flattens an instance of InstanceReplicaInstances from a JSON
// response object.
func flattenInstanceReplicaInstances(c *Client, i interface{}) *InstanceReplicaInstances {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaInstances{}
	r.Name = dcl.FlattenString(m["name"])
	r.Region = dcl.FlattenString(m["region"])

	return r
}

// expandInstanceServerCaCertMap expands the contents of InstanceServerCaCert into a JSON
// request object.
func expandInstanceServerCaCertMap(c *Client, f map[string]InstanceServerCaCert) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceServerCaCert(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceServerCaCertSlice expands the contents of InstanceServerCaCert into a JSON
// request object.
func expandInstanceServerCaCertSlice(c *Client, f []InstanceServerCaCert) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceServerCaCert(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceServerCaCertMap flattens the contents of InstanceServerCaCert from a JSON
// response object.
func flattenInstanceServerCaCertMap(c *Client, i interface{}) map[string]InstanceServerCaCert {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceServerCaCert{}
	}

	if len(a) == 0 {
		return map[string]InstanceServerCaCert{}
	}

	items := make(map[string]InstanceServerCaCert)
	for k, item := range a {
		items[k] = *flattenInstanceServerCaCert(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceServerCaCertSlice flattens the contents of InstanceServerCaCert from a JSON
// response object.
func flattenInstanceServerCaCertSlice(c *Client, i interface{}) []InstanceServerCaCert {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceServerCaCert{}
	}

	if len(a) == 0 {
		return []InstanceServerCaCert{}
	}

	items := make([]InstanceServerCaCert, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceServerCaCert(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceServerCaCert expands an instance of InstanceServerCaCert into a JSON
// request object.
func expandInstanceServerCaCert(c *Client, f *InstanceServerCaCert) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.CertSerialNumber; !dcl.IsEmptyValueIndirect(v) {
		m["certSerialNumber"] = v
	}
	if v := f.Cert; !dcl.IsEmptyValueIndirect(v) {
		m["cert"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.CommonName; !dcl.IsEmptyValueIndirect(v) {
		m["commonName"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		m["expirationTime"] = v
	}
	if v := f.Sha1Fingerprint; !dcl.IsEmptyValueIndirect(v) {
		m["sha1Fingerprint"] = v
	}
	if v := f.Instance; !dcl.IsEmptyValueIndirect(v) {
		m["instance"] = v
	}

	return m, nil
}

// flattenInstanceServerCaCert flattens an instance of InstanceServerCaCert from a JSON
// response object.
func flattenInstanceServerCaCert(c *Client, i interface{}) *InstanceServerCaCert {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceServerCaCert{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.CertSerialNumber = dcl.FlattenString(m["certSerialNumber"])
	r.Cert = dcl.FlattenString(m["cert"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.CommonName = dcl.FlattenString(m["commonName"])
	r.ExpirationTime = dcl.FlattenString(m["expirationTime"])
	r.Sha1Fingerprint = dcl.FlattenString(m["sha1Fingerprint"])
	r.Instance = dcl.FlattenString(m["instance"])

	return r
}

// expandInstanceOnPremisesConfigurationMap expands the contents of InstanceOnPremisesConfiguration into a JSON
// request object.
func expandInstanceOnPremisesConfigurationMap(c *Client, f map[string]InstanceOnPremisesConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceOnPremisesConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceOnPremisesConfigurationSlice expands the contents of InstanceOnPremisesConfiguration into a JSON
// request object.
func expandInstanceOnPremisesConfigurationSlice(c *Client, f []InstanceOnPremisesConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceOnPremisesConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceOnPremisesConfigurationMap flattens the contents of InstanceOnPremisesConfiguration from a JSON
// response object.
func flattenInstanceOnPremisesConfigurationMap(c *Client, i interface{}) map[string]InstanceOnPremisesConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceOnPremisesConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceOnPremisesConfiguration{}
	}

	items := make(map[string]InstanceOnPremisesConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceOnPremisesConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceOnPremisesConfigurationSlice flattens the contents of InstanceOnPremisesConfiguration from a JSON
// response object.
func flattenInstanceOnPremisesConfigurationSlice(c *Client, i interface{}) []InstanceOnPremisesConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceOnPremisesConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceOnPremisesConfiguration{}
	}

	items := make([]InstanceOnPremisesConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceOnPremisesConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceOnPremisesConfiguration expands an instance of InstanceOnPremisesConfiguration into a JSON
// request object.
func expandInstanceOnPremisesConfiguration(c *Client, f *InstanceOnPremisesConfiguration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HostPort; !dcl.IsEmptyValueIndirect(v) {
		m["hostPort"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}
	if v := f.CaCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["caCertificate"] = v
	}
	if v := f.ClientCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["clientCertificate"] = v
	}
	if v := f.ClientKey; !dcl.IsEmptyValueIndirect(v) {
		m["clientKey"] = v
	}
	if v := f.DumpFilePath; !dcl.IsEmptyValueIndirect(v) {
		m["dumpFilePath"] = v
	}
	if v := f.Database; !dcl.IsEmptyValueIndirect(v) {
		m["database"] = v
	}
	if v := f.ReplicatedDatabases; !dcl.IsEmptyValueIndirect(v) {
		m["replicatedDatabases"] = v
	}

	return m, nil
}

// flattenInstanceOnPremisesConfiguration flattens an instance of InstanceOnPremisesConfiguration from a JSON
// response object.
func flattenInstanceOnPremisesConfiguration(c *Client, i interface{}) *InstanceOnPremisesConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceOnPremisesConfiguration{}
	r.HostPort = dcl.FlattenString(m["hostPort"])
	r.Kind = dcl.FlattenString(m["kind"])
	r.Username = dcl.FlattenString(m["username"])
	r.Password = dcl.FlattenString(m["password"])
	r.CaCertificate = dcl.FlattenString(m["caCertificate"])
	r.ClientCertificate = dcl.FlattenString(m["clientCertificate"])
	r.ClientKey = dcl.FlattenString(m["clientKey"])
	r.DumpFilePath = dcl.FlattenString(m["dumpFilePath"])
	r.Database = dcl.FlattenString(m["database"])
	r.ReplicatedDatabases = dcl.FlattenStringSlice(m["replicatedDatabases"])

	return r
}

// expandInstanceDiskEncryptionStatusMap expands the contents of InstanceDiskEncryptionStatus into a JSON
// request object.
func expandInstanceDiskEncryptionStatusMap(c *Client, f map[string]InstanceDiskEncryptionStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceDiskEncryptionStatus(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceDiskEncryptionStatusSlice expands the contents of InstanceDiskEncryptionStatus into a JSON
// request object.
func expandInstanceDiskEncryptionStatusSlice(c *Client, f []InstanceDiskEncryptionStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceDiskEncryptionStatus(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceDiskEncryptionStatusMap flattens the contents of InstanceDiskEncryptionStatus from a JSON
// response object.
func flattenInstanceDiskEncryptionStatusMap(c *Client, i interface{}) map[string]InstanceDiskEncryptionStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceDiskEncryptionStatus{}
	}

	if len(a) == 0 {
		return map[string]InstanceDiskEncryptionStatus{}
	}

	items := make(map[string]InstanceDiskEncryptionStatus)
	for k, item := range a {
		items[k] = *flattenInstanceDiskEncryptionStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceDiskEncryptionStatusSlice flattens the contents of InstanceDiskEncryptionStatus from a JSON
// response object.
func flattenInstanceDiskEncryptionStatusSlice(c *Client, i interface{}) []InstanceDiskEncryptionStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDiskEncryptionStatus{}
	}

	if len(a) == 0 {
		return []InstanceDiskEncryptionStatus{}
	}

	items := make([]InstanceDiskEncryptionStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDiskEncryptionStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceDiskEncryptionStatus expands an instance of InstanceDiskEncryptionStatus into a JSON
// request object.
func expandInstanceDiskEncryptionStatus(c *Client, f *InstanceDiskEncryptionStatus) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyVersionName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyVersionName"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceDiskEncryptionStatus flattens an instance of InstanceDiskEncryptionStatus from a JSON
// response object.
func flattenInstanceDiskEncryptionStatus(c *Client, i interface{}) *InstanceDiskEncryptionStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceDiskEncryptionStatus{}
	r.KmsKeyVersionName = dcl.FlattenString(m["kmsKeyVersionName"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// flattenInstanceBackendTypeEnumSlice flattens the contents of InstanceBackendTypeEnum from a JSON
// response object.
func flattenInstanceBackendTypeEnumSlice(c *Client, i interface{}) []InstanceBackendTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceBackendTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceBackendTypeEnum{}
	}

	items := make([]InstanceBackendTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceBackendTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceBackendTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceBackendTypeEnum with the same value as that string.
func flattenInstanceBackendTypeEnum(i interface{}) *InstanceBackendTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceBackendTypeEnumRef("")
	}

	return InstanceBackendTypeEnumRef(s)
}

// flattenInstanceDatabaseVersionEnumSlice flattens the contents of InstanceDatabaseVersionEnum from a JSON
// response object.
func flattenInstanceDatabaseVersionEnumSlice(c *Client, i interface{}) []InstanceDatabaseVersionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDatabaseVersionEnum{}
	}

	if len(a) == 0 {
		return []InstanceDatabaseVersionEnum{}
	}

	items := make([]InstanceDatabaseVersionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDatabaseVersionEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceDatabaseVersionEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceDatabaseVersionEnum with the same value as that string.
func flattenInstanceDatabaseVersionEnum(i interface{}) *InstanceDatabaseVersionEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceDatabaseVersionEnumRef("")
	}

	return InstanceDatabaseVersionEnumRef(s)
}

// flattenInstanceInstanceTypeEnumSlice flattens the contents of InstanceInstanceTypeEnum from a JSON
// response object.
func flattenInstanceInstanceTypeEnumSlice(c *Client, i interface{}) []InstanceInstanceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceInstanceTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceInstanceTypeEnum{}
	}

	items := make([]InstanceInstanceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceInstanceTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceInstanceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceInstanceTypeEnum with the same value as that string.
func flattenInstanceInstanceTypeEnum(i interface{}) *InstanceInstanceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceInstanceTypeEnumRef("")
	}

	return InstanceInstanceTypeEnumRef(s)
}

// flattenInstanceIPAddressesTypeEnumSlice flattens the contents of InstanceIPAddressesTypeEnum from a JSON
// response object.
func flattenInstanceIPAddressesTypeEnumSlice(c *Client, i interface{}) []InstanceIPAddressesTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceIPAddressesTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceIPAddressesTypeEnum{}
	}

	items := make([]InstanceIPAddressesTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceIPAddressesTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceIPAddressesTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceIPAddressesTypeEnum with the same value as that string.
func flattenInstanceIPAddressesTypeEnum(i interface{}) *InstanceIPAddressesTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceIPAddressesTypeEnumRef("")
	}

	return InstanceIPAddressesTypeEnumRef(s)
}

// flattenInstanceSettingsAvailabilityTypeEnumSlice flattens the contents of InstanceSettingsAvailabilityTypeEnum from a JSON
// response object.
func flattenInstanceSettingsAvailabilityTypeEnumSlice(c *Client, i interface{}) []InstanceSettingsAvailabilityTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsAvailabilityTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsAvailabilityTypeEnum{}
	}

	items := make([]InstanceSettingsAvailabilityTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsAvailabilityTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceSettingsAvailabilityTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsAvailabilityTypeEnum with the same value as that string.
func flattenInstanceSettingsAvailabilityTypeEnum(i interface{}) *InstanceSettingsAvailabilityTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsAvailabilityTypeEnumRef("")
	}

	return InstanceSettingsAvailabilityTypeEnumRef(s)
}

// flattenInstanceSettingsPricingPlanEnumSlice flattens the contents of InstanceSettingsPricingPlanEnum from a JSON
// response object.
func flattenInstanceSettingsPricingPlanEnumSlice(c *Client, i interface{}) []InstanceSettingsPricingPlanEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsPricingPlanEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsPricingPlanEnum{}
	}

	items := make([]InstanceSettingsPricingPlanEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsPricingPlanEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceSettingsPricingPlanEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsPricingPlanEnum with the same value as that string.
func flattenInstanceSettingsPricingPlanEnum(i interface{}) *InstanceSettingsPricingPlanEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsPricingPlanEnumRef("")
	}

	return InstanceSettingsPricingPlanEnumRef(s)
}

// flattenInstanceSettingsReplicationTypeEnumSlice flattens the contents of InstanceSettingsReplicationTypeEnum from a JSON
// response object.
func flattenInstanceSettingsReplicationTypeEnumSlice(c *Client, i interface{}) []InstanceSettingsReplicationTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsReplicationTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsReplicationTypeEnum{}
	}

	items := make([]InstanceSettingsReplicationTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsReplicationTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceSettingsReplicationTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsReplicationTypeEnum with the same value as that string.
func flattenInstanceSettingsReplicationTypeEnum(i interface{}) *InstanceSettingsReplicationTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsReplicationTypeEnumRef("")
	}

	return InstanceSettingsReplicationTypeEnumRef(s)
}

// flattenInstanceSettingsActivationPolicyEnumSlice flattens the contents of InstanceSettingsActivationPolicyEnum from a JSON
// response object.
func flattenInstanceSettingsActivationPolicyEnumSlice(c *Client, i interface{}) []InstanceSettingsActivationPolicyEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsActivationPolicyEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsActivationPolicyEnum{}
	}

	items := make([]InstanceSettingsActivationPolicyEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsActivationPolicyEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceSettingsActivationPolicyEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsActivationPolicyEnum with the same value as that string.
func flattenInstanceSettingsActivationPolicyEnum(i interface{}) *InstanceSettingsActivationPolicyEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsActivationPolicyEnumRef("")
	}

	return InstanceSettingsActivationPolicyEnumRef(s)
}

// flattenInstanceSettingsDataDiskTypeEnumSlice flattens the contents of InstanceSettingsDataDiskTypeEnum from a JSON
// response object.
func flattenInstanceSettingsDataDiskTypeEnumSlice(c *Client, i interface{}) []InstanceSettingsDataDiskTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsDataDiskTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsDataDiskTypeEnum{}
	}

	items := make([]InstanceSettingsDataDiskTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsDataDiskTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceSettingsDataDiskTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsDataDiskTypeEnum with the same value as that string.
func flattenInstanceSettingsDataDiskTypeEnum(i interface{}) *InstanceSettingsDataDiskTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsDataDiskTypeEnumRef("")
	}

	return InstanceSettingsDataDiskTypeEnumRef(s)
}

// flattenInstanceSettingsMaintenanceWindowUpdateTrackEnumSlice flattens the contents of InstanceSettingsMaintenanceWindowUpdateTrackEnum from a JSON
// response object.
func flattenInstanceSettingsMaintenanceWindowUpdateTrackEnumSlice(c *Client, i interface{}) []InstanceSettingsMaintenanceWindowUpdateTrackEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsMaintenanceWindowUpdateTrackEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsMaintenanceWindowUpdateTrackEnum{}
	}

	items := make([]InstanceSettingsMaintenanceWindowUpdateTrackEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsMaintenanceWindowUpdateTrackEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceSettingsMaintenanceWindowUpdateTrackEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsMaintenanceWindowUpdateTrackEnum with the same value as that string.
func flattenInstanceSettingsMaintenanceWindowUpdateTrackEnum(i interface{}) *InstanceSettingsMaintenanceWindowUpdateTrackEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsMaintenanceWindowUpdateTrackEnumRef("")
	}

	return InstanceSettingsMaintenanceWindowUpdateTrackEnumRef(s)
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumSlice flattens the contents of InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumSlice(c *Client, i interface{}) []InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum{}
	}

	items := make([]InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum with the same value as that string.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(i interface{}) *InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumRef("")
	}

	return InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Instance) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInstance(b, c)
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
