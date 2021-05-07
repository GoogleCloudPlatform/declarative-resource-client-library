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
	"strings"
	"time"

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
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
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

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetInstanceTemplate(ctx, r.urlNormalized())
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
type createInstanceTemplateOperation struct {
	response map[string]interface{}
}

func (op *createInstanceTemplateOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
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
	op.response, _ = o.FirstResponse()

	if _, err := c.GetInstanceTemplate(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getInstanceTemplateRaw(ctx context.Context, r *InstanceTemplate) ([]byte, error) {

	u, err := instanceTemplateGetURL(c.Config.BasePath, r.urlNormalized())
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

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Properties = canonicalizeInstanceTemplateProperties(rawDesired.Properties, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
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
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
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

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.CanIPForward, initial.CanIPForward) || dcl.IsZeroValue(des.CanIPForward) {
		des.CanIPForward = initial.CanIPForward
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
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
	if dcl.StringCanonicalize(des.MinCpuPlatform, initial.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
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

	if dcl.BoolCanonicalize(des.CanIPForward, nw.CanIPForward) {
		nw.CanIPForward = des.CanIPForward
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	nw.Disks = canonicalizeNewInstanceTemplatePropertiesDisksSlice(c, des.Disks, nw.Disks)
	if dcl.IsZeroValue(nw.Labels) {
		nw.Labels = des.Labels
	}
	if dcl.NameToSelfLink(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	if dcl.StringCanonicalize(des.MinCpuPlatform, nw.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}
	if dcl.IsZeroValue(nw.Metadata) {
		nw.Metadata = des.Metadata
	}
	nw.ReservationAffinity = canonicalizeNewInstanceTemplatePropertiesReservationAffinity(c, des.ReservationAffinity, nw.ReservationAffinity)
	nw.GuestAccelerators = canonicalizeNewInstanceTemplatePropertiesGuestAcceleratorsSlice(c, des.GuestAccelerators, nw.GuestAccelerators)
	nw.NetworkInterfaces = canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesSlice(c, des.NetworkInterfaces, nw.NetworkInterfaces)
	nw.ShieldedInstanceConfig = canonicalizeNewInstanceTemplatePropertiesShieldedInstanceConfig(c, des.ShieldedInstanceConfig, nw.ShieldedInstanceConfig)
	nw.Scheduling = canonicalizeNewInstanceTemplatePropertiesScheduling(c, des.Scheduling, nw.Scheduling)
	nw.ServiceAccounts = canonicalizeNewInstanceTemplatePropertiesServiceAccountsSlice(c, des.ServiceAccounts, nw.ServiceAccounts)
	if dcl.IsZeroValue(nw.Tags) {
		nw.Tags = des.Tags
	}

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
			if diffs, _ := compareInstanceTemplatePropertiesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesSlice(c *Client, des, nw []InstanceTemplateProperties) []InstanceTemplateProperties {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplateProperties
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplateProperties(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesDisks(des, initial *InstanceTemplatePropertiesDisks, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisks {
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

	if dcl.BoolCanonicalize(des.AutoDelete, nw.AutoDelete) {
		nw.AutoDelete = des.AutoDelete
	}
	if dcl.BoolCanonicalize(des.Boot, nw.Boot) {
		nw.Boot = des.Boot
	}
	if dcl.StringCanonicalize(des.DeviceName, nw.DeviceName) {
		nw.DeviceName = des.DeviceName
	}
	nw.DiskEncryptionKey = canonicalizeNewInstanceTemplatePropertiesDisksDiskEncryptionKey(c, des.DiskEncryptionKey, nw.DiskEncryptionKey)
	if dcl.IsZeroValue(nw.Index) {
		nw.Index = des.Index
	}
	nw.InitializeParams = canonicalizeNewInstanceTemplatePropertiesDisksInitializeParams(c, des.InitializeParams, nw.InitializeParams)
	nw.GuestOsFeatures = canonicalizeNewInstanceTemplatePropertiesDisksGuestOsFeaturesSlice(c, des.GuestOsFeatures, nw.GuestOsFeatures)
	if dcl.IsZeroValue(nw.Interface) {
		nw.Interface = des.Interface
	}
	if dcl.IsZeroValue(nw.Mode) {
		nw.Mode = des.Mode
	}
	if dcl.NameToSelfLink(des.Source, nw.Source) {
		nw.Source = des.Source
	}
	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
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
			if diffs, _ := compareInstanceTemplatePropertiesDisksNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesDisksSlice(c *Client, des, nw []InstanceTemplatePropertiesDisks) []InstanceTemplatePropertiesDisks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesDisks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesDisks(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesDisksDiskEncryptionKey(des, initial *InstanceTemplatePropertiesDisksDiskEncryptionKey, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksDiskEncryptionKey {
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

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksDiskEncryptionKey(c *Client, des, nw *InstanceTemplatePropertiesDisksDiskEncryptionKey) *InstanceTemplatePropertiesDisksDiskEncryptionKey {
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

func canonicalizeNewInstanceTemplatePropertiesDisksDiskEncryptionKeySet(c *Client, des, nw []InstanceTemplatePropertiesDisksDiskEncryptionKey) []InstanceTemplatePropertiesDisksDiskEncryptionKey {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesDisksDiskEncryptionKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceTemplatePropertiesDisksDiskEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesDisksDiskEncryptionKeySlice(c *Client, des, nw []InstanceTemplatePropertiesDisksDiskEncryptionKey) []InstanceTemplatePropertiesDisksDiskEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesDisksDiskEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesDisksDiskEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesDisksInitializeParams(des, initial *InstanceTemplatePropertiesDisksInitializeParams, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksInitializeParams {
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
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.StringCanonicalize(des.SourceSnapshot, initial.SourceSnapshot) || dcl.IsZeroValue(des.SourceSnapshot) {
		des.SourceSnapshot = initial.SourceSnapshot
	}
	des.SourceSnapshotEncryptionKey = canonicalizeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(des.SourceSnapshotEncryptionKey, initial.SourceSnapshotEncryptionKey, opts...)
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.ResourcePolicies) {
		des.ResourcePolicies = initial.ResourcePolicies
	}
	if dcl.StringCanonicalize(des.OnUpdateAction, initial.OnUpdateAction) || dcl.IsZeroValue(des.OnUpdateAction) {
		des.OnUpdateAction = initial.OnUpdateAction
	}
	des.SourceImageEncryptionKey = canonicalizeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(des.SourceImageEncryptionKey, initial.SourceImageEncryptionKey, opts...)

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParams(c *Client, des, nw *InstanceTemplatePropertiesDisksInitializeParams) *InstanceTemplatePropertiesDisksInitializeParams {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.DiskName, nw.DiskName) {
		nw.DiskName = des.DiskName
	}
	if dcl.IsZeroValue(nw.DiskSizeGb) {
		nw.DiskSizeGb = des.DiskSizeGb
	}
	if dcl.NameToSelfLink(des.DiskType, nw.DiskType) {
		nw.DiskType = des.DiskType
	}
	if dcl.StringCanonicalize(des.SourceImage, nw.SourceImage) {
		nw.SourceImage = des.SourceImage
	}
	if dcl.IsZeroValue(nw.Labels) {
		nw.Labels = des.Labels
	}
	if dcl.StringCanonicalize(des.SourceSnapshot, nw.SourceSnapshot) {
		nw.SourceSnapshot = des.SourceSnapshot
	}
	nw.SourceSnapshotEncryptionKey = canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, des.SourceSnapshotEncryptionKey, nw.SourceSnapshotEncryptionKey)
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.IsZeroValue(nw.ResourcePolicies) {
		nw.ResourcePolicies = des.ResourcePolicies
	}
	if dcl.StringCanonicalize(des.OnUpdateAction, nw.OnUpdateAction) {
		nw.OnUpdateAction = des.OnUpdateAction
	}
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
			if diffs, _ := compareInstanceTemplatePropertiesDisksInitializeParamsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSlice(c *Client, des, nw []InstanceTemplatePropertiesDisksInitializeParams) []InstanceTemplatePropertiesDisksInitializeParams {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesDisksInitializeParams
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesDisksInitializeParams(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(des, initial *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
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
	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c *Client, des, nw *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RawKey, nw.RawKey) {
		nw.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.Sha256, nw.Sha256) {
		nw.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
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
			if diffs, _ := compareInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeySlice(c *Client, des, nw []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(des, initial *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
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
	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c *Client, des, nw *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RawKey, nw.RawKey) {
		nw.RawKey = des.RawKey
	}
	if dcl.StringCanonicalize(des.Sha256, nw.Sha256) {
		nw.Sha256 = des.Sha256
	}
	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
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
			if diffs, _ := compareInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeySlice(c *Client, des, nw []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesDisksGuestOsFeatures(des, initial *InstanceTemplatePropertiesDisksGuestOsFeatures, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesDisksGuestOsFeatures {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Type, initial.Type) || dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesDisksGuestOsFeatures(c *Client, des, nw *InstanceTemplatePropertiesDisksGuestOsFeatures) *InstanceTemplatePropertiesDisksGuestOsFeatures {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
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
			if diffs, _ := compareInstanceTemplatePropertiesDisksGuestOsFeaturesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesDisksGuestOsFeaturesSlice(c *Client, des, nw []InstanceTemplatePropertiesDisksGuestOsFeatures) []InstanceTemplatePropertiesDisksGuestOsFeatures {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesDisksGuestOsFeatures
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesDisksGuestOsFeatures(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesReservationAffinity(des, initial *InstanceTemplatePropertiesReservationAffinity, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesReservationAffinity {
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
	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesReservationAffinity(c *Client, des, nw *InstanceTemplatePropertiesReservationAffinity) *InstanceTemplatePropertiesReservationAffinity {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.IsZeroValue(nw.Value) {
		nw.Value = des.Value
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
			if diffs, _ := compareInstanceTemplatePropertiesReservationAffinityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesReservationAffinitySlice(c *Client, des, nw []InstanceTemplatePropertiesReservationAffinity) []InstanceTemplatePropertiesReservationAffinity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesReservationAffinity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesReservationAffinity(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesGuestAccelerators(des, initial *InstanceTemplatePropertiesGuestAccelerators, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesGuestAccelerators {
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

func canonicalizeNewInstanceTemplatePropertiesGuestAccelerators(c *Client, des, nw *InstanceTemplatePropertiesGuestAccelerators) *InstanceTemplatePropertiesGuestAccelerators {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AcceleratorCount) {
		nw.AcceleratorCount = des.AcceleratorCount
	}
	if dcl.StringCanonicalize(des.AcceleratorType, nw.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
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
			if diffs, _ := compareInstanceTemplatePropertiesGuestAcceleratorsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesGuestAcceleratorsSlice(c *Client, des, nw []InstanceTemplatePropertiesGuestAccelerators) []InstanceTemplatePropertiesGuestAccelerators {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesGuestAccelerators
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesGuestAccelerators(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesNetworkInterfaces(des, initial *InstanceTemplatePropertiesNetworkInterfaces, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesNetworkInterfaces {
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

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfaces(c *Client, des, nw *InstanceTemplatePropertiesNetworkInterfaces) *InstanceTemplatePropertiesNetworkInterfaces {
	if des == nil || nw == nil {
		return nw
	}

	nw.AccessConfigs = canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice(c, des.AccessConfigs, nw.AccessConfigs)
	nw.AliasIPRanges = canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice(c, des.AliasIPRanges, nw.AliasIPRanges)
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

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesSet(c *Client, des, nw []InstanceTemplatePropertiesNetworkInterfaces) []InstanceTemplatePropertiesNetworkInterfaces {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesNetworkInterfaces
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceTemplatePropertiesNetworkInterfacesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesSlice(c *Client, des, nw []InstanceTemplatePropertiesNetworkInterfaces) []InstanceTemplatePropertiesNetworkInterfaces {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesNetworkInterfaces
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesNetworkInterfaces(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(des, initial *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
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
	if dcl.NameToSelfLink(des.NatIP, initial.NatIP) || dcl.IsZeroValue(des.NatIP) {
		des.NatIP = initial.NatIP
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.BoolCanonicalize(des.SetPublicPtr, initial.SetPublicPtr) || dcl.IsZeroValue(des.SetPublicPtr) {
		des.SetPublicPtr = initial.SetPublicPtr
	}
	if dcl.StringCanonicalize(des.PublicPtrDomainName, initial.PublicPtrDomainName) || dcl.IsZeroValue(des.PublicPtrDomainName) {
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

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.NameToSelfLink(des.NatIP, nw.NatIP) {
		nw.NatIP = des.NatIP
	}
	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
	}
	if dcl.BoolCanonicalize(des.SetPublicPtr, nw.SetPublicPtr) {
		nw.SetPublicPtr = des.SetPublicPtr
	}
	if dcl.StringCanonicalize(des.PublicPtrDomainName, nw.PublicPtrDomainName) {
		nw.PublicPtrDomainName = des.PublicPtrDomainName
	}
	if dcl.IsZeroValue(nw.NetworkTier) {
		nw.NetworkTier = des.NetworkTier
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
			if diffs, _ := compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAccessConfigsSlice(c *Client, des, nw []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesNetworkInterfacesAccessConfigs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(des, initial *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
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

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c *Client, des, nw *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
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

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSet(c *Client, des, nw []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesSlice(c *Client, des, nw []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesShieldedInstanceConfig(des, initial *InstanceTemplatePropertiesShieldedInstanceConfig, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesShieldedInstanceConfig {
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

func canonicalizeNewInstanceTemplatePropertiesShieldedInstanceConfig(c *Client, des, nw *InstanceTemplatePropertiesShieldedInstanceConfig) *InstanceTemplatePropertiesShieldedInstanceConfig {
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

func canonicalizeNewInstanceTemplatePropertiesShieldedInstanceConfigSet(c *Client, des, nw []InstanceTemplatePropertiesShieldedInstanceConfig) []InstanceTemplatePropertiesShieldedInstanceConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceTemplatePropertiesShieldedInstanceConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceTemplatePropertiesShieldedInstanceConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesShieldedInstanceConfigSlice(c *Client, des, nw []InstanceTemplatePropertiesShieldedInstanceConfig) []InstanceTemplatePropertiesShieldedInstanceConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesShieldedInstanceConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesShieldedInstanceConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesScheduling(des, initial *InstanceTemplatePropertiesScheduling, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesScheduling {
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
	if dcl.IsZeroValue(des.NodeAffinities) {
		des.NodeAffinities = initial.NodeAffinities
	}

	return des
}

func canonicalizeNewInstanceTemplatePropertiesScheduling(c *Client, des, nw *InstanceTemplatePropertiesScheduling) *InstanceTemplatePropertiesScheduling {
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
	nw.NodeAffinities = canonicalizeNewInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice(c, des.NodeAffinities, nw.NodeAffinities)

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
			if diffs, _ := compareInstanceTemplatePropertiesSchedulingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesSchedulingSlice(c *Client, des, nw []InstanceTemplatePropertiesScheduling) []InstanceTemplatePropertiesScheduling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesScheduling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesScheduling(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesSchedulingNodeAffinities(des, initial *InstanceTemplatePropertiesSchedulingNodeAffinities, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesSchedulingNodeAffinities {
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

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.IsZeroValue(nw.Operator) {
		nw.Operator = des.Operator
	}
	if dcl.IsZeroValue(nw.Values) {
		nw.Values = des.Values
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
			if diffs, _ := compareInstanceTemplatePropertiesSchedulingNodeAffinitiesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesSchedulingNodeAffinitiesSlice(c *Client, des, nw []InstanceTemplatePropertiesSchedulingNodeAffinities) []InstanceTemplatePropertiesSchedulingNodeAffinities {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesSchedulingNodeAffinities
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesSchedulingNodeAffinities(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceTemplatePropertiesServiceAccounts(des, initial *InstanceTemplatePropertiesServiceAccounts, opts ...dcl.ApplyOption) *InstanceTemplatePropertiesServiceAccounts {
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

func canonicalizeNewInstanceTemplatePropertiesServiceAccounts(c *Client, des, nw *InstanceTemplatePropertiesServiceAccounts) *InstanceTemplatePropertiesServiceAccounts {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Email, nw.Email) {
		nw.Email = des.Email
	}
	if dcl.IsZeroValue(nw.Scopes) {
		nw.Scopes = des.Scopes
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
			if diffs, _ := compareInstanceTemplatePropertiesServiceAccountsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceTemplatePropertiesServiceAccountsSlice(c *Client, des, nw []InstanceTemplatePropertiesServiceAccounts) []InstanceTemplatePropertiesServiceAccounts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceTemplatePropertiesServiceAccounts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceTemplatePropertiesServiceAccounts(c, &d, &n))
	}

	return items
}

type instanceTemplateDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceTemplateApiOperation
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
func diffInstanceTemplate(c *Client, desired, actual *InstanceTemplate, opts ...dcl.ApplyOption) ([]instanceTemplateDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []instanceTemplateDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.CreationTimestamp, actual.CreationTimestamp, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceTemplateDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceTemplateDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceTemplateDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceTemplateDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceTemplateDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Properties, actual.Properties, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Properties")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceTemplateDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceTemplateDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
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
func compareInstanceTemplatePropertiesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplateProperties)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplateProperties)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplateProperties or *InstanceTemplateProperties", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplateProperties)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplateProperties)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplateProperties", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CanIPForward, actual.CanIPForward, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CanIPForward")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disks, actual.Disks, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesDisksNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Disks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinCpuPlatform, actual.MinCpuPlatform, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinCpuPlatform")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReservationAffinity, actual.ReservationAffinity, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesReservationAffinityNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReservationAffinity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GuestAccelerators, actual.GuestAccelerators, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesGuestAcceleratorsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GuestAccelerators")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NetworkInterfaces, actual.NetworkInterfaces, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesNetworkInterfacesNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NetworkInterfaces")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesShieldedInstanceConfigNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ShieldedInstanceConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scheduling, actual.Scheduling, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesSchedulingNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Scheduling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccounts, actual.ServiceAccounts, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesServiceAccountsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tags, actual.Tags, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesDisksNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesDisks)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesDisks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisks or *InstanceTemplatePropertiesDisks", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesDisks)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesDisks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisks", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AutoDelete, actual.AutoDelete, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AutoDelete")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Boot, actual.Boot, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Boot")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeviceName, actual.DeviceName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeviceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskEncryptionKey, actual.DiskEncryptionKey, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesDisksDiskEncryptionKeyNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Index, actual.Index, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Index")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InitializeParams, actual.InitializeParams, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesDisksInitializeParamsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InitializeParams")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GuestOsFeatures, actual.GuestOsFeatures, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesDisksGuestOsFeaturesNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GuestOsFeatures")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Interface, actual.Interface, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Interface")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Source, actual.Source, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Source")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesDisksDiskEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesDisksDiskEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesDisksDiskEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksDiskEncryptionKey or *InstanceTemplatePropertiesDisksDiskEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesDisksDiskEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesDisksDiskEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksDiskEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RsaEncryptedKey, actual.RsaEncryptedKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RsaEncryptedKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256, actual.Sha256, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesDisksInitializeParamsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesDisksInitializeParams)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesDisksInitializeParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksInitializeParams or *InstanceTemplatePropertiesDisksInitializeParams", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesDisksInitializeParams)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesDisksInitializeParams)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksInitializeParams", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DiskName, actual.DiskName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskSizeGb, actual.DiskSizeGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskType, actual.DiskType, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DiskType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceImage, actual.SourceImage, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceImage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceSnapshot, actual.SourceSnapshot, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceSnapshot")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceSnapshotEncryptionKey, actual.SourceSnapshotEncryptionKey, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceSnapshotEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourcePolicies, actual.ResourcePolicies, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourcePolicies")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OnUpdateAction, actual.OnUpdateAction, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OnUpdateAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceImageEncryptionKey, actual.SourceImageEncryptionKey, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceImageEncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey or *InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256, actual.Sha256, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey or *InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawKey, actual.RawKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RawKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256, actual.Sha256, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesDisksGuestOsFeaturesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesDisksGuestOsFeatures)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesDisksGuestOsFeatures)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksGuestOsFeatures or *InstanceTemplatePropertiesDisksGuestOsFeatures", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesDisksGuestOsFeatures)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesDisksGuestOsFeatures)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesDisksGuestOsFeatures", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesReservationAffinityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesReservationAffinity)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesReservationAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesReservationAffinity or *InstanceTemplatePropertiesReservationAffinity", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesReservationAffinity)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesReservationAffinity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesReservationAffinity", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesGuestAcceleratorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesGuestAccelerators)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesGuestAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesGuestAccelerators or *InstanceTemplatePropertiesGuestAccelerators", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesGuestAccelerators)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesGuestAccelerators)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesGuestAccelerators", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AcceleratorCount, actual.AcceleratorCount, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorType, actual.AcceleratorType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesNetworkInterfacesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesNetworkInterfaces)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesNetworkInterfaces)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesNetworkInterfaces or *InstanceTemplatePropertiesNetworkInterfaces", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesNetworkInterfaces)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesNetworkInterfaces)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesNetworkInterfaces", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AccessConfigs, actual.AccessConfigs, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AccessConfigs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AliasIPRanges, actual.AliasIPRanges, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AliasIPRanges")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NetworkIP, actual.NetworkIP, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NetworkIP")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subnetwork, actual.Subnetwork, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subnetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesNetworkInterfacesAccessConfigs)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesNetworkInterfacesAccessConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesNetworkInterfacesAccessConfigs or *InstanceTemplatePropertiesNetworkInterfacesAccessConfigs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesNetworkInterfacesAccessConfigs)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesNetworkInterfacesAccessConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesNetworkInterfacesAccessConfigs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NatIP, actual.NatIP, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NatIP")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SetPublicPtr, actual.SetPublicPtr, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SetPublicPtr")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublicPtrDomainName, actual.PublicPtrDomainName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PublicPtrDomainName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NetworkTier, actual.NetworkTier, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NetworkTier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges or *InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IPCidrRange, actual.IPCidrRange, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IPCidrRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubnetworkRangeName, actual.SubnetworkRangeName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubnetworkRangeName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesShieldedInstanceConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesShieldedInstanceConfig)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesShieldedInstanceConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesShieldedInstanceConfig or *InstanceTemplatePropertiesShieldedInstanceConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesShieldedInstanceConfig)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesShieldedInstanceConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesShieldedInstanceConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnableSecureBoot, actual.EnableSecureBoot, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableSecureBoot")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableVtpm, actual.EnableVtpm, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableVtpm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableIntegrityMonitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesSchedulingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesScheduling)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesScheduling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesScheduling or *InstanceTemplatePropertiesScheduling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesScheduling)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesScheduling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesScheduling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AutomaticRestart, actual.AutomaticRestart, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AutomaticRestart")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OnHostMaintenance, actual.OnHostMaintenance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OnHostMaintenance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Preemptible, actual.Preemptible, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Preemptible")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeAffinities, actual.NodeAffinities, dcl.Info{ObjectFunction: compareInstanceTemplatePropertiesSchedulingNodeAffinitiesNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeAffinities")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesSchedulingNodeAffinitiesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesSchedulingNodeAffinities)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesSchedulingNodeAffinities)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesSchedulingNodeAffinities or *InstanceTemplatePropertiesSchedulingNodeAffinities", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesSchedulingNodeAffinities)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesSchedulingNodeAffinities)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesSchedulingNodeAffinities", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Operator, actual.Operator, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Operator")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Values, actual.Values, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Values")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceTemplatePropertiesServiceAccountsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceTemplatePropertiesServiceAccounts)
	if !ok {
		desiredNotPointer, ok := d.(InstanceTemplatePropertiesServiceAccounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesServiceAccounts or *InstanceTemplatePropertiesServiceAccounts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceTemplatePropertiesServiceAccounts)
	if !ok {
		actualNotPointer, ok := a.(InstanceTemplatePropertiesServiceAccounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceTemplatePropertiesServiceAccounts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Email, actual.Email, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Email")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scopes, actual.Scopes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Scopes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *InstanceTemplate) urlNormalized() *InstanceTemplate {
	normalized := dcl.Copy(*r).(InstanceTemplate)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Name = dcl.SelfLinkToName(r.Name)
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
	return unmarshalMapInstanceTemplate(m, c)
}

func unmarshalMapInstanceTemplate(m map[string]interface{}, c *Client) (*InstanceTemplate, error) {
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
	} else if v != nil {
		m["properties"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
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

	res := &InstanceTemplate{}
	res.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	res.Description = dcl.FlattenString(m["description"])
	res.Id = dcl.FlattenInteger(m["id"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Name = dcl.FlattenString(m["name"])
	res.Properties = flattenInstanceTemplateProperties(c, m["properties"])
	res.Project = dcl.FlattenString(m["project"])

	return res
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
		items = append(items, *flattenInstanceTemplatePropertiesDisksInterfaceEnum(item.(interface{})))
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
		items = append(items, *flattenInstanceTemplatePropertiesDisksModeEnum(item.(interface{})))
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
		items = append(items, *flattenInstanceTemplatePropertiesDisksTypeEnum(item.(interface{})))
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
		items = append(items, *flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(item.(interface{})))
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
		items = append(items, *flattenInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(item.(interface{})))
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
		items = append(items, *flattenInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(item.(interface{})))
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

func convertFieldDiffToInstanceTemplateDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]instanceTemplateDiff, error) {
	var diffs []instanceTemplateDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := instanceTemplateDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameToinstanceTemplateApiOperation(op, opts...)
				if err != nil {
					return nil, err
				}
				diff.UpdateOp = op
			}
			diffs = append(diffs, diff)
		}
	}
	return diffs, nil
}

func convertOpNameToinstanceTemplateApiOperation(op string, opts ...dcl.ApplyOption) (instanceTemplateApiOperation, error) {
	switch op {

	case "updateInstanceTemplateUpdateOperation":
		return &updateInstanceTemplateUpdateOperation{}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
