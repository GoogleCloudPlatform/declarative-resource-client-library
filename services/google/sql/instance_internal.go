// Copyright 2020 Google LLC. All Rights Reserved.
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
package sql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mohae/deepcopy"
	"io/ioutil"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"reflect"
)

func (r *Instance) validate() error {

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
	return nil
}

func instanceGetURL(userBasePath string, r *Instance) string {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params)
}

func instanceListURL(userBasePath, project string) string {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/instances/", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params)
}

func instanceCreateURL(userBasePath, project string) string {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/instances", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params)
}

func instanceDeleteURL(userBasePath string, r *Instance) string {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params)
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

	var o Operation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInstanceRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u := instanceListURL(c.Config.BasePath, project)
	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InstanceMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err := dcl.AddQueryParams(u, m)
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
	for _, res := range resources {
		if f(res) {
			err := c.DeleteInstance(ctx, res)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

	u := instanceDeleteURL(c.Config.BasePath, r.urlNormalized())

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o Operation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config); err != nil {
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
	u := instanceCreateURL(c.Config.BasePath, project)

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o Operation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config); err != nil {
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
	u := instanceGetURL(c.Config.BasePath, r.urlNormalized())
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
	desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial)
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

func canonicalizeInstanceDesiredState(rawDesired, rawInitial *Instance) (*Instance, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MaxDiskSize = canonicalizeInstanceMaxDiskSize(rawDesired.MaxDiskSize, nil)
		rawDesired.CurrentDiskSize = canonicalizeInstanceCurrentDiskSize(rawDesired.CurrentDiskSize, nil)
		rawDesired.DiskEncryptionConfiguration = canonicalizeInstanceDiskEncryptionConfiguration(rawDesired.DiskEncryptionConfiguration, nil)
		rawDesired.FailoverReplica = canonicalizeInstanceFailoverReplica(rawDesired.FailoverReplica, nil)
		rawDesired.MasterInstance = canonicalizeInstanceMasterInstance(rawDesired.MasterInstance, nil)
		rawDesired.ReplicaConfiguration = canonicalizeInstanceReplicaConfiguration(rawDesired.ReplicaConfiguration, nil)
		rawDesired.ScheduledMaintenance = canonicalizeInstanceScheduledMaintenance(rawDesired.ScheduledMaintenance, nil)
		rawDesired.Settings = canonicalizeInstanceSettings(rawDesired.Settings, nil)

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
	rawDesired.MaxDiskSize = canonicalizeInstanceMaxDiskSize(rawDesired.MaxDiskSize, rawInitial.MaxDiskSize)
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
	rawDesired.CurrentDiskSize = canonicalizeInstanceCurrentDiskSize(rawDesired.CurrentDiskSize, rawInitial.CurrentDiskSize)
	rawDesired.DiskEncryptionConfiguration = canonicalizeInstanceDiskEncryptionConfiguration(rawDesired.DiskEncryptionConfiguration, rawInitial.DiskEncryptionConfiguration)
	rawDesired.FailoverReplica = canonicalizeInstanceFailoverReplica(rawDesired.FailoverReplica, rawInitial.FailoverReplica)
	if dcl.IsZeroValue(rawDesired.IPAddresses) {
		rawDesired.IPAddresses = rawInitial.IPAddresses
	}
	rawDesired.MasterInstance = canonicalizeInstanceMasterInstance(rawDesired.MasterInstance, rawInitial.MasterInstance)
	rawDesired.ReplicaConfiguration = canonicalizeInstanceReplicaConfiguration(rawDesired.ReplicaConfiguration, rawInitial.ReplicaConfiguration)
	rawDesired.ScheduledMaintenance = canonicalizeInstanceScheduledMaintenance(rawDesired.ScheduledMaintenance, rawInitial.ScheduledMaintenance)
	rawDesired.Settings = canonicalizeInstanceSettings(rawDesired.Settings, rawInitial.Settings)

	return rawDesired, nil
}

func canonicalizeInstanceNewState(rawNew, rawDesired *Instance) (*Instance, error) {

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
		rawNew.MaxDiskSize = canonicalizeNewInstanceMaxDiskSize(rawDesired.MaxDiskSize, rawNew.MaxDiskSize)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.NameToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

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
		rawNew.CurrentDiskSize = canonicalizeNewInstanceCurrentDiskSize(rawDesired.CurrentDiskSize, rawNew.CurrentDiskSize)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskEncryptionConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.DiskEncryptionConfiguration) {
		rawNew.DiskEncryptionConfiguration = rawDesired.DiskEncryptionConfiguration
	} else {
		rawNew.DiskEncryptionConfiguration = canonicalizeNewInstanceDiskEncryptionConfiguration(rawDesired.DiskEncryptionConfiguration, rawNew.DiskEncryptionConfiguration)
	}

	if dcl.IsEmptyValueIndirect(rawNew.FailoverReplica) && dcl.IsEmptyValueIndirect(rawDesired.FailoverReplica) {
		rawNew.FailoverReplica = rawDesired.FailoverReplica
	} else {
		rawNew.FailoverReplica = canonicalizeNewInstanceFailoverReplica(rawDesired.FailoverReplica, rawNew.FailoverReplica)
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPAddresses) && dcl.IsEmptyValueIndirect(rawDesired.IPAddresses) {
		rawNew.IPAddresses = rawDesired.IPAddresses
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MasterInstance) && dcl.IsEmptyValueIndirect(rawDesired.MasterInstance) {
		rawNew.MasterInstance = rawDesired.MasterInstance
	} else {
		rawNew.MasterInstance = canonicalizeNewInstanceMasterInstance(rawDesired.MasterInstance, rawNew.MasterInstance)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReplicaConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.ReplicaConfiguration) {
		rawNew.ReplicaConfiguration = rawDesired.ReplicaConfiguration
	} else {
		rawNew.ReplicaConfiguration = canonicalizeNewInstanceReplicaConfiguration(rawDesired.ReplicaConfiguration, rawNew.ReplicaConfiguration)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ScheduledMaintenance) && dcl.IsEmptyValueIndirect(rawDesired.ScheduledMaintenance) {
		rawNew.ScheduledMaintenance = rawDesired.ScheduledMaintenance
	} else {
		rawNew.ScheduledMaintenance = canonicalizeNewInstanceScheduledMaintenance(rawDesired.ScheduledMaintenance, rawNew.ScheduledMaintenance)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Settings) && dcl.IsEmptyValueIndirect(rawDesired.Settings) {
		rawNew.Settings = rawDesired.Settings
	} else {
		rawNew.Settings = canonicalizeNewInstanceSettings(rawDesired.Settings, rawNew.Settings)
	}

	return rawNew, nil
}

func canonicalizeInstanceMaxDiskSize(des, initial *InstanceMaxDiskSize) *InstanceMaxDiskSize {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceMaxDiskSize(des, nw *InstanceMaxDiskSize) *InstanceMaxDiskSize {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceCurrentDiskSize(des, initial *InstanceCurrentDiskSize) *InstanceCurrentDiskSize {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceCurrentDiskSize(des, nw *InstanceCurrentDiskSize) *InstanceCurrentDiskSize {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceDiskEncryptionConfiguration(des, initial *InstanceDiskEncryptionConfiguration) *InstanceDiskEncryptionConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewInstanceDiskEncryptionConfiguration(des, nw *InstanceDiskEncryptionConfiguration) *InstanceDiskEncryptionConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceFailoverReplica(des, initial *InstanceFailoverReplica) *InstanceFailoverReplica {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

	return des
}

func canonicalizeNewInstanceFailoverReplica(des, nw *InstanceFailoverReplica) *InstanceFailoverReplica {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceIPAddresses(des, initial *InstanceIPAddresses) *InstanceIPAddresses {
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
	if dcl.IsZeroValue(des.IPAddress) {
		des.IPAddress = initial.IPAddress
	}
	des.TimeToRetire = canonicalizeInstanceIPAddressesTimeToRetire(des.TimeToRetire, initial.TimeToRetire)

	return des
}

func canonicalizeNewInstanceIPAddresses(des, nw *InstanceIPAddresses) *InstanceIPAddresses {
	if des == nil || nw == nil {
		return nw
	}

	nw.TimeToRetire = canonicalizeNewInstanceIPAddressesTimeToRetire(des.TimeToRetire, nw.TimeToRetire)

	return nw
}

func canonicalizeInstanceIPAddressesTimeToRetire(des, initial *InstanceIPAddressesTimeToRetire) *InstanceIPAddressesTimeToRetire {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewInstanceIPAddressesTimeToRetire(des, nw *InstanceIPAddressesTimeToRetire) *InstanceIPAddressesTimeToRetire {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceMasterInstance(des, initial *InstanceMasterInstance) *InstanceMasterInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewInstanceMasterInstance(des, nw *InstanceMasterInstance) *InstanceMasterInstance {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceReplicaConfiguration(des, initial *InstanceReplicaConfiguration) *InstanceReplicaConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	des.MysqlReplicaConfiguration = canonicalizeInstanceReplicaConfigurationMysqlReplicaConfiguration(des.MysqlReplicaConfiguration, initial.MysqlReplicaConfiguration)
	if dcl.IsZeroValue(des.FailoverTarget) {
		des.FailoverTarget = initial.FailoverTarget
	}
	des.ReplicaPoolConfiguration = canonicalizeInstanceReplicaConfigurationReplicaPoolConfiguration(des.ReplicaPoolConfiguration, initial.ReplicaPoolConfiguration)

	return des
}

func canonicalizeNewInstanceReplicaConfiguration(des, nw *InstanceReplicaConfiguration) *InstanceReplicaConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	nw.MysqlReplicaConfiguration = canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfiguration(des.MysqlReplicaConfiguration, nw.MysqlReplicaConfiguration)
	nw.ReplicaPoolConfiguration = canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfiguration(des.ReplicaPoolConfiguration, nw.ReplicaPoolConfiguration)

	return nw
}

func canonicalizeInstanceReplicaConfigurationMysqlReplicaConfiguration(des, initial *InstanceReplicaConfigurationMysqlReplicaConfiguration) *InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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
	des.MasterHeartbeatPeriod = canonicalizeInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(des.MasterHeartbeatPeriod, initial.MasterHeartbeatPeriod)
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

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfiguration(des, nw *InstanceReplicaConfigurationMysqlReplicaConfiguration) *InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	nw.MasterHeartbeatPeriod = canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(des.MasterHeartbeatPeriod, nw.MasterHeartbeatPeriod)

	return nw
}

func canonicalizeInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(des, initial *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(des, nw *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceReplicaConfigurationReplicaPoolConfiguration(des, initial *InstanceReplicaConfigurationReplicaPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	des.StaticPoolConfiguration = canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(des.StaticPoolConfiguration, initial.StaticPoolConfiguration)
	des.AutoscalingPoolConfiguration = canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(des.AutoscalingPoolConfiguration, initial.AutoscalingPoolConfiguration)

	return des
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfiguration(des, nw *InstanceReplicaConfigurationReplicaPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	nw.StaticPoolConfiguration = canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(des.StaticPoolConfiguration, nw.StaticPoolConfiguration)
	nw.AutoscalingPoolConfiguration = canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(des.AutoscalingPoolConfiguration, nw.AutoscalingPoolConfiguration)

	return nw
}

func canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(des, initial *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(des, nw *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(des, initial *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(des, nw *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceScheduledMaintenance(des, initial *InstanceScheduledMaintenance) *InstanceScheduledMaintenance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.StartTime = canonicalizeInstanceScheduledMaintenanceStartTime(des.StartTime, initial.StartTime)
	if dcl.IsZeroValue(des.CanDefer) {
		des.CanDefer = initial.CanDefer
	}
	if dcl.IsZeroValue(des.CanReschedule) {
		des.CanReschedule = initial.CanReschedule
	}

	return des
}

func canonicalizeNewInstanceScheduledMaintenance(des, nw *InstanceScheduledMaintenance) *InstanceScheduledMaintenance {
	if des == nil || nw == nil {
		return nw
	}

	nw.StartTime = canonicalizeNewInstanceScheduledMaintenanceStartTime(des.StartTime, nw.StartTime)

	return nw
}

func canonicalizeInstanceScheduledMaintenanceStartTime(des, initial *InstanceScheduledMaintenanceStartTime) *InstanceScheduledMaintenanceStartTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

func canonicalizeNewInstanceScheduledMaintenanceStartTime(des, nw *InstanceScheduledMaintenanceStartTime) *InstanceScheduledMaintenanceStartTime {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeInstanceSettings(des, initial *InstanceSettings) *InstanceSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
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

	return des
}

func canonicalizeNewInstanceSettings(des, nw *InstanceSettings) *InstanceSettings {
	if des == nil || nw == nil {
		return nw
	}

	return nw
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
		c.Config.Logger.Infof("Detected diff in BackendType.\nDESIRED: %#v\nACTUAL: %#v", desired.BackendType, actual.BackendType)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "BackendType",
		})
	}
	if !dcl.IsZeroValue(desired.ConnectionName) && (dcl.IsZeroValue(actual.ConnectionName) || !reflect.DeepEqual(*desired.ConnectionName, *actual.ConnectionName)) {
		c.Config.Logger.Infof("Detected diff in ConnectionName.\nDESIRED: %#v\nACTUAL: %#v", desired.ConnectionName, actual.ConnectionName)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "ConnectionName",
		})
	}
	if !dcl.IsZeroValue(desired.DatabaseVersion) && (dcl.IsZeroValue(actual.DatabaseVersion) || !reflect.DeepEqual(*desired.DatabaseVersion, *actual.DatabaseVersion)) {
		c.Config.Logger.Infof("Detected diff in DatabaseVersion.\nDESIRED: %#v\nACTUAL: %#v", desired.DatabaseVersion, actual.DatabaseVersion)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "DatabaseVersion",
		})
	}
	if !dcl.IsZeroValue(desired.Etag) && (dcl.IsZeroValue(actual.Etag) || !reflect.DeepEqual(*desired.Etag, *actual.Etag)) {
		c.Config.Logger.Infof("Detected diff in Etag.\nDESIRED: %#v\nACTUAL: %#v", desired.Etag, actual.Etag)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Etag",
		})
	}
	if !dcl.IsZeroValue(desired.GceZone) && (dcl.IsZeroValue(actual.GceZone) || !reflect.DeepEqual(*desired.GceZone, *actual.GceZone)) {
		c.Config.Logger.Infof("Detected diff in GceZone.\nDESIRED: %#v\nACTUAL: %#v", desired.GceZone, actual.GceZone)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "GceZone",
		})
	}
	if !dcl.IsZeroValue(desired.InstanceType) && (dcl.IsZeroValue(actual.InstanceType) || !reflect.DeepEqual(*desired.InstanceType, *actual.InstanceType)) {
		c.Config.Logger.Infof("Detected diff in InstanceType.\nDESIRED: %#v\nACTUAL: %#v", desired.InstanceType, actual.InstanceType)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "InstanceType",
		})
	}
	if !dcl.IsZeroValue(desired.MasterInstanceName) && (dcl.IsZeroValue(actual.MasterInstanceName) || !reflect.DeepEqual(*desired.MasterInstanceName, *actual.MasterInstanceName)) {
		c.Config.Logger.Infof("Detected diff in MasterInstanceName.\nDESIRED: %#v\nACTUAL: %#v", desired.MasterInstanceName, actual.MasterInstanceName)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "MasterInstanceName",
		})
	}
	if compareInstanceMaxDiskSize(c, desired.MaxDiskSize, actual.MaxDiskSize) {
		c.Config.Logger.Infof("Detected diff in MaxDiskSize.\nDESIRED: %#v\nACTUAL: %#v", desired.MaxDiskSize, actual.MaxDiskSize)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "MaxDiskSize",
		})
	}
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %#v\nACTUAL: %#v", desired.Name, actual.Name)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Project) && !dcl.NameToSelfLink(desired.Project, actual.Project) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %#v\nACTUAL: %#v", desired.Project, actual.Project)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Project",
		})
	}
	if !dcl.IsZeroValue(desired.Region) && (dcl.IsZeroValue(actual.Region) || !reflect.DeepEqual(*desired.Region, *actual.Region)) {
		c.Config.Logger.Infof("Detected diff in Region.\nDESIRED: %#v\nACTUAL: %#v", desired.Region, actual.Region)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Region",
		})
	}
	if !dcl.IsZeroValue(desired.RootPassword) && (dcl.IsZeroValue(actual.RootPassword) || !reflect.DeepEqual(*desired.RootPassword, *actual.RootPassword)) {
		c.Config.Logger.Infof("Detected diff in RootPassword.\nDESIRED: %#v\nACTUAL: %#v", desired.RootPassword, actual.RootPassword)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "RootPassword",
		})
	}
	if compareInstanceCurrentDiskSize(c, desired.CurrentDiskSize, actual.CurrentDiskSize) {
		c.Config.Logger.Infof("Detected diff in CurrentDiskSize.\nDESIRED: %#v\nACTUAL: %#v", desired.CurrentDiskSize, actual.CurrentDiskSize)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "CurrentDiskSize",
		})
	}
	if compareInstanceDiskEncryptionConfiguration(c, desired.DiskEncryptionConfiguration, actual.DiskEncryptionConfiguration) {
		c.Config.Logger.Infof("Detected diff in DiskEncryptionConfiguration.\nDESIRED: %#v\nACTUAL: %#v", desired.DiskEncryptionConfiguration, actual.DiskEncryptionConfiguration)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "DiskEncryptionConfiguration",
		})
	}
	if compareInstanceFailoverReplica(c, desired.FailoverReplica, actual.FailoverReplica) {
		c.Config.Logger.Infof("Detected diff in FailoverReplica.\nDESIRED: %#v\nACTUAL: %#v", desired.FailoverReplica, actual.FailoverReplica)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "FailoverReplica",
		})
	}
	if compareInstanceIPAddressesSlice(c, desired.IPAddresses, actual.IPAddresses) {
		c.Config.Logger.Infof("Detected diff in IPAddresses.\nDESIRED: %#v\nACTUAL: %#v", desired.IPAddresses, actual.IPAddresses)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "IPAddresses",
		})
	}
	if compareInstanceMasterInstance(c, desired.MasterInstance, actual.MasterInstance) {
		c.Config.Logger.Infof("Detected diff in MasterInstance.\nDESIRED: %#v\nACTUAL: %#v", desired.MasterInstance, actual.MasterInstance)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "MasterInstance",
		})
	}
	if compareInstanceReplicaConfiguration(c, desired.ReplicaConfiguration, actual.ReplicaConfiguration) {
		c.Config.Logger.Infof("Detected diff in ReplicaConfiguration.\nDESIRED: %#v\nACTUAL: %#v", desired.ReplicaConfiguration, actual.ReplicaConfiguration)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "ReplicaConfiguration",
		})
	}
	if compareInstanceScheduledMaintenance(c, desired.ScheduledMaintenance, actual.ScheduledMaintenance) {
		c.Config.Logger.Infof("Detected diff in ScheduledMaintenance.\nDESIRED: %#v\nACTUAL: %#v", desired.ScheduledMaintenance, actual.ScheduledMaintenance)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "ScheduledMaintenance",
		})
	}
	if compareInstanceSettings(c, desired.Settings, actual.Settings) {
		c.Config.Logger.Infof("Detected diff in Settings.\nDESIRED: %#v\nACTUAL: %#v", desired.Settings, actual.Settings)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Settings",
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
		if compareInstanceMaxDiskSize(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
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
		if compareInstanceCurrentDiskSize(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
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
		if compareInstanceDiskEncryptionConfiguration(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.KmsKeyName, actual.KmsKeyName) && !dcl.IsZeroValue(desired.KmsKeyName) {
		c.Config.Logger.Infof("Diff in KmsKeyName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyName), dcl.SprintResource(actual.KmsKeyName))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
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
		if compareInstanceFailoverReplica(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Available == nil && desired.Available != nil && !dcl.IsEmptyValueIndirect(desired.Available) {
		return true
	}
	if !reflect.DeepEqual(desired.Available, actual.Available) && !dcl.IsZeroValue(desired.Available) {
		c.Config.Logger.Infof("Diff in Available. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Available), dcl.SprintResource(actual.Available))
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
		if compareInstanceIPAddresses(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	if actual.IPAddress == nil && desired.IPAddress != nil && !dcl.IsEmptyValueIndirect(desired.IPAddress) {
		return true
	}
	if !reflect.DeepEqual(desired.IPAddress, actual.IPAddress) && !dcl.IsZeroValue(desired.IPAddress) {
		c.Config.Logger.Infof("Diff in IPAddress. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPAddress), dcl.SprintResource(actual.IPAddress))
		return true
	}
	if actual.TimeToRetire == nil && desired.TimeToRetire != nil && !dcl.IsEmptyValueIndirect(desired.TimeToRetire) {
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
		if compareInstanceIPAddressesTimeToRetire(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
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
		if compareInstanceMasterInstance(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.Region == nil && desired.Region != nil && !dcl.IsEmptyValueIndirect(desired.Region) {
		return true
	}
	if !reflect.DeepEqual(desired.Region, actual.Region) && !dcl.IsZeroValue(desired.Region) {
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
		if compareInstanceReplicaConfiguration(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.MysqlReplicaConfiguration == nil && desired.MysqlReplicaConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.MysqlReplicaConfiguration) {
		return true
	}
	if compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c, desired.MysqlReplicaConfiguration, actual.MysqlReplicaConfiguration) && !dcl.IsZeroValue(desired.MysqlReplicaConfiguration) {
		c.Config.Logger.Infof("Diff in MysqlReplicaConfiguration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MysqlReplicaConfiguration), dcl.SprintResource(actual.MysqlReplicaConfiguration))
		return true
	}
	if actual.FailoverTarget == nil && desired.FailoverTarget != nil && !dcl.IsEmptyValueIndirect(desired.FailoverTarget) {
		return true
	}
	if !reflect.DeepEqual(desired.FailoverTarget, actual.FailoverTarget) && !dcl.IsZeroValue(desired.FailoverTarget) {
		c.Config.Logger.Infof("Diff in FailoverTarget. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FailoverTarget), dcl.SprintResource(actual.FailoverTarget))
		return true
	}
	if actual.ReplicaPoolConfiguration == nil && desired.ReplicaPoolConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.ReplicaPoolConfiguration) {
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
		if compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.DumpFilePath, actual.DumpFilePath) && !dcl.IsZeroValue(desired.DumpFilePath) {
		c.Config.Logger.Infof("Diff in DumpFilePath. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DumpFilePath), dcl.SprintResource(actual.DumpFilePath))
		return true
	}
	if actual.Username == nil && desired.Username != nil && !dcl.IsEmptyValueIndirect(desired.Username) {
		return true
	}
	if !reflect.DeepEqual(desired.Username, actual.Username) && !dcl.IsZeroValue(desired.Username) {
		c.Config.Logger.Infof("Diff in Username. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Username), dcl.SprintResource(actual.Username))
		return true
	}
	if actual.Password == nil && desired.Password != nil && !dcl.IsEmptyValueIndirect(desired.Password) {
		return true
	}
	if !reflect.DeepEqual(desired.Password, actual.Password) && !dcl.IsZeroValue(desired.Password) {
		c.Config.Logger.Infof("Diff in Password. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Password), dcl.SprintResource(actual.Password))
		return true
	}
	if actual.ConnectRetryInterval == nil && desired.ConnectRetryInterval != nil && !dcl.IsEmptyValueIndirect(desired.ConnectRetryInterval) {
		return true
	}
	if !reflect.DeepEqual(desired.ConnectRetryInterval, actual.ConnectRetryInterval) && !dcl.IsZeroValue(desired.ConnectRetryInterval) {
		c.Config.Logger.Infof("Diff in ConnectRetryInterval. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConnectRetryInterval), dcl.SprintResource(actual.ConnectRetryInterval))
		return true
	}
	if actual.MasterHeartbeatPeriod == nil && desired.MasterHeartbeatPeriod != nil && !dcl.IsEmptyValueIndirect(desired.MasterHeartbeatPeriod) {
		return true
	}
	if compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, desired.MasterHeartbeatPeriod, actual.MasterHeartbeatPeriod) && !dcl.IsZeroValue(desired.MasterHeartbeatPeriod) {
		c.Config.Logger.Infof("Diff in MasterHeartbeatPeriod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MasterHeartbeatPeriod), dcl.SprintResource(actual.MasterHeartbeatPeriod))
		return true
	}
	if actual.CaCertificate == nil && desired.CaCertificate != nil && !dcl.IsEmptyValueIndirect(desired.CaCertificate) {
		return true
	}
	if !reflect.DeepEqual(desired.CaCertificate, actual.CaCertificate) && !dcl.IsZeroValue(desired.CaCertificate) {
		c.Config.Logger.Infof("Diff in CaCertificate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CaCertificate), dcl.SprintResource(actual.CaCertificate))
		return true
	}
	if actual.ClientCertificate == nil && desired.ClientCertificate != nil && !dcl.IsEmptyValueIndirect(desired.ClientCertificate) {
		return true
	}
	if !reflect.DeepEqual(desired.ClientCertificate, actual.ClientCertificate) && !dcl.IsZeroValue(desired.ClientCertificate) {
		c.Config.Logger.Infof("Diff in ClientCertificate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientCertificate), dcl.SprintResource(actual.ClientCertificate))
		return true
	}
	if actual.ClientKey == nil && desired.ClientKey != nil && !dcl.IsEmptyValueIndirect(desired.ClientKey) {
		return true
	}
	if !reflect.DeepEqual(desired.ClientKey, actual.ClientKey) && !dcl.IsZeroValue(desired.ClientKey) {
		c.Config.Logger.Infof("Diff in ClientKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientKey), dcl.SprintResource(actual.ClientKey))
		return true
	}
	if actual.SslCipher == nil && desired.SslCipher != nil && !dcl.IsEmptyValueIndirect(desired.SslCipher) {
		return true
	}
	if !reflect.DeepEqual(desired.SslCipher, actual.SslCipher) && !dcl.IsZeroValue(desired.SslCipher) {
		c.Config.Logger.Infof("Diff in SslCipher. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SslCipher), dcl.SprintResource(actual.SslCipher))
		return true
	}
	if actual.VerifyServerCertificate == nil && desired.VerifyServerCertificate != nil && !dcl.IsEmptyValueIndirect(desired.VerifyServerCertificate) {
		return true
	}
	if !reflect.DeepEqual(desired.VerifyServerCertificate, actual.VerifyServerCertificate) && !dcl.IsZeroValue(desired.VerifyServerCertificate) {
		c.Config.Logger.Infof("Diff in VerifyServerCertificate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.VerifyServerCertificate), dcl.SprintResource(actual.VerifyServerCertificate))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
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
		if compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
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
		if compareInstanceReplicaConfigurationReplicaPoolConfiguration(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.StaticPoolConfiguration == nil && desired.StaticPoolConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.StaticPoolConfiguration) {
		return true
	}
	if compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, desired.StaticPoolConfiguration, actual.StaticPoolConfiguration) && !dcl.IsZeroValue(desired.StaticPoolConfiguration) {
		c.Config.Logger.Infof("Diff in StaticPoolConfiguration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StaticPoolConfiguration), dcl.SprintResource(actual.StaticPoolConfiguration))
		return true
	}
	if actual.AutoscalingPoolConfiguration == nil && desired.AutoscalingPoolConfiguration != nil && !dcl.IsEmptyValueIndirect(desired.AutoscalingPoolConfiguration) {
		return true
	}
	if compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, desired.AutoscalingPoolConfiguration, actual.AutoscalingPoolConfiguration) && !dcl.IsZeroValue(desired.AutoscalingPoolConfiguration) {
		c.Config.Logger.Infof("Diff in AutoscalingPoolConfiguration. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoscalingPoolConfiguration), dcl.SprintResource(actual.AutoscalingPoolConfiguration))
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
		if compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.ReplicaCount == nil && desired.ReplicaCount != nil && !dcl.IsEmptyValueIndirect(desired.ReplicaCount) {
		return true
	}
	if !reflect.DeepEqual(desired.ReplicaCount, actual.ReplicaCount) && !dcl.IsZeroValue(desired.ReplicaCount) {
		c.Config.Logger.Infof("Diff in ReplicaCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicaCount), dcl.SprintResource(actual.ReplicaCount))
		return true
	}
	if actual.ExposeReplicaIP == nil && desired.ExposeReplicaIP != nil && !dcl.IsEmptyValueIndirect(desired.ExposeReplicaIP) {
		return true
	}
	if !reflect.DeepEqual(desired.ExposeReplicaIP, actual.ExposeReplicaIP) && !dcl.IsZeroValue(desired.ExposeReplicaIP) {
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
		if compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.MinReplicaCount == nil && desired.MinReplicaCount != nil && !dcl.IsEmptyValueIndirect(desired.MinReplicaCount) {
		return true
	}
	if !reflect.DeepEqual(desired.MinReplicaCount, actual.MinReplicaCount) && !dcl.IsZeroValue(desired.MinReplicaCount) {
		c.Config.Logger.Infof("Diff in MinReplicaCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinReplicaCount), dcl.SprintResource(actual.MinReplicaCount))
		return true
	}
	if actual.MaxReplicaCount == nil && desired.MaxReplicaCount != nil && !dcl.IsEmptyValueIndirect(desired.MaxReplicaCount) {
		return true
	}
	if !reflect.DeepEqual(desired.MaxReplicaCount, actual.MaxReplicaCount) && !dcl.IsZeroValue(desired.MaxReplicaCount) {
		c.Config.Logger.Infof("Diff in MaxReplicaCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxReplicaCount), dcl.SprintResource(actual.MaxReplicaCount))
		return true
	}
	if actual.TargetCpuUtil == nil && desired.TargetCpuUtil != nil && !dcl.IsEmptyValueIndirect(desired.TargetCpuUtil) {
		return true
	}
	if !reflect.DeepEqual(desired.TargetCpuUtil, actual.TargetCpuUtil) && !dcl.IsZeroValue(desired.TargetCpuUtil) {
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
		if compareInstanceScheduledMaintenance(c, &desired[i], &desired[i]) {
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
		return true
	}
	if compareInstanceScheduledMaintenanceStartTime(c, desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) {
		c.Config.Logger.Infof("Diff in StartTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if actual.CanDefer == nil && desired.CanDefer != nil && !dcl.IsEmptyValueIndirect(desired.CanDefer) {
		return true
	}
	if !reflect.DeepEqual(desired.CanDefer, actual.CanDefer) && !dcl.IsZeroValue(desired.CanDefer) {
		c.Config.Logger.Infof("Diff in CanDefer. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CanDefer), dcl.SprintResource(actual.CanDefer))
		return true
	}
	if actual.CanReschedule == nil && desired.CanReschedule != nil && !dcl.IsEmptyValueIndirect(desired.CanReschedule) {
		return true
	}
	if !reflect.DeepEqual(desired.CanReschedule, actual.CanReschedule) && !dcl.IsZeroValue(desired.CanReschedule) {
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
		if compareInstanceScheduledMaintenanceStartTime(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
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
		if compareInstanceSettings(c, &desired[i], &desired[i]) {
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
		return true
	}
	if !dcl.SliceEquals(desired.AuthorizedGaeApplications, actual.AuthorizedGaeApplications) && !dcl.IsZeroValue(desired.AuthorizedGaeApplications) {
		c.Config.Logger.Infof("Diff in AuthorizedGaeApplications. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AuthorizedGaeApplications), dcl.SprintResource(actual.AuthorizedGaeApplications))
		return true
	}
	if actual.Tier == nil && desired.Tier != nil && !dcl.IsEmptyValueIndirect(desired.Tier) {
		return true
	}
	if !reflect.DeepEqual(desired.Tier, actual.Tier) && !dcl.IsZeroValue(desired.Tier) {
		c.Config.Logger.Infof("Diff in Tier. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tier), dcl.SprintResource(actual.Tier))
		return true
	}
	if actual.Kind == nil && desired.Kind != nil && !dcl.IsEmptyValueIndirect(desired.Kind) {
		return true
	}
	if !reflect.DeepEqual(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if actual.AvailabilityType == nil && desired.AvailabilityType != nil && !dcl.IsEmptyValueIndirect(desired.AvailabilityType) {
		return true
	}
	if !reflect.DeepEqual(desired.AvailabilityType, actual.AvailabilityType) && !dcl.IsZeroValue(desired.AvailabilityType) {
		c.Config.Logger.Infof("Diff in AvailabilityType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AvailabilityType), dcl.SprintResource(actual.AvailabilityType))
		return true
	}
	if actual.PricingPlan == nil && desired.PricingPlan != nil && !dcl.IsEmptyValueIndirect(desired.PricingPlan) {
		return true
	}
	if !reflect.DeepEqual(desired.PricingPlan, actual.PricingPlan) && !dcl.IsZeroValue(desired.PricingPlan) {
		c.Config.Logger.Infof("Diff in PricingPlan. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PricingPlan), dcl.SprintResource(actual.PricingPlan))
		return true
	}
	if actual.ReplicationType == nil && desired.ReplicationType != nil && !dcl.IsEmptyValueIndirect(desired.ReplicationType) {
		return true
	}
	if !reflect.DeepEqual(desired.ReplicationType, actual.ReplicationType) && !dcl.IsZeroValue(desired.ReplicationType) {
		c.Config.Logger.Infof("Diff in ReplicationType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicationType), dcl.SprintResource(actual.ReplicationType))
		return true
	}
	if actual.ActivationPolicy == nil && desired.ActivationPolicy != nil && !dcl.IsEmptyValueIndirect(desired.ActivationPolicy) {
		return true
	}
	if !reflect.DeepEqual(desired.ActivationPolicy, actual.ActivationPolicy) && !dcl.IsZeroValue(desired.ActivationPolicy) {
		c.Config.Logger.Infof("Diff in ActivationPolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ActivationPolicy), dcl.SprintResource(actual.ActivationPolicy))
		return true
	}
	if actual.StorageAutoResize == nil && desired.StorageAutoResize != nil && !dcl.IsEmptyValueIndirect(desired.StorageAutoResize) {
		return true
	}
	if !reflect.DeepEqual(desired.StorageAutoResize, actual.StorageAutoResize) && !dcl.IsZeroValue(desired.StorageAutoResize) {
		c.Config.Logger.Infof("Diff in StorageAutoResize. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StorageAutoResize), dcl.SprintResource(actual.StorageAutoResize))
		return true
	}
	if actual.DataDiskType == nil && desired.DataDiskType != nil && !dcl.IsEmptyValueIndirect(desired.DataDiskType) {
		return true
	}
	if !reflect.DeepEqual(desired.DataDiskType, actual.DataDiskType) && !dcl.IsZeroValue(desired.DataDiskType) {
		c.Config.Logger.Infof("Diff in DataDiskType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DataDiskType), dcl.SprintResource(actual.DataDiskType))
		return true
	}
	if actual.DatabaseReplicationEnabled == nil && desired.DatabaseReplicationEnabled != nil && !dcl.IsEmptyValueIndirect(desired.DatabaseReplicationEnabled) {
		return true
	}
	if !reflect.DeepEqual(desired.DatabaseReplicationEnabled, actual.DatabaseReplicationEnabled) && !dcl.IsZeroValue(desired.DatabaseReplicationEnabled) {
		c.Config.Logger.Infof("Diff in DatabaseReplicationEnabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DatabaseReplicationEnabled), dcl.SprintResource(actual.DatabaseReplicationEnabled))
		return true
	}
	if actual.CrashSafeReplicationEnabled == nil && desired.CrashSafeReplicationEnabled != nil && !dcl.IsEmptyValueIndirect(desired.CrashSafeReplicationEnabled) {
		return true
	}
	if !reflect.DeepEqual(desired.CrashSafeReplicationEnabled, actual.CrashSafeReplicationEnabled) && !dcl.IsZeroValue(desired.CrashSafeReplicationEnabled) {
		c.Config.Logger.Infof("Diff in CrashSafeReplicationEnabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrashSafeReplicationEnabled), dcl.SprintResource(actual.CrashSafeReplicationEnabled))
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

	return r
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

	return r
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

	return r
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

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Instance) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInstance(b, c)
		if err != nil {
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()

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
