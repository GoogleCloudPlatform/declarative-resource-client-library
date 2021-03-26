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
package redis

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

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "tier"); err != nil {
		return err
	}
	if err := dcl.Required(r, "memorySizeGb"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MaintenancePolicy) {
		if err := r.MaintenancePolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaintenanceSchedule) {
		if err := r.MaintenanceSchedule.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceServerCaCerts) validate() error {
	return nil
}
func (r *InstanceMaintenancePolicy) validate() error {
	return nil
}
func (r *InstanceMaintenancePolicyWeeklyMaintenanceWindow) validate() error {
	if err := dcl.Required(r, "day"); err != nil {
		return err
	}
	if err := dcl.Required(r, "startTime"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.StartTime) {
		if err := r.StartTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) validate() error {
	return nil
}
func (r *InstanceMaintenanceSchedule) validate() error {
	return nil
}

func instanceGetURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", "https://redis.googleapis.com/v1/", userBasePath, params), nil
}

func instanceListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances", "https://redis.googleapis.com/v1/", userBasePath, params), nil

}

func instanceCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances?instanceId={{name}}", "https://redis.googleapis.com/v1/", userBasePath, params), nil

}

func instanceDeleteURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", "https://redis.googleapis.com/v1/", userBasePath, params), nil
}

// instanceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceApiOperation interface {
	do(context.Context, *Instance, *Client) error
}

// newUpdateInstanceUpdateInstanceRequest creates a request for an
// Instance resource's UpdateInstance update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceUpdateInstanceRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/locations/%s/instances/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.LocationId; !dcl.IsEmptyValueIndirect(v) {
		req["locationId"] = v
	}
	if v := f.AlternativeLocationId; !dcl.IsEmptyValueIndirect(v) {
		req["alternativeLocationId"] = v
	}
	if v := f.RedisVersion; !dcl.IsEmptyValueIndirect(v) {
		req["redisVersion"] = v
	}
	if v := f.ReservedIPRange; !dcl.IsEmptyValueIndirect(v) {
		req["reservedIpRange"] = v
	}
	if v := f.RedisConfigs; !dcl.IsEmptyValueIndirect(v) {
		req["redisConfigs"] = v
	}
	if v := f.Tier; !dcl.IsEmptyValueIndirect(v) {
		req["tier"] = v
	}
	if v := f.MemorySizeGb; !dcl.IsEmptyValueIndirect(v) {
		req["memorySizeGb"] = v
	}
	if v := f.AuthorizedNetwork; !dcl.IsEmptyValueIndirect(v) {
		req["authorizedNetwork"] = v
	}
	if v := f.ConnectMode; !dcl.IsEmptyValueIndirect(v) {
		req["connectMode"] = v
	}
	if v := f.AuthEnabled; !dcl.IsEmptyValueIndirect(v) {
		req["authEnabled"] = v
	}
	if v := f.TransitEncryptionMode; !dcl.IsEmptyValueIndirect(v) {
		req["transitEncryptionMode"] = v
	}
	if v, err := expandInstanceMaintenancePolicy(c, f.MaintenancePolicy); err != nil {
		return nil, fmt.Errorf("error expanding MaintenancePolicy into maintenancePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["maintenancePolicy"] = v
	}
	return req, nil
}

// marshalUpdateInstanceUpdateInstanceRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceUpdateInstanceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceUpdateInstanceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceUpdateInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateInstance")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"name", "displayName", "labels", "locationId", "alternativeLocationId", "redisVersion", "reservedIpRange", "redisConfigs", "tier", "memorySizeGb", "authorizedNetwork", "connectMode", "authEnabled", "transitEncryptionMode", "maintenancePolicy"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceUpdateInstanceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceUpdateInstanceRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://redis.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInstanceRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := instanceListURL(c.Config.BasePath, project, location)
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
	Instances []map[string]interface{} `json:"instances"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listInstance(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Instance, string, error) {
	b, err := c.listInstanceRaw(ctx, project, location, pageToken, pageSize)
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
		res.Location = &location
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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://redis.googleapis.com/v1/", "GET"); err != nil {
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

	project, location, name := r.createFields()
	u, err := instanceCreateURL(c.Config.BasePath, project, location, name)

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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://redis.googleapis.com/v1/", "GET"); err != nil {
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
		rawDesired.MaintenancePolicy = canonicalizeInstanceMaintenancePolicy(rawDesired.MaintenancePolicy, nil, opts...)
		rawDesired.MaintenanceSchedule = canonicalizeInstanceMaintenanceSchedule(rawDesired.MaintenanceSchedule, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.StringCanonicalize(rawDesired.LocationId, rawInitial.LocationId) {
		rawDesired.LocationId = rawInitial.LocationId
	}
	if dcl.StringCanonicalize(rawDesired.AlternativeLocationId, rawInitial.AlternativeLocationId) {
		rawDesired.AlternativeLocationId = rawInitial.AlternativeLocationId
	}
	if dcl.StringCanonicalize(rawDesired.RedisVersion, rawInitial.RedisVersion) {
		rawDesired.RedisVersion = rawInitial.RedisVersion
	}
	if dcl.StringCanonicalize(rawDesired.ReservedIPRange, rawInitial.ReservedIPRange) {
		rawDesired.ReservedIPRange = rawInitial.ReservedIPRange
	}
	if dcl.StringCanonicalize(rawDesired.Host, rawInitial.Host) {
		rawDesired.Host = rawInitial.Host
	}
	if dcl.IsZeroValue(rawDesired.Port) {
		rawDesired.Port = rawInitial.Port
	}
	if dcl.StringCanonicalize(rawDesired.CurrentLocationId, rawInitial.CurrentLocationId) {
		rawDesired.CurrentLocationId = rawInitial.CurrentLocationId
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	if dcl.StringCanonicalize(rawDesired.StatusMessage, rawInitial.StatusMessage) {
		rawDesired.StatusMessage = rawInitial.StatusMessage
	}
	if dcl.IsZeroValue(rawDesired.RedisConfigs) {
		rawDesired.RedisConfigs = rawInitial.RedisConfigs
	}
	if dcl.IsZeroValue(rawDesired.Tier) {
		rawDesired.Tier = rawInitial.Tier
	}
	if dcl.IsZeroValue(rawDesired.MemorySizeGb) {
		rawDesired.MemorySizeGb = rawInitial.MemorySizeGb
	}
	if dcl.StringCanonicalize(rawDesired.AuthorizedNetwork, rawInitial.AuthorizedNetwork) {
		rawDesired.AuthorizedNetwork = rawInitial.AuthorizedNetwork
	}
	if dcl.StringCanonicalize(rawDesired.PersistenceIamIdentity, rawInitial.PersistenceIamIdentity) {
		rawDesired.PersistenceIamIdentity = rawInitial.PersistenceIamIdentity
	}
	if dcl.IsZeroValue(rawDesired.ConnectMode) {
		rawDesired.ConnectMode = rawInitial.ConnectMode
	}
	if dcl.BoolCanonicalize(rawDesired.AuthEnabled, rawInitial.AuthEnabled) {
		rawDesired.AuthEnabled = rawInitial.AuthEnabled
	}
	if dcl.IsZeroValue(rawDesired.ServerCaCerts) {
		rawDesired.ServerCaCerts = rawInitial.ServerCaCerts
	}
	if dcl.IsZeroValue(rawDesired.TransitEncryptionMode) {
		rawDesired.TransitEncryptionMode = rawInitial.TransitEncryptionMode
	}
	rawDesired.MaintenancePolicy = canonicalizeInstanceMaintenancePolicy(rawDesired.MaintenancePolicy, rawInitial.MaintenancePolicy, opts...)
	rawDesired.MaintenanceSchedule = canonicalizeInstanceMaintenanceSchedule(rawDesired.MaintenanceSchedule, rawInitial.MaintenanceSchedule, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.NameToSelfLink(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}

	return rawDesired, nil
}

func canonicalizeInstanceNewState(c *Client, rawNew, rawDesired *Instance) (*Instance, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LocationId) && dcl.IsEmptyValueIndirect(rawDesired.LocationId) {
		rawNew.LocationId = rawDesired.LocationId
	} else {
		if dcl.StringCanonicalize(rawDesired.LocationId, rawNew.LocationId) {
			rawNew.LocationId = rawDesired.LocationId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AlternativeLocationId) && dcl.IsEmptyValueIndirect(rawDesired.AlternativeLocationId) {
		rawNew.AlternativeLocationId = rawDesired.AlternativeLocationId
	} else {
		if dcl.StringCanonicalize(rawDesired.AlternativeLocationId, rawNew.AlternativeLocationId) {
			rawNew.AlternativeLocationId = rawDesired.AlternativeLocationId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RedisVersion) && dcl.IsEmptyValueIndirect(rawDesired.RedisVersion) {
		rawNew.RedisVersion = rawDesired.RedisVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.RedisVersion, rawNew.RedisVersion) {
			rawNew.RedisVersion = rawDesired.RedisVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReservedIPRange) && dcl.IsEmptyValueIndirect(rawDesired.ReservedIPRange) {
		rawNew.ReservedIPRange = rawDesired.ReservedIPRange
	} else {
		if dcl.StringCanonicalize(rawDesired.ReservedIPRange, rawNew.ReservedIPRange) {
			rawNew.ReservedIPRange = rawDesired.ReservedIPRange
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Host) && dcl.IsEmptyValueIndirect(rawDesired.Host) {
		rawNew.Host = rawDesired.Host
	} else {
		if dcl.StringCanonicalize(rawDesired.Host, rawNew.Host) {
			rawNew.Host = rawDesired.Host
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Port) && dcl.IsEmptyValueIndirect(rawDesired.Port) {
		rawNew.Port = rawDesired.Port
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CurrentLocationId) && dcl.IsEmptyValueIndirect(rawDesired.CurrentLocationId) {
		rawNew.CurrentLocationId = rawDesired.CurrentLocationId
	} else {
		if dcl.StringCanonicalize(rawDesired.CurrentLocationId, rawNew.CurrentLocationId) {
			rawNew.CurrentLocationId = rawDesired.CurrentLocationId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StatusMessage) && dcl.IsEmptyValueIndirect(rawDesired.StatusMessage) {
		rawNew.StatusMessage = rawDesired.StatusMessage
	} else {
		if dcl.StringCanonicalize(rawDesired.StatusMessage, rawNew.StatusMessage) {
			rawNew.StatusMessage = rawDesired.StatusMessage
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RedisConfigs) && dcl.IsEmptyValueIndirect(rawDesired.RedisConfigs) {
		rawNew.RedisConfigs = rawDesired.RedisConfigs
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Tier) && dcl.IsEmptyValueIndirect(rawDesired.Tier) {
		rawNew.Tier = rawDesired.Tier
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MemorySizeGb) && dcl.IsEmptyValueIndirect(rawDesired.MemorySizeGb) {
		rawNew.MemorySizeGb = rawDesired.MemorySizeGb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AuthorizedNetwork) && dcl.IsEmptyValueIndirect(rawDesired.AuthorizedNetwork) {
		rawNew.AuthorizedNetwork = rawDesired.AuthorizedNetwork
	} else {
		if dcl.StringCanonicalize(rawDesired.AuthorizedNetwork, rawNew.AuthorizedNetwork) {
			rawNew.AuthorizedNetwork = rawDesired.AuthorizedNetwork
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PersistenceIamIdentity) && dcl.IsEmptyValueIndirect(rawDesired.PersistenceIamIdentity) {
		rawNew.PersistenceIamIdentity = rawDesired.PersistenceIamIdentity
	} else {
		if dcl.StringCanonicalize(rawDesired.PersistenceIamIdentity, rawNew.PersistenceIamIdentity) {
			rawNew.PersistenceIamIdentity = rawDesired.PersistenceIamIdentity
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ConnectMode) && dcl.IsEmptyValueIndirect(rawDesired.ConnectMode) {
		rawNew.ConnectMode = rawDesired.ConnectMode
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AuthEnabled) && dcl.IsEmptyValueIndirect(rawDesired.AuthEnabled) {
		rawNew.AuthEnabled = rawDesired.AuthEnabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.AuthEnabled, rawNew.AuthEnabled) {
			rawNew.AuthEnabled = rawDesired.AuthEnabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServerCaCerts) && dcl.IsEmptyValueIndirect(rawDesired.ServerCaCerts) {
		rawNew.ServerCaCerts = rawDesired.ServerCaCerts
	} else {
		rawNew.ServerCaCerts = canonicalizeNewInstanceServerCaCertsSlice(c, rawDesired.ServerCaCerts, rawNew.ServerCaCerts)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TransitEncryptionMode) && dcl.IsEmptyValueIndirect(rawDesired.TransitEncryptionMode) {
		rawNew.TransitEncryptionMode = rawDesired.TransitEncryptionMode
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaintenancePolicy) && dcl.IsEmptyValueIndirect(rawDesired.MaintenancePolicy) {
		rawNew.MaintenancePolicy = rawDesired.MaintenancePolicy
	} else {
		rawNew.MaintenancePolicy = canonicalizeNewInstanceMaintenancePolicy(c, rawDesired.MaintenancePolicy, rawNew.MaintenancePolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaintenanceSchedule) && dcl.IsEmptyValueIndirect(rawDesired.MaintenanceSchedule) {
		rawNew.MaintenanceSchedule = rawDesired.MaintenanceSchedule
	} else {
		rawNew.MaintenanceSchedule = canonicalizeNewInstanceMaintenanceSchedule(c, rawDesired.MaintenanceSchedule, rawNew.MaintenanceSchedule)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.Region = rawDesired.Region

	return rawNew, nil
}

func canonicalizeInstanceServerCaCerts(des, initial *InstanceServerCaCerts, opts ...dcl.ApplyOption) *InstanceServerCaCerts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.SerialNumber, initial.SerialNumber) || dcl.IsZeroValue(des.SerialNumber) {
		des.SerialNumber = initial.SerialNumber
	}
	if dcl.StringCanonicalize(des.Cert, initial.Cert) || dcl.IsZeroValue(des.Cert) {
		des.Cert = initial.Cert
	}
	if dcl.IsZeroValue(des.CreateTime) {
		des.CreateTime = initial.CreateTime
	}
	if dcl.IsZeroValue(des.ExpireTime) {
		des.ExpireTime = initial.ExpireTime
	}
	if dcl.StringCanonicalize(des.Sha1Fingerprint, initial.Sha1Fingerprint) || dcl.IsZeroValue(des.Sha1Fingerprint) {
		des.Sha1Fingerprint = initial.Sha1Fingerprint
	}

	return des
}

func canonicalizeNewInstanceServerCaCerts(c *Client, des, nw *InstanceServerCaCerts) *InstanceServerCaCerts {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.SerialNumber, nw.SerialNumber) || dcl.IsZeroValue(des.SerialNumber) {
		nw.SerialNumber = des.SerialNumber
	}
	if dcl.StringCanonicalize(des.Cert, nw.Cert) || dcl.IsZeroValue(des.Cert) {
		nw.Cert = des.Cert
	}
	if dcl.StringCanonicalize(des.Sha1Fingerprint, nw.Sha1Fingerprint) || dcl.IsZeroValue(des.Sha1Fingerprint) {
		nw.Sha1Fingerprint = des.Sha1Fingerprint
	}

	return nw
}

func canonicalizeNewInstanceServerCaCertsSet(c *Client, des, nw []InstanceServerCaCerts) []InstanceServerCaCerts {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceServerCaCerts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceServerCaCerts(c, &d, &n) {
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

func canonicalizeNewInstanceServerCaCertsSlice(c *Client, des, nw []InstanceServerCaCerts) []InstanceServerCaCerts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceServerCaCerts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceServerCaCerts(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceMaintenancePolicy(des, initial *InstanceMaintenancePolicy, opts ...dcl.ApplyOption) *InstanceMaintenancePolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.CreateTime) {
		des.CreateTime = initial.CreateTime
	}
	if dcl.IsZeroValue(des.UpdateTime) {
		des.UpdateTime = initial.UpdateTime
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.WeeklyMaintenanceWindow) {
		des.WeeklyMaintenanceWindow = initial.WeeklyMaintenanceWindow
	}

	return des
}

func canonicalizeNewInstanceMaintenancePolicy(c *Client, des, nw *InstanceMaintenancePolicy) *InstanceMaintenancePolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Description, nw.Description) || dcl.IsZeroValue(des.Description) {
		nw.Description = des.Description
	}
	nw.WeeklyMaintenanceWindow = canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice(c, des.WeeklyMaintenanceWindow, nw.WeeklyMaintenanceWindow)

	return nw
}

func canonicalizeNewInstanceMaintenancePolicySet(c *Client, des, nw []InstanceMaintenancePolicy) []InstanceMaintenancePolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceMaintenancePolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceMaintenancePolicy(c, &d, &n) {
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

func canonicalizeNewInstanceMaintenancePolicySlice(c *Client, des, nw []InstanceMaintenancePolicy) []InstanceMaintenancePolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceMaintenancePolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceMaintenancePolicy(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceMaintenancePolicyWeeklyMaintenanceWindow(des, initial *InstanceMaintenancePolicyWeeklyMaintenanceWindow, opts ...dcl.ApplyOption) *InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Day) {
		des.Day = initial.Day
	}
	des.StartTime = canonicalizeInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(des.StartTime, initial.StartTime, opts...)
	if dcl.StringCanonicalize(des.Duration, initial.Duration) || dcl.IsZeroValue(des.Duration) {
		des.Duration = initial.Duration
	}

	return des
}

func canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindow(c *Client, des, nw *InstanceMaintenancePolicyWeeklyMaintenanceWindow) *InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	if des == nil || nw == nil {
		return nw
	}

	nw.StartTime = canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, des.StartTime, nw.StartTime)
	if dcl.StringCanonicalize(des.Duration, nw.Duration) || dcl.IsZeroValue(des.Duration) {
		nw.Duration = des.Duration
	}

	return nw
}

func canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowSet(c *Client, des, nw []InstanceMaintenancePolicyWeeklyMaintenanceWindow) []InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceMaintenancePolicyWeeklyMaintenanceWindow
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceMaintenancePolicyWeeklyMaintenanceWindow(c, &d, &n) {
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

func canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice(c *Client, des, nw []InstanceMaintenancePolicyWeeklyMaintenanceWindow) []InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceMaintenancePolicyWeeklyMaintenanceWindow
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindow(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(des, initial *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime, opts ...dcl.ApplyOption) *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Hours) {
		des.Hours = initial.Hours
	}
	if dcl.IsZeroValue(des.Minutes) {
		des.Minutes = initial.Minutes
	}
	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c *Client, des, nw *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSet(c *Client, des, nw []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, &d, &n) {
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

func canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSlice(c *Client, des, nw []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceMaintenanceSchedule(des, initial *InstanceMaintenanceSchedule, opts ...dcl.ApplyOption) *InstanceMaintenanceSchedule {
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
	if dcl.BoolCanonicalize(des.CanReschedule, initial.CanReschedule) || dcl.IsZeroValue(des.CanReschedule) {
		des.CanReschedule = initial.CanReschedule
	}

	return des
}

func canonicalizeNewInstanceMaintenanceSchedule(c *Client, des, nw *InstanceMaintenanceSchedule) *InstanceMaintenanceSchedule {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.CanReschedule, nw.CanReschedule) || dcl.IsZeroValue(des.CanReschedule) {
		nw.CanReschedule = des.CanReschedule
	}

	return nw
}

func canonicalizeNewInstanceMaintenanceScheduleSet(c *Client, des, nw []InstanceMaintenanceSchedule) []InstanceMaintenanceSchedule {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceMaintenanceSchedule
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceMaintenanceSchedule(c, &d, &n) {
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

func canonicalizeNewInstanceMaintenanceScheduleSlice(c *Client, des, nw []InstanceMaintenanceSchedule) []InstanceMaintenanceSchedule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceMaintenanceSchedule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceMaintenanceSchedule(c, &d, &n))
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
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "Name",
		})

	}
	if !dcl.IsZeroValue(desired.DisplayName) && !dcl.StringCanonicalize(desired.DisplayName, actual.DisplayName) {
		c.Config.Logger.Infof("Detected diff in DisplayName.\nDESIRED: %v\nACTUAL: %v", desired.DisplayName, actual.DisplayName)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "DisplayName",
		})

	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "Labels",
		})

	}
	if !dcl.IsZeroValue(desired.LocationId) && !dcl.StringCanonicalize(desired.LocationId, actual.LocationId) {
		c.Config.Logger.Infof("Detected diff in LocationId.\nDESIRED: %v\nACTUAL: %v", desired.LocationId, actual.LocationId)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "LocationId",
		})

	}
	if !dcl.IsZeroValue(desired.AlternativeLocationId) && !dcl.StringCanonicalize(desired.AlternativeLocationId, actual.AlternativeLocationId) {
		c.Config.Logger.Infof("Detected diff in AlternativeLocationId.\nDESIRED: %v\nACTUAL: %v", desired.AlternativeLocationId, actual.AlternativeLocationId)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "AlternativeLocationId",
		})

	}
	if !dcl.IsZeroValue(desired.RedisVersion) && !dcl.StringCanonicalize(desired.RedisVersion, actual.RedisVersion) {
		c.Config.Logger.Infof("Detected diff in RedisVersion.\nDESIRED: %v\nACTUAL: %v", desired.RedisVersion, actual.RedisVersion)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "RedisVersion",
		})

	}
	if !dcl.IsZeroValue(desired.ReservedIPRange) && !dcl.StringCanonicalize(desired.ReservedIPRange, actual.ReservedIPRange) {
		c.Config.Logger.Infof("Detected diff in ReservedIPRange.\nDESIRED: %v\nACTUAL: %v", desired.ReservedIPRange, actual.ReservedIPRange)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "ReservedIPRange",
		})

	}
	if !dcl.MapEquals(desired.RedisConfigs, actual.RedisConfigs, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in RedisConfigs.\nDESIRED: %v\nACTUAL: %v", desired.RedisConfigs, actual.RedisConfigs)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "RedisConfigs",
		})

	}
	if !reflect.DeepEqual(desired.Tier, actual.Tier) {
		c.Config.Logger.Infof("Detected diff in Tier.\nDESIRED: %v\nACTUAL: %v", desired.Tier, actual.Tier)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "Tier",
		})

	}
	if !reflect.DeepEqual(desired.MemorySizeGb, actual.MemorySizeGb) {
		c.Config.Logger.Infof("Detected diff in MemorySizeGb.\nDESIRED: %v\nACTUAL: %v", desired.MemorySizeGb, actual.MemorySizeGb)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "MemorySizeGb",
		})

	}
	if !dcl.IsZeroValue(desired.AuthorizedNetwork) && !dcl.StringCanonicalize(desired.AuthorizedNetwork, actual.AuthorizedNetwork) {
		c.Config.Logger.Infof("Detected diff in AuthorizedNetwork.\nDESIRED: %v\nACTUAL: %v", desired.AuthorizedNetwork, actual.AuthorizedNetwork)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "AuthorizedNetwork",
		})

	}
	if !reflect.DeepEqual(desired.ConnectMode, actual.ConnectMode) {
		c.Config.Logger.Infof("Detected diff in ConnectMode.\nDESIRED: %v\nACTUAL: %v", desired.ConnectMode, actual.ConnectMode)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "ConnectMode",
		})

	}
	if !dcl.IsZeroValue(desired.AuthEnabled) && !dcl.BoolCanonicalize(desired.AuthEnabled, actual.AuthEnabled) {
		c.Config.Logger.Infof("Detected diff in AuthEnabled.\nDESIRED: %v\nACTUAL: %v", desired.AuthEnabled, actual.AuthEnabled)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "AuthEnabled",
		})

	}
	if !reflect.DeepEqual(desired.TransitEncryptionMode, actual.TransitEncryptionMode) {
		c.Config.Logger.Infof("Detected diff in TransitEncryptionMode.\nDESIRED: %v\nACTUAL: %v", desired.TransitEncryptionMode, actual.TransitEncryptionMode)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "TransitEncryptionMode",
		})

	}
	if compareInstanceMaintenancePolicy(c, desired.MaintenancePolicy, actual.MaintenancePolicy) {
		c.Config.Logger.Infof("Detected diff in MaintenancePolicy.\nDESIRED: %v\nACTUAL: %v", desired.MaintenancePolicy, actual.MaintenancePolicy)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateInstanceOperation{},
			FieldName: "MaintenancePolicy",
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
func compareInstanceServerCaCerts(c *Client, desired, actual *InstanceServerCaCerts) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.SerialNumber == nil && desired.SerialNumber != nil && !dcl.IsEmptyValueIndirect(desired.SerialNumber) {
		c.Config.Logger.Infof("desired SerialNumber %s - but actually nil", dcl.SprintResource(desired.SerialNumber))
		return true
	}
	if !dcl.StringCanonicalize(desired.SerialNumber, actual.SerialNumber) && !dcl.IsZeroValue(desired.SerialNumber) {
		c.Config.Logger.Infof("Diff in SerialNumber. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SerialNumber), dcl.SprintResource(actual.SerialNumber))
		return true
	}
	if actual.Cert == nil && desired.Cert != nil && !dcl.IsEmptyValueIndirect(desired.Cert) {
		c.Config.Logger.Infof("desired Cert %s - but actually nil", dcl.SprintResource(desired.Cert))
		return true
	}
	if !dcl.StringCanonicalize(desired.Cert, actual.Cert) && !dcl.IsZeroValue(desired.Cert) {
		c.Config.Logger.Infof("Diff in Cert. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Cert), dcl.SprintResource(actual.Cert))
		return true
	}
	if actual.Sha1Fingerprint == nil && desired.Sha1Fingerprint != nil && !dcl.IsEmptyValueIndirect(desired.Sha1Fingerprint) {
		c.Config.Logger.Infof("desired Sha1Fingerprint %s - but actually nil", dcl.SprintResource(desired.Sha1Fingerprint))
		return true
	}
	if !dcl.StringCanonicalize(desired.Sha1Fingerprint, actual.Sha1Fingerprint) && !dcl.IsZeroValue(desired.Sha1Fingerprint) {
		c.Config.Logger.Infof("Diff in Sha1Fingerprint. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sha1Fingerprint), dcl.SprintResource(actual.Sha1Fingerprint))
		return true
	}
	return false
}

func compareInstanceServerCaCertsSlice(c *Client, desired, actual []InstanceServerCaCerts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceServerCaCerts, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceServerCaCerts(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceServerCaCerts, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceServerCaCertsMap(c *Client, desired, actual map[string]InstanceServerCaCerts) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceServerCaCerts, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceServerCaCerts, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceServerCaCerts(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceServerCaCerts, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceMaintenancePolicy(c *Client, desired, actual *InstanceMaintenancePolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !dcl.StringCanonicalize(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	if actual.WeeklyMaintenanceWindow == nil && desired.WeeklyMaintenanceWindow != nil && !dcl.IsEmptyValueIndirect(desired.WeeklyMaintenanceWindow) {
		c.Config.Logger.Infof("desired WeeklyMaintenanceWindow %s - but actually nil", dcl.SprintResource(desired.WeeklyMaintenanceWindow))
		return true
	}
	if compareInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice(c, desired.WeeklyMaintenanceWindow, actual.WeeklyMaintenanceWindow) && !dcl.IsZeroValue(desired.WeeklyMaintenanceWindow) {
		c.Config.Logger.Infof("Diff in WeeklyMaintenanceWindow. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WeeklyMaintenanceWindow), dcl.SprintResource(actual.WeeklyMaintenanceWindow))
		return true
	}
	return false
}

func compareInstanceMaintenancePolicySlice(c *Client, desired, actual []InstanceMaintenancePolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaintenancePolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceMaintenancePolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceMaintenancePolicyMap(c *Client, desired, actual map[string]InstanceMaintenancePolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaintenancePolicy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceMaintenancePolicy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicy, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindow(c *Client, desired, actual *InstanceMaintenancePolicyWeeklyMaintenanceWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Day == nil && desired.Day != nil && !dcl.IsEmptyValueIndirect(desired.Day) {
		c.Config.Logger.Infof("desired Day %s - but actually nil", dcl.SprintResource(desired.Day))
		return true
	}
	if !reflect.DeepEqual(desired.Day, actual.Day) && !dcl.IsZeroValue(desired.Day) {
		c.Config.Logger.Infof("Diff in Day. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Day), dcl.SprintResource(actual.Day))
		return true
	}
	if actual.StartTime == nil && desired.StartTime != nil && !dcl.IsEmptyValueIndirect(desired.StartTime) {
		c.Config.Logger.Infof("desired StartTime %s - but actually nil", dcl.SprintResource(desired.StartTime))
		return true
	}
	if compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) {
		c.Config.Logger.Infof("Diff in StartTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	return false
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice(c *Client, desired, actual []InstanceMaintenancePolicyWeeklyMaintenanceWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindow, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceMaintenancePolicyWeeklyMaintenanceWindow(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindow, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindowMap(c *Client, desired, actual map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindow, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindow, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceMaintenancePolicyWeeklyMaintenanceWindow(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindow, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c *Client, desired, actual *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Hours == nil && desired.Hours != nil && !dcl.IsEmptyValueIndirect(desired.Hours) {
		c.Config.Logger.Infof("desired Hours %s - but actually nil", dcl.SprintResource(desired.Hours))
		return true
	}
	if !reflect.DeepEqual(desired.Hours, actual.Hours) && !dcl.IsZeroValue(desired.Hours) {
		c.Config.Logger.Infof("Diff in Hours. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Hours), dcl.SprintResource(actual.Hours))
		return true
	}
	if actual.Minutes == nil && desired.Minutes != nil && !dcl.IsEmptyValueIndirect(desired.Minutes) {
		c.Config.Logger.Infof("desired Minutes %s - but actually nil", dcl.SprintResource(desired.Minutes))
		return true
	}
	if !reflect.DeepEqual(desired.Minutes, actual.Minutes) && !dcl.IsZeroValue(desired.Minutes) {
		c.Config.Logger.Infof("Diff in Minutes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Minutes), dcl.SprintResource(actual.Minutes))
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSlice(c *Client, desired, actual []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMap(c *Client, desired, actual map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceMaintenanceSchedule(c *Client, desired, actual *InstanceMaintenanceSchedule) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}

func compareInstanceMaintenanceScheduleSlice(c *Client, desired, actual []InstanceMaintenanceSchedule) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaintenanceSchedule, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceMaintenanceSchedule(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceMaintenanceSchedule, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceMaintenanceScheduleMap(c *Client, desired, actual map[string]InstanceMaintenanceSchedule) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaintenanceSchedule, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceMaintenanceSchedule, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceMaintenanceSchedule(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceMaintenanceSchedule, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceStateEnumSlice(c *Client, desired, actual []InstanceStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceStateEnum(c *Client, desired, actual *InstanceStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceTierEnumSlice(c *Client, desired, actual []InstanceTierEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTierEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTierEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTierEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTierEnum(c *Client, desired, actual *InstanceTierEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceConnectModeEnumSlice(c *Client, desired, actual []InstanceConnectModeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceConnectModeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceConnectModeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceConnectModeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceConnectModeEnum(c *Client, desired, actual *InstanceConnectModeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceTransitEncryptionModeEnumSlice(c *Client, desired, actual []InstanceTransitEncryptionModeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceTransitEncryptionModeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceTransitEncryptionModeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceTransitEncryptionModeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceTransitEncryptionModeEnum(c *Client, desired, actual *InstanceTransitEncryptionModeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumSlice(c *Client, desired, actual []InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(c *Client, desired, actual *InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Instance) urlNormalized() *Instance {
	normalized := deepcopy.Copy(*r).(Instance)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.LocationId = dcl.SelfLinkToName(r.LocationId)
	normalized.AlternativeLocationId = dcl.SelfLinkToName(r.AlternativeLocationId)
	normalized.RedisVersion = dcl.SelfLinkToName(r.RedisVersion)
	normalized.ReservedIPRange = dcl.SelfLinkToName(r.ReservedIPRange)
	normalized.Host = dcl.SelfLinkToName(r.Host)
	normalized.CurrentLocationId = dcl.SelfLinkToName(r.CurrentLocationId)
	normalized.StatusMessage = dcl.SelfLinkToName(r.StatusMessage)
	normalized.AuthorizedNetwork = dcl.SelfLinkToName(r.AuthorizedNetwork)
	normalized.PersistenceIamIdentity = dcl.SelfLinkToName(r.PersistenceIamIdentity)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	return &normalized
}

func (r *Instance) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateInstance" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", "https://redis.googleapis.com/v1/", userBasePath, fields), nil

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
	return unmarshalMapInstance(m, c)
}

func unmarshalMapInstance(m map[string]interface{}, c *Client) (*Instance, error) {

	return flattenInstance(c, m), nil
}

// expandInstance expands Instance into a JSON request object.
func expandInstance(c *Client, f *Instance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/instances/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.LocationId; !dcl.IsEmptyValueIndirect(v) {
		m["locationId"] = v
	}
	if v := f.AlternativeLocationId; !dcl.IsEmptyValueIndirect(v) {
		m["alternativeLocationId"] = v
	}
	if v := f.RedisVersion; !dcl.IsEmptyValueIndirect(v) {
		m["redisVersion"] = v
	}
	if v := f.ReservedIPRange; !dcl.IsEmptyValueIndirect(v) {
		m["reservedIpRange"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.CurrentLocationId; !dcl.IsEmptyValueIndirect(v) {
		m["currentLocationId"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.StatusMessage; !dcl.IsEmptyValueIndirect(v) {
		m["statusMessage"] = v
	}
	if v := f.RedisConfigs; !dcl.IsEmptyValueIndirect(v) {
		m["redisConfigs"] = v
	}
	if v := f.Tier; !dcl.IsEmptyValueIndirect(v) {
		m["tier"] = v
	}
	if v := f.MemorySizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["memorySizeGb"] = v
	}
	if v := f.AuthorizedNetwork; !dcl.IsEmptyValueIndirect(v) {
		m["authorizedNetwork"] = v
	}
	if v := f.PersistenceIamIdentity; !dcl.IsEmptyValueIndirect(v) {
		m["persistenceIamIdentity"] = v
	}
	if v := f.ConnectMode; !dcl.IsEmptyValueIndirect(v) {
		m["connectMode"] = v
	}
	if v := f.AuthEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["authEnabled"] = v
	}
	if v, err := expandInstanceServerCaCertsSlice(c, f.ServerCaCerts); err != nil {
		return nil, fmt.Errorf("error expanding ServerCaCerts into serverCaCerts: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["serverCaCerts"] = v
	}
	if v := f.TransitEncryptionMode; !dcl.IsEmptyValueIndirect(v) {
		m["transitEncryptionMode"] = v
	}
	if v, err := expandInstanceMaintenancePolicy(c, f.MaintenancePolicy); err != nil {
		return nil, fmt.Errorf("error expanding MaintenancePolicy into maintenancePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maintenancePolicy"] = v
	}
	if v, err := expandInstanceMaintenanceSchedule(c, f.MaintenanceSchedule); err != nil {
		return nil, fmt.Errorf("error expanding MaintenanceSchedule into maintenanceSchedule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maintenanceSchedule"] = v
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Region into region: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
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
	r.Name = dcl.FlattenString(m["name"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.LocationId = dcl.FlattenString(m["locationId"])
	r.AlternativeLocationId = dcl.FlattenString(m["alternativeLocationId"])
	r.RedisVersion = dcl.FlattenString(m["redisVersion"])
	r.ReservedIPRange = dcl.FlattenString(m["reservedIpRange"])
	r.Host = dcl.FlattenString(m["host"])
	r.Port = dcl.FlattenInteger(m["port"])
	r.CurrentLocationId = dcl.FlattenString(m["currentLocationId"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.State = flattenInstanceStateEnum(m["state"])
	r.StatusMessage = dcl.FlattenString(m["statusMessage"])
	r.RedisConfigs = dcl.FlattenKeyValuePairs(m["redisConfigs"])
	r.Tier = flattenInstanceTierEnum(m["tier"])
	r.MemorySizeGb = dcl.FlattenInteger(m["memorySizeGb"])
	r.AuthorizedNetwork = dcl.FlattenString(m["authorizedNetwork"])
	r.PersistenceIamIdentity = dcl.FlattenString(m["persistenceIamIdentity"])
	r.ConnectMode = flattenInstanceConnectModeEnum(m["connectMode"])
	r.AuthEnabled = dcl.FlattenBool(m["authEnabled"])
	r.ServerCaCerts = flattenInstanceServerCaCertsSlice(c, m["serverCaCerts"])
	r.TransitEncryptionMode = flattenInstanceTransitEncryptionModeEnum(m["transitEncryptionMode"])
	r.MaintenancePolicy = flattenInstanceMaintenancePolicy(c, m["maintenancePolicy"])
	r.MaintenanceSchedule = flattenInstanceMaintenanceSchedule(c, m["maintenanceSchedule"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])
	r.Region = dcl.FlattenString(m["region"])

	return r
}

// expandInstanceServerCaCertsMap expands the contents of InstanceServerCaCerts into a JSON
// request object.
func expandInstanceServerCaCertsMap(c *Client, f map[string]InstanceServerCaCerts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceServerCaCerts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceServerCaCertsSlice expands the contents of InstanceServerCaCerts into a JSON
// request object.
func expandInstanceServerCaCertsSlice(c *Client, f []InstanceServerCaCerts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceServerCaCerts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceServerCaCertsMap flattens the contents of InstanceServerCaCerts from a JSON
// response object.
func flattenInstanceServerCaCertsMap(c *Client, i interface{}) map[string]InstanceServerCaCerts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceServerCaCerts{}
	}

	if len(a) == 0 {
		return map[string]InstanceServerCaCerts{}
	}

	items := make(map[string]InstanceServerCaCerts)
	for k, item := range a {
		items[k] = *flattenInstanceServerCaCerts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceServerCaCertsSlice flattens the contents of InstanceServerCaCerts from a JSON
// response object.
func flattenInstanceServerCaCertsSlice(c *Client, i interface{}) []InstanceServerCaCerts {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceServerCaCerts{}
	}

	if len(a) == 0 {
		return []InstanceServerCaCerts{}
	}

	items := make([]InstanceServerCaCerts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceServerCaCerts(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceServerCaCerts expands an instance of InstanceServerCaCerts into a JSON
// request object.
func expandInstanceServerCaCerts(c *Client, f *InstanceServerCaCerts) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SerialNumber; !dcl.IsEmptyValueIndirect(v) {
		m["serialNumber"] = v
	}
	if v := f.Cert; !dcl.IsEmptyValueIndirect(v) {
		m["cert"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.ExpireTime; !dcl.IsEmptyValueIndirect(v) {
		m["expireTime"] = v
	}
	if v := f.Sha1Fingerprint; !dcl.IsEmptyValueIndirect(v) {
		m["sha1Fingerprint"] = v
	}

	return m, nil
}

// flattenInstanceServerCaCerts flattens an instance of InstanceServerCaCerts from a JSON
// response object.
func flattenInstanceServerCaCerts(c *Client, i interface{}) *InstanceServerCaCerts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceServerCaCerts{}
	r.SerialNumber = dcl.FlattenString(m["serialNumber"])
	r.Cert = dcl.FlattenString(m["cert"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.ExpireTime = dcl.FlattenString(m["expireTime"])
	r.Sha1Fingerprint = dcl.FlattenString(m["sha1Fingerprint"])

	return r
}

// expandInstanceMaintenancePolicyMap expands the contents of InstanceMaintenancePolicy into a JSON
// request object.
func expandInstanceMaintenancePolicyMap(c *Client, f map[string]InstanceMaintenancePolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceMaintenancePolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceMaintenancePolicySlice expands the contents of InstanceMaintenancePolicy into a JSON
// request object.
func expandInstanceMaintenancePolicySlice(c *Client, f []InstanceMaintenancePolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceMaintenancePolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceMaintenancePolicyMap flattens the contents of InstanceMaintenancePolicy from a JSON
// response object.
func flattenInstanceMaintenancePolicyMap(c *Client, i interface{}) map[string]InstanceMaintenancePolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceMaintenancePolicy{}
	}

	if len(a) == 0 {
		return map[string]InstanceMaintenancePolicy{}
	}

	items := make(map[string]InstanceMaintenancePolicy)
	for k, item := range a {
		items[k] = *flattenInstanceMaintenancePolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceMaintenancePolicySlice flattens the contents of InstanceMaintenancePolicy from a JSON
// response object.
func flattenInstanceMaintenancePolicySlice(c *Client, i interface{}) []InstanceMaintenancePolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceMaintenancePolicy{}
	}

	if len(a) == 0 {
		return []InstanceMaintenancePolicy{}
	}

	items := make([]InstanceMaintenancePolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceMaintenancePolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceMaintenancePolicy expands an instance of InstanceMaintenancePolicy into a JSON
// request object.
func expandInstanceMaintenancePolicy(c *Client, f *InstanceMaintenancePolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice(c, f.WeeklyMaintenanceWindow); err != nil {
		return nil, fmt.Errorf("error expanding WeeklyMaintenanceWindow into weeklyMaintenanceWindow: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["weeklyMaintenanceWindow"] = v
	}

	return m, nil
}

// flattenInstanceMaintenancePolicy flattens an instance of InstanceMaintenancePolicy from a JSON
// response object.
func flattenInstanceMaintenancePolicy(c *Client, i interface{}) *InstanceMaintenancePolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceMaintenancePolicy{}
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.Description = dcl.FlattenString(m["description"])
	r.WeeklyMaintenanceWindow = flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice(c, m["weeklyMaintenanceWindow"])

	return r
}

// expandInstanceMaintenancePolicyWeeklyMaintenanceWindowMap expands the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindow into a JSON
// request object.
func expandInstanceMaintenancePolicyWeeklyMaintenanceWindowMap(c *Client, f map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindow) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceMaintenancePolicyWeeklyMaintenanceWindow(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice expands the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindow into a JSON
// request object.
func expandInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice(c *Client, f []InstanceMaintenancePolicyWeeklyMaintenanceWindow) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceMaintenancePolicyWeeklyMaintenanceWindow(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowMap flattens the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindow from a JSON
// response object.
func flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowMap(c *Client, i interface{}) map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindow{}
	}

	if len(a) == 0 {
		return map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindow{}
	}

	items := make(map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindow)
	for k, item := range a {
		items[k] = *flattenInstanceMaintenancePolicyWeeklyMaintenanceWindow(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice flattens the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindow from a JSON
// response object.
func flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowSlice(c *Client, i interface{}) []InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceMaintenancePolicyWeeklyMaintenanceWindow{}
	}

	if len(a) == 0 {
		return []InstanceMaintenancePolicyWeeklyMaintenanceWindow{}
	}

	items := make([]InstanceMaintenancePolicyWeeklyMaintenanceWindow, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceMaintenancePolicyWeeklyMaintenanceWindow(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceMaintenancePolicyWeeklyMaintenanceWindow expands an instance of InstanceMaintenancePolicyWeeklyMaintenanceWindow into a JSON
// request object.
func expandInstanceMaintenancePolicyWeeklyMaintenanceWindow(c *Client, f *InstanceMaintenancePolicyWeeklyMaintenanceWindow) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Day; !dcl.IsEmptyValueIndirect(v) {
		m["day"] = v
	}
	if v, err := expandInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, f.StartTime); err != nil {
		return nil, fmt.Errorf("error expanding StartTime into startTime: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.Duration; !dcl.IsEmptyValueIndirect(v) {
		m["duration"] = v
	}

	return m, nil
}

// flattenInstanceMaintenancePolicyWeeklyMaintenanceWindow flattens an instance of InstanceMaintenancePolicyWeeklyMaintenanceWindow from a JSON
// response object.
func flattenInstanceMaintenancePolicyWeeklyMaintenanceWindow(c *Client, i interface{}) *InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceMaintenancePolicyWeeklyMaintenanceWindow{}
	r.Day = flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(m["day"])
	r.StartTime = flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, m["startTime"])
	r.Duration = dcl.FlattenString(m["duration"])

	return r
}

// expandInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMap expands the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime into a JSON
// request object.
func expandInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMap(c *Client, f map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSlice expands the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime into a JSON
// request object.
func expandInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSlice(c *Client, f []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMap flattens the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime from a JSON
// response object.
func flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMap(c *Client, i interface{}) map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{}
	}

	if len(a) == 0 {
		return map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{}
	}

	items := make(map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime)
	for k, item := range a {
		items[k] = *flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSlice flattens the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime from a JSON
// response object.
func flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSlice(c *Client, i interface{}) []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{}
	}

	if len(a) == 0 {
		return []InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{}
	}

	items := make([]InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime expands an instance of InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime into a JSON
// request object.
func expandInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c *Client, f *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Hours; !dcl.IsEmptyValueIndirect(v) {
		m["hours"] = v
	}
	if v := f.Minutes; !dcl.IsEmptyValueIndirect(v) {
		m["minutes"] = v
	}
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime flattens an instance of InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime from a JSON
// response object.
func flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c *Client, i interface{}) *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{}
	r.Hours = dcl.FlattenInteger(m["hours"])
	r.Minutes = dcl.FlattenInteger(m["minutes"])
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandInstanceMaintenanceScheduleMap expands the contents of InstanceMaintenanceSchedule into a JSON
// request object.
func expandInstanceMaintenanceScheduleMap(c *Client, f map[string]InstanceMaintenanceSchedule) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceMaintenanceSchedule(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceMaintenanceScheduleSlice expands the contents of InstanceMaintenanceSchedule into a JSON
// request object.
func expandInstanceMaintenanceScheduleSlice(c *Client, f []InstanceMaintenanceSchedule) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceMaintenanceSchedule(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceMaintenanceScheduleMap flattens the contents of InstanceMaintenanceSchedule from a JSON
// response object.
func flattenInstanceMaintenanceScheduleMap(c *Client, i interface{}) map[string]InstanceMaintenanceSchedule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceMaintenanceSchedule{}
	}

	if len(a) == 0 {
		return map[string]InstanceMaintenanceSchedule{}
	}

	items := make(map[string]InstanceMaintenanceSchedule)
	for k, item := range a {
		items[k] = *flattenInstanceMaintenanceSchedule(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceMaintenanceScheduleSlice flattens the contents of InstanceMaintenanceSchedule from a JSON
// response object.
func flattenInstanceMaintenanceScheduleSlice(c *Client, i interface{}) []InstanceMaintenanceSchedule {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceMaintenanceSchedule{}
	}

	if len(a) == 0 {
		return []InstanceMaintenanceSchedule{}
	}

	items := make([]InstanceMaintenanceSchedule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceMaintenanceSchedule(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceMaintenanceSchedule expands an instance of InstanceMaintenanceSchedule into a JSON
// request object.
func expandInstanceMaintenanceSchedule(c *Client, f *InstanceMaintenanceSchedule) (map[string]interface{}, error) {
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
	if v := f.CanReschedule; !dcl.IsEmptyValueIndirect(v) {
		m["canReschedule"] = v
	}

	return m, nil
}

// flattenInstanceMaintenanceSchedule flattens an instance of InstanceMaintenanceSchedule from a JSON
// response object.
func flattenInstanceMaintenanceSchedule(c *Client, i interface{}) *InstanceMaintenanceSchedule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceMaintenanceSchedule{}
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.EndTime = dcl.FlattenString(m["endTime"])
	r.CanReschedule = dcl.FlattenBool(m["canReschedule"])

	return r
}

// flattenInstanceStateEnumSlice flattens the contents of InstanceStateEnum from a JSON
// response object.
func flattenInstanceStateEnumSlice(c *Client, i interface{}) []InstanceStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceStateEnum{}
	}

	if len(a) == 0 {
		return []InstanceStateEnum{}
	}

	items := make([]InstanceStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceStateEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceStateEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceStateEnum with the same value as that string.
func flattenInstanceStateEnum(i interface{}) *InstanceStateEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceStateEnumRef("")
	}

	return InstanceStateEnumRef(s)
}

// flattenInstanceTierEnumSlice flattens the contents of InstanceTierEnum from a JSON
// response object.
func flattenInstanceTierEnumSlice(c *Client, i interface{}) []InstanceTierEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTierEnum{}
	}

	if len(a) == 0 {
		return []InstanceTierEnum{}
	}

	items := make([]InstanceTierEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTierEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceTierEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTierEnum with the same value as that string.
func flattenInstanceTierEnum(i interface{}) *InstanceTierEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTierEnumRef("")
	}

	return InstanceTierEnumRef(s)
}

// flattenInstanceConnectModeEnumSlice flattens the contents of InstanceConnectModeEnum from a JSON
// response object.
func flattenInstanceConnectModeEnumSlice(c *Client, i interface{}) []InstanceConnectModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceConnectModeEnum{}
	}

	if len(a) == 0 {
		return []InstanceConnectModeEnum{}
	}

	items := make([]InstanceConnectModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceConnectModeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceConnectModeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceConnectModeEnum with the same value as that string.
func flattenInstanceConnectModeEnum(i interface{}) *InstanceConnectModeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceConnectModeEnumRef("")
	}

	return InstanceConnectModeEnumRef(s)
}

// flattenInstanceTransitEncryptionModeEnumSlice flattens the contents of InstanceTransitEncryptionModeEnum from a JSON
// response object.
func flattenInstanceTransitEncryptionModeEnumSlice(c *Client, i interface{}) []InstanceTransitEncryptionModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTransitEncryptionModeEnum{}
	}

	if len(a) == 0 {
		return []InstanceTransitEncryptionModeEnum{}
	}

	items := make([]InstanceTransitEncryptionModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTransitEncryptionModeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceTransitEncryptionModeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTransitEncryptionModeEnum with the same value as that string.
func flattenInstanceTransitEncryptionModeEnum(i interface{}) *InstanceTransitEncryptionModeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTransitEncryptionModeEnumRef("")
	}

	return InstanceTransitEncryptionModeEnumRef(s)
}

// flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumSlice flattens the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum from a JSON
// response object.
func flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumSlice(c *Client, i interface{}) []InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum{}
	}

	if len(a) == 0 {
		return []InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum{}
	}

	items := make([]InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum with the same value as that string.
func flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(i interface{}) *InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumRef("")
	}

	return InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumRef(s)
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
