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
package compute

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Instance) validate() error {

	if err := dcl.Required(r, "zone"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Scheduling) {
		if err := r.Scheduling.validate(); err != nil {
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
func (r *InstanceDisks) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DiskEncryptionKey) {
		if err := r.DiskEncryptionKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.InitializeParams) {
		if err := r.InitializeParams.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceDisksDiskEncryptionKey) validate() error {
	return nil
}
func (r *InstanceDisksInitializeParams) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SourceImageEncryptionKey) {
		if err := r.SourceImageEncryptionKey.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceDisksInitializeParamsSourceImageEncryptionKey) validate() error {
	return nil
}
func (r *InstanceGuestAccelerators) validate() error {
	return nil
}
func (r *InstanceNetworkInterfaces) validate() error {
	return nil
}
func (r *InstanceNetworkInterfacesAccessConfigs) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	return nil
}
func (r *InstanceNetworkInterfacesAliasIPRanges) validate() error {
	return nil
}
func (r *InstanceScheduling) validate() error {
	return nil
}
func (r *InstanceServiceAccounts) validate() error {
	return nil
}
func (r *InstanceShieldedInstanceConfig) validate() error {
	return nil
}

func instanceGetURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"zone":    dcl.ValueOrEmptyString(r.Zone),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func instanceListURL(userBasePath, project, zone string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"zone":    zone,
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/instances", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func instanceCreateURL(userBasePath, project, zone string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"zone":    zone,
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/instances", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func instanceDeleteURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"zone":    dcl.ValueOrEmptyString(r.Zone),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func (r *Instance) SetPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"project": *n.Project,
		"zone":    *n.Zone,
		"name":    *n.Name,
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}/setIamPolicy", "https://www.googleapis.com/compute/v1/", userBasePath, fields)
}

func (r *Instance) getPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"project": *n.Project,
		"zone":    *n.Zone,
		"name":    *n.Name,
	}
	return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}/getIamPolicy", "https://www.googleapis.com/compute/v1/", userBasePath, fields)
}

func (r *Instance) IAMPolicyVersion() int {
	return 3
}

// instanceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceApiOperation interface {
	do(context.Context, *Instance, *Client) error
}

// newUpdateInstanceSetDeletionProtectionRequest creates a request for an
// Instance resource's setDeletionProtection update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceSetDeletionProtectionRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DeletionProtection; !dcl.IsEmptyValueIndirect(v) {
		req["deletionProtection"] = v
	}
	return req, nil
}

// marshalUpdateInstanceSetDeletionProtectionRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceSetDeletionProtectionRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceSetDeletionProtectionOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceSetDeletionProtectionOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setDeletionProtection")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceSetDeletionProtectionRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceSetDeletionProtectionRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateInstanceSetLabelsRequest creates a request for an
// Instance resource's setLabels update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceSetLabelsRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	b, err := c.getInstanceRaw(ctx, f.urlNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawLabelFingerprint, err := dcl.GetMapEntry(
		m,
		[]string{"labelFingerprint"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["labelFingerprint"] = rawLabelFingerprint.(string)
	}
	return req, nil
}

// marshalUpdateInstanceSetLabelsRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceSetLabelsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceSetLabelsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceSetLabelsOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setLabels")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceSetLabelsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceSetLabelsRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateInstanceSetMachineTypeRequest creates a request for an
// Instance resource's setMachineType update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceSetMachineTypeRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		req["machineType"] = v
	}
	return req, nil
}

// marshalUpdateInstanceSetMachineTypeRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceSetMachineTypeRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceSetMachineTypeOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceSetMachineTypeOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setMachineType")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceSetMachineTypeRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceSetMachineTypeRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateInstanceSetMetadataRequest creates a request for an
// Instance resource's setMetadata update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceSetMetadataRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.ListOfKeyValuesFromMap(f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into items: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["items"] = v
	}
	b, err := c.getInstanceRaw(ctx, f.urlNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawFingerprint, err := dcl.GetMapEntry(
		m,
		[]string{"metadata", "fingerprint"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["fingerprint"] = rawFingerprint.(string)
	}
	return req, nil
}

// marshalUpdateInstanceSetMetadataRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceSetMetadataRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceSetMetadataOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceSetMetadataOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setMetadata")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceSetMetadataRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceSetMetadataRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateInstanceSetTagsRequest creates a request for an
// Instance resource's setTags update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceSetTagsRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		req["items"] = v
	}
	b, err := c.getInstanceRaw(ctx, f.urlNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawFingerprint, err := dcl.GetMapEntry(
		m,
		[]string{"tags", "fingerprint"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["fingerprint"] = rawFingerprint.(string)
	}
	return req, nil
}

// marshalUpdateInstanceSetTagsRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceSetTagsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceSetTagsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceSetTagsOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setTags")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceSetTagsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceSetTagsRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateInstanceStartRequest creates a request for an
// Instance resource's start update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceStartRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateInstanceStartRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceStartRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceStartOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceStartOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "start")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceStartRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceStartRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateInstanceStopRequest creates a request for an
// Instance resource's stop update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceStopRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateInstanceStopRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceStopRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceStopOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceStopOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "stop")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceStopRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceStopRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
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
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateInstanceUpdateShieldedInstanceConfigRequest creates a request for an
// Instance resource's updateShieldedInstanceConfig update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceUpdateShieldedInstanceConfigRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandInstanceShieldedInstanceConfig(c, f.ShieldedInstanceConfig); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedInstanceConfig into shieldedInstanceConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["shieldedInstanceConfig"] = v
	}
	return req, nil
}

// marshalUpdateInstanceUpdateShieldedInstanceConfigRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceUpdateShieldedInstanceConfigRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceUpdateShieldedInstanceConfigOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceUpdateShieldedInstanceConfigOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "updateShieldedInstanceConfig")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceUpdateShieldedInstanceConfigRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceUpdateShieldedInstanceConfigRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInstanceRaw(ctx context.Context, project, zone, pageToken string, pageSize int32) ([]byte, error) {
	u, err := instanceListURL(c.Config.BasePath, project, zone)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listInstanceOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listInstance(ctx context.Context, project, zone, pageToken string, pageSize int32) ([]*Instance, string, error) {
	b, err := c.listInstanceRaw(ctx, project, zone, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInstanceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Instance
	for _, v := range m.Items {
		res := flattenInstance(c, v)
		res.Project = &project
		res.Zone = &zone
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
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetInstance(ctx, r.urlNormalized())
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
type createInstanceOperation struct {
	response map[string]interface{}
}

func (op *createInstanceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, zone := r.createFields()
	u, err := instanceCreateURL(c.Config.BasePath, project, zone)

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
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetInstance(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getInstanceRaw(ctx context.Context, r *Instance) ([]byte, error) {

	u, err := instanceGetURL(c.Config.BasePath, r.urlNormalized())
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

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Scheduling = canonicalizeInstanceScheduling(rawDesired.Scheduling, nil, opts...)
		rawDesired.ShieldedInstanceConfig = canonicalizeInstanceShieldedInstanceConfig(rawDesired.ShieldedInstanceConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.BoolCanonicalize(rawDesired.CanIPForward, rawInitial.CanIPForward) {
		rawDesired.CanIPForward = rawInitial.CanIPForward
	}
	if dcl.StringCanonicalize(rawDesired.CpuPlatform, rawInitial.CpuPlatform) {
		rawDesired.CpuPlatform = rawInitial.CpuPlatform
	}
	if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawInitial.CreationTimestamp) {
		rawDesired.CreationTimestamp = rawInitial.CreationTimestamp
	}
	if dcl.BoolCanonicalize(rawDesired.DeletionProtection, rawInitial.DeletionProtection) {
		rawDesired.DeletionProtection = rawInitial.DeletionProtection
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Disks) {
		rawDesired.Disks = rawInitial.Disks
	}
	if dcl.IsZeroValue(rawDesired.GuestAccelerators) {
		rawDesired.GuestAccelerators = rawInitial.GuestAccelerators
	}
	if dcl.StringCanonicalize(rawDesired.Hostname, rawInitial.Hostname) {
		rawDesired.Hostname = rawInitial.Hostname
	}
	if dcl.StringCanonicalize(rawDesired.Id, rawInitial.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.Metadata) {
		rawDesired.Metadata = rawInitial.Metadata
	}
	if dcl.NameToSelfLink(rawDesired.MachineType, rawInitial.MachineType) {
		rawDesired.MachineType = rawInitial.MachineType
	}
	if dcl.StringCanonicalize(rawDesired.MinCpuPlatform, rawInitial.MinCpuPlatform) {
		rawDesired.MinCpuPlatform = rawInitial.MinCpuPlatform
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.NetworkInterfaces) {
		rawDesired.NetworkInterfaces = rawInitial.NetworkInterfaces
	}
	rawDesired.Scheduling = canonicalizeInstanceScheduling(rawDesired.Scheduling, rawInitial.Scheduling, opts...)
	if dcl.IsZeroValue(rawDesired.ServiceAccounts) {
		rawDesired.ServiceAccounts = rawInitial.ServiceAccounts
	}
	rawDesired.ShieldedInstanceConfig = canonicalizeInstanceShieldedInstanceConfig(rawDesired.ShieldedInstanceConfig, rawInitial.ShieldedInstanceConfig, opts...)
	if dcl.IsZeroValue(rawDesired.Status) {
		rawDesired.Status = rawInitial.Status
	}
	if dcl.StringCanonicalize(rawDesired.StatusMessage, rawInitial.StatusMessage) {
		rawDesired.StatusMessage = rawInitial.StatusMessage
	}
	if dcl.IsZeroValue(rawDesired.Tags) {
		rawDesired.Tags = rawInitial.Tags
	}
	if dcl.NameToSelfLink(rawDesired.Zone, rawInitial.Zone) {
		rawDesired.Zone = rawInitial.Zone
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}

	return rawDesired, nil
}

func canonicalizeInstanceNewState(c *Client, rawNew, rawDesired *Instance) (*Instance, error) {

	if dcl.IsEmptyValueIndirect(rawNew.CanIPForward) && dcl.IsEmptyValueIndirect(rawDesired.CanIPForward) {
		rawNew.CanIPForward = rawDesired.CanIPForward
	} else {
		if dcl.BoolCanonicalize(rawDesired.CanIPForward, rawNew.CanIPForward) {
			rawNew.CanIPForward = rawDesired.CanIPForward
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CpuPlatform) && dcl.IsEmptyValueIndirect(rawDesired.CpuPlatform) {
		rawNew.CpuPlatform = rawDesired.CpuPlatform
	} else {
		if dcl.StringCanonicalize(rawDesired.CpuPlatform, rawNew.CpuPlatform) {
			rawNew.CpuPlatform = rawDesired.CpuPlatform
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
		if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawNew.CreationTimestamp) {
			rawNew.CreationTimestamp = rawDesired.CreationTimestamp
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeletionProtection) && dcl.IsEmptyValueIndirect(rawDesired.DeletionProtection) {
		rawNew.DeletionProtection = rawDesired.DeletionProtection
	} else {
		if dcl.BoolCanonicalize(rawDesired.DeletionProtection, rawNew.DeletionProtection) {
			rawNew.DeletionProtection = rawDesired.DeletionProtection
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	rawNew.Disks = rawDesired.Disks

	if dcl.IsEmptyValueIndirect(rawNew.GuestAccelerators) && dcl.IsEmptyValueIndirect(rawDesired.GuestAccelerators) {
		rawNew.GuestAccelerators = rawDesired.GuestAccelerators
	} else {
		rawNew.GuestAccelerators = canonicalizeNewInstanceGuestAcceleratorsSlice(c, rawDesired.GuestAccelerators, rawNew.GuestAccelerators)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Hostname) && dcl.IsEmptyValueIndirect(rawDesired.Hostname) {
		rawNew.Hostname = rawDesired.Hostname
	} else {
		if dcl.StringCanonicalize(rawDesired.Hostname, rawNew.Hostname) {
			rawNew.Hostname = rawDesired.Hostname
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Metadata) && dcl.IsEmptyValueIndirect(rawDesired.Metadata) {
		rawNew.Metadata = rawDesired.Metadata
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MachineType) && dcl.IsEmptyValueIndirect(rawDesired.MachineType) {
		rawNew.MachineType = rawDesired.MachineType
	} else {
		if dcl.NameToSelfLink(rawDesired.MachineType, rawNew.MachineType) {
			rawNew.MachineType = rawDesired.MachineType
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MinCpuPlatform) && dcl.IsEmptyValueIndirect(rawDesired.MinCpuPlatform) {
		rawNew.MinCpuPlatform = rawDesired.MinCpuPlatform
	} else {
		if dcl.StringCanonicalize(rawDesired.MinCpuPlatform, rawNew.MinCpuPlatform) {
			rawNew.MinCpuPlatform = rawDesired.MinCpuPlatform
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NetworkInterfaces) && dcl.IsEmptyValueIndirect(rawDesired.NetworkInterfaces) {
		rawNew.NetworkInterfaces = rawDesired.NetworkInterfaces
	} else {
		rawNew.NetworkInterfaces = canonicalizeNewInstanceNetworkInterfacesSlice(c, rawDesired.NetworkInterfaces, rawNew.NetworkInterfaces)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Scheduling) && dcl.IsEmptyValueIndirect(rawDesired.Scheduling) {
		rawNew.Scheduling = rawDesired.Scheduling
	} else {
		rawNew.Scheduling = canonicalizeNewInstanceScheduling(c, rawDesired.Scheduling, rawNew.Scheduling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccounts) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccounts) {
		rawNew.ServiceAccounts = rawDesired.ServiceAccounts
	} else {
		rawNew.ServiceAccounts = canonicalizeNewInstanceServiceAccountsSlice(c, rawDesired.ServiceAccounts, rawNew.ServiceAccounts)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ShieldedInstanceConfig) && dcl.IsEmptyValueIndirect(rawDesired.ShieldedInstanceConfig) {
		rawNew.ShieldedInstanceConfig = rawDesired.ShieldedInstanceConfig
	} else {
		rawNew.ShieldedInstanceConfig = canonicalizeNewInstanceShieldedInstanceConfig(c, rawDesired.ShieldedInstanceConfig, rawNew.ShieldedInstanceConfig)
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

	if dcl.IsEmptyValueIndirect(rawNew.Tags) && dcl.IsEmptyValueIndirect(rawDesired.Tags) {
		rawNew.Tags = rawDesired.Tags
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Zone) && dcl.IsEmptyValueIndirect(rawDesired.Zone) {
		rawNew.Zone = rawDesired.Zone
	} else {
		if dcl.NameToSelfLink(rawDesired.Zone, rawNew.Zone) {
			rawNew.Zone = rawDesired.Zone
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	return rawNew, nil
}

func canonicalizeInstanceDisks(des, initial *InstanceDisks, opts ...dcl.ApplyOption) *InstanceDisks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.AutoDelete, initial.AutoDelete) || dcl.IsZeroValue(des.AutoDelete) {
		des.AutoDelete = initial.AutoDelete
	}
	if dcl.BoolCanonicalize(des.Boot, initial.Boot) || dcl.IsZeroValue(des.Boot) {
		des.Boot = initial.Boot
	}
	if dcl.StringCanonicalize(des.DeviceName, initial.DeviceName) || dcl.IsZeroValue(des.DeviceName) {
		des.DeviceName = initial.DeviceName
	}
	des.DiskEncryptionKey = canonicalizeInstanceDisksDiskEncryptionKey(des.DiskEncryptionKey, initial.DiskEncryptionKey, opts...)
	if dcl.IsZeroValue(des.Index) {
		des.Index = initial.Index
	}
	des.InitializeParams = canonicalizeInstanceDisksInitializeParams(des.InitializeParams, initial.InitializeParams, opts...)
	if dcl.IsZeroValue(des.Interface) {
		des.Interface = initial.Interface
	}
	if dcl.IsZeroValue(des.Mode) {
		des.Mode = initial.Mode
	}
	if dcl.NameToSelfLink(des.Source, initial.Source) || dcl.IsZeroValue(des.Source) {
		des.Source = initial.Source
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewInstanceDisks(c *Client, des, nw *InstanceDisks) *InstanceDisks {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.AutoDelete, nw.AutoDelete) {
		nw.AutoDelete = des.AutoDelete
	}
	if dcl.BoolCanonicalize(des.Boot, nw.Boot) {
		nw.Boot = des.Boot
	}
	if dcl.StringCanonicalize(des.DeviceName, nw.DeviceName) {
		nw.DeviceName = des.DeviceName
	}
	nw.DiskEncryptionKey = canonicalizeNewInstanceDisksDiskEncryptionKey(c, des.DiskEncryptionKey, nw.DiskEncryptionKey)
	nw.InitializeParams = canonicalizeNewInstanceDisksInitializeParams(c, des.InitializeParams, nw.InitializeParams)
	if dcl.NameToSelfLink(des.Source, nw.Source) {
		nw.Source = des.Source
	}

	return nw
}

func canonicalizeNewInstanceDisksSet(c *Client, des, nw []InstanceDisks) []InstanceDisks {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceDisks
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceDisks(c, &d, &n) {
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

func canonicalizeNewInstanceDisksSlice(c *Client, des, nw []InstanceDisks) []InstanceDisks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceDisks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceDisks(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceDisksDiskEncryptionKey(des, initial *InstanceDisksDiskEncryptionKey, opts ...dcl.ApplyOption) *InstanceDisksDiskEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.RawKey, initial.RawKey) || dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.StringCanonicalize(des.RsaEncryptedKey, initial.RsaEncryptedKey) || dcl.IsZeroValue(des.RsaEncryptedKey) {
		des.RsaEncryptedKey = initial.RsaEncryptedKey
	}
	if dcl.StringCanonicalize(des.Sha256, initial.Sha256) || dcl.IsZeroValue(des.Sha256) {
		des.Sha256 = initial.Sha256
	}

	return des
}

func canonicalizeNewInstanceDisksDiskEncryptionKey(c *Client, des, nw *InstanceDisksDiskEncryptionKey) *InstanceDisksDiskEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RawKey, nw.RawKey) {
		nw.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.RsaEncryptedKey, nw.RsaEncryptedKey) {
		nw.RsaEncryptedKey = des.RsaEncryptedKey
	}
	if dcl.StringCanonicalize(des.Sha256, nw.Sha256) {
		nw.Sha256 = des.Sha256
	}

	return nw
}

func canonicalizeNewInstanceDisksDiskEncryptionKeySet(c *Client, des, nw []InstanceDisksDiskEncryptionKey) []InstanceDisksDiskEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceDisksDiskEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceDisksDiskEncryptionKey(c, &d, &n) {
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

func canonicalizeNewInstanceDisksDiskEncryptionKeySlice(c *Client, des, nw []InstanceDisksDiskEncryptionKey) []InstanceDisksDiskEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceDisksDiskEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceDisksDiskEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceDisksInitializeParams(des, initial *InstanceDisksInitializeParams, opts ...dcl.ApplyOption) *InstanceDisksInitializeParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.DiskName, initial.DiskName) || dcl.IsZeroValue(des.DiskName) {
		des.DiskName = initial.DiskName
	}
	if dcl.IsZeroValue(des.DiskSizeGb) {
		des.DiskSizeGb = initial.DiskSizeGb
	}
	if dcl.NameToSelfLink(des.DiskType, initial.DiskType) || dcl.IsZeroValue(des.DiskType) {
		des.DiskType = initial.DiskType
	}
	if dcl.StringCanonicalize(des.SourceImage, initial.SourceImage) || dcl.IsZeroValue(des.SourceImage) {
		des.SourceImage = initial.SourceImage
	}
	des.SourceImageEncryptionKey = canonicalizeInstanceDisksInitializeParamsSourceImageEncryptionKey(des.SourceImageEncryptionKey, initial.SourceImageEncryptionKey, opts...)

	return des
}

func canonicalizeNewInstanceDisksInitializeParams(c *Client, des, nw *InstanceDisksInitializeParams) *InstanceDisksInitializeParams {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.DiskName, nw.DiskName) {
		nw.DiskName = des.DiskName
	}
	if dcl.NameToSelfLink(des.DiskType, nw.DiskType) {
		nw.DiskType = des.DiskType
	}
	if dcl.StringCanonicalize(des.SourceImage, nw.SourceImage) {
		nw.SourceImage = des.SourceImage
	}
	nw.SourceImageEncryptionKey = canonicalizeNewInstanceDisksInitializeParamsSourceImageEncryptionKey(c, des.SourceImageEncryptionKey, nw.SourceImageEncryptionKey)

	return nw
}

func canonicalizeNewInstanceDisksInitializeParamsSet(c *Client, des, nw []InstanceDisksInitializeParams) []InstanceDisksInitializeParams {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceDisksInitializeParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceDisksInitializeParams(c, &d, &n) {
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

func canonicalizeNewInstanceDisksInitializeParamsSlice(c *Client, des, nw []InstanceDisksInitializeParams) []InstanceDisksInitializeParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceDisksInitializeParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceDisksInitializeParams(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceDisksInitializeParamsSourceImageEncryptionKey(des, initial *InstanceDisksInitializeParamsSourceImageEncryptionKey, opts ...dcl.ApplyOption) *InstanceDisksInitializeParamsSourceImageEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.RawKey, initial.RawKey) || dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.StringCanonicalize(des.Sha256, initial.Sha256) || dcl.IsZeroValue(des.Sha256) {
		des.Sha256 = initial.Sha256
	}

	return des
}

func canonicalizeNewInstanceDisksInitializeParamsSourceImageEncryptionKey(c *Client, des, nw *InstanceDisksInitializeParamsSourceImageEncryptionKey) *InstanceDisksInitializeParamsSourceImageEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RawKey, nw.RawKey) {
		nw.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.Sha256, nw.Sha256) {
		nw.Sha256 = des.Sha256
	}

	return nw
}

func canonicalizeNewInstanceDisksInitializeParamsSourceImageEncryptionKeySet(c *Client, des, nw []InstanceDisksInitializeParamsSourceImageEncryptionKey) []InstanceDisksInitializeParamsSourceImageEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceDisksInitializeParamsSourceImageEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceDisksInitializeParamsSourceImageEncryptionKey(c, &d, &n) {
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

func canonicalizeNewInstanceDisksInitializeParamsSourceImageEncryptionKeySlice(c *Client, des, nw []InstanceDisksInitializeParamsSourceImageEncryptionKey) []InstanceDisksInitializeParamsSourceImageEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceDisksInitializeParamsSourceImageEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceDisksInitializeParamsSourceImageEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGuestAccelerators(des, initial *InstanceGuestAccelerators, opts ...dcl.ApplyOption) *InstanceGuestAccelerators {
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

func canonicalizeNewInstanceGuestAccelerators(c *Client, des, nw *InstanceGuestAccelerators) *InstanceGuestAccelerators {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AcceleratorType, nw.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
	}

	return nw
}

func canonicalizeNewInstanceGuestAcceleratorsSet(c *Client, des, nw []InstanceGuestAccelerators) []InstanceGuestAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGuestAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGuestAccelerators(c, &d, &n) {
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

func canonicalizeNewInstanceGuestAcceleratorsSlice(c *Client, des, nw []InstanceGuestAccelerators) []InstanceGuestAccelerators {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceGuestAccelerators
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGuestAccelerators(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceNetworkInterfaces(des, initial *InstanceNetworkInterfaces, opts ...dcl.ApplyOption) *InstanceNetworkInterfaces {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AccessConfigs) {
		des.AccessConfigs = initial.AccessConfigs
	}
	if dcl.IsZeroValue(des.AliasIPRanges) {
		des.AliasIPRanges = initial.AliasIPRanges
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.NameToSelfLink(des.Network, initial.Network) || dcl.IsZeroValue(des.Network) {
		des.Network = initial.Network
	}
	if dcl.StringCanonicalize(des.NetworkIP, initial.NetworkIP) || dcl.IsZeroValue(des.NetworkIP) {
		des.NetworkIP = initial.NetworkIP
	}
	if dcl.NameToSelfLink(des.Subnetwork, initial.Subnetwork) || dcl.IsZeroValue(des.Subnetwork) {
		des.Subnetwork = initial.Subnetwork
	}

	return des
}

func canonicalizeNewInstanceNetworkInterfaces(c *Client, des, nw *InstanceNetworkInterfaces) *InstanceNetworkInterfaces {
	if des == nil || nw == nil {
		return nw
	}

	nw.AccessConfigs = canonicalizeNewInstanceNetworkInterfacesAccessConfigsSlice(c, des.AccessConfigs, nw.AccessConfigs)
	nw.AliasIPRanges = canonicalizeNewInstanceNetworkInterfacesAliasIPRangesSlice(c, des.AliasIPRanges, nw.AliasIPRanges)
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.NameToSelfLink(des.Network, nw.Network) {
		nw.Network = des.Network
	}
	if dcl.StringCanonicalize(des.NetworkIP, nw.NetworkIP) {
		nw.NetworkIP = des.NetworkIP
	}
	if dcl.NameToSelfLink(des.Subnetwork, nw.Subnetwork) {
		nw.Subnetwork = des.Subnetwork
	}

	return nw
}

func canonicalizeNewInstanceNetworkInterfacesSet(c *Client, des, nw []InstanceNetworkInterfaces) []InstanceNetworkInterfaces {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceNetworkInterfaces
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceNetworkInterfaces(c, &d, &n) {
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

func canonicalizeNewInstanceNetworkInterfacesSlice(c *Client, des, nw []InstanceNetworkInterfaces) []InstanceNetworkInterfaces {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceNetworkInterfaces
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceNetworkInterfaces(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceNetworkInterfacesAccessConfigs(des, initial *InstanceNetworkInterfacesAccessConfigs, opts ...dcl.ApplyOption) *InstanceNetworkInterfacesAccessConfigs {
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
	if dcl.StringCanonicalize(des.NatIP, initial.NatIP) || dcl.IsZeroValue(des.NatIP) {
		des.NatIP = initial.NatIP
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewInstanceNetworkInterfacesAccessConfigs(c *Client, des, nw *InstanceNetworkInterfacesAccessConfigs) *InstanceNetworkInterfacesAccessConfigs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.NatIP, nw.NatIP) {
		nw.NatIP = des.NatIP
	}

	return nw
}

func canonicalizeNewInstanceNetworkInterfacesAccessConfigsSet(c *Client, des, nw []InstanceNetworkInterfacesAccessConfigs) []InstanceNetworkInterfacesAccessConfigs {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceNetworkInterfacesAccessConfigs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceNetworkInterfacesAccessConfigs(c, &d, &n) {
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

func canonicalizeNewInstanceNetworkInterfacesAccessConfigsSlice(c *Client, des, nw []InstanceNetworkInterfacesAccessConfigs) []InstanceNetworkInterfacesAccessConfigs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceNetworkInterfacesAccessConfigs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceNetworkInterfacesAccessConfigs(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceNetworkInterfacesAliasIPRanges(des, initial *InstanceNetworkInterfacesAliasIPRanges, opts ...dcl.ApplyOption) *InstanceNetworkInterfacesAliasIPRanges {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.IPCidrRange, initial.IPCidrRange) || dcl.IsZeroValue(des.IPCidrRange) {
		des.IPCidrRange = initial.IPCidrRange
	}
	if dcl.StringCanonicalize(des.SubnetworkRangeName, initial.SubnetworkRangeName) || dcl.IsZeroValue(des.SubnetworkRangeName) {
		des.SubnetworkRangeName = initial.SubnetworkRangeName
	}

	return des
}

func canonicalizeNewInstanceNetworkInterfacesAliasIPRanges(c *Client, des, nw *InstanceNetworkInterfacesAliasIPRanges) *InstanceNetworkInterfacesAliasIPRanges {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.IPCidrRange, nw.IPCidrRange) {
		nw.IPCidrRange = des.IPCidrRange
	}
	if dcl.StringCanonicalize(des.SubnetworkRangeName, nw.SubnetworkRangeName) {
		nw.SubnetworkRangeName = des.SubnetworkRangeName
	}

	return nw
}

func canonicalizeNewInstanceNetworkInterfacesAliasIPRangesSet(c *Client, des, nw []InstanceNetworkInterfacesAliasIPRanges) []InstanceNetworkInterfacesAliasIPRanges {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceNetworkInterfacesAliasIPRanges
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceNetworkInterfacesAliasIPRanges(c, &d, &n) {
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

func canonicalizeNewInstanceNetworkInterfacesAliasIPRangesSlice(c *Client, des, nw []InstanceNetworkInterfacesAliasIPRanges) []InstanceNetworkInterfacesAliasIPRanges {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceNetworkInterfacesAliasIPRanges
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceNetworkInterfacesAliasIPRanges(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceScheduling(des, initial *InstanceScheduling, opts ...dcl.ApplyOption) *InstanceScheduling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.AutomaticRestart, initial.AutomaticRestart) || dcl.IsZeroValue(des.AutomaticRestart) {
		des.AutomaticRestart = initial.AutomaticRestart
	}
	if dcl.StringCanonicalize(des.OnHostMaintenance, initial.OnHostMaintenance) || dcl.IsZeroValue(des.OnHostMaintenance) {
		des.OnHostMaintenance = initial.OnHostMaintenance
	}
	if dcl.BoolCanonicalize(des.Preemptible, initial.Preemptible) || dcl.IsZeroValue(des.Preemptible) {
		des.Preemptible = initial.Preemptible
	}

	return des
}

func canonicalizeNewInstanceScheduling(c *Client, des, nw *InstanceScheduling) *InstanceScheduling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.AutomaticRestart, nw.AutomaticRestart) {
		nw.AutomaticRestart = des.AutomaticRestart
	}
	if dcl.StringCanonicalize(des.OnHostMaintenance, nw.OnHostMaintenance) {
		nw.OnHostMaintenance = des.OnHostMaintenance
	}
	if dcl.BoolCanonicalize(des.Preemptible, nw.Preemptible) {
		nw.Preemptible = des.Preemptible
	}

	return nw
}

func canonicalizeNewInstanceSchedulingSet(c *Client, des, nw []InstanceScheduling) []InstanceScheduling {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceScheduling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceScheduling(c, &d, &n) {
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

func canonicalizeNewInstanceSchedulingSlice(c *Client, des, nw []InstanceScheduling) []InstanceScheduling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceScheduling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceScheduling(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceServiceAccounts(des, initial *InstanceServiceAccounts, opts ...dcl.ApplyOption) *InstanceServiceAccounts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Email, initial.Email) || dcl.IsZeroValue(des.Email) {
		des.Email = initial.Email
	}
	if dcl.IsZeroValue(des.Scopes) {
		des.Scopes = initial.Scopes
	}

	return des
}

func canonicalizeNewInstanceServiceAccounts(c *Client, des, nw *InstanceServiceAccounts) *InstanceServiceAccounts {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Email, nw.Email) {
		nw.Email = des.Email
	}

	return nw
}

func canonicalizeNewInstanceServiceAccountsSet(c *Client, des, nw []InstanceServiceAccounts) []InstanceServiceAccounts {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceServiceAccounts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceServiceAccounts(c, &d, &n) {
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

func canonicalizeNewInstanceServiceAccountsSlice(c *Client, des, nw []InstanceServiceAccounts) []InstanceServiceAccounts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceServiceAccounts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceServiceAccounts(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceShieldedInstanceConfig(des, initial *InstanceShieldedInstanceConfig, opts ...dcl.ApplyOption) *InstanceShieldedInstanceConfig {
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
	if dcl.BoolCanonicalize(des.EnableVtpm, initial.EnableVtpm) || dcl.IsZeroValue(des.EnableVtpm) {
		des.EnableVtpm = initial.EnableVtpm
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, initial.EnableIntegrityMonitoring) || dcl.IsZeroValue(des.EnableIntegrityMonitoring) {
		des.EnableIntegrityMonitoring = initial.EnableIntegrityMonitoring
	}

	return des
}

func canonicalizeNewInstanceShieldedInstanceConfig(c *Client, des, nw *InstanceShieldedInstanceConfig) *InstanceShieldedInstanceConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnableSecureBoot, nw.EnableSecureBoot) {
		nw.EnableSecureBoot = des.EnableSecureBoot
	}
	if dcl.BoolCanonicalize(des.EnableVtpm, nw.EnableVtpm) {
		nw.EnableVtpm = des.EnableVtpm
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, nw.EnableIntegrityMonitoring) {
		nw.EnableIntegrityMonitoring = des.EnableIntegrityMonitoring
	}

	return nw
}

func canonicalizeNewInstanceShieldedInstanceConfigSet(c *Client, des, nw []InstanceShieldedInstanceConfig) []InstanceShieldedInstanceConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceShieldedInstanceConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceShieldedInstanceConfig(c, &d, &n) {
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

func canonicalizeNewInstanceShieldedInstanceConfigSlice(c *Client, des, nw []InstanceShieldedInstanceConfig) []InstanceShieldedInstanceConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceShieldedInstanceConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceShieldedInstanceConfig(c, &d, &n))
	}

	return items
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
	// New style diffs.
	if d, err := dcl.Diff(desired.CanIPForward, actual.CanIPForward, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "CanIPForward"})
	}

	if d, err := dcl.Diff(desired.CpuPlatform, actual.CpuPlatform, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "CpuPlatform"})
	}

	if d, err := dcl.Diff(desired.CreationTimestamp, actual.CreationTimestamp, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "CreationTimestamp"})
	}

	if d, err := dcl.Diff(desired.DeletionProtection, actual.DeletionProtection, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{
			UpdateOp: &updateInstanceSetDeletionProtectionOperation{}, FieldName: "DeletionProtection",
		})
	}

	if d, err := dcl.Diff(desired.Description, actual.Description, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "Description"})
	}

	if d, err := dcl.Diff(desired.Hostname, actual.Hostname, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "Hostname"})
	}

	if d, err := dcl.Diff(desired.Id, actual.Id, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "Id"})
	}

	if d, err := dcl.Diff(desired.Labels, actual.Labels, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{
			UpdateOp: &updateInstanceSetLabelsOperation{}, FieldName: "Labels",
		})
	}

	if d, err := dcl.Diff(desired.Metadata, actual.Metadata, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{
			UpdateOp: &updateInstanceSetMetadataOperation{}, FieldName: "Metadata",
		})
	}

	if d, err := dcl.Diff(desired.MachineType, actual.MachineType, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "ReferenceType"}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{
			UpdateOp: &updateInstanceSetMachineTypeOperation{}, FieldName: "MachineType",
		})
	}

	if d, err := dcl.Diff(desired.MinCpuPlatform, actual.MinCpuPlatform, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "MinCpuPlatform"})
	}

	if d, err := dcl.Diff(desired.Name, actual.Name, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "Name"})
	}

	if d, err := dcl.Diff(desired.StatusMessage, actual.StatusMessage, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "StatusMessage"})
	}

	if d, err := dcl.Diff(desired.Zone, actual.Zone, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "ReferenceType"}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "Zone"})
	}

	if d, err := dcl.Diff(desired.SelfLink, actual.SelfLink, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, FieldName: "SelfLink"})
	}

	if !dcl.IsZeroValue(desired.CanIPForward) && !dcl.BoolCanonicalize(desired.CanIPForward, actual.CanIPForward) {
		c.Config.Logger.Infof("Detected diff in CanIPForward.\nDESIRED: %v\nACTUAL: %v", desired.CanIPForward, actual.CanIPForward)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "CanIPForward",
		})
	}
	if !dcl.IsZeroValue(desired.DeletionProtection) && !dcl.BoolCanonicalize(desired.DeletionProtection, actual.DeletionProtection) {
		c.Config.Logger.Infof("Detected diff in DeletionProtection.\nDESIRED: %v\nACTUAL: %v", desired.DeletionProtection, actual.DeletionProtection)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceSetDeletionProtectionOperation{},
			FieldName: "DeletionProtection",
		})

	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if compareInstanceGuestAcceleratorsSlice(c, desired.GuestAccelerators, actual.GuestAccelerators) {
		c.Config.Logger.Infof("Detected diff in GuestAccelerators.\nDESIRED: %v\nACTUAL: %v", desired.GuestAccelerators, actual.GuestAccelerators)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "GuestAccelerators",
		})
	}
	if !dcl.IsZeroValue(desired.Hostname) && !dcl.StringCanonicalize(desired.Hostname, actual.Hostname) {
		c.Config.Logger.Infof("Detected diff in Hostname.\nDESIRED: %v\nACTUAL: %v", desired.Hostname, actual.Hostname)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Hostname",
		})
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceSetLabelsOperation{},
			FieldName: "Labels",
		})

	}
	if !dcl.MapEquals(desired.Metadata, actual.Metadata, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Metadata.\nDESIRED: %v\nACTUAL: %v", desired.Metadata, actual.Metadata)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceSetMetadataOperation{},
			FieldName: "Metadata",
		})

	}
	if !dcl.IsZeroValue(desired.MachineType) && !dcl.NameToSelfLink(desired.MachineType, actual.MachineType) {
		c.Config.Logger.Infof("Detected diff in MachineType.\nDESIRED: %v\nACTUAL: %v", desired.MachineType, actual.MachineType)
		switch *desired.Status {
		case "TERMINATED":
		case "RUNNING":
			diffs = append(diffs, instanceDiff{
				UpdateOp:  &updateInstanceStopOperation{},
				FieldName: "Status",
			})
		default:
			return nil, dcl.ApplyInfeasibleError{
				Message: fmt.Sprintf("Update blocked by invalid state: %#v.", desired.Status),
			}
		}

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceSetMachineTypeOperation{},
			FieldName: "MachineType",
		})

		switch *desired.Status {
		case "RUNNING":
			diffs = append(diffs, instanceDiff{
				UpdateOp:  &updateInstanceStartOperation{},
				FieldName: "Status",
			})
		default:
			c.Config.Logger.Warningf("desired state %q is not being reset to its desired value.", desired.Status)
		}
	}
	if !dcl.IsZeroValue(desired.MinCpuPlatform) && !dcl.StringCanonicalize(desired.MinCpuPlatform, actual.MinCpuPlatform) {
		c.Config.Logger.Infof("Detected diff in MinCpuPlatform.\nDESIRED: %v\nACTUAL: %v", desired.MinCpuPlatform, actual.MinCpuPlatform)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "MinCpuPlatform",
		})
	}
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if compareInstanceNetworkInterfacesSlice(c, desired.NetworkInterfaces, actual.NetworkInterfaces) {
		c.Config.Logger.Infof("Detected diff in NetworkInterfaces.\nDESIRED: %v\nACTUAL: %v", desired.NetworkInterfaces, actual.NetworkInterfaces)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "NetworkInterfaces",
		})
	}
	if compareInstanceScheduling(c, desired.Scheduling, actual.Scheduling) {
		c.Config.Logger.Infof("Detected diff in Scheduling.\nDESIRED: %v\nACTUAL: %v", desired.Scheduling, actual.Scheduling)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Scheduling",
		})
	}
	if compareInstanceServiceAccountsSlice(c, desired.ServiceAccounts, actual.ServiceAccounts) {
		c.Config.Logger.Infof("Detected diff in ServiceAccounts.\nDESIRED: %v\nACTUAL: %v", desired.ServiceAccounts, actual.ServiceAccounts)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "ServiceAccounts",
		})
	}
	if compareInstanceShieldedInstanceConfig(c, desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig) {
		c.Config.Logger.Infof("Detected diff in ShieldedInstanceConfig.\nDESIRED: %v\nACTUAL: %v", desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateShieldedInstanceConfigOperation{},
			FieldName: "ShieldedInstanceConfig",
		})

	}
	if !reflect.DeepEqual(desired.Status, actual.Status) {
		c.Config.Logger.Infof("Detected diff in Status.\nDESIRED: %v\nACTUAL: %v", desired.Status, actual.Status)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Status",
		})
	}
	if !dcl.StringSliceEquals(desired.Tags, actual.Tags) {
		c.Config.Logger.Infof("Detected diff in Tags.\nDESIRED: %v\nACTUAL: %v", desired.Tags, actual.Tags)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceSetTagsOperation{},
			FieldName: "Tags",
		})

	}
	if !dcl.IsZeroValue(desired.Zone) && !dcl.NameToSelfLink(desired.Zone, actual.Zone) {
		c.Config.Logger.Infof("Detected diff in Zone.\nDESIRED: %v\nACTUAL: %v", desired.Zone, actual.Zone)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Zone",
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
func compareInstanceDisks(c *Client, desired, actual *InstanceDisks) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AutoDelete == nil && desired.AutoDelete != nil && !dcl.IsEmptyValueIndirect(desired.AutoDelete) {
		c.Config.Logger.Infof("desired AutoDelete %s - but actually nil", dcl.SprintResource(desired.AutoDelete))
		return true
	}
	if !dcl.BoolCanonicalize(desired.AutoDelete, actual.AutoDelete) && !dcl.IsZeroValue(desired.AutoDelete) {
		c.Config.Logger.Infof("Diff in AutoDelete. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoDelete), dcl.SprintResource(actual.AutoDelete))
		return true
	}
	if actual.Boot == nil && desired.Boot != nil && !dcl.IsEmptyValueIndirect(desired.Boot) {
		c.Config.Logger.Infof("desired Boot %s - but actually nil", dcl.SprintResource(desired.Boot))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Boot, actual.Boot) && !dcl.IsZeroValue(desired.Boot) {
		c.Config.Logger.Infof("Diff in Boot. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Boot), dcl.SprintResource(actual.Boot))
		return true
	}
	if actual.DeviceName == nil && desired.DeviceName != nil && !dcl.IsEmptyValueIndirect(desired.DeviceName) {
		c.Config.Logger.Infof("desired DeviceName %s - but actually nil", dcl.SprintResource(desired.DeviceName))
		return true
	}
	if !dcl.StringCanonicalize(desired.DeviceName, actual.DeviceName) && !dcl.IsZeroValue(desired.DeviceName) {
		c.Config.Logger.Infof("Diff in DeviceName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DeviceName), dcl.SprintResource(actual.DeviceName))
		return true
	}
	if actual.DiskEncryptionKey == nil && desired.DiskEncryptionKey != nil && !dcl.IsEmptyValueIndirect(desired.DiskEncryptionKey) {
		c.Config.Logger.Infof("desired DiskEncryptionKey %s - but actually nil", dcl.SprintResource(desired.DiskEncryptionKey))
		return true
	}
	if compareInstanceDisksDiskEncryptionKey(c, desired.DiskEncryptionKey, actual.DiskEncryptionKey) && !dcl.IsZeroValue(desired.DiskEncryptionKey) {
		c.Config.Logger.Infof("Diff in DiskEncryptionKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskEncryptionKey), dcl.SprintResource(actual.DiskEncryptionKey))
		return true
	}
	if actual.Index == nil && desired.Index != nil && !dcl.IsEmptyValueIndirect(desired.Index) {
		c.Config.Logger.Infof("desired Index %s - but actually nil", dcl.SprintResource(desired.Index))
		return true
	}
	if !reflect.DeepEqual(desired.Index, actual.Index) && !dcl.IsZeroValue(desired.Index) {
		c.Config.Logger.Infof("Diff in Index. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Index), dcl.SprintResource(actual.Index))
		return true
	}
	if actual.Interface == nil && desired.Interface != nil && !dcl.IsEmptyValueIndirect(desired.Interface) {
		c.Config.Logger.Infof("desired Interface %s - but actually nil", dcl.SprintResource(desired.Interface))
		return true
	}
	if !reflect.DeepEqual(desired.Interface, actual.Interface) && !dcl.IsZeroValue(desired.Interface) {
		c.Config.Logger.Infof("Diff in Interface. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Interface), dcl.SprintResource(actual.Interface))
		return true
	}
	if actual.Mode == nil && desired.Mode != nil && !dcl.IsEmptyValueIndirect(desired.Mode) {
		c.Config.Logger.Infof("desired Mode %s - but actually nil", dcl.SprintResource(desired.Mode))
		return true
	}
	if !reflect.DeepEqual(desired.Mode, actual.Mode) && !dcl.IsZeroValue(desired.Mode) {
		c.Config.Logger.Infof("Diff in Mode. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Mode), dcl.SprintResource(actual.Mode))
		return true
	}
	if actual.Source == nil && desired.Source != nil && !dcl.IsEmptyValueIndirect(desired.Source) {
		c.Config.Logger.Infof("desired Source %s - but actually nil", dcl.SprintResource(desired.Source))
		return true
	}
	if !dcl.NameToSelfLink(desired.Source, actual.Source) && !dcl.IsZeroValue(desired.Source) {
		c.Config.Logger.Infof("Diff in Source. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Source), dcl.SprintResource(actual.Source))
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}

func compareInstanceDisksSlice(c *Client, desired, actual []InstanceDisks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisks, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDisks(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDisks, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDisksMap(c *Client, desired, actual map[string]InstanceDisks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisks, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceDisks, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceDisks(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceDisks, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceDisksDiskEncryptionKey(c *Client, desired, actual *InstanceDisksDiskEncryptionKey) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RawKey == nil && desired.RawKey != nil && !dcl.IsEmptyValueIndirect(desired.RawKey) {
		c.Config.Logger.Infof("desired RawKey %s - but actually nil", dcl.SprintResource(desired.RawKey))
		return true
	}
	if !dcl.StringCanonicalize(desired.RawKey, actual.RawKey) && !dcl.IsZeroValue(desired.RawKey) {
		c.Config.Logger.Infof("Diff in RawKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RawKey), dcl.SprintResource(actual.RawKey))
		return true
	}
	if actual.RsaEncryptedKey == nil && desired.RsaEncryptedKey != nil && !dcl.IsEmptyValueIndirect(desired.RsaEncryptedKey) {
		c.Config.Logger.Infof("desired RsaEncryptedKey %s - but actually nil", dcl.SprintResource(desired.RsaEncryptedKey))
		return true
	}
	if !dcl.StringCanonicalize(desired.RsaEncryptedKey, actual.RsaEncryptedKey) && !dcl.IsZeroValue(desired.RsaEncryptedKey) {
		c.Config.Logger.Infof("Diff in RsaEncryptedKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RsaEncryptedKey), dcl.SprintResource(actual.RsaEncryptedKey))
		return true
	}
	return false
}

func compareInstanceDisksDiskEncryptionKeySlice(c *Client, desired, actual []InstanceDisksDiskEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisksDiskEncryptionKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDisksDiskEncryptionKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDisksDiskEncryptionKey, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDisksDiskEncryptionKeyMap(c *Client, desired, actual map[string]InstanceDisksDiskEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisksDiskEncryptionKey, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceDisksDiskEncryptionKey, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceDisksDiskEncryptionKey(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceDisksDiskEncryptionKey, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceDisksInitializeParams(c *Client, desired, actual *InstanceDisksInitializeParams) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.DiskName == nil && desired.DiskName != nil && !dcl.IsEmptyValueIndirect(desired.DiskName) {
		c.Config.Logger.Infof("desired DiskName %s - but actually nil", dcl.SprintResource(desired.DiskName))
		return true
	}
	if !dcl.StringCanonicalize(desired.DiskName, actual.DiskName) && !dcl.IsZeroValue(desired.DiskName) {
		c.Config.Logger.Infof("Diff in DiskName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskName), dcl.SprintResource(actual.DiskName))
		return true
	}
	if actual.DiskSizeGb == nil && desired.DiskSizeGb != nil && !dcl.IsEmptyValueIndirect(desired.DiskSizeGb) {
		c.Config.Logger.Infof("desired DiskSizeGb %s - but actually nil", dcl.SprintResource(desired.DiskSizeGb))
		return true
	}
	if !reflect.DeepEqual(desired.DiskSizeGb, actual.DiskSizeGb) && !dcl.IsZeroValue(desired.DiskSizeGb) {
		c.Config.Logger.Infof("Diff in DiskSizeGb. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskSizeGb), dcl.SprintResource(actual.DiskSizeGb))
		return true
	}
	if actual.DiskType == nil && desired.DiskType != nil && !dcl.IsEmptyValueIndirect(desired.DiskType) {
		c.Config.Logger.Infof("desired DiskType %s - but actually nil", dcl.SprintResource(desired.DiskType))
		return true
	}
	if !dcl.NameToSelfLink(desired.DiskType, actual.DiskType) && !dcl.IsZeroValue(desired.DiskType) {
		c.Config.Logger.Infof("Diff in DiskType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskType), dcl.SprintResource(actual.DiskType))
		return true
	}
	if actual.SourceImage == nil && desired.SourceImage != nil && !dcl.IsEmptyValueIndirect(desired.SourceImage) {
		c.Config.Logger.Infof("desired SourceImage %s - but actually nil", dcl.SprintResource(desired.SourceImage))
		return true
	}
	if !dcl.StringCanonicalize(desired.SourceImage, actual.SourceImage) && !dcl.IsZeroValue(desired.SourceImage) {
		c.Config.Logger.Infof("Diff in SourceImage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceImage), dcl.SprintResource(actual.SourceImage))
		return true
	}
	if actual.SourceImageEncryptionKey == nil && desired.SourceImageEncryptionKey != nil && !dcl.IsEmptyValueIndirect(desired.SourceImageEncryptionKey) {
		c.Config.Logger.Infof("desired SourceImageEncryptionKey %s - but actually nil", dcl.SprintResource(desired.SourceImageEncryptionKey))
		return true
	}
	if compareInstanceDisksInitializeParamsSourceImageEncryptionKey(c, desired.SourceImageEncryptionKey, actual.SourceImageEncryptionKey) && !dcl.IsZeroValue(desired.SourceImageEncryptionKey) {
		c.Config.Logger.Infof("Diff in SourceImageEncryptionKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceImageEncryptionKey), dcl.SprintResource(actual.SourceImageEncryptionKey))
		return true
	}
	return false
}

func compareInstanceDisksInitializeParamsSlice(c *Client, desired, actual []InstanceDisksInitializeParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisksInitializeParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDisksInitializeParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDisksInitializeParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDisksInitializeParamsMap(c *Client, desired, actual map[string]InstanceDisksInitializeParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisksInitializeParams, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceDisksInitializeParams, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceDisksInitializeParams(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceDisksInitializeParams, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceDisksInitializeParamsSourceImageEncryptionKey(c *Client, desired, actual *InstanceDisksInitializeParamsSourceImageEncryptionKey) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RawKey == nil && desired.RawKey != nil && !dcl.IsEmptyValueIndirect(desired.RawKey) {
		c.Config.Logger.Infof("desired RawKey %s - but actually nil", dcl.SprintResource(desired.RawKey))
		return true
	}
	if !dcl.StringCanonicalize(desired.RawKey, actual.RawKey) && !dcl.IsZeroValue(desired.RawKey) {
		c.Config.Logger.Infof("Diff in RawKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RawKey), dcl.SprintResource(actual.RawKey))
		return true
	}
	return false
}

func compareInstanceDisksInitializeParamsSourceImageEncryptionKeySlice(c *Client, desired, actual []InstanceDisksInitializeParamsSourceImageEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisksInitializeParamsSourceImageEncryptionKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDisksInitializeParamsSourceImageEncryptionKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDisksInitializeParamsSourceImageEncryptionKey, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDisksInitializeParamsSourceImageEncryptionKeyMap(c *Client, desired, actual map[string]InstanceDisksInitializeParamsSourceImageEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisksInitializeParamsSourceImageEncryptionKey, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceDisksInitializeParamsSourceImageEncryptionKey, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceDisksInitializeParamsSourceImageEncryptionKey(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceDisksInitializeParamsSourceImageEncryptionKey, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGuestAccelerators(c *Client, desired, actual *InstanceGuestAccelerators) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AcceleratorCount == nil && desired.AcceleratorCount != nil && !dcl.IsEmptyValueIndirect(desired.AcceleratorCount) {
		c.Config.Logger.Infof("desired AcceleratorCount %s - but actually nil", dcl.SprintResource(desired.AcceleratorCount))
		return true
	}
	if !reflect.DeepEqual(desired.AcceleratorCount, actual.AcceleratorCount) && !dcl.IsZeroValue(desired.AcceleratorCount) {
		c.Config.Logger.Infof("Diff in AcceleratorCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorCount), dcl.SprintResource(actual.AcceleratorCount))
		return true
	}
	if actual.AcceleratorType == nil && desired.AcceleratorType != nil && !dcl.IsEmptyValueIndirect(desired.AcceleratorType) {
		c.Config.Logger.Infof("desired AcceleratorType %s - but actually nil", dcl.SprintResource(desired.AcceleratorType))
		return true
	}
	if !dcl.StringCanonicalize(desired.AcceleratorType, actual.AcceleratorType) && !dcl.IsZeroValue(desired.AcceleratorType) {
		c.Config.Logger.Infof("Diff in AcceleratorType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorType), dcl.SprintResource(actual.AcceleratorType))
		return true
	}
	return false
}

func compareInstanceGuestAcceleratorsSlice(c *Client, desired, actual []InstanceGuestAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGuestAccelerators, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGuestAccelerators(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGuestAccelerators, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGuestAcceleratorsMap(c *Client, desired, actual map[string]InstanceGuestAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGuestAccelerators, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGuestAccelerators, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGuestAccelerators(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGuestAccelerators, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceNetworkInterfaces(c *Client, desired, actual *InstanceNetworkInterfaces) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AccessConfigs == nil && desired.AccessConfigs != nil && !dcl.IsEmptyValueIndirect(desired.AccessConfigs) {
		c.Config.Logger.Infof("desired AccessConfigs %s - but actually nil", dcl.SprintResource(desired.AccessConfigs))
		return true
	}
	if compareInstanceNetworkInterfacesAccessConfigsSlice(c, desired.AccessConfigs, actual.AccessConfigs) && !dcl.IsZeroValue(desired.AccessConfigs) {
		c.Config.Logger.Infof("Diff in AccessConfigs. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AccessConfigs), dcl.SprintResource(actual.AccessConfigs))
		return true
	}
	if actual.AliasIPRanges == nil && desired.AliasIPRanges != nil && !dcl.IsEmptyValueIndirect(desired.AliasIPRanges) {
		c.Config.Logger.Infof("desired AliasIPRanges %s - but actually nil", dcl.SprintResource(desired.AliasIPRanges))
		return true
	}
	if compareInstanceNetworkInterfacesAliasIPRangesSlice(c, desired.AliasIPRanges, actual.AliasIPRanges) && !dcl.IsZeroValue(desired.AliasIPRanges) {
		c.Config.Logger.Infof("Diff in AliasIPRanges. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AliasIPRanges), dcl.SprintResource(actual.AliasIPRanges))
		return true
	}
	if actual.Network == nil && desired.Network != nil && !dcl.IsEmptyValueIndirect(desired.Network) {
		c.Config.Logger.Infof("desired Network %s - but actually nil", dcl.SprintResource(desired.Network))
		return true
	}
	if !dcl.NameToSelfLink(desired.Network, actual.Network) && !dcl.IsZeroValue(desired.Network) {
		c.Config.Logger.Infof("Diff in Network. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Network), dcl.SprintResource(actual.Network))
		return true
	}
	if actual.NetworkIP == nil && desired.NetworkIP != nil && !dcl.IsEmptyValueIndirect(desired.NetworkIP) {
		c.Config.Logger.Infof("desired NetworkIP %s - but actually nil", dcl.SprintResource(desired.NetworkIP))
		return true
	}
	if !dcl.StringCanonicalize(desired.NetworkIP, actual.NetworkIP) && !dcl.IsZeroValue(desired.NetworkIP) {
		c.Config.Logger.Infof("Diff in NetworkIP. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NetworkIP), dcl.SprintResource(actual.NetworkIP))
		return true
	}
	if actual.Subnetwork == nil && desired.Subnetwork != nil && !dcl.IsEmptyValueIndirect(desired.Subnetwork) {
		c.Config.Logger.Infof("desired Subnetwork %s - but actually nil", dcl.SprintResource(desired.Subnetwork))
		return true
	}
	if !dcl.NameToSelfLink(desired.Subnetwork, actual.Subnetwork) && !dcl.IsZeroValue(desired.Subnetwork) {
		c.Config.Logger.Infof("Diff in Subnetwork. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Subnetwork), dcl.SprintResource(actual.Subnetwork))
		return true
	}
	return false
}

func compareInstanceNetworkInterfacesSlice(c *Client, desired, actual []InstanceNetworkInterfaces) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceNetworkInterfaces, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceNetworkInterfaces(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfaces, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceNetworkInterfacesMap(c *Client, desired, actual map[string]InstanceNetworkInterfaces) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceNetworkInterfaces, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfaces, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceNetworkInterfaces(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfaces, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceNetworkInterfacesAccessConfigs(c *Client, desired, actual *InstanceNetworkInterfacesAccessConfigs) bool {
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
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.NatIP == nil && desired.NatIP != nil && !dcl.IsEmptyValueIndirect(desired.NatIP) {
		c.Config.Logger.Infof("desired NatIP %s - but actually nil", dcl.SprintResource(desired.NatIP))
		return true
	}
	if !dcl.StringCanonicalize(desired.NatIP, actual.NatIP) && !dcl.IsZeroValue(desired.NatIP) {
		c.Config.Logger.Infof("Diff in NatIP. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NatIP), dcl.SprintResource(actual.NatIP))
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}

func compareInstanceNetworkInterfacesAccessConfigsSlice(c *Client, desired, actual []InstanceNetworkInterfacesAccessConfigs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceNetworkInterfacesAccessConfigs, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceNetworkInterfacesAccessConfigs(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfacesAccessConfigs, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceNetworkInterfacesAccessConfigsMap(c *Client, desired, actual map[string]InstanceNetworkInterfacesAccessConfigs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceNetworkInterfacesAccessConfigs, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfacesAccessConfigs, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceNetworkInterfacesAccessConfigs(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfacesAccessConfigs, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceNetworkInterfacesAliasIPRanges(c *Client, desired, actual *InstanceNetworkInterfacesAliasIPRanges) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.IPCidrRange == nil && desired.IPCidrRange != nil && !dcl.IsEmptyValueIndirect(desired.IPCidrRange) {
		c.Config.Logger.Infof("desired IPCidrRange %s - but actually nil", dcl.SprintResource(desired.IPCidrRange))
		return true
	}
	if !dcl.StringCanonicalize(desired.IPCidrRange, actual.IPCidrRange) && !dcl.IsZeroValue(desired.IPCidrRange) {
		c.Config.Logger.Infof("Diff in IPCidrRange. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPCidrRange), dcl.SprintResource(actual.IPCidrRange))
		return true
	}
	if actual.SubnetworkRangeName == nil && desired.SubnetworkRangeName != nil && !dcl.IsEmptyValueIndirect(desired.SubnetworkRangeName) {
		c.Config.Logger.Infof("desired SubnetworkRangeName %s - but actually nil", dcl.SprintResource(desired.SubnetworkRangeName))
		return true
	}
	if !dcl.StringCanonicalize(desired.SubnetworkRangeName, actual.SubnetworkRangeName) && !dcl.IsZeroValue(desired.SubnetworkRangeName) {
		c.Config.Logger.Infof("Diff in SubnetworkRangeName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SubnetworkRangeName), dcl.SprintResource(actual.SubnetworkRangeName))
		return true
	}
	return false
}

func compareInstanceNetworkInterfacesAliasIPRangesSlice(c *Client, desired, actual []InstanceNetworkInterfacesAliasIPRanges) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceNetworkInterfacesAliasIPRanges, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceNetworkInterfacesAliasIPRanges(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfacesAliasIPRanges, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceNetworkInterfacesAliasIPRangesMap(c *Client, desired, actual map[string]InstanceNetworkInterfacesAliasIPRanges) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceNetworkInterfacesAliasIPRanges, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfacesAliasIPRanges, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceNetworkInterfacesAliasIPRanges(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfacesAliasIPRanges, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceScheduling(c *Client, desired, actual *InstanceScheduling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AutomaticRestart == nil && desired.AutomaticRestart != nil && !dcl.IsEmptyValueIndirect(desired.AutomaticRestart) {
		c.Config.Logger.Infof("desired AutomaticRestart %s - but actually nil", dcl.SprintResource(desired.AutomaticRestart))
		return true
	}
	if !dcl.BoolCanonicalize(desired.AutomaticRestart, actual.AutomaticRestart) && !dcl.IsZeroValue(desired.AutomaticRestart) {
		c.Config.Logger.Infof("Diff in AutomaticRestart. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutomaticRestart), dcl.SprintResource(actual.AutomaticRestart))
		return true
	}
	if actual.OnHostMaintenance == nil && desired.OnHostMaintenance != nil && !dcl.IsEmptyValueIndirect(desired.OnHostMaintenance) {
		c.Config.Logger.Infof("desired OnHostMaintenance %s - but actually nil", dcl.SprintResource(desired.OnHostMaintenance))
		return true
	}
	if !dcl.StringCanonicalize(desired.OnHostMaintenance, actual.OnHostMaintenance) && !dcl.IsZeroValue(desired.OnHostMaintenance) {
		c.Config.Logger.Infof("Diff in OnHostMaintenance. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OnHostMaintenance), dcl.SprintResource(actual.OnHostMaintenance))
		return true
	}
	if actual.Preemptible == nil && desired.Preemptible != nil && !dcl.IsEmptyValueIndirect(desired.Preemptible) {
		c.Config.Logger.Infof("desired Preemptible %s - but actually nil", dcl.SprintResource(desired.Preemptible))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Preemptible, actual.Preemptible) && !dcl.IsZeroValue(desired.Preemptible) {
		c.Config.Logger.Infof("Diff in Preemptible. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Preemptible), dcl.SprintResource(actual.Preemptible))
		return true
	}
	return false
}

func compareInstanceSchedulingSlice(c *Client, desired, actual []InstanceScheduling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceScheduling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceScheduling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceScheduling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSchedulingMap(c *Client, desired, actual map[string]InstanceScheduling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceScheduling, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceScheduling, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceScheduling(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceScheduling, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceServiceAccounts(c *Client, desired, actual *InstanceServiceAccounts) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Email == nil && desired.Email != nil && !dcl.IsEmptyValueIndirect(desired.Email) {
		c.Config.Logger.Infof("desired Email %s - but actually nil", dcl.SprintResource(desired.Email))
		return true
	}
	if !dcl.StringCanonicalize(desired.Email, actual.Email) && !dcl.IsZeroValue(desired.Email) {
		c.Config.Logger.Infof("Diff in Email. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Email), dcl.SprintResource(actual.Email))
		return true
	}
	if actual.Scopes == nil && desired.Scopes != nil && !dcl.IsEmptyValueIndirect(desired.Scopes) {
		c.Config.Logger.Infof("desired Scopes %s - but actually nil", dcl.SprintResource(desired.Scopes))
		return true
	}
	if !dcl.StringSliceEquals(desired.Scopes, actual.Scopes) && !dcl.IsZeroValue(desired.Scopes) {
		c.Config.Logger.Infof("Diff in Scopes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scopes), dcl.SprintResource(actual.Scopes))
		return true
	}
	return false
}

func compareInstanceServiceAccountsSlice(c *Client, desired, actual []InstanceServiceAccounts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceServiceAccounts, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceServiceAccounts(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceServiceAccounts, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceServiceAccountsMap(c *Client, desired, actual map[string]InstanceServiceAccounts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceServiceAccounts, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceServiceAccounts, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceServiceAccounts(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceServiceAccounts, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceShieldedInstanceConfig(c *Client, desired, actual *InstanceShieldedInstanceConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EnableSecureBoot == nil && desired.EnableSecureBoot != nil && !dcl.IsEmptyValueIndirect(desired.EnableSecureBoot) {
		c.Config.Logger.Infof("desired EnableSecureBoot %s - but actually nil", dcl.SprintResource(desired.EnableSecureBoot))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableSecureBoot, actual.EnableSecureBoot) && !dcl.IsZeroValue(desired.EnableSecureBoot) {
		c.Config.Logger.Infof("Diff in EnableSecureBoot. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableSecureBoot), dcl.SprintResource(actual.EnableSecureBoot))
		return true
	}
	if actual.EnableVtpm == nil && desired.EnableVtpm != nil && !dcl.IsEmptyValueIndirect(desired.EnableVtpm) {
		c.Config.Logger.Infof("desired EnableVtpm %s - but actually nil", dcl.SprintResource(desired.EnableVtpm))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableVtpm, actual.EnableVtpm) && !dcl.IsZeroValue(desired.EnableVtpm) {
		c.Config.Logger.Infof("Diff in EnableVtpm. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableVtpm), dcl.SprintResource(actual.EnableVtpm))
		return true
	}
	if actual.EnableIntegrityMonitoring == nil && desired.EnableIntegrityMonitoring != nil && !dcl.IsEmptyValueIndirect(desired.EnableIntegrityMonitoring) {
		c.Config.Logger.Infof("desired EnableIntegrityMonitoring %s - but actually nil", dcl.SprintResource(desired.EnableIntegrityMonitoring))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring) && !dcl.IsZeroValue(desired.EnableIntegrityMonitoring) {
		c.Config.Logger.Infof("Diff in EnableIntegrityMonitoring. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableIntegrityMonitoring), dcl.SprintResource(actual.EnableIntegrityMonitoring))
		return true
	}
	return false
}

func compareInstanceShieldedInstanceConfigSlice(c *Client, desired, actual []InstanceShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceShieldedInstanceConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceShieldedInstanceConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceShieldedInstanceConfigMap(c *Client, desired, actual map[string]InstanceShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceShieldedInstanceConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceShieldedInstanceConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceShieldedInstanceConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceDisksInterfaceEnumSlice(c *Client, desired, actual []InstanceDisksInterfaceEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisksInterfaceEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDisksInterfaceEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDisksInterfaceEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDisksInterfaceEnum(c *Client, desired, actual *InstanceDisksInterfaceEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceDisksModeEnumSlice(c *Client, desired, actual []InstanceDisksModeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisksModeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDisksModeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDisksModeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDisksModeEnum(c *Client, desired, actual *InstanceDisksModeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceDisksTypeEnumSlice(c *Client, desired, actual []InstanceDisksTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDisksTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDisksTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDisksTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDisksTypeEnum(c *Client, desired, actual *InstanceDisksTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceNetworkInterfacesAccessConfigsTypeEnumSlice(c *Client, desired, actual []InstanceNetworkInterfacesAccessConfigsTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceNetworkInterfacesAccessConfigsTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceNetworkInterfacesAccessConfigsTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceNetworkInterfacesAccessConfigsTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceNetworkInterfacesAccessConfigsTypeEnum(c *Client, desired, actual *InstanceNetworkInterfacesAccessConfigsTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceStatusEnumSlice(c *Client, desired, actual []InstanceStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceStatusEnum(c *Client, desired, actual *InstanceStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Instance) urlNormalized() *Instance {
	normalized := deepcopy.Copy(*r).(Instance)
	normalized.CpuPlatform = dcl.SelfLinkToName(r.CpuPlatform)
	normalized.CreationTimestamp = dcl.SelfLinkToName(r.CreationTimestamp)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Hostname = dcl.SelfLinkToName(r.Hostname)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.MachineType = dcl.SelfLinkToName(r.MachineType)
	normalized.MinCpuPlatform = dcl.SelfLinkToName(r.MinCpuPlatform)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.StatusMessage = dcl.SelfLinkToName(r.StatusMessage)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	return &normalized
}

func (r *Instance) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Zone), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Zone)
}

func (r *Instance) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Zone), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "setDeletionProtection" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"zone":    dcl.ValueOrEmptyString(n.Zone),
		}
		return dcl.URL("/projects/{{project}}/zones/{{zone}}/instances/{resourceId}/setDeletionProtection", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "setLabels" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"zone":    dcl.ValueOrEmptyString(n.Zone),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}/setLabels", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "setMachineType" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"zone":    dcl.ValueOrEmptyString(n.Zone),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}/setMachineType", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "setMetadata" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"zone":    dcl.ValueOrEmptyString(n.Zone),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}/setMetadata", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "setTags" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"zone":    dcl.ValueOrEmptyString(n.Zone),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}/setTags", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "start" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"zone":    dcl.ValueOrEmptyString(n.Zone),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}/start", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "stop" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"zone":    dcl.ValueOrEmptyString(n.Zone),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}/stop", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"zone":    dcl.ValueOrEmptyString(n.Zone),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/zones/{{zone}}/instances/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	if updateName == "updateShieldedInstanceConfig" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/instances/{{name}}/updateShieldedInstanceConfig", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

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
	dcl.MoveMapEntry(
		m,
		[]string{"tags"},
		[]string{"tags", "items"},
	)

	return json.Marshal(m)
}

// unmarshalInstance decodes JSON responses into the Instance resource schema.
func unmarshalInstance(b []byte, c *Client) (*Instance, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapInstance(m, c)
}

func unmarshalMapInstance(m map[string]interface{}, c *Client) (*Instance, error) {
	if v, err := dcl.MapFromListOfKeyValues(m, []string{"metadata", "items"}); err != nil {
		return nil, err
	} else {
		dcl.PutMapEntry(
			m,
			[]string{"metadata"},
			v,
		)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"tags", "items"},
		[]string{"tags"},
	)

	return flattenInstance(c, m), nil
}

// expandInstance expands Instance into a JSON request object.
func expandInstance(c *Client, f *Instance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.CanIPForward; !dcl.IsEmptyValueIndirect(v) {
		m["canIpForward"] = v
	}
	if v := f.CpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["cpuPlatform"] = v
	}
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v := f.DeletionProtection; !dcl.IsEmptyValueIndirect(v) {
		m["deletionProtection"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandInstanceDisksSlice(c, f.Disks); err != nil {
		return nil, fmt.Errorf("error expanding Disks into disks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["disks"] = v
	}
	if v, err := expandInstanceGuestAcceleratorsSlice(c, f.GuestAccelerators); err != nil {
		return nil, fmt.Errorf("error expanding GuestAccelerators into guestAccelerators: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["guestAccelerators"] = v
	}
	if v := f.Hostname; !dcl.IsEmptyValueIndirect(v) {
		m["hostname"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := dcl.ListOfKeyValuesFromMapInItemsStruct(f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandInstanceNetworkInterfacesSlice(c, f.NetworkInterfaces); err != nil {
		return nil, fmt.Errorf("error expanding NetworkInterfaces into networkInterfaces: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["networkInterfaces"] = v
	}
	if v, err := expandInstanceScheduling(c, f.Scheduling); err != nil {
		return nil, fmt.Errorf("error expanding Scheduling into scheduling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scheduling"] = v
	}
	if v, err := expandInstanceServiceAccountsSlice(c, f.ServiceAccounts); err != nil {
		return nil, fmt.Errorf("error expanding ServiceAccounts into serviceAccounts: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccounts"] = v
	}
	if v, err := expandInstanceShieldedInstanceConfig(c, f.ShieldedInstanceConfig); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedInstanceConfig into shieldedInstanceConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["shieldedInstanceConfig"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.StatusMessage; !dcl.IsEmptyValueIndirect(v) {
		m["statusMessage"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
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
	r.CanIPForward = dcl.FlattenBool(m["canIpForward"])
	r.CpuPlatform = dcl.FlattenString(m["cpuPlatform"])
	r.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	r.DeletionProtection = dcl.FlattenBool(m["deletionProtection"])
	r.Description = dcl.FlattenString(m["description"])
	r.Disks = flattenInstanceDisksSlice(c, m["disks"])
	r.GuestAccelerators = flattenInstanceGuestAcceleratorsSlice(c, m["guestAccelerators"])
	r.Hostname = dcl.FlattenString(m["hostname"])
	r.Id = dcl.FlattenString(m["id"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Metadata = dcl.FlattenKeyValuePairs(m["metadata"])
	r.MachineType = dcl.FlattenString(m["machineType"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])
	r.Name = dcl.FlattenString(m["name"])
	r.NetworkInterfaces = flattenInstanceNetworkInterfacesSlice(c, m["networkInterfaces"])
	r.Scheduling = flattenInstanceScheduling(c, m["scheduling"])
	r.ServiceAccounts = flattenInstanceServiceAccountsSlice(c, m["serviceAccounts"])
	r.ShieldedInstanceConfig = flattenInstanceShieldedInstanceConfig(c, m["shieldedInstanceConfig"])
	r.Status = flattenInstanceStatusEnum(m["status"])
	r.StatusMessage = dcl.FlattenString(m["statusMessage"])
	r.Tags = dcl.FlattenStringSlice(m["tags"])
	r.Zone = dcl.FlattenString(m["zone"])
	r.Project = dcl.FlattenString(m["project"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])

	return r
}

// expandInstanceDisksMap expands the contents of InstanceDisks into a JSON
// request object.
func expandInstanceDisksMap(c *Client, f map[string]InstanceDisks) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceDisks(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceDisksSlice expands the contents of InstanceDisks into a JSON
// request object.
func expandInstanceDisksSlice(c *Client, f []InstanceDisks) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceDisks(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceDisksMap flattens the contents of InstanceDisks from a JSON
// response object.
func flattenInstanceDisksMap(c *Client, i interface{}) map[string]InstanceDisks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceDisks{}
	}

	if len(a) == 0 {
		return map[string]InstanceDisks{}
	}

	items := make(map[string]InstanceDisks)
	for k, item := range a {
		items[k] = *flattenInstanceDisks(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceDisksSlice flattens the contents of InstanceDisks from a JSON
// response object.
func flattenInstanceDisksSlice(c *Client, i interface{}) []InstanceDisks {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDisks{}
	}

	if len(a) == 0 {
		return []InstanceDisks{}
	}

	items := make([]InstanceDisks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDisks(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceDisks expands an instance of InstanceDisks into a JSON
// request object.
func expandInstanceDisks(c *Client, f *InstanceDisks) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AutoDelete; !dcl.IsEmptyValueIndirect(v) {
		m["autoDelete"] = v
	}
	if v := f.Boot; !dcl.IsEmptyValueIndirect(v) {
		m["boot"] = v
	}
	if v := f.DeviceName; !dcl.IsEmptyValueIndirect(v) {
		m["deviceName"] = v
	}
	if v, err := expandInstanceDisksDiskEncryptionKey(c, f.DiskEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding DiskEncryptionKey into diskEncryptionKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["diskEncryptionKey"] = v
	}
	if v := f.Index; !dcl.IsEmptyValueIndirect(v) {
		m["index"] = v
	}
	if v, err := expandInstanceDisksInitializeParams(c, f.InitializeParams); err != nil {
		return nil, fmt.Errorf("error expanding InitializeParams into initializeParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["initializeParams"] = v
	}
	if v := f.Interface; !dcl.IsEmptyValueIndirect(v) {
		m["interface"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}
	if v := f.Source; !dcl.IsEmptyValueIndirect(v) {
		m["source"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenInstanceDisks flattens an instance of InstanceDisks from a JSON
// response object.
func flattenInstanceDisks(c *Client, i interface{}) *InstanceDisks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceDisks{}
	r.AutoDelete = dcl.FlattenBool(m["autoDelete"])
	r.Boot = dcl.FlattenBool(m["boot"])
	r.DeviceName = dcl.FlattenString(m["deviceName"])
	r.DiskEncryptionKey = flattenInstanceDisksDiskEncryptionKey(c, m["diskEncryptionKey"])
	r.Index = dcl.FlattenInteger(m["index"])
	r.InitializeParams = flattenInstanceDisksInitializeParams(c, m["initializeParams"])
	r.Interface = flattenInstanceDisksInterfaceEnum(m["interface"])
	r.Mode = flattenInstanceDisksModeEnum(m["mode"])
	r.Source = dcl.FlattenString(m["source"])
	r.Type = flattenInstanceDisksTypeEnum(m["type"])

	return r
}

// expandInstanceDisksDiskEncryptionKeyMap expands the contents of InstanceDisksDiskEncryptionKey into a JSON
// request object.
func expandInstanceDisksDiskEncryptionKeyMap(c *Client, f map[string]InstanceDisksDiskEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceDisksDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceDisksDiskEncryptionKeySlice expands the contents of InstanceDisksDiskEncryptionKey into a JSON
// request object.
func expandInstanceDisksDiskEncryptionKeySlice(c *Client, f []InstanceDisksDiskEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceDisksDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceDisksDiskEncryptionKeyMap flattens the contents of InstanceDisksDiskEncryptionKey from a JSON
// response object.
func flattenInstanceDisksDiskEncryptionKeyMap(c *Client, i interface{}) map[string]InstanceDisksDiskEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceDisksDiskEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]InstanceDisksDiskEncryptionKey{}
	}

	items := make(map[string]InstanceDisksDiskEncryptionKey)
	for k, item := range a {
		items[k] = *flattenInstanceDisksDiskEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceDisksDiskEncryptionKeySlice flattens the contents of InstanceDisksDiskEncryptionKey from a JSON
// response object.
func flattenInstanceDisksDiskEncryptionKeySlice(c *Client, i interface{}) []InstanceDisksDiskEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDisksDiskEncryptionKey{}
	}

	if len(a) == 0 {
		return []InstanceDisksDiskEncryptionKey{}
	}

	items := make([]InstanceDisksDiskEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDisksDiskEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceDisksDiskEncryptionKey expands an instance of InstanceDisksDiskEncryptionKey into a JSON
// request object.
func expandInstanceDisksDiskEncryptionKey(c *Client, f *InstanceDisksDiskEncryptionKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RawKey; !dcl.IsEmptyValueIndirect(v) {
		m["rawKey"] = v
	}
	if v := f.RsaEncryptedKey; !dcl.IsEmptyValueIndirect(v) {
		m["rsaEncryptedKey"] = v
	}
	if v := f.Sha256; !dcl.IsEmptyValueIndirect(v) {
		m["sha256"] = v
	}

	return m, nil
}

// flattenInstanceDisksDiskEncryptionKey flattens an instance of InstanceDisksDiskEncryptionKey from a JSON
// response object.
func flattenInstanceDisksDiskEncryptionKey(c *Client, i interface{}) *InstanceDisksDiskEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceDisksDiskEncryptionKey{}
	r.RawKey = dcl.FlattenString(m["rawKey"])
	r.RsaEncryptedKey = dcl.FlattenString(m["rsaEncryptedKey"])
	r.Sha256 = dcl.FlattenString(m["sha256"])

	return r
}

// expandInstanceDisksInitializeParamsMap expands the contents of InstanceDisksInitializeParams into a JSON
// request object.
func expandInstanceDisksInitializeParamsMap(c *Client, f map[string]InstanceDisksInitializeParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceDisksInitializeParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceDisksInitializeParamsSlice expands the contents of InstanceDisksInitializeParams into a JSON
// request object.
func expandInstanceDisksInitializeParamsSlice(c *Client, f []InstanceDisksInitializeParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceDisksInitializeParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceDisksInitializeParamsMap flattens the contents of InstanceDisksInitializeParams from a JSON
// response object.
func flattenInstanceDisksInitializeParamsMap(c *Client, i interface{}) map[string]InstanceDisksInitializeParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceDisksInitializeParams{}
	}

	if len(a) == 0 {
		return map[string]InstanceDisksInitializeParams{}
	}

	items := make(map[string]InstanceDisksInitializeParams)
	for k, item := range a {
		items[k] = *flattenInstanceDisksInitializeParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceDisksInitializeParamsSlice flattens the contents of InstanceDisksInitializeParams from a JSON
// response object.
func flattenInstanceDisksInitializeParamsSlice(c *Client, i interface{}) []InstanceDisksInitializeParams {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDisksInitializeParams{}
	}

	if len(a) == 0 {
		return []InstanceDisksInitializeParams{}
	}

	items := make([]InstanceDisksInitializeParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDisksInitializeParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceDisksInitializeParams expands an instance of InstanceDisksInitializeParams into a JSON
// request object.
func expandInstanceDisksInitializeParams(c *Client, f *InstanceDisksInitializeParams) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DiskName; !dcl.IsEmptyValueIndirect(v) {
		m["diskName"] = v
	}
	if v := f.DiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskSizeGb"] = v
	}
	if v := f.DiskType; !dcl.IsEmptyValueIndirect(v) {
		m["diskType"] = v
	}
	if v := f.SourceImage; !dcl.IsEmptyValueIndirect(v) {
		m["sourceImage"] = v
	}
	if v, err := expandInstanceDisksInitializeParamsSourceImageEncryptionKey(c, f.SourceImageEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding SourceImageEncryptionKey into sourceImageEncryptionKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sourceImageEncryptionKey"] = v
	}

	return m, nil
}

// flattenInstanceDisksInitializeParams flattens an instance of InstanceDisksInitializeParams from a JSON
// response object.
func flattenInstanceDisksInitializeParams(c *Client, i interface{}) *InstanceDisksInitializeParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceDisksInitializeParams{}
	r.DiskName = dcl.FlattenString(m["diskName"])
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.DiskType = dcl.FlattenString(m["diskType"])
	r.SourceImage = dcl.FlattenString(m["sourceImage"])
	r.SourceImageEncryptionKey = flattenInstanceDisksInitializeParamsSourceImageEncryptionKey(c, m["sourceImageEncryptionKey"])

	return r
}

// expandInstanceDisksInitializeParamsSourceImageEncryptionKeyMap expands the contents of InstanceDisksInitializeParamsSourceImageEncryptionKey into a JSON
// request object.
func expandInstanceDisksInitializeParamsSourceImageEncryptionKeyMap(c *Client, f map[string]InstanceDisksInitializeParamsSourceImageEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceDisksInitializeParamsSourceImageEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceDisksInitializeParamsSourceImageEncryptionKeySlice expands the contents of InstanceDisksInitializeParamsSourceImageEncryptionKey into a JSON
// request object.
func expandInstanceDisksInitializeParamsSourceImageEncryptionKeySlice(c *Client, f []InstanceDisksInitializeParamsSourceImageEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceDisksInitializeParamsSourceImageEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceDisksInitializeParamsSourceImageEncryptionKeyMap flattens the contents of InstanceDisksInitializeParamsSourceImageEncryptionKey from a JSON
// response object.
func flattenInstanceDisksInitializeParamsSourceImageEncryptionKeyMap(c *Client, i interface{}) map[string]InstanceDisksInitializeParamsSourceImageEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceDisksInitializeParamsSourceImageEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]InstanceDisksInitializeParamsSourceImageEncryptionKey{}
	}

	items := make(map[string]InstanceDisksInitializeParamsSourceImageEncryptionKey)
	for k, item := range a {
		items[k] = *flattenInstanceDisksInitializeParamsSourceImageEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceDisksInitializeParamsSourceImageEncryptionKeySlice flattens the contents of InstanceDisksInitializeParamsSourceImageEncryptionKey from a JSON
// response object.
func flattenInstanceDisksInitializeParamsSourceImageEncryptionKeySlice(c *Client, i interface{}) []InstanceDisksInitializeParamsSourceImageEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDisksInitializeParamsSourceImageEncryptionKey{}
	}

	if len(a) == 0 {
		return []InstanceDisksInitializeParamsSourceImageEncryptionKey{}
	}

	items := make([]InstanceDisksInitializeParamsSourceImageEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDisksInitializeParamsSourceImageEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceDisksInitializeParamsSourceImageEncryptionKey expands an instance of InstanceDisksInitializeParamsSourceImageEncryptionKey into a JSON
// request object.
func expandInstanceDisksInitializeParamsSourceImageEncryptionKey(c *Client, f *InstanceDisksInitializeParamsSourceImageEncryptionKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RawKey; !dcl.IsEmptyValueIndirect(v) {
		m["rawKey"] = v
	}
	if v := f.Sha256; !dcl.IsEmptyValueIndirect(v) {
		m["sha256"] = v
	}

	return m, nil
}

// flattenInstanceDisksInitializeParamsSourceImageEncryptionKey flattens an instance of InstanceDisksInitializeParamsSourceImageEncryptionKey from a JSON
// response object.
func flattenInstanceDisksInitializeParamsSourceImageEncryptionKey(c *Client, i interface{}) *InstanceDisksInitializeParamsSourceImageEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceDisksInitializeParamsSourceImageEncryptionKey{}
	r.RawKey = dcl.FlattenString(m["rawKey"])
	r.Sha256 = dcl.FlattenString(m["sha256"])

	return r
}

// expandInstanceGuestAcceleratorsMap expands the contents of InstanceGuestAccelerators into a JSON
// request object.
func expandInstanceGuestAcceleratorsMap(c *Client, f map[string]InstanceGuestAccelerators) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGuestAccelerators(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGuestAcceleratorsSlice expands the contents of InstanceGuestAccelerators into a JSON
// request object.
func expandInstanceGuestAcceleratorsSlice(c *Client, f []InstanceGuestAccelerators) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGuestAccelerators(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGuestAcceleratorsMap flattens the contents of InstanceGuestAccelerators from a JSON
// response object.
func flattenInstanceGuestAcceleratorsMap(c *Client, i interface{}) map[string]InstanceGuestAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGuestAccelerators{}
	}

	if len(a) == 0 {
		return map[string]InstanceGuestAccelerators{}
	}

	items := make(map[string]InstanceGuestAccelerators)
	for k, item := range a {
		items[k] = *flattenInstanceGuestAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGuestAcceleratorsSlice flattens the contents of InstanceGuestAccelerators from a JSON
// response object.
func flattenInstanceGuestAcceleratorsSlice(c *Client, i interface{}) []InstanceGuestAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGuestAccelerators{}
	}

	if len(a) == 0 {
		return []InstanceGuestAccelerators{}
	}

	items := make([]InstanceGuestAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGuestAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGuestAccelerators expands an instance of InstanceGuestAccelerators into a JSON
// request object.
func expandInstanceGuestAccelerators(c *Client, f *InstanceGuestAccelerators) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AcceleratorCount; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorCount"] = v
	}
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorType"] = v
	}

	return m, nil
}

// flattenInstanceGuestAccelerators flattens an instance of InstanceGuestAccelerators from a JSON
// response object.
func flattenInstanceGuestAccelerators(c *Client, i interface{}) *InstanceGuestAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGuestAccelerators{}
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])
	r.AcceleratorType = dcl.FlattenString(m["acceleratorType"])

	return r
}

// expandInstanceNetworkInterfacesMap expands the contents of InstanceNetworkInterfaces into a JSON
// request object.
func expandInstanceNetworkInterfacesMap(c *Client, f map[string]InstanceNetworkInterfaces) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceNetworkInterfaces(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceNetworkInterfacesSlice expands the contents of InstanceNetworkInterfaces into a JSON
// request object.
func expandInstanceNetworkInterfacesSlice(c *Client, f []InstanceNetworkInterfaces) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceNetworkInterfaces(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceNetworkInterfacesMap flattens the contents of InstanceNetworkInterfaces from a JSON
// response object.
func flattenInstanceNetworkInterfacesMap(c *Client, i interface{}) map[string]InstanceNetworkInterfaces {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceNetworkInterfaces{}
	}

	if len(a) == 0 {
		return map[string]InstanceNetworkInterfaces{}
	}

	items := make(map[string]InstanceNetworkInterfaces)
	for k, item := range a {
		items[k] = *flattenInstanceNetworkInterfaces(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceNetworkInterfacesSlice flattens the contents of InstanceNetworkInterfaces from a JSON
// response object.
func flattenInstanceNetworkInterfacesSlice(c *Client, i interface{}) []InstanceNetworkInterfaces {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworkInterfaces{}
	}

	if len(a) == 0 {
		return []InstanceNetworkInterfaces{}
	}

	items := make([]InstanceNetworkInterfaces, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworkInterfaces(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceNetworkInterfaces expands an instance of InstanceNetworkInterfaces into a JSON
// request object.
func expandInstanceNetworkInterfaces(c *Client, f *InstanceNetworkInterfaces) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInstanceNetworkInterfacesAccessConfigsSlice(c, f.AccessConfigs); err != nil {
		return nil, fmt.Errorf("error expanding AccessConfigs into accessConfigs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["accessConfigs"] = v
	}
	if v, err := expandInstanceNetworkInterfacesAliasIPRangesSlice(c, f.AliasIPRanges); err != nil {
		return nil, fmt.Errorf("error expanding AliasIPRanges into aliasIPRanges: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aliasIPRanges"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.NetworkIP; !dcl.IsEmptyValueIndirect(v) {
		m["networkIP"] = v
	}
	if v := f.Subnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["subnetwork"] = v
	}

	return m, nil
}

// flattenInstanceNetworkInterfaces flattens an instance of InstanceNetworkInterfaces from a JSON
// response object.
func flattenInstanceNetworkInterfaces(c *Client, i interface{}) *InstanceNetworkInterfaces {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceNetworkInterfaces{}
	r.AccessConfigs = flattenInstanceNetworkInterfacesAccessConfigsSlice(c, m["accessConfigs"])
	r.AliasIPRanges = flattenInstanceNetworkInterfacesAliasIPRangesSlice(c, m["aliasIPRanges"])
	r.Name = dcl.FlattenString(m["name"])
	r.Network = dcl.FlattenString(m["network"])
	r.NetworkIP = dcl.FlattenString(m["networkIP"])
	r.Subnetwork = dcl.FlattenString(m["subnetwork"])

	return r
}

// expandInstanceNetworkInterfacesAccessConfigsMap expands the contents of InstanceNetworkInterfacesAccessConfigs into a JSON
// request object.
func expandInstanceNetworkInterfacesAccessConfigsMap(c *Client, f map[string]InstanceNetworkInterfacesAccessConfigs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceNetworkInterfacesAccessConfigs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceNetworkInterfacesAccessConfigsSlice expands the contents of InstanceNetworkInterfacesAccessConfigs into a JSON
// request object.
func expandInstanceNetworkInterfacesAccessConfigsSlice(c *Client, f []InstanceNetworkInterfacesAccessConfigs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceNetworkInterfacesAccessConfigs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceNetworkInterfacesAccessConfigsMap flattens the contents of InstanceNetworkInterfacesAccessConfigs from a JSON
// response object.
func flattenInstanceNetworkInterfacesAccessConfigsMap(c *Client, i interface{}) map[string]InstanceNetworkInterfacesAccessConfigs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceNetworkInterfacesAccessConfigs{}
	}

	if len(a) == 0 {
		return map[string]InstanceNetworkInterfacesAccessConfigs{}
	}

	items := make(map[string]InstanceNetworkInterfacesAccessConfigs)
	for k, item := range a {
		items[k] = *flattenInstanceNetworkInterfacesAccessConfigs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceNetworkInterfacesAccessConfigsSlice flattens the contents of InstanceNetworkInterfacesAccessConfigs from a JSON
// response object.
func flattenInstanceNetworkInterfacesAccessConfigsSlice(c *Client, i interface{}) []InstanceNetworkInterfacesAccessConfigs {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworkInterfacesAccessConfigs{}
	}

	if len(a) == 0 {
		return []InstanceNetworkInterfacesAccessConfigs{}
	}

	items := make([]InstanceNetworkInterfacesAccessConfigs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworkInterfacesAccessConfigs(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceNetworkInterfacesAccessConfigs expands an instance of InstanceNetworkInterfacesAccessConfigs into a JSON
// request object.
func expandInstanceNetworkInterfacesAccessConfigs(c *Client, f *InstanceNetworkInterfacesAccessConfigs) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.NatIP; !dcl.IsEmptyValueIndirect(v) {
		m["natIP"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenInstanceNetworkInterfacesAccessConfigs flattens an instance of InstanceNetworkInterfacesAccessConfigs from a JSON
// response object.
func flattenInstanceNetworkInterfacesAccessConfigs(c *Client, i interface{}) *InstanceNetworkInterfacesAccessConfigs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceNetworkInterfacesAccessConfigs{}
	r.Name = dcl.FlattenString(m["name"])
	r.NatIP = dcl.FlattenString(m["natIP"])
	r.Type = flattenInstanceNetworkInterfacesAccessConfigsTypeEnum(m["type"])

	return r
}

// expandInstanceNetworkInterfacesAliasIPRangesMap expands the contents of InstanceNetworkInterfacesAliasIPRanges into a JSON
// request object.
func expandInstanceNetworkInterfacesAliasIPRangesMap(c *Client, f map[string]InstanceNetworkInterfacesAliasIPRanges) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceNetworkInterfacesAliasIPRanges(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceNetworkInterfacesAliasIPRangesSlice expands the contents of InstanceNetworkInterfacesAliasIPRanges into a JSON
// request object.
func expandInstanceNetworkInterfacesAliasIPRangesSlice(c *Client, f []InstanceNetworkInterfacesAliasIPRanges) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceNetworkInterfacesAliasIPRanges(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceNetworkInterfacesAliasIPRangesMap flattens the contents of InstanceNetworkInterfacesAliasIPRanges from a JSON
// response object.
func flattenInstanceNetworkInterfacesAliasIPRangesMap(c *Client, i interface{}) map[string]InstanceNetworkInterfacesAliasIPRanges {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceNetworkInterfacesAliasIPRanges{}
	}

	if len(a) == 0 {
		return map[string]InstanceNetworkInterfacesAliasIPRanges{}
	}

	items := make(map[string]InstanceNetworkInterfacesAliasIPRanges)
	for k, item := range a {
		items[k] = *flattenInstanceNetworkInterfacesAliasIPRanges(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceNetworkInterfacesAliasIPRangesSlice flattens the contents of InstanceNetworkInterfacesAliasIPRanges from a JSON
// response object.
func flattenInstanceNetworkInterfacesAliasIPRangesSlice(c *Client, i interface{}) []InstanceNetworkInterfacesAliasIPRanges {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworkInterfacesAliasIPRanges{}
	}

	if len(a) == 0 {
		return []InstanceNetworkInterfacesAliasIPRanges{}
	}

	items := make([]InstanceNetworkInterfacesAliasIPRanges, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworkInterfacesAliasIPRanges(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceNetworkInterfacesAliasIPRanges expands an instance of InstanceNetworkInterfacesAliasIPRanges into a JSON
// request object.
func expandInstanceNetworkInterfacesAliasIPRanges(c *Client, f *InstanceNetworkInterfacesAliasIPRanges) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IPCidrRange; !dcl.IsEmptyValueIndirect(v) {
		m["ipCidrRange"] = v
	}
	if v := f.SubnetworkRangeName; !dcl.IsEmptyValueIndirect(v) {
		m["subnetworkRangeName"] = v
	}

	return m, nil
}

// flattenInstanceNetworkInterfacesAliasIPRanges flattens an instance of InstanceNetworkInterfacesAliasIPRanges from a JSON
// response object.
func flattenInstanceNetworkInterfacesAliasIPRanges(c *Client, i interface{}) *InstanceNetworkInterfacesAliasIPRanges {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceNetworkInterfacesAliasIPRanges{}
	r.IPCidrRange = dcl.FlattenString(m["ipCidrRange"])
	r.SubnetworkRangeName = dcl.FlattenString(m["subnetworkRangeName"])

	return r
}

// expandInstanceSchedulingMap expands the contents of InstanceScheduling into a JSON
// request object.
func expandInstanceSchedulingMap(c *Client, f map[string]InstanceScheduling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceScheduling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSchedulingSlice expands the contents of InstanceScheduling into a JSON
// request object.
func expandInstanceSchedulingSlice(c *Client, f []InstanceScheduling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceScheduling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSchedulingMap flattens the contents of InstanceScheduling from a JSON
// response object.
func flattenInstanceSchedulingMap(c *Client, i interface{}) map[string]InstanceScheduling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceScheduling{}
	}

	if len(a) == 0 {
		return map[string]InstanceScheduling{}
	}

	items := make(map[string]InstanceScheduling)
	for k, item := range a {
		items[k] = *flattenInstanceScheduling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSchedulingSlice flattens the contents of InstanceScheduling from a JSON
// response object.
func flattenInstanceSchedulingSlice(c *Client, i interface{}) []InstanceScheduling {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceScheduling{}
	}

	if len(a) == 0 {
		return []InstanceScheduling{}
	}

	items := make([]InstanceScheduling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceScheduling(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceScheduling expands an instance of InstanceScheduling into a JSON
// request object.
func expandInstanceScheduling(c *Client, f *InstanceScheduling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AutomaticRestart; !dcl.IsEmptyValueIndirect(v) {
		m["automaticRestart"] = v
	}
	if v := f.OnHostMaintenance; !dcl.IsEmptyValueIndirect(v) {
		m["onHostMaintenance"] = v
	}
	if v := f.Preemptible; !dcl.IsEmptyValueIndirect(v) {
		m["preemptible"] = v
	}

	return m, nil
}

// flattenInstanceScheduling flattens an instance of InstanceScheduling from a JSON
// response object.
func flattenInstanceScheduling(c *Client, i interface{}) *InstanceScheduling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceScheduling{}
	r.AutomaticRestart = dcl.FlattenBool(m["automaticRestart"])
	r.OnHostMaintenance = dcl.FlattenString(m["onHostMaintenance"])
	r.Preemptible = dcl.FlattenBool(m["preemptible"])

	return r
}

// expandInstanceServiceAccountsMap expands the contents of InstanceServiceAccounts into a JSON
// request object.
func expandInstanceServiceAccountsMap(c *Client, f map[string]InstanceServiceAccounts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceServiceAccounts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceServiceAccountsSlice expands the contents of InstanceServiceAccounts into a JSON
// request object.
func expandInstanceServiceAccountsSlice(c *Client, f []InstanceServiceAccounts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceServiceAccounts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceServiceAccountsMap flattens the contents of InstanceServiceAccounts from a JSON
// response object.
func flattenInstanceServiceAccountsMap(c *Client, i interface{}) map[string]InstanceServiceAccounts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceServiceAccounts{}
	}

	if len(a) == 0 {
		return map[string]InstanceServiceAccounts{}
	}

	items := make(map[string]InstanceServiceAccounts)
	for k, item := range a {
		items[k] = *flattenInstanceServiceAccounts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceServiceAccountsSlice flattens the contents of InstanceServiceAccounts from a JSON
// response object.
func flattenInstanceServiceAccountsSlice(c *Client, i interface{}) []InstanceServiceAccounts {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceServiceAccounts{}
	}

	if len(a) == 0 {
		return []InstanceServiceAccounts{}
	}

	items := make([]InstanceServiceAccounts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceServiceAccounts(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceServiceAccounts expands an instance of InstanceServiceAccounts into a JSON
// request object.
func expandInstanceServiceAccounts(c *Client, f *InstanceServiceAccounts) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Email; !dcl.IsEmptyValueIndirect(v) {
		m["email"] = v
	}
	if v := f.Scopes; !dcl.IsEmptyValueIndirect(v) {
		m["scopes"] = v
	}

	return m, nil
}

// flattenInstanceServiceAccounts flattens an instance of InstanceServiceAccounts from a JSON
// response object.
func flattenInstanceServiceAccounts(c *Client, i interface{}) *InstanceServiceAccounts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceServiceAccounts{}
	r.Email = dcl.FlattenString(m["email"])
	r.Scopes = dcl.FlattenStringSlice(m["scopes"])

	return r
}

// expandInstanceShieldedInstanceConfigMap expands the contents of InstanceShieldedInstanceConfig into a JSON
// request object.
func expandInstanceShieldedInstanceConfigMap(c *Client, f map[string]InstanceShieldedInstanceConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceShieldedInstanceConfigSlice expands the contents of InstanceShieldedInstanceConfig into a JSON
// request object.
func expandInstanceShieldedInstanceConfigSlice(c *Client, f []InstanceShieldedInstanceConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceShieldedInstanceConfigMap flattens the contents of InstanceShieldedInstanceConfig from a JSON
// response object.
func flattenInstanceShieldedInstanceConfigMap(c *Client, i interface{}) map[string]InstanceShieldedInstanceConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceShieldedInstanceConfig{}
	}

	items := make(map[string]InstanceShieldedInstanceConfig)
	for k, item := range a {
		items[k] = *flattenInstanceShieldedInstanceConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceShieldedInstanceConfigSlice flattens the contents of InstanceShieldedInstanceConfig from a JSON
// response object.
func flattenInstanceShieldedInstanceConfigSlice(c *Client, i interface{}) []InstanceShieldedInstanceConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return []InstanceShieldedInstanceConfig{}
	}

	items := make([]InstanceShieldedInstanceConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceShieldedInstanceConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceShieldedInstanceConfig expands an instance of InstanceShieldedInstanceConfig into a JSON
// request object.
func expandInstanceShieldedInstanceConfig(c *Client, f *InstanceShieldedInstanceConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnableSecureBoot; !dcl.IsEmptyValueIndirect(v) {
		m["enableSecureBoot"] = v
	}
	if v := f.EnableVtpm; !dcl.IsEmptyValueIndirect(v) {
		m["enableVtpm"] = v
	}
	if v := f.EnableIntegrityMonitoring; !dcl.IsEmptyValueIndirect(v) {
		m["enableIntegrityMonitoring"] = v
	}

	return m, nil
}

// flattenInstanceShieldedInstanceConfig flattens an instance of InstanceShieldedInstanceConfig from a JSON
// response object.
func flattenInstanceShieldedInstanceConfig(c *Client, i interface{}) *InstanceShieldedInstanceConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceShieldedInstanceConfig{}
	r.EnableSecureBoot = dcl.FlattenBool(m["enableSecureBoot"])
	r.EnableVtpm = dcl.FlattenBool(m["enableVtpm"])
	r.EnableIntegrityMonitoring = dcl.FlattenBool(m["enableIntegrityMonitoring"])

	return r
}

// flattenInstanceDisksInterfaceEnumSlice flattens the contents of InstanceDisksInterfaceEnum from a JSON
// response object.
func flattenInstanceDisksInterfaceEnumSlice(c *Client, i interface{}) []InstanceDisksInterfaceEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDisksInterfaceEnum{}
	}

	if len(a) == 0 {
		return []InstanceDisksInterfaceEnum{}
	}

	items := make([]InstanceDisksInterfaceEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDisksInterfaceEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceDisksInterfaceEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceDisksInterfaceEnum with the same value as that string.
func flattenInstanceDisksInterfaceEnum(i interface{}) *InstanceDisksInterfaceEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceDisksInterfaceEnumRef("")
	}

	return InstanceDisksInterfaceEnumRef(s)
}

// flattenInstanceDisksModeEnumSlice flattens the contents of InstanceDisksModeEnum from a JSON
// response object.
func flattenInstanceDisksModeEnumSlice(c *Client, i interface{}) []InstanceDisksModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDisksModeEnum{}
	}

	if len(a) == 0 {
		return []InstanceDisksModeEnum{}
	}

	items := make([]InstanceDisksModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDisksModeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceDisksModeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceDisksModeEnum with the same value as that string.
func flattenInstanceDisksModeEnum(i interface{}) *InstanceDisksModeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceDisksModeEnumRef("")
	}

	return InstanceDisksModeEnumRef(s)
}

// flattenInstanceDisksTypeEnumSlice flattens the contents of InstanceDisksTypeEnum from a JSON
// response object.
func flattenInstanceDisksTypeEnumSlice(c *Client, i interface{}) []InstanceDisksTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDisksTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceDisksTypeEnum{}
	}

	items := make([]InstanceDisksTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDisksTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceDisksTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceDisksTypeEnum with the same value as that string.
func flattenInstanceDisksTypeEnum(i interface{}) *InstanceDisksTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceDisksTypeEnumRef("")
	}

	return InstanceDisksTypeEnumRef(s)
}

// flattenInstanceNetworkInterfacesAccessConfigsTypeEnumSlice flattens the contents of InstanceNetworkInterfacesAccessConfigsTypeEnum from a JSON
// response object.
func flattenInstanceNetworkInterfacesAccessConfigsTypeEnumSlice(c *Client, i interface{}) []InstanceNetworkInterfacesAccessConfigsTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworkInterfacesAccessConfigsTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceNetworkInterfacesAccessConfigsTypeEnum{}
	}

	items := make([]InstanceNetworkInterfacesAccessConfigsTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworkInterfacesAccessConfigsTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceNetworkInterfacesAccessConfigsTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceNetworkInterfacesAccessConfigsTypeEnum with the same value as that string.
func flattenInstanceNetworkInterfacesAccessConfigsTypeEnum(i interface{}) *InstanceNetworkInterfacesAccessConfigsTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceNetworkInterfacesAccessConfigsTypeEnumRef("")
	}

	return InstanceNetworkInterfacesAccessConfigsTypeEnumRef(s)
}

// flattenInstanceStatusEnumSlice flattens the contents of InstanceStatusEnum from a JSON
// response object.
func flattenInstanceStatusEnumSlice(c *Client, i interface{}) []InstanceStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceStatusEnum{}
	}

	if len(a) == 0 {
		return []InstanceStatusEnum{}
	}

	items := make([]InstanceStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceStatusEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceStatusEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceStatusEnum with the same value as that string.
func flattenInstanceStatusEnum(i interface{}) *InstanceStatusEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceStatusEnumRef("")
	}

	return InstanceStatusEnumRef(s)
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
		if nr.Zone == nil && ncr.Zone == nil {
			c.Config.Logger.Info("Both Zone fields null - considering equal.")
		} else if nr.Zone == nil || ncr.Zone == nil {
			c.Config.Logger.Info("Only one Zone field is null - considering unequal.")
			return false
		} else if *nr.Zone != *ncr.Zone {
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
