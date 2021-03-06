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
	"strings"
	"time"

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

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.RedisConfigs; !dcl.IsEmptyValueIndirect(v) {
		req["redisConfigs"] = v
	}
	if v := f.MemorySizeGb; !dcl.IsEmptyValueIndirect(v) {
		req["memorySizeGb"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/instances/%s", *f.Project, *f.Location, *f.Name)

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
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceUpdateInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateInstance")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
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
		res, err := unmarshalMapInstance(v, c)
		if err != nil {
			return nil, m.Token, err
		}
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
	r, err := c.GetInstance(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Instance not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetInstance checking for existence. error: %v", err)
		return err
	}

	u, err := instanceDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetInstance(ctx, r.URLNormalized())
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

	if _, err := c.GetInstance(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getInstanceRaw(ctx context.Context, r *Instance) ([]byte, error) {

	u, err := instanceGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) instanceDiffsForRawDesired(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (initial, desired *Instance, diffs []*dcl.FieldDiff, err error) {
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
	rawInitial, err := c.GetInstance(ctx, fetchState.URLNormalized())
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
	canonicalDesired := &Instance{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.StringCanonicalize(rawDesired.LocationId, rawInitial.LocationId) {
		canonicalDesired.LocationId = rawInitial.LocationId
	} else {
		canonicalDesired.LocationId = rawDesired.LocationId
	}
	if dcl.StringCanonicalize(rawDesired.AlternativeLocationId, rawInitial.AlternativeLocationId) {
		canonicalDesired.AlternativeLocationId = rawInitial.AlternativeLocationId
	} else {
		canonicalDesired.AlternativeLocationId = rawDesired.AlternativeLocationId
	}
	if dcl.StringCanonicalize(rawDesired.RedisVersion, rawInitial.RedisVersion) {
		canonicalDesired.RedisVersion = rawInitial.RedisVersion
	} else {
		canonicalDesired.RedisVersion = rawDesired.RedisVersion
	}
	if dcl.StringCanonicalize(rawDesired.ReservedIPRange, rawInitial.ReservedIPRange) {
		canonicalDesired.ReservedIPRange = rawInitial.ReservedIPRange
	} else {
		canonicalDesired.ReservedIPRange = rawDesired.ReservedIPRange
	}
	if dcl.IsZeroValue(rawDesired.RedisConfigs) {
		canonicalDesired.RedisConfigs = rawInitial.RedisConfigs
	} else {
		canonicalDesired.RedisConfigs = rawDesired.RedisConfigs
	}
	if dcl.IsZeroValue(rawDesired.Tier) {
		canonicalDesired.Tier = rawInitial.Tier
	} else {
		canonicalDesired.Tier = rawDesired.Tier
	}
	if dcl.IsZeroValue(rawDesired.MemorySizeGb) {
		canonicalDesired.MemorySizeGb = rawInitial.MemorySizeGb
	} else {
		canonicalDesired.MemorySizeGb = rawDesired.MemorySizeGb
	}
	if dcl.StringCanonicalize(rawDesired.AuthorizedNetwork, rawInitial.AuthorizedNetwork) {
		canonicalDesired.AuthorizedNetwork = rawInitial.AuthorizedNetwork
	} else {
		canonicalDesired.AuthorizedNetwork = rawDesired.AuthorizedNetwork
	}
	if dcl.IsZeroValue(rawDesired.ConnectMode) {
		canonicalDesired.ConnectMode = rawInitial.ConnectMode
	} else {
		canonicalDesired.ConnectMode = rawDesired.ConnectMode
	}
	if dcl.BoolCanonicalize(rawDesired.AuthEnabled, rawInitial.AuthEnabled) {
		canonicalDesired.AuthEnabled = rawInitial.AuthEnabled
	} else {
		canonicalDesired.AuthEnabled = rawDesired.AuthEnabled
	}
	if dcl.IsZeroValue(rawDesired.TransitEncryptionMode) {
		canonicalDesired.TransitEncryptionMode = rawInitial.TransitEncryptionMode
	} else {
		canonicalDesired.TransitEncryptionMode = rawDesired.TransitEncryptionMode
	}
	canonicalDesired.MaintenancePolicy = canonicalizeInstanceMaintenancePolicy(rawDesired.MaintenancePolicy, rawInitial.MaintenancePolicy, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeInstanceNewState(c *Client, rawNew, rawDesired *Instance) (*Instance, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
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

	cDes := &InstanceServerCaCerts{}

	if dcl.StringCanonicalize(des.SerialNumber, initial.SerialNumber) || dcl.IsZeroValue(des.SerialNumber) {
		cDes.SerialNumber = initial.SerialNumber
	} else {
		cDes.SerialNumber = des.SerialNumber
	}
	if dcl.StringCanonicalize(des.Cert, initial.Cert) || dcl.IsZeroValue(des.Cert) {
		cDes.Cert = initial.Cert
	} else {
		cDes.Cert = des.Cert
	}
	if dcl.StringCanonicalize(des.Sha1Fingerprint, initial.Sha1Fingerprint) || dcl.IsZeroValue(des.Sha1Fingerprint) {
		cDes.Sha1Fingerprint = initial.Sha1Fingerprint
	} else {
		cDes.Sha1Fingerprint = des.Sha1Fingerprint
	}

	return cDes
}

func canonicalizeNewInstanceServerCaCerts(c *Client, des, nw *InstanceServerCaCerts) *InstanceServerCaCerts {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.SerialNumber, nw.SerialNumber) {
		nw.SerialNumber = des.SerialNumber
	}
	if dcl.StringCanonicalize(des.Cert, nw.Cert) {
		nw.Cert = des.Cert
	}
	if dcl.IsZeroValue(nw.CreateTime) {
		nw.CreateTime = des.CreateTime
	}
	if dcl.IsZeroValue(nw.ExpireTime) {
		nw.ExpireTime = des.ExpireTime
	}
	if dcl.StringCanonicalize(des.Sha1Fingerprint, nw.Sha1Fingerprint) {
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
			if diffs, _ := compareInstanceServerCaCertsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &InstanceMaintenancePolicy{}

	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}
	if dcl.IsZeroValue(des.WeeklyMaintenanceWindow) {
		des.WeeklyMaintenanceWindow = initial.WeeklyMaintenanceWindow
	} else {
		cDes.WeeklyMaintenanceWindow = des.WeeklyMaintenanceWindow
	}

	return cDes
}

func canonicalizeNewInstanceMaintenancePolicy(c *Client, des, nw *InstanceMaintenancePolicy) *InstanceMaintenancePolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.CreateTime) {
		nw.CreateTime = des.CreateTime
	}
	if dcl.IsZeroValue(nw.UpdateTime) {
		nw.UpdateTime = des.UpdateTime
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
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
			if diffs, _ := compareInstanceMaintenancePolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &InstanceMaintenancePolicyWeeklyMaintenanceWindow{}

	if dcl.IsZeroValue(des.Day) {
		des.Day = initial.Day
	} else {
		cDes.Day = des.Day
	}
	cDes.StartTime = canonicalizeInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(des.StartTime, initial.StartTime, opts...)

	return cDes
}

func canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindow(c *Client, des, nw *InstanceMaintenancePolicyWeeklyMaintenanceWindow) *InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Day) {
		nw.Day = des.Day
	}
	nw.StartTime = canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c, des.StartTime, nw.StartTime)
	if dcl.StringCanonicalize(des.Duration, nw.Duration) {
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
			if diffs, _ := compareInstanceMaintenancePolicyWeeklyMaintenanceWindowNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{}

	if dcl.IsZeroValue(des.Hours) {
		des.Hours = initial.Hours
	} else {
		cDes.Hours = des.Hours
	}
	if dcl.IsZeroValue(des.Minutes) {
		des.Minutes = initial.Minutes
	} else {
		cDes.Minutes = des.Minutes
	}
	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	} else {
		cDes.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	} else {
		cDes.Nanos = des.Nanos
	}

	return cDes
}

func canonicalizeNewInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(c *Client, des, nw *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Hours) {
		nw.Hours = des.Hours
	}
	if dcl.IsZeroValue(nw.Minutes) {
		nw.Minutes = des.Minutes
	}
	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
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
			if diffs, _ := compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	cDes := &InstanceMaintenanceSchedule{}

	if dcl.BoolCanonicalize(des.CanReschedule, initial.CanReschedule) || dcl.IsZeroValue(des.CanReschedule) {
		cDes.CanReschedule = initial.CanReschedule
	} else {
		cDes.CanReschedule = des.CanReschedule
	}

	return cDes
}

func canonicalizeNewInstanceMaintenanceSchedule(c *Client, des, nw *InstanceMaintenanceSchedule) *InstanceMaintenanceSchedule {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.StartTime) {
		nw.StartTime = des.StartTime
	}
	if dcl.IsZeroValue(nw.EndTime) {
		nw.EndTime = des.EndTime
	}
	if dcl.BoolCanonicalize(des.CanReschedule, nw.CanReschedule) {
		nw.CanReschedule = des.CanReschedule
	}
	if dcl.IsZeroValue(nw.ScheduleDeadlineTime) {
		nw.ScheduleDeadlineTime = des.ScheduleDeadlineTime
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
			if diffs, _ := compareInstanceMaintenanceScheduleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
	}

	var items []InstanceMaintenanceSchedule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceMaintenanceSchedule(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffInstance(c *Client, desired, actual *Instance, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocationId, actual.LocationId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocationId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AlternativeLocationId, actual.AlternativeLocationId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AlternativeLocationId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RedisVersion, actual.RedisVersion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RedisVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReservedIPRange, actual.ReservedIPRange, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReservedIpRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CurrentLocationId, actual.CurrentLocationId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CurrentLocationId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StatusMessage, actual.StatusMessage, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StatusMessage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RedisConfigs, actual.RedisConfigs, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("RedisConfigs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tier, actual.Tier, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MemorySizeGb, actual.MemorySizeGb, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("MemorySizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthorizedNetwork, actual.AuthorizedNetwork, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuthorizedNetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PersistenceIamIdentity, actual.PersistenceIamIdentity, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PersistenceIamIdentity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConnectMode, actual.ConnectMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConnectMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthEnabled, actual.AuthEnabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuthEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServerCaCerts, actual.ServerCaCerts, dcl.Info{OutputOnly: true, ObjectFunction: compareInstanceServerCaCertsNewStyle, EmptyObject: EmptyInstanceServerCaCerts, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServerCaCerts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TransitEncryptionMode, actual.TransitEncryptionMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TransitEncryptionMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaintenancePolicy, actual.MaintenancePolicy, dcl.Info{ObjectFunction: compareInstanceMaintenancePolicyNewStyle, EmptyObject: EmptyInstanceMaintenancePolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaintenancePolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaintenanceSchedule, actual.MaintenanceSchedule, dcl.Info{OutputOnly: true, ObjectFunction: compareInstanceMaintenanceScheduleNewStyle, EmptyObject: EmptyInstanceMaintenanceSchedule, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaintenanceSchedule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareInstanceServerCaCertsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceServerCaCerts)
	if !ok {
		desiredNotPointer, ok := d.(InstanceServerCaCerts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceServerCaCerts or *InstanceServerCaCerts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceServerCaCerts)
	if !ok {
		actualNotPointer, ok := a.(InstanceServerCaCerts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceServerCaCerts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SerialNumber, actual.SerialNumber, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SerialNumber")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Cert, actual.Cert, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Cert")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpireTime, actual.ExpireTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpireTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha1Fingerprint, actual.Sha1Fingerprint, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha1Fingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceMaintenancePolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceMaintenancePolicy)
	if !ok {
		desiredNotPointer, ok := d.(InstanceMaintenancePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceMaintenancePolicy or *InstanceMaintenancePolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceMaintenancePolicy)
	if !ok {
		actualNotPointer, ok := a.(InstanceMaintenancePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceMaintenancePolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.WeeklyMaintenanceWindow, actual.WeeklyMaintenanceWindow, dcl.Info{ObjectFunction: compareInstanceMaintenancePolicyWeeklyMaintenanceWindowNewStyle, EmptyObject: EmptyInstanceMaintenancePolicyWeeklyMaintenanceWindow, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WeeklyMaintenanceWindow")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindowNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceMaintenancePolicyWeeklyMaintenanceWindow)
	if !ok {
		desiredNotPointer, ok := d.(InstanceMaintenancePolicyWeeklyMaintenanceWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceMaintenancePolicyWeeklyMaintenanceWindow or *InstanceMaintenancePolicyWeeklyMaintenanceWindow", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceMaintenancePolicyWeeklyMaintenanceWindow)
	if !ok {
		actualNotPointer, ok := a.(InstanceMaintenancePolicyWeeklyMaintenanceWindow)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceMaintenancePolicyWeeklyMaintenanceWindow", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Day, actual.Day, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Day")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StartTime, actual.StartTime, dcl.Info{ObjectFunction: compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeNewStyle, EmptyObject: EmptyInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Duration, actual.Duration, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Duration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime)
	if !ok {
		desiredNotPointer, ok := d.(InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime or *InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime)
	if !ok {
		actualNotPointer, ok := a.(InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Hours, actual.Hours, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Hours")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Minutes, actual.Minutes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Minutes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceMaintenanceScheduleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceMaintenanceSchedule)
	if !ok {
		desiredNotPointer, ok := d.(InstanceMaintenanceSchedule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceMaintenanceSchedule or *InstanceMaintenanceSchedule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceMaintenanceSchedule)
	if !ok {
		actualNotPointer, ok := a.(InstanceMaintenanceSchedule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceMaintenanceSchedule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StartTime, actual.StartTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndTime, actual.EndTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EndTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CanReschedule, actual.CanReschedule, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CanReschedule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ScheduleDeadlineTime, actual.ScheduleDeadlineTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ScheduleDeadlineTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Instance) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) createFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
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
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
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
	} else if v != nil {
		m["serverCaCerts"] = v
	}
	if v := f.TransitEncryptionMode; !dcl.IsEmptyValueIndirect(v) {
		m["transitEncryptionMode"] = v
	}
	if v, err := expandInstanceMaintenancePolicy(c, f.MaintenancePolicy); err != nil {
		return nil, fmt.Errorf("error expanding MaintenancePolicy into maintenancePolicy: %w", err)
	} else if v != nil {
		m["maintenancePolicy"] = v
	}
	if v, err := expandInstanceMaintenanceSchedule(c, f.MaintenanceSchedule); err != nil {
		return nil, fmt.Errorf("error expanding MaintenanceSchedule into maintenanceSchedule: %w", err)
	} else if v != nil {
		m["maintenanceSchedule"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
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

	res := &Instance{}
	res.Name = dcl.FlattenString(m["name"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.LocationId = dcl.FlattenString(m["locationId"])
	res.AlternativeLocationId = dcl.FlattenString(m["alternativeLocationId"])
	res.RedisVersion = dcl.FlattenString(m["redisVersion"])
	res.ReservedIPRange = dcl.FlattenString(m["reservedIpRange"])
	res.Host = dcl.FlattenString(m["host"])
	res.Port = dcl.FlattenInteger(m["port"])
	res.CurrentLocationId = dcl.FlattenString(m["currentLocationId"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.State = flattenInstanceStateEnum(m["state"])
	res.StatusMessage = dcl.FlattenString(m["statusMessage"])
	res.RedisConfigs = dcl.FlattenKeyValuePairs(m["redisConfigs"])
	res.Tier = flattenInstanceTierEnum(m["tier"])
	res.MemorySizeGb = dcl.FlattenInteger(m["memorySizeGb"])
	res.AuthorizedNetwork = dcl.FlattenString(m["authorizedNetwork"])
	res.PersistenceIamIdentity = dcl.FlattenString(m["persistenceIamIdentity"])
	res.ConnectMode = flattenInstanceConnectModeEnum(m["connectMode"])
	res.AuthEnabled = dcl.FlattenBool(m["authEnabled"])
	res.ServerCaCerts = flattenInstanceServerCaCertsSlice(c, m["serverCaCerts"])
	res.TransitEncryptionMode = flattenInstanceTransitEncryptionModeEnum(m["transitEncryptionMode"])
	res.MaintenancePolicy = flattenInstanceMaintenancePolicy(c, m["maintenancePolicy"])
	res.MaintenanceSchedule = flattenInstanceMaintenanceSchedule(c, m["maintenanceSchedule"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceServerCaCerts
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceMaintenancePolicy
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceMaintenancePolicyWeeklyMaintenanceWindow
	}
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime
	}
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
	if v := f.ScheduleDeadlineTime; !dcl.IsEmptyValueIndirect(v) {
		m["scheduleDeadlineTime"] = v
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceMaintenanceSchedule
	}
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.EndTime = dcl.FlattenString(m["endTime"])
	r.CanReschedule = dcl.FlattenBool(m["canReschedule"])
	r.ScheduleDeadlineTime = dcl.FlattenString(m["scheduleDeadlineTime"])

	return r
}

// flattenInstanceStateEnumMap flattens the contents of InstanceStateEnum from a JSON
// response object.
func flattenInstanceStateEnumMap(c *Client, i interface{}) map[string]InstanceStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceStateEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceStateEnum{}
	}

	items := make(map[string]InstanceStateEnum)
	for k, item := range a {
		items[k] = *flattenInstanceStateEnum(item.(interface{}))
	}

	return items
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

// flattenInstanceTierEnumMap flattens the contents of InstanceTierEnum from a JSON
// response object.
func flattenInstanceTierEnumMap(c *Client, i interface{}) map[string]InstanceTierEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTierEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceTierEnum{}
	}

	items := make(map[string]InstanceTierEnum)
	for k, item := range a {
		items[k] = *flattenInstanceTierEnum(item.(interface{}))
	}

	return items
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

// flattenInstanceConnectModeEnumMap flattens the contents of InstanceConnectModeEnum from a JSON
// response object.
func flattenInstanceConnectModeEnumMap(c *Client, i interface{}) map[string]InstanceConnectModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceConnectModeEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceConnectModeEnum{}
	}

	items := make(map[string]InstanceConnectModeEnum)
	for k, item := range a {
		items[k] = *flattenInstanceConnectModeEnum(item.(interface{}))
	}

	return items
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

// flattenInstanceTransitEncryptionModeEnumMap flattens the contents of InstanceTransitEncryptionModeEnum from a JSON
// response object.
func flattenInstanceTransitEncryptionModeEnumMap(c *Client, i interface{}) map[string]InstanceTransitEncryptionModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTransitEncryptionModeEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceTransitEncryptionModeEnum{}
	}

	items := make(map[string]InstanceTransitEncryptionModeEnum)
	for k, item := range a {
		items[k] = *flattenInstanceTransitEncryptionModeEnum(item.(interface{}))
	}

	return items
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

// flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumMap flattens the contents of InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum from a JSON
// response object.
func flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumMap(c *Client, i interface{}) map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum{}
	}

	items := make(map[string]InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum)
	for k, item := range a {
		items[k] = *flattenInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(item.(interface{}))
	}

	return items
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
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
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

type instanceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceApiOperation
}

func convertFieldDiffsToInstanceDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]instanceDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []instanceDiff
	// For each operation name, create a instanceDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := instanceDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToInstanceApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToInstanceApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (instanceApiOperation, error) {
	switch opName {

	case "updateInstanceUpdateInstanceOperation":
		return &updateInstanceUpdateInstanceOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
