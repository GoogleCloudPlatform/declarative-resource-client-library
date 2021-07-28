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

func (r *Instance) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.NetworkConfig) {
		if err := r.NetworkConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceNetworkConfig) validate() error {
	return nil
}
func (r *InstanceAvailableVersion) validate() error {
	return nil
}

func instanceGetURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", "https://datafusion.googleapis.com/v1beta1/", userBasePath, params), nil
}

func instanceListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances", "https://datafusion.googleapis.com/v1beta1/", userBasePath, params), nil

}

func instanceCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances?instanceId={{name}}", "https://datafusion.googleapis.com/v1beta1/", userBasePath, params), nil

}

func instanceDeleteURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", "https://datafusion.googleapis.com/v1beta1/", userBasePath, params), nil
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

	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		req["type"] = v
	}
	if v := f.EnableStackdriverLogging; !dcl.IsEmptyValueIndirect(v) {
		req["enableStackdriverLogging"] = v
	}
	if v := f.EnableStackdriverMonitoring; !dcl.IsEmptyValueIndirect(v) {
		req["enableStackdriverMonitoring"] = v
	}
	if v, err := expandInstanceNetworkConfig(c, f.NetworkConfig); err != nil {
		return nil, fmt.Errorf("error expanding NetworkConfig into networkConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["networkConfig"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Options; !dcl.IsEmptyValueIndirect(v) {
		req["options"] = v
	}
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		req["zone"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		req["version"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.DataprocServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		req["dataprocServiceAccount"] = v
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
	err = o.Wait(ctx, c.Config, "https://datafusion.googleapis.com/v1beta1/", "GET")

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
	if err := o.Wait(ctx, c.Config, "https://datafusion.googleapis.com/v1beta1/", "GET"); err != nil {
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
	if err := o.Wait(ctx, c.Config, "https://datafusion.googleapis.com/v1beta1/", "GET"); err != nil {
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
		rawDesired.NetworkConfig = canonicalizeInstanceNetworkConfig(rawDesired.NetworkConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Instance{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Type) {
		canonicalDesired.Type = rawInitial.Type
	} else {
		canonicalDesired.Type = rawDesired.Type
	}
	if dcl.BoolCanonicalize(rawDesired.EnableStackdriverLogging, rawInitial.EnableStackdriverLogging) {
		canonicalDesired.EnableStackdriverLogging = rawInitial.EnableStackdriverLogging
	} else {
		canonicalDesired.EnableStackdriverLogging = rawDesired.EnableStackdriverLogging
	}
	if dcl.BoolCanonicalize(rawDesired.EnableStackdriverMonitoring, rawInitial.EnableStackdriverMonitoring) {
		canonicalDesired.EnableStackdriverMonitoring = rawInitial.EnableStackdriverMonitoring
	} else {
		canonicalDesired.EnableStackdriverMonitoring = rawDesired.EnableStackdriverMonitoring
	}
	if dcl.BoolCanonicalize(rawDesired.PrivateInstance, rawInitial.PrivateInstance) {
		canonicalDesired.PrivateInstance = rawInitial.PrivateInstance
	} else {
		canonicalDesired.PrivateInstance = rawDesired.PrivateInstance
	}
	canonicalDesired.NetworkConfig = canonicalizeInstanceNetworkConfig(rawDesired.NetworkConfig, rawInitial.NetworkConfig, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.IsZeroValue(rawDesired.Options) {
		canonicalDesired.Options = rawInitial.Options
	} else {
		canonicalDesired.Options = rawDesired.Options
	}
	if dcl.StringCanonicalize(rawDesired.Zone, rawInitial.Zone) {
		canonicalDesired.Zone = rawInitial.Zone
	} else {
		canonicalDesired.Zone = rawDesired.Zone
	}
	if dcl.StringCanonicalize(rawDesired.Version, rawInitial.Version) {
		canonicalDesired.Version = rawInitial.Version
	} else {
		canonicalDesired.Version = rawDesired.Version
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.NameToSelfLink(rawDesired.DataprocServiceAccount, rawInitial.DataprocServiceAccount) {
		canonicalDesired.DataprocServiceAccount = rawInitial.DataprocServiceAccount
	} else {
		canonicalDesired.DataprocServiceAccount = rawDesired.DataprocServiceAccount
	}
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

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableStackdriverLogging) && dcl.IsEmptyValueIndirect(rawDesired.EnableStackdriverLogging) {
		rawNew.EnableStackdriverLogging = rawDesired.EnableStackdriverLogging
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableStackdriverLogging, rawNew.EnableStackdriverLogging) {
			rawNew.EnableStackdriverLogging = rawDesired.EnableStackdriverLogging
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableStackdriverMonitoring) && dcl.IsEmptyValueIndirect(rawDesired.EnableStackdriverMonitoring) {
		rawNew.EnableStackdriverMonitoring = rawDesired.EnableStackdriverMonitoring
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableStackdriverMonitoring, rawNew.EnableStackdriverMonitoring) {
			rawNew.EnableStackdriverMonitoring = rawDesired.EnableStackdriverMonitoring
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PrivateInstance) && dcl.IsEmptyValueIndirect(rawDesired.PrivateInstance) {
		rawNew.PrivateInstance = rawDesired.PrivateInstance
	} else {
		if dcl.BoolCanonicalize(rawDesired.PrivateInstance, rawNew.PrivateInstance) {
			rawNew.PrivateInstance = rawDesired.PrivateInstance
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NetworkConfig) && dcl.IsEmptyValueIndirect(rawDesired.NetworkConfig) {
		rawNew.NetworkConfig = rawDesired.NetworkConfig
	} else {
		rawNew.NetworkConfig = canonicalizeNewInstanceNetworkConfig(c, rawDesired.NetworkConfig, rawNew.NetworkConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Options) && dcl.IsEmptyValueIndirect(rawDesired.Options) {
		rawNew.Options = rawDesired.Options
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StateMessage) && dcl.IsEmptyValueIndirect(rawDesired.StateMessage) {
		rawNew.StateMessage = rawDesired.StateMessage
	} else {
		if dcl.StringCanonicalize(rawDesired.StateMessage, rawNew.StateMessage) {
			rawNew.StateMessage = rawDesired.StateMessage
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceEndpoint) && dcl.IsEmptyValueIndirect(rawDesired.ServiceEndpoint) {
		rawNew.ServiceEndpoint = rawDesired.ServiceEndpoint
	} else {
		if dcl.StringCanonicalize(rawDesired.ServiceEndpoint, rawNew.ServiceEndpoint) {
			rawNew.ServiceEndpoint = rawDesired.ServiceEndpoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Zone) && dcl.IsEmptyValueIndirect(rawDesired.Zone) {
		rawNew.Zone = rawDesired.Zone
	} else {
		if dcl.StringCanonicalize(rawDesired.Zone, rawNew.Zone) {
			rawNew.Zone = rawDesired.Zone
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Version) && dcl.IsEmptyValueIndirect(rawDesired.Version) {
		rawNew.Version = rawDesired.Version
	} else {
		if dcl.StringCanonicalize(rawDesired.Version, rawNew.Version) {
			rawNew.Version = rawDesired.Version
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AvailableVersion) && dcl.IsEmptyValueIndirect(rawDesired.AvailableVersion) {
		rawNew.AvailableVersion = rawDesired.AvailableVersion
	} else {
		rawNew.AvailableVersion = canonicalizeNewInstanceAvailableVersionSlice(c, rawDesired.AvailableVersion, rawNew.AvailableVersion)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ApiEndpoint) && dcl.IsEmptyValueIndirect(rawDesired.ApiEndpoint) {
		rawNew.ApiEndpoint = rawDesired.ApiEndpoint
	} else {
		if dcl.StringCanonicalize(rawDesired.ApiEndpoint, rawNew.ApiEndpoint) {
			rawNew.ApiEndpoint = rawDesired.ApiEndpoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GcsBucket) && dcl.IsEmptyValueIndirect(rawDesired.GcsBucket) {
		rawNew.GcsBucket = rawDesired.GcsBucket
	} else {
		if dcl.StringCanonicalize(rawDesired.GcsBucket, rawNew.GcsBucket) {
			rawNew.GcsBucket = rawDesired.GcsBucket
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.P4ServiceAccount) && dcl.IsEmptyValueIndirect(rawDesired.P4ServiceAccount) {
		rawNew.P4ServiceAccount = rawDesired.P4ServiceAccount
	} else {
		if dcl.StringCanonicalize(rawDesired.P4ServiceAccount, rawNew.P4ServiceAccount) {
			rawNew.P4ServiceAccount = rawDesired.P4ServiceAccount
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TenantProjectId) && dcl.IsEmptyValueIndirect(rawDesired.TenantProjectId) {
		rawNew.TenantProjectId = rawDesired.TenantProjectId
	} else {
		if dcl.NameToSelfLink(rawDesired.TenantProjectId, rawNew.TenantProjectId) {
			rawNew.TenantProjectId = rawDesired.TenantProjectId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DataprocServiceAccount) && dcl.IsEmptyValueIndirect(rawDesired.DataprocServiceAccount) {
		rawNew.DataprocServiceAccount = rawDesired.DataprocServiceAccount
	} else {
		if dcl.NameToSelfLink(rawDesired.DataprocServiceAccount, rawNew.DataprocServiceAccount) {
			rawNew.DataprocServiceAccount = rawDesired.DataprocServiceAccount
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeInstanceNetworkConfig(des, initial *InstanceNetworkConfig, opts ...dcl.ApplyOption) *InstanceNetworkConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceNetworkConfig{}

	if dcl.NameToSelfLink(des.Network, initial.Network) || dcl.IsZeroValue(des.Network) {
		cDes.Network = initial.Network
	} else {
		cDes.Network = des.Network
	}
	if dcl.StringCanonicalize(des.IPAllocation, initial.IPAllocation) || dcl.IsZeroValue(des.IPAllocation) {
		cDes.IPAllocation = initial.IPAllocation
	} else {
		cDes.IPAllocation = des.IPAllocation
	}

	return cDes
}

func canonicalizeNewInstanceNetworkConfig(c *Client, des, nw *InstanceNetworkConfig) *InstanceNetworkConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Network, nw.Network) {
		nw.Network = des.Network
	}
	if dcl.StringCanonicalize(des.IPAllocation, nw.IPAllocation) {
		nw.IPAllocation = des.IPAllocation
	}

	return nw
}

func canonicalizeNewInstanceNetworkConfigSet(c *Client, des, nw []InstanceNetworkConfig) []InstanceNetworkConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceNetworkConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceNetworkConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceNetworkConfigSlice(c *Client, des, nw []InstanceNetworkConfig) []InstanceNetworkConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceNetworkConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceNetworkConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceAvailableVersion(des, initial *InstanceAvailableVersion, opts ...dcl.ApplyOption) *InstanceAvailableVersion {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceAvailableVersion{}

	return cDes
}

func canonicalizeNewInstanceAvailableVersion(c *Client, des, nw *InstanceAvailableVersion) *InstanceAvailableVersion {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.VersionNumber, nw.VersionNumber) {
		nw.VersionNumber = des.VersionNumber
	}
	if dcl.BoolCanonicalize(des.DefaultVersion, nw.DefaultVersion) {
		nw.DefaultVersion = des.DefaultVersion
	}
	if dcl.IsZeroValue(nw.AvailableFeatures) {
		nw.AvailableFeatures = des.AvailableFeatures
	}

	return nw
}

func canonicalizeNewInstanceAvailableVersionSet(c *Client, des, nw []InstanceAvailableVersion) []InstanceAvailableVersion {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceAvailableVersion
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceAvailableVersionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceAvailableVersionSlice(c *Client, des, nw []InstanceAvailableVersion) []InstanceAvailableVersion {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceAvailableVersion
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceAvailableVersion(c, &d, &n))
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableStackdriverLogging, actual.EnableStackdriverLogging, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("EnableStackdriverLogging")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableStackdriverMonitoring, actual.EnableStackdriverMonitoring, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("EnableStackdriverMonitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateInstance, actual.PrivateInstance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrivateInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NetworkConfig, actual.NetworkConfig, dcl.Info{ObjectFunction: compareInstanceNetworkConfigNewStyle, EmptyObject: EmptyInstanceNetworkConfig, OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("NetworkConfig")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Options, actual.Options, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Options")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.StateMessage, actual.StateMessage, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StateMessage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceEndpoint, actual.ServiceEndpoint, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Zone")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.AvailableVersion, actual.AvailableVersion, dcl.Info{OutputOnly: true, ObjectFunction: compareInstanceAvailableVersionNewStyle, EmptyObject: EmptyInstanceAvailableVersion, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AvailableVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ApiEndpoint, actual.ApiEndpoint, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ApiEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GcsBucket, actual.GcsBucket, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GcsBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.P4ServiceAccount, actual.P4ServiceAccount, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("P4ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TenantProjectId, actual.TenantProjectId, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TenantProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DataprocServiceAccount, actual.DataprocServiceAccount, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("DataprocServiceAccount")); len(ds) != 0 || err != nil {
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
func compareInstanceNetworkConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceNetworkConfig)
	if !ok {
		desiredNotPointer, ok := d.(InstanceNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceNetworkConfig or *InstanceNetworkConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceNetworkConfig)
	if !ok {
		actualNotPointer, ok := a.(InstanceNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceNetworkConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPAllocation, actual.IPAllocation, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpAllocation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceAvailableVersionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceAvailableVersion)
	if !ok {
		desiredNotPointer, ok := d.(InstanceAvailableVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceAvailableVersion or *InstanceAvailableVersion", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceAvailableVersion)
	if !ok {
		actualNotPointer, ok := a.(InstanceAvailableVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceAvailableVersion", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.VersionNumber, actual.VersionNumber, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionNumber")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultVersion, actual.DefaultVersion, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AvailableFeatures, actual.AvailableFeatures, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AvailableFeatures")); len(ds) != 0 || err != nil {
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
		return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", "https://datafusion.googleapis.com/v1beta1/", userBasePath, fields), nil

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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.EnableStackdriverLogging; !dcl.IsEmptyValueIndirect(v) {
		m["enableStackdriverLogging"] = v
	}
	if v := f.EnableStackdriverMonitoring; !dcl.IsEmptyValueIndirect(v) {
		m["enableStackdriverMonitoring"] = v
	}
	if v := f.PrivateInstance; !dcl.IsEmptyValueIndirect(v) {
		m["privateInstance"] = v
	}
	if v, err := expandInstanceNetworkConfig(c, f.NetworkConfig); err != nil {
		return nil, fmt.Errorf("error expanding NetworkConfig into networkConfig: %w", err)
	} else if v != nil {
		m["networkConfig"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Options; !dcl.IsEmptyValueIndirect(v) {
		m["options"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.StateMessage; !dcl.IsEmptyValueIndirect(v) {
		m["stateMessage"] = v
	}
	if v := f.ServiceEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["serviceEndpoint"] = v
	}
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v, err := expandInstanceAvailableVersionSlice(c, f.AvailableVersion); err != nil {
		return nil, fmt.Errorf("error expanding AvailableVersion into availableVersion: %w", err)
	} else {
		m["availableVersion"] = v
	}
	if v := f.ApiEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["apiEndpoint"] = v
	}
	if v := f.GcsBucket; !dcl.IsEmptyValueIndirect(v) {
		m["gcsBucket"] = v
	}
	if v := f.P4ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["p4ServiceAccount"] = v
	}
	if v := f.TenantProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["tenantProjectId"] = v
	}
	if v := f.DataprocServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["dataprocServiceAccount"] = v
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
	res.Description = dcl.FlattenString(m["description"])
	res.Type = flattenInstanceTypeEnum(m["type"])
	res.EnableStackdriverLogging = dcl.FlattenBool(m["enableStackdriverLogging"])
	res.EnableStackdriverMonitoring = dcl.FlattenBool(m["enableStackdriverMonitoring"])
	res.PrivateInstance = dcl.FlattenBool(m["privateInstance"])
	res.NetworkConfig = flattenInstanceNetworkConfig(c, m["networkConfig"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Options = dcl.FlattenKeyValuePairs(m["options"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.State = flattenInstanceStateEnum(m["state"])
	res.StateMessage = dcl.FlattenString(m["stateMessage"])
	res.ServiceEndpoint = dcl.FlattenString(m["serviceEndpoint"])
	res.Zone = dcl.FlattenString(m["zone"])
	res.Version = dcl.FlattenString(m["version"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.AvailableVersion = flattenInstanceAvailableVersionSlice(c, m["availableVersion"])
	res.ApiEndpoint = dcl.FlattenString(m["apiEndpoint"])
	res.GcsBucket = dcl.FlattenString(m["gcsBucket"])
	res.P4ServiceAccount = dcl.FlattenString(m["p4ServiceAccount"])
	res.TenantProjectId = dcl.FlattenString(m["tenantProjectId"])
	res.DataprocServiceAccount = dcl.FlattenString(m["dataprocServiceAccount"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandInstanceNetworkConfigMap expands the contents of InstanceNetworkConfig into a JSON
// request object.
func expandInstanceNetworkConfigMap(c *Client, f map[string]InstanceNetworkConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceNetworkConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceNetworkConfigSlice expands the contents of InstanceNetworkConfig into a JSON
// request object.
func expandInstanceNetworkConfigSlice(c *Client, f []InstanceNetworkConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceNetworkConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceNetworkConfigMap flattens the contents of InstanceNetworkConfig from a JSON
// response object.
func flattenInstanceNetworkConfigMap(c *Client, i interface{}) map[string]InstanceNetworkConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceNetworkConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceNetworkConfig{}
	}

	items := make(map[string]InstanceNetworkConfig)
	for k, item := range a {
		items[k] = *flattenInstanceNetworkConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceNetworkConfigSlice flattens the contents of InstanceNetworkConfig from a JSON
// response object.
func flattenInstanceNetworkConfigSlice(c *Client, i interface{}) []InstanceNetworkConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworkConfig{}
	}

	if len(a) == 0 {
		return []InstanceNetworkConfig{}
	}

	items := make([]InstanceNetworkConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworkConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceNetworkConfig expands an instance of InstanceNetworkConfig into a JSON
// request object.
func expandInstanceNetworkConfig(c *Client, f *InstanceNetworkConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.IPAllocation; !dcl.IsEmptyValueIndirect(v) {
		m["ipAllocation"] = v
	}

	return m, nil
}

// flattenInstanceNetworkConfig flattens an instance of InstanceNetworkConfig from a JSON
// response object.
func flattenInstanceNetworkConfig(c *Client, i interface{}) *InstanceNetworkConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceNetworkConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceNetworkConfig
	}
	r.Network = dcl.FlattenString(m["network"])
	r.IPAllocation = dcl.FlattenString(m["ipAllocation"])

	return r
}

// expandInstanceAvailableVersionMap expands the contents of InstanceAvailableVersion into a JSON
// request object.
func expandInstanceAvailableVersionMap(c *Client, f map[string]InstanceAvailableVersion) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceAvailableVersion(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceAvailableVersionSlice expands the contents of InstanceAvailableVersion into a JSON
// request object.
func expandInstanceAvailableVersionSlice(c *Client, f []InstanceAvailableVersion) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceAvailableVersion(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceAvailableVersionMap flattens the contents of InstanceAvailableVersion from a JSON
// response object.
func flattenInstanceAvailableVersionMap(c *Client, i interface{}) map[string]InstanceAvailableVersion {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceAvailableVersion{}
	}

	if len(a) == 0 {
		return map[string]InstanceAvailableVersion{}
	}

	items := make(map[string]InstanceAvailableVersion)
	for k, item := range a {
		items[k] = *flattenInstanceAvailableVersion(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceAvailableVersionSlice flattens the contents of InstanceAvailableVersion from a JSON
// response object.
func flattenInstanceAvailableVersionSlice(c *Client, i interface{}) []InstanceAvailableVersion {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceAvailableVersion{}
	}

	if len(a) == 0 {
		return []InstanceAvailableVersion{}
	}

	items := make([]InstanceAvailableVersion, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceAvailableVersion(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceAvailableVersion expands an instance of InstanceAvailableVersion into a JSON
// request object.
func expandInstanceAvailableVersion(c *Client, f *InstanceAvailableVersion) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.VersionNumber; !dcl.IsEmptyValueIndirect(v) {
		m["versionNumber"] = v
	}
	if v := f.DefaultVersion; !dcl.IsEmptyValueIndirect(v) {
		m["defaultVersion"] = v
	}
	if v := f.AvailableFeatures; v != nil {
		m["availableFeatures"] = v
	}

	return m, nil
}

// flattenInstanceAvailableVersion flattens an instance of InstanceAvailableVersion from a JSON
// response object.
func flattenInstanceAvailableVersion(c *Client, i interface{}) *InstanceAvailableVersion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceAvailableVersion{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceAvailableVersion
	}
	r.VersionNumber = dcl.FlattenString(m["versionNumber"])
	r.DefaultVersion = dcl.FlattenBool(m["defaultVersion"])
	r.AvailableFeatures = dcl.FlattenStringSlice(m["availableFeatures"])

	return r
}

// flattenInstanceTypeEnumMap flattens the contents of InstanceTypeEnum from a JSON
// response object.
func flattenInstanceTypeEnumMap(c *Client, i interface{}) map[string]InstanceTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceTypeEnum{}
	}

	items := make(map[string]InstanceTypeEnum)
	for k, item := range a {
		items[k] = *flattenInstanceTypeEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceTypeEnumSlice flattens the contents of InstanceTypeEnum from a JSON
// response object.
func flattenInstanceTypeEnumSlice(c *Client, i interface{}) []InstanceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceTypeEnum{}
	}

	items := make([]InstanceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTypeEnum with the same value as that string.
func flattenInstanceTypeEnum(i interface{}) *InstanceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTypeEnumRef("")
	}

	return InstanceTypeEnumRef(s)
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
