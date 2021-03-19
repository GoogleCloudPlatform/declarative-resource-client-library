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

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *InstanceGroupManager) validate() error {

	if err := dcl.Required(r, "baseInstanceName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DistributionPolicy) {
		if err := r.DistributionPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CurrentActions) {
		if err := r.CurrentActions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.UpdatePolicy) {
		if err := r.UpdatePolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerDistributionPolicy) validate() error {
	return nil
}
func (r *InstanceGroupManagerDistributionPolicyZones) validate() error {
	return nil
}
func (r *InstanceGroupManagerCurrentActions) validate() error {
	return nil
}
func (r *InstanceGroupManagerVersions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TargetSize) {
		if err := r.TargetSize.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerVersionsTargetSize) validate() error {
	return nil
}
func (r *InstanceGroupManagerNamedPorts) validate() error {
	return nil
}
func (r *InstanceGroupManagerStatus) validate() error {
	if !dcl.IsEmptyValueIndirect(r.VersionTarget) {
		if err := r.VersionTarget.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerStatusVersionTarget) validate() error {
	return nil
}
func (r *InstanceGroupManagerAutoHealingPolicies) validate() error {
	return nil
}
func (r *InstanceGroupManagerUpdatePolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MaxSurge) {
		if err := r.MaxSurge.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerUpdatePolicyMaxSurge) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MaxUnavailable) {
		if err := r.MaxUnavailable.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) validate() error {
	return nil
}

func instanceGroupManagerGetURL(userBasePath string, r *InstanceGroupManager) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Get URL found")

}

func instanceGroupManagerListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid List URL found")

}

func instanceGroupManagerCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Create URL found")

}

func instanceGroupManagerDeleteURL(userBasePath string, r *InstanceGroupManager) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Delete URL found")

}

// instanceGroupManagerApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceGroupManagerApiOperation interface {
	do(context.Context, *InstanceGroupManager, *Client) error
}

// newUpdateInstanceGroupManagerSetInstanceTemplateRequest creates a request for an
// InstanceGroupManager resource's setInstanceTemplate update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceGroupManagerSetInstanceTemplateRequest(ctx context.Context, f *InstanceGroupManager, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.InstanceTemplate; !dcl.IsEmptyValueIndirect(v) {
		req["instanceTemplate"] = v
	}
	return req, nil
}

// marshalUpdateInstanceGroupManagerSetInstanceTemplateRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceGroupManagerSetInstanceTemplateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceGroupManagerSetInstanceTemplateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceGroupManagerSetInstanceTemplateOperation) do(ctx context.Context, r *InstanceGroupManager, c *Client) error {
	_, err := c.GetInstanceGroupManager(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setInstanceTemplate")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceGroupManagerSetInstanceTemplateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceGroupManagerSetInstanceTemplateRequest(c, req)
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

// newUpdateInstanceGroupManagerSetTargetPoolsRequest creates a request for an
// InstanceGroupManager resource's setTargetPools update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceGroupManagerSetTargetPoolsRequest(ctx context.Context, f *InstanceGroupManager, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.TargetPools; !dcl.IsEmptyValueIndirect(v) {
		req["targetPools"] = v
	}
	return req, nil
}

// marshalUpdateInstanceGroupManagerSetTargetPoolsRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceGroupManagerSetTargetPoolsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceGroupManagerSetTargetPoolsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceGroupManagerSetTargetPoolsOperation) do(ctx context.Context, r *InstanceGroupManager, c *Client) error {
	_, err := c.GetInstanceGroupManager(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setTargetPools")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceGroupManagerSetTargetPoolsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceGroupManagerSetTargetPoolsRequest(c, req)
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

func (c *Client) listInstanceGroupManagerRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := instanceGroupManagerListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InstanceGroupManagerMaxPage {
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

type listInstanceGroupManagerOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listInstanceGroupManager(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*InstanceGroupManager, string, error) {
	b, err := c.listInstanceGroupManagerRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInstanceGroupManagerOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*InstanceGroupManager
	for _, v := range m.Items {
		res := flattenInstanceGroupManager(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInstanceGroupManager(ctx context.Context, f func(*InstanceGroupManager) bool, resources []*InstanceGroupManager) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInstanceGroupManager(ctx, res)
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

type deleteInstanceGroupManagerOperation struct{}

func (op *deleteInstanceGroupManagerOperation) do(ctx context.Context, r *InstanceGroupManager, c *Client) error {

	_, err := c.GetInstanceGroupManager(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("InstanceGroupManager not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetInstanceGroupManager checking for existence. error: %v", err)
		return err
	}

	u, err := instanceGroupManagerDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	_, err = c.GetInstanceGroupManager(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createInstanceGroupManagerOperation struct {
	response map[string]interface{}
}

func (op *createInstanceGroupManagerOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createInstanceGroupManagerOperation) do(ctx context.Context, r *InstanceGroupManager, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := instanceGroupManagerCreateURL(c.Config.BasePath, project, location)

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

	if _, err := c.GetInstanceGroupManager(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getInstanceGroupManagerRaw(ctx context.Context, r *InstanceGroupManager) ([]byte, error) {

	u, err := instanceGroupManagerGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) instanceGroupManagerDiffsForRawDesired(ctx context.Context, rawDesired *InstanceGroupManager, opts ...dcl.ApplyOption) (initial, desired *InstanceGroupManager, diffs []instanceGroupManagerDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *InstanceGroupManager
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*InstanceGroupManager); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected InstanceGroupManager, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInstanceGroupManager(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a InstanceGroupManager resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve InstanceGroupManager resource: %v", err)
		}
		c.Config.Logger.Info("Found that InstanceGroupManager resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInstanceGroupManagerDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for InstanceGroupManager: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for InstanceGroupManager: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInstanceGroupManagerInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for InstanceGroupManager: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInstanceGroupManagerDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for InstanceGroupManager: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInstanceGroupManager(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInstanceGroupManagerInitialState(rawInitial, rawDesired *InstanceGroupManager) (*InstanceGroupManager, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInstanceGroupManagerDesiredState(rawDesired, rawInitial *InstanceGroupManager, opts ...dcl.ApplyOption) (*InstanceGroupManager, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DistributionPolicy = canonicalizeInstanceGroupManagerDistributionPolicy(rawDesired.DistributionPolicy, nil, opts...)
		rawDesired.CurrentActions = canonicalizeInstanceGroupManagerCurrentActions(rawDesired.CurrentActions, nil, opts...)
		rawDesired.Status = canonicalizeInstanceGroupManagerStatus(rawDesired.Status, nil, opts...)
		rawDesired.UpdatePolicy = canonicalizeInstanceGroupManagerUpdatePolicy(rawDesired.UpdatePolicy, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.BaseInstanceName, rawInitial.BaseInstanceName) {
		rawDesired.BaseInstanceName = rawInitial.BaseInstanceName
	}
	if dcl.IsZeroValue(rawDesired.CreationTimestamp) {
		rawDesired.CreationTimestamp = rawInitial.CreationTimestamp
	}
	rawDesired.DistributionPolicy = canonicalizeInstanceGroupManagerDistributionPolicy(rawDesired.DistributionPolicy, rawInitial.DistributionPolicy, opts...)
	rawDesired.CurrentActions = canonicalizeInstanceGroupManagerCurrentActions(rawDesired.CurrentActions, rawInitial.CurrentActions, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Versions) {
		rawDesired.Versions = rawInitial.Versions
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.NameToSelfLink(rawDesired.InstanceGroup, rawInitial.InstanceGroup) {
		rawDesired.InstanceGroup = rawInitial.InstanceGroup
	}
	if dcl.NameToSelfLink(rawDesired.InstanceTemplate, rawInitial.InstanceTemplate) {
		rawDesired.InstanceTemplate = rawInitial.InstanceTemplate
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.NamedPorts) {
		rawDesired.NamedPorts = rawInitial.NamedPorts
	}
	rawDesired.Status = canonicalizeInstanceGroupManagerStatus(rawDesired.Status, rawInitial.Status, opts...)
	if dcl.IsZeroValue(rawDesired.TargetPools) {
		rawDesired.TargetPools = rawInitial.TargetPools
	}
	if dcl.IsZeroValue(rawDesired.AutoHealingPolicies) {
		rawDesired.AutoHealingPolicies = rawInitial.AutoHealingPolicies
	}
	rawDesired.UpdatePolicy = canonicalizeInstanceGroupManagerUpdatePolicy(rawDesired.UpdatePolicy, rawInitial.UpdatePolicy, opts...)
	if dcl.IsZeroValue(rawDesired.TargetSize) {
		rawDesired.TargetSize = rawInitial.TargetSize
	}
	if dcl.StringCanonicalize(rawDesired.Zone, rawInitial.Zone) {
		rawDesired.Zone = rawInitial.Zone
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeInstanceGroupManagerNewState(c *Client, rawNew, rawDesired *InstanceGroupManager) (*InstanceGroupManager, error) {

	if dcl.IsEmptyValueIndirect(rawNew.BaseInstanceName) && dcl.IsEmptyValueIndirect(rawDesired.BaseInstanceName) {
		rawNew.BaseInstanceName = rawDesired.BaseInstanceName
	} else {
		if dcl.StringCanonicalize(rawDesired.BaseInstanceName, rawNew.BaseInstanceName) {
			rawNew.BaseInstanceName = rawDesired.BaseInstanceName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DistributionPolicy) && dcl.IsEmptyValueIndirect(rawDesired.DistributionPolicy) {
		rawNew.DistributionPolicy = rawDesired.DistributionPolicy
	} else {
		rawNew.DistributionPolicy = canonicalizeNewInstanceGroupManagerDistributionPolicy(c, rawDesired.DistributionPolicy, rawNew.DistributionPolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CurrentActions) && dcl.IsEmptyValueIndirect(rawDesired.CurrentActions) {
		rawNew.CurrentActions = rawDesired.CurrentActions
	} else {
		rawNew.CurrentActions = canonicalizeNewInstanceGroupManagerCurrentActions(c, rawDesired.CurrentActions, rawNew.CurrentActions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Versions) && dcl.IsEmptyValueIndirect(rawDesired.Versions) {
		rawNew.Versions = rawDesired.Versions
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.InstanceGroup) && dcl.IsEmptyValueIndirect(rawDesired.InstanceGroup) {
		rawNew.InstanceGroup = rawDesired.InstanceGroup
	} else {
		if dcl.NameToSelfLink(rawDesired.InstanceGroup, rawNew.InstanceGroup) {
			rawNew.InstanceGroup = rawDesired.InstanceGroup
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.InstanceTemplate) && dcl.IsEmptyValueIndirect(rawDesired.InstanceTemplate) {
		rawNew.InstanceTemplate = rawDesired.InstanceTemplate
	} else {
		if dcl.NameToSelfLink(rawDesired.InstanceTemplate, rawNew.InstanceTemplate) {
			rawNew.InstanceTemplate = rawDesired.InstanceTemplate
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NamedPorts) && dcl.IsEmptyValueIndirect(rawDesired.NamedPorts) {
		rawNew.NamedPorts = rawDesired.NamedPorts
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
		rawNew.Status = canonicalizeNewInstanceGroupManagerStatus(c, rawDesired.Status, rawNew.Status)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TargetPools) && dcl.IsEmptyValueIndirect(rawDesired.TargetPools) {
		rawNew.TargetPools = rawDesired.TargetPools
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AutoHealingPolicies) && dcl.IsEmptyValueIndirect(rawDesired.AutoHealingPolicies) {
		rawNew.AutoHealingPolicies = rawDesired.AutoHealingPolicies
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdatePolicy) && dcl.IsEmptyValueIndirect(rawDesired.UpdatePolicy) {
		rawNew.UpdatePolicy = rawDesired.UpdatePolicy
	} else {
		rawNew.UpdatePolicy = canonicalizeNewInstanceGroupManagerUpdatePolicy(c, rawDesired.UpdatePolicy, rawNew.UpdatePolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TargetSize) && dcl.IsEmptyValueIndirect(rawDesired.TargetSize) {
		rawNew.TargetSize = rawDesired.TargetSize
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Zone) && dcl.IsEmptyValueIndirect(rawDesired.Zone) {
		rawNew.Zone = rawDesired.Zone
	} else {
		if dcl.StringCanonicalize(rawDesired.Zone, rawNew.Zone) {
			rawNew.Zone = rawDesired.Zone
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeInstanceGroupManagerDistributionPolicy(des, initial *InstanceGroupManagerDistributionPolicy, opts ...dcl.ApplyOption) *InstanceGroupManagerDistributionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Zones) {
		des.Zones = initial.Zones
	}

	return des
}

func canonicalizeNewInstanceGroupManagerDistributionPolicy(c *Client, des, nw *InstanceGroupManagerDistributionPolicy) *InstanceGroupManagerDistributionPolicy {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerDistributionPolicySet(c *Client, des, nw []InstanceGroupManagerDistributionPolicy) []InstanceGroupManagerDistributionPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerDistributionPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerDistributionPolicy(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerDistributionPolicyZones(des, initial *InstanceGroupManagerDistributionPolicyZones, opts ...dcl.ApplyOption) *InstanceGroupManagerDistributionPolicyZones {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Zone, initial.Zone) || dcl.IsZeroValue(des.Zone) {
		des.Zone = initial.Zone
	}

	return des
}

func canonicalizeNewInstanceGroupManagerDistributionPolicyZones(c *Client, des, nw *InstanceGroupManagerDistributionPolicyZones) *InstanceGroupManagerDistributionPolicyZones {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Zone, nw.Zone) || dcl.IsZeroValue(des.Zone) {
		nw.Zone = des.Zone
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerDistributionPolicyZonesSet(c *Client, des, nw []InstanceGroupManagerDistributionPolicyZones) []InstanceGroupManagerDistributionPolicyZones {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerDistributionPolicyZones
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerDistributionPolicyZones(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerCurrentActions(des, initial *InstanceGroupManagerCurrentActions, opts ...dcl.ApplyOption) *InstanceGroupManagerCurrentActions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Abandoning) {
		des.Abandoning = initial.Abandoning
	}
	if dcl.IsZeroValue(des.Creating) {
		des.Creating = initial.Creating
	}
	if dcl.IsZeroValue(des.CreatingWithoutRetries) {
		des.CreatingWithoutRetries = initial.CreatingWithoutRetries
	}
	if dcl.IsZeroValue(des.Deleting) {
		des.Deleting = initial.Deleting
	}
	if dcl.IsZeroValue(des.None) {
		des.None = initial.None
	}
	if dcl.IsZeroValue(des.Recreating) {
		des.Recreating = initial.Recreating
	}
	if dcl.IsZeroValue(des.Refreshing) {
		des.Refreshing = initial.Refreshing
	}
	if dcl.IsZeroValue(des.Restarting) {
		des.Restarting = initial.Restarting
	}

	return des
}

func canonicalizeNewInstanceGroupManagerCurrentActions(c *Client, des, nw *InstanceGroupManagerCurrentActions) *InstanceGroupManagerCurrentActions {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerCurrentActionsSet(c *Client, des, nw []InstanceGroupManagerCurrentActions) []InstanceGroupManagerCurrentActions {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerCurrentActions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerCurrentActions(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerVersions(des, initial *InstanceGroupManagerVersions, opts ...dcl.ApplyOption) *InstanceGroupManagerVersions {
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
	if dcl.NameToSelfLink(des.InstanceTemplate, initial.InstanceTemplate) || dcl.IsZeroValue(des.InstanceTemplate) {
		des.InstanceTemplate = initial.InstanceTemplate
	}
	des.TargetSize = canonicalizeInstanceGroupManagerVersionsTargetSize(des.TargetSize, initial.TargetSize, opts...)

	return des
}

func canonicalizeNewInstanceGroupManagerVersions(c *Client, des, nw *InstanceGroupManagerVersions) *InstanceGroupManagerVersions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}
	if dcl.NameToSelfLink(des.InstanceTemplate, nw.InstanceTemplate) || dcl.IsZeroValue(des.InstanceTemplate) {
		nw.InstanceTemplate = des.InstanceTemplate
	}
	nw.TargetSize = canonicalizeNewInstanceGroupManagerVersionsTargetSize(c, des.TargetSize, nw.TargetSize)

	return nw
}

func canonicalizeNewInstanceGroupManagerVersionsSet(c *Client, des, nw []InstanceGroupManagerVersions) []InstanceGroupManagerVersions {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerVersions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerVersions(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerVersionsTargetSize(des, initial *InstanceGroupManagerVersionsTargetSize, opts ...dcl.ApplyOption) *InstanceGroupManagerVersionsTargetSize {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Fixed) {
		des.Fixed = initial.Fixed
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}
	if dcl.IsZeroValue(des.Calculated) {
		des.Calculated = initial.Calculated
	}

	return des
}

func canonicalizeNewInstanceGroupManagerVersionsTargetSize(c *Client, des, nw *InstanceGroupManagerVersionsTargetSize) *InstanceGroupManagerVersionsTargetSize {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerVersionsTargetSizeSet(c *Client, des, nw []InstanceGroupManagerVersionsTargetSize) []InstanceGroupManagerVersionsTargetSize {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerVersionsTargetSize
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerVersionsTargetSize(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerNamedPorts(des, initial *InstanceGroupManagerNamedPorts, opts ...dcl.ApplyOption) *InstanceGroupManagerNamedPorts {
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
	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}

	return des
}

func canonicalizeNewInstanceGroupManagerNamedPorts(c *Client, des, nw *InstanceGroupManagerNamedPorts) *InstanceGroupManagerNamedPorts {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) || dcl.IsZeroValue(des.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerNamedPortsSet(c *Client, des, nw []InstanceGroupManagerNamedPorts) []InstanceGroupManagerNamedPorts {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerNamedPorts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerNamedPorts(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerStatus(des, initial *InstanceGroupManagerStatus, opts ...dcl.ApplyOption) *InstanceGroupManagerStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.IsStable) {
		des.IsStable = initial.IsStable
	}
	des.VersionTarget = canonicalizeInstanceGroupManagerStatusVersionTarget(des.VersionTarget, initial.VersionTarget, opts...)
	if dcl.NameToSelfLink(des.Autoscalar, initial.Autoscalar) || dcl.IsZeroValue(des.Autoscalar) {
		des.Autoscalar = initial.Autoscalar
	}

	return des
}

func canonicalizeNewInstanceGroupManagerStatus(c *Client, des, nw *InstanceGroupManagerStatus) *InstanceGroupManagerStatus {
	if des == nil || nw == nil {
		return nw
	}

	nw.VersionTarget = canonicalizeNewInstanceGroupManagerStatusVersionTarget(c, des.VersionTarget, nw.VersionTarget)
	if dcl.NameToSelfLink(des.Autoscalar, nw.Autoscalar) || dcl.IsZeroValue(des.Autoscalar) {
		nw.Autoscalar = des.Autoscalar
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerStatusSet(c *Client, des, nw []InstanceGroupManagerStatus) []InstanceGroupManagerStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerStatus(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerStatusVersionTarget(des, initial *InstanceGroupManagerStatusVersionTarget, opts ...dcl.ApplyOption) *InstanceGroupManagerStatusVersionTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.IsReached) {
		des.IsReached = initial.IsReached
	}

	return des
}

func canonicalizeNewInstanceGroupManagerStatusVersionTarget(c *Client, des, nw *InstanceGroupManagerStatusVersionTarget) *InstanceGroupManagerStatusVersionTarget {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerStatusVersionTargetSet(c *Client, des, nw []InstanceGroupManagerStatusVersionTarget) []InstanceGroupManagerStatusVersionTarget {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerStatusVersionTarget
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerStatusVersionTarget(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerAutoHealingPolicies(des, initial *InstanceGroupManagerAutoHealingPolicies, opts ...dcl.ApplyOption) *InstanceGroupManagerAutoHealingPolicies {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.HealthCheck, initial.HealthCheck) || dcl.IsZeroValue(des.HealthCheck) {
		des.HealthCheck = initial.HealthCheck
	}
	if dcl.IsZeroValue(des.InitialDelaySec) {
		des.InitialDelaySec = initial.InitialDelaySec
	}

	return des
}

func canonicalizeNewInstanceGroupManagerAutoHealingPolicies(c *Client, des, nw *InstanceGroupManagerAutoHealingPolicies) *InstanceGroupManagerAutoHealingPolicies {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.HealthCheck, nw.HealthCheck) || dcl.IsZeroValue(des.HealthCheck) {
		nw.HealthCheck = des.HealthCheck
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerAutoHealingPoliciesSet(c *Client, des, nw []InstanceGroupManagerAutoHealingPolicies) []InstanceGroupManagerAutoHealingPolicies {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerAutoHealingPolicies
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerAutoHealingPolicies(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerUpdatePolicy(des, initial *InstanceGroupManagerUpdatePolicy, opts ...dcl.ApplyOption) *InstanceGroupManagerUpdatePolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.InstanceRedistributionType) {
		des.InstanceRedistributionType = initial.InstanceRedistributionType
	}
	if dcl.IsZeroValue(des.MinimalAction) {
		des.MinimalAction = initial.MinimalAction
	}
	des.MaxSurge = canonicalizeInstanceGroupManagerUpdatePolicyMaxSurge(des.MaxSurge, initial.MaxSurge, opts...)

	return des
}

func canonicalizeNewInstanceGroupManagerUpdatePolicy(c *Client, des, nw *InstanceGroupManagerUpdatePolicy) *InstanceGroupManagerUpdatePolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.MaxSurge = canonicalizeNewInstanceGroupManagerUpdatePolicyMaxSurge(c, des.MaxSurge, nw.MaxSurge)

	return nw
}

func canonicalizeNewInstanceGroupManagerUpdatePolicySet(c *Client, des, nw []InstanceGroupManagerUpdatePolicy) []InstanceGroupManagerUpdatePolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerUpdatePolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerUpdatePolicy(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerUpdatePolicyMaxSurge(des, initial *InstanceGroupManagerUpdatePolicyMaxSurge, opts ...dcl.ApplyOption) *InstanceGroupManagerUpdatePolicyMaxSurge {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Fixed) {
		des.Fixed = initial.Fixed
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}
	if dcl.IsZeroValue(des.Calculated) {
		des.Calculated = initial.Calculated
	}
	des.MaxUnavailable = canonicalizeInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(des.MaxUnavailable, initial.MaxUnavailable, opts...)

	return des
}

func canonicalizeNewInstanceGroupManagerUpdatePolicyMaxSurge(c *Client, des, nw *InstanceGroupManagerUpdatePolicyMaxSurge) *InstanceGroupManagerUpdatePolicyMaxSurge {
	if des == nil || nw == nil {
		return nw
	}

	nw.MaxUnavailable = canonicalizeNewInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, des.MaxUnavailable, nw.MaxUnavailable)

	return nw
}

func canonicalizeNewInstanceGroupManagerUpdatePolicyMaxSurgeSet(c *Client, des, nw []InstanceGroupManagerUpdatePolicyMaxSurge) []InstanceGroupManagerUpdatePolicyMaxSurge {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerUpdatePolicyMaxSurge
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerUpdatePolicyMaxSurge(c, &d, &n) {
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

func canonicalizeInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(des, initial *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable, opts ...dcl.ApplyOption) *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Fixed) {
		des.Fixed = initial.Fixed
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	}
	if dcl.IsZeroValue(des.Calculated) {
		des.Calculated = initial.Calculated
	}

	return des
}

func canonicalizeNewInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c *Client, des, nw *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableSet(c *Client, des, nw []InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) []InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, &d, &n) {
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

type instanceGroupManagerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceGroupManagerApiOperation
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
func diffInstanceGroupManager(c *Client, desired, actual *InstanceGroupManager, opts ...dcl.ApplyOption) ([]instanceGroupManagerDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []instanceGroupManagerDiff
	if !dcl.IsZeroValue(desired.BaseInstanceName) && !dcl.StringCanonicalize(desired.BaseInstanceName, actual.BaseInstanceName) {
		c.Config.Logger.Infof("Detected diff in BaseInstanceName.\nDESIRED: %v\nACTUAL: %v", desired.BaseInstanceName, actual.BaseInstanceName)
		diffs = append(diffs, instanceGroupManagerDiff{
			RequiresRecreate: true,
			FieldName:        "BaseInstanceName",
		})
	}
	if compareInstanceGroupManagerDistributionPolicy(c, desired.DistributionPolicy, actual.DistributionPolicy) {
		c.Config.Logger.Infof("Detected diff in DistributionPolicy.\nDESIRED: %v\nACTUAL: %v", desired.DistributionPolicy, actual.DistributionPolicy)
		diffs = append(diffs, instanceGroupManagerDiff{
			RequiresRecreate: true,
			FieldName:        "DistributionPolicy",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, instanceGroupManagerDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if compareInstanceGroupManagerVersionsSlice(c, desired.Versions, actual.Versions) {
		c.Config.Logger.Infof("Detected diff in Versions.\nDESIRED: %v\nACTUAL: %v", desired.Versions, actual.Versions)
		diffs = append(diffs, instanceGroupManagerDiff{
			RequiresRecreate: true,
			FieldName:        "Versions",
		})
	}
	if !dcl.IsZeroValue(desired.InstanceTemplate) && !dcl.NameToSelfLink(desired.InstanceTemplate, actual.InstanceTemplate) {
		c.Config.Logger.Infof("Detected diff in InstanceTemplate.\nDESIRED: %v\nACTUAL: %v", desired.InstanceTemplate, actual.InstanceTemplate)

		diffs = append(diffs, instanceGroupManagerDiff{
			UpdateOp:  &updateInstanceGroupManagerSetInstanceTemplateOperation{},
			FieldName: "InstanceTemplate",
		})

	}
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, instanceGroupManagerDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if compareInstanceGroupManagerNamedPortsSlice(c, desired.NamedPorts, actual.NamedPorts) {
		c.Config.Logger.Infof("Detected diff in NamedPorts.\nDESIRED: %v\nACTUAL: %v", desired.NamedPorts, actual.NamedPorts)
		diffs = append(diffs, instanceGroupManagerDiff{
			RequiresRecreate: true,
			FieldName:        "NamedPorts",
		})
	}
	if !reflect.DeepEqual(desired.TargetPools, actual.TargetPools) {
		c.Config.Logger.Infof("Detected diff in TargetPools.\nDESIRED: %v\nACTUAL: %v", desired.TargetPools, actual.TargetPools)

		diffs = append(diffs, instanceGroupManagerDiff{
			UpdateOp:  &updateInstanceGroupManagerSetTargetPoolsOperation{},
			FieldName: "TargetPools",
		})

	}
	if compareInstanceGroupManagerAutoHealingPoliciesSlice(c, desired.AutoHealingPolicies, actual.AutoHealingPolicies) {
		c.Config.Logger.Infof("Detected diff in AutoHealingPolicies.\nDESIRED: %v\nACTUAL: %v", desired.AutoHealingPolicies, actual.AutoHealingPolicies)
		diffs = append(diffs, instanceGroupManagerDiff{
			RequiresRecreate: true,
			FieldName:        "AutoHealingPolicies",
		})
	}
	if compareInstanceGroupManagerUpdatePolicy(c, desired.UpdatePolicy, actual.UpdatePolicy) {
		c.Config.Logger.Infof("Detected diff in UpdatePolicy.\nDESIRED: %v\nACTUAL: %v", desired.UpdatePolicy, actual.UpdatePolicy)
		diffs = append(diffs, instanceGroupManagerDiff{
			RequiresRecreate: true,
			FieldName:        "UpdatePolicy",
		})
	}
	if !reflect.DeepEqual(desired.TargetSize, actual.TargetSize) {
		c.Config.Logger.Infof("Detected diff in TargetSize.\nDESIRED: %v\nACTUAL: %v", desired.TargetSize, actual.TargetSize)
		diffs = append(diffs, instanceGroupManagerDiff{
			RequiresRecreate: true,
			FieldName:        "TargetSize",
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
	var deduped []instanceGroupManagerDiff
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
func compareInstanceGroupManagerDistributionPolicy(c *Client, desired, actual *InstanceGroupManagerDistributionPolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Zones == nil && desired.Zones != nil && !dcl.IsEmptyValueIndirect(desired.Zones) {
		c.Config.Logger.Infof("desired Zones %s - but actually nil", dcl.SprintResource(desired.Zones))
		return true
	}
	if compareInstanceGroupManagerDistributionPolicyZonesSlice(c, desired.Zones, actual.Zones) && !dcl.IsZeroValue(desired.Zones) {
		c.Config.Logger.Infof("Diff in Zones. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Zones), dcl.SprintResource(actual.Zones))
		return true
	}
	return false
}

func compareInstanceGroupManagerDistributionPolicySlice(c *Client, desired, actual []InstanceGroupManagerDistributionPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerDistributionPolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerDistributionPolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerDistributionPolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerDistributionPolicyMap(c *Client, desired, actual map[string]InstanceGroupManagerDistributionPolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerDistributionPolicy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerDistributionPolicy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerDistributionPolicy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerDistributionPolicy, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerDistributionPolicyZones(c *Client, desired, actual *InstanceGroupManagerDistributionPolicyZones) bool {
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
	if !dcl.StringCanonicalize(desired.Zone, actual.Zone) && !dcl.IsZeroValue(desired.Zone) {
		c.Config.Logger.Infof("Diff in Zone. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Zone), dcl.SprintResource(actual.Zone))
		return true
	}
	return false
}

func compareInstanceGroupManagerDistributionPolicyZonesSlice(c *Client, desired, actual []InstanceGroupManagerDistributionPolicyZones) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerDistributionPolicyZones, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerDistributionPolicyZones(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerDistributionPolicyZones, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerDistributionPolicyZonesMap(c *Client, desired, actual map[string]InstanceGroupManagerDistributionPolicyZones) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerDistributionPolicyZones, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerDistributionPolicyZones, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerDistributionPolicyZones(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerDistributionPolicyZones, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerCurrentActions(c *Client, desired, actual *InstanceGroupManagerCurrentActions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}

func compareInstanceGroupManagerCurrentActionsSlice(c *Client, desired, actual []InstanceGroupManagerCurrentActions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerCurrentActions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerCurrentActions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerCurrentActions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerCurrentActionsMap(c *Client, desired, actual map[string]InstanceGroupManagerCurrentActions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerCurrentActions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerCurrentActions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerCurrentActions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerCurrentActions, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerVersions(c *Client, desired, actual *InstanceGroupManagerVersions) bool {
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
	if actual.InstanceTemplate == nil && desired.InstanceTemplate != nil && !dcl.IsEmptyValueIndirect(desired.InstanceTemplate) {
		c.Config.Logger.Infof("desired InstanceTemplate %s - but actually nil", dcl.SprintResource(desired.InstanceTemplate))
		return true
	}
	if !dcl.NameToSelfLink(desired.InstanceTemplate, actual.InstanceTemplate) && !dcl.IsZeroValue(desired.InstanceTemplate) {
		c.Config.Logger.Infof("Diff in InstanceTemplate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InstanceTemplate), dcl.SprintResource(actual.InstanceTemplate))
		return true
	}
	if actual.TargetSize == nil && desired.TargetSize != nil && !dcl.IsEmptyValueIndirect(desired.TargetSize) {
		c.Config.Logger.Infof("desired TargetSize %s - but actually nil", dcl.SprintResource(desired.TargetSize))
		return true
	}
	if compareInstanceGroupManagerVersionsTargetSize(c, desired.TargetSize, actual.TargetSize) && !dcl.IsZeroValue(desired.TargetSize) {
		c.Config.Logger.Infof("Diff in TargetSize. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TargetSize), dcl.SprintResource(actual.TargetSize))
		return true
	}
	return false
}

func compareInstanceGroupManagerVersionsSlice(c *Client, desired, actual []InstanceGroupManagerVersions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerVersions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerVersions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerVersions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerVersionsMap(c *Client, desired, actual map[string]InstanceGroupManagerVersions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerVersions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerVersions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerVersions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerVersions, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerVersionsTargetSize(c *Client, desired, actual *InstanceGroupManagerVersionsTargetSize) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Fixed == nil && desired.Fixed != nil && !dcl.IsEmptyValueIndirect(desired.Fixed) {
		c.Config.Logger.Infof("desired Fixed %s - but actually nil", dcl.SprintResource(desired.Fixed))
		return true
	}
	if !reflect.DeepEqual(desired.Fixed, actual.Fixed) && !dcl.IsZeroValue(desired.Fixed) {
		c.Config.Logger.Infof("Diff in Fixed. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Fixed), dcl.SprintResource(actual.Fixed))
		return true
	}
	if actual.Percent == nil && desired.Percent != nil && !dcl.IsEmptyValueIndirect(desired.Percent) {
		c.Config.Logger.Infof("desired Percent %s - but actually nil", dcl.SprintResource(desired.Percent))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) {
		c.Config.Logger.Infof("Diff in Percent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	return false
}

func compareInstanceGroupManagerVersionsTargetSizeSlice(c *Client, desired, actual []InstanceGroupManagerVersionsTargetSize) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerVersionsTargetSize, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerVersionsTargetSize(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerVersionsTargetSize, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerVersionsTargetSizeMap(c *Client, desired, actual map[string]InstanceGroupManagerVersionsTargetSize) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerVersionsTargetSize, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerVersionsTargetSize, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerVersionsTargetSize(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerVersionsTargetSize, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerNamedPorts(c *Client, desired, actual *InstanceGroupManagerNamedPorts) bool {
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
	if actual.Port == nil && desired.Port != nil && !dcl.IsEmptyValueIndirect(desired.Port) {
		c.Config.Logger.Infof("desired Port %s - but actually nil", dcl.SprintResource(desired.Port))
		return true
	}
	if !reflect.DeepEqual(desired.Port, actual.Port) && !dcl.IsZeroValue(desired.Port) {
		c.Config.Logger.Infof("Diff in Port. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Port), dcl.SprintResource(actual.Port))
		return true
	}
	return false
}

func compareInstanceGroupManagerNamedPortsSlice(c *Client, desired, actual []InstanceGroupManagerNamedPorts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerNamedPorts, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerNamedPorts(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerNamedPorts, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerNamedPortsMap(c *Client, desired, actual map[string]InstanceGroupManagerNamedPorts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerNamedPorts, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerNamedPorts, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerNamedPorts(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerNamedPorts, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerStatus(c *Client, desired, actual *InstanceGroupManagerStatus) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.IsStable == nil && desired.IsStable != nil && !dcl.IsEmptyValueIndirect(desired.IsStable) {
		c.Config.Logger.Infof("desired IsStable %s - but actually nil", dcl.SprintResource(desired.IsStable))
		return true
	}
	if !reflect.DeepEqual(desired.IsStable, actual.IsStable) && !dcl.IsZeroValue(desired.IsStable) {
		c.Config.Logger.Infof("Diff in IsStable. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IsStable), dcl.SprintResource(actual.IsStable))
		return true
	}
	if actual.VersionTarget == nil && desired.VersionTarget != nil && !dcl.IsEmptyValueIndirect(desired.VersionTarget) {
		c.Config.Logger.Infof("desired VersionTarget %s - but actually nil", dcl.SprintResource(desired.VersionTarget))
		return true
	}
	if compareInstanceGroupManagerStatusVersionTarget(c, desired.VersionTarget, actual.VersionTarget) && !dcl.IsZeroValue(desired.VersionTarget) {
		c.Config.Logger.Infof("Diff in VersionTarget. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.VersionTarget), dcl.SprintResource(actual.VersionTarget))
		return true
	}
	if actual.Autoscalar == nil && desired.Autoscalar != nil && !dcl.IsEmptyValueIndirect(desired.Autoscalar) {
		c.Config.Logger.Infof("desired Autoscalar %s - but actually nil", dcl.SprintResource(desired.Autoscalar))
		return true
	}
	if !dcl.NameToSelfLink(desired.Autoscalar, actual.Autoscalar) && !dcl.IsZeroValue(desired.Autoscalar) {
		c.Config.Logger.Infof("Diff in Autoscalar. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Autoscalar), dcl.SprintResource(actual.Autoscalar))
		return true
	}
	return false
}

func compareInstanceGroupManagerStatusSlice(c *Client, desired, actual []InstanceGroupManagerStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerStatus, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerStatus(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerStatus, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerStatusMap(c *Client, desired, actual map[string]InstanceGroupManagerStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerStatus, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerStatus, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerStatus(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerStatus, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerStatusVersionTarget(c *Client, desired, actual *InstanceGroupManagerStatusVersionTarget) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.IsReached == nil && desired.IsReached != nil && !dcl.IsEmptyValueIndirect(desired.IsReached) {
		c.Config.Logger.Infof("desired IsReached %s - but actually nil", dcl.SprintResource(desired.IsReached))
		return true
	}
	if !reflect.DeepEqual(desired.IsReached, actual.IsReached) && !dcl.IsZeroValue(desired.IsReached) {
		c.Config.Logger.Infof("Diff in IsReached. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IsReached), dcl.SprintResource(actual.IsReached))
		return true
	}
	return false
}

func compareInstanceGroupManagerStatusVersionTargetSlice(c *Client, desired, actual []InstanceGroupManagerStatusVersionTarget) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerStatusVersionTarget, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerStatusVersionTarget(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerStatusVersionTarget, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerStatusVersionTargetMap(c *Client, desired, actual map[string]InstanceGroupManagerStatusVersionTarget) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerStatusVersionTarget, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerStatusVersionTarget, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerStatusVersionTarget(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerStatusVersionTarget, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerAutoHealingPolicies(c *Client, desired, actual *InstanceGroupManagerAutoHealingPolicies) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.HealthCheck == nil && desired.HealthCheck != nil && !dcl.IsEmptyValueIndirect(desired.HealthCheck) {
		c.Config.Logger.Infof("desired HealthCheck %s - but actually nil", dcl.SprintResource(desired.HealthCheck))
		return true
	}
	if !dcl.NameToSelfLink(desired.HealthCheck, actual.HealthCheck) && !dcl.IsZeroValue(desired.HealthCheck) {
		c.Config.Logger.Infof("Diff in HealthCheck. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HealthCheck), dcl.SprintResource(actual.HealthCheck))
		return true
	}
	if actual.InitialDelaySec == nil && desired.InitialDelaySec != nil && !dcl.IsEmptyValueIndirect(desired.InitialDelaySec) {
		c.Config.Logger.Infof("desired InitialDelaySec %s - but actually nil", dcl.SprintResource(desired.InitialDelaySec))
		return true
	}
	if !reflect.DeepEqual(desired.InitialDelaySec, actual.InitialDelaySec) && !dcl.IsZeroValue(desired.InitialDelaySec) {
		c.Config.Logger.Infof("Diff in InitialDelaySec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InitialDelaySec), dcl.SprintResource(actual.InitialDelaySec))
		return true
	}
	return false
}

func compareInstanceGroupManagerAutoHealingPoliciesSlice(c *Client, desired, actual []InstanceGroupManagerAutoHealingPolicies) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerAutoHealingPolicies, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerAutoHealingPolicies(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerAutoHealingPolicies, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerAutoHealingPoliciesMap(c *Client, desired, actual map[string]InstanceGroupManagerAutoHealingPolicies) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerAutoHealingPolicies, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerAutoHealingPolicies, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerAutoHealingPolicies(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerAutoHealingPolicies, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicy(c *Client, desired, actual *InstanceGroupManagerUpdatePolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.InstanceRedistributionType == nil && desired.InstanceRedistributionType != nil && !dcl.IsEmptyValueIndirect(desired.InstanceRedistributionType) {
		c.Config.Logger.Infof("desired InstanceRedistributionType %s - but actually nil", dcl.SprintResource(desired.InstanceRedistributionType))
		return true
	}
	if !reflect.DeepEqual(desired.InstanceRedistributionType, actual.InstanceRedistributionType) && !dcl.IsZeroValue(desired.InstanceRedistributionType) {
		c.Config.Logger.Infof("Diff in InstanceRedistributionType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InstanceRedistributionType), dcl.SprintResource(actual.InstanceRedistributionType))
		return true
	}
	if actual.MinimalAction == nil && desired.MinimalAction != nil && !dcl.IsEmptyValueIndirect(desired.MinimalAction) {
		c.Config.Logger.Infof("desired MinimalAction %s - but actually nil", dcl.SprintResource(desired.MinimalAction))
		return true
	}
	if !reflect.DeepEqual(desired.MinimalAction, actual.MinimalAction) && !dcl.IsZeroValue(desired.MinimalAction) {
		c.Config.Logger.Infof("Diff in MinimalAction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinimalAction), dcl.SprintResource(actual.MinimalAction))
		return true
	}
	if actual.MaxSurge == nil && desired.MaxSurge != nil && !dcl.IsEmptyValueIndirect(desired.MaxSurge) {
		c.Config.Logger.Infof("desired MaxSurge %s - but actually nil", dcl.SprintResource(desired.MaxSurge))
		return true
	}
	if compareInstanceGroupManagerUpdatePolicyMaxSurge(c, desired.MaxSurge, actual.MaxSurge) && !dcl.IsZeroValue(desired.MaxSurge) {
		c.Config.Logger.Infof("Diff in MaxSurge. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxSurge), dcl.SprintResource(actual.MaxSurge))
		return true
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicySlice(c *Client, desired, actual []InstanceGroupManagerUpdatePolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerUpdatePolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerUpdatePolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyMap(c *Client, desired, actual map[string]InstanceGroupManagerUpdatePolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerUpdatePolicy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerUpdatePolicy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicy, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyMaxSurge(c *Client, desired, actual *InstanceGroupManagerUpdatePolicyMaxSurge) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Fixed == nil && desired.Fixed != nil && !dcl.IsEmptyValueIndirect(desired.Fixed) {
		c.Config.Logger.Infof("desired Fixed %s - but actually nil", dcl.SprintResource(desired.Fixed))
		return true
	}
	if !reflect.DeepEqual(desired.Fixed, actual.Fixed) && !dcl.IsZeroValue(desired.Fixed) {
		c.Config.Logger.Infof("Diff in Fixed. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Fixed), dcl.SprintResource(actual.Fixed))
		return true
	}
	if actual.Percent == nil && desired.Percent != nil && !dcl.IsEmptyValueIndirect(desired.Percent) {
		c.Config.Logger.Infof("desired Percent %s - but actually nil", dcl.SprintResource(desired.Percent))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) {
		c.Config.Logger.Infof("Diff in Percent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	if actual.MaxUnavailable == nil && desired.MaxUnavailable != nil && !dcl.IsEmptyValueIndirect(desired.MaxUnavailable) {
		c.Config.Logger.Infof("desired MaxUnavailable %s - but actually nil", dcl.SprintResource(desired.MaxUnavailable))
		return true
	}
	if compareInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, desired.MaxUnavailable, actual.MaxUnavailable) && !dcl.IsZeroValue(desired.MaxUnavailable) {
		c.Config.Logger.Infof("Diff in MaxUnavailable. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxUnavailable), dcl.SprintResource(actual.MaxUnavailable))
		return true
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyMaxSurgeSlice(c *Client, desired, actual []InstanceGroupManagerUpdatePolicyMaxSurge) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerUpdatePolicyMaxSurge, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerUpdatePolicyMaxSurge(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicyMaxSurge, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyMaxSurgeMap(c *Client, desired, actual map[string]InstanceGroupManagerUpdatePolicyMaxSurge) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerUpdatePolicyMaxSurge, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicyMaxSurge, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerUpdatePolicyMaxSurge(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicyMaxSurge, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c *Client, desired, actual *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Fixed == nil && desired.Fixed != nil && !dcl.IsEmptyValueIndirect(desired.Fixed) {
		c.Config.Logger.Infof("desired Fixed %s - but actually nil", dcl.SprintResource(desired.Fixed))
		return true
	}
	if !reflect.DeepEqual(desired.Fixed, actual.Fixed) && !dcl.IsZeroValue(desired.Fixed) {
		c.Config.Logger.Infof("Diff in Fixed. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Fixed), dcl.SprintResource(actual.Fixed))
		return true
	}
	if actual.Percent == nil && desired.Percent != nil && !dcl.IsEmptyValueIndirect(desired.Percent) {
		c.Config.Logger.Infof("desired Percent %s - but actually nil", dcl.SprintResource(desired.Percent))
		return true
	}
	if !reflect.DeepEqual(desired.Percent, actual.Percent) && !dcl.IsZeroValue(desired.Percent) {
		c.Config.Logger.Infof("Diff in Percent. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Percent), dcl.SprintResource(actual.Percent))
		return true
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableSlice(c *Client, desired, actual []InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableMap(c *Client, desired, actual map[string]InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumSlice(c *Client, desired, actual []InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(c *Client, desired, actual *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceGroupManagerUpdatePolicyMinimalActionEnumSlice(c *Client, desired, actual []InstanceGroupManagerUpdatePolicyMinimalActionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceGroupManagerUpdatePolicyMinimalActionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceGroupManagerUpdatePolicyMinimalActionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceGroupManagerUpdatePolicyMinimalActionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceGroupManagerUpdatePolicyMinimalActionEnum(c *Client, desired, actual *InstanceGroupManagerUpdatePolicyMinimalActionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *InstanceGroupManager) urlNormalized() *InstanceGroupManager {
	normalized := deepcopy.Copy(*r).(InstanceGroupManager)
	normalized.BaseInstanceName = dcl.SelfLinkToName(r.BaseInstanceName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.InstanceGroup = dcl.SelfLinkToName(r.InstanceGroup)
	normalized.InstanceTemplate = dcl.SelfLinkToName(r.InstanceTemplate)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *InstanceGroupManager) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *InstanceGroupManager) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *InstanceGroupManager) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *InstanceGroupManager) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "setInstanceTemplate" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		if dcl.IsRegion(r.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers/{{name}}/setInstanceTemplate", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil
		}

		if dcl.IsZone(r.Location) {
			return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers/{{name}}/setInstanceTemplate", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil
		}

		return "", fmt.Errorf("No valid update URL for %s", updateName)

	}
	if updateName == "setTargetPools" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		if dcl.IsRegion(r.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers/{{name}}/setTargetPools", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil
		}

		if dcl.IsZone(r.Location) {
			return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers/{{name}}/setTargetPools", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil
		}

		return "", fmt.Errorf("No valid update URL for %s", updateName)

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the InstanceGroupManager resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *InstanceGroupManager) marshal(c *Client) ([]byte, error) {
	m, err := expandInstanceGroupManager(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling InstanceGroupManager: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalInstanceGroupManager decodes JSON responses into the InstanceGroupManager resource schema.
func unmarshalInstanceGroupManager(b []byte, c *Client) (*InstanceGroupManager, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapInstanceGroupManager(m, c)
}

func unmarshalMapInstanceGroupManager(m map[string]interface{}, c *Client) (*InstanceGroupManager, error) {

	return flattenInstanceGroupManager(c, m), nil
}

// expandInstanceGroupManager expands InstanceGroupManager into a JSON request object.
func expandInstanceGroupManager(c *Client, f *InstanceGroupManager) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.BaseInstanceName; !dcl.IsEmptyValueIndirect(v) {
		m["baseInstanceName"] = v
	}
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v, err := expandInstanceGroupManagerDistributionPolicy(c, f.DistributionPolicy); err != nil {
		return nil, fmt.Errorf("error expanding DistributionPolicy into distributionPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["distributionPolicy"] = v
	}
	if v, err := expandInstanceGroupManagerCurrentActions(c, f.CurrentActions); err != nil {
		return nil, fmt.Errorf("error expanding CurrentActions into currentActions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["currentActions"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandInstanceGroupManagerVersionsSlice(c, f.Versions); err != nil {
		return nil, fmt.Errorf("error expanding Versions into versions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["versions"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.InstanceGroup; !dcl.IsEmptyValueIndirect(v) {
		m["instanceGroup"] = v
	}
	if v := f.InstanceTemplate; !dcl.IsEmptyValueIndirect(v) {
		m["instanceTemplate"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandInstanceGroupManagerNamedPortsSlice(c, f.NamedPorts); err != nil {
		return nil, fmt.Errorf("error expanding NamedPorts into namedPorts: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["namedPorts"] = v
	}
	if v, err := expandInstanceGroupManagerStatus(c, f.Status); err != nil {
		return nil, fmt.Errorf("error expanding Status into status: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.TargetPools; !dcl.IsEmptyValueIndirect(v) {
		m["targetPools"] = v
	}
	if v, err := expandInstanceGroupManagerAutoHealingPoliciesSlice(c, f.AutoHealingPolicies); err != nil {
		return nil, fmt.Errorf("error expanding AutoHealingPolicies into autoHealingPolicies: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoHealingPolicies"] = v
	}
	if v, err := expandInstanceGroupManagerUpdatePolicy(c, f.UpdatePolicy); err != nil {
		return nil, fmt.Errorf("error expanding UpdatePolicy into updatePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["updatePolicy"] = v
	}
	if v := f.TargetSize; !dcl.IsEmptyValueIndirect(v) {
		m["targetSize"] = v
	}
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenInstanceGroupManager flattens InstanceGroupManager from a JSON request object into the
// InstanceGroupManager type.
func flattenInstanceGroupManager(c *Client, i interface{}) *InstanceGroupManager {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &InstanceGroupManager{}
	r.BaseInstanceName = dcl.FlattenString(m["baseInstanceName"])
	r.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	r.DistributionPolicy = flattenInstanceGroupManagerDistributionPolicy(c, m["distributionPolicy"])
	r.CurrentActions = flattenInstanceGroupManagerCurrentActions(c, m["currentActions"])
	r.Description = dcl.FlattenString(m["description"])
	r.Versions = flattenInstanceGroupManagerVersionsSlice(c, m["versions"])
	r.Id = dcl.FlattenInteger(m["id"])
	r.InstanceGroup = dcl.FlattenString(m["instanceGroup"])
	r.InstanceTemplate = dcl.FlattenString(m["instanceTemplate"])
	r.Name = dcl.FlattenString(m["name"])
	r.NamedPorts = flattenInstanceGroupManagerNamedPortsSlice(c, m["namedPorts"])
	r.Status = flattenInstanceGroupManagerStatus(c, m["status"])
	r.TargetPools = dcl.FlattenStringSlice(m["targetPools"])
	r.AutoHealingPolicies = flattenInstanceGroupManagerAutoHealingPoliciesSlice(c, m["autoHealingPolicies"])
	r.UpdatePolicy = flattenInstanceGroupManagerUpdatePolicy(c, m["updatePolicy"])
	r.TargetSize = dcl.FlattenInteger(m["targetSize"])
	r.Zone = dcl.FlattenString(m["zone"])
	r.Region = dcl.FlattenString(m["region"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandInstanceGroupManagerDistributionPolicyMap expands the contents of InstanceGroupManagerDistributionPolicy into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicyMap(c *Client, f map[string]InstanceGroupManagerDistributionPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerDistributionPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerDistributionPolicySlice expands the contents of InstanceGroupManagerDistributionPolicy into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicySlice(c *Client, f []InstanceGroupManagerDistributionPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerDistributionPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerDistributionPolicyMap flattens the contents of InstanceGroupManagerDistributionPolicy from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyMap(c *Client, i interface{}) map[string]InstanceGroupManagerDistributionPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerDistributionPolicy{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerDistributionPolicy{}
	}

	items := make(map[string]InstanceGroupManagerDistributionPolicy)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerDistributionPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerDistributionPolicySlice flattens the contents of InstanceGroupManagerDistributionPolicy from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicySlice(c *Client, i interface{}) []InstanceGroupManagerDistributionPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerDistributionPolicy{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerDistributionPolicy{}
	}

	items := make([]InstanceGroupManagerDistributionPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerDistributionPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerDistributionPolicy expands an instance of InstanceGroupManagerDistributionPolicy into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicy(c *Client, f *InstanceGroupManagerDistributionPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInstanceGroupManagerDistributionPolicyZonesSlice(c, f.Zones); err != nil {
		return nil, fmt.Errorf("error expanding Zones into zones: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["zones"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerDistributionPolicy flattens an instance of InstanceGroupManagerDistributionPolicy from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicy(c *Client, i interface{}) *InstanceGroupManagerDistributionPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerDistributionPolicy{}
	r.Zones = flattenInstanceGroupManagerDistributionPolicyZonesSlice(c, m["zones"])

	return r
}

// expandInstanceGroupManagerDistributionPolicyZonesMap expands the contents of InstanceGroupManagerDistributionPolicyZones into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicyZonesMap(c *Client, f map[string]InstanceGroupManagerDistributionPolicyZones) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerDistributionPolicyZones(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerDistributionPolicyZonesSlice expands the contents of InstanceGroupManagerDistributionPolicyZones into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicyZonesSlice(c *Client, f []InstanceGroupManagerDistributionPolicyZones) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerDistributionPolicyZones(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerDistributionPolicyZonesMap flattens the contents of InstanceGroupManagerDistributionPolicyZones from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyZonesMap(c *Client, i interface{}) map[string]InstanceGroupManagerDistributionPolicyZones {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerDistributionPolicyZones{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerDistributionPolicyZones{}
	}

	items := make(map[string]InstanceGroupManagerDistributionPolicyZones)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerDistributionPolicyZones(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerDistributionPolicyZonesSlice flattens the contents of InstanceGroupManagerDistributionPolicyZones from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyZonesSlice(c *Client, i interface{}) []InstanceGroupManagerDistributionPolicyZones {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerDistributionPolicyZones{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerDistributionPolicyZones{}
	}

	items := make([]InstanceGroupManagerDistributionPolicyZones, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerDistributionPolicyZones(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerDistributionPolicyZones expands an instance of InstanceGroupManagerDistributionPolicyZones into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicyZones(c *Client, f *InstanceGroupManagerDistributionPolicyZones) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerDistributionPolicyZones flattens an instance of InstanceGroupManagerDistributionPolicyZones from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyZones(c *Client, i interface{}) *InstanceGroupManagerDistributionPolicyZones {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerDistributionPolicyZones{}
	r.Zone = dcl.FlattenString(m["zone"])

	return r
}

// expandInstanceGroupManagerCurrentActionsMap expands the contents of InstanceGroupManagerCurrentActions into a JSON
// request object.
func expandInstanceGroupManagerCurrentActionsMap(c *Client, f map[string]InstanceGroupManagerCurrentActions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerCurrentActions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerCurrentActionsSlice expands the contents of InstanceGroupManagerCurrentActions into a JSON
// request object.
func expandInstanceGroupManagerCurrentActionsSlice(c *Client, f []InstanceGroupManagerCurrentActions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerCurrentActions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerCurrentActionsMap flattens the contents of InstanceGroupManagerCurrentActions from a JSON
// response object.
func flattenInstanceGroupManagerCurrentActionsMap(c *Client, i interface{}) map[string]InstanceGroupManagerCurrentActions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerCurrentActions{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerCurrentActions{}
	}

	items := make(map[string]InstanceGroupManagerCurrentActions)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerCurrentActions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerCurrentActionsSlice flattens the contents of InstanceGroupManagerCurrentActions from a JSON
// response object.
func flattenInstanceGroupManagerCurrentActionsSlice(c *Client, i interface{}) []InstanceGroupManagerCurrentActions {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerCurrentActions{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerCurrentActions{}
	}

	items := make([]InstanceGroupManagerCurrentActions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerCurrentActions(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerCurrentActions expands an instance of InstanceGroupManagerCurrentActions into a JSON
// request object.
func expandInstanceGroupManagerCurrentActions(c *Client, f *InstanceGroupManagerCurrentActions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Abandoning; !dcl.IsEmptyValueIndirect(v) {
		m["abandoning"] = v
	}
	if v := f.Creating; !dcl.IsEmptyValueIndirect(v) {
		m["creating"] = v
	}
	if v := f.CreatingWithoutRetries; !dcl.IsEmptyValueIndirect(v) {
		m["creatingWithoutRetries"] = v
	}
	if v := f.Deleting; !dcl.IsEmptyValueIndirect(v) {
		m["deleting"] = v
	}
	if v := f.None; !dcl.IsEmptyValueIndirect(v) {
		m["none"] = v
	}
	if v := f.Recreating; !dcl.IsEmptyValueIndirect(v) {
		m["recreating"] = v
	}
	if v := f.Refreshing; !dcl.IsEmptyValueIndirect(v) {
		m["refreshing"] = v
	}
	if v := f.Restarting; !dcl.IsEmptyValueIndirect(v) {
		m["restarting"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerCurrentActions flattens an instance of InstanceGroupManagerCurrentActions from a JSON
// response object.
func flattenInstanceGroupManagerCurrentActions(c *Client, i interface{}) *InstanceGroupManagerCurrentActions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerCurrentActions{}
	r.Abandoning = dcl.FlattenInteger(m["abandoning"])
	r.Creating = dcl.FlattenInteger(m["creating"])
	r.CreatingWithoutRetries = dcl.FlattenInteger(m["creatingWithoutRetries"])
	r.Deleting = dcl.FlattenInteger(m["deleting"])
	r.None = dcl.FlattenInteger(m["none"])
	r.Recreating = dcl.FlattenInteger(m["recreating"])
	r.Refreshing = dcl.FlattenInteger(m["refreshing"])
	r.Restarting = dcl.FlattenInteger(m["restarting"])

	return r
}

// expandInstanceGroupManagerVersionsMap expands the contents of InstanceGroupManagerVersions into a JSON
// request object.
func expandInstanceGroupManagerVersionsMap(c *Client, f map[string]InstanceGroupManagerVersions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerVersions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerVersionsSlice expands the contents of InstanceGroupManagerVersions into a JSON
// request object.
func expandInstanceGroupManagerVersionsSlice(c *Client, f []InstanceGroupManagerVersions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerVersions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerVersionsMap flattens the contents of InstanceGroupManagerVersions from a JSON
// response object.
func flattenInstanceGroupManagerVersionsMap(c *Client, i interface{}) map[string]InstanceGroupManagerVersions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerVersions{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerVersions{}
	}

	items := make(map[string]InstanceGroupManagerVersions)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerVersions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerVersionsSlice flattens the contents of InstanceGroupManagerVersions from a JSON
// response object.
func flattenInstanceGroupManagerVersionsSlice(c *Client, i interface{}) []InstanceGroupManagerVersions {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerVersions{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerVersions{}
	}

	items := make([]InstanceGroupManagerVersions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerVersions(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerVersions expands an instance of InstanceGroupManagerVersions into a JSON
// request object.
func expandInstanceGroupManagerVersions(c *Client, f *InstanceGroupManagerVersions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.InstanceTemplate; !dcl.IsEmptyValueIndirect(v) {
		m["instanceTemplate"] = v
	}
	if v, err := expandInstanceGroupManagerVersionsTargetSize(c, f.TargetSize); err != nil {
		return nil, fmt.Errorf("error expanding TargetSize into targetSize: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["targetSize"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerVersions flattens an instance of InstanceGroupManagerVersions from a JSON
// response object.
func flattenInstanceGroupManagerVersions(c *Client, i interface{}) *InstanceGroupManagerVersions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerVersions{}
	r.Name = dcl.FlattenString(m["name"])
	r.InstanceTemplate = dcl.FlattenString(m["instanceTemplate"])
	r.TargetSize = flattenInstanceGroupManagerVersionsTargetSize(c, m["targetSize"])

	return r
}

// expandInstanceGroupManagerVersionsTargetSizeMap expands the contents of InstanceGroupManagerVersionsTargetSize into a JSON
// request object.
func expandInstanceGroupManagerVersionsTargetSizeMap(c *Client, f map[string]InstanceGroupManagerVersionsTargetSize) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerVersionsTargetSize(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerVersionsTargetSizeSlice expands the contents of InstanceGroupManagerVersionsTargetSize into a JSON
// request object.
func expandInstanceGroupManagerVersionsTargetSizeSlice(c *Client, f []InstanceGroupManagerVersionsTargetSize) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerVersionsTargetSize(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerVersionsTargetSizeMap flattens the contents of InstanceGroupManagerVersionsTargetSize from a JSON
// response object.
func flattenInstanceGroupManagerVersionsTargetSizeMap(c *Client, i interface{}) map[string]InstanceGroupManagerVersionsTargetSize {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerVersionsTargetSize{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerVersionsTargetSize{}
	}

	items := make(map[string]InstanceGroupManagerVersionsTargetSize)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerVersionsTargetSize(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerVersionsTargetSizeSlice flattens the contents of InstanceGroupManagerVersionsTargetSize from a JSON
// response object.
func flattenInstanceGroupManagerVersionsTargetSizeSlice(c *Client, i interface{}) []InstanceGroupManagerVersionsTargetSize {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerVersionsTargetSize{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerVersionsTargetSize{}
	}

	items := make([]InstanceGroupManagerVersionsTargetSize, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerVersionsTargetSize(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerVersionsTargetSize expands an instance of InstanceGroupManagerVersionsTargetSize into a JSON
// request object.
func expandInstanceGroupManagerVersionsTargetSize(c *Client, f *InstanceGroupManagerVersionsTargetSize) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fixed; !dcl.IsEmptyValueIndirect(v) {
		m["fixed"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}
	if v := f.Calculated; !dcl.IsEmptyValueIndirect(v) {
		m["calculated"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerVersionsTargetSize flattens an instance of InstanceGroupManagerVersionsTargetSize from a JSON
// response object.
func flattenInstanceGroupManagerVersionsTargetSize(c *Client, i interface{}) *InstanceGroupManagerVersionsTargetSize {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerVersionsTargetSize{}
	r.Fixed = dcl.FlattenInteger(m["fixed"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Calculated = dcl.FlattenInteger(m["calculated"])

	return r
}

// expandInstanceGroupManagerNamedPortsMap expands the contents of InstanceGroupManagerNamedPorts into a JSON
// request object.
func expandInstanceGroupManagerNamedPortsMap(c *Client, f map[string]InstanceGroupManagerNamedPorts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerNamedPorts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerNamedPortsSlice expands the contents of InstanceGroupManagerNamedPorts into a JSON
// request object.
func expandInstanceGroupManagerNamedPortsSlice(c *Client, f []InstanceGroupManagerNamedPorts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerNamedPorts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerNamedPortsMap flattens the contents of InstanceGroupManagerNamedPorts from a JSON
// response object.
func flattenInstanceGroupManagerNamedPortsMap(c *Client, i interface{}) map[string]InstanceGroupManagerNamedPorts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerNamedPorts{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerNamedPorts{}
	}

	items := make(map[string]InstanceGroupManagerNamedPorts)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerNamedPorts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerNamedPortsSlice flattens the contents of InstanceGroupManagerNamedPorts from a JSON
// response object.
func flattenInstanceGroupManagerNamedPortsSlice(c *Client, i interface{}) []InstanceGroupManagerNamedPorts {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerNamedPorts{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerNamedPorts{}
	}

	items := make([]InstanceGroupManagerNamedPorts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerNamedPorts(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerNamedPorts expands an instance of InstanceGroupManagerNamedPorts into a JSON
// request object.
func expandInstanceGroupManagerNamedPorts(c *Client, f *InstanceGroupManagerNamedPorts) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerNamedPorts flattens an instance of InstanceGroupManagerNamedPorts from a JSON
// response object.
func flattenInstanceGroupManagerNamedPorts(c *Client, i interface{}) *InstanceGroupManagerNamedPorts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerNamedPorts{}
	r.Name = dcl.FlattenString(m["name"])
	r.Port = dcl.FlattenInteger(m["port"])

	return r
}

// expandInstanceGroupManagerStatusMap expands the contents of InstanceGroupManagerStatus into a JSON
// request object.
func expandInstanceGroupManagerStatusMap(c *Client, f map[string]InstanceGroupManagerStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerStatus(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerStatusSlice expands the contents of InstanceGroupManagerStatus into a JSON
// request object.
func expandInstanceGroupManagerStatusSlice(c *Client, f []InstanceGroupManagerStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerStatus(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerStatusMap flattens the contents of InstanceGroupManagerStatus from a JSON
// response object.
func flattenInstanceGroupManagerStatusMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatus{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatus{}
	}

	items := make(map[string]InstanceGroupManagerStatus)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatusSlice flattens the contents of InstanceGroupManagerStatus from a JSON
// response object.
func flattenInstanceGroupManagerStatusSlice(c *Client, i interface{}) []InstanceGroupManagerStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatus{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatus{}
	}

	items := make([]InstanceGroupManagerStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerStatus expands an instance of InstanceGroupManagerStatus into a JSON
// request object.
func expandInstanceGroupManagerStatus(c *Client, f *InstanceGroupManagerStatus) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IsStable; !dcl.IsEmptyValueIndirect(v) {
		m["isStable"] = v
	}
	if v, err := expandInstanceGroupManagerStatusVersionTarget(c, f.VersionTarget); err != nil {
		return nil, fmt.Errorf("error expanding VersionTarget into versionTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["versionTarget"] = v
	}
	if v := f.Autoscalar; !dcl.IsEmptyValueIndirect(v) {
		m["autoscalar"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerStatus flattens an instance of InstanceGroupManagerStatus from a JSON
// response object.
func flattenInstanceGroupManagerStatus(c *Client, i interface{}) *InstanceGroupManagerStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerStatus{}
	r.IsStable = dcl.FlattenBool(m["isStable"])
	r.VersionTarget = flattenInstanceGroupManagerStatusVersionTarget(c, m["versionTarget"])
	r.Autoscalar = dcl.FlattenString(m["autoscalar"])

	return r
}

// expandInstanceGroupManagerStatusVersionTargetMap expands the contents of InstanceGroupManagerStatusVersionTarget into a JSON
// request object.
func expandInstanceGroupManagerStatusVersionTargetMap(c *Client, f map[string]InstanceGroupManagerStatusVersionTarget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerStatusVersionTarget(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerStatusVersionTargetSlice expands the contents of InstanceGroupManagerStatusVersionTarget into a JSON
// request object.
func expandInstanceGroupManagerStatusVersionTargetSlice(c *Client, f []InstanceGroupManagerStatusVersionTarget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerStatusVersionTarget(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerStatusVersionTargetMap flattens the contents of InstanceGroupManagerStatusVersionTarget from a JSON
// response object.
func flattenInstanceGroupManagerStatusVersionTargetMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatusVersionTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatusVersionTarget{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatusVersionTarget{}
	}

	items := make(map[string]InstanceGroupManagerStatusVersionTarget)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatusVersionTarget(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatusVersionTargetSlice flattens the contents of InstanceGroupManagerStatusVersionTarget from a JSON
// response object.
func flattenInstanceGroupManagerStatusVersionTargetSlice(c *Client, i interface{}) []InstanceGroupManagerStatusVersionTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatusVersionTarget{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatusVersionTarget{}
	}

	items := make([]InstanceGroupManagerStatusVersionTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatusVersionTarget(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerStatusVersionTarget expands an instance of InstanceGroupManagerStatusVersionTarget into a JSON
// request object.
func expandInstanceGroupManagerStatusVersionTarget(c *Client, f *InstanceGroupManagerStatusVersionTarget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IsReached; !dcl.IsEmptyValueIndirect(v) {
		m["isReached"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerStatusVersionTarget flattens an instance of InstanceGroupManagerStatusVersionTarget from a JSON
// response object.
func flattenInstanceGroupManagerStatusVersionTarget(c *Client, i interface{}) *InstanceGroupManagerStatusVersionTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerStatusVersionTarget{}
	r.IsReached = dcl.FlattenBool(m["isReached"])

	return r
}

// expandInstanceGroupManagerAutoHealingPoliciesMap expands the contents of InstanceGroupManagerAutoHealingPolicies into a JSON
// request object.
func expandInstanceGroupManagerAutoHealingPoliciesMap(c *Client, f map[string]InstanceGroupManagerAutoHealingPolicies) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerAutoHealingPolicies(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerAutoHealingPoliciesSlice expands the contents of InstanceGroupManagerAutoHealingPolicies into a JSON
// request object.
func expandInstanceGroupManagerAutoHealingPoliciesSlice(c *Client, f []InstanceGroupManagerAutoHealingPolicies) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerAutoHealingPolicies(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerAutoHealingPoliciesMap flattens the contents of InstanceGroupManagerAutoHealingPolicies from a JSON
// response object.
func flattenInstanceGroupManagerAutoHealingPoliciesMap(c *Client, i interface{}) map[string]InstanceGroupManagerAutoHealingPolicies {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerAutoHealingPolicies{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerAutoHealingPolicies{}
	}

	items := make(map[string]InstanceGroupManagerAutoHealingPolicies)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerAutoHealingPolicies(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerAutoHealingPoliciesSlice flattens the contents of InstanceGroupManagerAutoHealingPolicies from a JSON
// response object.
func flattenInstanceGroupManagerAutoHealingPoliciesSlice(c *Client, i interface{}) []InstanceGroupManagerAutoHealingPolicies {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerAutoHealingPolicies{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerAutoHealingPolicies{}
	}

	items := make([]InstanceGroupManagerAutoHealingPolicies, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerAutoHealingPolicies(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerAutoHealingPolicies expands an instance of InstanceGroupManagerAutoHealingPolicies into a JSON
// request object.
func expandInstanceGroupManagerAutoHealingPolicies(c *Client, f *InstanceGroupManagerAutoHealingPolicies) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HealthCheck; !dcl.IsEmptyValueIndirect(v) {
		m["healthCheck"] = v
	}
	if v := f.InitialDelaySec; !dcl.IsEmptyValueIndirect(v) {
		m["initialDelaySec"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerAutoHealingPolicies flattens an instance of InstanceGroupManagerAutoHealingPolicies from a JSON
// response object.
func flattenInstanceGroupManagerAutoHealingPolicies(c *Client, i interface{}) *InstanceGroupManagerAutoHealingPolicies {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerAutoHealingPolicies{}
	r.HealthCheck = dcl.FlattenString(m["healthCheck"])
	r.InitialDelaySec = dcl.FlattenInteger(m["initialDelaySec"])

	return r
}

// expandInstanceGroupManagerUpdatePolicyMap expands the contents of InstanceGroupManagerUpdatePolicy into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicyMap(c *Client, f map[string]InstanceGroupManagerUpdatePolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerUpdatePolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerUpdatePolicySlice expands the contents of InstanceGroupManagerUpdatePolicy into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicySlice(c *Client, f []InstanceGroupManagerUpdatePolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerUpdatePolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerUpdatePolicyMap flattens the contents of InstanceGroupManagerUpdatePolicy from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMap(c *Client, i interface{}) map[string]InstanceGroupManagerUpdatePolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerUpdatePolicy{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerUpdatePolicy{}
	}

	items := make(map[string]InstanceGroupManagerUpdatePolicy)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerUpdatePolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicySlice flattens the contents of InstanceGroupManagerUpdatePolicy from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicySlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicy{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicy{}
	}

	items := make([]InstanceGroupManagerUpdatePolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerUpdatePolicy expands an instance of InstanceGroupManagerUpdatePolicy into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicy(c *Client, f *InstanceGroupManagerUpdatePolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.InstanceRedistributionType; !dcl.IsEmptyValueIndirect(v) {
		m["instanceRedistributionType"] = v
	}
	if v := f.MinimalAction; !dcl.IsEmptyValueIndirect(v) {
		m["minimalAction"] = v
	}
	if v, err := expandInstanceGroupManagerUpdatePolicyMaxSurge(c, f.MaxSurge); err != nil {
		return nil, fmt.Errorf("error expanding MaxSurge into maxSurge: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maxSurge"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerUpdatePolicy flattens an instance of InstanceGroupManagerUpdatePolicy from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicy(c *Client, i interface{}) *InstanceGroupManagerUpdatePolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerUpdatePolicy{}
	r.InstanceRedistributionType = flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(m["instanceRedistributionType"])
	r.MinimalAction = flattenInstanceGroupManagerUpdatePolicyMinimalActionEnum(m["minimalAction"])
	r.MaxSurge = flattenInstanceGroupManagerUpdatePolicyMaxSurge(c, m["maxSurge"])

	return r
}

// expandInstanceGroupManagerUpdatePolicyMaxSurgeMap expands the contents of InstanceGroupManagerUpdatePolicyMaxSurge into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicyMaxSurgeMap(c *Client, f map[string]InstanceGroupManagerUpdatePolicyMaxSurge) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerUpdatePolicyMaxSurge(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerUpdatePolicyMaxSurgeSlice expands the contents of InstanceGroupManagerUpdatePolicyMaxSurge into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicyMaxSurgeSlice(c *Client, f []InstanceGroupManagerUpdatePolicyMaxSurge) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerUpdatePolicyMaxSurge(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerUpdatePolicyMaxSurgeMap flattens the contents of InstanceGroupManagerUpdatePolicyMaxSurge from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMaxSurgeMap(c *Client, i interface{}) map[string]InstanceGroupManagerUpdatePolicyMaxSurge {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerUpdatePolicyMaxSurge{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerUpdatePolicyMaxSurge{}
	}

	items := make(map[string]InstanceGroupManagerUpdatePolicyMaxSurge)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerUpdatePolicyMaxSurge(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyMaxSurgeSlice flattens the contents of InstanceGroupManagerUpdatePolicyMaxSurge from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMaxSurgeSlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicyMaxSurge {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicyMaxSurge{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicyMaxSurge{}
	}

	items := make([]InstanceGroupManagerUpdatePolicyMaxSurge, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicyMaxSurge(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerUpdatePolicyMaxSurge expands an instance of InstanceGroupManagerUpdatePolicyMaxSurge into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicyMaxSurge(c *Client, f *InstanceGroupManagerUpdatePolicyMaxSurge) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fixed; !dcl.IsEmptyValueIndirect(v) {
		m["fixed"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}
	if v := f.Calculated; !dcl.IsEmptyValueIndirect(v) {
		m["calculated"] = v
	}
	if v, err := expandInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, f.MaxUnavailable); err != nil {
		return nil, fmt.Errorf("error expanding MaxUnavailable into maxUnavailable: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maxUnavailable"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerUpdatePolicyMaxSurge flattens an instance of InstanceGroupManagerUpdatePolicyMaxSurge from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMaxSurge(c *Client, i interface{}) *InstanceGroupManagerUpdatePolicyMaxSurge {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerUpdatePolicyMaxSurge{}
	r.Fixed = dcl.FlattenInteger(m["fixed"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Calculated = dcl.FlattenInteger(m["calculated"])
	r.MaxUnavailable = flattenInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, m["maxUnavailable"])

	return r
}

// expandInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableMap expands the contents of InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableMap(c *Client, f map[string]InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableSlice expands the contents of InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableSlice(c *Client, f []InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableMap flattens the contents of InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableMap(c *Client, i interface{}) map[string]InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable{}
	}

	items := make(map[string]InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableSlice flattens the contents of InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableSlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable{}
	}

	items := make([]InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable expands an instance of InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c *Client, f *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fixed; !dcl.IsEmptyValueIndirect(v) {
		m["fixed"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}
	if v := f.Calculated; !dcl.IsEmptyValueIndirect(v) {
		m["calculated"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable flattens an instance of InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(c *Client, i interface{}) *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable{}
	r.Fixed = dcl.FlattenInteger(m["fixed"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Calculated = dcl.FlattenInteger(m["calculated"])

	return r
}

// flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumSlice flattens the contents of InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumSlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum{}
	}

	items := make([]InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum with the same value as that string.
func flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(i interface{}) *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumRef("")
	}

	return InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumRef(s)
}

// flattenInstanceGroupManagerUpdatePolicyMinimalActionEnumSlice flattens the contents of InstanceGroupManagerUpdatePolicyMinimalActionEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMinimalActionEnumSlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicyMinimalActionEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicyMinimalActionEnum{}
	}

	items := make([]InstanceGroupManagerUpdatePolicyMinimalActionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicyMinimalActionEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyMinimalActionEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerUpdatePolicyMinimalActionEnum with the same value as that string.
func flattenInstanceGroupManagerUpdatePolicyMinimalActionEnum(i interface{}) *InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerUpdatePolicyMinimalActionEnumRef("")
	}

	return InstanceGroupManagerUpdatePolicyMinimalActionEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *InstanceGroupManager) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInstanceGroupManager(b, c)
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
