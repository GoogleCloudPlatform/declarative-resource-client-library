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

func (r *InstanceTemplate) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "properties"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Properties) {
		if err := r.Properties.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceTemplateProperties) validate() error {
	if err := dcl.Required(r, "machineType"); err != nil {
		return err
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
	if !dcl.IsEmptyValueIndirect(r.Scheduling) {
		if err := r.Scheduling.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceTemplatePropertiesDisks) validate() error {
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
func (r *InstanceTemplatePropertiesDisksDiskEncryptionKey) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesDisksInitializeParams) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SourceSnapshotEncryptionKey) {
		if err := r.SourceSnapshotEncryptionKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SourceImageEncryptionKey) {
		if err := r.SourceImageEncryptionKey.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesDisksGuestOsFeatures) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesReservationAffinity) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesGuestAccelerators) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesNetworkInterfaces) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	return nil
}
func (r *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesShieldedInstanceConfig) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesScheduling) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesSchedulingNodeAffinities) validate() error {
	return nil
}
func (r *InstanceTemplatePropertiesServiceAccounts) validate() error {
	return nil
}

func instanceTemplateGetURL(userBasePath string, r *InstanceTemplate) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/instanceTemplates/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func instanceTemplateListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/instanceTemplates", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func instanceTemplateCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/instanceTemplates", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func instanceTemplateDeleteURL(userBasePath string, r *InstanceTemplate) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/instanceTemplates/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// instanceTemplateApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceTemplateApiOperation interface {
	do(context.Context, *InstanceTemplate, *Client) error
}

// newUpdateInstanceTemplateUpdateRequest creates a request for an
// InstanceTemplate resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceTemplateUpdateRequest(ctx context.Context, f *InstanceTemplate, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateInstanceTemplateUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceTemplateUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceTemplateUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceTemplateUpdateOperation) do(ctx context.Context, r *InstanceTemplate, c *Client) error {
	_, err := c.GetInstanceTemplate(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceTemplateUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceTemplateUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInstanceTemplateRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := instanceTemplateListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InstanceTemplateMaxPage {
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

type listInstanceTemplateOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listInstanceTemplate(ctx context.Context, project, pageToken string, pageSize int32) ([]*InstanceTemplate, string, error) {
	b, err := c.listInstanceTemplateRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInstanceTemplateOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*InstanceTemplate
	for _, v := range m.Items {
		res := flattenInstanceTemplate(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInstanceTemplate(ctx context.Context, f func(*InstanceTemplate) bool, resources []*InstanceTemplate) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInstanceTemplate(ctx, res)
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

type deleteInstanceTemplateOperation struct{}

func (op *deleteInstanceTemplateOperation) do(ctx context.Context, r *InstanceTemplate, c *Client) error {

	_, err := c.GetInstanceTemplate(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("InstanceTemplate not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetInstanceTemplate checking for existence. error: %v", err)
		return err
	}

	u, err := instanceTemplateDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET"); err != nil {
		return err
	}
	_, err = c.GetInstanceTemplate(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createInstanceTemplateOperation struct{}

func (op *createInstanceTemplateOperation) do(ctx context.Context, r *InstanceTemplate, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := instanceTemplateCreateURL(c.Config.BasePath, project)

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
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetInstanceTemplate(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getInstanceTemplateRaw(ctx context.Context, r *InstanceTemplate) ([]byte, error) {

	u, err := instanceTemplateGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) instanceTemplateDiffsForRawDesired(ctx context.Context, rawDesired *InstanceTemplate, opts ...dcl.ApplyOption) (initial, desired *InstanceTemplate, diffs []instanceTemplateDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *InstanceTemplate
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*InstanceTemplate); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected InstanceTemplate, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInstanceTemplate(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a InstanceTemplate resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve InstanceTemplate resource: %v", err)
		}
		c.Config.Logger.Info("Found that InstanceTemplate resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInstanceTemplateDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for InstanceTemplate: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for InstanceTemplate: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInstanceTemplateInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for InstanceTemplate: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInstanceTemplateDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for InstanceTemplate: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInstanceTemplate(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInstanceTemplateInitialState(rawInitial, rawDesired *InstanceTemplate) (*InstanceTemplate, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInstanceTemplateDesiredState(rawDesired, rawInitial *InstanceTemplate, opts ...dcl.ApplyOption) (*InstanceTemplate, error) {

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*InstanceTemplate); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected InstanceTemplate, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Properties = canonicalizeInstanceTemplateProperties(rawDesired.Properties, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.CreationTimestamp) {
		rawDesired.CreationTimestamp = rawInitial.CreationTimestamp
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.IsZeroValue(rawDesired.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.Properties = canonicalizeInstanceTemplateProperties(rawDesired.Properties, rawInitial.Properties, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeInstanceTemplateNewState(c *Client, rawNew, rawDesired *InstanceTemplate) (*InstanceTemplate, error) {

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Properties) && dcl.IsEmptyValueIndirect(rawDesired.Properties) {
		rawNew.Properties = rawDesired.Properties
	} else {
		rawNew.Properties = canonicalizeNewInstanceTemplateProperties(c, rawDesired.Properties, rawNew.Properties)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeInstanceTemplateProperties(des, initial *InstanceTemplateProperties, opts ...dcl.ApplyOption) *InstanceTemplateProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.CanIPForward) {
		des.CanIPForward = initial.CanIPForward
	}
	if dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.Disks) {
		des.Disks = initial.Disks
	}
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.NameToSelfLink(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		des.MachineType = initial.MachineType
	}
	if dcl.IsZeroValue(des.MinCpuPlatform) {
		des.MinCpuPlatform = initial.MinCpuPlatform
	}
	if dcl.IsZeroValue(des.Metadata) {
		des.Metadata = initial.Metadata
	}
	des.ReservationAffinity = canonicalizeInstanceTemplatePropertiesReservationAffinity(des.ReservationAffinity, initial.ReservationAffinity, opts...)
	if dcl.IsZeroValue(des.GuestAccelerators) {
		des.GuestAccelerators = initial.GuestAccelerators
	}
	if dcl.IsZeroValue(des.NetworkInterfaces) {
		des.NetworkInterfaces = initial.NetworkInterfaces
	}
	des.ShieldedInstanceConfig = canonicalizeInstanceTemplatePropertiesShieldedInstanceConfig(des.ShieldedInstanceConfig, initial.ShieldedInstanceConfig, opts...)
	des.Scheduling = canonicalizeInstanceTemplatePropertiesScheduling(des.Scheduling, initial.Scheduling, opts...)
	if dcl.IsZeroValue(des.ServiceAccounts) {
		des.ServiceAccounts = initial.ServiceAccounts
	}
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	}

	return des
}

func canonicalizeNewInstanceTemplateProperties(c *Client, des, nw *InstanceTemplateProperties) *InstanceTemplateProperties {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.MachineType, nw.MachineType) || dcl.IsZeroValue(des.MachineType) {
		nw.MachineType = des.MachineType
	}
	nw.ReservationAffinity = canonicalizeNewInstanceTemplatePropertiesReservationAffinity(c, des.ReservationAffinity, nw.ReservationAffinity)
	nw.ShieldedInstanceConfig = canonicalizeNewInstanceTemplatePropertiesShieldedInstanceConfig(c, des.ShieldedInstanceConfig, nw.ShieldedInstanceConfig)
	nw.Scheduling = canonicalizeNewInstanceTemplatePropertiesScheduling(c, des.Scheduling, nw.Scheduling)

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesSet(c *Client, des, nw []InstanceTemplateProperties) []InstanceTemplateProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplateProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplateProperties(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesDisks(des, initial *InstanceTemplatePropertiesDisks, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AutoDelete) {
		des.AutoDelete = initial.AutoDelete
	}
	if dcl.IsZeroValue(des.Boot) {
		des.Boot = initial.Boot
	}
	if dcl.IsZeroValue(des.DeviceName) {
		des.DeviceName = initial.DeviceName
	}
	des.DiskEncryptionKey = canonicalizeInstanceTemplatePropertiesDisksDiskEncryptionKey(des.DiskEncryptionKey, initial.DiskEncryptionKey, opts...)
	if dcl.IsZeroValue(des.Index) {
		des.Index = initial.Index
	}
	des.InitializeParams = canonicalizeInstanceTemplatePropertiesDisksInitializeParams(des.InitializeParams, initial.InitializeParams, opts...)
	if dcl.IsZeroValue(des.GuestOsFeatures) {
		des.GuestOsFeatures = initial.GuestOsFeatures
	}
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

func canonicalizeNewInstanceTemplatePropertiesDisks(c *Client, des, nw *InstanceTemplatePropertiesDisks) *InstanceTemplatePropertiesDisks {
	if des == nil || nw == nil {
		return nw
	}

	nw.DiskEncryptionKey = canonicalizeNewInstanceTemplatePropertiesDisksDiskEncryptionKey(c, des.DiskEncryptionKey, nw.DiskEncryptionKey)
	nw.InitializeParams = canonicalizeNewInstanceTemplatePropertiesDisksInitializeParams(c, des.InitializeParams, nw.InitializeParams)
	if dcl.NameToSelfLink(des.Source, nw.Source) || dcl.IsZeroValue(des.Source) {
		nw.Source = des.Source
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesDisksSet(c *Client, des, nw []InstanceTemplatePropertiesDisks) []InstanceTemplatePropertiesDisks {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesDisks
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesDisks(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesDisksDiskEncryptionKey(des, initial *InstanceTemplatePropertiesDisksDiskEncryptionKey, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksDiskEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.IsZeroValue(des.RsaEncryptedKey) {
		des.RsaEncryptedKey = initial.RsaEncryptedKey
	}
	if dcl.IsZeroValue(des.Sha256) {
		des.Sha256 = initial.Sha256
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksDiskEncryptionKey(c *Client, des, nw *InstanceTemplatePropertiesDisksDiskEncryptionKey) *InstanceTemplatePropertiesDisksDiskEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesDisksDiskEncryptionKeySet(c *Client, des, nw []InstanceTemplatePropertiesDisksDiskEncryptionKey) []InstanceTemplatePropertiesDisksDiskEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesDisksDiskEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesDisksDiskEncryptionKey(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesDisksInitializeParams(des, initial *InstanceTemplatePropertiesDisksInitializeParams, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksInitializeParams {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DiskName) {
		des.DiskName = initial.DiskName
	}
	if dcl.IsZeroValue(des.DiskSizeGb) {
		des.DiskSizeGb = initial.DiskSizeGb
	}
	if dcl.NameToSelfLink(des.DiskType, initial.DiskType) || dcl.IsZeroValue(des.DiskType) {
		des.DiskType = initial.DiskType
	}
	if dcl.IsZeroValue(des.SourceImage) {
		des.SourceImage = initial.SourceImage
	}
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.SourceSnapshot) {
		des.SourceSnapshot = initial.SourceSnapshot
	}
	des.SourceSnapshotEncryptionKey = canonicalizeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(des.SourceSnapshotEncryptionKey, initial.SourceSnapshotEncryptionKey, opts...)
	if dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.ResourcePolicies) {
		des.ResourcePolicies = initial.ResourcePolicies
	}
	if dcl.IsZeroValue(des.OnUpdateAction) {
		des.OnUpdateAction = initial.OnUpdateAction
	}
	des.SourceImageEncryptionKey = canonicalizeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(des.SourceImageEncryptionKey, initial.SourceImageEncryptionKey, opts...)

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParams(c *Client, des, nw *InstanceTemplatePropertiesDisksInitializeParams) *InstanceTemplatePropertiesDisksInitializeParams {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.DiskType, nw.DiskType) || dcl.IsZeroValue(des.DiskType) {
		nw.DiskType = des.DiskType
	}
	nw.SourceSnapshotEncryptionKey = canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, des.SourceSnapshotEncryptionKey, nw.SourceSnapshotEncryptionKey)
	nw.SourceImageEncryptionKey = canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, des.SourceImageEncryptionKey, nw.SourceImageEncryptionKey)

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSet(c *Client, des, nw []InstanceTemplatePropertiesDisksInitializeParams) []InstanceTemplatePropertiesDisksInitializeParams {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesDisksInitializeParams
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesDisksInitializeParams(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(des, initial *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.IsZeroValue(des.Sha256) {
		des.Sha256 = initial.Sha256
	}
	if dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c *Client, des, nw *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeySet(c *Client, des, nw []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(des, initial *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RawKey) {
		des.RawKey = initial.RawKey
	}
	if dcl.IsZeroValue(des.Sha256) {
		des.Sha256 = initial.Sha256
	}
	if dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c *Client, des, nw *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeySet(c *Client, des, nw []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesDisksGuestOsFeatures(des, initial *InstanceTemplatePropertiesDisksGuestOsFeatures, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksGuestOsFeatures {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksGuestOsFeatures(c *Client, des, nw *InstanceTemplatePropertiesDisksGuestOsFeatures) *InstanceTemplatePropertiesDisksGuestOsFeatures {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesDisksGuestOsFeaturesSet(c *Client, des, nw []InstanceTemplatePropertiesDisksGuestOsFeatures) []InstanceTemplatePropertiesDisksGuestOsFeatures {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesDisksGuestOsFeatures
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesDisksGuestOsFeatures(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesReservationAffinity(des, initial *InstanceTemplatePropertiesReservationAffinity, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesReservationAffinity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesReservationAffinity(c *Client, des, nw *InstanceTemplatePropertiesReservationAffinity) *InstanceTemplatePropertiesReservationAffinity {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesReservationAffinitySet(c *Client, des, nw []InstanceTemplatePropertiesReservationAffinity) []InstanceTemplatePropertiesReservationAffinity {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesReservationAffinity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesReservationAffinity(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesGuestAccelerators(des, initial *InstanceTemplatePropertiesGuestAccelerators, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesGuestAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AcceleratorCount) {
		des.AcceleratorCount = initial.AcceleratorCount
	}
	if dcl.IsZeroValue(des.AcceleratorType) {
		des.AcceleratorType = initial.AcceleratorType
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesGuestAccelerators(c *Client, des, nw *InstanceTemplatePropertiesGuestAccelerators) *InstanceTemplatePropertiesGuestAccelerators {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesGuestAcceleratorsSet(c *Client, des, nw []InstanceTemplatePropertiesGuestAccelerators) []InstanceTemplatePropertiesGuestAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesGuestAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesGuestAccelerators(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesNetworkInterfaces(des, initial *InstanceTemplatePropertiesNetworkInterfaces, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesNetworkInterfaces {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
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
	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.NameToSelfLink(des.Network, initial.Network) || dcl.IsZeroValue(des.Network) {
		des.Network = initial.Network
	}
	if dcl.IsZeroValue(des.NetworkIP) {
		des.NetworkIP = initial.NetworkIP
	}
	if dcl.NameToSelfLink(des.Subnetwork, initial.Subnetwork) || dcl.IsZeroValue(des.Subnetwork) {
		des.Subnetwork = initial.Subnetwork
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfaces(c *Client, des, nw *InstanceTemplatePropertiesNetworkInterfaces) *InstanceTemplatePropertiesNetworkInterfaces {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Network, nw.Network) || dcl.IsZeroValue(des.Network) {
		nw.Network = des.Network
	}
	if dcl.NameToSelfLink(des.Subnetwork, nw.Subnetwork) || dcl.IsZeroValue(des.Subnetwork) {
		nw.Subnetwork = des.Subnetwork
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesSet(c *Client, des, nw []InstanceTemplatePropertiesNetworkInterfaces) []InstanceTemplatePropertiesNetworkInterfaces {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesNetworkInterfaces
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesNetworkInterfaces(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(des, initial *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.NameToSelfLink(des.NatIP, initial.NatIP) || dcl.IsZeroValue(des.NatIP) {
		des.NatIP = initial.NatIP
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.IsZeroValue(des.SetPublicPtr) {
		des.SetPublicPtr = initial.SetPublicPtr
	}
	if dcl.IsZeroValue(des.PublicPtrDomainName) {
		des.PublicPtrDomainName = initial.PublicPtrDomainName
	}
	if dcl.IsZeroValue(des.NetworkTier) {
		des.NetworkTier = initial.NetworkTier
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c *Client, des, nw *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.NatIP, nw.NatIP) || dcl.IsZeroValue(des.NatIP) {
		nw.NatIP = des.NatIP
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSet(c *Client, des, nw []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(des, initial *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.IPCidrRange) {
		des.IPCidrRange = initial.IPCidrRange
	}
	if dcl.IsZeroValue(des.SubnetworkRangeName) {
		des.SubnetworkRangeName = initial.SubnetworkRangeName
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c *Client, des, nw *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSet(c *Client, des, nw []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesShieldedInstanceConfig(des, initial *InstanceTemplatePropertiesShieldedInstanceConfig, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesShieldedInstanceConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.EnableSecureBoot) {
		des.EnableSecureBoot = initial.EnableSecureBoot
	}
	if dcl.IsZeroValue(des.EnableVtpm) {
		des.EnableVtpm = initial.EnableVtpm
	}
	if dcl.IsZeroValue(des.EnableIntegrityMonitoring) {
		des.EnableIntegrityMonitoring = initial.EnableIntegrityMonitoring
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesShieldedInstanceConfig(c *Client, des, nw *InstanceTemplatePropertiesShieldedInstanceConfig) *InstanceTemplatePropertiesShieldedInstanceConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesShieldedInstanceConfigSet(c *Client, des, nw []InstanceTemplatePropertiesShieldedInstanceConfig) []InstanceTemplatePropertiesShieldedInstanceConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesShieldedInstanceConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesShieldedInstanceConfig(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesScheduling(des, initial *InstanceTemplatePropertiesScheduling, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesScheduling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AutomaticRestart) {
		des.AutomaticRestart = initial.AutomaticRestart
	}
	if dcl.IsZeroValue(des.OnHostMaintenance) {
		des.OnHostMaintenance = initial.OnHostMaintenance
	}
	if dcl.IsZeroValue(des.Preemptible) {
		des.Preemptible = initial.Preemptible
	}
	if dcl.IsZeroValue(des.NodeAffinities) {
		des.NodeAffinities = initial.NodeAffinities
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesScheduling(c *Client, des, nw *InstanceTemplatePropertiesScheduling) *InstanceTemplatePropertiesScheduling {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesSchedulingSet(c *Client, des, nw []InstanceTemplatePropertiesScheduling) []InstanceTemplatePropertiesScheduling {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesScheduling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesScheduling(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesSchedulingNodeAffinities(des, initial *InstanceTemplatePropertiesSchedulingNodeAffinities, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesSchedulingNodeAffinities {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Operator) {
		des.Operator = initial.Operator
	}
	if dcl.IsZeroValue(des.Values) {
		des.Values = initial.Values
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesSchedulingNodeAffinities(c *Client, des, nw *InstanceTemplatePropertiesSchedulingNodeAffinities) *InstanceTemplatePropertiesSchedulingNodeAffinities {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesSchedulingNodeAffinitiesSet(c *Client, des, nw []InstanceTemplatePropertiesSchedulingNodeAffinities) []InstanceTemplatePropertiesSchedulingNodeAffinities {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesSchedulingNodeAffinities
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesSchedulingNodeAffinities(c, &d, &n) {
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

func canonicalizeInstanceTemplatePropertiesServiceAccounts(des, initial *InstanceTemplatePropertiesServiceAccounts, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesServiceAccounts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*InstanceTemplate)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Email) {
		des.Email = initial.Email
	}
	if dcl.IsZeroValue(des.Scopes) {
		des.Scopes = initial.Scopes
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesServiceAccounts(c *Client, des, nw *InstanceTemplatePropertiesServiceAccounts) *InstanceTemplatePropertiesServiceAccounts {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceTemplatePropertiesServiceAccountsSet(c *Client, des, nw []InstanceTemplatePropertiesServiceAccounts) []InstanceTemplatePropertiesServiceAccounts {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesServiceAccounts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceTemplatePropertiesServiceAccounts(c, &d, &n) {
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

type instanceTemplateDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceTemplateApiOperation
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
func diffInstanceTemplate(c *Client, desired, actual *InstanceTemplate, opts ...dcl.ApplyOption) ([]instanceTemplateDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []instanceTemplateDiff
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, instanceTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.SelfLink) && (dcl.IsZeroValue(actual.SelfLink) || !reflect.DeepEqual(*desired.SelfLink, *actual.SelfLink)) {
		c.Config.Logger.Infof("Detected diff in SelfLink.\nDESIRED: %v\nACTUAL: %v", desired.SelfLink, actual.SelfLink)
		diffs = append(diffs, instanceTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "SelfLink",
		})
	}
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, instanceTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if compareInstanceTemplateProperties(c, desired.Properties, actual.Properties) {
		c.Config.Logger.Infof("Detected diff in Properties.\nDESIRED: %v\nACTUAL: %v", desired.Properties, actual.Properties)
		diffs = append(diffs, instanceTemplateDiff{
			RequiresRecreate: true,
			FieldName:        "Properties",
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
	var deduped []instanceTemplateDiff
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
func compareInstanceTemplatePropertiesSlice(c *Client, desired, actual []InstanceTemplateProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplateProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplateProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplateProperties, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplateProperties(c *Client, desired, actual *InstanceTemplateProperties) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.CanIPForward == nil && desired.CanIPForward != nil && !dcl.IsEmptyValueIndirect(desired.CanIPForward) {
		c.Config.Logger.Infof("desired CanIPForward %s - but actually nil", dcl.SprintResource(desired.CanIPForward))
		return true
	}
	if !reflect.DeepEqual(desired.CanIPForward, actual.CanIPForward) && !dcl.IsZeroValue(desired.CanIPForward) && !(dcl.IsEmptyValueIndirect(desired.CanIPForward) && dcl.IsZeroValue(actual.CanIPForward)) {
		c.Config.Logger.Infof("Diff in CanIPForward. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CanIPForward), dcl.SprintResource(actual.CanIPForward))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !reflect.DeepEqual(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) && !(dcl.IsEmptyValueIndirect(desired.Description) && dcl.IsZeroValue(actual.Description)) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.Disks == nil && desired.Disks != nil && !dcl.IsEmptyValueIndirect(desired.Disks) {
		c.Config.Logger.Infof("desired Disks %s - but actually nil", dcl.SprintResource(desired.Disks))
		return true
	}
	if compareInstanceTemplatePropertiesDisksSlice(c, desired.Disks, actual.Disks) && !dcl.IsZeroValue(desired.Disks) {
		c.Config.Logger.Infof("Diff in Disks. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disks), dcl.SprintResource(actual.Disks))
		return true
	}
	if actual.Labels == nil && desired.Labels != nil && !dcl.IsEmptyValueIndirect(desired.Labels) {
		c.Config.Logger.Infof("desired Labels %s - but actually nil", dcl.SprintResource(desired.Labels))
		return true
	}
	if !reflect.DeepEqual(desired.Labels, actual.Labels) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	if actual.MachineType == nil && desired.MachineType != nil && !dcl.IsEmptyValueIndirect(desired.MachineType) {
		c.Config.Logger.Infof("desired MachineType %s - but actually nil", dcl.SprintResource(desired.MachineType))
		return true
	}
	if !dcl.NameToSelfLink(desired.MachineType, actual.MachineType) && !dcl.IsZeroValue(desired.MachineType) {
		c.Config.Logger.Infof("Diff in MachineType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MachineType), dcl.SprintResource(actual.MachineType))
		return true
	}
	if actual.MinCpuPlatform == nil && desired.MinCpuPlatform != nil && !dcl.IsEmptyValueIndirect(desired.MinCpuPlatform) {
		c.Config.Logger.Infof("desired MinCpuPlatform %s - but actually nil", dcl.SprintResource(desired.MinCpuPlatform))
		return true
	}
	if !reflect.DeepEqual(desired.MinCpuPlatform, actual.MinCpuPlatform) && !dcl.IsZeroValue(desired.MinCpuPlatform) && !(dcl.IsEmptyValueIndirect(desired.MinCpuPlatform) && dcl.IsZeroValue(actual.MinCpuPlatform)) {
		c.Config.Logger.Infof("Diff in MinCpuPlatform. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinCpuPlatform), dcl.SprintResource(actual.MinCpuPlatform))
		return true
	}
	if actual.Metadata == nil && desired.Metadata != nil && !dcl.IsEmptyValueIndirect(desired.Metadata) {
		c.Config.Logger.Infof("desired Metadata %s - but actually nil", dcl.SprintResource(desired.Metadata))
		return true
	}
	if !reflect.DeepEqual(desired.Metadata, actual.Metadata) && !dcl.IsZeroValue(desired.Metadata) {
		c.Config.Logger.Infof("Diff in Metadata. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Metadata), dcl.SprintResource(actual.Metadata))
		return true
	}
	if actual.ReservationAffinity == nil && desired.ReservationAffinity != nil && !dcl.IsEmptyValueIndirect(desired.ReservationAffinity) {
		c.Config.Logger.Infof("desired ReservationAffinity %s - but actually nil", dcl.SprintResource(desired.ReservationAffinity))
		return true
	}
	if compareInstanceTemplatePropertiesReservationAffinity(c, desired.ReservationAffinity, actual.ReservationAffinity) && !dcl.IsZeroValue(desired.ReservationAffinity) {
		c.Config.Logger.Infof("Diff in ReservationAffinity. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReservationAffinity), dcl.SprintResource(actual.ReservationAffinity))
		return true
	}
	if actual.GuestAccelerators == nil && desired.GuestAccelerators != nil && !dcl.IsEmptyValueIndirect(desired.GuestAccelerators) {
		c.Config.Logger.Infof("desired GuestAccelerators %s - but actually nil", dcl.SprintResource(desired.GuestAccelerators))
		return true
	}
	if compareInstanceTemplatePropertiesGuestAcceleratorsSlice(c, desired.GuestAccelerators, actual.GuestAccelerators) && !dcl.IsZeroValue(desired.GuestAccelerators) {
		c.Config.Logger.Infof("Diff in GuestAccelerators. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GuestAccelerators), dcl.SprintResource(actual.GuestAccelerators))
		return true
	}
	if actual.NetworkInterfaces == nil && desired.NetworkInterfaces != nil && !dcl.IsEmptyValueIndirect(desired.NetworkInterfaces) {
		c.Config.Logger.Infof("desired NetworkInterfaces %s - but actually nil", dcl.SprintResource(desired.NetworkInterfaces))
		return true
	}
	if compareInstanceTemplatePropertiesNetworkInterfacesSlice(c, desired.NetworkInterfaces, actual.NetworkInterfaces) && !dcl.IsZeroValue(desired.NetworkInterfaces) {
		c.Config.Logger.Infof("Diff in NetworkInterfaces. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NetworkInterfaces), dcl.SprintResource(actual.NetworkInterfaces))
		return true
	}
	if actual.ShieldedInstanceConfig == nil && desired.ShieldedInstanceConfig != nil && !dcl.IsEmptyValueIndirect(desired.ShieldedInstanceConfig) {
		c.Config.Logger.Infof("desired ShieldedInstanceConfig %s - but actually nil", dcl.SprintResource(desired.ShieldedInstanceConfig))
		return true
	}
	if compareInstanceTemplatePropertiesShieldedInstanceConfig(c, desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig) && !dcl.IsZeroValue(desired.ShieldedInstanceConfig) {
		c.Config.Logger.Infof("Diff in ShieldedInstanceConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ShieldedInstanceConfig), dcl.SprintResource(actual.ShieldedInstanceConfig))
		return true
	}
	if actual.Scheduling == nil && desired.Scheduling != nil && !dcl.IsEmptyValueIndirect(desired.Scheduling) {
		c.Config.Logger.Infof("desired Scheduling %s - but actually nil", dcl.SprintResource(desired.Scheduling))
		return true
	}
	if compareInstanceTemplatePropertiesScheduling(c, desired.Scheduling, actual.Scheduling) && !dcl.IsZeroValue(desired.Scheduling) {
		c.Config.Logger.Infof("Diff in Scheduling. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scheduling), dcl.SprintResource(actual.Scheduling))
		return true
	}
	if actual.ServiceAccounts == nil && desired.ServiceAccounts != nil && !dcl.IsEmptyValueIndirect(desired.ServiceAccounts) {
		c.Config.Logger.Infof("desired ServiceAccounts %s - but actually nil", dcl.SprintResource(desired.ServiceAccounts))
		return true
	}
	if compareInstanceTemplatePropertiesServiceAccountsSlice(c, desired.ServiceAccounts, actual.ServiceAccounts) && !dcl.IsZeroValue(desired.ServiceAccounts) {
		c.Config.Logger.Infof("Diff in ServiceAccounts. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccounts), dcl.SprintResource(actual.ServiceAccounts))
		return true
	}
	if actual.Tags == nil && desired.Tags != nil && !dcl.IsEmptyValueIndirect(desired.Tags) {
		c.Config.Logger.Infof("desired Tags %s - but actually nil", dcl.SprintResource(desired.Tags))
		return true
	}
	if !dcl.SliceEquals(desired.Tags, actual.Tags) && !dcl.IsZeroValue(desired.Tags) {
		c.Config.Logger.Infof("Diff in Tags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tags), dcl.SprintResource(actual.Tags))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesDisksSlice(c *Client, desired, actual []InstanceTemplatePropertiesDisks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesDisks, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesDisks(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesDisks, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesDisks(c *Client, desired, actual *InstanceTemplatePropertiesDisks) bool {
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
	if !reflect.DeepEqual(desired.AutoDelete, actual.AutoDelete) && !dcl.IsZeroValue(desired.AutoDelete) && !(dcl.IsEmptyValueIndirect(desired.AutoDelete) && dcl.IsZeroValue(actual.AutoDelete)) {
		c.Config.Logger.Infof("Diff in AutoDelete. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoDelete), dcl.SprintResource(actual.AutoDelete))
		return true
	}
	if actual.Boot == nil && desired.Boot != nil && !dcl.IsEmptyValueIndirect(desired.Boot) {
		c.Config.Logger.Infof("desired Boot %s - but actually nil", dcl.SprintResource(desired.Boot))
		return true
	}
	if !reflect.DeepEqual(desired.Boot, actual.Boot) && !dcl.IsZeroValue(desired.Boot) && !(dcl.IsEmptyValueIndirect(desired.Boot) && dcl.IsZeroValue(actual.Boot)) {
		c.Config.Logger.Infof("Diff in Boot. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Boot), dcl.SprintResource(actual.Boot))
		return true
	}
	if actual.DeviceName == nil && desired.DeviceName != nil && !dcl.IsEmptyValueIndirect(desired.DeviceName) {
		c.Config.Logger.Infof("desired DeviceName %s - but actually nil", dcl.SprintResource(desired.DeviceName))
		return true
	}
	if !reflect.DeepEqual(desired.DeviceName, actual.DeviceName) && !dcl.IsZeroValue(desired.DeviceName) && !(dcl.IsEmptyValueIndirect(desired.DeviceName) && dcl.IsZeroValue(actual.DeviceName)) {
		c.Config.Logger.Infof("Diff in DeviceName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DeviceName), dcl.SprintResource(actual.DeviceName))
		return true
	}
	if actual.DiskEncryptionKey == nil && desired.DiskEncryptionKey != nil && !dcl.IsEmptyValueIndirect(desired.DiskEncryptionKey) {
		c.Config.Logger.Infof("desired DiskEncryptionKey %s - but actually nil", dcl.SprintResource(desired.DiskEncryptionKey))
		return true
	}
	if compareInstanceTemplatePropertiesDisksDiskEncryptionKey(c, desired.DiskEncryptionKey, actual.DiskEncryptionKey) && !dcl.IsZeroValue(desired.DiskEncryptionKey) {
		c.Config.Logger.Infof("Diff in DiskEncryptionKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskEncryptionKey), dcl.SprintResource(actual.DiskEncryptionKey))
		return true
	}
	if actual.Index == nil && desired.Index != nil && !dcl.IsEmptyValueIndirect(desired.Index) {
		c.Config.Logger.Infof("desired Index %s - but actually nil", dcl.SprintResource(desired.Index))
		return true
	}
	if !reflect.DeepEqual(desired.Index, actual.Index) && !dcl.IsZeroValue(desired.Index) && !(dcl.IsEmptyValueIndirect(desired.Index) && dcl.IsZeroValue(actual.Index)) {
		c.Config.Logger.Infof("Diff in Index. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Index), dcl.SprintResource(actual.Index))
		return true
	}
	if actual.InitializeParams == nil && desired.InitializeParams != nil && !dcl.IsEmptyValueIndirect(desired.InitializeParams) {
		c.Config.Logger.Infof("desired InitializeParams %s - but actually nil", dcl.SprintResource(desired.InitializeParams))
		return true
	}
	if compareInstanceTemplatePropertiesDisksInitializeParams(c, desired.InitializeParams, actual.InitializeParams) && !dcl.IsZeroValue(desired.InitializeParams) {
		c.Config.Logger.Infof("Diff in InitializeParams. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InitializeParams), dcl.SprintResource(actual.InitializeParams))
		return true
	}
	if actual.GuestOsFeatures == nil && desired.GuestOsFeatures != nil && !dcl.IsEmptyValueIndirect(desired.GuestOsFeatures) {
		c.Config.Logger.Infof("desired GuestOsFeatures %s - but actually nil", dcl.SprintResource(desired.GuestOsFeatures))
		return true
	}
	if compareInstanceTemplatePropertiesDisksGuestOsFeaturesSlice(c, desired.GuestOsFeatures, actual.GuestOsFeatures) && !dcl.IsZeroValue(desired.GuestOsFeatures) {
		c.Config.Logger.Infof("Diff in GuestOsFeatures. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.GuestOsFeatures), dcl.SprintResource(actual.GuestOsFeatures))
		return true
	}
	if actual.Interface == nil && desired.Interface != nil && !dcl.IsEmptyValueIndirect(desired.Interface) {
		c.Config.Logger.Infof("desired Interface %s - but actually nil", dcl.SprintResource(desired.Interface))
		return true
	}
	if !reflect.DeepEqual(desired.Interface, actual.Interface) && !dcl.IsZeroValue(desired.Interface) && !(dcl.IsEmptyValueIndirect(desired.Interface) && dcl.IsZeroValue(actual.Interface)) {
		c.Config.Logger.Infof("Diff in Interface. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Interface), dcl.SprintResource(actual.Interface))
		return true
	}
	if actual.Mode == nil && desired.Mode != nil && !dcl.IsEmptyValueIndirect(desired.Mode) {
		c.Config.Logger.Infof("desired Mode %s - but actually nil", dcl.SprintResource(desired.Mode))
		return true
	}
	if !reflect.DeepEqual(desired.Mode, actual.Mode) && !dcl.IsZeroValue(desired.Mode) && !(dcl.IsEmptyValueIndirect(desired.Mode) && dcl.IsZeroValue(actual.Mode)) {
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
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) && !(dcl.IsEmptyValueIndirect(desired.Type) && dcl.IsZeroValue(actual.Type)) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesDisksDiskEncryptionKeySlice(c *Client, desired, actual []InstanceTemplatePropertiesDisksDiskEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesDisksDiskEncryptionKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesDisksDiskEncryptionKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesDisksDiskEncryptionKey, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesDisksDiskEncryptionKey(c *Client, desired, actual *InstanceTemplatePropertiesDisksDiskEncryptionKey) bool {
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
	if !reflect.DeepEqual(desired.RawKey, actual.RawKey) && !dcl.IsZeroValue(desired.RawKey) && !(dcl.IsEmptyValueIndirect(desired.RawKey) && dcl.IsZeroValue(actual.RawKey)) {
		c.Config.Logger.Infof("Diff in RawKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RawKey), dcl.SprintResource(actual.RawKey))
		return true
	}
	if actual.RsaEncryptedKey == nil && desired.RsaEncryptedKey != nil && !dcl.IsEmptyValueIndirect(desired.RsaEncryptedKey) {
		c.Config.Logger.Infof("desired RsaEncryptedKey %s - but actually nil", dcl.SprintResource(desired.RsaEncryptedKey))
		return true
	}
	if !reflect.DeepEqual(desired.RsaEncryptedKey, actual.RsaEncryptedKey) && !dcl.IsZeroValue(desired.RsaEncryptedKey) && !(dcl.IsEmptyValueIndirect(desired.RsaEncryptedKey) && dcl.IsZeroValue(actual.RsaEncryptedKey)) {
		c.Config.Logger.Infof("Diff in RsaEncryptedKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RsaEncryptedKey), dcl.SprintResource(actual.RsaEncryptedKey))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesDisksInitializeParamsSlice(c *Client, desired, actual []InstanceTemplatePropertiesDisksInitializeParams) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesDisksInitializeParams, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesDisksInitializeParams(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesDisksInitializeParams, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesDisksInitializeParams(c *Client, desired, actual *InstanceTemplatePropertiesDisksInitializeParams) bool {
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
	if !reflect.DeepEqual(desired.DiskName, actual.DiskName) && !dcl.IsZeroValue(desired.DiskName) && !(dcl.IsEmptyValueIndirect(desired.DiskName) && dcl.IsZeroValue(actual.DiskName)) {
		c.Config.Logger.Infof("Diff in DiskName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskName), dcl.SprintResource(actual.DiskName))
		return true
	}
	if actual.DiskSizeGb == nil && desired.DiskSizeGb != nil && !dcl.IsEmptyValueIndirect(desired.DiskSizeGb) {
		c.Config.Logger.Infof("desired DiskSizeGb %s - but actually nil", dcl.SprintResource(desired.DiskSizeGb))
		return true
	}
	if !reflect.DeepEqual(desired.DiskSizeGb, actual.DiskSizeGb) && !dcl.IsZeroValue(desired.DiskSizeGb) && !(dcl.IsEmptyValueIndirect(desired.DiskSizeGb) && dcl.IsZeroValue(actual.DiskSizeGb)) {
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
	if !reflect.DeepEqual(desired.SourceImage, actual.SourceImage) && !dcl.IsZeroValue(desired.SourceImage) && !(dcl.IsEmptyValueIndirect(desired.SourceImage) && dcl.IsZeroValue(actual.SourceImage)) {
		c.Config.Logger.Infof("Diff in SourceImage. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceImage), dcl.SprintResource(actual.SourceImage))
		return true
	}
	if actual.Labels == nil && desired.Labels != nil && !dcl.IsEmptyValueIndirect(desired.Labels) {
		c.Config.Logger.Infof("desired Labels %s - but actually nil", dcl.SprintResource(desired.Labels))
		return true
	}
	if !reflect.DeepEqual(desired.Labels, actual.Labels) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	if actual.SourceSnapshot == nil && desired.SourceSnapshot != nil && !dcl.IsEmptyValueIndirect(desired.SourceSnapshot) {
		c.Config.Logger.Infof("desired SourceSnapshot %s - but actually nil", dcl.SprintResource(desired.SourceSnapshot))
		return true
	}
	if !reflect.DeepEqual(desired.SourceSnapshot, actual.SourceSnapshot) && !dcl.IsZeroValue(desired.SourceSnapshot) && !(dcl.IsEmptyValueIndirect(desired.SourceSnapshot) && dcl.IsZeroValue(actual.SourceSnapshot)) {
		c.Config.Logger.Infof("Diff in SourceSnapshot. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceSnapshot), dcl.SprintResource(actual.SourceSnapshot))
		return true
	}
	if actual.SourceSnapshotEncryptionKey == nil && desired.SourceSnapshotEncryptionKey != nil && !dcl.IsEmptyValueIndirect(desired.SourceSnapshotEncryptionKey) {
		c.Config.Logger.Infof("desired SourceSnapshotEncryptionKey %s - but actually nil", dcl.SprintResource(desired.SourceSnapshotEncryptionKey))
		return true
	}
	if compareInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, desired.SourceSnapshotEncryptionKey, actual.SourceSnapshotEncryptionKey) && !dcl.IsZeroValue(desired.SourceSnapshotEncryptionKey) {
		c.Config.Logger.Infof("Diff in SourceSnapshotEncryptionKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceSnapshotEncryptionKey), dcl.SprintResource(actual.SourceSnapshotEncryptionKey))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !reflect.DeepEqual(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) && !(dcl.IsEmptyValueIndirect(desired.Description) && dcl.IsZeroValue(actual.Description)) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.ResourcePolicies == nil && desired.ResourcePolicies != nil && !dcl.IsEmptyValueIndirect(desired.ResourcePolicies) {
		c.Config.Logger.Infof("desired ResourcePolicies %s - but actually nil", dcl.SprintResource(desired.ResourcePolicies))
		return true
	}
	if !dcl.SliceEquals(desired.ResourcePolicies, actual.ResourcePolicies) && !dcl.IsZeroValue(desired.ResourcePolicies) {
		c.Config.Logger.Infof("Diff in ResourcePolicies. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourcePolicies), dcl.SprintResource(actual.ResourcePolicies))
		return true
	}
	if actual.OnUpdateAction == nil && desired.OnUpdateAction != nil && !dcl.IsEmptyValueIndirect(desired.OnUpdateAction) {
		c.Config.Logger.Infof("desired OnUpdateAction %s - but actually nil", dcl.SprintResource(desired.OnUpdateAction))
		return true
	}
	if !reflect.DeepEqual(desired.OnUpdateAction, actual.OnUpdateAction) && !dcl.IsZeroValue(desired.OnUpdateAction) && !(dcl.IsEmptyValueIndirect(desired.OnUpdateAction) && dcl.IsZeroValue(actual.OnUpdateAction)) {
		c.Config.Logger.Infof("Diff in OnUpdateAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OnUpdateAction), dcl.SprintResource(actual.OnUpdateAction))
		return true
	}
	if actual.SourceImageEncryptionKey == nil && desired.SourceImageEncryptionKey != nil && !dcl.IsEmptyValueIndirect(desired.SourceImageEncryptionKey) {
		c.Config.Logger.Infof("desired SourceImageEncryptionKey %s - but actually nil", dcl.SprintResource(desired.SourceImageEncryptionKey))
		return true
	}
	if compareInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, desired.SourceImageEncryptionKey, actual.SourceImageEncryptionKey) && !dcl.IsZeroValue(desired.SourceImageEncryptionKey) {
		c.Config.Logger.Infof("Diff in SourceImageEncryptionKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceImageEncryptionKey), dcl.SprintResource(actual.SourceImageEncryptionKey))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeySlice(c *Client, desired, actual []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c *Client, desired, actual *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) bool {
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
	if !reflect.DeepEqual(desired.RawKey, actual.RawKey) && !dcl.IsZeroValue(desired.RawKey) && !(dcl.IsEmptyValueIndirect(desired.RawKey) && dcl.IsZeroValue(actual.RawKey)) {
		c.Config.Logger.Infof("Diff in RawKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RawKey), dcl.SprintResource(actual.RawKey))
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
	return false
}
func compareInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeySlice(c *Client, desired, actual []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c *Client, desired, actual *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) bool {
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
	if !reflect.DeepEqual(desired.RawKey, actual.RawKey) && !dcl.IsZeroValue(desired.RawKey) && !(dcl.IsEmptyValueIndirect(desired.RawKey) && dcl.IsZeroValue(actual.RawKey)) {
		c.Config.Logger.Infof("Diff in RawKey. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RawKey), dcl.SprintResource(actual.RawKey))
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
	return false
}
func compareInstanceTemplatePropertiesDisksGuestOsFeaturesSlice(c *Client, desired, actual []InstanceTemplatePropertiesDisksGuestOsFeatures) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesDisksGuestOsFeatures, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesDisksGuestOsFeatures(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesDisksGuestOsFeatures, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesDisksGuestOsFeatures(c *Client, desired, actual *InstanceTemplatePropertiesDisksGuestOsFeatures) bool {
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
	return false
}
func compareInstanceTemplatePropertiesReservationAffinitySlice(c *Client, desired, actual []InstanceTemplatePropertiesReservationAffinity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesReservationAffinity, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesReservationAffinity(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesReservationAffinity, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesReservationAffinity(c *Client, desired, actual *InstanceTemplatePropertiesReservationAffinity) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !reflect.DeepEqual(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) && !(dcl.IsEmptyValueIndirect(desired.Key) && dcl.IsZeroValue(actual.Key)) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !dcl.SliceEquals(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesGuestAcceleratorsSlice(c *Client, desired, actual []InstanceTemplatePropertiesGuestAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesGuestAccelerators, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesGuestAccelerators(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesGuestAccelerators, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesGuestAccelerators(c *Client, desired, actual *InstanceTemplatePropertiesGuestAccelerators) bool {
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
	if !reflect.DeepEqual(desired.AcceleratorCount, actual.AcceleratorCount) && !dcl.IsZeroValue(desired.AcceleratorCount) && !(dcl.IsEmptyValueIndirect(desired.AcceleratorCount) && dcl.IsZeroValue(actual.AcceleratorCount)) {
		c.Config.Logger.Infof("Diff in AcceleratorCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorCount), dcl.SprintResource(actual.AcceleratorCount))
		return true
	}
	if actual.AcceleratorType == nil && desired.AcceleratorType != nil && !dcl.IsEmptyValueIndirect(desired.AcceleratorType) {
		c.Config.Logger.Infof("desired AcceleratorType %s - but actually nil", dcl.SprintResource(desired.AcceleratorType))
		return true
	}
	if !reflect.DeepEqual(desired.AcceleratorType, actual.AcceleratorType) && !dcl.IsZeroValue(desired.AcceleratorType) && !(dcl.IsEmptyValueIndirect(desired.AcceleratorType) && dcl.IsZeroValue(actual.AcceleratorType)) {
		c.Config.Logger.Infof("Diff in AcceleratorType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorType), dcl.SprintResource(actual.AcceleratorType))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesNetworkInterfacesSlice(c *Client, desired, actual []InstanceTemplatePropertiesNetworkInterfaces) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesNetworkInterfaces, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesNetworkInterfaces(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesNetworkInterfaces, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesNetworkInterfaces(c *Client, desired, actual *InstanceTemplatePropertiesNetworkInterfaces) bool {
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
	if compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice(c, desired.AccessConfigs, actual.AccessConfigs) && !dcl.IsZeroValue(desired.AccessConfigs) {
		c.Config.Logger.Infof("Diff in AccessConfigs. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AccessConfigs), dcl.SprintResource(actual.AccessConfigs))
		return true
	}
	if actual.AliasIPRanges == nil && desired.AliasIPRanges != nil && !dcl.IsEmptyValueIndirect(desired.AliasIPRanges) {
		c.Config.Logger.Infof("desired AliasIPRanges %s - but actually nil", dcl.SprintResource(desired.AliasIPRanges))
		return true
	}
	if compareInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice(c, desired.AliasIPRanges, actual.AliasIPRanges) && !dcl.IsZeroValue(desired.AliasIPRanges) {
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
	if !reflect.DeepEqual(desired.NetworkIP, actual.NetworkIP) && !dcl.IsZeroValue(desired.NetworkIP) && !(dcl.IsEmptyValueIndirect(desired.NetworkIP) && dcl.IsZeroValue(actual.NetworkIP)) {
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
func compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice(c *Client, desired, actual []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesNetworkInterfacesAccessConfigs, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesNetworkInterfacesAccessConfigs, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c *Client, desired, actual *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) bool {
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
	if actual.NatIP == nil && desired.NatIP != nil && !dcl.IsEmptyValueIndirect(desired.NatIP) {
		c.Config.Logger.Infof("desired NatIP %s - but actually nil", dcl.SprintResource(desired.NatIP))
		return true
	}
	if !dcl.NameToSelfLink(desired.NatIP, actual.NatIP) && !dcl.IsZeroValue(desired.NatIP) {
		c.Config.Logger.Infof("Diff in NatIP. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NatIP), dcl.SprintResource(actual.NatIP))
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
	if actual.SetPublicPtr == nil && desired.SetPublicPtr != nil && !dcl.IsEmptyValueIndirect(desired.SetPublicPtr) {
		c.Config.Logger.Infof("desired SetPublicPtr %s - but actually nil", dcl.SprintResource(desired.SetPublicPtr))
		return true
	}
	if !reflect.DeepEqual(desired.SetPublicPtr, actual.SetPublicPtr) && !dcl.IsZeroValue(desired.SetPublicPtr) && !(dcl.IsEmptyValueIndirect(desired.SetPublicPtr) && dcl.IsZeroValue(actual.SetPublicPtr)) {
		c.Config.Logger.Infof("Diff in SetPublicPtr. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SetPublicPtr), dcl.SprintResource(actual.SetPublicPtr))
		return true
	}
	if actual.PublicPtrDomainName == nil && desired.PublicPtrDomainName != nil && !dcl.IsEmptyValueIndirect(desired.PublicPtrDomainName) {
		c.Config.Logger.Infof("desired PublicPtrDomainName %s - but actually nil", dcl.SprintResource(desired.PublicPtrDomainName))
		return true
	}
	if !reflect.DeepEqual(desired.PublicPtrDomainName, actual.PublicPtrDomainName) && !dcl.IsZeroValue(desired.PublicPtrDomainName) && !(dcl.IsEmptyValueIndirect(desired.PublicPtrDomainName) && dcl.IsZeroValue(actual.PublicPtrDomainName)) {
		c.Config.Logger.Infof("Diff in PublicPtrDomainName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PublicPtrDomainName), dcl.SprintResource(actual.PublicPtrDomainName))
		return true
	}
	if actual.NetworkTier == nil && desired.NetworkTier != nil && !dcl.IsEmptyValueIndirect(desired.NetworkTier) {
		c.Config.Logger.Infof("desired NetworkTier %s - but actually nil", dcl.SprintResource(desired.NetworkTier))
		return true
	}
	if !reflect.DeepEqual(desired.NetworkTier, actual.NetworkTier) && !dcl.IsZeroValue(desired.NetworkTier) && !(dcl.IsEmptyValueIndirect(desired.NetworkTier) && dcl.IsZeroValue(actual.NetworkTier)) {
		c.Config.Logger.Infof("Diff in NetworkTier. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NetworkTier), dcl.SprintResource(actual.NetworkTier))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice(c *Client, desired, actual []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c *Client, desired, actual *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) bool {
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
	if !reflect.DeepEqual(desired.IPCidrRange, actual.IPCidrRange) && !dcl.IsZeroValue(desired.IPCidrRange) && !(dcl.IsEmptyValueIndirect(desired.IPCidrRange) && dcl.IsZeroValue(actual.IPCidrRange)) {
		c.Config.Logger.Infof("Diff in IPCidrRange. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPCidrRange), dcl.SprintResource(actual.IPCidrRange))
		return true
	}
	if actual.SubnetworkRangeName == nil && desired.SubnetworkRangeName != nil && !dcl.IsEmptyValueIndirect(desired.SubnetworkRangeName) {
		c.Config.Logger.Infof("desired SubnetworkRangeName %s - but actually nil", dcl.SprintResource(desired.SubnetworkRangeName))
		return true
	}
	if !reflect.DeepEqual(desired.SubnetworkRangeName, actual.SubnetworkRangeName) && !dcl.IsZeroValue(desired.SubnetworkRangeName) && !(dcl.IsEmptyValueIndirect(desired.SubnetworkRangeName) && dcl.IsZeroValue(actual.SubnetworkRangeName)) {
		c.Config.Logger.Infof("Diff in SubnetworkRangeName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SubnetworkRangeName), dcl.SprintResource(actual.SubnetworkRangeName))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesShieldedInstanceConfigSlice(c *Client, desired, actual []InstanceTemplatePropertiesShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesShieldedInstanceConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesShieldedInstanceConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesShieldedInstanceConfig(c *Client, desired, actual *InstanceTemplatePropertiesShieldedInstanceConfig) bool {
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
	if !reflect.DeepEqual(desired.EnableSecureBoot, actual.EnableSecureBoot) && !dcl.IsZeroValue(desired.EnableSecureBoot) && !(dcl.IsEmptyValueIndirect(desired.EnableSecureBoot) && dcl.IsZeroValue(actual.EnableSecureBoot)) {
		c.Config.Logger.Infof("Diff in EnableSecureBoot. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableSecureBoot), dcl.SprintResource(actual.EnableSecureBoot))
		return true
	}
	if actual.EnableVtpm == nil && desired.EnableVtpm != nil && !dcl.IsEmptyValueIndirect(desired.EnableVtpm) {
		c.Config.Logger.Infof("desired EnableVtpm %s - but actually nil", dcl.SprintResource(desired.EnableVtpm))
		return true
	}
	if !reflect.DeepEqual(desired.EnableVtpm, actual.EnableVtpm) && !dcl.IsZeroValue(desired.EnableVtpm) && !(dcl.IsEmptyValueIndirect(desired.EnableVtpm) && dcl.IsZeroValue(actual.EnableVtpm)) {
		c.Config.Logger.Infof("Diff in EnableVtpm. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableVtpm), dcl.SprintResource(actual.EnableVtpm))
		return true
	}
	if actual.EnableIntegrityMonitoring == nil && desired.EnableIntegrityMonitoring != nil && !dcl.IsEmptyValueIndirect(desired.EnableIntegrityMonitoring) {
		c.Config.Logger.Infof("desired EnableIntegrityMonitoring %s - but actually nil", dcl.SprintResource(desired.EnableIntegrityMonitoring))
		return true
	}
	if !reflect.DeepEqual(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring) && !dcl.IsZeroValue(desired.EnableIntegrityMonitoring) && !(dcl.IsEmptyValueIndirect(desired.EnableIntegrityMonitoring) && dcl.IsZeroValue(actual.EnableIntegrityMonitoring)) {
		c.Config.Logger.Infof("Diff in EnableIntegrityMonitoring. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableIntegrityMonitoring), dcl.SprintResource(actual.EnableIntegrityMonitoring))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesSchedulingSlice(c *Client, desired, actual []InstanceTemplatePropertiesScheduling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesScheduling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesScheduling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesScheduling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesScheduling(c *Client, desired, actual *InstanceTemplatePropertiesScheduling) bool {
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
	if !reflect.DeepEqual(desired.AutomaticRestart, actual.AutomaticRestart) && !dcl.IsZeroValue(desired.AutomaticRestart) && !(dcl.IsEmptyValueIndirect(desired.AutomaticRestart) && dcl.IsZeroValue(actual.AutomaticRestart)) {
		c.Config.Logger.Infof("Diff in AutomaticRestart. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutomaticRestart), dcl.SprintResource(actual.AutomaticRestart))
		return true
	}
	if actual.OnHostMaintenance == nil && desired.OnHostMaintenance != nil && !dcl.IsEmptyValueIndirect(desired.OnHostMaintenance) {
		c.Config.Logger.Infof("desired OnHostMaintenance %s - but actually nil", dcl.SprintResource(desired.OnHostMaintenance))
		return true
	}
	if !reflect.DeepEqual(desired.OnHostMaintenance, actual.OnHostMaintenance) && !dcl.IsZeroValue(desired.OnHostMaintenance) && !(dcl.IsEmptyValueIndirect(desired.OnHostMaintenance) && dcl.IsZeroValue(actual.OnHostMaintenance)) {
		c.Config.Logger.Infof("Diff in OnHostMaintenance. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OnHostMaintenance), dcl.SprintResource(actual.OnHostMaintenance))
		return true
	}
	if actual.Preemptible == nil && desired.Preemptible != nil && !dcl.IsEmptyValueIndirect(desired.Preemptible) {
		c.Config.Logger.Infof("desired Preemptible %s - but actually nil", dcl.SprintResource(desired.Preemptible))
		return true
	}
	if !reflect.DeepEqual(desired.Preemptible, actual.Preemptible) && !dcl.IsZeroValue(desired.Preemptible) && !(dcl.IsEmptyValueIndirect(desired.Preemptible) && dcl.IsZeroValue(actual.Preemptible)) {
		c.Config.Logger.Infof("Diff in Preemptible. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Preemptible), dcl.SprintResource(actual.Preemptible))
		return true
	}
	if actual.NodeAffinities == nil && desired.NodeAffinities != nil && !dcl.IsEmptyValueIndirect(desired.NodeAffinities) {
		c.Config.Logger.Infof("desired NodeAffinities %s - but actually nil", dcl.SprintResource(desired.NodeAffinities))
		return true
	}
	if compareInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice(c, desired.NodeAffinities, actual.NodeAffinities) && !dcl.IsZeroValue(desired.NodeAffinities) {
		c.Config.Logger.Infof("Diff in NodeAffinities. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NodeAffinities), dcl.SprintResource(actual.NodeAffinities))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice(c *Client, desired, actual []InstanceTemplatePropertiesSchedulingNodeAffinities) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesSchedulingNodeAffinities, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesSchedulingNodeAffinities(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesSchedulingNodeAffinities, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesSchedulingNodeAffinities(c *Client, desired, actual *InstanceTemplatePropertiesSchedulingNodeAffinities) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !reflect.DeepEqual(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) && !(dcl.IsEmptyValueIndirect(desired.Key) && dcl.IsZeroValue(actual.Key)) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.Operator == nil && desired.Operator != nil && !dcl.IsEmptyValueIndirect(desired.Operator) {
		c.Config.Logger.Infof("desired Operator %s - but actually nil", dcl.SprintResource(desired.Operator))
		return true
	}
	if !reflect.DeepEqual(desired.Operator, actual.Operator) && !dcl.IsZeroValue(desired.Operator) && !(dcl.IsEmptyValueIndirect(desired.Operator) && dcl.IsZeroValue(actual.Operator)) {
		c.Config.Logger.Infof("Diff in Operator. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Operator), dcl.SprintResource(actual.Operator))
		return true
	}
	if actual.Values == nil && desired.Values != nil && !dcl.IsEmptyValueIndirect(desired.Values) {
		c.Config.Logger.Infof("desired Values %s - but actually nil", dcl.SprintResource(desired.Values))
		return true
	}
	if !dcl.SliceEquals(desired.Values, actual.Values) && !dcl.IsZeroValue(desired.Values) {
		c.Config.Logger.Infof("Diff in Values. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Values), dcl.SprintResource(actual.Values))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesServiceAccountsSlice(c *Client, desired, actual []InstanceTemplatePropertiesServiceAccounts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesServiceAccounts, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesServiceAccounts(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesServiceAccounts, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesServiceAccounts(c *Client, desired, actual *InstanceTemplatePropertiesServiceAccounts) bool {
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
	if !reflect.DeepEqual(desired.Email, actual.Email) && !dcl.IsZeroValue(desired.Email) && !(dcl.IsEmptyValueIndirect(desired.Email) && dcl.IsZeroValue(actual.Email)) {
		c.Config.Logger.Infof("Diff in Email. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Email), dcl.SprintResource(actual.Email))
		return true
	}
	if actual.Scopes == nil && desired.Scopes != nil && !dcl.IsEmptyValueIndirect(desired.Scopes) {
		c.Config.Logger.Infof("desired Scopes %s - but actually nil", dcl.SprintResource(desired.Scopes))
		return true
	}
	if !dcl.SliceEquals(desired.Scopes, actual.Scopes) && !dcl.IsZeroValue(desired.Scopes) {
		c.Config.Logger.Infof("Diff in Scopes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Scopes), dcl.SprintResource(actual.Scopes))
		return true
	}
	return false
}
func compareInstanceTemplatePropertiesDisksInterfaceEnumSlice(c *Client, desired, actual []InstanceTemplatePropertiesDisksInterfaceEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesDisksInterfaceEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesDisksInterfaceEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesDisksInterfaceEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesDisksInterfaceEnum(c *Client, desired, actual *InstanceTemplatePropertiesDisksInterfaceEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceTemplatePropertiesDisksModeEnumSlice(c *Client, desired, actual []InstanceTemplatePropertiesDisksModeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesDisksModeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesDisksModeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesDisksModeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesDisksModeEnum(c *Client, desired, actual *InstanceTemplatePropertiesDisksModeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceTemplatePropertiesDisksTypeEnumSlice(c *Client, desired, actual []InstanceTemplatePropertiesDisksTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesDisksTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesDisksTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesDisksTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesDisksTypeEnum(c *Client, desired, actual *InstanceTemplatePropertiesDisksTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumSlice(c *Client, desired, actual []InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(c *Client, desired, actual *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumSlice(c *Client, desired, actual []InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(c *Client, desired, actual *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumSlice(c *Client, desired, actual []InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(c *Client, desired, actual *InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *InstanceTemplate) urlNormalized() *InstanceTemplate {
	normalized := deepcopy.Copy(*r).(InstanceTemplate)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *InstanceTemplate) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *InstanceTemplate) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *InstanceTemplate) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *InstanceTemplate) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/instanceTemplates/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the InstanceTemplate resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *InstanceTemplate) marshal(c *Client) ([]byte, error) {
	m, err := expandInstanceTemplate(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling InstanceTemplate: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"properties", "tags"},
		[]string{"properties", "tags", "items"},
	)

	return json.Marshal(m)
}

// unmarshalInstanceTemplate decodes JSON responses into the InstanceTemplate resource schema.
func unmarshalInstanceTemplate(b []byte, c *Client) (*InstanceTemplate, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	if v, err := dcl.MapFromListOfKeyValues(m, []string{"properties", "metadata", "items"}); err != nil {
		return nil, err
	} else {
		dcl.PutMapEntry(
			m,
			[]string{"properties", "metadata"},
			v,
		)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"properties", "tags", "items"},
		[]string{"properties", "tags"},
	)

	return flattenInstanceTemplate(c, m), nil
}

// expandInstanceTemplate expands InstanceTemplate into a JSON request object.
func expandInstanceTemplate(c *Client, f *InstanceTemplate) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandInstanceTemplateProperties(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenInstanceTemplate flattens InstanceTemplate from a JSON request object into the
// InstanceTemplate type.
func flattenInstanceTemplate(c *Client, i interface{}) *InstanceTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &InstanceTemplate{}
	r.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	r.Description = dcl.FlattenString(m["description"])
	r.Id = dcl.FlattenInteger(m["id"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Name = dcl.FlattenString(m["name"])
	r.Properties = flattenInstanceTemplateProperties(c, m["properties"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandInstanceTemplatePropertiesMap expands the contents of InstanceTemplateProperties into a JSON
// request object.
func expandInstanceTemplatePropertiesMap(c *Client, f map[string]InstanceTemplateProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplateProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesSlice expands the contents of InstanceTemplateProperties into a JSON
// request object.
func expandInstanceTemplatePropertiesSlice(c *Client, f []InstanceTemplateProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplateProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesMap flattens the contents of InstanceTemplateProperties from a JSON
// response object.
func flattenInstanceTemplatePropertiesMap(c *Client, i interface{}) map[string]InstanceTemplateProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplateProperties{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplateProperties{}
	}

	items := make(map[string]InstanceTemplateProperties)
	for k, item := range a {
		items[k] = *flattenInstanceTemplateProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesSlice flattens the contents of InstanceTemplateProperties from a JSON
// response object.
func flattenInstanceTemplatePropertiesSlice(c *Client, i interface{}) []InstanceTemplateProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplateProperties{}
	}

	if len(a) == 0 {
		return []InstanceTemplateProperties{}
	}

	items := make([]InstanceTemplateProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplateProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplateProperties expands an instance of InstanceTemplateProperties into a JSON
// request object.
func expandInstanceTemplateProperties(c *Client, f *InstanceTemplateProperties) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CanIPForward; !dcl.IsEmptyValueIndirect(v) {
		m["canIPForward"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandInstanceTemplatePropertiesDisksSlice(c, f.Disks); err != nil {
		return nil, fmt.Errorf("error expanding Disks into disks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["disks"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}
	if v, err := dcl.ListOfKeyValuesFromMapInItemsStruct(f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v, err := expandInstanceTemplatePropertiesReservationAffinity(c, f.ReservationAffinity); err != nil {
		return nil, fmt.Errorf("error expanding ReservationAffinity into reservationAffinity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reservationAffinity"] = v
	}
	if v, err := expandInstanceTemplatePropertiesGuestAcceleratorsSlice(c, f.GuestAccelerators); err != nil {
		return nil, fmt.Errorf("error expanding GuestAccelerators into guestAccelerators: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["guestAccelerators"] = v
	}
	if v, err := expandInstanceTemplatePropertiesNetworkInterfacesSlice(c, f.NetworkInterfaces); err != nil {
		return nil, fmt.Errorf("error expanding NetworkInterfaces into networkInterfaces: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["networkInterfaces"] = v
	}
	if v, err := expandInstanceTemplatePropertiesShieldedInstanceConfig(c, f.ShieldedInstanceConfig); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedInstanceConfig into shieldedInstanceConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["shieldedInstanceConfig"] = v
	}
	if v, err := expandInstanceTemplatePropertiesScheduling(c, f.Scheduling); err != nil {
		return nil, fmt.Errorf("error expanding Scheduling into scheduling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scheduling"] = v
	}
	if v, err := expandInstanceTemplatePropertiesServiceAccountsSlice(c, f.ServiceAccounts); err != nil {
		return nil, fmt.Errorf("error expanding ServiceAccounts into serviceAccounts: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccounts"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}

	return m, nil
}

// flattenInstanceTemplateProperties flattens an instance of InstanceTemplateProperties from a JSON
// response object.
func flattenInstanceTemplateProperties(c *Client, i interface{}) *InstanceTemplateProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplateProperties{}
	r.CanIPForward = dcl.FlattenBool(m["canIPForward"])
	r.Description = dcl.FlattenString(m["description"])
	r.Disks = flattenInstanceTemplatePropertiesDisksSlice(c, m["disks"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.MachineType = dcl.FlattenString(m["machineType"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])
	r.Metadata = dcl.FlattenKeyValuePairs(m["metadata"])
	r.ReservationAffinity = flattenInstanceTemplatePropertiesReservationAffinity(c, m["reservationAffinity"])
	r.GuestAccelerators = flattenInstanceTemplatePropertiesGuestAcceleratorsSlice(c, m["guestAccelerators"])
	r.NetworkInterfaces = flattenInstanceTemplatePropertiesNetworkInterfacesSlice(c, m["networkInterfaces"])
	r.ShieldedInstanceConfig = flattenInstanceTemplatePropertiesShieldedInstanceConfig(c, m["shieldedInstanceConfig"])
	r.Scheduling = flattenInstanceTemplatePropertiesScheduling(c, m["scheduling"])
	r.ServiceAccounts = flattenInstanceTemplatePropertiesServiceAccountsSlice(c, m["serviceAccounts"])
	r.Tags = dcl.FlattenStringSlice(m["tags"])

	return r
}

// expandInstanceTemplatePropertiesDisksMap expands the contents of InstanceTemplatePropertiesDisks into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksMap(c *Client, f map[string]InstanceTemplatePropertiesDisks) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesDisks(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesDisksSlice expands the contents of InstanceTemplatePropertiesDisks into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksSlice(c *Client, f []InstanceTemplatePropertiesDisks) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesDisks(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesDisksMap flattens the contents of InstanceTemplatePropertiesDisks from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesDisks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesDisks{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesDisks{}
	}

	items := make(map[string]InstanceTemplatePropertiesDisks)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesDisks(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesDisksSlice flattens the contents of InstanceTemplatePropertiesDisks from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksSlice(c *Client, i interface{}) []InstanceTemplatePropertiesDisks {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesDisks{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesDisks{}
	}

	items := make([]InstanceTemplatePropertiesDisks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesDisks(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesDisks expands an instance of InstanceTemplatePropertiesDisks into a JSON
// request object.
func expandInstanceTemplatePropertiesDisks(c *Client, f *InstanceTemplatePropertiesDisks) (map[string]interface{}, error) {
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
	if v, err := expandInstanceTemplatePropertiesDisksDiskEncryptionKey(c, f.DiskEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding DiskEncryptionKey into diskEncryptionKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["diskEncryptionKey"] = v
	}
	if v := f.Index; !dcl.IsEmptyValueIndirect(v) {
		m["index"] = v
	}
	if v, err := expandInstanceTemplatePropertiesDisksInitializeParams(c, f.InitializeParams); err != nil {
		return nil, fmt.Errorf("error expanding InitializeParams into initializeParams: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["initializeParams"] = v
	}
	if v, err := expandInstanceTemplatePropertiesDisksGuestOsFeaturesSlice(c, f.GuestOsFeatures); err != nil {
		return nil, fmt.Errorf("error expanding GuestOsFeatures into guestOsFeatures: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["guestOsFeatures"] = v
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

// flattenInstanceTemplatePropertiesDisks flattens an instance of InstanceTemplatePropertiesDisks from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisks(c *Client, i interface{}) *InstanceTemplatePropertiesDisks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesDisks{}
	r.AutoDelete = dcl.FlattenBool(m["autoDelete"])
	r.Boot = dcl.FlattenBool(m["boot"])
	r.DeviceName = dcl.FlattenString(m["deviceName"])
	r.DiskEncryptionKey = flattenInstanceTemplatePropertiesDisksDiskEncryptionKey(c, m["diskEncryptionKey"])
	r.Index = dcl.FlattenInteger(m["index"])
	r.InitializeParams = flattenInstanceTemplatePropertiesDisksInitializeParams(c, m["initializeParams"])
	r.GuestOsFeatures = flattenInstanceTemplatePropertiesDisksGuestOsFeaturesSlice(c, m["guestOsFeatures"])
	r.Interface = flattenInstanceTemplatePropertiesDisksInterfaceEnum(m["interface"])
	r.Mode = flattenInstanceTemplatePropertiesDisksModeEnum(m["mode"])
	r.Source = dcl.FlattenString(m["source"])
	r.Type = flattenInstanceTemplatePropertiesDisksTypeEnum(m["type"])

	return r
}

// expandInstanceTemplatePropertiesDisksDiskEncryptionKeyMap expands the contents of InstanceTemplatePropertiesDisksDiskEncryptionKey into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksDiskEncryptionKeyMap(c *Client, f map[string]InstanceTemplatePropertiesDisksDiskEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesDisksDiskEncryptionKeySlice expands the contents of InstanceTemplatePropertiesDisksDiskEncryptionKey into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksDiskEncryptionKeySlice(c *Client, f []InstanceTemplatePropertiesDisksDiskEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksDiskEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesDisksDiskEncryptionKeyMap flattens the contents of InstanceTemplatePropertiesDisksDiskEncryptionKey from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksDiskEncryptionKeyMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesDisksDiskEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesDisksDiskEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesDisksDiskEncryptionKey{}
	}

	items := make(map[string]InstanceTemplatePropertiesDisksDiskEncryptionKey)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesDisksDiskEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesDisksDiskEncryptionKeySlice flattens the contents of InstanceTemplatePropertiesDisksDiskEncryptionKey from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksDiskEncryptionKeySlice(c *Client, i interface{}) []InstanceTemplatePropertiesDisksDiskEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesDisksDiskEncryptionKey{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesDisksDiskEncryptionKey{}
	}

	items := make([]InstanceTemplatePropertiesDisksDiskEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesDisksDiskEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesDisksDiskEncryptionKey expands an instance of InstanceTemplatePropertiesDisksDiskEncryptionKey into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksDiskEncryptionKey(c *Client, f *InstanceTemplatePropertiesDisksDiskEncryptionKey) (map[string]interface{}, error) {
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

// flattenInstanceTemplatePropertiesDisksDiskEncryptionKey flattens an instance of InstanceTemplatePropertiesDisksDiskEncryptionKey from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksDiskEncryptionKey(c *Client, i interface{}) *InstanceTemplatePropertiesDisksDiskEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesDisksDiskEncryptionKey{}
	r.RawKey = dcl.FlattenString(m["rawKey"])
	r.RsaEncryptedKey = dcl.FlattenString(m["rsaEncryptedKey"])
	r.Sha256 = dcl.FlattenString(m["sha256"])

	return r
}

// expandInstanceTemplatePropertiesDisksInitializeParamsMap expands the contents of InstanceTemplatePropertiesDisksInitializeParams into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksInitializeParamsMap(c *Client, f map[string]InstanceTemplatePropertiesDisksInitializeParams) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksInitializeParams(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesDisksInitializeParamsSlice expands the contents of InstanceTemplatePropertiesDisksInitializeParams into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksInitializeParamsSlice(c *Client, f []InstanceTemplatePropertiesDisksInitializeParams) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksInitializeParams(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesDisksInitializeParamsMap flattens the contents of InstanceTemplatePropertiesDisksInitializeParams from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInitializeParamsMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesDisksInitializeParams {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesDisksInitializeParams{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesDisksInitializeParams{}
	}

	items := make(map[string]InstanceTemplatePropertiesDisksInitializeParams)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesDisksInitializeParams(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesDisksInitializeParamsSlice flattens the contents of InstanceTemplatePropertiesDisksInitializeParams from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInitializeParamsSlice(c *Client, i interface{}) []InstanceTemplatePropertiesDisksInitializeParams {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesDisksInitializeParams{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesDisksInitializeParams{}
	}

	items := make([]InstanceTemplatePropertiesDisksInitializeParams, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesDisksInitializeParams(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesDisksInitializeParams expands an instance of InstanceTemplatePropertiesDisksInitializeParams into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksInitializeParams(c *Client, f *InstanceTemplatePropertiesDisksInitializeParams) (map[string]interface{}, error) {
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
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.SourceSnapshot; !dcl.IsEmptyValueIndirect(v) {
		m["sourceSnapshot"] = v
	}
	if v, err := expandInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, f.SourceSnapshotEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding SourceSnapshotEncryptionKey into sourceSnapshotEncryptionKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sourceSnapshotEncryptionKey"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.ResourcePolicies; !dcl.IsEmptyValueIndirect(v) {
		m["resourcePolicies"] = v
	}
	if v := f.OnUpdateAction; !dcl.IsEmptyValueIndirect(v) {
		m["onUpdateAction"] = v
	}
	if v, err := expandInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, f.SourceImageEncryptionKey); err != nil {
		return nil, fmt.Errorf("error expanding SourceImageEncryptionKey into sourceImageEncryptionKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sourceImageEncryptionKey"] = v
	}

	return m, nil
}

// flattenInstanceTemplatePropertiesDisksInitializeParams flattens an instance of InstanceTemplatePropertiesDisksInitializeParams from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInitializeParams(c *Client, i interface{}) *InstanceTemplatePropertiesDisksInitializeParams {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesDisksInitializeParams{}
	r.DiskName = dcl.FlattenString(m["diskName"])
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.DiskType = dcl.FlattenString(m["diskType"])
	r.SourceImage = dcl.FlattenString(m["sourceImage"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.SourceSnapshot = dcl.FlattenString(m["sourceSnapshot"])
	r.SourceSnapshotEncryptionKey = flattenInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, m["sourceSnapshotEncryptionKey"])
	r.Description = dcl.FlattenString(m["description"])
	r.ResourcePolicies = dcl.FlattenStringSlice(m["resourcePolicies"])
	r.OnUpdateAction = dcl.FlattenString(m["onUpdateAction"])
	r.SourceImageEncryptionKey = flattenInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, m["sourceImageEncryptionKey"])

	return r
}

// expandInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyMap expands the contents of InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyMap(c *Client, f map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeySlice expands the contents of InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeySlice(c *Client, f []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyMap flattens the contents of InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{}
	}

	items := make(map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeySlice flattens the contents of InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeySlice(c *Client, i interface{}) []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{}
	}

	items := make([]InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey expands an instance of InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c *Client, f *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) (map[string]interface{}, error) {
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
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}

	return m, nil
}

// flattenInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey flattens an instance of InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c *Client, i interface{}) *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{}
	r.RawKey = dcl.FlattenString(m["rawKey"])
	r.Sha256 = dcl.FlattenString(m["sha256"])
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])

	return r
}

// expandInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyMap expands the contents of InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyMap(c *Client, f map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeySlice expands the contents of InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeySlice(c *Client, f []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyMap flattens the contents of InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{}
	}

	items := make(map[string]InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeySlice flattens the contents of InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeySlice(c *Client, i interface{}) []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{}
	}

	items := make([]InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey expands an instance of InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c *Client, f *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) (map[string]interface{}, error) {
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
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}

	return m, nil
}

// flattenInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey flattens an instance of InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c *Client, i interface{}) *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{}
	r.RawKey = dcl.FlattenString(m["rawKey"])
	r.Sha256 = dcl.FlattenString(m["sha256"])
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])

	return r
}

// expandInstanceTemplatePropertiesDisksGuestOsFeaturesMap expands the contents of InstanceTemplatePropertiesDisksGuestOsFeatures into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksGuestOsFeaturesMap(c *Client, f map[string]InstanceTemplatePropertiesDisksGuestOsFeatures) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksGuestOsFeatures(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesDisksGuestOsFeaturesSlice expands the contents of InstanceTemplatePropertiesDisksGuestOsFeatures into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksGuestOsFeaturesSlice(c *Client, f []InstanceTemplatePropertiesDisksGuestOsFeatures) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesDisksGuestOsFeatures(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesDisksGuestOsFeaturesMap flattens the contents of InstanceTemplatePropertiesDisksGuestOsFeatures from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksGuestOsFeaturesMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesDisksGuestOsFeatures {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesDisksGuestOsFeatures{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesDisksGuestOsFeatures{}
	}

	items := make(map[string]InstanceTemplatePropertiesDisksGuestOsFeatures)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesDisksGuestOsFeatures(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesDisksGuestOsFeaturesSlice flattens the contents of InstanceTemplatePropertiesDisksGuestOsFeatures from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksGuestOsFeaturesSlice(c *Client, i interface{}) []InstanceTemplatePropertiesDisksGuestOsFeatures {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesDisksGuestOsFeatures{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesDisksGuestOsFeatures{}
	}

	items := make([]InstanceTemplatePropertiesDisksGuestOsFeatures, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesDisksGuestOsFeatures(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesDisksGuestOsFeatures expands an instance of InstanceTemplatePropertiesDisksGuestOsFeatures into a JSON
// request object.
func expandInstanceTemplatePropertiesDisksGuestOsFeatures(c *Client, f *InstanceTemplatePropertiesDisksGuestOsFeatures) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenInstanceTemplatePropertiesDisksGuestOsFeatures flattens an instance of InstanceTemplatePropertiesDisksGuestOsFeatures from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksGuestOsFeatures(c *Client, i interface{}) *InstanceTemplatePropertiesDisksGuestOsFeatures {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesDisksGuestOsFeatures{}
	r.Type = dcl.FlattenString(m["type"])

	return r
}

// expandInstanceTemplatePropertiesReservationAffinityMap expands the contents of InstanceTemplatePropertiesReservationAffinity into a JSON
// request object.
func expandInstanceTemplatePropertiesReservationAffinityMap(c *Client, f map[string]InstanceTemplatePropertiesReservationAffinity) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesReservationAffinitySlice expands the contents of InstanceTemplatePropertiesReservationAffinity into a JSON
// request object.
func expandInstanceTemplatePropertiesReservationAffinitySlice(c *Client, f []InstanceTemplatePropertiesReservationAffinity) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesReservationAffinityMap flattens the contents of InstanceTemplatePropertiesReservationAffinity from a JSON
// response object.
func flattenInstanceTemplatePropertiesReservationAffinityMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesReservationAffinity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesReservationAffinity{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesReservationAffinity{}
	}

	items := make(map[string]InstanceTemplatePropertiesReservationAffinity)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesReservationAffinity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesReservationAffinitySlice flattens the contents of InstanceTemplatePropertiesReservationAffinity from a JSON
// response object.
func flattenInstanceTemplatePropertiesReservationAffinitySlice(c *Client, i interface{}) []InstanceTemplatePropertiesReservationAffinity {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesReservationAffinity{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesReservationAffinity{}
	}

	items := make([]InstanceTemplatePropertiesReservationAffinity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesReservationAffinity(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesReservationAffinity expands an instance of InstanceTemplatePropertiesReservationAffinity into a JSON
// request object.
func expandInstanceTemplatePropertiesReservationAffinity(c *Client, f *InstanceTemplatePropertiesReservationAffinity) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceTemplatePropertiesReservationAffinity flattens an instance of InstanceTemplatePropertiesReservationAffinity from a JSON
// response object.
func flattenInstanceTemplatePropertiesReservationAffinity(c *Client, i interface{}) *InstanceTemplatePropertiesReservationAffinity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesReservationAffinity{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenStringSlice(m["value"])

	return r
}

// expandInstanceTemplatePropertiesGuestAcceleratorsMap expands the contents of InstanceTemplatePropertiesGuestAccelerators into a JSON
// request object.
func expandInstanceTemplatePropertiesGuestAcceleratorsMap(c *Client, f map[string]InstanceTemplatePropertiesGuestAccelerators) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesGuestAccelerators(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesGuestAcceleratorsSlice expands the contents of InstanceTemplatePropertiesGuestAccelerators into a JSON
// request object.
func expandInstanceTemplatePropertiesGuestAcceleratorsSlice(c *Client, f []InstanceTemplatePropertiesGuestAccelerators) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesGuestAccelerators(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesGuestAcceleratorsMap flattens the contents of InstanceTemplatePropertiesGuestAccelerators from a JSON
// response object.
func flattenInstanceTemplatePropertiesGuestAcceleratorsMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesGuestAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesGuestAccelerators{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesGuestAccelerators{}
	}

	items := make(map[string]InstanceTemplatePropertiesGuestAccelerators)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesGuestAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesGuestAcceleratorsSlice flattens the contents of InstanceTemplatePropertiesGuestAccelerators from a JSON
// response object.
func flattenInstanceTemplatePropertiesGuestAcceleratorsSlice(c *Client, i interface{}) []InstanceTemplatePropertiesGuestAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesGuestAccelerators{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesGuestAccelerators{}
	}

	items := make([]InstanceTemplatePropertiesGuestAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesGuestAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesGuestAccelerators expands an instance of InstanceTemplatePropertiesGuestAccelerators into a JSON
// request object.
func expandInstanceTemplatePropertiesGuestAccelerators(c *Client, f *InstanceTemplatePropertiesGuestAccelerators) (map[string]interface{}, error) {
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

// flattenInstanceTemplatePropertiesGuestAccelerators flattens an instance of InstanceTemplatePropertiesGuestAccelerators from a JSON
// response object.
func flattenInstanceTemplatePropertiesGuestAccelerators(c *Client, i interface{}) *InstanceTemplatePropertiesGuestAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesGuestAccelerators{}
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])
	r.AcceleratorType = dcl.FlattenString(m["acceleratorType"])

	return r
}

// expandInstanceTemplatePropertiesNetworkInterfacesMap expands the contents of InstanceTemplatePropertiesNetworkInterfaces into a JSON
// request object.
func expandInstanceTemplatePropertiesNetworkInterfacesMap(c *Client, f map[string]InstanceTemplatePropertiesNetworkInterfaces) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesNetworkInterfaces(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesNetworkInterfacesSlice expands the contents of InstanceTemplatePropertiesNetworkInterfaces into a JSON
// request object.
func expandInstanceTemplatePropertiesNetworkInterfacesSlice(c *Client, f []InstanceTemplatePropertiesNetworkInterfaces) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesNetworkInterfaces(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesNetworkInterfacesMap flattens the contents of InstanceTemplatePropertiesNetworkInterfaces from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesNetworkInterfaces {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesNetworkInterfaces{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesNetworkInterfaces{}
	}

	items := make(map[string]InstanceTemplatePropertiesNetworkInterfaces)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesNetworkInterfaces(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesNetworkInterfacesSlice flattens the contents of InstanceTemplatePropertiesNetworkInterfaces from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesSlice(c *Client, i interface{}) []InstanceTemplatePropertiesNetworkInterfaces {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesNetworkInterfaces{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesNetworkInterfaces{}
	}

	items := make([]InstanceTemplatePropertiesNetworkInterfaces, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesNetworkInterfaces(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesNetworkInterfaces expands an instance of InstanceTemplatePropertiesNetworkInterfaces into a JSON
// request object.
func expandInstanceTemplatePropertiesNetworkInterfaces(c *Client, f *InstanceTemplatePropertiesNetworkInterfaces) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice(c, f.AccessConfigs); err != nil {
		return nil, fmt.Errorf("error expanding AccessConfigs into accessConfigs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["accessConfigs"] = v
	}
	if v, err := expandInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice(c, f.AliasIPRanges); err != nil {
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

// flattenInstanceTemplatePropertiesNetworkInterfaces flattens an instance of InstanceTemplatePropertiesNetworkInterfaces from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfaces(c *Client, i interface{}) *InstanceTemplatePropertiesNetworkInterfaces {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesNetworkInterfaces{}
	r.AccessConfigs = flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice(c, m["accessConfigs"])
	r.AliasIPRanges = flattenInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice(c, m["aliasIPRanges"])
	r.Name = dcl.FlattenString(m["name"])
	r.Network = dcl.FlattenString(m["network"])
	r.NetworkIP = dcl.FlattenString(m["networkIP"])
	r.Subnetwork = dcl.FlattenString(m["subnetwork"])

	return r
}

// expandInstanceTemplatePropertiesNetworkInterfacesAccessConfigsMap expands the contents of InstanceTemplatePropertiesNetworkInterfacesAccessConfigs into a JSON
// request object.
func expandInstanceTemplatePropertiesNetworkInterfacesAccessConfigsMap(c *Client, f map[string]InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice expands the contents of InstanceTemplatePropertiesNetworkInterfacesAccessConfigs into a JSON
// request object.
func expandInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice(c *Client, f []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsMap flattens the contents of InstanceTemplatePropertiesNetworkInterfacesAccessConfigs from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesNetworkInterfacesAccessConfigs{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesNetworkInterfacesAccessConfigs{}
	}

	items := make(map[string]InstanceTemplatePropertiesNetworkInterfacesAccessConfigs)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice flattens the contents of InstanceTemplatePropertiesNetworkInterfacesAccessConfigs from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice(c *Client, i interface{}) []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs{}
	}

	items := make([]InstanceTemplatePropertiesNetworkInterfacesAccessConfigs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesNetworkInterfacesAccessConfigs expands an instance of InstanceTemplatePropertiesNetworkInterfacesAccessConfigs into a JSON
// request object.
func expandInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c *Client, f *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) (map[string]interface{}, error) {
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
	if v := f.SetPublicPtr; !dcl.IsEmptyValueIndirect(v) {
		m["setPublicPtr"] = v
	}
	if v := f.PublicPtrDomainName; !dcl.IsEmptyValueIndirect(v) {
		m["publicPtrDomainName"] = v
	}
	if v := f.NetworkTier; !dcl.IsEmptyValueIndirect(v) {
		m["networkTier"] = v
	}

	return m, nil
}

// flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigs flattens an instance of InstanceTemplatePropertiesNetworkInterfacesAccessConfigs from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c *Client, i interface{}) *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesNetworkInterfacesAccessConfigs{}
	r.Name = dcl.FlattenString(m["name"])
	r.NatIP = dcl.FlattenString(m["natIP"])
	r.Type = flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(m["type"])
	r.SetPublicPtr = dcl.FlattenBool(m["setPublicPtr"])
	r.PublicPtrDomainName = dcl.FlattenString(m["publicPtrDomainName"])
	r.NetworkTier = flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(m["networkTier"])

	return r
}

// expandInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesMap expands the contents of InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges into a JSON
// request object.
func expandInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesMap(c *Client, f map[string]InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice expands the contents of InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges into a JSON
// request object.
func expandInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice(c *Client, f []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesMap flattens the contents of InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{}
	}

	items := make(map[string]InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice flattens the contents of InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice(c *Client, i interface{}) []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{}
	}

	items := make([]InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges expands an instance of InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges into a JSON
// request object.
func expandInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c *Client, f *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) (map[string]interface{}, error) {
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

// flattenInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges flattens an instance of InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c *Client, i interface{}) *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{}
	r.IPCidrRange = dcl.FlattenString(m["ipCidrRange"])
	r.SubnetworkRangeName = dcl.FlattenString(m["subnetworkRangeName"])

	return r
}

// expandInstanceTemplatePropertiesShieldedInstanceConfigMap expands the contents of InstanceTemplatePropertiesShieldedInstanceConfig into a JSON
// request object.
func expandInstanceTemplatePropertiesShieldedInstanceConfigMap(c *Client, f map[string]InstanceTemplatePropertiesShieldedInstanceConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesShieldedInstanceConfigSlice expands the contents of InstanceTemplatePropertiesShieldedInstanceConfig into a JSON
// request object.
func expandInstanceTemplatePropertiesShieldedInstanceConfigSlice(c *Client, f []InstanceTemplatePropertiesShieldedInstanceConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesShieldedInstanceConfigMap flattens the contents of InstanceTemplatePropertiesShieldedInstanceConfig from a JSON
// response object.
func flattenInstanceTemplatePropertiesShieldedInstanceConfigMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesShieldedInstanceConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesShieldedInstanceConfig{}
	}

	items := make(map[string]InstanceTemplatePropertiesShieldedInstanceConfig)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesShieldedInstanceConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesShieldedInstanceConfigSlice flattens the contents of InstanceTemplatePropertiesShieldedInstanceConfig from a JSON
// response object.
func flattenInstanceTemplatePropertiesShieldedInstanceConfigSlice(c *Client, i interface{}) []InstanceTemplatePropertiesShieldedInstanceConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesShieldedInstanceConfig{}
	}

	items := make([]InstanceTemplatePropertiesShieldedInstanceConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesShieldedInstanceConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesShieldedInstanceConfig expands an instance of InstanceTemplatePropertiesShieldedInstanceConfig into a JSON
// request object.
func expandInstanceTemplatePropertiesShieldedInstanceConfig(c *Client, f *InstanceTemplatePropertiesShieldedInstanceConfig) (map[string]interface{}, error) {
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

// flattenInstanceTemplatePropertiesShieldedInstanceConfig flattens an instance of InstanceTemplatePropertiesShieldedInstanceConfig from a JSON
// response object.
func flattenInstanceTemplatePropertiesShieldedInstanceConfig(c *Client, i interface{}) *InstanceTemplatePropertiesShieldedInstanceConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesShieldedInstanceConfig{}
	r.EnableSecureBoot = dcl.FlattenBool(m["enableSecureBoot"])
	r.EnableVtpm = dcl.FlattenBool(m["enableVtpm"])
	r.EnableIntegrityMonitoring = dcl.FlattenBool(m["enableIntegrityMonitoring"])

	return r
}

// expandInstanceTemplatePropertiesSchedulingMap expands the contents of InstanceTemplatePropertiesScheduling into a JSON
// request object.
func expandInstanceTemplatePropertiesSchedulingMap(c *Client, f map[string]InstanceTemplatePropertiesScheduling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesScheduling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesSchedulingSlice expands the contents of InstanceTemplatePropertiesScheduling into a JSON
// request object.
func expandInstanceTemplatePropertiesSchedulingSlice(c *Client, f []InstanceTemplatePropertiesScheduling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesScheduling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesSchedulingMap flattens the contents of InstanceTemplatePropertiesScheduling from a JSON
// response object.
func flattenInstanceTemplatePropertiesSchedulingMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesScheduling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesScheduling{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesScheduling{}
	}

	items := make(map[string]InstanceTemplatePropertiesScheduling)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesScheduling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesSchedulingSlice flattens the contents of InstanceTemplatePropertiesScheduling from a JSON
// response object.
func flattenInstanceTemplatePropertiesSchedulingSlice(c *Client, i interface{}) []InstanceTemplatePropertiesScheduling {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesScheduling{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesScheduling{}
	}

	items := make([]InstanceTemplatePropertiesScheduling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesScheduling(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesScheduling expands an instance of InstanceTemplatePropertiesScheduling into a JSON
// request object.
func expandInstanceTemplatePropertiesScheduling(c *Client, f *InstanceTemplatePropertiesScheduling) (map[string]interface{}, error) {
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
	if v, err := expandInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice(c, f.NodeAffinities); err != nil {
		return nil, fmt.Errorf("error expanding NodeAffinities into nodeAffinities: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["nodeAffinities"] = v
	}

	return m, nil
}

// flattenInstanceTemplatePropertiesScheduling flattens an instance of InstanceTemplatePropertiesScheduling from a JSON
// response object.
func flattenInstanceTemplatePropertiesScheduling(c *Client, i interface{}) *InstanceTemplatePropertiesScheduling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesScheduling{}
	r.AutomaticRestart = dcl.FlattenBool(m["automaticRestart"])
	r.OnHostMaintenance = dcl.FlattenString(m["onHostMaintenance"])
	r.Preemptible = dcl.FlattenBool(m["preemptible"])
	r.NodeAffinities = flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice(c, m["nodeAffinities"])

	return r
}

// expandInstanceTemplatePropertiesSchedulingNodeAffinitiesMap expands the contents of InstanceTemplatePropertiesSchedulingNodeAffinities into a JSON
// request object.
func expandInstanceTemplatePropertiesSchedulingNodeAffinitiesMap(c *Client, f map[string]InstanceTemplatePropertiesSchedulingNodeAffinities) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesSchedulingNodeAffinities(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice expands the contents of InstanceTemplatePropertiesSchedulingNodeAffinities into a JSON
// request object.
func expandInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice(c *Client, f []InstanceTemplatePropertiesSchedulingNodeAffinities) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesSchedulingNodeAffinities(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesMap flattens the contents of InstanceTemplatePropertiesSchedulingNodeAffinities from a JSON
// response object.
func flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesSchedulingNodeAffinities {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesSchedulingNodeAffinities{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesSchedulingNodeAffinities{}
	}

	items := make(map[string]InstanceTemplatePropertiesSchedulingNodeAffinities)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesSchedulingNodeAffinities(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice flattens the contents of InstanceTemplatePropertiesSchedulingNodeAffinities from a JSON
// response object.
func flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice(c *Client, i interface{}) []InstanceTemplatePropertiesSchedulingNodeAffinities {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesSchedulingNodeAffinities{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesSchedulingNodeAffinities{}
	}

	items := make([]InstanceTemplatePropertiesSchedulingNodeAffinities, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesSchedulingNodeAffinities(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesSchedulingNodeAffinities expands an instance of InstanceTemplatePropertiesSchedulingNodeAffinities into a JSON
// request object.
func expandInstanceTemplatePropertiesSchedulingNodeAffinities(c *Client, f *InstanceTemplatePropertiesSchedulingNodeAffinities) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Operator; !dcl.IsEmptyValueIndirect(v) {
		m["operator"] = v
	}
	if v := f.Values; !dcl.IsEmptyValueIndirect(v) {
		m["values"] = v
	}

	return m, nil
}

// flattenInstanceTemplatePropertiesSchedulingNodeAffinities flattens an instance of InstanceTemplatePropertiesSchedulingNodeAffinities from a JSON
// response object.
func flattenInstanceTemplatePropertiesSchedulingNodeAffinities(c *Client, i interface{}) *InstanceTemplatePropertiesSchedulingNodeAffinities {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesSchedulingNodeAffinities{}
	r.Key = dcl.FlattenString(m["key"])
	r.Operator = flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(m["operator"])
	r.Values = dcl.FlattenStringSlice(m["values"])

	return r
}

// expandInstanceTemplatePropertiesServiceAccountsMap expands the contents of InstanceTemplatePropertiesServiceAccounts into a JSON
// request object.
func expandInstanceTemplatePropertiesServiceAccountsMap(c *Client, f map[string]InstanceTemplatePropertiesServiceAccounts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceTemplatePropertiesServiceAccounts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceTemplatePropertiesServiceAccountsSlice expands the contents of InstanceTemplatePropertiesServiceAccounts into a JSON
// request object.
func expandInstanceTemplatePropertiesServiceAccountsSlice(c *Client, f []InstanceTemplatePropertiesServiceAccounts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceTemplatePropertiesServiceAccounts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceTemplatePropertiesServiceAccountsMap flattens the contents of InstanceTemplatePropertiesServiceAccounts from a JSON
// response object.
func flattenInstanceTemplatePropertiesServiceAccountsMap(c *Client, i interface{}) map[string]InstanceTemplatePropertiesServiceAccounts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTemplatePropertiesServiceAccounts{}
	}

	if len(a) == 0 {
		return map[string]InstanceTemplatePropertiesServiceAccounts{}
	}

	items := make(map[string]InstanceTemplatePropertiesServiceAccounts)
	for k, item := range a {
		items[k] = *flattenInstanceTemplatePropertiesServiceAccounts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceTemplatePropertiesServiceAccountsSlice flattens the contents of InstanceTemplatePropertiesServiceAccounts from a JSON
// response object.
func flattenInstanceTemplatePropertiesServiceAccountsSlice(c *Client, i interface{}) []InstanceTemplatePropertiesServiceAccounts {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesServiceAccounts{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesServiceAccounts{}
	}

	items := make([]InstanceTemplatePropertiesServiceAccounts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesServiceAccounts(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceTemplatePropertiesServiceAccounts expands an instance of InstanceTemplatePropertiesServiceAccounts into a JSON
// request object.
func expandInstanceTemplatePropertiesServiceAccounts(c *Client, f *InstanceTemplatePropertiesServiceAccounts) (map[string]interface{}, error) {
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

// flattenInstanceTemplatePropertiesServiceAccounts flattens an instance of InstanceTemplatePropertiesServiceAccounts from a JSON
// response object.
func flattenInstanceTemplatePropertiesServiceAccounts(c *Client, i interface{}) *InstanceTemplatePropertiesServiceAccounts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceTemplatePropertiesServiceAccounts{}
	r.Email = dcl.FlattenString(m["email"])
	r.Scopes = dcl.FlattenStringSlice(m["scopes"])

	return r
}

// flattenInstanceTemplatePropertiesDisksInterfaceEnumSlice flattens the contents of InstanceTemplatePropertiesDisksInterfaceEnum from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksInterfaceEnumSlice(c *Client, i interface{}) []InstanceTemplatePropertiesDisksInterfaceEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesDisksInterfaceEnum{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesDisksInterfaceEnum{}
	}

	items := make([]InstanceTemplatePropertiesDisksInterfaceEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesDisksInterfaceEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceTemplatePropertiesDisksInterfaceEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTemplatePropertiesDisksInterfaceEnum with the same value as that string.
func flattenInstanceTemplatePropertiesDisksInterfaceEnum(i interface{}) *InstanceTemplatePropertiesDisksInterfaceEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTemplatePropertiesDisksInterfaceEnumRef("")
	}

	return InstanceTemplatePropertiesDisksInterfaceEnumRef(s)
}

// flattenInstanceTemplatePropertiesDisksModeEnumSlice flattens the contents of InstanceTemplatePropertiesDisksModeEnum from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksModeEnumSlice(c *Client, i interface{}) []InstanceTemplatePropertiesDisksModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesDisksModeEnum{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesDisksModeEnum{}
	}

	items := make([]InstanceTemplatePropertiesDisksModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesDisksModeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceTemplatePropertiesDisksModeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTemplatePropertiesDisksModeEnum with the same value as that string.
func flattenInstanceTemplatePropertiesDisksModeEnum(i interface{}) *InstanceTemplatePropertiesDisksModeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTemplatePropertiesDisksModeEnumRef("")
	}

	return InstanceTemplatePropertiesDisksModeEnumRef(s)
}

// flattenInstanceTemplatePropertiesDisksTypeEnumSlice flattens the contents of InstanceTemplatePropertiesDisksTypeEnum from a JSON
// response object.
func flattenInstanceTemplatePropertiesDisksTypeEnumSlice(c *Client, i interface{}) []InstanceTemplatePropertiesDisksTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesDisksTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesDisksTypeEnum{}
	}

	items := make([]InstanceTemplatePropertiesDisksTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesDisksTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceTemplatePropertiesDisksTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTemplatePropertiesDisksTypeEnum with the same value as that string.
func flattenInstanceTemplatePropertiesDisksTypeEnum(i interface{}) *InstanceTemplatePropertiesDisksTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTemplatePropertiesDisksTypeEnumRef("")
	}

	return InstanceTemplatePropertiesDisksTypeEnumRef(s)
}

// flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumSlice flattens the contents of InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumSlice(c *Client, i interface{}) []InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum{}
	}

	items := make([]InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum with the same value as that string.
func flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(i interface{}) *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumRef("")
	}

	return InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumRef(s)
}

// flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumSlice flattens the contents of InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum from a JSON
// response object.
func flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumSlice(c *Client, i interface{}) []InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum{}
	}

	items := make([]InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum with the same value as that string.
func flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(i interface{}) *InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumRef("")
	}

	return InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumRef(s)
}

// flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumSlice flattens the contents of InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum from a JSON
// response object.
func flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumSlice(c *Client, i interface{}) []InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum{}
	}

	if len(a) == 0 {
		return []InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum{}
	}

	items := make([]InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum with the same value as that string.
func flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(i interface{}) *InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumRef("")
	}

	return InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *InstanceTemplate) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInstanceTemplate(b, c)
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
